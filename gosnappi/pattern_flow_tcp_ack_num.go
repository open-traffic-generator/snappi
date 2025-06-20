package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpAckNum *****
type patternFlowTcpAckNum struct {
	validation
	obj              *otg.PatternFlowTcpAckNum
	marshaller       marshalPatternFlowTcpAckNum
	unMarshaller     unMarshalPatternFlowTcpAckNum
	incrementHolder  PatternFlowTcpAckNumCounter
	decrementHolder  PatternFlowTcpAckNumCounter
	metricTagsHolder PatternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter
}

func NewPatternFlowTcpAckNum() PatternFlowTcpAckNum {
	obj := patternFlowTcpAckNum{obj: &otg.PatternFlowTcpAckNum{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpAckNum) msg() *otg.PatternFlowTcpAckNum {
	return obj.obj
}

func (obj *patternFlowTcpAckNum) setMsg(msg *otg.PatternFlowTcpAckNum) PatternFlowTcpAckNum {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpAckNum struct {
	obj *patternFlowTcpAckNum
}

type marshalPatternFlowTcpAckNum interface {
	// ToProto marshals PatternFlowTcpAckNum to protobuf object *otg.PatternFlowTcpAckNum
	ToProto() (*otg.PatternFlowTcpAckNum, error)
	// ToPbText marshals PatternFlowTcpAckNum to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpAckNum to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpAckNum to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpAckNum struct {
	obj *patternFlowTcpAckNum
}

type unMarshalPatternFlowTcpAckNum interface {
	// FromProto unmarshals PatternFlowTcpAckNum from protobuf object *otg.PatternFlowTcpAckNum
	FromProto(msg *otg.PatternFlowTcpAckNum) (PatternFlowTcpAckNum, error)
	// FromPbText unmarshals PatternFlowTcpAckNum from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpAckNum from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpAckNum from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpAckNum) Marshal() marshalPatternFlowTcpAckNum {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpAckNum{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpAckNum) Unmarshal() unMarshalPatternFlowTcpAckNum {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpAckNum{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpAckNum) ToProto() (*otg.PatternFlowTcpAckNum, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpAckNum) FromProto(msg *otg.PatternFlowTcpAckNum) (PatternFlowTcpAckNum, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpAckNum) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpAckNum) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpAckNum) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpAckNum) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpAckNum) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpAckNum) FromJson(value string) error {
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

func (obj *patternFlowTcpAckNum) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpAckNum) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpAckNum) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpAckNum) Clone() (PatternFlowTcpAckNum, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpAckNum()
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

func (obj *patternFlowTcpAckNum) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowTcpAckNum is acknowledgement number
type PatternFlowTcpAckNum interface {
	Validation
	// msg marshals PatternFlowTcpAckNum to protobuf object *otg.PatternFlowTcpAckNum
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpAckNum
	// setMsg unmarshals PatternFlowTcpAckNum from protobuf object *otg.PatternFlowTcpAckNum
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpAckNum) PatternFlowTcpAckNum
	// provides marshal interface
	Marshal() marshalPatternFlowTcpAckNum
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpAckNum
	// validate validates PatternFlowTcpAckNum
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpAckNum, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowTcpAckNumChoiceEnum, set in PatternFlowTcpAckNum
	Choice() PatternFlowTcpAckNumChoiceEnum
	// setChoice assigns PatternFlowTcpAckNumChoiceEnum provided by user to PatternFlowTcpAckNum
	setChoice(value PatternFlowTcpAckNumChoiceEnum) PatternFlowTcpAckNum
	// HasChoice checks if Choice has been set in PatternFlowTcpAckNum
	HasChoice() bool
	// Value returns uint32, set in PatternFlowTcpAckNum.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowTcpAckNum
	SetValue(value uint32) PatternFlowTcpAckNum
	// HasValue checks if Value has been set in PatternFlowTcpAckNum
	HasValue() bool
	// Values returns []uint32, set in PatternFlowTcpAckNum.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowTcpAckNum
	SetValues(value []uint32) PatternFlowTcpAckNum
	// Increment returns PatternFlowTcpAckNumCounter, set in PatternFlowTcpAckNum.
	// PatternFlowTcpAckNumCounter is integer counter pattern
	Increment() PatternFlowTcpAckNumCounter
	// SetIncrement assigns PatternFlowTcpAckNumCounter provided by user to PatternFlowTcpAckNum.
	// PatternFlowTcpAckNumCounter is integer counter pattern
	SetIncrement(value PatternFlowTcpAckNumCounter) PatternFlowTcpAckNum
	// HasIncrement checks if Increment has been set in PatternFlowTcpAckNum
	HasIncrement() bool
	// Decrement returns PatternFlowTcpAckNumCounter, set in PatternFlowTcpAckNum.
	// PatternFlowTcpAckNumCounter is integer counter pattern
	Decrement() PatternFlowTcpAckNumCounter
	// SetDecrement assigns PatternFlowTcpAckNumCounter provided by user to PatternFlowTcpAckNum.
	// PatternFlowTcpAckNumCounter is integer counter pattern
	SetDecrement(value PatternFlowTcpAckNumCounter) PatternFlowTcpAckNum
	// HasDecrement checks if Decrement has been set in PatternFlowTcpAckNum
	HasDecrement() bool
	// MetricTags returns PatternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIterIter, set in PatternFlowTcpAckNum
	MetricTags() PatternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter
	setNil()
}

type PatternFlowTcpAckNumChoiceEnum string

// Enum of Choice on PatternFlowTcpAckNum
var PatternFlowTcpAckNumChoice = struct {
	VALUE     PatternFlowTcpAckNumChoiceEnum
	VALUES    PatternFlowTcpAckNumChoiceEnum
	INCREMENT PatternFlowTcpAckNumChoiceEnum
	DECREMENT PatternFlowTcpAckNumChoiceEnum
}{
	VALUE:     PatternFlowTcpAckNumChoiceEnum("value"),
	VALUES:    PatternFlowTcpAckNumChoiceEnum("values"),
	INCREMENT: PatternFlowTcpAckNumChoiceEnum("increment"),
	DECREMENT: PatternFlowTcpAckNumChoiceEnum("decrement"),
}

func (obj *patternFlowTcpAckNum) Choice() PatternFlowTcpAckNumChoiceEnum {
	return PatternFlowTcpAckNumChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowTcpAckNum) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowTcpAckNum) setChoice(value PatternFlowTcpAckNumChoiceEnum) PatternFlowTcpAckNum {
	intValue, ok := otg.PatternFlowTcpAckNum_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowTcpAckNumChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowTcpAckNum_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowTcpAckNumChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowTcpAckNumChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowTcpAckNumChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowTcpAckNumCounter().msg()
	}

	if value == PatternFlowTcpAckNumChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowTcpAckNumCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpAckNum) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowTcpAckNumChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpAckNum) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowTcpAckNum object
func (obj *patternFlowTcpAckNum) SetValue(value uint32) PatternFlowTcpAckNum {
	obj.setChoice(PatternFlowTcpAckNumChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowTcpAckNum) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowTcpAckNum object
func (obj *patternFlowTcpAckNum) SetValues(value []uint32) PatternFlowTcpAckNum {
	obj.setChoice(PatternFlowTcpAckNumChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowTcpAckNumCounter
func (obj *patternFlowTcpAckNum) Increment() PatternFlowTcpAckNumCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowTcpAckNumChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowTcpAckNumCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowTcpAckNumCounter
func (obj *patternFlowTcpAckNum) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowTcpAckNumCounter value in the PatternFlowTcpAckNum object
func (obj *patternFlowTcpAckNum) SetIncrement(value PatternFlowTcpAckNumCounter) PatternFlowTcpAckNum {
	obj.setChoice(PatternFlowTcpAckNumChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowTcpAckNumCounter
func (obj *patternFlowTcpAckNum) Decrement() PatternFlowTcpAckNumCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowTcpAckNumChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowTcpAckNumCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowTcpAckNumCounter
func (obj *patternFlowTcpAckNum) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowTcpAckNumCounter value in the PatternFlowTcpAckNum object
func (obj *patternFlowTcpAckNum) SetDecrement(value PatternFlowTcpAckNumCounter) PatternFlowTcpAckNum {
	obj.setChoice(PatternFlowTcpAckNumChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowTcpAckNumMetricTag
func (obj *patternFlowTcpAckNum) MetricTags() PatternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowTcpAckNumMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter struct {
	obj                                *patternFlowTcpAckNum
	patternFlowTcpAckNumMetricTagSlice []PatternFlowTcpAckNumMetricTag
	fieldPtr                           *[]*otg.PatternFlowTcpAckNumMetricTag
}

func newPatternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter(ptr *[]*otg.PatternFlowTcpAckNumMetricTag) PatternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter {
	return &patternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter{fieldPtr: ptr}
}

type PatternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter interface {
	setMsg(*patternFlowTcpAckNum) PatternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter
	Items() []PatternFlowTcpAckNumMetricTag
	Add() PatternFlowTcpAckNumMetricTag
	Append(items ...PatternFlowTcpAckNumMetricTag) PatternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter
	Set(index int, newObj PatternFlowTcpAckNumMetricTag) PatternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter
	Clear() PatternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter
	clearHolderSlice() PatternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter
	appendHolderSlice(item PatternFlowTcpAckNumMetricTag) PatternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter
}

func (obj *patternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter) setMsg(msg *patternFlowTcpAckNum) PatternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowTcpAckNumMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter) Items() []PatternFlowTcpAckNumMetricTag {
	return obj.patternFlowTcpAckNumMetricTagSlice
}

func (obj *patternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter) Add() PatternFlowTcpAckNumMetricTag {
	newObj := &otg.PatternFlowTcpAckNumMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowTcpAckNumMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowTcpAckNumMetricTagSlice = append(obj.patternFlowTcpAckNumMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter) Append(items ...PatternFlowTcpAckNumMetricTag) PatternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowTcpAckNumMetricTagSlice = append(obj.patternFlowTcpAckNumMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter) Set(index int, newObj PatternFlowTcpAckNumMetricTag) PatternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowTcpAckNumMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter) Clear() PatternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowTcpAckNumMetricTag{}
		obj.patternFlowTcpAckNumMetricTagSlice = []PatternFlowTcpAckNumMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter) clearHolderSlice() PatternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter {
	if len(obj.patternFlowTcpAckNumMetricTagSlice) > 0 {
		obj.patternFlowTcpAckNumMetricTagSlice = []PatternFlowTcpAckNumMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter) appendHolderSlice(item PatternFlowTcpAckNumMetricTag) PatternFlowTcpAckNumPatternFlowTcpAckNumMetricTagIter {
	obj.patternFlowTcpAckNumMetricTagSlice = append(obj.patternFlowTcpAckNumMetricTagSlice, item)
	return obj
}

func (obj *patternFlowTcpAckNum) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
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
				obj.MetricTags().appendHolderSlice(&patternFlowTcpAckNumMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowTcpAckNum) setDefault() {
	var choices_set int = 0
	var choice PatternFlowTcpAckNumChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowTcpAckNumChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowTcpAckNumChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowTcpAckNumChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowTcpAckNumChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowTcpAckNumChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowTcpAckNum")
			}
		} else {
			intVal := otg.PatternFlowTcpAckNum_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowTcpAckNum_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
