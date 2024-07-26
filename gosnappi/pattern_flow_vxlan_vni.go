package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowVxlanVni *****
type patternFlowVxlanVni struct {
	validation
	obj              *otg.PatternFlowVxlanVni
	marshaller       marshalPatternFlowVxlanVni
	unMarshaller     unMarshalPatternFlowVxlanVni
	incrementHolder  PatternFlowVxlanVniCounter
	decrementHolder  PatternFlowVxlanVniCounter
	metricTagsHolder PatternFlowVxlanVniPatternFlowVxlanVniMetricTagIter
}

func NewPatternFlowVxlanVni() PatternFlowVxlanVni {
	obj := patternFlowVxlanVni{obj: &otg.PatternFlowVxlanVni{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowVxlanVni) msg() *otg.PatternFlowVxlanVni {
	return obj.obj
}

func (obj *patternFlowVxlanVni) setMsg(msg *otg.PatternFlowVxlanVni) PatternFlowVxlanVni {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowVxlanVni struct {
	obj *patternFlowVxlanVni
}

type marshalPatternFlowVxlanVni interface {
	// ToProto marshals PatternFlowVxlanVni to protobuf object *otg.PatternFlowVxlanVni
	ToProto() (*otg.PatternFlowVxlanVni, error)
	// ToPbText marshals PatternFlowVxlanVni to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowVxlanVni to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowVxlanVni to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowVxlanVni struct {
	obj *patternFlowVxlanVni
}

type unMarshalPatternFlowVxlanVni interface {
	// FromProto unmarshals PatternFlowVxlanVni from protobuf object *otg.PatternFlowVxlanVni
	FromProto(msg *otg.PatternFlowVxlanVni) (PatternFlowVxlanVni, error)
	// FromPbText unmarshals PatternFlowVxlanVni from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowVxlanVni from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowVxlanVni from JSON text
	FromJson(value string) error
}

func (obj *patternFlowVxlanVni) Marshal() marshalPatternFlowVxlanVni {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowVxlanVni{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowVxlanVni) Unmarshal() unMarshalPatternFlowVxlanVni {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowVxlanVni{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowVxlanVni) ToProto() (*otg.PatternFlowVxlanVni, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowVxlanVni) FromProto(msg *otg.PatternFlowVxlanVni) (PatternFlowVxlanVni, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowVxlanVni) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowVxlanVni) FromPbText(value string) error {
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

func (m *marshalpatternFlowVxlanVni) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowVxlanVni) FromYaml(value string) error {
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

func (m *marshalpatternFlowVxlanVni) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowVxlanVni) FromJson(value string) error {
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

func (obj *patternFlowVxlanVni) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowVxlanVni) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowVxlanVni) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowVxlanVni) Clone() (PatternFlowVxlanVni, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowVxlanVni()
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

func (obj *patternFlowVxlanVni) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowVxlanVni is vXLAN network id
type PatternFlowVxlanVni interface {
	Validation
	// msg marshals PatternFlowVxlanVni to protobuf object *otg.PatternFlowVxlanVni
	// and doesn't set defaults
	msg() *otg.PatternFlowVxlanVni
	// setMsg unmarshals PatternFlowVxlanVni from protobuf object *otg.PatternFlowVxlanVni
	// and doesn't set defaults
	setMsg(*otg.PatternFlowVxlanVni) PatternFlowVxlanVni
	// provides marshal interface
	Marshal() marshalPatternFlowVxlanVni
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowVxlanVni
	// validate validates PatternFlowVxlanVni
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowVxlanVni, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowVxlanVniChoiceEnum, set in PatternFlowVxlanVni
	Choice() PatternFlowVxlanVniChoiceEnum
	// setChoice assigns PatternFlowVxlanVniChoiceEnum provided by user to PatternFlowVxlanVni
	setChoice(value PatternFlowVxlanVniChoiceEnum) PatternFlowVxlanVni
	// HasChoice checks if Choice has been set in PatternFlowVxlanVni
	HasChoice() bool
	// Value returns uint32, set in PatternFlowVxlanVni.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowVxlanVni
	SetValue(value uint32) PatternFlowVxlanVni
	// HasValue checks if Value has been set in PatternFlowVxlanVni
	HasValue() bool
	// Values returns []uint32, set in PatternFlowVxlanVni.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowVxlanVni
	SetValues(value []uint32) PatternFlowVxlanVni
	// Auto returns uint32, set in PatternFlowVxlanVni.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowVxlanVni
	HasAuto() bool
	// Increment returns PatternFlowVxlanVniCounter, set in PatternFlowVxlanVni.
	// PatternFlowVxlanVniCounter is integer counter pattern
	Increment() PatternFlowVxlanVniCounter
	// SetIncrement assigns PatternFlowVxlanVniCounter provided by user to PatternFlowVxlanVni.
	// PatternFlowVxlanVniCounter is integer counter pattern
	SetIncrement(value PatternFlowVxlanVniCounter) PatternFlowVxlanVni
	// HasIncrement checks if Increment has been set in PatternFlowVxlanVni
	HasIncrement() bool
	// Decrement returns PatternFlowVxlanVniCounter, set in PatternFlowVxlanVni.
	// PatternFlowVxlanVniCounter is integer counter pattern
	Decrement() PatternFlowVxlanVniCounter
	// SetDecrement assigns PatternFlowVxlanVniCounter provided by user to PatternFlowVxlanVni.
	// PatternFlowVxlanVniCounter is integer counter pattern
	SetDecrement(value PatternFlowVxlanVniCounter) PatternFlowVxlanVni
	// HasDecrement checks if Decrement has been set in PatternFlowVxlanVni
	HasDecrement() bool
	// MetricTags returns PatternFlowVxlanVniPatternFlowVxlanVniMetricTagIterIter, set in PatternFlowVxlanVni
	MetricTags() PatternFlowVxlanVniPatternFlowVxlanVniMetricTagIter
	setNil()
}

type PatternFlowVxlanVniChoiceEnum string

// Enum of Choice on PatternFlowVxlanVni
var PatternFlowVxlanVniChoice = struct {
	VALUE     PatternFlowVxlanVniChoiceEnum
	VALUES    PatternFlowVxlanVniChoiceEnum
	AUTO      PatternFlowVxlanVniChoiceEnum
	INCREMENT PatternFlowVxlanVniChoiceEnum
	DECREMENT PatternFlowVxlanVniChoiceEnum
}{
	VALUE:     PatternFlowVxlanVniChoiceEnum("value"),
	VALUES:    PatternFlowVxlanVniChoiceEnum("values"),
	AUTO:      PatternFlowVxlanVniChoiceEnum("auto"),
	INCREMENT: PatternFlowVxlanVniChoiceEnum("increment"),
	DECREMENT: PatternFlowVxlanVniChoiceEnum("decrement"),
}

func (obj *patternFlowVxlanVni) Choice() PatternFlowVxlanVniChoiceEnum {
	return PatternFlowVxlanVniChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowVxlanVni) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowVxlanVni) setChoice(value PatternFlowVxlanVniChoiceEnum) PatternFlowVxlanVni {
	intValue, ok := otg.PatternFlowVxlanVni_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowVxlanVniChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowVxlanVni_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowVxlanVniChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowVxlanVniChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowVxlanVniChoice.AUTO {
		defaultValue := uint32(0)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowVxlanVniChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowVxlanVniCounter().msg()
	}

	if value == PatternFlowVxlanVniChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowVxlanVniCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowVxlanVni) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowVxlanVniChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowVxlanVni) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowVxlanVni object
func (obj *patternFlowVxlanVni) SetValue(value uint32) PatternFlowVxlanVni {
	obj.setChoice(PatternFlowVxlanVniChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowVxlanVni) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowVxlanVni object
func (obj *patternFlowVxlanVni) SetValues(value []uint32) PatternFlowVxlanVni {
	obj.setChoice(PatternFlowVxlanVniChoice.VALUES)
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
func (obj *patternFlowVxlanVni) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowVxlanVniChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowVxlanVni) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowVxlanVniCounter
func (obj *patternFlowVxlanVni) Increment() PatternFlowVxlanVniCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowVxlanVniChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowVxlanVniCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowVxlanVniCounter
func (obj *patternFlowVxlanVni) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowVxlanVniCounter value in the PatternFlowVxlanVni object
func (obj *patternFlowVxlanVni) SetIncrement(value PatternFlowVxlanVniCounter) PatternFlowVxlanVni {
	obj.setChoice(PatternFlowVxlanVniChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowVxlanVniCounter
func (obj *patternFlowVxlanVni) Decrement() PatternFlowVxlanVniCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowVxlanVniChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowVxlanVniCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowVxlanVniCounter
func (obj *patternFlowVxlanVni) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowVxlanVniCounter value in the PatternFlowVxlanVni object
func (obj *patternFlowVxlanVni) SetDecrement(value PatternFlowVxlanVniCounter) PatternFlowVxlanVni {
	obj.setChoice(PatternFlowVxlanVniChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowVxlanVniMetricTag
func (obj *patternFlowVxlanVni) MetricTags() PatternFlowVxlanVniPatternFlowVxlanVniMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowVxlanVniMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowVxlanVniPatternFlowVxlanVniMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowVxlanVniPatternFlowVxlanVniMetricTagIter struct {
	obj                               *patternFlowVxlanVni
	patternFlowVxlanVniMetricTagSlice []PatternFlowVxlanVniMetricTag
	fieldPtr                          *[]*otg.PatternFlowVxlanVniMetricTag
}

func newPatternFlowVxlanVniPatternFlowVxlanVniMetricTagIter(ptr *[]*otg.PatternFlowVxlanVniMetricTag) PatternFlowVxlanVniPatternFlowVxlanVniMetricTagIter {
	return &patternFlowVxlanVniPatternFlowVxlanVniMetricTagIter{fieldPtr: ptr}
}

type PatternFlowVxlanVniPatternFlowVxlanVniMetricTagIter interface {
	setMsg(*patternFlowVxlanVni) PatternFlowVxlanVniPatternFlowVxlanVniMetricTagIter
	Items() []PatternFlowVxlanVniMetricTag
	Add() PatternFlowVxlanVniMetricTag
	Append(items ...PatternFlowVxlanVniMetricTag) PatternFlowVxlanVniPatternFlowVxlanVniMetricTagIter
	Set(index int, newObj PatternFlowVxlanVniMetricTag) PatternFlowVxlanVniPatternFlowVxlanVniMetricTagIter
	Clear() PatternFlowVxlanVniPatternFlowVxlanVniMetricTagIter
	clearHolderSlice() PatternFlowVxlanVniPatternFlowVxlanVniMetricTagIter
	appendHolderSlice(item PatternFlowVxlanVniMetricTag) PatternFlowVxlanVniPatternFlowVxlanVniMetricTagIter
}

func (obj *patternFlowVxlanVniPatternFlowVxlanVniMetricTagIter) setMsg(msg *patternFlowVxlanVni) PatternFlowVxlanVniPatternFlowVxlanVniMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowVxlanVniMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowVxlanVniPatternFlowVxlanVniMetricTagIter) Items() []PatternFlowVxlanVniMetricTag {
	return obj.patternFlowVxlanVniMetricTagSlice
}

func (obj *patternFlowVxlanVniPatternFlowVxlanVniMetricTagIter) Add() PatternFlowVxlanVniMetricTag {
	newObj := &otg.PatternFlowVxlanVniMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowVxlanVniMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowVxlanVniMetricTagSlice = append(obj.patternFlowVxlanVniMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowVxlanVniPatternFlowVxlanVniMetricTagIter) Append(items ...PatternFlowVxlanVniMetricTag) PatternFlowVxlanVniPatternFlowVxlanVniMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowVxlanVniMetricTagSlice = append(obj.patternFlowVxlanVniMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowVxlanVniPatternFlowVxlanVniMetricTagIter) Set(index int, newObj PatternFlowVxlanVniMetricTag) PatternFlowVxlanVniPatternFlowVxlanVniMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowVxlanVniMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowVxlanVniPatternFlowVxlanVniMetricTagIter) Clear() PatternFlowVxlanVniPatternFlowVxlanVniMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowVxlanVniMetricTag{}
		obj.patternFlowVxlanVniMetricTagSlice = []PatternFlowVxlanVniMetricTag{}
	}
	return obj
}
func (obj *patternFlowVxlanVniPatternFlowVxlanVniMetricTagIter) clearHolderSlice() PatternFlowVxlanVniPatternFlowVxlanVniMetricTagIter {
	if len(obj.patternFlowVxlanVniMetricTagSlice) > 0 {
		obj.patternFlowVxlanVniMetricTagSlice = []PatternFlowVxlanVniMetricTag{}
	}
	return obj
}
func (obj *patternFlowVxlanVniPatternFlowVxlanVniMetricTagIter) appendHolderSlice(item PatternFlowVxlanVniMetricTag) PatternFlowVxlanVniPatternFlowVxlanVniMetricTagIter {
	obj.patternFlowVxlanVniMetricTagSlice = append(obj.patternFlowVxlanVniMetricTagSlice, item)
	return obj
}

func (obj *patternFlowVxlanVni) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVxlanVni.Value <= 16777215 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 16777215 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowVxlanVni.Values <= 16777215 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVxlanVni.Auto <= 16777215 but Got %d", *obj.obj.Auto))
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
				obj.MetricTags().appendHolderSlice(&patternFlowVxlanVniMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowVxlanVni) setDefault() {
	var choices_set int = 0
	var choice PatternFlowVxlanVniChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowVxlanVniChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowVxlanVniChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowVxlanVniChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowVxlanVniChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowVxlanVniChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowVxlanVniChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowVxlanVni")
			}
		} else {
			intVal := otg.PatternFlowVxlanVni_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowVxlanVni_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
