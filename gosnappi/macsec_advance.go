package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecAdvance *****
type macsecAdvance struct {
	validation
	obj          *otg.MacsecAdvance
	marshaller   marshalMacsecAdvance
	unMarshaller unMarshalMacsecAdvance
}

func NewMacsecAdvance() MacsecAdvance {
	obj := macsecAdvance{obj: &otg.MacsecAdvance{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecAdvance) msg() *otg.MacsecAdvance {
	return obj.obj
}

func (obj *macsecAdvance) setMsg(msg *otg.MacsecAdvance) MacsecAdvance {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecAdvance struct {
	obj *macsecAdvance
}

type marshalMacsecAdvance interface {
	// ToProto marshals MacsecAdvance to protobuf object *otg.MacsecAdvance
	ToProto() (*otg.MacsecAdvance, error)
	// ToPbText marshals MacsecAdvance to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecAdvance to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecAdvance to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecAdvance struct {
	obj *macsecAdvance
}

type unMarshalMacsecAdvance interface {
	// FromProto unmarshals MacsecAdvance from protobuf object *otg.MacsecAdvance
	FromProto(msg *otg.MacsecAdvance) (MacsecAdvance, error)
	// FromPbText unmarshals MacsecAdvance from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecAdvance from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecAdvance from JSON text
	FromJson(value string) error
}

func (obj *macsecAdvance) Marshal() marshalMacsecAdvance {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecAdvance{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecAdvance) Unmarshal() unMarshalMacsecAdvance {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecAdvance{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecAdvance) ToProto() (*otg.MacsecAdvance, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecAdvance) FromProto(msg *otg.MacsecAdvance) (MacsecAdvance, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecAdvance) ToPbText() (string, error) {
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

func (m *unMarshalmacsecAdvance) FromPbText(value string) error {
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

func (m *marshalmacsecAdvance) ToYaml() (string, error) {
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

func (m *unMarshalmacsecAdvance) FromYaml(value string) error {
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

func (m *marshalmacsecAdvance) ToJson() (string, error) {
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

func (m *unMarshalmacsecAdvance) FromJson(value string) error {
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

func (obj *macsecAdvance) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecAdvance) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecAdvance) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecAdvance) Clone() (MacsecAdvance, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecAdvance()
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

// MacsecAdvance is a container of advance properties for a MACsec interface.
type MacsecAdvance interface {
	Validation
	// msg marshals MacsecAdvance to protobuf object *otg.MacsecAdvance
	// and doesn't set defaults
	msg() *otg.MacsecAdvance
	// setMsg unmarshals MacsecAdvance from protobuf object *otg.MacsecAdvance
	// and doesn't set defaults
	setMsg(*otg.MacsecAdvance) MacsecAdvance
	// provides marshal interface
	Marshal() marshalMacsecAdvance
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecAdvance
	// validate validates MacsecAdvance
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecAdvance, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RekeyMode returns MacsecAdvanceRekeyModeEnum, set in MacsecAdvance
	RekeyMode() MacsecAdvanceRekeyModeEnum
	// SetRekeyMode assigns MacsecAdvanceRekeyModeEnum provided by user to MacsecAdvance
	SetRekeyMode(value MacsecAdvanceRekeyModeEnum) MacsecAdvance
	// HasRekeyMode checks if RekeyMode has been set in MacsecAdvance
	HasRekeyMode() bool
}

type MacsecAdvanceRekeyModeEnum string

// Enum of RekeyMode on MacsecAdvance
var MacsecAdvanceRekeyMode = struct {
	TIMER_BASED MacsecAdvanceRekeyModeEnum
	PN_BASED    MacsecAdvanceRekeyModeEnum
}{
	TIMER_BASED: MacsecAdvanceRekeyModeEnum("timer_based"),
	PN_BASED:    MacsecAdvanceRekeyModeEnum("pn_based"),
}

func (obj *macsecAdvance) RekeyMode() MacsecAdvanceRekeyModeEnum {
	return MacsecAdvanceRekeyModeEnum(obj.obj.RekeyMode.Enum().String())
}

// Rekey mode.
// RekeyMode returns a string
func (obj *macsecAdvance) HasRekeyMode() bool {
	return obj.obj.RekeyMode != nil
}

func (obj *macsecAdvance) SetRekeyMode(value MacsecAdvanceRekeyModeEnum) MacsecAdvance {
	intValue, ok := otg.MacsecAdvance_RekeyMode_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecAdvanceRekeyModeEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecAdvance_RekeyMode_Enum(intValue)
	obj.obj.RekeyMode = &enumValue

	return obj
}

func (obj *macsecAdvance) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *macsecAdvance) setDefault() {
	if obj.obj.RekeyMode == nil {
		obj.SetRekeyMode(MacsecAdvanceRekeyMode.TIMER_BASED)

	}

}
