package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv2PiggybackingFlag *****
type patternFlowGtpv2PiggybackingFlag struct {
	validation
	obj              *otg.PatternFlowGtpv2PiggybackingFlag
	marshaller       marshalPatternFlowGtpv2PiggybackingFlag
	unMarshaller     unMarshalPatternFlowGtpv2PiggybackingFlag
	incrementHolder  PatternFlowGtpv2PiggybackingFlagCounter
	decrementHolder  PatternFlowGtpv2PiggybackingFlagCounter
	metricTagsHolder PatternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter
}

func NewPatternFlowGtpv2PiggybackingFlag() PatternFlowGtpv2PiggybackingFlag {
	obj := patternFlowGtpv2PiggybackingFlag{obj: &otg.PatternFlowGtpv2PiggybackingFlag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv2PiggybackingFlag) msg() *otg.PatternFlowGtpv2PiggybackingFlag {
	return obj.obj
}

func (obj *patternFlowGtpv2PiggybackingFlag) setMsg(msg *otg.PatternFlowGtpv2PiggybackingFlag) PatternFlowGtpv2PiggybackingFlag {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv2PiggybackingFlag struct {
	obj *patternFlowGtpv2PiggybackingFlag
}

type marshalPatternFlowGtpv2PiggybackingFlag interface {
	// ToProto marshals PatternFlowGtpv2PiggybackingFlag to protobuf object *otg.PatternFlowGtpv2PiggybackingFlag
	ToProto() (*otg.PatternFlowGtpv2PiggybackingFlag, error)
	// ToPbText marshals PatternFlowGtpv2PiggybackingFlag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv2PiggybackingFlag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv2PiggybackingFlag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowGtpv2PiggybackingFlag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowGtpv2PiggybackingFlag struct {
	obj *patternFlowGtpv2PiggybackingFlag
}

type unMarshalPatternFlowGtpv2PiggybackingFlag interface {
	// FromProto unmarshals PatternFlowGtpv2PiggybackingFlag from protobuf object *otg.PatternFlowGtpv2PiggybackingFlag
	FromProto(msg *otg.PatternFlowGtpv2PiggybackingFlag) (PatternFlowGtpv2PiggybackingFlag, error)
	// FromPbText unmarshals PatternFlowGtpv2PiggybackingFlag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv2PiggybackingFlag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv2PiggybackingFlag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv2PiggybackingFlag) Marshal() marshalPatternFlowGtpv2PiggybackingFlag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv2PiggybackingFlag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv2PiggybackingFlag) Unmarshal() unMarshalPatternFlowGtpv2PiggybackingFlag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv2PiggybackingFlag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv2PiggybackingFlag) ToProto() (*otg.PatternFlowGtpv2PiggybackingFlag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv2PiggybackingFlag) FromProto(msg *otg.PatternFlowGtpv2PiggybackingFlag) (PatternFlowGtpv2PiggybackingFlag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv2PiggybackingFlag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2PiggybackingFlag) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv2PiggybackingFlag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2PiggybackingFlag) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv2PiggybackingFlag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowGtpv2PiggybackingFlag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2PiggybackingFlag) FromJson(value string) error {
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

func (obj *patternFlowGtpv2PiggybackingFlag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2PiggybackingFlag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2PiggybackingFlag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv2PiggybackingFlag) Clone() (PatternFlowGtpv2PiggybackingFlag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv2PiggybackingFlag()
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

func (obj *patternFlowGtpv2PiggybackingFlag) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGtpv2PiggybackingFlag is if piggybacking_flag is set to 1 then another GTP-C message with its own header shall be present at the end of the current message
type PatternFlowGtpv2PiggybackingFlag interface {
	Validation
	// msg marshals PatternFlowGtpv2PiggybackingFlag to protobuf object *otg.PatternFlowGtpv2PiggybackingFlag
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv2PiggybackingFlag
	// setMsg unmarshals PatternFlowGtpv2PiggybackingFlag from protobuf object *otg.PatternFlowGtpv2PiggybackingFlag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv2PiggybackingFlag) PatternFlowGtpv2PiggybackingFlag
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv2PiggybackingFlag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv2PiggybackingFlag
	// validate validates PatternFlowGtpv2PiggybackingFlag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv2PiggybackingFlag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGtpv2PiggybackingFlagChoiceEnum, set in PatternFlowGtpv2PiggybackingFlag
	Choice() PatternFlowGtpv2PiggybackingFlagChoiceEnum
	// setChoice assigns PatternFlowGtpv2PiggybackingFlagChoiceEnum provided by user to PatternFlowGtpv2PiggybackingFlag
	setChoice(value PatternFlowGtpv2PiggybackingFlagChoiceEnum) PatternFlowGtpv2PiggybackingFlag
	// HasChoice checks if Choice has been set in PatternFlowGtpv2PiggybackingFlag
	HasChoice() bool
	// Value returns uint32, set in PatternFlowGtpv2PiggybackingFlag.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowGtpv2PiggybackingFlag
	SetValue(value uint32) PatternFlowGtpv2PiggybackingFlag
	// HasValue checks if Value has been set in PatternFlowGtpv2PiggybackingFlag
	HasValue() bool
	// Values returns []uint32, set in PatternFlowGtpv2PiggybackingFlag.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowGtpv2PiggybackingFlag
	SetValues(value []uint32) PatternFlowGtpv2PiggybackingFlag
	// Increment returns PatternFlowGtpv2PiggybackingFlagCounter, set in PatternFlowGtpv2PiggybackingFlag.
	// PatternFlowGtpv2PiggybackingFlagCounter is integer counter pattern
	Increment() PatternFlowGtpv2PiggybackingFlagCounter
	// SetIncrement assigns PatternFlowGtpv2PiggybackingFlagCounter provided by user to PatternFlowGtpv2PiggybackingFlag.
	// PatternFlowGtpv2PiggybackingFlagCounter is integer counter pattern
	SetIncrement(value PatternFlowGtpv2PiggybackingFlagCounter) PatternFlowGtpv2PiggybackingFlag
	// HasIncrement checks if Increment has been set in PatternFlowGtpv2PiggybackingFlag
	HasIncrement() bool
	// Decrement returns PatternFlowGtpv2PiggybackingFlagCounter, set in PatternFlowGtpv2PiggybackingFlag.
	// PatternFlowGtpv2PiggybackingFlagCounter is integer counter pattern
	Decrement() PatternFlowGtpv2PiggybackingFlagCounter
	// SetDecrement assigns PatternFlowGtpv2PiggybackingFlagCounter provided by user to PatternFlowGtpv2PiggybackingFlag.
	// PatternFlowGtpv2PiggybackingFlagCounter is integer counter pattern
	SetDecrement(value PatternFlowGtpv2PiggybackingFlagCounter) PatternFlowGtpv2PiggybackingFlag
	// HasDecrement checks if Decrement has been set in PatternFlowGtpv2PiggybackingFlag
	HasDecrement() bool
	// MetricTags returns PatternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIterIter, set in PatternFlowGtpv2PiggybackingFlag
	MetricTags() PatternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter
	setNil()
}

type PatternFlowGtpv2PiggybackingFlagChoiceEnum string

// Enum of Choice on PatternFlowGtpv2PiggybackingFlag
var PatternFlowGtpv2PiggybackingFlagChoice = struct {
	VALUE     PatternFlowGtpv2PiggybackingFlagChoiceEnum
	VALUES    PatternFlowGtpv2PiggybackingFlagChoiceEnum
	INCREMENT PatternFlowGtpv2PiggybackingFlagChoiceEnum
	DECREMENT PatternFlowGtpv2PiggybackingFlagChoiceEnum
}{
	VALUE:     PatternFlowGtpv2PiggybackingFlagChoiceEnum("value"),
	VALUES:    PatternFlowGtpv2PiggybackingFlagChoiceEnum("values"),
	INCREMENT: PatternFlowGtpv2PiggybackingFlagChoiceEnum("increment"),
	DECREMENT: PatternFlowGtpv2PiggybackingFlagChoiceEnum("decrement"),
}

func (obj *patternFlowGtpv2PiggybackingFlag) Choice() PatternFlowGtpv2PiggybackingFlagChoiceEnum {
	return PatternFlowGtpv2PiggybackingFlagChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGtpv2PiggybackingFlag) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGtpv2PiggybackingFlag) setChoice(value PatternFlowGtpv2PiggybackingFlagChoiceEnum) PatternFlowGtpv2PiggybackingFlag {
	intValue, ok := otg.PatternFlowGtpv2PiggybackingFlag_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGtpv2PiggybackingFlagChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGtpv2PiggybackingFlag_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGtpv2PiggybackingFlagChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGtpv2PiggybackingFlagChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGtpv2PiggybackingFlagChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGtpv2PiggybackingFlagCounter().msg()
	}

	if value == PatternFlowGtpv2PiggybackingFlagChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGtpv2PiggybackingFlagCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv2PiggybackingFlag) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGtpv2PiggybackingFlagChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv2PiggybackingFlag) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowGtpv2PiggybackingFlag object
func (obj *patternFlowGtpv2PiggybackingFlag) SetValue(value uint32) PatternFlowGtpv2PiggybackingFlag {
	obj.setChoice(PatternFlowGtpv2PiggybackingFlagChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowGtpv2PiggybackingFlag) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowGtpv2PiggybackingFlag object
func (obj *patternFlowGtpv2PiggybackingFlag) SetValues(value []uint32) PatternFlowGtpv2PiggybackingFlag {
	obj.setChoice(PatternFlowGtpv2PiggybackingFlagChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGtpv2PiggybackingFlagCounter
func (obj *patternFlowGtpv2PiggybackingFlag) Increment() PatternFlowGtpv2PiggybackingFlagCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGtpv2PiggybackingFlagChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGtpv2PiggybackingFlagCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGtpv2PiggybackingFlagCounter
func (obj *patternFlowGtpv2PiggybackingFlag) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGtpv2PiggybackingFlagCounter value in the PatternFlowGtpv2PiggybackingFlag object
func (obj *patternFlowGtpv2PiggybackingFlag) SetIncrement(value PatternFlowGtpv2PiggybackingFlagCounter) PatternFlowGtpv2PiggybackingFlag {
	obj.setChoice(PatternFlowGtpv2PiggybackingFlagChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGtpv2PiggybackingFlagCounter
func (obj *patternFlowGtpv2PiggybackingFlag) Decrement() PatternFlowGtpv2PiggybackingFlagCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGtpv2PiggybackingFlagChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGtpv2PiggybackingFlagCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGtpv2PiggybackingFlagCounter
func (obj *patternFlowGtpv2PiggybackingFlag) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGtpv2PiggybackingFlagCounter value in the PatternFlowGtpv2PiggybackingFlag object
func (obj *patternFlowGtpv2PiggybackingFlag) SetDecrement(value PatternFlowGtpv2PiggybackingFlagCounter) PatternFlowGtpv2PiggybackingFlag {
	obj.setChoice(PatternFlowGtpv2PiggybackingFlagChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGtpv2PiggybackingFlagMetricTag
func (obj *patternFlowGtpv2PiggybackingFlag) MetricTags() PatternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGtpv2PiggybackingFlagMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter struct {
	obj                                            *patternFlowGtpv2PiggybackingFlag
	patternFlowGtpv2PiggybackingFlagMetricTagSlice []PatternFlowGtpv2PiggybackingFlagMetricTag
	fieldPtr                                       *[]*otg.PatternFlowGtpv2PiggybackingFlagMetricTag
}

func newPatternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter(ptr *[]*otg.PatternFlowGtpv2PiggybackingFlagMetricTag) PatternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter {
	return &patternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter{fieldPtr: ptr}
}

type PatternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter interface {
	setMsg(*patternFlowGtpv2PiggybackingFlag) PatternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter
	Items() []PatternFlowGtpv2PiggybackingFlagMetricTag
	Add() PatternFlowGtpv2PiggybackingFlagMetricTag
	Append(items ...PatternFlowGtpv2PiggybackingFlagMetricTag) PatternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter
	Set(index int, newObj PatternFlowGtpv2PiggybackingFlagMetricTag) PatternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter
	Clear() PatternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter
	clearHolderSlice() PatternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter
	appendHolderSlice(item PatternFlowGtpv2PiggybackingFlagMetricTag) PatternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter
}

func (obj *patternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter) setMsg(msg *patternFlowGtpv2PiggybackingFlag) PatternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGtpv2PiggybackingFlagMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter) Items() []PatternFlowGtpv2PiggybackingFlagMetricTag {
	return obj.patternFlowGtpv2PiggybackingFlagMetricTagSlice
}

func (obj *patternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter) Add() PatternFlowGtpv2PiggybackingFlagMetricTag {
	newObj := &otg.PatternFlowGtpv2PiggybackingFlagMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGtpv2PiggybackingFlagMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGtpv2PiggybackingFlagMetricTagSlice = append(obj.patternFlowGtpv2PiggybackingFlagMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter) Append(items ...PatternFlowGtpv2PiggybackingFlagMetricTag) PatternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGtpv2PiggybackingFlagMetricTagSlice = append(obj.patternFlowGtpv2PiggybackingFlagMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter) Set(index int, newObj PatternFlowGtpv2PiggybackingFlagMetricTag) PatternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGtpv2PiggybackingFlagMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter) Clear() PatternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGtpv2PiggybackingFlagMetricTag{}
		obj.patternFlowGtpv2PiggybackingFlagMetricTagSlice = []PatternFlowGtpv2PiggybackingFlagMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter) clearHolderSlice() PatternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter {
	if len(obj.patternFlowGtpv2PiggybackingFlagMetricTagSlice) > 0 {
		obj.patternFlowGtpv2PiggybackingFlagMetricTagSlice = []PatternFlowGtpv2PiggybackingFlagMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter) appendHolderSlice(item PatternFlowGtpv2PiggybackingFlagMetricTag) PatternFlowGtpv2PiggybackingFlagPatternFlowGtpv2PiggybackingFlagMetricTagIter {
	obj.patternFlowGtpv2PiggybackingFlagMetricTagSlice = append(obj.patternFlowGtpv2PiggybackingFlagMetricTagSlice, item)
	return obj
}

func (obj *patternFlowGtpv2PiggybackingFlag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2PiggybackingFlag.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowGtpv2PiggybackingFlag.Values <= 1 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowGtpv2PiggybackingFlagMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGtpv2PiggybackingFlag) setDefault() {
	var choices_set int = 0
	var choice PatternFlowGtpv2PiggybackingFlagChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowGtpv2PiggybackingFlagChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowGtpv2PiggybackingFlagChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowGtpv2PiggybackingFlagChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowGtpv2PiggybackingFlagChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowGtpv2PiggybackingFlagChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowGtpv2PiggybackingFlag")
			}
		} else {
			intVal := otg.PatternFlowGtpv2PiggybackingFlag_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowGtpv2PiggybackingFlag_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
