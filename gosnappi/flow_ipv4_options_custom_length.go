package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv4OptionsCustomLength *****
type flowIpv4OptionsCustomLength struct {
	validation
	obj          *otg.FlowIpv4OptionsCustomLength
	marshaller   marshalFlowIpv4OptionsCustomLength
	unMarshaller unMarshalFlowIpv4OptionsCustomLength
}

func NewFlowIpv4OptionsCustomLength() FlowIpv4OptionsCustomLength {
	obj := flowIpv4OptionsCustomLength{obj: &otg.FlowIpv4OptionsCustomLength{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv4OptionsCustomLength) msg() *otg.FlowIpv4OptionsCustomLength {
	return obj.obj
}

func (obj *flowIpv4OptionsCustomLength) setMsg(msg *otg.FlowIpv4OptionsCustomLength) FlowIpv4OptionsCustomLength {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv4OptionsCustomLength struct {
	obj *flowIpv4OptionsCustomLength
}

type marshalFlowIpv4OptionsCustomLength interface {
	// ToProto marshals FlowIpv4OptionsCustomLength to protobuf object *otg.FlowIpv4OptionsCustomLength
	ToProto() (*otg.FlowIpv4OptionsCustomLength, error)
	// ToPbText marshals FlowIpv4OptionsCustomLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv4OptionsCustomLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv4OptionsCustomLength to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv4OptionsCustomLength struct {
	obj *flowIpv4OptionsCustomLength
}

type unMarshalFlowIpv4OptionsCustomLength interface {
	// FromProto unmarshals FlowIpv4OptionsCustomLength from protobuf object *otg.FlowIpv4OptionsCustomLength
	FromProto(msg *otg.FlowIpv4OptionsCustomLength) (FlowIpv4OptionsCustomLength, error)
	// FromPbText unmarshals FlowIpv4OptionsCustomLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv4OptionsCustomLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv4OptionsCustomLength from JSON text
	FromJson(value string) error
}

func (obj *flowIpv4OptionsCustomLength) Marshal() marshalFlowIpv4OptionsCustomLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv4OptionsCustomLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv4OptionsCustomLength) Unmarshal() unMarshalFlowIpv4OptionsCustomLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv4OptionsCustomLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv4OptionsCustomLength) ToProto() (*otg.FlowIpv4OptionsCustomLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv4OptionsCustomLength) FromProto(msg *otg.FlowIpv4OptionsCustomLength) (FlowIpv4OptionsCustomLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv4OptionsCustomLength) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv4OptionsCustomLength) FromPbText(value string) error {
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

func (m *marshalflowIpv4OptionsCustomLength) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv4OptionsCustomLength) FromYaml(value string) error {
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

func (m *marshalflowIpv4OptionsCustomLength) ToJson() (string, error) {
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

func (m *unMarshalflowIpv4OptionsCustomLength) FromJson(value string) error {
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

func (obj *flowIpv4OptionsCustomLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv4OptionsCustomLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv4OptionsCustomLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv4OptionsCustomLength) Clone() (FlowIpv4OptionsCustomLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv4OptionsCustomLength()
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

// FlowIpv4OptionsCustomLength is length for custom options.
type FlowIpv4OptionsCustomLength interface {
	Validation
	// msg marshals FlowIpv4OptionsCustomLength to protobuf object *otg.FlowIpv4OptionsCustomLength
	// and doesn't set defaults
	msg() *otg.FlowIpv4OptionsCustomLength
	// setMsg unmarshals FlowIpv4OptionsCustomLength from protobuf object *otg.FlowIpv4OptionsCustomLength
	// and doesn't set defaults
	setMsg(*otg.FlowIpv4OptionsCustomLength) FlowIpv4OptionsCustomLength
	// provides marshal interface
	Marshal() marshalFlowIpv4OptionsCustomLength
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv4OptionsCustomLength
	// validate validates FlowIpv4OptionsCustomLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv4OptionsCustomLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowIpv4OptionsCustomLengthChoiceEnum, set in FlowIpv4OptionsCustomLength
	Choice() FlowIpv4OptionsCustomLengthChoiceEnum
	// setChoice assigns FlowIpv4OptionsCustomLengthChoiceEnum provided by user to FlowIpv4OptionsCustomLength
	setChoice(value FlowIpv4OptionsCustomLengthChoiceEnum) FlowIpv4OptionsCustomLength
	// HasChoice checks if Choice has been set in FlowIpv4OptionsCustomLength
	HasChoice() bool
	// Auto returns uint32, set in FlowIpv4OptionsCustomLength.
	Auto() uint32
	// HasAuto checks if Auto has been set in FlowIpv4OptionsCustomLength
	HasAuto() bool
	// Value returns uint32, set in FlowIpv4OptionsCustomLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to FlowIpv4OptionsCustomLength
	SetValue(value uint32) FlowIpv4OptionsCustomLength
	// HasValue checks if Value has been set in FlowIpv4OptionsCustomLength
	HasValue() bool
}

type FlowIpv4OptionsCustomLengthChoiceEnum string

// Enum of Choice on FlowIpv4OptionsCustomLength
var FlowIpv4OptionsCustomLengthChoice = struct {
	AUTO  FlowIpv4OptionsCustomLengthChoiceEnum
	VALUE FlowIpv4OptionsCustomLengthChoiceEnum
}{
	AUTO:  FlowIpv4OptionsCustomLengthChoiceEnum("auto"),
	VALUE: FlowIpv4OptionsCustomLengthChoiceEnum("value"),
}

func (obj *flowIpv4OptionsCustomLength) Choice() FlowIpv4OptionsCustomLengthChoiceEnum {
	return FlowIpv4OptionsCustomLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// auto or configured value.
// Choice returns a string
func (obj *flowIpv4OptionsCustomLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowIpv4OptionsCustomLength) setChoice(value FlowIpv4OptionsCustomLengthChoiceEnum) FlowIpv4OptionsCustomLength {
	intValue, ok := otg.FlowIpv4OptionsCustomLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowIpv4OptionsCustomLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowIpv4OptionsCustomLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Value = nil
	obj.obj.Auto = nil

	if value == FlowIpv4OptionsCustomLengthChoice.AUTO {
		defaultValue := uint32(0)
		obj.obj.Auto = &defaultValue
	}

	if value == FlowIpv4OptionsCustomLengthChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	return obj
}

// The OTG implementation can provide a system generated value for this property. If the OTG is unable to generate a value the default value must be used.
// Auto returns a uint32
func (obj *flowIpv4OptionsCustomLength) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(FlowIpv4OptionsCustomLengthChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated value for this property. If the OTG is unable to generate a value the default value must be used.
// Auto returns a uint32
func (obj *flowIpv4OptionsCustomLength) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Value returns a uint32
func (obj *flowIpv4OptionsCustomLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(FlowIpv4OptionsCustomLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *flowIpv4OptionsCustomLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the FlowIpv4OptionsCustomLength object
func (obj *flowIpv4OptionsCustomLength) SetValue(value uint32) FlowIpv4OptionsCustomLength {
	obj.setChoice(FlowIpv4OptionsCustomLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

func (obj *flowIpv4OptionsCustomLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *flowIpv4OptionsCustomLength) setDefault() {
	var choices_set int = 0
	var choice FlowIpv4OptionsCustomLengthChoiceEnum

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = FlowIpv4OptionsCustomLengthChoice.AUTO
	}

	if obj.obj.Value != nil {
		choices_set += 1
		choice = FlowIpv4OptionsCustomLengthChoice.VALUE
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowIpv4OptionsCustomLengthChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowIpv4OptionsCustomLength")
			}
		} else {
			intVal := otg.FlowIpv4OptionsCustomLength_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowIpv4OptionsCustomLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
