package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHPathTraceTlvTimestamp *****
type patternFlowIpv6SRHPathTraceTlvTimestamp struct {
	validation
	obj             *otg.PatternFlowIpv6SRHPathTraceTlvTimestamp
	marshaller      marshalPatternFlowIpv6SRHPathTraceTlvTimestamp
	unMarshaller    unMarshalPatternFlowIpv6SRHPathTraceTlvTimestamp
	incrementHolder PatternFlowIpv6SRHPathTraceTlvTimestampCounter
	decrementHolder PatternFlowIpv6SRHPathTraceTlvTimestampCounter
}

func NewPatternFlowIpv6SRHPathTraceTlvTimestamp() PatternFlowIpv6SRHPathTraceTlvTimestamp {
	obj := patternFlowIpv6SRHPathTraceTlvTimestamp{obj: &otg.PatternFlowIpv6SRHPathTraceTlvTimestamp{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvTimestamp) msg() *otg.PatternFlowIpv6SRHPathTraceTlvTimestamp {
	return obj.obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvTimestamp) setMsg(msg *otg.PatternFlowIpv6SRHPathTraceTlvTimestamp) PatternFlowIpv6SRHPathTraceTlvTimestamp {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHPathTraceTlvTimestamp struct {
	obj *patternFlowIpv6SRHPathTraceTlvTimestamp
}

type marshalPatternFlowIpv6SRHPathTraceTlvTimestamp interface {
	// ToProto marshals PatternFlowIpv6SRHPathTraceTlvTimestamp to protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvTimestamp
	ToProto() (*otg.PatternFlowIpv6SRHPathTraceTlvTimestamp, error)
	// ToPbText marshals PatternFlowIpv6SRHPathTraceTlvTimestamp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHPathTraceTlvTimestamp to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHPathTraceTlvTimestamp to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHPathTraceTlvTimestamp struct {
	obj *patternFlowIpv6SRHPathTraceTlvTimestamp
}

type unMarshalPatternFlowIpv6SRHPathTraceTlvTimestamp interface {
	// FromProto unmarshals PatternFlowIpv6SRHPathTraceTlvTimestamp from protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvTimestamp
	FromProto(msg *otg.PatternFlowIpv6SRHPathTraceTlvTimestamp) (PatternFlowIpv6SRHPathTraceTlvTimestamp, error)
	// FromPbText unmarshals PatternFlowIpv6SRHPathTraceTlvTimestamp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHPathTraceTlvTimestamp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHPathTraceTlvTimestamp from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHPathTraceTlvTimestamp) Marshal() marshalPatternFlowIpv6SRHPathTraceTlvTimestamp {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHPathTraceTlvTimestamp{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHPathTraceTlvTimestamp) Unmarshal() unMarshalPatternFlowIpv6SRHPathTraceTlvTimestamp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHPathTraceTlvTimestamp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHPathTraceTlvTimestamp) ToProto() (*otg.PatternFlowIpv6SRHPathTraceTlvTimestamp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvTimestamp) FromProto(msg *otg.PatternFlowIpv6SRHPathTraceTlvTimestamp) (PatternFlowIpv6SRHPathTraceTlvTimestamp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHPathTraceTlvTimestamp) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvTimestamp) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPathTraceTlvTimestamp) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvTimestamp) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPathTraceTlvTimestamp) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvTimestamp) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHPathTraceTlvTimestamp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPathTraceTlvTimestamp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPathTraceTlvTimestamp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHPathTraceTlvTimestamp) Clone() (PatternFlowIpv6SRHPathTraceTlvTimestamp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHPathTraceTlvTimestamp()
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

func (obj *patternFlowIpv6SRHPathTraceTlvTimestamp) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SRHPathTraceTlvTimestamp is timestamp inserted by the ingress node (draft-ietf-spring-srv6-path-tracing). Encoded as a 32-bit truncated PTP timestamp value.
type PatternFlowIpv6SRHPathTraceTlvTimestamp interface {
	Validation
	// msg marshals PatternFlowIpv6SRHPathTraceTlvTimestamp to protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvTimestamp
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHPathTraceTlvTimestamp
	// setMsg unmarshals PatternFlowIpv6SRHPathTraceTlvTimestamp from protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvTimestamp
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHPathTraceTlvTimestamp) PatternFlowIpv6SRHPathTraceTlvTimestamp
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHPathTraceTlvTimestamp
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHPathTraceTlvTimestamp
	// validate validates PatternFlowIpv6SRHPathTraceTlvTimestamp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHPathTraceTlvTimestamp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SRHPathTraceTlvTimestampChoiceEnum, set in PatternFlowIpv6SRHPathTraceTlvTimestamp
	Choice() PatternFlowIpv6SRHPathTraceTlvTimestampChoiceEnum
	// setChoice assigns PatternFlowIpv6SRHPathTraceTlvTimestampChoiceEnum provided by user to PatternFlowIpv6SRHPathTraceTlvTimestamp
	setChoice(value PatternFlowIpv6SRHPathTraceTlvTimestampChoiceEnum) PatternFlowIpv6SRHPathTraceTlvTimestamp
	// HasChoice checks if Choice has been set in PatternFlowIpv6SRHPathTraceTlvTimestamp
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SRHPathTraceTlvTimestamp.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvTimestamp
	SetValue(value uint32) PatternFlowIpv6SRHPathTraceTlvTimestamp
	// HasValue checks if Value has been set in PatternFlowIpv6SRHPathTraceTlvTimestamp
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SRHPathTraceTlvTimestamp.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvTimestamp
	SetValues(value []uint32) PatternFlowIpv6SRHPathTraceTlvTimestamp
	// Increment returns PatternFlowIpv6SRHPathTraceTlvTimestampCounter, set in PatternFlowIpv6SRHPathTraceTlvTimestamp.
	// PatternFlowIpv6SRHPathTraceTlvTimestampCounter is integer counter pattern
	Increment() PatternFlowIpv6SRHPathTraceTlvTimestampCounter
	// SetIncrement assigns PatternFlowIpv6SRHPathTraceTlvTimestampCounter provided by user to PatternFlowIpv6SRHPathTraceTlvTimestamp.
	// PatternFlowIpv6SRHPathTraceTlvTimestampCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SRHPathTraceTlvTimestampCounter) PatternFlowIpv6SRHPathTraceTlvTimestamp
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SRHPathTraceTlvTimestamp
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SRHPathTraceTlvTimestampCounter, set in PatternFlowIpv6SRHPathTraceTlvTimestamp.
	// PatternFlowIpv6SRHPathTraceTlvTimestampCounter is integer counter pattern
	Decrement() PatternFlowIpv6SRHPathTraceTlvTimestampCounter
	// SetDecrement assigns PatternFlowIpv6SRHPathTraceTlvTimestampCounter provided by user to PatternFlowIpv6SRHPathTraceTlvTimestamp.
	// PatternFlowIpv6SRHPathTraceTlvTimestampCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SRHPathTraceTlvTimestampCounter) PatternFlowIpv6SRHPathTraceTlvTimestamp
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SRHPathTraceTlvTimestamp
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SRHPathTraceTlvTimestampChoiceEnum string

// Enum of Choice on PatternFlowIpv6SRHPathTraceTlvTimestamp
var PatternFlowIpv6SRHPathTraceTlvTimestampChoice = struct {
	VALUE     PatternFlowIpv6SRHPathTraceTlvTimestampChoiceEnum
	VALUES    PatternFlowIpv6SRHPathTraceTlvTimestampChoiceEnum
	INCREMENT PatternFlowIpv6SRHPathTraceTlvTimestampChoiceEnum
	DECREMENT PatternFlowIpv6SRHPathTraceTlvTimestampChoiceEnum
}{
	VALUE:     PatternFlowIpv6SRHPathTraceTlvTimestampChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SRHPathTraceTlvTimestampChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SRHPathTraceTlvTimestampChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SRHPathTraceTlvTimestampChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SRHPathTraceTlvTimestamp) Choice() PatternFlowIpv6SRHPathTraceTlvTimestampChoiceEnum {
	return PatternFlowIpv6SRHPathTraceTlvTimestampChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SRHPathTraceTlvTimestamp) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SRHPathTraceTlvTimestamp) setChoice(value PatternFlowIpv6SRHPathTraceTlvTimestampChoiceEnum) PatternFlowIpv6SRHPathTraceTlvTimestamp {
	intValue, ok := otg.PatternFlowIpv6SRHPathTraceTlvTimestamp_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SRHPathTraceTlvTimestampChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SRHPathTraceTlvTimestamp_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SRHPathTraceTlvTimestampChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SRHPathTraceTlvTimestampChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SRHPathTraceTlvTimestampChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SRHPathTraceTlvTimestampCounter().msg()
	}

	if value == PatternFlowIpv6SRHPathTraceTlvTimestampChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SRHPathTraceTlvTimestampCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvTimestamp) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SRHPathTraceTlvTimestampChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvTimestamp) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SRHPathTraceTlvTimestamp object
func (obj *patternFlowIpv6SRHPathTraceTlvTimestamp) SetValue(value uint32) PatternFlowIpv6SRHPathTraceTlvTimestamp {
	obj.setChoice(PatternFlowIpv6SRHPathTraceTlvTimestampChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SRHPathTraceTlvTimestamp) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SRHPathTraceTlvTimestamp object
func (obj *patternFlowIpv6SRHPathTraceTlvTimestamp) SetValues(value []uint32) PatternFlowIpv6SRHPathTraceTlvTimestamp {
	obj.setChoice(PatternFlowIpv6SRHPathTraceTlvTimestampChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHPathTraceTlvTimestampCounter
func (obj *patternFlowIpv6SRHPathTraceTlvTimestamp) Increment() PatternFlowIpv6SRHPathTraceTlvTimestampCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SRHPathTraceTlvTimestampChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SRHPathTraceTlvTimestampCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHPathTraceTlvTimestampCounter
func (obj *patternFlowIpv6SRHPathTraceTlvTimestamp) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SRHPathTraceTlvTimestampCounter value in the PatternFlowIpv6SRHPathTraceTlvTimestamp object
func (obj *patternFlowIpv6SRHPathTraceTlvTimestamp) SetIncrement(value PatternFlowIpv6SRHPathTraceTlvTimestampCounter) PatternFlowIpv6SRHPathTraceTlvTimestamp {
	obj.setChoice(PatternFlowIpv6SRHPathTraceTlvTimestampChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHPathTraceTlvTimestampCounter
func (obj *patternFlowIpv6SRHPathTraceTlvTimestamp) Decrement() PatternFlowIpv6SRHPathTraceTlvTimestampCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SRHPathTraceTlvTimestampChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SRHPathTraceTlvTimestampCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHPathTraceTlvTimestampCounter
func (obj *patternFlowIpv6SRHPathTraceTlvTimestamp) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SRHPathTraceTlvTimestampCounter value in the PatternFlowIpv6SRHPathTraceTlvTimestamp object
func (obj *patternFlowIpv6SRHPathTraceTlvTimestamp) SetDecrement(value PatternFlowIpv6SRHPathTraceTlvTimestampCounter) PatternFlowIpv6SRHPathTraceTlvTimestamp {
	obj.setChoice(PatternFlowIpv6SRHPathTraceTlvTimestampChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvTimestamp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6SRHPathTraceTlvTimestamp) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SRHPathTraceTlvTimestampChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHPathTraceTlvTimestampChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SRHPathTraceTlvTimestampChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHPathTraceTlvTimestampChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHPathTraceTlvTimestampChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SRHPathTraceTlvTimestampChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SRHPathTraceTlvTimestamp")
			}
		} else {
			intVal := otg.PatternFlowIpv6SRHPathTraceTlvTimestamp_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SRHPathTraceTlvTimestamp_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
