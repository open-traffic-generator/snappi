package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RoCEv2Metric *****
type roCEv2Metric struct {
	validation
	obj          *otg.RoCEv2Metric
	marshaller   marshalRoCEv2Metric
	unMarshaller unMarshalRoCEv2Metric
}

func NewRoCEv2Metric() RoCEv2Metric {
	obj := roCEv2Metric{obj: &otg.RoCEv2Metric{}}
	obj.setDefault()
	return &obj
}

func (obj *roCEv2Metric) msg() *otg.RoCEv2Metric {
	return obj.obj
}

func (obj *roCEv2Metric) setMsg(msg *otg.RoCEv2Metric) RoCEv2Metric {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalroCEv2Metric struct {
	obj *roCEv2Metric
}

type marshalRoCEv2Metric interface {
	// ToProto marshals RoCEv2Metric to protobuf object *otg.RoCEv2Metric
	ToProto() (*otg.RoCEv2Metric, error)
	// ToPbText marshals RoCEv2Metric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RoCEv2Metric to YAML text
	ToYaml() (string, error)
	// ToJson marshals RoCEv2Metric to JSON text
	ToJson() (string, error)
}

type unMarshalroCEv2Metric struct {
	obj *roCEv2Metric
}

type unMarshalRoCEv2Metric interface {
	// FromProto unmarshals RoCEv2Metric from protobuf object *otg.RoCEv2Metric
	FromProto(msg *otg.RoCEv2Metric) (RoCEv2Metric, error)
	// FromPbText unmarshals RoCEv2Metric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RoCEv2Metric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RoCEv2Metric from JSON text
	FromJson(value string) error
}

func (obj *roCEv2Metric) Marshal() marshalRoCEv2Metric {
	if obj.marshaller == nil {
		obj.marshaller = &marshalroCEv2Metric{obj: obj}
	}
	return obj.marshaller
}

func (obj *roCEv2Metric) Unmarshal() unMarshalRoCEv2Metric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalroCEv2Metric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalroCEv2Metric) ToProto() (*otg.RoCEv2Metric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalroCEv2Metric) FromProto(msg *otg.RoCEv2Metric) (RoCEv2Metric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalroCEv2Metric) ToPbText() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (m *unMarshalroCEv2Metric) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalroCEv2Metric) ToYaml() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (m *unMarshalroCEv2Metric) FromYaml(value string) error {
	if value == "" {
		value = "{}"
	}
	data, err := yaml.YAMLToJSON([]byte(value))
	if err != nil {
		return err
	}
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	uError := opts.Unmarshal([]byte(data), m.obj.msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalroCEv2Metric) ToJson() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (m *unMarshalroCEv2Metric) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), m.obj.msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *roCEv2Metric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *roCEv2Metric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *roCEv2Metric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *roCEv2Metric) Clone() (RoCEv2Metric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRoCEv2Metric()
	data, err := proto.Marshal(obj.msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// RoCEv2Metric is roCEv2 per peer statistics information.
type RoCEv2Metric interface {
	Validation
	// msg marshals RoCEv2Metric to protobuf object *otg.RoCEv2Metric
	// and doesn't set defaults
	msg() *otg.RoCEv2Metric
	// setMsg unmarshals RoCEv2Metric from protobuf object *otg.RoCEv2Metric
	// and doesn't set defaults
	setMsg(*otg.RoCEv2Metric) RoCEv2Metric
	// provides marshal interface
	Marshal() marshalRoCEv2Metric
	// provides unmarshal interface
	Unmarshal() unMarshalRoCEv2Metric
	// validate validates RoCEv2Metric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RoCEv2Metric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in RoCEv2Metric.
	Name() string
	// SetName assigns string provided by user to RoCEv2Metric
	SetName(value string) RoCEv2Metric
	// HasName checks if Name has been set in RoCEv2Metric
	HasName() bool
	// SessionState returns RoCEv2MetricSessionStateEnum, set in RoCEv2Metric
	SessionState() RoCEv2MetricSessionStateEnum
	// SetSessionState assigns RoCEv2MetricSessionStateEnum provided by user to RoCEv2Metric
	SetSessionState(value RoCEv2MetricSessionStateEnum) RoCEv2Metric
	// HasSessionState checks if SessionState has been set in RoCEv2Metric
	HasSessionState() bool
	// QpConfigured returns uint64, set in RoCEv2Metric.
	QpConfigured() uint64
	// SetQpConfigured assigns uint64 provided by user to RoCEv2Metric
	SetQpConfigured(value uint64) RoCEv2Metric
	// HasQpConfigured checks if QpConfigured has been set in RoCEv2Metric
	HasQpConfigured() bool
	// QpUp returns uint64, set in RoCEv2Metric.
	QpUp() uint64
	// SetQpUp assigns uint64 provided by user to RoCEv2Metric
	SetQpUp(value uint64) RoCEv2Metric
	// HasQpUp checks if QpUp has been set in RoCEv2Metric
	HasQpUp() bool
	// QpDown returns uint64, set in RoCEv2Metric.
	QpDown() uint64
	// SetQpDown assigns uint64 provided by user to RoCEv2Metric
	SetQpDown(value uint64) RoCEv2Metric
	// HasQpDown checks if QpDown has been set in RoCEv2Metric
	HasQpDown() bool
	// ConnectRequestTx returns uint64, set in RoCEv2Metric.
	ConnectRequestTx() uint64
	// SetConnectRequestTx assigns uint64 provided by user to RoCEv2Metric
	SetConnectRequestTx(value uint64) RoCEv2Metric
	// HasConnectRequestTx checks if ConnectRequestTx has been set in RoCEv2Metric
	HasConnectRequestTx() bool
	// ConnectRequestRx returns uint64, set in RoCEv2Metric.
	ConnectRequestRx() uint64
	// SetConnectRequestRx assigns uint64 provided by user to RoCEv2Metric
	SetConnectRequestRx(value uint64) RoCEv2Metric
	// HasConnectRequestRx checks if ConnectRequestRx has been set in RoCEv2Metric
	HasConnectRequestRx() bool
	// ConnectReplyTx returns uint64, set in RoCEv2Metric.
	ConnectReplyTx() uint64
	// SetConnectReplyTx assigns uint64 provided by user to RoCEv2Metric
	SetConnectReplyTx(value uint64) RoCEv2Metric
	// HasConnectReplyTx checks if ConnectReplyTx has been set in RoCEv2Metric
	HasConnectReplyTx() bool
	// ConnectReplyRx returns uint64, set in RoCEv2Metric.
	ConnectReplyRx() uint64
	// SetConnectReplyRx assigns uint64 provided by user to RoCEv2Metric
	SetConnectReplyRx(value uint64) RoCEv2Metric
	// HasConnectReplyRx checks if ConnectReplyRx has been set in RoCEv2Metric
	HasConnectReplyRx() bool
	// ReadyTx returns uint64, set in RoCEv2Metric.
	ReadyTx() uint64
	// SetReadyTx assigns uint64 provided by user to RoCEv2Metric
	SetReadyTx(value uint64) RoCEv2Metric
	// HasReadyTx checks if ReadyTx has been set in RoCEv2Metric
	HasReadyTx() bool
	// ReadyRx returns uint64, set in RoCEv2Metric.
	ReadyRx() uint64
	// SetReadyRx assigns uint64 provided by user to RoCEv2Metric
	SetReadyRx(value uint64) RoCEv2Metric
	// HasReadyRx checks if ReadyRx has been set in RoCEv2Metric
	HasReadyRx() bool
	// DisconnectRequestTx returns uint64, set in RoCEv2Metric.
	DisconnectRequestTx() uint64
	// SetDisconnectRequestTx assigns uint64 provided by user to RoCEv2Metric
	SetDisconnectRequestTx(value uint64) RoCEv2Metric
	// HasDisconnectRequestTx checks if DisconnectRequestTx has been set in RoCEv2Metric
	HasDisconnectRequestTx() bool
	// DisconnectRequestRx returns uint64, set in RoCEv2Metric.
	DisconnectRequestRx() uint64
	// SetDisconnectRequestRx assigns uint64 provided by user to RoCEv2Metric
	SetDisconnectRequestRx(value uint64) RoCEv2Metric
	// HasDisconnectRequestRx checks if DisconnectRequestRx has been set in RoCEv2Metric
	HasDisconnectRequestRx() bool
	// DisconnectReplyTx returns uint64, set in RoCEv2Metric.
	DisconnectReplyTx() uint64
	// SetDisconnectReplyTx assigns uint64 provided by user to RoCEv2Metric
	SetDisconnectReplyTx(value uint64) RoCEv2Metric
	// HasDisconnectReplyTx checks if DisconnectReplyTx has been set in RoCEv2Metric
	HasDisconnectReplyTx() bool
	// DisconnectReplyRx returns uint64, set in RoCEv2Metric.
	DisconnectReplyRx() uint64
	// SetDisconnectReplyRx assigns uint64 provided by user to RoCEv2Metric
	SetDisconnectReplyRx(value uint64) RoCEv2Metric
	// HasDisconnectReplyRx checks if DisconnectReplyRx has been set in RoCEv2Metric
	HasDisconnectReplyRx() bool
	// RejectTx returns uint64, set in RoCEv2Metric.
	RejectTx() uint64
	// SetRejectTx assigns uint64 provided by user to RoCEv2Metric
	SetRejectTx(value uint64) RoCEv2Metric
	// HasRejectTx checks if RejectTx has been set in RoCEv2Metric
	HasRejectTx() bool
	// UnknownMsgRx returns uint64, set in RoCEv2Metric.
	UnknownMsgRx() uint64
	// SetUnknownMsgRx assigns uint64 provided by user to RoCEv2Metric
	SetUnknownMsgRx(value uint64) RoCEv2Metric
	// HasUnknownMsgRx checks if UnknownMsgRx has been set in RoCEv2Metric
	HasUnknownMsgRx() bool
}

// The name of a configured RoCEv2 peer.
// Name returns a string
func (obj *roCEv2Metric) Name() string {

	return *obj.obj.Name

}

// The name of a configured RoCEv2 peer.
// Name returns a string
func (obj *roCEv2Metric) HasName() bool {
	return obj.obj.Name != nil
}

// The name of a configured RoCEv2 peer.
// SetName sets the string value in the RoCEv2Metric object
func (obj *roCEv2Metric) SetName(value string) RoCEv2Metric {

	obj.obj.Name = &value
	return obj
}

type RoCEv2MetricSessionStateEnum string

// Enum of SessionState on RoCEv2Metric
var RoCEv2MetricSessionState = struct {
	UP   RoCEv2MetricSessionStateEnum
	DOWN RoCEv2MetricSessionStateEnum
}{
	UP:   RoCEv2MetricSessionStateEnum("up"),
	DOWN: RoCEv2MetricSessionStateEnum("down"),
}

func (obj *roCEv2Metric) SessionState() RoCEv2MetricSessionStateEnum {
	return RoCEv2MetricSessionStateEnum(obj.obj.SessionState.Enum().String())
}

// Session state as up or down. Up refers to an Established state and Down refers to any other state.
// SessionState returns a string
func (obj *roCEv2Metric) HasSessionState() bool {
	return obj.obj.SessionState != nil
}

func (obj *roCEv2Metric) SetSessionState(value RoCEv2MetricSessionStateEnum) RoCEv2Metric {
	intValue, ok := otg.RoCEv2Metric_SessionState_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on RoCEv2MetricSessionStateEnum", string(value)))
		return obj
	}
	enumValue := otg.RoCEv2Metric_SessionState_Enum(intValue)
	obj.obj.SessionState = &enumValue

	return obj
}

// Number of QPs configured on this port.
// QpConfigured returns a uint64
func (obj *roCEv2Metric) QpConfigured() uint64 {

	return *obj.obj.QpConfigured

}

// Number of QPs configured on this port.
// QpConfigured returns a uint64
func (obj *roCEv2Metric) HasQpConfigured() bool {
	return obj.obj.QpConfigured != nil
}

// Number of QPs configured on this port.
// SetQpConfigured sets the uint64 value in the RoCEv2Metric object
func (obj *roCEv2Metric) SetQpConfigured(value uint64) RoCEv2Metric {

	obj.obj.QpConfigured = &value
	return obj
}

// Number of QPs that are in UP state.
// QpUp returns a uint64
func (obj *roCEv2Metric) QpUp() uint64 {

	return *obj.obj.QpUp

}

// Number of QPs that are in UP state.
// QpUp returns a uint64
func (obj *roCEv2Metric) HasQpUp() bool {
	return obj.obj.QpUp != nil
}

// Number of QPs that are in UP state.
// SetQpUp sets the uint64 value in the RoCEv2Metric object
func (obj *roCEv2Metric) SetQpUp(value uint64) RoCEv2Metric {

	obj.obj.QpUp = &value
	return obj
}

// Number of QPs that have not come UP.
// QpDown returns a uint64
func (obj *roCEv2Metric) QpDown() uint64 {

	return *obj.obj.QpDown

}

// Number of QPs that have not come UP.
// QpDown returns a uint64
func (obj *roCEv2Metric) HasQpDown() bool {
	return obj.obj.QpDown != nil
}

// Number of QPs that have not come UP.
// SetQpDown sets the uint64 value in the RoCEv2Metric object
func (obj *roCEv2Metric) SetQpDown(value uint64) RoCEv2Metric {

	obj.obj.QpDown = &value
	return obj
}

// Number of REQ Message Transmitted.
// ConnectRequestTx returns a uint64
func (obj *roCEv2Metric) ConnectRequestTx() uint64 {

	return *obj.obj.ConnectRequestTx

}

// Number of REQ Message Transmitted.
// ConnectRequestTx returns a uint64
func (obj *roCEv2Metric) HasConnectRequestTx() bool {
	return obj.obj.ConnectRequestTx != nil
}

// Number of REQ Message Transmitted.
// SetConnectRequestTx sets the uint64 value in the RoCEv2Metric object
func (obj *roCEv2Metric) SetConnectRequestTx(value uint64) RoCEv2Metric {

	obj.obj.ConnectRequestTx = &value
	return obj
}

// Number of REQ Message Received.
// ConnectRequestRx returns a uint64
func (obj *roCEv2Metric) ConnectRequestRx() uint64 {

	return *obj.obj.ConnectRequestRx

}

// Number of REQ Message Received.
// ConnectRequestRx returns a uint64
func (obj *roCEv2Metric) HasConnectRequestRx() bool {
	return obj.obj.ConnectRequestRx != nil
}

// Number of REQ Message Received.
// SetConnectRequestRx sets the uint64 value in the RoCEv2Metric object
func (obj *roCEv2Metric) SetConnectRequestRx(value uint64) RoCEv2Metric {

	obj.obj.ConnectRequestRx = &value
	return obj
}

// Number of REP Message Transmitted.
// ConnectReplyTx returns a uint64
func (obj *roCEv2Metric) ConnectReplyTx() uint64 {

	return *obj.obj.ConnectReplyTx

}

// Number of REP Message Transmitted.
// ConnectReplyTx returns a uint64
func (obj *roCEv2Metric) HasConnectReplyTx() bool {
	return obj.obj.ConnectReplyTx != nil
}

// Number of REP Message Transmitted.
// SetConnectReplyTx sets the uint64 value in the RoCEv2Metric object
func (obj *roCEv2Metric) SetConnectReplyTx(value uint64) RoCEv2Metric {

	obj.obj.ConnectReplyTx = &value
	return obj
}

// Number of REP Message Received.
// ConnectReplyRx returns a uint64
func (obj *roCEv2Metric) ConnectReplyRx() uint64 {

	return *obj.obj.ConnectReplyRx

}

// Number of REP Message Received.
// ConnectReplyRx returns a uint64
func (obj *roCEv2Metric) HasConnectReplyRx() bool {
	return obj.obj.ConnectReplyRx != nil
}

// Number of REP Message Received.
// SetConnectReplyRx sets the uint64 value in the RoCEv2Metric object
func (obj *roCEv2Metric) SetConnectReplyRx(value uint64) RoCEv2Metric {

	obj.obj.ConnectReplyRx = &value
	return obj
}

// Number of RTU Message Transmitted.
// ReadyTx returns a uint64
func (obj *roCEv2Metric) ReadyTx() uint64 {

	return *obj.obj.ReadyTx

}

// Number of RTU Message Transmitted.
// ReadyTx returns a uint64
func (obj *roCEv2Metric) HasReadyTx() bool {
	return obj.obj.ReadyTx != nil
}

// Number of RTU Message Transmitted.
// SetReadyTx sets the uint64 value in the RoCEv2Metric object
func (obj *roCEv2Metric) SetReadyTx(value uint64) RoCEv2Metric {

	obj.obj.ReadyTx = &value
	return obj
}

// Number of RTU Message Received.
// ReadyRx returns a uint64
func (obj *roCEv2Metric) ReadyRx() uint64 {

	return *obj.obj.ReadyRx

}

// Number of RTU Message Received.
// ReadyRx returns a uint64
func (obj *roCEv2Metric) HasReadyRx() bool {
	return obj.obj.ReadyRx != nil
}

// Number of RTU Message Received.
// SetReadyRx sets the uint64 value in the RoCEv2Metric object
func (obj *roCEv2Metric) SetReadyRx(value uint64) RoCEv2Metric {

	obj.obj.ReadyRx = &value
	return obj
}

// Number of DREQ Message Transmitted.
// DisconnectRequestTx returns a uint64
func (obj *roCEv2Metric) DisconnectRequestTx() uint64 {

	return *obj.obj.DisconnectRequestTx

}

// Number of DREQ Message Transmitted.
// DisconnectRequestTx returns a uint64
func (obj *roCEv2Metric) HasDisconnectRequestTx() bool {
	return obj.obj.DisconnectRequestTx != nil
}

// Number of DREQ Message Transmitted.
// SetDisconnectRequestTx sets the uint64 value in the RoCEv2Metric object
func (obj *roCEv2Metric) SetDisconnectRequestTx(value uint64) RoCEv2Metric {

	obj.obj.DisconnectRequestTx = &value
	return obj
}

// Number of DREQ Message Received.
// DisconnectRequestRx returns a uint64
func (obj *roCEv2Metric) DisconnectRequestRx() uint64 {

	return *obj.obj.DisconnectRequestRx

}

// Number of DREQ Message Received.
// DisconnectRequestRx returns a uint64
func (obj *roCEv2Metric) HasDisconnectRequestRx() bool {
	return obj.obj.DisconnectRequestRx != nil
}

// Number of DREQ Message Received.
// SetDisconnectRequestRx sets the uint64 value in the RoCEv2Metric object
func (obj *roCEv2Metric) SetDisconnectRequestRx(value uint64) RoCEv2Metric {

	obj.obj.DisconnectRequestRx = &value
	return obj
}

// Number of DREP Message Transmitted.
// DisconnectReplyTx returns a uint64
func (obj *roCEv2Metric) DisconnectReplyTx() uint64 {

	return *obj.obj.DisconnectReplyTx

}

// Number of DREP Message Transmitted.
// DisconnectReplyTx returns a uint64
func (obj *roCEv2Metric) HasDisconnectReplyTx() bool {
	return obj.obj.DisconnectReplyTx != nil
}

// Number of DREP Message Transmitted.
// SetDisconnectReplyTx sets the uint64 value in the RoCEv2Metric object
func (obj *roCEv2Metric) SetDisconnectReplyTx(value uint64) RoCEv2Metric {

	obj.obj.DisconnectReplyTx = &value
	return obj
}

// Number of DREP Message Received.
// DisconnectReplyRx returns a uint64
func (obj *roCEv2Metric) DisconnectReplyRx() uint64 {

	return *obj.obj.DisconnectReplyRx

}

// Number of DREP Message Received.
// DisconnectReplyRx returns a uint64
func (obj *roCEv2Metric) HasDisconnectReplyRx() bool {
	return obj.obj.DisconnectReplyRx != nil
}

// Number of DREP Message Received.
// SetDisconnectReplyRx sets the uint64 value in the RoCEv2Metric object
func (obj *roCEv2Metric) SetDisconnectReplyRx(value uint64) RoCEv2Metric {

	obj.obj.DisconnectReplyRx = &value
	return obj
}

// Number of REJ Message Transmitted.
// RejectTx returns a uint64
func (obj *roCEv2Metric) RejectTx() uint64 {

	return *obj.obj.RejectTx

}

// Number of REJ Message Transmitted.
// RejectTx returns a uint64
func (obj *roCEv2Metric) HasRejectTx() bool {
	return obj.obj.RejectTx != nil
}

// Number of REJ Message Transmitted.
// SetRejectTx sets the uint64 value in the RoCEv2Metric object
func (obj *roCEv2Metric) SetRejectTx(value uint64) RoCEv2Metric {

	obj.obj.RejectTx = &value
	return obj
}

// Number of Unknow Message Received.
// UnknownMsgRx returns a uint64
func (obj *roCEv2Metric) UnknownMsgRx() uint64 {

	return *obj.obj.UnknownMsgRx

}

// Number of Unknow Message Received.
// UnknownMsgRx returns a uint64
func (obj *roCEv2Metric) HasUnknownMsgRx() bool {
	return obj.obj.UnknownMsgRx != nil
}

// Number of Unknow Message Received.
// SetUnknownMsgRx sets the uint64 value in the RoCEv2Metric object
func (obj *roCEv2Metric) SetUnknownMsgRx(value uint64) RoCEv2Metric {

	obj.obj.UnknownMsgRx = &value
	return obj
}

func (obj *roCEv2Metric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *roCEv2Metric) setDefault() {

}
