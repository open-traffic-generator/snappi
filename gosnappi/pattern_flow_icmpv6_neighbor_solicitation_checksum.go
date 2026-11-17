package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NeighborSolicitationChecksum *****
type patternFlowIcmpv6NeighborSolicitationChecksum struct {
	validation
	obj          *otg.PatternFlowIcmpv6NeighborSolicitationChecksum
	marshaller   marshalPatternFlowIcmpv6NeighborSolicitationChecksum
	unMarshaller unMarshalPatternFlowIcmpv6NeighborSolicitationChecksum
}

func NewPatternFlowIcmpv6NeighborSolicitationChecksum() PatternFlowIcmpv6NeighborSolicitationChecksum {
	obj := patternFlowIcmpv6NeighborSolicitationChecksum{obj: &otg.PatternFlowIcmpv6NeighborSolicitationChecksum{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NeighborSolicitationChecksum) msg() *otg.PatternFlowIcmpv6NeighborSolicitationChecksum {
	return obj.obj
}

func (obj *patternFlowIcmpv6NeighborSolicitationChecksum) setMsg(msg *otg.PatternFlowIcmpv6NeighborSolicitationChecksum) PatternFlowIcmpv6NeighborSolicitationChecksum {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NeighborSolicitationChecksum struct {
	obj *patternFlowIcmpv6NeighborSolicitationChecksum
}

type marshalPatternFlowIcmpv6NeighborSolicitationChecksum interface {
	// ToProto marshals PatternFlowIcmpv6NeighborSolicitationChecksum to protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationChecksum
	ToProto() (*otg.PatternFlowIcmpv6NeighborSolicitationChecksum, error)
	// ToPbText marshals PatternFlowIcmpv6NeighborSolicitationChecksum to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NeighborSolicitationChecksum to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NeighborSolicitationChecksum to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NeighborSolicitationChecksum struct {
	obj *patternFlowIcmpv6NeighborSolicitationChecksum
}

type unMarshalPatternFlowIcmpv6NeighborSolicitationChecksum interface {
	// FromProto unmarshals PatternFlowIcmpv6NeighborSolicitationChecksum from protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationChecksum
	FromProto(msg *otg.PatternFlowIcmpv6NeighborSolicitationChecksum) (PatternFlowIcmpv6NeighborSolicitationChecksum, error)
	// FromPbText unmarshals PatternFlowIcmpv6NeighborSolicitationChecksum from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NeighborSolicitationChecksum from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NeighborSolicitationChecksum from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NeighborSolicitationChecksum) Marshal() marshalPatternFlowIcmpv6NeighborSolicitationChecksum {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NeighborSolicitationChecksum{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NeighborSolicitationChecksum) Unmarshal() unMarshalPatternFlowIcmpv6NeighborSolicitationChecksum {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NeighborSolicitationChecksum{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NeighborSolicitationChecksum) ToProto() (*otg.PatternFlowIcmpv6NeighborSolicitationChecksum, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationChecksum) FromProto(msg *otg.PatternFlowIcmpv6NeighborSolicitationChecksum) (PatternFlowIcmpv6NeighborSolicitationChecksum, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NeighborSolicitationChecksum) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationChecksum) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NeighborSolicitationChecksum) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationChecksum) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NeighborSolicitationChecksum) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationChecksum) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NeighborSolicitationChecksum) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NeighborSolicitationChecksum) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NeighborSolicitationChecksum) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NeighborSolicitationChecksum) Clone() (PatternFlowIcmpv6NeighborSolicitationChecksum, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NeighborSolicitationChecksum()
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

// PatternFlowIcmpv6NeighborSolicitationChecksum is iCMPv6 checksum
type PatternFlowIcmpv6NeighborSolicitationChecksum interface {
	Validation
	// msg marshals PatternFlowIcmpv6NeighborSolicitationChecksum to protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationChecksum
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NeighborSolicitationChecksum
	// setMsg unmarshals PatternFlowIcmpv6NeighborSolicitationChecksum from protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationChecksum
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NeighborSolicitationChecksum) PatternFlowIcmpv6NeighborSolicitationChecksum
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NeighborSolicitationChecksum
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NeighborSolicitationChecksum
	// validate validates PatternFlowIcmpv6NeighborSolicitationChecksum
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NeighborSolicitationChecksum, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpv6NeighborSolicitationChecksumChoiceEnum, set in PatternFlowIcmpv6NeighborSolicitationChecksum
	Choice() PatternFlowIcmpv6NeighborSolicitationChecksumChoiceEnum
	// setChoice assigns PatternFlowIcmpv6NeighborSolicitationChecksumChoiceEnum provided by user to PatternFlowIcmpv6NeighborSolicitationChecksum
	setChoice(value PatternFlowIcmpv6NeighborSolicitationChecksumChoiceEnum) PatternFlowIcmpv6NeighborSolicitationChecksum
	// HasChoice checks if Choice has been set in PatternFlowIcmpv6NeighborSolicitationChecksum
	HasChoice() bool
	// Generated returns PatternFlowIcmpv6NeighborSolicitationChecksumGeneratedEnum, set in PatternFlowIcmpv6NeighborSolicitationChecksum
	Generated() PatternFlowIcmpv6NeighborSolicitationChecksumGeneratedEnum
	// SetGenerated assigns PatternFlowIcmpv6NeighborSolicitationChecksumGeneratedEnum provided by user to PatternFlowIcmpv6NeighborSolicitationChecksum
	SetGenerated(value PatternFlowIcmpv6NeighborSolicitationChecksumGeneratedEnum) PatternFlowIcmpv6NeighborSolicitationChecksum
	// HasGenerated checks if Generated has been set in PatternFlowIcmpv6NeighborSolicitationChecksum
	HasGenerated() bool
	// Custom returns uint32, set in PatternFlowIcmpv6NeighborSolicitationChecksum.
	Custom() uint32
	// SetCustom assigns uint32 provided by user to PatternFlowIcmpv6NeighborSolicitationChecksum
	SetCustom(value uint32) PatternFlowIcmpv6NeighborSolicitationChecksum
	// HasCustom checks if Custom has been set in PatternFlowIcmpv6NeighborSolicitationChecksum
	HasCustom() bool
}

type PatternFlowIcmpv6NeighborSolicitationChecksumChoiceEnum string

// Enum of Choice on PatternFlowIcmpv6NeighborSolicitationChecksum
var PatternFlowIcmpv6NeighborSolicitationChecksumChoice = struct {
	GENERATED PatternFlowIcmpv6NeighborSolicitationChecksumChoiceEnum
	CUSTOM    PatternFlowIcmpv6NeighborSolicitationChecksumChoiceEnum
}{
	GENERATED: PatternFlowIcmpv6NeighborSolicitationChecksumChoiceEnum("generated"),
	CUSTOM:    PatternFlowIcmpv6NeighborSolicitationChecksumChoiceEnum("custom"),
}

func (obj *patternFlowIcmpv6NeighborSolicitationChecksum) Choice() PatternFlowIcmpv6NeighborSolicitationChecksumChoiceEnum {
	return PatternFlowIcmpv6NeighborSolicitationChecksumChoiceEnum(obj.obj.Choice.Enum().String())
}

// The type of checksum
// Choice returns a string
func (obj *patternFlowIcmpv6NeighborSolicitationChecksum) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpv6NeighborSolicitationChecksum) setChoice(value PatternFlowIcmpv6NeighborSolicitationChecksumChoiceEnum) PatternFlowIcmpv6NeighborSolicitationChecksum {
	intValue, ok := otg.PatternFlowIcmpv6NeighborSolicitationChecksum_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpv6NeighborSolicitationChecksumChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpv6NeighborSolicitationChecksum_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Custom = nil
	obj.obj.Generated = otg.PatternFlowIcmpv6NeighborSolicitationChecksum_Generated_unspecified.Enum()
	return obj
}

type PatternFlowIcmpv6NeighborSolicitationChecksumGeneratedEnum string

// Enum of Generated on PatternFlowIcmpv6NeighborSolicitationChecksum
var PatternFlowIcmpv6NeighborSolicitationChecksumGenerated = struct {
	GOOD PatternFlowIcmpv6NeighborSolicitationChecksumGeneratedEnum
	BAD  PatternFlowIcmpv6NeighborSolicitationChecksumGeneratedEnum
}{
	GOOD: PatternFlowIcmpv6NeighborSolicitationChecksumGeneratedEnum("good"),
	BAD:  PatternFlowIcmpv6NeighborSolicitationChecksumGeneratedEnum("bad"),
}

func (obj *patternFlowIcmpv6NeighborSolicitationChecksum) Generated() PatternFlowIcmpv6NeighborSolicitationChecksumGeneratedEnum {
	return PatternFlowIcmpv6NeighborSolicitationChecksumGeneratedEnum(obj.obj.Generated.Enum().String())
}

// A system generated checksum value
// Generated returns a string
func (obj *patternFlowIcmpv6NeighborSolicitationChecksum) HasGenerated() bool {
	return obj.obj.Generated != nil
}

func (obj *patternFlowIcmpv6NeighborSolicitationChecksum) SetGenerated(value PatternFlowIcmpv6NeighborSolicitationChecksumGeneratedEnum) PatternFlowIcmpv6NeighborSolicitationChecksum {
	intValue, ok := otg.PatternFlowIcmpv6NeighborSolicitationChecksum_Generated_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpv6NeighborSolicitationChecksumGeneratedEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpv6NeighborSolicitationChecksum_Generated_Enum(intValue)
	obj.obj.Generated = &enumValue

	return obj
}

// A custom checksum value
// Custom returns a uint32
func (obj *patternFlowIcmpv6NeighborSolicitationChecksum) Custom() uint32 {

	if obj.obj.Custom == nil {
		obj.setChoice(PatternFlowIcmpv6NeighborSolicitationChecksumChoice.CUSTOM)
	}

	return *obj.obj.Custom

}

// A custom checksum value
// Custom returns a uint32
func (obj *patternFlowIcmpv6NeighborSolicitationChecksum) HasCustom() bool {
	return obj.obj.Custom != nil
}

// A custom checksum value
// SetCustom sets the uint32 value in the PatternFlowIcmpv6NeighborSolicitationChecksum object
func (obj *patternFlowIcmpv6NeighborSolicitationChecksum) SetCustom(value uint32) PatternFlowIcmpv6NeighborSolicitationChecksum {
	obj.setChoice(PatternFlowIcmpv6NeighborSolicitationChecksumChoice.CUSTOM)
	obj.obj.Custom = &value
	return obj
}

func (obj *patternFlowIcmpv6NeighborSolicitationChecksum) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {

		if *obj.obj.Custom > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NeighborSolicitationChecksum.Custom <= 65535 but Got %d", *obj.obj.Custom))
		}

	}

}

func (obj *patternFlowIcmpv6NeighborSolicitationChecksum) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpv6NeighborSolicitationChecksumChoiceEnum

	if obj.obj.Generated != nil && obj.obj.Generated.Number() != 0 {
		choices_set += 1
		choice = PatternFlowIcmpv6NeighborSolicitationChecksumChoice.GENERATED
	}

	if obj.obj.Custom != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NeighborSolicitationChecksumChoice.CUSTOM
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpv6NeighborSolicitationChecksumChoice.GENERATED)
			if obj.obj.Generated.Number() == 0 {
				obj.SetGenerated(PatternFlowIcmpv6NeighborSolicitationChecksumGenerated.GOOD)

			}

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpv6NeighborSolicitationChecksum")
			}
		} else {
			intVal := otg.PatternFlowIcmpv6NeighborSolicitationChecksum_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpv6NeighborSolicitationChecksum_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
