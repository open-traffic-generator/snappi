package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPLspTunnelFlag *****
type flowRSVPLspTunnelFlag struct {
	validation
	obj          *otg.FlowRSVPLspTunnelFlag
	marshaller   marshalFlowRSVPLspTunnelFlag
	unMarshaller unMarshalFlowRSVPLspTunnelFlag
}

func NewFlowRSVPLspTunnelFlag() FlowRSVPLspTunnelFlag {
	obj := flowRSVPLspTunnelFlag{obj: &otg.FlowRSVPLspTunnelFlag{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPLspTunnelFlag) msg() *otg.FlowRSVPLspTunnelFlag {
	return obj.obj
}

func (obj *flowRSVPLspTunnelFlag) setMsg(msg *otg.FlowRSVPLspTunnelFlag) FlowRSVPLspTunnelFlag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPLspTunnelFlag struct {
	obj *flowRSVPLspTunnelFlag
}

type marshalFlowRSVPLspTunnelFlag interface {
	// ToProto marshals FlowRSVPLspTunnelFlag to protobuf object *otg.FlowRSVPLspTunnelFlag
	ToProto() (*otg.FlowRSVPLspTunnelFlag, error)
	// ToPbText marshals FlowRSVPLspTunnelFlag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPLspTunnelFlag to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPLspTunnelFlag to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPLspTunnelFlag struct {
	obj *flowRSVPLspTunnelFlag
}

type unMarshalFlowRSVPLspTunnelFlag interface {
	// FromProto unmarshals FlowRSVPLspTunnelFlag from protobuf object *otg.FlowRSVPLspTunnelFlag
	FromProto(msg *otg.FlowRSVPLspTunnelFlag) (FlowRSVPLspTunnelFlag, error)
	// FromPbText unmarshals FlowRSVPLspTunnelFlag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPLspTunnelFlag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPLspTunnelFlag from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPLspTunnelFlag) Marshal() marshalFlowRSVPLspTunnelFlag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPLspTunnelFlag{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPLspTunnelFlag) Unmarshal() unMarshalFlowRSVPLspTunnelFlag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPLspTunnelFlag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPLspTunnelFlag) ToProto() (*otg.FlowRSVPLspTunnelFlag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPLspTunnelFlag) FromProto(msg *otg.FlowRSVPLspTunnelFlag) (FlowRSVPLspTunnelFlag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPLspTunnelFlag) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPLspTunnelFlag) FromPbText(value string) error {
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

func (m *marshalflowRSVPLspTunnelFlag) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPLspTunnelFlag) FromYaml(value string) error {
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

func (m *marshalflowRSVPLspTunnelFlag) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPLspTunnelFlag) FromJson(value string) error {
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

func (obj *flowRSVPLspTunnelFlag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPLspTunnelFlag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPLspTunnelFlag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPLspTunnelFlag) Clone() (FlowRSVPLspTunnelFlag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPLspTunnelFlag()
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

// FlowRSVPLspTunnelFlag is description is TBD
type FlowRSVPLspTunnelFlag interface {
	Validation
	// msg marshals FlowRSVPLspTunnelFlag to protobuf object *otg.FlowRSVPLspTunnelFlag
	// and doesn't set defaults
	msg() *otg.FlowRSVPLspTunnelFlag
	// setMsg unmarshals FlowRSVPLspTunnelFlag from protobuf object *otg.FlowRSVPLspTunnelFlag
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPLspTunnelFlag) FlowRSVPLspTunnelFlag
	// provides marshal interface
	Marshal() marshalFlowRSVPLspTunnelFlag
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPLspTunnelFlag
	// validate validates FlowRSVPLspTunnelFlag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPLspTunnelFlag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowRSVPLspTunnelFlagChoiceEnum, set in FlowRSVPLspTunnelFlag
	Choice() FlowRSVPLspTunnelFlagChoiceEnum
	// setChoice assigns FlowRSVPLspTunnelFlagChoiceEnum provided by user to FlowRSVPLspTunnelFlag
	setChoice(value FlowRSVPLspTunnelFlagChoiceEnum) FlowRSVPLspTunnelFlag
	// HasChoice checks if Choice has been set in FlowRSVPLspTunnelFlag
	HasChoice() bool
	// getter for LocalProtectionDesired to set choice.
	LocalProtectionDesired()
	// getter for SeStyleDesired to set choice.
	SeStyleDesired()
	// getter for LabelRecordingDesired to set choice.
	LabelRecordingDesired()
}

type FlowRSVPLspTunnelFlagChoiceEnum string

// Enum of Choice on FlowRSVPLspTunnelFlag
var FlowRSVPLspTunnelFlagChoice = struct {
	LOCAL_PROTECTION_DESIRED FlowRSVPLspTunnelFlagChoiceEnum
	LABEL_RECORDING_DESIRED  FlowRSVPLspTunnelFlagChoiceEnum
	SE_STYLE_DESIRED         FlowRSVPLspTunnelFlagChoiceEnum
}{
	LOCAL_PROTECTION_DESIRED: FlowRSVPLspTunnelFlagChoiceEnum("local_protection_desired"),
	LABEL_RECORDING_DESIRED:  FlowRSVPLspTunnelFlagChoiceEnum("label_recording_desired"),
	SE_STYLE_DESIRED:         FlowRSVPLspTunnelFlagChoiceEnum("se_style_desired"),
}

func (obj *flowRSVPLspTunnelFlag) Choice() FlowRSVPLspTunnelFlagChoiceEnum {
	return FlowRSVPLspTunnelFlagChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for LocalProtectionDesired to set choice
func (obj *flowRSVPLspTunnelFlag) LocalProtectionDesired() {
	obj.setChoice(FlowRSVPLspTunnelFlagChoice.LOCAL_PROTECTION_DESIRED)
}

// getter for SeStyleDesired to set choice
func (obj *flowRSVPLspTunnelFlag) SeStyleDesired() {
	obj.setChoice(FlowRSVPLspTunnelFlagChoice.SE_STYLE_DESIRED)
}

// getter for LabelRecordingDesired to set choice
func (obj *flowRSVPLspTunnelFlag) LabelRecordingDesired() {
	obj.setChoice(FlowRSVPLspTunnelFlagChoice.LABEL_RECORDING_DESIRED)
}

// description is TBD
// Choice returns a string
func (obj *flowRSVPLspTunnelFlag) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowRSVPLspTunnelFlag) setChoice(value FlowRSVPLspTunnelFlagChoiceEnum) FlowRSVPLspTunnelFlag {
	intValue, ok := otg.FlowRSVPLspTunnelFlag_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowRSVPLspTunnelFlagChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowRSVPLspTunnelFlag_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

func (obj *flowRSVPLspTunnelFlag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *flowRSVPLspTunnelFlag) setDefault() {
	var choices_set int = 0
	var choice FlowRSVPLspTunnelFlagChoiceEnum
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowRSVPLspTunnelFlagChoice.LOCAL_PROTECTION_DESIRED)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowRSVPLspTunnelFlag")
			}
		} else {
			intVal := otg.FlowRSVPLspTunnelFlag_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowRSVPLspTunnelFlag_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
