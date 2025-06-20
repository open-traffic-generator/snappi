package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4Identification *****
type patternFlowIpv4Identification struct {
	validation
	obj              *otg.PatternFlowIpv4Identification
	marshaller       marshalPatternFlowIpv4Identification
	unMarshaller     unMarshalPatternFlowIpv4Identification
	incrementHolder  PatternFlowIpv4IdentificationCounter
	decrementHolder  PatternFlowIpv4IdentificationCounter
	metricTagsHolder PatternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter
}

func NewPatternFlowIpv4Identification() PatternFlowIpv4Identification {
	obj := patternFlowIpv4Identification{obj: &otg.PatternFlowIpv4Identification{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4Identification) msg() *otg.PatternFlowIpv4Identification {
	return obj.obj
}

func (obj *patternFlowIpv4Identification) setMsg(msg *otg.PatternFlowIpv4Identification) PatternFlowIpv4Identification {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4Identification struct {
	obj *patternFlowIpv4Identification
}

type marshalPatternFlowIpv4Identification interface {
	// ToProto marshals PatternFlowIpv4Identification to protobuf object *otg.PatternFlowIpv4Identification
	ToProto() (*otg.PatternFlowIpv4Identification, error)
	// ToPbText marshals PatternFlowIpv4Identification to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4Identification to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4Identification to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4Identification struct {
	obj *patternFlowIpv4Identification
}

type unMarshalPatternFlowIpv4Identification interface {
	// FromProto unmarshals PatternFlowIpv4Identification from protobuf object *otg.PatternFlowIpv4Identification
	FromProto(msg *otg.PatternFlowIpv4Identification) (PatternFlowIpv4Identification, error)
	// FromPbText unmarshals PatternFlowIpv4Identification from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4Identification from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4Identification from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4Identification) Marshal() marshalPatternFlowIpv4Identification {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4Identification{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4Identification) Unmarshal() unMarshalPatternFlowIpv4Identification {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4Identification{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4Identification) ToProto() (*otg.PatternFlowIpv4Identification, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4Identification) FromProto(msg *otg.PatternFlowIpv4Identification) (PatternFlowIpv4Identification, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4Identification) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4Identification) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4Identification) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4Identification) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4Identification) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4Identification) FromJson(value string) error {
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

func (obj *patternFlowIpv4Identification) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4Identification) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4Identification) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4Identification) Clone() (PatternFlowIpv4Identification, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4Identification()
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

func (obj *patternFlowIpv4Identification) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4Identification is identification
type PatternFlowIpv4Identification interface {
	Validation
	// msg marshals PatternFlowIpv4Identification to protobuf object *otg.PatternFlowIpv4Identification
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4Identification
	// setMsg unmarshals PatternFlowIpv4Identification from protobuf object *otg.PatternFlowIpv4Identification
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4Identification) PatternFlowIpv4Identification
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4Identification
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4Identification
	// validate validates PatternFlowIpv4Identification
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4Identification, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4IdentificationChoiceEnum, set in PatternFlowIpv4Identification
	Choice() PatternFlowIpv4IdentificationChoiceEnum
	// setChoice assigns PatternFlowIpv4IdentificationChoiceEnum provided by user to PatternFlowIpv4Identification
	setChoice(value PatternFlowIpv4IdentificationChoiceEnum) PatternFlowIpv4Identification
	// HasChoice checks if Choice has been set in PatternFlowIpv4Identification
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv4Identification.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv4Identification
	SetValue(value uint32) PatternFlowIpv4Identification
	// HasValue checks if Value has been set in PatternFlowIpv4Identification
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv4Identification.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv4Identification
	SetValues(value []uint32) PatternFlowIpv4Identification
	// Increment returns PatternFlowIpv4IdentificationCounter, set in PatternFlowIpv4Identification.
	// PatternFlowIpv4IdentificationCounter is integer counter pattern
	Increment() PatternFlowIpv4IdentificationCounter
	// SetIncrement assigns PatternFlowIpv4IdentificationCounter provided by user to PatternFlowIpv4Identification.
	// PatternFlowIpv4IdentificationCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv4IdentificationCounter) PatternFlowIpv4Identification
	// HasIncrement checks if Increment has been set in PatternFlowIpv4Identification
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4IdentificationCounter, set in PatternFlowIpv4Identification.
	// PatternFlowIpv4IdentificationCounter is integer counter pattern
	Decrement() PatternFlowIpv4IdentificationCounter
	// SetDecrement assigns PatternFlowIpv4IdentificationCounter provided by user to PatternFlowIpv4Identification.
	// PatternFlowIpv4IdentificationCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv4IdentificationCounter) PatternFlowIpv4Identification
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4Identification
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIterIter, set in PatternFlowIpv4Identification
	MetricTags() PatternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter
	setNil()
}

type PatternFlowIpv4IdentificationChoiceEnum string

// Enum of Choice on PatternFlowIpv4Identification
var PatternFlowIpv4IdentificationChoice = struct {
	VALUE     PatternFlowIpv4IdentificationChoiceEnum
	VALUES    PatternFlowIpv4IdentificationChoiceEnum
	INCREMENT PatternFlowIpv4IdentificationChoiceEnum
	DECREMENT PatternFlowIpv4IdentificationChoiceEnum
}{
	VALUE:     PatternFlowIpv4IdentificationChoiceEnum("value"),
	VALUES:    PatternFlowIpv4IdentificationChoiceEnum("values"),
	INCREMENT: PatternFlowIpv4IdentificationChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4IdentificationChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4Identification) Choice() PatternFlowIpv4IdentificationChoiceEnum {
	return PatternFlowIpv4IdentificationChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4Identification) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4Identification) setChoice(value PatternFlowIpv4IdentificationChoiceEnum) PatternFlowIpv4Identification {
	intValue, ok := otg.PatternFlowIpv4Identification_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4IdentificationChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4Identification_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4IdentificationChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4IdentificationChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4IdentificationChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4IdentificationCounter().msg()
	}

	if value == PatternFlowIpv4IdentificationChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4IdentificationCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4Identification) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4IdentificationChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4Identification) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv4Identification object
func (obj *patternFlowIpv4Identification) SetValue(value uint32) PatternFlowIpv4Identification {
	obj.setChoice(PatternFlowIpv4IdentificationChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv4Identification) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv4Identification object
func (obj *patternFlowIpv4Identification) SetValues(value []uint32) PatternFlowIpv4Identification {
	obj.setChoice(PatternFlowIpv4IdentificationChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv4IdentificationCounter
func (obj *patternFlowIpv4Identification) Increment() PatternFlowIpv4IdentificationCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4IdentificationChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4IdentificationCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4IdentificationCounter
func (obj *patternFlowIpv4Identification) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4IdentificationCounter value in the PatternFlowIpv4Identification object
func (obj *patternFlowIpv4Identification) SetIncrement(value PatternFlowIpv4IdentificationCounter) PatternFlowIpv4Identification {
	obj.setChoice(PatternFlowIpv4IdentificationChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4IdentificationCounter
func (obj *patternFlowIpv4Identification) Decrement() PatternFlowIpv4IdentificationCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4IdentificationChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4IdentificationCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4IdentificationCounter
func (obj *patternFlowIpv4Identification) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4IdentificationCounter value in the PatternFlowIpv4Identification object
func (obj *patternFlowIpv4Identification) SetDecrement(value PatternFlowIpv4IdentificationCounter) PatternFlowIpv4Identification {
	obj.setChoice(PatternFlowIpv4IdentificationChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv4IdentificationMetricTag
func (obj *patternFlowIpv4Identification) MetricTags() PatternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv4IdentificationMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter struct {
	obj                                         *patternFlowIpv4Identification
	patternFlowIpv4IdentificationMetricTagSlice []PatternFlowIpv4IdentificationMetricTag
	fieldPtr                                    *[]*otg.PatternFlowIpv4IdentificationMetricTag
}

func newPatternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter(ptr *[]*otg.PatternFlowIpv4IdentificationMetricTag) PatternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter {
	return &patternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter interface {
	setMsg(*patternFlowIpv4Identification) PatternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter
	Items() []PatternFlowIpv4IdentificationMetricTag
	Add() PatternFlowIpv4IdentificationMetricTag
	Append(items ...PatternFlowIpv4IdentificationMetricTag) PatternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter
	Set(index int, newObj PatternFlowIpv4IdentificationMetricTag) PatternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter
	Clear() PatternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter
	clearHolderSlice() PatternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter
	appendHolderSlice(item PatternFlowIpv4IdentificationMetricTag) PatternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter
}

func (obj *patternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter) setMsg(msg *patternFlowIpv4Identification) PatternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv4IdentificationMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter) Items() []PatternFlowIpv4IdentificationMetricTag {
	return obj.patternFlowIpv4IdentificationMetricTagSlice
}

func (obj *patternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter) Add() PatternFlowIpv4IdentificationMetricTag {
	newObj := &otg.PatternFlowIpv4IdentificationMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv4IdentificationMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv4IdentificationMetricTagSlice = append(obj.patternFlowIpv4IdentificationMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter) Append(items ...PatternFlowIpv4IdentificationMetricTag) PatternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv4IdentificationMetricTagSlice = append(obj.patternFlowIpv4IdentificationMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter) Set(index int, newObj PatternFlowIpv4IdentificationMetricTag) PatternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv4IdentificationMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter) Clear() PatternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv4IdentificationMetricTag{}
		obj.patternFlowIpv4IdentificationMetricTagSlice = []PatternFlowIpv4IdentificationMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter) clearHolderSlice() PatternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter {
	if len(obj.patternFlowIpv4IdentificationMetricTagSlice) > 0 {
		obj.patternFlowIpv4IdentificationMetricTagSlice = []PatternFlowIpv4IdentificationMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter) appendHolderSlice(item PatternFlowIpv4IdentificationMetricTag) PatternFlowIpv4IdentificationPatternFlowIpv4IdentificationMetricTagIter {
	obj.patternFlowIpv4IdentificationMetricTagSlice = append(obj.patternFlowIpv4IdentificationMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIpv4Identification) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4Identification.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv4Identification.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv4IdentificationMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIpv4Identification) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv4IdentificationChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv4IdentificationChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv4IdentificationChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv4IdentificationChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv4IdentificationChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv4IdentificationChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv4Identification")
			}
		} else {
			intVal := otg.PatternFlowIpv4Identification_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv4Identification_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
