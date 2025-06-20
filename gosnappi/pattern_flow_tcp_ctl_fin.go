package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpCtlFin *****
type patternFlowTcpCtlFin struct {
	validation
	obj              *otg.PatternFlowTcpCtlFin
	marshaller       marshalPatternFlowTcpCtlFin
	unMarshaller     unMarshalPatternFlowTcpCtlFin
	incrementHolder  PatternFlowTcpCtlFinCounter
	decrementHolder  PatternFlowTcpCtlFinCounter
	metricTagsHolder PatternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter
}

func NewPatternFlowTcpCtlFin() PatternFlowTcpCtlFin {
	obj := patternFlowTcpCtlFin{obj: &otg.PatternFlowTcpCtlFin{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpCtlFin) msg() *otg.PatternFlowTcpCtlFin {
	return obj.obj
}

func (obj *patternFlowTcpCtlFin) setMsg(msg *otg.PatternFlowTcpCtlFin) PatternFlowTcpCtlFin {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpCtlFin struct {
	obj *patternFlowTcpCtlFin
}

type marshalPatternFlowTcpCtlFin interface {
	// ToProto marshals PatternFlowTcpCtlFin to protobuf object *otg.PatternFlowTcpCtlFin
	ToProto() (*otg.PatternFlowTcpCtlFin, error)
	// ToPbText marshals PatternFlowTcpCtlFin to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpCtlFin to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpCtlFin to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpCtlFin struct {
	obj *patternFlowTcpCtlFin
}

type unMarshalPatternFlowTcpCtlFin interface {
	// FromProto unmarshals PatternFlowTcpCtlFin from protobuf object *otg.PatternFlowTcpCtlFin
	FromProto(msg *otg.PatternFlowTcpCtlFin) (PatternFlowTcpCtlFin, error)
	// FromPbText unmarshals PatternFlowTcpCtlFin from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpCtlFin from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpCtlFin from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpCtlFin) Marshal() marshalPatternFlowTcpCtlFin {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpCtlFin{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpCtlFin) Unmarshal() unMarshalPatternFlowTcpCtlFin {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpCtlFin{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpCtlFin) ToProto() (*otg.PatternFlowTcpCtlFin, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpCtlFin) FromProto(msg *otg.PatternFlowTcpCtlFin) (PatternFlowTcpCtlFin, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpCtlFin) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlFin) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpCtlFin) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlFin) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpCtlFin) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpCtlFin) FromJson(value string) error {
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

func (obj *patternFlowTcpCtlFin) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlFin) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpCtlFin) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpCtlFin) Clone() (PatternFlowTcpCtlFin, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpCtlFin()
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

func (obj *patternFlowTcpCtlFin) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowTcpCtlFin is last packet from the sender.
type PatternFlowTcpCtlFin interface {
	Validation
	// msg marshals PatternFlowTcpCtlFin to protobuf object *otg.PatternFlowTcpCtlFin
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpCtlFin
	// setMsg unmarshals PatternFlowTcpCtlFin from protobuf object *otg.PatternFlowTcpCtlFin
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpCtlFin) PatternFlowTcpCtlFin
	// provides marshal interface
	Marshal() marshalPatternFlowTcpCtlFin
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpCtlFin
	// validate validates PatternFlowTcpCtlFin
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpCtlFin, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowTcpCtlFinChoiceEnum, set in PatternFlowTcpCtlFin
	Choice() PatternFlowTcpCtlFinChoiceEnum
	// setChoice assigns PatternFlowTcpCtlFinChoiceEnum provided by user to PatternFlowTcpCtlFin
	setChoice(value PatternFlowTcpCtlFinChoiceEnum) PatternFlowTcpCtlFin
	// HasChoice checks if Choice has been set in PatternFlowTcpCtlFin
	HasChoice() bool
	// Value returns uint32, set in PatternFlowTcpCtlFin.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowTcpCtlFin
	SetValue(value uint32) PatternFlowTcpCtlFin
	// HasValue checks if Value has been set in PatternFlowTcpCtlFin
	HasValue() bool
	// Values returns []uint32, set in PatternFlowTcpCtlFin.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowTcpCtlFin
	SetValues(value []uint32) PatternFlowTcpCtlFin
	// Increment returns PatternFlowTcpCtlFinCounter, set in PatternFlowTcpCtlFin.
	// PatternFlowTcpCtlFinCounter is integer counter pattern
	Increment() PatternFlowTcpCtlFinCounter
	// SetIncrement assigns PatternFlowTcpCtlFinCounter provided by user to PatternFlowTcpCtlFin.
	// PatternFlowTcpCtlFinCounter is integer counter pattern
	SetIncrement(value PatternFlowTcpCtlFinCounter) PatternFlowTcpCtlFin
	// HasIncrement checks if Increment has been set in PatternFlowTcpCtlFin
	HasIncrement() bool
	// Decrement returns PatternFlowTcpCtlFinCounter, set in PatternFlowTcpCtlFin.
	// PatternFlowTcpCtlFinCounter is integer counter pattern
	Decrement() PatternFlowTcpCtlFinCounter
	// SetDecrement assigns PatternFlowTcpCtlFinCounter provided by user to PatternFlowTcpCtlFin.
	// PatternFlowTcpCtlFinCounter is integer counter pattern
	SetDecrement(value PatternFlowTcpCtlFinCounter) PatternFlowTcpCtlFin
	// HasDecrement checks if Decrement has been set in PatternFlowTcpCtlFin
	HasDecrement() bool
	// MetricTags returns PatternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIterIter, set in PatternFlowTcpCtlFin
	MetricTags() PatternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter
	setNil()
}

type PatternFlowTcpCtlFinChoiceEnum string

// Enum of Choice on PatternFlowTcpCtlFin
var PatternFlowTcpCtlFinChoice = struct {
	VALUE     PatternFlowTcpCtlFinChoiceEnum
	VALUES    PatternFlowTcpCtlFinChoiceEnum
	INCREMENT PatternFlowTcpCtlFinChoiceEnum
	DECREMENT PatternFlowTcpCtlFinChoiceEnum
}{
	VALUE:     PatternFlowTcpCtlFinChoiceEnum("value"),
	VALUES:    PatternFlowTcpCtlFinChoiceEnum("values"),
	INCREMENT: PatternFlowTcpCtlFinChoiceEnum("increment"),
	DECREMENT: PatternFlowTcpCtlFinChoiceEnum("decrement"),
}

func (obj *patternFlowTcpCtlFin) Choice() PatternFlowTcpCtlFinChoiceEnum {
	return PatternFlowTcpCtlFinChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowTcpCtlFin) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowTcpCtlFin) setChoice(value PatternFlowTcpCtlFinChoiceEnum) PatternFlowTcpCtlFin {
	intValue, ok := otg.PatternFlowTcpCtlFin_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowTcpCtlFinChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowTcpCtlFin_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowTcpCtlFinChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowTcpCtlFinChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowTcpCtlFinChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowTcpCtlFinCounter().msg()
	}

	if value == PatternFlowTcpCtlFinChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowTcpCtlFinCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpCtlFin) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowTcpCtlFinChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpCtlFin) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowTcpCtlFin object
func (obj *patternFlowTcpCtlFin) SetValue(value uint32) PatternFlowTcpCtlFin {
	obj.setChoice(PatternFlowTcpCtlFinChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowTcpCtlFin) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowTcpCtlFin object
func (obj *patternFlowTcpCtlFin) SetValues(value []uint32) PatternFlowTcpCtlFin {
	obj.setChoice(PatternFlowTcpCtlFinChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowTcpCtlFinCounter
func (obj *patternFlowTcpCtlFin) Increment() PatternFlowTcpCtlFinCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowTcpCtlFinChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowTcpCtlFinCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowTcpCtlFinCounter
func (obj *patternFlowTcpCtlFin) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowTcpCtlFinCounter value in the PatternFlowTcpCtlFin object
func (obj *patternFlowTcpCtlFin) SetIncrement(value PatternFlowTcpCtlFinCounter) PatternFlowTcpCtlFin {
	obj.setChoice(PatternFlowTcpCtlFinChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowTcpCtlFinCounter
func (obj *patternFlowTcpCtlFin) Decrement() PatternFlowTcpCtlFinCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowTcpCtlFinChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowTcpCtlFinCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowTcpCtlFinCounter
func (obj *patternFlowTcpCtlFin) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowTcpCtlFinCounter value in the PatternFlowTcpCtlFin object
func (obj *patternFlowTcpCtlFin) SetDecrement(value PatternFlowTcpCtlFinCounter) PatternFlowTcpCtlFin {
	obj.setChoice(PatternFlowTcpCtlFinChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowTcpCtlFinMetricTag
func (obj *patternFlowTcpCtlFin) MetricTags() PatternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowTcpCtlFinMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter struct {
	obj                                *patternFlowTcpCtlFin
	patternFlowTcpCtlFinMetricTagSlice []PatternFlowTcpCtlFinMetricTag
	fieldPtr                           *[]*otg.PatternFlowTcpCtlFinMetricTag
}

func newPatternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter(ptr *[]*otg.PatternFlowTcpCtlFinMetricTag) PatternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter {
	return &patternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter{fieldPtr: ptr}
}

type PatternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter interface {
	setMsg(*patternFlowTcpCtlFin) PatternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter
	Items() []PatternFlowTcpCtlFinMetricTag
	Add() PatternFlowTcpCtlFinMetricTag
	Append(items ...PatternFlowTcpCtlFinMetricTag) PatternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter
	Set(index int, newObj PatternFlowTcpCtlFinMetricTag) PatternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter
	Clear() PatternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter
	clearHolderSlice() PatternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter
	appendHolderSlice(item PatternFlowTcpCtlFinMetricTag) PatternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter
}

func (obj *patternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter) setMsg(msg *patternFlowTcpCtlFin) PatternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowTcpCtlFinMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter) Items() []PatternFlowTcpCtlFinMetricTag {
	return obj.patternFlowTcpCtlFinMetricTagSlice
}

func (obj *patternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter) Add() PatternFlowTcpCtlFinMetricTag {
	newObj := &otg.PatternFlowTcpCtlFinMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowTcpCtlFinMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowTcpCtlFinMetricTagSlice = append(obj.patternFlowTcpCtlFinMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter) Append(items ...PatternFlowTcpCtlFinMetricTag) PatternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowTcpCtlFinMetricTagSlice = append(obj.patternFlowTcpCtlFinMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter) Set(index int, newObj PatternFlowTcpCtlFinMetricTag) PatternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowTcpCtlFinMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter) Clear() PatternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowTcpCtlFinMetricTag{}
		obj.patternFlowTcpCtlFinMetricTagSlice = []PatternFlowTcpCtlFinMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter) clearHolderSlice() PatternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter {
	if len(obj.patternFlowTcpCtlFinMetricTagSlice) > 0 {
		obj.patternFlowTcpCtlFinMetricTagSlice = []PatternFlowTcpCtlFinMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter) appendHolderSlice(item PatternFlowTcpCtlFinMetricTag) PatternFlowTcpCtlFinPatternFlowTcpCtlFinMetricTagIter {
	obj.patternFlowTcpCtlFinMetricTagSlice = append(obj.patternFlowTcpCtlFinMetricTagSlice, item)
	return obj
}

func (obj *patternFlowTcpCtlFin) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpCtlFin.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowTcpCtlFin.Values <= 1 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowTcpCtlFinMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowTcpCtlFin) setDefault() {
	var choices_set int = 0
	var choice PatternFlowTcpCtlFinChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowTcpCtlFinChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowTcpCtlFinChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowTcpCtlFinChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowTcpCtlFinChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowTcpCtlFinChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowTcpCtlFin")
			}
		} else {
			intVal := otg.PatternFlowTcpCtlFin_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowTcpCtlFin_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
