package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowVlanPriority *****
type patternFlowVlanPriority struct {
	validation
	obj              *otg.PatternFlowVlanPriority
	marshaller       marshalPatternFlowVlanPriority
	unMarshaller     unMarshalPatternFlowVlanPriority
	incrementHolder  PatternFlowVlanPriorityCounter
	decrementHolder  PatternFlowVlanPriorityCounter
	metricTagsHolder PatternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter
}

func NewPatternFlowVlanPriority() PatternFlowVlanPriority {
	obj := patternFlowVlanPriority{obj: &otg.PatternFlowVlanPriority{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowVlanPriority) msg() *otg.PatternFlowVlanPriority {
	return obj.obj
}

func (obj *patternFlowVlanPriority) setMsg(msg *otg.PatternFlowVlanPriority) PatternFlowVlanPriority {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowVlanPriority struct {
	obj *patternFlowVlanPriority
}

type marshalPatternFlowVlanPriority interface {
	// ToProto marshals PatternFlowVlanPriority to protobuf object *otg.PatternFlowVlanPriority
	ToProto() (*otg.PatternFlowVlanPriority, error)
	// ToPbText marshals PatternFlowVlanPriority to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowVlanPriority to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowVlanPriority to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowVlanPriority struct {
	obj *patternFlowVlanPriority
}

type unMarshalPatternFlowVlanPriority interface {
	// FromProto unmarshals PatternFlowVlanPriority from protobuf object *otg.PatternFlowVlanPriority
	FromProto(msg *otg.PatternFlowVlanPriority) (PatternFlowVlanPriority, error)
	// FromPbText unmarshals PatternFlowVlanPriority from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowVlanPriority from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowVlanPriority from JSON text
	FromJson(value string) error
}

func (obj *patternFlowVlanPriority) Marshal() marshalPatternFlowVlanPriority {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowVlanPriority{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowVlanPriority) Unmarshal() unMarshalPatternFlowVlanPriority {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowVlanPriority{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowVlanPriority) ToProto() (*otg.PatternFlowVlanPriority, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowVlanPriority) FromProto(msg *otg.PatternFlowVlanPriority) (PatternFlowVlanPriority, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowVlanPriority) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowVlanPriority) FromPbText(value string) error {
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

func (m *marshalpatternFlowVlanPriority) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowVlanPriority) FromYaml(value string) error {
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

func (m *marshalpatternFlowVlanPriority) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowVlanPriority) FromJson(value string) error {
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

func (obj *patternFlowVlanPriority) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowVlanPriority) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowVlanPriority) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowVlanPriority) Clone() (PatternFlowVlanPriority, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowVlanPriority()
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

func (obj *patternFlowVlanPriority) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowVlanPriority is priority code point
type PatternFlowVlanPriority interface {
	Validation
	// msg marshals PatternFlowVlanPriority to protobuf object *otg.PatternFlowVlanPriority
	// and doesn't set defaults
	msg() *otg.PatternFlowVlanPriority
	// setMsg unmarshals PatternFlowVlanPriority from protobuf object *otg.PatternFlowVlanPriority
	// and doesn't set defaults
	setMsg(*otg.PatternFlowVlanPriority) PatternFlowVlanPriority
	// provides marshal interface
	Marshal() marshalPatternFlowVlanPriority
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowVlanPriority
	// validate validates PatternFlowVlanPriority
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowVlanPriority, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowVlanPriorityChoiceEnum, set in PatternFlowVlanPriority
	Choice() PatternFlowVlanPriorityChoiceEnum
	// setChoice assigns PatternFlowVlanPriorityChoiceEnum provided by user to PatternFlowVlanPriority
	setChoice(value PatternFlowVlanPriorityChoiceEnum) PatternFlowVlanPriority
	// HasChoice checks if Choice has been set in PatternFlowVlanPriority
	HasChoice() bool
	// Value returns uint32, set in PatternFlowVlanPriority.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowVlanPriority
	SetValue(value uint32) PatternFlowVlanPriority
	// HasValue checks if Value has been set in PatternFlowVlanPriority
	HasValue() bool
	// Values returns []uint32, set in PatternFlowVlanPriority.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowVlanPriority
	SetValues(value []uint32) PatternFlowVlanPriority
	// Increment returns PatternFlowVlanPriorityCounter, set in PatternFlowVlanPriority.
	// PatternFlowVlanPriorityCounter is integer counter pattern
	Increment() PatternFlowVlanPriorityCounter
	// SetIncrement assigns PatternFlowVlanPriorityCounter provided by user to PatternFlowVlanPriority.
	// PatternFlowVlanPriorityCounter is integer counter pattern
	SetIncrement(value PatternFlowVlanPriorityCounter) PatternFlowVlanPriority
	// HasIncrement checks if Increment has been set in PatternFlowVlanPriority
	HasIncrement() bool
	// Decrement returns PatternFlowVlanPriorityCounter, set in PatternFlowVlanPriority.
	// PatternFlowVlanPriorityCounter is integer counter pattern
	Decrement() PatternFlowVlanPriorityCounter
	// SetDecrement assigns PatternFlowVlanPriorityCounter provided by user to PatternFlowVlanPriority.
	// PatternFlowVlanPriorityCounter is integer counter pattern
	SetDecrement(value PatternFlowVlanPriorityCounter) PatternFlowVlanPriority
	// HasDecrement checks if Decrement has been set in PatternFlowVlanPriority
	HasDecrement() bool
	// MetricTags returns PatternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIterIter, set in PatternFlowVlanPriority
	MetricTags() PatternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter
	setNil()
}

type PatternFlowVlanPriorityChoiceEnum string

// Enum of Choice on PatternFlowVlanPriority
var PatternFlowVlanPriorityChoice = struct {
	VALUE     PatternFlowVlanPriorityChoiceEnum
	VALUES    PatternFlowVlanPriorityChoiceEnum
	INCREMENT PatternFlowVlanPriorityChoiceEnum
	DECREMENT PatternFlowVlanPriorityChoiceEnum
}{
	VALUE:     PatternFlowVlanPriorityChoiceEnum("value"),
	VALUES:    PatternFlowVlanPriorityChoiceEnum("values"),
	INCREMENT: PatternFlowVlanPriorityChoiceEnum("increment"),
	DECREMENT: PatternFlowVlanPriorityChoiceEnum("decrement"),
}

func (obj *patternFlowVlanPriority) Choice() PatternFlowVlanPriorityChoiceEnum {
	return PatternFlowVlanPriorityChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowVlanPriority) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowVlanPriority) setChoice(value PatternFlowVlanPriorityChoiceEnum) PatternFlowVlanPriority {
	intValue, ok := otg.PatternFlowVlanPriority_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowVlanPriorityChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowVlanPriority_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowVlanPriorityChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowVlanPriorityChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowVlanPriorityChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowVlanPriorityCounter().msg()
	}

	if value == PatternFlowVlanPriorityChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowVlanPriorityCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowVlanPriority) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowVlanPriorityChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowVlanPriority) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowVlanPriority object
func (obj *patternFlowVlanPriority) SetValue(value uint32) PatternFlowVlanPriority {
	obj.setChoice(PatternFlowVlanPriorityChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowVlanPriority) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowVlanPriority object
func (obj *patternFlowVlanPriority) SetValues(value []uint32) PatternFlowVlanPriority {
	obj.setChoice(PatternFlowVlanPriorityChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowVlanPriorityCounter
func (obj *patternFlowVlanPriority) Increment() PatternFlowVlanPriorityCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowVlanPriorityChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowVlanPriorityCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowVlanPriorityCounter
func (obj *patternFlowVlanPriority) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowVlanPriorityCounter value in the PatternFlowVlanPriority object
func (obj *patternFlowVlanPriority) SetIncrement(value PatternFlowVlanPriorityCounter) PatternFlowVlanPriority {
	obj.setChoice(PatternFlowVlanPriorityChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowVlanPriorityCounter
func (obj *patternFlowVlanPriority) Decrement() PatternFlowVlanPriorityCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowVlanPriorityChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowVlanPriorityCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowVlanPriorityCounter
func (obj *patternFlowVlanPriority) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowVlanPriorityCounter value in the PatternFlowVlanPriority object
func (obj *patternFlowVlanPriority) SetDecrement(value PatternFlowVlanPriorityCounter) PatternFlowVlanPriority {
	obj.setChoice(PatternFlowVlanPriorityChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowVlanPriorityMetricTag
func (obj *patternFlowVlanPriority) MetricTags() PatternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowVlanPriorityMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter struct {
	obj                                   *patternFlowVlanPriority
	patternFlowVlanPriorityMetricTagSlice []PatternFlowVlanPriorityMetricTag
	fieldPtr                              *[]*otg.PatternFlowVlanPriorityMetricTag
}

func newPatternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter(ptr *[]*otg.PatternFlowVlanPriorityMetricTag) PatternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter {
	return &patternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter{fieldPtr: ptr}
}

type PatternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter interface {
	setMsg(*patternFlowVlanPriority) PatternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter
	Items() []PatternFlowVlanPriorityMetricTag
	Add() PatternFlowVlanPriorityMetricTag
	Append(items ...PatternFlowVlanPriorityMetricTag) PatternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter
	Set(index int, newObj PatternFlowVlanPriorityMetricTag) PatternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter
	Clear() PatternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter
	clearHolderSlice() PatternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter
	appendHolderSlice(item PatternFlowVlanPriorityMetricTag) PatternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter
}

func (obj *patternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter) setMsg(msg *patternFlowVlanPriority) PatternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowVlanPriorityMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter) Items() []PatternFlowVlanPriorityMetricTag {
	return obj.patternFlowVlanPriorityMetricTagSlice
}

func (obj *patternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter) Add() PatternFlowVlanPriorityMetricTag {
	newObj := &otg.PatternFlowVlanPriorityMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowVlanPriorityMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowVlanPriorityMetricTagSlice = append(obj.patternFlowVlanPriorityMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter) Append(items ...PatternFlowVlanPriorityMetricTag) PatternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowVlanPriorityMetricTagSlice = append(obj.patternFlowVlanPriorityMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter) Set(index int, newObj PatternFlowVlanPriorityMetricTag) PatternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowVlanPriorityMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter) Clear() PatternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowVlanPriorityMetricTag{}
		obj.patternFlowVlanPriorityMetricTagSlice = []PatternFlowVlanPriorityMetricTag{}
	}
	return obj
}
func (obj *patternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter) clearHolderSlice() PatternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter {
	if len(obj.patternFlowVlanPriorityMetricTagSlice) > 0 {
		obj.patternFlowVlanPriorityMetricTagSlice = []PatternFlowVlanPriorityMetricTag{}
	}
	return obj
}
func (obj *patternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter) appendHolderSlice(item PatternFlowVlanPriorityMetricTag) PatternFlowVlanPriorityPatternFlowVlanPriorityMetricTagIter {
	obj.patternFlowVlanPriorityMetricTagSlice = append(obj.patternFlowVlanPriorityMetricTagSlice, item)
	return obj
}

func (obj *patternFlowVlanPriority) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVlanPriority.Value <= 7 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 7 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowVlanPriority.Values <= 7 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

	if len(obj.obj.MetricTags) != 0 {

		if set_default {
			obj.MetricTags().clearHolderSlice()
			for _, item := range obj.obj.MetricTags {
				obj.MetricTags().appendHolderSlice(&patternFlowVlanPriorityMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowVlanPriority) setDefault() {
	var choices_set int = 0
	var choice PatternFlowVlanPriorityChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowVlanPriorityChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowVlanPriorityChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowVlanPriorityChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowVlanPriorityChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowVlanPriorityChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowVlanPriority")
			}
		} else {
			intVal := otg.PatternFlowVlanPriority_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowVlanPriority_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
