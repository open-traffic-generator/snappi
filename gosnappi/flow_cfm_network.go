package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfmNetwork *****
type flowCfmNetwork struct {
	validation
	obj          *otg.FlowCfmNetwork
	marshaller   marshalFlowCfmNetwork
	unMarshaller unMarshalFlowCfmNetwork
}

func NewFlowCfmNetwork() FlowCfmNetwork {
	obj := flowCfmNetwork{obj: &otg.FlowCfmNetwork{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfmNetwork) msg() *otg.FlowCfmNetwork {
	return obj.obj
}

func (obj *flowCfmNetwork) setMsg(msg *otg.FlowCfmNetwork) FlowCfmNetwork {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfmNetwork struct {
	obj *flowCfmNetwork
}

type marshalFlowCfmNetwork interface {
	// ToProto marshals FlowCfmNetwork to protobuf object *otg.FlowCfmNetwork
	ToProto() (*otg.FlowCfmNetwork, error)
	// ToPbText marshals FlowCfmNetwork to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfmNetwork to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfmNetwork to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfmNetwork struct {
	obj *flowCfmNetwork
}

type unMarshalFlowCfmNetwork interface {
	// FromProto unmarshals FlowCfmNetwork from protobuf object *otg.FlowCfmNetwork
	FromProto(msg *otg.FlowCfmNetwork) (FlowCfmNetwork, error)
	// FromPbText unmarshals FlowCfmNetwork from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfmNetwork from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfmNetwork from JSON text
	FromJson(value string) error
}

func (obj *flowCfmNetwork) Marshal() marshalFlowCfmNetwork {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfmNetwork{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfmNetwork) Unmarshal() unMarshalFlowCfmNetwork {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfmNetwork{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfmNetwork) ToProto() (*otg.FlowCfmNetwork, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfmNetwork) FromProto(msg *otg.FlowCfmNetwork) (FlowCfmNetwork, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfmNetwork) ToPbText() (string, error) {
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

func (m *unMarshalflowCfmNetwork) FromPbText(value string) error {
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

func (m *marshalflowCfmNetwork) ToYaml() (string, error) {
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

func (m *unMarshalflowCfmNetwork) FromYaml(value string) error {
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

func (m *marshalflowCfmNetwork) ToJson() (string, error) {
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

func (m *unMarshalflowCfmNetwork) FromJson(value string) error {
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

func (obj *flowCfmNetwork) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfmNetwork) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfmNetwork) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfmNetwork) Clone() (FlowCfmNetwork, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfmNetwork()
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

// FlowCfmNetwork is typed network address used as the Sender ID TLV Chassis ID Subtype 5 (IEEE 802.1AB-2016 Table 8-3). On the wire, the choice value encodes the leading IANA address family octet (1 = IPv4, 2 = IPv6), followed by the corresponding IP address bytes.
type FlowCfmNetwork interface {
	Validation
	// msg marshals FlowCfmNetwork to protobuf object *otg.FlowCfmNetwork
	// and doesn't set defaults
	msg() *otg.FlowCfmNetwork
	// setMsg unmarshals FlowCfmNetwork from protobuf object *otg.FlowCfmNetwork
	// and doesn't set defaults
	setMsg(*otg.FlowCfmNetwork) FlowCfmNetwork
	// provides marshal interface
	Marshal() marshalFlowCfmNetwork
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfmNetwork
	// validate validates FlowCfmNetwork
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfmNetwork, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowCfmNetworkChoiceEnum, set in FlowCfmNetwork
	Choice() FlowCfmNetworkChoiceEnum
	// setChoice assigns FlowCfmNetworkChoiceEnum provided by user to FlowCfmNetwork
	setChoice(value FlowCfmNetworkChoiceEnum) FlowCfmNetwork
	// HasChoice checks if Choice has been set in FlowCfmNetwork
	HasChoice() bool
	// Ipv4 returns string, set in FlowCfmNetwork.
	Ipv4() string
	// SetIpv4 assigns string provided by user to FlowCfmNetwork
	SetIpv4(value string) FlowCfmNetwork
	// HasIpv4 checks if Ipv4 has been set in FlowCfmNetwork
	HasIpv4() bool
	// Ipv6 returns string, set in FlowCfmNetwork.
	Ipv6() string
	// SetIpv6 assigns string provided by user to FlowCfmNetwork
	SetIpv6(value string) FlowCfmNetwork
	// HasIpv6 checks if Ipv6 has been set in FlowCfmNetwork
	HasIpv6() bool
}

type FlowCfmNetworkChoiceEnum string

// Enum of Choice on FlowCfmNetwork
var FlowCfmNetworkChoice = struct {
	IPV4 FlowCfmNetworkChoiceEnum
	IPV6 FlowCfmNetworkChoiceEnum
}{
	IPV4: FlowCfmNetworkChoiceEnum("ipv4"),
	IPV6: FlowCfmNetworkChoiceEnum("ipv6"),
}

func (obj *flowCfmNetwork) Choice() FlowCfmNetworkChoiceEnum {
	return FlowCfmNetworkChoiceEnum(obj.obj.Choice.Enum().String())
}

// IANA address family, encodes the first octet of the network address.
// Choice returns a string
func (obj *flowCfmNetwork) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowCfmNetwork) setChoice(value FlowCfmNetworkChoiceEnum) FlowCfmNetwork {
	intValue, ok := otg.FlowCfmNetwork_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowCfmNetworkChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowCfmNetwork_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Ipv6 = nil
	obj.obj.Ipv4 = nil

	if value == FlowCfmNetworkChoice.IPV4 {
		defaultValue := "0.0.0.0"
		obj.obj.Ipv4 = &defaultValue
	}

	if value == FlowCfmNetworkChoice.IPV6 {
		defaultValue := "::0"
		obj.obj.Ipv6 = &defaultValue
	}

	return obj
}

// Chassis ID IPv4 address
// Ipv4 returns a string
func (obj *flowCfmNetwork) Ipv4() string {

	if obj.obj.Ipv4 == nil {
		obj.setChoice(FlowCfmNetworkChoice.IPV4)
	}

	return *obj.obj.Ipv4

}

// Chassis ID IPv4 address
// Ipv4 returns a string
func (obj *flowCfmNetwork) HasIpv4() bool {
	return obj.obj.Ipv4 != nil
}

// Chassis ID IPv4 address
// SetIpv4 sets the string value in the FlowCfmNetwork object
func (obj *flowCfmNetwork) SetIpv4(value string) FlowCfmNetwork {
	obj.setChoice(FlowCfmNetworkChoice.IPV4)
	obj.obj.Ipv4 = &value
	return obj
}

// Chassis ID IPv6 address
// Ipv6 returns a string
func (obj *flowCfmNetwork) Ipv6() string {

	if obj.obj.Ipv6 == nil {
		obj.setChoice(FlowCfmNetworkChoice.IPV6)
	}

	return *obj.obj.Ipv6

}

// Chassis ID IPv6 address
// Ipv6 returns a string
func (obj *flowCfmNetwork) HasIpv6() bool {
	return obj.obj.Ipv6 != nil
}

// Chassis ID IPv6 address
// SetIpv6 sets the string value in the FlowCfmNetwork object
func (obj *flowCfmNetwork) SetIpv6(value string) FlowCfmNetwork {
	obj.setChoice(FlowCfmNetworkChoice.IPV6)
	obj.obj.Ipv6 = &value
	return obj
}

func (obj *flowCfmNetwork) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv4 != nil {

		err := obj.validateIpv4(obj.Ipv4())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowCfmNetwork.Ipv4"))
		}

	}

	if obj.obj.Ipv6 != nil {

		err := obj.validateIpv6(obj.Ipv6())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowCfmNetwork.Ipv6"))
		}

	}

}

func (obj *flowCfmNetwork) setDefault() {
	var choices_set int = 0
	var choice FlowCfmNetworkChoiceEnum

	if obj.obj.Ipv4 != nil {
		choices_set += 1
		choice = FlowCfmNetworkChoice.IPV4
	}

	if obj.obj.Ipv6 != nil {
		choices_set += 1
		choice = FlowCfmNetworkChoice.IPV6
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowCfmNetworkChoice.IPV4)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowCfmNetwork")
			}
		} else {
			intVal := otg.FlowCfmNetwork_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowCfmNetwork_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
