package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowEthernetEtherType *****
type patternFlowEthernetEtherType struct {
	validation
	obj              *otg.PatternFlowEthernetEtherType
	marshaller       marshalPatternFlowEthernetEtherType
	unMarshaller     unMarshalPatternFlowEthernetEtherType
	incrementHolder  PatternFlowEthernetEtherTypeCounter
	decrementHolder  PatternFlowEthernetEtherTypeCounter
	metricTagsHolder PatternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter
}

func NewPatternFlowEthernetEtherType() PatternFlowEthernetEtherType {
	obj := patternFlowEthernetEtherType{obj: &otg.PatternFlowEthernetEtherType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowEthernetEtherType) msg() *otg.PatternFlowEthernetEtherType {
	return obj.obj
}

func (obj *patternFlowEthernetEtherType) setMsg(msg *otg.PatternFlowEthernetEtherType) PatternFlowEthernetEtherType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowEthernetEtherType struct {
	obj *patternFlowEthernetEtherType
}

type marshalPatternFlowEthernetEtherType interface {
	// ToProto marshals PatternFlowEthernetEtherType to protobuf object *otg.PatternFlowEthernetEtherType
	ToProto() (*otg.PatternFlowEthernetEtherType, error)
	// ToPbText marshals PatternFlowEthernetEtherType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowEthernetEtherType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowEthernetEtherType to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowEthernetEtherType struct {
	obj *patternFlowEthernetEtherType
}

type unMarshalPatternFlowEthernetEtherType interface {
	// FromProto unmarshals PatternFlowEthernetEtherType from protobuf object *otg.PatternFlowEthernetEtherType
	FromProto(msg *otg.PatternFlowEthernetEtherType) (PatternFlowEthernetEtherType, error)
	// FromPbText unmarshals PatternFlowEthernetEtherType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowEthernetEtherType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowEthernetEtherType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowEthernetEtherType) Marshal() marshalPatternFlowEthernetEtherType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowEthernetEtherType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowEthernetEtherType) Unmarshal() unMarshalPatternFlowEthernetEtherType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowEthernetEtherType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowEthernetEtherType) ToProto() (*otg.PatternFlowEthernetEtherType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowEthernetEtherType) FromProto(msg *otg.PatternFlowEthernetEtherType) (PatternFlowEthernetEtherType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowEthernetEtherType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowEthernetEtherType) FromPbText(value string) error {
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

func (m *marshalpatternFlowEthernetEtherType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowEthernetEtherType) FromYaml(value string) error {
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

func (m *marshalpatternFlowEthernetEtherType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowEthernetEtherType) FromJson(value string) error {
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

func (obj *patternFlowEthernetEtherType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowEthernetEtherType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowEthernetEtherType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowEthernetEtherType) Clone() (PatternFlowEthernetEtherType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowEthernetEtherType()
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

func (obj *patternFlowEthernetEtherType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowEthernetEtherType is ethernet type
type PatternFlowEthernetEtherType interface {
	Validation
	// msg marshals PatternFlowEthernetEtherType to protobuf object *otg.PatternFlowEthernetEtherType
	// and doesn't set defaults
	msg() *otg.PatternFlowEthernetEtherType
	// setMsg unmarshals PatternFlowEthernetEtherType from protobuf object *otg.PatternFlowEthernetEtherType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowEthernetEtherType) PatternFlowEthernetEtherType
	// provides marshal interface
	Marshal() marshalPatternFlowEthernetEtherType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowEthernetEtherType
	// validate validates PatternFlowEthernetEtherType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowEthernetEtherType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowEthernetEtherTypeChoiceEnum, set in PatternFlowEthernetEtherType
	Choice() PatternFlowEthernetEtherTypeChoiceEnum
	// setChoice assigns PatternFlowEthernetEtherTypeChoiceEnum provided by user to PatternFlowEthernetEtherType
	setChoice(value PatternFlowEthernetEtherTypeChoiceEnum) PatternFlowEthernetEtherType
	// HasChoice checks if Choice has been set in PatternFlowEthernetEtherType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowEthernetEtherType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowEthernetEtherType
	SetValue(value uint32) PatternFlowEthernetEtherType
	// HasValue checks if Value has been set in PatternFlowEthernetEtherType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowEthernetEtherType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowEthernetEtherType
	SetValues(value []uint32) PatternFlowEthernetEtherType
	// Auto returns uint32, set in PatternFlowEthernetEtherType.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowEthernetEtherType
	HasAuto() bool
	// Increment returns PatternFlowEthernetEtherTypeCounter, set in PatternFlowEthernetEtherType.
	// PatternFlowEthernetEtherTypeCounter is integer counter pattern
	Increment() PatternFlowEthernetEtherTypeCounter
	// SetIncrement assigns PatternFlowEthernetEtherTypeCounter provided by user to PatternFlowEthernetEtherType.
	// PatternFlowEthernetEtherTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowEthernetEtherTypeCounter) PatternFlowEthernetEtherType
	// HasIncrement checks if Increment has been set in PatternFlowEthernetEtherType
	HasIncrement() bool
	// Decrement returns PatternFlowEthernetEtherTypeCounter, set in PatternFlowEthernetEtherType.
	// PatternFlowEthernetEtherTypeCounter is integer counter pattern
	Decrement() PatternFlowEthernetEtherTypeCounter
	// SetDecrement assigns PatternFlowEthernetEtherTypeCounter provided by user to PatternFlowEthernetEtherType.
	// PatternFlowEthernetEtherTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowEthernetEtherTypeCounter) PatternFlowEthernetEtherType
	// HasDecrement checks if Decrement has been set in PatternFlowEthernetEtherType
	HasDecrement() bool
	// MetricTags returns PatternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIterIter, set in PatternFlowEthernetEtherType
	MetricTags() PatternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter
	setNil()
}

type PatternFlowEthernetEtherTypeChoiceEnum string

// Enum of Choice on PatternFlowEthernetEtherType
var PatternFlowEthernetEtherTypeChoice = struct {
	VALUE     PatternFlowEthernetEtherTypeChoiceEnum
	VALUES    PatternFlowEthernetEtherTypeChoiceEnum
	AUTO      PatternFlowEthernetEtherTypeChoiceEnum
	INCREMENT PatternFlowEthernetEtherTypeChoiceEnum
	DECREMENT PatternFlowEthernetEtherTypeChoiceEnum
}{
	VALUE:     PatternFlowEthernetEtherTypeChoiceEnum("value"),
	VALUES:    PatternFlowEthernetEtherTypeChoiceEnum("values"),
	AUTO:      PatternFlowEthernetEtherTypeChoiceEnum("auto"),
	INCREMENT: PatternFlowEthernetEtherTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowEthernetEtherTypeChoiceEnum("decrement"),
}

func (obj *patternFlowEthernetEtherType) Choice() PatternFlowEthernetEtherTypeChoiceEnum {
	return PatternFlowEthernetEtherTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowEthernetEtherType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowEthernetEtherType) setChoice(value PatternFlowEthernetEtherTypeChoiceEnum) PatternFlowEthernetEtherType {
	intValue, ok := otg.PatternFlowEthernetEtherType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowEthernetEtherTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowEthernetEtherType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowEthernetEtherTypeChoice.VALUE {
		defaultValue := uint32(65535)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowEthernetEtherTypeChoice.VALUES {
		defaultValue := []uint32{65535}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowEthernetEtherTypeChoice.AUTO {
		defaultValue := uint32(65535)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowEthernetEtherTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowEthernetEtherTypeCounter().msg()
	}

	if value == PatternFlowEthernetEtherTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowEthernetEtherTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowEthernetEtherType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowEthernetEtherTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowEthernetEtherType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowEthernetEtherType object
func (obj *patternFlowEthernetEtherType) SetValue(value uint32) PatternFlowEthernetEtherType {
	obj.setChoice(PatternFlowEthernetEtherTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowEthernetEtherType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{65535})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowEthernetEtherType object
func (obj *patternFlowEthernetEtherType) SetValues(value []uint32) PatternFlowEthernetEtherType {
	obj.setChoice(PatternFlowEthernetEtherTypeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowEthernetEtherType) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowEthernetEtherTypeChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowEthernetEtherType) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowEthernetEtherTypeCounter
func (obj *patternFlowEthernetEtherType) Increment() PatternFlowEthernetEtherTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowEthernetEtherTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowEthernetEtherTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowEthernetEtherTypeCounter
func (obj *patternFlowEthernetEtherType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowEthernetEtherTypeCounter value in the PatternFlowEthernetEtherType object
func (obj *patternFlowEthernetEtherType) SetIncrement(value PatternFlowEthernetEtherTypeCounter) PatternFlowEthernetEtherType {
	obj.setChoice(PatternFlowEthernetEtherTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowEthernetEtherTypeCounter
func (obj *patternFlowEthernetEtherType) Decrement() PatternFlowEthernetEtherTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowEthernetEtherTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowEthernetEtherTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowEthernetEtherTypeCounter
func (obj *patternFlowEthernetEtherType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowEthernetEtherTypeCounter value in the PatternFlowEthernetEtherType object
func (obj *patternFlowEthernetEtherType) SetDecrement(value PatternFlowEthernetEtherTypeCounter) PatternFlowEthernetEtherType {
	obj.setChoice(PatternFlowEthernetEtherTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowEthernetEtherTypeMetricTag
func (obj *patternFlowEthernetEtherType) MetricTags() PatternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowEthernetEtherTypeMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter struct {
	obj                                        *patternFlowEthernetEtherType
	patternFlowEthernetEtherTypeMetricTagSlice []PatternFlowEthernetEtherTypeMetricTag
	fieldPtr                                   *[]*otg.PatternFlowEthernetEtherTypeMetricTag
}

func newPatternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter(ptr *[]*otg.PatternFlowEthernetEtherTypeMetricTag) PatternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter {
	return &patternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter{fieldPtr: ptr}
}

type PatternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter interface {
	setMsg(*patternFlowEthernetEtherType) PatternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter
	Items() []PatternFlowEthernetEtherTypeMetricTag
	Add() PatternFlowEthernetEtherTypeMetricTag
	Append(items ...PatternFlowEthernetEtherTypeMetricTag) PatternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter
	Set(index int, newObj PatternFlowEthernetEtherTypeMetricTag) PatternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter
	Clear() PatternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter
	clearHolderSlice() PatternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter
	appendHolderSlice(item PatternFlowEthernetEtherTypeMetricTag) PatternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter
}

func (obj *patternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter) setMsg(msg *patternFlowEthernetEtherType) PatternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowEthernetEtherTypeMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter) Items() []PatternFlowEthernetEtherTypeMetricTag {
	return obj.patternFlowEthernetEtherTypeMetricTagSlice
}

func (obj *patternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter) Add() PatternFlowEthernetEtherTypeMetricTag {
	newObj := &otg.PatternFlowEthernetEtherTypeMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowEthernetEtherTypeMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowEthernetEtherTypeMetricTagSlice = append(obj.patternFlowEthernetEtherTypeMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter) Append(items ...PatternFlowEthernetEtherTypeMetricTag) PatternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowEthernetEtherTypeMetricTagSlice = append(obj.patternFlowEthernetEtherTypeMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter) Set(index int, newObj PatternFlowEthernetEtherTypeMetricTag) PatternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowEthernetEtherTypeMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter) Clear() PatternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowEthernetEtherTypeMetricTag{}
		obj.patternFlowEthernetEtherTypeMetricTagSlice = []PatternFlowEthernetEtherTypeMetricTag{}
	}
	return obj
}
func (obj *patternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter) clearHolderSlice() PatternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter {
	if len(obj.patternFlowEthernetEtherTypeMetricTagSlice) > 0 {
		obj.patternFlowEthernetEtherTypeMetricTagSlice = []PatternFlowEthernetEtherTypeMetricTag{}
	}
	return obj
}
func (obj *patternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter) appendHolderSlice(item PatternFlowEthernetEtherTypeMetricTag) PatternFlowEthernetEtherTypePatternFlowEthernetEtherTypeMetricTagIter {
	obj.patternFlowEthernetEtherTypeMetricTagSlice = append(obj.patternFlowEthernetEtherTypeMetricTagSlice, item)
	return obj
}

func (obj *patternFlowEthernetEtherType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetEtherType.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowEthernetEtherType.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetEtherType.Auto <= 65535 but Got %d", *obj.obj.Auto))
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
				obj.MetricTags().appendHolderSlice(&patternFlowEthernetEtherTypeMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowEthernetEtherType) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowEthernetEtherTypeChoice.AUTO)

	}

}
