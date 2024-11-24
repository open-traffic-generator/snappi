package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4TosReliability *****
type patternFlowIpv4TosReliability struct {
	validation
	obj              *otg.PatternFlowIpv4TosReliability
	marshaller       marshalPatternFlowIpv4TosReliability
	unMarshaller     unMarshalPatternFlowIpv4TosReliability
	incrementHolder  PatternFlowIpv4TosReliabilityCounter
	decrementHolder  PatternFlowIpv4TosReliabilityCounter
	metricTagsHolder PatternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter
}

func NewPatternFlowIpv4TosReliability() PatternFlowIpv4TosReliability {
	obj := patternFlowIpv4TosReliability{obj: &otg.PatternFlowIpv4TosReliability{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4TosReliability) msg() *otg.PatternFlowIpv4TosReliability {
	return obj.obj
}

func (obj *patternFlowIpv4TosReliability) setMsg(msg *otg.PatternFlowIpv4TosReliability) PatternFlowIpv4TosReliability {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4TosReliability struct {
	obj *patternFlowIpv4TosReliability
}

type marshalPatternFlowIpv4TosReliability interface {
	// ToProto marshals PatternFlowIpv4TosReliability to protobuf object *otg.PatternFlowIpv4TosReliability
	ToProto() (*otg.PatternFlowIpv4TosReliability, error)
	// ToPbText marshals PatternFlowIpv4TosReliability to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4TosReliability to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4TosReliability to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4TosReliability struct {
	obj *patternFlowIpv4TosReliability
}

type unMarshalPatternFlowIpv4TosReliability interface {
	// FromProto unmarshals PatternFlowIpv4TosReliability from protobuf object *otg.PatternFlowIpv4TosReliability
	FromProto(msg *otg.PatternFlowIpv4TosReliability) (PatternFlowIpv4TosReliability, error)
	// FromPbText unmarshals PatternFlowIpv4TosReliability from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4TosReliability from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4TosReliability from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4TosReliability) Marshal() marshalPatternFlowIpv4TosReliability {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4TosReliability{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4TosReliability) Unmarshal() unMarshalPatternFlowIpv4TosReliability {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4TosReliability{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4TosReliability) ToProto() (*otg.PatternFlowIpv4TosReliability, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4TosReliability) FromProto(msg *otg.PatternFlowIpv4TosReliability) (PatternFlowIpv4TosReliability, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4TosReliability) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosReliability) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4TosReliability) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosReliability) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4TosReliability) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosReliability) FromJson(value string) error {
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

func (obj *patternFlowIpv4TosReliability) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosReliability) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosReliability) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4TosReliability) Clone() (PatternFlowIpv4TosReliability, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4TosReliability()
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

func (obj *patternFlowIpv4TosReliability) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4TosReliability is reliability
type PatternFlowIpv4TosReliability interface {
	Validation
	// msg marshals PatternFlowIpv4TosReliability to protobuf object *otg.PatternFlowIpv4TosReliability
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4TosReliability
	// setMsg unmarshals PatternFlowIpv4TosReliability from protobuf object *otg.PatternFlowIpv4TosReliability
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4TosReliability) PatternFlowIpv4TosReliability
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4TosReliability
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4TosReliability
	// validate validates PatternFlowIpv4TosReliability
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4TosReliability, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4TosReliabilityChoiceEnum, set in PatternFlowIpv4TosReliability
	Choice() PatternFlowIpv4TosReliabilityChoiceEnum
	// setChoice assigns PatternFlowIpv4TosReliabilityChoiceEnum provided by user to PatternFlowIpv4TosReliability
	setChoice(value PatternFlowIpv4TosReliabilityChoiceEnum) PatternFlowIpv4TosReliability
	// HasChoice checks if Choice has been set in PatternFlowIpv4TosReliability
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv4TosReliability.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv4TosReliability
	SetValue(value uint32) PatternFlowIpv4TosReliability
	// HasValue checks if Value has been set in PatternFlowIpv4TosReliability
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv4TosReliability.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv4TosReliability
	SetValues(value []uint32) PatternFlowIpv4TosReliability
	// Increment returns PatternFlowIpv4TosReliabilityCounter, set in PatternFlowIpv4TosReliability.
	// PatternFlowIpv4TosReliabilityCounter is integer counter pattern
	Increment() PatternFlowIpv4TosReliabilityCounter
	// SetIncrement assigns PatternFlowIpv4TosReliabilityCounter provided by user to PatternFlowIpv4TosReliability.
	// PatternFlowIpv4TosReliabilityCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv4TosReliabilityCounter) PatternFlowIpv4TosReliability
	// HasIncrement checks if Increment has been set in PatternFlowIpv4TosReliability
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4TosReliabilityCounter, set in PatternFlowIpv4TosReliability.
	// PatternFlowIpv4TosReliabilityCounter is integer counter pattern
	Decrement() PatternFlowIpv4TosReliabilityCounter
	// SetDecrement assigns PatternFlowIpv4TosReliabilityCounter provided by user to PatternFlowIpv4TosReliability.
	// PatternFlowIpv4TosReliabilityCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv4TosReliabilityCounter) PatternFlowIpv4TosReliability
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4TosReliability
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIterIter, set in PatternFlowIpv4TosReliability
	MetricTags() PatternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter
	setNil()
}

type PatternFlowIpv4TosReliabilityChoiceEnum string

// Enum of Choice on PatternFlowIpv4TosReliability
var PatternFlowIpv4TosReliabilityChoice = struct {
	VALUE     PatternFlowIpv4TosReliabilityChoiceEnum
	VALUES    PatternFlowIpv4TosReliabilityChoiceEnum
	INCREMENT PatternFlowIpv4TosReliabilityChoiceEnum
	DECREMENT PatternFlowIpv4TosReliabilityChoiceEnum
}{
	VALUE:     PatternFlowIpv4TosReliabilityChoiceEnum("value"),
	VALUES:    PatternFlowIpv4TosReliabilityChoiceEnum("values"),
	INCREMENT: PatternFlowIpv4TosReliabilityChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4TosReliabilityChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4TosReliability) Choice() PatternFlowIpv4TosReliabilityChoiceEnum {
	return PatternFlowIpv4TosReliabilityChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4TosReliability) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4TosReliability) setChoice(value PatternFlowIpv4TosReliabilityChoiceEnum) PatternFlowIpv4TosReliability {
	intValue, ok := otg.PatternFlowIpv4TosReliability_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4TosReliabilityChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4TosReliability_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4TosReliabilityChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4TosReliabilityChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4TosReliabilityChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4TosReliabilityCounter().msg()
	}

	if value == PatternFlowIpv4TosReliabilityChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4TosReliabilityCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4TosReliability) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4TosReliabilityChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4TosReliability) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv4TosReliability object
func (obj *patternFlowIpv4TosReliability) SetValue(value uint32) PatternFlowIpv4TosReliability {
	obj.setChoice(PatternFlowIpv4TosReliabilityChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv4TosReliability) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv4TosReliability object
func (obj *patternFlowIpv4TosReliability) SetValues(value []uint32) PatternFlowIpv4TosReliability {
	obj.setChoice(PatternFlowIpv4TosReliabilityChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv4TosReliabilityCounter
func (obj *patternFlowIpv4TosReliability) Increment() PatternFlowIpv4TosReliabilityCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4TosReliabilityChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4TosReliabilityCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4TosReliabilityCounter
func (obj *patternFlowIpv4TosReliability) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4TosReliabilityCounter value in the PatternFlowIpv4TosReliability object
func (obj *patternFlowIpv4TosReliability) SetIncrement(value PatternFlowIpv4TosReliabilityCounter) PatternFlowIpv4TosReliability {
	obj.setChoice(PatternFlowIpv4TosReliabilityChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4TosReliabilityCounter
func (obj *patternFlowIpv4TosReliability) Decrement() PatternFlowIpv4TosReliabilityCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4TosReliabilityChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4TosReliabilityCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4TosReliabilityCounter
func (obj *patternFlowIpv4TosReliability) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4TosReliabilityCounter value in the PatternFlowIpv4TosReliability object
func (obj *patternFlowIpv4TosReliability) SetDecrement(value PatternFlowIpv4TosReliabilityCounter) PatternFlowIpv4TosReliability {
	obj.setChoice(PatternFlowIpv4TosReliabilityChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv4TosReliabilityMetricTag
func (obj *patternFlowIpv4TosReliability) MetricTags() PatternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv4TosReliabilityMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter struct {
	obj                                         *patternFlowIpv4TosReliability
	patternFlowIpv4TosReliabilityMetricTagSlice []PatternFlowIpv4TosReliabilityMetricTag
	fieldPtr                                    *[]*otg.PatternFlowIpv4TosReliabilityMetricTag
}

func newPatternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter(ptr *[]*otg.PatternFlowIpv4TosReliabilityMetricTag) PatternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter {
	return &patternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter interface {
	setMsg(*patternFlowIpv4TosReliability) PatternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter
	Items() []PatternFlowIpv4TosReliabilityMetricTag
	Add() PatternFlowIpv4TosReliabilityMetricTag
	Append(items ...PatternFlowIpv4TosReliabilityMetricTag) PatternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter
	Set(index int, newObj PatternFlowIpv4TosReliabilityMetricTag) PatternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter
	Clear() PatternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter
	clearHolderSlice() PatternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter
	appendHolderSlice(item PatternFlowIpv4TosReliabilityMetricTag) PatternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter
}

func (obj *patternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter) setMsg(msg *patternFlowIpv4TosReliability) PatternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv4TosReliabilityMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter) Items() []PatternFlowIpv4TosReliabilityMetricTag {
	return obj.patternFlowIpv4TosReliabilityMetricTagSlice
}

func (obj *patternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter) Add() PatternFlowIpv4TosReliabilityMetricTag {
	newObj := &otg.PatternFlowIpv4TosReliabilityMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv4TosReliabilityMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv4TosReliabilityMetricTagSlice = append(obj.patternFlowIpv4TosReliabilityMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter) Append(items ...PatternFlowIpv4TosReliabilityMetricTag) PatternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv4TosReliabilityMetricTagSlice = append(obj.patternFlowIpv4TosReliabilityMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter) Set(index int, newObj PatternFlowIpv4TosReliabilityMetricTag) PatternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv4TosReliabilityMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter) Clear() PatternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv4TosReliabilityMetricTag{}
		obj.patternFlowIpv4TosReliabilityMetricTagSlice = []PatternFlowIpv4TosReliabilityMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter) clearHolderSlice() PatternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter {
	if len(obj.patternFlowIpv4TosReliabilityMetricTagSlice) > 0 {
		obj.patternFlowIpv4TosReliabilityMetricTagSlice = []PatternFlowIpv4TosReliabilityMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter) appendHolderSlice(item PatternFlowIpv4TosReliabilityMetricTag) PatternFlowIpv4TosReliabilityPatternFlowIpv4TosReliabilityMetricTagIter {
	obj.patternFlowIpv4TosReliabilityMetricTagSlice = append(obj.patternFlowIpv4TosReliabilityMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIpv4TosReliability) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosReliability.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv4TosReliability.Values <= 1 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv4TosReliabilityMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIpv4TosReliability) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv4TosReliabilityChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv4TosReliabilityChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv4TosReliabilityChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv4TosReliabilityChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv4TosReliabilityChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv4TosReliabilityChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv4TosReliability")
			}
		} else {
			intVal := otg.PatternFlowIpv4TosReliability_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv4TosReliability_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
