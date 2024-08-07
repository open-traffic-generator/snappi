package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4DontFragment *****
type patternFlowIpv4DontFragment struct {
	validation
	obj              *otg.PatternFlowIpv4DontFragment
	marshaller       marshalPatternFlowIpv4DontFragment
	unMarshaller     unMarshalPatternFlowIpv4DontFragment
	incrementHolder  PatternFlowIpv4DontFragmentCounter
	decrementHolder  PatternFlowIpv4DontFragmentCounter
	metricTagsHolder PatternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter
}

func NewPatternFlowIpv4DontFragment() PatternFlowIpv4DontFragment {
	obj := patternFlowIpv4DontFragment{obj: &otg.PatternFlowIpv4DontFragment{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4DontFragment) msg() *otg.PatternFlowIpv4DontFragment {
	return obj.obj
}

func (obj *patternFlowIpv4DontFragment) setMsg(msg *otg.PatternFlowIpv4DontFragment) PatternFlowIpv4DontFragment {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4DontFragment struct {
	obj *patternFlowIpv4DontFragment
}

type marshalPatternFlowIpv4DontFragment interface {
	// ToProto marshals PatternFlowIpv4DontFragment to protobuf object *otg.PatternFlowIpv4DontFragment
	ToProto() (*otg.PatternFlowIpv4DontFragment, error)
	// ToPbText marshals PatternFlowIpv4DontFragment to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4DontFragment to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4DontFragment to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4DontFragment struct {
	obj *patternFlowIpv4DontFragment
}

type unMarshalPatternFlowIpv4DontFragment interface {
	// FromProto unmarshals PatternFlowIpv4DontFragment from protobuf object *otg.PatternFlowIpv4DontFragment
	FromProto(msg *otg.PatternFlowIpv4DontFragment) (PatternFlowIpv4DontFragment, error)
	// FromPbText unmarshals PatternFlowIpv4DontFragment from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4DontFragment from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4DontFragment from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4DontFragment) Marshal() marshalPatternFlowIpv4DontFragment {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4DontFragment{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4DontFragment) Unmarshal() unMarshalPatternFlowIpv4DontFragment {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4DontFragment{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4DontFragment) ToProto() (*otg.PatternFlowIpv4DontFragment, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4DontFragment) FromProto(msg *otg.PatternFlowIpv4DontFragment) (PatternFlowIpv4DontFragment, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4DontFragment) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DontFragment) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4DontFragment) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DontFragment) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4DontFragment) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DontFragment) FromJson(value string) error {
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

func (obj *patternFlowIpv4DontFragment) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4DontFragment) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4DontFragment) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4DontFragment) Clone() (PatternFlowIpv4DontFragment, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4DontFragment()
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

func (obj *patternFlowIpv4DontFragment) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4DontFragment is dont fragment flag If the dont_fragment flag is set and fragmentation is required to route the packet then the packet is dropped.
type PatternFlowIpv4DontFragment interface {
	Validation
	// msg marshals PatternFlowIpv4DontFragment to protobuf object *otg.PatternFlowIpv4DontFragment
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4DontFragment
	// setMsg unmarshals PatternFlowIpv4DontFragment from protobuf object *otg.PatternFlowIpv4DontFragment
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4DontFragment) PatternFlowIpv4DontFragment
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4DontFragment
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4DontFragment
	// validate validates PatternFlowIpv4DontFragment
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4DontFragment, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4DontFragmentChoiceEnum, set in PatternFlowIpv4DontFragment
	Choice() PatternFlowIpv4DontFragmentChoiceEnum
	// setChoice assigns PatternFlowIpv4DontFragmentChoiceEnum provided by user to PatternFlowIpv4DontFragment
	setChoice(value PatternFlowIpv4DontFragmentChoiceEnum) PatternFlowIpv4DontFragment
	// HasChoice checks if Choice has been set in PatternFlowIpv4DontFragment
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv4DontFragment.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv4DontFragment
	SetValue(value uint32) PatternFlowIpv4DontFragment
	// HasValue checks if Value has been set in PatternFlowIpv4DontFragment
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv4DontFragment.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv4DontFragment
	SetValues(value []uint32) PatternFlowIpv4DontFragment
	// Increment returns PatternFlowIpv4DontFragmentCounter, set in PatternFlowIpv4DontFragment.
	// PatternFlowIpv4DontFragmentCounter is integer counter pattern
	Increment() PatternFlowIpv4DontFragmentCounter
	// SetIncrement assigns PatternFlowIpv4DontFragmentCounter provided by user to PatternFlowIpv4DontFragment.
	// PatternFlowIpv4DontFragmentCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv4DontFragmentCounter) PatternFlowIpv4DontFragment
	// HasIncrement checks if Increment has been set in PatternFlowIpv4DontFragment
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4DontFragmentCounter, set in PatternFlowIpv4DontFragment.
	// PatternFlowIpv4DontFragmentCounter is integer counter pattern
	Decrement() PatternFlowIpv4DontFragmentCounter
	// SetDecrement assigns PatternFlowIpv4DontFragmentCounter provided by user to PatternFlowIpv4DontFragment.
	// PatternFlowIpv4DontFragmentCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv4DontFragmentCounter) PatternFlowIpv4DontFragment
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4DontFragment
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIterIter, set in PatternFlowIpv4DontFragment
	MetricTags() PatternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter
	setNil()
}

type PatternFlowIpv4DontFragmentChoiceEnum string

// Enum of Choice on PatternFlowIpv4DontFragment
var PatternFlowIpv4DontFragmentChoice = struct {
	VALUE     PatternFlowIpv4DontFragmentChoiceEnum
	VALUES    PatternFlowIpv4DontFragmentChoiceEnum
	INCREMENT PatternFlowIpv4DontFragmentChoiceEnum
	DECREMENT PatternFlowIpv4DontFragmentChoiceEnum
}{
	VALUE:     PatternFlowIpv4DontFragmentChoiceEnum("value"),
	VALUES:    PatternFlowIpv4DontFragmentChoiceEnum("values"),
	INCREMENT: PatternFlowIpv4DontFragmentChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4DontFragmentChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4DontFragment) Choice() PatternFlowIpv4DontFragmentChoiceEnum {
	return PatternFlowIpv4DontFragmentChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4DontFragment) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4DontFragment) setChoice(value PatternFlowIpv4DontFragmentChoiceEnum) PatternFlowIpv4DontFragment {
	intValue, ok := otg.PatternFlowIpv4DontFragment_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4DontFragmentChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4DontFragment_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4DontFragmentChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4DontFragmentChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4DontFragmentChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4DontFragmentCounter().msg()
	}

	if value == PatternFlowIpv4DontFragmentChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4DontFragmentCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4DontFragment) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4DontFragmentChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4DontFragment) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv4DontFragment object
func (obj *patternFlowIpv4DontFragment) SetValue(value uint32) PatternFlowIpv4DontFragment {
	obj.setChoice(PatternFlowIpv4DontFragmentChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv4DontFragment) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv4DontFragment object
func (obj *patternFlowIpv4DontFragment) SetValues(value []uint32) PatternFlowIpv4DontFragment {
	obj.setChoice(PatternFlowIpv4DontFragmentChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv4DontFragmentCounter
func (obj *patternFlowIpv4DontFragment) Increment() PatternFlowIpv4DontFragmentCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4DontFragmentChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4DontFragmentCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4DontFragmentCounter
func (obj *patternFlowIpv4DontFragment) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4DontFragmentCounter value in the PatternFlowIpv4DontFragment object
func (obj *patternFlowIpv4DontFragment) SetIncrement(value PatternFlowIpv4DontFragmentCounter) PatternFlowIpv4DontFragment {
	obj.setChoice(PatternFlowIpv4DontFragmentChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4DontFragmentCounter
func (obj *patternFlowIpv4DontFragment) Decrement() PatternFlowIpv4DontFragmentCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4DontFragmentChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4DontFragmentCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4DontFragmentCounter
func (obj *patternFlowIpv4DontFragment) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4DontFragmentCounter value in the PatternFlowIpv4DontFragment object
func (obj *patternFlowIpv4DontFragment) SetDecrement(value PatternFlowIpv4DontFragmentCounter) PatternFlowIpv4DontFragment {
	obj.setChoice(PatternFlowIpv4DontFragmentChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv4DontFragmentMetricTag
func (obj *patternFlowIpv4DontFragment) MetricTags() PatternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv4DontFragmentMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter struct {
	obj                                       *patternFlowIpv4DontFragment
	patternFlowIpv4DontFragmentMetricTagSlice []PatternFlowIpv4DontFragmentMetricTag
	fieldPtr                                  *[]*otg.PatternFlowIpv4DontFragmentMetricTag
}

func newPatternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter(ptr *[]*otg.PatternFlowIpv4DontFragmentMetricTag) PatternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter {
	return &patternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter interface {
	setMsg(*patternFlowIpv4DontFragment) PatternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter
	Items() []PatternFlowIpv4DontFragmentMetricTag
	Add() PatternFlowIpv4DontFragmentMetricTag
	Append(items ...PatternFlowIpv4DontFragmentMetricTag) PatternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter
	Set(index int, newObj PatternFlowIpv4DontFragmentMetricTag) PatternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter
	Clear() PatternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter
	clearHolderSlice() PatternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter
	appendHolderSlice(item PatternFlowIpv4DontFragmentMetricTag) PatternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter
}

func (obj *patternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter) setMsg(msg *patternFlowIpv4DontFragment) PatternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv4DontFragmentMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter) Items() []PatternFlowIpv4DontFragmentMetricTag {
	return obj.patternFlowIpv4DontFragmentMetricTagSlice
}

func (obj *patternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter) Add() PatternFlowIpv4DontFragmentMetricTag {
	newObj := &otg.PatternFlowIpv4DontFragmentMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv4DontFragmentMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv4DontFragmentMetricTagSlice = append(obj.patternFlowIpv4DontFragmentMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter) Append(items ...PatternFlowIpv4DontFragmentMetricTag) PatternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv4DontFragmentMetricTagSlice = append(obj.patternFlowIpv4DontFragmentMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter) Set(index int, newObj PatternFlowIpv4DontFragmentMetricTag) PatternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv4DontFragmentMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter) Clear() PatternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv4DontFragmentMetricTag{}
		obj.patternFlowIpv4DontFragmentMetricTagSlice = []PatternFlowIpv4DontFragmentMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter) clearHolderSlice() PatternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter {
	if len(obj.patternFlowIpv4DontFragmentMetricTagSlice) > 0 {
		obj.patternFlowIpv4DontFragmentMetricTagSlice = []PatternFlowIpv4DontFragmentMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter) appendHolderSlice(item PatternFlowIpv4DontFragmentMetricTag) PatternFlowIpv4DontFragmentPatternFlowIpv4DontFragmentMetricTagIter {
	obj.patternFlowIpv4DontFragmentMetricTagSlice = append(obj.patternFlowIpv4DontFragmentMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIpv4DontFragment) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4DontFragment.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv4DontFragment.Values <= 1 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv4DontFragmentMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIpv4DontFragment) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv4DontFragmentChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv4DontFragmentChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv4DontFragmentChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv4DontFragmentChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv4DontFragmentChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv4DontFragmentChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv4DontFragment")
			}
		} else {
			intVal := otg.PatternFlowIpv4DontFragment_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv4DontFragment_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
