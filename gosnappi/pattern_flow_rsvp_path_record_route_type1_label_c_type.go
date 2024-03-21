package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathRecordRouteType1LabelCType *****
type patternFlowRSVPPathRecordRouteType1LabelCType struct {
	validation
	obj          *otg.PatternFlowRSVPPathRecordRouteType1LabelCType
	marshaller   marshalPatternFlowRSVPPathRecordRouteType1LabelCType
	unMarshaller unMarshalPatternFlowRSVPPathRecordRouteType1LabelCType
}

func NewPatternFlowRSVPPathRecordRouteType1LabelCType() PatternFlowRSVPPathRecordRouteType1LabelCType {
	obj := patternFlowRSVPPathRecordRouteType1LabelCType{obj: &otg.PatternFlowRSVPPathRecordRouteType1LabelCType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathRecordRouteType1LabelCType) msg() *otg.PatternFlowRSVPPathRecordRouteType1LabelCType {
	return obj.obj
}

func (obj *patternFlowRSVPPathRecordRouteType1LabelCType) setMsg(msg *otg.PatternFlowRSVPPathRecordRouteType1LabelCType) PatternFlowRSVPPathRecordRouteType1LabelCType {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathRecordRouteType1LabelCType struct {
	obj *patternFlowRSVPPathRecordRouteType1LabelCType
}

type marshalPatternFlowRSVPPathRecordRouteType1LabelCType interface {
	// ToProto marshals PatternFlowRSVPPathRecordRouteType1LabelCType to protobuf object *otg.PatternFlowRSVPPathRecordRouteType1LabelCType
	ToProto() (*otg.PatternFlowRSVPPathRecordRouteType1LabelCType, error)
	// ToPbText marshals PatternFlowRSVPPathRecordRouteType1LabelCType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathRecordRouteType1LabelCType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathRecordRouteType1LabelCType to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathRecordRouteType1LabelCType struct {
	obj *patternFlowRSVPPathRecordRouteType1LabelCType
}

type unMarshalPatternFlowRSVPPathRecordRouteType1LabelCType interface {
	// FromProto unmarshals PatternFlowRSVPPathRecordRouteType1LabelCType from protobuf object *otg.PatternFlowRSVPPathRecordRouteType1LabelCType
	FromProto(msg *otg.PatternFlowRSVPPathRecordRouteType1LabelCType) (PatternFlowRSVPPathRecordRouteType1LabelCType, error)
	// FromPbText unmarshals PatternFlowRSVPPathRecordRouteType1LabelCType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathRecordRouteType1LabelCType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathRecordRouteType1LabelCType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathRecordRouteType1LabelCType) Marshal() marshalPatternFlowRSVPPathRecordRouteType1LabelCType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathRecordRouteType1LabelCType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathRecordRouteType1LabelCType) Unmarshal() unMarshalPatternFlowRSVPPathRecordRouteType1LabelCType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathRecordRouteType1LabelCType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathRecordRouteType1LabelCType) ToProto() (*otg.PatternFlowRSVPPathRecordRouteType1LabelCType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathRecordRouteType1LabelCType) FromProto(msg *otg.PatternFlowRSVPPathRecordRouteType1LabelCType) (PatternFlowRSVPPathRecordRouteType1LabelCType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathRecordRouteType1LabelCType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRecordRouteType1LabelCType) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathRecordRouteType1LabelCType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRecordRouteType1LabelCType) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathRecordRouteType1LabelCType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRecordRouteType1LabelCType) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathRecordRouteType1LabelCType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathRecordRouteType1LabelCType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathRecordRouteType1LabelCType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathRecordRouteType1LabelCType) Clone() (PatternFlowRSVPPathRecordRouteType1LabelCType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathRecordRouteType1LabelCType()
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

// PatternFlowRSVPPathRecordRouteType1LabelCType is the C-Type of the included Label Object. Copied from the Label object.
type PatternFlowRSVPPathRecordRouteType1LabelCType interface {
	Validation
	// msg marshals PatternFlowRSVPPathRecordRouteType1LabelCType to protobuf object *otg.PatternFlowRSVPPathRecordRouteType1LabelCType
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathRecordRouteType1LabelCType
	// setMsg unmarshals PatternFlowRSVPPathRecordRouteType1LabelCType from protobuf object *otg.PatternFlowRSVPPathRecordRouteType1LabelCType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathRecordRouteType1LabelCType) PatternFlowRSVPPathRecordRouteType1LabelCType
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathRecordRouteType1LabelCType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathRecordRouteType1LabelCType
	// validate validates PatternFlowRSVPPathRecordRouteType1LabelCType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathRecordRouteType1LabelCType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathRecordRouteType1LabelCTypeChoiceEnum, set in PatternFlowRSVPPathRecordRouteType1LabelCType
	Choice() PatternFlowRSVPPathRecordRouteType1LabelCTypeChoiceEnum
	// setChoice assigns PatternFlowRSVPPathRecordRouteType1LabelCTypeChoiceEnum provided by user to PatternFlowRSVPPathRecordRouteType1LabelCType
	setChoice(value PatternFlowRSVPPathRecordRouteType1LabelCTypeChoiceEnum) PatternFlowRSVPPathRecordRouteType1LabelCType
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathRecordRouteType1LabelCType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRSVPPathRecordRouteType1LabelCType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRSVPPathRecordRouteType1LabelCType
	SetValue(value uint32) PatternFlowRSVPPathRecordRouteType1LabelCType
	// HasValue checks if Value has been set in PatternFlowRSVPPathRecordRouteType1LabelCType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRSVPPathRecordRouteType1LabelCType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRSVPPathRecordRouteType1LabelCType
	SetValues(value []uint32) PatternFlowRSVPPathRecordRouteType1LabelCType
}

type PatternFlowRSVPPathRecordRouteType1LabelCTypeChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathRecordRouteType1LabelCType
var PatternFlowRSVPPathRecordRouteType1LabelCTypeChoice = struct {
	VALUE  PatternFlowRSVPPathRecordRouteType1LabelCTypeChoiceEnum
	VALUES PatternFlowRSVPPathRecordRouteType1LabelCTypeChoiceEnum
}{
	VALUE:  PatternFlowRSVPPathRecordRouteType1LabelCTypeChoiceEnum("value"),
	VALUES: PatternFlowRSVPPathRecordRouteType1LabelCTypeChoiceEnum("values"),
}

func (obj *patternFlowRSVPPathRecordRouteType1LabelCType) Choice() PatternFlowRSVPPathRecordRouteType1LabelCTypeChoiceEnum {
	return PatternFlowRSVPPathRecordRouteType1LabelCTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathRecordRouteType1LabelCType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathRecordRouteType1LabelCType) setChoice(value PatternFlowRSVPPathRecordRouteType1LabelCTypeChoiceEnum) PatternFlowRSVPPathRecordRouteType1LabelCType {
	intValue, ok := otg.PatternFlowRSVPPathRecordRouteType1LabelCType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathRecordRouteType1LabelCTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathRecordRouteType1LabelCType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathRecordRouteType1LabelCTypeChoice.VALUE {
		defaultValue := uint32(1)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathRecordRouteType1LabelCTypeChoice.VALUES {
		defaultValue := []uint32{1}
		obj.obj.Values = defaultValue
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathRecordRouteType1LabelCType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathRecordRouteType1LabelCTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathRecordRouteType1LabelCType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRSVPPathRecordRouteType1LabelCType object
func (obj *patternFlowRSVPPathRecordRouteType1LabelCType) SetValue(value uint32) PatternFlowRSVPPathRecordRouteType1LabelCType {
	obj.setChoice(PatternFlowRSVPPathRecordRouteType1LabelCTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRSVPPathRecordRouteType1LabelCType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{1})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRSVPPathRecordRouteType1LabelCType object
func (obj *patternFlowRSVPPathRecordRouteType1LabelCType) SetValues(value []uint32) PatternFlowRSVPPathRecordRouteType1LabelCType {
	obj.setChoice(PatternFlowRSVPPathRecordRouteType1LabelCTypeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

func (obj *patternFlowRSVPPathRecordRouteType1LabelCType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathRecordRouteType1LabelCType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowRSVPPathRecordRouteType1LabelCType.Values <= 255 but Got %d", item))
			}

		}

	}

}

func (obj *patternFlowRSVPPathRecordRouteType1LabelCType) setDefault() {
	var choices_set int = 0
	var choice PatternFlowRSVPPathRecordRouteType1LabelCTypeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathRecordRouteType1LabelCTypeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowRSVPPathRecordRouteType1LabelCTypeChoice.VALUES
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowRSVPPathRecordRouteType1LabelCTypeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowRSVPPathRecordRouteType1LabelCType")
			}
		} else {
			intVal := otg.PatternFlowRSVPPathRecordRouteType1LabelCType_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowRSVPPathRecordRouteType1LabelCType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
