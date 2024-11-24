package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpCtlUrg *****
type patternFlowTcpCtlUrg struct {
	validation
	obj              *otg.PatternFlowTcpCtlUrg
	marshaller       marshalPatternFlowTcpCtlUrg
	unMarshaller     unMarshalPatternFlowTcpCtlUrg
	incrementHolder  PatternFlowTcpCtlUrgCounter
	decrementHolder  PatternFlowTcpCtlUrgCounter
	metricTagsHolder PatternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter
}

func NewPatternFlowTcpCtlUrg() PatternFlowTcpCtlUrg {
	obj := patternFlowTcpCtlUrg{obj: &otg.PatternFlowTcpCtlUrg{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpCtlUrg) msg() *otg.PatternFlowTcpCtlUrg {
	return obj.obj
}

func (obj *patternFlowTcpCtlUrg) setMsg(msg *otg.PatternFlowTcpCtlUrg) PatternFlowTcpCtlUrg {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpCtlUrg struct {
	obj *patternFlowTcpCtlUrg
}

type marshalPatternFlowTcpCtlUrg interface {
	// ToProto marshals PatternFlowTcpCtlUrg to protobuf object *otg.PatternFlowTcpCtlUrg
	ToProto() (*otg.PatternFlowTcpCtlUrg, error)
	// ToPbText marshals PatternFlowTcpCtlUrg to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpCtlUrg to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpCtlUrg to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowTcpCtlUrg to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowTcpCtlUrg struct {
	obj *patternFlowTcpCtlUrg
}

type unMarshalPatternFlowTcpCtlUrg interface {
	// FromProto unmarshals PatternFlowTcpCtlUrg from protobuf object *otg.PatternFlowTcpCtlUrg
	FromProto(msg *otg.PatternFlowTcpCtlUrg) (PatternFlowTcpCtlUrg, error)
	// FromPbText unmarshals PatternFlowTcpCtlUrg from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpCtlUrg from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpCtlUrg from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpCtlUrg) Marshal() marshalPatternFlowTcpCtlUrg {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpCtlUrg{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpCtlUrg) Unmarshal() unMarshalPatternFlowTcpCtlUrg {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpCtlUrg{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpCtlUrg) ToProto() (*otg.PatternFlowTcpCtlUrg, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpCtlUrg) FromProto(msg *otg.PatternFlowTcpCtlUrg) (PatternFlowTcpCtlUrg, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpCtlUrg) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlUrg) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpCtlUrg) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlUrg) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpCtlUrg) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowTcpCtlUrg) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlUrg) FromJson(value string) error {
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

func (obj *patternFlowTcpCtlUrg) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlUrg) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlUrg) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpCtlUrg) Clone() (PatternFlowTcpCtlUrg, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpCtlUrg()
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

func (obj *patternFlowTcpCtlUrg) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowTcpCtlUrg is a value of 1 indicates that the urgent pointer field is significant.
type PatternFlowTcpCtlUrg interface {
	Validation
	// msg marshals PatternFlowTcpCtlUrg to protobuf object *otg.PatternFlowTcpCtlUrg
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpCtlUrg
	// setMsg unmarshals PatternFlowTcpCtlUrg from protobuf object *otg.PatternFlowTcpCtlUrg
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpCtlUrg) PatternFlowTcpCtlUrg
	// provides marshal interface
	Marshal() marshalPatternFlowTcpCtlUrg
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpCtlUrg
	// validate validates PatternFlowTcpCtlUrg
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpCtlUrg, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowTcpCtlUrgChoiceEnum, set in PatternFlowTcpCtlUrg
	Choice() PatternFlowTcpCtlUrgChoiceEnum
	// setChoice assigns PatternFlowTcpCtlUrgChoiceEnum provided by user to PatternFlowTcpCtlUrg
	setChoice(value PatternFlowTcpCtlUrgChoiceEnum) PatternFlowTcpCtlUrg
	// HasChoice checks if Choice has been set in PatternFlowTcpCtlUrg
	HasChoice() bool
	// Value returns uint32, set in PatternFlowTcpCtlUrg.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowTcpCtlUrg
	SetValue(value uint32) PatternFlowTcpCtlUrg
	// HasValue checks if Value has been set in PatternFlowTcpCtlUrg
	HasValue() bool
	// Values returns []uint32, set in PatternFlowTcpCtlUrg.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowTcpCtlUrg
	SetValues(value []uint32) PatternFlowTcpCtlUrg
	// Increment returns PatternFlowTcpCtlUrgCounter, set in PatternFlowTcpCtlUrg.
	// PatternFlowTcpCtlUrgCounter is integer counter pattern
	Increment() PatternFlowTcpCtlUrgCounter
	// SetIncrement assigns PatternFlowTcpCtlUrgCounter provided by user to PatternFlowTcpCtlUrg.
	// PatternFlowTcpCtlUrgCounter is integer counter pattern
	SetIncrement(value PatternFlowTcpCtlUrgCounter) PatternFlowTcpCtlUrg
	// HasIncrement checks if Increment has been set in PatternFlowTcpCtlUrg
	HasIncrement() bool
	// Decrement returns PatternFlowTcpCtlUrgCounter, set in PatternFlowTcpCtlUrg.
	// PatternFlowTcpCtlUrgCounter is integer counter pattern
	Decrement() PatternFlowTcpCtlUrgCounter
	// SetDecrement assigns PatternFlowTcpCtlUrgCounter provided by user to PatternFlowTcpCtlUrg.
	// PatternFlowTcpCtlUrgCounter is integer counter pattern
	SetDecrement(value PatternFlowTcpCtlUrgCounter) PatternFlowTcpCtlUrg
	// HasDecrement checks if Decrement has been set in PatternFlowTcpCtlUrg
	HasDecrement() bool
	// MetricTags returns PatternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIterIter, set in PatternFlowTcpCtlUrg
	MetricTags() PatternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter
	setNil()
}

type PatternFlowTcpCtlUrgChoiceEnum string

// Enum of Choice on PatternFlowTcpCtlUrg
var PatternFlowTcpCtlUrgChoice = struct {
	VALUE     PatternFlowTcpCtlUrgChoiceEnum
	VALUES    PatternFlowTcpCtlUrgChoiceEnum
	INCREMENT PatternFlowTcpCtlUrgChoiceEnum
	DECREMENT PatternFlowTcpCtlUrgChoiceEnum
}{
	VALUE:     PatternFlowTcpCtlUrgChoiceEnum("value"),
	VALUES:    PatternFlowTcpCtlUrgChoiceEnum("values"),
	INCREMENT: PatternFlowTcpCtlUrgChoiceEnum("increment"),
	DECREMENT: PatternFlowTcpCtlUrgChoiceEnum("decrement"),
}

func (obj *patternFlowTcpCtlUrg) Choice() PatternFlowTcpCtlUrgChoiceEnum {
	return PatternFlowTcpCtlUrgChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowTcpCtlUrg) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowTcpCtlUrg) setChoice(value PatternFlowTcpCtlUrgChoiceEnum) PatternFlowTcpCtlUrg {
	intValue, ok := otg.PatternFlowTcpCtlUrg_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowTcpCtlUrgChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowTcpCtlUrg_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowTcpCtlUrgChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowTcpCtlUrgChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowTcpCtlUrgChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowTcpCtlUrgCounter().msg()
	}

	if value == PatternFlowTcpCtlUrgChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowTcpCtlUrgCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpCtlUrg) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowTcpCtlUrgChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpCtlUrg) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowTcpCtlUrg object
func (obj *patternFlowTcpCtlUrg) SetValue(value uint32) PatternFlowTcpCtlUrg {
	obj.setChoice(PatternFlowTcpCtlUrgChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowTcpCtlUrg) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowTcpCtlUrg object
func (obj *patternFlowTcpCtlUrg) SetValues(value []uint32) PatternFlowTcpCtlUrg {
	obj.setChoice(PatternFlowTcpCtlUrgChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowTcpCtlUrgCounter
func (obj *patternFlowTcpCtlUrg) Increment() PatternFlowTcpCtlUrgCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowTcpCtlUrgChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowTcpCtlUrgCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowTcpCtlUrgCounter
func (obj *patternFlowTcpCtlUrg) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowTcpCtlUrgCounter value in the PatternFlowTcpCtlUrg object
func (obj *patternFlowTcpCtlUrg) SetIncrement(value PatternFlowTcpCtlUrgCounter) PatternFlowTcpCtlUrg {
	obj.setChoice(PatternFlowTcpCtlUrgChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowTcpCtlUrgCounter
func (obj *patternFlowTcpCtlUrg) Decrement() PatternFlowTcpCtlUrgCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowTcpCtlUrgChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowTcpCtlUrgCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowTcpCtlUrgCounter
func (obj *patternFlowTcpCtlUrg) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowTcpCtlUrgCounter value in the PatternFlowTcpCtlUrg object
func (obj *patternFlowTcpCtlUrg) SetDecrement(value PatternFlowTcpCtlUrgCounter) PatternFlowTcpCtlUrg {
	obj.setChoice(PatternFlowTcpCtlUrgChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowTcpCtlUrgMetricTag
func (obj *patternFlowTcpCtlUrg) MetricTags() PatternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowTcpCtlUrgMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter struct {
	obj                                *patternFlowTcpCtlUrg
	patternFlowTcpCtlUrgMetricTagSlice []PatternFlowTcpCtlUrgMetricTag
	fieldPtr                           *[]*otg.PatternFlowTcpCtlUrgMetricTag
}

func newPatternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter(ptr *[]*otg.PatternFlowTcpCtlUrgMetricTag) PatternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter {
	return &patternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter{fieldPtr: ptr}
}

type PatternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter interface {
	setMsg(*patternFlowTcpCtlUrg) PatternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter
	Items() []PatternFlowTcpCtlUrgMetricTag
	Add() PatternFlowTcpCtlUrgMetricTag
	Append(items ...PatternFlowTcpCtlUrgMetricTag) PatternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter
	Set(index int, newObj PatternFlowTcpCtlUrgMetricTag) PatternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter
	Clear() PatternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter
	clearHolderSlice() PatternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter
	appendHolderSlice(item PatternFlowTcpCtlUrgMetricTag) PatternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter
}

func (obj *patternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter) setMsg(msg *patternFlowTcpCtlUrg) PatternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowTcpCtlUrgMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter) Items() []PatternFlowTcpCtlUrgMetricTag {
	return obj.patternFlowTcpCtlUrgMetricTagSlice
}

func (obj *patternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter) Add() PatternFlowTcpCtlUrgMetricTag {
	newObj := &otg.PatternFlowTcpCtlUrgMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowTcpCtlUrgMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowTcpCtlUrgMetricTagSlice = append(obj.patternFlowTcpCtlUrgMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter) Append(items ...PatternFlowTcpCtlUrgMetricTag) PatternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowTcpCtlUrgMetricTagSlice = append(obj.patternFlowTcpCtlUrgMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter) Set(index int, newObj PatternFlowTcpCtlUrgMetricTag) PatternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowTcpCtlUrgMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter) Clear() PatternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowTcpCtlUrgMetricTag{}
		obj.patternFlowTcpCtlUrgMetricTagSlice = []PatternFlowTcpCtlUrgMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter) clearHolderSlice() PatternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter {
	if len(obj.patternFlowTcpCtlUrgMetricTagSlice) > 0 {
		obj.patternFlowTcpCtlUrgMetricTagSlice = []PatternFlowTcpCtlUrgMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter) appendHolderSlice(item PatternFlowTcpCtlUrgMetricTag) PatternFlowTcpCtlUrgPatternFlowTcpCtlUrgMetricTagIter {
	obj.patternFlowTcpCtlUrgMetricTagSlice = append(obj.patternFlowTcpCtlUrgMetricTagSlice, item)
	return obj
}

func (obj *patternFlowTcpCtlUrg) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlUrg.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowTcpCtlUrg.Values <= 1 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowTcpCtlUrgMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowTcpCtlUrg) setDefault() {
	var choices_set int = 0
	var choice PatternFlowTcpCtlUrgChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowTcpCtlUrgChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowTcpCtlUrgChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowTcpCtlUrgChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowTcpCtlUrgChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowTcpCtlUrgChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowTcpCtlUrg")
			}
		} else {
			intVal := otg.PatternFlowTcpCtlUrg_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowTcpCtlUrg_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
