package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowEthernetSrc *****
type patternFlowEthernetSrc struct {
	validation
	obj              *otg.PatternFlowEthernetSrc
	marshaller       marshalPatternFlowEthernetSrc
	unMarshaller     unMarshalPatternFlowEthernetSrc
	incrementHolder  PatternFlowEthernetSrcCounter
	decrementHolder  PatternFlowEthernetSrcCounter
	metricTagsHolder PatternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter
}

func NewPatternFlowEthernetSrc() PatternFlowEthernetSrc {
	obj := patternFlowEthernetSrc{obj: &otg.PatternFlowEthernetSrc{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowEthernetSrc) msg() *otg.PatternFlowEthernetSrc {
	return obj.obj
}

func (obj *patternFlowEthernetSrc) setMsg(msg *otg.PatternFlowEthernetSrc) PatternFlowEthernetSrc {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowEthernetSrc struct {
	obj *patternFlowEthernetSrc
}

type marshalPatternFlowEthernetSrc interface {
	// ToProto marshals PatternFlowEthernetSrc to protobuf object *otg.PatternFlowEthernetSrc
	ToProto() (*otg.PatternFlowEthernetSrc, error)
	// ToPbText marshals PatternFlowEthernetSrc to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowEthernetSrc to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowEthernetSrc to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowEthernetSrc struct {
	obj *patternFlowEthernetSrc
}

type unMarshalPatternFlowEthernetSrc interface {
	// FromProto unmarshals PatternFlowEthernetSrc from protobuf object *otg.PatternFlowEthernetSrc
	FromProto(msg *otg.PatternFlowEthernetSrc) (PatternFlowEthernetSrc, error)
	// FromPbText unmarshals PatternFlowEthernetSrc from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowEthernetSrc from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowEthernetSrc from JSON text
	FromJson(value string) error
}

func (obj *patternFlowEthernetSrc) Marshal() marshalPatternFlowEthernetSrc {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowEthernetSrc{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowEthernetSrc) Unmarshal() unMarshalPatternFlowEthernetSrc {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowEthernetSrc{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowEthernetSrc) ToProto() (*otg.PatternFlowEthernetSrc, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowEthernetSrc) FromProto(msg *otg.PatternFlowEthernetSrc) (PatternFlowEthernetSrc, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowEthernetSrc) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowEthernetSrc) FromPbText(value string) error {
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

func (m *marshalpatternFlowEthernetSrc) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowEthernetSrc) FromYaml(value string) error {
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

func (m *marshalpatternFlowEthernetSrc) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowEthernetSrc) FromJson(value string) error {
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

func (obj *patternFlowEthernetSrc) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowEthernetSrc) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowEthernetSrc) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowEthernetSrc) Clone() (PatternFlowEthernetSrc, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowEthernetSrc()
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

func (obj *patternFlowEthernetSrc) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowEthernetSrc is source MAC address
type PatternFlowEthernetSrc interface {
	Validation
	// msg marshals PatternFlowEthernetSrc to protobuf object *otg.PatternFlowEthernetSrc
	// and doesn't set defaults
	msg() *otg.PatternFlowEthernetSrc
	// setMsg unmarshals PatternFlowEthernetSrc from protobuf object *otg.PatternFlowEthernetSrc
	// and doesn't set defaults
	setMsg(*otg.PatternFlowEthernetSrc) PatternFlowEthernetSrc
	// provides marshal interface
	Marshal() marshalPatternFlowEthernetSrc
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowEthernetSrc
	// validate validates PatternFlowEthernetSrc
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowEthernetSrc, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowEthernetSrcChoiceEnum, set in PatternFlowEthernetSrc
	Choice() PatternFlowEthernetSrcChoiceEnum
	// setChoice assigns PatternFlowEthernetSrcChoiceEnum provided by user to PatternFlowEthernetSrc
	setChoice(value PatternFlowEthernetSrcChoiceEnum) PatternFlowEthernetSrc
	// HasChoice checks if Choice has been set in PatternFlowEthernetSrc
	HasChoice() bool
	// Value returns string, set in PatternFlowEthernetSrc.
	Value() string
	// SetValue assigns string provided by user to PatternFlowEthernetSrc
	SetValue(value string) PatternFlowEthernetSrc
	// HasValue checks if Value has been set in PatternFlowEthernetSrc
	HasValue() bool
	// Values returns []string, set in PatternFlowEthernetSrc.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowEthernetSrc
	SetValues(value []string) PatternFlowEthernetSrc
	// Increment returns PatternFlowEthernetSrcCounter, set in PatternFlowEthernetSrc.
	// PatternFlowEthernetSrcCounter is mac counter pattern
	Increment() PatternFlowEthernetSrcCounter
	// SetIncrement assigns PatternFlowEthernetSrcCounter provided by user to PatternFlowEthernetSrc.
	// PatternFlowEthernetSrcCounter is mac counter pattern
	SetIncrement(value PatternFlowEthernetSrcCounter) PatternFlowEthernetSrc
	// HasIncrement checks if Increment has been set in PatternFlowEthernetSrc
	HasIncrement() bool
	// Decrement returns PatternFlowEthernetSrcCounter, set in PatternFlowEthernetSrc.
	// PatternFlowEthernetSrcCounter is mac counter pattern
	Decrement() PatternFlowEthernetSrcCounter
	// SetDecrement assigns PatternFlowEthernetSrcCounter provided by user to PatternFlowEthernetSrc.
	// PatternFlowEthernetSrcCounter is mac counter pattern
	SetDecrement(value PatternFlowEthernetSrcCounter) PatternFlowEthernetSrc
	// HasDecrement checks if Decrement has been set in PatternFlowEthernetSrc
	HasDecrement() bool
	// MetricTags returns PatternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIterIter, set in PatternFlowEthernetSrc
	MetricTags() PatternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter
	setNil()
}

type PatternFlowEthernetSrcChoiceEnum string

// Enum of Choice on PatternFlowEthernetSrc
var PatternFlowEthernetSrcChoice = struct {
	VALUE     PatternFlowEthernetSrcChoiceEnum
	VALUES    PatternFlowEthernetSrcChoiceEnum
	INCREMENT PatternFlowEthernetSrcChoiceEnum
	DECREMENT PatternFlowEthernetSrcChoiceEnum
}{
	VALUE:     PatternFlowEthernetSrcChoiceEnum("value"),
	VALUES:    PatternFlowEthernetSrcChoiceEnum("values"),
	INCREMENT: PatternFlowEthernetSrcChoiceEnum("increment"),
	DECREMENT: PatternFlowEthernetSrcChoiceEnum("decrement"),
}

func (obj *patternFlowEthernetSrc) Choice() PatternFlowEthernetSrcChoiceEnum {
	return PatternFlowEthernetSrcChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowEthernetSrc) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowEthernetSrc) setChoice(value PatternFlowEthernetSrcChoiceEnum) PatternFlowEthernetSrc {
	intValue, ok := otg.PatternFlowEthernetSrc_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowEthernetSrcChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowEthernetSrc_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowEthernetSrcChoice.VALUE {
		defaultValue := "00:00:00:00:00:00"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowEthernetSrcChoice.VALUES {
		defaultValue := []string{"00:00:00:00:00:00"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowEthernetSrcChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowEthernetSrcCounter().msg()
	}

	if value == PatternFlowEthernetSrcChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowEthernetSrcCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowEthernetSrc) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowEthernetSrcChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowEthernetSrc) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowEthernetSrc object
func (obj *patternFlowEthernetSrc) SetValue(value string) PatternFlowEthernetSrc {
	obj.setChoice(PatternFlowEthernetSrcChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowEthernetSrc) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"00:00:00:00:00:00"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowEthernetSrc object
func (obj *patternFlowEthernetSrc) SetValues(value []string) PatternFlowEthernetSrc {
	obj.setChoice(PatternFlowEthernetSrcChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowEthernetSrcCounter
func (obj *patternFlowEthernetSrc) Increment() PatternFlowEthernetSrcCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowEthernetSrcChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowEthernetSrcCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowEthernetSrcCounter
func (obj *patternFlowEthernetSrc) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowEthernetSrcCounter value in the PatternFlowEthernetSrc object
func (obj *patternFlowEthernetSrc) SetIncrement(value PatternFlowEthernetSrcCounter) PatternFlowEthernetSrc {
	obj.setChoice(PatternFlowEthernetSrcChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowEthernetSrcCounter
func (obj *patternFlowEthernetSrc) Decrement() PatternFlowEthernetSrcCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowEthernetSrcChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowEthernetSrcCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowEthernetSrcCounter
func (obj *patternFlowEthernetSrc) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowEthernetSrcCounter value in the PatternFlowEthernetSrc object
func (obj *patternFlowEthernetSrc) SetDecrement(value PatternFlowEthernetSrcCounter) PatternFlowEthernetSrc {
	obj.setChoice(PatternFlowEthernetSrcChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowEthernetSrcMetricTag
func (obj *patternFlowEthernetSrc) MetricTags() PatternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowEthernetSrcMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter struct {
	obj                                  *patternFlowEthernetSrc
	patternFlowEthernetSrcMetricTagSlice []PatternFlowEthernetSrcMetricTag
	fieldPtr                             *[]*otg.PatternFlowEthernetSrcMetricTag
}

func newPatternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter(ptr *[]*otg.PatternFlowEthernetSrcMetricTag) PatternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter {
	return &patternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter{fieldPtr: ptr}
}

type PatternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter interface {
	setMsg(*patternFlowEthernetSrc) PatternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter
	Items() []PatternFlowEthernetSrcMetricTag
	Add() PatternFlowEthernetSrcMetricTag
	Append(items ...PatternFlowEthernetSrcMetricTag) PatternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter
	Set(index int, newObj PatternFlowEthernetSrcMetricTag) PatternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter
	Clear() PatternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter
	clearHolderSlice() PatternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter
	appendHolderSlice(item PatternFlowEthernetSrcMetricTag) PatternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter
}

func (obj *patternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter) setMsg(msg *patternFlowEthernetSrc) PatternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowEthernetSrcMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter) Items() []PatternFlowEthernetSrcMetricTag {
	return obj.patternFlowEthernetSrcMetricTagSlice
}

func (obj *patternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter) Add() PatternFlowEthernetSrcMetricTag {
	newObj := &otg.PatternFlowEthernetSrcMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowEthernetSrcMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowEthernetSrcMetricTagSlice = append(obj.patternFlowEthernetSrcMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter) Append(items ...PatternFlowEthernetSrcMetricTag) PatternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowEthernetSrcMetricTagSlice = append(obj.patternFlowEthernetSrcMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter) Set(index int, newObj PatternFlowEthernetSrcMetricTag) PatternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowEthernetSrcMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter) Clear() PatternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowEthernetSrcMetricTag{}
		obj.patternFlowEthernetSrcMetricTagSlice = []PatternFlowEthernetSrcMetricTag{}
	}
	return obj
}
func (obj *patternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter) clearHolderSlice() PatternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter {
	if len(obj.patternFlowEthernetSrcMetricTagSlice) > 0 {
		obj.patternFlowEthernetSrcMetricTagSlice = []PatternFlowEthernetSrcMetricTag{}
	}
	return obj
}
func (obj *patternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter) appendHolderSlice(item PatternFlowEthernetSrcMetricTag) PatternFlowEthernetSrcPatternFlowEthernetSrcMetricTagIter {
	obj.patternFlowEthernetSrcMetricTagSlice = append(obj.patternFlowEthernetSrcMetricTagSlice, item)
	return obj
}

func (obj *patternFlowEthernetSrc) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetSrc.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateMacSlice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetSrc.Values"))
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
				obj.MetricTags().appendHolderSlice(&patternFlowEthernetSrcMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowEthernetSrc) setDefault() {
	var choices_set int = 0
	var choice PatternFlowEthernetSrcChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowEthernetSrcChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowEthernetSrcChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowEthernetSrcChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowEthernetSrcChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowEthernetSrcChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowEthernetSrc")
			}
		} else {
			intVal := otg.PatternFlowEthernetSrc_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowEthernetSrc_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
