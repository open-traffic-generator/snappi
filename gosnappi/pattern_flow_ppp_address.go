package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPppAddress *****
type patternFlowPppAddress struct {
	validation
	obj              *otg.PatternFlowPppAddress
	marshaller       marshalPatternFlowPppAddress
	unMarshaller     unMarshalPatternFlowPppAddress
	incrementHolder  PatternFlowPppAddressCounter
	decrementHolder  PatternFlowPppAddressCounter
	metricTagsHolder PatternFlowPppAddressPatternFlowPppAddressMetricTagIter
}

func NewPatternFlowPppAddress() PatternFlowPppAddress {
	obj := patternFlowPppAddress{obj: &otg.PatternFlowPppAddress{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPppAddress) msg() *otg.PatternFlowPppAddress {
	return obj.obj
}

func (obj *patternFlowPppAddress) setMsg(msg *otg.PatternFlowPppAddress) PatternFlowPppAddress {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPppAddress struct {
	obj *patternFlowPppAddress
}

type marshalPatternFlowPppAddress interface {
	// ToProto marshals PatternFlowPppAddress to protobuf object *otg.PatternFlowPppAddress
	ToProto() (*otg.PatternFlowPppAddress, error)
	// ToPbText marshals PatternFlowPppAddress to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPppAddress to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPppAddress to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPppAddress struct {
	obj *patternFlowPppAddress
}

type unMarshalPatternFlowPppAddress interface {
	// FromProto unmarshals PatternFlowPppAddress from protobuf object *otg.PatternFlowPppAddress
	FromProto(msg *otg.PatternFlowPppAddress) (PatternFlowPppAddress, error)
	// FromPbText unmarshals PatternFlowPppAddress from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPppAddress from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPppAddress from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPppAddress) Marshal() marshalPatternFlowPppAddress {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPppAddress{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPppAddress) Unmarshal() unMarshalPatternFlowPppAddress {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPppAddress{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPppAddress) ToProto() (*otg.PatternFlowPppAddress, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPppAddress) FromProto(msg *otg.PatternFlowPppAddress) (PatternFlowPppAddress, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPppAddress) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPppAddress) FromPbText(value string) error {
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

func (m *marshalpatternFlowPppAddress) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPppAddress) FromYaml(value string) error {
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

func (m *marshalpatternFlowPppAddress) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPppAddress) FromJson(value string) error {
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

func (obj *patternFlowPppAddress) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPppAddress) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPppAddress) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPppAddress) Clone() (PatternFlowPppAddress, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPppAddress()
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

func (obj *patternFlowPppAddress) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowPppAddress is pPP address
type PatternFlowPppAddress interface {
	Validation
	// msg marshals PatternFlowPppAddress to protobuf object *otg.PatternFlowPppAddress
	// and doesn't set defaults
	msg() *otg.PatternFlowPppAddress
	// setMsg unmarshals PatternFlowPppAddress from protobuf object *otg.PatternFlowPppAddress
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPppAddress) PatternFlowPppAddress
	// provides marshal interface
	Marshal() marshalPatternFlowPppAddress
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPppAddress
	// validate validates PatternFlowPppAddress
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPppAddress, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowPppAddressChoiceEnum, set in PatternFlowPppAddress
	Choice() PatternFlowPppAddressChoiceEnum
	// setChoice assigns PatternFlowPppAddressChoiceEnum provided by user to PatternFlowPppAddress
	setChoice(value PatternFlowPppAddressChoiceEnum) PatternFlowPppAddress
	// HasChoice checks if Choice has been set in PatternFlowPppAddress
	HasChoice() bool
	// Value returns uint32, set in PatternFlowPppAddress.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowPppAddress
	SetValue(value uint32) PatternFlowPppAddress
	// HasValue checks if Value has been set in PatternFlowPppAddress
	HasValue() bool
	// Values returns []uint32, set in PatternFlowPppAddress.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowPppAddress
	SetValues(value []uint32) PatternFlowPppAddress
	// Increment returns PatternFlowPppAddressCounter, set in PatternFlowPppAddress.
	// PatternFlowPppAddressCounter is integer counter pattern
	Increment() PatternFlowPppAddressCounter
	// SetIncrement assigns PatternFlowPppAddressCounter provided by user to PatternFlowPppAddress.
	// PatternFlowPppAddressCounter is integer counter pattern
	SetIncrement(value PatternFlowPppAddressCounter) PatternFlowPppAddress
	// HasIncrement checks if Increment has been set in PatternFlowPppAddress
	HasIncrement() bool
	// Decrement returns PatternFlowPppAddressCounter, set in PatternFlowPppAddress.
	// PatternFlowPppAddressCounter is integer counter pattern
	Decrement() PatternFlowPppAddressCounter
	// SetDecrement assigns PatternFlowPppAddressCounter provided by user to PatternFlowPppAddress.
	// PatternFlowPppAddressCounter is integer counter pattern
	SetDecrement(value PatternFlowPppAddressCounter) PatternFlowPppAddress
	// HasDecrement checks if Decrement has been set in PatternFlowPppAddress
	HasDecrement() bool
	// MetricTags returns PatternFlowPppAddressPatternFlowPppAddressMetricTagIterIter, set in PatternFlowPppAddress
	MetricTags() PatternFlowPppAddressPatternFlowPppAddressMetricTagIter
	setNil()
}

type PatternFlowPppAddressChoiceEnum string

// Enum of Choice on PatternFlowPppAddress
var PatternFlowPppAddressChoice = struct {
	VALUE     PatternFlowPppAddressChoiceEnum
	VALUES    PatternFlowPppAddressChoiceEnum
	INCREMENT PatternFlowPppAddressChoiceEnum
	DECREMENT PatternFlowPppAddressChoiceEnum
}{
	VALUE:     PatternFlowPppAddressChoiceEnum("value"),
	VALUES:    PatternFlowPppAddressChoiceEnum("values"),
	INCREMENT: PatternFlowPppAddressChoiceEnum("increment"),
	DECREMENT: PatternFlowPppAddressChoiceEnum("decrement"),
}

func (obj *patternFlowPppAddress) Choice() PatternFlowPppAddressChoiceEnum {
	return PatternFlowPppAddressChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowPppAddress) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowPppAddress) setChoice(value PatternFlowPppAddressChoiceEnum) PatternFlowPppAddress {
	intValue, ok := otg.PatternFlowPppAddress_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowPppAddressChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowPppAddress_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowPppAddressChoice.VALUE {
		defaultValue := uint32(255)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowPppAddressChoice.VALUES {
		defaultValue := []uint32{255}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowPppAddressChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowPppAddressCounter().msg()
	}

	if value == PatternFlowPppAddressChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowPppAddressCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowPppAddress) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowPppAddressChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowPppAddress) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowPppAddress object
func (obj *patternFlowPppAddress) SetValue(value uint32) PatternFlowPppAddress {
	obj.setChoice(PatternFlowPppAddressChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowPppAddress) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{255})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowPppAddress object
func (obj *patternFlowPppAddress) SetValues(value []uint32) PatternFlowPppAddress {
	obj.setChoice(PatternFlowPppAddressChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowPppAddressCounter
func (obj *patternFlowPppAddress) Increment() PatternFlowPppAddressCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowPppAddressChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowPppAddressCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowPppAddressCounter
func (obj *patternFlowPppAddress) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowPppAddressCounter value in the PatternFlowPppAddress object
func (obj *patternFlowPppAddress) SetIncrement(value PatternFlowPppAddressCounter) PatternFlowPppAddress {
	obj.setChoice(PatternFlowPppAddressChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowPppAddressCounter
func (obj *patternFlowPppAddress) Decrement() PatternFlowPppAddressCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowPppAddressChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowPppAddressCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowPppAddressCounter
func (obj *patternFlowPppAddress) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowPppAddressCounter value in the PatternFlowPppAddress object
func (obj *patternFlowPppAddress) SetDecrement(value PatternFlowPppAddressCounter) PatternFlowPppAddress {
	obj.setChoice(PatternFlowPppAddressChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowPppAddressMetricTag
func (obj *patternFlowPppAddress) MetricTags() PatternFlowPppAddressPatternFlowPppAddressMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowPppAddressMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowPppAddressPatternFlowPppAddressMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowPppAddressPatternFlowPppAddressMetricTagIter struct {
	obj                                 *patternFlowPppAddress
	patternFlowPppAddressMetricTagSlice []PatternFlowPppAddressMetricTag
	fieldPtr                            *[]*otg.PatternFlowPppAddressMetricTag
}

func newPatternFlowPppAddressPatternFlowPppAddressMetricTagIter(ptr *[]*otg.PatternFlowPppAddressMetricTag) PatternFlowPppAddressPatternFlowPppAddressMetricTagIter {
	return &patternFlowPppAddressPatternFlowPppAddressMetricTagIter{fieldPtr: ptr}
}

type PatternFlowPppAddressPatternFlowPppAddressMetricTagIter interface {
	setMsg(*patternFlowPppAddress) PatternFlowPppAddressPatternFlowPppAddressMetricTagIter
	Items() []PatternFlowPppAddressMetricTag
	Add() PatternFlowPppAddressMetricTag
	Append(items ...PatternFlowPppAddressMetricTag) PatternFlowPppAddressPatternFlowPppAddressMetricTagIter
	Set(index int, newObj PatternFlowPppAddressMetricTag) PatternFlowPppAddressPatternFlowPppAddressMetricTagIter
	Clear() PatternFlowPppAddressPatternFlowPppAddressMetricTagIter
	clearHolderSlice() PatternFlowPppAddressPatternFlowPppAddressMetricTagIter
	appendHolderSlice(item PatternFlowPppAddressMetricTag) PatternFlowPppAddressPatternFlowPppAddressMetricTagIter
}

func (obj *patternFlowPppAddressPatternFlowPppAddressMetricTagIter) setMsg(msg *patternFlowPppAddress) PatternFlowPppAddressPatternFlowPppAddressMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowPppAddressMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowPppAddressPatternFlowPppAddressMetricTagIter) Items() []PatternFlowPppAddressMetricTag {
	return obj.patternFlowPppAddressMetricTagSlice
}

func (obj *patternFlowPppAddressPatternFlowPppAddressMetricTagIter) Add() PatternFlowPppAddressMetricTag {
	newObj := &otg.PatternFlowPppAddressMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowPppAddressMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowPppAddressMetricTagSlice = append(obj.patternFlowPppAddressMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowPppAddressPatternFlowPppAddressMetricTagIter) Append(items ...PatternFlowPppAddressMetricTag) PatternFlowPppAddressPatternFlowPppAddressMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowPppAddressMetricTagSlice = append(obj.patternFlowPppAddressMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowPppAddressPatternFlowPppAddressMetricTagIter) Set(index int, newObj PatternFlowPppAddressMetricTag) PatternFlowPppAddressPatternFlowPppAddressMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowPppAddressMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowPppAddressPatternFlowPppAddressMetricTagIter) Clear() PatternFlowPppAddressPatternFlowPppAddressMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowPppAddressMetricTag{}
		obj.patternFlowPppAddressMetricTagSlice = []PatternFlowPppAddressMetricTag{}
	}
	return obj
}
func (obj *patternFlowPppAddressPatternFlowPppAddressMetricTagIter) clearHolderSlice() PatternFlowPppAddressPatternFlowPppAddressMetricTagIter {
	if len(obj.patternFlowPppAddressMetricTagSlice) > 0 {
		obj.patternFlowPppAddressMetricTagSlice = []PatternFlowPppAddressMetricTag{}
	}
	return obj
}
func (obj *patternFlowPppAddressPatternFlowPppAddressMetricTagIter) appendHolderSlice(item PatternFlowPppAddressMetricTag) PatternFlowPppAddressPatternFlowPppAddressMetricTagIter {
	obj.patternFlowPppAddressMetricTagSlice = append(obj.patternFlowPppAddressMetricTagSlice, item)
	return obj
}

func (obj *patternFlowPppAddress) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPppAddress.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowPppAddress.Values <= 255 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowPppAddressMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowPppAddress) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowPppAddressChoice.VALUE)

	}

}
