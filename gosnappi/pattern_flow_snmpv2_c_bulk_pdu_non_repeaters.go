package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowSnmpv2CBulkPDUNonRepeaters *****
type patternFlowSnmpv2CBulkPDUNonRepeaters struct {
	validation
	obj          *otg.PatternFlowSnmpv2CBulkPDUNonRepeaters
	marshaller   marshalPatternFlowSnmpv2CBulkPDUNonRepeaters
	unMarshaller unMarshalPatternFlowSnmpv2CBulkPDUNonRepeaters
}

func NewPatternFlowSnmpv2CBulkPDUNonRepeaters() PatternFlowSnmpv2CBulkPDUNonRepeaters {
	obj := patternFlowSnmpv2CBulkPDUNonRepeaters{obj: &otg.PatternFlowSnmpv2CBulkPDUNonRepeaters{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowSnmpv2CBulkPDUNonRepeaters) msg() *otg.PatternFlowSnmpv2CBulkPDUNonRepeaters {
	return obj.obj
}

func (obj *patternFlowSnmpv2CBulkPDUNonRepeaters) setMsg(msg *otg.PatternFlowSnmpv2CBulkPDUNonRepeaters) PatternFlowSnmpv2CBulkPDUNonRepeaters {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowSnmpv2CBulkPDUNonRepeaters struct {
	obj *patternFlowSnmpv2CBulkPDUNonRepeaters
}

type marshalPatternFlowSnmpv2CBulkPDUNonRepeaters interface {
	// ToProto marshals PatternFlowSnmpv2CBulkPDUNonRepeaters to protobuf object *otg.PatternFlowSnmpv2CBulkPDUNonRepeaters
	ToProto() (*otg.PatternFlowSnmpv2CBulkPDUNonRepeaters, error)
	// ToPbText marshals PatternFlowSnmpv2CBulkPDUNonRepeaters to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowSnmpv2CBulkPDUNonRepeaters to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowSnmpv2CBulkPDUNonRepeaters to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowSnmpv2CBulkPDUNonRepeaters struct {
	obj *patternFlowSnmpv2CBulkPDUNonRepeaters
}

type unMarshalPatternFlowSnmpv2CBulkPDUNonRepeaters interface {
	// FromProto unmarshals PatternFlowSnmpv2CBulkPDUNonRepeaters from protobuf object *otg.PatternFlowSnmpv2CBulkPDUNonRepeaters
	FromProto(msg *otg.PatternFlowSnmpv2CBulkPDUNonRepeaters) (PatternFlowSnmpv2CBulkPDUNonRepeaters, error)
	// FromPbText unmarshals PatternFlowSnmpv2CBulkPDUNonRepeaters from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowSnmpv2CBulkPDUNonRepeaters from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowSnmpv2CBulkPDUNonRepeaters from JSON text
	FromJson(value string) error
}

func (obj *patternFlowSnmpv2CBulkPDUNonRepeaters) Marshal() marshalPatternFlowSnmpv2CBulkPDUNonRepeaters {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowSnmpv2CBulkPDUNonRepeaters{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowSnmpv2CBulkPDUNonRepeaters) Unmarshal() unMarshalPatternFlowSnmpv2CBulkPDUNonRepeaters {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowSnmpv2CBulkPDUNonRepeaters{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowSnmpv2CBulkPDUNonRepeaters) ToProto() (*otg.PatternFlowSnmpv2CBulkPDUNonRepeaters, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowSnmpv2CBulkPDUNonRepeaters) FromProto(msg *otg.PatternFlowSnmpv2CBulkPDUNonRepeaters) (PatternFlowSnmpv2CBulkPDUNonRepeaters, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowSnmpv2CBulkPDUNonRepeaters) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CBulkPDUNonRepeaters) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalpatternFlowSnmpv2CBulkPDUNonRepeaters) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CBulkPDUNonRepeaters) FromYaml(value string) error {
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

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalpatternFlowSnmpv2CBulkPDUNonRepeaters) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CBulkPDUNonRepeaters) FromJson(value string) error {
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

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *patternFlowSnmpv2CBulkPDUNonRepeaters) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CBulkPDUNonRepeaters) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CBulkPDUNonRepeaters) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowSnmpv2CBulkPDUNonRepeaters) Clone() (PatternFlowSnmpv2CBulkPDUNonRepeaters, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowSnmpv2CBulkPDUNonRepeaters()
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

// PatternFlowSnmpv2CBulkPDUNonRepeaters is one variable binding in the Response-PDU is requested for the first non_repeaters variable bindings in the GetBulkRequest.
type PatternFlowSnmpv2CBulkPDUNonRepeaters interface {
	Validation
	// msg marshals PatternFlowSnmpv2CBulkPDUNonRepeaters to protobuf object *otg.PatternFlowSnmpv2CBulkPDUNonRepeaters
	// and doesn't set defaults
	msg() *otg.PatternFlowSnmpv2CBulkPDUNonRepeaters
	// setMsg unmarshals PatternFlowSnmpv2CBulkPDUNonRepeaters from protobuf object *otg.PatternFlowSnmpv2CBulkPDUNonRepeaters
	// and doesn't set defaults
	setMsg(*otg.PatternFlowSnmpv2CBulkPDUNonRepeaters) PatternFlowSnmpv2CBulkPDUNonRepeaters
	// provides marshal interface
	Marshal() marshalPatternFlowSnmpv2CBulkPDUNonRepeaters
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowSnmpv2CBulkPDUNonRepeaters
	// validate validates PatternFlowSnmpv2CBulkPDUNonRepeaters
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowSnmpv2CBulkPDUNonRepeaters, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowSnmpv2CBulkPDUNonRepeatersChoiceEnum, set in PatternFlowSnmpv2CBulkPDUNonRepeaters
	Choice() PatternFlowSnmpv2CBulkPDUNonRepeatersChoiceEnum
	// setChoice assigns PatternFlowSnmpv2CBulkPDUNonRepeatersChoiceEnum provided by user to PatternFlowSnmpv2CBulkPDUNonRepeaters
	setChoice(value PatternFlowSnmpv2CBulkPDUNonRepeatersChoiceEnum) PatternFlowSnmpv2CBulkPDUNonRepeaters
	// HasChoice checks if Choice has been set in PatternFlowSnmpv2CBulkPDUNonRepeaters
	HasChoice() bool
	// Value returns uint32, set in PatternFlowSnmpv2CBulkPDUNonRepeaters.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowSnmpv2CBulkPDUNonRepeaters
	SetValue(value uint32) PatternFlowSnmpv2CBulkPDUNonRepeaters
	// HasValue checks if Value has been set in PatternFlowSnmpv2CBulkPDUNonRepeaters
	HasValue() bool
	// Values returns []uint32, set in PatternFlowSnmpv2CBulkPDUNonRepeaters.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowSnmpv2CBulkPDUNonRepeaters
	SetValues(value []uint32) PatternFlowSnmpv2CBulkPDUNonRepeaters
}

type PatternFlowSnmpv2CBulkPDUNonRepeatersChoiceEnum string

// Enum of Choice on PatternFlowSnmpv2CBulkPDUNonRepeaters
var PatternFlowSnmpv2CBulkPDUNonRepeatersChoice = struct {
	VALUE  PatternFlowSnmpv2CBulkPDUNonRepeatersChoiceEnum
	VALUES PatternFlowSnmpv2CBulkPDUNonRepeatersChoiceEnum
}{
	VALUE:  PatternFlowSnmpv2CBulkPDUNonRepeatersChoiceEnum("value"),
	VALUES: PatternFlowSnmpv2CBulkPDUNonRepeatersChoiceEnum("values"),
}

func (obj *patternFlowSnmpv2CBulkPDUNonRepeaters) Choice() PatternFlowSnmpv2CBulkPDUNonRepeatersChoiceEnum {
	return PatternFlowSnmpv2CBulkPDUNonRepeatersChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowSnmpv2CBulkPDUNonRepeaters) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowSnmpv2CBulkPDUNonRepeaters) setChoice(value PatternFlowSnmpv2CBulkPDUNonRepeatersChoiceEnum) PatternFlowSnmpv2CBulkPDUNonRepeaters {
	intValue, ok := otg.PatternFlowSnmpv2CBulkPDUNonRepeaters_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowSnmpv2CBulkPDUNonRepeatersChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowSnmpv2CBulkPDUNonRepeaters_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowSnmpv2CBulkPDUNonRepeatersChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowSnmpv2CBulkPDUNonRepeatersChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowSnmpv2CBulkPDUNonRepeaters) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowSnmpv2CBulkPDUNonRepeatersChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowSnmpv2CBulkPDUNonRepeaters) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowSnmpv2CBulkPDUNonRepeaters object
func (obj *patternFlowSnmpv2CBulkPDUNonRepeaters) SetValue(value uint32) PatternFlowSnmpv2CBulkPDUNonRepeaters {
	obj.setChoice(PatternFlowSnmpv2CBulkPDUNonRepeatersChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowSnmpv2CBulkPDUNonRepeaters) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowSnmpv2CBulkPDUNonRepeaters object
func (obj *patternFlowSnmpv2CBulkPDUNonRepeaters) SetValues(value []uint32) PatternFlowSnmpv2CBulkPDUNonRepeaters {
	obj.setChoice(PatternFlowSnmpv2CBulkPDUNonRepeatersChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

func (obj *patternFlowSnmpv2CBulkPDUNonRepeaters) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *patternFlowSnmpv2CBulkPDUNonRepeaters) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowSnmpv2CBulkPDUNonRepeatersChoice.VALUE)

	}

}
