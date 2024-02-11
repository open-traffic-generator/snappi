package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPausePauseClass7 *****
type patternFlowPfcPausePauseClass7 struct {
	validation
	obj              *otg.PatternFlowPfcPausePauseClass7
	marshaller       marshalPatternFlowPfcPausePauseClass7
	unMarshaller     unMarshalPatternFlowPfcPausePauseClass7
	incrementHolder  PatternFlowPfcPausePauseClass7Counter
	decrementHolder  PatternFlowPfcPausePauseClass7Counter
	metricTagsHolder PatternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter
}

func NewPatternFlowPfcPausePauseClass7() PatternFlowPfcPausePauseClass7 {
	obj := patternFlowPfcPausePauseClass7{obj: &otg.PatternFlowPfcPausePauseClass7{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPausePauseClass7) msg() *otg.PatternFlowPfcPausePauseClass7 {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass7) setMsg(msg *otg.PatternFlowPfcPausePauseClass7) PatternFlowPfcPausePauseClass7 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPausePauseClass7 struct {
	obj *patternFlowPfcPausePauseClass7
}

type marshalPatternFlowPfcPausePauseClass7 interface {
	// ToProto marshals PatternFlowPfcPausePauseClass7 to protobuf object *otg.PatternFlowPfcPausePauseClass7
	ToProto() (*otg.PatternFlowPfcPausePauseClass7, error)
	// ToPbText marshals PatternFlowPfcPausePauseClass7 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPausePauseClass7 to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPausePauseClass7 to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPfcPausePauseClass7 struct {
	obj *patternFlowPfcPausePauseClass7
}

type unMarshalPatternFlowPfcPausePauseClass7 interface {
	// FromProto unmarshals PatternFlowPfcPausePauseClass7 from protobuf object *otg.PatternFlowPfcPausePauseClass7
	FromProto(msg *otg.PatternFlowPfcPausePauseClass7) (PatternFlowPfcPausePauseClass7, error)
	// FromPbText unmarshals PatternFlowPfcPausePauseClass7 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPausePauseClass7 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPausePauseClass7 from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPausePauseClass7) Marshal() marshalPatternFlowPfcPausePauseClass7 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPausePauseClass7{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPausePauseClass7) Unmarshal() unMarshalPatternFlowPfcPausePauseClass7 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPausePauseClass7{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPausePauseClass7) ToProto() (*otg.PatternFlowPfcPausePauseClass7, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPausePauseClass7) FromProto(msg *otg.PatternFlowPfcPausePauseClass7) (PatternFlowPfcPausePauseClass7, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPausePauseClass7) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass7) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass7) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass7) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass7) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass7) FromJson(value string) error {
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

func (obj *patternFlowPfcPausePauseClass7) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass7) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass7) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPausePauseClass7) Clone() (PatternFlowPfcPausePauseClass7, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPausePauseClass7()
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

func (obj *patternFlowPfcPausePauseClass7) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowPfcPausePauseClass7 is pause class 7
type PatternFlowPfcPausePauseClass7 interface {
	Validation
	// msg marshals PatternFlowPfcPausePauseClass7 to protobuf object *otg.PatternFlowPfcPausePauseClass7
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPausePauseClass7
	// setMsg unmarshals PatternFlowPfcPausePauseClass7 from protobuf object *otg.PatternFlowPfcPausePauseClass7
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPausePauseClass7) PatternFlowPfcPausePauseClass7
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPausePauseClass7
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPausePauseClass7
	// validate validates PatternFlowPfcPausePauseClass7
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPausePauseClass7, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowPfcPausePauseClass7ChoiceEnum, set in PatternFlowPfcPausePauseClass7
	Choice() PatternFlowPfcPausePauseClass7ChoiceEnum
	// setChoice assigns PatternFlowPfcPausePauseClass7ChoiceEnum provided by user to PatternFlowPfcPausePauseClass7
	setChoice(value PatternFlowPfcPausePauseClass7ChoiceEnum) PatternFlowPfcPausePauseClass7
	// HasChoice checks if Choice has been set in PatternFlowPfcPausePauseClass7
	HasChoice() bool
	// Value returns uint32, set in PatternFlowPfcPausePauseClass7.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowPfcPausePauseClass7
	SetValue(value uint32) PatternFlowPfcPausePauseClass7
	// HasValue checks if Value has been set in PatternFlowPfcPausePauseClass7
	HasValue() bool
	// Values returns []uint32, set in PatternFlowPfcPausePauseClass7.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowPfcPausePauseClass7
	SetValues(value []uint32) PatternFlowPfcPausePauseClass7
	// Increment returns PatternFlowPfcPausePauseClass7Counter, set in PatternFlowPfcPausePauseClass7.
	// PatternFlowPfcPausePauseClass7Counter is integer counter pattern
	Increment() PatternFlowPfcPausePauseClass7Counter
	// SetIncrement assigns PatternFlowPfcPausePauseClass7Counter provided by user to PatternFlowPfcPausePauseClass7.
	// PatternFlowPfcPausePauseClass7Counter is integer counter pattern
	SetIncrement(value PatternFlowPfcPausePauseClass7Counter) PatternFlowPfcPausePauseClass7
	// HasIncrement checks if Increment has been set in PatternFlowPfcPausePauseClass7
	HasIncrement() bool
	// Decrement returns PatternFlowPfcPausePauseClass7Counter, set in PatternFlowPfcPausePauseClass7.
	// PatternFlowPfcPausePauseClass7Counter is integer counter pattern
	Decrement() PatternFlowPfcPausePauseClass7Counter
	// SetDecrement assigns PatternFlowPfcPausePauseClass7Counter provided by user to PatternFlowPfcPausePauseClass7.
	// PatternFlowPfcPausePauseClass7Counter is integer counter pattern
	SetDecrement(value PatternFlowPfcPausePauseClass7Counter) PatternFlowPfcPausePauseClass7
	// HasDecrement checks if Decrement has been set in PatternFlowPfcPausePauseClass7
	HasDecrement() bool
	// MetricTags returns PatternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIterIter, set in PatternFlowPfcPausePauseClass7
	MetricTags() PatternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter
	setNil()
}

type PatternFlowPfcPausePauseClass7ChoiceEnum string

// Enum of Choice on PatternFlowPfcPausePauseClass7
var PatternFlowPfcPausePauseClass7Choice = struct {
	VALUE     PatternFlowPfcPausePauseClass7ChoiceEnum
	VALUES    PatternFlowPfcPausePauseClass7ChoiceEnum
	INCREMENT PatternFlowPfcPausePauseClass7ChoiceEnum
	DECREMENT PatternFlowPfcPausePauseClass7ChoiceEnum
}{
	VALUE:     PatternFlowPfcPausePauseClass7ChoiceEnum("value"),
	VALUES:    PatternFlowPfcPausePauseClass7ChoiceEnum("values"),
	INCREMENT: PatternFlowPfcPausePauseClass7ChoiceEnum("increment"),
	DECREMENT: PatternFlowPfcPausePauseClass7ChoiceEnum("decrement"),
}

func (obj *patternFlowPfcPausePauseClass7) Choice() PatternFlowPfcPausePauseClass7ChoiceEnum {
	return PatternFlowPfcPausePauseClass7ChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowPfcPausePauseClass7) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowPfcPausePauseClass7) setChoice(value PatternFlowPfcPausePauseClass7ChoiceEnum) PatternFlowPfcPausePauseClass7 {
	intValue, ok := otg.PatternFlowPfcPausePauseClass7_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowPfcPausePauseClass7ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowPfcPausePauseClass7_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowPfcPausePauseClass7Choice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowPfcPausePauseClass7Choice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowPfcPausePauseClass7Choice.INCREMENT {
		obj.obj.Increment = NewPatternFlowPfcPausePauseClass7Counter().msg()
	}

	if value == PatternFlowPfcPausePauseClass7Choice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowPfcPausePauseClass7Counter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowPfcPausePauseClass7) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowPfcPausePauseClass7Choice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowPfcPausePauseClass7) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowPfcPausePauseClass7 object
func (obj *patternFlowPfcPausePauseClass7) SetValue(value uint32) PatternFlowPfcPausePauseClass7 {
	obj.setChoice(PatternFlowPfcPausePauseClass7Choice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowPfcPausePauseClass7) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowPfcPausePauseClass7 object
func (obj *patternFlowPfcPausePauseClass7) SetValues(value []uint32) PatternFlowPfcPausePauseClass7 {
	obj.setChoice(PatternFlowPfcPausePauseClass7Choice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowPfcPausePauseClass7Counter
func (obj *patternFlowPfcPausePauseClass7) Increment() PatternFlowPfcPausePauseClass7Counter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowPfcPausePauseClass7Choice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowPfcPausePauseClass7Counter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowPfcPausePauseClass7Counter
func (obj *patternFlowPfcPausePauseClass7) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowPfcPausePauseClass7Counter value in the PatternFlowPfcPausePauseClass7 object
func (obj *patternFlowPfcPausePauseClass7) SetIncrement(value PatternFlowPfcPausePauseClass7Counter) PatternFlowPfcPausePauseClass7 {
	obj.setChoice(PatternFlowPfcPausePauseClass7Choice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowPfcPausePauseClass7Counter
func (obj *patternFlowPfcPausePauseClass7) Decrement() PatternFlowPfcPausePauseClass7Counter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowPfcPausePauseClass7Choice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowPfcPausePauseClass7Counter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowPfcPausePauseClass7Counter
func (obj *patternFlowPfcPausePauseClass7) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowPfcPausePauseClass7Counter value in the PatternFlowPfcPausePauseClass7 object
func (obj *patternFlowPfcPausePauseClass7) SetDecrement(value PatternFlowPfcPausePauseClass7Counter) PatternFlowPfcPausePauseClass7 {
	obj.setChoice(PatternFlowPfcPausePauseClass7Choice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowPfcPausePauseClass7MetricTag
func (obj *patternFlowPfcPausePauseClass7) MetricTags() PatternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowPfcPausePauseClass7MetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter struct {
	obj                                          *patternFlowPfcPausePauseClass7
	patternFlowPfcPausePauseClass7MetricTagSlice []PatternFlowPfcPausePauseClass7MetricTag
	fieldPtr                                     *[]*otg.PatternFlowPfcPausePauseClass7MetricTag
}

func newPatternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter(ptr *[]*otg.PatternFlowPfcPausePauseClass7MetricTag) PatternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter {
	return &patternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter{fieldPtr: ptr}
}

type PatternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter interface {
	setMsg(*patternFlowPfcPausePauseClass7) PatternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter
	Items() []PatternFlowPfcPausePauseClass7MetricTag
	Add() PatternFlowPfcPausePauseClass7MetricTag
	Append(items ...PatternFlowPfcPausePauseClass7MetricTag) PatternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter
	Set(index int, newObj PatternFlowPfcPausePauseClass7MetricTag) PatternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter
	Clear() PatternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter
	clearHolderSlice() PatternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter
	appendHolderSlice(item PatternFlowPfcPausePauseClass7MetricTag) PatternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter
}

func (obj *patternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter) setMsg(msg *patternFlowPfcPausePauseClass7) PatternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowPfcPausePauseClass7MetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter) Items() []PatternFlowPfcPausePauseClass7MetricTag {
	return obj.patternFlowPfcPausePauseClass7MetricTagSlice
}

func (obj *patternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter) Add() PatternFlowPfcPausePauseClass7MetricTag {
	newObj := &otg.PatternFlowPfcPausePauseClass7MetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowPfcPausePauseClass7MetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowPfcPausePauseClass7MetricTagSlice = append(obj.patternFlowPfcPausePauseClass7MetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter) Append(items ...PatternFlowPfcPausePauseClass7MetricTag) PatternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowPfcPausePauseClass7MetricTagSlice = append(obj.patternFlowPfcPausePauseClass7MetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter) Set(index int, newObj PatternFlowPfcPausePauseClass7MetricTag) PatternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowPfcPausePauseClass7MetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter) Clear() PatternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowPfcPausePauseClass7MetricTag{}
		obj.patternFlowPfcPausePauseClass7MetricTagSlice = []PatternFlowPfcPausePauseClass7MetricTag{}
	}
	return obj
}
func (obj *patternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter) clearHolderSlice() PatternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter {
	if len(obj.patternFlowPfcPausePauseClass7MetricTagSlice) > 0 {
		obj.patternFlowPfcPausePauseClass7MetricTagSlice = []PatternFlowPfcPausePauseClass7MetricTag{}
	}
	return obj
}
func (obj *patternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter) appendHolderSlice(item PatternFlowPfcPausePauseClass7MetricTag) PatternFlowPfcPausePauseClass7PatternFlowPfcPausePauseClass7MetricTagIter {
	obj.patternFlowPfcPausePauseClass7MetricTagSlice = append(obj.patternFlowPfcPausePauseClass7MetricTagSlice, item)
	return obj
}

func (obj *patternFlowPfcPausePauseClass7) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass7.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowPfcPausePauseClass7.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowPfcPausePauseClass7MetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowPfcPausePauseClass7) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowPfcPausePauseClass7Choice.VALUE)

	}

}
