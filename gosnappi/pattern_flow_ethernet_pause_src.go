package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowEthernetPauseSrc *****
type patternFlowEthernetPauseSrc struct {
	validation
	obj              *otg.PatternFlowEthernetPauseSrc
	marshaller       marshalPatternFlowEthernetPauseSrc
	unMarshaller     unMarshalPatternFlowEthernetPauseSrc
	incrementHolder  PatternFlowEthernetPauseSrcCounter
	decrementHolder  PatternFlowEthernetPauseSrcCounter
	metricTagsHolder PatternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter
}

func NewPatternFlowEthernetPauseSrc() PatternFlowEthernetPauseSrc {
	obj := patternFlowEthernetPauseSrc{obj: &otg.PatternFlowEthernetPauseSrc{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowEthernetPauseSrc) msg() *otg.PatternFlowEthernetPauseSrc {
	return obj.obj
}

func (obj *patternFlowEthernetPauseSrc) setMsg(msg *otg.PatternFlowEthernetPauseSrc) PatternFlowEthernetPauseSrc {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowEthernetPauseSrc struct {
	obj *patternFlowEthernetPauseSrc
}

type marshalPatternFlowEthernetPauseSrc interface {
	// ToProto marshals PatternFlowEthernetPauseSrc to protobuf object *otg.PatternFlowEthernetPauseSrc
	ToProto() (*otg.PatternFlowEthernetPauseSrc, error)
	// ToPbText marshals PatternFlowEthernetPauseSrc to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowEthernetPauseSrc to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowEthernetPauseSrc to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowEthernetPauseSrc struct {
	obj *patternFlowEthernetPauseSrc
}

type unMarshalPatternFlowEthernetPauseSrc interface {
	// FromProto unmarshals PatternFlowEthernetPauseSrc from protobuf object *otg.PatternFlowEthernetPauseSrc
	FromProto(msg *otg.PatternFlowEthernetPauseSrc) (PatternFlowEthernetPauseSrc, error)
	// FromPbText unmarshals PatternFlowEthernetPauseSrc from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowEthernetPauseSrc from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowEthernetPauseSrc from JSON text
	FromJson(value string) error
}

func (obj *patternFlowEthernetPauseSrc) Marshal() marshalPatternFlowEthernetPauseSrc {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowEthernetPauseSrc{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowEthernetPauseSrc) Unmarshal() unMarshalPatternFlowEthernetPauseSrc {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowEthernetPauseSrc{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowEthernetPauseSrc) ToProto() (*otg.PatternFlowEthernetPauseSrc, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowEthernetPauseSrc) FromProto(msg *otg.PatternFlowEthernetPauseSrc) (PatternFlowEthernetPauseSrc, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowEthernetPauseSrc) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseSrc) FromPbText(value string) error {
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

func (m *marshalpatternFlowEthernetPauseSrc) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseSrc) FromYaml(value string) error {
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

func (m *marshalpatternFlowEthernetPauseSrc) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseSrc) FromJson(value string) error {
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

func (obj *patternFlowEthernetPauseSrc) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseSrc) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseSrc) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowEthernetPauseSrc) Clone() (PatternFlowEthernetPauseSrc, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowEthernetPauseSrc()
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

func (obj *patternFlowEthernetPauseSrc) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowEthernetPauseSrc is source MAC address
type PatternFlowEthernetPauseSrc interface {
	Validation
	// msg marshals PatternFlowEthernetPauseSrc to protobuf object *otg.PatternFlowEthernetPauseSrc
	// and doesn't set defaults
	msg() *otg.PatternFlowEthernetPauseSrc
	// setMsg unmarshals PatternFlowEthernetPauseSrc from protobuf object *otg.PatternFlowEthernetPauseSrc
	// and doesn't set defaults
	setMsg(*otg.PatternFlowEthernetPauseSrc) PatternFlowEthernetPauseSrc
	// provides marshal interface
	Marshal() marshalPatternFlowEthernetPauseSrc
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowEthernetPauseSrc
	// validate validates PatternFlowEthernetPauseSrc
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowEthernetPauseSrc, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowEthernetPauseSrcChoiceEnum, set in PatternFlowEthernetPauseSrc
	Choice() PatternFlowEthernetPauseSrcChoiceEnum
	// setChoice assigns PatternFlowEthernetPauseSrcChoiceEnum provided by user to PatternFlowEthernetPauseSrc
	setChoice(value PatternFlowEthernetPauseSrcChoiceEnum) PatternFlowEthernetPauseSrc
	// HasChoice checks if Choice has been set in PatternFlowEthernetPauseSrc
	HasChoice() bool
	// Value returns string, set in PatternFlowEthernetPauseSrc.
	Value() string
	// SetValue assigns string provided by user to PatternFlowEthernetPauseSrc
	SetValue(value string) PatternFlowEthernetPauseSrc
	// HasValue checks if Value has been set in PatternFlowEthernetPauseSrc
	HasValue() bool
	// Values returns []string, set in PatternFlowEthernetPauseSrc.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowEthernetPauseSrc
	SetValues(value []string) PatternFlowEthernetPauseSrc
	// Increment returns PatternFlowEthernetPauseSrcCounter, set in PatternFlowEthernetPauseSrc.
	// PatternFlowEthernetPauseSrcCounter is mac counter pattern
	Increment() PatternFlowEthernetPauseSrcCounter
	// SetIncrement assigns PatternFlowEthernetPauseSrcCounter provided by user to PatternFlowEthernetPauseSrc.
	// PatternFlowEthernetPauseSrcCounter is mac counter pattern
	SetIncrement(value PatternFlowEthernetPauseSrcCounter) PatternFlowEthernetPauseSrc
	// HasIncrement checks if Increment has been set in PatternFlowEthernetPauseSrc
	HasIncrement() bool
	// Decrement returns PatternFlowEthernetPauseSrcCounter, set in PatternFlowEthernetPauseSrc.
	// PatternFlowEthernetPauseSrcCounter is mac counter pattern
	Decrement() PatternFlowEthernetPauseSrcCounter
	// SetDecrement assigns PatternFlowEthernetPauseSrcCounter provided by user to PatternFlowEthernetPauseSrc.
	// PatternFlowEthernetPauseSrcCounter is mac counter pattern
	SetDecrement(value PatternFlowEthernetPauseSrcCounter) PatternFlowEthernetPauseSrc
	// HasDecrement checks if Decrement has been set in PatternFlowEthernetPauseSrc
	HasDecrement() bool
	// MetricTags returns PatternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIterIter, set in PatternFlowEthernetPauseSrc
	MetricTags() PatternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter
	setNil()
}

type PatternFlowEthernetPauseSrcChoiceEnum string

// Enum of Choice on PatternFlowEthernetPauseSrc
var PatternFlowEthernetPauseSrcChoice = struct {
	VALUE     PatternFlowEthernetPauseSrcChoiceEnum
	VALUES    PatternFlowEthernetPauseSrcChoiceEnum
	INCREMENT PatternFlowEthernetPauseSrcChoiceEnum
	DECREMENT PatternFlowEthernetPauseSrcChoiceEnum
}{
	VALUE:     PatternFlowEthernetPauseSrcChoiceEnum("value"),
	VALUES:    PatternFlowEthernetPauseSrcChoiceEnum("values"),
	INCREMENT: PatternFlowEthernetPauseSrcChoiceEnum("increment"),
	DECREMENT: PatternFlowEthernetPauseSrcChoiceEnum("decrement"),
}

func (obj *patternFlowEthernetPauseSrc) Choice() PatternFlowEthernetPauseSrcChoiceEnum {
	return PatternFlowEthernetPauseSrcChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowEthernetPauseSrc) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowEthernetPauseSrc) setChoice(value PatternFlowEthernetPauseSrcChoiceEnum) PatternFlowEthernetPauseSrc {
	intValue, ok := otg.PatternFlowEthernetPauseSrc_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowEthernetPauseSrcChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowEthernetPauseSrc_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowEthernetPauseSrcChoice.VALUE {
		defaultValue := "00:00:00:00:00:00"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowEthernetPauseSrcChoice.VALUES {
		defaultValue := []string{"00:00:00:00:00:00"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowEthernetPauseSrcChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowEthernetPauseSrcCounter().msg()
	}

	if value == PatternFlowEthernetPauseSrcChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowEthernetPauseSrcCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowEthernetPauseSrc) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowEthernetPauseSrcChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowEthernetPauseSrc) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowEthernetPauseSrc object
func (obj *patternFlowEthernetPauseSrc) SetValue(value string) PatternFlowEthernetPauseSrc {
	obj.setChoice(PatternFlowEthernetPauseSrcChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowEthernetPauseSrc) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"00:00:00:00:00:00"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowEthernetPauseSrc object
func (obj *patternFlowEthernetPauseSrc) SetValues(value []string) PatternFlowEthernetPauseSrc {
	obj.setChoice(PatternFlowEthernetPauseSrcChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowEthernetPauseSrcCounter
func (obj *patternFlowEthernetPauseSrc) Increment() PatternFlowEthernetPauseSrcCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowEthernetPauseSrcChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowEthernetPauseSrcCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowEthernetPauseSrcCounter
func (obj *patternFlowEthernetPauseSrc) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowEthernetPauseSrcCounter value in the PatternFlowEthernetPauseSrc object
func (obj *patternFlowEthernetPauseSrc) SetIncrement(value PatternFlowEthernetPauseSrcCounter) PatternFlowEthernetPauseSrc {
	obj.setChoice(PatternFlowEthernetPauseSrcChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowEthernetPauseSrcCounter
func (obj *patternFlowEthernetPauseSrc) Decrement() PatternFlowEthernetPauseSrcCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowEthernetPauseSrcChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowEthernetPauseSrcCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowEthernetPauseSrcCounter
func (obj *patternFlowEthernetPauseSrc) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowEthernetPauseSrcCounter value in the PatternFlowEthernetPauseSrc object
func (obj *patternFlowEthernetPauseSrc) SetDecrement(value PatternFlowEthernetPauseSrcCounter) PatternFlowEthernetPauseSrc {
	obj.setChoice(PatternFlowEthernetPauseSrcChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowEthernetPauseSrcMetricTag
func (obj *patternFlowEthernetPauseSrc) MetricTags() PatternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowEthernetPauseSrcMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter struct {
	obj                                       *patternFlowEthernetPauseSrc
	patternFlowEthernetPauseSrcMetricTagSlice []PatternFlowEthernetPauseSrcMetricTag
	fieldPtr                                  *[]*otg.PatternFlowEthernetPauseSrcMetricTag
}

func newPatternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter(ptr *[]*otg.PatternFlowEthernetPauseSrcMetricTag) PatternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter {
	return &patternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter{fieldPtr: ptr}
}

type PatternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter interface {
	setMsg(*patternFlowEthernetPauseSrc) PatternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter
	Items() []PatternFlowEthernetPauseSrcMetricTag
	Add() PatternFlowEthernetPauseSrcMetricTag
	Append(items ...PatternFlowEthernetPauseSrcMetricTag) PatternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter
	Set(index int, newObj PatternFlowEthernetPauseSrcMetricTag) PatternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter
	Clear() PatternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter
	clearHolderSlice() PatternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter
	appendHolderSlice(item PatternFlowEthernetPauseSrcMetricTag) PatternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter
}

func (obj *patternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter) setMsg(msg *patternFlowEthernetPauseSrc) PatternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowEthernetPauseSrcMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter) Items() []PatternFlowEthernetPauseSrcMetricTag {
	return obj.patternFlowEthernetPauseSrcMetricTagSlice
}

func (obj *patternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter) Add() PatternFlowEthernetPauseSrcMetricTag {
	newObj := &otg.PatternFlowEthernetPauseSrcMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowEthernetPauseSrcMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowEthernetPauseSrcMetricTagSlice = append(obj.patternFlowEthernetPauseSrcMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter) Append(items ...PatternFlowEthernetPauseSrcMetricTag) PatternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowEthernetPauseSrcMetricTagSlice = append(obj.patternFlowEthernetPauseSrcMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter) Set(index int, newObj PatternFlowEthernetPauseSrcMetricTag) PatternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowEthernetPauseSrcMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter) Clear() PatternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowEthernetPauseSrcMetricTag{}
		obj.patternFlowEthernetPauseSrcMetricTagSlice = []PatternFlowEthernetPauseSrcMetricTag{}
	}
	return obj
}
func (obj *patternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter) clearHolderSlice() PatternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter {
	if len(obj.patternFlowEthernetPauseSrcMetricTagSlice) > 0 {
		obj.patternFlowEthernetPauseSrcMetricTagSlice = []PatternFlowEthernetPauseSrcMetricTag{}
	}
	return obj
}
func (obj *patternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter) appendHolderSlice(item PatternFlowEthernetPauseSrcMetricTag) PatternFlowEthernetPauseSrcPatternFlowEthernetPauseSrcMetricTagIter {
	obj.patternFlowEthernetPauseSrcMetricTagSlice = append(obj.patternFlowEthernetPauseSrcMetricTagSlice, item)
	return obj
}

func (obj *patternFlowEthernetPauseSrc) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetPauseSrc.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateMacSlice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetPauseSrc.Values"))
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
				obj.MetricTags().appendHolderSlice(&patternFlowEthernetPauseSrcMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowEthernetPauseSrc) setDefault() {
	var choices_set int = 0
	var choice PatternFlowEthernetPauseSrcChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowEthernetPauseSrcChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowEthernetPauseSrcChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowEthernetPauseSrcChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowEthernetPauseSrcChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowEthernetPauseSrcChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowEthernetPauseSrc")
			}
		} else {
			intVal := otg.PatternFlowEthernetPauseSrc_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowEthernetPauseSrc_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
