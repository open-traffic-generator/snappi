package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpCtlSyn *****
type patternFlowTcpCtlSyn struct {
	validation
	obj              *otg.PatternFlowTcpCtlSyn
	marshaller       marshalPatternFlowTcpCtlSyn
	unMarshaller     unMarshalPatternFlowTcpCtlSyn
	incrementHolder  PatternFlowTcpCtlSynCounter
	decrementHolder  PatternFlowTcpCtlSynCounter
	metricTagsHolder PatternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter
}

func NewPatternFlowTcpCtlSyn() PatternFlowTcpCtlSyn {
	obj := patternFlowTcpCtlSyn{obj: &otg.PatternFlowTcpCtlSyn{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpCtlSyn) msg() *otg.PatternFlowTcpCtlSyn {
	return obj.obj
}

func (obj *patternFlowTcpCtlSyn) setMsg(msg *otg.PatternFlowTcpCtlSyn) PatternFlowTcpCtlSyn {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpCtlSyn struct {
	obj *patternFlowTcpCtlSyn
}

type marshalPatternFlowTcpCtlSyn interface {
	// ToProto marshals PatternFlowTcpCtlSyn to protobuf object *otg.PatternFlowTcpCtlSyn
	ToProto() (*otg.PatternFlowTcpCtlSyn, error)
	// ToPbText marshals PatternFlowTcpCtlSyn to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpCtlSyn to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpCtlSyn to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowTcpCtlSyn to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowTcpCtlSyn struct {
	obj *patternFlowTcpCtlSyn
}

type unMarshalPatternFlowTcpCtlSyn interface {
	// FromProto unmarshals PatternFlowTcpCtlSyn from protobuf object *otg.PatternFlowTcpCtlSyn
	FromProto(msg *otg.PatternFlowTcpCtlSyn) (PatternFlowTcpCtlSyn, error)
	// FromPbText unmarshals PatternFlowTcpCtlSyn from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpCtlSyn from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpCtlSyn from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpCtlSyn) Marshal() marshalPatternFlowTcpCtlSyn {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpCtlSyn{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpCtlSyn) Unmarshal() unMarshalPatternFlowTcpCtlSyn {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpCtlSyn{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpCtlSyn) ToProto() (*otg.PatternFlowTcpCtlSyn, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpCtlSyn) FromProto(msg *otg.PatternFlowTcpCtlSyn) (PatternFlowTcpCtlSyn, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpCtlSyn) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlSyn) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpCtlSyn) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlSyn) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpCtlSyn) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowTcpCtlSyn) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlSyn) FromJson(value string) error {
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

func (obj *patternFlowTcpCtlSyn) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlSyn) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlSyn) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpCtlSyn) Clone() (PatternFlowTcpCtlSyn, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpCtlSyn()
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

func (obj *patternFlowTcpCtlSyn) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowTcpCtlSyn is synchronize sequenece numbers.
type PatternFlowTcpCtlSyn interface {
	Validation
	// msg marshals PatternFlowTcpCtlSyn to protobuf object *otg.PatternFlowTcpCtlSyn
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpCtlSyn
	// setMsg unmarshals PatternFlowTcpCtlSyn from protobuf object *otg.PatternFlowTcpCtlSyn
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpCtlSyn) PatternFlowTcpCtlSyn
	// provides marshal interface
	Marshal() marshalPatternFlowTcpCtlSyn
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpCtlSyn
	// validate validates PatternFlowTcpCtlSyn
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpCtlSyn, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowTcpCtlSynChoiceEnum, set in PatternFlowTcpCtlSyn
	Choice() PatternFlowTcpCtlSynChoiceEnum
	// setChoice assigns PatternFlowTcpCtlSynChoiceEnum provided by user to PatternFlowTcpCtlSyn
	setChoice(value PatternFlowTcpCtlSynChoiceEnum) PatternFlowTcpCtlSyn
	// HasChoice checks if Choice has been set in PatternFlowTcpCtlSyn
	HasChoice() bool
	// Value returns uint32, set in PatternFlowTcpCtlSyn.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowTcpCtlSyn
	SetValue(value uint32) PatternFlowTcpCtlSyn
	// HasValue checks if Value has been set in PatternFlowTcpCtlSyn
	HasValue() bool
	// Values returns []uint32, set in PatternFlowTcpCtlSyn.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowTcpCtlSyn
	SetValues(value []uint32) PatternFlowTcpCtlSyn
	// Increment returns PatternFlowTcpCtlSynCounter, set in PatternFlowTcpCtlSyn.
	// PatternFlowTcpCtlSynCounter is integer counter pattern
	Increment() PatternFlowTcpCtlSynCounter
	// SetIncrement assigns PatternFlowTcpCtlSynCounter provided by user to PatternFlowTcpCtlSyn.
	// PatternFlowTcpCtlSynCounter is integer counter pattern
	SetIncrement(value PatternFlowTcpCtlSynCounter) PatternFlowTcpCtlSyn
	// HasIncrement checks if Increment has been set in PatternFlowTcpCtlSyn
	HasIncrement() bool
	// Decrement returns PatternFlowTcpCtlSynCounter, set in PatternFlowTcpCtlSyn.
	// PatternFlowTcpCtlSynCounter is integer counter pattern
	Decrement() PatternFlowTcpCtlSynCounter
	// SetDecrement assigns PatternFlowTcpCtlSynCounter provided by user to PatternFlowTcpCtlSyn.
	// PatternFlowTcpCtlSynCounter is integer counter pattern
	SetDecrement(value PatternFlowTcpCtlSynCounter) PatternFlowTcpCtlSyn
	// HasDecrement checks if Decrement has been set in PatternFlowTcpCtlSyn
	HasDecrement() bool
	// MetricTags returns PatternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIterIter, set in PatternFlowTcpCtlSyn
	MetricTags() PatternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter
	setNil()
}

type PatternFlowTcpCtlSynChoiceEnum string

// Enum of Choice on PatternFlowTcpCtlSyn
var PatternFlowTcpCtlSynChoice = struct {
	VALUE     PatternFlowTcpCtlSynChoiceEnum
	VALUES    PatternFlowTcpCtlSynChoiceEnum
	INCREMENT PatternFlowTcpCtlSynChoiceEnum
	DECREMENT PatternFlowTcpCtlSynChoiceEnum
}{
	VALUE:     PatternFlowTcpCtlSynChoiceEnum("value"),
	VALUES:    PatternFlowTcpCtlSynChoiceEnum("values"),
	INCREMENT: PatternFlowTcpCtlSynChoiceEnum("increment"),
	DECREMENT: PatternFlowTcpCtlSynChoiceEnum("decrement"),
}

func (obj *patternFlowTcpCtlSyn) Choice() PatternFlowTcpCtlSynChoiceEnum {
	return PatternFlowTcpCtlSynChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowTcpCtlSyn) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowTcpCtlSyn) setChoice(value PatternFlowTcpCtlSynChoiceEnum) PatternFlowTcpCtlSyn {
	intValue, ok := otg.PatternFlowTcpCtlSyn_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowTcpCtlSynChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowTcpCtlSyn_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowTcpCtlSynChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowTcpCtlSynChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowTcpCtlSynChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowTcpCtlSynCounter().msg()
	}

	if value == PatternFlowTcpCtlSynChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowTcpCtlSynCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpCtlSyn) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowTcpCtlSynChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpCtlSyn) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowTcpCtlSyn object
func (obj *patternFlowTcpCtlSyn) SetValue(value uint32) PatternFlowTcpCtlSyn {
	obj.setChoice(PatternFlowTcpCtlSynChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowTcpCtlSyn) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowTcpCtlSyn object
func (obj *patternFlowTcpCtlSyn) SetValues(value []uint32) PatternFlowTcpCtlSyn {
	obj.setChoice(PatternFlowTcpCtlSynChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowTcpCtlSynCounter
func (obj *patternFlowTcpCtlSyn) Increment() PatternFlowTcpCtlSynCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowTcpCtlSynChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowTcpCtlSynCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowTcpCtlSynCounter
func (obj *patternFlowTcpCtlSyn) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowTcpCtlSynCounter value in the PatternFlowTcpCtlSyn object
func (obj *patternFlowTcpCtlSyn) SetIncrement(value PatternFlowTcpCtlSynCounter) PatternFlowTcpCtlSyn {
	obj.setChoice(PatternFlowTcpCtlSynChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowTcpCtlSynCounter
func (obj *patternFlowTcpCtlSyn) Decrement() PatternFlowTcpCtlSynCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowTcpCtlSynChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowTcpCtlSynCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowTcpCtlSynCounter
func (obj *patternFlowTcpCtlSyn) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowTcpCtlSynCounter value in the PatternFlowTcpCtlSyn object
func (obj *patternFlowTcpCtlSyn) SetDecrement(value PatternFlowTcpCtlSynCounter) PatternFlowTcpCtlSyn {
	obj.setChoice(PatternFlowTcpCtlSynChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowTcpCtlSynMetricTag
func (obj *patternFlowTcpCtlSyn) MetricTags() PatternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowTcpCtlSynMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter struct {
	obj                                *patternFlowTcpCtlSyn
	patternFlowTcpCtlSynMetricTagSlice []PatternFlowTcpCtlSynMetricTag
	fieldPtr                           *[]*otg.PatternFlowTcpCtlSynMetricTag
}

func newPatternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter(ptr *[]*otg.PatternFlowTcpCtlSynMetricTag) PatternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter {
	return &patternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter{fieldPtr: ptr}
}

type PatternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter interface {
	setMsg(*patternFlowTcpCtlSyn) PatternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter
	Items() []PatternFlowTcpCtlSynMetricTag
	Add() PatternFlowTcpCtlSynMetricTag
	Append(items ...PatternFlowTcpCtlSynMetricTag) PatternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter
	Set(index int, newObj PatternFlowTcpCtlSynMetricTag) PatternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter
	Clear() PatternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter
	clearHolderSlice() PatternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter
	appendHolderSlice(item PatternFlowTcpCtlSynMetricTag) PatternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter
}

func (obj *patternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter) setMsg(msg *patternFlowTcpCtlSyn) PatternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowTcpCtlSynMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter) Items() []PatternFlowTcpCtlSynMetricTag {
	return obj.patternFlowTcpCtlSynMetricTagSlice
}

func (obj *patternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter) Add() PatternFlowTcpCtlSynMetricTag {
	newObj := &otg.PatternFlowTcpCtlSynMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowTcpCtlSynMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowTcpCtlSynMetricTagSlice = append(obj.patternFlowTcpCtlSynMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter) Append(items ...PatternFlowTcpCtlSynMetricTag) PatternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowTcpCtlSynMetricTagSlice = append(obj.patternFlowTcpCtlSynMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter) Set(index int, newObj PatternFlowTcpCtlSynMetricTag) PatternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowTcpCtlSynMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter) Clear() PatternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowTcpCtlSynMetricTag{}
		obj.patternFlowTcpCtlSynMetricTagSlice = []PatternFlowTcpCtlSynMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter) clearHolderSlice() PatternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter {
	if len(obj.patternFlowTcpCtlSynMetricTagSlice) > 0 {
		obj.patternFlowTcpCtlSynMetricTagSlice = []PatternFlowTcpCtlSynMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter) appendHolderSlice(item PatternFlowTcpCtlSynMetricTag) PatternFlowTcpCtlSynPatternFlowTcpCtlSynMetricTagIter {
	obj.patternFlowTcpCtlSynMetricTagSlice = append(obj.patternFlowTcpCtlSynMetricTagSlice, item)
	return obj
}

func (obj *patternFlowTcpCtlSyn) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlSyn.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowTcpCtlSyn.Values <= 1 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowTcpCtlSynMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowTcpCtlSyn) setDefault() {
	var choices_set int = 0
	var choice PatternFlowTcpCtlSynChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowTcpCtlSynChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowTcpCtlSynChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowTcpCtlSynChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowTcpCtlSynChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowTcpCtlSynChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowTcpCtlSyn")
			}
		} else {
			intVal := otg.PatternFlowTcpCtlSyn_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowTcpCtlSyn_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
