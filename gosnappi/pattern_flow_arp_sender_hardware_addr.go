package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowArpSenderHardwareAddr *****
type patternFlowArpSenderHardwareAddr struct {
	validation
	obj              *otg.PatternFlowArpSenderHardwareAddr
	marshaller       marshalPatternFlowArpSenderHardwareAddr
	unMarshaller     unMarshalPatternFlowArpSenderHardwareAddr
	incrementHolder  PatternFlowArpSenderHardwareAddrCounter
	decrementHolder  PatternFlowArpSenderHardwareAddrCounter
	metricTagsHolder PatternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter
}

func NewPatternFlowArpSenderHardwareAddr() PatternFlowArpSenderHardwareAddr {
	obj := patternFlowArpSenderHardwareAddr{obj: &otg.PatternFlowArpSenderHardwareAddr{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowArpSenderHardwareAddr) msg() *otg.PatternFlowArpSenderHardwareAddr {
	return obj.obj
}

func (obj *patternFlowArpSenderHardwareAddr) setMsg(msg *otg.PatternFlowArpSenderHardwareAddr) PatternFlowArpSenderHardwareAddr {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowArpSenderHardwareAddr struct {
	obj *patternFlowArpSenderHardwareAddr
}

type marshalPatternFlowArpSenderHardwareAddr interface {
	// ToProto marshals PatternFlowArpSenderHardwareAddr to protobuf object *otg.PatternFlowArpSenderHardwareAddr
	ToProto() (*otg.PatternFlowArpSenderHardwareAddr, error)
	// ToPbText marshals PatternFlowArpSenderHardwareAddr to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowArpSenderHardwareAddr to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowArpSenderHardwareAddr to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowArpSenderHardwareAddr struct {
	obj *patternFlowArpSenderHardwareAddr
}

type unMarshalPatternFlowArpSenderHardwareAddr interface {
	// FromProto unmarshals PatternFlowArpSenderHardwareAddr from protobuf object *otg.PatternFlowArpSenderHardwareAddr
	FromProto(msg *otg.PatternFlowArpSenderHardwareAddr) (PatternFlowArpSenderHardwareAddr, error)
	// FromPbText unmarshals PatternFlowArpSenderHardwareAddr from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowArpSenderHardwareAddr from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowArpSenderHardwareAddr from JSON text
	FromJson(value string) error
}

func (obj *patternFlowArpSenderHardwareAddr) Marshal() marshalPatternFlowArpSenderHardwareAddr {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowArpSenderHardwareAddr{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowArpSenderHardwareAddr) Unmarshal() unMarshalPatternFlowArpSenderHardwareAddr {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowArpSenderHardwareAddr{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowArpSenderHardwareAddr) ToProto() (*otg.PatternFlowArpSenderHardwareAddr, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowArpSenderHardwareAddr) FromProto(msg *otg.PatternFlowArpSenderHardwareAddr) (PatternFlowArpSenderHardwareAddr, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowArpSenderHardwareAddr) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowArpSenderHardwareAddr) FromPbText(value string) error {
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

func (m *marshalpatternFlowArpSenderHardwareAddr) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowArpSenderHardwareAddr) FromYaml(value string) error {
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

func (m *marshalpatternFlowArpSenderHardwareAddr) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowArpSenderHardwareAddr) FromJson(value string) error {
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

func (obj *patternFlowArpSenderHardwareAddr) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowArpSenderHardwareAddr) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowArpSenderHardwareAddr) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowArpSenderHardwareAddr) Clone() (PatternFlowArpSenderHardwareAddr, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowArpSenderHardwareAddr()
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

func (obj *patternFlowArpSenderHardwareAddr) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowArpSenderHardwareAddr is media address of the sender
type PatternFlowArpSenderHardwareAddr interface {
	Validation
	// msg marshals PatternFlowArpSenderHardwareAddr to protobuf object *otg.PatternFlowArpSenderHardwareAddr
	// and doesn't set defaults
	msg() *otg.PatternFlowArpSenderHardwareAddr
	// setMsg unmarshals PatternFlowArpSenderHardwareAddr from protobuf object *otg.PatternFlowArpSenderHardwareAddr
	// and doesn't set defaults
	setMsg(*otg.PatternFlowArpSenderHardwareAddr) PatternFlowArpSenderHardwareAddr
	// provides marshal interface
	Marshal() marshalPatternFlowArpSenderHardwareAddr
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowArpSenderHardwareAddr
	// validate validates PatternFlowArpSenderHardwareAddr
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowArpSenderHardwareAddr, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowArpSenderHardwareAddrChoiceEnum, set in PatternFlowArpSenderHardwareAddr
	Choice() PatternFlowArpSenderHardwareAddrChoiceEnum
	// setChoice assigns PatternFlowArpSenderHardwareAddrChoiceEnum provided by user to PatternFlowArpSenderHardwareAddr
	setChoice(value PatternFlowArpSenderHardwareAddrChoiceEnum) PatternFlowArpSenderHardwareAddr
	// HasChoice checks if Choice has been set in PatternFlowArpSenderHardwareAddr
	HasChoice() bool
	// Value returns string, set in PatternFlowArpSenderHardwareAddr.
	Value() string
	// SetValue assigns string provided by user to PatternFlowArpSenderHardwareAddr
	SetValue(value string) PatternFlowArpSenderHardwareAddr
	// HasValue checks if Value has been set in PatternFlowArpSenderHardwareAddr
	HasValue() bool
	// Values returns []string, set in PatternFlowArpSenderHardwareAddr.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowArpSenderHardwareAddr
	SetValues(value []string) PatternFlowArpSenderHardwareAddr
	// Increment returns PatternFlowArpSenderHardwareAddrCounter, set in PatternFlowArpSenderHardwareAddr.
	// PatternFlowArpSenderHardwareAddrCounter is mac counter pattern
	Increment() PatternFlowArpSenderHardwareAddrCounter
	// SetIncrement assigns PatternFlowArpSenderHardwareAddrCounter provided by user to PatternFlowArpSenderHardwareAddr.
	// PatternFlowArpSenderHardwareAddrCounter is mac counter pattern
	SetIncrement(value PatternFlowArpSenderHardwareAddrCounter) PatternFlowArpSenderHardwareAddr
	// HasIncrement checks if Increment has been set in PatternFlowArpSenderHardwareAddr
	HasIncrement() bool
	// Decrement returns PatternFlowArpSenderHardwareAddrCounter, set in PatternFlowArpSenderHardwareAddr.
	// PatternFlowArpSenderHardwareAddrCounter is mac counter pattern
	Decrement() PatternFlowArpSenderHardwareAddrCounter
	// SetDecrement assigns PatternFlowArpSenderHardwareAddrCounter provided by user to PatternFlowArpSenderHardwareAddr.
	// PatternFlowArpSenderHardwareAddrCounter is mac counter pattern
	SetDecrement(value PatternFlowArpSenderHardwareAddrCounter) PatternFlowArpSenderHardwareAddr
	// HasDecrement checks if Decrement has been set in PatternFlowArpSenderHardwareAddr
	HasDecrement() bool
	// MetricTags returns PatternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIterIter, set in PatternFlowArpSenderHardwareAddr
	MetricTags() PatternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter
	setNil()
}

type PatternFlowArpSenderHardwareAddrChoiceEnum string

// Enum of Choice on PatternFlowArpSenderHardwareAddr
var PatternFlowArpSenderHardwareAddrChoice = struct {
	VALUE     PatternFlowArpSenderHardwareAddrChoiceEnum
	VALUES    PatternFlowArpSenderHardwareAddrChoiceEnum
	INCREMENT PatternFlowArpSenderHardwareAddrChoiceEnum
	DECREMENT PatternFlowArpSenderHardwareAddrChoiceEnum
}{
	VALUE:     PatternFlowArpSenderHardwareAddrChoiceEnum("value"),
	VALUES:    PatternFlowArpSenderHardwareAddrChoiceEnum("values"),
	INCREMENT: PatternFlowArpSenderHardwareAddrChoiceEnum("increment"),
	DECREMENT: PatternFlowArpSenderHardwareAddrChoiceEnum("decrement"),
}

func (obj *patternFlowArpSenderHardwareAddr) Choice() PatternFlowArpSenderHardwareAddrChoiceEnum {
	return PatternFlowArpSenderHardwareAddrChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowArpSenderHardwareAddr) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowArpSenderHardwareAddr) setChoice(value PatternFlowArpSenderHardwareAddrChoiceEnum) PatternFlowArpSenderHardwareAddr {
	intValue, ok := otg.PatternFlowArpSenderHardwareAddr_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowArpSenderHardwareAddrChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowArpSenderHardwareAddr_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowArpSenderHardwareAddrChoice.VALUE {
		defaultValue := "00:00:00:00:00:00"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowArpSenderHardwareAddrChoice.VALUES {
		defaultValue := []string{"00:00:00:00:00:00"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowArpSenderHardwareAddrChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowArpSenderHardwareAddrCounter().msg()
	}

	if value == PatternFlowArpSenderHardwareAddrChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowArpSenderHardwareAddrCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowArpSenderHardwareAddr) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowArpSenderHardwareAddrChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowArpSenderHardwareAddr) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowArpSenderHardwareAddr object
func (obj *patternFlowArpSenderHardwareAddr) SetValue(value string) PatternFlowArpSenderHardwareAddr {
	obj.setChoice(PatternFlowArpSenderHardwareAddrChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowArpSenderHardwareAddr) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"00:00:00:00:00:00"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowArpSenderHardwareAddr object
func (obj *patternFlowArpSenderHardwareAddr) SetValues(value []string) PatternFlowArpSenderHardwareAddr {
	obj.setChoice(PatternFlowArpSenderHardwareAddrChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowArpSenderHardwareAddrCounter
func (obj *patternFlowArpSenderHardwareAddr) Increment() PatternFlowArpSenderHardwareAddrCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowArpSenderHardwareAddrChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowArpSenderHardwareAddrCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowArpSenderHardwareAddrCounter
func (obj *patternFlowArpSenderHardwareAddr) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowArpSenderHardwareAddrCounter value in the PatternFlowArpSenderHardwareAddr object
func (obj *patternFlowArpSenderHardwareAddr) SetIncrement(value PatternFlowArpSenderHardwareAddrCounter) PatternFlowArpSenderHardwareAddr {
	obj.setChoice(PatternFlowArpSenderHardwareAddrChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowArpSenderHardwareAddrCounter
func (obj *patternFlowArpSenderHardwareAddr) Decrement() PatternFlowArpSenderHardwareAddrCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowArpSenderHardwareAddrChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowArpSenderHardwareAddrCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowArpSenderHardwareAddrCounter
func (obj *patternFlowArpSenderHardwareAddr) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowArpSenderHardwareAddrCounter value in the PatternFlowArpSenderHardwareAddr object
func (obj *patternFlowArpSenderHardwareAddr) SetDecrement(value PatternFlowArpSenderHardwareAddrCounter) PatternFlowArpSenderHardwareAddr {
	obj.setChoice(PatternFlowArpSenderHardwareAddrChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowArpSenderHardwareAddrMetricTag
func (obj *patternFlowArpSenderHardwareAddr) MetricTags() PatternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowArpSenderHardwareAddrMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter struct {
	obj                                            *patternFlowArpSenderHardwareAddr
	patternFlowArpSenderHardwareAddrMetricTagSlice []PatternFlowArpSenderHardwareAddrMetricTag
	fieldPtr                                       *[]*otg.PatternFlowArpSenderHardwareAddrMetricTag
}

func newPatternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter(ptr *[]*otg.PatternFlowArpSenderHardwareAddrMetricTag) PatternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter {
	return &patternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter{fieldPtr: ptr}
}

type PatternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter interface {
	setMsg(*patternFlowArpSenderHardwareAddr) PatternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter
	Items() []PatternFlowArpSenderHardwareAddrMetricTag
	Add() PatternFlowArpSenderHardwareAddrMetricTag
	Append(items ...PatternFlowArpSenderHardwareAddrMetricTag) PatternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter
	Set(index int, newObj PatternFlowArpSenderHardwareAddrMetricTag) PatternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter
	Clear() PatternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter
	clearHolderSlice() PatternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter
	appendHolderSlice(item PatternFlowArpSenderHardwareAddrMetricTag) PatternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter
}

func (obj *patternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter) setMsg(msg *patternFlowArpSenderHardwareAddr) PatternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowArpSenderHardwareAddrMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter) Items() []PatternFlowArpSenderHardwareAddrMetricTag {
	return obj.patternFlowArpSenderHardwareAddrMetricTagSlice
}

func (obj *patternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter) Add() PatternFlowArpSenderHardwareAddrMetricTag {
	newObj := &otg.PatternFlowArpSenderHardwareAddrMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowArpSenderHardwareAddrMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowArpSenderHardwareAddrMetricTagSlice = append(obj.patternFlowArpSenderHardwareAddrMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter) Append(items ...PatternFlowArpSenderHardwareAddrMetricTag) PatternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowArpSenderHardwareAddrMetricTagSlice = append(obj.patternFlowArpSenderHardwareAddrMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter) Set(index int, newObj PatternFlowArpSenderHardwareAddrMetricTag) PatternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowArpSenderHardwareAddrMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter) Clear() PatternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowArpSenderHardwareAddrMetricTag{}
		obj.patternFlowArpSenderHardwareAddrMetricTagSlice = []PatternFlowArpSenderHardwareAddrMetricTag{}
	}
	return obj
}
func (obj *patternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter) clearHolderSlice() PatternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter {
	if len(obj.patternFlowArpSenderHardwareAddrMetricTagSlice) > 0 {
		obj.patternFlowArpSenderHardwareAddrMetricTagSlice = []PatternFlowArpSenderHardwareAddrMetricTag{}
	}
	return obj
}
func (obj *patternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter) appendHolderSlice(item PatternFlowArpSenderHardwareAddrMetricTag) PatternFlowArpSenderHardwareAddrPatternFlowArpSenderHardwareAddrMetricTagIter {
	obj.patternFlowArpSenderHardwareAddrMetricTagSlice = append(obj.patternFlowArpSenderHardwareAddrMetricTagSlice, item)
	return obj
}

func (obj *patternFlowArpSenderHardwareAddr) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpSenderHardwareAddr.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateMacSlice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpSenderHardwareAddr.Values"))
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
				obj.MetricTags().appendHolderSlice(&patternFlowArpSenderHardwareAddrMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowArpSenderHardwareAddr) setDefault() {
	var choices_set int = 0
	var choice PatternFlowArpSenderHardwareAddrChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowArpSenderHardwareAddrChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowArpSenderHardwareAddrChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowArpSenderHardwareAddrChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowArpSenderHardwareAddrChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowArpSenderHardwareAddrChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowArpSenderHardwareAddr")
			}
		} else {
			intVal := otg.PatternFlowArpSenderHardwareAddr_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowArpSenderHardwareAddr_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
