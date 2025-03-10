package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpExtensionNextExtensionHeader *****
type patternFlowGtpExtensionNextExtensionHeader struct {
	validation
	obj              *otg.PatternFlowGtpExtensionNextExtensionHeader
	marshaller       marshalPatternFlowGtpExtensionNextExtensionHeader
	unMarshaller     unMarshalPatternFlowGtpExtensionNextExtensionHeader
	incrementHolder  PatternFlowGtpExtensionNextExtensionHeaderCounter
	decrementHolder  PatternFlowGtpExtensionNextExtensionHeaderCounter
	metricTagsHolder PatternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter
}

func NewPatternFlowGtpExtensionNextExtensionHeader() PatternFlowGtpExtensionNextExtensionHeader {
	obj := patternFlowGtpExtensionNextExtensionHeader{obj: &otg.PatternFlowGtpExtensionNextExtensionHeader{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpExtensionNextExtensionHeader) msg() *otg.PatternFlowGtpExtensionNextExtensionHeader {
	return obj.obj
}

func (obj *patternFlowGtpExtensionNextExtensionHeader) setMsg(msg *otg.PatternFlowGtpExtensionNextExtensionHeader) PatternFlowGtpExtensionNextExtensionHeader {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpExtensionNextExtensionHeader struct {
	obj *patternFlowGtpExtensionNextExtensionHeader
}

type marshalPatternFlowGtpExtensionNextExtensionHeader interface {
	// ToProto marshals PatternFlowGtpExtensionNextExtensionHeader to protobuf object *otg.PatternFlowGtpExtensionNextExtensionHeader
	ToProto() (*otg.PatternFlowGtpExtensionNextExtensionHeader, error)
	// ToPbText marshals PatternFlowGtpExtensionNextExtensionHeader to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpExtensionNextExtensionHeader to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpExtensionNextExtensionHeader to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowGtpExtensionNextExtensionHeader to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowGtpExtensionNextExtensionHeader struct {
	obj *patternFlowGtpExtensionNextExtensionHeader
}

type unMarshalPatternFlowGtpExtensionNextExtensionHeader interface {
	// FromProto unmarshals PatternFlowGtpExtensionNextExtensionHeader from protobuf object *otg.PatternFlowGtpExtensionNextExtensionHeader
	FromProto(msg *otg.PatternFlowGtpExtensionNextExtensionHeader) (PatternFlowGtpExtensionNextExtensionHeader, error)
	// FromPbText unmarshals PatternFlowGtpExtensionNextExtensionHeader from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpExtensionNextExtensionHeader from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpExtensionNextExtensionHeader from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpExtensionNextExtensionHeader) Marshal() marshalPatternFlowGtpExtensionNextExtensionHeader {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpExtensionNextExtensionHeader{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpExtensionNextExtensionHeader) Unmarshal() unMarshalPatternFlowGtpExtensionNextExtensionHeader {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpExtensionNextExtensionHeader{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpExtensionNextExtensionHeader) ToProto() (*otg.PatternFlowGtpExtensionNextExtensionHeader, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpExtensionNextExtensionHeader) FromProto(msg *otg.PatternFlowGtpExtensionNextExtensionHeader) (PatternFlowGtpExtensionNextExtensionHeader, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpExtensionNextExtensionHeader) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpExtensionNextExtensionHeader) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpExtensionNextExtensionHeader) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpExtensionNextExtensionHeader) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpExtensionNextExtensionHeader) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowGtpExtensionNextExtensionHeader) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpExtensionNextExtensionHeader) FromJson(value string) error {
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

func (obj *patternFlowGtpExtensionNextExtensionHeader) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpExtensionNextExtensionHeader) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpExtensionNextExtensionHeader) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpExtensionNextExtensionHeader) Clone() (PatternFlowGtpExtensionNextExtensionHeader, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpExtensionNextExtensionHeader()
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

func (obj *patternFlowGtpExtensionNextExtensionHeader) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGtpExtensionNextExtensionHeader is it states the type of the next extension, or 0 if no next  extension exists.  This permits chaining several next extension headers.
type PatternFlowGtpExtensionNextExtensionHeader interface {
	Validation
	// msg marshals PatternFlowGtpExtensionNextExtensionHeader to protobuf object *otg.PatternFlowGtpExtensionNextExtensionHeader
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpExtensionNextExtensionHeader
	// setMsg unmarshals PatternFlowGtpExtensionNextExtensionHeader from protobuf object *otg.PatternFlowGtpExtensionNextExtensionHeader
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpExtensionNextExtensionHeader) PatternFlowGtpExtensionNextExtensionHeader
	// provides marshal interface
	Marshal() marshalPatternFlowGtpExtensionNextExtensionHeader
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpExtensionNextExtensionHeader
	// validate validates PatternFlowGtpExtensionNextExtensionHeader
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpExtensionNextExtensionHeader, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGtpExtensionNextExtensionHeaderChoiceEnum, set in PatternFlowGtpExtensionNextExtensionHeader
	Choice() PatternFlowGtpExtensionNextExtensionHeaderChoiceEnum
	// setChoice assigns PatternFlowGtpExtensionNextExtensionHeaderChoiceEnum provided by user to PatternFlowGtpExtensionNextExtensionHeader
	setChoice(value PatternFlowGtpExtensionNextExtensionHeaderChoiceEnum) PatternFlowGtpExtensionNextExtensionHeader
	// HasChoice checks if Choice has been set in PatternFlowGtpExtensionNextExtensionHeader
	HasChoice() bool
	// Value returns uint32, set in PatternFlowGtpExtensionNextExtensionHeader.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowGtpExtensionNextExtensionHeader
	SetValue(value uint32) PatternFlowGtpExtensionNextExtensionHeader
	// HasValue checks if Value has been set in PatternFlowGtpExtensionNextExtensionHeader
	HasValue() bool
	// Values returns []uint32, set in PatternFlowGtpExtensionNextExtensionHeader.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowGtpExtensionNextExtensionHeader
	SetValues(value []uint32) PatternFlowGtpExtensionNextExtensionHeader
	// Increment returns PatternFlowGtpExtensionNextExtensionHeaderCounter, set in PatternFlowGtpExtensionNextExtensionHeader.
	// PatternFlowGtpExtensionNextExtensionHeaderCounter is integer counter pattern
	Increment() PatternFlowGtpExtensionNextExtensionHeaderCounter
	// SetIncrement assigns PatternFlowGtpExtensionNextExtensionHeaderCounter provided by user to PatternFlowGtpExtensionNextExtensionHeader.
	// PatternFlowGtpExtensionNextExtensionHeaderCounter is integer counter pattern
	SetIncrement(value PatternFlowGtpExtensionNextExtensionHeaderCounter) PatternFlowGtpExtensionNextExtensionHeader
	// HasIncrement checks if Increment has been set in PatternFlowGtpExtensionNextExtensionHeader
	HasIncrement() bool
	// Decrement returns PatternFlowGtpExtensionNextExtensionHeaderCounter, set in PatternFlowGtpExtensionNextExtensionHeader.
	// PatternFlowGtpExtensionNextExtensionHeaderCounter is integer counter pattern
	Decrement() PatternFlowGtpExtensionNextExtensionHeaderCounter
	// SetDecrement assigns PatternFlowGtpExtensionNextExtensionHeaderCounter provided by user to PatternFlowGtpExtensionNextExtensionHeader.
	// PatternFlowGtpExtensionNextExtensionHeaderCounter is integer counter pattern
	SetDecrement(value PatternFlowGtpExtensionNextExtensionHeaderCounter) PatternFlowGtpExtensionNextExtensionHeader
	// HasDecrement checks if Decrement has been set in PatternFlowGtpExtensionNextExtensionHeader
	HasDecrement() bool
	// MetricTags returns PatternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIterIter, set in PatternFlowGtpExtensionNextExtensionHeader
	MetricTags() PatternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter
	setNil()
}

type PatternFlowGtpExtensionNextExtensionHeaderChoiceEnum string

// Enum of Choice on PatternFlowGtpExtensionNextExtensionHeader
var PatternFlowGtpExtensionNextExtensionHeaderChoice = struct {
	VALUE     PatternFlowGtpExtensionNextExtensionHeaderChoiceEnum
	VALUES    PatternFlowGtpExtensionNextExtensionHeaderChoiceEnum
	INCREMENT PatternFlowGtpExtensionNextExtensionHeaderChoiceEnum
	DECREMENT PatternFlowGtpExtensionNextExtensionHeaderChoiceEnum
}{
	VALUE:     PatternFlowGtpExtensionNextExtensionHeaderChoiceEnum("value"),
	VALUES:    PatternFlowGtpExtensionNextExtensionHeaderChoiceEnum("values"),
	INCREMENT: PatternFlowGtpExtensionNextExtensionHeaderChoiceEnum("increment"),
	DECREMENT: PatternFlowGtpExtensionNextExtensionHeaderChoiceEnum("decrement"),
}

func (obj *patternFlowGtpExtensionNextExtensionHeader) Choice() PatternFlowGtpExtensionNextExtensionHeaderChoiceEnum {
	return PatternFlowGtpExtensionNextExtensionHeaderChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGtpExtensionNextExtensionHeader) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGtpExtensionNextExtensionHeader) setChoice(value PatternFlowGtpExtensionNextExtensionHeaderChoiceEnum) PatternFlowGtpExtensionNextExtensionHeader {
	intValue, ok := otg.PatternFlowGtpExtensionNextExtensionHeader_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGtpExtensionNextExtensionHeaderChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGtpExtensionNextExtensionHeader_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGtpExtensionNextExtensionHeaderChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGtpExtensionNextExtensionHeaderChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGtpExtensionNextExtensionHeaderChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGtpExtensionNextExtensionHeaderCounter().msg()
	}

	if value == PatternFlowGtpExtensionNextExtensionHeaderChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGtpExtensionNextExtensionHeaderCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpExtensionNextExtensionHeader) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGtpExtensionNextExtensionHeaderChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpExtensionNextExtensionHeader) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowGtpExtensionNextExtensionHeader object
func (obj *patternFlowGtpExtensionNextExtensionHeader) SetValue(value uint32) PatternFlowGtpExtensionNextExtensionHeader {
	obj.setChoice(PatternFlowGtpExtensionNextExtensionHeaderChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowGtpExtensionNextExtensionHeader) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowGtpExtensionNextExtensionHeader object
func (obj *patternFlowGtpExtensionNextExtensionHeader) SetValues(value []uint32) PatternFlowGtpExtensionNextExtensionHeader {
	obj.setChoice(PatternFlowGtpExtensionNextExtensionHeaderChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGtpExtensionNextExtensionHeaderCounter
func (obj *patternFlowGtpExtensionNextExtensionHeader) Increment() PatternFlowGtpExtensionNextExtensionHeaderCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGtpExtensionNextExtensionHeaderChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGtpExtensionNextExtensionHeaderCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGtpExtensionNextExtensionHeaderCounter
func (obj *patternFlowGtpExtensionNextExtensionHeader) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGtpExtensionNextExtensionHeaderCounter value in the PatternFlowGtpExtensionNextExtensionHeader object
func (obj *patternFlowGtpExtensionNextExtensionHeader) SetIncrement(value PatternFlowGtpExtensionNextExtensionHeaderCounter) PatternFlowGtpExtensionNextExtensionHeader {
	obj.setChoice(PatternFlowGtpExtensionNextExtensionHeaderChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGtpExtensionNextExtensionHeaderCounter
func (obj *patternFlowGtpExtensionNextExtensionHeader) Decrement() PatternFlowGtpExtensionNextExtensionHeaderCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGtpExtensionNextExtensionHeaderChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGtpExtensionNextExtensionHeaderCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGtpExtensionNextExtensionHeaderCounter
func (obj *patternFlowGtpExtensionNextExtensionHeader) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGtpExtensionNextExtensionHeaderCounter value in the PatternFlowGtpExtensionNextExtensionHeader object
func (obj *patternFlowGtpExtensionNextExtensionHeader) SetDecrement(value PatternFlowGtpExtensionNextExtensionHeaderCounter) PatternFlowGtpExtensionNextExtensionHeader {
	obj.setChoice(PatternFlowGtpExtensionNextExtensionHeaderChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGtpExtensionNextExtensionHeaderMetricTag
func (obj *patternFlowGtpExtensionNextExtensionHeader) MetricTags() PatternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGtpExtensionNextExtensionHeaderMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter struct {
	obj                                                      *patternFlowGtpExtensionNextExtensionHeader
	patternFlowGtpExtensionNextExtensionHeaderMetricTagSlice []PatternFlowGtpExtensionNextExtensionHeaderMetricTag
	fieldPtr                                                 *[]*otg.PatternFlowGtpExtensionNextExtensionHeaderMetricTag
}

func newPatternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter(ptr *[]*otg.PatternFlowGtpExtensionNextExtensionHeaderMetricTag) PatternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter {
	return &patternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter{fieldPtr: ptr}
}

type PatternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter interface {
	setMsg(*patternFlowGtpExtensionNextExtensionHeader) PatternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter
	Items() []PatternFlowGtpExtensionNextExtensionHeaderMetricTag
	Add() PatternFlowGtpExtensionNextExtensionHeaderMetricTag
	Append(items ...PatternFlowGtpExtensionNextExtensionHeaderMetricTag) PatternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter
	Set(index int, newObj PatternFlowGtpExtensionNextExtensionHeaderMetricTag) PatternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter
	Clear() PatternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter
	clearHolderSlice() PatternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter
	appendHolderSlice(item PatternFlowGtpExtensionNextExtensionHeaderMetricTag) PatternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter) setMsg(msg *patternFlowGtpExtensionNextExtensionHeader) PatternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGtpExtensionNextExtensionHeaderMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter) Items() []PatternFlowGtpExtensionNextExtensionHeaderMetricTag {
	return obj.patternFlowGtpExtensionNextExtensionHeaderMetricTagSlice
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter) Add() PatternFlowGtpExtensionNextExtensionHeaderMetricTag {
	newObj := &otg.PatternFlowGtpExtensionNextExtensionHeaderMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGtpExtensionNextExtensionHeaderMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGtpExtensionNextExtensionHeaderMetricTagSlice = append(obj.patternFlowGtpExtensionNextExtensionHeaderMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter) Append(items ...PatternFlowGtpExtensionNextExtensionHeaderMetricTag) PatternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGtpExtensionNextExtensionHeaderMetricTagSlice = append(obj.patternFlowGtpExtensionNextExtensionHeaderMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter) Set(index int, newObj PatternFlowGtpExtensionNextExtensionHeaderMetricTag) PatternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGtpExtensionNextExtensionHeaderMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter) Clear() PatternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGtpExtensionNextExtensionHeaderMetricTag{}
		obj.patternFlowGtpExtensionNextExtensionHeaderMetricTagSlice = []PatternFlowGtpExtensionNextExtensionHeaderMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter) clearHolderSlice() PatternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter {
	if len(obj.patternFlowGtpExtensionNextExtensionHeaderMetricTagSlice) > 0 {
		obj.patternFlowGtpExtensionNextExtensionHeaderMetricTagSlice = []PatternFlowGtpExtensionNextExtensionHeaderMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter) appendHolderSlice(item PatternFlowGtpExtensionNextExtensionHeaderMetricTag) PatternFlowGtpExtensionNextExtensionHeaderPatternFlowGtpExtensionNextExtensionHeaderMetricTagIter {
	obj.patternFlowGtpExtensionNextExtensionHeaderMetricTagSlice = append(obj.patternFlowGtpExtensionNextExtensionHeaderMetricTagSlice, item)
	return obj
}

func (obj *patternFlowGtpExtensionNextExtensionHeader) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpExtensionNextExtensionHeader.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowGtpExtensionNextExtensionHeader.Values <= 255 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowGtpExtensionNextExtensionHeaderMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGtpExtensionNextExtensionHeader) setDefault() {
	var choices_set int = 0
	var choice PatternFlowGtpExtensionNextExtensionHeaderChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowGtpExtensionNextExtensionHeaderChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowGtpExtensionNextExtensionHeaderChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowGtpExtensionNextExtensionHeaderChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowGtpExtensionNextExtensionHeaderChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowGtpExtensionNextExtensionHeaderChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowGtpExtensionNextExtensionHeader")
			}
		} else {
			intVal := otg.PatternFlowGtpExtensionNextExtensionHeader_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowGtpExtensionNextExtensionHeader_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
