package gosnappi

import (
	context "context"
	"fmt"
	log "log"
	net "net"
	"reflect"

	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	grpc "google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	mockServer *grpc.Server = nil
	mockConfig *otg.Config  = nil
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

type server struct {
	otg.UnimplementedOpenapiServer
}

func StartMockGrpcServer(location string) error {
	if mockServer != nil {
		log.Println("MockServer: Server already running")
		return nil
	}

	lis, err := net.Listen("tcp", location)

	if err != nil {
		log.Fatal("MockServer: Server failed to listen on address " + location)
	}
	svr := grpc.NewServer()
	log.Println("MockServer: started and listening on address " + location)
	otg.RegisterOpenapiServer(svr, &server{})
	log.Println("MockServer: Server subscribed with gRPC Protocol Service")

	go func() {
		if err := svr.Serve(lis); err != nil {
			log.Fatal("MockServer: Server failed to serve for incoming gRPC request.")

		}
	}()
	mockServer = svr
	return nil
}

func (s *server) SetConfig(ctx context.Context, req *otg.SetConfigRequest) (*otg.SetConfigResponse, error) {
	var resp *otg.SetConfigResponse
	var err error
	if reflect.TypeOf(req.Config) != reflect.TypeOf(mockConfig) {
		resp = nil
		err = fmt.Errorf("config did not match")
	} else {
		mockConfig = req.Config
		resp = &otg.SetConfigResponse{
			Warning: &otg.Warning{
				Warnings: []string{},
			},
		}
		err = nil
	}
	log.Printf("Got config: %v\n", req.Config)
	return resp, err
}

func (s *server) GetConfig(ctx context.Context, in *emptypb.Empty) (*otg.GetConfigResponse, error) {
	resp := &otg.GetConfigResponse{
		Config: mockConfig,
	}
	return resp, nil
}

func getBgpPeerNames(cfg *otg.Config) []string {
	names := []string{}
	for _, d := range cfg.Devices {
		if d == nil || d.Bgp == nil {
			continue
		}
		names = append(names, d.Bgp.Ipv4Interfaces[0].Peers[0].GetName())
	}

	return names
}

func getFlowNames(cfg *otg.Config) []string {
	names := []string{}
	for _, flow := range mockConfig.Flows {
		names = append(names, flow.Name)
	}

	return names
}

func getPortNames(cfg *otg.Config) []string {
	names := []string{}
	for _, port := range mockConfig.Ports {
		names = append(names, port.Name)
	}

	return names
}

func getRouteNames(cfg *otg.Config) []string {
	names := []string{}
	for _, d := range cfg.Devices {
		if d == nil || d.Bgp == nil {
			continue
		}
		for _, r := range d.Bgp.Ipv4Interfaces[0].Peers[0].GetV4Routes() {
			names = append(names, r.Name)
		}
	}
	return names
}

func isFlowMetricsDisabled(cfg *otg.Config) []string {
	names := []string{}
	for _, flow := range cfg.Flows {
		if flow.Metrics == nil || !*flow.Metrics.Enable {
			names = append(names, flow.Name)
		}
	}
	return names
}

func (s *server) GetMetrics(ctx context.Context, req *otg.GetMetricsRequest) (*otg.GetMetricsResponse, error) {
	var resp *otg.GetMetricsResponse
	var err error
	var tx int64 = 100
	metricsDisabledFlows := isFlowMetricsDisabled(mockConfig)
	if req.MetricsRequest.Flow != nil {
		f := &otg.FlowMetric{FramesTx: &tx}
		flowNames := getFlowNames(mockConfig)
		for _, req_flow := range req.MetricsRequest.Flow.FlowNames {
			res := contains(flowNames, req_flow)
			if !res {
				resp = nil
				errObj := NewError()
				errObj.Msg().Code = 13
				errObj.Msg().Errors = []string{"requested flow is not available in configured flows"}
				err = errObj
			} else if len(metricsDisabledFlows) > 0 {
				resp = nil
				errObj := NewError()
				errObj.Msg().Code = 13
				errObj.Msg().Errors = []string{"metrics not enabled for all the flows"}
				err = errObj
			} else {
				resp = &otg.GetMetricsResponse{
					MetricsResponse: &otg.MetricsResponse{
						FlowMetrics: []*otg.FlowMetric{f},
					},
				}
				err = nil
			}
		}
	} else if req.MetricsRequest.Port != nil {
		p := &otg.PortMetric{FramesTx: &tx}
		portNames := getPortNames(mockConfig)
		for _, req_port := range req.MetricsRequest.Port.PortNames {
			res := contains(portNames, req_port)
			if !res {
				resp = nil
				resp = nil
				errObj := NewError()
				errObj.Msg().Code = 13
				errObj.Msg().Errors = []string{"requested port is not available in configured ports"}
				err = errObj
			} else {
				resp = &otg.GetMetricsResponse{
					MetricsResponse: &otg.MetricsResponse{
						PortMetrics: []*otg.PortMetric{p},
					},
				}
				err = nil
			}
		}
	} else if req.MetricsRequest.Bgpv4 != nil {
		allNames := getBgpPeerNames(mockConfig)
		someNames := req.MetricsRequest.Bgpv4.PeerNames
		if len(someNames) == 0 {
			someNames = allNames
		} else {
			for _, name := range someNames {
				if !contains(allNames, name) {
					return nil, fmt.Errorf("%s is not a valid BGPv4 device", name)
				}
			}
		}

		metrics := []*otg.Bgpv4Metric{}

		for _, name := range someNames {
			one := int32(1)
			zero := int32(0)
			up := otg.Bgpv4Metric_SessionState_up
			metrics = append(metrics, &otg.Bgpv4Metric{
				Name:               &name,
				SessionState:       &up,
				RoutesAdvertised:   &one,
				RouteWithdrawsSent: &zero,
			})
		}

		resp = &otg.GetMetricsResponse{
			MetricsResponse: &otg.MetricsResponse{
				Bgpv4Metrics: metrics,
			},
		}
		err = nil
	}
	return resp, err
}

func (s *server) SetTransmitState(ctx context.Context, req *otg.SetTransmitStateRequest) (*otg.SetTransmitStateResponse, error) {
	var resp *otg.SetTransmitStateResponse
	var err error
	flowNames := getFlowNames(mockConfig)
	for _, req_flow := range req.TransmitState.FlowNames {
		res := contains(flowNames, req_flow)
		if !res {
			resp = nil
			err = fmt.Errorf("requested flow is not available in configured flows to start")
		} else {
			resp = &otg.SetTransmitStateResponse{
				Warning: &otg.Warning{
					Warnings: []string{},
				},
			}
			err = nil
		}
	}
	return resp, err
}

func (s *server) SetLinkState(ctx context.Context, req *otg.SetLinkStateRequest) (*otg.SetLinkStateResponse, error) {
	var resp *otg.SetLinkStateResponse
	var err error
	portNames := getPortNames(mockConfig)
	for _, req_port := range req.LinkState.PortNames {
		res := contains(portNames, req_port)
		if !res {
			resp = nil
			err = fmt.Errorf("requested port is not available in configured ports to do link down")
		} else {
			resp = &otg.SetLinkStateResponse{
				Warning: &otg.Warning{
					Warnings: []string{},
				},
			}
			err = nil
		}
	}
	return resp, err
}

func (s *server) SetCaptureState(ctx context.Context, req *otg.SetCaptureStateRequest) (*otg.SetCaptureStateResponse, error) {
	var resp *otg.SetCaptureStateResponse
	var err error
	portNames := getPortNames(mockConfig)
	for _, req_port := range req.CaptureState.PortNames {
		res := contains(portNames, req_port)
		if !res {
			resp = nil
			err = fmt.Errorf("requested port is not available in configured ports to start capture")
		} else {
			resp = &otg.SetCaptureStateResponse{
				Warning: &otg.Warning{
					Warnings: []string{},
				},
			}
			err = nil
		}
	}
	return resp, err
}

func (s *server) SetRouteState(ctx context.Context, req *otg.SetRouteStateRequest) (*otg.SetRouteStateResponse, error) {
	var resp *otg.SetRouteStateResponse
	var err error
	routeNames := getRouteNames(mockConfig)
	for _, req_route := range req.RouteState.Names {
		res := contains(routeNames, req_route)
		if !res {
			resp = nil
			err = fmt.Errorf("requested route is not available in configured routes to advertise")
		} else {
			resp = &otg.SetRouteStateResponse{
				Warning: &otg.Warning{
					Warnings: []string{},
				},
			}
			err = nil
		}
	}
	return resp, err
}
