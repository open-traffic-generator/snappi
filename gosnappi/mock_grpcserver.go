package gosnappi

import (
	context "context"
	"fmt"
	"io"
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
		names = append(names, *flow.Name)
	}

	return names
}

func getPortNames(cfg *otg.Config) []string {
	names := []string{}
	for _, port := range mockConfig.Ports {
		names = append(names, *port.Name)
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
			names = append(names, *r.Name)
		}
	}
	return names
}

func isFlowMetricsDisabled(cfg *otg.Config) []string {
	names := []string{}
	for _, flow := range cfg.Flows {
		if flow.Metrics == nil || !*flow.Metrics.Enable {
			names = append(names, *flow.Name)
		}
	}
	return names
}

func (s *server) GetMetrics(ctx context.Context, req *otg.GetMetricsRequest) (*otg.GetMetricsResponse, error) {
	var resp *otg.GetMetricsResponse
	var err error
	var tx uint64 = 100
	var errCode int32 = 13
	metricsDisabledFlows := isFlowMetricsDisabled(mockConfig)
	if req.MetricsRequest.Flow != nil {
		f := &otg.FlowMetric{FramesTx: &tx}
		flowNames := getFlowNames(mockConfig)
		for _, req_flow := range req.MetricsRequest.Flow.FlowNames {
			res := contains(flowNames, req_flow)
			if !res {
				resp = nil
				errObj := NewError()
				_ = errObj.SetCode(errCode)
				_ = errObj.SetErrors([]string{"requested flow is not available in configured flows"})
				err = errObj
			} else if len(metricsDisabledFlows) > 0 {
				resp = nil
				errObj := NewError()
				_ = errObj.SetCode(errCode)
				_ = errObj.SetErrors([]string{"metrics not enabled for all the flows"})
				err = errObj
			} else {
				ch := otg.MetricsResponse_Choice_Enum(int32(1))
				resp = &otg.GetMetricsResponse{
					MetricsResponse: &otg.MetricsResponse{
						Choice:      &ch,
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
				_ = errObj.SetCode(errCode)
				_ = errObj.SetErrors([]string{"requested port is not available in configured ports"})
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
			one := uint64(1)
			zero := uint64(0)
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

func (s *server) SetControlState(ctx context.Context, req *otg.SetControlStateRequest) (*otg.SetControlStateResponse, error) {
	switch choice := *req.ControlState.Choice; choice {
	case otg.ControlState_Choice_port:
		switch portChoice := *req.ControlState.Port.Choice; portChoice {
		case otg.StatePort_Choice_link:
			return s.SetLinkState(ctx, req)
		case otg.StatePort_Choice_capture:
			return s.SetCaptureState(ctx, req)
		case otg.StatePort_Choice_unspecified:
			return nil, fmt.Errorf("unspecified choice")
		}
	case otg.ControlState_Choice_protocol:
		switch protocolChoice := *req.ControlState.Protocol.Choice; protocolChoice {
		case otg.StateProtocol_Choice_route:
			return s.SetRouteState(ctx, req)
		case otg.StateProtocol_Choice_unspecified:
			return nil, fmt.Errorf("unspecified choice")
		}
	case otg.ControlState_Choice_traffic:
		return s.SetTransmitState(ctx, req)
	case otg.ControlState_Choice_unspecified:
		return nil, fmt.Errorf("unspecified choice")
	}

	return nil, fmt.Errorf("error in implementation")
}

func (s *server) SetTransmitState(ctx context.Context, req *otg.SetControlStateRequest) (*otg.SetControlStateResponse, error) {
	var resp *otg.SetControlStateResponse
	var err error
	flowNames := getFlowNames(mockConfig)
	for _, req_flow := range req.ControlState.Traffic.FlowTransmit.FlowNames {
		res := contains(flowNames, req_flow)
		if !res {
			resp = nil
			err = fmt.Errorf("requested flow is not available in configured flows to start")
		} else {
			resp = &otg.SetControlStateResponse{
				Warning: &otg.Warning{
					Warnings: []string{},
				},
			}
			err = nil
		}
	}
	return resp, err
}

func (s *server) SetLinkState(ctx context.Context, req *otg.SetControlStateRequest) (*otg.SetControlStateResponse, error) {
	var resp *otg.SetControlStateResponse
	var err error
	portNames := getPortNames(mockConfig)
	for _, req_port := range req.ControlState.Port.Link.PortNames {
		res := contains(portNames, req_port)
		if !res {
			resp = nil
			err = fmt.Errorf("requested port is not available in configured ports to do link down")
		} else {
			resp = &otg.SetControlStateResponse{
				Warning: &otg.Warning{
					Warnings: []string{},
				},
			}
			err = nil
		}
	}
	return resp, err
}

func (s *server) SetCaptureState(ctx context.Context, req *otg.SetControlStateRequest) (*otg.SetControlStateResponse, error) {
	var resp *otg.SetControlStateResponse
	var err error
	portNames := getPortNames(mockConfig)
	for _, req_port := range req.ControlState.Port.Capture.PortNames {
		res := contains(portNames, req_port)
		if !res {
			resp = nil
			err = fmt.Errorf("requested port is not available in configured ports to start capture")
		} else {
			resp = &otg.SetControlStateResponse{
				Warning: &otg.Warning{
					Warnings: []string{},
				},
			}
			err = nil
		}
	}
	return resp, err
}

func (s *server) SetRouteState(ctx context.Context, req *otg.SetControlStateRequest) (*otg.SetControlStateResponse, error) {
	var resp *otg.SetControlStateResponse
	var err error
	fmt.Println(req.ControlState.Protocol.Route.Names)
	routeNames := getRouteNames(mockConfig)
	for _, req_route := range req.ControlState.Protocol.Route.Names {
		res := contains(routeNames, req_route)
		if !res {
			resp = nil
			err = fmt.Errorf("requested route is not available in configured routes to advertise")
		} else {
			resp = &otg.SetControlStateResponse{
				Warning: &otg.Warning{
					Warnings: []string{},
				},
			}
			err = nil
		}
	}
	return resp, err
}

func (s *server) StreamSetConfig(srv otg.Openapi_StreamSetConfigServer) error {

	var resp *otg.SetConfigResponse
	var blob []byte
	for {
		data, err := srv.Recv()
		if data != nil {
			fmt.Println("Receiving chunk ")
		}
		if err != nil {
			if err == io.EOF {
				fmt.Println("transfer success")
				fmt.Printf("Transfer of %d bytes successful\n", len(blob))
				config := NewConfig()
				er := config.Unmarshal().FromPbText(string(blob))
				if er != nil {
					return er
				}
				yaml, _ := config.Marshal().ToYaml()
				fmt.Println(yaml)

				mockConfig, _ = config.Marshal().ToProto()
				resp = &otg.SetConfigResponse{
					Warning: &otg.Warning{
						Warnings: []string{},
					},
				}
				return srv.SendAndClose(resp)
			}
			return err
		}
		blob = append(blob, data.Datum...)
	}

}

func (s *server) StreamGetConfig(in *emptypb.Empty, srv otg.Openapi_StreamGetConfigServer) error {
	config := NewConfig()
	port1 := config.Ports().Add().SetName("p1")
	port2 := config.Ports().Add().SetName("p2")
	flow1 := config.Flows().Add().SetName("f1")
	flow1.TxRx().Port().SetTxName(port1.Name()).SetRxName(port2.Name())

	eth := flow1.Packet().Add().Ethernet()
	eth.Src()
	eth.Dst()
	eth.EtherType()

	ipv4 := flow1.Packet().Add().Ipv4()
	ipv4.Src()
	ipv4.Dst()
	ipv4.Version()
	ipv4.Priority()

	udp := flow1.Packet().Add().Udp()
	udp.SrcPort()
	udp.DstPort()
	udp.Checksum()

	flow1.Duration().Continuous()
	flow1.Rate()
	text, err := config.Marshal().ToPbText()
	if err != nil {
		return err
	}
	bytes := []byte(text)
	chunkSize := 50
	for i := 0; i < len(bytes); i += chunkSize {
		data := &otg.Data{}
		if i+chunkSize > len(bytes) {
			data.Datum = bytes[i:]
		} else {
			data.Datum = bytes[i : i+chunkSize]
		}
		if err := srv.Send(data); err != nil {
			fmt.Printf("Failed to send: %v\n", err)
			return err
		}
		fmt.Printf("Sent: %v\n", data)
	}

	fmt.Println("Finished streaming")
	return nil
}

func (s *server) StreamGetCapture(in *otg.GetCaptureRequest, srv otg.Openapi_StreamGetCaptureServer) error {
	fmt.Println(in)
	bytes := []byte("Hello how are you ?\nHello how are you ?\nHello how are you ?\nHello how are you ?\nHello how are you ?\nHello how are you ?\n")
	chunkSize := 10
	for i := 0; i < len(bytes); i += chunkSize {
		data := &otg.Data{}
		if i+chunkSize > len(bytes) {
			data.Datum = bytes[i:]
		} else {
			data.Datum = bytes[i : i+chunkSize]
		}
		if err := srv.Send(data); err != nil {
			fmt.Printf("Failed to send: %v\n", err)
			return err
		}
		fmt.Printf("Sent: %v\n", data)
	}

	fmt.Println("Finished streaming")
	return nil
}
