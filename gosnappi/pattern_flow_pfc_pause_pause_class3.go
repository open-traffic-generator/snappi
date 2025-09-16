package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPausePauseClass3 *****
type patternFlowPfcPausePauseClass3 struct {
	validation
	obj              *otg.PatternFlowPfcPausePauseClass3
	marshaller       marshalPatternFlowPfcPausePauseClass3
	unMarshaller     unMarshalPatternFlowPfcPausePauseClass3
	incrementHolder  PatternFlowPfcPausePauseClass3Counter
	decrementHolder  PatternFlowPfcPausePauseClass3Counter
	metricTagsHolder PatternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter
}

func NewPatternFlowPfcPausePauseClass3() PatternFlowPfcPausePauseClass3 {
	obj := patternFlowPfcPausePauseClass3{obj: &otg.PatternFlowPfcPausePauseClass3{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPausePauseClass3) msg() *otg.PatternFlowPfcPausePauseClass3 {
	return obj.obj
}

func (obj *patternFlowPfcPausePauseClass3) setMsg(msg *otg.PatternFlowPfcPausePauseClass3) PatternFlowPfcPausePauseClass3 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPausePauseClass3 struct {
	obj *patternFlowPfcPausePauseClass3
}

type marshalPatternFlowPfcPausePauseClass3 interface {
	// ToProto marshals PatternFlowPfcPausePauseClass3 to protobuf object *otg.PatternFlowPfcPausePauseClass3
	ToProto() (*otg.PatternFlowPfcPausePauseClass3, error)
	// ToPbText marshals PatternFlowPfcPausePauseClass3 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPausePauseClass3 to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPausePauseClass3 to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPfcPausePauseClass3 struct {
	obj *patternFlowPfcPausePauseClass3
}

type unMarshalPatternFlowPfcPausePauseClass3 interface {
	// FromProto unmarshals PatternFlowPfcPausePauseClass3 from protobuf object *otg.PatternFlowPfcPausePauseClass3
	FromProto(msg *otg.PatternFlowPfcPausePauseClass3) (PatternFlowPfcPausePauseClass3, error)
	// FromPbText unmarshals PatternFlowPfcPausePauseClass3 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPausePauseClass3 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPausePauseClass3 from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPausePauseClass3) Marshal() marshalPatternFlowPfcPausePauseClass3 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPausePauseClass3{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPausePauseClass3) Unmarshal() unMarshalPatternFlowPfcPausePauseClass3 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPausePauseClass3{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPausePauseClass3) ToProto() (*otg.PatternFlowPfcPausePauseClass3, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPausePauseClass3) FromProto(msg *otg.PatternFlowPfcPausePauseClass3) (PatternFlowPfcPausePauseClass3, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPausePauseClass3) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass3) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass3) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass3) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPausePauseClass3) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPausePauseClass3) FromJson(value string) error {
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

func (obj *patternFlowPfcPausePauseClass3) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass3) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPausePauseClass3) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPausePauseClass3) Clone() (PatternFlowPfcPausePauseClass3, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPausePauseClass3()
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

func (obj *patternFlowPfcPausePauseClass3) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowPfcPausePauseClass3 is pause class 3
type PatternFlowPfcPausePauseClass3 interface {
	Validation
	// msg marshals PatternFlowPfcPausePauseClass3 to protobuf object *otg.PatternFlowPfcPausePauseClass3
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPausePauseClass3
	// setMsg unmarshals PatternFlowPfcPausePauseClass3 from protobuf object *otg.PatternFlowPfcPausePauseClass3
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPausePauseClass3) PatternFlowPfcPausePauseClass3
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPausePauseClass3
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPausePauseClass3
	// validate validates PatternFlowPfcPausePauseClass3
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPausePauseClass3, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowPfcPausePauseClass3ChoiceEnum, set in PatternFlowPfcPausePauseClass3
	Choice() PatternFlowPfcPausePauseClass3ChoiceEnum
	// setChoice assigns PatternFlowPfcPausePauseClass3ChoiceEnum provided by user to PatternFlowPfcPausePauseClass3
	setChoice(value PatternFlowPfcPausePauseClass3ChoiceEnum) PatternFlowPfcPausePauseClass3
	// HasChoice checks if Choice has been set in PatternFlowPfcPausePauseClass3
	HasChoice() bool
	// Value returns uint32, set in PatternFlowPfcPausePauseClass3.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowPfcPausePauseClass3
	SetValue(value uint32) PatternFlowPfcPausePauseClass3
	// HasValue checks if Value has been set in PatternFlowPfcPausePauseClass3
	HasValue() bool
	// Values returns []uint32, set in PatternFlowPfcPausePauseClass3.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowPfcPausePauseClass3
	SetValues(value []uint32) PatternFlowPfcPausePauseClass3
	// Increment returns PatternFlowPfcPausePauseClass3Counter, set in PatternFlowPfcPausePauseClass3.
	// PatternFlowPfcPausePauseClass3Counter is integer counter pattern
	Increment() PatternFlowPfcPausePauseClass3Counter
	// SetIncrement assigns PatternFlowPfcPausePauseClass3Counter provided by user to PatternFlowPfcPausePauseClass3.
	// PatternFlowPfcPausePauseClass3Counter is integer counter pattern
	SetIncrement(value PatternFlowPfcPausePauseClass3Counter) PatternFlowPfcPausePauseClass3
	// HasIncrement checks if Increment has been set in PatternFlowPfcPausePauseClass3
	HasIncrement() bool
	// Decrement returns PatternFlowPfcPausePauseClass3Counter, set in PatternFlowPfcPausePauseClass3.
	// PatternFlowPfcPausePauseClass3Counter is integer counter pattern
	Decrement() PatternFlowPfcPausePauseClass3Counter
	// SetDecrement assigns PatternFlowPfcPausePauseClass3Counter provided by user to PatternFlowPfcPausePauseClass3.
	// PatternFlowPfcPausePauseClass3Counter is integer counter pattern
	SetDecrement(value PatternFlowPfcPausePauseClass3Counter) PatternFlowPfcPausePauseClass3
	// HasDecrement checks if Decrement has been set in PatternFlowPfcPausePauseClass3
	HasDecrement() bool
	// MetricTags returns PatternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIterIter, set in PatternFlowPfcPausePauseClass3
	MetricTags() PatternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter
	setNil()
}

type PatternFlowPfcPausePauseClass3ChoiceEnum string

// Enum of Choice on PatternFlowPfcPausePauseClass3
var PatternFlowPfcPausePauseClass3Choice = struct {
	VALUE     PatternFlowPfcPausePauseClass3ChoiceEnum
	VALUES    PatternFlowPfcPausePauseClass3ChoiceEnum
	INCREMENT PatternFlowPfcPausePauseClass3ChoiceEnum
	DECREMENT PatternFlowPfcPausePauseClass3ChoiceEnum
}{
	VALUE:     PatternFlowPfcPausePauseClass3ChoiceEnum("value"),
	VALUES:    PatternFlowPfcPausePauseClass3ChoiceEnum("values"),
	INCREMENT: PatternFlowPfcPausePauseClass3ChoiceEnum("increment"),
	DECREMENT: PatternFlowPfcPausePauseClass3ChoiceEnum("decrement"),
}

func (obj *patternFlowPfcPausePauseClass3) Choice() PatternFlowPfcPausePauseClass3ChoiceEnum {
	return PatternFlowPfcPausePauseClass3ChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowPfcPausePauseClass3) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowPfcPausePauseClass3) setChoice(value PatternFlowPfcPausePauseClass3ChoiceEnum) PatternFlowPfcPausePauseClass3 {
	intValue, ok := otg.PatternFlowPfcPausePauseClass3_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowPfcPausePauseClass3ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowPfcPausePauseClass3_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowPfcPausePauseClass3Choice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowPfcPausePauseClass3Choice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowPfcPausePauseClass3Choice.INCREMENT {
		obj.obj.Increment = NewPatternFlowPfcPausePauseClass3Counter().msg()
	}

	if value == PatternFlowPfcPausePauseClass3Choice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowPfcPausePauseClass3Counter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowPfcPausePauseClass3) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowPfcPausePauseClass3Choice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowPfcPausePauseClass3) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowPfcPausePauseClass3 object
func (obj *patternFlowPfcPausePauseClass3) SetValue(value uint32) PatternFlowPfcPausePauseClass3 {
	obj.setChoice(PatternFlowPfcPausePauseClass3Choice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowPfcPausePauseClass3) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowPfcPausePauseClass3 object
func (obj *patternFlowPfcPausePauseClass3) SetValues(value []uint32) PatternFlowPfcPausePauseClass3 {
	obj.setChoice(PatternFlowPfcPausePauseClass3Choice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowPfcPausePauseClass3Counter
func (obj *patternFlowPfcPausePauseClass3) Increment() PatternFlowPfcPausePauseClass3Counter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowPfcPausePauseClass3Choice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowPfcPausePauseClass3Counter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowPfcPausePauseClass3Counter
func (obj *patternFlowPfcPausePauseClass3) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowPfcPausePauseClass3Counter value in the PatternFlowPfcPausePauseClass3 object
func (obj *patternFlowPfcPausePauseClass3) SetIncrement(value PatternFlowPfcPausePauseClass3Counter) PatternFlowPfcPausePauseClass3 {
	obj.setChoice(PatternFlowPfcPausePauseClass3Choice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowPfcPausePauseClass3Counter
func (obj *patternFlowPfcPausePauseClass3) Decrement() PatternFlowPfcPausePauseClass3Counter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowPfcPausePauseClass3Choice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowPfcPausePauseClass3Counter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowPfcPausePauseClass3Counter
func (obj *patternFlowPfcPausePauseClass3) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowPfcPausePauseClass3Counter value in the PatternFlowPfcPausePauseClass3 object
func (obj *patternFlowPfcPausePauseClass3) SetDecrement(value PatternFlowPfcPausePauseClass3Counter) PatternFlowPfcPausePauseClass3 {
	obj.setChoice(PatternFlowPfcPausePauseClass3Choice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowPfcPausePauseClass3MetricTag
func (obj *patternFlowPfcPausePauseClass3) MetricTags() PatternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowPfcPausePauseClass3MetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter struct {
	obj                                          *patternFlowPfcPausePauseClass3
	patternFlowPfcPausePauseClass3MetricTagSlice []PatternFlowPfcPausePauseClass3MetricTag
	fieldPtr                                     *[]*otg.PatternFlowPfcPausePauseClass3MetricTag
}

func newPatternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter(ptr *[]*otg.PatternFlowPfcPausePauseClass3MetricTag) PatternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter {
	return &patternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter{fieldPtr: ptr}
}

type PatternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter interface {
	setMsg(*patternFlowPfcPausePauseClass3) PatternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter
	Items() []PatternFlowPfcPausePauseClass3MetricTag
	Add() PatternFlowPfcPausePauseClass3MetricTag
	Append(items ...PatternFlowPfcPausePauseClass3MetricTag) PatternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter
	Set(index int, newObj PatternFlowPfcPausePauseClass3MetricTag) PatternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter
	Clear() PatternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter
	clearHolderSlice() PatternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter
	appendHolderSlice(item PatternFlowPfcPausePauseClass3MetricTag) PatternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter
}

func (obj *patternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter) setMsg(msg *patternFlowPfcPausePauseClass3) PatternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowPfcPausePauseClass3MetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter) Items() []PatternFlowPfcPausePauseClass3MetricTag {
	return obj.patternFlowPfcPausePauseClass3MetricTagSlice
}

func (obj *patternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter) Add() PatternFlowPfcPausePauseClass3MetricTag {
	newObj := &otg.PatternFlowPfcPausePauseClass3MetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowPfcPausePauseClass3MetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowPfcPausePauseClass3MetricTagSlice = append(obj.patternFlowPfcPausePauseClass3MetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter) Append(items ...PatternFlowPfcPausePauseClass3MetricTag) PatternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowPfcPausePauseClass3MetricTagSlice = append(obj.patternFlowPfcPausePauseClass3MetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter) Set(index int, newObj PatternFlowPfcPausePauseClass3MetricTag) PatternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowPfcPausePauseClass3MetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter) Clear() PatternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowPfcPausePauseClass3MetricTag{}
		obj.patternFlowPfcPausePauseClass3MetricTagSlice = []PatternFlowPfcPausePauseClass3MetricTag{}
	}
	return obj
}
func (obj *patternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter) clearHolderSlice() PatternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter {
	if len(obj.patternFlowPfcPausePauseClass3MetricTagSlice) > 0 {
		obj.patternFlowPfcPausePauseClass3MetricTagSlice = []PatternFlowPfcPausePauseClass3MetricTag{}
	}
	return obj
}
func (obj *patternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter) appendHolderSlice(item PatternFlowPfcPausePauseClass3MetricTag) PatternFlowPfcPausePauseClass3PatternFlowPfcPausePauseClass3MetricTagIter {
	obj.patternFlowPfcPausePauseClass3MetricTagSlice = append(obj.patternFlowPfcPausePauseClass3MetricTagSlice, item)
	return obj
}

func (obj *patternFlowPfcPausePauseClass3) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPausePauseClass3.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowPfcPausePauseClass3.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowPfcPausePauseClass3MetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowPfcPausePauseClass3) setDefault() {
	var choices_set int = 0
	var choice PatternFlowPfcPausePauseClass3ChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowPfcPausePauseClass3Choice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowPfcPausePauseClass3Choice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowPfcPausePauseClass3Choice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowPfcPausePauseClass3Choice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowPfcPausePauseClass3Choice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowPfcPausePauseClass3")
			}
		} else {
			intVal := otg.PatternFlowPfcPausePauseClass3_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowPfcPausePauseClass3_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
