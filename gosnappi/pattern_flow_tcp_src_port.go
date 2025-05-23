package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpSrcPort *****
type patternFlowTcpSrcPort struct {
	validation
	obj              *otg.PatternFlowTcpSrcPort
	marshaller       marshalPatternFlowTcpSrcPort
	unMarshaller     unMarshalPatternFlowTcpSrcPort
	incrementHolder  PatternFlowTcpSrcPortCounter
	decrementHolder  PatternFlowTcpSrcPortCounter
	metricTagsHolder PatternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter
	randomHolder     PatternFlowTcpSrcPortRandom
}

func NewPatternFlowTcpSrcPort() PatternFlowTcpSrcPort {
	obj := patternFlowTcpSrcPort{obj: &otg.PatternFlowTcpSrcPort{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpSrcPort) msg() *otg.PatternFlowTcpSrcPort {
	return obj.obj
}

func (obj *patternFlowTcpSrcPort) setMsg(msg *otg.PatternFlowTcpSrcPort) PatternFlowTcpSrcPort {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpSrcPort struct {
	obj *patternFlowTcpSrcPort
}

type marshalPatternFlowTcpSrcPort interface {
	// ToProto marshals PatternFlowTcpSrcPort to protobuf object *otg.PatternFlowTcpSrcPort
	ToProto() (*otg.PatternFlowTcpSrcPort, error)
	// ToPbText marshals PatternFlowTcpSrcPort to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpSrcPort to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpSrcPort to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowTcpSrcPort to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowTcpSrcPort struct {
	obj *patternFlowTcpSrcPort
}

type unMarshalPatternFlowTcpSrcPort interface {
	// FromProto unmarshals PatternFlowTcpSrcPort from protobuf object *otg.PatternFlowTcpSrcPort
	FromProto(msg *otg.PatternFlowTcpSrcPort) (PatternFlowTcpSrcPort, error)
	// FromPbText unmarshals PatternFlowTcpSrcPort from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpSrcPort from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpSrcPort from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpSrcPort) Marshal() marshalPatternFlowTcpSrcPort {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpSrcPort{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpSrcPort) Unmarshal() unMarshalPatternFlowTcpSrcPort {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpSrcPort{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpSrcPort) ToProto() (*otg.PatternFlowTcpSrcPort, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpSrcPort) FromProto(msg *otg.PatternFlowTcpSrcPort) (PatternFlowTcpSrcPort, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpSrcPort) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpSrcPort) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpSrcPort) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpSrcPort) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpSrcPort) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowTcpSrcPort) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpSrcPort) FromJson(value string) error {
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

func (obj *patternFlowTcpSrcPort) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpSrcPort) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpSrcPort) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpSrcPort) Clone() (PatternFlowTcpSrcPort, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpSrcPort()
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

func (obj *patternFlowTcpSrcPort) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.randomHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowTcpSrcPort is source port
type PatternFlowTcpSrcPort interface {
	Validation
	// msg marshals PatternFlowTcpSrcPort to protobuf object *otg.PatternFlowTcpSrcPort
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpSrcPort
	// setMsg unmarshals PatternFlowTcpSrcPort from protobuf object *otg.PatternFlowTcpSrcPort
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpSrcPort) PatternFlowTcpSrcPort
	// provides marshal interface
	Marshal() marshalPatternFlowTcpSrcPort
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpSrcPort
	// validate validates PatternFlowTcpSrcPort
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpSrcPort, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowTcpSrcPortChoiceEnum, set in PatternFlowTcpSrcPort
	Choice() PatternFlowTcpSrcPortChoiceEnum
	// setChoice assigns PatternFlowTcpSrcPortChoiceEnum provided by user to PatternFlowTcpSrcPort
	setChoice(value PatternFlowTcpSrcPortChoiceEnum) PatternFlowTcpSrcPort
	// HasChoice checks if Choice has been set in PatternFlowTcpSrcPort
	HasChoice() bool
	// Value returns uint32, set in PatternFlowTcpSrcPort.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowTcpSrcPort
	SetValue(value uint32) PatternFlowTcpSrcPort
	// HasValue checks if Value has been set in PatternFlowTcpSrcPort
	HasValue() bool
	// Values returns []uint32, set in PatternFlowTcpSrcPort.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowTcpSrcPort
	SetValues(value []uint32) PatternFlowTcpSrcPort
	// Increment returns PatternFlowTcpSrcPortCounter, set in PatternFlowTcpSrcPort.
	// PatternFlowTcpSrcPortCounter is integer counter pattern
	Increment() PatternFlowTcpSrcPortCounter
	// SetIncrement assigns PatternFlowTcpSrcPortCounter provided by user to PatternFlowTcpSrcPort.
	// PatternFlowTcpSrcPortCounter is integer counter pattern
	SetIncrement(value PatternFlowTcpSrcPortCounter) PatternFlowTcpSrcPort
	// HasIncrement checks if Increment has been set in PatternFlowTcpSrcPort
	HasIncrement() bool
	// Decrement returns PatternFlowTcpSrcPortCounter, set in PatternFlowTcpSrcPort.
	// PatternFlowTcpSrcPortCounter is integer counter pattern
	Decrement() PatternFlowTcpSrcPortCounter
	// SetDecrement assigns PatternFlowTcpSrcPortCounter provided by user to PatternFlowTcpSrcPort.
	// PatternFlowTcpSrcPortCounter is integer counter pattern
	SetDecrement(value PatternFlowTcpSrcPortCounter) PatternFlowTcpSrcPort
	// HasDecrement checks if Decrement has been set in PatternFlowTcpSrcPort
	HasDecrement() bool
	// MetricTags returns PatternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIterIter, set in PatternFlowTcpSrcPort
	MetricTags() PatternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter
	// Random returns PatternFlowTcpSrcPortRandom, set in PatternFlowTcpSrcPort.
	// PatternFlowTcpSrcPortRandom is integer random pattern
	Random() PatternFlowTcpSrcPortRandom
	// SetRandom assigns PatternFlowTcpSrcPortRandom provided by user to PatternFlowTcpSrcPort.
	// PatternFlowTcpSrcPortRandom is integer random pattern
	SetRandom(value PatternFlowTcpSrcPortRandom) PatternFlowTcpSrcPort
	// HasRandom checks if Random has been set in PatternFlowTcpSrcPort
	HasRandom() bool
	setNil()
}

type PatternFlowTcpSrcPortChoiceEnum string

// Enum of Choice on PatternFlowTcpSrcPort
var PatternFlowTcpSrcPortChoice = struct {
	VALUE     PatternFlowTcpSrcPortChoiceEnum
	VALUES    PatternFlowTcpSrcPortChoiceEnum
	INCREMENT PatternFlowTcpSrcPortChoiceEnum
	DECREMENT PatternFlowTcpSrcPortChoiceEnum
	RANDOM    PatternFlowTcpSrcPortChoiceEnum
}{
	VALUE:     PatternFlowTcpSrcPortChoiceEnum("value"),
	VALUES:    PatternFlowTcpSrcPortChoiceEnum("values"),
	INCREMENT: PatternFlowTcpSrcPortChoiceEnum("increment"),
	DECREMENT: PatternFlowTcpSrcPortChoiceEnum("decrement"),
	RANDOM:    PatternFlowTcpSrcPortChoiceEnum("random"),
}

func (obj *patternFlowTcpSrcPort) Choice() PatternFlowTcpSrcPortChoiceEnum {
	return PatternFlowTcpSrcPortChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowTcpSrcPort) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowTcpSrcPort) setChoice(value PatternFlowTcpSrcPortChoiceEnum) PatternFlowTcpSrcPort {
	intValue, ok := otg.PatternFlowTcpSrcPort_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowTcpSrcPortChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowTcpSrcPort_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Random = nil
	obj.randomHolder = nil
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowTcpSrcPortChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowTcpSrcPortChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowTcpSrcPortChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowTcpSrcPortCounter().msg()
	}

	if value == PatternFlowTcpSrcPortChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowTcpSrcPortCounter().msg()
	}

	if value == PatternFlowTcpSrcPortChoice.RANDOM {
		obj.obj.Random = NewPatternFlowTcpSrcPortRandom().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpSrcPort) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowTcpSrcPortChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpSrcPort) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowTcpSrcPort object
func (obj *patternFlowTcpSrcPort) SetValue(value uint32) PatternFlowTcpSrcPort {
	obj.setChoice(PatternFlowTcpSrcPortChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowTcpSrcPort) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowTcpSrcPort object
func (obj *patternFlowTcpSrcPort) SetValues(value []uint32) PatternFlowTcpSrcPort {
	obj.setChoice(PatternFlowTcpSrcPortChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowTcpSrcPortCounter
func (obj *patternFlowTcpSrcPort) Increment() PatternFlowTcpSrcPortCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowTcpSrcPortChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowTcpSrcPortCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowTcpSrcPortCounter
func (obj *patternFlowTcpSrcPort) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowTcpSrcPortCounter value in the PatternFlowTcpSrcPort object
func (obj *patternFlowTcpSrcPort) SetIncrement(value PatternFlowTcpSrcPortCounter) PatternFlowTcpSrcPort {
	obj.setChoice(PatternFlowTcpSrcPortChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowTcpSrcPortCounter
func (obj *patternFlowTcpSrcPort) Decrement() PatternFlowTcpSrcPortCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowTcpSrcPortChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowTcpSrcPortCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowTcpSrcPortCounter
func (obj *patternFlowTcpSrcPort) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowTcpSrcPortCounter value in the PatternFlowTcpSrcPort object
func (obj *patternFlowTcpSrcPort) SetDecrement(value PatternFlowTcpSrcPortCounter) PatternFlowTcpSrcPort {
	obj.setChoice(PatternFlowTcpSrcPortChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowTcpSrcPortMetricTag
func (obj *patternFlowTcpSrcPort) MetricTags() PatternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowTcpSrcPortMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter struct {
	obj                                 *patternFlowTcpSrcPort
	patternFlowTcpSrcPortMetricTagSlice []PatternFlowTcpSrcPortMetricTag
	fieldPtr                            *[]*otg.PatternFlowTcpSrcPortMetricTag
}

func newPatternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter(ptr *[]*otg.PatternFlowTcpSrcPortMetricTag) PatternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter {
	return &patternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter{fieldPtr: ptr}
}

type PatternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter interface {
	setMsg(*patternFlowTcpSrcPort) PatternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter
	Items() []PatternFlowTcpSrcPortMetricTag
	Add() PatternFlowTcpSrcPortMetricTag
	Append(items ...PatternFlowTcpSrcPortMetricTag) PatternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter
	Set(index int, newObj PatternFlowTcpSrcPortMetricTag) PatternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter
	Clear() PatternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter
	clearHolderSlice() PatternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter
	appendHolderSlice(item PatternFlowTcpSrcPortMetricTag) PatternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter
}

func (obj *patternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter) setMsg(msg *patternFlowTcpSrcPort) PatternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowTcpSrcPortMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter) Items() []PatternFlowTcpSrcPortMetricTag {
	return obj.patternFlowTcpSrcPortMetricTagSlice
}

func (obj *patternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter) Add() PatternFlowTcpSrcPortMetricTag {
	newObj := &otg.PatternFlowTcpSrcPortMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowTcpSrcPortMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowTcpSrcPortMetricTagSlice = append(obj.patternFlowTcpSrcPortMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter) Append(items ...PatternFlowTcpSrcPortMetricTag) PatternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowTcpSrcPortMetricTagSlice = append(obj.patternFlowTcpSrcPortMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter) Set(index int, newObj PatternFlowTcpSrcPortMetricTag) PatternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowTcpSrcPortMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter) Clear() PatternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowTcpSrcPortMetricTag{}
		obj.patternFlowTcpSrcPortMetricTagSlice = []PatternFlowTcpSrcPortMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter) clearHolderSlice() PatternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter {
	if len(obj.patternFlowTcpSrcPortMetricTagSlice) > 0 {
		obj.patternFlowTcpSrcPortMetricTagSlice = []PatternFlowTcpSrcPortMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter) appendHolderSlice(item PatternFlowTcpSrcPortMetricTag) PatternFlowTcpSrcPortPatternFlowTcpSrcPortMetricTagIter {
	obj.patternFlowTcpSrcPortMetricTagSlice = append(obj.patternFlowTcpSrcPortMetricTagSlice, item)
	return obj
}

// description is TBD
// Random returns a PatternFlowTcpSrcPortRandom
func (obj *patternFlowTcpSrcPort) Random() PatternFlowTcpSrcPortRandom {
	if obj.obj.Random == nil {
		obj.setChoice(PatternFlowTcpSrcPortChoice.RANDOM)
	}
	if obj.randomHolder == nil {
		obj.randomHolder = &patternFlowTcpSrcPortRandom{obj: obj.obj.Random}
	}
	return obj.randomHolder
}

// description is TBD
// Random returns a PatternFlowTcpSrcPortRandom
func (obj *patternFlowTcpSrcPort) HasRandom() bool {
	return obj.obj.Random != nil
}

// description is TBD
// SetRandom sets the PatternFlowTcpSrcPortRandom value in the PatternFlowTcpSrcPort object
func (obj *patternFlowTcpSrcPort) SetRandom(value PatternFlowTcpSrcPortRandom) PatternFlowTcpSrcPort {
	obj.setChoice(PatternFlowTcpSrcPortChoice.RANDOM)
	obj.randomHolder = nil
	obj.obj.Random = value.msg()

	return obj
}

func (obj *patternFlowTcpSrcPort) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpSrcPort.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowTcpSrcPort.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowTcpSrcPortMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.Random != nil {

		obj.Random().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowTcpSrcPort) setDefault() {
	var choices_set int = 0
	var choice PatternFlowTcpSrcPortChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowTcpSrcPortChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowTcpSrcPortChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowTcpSrcPortChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowTcpSrcPortChoice.DECREMENT
	}

	if obj.obj.Random != nil {
		choices_set += 1
		choice = PatternFlowTcpSrcPortChoice.RANDOM
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowTcpSrcPortChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowTcpSrcPort")
			}
		} else {
			intVal := otg.PatternFlowTcpSrcPort_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowTcpSrcPort_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
