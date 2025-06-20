package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowMplsLabel *****
type patternFlowMplsLabel struct {
	validation
	obj              *otg.PatternFlowMplsLabel
	marshaller       marshalPatternFlowMplsLabel
	unMarshaller     unMarshalPatternFlowMplsLabel
	incrementHolder  PatternFlowMplsLabelCounter
	decrementHolder  PatternFlowMplsLabelCounter
	metricTagsHolder PatternFlowMplsLabelPatternFlowMplsLabelMetricTagIter
}

func NewPatternFlowMplsLabel() PatternFlowMplsLabel {
	obj := patternFlowMplsLabel{obj: &otg.PatternFlowMplsLabel{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowMplsLabel) msg() *otg.PatternFlowMplsLabel {
	return obj.obj
}

func (obj *patternFlowMplsLabel) setMsg(msg *otg.PatternFlowMplsLabel) PatternFlowMplsLabel {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowMplsLabel struct {
	obj *patternFlowMplsLabel
}

type marshalPatternFlowMplsLabel interface {
	// ToProto marshals PatternFlowMplsLabel to protobuf object *otg.PatternFlowMplsLabel
	ToProto() (*otg.PatternFlowMplsLabel, error)
	// ToPbText marshals PatternFlowMplsLabel to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowMplsLabel to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowMplsLabel to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowMplsLabel struct {
	obj *patternFlowMplsLabel
}

type unMarshalPatternFlowMplsLabel interface {
	// FromProto unmarshals PatternFlowMplsLabel from protobuf object *otg.PatternFlowMplsLabel
	FromProto(msg *otg.PatternFlowMplsLabel) (PatternFlowMplsLabel, error)
	// FromPbText unmarshals PatternFlowMplsLabel from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowMplsLabel from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowMplsLabel from JSON text
	FromJson(value string) error
}

func (obj *patternFlowMplsLabel) Marshal() marshalPatternFlowMplsLabel {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowMplsLabel{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowMplsLabel) Unmarshal() unMarshalPatternFlowMplsLabel {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowMplsLabel{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowMplsLabel) ToProto() (*otg.PatternFlowMplsLabel, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowMplsLabel) FromProto(msg *otg.PatternFlowMplsLabel) (PatternFlowMplsLabel, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowMplsLabel) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowMplsLabel) FromPbText(value string) error {
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

func (m *marshalpatternFlowMplsLabel) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowMplsLabel) FromYaml(value string) error {
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

func (m *marshalpatternFlowMplsLabel) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowMplsLabel) FromJson(value string) error {
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

func (obj *patternFlowMplsLabel) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowMplsLabel) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowMplsLabel) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowMplsLabel) Clone() (PatternFlowMplsLabel, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowMplsLabel()
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

func (obj *patternFlowMplsLabel) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowMplsLabel is label of routers
type PatternFlowMplsLabel interface {
	Validation
	// msg marshals PatternFlowMplsLabel to protobuf object *otg.PatternFlowMplsLabel
	// and doesn't set defaults
	msg() *otg.PatternFlowMplsLabel
	// setMsg unmarshals PatternFlowMplsLabel from protobuf object *otg.PatternFlowMplsLabel
	// and doesn't set defaults
	setMsg(*otg.PatternFlowMplsLabel) PatternFlowMplsLabel
	// provides marshal interface
	Marshal() marshalPatternFlowMplsLabel
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowMplsLabel
	// validate validates PatternFlowMplsLabel
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowMplsLabel, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowMplsLabelChoiceEnum, set in PatternFlowMplsLabel
	Choice() PatternFlowMplsLabelChoiceEnum
	// setChoice assigns PatternFlowMplsLabelChoiceEnum provided by user to PatternFlowMplsLabel
	setChoice(value PatternFlowMplsLabelChoiceEnum) PatternFlowMplsLabel
	// HasChoice checks if Choice has been set in PatternFlowMplsLabel
	HasChoice() bool
	// Value returns uint32, set in PatternFlowMplsLabel.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowMplsLabel
	SetValue(value uint32) PatternFlowMplsLabel
	// HasValue checks if Value has been set in PatternFlowMplsLabel
	HasValue() bool
	// Values returns []uint32, set in PatternFlowMplsLabel.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowMplsLabel
	SetValues(value []uint32) PatternFlowMplsLabel
	// Auto returns uint32, set in PatternFlowMplsLabel.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowMplsLabel
	HasAuto() bool
	// Increment returns PatternFlowMplsLabelCounter, set in PatternFlowMplsLabel.
	// PatternFlowMplsLabelCounter is integer counter pattern
	Increment() PatternFlowMplsLabelCounter
	// SetIncrement assigns PatternFlowMplsLabelCounter provided by user to PatternFlowMplsLabel.
	// PatternFlowMplsLabelCounter is integer counter pattern
	SetIncrement(value PatternFlowMplsLabelCounter) PatternFlowMplsLabel
	// HasIncrement checks if Increment has been set in PatternFlowMplsLabel
	HasIncrement() bool
	// Decrement returns PatternFlowMplsLabelCounter, set in PatternFlowMplsLabel.
	// PatternFlowMplsLabelCounter is integer counter pattern
	Decrement() PatternFlowMplsLabelCounter
	// SetDecrement assigns PatternFlowMplsLabelCounter provided by user to PatternFlowMplsLabel.
	// PatternFlowMplsLabelCounter is integer counter pattern
	SetDecrement(value PatternFlowMplsLabelCounter) PatternFlowMplsLabel
	// HasDecrement checks if Decrement has been set in PatternFlowMplsLabel
	HasDecrement() bool
	// MetricTags returns PatternFlowMplsLabelPatternFlowMplsLabelMetricTagIterIter, set in PatternFlowMplsLabel
	MetricTags() PatternFlowMplsLabelPatternFlowMplsLabelMetricTagIter
	setNil()
}

type PatternFlowMplsLabelChoiceEnum string

// Enum of Choice on PatternFlowMplsLabel
var PatternFlowMplsLabelChoice = struct {
	VALUE     PatternFlowMplsLabelChoiceEnum
	VALUES    PatternFlowMplsLabelChoiceEnum
	AUTO      PatternFlowMplsLabelChoiceEnum
	INCREMENT PatternFlowMplsLabelChoiceEnum
	DECREMENT PatternFlowMplsLabelChoiceEnum
}{
	VALUE:     PatternFlowMplsLabelChoiceEnum("value"),
	VALUES:    PatternFlowMplsLabelChoiceEnum("values"),
	AUTO:      PatternFlowMplsLabelChoiceEnum("auto"),
	INCREMENT: PatternFlowMplsLabelChoiceEnum("increment"),
	DECREMENT: PatternFlowMplsLabelChoiceEnum("decrement"),
}

func (obj *patternFlowMplsLabel) Choice() PatternFlowMplsLabelChoiceEnum {
	return PatternFlowMplsLabelChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowMplsLabel) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowMplsLabel) setChoice(value PatternFlowMplsLabelChoiceEnum) PatternFlowMplsLabel {
	intValue, ok := otg.PatternFlowMplsLabel_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowMplsLabelChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowMplsLabel_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowMplsLabelChoice.VALUE {
		defaultValue := uint32(16)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowMplsLabelChoice.VALUES {
		defaultValue := []uint32{16}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowMplsLabelChoice.AUTO {
		defaultValue := uint32(16)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowMplsLabelChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowMplsLabelCounter().msg()
	}

	if value == PatternFlowMplsLabelChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowMplsLabelCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowMplsLabel) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowMplsLabelChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowMplsLabel) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowMplsLabel object
func (obj *patternFlowMplsLabel) SetValue(value uint32) PatternFlowMplsLabel {
	obj.setChoice(PatternFlowMplsLabelChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowMplsLabel) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{16})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowMplsLabel object
func (obj *patternFlowMplsLabel) SetValues(value []uint32) PatternFlowMplsLabel {
	obj.setChoice(PatternFlowMplsLabelChoice.VALUES)
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
func (obj *patternFlowMplsLabel) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowMplsLabelChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowMplsLabel) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowMplsLabelCounter
func (obj *patternFlowMplsLabel) Increment() PatternFlowMplsLabelCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowMplsLabelChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowMplsLabelCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowMplsLabelCounter
func (obj *patternFlowMplsLabel) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowMplsLabelCounter value in the PatternFlowMplsLabel object
func (obj *patternFlowMplsLabel) SetIncrement(value PatternFlowMplsLabelCounter) PatternFlowMplsLabel {
	obj.setChoice(PatternFlowMplsLabelChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowMplsLabelCounter
func (obj *patternFlowMplsLabel) Decrement() PatternFlowMplsLabelCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowMplsLabelChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowMplsLabelCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowMplsLabelCounter
func (obj *patternFlowMplsLabel) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowMplsLabelCounter value in the PatternFlowMplsLabel object
func (obj *patternFlowMplsLabel) SetDecrement(value PatternFlowMplsLabelCounter) PatternFlowMplsLabel {
	obj.setChoice(PatternFlowMplsLabelChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowMplsLabelMetricTag
func (obj *patternFlowMplsLabel) MetricTags() PatternFlowMplsLabelPatternFlowMplsLabelMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowMplsLabelMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowMplsLabelPatternFlowMplsLabelMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowMplsLabelPatternFlowMplsLabelMetricTagIter struct {
	obj                                *patternFlowMplsLabel
	patternFlowMplsLabelMetricTagSlice []PatternFlowMplsLabelMetricTag
	fieldPtr                           *[]*otg.PatternFlowMplsLabelMetricTag
}

func newPatternFlowMplsLabelPatternFlowMplsLabelMetricTagIter(ptr *[]*otg.PatternFlowMplsLabelMetricTag) PatternFlowMplsLabelPatternFlowMplsLabelMetricTagIter {
	return &patternFlowMplsLabelPatternFlowMplsLabelMetricTagIter{fieldPtr: ptr}
}

type PatternFlowMplsLabelPatternFlowMplsLabelMetricTagIter interface {
	setMsg(*patternFlowMplsLabel) PatternFlowMplsLabelPatternFlowMplsLabelMetricTagIter
	Items() []PatternFlowMplsLabelMetricTag
	Add() PatternFlowMplsLabelMetricTag
	Append(items ...PatternFlowMplsLabelMetricTag) PatternFlowMplsLabelPatternFlowMplsLabelMetricTagIter
	Set(index int, newObj PatternFlowMplsLabelMetricTag) PatternFlowMplsLabelPatternFlowMplsLabelMetricTagIter
	Clear() PatternFlowMplsLabelPatternFlowMplsLabelMetricTagIter
	clearHolderSlice() PatternFlowMplsLabelPatternFlowMplsLabelMetricTagIter
	appendHolderSlice(item PatternFlowMplsLabelMetricTag) PatternFlowMplsLabelPatternFlowMplsLabelMetricTagIter
}

func (obj *patternFlowMplsLabelPatternFlowMplsLabelMetricTagIter) setMsg(msg *patternFlowMplsLabel) PatternFlowMplsLabelPatternFlowMplsLabelMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowMplsLabelMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowMplsLabelPatternFlowMplsLabelMetricTagIter) Items() []PatternFlowMplsLabelMetricTag {
	return obj.patternFlowMplsLabelMetricTagSlice
}

func (obj *patternFlowMplsLabelPatternFlowMplsLabelMetricTagIter) Add() PatternFlowMplsLabelMetricTag {
	newObj := &otg.PatternFlowMplsLabelMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowMplsLabelMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowMplsLabelMetricTagSlice = append(obj.patternFlowMplsLabelMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowMplsLabelPatternFlowMplsLabelMetricTagIter) Append(items ...PatternFlowMplsLabelMetricTag) PatternFlowMplsLabelPatternFlowMplsLabelMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowMplsLabelMetricTagSlice = append(obj.patternFlowMplsLabelMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowMplsLabelPatternFlowMplsLabelMetricTagIter) Set(index int, newObj PatternFlowMplsLabelMetricTag) PatternFlowMplsLabelPatternFlowMplsLabelMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowMplsLabelMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowMplsLabelPatternFlowMplsLabelMetricTagIter) Clear() PatternFlowMplsLabelPatternFlowMplsLabelMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowMplsLabelMetricTag{}
		obj.patternFlowMplsLabelMetricTagSlice = []PatternFlowMplsLabelMetricTag{}
	}
	return obj
}
func (obj *patternFlowMplsLabelPatternFlowMplsLabelMetricTagIter) clearHolderSlice() PatternFlowMplsLabelPatternFlowMplsLabelMetricTagIter {
	if len(obj.patternFlowMplsLabelMetricTagSlice) > 0 {
		obj.patternFlowMplsLabelMetricTagSlice = []PatternFlowMplsLabelMetricTag{}
	}
	return obj
}
func (obj *patternFlowMplsLabelPatternFlowMplsLabelMetricTagIter) appendHolderSlice(item PatternFlowMplsLabelMetricTag) PatternFlowMplsLabelPatternFlowMplsLabelMetricTagIter {
	obj.patternFlowMplsLabelMetricTagSlice = append(obj.patternFlowMplsLabelMetricTagSlice, item)
	return obj
}

func (obj *patternFlowMplsLabel) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowMplsLabel.Value <= 1048575 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1048575 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowMplsLabel.Values <= 1048575 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowMplsLabel.Auto <= 1048575 but Got %d", *obj.obj.Auto))
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
				obj.MetricTags().appendHolderSlice(&patternFlowMplsLabelMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowMplsLabel) setDefault() {
	var choices_set int = 0
	var choice PatternFlowMplsLabelChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowMplsLabelChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowMplsLabelChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowMplsLabelChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowMplsLabelChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowMplsLabelChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowMplsLabelChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowMplsLabel")
			}
		} else {
			intVal := otg.PatternFlowMplsLabel_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowMplsLabel_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
