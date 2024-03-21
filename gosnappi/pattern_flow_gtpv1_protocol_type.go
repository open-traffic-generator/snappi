package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1ProtocolType *****
type patternFlowGtpv1ProtocolType struct {
	validation
	obj              *otg.PatternFlowGtpv1ProtocolType
	marshaller       marshalPatternFlowGtpv1ProtocolType
	unMarshaller     unMarshalPatternFlowGtpv1ProtocolType
	incrementHolder  PatternFlowGtpv1ProtocolTypeCounter
	decrementHolder  PatternFlowGtpv1ProtocolTypeCounter
	metricTagsHolder PatternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter
}

func NewPatternFlowGtpv1ProtocolType() PatternFlowGtpv1ProtocolType {
	obj := patternFlowGtpv1ProtocolType{obj: &otg.PatternFlowGtpv1ProtocolType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1ProtocolType) msg() *otg.PatternFlowGtpv1ProtocolType {
	return obj.obj
}

func (obj *patternFlowGtpv1ProtocolType) setMsg(msg *otg.PatternFlowGtpv1ProtocolType) PatternFlowGtpv1ProtocolType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1ProtocolType struct {
	obj *patternFlowGtpv1ProtocolType
}

type marshalPatternFlowGtpv1ProtocolType interface {
	// ToProto marshals PatternFlowGtpv1ProtocolType to protobuf object *otg.PatternFlowGtpv1ProtocolType
	ToProto() (*otg.PatternFlowGtpv1ProtocolType, error)
	// ToPbText marshals PatternFlowGtpv1ProtocolType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1ProtocolType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1ProtocolType to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv1ProtocolType struct {
	obj *patternFlowGtpv1ProtocolType
}

type unMarshalPatternFlowGtpv1ProtocolType interface {
	// FromProto unmarshals PatternFlowGtpv1ProtocolType from protobuf object *otg.PatternFlowGtpv1ProtocolType
	FromProto(msg *otg.PatternFlowGtpv1ProtocolType) (PatternFlowGtpv1ProtocolType, error)
	// FromPbText unmarshals PatternFlowGtpv1ProtocolType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1ProtocolType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1ProtocolType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1ProtocolType) Marshal() marshalPatternFlowGtpv1ProtocolType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1ProtocolType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1ProtocolType) Unmarshal() unMarshalPatternFlowGtpv1ProtocolType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1ProtocolType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1ProtocolType) ToProto() (*otg.PatternFlowGtpv1ProtocolType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1ProtocolType) FromProto(msg *otg.PatternFlowGtpv1ProtocolType) (PatternFlowGtpv1ProtocolType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1ProtocolType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1ProtocolType) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1ProtocolType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1ProtocolType) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1ProtocolType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1ProtocolType) FromJson(value string) error {
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

func (obj *patternFlowGtpv1ProtocolType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1ProtocolType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1ProtocolType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1ProtocolType) Clone() (PatternFlowGtpv1ProtocolType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1ProtocolType()
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

func (obj *patternFlowGtpv1ProtocolType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGtpv1ProtocolType is protocol type, GTP is 1, GTP' is 0
type PatternFlowGtpv1ProtocolType interface {
	Validation
	// msg marshals PatternFlowGtpv1ProtocolType to protobuf object *otg.PatternFlowGtpv1ProtocolType
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1ProtocolType
	// setMsg unmarshals PatternFlowGtpv1ProtocolType from protobuf object *otg.PatternFlowGtpv1ProtocolType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1ProtocolType) PatternFlowGtpv1ProtocolType
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1ProtocolType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1ProtocolType
	// validate validates PatternFlowGtpv1ProtocolType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1ProtocolType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGtpv1ProtocolTypeChoiceEnum, set in PatternFlowGtpv1ProtocolType
	Choice() PatternFlowGtpv1ProtocolTypeChoiceEnum
	// setChoice assigns PatternFlowGtpv1ProtocolTypeChoiceEnum provided by user to PatternFlowGtpv1ProtocolType
	setChoice(value PatternFlowGtpv1ProtocolTypeChoiceEnum) PatternFlowGtpv1ProtocolType
	// HasChoice checks if Choice has been set in PatternFlowGtpv1ProtocolType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowGtpv1ProtocolType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowGtpv1ProtocolType
	SetValue(value uint32) PatternFlowGtpv1ProtocolType
	// HasValue checks if Value has been set in PatternFlowGtpv1ProtocolType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowGtpv1ProtocolType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowGtpv1ProtocolType
	SetValues(value []uint32) PatternFlowGtpv1ProtocolType
	// Increment returns PatternFlowGtpv1ProtocolTypeCounter, set in PatternFlowGtpv1ProtocolType.
	// PatternFlowGtpv1ProtocolTypeCounter is integer counter pattern
	Increment() PatternFlowGtpv1ProtocolTypeCounter
	// SetIncrement assigns PatternFlowGtpv1ProtocolTypeCounter provided by user to PatternFlowGtpv1ProtocolType.
	// PatternFlowGtpv1ProtocolTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowGtpv1ProtocolTypeCounter) PatternFlowGtpv1ProtocolType
	// HasIncrement checks if Increment has been set in PatternFlowGtpv1ProtocolType
	HasIncrement() bool
	// Decrement returns PatternFlowGtpv1ProtocolTypeCounter, set in PatternFlowGtpv1ProtocolType.
	// PatternFlowGtpv1ProtocolTypeCounter is integer counter pattern
	Decrement() PatternFlowGtpv1ProtocolTypeCounter
	// SetDecrement assigns PatternFlowGtpv1ProtocolTypeCounter provided by user to PatternFlowGtpv1ProtocolType.
	// PatternFlowGtpv1ProtocolTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowGtpv1ProtocolTypeCounter) PatternFlowGtpv1ProtocolType
	// HasDecrement checks if Decrement has been set in PatternFlowGtpv1ProtocolType
	HasDecrement() bool
	// MetricTags returns PatternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIterIter, set in PatternFlowGtpv1ProtocolType
	MetricTags() PatternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter
	setNil()
}

type PatternFlowGtpv1ProtocolTypeChoiceEnum string

// Enum of Choice on PatternFlowGtpv1ProtocolType
var PatternFlowGtpv1ProtocolTypeChoice = struct {
	VALUE     PatternFlowGtpv1ProtocolTypeChoiceEnum
	VALUES    PatternFlowGtpv1ProtocolTypeChoiceEnum
	INCREMENT PatternFlowGtpv1ProtocolTypeChoiceEnum
	DECREMENT PatternFlowGtpv1ProtocolTypeChoiceEnum
}{
	VALUE:     PatternFlowGtpv1ProtocolTypeChoiceEnum("value"),
	VALUES:    PatternFlowGtpv1ProtocolTypeChoiceEnum("values"),
	INCREMENT: PatternFlowGtpv1ProtocolTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowGtpv1ProtocolTypeChoiceEnum("decrement"),
}

func (obj *patternFlowGtpv1ProtocolType) Choice() PatternFlowGtpv1ProtocolTypeChoiceEnum {
	return PatternFlowGtpv1ProtocolTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGtpv1ProtocolType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGtpv1ProtocolType) setChoice(value PatternFlowGtpv1ProtocolTypeChoiceEnum) PatternFlowGtpv1ProtocolType {
	intValue, ok := otg.PatternFlowGtpv1ProtocolType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGtpv1ProtocolTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGtpv1ProtocolType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGtpv1ProtocolTypeChoice.VALUE {
		defaultValue := uint32(1)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGtpv1ProtocolTypeChoice.VALUES {
		defaultValue := []uint32{1}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGtpv1ProtocolTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGtpv1ProtocolTypeCounter().msg()
	}

	if value == PatternFlowGtpv1ProtocolTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGtpv1ProtocolTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv1ProtocolType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGtpv1ProtocolTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv1ProtocolType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowGtpv1ProtocolType object
func (obj *patternFlowGtpv1ProtocolType) SetValue(value uint32) PatternFlowGtpv1ProtocolType {
	obj.setChoice(PatternFlowGtpv1ProtocolTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowGtpv1ProtocolType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{1})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowGtpv1ProtocolType object
func (obj *patternFlowGtpv1ProtocolType) SetValues(value []uint32) PatternFlowGtpv1ProtocolType {
	obj.setChoice(PatternFlowGtpv1ProtocolTypeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGtpv1ProtocolTypeCounter
func (obj *patternFlowGtpv1ProtocolType) Increment() PatternFlowGtpv1ProtocolTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGtpv1ProtocolTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGtpv1ProtocolTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGtpv1ProtocolTypeCounter
func (obj *patternFlowGtpv1ProtocolType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGtpv1ProtocolTypeCounter value in the PatternFlowGtpv1ProtocolType object
func (obj *patternFlowGtpv1ProtocolType) SetIncrement(value PatternFlowGtpv1ProtocolTypeCounter) PatternFlowGtpv1ProtocolType {
	obj.setChoice(PatternFlowGtpv1ProtocolTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGtpv1ProtocolTypeCounter
func (obj *patternFlowGtpv1ProtocolType) Decrement() PatternFlowGtpv1ProtocolTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGtpv1ProtocolTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGtpv1ProtocolTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGtpv1ProtocolTypeCounter
func (obj *patternFlowGtpv1ProtocolType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGtpv1ProtocolTypeCounter value in the PatternFlowGtpv1ProtocolType object
func (obj *patternFlowGtpv1ProtocolType) SetDecrement(value PatternFlowGtpv1ProtocolTypeCounter) PatternFlowGtpv1ProtocolType {
	obj.setChoice(PatternFlowGtpv1ProtocolTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGtpv1ProtocolTypeMetricTag
func (obj *patternFlowGtpv1ProtocolType) MetricTags() PatternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGtpv1ProtocolTypeMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter struct {
	obj                                        *patternFlowGtpv1ProtocolType
	patternFlowGtpv1ProtocolTypeMetricTagSlice []PatternFlowGtpv1ProtocolTypeMetricTag
	fieldPtr                                   *[]*otg.PatternFlowGtpv1ProtocolTypeMetricTag
}

func newPatternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter(ptr *[]*otg.PatternFlowGtpv1ProtocolTypeMetricTag) PatternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter {
	return &patternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter{fieldPtr: ptr}
}

type PatternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter interface {
	setMsg(*patternFlowGtpv1ProtocolType) PatternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter
	Items() []PatternFlowGtpv1ProtocolTypeMetricTag
	Add() PatternFlowGtpv1ProtocolTypeMetricTag
	Append(items ...PatternFlowGtpv1ProtocolTypeMetricTag) PatternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter
	Set(index int, newObj PatternFlowGtpv1ProtocolTypeMetricTag) PatternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter
	Clear() PatternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter
	clearHolderSlice() PatternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter
	appendHolderSlice(item PatternFlowGtpv1ProtocolTypeMetricTag) PatternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter
}

func (obj *patternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter) setMsg(msg *patternFlowGtpv1ProtocolType) PatternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGtpv1ProtocolTypeMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter) Items() []PatternFlowGtpv1ProtocolTypeMetricTag {
	return obj.patternFlowGtpv1ProtocolTypeMetricTagSlice
}

func (obj *patternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter) Add() PatternFlowGtpv1ProtocolTypeMetricTag {
	newObj := &otg.PatternFlowGtpv1ProtocolTypeMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGtpv1ProtocolTypeMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGtpv1ProtocolTypeMetricTagSlice = append(obj.patternFlowGtpv1ProtocolTypeMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter) Append(items ...PatternFlowGtpv1ProtocolTypeMetricTag) PatternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGtpv1ProtocolTypeMetricTagSlice = append(obj.patternFlowGtpv1ProtocolTypeMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter) Set(index int, newObj PatternFlowGtpv1ProtocolTypeMetricTag) PatternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGtpv1ProtocolTypeMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter) Clear() PatternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGtpv1ProtocolTypeMetricTag{}
		obj.patternFlowGtpv1ProtocolTypeMetricTagSlice = []PatternFlowGtpv1ProtocolTypeMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter) clearHolderSlice() PatternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter {
	if len(obj.patternFlowGtpv1ProtocolTypeMetricTagSlice) > 0 {
		obj.patternFlowGtpv1ProtocolTypeMetricTagSlice = []PatternFlowGtpv1ProtocolTypeMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter) appendHolderSlice(item PatternFlowGtpv1ProtocolTypeMetricTag) PatternFlowGtpv1ProtocolTypePatternFlowGtpv1ProtocolTypeMetricTagIter {
	obj.patternFlowGtpv1ProtocolTypeMetricTagSlice = append(obj.patternFlowGtpv1ProtocolTypeMetricTagSlice, item)
	return obj
}

func (obj *patternFlowGtpv1ProtocolType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1ProtocolType.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowGtpv1ProtocolType.Values <= 1 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowGtpv1ProtocolTypeMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGtpv1ProtocolType) setDefault() {
	var choices_set int = 0
	var choice PatternFlowGtpv1ProtocolTypeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowGtpv1ProtocolTypeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowGtpv1ProtocolTypeChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowGtpv1ProtocolTypeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowGtpv1ProtocolTypeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowGtpv1ProtocolTypeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowGtpv1ProtocolType")
			}
		} else {
			intVal := otg.PatternFlowGtpv1ProtocolType_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowGtpv1ProtocolType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
