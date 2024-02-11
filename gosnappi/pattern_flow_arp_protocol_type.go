package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowArpProtocolType *****
type patternFlowArpProtocolType struct {
	validation
	obj              *otg.PatternFlowArpProtocolType
	marshaller       marshalPatternFlowArpProtocolType
	unMarshaller     unMarshalPatternFlowArpProtocolType
	incrementHolder  PatternFlowArpProtocolTypeCounter
	decrementHolder  PatternFlowArpProtocolTypeCounter
	metricTagsHolder PatternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter
}

func NewPatternFlowArpProtocolType() PatternFlowArpProtocolType {
	obj := patternFlowArpProtocolType{obj: &otg.PatternFlowArpProtocolType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowArpProtocolType) msg() *otg.PatternFlowArpProtocolType {
	return obj.obj
}

func (obj *patternFlowArpProtocolType) setMsg(msg *otg.PatternFlowArpProtocolType) PatternFlowArpProtocolType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowArpProtocolType struct {
	obj *patternFlowArpProtocolType
}

type marshalPatternFlowArpProtocolType interface {
	// ToProto marshals PatternFlowArpProtocolType to protobuf object *otg.PatternFlowArpProtocolType
	ToProto() (*otg.PatternFlowArpProtocolType, error)
	// ToPbText marshals PatternFlowArpProtocolType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowArpProtocolType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowArpProtocolType to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowArpProtocolType struct {
	obj *patternFlowArpProtocolType
}

type unMarshalPatternFlowArpProtocolType interface {
	// FromProto unmarshals PatternFlowArpProtocolType from protobuf object *otg.PatternFlowArpProtocolType
	FromProto(msg *otg.PatternFlowArpProtocolType) (PatternFlowArpProtocolType, error)
	// FromPbText unmarshals PatternFlowArpProtocolType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowArpProtocolType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowArpProtocolType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowArpProtocolType) Marshal() marshalPatternFlowArpProtocolType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowArpProtocolType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowArpProtocolType) Unmarshal() unMarshalPatternFlowArpProtocolType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowArpProtocolType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowArpProtocolType) ToProto() (*otg.PatternFlowArpProtocolType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowArpProtocolType) FromProto(msg *otg.PatternFlowArpProtocolType) (PatternFlowArpProtocolType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowArpProtocolType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowArpProtocolType) FromPbText(value string) error {
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

func (m *marshalpatternFlowArpProtocolType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowArpProtocolType) FromYaml(value string) error {
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

func (m *marshalpatternFlowArpProtocolType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowArpProtocolType) FromJson(value string) error {
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

func (obj *patternFlowArpProtocolType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowArpProtocolType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowArpProtocolType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowArpProtocolType) Clone() (PatternFlowArpProtocolType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowArpProtocolType()
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

func (obj *patternFlowArpProtocolType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowArpProtocolType is the internetwork protocol for which the ARP request is intended
type PatternFlowArpProtocolType interface {
	Validation
	// msg marshals PatternFlowArpProtocolType to protobuf object *otg.PatternFlowArpProtocolType
	// and doesn't set defaults
	msg() *otg.PatternFlowArpProtocolType
	// setMsg unmarshals PatternFlowArpProtocolType from protobuf object *otg.PatternFlowArpProtocolType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowArpProtocolType) PatternFlowArpProtocolType
	// provides marshal interface
	Marshal() marshalPatternFlowArpProtocolType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowArpProtocolType
	// validate validates PatternFlowArpProtocolType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowArpProtocolType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowArpProtocolTypeChoiceEnum, set in PatternFlowArpProtocolType
	Choice() PatternFlowArpProtocolTypeChoiceEnum
	// setChoice assigns PatternFlowArpProtocolTypeChoiceEnum provided by user to PatternFlowArpProtocolType
	setChoice(value PatternFlowArpProtocolTypeChoiceEnum) PatternFlowArpProtocolType
	// HasChoice checks if Choice has been set in PatternFlowArpProtocolType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowArpProtocolType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowArpProtocolType
	SetValue(value uint32) PatternFlowArpProtocolType
	// HasValue checks if Value has been set in PatternFlowArpProtocolType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowArpProtocolType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowArpProtocolType
	SetValues(value []uint32) PatternFlowArpProtocolType
	// Increment returns PatternFlowArpProtocolTypeCounter, set in PatternFlowArpProtocolType.
	// PatternFlowArpProtocolTypeCounter is integer counter pattern
	Increment() PatternFlowArpProtocolTypeCounter
	// SetIncrement assigns PatternFlowArpProtocolTypeCounter provided by user to PatternFlowArpProtocolType.
	// PatternFlowArpProtocolTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowArpProtocolTypeCounter) PatternFlowArpProtocolType
	// HasIncrement checks if Increment has been set in PatternFlowArpProtocolType
	HasIncrement() bool
	// Decrement returns PatternFlowArpProtocolTypeCounter, set in PatternFlowArpProtocolType.
	// PatternFlowArpProtocolTypeCounter is integer counter pattern
	Decrement() PatternFlowArpProtocolTypeCounter
	// SetDecrement assigns PatternFlowArpProtocolTypeCounter provided by user to PatternFlowArpProtocolType.
	// PatternFlowArpProtocolTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowArpProtocolTypeCounter) PatternFlowArpProtocolType
	// HasDecrement checks if Decrement has been set in PatternFlowArpProtocolType
	HasDecrement() bool
	// MetricTags returns PatternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIterIter, set in PatternFlowArpProtocolType
	MetricTags() PatternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter
	setNil()
}

type PatternFlowArpProtocolTypeChoiceEnum string

// Enum of Choice on PatternFlowArpProtocolType
var PatternFlowArpProtocolTypeChoice = struct {
	VALUE     PatternFlowArpProtocolTypeChoiceEnum
	VALUES    PatternFlowArpProtocolTypeChoiceEnum
	INCREMENT PatternFlowArpProtocolTypeChoiceEnum
	DECREMENT PatternFlowArpProtocolTypeChoiceEnum
}{
	VALUE:     PatternFlowArpProtocolTypeChoiceEnum("value"),
	VALUES:    PatternFlowArpProtocolTypeChoiceEnum("values"),
	INCREMENT: PatternFlowArpProtocolTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowArpProtocolTypeChoiceEnum("decrement"),
}

func (obj *patternFlowArpProtocolType) Choice() PatternFlowArpProtocolTypeChoiceEnum {
	return PatternFlowArpProtocolTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowArpProtocolType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowArpProtocolType) setChoice(value PatternFlowArpProtocolTypeChoiceEnum) PatternFlowArpProtocolType {
	intValue, ok := otg.PatternFlowArpProtocolType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowArpProtocolTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowArpProtocolType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowArpProtocolTypeChoice.VALUE {
		defaultValue := uint32(2048)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowArpProtocolTypeChoice.VALUES {
		defaultValue := []uint32{2048}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowArpProtocolTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowArpProtocolTypeCounter().msg()
	}

	if value == PatternFlowArpProtocolTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowArpProtocolTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowArpProtocolType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowArpProtocolTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowArpProtocolType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowArpProtocolType object
func (obj *patternFlowArpProtocolType) SetValue(value uint32) PatternFlowArpProtocolType {
	obj.setChoice(PatternFlowArpProtocolTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowArpProtocolType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{2048})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowArpProtocolType object
func (obj *patternFlowArpProtocolType) SetValues(value []uint32) PatternFlowArpProtocolType {
	obj.setChoice(PatternFlowArpProtocolTypeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowArpProtocolTypeCounter
func (obj *patternFlowArpProtocolType) Increment() PatternFlowArpProtocolTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowArpProtocolTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowArpProtocolTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowArpProtocolTypeCounter
func (obj *patternFlowArpProtocolType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowArpProtocolTypeCounter value in the PatternFlowArpProtocolType object
func (obj *patternFlowArpProtocolType) SetIncrement(value PatternFlowArpProtocolTypeCounter) PatternFlowArpProtocolType {
	obj.setChoice(PatternFlowArpProtocolTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowArpProtocolTypeCounter
func (obj *patternFlowArpProtocolType) Decrement() PatternFlowArpProtocolTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowArpProtocolTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowArpProtocolTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowArpProtocolTypeCounter
func (obj *patternFlowArpProtocolType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowArpProtocolTypeCounter value in the PatternFlowArpProtocolType object
func (obj *patternFlowArpProtocolType) SetDecrement(value PatternFlowArpProtocolTypeCounter) PatternFlowArpProtocolType {
	obj.setChoice(PatternFlowArpProtocolTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowArpProtocolTypeMetricTag
func (obj *patternFlowArpProtocolType) MetricTags() PatternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowArpProtocolTypeMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter struct {
	obj                                      *patternFlowArpProtocolType
	patternFlowArpProtocolTypeMetricTagSlice []PatternFlowArpProtocolTypeMetricTag
	fieldPtr                                 *[]*otg.PatternFlowArpProtocolTypeMetricTag
}

func newPatternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter(ptr *[]*otg.PatternFlowArpProtocolTypeMetricTag) PatternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter {
	return &patternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter{fieldPtr: ptr}
}

type PatternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter interface {
	setMsg(*patternFlowArpProtocolType) PatternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter
	Items() []PatternFlowArpProtocolTypeMetricTag
	Add() PatternFlowArpProtocolTypeMetricTag
	Append(items ...PatternFlowArpProtocolTypeMetricTag) PatternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter
	Set(index int, newObj PatternFlowArpProtocolTypeMetricTag) PatternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter
	Clear() PatternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter
	clearHolderSlice() PatternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter
	appendHolderSlice(item PatternFlowArpProtocolTypeMetricTag) PatternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter
}

func (obj *patternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter) setMsg(msg *patternFlowArpProtocolType) PatternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowArpProtocolTypeMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter) Items() []PatternFlowArpProtocolTypeMetricTag {
	return obj.patternFlowArpProtocolTypeMetricTagSlice
}

func (obj *patternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter) Add() PatternFlowArpProtocolTypeMetricTag {
	newObj := &otg.PatternFlowArpProtocolTypeMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowArpProtocolTypeMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowArpProtocolTypeMetricTagSlice = append(obj.patternFlowArpProtocolTypeMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter) Append(items ...PatternFlowArpProtocolTypeMetricTag) PatternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowArpProtocolTypeMetricTagSlice = append(obj.patternFlowArpProtocolTypeMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter) Set(index int, newObj PatternFlowArpProtocolTypeMetricTag) PatternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowArpProtocolTypeMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter) Clear() PatternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowArpProtocolTypeMetricTag{}
		obj.patternFlowArpProtocolTypeMetricTagSlice = []PatternFlowArpProtocolTypeMetricTag{}
	}
	return obj
}
func (obj *patternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter) clearHolderSlice() PatternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter {
	if len(obj.patternFlowArpProtocolTypeMetricTagSlice) > 0 {
		obj.patternFlowArpProtocolTypeMetricTagSlice = []PatternFlowArpProtocolTypeMetricTag{}
	}
	return obj
}
func (obj *patternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter) appendHolderSlice(item PatternFlowArpProtocolTypeMetricTag) PatternFlowArpProtocolTypePatternFlowArpProtocolTypeMetricTagIter {
	obj.patternFlowArpProtocolTypeMetricTagSlice = append(obj.patternFlowArpProtocolTypeMetricTagSlice, item)
	return obj
}

func (obj *patternFlowArpProtocolType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpProtocolType.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowArpProtocolType.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowArpProtocolTypeMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowArpProtocolType) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowArpProtocolTypeChoice.VALUE)

	}

}
