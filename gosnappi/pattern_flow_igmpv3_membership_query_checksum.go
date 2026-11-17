package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3MembershipQueryChecksum *****
type patternFlowIgmpv3MembershipQueryChecksum struct {
	validation
	obj          *otg.PatternFlowIgmpv3MembershipQueryChecksum
	marshaller   marshalPatternFlowIgmpv3MembershipQueryChecksum
	unMarshaller unMarshalPatternFlowIgmpv3MembershipQueryChecksum
}

func NewPatternFlowIgmpv3MembershipQueryChecksum() PatternFlowIgmpv3MembershipQueryChecksum {
	obj := patternFlowIgmpv3MembershipQueryChecksum{obj: &otg.PatternFlowIgmpv3MembershipQueryChecksum{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3MembershipQueryChecksum) msg() *otg.PatternFlowIgmpv3MembershipQueryChecksum {
	return obj.obj
}

func (obj *patternFlowIgmpv3MembershipQueryChecksum) setMsg(msg *otg.PatternFlowIgmpv3MembershipQueryChecksum) PatternFlowIgmpv3MembershipQueryChecksum {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3MembershipQueryChecksum struct {
	obj *patternFlowIgmpv3MembershipQueryChecksum
}

type marshalPatternFlowIgmpv3MembershipQueryChecksum interface {
	// ToProto marshals PatternFlowIgmpv3MembershipQueryChecksum to protobuf object *otg.PatternFlowIgmpv3MembershipQueryChecksum
	ToProto() (*otg.PatternFlowIgmpv3MembershipQueryChecksum, error)
	// ToPbText marshals PatternFlowIgmpv3MembershipQueryChecksum to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3MembershipQueryChecksum to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3MembershipQueryChecksum to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3MembershipQueryChecksum struct {
	obj *patternFlowIgmpv3MembershipQueryChecksum
}

type unMarshalPatternFlowIgmpv3MembershipQueryChecksum interface {
	// FromProto unmarshals PatternFlowIgmpv3MembershipQueryChecksum from protobuf object *otg.PatternFlowIgmpv3MembershipQueryChecksum
	FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryChecksum) (PatternFlowIgmpv3MembershipQueryChecksum, error)
	// FromPbText unmarshals PatternFlowIgmpv3MembershipQueryChecksum from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3MembershipQueryChecksum from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3MembershipQueryChecksum from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3MembershipQueryChecksum) Marshal() marshalPatternFlowIgmpv3MembershipQueryChecksum {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3MembershipQueryChecksum{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3MembershipQueryChecksum) Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryChecksum {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3MembershipQueryChecksum{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3MembershipQueryChecksum) ToProto() (*otg.PatternFlowIgmpv3MembershipQueryChecksum, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3MembershipQueryChecksum) FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryChecksum) (PatternFlowIgmpv3MembershipQueryChecksum, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3MembershipQueryChecksum) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryChecksum) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryChecksum) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryChecksum) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryChecksum) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryChecksum) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3MembershipQueryChecksum) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryChecksum) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryChecksum) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3MembershipQueryChecksum) Clone() (PatternFlowIgmpv3MembershipQueryChecksum, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3MembershipQueryChecksum()
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

// PatternFlowIgmpv3MembershipQueryChecksum is checksum of the whole IGMP message, a 16-bit one's complement sum with no pseudo-header.
type PatternFlowIgmpv3MembershipQueryChecksum interface {
	Validation
	// msg marshals PatternFlowIgmpv3MembershipQueryChecksum to protobuf object *otg.PatternFlowIgmpv3MembershipQueryChecksum
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3MembershipQueryChecksum
	// setMsg unmarshals PatternFlowIgmpv3MembershipQueryChecksum from protobuf object *otg.PatternFlowIgmpv3MembershipQueryChecksum
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3MembershipQueryChecksum) PatternFlowIgmpv3MembershipQueryChecksum
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3MembershipQueryChecksum
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryChecksum
	// validate validates PatternFlowIgmpv3MembershipQueryChecksum
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3MembershipQueryChecksum, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIgmpv3MembershipQueryChecksumChoiceEnum, set in PatternFlowIgmpv3MembershipQueryChecksum
	Choice() PatternFlowIgmpv3MembershipQueryChecksumChoiceEnum
	// setChoice assigns PatternFlowIgmpv3MembershipQueryChecksumChoiceEnum provided by user to PatternFlowIgmpv3MembershipQueryChecksum
	setChoice(value PatternFlowIgmpv3MembershipQueryChecksumChoiceEnum) PatternFlowIgmpv3MembershipQueryChecksum
	// HasChoice checks if Choice has been set in PatternFlowIgmpv3MembershipQueryChecksum
	HasChoice() bool
	// Generated returns PatternFlowIgmpv3MembershipQueryChecksumGeneratedEnum, set in PatternFlowIgmpv3MembershipQueryChecksum
	Generated() PatternFlowIgmpv3MembershipQueryChecksumGeneratedEnum
	// SetGenerated assigns PatternFlowIgmpv3MembershipQueryChecksumGeneratedEnum provided by user to PatternFlowIgmpv3MembershipQueryChecksum
	SetGenerated(value PatternFlowIgmpv3MembershipQueryChecksumGeneratedEnum) PatternFlowIgmpv3MembershipQueryChecksum
	// HasGenerated checks if Generated has been set in PatternFlowIgmpv3MembershipQueryChecksum
	HasGenerated() bool
	// Custom returns uint32, set in PatternFlowIgmpv3MembershipQueryChecksum.
	Custom() uint32
	// SetCustom assigns uint32 provided by user to PatternFlowIgmpv3MembershipQueryChecksum
	SetCustom(value uint32) PatternFlowIgmpv3MembershipQueryChecksum
	// HasCustom checks if Custom has been set in PatternFlowIgmpv3MembershipQueryChecksum
	HasCustom() bool
}

type PatternFlowIgmpv3MembershipQueryChecksumChoiceEnum string

// Enum of Choice on PatternFlowIgmpv3MembershipQueryChecksum
var PatternFlowIgmpv3MembershipQueryChecksumChoice = struct {
	GENERATED PatternFlowIgmpv3MembershipQueryChecksumChoiceEnum
	CUSTOM    PatternFlowIgmpv3MembershipQueryChecksumChoiceEnum
}{
	GENERATED: PatternFlowIgmpv3MembershipQueryChecksumChoiceEnum("generated"),
	CUSTOM:    PatternFlowIgmpv3MembershipQueryChecksumChoiceEnum("custom"),
}

func (obj *patternFlowIgmpv3MembershipQueryChecksum) Choice() PatternFlowIgmpv3MembershipQueryChecksumChoiceEnum {
	return PatternFlowIgmpv3MembershipQueryChecksumChoiceEnum(obj.obj.Choice.Enum().String())
}

// The type of checksum
// Choice returns a string
func (obj *patternFlowIgmpv3MembershipQueryChecksum) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIgmpv3MembershipQueryChecksum) setChoice(value PatternFlowIgmpv3MembershipQueryChecksumChoiceEnum) PatternFlowIgmpv3MembershipQueryChecksum {
	intValue, ok := otg.PatternFlowIgmpv3MembershipQueryChecksum_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv3MembershipQueryChecksumChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv3MembershipQueryChecksum_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Custom = nil
	obj.obj.Generated = otg.PatternFlowIgmpv3MembershipQueryChecksum_Generated_unspecified.Enum()
	return obj
}

type PatternFlowIgmpv3MembershipQueryChecksumGeneratedEnum string

// Enum of Generated on PatternFlowIgmpv3MembershipQueryChecksum
var PatternFlowIgmpv3MembershipQueryChecksumGenerated = struct {
	GOOD PatternFlowIgmpv3MembershipQueryChecksumGeneratedEnum
	BAD  PatternFlowIgmpv3MembershipQueryChecksumGeneratedEnum
}{
	GOOD: PatternFlowIgmpv3MembershipQueryChecksumGeneratedEnum("good"),
	BAD:  PatternFlowIgmpv3MembershipQueryChecksumGeneratedEnum("bad"),
}

func (obj *patternFlowIgmpv3MembershipQueryChecksum) Generated() PatternFlowIgmpv3MembershipQueryChecksumGeneratedEnum {
	return PatternFlowIgmpv3MembershipQueryChecksumGeneratedEnum(obj.obj.Generated.Enum().String())
}

// A system generated checksum value
// Generated returns a string
func (obj *patternFlowIgmpv3MembershipQueryChecksum) HasGenerated() bool {
	return obj.obj.Generated != nil
}

func (obj *patternFlowIgmpv3MembershipQueryChecksum) SetGenerated(value PatternFlowIgmpv3MembershipQueryChecksumGeneratedEnum) PatternFlowIgmpv3MembershipQueryChecksum {
	intValue, ok := otg.PatternFlowIgmpv3MembershipQueryChecksum_Generated_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv3MembershipQueryChecksumGeneratedEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv3MembershipQueryChecksum_Generated_Enum(intValue)
	obj.obj.Generated = &enumValue

	return obj
}

// A custom checksum value
// Custom returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryChecksum) Custom() uint32 {

	if obj.obj.Custom == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipQueryChecksumChoice.CUSTOM)
	}

	return *obj.obj.Custom

}

// A custom checksum value
// Custom returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryChecksum) HasCustom() bool {
	return obj.obj.Custom != nil
}

// A custom checksum value
// SetCustom sets the uint32 value in the PatternFlowIgmpv3MembershipQueryChecksum object
func (obj *patternFlowIgmpv3MembershipQueryChecksum) SetCustom(value uint32) PatternFlowIgmpv3MembershipQueryChecksum {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryChecksumChoice.CUSTOM)
	obj.obj.Custom = &value
	return obj
}

func (obj *patternFlowIgmpv3MembershipQueryChecksum) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {

		if *obj.obj.Custom > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQueryChecksum.Custom <= 65535 but Got %d", *obj.obj.Custom))
		}

	}

}

func (obj *patternFlowIgmpv3MembershipQueryChecksum) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIgmpv3MembershipQueryChecksumChoiceEnum

	if obj.obj.Generated != nil && obj.obj.Generated.Number() != 0 {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryChecksumChoice.GENERATED
	}

	if obj.obj.Custom != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryChecksumChoice.CUSTOM
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIgmpv3MembershipQueryChecksumChoice.GENERATED)
			if obj.obj.Generated.Number() == 0 {
				obj.SetGenerated(PatternFlowIgmpv3MembershipQueryChecksumGenerated.GOOD)

			}

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIgmpv3MembershipQueryChecksum")
			}
		} else {
			intVal := otg.PatternFlowIgmpv3MembershipQueryChecksum_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIgmpv3MembershipQueryChecksum_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
