package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecMetric *****
type macsecMetric struct {
	validation
	obj          *otg.MacsecMetric
	marshaller   marshalMacsecMetric
	unMarshaller unMarshalMacsecMetric
}

func NewMacsecMetric() MacsecMetric {
	obj := macsecMetric{obj: &otg.MacsecMetric{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecMetric) msg() *otg.MacsecMetric {
	return obj.obj
}

func (obj *macsecMetric) setMsg(msg *otg.MacsecMetric) MacsecMetric {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecMetric struct {
	obj *macsecMetric
}

type marshalMacsecMetric interface {
	// ToProto marshals MacsecMetric to protobuf object *otg.MacsecMetric
	ToProto() (*otg.MacsecMetric, error)
	// ToPbText marshals MacsecMetric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecMetric to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecMetric to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals MacsecMetric to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalmacsecMetric struct {
	obj *macsecMetric
}

type unMarshalMacsecMetric interface {
	// FromProto unmarshals MacsecMetric from protobuf object *otg.MacsecMetric
	FromProto(msg *otg.MacsecMetric) (MacsecMetric, error)
	// FromPbText unmarshals MacsecMetric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecMetric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecMetric from JSON text
	FromJson(value string) error
}

func (obj *macsecMetric) Marshal() marshalMacsecMetric {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecMetric{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecMetric) Unmarshal() unMarshalMacsecMetric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecMetric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecMetric) ToProto() (*otg.MacsecMetric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecMetric) FromProto(msg *otg.MacsecMetric) (MacsecMetric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecMetric) ToPbText() (string, error) {
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

func (m *unMarshalmacsecMetric) FromPbText(value string) error {
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

func (m *marshalmacsecMetric) ToYaml() (string, error) {
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

func (m *unMarshalmacsecMetric) FromYaml(value string) error {
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

func (m *marshalmacsecMetric) ToJsonRaw() (string, error) {
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

func (m *marshalmacsecMetric) ToJson() (string, error) {
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

func (m *unMarshalmacsecMetric) FromJson(value string) error {
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

func (obj *macsecMetric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecMetric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecMetric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecMetric) Clone() (MacsecMetric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecMetric()
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

// MacsecMetric is mACsec per secure entity(secY) statistics information.
type MacsecMetric interface {
	Validation
	// msg marshals MacsecMetric to protobuf object *otg.MacsecMetric
	// and doesn't set defaults
	msg() *otg.MacsecMetric
	// setMsg unmarshals MacsecMetric from protobuf object *otg.MacsecMetric
	// and doesn't set defaults
	setMsg(*otg.MacsecMetric) MacsecMetric
	// provides marshal interface
	Marshal() marshalMacsecMetric
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecMetric
	// validate validates MacsecMetric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecMetric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in MacsecMetric.
	Name() string
	// SetName assigns string provided by user to MacsecMetric
	SetName(value string) MacsecMetric
	// HasName checks if Name has been set in MacsecMetric
	HasName() bool
	// SessionState returns MacsecMetricSessionStateEnum, set in MacsecMetric
	SessionState() MacsecMetricSessionStateEnum
	// SetSessionState assigns MacsecMetricSessionStateEnum provided by user to MacsecMetric
	SetSessionState(value MacsecMetricSessionStateEnum) MacsecMetric
	// HasSessionState checks if SessionState has been set in MacsecMetric
	HasSessionState() bool
	// SessionFlapCount returns uint64, set in MacsecMetric.
	SessionFlapCount() uint64
	// SetSessionFlapCount assigns uint64 provided by user to MacsecMetric
	SetSessionFlapCount(value uint64) MacsecMetric
	// HasSessionFlapCount checks if SessionFlapCount has been set in MacsecMetric
	HasSessionFlapCount() bool
	// OutPktsProtected returns uint64, set in MacsecMetric.
	OutPktsProtected() uint64
	// SetOutPktsProtected assigns uint64 provided by user to MacsecMetric
	SetOutPktsProtected(value uint64) MacsecMetric
	// HasOutPktsProtected checks if OutPktsProtected has been set in MacsecMetric
	HasOutPktsProtected() bool
	// OutPktsEncrypted returns uint64, set in MacsecMetric.
	OutPktsEncrypted() uint64
	// SetOutPktsEncrypted assigns uint64 provided by user to MacsecMetric
	SetOutPktsEncrypted(value uint64) MacsecMetric
	// HasOutPktsEncrypted checks if OutPktsEncrypted has been set in MacsecMetric
	HasOutPktsEncrypted() bool
	// InPktsOk returns uint64, set in MacsecMetric.
	InPktsOk() uint64
	// SetInPktsOk assigns uint64 provided by user to MacsecMetric
	SetInPktsOk(value uint64) MacsecMetric
	// HasInPktsOk checks if InPktsOk has been set in MacsecMetric
	HasInPktsOk() bool
	// BadPktsRx returns uint64, set in MacsecMetric.
	BadPktsRx() uint64
	// SetBadPktsRx assigns uint64 provided by user to MacsecMetric
	SetBadPktsRx(value uint64) MacsecMetric
	// HasBadPktsRx checks if BadPktsRx has been set in MacsecMetric
	HasBadPktsRx() bool
	// InPktsBadTag returns uint64, set in MacsecMetric.
	InPktsBadTag() uint64
	// SetInPktsBadTag assigns uint64 provided by user to MacsecMetric
	SetInPktsBadTag(value uint64) MacsecMetric
	// HasInPktsBadTag checks if InPktsBadTag has been set in MacsecMetric
	HasInPktsBadTag() bool
	// InPktsLate returns uint64, set in MacsecMetric.
	InPktsLate() uint64
	// SetInPktsLate assigns uint64 provided by user to MacsecMetric
	SetInPktsLate(value uint64) MacsecMetric
	// HasInPktsLate checks if InPktsLate has been set in MacsecMetric
	HasInPktsLate() bool
	// InPktsNoSci returns uint64, set in MacsecMetric.
	InPktsNoSci() uint64
	// SetInPktsNoSci assigns uint64 provided by user to MacsecMetric
	SetInPktsNoSci(value uint64) MacsecMetric
	// HasInPktsNoSci checks if InPktsNoSci has been set in MacsecMetric
	HasInPktsNoSci() bool
	// InPktsNotUsingSa returns uint64, set in MacsecMetric.
	InPktsNotUsingSa() uint64
	// SetInPktsNotUsingSa assigns uint64 provided by user to MacsecMetric
	SetInPktsNotUsingSa(value uint64) MacsecMetric
	// HasInPktsNotUsingSa checks if InPktsNotUsingSa has been set in MacsecMetric
	HasInPktsNotUsingSa() bool
	// InPktsNotValid returns uint64, set in MacsecMetric.
	InPktsNotValid() uint64
	// SetInPktsNotValid assigns uint64 provided by user to MacsecMetric
	SetInPktsNotValid(value uint64) MacsecMetric
	// HasInPktsNotValid checks if InPktsNotValid has been set in MacsecMetric
	HasInPktsNotValid() bool
	// InPktsUnknownSci returns uint64, set in MacsecMetric.
	InPktsUnknownSci() uint64
	// SetInPktsUnknownSci assigns uint64 provided by user to MacsecMetric
	SetInPktsUnknownSci(value uint64) MacsecMetric
	// HasInPktsUnknownSci checks if InPktsUnknownSci has been set in MacsecMetric
	HasInPktsUnknownSci() bool
	// InPktsUnusedSa returns uint64, set in MacsecMetric.
	InPktsUnusedSa() uint64
	// SetInPktsUnusedSa assigns uint64 provided by user to MacsecMetric
	SetInPktsUnusedSa(value uint64) MacsecMetric
	// HasInPktsUnusedSa checks if InPktsUnusedSa has been set in MacsecMetric
	HasInPktsUnusedSa() bool
	// InPktsInvalid returns uint64, set in MacsecMetric.
	InPktsInvalid() uint64
	// SetInPktsInvalid assigns uint64 provided by user to MacsecMetric
	SetInPktsInvalid(value uint64) MacsecMetric
	// HasInPktsInvalid checks if InPktsInvalid has been set in MacsecMetric
	HasInPktsInvalid() bool
	// InPktsUntagged returns uint64, set in MacsecMetric.
	InPktsUntagged() uint64
	// SetInPktsUntagged assigns uint64 provided by user to MacsecMetric
	SetInPktsUntagged(value uint64) MacsecMetric
	// HasInPktsUntagged checks if InPktsUntagged has been set in MacsecMetric
	HasInPktsUntagged() bool
	// OutOctetsProtected returns uint64, set in MacsecMetric.
	OutOctetsProtected() uint64
	// SetOutOctetsProtected assigns uint64 provided by user to MacsecMetric
	SetOutOctetsProtected(value uint64) MacsecMetric
	// HasOutOctetsProtected checks if OutOctetsProtected has been set in MacsecMetric
	HasOutOctetsProtected() bool
	// OutOctetsEncrypted returns uint64, set in MacsecMetric.
	OutOctetsEncrypted() uint64
	// SetOutOctetsEncrypted assigns uint64 provided by user to MacsecMetric
	SetOutOctetsEncrypted(value uint64) MacsecMetric
	// HasOutOctetsEncrypted checks if OutOctetsEncrypted has been set in MacsecMetric
	HasOutOctetsEncrypted() bool
	// InOctetsValidated returns uint64, set in MacsecMetric.
	InOctetsValidated() uint64
	// SetInOctetsValidated assigns uint64 provided by user to MacsecMetric
	SetInOctetsValidated(value uint64) MacsecMetric
	// HasInOctetsValidated checks if InOctetsValidated has been set in MacsecMetric
	HasInOctetsValidated() bool
	// InOctetsDecrypted returns uint64, set in MacsecMetric.
	InOctetsDecrypted() uint64
	// SetInOctetsDecrypted assigns uint64 provided by user to MacsecMetric
	SetInOctetsDecrypted(value uint64) MacsecMetric
	// HasInOctetsDecrypted checks if InOctetsDecrypted has been set in MacsecMetric
	HasInOctetsDecrypted() bool
}

// The name of a configured MACsec secure entity(secY).
// Name returns a string
func (obj *macsecMetric) Name() string {

	return *obj.obj.Name

}

// The name of a configured MACsec secure entity(secY).
// Name returns a string
func (obj *macsecMetric) HasName() bool {
	return obj.obj.Name != nil
}

// The name of a configured MACsec secure entity(secY).
// SetName sets the string value in the MacsecMetric object
func (obj *macsecMetric) SetName(value string) MacsecMetric {

	obj.obj.Name = &value
	return obj
}

type MacsecMetricSessionStateEnum string

// Enum of SessionState on MacsecMetric
var MacsecMetricSessionState = struct {
	UP   MacsecMetricSessionStateEnum
	DOWN MacsecMetricSessionStateEnum
}{
	UP:   MacsecMetricSessionStateEnum("up"),
	DOWN: MacsecMetricSessionStateEnum("down"),
}

func (obj *macsecMetric) SessionState() MacsecMetricSessionStateEnum {
	return MacsecMetricSessionStateEnum(obj.obj.SessionState.Enum().String())
}

// Session state as up or down. Up refers to an Established state and Down refers to any other state.
// SessionState returns a string
func (obj *macsecMetric) HasSessionState() bool {
	return obj.obj.SessionState != nil
}

func (obj *macsecMetric) SetSessionState(value MacsecMetricSessionStateEnum) MacsecMetric {
	intValue, ok := otg.MacsecMetric_SessionState_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecMetricSessionStateEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecMetric_SessionState_Enum(intValue)
	obj.obj.SessionState = &enumValue

	return obj
}

// Number of times the session went from Up to Down state.
// SessionFlapCount returns a uint64
func (obj *macsecMetric) SessionFlapCount() uint64 {

	return *obj.obj.SessionFlapCount

}

// Number of times the session went from Up to Down state.
// SessionFlapCount returns a uint64
func (obj *macsecMetric) HasSessionFlapCount() bool {
	return obj.obj.SessionFlapCount != nil
}

// Number of times the session went from Up to Down state.
// SetSessionFlapCount sets the uint64 value in the MacsecMetric object
func (obj *macsecMetric) SetSessionFlapCount(value uint64) MacsecMetric {

	obj.obj.SessionFlapCount = &value
	return obj
}

// OutPktsProtected, the number of protected packets transmitted.
// OutPktsProtected returns a uint64
func (obj *macsecMetric) OutPktsProtected() uint64 {

	return *obj.obj.OutPktsProtected

}

// OutPktsProtected, the number of protected packets transmitted.
// OutPktsProtected returns a uint64
func (obj *macsecMetric) HasOutPktsProtected() bool {
	return obj.obj.OutPktsProtected != nil
}

// OutPktsProtected, the number of protected packets transmitted.
// SetOutPktsProtected sets the uint64 value in the MacsecMetric object
func (obj *macsecMetric) SetOutPktsProtected(value uint64) MacsecMetric {

	obj.obj.OutPktsProtected = &value
	return obj
}

// OutPktsEncrypted, the number of encrypted packets transmitted.
// OutPktsEncrypted returns a uint64
func (obj *macsecMetric) OutPktsEncrypted() uint64 {

	return *obj.obj.OutPktsEncrypted

}

// OutPktsEncrypted, the number of encrypted packets transmitted.
// OutPktsEncrypted returns a uint64
func (obj *macsecMetric) HasOutPktsEncrypted() bool {
	return obj.obj.OutPktsEncrypted != nil
}

// OutPktsEncrypted, the number of encrypted packets transmitted.
// SetOutPktsEncrypted sets the uint64 value in the MacsecMetric object
func (obj *macsecMetric) SetOutPktsEncrypted(value uint64) MacsecMetric {

	obj.obj.OutPktsEncrypted = &value
	return obj
}

// InPktsOk, the number of valid packets received.
// InPktsOk returns a uint64
func (obj *macsecMetric) InPktsOk() uint64 {

	return *obj.obj.InPktsOk

}

// InPktsOk, the number of valid packets received.
// InPktsOk returns a uint64
func (obj *macsecMetric) HasInPktsOk() bool {
	return obj.obj.InPktsOk != nil
}

// InPktsOk, the number of valid packets received.
// SetInPktsOk sets the uint64 value in the MacsecMetric object
func (obj *macsecMetric) SetInPktsOk(value uint64) MacsecMetric {

	obj.obj.InPktsOk = &value
	return obj
}

// The number of bad packets received.
// BadPktsRx returns a uint64
func (obj *macsecMetric) BadPktsRx() uint64 {

	return *obj.obj.BadPktsRx

}

// The number of bad packets received.
// BadPktsRx returns a uint64
func (obj *macsecMetric) HasBadPktsRx() bool {
	return obj.obj.BadPktsRx != nil
}

// The number of bad packets received.
// SetBadPktsRx sets the uint64 value in the MacsecMetric object
func (obj *macsecMetric) SetBadPktsRx(value uint64) MacsecMetric {

	obj.obj.BadPktsRx = &value
	return obj
}

// InPktsBadTag, the number of packets discarded due to bad tag/ICV.
// InPktsBadTag returns a uint64
func (obj *macsecMetric) InPktsBadTag() uint64 {

	return *obj.obj.InPktsBadTag

}

// InPktsBadTag, the number of packets discarded due to bad tag/ICV.
// InPktsBadTag returns a uint64
func (obj *macsecMetric) HasInPktsBadTag() bool {
	return obj.obj.InPktsBadTag != nil
}

// InPktsBadTag, the number of packets discarded due to bad tag/ICV.
// SetInPktsBadTag sets the uint64 value in the MacsecMetric object
func (obj *macsecMetric) SetInPktsBadTag(value uint64) MacsecMetric {

	obj.obj.InPktsBadTag = &value
	return obj
}

// InPktsLate, the number of packets discarded out of window.
// InPktsLate returns a uint64
func (obj *macsecMetric) InPktsLate() uint64 {

	return *obj.obj.InPktsLate

}

// InPktsLate, the number of packets discarded out of window.
// InPktsLate returns a uint64
func (obj *macsecMetric) HasInPktsLate() bool {
	return obj.obj.InPktsLate != nil
}

// InPktsLate, the number of packets discarded out of window.
// SetInPktsLate sets the uint64 value in the MacsecMetric object
func (obj *macsecMetric) SetInPktsLate(value uint64) MacsecMetric {

	obj.obj.InPktsLate = &value
	return obj
}

// InPktsNoSCI,the number of packets discarded due to unknown SCI.
// InPktsNoSci returns a uint64
func (obj *macsecMetric) InPktsNoSci() uint64 {

	return *obj.obj.InPktsNoSci

}

// InPktsNoSCI,the number of packets discarded due to unknown SCI.
// InPktsNoSci returns a uint64
func (obj *macsecMetric) HasInPktsNoSci() bool {
	return obj.obj.InPktsNoSci != nil
}

// InPktsNoSCI,the number of packets discarded due to unknown SCI.
// SetInPktsNoSci sets the uint64 value in the MacsecMetric object
func (obj *macsecMetric) SetInPktsNoSci(value uint64) MacsecMetric {

	obj.obj.InPktsNoSci = &value
	return obj
}

// InPktsNotUsingSA, the number of packets discarded due to unused SA.
// InPktsNotUsingSa returns a uint64
func (obj *macsecMetric) InPktsNotUsingSa() uint64 {

	return *obj.obj.InPktsNotUsingSa

}

// InPktsNotUsingSA, the number of packets discarded due to unused SA.
// InPktsNotUsingSa returns a uint64
func (obj *macsecMetric) HasInPktsNotUsingSa() bool {
	return obj.obj.InPktsNotUsingSa != nil
}

// InPktsNotUsingSA, the number of packets discarded due to unused SA.
// SetInPktsNotUsingSa sets the uint64 value in the MacsecMetric object
func (obj *macsecMetric) SetInPktsNotUsingSa(value uint64) MacsecMetric {

	obj.obj.InPktsNotUsingSa = &value
	return obj
}

// InPktsNotValid, the number of packets discarded due to invalid ICV.
// InPktsNotValid returns a uint64
func (obj *macsecMetric) InPktsNotValid() uint64 {

	return *obj.obj.InPktsNotValid

}

// InPktsNotValid, the number of packets discarded due to invalid ICV.
// InPktsNotValid returns a uint64
func (obj *macsecMetric) HasInPktsNotValid() bool {
	return obj.obj.InPktsNotValid != nil
}

// InPktsNotValid, the number of packets discarded due to invalid ICV.
// SetInPktsNotValid sets the uint64 value in the MacsecMetric object
func (obj *macsecMetric) SetInPktsNotValid(value uint64) MacsecMetric {

	obj.obj.InPktsNotValid = &value
	return obj
}

// InPktsUnknownSCI, the number of packets received with unknown SCI.
// InPktsUnknownSci returns a uint64
func (obj *macsecMetric) InPktsUnknownSci() uint64 {

	return *obj.obj.InPktsUnknownSci

}

// InPktsUnknownSCI, the number of packets received with unknown SCI.
// InPktsUnknownSci returns a uint64
func (obj *macsecMetric) HasInPktsUnknownSci() bool {
	return obj.obj.InPktsUnknownSci != nil
}

// InPktsUnknownSCI, the number of packets received with unknown SCI.
// SetInPktsUnknownSci sets the uint64 value in the MacsecMetric object
func (obj *macsecMetric) SetInPktsUnknownSci(value uint64) MacsecMetric {

	obj.obj.InPktsUnknownSci = &value
	return obj
}

// InPktsUnusedSA, the number of packets received with unused SA.
// InPktsUnusedSa returns a uint64
func (obj *macsecMetric) InPktsUnusedSa() uint64 {

	return *obj.obj.InPktsUnusedSa

}

// InPktsUnusedSA, the number of packets received with unused SA.
// InPktsUnusedSa returns a uint64
func (obj *macsecMetric) HasInPktsUnusedSa() bool {
	return obj.obj.InPktsUnusedSa != nil
}

// InPktsUnusedSA, the number of packets received with unused SA.
// SetInPktsUnusedSa sets the uint64 value in the MacsecMetric object
func (obj *macsecMetric) SetInPktsUnusedSa(value uint64) MacsecMetric {

	obj.obj.InPktsUnusedSa = &value
	return obj
}

// InPktsInvalid, the number of packets received with invalid ICV.
// InPktsInvalid returns a uint64
func (obj *macsecMetric) InPktsInvalid() uint64 {

	return *obj.obj.InPktsInvalid

}

// InPktsInvalid, the number of packets received with invalid ICV.
// InPktsInvalid returns a uint64
func (obj *macsecMetric) HasInPktsInvalid() bool {
	return obj.obj.InPktsInvalid != nil
}

// InPktsInvalid, the number of packets received with invalid ICV.
// SetInPktsInvalid sets the uint64 value in the MacsecMetric object
func (obj *macsecMetric) SetInPktsInvalid(value uint64) MacsecMetric {

	obj.obj.InPktsInvalid = &value
	return obj
}

// InPktsUntagged, the number of non-MACsec packets received.
// InPktsUntagged returns a uint64
func (obj *macsecMetric) InPktsUntagged() uint64 {

	return *obj.obj.InPktsUntagged

}

// InPktsUntagged, the number of non-MACsec packets received.
// InPktsUntagged returns a uint64
func (obj *macsecMetric) HasInPktsUntagged() bool {
	return obj.obj.InPktsUntagged != nil
}

// InPktsUntagged, the number of non-MACsec packets received.
// SetInPktsUntagged sets the uint64 value in the MacsecMetric object
func (obj *macsecMetric) SetInPktsUntagged(value uint64) MacsecMetric {

	obj.obj.InPktsUntagged = &value
	return obj
}

// OutOctetsProtected, the number of bytes transmitted as protected.
// OutOctetsProtected returns a uint64
func (obj *macsecMetric) OutOctetsProtected() uint64 {

	return *obj.obj.OutOctetsProtected

}

// OutOctetsProtected, the number of bytes transmitted as protected.
// OutOctetsProtected returns a uint64
func (obj *macsecMetric) HasOutOctetsProtected() bool {
	return obj.obj.OutOctetsProtected != nil
}

// OutOctetsProtected, the number of bytes transmitted as protected.
// SetOutOctetsProtected sets the uint64 value in the MacsecMetric object
func (obj *macsecMetric) SetOutOctetsProtected(value uint64) MacsecMetric {

	obj.obj.OutOctetsProtected = &value
	return obj
}

// OutOctetsEncrypted, the number of bytes transmitted as encrypted.
// OutOctetsEncrypted returns a uint64
func (obj *macsecMetric) OutOctetsEncrypted() uint64 {

	return *obj.obj.OutOctetsEncrypted

}

// OutOctetsEncrypted, the number of bytes transmitted as encrypted.
// OutOctetsEncrypted returns a uint64
func (obj *macsecMetric) HasOutOctetsEncrypted() bool {
	return obj.obj.OutOctetsEncrypted != nil
}

// OutOctetsEncrypted, the number of bytes transmitted as encrypted.
// SetOutOctetsEncrypted sets the uint64 value in the MacsecMetric object
func (obj *macsecMetric) SetOutOctetsEncrypted(value uint64) MacsecMetric {

	obj.obj.OutOctetsEncrypted = &value
	return obj
}

// InOctetsValidated, the number of received bytes validated.
// InOctetsValidated returns a uint64
func (obj *macsecMetric) InOctetsValidated() uint64 {

	return *obj.obj.InOctetsValidated

}

// InOctetsValidated, the number of received bytes validated.
// InOctetsValidated returns a uint64
func (obj *macsecMetric) HasInOctetsValidated() bool {
	return obj.obj.InOctetsValidated != nil
}

// InOctetsValidated, the number of received bytes validated.
// SetInOctetsValidated sets the uint64 value in the MacsecMetric object
func (obj *macsecMetric) SetInOctetsValidated(value uint64) MacsecMetric {

	obj.obj.InOctetsValidated = &value
	return obj
}

// InOctetsDecrypted, the number of received bytes decrypted.
// InOctetsDecrypted returns a uint64
func (obj *macsecMetric) InOctetsDecrypted() uint64 {

	return *obj.obj.InOctetsDecrypted

}

// InOctetsDecrypted, the number of received bytes decrypted.
// InOctetsDecrypted returns a uint64
func (obj *macsecMetric) HasInOctetsDecrypted() bool {
	return obj.obj.InOctetsDecrypted != nil
}

// InOctetsDecrypted, the number of received bytes decrypted.
// SetInOctetsDecrypted sets the uint64 value in the MacsecMetric object
func (obj *macsecMetric) SetInOctetsDecrypted(value uint64) MacsecMetric {

	obj.obj.InOctetsDecrypted = &value
	return obj
}

func (obj *macsecMetric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *macsecMetric) setDefault() {

}
