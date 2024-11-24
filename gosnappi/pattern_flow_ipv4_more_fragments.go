package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4MoreFragments *****
type patternFlowIpv4MoreFragments struct {
	validation
	obj              *otg.PatternFlowIpv4MoreFragments
	marshaller       marshalPatternFlowIpv4MoreFragments
	unMarshaller     unMarshalPatternFlowIpv4MoreFragments
	incrementHolder  PatternFlowIpv4MoreFragmentsCounter
	decrementHolder  PatternFlowIpv4MoreFragmentsCounter
	metricTagsHolder PatternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter
}

func NewPatternFlowIpv4MoreFragments() PatternFlowIpv4MoreFragments {
	obj := patternFlowIpv4MoreFragments{obj: &otg.PatternFlowIpv4MoreFragments{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4MoreFragments) msg() *otg.PatternFlowIpv4MoreFragments {
	return obj.obj
}

func (obj *patternFlowIpv4MoreFragments) setMsg(msg *otg.PatternFlowIpv4MoreFragments) PatternFlowIpv4MoreFragments {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4MoreFragments struct {
	obj *patternFlowIpv4MoreFragments
}

type marshalPatternFlowIpv4MoreFragments interface {
	// ToProto marshals PatternFlowIpv4MoreFragments to protobuf object *otg.PatternFlowIpv4MoreFragments
	ToProto() (*otg.PatternFlowIpv4MoreFragments, error)
	// ToPbText marshals PatternFlowIpv4MoreFragments to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4MoreFragments to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4MoreFragments to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv4MoreFragments to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv4MoreFragments struct {
	obj *patternFlowIpv4MoreFragments
}

type unMarshalPatternFlowIpv4MoreFragments interface {
	// FromProto unmarshals PatternFlowIpv4MoreFragments from protobuf object *otg.PatternFlowIpv4MoreFragments
	FromProto(msg *otg.PatternFlowIpv4MoreFragments) (PatternFlowIpv4MoreFragments, error)
	// FromPbText unmarshals PatternFlowIpv4MoreFragments from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4MoreFragments from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4MoreFragments from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4MoreFragments) Marshal() marshalPatternFlowIpv4MoreFragments {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4MoreFragments{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4MoreFragments) Unmarshal() unMarshalPatternFlowIpv4MoreFragments {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4MoreFragments{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4MoreFragments) ToProto() (*otg.PatternFlowIpv4MoreFragments, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4MoreFragments) FromProto(msg *otg.PatternFlowIpv4MoreFragments) (PatternFlowIpv4MoreFragments, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4MoreFragments) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4MoreFragments) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4MoreFragments) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4MoreFragments) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4MoreFragments) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv4MoreFragments) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4MoreFragments) FromJson(value string) error {
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

func (obj *patternFlowIpv4MoreFragments) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4MoreFragments) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4MoreFragments) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4MoreFragments) Clone() (PatternFlowIpv4MoreFragments, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4MoreFragments()
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

func (obj *patternFlowIpv4MoreFragments) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4MoreFragments is more fragments flag
type PatternFlowIpv4MoreFragments interface {
	Validation
	// msg marshals PatternFlowIpv4MoreFragments to protobuf object *otg.PatternFlowIpv4MoreFragments
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4MoreFragments
	// setMsg unmarshals PatternFlowIpv4MoreFragments from protobuf object *otg.PatternFlowIpv4MoreFragments
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4MoreFragments) PatternFlowIpv4MoreFragments
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4MoreFragments
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4MoreFragments
	// validate validates PatternFlowIpv4MoreFragments
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4MoreFragments, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4MoreFragmentsChoiceEnum, set in PatternFlowIpv4MoreFragments
	Choice() PatternFlowIpv4MoreFragmentsChoiceEnum
	// setChoice assigns PatternFlowIpv4MoreFragmentsChoiceEnum provided by user to PatternFlowIpv4MoreFragments
	setChoice(value PatternFlowIpv4MoreFragmentsChoiceEnum) PatternFlowIpv4MoreFragments
	// HasChoice checks if Choice has been set in PatternFlowIpv4MoreFragments
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv4MoreFragments.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv4MoreFragments
	SetValue(value uint32) PatternFlowIpv4MoreFragments
	// HasValue checks if Value has been set in PatternFlowIpv4MoreFragments
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv4MoreFragments.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv4MoreFragments
	SetValues(value []uint32) PatternFlowIpv4MoreFragments
	// Increment returns PatternFlowIpv4MoreFragmentsCounter, set in PatternFlowIpv4MoreFragments.
	// PatternFlowIpv4MoreFragmentsCounter is integer counter pattern
	Increment() PatternFlowIpv4MoreFragmentsCounter
	// SetIncrement assigns PatternFlowIpv4MoreFragmentsCounter provided by user to PatternFlowIpv4MoreFragments.
	// PatternFlowIpv4MoreFragmentsCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv4MoreFragmentsCounter) PatternFlowIpv4MoreFragments
	// HasIncrement checks if Increment has been set in PatternFlowIpv4MoreFragments
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4MoreFragmentsCounter, set in PatternFlowIpv4MoreFragments.
	// PatternFlowIpv4MoreFragmentsCounter is integer counter pattern
	Decrement() PatternFlowIpv4MoreFragmentsCounter
	// SetDecrement assigns PatternFlowIpv4MoreFragmentsCounter provided by user to PatternFlowIpv4MoreFragments.
	// PatternFlowIpv4MoreFragmentsCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv4MoreFragmentsCounter) PatternFlowIpv4MoreFragments
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4MoreFragments
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIterIter, set in PatternFlowIpv4MoreFragments
	MetricTags() PatternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter
	setNil()
}

type PatternFlowIpv4MoreFragmentsChoiceEnum string

// Enum of Choice on PatternFlowIpv4MoreFragments
var PatternFlowIpv4MoreFragmentsChoice = struct {
	VALUE     PatternFlowIpv4MoreFragmentsChoiceEnum
	VALUES    PatternFlowIpv4MoreFragmentsChoiceEnum
	INCREMENT PatternFlowIpv4MoreFragmentsChoiceEnum
	DECREMENT PatternFlowIpv4MoreFragmentsChoiceEnum
}{
	VALUE:     PatternFlowIpv4MoreFragmentsChoiceEnum("value"),
	VALUES:    PatternFlowIpv4MoreFragmentsChoiceEnum("values"),
	INCREMENT: PatternFlowIpv4MoreFragmentsChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4MoreFragmentsChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4MoreFragments) Choice() PatternFlowIpv4MoreFragmentsChoiceEnum {
	return PatternFlowIpv4MoreFragmentsChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4MoreFragments) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4MoreFragments) setChoice(value PatternFlowIpv4MoreFragmentsChoiceEnum) PatternFlowIpv4MoreFragments {
	intValue, ok := otg.PatternFlowIpv4MoreFragments_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4MoreFragmentsChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4MoreFragments_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4MoreFragmentsChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4MoreFragmentsChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4MoreFragmentsChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4MoreFragmentsCounter().msg()
	}

	if value == PatternFlowIpv4MoreFragmentsChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4MoreFragmentsCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4MoreFragments) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4MoreFragmentsChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4MoreFragments) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv4MoreFragments object
func (obj *patternFlowIpv4MoreFragments) SetValue(value uint32) PatternFlowIpv4MoreFragments {
	obj.setChoice(PatternFlowIpv4MoreFragmentsChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv4MoreFragments) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv4MoreFragments object
func (obj *patternFlowIpv4MoreFragments) SetValues(value []uint32) PatternFlowIpv4MoreFragments {
	obj.setChoice(PatternFlowIpv4MoreFragmentsChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv4MoreFragmentsCounter
func (obj *patternFlowIpv4MoreFragments) Increment() PatternFlowIpv4MoreFragmentsCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4MoreFragmentsChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4MoreFragmentsCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4MoreFragmentsCounter
func (obj *patternFlowIpv4MoreFragments) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4MoreFragmentsCounter value in the PatternFlowIpv4MoreFragments object
func (obj *patternFlowIpv4MoreFragments) SetIncrement(value PatternFlowIpv4MoreFragmentsCounter) PatternFlowIpv4MoreFragments {
	obj.setChoice(PatternFlowIpv4MoreFragmentsChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4MoreFragmentsCounter
func (obj *patternFlowIpv4MoreFragments) Decrement() PatternFlowIpv4MoreFragmentsCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4MoreFragmentsChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4MoreFragmentsCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4MoreFragmentsCounter
func (obj *patternFlowIpv4MoreFragments) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4MoreFragmentsCounter value in the PatternFlowIpv4MoreFragments object
func (obj *patternFlowIpv4MoreFragments) SetDecrement(value PatternFlowIpv4MoreFragmentsCounter) PatternFlowIpv4MoreFragments {
	obj.setChoice(PatternFlowIpv4MoreFragmentsChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv4MoreFragmentsMetricTag
func (obj *patternFlowIpv4MoreFragments) MetricTags() PatternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv4MoreFragmentsMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter struct {
	obj                                        *patternFlowIpv4MoreFragments
	patternFlowIpv4MoreFragmentsMetricTagSlice []PatternFlowIpv4MoreFragmentsMetricTag
	fieldPtr                                   *[]*otg.PatternFlowIpv4MoreFragmentsMetricTag
}

func newPatternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter(ptr *[]*otg.PatternFlowIpv4MoreFragmentsMetricTag) PatternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter {
	return &patternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter interface {
	setMsg(*patternFlowIpv4MoreFragments) PatternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter
	Items() []PatternFlowIpv4MoreFragmentsMetricTag
	Add() PatternFlowIpv4MoreFragmentsMetricTag
	Append(items ...PatternFlowIpv4MoreFragmentsMetricTag) PatternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter
	Set(index int, newObj PatternFlowIpv4MoreFragmentsMetricTag) PatternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter
	Clear() PatternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter
	clearHolderSlice() PatternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter
	appendHolderSlice(item PatternFlowIpv4MoreFragmentsMetricTag) PatternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter
}

func (obj *patternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter) setMsg(msg *patternFlowIpv4MoreFragments) PatternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv4MoreFragmentsMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter) Items() []PatternFlowIpv4MoreFragmentsMetricTag {
	return obj.patternFlowIpv4MoreFragmentsMetricTagSlice
}

func (obj *patternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter) Add() PatternFlowIpv4MoreFragmentsMetricTag {
	newObj := &otg.PatternFlowIpv4MoreFragmentsMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv4MoreFragmentsMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv4MoreFragmentsMetricTagSlice = append(obj.patternFlowIpv4MoreFragmentsMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter) Append(items ...PatternFlowIpv4MoreFragmentsMetricTag) PatternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv4MoreFragmentsMetricTagSlice = append(obj.patternFlowIpv4MoreFragmentsMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter) Set(index int, newObj PatternFlowIpv4MoreFragmentsMetricTag) PatternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv4MoreFragmentsMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter) Clear() PatternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv4MoreFragmentsMetricTag{}
		obj.patternFlowIpv4MoreFragmentsMetricTagSlice = []PatternFlowIpv4MoreFragmentsMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter) clearHolderSlice() PatternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter {
	if len(obj.patternFlowIpv4MoreFragmentsMetricTagSlice) > 0 {
		obj.patternFlowIpv4MoreFragmentsMetricTagSlice = []PatternFlowIpv4MoreFragmentsMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter) appendHolderSlice(item PatternFlowIpv4MoreFragmentsMetricTag) PatternFlowIpv4MoreFragmentsPatternFlowIpv4MoreFragmentsMetricTagIter {
	obj.patternFlowIpv4MoreFragmentsMetricTagSlice = append(obj.patternFlowIpv4MoreFragmentsMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIpv4MoreFragments) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4MoreFragments.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv4MoreFragments.Values <= 1 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv4MoreFragmentsMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIpv4MoreFragments) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv4MoreFragmentsChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv4MoreFragmentsChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv4MoreFragmentsChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv4MoreFragmentsChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv4MoreFragmentsChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv4MoreFragmentsChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv4MoreFragments")
			}
		} else {
			intVal := otg.PatternFlowIpv4MoreFragments_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv4MoreFragments_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
