package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingUsidLastEntry *****
type patternFlowIpv6SegmentRoutingUsidLastEntry struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingUsidLastEntry
	marshaller      marshalPatternFlowIpv6SegmentRoutingUsidLastEntry
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingUsidLastEntry
	incrementHolder PatternFlowIpv6SegmentRoutingUsidLastEntryCounter
	decrementHolder PatternFlowIpv6SegmentRoutingUsidLastEntryCounter
}

func NewPatternFlowIpv6SegmentRoutingUsidLastEntry() PatternFlowIpv6SegmentRoutingUsidLastEntry {
	obj := patternFlowIpv6SegmentRoutingUsidLastEntry{obj: &otg.PatternFlowIpv6SegmentRoutingUsidLastEntry{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidLastEntry) msg() *otg.PatternFlowIpv6SegmentRoutingUsidLastEntry {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidLastEntry) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingUsidLastEntry) PatternFlowIpv6SegmentRoutingUsidLastEntry {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingUsidLastEntry struct {
	obj *patternFlowIpv6SegmentRoutingUsidLastEntry
}

type marshalPatternFlowIpv6SegmentRoutingUsidLastEntry interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingUsidLastEntry to protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidLastEntry
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingUsidLastEntry, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingUsidLastEntry to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingUsidLastEntry to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingUsidLastEntry to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingUsidLastEntry struct {
	obj *patternFlowIpv6SegmentRoutingUsidLastEntry
}

type unMarshalPatternFlowIpv6SegmentRoutingUsidLastEntry interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingUsidLastEntry from protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidLastEntry
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingUsidLastEntry) (PatternFlowIpv6SegmentRoutingUsidLastEntry, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingUsidLastEntry from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingUsidLastEntry from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingUsidLastEntry from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingUsidLastEntry) Marshal() marshalPatternFlowIpv6SegmentRoutingUsidLastEntry {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingUsidLastEntry{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingUsidLastEntry) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingUsidLastEntry {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingUsidLastEntry{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingUsidLastEntry) ToProto() (*otg.PatternFlowIpv6SegmentRoutingUsidLastEntry, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidLastEntry) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingUsidLastEntry) (PatternFlowIpv6SegmentRoutingUsidLastEntry, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingUsidLastEntry) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidLastEntry) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingUsidLastEntry) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidLastEntry) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingUsidLastEntry) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidLastEntry) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingUsidLastEntry) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingUsidLastEntry) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingUsidLastEntry) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingUsidLastEntry) Clone() (PatternFlowIpv6SegmentRoutingUsidLastEntry, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingUsidLastEntry()
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

func (obj *patternFlowIpv6SegmentRoutingUsidLastEntry) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingUsidLastEntry is zero-based index of the last element in the uSID container list (RFC 8754 Section 2.1). When auto is assigned the value is set to one less than the number of containers specified.
type PatternFlowIpv6SegmentRoutingUsidLastEntry interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingUsidLastEntry to protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidLastEntry
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingUsidLastEntry
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingUsidLastEntry from protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidLastEntry
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingUsidLastEntry) PatternFlowIpv6SegmentRoutingUsidLastEntry
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingUsidLastEntry
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingUsidLastEntry
	// validate validates PatternFlowIpv6SegmentRoutingUsidLastEntry
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingUsidLastEntry, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingUsidLastEntryChoiceEnum, set in PatternFlowIpv6SegmentRoutingUsidLastEntry
	Choice() PatternFlowIpv6SegmentRoutingUsidLastEntryChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingUsidLastEntryChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingUsidLastEntry
	setChoice(value PatternFlowIpv6SegmentRoutingUsidLastEntryChoiceEnum) PatternFlowIpv6SegmentRoutingUsidLastEntry
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingUsidLastEntry
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SegmentRoutingUsidLastEntry.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingUsidLastEntry
	SetValue(value uint32) PatternFlowIpv6SegmentRoutingUsidLastEntry
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingUsidLastEntry
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SegmentRoutingUsidLastEntry.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SegmentRoutingUsidLastEntry
	SetValues(value []uint32) PatternFlowIpv6SegmentRoutingUsidLastEntry
	// Auto returns uint32, set in PatternFlowIpv6SegmentRoutingUsidLastEntry.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIpv6SegmentRoutingUsidLastEntry
	HasAuto() bool
	// Increment returns PatternFlowIpv6SegmentRoutingUsidLastEntryCounter, set in PatternFlowIpv6SegmentRoutingUsidLastEntry.
	// PatternFlowIpv6SegmentRoutingUsidLastEntryCounter is integer counter pattern
	Increment() PatternFlowIpv6SegmentRoutingUsidLastEntryCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingUsidLastEntryCounter provided by user to PatternFlowIpv6SegmentRoutingUsidLastEntry.
	// PatternFlowIpv6SegmentRoutingUsidLastEntryCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingUsidLastEntryCounter) PatternFlowIpv6SegmentRoutingUsidLastEntry
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingUsidLastEntry
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingUsidLastEntryCounter, set in PatternFlowIpv6SegmentRoutingUsidLastEntry.
	// PatternFlowIpv6SegmentRoutingUsidLastEntryCounter is integer counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingUsidLastEntryCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingUsidLastEntryCounter provided by user to PatternFlowIpv6SegmentRoutingUsidLastEntry.
	// PatternFlowIpv6SegmentRoutingUsidLastEntryCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingUsidLastEntryCounter) PatternFlowIpv6SegmentRoutingUsidLastEntry
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingUsidLastEntry
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingUsidLastEntryChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingUsidLastEntry
var PatternFlowIpv6SegmentRoutingUsidLastEntryChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingUsidLastEntryChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingUsidLastEntryChoiceEnum
	AUTO      PatternFlowIpv6SegmentRoutingUsidLastEntryChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingUsidLastEntryChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingUsidLastEntryChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingUsidLastEntryChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingUsidLastEntryChoiceEnum("values"),
	AUTO:      PatternFlowIpv6SegmentRoutingUsidLastEntryChoiceEnum("auto"),
	INCREMENT: PatternFlowIpv6SegmentRoutingUsidLastEntryChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingUsidLastEntryChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingUsidLastEntry) Choice() PatternFlowIpv6SegmentRoutingUsidLastEntryChoiceEnum {
	return PatternFlowIpv6SegmentRoutingUsidLastEntryChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingUsidLastEntry) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingUsidLastEntry) setChoice(value PatternFlowIpv6SegmentRoutingUsidLastEntryChoiceEnum) PatternFlowIpv6SegmentRoutingUsidLastEntry {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingUsidLastEntry_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingUsidLastEntryChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingUsidLastEntry_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingUsidLastEntryChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingUsidLastEntryChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingUsidLastEntryChoice.AUTO {
		defaultValue := uint32(0)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingUsidLastEntryChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingUsidLastEntryCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingUsidLastEntryChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingUsidLastEntryCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidLastEntry) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingUsidLastEntryChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidLastEntry) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SegmentRoutingUsidLastEntry object
func (obj *patternFlowIpv6SegmentRoutingUsidLastEntry) SetValue(value uint32) PatternFlowIpv6SegmentRoutingUsidLastEntry {
	obj.setChoice(PatternFlowIpv6SegmentRoutingUsidLastEntryChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SegmentRoutingUsidLastEntry) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SegmentRoutingUsidLastEntry object
func (obj *patternFlowIpv6SegmentRoutingUsidLastEntry) SetValues(value []uint32) PatternFlowIpv6SegmentRoutingUsidLastEntry {
	obj.setChoice(PatternFlowIpv6SegmentRoutingUsidLastEntryChoice.VALUES)
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
func (obj *patternFlowIpv6SegmentRoutingUsidLastEntry) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingUsidLastEntryChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidLastEntry) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingUsidLastEntryCounter
func (obj *patternFlowIpv6SegmentRoutingUsidLastEntry) Increment() PatternFlowIpv6SegmentRoutingUsidLastEntryCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingUsidLastEntryChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingUsidLastEntryCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingUsidLastEntryCounter
func (obj *patternFlowIpv6SegmentRoutingUsidLastEntry) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingUsidLastEntryCounter value in the PatternFlowIpv6SegmentRoutingUsidLastEntry object
func (obj *patternFlowIpv6SegmentRoutingUsidLastEntry) SetIncrement(value PatternFlowIpv6SegmentRoutingUsidLastEntryCounter) PatternFlowIpv6SegmentRoutingUsidLastEntry {
	obj.setChoice(PatternFlowIpv6SegmentRoutingUsidLastEntryChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingUsidLastEntryCounter
func (obj *patternFlowIpv6SegmentRoutingUsidLastEntry) Decrement() PatternFlowIpv6SegmentRoutingUsidLastEntryCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingUsidLastEntryChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingUsidLastEntryCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingUsidLastEntryCounter
func (obj *patternFlowIpv6SegmentRoutingUsidLastEntry) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingUsidLastEntryCounter value in the PatternFlowIpv6SegmentRoutingUsidLastEntry object
func (obj *patternFlowIpv6SegmentRoutingUsidLastEntry) SetDecrement(value PatternFlowIpv6SegmentRoutingUsidLastEntryCounter) PatternFlowIpv6SegmentRoutingUsidLastEntry {
	obj.setChoice(PatternFlowIpv6SegmentRoutingUsidLastEntryChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidLastEntry) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingUsidLastEntry.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SegmentRoutingUsidLastEntry.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingUsidLastEntry.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6SegmentRoutingUsidLastEntry) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingUsidLastEntryChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingUsidLastEntryChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingUsidLastEntryChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingUsidLastEntryChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingUsidLastEntryChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingUsidLastEntryChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingUsidLastEntryChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingUsidLastEntry")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingUsidLastEntry_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingUsidLastEntry_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
