package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingLastEntry *****
type patternFlowIpv6SegmentRoutingLastEntry struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingLastEntry
	marshaller      marshalPatternFlowIpv6SegmentRoutingLastEntry
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingLastEntry
	incrementHolder PatternFlowIpv6SegmentRoutingLastEntryCounter
	decrementHolder PatternFlowIpv6SegmentRoutingLastEntryCounter
}

func NewPatternFlowIpv6SegmentRoutingLastEntry() PatternFlowIpv6SegmentRoutingLastEntry {
	obj := patternFlowIpv6SegmentRoutingLastEntry{obj: &otg.PatternFlowIpv6SegmentRoutingLastEntry{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingLastEntry) msg() *otg.PatternFlowIpv6SegmentRoutingLastEntry {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingLastEntry) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingLastEntry) PatternFlowIpv6SegmentRoutingLastEntry {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingLastEntry struct {
	obj *patternFlowIpv6SegmentRoutingLastEntry
}

type marshalPatternFlowIpv6SegmentRoutingLastEntry interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingLastEntry to protobuf object *otg.PatternFlowIpv6SegmentRoutingLastEntry
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingLastEntry, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingLastEntry to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingLastEntry to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingLastEntry to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingLastEntry struct {
	obj *patternFlowIpv6SegmentRoutingLastEntry
}

type unMarshalPatternFlowIpv6SegmentRoutingLastEntry interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingLastEntry from protobuf object *otg.PatternFlowIpv6SegmentRoutingLastEntry
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingLastEntry) (PatternFlowIpv6SegmentRoutingLastEntry, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingLastEntry from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingLastEntry from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingLastEntry from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingLastEntry) Marshal() marshalPatternFlowIpv6SegmentRoutingLastEntry {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingLastEntry{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingLastEntry) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingLastEntry {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingLastEntry{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingLastEntry) ToProto() (*otg.PatternFlowIpv6SegmentRoutingLastEntry, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingLastEntry) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingLastEntry) (PatternFlowIpv6SegmentRoutingLastEntry, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingLastEntry) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingLastEntry) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingLastEntry) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingLastEntry) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingLastEntry) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingLastEntry) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingLastEntry) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingLastEntry) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingLastEntry) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingLastEntry) Clone() (PatternFlowIpv6SegmentRoutingLastEntry, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingLastEntry()
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

func (obj *patternFlowIpv6SegmentRoutingLastEntry) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingLastEntry is 8-bit unsigned integer that contains the zero-based index of the last element in the Segment List array. For example, if the Segment List contains 3 addresses (at indices 0, 1, 2), the value of Last Entry is 2. When auto is assigned the value should be automatically set by the implementation to one less than the number of segments specified.
type PatternFlowIpv6SegmentRoutingLastEntry interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingLastEntry to protobuf object *otg.PatternFlowIpv6SegmentRoutingLastEntry
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingLastEntry
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingLastEntry from protobuf object *otg.PatternFlowIpv6SegmentRoutingLastEntry
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingLastEntry) PatternFlowIpv6SegmentRoutingLastEntry
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingLastEntry
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingLastEntry
	// validate validates PatternFlowIpv6SegmentRoutingLastEntry
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingLastEntry, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingLastEntryChoiceEnum, set in PatternFlowIpv6SegmentRoutingLastEntry
	Choice() PatternFlowIpv6SegmentRoutingLastEntryChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingLastEntryChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingLastEntry
	setChoice(value PatternFlowIpv6SegmentRoutingLastEntryChoiceEnum) PatternFlowIpv6SegmentRoutingLastEntry
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingLastEntry
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SegmentRoutingLastEntry.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingLastEntry
	SetValue(value uint32) PatternFlowIpv6SegmentRoutingLastEntry
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingLastEntry
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SegmentRoutingLastEntry.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SegmentRoutingLastEntry
	SetValues(value []uint32) PatternFlowIpv6SegmentRoutingLastEntry
	// Auto returns uint32, set in PatternFlowIpv6SegmentRoutingLastEntry.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIpv6SegmentRoutingLastEntry
	HasAuto() bool
	// Increment returns PatternFlowIpv6SegmentRoutingLastEntryCounter, set in PatternFlowIpv6SegmentRoutingLastEntry.
	// PatternFlowIpv6SegmentRoutingLastEntryCounter is integer counter pattern
	Increment() PatternFlowIpv6SegmentRoutingLastEntryCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingLastEntryCounter provided by user to PatternFlowIpv6SegmentRoutingLastEntry.
	// PatternFlowIpv6SegmentRoutingLastEntryCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingLastEntryCounter) PatternFlowIpv6SegmentRoutingLastEntry
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingLastEntry
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingLastEntryCounter, set in PatternFlowIpv6SegmentRoutingLastEntry.
	// PatternFlowIpv6SegmentRoutingLastEntryCounter is integer counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingLastEntryCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingLastEntryCounter provided by user to PatternFlowIpv6SegmentRoutingLastEntry.
	// PatternFlowIpv6SegmentRoutingLastEntryCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingLastEntryCounter) PatternFlowIpv6SegmentRoutingLastEntry
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingLastEntry
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingLastEntryChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingLastEntry
var PatternFlowIpv6SegmentRoutingLastEntryChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingLastEntryChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingLastEntryChoiceEnum
	AUTO      PatternFlowIpv6SegmentRoutingLastEntryChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingLastEntryChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingLastEntryChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingLastEntryChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingLastEntryChoiceEnum("values"),
	AUTO:      PatternFlowIpv6SegmentRoutingLastEntryChoiceEnum("auto"),
	INCREMENT: PatternFlowIpv6SegmentRoutingLastEntryChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingLastEntryChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingLastEntry) Choice() PatternFlowIpv6SegmentRoutingLastEntryChoiceEnum {
	return PatternFlowIpv6SegmentRoutingLastEntryChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingLastEntry) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingLastEntry) setChoice(value PatternFlowIpv6SegmentRoutingLastEntryChoiceEnum) PatternFlowIpv6SegmentRoutingLastEntry {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingLastEntry_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingLastEntryChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingLastEntry_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingLastEntryChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingLastEntryChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingLastEntryChoice.AUTO {
		defaultValue := uint32(0)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingLastEntryChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingLastEntryCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingLastEntryChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingLastEntryCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingLastEntry) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingLastEntryChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingLastEntry) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SegmentRoutingLastEntry object
func (obj *patternFlowIpv6SegmentRoutingLastEntry) SetValue(value uint32) PatternFlowIpv6SegmentRoutingLastEntry {
	obj.setChoice(PatternFlowIpv6SegmentRoutingLastEntryChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SegmentRoutingLastEntry) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SegmentRoutingLastEntry object
func (obj *patternFlowIpv6SegmentRoutingLastEntry) SetValues(value []uint32) PatternFlowIpv6SegmentRoutingLastEntry {
	obj.setChoice(PatternFlowIpv6SegmentRoutingLastEntryChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIpv6SegmentRoutingLastEntry) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingLastEntryChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIpv6SegmentRoutingLastEntry) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingLastEntryCounter
func (obj *patternFlowIpv6SegmentRoutingLastEntry) Increment() PatternFlowIpv6SegmentRoutingLastEntryCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingLastEntryChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingLastEntryCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingLastEntryCounter
func (obj *patternFlowIpv6SegmentRoutingLastEntry) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingLastEntryCounter value in the PatternFlowIpv6SegmentRoutingLastEntry object
func (obj *patternFlowIpv6SegmentRoutingLastEntry) SetIncrement(value PatternFlowIpv6SegmentRoutingLastEntryCounter) PatternFlowIpv6SegmentRoutingLastEntry {
	obj.setChoice(PatternFlowIpv6SegmentRoutingLastEntryChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingLastEntryCounter
func (obj *patternFlowIpv6SegmentRoutingLastEntry) Decrement() PatternFlowIpv6SegmentRoutingLastEntryCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingLastEntryChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingLastEntryCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingLastEntryCounter
func (obj *patternFlowIpv6SegmentRoutingLastEntry) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingLastEntryCounter value in the PatternFlowIpv6SegmentRoutingLastEntry object
func (obj *patternFlowIpv6SegmentRoutingLastEntry) SetDecrement(value PatternFlowIpv6SegmentRoutingLastEntryCounter) PatternFlowIpv6SegmentRoutingLastEntry {
	obj.setChoice(PatternFlowIpv6SegmentRoutingLastEntryChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingLastEntry) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingLastEntry.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SegmentRoutingLastEntry.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingLastEntry.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6SegmentRoutingLastEntry) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingLastEntryChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingLastEntryChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingLastEntryChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingLastEntryChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingLastEntryChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingLastEntryChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingLastEntryChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingLastEntry")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingLastEntry_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingLastEntry_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
