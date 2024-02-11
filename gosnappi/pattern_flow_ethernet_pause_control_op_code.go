package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowEthernetPauseControlOpCode *****
type patternFlowEthernetPauseControlOpCode struct {
	validation
	obj              *otg.PatternFlowEthernetPauseControlOpCode
	marshaller       marshalPatternFlowEthernetPauseControlOpCode
	unMarshaller     unMarshalPatternFlowEthernetPauseControlOpCode
	incrementHolder  PatternFlowEthernetPauseControlOpCodeCounter
	decrementHolder  PatternFlowEthernetPauseControlOpCodeCounter
	metricTagsHolder PatternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter
}

func NewPatternFlowEthernetPauseControlOpCode() PatternFlowEthernetPauseControlOpCode {
	obj := patternFlowEthernetPauseControlOpCode{obj: &otg.PatternFlowEthernetPauseControlOpCode{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowEthernetPauseControlOpCode) msg() *otg.PatternFlowEthernetPauseControlOpCode {
	return obj.obj
}

func (obj *patternFlowEthernetPauseControlOpCode) setMsg(msg *otg.PatternFlowEthernetPauseControlOpCode) PatternFlowEthernetPauseControlOpCode {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowEthernetPauseControlOpCode struct {
	obj *patternFlowEthernetPauseControlOpCode
}

type marshalPatternFlowEthernetPauseControlOpCode interface {
	// ToProto marshals PatternFlowEthernetPauseControlOpCode to protobuf object *otg.PatternFlowEthernetPauseControlOpCode
	ToProto() (*otg.PatternFlowEthernetPauseControlOpCode, error)
	// ToPbText marshals PatternFlowEthernetPauseControlOpCode to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowEthernetPauseControlOpCode to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowEthernetPauseControlOpCode to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowEthernetPauseControlOpCode struct {
	obj *patternFlowEthernetPauseControlOpCode
}

type unMarshalPatternFlowEthernetPauseControlOpCode interface {
	// FromProto unmarshals PatternFlowEthernetPauseControlOpCode from protobuf object *otg.PatternFlowEthernetPauseControlOpCode
	FromProto(msg *otg.PatternFlowEthernetPauseControlOpCode) (PatternFlowEthernetPauseControlOpCode, error)
	// FromPbText unmarshals PatternFlowEthernetPauseControlOpCode from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowEthernetPauseControlOpCode from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowEthernetPauseControlOpCode from JSON text
	FromJson(value string) error
}

func (obj *patternFlowEthernetPauseControlOpCode) Marshal() marshalPatternFlowEthernetPauseControlOpCode {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowEthernetPauseControlOpCode{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowEthernetPauseControlOpCode) Unmarshal() unMarshalPatternFlowEthernetPauseControlOpCode {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowEthernetPauseControlOpCode{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowEthernetPauseControlOpCode) ToProto() (*otg.PatternFlowEthernetPauseControlOpCode, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowEthernetPauseControlOpCode) FromProto(msg *otg.PatternFlowEthernetPauseControlOpCode) (PatternFlowEthernetPauseControlOpCode, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowEthernetPauseControlOpCode) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseControlOpCode) FromPbText(value string) error {
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

func (m *marshalpatternFlowEthernetPauseControlOpCode) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseControlOpCode) FromYaml(value string) error {
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

func (m *marshalpatternFlowEthernetPauseControlOpCode) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseControlOpCode) FromJson(value string) error {
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

func (obj *patternFlowEthernetPauseControlOpCode) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseControlOpCode) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseControlOpCode) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowEthernetPauseControlOpCode) Clone() (PatternFlowEthernetPauseControlOpCode, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowEthernetPauseControlOpCode()
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

func (obj *patternFlowEthernetPauseControlOpCode) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowEthernetPauseControlOpCode is control operation code
type PatternFlowEthernetPauseControlOpCode interface {
	Validation
	// msg marshals PatternFlowEthernetPauseControlOpCode to protobuf object *otg.PatternFlowEthernetPauseControlOpCode
	// and doesn't set defaults
	msg() *otg.PatternFlowEthernetPauseControlOpCode
	// setMsg unmarshals PatternFlowEthernetPauseControlOpCode from protobuf object *otg.PatternFlowEthernetPauseControlOpCode
	// and doesn't set defaults
	setMsg(*otg.PatternFlowEthernetPauseControlOpCode) PatternFlowEthernetPauseControlOpCode
	// provides marshal interface
	Marshal() marshalPatternFlowEthernetPauseControlOpCode
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowEthernetPauseControlOpCode
	// validate validates PatternFlowEthernetPauseControlOpCode
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowEthernetPauseControlOpCode, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowEthernetPauseControlOpCodeChoiceEnum, set in PatternFlowEthernetPauseControlOpCode
	Choice() PatternFlowEthernetPauseControlOpCodeChoiceEnum
	// setChoice assigns PatternFlowEthernetPauseControlOpCodeChoiceEnum provided by user to PatternFlowEthernetPauseControlOpCode
	setChoice(value PatternFlowEthernetPauseControlOpCodeChoiceEnum) PatternFlowEthernetPauseControlOpCode
	// HasChoice checks if Choice has been set in PatternFlowEthernetPauseControlOpCode
	HasChoice() bool
	// Value returns uint32, set in PatternFlowEthernetPauseControlOpCode.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowEthernetPauseControlOpCode
	SetValue(value uint32) PatternFlowEthernetPauseControlOpCode
	// HasValue checks if Value has been set in PatternFlowEthernetPauseControlOpCode
	HasValue() bool
	// Values returns []uint32, set in PatternFlowEthernetPauseControlOpCode.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowEthernetPauseControlOpCode
	SetValues(value []uint32) PatternFlowEthernetPauseControlOpCode
	// Increment returns PatternFlowEthernetPauseControlOpCodeCounter, set in PatternFlowEthernetPauseControlOpCode.
	// PatternFlowEthernetPauseControlOpCodeCounter is integer counter pattern
	Increment() PatternFlowEthernetPauseControlOpCodeCounter
	// SetIncrement assigns PatternFlowEthernetPauseControlOpCodeCounter provided by user to PatternFlowEthernetPauseControlOpCode.
	// PatternFlowEthernetPauseControlOpCodeCounter is integer counter pattern
	SetIncrement(value PatternFlowEthernetPauseControlOpCodeCounter) PatternFlowEthernetPauseControlOpCode
	// HasIncrement checks if Increment has been set in PatternFlowEthernetPauseControlOpCode
	HasIncrement() bool
	// Decrement returns PatternFlowEthernetPauseControlOpCodeCounter, set in PatternFlowEthernetPauseControlOpCode.
	// PatternFlowEthernetPauseControlOpCodeCounter is integer counter pattern
	Decrement() PatternFlowEthernetPauseControlOpCodeCounter
	// SetDecrement assigns PatternFlowEthernetPauseControlOpCodeCounter provided by user to PatternFlowEthernetPauseControlOpCode.
	// PatternFlowEthernetPauseControlOpCodeCounter is integer counter pattern
	SetDecrement(value PatternFlowEthernetPauseControlOpCodeCounter) PatternFlowEthernetPauseControlOpCode
	// HasDecrement checks if Decrement has been set in PatternFlowEthernetPauseControlOpCode
	HasDecrement() bool
	// MetricTags returns PatternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIterIter, set in PatternFlowEthernetPauseControlOpCode
	MetricTags() PatternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter
	setNil()
}

type PatternFlowEthernetPauseControlOpCodeChoiceEnum string

// Enum of Choice on PatternFlowEthernetPauseControlOpCode
var PatternFlowEthernetPauseControlOpCodeChoice = struct {
	VALUE     PatternFlowEthernetPauseControlOpCodeChoiceEnum
	VALUES    PatternFlowEthernetPauseControlOpCodeChoiceEnum
	INCREMENT PatternFlowEthernetPauseControlOpCodeChoiceEnum
	DECREMENT PatternFlowEthernetPauseControlOpCodeChoiceEnum
}{
	VALUE:     PatternFlowEthernetPauseControlOpCodeChoiceEnum("value"),
	VALUES:    PatternFlowEthernetPauseControlOpCodeChoiceEnum("values"),
	INCREMENT: PatternFlowEthernetPauseControlOpCodeChoiceEnum("increment"),
	DECREMENT: PatternFlowEthernetPauseControlOpCodeChoiceEnum("decrement"),
}

func (obj *patternFlowEthernetPauseControlOpCode) Choice() PatternFlowEthernetPauseControlOpCodeChoiceEnum {
	return PatternFlowEthernetPauseControlOpCodeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowEthernetPauseControlOpCode) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowEthernetPauseControlOpCode) setChoice(value PatternFlowEthernetPauseControlOpCodeChoiceEnum) PatternFlowEthernetPauseControlOpCode {
	intValue, ok := otg.PatternFlowEthernetPauseControlOpCode_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowEthernetPauseControlOpCodeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowEthernetPauseControlOpCode_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowEthernetPauseControlOpCodeChoice.VALUE {
		defaultValue := uint32(1)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowEthernetPauseControlOpCodeChoice.VALUES {
		defaultValue := []uint32{1}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowEthernetPauseControlOpCodeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowEthernetPauseControlOpCodeCounter().msg()
	}

	if value == PatternFlowEthernetPauseControlOpCodeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowEthernetPauseControlOpCodeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowEthernetPauseControlOpCode) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowEthernetPauseControlOpCodeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowEthernetPauseControlOpCode) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowEthernetPauseControlOpCode object
func (obj *patternFlowEthernetPauseControlOpCode) SetValue(value uint32) PatternFlowEthernetPauseControlOpCode {
	obj.setChoice(PatternFlowEthernetPauseControlOpCodeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowEthernetPauseControlOpCode) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{1})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowEthernetPauseControlOpCode object
func (obj *patternFlowEthernetPauseControlOpCode) SetValues(value []uint32) PatternFlowEthernetPauseControlOpCode {
	obj.setChoice(PatternFlowEthernetPauseControlOpCodeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowEthernetPauseControlOpCodeCounter
func (obj *patternFlowEthernetPauseControlOpCode) Increment() PatternFlowEthernetPauseControlOpCodeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowEthernetPauseControlOpCodeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowEthernetPauseControlOpCodeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowEthernetPauseControlOpCodeCounter
func (obj *patternFlowEthernetPauseControlOpCode) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowEthernetPauseControlOpCodeCounter value in the PatternFlowEthernetPauseControlOpCode object
func (obj *patternFlowEthernetPauseControlOpCode) SetIncrement(value PatternFlowEthernetPauseControlOpCodeCounter) PatternFlowEthernetPauseControlOpCode {
	obj.setChoice(PatternFlowEthernetPauseControlOpCodeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowEthernetPauseControlOpCodeCounter
func (obj *patternFlowEthernetPauseControlOpCode) Decrement() PatternFlowEthernetPauseControlOpCodeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowEthernetPauseControlOpCodeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowEthernetPauseControlOpCodeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowEthernetPauseControlOpCodeCounter
func (obj *patternFlowEthernetPauseControlOpCode) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowEthernetPauseControlOpCodeCounter value in the PatternFlowEthernetPauseControlOpCode object
func (obj *patternFlowEthernetPauseControlOpCode) SetDecrement(value PatternFlowEthernetPauseControlOpCodeCounter) PatternFlowEthernetPauseControlOpCode {
	obj.setChoice(PatternFlowEthernetPauseControlOpCodeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowEthernetPauseControlOpCodeMetricTag
func (obj *patternFlowEthernetPauseControlOpCode) MetricTags() PatternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowEthernetPauseControlOpCodeMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter struct {
	obj                                                 *patternFlowEthernetPauseControlOpCode
	patternFlowEthernetPauseControlOpCodeMetricTagSlice []PatternFlowEthernetPauseControlOpCodeMetricTag
	fieldPtr                                            *[]*otg.PatternFlowEthernetPauseControlOpCodeMetricTag
}

func newPatternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter(ptr *[]*otg.PatternFlowEthernetPauseControlOpCodeMetricTag) PatternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter {
	return &patternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter{fieldPtr: ptr}
}

type PatternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter interface {
	setMsg(*patternFlowEthernetPauseControlOpCode) PatternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter
	Items() []PatternFlowEthernetPauseControlOpCodeMetricTag
	Add() PatternFlowEthernetPauseControlOpCodeMetricTag
	Append(items ...PatternFlowEthernetPauseControlOpCodeMetricTag) PatternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter
	Set(index int, newObj PatternFlowEthernetPauseControlOpCodeMetricTag) PatternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter
	Clear() PatternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter
	clearHolderSlice() PatternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter
	appendHolderSlice(item PatternFlowEthernetPauseControlOpCodeMetricTag) PatternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter
}

func (obj *patternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter) setMsg(msg *patternFlowEthernetPauseControlOpCode) PatternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowEthernetPauseControlOpCodeMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter) Items() []PatternFlowEthernetPauseControlOpCodeMetricTag {
	return obj.patternFlowEthernetPauseControlOpCodeMetricTagSlice
}

func (obj *patternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter) Add() PatternFlowEthernetPauseControlOpCodeMetricTag {
	newObj := &otg.PatternFlowEthernetPauseControlOpCodeMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowEthernetPauseControlOpCodeMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowEthernetPauseControlOpCodeMetricTagSlice = append(obj.patternFlowEthernetPauseControlOpCodeMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter) Append(items ...PatternFlowEthernetPauseControlOpCodeMetricTag) PatternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowEthernetPauseControlOpCodeMetricTagSlice = append(obj.patternFlowEthernetPauseControlOpCodeMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter) Set(index int, newObj PatternFlowEthernetPauseControlOpCodeMetricTag) PatternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowEthernetPauseControlOpCodeMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter) Clear() PatternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowEthernetPauseControlOpCodeMetricTag{}
		obj.patternFlowEthernetPauseControlOpCodeMetricTagSlice = []PatternFlowEthernetPauseControlOpCodeMetricTag{}
	}
	return obj
}
func (obj *patternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter) clearHolderSlice() PatternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter {
	if len(obj.patternFlowEthernetPauseControlOpCodeMetricTagSlice) > 0 {
		obj.patternFlowEthernetPauseControlOpCodeMetricTagSlice = []PatternFlowEthernetPauseControlOpCodeMetricTag{}
	}
	return obj
}
func (obj *patternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter) appendHolderSlice(item PatternFlowEthernetPauseControlOpCodeMetricTag) PatternFlowEthernetPauseControlOpCodePatternFlowEthernetPauseControlOpCodeMetricTagIter {
	obj.patternFlowEthernetPauseControlOpCodeMetricTagSlice = append(obj.patternFlowEthernetPauseControlOpCodeMetricTagSlice, item)
	return obj
}

func (obj *patternFlowEthernetPauseControlOpCode) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetPauseControlOpCode.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowEthernetPauseControlOpCode.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowEthernetPauseControlOpCodeMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowEthernetPauseControlOpCode) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowEthernetPauseControlOpCodeChoice.VALUE)

	}

}
