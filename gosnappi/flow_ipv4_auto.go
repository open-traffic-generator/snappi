package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv4Auto *****
type flowIpv4Auto struct {
	validation
	obj          *otg.FlowIpv4Auto
	marshaller   marshalFlowIpv4Auto
	unMarshaller unMarshalFlowIpv4Auto
}

func NewFlowIpv4Auto() FlowIpv4Auto {
	obj := flowIpv4Auto{obj: &otg.FlowIpv4Auto{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv4Auto) msg() *otg.FlowIpv4Auto {
	return obj.obj
}

func (obj *flowIpv4Auto) setMsg(msg *otg.FlowIpv4Auto) FlowIpv4Auto {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv4Auto struct {
	obj *flowIpv4Auto
}

type marshalFlowIpv4Auto interface {
	// ToProto marshals FlowIpv4Auto to protobuf object *otg.FlowIpv4Auto
	ToProto() (*otg.FlowIpv4Auto, error)
	// ToPbText marshals FlowIpv4Auto to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv4Auto to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv4Auto to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv4Auto struct {
	obj *flowIpv4Auto
}

type unMarshalFlowIpv4Auto interface {
	// FromProto unmarshals FlowIpv4Auto from protobuf object *otg.FlowIpv4Auto
	FromProto(msg *otg.FlowIpv4Auto) (FlowIpv4Auto, error)
	// FromPbText unmarshals FlowIpv4Auto from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv4Auto from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv4Auto from JSON text
	FromJson(value string) error
}

func (obj *flowIpv4Auto) Marshal() marshalFlowIpv4Auto {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv4Auto{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv4Auto) Unmarshal() unMarshalFlowIpv4Auto {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv4Auto{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv4Auto) ToProto() (*otg.FlowIpv4Auto, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv4Auto) FromProto(msg *otg.FlowIpv4Auto) (FlowIpv4Auto, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv4Auto) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv4Auto) FromPbText(value string) error {
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

func (m *marshalflowIpv4Auto) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv4Auto) FromYaml(value string) error {
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

func (m *marshalflowIpv4Auto) ToJson() (string, error) {
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

func (m *unMarshalflowIpv4Auto) FromJson(value string) error {
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

func (obj *flowIpv4Auto) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv4Auto) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv4Auto) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv4Auto) Clone() (FlowIpv4Auto, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv4Auto()
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

// FlowIpv4Auto is the OTG implementation can provide a system generated, value for this property.
type FlowIpv4Auto interface {
	Validation
	// msg marshals FlowIpv4Auto to protobuf object *otg.FlowIpv4Auto
	// and doesn't set defaults
	msg() *otg.FlowIpv4Auto
	// setMsg unmarshals FlowIpv4Auto from protobuf object *otg.FlowIpv4Auto
	// and doesn't set defaults
	setMsg(*otg.FlowIpv4Auto) FlowIpv4Auto
	// provides marshal interface
	Marshal() marshalFlowIpv4Auto
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv4Auto
	// validate validates FlowIpv4Auto
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv4Auto, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowIpv4AutoChoiceEnum, set in FlowIpv4Auto
	Choice() FlowIpv4AutoChoiceEnum
	// setChoice assigns FlowIpv4AutoChoiceEnum provided by user to FlowIpv4Auto
	setChoice(value FlowIpv4AutoChoiceEnum) FlowIpv4Auto
	// getter for Dhcp to set choice.
	Dhcp()
}

type FlowIpv4AutoChoiceEnum string

// Enum of Choice on FlowIpv4Auto
var FlowIpv4AutoChoice = struct {
	DHCP FlowIpv4AutoChoiceEnum
}{
	DHCP: FlowIpv4AutoChoiceEnum("dhcp"),
}

func (obj *flowIpv4Auto) Choice() FlowIpv4AutoChoiceEnum {
	return FlowIpv4AutoChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for Dhcp to set choice
func (obj *flowIpv4Auto) Dhcp() {
	obj.setChoice(FlowIpv4AutoChoice.DHCP)
}

func (obj *flowIpv4Auto) setChoice(value FlowIpv4AutoChoiceEnum) FlowIpv4Auto {
	intValue, ok := otg.FlowIpv4Auto_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowIpv4AutoChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowIpv4Auto_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

func (obj *flowIpv4Auto) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface FlowIpv4Auto")
	}
}

func (obj *flowIpv4Auto) setDefault() {
	var choices_set int = 0
	var choice FlowIpv4AutoChoiceEnum
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowIpv4Auto")
			}
		} else {
			intVal := otg.FlowIpv4Auto_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowIpv4Auto_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
