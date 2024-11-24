package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1PnFlag *****
type patternFlowGtpv1PnFlag struct {
	validation
	obj              *otg.PatternFlowGtpv1PnFlag
	marshaller       marshalPatternFlowGtpv1PnFlag
	unMarshaller     unMarshalPatternFlowGtpv1PnFlag
	incrementHolder  PatternFlowGtpv1PnFlagCounter
	decrementHolder  PatternFlowGtpv1PnFlagCounter
	metricTagsHolder PatternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter
}

func NewPatternFlowGtpv1PnFlag() PatternFlowGtpv1PnFlag {
	obj := patternFlowGtpv1PnFlag{obj: &otg.PatternFlowGtpv1PnFlag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1PnFlag) msg() *otg.PatternFlowGtpv1PnFlag {
	return obj.obj
}

func (obj *patternFlowGtpv1PnFlag) setMsg(msg *otg.PatternFlowGtpv1PnFlag) PatternFlowGtpv1PnFlag {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1PnFlag struct {
	obj *patternFlowGtpv1PnFlag
}

type marshalPatternFlowGtpv1PnFlag interface {
	// ToProto marshals PatternFlowGtpv1PnFlag to protobuf object *otg.PatternFlowGtpv1PnFlag
	ToProto() (*otg.PatternFlowGtpv1PnFlag, error)
	// ToPbText marshals PatternFlowGtpv1PnFlag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1PnFlag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1PnFlag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowGtpv1PnFlag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowGtpv1PnFlag struct {
	obj *patternFlowGtpv1PnFlag
}

type unMarshalPatternFlowGtpv1PnFlag interface {
	// FromProto unmarshals PatternFlowGtpv1PnFlag from protobuf object *otg.PatternFlowGtpv1PnFlag
	FromProto(msg *otg.PatternFlowGtpv1PnFlag) (PatternFlowGtpv1PnFlag, error)
	// FromPbText unmarshals PatternFlowGtpv1PnFlag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1PnFlag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1PnFlag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1PnFlag) Marshal() marshalPatternFlowGtpv1PnFlag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1PnFlag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1PnFlag) Unmarshal() unMarshalPatternFlowGtpv1PnFlag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1PnFlag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1PnFlag) ToProto() (*otg.PatternFlowGtpv1PnFlag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1PnFlag) FromProto(msg *otg.PatternFlowGtpv1PnFlag) (PatternFlowGtpv1PnFlag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1PnFlag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1PnFlag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1PnFlag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1PnFlag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1PnFlag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowGtpv1PnFlag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1PnFlag) FromJson(value string) error {
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

func (obj *patternFlowGtpv1PnFlag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1PnFlag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1PnFlag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1PnFlag) Clone() (PatternFlowGtpv1PnFlag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1PnFlag()
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

func (obj *patternFlowGtpv1PnFlag) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGtpv1PnFlag is n-PDU field present
type PatternFlowGtpv1PnFlag interface {
	Validation
	// msg marshals PatternFlowGtpv1PnFlag to protobuf object *otg.PatternFlowGtpv1PnFlag
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1PnFlag
	// setMsg unmarshals PatternFlowGtpv1PnFlag from protobuf object *otg.PatternFlowGtpv1PnFlag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1PnFlag) PatternFlowGtpv1PnFlag
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1PnFlag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1PnFlag
	// validate validates PatternFlowGtpv1PnFlag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1PnFlag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGtpv1PnFlagChoiceEnum, set in PatternFlowGtpv1PnFlag
	Choice() PatternFlowGtpv1PnFlagChoiceEnum
	// setChoice assigns PatternFlowGtpv1PnFlagChoiceEnum provided by user to PatternFlowGtpv1PnFlag
	setChoice(value PatternFlowGtpv1PnFlagChoiceEnum) PatternFlowGtpv1PnFlag
	// HasChoice checks if Choice has been set in PatternFlowGtpv1PnFlag
	HasChoice() bool
	// Value returns uint32, set in PatternFlowGtpv1PnFlag.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowGtpv1PnFlag
	SetValue(value uint32) PatternFlowGtpv1PnFlag
	// HasValue checks if Value has been set in PatternFlowGtpv1PnFlag
	HasValue() bool
	// Values returns []uint32, set in PatternFlowGtpv1PnFlag.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowGtpv1PnFlag
	SetValues(value []uint32) PatternFlowGtpv1PnFlag
	// Increment returns PatternFlowGtpv1PnFlagCounter, set in PatternFlowGtpv1PnFlag.
	// PatternFlowGtpv1PnFlagCounter is integer counter pattern
	Increment() PatternFlowGtpv1PnFlagCounter
	// SetIncrement assigns PatternFlowGtpv1PnFlagCounter provided by user to PatternFlowGtpv1PnFlag.
	// PatternFlowGtpv1PnFlagCounter is integer counter pattern
	SetIncrement(value PatternFlowGtpv1PnFlagCounter) PatternFlowGtpv1PnFlag
	// HasIncrement checks if Increment has been set in PatternFlowGtpv1PnFlag
	HasIncrement() bool
	// Decrement returns PatternFlowGtpv1PnFlagCounter, set in PatternFlowGtpv1PnFlag.
	// PatternFlowGtpv1PnFlagCounter is integer counter pattern
	Decrement() PatternFlowGtpv1PnFlagCounter
	// SetDecrement assigns PatternFlowGtpv1PnFlagCounter provided by user to PatternFlowGtpv1PnFlag.
	// PatternFlowGtpv1PnFlagCounter is integer counter pattern
	SetDecrement(value PatternFlowGtpv1PnFlagCounter) PatternFlowGtpv1PnFlag
	// HasDecrement checks if Decrement has been set in PatternFlowGtpv1PnFlag
	HasDecrement() bool
	// MetricTags returns PatternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIterIter, set in PatternFlowGtpv1PnFlag
	MetricTags() PatternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter
	setNil()
}

type PatternFlowGtpv1PnFlagChoiceEnum string

// Enum of Choice on PatternFlowGtpv1PnFlag
var PatternFlowGtpv1PnFlagChoice = struct {
	VALUE     PatternFlowGtpv1PnFlagChoiceEnum
	VALUES    PatternFlowGtpv1PnFlagChoiceEnum
	INCREMENT PatternFlowGtpv1PnFlagChoiceEnum
	DECREMENT PatternFlowGtpv1PnFlagChoiceEnum
}{
	VALUE:     PatternFlowGtpv1PnFlagChoiceEnum("value"),
	VALUES:    PatternFlowGtpv1PnFlagChoiceEnum("values"),
	INCREMENT: PatternFlowGtpv1PnFlagChoiceEnum("increment"),
	DECREMENT: PatternFlowGtpv1PnFlagChoiceEnum("decrement"),
}

func (obj *patternFlowGtpv1PnFlag) Choice() PatternFlowGtpv1PnFlagChoiceEnum {
	return PatternFlowGtpv1PnFlagChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGtpv1PnFlag) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGtpv1PnFlag) setChoice(value PatternFlowGtpv1PnFlagChoiceEnum) PatternFlowGtpv1PnFlag {
	intValue, ok := otg.PatternFlowGtpv1PnFlag_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGtpv1PnFlagChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGtpv1PnFlag_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGtpv1PnFlagChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGtpv1PnFlagChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGtpv1PnFlagChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGtpv1PnFlagCounter().msg()
	}

	if value == PatternFlowGtpv1PnFlagChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGtpv1PnFlagCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv1PnFlag) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGtpv1PnFlagChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv1PnFlag) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowGtpv1PnFlag object
func (obj *patternFlowGtpv1PnFlag) SetValue(value uint32) PatternFlowGtpv1PnFlag {
	obj.setChoice(PatternFlowGtpv1PnFlagChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowGtpv1PnFlag) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowGtpv1PnFlag object
func (obj *patternFlowGtpv1PnFlag) SetValues(value []uint32) PatternFlowGtpv1PnFlag {
	obj.setChoice(PatternFlowGtpv1PnFlagChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGtpv1PnFlagCounter
func (obj *patternFlowGtpv1PnFlag) Increment() PatternFlowGtpv1PnFlagCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGtpv1PnFlagChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGtpv1PnFlagCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGtpv1PnFlagCounter
func (obj *patternFlowGtpv1PnFlag) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGtpv1PnFlagCounter value in the PatternFlowGtpv1PnFlag object
func (obj *patternFlowGtpv1PnFlag) SetIncrement(value PatternFlowGtpv1PnFlagCounter) PatternFlowGtpv1PnFlag {
	obj.setChoice(PatternFlowGtpv1PnFlagChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGtpv1PnFlagCounter
func (obj *patternFlowGtpv1PnFlag) Decrement() PatternFlowGtpv1PnFlagCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGtpv1PnFlagChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGtpv1PnFlagCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGtpv1PnFlagCounter
func (obj *patternFlowGtpv1PnFlag) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGtpv1PnFlagCounter value in the PatternFlowGtpv1PnFlag object
func (obj *patternFlowGtpv1PnFlag) SetDecrement(value PatternFlowGtpv1PnFlagCounter) PatternFlowGtpv1PnFlag {
	obj.setChoice(PatternFlowGtpv1PnFlagChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGtpv1PnFlagMetricTag
func (obj *patternFlowGtpv1PnFlag) MetricTags() PatternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGtpv1PnFlagMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter struct {
	obj                                  *patternFlowGtpv1PnFlag
	patternFlowGtpv1PnFlagMetricTagSlice []PatternFlowGtpv1PnFlagMetricTag
	fieldPtr                             *[]*otg.PatternFlowGtpv1PnFlagMetricTag
}

func newPatternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter(ptr *[]*otg.PatternFlowGtpv1PnFlagMetricTag) PatternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter {
	return &patternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter{fieldPtr: ptr}
}

type PatternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter interface {
	setMsg(*patternFlowGtpv1PnFlag) PatternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter
	Items() []PatternFlowGtpv1PnFlagMetricTag
	Add() PatternFlowGtpv1PnFlagMetricTag
	Append(items ...PatternFlowGtpv1PnFlagMetricTag) PatternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter
	Set(index int, newObj PatternFlowGtpv1PnFlagMetricTag) PatternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter
	Clear() PatternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter
	clearHolderSlice() PatternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter
	appendHolderSlice(item PatternFlowGtpv1PnFlagMetricTag) PatternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter
}

func (obj *patternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter) setMsg(msg *patternFlowGtpv1PnFlag) PatternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGtpv1PnFlagMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter) Items() []PatternFlowGtpv1PnFlagMetricTag {
	return obj.patternFlowGtpv1PnFlagMetricTagSlice
}

func (obj *patternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter) Add() PatternFlowGtpv1PnFlagMetricTag {
	newObj := &otg.PatternFlowGtpv1PnFlagMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGtpv1PnFlagMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGtpv1PnFlagMetricTagSlice = append(obj.patternFlowGtpv1PnFlagMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter) Append(items ...PatternFlowGtpv1PnFlagMetricTag) PatternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGtpv1PnFlagMetricTagSlice = append(obj.patternFlowGtpv1PnFlagMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter) Set(index int, newObj PatternFlowGtpv1PnFlagMetricTag) PatternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGtpv1PnFlagMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter) Clear() PatternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGtpv1PnFlagMetricTag{}
		obj.patternFlowGtpv1PnFlagMetricTagSlice = []PatternFlowGtpv1PnFlagMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter) clearHolderSlice() PatternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter {
	if len(obj.patternFlowGtpv1PnFlagMetricTagSlice) > 0 {
		obj.patternFlowGtpv1PnFlagMetricTagSlice = []PatternFlowGtpv1PnFlagMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter) appendHolderSlice(item PatternFlowGtpv1PnFlagMetricTag) PatternFlowGtpv1PnFlagPatternFlowGtpv1PnFlagMetricTagIter {
	obj.patternFlowGtpv1PnFlagMetricTagSlice = append(obj.patternFlowGtpv1PnFlagMetricTagSlice, item)
	return obj
}

func (obj *patternFlowGtpv1PnFlag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1PnFlag.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowGtpv1PnFlag.Values <= 1 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowGtpv1PnFlagMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGtpv1PnFlag) setDefault() {
	var choices_set int = 0
	var choice PatternFlowGtpv1PnFlagChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowGtpv1PnFlagChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowGtpv1PnFlagChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowGtpv1PnFlagChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowGtpv1PnFlagChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowGtpv1PnFlagChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowGtpv1PnFlag")
			}
		} else {
			intVal := otg.PatternFlowGtpv1PnFlag_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowGtpv1PnFlag_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
