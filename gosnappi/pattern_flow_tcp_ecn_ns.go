package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpEcnNs *****
type patternFlowTcpEcnNs struct {
	validation
	obj              *otg.PatternFlowTcpEcnNs
	marshaller       marshalPatternFlowTcpEcnNs
	unMarshaller     unMarshalPatternFlowTcpEcnNs
	incrementHolder  PatternFlowTcpEcnNsCounter
	decrementHolder  PatternFlowTcpEcnNsCounter
	metricTagsHolder PatternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter
}

func NewPatternFlowTcpEcnNs() PatternFlowTcpEcnNs {
	obj := patternFlowTcpEcnNs{obj: &otg.PatternFlowTcpEcnNs{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpEcnNs) msg() *otg.PatternFlowTcpEcnNs {
	return obj.obj
}

func (obj *patternFlowTcpEcnNs) setMsg(msg *otg.PatternFlowTcpEcnNs) PatternFlowTcpEcnNs {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpEcnNs struct {
	obj *patternFlowTcpEcnNs
}

type marshalPatternFlowTcpEcnNs interface {
	// ToProto marshals PatternFlowTcpEcnNs to protobuf object *otg.PatternFlowTcpEcnNs
	ToProto() (*otg.PatternFlowTcpEcnNs, error)
	// ToPbText marshals PatternFlowTcpEcnNs to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpEcnNs to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpEcnNs to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowTcpEcnNs to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowTcpEcnNs struct {
	obj *patternFlowTcpEcnNs
}

type unMarshalPatternFlowTcpEcnNs interface {
	// FromProto unmarshals PatternFlowTcpEcnNs from protobuf object *otg.PatternFlowTcpEcnNs
	FromProto(msg *otg.PatternFlowTcpEcnNs) (PatternFlowTcpEcnNs, error)
	// FromPbText unmarshals PatternFlowTcpEcnNs from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpEcnNs from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpEcnNs from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpEcnNs) Marshal() marshalPatternFlowTcpEcnNs {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpEcnNs{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpEcnNs) Unmarshal() unMarshalPatternFlowTcpEcnNs {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpEcnNs{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpEcnNs) ToProto() (*otg.PatternFlowTcpEcnNs, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpEcnNs) FromProto(msg *otg.PatternFlowTcpEcnNs) (PatternFlowTcpEcnNs, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpEcnNs) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpEcnNs) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpEcnNs) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpEcnNs) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpEcnNs) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowTcpEcnNs) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpEcnNs) FromJson(value string) error {
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

func (obj *patternFlowTcpEcnNs) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpEcnNs) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpEcnNs) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpEcnNs) Clone() (PatternFlowTcpEcnNs, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpEcnNs()
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

func (obj *patternFlowTcpEcnNs) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowTcpEcnNs is explicit congestion notification, concealment protection.
type PatternFlowTcpEcnNs interface {
	Validation
	// msg marshals PatternFlowTcpEcnNs to protobuf object *otg.PatternFlowTcpEcnNs
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpEcnNs
	// setMsg unmarshals PatternFlowTcpEcnNs from protobuf object *otg.PatternFlowTcpEcnNs
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpEcnNs) PatternFlowTcpEcnNs
	// provides marshal interface
	Marshal() marshalPatternFlowTcpEcnNs
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpEcnNs
	// validate validates PatternFlowTcpEcnNs
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpEcnNs, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowTcpEcnNsChoiceEnum, set in PatternFlowTcpEcnNs
	Choice() PatternFlowTcpEcnNsChoiceEnum
	// setChoice assigns PatternFlowTcpEcnNsChoiceEnum provided by user to PatternFlowTcpEcnNs
	setChoice(value PatternFlowTcpEcnNsChoiceEnum) PatternFlowTcpEcnNs
	// HasChoice checks if Choice has been set in PatternFlowTcpEcnNs
	HasChoice() bool
	// Value returns uint32, set in PatternFlowTcpEcnNs.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowTcpEcnNs
	SetValue(value uint32) PatternFlowTcpEcnNs
	// HasValue checks if Value has been set in PatternFlowTcpEcnNs
	HasValue() bool
	// Values returns []uint32, set in PatternFlowTcpEcnNs.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowTcpEcnNs
	SetValues(value []uint32) PatternFlowTcpEcnNs
	// Increment returns PatternFlowTcpEcnNsCounter, set in PatternFlowTcpEcnNs.
	// PatternFlowTcpEcnNsCounter is integer counter pattern
	Increment() PatternFlowTcpEcnNsCounter
	// SetIncrement assigns PatternFlowTcpEcnNsCounter provided by user to PatternFlowTcpEcnNs.
	// PatternFlowTcpEcnNsCounter is integer counter pattern
	SetIncrement(value PatternFlowTcpEcnNsCounter) PatternFlowTcpEcnNs
	// HasIncrement checks if Increment has been set in PatternFlowTcpEcnNs
	HasIncrement() bool
	// Decrement returns PatternFlowTcpEcnNsCounter, set in PatternFlowTcpEcnNs.
	// PatternFlowTcpEcnNsCounter is integer counter pattern
	Decrement() PatternFlowTcpEcnNsCounter
	// SetDecrement assigns PatternFlowTcpEcnNsCounter provided by user to PatternFlowTcpEcnNs.
	// PatternFlowTcpEcnNsCounter is integer counter pattern
	SetDecrement(value PatternFlowTcpEcnNsCounter) PatternFlowTcpEcnNs
	// HasDecrement checks if Decrement has been set in PatternFlowTcpEcnNs
	HasDecrement() bool
	// MetricTags returns PatternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIterIter, set in PatternFlowTcpEcnNs
	MetricTags() PatternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter
	setNil()
}

type PatternFlowTcpEcnNsChoiceEnum string

// Enum of Choice on PatternFlowTcpEcnNs
var PatternFlowTcpEcnNsChoice = struct {
	VALUE     PatternFlowTcpEcnNsChoiceEnum
	VALUES    PatternFlowTcpEcnNsChoiceEnum
	INCREMENT PatternFlowTcpEcnNsChoiceEnum
	DECREMENT PatternFlowTcpEcnNsChoiceEnum
}{
	VALUE:     PatternFlowTcpEcnNsChoiceEnum("value"),
	VALUES:    PatternFlowTcpEcnNsChoiceEnum("values"),
	INCREMENT: PatternFlowTcpEcnNsChoiceEnum("increment"),
	DECREMENT: PatternFlowTcpEcnNsChoiceEnum("decrement"),
}

func (obj *patternFlowTcpEcnNs) Choice() PatternFlowTcpEcnNsChoiceEnum {
	return PatternFlowTcpEcnNsChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowTcpEcnNs) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowTcpEcnNs) setChoice(value PatternFlowTcpEcnNsChoiceEnum) PatternFlowTcpEcnNs {
	intValue, ok := otg.PatternFlowTcpEcnNs_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowTcpEcnNsChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowTcpEcnNs_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowTcpEcnNsChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowTcpEcnNsChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowTcpEcnNsChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowTcpEcnNsCounter().msg()
	}

	if value == PatternFlowTcpEcnNsChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowTcpEcnNsCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpEcnNs) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowTcpEcnNsChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpEcnNs) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowTcpEcnNs object
func (obj *patternFlowTcpEcnNs) SetValue(value uint32) PatternFlowTcpEcnNs {
	obj.setChoice(PatternFlowTcpEcnNsChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowTcpEcnNs) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowTcpEcnNs object
func (obj *patternFlowTcpEcnNs) SetValues(value []uint32) PatternFlowTcpEcnNs {
	obj.setChoice(PatternFlowTcpEcnNsChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowTcpEcnNsCounter
func (obj *patternFlowTcpEcnNs) Increment() PatternFlowTcpEcnNsCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowTcpEcnNsChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowTcpEcnNsCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowTcpEcnNsCounter
func (obj *patternFlowTcpEcnNs) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowTcpEcnNsCounter value in the PatternFlowTcpEcnNs object
func (obj *patternFlowTcpEcnNs) SetIncrement(value PatternFlowTcpEcnNsCounter) PatternFlowTcpEcnNs {
	obj.setChoice(PatternFlowTcpEcnNsChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowTcpEcnNsCounter
func (obj *patternFlowTcpEcnNs) Decrement() PatternFlowTcpEcnNsCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowTcpEcnNsChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowTcpEcnNsCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowTcpEcnNsCounter
func (obj *patternFlowTcpEcnNs) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowTcpEcnNsCounter value in the PatternFlowTcpEcnNs object
func (obj *patternFlowTcpEcnNs) SetDecrement(value PatternFlowTcpEcnNsCounter) PatternFlowTcpEcnNs {
	obj.setChoice(PatternFlowTcpEcnNsChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowTcpEcnNsMetricTag
func (obj *patternFlowTcpEcnNs) MetricTags() PatternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowTcpEcnNsMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter struct {
	obj                               *patternFlowTcpEcnNs
	patternFlowTcpEcnNsMetricTagSlice []PatternFlowTcpEcnNsMetricTag
	fieldPtr                          *[]*otg.PatternFlowTcpEcnNsMetricTag
}

func newPatternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter(ptr *[]*otg.PatternFlowTcpEcnNsMetricTag) PatternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter {
	return &patternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter{fieldPtr: ptr}
}

type PatternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter interface {
	setMsg(*patternFlowTcpEcnNs) PatternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter
	Items() []PatternFlowTcpEcnNsMetricTag
	Add() PatternFlowTcpEcnNsMetricTag
	Append(items ...PatternFlowTcpEcnNsMetricTag) PatternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter
	Set(index int, newObj PatternFlowTcpEcnNsMetricTag) PatternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter
	Clear() PatternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter
	clearHolderSlice() PatternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter
	appendHolderSlice(item PatternFlowTcpEcnNsMetricTag) PatternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter
}

func (obj *patternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter) setMsg(msg *patternFlowTcpEcnNs) PatternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowTcpEcnNsMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter) Items() []PatternFlowTcpEcnNsMetricTag {
	return obj.patternFlowTcpEcnNsMetricTagSlice
}

func (obj *patternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter) Add() PatternFlowTcpEcnNsMetricTag {
	newObj := &otg.PatternFlowTcpEcnNsMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowTcpEcnNsMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowTcpEcnNsMetricTagSlice = append(obj.patternFlowTcpEcnNsMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter) Append(items ...PatternFlowTcpEcnNsMetricTag) PatternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowTcpEcnNsMetricTagSlice = append(obj.patternFlowTcpEcnNsMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter) Set(index int, newObj PatternFlowTcpEcnNsMetricTag) PatternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowTcpEcnNsMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter) Clear() PatternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowTcpEcnNsMetricTag{}
		obj.patternFlowTcpEcnNsMetricTagSlice = []PatternFlowTcpEcnNsMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter) clearHolderSlice() PatternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter {
	if len(obj.patternFlowTcpEcnNsMetricTagSlice) > 0 {
		obj.patternFlowTcpEcnNsMetricTagSlice = []PatternFlowTcpEcnNsMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter) appendHolderSlice(item PatternFlowTcpEcnNsMetricTag) PatternFlowTcpEcnNsPatternFlowTcpEcnNsMetricTagIter {
	obj.patternFlowTcpEcnNsMetricTagSlice = append(obj.patternFlowTcpEcnNsMetricTagSlice, item)
	return obj
}

func (obj *patternFlowTcpEcnNs) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpEcnNs.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowTcpEcnNs.Values <= 1 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowTcpEcnNsMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowTcpEcnNs) setDefault() {
	var choices_set int = 0
	var choice PatternFlowTcpEcnNsChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowTcpEcnNsChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowTcpEcnNsChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowTcpEcnNsChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowTcpEcnNsChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowTcpEcnNsChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowTcpEcnNs")
			}
		} else {
			intVal := otg.PatternFlowTcpEcnNs_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowTcpEcnNs_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
