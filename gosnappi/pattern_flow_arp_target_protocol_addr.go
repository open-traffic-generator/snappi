package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowArpTargetProtocolAddr *****
type patternFlowArpTargetProtocolAddr struct {
	validation
	obj              *otg.PatternFlowArpTargetProtocolAddr
	marshaller       marshalPatternFlowArpTargetProtocolAddr
	unMarshaller     unMarshalPatternFlowArpTargetProtocolAddr
	incrementHolder  PatternFlowArpTargetProtocolAddrCounter
	decrementHolder  PatternFlowArpTargetProtocolAddrCounter
	metricTagsHolder PatternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter
}

func NewPatternFlowArpTargetProtocolAddr() PatternFlowArpTargetProtocolAddr {
	obj := patternFlowArpTargetProtocolAddr{obj: &otg.PatternFlowArpTargetProtocolAddr{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowArpTargetProtocolAddr) msg() *otg.PatternFlowArpTargetProtocolAddr {
	return obj.obj
}

func (obj *patternFlowArpTargetProtocolAddr) setMsg(msg *otg.PatternFlowArpTargetProtocolAddr) PatternFlowArpTargetProtocolAddr {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowArpTargetProtocolAddr struct {
	obj *patternFlowArpTargetProtocolAddr
}

type marshalPatternFlowArpTargetProtocolAddr interface {
	// ToProto marshals PatternFlowArpTargetProtocolAddr to protobuf object *otg.PatternFlowArpTargetProtocolAddr
	ToProto() (*otg.PatternFlowArpTargetProtocolAddr, error)
	// ToPbText marshals PatternFlowArpTargetProtocolAddr to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowArpTargetProtocolAddr to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowArpTargetProtocolAddr to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowArpTargetProtocolAddr to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowArpTargetProtocolAddr struct {
	obj *patternFlowArpTargetProtocolAddr
}

type unMarshalPatternFlowArpTargetProtocolAddr interface {
	// FromProto unmarshals PatternFlowArpTargetProtocolAddr from protobuf object *otg.PatternFlowArpTargetProtocolAddr
	FromProto(msg *otg.PatternFlowArpTargetProtocolAddr) (PatternFlowArpTargetProtocolAddr, error)
	// FromPbText unmarshals PatternFlowArpTargetProtocolAddr from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowArpTargetProtocolAddr from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowArpTargetProtocolAddr from JSON text
	FromJson(value string) error
}

func (obj *patternFlowArpTargetProtocolAddr) Marshal() marshalPatternFlowArpTargetProtocolAddr {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowArpTargetProtocolAddr{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowArpTargetProtocolAddr) Unmarshal() unMarshalPatternFlowArpTargetProtocolAddr {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowArpTargetProtocolAddr{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowArpTargetProtocolAddr) ToProto() (*otg.PatternFlowArpTargetProtocolAddr, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowArpTargetProtocolAddr) FromProto(msg *otg.PatternFlowArpTargetProtocolAddr) (PatternFlowArpTargetProtocolAddr, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowArpTargetProtocolAddr) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowArpTargetProtocolAddr) FromPbText(value string) error {
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

func (m *marshalpatternFlowArpTargetProtocolAddr) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowArpTargetProtocolAddr) FromYaml(value string) error {
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

func (m *marshalpatternFlowArpTargetProtocolAddr) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalpatternFlowArpTargetProtocolAddr) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowArpTargetProtocolAddr) FromJson(value string) error {
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

func (obj *patternFlowArpTargetProtocolAddr) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowArpTargetProtocolAddr) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowArpTargetProtocolAddr) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowArpTargetProtocolAddr) Clone() (PatternFlowArpTargetProtocolAddr, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowArpTargetProtocolAddr()
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

func (obj *patternFlowArpTargetProtocolAddr) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowArpTargetProtocolAddr is internetwork address of the target
type PatternFlowArpTargetProtocolAddr interface {
	Validation
	// msg marshals PatternFlowArpTargetProtocolAddr to protobuf object *otg.PatternFlowArpTargetProtocolAddr
	// and doesn't set defaults
	msg() *otg.PatternFlowArpTargetProtocolAddr
	// setMsg unmarshals PatternFlowArpTargetProtocolAddr from protobuf object *otg.PatternFlowArpTargetProtocolAddr
	// and doesn't set defaults
	setMsg(*otg.PatternFlowArpTargetProtocolAddr) PatternFlowArpTargetProtocolAddr
	// provides marshal interface
	Marshal() marshalPatternFlowArpTargetProtocolAddr
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowArpTargetProtocolAddr
	// validate validates PatternFlowArpTargetProtocolAddr
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowArpTargetProtocolAddr, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowArpTargetProtocolAddrChoiceEnum, set in PatternFlowArpTargetProtocolAddr
	Choice() PatternFlowArpTargetProtocolAddrChoiceEnum
	// setChoice assigns PatternFlowArpTargetProtocolAddrChoiceEnum provided by user to PatternFlowArpTargetProtocolAddr
	setChoice(value PatternFlowArpTargetProtocolAddrChoiceEnum) PatternFlowArpTargetProtocolAddr
	// HasChoice checks if Choice has been set in PatternFlowArpTargetProtocolAddr
	HasChoice() bool
	// Value returns string, set in PatternFlowArpTargetProtocolAddr.
	Value() string
	// SetValue assigns string provided by user to PatternFlowArpTargetProtocolAddr
	SetValue(value string) PatternFlowArpTargetProtocolAddr
	// HasValue checks if Value has been set in PatternFlowArpTargetProtocolAddr
	HasValue() bool
	// Values returns []string, set in PatternFlowArpTargetProtocolAddr.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowArpTargetProtocolAddr
	SetValues(value []string) PatternFlowArpTargetProtocolAddr
	// Increment returns PatternFlowArpTargetProtocolAddrCounter, set in PatternFlowArpTargetProtocolAddr.
	// PatternFlowArpTargetProtocolAddrCounter is ipv4 counter pattern
	Increment() PatternFlowArpTargetProtocolAddrCounter
	// SetIncrement assigns PatternFlowArpTargetProtocolAddrCounter provided by user to PatternFlowArpTargetProtocolAddr.
	// PatternFlowArpTargetProtocolAddrCounter is ipv4 counter pattern
	SetIncrement(value PatternFlowArpTargetProtocolAddrCounter) PatternFlowArpTargetProtocolAddr
	// HasIncrement checks if Increment has been set in PatternFlowArpTargetProtocolAddr
	HasIncrement() bool
	// Decrement returns PatternFlowArpTargetProtocolAddrCounter, set in PatternFlowArpTargetProtocolAddr.
	// PatternFlowArpTargetProtocolAddrCounter is ipv4 counter pattern
	Decrement() PatternFlowArpTargetProtocolAddrCounter
	// SetDecrement assigns PatternFlowArpTargetProtocolAddrCounter provided by user to PatternFlowArpTargetProtocolAddr.
	// PatternFlowArpTargetProtocolAddrCounter is ipv4 counter pattern
	SetDecrement(value PatternFlowArpTargetProtocolAddrCounter) PatternFlowArpTargetProtocolAddr
	// HasDecrement checks if Decrement has been set in PatternFlowArpTargetProtocolAddr
	HasDecrement() bool
	// MetricTags returns PatternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIterIter, set in PatternFlowArpTargetProtocolAddr
	MetricTags() PatternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter
	setNil()
}

type PatternFlowArpTargetProtocolAddrChoiceEnum string

// Enum of Choice on PatternFlowArpTargetProtocolAddr
var PatternFlowArpTargetProtocolAddrChoice = struct {
	VALUE     PatternFlowArpTargetProtocolAddrChoiceEnum
	VALUES    PatternFlowArpTargetProtocolAddrChoiceEnum
	INCREMENT PatternFlowArpTargetProtocolAddrChoiceEnum
	DECREMENT PatternFlowArpTargetProtocolAddrChoiceEnum
}{
	VALUE:     PatternFlowArpTargetProtocolAddrChoiceEnum("value"),
	VALUES:    PatternFlowArpTargetProtocolAddrChoiceEnum("values"),
	INCREMENT: PatternFlowArpTargetProtocolAddrChoiceEnum("increment"),
	DECREMENT: PatternFlowArpTargetProtocolAddrChoiceEnum("decrement"),
}

func (obj *patternFlowArpTargetProtocolAddr) Choice() PatternFlowArpTargetProtocolAddrChoiceEnum {
	return PatternFlowArpTargetProtocolAddrChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowArpTargetProtocolAddr) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowArpTargetProtocolAddr) setChoice(value PatternFlowArpTargetProtocolAddrChoiceEnum) PatternFlowArpTargetProtocolAddr {
	intValue, ok := otg.PatternFlowArpTargetProtocolAddr_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowArpTargetProtocolAddrChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowArpTargetProtocolAddr_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowArpTargetProtocolAddrChoice.VALUE {
		defaultValue := "0.0.0.0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowArpTargetProtocolAddrChoice.VALUES {
		defaultValue := []string{"0.0.0.0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowArpTargetProtocolAddrChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowArpTargetProtocolAddrCounter().msg()
	}

	if value == PatternFlowArpTargetProtocolAddrChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowArpTargetProtocolAddrCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowArpTargetProtocolAddr) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowArpTargetProtocolAddrChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowArpTargetProtocolAddr) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowArpTargetProtocolAddr object
func (obj *patternFlowArpTargetProtocolAddr) SetValue(value string) PatternFlowArpTargetProtocolAddr {
	obj.setChoice(PatternFlowArpTargetProtocolAddrChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowArpTargetProtocolAddr) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"0.0.0.0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowArpTargetProtocolAddr object
func (obj *patternFlowArpTargetProtocolAddr) SetValues(value []string) PatternFlowArpTargetProtocolAddr {
	obj.setChoice(PatternFlowArpTargetProtocolAddrChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowArpTargetProtocolAddrCounter
func (obj *patternFlowArpTargetProtocolAddr) Increment() PatternFlowArpTargetProtocolAddrCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowArpTargetProtocolAddrChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowArpTargetProtocolAddrCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowArpTargetProtocolAddrCounter
func (obj *patternFlowArpTargetProtocolAddr) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowArpTargetProtocolAddrCounter value in the PatternFlowArpTargetProtocolAddr object
func (obj *patternFlowArpTargetProtocolAddr) SetIncrement(value PatternFlowArpTargetProtocolAddrCounter) PatternFlowArpTargetProtocolAddr {
	obj.setChoice(PatternFlowArpTargetProtocolAddrChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowArpTargetProtocolAddrCounter
func (obj *patternFlowArpTargetProtocolAddr) Decrement() PatternFlowArpTargetProtocolAddrCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowArpTargetProtocolAddrChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowArpTargetProtocolAddrCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowArpTargetProtocolAddrCounter
func (obj *patternFlowArpTargetProtocolAddr) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowArpTargetProtocolAddrCounter value in the PatternFlowArpTargetProtocolAddr object
func (obj *patternFlowArpTargetProtocolAddr) SetDecrement(value PatternFlowArpTargetProtocolAddrCounter) PatternFlowArpTargetProtocolAddr {
	obj.setChoice(PatternFlowArpTargetProtocolAddrChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowArpTargetProtocolAddrMetricTag
func (obj *patternFlowArpTargetProtocolAddr) MetricTags() PatternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowArpTargetProtocolAddrMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter struct {
	obj                                            *patternFlowArpTargetProtocolAddr
	patternFlowArpTargetProtocolAddrMetricTagSlice []PatternFlowArpTargetProtocolAddrMetricTag
	fieldPtr                                       *[]*otg.PatternFlowArpTargetProtocolAddrMetricTag
}

func newPatternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter(ptr *[]*otg.PatternFlowArpTargetProtocolAddrMetricTag) PatternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter {
	return &patternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter{fieldPtr: ptr}
}

type PatternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter interface {
	setMsg(*patternFlowArpTargetProtocolAddr) PatternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter
	Items() []PatternFlowArpTargetProtocolAddrMetricTag
	Add() PatternFlowArpTargetProtocolAddrMetricTag
	Append(items ...PatternFlowArpTargetProtocolAddrMetricTag) PatternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter
	Set(index int, newObj PatternFlowArpTargetProtocolAddrMetricTag) PatternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter
	Clear() PatternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter
	clearHolderSlice() PatternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter
	appendHolderSlice(item PatternFlowArpTargetProtocolAddrMetricTag) PatternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter
}

func (obj *patternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter) setMsg(msg *patternFlowArpTargetProtocolAddr) PatternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowArpTargetProtocolAddrMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter) Items() []PatternFlowArpTargetProtocolAddrMetricTag {
	return obj.patternFlowArpTargetProtocolAddrMetricTagSlice
}

func (obj *patternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter) Add() PatternFlowArpTargetProtocolAddrMetricTag {
	newObj := &otg.PatternFlowArpTargetProtocolAddrMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowArpTargetProtocolAddrMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowArpTargetProtocolAddrMetricTagSlice = append(obj.patternFlowArpTargetProtocolAddrMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter) Append(items ...PatternFlowArpTargetProtocolAddrMetricTag) PatternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowArpTargetProtocolAddrMetricTagSlice = append(obj.patternFlowArpTargetProtocolAddrMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter) Set(index int, newObj PatternFlowArpTargetProtocolAddrMetricTag) PatternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowArpTargetProtocolAddrMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter) Clear() PatternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowArpTargetProtocolAddrMetricTag{}
		obj.patternFlowArpTargetProtocolAddrMetricTagSlice = []PatternFlowArpTargetProtocolAddrMetricTag{}
	}
	return obj
}
func (obj *patternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter) clearHolderSlice() PatternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter {
	if len(obj.patternFlowArpTargetProtocolAddrMetricTagSlice) > 0 {
		obj.patternFlowArpTargetProtocolAddrMetricTagSlice = []PatternFlowArpTargetProtocolAddrMetricTag{}
	}
	return obj
}
func (obj *patternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter) appendHolderSlice(item PatternFlowArpTargetProtocolAddrMetricTag) PatternFlowArpTargetProtocolAddrPatternFlowArpTargetProtocolAddrMetricTagIter {
	obj.patternFlowArpTargetProtocolAddrMetricTagSlice = append(obj.patternFlowArpTargetProtocolAddrMetricTagSlice, item)
	return obj
}

func (obj *patternFlowArpTargetProtocolAddr) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpTargetProtocolAddr.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpTargetProtocolAddr.Values"))
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
				obj.MetricTags().appendHolderSlice(&patternFlowArpTargetProtocolAddrMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowArpTargetProtocolAddr) setDefault() {
	var choices_set int = 0
	var choice PatternFlowArpTargetProtocolAddrChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowArpTargetProtocolAddrChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowArpTargetProtocolAddrChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowArpTargetProtocolAddrChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowArpTargetProtocolAddrChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowArpTargetProtocolAddrChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowArpTargetProtocolAddr")
			}
		} else {
			intVal := otg.PatternFlowArpTargetProtocolAddr_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowArpTargetProtocolAddr_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
