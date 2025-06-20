package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPRouteRecordLength *****
type flowRSVPRouteRecordLength struct {
	validation
	obj          *otg.FlowRSVPRouteRecordLength
	marshaller   marshalFlowRSVPRouteRecordLength
	unMarshaller unMarshalFlowRSVPRouteRecordLength
}

func NewFlowRSVPRouteRecordLength() FlowRSVPRouteRecordLength {
	obj := flowRSVPRouteRecordLength{obj: &otg.FlowRSVPRouteRecordLength{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPRouteRecordLength) msg() *otg.FlowRSVPRouteRecordLength {
	return obj.obj
}

func (obj *flowRSVPRouteRecordLength) setMsg(msg *otg.FlowRSVPRouteRecordLength) FlowRSVPRouteRecordLength {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPRouteRecordLength struct {
	obj *flowRSVPRouteRecordLength
}

type marshalFlowRSVPRouteRecordLength interface {
	// ToProto marshals FlowRSVPRouteRecordLength to protobuf object *otg.FlowRSVPRouteRecordLength
	ToProto() (*otg.FlowRSVPRouteRecordLength, error)
	// ToPbText marshals FlowRSVPRouteRecordLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPRouteRecordLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPRouteRecordLength to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPRouteRecordLength struct {
	obj *flowRSVPRouteRecordLength
}

type unMarshalFlowRSVPRouteRecordLength interface {
	// FromProto unmarshals FlowRSVPRouteRecordLength from protobuf object *otg.FlowRSVPRouteRecordLength
	FromProto(msg *otg.FlowRSVPRouteRecordLength) (FlowRSVPRouteRecordLength, error)
	// FromPbText unmarshals FlowRSVPRouteRecordLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPRouteRecordLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPRouteRecordLength from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPRouteRecordLength) Marshal() marshalFlowRSVPRouteRecordLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPRouteRecordLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPRouteRecordLength) Unmarshal() unMarshalFlowRSVPRouteRecordLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPRouteRecordLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPRouteRecordLength) ToProto() (*otg.FlowRSVPRouteRecordLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPRouteRecordLength) FromProto(msg *otg.FlowRSVPRouteRecordLength) (FlowRSVPRouteRecordLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPRouteRecordLength) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPRouteRecordLength) FromPbText(value string) error {
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

func (m *marshalflowRSVPRouteRecordLength) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPRouteRecordLength) FromYaml(value string) error {
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

func (m *marshalflowRSVPRouteRecordLength) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPRouteRecordLength) FromJson(value string) error {
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

func (obj *flowRSVPRouteRecordLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPRouteRecordLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPRouteRecordLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPRouteRecordLength) Clone() (FlowRSVPRouteRecordLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPRouteRecordLength()
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

// FlowRSVPRouteRecordLength is description is TBD
type FlowRSVPRouteRecordLength interface {
	Validation
	// msg marshals FlowRSVPRouteRecordLength to protobuf object *otg.FlowRSVPRouteRecordLength
	// and doesn't set defaults
	msg() *otg.FlowRSVPRouteRecordLength
	// setMsg unmarshals FlowRSVPRouteRecordLength from protobuf object *otg.FlowRSVPRouteRecordLength
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPRouteRecordLength) FlowRSVPRouteRecordLength
	// provides marshal interface
	Marshal() marshalFlowRSVPRouteRecordLength
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPRouteRecordLength
	// validate validates FlowRSVPRouteRecordLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPRouteRecordLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowRSVPRouteRecordLengthChoiceEnum, set in FlowRSVPRouteRecordLength
	Choice() FlowRSVPRouteRecordLengthChoiceEnum
	// setChoice assigns FlowRSVPRouteRecordLengthChoiceEnum provided by user to FlowRSVPRouteRecordLength
	setChoice(value FlowRSVPRouteRecordLengthChoiceEnum) FlowRSVPRouteRecordLength
	// HasChoice checks if Choice has been set in FlowRSVPRouteRecordLength
	HasChoice() bool
	// Auto returns uint32, set in FlowRSVPRouteRecordLength.
	Auto() uint32
	// HasAuto checks if Auto has been set in FlowRSVPRouteRecordLength
	HasAuto() bool
	// Value returns uint32, set in FlowRSVPRouteRecordLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to FlowRSVPRouteRecordLength
	SetValue(value uint32) FlowRSVPRouteRecordLength
	// HasValue checks if Value has been set in FlowRSVPRouteRecordLength
	HasValue() bool
}

type FlowRSVPRouteRecordLengthChoiceEnum string

// Enum of Choice on FlowRSVPRouteRecordLength
var FlowRSVPRouteRecordLengthChoice = struct {
	AUTO  FlowRSVPRouteRecordLengthChoiceEnum
	VALUE FlowRSVPRouteRecordLengthChoiceEnum
}{
	AUTO:  FlowRSVPRouteRecordLengthChoiceEnum("auto"),
	VALUE: FlowRSVPRouteRecordLengthChoiceEnum("value"),
}

func (obj *flowRSVPRouteRecordLength) Choice() FlowRSVPRouteRecordLengthChoiceEnum {
	return FlowRSVPRouteRecordLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// auto or configured value.
// Choice returns a string
func (obj *flowRSVPRouteRecordLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowRSVPRouteRecordLength) setChoice(value FlowRSVPRouteRecordLengthChoiceEnum) FlowRSVPRouteRecordLength {
	intValue, ok := otg.FlowRSVPRouteRecordLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowRSVPRouteRecordLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowRSVPRouteRecordLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Value = nil
	obj.obj.Auto = nil

	if value == FlowRSVPRouteRecordLengthChoice.AUTO {
		defaultValue := uint32(8)
		obj.obj.Auto = &defaultValue
	}

	if value == FlowRSVPRouteRecordLengthChoice.VALUE {
		defaultValue := uint32(8)
		obj.obj.Value = &defaultValue
	}

	return obj
}

// The OTG implementation will provide a system generated value for this property.  If the OTG implementation is unable to generate a value the default value must be used.
// Auto returns a uint32
func (obj *flowRSVPRouteRecordLength) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(FlowRSVPRouteRecordLengthChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation will provide a system generated value for this property.  If the OTG implementation is unable to generate a value the default value must be used.
// Auto returns a uint32
func (obj *flowRSVPRouteRecordLength) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Value returns a uint32
func (obj *flowRSVPRouteRecordLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(FlowRSVPRouteRecordLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *flowRSVPRouteRecordLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the FlowRSVPRouteRecordLength object
func (obj *flowRSVPRouteRecordLength) SetValue(value uint32) FlowRSVPRouteRecordLength {
	obj.setChoice(FlowRSVPRouteRecordLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

func (obj *flowRSVPRouteRecordLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowRSVPRouteRecordLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

}

func (obj *flowRSVPRouteRecordLength) setDefault() {
	var choices_set int = 0
	var choice FlowRSVPRouteRecordLengthChoiceEnum

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = FlowRSVPRouteRecordLengthChoice.AUTO
	}

	if obj.obj.Value != nil {
		choices_set += 1
		choice = FlowRSVPRouteRecordLengthChoice.VALUE
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowRSVPRouteRecordLengthChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowRSVPRouteRecordLength")
			}
		} else {
			intVal := otg.FlowRSVPRouteRecordLength_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowRSVPRouteRecordLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
