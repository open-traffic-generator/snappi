package gosnappi

import (
	"io/ioutil"
	"log"
	"net/http"
)

type HttpServer struct {
	serverLocation string
	Api            Api
	Config         Config
}

var (
	httpServer HttpServer = HttpServer{}
)

func StartMockHttpServer(location string) {
	httpServer.serverLocation = location
	httpServer.Api = NewApi()
	httpServer.Config = NewConfig()

	http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			body, _ := ioutil.ReadAll(r.Body)
			httpServer.Config.Unmarshal().FromJson(string(body))
			response := NewSetConfigResponse()
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			resp, _ := response.Warning().Marshal().ToJson()
			w.Write([]byte(resp))
		case http.MethodGet:
			// config := httpServer.Config
			response := NewGetConfigResponse()
			httpServerConfig, _ := httpServer.Config.Marshal().ToJson()
			response.Config().Unmarshal().FromJson(httpServerConfig)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			resp, _ := response.Config().Marshal().ToJson()
			w.Write([]byte(resp))
		}
	})

	http.HandleFunc("/monitor/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case http.MethodPost:
			body, _ := ioutil.ReadAll(r.Body)
			metricsReq := NewMetricsRequest()
			metricsReq.Unmarshal().FromJson(string(body))
			if metricsReq.Choice() == MetricsRequestChoice.FLOW {
				flow_responses := NewGetMetricsResponse().MetricsResponse()
				for _, flow_name := range metricsReq.Flow().FlowNames() {
					flow_rsp := flow_responses.FlowMetrics().Add()
					flow_rsp.SetName(flow_name)
					flow_rsp.SetBytesTx(1000)
					flow_rsp.SetBytesRx(1000)
				}
				response := NewGetMetricsResponse().MetricsResponse()
				flow_resp, _ := flow_responses.Marshal().ToJson()
				response.Unmarshal().FromJson(flow_resp)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(flow_resp))
			} else if metricsReq.Choice() == MetricsRequestChoice.PORT {
				port_response := NewGetMetricsResponse().MetricsResponse()
				for _, port_name := range metricsReq.Port().PortNames() {
					port_rsp := port_response.PortMetrics().Add()
					port_rsp.SetName(port_name)
					port_rsp.SetBytesTx(2000)
					port_rsp.SetBytesRx(2000)
				}
				response := NewGetMetricsResponse().MetricsResponse()
				port_resp, _ := port_response.Marshal().ToJson()
				response.Unmarshal().FromJson(port_resp)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(port_resp))
			} else if metricsReq.Choice() == MetricsRequestChoice.BGPV4 {
				bgpv4_response := NewGetMetricsResponse().MetricsResponse()
				for _, bgpv4_name := range metricsReq.Bgpv4().PeerNames() {
					bgpv4_rsp := bgpv4_response.Bgpv4Metrics().Add()
					bgpv4_rsp.SetName(bgpv4_name)
					bgpv4_rsp.SetRoutesAdvertised(80)
				}
				response := NewGetMetricsResponse().MetricsResponse()
				bgpv4_resp, _ := bgpv4_response.Marshal().ToJson()
				response.Unmarshal().FromJson(bgpv4_resp)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(bgpv4_resp))
			}
		}
	})

	go func() {
		if err := http.ListenAndServe(httpServer.serverLocation, nil); err != nil {
			log.Fatal("Server failed to serve incoming HTTP request.")
		}
	}()
}
