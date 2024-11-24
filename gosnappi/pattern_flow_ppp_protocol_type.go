package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPppProtocolType *****
type patternFlowPppProtocolType struct {
	validation
	obj              *otg.PatternFlowPppProtocolType
	marshaller       marshalPatternFlowPppProtocolType
	unMarshaller     unMarshalPatternFlowPppProtocolType
	incrementHolder  PatternFlowPppProtocolTypeCounter
	decrementHolder  PatternFlowPppProtocolTypeCounter
	metricTagsHolder PatternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter
}

func NewPatternFlowPppProtocolType() PatternFlowPppProtocolType {
	obj := patternFlowPppProtocolType{obj: &otg.PatternFlowPppProtocolType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPppProtocolType) msg() *otg.PatternFlowPppProtocolType {
	return obj.obj
}

func (obj *patternFlowPppProtocolType) setMsg(msg *otg.PatternFlowPppProtocolType) PatternFlowPppProtocolType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPppProtocolType struct {
	obj *patternFlowPppProtocolType
}

type marshalPatternFlowPppProtocolType interface {
	// ToProto marshals PatternFlowPppProtocolType to protobuf object *otg.PatternFlowPppProtocolType
	ToProto() (*otg.PatternFlowPppProtocolType, error)
	// ToPbText marshals PatternFlowPppProtocolType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPppProtocolType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPppProtocolType to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowPppProtocolType to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowPppProtocolType struct {
	obj *patternFlowPppProtocolType
}

type unMarshalPatternFlowPppProtocolType interface {
	// FromProto unmarshals PatternFlowPppProtocolType from protobuf object *otg.PatternFlowPppProtocolType
	FromProto(msg *otg.PatternFlowPppProtocolType) (PatternFlowPppProtocolType, error)
	// FromPbText unmarshals PatternFlowPppProtocolType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPppProtocolType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPppProtocolType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPppProtocolType) Marshal() marshalPatternFlowPppProtocolType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPppProtocolType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPppProtocolType) Unmarshal() unMarshalPatternFlowPppProtocolType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPppProtocolType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPppProtocolType) ToProto() (*otg.PatternFlowPppProtocolType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPppProtocolType) FromProto(msg *otg.PatternFlowPppProtocolType) (PatternFlowPppProtocolType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPppProtocolType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPppProtocolType) FromPbText(value string) error {
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

func (m *marshalpatternFlowPppProtocolType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPppProtocolType) FromYaml(value string) error {
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

func (m *marshalpatternFlowPppProtocolType) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowPppProtocolType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPppProtocolType) FromJson(value string) error {
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

func (obj *patternFlowPppProtocolType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPppProtocolType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPppProtocolType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPppProtocolType) Clone() (PatternFlowPppProtocolType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPppProtocolType()
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

func (obj *patternFlowPppProtocolType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowPppProtocolType is pPP protocol type
type PatternFlowPppProtocolType interface {
	Validation
	// msg marshals PatternFlowPppProtocolType to protobuf object *otg.PatternFlowPppProtocolType
	// and doesn't set defaults
	msg() *otg.PatternFlowPppProtocolType
	// setMsg unmarshals PatternFlowPppProtocolType from protobuf object *otg.PatternFlowPppProtocolType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPppProtocolType) PatternFlowPppProtocolType
	// provides marshal interface
	Marshal() marshalPatternFlowPppProtocolType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPppProtocolType
	// validate validates PatternFlowPppProtocolType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPppProtocolType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowPppProtocolTypeChoiceEnum, set in PatternFlowPppProtocolType
	Choice() PatternFlowPppProtocolTypeChoiceEnum
	// setChoice assigns PatternFlowPppProtocolTypeChoiceEnum provided by user to PatternFlowPppProtocolType
	setChoice(value PatternFlowPppProtocolTypeChoiceEnum) PatternFlowPppProtocolType
	// HasChoice checks if Choice has been set in PatternFlowPppProtocolType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowPppProtocolType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowPppProtocolType
	SetValue(value uint32) PatternFlowPppProtocolType
	// HasValue checks if Value has been set in PatternFlowPppProtocolType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowPppProtocolType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowPppProtocolType
	SetValues(value []uint32) PatternFlowPppProtocolType
	// Auto returns uint32, set in PatternFlowPppProtocolType.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowPppProtocolType
	HasAuto() bool
	// Increment returns PatternFlowPppProtocolTypeCounter, set in PatternFlowPppProtocolType.
	// PatternFlowPppProtocolTypeCounter is integer counter pattern
	Increment() PatternFlowPppProtocolTypeCounter
	// SetIncrement assigns PatternFlowPppProtocolTypeCounter provided by user to PatternFlowPppProtocolType.
	// PatternFlowPppProtocolTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowPppProtocolTypeCounter) PatternFlowPppProtocolType
	// HasIncrement checks if Increment has been set in PatternFlowPppProtocolType
	HasIncrement() bool
	// Decrement returns PatternFlowPppProtocolTypeCounter, set in PatternFlowPppProtocolType.
	// PatternFlowPppProtocolTypeCounter is integer counter pattern
	Decrement() PatternFlowPppProtocolTypeCounter
	// SetDecrement assigns PatternFlowPppProtocolTypeCounter provided by user to PatternFlowPppProtocolType.
	// PatternFlowPppProtocolTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowPppProtocolTypeCounter) PatternFlowPppProtocolType
	// HasDecrement checks if Decrement has been set in PatternFlowPppProtocolType
	HasDecrement() bool
	// MetricTags returns PatternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIterIter, set in PatternFlowPppProtocolType
	MetricTags() PatternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter
	setNil()
}

type PatternFlowPppProtocolTypeChoiceEnum string

// Enum of Choice on PatternFlowPppProtocolType
var PatternFlowPppProtocolTypeChoice = struct {
	VALUE     PatternFlowPppProtocolTypeChoiceEnum
	VALUES    PatternFlowPppProtocolTypeChoiceEnum
	AUTO      PatternFlowPppProtocolTypeChoiceEnum
	INCREMENT PatternFlowPppProtocolTypeChoiceEnum
	DECREMENT PatternFlowPppProtocolTypeChoiceEnum
}{
	VALUE:     PatternFlowPppProtocolTypeChoiceEnum("value"),
	VALUES:    PatternFlowPppProtocolTypeChoiceEnum("values"),
	AUTO:      PatternFlowPppProtocolTypeChoiceEnum("auto"),
	INCREMENT: PatternFlowPppProtocolTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowPppProtocolTypeChoiceEnum("decrement"),
}

func (obj *patternFlowPppProtocolType) Choice() PatternFlowPppProtocolTypeChoiceEnum {
	return PatternFlowPppProtocolTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowPppProtocolType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowPppProtocolType) setChoice(value PatternFlowPppProtocolTypeChoiceEnum) PatternFlowPppProtocolType {
	intValue, ok := otg.PatternFlowPppProtocolType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowPppProtocolTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowPppProtocolType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowPppProtocolTypeChoice.VALUE {
		defaultValue := uint32(33)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowPppProtocolTypeChoice.VALUES {
		defaultValue := []uint32{33}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowPppProtocolTypeChoice.AUTO {
		defaultValue := uint32(33)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowPppProtocolTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowPppProtocolTypeCounter().msg()
	}

	if value == PatternFlowPppProtocolTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowPppProtocolTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowPppProtocolType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowPppProtocolTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowPppProtocolType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowPppProtocolType object
func (obj *patternFlowPppProtocolType) SetValue(value uint32) PatternFlowPppProtocolType {
	obj.setChoice(PatternFlowPppProtocolTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowPppProtocolType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{33})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowPppProtocolType object
func (obj *patternFlowPppProtocolType) SetValues(value []uint32) PatternFlowPppProtocolType {
	obj.setChoice(PatternFlowPppProtocolTypeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowPppProtocolType) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowPppProtocolTypeChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowPppProtocolType) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowPppProtocolTypeCounter
func (obj *patternFlowPppProtocolType) Increment() PatternFlowPppProtocolTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowPppProtocolTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowPppProtocolTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowPppProtocolTypeCounter
func (obj *patternFlowPppProtocolType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowPppProtocolTypeCounter value in the PatternFlowPppProtocolType object
func (obj *patternFlowPppProtocolType) SetIncrement(value PatternFlowPppProtocolTypeCounter) PatternFlowPppProtocolType {
	obj.setChoice(PatternFlowPppProtocolTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowPppProtocolTypeCounter
func (obj *patternFlowPppProtocolType) Decrement() PatternFlowPppProtocolTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowPppProtocolTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowPppProtocolTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowPppProtocolTypeCounter
func (obj *patternFlowPppProtocolType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowPppProtocolTypeCounter value in the PatternFlowPppProtocolType object
func (obj *patternFlowPppProtocolType) SetDecrement(value PatternFlowPppProtocolTypeCounter) PatternFlowPppProtocolType {
	obj.setChoice(PatternFlowPppProtocolTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowPppProtocolTypeMetricTag
func (obj *patternFlowPppProtocolType) MetricTags() PatternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowPppProtocolTypeMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter struct {
	obj                                      *patternFlowPppProtocolType
	patternFlowPppProtocolTypeMetricTagSlice []PatternFlowPppProtocolTypeMetricTag
	fieldPtr                                 *[]*otg.PatternFlowPppProtocolTypeMetricTag
}

func newPatternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter(ptr *[]*otg.PatternFlowPppProtocolTypeMetricTag) PatternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter {
	return &patternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter{fieldPtr: ptr}
}

type PatternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter interface {
	setMsg(*patternFlowPppProtocolType) PatternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter
	Items() []PatternFlowPppProtocolTypeMetricTag
	Add() PatternFlowPppProtocolTypeMetricTag
	Append(items ...PatternFlowPppProtocolTypeMetricTag) PatternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter
	Set(index int, newObj PatternFlowPppProtocolTypeMetricTag) PatternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter
	Clear() PatternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter
	clearHolderSlice() PatternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter
	appendHolderSlice(item PatternFlowPppProtocolTypeMetricTag) PatternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter
}

func (obj *patternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter) setMsg(msg *patternFlowPppProtocolType) PatternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowPppProtocolTypeMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter) Items() []PatternFlowPppProtocolTypeMetricTag {
	return obj.patternFlowPppProtocolTypeMetricTagSlice
}

func (obj *patternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter) Add() PatternFlowPppProtocolTypeMetricTag {
	newObj := &otg.PatternFlowPppProtocolTypeMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowPppProtocolTypeMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowPppProtocolTypeMetricTagSlice = append(obj.patternFlowPppProtocolTypeMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter) Append(items ...PatternFlowPppProtocolTypeMetricTag) PatternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowPppProtocolTypeMetricTagSlice = append(obj.patternFlowPppProtocolTypeMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter) Set(index int, newObj PatternFlowPppProtocolTypeMetricTag) PatternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowPppProtocolTypeMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter) Clear() PatternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowPppProtocolTypeMetricTag{}
		obj.patternFlowPppProtocolTypeMetricTagSlice = []PatternFlowPppProtocolTypeMetricTag{}
	}
	return obj
}
func (obj *patternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter) clearHolderSlice() PatternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter {
	if len(obj.patternFlowPppProtocolTypeMetricTagSlice) > 0 {
		obj.patternFlowPppProtocolTypeMetricTagSlice = []PatternFlowPppProtocolTypeMetricTag{}
	}
	return obj
}
func (obj *patternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter) appendHolderSlice(item PatternFlowPppProtocolTypeMetricTag) PatternFlowPppProtocolTypePatternFlowPppProtocolTypeMetricTagIter {
	obj.patternFlowPppProtocolTypeMetricTagSlice = append(obj.patternFlowPppProtocolTypeMetricTagSlice, item)
	return obj
}

func (obj *patternFlowPppProtocolType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPppProtocolType.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowPppProtocolType.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPppProtocolType.Auto <= 65535 but Got %d", *obj.obj.Auto))
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
				obj.MetricTags().appendHolderSlice(&patternFlowPppProtocolTypeMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowPppProtocolType) setDefault() {
	var choices_set int = 0
	var choice PatternFlowPppProtocolTypeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowPppProtocolTypeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowPppProtocolTypeChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowPppProtocolTypeChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowPppProtocolTypeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowPppProtocolTypeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowPppProtocolTypeChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowPppProtocolType")
			}
		} else {
			intVal := otg.PatternFlowPppProtocolType_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowPppProtocolType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
