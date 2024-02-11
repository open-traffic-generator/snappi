package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpCtlAck *****
type patternFlowTcpCtlAck struct {
	validation
	obj              *otg.PatternFlowTcpCtlAck
	marshaller       marshalPatternFlowTcpCtlAck
	unMarshaller     unMarshalPatternFlowTcpCtlAck
	incrementHolder  PatternFlowTcpCtlAckCounter
	decrementHolder  PatternFlowTcpCtlAckCounter
	metricTagsHolder PatternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter
}

func NewPatternFlowTcpCtlAck() PatternFlowTcpCtlAck {
	obj := patternFlowTcpCtlAck{obj: &otg.PatternFlowTcpCtlAck{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpCtlAck) msg() *otg.PatternFlowTcpCtlAck {
	return obj.obj
}

func (obj *patternFlowTcpCtlAck) setMsg(msg *otg.PatternFlowTcpCtlAck) PatternFlowTcpCtlAck {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpCtlAck struct {
	obj *patternFlowTcpCtlAck
}

type marshalPatternFlowTcpCtlAck interface {
	// ToProto marshals PatternFlowTcpCtlAck to protobuf object *otg.PatternFlowTcpCtlAck
	ToProto() (*otg.PatternFlowTcpCtlAck, error)
	// ToPbText marshals PatternFlowTcpCtlAck to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpCtlAck to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpCtlAck to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpCtlAck struct {
	obj *patternFlowTcpCtlAck
}

type unMarshalPatternFlowTcpCtlAck interface {
	// FromProto unmarshals PatternFlowTcpCtlAck from protobuf object *otg.PatternFlowTcpCtlAck
	FromProto(msg *otg.PatternFlowTcpCtlAck) (PatternFlowTcpCtlAck, error)
	// FromPbText unmarshals PatternFlowTcpCtlAck from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpCtlAck from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpCtlAck from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpCtlAck) Marshal() marshalPatternFlowTcpCtlAck {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpCtlAck{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpCtlAck) Unmarshal() unMarshalPatternFlowTcpCtlAck {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpCtlAck{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpCtlAck) ToProto() (*otg.PatternFlowTcpCtlAck, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpCtlAck) FromProto(msg *otg.PatternFlowTcpCtlAck) (PatternFlowTcpCtlAck, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpCtlAck) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlAck) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpCtlAck) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlAck) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpCtlAck) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlAck) FromJson(value string) error {
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

func (obj *patternFlowTcpCtlAck) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlAck) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlAck) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpCtlAck) Clone() (PatternFlowTcpCtlAck, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpCtlAck()
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

func (obj *patternFlowTcpCtlAck) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowTcpCtlAck is a value of 1 indicates that the ackknowledgment field is significant.
type PatternFlowTcpCtlAck interface {
	Validation
	// msg marshals PatternFlowTcpCtlAck to protobuf object *otg.PatternFlowTcpCtlAck
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpCtlAck
	// setMsg unmarshals PatternFlowTcpCtlAck from protobuf object *otg.PatternFlowTcpCtlAck
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpCtlAck) PatternFlowTcpCtlAck
	// provides marshal interface
	Marshal() marshalPatternFlowTcpCtlAck
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpCtlAck
	// validate validates PatternFlowTcpCtlAck
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpCtlAck, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowTcpCtlAckChoiceEnum, set in PatternFlowTcpCtlAck
	Choice() PatternFlowTcpCtlAckChoiceEnum
	// setChoice assigns PatternFlowTcpCtlAckChoiceEnum provided by user to PatternFlowTcpCtlAck
	setChoice(value PatternFlowTcpCtlAckChoiceEnum) PatternFlowTcpCtlAck
	// HasChoice checks if Choice has been set in PatternFlowTcpCtlAck
	HasChoice() bool
	// Value returns uint32, set in PatternFlowTcpCtlAck.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowTcpCtlAck
	SetValue(value uint32) PatternFlowTcpCtlAck
	// HasValue checks if Value has been set in PatternFlowTcpCtlAck
	HasValue() bool
	// Values returns []uint32, set in PatternFlowTcpCtlAck.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowTcpCtlAck
	SetValues(value []uint32) PatternFlowTcpCtlAck
	// Increment returns PatternFlowTcpCtlAckCounter, set in PatternFlowTcpCtlAck.
	// PatternFlowTcpCtlAckCounter is integer counter pattern
	Increment() PatternFlowTcpCtlAckCounter
	// SetIncrement assigns PatternFlowTcpCtlAckCounter provided by user to PatternFlowTcpCtlAck.
	// PatternFlowTcpCtlAckCounter is integer counter pattern
	SetIncrement(value PatternFlowTcpCtlAckCounter) PatternFlowTcpCtlAck
	// HasIncrement checks if Increment has been set in PatternFlowTcpCtlAck
	HasIncrement() bool
	// Decrement returns PatternFlowTcpCtlAckCounter, set in PatternFlowTcpCtlAck.
	// PatternFlowTcpCtlAckCounter is integer counter pattern
	Decrement() PatternFlowTcpCtlAckCounter
	// SetDecrement assigns PatternFlowTcpCtlAckCounter provided by user to PatternFlowTcpCtlAck.
	// PatternFlowTcpCtlAckCounter is integer counter pattern
	SetDecrement(value PatternFlowTcpCtlAckCounter) PatternFlowTcpCtlAck
	// HasDecrement checks if Decrement has been set in PatternFlowTcpCtlAck
	HasDecrement() bool
	// MetricTags returns PatternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIterIter, set in PatternFlowTcpCtlAck
	MetricTags() PatternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter
	setNil()
}

type PatternFlowTcpCtlAckChoiceEnum string

// Enum of Choice on PatternFlowTcpCtlAck
var PatternFlowTcpCtlAckChoice = struct {
	VALUE     PatternFlowTcpCtlAckChoiceEnum
	VALUES    PatternFlowTcpCtlAckChoiceEnum
	INCREMENT PatternFlowTcpCtlAckChoiceEnum
	DECREMENT PatternFlowTcpCtlAckChoiceEnum
}{
	VALUE:     PatternFlowTcpCtlAckChoiceEnum("value"),
	VALUES:    PatternFlowTcpCtlAckChoiceEnum("values"),
	INCREMENT: PatternFlowTcpCtlAckChoiceEnum("increment"),
	DECREMENT: PatternFlowTcpCtlAckChoiceEnum("decrement"),
}

func (obj *patternFlowTcpCtlAck) Choice() PatternFlowTcpCtlAckChoiceEnum {
	return PatternFlowTcpCtlAckChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowTcpCtlAck) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowTcpCtlAck) setChoice(value PatternFlowTcpCtlAckChoiceEnum) PatternFlowTcpCtlAck {
	intValue, ok := otg.PatternFlowTcpCtlAck_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowTcpCtlAckChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowTcpCtlAck_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowTcpCtlAckChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowTcpCtlAckChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowTcpCtlAckChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowTcpCtlAckCounter().msg()
	}

	if value == PatternFlowTcpCtlAckChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowTcpCtlAckCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpCtlAck) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowTcpCtlAckChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpCtlAck) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowTcpCtlAck object
func (obj *patternFlowTcpCtlAck) SetValue(value uint32) PatternFlowTcpCtlAck {
	obj.setChoice(PatternFlowTcpCtlAckChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowTcpCtlAck) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowTcpCtlAck object
func (obj *patternFlowTcpCtlAck) SetValues(value []uint32) PatternFlowTcpCtlAck {
	obj.setChoice(PatternFlowTcpCtlAckChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowTcpCtlAckCounter
func (obj *patternFlowTcpCtlAck) Increment() PatternFlowTcpCtlAckCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowTcpCtlAckChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowTcpCtlAckCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowTcpCtlAckCounter
func (obj *patternFlowTcpCtlAck) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowTcpCtlAckCounter value in the PatternFlowTcpCtlAck object
func (obj *patternFlowTcpCtlAck) SetIncrement(value PatternFlowTcpCtlAckCounter) PatternFlowTcpCtlAck {
	obj.setChoice(PatternFlowTcpCtlAckChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowTcpCtlAckCounter
func (obj *patternFlowTcpCtlAck) Decrement() PatternFlowTcpCtlAckCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowTcpCtlAckChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowTcpCtlAckCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowTcpCtlAckCounter
func (obj *patternFlowTcpCtlAck) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowTcpCtlAckCounter value in the PatternFlowTcpCtlAck object
func (obj *patternFlowTcpCtlAck) SetDecrement(value PatternFlowTcpCtlAckCounter) PatternFlowTcpCtlAck {
	obj.setChoice(PatternFlowTcpCtlAckChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowTcpCtlAckMetricTag
func (obj *patternFlowTcpCtlAck) MetricTags() PatternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowTcpCtlAckMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter struct {
	obj                                *patternFlowTcpCtlAck
	patternFlowTcpCtlAckMetricTagSlice []PatternFlowTcpCtlAckMetricTag
	fieldPtr                           *[]*otg.PatternFlowTcpCtlAckMetricTag
}

func newPatternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter(ptr *[]*otg.PatternFlowTcpCtlAckMetricTag) PatternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter {
	return &patternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter{fieldPtr: ptr}
}

type PatternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter interface {
	setMsg(*patternFlowTcpCtlAck) PatternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter
	Items() []PatternFlowTcpCtlAckMetricTag
	Add() PatternFlowTcpCtlAckMetricTag
	Append(items ...PatternFlowTcpCtlAckMetricTag) PatternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter
	Set(index int, newObj PatternFlowTcpCtlAckMetricTag) PatternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter
	Clear() PatternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter
	clearHolderSlice() PatternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter
	appendHolderSlice(item PatternFlowTcpCtlAckMetricTag) PatternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter
}

func (obj *patternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter) setMsg(msg *patternFlowTcpCtlAck) PatternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowTcpCtlAckMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter) Items() []PatternFlowTcpCtlAckMetricTag {
	return obj.patternFlowTcpCtlAckMetricTagSlice
}

func (obj *patternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter) Add() PatternFlowTcpCtlAckMetricTag {
	newObj := &otg.PatternFlowTcpCtlAckMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowTcpCtlAckMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowTcpCtlAckMetricTagSlice = append(obj.patternFlowTcpCtlAckMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter) Append(items ...PatternFlowTcpCtlAckMetricTag) PatternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowTcpCtlAckMetricTagSlice = append(obj.patternFlowTcpCtlAckMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter) Set(index int, newObj PatternFlowTcpCtlAckMetricTag) PatternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowTcpCtlAckMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter) Clear() PatternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowTcpCtlAckMetricTag{}
		obj.patternFlowTcpCtlAckMetricTagSlice = []PatternFlowTcpCtlAckMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter) clearHolderSlice() PatternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter {
	if len(obj.patternFlowTcpCtlAckMetricTagSlice) > 0 {
		obj.patternFlowTcpCtlAckMetricTagSlice = []PatternFlowTcpCtlAckMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter) appendHolderSlice(item PatternFlowTcpCtlAckMetricTag) PatternFlowTcpCtlAckPatternFlowTcpCtlAckMetricTagIter {
	obj.patternFlowTcpCtlAckMetricTagSlice = append(obj.patternFlowTcpCtlAckMetricTagSlice, item)
	return obj
}

func (obj *patternFlowTcpCtlAck) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlAck.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowTcpCtlAck.Values <= 1 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowTcpCtlAckMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowTcpCtlAck) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowTcpCtlAckChoice.VALUE)

	}

}
