package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowVxlanFlags *****
type patternFlowVxlanFlags struct {
	validation
	obj              *otg.PatternFlowVxlanFlags
	marshaller       marshalPatternFlowVxlanFlags
	unMarshaller     unMarshalPatternFlowVxlanFlags
	incrementHolder  PatternFlowVxlanFlagsCounter
	decrementHolder  PatternFlowVxlanFlagsCounter
	metricTagsHolder PatternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter
}

func NewPatternFlowVxlanFlags() PatternFlowVxlanFlags {
	obj := patternFlowVxlanFlags{obj: &otg.PatternFlowVxlanFlags{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowVxlanFlags) msg() *otg.PatternFlowVxlanFlags {
	return obj.obj
}

func (obj *patternFlowVxlanFlags) setMsg(msg *otg.PatternFlowVxlanFlags) PatternFlowVxlanFlags {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowVxlanFlags struct {
	obj *patternFlowVxlanFlags
}

type marshalPatternFlowVxlanFlags interface {
	// ToProto marshals PatternFlowVxlanFlags to protobuf object *otg.PatternFlowVxlanFlags
	ToProto() (*otg.PatternFlowVxlanFlags, error)
	// ToPbText marshals PatternFlowVxlanFlags to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowVxlanFlags to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowVxlanFlags to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowVxlanFlags to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowVxlanFlags struct {
	obj *patternFlowVxlanFlags
}

type unMarshalPatternFlowVxlanFlags interface {
	// FromProto unmarshals PatternFlowVxlanFlags from protobuf object *otg.PatternFlowVxlanFlags
	FromProto(msg *otg.PatternFlowVxlanFlags) (PatternFlowVxlanFlags, error)
	// FromPbText unmarshals PatternFlowVxlanFlags from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowVxlanFlags from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowVxlanFlags from JSON text
	FromJson(value string) error
}

func (obj *patternFlowVxlanFlags) Marshal() marshalPatternFlowVxlanFlags {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowVxlanFlags{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowVxlanFlags) Unmarshal() unMarshalPatternFlowVxlanFlags {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowVxlanFlags{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowVxlanFlags) ToProto() (*otg.PatternFlowVxlanFlags, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowVxlanFlags) FromProto(msg *otg.PatternFlowVxlanFlags) (PatternFlowVxlanFlags, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowVxlanFlags) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowVxlanFlags) FromPbText(value string) error {
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

func (m *marshalpatternFlowVxlanFlags) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowVxlanFlags) FromYaml(value string) error {
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

func (m *marshalpatternFlowVxlanFlags) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowVxlanFlags) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowVxlanFlags) FromJson(value string) error {
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

func (obj *patternFlowVxlanFlags) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowVxlanFlags) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowVxlanFlags) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowVxlanFlags) Clone() (PatternFlowVxlanFlags, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowVxlanFlags()
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

func (obj *patternFlowVxlanFlags) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowVxlanFlags is flags field with a bit format of RRRRIRRR. The I flag MUST be set to 1 for a valid vxlan network id (VNI).   The other 7 bits (designated "R") are reserved fields and MUST be  set to zero on transmission and ignored on receipt.
type PatternFlowVxlanFlags interface {
	Validation
	// msg marshals PatternFlowVxlanFlags to protobuf object *otg.PatternFlowVxlanFlags
	// and doesn't set defaults
	msg() *otg.PatternFlowVxlanFlags
	// setMsg unmarshals PatternFlowVxlanFlags from protobuf object *otg.PatternFlowVxlanFlags
	// and doesn't set defaults
	setMsg(*otg.PatternFlowVxlanFlags) PatternFlowVxlanFlags
	// provides marshal interface
	Marshal() marshalPatternFlowVxlanFlags
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowVxlanFlags
	// validate validates PatternFlowVxlanFlags
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowVxlanFlags, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowVxlanFlagsChoiceEnum, set in PatternFlowVxlanFlags
	Choice() PatternFlowVxlanFlagsChoiceEnum
	// setChoice assigns PatternFlowVxlanFlagsChoiceEnum provided by user to PatternFlowVxlanFlags
	setChoice(value PatternFlowVxlanFlagsChoiceEnum) PatternFlowVxlanFlags
	// HasChoice checks if Choice has been set in PatternFlowVxlanFlags
	HasChoice() bool
	// Value returns uint32, set in PatternFlowVxlanFlags.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowVxlanFlags
	SetValue(value uint32) PatternFlowVxlanFlags
	// HasValue checks if Value has been set in PatternFlowVxlanFlags
	HasValue() bool
	// Values returns []uint32, set in PatternFlowVxlanFlags.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowVxlanFlags
	SetValues(value []uint32) PatternFlowVxlanFlags
	// Increment returns PatternFlowVxlanFlagsCounter, set in PatternFlowVxlanFlags.
	// PatternFlowVxlanFlagsCounter is integer counter pattern
	Increment() PatternFlowVxlanFlagsCounter
	// SetIncrement assigns PatternFlowVxlanFlagsCounter provided by user to PatternFlowVxlanFlags.
	// PatternFlowVxlanFlagsCounter is integer counter pattern
	SetIncrement(value PatternFlowVxlanFlagsCounter) PatternFlowVxlanFlags
	// HasIncrement checks if Increment has been set in PatternFlowVxlanFlags
	HasIncrement() bool
	// Decrement returns PatternFlowVxlanFlagsCounter, set in PatternFlowVxlanFlags.
	// PatternFlowVxlanFlagsCounter is integer counter pattern
	Decrement() PatternFlowVxlanFlagsCounter
	// SetDecrement assigns PatternFlowVxlanFlagsCounter provided by user to PatternFlowVxlanFlags.
	// PatternFlowVxlanFlagsCounter is integer counter pattern
	SetDecrement(value PatternFlowVxlanFlagsCounter) PatternFlowVxlanFlags
	// HasDecrement checks if Decrement has been set in PatternFlowVxlanFlags
	HasDecrement() bool
	// MetricTags returns PatternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIterIter, set in PatternFlowVxlanFlags
	MetricTags() PatternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter
	setNil()
}

type PatternFlowVxlanFlagsChoiceEnum string

// Enum of Choice on PatternFlowVxlanFlags
var PatternFlowVxlanFlagsChoice = struct {
	VALUE     PatternFlowVxlanFlagsChoiceEnum
	VALUES    PatternFlowVxlanFlagsChoiceEnum
	INCREMENT PatternFlowVxlanFlagsChoiceEnum
	DECREMENT PatternFlowVxlanFlagsChoiceEnum
}{
	VALUE:     PatternFlowVxlanFlagsChoiceEnum("value"),
	VALUES:    PatternFlowVxlanFlagsChoiceEnum("values"),
	INCREMENT: PatternFlowVxlanFlagsChoiceEnum("increment"),
	DECREMENT: PatternFlowVxlanFlagsChoiceEnum("decrement"),
}

func (obj *patternFlowVxlanFlags) Choice() PatternFlowVxlanFlagsChoiceEnum {
	return PatternFlowVxlanFlagsChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowVxlanFlags) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowVxlanFlags) setChoice(value PatternFlowVxlanFlagsChoiceEnum) PatternFlowVxlanFlags {
	intValue, ok := otg.PatternFlowVxlanFlags_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowVxlanFlagsChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowVxlanFlags_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowVxlanFlagsChoice.VALUE {
		defaultValue := uint32(8)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowVxlanFlagsChoice.VALUES {
		defaultValue := []uint32{8}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowVxlanFlagsChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowVxlanFlagsCounter().msg()
	}

	if value == PatternFlowVxlanFlagsChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowVxlanFlagsCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowVxlanFlags) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowVxlanFlagsChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowVxlanFlags) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowVxlanFlags object
func (obj *patternFlowVxlanFlags) SetValue(value uint32) PatternFlowVxlanFlags {
	obj.setChoice(PatternFlowVxlanFlagsChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowVxlanFlags) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{8})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowVxlanFlags object
func (obj *patternFlowVxlanFlags) SetValues(value []uint32) PatternFlowVxlanFlags {
	obj.setChoice(PatternFlowVxlanFlagsChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowVxlanFlagsCounter
func (obj *patternFlowVxlanFlags) Increment() PatternFlowVxlanFlagsCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowVxlanFlagsChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowVxlanFlagsCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowVxlanFlagsCounter
func (obj *patternFlowVxlanFlags) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowVxlanFlagsCounter value in the PatternFlowVxlanFlags object
func (obj *patternFlowVxlanFlags) SetIncrement(value PatternFlowVxlanFlagsCounter) PatternFlowVxlanFlags {
	obj.setChoice(PatternFlowVxlanFlagsChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowVxlanFlagsCounter
func (obj *patternFlowVxlanFlags) Decrement() PatternFlowVxlanFlagsCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowVxlanFlagsChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowVxlanFlagsCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowVxlanFlagsCounter
func (obj *patternFlowVxlanFlags) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowVxlanFlagsCounter value in the PatternFlowVxlanFlags object
func (obj *patternFlowVxlanFlags) SetDecrement(value PatternFlowVxlanFlagsCounter) PatternFlowVxlanFlags {
	obj.setChoice(PatternFlowVxlanFlagsChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowVxlanFlagsMetricTag
func (obj *patternFlowVxlanFlags) MetricTags() PatternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowVxlanFlagsMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter struct {
	obj                                 *patternFlowVxlanFlags
	patternFlowVxlanFlagsMetricTagSlice []PatternFlowVxlanFlagsMetricTag
	fieldPtr                            *[]*otg.PatternFlowVxlanFlagsMetricTag
}

func newPatternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter(ptr *[]*otg.PatternFlowVxlanFlagsMetricTag) PatternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter {
	return &patternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter{fieldPtr: ptr}
}

type PatternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter interface {
	setMsg(*patternFlowVxlanFlags) PatternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter
	Items() []PatternFlowVxlanFlagsMetricTag
	Add() PatternFlowVxlanFlagsMetricTag
	Append(items ...PatternFlowVxlanFlagsMetricTag) PatternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter
	Set(index int, newObj PatternFlowVxlanFlagsMetricTag) PatternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter
	Clear() PatternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter
	clearHolderSlice() PatternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter
	appendHolderSlice(item PatternFlowVxlanFlagsMetricTag) PatternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter
}

func (obj *patternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter) setMsg(msg *patternFlowVxlanFlags) PatternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowVxlanFlagsMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter) Items() []PatternFlowVxlanFlagsMetricTag {
	return obj.patternFlowVxlanFlagsMetricTagSlice
}

func (obj *patternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter) Add() PatternFlowVxlanFlagsMetricTag {
	newObj := &otg.PatternFlowVxlanFlagsMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowVxlanFlagsMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowVxlanFlagsMetricTagSlice = append(obj.patternFlowVxlanFlagsMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter) Append(items ...PatternFlowVxlanFlagsMetricTag) PatternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowVxlanFlagsMetricTagSlice = append(obj.patternFlowVxlanFlagsMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter) Set(index int, newObj PatternFlowVxlanFlagsMetricTag) PatternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowVxlanFlagsMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter) Clear() PatternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowVxlanFlagsMetricTag{}
		obj.patternFlowVxlanFlagsMetricTagSlice = []PatternFlowVxlanFlagsMetricTag{}
	}
	return obj
}
func (obj *patternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter) clearHolderSlice() PatternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter {
	if len(obj.patternFlowVxlanFlagsMetricTagSlice) > 0 {
		obj.patternFlowVxlanFlagsMetricTagSlice = []PatternFlowVxlanFlagsMetricTag{}
	}
	return obj
}
func (obj *patternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter) appendHolderSlice(item PatternFlowVxlanFlagsMetricTag) PatternFlowVxlanFlagsPatternFlowVxlanFlagsMetricTagIter {
	obj.patternFlowVxlanFlagsMetricTagSlice = append(obj.patternFlowVxlanFlagsMetricTagSlice, item)
	return obj
}

func (obj *patternFlowVxlanFlags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVxlanFlags.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowVxlanFlags.Values <= 255 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowVxlanFlagsMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowVxlanFlags) setDefault() {
	var choices_set int = 0
	var choice PatternFlowVxlanFlagsChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowVxlanFlagsChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowVxlanFlagsChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowVxlanFlagsChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowVxlanFlagsChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowVxlanFlagsChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowVxlanFlags")
			}
		} else {
			intVal := otg.PatternFlowVxlanFlags_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowVxlanFlags_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
