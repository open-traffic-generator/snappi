package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPausePauseClass1 *****
type patternFlowPfcPausePauseClass1 struct {
	validation
	obj              *otg.PatternFlowPfcPausePauseClass1
	marshaller       marshalPatternFlowPfcPausePauseClass1
	unMarshaller     unMarshalPatternFlowPfcPausePauseClass1
	incrementHolder  PatternFlowPfcPausePauseClass1Counter
	decrementHolder  PatternFlowPfcPausePauseClass1Counter
	metricTagsHolder PatternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter
}

func NewPatternFlowPfcPausePauseClass1() PatternFlowPfcPausePauseClass1 {
	obj := patternFlowPfcPausePauseClass1{obj: &otg.PatternFlowPfcPausePauseClass1{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPausePauseClass1) msg() *otg.PatternFlowPfcPausePauseClass1 {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass1) setMsg(msg *otg.PatternFlowPfcPausePauseClass1) PatternFlowPfcPausePauseClass1 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPausePauseClass1 struct {
	obj *patternFlowPfcPausePauseClass1
}

type marshalPatternFlowPfcPausePauseClass1 interface {
	// ToProto marshals PatternFlowPfcPausePauseClass1 to protobuf object *otg.PatternFlowPfcPausePauseClass1
	ToProto() (*otg.PatternFlowPfcPausePauseClass1, error)
	// ToPbText marshals PatternFlowPfcPausePauseClass1 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPausePauseClass1 to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPausePauseClass1 to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowPfcPausePauseClass1 to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowPfcPausePauseClass1 struct {
	obj *patternFlowPfcPausePauseClass1
}

type unMarshalPatternFlowPfcPausePauseClass1 interface {
	// FromProto unmarshals PatternFlowPfcPausePauseClass1 from protobuf object *otg.PatternFlowPfcPausePauseClass1
	FromProto(msg *otg.PatternFlowPfcPausePauseClass1) (PatternFlowPfcPausePauseClass1, error)
	// FromPbText unmarshals PatternFlowPfcPausePauseClass1 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPausePauseClass1 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPausePauseClass1 from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPausePauseClass1) Marshal() marshalPatternFlowPfcPausePauseClass1 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPausePauseClass1{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPausePauseClass1) Unmarshal() unMarshalPatternFlowPfcPausePauseClass1 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPausePauseClass1{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPausePauseClass1) ToProto() (*otg.PatternFlowPfcPausePauseClass1, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPausePauseClass1) FromProto(msg *otg.PatternFlowPfcPausePauseClass1) (PatternFlowPfcPausePauseClass1, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPausePauseClass1) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass1) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass1) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass1) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass1) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowPfcPausePauseClass1) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass1) FromJson(value string) error {
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

func (obj *patternFlowPfcPausePauseClass1) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass1) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass1) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPausePauseClass1) Clone() (PatternFlowPfcPausePauseClass1, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPausePauseClass1()
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

func (obj *patternFlowPfcPausePauseClass1) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowPfcPausePauseClass1 is pause class 1
type PatternFlowPfcPausePauseClass1 interface {
	Validation
	// msg marshals PatternFlowPfcPausePauseClass1 to protobuf object *otg.PatternFlowPfcPausePauseClass1
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPausePauseClass1
	// setMsg unmarshals PatternFlowPfcPausePauseClass1 from protobuf object *otg.PatternFlowPfcPausePauseClass1
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPausePauseClass1) PatternFlowPfcPausePauseClass1
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPausePauseClass1
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPausePauseClass1
	// validate validates PatternFlowPfcPausePauseClass1
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPausePauseClass1, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowPfcPausePauseClass1ChoiceEnum, set in PatternFlowPfcPausePauseClass1
	Choice() PatternFlowPfcPausePauseClass1ChoiceEnum
	// setChoice assigns PatternFlowPfcPausePauseClass1ChoiceEnum provided by user to PatternFlowPfcPausePauseClass1
	setChoice(value PatternFlowPfcPausePauseClass1ChoiceEnum) PatternFlowPfcPausePauseClass1
	// HasChoice checks if Choice has been set in PatternFlowPfcPausePauseClass1
	HasChoice() bool
	// Value returns uint32, set in PatternFlowPfcPausePauseClass1.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowPfcPausePauseClass1
	SetValue(value uint32) PatternFlowPfcPausePauseClass1
	// HasValue checks if Value has been set in PatternFlowPfcPausePauseClass1
	HasValue() bool
	// Values returns []uint32, set in PatternFlowPfcPausePauseClass1.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowPfcPausePauseClass1
	SetValues(value []uint32) PatternFlowPfcPausePauseClass1
	// Increment returns PatternFlowPfcPausePauseClass1Counter, set in PatternFlowPfcPausePauseClass1.
	// PatternFlowPfcPausePauseClass1Counter is integer counter pattern
	Increment() PatternFlowPfcPausePauseClass1Counter
	// SetIncrement assigns PatternFlowPfcPausePauseClass1Counter provided by user to PatternFlowPfcPausePauseClass1.
	// PatternFlowPfcPausePauseClass1Counter is integer counter pattern
	SetIncrement(value PatternFlowPfcPausePauseClass1Counter) PatternFlowPfcPausePauseClass1
	// HasIncrement checks if Increment has been set in PatternFlowPfcPausePauseClass1
	HasIncrement() bool
	// Decrement returns PatternFlowPfcPausePauseClass1Counter, set in PatternFlowPfcPausePauseClass1.
	// PatternFlowPfcPausePauseClass1Counter is integer counter pattern
	Decrement() PatternFlowPfcPausePauseClass1Counter
	// SetDecrement assigns PatternFlowPfcPausePauseClass1Counter provided by user to PatternFlowPfcPausePauseClass1.
	// PatternFlowPfcPausePauseClass1Counter is integer counter pattern
	SetDecrement(value PatternFlowPfcPausePauseClass1Counter) PatternFlowPfcPausePauseClass1
	// HasDecrement checks if Decrement has been set in PatternFlowPfcPausePauseClass1
	HasDecrement() bool
	// MetricTags returns PatternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIterIter, set in PatternFlowPfcPausePauseClass1
	MetricTags() PatternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter
	setNil()
}

type PatternFlowPfcPausePauseClass1ChoiceEnum string

// Enum of Choice on PatternFlowPfcPausePauseClass1
var PatternFlowPfcPausePauseClass1Choice = struct {
	VALUE     PatternFlowPfcPausePauseClass1ChoiceEnum
	VALUES    PatternFlowPfcPausePauseClass1ChoiceEnum
	INCREMENT PatternFlowPfcPausePauseClass1ChoiceEnum
	DECREMENT PatternFlowPfcPausePauseClass1ChoiceEnum
}{
	VALUE:     PatternFlowPfcPausePauseClass1ChoiceEnum("value"),
	VALUES:    PatternFlowPfcPausePauseClass1ChoiceEnum("values"),
	INCREMENT: PatternFlowPfcPausePauseClass1ChoiceEnum("increment"),
	DECREMENT: PatternFlowPfcPausePauseClass1ChoiceEnum("decrement"),
}

func (obj *patternFlowPfcPausePauseClass1) Choice() PatternFlowPfcPausePauseClass1ChoiceEnum {
	return PatternFlowPfcPausePauseClass1ChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowPfcPausePauseClass1) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowPfcPausePauseClass1) setChoice(value PatternFlowPfcPausePauseClass1ChoiceEnum) PatternFlowPfcPausePauseClass1 {
	intValue, ok := otg.PatternFlowPfcPausePauseClass1_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowPfcPausePauseClass1ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowPfcPausePauseClass1_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowPfcPausePauseClass1Choice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowPfcPausePauseClass1Choice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowPfcPausePauseClass1Choice.INCREMENT {
		obj.obj.Increment = NewPatternFlowPfcPausePauseClass1Counter().msg()
	}

	if value == PatternFlowPfcPausePauseClass1Choice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowPfcPausePauseClass1Counter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowPfcPausePauseClass1) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowPfcPausePauseClass1Choice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowPfcPausePauseClass1) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowPfcPausePauseClass1 object
func (obj *patternFlowPfcPausePauseClass1) SetValue(value uint32) PatternFlowPfcPausePauseClass1 {
	obj.setChoice(PatternFlowPfcPausePauseClass1Choice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowPfcPausePauseClass1) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowPfcPausePauseClass1 object
func (obj *patternFlowPfcPausePauseClass1) SetValues(value []uint32) PatternFlowPfcPausePauseClass1 {
	obj.setChoice(PatternFlowPfcPausePauseClass1Choice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowPfcPausePauseClass1Counter
func (obj *patternFlowPfcPausePauseClass1) Increment() PatternFlowPfcPausePauseClass1Counter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowPfcPausePauseClass1Choice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowPfcPausePauseClass1Counter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowPfcPausePauseClass1Counter
func (obj *patternFlowPfcPausePauseClass1) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowPfcPausePauseClass1Counter value in the PatternFlowPfcPausePauseClass1 object
func (obj *patternFlowPfcPausePauseClass1) SetIncrement(value PatternFlowPfcPausePauseClass1Counter) PatternFlowPfcPausePauseClass1 {
	obj.setChoice(PatternFlowPfcPausePauseClass1Choice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowPfcPausePauseClass1Counter
func (obj *patternFlowPfcPausePauseClass1) Decrement() PatternFlowPfcPausePauseClass1Counter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowPfcPausePauseClass1Choice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowPfcPausePauseClass1Counter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowPfcPausePauseClass1Counter
func (obj *patternFlowPfcPausePauseClass1) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowPfcPausePauseClass1Counter value in the PatternFlowPfcPausePauseClass1 object
func (obj *patternFlowPfcPausePauseClass1) SetDecrement(value PatternFlowPfcPausePauseClass1Counter) PatternFlowPfcPausePauseClass1 {
	obj.setChoice(PatternFlowPfcPausePauseClass1Choice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowPfcPausePauseClass1MetricTag
func (obj *patternFlowPfcPausePauseClass1) MetricTags() PatternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowPfcPausePauseClass1MetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter struct {
	obj                                          *patternFlowPfcPausePauseClass1
	patternFlowPfcPausePauseClass1MetricTagSlice []PatternFlowPfcPausePauseClass1MetricTag
	fieldPtr                                     *[]*otg.PatternFlowPfcPausePauseClass1MetricTag
}

func newPatternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter(ptr *[]*otg.PatternFlowPfcPausePauseClass1MetricTag) PatternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter {
	return &patternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter{fieldPtr: ptr}
}

type PatternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter interface {
	setMsg(*patternFlowPfcPausePauseClass1) PatternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter
	Items() []PatternFlowPfcPausePauseClass1MetricTag
	Add() PatternFlowPfcPausePauseClass1MetricTag
	Append(items ...PatternFlowPfcPausePauseClass1MetricTag) PatternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter
	Set(index int, newObj PatternFlowPfcPausePauseClass1MetricTag) PatternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter
	Clear() PatternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter
	clearHolderSlice() PatternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter
	appendHolderSlice(item PatternFlowPfcPausePauseClass1MetricTag) PatternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter
}

func (obj *patternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter) setMsg(msg *patternFlowPfcPausePauseClass1) PatternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowPfcPausePauseClass1MetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter) Items() []PatternFlowPfcPausePauseClass1MetricTag {
	return obj.patternFlowPfcPausePauseClass1MetricTagSlice
}

func (obj *patternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter) Add() PatternFlowPfcPausePauseClass1MetricTag {
	newObj := &otg.PatternFlowPfcPausePauseClass1MetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowPfcPausePauseClass1MetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowPfcPausePauseClass1MetricTagSlice = append(obj.patternFlowPfcPausePauseClass1MetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter) Append(items ...PatternFlowPfcPausePauseClass1MetricTag) PatternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowPfcPausePauseClass1MetricTagSlice = append(obj.patternFlowPfcPausePauseClass1MetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter) Set(index int, newObj PatternFlowPfcPausePauseClass1MetricTag) PatternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowPfcPausePauseClass1MetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter) Clear() PatternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowPfcPausePauseClass1MetricTag{}
		obj.patternFlowPfcPausePauseClass1MetricTagSlice = []PatternFlowPfcPausePauseClass1MetricTag{}
	}
	return obj
}
func (obj *patternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter) clearHolderSlice() PatternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter {
	if len(obj.patternFlowPfcPausePauseClass1MetricTagSlice) > 0 {
		obj.patternFlowPfcPausePauseClass1MetricTagSlice = []PatternFlowPfcPausePauseClass1MetricTag{}
	}
	return obj
}
func (obj *patternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter) appendHolderSlice(item PatternFlowPfcPausePauseClass1MetricTag) PatternFlowPfcPausePauseClass1PatternFlowPfcPausePauseClass1MetricTagIter {
	obj.patternFlowPfcPausePauseClass1MetricTagSlice = append(obj.patternFlowPfcPausePauseClass1MetricTagSlice, item)
	return obj
}

func (obj *patternFlowPfcPausePauseClass1) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass1.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowPfcPausePauseClass1.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowPfcPausePauseClass1MetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowPfcPausePauseClass1) setDefault() {
	var choices_set int = 0
	var choice PatternFlowPfcPausePauseClass1ChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowPfcPausePauseClass1Choice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowPfcPausePauseClass1Choice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowPfcPausePauseClass1Choice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowPfcPausePauseClass1Choice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowPfcPausePauseClass1Choice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowPfcPausePauseClass1")
			}
		} else {
			intVal := otg.PatternFlowPfcPausePauseClass1_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowPfcPausePauseClass1_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
