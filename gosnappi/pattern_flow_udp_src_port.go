package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowUdpSrcPort *****
type patternFlowUdpSrcPort struct {
	validation
	obj              *otg.PatternFlowUdpSrcPort
	marshaller       marshalPatternFlowUdpSrcPort
	unMarshaller     unMarshalPatternFlowUdpSrcPort
	incrementHolder  PatternFlowUdpSrcPortCounter
	decrementHolder  PatternFlowUdpSrcPortCounter
	metricTagsHolder PatternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter
}

func NewPatternFlowUdpSrcPort() PatternFlowUdpSrcPort {
	obj := patternFlowUdpSrcPort{obj: &otg.PatternFlowUdpSrcPort{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowUdpSrcPort) msg() *otg.PatternFlowUdpSrcPort {
	return obj.obj
}

func (obj *patternFlowUdpSrcPort) setMsg(msg *otg.PatternFlowUdpSrcPort) PatternFlowUdpSrcPort {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowUdpSrcPort struct {
	obj *patternFlowUdpSrcPort
}

type marshalPatternFlowUdpSrcPort interface {
	// ToProto marshals PatternFlowUdpSrcPort to protobuf object *otg.PatternFlowUdpSrcPort
	ToProto() (*otg.PatternFlowUdpSrcPort, error)
	// ToPbText marshals PatternFlowUdpSrcPort to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowUdpSrcPort to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowUdpSrcPort to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowUdpSrcPort struct {
	obj *patternFlowUdpSrcPort
}

type unMarshalPatternFlowUdpSrcPort interface {
	// FromProto unmarshals PatternFlowUdpSrcPort from protobuf object *otg.PatternFlowUdpSrcPort
	FromProto(msg *otg.PatternFlowUdpSrcPort) (PatternFlowUdpSrcPort, error)
	// FromPbText unmarshals PatternFlowUdpSrcPort from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowUdpSrcPort from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowUdpSrcPort from JSON text
	FromJson(value string) error
}

func (obj *patternFlowUdpSrcPort) Marshal() marshalPatternFlowUdpSrcPort {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowUdpSrcPort{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowUdpSrcPort) Unmarshal() unMarshalPatternFlowUdpSrcPort {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowUdpSrcPort{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowUdpSrcPort) ToProto() (*otg.PatternFlowUdpSrcPort, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowUdpSrcPort) FromProto(msg *otg.PatternFlowUdpSrcPort) (PatternFlowUdpSrcPort, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowUdpSrcPort) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowUdpSrcPort) FromPbText(value string) error {
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

func (m *marshalpatternFlowUdpSrcPort) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowUdpSrcPort) FromYaml(value string) error {
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

func (m *marshalpatternFlowUdpSrcPort) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowUdpSrcPort) FromJson(value string) error {
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

func (obj *patternFlowUdpSrcPort) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowUdpSrcPort) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowUdpSrcPort) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowUdpSrcPort) Clone() (PatternFlowUdpSrcPort, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowUdpSrcPort()
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

func (obj *patternFlowUdpSrcPort) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowUdpSrcPort is source port
type PatternFlowUdpSrcPort interface {
	Validation
	// msg marshals PatternFlowUdpSrcPort to protobuf object *otg.PatternFlowUdpSrcPort
	// and doesn't set defaults
	msg() *otg.PatternFlowUdpSrcPort
	// setMsg unmarshals PatternFlowUdpSrcPort from protobuf object *otg.PatternFlowUdpSrcPort
	// and doesn't set defaults
	setMsg(*otg.PatternFlowUdpSrcPort) PatternFlowUdpSrcPort
	// provides marshal interface
	Marshal() marshalPatternFlowUdpSrcPort
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowUdpSrcPort
	// validate validates PatternFlowUdpSrcPort
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowUdpSrcPort, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowUdpSrcPortChoiceEnum, set in PatternFlowUdpSrcPort
	Choice() PatternFlowUdpSrcPortChoiceEnum
	// setChoice assigns PatternFlowUdpSrcPortChoiceEnum provided by user to PatternFlowUdpSrcPort
	setChoice(value PatternFlowUdpSrcPortChoiceEnum) PatternFlowUdpSrcPort
	// HasChoice checks if Choice has been set in PatternFlowUdpSrcPort
	HasChoice() bool
	// Value returns uint32, set in PatternFlowUdpSrcPort.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowUdpSrcPort
	SetValue(value uint32) PatternFlowUdpSrcPort
	// HasValue checks if Value has been set in PatternFlowUdpSrcPort
	HasValue() bool
	// Values returns []uint32, set in PatternFlowUdpSrcPort.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowUdpSrcPort
	SetValues(value []uint32) PatternFlowUdpSrcPort
	// Increment returns PatternFlowUdpSrcPortCounter, set in PatternFlowUdpSrcPort.
	// PatternFlowUdpSrcPortCounter is integer counter pattern
	Increment() PatternFlowUdpSrcPortCounter
	// SetIncrement assigns PatternFlowUdpSrcPortCounter provided by user to PatternFlowUdpSrcPort.
	// PatternFlowUdpSrcPortCounter is integer counter pattern
	SetIncrement(value PatternFlowUdpSrcPortCounter) PatternFlowUdpSrcPort
	// HasIncrement checks if Increment has been set in PatternFlowUdpSrcPort
	HasIncrement() bool
	// Decrement returns PatternFlowUdpSrcPortCounter, set in PatternFlowUdpSrcPort.
	// PatternFlowUdpSrcPortCounter is integer counter pattern
	Decrement() PatternFlowUdpSrcPortCounter
	// SetDecrement assigns PatternFlowUdpSrcPortCounter provided by user to PatternFlowUdpSrcPort.
	// PatternFlowUdpSrcPortCounter is integer counter pattern
	SetDecrement(value PatternFlowUdpSrcPortCounter) PatternFlowUdpSrcPort
	// HasDecrement checks if Decrement has been set in PatternFlowUdpSrcPort
	HasDecrement() bool
	// MetricTags returns PatternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIterIter, set in PatternFlowUdpSrcPort
	MetricTags() PatternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter
	setNil()
}

type PatternFlowUdpSrcPortChoiceEnum string

// Enum of Choice on PatternFlowUdpSrcPort
var PatternFlowUdpSrcPortChoice = struct {
	VALUE     PatternFlowUdpSrcPortChoiceEnum
	VALUES    PatternFlowUdpSrcPortChoiceEnum
	INCREMENT PatternFlowUdpSrcPortChoiceEnum
	DECREMENT PatternFlowUdpSrcPortChoiceEnum
}{
	VALUE:     PatternFlowUdpSrcPortChoiceEnum("value"),
	VALUES:    PatternFlowUdpSrcPortChoiceEnum("values"),
	INCREMENT: PatternFlowUdpSrcPortChoiceEnum("increment"),
	DECREMENT: PatternFlowUdpSrcPortChoiceEnum("decrement"),
}

func (obj *patternFlowUdpSrcPort) Choice() PatternFlowUdpSrcPortChoiceEnum {
	return PatternFlowUdpSrcPortChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowUdpSrcPort) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowUdpSrcPort) setChoice(value PatternFlowUdpSrcPortChoiceEnum) PatternFlowUdpSrcPort {
	intValue, ok := otg.PatternFlowUdpSrcPort_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowUdpSrcPortChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowUdpSrcPort_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowUdpSrcPortChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowUdpSrcPortChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowUdpSrcPortChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowUdpSrcPortCounter().msg()
	}

	if value == PatternFlowUdpSrcPortChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowUdpSrcPortCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowUdpSrcPort) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowUdpSrcPortChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowUdpSrcPort) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowUdpSrcPort object
func (obj *patternFlowUdpSrcPort) SetValue(value uint32) PatternFlowUdpSrcPort {
	obj.setChoice(PatternFlowUdpSrcPortChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowUdpSrcPort) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowUdpSrcPort object
func (obj *patternFlowUdpSrcPort) SetValues(value []uint32) PatternFlowUdpSrcPort {
	obj.setChoice(PatternFlowUdpSrcPortChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowUdpSrcPortCounter
func (obj *patternFlowUdpSrcPort) Increment() PatternFlowUdpSrcPortCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowUdpSrcPortChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowUdpSrcPortCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowUdpSrcPortCounter
func (obj *patternFlowUdpSrcPort) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowUdpSrcPortCounter value in the PatternFlowUdpSrcPort object
func (obj *patternFlowUdpSrcPort) SetIncrement(value PatternFlowUdpSrcPortCounter) PatternFlowUdpSrcPort {
	obj.setChoice(PatternFlowUdpSrcPortChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowUdpSrcPortCounter
func (obj *patternFlowUdpSrcPort) Decrement() PatternFlowUdpSrcPortCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowUdpSrcPortChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowUdpSrcPortCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowUdpSrcPortCounter
func (obj *patternFlowUdpSrcPort) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowUdpSrcPortCounter value in the PatternFlowUdpSrcPort object
func (obj *patternFlowUdpSrcPort) SetDecrement(value PatternFlowUdpSrcPortCounter) PatternFlowUdpSrcPort {
	obj.setChoice(PatternFlowUdpSrcPortChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowUdpSrcPortMetricTag
func (obj *patternFlowUdpSrcPort) MetricTags() PatternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowUdpSrcPortMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter struct {
	obj                                 *patternFlowUdpSrcPort
	patternFlowUdpSrcPortMetricTagSlice []PatternFlowUdpSrcPortMetricTag
	fieldPtr                            *[]*otg.PatternFlowUdpSrcPortMetricTag
}

func newPatternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter(ptr *[]*otg.PatternFlowUdpSrcPortMetricTag) PatternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter {
	return &patternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter{fieldPtr: ptr}
}

type PatternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter interface {
	setMsg(*patternFlowUdpSrcPort) PatternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter
	Items() []PatternFlowUdpSrcPortMetricTag
	Add() PatternFlowUdpSrcPortMetricTag
	Append(items ...PatternFlowUdpSrcPortMetricTag) PatternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter
	Set(index int, newObj PatternFlowUdpSrcPortMetricTag) PatternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter
	Clear() PatternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter
	clearHolderSlice() PatternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter
	appendHolderSlice(item PatternFlowUdpSrcPortMetricTag) PatternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter
}

func (obj *patternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter) setMsg(msg *patternFlowUdpSrcPort) PatternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowUdpSrcPortMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter) Items() []PatternFlowUdpSrcPortMetricTag {
	return obj.patternFlowUdpSrcPortMetricTagSlice
}

func (obj *patternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter) Add() PatternFlowUdpSrcPortMetricTag {
	newObj := &otg.PatternFlowUdpSrcPortMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowUdpSrcPortMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowUdpSrcPortMetricTagSlice = append(obj.patternFlowUdpSrcPortMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter) Append(items ...PatternFlowUdpSrcPortMetricTag) PatternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowUdpSrcPortMetricTagSlice = append(obj.patternFlowUdpSrcPortMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter) Set(index int, newObj PatternFlowUdpSrcPortMetricTag) PatternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowUdpSrcPortMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter) Clear() PatternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowUdpSrcPortMetricTag{}
		obj.patternFlowUdpSrcPortMetricTagSlice = []PatternFlowUdpSrcPortMetricTag{}
	}
	return obj
}
func (obj *patternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter) clearHolderSlice() PatternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter {
	if len(obj.patternFlowUdpSrcPortMetricTagSlice) > 0 {
		obj.patternFlowUdpSrcPortMetricTagSlice = []PatternFlowUdpSrcPortMetricTag{}
	}
	return obj
}
func (obj *patternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter) appendHolderSlice(item PatternFlowUdpSrcPortMetricTag) PatternFlowUdpSrcPortPatternFlowUdpSrcPortMetricTagIter {
	obj.patternFlowUdpSrcPortMetricTagSlice = append(obj.patternFlowUdpSrcPortMetricTagSlice, item)
	return obj
}

func (obj *patternFlowUdpSrcPort) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowUdpSrcPort.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowUdpSrcPort.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowUdpSrcPortMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowUdpSrcPort) setDefault() {
	var choices_set int = 0
	var choice PatternFlowUdpSrcPortChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowUdpSrcPortChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowUdpSrcPortChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowUdpSrcPortChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowUdpSrcPortChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowUdpSrcPortChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowUdpSrcPort")
			}
		} else {
			intVal := otg.PatternFlowUdpSrcPort_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowUdpSrcPort_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
