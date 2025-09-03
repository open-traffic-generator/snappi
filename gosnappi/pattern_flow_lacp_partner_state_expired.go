package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerStateExpired *****
type patternFlowLacpPartnerStateExpired struct {
	validation
	obj             *otg.PatternFlowLacpPartnerStateExpired
	marshaller      marshalPatternFlowLacpPartnerStateExpired
	unMarshaller    unMarshalPatternFlowLacpPartnerStateExpired
	incrementHolder PatternFlowLacpPartnerStateExpiredCounter
	decrementHolder PatternFlowLacpPartnerStateExpiredCounter
}

func NewPatternFlowLacpPartnerStateExpired() PatternFlowLacpPartnerStateExpired {
	obj := patternFlowLacpPartnerStateExpired{obj: &otg.PatternFlowLacpPartnerStateExpired{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerStateExpired) msg() *otg.PatternFlowLacpPartnerStateExpired {
	return obj.obj
}

func (obj *patternFlowLacpPartnerStateExpired) setMsg(msg *otg.PatternFlowLacpPartnerStateExpired) PatternFlowLacpPartnerStateExpired {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerStateExpired struct {
	obj *patternFlowLacpPartnerStateExpired
}

type marshalPatternFlowLacpPartnerStateExpired interface {
	// ToProto marshals PatternFlowLacpPartnerStateExpired to protobuf object *otg.PatternFlowLacpPartnerStateExpired
	ToProto() (*otg.PatternFlowLacpPartnerStateExpired, error)
	// ToPbText marshals PatternFlowLacpPartnerStateExpired to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerStateExpired to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerStateExpired to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerStateExpired struct {
	obj *patternFlowLacpPartnerStateExpired
}

type unMarshalPatternFlowLacpPartnerStateExpired interface {
	// FromProto unmarshals PatternFlowLacpPartnerStateExpired from protobuf object *otg.PatternFlowLacpPartnerStateExpired
	FromProto(msg *otg.PatternFlowLacpPartnerStateExpired) (PatternFlowLacpPartnerStateExpired, error)
	// FromPbText unmarshals PatternFlowLacpPartnerStateExpired from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerStateExpired from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerStateExpired from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerStateExpired) Marshal() marshalPatternFlowLacpPartnerStateExpired {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerStateExpired{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerStateExpired) Unmarshal() unMarshalPatternFlowLacpPartnerStateExpired {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerStateExpired{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerStateExpired) ToProto() (*otg.PatternFlowLacpPartnerStateExpired, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerStateExpired) FromProto(msg *otg.PatternFlowLacpPartnerStateExpired) (PatternFlowLacpPartnerStateExpired, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerStateExpired) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateExpired) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateExpired) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateExpired) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateExpired) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateExpired) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerStateExpired) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateExpired) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateExpired) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerStateExpired) Clone() (PatternFlowLacpPartnerStateExpired, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerStateExpired()
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

func (obj *patternFlowLacpPartnerStateExpired) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpPartnerStateExpired is expired (1=Expired, 0=Not Expired)
type PatternFlowLacpPartnerStateExpired interface {
	Validation
	// msg marshals PatternFlowLacpPartnerStateExpired to protobuf object *otg.PatternFlowLacpPartnerStateExpired
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerStateExpired
	// setMsg unmarshals PatternFlowLacpPartnerStateExpired from protobuf object *otg.PatternFlowLacpPartnerStateExpired
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerStateExpired) PatternFlowLacpPartnerStateExpired
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerStateExpired
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerStateExpired
	// validate validates PatternFlowLacpPartnerStateExpired
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerStateExpired, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpPartnerStateExpiredChoiceEnum, set in PatternFlowLacpPartnerStateExpired
	Choice() PatternFlowLacpPartnerStateExpiredChoiceEnum
	// setChoice assigns PatternFlowLacpPartnerStateExpiredChoiceEnum provided by user to PatternFlowLacpPartnerStateExpired
	setChoice(value PatternFlowLacpPartnerStateExpiredChoiceEnum) PatternFlowLacpPartnerStateExpired
	// HasChoice checks if Choice has been set in PatternFlowLacpPartnerStateExpired
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpPartnerStateExpired.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpPartnerStateExpired
	SetValue(value uint32) PatternFlowLacpPartnerStateExpired
	// HasValue checks if Value has been set in PatternFlowLacpPartnerStateExpired
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpPartnerStateExpired.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpPartnerStateExpired
	SetValues(value []uint32) PatternFlowLacpPartnerStateExpired
	// Increment returns PatternFlowLacpPartnerStateExpiredCounter, set in PatternFlowLacpPartnerStateExpired.
	// PatternFlowLacpPartnerStateExpiredCounter is integer counter pattern
	Increment() PatternFlowLacpPartnerStateExpiredCounter
	// SetIncrement assigns PatternFlowLacpPartnerStateExpiredCounter provided by user to PatternFlowLacpPartnerStateExpired.
	// PatternFlowLacpPartnerStateExpiredCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpPartnerStateExpiredCounter) PatternFlowLacpPartnerStateExpired
	// HasIncrement checks if Increment has been set in PatternFlowLacpPartnerStateExpired
	HasIncrement() bool
	// Decrement returns PatternFlowLacpPartnerStateExpiredCounter, set in PatternFlowLacpPartnerStateExpired.
	// PatternFlowLacpPartnerStateExpiredCounter is integer counter pattern
	Decrement() PatternFlowLacpPartnerStateExpiredCounter
	// SetDecrement assigns PatternFlowLacpPartnerStateExpiredCounter provided by user to PatternFlowLacpPartnerStateExpired.
	// PatternFlowLacpPartnerStateExpiredCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpPartnerStateExpiredCounter) PatternFlowLacpPartnerStateExpired
	// HasDecrement checks if Decrement has been set in PatternFlowLacpPartnerStateExpired
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpPartnerStateExpiredChoiceEnum string

// Enum of Choice on PatternFlowLacpPartnerStateExpired
var PatternFlowLacpPartnerStateExpiredChoice = struct {
	VALUE     PatternFlowLacpPartnerStateExpiredChoiceEnum
	VALUES    PatternFlowLacpPartnerStateExpiredChoiceEnum
	INCREMENT PatternFlowLacpPartnerStateExpiredChoiceEnum
	DECREMENT PatternFlowLacpPartnerStateExpiredChoiceEnum
}{
	VALUE:     PatternFlowLacpPartnerStateExpiredChoiceEnum("value"),
	VALUES:    PatternFlowLacpPartnerStateExpiredChoiceEnum("values"),
	INCREMENT: PatternFlowLacpPartnerStateExpiredChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpPartnerStateExpiredChoiceEnum("decrement"),
}

func (obj *patternFlowLacpPartnerStateExpired) Choice() PatternFlowLacpPartnerStateExpiredChoiceEnum {
	return PatternFlowLacpPartnerStateExpiredChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpPartnerStateExpired) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpPartnerStateExpired) setChoice(value PatternFlowLacpPartnerStateExpiredChoiceEnum) PatternFlowLacpPartnerStateExpired {
	intValue, ok := otg.PatternFlowLacpPartnerStateExpired_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpPartnerStateExpiredChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpPartnerStateExpired_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpPartnerStateExpiredChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpPartnerStateExpiredChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpPartnerStateExpiredChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpPartnerStateExpiredCounter().msg()
	}

	if value == PatternFlowLacpPartnerStateExpiredChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpPartnerStateExpiredCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerStateExpired) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpPartnerStateExpiredChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerStateExpired) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpPartnerStateExpired object
func (obj *patternFlowLacpPartnerStateExpired) SetValue(value uint32) PatternFlowLacpPartnerStateExpired {
	obj.setChoice(PatternFlowLacpPartnerStateExpiredChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpPartnerStateExpired) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpPartnerStateExpired object
func (obj *patternFlowLacpPartnerStateExpired) SetValues(value []uint32) PatternFlowLacpPartnerStateExpired {
	obj.setChoice(PatternFlowLacpPartnerStateExpiredChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerStateExpiredCounter
func (obj *patternFlowLacpPartnerStateExpired) Increment() PatternFlowLacpPartnerStateExpiredCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpPartnerStateExpiredChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpPartnerStateExpiredCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerStateExpiredCounter
func (obj *patternFlowLacpPartnerStateExpired) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpPartnerStateExpiredCounter value in the PatternFlowLacpPartnerStateExpired object
func (obj *patternFlowLacpPartnerStateExpired) SetIncrement(value PatternFlowLacpPartnerStateExpiredCounter) PatternFlowLacpPartnerStateExpired {
	obj.setChoice(PatternFlowLacpPartnerStateExpiredChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerStateExpiredCounter
func (obj *patternFlowLacpPartnerStateExpired) Decrement() PatternFlowLacpPartnerStateExpiredCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpPartnerStateExpiredChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpPartnerStateExpiredCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerStateExpiredCounter
func (obj *patternFlowLacpPartnerStateExpired) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpPartnerStateExpiredCounter value in the PatternFlowLacpPartnerStateExpired object
func (obj *patternFlowLacpPartnerStateExpired) SetDecrement(value PatternFlowLacpPartnerStateExpiredCounter) PatternFlowLacpPartnerStateExpired {
	obj.setChoice(PatternFlowLacpPartnerStateExpiredChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpPartnerStateExpired) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateExpired.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpPartnerStateExpired.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpPartnerStateExpired) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpPartnerStateExpiredChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateExpiredChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateExpiredChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateExpiredChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateExpiredChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpPartnerStateExpiredChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpPartnerStateExpired")
			}
		} else {
			intVal := otg.PatternFlowLacpPartnerStateExpired_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpPartnerStateExpired_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
