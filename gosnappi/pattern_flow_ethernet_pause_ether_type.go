package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowEthernetPauseEtherType *****
type patternFlowEthernetPauseEtherType struct {
	validation
	obj              *otg.PatternFlowEthernetPauseEtherType
	marshaller       marshalPatternFlowEthernetPauseEtherType
	unMarshaller     unMarshalPatternFlowEthernetPauseEtherType
	incrementHolder  PatternFlowEthernetPauseEtherTypeCounter
	decrementHolder  PatternFlowEthernetPauseEtherTypeCounter
	metricTagsHolder PatternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter
}

func NewPatternFlowEthernetPauseEtherType() PatternFlowEthernetPauseEtherType {
	obj := patternFlowEthernetPauseEtherType{obj: &otg.PatternFlowEthernetPauseEtherType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowEthernetPauseEtherType) msg() *otg.PatternFlowEthernetPauseEtherType {
	return obj.obj
}

func (obj *patternFlowEthernetPauseEtherType) setMsg(msg *otg.PatternFlowEthernetPauseEtherType) PatternFlowEthernetPauseEtherType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowEthernetPauseEtherType struct {
	obj *patternFlowEthernetPauseEtherType
}

type marshalPatternFlowEthernetPauseEtherType interface {
	// ToProto marshals PatternFlowEthernetPauseEtherType to protobuf object *otg.PatternFlowEthernetPauseEtherType
	ToProto() (*otg.PatternFlowEthernetPauseEtherType, error)
	// ToPbText marshals PatternFlowEthernetPauseEtherType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowEthernetPauseEtherType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowEthernetPauseEtherType to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowEthernetPauseEtherType to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowEthernetPauseEtherType struct {
	obj *patternFlowEthernetPauseEtherType
}

type unMarshalPatternFlowEthernetPauseEtherType interface {
	// FromProto unmarshals PatternFlowEthernetPauseEtherType from protobuf object *otg.PatternFlowEthernetPauseEtherType
	FromProto(msg *otg.PatternFlowEthernetPauseEtherType) (PatternFlowEthernetPauseEtherType, error)
	// FromPbText unmarshals PatternFlowEthernetPauseEtherType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowEthernetPauseEtherType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowEthernetPauseEtherType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowEthernetPauseEtherType) Marshal() marshalPatternFlowEthernetPauseEtherType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowEthernetPauseEtherType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowEthernetPauseEtherType) Unmarshal() unMarshalPatternFlowEthernetPauseEtherType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowEthernetPauseEtherType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowEthernetPauseEtherType) ToProto() (*otg.PatternFlowEthernetPauseEtherType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowEthernetPauseEtherType) FromProto(msg *otg.PatternFlowEthernetPauseEtherType) (PatternFlowEthernetPauseEtherType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowEthernetPauseEtherType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseEtherType) FromPbText(value string) error {
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

func (m *marshalpatternFlowEthernetPauseEtherType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseEtherType) FromYaml(value string) error {
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

func (m *marshalpatternFlowEthernetPauseEtherType) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowEthernetPauseEtherType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseEtherType) FromJson(value string) error {
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

func (obj *patternFlowEthernetPauseEtherType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseEtherType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseEtherType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowEthernetPauseEtherType) Clone() (PatternFlowEthernetPauseEtherType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowEthernetPauseEtherType()
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

func (obj *patternFlowEthernetPauseEtherType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowEthernetPauseEtherType is ethernet type
type PatternFlowEthernetPauseEtherType interface {
	Validation
	// msg marshals PatternFlowEthernetPauseEtherType to protobuf object *otg.PatternFlowEthernetPauseEtherType
	// and doesn't set defaults
	msg() *otg.PatternFlowEthernetPauseEtherType
	// setMsg unmarshals PatternFlowEthernetPauseEtherType from protobuf object *otg.PatternFlowEthernetPauseEtherType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowEthernetPauseEtherType) PatternFlowEthernetPauseEtherType
	// provides marshal interface
	Marshal() marshalPatternFlowEthernetPauseEtherType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowEthernetPauseEtherType
	// validate validates PatternFlowEthernetPauseEtherType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowEthernetPauseEtherType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowEthernetPauseEtherTypeChoiceEnum, set in PatternFlowEthernetPauseEtherType
	Choice() PatternFlowEthernetPauseEtherTypeChoiceEnum
	// setChoice assigns PatternFlowEthernetPauseEtherTypeChoiceEnum provided by user to PatternFlowEthernetPauseEtherType
	setChoice(value PatternFlowEthernetPauseEtherTypeChoiceEnum) PatternFlowEthernetPauseEtherType
	// HasChoice checks if Choice has been set in PatternFlowEthernetPauseEtherType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowEthernetPauseEtherType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowEthernetPauseEtherType
	SetValue(value uint32) PatternFlowEthernetPauseEtherType
	// HasValue checks if Value has been set in PatternFlowEthernetPauseEtherType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowEthernetPauseEtherType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowEthernetPauseEtherType
	SetValues(value []uint32) PatternFlowEthernetPauseEtherType
	// Increment returns PatternFlowEthernetPauseEtherTypeCounter, set in PatternFlowEthernetPauseEtherType.
	// PatternFlowEthernetPauseEtherTypeCounter is integer counter pattern
	Increment() PatternFlowEthernetPauseEtherTypeCounter
	// SetIncrement assigns PatternFlowEthernetPauseEtherTypeCounter provided by user to PatternFlowEthernetPauseEtherType.
	// PatternFlowEthernetPauseEtherTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowEthernetPauseEtherTypeCounter) PatternFlowEthernetPauseEtherType
	// HasIncrement checks if Increment has been set in PatternFlowEthernetPauseEtherType
	HasIncrement() bool
	// Decrement returns PatternFlowEthernetPauseEtherTypeCounter, set in PatternFlowEthernetPauseEtherType.
	// PatternFlowEthernetPauseEtherTypeCounter is integer counter pattern
	Decrement() PatternFlowEthernetPauseEtherTypeCounter
	// SetDecrement assigns PatternFlowEthernetPauseEtherTypeCounter provided by user to PatternFlowEthernetPauseEtherType.
	// PatternFlowEthernetPauseEtherTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowEthernetPauseEtherTypeCounter) PatternFlowEthernetPauseEtherType
	// HasDecrement checks if Decrement has been set in PatternFlowEthernetPauseEtherType
	HasDecrement() bool
	// MetricTags returns PatternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIterIter, set in PatternFlowEthernetPauseEtherType
	MetricTags() PatternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter
	setNil()
}

type PatternFlowEthernetPauseEtherTypeChoiceEnum string

// Enum of Choice on PatternFlowEthernetPauseEtherType
var PatternFlowEthernetPauseEtherTypeChoice = struct {
	VALUE     PatternFlowEthernetPauseEtherTypeChoiceEnum
	VALUES    PatternFlowEthernetPauseEtherTypeChoiceEnum
	INCREMENT PatternFlowEthernetPauseEtherTypeChoiceEnum
	DECREMENT PatternFlowEthernetPauseEtherTypeChoiceEnum
}{
	VALUE:     PatternFlowEthernetPauseEtherTypeChoiceEnum("value"),
	VALUES:    PatternFlowEthernetPauseEtherTypeChoiceEnum("values"),
	INCREMENT: PatternFlowEthernetPauseEtherTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowEthernetPauseEtherTypeChoiceEnum("decrement"),
}

func (obj *patternFlowEthernetPauseEtherType) Choice() PatternFlowEthernetPauseEtherTypeChoiceEnum {
	return PatternFlowEthernetPauseEtherTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowEthernetPauseEtherType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowEthernetPauseEtherType) setChoice(value PatternFlowEthernetPauseEtherTypeChoiceEnum) PatternFlowEthernetPauseEtherType {
	intValue, ok := otg.PatternFlowEthernetPauseEtherType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowEthernetPauseEtherTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowEthernetPauseEtherType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowEthernetPauseEtherTypeChoice.VALUE {
		defaultValue := uint32(34824)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowEthernetPauseEtherTypeChoice.VALUES {
		defaultValue := []uint32{34824}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowEthernetPauseEtherTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowEthernetPauseEtherTypeCounter().msg()
	}

	if value == PatternFlowEthernetPauseEtherTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowEthernetPauseEtherTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowEthernetPauseEtherType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowEthernetPauseEtherTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowEthernetPauseEtherType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowEthernetPauseEtherType object
func (obj *patternFlowEthernetPauseEtherType) SetValue(value uint32) PatternFlowEthernetPauseEtherType {
	obj.setChoice(PatternFlowEthernetPauseEtherTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowEthernetPauseEtherType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{34824})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowEthernetPauseEtherType object
func (obj *patternFlowEthernetPauseEtherType) SetValues(value []uint32) PatternFlowEthernetPauseEtherType {
	obj.setChoice(PatternFlowEthernetPauseEtherTypeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowEthernetPauseEtherTypeCounter
func (obj *patternFlowEthernetPauseEtherType) Increment() PatternFlowEthernetPauseEtherTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowEthernetPauseEtherTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowEthernetPauseEtherTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowEthernetPauseEtherTypeCounter
func (obj *patternFlowEthernetPauseEtherType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowEthernetPauseEtherTypeCounter value in the PatternFlowEthernetPauseEtherType object
func (obj *patternFlowEthernetPauseEtherType) SetIncrement(value PatternFlowEthernetPauseEtherTypeCounter) PatternFlowEthernetPauseEtherType {
	obj.setChoice(PatternFlowEthernetPauseEtherTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowEthernetPauseEtherTypeCounter
func (obj *patternFlowEthernetPauseEtherType) Decrement() PatternFlowEthernetPauseEtherTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowEthernetPauseEtherTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowEthernetPauseEtherTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowEthernetPauseEtherTypeCounter
func (obj *patternFlowEthernetPauseEtherType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowEthernetPauseEtherTypeCounter value in the PatternFlowEthernetPauseEtherType object
func (obj *patternFlowEthernetPauseEtherType) SetDecrement(value PatternFlowEthernetPauseEtherTypeCounter) PatternFlowEthernetPauseEtherType {
	obj.setChoice(PatternFlowEthernetPauseEtherTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowEthernetPauseEtherTypeMetricTag
func (obj *patternFlowEthernetPauseEtherType) MetricTags() PatternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowEthernetPauseEtherTypeMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter struct {
	obj                                             *patternFlowEthernetPauseEtherType
	patternFlowEthernetPauseEtherTypeMetricTagSlice []PatternFlowEthernetPauseEtherTypeMetricTag
	fieldPtr                                        *[]*otg.PatternFlowEthernetPauseEtherTypeMetricTag
}

func newPatternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter(ptr *[]*otg.PatternFlowEthernetPauseEtherTypeMetricTag) PatternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter {
	return &patternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter{fieldPtr: ptr}
}

type PatternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter interface {
	setMsg(*patternFlowEthernetPauseEtherType) PatternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter
	Items() []PatternFlowEthernetPauseEtherTypeMetricTag
	Add() PatternFlowEthernetPauseEtherTypeMetricTag
	Append(items ...PatternFlowEthernetPauseEtherTypeMetricTag) PatternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter
	Set(index int, newObj PatternFlowEthernetPauseEtherTypeMetricTag) PatternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter
	Clear() PatternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter
	clearHolderSlice() PatternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter
	appendHolderSlice(item PatternFlowEthernetPauseEtherTypeMetricTag) PatternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter
}

func (obj *patternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter) setMsg(msg *patternFlowEthernetPauseEtherType) PatternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowEthernetPauseEtherTypeMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter) Items() []PatternFlowEthernetPauseEtherTypeMetricTag {
	return obj.patternFlowEthernetPauseEtherTypeMetricTagSlice
}

func (obj *patternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter) Add() PatternFlowEthernetPauseEtherTypeMetricTag {
	newObj := &otg.PatternFlowEthernetPauseEtherTypeMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowEthernetPauseEtherTypeMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowEthernetPauseEtherTypeMetricTagSlice = append(obj.patternFlowEthernetPauseEtherTypeMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter) Append(items ...PatternFlowEthernetPauseEtherTypeMetricTag) PatternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowEthernetPauseEtherTypeMetricTagSlice = append(obj.patternFlowEthernetPauseEtherTypeMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter) Set(index int, newObj PatternFlowEthernetPauseEtherTypeMetricTag) PatternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowEthernetPauseEtherTypeMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter) Clear() PatternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowEthernetPauseEtherTypeMetricTag{}
		obj.patternFlowEthernetPauseEtherTypeMetricTagSlice = []PatternFlowEthernetPauseEtherTypeMetricTag{}
	}
	return obj
}
func (obj *patternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter) clearHolderSlice() PatternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter {
	if len(obj.patternFlowEthernetPauseEtherTypeMetricTagSlice) > 0 {
		obj.patternFlowEthernetPauseEtherTypeMetricTagSlice = []PatternFlowEthernetPauseEtherTypeMetricTag{}
	}
	return obj
}
func (obj *patternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter) appendHolderSlice(item PatternFlowEthernetPauseEtherTypeMetricTag) PatternFlowEthernetPauseEtherTypePatternFlowEthernetPauseEtherTypeMetricTagIter {
	obj.patternFlowEthernetPauseEtherTypeMetricTagSlice = append(obj.patternFlowEthernetPauseEtherTypeMetricTagSlice, item)
	return obj
}

func (obj *patternFlowEthernetPauseEtherType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetPauseEtherType.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowEthernetPauseEtherType.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowEthernetPauseEtherTypeMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowEthernetPauseEtherType) setDefault() {
	var choices_set int = 0
	var choice PatternFlowEthernetPauseEtherTypeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowEthernetPauseEtherTypeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowEthernetPauseEtherTypeChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowEthernetPauseEtherTypeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowEthernetPauseEtherTypeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowEthernetPauseEtherTypeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowEthernetPauseEtherType")
			}
		} else {
			intVal := otg.PatternFlowEthernetPauseEtherType_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowEthernetPauseEtherType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
