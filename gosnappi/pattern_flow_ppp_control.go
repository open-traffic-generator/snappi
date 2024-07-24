package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPppControl *****
type patternFlowPppControl struct {
	validation
	obj              *otg.PatternFlowPppControl
	marshaller       marshalPatternFlowPppControl
	unMarshaller     unMarshalPatternFlowPppControl
	incrementHolder  PatternFlowPppControlCounter
	decrementHolder  PatternFlowPppControlCounter
	metricTagsHolder PatternFlowPppControlPatternFlowPppControlMetricTagIter
}

func NewPatternFlowPppControl() PatternFlowPppControl {
	obj := patternFlowPppControl{obj: &otg.PatternFlowPppControl{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPppControl) msg() *otg.PatternFlowPppControl {
	return obj.obj
}

func (obj *patternFlowPppControl) setMsg(msg *otg.PatternFlowPppControl) PatternFlowPppControl {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPppControl struct {
	obj *patternFlowPppControl
}

type marshalPatternFlowPppControl interface {
	// ToProto marshals PatternFlowPppControl to protobuf object *otg.PatternFlowPppControl
	ToProto() (*otg.PatternFlowPppControl, error)
	// ToPbText marshals PatternFlowPppControl to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPppControl to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPppControl to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPppControl struct {
	obj *patternFlowPppControl
}

type unMarshalPatternFlowPppControl interface {
	// FromProto unmarshals PatternFlowPppControl from protobuf object *otg.PatternFlowPppControl
	FromProto(msg *otg.PatternFlowPppControl) (PatternFlowPppControl, error)
	// FromPbText unmarshals PatternFlowPppControl from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPppControl from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPppControl from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPppControl) Marshal() marshalPatternFlowPppControl {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPppControl{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPppControl) Unmarshal() unMarshalPatternFlowPppControl {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPppControl{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPppControl) ToProto() (*otg.PatternFlowPppControl, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPppControl) FromProto(msg *otg.PatternFlowPppControl) (PatternFlowPppControl, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPppControl) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPppControl) FromPbText(value string) error {
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

func (m *marshalpatternFlowPppControl) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPppControl) FromYaml(value string) error {
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

func (m *marshalpatternFlowPppControl) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPppControl) FromJson(value string) error {
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

func (obj *patternFlowPppControl) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPppControl) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPppControl) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPppControl) Clone() (PatternFlowPppControl, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPppControl()
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

func (obj *patternFlowPppControl) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowPppControl is pPP control
type PatternFlowPppControl interface {
	Validation
	// msg marshals PatternFlowPppControl to protobuf object *otg.PatternFlowPppControl
	// and doesn't set defaults
	msg() *otg.PatternFlowPppControl
	// setMsg unmarshals PatternFlowPppControl from protobuf object *otg.PatternFlowPppControl
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPppControl) PatternFlowPppControl
	// provides marshal interface
	Marshal() marshalPatternFlowPppControl
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPppControl
	// validate validates PatternFlowPppControl
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPppControl, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowPppControlChoiceEnum, set in PatternFlowPppControl
	Choice() PatternFlowPppControlChoiceEnum
	// setChoice assigns PatternFlowPppControlChoiceEnum provided by user to PatternFlowPppControl
	setChoice(value PatternFlowPppControlChoiceEnum) PatternFlowPppControl
	// HasChoice checks if Choice has been set in PatternFlowPppControl
	HasChoice() bool
	// Value returns uint32, set in PatternFlowPppControl.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowPppControl
	SetValue(value uint32) PatternFlowPppControl
	// HasValue checks if Value has been set in PatternFlowPppControl
	HasValue() bool
	// Values returns []uint32, set in PatternFlowPppControl.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowPppControl
	SetValues(value []uint32) PatternFlowPppControl
	// Increment returns PatternFlowPppControlCounter, set in PatternFlowPppControl.
	// PatternFlowPppControlCounter is integer counter pattern
	Increment() PatternFlowPppControlCounter
	// SetIncrement assigns PatternFlowPppControlCounter provided by user to PatternFlowPppControl.
	// PatternFlowPppControlCounter is integer counter pattern
	SetIncrement(value PatternFlowPppControlCounter) PatternFlowPppControl
	// HasIncrement checks if Increment has been set in PatternFlowPppControl
	HasIncrement() bool
	// Decrement returns PatternFlowPppControlCounter, set in PatternFlowPppControl.
	// PatternFlowPppControlCounter is integer counter pattern
	Decrement() PatternFlowPppControlCounter
	// SetDecrement assigns PatternFlowPppControlCounter provided by user to PatternFlowPppControl.
	// PatternFlowPppControlCounter is integer counter pattern
	SetDecrement(value PatternFlowPppControlCounter) PatternFlowPppControl
	// HasDecrement checks if Decrement has been set in PatternFlowPppControl
	HasDecrement() bool
	// MetricTags returns PatternFlowPppControlPatternFlowPppControlMetricTagIterIter, set in PatternFlowPppControl
	MetricTags() PatternFlowPppControlPatternFlowPppControlMetricTagIter
	setNil()
}

type PatternFlowPppControlChoiceEnum string

// Enum of Choice on PatternFlowPppControl
var PatternFlowPppControlChoice = struct {
	VALUE     PatternFlowPppControlChoiceEnum
	VALUES    PatternFlowPppControlChoiceEnum
	INCREMENT PatternFlowPppControlChoiceEnum
	DECREMENT PatternFlowPppControlChoiceEnum
}{
	VALUE:     PatternFlowPppControlChoiceEnum("value"),
	VALUES:    PatternFlowPppControlChoiceEnum("values"),
	INCREMENT: PatternFlowPppControlChoiceEnum("increment"),
	DECREMENT: PatternFlowPppControlChoiceEnum("decrement"),
}

func (obj *patternFlowPppControl) Choice() PatternFlowPppControlChoiceEnum {
	return PatternFlowPppControlChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowPppControl) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowPppControl) setChoice(value PatternFlowPppControlChoiceEnum) PatternFlowPppControl {
	intValue, ok := otg.PatternFlowPppControl_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowPppControlChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowPppControl_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowPppControlChoice.VALUE {
		defaultValue := uint32(3)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowPppControlChoice.VALUES {
		defaultValue := []uint32{3}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowPppControlChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowPppControlCounter().msg()
	}

	if value == PatternFlowPppControlChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowPppControlCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowPppControl) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowPppControlChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowPppControl) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowPppControl object
func (obj *patternFlowPppControl) SetValue(value uint32) PatternFlowPppControl {
	obj.setChoice(PatternFlowPppControlChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowPppControl) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{3})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowPppControl object
func (obj *patternFlowPppControl) SetValues(value []uint32) PatternFlowPppControl {
	obj.setChoice(PatternFlowPppControlChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowPppControlCounter
func (obj *patternFlowPppControl) Increment() PatternFlowPppControlCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowPppControlChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowPppControlCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowPppControlCounter
func (obj *patternFlowPppControl) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowPppControlCounter value in the PatternFlowPppControl object
func (obj *patternFlowPppControl) SetIncrement(value PatternFlowPppControlCounter) PatternFlowPppControl {
	obj.setChoice(PatternFlowPppControlChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowPppControlCounter
func (obj *patternFlowPppControl) Decrement() PatternFlowPppControlCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowPppControlChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowPppControlCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowPppControlCounter
func (obj *patternFlowPppControl) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowPppControlCounter value in the PatternFlowPppControl object
func (obj *patternFlowPppControl) SetDecrement(value PatternFlowPppControlCounter) PatternFlowPppControl {
	obj.setChoice(PatternFlowPppControlChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowPppControlMetricTag
func (obj *patternFlowPppControl) MetricTags() PatternFlowPppControlPatternFlowPppControlMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowPppControlMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowPppControlPatternFlowPppControlMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowPppControlPatternFlowPppControlMetricTagIter struct {
	obj                                 *patternFlowPppControl
	patternFlowPppControlMetricTagSlice []PatternFlowPppControlMetricTag
	fieldPtr                            *[]*otg.PatternFlowPppControlMetricTag
}

func newPatternFlowPppControlPatternFlowPppControlMetricTagIter(ptr *[]*otg.PatternFlowPppControlMetricTag) PatternFlowPppControlPatternFlowPppControlMetricTagIter {
	return &patternFlowPppControlPatternFlowPppControlMetricTagIter{fieldPtr: ptr}
}

type PatternFlowPppControlPatternFlowPppControlMetricTagIter interface {
	setMsg(*patternFlowPppControl) PatternFlowPppControlPatternFlowPppControlMetricTagIter
	Items() []PatternFlowPppControlMetricTag
	Add() PatternFlowPppControlMetricTag
	Append(items ...PatternFlowPppControlMetricTag) PatternFlowPppControlPatternFlowPppControlMetricTagIter
	Set(index int, newObj PatternFlowPppControlMetricTag) PatternFlowPppControlPatternFlowPppControlMetricTagIter
	Clear() PatternFlowPppControlPatternFlowPppControlMetricTagIter
	clearHolderSlice() PatternFlowPppControlPatternFlowPppControlMetricTagIter
	appendHolderSlice(item PatternFlowPppControlMetricTag) PatternFlowPppControlPatternFlowPppControlMetricTagIter
}

func (obj *patternFlowPppControlPatternFlowPppControlMetricTagIter) setMsg(msg *patternFlowPppControl) PatternFlowPppControlPatternFlowPppControlMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowPppControlMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowPppControlPatternFlowPppControlMetricTagIter) Items() []PatternFlowPppControlMetricTag {
	return obj.patternFlowPppControlMetricTagSlice
}

func (obj *patternFlowPppControlPatternFlowPppControlMetricTagIter) Add() PatternFlowPppControlMetricTag {
	newObj := &otg.PatternFlowPppControlMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowPppControlMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowPppControlMetricTagSlice = append(obj.patternFlowPppControlMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowPppControlPatternFlowPppControlMetricTagIter) Append(items ...PatternFlowPppControlMetricTag) PatternFlowPppControlPatternFlowPppControlMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowPppControlMetricTagSlice = append(obj.patternFlowPppControlMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowPppControlPatternFlowPppControlMetricTagIter) Set(index int, newObj PatternFlowPppControlMetricTag) PatternFlowPppControlPatternFlowPppControlMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowPppControlMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowPppControlPatternFlowPppControlMetricTagIter) Clear() PatternFlowPppControlPatternFlowPppControlMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowPppControlMetricTag{}
		obj.patternFlowPppControlMetricTagSlice = []PatternFlowPppControlMetricTag{}
	}
	return obj
}
func (obj *patternFlowPppControlPatternFlowPppControlMetricTagIter) clearHolderSlice() PatternFlowPppControlPatternFlowPppControlMetricTagIter {
	if len(obj.patternFlowPppControlMetricTagSlice) > 0 {
		obj.patternFlowPppControlMetricTagSlice = []PatternFlowPppControlMetricTag{}
	}
	return obj
}
func (obj *patternFlowPppControlPatternFlowPppControlMetricTagIter) appendHolderSlice(item PatternFlowPppControlMetricTag) PatternFlowPppControlPatternFlowPppControlMetricTagIter {
	obj.patternFlowPppControlMetricTagSlice = append(obj.patternFlowPppControlMetricTagSlice, item)
	return obj
}

func (obj *patternFlowPppControl) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPppControl.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowPppControl.Values <= 255 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowPppControlMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowPppControl) setDefault() {
	var choices_set int = 0
	var choice PatternFlowPppControlChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowPppControlChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowPppControlChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowPppControlChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowPppControlChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowPppControlChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowPppControl")
			}
		} else {
			intVal := otg.PatternFlowPppControl_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowPppControl_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
