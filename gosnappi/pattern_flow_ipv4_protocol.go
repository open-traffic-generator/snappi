package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4Protocol *****
type patternFlowIpv4Protocol struct {
	validation
	obj              *otg.PatternFlowIpv4Protocol
	marshaller       marshalPatternFlowIpv4Protocol
	unMarshaller     unMarshalPatternFlowIpv4Protocol
	incrementHolder  PatternFlowIpv4ProtocolCounter
	decrementHolder  PatternFlowIpv4ProtocolCounter
	metricTagsHolder PatternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter
}

func NewPatternFlowIpv4Protocol() PatternFlowIpv4Protocol {
	obj := patternFlowIpv4Protocol{obj: &otg.PatternFlowIpv4Protocol{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4Protocol) msg() *otg.PatternFlowIpv4Protocol {
	return obj.obj
}

func (obj *patternFlowIpv4Protocol) setMsg(msg *otg.PatternFlowIpv4Protocol) PatternFlowIpv4Protocol {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4Protocol struct {
	obj *patternFlowIpv4Protocol
}

type marshalPatternFlowIpv4Protocol interface {
	// ToProto marshals PatternFlowIpv4Protocol to protobuf object *otg.PatternFlowIpv4Protocol
	ToProto() (*otg.PatternFlowIpv4Protocol, error)
	// ToPbText marshals PatternFlowIpv4Protocol to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4Protocol to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4Protocol to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4Protocol struct {
	obj *patternFlowIpv4Protocol
}

type unMarshalPatternFlowIpv4Protocol interface {
	// FromProto unmarshals PatternFlowIpv4Protocol from protobuf object *otg.PatternFlowIpv4Protocol
	FromProto(msg *otg.PatternFlowIpv4Protocol) (PatternFlowIpv4Protocol, error)
	// FromPbText unmarshals PatternFlowIpv4Protocol from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4Protocol from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4Protocol from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4Protocol) Marshal() marshalPatternFlowIpv4Protocol {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4Protocol{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4Protocol) Unmarshal() unMarshalPatternFlowIpv4Protocol {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4Protocol{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4Protocol) ToProto() (*otg.PatternFlowIpv4Protocol, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4Protocol) FromProto(msg *otg.PatternFlowIpv4Protocol) (PatternFlowIpv4Protocol, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4Protocol) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4Protocol) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4Protocol) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4Protocol) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4Protocol) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4Protocol) FromJson(value string) error {
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

func (obj *patternFlowIpv4Protocol) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4Protocol) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4Protocol) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4Protocol) Clone() (PatternFlowIpv4Protocol, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4Protocol()
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

func (obj *patternFlowIpv4Protocol) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4Protocol is protocol, default is 61 any host internal protocol
type PatternFlowIpv4Protocol interface {
	Validation
	// msg marshals PatternFlowIpv4Protocol to protobuf object *otg.PatternFlowIpv4Protocol
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4Protocol
	// setMsg unmarshals PatternFlowIpv4Protocol from protobuf object *otg.PatternFlowIpv4Protocol
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4Protocol) PatternFlowIpv4Protocol
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4Protocol
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4Protocol
	// validate validates PatternFlowIpv4Protocol
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4Protocol, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4ProtocolChoiceEnum, set in PatternFlowIpv4Protocol
	Choice() PatternFlowIpv4ProtocolChoiceEnum
	// setChoice assigns PatternFlowIpv4ProtocolChoiceEnum provided by user to PatternFlowIpv4Protocol
	setChoice(value PatternFlowIpv4ProtocolChoiceEnum) PatternFlowIpv4Protocol
	// HasChoice checks if Choice has been set in PatternFlowIpv4Protocol
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv4Protocol.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv4Protocol
	SetValue(value uint32) PatternFlowIpv4Protocol
	// HasValue checks if Value has been set in PatternFlowIpv4Protocol
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv4Protocol.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv4Protocol
	SetValues(value []uint32) PatternFlowIpv4Protocol
	// Auto returns uint32, set in PatternFlowIpv4Protocol.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIpv4Protocol
	HasAuto() bool
	// Increment returns PatternFlowIpv4ProtocolCounter, set in PatternFlowIpv4Protocol.
	// PatternFlowIpv4ProtocolCounter is integer counter pattern
	Increment() PatternFlowIpv4ProtocolCounter
	// SetIncrement assigns PatternFlowIpv4ProtocolCounter provided by user to PatternFlowIpv4Protocol.
	// PatternFlowIpv4ProtocolCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv4ProtocolCounter) PatternFlowIpv4Protocol
	// HasIncrement checks if Increment has been set in PatternFlowIpv4Protocol
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4ProtocolCounter, set in PatternFlowIpv4Protocol.
	// PatternFlowIpv4ProtocolCounter is integer counter pattern
	Decrement() PatternFlowIpv4ProtocolCounter
	// SetDecrement assigns PatternFlowIpv4ProtocolCounter provided by user to PatternFlowIpv4Protocol.
	// PatternFlowIpv4ProtocolCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv4ProtocolCounter) PatternFlowIpv4Protocol
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4Protocol
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIterIter, set in PatternFlowIpv4Protocol
	MetricTags() PatternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter
	setNil()
}

type PatternFlowIpv4ProtocolChoiceEnum string

// Enum of Choice on PatternFlowIpv4Protocol
var PatternFlowIpv4ProtocolChoice = struct {
	VALUE     PatternFlowIpv4ProtocolChoiceEnum
	VALUES    PatternFlowIpv4ProtocolChoiceEnum
	AUTO      PatternFlowIpv4ProtocolChoiceEnum
	INCREMENT PatternFlowIpv4ProtocolChoiceEnum
	DECREMENT PatternFlowIpv4ProtocolChoiceEnum
}{
	VALUE:     PatternFlowIpv4ProtocolChoiceEnum("value"),
	VALUES:    PatternFlowIpv4ProtocolChoiceEnum("values"),
	AUTO:      PatternFlowIpv4ProtocolChoiceEnum("auto"),
	INCREMENT: PatternFlowIpv4ProtocolChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4ProtocolChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4Protocol) Choice() PatternFlowIpv4ProtocolChoiceEnum {
	return PatternFlowIpv4ProtocolChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4Protocol) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4Protocol) setChoice(value PatternFlowIpv4ProtocolChoiceEnum) PatternFlowIpv4Protocol {
	intValue, ok := otg.PatternFlowIpv4Protocol_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4ProtocolChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4Protocol_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4ProtocolChoice.VALUE {
		defaultValue := uint32(61)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4ProtocolChoice.VALUES {
		defaultValue := []uint32{61}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4ProtocolChoice.AUTO {
		defaultValue := uint32(61)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIpv4ProtocolChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4ProtocolCounter().msg()
	}

	if value == PatternFlowIpv4ProtocolChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4ProtocolCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4Protocol) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4ProtocolChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4Protocol) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv4Protocol object
func (obj *patternFlowIpv4Protocol) SetValue(value uint32) PatternFlowIpv4Protocol {
	obj.setChoice(PatternFlowIpv4ProtocolChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv4Protocol) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{61})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv4Protocol object
func (obj *patternFlowIpv4Protocol) SetValues(value []uint32) PatternFlowIpv4Protocol {
	obj.setChoice(PatternFlowIpv4ProtocolChoice.VALUES)
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
func (obj *patternFlowIpv4Protocol) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIpv4ProtocolChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIpv4Protocol) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIpv4ProtocolCounter
func (obj *patternFlowIpv4Protocol) Increment() PatternFlowIpv4ProtocolCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4ProtocolChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4ProtocolCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4ProtocolCounter
func (obj *patternFlowIpv4Protocol) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4ProtocolCounter value in the PatternFlowIpv4Protocol object
func (obj *patternFlowIpv4Protocol) SetIncrement(value PatternFlowIpv4ProtocolCounter) PatternFlowIpv4Protocol {
	obj.setChoice(PatternFlowIpv4ProtocolChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4ProtocolCounter
func (obj *patternFlowIpv4Protocol) Decrement() PatternFlowIpv4ProtocolCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4ProtocolChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4ProtocolCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4ProtocolCounter
func (obj *patternFlowIpv4Protocol) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4ProtocolCounter value in the PatternFlowIpv4Protocol object
func (obj *patternFlowIpv4Protocol) SetDecrement(value PatternFlowIpv4ProtocolCounter) PatternFlowIpv4Protocol {
	obj.setChoice(PatternFlowIpv4ProtocolChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv4ProtocolMetricTag
func (obj *patternFlowIpv4Protocol) MetricTags() PatternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv4ProtocolMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter struct {
	obj                                   *patternFlowIpv4Protocol
	patternFlowIpv4ProtocolMetricTagSlice []PatternFlowIpv4ProtocolMetricTag
	fieldPtr                              *[]*otg.PatternFlowIpv4ProtocolMetricTag
}

func newPatternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter(ptr *[]*otg.PatternFlowIpv4ProtocolMetricTag) PatternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter {
	return &patternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter interface {
	setMsg(*patternFlowIpv4Protocol) PatternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter
	Items() []PatternFlowIpv4ProtocolMetricTag
	Add() PatternFlowIpv4ProtocolMetricTag
	Append(items ...PatternFlowIpv4ProtocolMetricTag) PatternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter
	Set(index int, newObj PatternFlowIpv4ProtocolMetricTag) PatternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter
	Clear() PatternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter
	clearHolderSlice() PatternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter
	appendHolderSlice(item PatternFlowIpv4ProtocolMetricTag) PatternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter
}

func (obj *patternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter) setMsg(msg *patternFlowIpv4Protocol) PatternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv4ProtocolMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter) Items() []PatternFlowIpv4ProtocolMetricTag {
	return obj.patternFlowIpv4ProtocolMetricTagSlice
}

func (obj *patternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter) Add() PatternFlowIpv4ProtocolMetricTag {
	newObj := &otg.PatternFlowIpv4ProtocolMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv4ProtocolMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv4ProtocolMetricTagSlice = append(obj.patternFlowIpv4ProtocolMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter) Append(items ...PatternFlowIpv4ProtocolMetricTag) PatternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv4ProtocolMetricTagSlice = append(obj.patternFlowIpv4ProtocolMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter) Set(index int, newObj PatternFlowIpv4ProtocolMetricTag) PatternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv4ProtocolMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter) Clear() PatternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv4ProtocolMetricTag{}
		obj.patternFlowIpv4ProtocolMetricTagSlice = []PatternFlowIpv4ProtocolMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter) clearHolderSlice() PatternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter {
	if len(obj.patternFlowIpv4ProtocolMetricTagSlice) > 0 {
		obj.patternFlowIpv4ProtocolMetricTagSlice = []PatternFlowIpv4ProtocolMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter) appendHolderSlice(item PatternFlowIpv4ProtocolMetricTag) PatternFlowIpv4ProtocolPatternFlowIpv4ProtocolMetricTagIter {
	obj.patternFlowIpv4ProtocolMetricTagSlice = append(obj.patternFlowIpv4ProtocolMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIpv4Protocol) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4Protocol.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv4Protocol.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4Protocol.Auto <= 255 but Got %d", *obj.obj.Auto))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv4ProtocolMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIpv4Protocol) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowIpv4ProtocolChoice.AUTO)

	}

}
