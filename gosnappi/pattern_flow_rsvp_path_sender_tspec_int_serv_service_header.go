package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTspecIntServServiceHeader *****
type patternFlowRSVPPathSenderTspecIntServServiceHeader struct {
	validation
	obj             *otg.PatternFlowRSVPPathSenderTspecIntServServiceHeader
	marshaller      marshalPatternFlowRSVPPathSenderTspecIntServServiceHeader
	unMarshaller    unMarshalPatternFlowRSVPPathSenderTspecIntServServiceHeader
	incrementHolder PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
	decrementHolder PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
}

func NewPatternFlowRSVPPathSenderTspecIntServServiceHeader() PatternFlowRSVPPathSenderTspecIntServServiceHeader {
	obj := patternFlowRSVPPathSenderTspecIntServServiceHeader{obj: &otg.PatternFlowRSVPPathSenderTspecIntServServiceHeader{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeader) msg() *otg.PatternFlowRSVPPathSenderTspecIntServServiceHeader {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeader) setMsg(msg *otg.PatternFlowRSVPPathSenderTspecIntServServiceHeader) PatternFlowRSVPPathSenderTspecIntServServiceHeader {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTspecIntServServiceHeader struct {
	obj *patternFlowRSVPPathSenderTspecIntServServiceHeader
}

type marshalPatternFlowRSVPPathSenderTspecIntServServiceHeader interface {
	// ToProto marshals PatternFlowRSVPPathSenderTspecIntServServiceHeader to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServServiceHeader
	ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServServiceHeader, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTspecIntServServiceHeader to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTspecIntServServiceHeader to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTspecIntServServiceHeader to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowRSVPPathSenderTspecIntServServiceHeader to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTspecIntServServiceHeader struct {
	obj *patternFlowRSVPPathSenderTspecIntServServiceHeader
}

type unMarshalPatternFlowRSVPPathSenderTspecIntServServiceHeader interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTspecIntServServiceHeader from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServServiceHeader
	FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServServiceHeader) (PatternFlowRSVPPathSenderTspecIntServServiceHeader, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTspecIntServServiceHeader from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTspecIntServServiceHeader from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTspecIntServServiceHeader from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeader) Marshal() marshalPatternFlowRSVPPathSenderTspecIntServServiceHeader {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTspecIntServServiceHeader{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeader) Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServServiceHeader {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTspecIntServServiceHeader{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServServiceHeader) ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServServiceHeader, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServServiceHeader) FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServServiceHeader) (PatternFlowRSVPPathSenderTspecIntServServiceHeader, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServServiceHeader) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServServiceHeader) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServServiceHeader) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServServiceHeader) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServServiceHeader) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServServiceHeader) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServServiceHeader) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeader) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeader) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeader) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeader) Clone() (PatternFlowRSVPPathSenderTspecIntServServiceHeader, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTspecIntServServiceHeader()
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

func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeader) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathSenderTspecIntServServiceHeader is service header, service number - '1' (Generic information) if in a PATH message.
type PatternFlowRSVPPathSenderTspecIntServServiceHeader interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTspecIntServServiceHeader to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServServiceHeader
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTspecIntServServiceHeader
	// setMsg unmarshals PatternFlowRSVPPathSenderTspecIntServServiceHeader from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServServiceHeader
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTspecIntServServiceHeader) PatternFlowRSVPPathSenderTspecIntServServiceHeader
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTspecIntServServiceHeader
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServServiceHeader
	// validate validates PatternFlowRSVPPathSenderTspecIntServServiceHeader
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTspecIntServServiceHeader, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoiceEnum, set in PatternFlowRSVPPathSenderTspecIntServServiceHeader
	Choice() PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoiceEnum
	// setChoice assigns PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoiceEnum provided by user to PatternFlowRSVPPathSenderTspecIntServServiceHeader
	setChoice(value PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoiceEnum) PatternFlowRSVPPathSenderTspecIntServServiceHeader
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathSenderTspecIntServServiceHeader
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRSVPPathSenderTspecIntServServiceHeader.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServServiceHeader
	SetValue(value uint32) PatternFlowRSVPPathSenderTspecIntServServiceHeader
	// HasValue checks if Value has been set in PatternFlowRSVPPathSenderTspecIntServServiceHeader
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRSVPPathSenderTspecIntServServiceHeader.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServServiceHeader
	SetValues(value []uint32) PatternFlowRSVPPathSenderTspecIntServServiceHeader
	// Increment returns PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter, set in PatternFlowRSVPPathSenderTspecIntServServiceHeader.
	// PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter is integer counter pattern
	Increment() PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
	// SetIncrement assigns PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter provided by user to PatternFlowRSVPPathSenderTspecIntServServiceHeader.
	// PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter is integer counter pattern
	SetIncrement(value PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) PatternFlowRSVPPathSenderTspecIntServServiceHeader
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathSenderTspecIntServServiceHeader
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter, set in PatternFlowRSVPPathSenderTspecIntServServiceHeader.
	// PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter is integer counter pattern
	Decrement() PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
	// SetDecrement assigns PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter provided by user to PatternFlowRSVPPathSenderTspecIntServServiceHeader.
	// PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter is integer counter pattern
	SetDecrement(value PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) PatternFlowRSVPPathSenderTspecIntServServiceHeader
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathSenderTspecIntServServiceHeader
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathSenderTspecIntServServiceHeader
var PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoice = struct {
	VALUE     PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoiceEnum
	VALUES    PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoiceEnum
	INCREMENT PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoiceEnum
	DECREMENT PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeader) Choice() PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoiceEnum {
	return PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeader) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeader) setChoice(value PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoiceEnum) PatternFlowRSVPPathSenderTspecIntServServiceHeader {
	intValue, ok := otg.PatternFlowRSVPPathSenderTspecIntServServiceHeader_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathSenderTspecIntServServiceHeader_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoice.VALUE {
		defaultValue := uint32(1)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoice.VALUES {
		defaultValue := []uint32{1}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter().msg()
	}

	if value == PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeader) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeader) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServServiceHeader object
func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeader) SetValue(value uint32) PatternFlowRSVPPathSenderTspecIntServServiceHeader {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeader) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{1})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRSVPPathSenderTspecIntServServiceHeader object
func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeader) SetValues(value []uint32) PatternFlowRSVPPathSenderTspecIntServServiceHeader {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeader) Increment() PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathSenderTspecIntServServiceHeaderCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeader) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter value in the PatternFlowRSVPPathSenderTspecIntServServiceHeader object
func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeader) SetIncrement(value PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) PatternFlowRSVPPathSenderTspecIntServServiceHeader {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeader) Decrement() PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathSenderTspecIntServServiceHeaderCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter
func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeader) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter value in the PatternFlowRSVPPathSenderTspecIntServServiceHeader object
func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeader) SetDecrement(value PatternFlowRSVPPathSenderTspecIntServServiceHeaderCounter) PatternFlowRSVPPathSenderTspecIntServServiceHeader {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeader) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServServiceHeader.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowRSVPPathSenderTspecIntServServiceHeader.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowRSVPPathSenderTspecIntServServiceHeader) setDefault() {
	var choices_set int = 0
	var choice PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowRSVPPathSenderTspecIntServServiceHeaderChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowRSVPPathSenderTspecIntServServiceHeader")
			}
		} else {
			intVal := otg.PatternFlowRSVPPathSenderTspecIntServServiceHeader_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowRSVPPathSenderTspecIntServServiceHeader_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
