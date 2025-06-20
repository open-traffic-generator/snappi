package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowVlanCfi *****
type patternFlowVlanCfi struct {
	validation
	obj              *otg.PatternFlowVlanCfi
	marshaller       marshalPatternFlowVlanCfi
	unMarshaller     unMarshalPatternFlowVlanCfi
	incrementHolder  PatternFlowVlanCfiCounter
	decrementHolder  PatternFlowVlanCfiCounter
	metricTagsHolder PatternFlowVlanCfiPatternFlowVlanCfiMetricTagIter
}

func NewPatternFlowVlanCfi() PatternFlowVlanCfi {
	obj := patternFlowVlanCfi{obj: &otg.PatternFlowVlanCfi{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowVlanCfi) msg() *otg.PatternFlowVlanCfi {
	return obj.obj
}

func (obj *patternFlowVlanCfi) setMsg(msg *otg.PatternFlowVlanCfi) PatternFlowVlanCfi {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowVlanCfi struct {
	obj *patternFlowVlanCfi
}

type marshalPatternFlowVlanCfi interface {
	// ToProto marshals PatternFlowVlanCfi to protobuf object *otg.PatternFlowVlanCfi
	ToProto() (*otg.PatternFlowVlanCfi, error)
	// ToPbText marshals PatternFlowVlanCfi to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowVlanCfi to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowVlanCfi to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowVlanCfi struct {
	obj *patternFlowVlanCfi
}

type unMarshalPatternFlowVlanCfi interface {
	// FromProto unmarshals PatternFlowVlanCfi from protobuf object *otg.PatternFlowVlanCfi
	FromProto(msg *otg.PatternFlowVlanCfi) (PatternFlowVlanCfi, error)
	// FromPbText unmarshals PatternFlowVlanCfi from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowVlanCfi from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowVlanCfi from JSON text
	FromJson(value string) error
}

func (obj *patternFlowVlanCfi) Marshal() marshalPatternFlowVlanCfi {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowVlanCfi{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowVlanCfi) Unmarshal() unMarshalPatternFlowVlanCfi {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowVlanCfi{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowVlanCfi) ToProto() (*otg.PatternFlowVlanCfi, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowVlanCfi) FromProto(msg *otg.PatternFlowVlanCfi) (PatternFlowVlanCfi, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowVlanCfi) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowVlanCfi) FromPbText(value string) error {
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

func (m *marshalpatternFlowVlanCfi) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowVlanCfi) FromYaml(value string) error {
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

func (m *marshalpatternFlowVlanCfi) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowVlanCfi) FromJson(value string) error {
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

func (obj *patternFlowVlanCfi) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowVlanCfi) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowVlanCfi) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowVlanCfi) Clone() (PatternFlowVlanCfi, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowVlanCfi()
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

func (obj *patternFlowVlanCfi) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowVlanCfi is canonical format indicator or drop elegible indicator
type PatternFlowVlanCfi interface {
	Validation
	// msg marshals PatternFlowVlanCfi to protobuf object *otg.PatternFlowVlanCfi
	// and doesn't set defaults
	msg() *otg.PatternFlowVlanCfi
	// setMsg unmarshals PatternFlowVlanCfi from protobuf object *otg.PatternFlowVlanCfi
	// and doesn't set defaults
	setMsg(*otg.PatternFlowVlanCfi) PatternFlowVlanCfi
	// provides marshal interface
	Marshal() marshalPatternFlowVlanCfi
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowVlanCfi
	// validate validates PatternFlowVlanCfi
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowVlanCfi, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowVlanCfiChoiceEnum, set in PatternFlowVlanCfi
	Choice() PatternFlowVlanCfiChoiceEnum
	// setChoice assigns PatternFlowVlanCfiChoiceEnum provided by user to PatternFlowVlanCfi
	setChoice(value PatternFlowVlanCfiChoiceEnum) PatternFlowVlanCfi
	// HasChoice checks if Choice has been set in PatternFlowVlanCfi
	HasChoice() bool
	// Value returns uint32, set in PatternFlowVlanCfi.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowVlanCfi
	SetValue(value uint32) PatternFlowVlanCfi
	// HasValue checks if Value has been set in PatternFlowVlanCfi
	HasValue() bool
	// Values returns []uint32, set in PatternFlowVlanCfi.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowVlanCfi
	SetValues(value []uint32) PatternFlowVlanCfi
	// Increment returns PatternFlowVlanCfiCounter, set in PatternFlowVlanCfi.
	// PatternFlowVlanCfiCounter is integer counter pattern
	Increment() PatternFlowVlanCfiCounter
	// SetIncrement assigns PatternFlowVlanCfiCounter provided by user to PatternFlowVlanCfi.
	// PatternFlowVlanCfiCounter is integer counter pattern
	SetIncrement(value PatternFlowVlanCfiCounter) PatternFlowVlanCfi
	// HasIncrement checks if Increment has been set in PatternFlowVlanCfi
	HasIncrement() bool
	// Decrement returns PatternFlowVlanCfiCounter, set in PatternFlowVlanCfi.
	// PatternFlowVlanCfiCounter is integer counter pattern
	Decrement() PatternFlowVlanCfiCounter
	// SetDecrement assigns PatternFlowVlanCfiCounter provided by user to PatternFlowVlanCfi.
	// PatternFlowVlanCfiCounter is integer counter pattern
	SetDecrement(value PatternFlowVlanCfiCounter) PatternFlowVlanCfi
	// HasDecrement checks if Decrement has been set in PatternFlowVlanCfi
	HasDecrement() bool
	// MetricTags returns PatternFlowVlanCfiPatternFlowVlanCfiMetricTagIterIter, set in PatternFlowVlanCfi
	MetricTags() PatternFlowVlanCfiPatternFlowVlanCfiMetricTagIter
	setNil()
}

type PatternFlowVlanCfiChoiceEnum string

// Enum of Choice on PatternFlowVlanCfi
var PatternFlowVlanCfiChoice = struct {
	VALUE     PatternFlowVlanCfiChoiceEnum
	VALUES    PatternFlowVlanCfiChoiceEnum
	INCREMENT PatternFlowVlanCfiChoiceEnum
	DECREMENT PatternFlowVlanCfiChoiceEnum
}{
	VALUE:     PatternFlowVlanCfiChoiceEnum("value"),
	VALUES:    PatternFlowVlanCfiChoiceEnum("values"),
	INCREMENT: PatternFlowVlanCfiChoiceEnum("increment"),
	DECREMENT: PatternFlowVlanCfiChoiceEnum("decrement"),
}

func (obj *patternFlowVlanCfi) Choice() PatternFlowVlanCfiChoiceEnum {
	return PatternFlowVlanCfiChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowVlanCfi) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowVlanCfi) setChoice(value PatternFlowVlanCfiChoiceEnum) PatternFlowVlanCfi {
	intValue, ok := otg.PatternFlowVlanCfi_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowVlanCfiChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowVlanCfi_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowVlanCfiChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowVlanCfiChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowVlanCfiChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowVlanCfiCounter().msg()
	}

	if value == PatternFlowVlanCfiChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowVlanCfiCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowVlanCfi) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowVlanCfiChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowVlanCfi) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowVlanCfi object
func (obj *patternFlowVlanCfi) SetValue(value uint32) PatternFlowVlanCfi {
	obj.setChoice(PatternFlowVlanCfiChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowVlanCfi) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowVlanCfi object
func (obj *patternFlowVlanCfi) SetValues(value []uint32) PatternFlowVlanCfi {
	obj.setChoice(PatternFlowVlanCfiChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowVlanCfiCounter
func (obj *patternFlowVlanCfi) Increment() PatternFlowVlanCfiCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowVlanCfiChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowVlanCfiCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowVlanCfiCounter
func (obj *patternFlowVlanCfi) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowVlanCfiCounter value in the PatternFlowVlanCfi object
func (obj *patternFlowVlanCfi) SetIncrement(value PatternFlowVlanCfiCounter) PatternFlowVlanCfi {
	obj.setChoice(PatternFlowVlanCfiChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowVlanCfiCounter
func (obj *patternFlowVlanCfi) Decrement() PatternFlowVlanCfiCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowVlanCfiChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowVlanCfiCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowVlanCfiCounter
func (obj *patternFlowVlanCfi) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowVlanCfiCounter value in the PatternFlowVlanCfi object
func (obj *patternFlowVlanCfi) SetDecrement(value PatternFlowVlanCfiCounter) PatternFlowVlanCfi {
	obj.setChoice(PatternFlowVlanCfiChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowVlanCfiMetricTag
func (obj *patternFlowVlanCfi) MetricTags() PatternFlowVlanCfiPatternFlowVlanCfiMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowVlanCfiMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowVlanCfiPatternFlowVlanCfiMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowVlanCfiPatternFlowVlanCfiMetricTagIter struct {
	obj                              *patternFlowVlanCfi
	patternFlowVlanCfiMetricTagSlice []PatternFlowVlanCfiMetricTag
	fieldPtr                         *[]*otg.PatternFlowVlanCfiMetricTag
}

func newPatternFlowVlanCfiPatternFlowVlanCfiMetricTagIter(ptr *[]*otg.PatternFlowVlanCfiMetricTag) PatternFlowVlanCfiPatternFlowVlanCfiMetricTagIter {
	return &patternFlowVlanCfiPatternFlowVlanCfiMetricTagIter{fieldPtr: ptr}
}

type PatternFlowVlanCfiPatternFlowVlanCfiMetricTagIter interface {
	setMsg(*patternFlowVlanCfi) PatternFlowVlanCfiPatternFlowVlanCfiMetricTagIter
	Items() []PatternFlowVlanCfiMetricTag
	Add() PatternFlowVlanCfiMetricTag
	Append(items ...PatternFlowVlanCfiMetricTag) PatternFlowVlanCfiPatternFlowVlanCfiMetricTagIter
	Set(index int, newObj PatternFlowVlanCfiMetricTag) PatternFlowVlanCfiPatternFlowVlanCfiMetricTagIter
	Clear() PatternFlowVlanCfiPatternFlowVlanCfiMetricTagIter
	clearHolderSlice() PatternFlowVlanCfiPatternFlowVlanCfiMetricTagIter
	appendHolderSlice(item PatternFlowVlanCfiMetricTag) PatternFlowVlanCfiPatternFlowVlanCfiMetricTagIter
}

func (obj *patternFlowVlanCfiPatternFlowVlanCfiMetricTagIter) setMsg(msg *patternFlowVlanCfi) PatternFlowVlanCfiPatternFlowVlanCfiMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowVlanCfiMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowVlanCfiPatternFlowVlanCfiMetricTagIter) Items() []PatternFlowVlanCfiMetricTag {
	return obj.patternFlowVlanCfiMetricTagSlice
}

func (obj *patternFlowVlanCfiPatternFlowVlanCfiMetricTagIter) Add() PatternFlowVlanCfiMetricTag {
	newObj := &otg.PatternFlowVlanCfiMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowVlanCfiMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowVlanCfiMetricTagSlice = append(obj.patternFlowVlanCfiMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowVlanCfiPatternFlowVlanCfiMetricTagIter) Append(items ...PatternFlowVlanCfiMetricTag) PatternFlowVlanCfiPatternFlowVlanCfiMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowVlanCfiMetricTagSlice = append(obj.patternFlowVlanCfiMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowVlanCfiPatternFlowVlanCfiMetricTagIter) Set(index int, newObj PatternFlowVlanCfiMetricTag) PatternFlowVlanCfiPatternFlowVlanCfiMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowVlanCfiMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowVlanCfiPatternFlowVlanCfiMetricTagIter) Clear() PatternFlowVlanCfiPatternFlowVlanCfiMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowVlanCfiMetricTag{}
		obj.patternFlowVlanCfiMetricTagSlice = []PatternFlowVlanCfiMetricTag{}
	}
	return obj
}
func (obj *patternFlowVlanCfiPatternFlowVlanCfiMetricTagIter) clearHolderSlice() PatternFlowVlanCfiPatternFlowVlanCfiMetricTagIter {
	if len(obj.patternFlowVlanCfiMetricTagSlice) > 0 {
		obj.patternFlowVlanCfiMetricTagSlice = []PatternFlowVlanCfiMetricTag{}
	}
	return obj
}
func (obj *patternFlowVlanCfiPatternFlowVlanCfiMetricTagIter) appendHolderSlice(item PatternFlowVlanCfiMetricTag) PatternFlowVlanCfiPatternFlowVlanCfiMetricTagIter {
	obj.patternFlowVlanCfiMetricTagSlice = append(obj.patternFlowVlanCfiMetricTagSlice, item)
	return obj
}

func (obj *patternFlowVlanCfi) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVlanCfi.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowVlanCfi.Values <= 1 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowVlanCfiMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowVlanCfi) setDefault() {
	var choices_set int = 0
	var choice PatternFlowVlanCfiChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowVlanCfiChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowVlanCfiChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowVlanCfiChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowVlanCfiChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowVlanCfiChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowVlanCfi")
			}
		} else {
			intVal := otg.PatternFlowVlanCfi_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowVlanCfi_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
