package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv2TeidFlag *****
type patternFlowGtpv2TeidFlag struct {
	validation
	obj              *otg.PatternFlowGtpv2TeidFlag
	marshaller       marshalPatternFlowGtpv2TeidFlag
	unMarshaller     unMarshalPatternFlowGtpv2TeidFlag
	incrementHolder  PatternFlowGtpv2TeidFlagCounter
	decrementHolder  PatternFlowGtpv2TeidFlagCounter
	metricTagsHolder PatternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter
}

func NewPatternFlowGtpv2TeidFlag() PatternFlowGtpv2TeidFlag {
	obj := patternFlowGtpv2TeidFlag{obj: &otg.PatternFlowGtpv2TeidFlag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv2TeidFlag) msg() *otg.PatternFlowGtpv2TeidFlag {
	return obj.obj
}

func (obj *patternFlowGtpv2TeidFlag) setMsg(msg *otg.PatternFlowGtpv2TeidFlag) PatternFlowGtpv2TeidFlag {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv2TeidFlag struct {
	obj *patternFlowGtpv2TeidFlag
}

type marshalPatternFlowGtpv2TeidFlag interface {
	// ToProto marshals PatternFlowGtpv2TeidFlag to protobuf object *otg.PatternFlowGtpv2TeidFlag
	ToProto() (*otg.PatternFlowGtpv2TeidFlag, error)
	// ToPbText marshals PatternFlowGtpv2TeidFlag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv2TeidFlag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv2TeidFlag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv2TeidFlag struct {
	obj *patternFlowGtpv2TeidFlag
}

type unMarshalPatternFlowGtpv2TeidFlag interface {
	// FromProto unmarshals PatternFlowGtpv2TeidFlag from protobuf object *otg.PatternFlowGtpv2TeidFlag
	FromProto(msg *otg.PatternFlowGtpv2TeidFlag) (PatternFlowGtpv2TeidFlag, error)
	// FromPbText unmarshals PatternFlowGtpv2TeidFlag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv2TeidFlag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv2TeidFlag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv2TeidFlag) Marshal() marshalPatternFlowGtpv2TeidFlag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv2TeidFlag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv2TeidFlag) Unmarshal() unMarshalPatternFlowGtpv2TeidFlag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv2TeidFlag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv2TeidFlag) ToProto() (*otg.PatternFlowGtpv2TeidFlag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv2TeidFlag) FromProto(msg *otg.PatternFlowGtpv2TeidFlag) (PatternFlowGtpv2TeidFlag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv2TeidFlag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2TeidFlag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv2TeidFlag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2TeidFlag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv2TeidFlag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2TeidFlag) FromJson(value string) error {
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

func (obj *patternFlowGtpv2TeidFlag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2TeidFlag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2TeidFlag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv2TeidFlag) Clone() (PatternFlowGtpv2TeidFlag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv2TeidFlag()
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

func (obj *patternFlowGtpv2TeidFlag) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGtpv2TeidFlag is if teid_flag is set to 1 then the TEID field will be present  between the message length and the sequence number. All messages except Echo and Echo reply require TEID to be present
type PatternFlowGtpv2TeidFlag interface {
	Validation
	// msg marshals PatternFlowGtpv2TeidFlag to protobuf object *otg.PatternFlowGtpv2TeidFlag
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv2TeidFlag
	// setMsg unmarshals PatternFlowGtpv2TeidFlag from protobuf object *otg.PatternFlowGtpv2TeidFlag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv2TeidFlag) PatternFlowGtpv2TeidFlag
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv2TeidFlag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv2TeidFlag
	// validate validates PatternFlowGtpv2TeidFlag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv2TeidFlag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGtpv2TeidFlagChoiceEnum, set in PatternFlowGtpv2TeidFlag
	Choice() PatternFlowGtpv2TeidFlagChoiceEnum
	// setChoice assigns PatternFlowGtpv2TeidFlagChoiceEnum provided by user to PatternFlowGtpv2TeidFlag
	setChoice(value PatternFlowGtpv2TeidFlagChoiceEnum) PatternFlowGtpv2TeidFlag
	// HasChoice checks if Choice has been set in PatternFlowGtpv2TeidFlag
	HasChoice() bool
	// Value returns uint32, set in PatternFlowGtpv2TeidFlag.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowGtpv2TeidFlag
	SetValue(value uint32) PatternFlowGtpv2TeidFlag
	// HasValue checks if Value has been set in PatternFlowGtpv2TeidFlag
	HasValue() bool
	// Values returns []uint32, set in PatternFlowGtpv2TeidFlag.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowGtpv2TeidFlag
	SetValues(value []uint32) PatternFlowGtpv2TeidFlag
	// Increment returns PatternFlowGtpv2TeidFlagCounter, set in PatternFlowGtpv2TeidFlag.
	// PatternFlowGtpv2TeidFlagCounter is integer counter pattern
	Increment() PatternFlowGtpv2TeidFlagCounter
	// SetIncrement assigns PatternFlowGtpv2TeidFlagCounter provided by user to PatternFlowGtpv2TeidFlag.
	// PatternFlowGtpv2TeidFlagCounter is integer counter pattern
	SetIncrement(value PatternFlowGtpv2TeidFlagCounter) PatternFlowGtpv2TeidFlag
	// HasIncrement checks if Increment has been set in PatternFlowGtpv2TeidFlag
	HasIncrement() bool
	// Decrement returns PatternFlowGtpv2TeidFlagCounter, set in PatternFlowGtpv2TeidFlag.
	// PatternFlowGtpv2TeidFlagCounter is integer counter pattern
	Decrement() PatternFlowGtpv2TeidFlagCounter
	// SetDecrement assigns PatternFlowGtpv2TeidFlagCounter provided by user to PatternFlowGtpv2TeidFlag.
	// PatternFlowGtpv2TeidFlagCounter is integer counter pattern
	SetDecrement(value PatternFlowGtpv2TeidFlagCounter) PatternFlowGtpv2TeidFlag
	// HasDecrement checks if Decrement has been set in PatternFlowGtpv2TeidFlag
	HasDecrement() bool
	// MetricTags returns PatternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIterIter, set in PatternFlowGtpv2TeidFlag
	MetricTags() PatternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter
	setNil()
}

type PatternFlowGtpv2TeidFlagChoiceEnum string

// Enum of Choice on PatternFlowGtpv2TeidFlag
var PatternFlowGtpv2TeidFlagChoice = struct {
	VALUE     PatternFlowGtpv2TeidFlagChoiceEnum
	VALUES    PatternFlowGtpv2TeidFlagChoiceEnum
	INCREMENT PatternFlowGtpv2TeidFlagChoiceEnum
	DECREMENT PatternFlowGtpv2TeidFlagChoiceEnum
}{
	VALUE:     PatternFlowGtpv2TeidFlagChoiceEnum("value"),
	VALUES:    PatternFlowGtpv2TeidFlagChoiceEnum("values"),
	INCREMENT: PatternFlowGtpv2TeidFlagChoiceEnum("increment"),
	DECREMENT: PatternFlowGtpv2TeidFlagChoiceEnum("decrement"),
}

func (obj *patternFlowGtpv2TeidFlag) Choice() PatternFlowGtpv2TeidFlagChoiceEnum {
	return PatternFlowGtpv2TeidFlagChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGtpv2TeidFlag) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGtpv2TeidFlag) setChoice(value PatternFlowGtpv2TeidFlagChoiceEnum) PatternFlowGtpv2TeidFlag {
	intValue, ok := otg.PatternFlowGtpv2TeidFlag_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGtpv2TeidFlagChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGtpv2TeidFlag_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGtpv2TeidFlagChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGtpv2TeidFlagChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGtpv2TeidFlagChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGtpv2TeidFlagCounter().msg()
	}

	if value == PatternFlowGtpv2TeidFlagChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGtpv2TeidFlagCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv2TeidFlag) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGtpv2TeidFlagChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv2TeidFlag) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowGtpv2TeidFlag object
func (obj *patternFlowGtpv2TeidFlag) SetValue(value uint32) PatternFlowGtpv2TeidFlag {
	obj.setChoice(PatternFlowGtpv2TeidFlagChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowGtpv2TeidFlag) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowGtpv2TeidFlag object
func (obj *patternFlowGtpv2TeidFlag) SetValues(value []uint32) PatternFlowGtpv2TeidFlag {
	obj.setChoice(PatternFlowGtpv2TeidFlagChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGtpv2TeidFlagCounter
func (obj *patternFlowGtpv2TeidFlag) Increment() PatternFlowGtpv2TeidFlagCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGtpv2TeidFlagChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGtpv2TeidFlagCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGtpv2TeidFlagCounter
func (obj *patternFlowGtpv2TeidFlag) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGtpv2TeidFlagCounter value in the PatternFlowGtpv2TeidFlag object
func (obj *patternFlowGtpv2TeidFlag) SetIncrement(value PatternFlowGtpv2TeidFlagCounter) PatternFlowGtpv2TeidFlag {
	obj.setChoice(PatternFlowGtpv2TeidFlagChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGtpv2TeidFlagCounter
func (obj *patternFlowGtpv2TeidFlag) Decrement() PatternFlowGtpv2TeidFlagCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGtpv2TeidFlagChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGtpv2TeidFlagCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGtpv2TeidFlagCounter
func (obj *patternFlowGtpv2TeidFlag) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGtpv2TeidFlagCounter value in the PatternFlowGtpv2TeidFlag object
func (obj *patternFlowGtpv2TeidFlag) SetDecrement(value PatternFlowGtpv2TeidFlagCounter) PatternFlowGtpv2TeidFlag {
	obj.setChoice(PatternFlowGtpv2TeidFlagChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGtpv2TeidFlagMetricTag
func (obj *patternFlowGtpv2TeidFlag) MetricTags() PatternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGtpv2TeidFlagMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter struct {
	obj                                    *patternFlowGtpv2TeidFlag
	patternFlowGtpv2TeidFlagMetricTagSlice []PatternFlowGtpv2TeidFlagMetricTag
	fieldPtr                               *[]*otg.PatternFlowGtpv2TeidFlagMetricTag
}

func newPatternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter(ptr *[]*otg.PatternFlowGtpv2TeidFlagMetricTag) PatternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter {
	return &patternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter{fieldPtr: ptr}
}

type PatternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter interface {
	setMsg(*patternFlowGtpv2TeidFlag) PatternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter
	Items() []PatternFlowGtpv2TeidFlagMetricTag
	Add() PatternFlowGtpv2TeidFlagMetricTag
	Append(items ...PatternFlowGtpv2TeidFlagMetricTag) PatternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter
	Set(index int, newObj PatternFlowGtpv2TeidFlagMetricTag) PatternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter
	Clear() PatternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter
	clearHolderSlice() PatternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter
	appendHolderSlice(item PatternFlowGtpv2TeidFlagMetricTag) PatternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter
}

func (obj *patternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter) setMsg(msg *patternFlowGtpv2TeidFlag) PatternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGtpv2TeidFlagMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter) Items() []PatternFlowGtpv2TeidFlagMetricTag {
	return obj.patternFlowGtpv2TeidFlagMetricTagSlice
}

func (obj *patternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter) Add() PatternFlowGtpv2TeidFlagMetricTag {
	newObj := &otg.PatternFlowGtpv2TeidFlagMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGtpv2TeidFlagMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGtpv2TeidFlagMetricTagSlice = append(obj.patternFlowGtpv2TeidFlagMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter) Append(items ...PatternFlowGtpv2TeidFlagMetricTag) PatternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGtpv2TeidFlagMetricTagSlice = append(obj.patternFlowGtpv2TeidFlagMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter) Set(index int, newObj PatternFlowGtpv2TeidFlagMetricTag) PatternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGtpv2TeidFlagMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter) Clear() PatternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGtpv2TeidFlagMetricTag{}
		obj.patternFlowGtpv2TeidFlagMetricTagSlice = []PatternFlowGtpv2TeidFlagMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter) clearHolderSlice() PatternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter {
	if len(obj.patternFlowGtpv2TeidFlagMetricTagSlice) > 0 {
		obj.patternFlowGtpv2TeidFlagMetricTagSlice = []PatternFlowGtpv2TeidFlagMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter) appendHolderSlice(item PatternFlowGtpv2TeidFlagMetricTag) PatternFlowGtpv2TeidFlagPatternFlowGtpv2TeidFlagMetricTagIter {
	obj.patternFlowGtpv2TeidFlagMetricTagSlice = append(obj.patternFlowGtpv2TeidFlagMetricTagSlice, item)
	return obj
}

func (obj *patternFlowGtpv2TeidFlag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2TeidFlag.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowGtpv2TeidFlag.Values <= 1 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowGtpv2TeidFlagMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGtpv2TeidFlag) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowGtpv2TeidFlagChoice.VALUE)

	}

}
