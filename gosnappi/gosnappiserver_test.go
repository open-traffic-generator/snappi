package gosnappi_test

import (
	"fmt"

	"net/http"
	"testing"

	gosnappi "github.com/open-traffic-generator/snappi/gosnappi"
	"github.com/open-traffic-generator/snappi/gosnappi/httpapi"

	"github.com/open-traffic-generator/snappi/gosnappi/httpapi/controllers"
	"github.com/open-traffic-generator/snappi/gosnappi/httpapi/interfaces"
)

//////////////////////////////////////////////////////////////////
// 			Strating gosnappi server with specific roure
//////////////////////////////////////////////////////////////////

func StartHttpServer() {
	fmt.Println("!!!!!!!!	Try to start gosnappi HTTP server	!!!!!!!!!!!!!")
	configHandler := NewConfigurationHandler()
	// Initiate other handler and append it
	controllers := []httpapi.HttpController{
		configHandler.GetController(),
	}
	router := httpapi.AppendRoutes(nil, controllers...)

	go func() {
		fmt.Println("Generated Http Server serving incoming HTTP requests on 127.0.0.1:50071.")
		if err := http.ListenAndServe("127.0.0.1:50071", router); err != nil {
			fmt.Println("Generated Http Server failed to serve incoming HTTP request.")
		}
	}()
}

//////////////////////////////////////////////////////////////////
// 		Initiating gosnappi handler and defined demo business logic
//////////////////////////////////////////////////////////////////

type configurationHandler struct {
	controller interfaces.ConfigurationController
}

func NewConfigurationHandler() interfaces.ConfigurationHandler {
	handler := new(configurationHandler)
	handler.controller = controllers.NewHttpConfigurationController(handler)
	return handler
}

func (h *configurationHandler) GetController() interfaces.ConfigurationController {
	return h.controller
}

func (h *configurationHandler) SetConfig(rbody gosnappi.Config, r *http.Request) (gosnappi.SetConfigResponse, error) {
	fmt.Println("SetConfig: We need to put our business logic ...")
	result := gosnappi.NewSetConfigResponse()
	result.Warning()
	return result, nil
}

func (h *configurationHandler) GetConfig(r *http.Request) (gosnappi.GetConfigResponse, error) {
	fmt.Println("GetConfig: We need to put our business logic ...")
	result := gosnappi.NewGetConfigResponse()
	result.Config()
	return result, nil

}

func (h *configurationHandler) UpdateConfig(rbody gosnappi.ConfigUpdate, r *http.Request) (gosnappi.UpdateConfigResponse, error) {
	fmt.Println("UpdateConfig: We need to put our business logic ...")
	result := gosnappi.NewUpdateConfigResponse()
	result.Warning()
	return result, nil
}

func (h *configurationHandler) AppendConfig(rbody gosnappi.ConfigAppend, r *http.Request) (gosnappi.AppendConfigResponse, error) {
	fmt.Println("AppendConfig: We need to put our business logic ...")
	result := gosnappi.NewAppendConfigResponse()
	result.Warning()
	return result, nil
}

func (h *configurationHandler) DeleteConfig(rbody gosnappi.ConfigDelete, r *http.Request) (gosnappi.DeleteConfigResponse, error) {
	fmt.Println("DeleteConfig: We need to put our business logic ...")
	result := gosnappi.NewDeleteConfigResponse()
	result.Warning()
	return result, nil
}

//////////////////////////////////////////////////////////////////
// 		Test Code where init gosnappi server and access it through gosnappi package
//////////////////////////////////////////////////////////////////

func init() {
	StartHttpServer()
}

func TestSnappiServer(t *testing.T) {
	api := gosnappi.NewApi()
	api.NewHttpTransport().SetLocation("http://127.0.0.1:50071")
	config := config1(api)
	state, err := api.SetConfig(config)
	data, _ := state.Marshal().ToJson()
	fmt.Println("state : ", data)
	fmt.Println("err: ", err)
}
