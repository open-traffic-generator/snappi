package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGreVersion *****
type patternFlowGreVersion struct {
	validation
	obj              *otg.PatternFlowGreVersion
	marshaller       marshalPatternFlowGreVersion
	unMarshaller     unMarshalPatternFlowGreVersion
	incrementHolder  PatternFlowGreVersionCounter
	decrementHolder  PatternFlowGreVersionCounter
	metricTagsHolder PatternFlowGreVersionPatternFlowGreVersionMetricTagIter
}

func NewPatternFlowGreVersion() PatternFlowGreVersion {
	obj := patternFlowGreVersion{obj: &otg.PatternFlowGreVersion{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGreVersion) msg() *otg.PatternFlowGreVersion {
	return obj.obj
}

func (obj *patternFlowGreVersion) setMsg(msg *otg.PatternFlowGreVersion) PatternFlowGreVersion {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGreVersion struct {
	obj *patternFlowGreVersion
}

type marshalPatternFlowGreVersion interface {
	// ToProto marshals PatternFlowGreVersion to protobuf object *otg.PatternFlowGreVersion
	ToProto() (*otg.PatternFlowGreVersion, error)
	// ToPbText marshals PatternFlowGreVersion to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGreVersion to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGreVersion to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowGreVersion to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowGreVersion struct {
	obj *patternFlowGreVersion
}

type unMarshalPatternFlowGreVersion interface {
	// FromProto unmarshals PatternFlowGreVersion from protobuf object *otg.PatternFlowGreVersion
	FromProto(msg *otg.PatternFlowGreVersion) (PatternFlowGreVersion, error)
	// FromPbText unmarshals PatternFlowGreVersion from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGreVersion from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGreVersion from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGreVersion) Marshal() marshalPatternFlowGreVersion {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGreVersion{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGreVersion) Unmarshal() unMarshalPatternFlowGreVersion {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGreVersion{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGreVersion) ToProto() (*otg.PatternFlowGreVersion, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGreVersion) FromProto(msg *otg.PatternFlowGreVersion) (PatternFlowGreVersion, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGreVersion) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGreVersion) FromPbText(value string) error {
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

func (m *marshalpatternFlowGreVersion) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGreVersion) FromYaml(value string) error {
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

func (m *marshalpatternFlowGreVersion) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowGreVersion) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGreVersion) FromJson(value string) error {
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

func (obj *patternFlowGreVersion) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGreVersion) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGreVersion) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGreVersion) Clone() (PatternFlowGreVersion, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGreVersion()
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

func (obj *patternFlowGreVersion) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGreVersion is gRE version number
type PatternFlowGreVersion interface {
	Validation
	// msg marshals PatternFlowGreVersion to protobuf object *otg.PatternFlowGreVersion
	// and doesn't set defaults
	msg() *otg.PatternFlowGreVersion
	// setMsg unmarshals PatternFlowGreVersion from protobuf object *otg.PatternFlowGreVersion
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGreVersion) PatternFlowGreVersion
	// provides marshal interface
	Marshal() marshalPatternFlowGreVersion
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGreVersion
	// validate validates PatternFlowGreVersion
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGreVersion, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGreVersionChoiceEnum, set in PatternFlowGreVersion
	Choice() PatternFlowGreVersionChoiceEnum
	// setChoice assigns PatternFlowGreVersionChoiceEnum provided by user to PatternFlowGreVersion
	setChoice(value PatternFlowGreVersionChoiceEnum) PatternFlowGreVersion
	// HasChoice checks if Choice has been set in PatternFlowGreVersion
	HasChoice() bool
	// Value returns uint32, set in PatternFlowGreVersion.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowGreVersion
	SetValue(value uint32) PatternFlowGreVersion
	// HasValue checks if Value has been set in PatternFlowGreVersion
	HasValue() bool
	// Values returns []uint32, set in PatternFlowGreVersion.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowGreVersion
	SetValues(value []uint32) PatternFlowGreVersion
	// Increment returns PatternFlowGreVersionCounter, set in PatternFlowGreVersion.
	// PatternFlowGreVersionCounter is integer counter pattern
	Increment() PatternFlowGreVersionCounter
	// SetIncrement assigns PatternFlowGreVersionCounter provided by user to PatternFlowGreVersion.
	// PatternFlowGreVersionCounter is integer counter pattern
	SetIncrement(value PatternFlowGreVersionCounter) PatternFlowGreVersion
	// HasIncrement checks if Increment has been set in PatternFlowGreVersion
	HasIncrement() bool
	// Decrement returns PatternFlowGreVersionCounter, set in PatternFlowGreVersion.
	// PatternFlowGreVersionCounter is integer counter pattern
	Decrement() PatternFlowGreVersionCounter
	// SetDecrement assigns PatternFlowGreVersionCounter provided by user to PatternFlowGreVersion.
	// PatternFlowGreVersionCounter is integer counter pattern
	SetDecrement(value PatternFlowGreVersionCounter) PatternFlowGreVersion
	// HasDecrement checks if Decrement has been set in PatternFlowGreVersion
	HasDecrement() bool
	// MetricTags returns PatternFlowGreVersionPatternFlowGreVersionMetricTagIterIter, set in PatternFlowGreVersion
	MetricTags() PatternFlowGreVersionPatternFlowGreVersionMetricTagIter
	setNil()
}

type PatternFlowGreVersionChoiceEnum string

// Enum of Choice on PatternFlowGreVersion
var PatternFlowGreVersionChoice = struct {
	VALUE     PatternFlowGreVersionChoiceEnum
	VALUES    PatternFlowGreVersionChoiceEnum
	INCREMENT PatternFlowGreVersionChoiceEnum
	DECREMENT PatternFlowGreVersionChoiceEnum
}{
	VALUE:     PatternFlowGreVersionChoiceEnum("value"),
	VALUES:    PatternFlowGreVersionChoiceEnum("values"),
	INCREMENT: PatternFlowGreVersionChoiceEnum("increment"),
	DECREMENT: PatternFlowGreVersionChoiceEnum("decrement"),
}

func (obj *patternFlowGreVersion) Choice() PatternFlowGreVersionChoiceEnum {
	return PatternFlowGreVersionChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGreVersion) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGreVersion) setChoice(value PatternFlowGreVersionChoiceEnum) PatternFlowGreVersion {
	intValue, ok := otg.PatternFlowGreVersion_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGreVersionChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGreVersion_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGreVersionChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGreVersionChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGreVersionChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGreVersionCounter().msg()
	}

	if value == PatternFlowGreVersionChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGreVersionCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGreVersion) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGreVersionChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGreVersion) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowGreVersion object
func (obj *patternFlowGreVersion) SetValue(value uint32) PatternFlowGreVersion {
	obj.setChoice(PatternFlowGreVersionChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowGreVersion) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowGreVersion object
func (obj *patternFlowGreVersion) SetValues(value []uint32) PatternFlowGreVersion {
	obj.setChoice(PatternFlowGreVersionChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGreVersionCounter
func (obj *patternFlowGreVersion) Increment() PatternFlowGreVersionCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGreVersionChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGreVersionCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGreVersionCounter
func (obj *patternFlowGreVersion) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGreVersionCounter value in the PatternFlowGreVersion object
func (obj *patternFlowGreVersion) SetIncrement(value PatternFlowGreVersionCounter) PatternFlowGreVersion {
	obj.setChoice(PatternFlowGreVersionChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGreVersionCounter
func (obj *patternFlowGreVersion) Decrement() PatternFlowGreVersionCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGreVersionChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGreVersionCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGreVersionCounter
func (obj *patternFlowGreVersion) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGreVersionCounter value in the PatternFlowGreVersion object
func (obj *patternFlowGreVersion) SetDecrement(value PatternFlowGreVersionCounter) PatternFlowGreVersion {
	obj.setChoice(PatternFlowGreVersionChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGreVersionMetricTag
func (obj *patternFlowGreVersion) MetricTags() PatternFlowGreVersionPatternFlowGreVersionMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGreVersionMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGreVersionPatternFlowGreVersionMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGreVersionPatternFlowGreVersionMetricTagIter struct {
	obj                                 *patternFlowGreVersion
	patternFlowGreVersionMetricTagSlice []PatternFlowGreVersionMetricTag
	fieldPtr                            *[]*otg.PatternFlowGreVersionMetricTag
}

func newPatternFlowGreVersionPatternFlowGreVersionMetricTagIter(ptr *[]*otg.PatternFlowGreVersionMetricTag) PatternFlowGreVersionPatternFlowGreVersionMetricTagIter {
	return &patternFlowGreVersionPatternFlowGreVersionMetricTagIter{fieldPtr: ptr}
}

type PatternFlowGreVersionPatternFlowGreVersionMetricTagIter interface {
	setMsg(*patternFlowGreVersion) PatternFlowGreVersionPatternFlowGreVersionMetricTagIter
	Items() []PatternFlowGreVersionMetricTag
	Add() PatternFlowGreVersionMetricTag
	Append(items ...PatternFlowGreVersionMetricTag) PatternFlowGreVersionPatternFlowGreVersionMetricTagIter
	Set(index int, newObj PatternFlowGreVersionMetricTag) PatternFlowGreVersionPatternFlowGreVersionMetricTagIter
	Clear() PatternFlowGreVersionPatternFlowGreVersionMetricTagIter
	clearHolderSlice() PatternFlowGreVersionPatternFlowGreVersionMetricTagIter
	appendHolderSlice(item PatternFlowGreVersionMetricTag) PatternFlowGreVersionPatternFlowGreVersionMetricTagIter
}

func (obj *patternFlowGreVersionPatternFlowGreVersionMetricTagIter) setMsg(msg *patternFlowGreVersion) PatternFlowGreVersionPatternFlowGreVersionMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGreVersionMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGreVersionPatternFlowGreVersionMetricTagIter) Items() []PatternFlowGreVersionMetricTag {
	return obj.patternFlowGreVersionMetricTagSlice
}

func (obj *patternFlowGreVersionPatternFlowGreVersionMetricTagIter) Add() PatternFlowGreVersionMetricTag {
	newObj := &otg.PatternFlowGreVersionMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGreVersionMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGreVersionMetricTagSlice = append(obj.patternFlowGreVersionMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGreVersionPatternFlowGreVersionMetricTagIter) Append(items ...PatternFlowGreVersionMetricTag) PatternFlowGreVersionPatternFlowGreVersionMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGreVersionMetricTagSlice = append(obj.patternFlowGreVersionMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGreVersionPatternFlowGreVersionMetricTagIter) Set(index int, newObj PatternFlowGreVersionMetricTag) PatternFlowGreVersionPatternFlowGreVersionMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGreVersionMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGreVersionPatternFlowGreVersionMetricTagIter) Clear() PatternFlowGreVersionPatternFlowGreVersionMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGreVersionMetricTag{}
		obj.patternFlowGreVersionMetricTagSlice = []PatternFlowGreVersionMetricTag{}
	}
	return obj
}
func (obj *patternFlowGreVersionPatternFlowGreVersionMetricTagIter) clearHolderSlice() PatternFlowGreVersionPatternFlowGreVersionMetricTagIter {
	if len(obj.patternFlowGreVersionMetricTagSlice) > 0 {
		obj.patternFlowGreVersionMetricTagSlice = []PatternFlowGreVersionMetricTag{}
	}
	return obj
}
func (obj *patternFlowGreVersionPatternFlowGreVersionMetricTagIter) appendHolderSlice(item PatternFlowGreVersionMetricTag) PatternFlowGreVersionPatternFlowGreVersionMetricTagIter {
	obj.patternFlowGreVersionMetricTagSlice = append(obj.patternFlowGreVersionMetricTagSlice, item)
	return obj
}

func (obj *patternFlowGreVersion) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGreVersion.Value <= 7 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 7 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowGreVersion.Values <= 7 but Got %d", item))
			}

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
				obj.MetricTags().appendHolderSlice(&patternFlowGreVersionMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGreVersion) setDefault() {
	var choices_set int = 0
	var choice PatternFlowGreVersionChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowGreVersionChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowGreVersionChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowGreVersionChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowGreVersionChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowGreVersionChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowGreVersion")
			}
		} else {
			intVal := otg.PatternFlowGreVersion_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowGreVersion_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
