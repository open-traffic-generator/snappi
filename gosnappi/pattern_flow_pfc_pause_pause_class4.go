package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPausePauseClass4 *****
type patternFlowPfcPausePauseClass4 struct {
	validation
	obj              *otg.PatternFlowPfcPausePauseClass4
	marshaller       marshalPatternFlowPfcPausePauseClass4
	unMarshaller     unMarshalPatternFlowPfcPausePauseClass4
	incrementHolder  PatternFlowPfcPausePauseClass4Counter
	decrementHolder  PatternFlowPfcPausePauseClass4Counter
	metricTagsHolder PatternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter
}

func NewPatternFlowPfcPausePauseClass4() PatternFlowPfcPausePauseClass4 {
	obj := patternFlowPfcPausePauseClass4{obj: &otg.PatternFlowPfcPausePauseClass4{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPausePauseClass4) msg() *otg.PatternFlowPfcPausePauseClass4 {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass4) setMsg(msg *otg.PatternFlowPfcPausePauseClass4) PatternFlowPfcPausePauseClass4 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPausePauseClass4 struct {
	obj *patternFlowPfcPausePauseClass4
}

type marshalPatternFlowPfcPausePauseClass4 interface {
	// ToProto marshals PatternFlowPfcPausePauseClass4 to protobuf object *otg.PatternFlowPfcPausePauseClass4
	ToProto() (*otg.PatternFlowPfcPausePauseClass4, error)
	// ToPbText marshals PatternFlowPfcPausePauseClass4 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPausePauseClass4 to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPausePauseClass4 to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPfcPausePauseClass4 struct {
	obj *patternFlowPfcPausePauseClass4
}

type unMarshalPatternFlowPfcPausePauseClass4 interface {
	// FromProto unmarshals PatternFlowPfcPausePauseClass4 from protobuf object *otg.PatternFlowPfcPausePauseClass4
	FromProto(msg *otg.PatternFlowPfcPausePauseClass4) (PatternFlowPfcPausePauseClass4, error)
	// FromPbText unmarshals PatternFlowPfcPausePauseClass4 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPausePauseClass4 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPausePauseClass4 from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPausePauseClass4) Marshal() marshalPatternFlowPfcPausePauseClass4 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPausePauseClass4{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPausePauseClass4) Unmarshal() unMarshalPatternFlowPfcPausePauseClass4 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPausePauseClass4{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPausePauseClass4) ToProto() (*otg.PatternFlowPfcPausePauseClass4, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPausePauseClass4) FromProto(msg *otg.PatternFlowPfcPausePauseClass4) (PatternFlowPfcPausePauseClass4, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPausePauseClass4) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass4) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass4) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass4) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass4) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass4) FromJson(value string) error {
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

func (obj *patternFlowPfcPausePauseClass4) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass4) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass4) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPausePauseClass4) Clone() (PatternFlowPfcPausePauseClass4, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPausePauseClass4()
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

func (obj *patternFlowPfcPausePauseClass4) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowPfcPausePauseClass4 is pause class 4
type PatternFlowPfcPausePauseClass4 interface {
	Validation
	// msg marshals PatternFlowPfcPausePauseClass4 to protobuf object *otg.PatternFlowPfcPausePauseClass4
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPausePauseClass4
	// setMsg unmarshals PatternFlowPfcPausePauseClass4 from protobuf object *otg.PatternFlowPfcPausePauseClass4
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPausePauseClass4) PatternFlowPfcPausePauseClass4
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPausePauseClass4
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPausePauseClass4
	// validate validates PatternFlowPfcPausePauseClass4
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPausePauseClass4, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowPfcPausePauseClass4ChoiceEnum, set in PatternFlowPfcPausePauseClass4
	Choice() PatternFlowPfcPausePauseClass4ChoiceEnum
	// setChoice assigns PatternFlowPfcPausePauseClass4ChoiceEnum provided by user to PatternFlowPfcPausePauseClass4
	setChoice(value PatternFlowPfcPausePauseClass4ChoiceEnum) PatternFlowPfcPausePauseClass4
	// HasChoice checks if Choice has been set in PatternFlowPfcPausePauseClass4
	HasChoice() bool
	// Value returns uint32, set in PatternFlowPfcPausePauseClass4.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowPfcPausePauseClass4
	SetValue(value uint32) PatternFlowPfcPausePauseClass4
	// HasValue checks if Value has been set in PatternFlowPfcPausePauseClass4
	HasValue() bool
	// Values returns []uint32, set in PatternFlowPfcPausePauseClass4.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowPfcPausePauseClass4
	SetValues(value []uint32) PatternFlowPfcPausePauseClass4
	// Increment returns PatternFlowPfcPausePauseClass4Counter, set in PatternFlowPfcPausePauseClass4.
	// PatternFlowPfcPausePauseClass4Counter is integer counter pattern
	Increment() PatternFlowPfcPausePauseClass4Counter
	// SetIncrement assigns PatternFlowPfcPausePauseClass4Counter provided by user to PatternFlowPfcPausePauseClass4.
	// PatternFlowPfcPausePauseClass4Counter is integer counter pattern
	SetIncrement(value PatternFlowPfcPausePauseClass4Counter) PatternFlowPfcPausePauseClass4
	// HasIncrement checks if Increment has been set in PatternFlowPfcPausePauseClass4
	HasIncrement() bool
	// Decrement returns PatternFlowPfcPausePauseClass4Counter, set in PatternFlowPfcPausePauseClass4.
	// PatternFlowPfcPausePauseClass4Counter is integer counter pattern
	Decrement() PatternFlowPfcPausePauseClass4Counter
	// SetDecrement assigns PatternFlowPfcPausePauseClass4Counter provided by user to PatternFlowPfcPausePauseClass4.
	// PatternFlowPfcPausePauseClass4Counter is integer counter pattern
	SetDecrement(value PatternFlowPfcPausePauseClass4Counter) PatternFlowPfcPausePauseClass4
	// HasDecrement checks if Decrement has been set in PatternFlowPfcPausePauseClass4
	HasDecrement() bool
	// MetricTags returns PatternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIterIter, set in PatternFlowPfcPausePauseClass4
	MetricTags() PatternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter
	setNil()
}

type PatternFlowPfcPausePauseClass4ChoiceEnum string

// Enum of Choice on PatternFlowPfcPausePauseClass4
var PatternFlowPfcPausePauseClass4Choice = struct {
	VALUE     PatternFlowPfcPausePauseClass4ChoiceEnum
	VALUES    PatternFlowPfcPausePauseClass4ChoiceEnum
	INCREMENT PatternFlowPfcPausePauseClass4ChoiceEnum
	DECREMENT PatternFlowPfcPausePauseClass4ChoiceEnum
}{
	VALUE:     PatternFlowPfcPausePauseClass4ChoiceEnum("value"),
	VALUES:    PatternFlowPfcPausePauseClass4ChoiceEnum("values"),
	INCREMENT: PatternFlowPfcPausePauseClass4ChoiceEnum("increment"),
	DECREMENT: PatternFlowPfcPausePauseClass4ChoiceEnum("decrement"),
}

func (obj *patternFlowPfcPausePauseClass4) Choice() PatternFlowPfcPausePauseClass4ChoiceEnum {
	return PatternFlowPfcPausePauseClass4ChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowPfcPausePauseClass4) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowPfcPausePauseClass4) setChoice(value PatternFlowPfcPausePauseClass4ChoiceEnum) PatternFlowPfcPausePauseClass4 {
	intValue, ok := otg.PatternFlowPfcPausePauseClass4_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowPfcPausePauseClass4ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowPfcPausePauseClass4_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowPfcPausePauseClass4Choice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowPfcPausePauseClass4Choice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowPfcPausePauseClass4Choice.INCREMENT {
		obj.obj.Increment = NewPatternFlowPfcPausePauseClass4Counter().msg()
	}

	if value == PatternFlowPfcPausePauseClass4Choice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowPfcPausePauseClass4Counter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowPfcPausePauseClass4) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowPfcPausePauseClass4Choice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowPfcPausePauseClass4) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowPfcPausePauseClass4 object
func (obj *patternFlowPfcPausePauseClass4) SetValue(value uint32) PatternFlowPfcPausePauseClass4 {
	obj.setChoice(PatternFlowPfcPausePauseClass4Choice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowPfcPausePauseClass4) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowPfcPausePauseClass4 object
func (obj *patternFlowPfcPausePauseClass4) SetValues(value []uint32) PatternFlowPfcPausePauseClass4 {
	obj.setChoice(PatternFlowPfcPausePauseClass4Choice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowPfcPausePauseClass4Counter
func (obj *patternFlowPfcPausePauseClass4) Increment() PatternFlowPfcPausePauseClass4Counter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowPfcPausePauseClass4Choice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowPfcPausePauseClass4Counter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowPfcPausePauseClass4Counter
func (obj *patternFlowPfcPausePauseClass4) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowPfcPausePauseClass4Counter value in the PatternFlowPfcPausePauseClass4 object
func (obj *patternFlowPfcPausePauseClass4) SetIncrement(value PatternFlowPfcPausePauseClass4Counter) PatternFlowPfcPausePauseClass4 {
	obj.setChoice(PatternFlowPfcPausePauseClass4Choice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowPfcPausePauseClass4Counter
func (obj *patternFlowPfcPausePauseClass4) Decrement() PatternFlowPfcPausePauseClass4Counter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowPfcPausePauseClass4Choice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowPfcPausePauseClass4Counter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowPfcPausePauseClass4Counter
func (obj *patternFlowPfcPausePauseClass4) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowPfcPausePauseClass4Counter value in the PatternFlowPfcPausePauseClass4 object
func (obj *patternFlowPfcPausePauseClass4) SetDecrement(value PatternFlowPfcPausePauseClass4Counter) PatternFlowPfcPausePauseClass4 {
	obj.setChoice(PatternFlowPfcPausePauseClass4Choice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowPfcPausePauseClass4MetricTag
func (obj *patternFlowPfcPausePauseClass4) MetricTags() PatternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowPfcPausePauseClass4MetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter struct {
	obj                                          *patternFlowPfcPausePauseClass4
	patternFlowPfcPausePauseClass4MetricTagSlice []PatternFlowPfcPausePauseClass4MetricTag
	fieldPtr                                     *[]*otg.PatternFlowPfcPausePauseClass4MetricTag
}

func newPatternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter(ptr *[]*otg.PatternFlowPfcPausePauseClass4MetricTag) PatternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter {
	return &patternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter{fieldPtr: ptr}
}

type PatternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter interface {
	setMsg(*patternFlowPfcPausePauseClass4) PatternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter
	Items() []PatternFlowPfcPausePauseClass4MetricTag
	Add() PatternFlowPfcPausePauseClass4MetricTag
	Append(items ...PatternFlowPfcPausePauseClass4MetricTag) PatternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter
	Set(index int, newObj PatternFlowPfcPausePauseClass4MetricTag) PatternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter
	Clear() PatternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter
	clearHolderSlice() PatternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter
	appendHolderSlice(item PatternFlowPfcPausePauseClass4MetricTag) PatternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter
}

func (obj *patternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter) setMsg(msg *patternFlowPfcPausePauseClass4) PatternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowPfcPausePauseClass4MetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter) Items() []PatternFlowPfcPausePauseClass4MetricTag {
	return obj.patternFlowPfcPausePauseClass4MetricTagSlice
}

func (obj *patternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter) Add() PatternFlowPfcPausePauseClass4MetricTag {
	newObj := &otg.PatternFlowPfcPausePauseClass4MetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowPfcPausePauseClass4MetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowPfcPausePauseClass4MetricTagSlice = append(obj.patternFlowPfcPausePauseClass4MetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter) Append(items ...PatternFlowPfcPausePauseClass4MetricTag) PatternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowPfcPausePauseClass4MetricTagSlice = append(obj.patternFlowPfcPausePauseClass4MetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter) Set(index int, newObj PatternFlowPfcPausePauseClass4MetricTag) PatternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowPfcPausePauseClass4MetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter) Clear() PatternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowPfcPausePauseClass4MetricTag{}
		obj.patternFlowPfcPausePauseClass4MetricTagSlice = []PatternFlowPfcPausePauseClass4MetricTag{}
	}
	return obj
}
func (obj *patternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter) clearHolderSlice() PatternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter {
	if len(obj.patternFlowPfcPausePauseClass4MetricTagSlice) > 0 {
		obj.patternFlowPfcPausePauseClass4MetricTagSlice = []PatternFlowPfcPausePauseClass4MetricTag{}
	}
	return obj
}
func (obj *patternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter) appendHolderSlice(item PatternFlowPfcPausePauseClass4MetricTag) PatternFlowPfcPausePauseClass4PatternFlowPfcPausePauseClass4MetricTagIter {
	obj.patternFlowPfcPausePauseClass4MetricTagSlice = append(obj.patternFlowPfcPausePauseClass4MetricTagSlice, item)
	return obj
}

func (obj *patternFlowPfcPausePauseClass4) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass4.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowPfcPausePauseClass4.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowPfcPausePauseClass4MetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowPfcPausePauseClass4) setDefault() {
	var choices_set int = 0
	var choice PatternFlowPfcPausePauseClass4ChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowPfcPausePauseClass4Choice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowPfcPausePauseClass4Choice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowPfcPausePauseClass4Choice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowPfcPausePauseClass4Choice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowPfcPausePauseClass4Choice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowPfcPausePauseClass4")
			}
		} else {
			intVal := otg.PatternFlowPfcPausePauseClass4_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowPfcPausePauseClass4_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
