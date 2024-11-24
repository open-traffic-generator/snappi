package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1NPduNumber *****
type patternFlowGtpv1NPduNumber struct {
	validation
	obj              *otg.PatternFlowGtpv1NPduNumber
	marshaller       marshalPatternFlowGtpv1NPduNumber
	unMarshaller     unMarshalPatternFlowGtpv1NPduNumber
	incrementHolder  PatternFlowGtpv1NPduNumberCounter
	decrementHolder  PatternFlowGtpv1NPduNumberCounter
	metricTagsHolder PatternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter
}

func NewPatternFlowGtpv1NPduNumber() PatternFlowGtpv1NPduNumber {
	obj := patternFlowGtpv1NPduNumber{obj: &otg.PatternFlowGtpv1NPduNumber{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1NPduNumber) msg() *otg.PatternFlowGtpv1NPduNumber {
	return obj.obj
}

func (obj *patternFlowGtpv1NPduNumber) setMsg(msg *otg.PatternFlowGtpv1NPduNumber) PatternFlowGtpv1NPduNumber {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1NPduNumber struct {
	obj *patternFlowGtpv1NPduNumber
}

type marshalPatternFlowGtpv1NPduNumber interface {
	// ToProto marshals PatternFlowGtpv1NPduNumber to protobuf object *otg.PatternFlowGtpv1NPduNumber
	ToProto() (*otg.PatternFlowGtpv1NPduNumber, error)
	// ToPbText marshals PatternFlowGtpv1NPduNumber to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1NPduNumber to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1NPduNumber to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowGtpv1NPduNumber to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowGtpv1NPduNumber struct {
	obj *patternFlowGtpv1NPduNumber
}

type unMarshalPatternFlowGtpv1NPduNumber interface {
	// FromProto unmarshals PatternFlowGtpv1NPduNumber from protobuf object *otg.PatternFlowGtpv1NPduNumber
	FromProto(msg *otg.PatternFlowGtpv1NPduNumber) (PatternFlowGtpv1NPduNumber, error)
	// FromPbText unmarshals PatternFlowGtpv1NPduNumber from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1NPduNumber from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1NPduNumber from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1NPduNumber) Marshal() marshalPatternFlowGtpv1NPduNumber {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1NPduNumber{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1NPduNumber) Unmarshal() unMarshalPatternFlowGtpv1NPduNumber {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1NPduNumber{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1NPduNumber) ToProto() (*otg.PatternFlowGtpv1NPduNumber, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1NPduNumber) FromProto(msg *otg.PatternFlowGtpv1NPduNumber) (PatternFlowGtpv1NPduNumber, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1NPduNumber) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1NPduNumber) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1NPduNumber) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1NPduNumber) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1NPduNumber) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowGtpv1NPduNumber) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1NPduNumber) FromJson(value string) error {
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

func (obj *patternFlowGtpv1NPduNumber) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1NPduNumber) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1NPduNumber) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1NPduNumber) Clone() (PatternFlowGtpv1NPduNumber, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1NPduNumber()
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

func (obj *patternFlowGtpv1NPduNumber) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGtpv1NPduNumber is n-PDU number. Exists if any of the e_flag, s_flag, or pn_flag bits are on.  Must be interpreted only if the pn_flag bit is on.
type PatternFlowGtpv1NPduNumber interface {
	Validation
	// msg marshals PatternFlowGtpv1NPduNumber to protobuf object *otg.PatternFlowGtpv1NPduNumber
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1NPduNumber
	// setMsg unmarshals PatternFlowGtpv1NPduNumber from protobuf object *otg.PatternFlowGtpv1NPduNumber
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1NPduNumber) PatternFlowGtpv1NPduNumber
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1NPduNumber
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1NPduNumber
	// validate validates PatternFlowGtpv1NPduNumber
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1NPduNumber, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGtpv1NPduNumberChoiceEnum, set in PatternFlowGtpv1NPduNumber
	Choice() PatternFlowGtpv1NPduNumberChoiceEnum
	// setChoice assigns PatternFlowGtpv1NPduNumberChoiceEnum provided by user to PatternFlowGtpv1NPduNumber
	setChoice(value PatternFlowGtpv1NPduNumberChoiceEnum) PatternFlowGtpv1NPduNumber
	// HasChoice checks if Choice has been set in PatternFlowGtpv1NPduNumber
	HasChoice() bool
	// Value returns uint32, set in PatternFlowGtpv1NPduNumber.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowGtpv1NPduNumber
	SetValue(value uint32) PatternFlowGtpv1NPduNumber
	// HasValue checks if Value has been set in PatternFlowGtpv1NPduNumber
	HasValue() bool
	// Values returns []uint32, set in PatternFlowGtpv1NPduNumber.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowGtpv1NPduNumber
	SetValues(value []uint32) PatternFlowGtpv1NPduNumber
	// Increment returns PatternFlowGtpv1NPduNumberCounter, set in PatternFlowGtpv1NPduNumber.
	// PatternFlowGtpv1NPduNumberCounter is integer counter pattern
	Increment() PatternFlowGtpv1NPduNumberCounter
	// SetIncrement assigns PatternFlowGtpv1NPduNumberCounter provided by user to PatternFlowGtpv1NPduNumber.
	// PatternFlowGtpv1NPduNumberCounter is integer counter pattern
	SetIncrement(value PatternFlowGtpv1NPduNumberCounter) PatternFlowGtpv1NPduNumber
	// HasIncrement checks if Increment has been set in PatternFlowGtpv1NPduNumber
	HasIncrement() bool
	// Decrement returns PatternFlowGtpv1NPduNumberCounter, set in PatternFlowGtpv1NPduNumber.
	// PatternFlowGtpv1NPduNumberCounter is integer counter pattern
	Decrement() PatternFlowGtpv1NPduNumberCounter
	// SetDecrement assigns PatternFlowGtpv1NPduNumberCounter provided by user to PatternFlowGtpv1NPduNumber.
	// PatternFlowGtpv1NPduNumberCounter is integer counter pattern
	SetDecrement(value PatternFlowGtpv1NPduNumberCounter) PatternFlowGtpv1NPduNumber
	// HasDecrement checks if Decrement has been set in PatternFlowGtpv1NPduNumber
	HasDecrement() bool
	// MetricTags returns PatternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIterIter, set in PatternFlowGtpv1NPduNumber
	MetricTags() PatternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter
	setNil()
}

type PatternFlowGtpv1NPduNumberChoiceEnum string

// Enum of Choice on PatternFlowGtpv1NPduNumber
var PatternFlowGtpv1NPduNumberChoice = struct {
	VALUE     PatternFlowGtpv1NPduNumberChoiceEnum
	VALUES    PatternFlowGtpv1NPduNumberChoiceEnum
	INCREMENT PatternFlowGtpv1NPduNumberChoiceEnum
	DECREMENT PatternFlowGtpv1NPduNumberChoiceEnum
}{
	VALUE:     PatternFlowGtpv1NPduNumberChoiceEnum("value"),
	VALUES:    PatternFlowGtpv1NPduNumberChoiceEnum("values"),
	INCREMENT: PatternFlowGtpv1NPduNumberChoiceEnum("increment"),
	DECREMENT: PatternFlowGtpv1NPduNumberChoiceEnum("decrement"),
}

func (obj *patternFlowGtpv1NPduNumber) Choice() PatternFlowGtpv1NPduNumberChoiceEnum {
	return PatternFlowGtpv1NPduNumberChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGtpv1NPduNumber) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGtpv1NPduNumber) setChoice(value PatternFlowGtpv1NPduNumberChoiceEnum) PatternFlowGtpv1NPduNumber {
	intValue, ok := otg.PatternFlowGtpv1NPduNumber_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGtpv1NPduNumberChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGtpv1NPduNumber_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGtpv1NPduNumberChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGtpv1NPduNumberChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGtpv1NPduNumberChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGtpv1NPduNumberCounter().msg()
	}

	if value == PatternFlowGtpv1NPduNumberChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGtpv1NPduNumberCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv1NPduNumber) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGtpv1NPduNumberChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv1NPduNumber) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowGtpv1NPduNumber object
func (obj *patternFlowGtpv1NPduNumber) SetValue(value uint32) PatternFlowGtpv1NPduNumber {
	obj.setChoice(PatternFlowGtpv1NPduNumberChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowGtpv1NPduNumber) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowGtpv1NPduNumber object
func (obj *patternFlowGtpv1NPduNumber) SetValues(value []uint32) PatternFlowGtpv1NPduNumber {
	obj.setChoice(PatternFlowGtpv1NPduNumberChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGtpv1NPduNumberCounter
func (obj *patternFlowGtpv1NPduNumber) Increment() PatternFlowGtpv1NPduNumberCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGtpv1NPduNumberChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGtpv1NPduNumberCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGtpv1NPduNumberCounter
func (obj *patternFlowGtpv1NPduNumber) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGtpv1NPduNumberCounter value in the PatternFlowGtpv1NPduNumber object
func (obj *patternFlowGtpv1NPduNumber) SetIncrement(value PatternFlowGtpv1NPduNumberCounter) PatternFlowGtpv1NPduNumber {
	obj.setChoice(PatternFlowGtpv1NPduNumberChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGtpv1NPduNumberCounter
func (obj *patternFlowGtpv1NPduNumber) Decrement() PatternFlowGtpv1NPduNumberCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGtpv1NPduNumberChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGtpv1NPduNumberCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGtpv1NPduNumberCounter
func (obj *patternFlowGtpv1NPduNumber) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGtpv1NPduNumberCounter value in the PatternFlowGtpv1NPduNumber object
func (obj *patternFlowGtpv1NPduNumber) SetDecrement(value PatternFlowGtpv1NPduNumberCounter) PatternFlowGtpv1NPduNumber {
	obj.setChoice(PatternFlowGtpv1NPduNumberChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGtpv1NPduNumberMetricTag
func (obj *patternFlowGtpv1NPduNumber) MetricTags() PatternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGtpv1NPduNumberMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter struct {
	obj                                      *patternFlowGtpv1NPduNumber
	patternFlowGtpv1NPduNumberMetricTagSlice []PatternFlowGtpv1NPduNumberMetricTag
	fieldPtr                                 *[]*otg.PatternFlowGtpv1NPduNumberMetricTag
}

func newPatternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter(ptr *[]*otg.PatternFlowGtpv1NPduNumberMetricTag) PatternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter {
	return &patternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter{fieldPtr: ptr}
}

type PatternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter interface {
	setMsg(*patternFlowGtpv1NPduNumber) PatternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter
	Items() []PatternFlowGtpv1NPduNumberMetricTag
	Add() PatternFlowGtpv1NPduNumberMetricTag
	Append(items ...PatternFlowGtpv1NPduNumberMetricTag) PatternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter
	Set(index int, newObj PatternFlowGtpv1NPduNumberMetricTag) PatternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter
	Clear() PatternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter
	clearHolderSlice() PatternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter
	appendHolderSlice(item PatternFlowGtpv1NPduNumberMetricTag) PatternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter
}

func (obj *patternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter) setMsg(msg *patternFlowGtpv1NPduNumber) PatternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGtpv1NPduNumberMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter) Items() []PatternFlowGtpv1NPduNumberMetricTag {
	return obj.patternFlowGtpv1NPduNumberMetricTagSlice
}

func (obj *patternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter) Add() PatternFlowGtpv1NPduNumberMetricTag {
	newObj := &otg.PatternFlowGtpv1NPduNumberMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGtpv1NPduNumberMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGtpv1NPduNumberMetricTagSlice = append(obj.patternFlowGtpv1NPduNumberMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter) Append(items ...PatternFlowGtpv1NPduNumberMetricTag) PatternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGtpv1NPduNumberMetricTagSlice = append(obj.patternFlowGtpv1NPduNumberMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter) Set(index int, newObj PatternFlowGtpv1NPduNumberMetricTag) PatternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGtpv1NPduNumberMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter) Clear() PatternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGtpv1NPduNumberMetricTag{}
		obj.patternFlowGtpv1NPduNumberMetricTagSlice = []PatternFlowGtpv1NPduNumberMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter) clearHolderSlice() PatternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter {
	if len(obj.patternFlowGtpv1NPduNumberMetricTagSlice) > 0 {
		obj.patternFlowGtpv1NPduNumberMetricTagSlice = []PatternFlowGtpv1NPduNumberMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter) appendHolderSlice(item PatternFlowGtpv1NPduNumberMetricTag) PatternFlowGtpv1NPduNumberPatternFlowGtpv1NPduNumberMetricTagIter {
	obj.patternFlowGtpv1NPduNumberMetricTagSlice = append(obj.patternFlowGtpv1NPduNumberMetricTagSlice, item)
	return obj
}

func (obj *patternFlowGtpv1NPduNumber) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1NPduNumber.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowGtpv1NPduNumber.Values <= 255 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowGtpv1NPduNumberMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGtpv1NPduNumber) setDefault() {
	var choices_set int = 0
	var choice PatternFlowGtpv1NPduNumberChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowGtpv1NPduNumberChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowGtpv1NPduNumberChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowGtpv1NPduNumberChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowGtpv1NPduNumberChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowGtpv1NPduNumberChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowGtpv1NPduNumber")
			}
		} else {
			intVal := otg.PatternFlowGtpv1NPduNumber_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowGtpv1NPduNumber_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
