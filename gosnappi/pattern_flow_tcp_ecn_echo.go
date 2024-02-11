package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpEcnEcho *****
type patternFlowTcpEcnEcho struct {
	validation
	obj              *otg.PatternFlowTcpEcnEcho
	marshaller       marshalPatternFlowTcpEcnEcho
	unMarshaller     unMarshalPatternFlowTcpEcnEcho
	incrementHolder  PatternFlowTcpEcnEchoCounter
	decrementHolder  PatternFlowTcpEcnEchoCounter
	metricTagsHolder PatternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter
}

func NewPatternFlowTcpEcnEcho() PatternFlowTcpEcnEcho {
	obj := patternFlowTcpEcnEcho{obj: &otg.PatternFlowTcpEcnEcho{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpEcnEcho) msg() *otg.PatternFlowTcpEcnEcho {
	return obj.obj
}

func (obj *patternFlowTcpEcnEcho) setMsg(msg *otg.PatternFlowTcpEcnEcho) PatternFlowTcpEcnEcho {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpEcnEcho struct {
	obj *patternFlowTcpEcnEcho
}

type marshalPatternFlowTcpEcnEcho interface {
	// ToProto marshals PatternFlowTcpEcnEcho to protobuf object *otg.PatternFlowTcpEcnEcho
	ToProto() (*otg.PatternFlowTcpEcnEcho, error)
	// ToPbText marshals PatternFlowTcpEcnEcho to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpEcnEcho to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpEcnEcho to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpEcnEcho struct {
	obj *patternFlowTcpEcnEcho
}

type unMarshalPatternFlowTcpEcnEcho interface {
	// FromProto unmarshals PatternFlowTcpEcnEcho from protobuf object *otg.PatternFlowTcpEcnEcho
	FromProto(msg *otg.PatternFlowTcpEcnEcho) (PatternFlowTcpEcnEcho, error)
	// FromPbText unmarshals PatternFlowTcpEcnEcho from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpEcnEcho from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpEcnEcho from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpEcnEcho) Marshal() marshalPatternFlowTcpEcnEcho {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpEcnEcho{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpEcnEcho) Unmarshal() unMarshalPatternFlowTcpEcnEcho {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpEcnEcho{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpEcnEcho) ToProto() (*otg.PatternFlowTcpEcnEcho, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpEcnEcho) FromProto(msg *otg.PatternFlowTcpEcnEcho) (PatternFlowTcpEcnEcho, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpEcnEcho) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpEcnEcho) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpEcnEcho) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpEcnEcho) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpEcnEcho) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpEcnEcho) FromJson(value string) error {
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

func (obj *patternFlowTcpEcnEcho) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpEcnEcho) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpEcnEcho) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpEcnEcho) Clone() (PatternFlowTcpEcnEcho, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpEcnEcho()
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

func (obj *patternFlowTcpEcnEcho) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowTcpEcnEcho is explicit congestion notification, echo. 1 indicates the peer is ecn capable. 0 indicates that a packet with ipv4.ecn = 11 in the ip header was  received during normal transmission.
type PatternFlowTcpEcnEcho interface {
	Validation
	// msg marshals PatternFlowTcpEcnEcho to protobuf object *otg.PatternFlowTcpEcnEcho
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpEcnEcho
	// setMsg unmarshals PatternFlowTcpEcnEcho from protobuf object *otg.PatternFlowTcpEcnEcho
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpEcnEcho) PatternFlowTcpEcnEcho
	// provides marshal interface
	Marshal() marshalPatternFlowTcpEcnEcho
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpEcnEcho
	// validate validates PatternFlowTcpEcnEcho
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpEcnEcho, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowTcpEcnEchoChoiceEnum, set in PatternFlowTcpEcnEcho
	Choice() PatternFlowTcpEcnEchoChoiceEnum
	// setChoice assigns PatternFlowTcpEcnEchoChoiceEnum provided by user to PatternFlowTcpEcnEcho
	setChoice(value PatternFlowTcpEcnEchoChoiceEnum) PatternFlowTcpEcnEcho
	// HasChoice checks if Choice has been set in PatternFlowTcpEcnEcho
	HasChoice() bool
	// Value returns uint32, set in PatternFlowTcpEcnEcho.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowTcpEcnEcho
	SetValue(value uint32) PatternFlowTcpEcnEcho
	// HasValue checks if Value has been set in PatternFlowTcpEcnEcho
	HasValue() bool
	// Values returns []uint32, set in PatternFlowTcpEcnEcho.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowTcpEcnEcho
	SetValues(value []uint32) PatternFlowTcpEcnEcho
	// Increment returns PatternFlowTcpEcnEchoCounter, set in PatternFlowTcpEcnEcho.
	// PatternFlowTcpEcnEchoCounter is integer counter pattern
	Increment() PatternFlowTcpEcnEchoCounter
	// SetIncrement assigns PatternFlowTcpEcnEchoCounter provided by user to PatternFlowTcpEcnEcho.
	// PatternFlowTcpEcnEchoCounter is integer counter pattern
	SetIncrement(value PatternFlowTcpEcnEchoCounter) PatternFlowTcpEcnEcho
	// HasIncrement checks if Increment has been set in PatternFlowTcpEcnEcho
	HasIncrement() bool
	// Decrement returns PatternFlowTcpEcnEchoCounter, set in PatternFlowTcpEcnEcho.
	// PatternFlowTcpEcnEchoCounter is integer counter pattern
	Decrement() PatternFlowTcpEcnEchoCounter
	// SetDecrement assigns PatternFlowTcpEcnEchoCounter provided by user to PatternFlowTcpEcnEcho.
	// PatternFlowTcpEcnEchoCounter is integer counter pattern
	SetDecrement(value PatternFlowTcpEcnEchoCounter) PatternFlowTcpEcnEcho
	// HasDecrement checks if Decrement has been set in PatternFlowTcpEcnEcho
	HasDecrement() bool
	// MetricTags returns PatternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIterIter, set in PatternFlowTcpEcnEcho
	MetricTags() PatternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter
	setNil()
}

type PatternFlowTcpEcnEchoChoiceEnum string

// Enum of Choice on PatternFlowTcpEcnEcho
var PatternFlowTcpEcnEchoChoice = struct {
	VALUE     PatternFlowTcpEcnEchoChoiceEnum
	VALUES    PatternFlowTcpEcnEchoChoiceEnum
	INCREMENT PatternFlowTcpEcnEchoChoiceEnum
	DECREMENT PatternFlowTcpEcnEchoChoiceEnum
}{
	VALUE:     PatternFlowTcpEcnEchoChoiceEnum("value"),
	VALUES:    PatternFlowTcpEcnEchoChoiceEnum("values"),
	INCREMENT: PatternFlowTcpEcnEchoChoiceEnum("increment"),
	DECREMENT: PatternFlowTcpEcnEchoChoiceEnum("decrement"),
}

func (obj *patternFlowTcpEcnEcho) Choice() PatternFlowTcpEcnEchoChoiceEnum {
	return PatternFlowTcpEcnEchoChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowTcpEcnEcho) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowTcpEcnEcho) setChoice(value PatternFlowTcpEcnEchoChoiceEnum) PatternFlowTcpEcnEcho {
	intValue, ok := otg.PatternFlowTcpEcnEcho_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowTcpEcnEchoChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowTcpEcnEcho_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowTcpEcnEchoChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowTcpEcnEchoChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowTcpEcnEchoChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowTcpEcnEchoCounter().msg()
	}

	if value == PatternFlowTcpEcnEchoChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowTcpEcnEchoCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpEcnEcho) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowTcpEcnEchoChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpEcnEcho) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowTcpEcnEcho object
func (obj *patternFlowTcpEcnEcho) SetValue(value uint32) PatternFlowTcpEcnEcho {
	obj.setChoice(PatternFlowTcpEcnEchoChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowTcpEcnEcho) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowTcpEcnEcho object
func (obj *patternFlowTcpEcnEcho) SetValues(value []uint32) PatternFlowTcpEcnEcho {
	obj.setChoice(PatternFlowTcpEcnEchoChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowTcpEcnEchoCounter
func (obj *patternFlowTcpEcnEcho) Increment() PatternFlowTcpEcnEchoCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowTcpEcnEchoChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowTcpEcnEchoCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowTcpEcnEchoCounter
func (obj *patternFlowTcpEcnEcho) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowTcpEcnEchoCounter value in the PatternFlowTcpEcnEcho object
func (obj *patternFlowTcpEcnEcho) SetIncrement(value PatternFlowTcpEcnEchoCounter) PatternFlowTcpEcnEcho {
	obj.setChoice(PatternFlowTcpEcnEchoChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowTcpEcnEchoCounter
func (obj *patternFlowTcpEcnEcho) Decrement() PatternFlowTcpEcnEchoCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowTcpEcnEchoChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowTcpEcnEchoCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowTcpEcnEchoCounter
func (obj *patternFlowTcpEcnEcho) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowTcpEcnEchoCounter value in the PatternFlowTcpEcnEcho object
func (obj *patternFlowTcpEcnEcho) SetDecrement(value PatternFlowTcpEcnEchoCounter) PatternFlowTcpEcnEcho {
	obj.setChoice(PatternFlowTcpEcnEchoChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowTcpEcnEchoMetricTag
func (obj *patternFlowTcpEcnEcho) MetricTags() PatternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowTcpEcnEchoMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter struct {
	obj                                 *patternFlowTcpEcnEcho
	patternFlowTcpEcnEchoMetricTagSlice []PatternFlowTcpEcnEchoMetricTag
	fieldPtr                            *[]*otg.PatternFlowTcpEcnEchoMetricTag
}

func newPatternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter(ptr *[]*otg.PatternFlowTcpEcnEchoMetricTag) PatternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter {
	return &patternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter{fieldPtr: ptr}
}

type PatternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter interface {
	setMsg(*patternFlowTcpEcnEcho) PatternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter
	Items() []PatternFlowTcpEcnEchoMetricTag
	Add() PatternFlowTcpEcnEchoMetricTag
	Append(items ...PatternFlowTcpEcnEchoMetricTag) PatternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter
	Set(index int, newObj PatternFlowTcpEcnEchoMetricTag) PatternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter
	Clear() PatternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter
	clearHolderSlice() PatternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter
	appendHolderSlice(item PatternFlowTcpEcnEchoMetricTag) PatternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter
}

func (obj *patternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter) setMsg(msg *patternFlowTcpEcnEcho) PatternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowTcpEcnEchoMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter) Items() []PatternFlowTcpEcnEchoMetricTag {
	return obj.patternFlowTcpEcnEchoMetricTagSlice
}

func (obj *patternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter) Add() PatternFlowTcpEcnEchoMetricTag {
	newObj := &otg.PatternFlowTcpEcnEchoMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowTcpEcnEchoMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowTcpEcnEchoMetricTagSlice = append(obj.patternFlowTcpEcnEchoMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter) Append(items ...PatternFlowTcpEcnEchoMetricTag) PatternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowTcpEcnEchoMetricTagSlice = append(obj.patternFlowTcpEcnEchoMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter) Set(index int, newObj PatternFlowTcpEcnEchoMetricTag) PatternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowTcpEcnEchoMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter) Clear() PatternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowTcpEcnEchoMetricTag{}
		obj.patternFlowTcpEcnEchoMetricTagSlice = []PatternFlowTcpEcnEchoMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter) clearHolderSlice() PatternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter {
	if len(obj.patternFlowTcpEcnEchoMetricTagSlice) > 0 {
		obj.patternFlowTcpEcnEchoMetricTagSlice = []PatternFlowTcpEcnEchoMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter) appendHolderSlice(item PatternFlowTcpEcnEchoMetricTag) PatternFlowTcpEcnEchoPatternFlowTcpEcnEchoMetricTagIter {
	obj.patternFlowTcpEcnEchoMetricTagSlice = append(obj.patternFlowTcpEcnEchoMetricTagSlice, item)
	return obj
}

func (obj *patternFlowTcpEcnEcho) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpEcnEcho.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowTcpEcnEcho.Values <= 1 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowTcpEcnEchoMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowTcpEcnEcho) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowTcpEcnEchoChoice.VALUE)

	}

}
