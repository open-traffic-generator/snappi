package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RoCEv2IPv4Metric *****
type roCEv2IPv4Metric struct {
	validation
	obj          *otg.RoCEv2IPv4Metric
	marshaller   marshalRoCEv2IPv4Metric
	unMarshaller unMarshalRoCEv2IPv4Metric
}

func NewRoCEv2IPv4Metric() RoCEv2IPv4Metric {
	obj := roCEv2IPv4Metric{obj: &otg.RoCEv2IPv4Metric{}}
	obj.setDefault()
	return &obj
}

func (obj *roCEv2IPv4Metric) msg() *otg.RoCEv2IPv4Metric {
	return obj.obj
}

func (obj *roCEv2IPv4Metric) setMsg(msg *otg.RoCEv2IPv4Metric) RoCEv2IPv4Metric {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalroCEv2IPv4Metric struct {
	obj *roCEv2IPv4Metric
}

type marshalRoCEv2IPv4Metric interface {
	// ToProto marshals RoCEv2IPv4Metric to protobuf object *otg.RoCEv2IPv4Metric
	ToProto() (*otg.RoCEv2IPv4Metric, error)
	// ToPbText marshals RoCEv2IPv4Metric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RoCEv2IPv4Metric to YAML text
	ToYaml() (string, error)
	// ToJson marshals RoCEv2IPv4Metric to JSON text
	ToJson() (string, error)
}

type unMarshalroCEv2IPv4Metric struct {
	obj *roCEv2IPv4Metric
}

type unMarshalRoCEv2IPv4Metric interface {
	// FromProto unmarshals RoCEv2IPv4Metric from protobuf object *otg.RoCEv2IPv4Metric
	FromProto(msg *otg.RoCEv2IPv4Metric) (RoCEv2IPv4Metric, error)
	// FromPbText unmarshals RoCEv2IPv4Metric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RoCEv2IPv4Metric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RoCEv2IPv4Metric from JSON text
	FromJson(value string) error
}

func (obj *roCEv2IPv4Metric) Marshal() marshalRoCEv2IPv4Metric {
	if obj.marshaller == nil {
		obj.marshaller = &marshalroCEv2IPv4Metric{obj: obj}
	}
	return obj.marshaller
}

func (obj *roCEv2IPv4Metric) Unmarshal() unMarshalRoCEv2IPv4Metric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalroCEv2IPv4Metric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalroCEv2IPv4Metric) ToProto() (*otg.RoCEv2IPv4Metric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalroCEv2IPv4Metric) FromProto(msg *otg.RoCEv2IPv4Metric) (RoCEv2IPv4Metric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalroCEv2IPv4Metric) ToPbText() (string, error) {
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

func (m *unMarshalroCEv2IPv4Metric) FromPbText(value string) error {
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

func (m *marshalroCEv2IPv4Metric) ToYaml() (string, error) {
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

func (m *unMarshalroCEv2IPv4Metric) FromYaml(value string) error {
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

func (m *marshalroCEv2IPv4Metric) ToJson() (string, error) {
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

func (m *unMarshalroCEv2IPv4Metric) FromJson(value string) error {
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

func (obj *roCEv2IPv4Metric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *roCEv2IPv4Metric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *roCEv2IPv4Metric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *roCEv2IPv4Metric) Clone() (RoCEv2IPv4Metric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRoCEv2IPv4Metric()
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

// RoCEv2IPv4Metric is roCEv2 per peer statistics information.
type RoCEv2IPv4Metric interface {
	Validation
	// msg marshals RoCEv2IPv4Metric to protobuf object *otg.RoCEv2IPv4Metric
	// and doesn't set defaults
	msg() *otg.RoCEv2IPv4Metric
	// setMsg unmarshals RoCEv2IPv4Metric from protobuf object *otg.RoCEv2IPv4Metric
	// and doesn't set defaults
	setMsg(*otg.RoCEv2IPv4Metric) RoCEv2IPv4Metric
	// provides marshal interface
	Marshal() marshalRoCEv2IPv4Metric
	// provides unmarshal interface
	Unmarshal() unMarshalRoCEv2IPv4Metric
	// validate validates RoCEv2IPv4Metric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RoCEv2IPv4Metric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in RoCEv2IPv4Metric.
	Name() string
	// SetName assigns string provided by user to RoCEv2IPv4Metric
	SetName(value string) RoCEv2IPv4Metric
	// HasName checks if Name has been set in RoCEv2IPv4Metric
	HasName() bool
	// SessionState returns RoCEv2IPv4MetricSessionStateEnum, set in RoCEv2IPv4Metric
	SessionState() RoCEv2IPv4MetricSessionStateEnum
	// SetSessionState assigns RoCEv2IPv4MetricSessionStateEnum provided by user to RoCEv2IPv4Metric
	SetSessionState(value RoCEv2IPv4MetricSessionStateEnum) RoCEv2IPv4Metric
	// HasSessionState checks if SessionState has been set in RoCEv2IPv4Metric
	HasSessionState() bool
	// QpConfigured returns uint64, set in RoCEv2IPv4Metric.
	QpConfigured() uint64
	// SetQpConfigured assigns uint64 provided by user to RoCEv2IPv4Metric
	SetQpConfigured(value uint64) RoCEv2IPv4Metric
	// HasQpConfigured checks if QpConfigured has been set in RoCEv2IPv4Metric
	HasQpConfigured() bool
	// QpUp returns uint64, set in RoCEv2IPv4Metric.
	QpUp() uint64
	// SetQpUp assigns uint64 provided by user to RoCEv2IPv4Metric
	SetQpUp(value uint64) RoCEv2IPv4Metric
	// HasQpUp checks if QpUp has been set in RoCEv2IPv4Metric
	HasQpUp() bool
	// QpDown returns uint64, set in RoCEv2IPv4Metric.
	QpDown() uint64
	// SetQpDown assigns uint64 provided by user to RoCEv2IPv4Metric
	SetQpDown(value uint64) RoCEv2IPv4Metric
	// HasQpDown checks if QpDown has been set in RoCEv2IPv4Metric
	HasQpDown() bool
	// ConnectRequestTx returns uint64, set in RoCEv2IPv4Metric.
	ConnectRequestTx() uint64
	// SetConnectRequestTx assigns uint64 provided by user to RoCEv2IPv4Metric
	SetConnectRequestTx(value uint64) RoCEv2IPv4Metric
	// HasConnectRequestTx checks if ConnectRequestTx has been set in RoCEv2IPv4Metric
	HasConnectRequestTx() bool
	// ConnectRequestRx returns uint64, set in RoCEv2IPv4Metric.
	ConnectRequestRx() uint64
	// SetConnectRequestRx assigns uint64 provided by user to RoCEv2IPv4Metric
	SetConnectRequestRx(value uint64) RoCEv2IPv4Metric
	// HasConnectRequestRx checks if ConnectRequestRx has been set in RoCEv2IPv4Metric
	HasConnectRequestRx() bool
	// ConnectReplyTx returns uint64, set in RoCEv2IPv4Metric.
	ConnectReplyTx() uint64
	// SetConnectReplyTx assigns uint64 provided by user to RoCEv2IPv4Metric
	SetConnectReplyTx(value uint64) RoCEv2IPv4Metric
	// HasConnectReplyTx checks if ConnectReplyTx has been set in RoCEv2IPv4Metric
	HasConnectReplyTx() bool
	// ConnectReplyRx returns uint64, set in RoCEv2IPv4Metric.
	ConnectReplyRx() uint64
	// SetConnectReplyRx assigns uint64 provided by user to RoCEv2IPv4Metric
	SetConnectReplyRx(value uint64) RoCEv2IPv4Metric
	// HasConnectReplyRx checks if ConnectReplyRx has been set in RoCEv2IPv4Metric
	HasConnectReplyRx() bool
	// ReadyTx returns uint64, set in RoCEv2IPv4Metric.
	ReadyTx() uint64
	// SetReadyTx assigns uint64 provided by user to RoCEv2IPv4Metric
	SetReadyTx(value uint64) RoCEv2IPv4Metric
	// HasReadyTx checks if ReadyTx has been set in RoCEv2IPv4Metric
	HasReadyTx() bool
	// ReadyRx returns uint64, set in RoCEv2IPv4Metric.
	ReadyRx() uint64
	// SetReadyRx assigns uint64 provided by user to RoCEv2IPv4Metric
	SetReadyRx(value uint64) RoCEv2IPv4Metric
	// HasReadyRx checks if ReadyRx has been set in RoCEv2IPv4Metric
	HasReadyRx() bool
	// DisconnectRequestTx returns uint64, set in RoCEv2IPv4Metric.
	DisconnectRequestTx() uint64
	// SetDisconnectRequestTx assigns uint64 provided by user to RoCEv2IPv4Metric
	SetDisconnectRequestTx(value uint64) RoCEv2IPv4Metric
	// HasDisconnectRequestTx checks if DisconnectRequestTx has been set in RoCEv2IPv4Metric
	HasDisconnectRequestTx() bool
	// DisconnectRequestRx returns uint64, set in RoCEv2IPv4Metric.
	DisconnectRequestRx() uint64
	// SetDisconnectRequestRx assigns uint64 provided by user to RoCEv2IPv4Metric
	SetDisconnectRequestRx(value uint64) RoCEv2IPv4Metric
	// HasDisconnectRequestRx checks if DisconnectRequestRx has been set in RoCEv2IPv4Metric
	HasDisconnectRequestRx() bool
	// DisconnectReplyTx returns uint64, set in RoCEv2IPv4Metric.
	DisconnectReplyTx() uint64
	// SetDisconnectReplyTx assigns uint64 provided by user to RoCEv2IPv4Metric
	SetDisconnectReplyTx(value uint64) RoCEv2IPv4Metric
	// HasDisconnectReplyTx checks if DisconnectReplyTx has been set in RoCEv2IPv4Metric
	HasDisconnectReplyTx() bool
	// DisconnectReplyRx returns uint64, set in RoCEv2IPv4Metric.
	DisconnectReplyRx() uint64
	// SetDisconnectReplyRx assigns uint64 provided by user to RoCEv2IPv4Metric
	SetDisconnectReplyRx(value uint64) RoCEv2IPv4Metric
	// HasDisconnectReplyRx checks if DisconnectReplyRx has been set in RoCEv2IPv4Metric
	HasDisconnectReplyRx() bool
	// RejectTx returns uint64, set in RoCEv2IPv4Metric.
	RejectTx() uint64
	// SetRejectTx assigns uint64 provided by user to RoCEv2IPv4Metric
	SetRejectTx(value uint64) RoCEv2IPv4Metric
	// HasRejectTx checks if RejectTx has been set in RoCEv2IPv4Metric
	HasRejectTx() bool
	// UnknownMsgRx returns uint64, set in RoCEv2IPv4Metric.
	UnknownMsgRx() uint64
	// SetUnknownMsgRx assigns uint64 provided by user to RoCEv2IPv4Metric
	SetUnknownMsgRx(value uint64) RoCEv2IPv4Metric
	// HasUnknownMsgRx checks if UnknownMsgRx has been set in RoCEv2IPv4Metric
	HasUnknownMsgRx() bool
}

// The name of a configured RoCEv2 peer.
// Name returns a string
func (obj *roCEv2IPv4Metric) Name() string {

	return *obj.obj.Name

}

// The name of a configured RoCEv2 peer.
// Name returns a string
func (obj *roCEv2IPv4Metric) HasName() bool {
	return obj.obj.Name != nil
}

// The name of a configured RoCEv2 peer.
// SetName sets the string value in the RoCEv2IPv4Metric object
func (obj *roCEv2IPv4Metric) SetName(value string) RoCEv2IPv4Metric {

	obj.obj.Name = &value
	return obj
}

type RoCEv2IPv4MetricSessionStateEnum string

// Enum of SessionState on RoCEv2IPv4Metric
var RoCEv2IPv4MetricSessionState = struct {
	UP   RoCEv2IPv4MetricSessionStateEnum
	DOWN RoCEv2IPv4MetricSessionStateEnum
}{
	UP:   RoCEv2IPv4MetricSessionStateEnum("up"),
	DOWN: RoCEv2IPv4MetricSessionStateEnum("down"),
}

func (obj *roCEv2IPv4Metric) SessionState() RoCEv2IPv4MetricSessionStateEnum {
	return RoCEv2IPv4MetricSessionStateEnum(obj.obj.SessionState.Enum().String())
}

// Session state as up or down. Up refers to an Established state and Down refers to any other state.
// SessionState returns a string
func (obj *roCEv2IPv4Metric) HasSessionState() bool {
	return obj.obj.SessionState != nil
}

func (obj *roCEv2IPv4Metric) SetSessionState(value RoCEv2IPv4MetricSessionStateEnum) RoCEv2IPv4Metric {
	intValue, ok := otg.RoCEv2IPv4Metric_SessionState_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on RoCEv2IPv4MetricSessionStateEnum", string(value)))
		return obj
	}
	enumValue := otg.RoCEv2IPv4Metric_SessionState_Enum(intValue)
	obj.obj.SessionState = &enumValue

	return obj
}

// Number of QPs configured on this port.
// QpConfigured returns a uint64
func (obj *roCEv2IPv4Metric) QpConfigured() uint64 {

	return *obj.obj.QpConfigured

}

// Number of QPs configured on this port.
// QpConfigured returns a uint64
func (obj *roCEv2IPv4Metric) HasQpConfigured() bool {
	return obj.obj.QpConfigured != nil
}

// Number of QPs configured on this port.
// SetQpConfigured sets the uint64 value in the RoCEv2IPv4Metric object
func (obj *roCEv2IPv4Metric) SetQpConfigured(value uint64) RoCEv2IPv4Metric {

	obj.obj.QpConfigured = &value
	return obj
}

// Number of QPs that are in UP state.
// QpUp returns a uint64
func (obj *roCEv2IPv4Metric) QpUp() uint64 {

	return *obj.obj.QpUp

}

// Number of QPs that are in UP state.
// QpUp returns a uint64
func (obj *roCEv2IPv4Metric) HasQpUp() bool {
	return obj.obj.QpUp != nil
}

// Number of QPs that are in UP state.
// SetQpUp sets the uint64 value in the RoCEv2IPv4Metric object
func (obj *roCEv2IPv4Metric) SetQpUp(value uint64) RoCEv2IPv4Metric {

	obj.obj.QpUp = &value
	return obj
}

// Number of QPs that have not come UP.
// QpDown returns a uint64
func (obj *roCEv2IPv4Metric) QpDown() uint64 {

	return *obj.obj.QpDown

}

// Number of QPs that have not come UP.
// QpDown returns a uint64
func (obj *roCEv2IPv4Metric) HasQpDown() bool {
	return obj.obj.QpDown != nil
}

// Number of QPs that have not come UP.
// SetQpDown sets the uint64 value in the RoCEv2IPv4Metric object
func (obj *roCEv2IPv4Metric) SetQpDown(value uint64) RoCEv2IPv4Metric {

	obj.obj.QpDown = &value
	return obj
}

// Number of REQ Message Transmitted.
// ConnectRequestTx returns a uint64
func (obj *roCEv2IPv4Metric) ConnectRequestTx() uint64 {

	return *obj.obj.ConnectRequestTx

}

// Number of REQ Message Transmitted.
// ConnectRequestTx returns a uint64
func (obj *roCEv2IPv4Metric) HasConnectRequestTx() bool {
	return obj.obj.ConnectRequestTx != nil
}

// Number of REQ Message Transmitted.
// SetConnectRequestTx sets the uint64 value in the RoCEv2IPv4Metric object
func (obj *roCEv2IPv4Metric) SetConnectRequestTx(value uint64) RoCEv2IPv4Metric {

	obj.obj.ConnectRequestTx = &value
	return obj
}

// Number of REQ Message Received.
// ConnectRequestRx returns a uint64
func (obj *roCEv2IPv4Metric) ConnectRequestRx() uint64 {

	return *obj.obj.ConnectRequestRx

}

// Number of REQ Message Received.
// ConnectRequestRx returns a uint64
func (obj *roCEv2IPv4Metric) HasConnectRequestRx() bool {
	return obj.obj.ConnectRequestRx != nil
}

// Number of REQ Message Received.
// SetConnectRequestRx sets the uint64 value in the RoCEv2IPv4Metric object
func (obj *roCEv2IPv4Metric) SetConnectRequestRx(value uint64) RoCEv2IPv4Metric {

	obj.obj.ConnectRequestRx = &value
	return obj
}

// Number of REP Message Transmitted.
// ConnectReplyTx returns a uint64
func (obj *roCEv2IPv4Metric) ConnectReplyTx() uint64 {

	return *obj.obj.ConnectReplyTx

}

// Number of REP Message Transmitted.
// ConnectReplyTx returns a uint64
func (obj *roCEv2IPv4Metric) HasConnectReplyTx() bool {
	return obj.obj.ConnectReplyTx != nil
}

// Number of REP Message Transmitted.
// SetConnectReplyTx sets the uint64 value in the RoCEv2IPv4Metric object
func (obj *roCEv2IPv4Metric) SetConnectReplyTx(value uint64) RoCEv2IPv4Metric {

	obj.obj.ConnectReplyTx = &value
	return obj
}

// Number of REP Message Received.
// ConnectReplyRx returns a uint64
func (obj *roCEv2IPv4Metric) ConnectReplyRx() uint64 {

	return *obj.obj.ConnectReplyRx

}

// Number of REP Message Received.
// ConnectReplyRx returns a uint64
func (obj *roCEv2IPv4Metric) HasConnectReplyRx() bool {
	return obj.obj.ConnectReplyRx != nil
}

// Number of REP Message Received.
// SetConnectReplyRx sets the uint64 value in the RoCEv2IPv4Metric object
func (obj *roCEv2IPv4Metric) SetConnectReplyRx(value uint64) RoCEv2IPv4Metric {

	obj.obj.ConnectReplyRx = &value
	return obj
}

// Number of RTU Message Transmitted.
// ReadyTx returns a uint64
func (obj *roCEv2IPv4Metric) ReadyTx() uint64 {

	return *obj.obj.ReadyTx

}

// Number of RTU Message Transmitted.
// ReadyTx returns a uint64
func (obj *roCEv2IPv4Metric) HasReadyTx() bool {
	return obj.obj.ReadyTx != nil
}

// Number of RTU Message Transmitted.
// SetReadyTx sets the uint64 value in the RoCEv2IPv4Metric object
func (obj *roCEv2IPv4Metric) SetReadyTx(value uint64) RoCEv2IPv4Metric {

	obj.obj.ReadyTx = &value
	return obj
}

// Number of RTU Message Received.
// ReadyRx returns a uint64
func (obj *roCEv2IPv4Metric) ReadyRx() uint64 {

	return *obj.obj.ReadyRx

}

// Number of RTU Message Received.
// ReadyRx returns a uint64
func (obj *roCEv2IPv4Metric) HasReadyRx() bool {
	return obj.obj.ReadyRx != nil
}

// Number of RTU Message Received.
// SetReadyRx sets the uint64 value in the RoCEv2IPv4Metric object
func (obj *roCEv2IPv4Metric) SetReadyRx(value uint64) RoCEv2IPv4Metric {

	obj.obj.ReadyRx = &value
	return obj
}

// Number of DREQ Message Transmitted.
// DisconnectRequestTx returns a uint64
func (obj *roCEv2IPv4Metric) DisconnectRequestTx() uint64 {

	return *obj.obj.DisconnectRequestTx

}

// Number of DREQ Message Transmitted.
// DisconnectRequestTx returns a uint64
func (obj *roCEv2IPv4Metric) HasDisconnectRequestTx() bool {
	return obj.obj.DisconnectRequestTx != nil
}

// Number of DREQ Message Transmitted.
// SetDisconnectRequestTx sets the uint64 value in the RoCEv2IPv4Metric object
func (obj *roCEv2IPv4Metric) SetDisconnectRequestTx(value uint64) RoCEv2IPv4Metric {

	obj.obj.DisconnectRequestTx = &value
	return obj
}

// Number of DREQ Message Received.
// DisconnectRequestRx returns a uint64
func (obj *roCEv2IPv4Metric) DisconnectRequestRx() uint64 {

	return *obj.obj.DisconnectRequestRx

}

// Number of DREQ Message Received.
// DisconnectRequestRx returns a uint64
func (obj *roCEv2IPv4Metric) HasDisconnectRequestRx() bool {
	return obj.obj.DisconnectRequestRx != nil
}

// Number of DREQ Message Received.
// SetDisconnectRequestRx sets the uint64 value in the RoCEv2IPv4Metric object
func (obj *roCEv2IPv4Metric) SetDisconnectRequestRx(value uint64) RoCEv2IPv4Metric {

	obj.obj.DisconnectRequestRx = &value
	return obj
}

// Number of DREP Message Transmitted.
// DisconnectReplyTx returns a uint64
func (obj *roCEv2IPv4Metric) DisconnectReplyTx() uint64 {

	return *obj.obj.DisconnectReplyTx

}

// Number of DREP Message Transmitted.
// DisconnectReplyTx returns a uint64
func (obj *roCEv2IPv4Metric) HasDisconnectReplyTx() bool {
	return obj.obj.DisconnectReplyTx != nil
}

// Number of DREP Message Transmitted.
// SetDisconnectReplyTx sets the uint64 value in the RoCEv2IPv4Metric object
func (obj *roCEv2IPv4Metric) SetDisconnectReplyTx(value uint64) RoCEv2IPv4Metric {

	obj.obj.DisconnectReplyTx = &value
	return obj
}

// Number of DREP Message Received.
// DisconnectReplyRx returns a uint64
func (obj *roCEv2IPv4Metric) DisconnectReplyRx() uint64 {

	return *obj.obj.DisconnectReplyRx

}

// Number of DREP Message Received.
// DisconnectReplyRx returns a uint64
func (obj *roCEv2IPv4Metric) HasDisconnectReplyRx() bool {
	return obj.obj.DisconnectReplyRx != nil
}

// Number of DREP Message Received.
// SetDisconnectReplyRx sets the uint64 value in the RoCEv2IPv4Metric object
func (obj *roCEv2IPv4Metric) SetDisconnectReplyRx(value uint64) RoCEv2IPv4Metric {

	obj.obj.DisconnectReplyRx = &value
	return obj
}

// Number of REJ Message Transmitted.
// RejectTx returns a uint64
func (obj *roCEv2IPv4Metric) RejectTx() uint64 {

	return *obj.obj.RejectTx

}

// Number of REJ Message Transmitted.
// RejectTx returns a uint64
func (obj *roCEv2IPv4Metric) HasRejectTx() bool {
	return obj.obj.RejectTx != nil
}

// Number of REJ Message Transmitted.
// SetRejectTx sets the uint64 value in the RoCEv2IPv4Metric object
func (obj *roCEv2IPv4Metric) SetRejectTx(value uint64) RoCEv2IPv4Metric {

	obj.obj.RejectTx = &value
	return obj
}

// Number of Unknow Message Received.
// UnknownMsgRx returns a uint64
func (obj *roCEv2IPv4Metric) UnknownMsgRx() uint64 {

	return *obj.obj.UnknownMsgRx

}

// Number of Unknow Message Received.
// UnknownMsgRx returns a uint64
func (obj *roCEv2IPv4Metric) HasUnknownMsgRx() bool {
	return obj.obj.UnknownMsgRx != nil
}

// Number of Unknow Message Received.
// SetUnknownMsgRx sets the uint64 value in the RoCEv2IPv4Metric object
func (obj *roCEv2IPv4Metric) SetUnknownMsgRx(value uint64) RoCEv2IPv4Metric {

	obj.obj.UnknownMsgRx = &value
	return obj
}

func (obj *roCEv2IPv4Metric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *roCEv2IPv4Metric) setDefault() {

}
