package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3MembershipReportChecksum *****
type patternFlowIgmpv3MembershipReportChecksum struct {
	validation
	obj          *otg.PatternFlowIgmpv3MembershipReportChecksum
	marshaller   marshalPatternFlowIgmpv3MembershipReportChecksum
	unMarshaller unMarshalPatternFlowIgmpv3MembershipReportChecksum
}

func NewPatternFlowIgmpv3MembershipReportChecksum() PatternFlowIgmpv3MembershipReportChecksum {
	obj := patternFlowIgmpv3MembershipReportChecksum{obj: &otg.PatternFlowIgmpv3MembershipReportChecksum{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3MembershipReportChecksum) msg() *otg.PatternFlowIgmpv3MembershipReportChecksum {
	return obj.obj
}

func (obj *patternFlowIgmpv3MembershipReportChecksum) setMsg(msg *otg.PatternFlowIgmpv3MembershipReportChecksum) PatternFlowIgmpv3MembershipReportChecksum {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3MembershipReportChecksum struct {
	obj *patternFlowIgmpv3MembershipReportChecksum
}

type marshalPatternFlowIgmpv3MembershipReportChecksum interface {
	// ToProto marshals PatternFlowIgmpv3MembershipReportChecksum to protobuf object *otg.PatternFlowIgmpv3MembershipReportChecksum
	ToProto() (*otg.PatternFlowIgmpv3MembershipReportChecksum, error)
	// ToPbText marshals PatternFlowIgmpv3MembershipReportChecksum to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3MembershipReportChecksum to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3MembershipReportChecksum to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3MembershipReportChecksum struct {
	obj *patternFlowIgmpv3MembershipReportChecksum
}

type unMarshalPatternFlowIgmpv3MembershipReportChecksum interface {
	// FromProto unmarshals PatternFlowIgmpv3MembershipReportChecksum from protobuf object *otg.PatternFlowIgmpv3MembershipReportChecksum
	FromProto(msg *otg.PatternFlowIgmpv3MembershipReportChecksum) (PatternFlowIgmpv3MembershipReportChecksum, error)
	// FromPbText unmarshals PatternFlowIgmpv3MembershipReportChecksum from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3MembershipReportChecksum from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3MembershipReportChecksum from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3MembershipReportChecksum) Marshal() marshalPatternFlowIgmpv3MembershipReportChecksum {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3MembershipReportChecksum{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3MembershipReportChecksum) Unmarshal() unMarshalPatternFlowIgmpv3MembershipReportChecksum {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3MembershipReportChecksum{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3MembershipReportChecksum) ToProto() (*otg.PatternFlowIgmpv3MembershipReportChecksum, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3MembershipReportChecksum) FromProto(msg *otg.PatternFlowIgmpv3MembershipReportChecksum) (PatternFlowIgmpv3MembershipReportChecksum, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3MembershipReportChecksum) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipReportChecksum) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipReportChecksum) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipReportChecksum) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipReportChecksum) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipReportChecksum) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3MembershipReportChecksum) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipReportChecksum) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipReportChecksum) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3MembershipReportChecksum) Clone() (PatternFlowIgmpv3MembershipReportChecksum, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3MembershipReportChecksum()
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

// PatternFlowIgmpv3MembershipReportChecksum is checksum of the whole IGMP message, a 16-bit one's complement sum with no pseudo-header.
type PatternFlowIgmpv3MembershipReportChecksum interface {
	Validation
	// msg marshals PatternFlowIgmpv3MembershipReportChecksum to protobuf object *otg.PatternFlowIgmpv3MembershipReportChecksum
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3MembershipReportChecksum
	// setMsg unmarshals PatternFlowIgmpv3MembershipReportChecksum from protobuf object *otg.PatternFlowIgmpv3MembershipReportChecksum
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3MembershipReportChecksum) PatternFlowIgmpv3MembershipReportChecksum
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3MembershipReportChecksum
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3MembershipReportChecksum
	// validate validates PatternFlowIgmpv3MembershipReportChecksum
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3MembershipReportChecksum, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIgmpv3MembershipReportChecksumChoiceEnum, set in PatternFlowIgmpv3MembershipReportChecksum
	Choice() PatternFlowIgmpv3MembershipReportChecksumChoiceEnum
	// setChoice assigns PatternFlowIgmpv3MembershipReportChecksumChoiceEnum provided by user to PatternFlowIgmpv3MembershipReportChecksum
	setChoice(value PatternFlowIgmpv3MembershipReportChecksumChoiceEnum) PatternFlowIgmpv3MembershipReportChecksum
	// HasChoice checks if Choice has been set in PatternFlowIgmpv3MembershipReportChecksum
	HasChoice() bool
	// Generated returns PatternFlowIgmpv3MembershipReportChecksumGeneratedEnum, set in PatternFlowIgmpv3MembershipReportChecksum
	Generated() PatternFlowIgmpv3MembershipReportChecksumGeneratedEnum
	// SetGenerated assigns PatternFlowIgmpv3MembershipReportChecksumGeneratedEnum provided by user to PatternFlowIgmpv3MembershipReportChecksum
	SetGenerated(value PatternFlowIgmpv3MembershipReportChecksumGeneratedEnum) PatternFlowIgmpv3MembershipReportChecksum
	// HasGenerated checks if Generated has been set in PatternFlowIgmpv3MembershipReportChecksum
	HasGenerated() bool
	// Custom returns uint32, set in PatternFlowIgmpv3MembershipReportChecksum.
	Custom() uint32
	// SetCustom assigns uint32 provided by user to PatternFlowIgmpv3MembershipReportChecksum
	SetCustom(value uint32) PatternFlowIgmpv3MembershipReportChecksum
	// HasCustom checks if Custom has been set in PatternFlowIgmpv3MembershipReportChecksum
	HasCustom() bool
}

type PatternFlowIgmpv3MembershipReportChecksumChoiceEnum string

// Enum of Choice on PatternFlowIgmpv3MembershipReportChecksum
var PatternFlowIgmpv3MembershipReportChecksumChoice = struct {
	GENERATED PatternFlowIgmpv3MembershipReportChecksumChoiceEnum
	CUSTOM    PatternFlowIgmpv3MembershipReportChecksumChoiceEnum
}{
	GENERATED: PatternFlowIgmpv3MembershipReportChecksumChoiceEnum("generated"),
	CUSTOM:    PatternFlowIgmpv3MembershipReportChecksumChoiceEnum("custom"),
}

func (obj *patternFlowIgmpv3MembershipReportChecksum) Choice() PatternFlowIgmpv3MembershipReportChecksumChoiceEnum {
	return PatternFlowIgmpv3MembershipReportChecksumChoiceEnum(obj.obj.Choice.Enum().String())
}

// The type of checksum
// Choice returns a string
func (obj *patternFlowIgmpv3MembershipReportChecksum) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIgmpv3MembershipReportChecksum) setChoice(value PatternFlowIgmpv3MembershipReportChecksumChoiceEnum) PatternFlowIgmpv3MembershipReportChecksum {
	intValue, ok := otg.PatternFlowIgmpv3MembershipReportChecksum_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv3MembershipReportChecksumChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv3MembershipReportChecksum_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Custom = nil
	obj.obj.Generated = otg.PatternFlowIgmpv3MembershipReportChecksum_Generated_unspecified.Enum()
	return obj
}

type PatternFlowIgmpv3MembershipReportChecksumGeneratedEnum string

// Enum of Generated on PatternFlowIgmpv3MembershipReportChecksum
var PatternFlowIgmpv3MembershipReportChecksumGenerated = struct {
	GOOD PatternFlowIgmpv3MembershipReportChecksumGeneratedEnum
	BAD  PatternFlowIgmpv3MembershipReportChecksumGeneratedEnum
}{
	GOOD: PatternFlowIgmpv3MembershipReportChecksumGeneratedEnum("good"),
	BAD:  PatternFlowIgmpv3MembershipReportChecksumGeneratedEnum("bad"),
}

func (obj *patternFlowIgmpv3MembershipReportChecksum) Generated() PatternFlowIgmpv3MembershipReportChecksumGeneratedEnum {
	return PatternFlowIgmpv3MembershipReportChecksumGeneratedEnum(obj.obj.Generated.Enum().String())
}

// A system generated checksum value
// Generated returns a string
func (obj *patternFlowIgmpv3MembershipReportChecksum) HasGenerated() bool {
	return obj.obj.Generated != nil
}

func (obj *patternFlowIgmpv3MembershipReportChecksum) SetGenerated(value PatternFlowIgmpv3MembershipReportChecksumGeneratedEnum) PatternFlowIgmpv3MembershipReportChecksum {
	intValue, ok := otg.PatternFlowIgmpv3MembershipReportChecksum_Generated_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv3MembershipReportChecksumGeneratedEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv3MembershipReportChecksum_Generated_Enum(intValue)
	obj.obj.Generated = &enumValue

	return obj
}

// A custom checksum value
// Custom returns a uint32
func (obj *patternFlowIgmpv3MembershipReportChecksum) Custom() uint32 {

	if obj.obj.Custom == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipReportChecksumChoice.CUSTOM)
	}

	return *obj.obj.Custom

}

// A custom checksum value
// Custom returns a uint32
func (obj *patternFlowIgmpv3MembershipReportChecksum) HasCustom() bool {
	return obj.obj.Custom != nil
}

// A custom checksum value
// SetCustom sets the uint32 value in the PatternFlowIgmpv3MembershipReportChecksum object
func (obj *patternFlowIgmpv3MembershipReportChecksum) SetCustom(value uint32) PatternFlowIgmpv3MembershipReportChecksum {
	obj.setChoice(PatternFlowIgmpv3MembershipReportChecksumChoice.CUSTOM)
	obj.obj.Custom = &value
	return obj
}

func (obj *patternFlowIgmpv3MembershipReportChecksum) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {

		if *obj.obj.Custom > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipReportChecksum.Custom <= 65535 but Got %d", *obj.obj.Custom))
		}

	}

}

func (obj *patternFlowIgmpv3MembershipReportChecksum) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIgmpv3MembershipReportChecksumChoiceEnum

	if obj.obj.Generated != nil && obj.obj.Generated.Number() != 0 {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipReportChecksumChoice.GENERATED
	}

	if obj.obj.Custom != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipReportChecksumChoice.CUSTOM
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIgmpv3MembershipReportChecksumChoice.GENERATED)
			if obj.obj.Generated.Number() == 0 {
				obj.SetGenerated(PatternFlowIgmpv3MembershipReportChecksumGenerated.GOOD)

			}

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIgmpv3MembershipReportChecksum")
			}
		} else {
			intVal := otg.PatternFlowIgmpv3MembershipReportChecksum_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIgmpv3MembershipReportChecksum_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
