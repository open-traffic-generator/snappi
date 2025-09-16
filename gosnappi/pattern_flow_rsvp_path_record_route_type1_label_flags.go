package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathRecordRouteType1LabelFlags *****
type patternFlowRSVPPathRecordRouteType1LabelFlags struct {
	validation
	obj          *otg.PatternFlowRSVPPathRecordRouteType1LabelFlags
	marshaller   marshalPatternFlowRSVPPathRecordRouteType1LabelFlags
	unMarshaller unMarshalPatternFlowRSVPPathRecordRouteType1LabelFlags
}

func NewPatternFlowRSVPPathRecordRouteType1LabelFlags() PatternFlowRSVPPathRecordRouteType1LabelFlags {
	obj := patternFlowRSVPPathRecordRouteType1LabelFlags{obj: &otg.PatternFlowRSVPPathRecordRouteType1LabelFlags{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathRecordRouteType1LabelFlags) msg() *otg.PatternFlowRSVPPathRecordRouteType1LabelFlags {
	return obj.obj
}

func (obj *patternFlowRSVPPathRecordRouteType1LabelFlags) setMsg(msg *otg.PatternFlowRSVPPathRecordRouteType1LabelFlags) PatternFlowRSVPPathRecordRouteType1LabelFlags {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathRecordRouteType1LabelFlags struct {
	obj *patternFlowRSVPPathRecordRouteType1LabelFlags
}

type marshalPatternFlowRSVPPathRecordRouteType1LabelFlags interface {
	// ToProto marshals PatternFlowRSVPPathRecordRouteType1LabelFlags to protobuf object *otg.PatternFlowRSVPPathRecordRouteType1LabelFlags
	ToProto() (*otg.PatternFlowRSVPPathRecordRouteType1LabelFlags, error)
	// ToPbText marshals PatternFlowRSVPPathRecordRouteType1LabelFlags to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathRecordRouteType1LabelFlags to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathRecordRouteType1LabelFlags to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathRecordRouteType1LabelFlags struct {
	obj *patternFlowRSVPPathRecordRouteType1LabelFlags
}

type unMarshalPatternFlowRSVPPathRecordRouteType1LabelFlags interface {
	// FromProto unmarshals PatternFlowRSVPPathRecordRouteType1LabelFlags from protobuf object *otg.PatternFlowRSVPPathRecordRouteType1LabelFlags
	FromProto(msg *otg.PatternFlowRSVPPathRecordRouteType1LabelFlags) (PatternFlowRSVPPathRecordRouteType1LabelFlags, error)
	// FromPbText unmarshals PatternFlowRSVPPathRecordRouteType1LabelFlags from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathRecordRouteType1LabelFlags from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathRecordRouteType1LabelFlags from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathRecordRouteType1LabelFlags) Marshal() marshalPatternFlowRSVPPathRecordRouteType1LabelFlags {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathRecordRouteType1LabelFlags{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathRecordRouteType1LabelFlags) Unmarshal() unMarshalPatternFlowRSVPPathRecordRouteType1LabelFlags {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathRecordRouteType1LabelFlags{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathRecordRouteType1LabelFlags) ToProto() (*otg.PatternFlowRSVPPathRecordRouteType1LabelFlags, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathRecordRouteType1LabelFlags) FromProto(msg *otg.PatternFlowRSVPPathRecordRouteType1LabelFlags) (PatternFlowRSVPPathRecordRouteType1LabelFlags, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathRecordRouteType1LabelFlags) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRecordRouteType1LabelFlags) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathRecordRouteType1LabelFlags) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRecordRouteType1LabelFlags) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathRecordRouteType1LabelFlags) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRecordRouteType1LabelFlags) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathRecordRouteType1LabelFlags) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathRecordRouteType1LabelFlags) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathRecordRouteType1LabelFlags) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathRecordRouteType1LabelFlags) Clone() (PatternFlowRSVPPathRecordRouteType1LabelFlags, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathRecordRouteType1LabelFlags()
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

// PatternFlowRSVPPathRecordRouteType1LabelFlags is 0x01 = Global label. This flag indicates that the label will be understood if received on any interface.
type PatternFlowRSVPPathRecordRouteType1LabelFlags interface {
	Validation
	// msg marshals PatternFlowRSVPPathRecordRouteType1LabelFlags to protobuf object *otg.PatternFlowRSVPPathRecordRouteType1LabelFlags
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathRecordRouteType1LabelFlags
	// setMsg unmarshals PatternFlowRSVPPathRecordRouteType1LabelFlags from protobuf object *otg.PatternFlowRSVPPathRecordRouteType1LabelFlags
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathRecordRouteType1LabelFlags) PatternFlowRSVPPathRecordRouteType1LabelFlags
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathRecordRouteType1LabelFlags
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathRecordRouteType1LabelFlags
	// validate validates PatternFlowRSVPPathRecordRouteType1LabelFlags
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathRecordRouteType1LabelFlags, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathRecordRouteType1LabelFlagsChoiceEnum, set in PatternFlowRSVPPathRecordRouteType1LabelFlags
	Choice() PatternFlowRSVPPathRecordRouteType1LabelFlagsChoiceEnum
	// setChoice assigns PatternFlowRSVPPathRecordRouteType1LabelFlagsChoiceEnum provided by user to PatternFlowRSVPPathRecordRouteType1LabelFlags
	setChoice(value PatternFlowRSVPPathRecordRouteType1LabelFlagsChoiceEnum) PatternFlowRSVPPathRecordRouteType1LabelFlags
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathRecordRouteType1LabelFlags
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRSVPPathRecordRouteType1LabelFlags.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRSVPPathRecordRouteType1LabelFlags
	SetValue(value uint32) PatternFlowRSVPPathRecordRouteType1LabelFlags
	// HasValue checks if Value has been set in PatternFlowRSVPPathRecordRouteType1LabelFlags
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRSVPPathRecordRouteType1LabelFlags.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRSVPPathRecordRouteType1LabelFlags
	SetValues(value []uint32) PatternFlowRSVPPathRecordRouteType1LabelFlags
}

type PatternFlowRSVPPathRecordRouteType1LabelFlagsChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathRecordRouteType1LabelFlags
var PatternFlowRSVPPathRecordRouteType1LabelFlagsChoice = struct {
	VALUE  PatternFlowRSVPPathRecordRouteType1LabelFlagsChoiceEnum
	VALUES PatternFlowRSVPPathRecordRouteType1LabelFlagsChoiceEnum
}{
	VALUE:  PatternFlowRSVPPathRecordRouteType1LabelFlagsChoiceEnum("value"),
	VALUES: PatternFlowRSVPPathRecordRouteType1LabelFlagsChoiceEnum("values"),
}

func (obj *patternFlowRSVPPathRecordRouteType1LabelFlags) Choice() PatternFlowRSVPPathRecordRouteType1LabelFlagsChoiceEnum {
	return PatternFlowRSVPPathRecordRouteType1LabelFlagsChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathRecordRouteType1LabelFlags) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathRecordRouteType1LabelFlags) setChoice(value PatternFlowRSVPPathRecordRouteType1LabelFlagsChoiceEnum) PatternFlowRSVPPathRecordRouteType1LabelFlags {
	intValue, ok := otg.PatternFlowRSVPPathRecordRouteType1LabelFlags_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathRecordRouteType1LabelFlagsChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathRecordRouteType1LabelFlags_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathRecordRouteType1LabelFlagsChoice.VALUE {
		defaultValue := uint32(1)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathRecordRouteType1LabelFlagsChoice.VALUES {
		defaultValue := []uint32{1}
		obj.obj.Values = defaultValue
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathRecordRouteType1LabelFlags) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathRecordRouteType1LabelFlagsChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathRecordRouteType1LabelFlags) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRSVPPathRecordRouteType1LabelFlags object
func (obj *patternFlowRSVPPathRecordRouteType1LabelFlags) SetValue(value uint32) PatternFlowRSVPPathRecordRouteType1LabelFlags {
	obj.setChoice(PatternFlowRSVPPathRecordRouteType1LabelFlagsChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRSVPPathRecordRouteType1LabelFlags) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{1})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRSVPPathRecordRouteType1LabelFlags object
func (obj *patternFlowRSVPPathRecordRouteType1LabelFlags) SetValues(value []uint32) PatternFlowRSVPPathRecordRouteType1LabelFlags {
	obj.setChoice(PatternFlowRSVPPathRecordRouteType1LabelFlagsChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

func (obj *patternFlowRSVPPathRecordRouteType1LabelFlags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathRecordRouteType1LabelFlags.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowRSVPPathRecordRouteType1LabelFlags.Values <= 255 but Got %d", item))
			}

		}

	}

}

func (obj *patternFlowRSVPPathRecordRouteType1LabelFlags) setDefault() {
	var choices_set int = 0
	var choice PatternFlowRSVPPathRecordRouteType1LabelFlagsChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathRecordRouteType1LabelFlagsChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowRSVPPathRecordRouteType1LabelFlagsChoice.VALUES
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowRSVPPathRecordRouteType1LabelFlagsChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowRSVPPathRecordRouteType1LabelFlags")
			}
		} else {
			intVal := otg.PatternFlowRSVPPathRecordRouteType1LabelFlags_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowRSVPPathRecordRouteType1LabelFlags_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
