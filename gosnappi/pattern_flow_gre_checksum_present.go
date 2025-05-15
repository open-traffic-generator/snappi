package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGreChecksumPresent *****
type patternFlowGreChecksumPresent struct {
	validation
	obj              *otg.PatternFlowGreChecksumPresent
	marshaller       marshalPatternFlowGreChecksumPresent
	unMarshaller     unMarshalPatternFlowGreChecksumPresent
	incrementHolder  PatternFlowGreChecksumPresentCounter
	decrementHolder  PatternFlowGreChecksumPresentCounter
	metricTagsHolder PatternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter
}

func NewPatternFlowGreChecksumPresent() PatternFlowGreChecksumPresent {
	obj := patternFlowGreChecksumPresent{obj: &otg.PatternFlowGreChecksumPresent{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGreChecksumPresent) msg() *otg.PatternFlowGreChecksumPresent {
	return obj.obj
}

func (obj *patternFlowGreChecksumPresent) setMsg(msg *otg.PatternFlowGreChecksumPresent) PatternFlowGreChecksumPresent {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGreChecksumPresent struct {
	obj *patternFlowGreChecksumPresent
}

type marshalPatternFlowGreChecksumPresent interface {
	// ToProto marshals PatternFlowGreChecksumPresent to protobuf object *otg.PatternFlowGreChecksumPresent
	ToProto() (*otg.PatternFlowGreChecksumPresent, error)
	// ToPbText marshals PatternFlowGreChecksumPresent to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGreChecksumPresent to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGreChecksumPresent to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGreChecksumPresent struct {
	obj *patternFlowGreChecksumPresent
}

type unMarshalPatternFlowGreChecksumPresent interface {
	// FromProto unmarshals PatternFlowGreChecksumPresent from protobuf object *otg.PatternFlowGreChecksumPresent
	FromProto(msg *otg.PatternFlowGreChecksumPresent) (PatternFlowGreChecksumPresent, error)
	// FromPbText unmarshals PatternFlowGreChecksumPresent from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGreChecksumPresent from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGreChecksumPresent from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGreChecksumPresent) Marshal() marshalPatternFlowGreChecksumPresent {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGreChecksumPresent{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGreChecksumPresent) Unmarshal() unMarshalPatternFlowGreChecksumPresent {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGreChecksumPresent{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGreChecksumPresent) ToProto() (*otg.PatternFlowGreChecksumPresent, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGreChecksumPresent) FromProto(msg *otg.PatternFlowGreChecksumPresent) (PatternFlowGreChecksumPresent, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGreChecksumPresent) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGreChecksumPresent) FromPbText(value string) error {
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

func (m *marshalpatternFlowGreChecksumPresent) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGreChecksumPresent) FromYaml(value string) error {
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

func (m *marshalpatternFlowGreChecksumPresent) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGreChecksumPresent) FromJson(value string) error {
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

func (obj *patternFlowGreChecksumPresent) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGreChecksumPresent) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGreChecksumPresent) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGreChecksumPresent) Clone() (PatternFlowGreChecksumPresent, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGreChecksumPresent()
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

func (obj *patternFlowGreChecksumPresent) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGreChecksumPresent is checksum present bit
type PatternFlowGreChecksumPresent interface {
	Validation
	// msg marshals PatternFlowGreChecksumPresent to protobuf object *otg.PatternFlowGreChecksumPresent
	// and doesn't set defaults
	msg() *otg.PatternFlowGreChecksumPresent
	// setMsg unmarshals PatternFlowGreChecksumPresent from protobuf object *otg.PatternFlowGreChecksumPresent
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGreChecksumPresent) PatternFlowGreChecksumPresent
	// provides marshal interface
	Marshal() marshalPatternFlowGreChecksumPresent
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGreChecksumPresent
	// validate validates PatternFlowGreChecksumPresent
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGreChecksumPresent, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGreChecksumPresentChoiceEnum, set in PatternFlowGreChecksumPresent
	Choice() PatternFlowGreChecksumPresentChoiceEnum
	// setChoice assigns PatternFlowGreChecksumPresentChoiceEnum provided by user to PatternFlowGreChecksumPresent
	setChoice(value PatternFlowGreChecksumPresentChoiceEnum) PatternFlowGreChecksumPresent
	// HasChoice checks if Choice has been set in PatternFlowGreChecksumPresent
	HasChoice() bool
	// Value returns uint32, set in PatternFlowGreChecksumPresent.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowGreChecksumPresent
	SetValue(value uint32) PatternFlowGreChecksumPresent
	// HasValue checks if Value has been set in PatternFlowGreChecksumPresent
	HasValue() bool
	// Values returns []uint32, set in PatternFlowGreChecksumPresent.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowGreChecksumPresent
	SetValues(value []uint32) PatternFlowGreChecksumPresent
	// Increment returns PatternFlowGreChecksumPresentCounter, set in PatternFlowGreChecksumPresent.
	// PatternFlowGreChecksumPresentCounter is integer counter pattern
	Increment() PatternFlowGreChecksumPresentCounter
	// SetIncrement assigns PatternFlowGreChecksumPresentCounter provided by user to PatternFlowGreChecksumPresent.
	// PatternFlowGreChecksumPresentCounter is integer counter pattern
	SetIncrement(value PatternFlowGreChecksumPresentCounter) PatternFlowGreChecksumPresent
	// HasIncrement checks if Increment has been set in PatternFlowGreChecksumPresent
	HasIncrement() bool
	// Decrement returns PatternFlowGreChecksumPresentCounter, set in PatternFlowGreChecksumPresent.
	// PatternFlowGreChecksumPresentCounter is integer counter pattern
	Decrement() PatternFlowGreChecksumPresentCounter
	// SetDecrement assigns PatternFlowGreChecksumPresentCounter provided by user to PatternFlowGreChecksumPresent.
	// PatternFlowGreChecksumPresentCounter is integer counter pattern
	SetDecrement(value PatternFlowGreChecksumPresentCounter) PatternFlowGreChecksumPresent
	// HasDecrement checks if Decrement has been set in PatternFlowGreChecksumPresent
	HasDecrement() bool
	// MetricTags returns PatternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIterIter, set in PatternFlowGreChecksumPresent
	MetricTags() PatternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter
	setNil()
}

type PatternFlowGreChecksumPresentChoiceEnum string

// Enum of Choice on PatternFlowGreChecksumPresent
var PatternFlowGreChecksumPresentChoice = struct {
	VALUE     PatternFlowGreChecksumPresentChoiceEnum
	VALUES    PatternFlowGreChecksumPresentChoiceEnum
	INCREMENT PatternFlowGreChecksumPresentChoiceEnum
	DECREMENT PatternFlowGreChecksumPresentChoiceEnum
}{
	VALUE:     PatternFlowGreChecksumPresentChoiceEnum("value"),
	VALUES:    PatternFlowGreChecksumPresentChoiceEnum("values"),
	INCREMENT: PatternFlowGreChecksumPresentChoiceEnum("increment"),
	DECREMENT: PatternFlowGreChecksumPresentChoiceEnum("decrement"),
}

func (obj *patternFlowGreChecksumPresent) Choice() PatternFlowGreChecksumPresentChoiceEnum {
	return PatternFlowGreChecksumPresentChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGreChecksumPresent) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGreChecksumPresent) setChoice(value PatternFlowGreChecksumPresentChoiceEnum) PatternFlowGreChecksumPresent {
	intValue, ok := otg.PatternFlowGreChecksumPresent_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGreChecksumPresentChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGreChecksumPresent_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGreChecksumPresentChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGreChecksumPresentChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGreChecksumPresentChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGreChecksumPresentCounter().msg()
	}

	if value == PatternFlowGreChecksumPresentChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGreChecksumPresentCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGreChecksumPresent) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGreChecksumPresentChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGreChecksumPresent) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowGreChecksumPresent object
func (obj *patternFlowGreChecksumPresent) SetValue(value uint32) PatternFlowGreChecksumPresent {
	obj.setChoice(PatternFlowGreChecksumPresentChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowGreChecksumPresent) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowGreChecksumPresent object
func (obj *patternFlowGreChecksumPresent) SetValues(value []uint32) PatternFlowGreChecksumPresent {
	obj.setChoice(PatternFlowGreChecksumPresentChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGreChecksumPresentCounter
func (obj *patternFlowGreChecksumPresent) Increment() PatternFlowGreChecksumPresentCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGreChecksumPresentChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGreChecksumPresentCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGreChecksumPresentCounter
func (obj *patternFlowGreChecksumPresent) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGreChecksumPresentCounter value in the PatternFlowGreChecksumPresent object
func (obj *patternFlowGreChecksumPresent) SetIncrement(value PatternFlowGreChecksumPresentCounter) PatternFlowGreChecksumPresent {
	obj.setChoice(PatternFlowGreChecksumPresentChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGreChecksumPresentCounter
func (obj *patternFlowGreChecksumPresent) Decrement() PatternFlowGreChecksumPresentCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGreChecksumPresentChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGreChecksumPresentCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGreChecksumPresentCounter
func (obj *patternFlowGreChecksumPresent) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGreChecksumPresentCounter value in the PatternFlowGreChecksumPresent object
func (obj *patternFlowGreChecksumPresent) SetDecrement(value PatternFlowGreChecksumPresentCounter) PatternFlowGreChecksumPresent {
	obj.setChoice(PatternFlowGreChecksumPresentChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGreChecksumPresentMetricTag
func (obj *patternFlowGreChecksumPresent) MetricTags() PatternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGreChecksumPresentMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter struct {
	obj                                         *patternFlowGreChecksumPresent
	patternFlowGreChecksumPresentMetricTagSlice []PatternFlowGreChecksumPresentMetricTag
	fieldPtr                                    *[]*otg.PatternFlowGreChecksumPresentMetricTag
}

func newPatternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter(ptr *[]*otg.PatternFlowGreChecksumPresentMetricTag) PatternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter {
	return &patternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter{fieldPtr: ptr}
}

type PatternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter interface {
	setMsg(*patternFlowGreChecksumPresent) PatternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter
	Items() []PatternFlowGreChecksumPresentMetricTag
	Add() PatternFlowGreChecksumPresentMetricTag
	Append(items ...PatternFlowGreChecksumPresentMetricTag) PatternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter
	Set(index int, newObj PatternFlowGreChecksumPresentMetricTag) PatternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter
	Clear() PatternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter
	clearHolderSlice() PatternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter
	appendHolderSlice(item PatternFlowGreChecksumPresentMetricTag) PatternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter
}

func (obj *patternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter) setMsg(msg *patternFlowGreChecksumPresent) PatternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGreChecksumPresentMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter) Items() []PatternFlowGreChecksumPresentMetricTag {
	return obj.patternFlowGreChecksumPresentMetricTagSlice
}

func (obj *patternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter) Add() PatternFlowGreChecksumPresentMetricTag {
	newObj := &otg.PatternFlowGreChecksumPresentMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGreChecksumPresentMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGreChecksumPresentMetricTagSlice = append(obj.patternFlowGreChecksumPresentMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter) Append(items ...PatternFlowGreChecksumPresentMetricTag) PatternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGreChecksumPresentMetricTagSlice = append(obj.patternFlowGreChecksumPresentMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter) Set(index int, newObj PatternFlowGreChecksumPresentMetricTag) PatternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGreChecksumPresentMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter) Clear() PatternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGreChecksumPresentMetricTag{}
		obj.patternFlowGreChecksumPresentMetricTagSlice = []PatternFlowGreChecksumPresentMetricTag{}
	}
	return obj
}
func (obj *patternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter) clearHolderSlice() PatternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter {
	if len(obj.patternFlowGreChecksumPresentMetricTagSlice) > 0 {
		obj.patternFlowGreChecksumPresentMetricTagSlice = []PatternFlowGreChecksumPresentMetricTag{}
	}
	return obj
}
func (obj *patternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter) appendHolderSlice(item PatternFlowGreChecksumPresentMetricTag) PatternFlowGreChecksumPresentPatternFlowGreChecksumPresentMetricTagIter {
	obj.patternFlowGreChecksumPresentMetricTagSlice = append(obj.patternFlowGreChecksumPresentMetricTagSlice, item)
	return obj
}

func (obj *patternFlowGreChecksumPresent) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGreChecksumPresent.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowGreChecksumPresent.Values <= 1 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowGreChecksumPresentMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGreChecksumPresent) setDefault() {
	var choices_set int = 0
	var choice PatternFlowGreChecksumPresentChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowGreChecksumPresentChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowGreChecksumPresentChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowGreChecksumPresentChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowGreChecksumPresentChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowGreChecksumPresentChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowGreChecksumPresent")
			}
		} else {
			intVal := otg.PatternFlowGreChecksumPresent_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowGreChecksumPresent_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
