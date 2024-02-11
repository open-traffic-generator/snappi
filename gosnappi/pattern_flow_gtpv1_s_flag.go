package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1SFlag *****
type patternFlowGtpv1SFlag struct {
	validation
	obj              *otg.PatternFlowGtpv1SFlag
	marshaller       marshalPatternFlowGtpv1SFlag
	unMarshaller     unMarshalPatternFlowGtpv1SFlag
	incrementHolder  PatternFlowGtpv1SFlagCounter
	decrementHolder  PatternFlowGtpv1SFlagCounter
	metricTagsHolder PatternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter
}

func NewPatternFlowGtpv1SFlag() PatternFlowGtpv1SFlag {
	obj := patternFlowGtpv1SFlag{obj: &otg.PatternFlowGtpv1SFlag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1SFlag) msg() *otg.PatternFlowGtpv1SFlag {
	return obj.obj
}

func (obj *patternFlowGtpv1SFlag) setMsg(msg *otg.PatternFlowGtpv1SFlag) PatternFlowGtpv1SFlag {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1SFlag struct {
	obj *patternFlowGtpv1SFlag
}

type marshalPatternFlowGtpv1SFlag interface {
	// ToProto marshals PatternFlowGtpv1SFlag to protobuf object *otg.PatternFlowGtpv1SFlag
	ToProto() (*otg.PatternFlowGtpv1SFlag, error)
	// ToPbText marshals PatternFlowGtpv1SFlag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1SFlag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1SFlag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv1SFlag struct {
	obj *patternFlowGtpv1SFlag
}

type unMarshalPatternFlowGtpv1SFlag interface {
	// FromProto unmarshals PatternFlowGtpv1SFlag from protobuf object *otg.PatternFlowGtpv1SFlag
	FromProto(msg *otg.PatternFlowGtpv1SFlag) (PatternFlowGtpv1SFlag, error)
	// FromPbText unmarshals PatternFlowGtpv1SFlag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1SFlag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1SFlag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1SFlag) Marshal() marshalPatternFlowGtpv1SFlag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1SFlag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1SFlag) Unmarshal() unMarshalPatternFlowGtpv1SFlag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1SFlag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1SFlag) ToProto() (*otg.PatternFlowGtpv1SFlag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1SFlag) FromProto(msg *otg.PatternFlowGtpv1SFlag) (PatternFlowGtpv1SFlag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1SFlag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1SFlag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1SFlag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1SFlag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1SFlag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1SFlag) FromJson(value string) error {
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

func (obj *patternFlowGtpv1SFlag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1SFlag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1SFlag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1SFlag) Clone() (PatternFlowGtpv1SFlag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1SFlag()
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

func (obj *patternFlowGtpv1SFlag) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGtpv1SFlag is sequence number field present
type PatternFlowGtpv1SFlag interface {
	Validation
	// msg marshals PatternFlowGtpv1SFlag to protobuf object *otg.PatternFlowGtpv1SFlag
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1SFlag
	// setMsg unmarshals PatternFlowGtpv1SFlag from protobuf object *otg.PatternFlowGtpv1SFlag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1SFlag) PatternFlowGtpv1SFlag
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1SFlag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1SFlag
	// validate validates PatternFlowGtpv1SFlag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1SFlag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGtpv1SFlagChoiceEnum, set in PatternFlowGtpv1SFlag
	Choice() PatternFlowGtpv1SFlagChoiceEnum
	// setChoice assigns PatternFlowGtpv1SFlagChoiceEnum provided by user to PatternFlowGtpv1SFlag
	setChoice(value PatternFlowGtpv1SFlagChoiceEnum) PatternFlowGtpv1SFlag
	// HasChoice checks if Choice has been set in PatternFlowGtpv1SFlag
	HasChoice() bool
	// Value returns uint32, set in PatternFlowGtpv1SFlag.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowGtpv1SFlag
	SetValue(value uint32) PatternFlowGtpv1SFlag
	// HasValue checks if Value has been set in PatternFlowGtpv1SFlag
	HasValue() bool
	// Values returns []uint32, set in PatternFlowGtpv1SFlag.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowGtpv1SFlag
	SetValues(value []uint32) PatternFlowGtpv1SFlag
	// Increment returns PatternFlowGtpv1SFlagCounter, set in PatternFlowGtpv1SFlag.
	// PatternFlowGtpv1SFlagCounter is integer counter pattern
	Increment() PatternFlowGtpv1SFlagCounter
	// SetIncrement assigns PatternFlowGtpv1SFlagCounter provided by user to PatternFlowGtpv1SFlag.
	// PatternFlowGtpv1SFlagCounter is integer counter pattern
	SetIncrement(value PatternFlowGtpv1SFlagCounter) PatternFlowGtpv1SFlag
	// HasIncrement checks if Increment has been set in PatternFlowGtpv1SFlag
	HasIncrement() bool
	// Decrement returns PatternFlowGtpv1SFlagCounter, set in PatternFlowGtpv1SFlag.
	// PatternFlowGtpv1SFlagCounter is integer counter pattern
	Decrement() PatternFlowGtpv1SFlagCounter
	// SetDecrement assigns PatternFlowGtpv1SFlagCounter provided by user to PatternFlowGtpv1SFlag.
	// PatternFlowGtpv1SFlagCounter is integer counter pattern
	SetDecrement(value PatternFlowGtpv1SFlagCounter) PatternFlowGtpv1SFlag
	// HasDecrement checks if Decrement has been set in PatternFlowGtpv1SFlag
	HasDecrement() bool
	// MetricTags returns PatternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIterIter, set in PatternFlowGtpv1SFlag
	MetricTags() PatternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter
	setNil()
}

type PatternFlowGtpv1SFlagChoiceEnum string

// Enum of Choice on PatternFlowGtpv1SFlag
var PatternFlowGtpv1SFlagChoice = struct {
	VALUE     PatternFlowGtpv1SFlagChoiceEnum
	VALUES    PatternFlowGtpv1SFlagChoiceEnum
	INCREMENT PatternFlowGtpv1SFlagChoiceEnum
	DECREMENT PatternFlowGtpv1SFlagChoiceEnum
}{
	VALUE:     PatternFlowGtpv1SFlagChoiceEnum("value"),
	VALUES:    PatternFlowGtpv1SFlagChoiceEnum("values"),
	INCREMENT: PatternFlowGtpv1SFlagChoiceEnum("increment"),
	DECREMENT: PatternFlowGtpv1SFlagChoiceEnum("decrement"),
}

func (obj *patternFlowGtpv1SFlag) Choice() PatternFlowGtpv1SFlagChoiceEnum {
	return PatternFlowGtpv1SFlagChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGtpv1SFlag) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGtpv1SFlag) setChoice(value PatternFlowGtpv1SFlagChoiceEnum) PatternFlowGtpv1SFlag {
	intValue, ok := otg.PatternFlowGtpv1SFlag_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGtpv1SFlagChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGtpv1SFlag_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGtpv1SFlagChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGtpv1SFlagChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGtpv1SFlagChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGtpv1SFlagCounter().msg()
	}

	if value == PatternFlowGtpv1SFlagChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGtpv1SFlagCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv1SFlag) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGtpv1SFlagChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv1SFlag) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowGtpv1SFlag object
func (obj *patternFlowGtpv1SFlag) SetValue(value uint32) PatternFlowGtpv1SFlag {
	obj.setChoice(PatternFlowGtpv1SFlagChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowGtpv1SFlag) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowGtpv1SFlag object
func (obj *patternFlowGtpv1SFlag) SetValues(value []uint32) PatternFlowGtpv1SFlag {
	obj.setChoice(PatternFlowGtpv1SFlagChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGtpv1SFlagCounter
func (obj *patternFlowGtpv1SFlag) Increment() PatternFlowGtpv1SFlagCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGtpv1SFlagChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGtpv1SFlagCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGtpv1SFlagCounter
func (obj *patternFlowGtpv1SFlag) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGtpv1SFlagCounter value in the PatternFlowGtpv1SFlag object
func (obj *patternFlowGtpv1SFlag) SetIncrement(value PatternFlowGtpv1SFlagCounter) PatternFlowGtpv1SFlag {
	obj.setChoice(PatternFlowGtpv1SFlagChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGtpv1SFlagCounter
func (obj *patternFlowGtpv1SFlag) Decrement() PatternFlowGtpv1SFlagCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGtpv1SFlagChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGtpv1SFlagCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGtpv1SFlagCounter
func (obj *patternFlowGtpv1SFlag) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGtpv1SFlagCounter value in the PatternFlowGtpv1SFlag object
func (obj *patternFlowGtpv1SFlag) SetDecrement(value PatternFlowGtpv1SFlagCounter) PatternFlowGtpv1SFlag {
	obj.setChoice(PatternFlowGtpv1SFlagChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGtpv1SFlagMetricTag
func (obj *patternFlowGtpv1SFlag) MetricTags() PatternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGtpv1SFlagMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter struct {
	obj                                 *patternFlowGtpv1SFlag
	patternFlowGtpv1SFlagMetricTagSlice []PatternFlowGtpv1SFlagMetricTag
	fieldPtr                            *[]*otg.PatternFlowGtpv1SFlagMetricTag
}

func newPatternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter(ptr *[]*otg.PatternFlowGtpv1SFlagMetricTag) PatternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter {
	return &patternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter{fieldPtr: ptr}
}

type PatternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter interface {
	setMsg(*patternFlowGtpv1SFlag) PatternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter
	Items() []PatternFlowGtpv1SFlagMetricTag
	Add() PatternFlowGtpv1SFlagMetricTag
	Append(items ...PatternFlowGtpv1SFlagMetricTag) PatternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter
	Set(index int, newObj PatternFlowGtpv1SFlagMetricTag) PatternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter
	Clear() PatternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter
	clearHolderSlice() PatternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter
	appendHolderSlice(item PatternFlowGtpv1SFlagMetricTag) PatternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter
}

func (obj *patternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter) setMsg(msg *patternFlowGtpv1SFlag) PatternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGtpv1SFlagMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter) Items() []PatternFlowGtpv1SFlagMetricTag {
	return obj.patternFlowGtpv1SFlagMetricTagSlice
}

func (obj *patternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter) Add() PatternFlowGtpv1SFlagMetricTag {
	newObj := &otg.PatternFlowGtpv1SFlagMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGtpv1SFlagMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGtpv1SFlagMetricTagSlice = append(obj.patternFlowGtpv1SFlagMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter) Append(items ...PatternFlowGtpv1SFlagMetricTag) PatternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGtpv1SFlagMetricTagSlice = append(obj.patternFlowGtpv1SFlagMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter) Set(index int, newObj PatternFlowGtpv1SFlagMetricTag) PatternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGtpv1SFlagMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter) Clear() PatternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGtpv1SFlagMetricTag{}
		obj.patternFlowGtpv1SFlagMetricTagSlice = []PatternFlowGtpv1SFlagMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter) clearHolderSlice() PatternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter {
	if len(obj.patternFlowGtpv1SFlagMetricTagSlice) > 0 {
		obj.patternFlowGtpv1SFlagMetricTagSlice = []PatternFlowGtpv1SFlagMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter) appendHolderSlice(item PatternFlowGtpv1SFlagMetricTag) PatternFlowGtpv1SFlagPatternFlowGtpv1SFlagMetricTagIter {
	obj.patternFlowGtpv1SFlagMetricTagSlice = append(obj.patternFlowGtpv1SFlagMetricTagSlice, item)
	return obj
}

func (obj *patternFlowGtpv1SFlag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1SFlag.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowGtpv1SFlag.Values <= 1 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowGtpv1SFlagMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGtpv1SFlag) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowGtpv1SFlagChoice.VALUE)

	}

}
