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
			StatusCode_400: &snappipb.SetConfigResponse_StatusCode400{
				BadRequest: &snappipb.BadRequest{
					ResponseError: &snappipb.ResponseError{
						Errors: []string{"Invalid Config object received"},
					},
				},
			},
		}
	} else {
		mockConfig = req.Config
		resp = &snappipb.SetConfigResponse{
			StatusCode_200: &snappipb.SetConfigResponse_StatusCode200{
				Success: &snappipb.Success{},
			},
		}
	}
	log.Printf("Got config: %v\n", req.Config)
	return resp, nil
}

func (s *server) GetConfig(ctx context.Context, in *emptypb.Empty) (*snappipb.GetConfigResponse, error) {
	resp := &snappipb.GetConfigResponse{
		StatusCode_200: &snappipb.GetConfigResponse_StatusCode200{
			Config: mockConfig,
		},
	}
	return resp, nil
}

func getBgpDeviceNames(cfg *snappipb.Config) []string {
	names := []string{}
	for _, d := range cfg.Devices {
		if d == nil || d.Ethernet == nil || d.Ethernet.Ipv4 == nil || d.Ethernet.Ipv4.Bgpv4 == nil {
			continue
		}
		names = append(names, d.Ethernet.Ipv4.Bgpv4.Name)
	}

	return names
}

func (s *server) GetMetrics(ctx context.Context, req *snappipb.GetMetricsRequest) (*snappipb.GetMetricsResponse, error) {
	var resp *snappipb.GetMetricsResponse
	var tx int32 = 100
	if req.MetricsRequest.Flow != nil {
		f := &snappipb.FlowMetric{FramesTx: &tx}
		flowNames := []string{}
		for _, flow := range mockConfig.Flows {
			flowNames = append(flowNames, flow.Name)
		}
		for _, req_flow := range req.MetricsRequest.Flow.FlowNames {
			res := contains(flowNames, req_flow)
			if res == false {
				resp = &snappipb.GetMetricsResponse{
					StatusCode_400: &snappipb.GetMetricsResponse_StatusCode400{
						BadRequest: &snappipb.BadRequest{
							ResponseError: &snappipb.ResponseError{
								Errors: []string{"requested flow is not available in configured flows"},
							},
						},
					},
				}
			} else {
				resp = &snappipb.GetMetricsResponse{
					StatusCode_200: &snappipb.GetMetricsResponse_StatusCode200{
						MetricsResponse: &snappipb.MetricsResponse{
							FlowMetrics: []*snappipb.FlowMetric{f},
						},
					},
				}
			}
		}
	} else if req.MetricsRequest.Port != nil {
		p := &snappipb.PortMetric{FramesTx: &tx}
		portNames := []string{}
		for _, port := range mockConfig.Ports {
			portNames = append(portNames, port.Name)
		}
		for _, req_port := range req.MetricsRequest.Port.PortNames {
			res := contains(portNames, req_port)
			if res == false {
				resp = &snappipb.GetMetricsResponse{
					StatusCode_400: &snappipb.GetMetricsResponse_StatusCode400{
						BadRequest: &snappipb.BadRequest{
							ResponseError: &snappipb.ResponseError{
								Errors: []string{"requested port is not available in configured ports"},
							},
						},
					},
				}
			} else {
				resp = &snappipb.GetMetricsResponse{
					StatusCode_200: &snappipb.GetMetricsResponse_StatusCode200{
						MetricsResponse: &snappipb.MetricsResponse{
							PortMetrics: []*snappipb.PortMetric{p},
						},
					},
				}
			}
		}
	} else if req.MetricsRequest.Bgpv4 != nil {
		allNames := getBgpDeviceNames(mockConfig)
		someNames := req.MetricsRequest.Bgpv4.DeviceNames
		if len(someNames) == 0 {
			someNames = allNames
		} else {
			for _, name := range someNames {
				if !contains(allNames, name) {
					return &snappipb.GetMetricsResponse{
						StatusCode_400: &snappipb.GetMetricsResponse_StatusCode400{
							BadRequest: &snappipb.BadRequest{
								ResponseError: &snappipb.ResponseError{
									Errors: []string{name + " is not a valid BGPv4 device"},
								},
							},
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
				Name:             &name,
				SessionState:     &up,
				RoutesAdvertised: &one,
				RoutesWithdrawn:  &zero,
			})
		}

		resp = &snappipb.GetMetricsResponse{
			StatusCode_200: &snappipb.GetMetricsResponse_StatusCode200{
				MetricsResponse: &snappipb.MetricsResponse{
					Bgpv4Metrics: metrics,
				},
			},
		}
	}
	return resp, nil
}
