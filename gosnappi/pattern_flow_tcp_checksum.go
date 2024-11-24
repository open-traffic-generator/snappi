package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpChecksum *****
type patternFlowTcpChecksum struct {
	validation
	obj          *otg.PatternFlowTcpChecksum
	marshaller   marshalPatternFlowTcpChecksum
	unMarshaller unMarshalPatternFlowTcpChecksum
}

func NewPatternFlowTcpChecksum() PatternFlowTcpChecksum {
	obj := patternFlowTcpChecksum{obj: &otg.PatternFlowTcpChecksum{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpChecksum) msg() *otg.PatternFlowTcpChecksum {
	return obj.obj
}

func (obj *patternFlowTcpChecksum) setMsg(msg *otg.PatternFlowTcpChecksum) PatternFlowTcpChecksum {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpChecksum struct {
	obj *patternFlowTcpChecksum
}

type marshalPatternFlowTcpChecksum interface {
	// ToProto marshals PatternFlowTcpChecksum to protobuf object *otg.PatternFlowTcpChecksum
	ToProto() (*otg.PatternFlowTcpChecksum, error)
	// ToPbText marshals PatternFlowTcpChecksum to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpChecksum to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpChecksum to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowTcpChecksum to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowTcpChecksum struct {
	obj *patternFlowTcpChecksum
}

type unMarshalPatternFlowTcpChecksum interface {
	// FromProto unmarshals PatternFlowTcpChecksum from protobuf object *otg.PatternFlowTcpChecksum
	FromProto(msg *otg.PatternFlowTcpChecksum) (PatternFlowTcpChecksum, error)
	// FromPbText unmarshals PatternFlowTcpChecksum from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpChecksum from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpChecksum from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpChecksum) Marshal() marshalPatternFlowTcpChecksum {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpChecksum{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpChecksum) Unmarshal() unMarshalPatternFlowTcpChecksum {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpChecksum{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpChecksum) ToProto() (*otg.PatternFlowTcpChecksum, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpChecksum) FromProto(msg *otg.PatternFlowTcpChecksum) (PatternFlowTcpChecksum, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpChecksum) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpChecksum) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpChecksum) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpChecksum) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpChecksum) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowTcpChecksum) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpChecksum) FromJson(value string) error {
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

func (obj *patternFlowTcpChecksum) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpChecksum) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpChecksum) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpChecksum) Clone() (PatternFlowTcpChecksum, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpChecksum()
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

// PatternFlowTcpChecksum is the one's complement of the one's complement sum of all 16 bit words in header and text.  An all-zero value means that no checksum will be transmitted.   While computing the checksum, the checksum field itself is replaced with zeros.
type PatternFlowTcpChecksum interface {
	Validation
	// msg marshals PatternFlowTcpChecksum to protobuf object *otg.PatternFlowTcpChecksum
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpChecksum
	// setMsg unmarshals PatternFlowTcpChecksum from protobuf object *otg.PatternFlowTcpChecksum
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpChecksum) PatternFlowTcpChecksum
	// provides marshal interface
	Marshal() marshalPatternFlowTcpChecksum
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpChecksum
	// validate validates PatternFlowTcpChecksum
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpChecksum, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowTcpChecksumChoiceEnum, set in PatternFlowTcpChecksum
	Choice() PatternFlowTcpChecksumChoiceEnum
	// setChoice assigns PatternFlowTcpChecksumChoiceEnum provided by user to PatternFlowTcpChecksum
	setChoice(value PatternFlowTcpChecksumChoiceEnum) PatternFlowTcpChecksum
	// HasChoice checks if Choice has been set in PatternFlowTcpChecksum
	HasChoice() bool
	// Generated returns PatternFlowTcpChecksumGeneratedEnum, set in PatternFlowTcpChecksum
	Generated() PatternFlowTcpChecksumGeneratedEnum
	// SetGenerated assigns PatternFlowTcpChecksumGeneratedEnum provided by user to PatternFlowTcpChecksum
	SetGenerated(value PatternFlowTcpChecksumGeneratedEnum) PatternFlowTcpChecksum
	// HasGenerated checks if Generated has been set in PatternFlowTcpChecksum
	HasGenerated() bool
	// Custom returns uint32, set in PatternFlowTcpChecksum.
	Custom() uint32
	// SetCustom assigns uint32 provided by user to PatternFlowTcpChecksum
	SetCustom(value uint32) PatternFlowTcpChecksum
	// HasCustom checks if Custom has been set in PatternFlowTcpChecksum
	HasCustom() bool
}

type PatternFlowTcpChecksumChoiceEnum string

// Enum of Choice on PatternFlowTcpChecksum
var PatternFlowTcpChecksumChoice = struct {
	GENERATED PatternFlowTcpChecksumChoiceEnum
	CUSTOM    PatternFlowTcpChecksumChoiceEnum
}{
	GENERATED: PatternFlowTcpChecksumChoiceEnum("generated"),
	CUSTOM:    PatternFlowTcpChecksumChoiceEnum("custom"),
}

func (obj *patternFlowTcpChecksum) Choice() PatternFlowTcpChecksumChoiceEnum {
	return PatternFlowTcpChecksumChoiceEnum(obj.obj.Choice.Enum().String())
}

// The type of checksum
// Choice returns a string
func (obj *patternFlowTcpChecksum) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowTcpChecksum) setChoice(value PatternFlowTcpChecksumChoiceEnum) PatternFlowTcpChecksum {
	intValue, ok := otg.PatternFlowTcpChecksum_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowTcpChecksumChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowTcpChecksum_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Custom = nil
	obj.obj.Generated = otg.PatternFlowTcpChecksum_Generated_unspecified.Enum()
	return obj
}

type PatternFlowTcpChecksumGeneratedEnum string

// Enum of Generated on PatternFlowTcpChecksum
var PatternFlowTcpChecksumGenerated = struct {
	GOOD PatternFlowTcpChecksumGeneratedEnum
	BAD  PatternFlowTcpChecksumGeneratedEnum
}{
	GOOD: PatternFlowTcpChecksumGeneratedEnum("good"),
	BAD:  PatternFlowTcpChecksumGeneratedEnum("bad"),
}

func (obj *patternFlowTcpChecksum) Generated() PatternFlowTcpChecksumGeneratedEnum {
	return PatternFlowTcpChecksumGeneratedEnum(obj.obj.Generated.Enum().String())
}

// A system generated checksum value
// Generated returns a string
func (obj *patternFlowTcpChecksum) HasGenerated() bool {
	return obj.obj.Generated != nil
}

func (obj *patternFlowTcpChecksum) SetGenerated(value PatternFlowTcpChecksumGeneratedEnum) PatternFlowTcpChecksum {
	intValue, ok := otg.PatternFlowTcpChecksum_Generated_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowTcpChecksumGeneratedEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowTcpChecksum_Generated_Enum(intValue)
	obj.obj.Generated = &enumValue

	return obj
}

// A custom checksum value
// Custom returns a uint32
func (obj *patternFlowTcpChecksum) Custom() uint32 {

	if obj.obj.Custom == nil {
		obj.setChoice(PatternFlowTcpChecksumChoice.CUSTOM)
	}

	return *obj.obj.Custom

}

// A custom checksum value
// Custom returns a uint32
func (obj *patternFlowTcpChecksum) HasCustom() bool {
	return obj.obj.Custom != nil
}

// A custom checksum value
// SetCustom sets the uint32 value in the PatternFlowTcpChecksum object
func (obj *patternFlowTcpChecksum) SetCustom(value uint32) PatternFlowTcpChecksum {
	obj.setChoice(PatternFlowTcpChecksumChoice.CUSTOM)
	obj.obj.Custom = &value
	return obj
}

func (obj *patternFlowTcpChecksum) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {

		if *obj.obj.Custom > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpChecksum.Custom <= 65535 but Got %d", *obj.obj.Custom))
		}

	}

}

func (obj *patternFlowTcpChecksum) setDefault() {
	var choices_set int = 0
	var choice PatternFlowTcpChecksumChoiceEnum

	if obj.obj.Generated != nil && obj.obj.Generated.Number() != 0 {
		choices_set += 1
		choice = PatternFlowTcpChecksumChoice.GENERATED
	}

	if obj.obj.Custom != nil {
		choices_set += 1
		choice = PatternFlowTcpChecksumChoice.CUSTOM
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowTcpChecksumChoice.GENERATED)
			if obj.obj.Generated.Number() == 0 {
				obj.SetGenerated(PatternFlowTcpChecksumGenerated.GOOD)

			}

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowTcpChecksum")
			}
		} else {
			intVal := otg.PatternFlowTcpChecksum_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowTcpChecksum_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
