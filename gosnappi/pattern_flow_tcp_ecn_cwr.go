package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpEcnCwr *****
type patternFlowTcpEcnCwr struct {
	validation
	obj              *otg.PatternFlowTcpEcnCwr
	marshaller       marshalPatternFlowTcpEcnCwr
	unMarshaller     unMarshalPatternFlowTcpEcnCwr
	incrementHolder  PatternFlowTcpEcnCwrCounter
	decrementHolder  PatternFlowTcpEcnCwrCounter
	metricTagsHolder PatternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter
}

func NewPatternFlowTcpEcnCwr() PatternFlowTcpEcnCwr {
	obj := patternFlowTcpEcnCwr{obj: &otg.PatternFlowTcpEcnCwr{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpEcnCwr) msg() *otg.PatternFlowTcpEcnCwr {
	return obj.obj
}

func (obj *patternFlowTcpEcnCwr) setMsg(msg *otg.PatternFlowTcpEcnCwr) PatternFlowTcpEcnCwr {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpEcnCwr struct {
	obj *patternFlowTcpEcnCwr
}

type marshalPatternFlowTcpEcnCwr interface {
	// ToProto marshals PatternFlowTcpEcnCwr to protobuf object *otg.PatternFlowTcpEcnCwr
	ToProto() (*otg.PatternFlowTcpEcnCwr, error)
	// ToPbText marshals PatternFlowTcpEcnCwr to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpEcnCwr to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpEcnCwr to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpEcnCwr struct {
	obj *patternFlowTcpEcnCwr
}

type unMarshalPatternFlowTcpEcnCwr interface {
	// FromProto unmarshals PatternFlowTcpEcnCwr from protobuf object *otg.PatternFlowTcpEcnCwr
	FromProto(msg *otg.PatternFlowTcpEcnCwr) (PatternFlowTcpEcnCwr, error)
	// FromPbText unmarshals PatternFlowTcpEcnCwr from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpEcnCwr from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpEcnCwr from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpEcnCwr) Marshal() marshalPatternFlowTcpEcnCwr {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpEcnCwr{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpEcnCwr) Unmarshal() unMarshalPatternFlowTcpEcnCwr {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpEcnCwr{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpEcnCwr) ToProto() (*otg.PatternFlowTcpEcnCwr, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpEcnCwr) FromProto(msg *otg.PatternFlowTcpEcnCwr) (PatternFlowTcpEcnCwr, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpEcnCwr) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpEcnCwr) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpEcnCwr) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpEcnCwr) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpEcnCwr) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpEcnCwr) FromJson(value string) error {
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

func (obj *patternFlowTcpEcnCwr) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpEcnCwr) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpEcnCwr) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpEcnCwr) Clone() (PatternFlowTcpEcnCwr, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpEcnCwr()
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

func (obj *patternFlowTcpEcnCwr) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowTcpEcnCwr is explicit congestion notification, congestion window reduced.
type PatternFlowTcpEcnCwr interface {
	Validation
	// msg marshals PatternFlowTcpEcnCwr to protobuf object *otg.PatternFlowTcpEcnCwr
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpEcnCwr
	// setMsg unmarshals PatternFlowTcpEcnCwr from protobuf object *otg.PatternFlowTcpEcnCwr
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpEcnCwr) PatternFlowTcpEcnCwr
	// provides marshal interface
	Marshal() marshalPatternFlowTcpEcnCwr
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpEcnCwr
	// validate validates PatternFlowTcpEcnCwr
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpEcnCwr, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowTcpEcnCwrChoiceEnum, set in PatternFlowTcpEcnCwr
	Choice() PatternFlowTcpEcnCwrChoiceEnum
	// setChoice assigns PatternFlowTcpEcnCwrChoiceEnum provided by user to PatternFlowTcpEcnCwr
	setChoice(value PatternFlowTcpEcnCwrChoiceEnum) PatternFlowTcpEcnCwr
	// HasChoice checks if Choice has been set in PatternFlowTcpEcnCwr
	HasChoice() bool
	// Value returns uint32, set in PatternFlowTcpEcnCwr.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowTcpEcnCwr
	SetValue(value uint32) PatternFlowTcpEcnCwr
	// HasValue checks if Value has been set in PatternFlowTcpEcnCwr
	HasValue() bool
	// Values returns []uint32, set in PatternFlowTcpEcnCwr.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowTcpEcnCwr
	SetValues(value []uint32) PatternFlowTcpEcnCwr
	// Increment returns PatternFlowTcpEcnCwrCounter, set in PatternFlowTcpEcnCwr.
	// PatternFlowTcpEcnCwrCounter is integer counter pattern
	Increment() PatternFlowTcpEcnCwrCounter
	// SetIncrement assigns PatternFlowTcpEcnCwrCounter provided by user to PatternFlowTcpEcnCwr.
	// PatternFlowTcpEcnCwrCounter is integer counter pattern
	SetIncrement(value PatternFlowTcpEcnCwrCounter) PatternFlowTcpEcnCwr
	// HasIncrement checks if Increment has been set in PatternFlowTcpEcnCwr
	HasIncrement() bool
	// Decrement returns PatternFlowTcpEcnCwrCounter, set in PatternFlowTcpEcnCwr.
	// PatternFlowTcpEcnCwrCounter is integer counter pattern
	Decrement() PatternFlowTcpEcnCwrCounter
	// SetDecrement assigns PatternFlowTcpEcnCwrCounter provided by user to PatternFlowTcpEcnCwr.
	// PatternFlowTcpEcnCwrCounter is integer counter pattern
	SetDecrement(value PatternFlowTcpEcnCwrCounter) PatternFlowTcpEcnCwr
	// HasDecrement checks if Decrement has been set in PatternFlowTcpEcnCwr
	HasDecrement() bool
	// MetricTags returns PatternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIterIter, set in PatternFlowTcpEcnCwr
	MetricTags() PatternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter
	setNil()
}

type PatternFlowTcpEcnCwrChoiceEnum string

// Enum of Choice on PatternFlowTcpEcnCwr
var PatternFlowTcpEcnCwrChoice = struct {
	VALUE     PatternFlowTcpEcnCwrChoiceEnum
	VALUES    PatternFlowTcpEcnCwrChoiceEnum
	INCREMENT PatternFlowTcpEcnCwrChoiceEnum
	DECREMENT PatternFlowTcpEcnCwrChoiceEnum
}{
	VALUE:     PatternFlowTcpEcnCwrChoiceEnum("value"),
	VALUES:    PatternFlowTcpEcnCwrChoiceEnum("values"),
	INCREMENT: PatternFlowTcpEcnCwrChoiceEnum("increment"),
	DECREMENT: PatternFlowTcpEcnCwrChoiceEnum("decrement"),
}

func (obj *patternFlowTcpEcnCwr) Choice() PatternFlowTcpEcnCwrChoiceEnum {
	return PatternFlowTcpEcnCwrChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowTcpEcnCwr) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowTcpEcnCwr) setChoice(value PatternFlowTcpEcnCwrChoiceEnum) PatternFlowTcpEcnCwr {
	intValue, ok := otg.PatternFlowTcpEcnCwr_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowTcpEcnCwrChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowTcpEcnCwr_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowTcpEcnCwrChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowTcpEcnCwrChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowTcpEcnCwrChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowTcpEcnCwrCounter().msg()
	}

	if value == PatternFlowTcpEcnCwrChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowTcpEcnCwrCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpEcnCwr) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowTcpEcnCwrChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpEcnCwr) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowTcpEcnCwr object
func (obj *patternFlowTcpEcnCwr) SetValue(value uint32) PatternFlowTcpEcnCwr {
	obj.setChoice(PatternFlowTcpEcnCwrChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowTcpEcnCwr) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowTcpEcnCwr object
func (obj *patternFlowTcpEcnCwr) SetValues(value []uint32) PatternFlowTcpEcnCwr {
	obj.setChoice(PatternFlowTcpEcnCwrChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowTcpEcnCwrCounter
func (obj *patternFlowTcpEcnCwr) Increment() PatternFlowTcpEcnCwrCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowTcpEcnCwrChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowTcpEcnCwrCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowTcpEcnCwrCounter
func (obj *patternFlowTcpEcnCwr) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowTcpEcnCwrCounter value in the PatternFlowTcpEcnCwr object
func (obj *patternFlowTcpEcnCwr) SetIncrement(value PatternFlowTcpEcnCwrCounter) PatternFlowTcpEcnCwr {
	obj.setChoice(PatternFlowTcpEcnCwrChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowTcpEcnCwrCounter
func (obj *patternFlowTcpEcnCwr) Decrement() PatternFlowTcpEcnCwrCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowTcpEcnCwrChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowTcpEcnCwrCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowTcpEcnCwrCounter
func (obj *patternFlowTcpEcnCwr) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowTcpEcnCwrCounter value in the PatternFlowTcpEcnCwr object
func (obj *patternFlowTcpEcnCwr) SetDecrement(value PatternFlowTcpEcnCwrCounter) PatternFlowTcpEcnCwr {
	obj.setChoice(PatternFlowTcpEcnCwrChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowTcpEcnCwrMetricTag
func (obj *patternFlowTcpEcnCwr) MetricTags() PatternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowTcpEcnCwrMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter struct {
	obj                                *patternFlowTcpEcnCwr
	patternFlowTcpEcnCwrMetricTagSlice []PatternFlowTcpEcnCwrMetricTag
	fieldPtr                           *[]*otg.PatternFlowTcpEcnCwrMetricTag
}

func newPatternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter(ptr *[]*otg.PatternFlowTcpEcnCwrMetricTag) PatternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter {
	return &patternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter{fieldPtr: ptr}
}

type PatternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter interface {
	setMsg(*patternFlowTcpEcnCwr) PatternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter
	Items() []PatternFlowTcpEcnCwrMetricTag
	Add() PatternFlowTcpEcnCwrMetricTag
	Append(items ...PatternFlowTcpEcnCwrMetricTag) PatternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter
	Set(index int, newObj PatternFlowTcpEcnCwrMetricTag) PatternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter
	Clear() PatternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter
	clearHolderSlice() PatternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter
	appendHolderSlice(item PatternFlowTcpEcnCwrMetricTag) PatternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter
}

func (obj *patternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter) setMsg(msg *patternFlowTcpEcnCwr) PatternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowTcpEcnCwrMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter) Items() []PatternFlowTcpEcnCwrMetricTag {
	return obj.patternFlowTcpEcnCwrMetricTagSlice
}

func (obj *patternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter) Add() PatternFlowTcpEcnCwrMetricTag {
	newObj := &otg.PatternFlowTcpEcnCwrMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowTcpEcnCwrMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowTcpEcnCwrMetricTagSlice = append(obj.patternFlowTcpEcnCwrMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter) Append(items ...PatternFlowTcpEcnCwrMetricTag) PatternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowTcpEcnCwrMetricTagSlice = append(obj.patternFlowTcpEcnCwrMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter) Set(index int, newObj PatternFlowTcpEcnCwrMetricTag) PatternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowTcpEcnCwrMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter) Clear() PatternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowTcpEcnCwrMetricTag{}
		obj.patternFlowTcpEcnCwrMetricTagSlice = []PatternFlowTcpEcnCwrMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter) clearHolderSlice() PatternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter {
	if len(obj.patternFlowTcpEcnCwrMetricTagSlice) > 0 {
		obj.patternFlowTcpEcnCwrMetricTagSlice = []PatternFlowTcpEcnCwrMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter) appendHolderSlice(item PatternFlowTcpEcnCwrMetricTag) PatternFlowTcpEcnCwrPatternFlowTcpEcnCwrMetricTagIter {
	obj.patternFlowTcpEcnCwrMetricTagSlice = append(obj.patternFlowTcpEcnCwrMetricTagSlice, item)
	return obj
}

func (obj *patternFlowTcpEcnCwr) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpEcnCwr.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowTcpEcnCwr.Values <= 1 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowTcpEcnCwrMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowTcpEcnCwr) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowTcpEcnCwrChoice.VALUE)

	}

}
