package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4HeaderChecksum *****
type patternFlowIpv4HeaderChecksum struct {
	validation
	obj          *otg.PatternFlowIpv4HeaderChecksum
	marshaller   marshalPatternFlowIpv4HeaderChecksum
	unMarshaller unMarshalPatternFlowIpv4HeaderChecksum
}

func NewPatternFlowIpv4HeaderChecksum() PatternFlowIpv4HeaderChecksum {
	obj := patternFlowIpv4HeaderChecksum{obj: &otg.PatternFlowIpv4HeaderChecksum{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4HeaderChecksum) msg() *otg.PatternFlowIpv4HeaderChecksum {
	return obj.obj
}

func (obj *patternFlowIpv4HeaderChecksum) setMsg(msg *otg.PatternFlowIpv4HeaderChecksum) PatternFlowIpv4HeaderChecksum {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4HeaderChecksum struct {
	obj *patternFlowIpv4HeaderChecksum
}

type marshalPatternFlowIpv4HeaderChecksum interface {
	// ToProto marshals PatternFlowIpv4HeaderChecksum to protobuf object *otg.PatternFlowIpv4HeaderChecksum
	ToProto() (*otg.PatternFlowIpv4HeaderChecksum, error)
	// ToPbText marshals PatternFlowIpv4HeaderChecksum to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4HeaderChecksum to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4HeaderChecksum to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv4HeaderChecksum to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv4HeaderChecksum struct {
	obj *patternFlowIpv4HeaderChecksum
}

type unMarshalPatternFlowIpv4HeaderChecksum interface {
	// FromProto unmarshals PatternFlowIpv4HeaderChecksum from protobuf object *otg.PatternFlowIpv4HeaderChecksum
	FromProto(msg *otg.PatternFlowIpv4HeaderChecksum) (PatternFlowIpv4HeaderChecksum, error)
	// FromPbText unmarshals PatternFlowIpv4HeaderChecksum from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4HeaderChecksum from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4HeaderChecksum from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4HeaderChecksum) Marshal() marshalPatternFlowIpv4HeaderChecksum {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4HeaderChecksum{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4HeaderChecksum) Unmarshal() unMarshalPatternFlowIpv4HeaderChecksum {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4HeaderChecksum{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4HeaderChecksum) ToProto() (*otg.PatternFlowIpv4HeaderChecksum, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4HeaderChecksum) FromProto(msg *otg.PatternFlowIpv4HeaderChecksum) (PatternFlowIpv4HeaderChecksum, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4HeaderChecksum) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4HeaderChecksum) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4HeaderChecksum) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4HeaderChecksum) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4HeaderChecksum) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv4HeaderChecksum) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4HeaderChecksum) FromJson(value string) error {
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

func (obj *patternFlowIpv4HeaderChecksum) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4HeaderChecksum) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4HeaderChecksum) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4HeaderChecksum) Clone() (PatternFlowIpv4HeaderChecksum, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4HeaderChecksum()
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

// PatternFlowIpv4HeaderChecksum is header checksum
type PatternFlowIpv4HeaderChecksum interface {
	Validation
	// msg marshals PatternFlowIpv4HeaderChecksum to protobuf object *otg.PatternFlowIpv4HeaderChecksum
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4HeaderChecksum
	// setMsg unmarshals PatternFlowIpv4HeaderChecksum from protobuf object *otg.PatternFlowIpv4HeaderChecksum
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4HeaderChecksum) PatternFlowIpv4HeaderChecksum
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4HeaderChecksum
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4HeaderChecksum
	// validate validates PatternFlowIpv4HeaderChecksum
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4HeaderChecksum, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4HeaderChecksumChoiceEnum, set in PatternFlowIpv4HeaderChecksum
	Choice() PatternFlowIpv4HeaderChecksumChoiceEnum
	// setChoice assigns PatternFlowIpv4HeaderChecksumChoiceEnum provided by user to PatternFlowIpv4HeaderChecksum
	setChoice(value PatternFlowIpv4HeaderChecksumChoiceEnum) PatternFlowIpv4HeaderChecksum
	// HasChoice checks if Choice has been set in PatternFlowIpv4HeaderChecksum
	HasChoice() bool
	// Generated returns PatternFlowIpv4HeaderChecksumGeneratedEnum, set in PatternFlowIpv4HeaderChecksum
	Generated() PatternFlowIpv4HeaderChecksumGeneratedEnum
	// SetGenerated assigns PatternFlowIpv4HeaderChecksumGeneratedEnum provided by user to PatternFlowIpv4HeaderChecksum
	SetGenerated(value PatternFlowIpv4HeaderChecksumGeneratedEnum) PatternFlowIpv4HeaderChecksum
	// HasGenerated checks if Generated has been set in PatternFlowIpv4HeaderChecksum
	HasGenerated() bool
	// Custom returns uint32, set in PatternFlowIpv4HeaderChecksum.
	Custom() uint32
	// SetCustom assigns uint32 provided by user to PatternFlowIpv4HeaderChecksum
	SetCustom(value uint32) PatternFlowIpv4HeaderChecksum
	// HasCustom checks if Custom has been set in PatternFlowIpv4HeaderChecksum
	HasCustom() bool
}

type PatternFlowIpv4HeaderChecksumChoiceEnum string

// Enum of Choice on PatternFlowIpv4HeaderChecksum
var PatternFlowIpv4HeaderChecksumChoice = struct {
	GENERATED PatternFlowIpv4HeaderChecksumChoiceEnum
	CUSTOM    PatternFlowIpv4HeaderChecksumChoiceEnum
}{
	GENERATED: PatternFlowIpv4HeaderChecksumChoiceEnum("generated"),
	CUSTOM:    PatternFlowIpv4HeaderChecksumChoiceEnum("custom"),
}

func (obj *patternFlowIpv4HeaderChecksum) Choice() PatternFlowIpv4HeaderChecksumChoiceEnum {
	return PatternFlowIpv4HeaderChecksumChoiceEnum(obj.obj.Choice.Enum().String())
}

// The type of checksum
// Choice returns a string
func (obj *patternFlowIpv4HeaderChecksum) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4HeaderChecksum) setChoice(value PatternFlowIpv4HeaderChecksumChoiceEnum) PatternFlowIpv4HeaderChecksum {
	intValue, ok := otg.PatternFlowIpv4HeaderChecksum_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4HeaderChecksumChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4HeaderChecksum_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Custom = nil
	obj.obj.Generated = otg.PatternFlowIpv4HeaderChecksum_Generated_unspecified.Enum()
	return obj
}

type PatternFlowIpv4HeaderChecksumGeneratedEnum string

// Enum of Generated on PatternFlowIpv4HeaderChecksum
var PatternFlowIpv4HeaderChecksumGenerated = struct {
	GOOD PatternFlowIpv4HeaderChecksumGeneratedEnum
	BAD  PatternFlowIpv4HeaderChecksumGeneratedEnum
}{
	GOOD: PatternFlowIpv4HeaderChecksumGeneratedEnum("good"),
	BAD:  PatternFlowIpv4HeaderChecksumGeneratedEnum("bad"),
}

func (obj *patternFlowIpv4HeaderChecksum) Generated() PatternFlowIpv4HeaderChecksumGeneratedEnum {
	return PatternFlowIpv4HeaderChecksumGeneratedEnum(obj.obj.Generated.Enum().String())
}

// A system generated checksum value
// Generated returns a string
func (obj *patternFlowIpv4HeaderChecksum) HasGenerated() bool {
	return obj.obj.Generated != nil
}

func (obj *patternFlowIpv4HeaderChecksum) SetGenerated(value PatternFlowIpv4HeaderChecksumGeneratedEnum) PatternFlowIpv4HeaderChecksum {
	intValue, ok := otg.PatternFlowIpv4HeaderChecksum_Generated_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4HeaderChecksumGeneratedEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4HeaderChecksum_Generated_Enum(intValue)
	obj.obj.Generated = &enumValue

	return obj
}

// A custom checksum value
// Custom returns a uint32
func (obj *patternFlowIpv4HeaderChecksum) Custom() uint32 {

	if obj.obj.Custom == nil {
		obj.setChoice(PatternFlowIpv4HeaderChecksumChoice.CUSTOM)
	}

	return *obj.obj.Custom

}

// A custom checksum value
// Custom returns a uint32
func (obj *patternFlowIpv4HeaderChecksum) HasCustom() bool {
	return obj.obj.Custom != nil
}

// A custom checksum value
// SetCustom sets the uint32 value in the PatternFlowIpv4HeaderChecksum object
func (obj *patternFlowIpv4HeaderChecksum) SetCustom(value uint32) PatternFlowIpv4HeaderChecksum {
	obj.setChoice(PatternFlowIpv4HeaderChecksumChoice.CUSTOM)
	obj.obj.Custom = &value
	return obj
}

func (obj *patternFlowIpv4HeaderChecksum) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {

		if *obj.obj.Custom > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4HeaderChecksum.Custom <= 65535 but Got %d", *obj.obj.Custom))
		}

	}

}

func (obj *patternFlowIpv4HeaderChecksum) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv4HeaderChecksumChoiceEnum

	if obj.obj.Generated != nil && obj.obj.Generated.Number() != 0 {
		choices_set += 1
		choice = PatternFlowIpv4HeaderChecksumChoice.GENERATED
	}

	if obj.obj.Custom != nil {
		choices_set += 1
		choice = PatternFlowIpv4HeaderChecksumChoice.CUSTOM
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv4HeaderChecksumChoice.GENERATED)
			if obj.obj.Generated.Number() == 0 {
				obj.SetGenerated(PatternFlowIpv4HeaderChecksumGenerated.GOOD)

			}

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv4HeaderChecksum")
			}
		} else {
			intVal := otg.PatternFlowIpv4HeaderChecksum_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv4HeaderChecksum_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
