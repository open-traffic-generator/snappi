package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4Src *****
type patternFlowIpv4Src struct {
	validation
	obj              *otg.PatternFlowIpv4Src
	marshaller       marshalPatternFlowIpv4Src
	unMarshaller     unMarshalPatternFlowIpv4Src
	incrementHolder  PatternFlowIpv4SrcCounter
	decrementHolder  PatternFlowIpv4SrcCounter
	metricTagsHolder PatternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter
	autoHolder       FlowIpv4Auto
}

func NewPatternFlowIpv4Src() PatternFlowIpv4Src {
	obj := patternFlowIpv4Src{obj: &otg.PatternFlowIpv4Src{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4Src) msg() *otg.PatternFlowIpv4Src {
	return obj.obj
}

func (obj *patternFlowIpv4Src) setMsg(msg *otg.PatternFlowIpv4Src) PatternFlowIpv4Src {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4Src struct {
	obj *patternFlowIpv4Src
}

type marshalPatternFlowIpv4Src interface {
	// ToProto marshals PatternFlowIpv4Src to protobuf object *otg.PatternFlowIpv4Src
	ToProto() (*otg.PatternFlowIpv4Src, error)
	// ToPbText marshals PatternFlowIpv4Src to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4Src to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4Src to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4Src struct {
	obj *patternFlowIpv4Src
}

type unMarshalPatternFlowIpv4Src interface {
	// FromProto unmarshals PatternFlowIpv4Src from protobuf object *otg.PatternFlowIpv4Src
	FromProto(msg *otg.PatternFlowIpv4Src) (PatternFlowIpv4Src, error)
	// FromPbText unmarshals PatternFlowIpv4Src from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4Src from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4Src from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4Src) Marshal() marshalPatternFlowIpv4Src {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4Src{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4Src) Unmarshal() unMarshalPatternFlowIpv4Src {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4Src{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4Src) ToProto() (*otg.PatternFlowIpv4Src, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4Src) FromProto(msg *otg.PatternFlowIpv4Src) (PatternFlowIpv4Src, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4Src) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4Src) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4Src) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4Src) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4Src) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4Src) FromJson(value string) error {
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

func (obj *patternFlowIpv4Src) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4Src) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4Src) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4Src) Clone() (PatternFlowIpv4Src, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4Src()
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

func (obj *patternFlowIpv4Src) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.autoHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4Src is source address
type PatternFlowIpv4Src interface {
	Validation
	// msg marshals PatternFlowIpv4Src to protobuf object *otg.PatternFlowIpv4Src
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4Src
	// setMsg unmarshals PatternFlowIpv4Src from protobuf object *otg.PatternFlowIpv4Src
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4Src) PatternFlowIpv4Src
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4Src
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4Src
	// validate validates PatternFlowIpv4Src
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4Src, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4SrcChoiceEnum, set in PatternFlowIpv4Src
	Choice() PatternFlowIpv4SrcChoiceEnum
	// setChoice assigns PatternFlowIpv4SrcChoiceEnum provided by user to PatternFlowIpv4Src
	setChoice(value PatternFlowIpv4SrcChoiceEnum) PatternFlowIpv4Src
	// HasChoice checks if Choice has been set in PatternFlowIpv4Src
	HasChoice() bool
	// Value returns string, set in PatternFlowIpv4Src.
	Value() string
	// SetValue assigns string provided by user to PatternFlowIpv4Src
	SetValue(value string) PatternFlowIpv4Src
	// HasValue checks if Value has been set in PatternFlowIpv4Src
	HasValue() bool
	// Values returns []string, set in PatternFlowIpv4Src.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowIpv4Src
	SetValues(value []string) PatternFlowIpv4Src
	// Increment returns PatternFlowIpv4SrcCounter, set in PatternFlowIpv4Src.
	// PatternFlowIpv4SrcCounter is ipv4 counter pattern
	Increment() PatternFlowIpv4SrcCounter
	// SetIncrement assigns PatternFlowIpv4SrcCounter provided by user to PatternFlowIpv4Src.
	// PatternFlowIpv4SrcCounter is ipv4 counter pattern
	SetIncrement(value PatternFlowIpv4SrcCounter) PatternFlowIpv4Src
	// HasIncrement checks if Increment has been set in PatternFlowIpv4Src
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4SrcCounter, set in PatternFlowIpv4Src.
	// PatternFlowIpv4SrcCounter is ipv4 counter pattern
	Decrement() PatternFlowIpv4SrcCounter
	// SetDecrement assigns PatternFlowIpv4SrcCounter provided by user to PatternFlowIpv4Src.
	// PatternFlowIpv4SrcCounter is ipv4 counter pattern
	SetDecrement(value PatternFlowIpv4SrcCounter) PatternFlowIpv4Src
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4Src
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIterIter, set in PatternFlowIpv4Src
	MetricTags() PatternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter
	// Auto returns FlowIpv4Auto, set in PatternFlowIpv4Src.
	// FlowIpv4Auto is the OTG implementation can provide a system generated, value for this property.
	Auto() FlowIpv4Auto
	// HasAuto checks if Auto has been set in PatternFlowIpv4Src
	HasAuto() bool
	setNil()
}

type PatternFlowIpv4SrcChoiceEnum string

// Enum of Choice on PatternFlowIpv4Src
var PatternFlowIpv4SrcChoice = struct {
	VALUE     PatternFlowIpv4SrcChoiceEnum
	VALUES    PatternFlowIpv4SrcChoiceEnum
	AUTO      PatternFlowIpv4SrcChoiceEnum
	INCREMENT PatternFlowIpv4SrcChoiceEnum
	DECREMENT PatternFlowIpv4SrcChoiceEnum
}{
	VALUE:     PatternFlowIpv4SrcChoiceEnum("value"),
	VALUES:    PatternFlowIpv4SrcChoiceEnum("values"),
	AUTO:      PatternFlowIpv4SrcChoiceEnum("auto"),
	INCREMENT: PatternFlowIpv4SrcChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4SrcChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4Src) Choice() PatternFlowIpv4SrcChoiceEnum {
	return PatternFlowIpv4SrcChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4Src) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4Src) setChoice(value PatternFlowIpv4SrcChoiceEnum) PatternFlowIpv4Src {
	intValue, ok := otg.PatternFlowIpv4Src_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4SrcChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4Src_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Auto = nil
	obj.autoHolder = nil
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4SrcChoice.VALUE {
		defaultValue := "0.0.0.0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4SrcChoice.VALUES {
		defaultValue := []string{"0.0.0.0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4SrcChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4SrcCounter().msg()
	}

	if value == PatternFlowIpv4SrcChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4SrcCounter().msg()
	}

	if value == PatternFlowIpv4SrcChoice.AUTO {
		obj.obj.Auto = NewFlowIpv4Auto().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowIpv4Src) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4SrcChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowIpv4Src) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowIpv4Src object
func (obj *patternFlowIpv4Src) SetValue(value string) PatternFlowIpv4Src {
	obj.setChoice(PatternFlowIpv4SrcChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowIpv4Src) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"0.0.0.0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowIpv4Src object
func (obj *patternFlowIpv4Src) SetValues(value []string) PatternFlowIpv4Src {
	obj.setChoice(PatternFlowIpv4SrcChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv4SrcCounter
func (obj *patternFlowIpv4Src) Increment() PatternFlowIpv4SrcCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4SrcChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4SrcCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4SrcCounter
func (obj *patternFlowIpv4Src) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4SrcCounter value in the PatternFlowIpv4Src object
func (obj *patternFlowIpv4Src) SetIncrement(value PatternFlowIpv4SrcCounter) PatternFlowIpv4Src {
	obj.setChoice(PatternFlowIpv4SrcChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4SrcCounter
func (obj *patternFlowIpv4Src) Decrement() PatternFlowIpv4SrcCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4SrcChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4SrcCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4SrcCounter
func (obj *patternFlowIpv4Src) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4SrcCounter value in the PatternFlowIpv4Src object
func (obj *patternFlowIpv4Src) SetDecrement(value PatternFlowIpv4SrcCounter) PatternFlowIpv4Src {
	obj.setChoice(PatternFlowIpv4SrcChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv4SrcMetricTag
func (obj *patternFlowIpv4Src) MetricTags() PatternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv4SrcMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter struct {
	obj                              *patternFlowIpv4Src
	patternFlowIpv4SrcMetricTagSlice []PatternFlowIpv4SrcMetricTag
	fieldPtr                         *[]*otg.PatternFlowIpv4SrcMetricTag
}

func newPatternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter(ptr *[]*otg.PatternFlowIpv4SrcMetricTag) PatternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter {
	return &patternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter interface {
	setMsg(*patternFlowIpv4Src) PatternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter
	Items() []PatternFlowIpv4SrcMetricTag
	Add() PatternFlowIpv4SrcMetricTag
	Append(items ...PatternFlowIpv4SrcMetricTag) PatternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter
	Set(index int, newObj PatternFlowIpv4SrcMetricTag) PatternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter
	Clear() PatternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter
	clearHolderSlice() PatternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter
	appendHolderSlice(item PatternFlowIpv4SrcMetricTag) PatternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter
}

func (obj *patternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter) setMsg(msg *patternFlowIpv4Src) PatternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv4SrcMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter) Items() []PatternFlowIpv4SrcMetricTag {
	return obj.patternFlowIpv4SrcMetricTagSlice
}

func (obj *patternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter) Add() PatternFlowIpv4SrcMetricTag {
	newObj := &otg.PatternFlowIpv4SrcMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv4SrcMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv4SrcMetricTagSlice = append(obj.patternFlowIpv4SrcMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter) Append(items ...PatternFlowIpv4SrcMetricTag) PatternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv4SrcMetricTagSlice = append(obj.patternFlowIpv4SrcMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter) Set(index int, newObj PatternFlowIpv4SrcMetricTag) PatternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv4SrcMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter) Clear() PatternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv4SrcMetricTag{}
		obj.patternFlowIpv4SrcMetricTagSlice = []PatternFlowIpv4SrcMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter) clearHolderSlice() PatternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter {
	if len(obj.patternFlowIpv4SrcMetricTagSlice) > 0 {
		obj.patternFlowIpv4SrcMetricTagSlice = []PatternFlowIpv4SrcMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter) appendHolderSlice(item PatternFlowIpv4SrcMetricTag) PatternFlowIpv4SrcPatternFlowIpv4SrcMetricTagIter {
	obj.patternFlowIpv4SrcMetricTagSlice = append(obj.patternFlowIpv4SrcMetricTagSlice, item)
	return obj
}

// description is TBD
// Auto returns a FlowIpv4Auto
func (obj *patternFlowIpv4Src) Auto() FlowIpv4Auto {
	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIpv4SrcChoice.AUTO)
	}
	if obj.autoHolder == nil {
		obj.autoHolder = &flowIpv4Auto{obj: obj.obj.Auto}
	}
	return obj.autoHolder
}

// description is TBD
// Auto returns a FlowIpv4Auto
func (obj *patternFlowIpv4Src) HasAuto() bool {
	return obj.obj.Auto != nil
}

func (obj *patternFlowIpv4Src) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv4Src.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv4Src.Values"))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv4SrcMetricTag{obj: item})
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

func (obj *patternFlowIpv4Src) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv4SrcChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv4SrcChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv4SrcChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIpv4SrcChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv4SrcChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv4SrcChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv4SrcChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv4Src")
			}
		} else {
			intVal := otg.PatternFlowIpv4Src_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv4Src_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
