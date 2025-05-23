package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4DscpEcn *****
type patternFlowIpv4DscpEcn struct {
	validation
	obj              *otg.PatternFlowIpv4DscpEcn
	marshaller       marshalPatternFlowIpv4DscpEcn
	unMarshaller     unMarshalPatternFlowIpv4DscpEcn
	incrementHolder  PatternFlowIpv4DscpEcnCounter
	decrementHolder  PatternFlowIpv4DscpEcnCounter
	metricTagsHolder PatternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter
}

func NewPatternFlowIpv4DscpEcn() PatternFlowIpv4DscpEcn {
	obj := patternFlowIpv4DscpEcn{obj: &otg.PatternFlowIpv4DscpEcn{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4DscpEcn) msg() *otg.PatternFlowIpv4DscpEcn {
	return obj.obj
}

func (obj *patternFlowIpv4DscpEcn) setMsg(msg *otg.PatternFlowIpv4DscpEcn) PatternFlowIpv4DscpEcn {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4DscpEcn struct {
	obj *patternFlowIpv4DscpEcn
}

type marshalPatternFlowIpv4DscpEcn interface {
	// ToProto marshals PatternFlowIpv4DscpEcn to protobuf object *otg.PatternFlowIpv4DscpEcn
	ToProto() (*otg.PatternFlowIpv4DscpEcn, error)
	// ToPbText marshals PatternFlowIpv4DscpEcn to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4DscpEcn to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4DscpEcn to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv4DscpEcn to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv4DscpEcn struct {
	obj *patternFlowIpv4DscpEcn
}

type unMarshalPatternFlowIpv4DscpEcn interface {
	// FromProto unmarshals PatternFlowIpv4DscpEcn from protobuf object *otg.PatternFlowIpv4DscpEcn
	FromProto(msg *otg.PatternFlowIpv4DscpEcn) (PatternFlowIpv4DscpEcn, error)
	// FromPbText unmarshals PatternFlowIpv4DscpEcn from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4DscpEcn from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4DscpEcn from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4DscpEcn) Marshal() marshalPatternFlowIpv4DscpEcn {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4DscpEcn{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4DscpEcn) Unmarshal() unMarshalPatternFlowIpv4DscpEcn {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4DscpEcn{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4DscpEcn) ToProto() (*otg.PatternFlowIpv4DscpEcn, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4DscpEcn) FromProto(msg *otg.PatternFlowIpv4DscpEcn) (PatternFlowIpv4DscpEcn, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4DscpEcn) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DscpEcn) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4DscpEcn) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DscpEcn) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4DscpEcn) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv4DscpEcn) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DscpEcn) FromJson(value string) error {
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

func (obj *patternFlowIpv4DscpEcn) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4DscpEcn) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4DscpEcn) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4DscpEcn) Clone() (PatternFlowIpv4DscpEcn, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4DscpEcn()
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

func (obj *patternFlowIpv4DscpEcn) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4DscpEcn is explicit congestion notification
type PatternFlowIpv4DscpEcn interface {
	Validation
	// msg marshals PatternFlowIpv4DscpEcn to protobuf object *otg.PatternFlowIpv4DscpEcn
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4DscpEcn
	// setMsg unmarshals PatternFlowIpv4DscpEcn from protobuf object *otg.PatternFlowIpv4DscpEcn
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4DscpEcn) PatternFlowIpv4DscpEcn
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4DscpEcn
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4DscpEcn
	// validate validates PatternFlowIpv4DscpEcn
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4DscpEcn, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4DscpEcnChoiceEnum, set in PatternFlowIpv4DscpEcn
	Choice() PatternFlowIpv4DscpEcnChoiceEnum
	// setChoice assigns PatternFlowIpv4DscpEcnChoiceEnum provided by user to PatternFlowIpv4DscpEcn
	setChoice(value PatternFlowIpv4DscpEcnChoiceEnum) PatternFlowIpv4DscpEcn
	// HasChoice checks if Choice has been set in PatternFlowIpv4DscpEcn
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv4DscpEcn.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv4DscpEcn
	SetValue(value uint32) PatternFlowIpv4DscpEcn
	// HasValue checks if Value has been set in PatternFlowIpv4DscpEcn
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv4DscpEcn.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv4DscpEcn
	SetValues(value []uint32) PatternFlowIpv4DscpEcn
	// Increment returns PatternFlowIpv4DscpEcnCounter, set in PatternFlowIpv4DscpEcn.
	// PatternFlowIpv4DscpEcnCounter is integer counter pattern
	Increment() PatternFlowIpv4DscpEcnCounter
	// SetIncrement assigns PatternFlowIpv4DscpEcnCounter provided by user to PatternFlowIpv4DscpEcn.
	// PatternFlowIpv4DscpEcnCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv4DscpEcnCounter) PatternFlowIpv4DscpEcn
	// HasIncrement checks if Increment has been set in PatternFlowIpv4DscpEcn
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4DscpEcnCounter, set in PatternFlowIpv4DscpEcn.
	// PatternFlowIpv4DscpEcnCounter is integer counter pattern
	Decrement() PatternFlowIpv4DscpEcnCounter
	// SetDecrement assigns PatternFlowIpv4DscpEcnCounter provided by user to PatternFlowIpv4DscpEcn.
	// PatternFlowIpv4DscpEcnCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv4DscpEcnCounter) PatternFlowIpv4DscpEcn
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4DscpEcn
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIterIter, set in PatternFlowIpv4DscpEcn
	MetricTags() PatternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter
	setNil()
}

type PatternFlowIpv4DscpEcnChoiceEnum string

// Enum of Choice on PatternFlowIpv4DscpEcn
var PatternFlowIpv4DscpEcnChoice = struct {
	VALUE     PatternFlowIpv4DscpEcnChoiceEnum
	VALUES    PatternFlowIpv4DscpEcnChoiceEnum
	INCREMENT PatternFlowIpv4DscpEcnChoiceEnum
	DECREMENT PatternFlowIpv4DscpEcnChoiceEnum
}{
	VALUE:     PatternFlowIpv4DscpEcnChoiceEnum("value"),
	VALUES:    PatternFlowIpv4DscpEcnChoiceEnum("values"),
	INCREMENT: PatternFlowIpv4DscpEcnChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4DscpEcnChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4DscpEcn) Choice() PatternFlowIpv4DscpEcnChoiceEnum {
	return PatternFlowIpv4DscpEcnChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4DscpEcn) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4DscpEcn) setChoice(value PatternFlowIpv4DscpEcnChoiceEnum) PatternFlowIpv4DscpEcn {
	intValue, ok := otg.PatternFlowIpv4DscpEcn_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4DscpEcnChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4DscpEcn_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4DscpEcnChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4DscpEcnChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4DscpEcnChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4DscpEcnCounter().msg()
	}

	if value == PatternFlowIpv4DscpEcnChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4DscpEcnCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4DscpEcn) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4DscpEcnChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4DscpEcn) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv4DscpEcn object
func (obj *patternFlowIpv4DscpEcn) SetValue(value uint32) PatternFlowIpv4DscpEcn {
	obj.setChoice(PatternFlowIpv4DscpEcnChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv4DscpEcn) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv4DscpEcn object
func (obj *patternFlowIpv4DscpEcn) SetValues(value []uint32) PatternFlowIpv4DscpEcn {
	obj.setChoice(PatternFlowIpv4DscpEcnChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv4DscpEcnCounter
func (obj *patternFlowIpv4DscpEcn) Increment() PatternFlowIpv4DscpEcnCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4DscpEcnChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4DscpEcnCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4DscpEcnCounter
func (obj *patternFlowIpv4DscpEcn) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4DscpEcnCounter value in the PatternFlowIpv4DscpEcn object
func (obj *patternFlowIpv4DscpEcn) SetIncrement(value PatternFlowIpv4DscpEcnCounter) PatternFlowIpv4DscpEcn {
	obj.setChoice(PatternFlowIpv4DscpEcnChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4DscpEcnCounter
func (obj *patternFlowIpv4DscpEcn) Decrement() PatternFlowIpv4DscpEcnCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4DscpEcnChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4DscpEcnCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4DscpEcnCounter
func (obj *patternFlowIpv4DscpEcn) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4DscpEcnCounter value in the PatternFlowIpv4DscpEcn object
func (obj *patternFlowIpv4DscpEcn) SetDecrement(value PatternFlowIpv4DscpEcnCounter) PatternFlowIpv4DscpEcn {
	obj.setChoice(PatternFlowIpv4DscpEcnChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv4DscpEcnMetricTag
func (obj *patternFlowIpv4DscpEcn) MetricTags() PatternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv4DscpEcnMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter struct {
	obj                                  *patternFlowIpv4DscpEcn
	patternFlowIpv4DscpEcnMetricTagSlice []PatternFlowIpv4DscpEcnMetricTag
	fieldPtr                             *[]*otg.PatternFlowIpv4DscpEcnMetricTag
}

func newPatternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter(ptr *[]*otg.PatternFlowIpv4DscpEcnMetricTag) PatternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter {
	return &patternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter interface {
	setMsg(*patternFlowIpv4DscpEcn) PatternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter
	Items() []PatternFlowIpv4DscpEcnMetricTag
	Add() PatternFlowIpv4DscpEcnMetricTag
	Append(items ...PatternFlowIpv4DscpEcnMetricTag) PatternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter
	Set(index int, newObj PatternFlowIpv4DscpEcnMetricTag) PatternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter
	Clear() PatternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter
	clearHolderSlice() PatternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter
	appendHolderSlice(item PatternFlowIpv4DscpEcnMetricTag) PatternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter
}

func (obj *patternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter) setMsg(msg *patternFlowIpv4DscpEcn) PatternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv4DscpEcnMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter) Items() []PatternFlowIpv4DscpEcnMetricTag {
	return obj.patternFlowIpv4DscpEcnMetricTagSlice
}

func (obj *patternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter) Add() PatternFlowIpv4DscpEcnMetricTag {
	newObj := &otg.PatternFlowIpv4DscpEcnMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv4DscpEcnMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv4DscpEcnMetricTagSlice = append(obj.patternFlowIpv4DscpEcnMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter) Append(items ...PatternFlowIpv4DscpEcnMetricTag) PatternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv4DscpEcnMetricTagSlice = append(obj.patternFlowIpv4DscpEcnMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter) Set(index int, newObj PatternFlowIpv4DscpEcnMetricTag) PatternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv4DscpEcnMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter) Clear() PatternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv4DscpEcnMetricTag{}
		obj.patternFlowIpv4DscpEcnMetricTagSlice = []PatternFlowIpv4DscpEcnMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter) clearHolderSlice() PatternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter {
	if len(obj.patternFlowIpv4DscpEcnMetricTagSlice) > 0 {
		obj.patternFlowIpv4DscpEcnMetricTagSlice = []PatternFlowIpv4DscpEcnMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter) appendHolderSlice(item PatternFlowIpv4DscpEcnMetricTag) PatternFlowIpv4DscpEcnPatternFlowIpv4DscpEcnMetricTagIter {
	obj.patternFlowIpv4DscpEcnMetricTagSlice = append(obj.patternFlowIpv4DscpEcnMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIpv4DscpEcn) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 3 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4DscpEcn.Value <= 3 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 3 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv4DscpEcn.Values <= 3 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv4DscpEcnMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIpv4DscpEcn) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv4DscpEcnChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv4DscpEcnChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv4DscpEcnChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv4DscpEcnChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv4DscpEcnChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv4DscpEcnChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv4DscpEcn")
			}
		} else {
			intVal := otg.PatternFlowIpv4DscpEcn_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv4DscpEcn_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
