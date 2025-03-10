package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 *****
type patternFlowRSVPPathSessionExtTunnelIdAsIpv4 struct {
	validation
	obj             *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4
	marshaller      marshalPatternFlowRSVPPathSessionExtTunnelIdAsIpv4
	unMarshaller    unMarshalPatternFlowRSVPPathSessionExtTunnelIdAsIpv4
	incrementHolder PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
	decrementHolder PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
}

func NewPatternFlowRSVPPathSessionExtTunnelIdAsIpv4() PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 {
	obj := patternFlowRSVPPathSessionExtTunnelIdAsIpv4{obj: &otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4) msg() *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 {
	return obj.obj
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4) setMsg(msg *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4) PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSessionExtTunnelIdAsIpv4 struct {
	obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4
}

type marshalPatternFlowRSVPPathSessionExtTunnelIdAsIpv4 interface {
	// ToProto marshals PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 to protobuf object *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4
	ToProto() (*otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4, error)
	// ToPbText marshals PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowRSVPPathSessionExtTunnelIdAsIpv4 struct {
	obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4
}

type unMarshalPatternFlowRSVPPathSessionExtTunnelIdAsIpv4 interface {
	// FromProto unmarshals PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 from protobuf object *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4
	FromProto(msg *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4) (PatternFlowRSVPPathSessionExtTunnelIdAsIpv4, error)
	// FromPbText unmarshals PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4) Marshal() marshalPatternFlowRSVPPathSessionExtTunnelIdAsIpv4 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSessionExtTunnelIdAsIpv4{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4) Unmarshal() unMarshalPatternFlowRSVPPathSessionExtTunnelIdAsIpv4 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSessionExtTunnelIdAsIpv4{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSessionExtTunnelIdAsIpv4) ToProto() (*otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSessionExtTunnelIdAsIpv4) FromProto(msg *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4) (PatternFlowRSVPPathSessionExtTunnelIdAsIpv4, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSessionExtTunnelIdAsIpv4) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionExtTunnelIdAsIpv4) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSessionExtTunnelIdAsIpv4) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionExtTunnelIdAsIpv4) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSessionExtTunnelIdAsIpv4) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowRSVPPathSessionExtTunnelIdAsIpv4) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionExtTunnelIdAsIpv4) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4) Clone() (PatternFlowRSVPPathSessionExtTunnelIdAsIpv4, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSessionExtTunnelIdAsIpv4()
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

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 is iPv4 address of the ingress endpoint for the tunnel.
type PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 interface {
	Validation
	// msg marshals PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 to protobuf object *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4
	// setMsg unmarshals PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 from protobuf object *otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4) PatternFlowRSVPPathSessionExtTunnelIdAsIpv4
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSessionExtTunnelIdAsIpv4
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSessionExtTunnelIdAsIpv4
	// validate validates PatternFlowRSVPPathSessionExtTunnelIdAsIpv4
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSessionExtTunnelIdAsIpv4, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathSessionExtTunnelIdAsIpv4ChoiceEnum, set in PatternFlowRSVPPathSessionExtTunnelIdAsIpv4
	Choice() PatternFlowRSVPPathSessionExtTunnelIdAsIpv4ChoiceEnum
	// setChoice assigns PatternFlowRSVPPathSessionExtTunnelIdAsIpv4ChoiceEnum provided by user to PatternFlowRSVPPathSessionExtTunnelIdAsIpv4
	setChoice(value PatternFlowRSVPPathSessionExtTunnelIdAsIpv4ChoiceEnum) PatternFlowRSVPPathSessionExtTunnelIdAsIpv4
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathSessionExtTunnelIdAsIpv4
	HasChoice() bool
	// Value returns string, set in PatternFlowRSVPPathSessionExtTunnelIdAsIpv4.
	Value() string
	// SetValue assigns string provided by user to PatternFlowRSVPPathSessionExtTunnelIdAsIpv4
	SetValue(value string) PatternFlowRSVPPathSessionExtTunnelIdAsIpv4
	// HasValue checks if Value has been set in PatternFlowRSVPPathSessionExtTunnelIdAsIpv4
	HasValue() bool
	// Values returns []string, set in PatternFlowRSVPPathSessionExtTunnelIdAsIpv4.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowRSVPPathSessionExtTunnelIdAsIpv4
	SetValues(value []string) PatternFlowRSVPPathSessionExtTunnelIdAsIpv4
	// Increment returns PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter, set in PatternFlowRSVPPathSessionExtTunnelIdAsIpv4.
	// PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter is ipv4 counter pattern
	Increment() PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
	// SetIncrement assigns PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter provided by user to PatternFlowRSVPPathSessionExtTunnelIdAsIpv4.
	// PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter is ipv4 counter pattern
	SetIncrement(value PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) PatternFlowRSVPPathSessionExtTunnelIdAsIpv4
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathSessionExtTunnelIdAsIpv4
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter, set in PatternFlowRSVPPathSessionExtTunnelIdAsIpv4.
	// PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter is ipv4 counter pattern
	Decrement() PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
	// SetDecrement assigns PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter provided by user to PatternFlowRSVPPathSessionExtTunnelIdAsIpv4.
	// PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter is ipv4 counter pattern
	SetDecrement(value PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) PatternFlowRSVPPathSessionExtTunnelIdAsIpv4
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathSessionExtTunnelIdAsIpv4
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathSessionExtTunnelIdAsIpv4ChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathSessionExtTunnelIdAsIpv4
var PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Choice = struct {
	VALUE     PatternFlowRSVPPathSessionExtTunnelIdAsIpv4ChoiceEnum
	VALUES    PatternFlowRSVPPathSessionExtTunnelIdAsIpv4ChoiceEnum
	INCREMENT PatternFlowRSVPPathSessionExtTunnelIdAsIpv4ChoiceEnum
	DECREMENT PatternFlowRSVPPathSessionExtTunnelIdAsIpv4ChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathSessionExtTunnelIdAsIpv4ChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathSessionExtTunnelIdAsIpv4ChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathSessionExtTunnelIdAsIpv4ChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathSessionExtTunnelIdAsIpv4ChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4) Choice() PatternFlowRSVPPathSessionExtTunnelIdAsIpv4ChoiceEnum {
	return PatternFlowRSVPPathSessionExtTunnelIdAsIpv4ChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4) setChoice(value PatternFlowRSVPPathSessionExtTunnelIdAsIpv4ChoiceEnum) PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 {
	intValue, ok := otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathSessionExtTunnelIdAsIpv4ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Choice.VALUE {
		defaultValue := "0.0.0.0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Choice.VALUES {
		defaultValue := []string{"0.0.0.0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Choice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter().msg()
	}

	if value == PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Choice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Choice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 object
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4) SetValue(value string) PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 {
	obj.setChoice(PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Choice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"0.0.0.0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 object
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4) SetValues(value []string) PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 {
	obj.setChoice(PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Choice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4) Increment() PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Choice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter value in the PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 object
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4) SetIncrement(value PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 {
	obj.setChoice(PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Choice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4) Decrement() PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Choice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter value in the PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 object
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4) SetDecrement(value PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Counter) PatternFlowRSVPPathSessionExtTunnelIdAsIpv4 {
	obj.setChoice(PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Choice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowRSVPPathSessionExtTunnelIdAsIpv4.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowRSVPPathSessionExtTunnelIdAsIpv4.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsIpv4) setDefault() {
	var choices_set int = 0
	var choice PatternFlowRSVPPathSessionExtTunnelIdAsIpv4ChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Choice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Choice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Choice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Choice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowRSVPPathSessionExtTunnelIdAsIpv4Choice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowRSVPPathSessionExtTunnelIdAsIpv4")
			}
		} else {
			intVal := otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowRSVPPathSessionExtTunnelIdAsIpv4_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
