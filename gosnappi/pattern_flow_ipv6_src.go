package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6Src *****
type patternFlowIpv6Src struct {
	validation
	obj              *otg.PatternFlowIpv6Src
	marshaller       marshalPatternFlowIpv6Src
	unMarshaller     unMarshalPatternFlowIpv6Src
	incrementHolder  PatternFlowIpv6SrcCounter
	decrementHolder  PatternFlowIpv6SrcCounter
	metricTagsHolder PatternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter
	autoHolder       FlowIpv6Auto
}

func NewPatternFlowIpv6Src() PatternFlowIpv6Src {
	obj := patternFlowIpv6Src{obj: &otg.PatternFlowIpv6Src{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6Src) msg() *otg.PatternFlowIpv6Src {
	return obj.obj
}

func (obj *patternFlowIpv6Src) setMsg(msg *otg.PatternFlowIpv6Src) PatternFlowIpv6Src {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6Src struct {
	obj *patternFlowIpv6Src
}

type marshalPatternFlowIpv6Src interface {
	// ToProto marshals PatternFlowIpv6Src to protobuf object *otg.PatternFlowIpv6Src
	ToProto() (*otg.PatternFlowIpv6Src, error)
	// ToPbText marshals PatternFlowIpv6Src to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6Src to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6Src to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv6Src to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv6Src struct {
	obj *patternFlowIpv6Src
}

type unMarshalPatternFlowIpv6Src interface {
	// FromProto unmarshals PatternFlowIpv6Src from protobuf object *otg.PatternFlowIpv6Src
	FromProto(msg *otg.PatternFlowIpv6Src) (PatternFlowIpv6Src, error)
	// FromPbText unmarshals PatternFlowIpv6Src from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6Src from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6Src from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6Src) Marshal() marshalPatternFlowIpv6Src {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6Src{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6Src) Unmarshal() unMarshalPatternFlowIpv6Src {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6Src{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6Src) ToProto() (*otg.PatternFlowIpv6Src, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6Src) FromProto(msg *otg.PatternFlowIpv6Src) (PatternFlowIpv6Src, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6Src) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6Src) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6Src) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6Src) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6Src) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv6Src) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6Src) FromJson(value string) error {
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

func (obj *patternFlowIpv6Src) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6Src) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6Src) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6Src) Clone() (PatternFlowIpv6Src, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6Src()
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

func (obj *patternFlowIpv6Src) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.autoHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6Src is source address
type PatternFlowIpv6Src interface {
	Validation
	// msg marshals PatternFlowIpv6Src to protobuf object *otg.PatternFlowIpv6Src
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6Src
	// setMsg unmarshals PatternFlowIpv6Src from protobuf object *otg.PatternFlowIpv6Src
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6Src) PatternFlowIpv6Src
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6Src
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6Src
	// validate validates PatternFlowIpv6Src
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6Src, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SrcChoiceEnum, set in PatternFlowIpv6Src
	Choice() PatternFlowIpv6SrcChoiceEnum
	// setChoice assigns PatternFlowIpv6SrcChoiceEnum provided by user to PatternFlowIpv6Src
	setChoice(value PatternFlowIpv6SrcChoiceEnum) PatternFlowIpv6Src
	// HasChoice checks if Choice has been set in PatternFlowIpv6Src
	HasChoice() bool
	// Value returns string, set in PatternFlowIpv6Src.
	Value() string
	// SetValue assigns string provided by user to PatternFlowIpv6Src
	SetValue(value string) PatternFlowIpv6Src
	// HasValue checks if Value has been set in PatternFlowIpv6Src
	HasValue() bool
	// Values returns []string, set in PatternFlowIpv6Src.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowIpv6Src
	SetValues(value []string) PatternFlowIpv6Src
	// Increment returns PatternFlowIpv6SrcCounter, set in PatternFlowIpv6Src.
	// PatternFlowIpv6SrcCounter is ipv6 counter pattern
	Increment() PatternFlowIpv6SrcCounter
	// SetIncrement assigns PatternFlowIpv6SrcCounter provided by user to PatternFlowIpv6Src.
	// PatternFlowIpv6SrcCounter is ipv6 counter pattern
	SetIncrement(value PatternFlowIpv6SrcCounter) PatternFlowIpv6Src
	// HasIncrement checks if Increment has been set in PatternFlowIpv6Src
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SrcCounter, set in PatternFlowIpv6Src.
	// PatternFlowIpv6SrcCounter is ipv6 counter pattern
	Decrement() PatternFlowIpv6SrcCounter
	// SetDecrement assigns PatternFlowIpv6SrcCounter provided by user to PatternFlowIpv6Src.
	// PatternFlowIpv6SrcCounter is ipv6 counter pattern
	SetDecrement(value PatternFlowIpv6SrcCounter) PatternFlowIpv6Src
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6Src
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIterIter, set in PatternFlowIpv6Src
	MetricTags() PatternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter
	// Auto returns FlowIpv6Auto, set in PatternFlowIpv6Src.
	// FlowIpv6Auto is the OTG implementation can provide a system generated, value for this property.
	Auto() FlowIpv6Auto
	// HasAuto checks if Auto has been set in PatternFlowIpv6Src
	HasAuto() bool
	setNil()
}

type PatternFlowIpv6SrcChoiceEnum string

// Enum of Choice on PatternFlowIpv6Src
var PatternFlowIpv6SrcChoice = struct {
	VALUE     PatternFlowIpv6SrcChoiceEnum
	VALUES    PatternFlowIpv6SrcChoiceEnum
	INCREMENT PatternFlowIpv6SrcChoiceEnum
	DECREMENT PatternFlowIpv6SrcChoiceEnum
	AUTO      PatternFlowIpv6SrcChoiceEnum
}{
	VALUE:     PatternFlowIpv6SrcChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SrcChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SrcChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SrcChoiceEnum("decrement"),
	AUTO:      PatternFlowIpv6SrcChoiceEnum("auto"),
}

func (obj *patternFlowIpv6Src) Choice() PatternFlowIpv6SrcChoiceEnum {
	return PatternFlowIpv6SrcChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6Src) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6Src) setChoice(value PatternFlowIpv6SrcChoiceEnum) PatternFlowIpv6Src {
	intValue, ok := otg.PatternFlowIpv6Src_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SrcChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6Src_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Auto = nil
	obj.autoHolder = nil
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SrcChoice.VALUE {
		defaultValue := "::0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SrcChoice.VALUES {
		defaultValue := []string{"::0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SrcChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SrcCounter().msg()
	}

	if value == PatternFlowIpv6SrcChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SrcCounter().msg()
	}

	if value == PatternFlowIpv6SrcChoice.AUTO {
		obj.obj.Auto = NewFlowIpv6Auto().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowIpv6Src) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SrcChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowIpv6Src) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowIpv6Src object
func (obj *patternFlowIpv6Src) SetValue(value string) PatternFlowIpv6Src {
	obj.setChoice(PatternFlowIpv6SrcChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowIpv6Src) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"::0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowIpv6Src object
func (obj *patternFlowIpv6Src) SetValues(value []string) PatternFlowIpv6Src {
	obj.setChoice(PatternFlowIpv6SrcChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SrcCounter
func (obj *patternFlowIpv6Src) Increment() PatternFlowIpv6SrcCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SrcChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SrcCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SrcCounter
func (obj *patternFlowIpv6Src) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SrcCounter value in the PatternFlowIpv6Src object
func (obj *patternFlowIpv6Src) SetIncrement(value PatternFlowIpv6SrcCounter) PatternFlowIpv6Src {
	obj.setChoice(PatternFlowIpv6SrcChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SrcCounter
func (obj *patternFlowIpv6Src) Decrement() PatternFlowIpv6SrcCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SrcChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SrcCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SrcCounter
func (obj *patternFlowIpv6Src) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SrcCounter value in the PatternFlowIpv6Src object
func (obj *patternFlowIpv6Src) SetDecrement(value PatternFlowIpv6SrcCounter) PatternFlowIpv6Src {
	obj.setChoice(PatternFlowIpv6SrcChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv6SrcMetricTag
func (obj *patternFlowIpv6Src) MetricTags() PatternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv6SrcMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter struct {
	obj                              *patternFlowIpv6Src
	patternFlowIpv6SrcMetricTagSlice []PatternFlowIpv6SrcMetricTag
	fieldPtr                         *[]*otg.PatternFlowIpv6SrcMetricTag
}

func newPatternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter(ptr *[]*otg.PatternFlowIpv6SrcMetricTag) PatternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter {
	return &patternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter interface {
	setMsg(*patternFlowIpv6Src) PatternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter
	Items() []PatternFlowIpv6SrcMetricTag
	Add() PatternFlowIpv6SrcMetricTag
	Append(items ...PatternFlowIpv6SrcMetricTag) PatternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter
	Set(index int, newObj PatternFlowIpv6SrcMetricTag) PatternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter
	Clear() PatternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter
	clearHolderSlice() PatternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter
	appendHolderSlice(item PatternFlowIpv6SrcMetricTag) PatternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter
}

func (obj *patternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter) setMsg(msg *patternFlowIpv6Src) PatternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv6SrcMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter) Items() []PatternFlowIpv6SrcMetricTag {
	return obj.patternFlowIpv6SrcMetricTagSlice
}

func (obj *patternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter) Add() PatternFlowIpv6SrcMetricTag {
	newObj := &otg.PatternFlowIpv6SrcMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv6SrcMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv6SrcMetricTagSlice = append(obj.patternFlowIpv6SrcMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter) Append(items ...PatternFlowIpv6SrcMetricTag) PatternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv6SrcMetricTagSlice = append(obj.patternFlowIpv6SrcMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter) Set(index int, newObj PatternFlowIpv6SrcMetricTag) PatternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv6SrcMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter) Clear() PatternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv6SrcMetricTag{}
		obj.patternFlowIpv6SrcMetricTagSlice = []PatternFlowIpv6SrcMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter) clearHolderSlice() PatternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter {
	if len(obj.patternFlowIpv6SrcMetricTagSlice) > 0 {
		obj.patternFlowIpv6SrcMetricTagSlice = []PatternFlowIpv6SrcMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter) appendHolderSlice(item PatternFlowIpv6SrcMetricTag) PatternFlowIpv6SrcPatternFlowIpv6SrcMetricTagIter {
	obj.patternFlowIpv6SrcMetricTagSlice = append(obj.patternFlowIpv6SrcMetricTagSlice, item)
	return obj
}

// description is TBD
// Auto returns a FlowIpv6Auto
func (obj *patternFlowIpv6Src) Auto() FlowIpv6Auto {
	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIpv6SrcChoice.AUTO)
	}
	if obj.autoHolder == nil {
		obj.autoHolder = &flowIpv6Auto{obj: obj.obj.Auto}
	}
	return obj.autoHolder
}

// description is TBD
// Auto returns a FlowIpv6Auto
func (obj *patternFlowIpv6Src) HasAuto() bool {
	return obj.obj.Auto != nil
}

func (obj *patternFlowIpv6Src) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv6(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6Src.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv6Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv6Src.Values"))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv6SrcMetricTag{obj: item})
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

func (obj *patternFlowIpv6Src) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SrcChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SrcChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SrcChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SrcChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SrcChoice.DECREMENT
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIpv6SrcChoice.AUTO
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SrcChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6Src")
			}
		} else {
			intVal := otg.PatternFlowIpv6Src_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6Src_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
