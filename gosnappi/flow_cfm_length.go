package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfmLength *****
type flowCfmLength struct {
	validation
	obj          *otg.FlowCfmLength
	marshaller   marshalFlowCfmLength
	unMarshaller unMarshalFlowCfmLength
}

func NewFlowCfmLength() FlowCfmLength {
	obj := flowCfmLength{obj: &otg.FlowCfmLength{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfmLength) msg() *otg.FlowCfmLength {
	return obj.obj
}

func (obj *flowCfmLength) setMsg(msg *otg.FlowCfmLength) FlowCfmLength {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfmLength struct {
	obj *flowCfmLength
}

type marshalFlowCfmLength interface {
	// ToProto marshals FlowCfmLength to protobuf object *otg.FlowCfmLength
	ToProto() (*otg.FlowCfmLength, error)
	// ToPbText marshals FlowCfmLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfmLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfmLength to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfmLength struct {
	obj *flowCfmLength
}

type unMarshalFlowCfmLength interface {
	// FromProto unmarshals FlowCfmLength from protobuf object *otg.FlowCfmLength
	FromProto(msg *otg.FlowCfmLength) (FlowCfmLength, error)
	// FromPbText unmarshals FlowCfmLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfmLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfmLength from JSON text
	FromJson(value string) error
}

func (obj *flowCfmLength) Marshal() marshalFlowCfmLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfmLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfmLength) Unmarshal() unMarshalFlowCfmLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfmLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfmLength) ToProto() (*otg.FlowCfmLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfmLength) FromProto(msg *otg.FlowCfmLength) (FlowCfmLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfmLength) ToPbText() (string, error) {
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

func (m *unMarshalflowCfmLength) FromPbText(value string) error {
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

func (m *marshalflowCfmLength) ToYaml() (string, error) {
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

func (m *unMarshalflowCfmLength) FromYaml(value string) error {
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

func (m *marshalflowCfmLength) ToJson() (string, error) {
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

func (m *unMarshalflowCfmLength) FromJson(value string) error {
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

func (obj *flowCfmLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfmLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfmLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfmLength) Clone() (FlowCfmLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfmLength()
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

// FlowCfmLength is length field with auto-compute or explicit-value choice. When auto, the implementation derives the value from the associated data field.
type FlowCfmLength interface {
	Validation
	// msg marshals FlowCfmLength to protobuf object *otg.FlowCfmLength
	// and doesn't set defaults
	msg() *otg.FlowCfmLength
	// setMsg unmarshals FlowCfmLength from protobuf object *otg.FlowCfmLength
	// and doesn't set defaults
	setMsg(*otg.FlowCfmLength) FlowCfmLength
	// provides marshal interface
	Marshal() marshalFlowCfmLength
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfmLength
	// validate validates FlowCfmLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfmLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowCfmLengthChoiceEnum, set in FlowCfmLength
	Choice() FlowCfmLengthChoiceEnum
	// setChoice assigns FlowCfmLengthChoiceEnum provided by user to FlowCfmLength
	setChoice(value FlowCfmLengthChoiceEnum) FlowCfmLength
	// HasChoice checks if Choice has been set in FlowCfmLength
	HasChoice() bool
	// Auto returns uint32, set in FlowCfmLength.
	Auto() uint32
	// HasAuto checks if Auto has been set in FlowCfmLength
	HasAuto() bool
	// Value returns uint32, set in FlowCfmLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to FlowCfmLength
	SetValue(value uint32) FlowCfmLength
	// HasValue checks if Value has been set in FlowCfmLength
	HasValue() bool
}

type FlowCfmLengthChoiceEnum string

// Enum of Choice on FlowCfmLength
var FlowCfmLengthChoice = struct {
	AUTO  FlowCfmLengthChoiceEnum
	VALUE FlowCfmLengthChoiceEnum
}{
	AUTO:  FlowCfmLengthChoiceEnum("auto"),
	VALUE: FlowCfmLengthChoiceEnum("value"),
}

func (obj *flowCfmLength) Choice() FlowCfmLengthChoiceEnum {
	return FlowCfmLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// auto or configured value.
// Choice returns a string
func (obj *flowCfmLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowCfmLength) setChoice(value FlowCfmLengthChoiceEnum) FlowCfmLength {
	intValue, ok := otg.FlowCfmLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowCfmLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowCfmLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Value = nil
	obj.obj.Auto = nil

	if value == FlowCfmLengthChoice.AUTO {
		defaultValue := uint32(1)
		obj.obj.Auto = &defaultValue
	}

	if value == FlowCfmLengthChoice.VALUE {
		defaultValue := uint32(1)
		obj.obj.Value = &defaultValue
	}

	return obj
}

// The OTG implementation will provide a system generated value for this property.  If the OTG implementation is unable to generate a value the default value must be used.
// Auto returns a uint32
func (obj *flowCfmLength) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(FlowCfmLengthChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation will provide a system generated value for this property.  If the OTG implementation is unable to generate a value the default value must be used.
// Auto returns a uint32
func (obj *flowCfmLength) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Value returns a uint32
func (obj *flowCfmLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(FlowCfmLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *flowCfmLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the FlowCfmLength object
func (obj *flowCfmLength) SetValue(value uint32) FlowCfmLength {
	obj.setChoice(FlowCfmLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

func (obj *flowCfmLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowCfmLength.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowCfmLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

}

func (obj *flowCfmLength) setDefault() {
	var choices_set int = 0
	var choice FlowCfmLengthChoiceEnum

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = FlowCfmLengthChoice.AUTO
	}

	if obj.obj.Value != nil {
		choices_set += 1
		choice = FlowCfmLengthChoice.VALUE
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowCfmLengthChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowCfmLength")
			}
		} else {
			intVal := otg.FlowCfmLength_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowCfmLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
