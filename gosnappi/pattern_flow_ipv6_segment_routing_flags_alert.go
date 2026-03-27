package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingFlagsAlert *****
type patternFlowIpv6SegmentRoutingFlagsAlert struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingFlagsAlert
	marshaller      marshalPatternFlowIpv6SegmentRoutingFlagsAlert
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingFlagsAlert
	incrementHolder PatternFlowIpv6SegmentRoutingFlagsAlertCounter
	decrementHolder PatternFlowIpv6SegmentRoutingFlagsAlertCounter
}

func NewPatternFlowIpv6SegmentRoutingFlagsAlert() PatternFlowIpv6SegmentRoutingFlagsAlert {
	obj := patternFlowIpv6SegmentRoutingFlagsAlert{obj: &otg.PatternFlowIpv6SegmentRoutingFlagsAlert{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsAlert) msg() *otg.PatternFlowIpv6SegmentRoutingFlagsAlert {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsAlert) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingFlagsAlert) PatternFlowIpv6SegmentRoutingFlagsAlert {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingFlagsAlert struct {
	obj *patternFlowIpv6SegmentRoutingFlagsAlert
}

type marshalPatternFlowIpv6SegmentRoutingFlagsAlert interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingFlagsAlert to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsAlert
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsAlert, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingFlagsAlert to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingFlagsAlert to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingFlagsAlert to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingFlagsAlert struct {
	obj *patternFlowIpv6SegmentRoutingFlagsAlert
}

type unMarshalPatternFlowIpv6SegmentRoutingFlagsAlert interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingFlagsAlert from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsAlert
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsAlert) (PatternFlowIpv6SegmentRoutingFlagsAlert, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingFlagsAlert from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingFlagsAlert from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingFlagsAlert from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingFlagsAlert) Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsAlert {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingFlagsAlert{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingFlagsAlert) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsAlert {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingFlagsAlert{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsAlert) ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsAlert, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsAlert) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsAlert) (PatternFlowIpv6SegmentRoutingFlagsAlert, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsAlert) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsAlert) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsAlert) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsAlert) FromYaml(value string) error {
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
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsAlert) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsAlert) FromJson(value string) error {
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
	m.obj.setNil()
	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *patternFlowIpv6SegmentRoutingFlagsAlert) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsAlert) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsAlert) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingFlagsAlert) Clone() (PatternFlowIpv6SegmentRoutingFlagsAlert, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingFlagsAlert()
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

func (obj *patternFlowIpv6SegmentRoutingFlagsAlert) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingFlagsAlert is alert Flag. Indicates the presence of important TLVs that must be inspected.
type PatternFlowIpv6SegmentRoutingFlagsAlert interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingFlagsAlert to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsAlert
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingFlagsAlert
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingFlagsAlert from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsAlert
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingFlagsAlert) PatternFlowIpv6SegmentRoutingFlagsAlert
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsAlert
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsAlert
	// validate validates PatternFlowIpv6SegmentRoutingFlagsAlert
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingFlagsAlert, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingFlagsAlertChoiceEnum, set in PatternFlowIpv6SegmentRoutingFlagsAlert
	Choice() PatternFlowIpv6SegmentRoutingFlagsAlertChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingFlagsAlertChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingFlagsAlert
	setChoice(value PatternFlowIpv6SegmentRoutingFlagsAlertChoiceEnum) PatternFlowIpv6SegmentRoutingFlagsAlert
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingFlagsAlert
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsAlert.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsAlert
	SetValue(value uint32) PatternFlowIpv6SegmentRoutingFlagsAlert
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingFlagsAlert
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SegmentRoutingFlagsAlert.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsAlert
	SetValues(value []uint32) PatternFlowIpv6SegmentRoutingFlagsAlert
	// Increment returns PatternFlowIpv6SegmentRoutingFlagsAlertCounter, set in PatternFlowIpv6SegmentRoutingFlagsAlert.
	// PatternFlowIpv6SegmentRoutingFlagsAlertCounter is integer counter pattern
	Increment() PatternFlowIpv6SegmentRoutingFlagsAlertCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingFlagsAlertCounter provided by user to PatternFlowIpv6SegmentRoutingFlagsAlert.
	// PatternFlowIpv6SegmentRoutingFlagsAlertCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingFlagsAlertCounter) PatternFlowIpv6SegmentRoutingFlagsAlert
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingFlagsAlert
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingFlagsAlertCounter, set in PatternFlowIpv6SegmentRoutingFlagsAlert.
	// PatternFlowIpv6SegmentRoutingFlagsAlertCounter is integer counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingFlagsAlertCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingFlagsAlertCounter provided by user to PatternFlowIpv6SegmentRoutingFlagsAlert.
	// PatternFlowIpv6SegmentRoutingFlagsAlertCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingFlagsAlertCounter) PatternFlowIpv6SegmentRoutingFlagsAlert
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingFlagsAlert
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingFlagsAlertChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingFlagsAlert
var PatternFlowIpv6SegmentRoutingFlagsAlertChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingFlagsAlertChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingFlagsAlertChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingFlagsAlertChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingFlagsAlertChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingFlagsAlertChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingFlagsAlertChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SegmentRoutingFlagsAlertChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingFlagsAlertChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingFlagsAlert) Choice() PatternFlowIpv6SegmentRoutingFlagsAlertChoiceEnum {
	return PatternFlowIpv6SegmentRoutingFlagsAlertChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingFlagsAlert) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingFlagsAlert) setChoice(value PatternFlowIpv6SegmentRoutingFlagsAlertChoiceEnum) PatternFlowIpv6SegmentRoutingFlagsAlert {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingFlagsAlert_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingFlagsAlertChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingFlagsAlert_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingFlagsAlertChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingFlagsAlertChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingFlagsAlertChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingFlagsAlertCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingFlagsAlertChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingFlagsAlertCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsAlert) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsAlertChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsAlert) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsAlert object
func (obj *patternFlowIpv6SegmentRoutingFlagsAlert) SetValue(value uint32) PatternFlowIpv6SegmentRoutingFlagsAlert {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsAlertChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsAlert) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SegmentRoutingFlagsAlert object
func (obj *patternFlowIpv6SegmentRoutingFlagsAlert) SetValues(value []uint32) PatternFlowIpv6SegmentRoutingFlagsAlert {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsAlertChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingFlagsAlertCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsAlert) Increment() PatternFlowIpv6SegmentRoutingFlagsAlertCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsAlertChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingFlagsAlertCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingFlagsAlertCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsAlert) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingFlagsAlertCounter value in the PatternFlowIpv6SegmentRoutingFlagsAlert object
func (obj *patternFlowIpv6SegmentRoutingFlagsAlert) SetIncrement(value PatternFlowIpv6SegmentRoutingFlagsAlertCounter) PatternFlowIpv6SegmentRoutingFlagsAlert {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsAlertChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingFlagsAlertCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsAlert) Decrement() PatternFlowIpv6SegmentRoutingFlagsAlertCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsAlertChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingFlagsAlertCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingFlagsAlertCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsAlert) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingFlagsAlertCounter value in the PatternFlowIpv6SegmentRoutingFlagsAlert object
func (obj *patternFlowIpv6SegmentRoutingFlagsAlert) SetDecrement(value PatternFlowIpv6SegmentRoutingFlagsAlertCounter) PatternFlowIpv6SegmentRoutingFlagsAlert {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsAlertChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsAlert) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsAlert.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SegmentRoutingFlagsAlert.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6SegmentRoutingFlagsAlert) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingFlagsAlertChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsAlertChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsAlertChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsAlertChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsAlertChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsAlertChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingFlagsAlert")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingFlagsAlert_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingFlagsAlert_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
