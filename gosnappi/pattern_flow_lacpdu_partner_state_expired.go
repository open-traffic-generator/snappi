package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerStateExpired *****
type patternFlowLacpduPartnerStateExpired struct {
	validation
	obj             *otg.PatternFlowLacpduPartnerStateExpired
	marshaller      marshalPatternFlowLacpduPartnerStateExpired
	unMarshaller    unMarshalPatternFlowLacpduPartnerStateExpired
	incrementHolder PatternFlowLacpduPartnerStateExpiredCounter
	decrementHolder PatternFlowLacpduPartnerStateExpiredCounter
}

func NewPatternFlowLacpduPartnerStateExpired() PatternFlowLacpduPartnerStateExpired {
	obj := patternFlowLacpduPartnerStateExpired{obj: &otg.PatternFlowLacpduPartnerStateExpired{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerStateExpired) msg() *otg.PatternFlowLacpduPartnerStateExpired {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerStateExpired) setMsg(msg *otg.PatternFlowLacpduPartnerStateExpired) PatternFlowLacpduPartnerStateExpired {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerStateExpired struct {
	obj *patternFlowLacpduPartnerStateExpired
}

type marshalPatternFlowLacpduPartnerStateExpired interface {
	// ToProto marshals PatternFlowLacpduPartnerStateExpired to protobuf object *otg.PatternFlowLacpduPartnerStateExpired
	ToProto() (*otg.PatternFlowLacpduPartnerStateExpired, error)
	// ToPbText marshals PatternFlowLacpduPartnerStateExpired to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerStateExpired to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerStateExpired to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerStateExpired struct {
	obj *patternFlowLacpduPartnerStateExpired
}

type unMarshalPatternFlowLacpduPartnerStateExpired interface {
	// FromProto unmarshals PatternFlowLacpduPartnerStateExpired from protobuf object *otg.PatternFlowLacpduPartnerStateExpired
	FromProto(msg *otg.PatternFlowLacpduPartnerStateExpired) (PatternFlowLacpduPartnerStateExpired, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerStateExpired from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerStateExpired from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerStateExpired from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerStateExpired) Marshal() marshalPatternFlowLacpduPartnerStateExpired {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerStateExpired{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerStateExpired) Unmarshal() unMarshalPatternFlowLacpduPartnerStateExpired {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerStateExpired{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerStateExpired) ToProto() (*otg.PatternFlowLacpduPartnerStateExpired, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerStateExpired) FromProto(msg *otg.PatternFlowLacpduPartnerStateExpired) (PatternFlowLacpduPartnerStateExpired, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerStateExpired) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateExpired) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateExpired) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateExpired) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateExpired) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateExpired) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerStateExpired) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateExpired) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateExpired) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerStateExpired) Clone() (PatternFlowLacpduPartnerStateExpired, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerStateExpired()
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

func (obj *patternFlowLacpduPartnerStateExpired) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduPartnerStateExpired is expired (1=Expired, 0=Not Expired)
type PatternFlowLacpduPartnerStateExpired interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerStateExpired to protobuf object *otg.PatternFlowLacpduPartnerStateExpired
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerStateExpired
	// setMsg unmarshals PatternFlowLacpduPartnerStateExpired from protobuf object *otg.PatternFlowLacpduPartnerStateExpired
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerStateExpired) PatternFlowLacpduPartnerStateExpired
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerStateExpired
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerStateExpired
	// validate validates PatternFlowLacpduPartnerStateExpired
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerStateExpired, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduPartnerStateExpiredChoiceEnum, set in PatternFlowLacpduPartnerStateExpired
	Choice() PatternFlowLacpduPartnerStateExpiredChoiceEnum
	// setChoice assigns PatternFlowLacpduPartnerStateExpiredChoiceEnum provided by user to PatternFlowLacpduPartnerStateExpired
	setChoice(value PatternFlowLacpduPartnerStateExpiredChoiceEnum) PatternFlowLacpduPartnerStateExpired
	// HasChoice checks if Choice has been set in PatternFlowLacpduPartnerStateExpired
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduPartnerStateExpired.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduPartnerStateExpired
	SetValue(value uint32) PatternFlowLacpduPartnerStateExpired
	// HasValue checks if Value has been set in PatternFlowLacpduPartnerStateExpired
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduPartnerStateExpired.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduPartnerStateExpired
	SetValues(value []uint32) PatternFlowLacpduPartnerStateExpired
	// Increment returns PatternFlowLacpduPartnerStateExpiredCounter, set in PatternFlowLacpduPartnerStateExpired.
	// PatternFlowLacpduPartnerStateExpiredCounter is integer counter pattern
	Increment() PatternFlowLacpduPartnerStateExpiredCounter
	// SetIncrement assigns PatternFlowLacpduPartnerStateExpiredCounter provided by user to PatternFlowLacpduPartnerStateExpired.
	// PatternFlowLacpduPartnerStateExpiredCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduPartnerStateExpiredCounter) PatternFlowLacpduPartnerStateExpired
	// HasIncrement checks if Increment has been set in PatternFlowLacpduPartnerStateExpired
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduPartnerStateExpiredCounter, set in PatternFlowLacpduPartnerStateExpired.
	// PatternFlowLacpduPartnerStateExpiredCounter is integer counter pattern
	Decrement() PatternFlowLacpduPartnerStateExpiredCounter
	// SetDecrement assigns PatternFlowLacpduPartnerStateExpiredCounter provided by user to PatternFlowLacpduPartnerStateExpired.
	// PatternFlowLacpduPartnerStateExpiredCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduPartnerStateExpiredCounter) PatternFlowLacpduPartnerStateExpired
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduPartnerStateExpired
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduPartnerStateExpiredChoiceEnum string

// Enum of Choice on PatternFlowLacpduPartnerStateExpired
var PatternFlowLacpduPartnerStateExpiredChoice = struct {
	VALUE     PatternFlowLacpduPartnerStateExpiredChoiceEnum
	VALUES    PatternFlowLacpduPartnerStateExpiredChoiceEnum
	INCREMENT PatternFlowLacpduPartnerStateExpiredChoiceEnum
	DECREMENT PatternFlowLacpduPartnerStateExpiredChoiceEnum
}{
	VALUE:     PatternFlowLacpduPartnerStateExpiredChoiceEnum("value"),
	VALUES:    PatternFlowLacpduPartnerStateExpiredChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduPartnerStateExpiredChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduPartnerStateExpiredChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduPartnerStateExpired) Choice() PatternFlowLacpduPartnerStateExpiredChoiceEnum {
	return PatternFlowLacpduPartnerStateExpiredChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduPartnerStateExpired) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduPartnerStateExpired) setChoice(value PatternFlowLacpduPartnerStateExpiredChoiceEnum) PatternFlowLacpduPartnerStateExpired {
	intValue, ok := otg.PatternFlowLacpduPartnerStateExpired_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduPartnerStateExpiredChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduPartnerStateExpired_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduPartnerStateExpiredChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduPartnerStateExpiredChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduPartnerStateExpiredChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduPartnerStateExpiredCounter().msg()
	}

	if value == PatternFlowLacpduPartnerStateExpiredChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduPartnerStateExpiredCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerStateExpired) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduPartnerStateExpiredChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerStateExpired) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduPartnerStateExpired object
func (obj *patternFlowLacpduPartnerStateExpired) SetValue(value uint32) PatternFlowLacpduPartnerStateExpired {
	obj.setChoice(PatternFlowLacpduPartnerStateExpiredChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduPartnerStateExpired) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduPartnerStateExpired object
func (obj *patternFlowLacpduPartnerStateExpired) SetValues(value []uint32) PatternFlowLacpduPartnerStateExpired {
	obj.setChoice(PatternFlowLacpduPartnerStateExpiredChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerStateExpiredCounter
func (obj *patternFlowLacpduPartnerStateExpired) Increment() PatternFlowLacpduPartnerStateExpiredCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduPartnerStateExpiredChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduPartnerStateExpiredCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerStateExpiredCounter
func (obj *patternFlowLacpduPartnerStateExpired) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduPartnerStateExpiredCounter value in the PatternFlowLacpduPartnerStateExpired object
func (obj *patternFlowLacpduPartnerStateExpired) SetIncrement(value PatternFlowLacpduPartnerStateExpiredCounter) PatternFlowLacpduPartnerStateExpired {
	obj.setChoice(PatternFlowLacpduPartnerStateExpiredChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerStateExpiredCounter
func (obj *patternFlowLacpduPartnerStateExpired) Decrement() PatternFlowLacpduPartnerStateExpiredCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduPartnerStateExpiredChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduPartnerStateExpiredCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerStateExpiredCounter
func (obj *patternFlowLacpduPartnerStateExpired) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduPartnerStateExpiredCounter value in the PatternFlowLacpduPartnerStateExpired object
func (obj *patternFlowLacpduPartnerStateExpired) SetDecrement(value PatternFlowLacpduPartnerStateExpiredCounter) PatternFlowLacpduPartnerStateExpired {
	obj.setChoice(PatternFlowLacpduPartnerStateExpiredChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduPartnerStateExpired) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateExpired.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduPartnerStateExpired.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpduPartnerStateExpired) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduPartnerStateExpiredChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateExpiredChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateExpiredChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateExpiredChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateExpiredChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduPartnerStateExpiredChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduPartnerStateExpired")
			}
		} else {
			intVal := otg.PatternFlowLacpduPartnerStateExpired_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduPartnerStateExpired_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
