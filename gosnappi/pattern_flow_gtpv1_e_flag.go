package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1EFlag *****
type patternFlowGtpv1EFlag struct {
	validation
	obj              *otg.PatternFlowGtpv1EFlag
	marshaller       marshalPatternFlowGtpv1EFlag
	unMarshaller     unMarshalPatternFlowGtpv1EFlag
	incrementHolder  PatternFlowGtpv1EFlagCounter
	decrementHolder  PatternFlowGtpv1EFlagCounter
	metricTagsHolder PatternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter
}

func NewPatternFlowGtpv1EFlag() PatternFlowGtpv1EFlag {
	obj := patternFlowGtpv1EFlag{obj: &otg.PatternFlowGtpv1EFlag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1EFlag) msg() *otg.PatternFlowGtpv1EFlag {
	return obj.obj
}

func (obj *patternFlowGtpv1EFlag) setMsg(msg *otg.PatternFlowGtpv1EFlag) PatternFlowGtpv1EFlag {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1EFlag struct {
	obj *patternFlowGtpv1EFlag
}

type marshalPatternFlowGtpv1EFlag interface {
	// ToProto marshals PatternFlowGtpv1EFlag to protobuf object *otg.PatternFlowGtpv1EFlag
	ToProto() (*otg.PatternFlowGtpv1EFlag, error)
	// ToPbText marshals PatternFlowGtpv1EFlag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1EFlag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1EFlag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv1EFlag struct {
	obj *patternFlowGtpv1EFlag
}

type unMarshalPatternFlowGtpv1EFlag interface {
	// FromProto unmarshals PatternFlowGtpv1EFlag from protobuf object *otg.PatternFlowGtpv1EFlag
	FromProto(msg *otg.PatternFlowGtpv1EFlag) (PatternFlowGtpv1EFlag, error)
	// FromPbText unmarshals PatternFlowGtpv1EFlag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1EFlag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1EFlag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1EFlag) Marshal() marshalPatternFlowGtpv1EFlag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1EFlag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1EFlag) Unmarshal() unMarshalPatternFlowGtpv1EFlag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1EFlag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1EFlag) ToProto() (*otg.PatternFlowGtpv1EFlag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1EFlag) FromProto(msg *otg.PatternFlowGtpv1EFlag) (PatternFlowGtpv1EFlag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1EFlag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1EFlag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1EFlag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1EFlag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1EFlag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1EFlag) FromJson(value string) error {
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

func (obj *patternFlowGtpv1EFlag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1EFlag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1EFlag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1EFlag) Clone() (PatternFlowGtpv1EFlag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1EFlag()
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

func (obj *patternFlowGtpv1EFlag) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGtpv1EFlag is extension header field present
type PatternFlowGtpv1EFlag interface {
	Validation
	// msg marshals PatternFlowGtpv1EFlag to protobuf object *otg.PatternFlowGtpv1EFlag
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1EFlag
	// setMsg unmarshals PatternFlowGtpv1EFlag from protobuf object *otg.PatternFlowGtpv1EFlag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1EFlag) PatternFlowGtpv1EFlag
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1EFlag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1EFlag
	// validate validates PatternFlowGtpv1EFlag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1EFlag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGtpv1EFlagChoiceEnum, set in PatternFlowGtpv1EFlag
	Choice() PatternFlowGtpv1EFlagChoiceEnum
	// setChoice assigns PatternFlowGtpv1EFlagChoiceEnum provided by user to PatternFlowGtpv1EFlag
	setChoice(value PatternFlowGtpv1EFlagChoiceEnum) PatternFlowGtpv1EFlag
	// HasChoice checks if Choice has been set in PatternFlowGtpv1EFlag
	HasChoice() bool
	// Value returns uint32, set in PatternFlowGtpv1EFlag.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowGtpv1EFlag
	SetValue(value uint32) PatternFlowGtpv1EFlag
	// HasValue checks if Value has been set in PatternFlowGtpv1EFlag
	HasValue() bool
	// Values returns []uint32, set in PatternFlowGtpv1EFlag.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowGtpv1EFlag
	SetValues(value []uint32) PatternFlowGtpv1EFlag
	// Increment returns PatternFlowGtpv1EFlagCounter, set in PatternFlowGtpv1EFlag.
	// PatternFlowGtpv1EFlagCounter is integer counter pattern
	Increment() PatternFlowGtpv1EFlagCounter
	// SetIncrement assigns PatternFlowGtpv1EFlagCounter provided by user to PatternFlowGtpv1EFlag.
	// PatternFlowGtpv1EFlagCounter is integer counter pattern
	SetIncrement(value PatternFlowGtpv1EFlagCounter) PatternFlowGtpv1EFlag
	// HasIncrement checks if Increment has been set in PatternFlowGtpv1EFlag
	HasIncrement() bool
	// Decrement returns PatternFlowGtpv1EFlagCounter, set in PatternFlowGtpv1EFlag.
	// PatternFlowGtpv1EFlagCounter is integer counter pattern
	Decrement() PatternFlowGtpv1EFlagCounter
	// SetDecrement assigns PatternFlowGtpv1EFlagCounter provided by user to PatternFlowGtpv1EFlag.
	// PatternFlowGtpv1EFlagCounter is integer counter pattern
	SetDecrement(value PatternFlowGtpv1EFlagCounter) PatternFlowGtpv1EFlag
	// HasDecrement checks if Decrement has been set in PatternFlowGtpv1EFlag
	HasDecrement() bool
	// MetricTags returns PatternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIterIter, set in PatternFlowGtpv1EFlag
	MetricTags() PatternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter
	setNil()
}

type PatternFlowGtpv1EFlagChoiceEnum string

// Enum of Choice on PatternFlowGtpv1EFlag
var PatternFlowGtpv1EFlagChoice = struct {
	VALUE     PatternFlowGtpv1EFlagChoiceEnum
	VALUES    PatternFlowGtpv1EFlagChoiceEnum
	INCREMENT PatternFlowGtpv1EFlagChoiceEnum
	DECREMENT PatternFlowGtpv1EFlagChoiceEnum
}{
	VALUE:     PatternFlowGtpv1EFlagChoiceEnum("value"),
	VALUES:    PatternFlowGtpv1EFlagChoiceEnum("values"),
	INCREMENT: PatternFlowGtpv1EFlagChoiceEnum("increment"),
	DECREMENT: PatternFlowGtpv1EFlagChoiceEnum("decrement"),
}

func (obj *patternFlowGtpv1EFlag) Choice() PatternFlowGtpv1EFlagChoiceEnum {
	return PatternFlowGtpv1EFlagChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGtpv1EFlag) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGtpv1EFlag) setChoice(value PatternFlowGtpv1EFlagChoiceEnum) PatternFlowGtpv1EFlag {
	intValue, ok := otg.PatternFlowGtpv1EFlag_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGtpv1EFlagChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGtpv1EFlag_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGtpv1EFlagChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGtpv1EFlagChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGtpv1EFlagChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGtpv1EFlagCounter().msg()
	}

	if value == PatternFlowGtpv1EFlagChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGtpv1EFlagCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv1EFlag) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGtpv1EFlagChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv1EFlag) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowGtpv1EFlag object
func (obj *patternFlowGtpv1EFlag) SetValue(value uint32) PatternFlowGtpv1EFlag {
	obj.setChoice(PatternFlowGtpv1EFlagChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowGtpv1EFlag) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowGtpv1EFlag object
func (obj *patternFlowGtpv1EFlag) SetValues(value []uint32) PatternFlowGtpv1EFlag {
	obj.setChoice(PatternFlowGtpv1EFlagChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGtpv1EFlagCounter
func (obj *patternFlowGtpv1EFlag) Increment() PatternFlowGtpv1EFlagCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGtpv1EFlagChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGtpv1EFlagCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGtpv1EFlagCounter
func (obj *patternFlowGtpv1EFlag) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGtpv1EFlagCounter value in the PatternFlowGtpv1EFlag object
func (obj *patternFlowGtpv1EFlag) SetIncrement(value PatternFlowGtpv1EFlagCounter) PatternFlowGtpv1EFlag {
	obj.setChoice(PatternFlowGtpv1EFlagChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGtpv1EFlagCounter
func (obj *patternFlowGtpv1EFlag) Decrement() PatternFlowGtpv1EFlagCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGtpv1EFlagChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGtpv1EFlagCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGtpv1EFlagCounter
func (obj *patternFlowGtpv1EFlag) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGtpv1EFlagCounter value in the PatternFlowGtpv1EFlag object
func (obj *patternFlowGtpv1EFlag) SetDecrement(value PatternFlowGtpv1EFlagCounter) PatternFlowGtpv1EFlag {
	obj.setChoice(PatternFlowGtpv1EFlagChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGtpv1EFlagMetricTag
func (obj *patternFlowGtpv1EFlag) MetricTags() PatternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGtpv1EFlagMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter struct {
	obj                                 *patternFlowGtpv1EFlag
	patternFlowGtpv1EFlagMetricTagSlice []PatternFlowGtpv1EFlagMetricTag
	fieldPtr                            *[]*otg.PatternFlowGtpv1EFlagMetricTag
}

func newPatternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter(ptr *[]*otg.PatternFlowGtpv1EFlagMetricTag) PatternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter {
	return &patternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter{fieldPtr: ptr}
}

type PatternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter interface {
	setMsg(*patternFlowGtpv1EFlag) PatternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter
	Items() []PatternFlowGtpv1EFlagMetricTag
	Add() PatternFlowGtpv1EFlagMetricTag
	Append(items ...PatternFlowGtpv1EFlagMetricTag) PatternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter
	Set(index int, newObj PatternFlowGtpv1EFlagMetricTag) PatternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter
	Clear() PatternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter
	clearHolderSlice() PatternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter
	appendHolderSlice(item PatternFlowGtpv1EFlagMetricTag) PatternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter
}

func (obj *patternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter) setMsg(msg *patternFlowGtpv1EFlag) PatternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGtpv1EFlagMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter) Items() []PatternFlowGtpv1EFlagMetricTag {
	return obj.patternFlowGtpv1EFlagMetricTagSlice
}

func (obj *patternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter) Add() PatternFlowGtpv1EFlagMetricTag {
	newObj := &otg.PatternFlowGtpv1EFlagMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGtpv1EFlagMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGtpv1EFlagMetricTagSlice = append(obj.patternFlowGtpv1EFlagMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter) Append(items ...PatternFlowGtpv1EFlagMetricTag) PatternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGtpv1EFlagMetricTagSlice = append(obj.patternFlowGtpv1EFlagMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter) Set(index int, newObj PatternFlowGtpv1EFlagMetricTag) PatternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGtpv1EFlagMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter) Clear() PatternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGtpv1EFlagMetricTag{}
		obj.patternFlowGtpv1EFlagMetricTagSlice = []PatternFlowGtpv1EFlagMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter) clearHolderSlice() PatternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter {
	if len(obj.patternFlowGtpv1EFlagMetricTagSlice) > 0 {
		obj.patternFlowGtpv1EFlagMetricTagSlice = []PatternFlowGtpv1EFlagMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter) appendHolderSlice(item PatternFlowGtpv1EFlagMetricTag) PatternFlowGtpv1EFlagPatternFlowGtpv1EFlagMetricTagIter {
	obj.patternFlowGtpv1EFlagMetricTagSlice = append(obj.patternFlowGtpv1EFlagMetricTagSlice, item)
	return obj
}

func (obj *patternFlowGtpv1EFlag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1EFlag.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowGtpv1EFlag.Values <= 1 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowGtpv1EFlagMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGtpv1EFlag) setDefault() {
	var choices_set int = 0
	var choice PatternFlowGtpv1EFlagChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowGtpv1EFlagChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowGtpv1EFlagChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowGtpv1EFlagChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowGtpv1EFlagChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowGtpv1EFlagChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowGtpv1EFlag")
			}
		} else {
			intVal := otg.PatternFlowGtpv1EFlag_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowGtpv1EFlag_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
