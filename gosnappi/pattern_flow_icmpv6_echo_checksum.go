package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6EchoChecksum *****
type patternFlowIcmpv6EchoChecksum struct {
	validation
	obj          *otg.PatternFlowIcmpv6EchoChecksum
	marshaller   marshalPatternFlowIcmpv6EchoChecksum
	unMarshaller unMarshalPatternFlowIcmpv6EchoChecksum
}

func NewPatternFlowIcmpv6EchoChecksum() PatternFlowIcmpv6EchoChecksum {
	obj := patternFlowIcmpv6EchoChecksum{obj: &otg.PatternFlowIcmpv6EchoChecksum{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6EchoChecksum) msg() *otg.PatternFlowIcmpv6EchoChecksum {
	return obj.obj
}

func (obj *patternFlowIcmpv6EchoChecksum) setMsg(msg *otg.PatternFlowIcmpv6EchoChecksum) PatternFlowIcmpv6EchoChecksum {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6EchoChecksum struct {
	obj *patternFlowIcmpv6EchoChecksum
}

type marshalPatternFlowIcmpv6EchoChecksum interface {
	// ToProto marshals PatternFlowIcmpv6EchoChecksum to protobuf object *otg.PatternFlowIcmpv6EchoChecksum
	ToProto() (*otg.PatternFlowIcmpv6EchoChecksum, error)
	// ToPbText marshals PatternFlowIcmpv6EchoChecksum to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6EchoChecksum to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6EchoChecksum to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIcmpv6EchoChecksum to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIcmpv6EchoChecksum struct {
	obj *patternFlowIcmpv6EchoChecksum
}

type unMarshalPatternFlowIcmpv6EchoChecksum interface {
	// FromProto unmarshals PatternFlowIcmpv6EchoChecksum from protobuf object *otg.PatternFlowIcmpv6EchoChecksum
	FromProto(msg *otg.PatternFlowIcmpv6EchoChecksum) (PatternFlowIcmpv6EchoChecksum, error)
	// FromPbText unmarshals PatternFlowIcmpv6EchoChecksum from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6EchoChecksum from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6EchoChecksum from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6EchoChecksum) Marshal() marshalPatternFlowIcmpv6EchoChecksum {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6EchoChecksum{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6EchoChecksum) Unmarshal() unMarshalPatternFlowIcmpv6EchoChecksum {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6EchoChecksum{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6EchoChecksum) ToProto() (*otg.PatternFlowIcmpv6EchoChecksum, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6EchoChecksum) FromProto(msg *otg.PatternFlowIcmpv6EchoChecksum) (PatternFlowIcmpv6EchoChecksum, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6EchoChecksum) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoChecksum) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6EchoChecksum) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoChecksum) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6EchoChecksum) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIcmpv6EchoChecksum) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoChecksum) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6EchoChecksum) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6EchoChecksum) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6EchoChecksum) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6EchoChecksum) Clone() (PatternFlowIcmpv6EchoChecksum, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6EchoChecksum()
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

// PatternFlowIcmpv6EchoChecksum is iCMPv6 checksum
type PatternFlowIcmpv6EchoChecksum interface {
	Validation
	// msg marshals PatternFlowIcmpv6EchoChecksum to protobuf object *otg.PatternFlowIcmpv6EchoChecksum
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6EchoChecksum
	// setMsg unmarshals PatternFlowIcmpv6EchoChecksum from protobuf object *otg.PatternFlowIcmpv6EchoChecksum
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6EchoChecksum) PatternFlowIcmpv6EchoChecksum
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6EchoChecksum
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6EchoChecksum
	// validate validates PatternFlowIcmpv6EchoChecksum
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6EchoChecksum, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpv6EchoChecksumChoiceEnum, set in PatternFlowIcmpv6EchoChecksum
	Choice() PatternFlowIcmpv6EchoChecksumChoiceEnum
	// setChoice assigns PatternFlowIcmpv6EchoChecksumChoiceEnum provided by user to PatternFlowIcmpv6EchoChecksum
	setChoice(value PatternFlowIcmpv6EchoChecksumChoiceEnum) PatternFlowIcmpv6EchoChecksum
	// HasChoice checks if Choice has been set in PatternFlowIcmpv6EchoChecksum
	HasChoice() bool
	// Generated returns PatternFlowIcmpv6EchoChecksumGeneratedEnum, set in PatternFlowIcmpv6EchoChecksum
	Generated() PatternFlowIcmpv6EchoChecksumGeneratedEnum
	// SetGenerated assigns PatternFlowIcmpv6EchoChecksumGeneratedEnum provided by user to PatternFlowIcmpv6EchoChecksum
	SetGenerated(value PatternFlowIcmpv6EchoChecksumGeneratedEnum) PatternFlowIcmpv6EchoChecksum
	// HasGenerated checks if Generated has been set in PatternFlowIcmpv6EchoChecksum
	HasGenerated() bool
	// Custom returns uint32, set in PatternFlowIcmpv6EchoChecksum.
	Custom() uint32
	// SetCustom assigns uint32 provided by user to PatternFlowIcmpv6EchoChecksum
	SetCustom(value uint32) PatternFlowIcmpv6EchoChecksum
	// HasCustom checks if Custom has been set in PatternFlowIcmpv6EchoChecksum
	HasCustom() bool
}

type PatternFlowIcmpv6EchoChecksumChoiceEnum string

// Enum of Choice on PatternFlowIcmpv6EchoChecksum
var PatternFlowIcmpv6EchoChecksumChoice = struct {
	GENERATED PatternFlowIcmpv6EchoChecksumChoiceEnum
	CUSTOM    PatternFlowIcmpv6EchoChecksumChoiceEnum
}{
	GENERATED: PatternFlowIcmpv6EchoChecksumChoiceEnum("generated"),
	CUSTOM:    PatternFlowIcmpv6EchoChecksumChoiceEnum("custom"),
}

func (obj *patternFlowIcmpv6EchoChecksum) Choice() PatternFlowIcmpv6EchoChecksumChoiceEnum {
	return PatternFlowIcmpv6EchoChecksumChoiceEnum(obj.obj.Choice.Enum().String())
}

// The type of checksum
// Choice returns a string
func (obj *patternFlowIcmpv6EchoChecksum) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpv6EchoChecksum) setChoice(value PatternFlowIcmpv6EchoChecksumChoiceEnum) PatternFlowIcmpv6EchoChecksum {
	intValue, ok := otg.PatternFlowIcmpv6EchoChecksum_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpv6EchoChecksumChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpv6EchoChecksum_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Custom = nil
	obj.obj.Generated = otg.PatternFlowIcmpv6EchoChecksum_Generated_unspecified.Enum()
	return obj
}

type PatternFlowIcmpv6EchoChecksumGeneratedEnum string

// Enum of Generated on PatternFlowIcmpv6EchoChecksum
var PatternFlowIcmpv6EchoChecksumGenerated = struct {
	GOOD PatternFlowIcmpv6EchoChecksumGeneratedEnum
	BAD  PatternFlowIcmpv6EchoChecksumGeneratedEnum
}{
	GOOD: PatternFlowIcmpv6EchoChecksumGeneratedEnum("good"),
	BAD:  PatternFlowIcmpv6EchoChecksumGeneratedEnum("bad"),
}

func (obj *patternFlowIcmpv6EchoChecksum) Generated() PatternFlowIcmpv6EchoChecksumGeneratedEnum {
	return PatternFlowIcmpv6EchoChecksumGeneratedEnum(obj.obj.Generated.Enum().String())
}

// A system generated checksum value
// Generated returns a string
func (obj *patternFlowIcmpv6EchoChecksum) HasGenerated() bool {
	return obj.obj.Generated != nil
}

func (obj *patternFlowIcmpv6EchoChecksum) SetGenerated(value PatternFlowIcmpv6EchoChecksumGeneratedEnum) PatternFlowIcmpv6EchoChecksum {
	intValue, ok := otg.PatternFlowIcmpv6EchoChecksum_Generated_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpv6EchoChecksumGeneratedEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpv6EchoChecksum_Generated_Enum(intValue)
	obj.obj.Generated = &enumValue

	return obj
}

// A custom checksum value
// Custom returns a uint32
func (obj *patternFlowIcmpv6EchoChecksum) Custom() uint32 {

	if obj.obj.Custom == nil {
		obj.setChoice(PatternFlowIcmpv6EchoChecksumChoice.CUSTOM)
	}

	return *obj.obj.Custom

}

// A custom checksum value
// Custom returns a uint32
func (obj *patternFlowIcmpv6EchoChecksum) HasCustom() bool {
	return obj.obj.Custom != nil
}

// A custom checksum value
// SetCustom sets the uint32 value in the PatternFlowIcmpv6EchoChecksum object
func (obj *patternFlowIcmpv6EchoChecksum) SetCustom(value uint32) PatternFlowIcmpv6EchoChecksum {
	obj.setChoice(PatternFlowIcmpv6EchoChecksumChoice.CUSTOM)
	obj.obj.Custom = &value
	return obj
}

func (obj *patternFlowIcmpv6EchoChecksum) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {

		if *obj.obj.Custom > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoChecksum.Custom <= 65535 but Got %d", *obj.obj.Custom))
		}

	}

}

func (obj *patternFlowIcmpv6EchoChecksum) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpv6EchoChecksumChoiceEnum

	if obj.obj.Generated != nil && obj.obj.Generated.Number() != 0 {
		choices_set += 1
		choice = PatternFlowIcmpv6EchoChecksumChoice.GENERATED
	}

	if obj.obj.Custom != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6EchoChecksumChoice.CUSTOM
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpv6EchoChecksumChoice.GENERATED)
			if obj.obj.Generated.Number() == 0 {
				obj.SetGenerated(PatternFlowIcmpv6EchoChecksumGenerated.GOOD)

			}

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpv6EchoChecksum")
			}
		} else {
			intVal := otg.PatternFlowIcmpv6EchoChecksum_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpv6EchoChecksum_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
