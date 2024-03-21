package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPausePauseClass5 *****
type patternFlowPfcPausePauseClass5 struct {
	validation
	obj              *otg.PatternFlowPfcPausePauseClass5
	marshaller       marshalPatternFlowPfcPausePauseClass5
	unMarshaller     unMarshalPatternFlowPfcPausePauseClass5
	incrementHolder  PatternFlowPfcPausePauseClass5Counter
	decrementHolder  PatternFlowPfcPausePauseClass5Counter
	metricTagsHolder PatternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter
}

func NewPatternFlowPfcPausePauseClass5() PatternFlowPfcPausePauseClass5 {
	obj := patternFlowPfcPausePauseClass5{obj: &otg.PatternFlowPfcPausePauseClass5{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPausePauseClass5) msg() *otg.PatternFlowPfcPausePauseClass5 {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass5) setMsg(msg *otg.PatternFlowPfcPausePauseClass5) PatternFlowPfcPausePauseClass5 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPausePauseClass5 struct {
	obj *patternFlowPfcPausePauseClass5
}

type marshalPatternFlowPfcPausePauseClass5 interface {
	// ToProto marshals PatternFlowPfcPausePauseClass5 to protobuf object *otg.PatternFlowPfcPausePauseClass5
	ToProto() (*otg.PatternFlowPfcPausePauseClass5, error)
	// ToPbText marshals PatternFlowPfcPausePauseClass5 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPausePauseClass5 to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPausePauseClass5 to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPfcPausePauseClass5 struct {
	obj *patternFlowPfcPausePauseClass5
}

type unMarshalPatternFlowPfcPausePauseClass5 interface {
	// FromProto unmarshals PatternFlowPfcPausePauseClass5 from protobuf object *otg.PatternFlowPfcPausePauseClass5
	FromProto(msg *otg.PatternFlowPfcPausePauseClass5) (PatternFlowPfcPausePauseClass5, error)
	// FromPbText unmarshals PatternFlowPfcPausePauseClass5 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPausePauseClass5 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPausePauseClass5 from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPausePauseClass5) Marshal() marshalPatternFlowPfcPausePauseClass5 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPausePauseClass5{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPausePauseClass5) Unmarshal() unMarshalPatternFlowPfcPausePauseClass5 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPausePauseClass5{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPausePauseClass5) ToProto() (*otg.PatternFlowPfcPausePauseClass5, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPausePauseClass5) FromProto(msg *otg.PatternFlowPfcPausePauseClass5) (PatternFlowPfcPausePauseClass5, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPausePauseClass5) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass5) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass5) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass5) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass5) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass5) FromJson(value string) error {
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

func (obj *patternFlowPfcPausePauseClass5) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass5) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass5) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPausePauseClass5) Clone() (PatternFlowPfcPausePauseClass5, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPausePauseClass5()
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

func (obj *patternFlowPfcPausePauseClass5) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowPfcPausePauseClass5 is pause class 5
type PatternFlowPfcPausePauseClass5 interface {
	Validation
	// msg marshals PatternFlowPfcPausePauseClass5 to protobuf object *otg.PatternFlowPfcPausePauseClass5
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPausePauseClass5
	// setMsg unmarshals PatternFlowPfcPausePauseClass5 from protobuf object *otg.PatternFlowPfcPausePauseClass5
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPausePauseClass5) PatternFlowPfcPausePauseClass5
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPausePauseClass5
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPausePauseClass5
	// validate validates PatternFlowPfcPausePauseClass5
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPausePauseClass5, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowPfcPausePauseClass5ChoiceEnum, set in PatternFlowPfcPausePauseClass5
	Choice() PatternFlowPfcPausePauseClass5ChoiceEnum
	// setChoice assigns PatternFlowPfcPausePauseClass5ChoiceEnum provided by user to PatternFlowPfcPausePauseClass5
	setChoice(value PatternFlowPfcPausePauseClass5ChoiceEnum) PatternFlowPfcPausePauseClass5
	// HasChoice checks if Choice has been set in PatternFlowPfcPausePauseClass5
	HasChoice() bool
	// Value returns uint32, set in PatternFlowPfcPausePauseClass5.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowPfcPausePauseClass5
	SetValue(value uint32) PatternFlowPfcPausePauseClass5
	// HasValue checks if Value has been set in PatternFlowPfcPausePauseClass5
	HasValue() bool
	// Values returns []uint32, set in PatternFlowPfcPausePauseClass5.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowPfcPausePauseClass5
	SetValues(value []uint32) PatternFlowPfcPausePauseClass5
	// Increment returns PatternFlowPfcPausePauseClass5Counter, set in PatternFlowPfcPausePauseClass5.
	// PatternFlowPfcPausePauseClass5Counter is integer counter pattern
	Increment() PatternFlowPfcPausePauseClass5Counter
	// SetIncrement assigns PatternFlowPfcPausePauseClass5Counter provided by user to PatternFlowPfcPausePauseClass5.
	// PatternFlowPfcPausePauseClass5Counter is integer counter pattern
	SetIncrement(value PatternFlowPfcPausePauseClass5Counter) PatternFlowPfcPausePauseClass5
	// HasIncrement checks if Increment has been set in PatternFlowPfcPausePauseClass5
	HasIncrement() bool
	// Decrement returns PatternFlowPfcPausePauseClass5Counter, set in PatternFlowPfcPausePauseClass5.
	// PatternFlowPfcPausePauseClass5Counter is integer counter pattern
	Decrement() PatternFlowPfcPausePauseClass5Counter
	// SetDecrement assigns PatternFlowPfcPausePauseClass5Counter provided by user to PatternFlowPfcPausePauseClass5.
	// PatternFlowPfcPausePauseClass5Counter is integer counter pattern
	SetDecrement(value PatternFlowPfcPausePauseClass5Counter) PatternFlowPfcPausePauseClass5
	// HasDecrement checks if Decrement has been set in PatternFlowPfcPausePauseClass5
	HasDecrement() bool
	// MetricTags returns PatternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIterIter, set in PatternFlowPfcPausePauseClass5
	MetricTags() PatternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter
	setNil()
}

type PatternFlowPfcPausePauseClass5ChoiceEnum string

// Enum of Choice on PatternFlowPfcPausePauseClass5
var PatternFlowPfcPausePauseClass5Choice = struct {
	VALUE     PatternFlowPfcPausePauseClass5ChoiceEnum
	VALUES    PatternFlowPfcPausePauseClass5ChoiceEnum
	INCREMENT PatternFlowPfcPausePauseClass5ChoiceEnum
	DECREMENT PatternFlowPfcPausePauseClass5ChoiceEnum
}{
	VALUE:     PatternFlowPfcPausePauseClass5ChoiceEnum("value"),
	VALUES:    PatternFlowPfcPausePauseClass5ChoiceEnum("values"),
	INCREMENT: PatternFlowPfcPausePauseClass5ChoiceEnum("increment"),
	DECREMENT: PatternFlowPfcPausePauseClass5ChoiceEnum("decrement"),
}

func (obj *patternFlowPfcPausePauseClass5) Choice() PatternFlowPfcPausePauseClass5ChoiceEnum {
	return PatternFlowPfcPausePauseClass5ChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowPfcPausePauseClass5) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowPfcPausePauseClass5) setChoice(value PatternFlowPfcPausePauseClass5ChoiceEnum) PatternFlowPfcPausePauseClass5 {
	intValue, ok := otg.PatternFlowPfcPausePauseClass5_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowPfcPausePauseClass5ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowPfcPausePauseClass5_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowPfcPausePauseClass5Choice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowPfcPausePauseClass5Choice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowPfcPausePauseClass5Choice.INCREMENT {
		obj.obj.Increment = NewPatternFlowPfcPausePauseClass5Counter().msg()
	}

	if value == PatternFlowPfcPausePauseClass5Choice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowPfcPausePauseClass5Counter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowPfcPausePauseClass5) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowPfcPausePauseClass5Choice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowPfcPausePauseClass5) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowPfcPausePauseClass5 object
func (obj *patternFlowPfcPausePauseClass5) SetValue(value uint32) PatternFlowPfcPausePauseClass5 {
	obj.setChoice(PatternFlowPfcPausePauseClass5Choice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowPfcPausePauseClass5) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowPfcPausePauseClass5 object
func (obj *patternFlowPfcPausePauseClass5) SetValues(value []uint32) PatternFlowPfcPausePauseClass5 {
	obj.setChoice(PatternFlowPfcPausePauseClass5Choice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowPfcPausePauseClass5Counter
func (obj *patternFlowPfcPausePauseClass5) Increment() PatternFlowPfcPausePauseClass5Counter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowPfcPausePauseClass5Choice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowPfcPausePauseClass5Counter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowPfcPausePauseClass5Counter
func (obj *patternFlowPfcPausePauseClass5) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowPfcPausePauseClass5Counter value in the PatternFlowPfcPausePauseClass5 object
func (obj *patternFlowPfcPausePauseClass5) SetIncrement(value PatternFlowPfcPausePauseClass5Counter) PatternFlowPfcPausePauseClass5 {
	obj.setChoice(PatternFlowPfcPausePauseClass5Choice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowPfcPausePauseClass5Counter
func (obj *patternFlowPfcPausePauseClass5) Decrement() PatternFlowPfcPausePauseClass5Counter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowPfcPausePauseClass5Choice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowPfcPausePauseClass5Counter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowPfcPausePauseClass5Counter
func (obj *patternFlowPfcPausePauseClass5) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowPfcPausePauseClass5Counter value in the PatternFlowPfcPausePauseClass5 object
func (obj *patternFlowPfcPausePauseClass5) SetDecrement(value PatternFlowPfcPausePauseClass5Counter) PatternFlowPfcPausePauseClass5 {
	obj.setChoice(PatternFlowPfcPausePauseClass5Choice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowPfcPausePauseClass5MetricTag
func (obj *patternFlowPfcPausePauseClass5) MetricTags() PatternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowPfcPausePauseClass5MetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter struct {
	obj                                          *patternFlowPfcPausePauseClass5
	patternFlowPfcPausePauseClass5MetricTagSlice []PatternFlowPfcPausePauseClass5MetricTag
	fieldPtr                                     *[]*otg.PatternFlowPfcPausePauseClass5MetricTag
}

func newPatternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter(ptr *[]*otg.PatternFlowPfcPausePauseClass5MetricTag) PatternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter {
	return &patternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter{fieldPtr: ptr}
}

type PatternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter interface {
	setMsg(*patternFlowPfcPausePauseClass5) PatternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter
	Items() []PatternFlowPfcPausePauseClass5MetricTag
	Add() PatternFlowPfcPausePauseClass5MetricTag
	Append(items ...PatternFlowPfcPausePauseClass5MetricTag) PatternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter
	Set(index int, newObj PatternFlowPfcPausePauseClass5MetricTag) PatternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter
	Clear() PatternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter
	clearHolderSlice() PatternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter
	appendHolderSlice(item PatternFlowPfcPausePauseClass5MetricTag) PatternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter
}

func (obj *patternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter) setMsg(msg *patternFlowPfcPausePauseClass5) PatternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowPfcPausePauseClass5MetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter) Items() []PatternFlowPfcPausePauseClass5MetricTag {
	return obj.patternFlowPfcPausePauseClass5MetricTagSlice
}

func (obj *patternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter) Add() PatternFlowPfcPausePauseClass5MetricTag {
	newObj := &otg.PatternFlowPfcPausePauseClass5MetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowPfcPausePauseClass5MetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowPfcPausePauseClass5MetricTagSlice = append(obj.patternFlowPfcPausePauseClass5MetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter) Append(items ...PatternFlowPfcPausePauseClass5MetricTag) PatternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowPfcPausePauseClass5MetricTagSlice = append(obj.patternFlowPfcPausePauseClass5MetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter) Set(index int, newObj PatternFlowPfcPausePauseClass5MetricTag) PatternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowPfcPausePauseClass5MetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter) Clear() PatternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowPfcPausePauseClass5MetricTag{}
		obj.patternFlowPfcPausePauseClass5MetricTagSlice = []PatternFlowPfcPausePauseClass5MetricTag{}
	}
	return obj
}
func (obj *patternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter) clearHolderSlice() PatternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter {
	if len(obj.patternFlowPfcPausePauseClass5MetricTagSlice) > 0 {
		obj.patternFlowPfcPausePauseClass5MetricTagSlice = []PatternFlowPfcPausePauseClass5MetricTag{}
	}
	return obj
}
func (obj *patternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter) appendHolderSlice(item PatternFlowPfcPausePauseClass5MetricTag) PatternFlowPfcPausePauseClass5PatternFlowPfcPausePauseClass5MetricTagIter {
	obj.patternFlowPfcPausePauseClass5MetricTagSlice = append(obj.patternFlowPfcPausePauseClass5MetricTagSlice, item)
	return obj
}

func (obj *patternFlowPfcPausePauseClass5) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass5.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowPfcPausePauseClass5.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowPfcPausePauseClass5MetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowPfcPausePauseClass5) setDefault() {
	var choices_set int = 0
	var choice PatternFlowPfcPausePauseClass5ChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowPfcPausePauseClass5Choice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowPfcPausePauseClass5Choice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowPfcPausePauseClass5Choice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowPfcPausePauseClass5Choice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowPfcPausePauseClass5Choice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowPfcPausePauseClass5")
			}
		} else {
			intVal := otg.PatternFlowPfcPausePauseClass5_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowPfcPausePauseClass5_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
