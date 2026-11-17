package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3MembershipQueryQrv *****
type patternFlowIgmpv3MembershipQueryQrv struct {
	validation
	obj             *otg.PatternFlowIgmpv3MembershipQueryQrv
	marshaller      marshalPatternFlowIgmpv3MembershipQueryQrv
	unMarshaller    unMarshalPatternFlowIgmpv3MembershipQueryQrv
	incrementHolder PatternFlowIgmpv3MembershipQueryQrvCounter
	decrementHolder PatternFlowIgmpv3MembershipQueryQrvCounter
}

func NewPatternFlowIgmpv3MembershipQueryQrv() PatternFlowIgmpv3MembershipQueryQrv {
	obj := patternFlowIgmpv3MembershipQueryQrv{obj: &otg.PatternFlowIgmpv3MembershipQueryQrv{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3MembershipQueryQrv) msg() *otg.PatternFlowIgmpv3MembershipQueryQrv {
	return obj.obj
}

func (obj *patternFlowIgmpv3MembershipQueryQrv) setMsg(msg *otg.PatternFlowIgmpv3MembershipQueryQrv) PatternFlowIgmpv3MembershipQueryQrv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3MembershipQueryQrv struct {
	obj *patternFlowIgmpv3MembershipQueryQrv
}

type marshalPatternFlowIgmpv3MembershipQueryQrv interface {
	// ToProto marshals PatternFlowIgmpv3MembershipQueryQrv to protobuf object *otg.PatternFlowIgmpv3MembershipQueryQrv
	ToProto() (*otg.PatternFlowIgmpv3MembershipQueryQrv, error)
	// ToPbText marshals PatternFlowIgmpv3MembershipQueryQrv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3MembershipQueryQrv to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3MembershipQueryQrv to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3MembershipQueryQrv struct {
	obj *patternFlowIgmpv3MembershipQueryQrv
}

type unMarshalPatternFlowIgmpv3MembershipQueryQrv interface {
	// FromProto unmarshals PatternFlowIgmpv3MembershipQueryQrv from protobuf object *otg.PatternFlowIgmpv3MembershipQueryQrv
	FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryQrv) (PatternFlowIgmpv3MembershipQueryQrv, error)
	// FromPbText unmarshals PatternFlowIgmpv3MembershipQueryQrv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3MembershipQueryQrv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3MembershipQueryQrv from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3MembershipQueryQrv) Marshal() marshalPatternFlowIgmpv3MembershipQueryQrv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3MembershipQueryQrv{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3MembershipQueryQrv) Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryQrv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3MembershipQueryQrv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3MembershipQueryQrv) ToProto() (*otg.PatternFlowIgmpv3MembershipQueryQrv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3MembershipQueryQrv) FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryQrv) (PatternFlowIgmpv3MembershipQueryQrv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3MembershipQueryQrv) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryQrv) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryQrv) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryQrv) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryQrv) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryQrv) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3MembershipQueryQrv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryQrv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryQrv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3MembershipQueryQrv) Clone() (PatternFlowIgmpv3MembershipQueryQrv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3MembershipQueryQrv()
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

func (obj *patternFlowIgmpv3MembershipQueryQrv) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIgmpv3MembershipQueryQrv is querier's robustness variable. Zero means the receiver uses its own value, otherwise it is the value used by the querier.
type PatternFlowIgmpv3MembershipQueryQrv interface {
	Validation
	// msg marshals PatternFlowIgmpv3MembershipQueryQrv to protobuf object *otg.PatternFlowIgmpv3MembershipQueryQrv
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3MembershipQueryQrv
	// setMsg unmarshals PatternFlowIgmpv3MembershipQueryQrv from protobuf object *otg.PatternFlowIgmpv3MembershipQueryQrv
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3MembershipQueryQrv) PatternFlowIgmpv3MembershipQueryQrv
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3MembershipQueryQrv
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryQrv
	// validate validates PatternFlowIgmpv3MembershipQueryQrv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3MembershipQueryQrv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIgmpv3MembershipQueryQrvChoiceEnum, set in PatternFlowIgmpv3MembershipQueryQrv
	Choice() PatternFlowIgmpv3MembershipQueryQrvChoiceEnum
	// setChoice assigns PatternFlowIgmpv3MembershipQueryQrvChoiceEnum provided by user to PatternFlowIgmpv3MembershipQueryQrv
	setChoice(value PatternFlowIgmpv3MembershipQueryQrvChoiceEnum) PatternFlowIgmpv3MembershipQueryQrv
	// HasChoice checks if Choice has been set in PatternFlowIgmpv3MembershipQueryQrv
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIgmpv3MembershipQueryQrv.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIgmpv3MembershipQueryQrv
	SetValue(value uint32) PatternFlowIgmpv3MembershipQueryQrv
	// HasValue checks if Value has been set in PatternFlowIgmpv3MembershipQueryQrv
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIgmpv3MembershipQueryQrv.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIgmpv3MembershipQueryQrv
	SetValues(value []uint32) PatternFlowIgmpv3MembershipQueryQrv
	// Increment returns PatternFlowIgmpv3MembershipQueryQrvCounter, set in PatternFlowIgmpv3MembershipQueryQrv.
	// PatternFlowIgmpv3MembershipQueryQrvCounter is integer counter pattern
	Increment() PatternFlowIgmpv3MembershipQueryQrvCounter
	// SetIncrement assigns PatternFlowIgmpv3MembershipQueryQrvCounter provided by user to PatternFlowIgmpv3MembershipQueryQrv.
	// PatternFlowIgmpv3MembershipQueryQrvCounter is integer counter pattern
	SetIncrement(value PatternFlowIgmpv3MembershipQueryQrvCounter) PatternFlowIgmpv3MembershipQueryQrv
	// HasIncrement checks if Increment has been set in PatternFlowIgmpv3MembershipQueryQrv
	HasIncrement() bool
	// Decrement returns PatternFlowIgmpv3MembershipQueryQrvCounter, set in PatternFlowIgmpv3MembershipQueryQrv.
	// PatternFlowIgmpv3MembershipQueryQrvCounter is integer counter pattern
	Decrement() PatternFlowIgmpv3MembershipQueryQrvCounter
	// SetDecrement assigns PatternFlowIgmpv3MembershipQueryQrvCounter provided by user to PatternFlowIgmpv3MembershipQueryQrv.
	// PatternFlowIgmpv3MembershipQueryQrvCounter is integer counter pattern
	SetDecrement(value PatternFlowIgmpv3MembershipQueryQrvCounter) PatternFlowIgmpv3MembershipQueryQrv
	// HasDecrement checks if Decrement has been set in PatternFlowIgmpv3MembershipQueryQrv
	HasDecrement() bool
	setNil()
}

type PatternFlowIgmpv3MembershipQueryQrvChoiceEnum string

// Enum of Choice on PatternFlowIgmpv3MembershipQueryQrv
var PatternFlowIgmpv3MembershipQueryQrvChoice = struct {
	VALUE     PatternFlowIgmpv3MembershipQueryQrvChoiceEnum
	VALUES    PatternFlowIgmpv3MembershipQueryQrvChoiceEnum
	INCREMENT PatternFlowIgmpv3MembershipQueryQrvChoiceEnum
	DECREMENT PatternFlowIgmpv3MembershipQueryQrvChoiceEnum
}{
	VALUE:     PatternFlowIgmpv3MembershipQueryQrvChoiceEnum("value"),
	VALUES:    PatternFlowIgmpv3MembershipQueryQrvChoiceEnum("values"),
	INCREMENT: PatternFlowIgmpv3MembershipQueryQrvChoiceEnum("increment"),
	DECREMENT: PatternFlowIgmpv3MembershipQueryQrvChoiceEnum("decrement"),
}

func (obj *patternFlowIgmpv3MembershipQueryQrv) Choice() PatternFlowIgmpv3MembershipQueryQrvChoiceEnum {
	return PatternFlowIgmpv3MembershipQueryQrvChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIgmpv3MembershipQueryQrv) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIgmpv3MembershipQueryQrv) setChoice(value PatternFlowIgmpv3MembershipQueryQrvChoiceEnum) PatternFlowIgmpv3MembershipQueryQrv {
	intValue, ok := otg.PatternFlowIgmpv3MembershipQueryQrv_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv3MembershipQueryQrvChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv3MembershipQueryQrv_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIgmpv3MembershipQueryQrvChoice.VALUE {
		defaultValue := uint32(2)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIgmpv3MembershipQueryQrvChoice.VALUES {
		defaultValue := []uint32{2}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIgmpv3MembershipQueryQrvChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIgmpv3MembershipQueryQrvCounter().msg()
	}

	if value == PatternFlowIgmpv3MembershipQueryQrvChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIgmpv3MembershipQueryQrvCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryQrv) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipQueryQrvChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryQrv) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIgmpv3MembershipQueryQrv object
func (obj *patternFlowIgmpv3MembershipQueryQrv) SetValue(value uint32) PatternFlowIgmpv3MembershipQueryQrv {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryQrvChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIgmpv3MembershipQueryQrv) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{2})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIgmpv3MembershipQueryQrv object
func (obj *patternFlowIgmpv3MembershipQueryQrv) SetValues(value []uint32) PatternFlowIgmpv3MembershipQueryQrv {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryQrvChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIgmpv3MembershipQueryQrvCounter
func (obj *patternFlowIgmpv3MembershipQueryQrv) Increment() PatternFlowIgmpv3MembershipQueryQrvCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipQueryQrvChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIgmpv3MembershipQueryQrvCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIgmpv3MembershipQueryQrvCounter
func (obj *patternFlowIgmpv3MembershipQueryQrv) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIgmpv3MembershipQueryQrvCounter value in the PatternFlowIgmpv3MembershipQueryQrv object
func (obj *patternFlowIgmpv3MembershipQueryQrv) SetIncrement(value PatternFlowIgmpv3MembershipQueryQrvCounter) PatternFlowIgmpv3MembershipQueryQrv {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryQrvChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3MembershipQueryQrvCounter
func (obj *patternFlowIgmpv3MembershipQueryQrv) Decrement() PatternFlowIgmpv3MembershipQueryQrvCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipQueryQrvChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIgmpv3MembershipQueryQrvCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3MembershipQueryQrvCounter
func (obj *patternFlowIgmpv3MembershipQueryQrv) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIgmpv3MembershipQueryQrvCounter value in the PatternFlowIgmpv3MembershipQueryQrv object
func (obj *patternFlowIgmpv3MembershipQueryQrv) SetDecrement(value PatternFlowIgmpv3MembershipQueryQrvCounter) PatternFlowIgmpv3MembershipQueryQrv {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryQrvChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIgmpv3MembershipQueryQrv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQueryQrv.Value <= 7 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 7 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIgmpv3MembershipQueryQrv.Values <= 7 but Got %d", item))
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

func (obj *patternFlowIgmpv3MembershipQueryQrv) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIgmpv3MembershipQueryQrvChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryQrvChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryQrvChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryQrvChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryQrvChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIgmpv3MembershipQueryQrvChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIgmpv3MembershipQueryQrv")
			}
		} else {
			intVal := otg.PatternFlowIgmpv3MembershipQueryQrv_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIgmpv3MembershipQueryQrv_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
