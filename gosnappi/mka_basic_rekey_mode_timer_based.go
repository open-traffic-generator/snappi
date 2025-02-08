package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MkaBasicRekeyModeTimerBased *****
type mkaBasicRekeyModeTimerBased struct {
	validation
	obj          *otg.MkaBasicRekeyModeTimerBased
	marshaller   marshalMkaBasicRekeyModeTimerBased
	unMarshaller unMarshalMkaBasicRekeyModeTimerBased
}

func NewMkaBasicRekeyModeTimerBased() MkaBasicRekeyModeTimerBased {
	obj := mkaBasicRekeyModeTimerBased{obj: &otg.MkaBasicRekeyModeTimerBased{}}
	obj.setDefault()
	return &obj
}

func (obj *mkaBasicRekeyModeTimerBased) msg() *otg.MkaBasicRekeyModeTimerBased {
	return obj.obj
}

func (obj *mkaBasicRekeyModeTimerBased) setMsg(msg *otg.MkaBasicRekeyModeTimerBased) MkaBasicRekeyModeTimerBased {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmkaBasicRekeyModeTimerBased struct {
	obj *mkaBasicRekeyModeTimerBased
}

type marshalMkaBasicRekeyModeTimerBased interface {
	// ToProto marshals MkaBasicRekeyModeTimerBased to protobuf object *otg.MkaBasicRekeyModeTimerBased
	ToProto() (*otg.MkaBasicRekeyModeTimerBased, error)
	// ToPbText marshals MkaBasicRekeyModeTimerBased to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MkaBasicRekeyModeTimerBased to YAML text
	ToYaml() (string, error)
	// ToJson marshals MkaBasicRekeyModeTimerBased to JSON text
	ToJson() (string, error)
}

type unMarshalmkaBasicRekeyModeTimerBased struct {
	obj *mkaBasicRekeyModeTimerBased
}

type unMarshalMkaBasicRekeyModeTimerBased interface {
	// FromProto unmarshals MkaBasicRekeyModeTimerBased from protobuf object *otg.MkaBasicRekeyModeTimerBased
	FromProto(msg *otg.MkaBasicRekeyModeTimerBased) (MkaBasicRekeyModeTimerBased, error)
	// FromPbText unmarshals MkaBasicRekeyModeTimerBased from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MkaBasicRekeyModeTimerBased from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MkaBasicRekeyModeTimerBased from JSON text
	FromJson(value string) error
}

func (obj *mkaBasicRekeyModeTimerBased) Marshal() marshalMkaBasicRekeyModeTimerBased {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmkaBasicRekeyModeTimerBased{obj: obj}
	}
	return obj.marshaller
}

func (obj *mkaBasicRekeyModeTimerBased) Unmarshal() unMarshalMkaBasicRekeyModeTimerBased {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmkaBasicRekeyModeTimerBased{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmkaBasicRekeyModeTimerBased) ToProto() (*otg.MkaBasicRekeyModeTimerBased, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmkaBasicRekeyModeTimerBased) FromProto(msg *otg.MkaBasicRekeyModeTimerBased) (MkaBasicRekeyModeTimerBased, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmkaBasicRekeyModeTimerBased) ToPbText() (string, error) {
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

func (m *unMarshalmkaBasicRekeyModeTimerBased) FromPbText(value string) error {
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

func (m *marshalmkaBasicRekeyModeTimerBased) ToYaml() (string, error) {
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

func (m *unMarshalmkaBasicRekeyModeTimerBased) FromYaml(value string) error {
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

func (m *marshalmkaBasicRekeyModeTimerBased) ToJson() (string, error) {
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

func (m *unMarshalmkaBasicRekeyModeTimerBased) FromJson(value string) error {
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

func (obj *mkaBasicRekeyModeTimerBased) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *mkaBasicRekeyModeTimerBased) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *mkaBasicRekeyModeTimerBased) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *mkaBasicRekeyModeTimerBased) Clone() (MkaBasicRekeyModeTimerBased, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMkaBasicRekeyModeTimerBased()
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

// MkaBasicRekeyModeTimerBased is timer based periodic rekey properties.
type MkaBasicRekeyModeTimerBased interface {
	Validation
	// msg marshals MkaBasicRekeyModeTimerBased to protobuf object *otg.MkaBasicRekeyModeTimerBased
	// and doesn't set defaults
	msg() *otg.MkaBasicRekeyModeTimerBased
	// setMsg unmarshals MkaBasicRekeyModeTimerBased from protobuf object *otg.MkaBasicRekeyModeTimerBased
	// and doesn't set defaults
	setMsg(*otg.MkaBasicRekeyModeTimerBased) MkaBasicRekeyModeTimerBased
	// provides marshal interface
	Marshal() marshalMkaBasicRekeyModeTimerBased
	// provides unmarshal interface
	Unmarshal() unMarshalMkaBasicRekeyModeTimerBased
	// validate validates MkaBasicRekeyModeTimerBased
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MkaBasicRekeyModeTimerBased, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RekeyCount returns MkaBasicRekeyModeTimerBasedRekeyCountEnum, set in MkaBasicRekeyModeTimerBased
	RekeyCount() MkaBasicRekeyModeTimerBasedRekeyCountEnum
	// SetRekeyCount assigns MkaBasicRekeyModeTimerBasedRekeyCountEnum provided by user to MkaBasicRekeyModeTimerBased
	SetRekeyCount(value MkaBasicRekeyModeTimerBasedRekeyCountEnum) MkaBasicRekeyModeTimerBased
	// HasRekeyCount checks if RekeyCount has been set in MkaBasicRekeyModeTimerBased
	HasRekeyCount() bool
	// RekeyInterval returns uint32, set in MkaBasicRekeyModeTimerBased.
	RekeyInterval() uint32
	// SetRekeyInterval assigns uint32 provided by user to MkaBasicRekeyModeTimerBased
	SetRekeyInterval(value uint32) MkaBasicRekeyModeTimerBased
	// HasRekeyInterval checks if RekeyInterval has been set in MkaBasicRekeyModeTimerBased
	HasRekeyInterval() bool
}

type MkaBasicRekeyModeTimerBasedRekeyCountEnum string

// Enum of RekeyCount on MkaBasicRekeyModeTimerBased
var MkaBasicRekeyModeTimerBasedRekeyCount = struct {
	CONTINUOUS  MkaBasicRekeyModeTimerBasedRekeyCountEnum
	FIXED_COUNT MkaBasicRekeyModeTimerBasedRekeyCountEnum
}{
	CONTINUOUS:  MkaBasicRekeyModeTimerBasedRekeyCountEnum("continuous"),
	FIXED_COUNT: MkaBasicRekeyModeTimerBasedRekeyCountEnum("fixed_count"),
}

func (obj *mkaBasicRekeyModeTimerBased) RekeyCount() MkaBasicRekeyModeTimerBasedRekeyCountEnum {
	return MkaBasicRekeyModeTimerBasedRekeyCountEnum(obj.obj.RekeyCount.Enum().String())
}

// Periodic Rekey count.
// RekeyCount returns a string
func (obj *mkaBasicRekeyModeTimerBased) HasRekeyCount() bool {
	return obj.obj.RekeyCount != nil
}

func (obj *mkaBasicRekeyModeTimerBased) SetRekeyCount(value MkaBasicRekeyModeTimerBasedRekeyCountEnum) MkaBasicRekeyModeTimerBased {
	intValue, ok := otg.MkaBasicRekeyModeTimerBased_RekeyCount_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MkaBasicRekeyModeTimerBasedRekeyCountEnum", string(value)))
		return obj
	}
	enumValue := otg.MkaBasicRekeyModeTimerBased_RekeyCount_Enum(intValue)
	obj.obj.RekeyCount = &enumValue

	return obj
}

// Periodic Rekey Interval (sec).
// RekeyInterval returns a uint32
func (obj *mkaBasicRekeyModeTimerBased) RekeyInterval() uint32 {

	return *obj.obj.RekeyInterval

}

// Periodic Rekey Interval (sec).
// RekeyInterval returns a uint32
func (obj *mkaBasicRekeyModeTimerBased) HasRekeyInterval() bool {
	return obj.obj.RekeyInterval != nil
}

// Periodic Rekey Interval (sec).
// SetRekeyInterval sets the uint32 value in the MkaBasicRekeyModeTimerBased object
func (obj *mkaBasicRekeyModeTimerBased) SetRekeyInterval(value uint32) MkaBasicRekeyModeTimerBased {

	obj.obj.RekeyInterval = &value
	return obj
}

func (obj *mkaBasicRekeyModeTimerBased) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RekeyInterval != nil {

		if *obj.obj.RekeyInterval < 30 || *obj.obj.RekeyInterval > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("30 <= MkaBasicRekeyModeTimerBased.RekeyInterval <= 65535 but Got %d", *obj.obj.RekeyInterval))
		}

	}

}

func (obj *mkaBasicRekeyModeTimerBased) setDefault() {
	if obj.obj.RekeyCount == nil {
		obj.SetRekeyCount(MkaBasicRekeyModeTimerBasedRekeyCount.CONTINUOUS)

	}
	if obj.obj.RekeyInterval == nil {
		obj.SetRekeyInterval(300)
	}

}
