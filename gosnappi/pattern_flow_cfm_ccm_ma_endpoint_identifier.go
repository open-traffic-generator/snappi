package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmCcmMaEndpointIdentifier *****
type patternFlowCfmCcmMaEndpointIdentifier struct {
	validation
	obj             *otg.PatternFlowCfmCcmMaEndpointIdentifier
	marshaller      marshalPatternFlowCfmCcmMaEndpointIdentifier
	unMarshaller    unMarshalPatternFlowCfmCcmMaEndpointIdentifier
	incrementHolder PatternFlowCfmCcmMaEndpointIdentifierCounter
	decrementHolder PatternFlowCfmCcmMaEndpointIdentifierCounter
}

func NewPatternFlowCfmCcmMaEndpointIdentifier() PatternFlowCfmCcmMaEndpointIdentifier {
	obj := patternFlowCfmCcmMaEndpointIdentifier{obj: &otg.PatternFlowCfmCcmMaEndpointIdentifier{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmCcmMaEndpointIdentifier) msg() *otg.PatternFlowCfmCcmMaEndpointIdentifier {
	return obj.obj
}

func (obj *patternFlowCfmCcmMaEndpointIdentifier) setMsg(msg *otg.PatternFlowCfmCcmMaEndpointIdentifier) PatternFlowCfmCcmMaEndpointIdentifier {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmCcmMaEndpointIdentifier struct {
	obj *patternFlowCfmCcmMaEndpointIdentifier
}

type marshalPatternFlowCfmCcmMaEndpointIdentifier interface {
	// ToProto marshals PatternFlowCfmCcmMaEndpointIdentifier to protobuf object *otg.PatternFlowCfmCcmMaEndpointIdentifier
	ToProto() (*otg.PatternFlowCfmCcmMaEndpointIdentifier, error)
	// ToPbText marshals PatternFlowCfmCcmMaEndpointIdentifier to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmCcmMaEndpointIdentifier to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmCcmMaEndpointIdentifier to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmCcmMaEndpointIdentifier struct {
	obj *patternFlowCfmCcmMaEndpointIdentifier
}

type unMarshalPatternFlowCfmCcmMaEndpointIdentifier interface {
	// FromProto unmarshals PatternFlowCfmCcmMaEndpointIdentifier from protobuf object *otg.PatternFlowCfmCcmMaEndpointIdentifier
	FromProto(msg *otg.PatternFlowCfmCcmMaEndpointIdentifier) (PatternFlowCfmCcmMaEndpointIdentifier, error)
	// FromPbText unmarshals PatternFlowCfmCcmMaEndpointIdentifier from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmCcmMaEndpointIdentifier from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmCcmMaEndpointIdentifier from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmCcmMaEndpointIdentifier) Marshal() marshalPatternFlowCfmCcmMaEndpointIdentifier {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmCcmMaEndpointIdentifier{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmCcmMaEndpointIdentifier) Unmarshal() unMarshalPatternFlowCfmCcmMaEndpointIdentifier {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmCcmMaEndpointIdentifier{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmCcmMaEndpointIdentifier) ToProto() (*otg.PatternFlowCfmCcmMaEndpointIdentifier, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmCcmMaEndpointIdentifier) FromProto(msg *otg.PatternFlowCfmCcmMaEndpointIdentifier) (PatternFlowCfmCcmMaEndpointIdentifier, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmCcmMaEndpointIdentifier) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmCcmMaEndpointIdentifier) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmCcmMaEndpointIdentifier) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmCcmMaEndpointIdentifier) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmCcmMaEndpointIdentifier) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmCcmMaEndpointIdentifier) FromJson(value string) error {
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

func (obj *patternFlowCfmCcmMaEndpointIdentifier) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmCcmMaEndpointIdentifier) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmCcmMaEndpointIdentifier) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmCcmMaEndpointIdentifier) Clone() (PatternFlowCfmCcmMaEndpointIdentifier, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmCcmMaEndpointIdentifier()
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

func (obj *patternFlowCfmCcmMaEndpointIdentifier) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowCfmCcmMaEndpointIdentifier is maintenance association endPoint identifier
type PatternFlowCfmCcmMaEndpointIdentifier interface {
	Validation
	// msg marshals PatternFlowCfmCcmMaEndpointIdentifier to protobuf object *otg.PatternFlowCfmCcmMaEndpointIdentifier
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmCcmMaEndpointIdentifier
	// setMsg unmarshals PatternFlowCfmCcmMaEndpointIdentifier from protobuf object *otg.PatternFlowCfmCcmMaEndpointIdentifier
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmCcmMaEndpointIdentifier) PatternFlowCfmCcmMaEndpointIdentifier
	// provides marshal interface
	Marshal() marshalPatternFlowCfmCcmMaEndpointIdentifier
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmCcmMaEndpointIdentifier
	// validate validates PatternFlowCfmCcmMaEndpointIdentifier
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmCcmMaEndpointIdentifier, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowCfmCcmMaEndpointIdentifierChoiceEnum, set in PatternFlowCfmCcmMaEndpointIdentifier
	Choice() PatternFlowCfmCcmMaEndpointIdentifierChoiceEnum
	// setChoice assigns PatternFlowCfmCcmMaEndpointIdentifierChoiceEnum provided by user to PatternFlowCfmCcmMaEndpointIdentifier
	setChoice(value PatternFlowCfmCcmMaEndpointIdentifierChoiceEnum) PatternFlowCfmCcmMaEndpointIdentifier
	// HasChoice checks if Choice has been set in PatternFlowCfmCcmMaEndpointIdentifier
	HasChoice() bool
	// Value returns uint32, set in PatternFlowCfmCcmMaEndpointIdentifier.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowCfmCcmMaEndpointIdentifier
	SetValue(value uint32) PatternFlowCfmCcmMaEndpointIdentifier
	// HasValue checks if Value has been set in PatternFlowCfmCcmMaEndpointIdentifier
	HasValue() bool
	// Values returns []uint32, set in PatternFlowCfmCcmMaEndpointIdentifier.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowCfmCcmMaEndpointIdentifier
	SetValues(value []uint32) PatternFlowCfmCcmMaEndpointIdentifier
	// Increment returns PatternFlowCfmCcmMaEndpointIdentifierCounter, set in PatternFlowCfmCcmMaEndpointIdentifier.
	// PatternFlowCfmCcmMaEndpointIdentifierCounter is integer counter pattern
	Increment() PatternFlowCfmCcmMaEndpointIdentifierCounter
	// SetIncrement assigns PatternFlowCfmCcmMaEndpointIdentifierCounter provided by user to PatternFlowCfmCcmMaEndpointIdentifier.
	// PatternFlowCfmCcmMaEndpointIdentifierCounter is integer counter pattern
	SetIncrement(value PatternFlowCfmCcmMaEndpointIdentifierCounter) PatternFlowCfmCcmMaEndpointIdentifier
	// HasIncrement checks if Increment has been set in PatternFlowCfmCcmMaEndpointIdentifier
	HasIncrement() bool
	// Decrement returns PatternFlowCfmCcmMaEndpointIdentifierCounter, set in PatternFlowCfmCcmMaEndpointIdentifier.
	// PatternFlowCfmCcmMaEndpointIdentifierCounter is integer counter pattern
	Decrement() PatternFlowCfmCcmMaEndpointIdentifierCounter
	// SetDecrement assigns PatternFlowCfmCcmMaEndpointIdentifierCounter provided by user to PatternFlowCfmCcmMaEndpointIdentifier.
	// PatternFlowCfmCcmMaEndpointIdentifierCounter is integer counter pattern
	SetDecrement(value PatternFlowCfmCcmMaEndpointIdentifierCounter) PatternFlowCfmCcmMaEndpointIdentifier
	// HasDecrement checks if Decrement has been set in PatternFlowCfmCcmMaEndpointIdentifier
	HasDecrement() bool
	setNil()
}

type PatternFlowCfmCcmMaEndpointIdentifierChoiceEnum string

// Enum of Choice on PatternFlowCfmCcmMaEndpointIdentifier
var PatternFlowCfmCcmMaEndpointIdentifierChoice = struct {
	VALUE     PatternFlowCfmCcmMaEndpointIdentifierChoiceEnum
	VALUES    PatternFlowCfmCcmMaEndpointIdentifierChoiceEnum
	INCREMENT PatternFlowCfmCcmMaEndpointIdentifierChoiceEnum
	DECREMENT PatternFlowCfmCcmMaEndpointIdentifierChoiceEnum
}{
	VALUE:     PatternFlowCfmCcmMaEndpointIdentifierChoiceEnum("value"),
	VALUES:    PatternFlowCfmCcmMaEndpointIdentifierChoiceEnum("values"),
	INCREMENT: PatternFlowCfmCcmMaEndpointIdentifierChoiceEnum("increment"),
	DECREMENT: PatternFlowCfmCcmMaEndpointIdentifierChoiceEnum("decrement"),
}

func (obj *patternFlowCfmCcmMaEndpointIdentifier) Choice() PatternFlowCfmCcmMaEndpointIdentifierChoiceEnum {
	return PatternFlowCfmCcmMaEndpointIdentifierChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowCfmCcmMaEndpointIdentifier) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowCfmCcmMaEndpointIdentifier) setChoice(value PatternFlowCfmCcmMaEndpointIdentifierChoiceEnum) PatternFlowCfmCcmMaEndpointIdentifier {
	intValue, ok := otg.PatternFlowCfmCcmMaEndpointIdentifier_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowCfmCcmMaEndpointIdentifierChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowCfmCcmMaEndpointIdentifier_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowCfmCcmMaEndpointIdentifierChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowCfmCcmMaEndpointIdentifierChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowCfmCcmMaEndpointIdentifierChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowCfmCcmMaEndpointIdentifierCounter().msg()
	}

	if value == PatternFlowCfmCcmMaEndpointIdentifierChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowCfmCcmMaEndpointIdentifierCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowCfmCcmMaEndpointIdentifier) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowCfmCcmMaEndpointIdentifierChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowCfmCcmMaEndpointIdentifier) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowCfmCcmMaEndpointIdentifier object
func (obj *patternFlowCfmCcmMaEndpointIdentifier) SetValue(value uint32) PatternFlowCfmCcmMaEndpointIdentifier {
	obj.setChoice(PatternFlowCfmCcmMaEndpointIdentifierChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowCfmCcmMaEndpointIdentifier) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowCfmCcmMaEndpointIdentifier object
func (obj *patternFlowCfmCcmMaEndpointIdentifier) SetValues(value []uint32) PatternFlowCfmCcmMaEndpointIdentifier {
	obj.setChoice(PatternFlowCfmCcmMaEndpointIdentifierChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowCfmCcmMaEndpointIdentifierCounter
func (obj *patternFlowCfmCcmMaEndpointIdentifier) Increment() PatternFlowCfmCcmMaEndpointIdentifierCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowCfmCcmMaEndpointIdentifierChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowCfmCcmMaEndpointIdentifierCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowCfmCcmMaEndpointIdentifierCounter
func (obj *patternFlowCfmCcmMaEndpointIdentifier) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowCfmCcmMaEndpointIdentifierCounter value in the PatternFlowCfmCcmMaEndpointIdentifier object
func (obj *patternFlowCfmCcmMaEndpointIdentifier) SetIncrement(value PatternFlowCfmCcmMaEndpointIdentifierCounter) PatternFlowCfmCcmMaEndpointIdentifier {
	obj.setChoice(PatternFlowCfmCcmMaEndpointIdentifierChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowCfmCcmMaEndpointIdentifierCounter
func (obj *patternFlowCfmCcmMaEndpointIdentifier) Decrement() PatternFlowCfmCcmMaEndpointIdentifierCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowCfmCcmMaEndpointIdentifierChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowCfmCcmMaEndpointIdentifierCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowCfmCcmMaEndpointIdentifierCounter
func (obj *patternFlowCfmCcmMaEndpointIdentifier) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowCfmCcmMaEndpointIdentifierCounter value in the PatternFlowCfmCcmMaEndpointIdentifier object
func (obj *patternFlowCfmCcmMaEndpointIdentifier) SetDecrement(value PatternFlowCfmCcmMaEndpointIdentifierCounter) PatternFlowCfmCcmMaEndpointIdentifier {
	obj.setChoice(PatternFlowCfmCcmMaEndpointIdentifierChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowCfmCcmMaEndpointIdentifier) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmCcmMaEndpointIdentifier.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowCfmCcmMaEndpointIdentifier.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowCfmCcmMaEndpointIdentifier) setDefault() {
	var choices_set int = 0
	var choice PatternFlowCfmCcmMaEndpointIdentifierChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowCfmCcmMaEndpointIdentifierChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowCfmCcmMaEndpointIdentifierChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowCfmCcmMaEndpointIdentifierChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowCfmCcmMaEndpointIdentifierChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowCfmCcmMaEndpointIdentifierChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowCfmCcmMaEndpointIdentifier")
			}
		} else {
			intVal := otg.PatternFlowCfmCcmMaEndpointIdentifier_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowCfmCcmMaEndpointIdentifier_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
