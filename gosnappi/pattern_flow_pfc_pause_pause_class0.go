package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPausePauseClass0 *****
type patternFlowPfcPausePauseClass0 struct {
	validation
	obj              *otg.PatternFlowPfcPausePauseClass0
	marshaller       marshalPatternFlowPfcPausePauseClass0
	unMarshaller     unMarshalPatternFlowPfcPausePauseClass0
	incrementHolder  PatternFlowPfcPausePauseClass0Counter
	decrementHolder  PatternFlowPfcPausePauseClass0Counter
	metricTagsHolder PatternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter
}

func NewPatternFlowPfcPausePauseClass0() PatternFlowPfcPausePauseClass0 {
	obj := patternFlowPfcPausePauseClass0{obj: &otg.PatternFlowPfcPausePauseClass0{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPausePauseClass0) msg() *otg.PatternFlowPfcPausePauseClass0 {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass0) setMsg(msg *otg.PatternFlowPfcPausePauseClass0) PatternFlowPfcPausePauseClass0 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPausePauseClass0 struct {
	obj *patternFlowPfcPausePauseClass0
}

type marshalPatternFlowPfcPausePauseClass0 interface {
	// ToProto marshals PatternFlowPfcPausePauseClass0 to protobuf object *otg.PatternFlowPfcPausePauseClass0
	ToProto() (*otg.PatternFlowPfcPausePauseClass0, error)
	// ToPbText marshals PatternFlowPfcPausePauseClass0 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPausePauseClass0 to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPausePauseClass0 to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowPfcPausePauseClass0 to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowPfcPausePauseClass0 struct {
	obj *patternFlowPfcPausePauseClass0
}

type unMarshalPatternFlowPfcPausePauseClass0 interface {
	// FromProto unmarshals PatternFlowPfcPausePauseClass0 from protobuf object *otg.PatternFlowPfcPausePauseClass0
	FromProto(msg *otg.PatternFlowPfcPausePauseClass0) (PatternFlowPfcPausePauseClass0, error)
	// FromPbText unmarshals PatternFlowPfcPausePauseClass0 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPausePauseClass0 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPausePauseClass0 from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPausePauseClass0) Marshal() marshalPatternFlowPfcPausePauseClass0 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPausePauseClass0{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPausePauseClass0) Unmarshal() unMarshalPatternFlowPfcPausePauseClass0 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPausePauseClass0{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPausePauseClass0) ToProto() (*otg.PatternFlowPfcPausePauseClass0, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPausePauseClass0) FromProto(msg *otg.PatternFlowPfcPausePauseClass0) (PatternFlowPfcPausePauseClass0, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPausePauseClass0) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass0) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass0) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass0) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass0) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowPfcPausePauseClass0) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass0) FromJson(value string) error {
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

func (obj *patternFlowPfcPausePauseClass0) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass0) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass0) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPausePauseClass0) Clone() (PatternFlowPfcPausePauseClass0, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPausePauseClass0()
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

func (obj *patternFlowPfcPausePauseClass0) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowPfcPausePauseClass0 is pause class 0
type PatternFlowPfcPausePauseClass0 interface {
	Validation
	// msg marshals PatternFlowPfcPausePauseClass0 to protobuf object *otg.PatternFlowPfcPausePauseClass0
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPausePauseClass0
	// setMsg unmarshals PatternFlowPfcPausePauseClass0 from protobuf object *otg.PatternFlowPfcPausePauseClass0
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPausePauseClass0) PatternFlowPfcPausePauseClass0
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPausePauseClass0
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPausePauseClass0
	// validate validates PatternFlowPfcPausePauseClass0
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPausePauseClass0, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowPfcPausePauseClass0ChoiceEnum, set in PatternFlowPfcPausePauseClass0
	Choice() PatternFlowPfcPausePauseClass0ChoiceEnum
	// setChoice assigns PatternFlowPfcPausePauseClass0ChoiceEnum provided by user to PatternFlowPfcPausePauseClass0
	setChoice(value PatternFlowPfcPausePauseClass0ChoiceEnum) PatternFlowPfcPausePauseClass0
	// HasChoice checks if Choice has been set in PatternFlowPfcPausePauseClass0
	HasChoice() bool
	// Value returns uint32, set in PatternFlowPfcPausePauseClass0.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowPfcPausePauseClass0
	SetValue(value uint32) PatternFlowPfcPausePauseClass0
	// HasValue checks if Value has been set in PatternFlowPfcPausePauseClass0
	HasValue() bool
	// Values returns []uint32, set in PatternFlowPfcPausePauseClass0.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowPfcPausePauseClass0
	SetValues(value []uint32) PatternFlowPfcPausePauseClass0
	// Increment returns PatternFlowPfcPausePauseClass0Counter, set in PatternFlowPfcPausePauseClass0.
	// PatternFlowPfcPausePauseClass0Counter is integer counter pattern
	Increment() PatternFlowPfcPausePauseClass0Counter
	// SetIncrement assigns PatternFlowPfcPausePauseClass0Counter provided by user to PatternFlowPfcPausePauseClass0.
	// PatternFlowPfcPausePauseClass0Counter is integer counter pattern
	SetIncrement(value PatternFlowPfcPausePauseClass0Counter) PatternFlowPfcPausePauseClass0
	// HasIncrement checks if Increment has been set in PatternFlowPfcPausePauseClass0
	HasIncrement() bool
	// Decrement returns PatternFlowPfcPausePauseClass0Counter, set in PatternFlowPfcPausePauseClass0.
	// PatternFlowPfcPausePauseClass0Counter is integer counter pattern
	Decrement() PatternFlowPfcPausePauseClass0Counter
	// SetDecrement assigns PatternFlowPfcPausePauseClass0Counter provided by user to PatternFlowPfcPausePauseClass0.
	// PatternFlowPfcPausePauseClass0Counter is integer counter pattern
	SetDecrement(value PatternFlowPfcPausePauseClass0Counter) PatternFlowPfcPausePauseClass0
	// HasDecrement checks if Decrement has been set in PatternFlowPfcPausePauseClass0
	HasDecrement() bool
	// MetricTags returns PatternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIterIter, set in PatternFlowPfcPausePauseClass0
	MetricTags() PatternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter
	setNil()
}

type PatternFlowPfcPausePauseClass0ChoiceEnum string

// Enum of Choice on PatternFlowPfcPausePauseClass0
var PatternFlowPfcPausePauseClass0Choice = struct {
	VALUE     PatternFlowPfcPausePauseClass0ChoiceEnum
	VALUES    PatternFlowPfcPausePauseClass0ChoiceEnum
	INCREMENT PatternFlowPfcPausePauseClass0ChoiceEnum
	DECREMENT PatternFlowPfcPausePauseClass0ChoiceEnum
}{
	VALUE:     PatternFlowPfcPausePauseClass0ChoiceEnum("value"),
	VALUES:    PatternFlowPfcPausePauseClass0ChoiceEnum("values"),
	INCREMENT: PatternFlowPfcPausePauseClass0ChoiceEnum("increment"),
	DECREMENT: PatternFlowPfcPausePauseClass0ChoiceEnum("decrement"),
}

func (obj *patternFlowPfcPausePauseClass0) Choice() PatternFlowPfcPausePauseClass0ChoiceEnum {
	return PatternFlowPfcPausePauseClass0ChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowPfcPausePauseClass0) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowPfcPausePauseClass0) setChoice(value PatternFlowPfcPausePauseClass0ChoiceEnum) PatternFlowPfcPausePauseClass0 {
	intValue, ok := otg.PatternFlowPfcPausePauseClass0_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowPfcPausePauseClass0ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowPfcPausePauseClass0_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowPfcPausePauseClass0Choice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowPfcPausePauseClass0Choice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowPfcPausePauseClass0Choice.INCREMENT {
		obj.obj.Increment = NewPatternFlowPfcPausePauseClass0Counter().msg()
	}

	if value == PatternFlowPfcPausePauseClass0Choice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowPfcPausePauseClass0Counter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowPfcPausePauseClass0) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowPfcPausePauseClass0Choice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowPfcPausePauseClass0) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowPfcPausePauseClass0 object
func (obj *patternFlowPfcPausePauseClass0) SetValue(value uint32) PatternFlowPfcPausePauseClass0 {
	obj.setChoice(PatternFlowPfcPausePauseClass0Choice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowPfcPausePauseClass0) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowPfcPausePauseClass0 object
func (obj *patternFlowPfcPausePauseClass0) SetValues(value []uint32) PatternFlowPfcPausePauseClass0 {
	obj.setChoice(PatternFlowPfcPausePauseClass0Choice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowPfcPausePauseClass0Counter
func (obj *patternFlowPfcPausePauseClass0) Increment() PatternFlowPfcPausePauseClass0Counter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowPfcPausePauseClass0Choice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowPfcPausePauseClass0Counter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowPfcPausePauseClass0Counter
func (obj *patternFlowPfcPausePauseClass0) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowPfcPausePauseClass0Counter value in the PatternFlowPfcPausePauseClass0 object
func (obj *patternFlowPfcPausePauseClass0) SetIncrement(value PatternFlowPfcPausePauseClass0Counter) PatternFlowPfcPausePauseClass0 {
	obj.setChoice(PatternFlowPfcPausePauseClass0Choice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowPfcPausePauseClass0Counter
func (obj *patternFlowPfcPausePauseClass0) Decrement() PatternFlowPfcPausePauseClass0Counter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowPfcPausePauseClass0Choice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowPfcPausePauseClass0Counter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowPfcPausePauseClass0Counter
func (obj *patternFlowPfcPausePauseClass0) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowPfcPausePauseClass0Counter value in the PatternFlowPfcPausePauseClass0 object
func (obj *patternFlowPfcPausePauseClass0) SetDecrement(value PatternFlowPfcPausePauseClass0Counter) PatternFlowPfcPausePauseClass0 {
	obj.setChoice(PatternFlowPfcPausePauseClass0Choice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowPfcPausePauseClass0MetricTag
func (obj *patternFlowPfcPausePauseClass0) MetricTags() PatternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowPfcPausePauseClass0MetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter struct {
	obj                                          *patternFlowPfcPausePauseClass0
	patternFlowPfcPausePauseClass0MetricTagSlice []PatternFlowPfcPausePauseClass0MetricTag
	fieldPtr                                     *[]*otg.PatternFlowPfcPausePauseClass0MetricTag
}

func newPatternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter(ptr *[]*otg.PatternFlowPfcPausePauseClass0MetricTag) PatternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter {
	return &patternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter{fieldPtr: ptr}
}

type PatternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter interface {
	setMsg(*patternFlowPfcPausePauseClass0) PatternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter
	Items() []PatternFlowPfcPausePauseClass0MetricTag
	Add() PatternFlowPfcPausePauseClass0MetricTag
	Append(items ...PatternFlowPfcPausePauseClass0MetricTag) PatternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter
	Set(index int, newObj PatternFlowPfcPausePauseClass0MetricTag) PatternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter
	Clear() PatternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter
	clearHolderSlice() PatternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter
	appendHolderSlice(item PatternFlowPfcPausePauseClass0MetricTag) PatternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter
}

func (obj *patternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter) setMsg(msg *patternFlowPfcPausePauseClass0) PatternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowPfcPausePauseClass0MetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter) Items() []PatternFlowPfcPausePauseClass0MetricTag {
	return obj.patternFlowPfcPausePauseClass0MetricTagSlice
}

func (obj *patternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter) Add() PatternFlowPfcPausePauseClass0MetricTag {
	newObj := &otg.PatternFlowPfcPausePauseClass0MetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowPfcPausePauseClass0MetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowPfcPausePauseClass0MetricTagSlice = append(obj.patternFlowPfcPausePauseClass0MetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter) Append(items ...PatternFlowPfcPausePauseClass0MetricTag) PatternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowPfcPausePauseClass0MetricTagSlice = append(obj.patternFlowPfcPausePauseClass0MetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter) Set(index int, newObj PatternFlowPfcPausePauseClass0MetricTag) PatternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowPfcPausePauseClass0MetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter) Clear() PatternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowPfcPausePauseClass0MetricTag{}
		obj.patternFlowPfcPausePauseClass0MetricTagSlice = []PatternFlowPfcPausePauseClass0MetricTag{}
	}
	return obj
}
func (obj *patternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter) clearHolderSlice() PatternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter {
	if len(obj.patternFlowPfcPausePauseClass0MetricTagSlice) > 0 {
		obj.patternFlowPfcPausePauseClass0MetricTagSlice = []PatternFlowPfcPausePauseClass0MetricTag{}
	}
	return obj
}
func (obj *patternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter) appendHolderSlice(item PatternFlowPfcPausePauseClass0MetricTag) PatternFlowPfcPausePauseClass0PatternFlowPfcPausePauseClass0MetricTagIter {
	obj.patternFlowPfcPausePauseClass0MetricTagSlice = append(obj.patternFlowPfcPausePauseClass0MetricTagSlice, item)
	return obj
}

func (obj *patternFlowPfcPausePauseClass0) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass0.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowPfcPausePauseClass0.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowPfcPausePauseClass0MetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowPfcPausePauseClass0) setDefault() {
	var choices_set int = 0
	var choice PatternFlowPfcPausePauseClass0ChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowPfcPausePauseClass0Choice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowPfcPausePauseClass0Choice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowPfcPausePauseClass0Choice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowPfcPausePauseClass0Choice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowPfcPausePauseClass0Choice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowPfcPausePauseClass0")
			}
		} else {
			intVal := otg.PatternFlowPfcPausePauseClass0_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowPfcPausePauseClass0_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
