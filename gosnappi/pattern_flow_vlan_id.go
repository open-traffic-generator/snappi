package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowVlanId *****
type patternFlowVlanId struct {
	validation
	obj              *otg.PatternFlowVlanId
	marshaller       marshalPatternFlowVlanId
	unMarshaller     unMarshalPatternFlowVlanId
	incrementHolder  PatternFlowVlanIdCounter
	decrementHolder  PatternFlowVlanIdCounter
	metricTagsHolder PatternFlowVlanIdPatternFlowVlanIdMetricTagIter
}

func NewPatternFlowVlanId() PatternFlowVlanId {
	obj := patternFlowVlanId{obj: &otg.PatternFlowVlanId{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowVlanId) msg() *otg.PatternFlowVlanId {
	return obj.obj
}

func (obj *patternFlowVlanId) setMsg(msg *otg.PatternFlowVlanId) PatternFlowVlanId {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowVlanId struct {
	obj *patternFlowVlanId
}

type marshalPatternFlowVlanId interface {
	// ToProto marshals PatternFlowVlanId to protobuf object *otg.PatternFlowVlanId
	ToProto() (*otg.PatternFlowVlanId, error)
	// ToPbText marshals PatternFlowVlanId to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowVlanId to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowVlanId to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowVlanId to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowVlanId struct {
	obj *patternFlowVlanId
}

type unMarshalPatternFlowVlanId interface {
	// FromProto unmarshals PatternFlowVlanId from protobuf object *otg.PatternFlowVlanId
	FromProto(msg *otg.PatternFlowVlanId) (PatternFlowVlanId, error)
	// FromPbText unmarshals PatternFlowVlanId from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowVlanId from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowVlanId from JSON text
	FromJson(value string) error
}

func (obj *patternFlowVlanId) Marshal() marshalPatternFlowVlanId {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowVlanId{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowVlanId) Unmarshal() unMarshalPatternFlowVlanId {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowVlanId{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowVlanId) ToProto() (*otg.PatternFlowVlanId, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowVlanId) FromProto(msg *otg.PatternFlowVlanId) (PatternFlowVlanId, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowVlanId) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowVlanId) FromPbText(value string) error {
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

func (m *marshalpatternFlowVlanId) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowVlanId) FromYaml(value string) error {
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

func (m *marshalpatternFlowVlanId) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowVlanId) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowVlanId) FromJson(value string) error {
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

func (obj *patternFlowVlanId) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowVlanId) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowVlanId) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowVlanId) Clone() (PatternFlowVlanId, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowVlanId()
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

func (obj *patternFlowVlanId) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowVlanId is vlan identifier
type PatternFlowVlanId interface {
	Validation
	// msg marshals PatternFlowVlanId to protobuf object *otg.PatternFlowVlanId
	// and doesn't set defaults
	msg() *otg.PatternFlowVlanId
	// setMsg unmarshals PatternFlowVlanId from protobuf object *otg.PatternFlowVlanId
	// and doesn't set defaults
	setMsg(*otg.PatternFlowVlanId) PatternFlowVlanId
	// provides marshal interface
	Marshal() marshalPatternFlowVlanId
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowVlanId
	// validate validates PatternFlowVlanId
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowVlanId, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowVlanIdChoiceEnum, set in PatternFlowVlanId
	Choice() PatternFlowVlanIdChoiceEnum
	// setChoice assigns PatternFlowVlanIdChoiceEnum provided by user to PatternFlowVlanId
	setChoice(value PatternFlowVlanIdChoiceEnum) PatternFlowVlanId
	// HasChoice checks if Choice has been set in PatternFlowVlanId
	HasChoice() bool
	// Value returns uint32, set in PatternFlowVlanId.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowVlanId
	SetValue(value uint32) PatternFlowVlanId
	// HasValue checks if Value has been set in PatternFlowVlanId
	HasValue() bool
	// Values returns []uint32, set in PatternFlowVlanId.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowVlanId
	SetValues(value []uint32) PatternFlowVlanId
	// Increment returns PatternFlowVlanIdCounter, set in PatternFlowVlanId.
	// PatternFlowVlanIdCounter is integer counter pattern
	Increment() PatternFlowVlanIdCounter
	// SetIncrement assigns PatternFlowVlanIdCounter provided by user to PatternFlowVlanId.
	// PatternFlowVlanIdCounter is integer counter pattern
	SetIncrement(value PatternFlowVlanIdCounter) PatternFlowVlanId
	// HasIncrement checks if Increment has been set in PatternFlowVlanId
	HasIncrement() bool
	// Decrement returns PatternFlowVlanIdCounter, set in PatternFlowVlanId.
	// PatternFlowVlanIdCounter is integer counter pattern
	Decrement() PatternFlowVlanIdCounter
	// SetDecrement assigns PatternFlowVlanIdCounter provided by user to PatternFlowVlanId.
	// PatternFlowVlanIdCounter is integer counter pattern
	SetDecrement(value PatternFlowVlanIdCounter) PatternFlowVlanId
	// HasDecrement checks if Decrement has been set in PatternFlowVlanId
	HasDecrement() bool
	// MetricTags returns PatternFlowVlanIdPatternFlowVlanIdMetricTagIterIter, set in PatternFlowVlanId
	MetricTags() PatternFlowVlanIdPatternFlowVlanIdMetricTagIter
	setNil()
}

type PatternFlowVlanIdChoiceEnum string

// Enum of Choice on PatternFlowVlanId
var PatternFlowVlanIdChoice = struct {
	VALUE     PatternFlowVlanIdChoiceEnum
	VALUES    PatternFlowVlanIdChoiceEnum
	INCREMENT PatternFlowVlanIdChoiceEnum
	DECREMENT PatternFlowVlanIdChoiceEnum
}{
	VALUE:     PatternFlowVlanIdChoiceEnum("value"),
	VALUES:    PatternFlowVlanIdChoiceEnum("values"),
	INCREMENT: PatternFlowVlanIdChoiceEnum("increment"),
	DECREMENT: PatternFlowVlanIdChoiceEnum("decrement"),
}

func (obj *patternFlowVlanId) Choice() PatternFlowVlanIdChoiceEnum {
	return PatternFlowVlanIdChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowVlanId) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowVlanId) setChoice(value PatternFlowVlanIdChoiceEnum) PatternFlowVlanId {
	intValue, ok := otg.PatternFlowVlanId_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowVlanIdChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowVlanId_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowVlanIdChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowVlanIdChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowVlanIdChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowVlanIdCounter().msg()
	}

	if value == PatternFlowVlanIdChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowVlanIdCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowVlanId) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowVlanIdChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowVlanId) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowVlanId object
func (obj *patternFlowVlanId) SetValue(value uint32) PatternFlowVlanId {
	obj.setChoice(PatternFlowVlanIdChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowVlanId) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowVlanId object
func (obj *patternFlowVlanId) SetValues(value []uint32) PatternFlowVlanId {
	obj.setChoice(PatternFlowVlanIdChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowVlanIdCounter
func (obj *patternFlowVlanId) Increment() PatternFlowVlanIdCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowVlanIdChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowVlanIdCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowVlanIdCounter
func (obj *patternFlowVlanId) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowVlanIdCounter value in the PatternFlowVlanId object
func (obj *patternFlowVlanId) SetIncrement(value PatternFlowVlanIdCounter) PatternFlowVlanId {
	obj.setChoice(PatternFlowVlanIdChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowVlanIdCounter
func (obj *patternFlowVlanId) Decrement() PatternFlowVlanIdCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowVlanIdChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowVlanIdCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowVlanIdCounter
func (obj *patternFlowVlanId) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowVlanIdCounter value in the PatternFlowVlanId object
func (obj *patternFlowVlanId) SetDecrement(value PatternFlowVlanIdCounter) PatternFlowVlanId {
	obj.setChoice(PatternFlowVlanIdChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowVlanIdMetricTag
func (obj *patternFlowVlanId) MetricTags() PatternFlowVlanIdPatternFlowVlanIdMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowVlanIdMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowVlanIdPatternFlowVlanIdMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowVlanIdPatternFlowVlanIdMetricTagIter struct {
	obj                             *patternFlowVlanId
	patternFlowVlanIdMetricTagSlice []PatternFlowVlanIdMetricTag
	fieldPtr                        *[]*otg.PatternFlowVlanIdMetricTag
}

func newPatternFlowVlanIdPatternFlowVlanIdMetricTagIter(ptr *[]*otg.PatternFlowVlanIdMetricTag) PatternFlowVlanIdPatternFlowVlanIdMetricTagIter {
	return &patternFlowVlanIdPatternFlowVlanIdMetricTagIter{fieldPtr: ptr}
}

type PatternFlowVlanIdPatternFlowVlanIdMetricTagIter interface {
	setMsg(*patternFlowVlanId) PatternFlowVlanIdPatternFlowVlanIdMetricTagIter
	Items() []PatternFlowVlanIdMetricTag
	Add() PatternFlowVlanIdMetricTag
	Append(items ...PatternFlowVlanIdMetricTag) PatternFlowVlanIdPatternFlowVlanIdMetricTagIter
	Set(index int, newObj PatternFlowVlanIdMetricTag) PatternFlowVlanIdPatternFlowVlanIdMetricTagIter
	Clear() PatternFlowVlanIdPatternFlowVlanIdMetricTagIter
	clearHolderSlice() PatternFlowVlanIdPatternFlowVlanIdMetricTagIter
	appendHolderSlice(item PatternFlowVlanIdMetricTag) PatternFlowVlanIdPatternFlowVlanIdMetricTagIter
}

func (obj *patternFlowVlanIdPatternFlowVlanIdMetricTagIter) setMsg(msg *patternFlowVlanId) PatternFlowVlanIdPatternFlowVlanIdMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowVlanIdMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowVlanIdPatternFlowVlanIdMetricTagIter) Items() []PatternFlowVlanIdMetricTag {
	return obj.patternFlowVlanIdMetricTagSlice
}

func (obj *patternFlowVlanIdPatternFlowVlanIdMetricTagIter) Add() PatternFlowVlanIdMetricTag {
	newObj := &otg.PatternFlowVlanIdMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowVlanIdMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowVlanIdMetricTagSlice = append(obj.patternFlowVlanIdMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowVlanIdPatternFlowVlanIdMetricTagIter) Append(items ...PatternFlowVlanIdMetricTag) PatternFlowVlanIdPatternFlowVlanIdMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowVlanIdMetricTagSlice = append(obj.patternFlowVlanIdMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowVlanIdPatternFlowVlanIdMetricTagIter) Set(index int, newObj PatternFlowVlanIdMetricTag) PatternFlowVlanIdPatternFlowVlanIdMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowVlanIdMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowVlanIdPatternFlowVlanIdMetricTagIter) Clear() PatternFlowVlanIdPatternFlowVlanIdMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowVlanIdMetricTag{}
		obj.patternFlowVlanIdMetricTagSlice = []PatternFlowVlanIdMetricTag{}
	}
	return obj
}
func (obj *patternFlowVlanIdPatternFlowVlanIdMetricTagIter) clearHolderSlice() PatternFlowVlanIdPatternFlowVlanIdMetricTagIter {
	if len(obj.patternFlowVlanIdMetricTagSlice) > 0 {
		obj.patternFlowVlanIdMetricTagSlice = []PatternFlowVlanIdMetricTag{}
	}
	return obj
}
func (obj *patternFlowVlanIdPatternFlowVlanIdMetricTagIter) appendHolderSlice(item PatternFlowVlanIdMetricTag) PatternFlowVlanIdPatternFlowVlanIdMetricTagIter {
	obj.patternFlowVlanIdMetricTagSlice = append(obj.patternFlowVlanIdMetricTagSlice, item)
	return obj
}

func (obj *patternFlowVlanId) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 4095 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVlanId.Value <= 4095 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 4095 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowVlanId.Values <= 4095 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowVlanIdMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowVlanId) setDefault() {
	var choices_set int = 0
	var choice PatternFlowVlanIdChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowVlanIdChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowVlanIdChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowVlanIdChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowVlanIdChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowVlanIdChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowVlanId")
			}
		} else {
			intVal := otg.PatternFlowVlanId_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowVlanId_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
