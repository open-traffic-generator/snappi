package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPauseEtherType *****
type patternFlowPfcPauseEtherType struct {
	validation
	obj              *otg.PatternFlowPfcPauseEtherType
	marshaller       marshalPatternFlowPfcPauseEtherType
	unMarshaller     unMarshalPatternFlowPfcPauseEtherType
	incrementHolder  PatternFlowPfcPauseEtherTypeCounter
	decrementHolder  PatternFlowPfcPauseEtherTypeCounter
	metricTagsHolder PatternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter
}

func NewPatternFlowPfcPauseEtherType() PatternFlowPfcPauseEtherType {
	obj := patternFlowPfcPauseEtherType{obj: &otg.PatternFlowPfcPauseEtherType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPauseEtherType) msg() *otg.PatternFlowPfcPauseEtherType {
	return obj.obj
}

func (obj *patternFlowPfcPauseEtherType) setMsg(msg *otg.PatternFlowPfcPauseEtherType) PatternFlowPfcPauseEtherType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPauseEtherType struct {
	obj *patternFlowPfcPauseEtherType
}

type marshalPatternFlowPfcPauseEtherType interface {
	// ToProto marshals PatternFlowPfcPauseEtherType to protobuf object *otg.PatternFlowPfcPauseEtherType
	ToProto() (*otg.PatternFlowPfcPauseEtherType, error)
	// ToPbText marshals PatternFlowPfcPauseEtherType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPauseEtherType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPauseEtherType to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPfcPauseEtherType struct {
	obj *patternFlowPfcPauseEtherType
}

type unMarshalPatternFlowPfcPauseEtherType interface {
	// FromProto unmarshals PatternFlowPfcPauseEtherType from protobuf object *otg.PatternFlowPfcPauseEtherType
	FromProto(msg *otg.PatternFlowPfcPauseEtherType) (PatternFlowPfcPauseEtherType, error)
	// FromPbText unmarshals PatternFlowPfcPauseEtherType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPauseEtherType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPauseEtherType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPauseEtherType) Marshal() marshalPatternFlowPfcPauseEtherType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPauseEtherType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPauseEtherType) Unmarshal() unMarshalPatternFlowPfcPauseEtherType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPauseEtherType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPauseEtherType) ToProto() (*otg.PatternFlowPfcPauseEtherType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPauseEtherType) FromProto(msg *otg.PatternFlowPfcPauseEtherType) (PatternFlowPfcPauseEtherType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPauseEtherType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseEtherType) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPauseEtherType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseEtherType) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPauseEtherType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseEtherType) FromJson(value string) error {
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

func (obj *patternFlowPfcPauseEtherType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseEtherType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseEtherType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPauseEtherType) Clone() (PatternFlowPfcPauseEtherType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPauseEtherType()
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

func (obj *patternFlowPfcPauseEtherType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowPfcPauseEtherType is ethernet type
type PatternFlowPfcPauseEtherType interface {
	Validation
	// msg marshals PatternFlowPfcPauseEtherType to protobuf object *otg.PatternFlowPfcPauseEtherType
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPauseEtherType
	// setMsg unmarshals PatternFlowPfcPauseEtherType from protobuf object *otg.PatternFlowPfcPauseEtherType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPauseEtherType) PatternFlowPfcPauseEtherType
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPauseEtherType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPauseEtherType
	// validate validates PatternFlowPfcPauseEtherType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPauseEtherType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowPfcPauseEtherTypeChoiceEnum, set in PatternFlowPfcPauseEtherType
	Choice() PatternFlowPfcPauseEtherTypeChoiceEnum
	// setChoice assigns PatternFlowPfcPauseEtherTypeChoiceEnum provided by user to PatternFlowPfcPauseEtherType
	setChoice(value PatternFlowPfcPauseEtherTypeChoiceEnum) PatternFlowPfcPauseEtherType
	// HasChoice checks if Choice has been set in PatternFlowPfcPauseEtherType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowPfcPauseEtherType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowPfcPauseEtherType
	SetValue(value uint32) PatternFlowPfcPauseEtherType
	// HasValue checks if Value has been set in PatternFlowPfcPauseEtherType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowPfcPauseEtherType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowPfcPauseEtherType
	SetValues(value []uint32) PatternFlowPfcPauseEtherType
	// Increment returns PatternFlowPfcPauseEtherTypeCounter, set in PatternFlowPfcPauseEtherType.
	// PatternFlowPfcPauseEtherTypeCounter is integer counter pattern
	Increment() PatternFlowPfcPauseEtherTypeCounter
	// SetIncrement assigns PatternFlowPfcPauseEtherTypeCounter provided by user to PatternFlowPfcPauseEtherType.
	// PatternFlowPfcPauseEtherTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowPfcPauseEtherTypeCounter) PatternFlowPfcPauseEtherType
	// HasIncrement checks if Increment has been set in PatternFlowPfcPauseEtherType
	HasIncrement() bool
	// Decrement returns PatternFlowPfcPauseEtherTypeCounter, set in PatternFlowPfcPauseEtherType.
	// PatternFlowPfcPauseEtherTypeCounter is integer counter pattern
	Decrement() PatternFlowPfcPauseEtherTypeCounter
	// SetDecrement assigns PatternFlowPfcPauseEtherTypeCounter provided by user to PatternFlowPfcPauseEtherType.
	// PatternFlowPfcPauseEtherTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowPfcPauseEtherTypeCounter) PatternFlowPfcPauseEtherType
	// HasDecrement checks if Decrement has been set in PatternFlowPfcPauseEtherType
	HasDecrement() bool
	// MetricTags returns PatternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIterIter, set in PatternFlowPfcPauseEtherType
	MetricTags() PatternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter
	setNil()
}

type PatternFlowPfcPauseEtherTypeChoiceEnum string

// Enum of Choice on PatternFlowPfcPauseEtherType
var PatternFlowPfcPauseEtherTypeChoice = struct {
	VALUE     PatternFlowPfcPauseEtherTypeChoiceEnum
	VALUES    PatternFlowPfcPauseEtherTypeChoiceEnum
	INCREMENT PatternFlowPfcPauseEtherTypeChoiceEnum
	DECREMENT PatternFlowPfcPauseEtherTypeChoiceEnum
}{
	VALUE:     PatternFlowPfcPauseEtherTypeChoiceEnum("value"),
	VALUES:    PatternFlowPfcPauseEtherTypeChoiceEnum("values"),
	INCREMENT: PatternFlowPfcPauseEtherTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowPfcPauseEtherTypeChoiceEnum("decrement"),
}

func (obj *patternFlowPfcPauseEtherType) Choice() PatternFlowPfcPauseEtherTypeChoiceEnum {
	return PatternFlowPfcPauseEtherTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowPfcPauseEtherType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowPfcPauseEtherType) setChoice(value PatternFlowPfcPauseEtherTypeChoiceEnum) PatternFlowPfcPauseEtherType {
	intValue, ok := otg.PatternFlowPfcPauseEtherType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowPfcPauseEtherTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowPfcPauseEtherType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowPfcPauseEtherTypeChoice.VALUE {
		defaultValue := uint32(34824)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowPfcPauseEtherTypeChoice.VALUES {
		defaultValue := []uint32{34824}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowPfcPauseEtherTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowPfcPauseEtherTypeCounter().msg()
	}

	if value == PatternFlowPfcPauseEtherTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowPfcPauseEtherTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowPfcPauseEtherType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowPfcPauseEtherTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowPfcPauseEtherType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowPfcPauseEtherType object
func (obj *patternFlowPfcPauseEtherType) SetValue(value uint32) PatternFlowPfcPauseEtherType {
	obj.setChoice(PatternFlowPfcPauseEtherTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowPfcPauseEtherType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{34824})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowPfcPauseEtherType object
func (obj *patternFlowPfcPauseEtherType) SetValues(value []uint32) PatternFlowPfcPauseEtherType {
	obj.setChoice(PatternFlowPfcPauseEtherTypeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowPfcPauseEtherTypeCounter
func (obj *patternFlowPfcPauseEtherType) Increment() PatternFlowPfcPauseEtherTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowPfcPauseEtherTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowPfcPauseEtherTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowPfcPauseEtherTypeCounter
func (obj *patternFlowPfcPauseEtherType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowPfcPauseEtherTypeCounter value in the PatternFlowPfcPauseEtherType object
func (obj *patternFlowPfcPauseEtherType) SetIncrement(value PatternFlowPfcPauseEtherTypeCounter) PatternFlowPfcPauseEtherType {
	obj.setChoice(PatternFlowPfcPauseEtherTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowPfcPauseEtherTypeCounter
func (obj *patternFlowPfcPauseEtherType) Decrement() PatternFlowPfcPauseEtherTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowPfcPauseEtherTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowPfcPauseEtherTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowPfcPauseEtherTypeCounter
func (obj *patternFlowPfcPauseEtherType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowPfcPauseEtherTypeCounter value in the PatternFlowPfcPauseEtherType object
func (obj *patternFlowPfcPauseEtherType) SetDecrement(value PatternFlowPfcPauseEtherTypeCounter) PatternFlowPfcPauseEtherType {
	obj.setChoice(PatternFlowPfcPauseEtherTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowPfcPauseEtherTypeMetricTag
func (obj *patternFlowPfcPauseEtherType) MetricTags() PatternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowPfcPauseEtherTypeMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter struct {
	obj                                        *patternFlowPfcPauseEtherType
	patternFlowPfcPauseEtherTypeMetricTagSlice []PatternFlowPfcPauseEtherTypeMetricTag
	fieldPtr                                   *[]*otg.PatternFlowPfcPauseEtherTypeMetricTag
}

func newPatternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter(ptr *[]*otg.PatternFlowPfcPauseEtherTypeMetricTag) PatternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter {
	return &patternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter{fieldPtr: ptr}
}

type PatternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter interface {
	setMsg(*patternFlowPfcPauseEtherType) PatternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter
	Items() []PatternFlowPfcPauseEtherTypeMetricTag
	Add() PatternFlowPfcPauseEtherTypeMetricTag
	Append(items ...PatternFlowPfcPauseEtherTypeMetricTag) PatternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter
	Set(index int, newObj PatternFlowPfcPauseEtherTypeMetricTag) PatternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter
	Clear() PatternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter
	clearHolderSlice() PatternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter
	appendHolderSlice(item PatternFlowPfcPauseEtherTypeMetricTag) PatternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter
}

func (obj *patternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter) setMsg(msg *patternFlowPfcPauseEtherType) PatternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowPfcPauseEtherTypeMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter) Items() []PatternFlowPfcPauseEtherTypeMetricTag {
	return obj.patternFlowPfcPauseEtherTypeMetricTagSlice
}

func (obj *patternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter) Add() PatternFlowPfcPauseEtherTypeMetricTag {
	newObj := &otg.PatternFlowPfcPauseEtherTypeMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowPfcPauseEtherTypeMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowPfcPauseEtherTypeMetricTagSlice = append(obj.patternFlowPfcPauseEtherTypeMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter) Append(items ...PatternFlowPfcPauseEtherTypeMetricTag) PatternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowPfcPauseEtherTypeMetricTagSlice = append(obj.patternFlowPfcPauseEtherTypeMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter) Set(index int, newObj PatternFlowPfcPauseEtherTypeMetricTag) PatternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowPfcPauseEtherTypeMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter) Clear() PatternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowPfcPauseEtherTypeMetricTag{}
		obj.patternFlowPfcPauseEtherTypeMetricTagSlice = []PatternFlowPfcPauseEtherTypeMetricTag{}
	}
	return obj
}
func (obj *patternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter) clearHolderSlice() PatternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter {
	if len(obj.patternFlowPfcPauseEtherTypeMetricTagSlice) > 0 {
		obj.patternFlowPfcPauseEtherTypeMetricTagSlice = []PatternFlowPfcPauseEtherTypeMetricTag{}
	}
	return obj
}
func (obj *patternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter) appendHolderSlice(item PatternFlowPfcPauseEtherTypeMetricTag) PatternFlowPfcPauseEtherTypePatternFlowPfcPauseEtherTypeMetricTagIter {
	obj.patternFlowPfcPauseEtherTypeMetricTagSlice = append(obj.patternFlowPfcPauseEtherTypeMetricTagSlice, item)
	return obj
}

func (obj *patternFlowPfcPauseEtherType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPauseEtherType.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowPfcPauseEtherType.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowPfcPauseEtherTypeMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowPfcPauseEtherType) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowPfcPauseEtherTypeChoice.VALUE)

	}

}
