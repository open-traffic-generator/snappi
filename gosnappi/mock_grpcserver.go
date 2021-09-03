package gosnappi

import (
	context "context"
	log "log"
	net "net"
	"reflect"

	snappipb "github.com/open-traffic-generator/snappi/gosnappi/snappipb"
	grpc "google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	mockServer *grpc.Server     = nil
	mockConfig *snappipb.Config = nil
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
	snappipb.UnimplementedOpenapiServer
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
	snappipb.RegisterOpenapiServer(svr, &server{})
	log.Println("MockServer: Server subscribed with gRPC Protocol Service")

	go func() {
		if err := svr.Serve(lis); err != nil {
			log.Fatal("MockServer: Server failed to serve for incoming gRPC request.")

		}
	}()
	mockServer = svr
	return nil
}

func (s *server) SetConfig(ctx context.Context, req *snappipb.SetConfigRequest) (*snappipb.SetConfigResponse, error) {
	var resp *snappipb.SetConfigResponse
	if reflect.TypeOf(req.Config) != reflect.TypeOf(mockConfig) {
		resp = &snappipb.SetConfigResponse{
			StatusCode_400: &snappipb.ResponseError{
				Errors: []string{"Invalid Config object received"},
			},
		}
	} else {
		mockConfig = req.Config
		resp = &snappipb.SetConfigResponse{
			StatusCode_200: &snappipb.ResponseWarning{
				Warnings: []string{},
			},
		}
	}
	log.Printf("Got config: %v\n", req.Config)
	return resp, nil
}

func (s *server) GetConfig(ctx context.Context, in *emptypb.Empty) (*snappipb.GetConfigResponse, error) {
	resp := &snappipb.GetConfigResponse{
		StatusCode_200: mockConfig,
	}
	return resp, nil
}

func getBgpPeerNames(cfg *snappipb.Config) []string {
	names := []string{}
	for _, d := range cfg.Devices {
		if d == nil || d.Ethernet == nil || d.Ethernet.Ipv4 == nil || d.Ethernet.Ipv4.Bgpv4 == nil {
			continue
		}
		names = append(names, d.Ethernet.Ipv4.Bgpv4.Name)
	}

	return names
}

func getFlowNames(cfg *snappipb.Config) []string {
	names := []string{}
	for _, flow := range mockConfig.Flows {
		names = append(names, flow.Name)
	}

	return names
}

func getPortNames(cfg *snappipb.Config) []string {
	names := []string{}
	for _, port := range mockConfig.Ports {
		names = append(names, port.Name)
	}

	return names
}

func getRouteNames(cfg *snappipb.Config) []string {
	names := []string{}
	for _, d := range cfg.Devices {
		if d == nil || d.Ethernet == nil || d.Ethernet.Ipv4 == nil || d.Ethernet.Ipv4.Bgpv4 == nil {
			continue
		}
		for _, r := range d.Ethernet.Ipv4.Bgpv4.Bgpv4Routes {
			names = append(names, r.Name)
		}
	}
	return names
}

func isFlowMetricsDisabled(cfg *snappipb.Config) []string {
	names := []string{}
	for _, flow := range mockConfig.Flows {
		if flow.Metrics == nil {
			names = append(names, flow.Name)
		}
	}
	return names
}

func (s *server) GetMetrics(ctx context.Context, req *snappipb.GetMetricsRequest) (*snappipb.GetMetricsResponse, error) {
	var resp *snappipb.GetMetricsResponse
	var tx int32 = 100
	metricsDisabledFlows := isFlowMetricsDisabled(mockConfig)
	if req.MetricsRequest.Flow != nil {
		f := &snappipb.FlowMetric{FramesTx: &tx}
		flowNames := getFlowNames(mockConfig)
		for _, req_flow := range req.MetricsRequest.Flow.FlowNames {
			res := contains(flowNames, req_flow)
			if !res {
				resp = &snappipb.GetMetricsResponse{
					StatusCode_500: &snappipb.ResponseError{
						Errors: []string{"requested flow is not available in configured flows"},
					},
				}
			} else if len(metricsDisabledFlows) > 0 {
				resp = &snappipb.GetMetricsResponse{
					StatusCode_400: &snappipb.ResponseError{
						Errors: []string{"metrics not enabled for all the flows"},
					},
				}
			} else {
				resp = &snappipb.GetMetricsResponse{
					StatusCode_200: &snappipb.MetricsResponse{
						FlowMetrics: []*snappipb.FlowMetric{f},
					},
				}
			}
		}
	} else if req.MetricsRequest.Port != nil {
		p := &snappipb.PortMetric{FramesTx: &tx}
		portNames := getPortNames(mockConfig)
		for _, req_port := range req.MetricsRequest.Port.PortNames {
			res := contains(portNames, req_port)
			if !res {
				resp = &snappipb.GetMetricsResponse{
					StatusCode_400: &snappipb.ResponseError{
						Errors: []string{"requested port is not available in configured ports"},
					},
				}
			} else {
				resp = &snappipb.GetMetricsResponse{
					StatusCode_200: &snappipb.MetricsResponse{
						PortMetrics: []*snappipb.PortMetric{p},
					},
				}
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
					return &snappipb.GetMetricsResponse{
						StatusCode_400: &snappipb.ResponseError{
							Errors: []string{name + " is not a valid BGPv4 device"},
						},
					}, nil
				}
			}
		}

		metrics := []*snappipb.Bgpv4Metric{}

		for _, name := range someNames {
			one := int32(1)
			zero := int32(0)
			up := snappipb.Bgpv4Metric_SessionState_up
			metrics = append(metrics, &snappipb.Bgpv4Metric{
				Name:               &name,
				SessionState:       &up,
				RoutesAdvertised:   &one,
				RouteWithdrawsSent: &zero,
			})
		}

		resp = &snappipb.GetMetricsResponse{
			StatusCode_200: &snappipb.MetricsResponse{
				Bgpv4Metrics: metrics,
			},
		}
	}
	return resp, nil
}

func (s *server) SetTransmitState(ctx context.Context, req *snappipb.SetTransmitStateRequest) (*snappipb.SetTransmitStateResponse, error) {
	var resp *snappipb.SetTransmitStateResponse
	flowNames := getFlowNames(mockConfig)
	for _, req_flow := range req.TransmitState.FlowNames {
		res := contains(flowNames, req_flow)
		if !res {
			resp = &snappipb.SetTransmitStateResponse{
				StatusCode_400: &snappipb.ResponseError{
					Errors: []string{"requested flow is not available in configured flows to start"},
				},
			}
		} else {
			resp = &snappipb.SetTransmitStateResponse{
				StatusCode_200: &snappipb.ResponseWarning{
					Warnings: []string{},
				},
			}
		}
	}
	return resp, nil
}

func (s *server) SetLinkState(ctx context.Context, req *snappipb.SetLinkStateRequest) (*snappipb.SetLinkStateResponse, error) {
	var resp *snappipb.SetLinkStateResponse
	portNames := getPortNames(mockConfig)
	for _, req_port := range req.LinkState.PortNames {
		res := contains(portNames, req_port)
		if !res {
			resp = &snappipb.SetLinkStateResponse{
				StatusCode_400: &snappipb.ResponseError{
					Errors: []string{"requested port is not available in configured ports to do link down"},
				},
			}
		} else {
			resp = &snappipb.SetLinkStateResponse{
				StatusCode_200: &snappipb.ResponseWarning{
					Warnings: []string{},
				},
			}
		}
	}
	return resp, nil
}

func (s *server) SetCaptureState(ctx context.Context, req *snappipb.SetCaptureStateRequest) (*snappipb.SetCaptureStateResponse, error) {
	var resp *snappipb.SetCaptureStateResponse
	portNames := getPortNames(mockConfig)
	for _, req_port := range req.CaptureState.PortNames {
		res := contains(portNames, req_port)
		if !res {
			resp = &snappipb.SetCaptureStateResponse{
				StatusCode_400: &snappipb.ResponseError{
					Errors: []string{"requested port is not available in configured ports to start capture"},
				},
			}
		} else {
			resp = &snappipb.SetCaptureStateResponse{
				StatusCode_200: &snappipb.ResponseWarning{
					Warnings: []string{},
				},
			}
		}
	}
	return resp, nil
}

func (s *server) SetRouteState(ctx context.Context, req *snappipb.SetRouteStateRequest) (*snappipb.SetRouteStateResponse, error) {
	var resp *snappipb.SetRouteStateResponse
	routeNames := getRouteNames(mockConfig)
	for _, req_route := range req.RouteState.Names {
		res := contains(routeNames, req_route)
		if !res {
			resp = &snappipb.SetRouteStateResponse{
				StatusCode_400: &snappipb.ResponseError{
					Errors: []string{"requested route is not available in configured routes to advertise"},
				},
			}
		} else {
			resp = &snappipb.SetRouteStateResponse{
				StatusCode_200: &snappipb.ResponseWarning{
					Warnings: []string{},
				},
			}
		}
	}
	return resp, nil
}
