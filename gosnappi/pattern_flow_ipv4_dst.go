package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4Dst *****
type patternFlowIpv4Dst struct {
	validation
	obj              *otg.PatternFlowIpv4Dst
	marshaller       marshalPatternFlowIpv4Dst
	unMarshaller     unMarshalPatternFlowIpv4Dst
	incrementHolder  PatternFlowIpv4DstCounter
	decrementHolder  PatternFlowIpv4DstCounter
	metricTagsHolder PatternFlowIpv4DstPatternFlowIpv4DstMetricTagIter
	autoHolder       FlowIpv4Auto
	randomHolder     PatternFlowIpv4DstRandom
}

func NewPatternFlowIpv4Dst() PatternFlowIpv4Dst {
	obj := patternFlowIpv4Dst{obj: &otg.PatternFlowIpv4Dst{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4Dst) msg() *otg.PatternFlowIpv4Dst {
	return obj.obj
}

func (obj *patternFlowIpv4Dst) setMsg(msg *otg.PatternFlowIpv4Dst) PatternFlowIpv4Dst {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4Dst struct {
	obj *patternFlowIpv4Dst
}

type marshalPatternFlowIpv4Dst interface {
	// ToProto marshals PatternFlowIpv4Dst to protobuf object *otg.PatternFlowIpv4Dst
	ToProto() (*otg.PatternFlowIpv4Dst, error)
	// ToPbText marshals PatternFlowIpv4Dst to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4Dst to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4Dst to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv4Dst to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv4Dst struct {
	obj *patternFlowIpv4Dst
}

type unMarshalPatternFlowIpv4Dst interface {
	// FromProto unmarshals PatternFlowIpv4Dst from protobuf object *otg.PatternFlowIpv4Dst
	FromProto(msg *otg.PatternFlowIpv4Dst) (PatternFlowIpv4Dst, error)
	// FromPbText unmarshals PatternFlowIpv4Dst from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4Dst from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4Dst from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4Dst) Marshal() marshalPatternFlowIpv4Dst {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4Dst{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4Dst) Unmarshal() unMarshalPatternFlowIpv4Dst {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4Dst{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4Dst) ToProto() (*otg.PatternFlowIpv4Dst, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4Dst) FromProto(msg *otg.PatternFlowIpv4Dst) (PatternFlowIpv4Dst, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4Dst) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4Dst) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4Dst) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4Dst) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4Dst) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv4Dst) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4Dst) FromJson(value string) error {
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

func (obj *patternFlowIpv4Dst) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4Dst) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4Dst) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4Dst) Clone() (PatternFlowIpv4Dst, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4Dst()
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

func (obj *patternFlowIpv4Dst) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.autoHolder = nil
	obj.randomHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4Dst is destination address
type PatternFlowIpv4Dst interface {
	Validation
	// msg marshals PatternFlowIpv4Dst to protobuf object *otg.PatternFlowIpv4Dst
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4Dst
	// setMsg unmarshals PatternFlowIpv4Dst from protobuf object *otg.PatternFlowIpv4Dst
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4Dst) PatternFlowIpv4Dst
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4Dst
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4Dst
	// validate validates PatternFlowIpv4Dst
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4Dst, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4DstChoiceEnum, set in PatternFlowIpv4Dst
	Choice() PatternFlowIpv4DstChoiceEnum
	// setChoice assigns PatternFlowIpv4DstChoiceEnum provided by user to PatternFlowIpv4Dst
	setChoice(value PatternFlowIpv4DstChoiceEnum) PatternFlowIpv4Dst
	// HasChoice checks if Choice has been set in PatternFlowIpv4Dst
	HasChoice() bool
	// Value returns string, set in PatternFlowIpv4Dst.
	Value() string
	// SetValue assigns string provided by user to PatternFlowIpv4Dst
	SetValue(value string) PatternFlowIpv4Dst
	// HasValue checks if Value has been set in PatternFlowIpv4Dst
	HasValue() bool
	// Values returns []string, set in PatternFlowIpv4Dst.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowIpv4Dst
	SetValues(value []string) PatternFlowIpv4Dst
	// Increment returns PatternFlowIpv4DstCounter, set in PatternFlowIpv4Dst.
	// PatternFlowIpv4DstCounter is ipv4 counter pattern
	Increment() PatternFlowIpv4DstCounter
	// SetIncrement assigns PatternFlowIpv4DstCounter provided by user to PatternFlowIpv4Dst.
	// PatternFlowIpv4DstCounter is ipv4 counter pattern
	SetIncrement(value PatternFlowIpv4DstCounter) PatternFlowIpv4Dst
	// HasIncrement checks if Increment has been set in PatternFlowIpv4Dst
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4DstCounter, set in PatternFlowIpv4Dst.
	// PatternFlowIpv4DstCounter is ipv4 counter pattern
	Decrement() PatternFlowIpv4DstCounter
	// SetDecrement assigns PatternFlowIpv4DstCounter provided by user to PatternFlowIpv4Dst.
	// PatternFlowIpv4DstCounter is ipv4 counter pattern
	SetDecrement(value PatternFlowIpv4DstCounter) PatternFlowIpv4Dst
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4Dst
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv4DstPatternFlowIpv4DstMetricTagIterIter, set in PatternFlowIpv4Dst
	MetricTags() PatternFlowIpv4DstPatternFlowIpv4DstMetricTagIter
	// Auto returns FlowIpv4Auto, set in PatternFlowIpv4Dst.
	// FlowIpv4Auto is the OTG implementation can provide a system generated, value for this property.
	Auto() FlowIpv4Auto
	// HasAuto checks if Auto has been set in PatternFlowIpv4Dst
	HasAuto() bool
	// Random returns PatternFlowIpv4DstRandom, set in PatternFlowIpv4Dst.
	// PatternFlowIpv4DstRandom is ipv4 random pattern
	Random() PatternFlowIpv4DstRandom
	// SetRandom assigns PatternFlowIpv4DstRandom provided by user to PatternFlowIpv4Dst.
	// PatternFlowIpv4DstRandom is ipv4 random pattern
	SetRandom(value PatternFlowIpv4DstRandom) PatternFlowIpv4Dst
	// HasRandom checks if Random has been set in PatternFlowIpv4Dst
	HasRandom() bool
	setNil()
}

type PatternFlowIpv4DstChoiceEnum string

// Enum of Choice on PatternFlowIpv4Dst
var PatternFlowIpv4DstChoice = struct {
	VALUE     PatternFlowIpv4DstChoiceEnum
	VALUES    PatternFlowIpv4DstChoiceEnum
	INCREMENT PatternFlowIpv4DstChoiceEnum
	DECREMENT PatternFlowIpv4DstChoiceEnum
	AUTO      PatternFlowIpv4DstChoiceEnum
	RANDOM    PatternFlowIpv4DstChoiceEnum
}{
	VALUE:     PatternFlowIpv4DstChoiceEnum("value"),
	VALUES:    PatternFlowIpv4DstChoiceEnum("values"),
	INCREMENT: PatternFlowIpv4DstChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4DstChoiceEnum("decrement"),
	AUTO:      PatternFlowIpv4DstChoiceEnum("auto"),
	RANDOM:    PatternFlowIpv4DstChoiceEnum("random"),
}

func (obj *patternFlowIpv4Dst) Choice() PatternFlowIpv4DstChoiceEnum {
	return PatternFlowIpv4DstChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4Dst) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4Dst) setChoice(value PatternFlowIpv4DstChoiceEnum) PatternFlowIpv4Dst {
	intValue, ok := otg.PatternFlowIpv4Dst_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4DstChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4Dst_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Random = nil
	obj.randomHolder = nil
	obj.obj.Auto = nil
	obj.autoHolder = nil
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4DstChoice.VALUE {
		defaultValue := "0.0.0.0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4DstChoice.VALUES {
		defaultValue := []string{"0.0.0.0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4DstChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4DstCounter().msg()
	}

	if value == PatternFlowIpv4DstChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4DstCounter().msg()
	}

	if value == PatternFlowIpv4DstChoice.AUTO {
		obj.obj.Auto = NewFlowIpv4Auto().msg()
	}

	if value == PatternFlowIpv4DstChoice.RANDOM {
		obj.obj.Random = NewPatternFlowIpv4DstRandom().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowIpv4Dst) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4DstChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowIpv4Dst) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowIpv4Dst object
func (obj *patternFlowIpv4Dst) SetValue(value string) PatternFlowIpv4Dst {
	obj.setChoice(PatternFlowIpv4DstChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowIpv4Dst) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"0.0.0.0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowIpv4Dst object
func (obj *patternFlowIpv4Dst) SetValues(value []string) PatternFlowIpv4Dst {
	obj.setChoice(PatternFlowIpv4DstChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv4DstCounter
func (obj *patternFlowIpv4Dst) Increment() PatternFlowIpv4DstCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4DstChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4DstCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4DstCounter
func (obj *patternFlowIpv4Dst) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4DstCounter value in the PatternFlowIpv4Dst object
func (obj *patternFlowIpv4Dst) SetIncrement(value PatternFlowIpv4DstCounter) PatternFlowIpv4Dst {
	obj.setChoice(PatternFlowIpv4DstChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4DstCounter
func (obj *patternFlowIpv4Dst) Decrement() PatternFlowIpv4DstCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4DstChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4DstCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4DstCounter
func (obj *patternFlowIpv4Dst) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4DstCounter value in the PatternFlowIpv4Dst object
func (obj *patternFlowIpv4Dst) SetDecrement(value PatternFlowIpv4DstCounter) PatternFlowIpv4Dst {
	obj.setChoice(PatternFlowIpv4DstChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv4DstMetricTag
func (obj *patternFlowIpv4Dst) MetricTags() PatternFlowIpv4DstPatternFlowIpv4DstMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv4DstMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv4DstPatternFlowIpv4DstMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv4DstPatternFlowIpv4DstMetricTagIter struct {
	obj                              *patternFlowIpv4Dst
	patternFlowIpv4DstMetricTagSlice []PatternFlowIpv4DstMetricTag
	fieldPtr                         *[]*otg.PatternFlowIpv4DstMetricTag
}

func newPatternFlowIpv4DstPatternFlowIpv4DstMetricTagIter(ptr *[]*otg.PatternFlowIpv4DstMetricTag) PatternFlowIpv4DstPatternFlowIpv4DstMetricTagIter {
	return &patternFlowIpv4DstPatternFlowIpv4DstMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv4DstPatternFlowIpv4DstMetricTagIter interface {
	setMsg(*patternFlowIpv4Dst) PatternFlowIpv4DstPatternFlowIpv4DstMetricTagIter
	Items() []PatternFlowIpv4DstMetricTag
	Add() PatternFlowIpv4DstMetricTag
	Append(items ...PatternFlowIpv4DstMetricTag) PatternFlowIpv4DstPatternFlowIpv4DstMetricTagIter
	Set(index int, newObj PatternFlowIpv4DstMetricTag) PatternFlowIpv4DstPatternFlowIpv4DstMetricTagIter
	Clear() PatternFlowIpv4DstPatternFlowIpv4DstMetricTagIter
	clearHolderSlice() PatternFlowIpv4DstPatternFlowIpv4DstMetricTagIter
	appendHolderSlice(item PatternFlowIpv4DstMetricTag) PatternFlowIpv4DstPatternFlowIpv4DstMetricTagIter
}

func (obj *patternFlowIpv4DstPatternFlowIpv4DstMetricTagIter) setMsg(msg *patternFlowIpv4Dst) PatternFlowIpv4DstPatternFlowIpv4DstMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv4DstMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv4DstPatternFlowIpv4DstMetricTagIter) Items() []PatternFlowIpv4DstMetricTag {
	return obj.patternFlowIpv4DstMetricTagSlice
}

func (obj *patternFlowIpv4DstPatternFlowIpv4DstMetricTagIter) Add() PatternFlowIpv4DstMetricTag {
	newObj := &otg.PatternFlowIpv4DstMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv4DstMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv4DstMetricTagSlice = append(obj.patternFlowIpv4DstMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv4DstPatternFlowIpv4DstMetricTagIter) Append(items ...PatternFlowIpv4DstMetricTag) PatternFlowIpv4DstPatternFlowIpv4DstMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv4DstMetricTagSlice = append(obj.patternFlowIpv4DstMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv4DstPatternFlowIpv4DstMetricTagIter) Set(index int, newObj PatternFlowIpv4DstMetricTag) PatternFlowIpv4DstPatternFlowIpv4DstMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv4DstMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv4DstPatternFlowIpv4DstMetricTagIter) Clear() PatternFlowIpv4DstPatternFlowIpv4DstMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv4DstMetricTag{}
		obj.patternFlowIpv4DstMetricTagSlice = []PatternFlowIpv4DstMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4DstPatternFlowIpv4DstMetricTagIter) clearHolderSlice() PatternFlowIpv4DstPatternFlowIpv4DstMetricTagIter {
	if len(obj.patternFlowIpv4DstMetricTagSlice) > 0 {
		obj.patternFlowIpv4DstMetricTagSlice = []PatternFlowIpv4DstMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4DstPatternFlowIpv4DstMetricTagIter) appendHolderSlice(item PatternFlowIpv4DstMetricTag) PatternFlowIpv4DstPatternFlowIpv4DstMetricTagIter {
	obj.patternFlowIpv4DstMetricTagSlice = append(obj.patternFlowIpv4DstMetricTagSlice, item)
	return obj
}

// description is TBD
// Auto returns a FlowIpv4Auto
func (obj *patternFlowIpv4Dst) Auto() FlowIpv4Auto {
	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIpv4DstChoice.AUTO)
	}
	if obj.autoHolder == nil {
		obj.autoHolder = &flowIpv4Auto{obj: obj.obj.Auto}
	}
	return obj.autoHolder
}

// description is TBD
// Auto returns a FlowIpv4Auto
func (obj *patternFlowIpv4Dst) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Random returns a PatternFlowIpv4DstRandom
func (obj *patternFlowIpv4Dst) Random() PatternFlowIpv4DstRandom {
	if obj.obj.Random == nil {
		obj.setChoice(PatternFlowIpv4DstChoice.RANDOM)
	}
	if obj.randomHolder == nil {
		obj.randomHolder = &patternFlowIpv4DstRandom{obj: obj.obj.Random}
	}
	return obj.randomHolder
}

// description is TBD
// Random returns a PatternFlowIpv4DstRandom
func (obj *patternFlowIpv4Dst) HasRandom() bool {
	return obj.obj.Random != nil
}

// description is TBD
// SetRandom sets the PatternFlowIpv4DstRandom value in the PatternFlowIpv4Dst object
func (obj *patternFlowIpv4Dst) SetRandom(value PatternFlowIpv4DstRandom) PatternFlowIpv4Dst {
	obj.setChoice(PatternFlowIpv4DstChoice.RANDOM)
	obj.randomHolder = nil
	obj.obj.Random = value.msg()

	return obj
}

func (obj *patternFlowIpv4Dst) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv4Dst.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowIpv4Dst.Values"))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv4DstMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.Auto != nil {

		obj.Auto().validateObj(vObj, set_default)
	}

	if obj.obj.Random != nil {

		obj.Random().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv4Dst) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv4DstChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv4DstChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv4DstChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv4DstChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv4DstChoice.DECREMENT
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIpv4DstChoice.AUTO
	}

	if obj.obj.Random != nil {
		choices_set += 1
		choice = PatternFlowIpv4DstChoice.RANDOM
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv4DstChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv4Dst")
			}
		} else {
			intVal := otg.PatternFlowIpv4Dst_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv4Dst_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
