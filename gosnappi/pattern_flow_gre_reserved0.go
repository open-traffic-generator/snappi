package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGreReserved0 *****
type patternFlowGreReserved0 struct {
	validation
	obj              *otg.PatternFlowGreReserved0
	marshaller       marshalPatternFlowGreReserved0
	unMarshaller     unMarshalPatternFlowGreReserved0
	incrementHolder  PatternFlowGreReserved0Counter
	decrementHolder  PatternFlowGreReserved0Counter
	metricTagsHolder PatternFlowGreReserved0PatternFlowGreReserved0MetricTagIter
}

func NewPatternFlowGreReserved0() PatternFlowGreReserved0 {
	obj := patternFlowGreReserved0{obj: &otg.PatternFlowGreReserved0{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGreReserved0) msg() *otg.PatternFlowGreReserved0 {
	return obj.obj
}

func (obj *patternFlowGreReserved0) setMsg(msg *otg.PatternFlowGreReserved0) PatternFlowGreReserved0 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGreReserved0 struct {
	obj *patternFlowGreReserved0
}

type marshalPatternFlowGreReserved0 interface {
	// ToProto marshals PatternFlowGreReserved0 to protobuf object *otg.PatternFlowGreReserved0
	ToProto() (*otg.PatternFlowGreReserved0, error)
	// ToPbText marshals PatternFlowGreReserved0 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGreReserved0 to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGreReserved0 to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGreReserved0 struct {
	obj *patternFlowGreReserved0
}

type unMarshalPatternFlowGreReserved0 interface {
	// FromProto unmarshals PatternFlowGreReserved0 from protobuf object *otg.PatternFlowGreReserved0
	FromProto(msg *otg.PatternFlowGreReserved0) (PatternFlowGreReserved0, error)
	// FromPbText unmarshals PatternFlowGreReserved0 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGreReserved0 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGreReserved0 from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGreReserved0) Marshal() marshalPatternFlowGreReserved0 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGreReserved0{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGreReserved0) Unmarshal() unMarshalPatternFlowGreReserved0 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGreReserved0{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGreReserved0) ToProto() (*otg.PatternFlowGreReserved0, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGreReserved0) FromProto(msg *otg.PatternFlowGreReserved0) (PatternFlowGreReserved0, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGreReserved0) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGreReserved0) FromPbText(value string) error {
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

func (m *marshalpatternFlowGreReserved0) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGreReserved0) FromYaml(value string) error {
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

func (m *marshalpatternFlowGreReserved0) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGreReserved0) FromJson(value string) error {
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

func (obj *patternFlowGreReserved0) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGreReserved0) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGreReserved0) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGreReserved0) Clone() (PatternFlowGreReserved0, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGreReserved0()
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

func (obj *patternFlowGreReserved0) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGreReserved0 is reserved bits
type PatternFlowGreReserved0 interface {
	Validation
	// msg marshals PatternFlowGreReserved0 to protobuf object *otg.PatternFlowGreReserved0
	// and doesn't set defaults
	msg() *otg.PatternFlowGreReserved0
	// setMsg unmarshals PatternFlowGreReserved0 from protobuf object *otg.PatternFlowGreReserved0
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGreReserved0) PatternFlowGreReserved0
	// provides marshal interface
	Marshal() marshalPatternFlowGreReserved0
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGreReserved0
	// validate validates PatternFlowGreReserved0
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGreReserved0, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGreReserved0ChoiceEnum, set in PatternFlowGreReserved0
	Choice() PatternFlowGreReserved0ChoiceEnum
	// setChoice assigns PatternFlowGreReserved0ChoiceEnum provided by user to PatternFlowGreReserved0
	setChoice(value PatternFlowGreReserved0ChoiceEnum) PatternFlowGreReserved0
	// HasChoice checks if Choice has been set in PatternFlowGreReserved0
	HasChoice() bool
	// Value returns uint32, set in PatternFlowGreReserved0.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowGreReserved0
	SetValue(value uint32) PatternFlowGreReserved0
	// HasValue checks if Value has been set in PatternFlowGreReserved0
	HasValue() bool
	// Values returns []uint32, set in PatternFlowGreReserved0.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowGreReserved0
	SetValues(value []uint32) PatternFlowGreReserved0
	// Increment returns PatternFlowGreReserved0Counter, set in PatternFlowGreReserved0.
	// PatternFlowGreReserved0Counter is integer counter pattern
	Increment() PatternFlowGreReserved0Counter
	// SetIncrement assigns PatternFlowGreReserved0Counter provided by user to PatternFlowGreReserved0.
	// PatternFlowGreReserved0Counter is integer counter pattern
	SetIncrement(value PatternFlowGreReserved0Counter) PatternFlowGreReserved0
	// HasIncrement checks if Increment has been set in PatternFlowGreReserved0
	HasIncrement() bool
	// Decrement returns PatternFlowGreReserved0Counter, set in PatternFlowGreReserved0.
	// PatternFlowGreReserved0Counter is integer counter pattern
	Decrement() PatternFlowGreReserved0Counter
	// SetDecrement assigns PatternFlowGreReserved0Counter provided by user to PatternFlowGreReserved0.
	// PatternFlowGreReserved0Counter is integer counter pattern
	SetDecrement(value PatternFlowGreReserved0Counter) PatternFlowGreReserved0
	// HasDecrement checks if Decrement has been set in PatternFlowGreReserved0
	HasDecrement() bool
	// MetricTags returns PatternFlowGreReserved0PatternFlowGreReserved0MetricTagIterIter, set in PatternFlowGreReserved0
	MetricTags() PatternFlowGreReserved0PatternFlowGreReserved0MetricTagIter
	setNil()
}

type PatternFlowGreReserved0ChoiceEnum string

// Enum of Choice on PatternFlowGreReserved0
var PatternFlowGreReserved0Choice = struct {
	VALUE     PatternFlowGreReserved0ChoiceEnum
	VALUES    PatternFlowGreReserved0ChoiceEnum
	INCREMENT PatternFlowGreReserved0ChoiceEnum
	DECREMENT PatternFlowGreReserved0ChoiceEnum
}{
	VALUE:     PatternFlowGreReserved0ChoiceEnum("value"),
	VALUES:    PatternFlowGreReserved0ChoiceEnum("values"),
	INCREMENT: PatternFlowGreReserved0ChoiceEnum("increment"),
	DECREMENT: PatternFlowGreReserved0ChoiceEnum("decrement"),
}

func (obj *patternFlowGreReserved0) Choice() PatternFlowGreReserved0ChoiceEnum {
	return PatternFlowGreReserved0ChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGreReserved0) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGreReserved0) setChoice(value PatternFlowGreReserved0ChoiceEnum) PatternFlowGreReserved0 {
	intValue, ok := otg.PatternFlowGreReserved0_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGreReserved0ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGreReserved0_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGreReserved0Choice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGreReserved0Choice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGreReserved0Choice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGreReserved0Counter().msg()
	}

	if value == PatternFlowGreReserved0Choice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGreReserved0Counter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGreReserved0) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGreReserved0Choice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGreReserved0) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowGreReserved0 object
func (obj *patternFlowGreReserved0) SetValue(value uint32) PatternFlowGreReserved0 {
	obj.setChoice(PatternFlowGreReserved0Choice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowGreReserved0) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowGreReserved0 object
func (obj *patternFlowGreReserved0) SetValues(value []uint32) PatternFlowGreReserved0 {
	obj.setChoice(PatternFlowGreReserved0Choice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGreReserved0Counter
func (obj *patternFlowGreReserved0) Increment() PatternFlowGreReserved0Counter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGreReserved0Choice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGreReserved0Counter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGreReserved0Counter
func (obj *patternFlowGreReserved0) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGreReserved0Counter value in the PatternFlowGreReserved0 object
func (obj *patternFlowGreReserved0) SetIncrement(value PatternFlowGreReserved0Counter) PatternFlowGreReserved0 {
	obj.setChoice(PatternFlowGreReserved0Choice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGreReserved0Counter
func (obj *patternFlowGreReserved0) Decrement() PatternFlowGreReserved0Counter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGreReserved0Choice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGreReserved0Counter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGreReserved0Counter
func (obj *patternFlowGreReserved0) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGreReserved0Counter value in the PatternFlowGreReserved0 object
func (obj *patternFlowGreReserved0) SetDecrement(value PatternFlowGreReserved0Counter) PatternFlowGreReserved0 {
	obj.setChoice(PatternFlowGreReserved0Choice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGreReserved0MetricTag
func (obj *patternFlowGreReserved0) MetricTags() PatternFlowGreReserved0PatternFlowGreReserved0MetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGreReserved0MetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGreReserved0PatternFlowGreReserved0MetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGreReserved0PatternFlowGreReserved0MetricTagIter struct {
	obj                                   *patternFlowGreReserved0
	patternFlowGreReserved0MetricTagSlice []PatternFlowGreReserved0MetricTag
	fieldPtr                              *[]*otg.PatternFlowGreReserved0MetricTag
}

func newPatternFlowGreReserved0PatternFlowGreReserved0MetricTagIter(ptr *[]*otg.PatternFlowGreReserved0MetricTag) PatternFlowGreReserved0PatternFlowGreReserved0MetricTagIter {
	return &patternFlowGreReserved0PatternFlowGreReserved0MetricTagIter{fieldPtr: ptr}
}

type PatternFlowGreReserved0PatternFlowGreReserved0MetricTagIter interface {
	setMsg(*patternFlowGreReserved0) PatternFlowGreReserved0PatternFlowGreReserved0MetricTagIter
	Items() []PatternFlowGreReserved0MetricTag
	Add() PatternFlowGreReserved0MetricTag
	Append(items ...PatternFlowGreReserved0MetricTag) PatternFlowGreReserved0PatternFlowGreReserved0MetricTagIter
	Set(index int, newObj PatternFlowGreReserved0MetricTag) PatternFlowGreReserved0PatternFlowGreReserved0MetricTagIter
	Clear() PatternFlowGreReserved0PatternFlowGreReserved0MetricTagIter
	clearHolderSlice() PatternFlowGreReserved0PatternFlowGreReserved0MetricTagIter
	appendHolderSlice(item PatternFlowGreReserved0MetricTag) PatternFlowGreReserved0PatternFlowGreReserved0MetricTagIter
}

func (obj *patternFlowGreReserved0PatternFlowGreReserved0MetricTagIter) setMsg(msg *patternFlowGreReserved0) PatternFlowGreReserved0PatternFlowGreReserved0MetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGreReserved0MetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGreReserved0PatternFlowGreReserved0MetricTagIter) Items() []PatternFlowGreReserved0MetricTag {
	return obj.patternFlowGreReserved0MetricTagSlice
}

func (obj *patternFlowGreReserved0PatternFlowGreReserved0MetricTagIter) Add() PatternFlowGreReserved0MetricTag {
	newObj := &otg.PatternFlowGreReserved0MetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGreReserved0MetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGreReserved0MetricTagSlice = append(obj.patternFlowGreReserved0MetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGreReserved0PatternFlowGreReserved0MetricTagIter) Append(items ...PatternFlowGreReserved0MetricTag) PatternFlowGreReserved0PatternFlowGreReserved0MetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGreReserved0MetricTagSlice = append(obj.patternFlowGreReserved0MetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGreReserved0PatternFlowGreReserved0MetricTagIter) Set(index int, newObj PatternFlowGreReserved0MetricTag) PatternFlowGreReserved0PatternFlowGreReserved0MetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGreReserved0MetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGreReserved0PatternFlowGreReserved0MetricTagIter) Clear() PatternFlowGreReserved0PatternFlowGreReserved0MetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGreReserved0MetricTag{}
		obj.patternFlowGreReserved0MetricTagSlice = []PatternFlowGreReserved0MetricTag{}
	}
	return obj
}
func (obj *patternFlowGreReserved0PatternFlowGreReserved0MetricTagIter) clearHolderSlice() PatternFlowGreReserved0PatternFlowGreReserved0MetricTagIter {
	if len(obj.patternFlowGreReserved0MetricTagSlice) > 0 {
		obj.patternFlowGreReserved0MetricTagSlice = []PatternFlowGreReserved0MetricTag{}
	}
	return obj
}
func (obj *patternFlowGreReserved0PatternFlowGreReserved0MetricTagIter) appendHolderSlice(item PatternFlowGreReserved0MetricTag) PatternFlowGreReserved0PatternFlowGreReserved0MetricTagIter {
	obj.patternFlowGreReserved0MetricTagSlice = append(obj.patternFlowGreReserved0MetricTagSlice, item)
	return obj
}

func (obj *patternFlowGreReserved0) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 4095 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGreReserved0.Value <= 4095 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 4095 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowGreReserved0.Values <= 4095 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowGreReserved0MetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGreReserved0) setDefault() {
	var choices_set int = 0
	var choice PatternFlowGreReserved0ChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowGreReserved0Choice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowGreReserved0Choice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowGreReserved0Choice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowGreReserved0Choice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowGreReserved0Choice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowGreReserved0")
			}
		} else {
			intVal := otg.PatternFlowGreReserved0_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowGreReserved0_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
