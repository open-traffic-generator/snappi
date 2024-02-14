package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPRecordRouteIPv4Flag *****
type flowRSVPRecordRouteIPv4Flag struct {
	validation
	obj          *otg.FlowRSVPRecordRouteIPv4Flag
	marshaller   marshalFlowRSVPRecordRouteIPv4Flag
	unMarshaller unMarshalFlowRSVPRecordRouteIPv4Flag
}

func NewFlowRSVPRecordRouteIPv4Flag() FlowRSVPRecordRouteIPv4Flag {
	obj := flowRSVPRecordRouteIPv4Flag{obj: &otg.FlowRSVPRecordRouteIPv4Flag{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPRecordRouteIPv4Flag) msg() *otg.FlowRSVPRecordRouteIPv4Flag {
	return obj.obj
}

func (obj *flowRSVPRecordRouteIPv4Flag) setMsg(msg *otg.FlowRSVPRecordRouteIPv4Flag) FlowRSVPRecordRouteIPv4Flag {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPRecordRouteIPv4Flag struct {
	obj *flowRSVPRecordRouteIPv4Flag
}

type marshalFlowRSVPRecordRouteIPv4Flag interface {
	// ToProto marshals FlowRSVPRecordRouteIPv4Flag to protobuf object *otg.FlowRSVPRecordRouteIPv4Flag
	ToProto() (*otg.FlowRSVPRecordRouteIPv4Flag, error)
	// ToPbText marshals FlowRSVPRecordRouteIPv4Flag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPRecordRouteIPv4Flag to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPRecordRouteIPv4Flag to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPRecordRouteIPv4Flag struct {
	obj *flowRSVPRecordRouteIPv4Flag
}

type unMarshalFlowRSVPRecordRouteIPv4Flag interface {
	// FromProto unmarshals FlowRSVPRecordRouteIPv4Flag from protobuf object *otg.FlowRSVPRecordRouteIPv4Flag
	FromProto(msg *otg.FlowRSVPRecordRouteIPv4Flag) (FlowRSVPRecordRouteIPv4Flag, error)
	// FromPbText unmarshals FlowRSVPRecordRouteIPv4Flag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPRecordRouteIPv4Flag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPRecordRouteIPv4Flag from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPRecordRouteIPv4Flag) Marshal() marshalFlowRSVPRecordRouteIPv4Flag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPRecordRouteIPv4Flag{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPRecordRouteIPv4Flag) Unmarshal() unMarshalFlowRSVPRecordRouteIPv4Flag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPRecordRouteIPv4Flag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPRecordRouteIPv4Flag) ToProto() (*otg.FlowRSVPRecordRouteIPv4Flag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPRecordRouteIPv4Flag) FromProto(msg *otg.FlowRSVPRecordRouteIPv4Flag) (FlowRSVPRecordRouteIPv4Flag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPRecordRouteIPv4Flag) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPRecordRouteIPv4Flag) FromPbText(value string) error {
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

func (m *marshalflowRSVPRecordRouteIPv4Flag) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPRecordRouteIPv4Flag) FromYaml(value string) error {
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

func (m *marshalflowRSVPRecordRouteIPv4Flag) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPRecordRouteIPv4Flag) FromJson(value string) error {
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

func (obj *flowRSVPRecordRouteIPv4Flag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPRecordRouteIPv4Flag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPRecordRouteIPv4Flag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPRecordRouteIPv4Flag) Clone() (FlowRSVPRecordRouteIPv4Flag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPRecordRouteIPv4Flag()
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

// FlowRSVPRecordRouteIPv4Flag is description is TBD
type FlowRSVPRecordRouteIPv4Flag interface {
	Validation
	// msg marshals FlowRSVPRecordRouteIPv4Flag to protobuf object *otg.FlowRSVPRecordRouteIPv4Flag
	// and doesn't set defaults
	msg() *otg.FlowRSVPRecordRouteIPv4Flag
	// setMsg unmarshals FlowRSVPRecordRouteIPv4Flag from protobuf object *otg.FlowRSVPRecordRouteIPv4Flag
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPRecordRouteIPv4Flag) FlowRSVPRecordRouteIPv4Flag
	// provides marshal interface
	Marshal() marshalFlowRSVPRecordRouteIPv4Flag
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPRecordRouteIPv4Flag
	// validate validates FlowRSVPRecordRouteIPv4Flag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPRecordRouteIPv4Flag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowRSVPRecordRouteIPv4FlagChoiceEnum, set in FlowRSVPRecordRouteIPv4Flag
	Choice() FlowRSVPRecordRouteIPv4FlagChoiceEnum
	// setChoice assigns FlowRSVPRecordRouteIPv4FlagChoiceEnum provided by user to FlowRSVPRecordRouteIPv4Flag
	setChoice(value FlowRSVPRecordRouteIPv4FlagChoiceEnum) FlowRSVPRecordRouteIPv4Flag
	// HasChoice checks if Choice has been set in FlowRSVPRecordRouteIPv4Flag
	HasChoice() bool
	// getter for LocalProtectionAvailable to set choice.
	LocalProtectionAvailable()
	// getter for LocalProtectionInUse to set choice.
	LocalProtectionInUse()
}

type FlowRSVPRecordRouteIPv4FlagChoiceEnum string

// Enum of Choice on FlowRSVPRecordRouteIPv4Flag
var FlowRSVPRecordRouteIPv4FlagChoice = struct {
	LOCAL_PROTECTION_AVAILABLE FlowRSVPRecordRouteIPv4FlagChoiceEnum
	LOCAL_PROTECTION_IN_USE    FlowRSVPRecordRouteIPv4FlagChoiceEnum
}{
	LOCAL_PROTECTION_AVAILABLE: FlowRSVPRecordRouteIPv4FlagChoiceEnum("local_protection_available"),
	LOCAL_PROTECTION_IN_USE:    FlowRSVPRecordRouteIPv4FlagChoiceEnum("local_protection_in_use"),
}

func (obj *flowRSVPRecordRouteIPv4Flag) Choice() FlowRSVPRecordRouteIPv4FlagChoiceEnum {
	return FlowRSVPRecordRouteIPv4FlagChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for LocalProtectionAvailable to set choice
func (obj *flowRSVPRecordRouteIPv4Flag) LocalProtectionAvailable() {
	obj.setChoice(FlowRSVPRecordRouteIPv4FlagChoice.LOCAL_PROTECTION_AVAILABLE)
}

// getter for LocalProtectionInUse to set choice
func (obj *flowRSVPRecordRouteIPv4Flag) LocalProtectionInUse() {
	obj.setChoice(FlowRSVPRecordRouteIPv4FlagChoice.LOCAL_PROTECTION_IN_USE)
}

// description is TBD
// Choice returns a string
func (obj *flowRSVPRecordRouteIPv4Flag) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowRSVPRecordRouteIPv4Flag) setChoice(value FlowRSVPRecordRouteIPv4FlagChoiceEnum) FlowRSVPRecordRouteIPv4Flag {
	intValue, ok := otg.FlowRSVPRecordRouteIPv4Flag_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowRSVPRecordRouteIPv4FlagChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowRSVPRecordRouteIPv4Flag_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

func (obj *flowRSVPRecordRouteIPv4Flag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *flowRSVPRecordRouteIPv4Flag) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(FlowRSVPRecordRouteIPv4FlagChoice.LOCAL_PROTECTION_AVAILABLE)

	}

}
