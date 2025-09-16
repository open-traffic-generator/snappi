package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPausePauseClass6 *****
type patternFlowPfcPausePauseClass6 struct {
	validation
	obj              *otg.PatternFlowPfcPausePauseClass6
	marshaller       marshalPatternFlowPfcPausePauseClass6
	unMarshaller     unMarshalPatternFlowPfcPausePauseClass6
	incrementHolder  PatternFlowPfcPausePauseClass6Counter
	decrementHolder  PatternFlowPfcPausePauseClass6Counter
	metricTagsHolder PatternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter
}

func NewPatternFlowPfcPausePauseClass6() PatternFlowPfcPausePauseClass6 {
	obj := patternFlowPfcPausePauseClass6{obj: &otg.PatternFlowPfcPausePauseClass6{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPausePauseClass6) msg() *otg.PatternFlowPfcPausePauseClass6 {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass6) setMsg(msg *otg.PatternFlowPfcPausePauseClass6) PatternFlowPfcPausePauseClass6 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPausePauseClass6 struct {
	obj *patternFlowPfcPausePauseClass6
}

type marshalPatternFlowPfcPausePauseClass6 interface {
	// ToProto marshals PatternFlowPfcPausePauseClass6 to protobuf object *otg.PatternFlowPfcPausePauseClass6
	ToProto() (*otg.PatternFlowPfcPausePauseClass6, error)
	// ToPbText marshals PatternFlowPfcPausePauseClass6 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPausePauseClass6 to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPausePauseClass6 to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPfcPausePauseClass6 struct {
	obj *patternFlowPfcPausePauseClass6
}

type unMarshalPatternFlowPfcPausePauseClass6 interface {
	// FromProto unmarshals PatternFlowPfcPausePauseClass6 from protobuf object *otg.PatternFlowPfcPausePauseClass6
	FromProto(msg *otg.PatternFlowPfcPausePauseClass6) (PatternFlowPfcPausePauseClass6, error)
	// FromPbText unmarshals PatternFlowPfcPausePauseClass6 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPausePauseClass6 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPausePauseClass6 from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPausePauseClass6) Marshal() marshalPatternFlowPfcPausePauseClass6 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPausePauseClass6{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPausePauseClass6) Unmarshal() unMarshalPatternFlowPfcPausePauseClass6 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPausePauseClass6{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPausePauseClass6) ToProto() (*otg.PatternFlowPfcPausePauseClass6, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPausePauseClass6) FromProto(msg *otg.PatternFlowPfcPausePauseClass6) (PatternFlowPfcPausePauseClass6, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPausePauseClass6) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass6) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass6) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass6) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass6) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass6) FromJson(value string) error {
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

func (obj *patternFlowPfcPausePauseClass6) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass6) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass6) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPausePauseClass6) Clone() (PatternFlowPfcPausePauseClass6, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPausePauseClass6()
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

func (obj *patternFlowPfcPausePauseClass6) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowPfcPausePauseClass6 is pause class 6
type PatternFlowPfcPausePauseClass6 interface {
	Validation
	// msg marshals PatternFlowPfcPausePauseClass6 to protobuf object *otg.PatternFlowPfcPausePauseClass6
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPausePauseClass6
	// setMsg unmarshals PatternFlowPfcPausePauseClass6 from protobuf object *otg.PatternFlowPfcPausePauseClass6
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPausePauseClass6) PatternFlowPfcPausePauseClass6
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPausePauseClass6
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPausePauseClass6
	// validate validates PatternFlowPfcPausePauseClass6
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPausePauseClass6, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowPfcPausePauseClass6ChoiceEnum, set in PatternFlowPfcPausePauseClass6
	Choice() PatternFlowPfcPausePauseClass6ChoiceEnum
	// setChoice assigns PatternFlowPfcPausePauseClass6ChoiceEnum provided by user to PatternFlowPfcPausePauseClass6
	setChoice(value PatternFlowPfcPausePauseClass6ChoiceEnum) PatternFlowPfcPausePauseClass6
	// HasChoice checks if Choice has been set in PatternFlowPfcPausePauseClass6
	HasChoice() bool
	// Value returns uint32, set in PatternFlowPfcPausePauseClass6.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowPfcPausePauseClass6
	SetValue(value uint32) PatternFlowPfcPausePauseClass6
	// HasValue checks if Value has been set in PatternFlowPfcPausePauseClass6
	HasValue() bool
	// Values returns []uint32, set in PatternFlowPfcPausePauseClass6.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowPfcPausePauseClass6
	SetValues(value []uint32) PatternFlowPfcPausePauseClass6
	// Increment returns PatternFlowPfcPausePauseClass6Counter, set in PatternFlowPfcPausePauseClass6.
	// PatternFlowPfcPausePauseClass6Counter is integer counter pattern
	Increment() PatternFlowPfcPausePauseClass6Counter
	// SetIncrement assigns PatternFlowPfcPausePauseClass6Counter provided by user to PatternFlowPfcPausePauseClass6.
	// PatternFlowPfcPausePauseClass6Counter is integer counter pattern
	SetIncrement(value PatternFlowPfcPausePauseClass6Counter) PatternFlowPfcPausePauseClass6
	// HasIncrement checks if Increment has been set in PatternFlowPfcPausePauseClass6
	HasIncrement() bool
	// Decrement returns PatternFlowPfcPausePauseClass6Counter, set in PatternFlowPfcPausePauseClass6.
	// PatternFlowPfcPausePauseClass6Counter is integer counter pattern
	Decrement() PatternFlowPfcPausePauseClass6Counter
	// SetDecrement assigns PatternFlowPfcPausePauseClass6Counter provided by user to PatternFlowPfcPausePauseClass6.
	// PatternFlowPfcPausePauseClass6Counter is integer counter pattern
	SetDecrement(value PatternFlowPfcPausePauseClass6Counter) PatternFlowPfcPausePauseClass6
	// HasDecrement checks if Decrement has been set in PatternFlowPfcPausePauseClass6
	HasDecrement() bool
	// MetricTags returns PatternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIterIter, set in PatternFlowPfcPausePauseClass6
	MetricTags() PatternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter
	setNil()
}

type PatternFlowPfcPausePauseClass6ChoiceEnum string

// Enum of Choice on PatternFlowPfcPausePauseClass6
var PatternFlowPfcPausePauseClass6Choice = struct {
	VALUE     PatternFlowPfcPausePauseClass6ChoiceEnum
	VALUES    PatternFlowPfcPausePauseClass6ChoiceEnum
	INCREMENT PatternFlowPfcPausePauseClass6ChoiceEnum
	DECREMENT PatternFlowPfcPausePauseClass6ChoiceEnum
}{
	VALUE:     PatternFlowPfcPausePauseClass6ChoiceEnum("value"),
	VALUES:    PatternFlowPfcPausePauseClass6ChoiceEnum("values"),
	INCREMENT: PatternFlowPfcPausePauseClass6ChoiceEnum("increment"),
	DECREMENT: PatternFlowPfcPausePauseClass6ChoiceEnum("decrement"),
}

func (obj *patternFlowPfcPausePauseClass6) Choice() PatternFlowPfcPausePauseClass6ChoiceEnum {
	return PatternFlowPfcPausePauseClass6ChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowPfcPausePauseClass6) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowPfcPausePauseClass6) setChoice(value PatternFlowPfcPausePauseClass6ChoiceEnum) PatternFlowPfcPausePauseClass6 {
	intValue, ok := otg.PatternFlowPfcPausePauseClass6_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowPfcPausePauseClass6ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowPfcPausePauseClass6_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowPfcPausePauseClass6Choice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowPfcPausePauseClass6Choice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowPfcPausePauseClass6Choice.INCREMENT {
		obj.obj.Increment = NewPatternFlowPfcPausePauseClass6Counter().msg()
	}

	if value == PatternFlowPfcPausePauseClass6Choice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowPfcPausePauseClass6Counter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowPfcPausePauseClass6) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowPfcPausePauseClass6Choice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowPfcPausePauseClass6) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowPfcPausePauseClass6 object
func (obj *patternFlowPfcPausePauseClass6) SetValue(value uint32) PatternFlowPfcPausePauseClass6 {
	obj.setChoice(PatternFlowPfcPausePauseClass6Choice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowPfcPausePauseClass6) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowPfcPausePauseClass6 object
func (obj *patternFlowPfcPausePauseClass6) SetValues(value []uint32) PatternFlowPfcPausePauseClass6 {
	obj.setChoice(PatternFlowPfcPausePauseClass6Choice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowPfcPausePauseClass6Counter
func (obj *patternFlowPfcPausePauseClass6) Increment() PatternFlowPfcPausePauseClass6Counter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowPfcPausePauseClass6Choice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowPfcPausePauseClass6Counter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowPfcPausePauseClass6Counter
func (obj *patternFlowPfcPausePauseClass6) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowPfcPausePauseClass6Counter value in the PatternFlowPfcPausePauseClass6 object
func (obj *patternFlowPfcPausePauseClass6) SetIncrement(value PatternFlowPfcPausePauseClass6Counter) PatternFlowPfcPausePauseClass6 {
	obj.setChoice(PatternFlowPfcPausePauseClass6Choice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowPfcPausePauseClass6Counter
func (obj *patternFlowPfcPausePauseClass6) Decrement() PatternFlowPfcPausePauseClass6Counter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowPfcPausePauseClass6Choice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowPfcPausePauseClass6Counter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowPfcPausePauseClass6Counter
func (obj *patternFlowPfcPausePauseClass6) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowPfcPausePauseClass6Counter value in the PatternFlowPfcPausePauseClass6 object
func (obj *patternFlowPfcPausePauseClass6) SetDecrement(value PatternFlowPfcPausePauseClass6Counter) PatternFlowPfcPausePauseClass6 {
	obj.setChoice(PatternFlowPfcPausePauseClass6Choice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowPfcPausePauseClass6MetricTag
func (obj *patternFlowPfcPausePauseClass6) MetricTags() PatternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowPfcPausePauseClass6MetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter struct {
	obj                                          *patternFlowPfcPausePauseClass6
	patternFlowPfcPausePauseClass6MetricTagSlice []PatternFlowPfcPausePauseClass6MetricTag
	fieldPtr                                     *[]*otg.PatternFlowPfcPausePauseClass6MetricTag
}

func newPatternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter(ptr *[]*otg.PatternFlowPfcPausePauseClass6MetricTag) PatternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter {
	return &patternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter{fieldPtr: ptr}
}

type PatternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter interface {
	setMsg(*patternFlowPfcPausePauseClass6) PatternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter
	Items() []PatternFlowPfcPausePauseClass6MetricTag
	Add() PatternFlowPfcPausePauseClass6MetricTag
	Append(items ...PatternFlowPfcPausePauseClass6MetricTag) PatternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter
	Set(index int, newObj PatternFlowPfcPausePauseClass6MetricTag) PatternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter
	Clear() PatternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter
	clearHolderSlice() PatternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter
	appendHolderSlice(item PatternFlowPfcPausePauseClass6MetricTag) PatternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter
}

func (obj *patternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter) setMsg(msg *patternFlowPfcPausePauseClass6) PatternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowPfcPausePauseClass6MetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter) Items() []PatternFlowPfcPausePauseClass6MetricTag {
	return obj.patternFlowPfcPausePauseClass6MetricTagSlice
}

func (obj *patternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter) Add() PatternFlowPfcPausePauseClass6MetricTag {
	newObj := &otg.PatternFlowPfcPausePauseClass6MetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowPfcPausePauseClass6MetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowPfcPausePauseClass6MetricTagSlice = append(obj.patternFlowPfcPausePauseClass6MetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter) Append(items ...PatternFlowPfcPausePauseClass6MetricTag) PatternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowPfcPausePauseClass6MetricTagSlice = append(obj.patternFlowPfcPausePauseClass6MetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter) Set(index int, newObj PatternFlowPfcPausePauseClass6MetricTag) PatternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowPfcPausePauseClass6MetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter) Clear() PatternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowPfcPausePauseClass6MetricTag{}
		obj.patternFlowPfcPausePauseClass6MetricTagSlice = []PatternFlowPfcPausePauseClass6MetricTag{}
	}
	return obj
}
func (obj *patternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter) clearHolderSlice() PatternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter {
	if len(obj.patternFlowPfcPausePauseClass6MetricTagSlice) > 0 {
		obj.patternFlowPfcPausePauseClass6MetricTagSlice = []PatternFlowPfcPausePauseClass6MetricTag{}
	}
	return obj
}
func (obj *patternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter) appendHolderSlice(item PatternFlowPfcPausePauseClass6MetricTag) PatternFlowPfcPausePauseClass6PatternFlowPfcPausePauseClass6MetricTagIter {
	obj.patternFlowPfcPausePauseClass6MetricTagSlice = append(obj.patternFlowPfcPausePauseClass6MetricTagSlice, item)
	return obj
}

func (obj *patternFlowPfcPausePauseClass6) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass6.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowPfcPausePauseClass6.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowPfcPausePauseClass6MetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowPfcPausePauseClass6) setDefault() {
	var choices_set int = 0
	var choice PatternFlowPfcPausePauseClass6ChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowPfcPausePauseClass6Choice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowPfcPausePauseClass6Choice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowPfcPausePauseClass6Choice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowPfcPausePauseClass6Choice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowPfcPausePauseClass6Choice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowPfcPausePauseClass6")
			}
		} else {
			intVal := otg.PatternFlowPfcPausePauseClass6_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowPfcPausePauseClass6_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
