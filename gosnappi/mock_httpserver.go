package gosnappi

import (
	"io/ioutil"
	"log"
	"net/http"
)

type HttpServer struct {
	serverLocation string
	Api            GosnappiApi
	Config         Config
}

var (
	httpServer HttpServer = HttpServer{}
)

func StartMockHttpServer(location string) {
	httpServer.serverLocation = location
	httpServer.Api = NewApi()
	httpServer.Config = httpServer.Api.NewConfig()

	http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			body, _ := ioutil.ReadAll(r.Body)
			httpServer.Config.FromJson(string(body))
			response := httpServer.Api.NewSetConfigResponse_StatusCode200()
			response.Success()
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response.ToJson()))
		case http.MethodGet:
			config := httpServer.Config
			response := httpServer.Api.NewGetConfigResponse_StatusCode200()
			response.Config().FromJson(config.ToJson())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response.ToJson()))
		}
	})

	http.HandleFunc("/results/metrics", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			body, _ := ioutil.ReadAll(r.Body)
			metricsReq := httpServer.Api.NewMetricsRequest()
			metricsReq.FromJson(string(body))
			if metricsReq.Choice() == MetricsRequestChoice.FLOW {
				flow_responses := httpServer.Api.NewGetMetricsResponse_StatusCode200().MetricsResponse()
				for _, flow_name := range metricsReq.Flow().FlowNames() {
					flow_rsp := flow_responses.FlowMetrics().Add()
					flow_rsp.SetName(flow_name)
					flow_rsp.SetBytesTx(1000)
					flow_rsp.SetBytesRx(1000)
				}
				response := httpServer.Api.NewGetMetricsResponse_StatusCode200().MetricsResponse()
				response.FromJson(flow_responses.ToJson())
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(response.ToJson()))
			}

		}
	})

	go func() {
		if err := http.ListenAndServe(httpServer.serverLocation, nil); err != nil {
			log.Fatal("Server failed to serve incoming HTTP request.")
		}
	}()
}
