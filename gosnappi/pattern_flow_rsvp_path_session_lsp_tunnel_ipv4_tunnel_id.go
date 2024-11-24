package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId *****
type patternFlowRSVPPathSessionLspTunnelIpv4TunnelId struct {
	validation
	obj             *otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
	marshaller      marshalPatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
	unMarshaller    unMarshalPatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
	incrementHolder PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
	decrementHolder PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
}

func NewPatternFlowRSVPPathSessionLspTunnelIpv4TunnelId() PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId {
	obj := patternFlowRSVPPathSessionLspTunnelIpv4TunnelId{obj: &otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelId) msg() *otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId {
	return obj.obj
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelId) setMsg(msg *otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId) PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSessionLspTunnelIpv4TunnelId struct {
	obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelId
}

type marshalPatternFlowRSVPPathSessionLspTunnelIpv4TunnelId interface {
	// ToProto marshals PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId to protobuf object *otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
	ToProto() (*otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId, error)
	// ToPbText marshals PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4TunnelId struct {
	obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelId
}

type unMarshalPatternFlowRSVPPathSessionLspTunnelIpv4TunnelId interface {
	// FromProto unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId from protobuf object *otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
	FromProto(msg *otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId) (PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId, error)
	// FromPbText unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelId) Marshal() marshalPatternFlowRSVPPathSessionLspTunnelIpv4TunnelId {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSessionLspTunnelIpv4TunnelId{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelId) Unmarshal() unMarshalPatternFlowRSVPPathSessionLspTunnelIpv4TunnelId {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4TunnelId{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSessionLspTunnelIpv4TunnelId) ToProto() (*otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4TunnelId) FromProto(msg *otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId) (PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSessionLspTunnelIpv4TunnelId) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4TunnelId) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSessionLspTunnelIpv4TunnelId) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4TunnelId) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSessionLspTunnelIpv4TunnelId) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowRSVPPathSessionLspTunnelIpv4TunnelId) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4TunnelId) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelId) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelId) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelId) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelId) Clone() (PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSessionLspTunnelIpv4TunnelId()
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

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelId) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId is a 16-bit identifier used in the SESSION that remains constant over the life of the tunnel.
type PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId interface {
	Validation
	// msg marshals PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId to protobuf object *otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
	// setMsg unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId from protobuf object *otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId) PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
	// validate validates PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoiceEnum, set in PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
	Choice() PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoiceEnum
	// setChoice assigns PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoiceEnum provided by user to PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
	setChoice(value PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoiceEnum) PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
	SetValue(value uint32) PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
	// HasValue checks if Value has been set in PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
	SetValues(value []uint32) PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
	// Increment returns PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter, set in PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId.
	// PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter is integer counter pattern
	Increment() PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
	// SetIncrement assigns PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter provided by user to PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId.
	// PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter is integer counter pattern
	SetIncrement(value PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter, set in PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId.
	// PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter is integer counter pattern
	Decrement() PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
	// SetDecrement assigns PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter provided by user to PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId.
	// PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter is integer counter pattern
	SetDecrement(value PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId
var PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoice = struct {
	VALUE     PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoiceEnum
	VALUES    PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoiceEnum
	INCREMENT PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoiceEnum
	DECREMENT PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelId) Choice() PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoiceEnum {
	return PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelId) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelId) setChoice(value PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoiceEnum) PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId {
	intValue, ok := otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoice.VALUE {
		defaultValue := uint32(1)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoice.VALUES {
		defaultValue := []uint32{1}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter().msg()
	}

	if value == PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelId) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelId) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId object
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelId) SetValue(value uint32) PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId {
	obj.setChoice(PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelId) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{1})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId object
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelId) SetValues(value []uint32) PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId {
	obj.setChoice(PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelId) Increment() PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelId) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter value in the PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId object
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelId) SetIncrement(value PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId {
	obj.setChoice(PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelId) Decrement() PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelId) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter value in the PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId object
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelId) SetDecrement(value PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdCounter) PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId {
	obj.setChoice(PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelId) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4TunnelId) setDefault() {
	var choices_set int = 0
	var choice PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowRSVPPathSessionLspTunnelIpv4TunnelIdChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId")
			}
		} else {
			intVal := otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowRSVPPathSessionLspTunnelIpv4TunnelId_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
