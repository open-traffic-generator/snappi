package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGreProtocol *****
type patternFlowGreProtocol struct {
	validation
	obj              *otg.PatternFlowGreProtocol
	marshaller       marshalPatternFlowGreProtocol
	unMarshaller     unMarshalPatternFlowGreProtocol
	incrementHolder  PatternFlowGreProtocolCounter
	decrementHolder  PatternFlowGreProtocolCounter
	metricTagsHolder PatternFlowGreProtocolPatternFlowGreProtocolMetricTagIter
}

func NewPatternFlowGreProtocol() PatternFlowGreProtocol {
	obj := patternFlowGreProtocol{obj: &otg.PatternFlowGreProtocol{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGreProtocol) msg() *otg.PatternFlowGreProtocol {
	return obj.obj
}

func (obj *patternFlowGreProtocol) setMsg(msg *otg.PatternFlowGreProtocol) PatternFlowGreProtocol {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGreProtocol struct {
	obj *patternFlowGreProtocol
}

type marshalPatternFlowGreProtocol interface {
	// ToProto marshals PatternFlowGreProtocol to protobuf object *otg.PatternFlowGreProtocol
	ToProto() (*otg.PatternFlowGreProtocol, error)
	// ToPbText marshals PatternFlowGreProtocol to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGreProtocol to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGreProtocol to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGreProtocol struct {
	obj *patternFlowGreProtocol
}

type unMarshalPatternFlowGreProtocol interface {
	// FromProto unmarshals PatternFlowGreProtocol from protobuf object *otg.PatternFlowGreProtocol
	FromProto(msg *otg.PatternFlowGreProtocol) (PatternFlowGreProtocol, error)
	// FromPbText unmarshals PatternFlowGreProtocol from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGreProtocol from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGreProtocol from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGreProtocol) Marshal() marshalPatternFlowGreProtocol {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGreProtocol{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGreProtocol) Unmarshal() unMarshalPatternFlowGreProtocol {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGreProtocol{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGreProtocol) ToProto() (*otg.PatternFlowGreProtocol, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGreProtocol) FromProto(msg *otg.PatternFlowGreProtocol) (PatternFlowGreProtocol, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGreProtocol) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGreProtocol) FromPbText(value string) error {
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

func (m *marshalpatternFlowGreProtocol) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGreProtocol) FromYaml(value string) error {
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

func (m *marshalpatternFlowGreProtocol) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGreProtocol) FromJson(value string) error {
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

func (obj *patternFlowGreProtocol) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGreProtocol) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGreProtocol) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGreProtocol) Clone() (PatternFlowGreProtocol, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGreProtocol()
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

func (obj *patternFlowGreProtocol) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGreProtocol is protocol type of encapsulated payload
type PatternFlowGreProtocol interface {
	Validation
	// msg marshals PatternFlowGreProtocol to protobuf object *otg.PatternFlowGreProtocol
	// and doesn't set defaults
	msg() *otg.PatternFlowGreProtocol
	// setMsg unmarshals PatternFlowGreProtocol from protobuf object *otg.PatternFlowGreProtocol
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGreProtocol) PatternFlowGreProtocol
	// provides marshal interface
	Marshal() marshalPatternFlowGreProtocol
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGreProtocol
	// validate validates PatternFlowGreProtocol
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGreProtocol, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGreProtocolChoiceEnum, set in PatternFlowGreProtocol
	Choice() PatternFlowGreProtocolChoiceEnum
	// setChoice assigns PatternFlowGreProtocolChoiceEnum provided by user to PatternFlowGreProtocol
	setChoice(value PatternFlowGreProtocolChoiceEnum) PatternFlowGreProtocol
	// HasChoice checks if Choice has been set in PatternFlowGreProtocol
	HasChoice() bool
	// Value returns uint32, set in PatternFlowGreProtocol.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowGreProtocol
	SetValue(value uint32) PatternFlowGreProtocol
	// HasValue checks if Value has been set in PatternFlowGreProtocol
	HasValue() bool
	// Values returns []uint32, set in PatternFlowGreProtocol.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowGreProtocol
	SetValues(value []uint32) PatternFlowGreProtocol
	// Increment returns PatternFlowGreProtocolCounter, set in PatternFlowGreProtocol.
	// PatternFlowGreProtocolCounter is integer counter pattern
	Increment() PatternFlowGreProtocolCounter
	// SetIncrement assigns PatternFlowGreProtocolCounter provided by user to PatternFlowGreProtocol.
	// PatternFlowGreProtocolCounter is integer counter pattern
	SetIncrement(value PatternFlowGreProtocolCounter) PatternFlowGreProtocol
	// HasIncrement checks if Increment has been set in PatternFlowGreProtocol
	HasIncrement() bool
	// Decrement returns PatternFlowGreProtocolCounter, set in PatternFlowGreProtocol.
	// PatternFlowGreProtocolCounter is integer counter pattern
	Decrement() PatternFlowGreProtocolCounter
	// SetDecrement assigns PatternFlowGreProtocolCounter provided by user to PatternFlowGreProtocol.
	// PatternFlowGreProtocolCounter is integer counter pattern
	SetDecrement(value PatternFlowGreProtocolCounter) PatternFlowGreProtocol
	// HasDecrement checks if Decrement has been set in PatternFlowGreProtocol
	HasDecrement() bool
	// MetricTags returns PatternFlowGreProtocolPatternFlowGreProtocolMetricTagIterIter, set in PatternFlowGreProtocol
	MetricTags() PatternFlowGreProtocolPatternFlowGreProtocolMetricTagIter
	setNil()
}

type PatternFlowGreProtocolChoiceEnum string

// Enum of Choice on PatternFlowGreProtocol
var PatternFlowGreProtocolChoice = struct {
	VALUE     PatternFlowGreProtocolChoiceEnum
	VALUES    PatternFlowGreProtocolChoiceEnum
	INCREMENT PatternFlowGreProtocolChoiceEnum
	DECREMENT PatternFlowGreProtocolChoiceEnum
}{
	VALUE:     PatternFlowGreProtocolChoiceEnum("value"),
	VALUES:    PatternFlowGreProtocolChoiceEnum("values"),
	INCREMENT: PatternFlowGreProtocolChoiceEnum("increment"),
	DECREMENT: PatternFlowGreProtocolChoiceEnum("decrement"),
}

func (obj *patternFlowGreProtocol) Choice() PatternFlowGreProtocolChoiceEnum {
	return PatternFlowGreProtocolChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGreProtocol) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGreProtocol) setChoice(value PatternFlowGreProtocolChoiceEnum) PatternFlowGreProtocol {
	intValue, ok := otg.PatternFlowGreProtocol_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGreProtocolChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGreProtocol_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGreProtocolChoice.VALUE {
		defaultValue := uint32(2048)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGreProtocolChoice.VALUES {
		defaultValue := []uint32{2048}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGreProtocolChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGreProtocolCounter().msg()
	}

	if value == PatternFlowGreProtocolChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGreProtocolCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGreProtocol) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGreProtocolChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGreProtocol) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowGreProtocol object
func (obj *patternFlowGreProtocol) SetValue(value uint32) PatternFlowGreProtocol {
	obj.setChoice(PatternFlowGreProtocolChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowGreProtocol) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{2048})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowGreProtocol object
func (obj *patternFlowGreProtocol) SetValues(value []uint32) PatternFlowGreProtocol {
	obj.setChoice(PatternFlowGreProtocolChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGreProtocolCounter
func (obj *patternFlowGreProtocol) Increment() PatternFlowGreProtocolCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGreProtocolChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGreProtocolCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGreProtocolCounter
func (obj *patternFlowGreProtocol) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGreProtocolCounter value in the PatternFlowGreProtocol object
func (obj *patternFlowGreProtocol) SetIncrement(value PatternFlowGreProtocolCounter) PatternFlowGreProtocol {
	obj.setChoice(PatternFlowGreProtocolChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGreProtocolCounter
func (obj *patternFlowGreProtocol) Decrement() PatternFlowGreProtocolCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGreProtocolChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGreProtocolCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGreProtocolCounter
func (obj *patternFlowGreProtocol) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGreProtocolCounter value in the PatternFlowGreProtocol object
func (obj *patternFlowGreProtocol) SetDecrement(value PatternFlowGreProtocolCounter) PatternFlowGreProtocol {
	obj.setChoice(PatternFlowGreProtocolChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGreProtocolMetricTag
func (obj *patternFlowGreProtocol) MetricTags() PatternFlowGreProtocolPatternFlowGreProtocolMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGreProtocolMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGreProtocolPatternFlowGreProtocolMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGreProtocolPatternFlowGreProtocolMetricTagIter struct {
	obj                                  *patternFlowGreProtocol
	patternFlowGreProtocolMetricTagSlice []PatternFlowGreProtocolMetricTag
	fieldPtr                             *[]*otg.PatternFlowGreProtocolMetricTag
}

func newPatternFlowGreProtocolPatternFlowGreProtocolMetricTagIter(ptr *[]*otg.PatternFlowGreProtocolMetricTag) PatternFlowGreProtocolPatternFlowGreProtocolMetricTagIter {
	return &patternFlowGreProtocolPatternFlowGreProtocolMetricTagIter{fieldPtr: ptr}
}

type PatternFlowGreProtocolPatternFlowGreProtocolMetricTagIter interface {
	setMsg(*patternFlowGreProtocol) PatternFlowGreProtocolPatternFlowGreProtocolMetricTagIter
	Items() []PatternFlowGreProtocolMetricTag
	Add() PatternFlowGreProtocolMetricTag
	Append(items ...PatternFlowGreProtocolMetricTag) PatternFlowGreProtocolPatternFlowGreProtocolMetricTagIter
	Set(index int, newObj PatternFlowGreProtocolMetricTag) PatternFlowGreProtocolPatternFlowGreProtocolMetricTagIter
	Clear() PatternFlowGreProtocolPatternFlowGreProtocolMetricTagIter
	clearHolderSlice() PatternFlowGreProtocolPatternFlowGreProtocolMetricTagIter
	appendHolderSlice(item PatternFlowGreProtocolMetricTag) PatternFlowGreProtocolPatternFlowGreProtocolMetricTagIter
}

func (obj *patternFlowGreProtocolPatternFlowGreProtocolMetricTagIter) setMsg(msg *patternFlowGreProtocol) PatternFlowGreProtocolPatternFlowGreProtocolMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGreProtocolMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGreProtocolPatternFlowGreProtocolMetricTagIter) Items() []PatternFlowGreProtocolMetricTag {
	return obj.patternFlowGreProtocolMetricTagSlice
}

func (obj *patternFlowGreProtocolPatternFlowGreProtocolMetricTagIter) Add() PatternFlowGreProtocolMetricTag {
	newObj := &otg.PatternFlowGreProtocolMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGreProtocolMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGreProtocolMetricTagSlice = append(obj.patternFlowGreProtocolMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGreProtocolPatternFlowGreProtocolMetricTagIter) Append(items ...PatternFlowGreProtocolMetricTag) PatternFlowGreProtocolPatternFlowGreProtocolMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGreProtocolMetricTagSlice = append(obj.patternFlowGreProtocolMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGreProtocolPatternFlowGreProtocolMetricTagIter) Set(index int, newObj PatternFlowGreProtocolMetricTag) PatternFlowGreProtocolPatternFlowGreProtocolMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGreProtocolMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGreProtocolPatternFlowGreProtocolMetricTagIter) Clear() PatternFlowGreProtocolPatternFlowGreProtocolMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGreProtocolMetricTag{}
		obj.patternFlowGreProtocolMetricTagSlice = []PatternFlowGreProtocolMetricTag{}
	}
	return obj
}
func (obj *patternFlowGreProtocolPatternFlowGreProtocolMetricTagIter) clearHolderSlice() PatternFlowGreProtocolPatternFlowGreProtocolMetricTagIter {
	if len(obj.patternFlowGreProtocolMetricTagSlice) > 0 {
		obj.patternFlowGreProtocolMetricTagSlice = []PatternFlowGreProtocolMetricTag{}
	}
	return obj
}
func (obj *patternFlowGreProtocolPatternFlowGreProtocolMetricTagIter) appendHolderSlice(item PatternFlowGreProtocolMetricTag) PatternFlowGreProtocolPatternFlowGreProtocolMetricTagIter {
	obj.patternFlowGreProtocolMetricTagSlice = append(obj.patternFlowGreProtocolMetricTagSlice, item)
	return obj
}

func (obj *patternFlowGreProtocol) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGreProtocol.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowGreProtocol.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowGreProtocolMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGreProtocol) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowGreProtocolChoice.VALUE)

	}

}
