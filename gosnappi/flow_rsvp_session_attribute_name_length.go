package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPSessionAttributeNameLength *****
type flowRSVPSessionAttributeNameLength struct {
	validation
	obj          *otg.FlowRSVPSessionAttributeNameLength
	marshaller   marshalFlowRSVPSessionAttributeNameLength
	unMarshaller unMarshalFlowRSVPSessionAttributeNameLength
}

func NewFlowRSVPSessionAttributeNameLength() FlowRSVPSessionAttributeNameLength {
	obj := flowRSVPSessionAttributeNameLength{obj: &otg.FlowRSVPSessionAttributeNameLength{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPSessionAttributeNameLength) msg() *otg.FlowRSVPSessionAttributeNameLength {
	return obj.obj
}

func (obj *flowRSVPSessionAttributeNameLength) setMsg(msg *otg.FlowRSVPSessionAttributeNameLength) FlowRSVPSessionAttributeNameLength {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPSessionAttributeNameLength struct {
	obj *flowRSVPSessionAttributeNameLength
}

type marshalFlowRSVPSessionAttributeNameLength interface {
	// ToProto marshals FlowRSVPSessionAttributeNameLength to protobuf object *otg.FlowRSVPSessionAttributeNameLength
	ToProto() (*otg.FlowRSVPSessionAttributeNameLength, error)
	// ToPbText marshals FlowRSVPSessionAttributeNameLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPSessionAttributeNameLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPSessionAttributeNameLength to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPSessionAttributeNameLength struct {
	obj *flowRSVPSessionAttributeNameLength
}

type unMarshalFlowRSVPSessionAttributeNameLength interface {
	// FromProto unmarshals FlowRSVPSessionAttributeNameLength from protobuf object *otg.FlowRSVPSessionAttributeNameLength
	FromProto(msg *otg.FlowRSVPSessionAttributeNameLength) (FlowRSVPSessionAttributeNameLength, error)
	// FromPbText unmarshals FlowRSVPSessionAttributeNameLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPSessionAttributeNameLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPSessionAttributeNameLength from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPSessionAttributeNameLength) Marshal() marshalFlowRSVPSessionAttributeNameLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPSessionAttributeNameLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPSessionAttributeNameLength) Unmarshal() unMarshalFlowRSVPSessionAttributeNameLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPSessionAttributeNameLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPSessionAttributeNameLength) ToProto() (*otg.FlowRSVPSessionAttributeNameLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPSessionAttributeNameLength) FromProto(msg *otg.FlowRSVPSessionAttributeNameLength) (FlowRSVPSessionAttributeNameLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPSessionAttributeNameLength) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPSessionAttributeNameLength) FromPbText(value string) error {
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

func (m *marshalflowRSVPSessionAttributeNameLength) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPSessionAttributeNameLength) FromYaml(value string) error {
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

func (m *marshalflowRSVPSessionAttributeNameLength) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPSessionAttributeNameLength) FromJson(value string) error {
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

func (obj *flowRSVPSessionAttributeNameLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPSessionAttributeNameLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPSessionAttributeNameLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPSessionAttributeNameLength) Clone() (FlowRSVPSessionAttributeNameLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPSessionAttributeNameLength()
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

// FlowRSVPSessionAttributeNameLength is description is TBD
type FlowRSVPSessionAttributeNameLength interface {
	Validation
	// msg marshals FlowRSVPSessionAttributeNameLength to protobuf object *otg.FlowRSVPSessionAttributeNameLength
	// and doesn't set defaults
	msg() *otg.FlowRSVPSessionAttributeNameLength
	// setMsg unmarshals FlowRSVPSessionAttributeNameLength from protobuf object *otg.FlowRSVPSessionAttributeNameLength
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPSessionAttributeNameLength) FlowRSVPSessionAttributeNameLength
	// provides marshal interface
	Marshal() marshalFlowRSVPSessionAttributeNameLength
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPSessionAttributeNameLength
	// validate validates FlowRSVPSessionAttributeNameLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPSessionAttributeNameLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowRSVPSessionAttributeNameLengthChoiceEnum, set in FlowRSVPSessionAttributeNameLength
	Choice() FlowRSVPSessionAttributeNameLengthChoiceEnum
	// setChoice assigns FlowRSVPSessionAttributeNameLengthChoiceEnum provided by user to FlowRSVPSessionAttributeNameLength
	setChoice(value FlowRSVPSessionAttributeNameLengthChoiceEnum) FlowRSVPSessionAttributeNameLength
	// HasChoice checks if Choice has been set in FlowRSVPSessionAttributeNameLength
	HasChoice() bool
	// Auto returns uint32, set in FlowRSVPSessionAttributeNameLength.
	Auto() uint32
	// HasAuto checks if Auto has been set in FlowRSVPSessionAttributeNameLength
	HasAuto() bool
	// Value returns uint32, set in FlowRSVPSessionAttributeNameLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to FlowRSVPSessionAttributeNameLength
	SetValue(value uint32) FlowRSVPSessionAttributeNameLength
	// HasValue checks if Value has been set in FlowRSVPSessionAttributeNameLength
	HasValue() bool
}

type FlowRSVPSessionAttributeNameLengthChoiceEnum string

// Enum of Choice on FlowRSVPSessionAttributeNameLength
var FlowRSVPSessionAttributeNameLengthChoice = struct {
	AUTO  FlowRSVPSessionAttributeNameLengthChoiceEnum
	VALUE FlowRSVPSessionAttributeNameLengthChoiceEnum
}{
	AUTO:  FlowRSVPSessionAttributeNameLengthChoiceEnum("auto"),
	VALUE: FlowRSVPSessionAttributeNameLengthChoiceEnum("value"),
}

func (obj *flowRSVPSessionAttributeNameLength) Choice() FlowRSVPSessionAttributeNameLengthChoiceEnum {
	return FlowRSVPSessionAttributeNameLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// auto or configured value.
// Choice returns a string
func (obj *flowRSVPSessionAttributeNameLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowRSVPSessionAttributeNameLength) setChoice(value FlowRSVPSessionAttributeNameLengthChoiceEnum) FlowRSVPSessionAttributeNameLength {
	intValue, ok := otg.FlowRSVPSessionAttributeNameLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowRSVPSessionAttributeNameLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowRSVPSessionAttributeNameLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Value = nil
	obj.obj.Auto = nil

	if value == FlowRSVPSessionAttributeNameLengthChoice.AUTO {
		defaultValue := uint32(0)
		obj.obj.Auto = &defaultValue
	}

	if value == FlowRSVPSessionAttributeNameLengthChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	return obj
}

// OTG will provide a system generated value for this property.  If OTG is unable to generate a value the default value must be used.
// Auto returns a uint32
func (obj *flowRSVPSessionAttributeNameLength) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(FlowRSVPSessionAttributeNameLengthChoice.AUTO)
	}

	return *obj.obj.Auto

}

// OTG will provide a system generated value for this property.  If OTG is unable to generate a value the default value must be used.
// Auto returns a uint32
func (obj *flowRSVPSessionAttributeNameLength) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Value returns a uint32
func (obj *flowRSVPSessionAttributeNameLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(FlowRSVPSessionAttributeNameLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *flowRSVPSessionAttributeNameLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the FlowRSVPSessionAttributeNameLength object
func (obj *flowRSVPSessionAttributeNameLength) SetValue(value uint32) FlowRSVPSessionAttributeNameLength {
	obj.setChoice(FlowRSVPSessionAttributeNameLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

func (obj *flowRSVPSessionAttributeNameLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowRSVPSessionAttributeNameLength.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowRSVPSessionAttributeNameLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

}

func (obj *flowRSVPSessionAttributeNameLength) setDefault() {
	var choices_set int = 0
	var choice FlowRSVPSessionAttributeNameLengthChoiceEnum

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = FlowRSVPSessionAttributeNameLengthChoice.AUTO
	}

	if obj.obj.Value != nil {
		choices_set += 1
		choice = FlowRSVPSessionAttributeNameLengthChoice.VALUE
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowRSVPSessionAttributeNameLengthChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowRSVPSessionAttributeNameLength")
			}
		} else {
			intVal := otg.FlowRSVPSessionAttributeNameLength_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowRSVPSessionAttributeNameLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
