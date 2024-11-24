package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPauseControlOpCode *****
type patternFlowPfcPauseControlOpCode struct {
	validation
	obj              *otg.PatternFlowPfcPauseControlOpCode
	marshaller       marshalPatternFlowPfcPauseControlOpCode
	unMarshaller     unMarshalPatternFlowPfcPauseControlOpCode
	incrementHolder  PatternFlowPfcPauseControlOpCodeCounter
	decrementHolder  PatternFlowPfcPauseControlOpCodeCounter
	metricTagsHolder PatternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter
}

func NewPatternFlowPfcPauseControlOpCode() PatternFlowPfcPauseControlOpCode {
	obj := patternFlowPfcPauseControlOpCode{obj: &otg.PatternFlowPfcPauseControlOpCode{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPauseControlOpCode) msg() *otg.PatternFlowPfcPauseControlOpCode {
	return obj.obj
}

func (obj *patternFlowPfcPauseControlOpCode) setMsg(msg *otg.PatternFlowPfcPauseControlOpCode) PatternFlowPfcPauseControlOpCode {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPauseControlOpCode struct {
	obj *patternFlowPfcPauseControlOpCode
}

type marshalPatternFlowPfcPauseControlOpCode interface {
	// ToProto marshals PatternFlowPfcPauseControlOpCode to protobuf object *otg.PatternFlowPfcPauseControlOpCode
	ToProto() (*otg.PatternFlowPfcPauseControlOpCode, error)
	// ToPbText marshals PatternFlowPfcPauseControlOpCode to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPauseControlOpCode to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPauseControlOpCode to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowPfcPauseControlOpCode to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowPfcPauseControlOpCode struct {
	obj *patternFlowPfcPauseControlOpCode
}

type unMarshalPatternFlowPfcPauseControlOpCode interface {
	// FromProto unmarshals PatternFlowPfcPauseControlOpCode from protobuf object *otg.PatternFlowPfcPauseControlOpCode
	FromProto(msg *otg.PatternFlowPfcPauseControlOpCode) (PatternFlowPfcPauseControlOpCode, error)
	// FromPbText unmarshals PatternFlowPfcPauseControlOpCode from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPauseControlOpCode from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPauseControlOpCode from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPauseControlOpCode) Marshal() marshalPatternFlowPfcPauseControlOpCode {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPauseControlOpCode{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPauseControlOpCode) Unmarshal() unMarshalPatternFlowPfcPauseControlOpCode {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPauseControlOpCode{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPauseControlOpCode) ToProto() (*otg.PatternFlowPfcPauseControlOpCode, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPauseControlOpCode) FromProto(msg *otg.PatternFlowPfcPauseControlOpCode) (PatternFlowPfcPauseControlOpCode, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPauseControlOpCode) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseControlOpCode) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPauseControlOpCode) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseControlOpCode) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPauseControlOpCode) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowPfcPauseControlOpCode) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseControlOpCode) FromJson(value string) error {
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

func (obj *patternFlowPfcPauseControlOpCode) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseControlOpCode) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseControlOpCode) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPauseControlOpCode) Clone() (PatternFlowPfcPauseControlOpCode, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPauseControlOpCode()
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

func (obj *patternFlowPfcPauseControlOpCode) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowPfcPauseControlOpCode is control operation code
type PatternFlowPfcPauseControlOpCode interface {
	Validation
	// msg marshals PatternFlowPfcPauseControlOpCode to protobuf object *otg.PatternFlowPfcPauseControlOpCode
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPauseControlOpCode
	// setMsg unmarshals PatternFlowPfcPauseControlOpCode from protobuf object *otg.PatternFlowPfcPauseControlOpCode
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPauseControlOpCode) PatternFlowPfcPauseControlOpCode
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPauseControlOpCode
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPauseControlOpCode
	// validate validates PatternFlowPfcPauseControlOpCode
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPauseControlOpCode, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowPfcPauseControlOpCodeChoiceEnum, set in PatternFlowPfcPauseControlOpCode
	Choice() PatternFlowPfcPauseControlOpCodeChoiceEnum
	// setChoice assigns PatternFlowPfcPauseControlOpCodeChoiceEnum provided by user to PatternFlowPfcPauseControlOpCode
	setChoice(value PatternFlowPfcPauseControlOpCodeChoiceEnum) PatternFlowPfcPauseControlOpCode
	// HasChoice checks if Choice has been set in PatternFlowPfcPauseControlOpCode
	HasChoice() bool
	// Value returns uint32, set in PatternFlowPfcPauseControlOpCode.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowPfcPauseControlOpCode
	SetValue(value uint32) PatternFlowPfcPauseControlOpCode
	// HasValue checks if Value has been set in PatternFlowPfcPauseControlOpCode
	HasValue() bool
	// Values returns []uint32, set in PatternFlowPfcPauseControlOpCode.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowPfcPauseControlOpCode
	SetValues(value []uint32) PatternFlowPfcPauseControlOpCode
	// Increment returns PatternFlowPfcPauseControlOpCodeCounter, set in PatternFlowPfcPauseControlOpCode.
	// PatternFlowPfcPauseControlOpCodeCounter is integer counter pattern
	Increment() PatternFlowPfcPauseControlOpCodeCounter
	// SetIncrement assigns PatternFlowPfcPauseControlOpCodeCounter provided by user to PatternFlowPfcPauseControlOpCode.
	// PatternFlowPfcPauseControlOpCodeCounter is integer counter pattern
	SetIncrement(value PatternFlowPfcPauseControlOpCodeCounter) PatternFlowPfcPauseControlOpCode
	// HasIncrement checks if Increment has been set in PatternFlowPfcPauseControlOpCode
	HasIncrement() bool
	// Decrement returns PatternFlowPfcPauseControlOpCodeCounter, set in PatternFlowPfcPauseControlOpCode.
	// PatternFlowPfcPauseControlOpCodeCounter is integer counter pattern
	Decrement() PatternFlowPfcPauseControlOpCodeCounter
	// SetDecrement assigns PatternFlowPfcPauseControlOpCodeCounter provided by user to PatternFlowPfcPauseControlOpCode.
	// PatternFlowPfcPauseControlOpCodeCounter is integer counter pattern
	SetDecrement(value PatternFlowPfcPauseControlOpCodeCounter) PatternFlowPfcPauseControlOpCode
	// HasDecrement checks if Decrement has been set in PatternFlowPfcPauseControlOpCode
	HasDecrement() bool
	// MetricTags returns PatternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIterIter, set in PatternFlowPfcPauseControlOpCode
	MetricTags() PatternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter
	setNil()
}

type PatternFlowPfcPauseControlOpCodeChoiceEnum string

// Enum of Choice on PatternFlowPfcPauseControlOpCode
var PatternFlowPfcPauseControlOpCodeChoice = struct {
	VALUE     PatternFlowPfcPauseControlOpCodeChoiceEnum
	VALUES    PatternFlowPfcPauseControlOpCodeChoiceEnum
	INCREMENT PatternFlowPfcPauseControlOpCodeChoiceEnum
	DECREMENT PatternFlowPfcPauseControlOpCodeChoiceEnum
}{
	VALUE:     PatternFlowPfcPauseControlOpCodeChoiceEnum("value"),
	VALUES:    PatternFlowPfcPauseControlOpCodeChoiceEnum("values"),
	INCREMENT: PatternFlowPfcPauseControlOpCodeChoiceEnum("increment"),
	DECREMENT: PatternFlowPfcPauseControlOpCodeChoiceEnum("decrement"),
}

func (obj *patternFlowPfcPauseControlOpCode) Choice() PatternFlowPfcPauseControlOpCodeChoiceEnum {
	return PatternFlowPfcPauseControlOpCodeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowPfcPauseControlOpCode) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowPfcPauseControlOpCode) setChoice(value PatternFlowPfcPauseControlOpCodeChoiceEnum) PatternFlowPfcPauseControlOpCode {
	intValue, ok := otg.PatternFlowPfcPauseControlOpCode_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowPfcPauseControlOpCodeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowPfcPauseControlOpCode_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowPfcPauseControlOpCodeChoice.VALUE {
		defaultValue := uint32(257)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowPfcPauseControlOpCodeChoice.VALUES {
		defaultValue := []uint32{257}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowPfcPauseControlOpCodeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowPfcPauseControlOpCodeCounter().msg()
	}

	if value == PatternFlowPfcPauseControlOpCodeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowPfcPauseControlOpCodeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowPfcPauseControlOpCode) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowPfcPauseControlOpCodeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowPfcPauseControlOpCode) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowPfcPauseControlOpCode object
func (obj *patternFlowPfcPauseControlOpCode) SetValue(value uint32) PatternFlowPfcPauseControlOpCode {
	obj.setChoice(PatternFlowPfcPauseControlOpCodeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowPfcPauseControlOpCode) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{257})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowPfcPauseControlOpCode object
func (obj *patternFlowPfcPauseControlOpCode) SetValues(value []uint32) PatternFlowPfcPauseControlOpCode {
	obj.setChoice(PatternFlowPfcPauseControlOpCodeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowPfcPauseControlOpCodeCounter
func (obj *patternFlowPfcPauseControlOpCode) Increment() PatternFlowPfcPauseControlOpCodeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowPfcPauseControlOpCodeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowPfcPauseControlOpCodeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowPfcPauseControlOpCodeCounter
func (obj *patternFlowPfcPauseControlOpCode) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowPfcPauseControlOpCodeCounter value in the PatternFlowPfcPauseControlOpCode object
func (obj *patternFlowPfcPauseControlOpCode) SetIncrement(value PatternFlowPfcPauseControlOpCodeCounter) PatternFlowPfcPauseControlOpCode {
	obj.setChoice(PatternFlowPfcPauseControlOpCodeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowPfcPauseControlOpCodeCounter
func (obj *patternFlowPfcPauseControlOpCode) Decrement() PatternFlowPfcPauseControlOpCodeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowPfcPauseControlOpCodeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowPfcPauseControlOpCodeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowPfcPauseControlOpCodeCounter
func (obj *patternFlowPfcPauseControlOpCode) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowPfcPauseControlOpCodeCounter value in the PatternFlowPfcPauseControlOpCode object
func (obj *patternFlowPfcPauseControlOpCode) SetDecrement(value PatternFlowPfcPauseControlOpCodeCounter) PatternFlowPfcPauseControlOpCode {
	obj.setChoice(PatternFlowPfcPauseControlOpCodeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowPfcPauseControlOpCodeMetricTag
func (obj *patternFlowPfcPauseControlOpCode) MetricTags() PatternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowPfcPauseControlOpCodeMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter struct {
	obj                                            *patternFlowPfcPauseControlOpCode
	patternFlowPfcPauseControlOpCodeMetricTagSlice []PatternFlowPfcPauseControlOpCodeMetricTag
	fieldPtr                                       *[]*otg.PatternFlowPfcPauseControlOpCodeMetricTag
}

func newPatternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter(ptr *[]*otg.PatternFlowPfcPauseControlOpCodeMetricTag) PatternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter {
	return &patternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter{fieldPtr: ptr}
}

type PatternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter interface {
	setMsg(*patternFlowPfcPauseControlOpCode) PatternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter
	Items() []PatternFlowPfcPauseControlOpCodeMetricTag
	Add() PatternFlowPfcPauseControlOpCodeMetricTag
	Append(items ...PatternFlowPfcPauseControlOpCodeMetricTag) PatternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter
	Set(index int, newObj PatternFlowPfcPauseControlOpCodeMetricTag) PatternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter
	Clear() PatternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter
	clearHolderSlice() PatternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter
	appendHolderSlice(item PatternFlowPfcPauseControlOpCodeMetricTag) PatternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter
}

func (obj *patternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter) setMsg(msg *patternFlowPfcPauseControlOpCode) PatternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowPfcPauseControlOpCodeMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter) Items() []PatternFlowPfcPauseControlOpCodeMetricTag {
	return obj.patternFlowPfcPauseControlOpCodeMetricTagSlice
}

func (obj *patternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter) Add() PatternFlowPfcPauseControlOpCodeMetricTag {
	newObj := &otg.PatternFlowPfcPauseControlOpCodeMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowPfcPauseControlOpCodeMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowPfcPauseControlOpCodeMetricTagSlice = append(obj.patternFlowPfcPauseControlOpCodeMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter) Append(items ...PatternFlowPfcPauseControlOpCodeMetricTag) PatternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowPfcPauseControlOpCodeMetricTagSlice = append(obj.patternFlowPfcPauseControlOpCodeMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter) Set(index int, newObj PatternFlowPfcPauseControlOpCodeMetricTag) PatternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowPfcPauseControlOpCodeMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter) Clear() PatternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowPfcPauseControlOpCodeMetricTag{}
		obj.patternFlowPfcPauseControlOpCodeMetricTagSlice = []PatternFlowPfcPauseControlOpCodeMetricTag{}
	}
	return obj
}
func (obj *patternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter) clearHolderSlice() PatternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter {
	if len(obj.patternFlowPfcPauseControlOpCodeMetricTagSlice) > 0 {
		obj.patternFlowPfcPauseControlOpCodeMetricTagSlice = []PatternFlowPfcPauseControlOpCodeMetricTag{}
	}
	return obj
}
func (obj *patternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter) appendHolderSlice(item PatternFlowPfcPauseControlOpCodeMetricTag) PatternFlowPfcPauseControlOpCodePatternFlowPfcPauseControlOpCodeMetricTagIter {
	obj.patternFlowPfcPauseControlOpCodeMetricTagSlice = append(obj.patternFlowPfcPauseControlOpCodeMetricTagSlice, item)
	return obj
}

func (obj *patternFlowPfcPauseControlOpCode) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPauseControlOpCode.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowPfcPauseControlOpCode.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowPfcPauseControlOpCodeMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowPfcPauseControlOpCode) setDefault() {
	var choices_set int = 0
	var choice PatternFlowPfcPauseControlOpCodeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowPfcPauseControlOpCodeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowPfcPauseControlOpCodeChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowPfcPauseControlOpCodeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowPfcPauseControlOpCodeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowPfcPauseControlOpCodeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowPfcPauseControlOpCode")
			}
		} else {
			intVal := otg.PatternFlowPfcPauseControlOpCode_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowPfcPauseControlOpCode_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
