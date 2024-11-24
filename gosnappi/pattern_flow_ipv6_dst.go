package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6Dst *****
type patternFlowIpv6Dst struct {
	validation
	obj              *otg.PatternFlowIpv6Dst
	marshaller       marshalPatternFlowIpv6Dst
	unMarshaller     unMarshalPatternFlowIpv6Dst
	incrementHolder  PatternFlowIpv6DstCounter
	decrementHolder  PatternFlowIpv6DstCounter
	metricTagsHolder PatternFlowIpv6DstPatternFlowIpv6DstMetricTagIter
	autoHolder       FlowIpv6Auto
}

func NewPatternFlowIpv6Dst() PatternFlowIpv6Dst {
	obj := patternFlowIpv6Dst{obj: &otg.PatternFlowIpv6Dst{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6Dst) msg() *otg.PatternFlowIpv6Dst {
	return obj.obj
}

func (obj *patternFlowIpv6Dst) setMsg(msg *otg.PatternFlowIpv6Dst) PatternFlowIpv6Dst {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6Dst struct {
	obj *patternFlowIpv6Dst
}

type marshalPatternFlowIpv6Dst interface {
	// ToProto marshals PatternFlowIpv6Dst to protobuf object *otg.PatternFlowIpv6Dst
	ToProto() (*otg.PatternFlowIpv6Dst, error)
	// ToPbText marshals PatternFlowIpv6Dst to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6Dst to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6Dst to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv6Dst to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv6Dst struct {
	obj *patternFlowIpv6Dst
}

type unMarshalPatternFlowIpv6Dst interface {
	// FromProto unmarshals PatternFlowIpv6Dst from protobuf object *otg.PatternFlowIpv6Dst
	FromProto(msg *otg.PatternFlowIpv6Dst) (PatternFlowIpv6Dst, error)
	// FromPbText unmarshals PatternFlowIpv6Dst from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6Dst from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6Dst from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6Dst) Marshal() marshalPatternFlowIpv6Dst {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6Dst{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6Dst) Unmarshal() unMarshalPatternFlowIpv6Dst {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6Dst{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6Dst) ToProto() (*otg.PatternFlowIpv6Dst, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6Dst) FromProto(msg *otg.PatternFlowIpv6Dst) (PatternFlowIpv6Dst, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6Dst) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6Dst) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6Dst) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6Dst) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6Dst) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv6Dst) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6Dst) FromJson(value string) error {
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

func (obj *patternFlowIpv6Dst) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6Dst) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6Dst) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6Dst) Clone() (PatternFlowIpv6Dst, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6Dst()
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

func (obj *patternFlowIpv6Dst) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.autoHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6Dst is destination address
type PatternFlowIpv6Dst interface {
	Validation
	// msg marshals PatternFlowIpv6Dst to protobuf object *otg.PatternFlowIpv6Dst
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6Dst
	// setMsg unmarshals PatternFlowIpv6Dst from protobuf object *otg.PatternFlowIpv6Dst
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6Dst) PatternFlowIpv6Dst
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6Dst
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6Dst
	// validate validates PatternFlowIpv6Dst
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6Dst, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6DstChoiceEnum, set in PatternFlowIpv6Dst
	Choice() PatternFlowIpv6DstChoiceEnum
	// setChoice assigns PatternFlowIpv6DstChoiceEnum provided by user to PatternFlowIpv6Dst
	setChoice(value PatternFlowIpv6DstChoiceEnum) PatternFlowIpv6Dst
	// HasChoice checks if Choice has been set in PatternFlowIpv6Dst
	HasChoice() bool
	// Value returns string, set in PatternFlowIpv6Dst.
	Value() string
	// SetValue assigns string provided by user to PatternFlowIpv6Dst
	SetValue(value string) PatternFlowIpv6Dst
	// HasValue checks if Value has been set in PatternFlowIpv6Dst
	HasValue() bool
	// Values returns []string, set in PatternFlowIpv6Dst.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowIpv6Dst
	SetValues(value []string) PatternFlowIpv6Dst
	// Increment returns PatternFlowIpv6DstCounter, set in PatternFlowIpv6Dst.
	// PatternFlowIpv6DstCounter is ipv6 counter pattern
	Increment() PatternFlowIpv6DstCounter
	// SetIncrement assigns PatternFlowIpv6DstCounter provided by user to PatternFlowIpv6Dst.
	// PatternFlowIpv6DstCounter is ipv6 counter pattern
	SetIncrement(value PatternFlowIpv6DstCounter) PatternFlowIpv6Dst
	// HasIncrement checks if Increment has been set in PatternFlowIpv6Dst
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6DstCounter, set in PatternFlowIpv6Dst.
	// PatternFlowIpv6DstCounter is ipv6 counter pattern
	Decrement() PatternFlowIpv6DstCounter
	// SetDecrement assigns PatternFlowIpv6DstCounter provided by user to PatternFlowIpv6Dst.
	// PatternFlowIpv6DstCounter is ipv6 counter pattern
	SetDecrement(value PatternFlowIpv6DstCounter) PatternFlowIpv6Dst
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6Dst
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv6DstPatternFlowIpv6DstMetricTagIterIter, set in PatternFlowIpv6Dst
	MetricTags() PatternFlowIpv6DstPatternFlowIpv6DstMetricTagIter
	// Auto returns FlowIpv6Auto, set in PatternFlowIpv6Dst.
	// FlowIpv6Auto is the OTG implementation can provide a system generated, value for this property.
	Auto() FlowIpv6Auto
	// HasAuto checks if Auto has been set in PatternFlowIpv6Dst
	HasAuto() bool
	setNil()
}

type PatternFlowIpv6DstChoiceEnum string

// Enum of Choice on PatternFlowIpv6Dst
var PatternFlowIpv6DstChoice = struct {
	VALUE     PatternFlowIpv6DstChoiceEnum
	VALUES    PatternFlowIpv6DstChoiceEnum
	INCREMENT PatternFlowIpv6DstChoiceEnum
	DECREMENT PatternFlowIpv6DstChoiceEnum
	AUTO      PatternFlowIpv6DstChoiceEnum
}{
	VALUE:     PatternFlowIpv6DstChoiceEnum("value"),
	VALUES:    PatternFlowIpv6DstChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6DstChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6DstChoiceEnum("decrement"),
	AUTO:      PatternFlowIpv6DstChoiceEnum("auto"),
}

func (obj *patternFlowIpv6Dst) Choice() PatternFlowIpv6DstChoiceEnum {
	return PatternFlowIpv6DstChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6Dst) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6Dst) setChoice(value PatternFlowIpv6DstChoiceEnum) PatternFlowIpv6Dst {
	intValue, ok := otg.PatternFlowIpv6Dst_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6DstChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6Dst_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Auto = nil
	obj.autoHolder = nil
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6DstChoice.VALUE {
		defaultValue := "::0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6DstChoice.VALUES {
		defaultValue := []string{"::0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6DstChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6DstCounter().msg()
	}

	if value == PatternFlowIpv6DstChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6DstCounter().msg()
	}

	if value == PatternFlowIpv6DstChoice.AUTO {
		obj.obj.Auto = NewFlowIpv6Auto().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowIpv6Dst) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6DstChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowIpv6Dst) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowIpv6Dst object
func (obj *patternFlowIpv6Dst) SetValue(value string) PatternFlowIpv6Dst {
	obj.setChoice(PatternFlowIpv6DstChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowIpv6Dst) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"::0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowIpv6Dst object
func (obj *patternFlowIpv6Dst) SetValues(value []string) PatternFlowIpv6Dst {
	obj.setChoice(PatternFlowIpv6DstChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6DstCounter
func (obj *patternFlowIpv6Dst) Increment() PatternFlowIpv6DstCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6DstChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6DstCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6DstCounter
func (obj *patternFlowIpv6Dst) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6DstCounter value in the PatternFlowIpv6Dst object
func (obj *patternFlowIpv6Dst) SetIncrement(value PatternFlowIpv6DstCounter) PatternFlowIpv6Dst {
	obj.setChoice(PatternFlowIpv6DstChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6DstCounter
func (obj *patternFlowIpv6Dst) Decrement() PatternFlowIpv6DstCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6DstChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6DstCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6DstCounter
func (obj *patternFlowIpv6Dst) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6DstCounter value in the PatternFlowIpv6Dst object
func (obj *patternFlowIpv6Dst) SetDecrement(value PatternFlowIpv6DstCounter) PatternFlowIpv6Dst {
	obj.setChoice(PatternFlowIpv6DstChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv6DstMetricTag
func (obj *patternFlowIpv6Dst) MetricTags() PatternFlowIpv6DstPatternFlowIpv6DstMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv6DstMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv6DstPatternFlowIpv6DstMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv6DstPatternFlowIpv6DstMetricTagIter struct {
	obj                              *patternFlowIpv6Dst
	patternFlowIpv6DstMetricTagSlice []PatternFlowIpv6DstMetricTag
	fieldPtr                         *[]*otg.PatternFlowIpv6DstMetricTag
}

func newPatternFlowIpv6DstPatternFlowIpv6DstMetricTagIter(ptr *[]*otg.PatternFlowIpv6DstMetricTag) PatternFlowIpv6DstPatternFlowIpv6DstMetricTagIter {
	return &patternFlowIpv6DstPatternFlowIpv6DstMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv6DstPatternFlowIpv6DstMetricTagIter interface {
	setMsg(*patternFlowIpv6Dst) PatternFlowIpv6DstPatternFlowIpv6DstMetricTagIter
	Items() []PatternFlowIpv6DstMetricTag
	Add() PatternFlowIpv6DstMetricTag
	Append(items ...PatternFlowIpv6DstMetricTag) PatternFlowIpv6DstPatternFlowIpv6DstMetricTagIter
	Set(index int, newObj PatternFlowIpv6DstMetricTag) PatternFlowIpv6DstPatternFlowIpv6DstMetricTagIter
	Clear() PatternFlowIpv6DstPatternFlowIpv6DstMetricTagIter
	clearHolderSlice() PatternFlowIpv6DstPatternFlowIpv6DstMetricTagIter
	appendHolderSlice(item PatternFlowIpv6DstMetricTag) PatternFlowIpv6DstPatternFlowIpv6DstMetricTagIter
}

func (obj *patternFlowIpv6DstPatternFlowIpv6DstMetricTagIter) setMsg(msg *patternFlowIpv6Dst) PatternFlowIpv6DstPatternFlowIpv6DstMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv6DstMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv6DstPatternFlowIpv6DstMetricTagIter) Items() []PatternFlowIpv6DstMetricTag {
	return obj.patternFlowIpv6DstMetricTagSlice
}

func (obj *patternFlowIpv6DstPatternFlowIpv6DstMetricTagIter) Add() PatternFlowIpv6DstMetricTag {
	newObj := &otg.PatternFlowIpv6DstMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv6DstMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv6DstMetricTagSlice = append(obj.patternFlowIpv6DstMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv6DstPatternFlowIpv6DstMetricTagIter) Append(items ...PatternFlowIpv6DstMetricTag) PatternFlowIpv6DstPatternFlowIpv6DstMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv6DstMetricTagSlice = append(obj.patternFlowIpv6DstMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv6DstPatternFlowIpv6DstMetricTagIter) Set(index int, newObj PatternFlowIpv6DstMetricTag) PatternFlowIpv6DstPatternFlowIpv6DstMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv6DstMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv6DstPatternFlowIpv6DstMetricTagIter) Clear() PatternFlowIpv6DstPatternFlowIpv6DstMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv6DstMetricTag{}
		obj.patternFlowIpv6DstMetricTagSlice = []PatternFlowIpv6DstMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv6DstPatternFlowIpv6DstMetricTagIter) clearHolderSlice() PatternFlowIpv6DstPatternFlowIpv6DstMetricTagIter {
	if len(obj.patternFlowIpv6DstMetricTagSlice) > 0 {
		obj.patternFlowIpv6DstMetricTagSlice = []PatternFlowIpv6DstMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv6DstPatternFlowIpv6DstMetricTagIter) appendHolderSlice(item PatternFlowIpv6DstMetricTag) PatternFlowIpv6DstPatternFlowIpv6DstMetricTagIter {
	obj.patternFlowIpv6DstMetricTagSlice = append(obj.patternFlowIpv6DstMetricTagSlice, item)
	return obj
}

// description is TBD
// Auto returns a FlowIpv6Auto
func (obj *patternFlowIpv6Dst) Auto() FlowIpv6Auto {
	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIpv6DstChoice.AUTO)
	}
	if obj.autoHolder == nil {
		obj.autoHolder = &flowIpv6Auto{obj: obj.obj.Auto}
	}
	return obj.autoHolder
}

// description is TBD
// Auto returns a FlowIpv6Auto
func (obj *patternFlowIpv6Dst) HasAuto() bool {
	return obj.obj.Auto != nil
}

func (obj *patternFlowIpv6Dst) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv6(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6Dst.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv6Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6Dst.Values"))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv6DstMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.Auto != nil {

		obj.Auto().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6Dst) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6DstChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6DstChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6DstChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6DstChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6DstChoice.DECREMENT
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIpv6DstChoice.AUTO
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6DstChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6Dst")
			}
		} else {
			intVal := otg.PatternFlowIpv6Dst_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6Dst_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
