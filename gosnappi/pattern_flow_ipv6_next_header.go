package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6NextHeader *****
type patternFlowIpv6NextHeader struct {
	validation
	obj              *otg.PatternFlowIpv6NextHeader
	marshaller       marshalPatternFlowIpv6NextHeader
	unMarshaller     unMarshalPatternFlowIpv6NextHeader
	incrementHolder  PatternFlowIpv6NextHeaderCounter
	decrementHolder  PatternFlowIpv6NextHeaderCounter
	metricTagsHolder PatternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter
}

func NewPatternFlowIpv6NextHeader() PatternFlowIpv6NextHeader {
	obj := patternFlowIpv6NextHeader{obj: &otg.PatternFlowIpv6NextHeader{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6NextHeader) msg() *otg.PatternFlowIpv6NextHeader {
	return obj.obj
}

func (obj *patternFlowIpv6NextHeader) setMsg(msg *otg.PatternFlowIpv6NextHeader) PatternFlowIpv6NextHeader {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6NextHeader struct {
	obj *patternFlowIpv6NextHeader
}

type marshalPatternFlowIpv6NextHeader interface {
	// ToProto marshals PatternFlowIpv6NextHeader to protobuf object *otg.PatternFlowIpv6NextHeader
	ToProto() (*otg.PatternFlowIpv6NextHeader, error)
	// ToPbText marshals PatternFlowIpv6NextHeader to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6NextHeader to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6NextHeader to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6NextHeader struct {
	obj *patternFlowIpv6NextHeader
}

type unMarshalPatternFlowIpv6NextHeader interface {
	// FromProto unmarshals PatternFlowIpv6NextHeader from protobuf object *otg.PatternFlowIpv6NextHeader
	FromProto(msg *otg.PatternFlowIpv6NextHeader) (PatternFlowIpv6NextHeader, error)
	// FromPbText unmarshals PatternFlowIpv6NextHeader from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6NextHeader from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6NextHeader from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6NextHeader) Marshal() marshalPatternFlowIpv6NextHeader {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6NextHeader{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6NextHeader) Unmarshal() unMarshalPatternFlowIpv6NextHeader {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6NextHeader{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6NextHeader) ToProto() (*otg.PatternFlowIpv6NextHeader, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6NextHeader) FromProto(msg *otg.PatternFlowIpv6NextHeader) (PatternFlowIpv6NextHeader, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6NextHeader) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6NextHeader) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6NextHeader) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6NextHeader) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6NextHeader) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6NextHeader) FromJson(value string) error {
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

func (obj *patternFlowIpv6NextHeader) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6NextHeader) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6NextHeader) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6NextHeader) Clone() (PatternFlowIpv6NextHeader, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6NextHeader()
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

func (obj *patternFlowIpv6NextHeader) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6NextHeader is next header
type PatternFlowIpv6NextHeader interface {
	Validation
	// msg marshals PatternFlowIpv6NextHeader to protobuf object *otg.PatternFlowIpv6NextHeader
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6NextHeader
	// setMsg unmarshals PatternFlowIpv6NextHeader from protobuf object *otg.PatternFlowIpv6NextHeader
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6NextHeader) PatternFlowIpv6NextHeader
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6NextHeader
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6NextHeader
	// validate validates PatternFlowIpv6NextHeader
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6NextHeader, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6NextHeaderChoiceEnum, set in PatternFlowIpv6NextHeader
	Choice() PatternFlowIpv6NextHeaderChoiceEnum
	// setChoice assigns PatternFlowIpv6NextHeaderChoiceEnum provided by user to PatternFlowIpv6NextHeader
	setChoice(value PatternFlowIpv6NextHeaderChoiceEnum) PatternFlowIpv6NextHeader
	// HasChoice checks if Choice has been set in PatternFlowIpv6NextHeader
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6NextHeader.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6NextHeader
	SetValue(value uint32) PatternFlowIpv6NextHeader
	// HasValue checks if Value has been set in PatternFlowIpv6NextHeader
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6NextHeader.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6NextHeader
	SetValues(value []uint32) PatternFlowIpv6NextHeader
	// Auto returns uint32, set in PatternFlowIpv6NextHeader.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIpv6NextHeader
	HasAuto() bool
	// Increment returns PatternFlowIpv6NextHeaderCounter, set in PatternFlowIpv6NextHeader.
	// PatternFlowIpv6NextHeaderCounter is integer counter pattern
	Increment() PatternFlowIpv6NextHeaderCounter
	// SetIncrement assigns PatternFlowIpv6NextHeaderCounter provided by user to PatternFlowIpv6NextHeader.
	// PatternFlowIpv6NextHeaderCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6NextHeaderCounter) PatternFlowIpv6NextHeader
	// HasIncrement checks if Increment has been set in PatternFlowIpv6NextHeader
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6NextHeaderCounter, set in PatternFlowIpv6NextHeader.
	// PatternFlowIpv6NextHeaderCounter is integer counter pattern
	Decrement() PatternFlowIpv6NextHeaderCounter
	// SetDecrement assigns PatternFlowIpv6NextHeaderCounter provided by user to PatternFlowIpv6NextHeader.
	// PatternFlowIpv6NextHeaderCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6NextHeaderCounter) PatternFlowIpv6NextHeader
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6NextHeader
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIterIter, set in PatternFlowIpv6NextHeader
	MetricTags() PatternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter
	setNil()
}

type PatternFlowIpv6NextHeaderChoiceEnum string

// Enum of Choice on PatternFlowIpv6NextHeader
var PatternFlowIpv6NextHeaderChoice = struct {
	VALUE     PatternFlowIpv6NextHeaderChoiceEnum
	VALUES    PatternFlowIpv6NextHeaderChoiceEnum
	AUTO      PatternFlowIpv6NextHeaderChoiceEnum
	INCREMENT PatternFlowIpv6NextHeaderChoiceEnum
	DECREMENT PatternFlowIpv6NextHeaderChoiceEnum
}{
	VALUE:     PatternFlowIpv6NextHeaderChoiceEnum("value"),
	VALUES:    PatternFlowIpv6NextHeaderChoiceEnum("values"),
	AUTO:      PatternFlowIpv6NextHeaderChoiceEnum("auto"),
	INCREMENT: PatternFlowIpv6NextHeaderChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6NextHeaderChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6NextHeader) Choice() PatternFlowIpv6NextHeaderChoiceEnum {
	return PatternFlowIpv6NextHeaderChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6NextHeader) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6NextHeader) setChoice(value PatternFlowIpv6NextHeaderChoiceEnum) PatternFlowIpv6NextHeader {
	intValue, ok := otg.PatternFlowIpv6NextHeader_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6NextHeaderChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6NextHeader_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6NextHeaderChoice.VALUE {
		defaultValue := uint32(59)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6NextHeaderChoice.VALUES {
		defaultValue := []uint32{59}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6NextHeaderChoice.AUTO {
		defaultValue := uint32(59)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIpv6NextHeaderChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6NextHeaderCounter().msg()
	}

	if value == PatternFlowIpv6NextHeaderChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6NextHeaderCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6NextHeader) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6NextHeaderChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6NextHeader) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6NextHeader object
func (obj *patternFlowIpv6NextHeader) SetValue(value uint32) PatternFlowIpv6NextHeader {
	obj.setChoice(PatternFlowIpv6NextHeaderChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6NextHeader) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{59})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6NextHeader object
func (obj *patternFlowIpv6NextHeader) SetValues(value []uint32) PatternFlowIpv6NextHeader {
	obj.setChoice(PatternFlowIpv6NextHeaderChoice.VALUES)
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
func (obj *patternFlowIpv6NextHeader) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIpv6NextHeaderChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIpv6NextHeader) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIpv6NextHeaderCounter
func (obj *patternFlowIpv6NextHeader) Increment() PatternFlowIpv6NextHeaderCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6NextHeaderChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6NextHeaderCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6NextHeaderCounter
func (obj *patternFlowIpv6NextHeader) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6NextHeaderCounter value in the PatternFlowIpv6NextHeader object
func (obj *patternFlowIpv6NextHeader) SetIncrement(value PatternFlowIpv6NextHeaderCounter) PatternFlowIpv6NextHeader {
	obj.setChoice(PatternFlowIpv6NextHeaderChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6NextHeaderCounter
func (obj *patternFlowIpv6NextHeader) Decrement() PatternFlowIpv6NextHeaderCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6NextHeaderChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6NextHeaderCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6NextHeaderCounter
func (obj *patternFlowIpv6NextHeader) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6NextHeaderCounter value in the PatternFlowIpv6NextHeader object
func (obj *patternFlowIpv6NextHeader) SetDecrement(value PatternFlowIpv6NextHeaderCounter) PatternFlowIpv6NextHeader {
	obj.setChoice(PatternFlowIpv6NextHeaderChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv6NextHeaderMetricTag
func (obj *patternFlowIpv6NextHeader) MetricTags() PatternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv6NextHeaderMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter struct {
	obj                                     *patternFlowIpv6NextHeader
	patternFlowIpv6NextHeaderMetricTagSlice []PatternFlowIpv6NextHeaderMetricTag
	fieldPtr                                *[]*otg.PatternFlowIpv6NextHeaderMetricTag
}

func newPatternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter(ptr *[]*otg.PatternFlowIpv6NextHeaderMetricTag) PatternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter {
	return &patternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter interface {
	setMsg(*patternFlowIpv6NextHeader) PatternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter
	Items() []PatternFlowIpv6NextHeaderMetricTag
	Add() PatternFlowIpv6NextHeaderMetricTag
	Append(items ...PatternFlowIpv6NextHeaderMetricTag) PatternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter
	Set(index int, newObj PatternFlowIpv6NextHeaderMetricTag) PatternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter
	Clear() PatternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter
	clearHolderSlice() PatternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter
	appendHolderSlice(item PatternFlowIpv6NextHeaderMetricTag) PatternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter
}

func (obj *patternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter) setMsg(msg *patternFlowIpv6NextHeader) PatternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv6NextHeaderMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter) Items() []PatternFlowIpv6NextHeaderMetricTag {
	return obj.patternFlowIpv6NextHeaderMetricTagSlice
}

func (obj *patternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter) Add() PatternFlowIpv6NextHeaderMetricTag {
	newObj := &otg.PatternFlowIpv6NextHeaderMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv6NextHeaderMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv6NextHeaderMetricTagSlice = append(obj.patternFlowIpv6NextHeaderMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter) Append(items ...PatternFlowIpv6NextHeaderMetricTag) PatternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv6NextHeaderMetricTagSlice = append(obj.patternFlowIpv6NextHeaderMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter) Set(index int, newObj PatternFlowIpv6NextHeaderMetricTag) PatternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv6NextHeaderMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter) Clear() PatternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv6NextHeaderMetricTag{}
		obj.patternFlowIpv6NextHeaderMetricTagSlice = []PatternFlowIpv6NextHeaderMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter) clearHolderSlice() PatternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter {
	if len(obj.patternFlowIpv6NextHeaderMetricTagSlice) > 0 {
		obj.patternFlowIpv6NextHeaderMetricTagSlice = []PatternFlowIpv6NextHeaderMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter) appendHolderSlice(item PatternFlowIpv6NextHeaderMetricTag) PatternFlowIpv6NextHeaderPatternFlowIpv6NextHeaderMetricTagIter {
	obj.patternFlowIpv6NextHeaderMetricTagSlice = append(obj.patternFlowIpv6NextHeaderMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIpv6NextHeader) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6NextHeader.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6NextHeader.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6NextHeader.Auto <= 255 but Got %d", *obj.obj.Auto))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv6NextHeaderMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIpv6NextHeader) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowIpv6NextHeaderChoice.AUTO)

	}

}
