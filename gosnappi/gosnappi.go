/* Open Traffic Generator API 1.28.0
 * Open Traffic Generator API defines a model-driven, vendor-neutral and standard
 * interface for emulating layer 2-7 network devices and generating test traffic.
 *
 * Contributions can be made in the following ways:
 * - [open an issue](https://github.com/open-traffic-generator/models/issues) in the
 * models repository
 * - [fork the models repository](https://github.com/open-traffic-generator/models)
 * and submit a PR
 * License: MIT */

package gosnappi

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"

	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// function related to error handling
func FromError(err error) (Error, bool) {
	if rErr, ok := err.(Error); ok {
		return rErr, true
	}

	rErr := NewError()
	if err := rErr.Unmarshal().FromJson(err.Error()); err == nil {
		return rErr, true
	}

	return fromGrpcError(err)
}

func setResponseErr(obj Error, code int32, message string) {
	errors := []string{}
	errors = append(errors, message)
	obj.msg().Code = &code
	obj.msg().Errors = errors
}

// parses and return errors for grpc response
func fromGrpcError(err error) (Error, bool) {
	st, ok := status.FromError(err)
	if ok {
		rErr := NewError()
		if err := rErr.Unmarshal().FromJson(st.Message()); err == nil {
			var code = int32(st.Code())
			rErr.msg().Code = &code
			return rErr, true
		}

		setResponseErr(rErr, int32(st.Code()), st.Message())
		return rErr, true
	}

	return nil, false
}

// parses and return errors for http responses
func fromHttpError(statusCode int, body []byte) Error {
	rErr := NewError()
	bStr := string(body)
	if err := rErr.Unmarshal().FromJson(bStr); err == nil {
		return rErr
	}

	setResponseErr(rErr, int32(statusCode), bStr)

	return rErr
}

type versionMeta struct {
	checkVersion  bool
	localVersion  Version
	remoteVersion Version
	checkError    error
	clientName    string
	clientAppVer  string
	serverName    string
}
type gosnappiApi struct {
	apiSt
	grpcClient  otg.OpenapiClient
	httpClient  httpClient
	versionMeta *versionMeta
}

// grpcConnect builds up a grpc connection
func (api *gosnappiApi) grpcConnect() error {
	if api.grpcClient == nil {
		if api.grpc.clientConnection == nil {
			ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.dialTimeout)
			defer cancelFunc()
			conn, err := grpc.DialContext(ctx, api.grpc.location, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				return err
			}
			api.grpcClient = otg.NewOpenapiClient(conn)
			api.grpc.clientConnection = conn
		} else {
			api.grpcClient = otg.NewOpenapiClient(api.grpc.clientConnection)
		}
	}
	return nil
}

func (api *gosnappiApi) grpcClose() error {
	if api.grpc != nil {
		if api.grpc.clientConnection != nil {
			err := api.grpc.clientConnection.Close()
			if err != nil {
				return err
			}
		}
	}
	api.grpcClient = nil
	api.grpc = nil
	return nil
}

func (api *gosnappiApi) Close() error {
	if api.hasGrpcTransport() {
		err := api.grpcClose()
		return err
	}
	if api.hasHttpTransport() {
		err := api.http.conn.(*net.TCPConn).SetLinger(0)
		api.http.conn.Close()
		api.http.conn = nil
		api.http = nil
		api.httpClient.client = nil
		return err
	}
	return nil
}

// NewApi returns a new instance of the top level interface hierarchy
func NewApi() Api {
	api := gosnappiApi{}
	api.versionMeta = &versionMeta{checkVersion: false}
	return &api
}

// httpConnect builds up a http connection
func (api *gosnappiApi) httpConnect() error {
	if api.httpClient.client == nil {
		tr := http.Transport{
			DialTLSContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				tcpConn, err := (&net.Dialer{}).DialContext(ctx, network, addr)
				if err != nil {
					return nil, err
				}
				tlsConn := tls.Client(tcpConn, &tls.Config{InsecureSkipVerify: !api.http.verify})
				err = tlsConn.Handshake()
				if err != nil {
					return nil, err
				}
				api.http.conn = tcpConn
				return tlsConn, nil
			},
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				tcpConn, err := (&net.Dialer{}).DialContext(ctx, network, addr)
				if err != nil {
					return nil, err
				}
				api.http.conn = tcpConn
				return tcpConn, nil
			},
		}
		client := httpClient{
			client: &http.Client{
				Transport: &tr,
			},
			ctx: context.Background(),
		}
		api.httpClient = client
	}
	return nil
}

func (api *gosnappiApi) httpSendRecv(urlPath string, jsonBody string, method string) (*http.Response, error) {
	err := api.httpConnect()
	if err != nil {
		return nil, err
	}
	httpClient := api.httpClient
	var bodyReader = bytes.NewReader([]byte(jsonBody))
	queryUrl, err := url.Parse(api.http.location)
	if err != nil {
		return nil, err
	}
	queryUrl, _ = queryUrl.Parse(urlPath)
	req, _ := http.NewRequest(method, queryUrl.String(), bodyReader)
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(httpClient.ctx)
	response, err := httpClient.client.Do(req)
	return response, err
}

// GosnappiApi open Traffic Generator API defines a model-driven, vendor-neutral and standard
// interface for emulating layer 2-7 network devices and generating test traffic.
//
// Contributions can be made in the following ways:
// - [open an issue](https://github.com/open-traffic-generator/models/issues) in the models repository
// - [fork the models repository](https://github.com/open-traffic-generator/models) and submit a PR
type Api interface {
	api
	// SetConfig sets configuration resources on the traffic generator.
	SetConfig(config Config) (Warning, error)
	// GetConfig description is TBD
	GetConfig() (Config, error)
	// UpdateConfig updates specific attributes of resources configured on the traffic generator. The fetched configuration shall reflect the updates applied successfully.
	// The Response.Warnings in the Success response is available for implementers to disclose additional information about a state change including any implicit changes that are outside the scope of the state change.
	UpdateConfig(configUpdate ConfigUpdate) (Warning, error)
	// AppendConfig append new attributes of resources to existing configuration on the traffic generator. Resource names should not be part of existing configuration of that resource type; it should be unique for the operation to succeed. The fetched configuration shall also reflect the new configuration applied successfully.
	AppendConfig(configAppend ConfigAppend) (Warning, error)
	// DeleteConfig delete attributes of resources from existing configuration on the traffic generator. Resource names should already be part of existing configuration of that resource type; for the operation to succeed. The fetched configuration shall not reflect the removed configuration deleted successfully.
	DeleteConfig(configDelete ConfigDelete) (Warning, error)
	// SetControlState sets the operational state of configured resources.
	SetControlState(controlState ControlState) (Warning, error)
	// SetControlAction triggers actions against configured resources.
	SetControlAction(controlAction ControlAction) (ControlActionResponse, error)
	// GetMetrics description is TBD
	GetMetrics(metricsRequest MetricsRequest) (MetricsResponse, error)
	// GetStates description is TBD
	GetStates(statesRequest StatesRequest) (StatesResponse, error)
	// GetCapture description is TBD
	GetCapture(captureRequest CaptureRequest) ([]byte, error)
	// GetVersion description is TBD
	GetVersion() (Version, error)
	// GetLocalVersion provides version details of local client
	GetLocalVersion() Version
	// GetRemoteVersion provides version details received from remote server
	GetRemoteVersion() (Version, error)
	// SetVersionCompatibilityCheck allows enabling or disabling automatic version
	// compatibility check between client and server API spec version upon API call
	SetVersionCompatibilityCheck(bool)
	// ability to set component names and specific version for version check error msgs
	SetComponentInformation(string, string, string)
	// CheckVersionCompatibility compares API spec version for local client and remote server,
	// and returns an error if they are not compatible according to Semantic Versioning 2.0.0
	CheckVersionCompatibility() error
}

func (api *gosnappiApi) GetLocalVersion() Version {
	if api.versionMeta.localVersion == nil {
		api.versionMeta.localVersion = NewVersion().SetApiSpecVersion("1.28.0").SetSdkVersion("1.28.2")
	}

	return api.versionMeta.localVersion
}

func (api *gosnappiApi) GetRemoteVersion() (Version, error) {
	if api.versionMeta.remoteVersion == nil {
		v, err := api.GetVersion()
		if err != nil {
			return nil, fmt.Errorf("could not fetch remote version: %v", err)
		}

		api.versionMeta.remoteVersion = v
	}

	return api.versionMeta.remoteVersion, nil
}

func (api *gosnappiApi) SetVersionCompatibilityCheck(v bool) {
	api.versionMeta.checkVersion = v
}

func (api *gosnappiApi) SetComponentInformation(clientName string, clientVer string, serverName string) {
	api.versionMeta.clientName = clientName
	api.versionMeta.clientAppVer = clientVer
	api.versionMeta.serverName = serverName
}

func (api *gosnappiApi) checkLocalRemoteVersionCompatibility() (error, error) {
	localVer := api.GetLocalVersion()
	remoteVer, err := api.GetRemoteVersion()
	if err != nil {
		return nil, err
	}
	err = checkClientServerVersionCompatibility(localVer.ApiSpecVersion(), remoteVer.ApiSpecVersion(), "API spec")
	if err != nil {
		if api.versionMeta.clientName != "" {
			return fmt.Errorf(
				"%s %s is not compatible with %s %s", api.versionMeta.clientName, api.versionMeta.clientAppVer,
				api.versionMeta.serverName, remoteVer.AppVersion()), nil
		} else {
			return fmt.Errorf(
				"client SDK version '%s' is not compatible with server SDK version '%s': %v",
				localVer.SdkVersion(), remoteVer.SdkVersion(), err,
			), nil
		}
	}

	return nil, nil
}

func (api *gosnappiApi) checkLocalRemoteVersionCompatibilityOnce() error {
	if !api.versionMeta.checkVersion {
		return nil
	}

	if api.versionMeta.checkError != nil {
		return api.versionMeta.checkError
	}

	compatErr, apiErr := api.checkLocalRemoteVersionCompatibility()
	if compatErr != nil {
		api.versionMeta.checkError = compatErr
		return compatErr
	}
	if apiErr != nil {
		api.versionMeta.checkError = nil
		return apiErr
	}

	api.versionMeta.checkVersion = false
	api.versionMeta.checkError = nil
	return nil
}

func (api *gosnappiApi) CheckVersionCompatibility() error {
	compatErr, apiErr := api.checkLocalRemoteVersionCompatibility()
	if compatErr != nil {
		return fmt.Errorf("version error: %v", compatErr)
	}
	if apiErr != nil {
		return apiErr
	}

	return nil
}

func (api *gosnappiApi) SetConfig(config Config) (Warning, error) {

	if err := config.validate(); err != nil {
		return nil, err
	}

	if err := api.checkLocalRemoteVersionCompatibilityOnce(); err != nil {
		return nil, err
	}
	if api.hasHttpTransport() {
		return api.httpSetConfig(config)
	}
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := otg.SetConfigRequest{Config: config.msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.SetConfig(ctx, &request)
	if err != nil {
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewWarning()
	if resp.GetWarning() != nil {
		return ret.setMsg(resp.GetWarning()), nil
	}

	return ret, nil
}

func (api *gosnappiApi) GetConfig() (Config, error) {

	if err := api.checkLocalRemoteVersionCompatibilityOnce(); err != nil {
		return nil, err
	}
	if api.hasHttpTransport() {
		return api.httpGetConfig()
	}
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := emptypb.Empty{}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.GetConfig(ctx, &request)
	if err != nil {
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewConfig()
	if resp.GetConfig() != nil {
		return ret.setMsg(resp.GetConfig()), nil
	}

	return ret, nil
}

func (api *gosnappiApi) UpdateConfig(configUpdate ConfigUpdate) (Warning, error) {

	if err := configUpdate.validate(); err != nil {
		return nil, err
	}

	if err := api.checkLocalRemoteVersionCompatibilityOnce(); err != nil {
		return nil, err
	}
	if api.hasHttpTransport() {
		return api.httpUpdateConfig(configUpdate)
	}
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := otg.UpdateConfigRequest{ConfigUpdate: configUpdate.msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.UpdateConfig(ctx, &request)
	if err != nil {
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewWarning()
	if resp.GetWarning() != nil {
		return ret.setMsg(resp.GetWarning()), nil
	}

	return ret, nil
}

func (api *gosnappiApi) AppendConfig(configAppend ConfigAppend) (Warning, error) {

	if err := configAppend.validate(); err != nil {
		return nil, err
	}

	if err := api.checkLocalRemoteVersionCompatibilityOnce(); err != nil {
		return nil, err
	}
	if api.hasHttpTransport() {
		return api.httpAppendConfig(configAppend)
	}
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := otg.AppendConfigRequest{ConfigAppend: configAppend.msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.AppendConfig(ctx, &request)
	if err != nil {
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewWarning()
	if resp.GetWarning() != nil {
		return ret.setMsg(resp.GetWarning()), nil
	}

	return ret, nil
}

func (api *gosnappiApi) DeleteConfig(configDelete ConfigDelete) (Warning, error) {

	if err := configDelete.validate(); err != nil {
		return nil, err
	}

	if err := api.checkLocalRemoteVersionCompatibilityOnce(); err != nil {
		return nil, err
	}
	if api.hasHttpTransport() {
		return api.httpDeleteConfig(configDelete)
	}
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := otg.DeleteConfigRequest{ConfigDelete: configDelete.msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.DeleteConfig(ctx, &request)
	if err != nil {
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewWarning()
	if resp.GetWarning() != nil {
		return ret.setMsg(resp.GetWarning()), nil
	}

	return ret, nil
}

func (api *gosnappiApi) SetControlState(controlState ControlState) (Warning, error) {

	if err := controlState.validate(); err != nil {
		return nil, err
	}

	if err := api.checkLocalRemoteVersionCompatibilityOnce(); err != nil {
		return nil, err
	}
	if api.hasHttpTransport() {
		return api.httpSetControlState(controlState)
	}
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := otg.SetControlStateRequest{ControlState: controlState.msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.SetControlState(ctx, &request)
	if err != nil {
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewWarning()
	if resp.GetWarning() != nil {
		return ret.setMsg(resp.GetWarning()), nil
	}

	return ret, nil
}

func (api *gosnappiApi) SetControlAction(controlAction ControlAction) (ControlActionResponse, error) {

	if err := controlAction.validate(); err != nil {
		return nil, err
	}

	if err := api.checkLocalRemoteVersionCompatibilityOnce(); err != nil {
		return nil, err
	}
	if api.hasHttpTransport() {
		return api.httpSetControlAction(controlAction)
	}
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := otg.SetControlActionRequest{ControlAction: controlAction.msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.SetControlAction(ctx, &request)
	if err != nil {
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewControlActionResponse()
	if resp.GetControlActionResponse() != nil {
		return ret.setMsg(resp.GetControlActionResponse()), nil
	}

	return ret, nil
}

func (api *gosnappiApi) GetMetrics(metricsRequest MetricsRequest) (MetricsResponse, error) {

	if err := metricsRequest.validate(); err != nil {
		return nil, err
	}

	if err := api.checkLocalRemoteVersionCompatibilityOnce(); err != nil {
		return nil, err
	}
	if api.hasHttpTransport() {
		return api.httpGetMetrics(metricsRequest)
	}
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := otg.GetMetricsRequest{MetricsRequest: metricsRequest.msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.GetMetrics(ctx, &request)
	if err != nil {
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewMetricsResponse()
	if resp.GetMetricsResponse() != nil {
		return ret.setMsg(resp.GetMetricsResponse()), nil
	}

	return ret, nil
}

func (api *gosnappiApi) GetStates(statesRequest StatesRequest) (StatesResponse, error) {

	if err := statesRequest.validate(); err != nil {
		return nil, err
	}

	if err := api.checkLocalRemoteVersionCompatibilityOnce(); err != nil {
		return nil, err
	}
	if api.hasHttpTransport() {
		return api.httpGetStates(statesRequest)
	}
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := otg.GetStatesRequest{StatesRequest: statesRequest.msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.GetStates(ctx, &request)
	if err != nil {
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewStatesResponse()
	if resp.GetStatesResponse() != nil {
		return ret.setMsg(resp.GetStatesResponse()), nil
	}

	return ret, nil
}

func (api *gosnappiApi) GetCapture(captureRequest CaptureRequest) ([]byte, error) {

	if err := captureRequest.validate(); err != nil {
		return nil, err
	}

	if err := api.checkLocalRemoteVersionCompatibilityOnce(); err != nil {
		return nil, err
	}
	if api.hasHttpTransport() {
		return api.httpGetCapture(captureRequest)
	}
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := otg.GetCaptureRequest{CaptureRequest: captureRequest.msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.GetCapture(ctx, &request)
	if err != nil {
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	if resp.ResponseBytes != nil {
		return resp.ResponseBytes, nil
	}
	return nil, nil
}

func (api *gosnappiApi) GetVersion() (Version, error) {

	if api.hasHttpTransport() {
		return api.httpGetVersion()
	}
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := emptypb.Empty{}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.GetVersion(ctx, &request)
	if err != nil {
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewVersion()
	if resp.GetVersion() != nil {
		return ret.setMsg(resp.GetVersion()), nil
	}

	return ret, nil
}

func (api *gosnappiApi) httpSetConfig(config Config) (Warning, error) {
	configJson, err := config.Marshal().ToJson()
	if err != nil {
		return nil, err
	}
	resp, err := api.httpSendRecv("config", configJson, "POST")

	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		obj := NewSetConfigResponse().Warning()
		if err := obj.Unmarshal().FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, fromHttpError(resp.StatusCode, bodyBytes)
	}
}

func (api *gosnappiApi) httpGetConfig() (Config, error) {
	resp, err := api.httpSendRecv("config", "", "GET")
	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		obj := NewGetConfigResponse().Config()
		if err := obj.Unmarshal().FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, fromHttpError(resp.StatusCode, bodyBytes)
	}
}

func (api *gosnappiApi) httpUpdateConfig(configUpdate ConfigUpdate) (Warning, error) {
	configUpdateJson, err := configUpdate.Marshal().ToJson()
	if err != nil {
		return nil, err
	}
	resp, err := api.httpSendRecv("config", configUpdateJson, "PATCH")

	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		obj := NewUpdateConfigResponse().Warning()
		if err := obj.Unmarshal().FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, fromHttpError(resp.StatusCode, bodyBytes)
	}
}

func (api *gosnappiApi) httpAppendConfig(configAppend ConfigAppend) (Warning, error) {
	configAppendJson, err := configAppend.Marshal().ToJson()
	if err != nil {
		return nil, err
	}
	resp, err := api.httpSendRecv("config/append", configAppendJson, "PATCH")

	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		obj := NewAppendConfigResponse().Warning()
		if err := obj.Unmarshal().FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, fromHttpError(resp.StatusCode, bodyBytes)
	}
}

func (api *gosnappiApi) httpDeleteConfig(configDelete ConfigDelete) (Warning, error) {
	configDeleteJson, err := configDelete.Marshal().ToJson()
	if err != nil {
		return nil, err
	}
	resp, err := api.httpSendRecv("config/delete", configDeleteJson, "PATCH")

	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		obj := NewDeleteConfigResponse().Warning()
		if err := obj.Unmarshal().FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, fromHttpError(resp.StatusCode, bodyBytes)
	}
}

func (api *gosnappiApi) httpSetControlState(controlState ControlState) (Warning, error) {
	controlStateJson, err := controlState.Marshal().ToJson()
	if err != nil {
		return nil, err
	}
	resp, err := api.httpSendRecv("control/state", controlStateJson, "POST")

	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		obj := NewSetControlStateResponse().Warning()
		if err := obj.Unmarshal().FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, fromHttpError(resp.StatusCode, bodyBytes)
	}
}

func (api *gosnappiApi) httpSetControlAction(controlAction ControlAction) (ControlActionResponse, error) {
	controlActionJson, err := controlAction.Marshal().ToJson()
	if err != nil {
		return nil, err
	}
	resp, err := api.httpSendRecv("control/action", controlActionJson, "POST")

	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		obj := NewSetControlActionResponse().ControlActionResponse()
		if err := obj.Unmarshal().FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, fromHttpError(resp.StatusCode, bodyBytes)
	}
}

func (api *gosnappiApi) httpGetMetrics(metricsRequest MetricsRequest) (MetricsResponse, error) {
	metricsRequestJson, err := metricsRequest.Marshal().ToJson()
	if err != nil {
		return nil, err
	}
	resp, err := api.httpSendRecv("monitor/metrics", metricsRequestJson, "POST")

	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		obj := NewGetMetricsResponse().MetricsResponse()
		if err := obj.Unmarshal().FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, fromHttpError(resp.StatusCode, bodyBytes)
	}
}

func (api *gosnappiApi) httpGetStates(statesRequest StatesRequest) (StatesResponse, error) {
	statesRequestJson, err := statesRequest.Marshal().ToJson()
	if err != nil {
		return nil, err
	}
	resp, err := api.httpSendRecv("monitor/states", statesRequestJson, "POST")

	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		obj := NewGetStatesResponse().StatesResponse()
		if err := obj.Unmarshal().FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, fromHttpError(resp.StatusCode, bodyBytes)
	}
}

func (api *gosnappiApi) httpGetCapture(captureRequest CaptureRequest) ([]byte, error) {
	captureRequestJson, err := captureRequest.Marshal().ToJson()
	if err != nil {
		return nil, err
	}
	resp, err := api.httpSendRecv("monitor/capture", captureRequestJson, "POST")

	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		return bodyBytes, nil
	} else {
		return nil, fromHttpError(resp.StatusCode, bodyBytes)
	}
}

func (api *gosnappiApi) httpGetVersion() (Version, error) {
	resp, err := api.httpSendRecv("capabilities/version", "", "GET")
	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		obj := NewGetVersionResponse().Version()
		if err := obj.Unmarshal().FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, fromHttpError(resp.StatusCode, bodyBytes)
	}
}
