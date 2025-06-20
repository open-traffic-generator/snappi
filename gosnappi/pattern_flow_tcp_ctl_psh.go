package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpCtlPsh *****
type patternFlowTcpCtlPsh struct {
	validation
	obj              *otg.PatternFlowTcpCtlPsh
	marshaller       marshalPatternFlowTcpCtlPsh
	unMarshaller     unMarshalPatternFlowTcpCtlPsh
	incrementHolder  PatternFlowTcpCtlPshCounter
	decrementHolder  PatternFlowTcpCtlPshCounter
	metricTagsHolder PatternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter
}

func NewPatternFlowTcpCtlPsh() PatternFlowTcpCtlPsh {
	obj := patternFlowTcpCtlPsh{obj: &otg.PatternFlowTcpCtlPsh{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpCtlPsh) msg() *otg.PatternFlowTcpCtlPsh {
	return obj.obj
}

func (obj *patternFlowTcpCtlPsh) setMsg(msg *otg.PatternFlowTcpCtlPsh) PatternFlowTcpCtlPsh {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpCtlPsh struct {
	obj *patternFlowTcpCtlPsh
}

type marshalPatternFlowTcpCtlPsh interface {
	// ToProto marshals PatternFlowTcpCtlPsh to protobuf object *otg.PatternFlowTcpCtlPsh
	ToProto() (*otg.PatternFlowTcpCtlPsh, error)
	// ToPbText marshals PatternFlowTcpCtlPsh to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpCtlPsh to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpCtlPsh to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpCtlPsh struct {
	obj *patternFlowTcpCtlPsh
}

type unMarshalPatternFlowTcpCtlPsh interface {
	// FromProto unmarshals PatternFlowTcpCtlPsh from protobuf object *otg.PatternFlowTcpCtlPsh
	FromProto(msg *otg.PatternFlowTcpCtlPsh) (PatternFlowTcpCtlPsh, error)
	// FromPbText unmarshals PatternFlowTcpCtlPsh from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpCtlPsh from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpCtlPsh from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpCtlPsh) Marshal() marshalPatternFlowTcpCtlPsh {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpCtlPsh{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpCtlPsh) Unmarshal() unMarshalPatternFlowTcpCtlPsh {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpCtlPsh{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpCtlPsh) ToProto() (*otg.PatternFlowTcpCtlPsh, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpCtlPsh) FromProto(msg *otg.PatternFlowTcpCtlPsh) (PatternFlowTcpCtlPsh, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpCtlPsh) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlPsh) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpCtlPsh) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlPsh) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpCtlPsh) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlPsh) FromJson(value string) error {
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

func (obj *patternFlowTcpCtlPsh) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlPsh) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlPsh) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpCtlPsh) Clone() (PatternFlowTcpCtlPsh, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpCtlPsh()
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

func (obj *patternFlowTcpCtlPsh) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowTcpCtlPsh is asks to push the buffered data to the receiving application.
type PatternFlowTcpCtlPsh interface {
	Validation
	// msg marshals PatternFlowTcpCtlPsh to protobuf object *otg.PatternFlowTcpCtlPsh
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpCtlPsh
	// setMsg unmarshals PatternFlowTcpCtlPsh from protobuf object *otg.PatternFlowTcpCtlPsh
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpCtlPsh) PatternFlowTcpCtlPsh
	// provides marshal interface
	Marshal() marshalPatternFlowTcpCtlPsh
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpCtlPsh
	// validate validates PatternFlowTcpCtlPsh
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpCtlPsh, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowTcpCtlPshChoiceEnum, set in PatternFlowTcpCtlPsh
	Choice() PatternFlowTcpCtlPshChoiceEnum
	// setChoice assigns PatternFlowTcpCtlPshChoiceEnum provided by user to PatternFlowTcpCtlPsh
	setChoice(value PatternFlowTcpCtlPshChoiceEnum) PatternFlowTcpCtlPsh
	// HasChoice checks if Choice has been set in PatternFlowTcpCtlPsh
	HasChoice() bool
	// Value returns uint32, set in PatternFlowTcpCtlPsh.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowTcpCtlPsh
	SetValue(value uint32) PatternFlowTcpCtlPsh
	// HasValue checks if Value has been set in PatternFlowTcpCtlPsh
	HasValue() bool
	// Values returns []uint32, set in PatternFlowTcpCtlPsh.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowTcpCtlPsh
	SetValues(value []uint32) PatternFlowTcpCtlPsh
	// Increment returns PatternFlowTcpCtlPshCounter, set in PatternFlowTcpCtlPsh.
	// PatternFlowTcpCtlPshCounter is integer counter pattern
	Increment() PatternFlowTcpCtlPshCounter
	// SetIncrement assigns PatternFlowTcpCtlPshCounter provided by user to PatternFlowTcpCtlPsh.
	// PatternFlowTcpCtlPshCounter is integer counter pattern
	SetIncrement(value PatternFlowTcpCtlPshCounter) PatternFlowTcpCtlPsh
	// HasIncrement checks if Increment has been set in PatternFlowTcpCtlPsh
	HasIncrement() bool
	// Decrement returns PatternFlowTcpCtlPshCounter, set in PatternFlowTcpCtlPsh.
	// PatternFlowTcpCtlPshCounter is integer counter pattern
	Decrement() PatternFlowTcpCtlPshCounter
	// SetDecrement assigns PatternFlowTcpCtlPshCounter provided by user to PatternFlowTcpCtlPsh.
	// PatternFlowTcpCtlPshCounter is integer counter pattern
	SetDecrement(value PatternFlowTcpCtlPshCounter) PatternFlowTcpCtlPsh
	// HasDecrement checks if Decrement has been set in PatternFlowTcpCtlPsh
	HasDecrement() bool
	// MetricTags returns PatternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIterIter, set in PatternFlowTcpCtlPsh
	MetricTags() PatternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter
	setNil()
}

type PatternFlowTcpCtlPshChoiceEnum string

// Enum of Choice on PatternFlowTcpCtlPsh
var PatternFlowTcpCtlPshChoice = struct {
	VALUE     PatternFlowTcpCtlPshChoiceEnum
	VALUES    PatternFlowTcpCtlPshChoiceEnum
	INCREMENT PatternFlowTcpCtlPshChoiceEnum
	DECREMENT PatternFlowTcpCtlPshChoiceEnum
}{
	VALUE:     PatternFlowTcpCtlPshChoiceEnum("value"),
	VALUES:    PatternFlowTcpCtlPshChoiceEnum("values"),
	INCREMENT: PatternFlowTcpCtlPshChoiceEnum("increment"),
	DECREMENT: PatternFlowTcpCtlPshChoiceEnum("decrement"),
}

func (obj *patternFlowTcpCtlPsh) Choice() PatternFlowTcpCtlPshChoiceEnum {
	return PatternFlowTcpCtlPshChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowTcpCtlPsh) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowTcpCtlPsh) setChoice(value PatternFlowTcpCtlPshChoiceEnum) PatternFlowTcpCtlPsh {
	intValue, ok := otg.PatternFlowTcpCtlPsh_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowTcpCtlPshChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowTcpCtlPsh_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowTcpCtlPshChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowTcpCtlPshChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowTcpCtlPshChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowTcpCtlPshCounter().msg()
	}

	if value == PatternFlowTcpCtlPshChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowTcpCtlPshCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpCtlPsh) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowTcpCtlPshChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpCtlPsh) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowTcpCtlPsh object
func (obj *patternFlowTcpCtlPsh) SetValue(value uint32) PatternFlowTcpCtlPsh {
	obj.setChoice(PatternFlowTcpCtlPshChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowTcpCtlPsh) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowTcpCtlPsh object
func (obj *patternFlowTcpCtlPsh) SetValues(value []uint32) PatternFlowTcpCtlPsh {
	obj.setChoice(PatternFlowTcpCtlPshChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowTcpCtlPshCounter
func (obj *patternFlowTcpCtlPsh) Increment() PatternFlowTcpCtlPshCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowTcpCtlPshChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowTcpCtlPshCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowTcpCtlPshCounter
func (obj *patternFlowTcpCtlPsh) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowTcpCtlPshCounter value in the PatternFlowTcpCtlPsh object
func (obj *patternFlowTcpCtlPsh) SetIncrement(value PatternFlowTcpCtlPshCounter) PatternFlowTcpCtlPsh {
	obj.setChoice(PatternFlowTcpCtlPshChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowTcpCtlPshCounter
func (obj *patternFlowTcpCtlPsh) Decrement() PatternFlowTcpCtlPshCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowTcpCtlPshChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowTcpCtlPshCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowTcpCtlPshCounter
func (obj *patternFlowTcpCtlPsh) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowTcpCtlPshCounter value in the PatternFlowTcpCtlPsh object
func (obj *patternFlowTcpCtlPsh) SetDecrement(value PatternFlowTcpCtlPshCounter) PatternFlowTcpCtlPsh {
	obj.setChoice(PatternFlowTcpCtlPshChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowTcpCtlPshMetricTag
func (obj *patternFlowTcpCtlPsh) MetricTags() PatternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowTcpCtlPshMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter struct {
	obj                                *patternFlowTcpCtlPsh
	patternFlowTcpCtlPshMetricTagSlice []PatternFlowTcpCtlPshMetricTag
	fieldPtr                           *[]*otg.PatternFlowTcpCtlPshMetricTag
}

func newPatternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter(ptr *[]*otg.PatternFlowTcpCtlPshMetricTag) PatternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter {
	return &patternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter{fieldPtr: ptr}
}

type PatternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter interface {
	setMsg(*patternFlowTcpCtlPsh) PatternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter
	Items() []PatternFlowTcpCtlPshMetricTag
	Add() PatternFlowTcpCtlPshMetricTag
	Append(items ...PatternFlowTcpCtlPshMetricTag) PatternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter
	Set(index int, newObj PatternFlowTcpCtlPshMetricTag) PatternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter
	Clear() PatternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter
	clearHolderSlice() PatternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter
	appendHolderSlice(item PatternFlowTcpCtlPshMetricTag) PatternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter
}

func (obj *patternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter) setMsg(msg *patternFlowTcpCtlPsh) PatternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowTcpCtlPshMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter) Items() []PatternFlowTcpCtlPshMetricTag {
	return obj.patternFlowTcpCtlPshMetricTagSlice
}

func (obj *patternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter) Add() PatternFlowTcpCtlPshMetricTag {
	newObj := &otg.PatternFlowTcpCtlPshMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowTcpCtlPshMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowTcpCtlPshMetricTagSlice = append(obj.patternFlowTcpCtlPshMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter) Append(items ...PatternFlowTcpCtlPshMetricTag) PatternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowTcpCtlPshMetricTagSlice = append(obj.patternFlowTcpCtlPshMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter) Set(index int, newObj PatternFlowTcpCtlPshMetricTag) PatternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowTcpCtlPshMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter) Clear() PatternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowTcpCtlPshMetricTag{}
		obj.patternFlowTcpCtlPshMetricTagSlice = []PatternFlowTcpCtlPshMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter) clearHolderSlice() PatternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter {
	if len(obj.patternFlowTcpCtlPshMetricTagSlice) > 0 {
		obj.patternFlowTcpCtlPshMetricTagSlice = []PatternFlowTcpCtlPshMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter) appendHolderSlice(item PatternFlowTcpCtlPshMetricTag) PatternFlowTcpCtlPshPatternFlowTcpCtlPshMetricTagIter {
	obj.patternFlowTcpCtlPshMetricTagSlice = append(obj.patternFlowTcpCtlPshMetricTagSlice, item)
	return obj
}

func (obj *patternFlowTcpCtlPsh) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlPsh.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowTcpCtlPsh.Values <= 1 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowTcpCtlPshMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowTcpCtlPsh) setDefault() {
	var choices_set int = 0
	var choice PatternFlowTcpCtlPshChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowTcpCtlPshChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowTcpCtlPshChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowTcpCtlPshChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowTcpCtlPshChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowTcpCtlPshChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowTcpCtlPsh")
			}
		} else {
			intVal := otg.PatternFlowTcpCtlPsh_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowTcpCtlPsh_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
