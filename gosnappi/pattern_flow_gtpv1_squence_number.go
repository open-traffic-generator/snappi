package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1SquenceNumber *****
type patternFlowGtpv1SquenceNumber struct {
	validation
	obj              *otg.PatternFlowGtpv1SquenceNumber
	marshaller       marshalPatternFlowGtpv1SquenceNumber
	unMarshaller     unMarshalPatternFlowGtpv1SquenceNumber
	incrementHolder  PatternFlowGtpv1SquenceNumberCounter
	decrementHolder  PatternFlowGtpv1SquenceNumberCounter
	metricTagsHolder PatternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter
}

func NewPatternFlowGtpv1SquenceNumber() PatternFlowGtpv1SquenceNumber {
	obj := patternFlowGtpv1SquenceNumber{obj: &otg.PatternFlowGtpv1SquenceNumber{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1SquenceNumber) msg() *otg.PatternFlowGtpv1SquenceNumber {
	return obj.obj
}

func (obj *patternFlowGtpv1SquenceNumber) setMsg(msg *otg.PatternFlowGtpv1SquenceNumber) PatternFlowGtpv1SquenceNumber {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1SquenceNumber struct {
	obj *patternFlowGtpv1SquenceNumber
}

type marshalPatternFlowGtpv1SquenceNumber interface {
	// ToProto marshals PatternFlowGtpv1SquenceNumber to protobuf object *otg.PatternFlowGtpv1SquenceNumber
	ToProto() (*otg.PatternFlowGtpv1SquenceNumber, error)
	// ToPbText marshals PatternFlowGtpv1SquenceNumber to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1SquenceNumber to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1SquenceNumber to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowGtpv1SquenceNumber to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowGtpv1SquenceNumber struct {
	obj *patternFlowGtpv1SquenceNumber
}

type unMarshalPatternFlowGtpv1SquenceNumber interface {
	// FromProto unmarshals PatternFlowGtpv1SquenceNumber from protobuf object *otg.PatternFlowGtpv1SquenceNumber
	FromProto(msg *otg.PatternFlowGtpv1SquenceNumber) (PatternFlowGtpv1SquenceNumber, error)
	// FromPbText unmarshals PatternFlowGtpv1SquenceNumber from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1SquenceNumber from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1SquenceNumber from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1SquenceNumber) Marshal() marshalPatternFlowGtpv1SquenceNumber {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1SquenceNumber{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1SquenceNumber) Unmarshal() unMarshalPatternFlowGtpv1SquenceNumber {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1SquenceNumber{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1SquenceNumber) ToProto() (*otg.PatternFlowGtpv1SquenceNumber, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1SquenceNumber) FromProto(msg *otg.PatternFlowGtpv1SquenceNumber) (PatternFlowGtpv1SquenceNumber, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1SquenceNumber) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1SquenceNumber) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1SquenceNumber) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1SquenceNumber) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1SquenceNumber) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowGtpv1SquenceNumber) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1SquenceNumber) FromJson(value string) error {
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

func (obj *patternFlowGtpv1SquenceNumber) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1SquenceNumber) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1SquenceNumber) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1SquenceNumber) Clone() (PatternFlowGtpv1SquenceNumber, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1SquenceNumber()
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

func (obj *patternFlowGtpv1SquenceNumber) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGtpv1SquenceNumber is sequence number. Exists if any of the e_flag, s_flag, or pn_flag bits are on.  Must be interpreted only if the s_flag bit is on.
type PatternFlowGtpv1SquenceNumber interface {
	Validation
	// msg marshals PatternFlowGtpv1SquenceNumber to protobuf object *otg.PatternFlowGtpv1SquenceNumber
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1SquenceNumber
	// setMsg unmarshals PatternFlowGtpv1SquenceNumber from protobuf object *otg.PatternFlowGtpv1SquenceNumber
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1SquenceNumber) PatternFlowGtpv1SquenceNumber
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1SquenceNumber
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1SquenceNumber
	// validate validates PatternFlowGtpv1SquenceNumber
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1SquenceNumber, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGtpv1SquenceNumberChoiceEnum, set in PatternFlowGtpv1SquenceNumber
	Choice() PatternFlowGtpv1SquenceNumberChoiceEnum
	// setChoice assigns PatternFlowGtpv1SquenceNumberChoiceEnum provided by user to PatternFlowGtpv1SquenceNumber
	setChoice(value PatternFlowGtpv1SquenceNumberChoiceEnum) PatternFlowGtpv1SquenceNumber
	// HasChoice checks if Choice has been set in PatternFlowGtpv1SquenceNumber
	HasChoice() bool
	// Value returns uint32, set in PatternFlowGtpv1SquenceNumber.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowGtpv1SquenceNumber
	SetValue(value uint32) PatternFlowGtpv1SquenceNumber
	// HasValue checks if Value has been set in PatternFlowGtpv1SquenceNumber
	HasValue() bool
	// Values returns []uint32, set in PatternFlowGtpv1SquenceNumber.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowGtpv1SquenceNumber
	SetValues(value []uint32) PatternFlowGtpv1SquenceNumber
	// Increment returns PatternFlowGtpv1SquenceNumberCounter, set in PatternFlowGtpv1SquenceNumber.
	// PatternFlowGtpv1SquenceNumberCounter is integer counter pattern
	Increment() PatternFlowGtpv1SquenceNumberCounter
	// SetIncrement assigns PatternFlowGtpv1SquenceNumberCounter provided by user to PatternFlowGtpv1SquenceNumber.
	// PatternFlowGtpv1SquenceNumberCounter is integer counter pattern
	SetIncrement(value PatternFlowGtpv1SquenceNumberCounter) PatternFlowGtpv1SquenceNumber
	// HasIncrement checks if Increment has been set in PatternFlowGtpv1SquenceNumber
	HasIncrement() bool
	// Decrement returns PatternFlowGtpv1SquenceNumberCounter, set in PatternFlowGtpv1SquenceNumber.
	// PatternFlowGtpv1SquenceNumberCounter is integer counter pattern
	Decrement() PatternFlowGtpv1SquenceNumberCounter
	// SetDecrement assigns PatternFlowGtpv1SquenceNumberCounter provided by user to PatternFlowGtpv1SquenceNumber.
	// PatternFlowGtpv1SquenceNumberCounter is integer counter pattern
	SetDecrement(value PatternFlowGtpv1SquenceNumberCounter) PatternFlowGtpv1SquenceNumber
	// HasDecrement checks if Decrement has been set in PatternFlowGtpv1SquenceNumber
	HasDecrement() bool
	// MetricTags returns PatternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIterIter, set in PatternFlowGtpv1SquenceNumber
	MetricTags() PatternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter
	setNil()
}

type PatternFlowGtpv1SquenceNumberChoiceEnum string

// Enum of Choice on PatternFlowGtpv1SquenceNumber
var PatternFlowGtpv1SquenceNumberChoice = struct {
	VALUE     PatternFlowGtpv1SquenceNumberChoiceEnum
	VALUES    PatternFlowGtpv1SquenceNumberChoiceEnum
	INCREMENT PatternFlowGtpv1SquenceNumberChoiceEnum
	DECREMENT PatternFlowGtpv1SquenceNumberChoiceEnum
}{
	VALUE:     PatternFlowGtpv1SquenceNumberChoiceEnum("value"),
	VALUES:    PatternFlowGtpv1SquenceNumberChoiceEnum("values"),
	INCREMENT: PatternFlowGtpv1SquenceNumberChoiceEnum("increment"),
	DECREMENT: PatternFlowGtpv1SquenceNumberChoiceEnum("decrement"),
}

func (obj *patternFlowGtpv1SquenceNumber) Choice() PatternFlowGtpv1SquenceNumberChoiceEnum {
	return PatternFlowGtpv1SquenceNumberChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGtpv1SquenceNumber) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGtpv1SquenceNumber) setChoice(value PatternFlowGtpv1SquenceNumberChoiceEnum) PatternFlowGtpv1SquenceNumber {
	intValue, ok := otg.PatternFlowGtpv1SquenceNumber_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGtpv1SquenceNumberChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGtpv1SquenceNumber_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGtpv1SquenceNumberChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGtpv1SquenceNumberChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGtpv1SquenceNumberChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGtpv1SquenceNumberCounter().msg()
	}

	if value == PatternFlowGtpv1SquenceNumberChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGtpv1SquenceNumberCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv1SquenceNumber) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGtpv1SquenceNumberChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv1SquenceNumber) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowGtpv1SquenceNumber object
func (obj *patternFlowGtpv1SquenceNumber) SetValue(value uint32) PatternFlowGtpv1SquenceNumber {
	obj.setChoice(PatternFlowGtpv1SquenceNumberChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowGtpv1SquenceNumber) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowGtpv1SquenceNumber object
func (obj *patternFlowGtpv1SquenceNumber) SetValues(value []uint32) PatternFlowGtpv1SquenceNumber {
	obj.setChoice(PatternFlowGtpv1SquenceNumberChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGtpv1SquenceNumberCounter
func (obj *patternFlowGtpv1SquenceNumber) Increment() PatternFlowGtpv1SquenceNumberCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGtpv1SquenceNumberChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGtpv1SquenceNumberCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGtpv1SquenceNumberCounter
func (obj *patternFlowGtpv1SquenceNumber) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGtpv1SquenceNumberCounter value in the PatternFlowGtpv1SquenceNumber object
func (obj *patternFlowGtpv1SquenceNumber) SetIncrement(value PatternFlowGtpv1SquenceNumberCounter) PatternFlowGtpv1SquenceNumber {
	obj.setChoice(PatternFlowGtpv1SquenceNumberChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGtpv1SquenceNumberCounter
func (obj *patternFlowGtpv1SquenceNumber) Decrement() PatternFlowGtpv1SquenceNumberCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGtpv1SquenceNumberChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGtpv1SquenceNumberCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGtpv1SquenceNumberCounter
func (obj *patternFlowGtpv1SquenceNumber) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGtpv1SquenceNumberCounter value in the PatternFlowGtpv1SquenceNumber object
func (obj *patternFlowGtpv1SquenceNumber) SetDecrement(value PatternFlowGtpv1SquenceNumberCounter) PatternFlowGtpv1SquenceNumber {
	obj.setChoice(PatternFlowGtpv1SquenceNumberChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGtpv1SquenceNumberMetricTag
func (obj *patternFlowGtpv1SquenceNumber) MetricTags() PatternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGtpv1SquenceNumberMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter struct {
	obj                                         *patternFlowGtpv1SquenceNumber
	patternFlowGtpv1SquenceNumberMetricTagSlice []PatternFlowGtpv1SquenceNumberMetricTag
	fieldPtr                                    *[]*otg.PatternFlowGtpv1SquenceNumberMetricTag
}

func newPatternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter(ptr *[]*otg.PatternFlowGtpv1SquenceNumberMetricTag) PatternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter {
	return &patternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter{fieldPtr: ptr}
}

type PatternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter interface {
	setMsg(*patternFlowGtpv1SquenceNumber) PatternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter
	Items() []PatternFlowGtpv1SquenceNumberMetricTag
	Add() PatternFlowGtpv1SquenceNumberMetricTag
	Append(items ...PatternFlowGtpv1SquenceNumberMetricTag) PatternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter
	Set(index int, newObj PatternFlowGtpv1SquenceNumberMetricTag) PatternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter
	Clear() PatternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter
	clearHolderSlice() PatternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter
	appendHolderSlice(item PatternFlowGtpv1SquenceNumberMetricTag) PatternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter
}

func (obj *patternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter) setMsg(msg *patternFlowGtpv1SquenceNumber) PatternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGtpv1SquenceNumberMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter) Items() []PatternFlowGtpv1SquenceNumberMetricTag {
	return obj.patternFlowGtpv1SquenceNumberMetricTagSlice
}

func (obj *patternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter) Add() PatternFlowGtpv1SquenceNumberMetricTag {
	newObj := &otg.PatternFlowGtpv1SquenceNumberMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGtpv1SquenceNumberMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGtpv1SquenceNumberMetricTagSlice = append(obj.patternFlowGtpv1SquenceNumberMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter) Append(items ...PatternFlowGtpv1SquenceNumberMetricTag) PatternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGtpv1SquenceNumberMetricTagSlice = append(obj.patternFlowGtpv1SquenceNumberMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter) Set(index int, newObj PatternFlowGtpv1SquenceNumberMetricTag) PatternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGtpv1SquenceNumberMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter) Clear() PatternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGtpv1SquenceNumberMetricTag{}
		obj.patternFlowGtpv1SquenceNumberMetricTagSlice = []PatternFlowGtpv1SquenceNumberMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter) clearHolderSlice() PatternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter {
	if len(obj.patternFlowGtpv1SquenceNumberMetricTagSlice) > 0 {
		obj.patternFlowGtpv1SquenceNumberMetricTagSlice = []PatternFlowGtpv1SquenceNumberMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter) appendHolderSlice(item PatternFlowGtpv1SquenceNumberMetricTag) PatternFlowGtpv1SquenceNumberPatternFlowGtpv1SquenceNumberMetricTagIter {
	obj.patternFlowGtpv1SquenceNumberMetricTagSlice = append(obj.patternFlowGtpv1SquenceNumberMetricTagSlice, item)
	return obj
}

func (obj *patternFlowGtpv1SquenceNumber) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1SquenceNumber.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowGtpv1SquenceNumber.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowGtpv1SquenceNumberMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGtpv1SquenceNumber) setDefault() {
	var choices_set int = 0
	var choice PatternFlowGtpv1SquenceNumberChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowGtpv1SquenceNumberChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowGtpv1SquenceNumberChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowGtpv1SquenceNumberChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowGtpv1SquenceNumberChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowGtpv1SquenceNumberChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowGtpv1SquenceNumber")
			}
		} else {
			intVal := otg.PatternFlowGtpv1SquenceNumber_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowGtpv1SquenceNumber_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
