package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize *****
type patternFlowRSVPPathSenderTspecIntServMaximumPacketSize struct {
	validation
	obj             *otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
	marshaller      marshalPatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
	unMarshaller    unMarshalPatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
	incrementHolder PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
	decrementHolder PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
}

func NewPatternFlowRSVPPathSenderTspecIntServMaximumPacketSize() PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize {
	obj := patternFlowRSVPPathSenderTspecIntServMaximumPacketSize{obj: &otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSize) msg() *otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSize) setMsg(msg *otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize) PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTspecIntServMaximumPacketSize struct {
	obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSize
}

type marshalPatternFlowRSVPPathSenderTspecIntServMaximumPacketSize interface {
	// ToProto marshals PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
	ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTspecIntServMaximumPacketSize struct {
	obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSize
}

type unMarshalPatternFlowRSVPPathSenderTspecIntServMaximumPacketSize interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
	FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize) (PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSize) Marshal() marshalPatternFlowRSVPPathSenderTspecIntServMaximumPacketSize {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTspecIntServMaximumPacketSize{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSize) Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServMaximumPacketSize {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTspecIntServMaximumPacketSize{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServMaximumPacketSize) ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServMaximumPacketSize) FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize) (PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServMaximumPacketSize) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServMaximumPacketSize) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServMaximumPacketSize) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServMaximumPacketSize) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServMaximumPacketSize) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServMaximumPacketSize) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServMaximumPacketSize) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSize) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSize) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSize) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSize) Clone() (PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTspecIntServMaximumPacketSize()
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

func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSize) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize is the maximum packet size parameter should be set to the size of the largest packet the application might wish to generate. This value must, by definition, be equal to or larger than the value of The minimum policed unit.
type PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
	// setMsg unmarshals PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize) PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
	// validate validates PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoiceEnum, set in PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
	Choice() PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoiceEnum
	// setChoice assigns PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoiceEnum provided by user to PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
	setChoice(value PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoiceEnum) PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
	SetValue(value uint32) PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
	// HasValue checks if Value has been set in PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
	SetValues(value []uint32) PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
	// Increment returns PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter, set in PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize.
	// PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter is integer counter pattern
	Increment() PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
	// SetIncrement assigns PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter provided by user to PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize.
	// PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter is integer counter pattern
	SetIncrement(value PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter, set in PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize.
	// PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter is integer counter pattern
	Decrement() PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
	// SetDecrement assigns PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter provided by user to PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize.
	// PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter is integer counter pattern
	SetDecrement(value PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize
var PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoice = struct {
	VALUE     PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoiceEnum
	VALUES    PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoiceEnum
	INCREMENT PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoiceEnum
	DECREMENT PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSize) Choice() PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoiceEnum {
	return PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSize) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSize) setChoice(value PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoiceEnum) PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize {
	intValue, ok := otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter().msg()
	}

	if value == PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSize) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSize) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize object
func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSize) SetValue(value uint32) PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSize) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize object
func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSize) SetValues(value []uint32) PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSize) Increment() PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSize) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter value in the PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize object
func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSize) SetIncrement(value PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSize) Decrement() PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter
func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSize) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter value in the PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize object
func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSize) SetDecrement(value PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeCounter) PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSize) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowRSVPPathSenderTspecIntServMaximumPacketSize) setDefault() {
	var choices_set int = 0
	var choice PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowRSVPPathSenderTspecIntServMaximumPacketSizeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize")
			}
		} else {
			intVal := otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowRSVPPathSenderTspecIntServMaximumPacketSize_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
