package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpCtlRst *****
type patternFlowTcpCtlRst struct {
	validation
	obj              *otg.PatternFlowTcpCtlRst
	marshaller       marshalPatternFlowTcpCtlRst
	unMarshaller     unMarshalPatternFlowTcpCtlRst
	incrementHolder  PatternFlowTcpCtlRstCounter
	decrementHolder  PatternFlowTcpCtlRstCounter
	metricTagsHolder PatternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter
}

func NewPatternFlowTcpCtlRst() PatternFlowTcpCtlRst {
	obj := patternFlowTcpCtlRst{obj: &otg.PatternFlowTcpCtlRst{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpCtlRst) msg() *otg.PatternFlowTcpCtlRst {
	return obj.obj
}

func (obj *patternFlowTcpCtlRst) setMsg(msg *otg.PatternFlowTcpCtlRst) PatternFlowTcpCtlRst {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpCtlRst struct {
	obj *patternFlowTcpCtlRst
}

type marshalPatternFlowTcpCtlRst interface {
	// ToProto marshals PatternFlowTcpCtlRst to protobuf object *otg.PatternFlowTcpCtlRst
	ToProto() (*otg.PatternFlowTcpCtlRst, error)
	// ToPbText marshals PatternFlowTcpCtlRst to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpCtlRst to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpCtlRst to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowTcpCtlRst to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowTcpCtlRst struct {
	obj *patternFlowTcpCtlRst
}

type unMarshalPatternFlowTcpCtlRst interface {
	// FromProto unmarshals PatternFlowTcpCtlRst from protobuf object *otg.PatternFlowTcpCtlRst
	FromProto(msg *otg.PatternFlowTcpCtlRst) (PatternFlowTcpCtlRst, error)
	// FromPbText unmarshals PatternFlowTcpCtlRst from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpCtlRst from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpCtlRst from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpCtlRst) Marshal() marshalPatternFlowTcpCtlRst {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpCtlRst{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpCtlRst) Unmarshal() unMarshalPatternFlowTcpCtlRst {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpCtlRst{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpCtlRst) ToProto() (*otg.PatternFlowTcpCtlRst, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpCtlRst) FromProto(msg *otg.PatternFlowTcpCtlRst) (PatternFlowTcpCtlRst, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpCtlRst) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlRst) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpCtlRst) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlRst) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpCtlRst) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowTcpCtlRst) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlRst) FromJson(value string) error {
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

func (obj *patternFlowTcpCtlRst) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlRst) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlRst) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpCtlRst) Clone() (PatternFlowTcpCtlRst, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpCtlRst()
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

func (obj *patternFlowTcpCtlRst) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowTcpCtlRst is reset the connection.
type PatternFlowTcpCtlRst interface {
	Validation
	// msg marshals PatternFlowTcpCtlRst to protobuf object *otg.PatternFlowTcpCtlRst
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpCtlRst
	// setMsg unmarshals PatternFlowTcpCtlRst from protobuf object *otg.PatternFlowTcpCtlRst
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpCtlRst) PatternFlowTcpCtlRst
	// provides marshal interface
	Marshal() marshalPatternFlowTcpCtlRst
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpCtlRst
	// validate validates PatternFlowTcpCtlRst
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpCtlRst, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowTcpCtlRstChoiceEnum, set in PatternFlowTcpCtlRst
	Choice() PatternFlowTcpCtlRstChoiceEnum
	// setChoice assigns PatternFlowTcpCtlRstChoiceEnum provided by user to PatternFlowTcpCtlRst
	setChoice(value PatternFlowTcpCtlRstChoiceEnum) PatternFlowTcpCtlRst
	// HasChoice checks if Choice has been set in PatternFlowTcpCtlRst
	HasChoice() bool
	// Value returns uint32, set in PatternFlowTcpCtlRst.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowTcpCtlRst
	SetValue(value uint32) PatternFlowTcpCtlRst
	// HasValue checks if Value has been set in PatternFlowTcpCtlRst
	HasValue() bool
	// Values returns []uint32, set in PatternFlowTcpCtlRst.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowTcpCtlRst
	SetValues(value []uint32) PatternFlowTcpCtlRst
	// Increment returns PatternFlowTcpCtlRstCounter, set in PatternFlowTcpCtlRst.
	// PatternFlowTcpCtlRstCounter is integer counter pattern
	Increment() PatternFlowTcpCtlRstCounter
	// SetIncrement assigns PatternFlowTcpCtlRstCounter provided by user to PatternFlowTcpCtlRst.
	// PatternFlowTcpCtlRstCounter is integer counter pattern
	SetIncrement(value PatternFlowTcpCtlRstCounter) PatternFlowTcpCtlRst
	// HasIncrement checks if Increment has been set in PatternFlowTcpCtlRst
	HasIncrement() bool
	// Decrement returns PatternFlowTcpCtlRstCounter, set in PatternFlowTcpCtlRst.
	// PatternFlowTcpCtlRstCounter is integer counter pattern
	Decrement() PatternFlowTcpCtlRstCounter
	// SetDecrement assigns PatternFlowTcpCtlRstCounter provided by user to PatternFlowTcpCtlRst.
	// PatternFlowTcpCtlRstCounter is integer counter pattern
	SetDecrement(value PatternFlowTcpCtlRstCounter) PatternFlowTcpCtlRst
	// HasDecrement checks if Decrement has been set in PatternFlowTcpCtlRst
	HasDecrement() bool
	// MetricTags returns PatternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIterIter, set in PatternFlowTcpCtlRst
	MetricTags() PatternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter
	setNil()
}

type PatternFlowTcpCtlRstChoiceEnum string

// Enum of Choice on PatternFlowTcpCtlRst
var PatternFlowTcpCtlRstChoice = struct {
	VALUE     PatternFlowTcpCtlRstChoiceEnum
	VALUES    PatternFlowTcpCtlRstChoiceEnum
	INCREMENT PatternFlowTcpCtlRstChoiceEnum
	DECREMENT PatternFlowTcpCtlRstChoiceEnum
}{
	VALUE:     PatternFlowTcpCtlRstChoiceEnum("value"),
	VALUES:    PatternFlowTcpCtlRstChoiceEnum("values"),
	INCREMENT: PatternFlowTcpCtlRstChoiceEnum("increment"),
	DECREMENT: PatternFlowTcpCtlRstChoiceEnum("decrement"),
}

func (obj *patternFlowTcpCtlRst) Choice() PatternFlowTcpCtlRstChoiceEnum {
	return PatternFlowTcpCtlRstChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowTcpCtlRst) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowTcpCtlRst) setChoice(value PatternFlowTcpCtlRstChoiceEnum) PatternFlowTcpCtlRst {
	intValue, ok := otg.PatternFlowTcpCtlRst_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowTcpCtlRstChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowTcpCtlRst_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowTcpCtlRstChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowTcpCtlRstChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowTcpCtlRstChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowTcpCtlRstCounter().msg()
	}

	if value == PatternFlowTcpCtlRstChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowTcpCtlRstCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpCtlRst) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowTcpCtlRstChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpCtlRst) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowTcpCtlRst object
func (obj *patternFlowTcpCtlRst) SetValue(value uint32) PatternFlowTcpCtlRst {
	obj.setChoice(PatternFlowTcpCtlRstChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowTcpCtlRst) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowTcpCtlRst object
func (obj *patternFlowTcpCtlRst) SetValues(value []uint32) PatternFlowTcpCtlRst {
	obj.setChoice(PatternFlowTcpCtlRstChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowTcpCtlRstCounter
func (obj *patternFlowTcpCtlRst) Increment() PatternFlowTcpCtlRstCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowTcpCtlRstChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowTcpCtlRstCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowTcpCtlRstCounter
func (obj *patternFlowTcpCtlRst) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowTcpCtlRstCounter value in the PatternFlowTcpCtlRst object
func (obj *patternFlowTcpCtlRst) SetIncrement(value PatternFlowTcpCtlRstCounter) PatternFlowTcpCtlRst {
	obj.setChoice(PatternFlowTcpCtlRstChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowTcpCtlRstCounter
func (obj *patternFlowTcpCtlRst) Decrement() PatternFlowTcpCtlRstCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowTcpCtlRstChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowTcpCtlRstCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowTcpCtlRstCounter
func (obj *patternFlowTcpCtlRst) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowTcpCtlRstCounter value in the PatternFlowTcpCtlRst object
func (obj *patternFlowTcpCtlRst) SetDecrement(value PatternFlowTcpCtlRstCounter) PatternFlowTcpCtlRst {
	obj.setChoice(PatternFlowTcpCtlRstChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowTcpCtlRstMetricTag
func (obj *patternFlowTcpCtlRst) MetricTags() PatternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowTcpCtlRstMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter struct {
	obj                                *patternFlowTcpCtlRst
	patternFlowTcpCtlRstMetricTagSlice []PatternFlowTcpCtlRstMetricTag
	fieldPtr                           *[]*otg.PatternFlowTcpCtlRstMetricTag
}

func newPatternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter(ptr *[]*otg.PatternFlowTcpCtlRstMetricTag) PatternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter {
	return &patternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter{fieldPtr: ptr}
}

type PatternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter interface {
	setMsg(*patternFlowTcpCtlRst) PatternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter
	Items() []PatternFlowTcpCtlRstMetricTag
	Add() PatternFlowTcpCtlRstMetricTag
	Append(items ...PatternFlowTcpCtlRstMetricTag) PatternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter
	Set(index int, newObj PatternFlowTcpCtlRstMetricTag) PatternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter
	Clear() PatternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter
	clearHolderSlice() PatternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter
	appendHolderSlice(item PatternFlowTcpCtlRstMetricTag) PatternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter
}

func (obj *patternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter) setMsg(msg *patternFlowTcpCtlRst) PatternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowTcpCtlRstMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter) Items() []PatternFlowTcpCtlRstMetricTag {
	return obj.patternFlowTcpCtlRstMetricTagSlice
}

func (obj *patternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter) Add() PatternFlowTcpCtlRstMetricTag {
	newObj := &otg.PatternFlowTcpCtlRstMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowTcpCtlRstMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowTcpCtlRstMetricTagSlice = append(obj.patternFlowTcpCtlRstMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter) Append(items ...PatternFlowTcpCtlRstMetricTag) PatternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowTcpCtlRstMetricTagSlice = append(obj.patternFlowTcpCtlRstMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter) Set(index int, newObj PatternFlowTcpCtlRstMetricTag) PatternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowTcpCtlRstMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter) Clear() PatternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowTcpCtlRstMetricTag{}
		obj.patternFlowTcpCtlRstMetricTagSlice = []PatternFlowTcpCtlRstMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter) clearHolderSlice() PatternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter {
	if len(obj.patternFlowTcpCtlRstMetricTagSlice) > 0 {
		obj.patternFlowTcpCtlRstMetricTagSlice = []PatternFlowTcpCtlRstMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter) appendHolderSlice(item PatternFlowTcpCtlRstMetricTag) PatternFlowTcpCtlRstPatternFlowTcpCtlRstMetricTagIter {
	obj.patternFlowTcpCtlRstMetricTagSlice = append(obj.patternFlowTcpCtlRstMetricTagSlice, item)
	return obj
}

func (obj *patternFlowTcpCtlRst) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlRst.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowTcpCtlRst.Values <= 1 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowTcpCtlRstMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowTcpCtlRst) setDefault() {
	var choices_set int = 0
	var choice PatternFlowTcpCtlRstChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowTcpCtlRstChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowTcpCtlRstChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowTcpCtlRstChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowTcpCtlRstChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowTcpCtlRstChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowTcpCtlRst")
			}
		} else {
			intVal := otg.PatternFlowTcpCtlRst_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowTcpCtlRst_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
