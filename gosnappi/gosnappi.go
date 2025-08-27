/* Open Traffic Generator API 1.34.0
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
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

var loggerSt = &logger{}
var logs = loggerSt.getLogger("otg")

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
			var opts []grpc.DialOption
			opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if api.Telemetry().isOTLPEnabled() {
				opts = append(opts, grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
			}
			conn, err := grpc.DialContext(ctx, api.grpc.location, opts...)
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
	api.tracer = &telemetry{transport: "HTTP", serviceName: "go-snappi"}
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

		var client httpClient
		if api.Telemetry().isOTLPEnabled() {
			client = httpClient{
				client: &http.Client{
					Transport: otelhttp.NewTransport(&tr),
				},
				ctx: api.Telemetry().getRootContext(),
			}
		} else {
			client = httpClient{
				client: &http.Client{
					Transport: &tr,
				},
				ctx: context.Background(),
			}
		}

		api.httpClient = client
	}
	return nil
}

func (api *gosnappiApi) httpSendRecv(ctx context.Context, urlPath string, jsonBody string, method string) (*http.Response, error) {
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
	if ctx != nil {
		req = req.WithContext(ctx)
	} else {
		req = req.WithContext(httpClient.ctx)
	}
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
	// streaming method for SetConfig
	streamSetConfig(context.Context, []byte) (*otg.SetConfigResponse, error)
	// GetConfig description is TBD
	GetConfig() (Config, error)
	// streaming method for GetConfig
	streamGetConfig(context.Context, *emptypb.Empty) (Config, error)
	// UpdateConfig updates specific attributes of resources configured on the traffic generator. The fetched configuration shall reflect the updates applied successfully.
	// The Response.Warnings in the Success response is available for implementers to disclose additional information about a state change including any implicit changes that are outside the scope of the state change.
	UpdateConfig(configUpdate ConfigUpdate) (Warning, error)
	// AppendConfig append new attributes of resources to existing configuration on the traffic generator. Resource names should not be part of existing configuration of that resource type; it should be unique for the operation to succeed. A failed append might leave the configuration in an undefined state and if the error is due to some invalid or unsupported configuration in the appended resources, it is expected that the user fix the error and  restart from SetConfig operation. The fetched configuration shall also reflect the new configuration applied successfully.
	AppendConfig(configAppend ConfigAppend) (Warning, error)
	// DeleteConfig delete attributes of resources from existing configuration on the traffic generator. Resource names should already be part of existing configuration of that resource type; for the operation to succeed. A failed delete will leave the configuration in an undefined state and if the error is due to some invalid or unsupported configuration in the deleted  resources, it is expected that the user fix the error and restart from SetConfig operation. On successful deletion the fetched configuration shall not reflect the removed configuration.
	DeleteConfig(configDelete ConfigDelete) (Warning, error)
	// SetControlState sets the operational state of configured resources.
	SetControlState(controlState ControlState) (Warning, error)
	// streaming method for SetControlState
	streamSetControlState(context.Context, []byte) (*otg.SetControlStateResponse, error)
	// SetControlAction triggers actions against configured resources.
	SetControlAction(controlAction ControlAction) (ControlActionResponse, error)
	// streaming method for SetControlAction
	streamSetControlAction(context.Context, []byte) (*otg.SetControlActionResponse, error)
	// GetMetrics description is TBD
	GetMetrics(metricsRequest MetricsRequest) (MetricsResponse, error)
	// streaming method for GetMetrics
	streamGetMetrics(context.Context, *otg.GetMetricsRequest) (MetricsResponse, error)
	// GetStates description is TBD
	GetStates(statesRequest StatesRequest) (StatesResponse, error)
	// streaming method for GetStates
	streamGetStates(context.Context, *otg.GetStatesRequest) (StatesResponse, error)
	// GetCapture description is TBD
	GetCapture(captureRequest CaptureRequest) ([]byte, error)
	// streaming method for GetCapture
	streamGetCapture(context.Context, *otg.GetCaptureRequest) ([]byte, error)
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
		api.versionMeta.localVersion = NewVersion().SetApiSpecVersion("1.34.0").SetSdkVersion("1.34.1")
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

func (api *gosnappiApi) streamSetConfig(ctx context.Context, data []byte) (*otg.SetConfigResponse, error) {
	chunkSize := api.grpc.chunkSize
	streamClient, err := api.grpcClient.StreamSetConfig(ctx)
	if err != nil {
		return nil, err
	}
	bytes := []byte(data)
	var i uint64
	for i = 0; i < uint64(len(bytes)); i += chunkSize {
		data := &otg.Data{}
		data.ChunkSize = uint64(chunkSize)
		if i+chunkSize > uint64(len(bytes)) {
			data.Datum = bytes[i:]
		} else {
			data.Datum = bytes[i : i+chunkSize]
		}
		if err := streamClient.Send(data); err != nil {
			return nil, err
		}
	}
	res, err := streamClient.CloseAndRecv()
	if err != nil {
		return nil, err
	}
	return res, nil
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

	// adding spans grpc transport for OTLP instrumentation
	parentCtx := api.Telemetry().getRootContext()
	newCtx, span := api.Telemetry().NewSpan(parentCtx, "SetConfig", trace.WithSpanKind(trace.SpanKindClient))
	defer api.Telemetry().CloseSpan(span)

	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := otg.SetConfigRequest{Config: config.msg()}
	if newCtx == nil {
		newCtx = context.Background()
	}
	ctx, cancelFunc := context.WithTimeout(newCtx, api.grpc.requestTimeout)
	defer cancelFunc()
	var resp *otg.SetConfigResponse
	var err error
	if api.grpc.enableGrpcStreaming {
		str, er := proto.Marshal(config.msg())
		if er != nil {
			return nil, er
		}
		resp, err = api.streamSetConfig(ctx, str)
	} else {

		resp, err = api.grpcClient.SetConfig(ctx, &request)
	}
	if err != nil {
		api.Telemetry().SetSpanStatus(span, codes.Error, err.Error())
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewWarning()
	if resp.GetWarning() != nil {
		ret.setMsg(resp.GetWarning())
		return ret, nil
	}

	return ret, nil
}

func (api *gosnappiApi) streamGetConfig(ctx context.Context, req *emptypb.Empty) (Config, error) {
	streamClient, err := api.grpcClient.StreamGetConfig(ctx, req)
	if err != nil {
		return nil, err
	}
	var bytes []byte
	for {
		resp, err := streamClient.Recv()
		if err == io.EOF {
			res := NewConfig()
			m_err := proto.Unmarshal(bytes, res.msg())
			if m_err != nil {
				return nil, m_err
			} else {
				return res, nil
			}

		}
		if err != nil {
			return nil, err
		}
		bytes = append(bytes, resp.Datum...)
	}
}

func (api *gosnappiApi) GetConfig() (Config, error) {

	if err := api.checkLocalRemoteVersionCompatibilityOnce(); err != nil {
		return nil, err
	}

	if api.hasHttpTransport() {
		return api.httpGetConfig()
	}

	// adding spans grpc transport for OTLP instrumentation
	parentCtx := api.Telemetry().getRootContext()
	newCtx, span := api.Telemetry().NewSpan(parentCtx, "GetConfig", trace.WithSpanKind(trace.SpanKindClient))
	defer api.Telemetry().CloseSpan(span)

	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := emptypb.Empty{}
	if newCtx == nil {
		newCtx = context.Background()
	}
	ctx, cancelFunc := context.WithTimeout(newCtx, api.grpc.requestTimeout)
	defer cancelFunc()
	var resp *otg.GetConfigResponse
	var err error
	if api.grpc.enableGrpcStreaming {
		return api.streamGetConfig(ctx, &request)
	} else {

		resp, err = api.grpcClient.GetConfig(ctx, &request)
	}
	if err != nil {
		api.Telemetry().SetSpanStatus(span, codes.Error, err.Error())
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewConfig()
	if resp.GetConfig() != nil {
		ret.setMsg(resp.GetConfig())
		return ret, nil
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

	// adding spans grpc transport for OTLP instrumentation
	parentCtx := api.Telemetry().getRootContext()
	newCtx, span := api.Telemetry().NewSpan(parentCtx, "UpdateConfig", trace.WithSpanKind(trace.SpanKindClient))
	defer api.Telemetry().CloseSpan(span)

	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := otg.UpdateConfigRequest{ConfigUpdate: configUpdate.msg()}
	if newCtx == nil {
		newCtx = context.Background()
	}
	ctx, cancelFunc := context.WithTimeout(newCtx, api.grpc.requestTimeout)
	defer cancelFunc()

	resp, err := api.grpcClient.UpdateConfig(ctx, &request)

	if err != nil {
		api.Telemetry().SetSpanStatus(span, codes.Error, err.Error())
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewWarning()
	if resp.GetWarning() != nil {
		ret.setMsg(resp.GetWarning())
		return ret, nil
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

	// adding spans grpc transport for OTLP instrumentation
	parentCtx := api.Telemetry().getRootContext()
	newCtx, span := api.Telemetry().NewSpan(parentCtx, "AppendConfig", trace.WithSpanKind(trace.SpanKindClient))
	defer api.Telemetry().CloseSpan(span)

	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := otg.AppendConfigRequest{ConfigAppend: configAppend.msg()}
	if newCtx == nil {
		newCtx = context.Background()
	}
	ctx, cancelFunc := context.WithTimeout(newCtx, api.grpc.requestTimeout)
	defer cancelFunc()

	resp, err := api.grpcClient.AppendConfig(ctx, &request)

	if err != nil {
		api.Telemetry().SetSpanStatus(span, codes.Error, err.Error())
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewWarning()
	if resp.GetWarning() != nil {
		ret.setMsg(resp.GetWarning())
		return ret, nil
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

	// adding spans grpc transport for OTLP instrumentation
	parentCtx := api.Telemetry().getRootContext()
	newCtx, span := api.Telemetry().NewSpan(parentCtx, "DeleteConfig", trace.WithSpanKind(trace.SpanKindClient))
	defer api.Telemetry().CloseSpan(span)

	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := otg.DeleteConfigRequest{ConfigDelete: configDelete.msg()}
	if newCtx == nil {
		newCtx = context.Background()
	}
	ctx, cancelFunc := context.WithTimeout(newCtx, api.grpc.requestTimeout)
	defer cancelFunc()

	resp, err := api.grpcClient.DeleteConfig(ctx, &request)

	if err != nil {
		api.Telemetry().SetSpanStatus(span, codes.Error, err.Error())
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewWarning()
	if resp.GetWarning() != nil {
		ret.setMsg(resp.GetWarning())
		return ret, nil
	}

	return ret, nil
}

func (api *gosnappiApi) streamSetControlState(ctx context.Context, data []byte) (*otg.SetControlStateResponse, error) {
	chunkSize := api.grpc.chunkSize
	streamClient, err := api.grpcClient.StreamSetControlState(ctx)
	if err != nil {
		return nil, err
	}
	bytes := []byte(data)
	var i uint64
	for i = 0; i < uint64(len(bytes)); i += chunkSize {
		data := &otg.Data{}
		data.ChunkSize = uint64(chunkSize)
		if i+chunkSize > uint64(len(bytes)) {
			data.Datum = bytes[i:]
		} else {
			data.Datum = bytes[i : i+chunkSize]
		}
		if err := streamClient.Send(data); err != nil {
			return nil, err
		}
	}
	res, err := streamClient.CloseAndRecv()
	if err != nil {
		return nil, err
	}
	return res, nil
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

	// adding spans grpc transport for OTLP instrumentation
	parentCtx := api.Telemetry().getRootContext()
	newCtx, span := api.Telemetry().NewSpan(parentCtx, "SetControlState", trace.WithSpanKind(trace.SpanKindClient))
	defer api.Telemetry().CloseSpan(span)

	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := otg.SetControlStateRequest{ControlState: controlState.msg()}
	if newCtx == nil {
		newCtx = context.Background()
	}
	ctx, cancelFunc := context.WithTimeout(newCtx, api.grpc.requestTimeout)
	defer cancelFunc()
	var resp *otg.SetControlStateResponse
	var err error
	if api.grpc.enableGrpcStreaming {
		str, er := proto.Marshal(controlState.msg())
		if er != nil {
			return nil, er
		}
		resp, err = api.streamSetControlState(ctx, str)
	} else {

		resp, err = api.grpcClient.SetControlState(ctx, &request)
	}
	if err != nil {
		api.Telemetry().SetSpanStatus(span, codes.Error, err.Error())
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewWarning()
	if resp.GetWarning() != nil {
		ret.setMsg(resp.GetWarning())
		return ret, nil
	}

	return ret, nil
}

func (api *gosnappiApi) streamSetControlAction(ctx context.Context, data []byte) (*otg.SetControlActionResponse, error) {
	chunkSize := api.grpc.chunkSize
	streamClient, err := api.grpcClient.StreamSetControlAction(ctx)
	if err != nil {
		return nil, err
	}
	bytes := []byte(data)
	var i uint64
	for i = 0; i < uint64(len(bytes)); i += chunkSize {
		data := &otg.Data{}
		data.ChunkSize = uint64(chunkSize)
		if i+chunkSize > uint64(len(bytes)) {
			data.Datum = bytes[i:]
		} else {
			data.Datum = bytes[i : i+chunkSize]
		}
		if err := streamClient.Send(data); err != nil {
			return nil, err
		}
	}
	res, err := streamClient.CloseAndRecv()
	if err != nil {
		return nil, err
	}
	return res, nil
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

	// adding spans grpc transport for OTLP instrumentation
	parentCtx := api.Telemetry().getRootContext()
	newCtx, span := api.Telemetry().NewSpan(parentCtx, "SetControlAction", trace.WithSpanKind(trace.SpanKindClient))
	defer api.Telemetry().CloseSpan(span)

	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := otg.SetControlActionRequest{ControlAction: controlAction.msg()}
	if newCtx == nil {
		newCtx = context.Background()
	}
	ctx, cancelFunc := context.WithTimeout(newCtx, api.grpc.requestTimeout)
	defer cancelFunc()
	var resp *otg.SetControlActionResponse
	var err error
	if api.grpc.enableGrpcStreaming {
		str, er := proto.Marshal(controlAction.msg())
		if er != nil {
			return nil, er
		}
		resp, err = api.streamSetControlAction(ctx, str)
	} else {

		resp, err = api.grpcClient.SetControlAction(ctx, &request)
	}
	if err != nil {
		api.Telemetry().SetSpanStatus(span, codes.Error, err.Error())
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewControlActionResponse()
	if resp.GetControlActionResponse() != nil {
		ret.setMsg(resp.GetControlActionResponse())
		return ret, nil
	}

	return ret, nil
}

func (api *gosnappiApi) streamGetMetrics(ctx context.Context, req *otg.GetMetricsRequest) (MetricsResponse, error) {
	streamClient, err := api.grpcClient.StreamGetMetrics(ctx, req)
	if err != nil {
		return nil, err
	}
	var bytes []byte
	for {
		resp, err := streamClient.Recv()
		if err == io.EOF {
			res := NewMetricsResponse()
			m_err := proto.Unmarshal(bytes, res.msg())
			if m_err != nil {
				return nil, m_err
			} else {
				return res, nil
			}

		}
		if err != nil {
			return nil, err
		}
		bytes = append(bytes, resp.Datum...)
	}
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

	// adding spans grpc transport for OTLP instrumentation
	parentCtx := api.Telemetry().getRootContext()
	newCtx, span := api.Telemetry().NewSpan(parentCtx, "GetMetrics", trace.WithSpanKind(trace.SpanKindClient))
	defer api.Telemetry().CloseSpan(span)

	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := otg.GetMetricsRequest{MetricsRequest: metricsRequest.msg()}
	if newCtx == nil {
		newCtx = context.Background()
	}
	ctx, cancelFunc := context.WithTimeout(newCtx, api.grpc.requestTimeout)
	defer cancelFunc()
	var resp *otg.GetMetricsResponse
	var err error
	if api.grpc.enableGrpcStreaming {
		return api.streamGetMetrics(ctx, &request)
	} else {

		resp, err = api.grpcClient.GetMetrics(ctx, &request)
	}
	if err != nil {
		api.Telemetry().SetSpanStatus(span, codes.Error, err.Error())
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewMetricsResponse()
	if resp.GetMetricsResponse() != nil {
		ret.setMsg(resp.GetMetricsResponse())
		return ret, nil
	}

	return ret, nil
}

func (api *gosnappiApi) streamGetStates(ctx context.Context, req *otg.GetStatesRequest) (StatesResponse, error) {
	streamClient, err := api.grpcClient.StreamGetStates(ctx, req)
	if err != nil {
		return nil, err
	}
	var bytes []byte
	for {
		resp, err := streamClient.Recv()
		if err == io.EOF {
			res := NewStatesResponse()
			m_err := proto.Unmarshal(bytes, res.msg())
			if m_err != nil {
				return nil, m_err
			} else {
				return res, nil
			}

		}
		if err != nil {
			return nil, err
		}
		bytes = append(bytes, resp.Datum...)
	}
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

	// adding spans grpc transport for OTLP instrumentation
	parentCtx := api.Telemetry().getRootContext()
	newCtx, span := api.Telemetry().NewSpan(parentCtx, "GetStates", trace.WithSpanKind(trace.SpanKindClient))
	defer api.Telemetry().CloseSpan(span)

	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := otg.GetStatesRequest{StatesRequest: statesRequest.msg()}
	if newCtx == nil {
		newCtx = context.Background()
	}
	ctx, cancelFunc := context.WithTimeout(newCtx, api.grpc.requestTimeout)
	defer cancelFunc()
	var resp *otg.GetStatesResponse
	var err error
	if api.grpc.enableGrpcStreaming {
		return api.streamGetStates(ctx, &request)
	} else {

		resp, err = api.grpcClient.GetStates(ctx, &request)
	}
	if err != nil {
		api.Telemetry().SetSpanStatus(span, codes.Error, err.Error())
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewStatesResponse()
	if resp.GetStatesResponse() != nil {
		ret.setMsg(resp.GetStatesResponse())
		return ret, nil
	}

	return ret, nil
}

func (api *gosnappiApi) streamGetCapture(ctx context.Context, req *otg.GetCaptureRequest) ([]byte, error) {
	streamClient, err := api.grpcClient.StreamGetCapture(ctx, req)
	if err != nil {
		return nil, err
	}
	var bytes []byte
	for {
		resp, err := streamClient.Recv()
		if err == io.EOF {
			return bytes, nil
		}
		if err != nil {
			return nil, err
		}
		bytes = append(bytes, resp.Datum...)
	}
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

	// adding spans grpc transport for OTLP instrumentation
	parentCtx := api.Telemetry().getRootContext()
	newCtx, span := api.Telemetry().NewSpan(parentCtx, "GetCapture", trace.WithSpanKind(trace.SpanKindClient))
	defer api.Telemetry().CloseSpan(span)

	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := otg.GetCaptureRequest{CaptureRequest: captureRequest.msg()}
	if newCtx == nil {
		newCtx = context.Background()
	}
	ctx, cancelFunc := context.WithTimeout(newCtx, api.grpc.requestTimeout)
	defer cancelFunc()
	var resp *otg.GetCaptureResponse
	var err error
	if api.grpc.enableGrpcStreaming {
		return api.streamGetCapture(ctx, &request)
	} else {

		resp, err = api.grpcClient.GetCapture(ctx, &request)
	}
	if err != nil {
		api.Telemetry().SetSpanStatus(span, codes.Error, err.Error())
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

	// adding spans grpc transport for OTLP instrumentation
	parentCtx := api.Telemetry().getRootContext()
	newCtx, span := api.Telemetry().NewSpan(parentCtx, "GetVersion", trace.WithSpanKind(trace.SpanKindClient))
	defer api.Telemetry().CloseSpan(span)

	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := emptypb.Empty{}
	if newCtx == nil {
		newCtx = context.Background()
	}
	ctx, cancelFunc := context.WithTimeout(newCtx, api.grpc.requestTimeout)
	defer cancelFunc()

	resp, err := api.grpcClient.GetVersion(ctx, &request)

	if err != nil {
		api.Telemetry().SetSpanStatus(span, codes.Error, err.Error())
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewVersion()
	if resp.GetVersion() != nil {
		ret.setMsg(resp.GetVersion())
		return ret, nil
	}

	return ret, nil
}

func (api *gosnappiApi) httpSetConfig(config Config) (Warning, error) {
	configJson, err := config.Marshal().ToJson()
	if err != nil {
		return nil, err
	}
	parentCtx := api.Telemetry().getRootContext()
	newCtx, span := api.Telemetry().NewSpan(parentCtx, "SetConfig", trace.WithSpanKind(trace.SpanKindClient))
	defer api.Telemetry().CloseSpan(span)
	resp, err := api.httpSendRecv(newCtx, "config", configJson, "POST")

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
		err := fromHttpError(resp.StatusCode, bodyBytes)
		api.Telemetry().SetSpanStatus(span, codes.Error, err.Error())
		return nil, err
	}
}

func (api *gosnappiApi) httpGetConfig() (Config, error) {
	parentCtx := api.Telemetry().getRootContext()
	newCtx, span := api.Telemetry().NewSpan(parentCtx, "GetConfig", trace.WithSpanKind(trace.SpanKindClient))
	defer api.Telemetry().CloseSpan(span)
	resp, err := api.httpSendRecv(newCtx, "config", "", "GET")
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
		err := fromHttpError(resp.StatusCode, bodyBytes)
		api.Telemetry().SetSpanStatus(span, codes.Error, err.Error())
		return nil, err
	}
}

func (api *gosnappiApi) httpUpdateConfig(configUpdate ConfigUpdate) (Warning, error) {
	configUpdateJson, err := configUpdate.Marshal().ToJson()
	if err != nil {
		return nil, err
	}
	parentCtx := api.Telemetry().getRootContext()
	newCtx, span := api.Telemetry().NewSpan(parentCtx, "UpdateConfig", trace.WithSpanKind(trace.SpanKindClient))
	defer api.Telemetry().CloseSpan(span)
	resp, err := api.httpSendRecv(newCtx, "config", configUpdateJson, "PATCH")

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
		err := fromHttpError(resp.StatusCode, bodyBytes)
		api.Telemetry().SetSpanStatus(span, codes.Error, err.Error())
		return nil, err
	}
}

func (api *gosnappiApi) httpAppendConfig(configAppend ConfigAppend) (Warning, error) {
	configAppendJson, err := configAppend.Marshal().ToJson()
	if err != nil {
		return nil, err
	}
	parentCtx := api.Telemetry().getRootContext()
	newCtx, span := api.Telemetry().NewSpan(parentCtx, "AppendConfig", trace.WithSpanKind(trace.SpanKindClient))
	defer api.Telemetry().CloseSpan(span)
	resp, err := api.httpSendRecv(newCtx, "config/append", configAppendJson, "PATCH")

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
		err := fromHttpError(resp.StatusCode, bodyBytes)
		api.Telemetry().SetSpanStatus(span, codes.Error, err.Error())
		return nil, err
	}
}

func (api *gosnappiApi) httpDeleteConfig(configDelete ConfigDelete) (Warning, error) {
	configDeleteJson, err := configDelete.Marshal().ToJson()
	if err != nil {
		return nil, err
	}
	parentCtx := api.Telemetry().getRootContext()
	newCtx, span := api.Telemetry().NewSpan(parentCtx, "DeleteConfig", trace.WithSpanKind(trace.SpanKindClient))
	defer api.Telemetry().CloseSpan(span)
	resp, err := api.httpSendRecv(newCtx, "config/delete", configDeleteJson, "PATCH")

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
		err := fromHttpError(resp.StatusCode, bodyBytes)
		api.Telemetry().SetSpanStatus(span, codes.Error, err.Error())
		return nil, err
	}
}

func (api *gosnappiApi) httpSetControlState(controlState ControlState) (Warning, error) {
	controlStateJson, err := controlState.Marshal().ToJson()
	if err != nil {
		return nil, err
	}
	parentCtx := api.Telemetry().getRootContext()
	newCtx, span := api.Telemetry().NewSpan(parentCtx, "SetControlState", trace.WithSpanKind(trace.SpanKindClient))
	defer api.Telemetry().CloseSpan(span)
	resp, err := api.httpSendRecv(newCtx, "control/state", controlStateJson, "POST")

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
		err := fromHttpError(resp.StatusCode, bodyBytes)
		api.Telemetry().SetSpanStatus(span, codes.Error, err.Error())
		return nil, err
	}
}

func (api *gosnappiApi) httpSetControlAction(controlAction ControlAction) (ControlActionResponse, error) {
	controlActionJson, err := controlAction.Marshal().ToJson()
	if err != nil {
		return nil, err
	}
	parentCtx := api.Telemetry().getRootContext()
	newCtx, span := api.Telemetry().NewSpan(parentCtx, "SetControlAction", trace.WithSpanKind(trace.SpanKindClient))
	defer api.Telemetry().CloseSpan(span)
	resp, err := api.httpSendRecv(newCtx, "control/action", controlActionJson, "POST")

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
		err := fromHttpError(resp.StatusCode, bodyBytes)
		api.Telemetry().SetSpanStatus(span, codes.Error, err.Error())
		return nil, err
	}
}

func (api *gosnappiApi) httpGetMetrics(metricsRequest MetricsRequest) (MetricsResponse, error) {
	metricsRequestJson, err := metricsRequest.Marshal().ToJson()
	if err != nil {
		return nil, err
	}
	parentCtx := api.Telemetry().getRootContext()
	newCtx, span := api.Telemetry().NewSpan(parentCtx, "GetMetrics", trace.WithSpanKind(trace.SpanKindClient))
	defer api.Telemetry().CloseSpan(span)
	resp, err := api.httpSendRecv(newCtx, "monitor/metrics", metricsRequestJson, "POST")

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
		err := fromHttpError(resp.StatusCode, bodyBytes)
		api.Telemetry().SetSpanStatus(span, codes.Error, err.Error())
		return nil, err
	}
}

func (api *gosnappiApi) httpGetStates(statesRequest StatesRequest) (StatesResponse, error) {
	statesRequestJson, err := statesRequest.Marshal().ToJson()
	if err != nil {
		return nil, err
	}
	parentCtx := api.Telemetry().getRootContext()
	newCtx, span := api.Telemetry().NewSpan(parentCtx, "GetStates", trace.WithSpanKind(trace.SpanKindClient))
	defer api.Telemetry().CloseSpan(span)
	resp, err := api.httpSendRecv(newCtx, "monitor/states", statesRequestJson, "POST")

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
		err := fromHttpError(resp.StatusCode, bodyBytes)
		api.Telemetry().SetSpanStatus(span, codes.Error, err.Error())
		return nil, err
	}
}

func (api *gosnappiApi) httpGetCapture(captureRequest CaptureRequest) ([]byte, error) {
	captureRequestJson, err := captureRequest.Marshal().ToJson()
	if err != nil {
		return nil, err
	}
	parentCtx := api.Telemetry().getRootContext()
	newCtx, span := api.Telemetry().NewSpan(parentCtx, "GetCapture", trace.WithSpanKind(trace.SpanKindClient))
	defer api.Telemetry().CloseSpan(span)
	resp, err := api.httpSendRecv(newCtx, "monitor/capture", captureRequestJson, "POST")

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
		err := fromHttpError(resp.StatusCode, bodyBytes)
		api.Telemetry().SetSpanStatus(span, codes.Error, err.Error())
		return nil, err
	}
}

func (api *gosnappiApi) httpGetVersion() (Version, error) {
	parentCtx := api.Telemetry().getRootContext()
	newCtx, span := api.Telemetry().NewSpan(parentCtx, "GetVersion", trace.WithSpanKind(trace.SpanKindClient))
	defer api.Telemetry().CloseSpan(span)
	resp, err := api.httpSendRecv(newCtx, "capabilities/version", "", "GET")
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
		err := fromHttpError(resp.StatusCode, bodyBytes)
		api.Telemetry().SetSpanStatus(span, codes.Error, err.Error())
		return nil, err
	}
}
