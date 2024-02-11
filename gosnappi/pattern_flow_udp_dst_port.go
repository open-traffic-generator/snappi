package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowUdpDstPort *****
type patternFlowUdpDstPort struct {
	validation
	obj              *otg.PatternFlowUdpDstPort
	marshaller       marshalPatternFlowUdpDstPort
	unMarshaller     unMarshalPatternFlowUdpDstPort
	incrementHolder  PatternFlowUdpDstPortCounter
	decrementHolder  PatternFlowUdpDstPortCounter
	metricTagsHolder PatternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter
}

func NewPatternFlowUdpDstPort() PatternFlowUdpDstPort {
	obj := patternFlowUdpDstPort{obj: &otg.PatternFlowUdpDstPort{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowUdpDstPort) msg() *otg.PatternFlowUdpDstPort {
	return obj.obj
}

func (obj *patternFlowUdpDstPort) setMsg(msg *otg.PatternFlowUdpDstPort) PatternFlowUdpDstPort {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowUdpDstPort struct {
	obj *patternFlowUdpDstPort
}

type marshalPatternFlowUdpDstPort interface {
	// ToProto marshals PatternFlowUdpDstPort to protobuf object *otg.PatternFlowUdpDstPort
	ToProto() (*otg.PatternFlowUdpDstPort, error)
	// ToPbText marshals PatternFlowUdpDstPort to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowUdpDstPort to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowUdpDstPort to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowUdpDstPort struct {
	obj *patternFlowUdpDstPort
}

type unMarshalPatternFlowUdpDstPort interface {
	// FromProto unmarshals PatternFlowUdpDstPort from protobuf object *otg.PatternFlowUdpDstPort
	FromProto(msg *otg.PatternFlowUdpDstPort) (PatternFlowUdpDstPort, error)
	// FromPbText unmarshals PatternFlowUdpDstPort from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowUdpDstPort from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowUdpDstPort from JSON text
	FromJson(value string) error
}

func (obj *patternFlowUdpDstPort) Marshal() marshalPatternFlowUdpDstPort {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowUdpDstPort{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowUdpDstPort) Unmarshal() unMarshalPatternFlowUdpDstPort {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowUdpDstPort{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowUdpDstPort) ToProto() (*otg.PatternFlowUdpDstPort, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowUdpDstPort) FromProto(msg *otg.PatternFlowUdpDstPort) (PatternFlowUdpDstPort, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowUdpDstPort) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowUdpDstPort) FromPbText(value string) error {
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

func (m *marshalpatternFlowUdpDstPort) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowUdpDstPort) FromYaml(value string) error {
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

func (m *marshalpatternFlowUdpDstPort) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowUdpDstPort) FromJson(value string) error {
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

func (obj *patternFlowUdpDstPort) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowUdpDstPort) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowUdpDstPort) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowUdpDstPort) Clone() (PatternFlowUdpDstPort, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowUdpDstPort()
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

func (obj *patternFlowUdpDstPort) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowUdpDstPort is destination port
type PatternFlowUdpDstPort interface {
	Validation
	// msg marshals PatternFlowUdpDstPort to protobuf object *otg.PatternFlowUdpDstPort
	// and doesn't set defaults
	msg() *otg.PatternFlowUdpDstPort
	// setMsg unmarshals PatternFlowUdpDstPort from protobuf object *otg.PatternFlowUdpDstPort
	// and doesn't set defaults
	setMsg(*otg.PatternFlowUdpDstPort) PatternFlowUdpDstPort
	// provides marshal interface
	Marshal() marshalPatternFlowUdpDstPort
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowUdpDstPort
	// validate validates PatternFlowUdpDstPort
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowUdpDstPort, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowUdpDstPortChoiceEnum, set in PatternFlowUdpDstPort
	Choice() PatternFlowUdpDstPortChoiceEnum
	// setChoice assigns PatternFlowUdpDstPortChoiceEnum provided by user to PatternFlowUdpDstPort
	setChoice(value PatternFlowUdpDstPortChoiceEnum) PatternFlowUdpDstPort
	// HasChoice checks if Choice has been set in PatternFlowUdpDstPort
	HasChoice() bool
	// Value returns uint32, set in PatternFlowUdpDstPort.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowUdpDstPort
	SetValue(value uint32) PatternFlowUdpDstPort
	// HasValue checks if Value has been set in PatternFlowUdpDstPort
	HasValue() bool
	// Values returns []uint32, set in PatternFlowUdpDstPort.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowUdpDstPort
	SetValues(value []uint32) PatternFlowUdpDstPort
	// Increment returns PatternFlowUdpDstPortCounter, set in PatternFlowUdpDstPort.
	// PatternFlowUdpDstPortCounter is integer counter pattern
	Increment() PatternFlowUdpDstPortCounter
	// SetIncrement assigns PatternFlowUdpDstPortCounter provided by user to PatternFlowUdpDstPort.
	// PatternFlowUdpDstPortCounter is integer counter pattern
	SetIncrement(value PatternFlowUdpDstPortCounter) PatternFlowUdpDstPort
	// HasIncrement checks if Increment has been set in PatternFlowUdpDstPort
	HasIncrement() bool
	// Decrement returns PatternFlowUdpDstPortCounter, set in PatternFlowUdpDstPort.
	// PatternFlowUdpDstPortCounter is integer counter pattern
	Decrement() PatternFlowUdpDstPortCounter
	// SetDecrement assigns PatternFlowUdpDstPortCounter provided by user to PatternFlowUdpDstPort.
	// PatternFlowUdpDstPortCounter is integer counter pattern
	SetDecrement(value PatternFlowUdpDstPortCounter) PatternFlowUdpDstPort
	// HasDecrement checks if Decrement has been set in PatternFlowUdpDstPort
	HasDecrement() bool
	// MetricTags returns PatternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIterIter, set in PatternFlowUdpDstPort
	MetricTags() PatternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter
	setNil()
}

type PatternFlowUdpDstPortChoiceEnum string

// Enum of Choice on PatternFlowUdpDstPort
var PatternFlowUdpDstPortChoice = struct {
	VALUE     PatternFlowUdpDstPortChoiceEnum
	VALUES    PatternFlowUdpDstPortChoiceEnum
	INCREMENT PatternFlowUdpDstPortChoiceEnum
	DECREMENT PatternFlowUdpDstPortChoiceEnum
}{
	VALUE:     PatternFlowUdpDstPortChoiceEnum("value"),
	VALUES:    PatternFlowUdpDstPortChoiceEnum("values"),
	INCREMENT: PatternFlowUdpDstPortChoiceEnum("increment"),
	DECREMENT: PatternFlowUdpDstPortChoiceEnum("decrement"),
}

func (obj *patternFlowUdpDstPort) Choice() PatternFlowUdpDstPortChoiceEnum {
	return PatternFlowUdpDstPortChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowUdpDstPort) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowUdpDstPort) setChoice(value PatternFlowUdpDstPortChoiceEnum) PatternFlowUdpDstPort {
	intValue, ok := otg.PatternFlowUdpDstPort_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowUdpDstPortChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowUdpDstPort_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowUdpDstPortChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowUdpDstPortChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowUdpDstPortChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowUdpDstPortCounter().msg()
	}

	if value == PatternFlowUdpDstPortChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowUdpDstPortCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowUdpDstPort) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowUdpDstPortChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowUdpDstPort) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowUdpDstPort object
func (obj *patternFlowUdpDstPort) SetValue(value uint32) PatternFlowUdpDstPort {
	obj.setChoice(PatternFlowUdpDstPortChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowUdpDstPort) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowUdpDstPort object
func (obj *patternFlowUdpDstPort) SetValues(value []uint32) PatternFlowUdpDstPort {
	obj.setChoice(PatternFlowUdpDstPortChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowUdpDstPortCounter
func (obj *patternFlowUdpDstPort) Increment() PatternFlowUdpDstPortCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowUdpDstPortChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowUdpDstPortCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowUdpDstPortCounter
func (obj *patternFlowUdpDstPort) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowUdpDstPortCounter value in the PatternFlowUdpDstPort object
func (obj *patternFlowUdpDstPort) SetIncrement(value PatternFlowUdpDstPortCounter) PatternFlowUdpDstPort {
	obj.setChoice(PatternFlowUdpDstPortChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowUdpDstPortCounter
func (obj *patternFlowUdpDstPort) Decrement() PatternFlowUdpDstPortCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowUdpDstPortChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowUdpDstPortCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowUdpDstPortCounter
func (obj *patternFlowUdpDstPort) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowUdpDstPortCounter value in the PatternFlowUdpDstPort object
func (obj *patternFlowUdpDstPort) SetDecrement(value PatternFlowUdpDstPortCounter) PatternFlowUdpDstPort {
	obj.setChoice(PatternFlowUdpDstPortChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowUdpDstPortMetricTag
func (obj *patternFlowUdpDstPort) MetricTags() PatternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowUdpDstPortMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter struct {
	obj                                 *patternFlowUdpDstPort
	patternFlowUdpDstPortMetricTagSlice []PatternFlowUdpDstPortMetricTag
	fieldPtr                            *[]*otg.PatternFlowUdpDstPortMetricTag
}

func newPatternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter(ptr *[]*otg.PatternFlowUdpDstPortMetricTag) PatternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter {
	return &patternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter{fieldPtr: ptr}
}

type PatternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter interface {
	setMsg(*patternFlowUdpDstPort) PatternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter
	Items() []PatternFlowUdpDstPortMetricTag
	Add() PatternFlowUdpDstPortMetricTag
	Append(items ...PatternFlowUdpDstPortMetricTag) PatternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter
	Set(index int, newObj PatternFlowUdpDstPortMetricTag) PatternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter
	Clear() PatternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter
	clearHolderSlice() PatternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter
	appendHolderSlice(item PatternFlowUdpDstPortMetricTag) PatternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter
}

func (obj *patternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter) setMsg(msg *patternFlowUdpDstPort) PatternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowUdpDstPortMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter) Items() []PatternFlowUdpDstPortMetricTag {
	return obj.patternFlowUdpDstPortMetricTagSlice
}

func (obj *patternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter) Add() PatternFlowUdpDstPortMetricTag {
	newObj := &otg.PatternFlowUdpDstPortMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowUdpDstPortMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowUdpDstPortMetricTagSlice = append(obj.patternFlowUdpDstPortMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter) Append(items ...PatternFlowUdpDstPortMetricTag) PatternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowUdpDstPortMetricTagSlice = append(obj.patternFlowUdpDstPortMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter) Set(index int, newObj PatternFlowUdpDstPortMetricTag) PatternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowUdpDstPortMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter) Clear() PatternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowUdpDstPortMetricTag{}
		obj.patternFlowUdpDstPortMetricTagSlice = []PatternFlowUdpDstPortMetricTag{}
	}
	return obj
}
func (obj *patternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter) clearHolderSlice() PatternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter {
	if len(obj.patternFlowUdpDstPortMetricTagSlice) > 0 {
		obj.patternFlowUdpDstPortMetricTagSlice = []PatternFlowUdpDstPortMetricTag{}
	}
	return obj
}
func (obj *patternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter) appendHolderSlice(item PatternFlowUdpDstPortMetricTag) PatternFlowUdpDstPortPatternFlowUdpDstPortMetricTagIter {
	obj.patternFlowUdpDstPortMetricTagSlice = append(obj.patternFlowUdpDstPortMetricTagSlice, item)
	return obj
}

func (obj *patternFlowUdpDstPort) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowUdpDstPort.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowUdpDstPort.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowUdpDstPortMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowUdpDstPort) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowUdpDstPortChoice.VALUE)

	}

}
