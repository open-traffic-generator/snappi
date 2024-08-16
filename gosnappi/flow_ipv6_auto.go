package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6Auto *****
type flowIpv6Auto struct {
	validation
	obj          *otg.FlowIpv6Auto
	marshaller   marshalFlowIpv6Auto
	unMarshaller unMarshalFlowIpv6Auto
}

func NewFlowIpv6Auto() FlowIpv6Auto {
	obj := flowIpv6Auto{obj: &otg.FlowIpv6Auto{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6Auto) msg() *otg.FlowIpv6Auto {
	return obj.obj
}

func (obj *flowIpv6Auto) setMsg(msg *otg.FlowIpv6Auto) FlowIpv6Auto {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6Auto struct {
	obj *flowIpv6Auto
}

type marshalFlowIpv6Auto interface {
	// ToProto marshals FlowIpv6Auto to protobuf object *otg.FlowIpv6Auto
	ToProto() (*otg.FlowIpv6Auto, error)
	// ToPbText marshals FlowIpv6Auto to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6Auto to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6Auto to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6Auto struct {
	obj *flowIpv6Auto
}

type unMarshalFlowIpv6Auto interface {
	// FromProto unmarshals FlowIpv6Auto from protobuf object *otg.FlowIpv6Auto
	FromProto(msg *otg.FlowIpv6Auto) (FlowIpv6Auto, error)
	// FromPbText unmarshals FlowIpv6Auto from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6Auto from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6Auto from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6Auto) Marshal() marshalFlowIpv6Auto {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6Auto{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6Auto) Unmarshal() unMarshalFlowIpv6Auto {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6Auto{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6Auto) ToProto() (*otg.FlowIpv6Auto, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6Auto) FromProto(msg *otg.FlowIpv6Auto) (FlowIpv6Auto, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6Auto) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6Auto) FromPbText(value string) error {
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

func (m *marshalflowIpv6Auto) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6Auto) FromYaml(value string) error {
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

func (m *marshalflowIpv6Auto) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6Auto) FromJson(value string) error {
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

func (obj *flowIpv6Auto) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6Auto) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6Auto) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6Auto) Clone() (FlowIpv6Auto, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6Auto()
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

// FlowIpv6Auto is the OTG implementation can provide a system generated, value for this property.
type FlowIpv6Auto interface {
	Validation
	// msg marshals FlowIpv6Auto to protobuf object *otg.FlowIpv6Auto
	// and doesn't set defaults
	msg() *otg.FlowIpv6Auto
	// setMsg unmarshals FlowIpv6Auto from protobuf object *otg.FlowIpv6Auto
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6Auto) FlowIpv6Auto
	// provides marshal interface
	Marshal() marshalFlowIpv6Auto
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6Auto
	// validate validates FlowIpv6Auto
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6Auto, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowIpv6AutoChoiceEnum, set in FlowIpv6Auto
	Choice() FlowIpv6AutoChoiceEnum
	// setChoice assigns FlowIpv6AutoChoiceEnum provided by user to FlowIpv6Auto
	setChoice(value FlowIpv6AutoChoiceEnum) FlowIpv6Auto
	// getter for Dhcp to set choice.
	Dhcp()
}

type FlowIpv6AutoChoiceEnum string

// Enum of Choice on FlowIpv6Auto
var FlowIpv6AutoChoice = struct {
	DHCP FlowIpv6AutoChoiceEnum
}{
	DHCP: FlowIpv6AutoChoiceEnum("dhcp"),
}

func (obj *flowIpv6Auto) Choice() FlowIpv6AutoChoiceEnum {
	return FlowIpv6AutoChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for Dhcp to set choice
func (obj *flowIpv6Auto) Dhcp() {
	obj.setChoice(FlowIpv6AutoChoice.DHCP)
}

func (obj *flowIpv6Auto) setChoice(value FlowIpv6AutoChoiceEnum) FlowIpv6Auto {
	intValue, ok := otg.FlowIpv6Auto_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowIpv6AutoChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowIpv6Auto_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

func (obj *flowIpv6Auto) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface FlowIpv6Auto")
	}
}

func (obj *flowIpv6Auto) setDefault() {
	var choices_set int = 0
	var choice FlowIpv6AutoChoiceEnum
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowIpv6Auto")
			}
		} else {
			intVal := otg.FlowIpv6Auto_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowIpv6Auto_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
