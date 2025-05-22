package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2IPv4MetricPerPeer *****
type rocev2IPv4MetricPerPeer struct {
	validation
	obj          *otg.Rocev2IPv4MetricPerPeer
	marshaller   marshalRocev2IPv4MetricPerPeer
	unMarshaller unMarshalRocev2IPv4MetricPerPeer
}

func NewRocev2IPv4MetricPerPeer() Rocev2IPv4MetricPerPeer {
	obj := rocev2IPv4MetricPerPeer{obj: &otg.Rocev2IPv4MetricPerPeer{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2IPv4MetricPerPeer) msg() *otg.Rocev2IPv4MetricPerPeer {
	return obj.obj
}

func (obj *rocev2IPv4MetricPerPeer) setMsg(msg *otg.Rocev2IPv4MetricPerPeer) Rocev2IPv4MetricPerPeer {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2IPv4MetricPerPeer struct {
	obj *rocev2IPv4MetricPerPeer
}

type marshalRocev2IPv4MetricPerPeer interface {
	// ToProto marshals Rocev2IPv4MetricPerPeer to protobuf object *otg.Rocev2IPv4MetricPerPeer
	ToProto() (*otg.Rocev2IPv4MetricPerPeer, error)
	// ToPbText marshals Rocev2IPv4MetricPerPeer to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2IPv4MetricPerPeer to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2IPv4MetricPerPeer to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Rocev2IPv4MetricPerPeer to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalrocev2IPv4MetricPerPeer struct {
	obj *rocev2IPv4MetricPerPeer
}

type unMarshalRocev2IPv4MetricPerPeer interface {
	// FromProto unmarshals Rocev2IPv4MetricPerPeer from protobuf object *otg.Rocev2IPv4MetricPerPeer
	FromProto(msg *otg.Rocev2IPv4MetricPerPeer) (Rocev2IPv4MetricPerPeer, error)
	// FromPbText unmarshals Rocev2IPv4MetricPerPeer from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2IPv4MetricPerPeer from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2IPv4MetricPerPeer from JSON text
	FromJson(value string) error
}

func (obj *rocev2IPv4MetricPerPeer) Marshal() marshalRocev2IPv4MetricPerPeer {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2IPv4MetricPerPeer{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2IPv4MetricPerPeer) Unmarshal() unMarshalRocev2IPv4MetricPerPeer {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2IPv4MetricPerPeer{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2IPv4MetricPerPeer) ToProto() (*otg.Rocev2IPv4MetricPerPeer, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2IPv4MetricPerPeer) FromProto(msg *otg.Rocev2IPv4MetricPerPeer) (Rocev2IPv4MetricPerPeer, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2IPv4MetricPerPeer) ToPbText() (string, error) {
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

func (m *unMarshalrocev2IPv4MetricPerPeer) FromPbText(value string) error {
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

func (m *marshalrocev2IPv4MetricPerPeer) ToYaml() (string, error) {
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

func (m *unMarshalrocev2IPv4MetricPerPeer) FromYaml(value string) error {
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

func (m *marshalrocev2IPv4MetricPerPeer) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalrocev2IPv4MetricPerPeer) ToJson() (string, error) {
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

func (m *unMarshalrocev2IPv4MetricPerPeer) FromJson(value string) error {
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

func (obj *rocev2IPv4MetricPerPeer) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2IPv4MetricPerPeer) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2IPv4MetricPerPeer) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2IPv4MetricPerPeer) Clone() (Rocev2IPv4MetricPerPeer, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2IPv4MetricPerPeer()
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

// Rocev2IPv4MetricPerPeer is roCEv2 per peer statistics information.
type Rocev2IPv4MetricPerPeer interface {
	Validation
	// msg marshals Rocev2IPv4MetricPerPeer to protobuf object *otg.Rocev2IPv4MetricPerPeer
	// and doesn't set defaults
	msg() *otg.Rocev2IPv4MetricPerPeer
	// setMsg unmarshals Rocev2IPv4MetricPerPeer from protobuf object *otg.Rocev2IPv4MetricPerPeer
	// and doesn't set defaults
	setMsg(*otg.Rocev2IPv4MetricPerPeer) Rocev2IPv4MetricPerPeer
	// provides marshal interface
	Marshal() marshalRocev2IPv4MetricPerPeer
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2IPv4MetricPerPeer
	// validate validates Rocev2IPv4MetricPerPeer
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2IPv4MetricPerPeer, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in Rocev2IPv4MetricPerPeer.
	Name() string
	// SetName assigns string provided by user to Rocev2IPv4MetricPerPeer
	SetName(value string) Rocev2IPv4MetricPerPeer
	// HasName checks if Name has been set in Rocev2IPv4MetricPerPeer
	HasName() bool
	// QpConfigured returns uint64, set in Rocev2IPv4MetricPerPeer.
	QpConfigured() uint64
	// SetQpConfigured assigns uint64 provided by user to Rocev2IPv4MetricPerPeer
	SetQpConfigured(value uint64) Rocev2IPv4MetricPerPeer
	// HasQpConfigured checks if QpConfigured has been set in Rocev2IPv4MetricPerPeer
	HasQpConfigured() bool
	// QpUp returns uint64, set in Rocev2IPv4MetricPerPeer.
	QpUp() uint64
	// SetQpUp assigns uint64 provided by user to Rocev2IPv4MetricPerPeer
	SetQpUp(value uint64) Rocev2IPv4MetricPerPeer
	// HasQpUp checks if QpUp has been set in Rocev2IPv4MetricPerPeer
	HasQpUp() bool
	// QpDown returns uint64, set in Rocev2IPv4MetricPerPeer.
	QpDown() uint64
	// SetQpDown assigns uint64 provided by user to Rocev2IPv4MetricPerPeer
	SetQpDown(value uint64) Rocev2IPv4MetricPerPeer
	// HasQpDown checks if QpDown has been set in Rocev2IPv4MetricPerPeer
	HasQpDown() bool
	// ConnectRequestTx returns uint64, set in Rocev2IPv4MetricPerPeer.
	ConnectRequestTx() uint64
	// SetConnectRequestTx assigns uint64 provided by user to Rocev2IPv4MetricPerPeer
	SetConnectRequestTx(value uint64) Rocev2IPv4MetricPerPeer
	// HasConnectRequestTx checks if ConnectRequestTx has been set in Rocev2IPv4MetricPerPeer
	HasConnectRequestTx() bool
	// ConnectRequestRx returns uint64, set in Rocev2IPv4MetricPerPeer.
	ConnectRequestRx() uint64
	// SetConnectRequestRx assigns uint64 provided by user to Rocev2IPv4MetricPerPeer
	SetConnectRequestRx(value uint64) Rocev2IPv4MetricPerPeer
	// HasConnectRequestRx checks if ConnectRequestRx has been set in Rocev2IPv4MetricPerPeer
	HasConnectRequestRx() bool
	// ConnectReplyTx returns uint64, set in Rocev2IPv4MetricPerPeer.
	ConnectReplyTx() uint64
	// SetConnectReplyTx assigns uint64 provided by user to Rocev2IPv4MetricPerPeer
	SetConnectReplyTx(value uint64) Rocev2IPv4MetricPerPeer
	// HasConnectReplyTx checks if ConnectReplyTx has been set in Rocev2IPv4MetricPerPeer
	HasConnectReplyTx() bool
	// ConnectReplyRx returns uint64, set in Rocev2IPv4MetricPerPeer.
	ConnectReplyRx() uint64
	// SetConnectReplyRx assigns uint64 provided by user to Rocev2IPv4MetricPerPeer
	SetConnectReplyRx(value uint64) Rocev2IPv4MetricPerPeer
	// HasConnectReplyRx checks if ConnectReplyRx has been set in Rocev2IPv4MetricPerPeer
	HasConnectReplyRx() bool
	// ReadyTx returns uint64, set in Rocev2IPv4MetricPerPeer.
	ReadyTx() uint64
	// SetReadyTx assigns uint64 provided by user to Rocev2IPv4MetricPerPeer
	SetReadyTx(value uint64) Rocev2IPv4MetricPerPeer
	// HasReadyTx checks if ReadyTx has been set in Rocev2IPv4MetricPerPeer
	HasReadyTx() bool
	// ReadyRx returns uint64, set in Rocev2IPv4MetricPerPeer.
	ReadyRx() uint64
	// SetReadyRx assigns uint64 provided by user to Rocev2IPv4MetricPerPeer
	SetReadyRx(value uint64) Rocev2IPv4MetricPerPeer
	// HasReadyRx checks if ReadyRx has been set in Rocev2IPv4MetricPerPeer
	HasReadyRx() bool
	// DisconnectRequestTx returns uint64, set in Rocev2IPv4MetricPerPeer.
	DisconnectRequestTx() uint64
	// SetDisconnectRequestTx assigns uint64 provided by user to Rocev2IPv4MetricPerPeer
	SetDisconnectRequestTx(value uint64) Rocev2IPv4MetricPerPeer
	// HasDisconnectRequestTx checks if DisconnectRequestTx has been set in Rocev2IPv4MetricPerPeer
	HasDisconnectRequestTx() bool
	// DisconnectRequestRx returns uint64, set in Rocev2IPv4MetricPerPeer.
	DisconnectRequestRx() uint64
	// SetDisconnectRequestRx assigns uint64 provided by user to Rocev2IPv4MetricPerPeer
	SetDisconnectRequestRx(value uint64) Rocev2IPv4MetricPerPeer
	// HasDisconnectRequestRx checks if DisconnectRequestRx has been set in Rocev2IPv4MetricPerPeer
	HasDisconnectRequestRx() bool
	// DisconnectReplyTx returns uint64, set in Rocev2IPv4MetricPerPeer.
	DisconnectReplyTx() uint64
	// SetDisconnectReplyTx assigns uint64 provided by user to Rocev2IPv4MetricPerPeer
	SetDisconnectReplyTx(value uint64) Rocev2IPv4MetricPerPeer
	// HasDisconnectReplyTx checks if DisconnectReplyTx has been set in Rocev2IPv4MetricPerPeer
	HasDisconnectReplyTx() bool
	// DisconnectReplyRx returns uint64, set in Rocev2IPv4MetricPerPeer.
	DisconnectReplyRx() uint64
	// SetDisconnectReplyRx assigns uint64 provided by user to Rocev2IPv4MetricPerPeer
	SetDisconnectReplyRx(value uint64) Rocev2IPv4MetricPerPeer
	// HasDisconnectReplyRx checks if DisconnectReplyRx has been set in Rocev2IPv4MetricPerPeer
	HasDisconnectReplyRx() bool
	// RejectTx returns uint64, set in Rocev2IPv4MetricPerPeer.
	RejectTx() uint64
	// SetRejectTx assigns uint64 provided by user to Rocev2IPv4MetricPerPeer
	SetRejectTx(value uint64) Rocev2IPv4MetricPerPeer
	// HasRejectTx checks if RejectTx has been set in Rocev2IPv4MetricPerPeer
	HasRejectTx() bool
	// UnknownMsgRx returns uint64, set in Rocev2IPv4MetricPerPeer.
	UnknownMsgRx() uint64
	// SetUnknownMsgRx assigns uint64 provided by user to Rocev2IPv4MetricPerPeer
	SetUnknownMsgRx(value uint64) Rocev2IPv4MetricPerPeer
	// HasUnknownMsgRx checks if UnknownMsgRx has been set in Rocev2IPv4MetricPerPeer
	HasUnknownMsgRx() bool
}

// The name of a configured RoCEv2 peer.
// Name returns a string
func (obj *rocev2IPv4MetricPerPeer) Name() string {

	return *obj.obj.Name

}

// The name of a configured RoCEv2 peer.
// Name returns a string
func (obj *rocev2IPv4MetricPerPeer) HasName() bool {
	return obj.obj.Name != nil
}

// The name of a configured RoCEv2 peer.
// SetName sets the string value in the Rocev2IPv4MetricPerPeer object
func (obj *rocev2IPv4MetricPerPeer) SetName(value string) Rocev2IPv4MetricPerPeer {

	obj.obj.Name = &value
	return obj
}

// Number of QPs configured on this port.
// QpConfigured returns a uint64
func (obj *rocev2IPv4MetricPerPeer) QpConfigured() uint64 {

	return *obj.obj.QpConfigured

}

// Number of QPs configured on this port.
// QpConfigured returns a uint64
func (obj *rocev2IPv4MetricPerPeer) HasQpConfigured() bool {
	return obj.obj.QpConfigured != nil
}

// Number of QPs configured on this port.
// SetQpConfigured sets the uint64 value in the Rocev2IPv4MetricPerPeer object
func (obj *rocev2IPv4MetricPerPeer) SetQpConfigured(value uint64) Rocev2IPv4MetricPerPeer {

	obj.obj.QpConfigured = &value
	return obj
}

// Number of QPs that are in UP state.
// QpUp returns a uint64
func (obj *rocev2IPv4MetricPerPeer) QpUp() uint64 {

	return *obj.obj.QpUp

}

// Number of QPs that are in UP state.
// QpUp returns a uint64
func (obj *rocev2IPv4MetricPerPeer) HasQpUp() bool {
	return obj.obj.QpUp != nil
}

// Number of QPs that are in UP state.
// SetQpUp sets the uint64 value in the Rocev2IPv4MetricPerPeer object
func (obj *rocev2IPv4MetricPerPeer) SetQpUp(value uint64) Rocev2IPv4MetricPerPeer {

	obj.obj.QpUp = &value
	return obj
}

// Number of QPs that have not come UP.
// QpDown returns a uint64
func (obj *rocev2IPv4MetricPerPeer) QpDown() uint64 {

	return *obj.obj.QpDown

}

// Number of QPs that have not come UP.
// QpDown returns a uint64
func (obj *rocev2IPv4MetricPerPeer) HasQpDown() bool {
	return obj.obj.QpDown != nil
}

// Number of QPs that have not come UP.
// SetQpDown sets the uint64 value in the Rocev2IPv4MetricPerPeer object
func (obj *rocev2IPv4MetricPerPeer) SetQpDown(value uint64) Rocev2IPv4MetricPerPeer {

	obj.obj.QpDown = &value
	return obj
}

// Number of REQ Message Transmitted.
// ConnectRequestTx returns a uint64
func (obj *rocev2IPv4MetricPerPeer) ConnectRequestTx() uint64 {

	return *obj.obj.ConnectRequestTx

}

// Number of REQ Message Transmitted.
// ConnectRequestTx returns a uint64
func (obj *rocev2IPv4MetricPerPeer) HasConnectRequestTx() bool {
	return obj.obj.ConnectRequestTx != nil
}

// Number of REQ Message Transmitted.
// SetConnectRequestTx sets the uint64 value in the Rocev2IPv4MetricPerPeer object
func (obj *rocev2IPv4MetricPerPeer) SetConnectRequestTx(value uint64) Rocev2IPv4MetricPerPeer {

	obj.obj.ConnectRequestTx = &value
	return obj
}

// Number of REQ Message Received.
// ConnectRequestRx returns a uint64
func (obj *rocev2IPv4MetricPerPeer) ConnectRequestRx() uint64 {

	return *obj.obj.ConnectRequestRx

}

// Number of REQ Message Received.
// ConnectRequestRx returns a uint64
func (obj *rocev2IPv4MetricPerPeer) HasConnectRequestRx() bool {
	return obj.obj.ConnectRequestRx != nil
}

// Number of REQ Message Received.
// SetConnectRequestRx sets the uint64 value in the Rocev2IPv4MetricPerPeer object
func (obj *rocev2IPv4MetricPerPeer) SetConnectRequestRx(value uint64) Rocev2IPv4MetricPerPeer {

	obj.obj.ConnectRequestRx = &value
	return obj
}

// Number of REP Message Transmitted.
// ConnectReplyTx returns a uint64
func (obj *rocev2IPv4MetricPerPeer) ConnectReplyTx() uint64 {

	return *obj.obj.ConnectReplyTx

}

// Number of REP Message Transmitted.
// ConnectReplyTx returns a uint64
func (obj *rocev2IPv4MetricPerPeer) HasConnectReplyTx() bool {
	return obj.obj.ConnectReplyTx != nil
}

// Number of REP Message Transmitted.
// SetConnectReplyTx sets the uint64 value in the Rocev2IPv4MetricPerPeer object
func (obj *rocev2IPv4MetricPerPeer) SetConnectReplyTx(value uint64) Rocev2IPv4MetricPerPeer {

	obj.obj.ConnectReplyTx = &value
	return obj
}

// Number of REP Message Received.
// ConnectReplyRx returns a uint64
func (obj *rocev2IPv4MetricPerPeer) ConnectReplyRx() uint64 {

	return *obj.obj.ConnectReplyRx

}

// Number of REP Message Received.
// ConnectReplyRx returns a uint64
func (obj *rocev2IPv4MetricPerPeer) HasConnectReplyRx() bool {
	return obj.obj.ConnectReplyRx != nil
}

// Number of REP Message Received.
// SetConnectReplyRx sets the uint64 value in the Rocev2IPv4MetricPerPeer object
func (obj *rocev2IPv4MetricPerPeer) SetConnectReplyRx(value uint64) Rocev2IPv4MetricPerPeer {

	obj.obj.ConnectReplyRx = &value
	return obj
}

// Number of RTU Message Transmitted.
// ReadyTx returns a uint64
func (obj *rocev2IPv4MetricPerPeer) ReadyTx() uint64 {

	return *obj.obj.ReadyTx

}

// Number of RTU Message Transmitted.
// ReadyTx returns a uint64
func (obj *rocev2IPv4MetricPerPeer) HasReadyTx() bool {
	return obj.obj.ReadyTx != nil
}

// Number of RTU Message Transmitted.
// SetReadyTx sets the uint64 value in the Rocev2IPv4MetricPerPeer object
func (obj *rocev2IPv4MetricPerPeer) SetReadyTx(value uint64) Rocev2IPv4MetricPerPeer {

	obj.obj.ReadyTx = &value
	return obj
}

// Number of RTU Message Received.
// ReadyRx returns a uint64
func (obj *rocev2IPv4MetricPerPeer) ReadyRx() uint64 {

	return *obj.obj.ReadyRx

}

// Number of RTU Message Received.
// ReadyRx returns a uint64
func (obj *rocev2IPv4MetricPerPeer) HasReadyRx() bool {
	return obj.obj.ReadyRx != nil
}

// Number of RTU Message Received.
// SetReadyRx sets the uint64 value in the Rocev2IPv4MetricPerPeer object
func (obj *rocev2IPv4MetricPerPeer) SetReadyRx(value uint64) Rocev2IPv4MetricPerPeer {

	obj.obj.ReadyRx = &value
	return obj
}

// Number of DREQ Message Transmitted.
// DisconnectRequestTx returns a uint64
func (obj *rocev2IPv4MetricPerPeer) DisconnectRequestTx() uint64 {

	return *obj.obj.DisconnectRequestTx

}

// Number of DREQ Message Transmitted.
// DisconnectRequestTx returns a uint64
func (obj *rocev2IPv4MetricPerPeer) HasDisconnectRequestTx() bool {
	return obj.obj.DisconnectRequestTx != nil
}

// Number of DREQ Message Transmitted.
// SetDisconnectRequestTx sets the uint64 value in the Rocev2IPv4MetricPerPeer object
func (obj *rocev2IPv4MetricPerPeer) SetDisconnectRequestTx(value uint64) Rocev2IPv4MetricPerPeer {

	obj.obj.DisconnectRequestTx = &value
	return obj
}

// Number of DREQ Message Received.
// DisconnectRequestRx returns a uint64
func (obj *rocev2IPv4MetricPerPeer) DisconnectRequestRx() uint64 {

	return *obj.obj.DisconnectRequestRx

}

// Number of DREQ Message Received.
// DisconnectRequestRx returns a uint64
func (obj *rocev2IPv4MetricPerPeer) HasDisconnectRequestRx() bool {
	return obj.obj.DisconnectRequestRx != nil
}

// Number of DREQ Message Received.
// SetDisconnectRequestRx sets the uint64 value in the Rocev2IPv4MetricPerPeer object
func (obj *rocev2IPv4MetricPerPeer) SetDisconnectRequestRx(value uint64) Rocev2IPv4MetricPerPeer {

	obj.obj.DisconnectRequestRx = &value
	return obj
}

// Number of DREP Message Transmitted.
// DisconnectReplyTx returns a uint64
func (obj *rocev2IPv4MetricPerPeer) DisconnectReplyTx() uint64 {

	return *obj.obj.DisconnectReplyTx

}

// Number of DREP Message Transmitted.
// DisconnectReplyTx returns a uint64
func (obj *rocev2IPv4MetricPerPeer) HasDisconnectReplyTx() bool {
	return obj.obj.DisconnectReplyTx != nil
}

// Number of DREP Message Transmitted.
// SetDisconnectReplyTx sets the uint64 value in the Rocev2IPv4MetricPerPeer object
func (obj *rocev2IPv4MetricPerPeer) SetDisconnectReplyTx(value uint64) Rocev2IPv4MetricPerPeer {

	obj.obj.DisconnectReplyTx = &value
	return obj
}

// Number of DREP Message Received.
// DisconnectReplyRx returns a uint64
func (obj *rocev2IPv4MetricPerPeer) DisconnectReplyRx() uint64 {

	return *obj.obj.DisconnectReplyRx

}

// Number of DREP Message Received.
// DisconnectReplyRx returns a uint64
func (obj *rocev2IPv4MetricPerPeer) HasDisconnectReplyRx() bool {
	return obj.obj.DisconnectReplyRx != nil
}

// Number of DREP Message Received.
// SetDisconnectReplyRx sets the uint64 value in the Rocev2IPv4MetricPerPeer object
func (obj *rocev2IPv4MetricPerPeer) SetDisconnectReplyRx(value uint64) Rocev2IPv4MetricPerPeer {

	obj.obj.DisconnectReplyRx = &value
	return obj
}

// Number of REJ Message Transmitted.
// RejectTx returns a uint64
func (obj *rocev2IPv4MetricPerPeer) RejectTx() uint64 {

	return *obj.obj.RejectTx

}

// Number of REJ Message Transmitted.
// RejectTx returns a uint64
func (obj *rocev2IPv4MetricPerPeer) HasRejectTx() bool {
	return obj.obj.RejectTx != nil
}

// Number of REJ Message Transmitted.
// SetRejectTx sets the uint64 value in the Rocev2IPv4MetricPerPeer object
func (obj *rocev2IPv4MetricPerPeer) SetRejectTx(value uint64) Rocev2IPv4MetricPerPeer {

	obj.obj.RejectTx = &value
	return obj
}

// Number of Unknown Message Received.
// UnknownMsgRx returns a uint64
func (obj *rocev2IPv4MetricPerPeer) UnknownMsgRx() uint64 {

	return *obj.obj.UnknownMsgRx

}

// Number of Unknown Message Received.
// UnknownMsgRx returns a uint64
func (obj *rocev2IPv4MetricPerPeer) HasUnknownMsgRx() bool {
	return obj.obj.UnknownMsgRx != nil
}

// Number of Unknown Message Received.
// SetUnknownMsgRx sets the uint64 value in the Rocev2IPv4MetricPerPeer object
func (obj *rocev2IPv4MetricPerPeer) SetUnknownMsgRx(value uint64) Rocev2IPv4MetricPerPeer {

	obj.obj.UnknownMsgRx = &value
	return obj
}

func (obj *rocev2IPv4MetricPerPeer) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *rocev2IPv4MetricPerPeer) setDefault() {

}
