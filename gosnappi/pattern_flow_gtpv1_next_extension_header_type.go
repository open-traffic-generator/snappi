package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1NextExtensionHeaderType *****
type patternFlowGtpv1NextExtensionHeaderType struct {
	validation
	obj              *otg.PatternFlowGtpv1NextExtensionHeaderType
	marshaller       marshalPatternFlowGtpv1NextExtensionHeaderType
	unMarshaller     unMarshalPatternFlowGtpv1NextExtensionHeaderType
	incrementHolder  PatternFlowGtpv1NextExtensionHeaderTypeCounter
	decrementHolder  PatternFlowGtpv1NextExtensionHeaderTypeCounter
	metricTagsHolder PatternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter
}

func NewPatternFlowGtpv1NextExtensionHeaderType() PatternFlowGtpv1NextExtensionHeaderType {
	obj := patternFlowGtpv1NextExtensionHeaderType{obj: &otg.PatternFlowGtpv1NextExtensionHeaderType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1NextExtensionHeaderType) msg() *otg.PatternFlowGtpv1NextExtensionHeaderType {
	return obj.obj
}

func (obj *patternFlowGtpv1NextExtensionHeaderType) setMsg(msg *otg.PatternFlowGtpv1NextExtensionHeaderType) PatternFlowGtpv1NextExtensionHeaderType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1NextExtensionHeaderType struct {
	obj *patternFlowGtpv1NextExtensionHeaderType
}

type marshalPatternFlowGtpv1NextExtensionHeaderType interface {
	// ToProto marshals PatternFlowGtpv1NextExtensionHeaderType to protobuf object *otg.PatternFlowGtpv1NextExtensionHeaderType
	ToProto() (*otg.PatternFlowGtpv1NextExtensionHeaderType, error)
	// ToPbText marshals PatternFlowGtpv1NextExtensionHeaderType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1NextExtensionHeaderType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1NextExtensionHeaderType to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv1NextExtensionHeaderType struct {
	obj *patternFlowGtpv1NextExtensionHeaderType
}

type unMarshalPatternFlowGtpv1NextExtensionHeaderType interface {
	// FromProto unmarshals PatternFlowGtpv1NextExtensionHeaderType from protobuf object *otg.PatternFlowGtpv1NextExtensionHeaderType
	FromProto(msg *otg.PatternFlowGtpv1NextExtensionHeaderType) (PatternFlowGtpv1NextExtensionHeaderType, error)
	// FromPbText unmarshals PatternFlowGtpv1NextExtensionHeaderType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1NextExtensionHeaderType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1NextExtensionHeaderType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1NextExtensionHeaderType) Marshal() marshalPatternFlowGtpv1NextExtensionHeaderType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1NextExtensionHeaderType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1NextExtensionHeaderType) Unmarshal() unMarshalPatternFlowGtpv1NextExtensionHeaderType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1NextExtensionHeaderType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1NextExtensionHeaderType) ToProto() (*otg.PatternFlowGtpv1NextExtensionHeaderType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1NextExtensionHeaderType) FromProto(msg *otg.PatternFlowGtpv1NextExtensionHeaderType) (PatternFlowGtpv1NextExtensionHeaderType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1NextExtensionHeaderType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1NextExtensionHeaderType) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1NextExtensionHeaderType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1NextExtensionHeaderType) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1NextExtensionHeaderType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1NextExtensionHeaderType) FromJson(value string) error {
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

func (obj *patternFlowGtpv1NextExtensionHeaderType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1NextExtensionHeaderType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1NextExtensionHeaderType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1NextExtensionHeaderType) Clone() (PatternFlowGtpv1NextExtensionHeaderType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1NextExtensionHeaderType()
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

func (obj *patternFlowGtpv1NextExtensionHeaderType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGtpv1NextExtensionHeaderType is next extension header. Exists if any of the e_flag, s_flag, or pn_flag bits are on.  Must be interpreted only if the e_flag bit is on.
type PatternFlowGtpv1NextExtensionHeaderType interface {
	Validation
	// msg marshals PatternFlowGtpv1NextExtensionHeaderType to protobuf object *otg.PatternFlowGtpv1NextExtensionHeaderType
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1NextExtensionHeaderType
	// setMsg unmarshals PatternFlowGtpv1NextExtensionHeaderType from protobuf object *otg.PatternFlowGtpv1NextExtensionHeaderType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1NextExtensionHeaderType) PatternFlowGtpv1NextExtensionHeaderType
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1NextExtensionHeaderType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1NextExtensionHeaderType
	// validate validates PatternFlowGtpv1NextExtensionHeaderType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1NextExtensionHeaderType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGtpv1NextExtensionHeaderTypeChoiceEnum, set in PatternFlowGtpv1NextExtensionHeaderType
	Choice() PatternFlowGtpv1NextExtensionHeaderTypeChoiceEnum
	// setChoice assigns PatternFlowGtpv1NextExtensionHeaderTypeChoiceEnum provided by user to PatternFlowGtpv1NextExtensionHeaderType
	setChoice(value PatternFlowGtpv1NextExtensionHeaderTypeChoiceEnum) PatternFlowGtpv1NextExtensionHeaderType
	// HasChoice checks if Choice has been set in PatternFlowGtpv1NextExtensionHeaderType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowGtpv1NextExtensionHeaderType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowGtpv1NextExtensionHeaderType
	SetValue(value uint32) PatternFlowGtpv1NextExtensionHeaderType
	// HasValue checks if Value has been set in PatternFlowGtpv1NextExtensionHeaderType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowGtpv1NextExtensionHeaderType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowGtpv1NextExtensionHeaderType
	SetValues(value []uint32) PatternFlowGtpv1NextExtensionHeaderType
	// Increment returns PatternFlowGtpv1NextExtensionHeaderTypeCounter, set in PatternFlowGtpv1NextExtensionHeaderType.
	// PatternFlowGtpv1NextExtensionHeaderTypeCounter is integer counter pattern
	Increment() PatternFlowGtpv1NextExtensionHeaderTypeCounter
	// SetIncrement assigns PatternFlowGtpv1NextExtensionHeaderTypeCounter provided by user to PatternFlowGtpv1NextExtensionHeaderType.
	// PatternFlowGtpv1NextExtensionHeaderTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowGtpv1NextExtensionHeaderTypeCounter) PatternFlowGtpv1NextExtensionHeaderType
	// HasIncrement checks if Increment has been set in PatternFlowGtpv1NextExtensionHeaderType
	HasIncrement() bool
	// Decrement returns PatternFlowGtpv1NextExtensionHeaderTypeCounter, set in PatternFlowGtpv1NextExtensionHeaderType.
	// PatternFlowGtpv1NextExtensionHeaderTypeCounter is integer counter pattern
	Decrement() PatternFlowGtpv1NextExtensionHeaderTypeCounter
	// SetDecrement assigns PatternFlowGtpv1NextExtensionHeaderTypeCounter provided by user to PatternFlowGtpv1NextExtensionHeaderType.
	// PatternFlowGtpv1NextExtensionHeaderTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowGtpv1NextExtensionHeaderTypeCounter) PatternFlowGtpv1NextExtensionHeaderType
	// HasDecrement checks if Decrement has been set in PatternFlowGtpv1NextExtensionHeaderType
	HasDecrement() bool
	// MetricTags returns PatternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIterIter, set in PatternFlowGtpv1NextExtensionHeaderType
	MetricTags() PatternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter
	setNil()
}

type PatternFlowGtpv1NextExtensionHeaderTypeChoiceEnum string

// Enum of Choice on PatternFlowGtpv1NextExtensionHeaderType
var PatternFlowGtpv1NextExtensionHeaderTypeChoice = struct {
	VALUE     PatternFlowGtpv1NextExtensionHeaderTypeChoiceEnum
	VALUES    PatternFlowGtpv1NextExtensionHeaderTypeChoiceEnum
	INCREMENT PatternFlowGtpv1NextExtensionHeaderTypeChoiceEnum
	DECREMENT PatternFlowGtpv1NextExtensionHeaderTypeChoiceEnum
}{
	VALUE:     PatternFlowGtpv1NextExtensionHeaderTypeChoiceEnum("value"),
	VALUES:    PatternFlowGtpv1NextExtensionHeaderTypeChoiceEnum("values"),
	INCREMENT: PatternFlowGtpv1NextExtensionHeaderTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowGtpv1NextExtensionHeaderTypeChoiceEnum("decrement"),
}

func (obj *patternFlowGtpv1NextExtensionHeaderType) Choice() PatternFlowGtpv1NextExtensionHeaderTypeChoiceEnum {
	return PatternFlowGtpv1NextExtensionHeaderTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGtpv1NextExtensionHeaderType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGtpv1NextExtensionHeaderType) setChoice(value PatternFlowGtpv1NextExtensionHeaderTypeChoiceEnum) PatternFlowGtpv1NextExtensionHeaderType {
	intValue, ok := otg.PatternFlowGtpv1NextExtensionHeaderType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGtpv1NextExtensionHeaderTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGtpv1NextExtensionHeaderType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGtpv1NextExtensionHeaderTypeChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGtpv1NextExtensionHeaderTypeChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGtpv1NextExtensionHeaderTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGtpv1NextExtensionHeaderTypeCounter().msg()
	}

	if value == PatternFlowGtpv1NextExtensionHeaderTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGtpv1NextExtensionHeaderTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv1NextExtensionHeaderType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGtpv1NextExtensionHeaderTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv1NextExtensionHeaderType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowGtpv1NextExtensionHeaderType object
func (obj *patternFlowGtpv1NextExtensionHeaderType) SetValue(value uint32) PatternFlowGtpv1NextExtensionHeaderType {
	obj.setChoice(PatternFlowGtpv1NextExtensionHeaderTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowGtpv1NextExtensionHeaderType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowGtpv1NextExtensionHeaderType object
func (obj *patternFlowGtpv1NextExtensionHeaderType) SetValues(value []uint32) PatternFlowGtpv1NextExtensionHeaderType {
	obj.setChoice(PatternFlowGtpv1NextExtensionHeaderTypeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGtpv1NextExtensionHeaderTypeCounter
func (obj *patternFlowGtpv1NextExtensionHeaderType) Increment() PatternFlowGtpv1NextExtensionHeaderTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGtpv1NextExtensionHeaderTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGtpv1NextExtensionHeaderTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGtpv1NextExtensionHeaderTypeCounter
func (obj *patternFlowGtpv1NextExtensionHeaderType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGtpv1NextExtensionHeaderTypeCounter value in the PatternFlowGtpv1NextExtensionHeaderType object
func (obj *patternFlowGtpv1NextExtensionHeaderType) SetIncrement(value PatternFlowGtpv1NextExtensionHeaderTypeCounter) PatternFlowGtpv1NextExtensionHeaderType {
	obj.setChoice(PatternFlowGtpv1NextExtensionHeaderTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGtpv1NextExtensionHeaderTypeCounter
func (obj *patternFlowGtpv1NextExtensionHeaderType) Decrement() PatternFlowGtpv1NextExtensionHeaderTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGtpv1NextExtensionHeaderTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGtpv1NextExtensionHeaderTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGtpv1NextExtensionHeaderTypeCounter
func (obj *patternFlowGtpv1NextExtensionHeaderType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGtpv1NextExtensionHeaderTypeCounter value in the PatternFlowGtpv1NextExtensionHeaderType object
func (obj *patternFlowGtpv1NextExtensionHeaderType) SetDecrement(value PatternFlowGtpv1NextExtensionHeaderTypeCounter) PatternFlowGtpv1NextExtensionHeaderType {
	obj.setChoice(PatternFlowGtpv1NextExtensionHeaderTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGtpv1NextExtensionHeaderTypeMetricTag
func (obj *patternFlowGtpv1NextExtensionHeaderType) MetricTags() PatternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGtpv1NextExtensionHeaderTypeMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter struct {
	obj                                                   *patternFlowGtpv1NextExtensionHeaderType
	patternFlowGtpv1NextExtensionHeaderTypeMetricTagSlice []PatternFlowGtpv1NextExtensionHeaderTypeMetricTag
	fieldPtr                                              *[]*otg.PatternFlowGtpv1NextExtensionHeaderTypeMetricTag
}

func newPatternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter(ptr *[]*otg.PatternFlowGtpv1NextExtensionHeaderTypeMetricTag) PatternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter {
	return &patternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter{fieldPtr: ptr}
}

type PatternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter interface {
	setMsg(*patternFlowGtpv1NextExtensionHeaderType) PatternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter
	Items() []PatternFlowGtpv1NextExtensionHeaderTypeMetricTag
	Add() PatternFlowGtpv1NextExtensionHeaderTypeMetricTag
	Append(items ...PatternFlowGtpv1NextExtensionHeaderTypeMetricTag) PatternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter
	Set(index int, newObj PatternFlowGtpv1NextExtensionHeaderTypeMetricTag) PatternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter
	Clear() PatternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter
	clearHolderSlice() PatternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter
	appendHolderSlice(item PatternFlowGtpv1NextExtensionHeaderTypeMetricTag) PatternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter) setMsg(msg *patternFlowGtpv1NextExtensionHeaderType) PatternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGtpv1NextExtensionHeaderTypeMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter) Items() []PatternFlowGtpv1NextExtensionHeaderTypeMetricTag {
	return obj.patternFlowGtpv1NextExtensionHeaderTypeMetricTagSlice
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter) Add() PatternFlowGtpv1NextExtensionHeaderTypeMetricTag {
	newObj := &otg.PatternFlowGtpv1NextExtensionHeaderTypeMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGtpv1NextExtensionHeaderTypeMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGtpv1NextExtensionHeaderTypeMetricTagSlice = append(obj.patternFlowGtpv1NextExtensionHeaderTypeMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter) Append(items ...PatternFlowGtpv1NextExtensionHeaderTypeMetricTag) PatternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGtpv1NextExtensionHeaderTypeMetricTagSlice = append(obj.patternFlowGtpv1NextExtensionHeaderTypeMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter) Set(index int, newObj PatternFlowGtpv1NextExtensionHeaderTypeMetricTag) PatternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGtpv1NextExtensionHeaderTypeMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter) Clear() PatternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGtpv1NextExtensionHeaderTypeMetricTag{}
		obj.patternFlowGtpv1NextExtensionHeaderTypeMetricTagSlice = []PatternFlowGtpv1NextExtensionHeaderTypeMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter) clearHolderSlice() PatternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter {
	if len(obj.patternFlowGtpv1NextExtensionHeaderTypeMetricTagSlice) > 0 {
		obj.patternFlowGtpv1NextExtensionHeaderTypeMetricTagSlice = []PatternFlowGtpv1NextExtensionHeaderTypeMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter) appendHolderSlice(item PatternFlowGtpv1NextExtensionHeaderTypeMetricTag) PatternFlowGtpv1NextExtensionHeaderTypePatternFlowGtpv1NextExtensionHeaderTypeMetricTagIter {
	obj.patternFlowGtpv1NextExtensionHeaderTypeMetricTagSlice = append(obj.patternFlowGtpv1NextExtensionHeaderTypeMetricTagSlice, item)
	return obj
}

func (obj *patternFlowGtpv1NextExtensionHeaderType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1NextExtensionHeaderType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowGtpv1NextExtensionHeaderType.Values <= 255 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowGtpv1NextExtensionHeaderTypeMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGtpv1NextExtensionHeaderType) setDefault() {
	var choices_set int = 0
	var choice PatternFlowGtpv1NextExtensionHeaderTypeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowGtpv1NextExtensionHeaderTypeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowGtpv1NextExtensionHeaderTypeChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowGtpv1NextExtensionHeaderTypeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowGtpv1NextExtensionHeaderTypeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowGtpv1NextExtensionHeaderTypeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowGtpv1NextExtensionHeaderType")
			}
		} else {
			intVal := otg.PatternFlowGtpv1NextExtensionHeaderType_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowGtpv1NextExtensionHeaderType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
