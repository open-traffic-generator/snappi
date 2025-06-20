package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv1GroupAddress *****
type patternFlowIgmpv1GroupAddress struct {
	validation
	obj              *otg.PatternFlowIgmpv1GroupAddress
	marshaller       marshalPatternFlowIgmpv1GroupAddress
	unMarshaller     unMarshalPatternFlowIgmpv1GroupAddress
	incrementHolder  PatternFlowIgmpv1GroupAddressCounter
	decrementHolder  PatternFlowIgmpv1GroupAddressCounter
	metricTagsHolder PatternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter
}

func NewPatternFlowIgmpv1GroupAddress() PatternFlowIgmpv1GroupAddress {
	obj := patternFlowIgmpv1GroupAddress{obj: &otg.PatternFlowIgmpv1GroupAddress{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv1GroupAddress) msg() *otg.PatternFlowIgmpv1GroupAddress {
	return obj.obj
}

func (obj *patternFlowIgmpv1GroupAddress) setMsg(msg *otg.PatternFlowIgmpv1GroupAddress) PatternFlowIgmpv1GroupAddress {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv1GroupAddress struct {
	obj *patternFlowIgmpv1GroupAddress
}

type marshalPatternFlowIgmpv1GroupAddress interface {
	// ToProto marshals PatternFlowIgmpv1GroupAddress to protobuf object *otg.PatternFlowIgmpv1GroupAddress
	ToProto() (*otg.PatternFlowIgmpv1GroupAddress, error)
	// ToPbText marshals PatternFlowIgmpv1GroupAddress to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv1GroupAddress to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv1GroupAddress to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv1GroupAddress struct {
	obj *patternFlowIgmpv1GroupAddress
}

type unMarshalPatternFlowIgmpv1GroupAddress interface {
	// FromProto unmarshals PatternFlowIgmpv1GroupAddress from protobuf object *otg.PatternFlowIgmpv1GroupAddress
	FromProto(msg *otg.PatternFlowIgmpv1GroupAddress) (PatternFlowIgmpv1GroupAddress, error)
	// FromPbText unmarshals PatternFlowIgmpv1GroupAddress from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv1GroupAddress from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv1GroupAddress from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv1GroupAddress) Marshal() marshalPatternFlowIgmpv1GroupAddress {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv1GroupAddress{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv1GroupAddress) Unmarshal() unMarshalPatternFlowIgmpv1GroupAddress {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv1GroupAddress{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv1GroupAddress) ToProto() (*otg.PatternFlowIgmpv1GroupAddress, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv1GroupAddress) FromProto(msg *otg.PatternFlowIgmpv1GroupAddress) (PatternFlowIgmpv1GroupAddress, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv1GroupAddress) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1GroupAddress) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv1GroupAddress) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1GroupAddress) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv1GroupAddress) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1GroupAddress) FromJson(value string) error {
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

func (obj *patternFlowIgmpv1GroupAddress) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv1GroupAddress) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv1GroupAddress) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv1GroupAddress) Clone() (PatternFlowIgmpv1GroupAddress, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv1GroupAddress()
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

func (obj *patternFlowIgmpv1GroupAddress) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIgmpv1GroupAddress is group address
type PatternFlowIgmpv1GroupAddress interface {
	Validation
	// msg marshals PatternFlowIgmpv1GroupAddress to protobuf object *otg.PatternFlowIgmpv1GroupAddress
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv1GroupAddress
	// setMsg unmarshals PatternFlowIgmpv1GroupAddress from protobuf object *otg.PatternFlowIgmpv1GroupAddress
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv1GroupAddress) PatternFlowIgmpv1GroupAddress
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv1GroupAddress
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv1GroupAddress
	// validate validates PatternFlowIgmpv1GroupAddress
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv1GroupAddress, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIgmpv1GroupAddressChoiceEnum, set in PatternFlowIgmpv1GroupAddress
	Choice() PatternFlowIgmpv1GroupAddressChoiceEnum
	// setChoice assigns PatternFlowIgmpv1GroupAddressChoiceEnum provided by user to PatternFlowIgmpv1GroupAddress
	setChoice(value PatternFlowIgmpv1GroupAddressChoiceEnum) PatternFlowIgmpv1GroupAddress
	// HasChoice checks if Choice has been set in PatternFlowIgmpv1GroupAddress
	HasChoice() bool
	// Value returns string, set in PatternFlowIgmpv1GroupAddress.
	Value() string
	// SetValue assigns string provided by user to PatternFlowIgmpv1GroupAddress
	SetValue(value string) PatternFlowIgmpv1GroupAddress
	// HasValue checks if Value has been set in PatternFlowIgmpv1GroupAddress
	HasValue() bool
	// Values returns []string, set in PatternFlowIgmpv1GroupAddress.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowIgmpv1GroupAddress
	SetValues(value []string) PatternFlowIgmpv1GroupAddress
	// Increment returns PatternFlowIgmpv1GroupAddressCounter, set in PatternFlowIgmpv1GroupAddress.
	// PatternFlowIgmpv1GroupAddressCounter is ipv4 counter pattern
	Increment() PatternFlowIgmpv1GroupAddressCounter
	// SetIncrement assigns PatternFlowIgmpv1GroupAddressCounter provided by user to PatternFlowIgmpv1GroupAddress.
	// PatternFlowIgmpv1GroupAddressCounter is ipv4 counter pattern
	SetIncrement(value PatternFlowIgmpv1GroupAddressCounter) PatternFlowIgmpv1GroupAddress
	// HasIncrement checks if Increment has been set in PatternFlowIgmpv1GroupAddress
	HasIncrement() bool
	// Decrement returns PatternFlowIgmpv1GroupAddressCounter, set in PatternFlowIgmpv1GroupAddress.
	// PatternFlowIgmpv1GroupAddressCounter is ipv4 counter pattern
	Decrement() PatternFlowIgmpv1GroupAddressCounter
	// SetDecrement assigns PatternFlowIgmpv1GroupAddressCounter provided by user to PatternFlowIgmpv1GroupAddress.
	// PatternFlowIgmpv1GroupAddressCounter is ipv4 counter pattern
	SetDecrement(value PatternFlowIgmpv1GroupAddressCounter) PatternFlowIgmpv1GroupAddress
	// HasDecrement checks if Decrement has been set in PatternFlowIgmpv1GroupAddress
	HasDecrement() bool
	// MetricTags returns PatternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIterIter, set in PatternFlowIgmpv1GroupAddress
	MetricTags() PatternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter
	setNil()
}

type PatternFlowIgmpv1GroupAddressChoiceEnum string

// Enum of Choice on PatternFlowIgmpv1GroupAddress
var PatternFlowIgmpv1GroupAddressChoice = struct {
	VALUE     PatternFlowIgmpv1GroupAddressChoiceEnum
	VALUES    PatternFlowIgmpv1GroupAddressChoiceEnum
	INCREMENT PatternFlowIgmpv1GroupAddressChoiceEnum
	DECREMENT PatternFlowIgmpv1GroupAddressChoiceEnum
}{
	VALUE:     PatternFlowIgmpv1GroupAddressChoiceEnum("value"),
	VALUES:    PatternFlowIgmpv1GroupAddressChoiceEnum("values"),
	INCREMENT: PatternFlowIgmpv1GroupAddressChoiceEnum("increment"),
	DECREMENT: PatternFlowIgmpv1GroupAddressChoiceEnum("decrement"),
}

func (obj *patternFlowIgmpv1GroupAddress) Choice() PatternFlowIgmpv1GroupAddressChoiceEnum {
	return PatternFlowIgmpv1GroupAddressChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIgmpv1GroupAddress) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIgmpv1GroupAddress) setChoice(value PatternFlowIgmpv1GroupAddressChoiceEnum) PatternFlowIgmpv1GroupAddress {
	intValue, ok := otg.PatternFlowIgmpv1GroupAddress_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv1GroupAddressChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv1GroupAddress_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIgmpv1GroupAddressChoice.VALUE {
		defaultValue := "0.0.0.0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIgmpv1GroupAddressChoice.VALUES {
		defaultValue := []string{"0.0.0.0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIgmpv1GroupAddressChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIgmpv1GroupAddressCounter().msg()
	}

	if value == PatternFlowIgmpv1GroupAddressChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIgmpv1GroupAddressCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowIgmpv1GroupAddress) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIgmpv1GroupAddressChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowIgmpv1GroupAddress) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowIgmpv1GroupAddress object
func (obj *patternFlowIgmpv1GroupAddress) SetValue(value string) PatternFlowIgmpv1GroupAddress {
	obj.setChoice(PatternFlowIgmpv1GroupAddressChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowIgmpv1GroupAddress) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"0.0.0.0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowIgmpv1GroupAddress object
func (obj *patternFlowIgmpv1GroupAddress) SetValues(value []string) PatternFlowIgmpv1GroupAddress {
	obj.setChoice(PatternFlowIgmpv1GroupAddressChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIgmpv1GroupAddressCounter
func (obj *patternFlowIgmpv1GroupAddress) Increment() PatternFlowIgmpv1GroupAddressCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIgmpv1GroupAddressChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIgmpv1GroupAddressCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIgmpv1GroupAddressCounter
func (obj *patternFlowIgmpv1GroupAddress) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIgmpv1GroupAddressCounter value in the PatternFlowIgmpv1GroupAddress object
func (obj *patternFlowIgmpv1GroupAddress) SetIncrement(value PatternFlowIgmpv1GroupAddressCounter) PatternFlowIgmpv1GroupAddress {
	obj.setChoice(PatternFlowIgmpv1GroupAddressChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIgmpv1GroupAddressCounter
func (obj *patternFlowIgmpv1GroupAddress) Decrement() PatternFlowIgmpv1GroupAddressCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIgmpv1GroupAddressChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIgmpv1GroupAddressCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIgmpv1GroupAddressCounter
func (obj *patternFlowIgmpv1GroupAddress) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIgmpv1GroupAddressCounter value in the PatternFlowIgmpv1GroupAddress object
func (obj *patternFlowIgmpv1GroupAddress) SetDecrement(value PatternFlowIgmpv1GroupAddressCounter) PatternFlowIgmpv1GroupAddress {
	obj.setChoice(PatternFlowIgmpv1GroupAddressChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIgmpv1GroupAddressMetricTag
func (obj *patternFlowIgmpv1GroupAddress) MetricTags() PatternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIgmpv1GroupAddressMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter struct {
	obj                                         *patternFlowIgmpv1GroupAddress
	patternFlowIgmpv1GroupAddressMetricTagSlice []PatternFlowIgmpv1GroupAddressMetricTag
	fieldPtr                                    *[]*otg.PatternFlowIgmpv1GroupAddressMetricTag
}

func newPatternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter(ptr *[]*otg.PatternFlowIgmpv1GroupAddressMetricTag) PatternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter {
	return &patternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter interface {
	setMsg(*patternFlowIgmpv1GroupAddress) PatternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter
	Items() []PatternFlowIgmpv1GroupAddressMetricTag
	Add() PatternFlowIgmpv1GroupAddressMetricTag
	Append(items ...PatternFlowIgmpv1GroupAddressMetricTag) PatternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter
	Set(index int, newObj PatternFlowIgmpv1GroupAddressMetricTag) PatternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter
	Clear() PatternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter
	clearHolderSlice() PatternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter
	appendHolderSlice(item PatternFlowIgmpv1GroupAddressMetricTag) PatternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter
}

func (obj *patternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter) setMsg(msg *patternFlowIgmpv1GroupAddress) PatternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIgmpv1GroupAddressMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter) Items() []PatternFlowIgmpv1GroupAddressMetricTag {
	return obj.patternFlowIgmpv1GroupAddressMetricTagSlice
}

func (obj *patternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter) Add() PatternFlowIgmpv1GroupAddressMetricTag {
	newObj := &otg.PatternFlowIgmpv1GroupAddressMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIgmpv1GroupAddressMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIgmpv1GroupAddressMetricTagSlice = append(obj.patternFlowIgmpv1GroupAddressMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter) Append(items ...PatternFlowIgmpv1GroupAddressMetricTag) PatternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIgmpv1GroupAddressMetricTagSlice = append(obj.patternFlowIgmpv1GroupAddressMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter) Set(index int, newObj PatternFlowIgmpv1GroupAddressMetricTag) PatternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIgmpv1GroupAddressMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter) Clear() PatternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIgmpv1GroupAddressMetricTag{}
		obj.patternFlowIgmpv1GroupAddressMetricTagSlice = []PatternFlowIgmpv1GroupAddressMetricTag{}
	}
	return obj
}
func (obj *patternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter) clearHolderSlice() PatternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter {
	if len(obj.patternFlowIgmpv1GroupAddressMetricTagSlice) > 0 {
		obj.patternFlowIgmpv1GroupAddressMetricTagSlice = []PatternFlowIgmpv1GroupAddressMetricTag{}
	}
	return obj
}
func (obj *patternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter) appendHolderSlice(item PatternFlowIgmpv1GroupAddressMetricTag) PatternFlowIgmpv1GroupAddressPatternFlowIgmpv1GroupAddressMetricTagIter {
	obj.patternFlowIgmpv1GroupAddressMetricTagSlice = append(obj.patternFlowIgmpv1GroupAddressMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIgmpv1GroupAddress) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIgmpv1GroupAddress.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIgmpv1GroupAddress.Values"))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIgmpv1GroupAddressMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIgmpv1GroupAddress) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIgmpv1GroupAddressChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIgmpv1GroupAddressChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIgmpv1GroupAddressChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIgmpv1GroupAddressChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIgmpv1GroupAddressChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIgmpv1GroupAddressChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIgmpv1GroupAddress")
			}
		} else {
			intVal := otg.PatternFlowIgmpv1GroupAddress_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIgmpv1GroupAddress_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
