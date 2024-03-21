package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowArpSenderProtocolAddr *****
type patternFlowArpSenderProtocolAddr struct {
	validation
	obj              *otg.PatternFlowArpSenderProtocolAddr
	marshaller       marshalPatternFlowArpSenderProtocolAddr
	unMarshaller     unMarshalPatternFlowArpSenderProtocolAddr
	incrementHolder  PatternFlowArpSenderProtocolAddrCounter
	decrementHolder  PatternFlowArpSenderProtocolAddrCounter
	metricTagsHolder PatternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter
}

func NewPatternFlowArpSenderProtocolAddr() PatternFlowArpSenderProtocolAddr {
	obj := patternFlowArpSenderProtocolAddr{obj: &otg.PatternFlowArpSenderProtocolAddr{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowArpSenderProtocolAddr) msg() *otg.PatternFlowArpSenderProtocolAddr {
	return obj.obj
}

func (obj *patternFlowArpSenderProtocolAddr) setMsg(msg *otg.PatternFlowArpSenderProtocolAddr) PatternFlowArpSenderProtocolAddr {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowArpSenderProtocolAddr struct {
	obj *patternFlowArpSenderProtocolAddr
}

type marshalPatternFlowArpSenderProtocolAddr interface {
	// ToProto marshals PatternFlowArpSenderProtocolAddr to protobuf object *otg.PatternFlowArpSenderProtocolAddr
	ToProto() (*otg.PatternFlowArpSenderProtocolAddr, error)
	// ToPbText marshals PatternFlowArpSenderProtocolAddr to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowArpSenderProtocolAddr to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowArpSenderProtocolAddr to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowArpSenderProtocolAddr struct {
	obj *patternFlowArpSenderProtocolAddr
}

type unMarshalPatternFlowArpSenderProtocolAddr interface {
	// FromProto unmarshals PatternFlowArpSenderProtocolAddr from protobuf object *otg.PatternFlowArpSenderProtocolAddr
	FromProto(msg *otg.PatternFlowArpSenderProtocolAddr) (PatternFlowArpSenderProtocolAddr, error)
	// FromPbText unmarshals PatternFlowArpSenderProtocolAddr from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowArpSenderProtocolAddr from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowArpSenderProtocolAddr from JSON text
	FromJson(value string) error
}

func (obj *patternFlowArpSenderProtocolAddr) Marshal() marshalPatternFlowArpSenderProtocolAddr {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowArpSenderProtocolAddr{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowArpSenderProtocolAddr) Unmarshal() unMarshalPatternFlowArpSenderProtocolAddr {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowArpSenderProtocolAddr{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowArpSenderProtocolAddr) ToProto() (*otg.PatternFlowArpSenderProtocolAddr, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowArpSenderProtocolAddr) FromProto(msg *otg.PatternFlowArpSenderProtocolAddr) (PatternFlowArpSenderProtocolAddr, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowArpSenderProtocolAddr) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowArpSenderProtocolAddr) FromPbText(value string) error {
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

func (m *marshalpatternFlowArpSenderProtocolAddr) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowArpSenderProtocolAddr) FromYaml(value string) error {
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

func (m *marshalpatternFlowArpSenderProtocolAddr) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowArpSenderProtocolAddr) FromJson(value string) error {
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

func (obj *patternFlowArpSenderProtocolAddr) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowArpSenderProtocolAddr) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowArpSenderProtocolAddr) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowArpSenderProtocolAddr) Clone() (PatternFlowArpSenderProtocolAddr, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowArpSenderProtocolAddr()
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

func (obj *patternFlowArpSenderProtocolAddr) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowArpSenderProtocolAddr is internetwork address of the sender
type PatternFlowArpSenderProtocolAddr interface {
	Validation
	// msg marshals PatternFlowArpSenderProtocolAddr to protobuf object *otg.PatternFlowArpSenderProtocolAddr
	// and doesn't set defaults
	msg() *otg.PatternFlowArpSenderProtocolAddr
	// setMsg unmarshals PatternFlowArpSenderProtocolAddr from protobuf object *otg.PatternFlowArpSenderProtocolAddr
	// and doesn't set defaults
	setMsg(*otg.PatternFlowArpSenderProtocolAddr) PatternFlowArpSenderProtocolAddr
	// provides marshal interface
	Marshal() marshalPatternFlowArpSenderProtocolAddr
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowArpSenderProtocolAddr
	// validate validates PatternFlowArpSenderProtocolAddr
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowArpSenderProtocolAddr, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowArpSenderProtocolAddrChoiceEnum, set in PatternFlowArpSenderProtocolAddr
	Choice() PatternFlowArpSenderProtocolAddrChoiceEnum
	// setChoice assigns PatternFlowArpSenderProtocolAddrChoiceEnum provided by user to PatternFlowArpSenderProtocolAddr
	setChoice(value PatternFlowArpSenderProtocolAddrChoiceEnum) PatternFlowArpSenderProtocolAddr
	// HasChoice checks if Choice has been set in PatternFlowArpSenderProtocolAddr
	HasChoice() bool
	// Value returns string, set in PatternFlowArpSenderProtocolAddr.
	Value() string
	// SetValue assigns string provided by user to PatternFlowArpSenderProtocolAddr
	SetValue(value string) PatternFlowArpSenderProtocolAddr
	// HasValue checks if Value has been set in PatternFlowArpSenderProtocolAddr
	HasValue() bool
	// Values returns []string, set in PatternFlowArpSenderProtocolAddr.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowArpSenderProtocolAddr
	SetValues(value []string) PatternFlowArpSenderProtocolAddr
	// Increment returns PatternFlowArpSenderProtocolAddrCounter, set in PatternFlowArpSenderProtocolAddr.
	// PatternFlowArpSenderProtocolAddrCounter is ipv4 counter pattern
	Increment() PatternFlowArpSenderProtocolAddrCounter
	// SetIncrement assigns PatternFlowArpSenderProtocolAddrCounter provided by user to PatternFlowArpSenderProtocolAddr.
	// PatternFlowArpSenderProtocolAddrCounter is ipv4 counter pattern
	SetIncrement(value PatternFlowArpSenderProtocolAddrCounter) PatternFlowArpSenderProtocolAddr
	// HasIncrement checks if Increment has been set in PatternFlowArpSenderProtocolAddr
	HasIncrement() bool
	// Decrement returns PatternFlowArpSenderProtocolAddrCounter, set in PatternFlowArpSenderProtocolAddr.
	// PatternFlowArpSenderProtocolAddrCounter is ipv4 counter pattern
	Decrement() PatternFlowArpSenderProtocolAddrCounter
	// SetDecrement assigns PatternFlowArpSenderProtocolAddrCounter provided by user to PatternFlowArpSenderProtocolAddr.
	// PatternFlowArpSenderProtocolAddrCounter is ipv4 counter pattern
	SetDecrement(value PatternFlowArpSenderProtocolAddrCounter) PatternFlowArpSenderProtocolAddr
	// HasDecrement checks if Decrement has been set in PatternFlowArpSenderProtocolAddr
	HasDecrement() bool
	// MetricTags returns PatternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIterIter, set in PatternFlowArpSenderProtocolAddr
	MetricTags() PatternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter
	setNil()
}

type PatternFlowArpSenderProtocolAddrChoiceEnum string

// Enum of Choice on PatternFlowArpSenderProtocolAddr
var PatternFlowArpSenderProtocolAddrChoice = struct {
	VALUE     PatternFlowArpSenderProtocolAddrChoiceEnum
	VALUES    PatternFlowArpSenderProtocolAddrChoiceEnum
	INCREMENT PatternFlowArpSenderProtocolAddrChoiceEnum
	DECREMENT PatternFlowArpSenderProtocolAddrChoiceEnum
}{
	VALUE:     PatternFlowArpSenderProtocolAddrChoiceEnum("value"),
	VALUES:    PatternFlowArpSenderProtocolAddrChoiceEnum("values"),
	INCREMENT: PatternFlowArpSenderProtocolAddrChoiceEnum("increment"),
	DECREMENT: PatternFlowArpSenderProtocolAddrChoiceEnum("decrement"),
}

func (obj *patternFlowArpSenderProtocolAddr) Choice() PatternFlowArpSenderProtocolAddrChoiceEnum {
	return PatternFlowArpSenderProtocolAddrChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowArpSenderProtocolAddr) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowArpSenderProtocolAddr) setChoice(value PatternFlowArpSenderProtocolAddrChoiceEnum) PatternFlowArpSenderProtocolAddr {
	intValue, ok := otg.PatternFlowArpSenderProtocolAddr_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowArpSenderProtocolAddrChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowArpSenderProtocolAddr_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowArpSenderProtocolAddrChoice.VALUE {
		defaultValue := "0.0.0.0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowArpSenderProtocolAddrChoice.VALUES {
		defaultValue := []string{"0.0.0.0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowArpSenderProtocolAddrChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowArpSenderProtocolAddrCounter().msg()
	}

	if value == PatternFlowArpSenderProtocolAddrChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowArpSenderProtocolAddrCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowArpSenderProtocolAddr) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowArpSenderProtocolAddrChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowArpSenderProtocolAddr) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowArpSenderProtocolAddr object
func (obj *patternFlowArpSenderProtocolAddr) SetValue(value string) PatternFlowArpSenderProtocolAddr {
	obj.setChoice(PatternFlowArpSenderProtocolAddrChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowArpSenderProtocolAddr) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"0.0.0.0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowArpSenderProtocolAddr object
func (obj *patternFlowArpSenderProtocolAddr) SetValues(value []string) PatternFlowArpSenderProtocolAddr {
	obj.setChoice(PatternFlowArpSenderProtocolAddrChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowArpSenderProtocolAddrCounter
func (obj *patternFlowArpSenderProtocolAddr) Increment() PatternFlowArpSenderProtocolAddrCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowArpSenderProtocolAddrChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowArpSenderProtocolAddrCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowArpSenderProtocolAddrCounter
func (obj *patternFlowArpSenderProtocolAddr) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowArpSenderProtocolAddrCounter value in the PatternFlowArpSenderProtocolAddr object
func (obj *patternFlowArpSenderProtocolAddr) SetIncrement(value PatternFlowArpSenderProtocolAddrCounter) PatternFlowArpSenderProtocolAddr {
	obj.setChoice(PatternFlowArpSenderProtocolAddrChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowArpSenderProtocolAddrCounter
func (obj *patternFlowArpSenderProtocolAddr) Decrement() PatternFlowArpSenderProtocolAddrCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowArpSenderProtocolAddrChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowArpSenderProtocolAddrCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowArpSenderProtocolAddrCounter
func (obj *patternFlowArpSenderProtocolAddr) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowArpSenderProtocolAddrCounter value in the PatternFlowArpSenderProtocolAddr object
func (obj *patternFlowArpSenderProtocolAddr) SetDecrement(value PatternFlowArpSenderProtocolAddrCounter) PatternFlowArpSenderProtocolAddr {
	obj.setChoice(PatternFlowArpSenderProtocolAddrChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowArpSenderProtocolAddrMetricTag
func (obj *patternFlowArpSenderProtocolAddr) MetricTags() PatternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowArpSenderProtocolAddrMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter struct {
	obj                                            *patternFlowArpSenderProtocolAddr
	patternFlowArpSenderProtocolAddrMetricTagSlice []PatternFlowArpSenderProtocolAddrMetricTag
	fieldPtr                                       *[]*otg.PatternFlowArpSenderProtocolAddrMetricTag
}

func newPatternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter(ptr *[]*otg.PatternFlowArpSenderProtocolAddrMetricTag) PatternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter {
	return &patternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter{fieldPtr: ptr}
}

type PatternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter interface {
	setMsg(*patternFlowArpSenderProtocolAddr) PatternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter
	Items() []PatternFlowArpSenderProtocolAddrMetricTag
	Add() PatternFlowArpSenderProtocolAddrMetricTag
	Append(items ...PatternFlowArpSenderProtocolAddrMetricTag) PatternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter
	Set(index int, newObj PatternFlowArpSenderProtocolAddrMetricTag) PatternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter
	Clear() PatternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter
	clearHolderSlice() PatternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter
	appendHolderSlice(item PatternFlowArpSenderProtocolAddrMetricTag) PatternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter
}

func (obj *patternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter) setMsg(msg *patternFlowArpSenderProtocolAddr) PatternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowArpSenderProtocolAddrMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter) Items() []PatternFlowArpSenderProtocolAddrMetricTag {
	return obj.patternFlowArpSenderProtocolAddrMetricTagSlice
}

func (obj *patternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter) Add() PatternFlowArpSenderProtocolAddrMetricTag {
	newObj := &otg.PatternFlowArpSenderProtocolAddrMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowArpSenderProtocolAddrMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowArpSenderProtocolAddrMetricTagSlice = append(obj.patternFlowArpSenderProtocolAddrMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter) Append(items ...PatternFlowArpSenderProtocolAddrMetricTag) PatternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowArpSenderProtocolAddrMetricTagSlice = append(obj.patternFlowArpSenderProtocolAddrMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter) Set(index int, newObj PatternFlowArpSenderProtocolAddrMetricTag) PatternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowArpSenderProtocolAddrMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter) Clear() PatternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowArpSenderProtocolAddrMetricTag{}
		obj.patternFlowArpSenderProtocolAddrMetricTagSlice = []PatternFlowArpSenderProtocolAddrMetricTag{}
	}
	return obj
}
func (obj *patternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter) clearHolderSlice() PatternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter {
	if len(obj.patternFlowArpSenderProtocolAddrMetricTagSlice) > 0 {
		obj.patternFlowArpSenderProtocolAddrMetricTagSlice = []PatternFlowArpSenderProtocolAddrMetricTag{}
	}
	return obj
}
func (obj *patternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter) appendHolderSlice(item PatternFlowArpSenderProtocolAddrMetricTag) PatternFlowArpSenderProtocolAddrPatternFlowArpSenderProtocolAddrMetricTagIter {
	obj.patternFlowArpSenderProtocolAddrMetricTagSlice = append(obj.patternFlowArpSenderProtocolAddrMetricTagSlice, item)
	return obj
}

func (obj *patternFlowArpSenderProtocolAddr) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpSenderProtocolAddr.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpSenderProtocolAddr.Values"))
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
				obj.MetricTags().appendHolderSlice(&patternFlowArpSenderProtocolAddrMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowArpSenderProtocolAddr) setDefault() {
	var choices_set int = 0
	var choice PatternFlowArpSenderProtocolAddrChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowArpSenderProtocolAddrChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowArpSenderProtocolAddrChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowArpSenderProtocolAddrChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowArpSenderProtocolAddrChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowArpSenderProtocolAddrChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowArpSenderProtocolAddr")
			}
		} else {
			intVal := otg.PatternFlowArpSenderProtocolAddr_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowArpSenderProtocolAddr_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
