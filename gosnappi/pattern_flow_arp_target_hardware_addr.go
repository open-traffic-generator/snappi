package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowArpTargetHardwareAddr *****
type patternFlowArpTargetHardwareAddr struct {
	validation
	obj              *otg.PatternFlowArpTargetHardwareAddr
	marshaller       marshalPatternFlowArpTargetHardwareAddr
	unMarshaller     unMarshalPatternFlowArpTargetHardwareAddr
	incrementHolder  PatternFlowArpTargetHardwareAddrCounter
	decrementHolder  PatternFlowArpTargetHardwareAddrCounter
	metricTagsHolder PatternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter
}

func NewPatternFlowArpTargetHardwareAddr() PatternFlowArpTargetHardwareAddr {
	obj := patternFlowArpTargetHardwareAddr{obj: &otg.PatternFlowArpTargetHardwareAddr{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowArpTargetHardwareAddr) msg() *otg.PatternFlowArpTargetHardwareAddr {
	return obj.obj
}

func (obj *patternFlowArpTargetHardwareAddr) setMsg(msg *otg.PatternFlowArpTargetHardwareAddr) PatternFlowArpTargetHardwareAddr {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowArpTargetHardwareAddr struct {
	obj *patternFlowArpTargetHardwareAddr
}

type marshalPatternFlowArpTargetHardwareAddr interface {
	// ToProto marshals PatternFlowArpTargetHardwareAddr to protobuf object *otg.PatternFlowArpTargetHardwareAddr
	ToProto() (*otg.PatternFlowArpTargetHardwareAddr, error)
	// ToPbText marshals PatternFlowArpTargetHardwareAddr to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowArpTargetHardwareAddr to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowArpTargetHardwareAddr to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowArpTargetHardwareAddr struct {
	obj *patternFlowArpTargetHardwareAddr
}

type unMarshalPatternFlowArpTargetHardwareAddr interface {
	// FromProto unmarshals PatternFlowArpTargetHardwareAddr from protobuf object *otg.PatternFlowArpTargetHardwareAddr
	FromProto(msg *otg.PatternFlowArpTargetHardwareAddr) (PatternFlowArpTargetHardwareAddr, error)
	// FromPbText unmarshals PatternFlowArpTargetHardwareAddr from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowArpTargetHardwareAddr from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowArpTargetHardwareAddr from JSON text
	FromJson(value string) error
}

func (obj *patternFlowArpTargetHardwareAddr) Marshal() marshalPatternFlowArpTargetHardwareAddr {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowArpTargetHardwareAddr{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowArpTargetHardwareAddr) Unmarshal() unMarshalPatternFlowArpTargetHardwareAddr {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowArpTargetHardwareAddr{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowArpTargetHardwareAddr) ToProto() (*otg.PatternFlowArpTargetHardwareAddr, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowArpTargetHardwareAddr) FromProto(msg *otg.PatternFlowArpTargetHardwareAddr) (PatternFlowArpTargetHardwareAddr, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowArpTargetHardwareAddr) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowArpTargetHardwareAddr) FromPbText(value string) error {
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

func (m *marshalpatternFlowArpTargetHardwareAddr) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowArpTargetHardwareAddr) FromYaml(value string) error {
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

func (m *marshalpatternFlowArpTargetHardwareAddr) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowArpTargetHardwareAddr) FromJson(value string) error {
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

func (obj *patternFlowArpTargetHardwareAddr) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowArpTargetHardwareAddr) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowArpTargetHardwareAddr) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowArpTargetHardwareAddr) Clone() (PatternFlowArpTargetHardwareAddr, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowArpTargetHardwareAddr()
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

func (obj *patternFlowArpTargetHardwareAddr) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowArpTargetHardwareAddr is media address of the target
type PatternFlowArpTargetHardwareAddr interface {
	Validation
	// msg marshals PatternFlowArpTargetHardwareAddr to protobuf object *otg.PatternFlowArpTargetHardwareAddr
	// and doesn't set defaults
	msg() *otg.PatternFlowArpTargetHardwareAddr
	// setMsg unmarshals PatternFlowArpTargetHardwareAddr from protobuf object *otg.PatternFlowArpTargetHardwareAddr
	// and doesn't set defaults
	setMsg(*otg.PatternFlowArpTargetHardwareAddr) PatternFlowArpTargetHardwareAddr
	// provides marshal interface
	Marshal() marshalPatternFlowArpTargetHardwareAddr
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowArpTargetHardwareAddr
	// validate validates PatternFlowArpTargetHardwareAddr
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowArpTargetHardwareAddr, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowArpTargetHardwareAddrChoiceEnum, set in PatternFlowArpTargetHardwareAddr
	Choice() PatternFlowArpTargetHardwareAddrChoiceEnum
	// setChoice assigns PatternFlowArpTargetHardwareAddrChoiceEnum provided by user to PatternFlowArpTargetHardwareAddr
	setChoice(value PatternFlowArpTargetHardwareAddrChoiceEnum) PatternFlowArpTargetHardwareAddr
	// HasChoice checks if Choice has been set in PatternFlowArpTargetHardwareAddr
	HasChoice() bool
	// Value returns string, set in PatternFlowArpTargetHardwareAddr.
	Value() string
	// SetValue assigns string provided by user to PatternFlowArpTargetHardwareAddr
	SetValue(value string) PatternFlowArpTargetHardwareAddr
	// HasValue checks if Value has been set in PatternFlowArpTargetHardwareAddr
	HasValue() bool
	// Values returns []string, set in PatternFlowArpTargetHardwareAddr.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowArpTargetHardwareAddr
	SetValues(value []string) PatternFlowArpTargetHardwareAddr
	// Increment returns PatternFlowArpTargetHardwareAddrCounter, set in PatternFlowArpTargetHardwareAddr.
	// PatternFlowArpTargetHardwareAddrCounter is mac counter pattern
	Increment() PatternFlowArpTargetHardwareAddrCounter
	// SetIncrement assigns PatternFlowArpTargetHardwareAddrCounter provided by user to PatternFlowArpTargetHardwareAddr.
	// PatternFlowArpTargetHardwareAddrCounter is mac counter pattern
	SetIncrement(value PatternFlowArpTargetHardwareAddrCounter) PatternFlowArpTargetHardwareAddr
	// HasIncrement checks if Increment has been set in PatternFlowArpTargetHardwareAddr
	HasIncrement() bool
	// Decrement returns PatternFlowArpTargetHardwareAddrCounter, set in PatternFlowArpTargetHardwareAddr.
	// PatternFlowArpTargetHardwareAddrCounter is mac counter pattern
	Decrement() PatternFlowArpTargetHardwareAddrCounter
	// SetDecrement assigns PatternFlowArpTargetHardwareAddrCounter provided by user to PatternFlowArpTargetHardwareAddr.
	// PatternFlowArpTargetHardwareAddrCounter is mac counter pattern
	SetDecrement(value PatternFlowArpTargetHardwareAddrCounter) PatternFlowArpTargetHardwareAddr
	// HasDecrement checks if Decrement has been set in PatternFlowArpTargetHardwareAddr
	HasDecrement() bool
	// MetricTags returns PatternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIterIter, set in PatternFlowArpTargetHardwareAddr
	MetricTags() PatternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter
	setNil()
}

type PatternFlowArpTargetHardwareAddrChoiceEnum string

// Enum of Choice on PatternFlowArpTargetHardwareAddr
var PatternFlowArpTargetHardwareAddrChoice = struct {
	VALUE     PatternFlowArpTargetHardwareAddrChoiceEnum
	VALUES    PatternFlowArpTargetHardwareAddrChoiceEnum
	INCREMENT PatternFlowArpTargetHardwareAddrChoiceEnum
	DECREMENT PatternFlowArpTargetHardwareAddrChoiceEnum
}{
	VALUE:     PatternFlowArpTargetHardwareAddrChoiceEnum("value"),
	VALUES:    PatternFlowArpTargetHardwareAddrChoiceEnum("values"),
	INCREMENT: PatternFlowArpTargetHardwareAddrChoiceEnum("increment"),
	DECREMENT: PatternFlowArpTargetHardwareAddrChoiceEnum("decrement"),
}

func (obj *patternFlowArpTargetHardwareAddr) Choice() PatternFlowArpTargetHardwareAddrChoiceEnum {
	return PatternFlowArpTargetHardwareAddrChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowArpTargetHardwareAddr) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowArpTargetHardwareAddr) setChoice(value PatternFlowArpTargetHardwareAddrChoiceEnum) PatternFlowArpTargetHardwareAddr {
	intValue, ok := otg.PatternFlowArpTargetHardwareAddr_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowArpTargetHardwareAddrChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowArpTargetHardwareAddr_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowArpTargetHardwareAddrChoice.VALUE {
		defaultValue := "00:00:00:00:00:00"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowArpTargetHardwareAddrChoice.VALUES {
		defaultValue := []string{"00:00:00:00:00:00"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowArpTargetHardwareAddrChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowArpTargetHardwareAddrCounter().msg()
	}

	if value == PatternFlowArpTargetHardwareAddrChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowArpTargetHardwareAddrCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowArpTargetHardwareAddr) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowArpTargetHardwareAddrChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowArpTargetHardwareAddr) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowArpTargetHardwareAddr object
func (obj *patternFlowArpTargetHardwareAddr) SetValue(value string) PatternFlowArpTargetHardwareAddr {
	obj.setChoice(PatternFlowArpTargetHardwareAddrChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowArpTargetHardwareAddr) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"00:00:00:00:00:00"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowArpTargetHardwareAddr object
func (obj *patternFlowArpTargetHardwareAddr) SetValues(value []string) PatternFlowArpTargetHardwareAddr {
	obj.setChoice(PatternFlowArpTargetHardwareAddrChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowArpTargetHardwareAddrCounter
func (obj *patternFlowArpTargetHardwareAddr) Increment() PatternFlowArpTargetHardwareAddrCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowArpTargetHardwareAddrChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowArpTargetHardwareAddrCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowArpTargetHardwareAddrCounter
func (obj *patternFlowArpTargetHardwareAddr) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowArpTargetHardwareAddrCounter value in the PatternFlowArpTargetHardwareAddr object
func (obj *patternFlowArpTargetHardwareAddr) SetIncrement(value PatternFlowArpTargetHardwareAddrCounter) PatternFlowArpTargetHardwareAddr {
	obj.setChoice(PatternFlowArpTargetHardwareAddrChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowArpTargetHardwareAddrCounter
func (obj *patternFlowArpTargetHardwareAddr) Decrement() PatternFlowArpTargetHardwareAddrCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowArpTargetHardwareAddrChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowArpTargetHardwareAddrCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowArpTargetHardwareAddrCounter
func (obj *patternFlowArpTargetHardwareAddr) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowArpTargetHardwareAddrCounter value in the PatternFlowArpTargetHardwareAddr object
func (obj *patternFlowArpTargetHardwareAddr) SetDecrement(value PatternFlowArpTargetHardwareAddrCounter) PatternFlowArpTargetHardwareAddr {
	obj.setChoice(PatternFlowArpTargetHardwareAddrChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowArpTargetHardwareAddrMetricTag
func (obj *patternFlowArpTargetHardwareAddr) MetricTags() PatternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowArpTargetHardwareAddrMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter struct {
	obj                                            *patternFlowArpTargetHardwareAddr
	patternFlowArpTargetHardwareAddrMetricTagSlice []PatternFlowArpTargetHardwareAddrMetricTag
	fieldPtr                                       *[]*otg.PatternFlowArpTargetHardwareAddrMetricTag
}

func newPatternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter(ptr *[]*otg.PatternFlowArpTargetHardwareAddrMetricTag) PatternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter {
	return &patternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter{fieldPtr: ptr}
}

type PatternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter interface {
	setMsg(*patternFlowArpTargetHardwareAddr) PatternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter
	Items() []PatternFlowArpTargetHardwareAddrMetricTag
	Add() PatternFlowArpTargetHardwareAddrMetricTag
	Append(items ...PatternFlowArpTargetHardwareAddrMetricTag) PatternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter
	Set(index int, newObj PatternFlowArpTargetHardwareAddrMetricTag) PatternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter
	Clear() PatternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter
	clearHolderSlice() PatternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter
	appendHolderSlice(item PatternFlowArpTargetHardwareAddrMetricTag) PatternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter
}

func (obj *patternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter) setMsg(msg *patternFlowArpTargetHardwareAddr) PatternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowArpTargetHardwareAddrMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter) Items() []PatternFlowArpTargetHardwareAddrMetricTag {
	return obj.patternFlowArpTargetHardwareAddrMetricTagSlice
}

func (obj *patternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter) Add() PatternFlowArpTargetHardwareAddrMetricTag {
	newObj := &otg.PatternFlowArpTargetHardwareAddrMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowArpTargetHardwareAddrMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowArpTargetHardwareAddrMetricTagSlice = append(obj.patternFlowArpTargetHardwareAddrMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter) Append(items ...PatternFlowArpTargetHardwareAddrMetricTag) PatternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowArpTargetHardwareAddrMetricTagSlice = append(obj.patternFlowArpTargetHardwareAddrMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter) Set(index int, newObj PatternFlowArpTargetHardwareAddrMetricTag) PatternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowArpTargetHardwareAddrMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter) Clear() PatternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowArpTargetHardwareAddrMetricTag{}
		obj.patternFlowArpTargetHardwareAddrMetricTagSlice = []PatternFlowArpTargetHardwareAddrMetricTag{}
	}
	return obj
}
func (obj *patternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter) clearHolderSlice() PatternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter {
	if len(obj.patternFlowArpTargetHardwareAddrMetricTagSlice) > 0 {
		obj.patternFlowArpTargetHardwareAddrMetricTagSlice = []PatternFlowArpTargetHardwareAddrMetricTag{}
	}
	return obj
}
func (obj *patternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter) appendHolderSlice(item PatternFlowArpTargetHardwareAddrMetricTag) PatternFlowArpTargetHardwareAddrPatternFlowArpTargetHardwareAddrMetricTagIter {
	obj.patternFlowArpTargetHardwareAddrMetricTagSlice = append(obj.patternFlowArpTargetHardwareAddrMetricTagSlice, item)
	return obj
}

func (obj *patternFlowArpTargetHardwareAddr) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpTargetHardwareAddr.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateMacSlice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowArpTargetHardwareAddr.Values"))
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
				obj.MetricTags().appendHolderSlice(&patternFlowArpTargetHardwareAddrMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowArpTargetHardwareAddr) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowArpTargetHardwareAddrChoice.VALUE)

	}

}
