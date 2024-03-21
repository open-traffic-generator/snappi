package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowVlanTpid *****
type patternFlowVlanTpid struct {
	validation
	obj              *otg.PatternFlowVlanTpid
	marshaller       marshalPatternFlowVlanTpid
	unMarshaller     unMarshalPatternFlowVlanTpid
	incrementHolder  PatternFlowVlanTpidCounter
	decrementHolder  PatternFlowVlanTpidCounter
	metricTagsHolder PatternFlowVlanTpidPatternFlowVlanTpidMetricTagIter
}

func NewPatternFlowVlanTpid() PatternFlowVlanTpid {
	obj := patternFlowVlanTpid{obj: &otg.PatternFlowVlanTpid{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowVlanTpid) msg() *otg.PatternFlowVlanTpid {
	return obj.obj
}

func (obj *patternFlowVlanTpid) setMsg(msg *otg.PatternFlowVlanTpid) PatternFlowVlanTpid {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowVlanTpid struct {
	obj *patternFlowVlanTpid
}

type marshalPatternFlowVlanTpid interface {
	// ToProto marshals PatternFlowVlanTpid to protobuf object *otg.PatternFlowVlanTpid
	ToProto() (*otg.PatternFlowVlanTpid, error)
	// ToPbText marshals PatternFlowVlanTpid to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowVlanTpid to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowVlanTpid to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowVlanTpid struct {
	obj *patternFlowVlanTpid
}

type unMarshalPatternFlowVlanTpid interface {
	// FromProto unmarshals PatternFlowVlanTpid from protobuf object *otg.PatternFlowVlanTpid
	FromProto(msg *otg.PatternFlowVlanTpid) (PatternFlowVlanTpid, error)
	// FromPbText unmarshals PatternFlowVlanTpid from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowVlanTpid from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowVlanTpid from JSON text
	FromJson(value string) error
}

func (obj *patternFlowVlanTpid) Marshal() marshalPatternFlowVlanTpid {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowVlanTpid{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowVlanTpid) Unmarshal() unMarshalPatternFlowVlanTpid {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowVlanTpid{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowVlanTpid) ToProto() (*otg.PatternFlowVlanTpid, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowVlanTpid) FromProto(msg *otg.PatternFlowVlanTpid) (PatternFlowVlanTpid, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowVlanTpid) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowVlanTpid) FromPbText(value string) error {
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

func (m *marshalpatternFlowVlanTpid) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowVlanTpid) FromYaml(value string) error {
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

func (m *marshalpatternFlowVlanTpid) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowVlanTpid) FromJson(value string) error {
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

func (obj *patternFlowVlanTpid) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowVlanTpid) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowVlanTpid) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowVlanTpid) Clone() (PatternFlowVlanTpid, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowVlanTpid()
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

func (obj *patternFlowVlanTpid) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowVlanTpid is protocol identifier
type PatternFlowVlanTpid interface {
	Validation
	// msg marshals PatternFlowVlanTpid to protobuf object *otg.PatternFlowVlanTpid
	// and doesn't set defaults
	msg() *otg.PatternFlowVlanTpid
	// setMsg unmarshals PatternFlowVlanTpid from protobuf object *otg.PatternFlowVlanTpid
	// and doesn't set defaults
	setMsg(*otg.PatternFlowVlanTpid) PatternFlowVlanTpid
	// provides marshal interface
	Marshal() marshalPatternFlowVlanTpid
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowVlanTpid
	// validate validates PatternFlowVlanTpid
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowVlanTpid, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowVlanTpidChoiceEnum, set in PatternFlowVlanTpid
	Choice() PatternFlowVlanTpidChoiceEnum
	// setChoice assigns PatternFlowVlanTpidChoiceEnum provided by user to PatternFlowVlanTpid
	setChoice(value PatternFlowVlanTpidChoiceEnum) PatternFlowVlanTpid
	// HasChoice checks if Choice has been set in PatternFlowVlanTpid
	HasChoice() bool
	// Value returns uint32, set in PatternFlowVlanTpid.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowVlanTpid
	SetValue(value uint32) PatternFlowVlanTpid
	// HasValue checks if Value has been set in PatternFlowVlanTpid
	HasValue() bool
	// Values returns []uint32, set in PatternFlowVlanTpid.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowVlanTpid
	SetValues(value []uint32) PatternFlowVlanTpid
	// Increment returns PatternFlowVlanTpidCounter, set in PatternFlowVlanTpid.
	// PatternFlowVlanTpidCounter is integer counter pattern
	Increment() PatternFlowVlanTpidCounter
	// SetIncrement assigns PatternFlowVlanTpidCounter provided by user to PatternFlowVlanTpid.
	// PatternFlowVlanTpidCounter is integer counter pattern
	SetIncrement(value PatternFlowVlanTpidCounter) PatternFlowVlanTpid
	// HasIncrement checks if Increment has been set in PatternFlowVlanTpid
	HasIncrement() bool
	// Decrement returns PatternFlowVlanTpidCounter, set in PatternFlowVlanTpid.
	// PatternFlowVlanTpidCounter is integer counter pattern
	Decrement() PatternFlowVlanTpidCounter
	// SetDecrement assigns PatternFlowVlanTpidCounter provided by user to PatternFlowVlanTpid.
	// PatternFlowVlanTpidCounter is integer counter pattern
	SetDecrement(value PatternFlowVlanTpidCounter) PatternFlowVlanTpid
	// HasDecrement checks if Decrement has been set in PatternFlowVlanTpid
	HasDecrement() bool
	// MetricTags returns PatternFlowVlanTpidPatternFlowVlanTpidMetricTagIterIter, set in PatternFlowVlanTpid
	MetricTags() PatternFlowVlanTpidPatternFlowVlanTpidMetricTagIter
	setNil()
}

type PatternFlowVlanTpidChoiceEnum string

// Enum of Choice on PatternFlowVlanTpid
var PatternFlowVlanTpidChoice = struct {
	VALUE     PatternFlowVlanTpidChoiceEnum
	VALUES    PatternFlowVlanTpidChoiceEnum
	INCREMENT PatternFlowVlanTpidChoiceEnum
	DECREMENT PatternFlowVlanTpidChoiceEnum
}{
	VALUE:     PatternFlowVlanTpidChoiceEnum("value"),
	VALUES:    PatternFlowVlanTpidChoiceEnum("values"),
	INCREMENT: PatternFlowVlanTpidChoiceEnum("increment"),
	DECREMENT: PatternFlowVlanTpidChoiceEnum("decrement"),
}

func (obj *patternFlowVlanTpid) Choice() PatternFlowVlanTpidChoiceEnum {
	return PatternFlowVlanTpidChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowVlanTpid) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowVlanTpid) setChoice(value PatternFlowVlanTpidChoiceEnum) PatternFlowVlanTpid {
	intValue, ok := otg.PatternFlowVlanTpid_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowVlanTpidChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowVlanTpid_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowVlanTpidChoice.VALUE {
		defaultValue := uint32(65535)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowVlanTpidChoice.VALUES {
		defaultValue := []uint32{65535}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowVlanTpidChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowVlanTpidCounter().msg()
	}

	if value == PatternFlowVlanTpidChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowVlanTpidCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowVlanTpid) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowVlanTpidChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowVlanTpid) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowVlanTpid object
func (obj *patternFlowVlanTpid) SetValue(value uint32) PatternFlowVlanTpid {
	obj.setChoice(PatternFlowVlanTpidChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowVlanTpid) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{65535})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowVlanTpid object
func (obj *patternFlowVlanTpid) SetValues(value []uint32) PatternFlowVlanTpid {
	obj.setChoice(PatternFlowVlanTpidChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowVlanTpidCounter
func (obj *patternFlowVlanTpid) Increment() PatternFlowVlanTpidCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowVlanTpidChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowVlanTpidCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowVlanTpidCounter
func (obj *patternFlowVlanTpid) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowVlanTpidCounter value in the PatternFlowVlanTpid object
func (obj *patternFlowVlanTpid) SetIncrement(value PatternFlowVlanTpidCounter) PatternFlowVlanTpid {
	obj.setChoice(PatternFlowVlanTpidChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowVlanTpidCounter
func (obj *patternFlowVlanTpid) Decrement() PatternFlowVlanTpidCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowVlanTpidChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowVlanTpidCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowVlanTpidCounter
func (obj *patternFlowVlanTpid) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowVlanTpidCounter value in the PatternFlowVlanTpid object
func (obj *patternFlowVlanTpid) SetDecrement(value PatternFlowVlanTpidCounter) PatternFlowVlanTpid {
	obj.setChoice(PatternFlowVlanTpidChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowVlanTpidMetricTag
func (obj *patternFlowVlanTpid) MetricTags() PatternFlowVlanTpidPatternFlowVlanTpidMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowVlanTpidMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowVlanTpidPatternFlowVlanTpidMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowVlanTpidPatternFlowVlanTpidMetricTagIter struct {
	obj                               *patternFlowVlanTpid
	patternFlowVlanTpidMetricTagSlice []PatternFlowVlanTpidMetricTag
	fieldPtr                          *[]*otg.PatternFlowVlanTpidMetricTag
}

func newPatternFlowVlanTpidPatternFlowVlanTpidMetricTagIter(ptr *[]*otg.PatternFlowVlanTpidMetricTag) PatternFlowVlanTpidPatternFlowVlanTpidMetricTagIter {
	return &patternFlowVlanTpidPatternFlowVlanTpidMetricTagIter{fieldPtr: ptr}
}

type PatternFlowVlanTpidPatternFlowVlanTpidMetricTagIter interface {
	setMsg(*patternFlowVlanTpid) PatternFlowVlanTpidPatternFlowVlanTpidMetricTagIter
	Items() []PatternFlowVlanTpidMetricTag
	Add() PatternFlowVlanTpidMetricTag
	Append(items ...PatternFlowVlanTpidMetricTag) PatternFlowVlanTpidPatternFlowVlanTpidMetricTagIter
	Set(index int, newObj PatternFlowVlanTpidMetricTag) PatternFlowVlanTpidPatternFlowVlanTpidMetricTagIter
	Clear() PatternFlowVlanTpidPatternFlowVlanTpidMetricTagIter
	clearHolderSlice() PatternFlowVlanTpidPatternFlowVlanTpidMetricTagIter
	appendHolderSlice(item PatternFlowVlanTpidMetricTag) PatternFlowVlanTpidPatternFlowVlanTpidMetricTagIter
}

func (obj *patternFlowVlanTpidPatternFlowVlanTpidMetricTagIter) setMsg(msg *patternFlowVlanTpid) PatternFlowVlanTpidPatternFlowVlanTpidMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowVlanTpidMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowVlanTpidPatternFlowVlanTpidMetricTagIter) Items() []PatternFlowVlanTpidMetricTag {
	return obj.patternFlowVlanTpidMetricTagSlice
}

func (obj *patternFlowVlanTpidPatternFlowVlanTpidMetricTagIter) Add() PatternFlowVlanTpidMetricTag {
	newObj := &otg.PatternFlowVlanTpidMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowVlanTpidMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowVlanTpidMetricTagSlice = append(obj.patternFlowVlanTpidMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowVlanTpidPatternFlowVlanTpidMetricTagIter) Append(items ...PatternFlowVlanTpidMetricTag) PatternFlowVlanTpidPatternFlowVlanTpidMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowVlanTpidMetricTagSlice = append(obj.patternFlowVlanTpidMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowVlanTpidPatternFlowVlanTpidMetricTagIter) Set(index int, newObj PatternFlowVlanTpidMetricTag) PatternFlowVlanTpidPatternFlowVlanTpidMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowVlanTpidMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowVlanTpidPatternFlowVlanTpidMetricTagIter) Clear() PatternFlowVlanTpidPatternFlowVlanTpidMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowVlanTpidMetricTag{}
		obj.patternFlowVlanTpidMetricTagSlice = []PatternFlowVlanTpidMetricTag{}
	}
	return obj
}
func (obj *patternFlowVlanTpidPatternFlowVlanTpidMetricTagIter) clearHolderSlice() PatternFlowVlanTpidPatternFlowVlanTpidMetricTagIter {
	if len(obj.patternFlowVlanTpidMetricTagSlice) > 0 {
		obj.patternFlowVlanTpidMetricTagSlice = []PatternFlowVlanTpidMetricTag{}
	}
	return obj
}
func (obj *patternFlowVlanTpidPatternFlowVlanTpidMetricTagIter) appendHolderSlice(item PatternFlowVlanTpidMetricTag) PatternFlowVlanTpidPatternFlowVlanTpidMetricTagIter {
	obj.patternFlowVlanTpidMetricTagSlice = append(obj.patternFlowVlanTpidMetricTagSlice, item)
	return obj
}

func (obj *patternFlowVlanTpid) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVlanTpid.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowVlanTpid.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowVlanTpidMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowVlanTpid) setDefault() {
	var choices_set int = 0
	var choice PatternFlowVlanTpidChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowVlanTpidChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowVlanTpidChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowVlanTpidChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowVlanTpidChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowVlanTpidChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowVlanTpid")
			}
		} else {
			intVal := otg.PatternFlowVlanTpid_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowVlanTpid_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
