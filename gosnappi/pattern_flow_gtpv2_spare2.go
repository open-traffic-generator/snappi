package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv2Spare2 *****
type patternFlowGtpv2Spare2 struct {
	validation
	obj              *otg.PatternFlowGtpv2Spare2
	marshaller       marshalPatternFlowGtpv2Spare2
	unMarshaller     unMarshalPatternFlowGtpv2Spare2
	incrementHolder  PatternFlowGtpv2Spare2Counter
	decrementHolder  PatternFlowGtpv2Spare2Counter
	metricTagsHolder PatternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter
}

func NewPatternFlowGtpv2Spare2() PatternFlowGtpv2Spare2 {
	obj := patternFlowGtpv2Spare2{obj: &otg.PatternFlowGtpv2Spare2{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv2Spare2) msg() *otg.PatternFlowGtpv2Spare2 {
	return obj.obj
}

func (obj *patternFlowGtpv2Spare2) setMsg(msg *otg.PatternFlowGtpv2Spare2) PatternFlowGtpv2Spare2 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv2Spare2 struct {
	obj *patternFlowGtpv2Spare2
}

type marshalPatternFlowGtpv2Spare2 interface {
	// ToProto marshals PatternFlowGtpv2Spare2 to protobuf object *otg.PatternFlowGtpv2Spare2
	ToProto() (*otg.PatternFlowGtpv2Spare2, error)
	// ToPbText marshals PatternFlowGtpv2Spare2 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv2Spare2 to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv2Spare2 to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv2Spare2 struct {
	obj *patternFlowGtpv2Spare2
}

type unMarshalPatternFlowGtpv2Spare2 interface {
	// FromProto unmarshals PatternFlowGtpv2Spare2 from protobuf object *otg.PatternFlowGtpv2Spare2
	FromProto(msg *otg.PatternFlowGtpv2Spare2) (PatternFlowGtpv2Spare2, error)
	// FromPbText unmarshals PatternFlowGtpv2Spare2 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv2Spare2 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv2Spare2 from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv2Spare2) Marshal() marshalPatternFlowGtpv2Spare2 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv2Spare2{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv2Spare2) Unmarshal() unMarshalPatternFlowGtpv2Spare2 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv2Spare2{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv2Spare2) ToProto() (*otg.PatternFlowGtpv2Spare2, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv2Spare2) FromProto(msg *otg.PatternFlowGtpv2Spare2) (PatternFlowGtpv2Spare2, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv2Spare2) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2Spare2) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv2Spare2) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2Spare2) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv2Spare2) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2Spare2) FromJson(value string) error {
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

func (obj *patternFlowGtpv2Spare2) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2Spare2) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2Spare2) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv2Spare2) Clone() (PatternFlowGtpv2Spare2, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv2Spare2()
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

func (obj *patternFlowGtpv2Spare2) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGtpv2Spare2 is reserved field
type PatternFlowGtpv2Spare2 interface {
	Validation
	// msg marshals PatternFlowGtpv2Spare2 to protobuf object *otg.PatternFlowGtpv2Spare2
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv2Spare2
	// setMsg unmarshals PatternFlowGtpv2Spare2 from protobuf object *otg.PatternFlowGtpv2Spare2
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv2Spare2) PatternFlowGtpv2Spare2
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv2Spare2
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv2Spare2
	// validate validates PatternFlowGtpv2Spare2
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv2Spare2, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGtpv2Spare2ChoiceEnum, set in PatternFlowGtpv2Spare2
	Choice() PatternFlowGtpv2Spare2ChoiceEnum
	// setChoice assigns PatternFlowGtpv2Spare2ChoiceEnum provided by user to PatternFlowGtpv2Spare2
	setChoice(value PatternFlowGtpv2Spare2ChoiceEnum) PatternFlowGtpv2Spare2
	// HasChoice checks if Choice has been set in PatternFlowGtpv2Spare2
	HasChoice() bool
	// Value returns uint32, set in PatternFlowGtpv2Spare2.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowGtpv2Spare2
	SetValue(value uint32) PatternFlowGtpv2Spare2
	// HasValue checks if Value has been set in PatternFlowGtpv2Spare2
	HasValue() bool
	// Values returns []uint32, set in PatternFlowGtpv2Spare2.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowGtpv2Spare2
	SetValues(value []uint32) PatternFlowGtpv2Spare2
	// Increment returns PatternFlowGtpv2Spare2Counter, set in PatternFlowGtpv2Spare2.
	// PatternFlowGtpv2Spare2Counter is integer counter pattern
	Increment() PatternFlowGtpv2Spare2Counter
	// SetIncrement assigns PatternFlowGtpv2Spare2Counter provided by user to PatternFlowGtpv2Spare2.
	// PatternFlowGtpv2Spare2Counter is integer counter pattern
	SetIncrement(value PatternFlowGtpv2Spare2Counter) PatternFlowGtpv2Spare2
	// HasIncrement checks if Increment has been set in PatternFlowGtpv2Spare2
	HasIncrement() bool
	// Decrement returns PatternFlowGtpv2Spare2Counter, set in PatternFlowGtpv2Spare2.
	// PatternFlowGtpv2Spare2Counter is integer counter pattern
	Decrement() PatternFlowGtpv2Spare2Counter
	// SetDecrement assigns PatternFlowGtpv2Spare2Counter provided by user to PatternFlowGtpv2Spare2.
	// PatternFlowGtpv2Spare2Counter is integer counter pattern
	SetDecrement(value PatternFlowGtpv2Spare2Counter) PatternFlowGtpv2Spare2
	// HasDecrement checks if Decrement has been set in PatternFlowGtpv2Spare2
	HasDecrement() bool
	// MetricTags returns PatternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIterIter, set in PatternFlowGtpv2Spare2
	MetricTags() PatternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter
	setNil()
}

type PatternFlowGtpv2Spare2ChoiceEnum string

// Enum of Choice on PatternFlowGtpv2Spare2
var PatternFlowGtpv2Spare2Choice = struct {
	VALUE     PatternFlowGtpv2Spare2ChoiceEnum
	VALUES    PatternFlowGtpv2Spare2ChoiceEnum
	INCREMENT PatternFlowGtpv2Spare2ChoiceEnum
	DECREMENT PatternFlowGtpv2Spare2ChoiceEnum
}{
	VALUE:     PatternFlowGtpv2Spare2ChoiceEnum("value"),
	VALUES:    PatternFlowGtpv2Spare2ChoiceEnum("values"),
	INCREMENT: PatternFlowGtpv2Spare2ChoiceEnum("increment"),
	DECREMENT: PatternFlowGtpv2Spare2ChoiceEnum("decrement"),
}

func (obj *patternFlowGtpv2Spare2) Choice() PatternFlowGtpv2Spare2ChoiceEnum {
	return PatternFlowGtpv2Spare2ChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGtpv2Spare2) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGtpv2Spare2) setChoice(value PatternFlowGtpv2Spare2ChoiceEnum) PatternFlowGtpv2Spare2 {
	intValue, ok := otg.PatternFlowGtpv2Spare2_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGtpv2Spare2ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGtpv2Spare2_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGtpv2Spare2Choice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGtpv2Spare2Choice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGtpv2Spare2Choice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGtpv2Spare2Counter().msg()
	}

	if value == PatternFlowGtpv2Spare2Choice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGtpv2Spare2Counter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv2Spare2) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGtpv2Spare2Choice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv2Spare2) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowGtpv2Spare2 object
func (obj *patternFlowGtpv2Spare2) SetValue(value uint32) PatternFlowGtpv2Spare2 {
	obj.setChoice(PatternFlowGtpv2Spare2Choice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowGtpv2Spare2) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowGtpv2Spare2 object
func (obj *patternFlowGtpv2Spare2) SetValues(value []uint32) PatternFlowGtpv2Spare2 {
	obj.setChoice(PatternFlowGtpv2Spare2Choice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGtpv2Spare2Counter
func (obj *patternFlowGtpv2Spare2) Increment() PatternFlowGtpv2Spare2Counter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGtpv2Spare2Choice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGtpv2Spare2Counter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGtpv2Spare2Counter
func (obj *patternFlowGtpv2Spare2) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGtpv2Spare2Counter value in the PatternFlowGtpv2Spare2 object
func (obj *patternFlowGtpv2Spare2) SetIncrement(value PatternFlowGtpv2Spare2Counter) PatternFlowGtpv2Spare2 {
	obj.setChoice(PatternFlowGtpv2Spare2Choice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGtpv2Spare2Counter
func (obj *patternFlowGtpv2Spare2) Decrement() PatternFlowGtpv2Spare2Counter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGtpv2Spare2Choice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGtpv2Spare2Counter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGtpv2Spare2Counter
func (obj *patternFlowGtpv2Spare2) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGtpv2Spare2Counter value in the PatternFlowGtpv2Spare2 object
func (obj *patternFlowGtpv2Spare2) SetDecrement(value PatternFlowGtpv2Spare2Counter) PatternFlowGtpv2Spare2 {
	obj.setChoice(PatternFlowGtpv2Spare2Choice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGtpv2Spare2MetricTag
func (obj *patternFlowGtpv2Spare2) MetricTags() PatternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGtpv2Spare2MetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter struct {
	obj                                  *patternFlowGtpv2Spare2
	patternFlowGtpv2Spare2MetricTagSlice []PatternFlowGtpv2Spare2MetricTag
	fieldPtr                             *[]*otg.PatternFlowGtpv2Spare2MetricTag
}

func newPatternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter(ptr *[]*otg.PatternFlowGtpv2Spare2MetricTag) PatternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter {
	return &patternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter{fieldPtr: ptr}
}

type PatternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter interface {
	setMsg(*patternFlowGtpv2Spare2) PatternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter
	Items() []PatternFlowGtpv2Spare2MetricTag
	Add() PatternFlowGtpv2Spare2MetricTag
	Append(items ...PatternFlowGtpv2Spare2MetricTag) PatternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter
	Set(index int, newObj PatternFlowGtpv2Spare2MetricTag) PatternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter
	Clear() PatternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter
	clearHolderSlice() PatternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter
	appendHolderSlice(item PatternFlowGtpv2Spare2MetricTag) PatternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter
}

func (obj *patternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter) setMsg(msg *patternFlowGtpv2Spare2) PatternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGtpv2Spare2MetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter) Items() []PatternFlowGtpv2Spare2MetricTag {
	return obj.patternFlowGtpv2Spare2MetricTagSlice
}

func (obj *patternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter) Add() PatternFlowGtpv2Spare2MetricTag {
	newObj := &otg.PatternFlowGtpv2Spare2MetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGtpv2Spare2MetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGtpv2Spare2MetricTagSlice = append(obj.patternFlowGtpv2Spare2MetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter) Append(items ...PatternFlowGtpv2Spare2MetricTag) PatternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGtpv2Spare2MetricTagSlice = append(obj.patternFlowGtpv2Spare2MetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter) Set(index int, newObj PatternFlowGtpv2Spare2MetricTag) PatternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGtpv2Spare2MetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter) Clear() PatternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGtpv2Spare2MetricTag{}
		obj.patternFlowGtpv2Spare2MetricTagSlice = []PatternFlowGtpv2Spare2MetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter) clearHolderSlice() PatternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter {
	if len(obj.patternFlowGtpv2Spare2MetricTagSlice) > 0 {
		obj.patternFlowGtpv2Spare2MetricTagSlice = []PatternFlowGtpv2Spare2MetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter) appendHolderSlice(item PatternFlowGtpv2Spare2MetricTag) PatternFlowGtpv2Spare2PatternFlowGtpv2Spare2MetricTagIter {
	obj.patternFlowGtpv2Spare2MetricTagSlice = append(obj.patternFlowGtpv2Spare2MetricTagSlice, item)
	return obj
}

func (obj *patternFlowGtpv2Spare2) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2Spare2.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowGtpv2Spare2.Values <= 255 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowGtpv2Spare2MetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGtpv2Spare2) setDefault() {
	var choices_set int = 0
	var choice PatternFlowGtpv2Spare2ChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowGtpv2Spare2Choice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowGtpv2Spare2Choice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowGtpv2Spare2Choice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowGtpv2Spare2Choice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowGtpv2Spare2Choice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowGtpv2Spare2")
			}
		} else {
			intVal := otg.PatternFlowGtpv2Spare2_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowGtpv2Spare2_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
