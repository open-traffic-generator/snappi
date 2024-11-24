package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpSeqNum *****
type patternFlowTcpSeqNum struct {
	validation
	obj              *otg.PatternFlowTcpSeqNum
	marshaller       marshalPatternFlowTcpSeqNum
	unMarshaller     unMarshalPatternFlowTcpSeqNum
	incrementHolder  PatternFlowTcpSeqNumCounter
	decrementHolder  PatternFlowTcpSeqNumCounter
	metricTagsHolder PatternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter
}

func NewPatternFlowTcpSeqNum() PatternFlowTcpSeqNum {
	obj := patternFlowTcpSeqNum{obj: &otg.PatternFlowTcpSeqNum{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpSeqNum) msg() *otg.PatternFlowTcpSeqNum {
	return obj.obj
}

func (obj *patternFlowTcpSeqNum) setMsg(msg *otg.PatternFlowTcpSeqNum) PatternFlowTcpSeqNum {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpSeqNum struct {
	obj *patternFlowTcpSeqNum
}

type marshalPatternFlowTcpSeqNum interface {
	// ToProto marshals PatternFlowTcpSeqNum to protobuf object *otg.PatternFlowTcpSeqNum
	ToProto() (*otg.PatternFlowTcpSeqNum, error)
	// ToPbText marshals PatternFlowTcpSeqNum to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpSeqNum to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpSeqNum to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowTcpSeqNum to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowTcpSeqNum struct {
	obj *patternFlowTcpSeqNum
}

type unMarshalPatternFlowTcpSeqNum interface {
	// FromProto unmarshals PatternFlowTcpSeqNum from protobuf object *otg.PatternFlowTcpSeqNum
	FromProto(msg *otg.PatternFlowTcpSeqNum) (PatternFlowTcpSeqNum, error)
	// FromPbText unmarshals PatternFlowTcpSeqNum from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpSeqNum from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpSeqNum from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpSeqNum) Marshal() marshalPatternFlowTcpSeqNum {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpSeqNum{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpSeqNum) Unmarshal() unMarshalPatternFlowTcpSeqNum {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpSeqNum{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpSeqNum) ToProto() (*otg.PatternFlowTcpSeqNum, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpSeqNum) FromProto(msg *otg.PatternFlowTcpSeqNum) (PatternFlowTcpSeqNum, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpSeqNum) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpSeqNum) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpSeqNum) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpSeqNum) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpSeqNum) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalpatternFlowTcpSeqNum) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpSeqNum) FromJson(value string) error {
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

func (obj *patternFlowTcpSeqNum) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpSeqNum) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpSeqNum) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpSeqNum) Clone() (PatternFlowTcpSeqNum, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpSeqNum()
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

func (obj *patternFlowTcpSeqNum) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowTcpSeqNum is sequence number
type PatternFlowTcpSeqNum interface {
	Validation
	// msg marshals PatternFlowTcpSeqNum to protobuf object *otg.PatternFlowTcpSeqNum
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpSeqNum
	// setMsg unmarshals PatternFlowTcpSeqNum from protobuf object *otg.PatternFlowTcpSeqNum
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpSeqNum) PatternFlowTcpSeqNum
	// provides marshal interface
	Marshal() marshalPatternFlowTcpSeqNum
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpSeqNum
	// validate validates PatternFlowTcpSeqNum
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpSeqNum, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowTcpSeqNumChoiceEnum, set in PatternFlowTcpSeqNum
	Choice() PatternFlowTcpSeqNumChoiceEnum
	// setChoice assigns PatternFlowTcpSeqNumChoiceEnum provided by user to PatternFlowTcpSeqNum
	setChoice(value PatternFlowTcpSeqNumChoiceEnum) PatternFlowTcpSeqNum
	// HasChoice checks if Choice has been set in PatternFlowTcpSeqNum
	HasChoice() bool
	// Value returns uint32, set in PatternFlowTcpSeqNum.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowTcpSeqNum
	SetValue(value uint32) PatternFlowTcpSeqNum
	// HasValue checks if Value has been set in PatternFlowTcpSeqNum
	HasValue() bool
	// Values returns []uint32, set in PatternFlowTcpSeqNum.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowTcpSeqNum
	SetValues(value []uint32) PatternFlowTcpSeqNum
	// Increment returns PatternFlowTcpSeqNumCounter, set in PatternFlowTcpSeqNum.
	// PatternFlowTcpSeqNumCounter is integer counter pattern
	Increment() PatternFlowTcpSeqNumCounter
	// SetIncrement assigns PatternFlowTcpSeqNumCounter provided by user to PatternFlowTcpSeqNum.
	// PatternFlowTcpSeqNumCounter is integer counter pattern
	SetIncrement(value PatternFlowTcpSeqNumCounter) PatternFlowTcpSeqNum
	// HasIncrement checks if Increment has been set in PatternFlowTcpSeqNum
	HasIncrement() bool
	// Decrement returns PatternFlowTcpSeqNumCounter, set in PatternFlowTcpSeqNum.
	// PatternFlowTcpSeqNumCounter is integer counter pattern
	Decrement() PatternFlowTcpSeqNumCounter
	// SetDecrement assigns PatternFlowTcpSeqNumCounter provided by user to PatternFlowTcpSeqNum.
	// PatternFlowTcpSeqNumCounter is integer counter pattern
	SetDecrement(value PatternFlowTcpSeqNumCounter) PatternFlowTcpSeqNum
	// HasDecrement checks if Decrement has been set in PatternFlowTcpSeqNum
	HasDecrement() bool
	// MetricTags returns PatternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIterIter, set in PatternFlowTcpSeqNum
	MetricTags() PatternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter
	setNil()
}

type PatternFlowTcpSeqNumChoiceEnum string

// Enum of Choice on PatternFlowTcpSeqNum
var PatternFlowTcpSeqNumChoice = struct {
	VALUE     PatternFlowTcpSeqNumChoiceEnum
	VALUES    PatternFlowTcpSeqNumChoiceEnum
	INCREMENT PatternFlowTcpSeqNumChoiceEnum
	DECREMENT PatternFlowTcpSeqNumChoiceEnum
}{
	VALUE:     PatternFlowTcpSeqNumChoiceEnum("value"),
	VALUES:    PatternFlowTcpSeqNumChoiceEnum("values"),
	INCREMENT: PatternFlowTcpSeqNumChoiceEnum("increment"),
	DECREMENT: PatternFlowTcpSeqNumChoiceEnum("decrement"),
}

func (obj *patternFlowTcpSeqNum) Choice() PatternFlowTcpSeqNumChoiceEnum {
	return PatternFlowTcpSeqNumChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowTcpSeqNum) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowTcpSeqNum) setChoice(value PatternFlowTcpSeqNumChoiceEnum) PatternFlowTcpSeqNum {
	intValue, ok := otg.PatternFlowTcpSeqNum_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowTcpSeqNumChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowTcpSeqNum_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowTcpSeqNumChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowTcpSeqNumChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowTcpSeqNumChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowTcpSeqNumCounter().msg()
	}

	if value == PatternFlowTcpSeqNumChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowTcpSeqNumCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpSeqNum) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowTcpSeqNumChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpSeqNum) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowTcpSeqNum object
func (obj *patternFlowTcpSeqNum) SetValue(value uint32) PatternFlowTcpSeqNum {
	obj.setChoice(PatternFlowTcpSeqNumChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowTcpSeqNum) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowTcpSeqNum object
func (obj *patternFlowTcpSeqNum) SetValues(value []uint32) PatternFlowTcpSeqNum {
	obj.setChoice(PatternFlowTcpSeqNumChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowTcpSeqNumCounter
func (obj *patternFlowTcpSeqNum) Increment() PatternFlowTcpSeqNumCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowTcpSeqNumChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowTcpSeqNumCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowTcpSeqNumCounter
func (obj *patternFlowTcpSeqNum) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowTcpSeqNumCounter value in the PatternFlowTcpSeqNum object
func (obj *patternFlowTcpSeqNum) SetIncrement(value PatternFlowTcpSeqNumCounter) PatternFlowTcpSeqNum {
	obj.setChoice(PatternFlowTcpSeqNumChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowTcpSeqNumCounter
func (obj *patternFlowTcpSeqNum) Decrement() PatternFlowTcpSeqNumCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowTcpSeqNumChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowTcpSeqNumCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowTcpSeqNumCounter
func (obj *patternFlowTcpSeqNum) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowTcpSeqNumCounter value in the PatternFlowTcpSeqNum object
func (obj *patternFlowTcpSeqNum) SetDecrement(value PatternFlowTcpSeqNumCounter) PatternFlowTcpSeqNum {
	obj.setChoice(PatternFlowTcpSeqNumChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowTcpSeqNumMetricTag
func (obj *patternFlowTcpSeqNum) MetricTags() PatternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowTcpSeqNumMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter struct {
	obj                                *patternFlowTcpSeqNum
	patternFlowTcpSeqNumMetricTagSlice []PatternFlowTcpSeqNumMetricTag
	fieldPtr                           *[]*otg.PatternFlowTcpSeqNumMetricTag
}

func newPatternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter(ptr *[]*otg.PatternFlowTcpSeqNumMetricTag) PatternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter {
	return &patternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter{fieldPtr: ptr}
}

type PatternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter interface {
	setMsg(*patternFlowTcpSeqNum) PatternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter
	Items() []PatternFlowTcpSeqNumMetricTag
	Add() PatternFlowTcpSeqNumMetricTag
	Append(items ...PatternFlowTcpSeqNumMetricTag) PatternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter
	Set(index int, newObj PatternFlowTcpSeqNumMetricTag) PatternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter
	Clear() PatternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter
	clearHolderSlice() PatternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter
	appendHolderSlice(item PatternFlowTcpSeqNumMetricTag) PatternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter
}

func (obj *patternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter) setMsg(msg *patternFlowTcpSeqNum) PatternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowTcpSeqNumMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter) Items() []PatternFlowTcpSeqNumMetricTag {
	return obj.patternFlowTcpSeqNumMetricTagSlice
}

func (obj *patternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter) Add() PatternFlowTcpSeqNumMetricTag {
	newObj := &otg.PatternFlowTcpSeqNumMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowTcpSeqNumMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowTcpSeqNumMetricTagSlice = append(obj.patternFlowTcpSeqNumMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter) Append(items ...PatternFlowTcpSeqNumMetricTag) PatternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowTcpSeqNumMetricTagSlice = append(obj.patternFlowTcpSeqNumMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter) Set(index int, newObj PatternFlowTcpSeqNumMetricTag) PatternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowTcpSeqNumMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter) Clear() PatternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowTcpSeqNumMetricTag{}
		obj.patternFlowTcpSeqNumMetricTagSlice = []PatternFlowTcpSeqNumMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter) clearHolderSlice() PatternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter {
	if len(obj.patternFlowTcpSeqNumMetricTagSlice) > 0 {
		obj.patternFlowTcpSeqNumMetricTagSlice = []PatternFlowTcpSeqNumMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter) appendHolderSlice(item PatternFlowTcpSeqNumMetricTag) PatternFlowTcpSeqNumPatternFlowTcpSeqNumMetricTagIter {
	obj.patternFlowTcpSeqNumMetricTagSlice = append(obj.patternFlowTcpSeqNumMetricTagSlice, item)
	return obj
}

func (obj *patternFlowTcpSeqNum) validateObj(vObj *validation, set_default bool) {
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
				obj.MetricTags().appendHolderSlice(&patternFlowTcpSeqNumMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowTcpSeqNum) setDefault() {
	var choices_set int = 0
	var choice PatternFlowTcpSeqNumChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowTcpSeqNumChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowTcpSeqNumChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowTcpSeqNumChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowTcpSeqNumChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowTcpSeqNumChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowTcpSeqNum")
			}
		} else {
			intVal := otg.PatternFlowTcpSeqNum_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowTcpSeqNum_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
