package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPObjectLength *****
type flowRSVPObjectLength struct {
	validation
	obj          *otg.FlowRSVPObjectLength
	marshaller   marshalFlowRSVPObjectLength
	unMarshaller unMarshalFlowRSVPObjectLength
}

func NewFlowRSVPObjectLength() FlowRSVPObjectLength {
	obj := flowRSVPObjectLength{obj: &otg.FlowRSVPObjectLength{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPObjectLength) msg() *otg.FlowRSVPObjectLength {
	return obj.obj
}

func (obj *flowRSVPObjectLength) setMsg(msg *otg.FlowRSVPObjectLength) FlowRSVPObjectLength {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPObjectLength struct {
	obj *flowRSVPObjectLength
}

type marshalFlowRSVPObjectLength interface {
	// ToProto marshals FlowRSVPObjectLength to protobuf object *otg.FlowRSVPObjectLength
	ToProto() (*otg.FlowRSVPObjectLength, error)
	// ToPbText marshals FlowRSVPObjectLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPObjectLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPObjectLength to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPObjectLength struct {
	obj *flowRSVPObjectLength
}

type unMarshalFlowRSVPObjectLength interface {
	// FromProto unmarshals FlowRSVPObjectLength from protobuf object *otg.FlowRSVPObjectLength
	FromProto(msg *otg.FlowRSVPObjectLength) (FlowRSVPObjectLength, error)
	// FromPbText unmarshals FlowRSVPObjectLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPObjectLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPObjectLength from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPObjectLength) Marshal() marshalFlowRSVPObjectLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPObjectLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPObjectLength) Unmarshal() unMarshalFlowRSVPObjectLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPObjectLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPObjectLength) ToProto() (*otg.FlowRSVPObjectLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPObjectLength) FromProto(msg *otg.FlowRSVPObjectLength) (FlowRSVPObjectLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPObjectLength) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPObjectLength) FromPbText(value string) error {
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

func (m *marshalflowRSVPObjectLength) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPObjectLength) FromYaml(value string) error {
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

func (m *marshalflowRSVPObjectLength) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPObjectLength) FromJson(value string) error {
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

func (obj *flowRSVPObjectLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPObjectLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPObjectLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPObjectLength) Clone() (FlowRSVPObjectLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPObjectLength()
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

// FlowRSVPObjectLength is description is TBD
type FlowRSVPObjectLength interface {
	Validation
	// msg marshals FlowRSVPObjectLength to protobuf object *otg.FlowRSVPObjectLength
	// and doesn't set defaults
	msg() *otg.FlowRSVPObjectLength
	// setMsg unmarshals FlowRSVPObjectLength from protobuf object *otg.FlowRSVPObjectLength
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPObjectLength) FlowRSVPObjectLength
	// provides marshal interface
	Marshal() marshalFlowRSVPObjectLength
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPObjectLength
	// validate validates FlowRSVPObjectLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPObjectLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowRSVPObjectLengthChoiceEnum, set in FlowRSVPObjectLength
	Choice() FlowRSVPObjectLengthChoiceEnum
	// setChoice assigns FlowRSVPObjectLengthChoiceEnum provided by user to FlowRSVPObjectLength
	setChoice(value FlowRSVPObjectLengthChoiceEnum) FlowRSVPObjectLength
	// HasChoice checks if Choice has been set in FlowRSVPObjectLength
	HasChoice() bool
	// Auto returns uint32, set in FlowRSVPObjectLength.
	Auto() uint32
	// HasAuto checks if Auto has been set in FlowRSVPObjectLength
	HasAuto() bool
	// Value returns uint32, set in FlowRSVPObjectLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to FlowRSVPObjectLength
	SetValue(value uint32) FlowRSVPObjectLength
	// HasValue checks if Value has been set in FlowRSVPObjectLength
	HasValue() bool
}

type FlowRSVPObjectLengthChoiceEnum string

// Enum of Choice on FlowRSVPObjectLength
var FlowRSVPObjectLengthChoice = struct {
	AUTO  FlowRSVPObjectLengthChoiceEnum
	VALUE FlowRSVPObjectLengthChoiceEnum
}{
	AUTO:  FlowRSVPObjectLengthChoiceEnum("auto"),
	VALUE: FlowRSVPObjectLengthChoiceEnum("value"),
}

func (obj *flowRSVPObjectLength) Choice() FlowRSVPObjectLengthChoiceEnum {
	return FlowRSVPObjectLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// auto or configured value.
// Choice returns a string
func (obj *flowRSVPObjectLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowRSVPObjectLength) setChoice(value FlowRSVPObjectLengthChoiceEnum) FlowRSVPObjectLength {
	intValue, ok := otg.FlowRSVPObjectLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowRSVPObjectLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowRSVPObjectLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Value = nil
	obj.obj.Auto = nil

	if value == FlowRSVPObjectLengthChoice.AUTO {
		defaultValue := uint32(4)
		obj.obj.Auto = &defaultValue
	}

	if value == FlowRSVPObjectLengthChoice.VALUE {
		defaultValue := uint32(4)
		obj.obj.Value = &defaultValue
	}

	return obj
}

// OTG will provide a system generated value for this property.  If OTG is unable to generate a value the default value must be used.
// Auto returns a uint32
func (obj *flowRSVPObjectLength) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(FlowRSVPObjectLengthChoice.AUTO)
	}

	return *obj.obj.Auto

}

// OTG will provide a system generated value for this property.  If OTG is unable to generate a value the default value must be used.
// Auto returns a uint32
func (obj *flowRSVPObjectLength) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Value returns a uint32
func (obj *flowRSVPObjectLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(FlowRSVPObjectLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *flowRSVPObjectLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the FlowRSVPObjectLength object
func (obj *flowRSVPObjectLength) SetValue(value uint32) FlowRSVPObjectLength {
	obj.setChoice(FlowRSVPObjectLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

func (obj *flowRSVPObjectLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto < 4 || *obj.obj.Auto > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("4 <= FlowRSVPObjectLength.Auto <= 65535 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Value != nil {

		if *obj.obj.Value < 4 || *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("4 <= FlowRSVPObjectLength.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

}

func (obj *flowRSVPObjectLength) setDefault() {
	var choices_set int = 0
	var choice FlowRSVPObjectLengthChoiceEnum

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = FlowRSVPObjectLengthChoice.AUTO
	}

	if obj.obj.Value != nil {
		choices_set += 1
		choice = FlowRSVPObjectLengthChoice.VALUE
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowRSVPObjectLengthChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowRSVPObjectLength")
			}
		} else {
			intVal := otg.FlowRSVPObjectLength_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowRSVPObjectLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
