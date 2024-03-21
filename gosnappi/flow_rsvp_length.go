package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPLength *****
type flowRSVPLength struct {
	validation
	obj          *otg.FlowRSVPLength
	marshaller   marshalFlowRSVPLength
	unMarshaller unMarshalFlowRSVPLength
}

func NewFlowRSVPLength() FlowRSVPLength {
	obj := flowRSVPLength{obj: &otg.FlowRSVPLength{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPLength) msg() *otg.FlowRSVPLength {
	return obj.obj
}

func (obj *flowRSVPLength) setMsg(msg *otg.FlowRSVPLength) FlowRSVPLength {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPLength struct {
	obj *flowRSVPLength
}

type marshalFlowRSVPLength interface {
	// ToProto marshals FlowRSVPLength to protobuf object *otg.FlowRSVPLength
	ToProto() (*otg.FlowRSVPLength, error)
	// ToPbText marshals FlowRSVPLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPLength to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPLength struct {
	obj *flowRSVPLength
}

type unMarshalFlowRSVPLength interface {
	// FromProto unmarshals FlowRSVPLength from protobuf object *otg.FlowRSVPLength
	FromProto(msg *otg.FlowRSVPLength) (FlowRSVPLength, error)
	// FromPbText unmarshals FlowRSVPLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPLength from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPLength) Marshal() marshalFlowRSVPLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPLength) Unmarshal() unMarshalFlowRSVPLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPLength) ToProto() (*otg.FlowRSVPLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPLength) FromProto(msg *otg.FlowRSVPLength) (FlowRSVPLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPLength) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPLength) FromPbText(value string) error {
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

func (m *marshalflowRSVPLength) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPLength) FromYaml(value string) error {
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

func (m *marshalflowRSVPLength) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPLength) FromJson(value string) error {
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

func (obj *flowRSVPLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPLength) Clone() (FlowRSVPLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPLength()
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

// FlowRSVPLength is description is TBD
type FlowRSVPLength interface {
	Validation
	// msg marshals FlowRSVPLength to protobuf object *otg.FlowRSVPLength
	// and doesn't set defaults
	msg() *otg.FlowRSVPLength
	// setMsg unmarshals FlowRSVPLength from protobuf object *otg.FlowRSVPLength
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPLength) FlowRSVPLength
	// provides marshal interface
	Marshal() marshalFlowRSVPLength
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPLength
	// validate validates FlowRSVPLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowRSVPLengthChoiceEnum, set in FlowRSVPLength
	Choice() FlowRSVPLengthChoiceEnum
	// setChoice assigns FlowRSVPLengthChoiceEnum provided by user to FlowRSVPLength
	setChoice(value FlowRSVPLengthChoiceEnum) FlowRSVPLength
	// HasChoice checks if Choice has been set in FlowRSVPLength
	HasChoice() bool
	// Auto returns uint32, set in FlowRSVPLength.
	Auto() uint32
	// HasAuto checks if Auto has been set in FlowRSVPLength
	HasAuto() bool
	// Value returns uint32, set in FlowRSVPLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to FlowRSVPLength
	SetValue(value uint32) FlowRSVPLength
	// HasValue checks if Value has been set in FlowRSVPLength
	HasValue() bool
}

type FlowRSVPLengthChoiceEnum string

// Enum of Choice on FlowRSVPLength
var FlowRSVPLengthChoice = struct {
	AUTO  FlowRSVPLengthChoiceEnum
	VALUE FlowRSVPLengthChoiceEnum
}{
	AUTO:  FlowRSVPLengthChoiceEnum("auto"),
	VALUE: FlowRSVPLengthChoiceEnum("value"),
}

func (obj *flowRSVPLength) Choice() FlowRSVPLengthChoiceEnum {
	return FlowRSVPLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// auto or configured value.
// Choice returns a string
func (obj *flowRSVPLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowRSVPLength) setChoice(value FlowRSVPLengthChoiceEnum) FlowRSVPLength {
	intValue, ok := otg.FlowRSVPLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowRSVPLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowRSVPLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Value = nil
	obj.obj.Auto = nil

	if value == FlowRSVPLengthChoice.AUTO {
		defaultValue := uint32(0)
		obj.obj.Auto = &defaultValue
	}

	if value == FlowRSVPLengthChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	return obj
}

// OTG will provide a system generated value for this property.  If OTG is unable to generate a value the default value must be used.
// Auto returns a uint32
func (obj *flowRSVPLength) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(FlowRSVPLengthChoice.AUTO)
	}

	return *obj.obj.Auto

}

// OTG will provide a system generated value for this property.  If OTG is unable to generate a value the default value must be used.
// Auto returns a uint32
func (obj *flowRSVPLength) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Value returns a uint32
func (obj *flowRSVPLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(FlowRSVPLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *flowRSVPLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the FlowRSVPLength object
func (obj *flowRSVPLength) SetValue(value uint32) FlowRSVPLength {
	obj.setChoice(FlowRSVPLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

func (obj *flowRSVPLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowRSVPLength.Auto <= 65535 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowRSVPLength.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

}

func (obj *flowRSVPLength) setDefault() {
	var choices_set int = 0
	var choice FlowRSVPLengthChoiceEnum

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = FlowRSVPLengthChoice.AUTO
	}

	if obj.obj.Value != nil {
		choices_set += 1
		choice = FlowRSVPLengthChoice.VALUE
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowRSVPLengthChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowRSVPLength")
			}
		} else {
			intVal := otg.FlowRSVPLength_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowRSVPLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
