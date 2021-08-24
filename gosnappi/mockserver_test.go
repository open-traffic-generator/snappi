package gosnappi_test

import (
	context "context"
	fmt "fmt"
	log "log"
	net "net"
	"reflect"

	snappipb "github.com/open-traffic-generator/snappi/gosnappi/snappipb"
	grpc "google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	testport   uint             = 50001
	testserver *grpc.Server     = nil
	config     *snappipb.Config = nil
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

func StartMockServer() error {
	if testserver != nil {
		log.Print("MockServer: Server already running")
		return nil
	}

	addr := fmt.Sprintf("127.0.0.1:%d", testport)

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal(fmt.Sprintf("MockServer: Server failed to listen on address %s", addr))
	}
	svr := grpc.NewServer()
	log.Print(fmt.Sprintf("MockServer: started and listening on address %s", addr))
	snappipb.RegisterOpenapiServer(svr, &server{})
	log.Print("MockServer: Server subscribed with gRPC Protocol Service")

	go func() {
		if err := svr.Serve(lis); err != nil {
			log.Fatal("MockServer: Server failed to serve for incoming gRPC request.")

		}
	}()
	testserver = svr
	return nil
}

func (s *server) SetConfig(ctx context.Context, req *snappipb.SetConfigRequest) (*snappipb.SetConfigResponse, error) {
	var resp *snappipb.SetConfigResponse
	if reflect.TypeOf(req.Config) != reflect.TypeOf(config) {
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
		config = req.Config
		resp = &snappipb.SetConfigResponse{
			StatusCode_200: &snappipb.SetConfigResponse_StatusCode200{
				Success: &snappipb.Success{},
			},
		}
	}
	fmt.Println(req.Config)
	return resp, nil
}

func (s *server) GetConfig(ctx context.Context, in *emptypb.Empty) (*snappipb.GetConfigResponse, error) {
	resp := &snappipb.GetConfigResponse{
		StatusCode_200: &snappipb.GetConfigResponse_StatusCode200{
			Config: config,
		},
	}
	return resp, nil
}

func (s *server) GetMetrics(ctx context.Context, req *snappipb.GetMetricsRequest) (*snappipb.GetMetricsResponse, error) {
	var resp *snappipb.GetMetricsResponse
	var tx int32 = 100
	f := &snappipb.FlowMetric{FramesTx: &tx}
	resp = &snappipb.GetMetricsResponse{
		StatusCode_200: &snappipb.GetMetricsResponse_StatusCode200{
			MetricsResponse: &snappipb.MetricsResponse{
				FlowMetrics: []*snappipb.FlowMetric{f},
			},
		},
	}

	flowNames := []string{}
	for _, flow := range config.Flows {
		flowNames = append(flowNames, flow.Name)
	}

	for _, req_flow := range req.MetricsRequest.Flow.FlowNames {
		res := contains(flowNames, req_flow)
		if res == false {
			resp = &snappipb.GetMetricsResponse{
				StatusCode_400: &snappipb.GetMetricsResponse_StatusCode400{
					BadRequest: &snappipb.BadRequest{
						ResponseError: &snappipb.ResponseError{
							Errors: []string{"requested flow in not available in configured flows"},
						},
					},
				},
			}
		}
	}
	return resp, nil
}
