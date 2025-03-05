package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MkaMetric *****
type mkaMetric struct {
	validation
	obj          *otg.MkaMetric
	marshaller   marshalMkaMetric
	unMarshaller unMarshalMkaMetric
}

func NewMkaMetric() MkaMetric {
	obj := mkaMetric{obj: &otg.MkaMetric{}}
	obj.setDefault()
	return &obj
}

func (obj *mkaMetric) msg() *otg.MkaMetric {
	return obj.obj
}

func (obj *mkaMetric) setMsg(msg *otg.MkaMetric) MkaMetric {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmkaMetric struct {
	obj *mkaMetric
}

type marshalMkaMetric interface {
	// ToProto marshals MkaMetric to protobuf object *otg.MkaMetric
	ToProto() (*otg.MkaMetric, error)
	// ToPbText marshals MkaMetric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MkaMetric to YAML text
	ToYaml() (string, error)
	// ToJson marshals MkaMetric to JSON text
	ToJson() (string, error)
}

type unMarshalmkaMetric struct {
	obj *mkaMetric
}

type unMarshalMkaMetric interface {
	// FromProto unmarshals MkaMetric from protobuf object *otg.MkaMetric
	FromProto(msg *otg.MkaMetric) (MkaMetric, error)
	// FromPbText unmarshals MkaMetric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MkaMetric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MkaMetric from JSON text
	FromJson(value string) error
}

func (obj *mkaMetric) Marshal() marshalMkaMetric {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmkaMetric{obj: obj}
	}
	return obj.marshaller
}

func (obj *mkaMetric) Unmarshal() unMarshalMkaMetric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmkaMetric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmkaMetric) ToProto() (*otg.MkaMetric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmkaMetric) FromProto(msg *otg.MkaMetric) (MkaMetric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmkaMetric) ToPbText() (string, error) {
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

func (m *unMarshalmkaMetric) FromPbText(value string) error {
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

func (m *marshalmkaMetric) ToYaml() (string, error) {
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

func (m *unMarshalmkaMetric) FromYaml(value string) error {
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

func (m *marshalmkaMetric) ToJson() (string, error) {
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

func (m *unMarshalmkaMetric) FromJson(value string) error {
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

func (obj *mkaMetric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *mkaMetric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *mkaMetric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *mkaMetric) Clone() (MkaMetric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMkaMetric()
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

// MkaMetric is mKA per peer statistics information.
type MkaMetric interface {
	Validation
	// msg marshals MkaMetric to protobuf object *otg.MkaMetric
	// and doesn't set defaults
	msg() *otg.MkaMetric
	// setMsg unmarshals MkaMetric from protobuf object *otg.MkaMetric
	// and doesn't set defaults
	setMsg(*otg.MkaMetric) MkaMetric
	// provides marshal interface
	Marshal() marshalMkaMetric
	// provides unmarshal interface
	Unmarshal() unMarshalMkaMetric
	// validate validates MkaMetric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MkaMetric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in MkaMetric.
	Name() string
	// SetName assigns string provided by user to MkaMetric
	SetName(value string) MkaMetric
	// HasName checks if Name has been set in MkaMetric
	HasName() bool
	// SessionState returns MkaMetricSessionStateEnum, set in MkaMetric
	SessionState() MkaMetricSessionStateEnum
	// SetSessionState assigns MkaMetricSessionStateEnum provided by user to MkaMetric
	SetSessionState(value MkaMetricSessionStateEnum) MkaMetric
	// HasSessionState checks if SessionState has been set in MkaMetric
	HasSessionState() bool
	// SessionFlapCount returns uint64, set in MkaMetric.
	SessionFlapCount() uint64
	// SetSessionFlapCount assigns uint64 provided by user to MkaMetric
	SetSessionFlapCount(value uint64) MkaMetric
	// HasSessionFlapCount checks if SessionFlapCount has been set in MkaMetric
	HasSessionFlapCount() bool
	// MkpduTx returns uint64, set in MkaMetric.
	MkpduTx() uint64
	// SetMkpduTx assigns uint64 provided by user to MkaMetric
	SetMkpduTx(value uint64) MkaMetric
	// HasMkpduTx checks if MkpduTx has been set in MkaMetric
	HasMkpduTx() bool
	// MkpduRx returns uint64, set in MkaMetric.
	MkpduRx() uint64
	// SetMkpduRx assigns uint64 provided by user to MkaMetric
	SetMkpduRx(value uint64) MkaMetric
	// HasMkpduRx checks if MkpduRx has been set in MkaMetric
	HasMkpduRx() bool
	// LivePeerCount returns uint64, set in MkaMetric.
	LivePeerCount() uint64
	// SetLivePeerCount assigns uint64 provided by user to MkaMetric
	SetLivePeerCount(value uint64) MkaMetric
	// HasLivePeerCount checks if LivePeerCount has been set in MkaMetric
	HasLivePeerCount() bool
	// PotentialPeerCount returns uint64, set in MkaMetric.
	PotentialPeerCount() uint64
	// SetPotentialPeerCount assigns uint64 provided by user to MkaMetric
	SetPotentialPeerCount(value uint64) MkaMetric
	// HasPotentialPeerCount checks if PotentialPeerCount has been set in MkaMetric
	HasPotentialPeerCount() bool
	// LatestKeyTxPeerCount returns uint64, set in MkaMetric.
	LatestKeyTxPeerCount() uint64
	// SetLatestKeyTxPeerCount assigns uint64 provided by user to MkaMetric
	SetLatestKeyTxPeerCount(value uint64) MkaMetric
	// HasLatestKeyTxPeerCount checks if LatestKeyTxPeerCount has been set in MkaMetric
	HasLatestKeyTxPeerCount() bool
	// LatestKeyRxPeerCount returns uint64, set in MkaMetric.
	LatestKeyRxPeerCount() uint64
	// SetLatestKeyRxPeerCount assigns uint64 provided by user to MkaMetric
	SetLatestKeyRxPeerCount(value uint64) MkaMetric
	// HasLatestKeyRxPeerCount checks if LatestKeyRxPeerCount has been set in MkaMetric
	HasLatestKeyRxPeerCount() bool
	// MalformedMkpdu returns uint64, set in MkaMetric.
	MalformedMkpdu() uint64
	// SetMalformedMkpdu assigns uint64 provided by user to MkaMetric
	SetMalformedMkpdu(value uint64) MkaMetric
	// HasMalformedMkpdu checks if MalformedMkpdu has been set in MkaMetric
	HasMalformedMkpdu() bool
	// IcvMismatch returns uint64, set in MkaMetric.
	IcvMismatch() uint64
	// SetIcvMismatch assigns uint64 provided by user to MkaMetric
	SetIcvMismatch(value uint64) MkaMetric
	// HasIcvMismatch checks if IcvMismatch has been set in MkaMetric
	HasIcvMismatch() bool
}

// The name of a configured MKA peer.
// Name returns a string
func (obj *mkaMetric) Name() string {

	return *obj.obj.Name

}

// The name of a configured MKA peer.
// Name returns a string
func (obj *mkaMetric) HasName() bool {
	return obj.obj.Name != nil
}

// The name of a configured MKA peer.
// SetName sets the string value in the MkaMetric object
func (obj *mkaMetric) SetName(value string) MkaMetric {

	obj.obj.Name = &value
	return obj
}

type MkaMetricSessionStateEnum string

// Enum of SessionState on MkaMetric
var MkaMetricSessionState = struct {
	UP   MkaMetricSessionStateEnum
	DOWN MkaMetricSessionStateEnum
}{
	UP:   MkaMetricSessionStateEnum("up"),
	DOWN: MkaMetricSessionStateEnum("down"),
}

func (obj *mkaMetric) SessionState() MkaMetricSessionStateEnum {
	return MkaMetricSessionStateEnum(obj.obj.SessionState.Enum().String())
}

// Session state as up or down. Up refers to an Established state and Down refers to any other state.
// SessionState returns a string
func (obj *mkaMetric) HasSessionState() bool {
	return obj.obj.SessionState != nil
}

func (obj *mkaMetric) SetSessionState(value MkaMetricSessionStateEnum) MkaMetric {
	intValue, ok := otg.MkaMetric_SessionState_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MkaMetricSessionStateEnum", string(value)))
		return obj
	}
	enumValue := otg.MkaMetric_SessionState_Enum(intValue)
	obj.obj.SessionState = &enumValue

	return obj
}

// Number of times the session went from Up to Down state.
// SessionFlapCount returns a uint64
func (obj *mkaMetric) SessionFlapCount() uint64 {

	return *obj.obj.SessionFlapCount

}

// Number of times the session went from Up to Down state.
// SessionFlapCount returns a uint64
func (obj *mkaMetric) HasSessionFlapCount() bool {
	return obj.obj.SessionFlapCount != nil
}

// Number of times the session went from Up to Down state.
// SetSessionFlapCount sets the uint64 value in the MkaMetric object
func (obj *mkaMetric) SetSessionFlapCount(value uint64) MkaMetric {

	obj.obj.SessionFlapCount = &value
	return obj
}

// Number of MKA protocol data unit(MKPDU) frames Tx.
// MkpduTx returns a uint64
func (obj *mkaMetric) MkpduTx() uint64 {

	return *obj.obj.MkpduTx

}

// Number of MKA protocol data unit(MKPDU) frames Tx.
// MkpduTx returns a uint64
func (obj *mkaMetric) HasMkpduTx() bool {
	return obj.obj.MkpduTx != nil
}

// Number of MKA protocol data unit(MKPDU) frames Tx.
// SetMkpduTx sets the uint64 value in the MkaMetric object
func (obj *mkaMetric) SetMkpduTx(value uint64) MkaMetric {

	obj.obj.MkpduTx = &value
	return obj
}

// Number of MKA protocol data unit(MKPDU) frames Rx.
// MkpduRx returns a uint64
func (obj *mkaMetric) MkpduRx() uint64 {

	return *obj.obj.MkpduRx

}

// Number of MKA protocol data unit(MKPDU) frames Rx.
// MkpduRx returns a uint64
func (obj *mkaMetric) HasMkpduRx() bool {
	return obj.obj.MkpduRx != nil
}

// Number of MKA protocol data unit(MKPDU) frames Rx.
// SetMkpduRx sets the uint64 value in the MkaMetric object
func (obj *mkaMetric) SetMkpduRx(value uint64) MkaMetric {

	obj.obj.MkpduRx = &value
	return obj
}

// Number of MKA live peers.
// LivePeerCount returns a uint64
func (obj *mkaMetric) LivePeerCount() uint64 {

	return *obj.obj.LivePeerCount

}

// Number of MKA live peers.
// LivePeerCount returns a uint64
func (obj *mkaMetric) HasLivePeerCount() bool {
	return obj.obj.LivePeerCount != nil
}

// Number of MKA live peers.
// SetLivePeerCount sets the uint64 value in the MkaMetric object
func (obj *mkaMetric) SetLivePeerCount(value uint64) MkaMetric {

	obj.obj.LivePeerCount = &value
	return obj
}

// Number of MKA potential peers.
// PotentialPeerCount returns a uint64
func (obj *mkaMetric) PotentialPeerCount() uint64 {

	return *obj.obj.PotentialPeerCount

}

// Number of MKA potential peers.
// PotentialPeerCount returns a uint64
func (obj *mkaMetric) HasPotentialPeerCount() bool {
	return obj.obj.PotentialPeerCount != nil
}

// Number of MKA potential peers.
// SetPotentialPeerCount sets the uint64 value in the MkaMetric object
func (obj *mkaMetric) SetPotentialPeerCount(value uint64) MkaMetric {

	obj.obj.PotentialPeerCount = &value
	return obj
}

// Number of MKA latest key Tx peers.
// LatestKeyTxPeerCount returns a uint64
func (obj *mkaMetric) LatestKeyTxPeerCount() uint64 {

	return *obj.obj.LatestKeyTxPeerCount

}

// Number of MKA latest key Tx peers.
// LatestKeyTxPeerCount returns a uint64
func (obj *mkaMetric) HasLatestKeyTxPeerCount() bool {
	return obj.obj.LatestKeyTxPeerCount != nil
}

// Number of MKA latest key Tx peers.
// SetLatestKeyTxPeerCount sets the uint64 value in the MkaMetric object
func (obj *mkaMetric) SetLatestKeyTxPeerCount(value uint64) MkaMetric {

	obj.obj.LatestKeyTxPeerCount = &value
	return obj
}

// Number of MKA latest key Rx peers.
// LatestKeyRxPeerCount returns a uint64
func (obj *mkaMetric) LatestKeyRxPeerCount() uint64 {

	return *obj.obj.LatestKeyRxPeerCount

}

// Number of MKA latest key Rx peers.
// LatestKeyRxPeerCount returns a uint64
func (obj *mkaMetric) HasLatestKeyRxPeerCount() bool {
	return obj.obj.LatestKeyRxPeerCount != nil
}

// Number of MKA latest key Rx peers.
// SetLatestKeyRxPeerCount sets the uint64 value in the MkaMetric object
func (obj *mkaMetric) SetLatestKeyRxPeerCount(value uint64) MkaMetric {

	obj.obj.LatestKeyRxPeerCount = &value
	return obj
}

// Number of malformed MKA Protocol Data Unit(MKPDU) frames Rx.
// MalformedMkpdu returns a uint64
func (obj *mkaMetric) MalformedMkpdu() uint64 {

	return *obj.obj.MalformedMkpdu

}

// Number of malformed MKA Protocol Data Unit(MKPDU) frames Rx.
// MalformedMkpdu returns a uint64
func (obj *mkaMetric) HasMalformedMkpdu() bool {
	return obj.obj.MalformedMkpdu != nil
}

// Number of malformed MKA Protocol Data Unit(MKPDU) frames Rx.
// SetMalformedMkpdu sets the uint64 value in the MkaMetric object
func (obj *mkaMetric) SetMalformedMkpdu(value uint64) MkaMetric {

	obj.obj.MalformedMkpdu = &value
	return obj
}

// Number of MKA Protocol Data Unit(MKPDU) frames with ICV mismatch Rx.
// IcvMismatch returns a uint64
func (obj *mkaMetric) IcvMismatch() uint64 {

	return *obj.obj.IcvMismatch

}

// Number of MKA Protocol Data Unit(MKPDU) frames with ICV mismatch Rx.
// IcvMismatch returns a uint64
func (obj *mkaMetric) HasIcvMismatch() bool {
	return obj.obj.IcvMismatch != nil
}

// Number of MKA Protocol Data Unit(MKPDU) frames with ICV mismatch Rx.
// SetIcvMismatch sets the uint64 value in the MkaMetric object
func (obj *mkaMetric) SetIcvMismatch(value uint64) MkaMetric {

	obj.obj.IcvMismatch = &value
	return obj
}

func (obj *mkaMetric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *mkaMetric) setDefault() {

}
