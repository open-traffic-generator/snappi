package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpWindow *****
type patternFlowTcpWindow struct {
	validation
	obj              *otg.PatternFlowTcpWindow
	marshaller       marshalPatternFlowTcpWindow
	unMarshaller     unMarshalPatternFlowTcpWindow
	incrementHolder  PatternFlowTcpWindowCounter
	decrementHolder  PatternFlowTcpWindowCounter
	metricTagsHolder PatternFlowTcpWindowPatternFlowTcpWindowMetricTagIter
}

func NewPatternFlowTcpWindow() PatternFlowTcpWindow {
	obj := patternFlowTcpWindow{obj: &otg.PatternFlowTcpWindow{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpWindow) msg() *otg.PatternFlowTcpWindow {
	return obj.obj
}

func (obj *patternFlowTcpWindow) setMsg(msg *otg.PatternFlowTcpWindow) PatternFlowTcpWindow {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpWindow struct {
	obj *patternFlowTcpWindow
}

type marshalPatternFlowTcpWindow interface {
	// ToProto marshals PatternFlowTcpWindow to protobuf object *otg.PatternFlowTcpWindow
	ToProto() (*otg.PatternFlowTcpWindow, error)
	// ToPbText marshals PatternFlowTcpWindow to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpWindow to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpWindow to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpWindow struct {
	obj *patternFlowTcpWindow
}

type unMarshalPatternFlowTcpWindow interface {
	// FromProto unmarshals PatternFlowTcpWindow from protobuf object *otg.PatternFlowTcpWindow
	FromProto(msg *otg.PatternFlowTcpWindow) (PatternFlowTcpWindow, error)
	// FromPbText unmarshals PatternFlowTcpWindow from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpWindow from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpWindow from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpWindow) Marshal() marshalPatternFlowTcpWindow {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpWindow{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpWindow) Unmarshal() unMarshalPatternFlowTcpWindow {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpWindow{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpWindow) ToProto() (*otg.PatternFlowTcpWindow, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpWindow) FromProto(msg *otg.PatternFlowTcpWindow) (PatternFlowTcpWindow, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpWindow) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpWindow) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpWindow) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpWindow) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpWindow) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpWindow) FromJson(value string) error {
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

func (obj *patternFlowTcpWindow) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpWindow) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpWindow) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpWindow) Clone() (PatternFlowTcpWindow, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpWindow()
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

func (obj *patternFlowTcpWindow) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowTcpWindow is tcp connection window.
type PatternFlowTcpWindow interface {
	Validation
	// msg marshals PatternFlowTcpWindow to protobuf object *otg.PatternFlowTcpWindow
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpWindow
	// setMsg unmarshals PatternFlowTcpWindow from protobuf object *otg.PatternFlowTcpWindow
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpWindow) PatternFlowTcpWindow
	// provides marshal interface
	Marshal() marshalPatternFlowTcpWindow
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpWindow
	// validate validates PatternFlowTcpWindow
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpWindow, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowTcpWindowChoiceEnum, set in PatternFlowTcpWindow
	Choice() PatternFlowTcpWindowChoiceEnum
	// setChoice assigns PatternFlowTcpWindowChoiceEnum provided by user to PatternFlowTcpWindow
	setChoice(value PatternFlowTcpWindowChoiceEnum) PatternFlowTcpWindow
	// HasChoice checks if Choice has been set in PatternFlowTcpWindow
	HasChoice() bool
	// Value returns uint32, set in PatternFlowTcpWindow.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowTcpWindow
	SetValue(value uint32) PatternFlowTcpWindow
	// HasValue checks if Value has been set in PatternFlowTcpWindow
	HasValue() bool
	// Values returns []uint32, set in PatternFlowTcpWindow.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowTcpWindow
	SetValues(value []uint32) PatternFlowTcpWindow
	// Increment returns PatternFlowTcpWindowCounter, set in PatternFlowTcpWindow.
	// PatternFlowTcpWindowCounter is integer counter pattern
	Increment() PatternFlowTcpWindowCounter
	// SetIncrement assigns PatternFlowTcpWindowCounter provided by user to PatternFlowTcpWindow.
	// PatternFlowTcpWindowCounter is integer counter pattern
	SetIncrement(value PatternFlowTcpWindowCounter) PatternFlowTcpWindow
	// HasIncrement checks if Increment has been set in PatternFlowTcpWindow
	HasIncrement() bool
	// Decrement returns PatternFlowTcpWindowCounter, set in PatternFlowTcpWindow.
	// PatternFlowTcpWindowCounter is integer counter pattern
	Decrement() PatternFlowTcpWindowCounter
	// SetDecrement assigns PatternFlowTcpWindowCounter provided by user to PatternFlowTcpWindow.
	// PatternFlowTcpWindowCounter is integer counter pattern
	SetDecrement(value PatternFlowTcpWindowCounter) PatternFlowTcpWindow
	// HasDecrement checks if Decrement has been set in PatternFlowTcpWindow
	HasDecrement() bool
	// MetricTags returns PatternFlowTcpWindowPatternFlowTcpWindowMetricTagIterIter, set in PatternFlowTcpWindow
	MetricTags() PatternFlowTcpWindowPatternFlowTcpWindowMetricTagIter
	setNil()
}

type PatternFlowTcpWindowChoiceEnum string

// Enum of Choice on PatternFlowTcpWindow
var PatternFlowTcpWindowChoice = struct {
	VALUE     PatternFlowTcpWindowChoiceEnum
	VALUES    PatternFlowTcpWindowChoiceEnum
	INCREMENT PatternFlowTcpWindowChoiceEnum
	DECREMENT PatternFlowTcpWindowChoiceEnum
}{
	VALUE:     PatternFlowTcpWindowChoiceEnum("value"),
	VALUES:    PatternFlowTcpWindowChoiceEnum("values"),
	INCREMENT: PatternFlowTcpWindowChoiceEnum("increment"),
	DECREMENT: PatternFlowTcpWindowChoiceEnum("decrement"),
}

func (obj *patternFlowTcpWindow) Choice() PatternFlowTcpWindowChoiceEnum {
	return PatternFlowTcpWindowChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowTcpWindow) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowTcpWindow) setChoice(value PatternFlowTcpWindowChoiceEnum) PatternFlowTcpWindow {
	intValue, ok := otg.PatternFlowTcpWindow_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowTcpWindowChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowTcpWindow_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowTcpWindowChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowTcpWindowChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowTcpWindowChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowTcpWindowCounter().msg()
	}

	if value == PatternFlowTcpWindowChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowTcpWindowCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpWindow) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowTcpWindowChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpWindow) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowTcpWindow object
func (obj *patternFlowTcpWindow) SetValue(value uint32) PatternFlowTcpWindow {
	obj.setChoice(PatternFlowTcpWindowChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowTcpWindow) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowTcpWindow object
func (obj *patternFlowTcpWindow) SetValues(value []uint32) PatternFlowTcpWindow {
	obj.setChoice(PatternFlowTcpWindowChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowTcpWindowCounter
func (obj *patternFlowTcpWindow) Increment() PatternFlowTcpWindowCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowTcpWindowChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowTcpWindowCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowTcpWindowCounter
func (obj *patternFlowTcpWindow) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowTcpWindowCounter value in the PatternFlowTcpWindow object
func (obj *patternFlowTcpWindow) SetIncrement(value PatternFlowTcpWindowCounter) PatternFlowTcpWindow {
	obj.setChoice(PatternFlowTcpWindowChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowTcpWindowCounter
func (obj *patternFlowTcpWindow) Decrement() PatternFlowTcpWindowCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowTcpWindowChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowTcpWindowCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowTcpWindowCounter
func (obj *patternFlowTcpWindow) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowTcpWindowCounter value in the PatternFlowTcpWindow object
func (obj *patternFlowTcpWindow) SetDecrement(value PatternFlowTcpWindowCounter) PatternFlowTcpWindow {
	obj.setChoice(PatternFlowTcpWindowChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowTcpWindowMetricTag
func (obj *patternFlowTcpWindow) MetricTags() PatternFlowTcpWindowPatternFlowTcpWindowMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowTcpWindowMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowTcpWindowPatternFlowTcpWindowMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowTcpWindowPatternFlowTcpWindowMetricTagIter struct {
	obj                                *patternFlowTcpWindow
	patternFlowTcpWindowMetricTagSlice []PatternFlowTcpWindowMetricTag
	fieldPtr                           *[]*otg.PatternFlowTcpWindowMetricTag
}

func newPatternFlowTcpWindowPatternFlowTcpWindowMetricTagIter(ptr *[]*otg.PatternFlowTcpWindowMetricTag) PatternFlowTcpWindowPatternFlowTcpWindowMetricTagIter {
	return &patternFlowTcpWindowPatternFlowTcpWindowMetricTagIter{fieldPtr: ptr}
}

type PatternFlowTcpWindowPatternFlowTcpWindowMetricTagIter interface {
	setMsg(*patternFlowTcpWindow) PatternFlowTcpWindowPatternFlowTcpWindowMetricTagIter
	Items() []PatternFlowTcpWindowMetricTag
	Add() PatternFlowTcpWindowMetricTag
	Append(items ...PatternFlowTcpWindowMetricTag) PatternFlowTcpWindowPatternFlowTcpWindowMetricTagIter
	Set(index int, newObj PatternFlowTcpWindowMetricTag) PatternFlowTcpWindowPatternFlowTcpWindowMetricTagIter
	Clear() PatternFlowTcpWindowPatternFlowTcpWindowMetricTagIter
	clearHolderSlice() PatternFlowTcpWindowPatternFlowTcpWindowMetricTagIter
	appendHolderSlice(item PatternFlowTcpWindowMetricTag) PatternFlowTcpWindowPatternFlowTcpWindowMetricTagIter
}

func (obj *patternFlowTcpWindowPatternFlowTcpWindowMetricTagIter) setMsg(msg *patternFlowTcpWindow) PatternFlowTcpWindowPatternFlowTcpWindowMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowTcpWindowMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowTcpWindowPatternFlowTcpWindowMetricTagIter) Items() []PatternFlowTcpWindowMetricTag {
	return obj.patternFlowTcpWindowMetricTagSlice
}

func (obj *patternFlowTcpWindowPatternFlowTcpWindowMetricTagIter) Add() PatternFlowTcpWindowMetricTag {
	newObj := &otg.PatternFlowTcpWindowMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowTcpWindowMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowTcpWindowMetricTagSlice = append(obj.patternFlowTcpWindowMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowTcpWindowPatternFlowTcpWindowMetricTagIter) Append(items ...PatternFlowTcpWindowMetricTag) PatternFlowTcpWindowPatternFlowTcpWindowMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowTcpWindowMetricTagSlice = append(obj.patternFlowTcpWindowMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowTcpWindowPatternFlowTcpWindowMetricTagIter) Set(index int, newObj PatternFlowTcpWindowMetricTag) PatternFlowTcpWindowPatternFlowTcpWindowMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowTcpWindowMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowTcpWindowPatternFlowTcpWindowMetricTagIter) Clear() PatternFlowTcpWindowPatternFlowTcpWindowMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowTcpWindowMetricTag{}
		obj.patternFlowTcpWindowMetricTagSlice = []PatternFlowTcpWindowMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpWindowPatternFlowTcpWindowMetricTagIter) clearHolderSlice() PatternFlowTcpWindowPatternFlowTcpWindowMetricTagIter {
	if len(obj.patternFlowTcpWindowMetricTagSlice) > 0 {
		obj.patternFlowTcpWindowMetricTagSlice = []PatternFlowTcpWindowMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpWindowPatternFlowTcpWindowMetricTagIter) appendHolderSlice(item PatternFlowTcpWindowMetricTag) PatternFlowTcpWindowPatternFlowTcpWindowMetricTagIter {
	obj.patternFlowTcpWindowMetricTagSlice = append(obj.patternFlowTcpWindowMetricTagSlice, item)
	return obj
}

func (obj *patternFlowTcpWindow) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpWindow.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowTcpWindow.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowTcpWindowMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowTcpWindow) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowTcpWindowChoice.VALUE)

	}

}
