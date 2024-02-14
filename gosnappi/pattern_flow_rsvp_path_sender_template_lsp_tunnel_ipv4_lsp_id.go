package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId *****
type patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId struct {
	validation
	obj             *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
	marshaller      marshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
	unMarshaller    unMarshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
	incrementHolder PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
	decrementHolder PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
}

func NewPatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId {
	obj := patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId{obj: &otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) msg() *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) setMsg(msg *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId struct {
	obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
}

type marshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId interface {
	// ToProto marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId to protobuf object *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
	ToProto() (*otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId struct {
	obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
}

type unMarshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId from protobuf object *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
	FromProto(msg *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) (PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) Marshal() marshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) Unmarshal() unMarshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) ToProto() (*otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) FromProto(msg *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) (PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) Clone() (PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId()
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

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId is a 16-bit identifier used in the SENDER_TEMPLATE that can be changed to allow a sender to share resources with itself.
type PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId to protobuf object *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
	// setMsg unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId from protobuf object *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
	// validate validates PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoiceEnum, set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
	Choice() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoiceEnum
	// setChoice assigns PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoiceEnum provided by user to PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
	setChoice(value PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoiceEnum) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
	SetValue(value uint32) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
	// HasValue checks if Value has been set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
	SetValues(value []uint32) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
	// Increment returns PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter, set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId.
	// PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter is integer counter pattern
	Increment() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
	// SetIncrement assigns PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter provided by user to PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId.
	// PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter is integer counter pattern
	SetIncrement(value PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter, set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId.
	// PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter is integer counter pattern
	Decrement() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
	// SetDecrement assigns PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter provided by user to PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId.
	// PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter is integer counter pattern
	SetDecrement(value PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId
var PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoice = struct {
	VALUE     PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoiceEnum
	VALUES    PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoiceEnum
	INCREMENT PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoiceEnum
	DECREMENT PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) Choice() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoiceEnum {
	return PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) setChoice(value PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoiceEnum) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId {
	intValue, ok := otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoice.VALUE {
		defaultValue := uint32(1)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoice.VALUES {
		defaultValue := []uint32{1}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter().msg()
	}

	if value == PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId object
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) SetValue(value uint32) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId {
	obj.setChoice(PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{1})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId object
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) SetValues(value []uint32) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId {
	obj.setChoice(PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) Increment() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter value in the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId object
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) SetIncrement(value PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId {
	obj.setChoice(PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) Decrement() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter value in the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId object
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) SetDecrement(value PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdCounter) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId {
	obj.setChoice(PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4LspId) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTemplateLspTunnelIpv4LspIdChoice.VALUE)

	}

}
