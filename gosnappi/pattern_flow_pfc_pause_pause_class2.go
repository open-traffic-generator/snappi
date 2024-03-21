package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPausePauseClass2 *****
type patternFlowPfcPausePauseClass2 struct {
	validation
	obj              *otg.PatternFlowPfcPausePauseClass2
	marshaller       marshalPatternFlowPfcPausePauseClass2
	unMarshaller     unMarshalPatternFlowPfcPausePauseClass2
	incrementHolder  PatternFlowPfcPausePauseClass2Counter
	decrementHolder  PatternFlowPfcPausePauseClass2Counter
	metricTagsHolder PatternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter
}

func NewPatternFlowPfcPausePauseClass2() PatternFlowPfcPausePauseClass2 {
	obj := patternFlowPfcPausePauseClass2{obj: &otg.PatternFlowPfcPausePauseClass2{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPausePauseClass2) msg() *otg.PatternFlowPfcPausePauseClass2 {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass2) setMsg(msg *otg.PatternFlowPfcPausePauseClass2) PatternFlowPfcPausePauseClass2 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPausePauseClass2 struct {
	obj *patternFlowPfcPausePauseClass2
}

type marshalPatternFlowPfcPausePauseClass2 interface {
	// ToProto marshals PatternFlowPfcPausePauseClass2 to protobuf object *otg.PatternFlowPfcPausePauseClass2
	ToProto() (*otg.PatternFlowPfcPausePauseClass2, error)
	// ToPbText marshals PatternFlowPfcPausePauseClass2 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPausePauseClass2 to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPausePauseClass2 to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPfcPausePauseClass2 struct {
	obj *patternFlowPfcPausePauseClass2
}

type unMarshalPatternFlowPfcPausePauseClass2 interface {
	// FromProto unmarshals PatternFlowPfcPausePauseClass2 from protobuf object *otg.PatternFlowPfcPausePauseClass2
	FromProto(msg *otg.PatternFlowPfcPausePauseClass2) (PatternFlowPfcPausePauseClass2, error)
	// FromPbText unmarshals PatternFlowPfcPausePauseClass2 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPausePauseClass2 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPausePauseClass2 from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPausePauseClass2) Marshal() marshalPatternFlowPfcPausePauseClass2 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPausePauseClass2{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPausePauseClass2) Unmarshal() unMarshalPatternFlowPfcPausePauseClass2 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPausePauseClass2{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPausePauseClass2) ToProto() (*otg.PatternFlowPfcPausePauseClass2, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPausePauseClass2) FromProto(msg *otg.PatternFlowPfcPausePauseClass2) (PatternFlowPfcPausePauseClass2, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPausePauseClass2) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass2) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass2) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass2) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass2) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass2) FromJson(value string) error {
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

func (obj *patternFlowPfcPausePauseClass2) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass2) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass2) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPausePauseClass2) Clone() (PatternFlowPfcPausePauseClass2, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPausePauseClass2()
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

func (obj *patternFlowPfcPausePauseClass2) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowPfcPausePauseClass2 is pause class 2
type PatternFlowPfcPausePauseClass2 interface {
	Validation
	// msg marshals PatternFlowPfcPausePauseClass2 to protobuf object *otg.PatternFlowPfcPausePauseClass2
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPausePauseClass2
	// setMsg unmarshals PatternFlowPfcPausePauseClass2 from protobuf object *otg.PatternFlowPfcPausePauseClass2
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPausePauseClass2) PatternFlowPfcPausePauseClass2
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPausePauseClass2
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPausePauseClass2
	// validate validates PatternFlowPfcPausePauseClass2
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPausePauseClass2, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowPfcPausePauseClass2ChoiceEnum, set in PatternFlowPfcPausePauseClass2
	Choice() PatternFlowPfcPausePauseClass2ChoiceEnum
	// setChoice assigns PatternFlowPfcPausePauseClass2ChoiceEnum provided by user to PatternFlowPfcPausePauseClass2
	setChoice(value PatternFlowPfcPausePauseClass2ChoiceEnum) PatternFlowPfcPausePauseClass2
	// HasChoice checks if Choice has been set in PatternFlowPfcPausePauseClass2
	HasChoice() bool
	// Value returns uint32, set in PatternFlowPfcPausePauseClass2.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowPfcPausePauseClass2
	SetValue(value uint32) PatternFlowPfcPausePauseClass2
	// HasValue checks if Value has been set in PatternFlowPfcPausePauseClass2
	HasValue() bool
	// Values returns []uint32, set in PatternFlowPfcPausePauseClass2.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowPfcPausePauseClass2
	SetValues(value []uint32) PatternFlowPfcPausePauseClass2
	// Increment returns PatternFlowPfcPausePauseClass2Counter, set in PatternFlowPfcPausePauseClass2.
	// PatternFlowPfcPausePauseClass2Counter is integer counter pattern
	Increment() PatternFlowPfcPausePauseClass2Counter
	// SetIncrement assigns PatternFlowPfcPausePauseClass2Counter provided by user to PatternFlowPfcPausePauseClass2.
	// PatternFlowPfcPausePauseClass2Counter is integer counter pattern
	SetIncrement(value PatternFlowPfcPausePauseClass2Counter) PatternFlowPfcPausePauseClass2
	// HasIncrement checks if Increment has been set in PatternFlowPfcPausePauseClass2
	HasIncrement() bool
	// Decrement returns PatternFlowPfcPausePauseClass2Counter, set in PatternFlowPfcPausePauseClass2.
	// PatternFlowPfcPausePauseClass2Counter is integer counter pattern
	Decrement() PatternFlowPfcPausePauseClass2Counter
	// SetDecrement assigns PatternFlowPfcPausePauseClass2Counter provided by user to PatternFlowPfcPausePauseClass2.
	// PatternFlowPfcPausePauseClass2Counter is integer counter pattern
	SetDecrement(value PatternFlowPfcPausePauseClass2Counter) PatternFlowPfcPausePauseClass2
	// HasDecrement checks if Decrement has been set in PatternFlowPfcPausePauseClass2
	HasDecrement() bool
	// MetricTags returns PatternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIterIter, set in PatternFlowPfcPausePauseClass2
	MetricTags() PatternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter
	setNil()
}

type PatternFlowPfcPausePauseClass2ChoiceEnum string

// Enum of Choice on PatternFlowPfcPausePauseClass2
var PatternFlowPfcPausePauseClass2Choice = struct {
	VALUE     PatternFlowPfcPausePauseClass2ChoiceEnum
	VALUES    PatternFlowPfcPausePauseClass2ChoiceEnum
	INCREMENT PatternFlowPfcPausePauseClass2ChoiceEnum
	DECREMENT PatternFlowPfcPausePauseClass2ChoiceEnum
}{
	VALUE:     PatternFlowPfcPausePauseClass2ChoiceEnum("value"),
	VALUES:    PatternFlowPfcPausePauseClass2ChoiceEnum("values"),
	INCREMENT: PatternFlowPfcPausePauseClass2ChoiceEnum("increment"),
	DECREMENT: PatternFlowPfcPausePauseClass2ChoiceEnum("decrement"),
}

func (obj *patternFlowPfcPausePauseClass2) Choice() PatternFlowPfcPausePauseClass2ChoiceEnum {
	return PatternFlowPfcPausePauseClass2ChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowPfcPausePauseClass2) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowPfcPausePauseClass2) setChoice(value PatternFlowPfcPausePauseClass2ChoiceEnum) PatternFlowPfcPausePauseClass2 {
	intValue, ok := otg.PatternFlowPfcPausePauseClass2_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowPfcPausePauseClass2ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowPfcPausePauseClass2_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowPfcPausePauseClass2Choice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowPfcPausePauseClass2Choice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowPfcPausePauseClass2Choice.INCREMENT {
		obj.obj.Increment = NewPatternFlowPfcPausePauseClass2Counter().msg()
	}

	if value == PatternFlowPfcPausePauseClass2Choice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowPfcPausePauseClass2Counter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowPfcPausePauseClass2) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowPfcPausePauseClass2Choice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowPfcPausePauseClass2) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowPfcPausePauseClass2 object
func (obj *patternFlowPfcPausePauseClass2) SetValue(value uint32) PatternFlowPfcPausePauseClass2 {
	obj.setChoice(PatternFlowPfcPausePauseClass2Choice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowPfcPausePauseClass2) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowPfcPausePauseClass2 object
func (obj *patternFlowPfcPausePauseClass2) SetValues(value []uint32) PatternFlowPfcPausePauseClass2 {
	obj.setChoice(PatternFlowPfcPausePauseClass2Choice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowPfcPausePauseClass2Counter
func (obj *patternFlowPfcPausePauseClass2) Increment() PatternFlowPfcPausePauseClass2Counter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowPfcPausePauseClass2Choice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowPfcPausePauseClass2Counter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowPfcPausePauseClass2Counter
func (obj *patternFlowPfcPausePauseClass2) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowPfcPausePauseClass2Counter value in the PatternFlowPfcPausePauseClass2 object
func (obj *patternFlowPfcPausePauseClass2) SetIncrement(value PatternFlowPfcPausePauseClass2Counter) PatternFlowPfcPausePauseClass2 {
	obj.setChoice(PatternFlowPfcPausePauseClass2Choice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowPfcPausePauseClass2Counter
func (obj *patternFlowPfcPausePauseClass2) Decrement() PatternFlowPfcPausePauseClass2Counter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowPfcPausePauseClass2Choice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowPfcPausePauseClass2Counter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowPfcPausePauseClass2Counter
func (obj *patternFlowPfcPausePauseClass2) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowPfcPausePauseClass2Counter value in the PatternFlowPfcPausePauseClass2 object
func (obj *patternFlowPfcPausePauseClass2) SetDecrement(value PatternFlowPfcPausePauseClass2Counter) PatternFlowPfcPausePauseClass2 {
	obj.setChoice(PatternFlowPfcPausePauseClass2Choice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowPfcPausePauseClass2MetricTag
func (obj *patternFlowPfcPausePauseClass2) MetricTags() PatternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowPfcPausePauseClass2MetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter struct {
	obj                                          *patternFlowPfcPausePauseClass2
	patternFlowPfcPausePauseClass2MetricTagSlice []PatternFlowPfcPausePauseClass2MetricTag
	fieldPtr                                     *[]*otg.PatternFlowPfcPausePauseClass2MetricTag
}

func newPatternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter(ptr *[]*otg.PatternFlowPfcPausePauseClass2MetricTag) PatternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter {
	return &patternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter{fieldPtr: ptr}
}

type PatternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter interface {
	setMsg(*patternFlowPfcPausePauseClass2) PatternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter
	Items() []PatternFlowPfcPausePauseClass2MetricTag
	Add() PatternFlowPfcPausePauseClass2MetricTag
	Append(items ...PatternFlowPfcPausePauseClass2MetricTag) PatternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter
	Set(index int, newObj PatternFlowPfcPausePauseClass2MetricTag) PatternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter
	Clear() PatternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter
	clearHolderSlice() PatternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter
	appendHolderSlice(item PatternFlowPfcPausePauseClass2MetricTag) PatternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter
}

func (obj *patternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter) setMsg(msg *patternFlowPfcPausePauseClass2) PatternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowPfcPausePauseClass2MetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter) Items() []PatternFlowPfcPausePauseClass2MetricTag {
	return obj.patternFlowPfcPausePauseClass2MetricTagSlice
}

func (obj *patternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter) Add() PatternFlowPfcPausePauseClass2MetricTag {
	newObj := &otg.PatternFlowPfcPausePauseClass2MetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowPfcPausePauseClass2MetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowPfcPausePauseClass2MetricTagSlice = append(obj.patternFlowPfcPausePauseClass2MetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter) Append(items ...PatternFlowPfcPausePauseClass2MetricTag) PatternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowPfcPausePauseClass2MetricTagSlice = append(obj.patternFlowPfcPausePauseClass2MetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter) Set(index int, newObj PatternFlowPfcPausePauseClass2MetricTag) PatternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowPfcPausePauseClass2MetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter) Clear() PatternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowPfcPausePauseClass2MetricTag{}
		obj.patternFlowPfcPausePauseClass2MetricTagSlice = []PatternFlowPfcPausePauseClass2MetricTag{}
	}
	return obj
}
func (obj *patternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter) clearHolderSlice() PatternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter {
	if len(obj.patternFlowPfcPausePauseClass2MetricTagSlice) > 0 {
		obj.patternFlowPfcPausePauseClass2MetricTagSlice = []PatternFlowPfcPausePauseClass2MetricTag{}
	}
	return obj
}
func (obj *patternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter) appendHolderSlice(item PatternFlowPfcPausePauseClass2MetricTag) PatternFlowPfcPausePauseClass2PatternFlowPfcPausePauseClass2MetricTagIter {
	obj.patternFlowPfcPausePauseClass2MetricTagSlice = append(obj.patternFlowPfcPausePauseClass2MetricTagSlice, item)
	return obj
}

func (obj *patternFlowPfcPausePauseClass2) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass2.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowPfcPausePauseClass2.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowPfcPausePauseClass2MetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowPfcPausePauseClass2) setDefault() {
	var choices_set int = 0
	var choice PatternFlowPfcPausePauseClass2ChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowPfcPausePauseClass2Choice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowPfcPausePauseClass2Choice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowPfcPausePauseClass2Choice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowPfcPausePauseClass2Choice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowPfcPausePauseClass2Choice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowPfcPausePauseClass2")
			}
		} else {
			intVal := otg.PatternFlowPfcPausePauseClass2_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowPfcPausePauseClass2_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
