package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPExplicitRouteASNumberLength *****
type flowRSVPExplicitRouteASNumberLength struct {
	validation
	obj          *otg.FlowRSVPExplicitRouteASNumberLength
	marshaller   marshalFlowRSVPExplicitRouteASNumberLength
	unMarshaller unMarshalFlowRSVPExplicitRouteASNumberLength
}

func NewFlowRSVPExplicitRouteASNumberLength() FlowRSVPExplicitRouteASNumberLength {
	obj := flowRSVPExplicitRouteASNumberLength{obj: &otg.FlowRSVPExplicitRouteASNumberLength{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPExplicitRouteASNumberLength) msg() *otg.FlowRSVPExplicitRouteASNumberLength {
	return obj.obj
}

func (obj *flowRSVPExplicitRouteASNumberLength) setMsg(msg *otg.FlowRSVPExplicitRouteASNumberLength) FlowRSVPExplicitRouteASNumberLength {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPExplicitRouteASNumberLength struct {
	obj *flowRSVPExplicitRouteASNumberLength
}

type marshalFlowRSVPExplicitRouteASNumberLength interface {
	// ToProto marshals FlowRSVPExplicitRouteASNumberLength to protobuf object *otg.FlowRSVPExplicitRouteASNumberLength
	ToProto() (*otg.FlowRSVPExplicitRouteASNumberLength, error)
	// ToPbText marshals FlowRSVPExplicitRouteASNumberLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPExplicitRouteASNumberLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPExplicitRouteASNumberLength to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPExplicitRouteASNumberLength struct {
	obj *flowRSVPExplicitRouteASNumberLength
}

type unMarshalFlowRSVPExplicitRouteASNumberLength interface {
	// FromProto unmarshals FlowRSVPExplicitRouteASNumberLength from protobuf object *otg.FlowRSVPExplicitRouteASNumberLength
	FromProto(msg *otg.FlowRSVPExplicitRouteASNumberLength) (FlowRSVPExplicitRouteASNumberLength, error)
	// FromPbText unmarshals FlowRSVPExplicitRouteASNumberLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPExplicitRouteASNumberLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPExplicitRouteASNumberLength from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPExplicitRouteASNumberLength) Marshal() marshalFlowRSVPExplicitRouteASNumberLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPExplicitRouteASNumberLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPExplicitRouteASNumberLength) Unmarshal() unMarshalFlowRSVPExplicitRouteASNumberLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPExplicitRouteASNumberLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPExplicitRouteASNumberLength) ToProto() (*otg.FlowRSVPExplicitRouteASNumberLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPExplicitRouteASNumberLength) FromProto(msg *otg.FlowRSVPExplicitRouteASNumberLength) (FlowRSVPExplicitRouteASNumberLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPExplicitRouteASNumberLength) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPExplicitRouteASNumberLength) FromPbText(value string) error {
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

func (m *marshalflowRSVPExplicitRouteASNumberLength) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPExplicitRouteASNumberLength) FromYaml(value string) error {
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

func (m *marshalflowRSVPExplicitRouteASNumberLength) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPExplicitRouteASNumberLength) FromJson(value string) error {
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

func (obj *flowRSVPExplicitRouteASNumberLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPExplicitRouteASNumberLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPExplicitRouteASNumberLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPExplicitRouteASNumberLength) Clone() (FlowRSVPExplicitRouteASNumberLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPExplicitRouteASNumberLength()
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

// FlowRSVPExplicitRouteASNumberLength is description is TBD
type FlowRSVPExplicitRouteASNumberLength interface {
	Validation
	// msg marshals FlowRSVPExplicitRouteASNumberLength to protobuf object *otg.FlowRSVPExplicitRouteASNumberLength
	// and doesn't set defaults
	msg() *otg.FlowRSVPExplicitRouteASNumberLength
	// setMsg unmarshals FlowRSVPExplicitRouteASNumberLength from protobuf object *otg.FlowRSVPExplicitRouteASNumberLength
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPExplicitRouteASNumberLength) FlowRSVPExplicitRouteASNumberLength
	// provides marshal interface
	Marshal() marshalFlowRSVPExplicitRouteASNumberLength
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPExplicitRouteASNumberLength
	// validate validates FlowRSVPExplicitRouteASNumberLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPExplicitRouteASNumberLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowRSVPExplicitRouteASNumberLengthChoiceEnum, set in FlowRSVPExplicitRouteASNumberLength
	Choice() FlowRSVPExplicitRouteASNumberLengthChoiceEnum
	// setChoice assigns FlowRSVPExplicitRouteASNumberLengthChoiceEnum provided by user to FlowRSVPExplicitRouteASNumberLength
	setChoice(value FlowRSVPExplicitRouteASNumberLengthChoiceEnum) FlowRSVPExplicitRouteASNumberLength
	// HasChoice checks if Choice has been set in FlowRSVPExplicitRouteASNumberLength
	HasChoice() bool
	// Auto returns uint32, set in FlowRSVPExplicitRouteASNumberLength.
	Auto() uint32
	// HasAuto checks if Auto has been set in FlowRSVPExplicitRouteASNumberLength
	HasAuto() bool
	// Value returns uint32, set in FlowRSVPExplicitRouteASNumberLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to FlowRSVPExplicitRouteASNumberLength
	SetValue(value uint32) FlowRSVPExplicitRouteASNumberLength
	// HasValue checks if Value has been set in FlowRSVPExplicitRouteASNumberLength
	HasValue() bool
}

type FlowRSVPExplicitRouteASNumberLengthChoiceEnum string

// Enum of Choice on FlowRSVPExplicitRouteASNumberLength
var FlowRSVPExplicitRouteASNumberLengthChoice = struct {
	AUTO  FlowRSVPExplicitRouteASNumberLengthChoiceEnum
	VALUE FlowRSVPExplicitRouteASNumberLengthChoiceEnum
}{
	AUTO:  FlowRSVPExplicitRouteASNumberLengthChoiceEnum("auto"),
	VALUE: FlowRSVPExplicitRouteASNumberLengthChoiceEnum("value"),
}

func (obj *flowRSVPExplicitRouteASNumberLength) Choice() FlowRSVPExplicitRouteASNumberLengthChoiceEnum {
	return FlowRSVPExplicitRouteASNumberLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// auto or configured value.
// Choice returns a string
func (obj *flowRSVPExplicitRouteASNumberLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowRSVPExplicitRouteASNumberLength) setChoice(value FlowRSVPExplicitRouteASNumberLengthChoiceEnum) FlowRSVPExplicitRouteASNumberLength {
	intValue, ok := otg.FlowRSVPExplicitRouteASNumberLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowRSVPExplicitRouteASNumberLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowRSVPExplicitRouteASNumberLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Value = nil
	obj.obj.Auto = nil

	if value == FlowRSVPExplicitRouteASNumberLengthChoice.AUTO {
		defaultValue := uint32(4)
		obj.obj.Auto = &defaultValue
	}

	if value == FlowRSVPExplicitRouteASNumberLengthChoice.VALUE {
		defaultValue := uint32(4)
		obj.obj.Value = &defaultValue
	}

	return obj
}

// The OTG implementation will provide a system generated value for this property.  If the OTG implementation is unable to generate a value the default value must be used.
// Auto returns a uint32
func (obj *flowRSVPExplicitRouteASNumberLength) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(FlowRSVPExplicitRouteASNumberLengthChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation will provide a system generated value for this property.  If the OTG implementation is unable to generate a value the default value must be used.
// Auto returns a uint32
func (obj *flowRSVPExplicitRouteASNumberLength) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Value returns a uint32
func (obj *flowRSVPExplicitRouteASNumberLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(FlowRSVPExplicitRouteASNumberLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *flowRSVPExplicitRouteASNumberLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the FlowRSVPExplicitRouteASNumberLength object
func (obj *flowRSVPExplicitRouteASNumberLength) SetValue(value uint32) FlowRSVPExplicitRouteASNumberLength {
	obj.setChoice(FlowRSVPExplicitRouteASNumberLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

func (obj *flowRSVPExplicitRouteASNumberLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowRSVPExplicitRouteASNumberLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

}

func (obj *flowRSVPExplicitRouteASNumberLength) setDefault() {
	var choices_set int = 0
	var choice FlowRSVPExplicitRouteASNumberLengthChoiceEnum

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = FlowRSVPExplicitRouteASNumberLengthChoice.AUTO
	}

	if obj.obj.Value != nil {
		choices_set += 1
		choice = FlowRSVPExplicitRouteASNumberLengthChoice.VALUE
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowRSVPExplicitRouteASNumberLengthChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowRSVPExplicitRouteASNumberLength")
			}
		} else {
			intVal := otg.FlowRSVPExplicitRouteASNumberLength_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowRSVPExplicitRouteASNumberLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
