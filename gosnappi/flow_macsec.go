package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowMacsec *****
type flowMacsec struct {
	validation
	obj          *otg.FlowMacsec
	marshaller   marshalFlowMacsec
	unMarshaller unMarshalFlowMacsec
}

func NewFlowMacsec() FlowMacsec {
	obj := flowMacsec{obj: &otg.FlowMacsec{}}
	obj.setDefault()
	return &obj
}

func (obj *flowMacsec) msg() *otg.FlowMacsec {
	return obj.obj
}

func (obj *flowMacsec) setMsg(msg *otg.FlowMacsec) FlowMacsec {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowMacsec struct {
	obj *flowMacsec
}

type marshalFlowMacsec interface {
	// ToProto marshals FlowMacsec to protobuf object *otg.FlowMacsec
	ToProto() (*otg.FlowMacsec, error)
	// ToPbText marshals FlowMacsec to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowMacsec to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowMacsec to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowMacsec to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowMacsec struct {
	obj *flowMacsec
}

type unMarshalFlowMacsec interface {
	// FromProto unmarshals FlowMacsec from protobuf object *otg.FlowMacsec
	FromProto(msg *otg.FlowMacsec) (FlowMacsec, error)
	// FromPbText unmarshals FlowMacsec from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowMacsec from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowMacsec from JSON text
	FromJson(value string) error
}

func (obj *flowMacsec) Marshal() marshalFlowMacsec {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowMacsec{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowMacsec) Unmarshal() unMarshalFlowMacsec {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowMacsec{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowMacsec) ToProto() (*otg.FlowMacsec, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowMacsec) FromProto(msg *otg.FlowMacsec) (FlowMacsec, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowMacsec) ToPbText() (string, error) {
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

func (m *unMarshalflowMacsec) FromPbText(value string) error {
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

func (m *marshalflowMacsec) ToYaml() (string, error) {
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

func (m *unMarshalflowMacsec) FromYaml(value string) error {
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

func (m *marshalflowMacsec) ToJsonRaw() (string, error) {
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

func (m *marshalflowMacsec) ToJson() (string, error) {
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

func (m *unMarshalflowMacsec) FromJson(value string) error {
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

func (obj *flowMacsec) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowMacsec) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowMacsec) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowMacsec) Clone() (FlowMacsec, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowMacsec()
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

// FlowMacsec is mACsec packet header.
type FlowMacsec interface {
	Validation
	// msg marshals FlowMacsec to protobuf object *otg.FlowMacsec
	// and doesn't set defaults
	msg() *otg.FlowMacsec
	// setMsg unmarshals FlowMacsec from protobuf object *otg.FlowMacsec
	// and doesn't set defaults
	setMsg(*otg.FlowMacsec) FlowMacsec
	// provides marshal interface
	Marshal() marshalFlowMacsec
	// provides unmarshal interface
	Unmarshal() unMarshalFlowMacsec
	// validate validates FlowMacsec
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowMacsec, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowMacsecChoiceEnum, set in FlowMacsec
	Choice() FlowMacsecChoiceEnum
	// setChoice assigns FlowMacsecChoiceEnum provided by user to FlowMacsec
	setChoice(value FlowMacsecChoiceEnum) FlowMacsec
	// HasChoice checks if Choice has been set in FlowMacsec
	HasChoice() bool
	// getter for Auto to set choice.
	Auto()
}

type FlowMacsecChoiceEnum string

// Enum of Choice on FlowMacsec
var FlowMacsecChoice = struct {
	AUTO FlowMacsecChoiceEnum
}{
	AUTO: FlowMacsecChoiceEnum("auto"),
}

func (obj *flowMacsec) Choice() FlowMacsecChoiceEnum {
	return FlowMacsecChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for Auto to set choice
func (obj *flowMacsec) Auto() {
	obj.setChoice(FlowMacsecChoice.AUTO)
}

// Currently only auto choice is allowed. If choice is auto, MACsec header is autogenerated. If auto choice is selected, MACsec protocol must be configured in device; flow.tx_rx.choice must be of type 'device' and flow.tx_rx.device.tx_names[0] must be chosen to be an endpoint that is on or behind a MACSec enabled ethernet to be able to correctly auto-fill the fields of the MACsec header. If one of the conditions is not true, the implementation should return an error specifying the issue. A custom choice can be added in future to allow user to set specific MACsec header fields and/ or to generate flow.tx_rx.port type of traffic with MACSec header fields explicitly specified by the user.
// Choice returns a string
func (obj *flowMacsec) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowMacsec) setChoice(value FlowMacsecChoiceEnum) FlowMacsec {
	intValue, ok := otg.FlowMacsec_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowMacsecChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowMacsec_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

func (obj *flowMacsec) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *flowMacsec) setDefault() {
	var choices_set int = 0
	var choice FlowMacsecChoiceEnum
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowMacsecChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowMacsec")
			}
		} else {
			intVal := otg.FlowMacsec_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowMacsec_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
