package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpDstPort *****
type patternFlowTcpDstPort struct {
	validation
	obj              *otg.PatternFlowTcpDstPort
	marshaller       marshalPatternFlowTcpDstPort
	unMarshaller     unMarshalPatternFlowTcpDstPort
	incrementHolder  PatternFlowTcpDstPortCounter
	decrementHolder  PatternFlowTcpDstPortCounter
	metricTagsHolder PatternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter
}

func NewPatternFlowTcpDstPort() PatternFlowTcpDstPort {
	obj := patternFlowTcpDstPort{obj: &otg.PatternFlowTcpDstPort{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpDstPort) msg() *otg.PatternFlowTcpDstPort {
	return obj.obj
}

func (obj *patternFlowTcpDstPort) setMsg(msg *otg.PatternFlowTcpDstPort) PatternFlowTcpDstPort {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpDstPort struct {
	obj *patternFlowTcpDstPort
}

type marshalPatternFlowTcpDstPort interface {
	// ToProto marshals PatternFlowTcpDstPort to protobuf object *otg.PatternFlowTcpDstPort
	ToProto() (*otg.PatternFlowTcpDstPort, error)
	// ToPbText marshals PatternFlowTcpDstPort to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpDstPort to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpDstPort to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpDstPort struct {
	obj *patternFlowTcpDstPort
}

type unMarshalPatternFlowTcpDstPort interface {
	// FromProto unmarshals PatternFlowTcpDstPort from protobuf object *otg.PatternFlowTcpDstPort
	FromProto(msg *otg.PatternFlowTcpDstPort) (PatternFlowTcpDstPort, error)
	// FromPbText unmarshals PatternFlowTcpDstPort from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpDstPort from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpDstPort from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpDstPort) Marshal() marshalPatternFlowTcpDstPort {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpDstPort{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpDstPort) Unmarshal() unMarshalPatternFlowTcpDstPort {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpDstPort{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpDstPort) ToProto() (*otg.PatternFlowTcpDstPort, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpDstPort) FromProto(msg *otg.PatternFlowTcpDstPort) (PatternFlowTcpDstPort, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpDstPort) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpDstPort) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpDstPort) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpDstPort) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpDstPort) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpDstPort) FromJson(value string) error {
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

func (obj *patternFlowTcpDstPort) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpDstPort) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpDstPort) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpDstPort) Clone() (PatternFlowTcpDstPort, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpDstPort()
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

func (obj *patternFlowTcpDstPort) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowTcpDstPort is destination port
type PatternFlowTcpDstPort interface {
	Validation
	// msg marshals PatternFlowTcpDstPort to protobuf object *otg.PatternFlowTcpDstPort
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpDstPort
	// setMsg unmarshals PatternFlowTcpDstPort from protobuf object *otg.PatternFlowTcpDstPort
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpDstPort) PatternFlowTcpDstPort
	// provides marshal interface
	Marshal() marshalPatternFlowTcpDstPort
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpDstPort
	// validate validates PatternFlowTcpDstPort
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpDstPort, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowTcpDstPortChoiceEnum, set in PatternFlowTcpDstPort
	Choice() PatternFlowTcpDstPortChoiceEnum
	// setChoice assigns PatternFlowTcpDstPortChoiceEnum provided by user to PatternFlowTcpDstPort
	setChoice(value PatternFlowTcpDstPortChoiceEnum) PatternFlowTcpDstPort
	// HasChoice checks if Choice has been set in PatternFlowTcpDstPort
	HasChoice() bool
	// Value returns uint32, set in PatternFlowTcpDstPort.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowTcpDstPort
	SetValue(value uint32) PatternFlowTcpDstPort
	// HasValue checks if Value has been set in PatternFlowTcpDstPort
	HasValue() bool
	// Values returns []uint32, set in PatternFlowTcpDstPort.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowTcpDstPort
	SetValues(value []uint32) PatternFlowTcpDstPort
	// Increment returns PatternFlowTcpDstPortCounter, set in PatternFlowTcpDstPort.
	// PatternFlowTcpDstPortCounter is integer counter pattern
	Increment() PatternFlowTcpDstPortCounter
	// SetIncrement assigns PatternFlowTcpDstPortCounter provided by user to PatternFlowTcpDstPort.
	// PatternFlowTcpDstPortCounter is integer counter pattern
	SetIncrement(value PatternFlowTcpDstPortCounter) PatternFlowTcpDstPort
	// HasIncrement checks if Increment has been set in PatternFlowTcpDstPort
	HasIncrement() bool
	// Decrement returns PatternFlowTcpDstPortCounter, set in PatternFlowTcpDstPort.
	// PatternFlowTcpDstPortCounter is integer counter pattern
	Decrement() PatternFlowTcpDstPortCounter
	// SetDecrement assigns PatternFlowTcpDstPortCounter provided by user to PatternFlowTcpDstPort.
	// PatternFlowTcpDstPortCounter is integer counter pattern
	SetDecrement(value PatternFlowTcpDstPortCounter) PatternFlowTcpDstPort
	// HasDecrement checks if Decrement has been set in PatternFlowTcpDstPort
	HasDecrement() bool
	// MetricTags returns PatternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIterIter, set in PatternFlowTcpDstPort
	MetricTags() PatternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter
	setNil()
}

type PatternFlowTcpDstPortChoiceEnum string

// Enum of Choice on PatternFlowTcpDstPort
var PatternFlowTcpDstPortChoice = struct {
	VALUE     PatternFlowTcpDstPortChoiceEnum
	VALUES    PatternFlowTcpDstPortChoiceEnum
	INCREMENT PatternFlowTcpDstPortChoiceEnum
	DECREMENT PatternFlowTcpDstPortChoiceEnum
}{
	VALUE:     PatternFlowTcpDstPortChoiceEnum("value"),
	VALUES:    PatternFlowTcpDstPortChoiceEnum("values"),
	INCREMENT: PatternFlowTcpDstPortChoiceEnum("increment"),
	DECREMENT: PatternFlowTcpDstPortChoiceEnum("decrement"),
}

func (obj *patternFlowTcpDstPort) Choice() PatternFlowTcpDstPortChoiceEnum {
	return PatternFlowTcpDstPortChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowTcpDstPort) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowTcpDstPort) setChoice(value PatternFlowTcpDstPortChoiceEnum) PatternFlowTcpDstPort {
	intValue, ok := otg.PatternFlowTcpDstPort_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowTcpDstPortChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowTcpDstPort_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowTcpDstPortChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowTcpDstPortChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowTcpDstPortChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowTcpDstPortCounter().msg()
	}

	if value == PatternFlowTcpDstPortChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowTcpDstPortCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpDstPort) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowTcpDstPortChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpDstPort) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowTcpDstPort object
func (obj *patternFlowTcpDstPort) SetValue(value uint32) PatternFlowTcpDstPort {
	obj.setChoice(PatternFlowTcpDstPortChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowTcpDstPort) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowTcpDstPort object
func (obj *patternFlowTcpDstPort) SetValues(value []uint32) PatternFlowTcpDstPort {
	obj.setChoice(PatternFlowTcpDstPortChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowTcpDstPortCounter
func (obj *patternFlowTcpDstPort) Increment() PatternFlowTcpDstPortCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowTcpDstPortChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowTcpDstPortCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowTcpDstPortCounter
func (obj *patternFlowTcpDstPort) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowTcpDstPortCounter value in the PatternFlowTcpDstPort object
func (obj *patternFlowTcpDstPort) SetIncrement(value PatternFlowTcpDstPortCounter) PatternFlowTcpDstPort {
	obj.setChoice(PatternFlowTcpDstPortChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowTcpDstPortCounter
func (obj *patternFlowTcpDstPort) Decrement() PatternFlowTcpDstPortCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowTcpDstPortChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowTcpDstPortCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowTcpDstPortCounter
func (obj *patternFlowTcpDstPort) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowTcpDstPortCounter value in the PatternFlowTcpDstPort object
func (obj *patternFlowTcpDstPort) SetDecrement(value PatternFlowTcpDstPortCounter) PatternFlowTcpDstPort {
	obj.setChoice(PatternFlowTcpDstPortChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowTcpDstPortMetricTag
func (obj *patternFlowTcpDstPort) MetricTags() PatternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowTcpDstPortMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter struct {
	obj                                 *patternFlowTcpDstPort
	patternFlowTcpDstPortMetricTagSlice []PatternFlowTcpDstPortMetricTag
	fieldPtr                            *[]*otg.PatternFlowTcpDstPortMetricTag
}

func newPatternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter(ptr *[]*otg.PatternFlowTcpDstPortMetricTag) PatternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter {
	return &patternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter{fieldPtr: ptr}
}

type PatternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter interface {
	setMsg(*patternFlowTcpDstPort) PatternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter
	Items() []PatternFlowTcpDstPortMetricTag
	Add() PatternFlowTcpDstPortMetricTag
	Append(items ...PatternFlowTcpDstPortMetricTag) PatternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter
	Set(index int, newObj PatternFlowTcpDstPortMetricTag) PatternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter
	Clear() PatternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter
	clearHolderSlice() PatternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter
	appendHolderSlice(item PatternFlowTcpDstPortMetricTag) PatternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter
}

func (obj *patternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter) setMsg(msg *patternFlowTcpDstPort) PatternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowTcpDstPortMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter) Items() []PatternFlowTcpDstPortMetricTag {
	return obj.patternFlowTcpDstPortMetricTagSlice
}

func (obj *patternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter) Add() PatternFlowTcpDstPortMetricTag {
	newObj := &otg.PatternFlowTcpDstPortMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowTcpDstPortMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowTcpDstPortMetricTagSlice = append(obj.patternFlowTcpDstPortMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter) Append(items ...PatternFlowTcpDstPortMetricTag) PatternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowTcpDstPortMetricTagSlice = append(obj.patternFlowTcpDstPortMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter) Set(index int, newObj PatternFlowTcpDstPortMetricTag) PatternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowTcpDstPortMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter) Clear() PatternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowTcpDstPortMetricTag{}
		obj.patternFlowTcpDstPortMetricTagSlice = []PatternFlowTcpDstPortMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter) clearHolderSlice() PatternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter {
	if len(obj.patternFlowTcpDstPortMetricTagSlice) > 0 {
		obj.patternFlowTcpDstPortMetricTagSlice = []PatternFlowTcpDstPortMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter) appendHolderSlice(item PatternFlowTcpDstPortMetricTag) PatternFlowTcpDstPortPatternFlowTcpDstPortMetricTagIter {
	obj.patternFlowTcpDstPortMetricTagSlice = append(obj.patternFlowTcpDstPortMetricTagSlice, item)
	return obj
}

func (obj *patternFlowTcpDstPort) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpDstPort.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowTcpDstPort.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowTcpDstPortMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowTcpDstPort) setDefault() {
	var choices_set int = 0
	var choice PatternFlowTcpDstPortChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowTcpDstPortChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowTcpDstPortChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowTcpDstPortChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowTcpDstPortChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowTcpDstPortChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowTcpDstPort")
			}
		} else {
			intVal := otg.PatternFlowTcpDstPort_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowTcpDstPort_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
