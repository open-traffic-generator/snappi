package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** EthernetConnection *****
type ethernetConnection struct {
	validation
	obj          *otg.EthernetConnection
	marshaller   marshalEthernetConnection
	unMarshaller unMarshalEthernetConnection
}

func NewEthernetConnection() EthernetConnection {
	obj := ethernetConnection{obj: &otg.EthernetConnection{}}
	obj.setDefault()
	return &obj
}

func (obj *ethernetConnection) msg() *otg.EthernetConnection {
	return obj.obj
}

func (obj *ethernetConnection) setMsg(msg *otg.EthernetConnection) EthernetConnection {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalethernetConnection struct {
	obj *ethernetConnection
}

type marshalEthernetConnection interface {
	// ToProto marshals EthernetConnection to protobuf object *otg.EthernetConnection
	ToProto() (*otg.EthernetConnection, error)
	// ToPbText marshals EthernetConnection to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals EthernetConnection to YAML text
	ToYaml() (string, error)
	// ToJson marshals EthernetConnection to JSON text
	ToJson() (string, error)
}

type unMarshalethernetConnection struct {
	obj *ethernetConnection
}

type unMarshalEthernetConnection interface {
	// FromProto unmarshals EthernetConnection from protobuf object *otg.EthernetConnection
	FromProto(msg *otg.EthernetConnection) (EthernetConnection, error)
	// FromPbText unmarshals EthernetConnection from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals EthernetConnection from YAML text
	FromYaml(value string) error
	// FromJson unmarshals EthernetConnection from JSON text
	FromJson(value string) error
}

func (obj *ethernetConnection) Marshal() marshalEthernetConnection {
	if obj.marshaller == nil {
		obj.marshaller = &marshalethernetConnection{obj: obj}
	}
	return obj.marshaller
}

func (obj *ethernetConnection) Unmarshal() unMarshalEthernetConnection {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalethernetConnection{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalethernetConnection) ToProto() (*otg.EthernetConnection, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalethernetConnection) FromProto(msg *otg.EthernetConnection) (EthernetConnection, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalethernetConnection) ToPbText() (string, error) {
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

func (m *unMarshalethernetConnection) FromPbText(value string) error {
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

func (m *marshalethernetConnection) ToYaml() (string, error) {
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

func (m *unMarshalethernetConnection) FromYaml(value string) error {
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

func (m *marshalethernetConnection) ToJson() (string, error) {
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

func (m *unMarshalethernetConnection) FromJson(value string) error {
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

func (obj *ethernetConnection) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ethernetConnection) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ethernetConnection) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ethernetConnection) Clone() (EthernetConnection, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewEthernetConnection()
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

// EthernetConnection is ethernet interface connection to a port, LAG or VXLAN tunnel.
type EthernetConnection interface {
	Validation
	// msg marshals EthernetConnection to protobuf object *otg.EthernetConnection
	// and doesn't set defaults
	msg() *otg.EthernetConnection
	// setMsg unmarshals EthernetConnection from protobuf object *otg.EthernetConnection
	// and doesn't set defaults
	setMsg(*otg.EthernetConnection) EthernetConnection
	// provides marshal interface
	Marshal() marshalEthernetConnection
	// provides unmarshal interface
	Unmarshal() unMarshalEthernetConnection
	// validate validates EthernetConnection
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (EthernetConnection, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns EthernetConnectionChoiceEnum, set in EthernetConnection
	Choice() EthernetConnectionChoiceEnum
	// setChoice assigns EthernetConnectionChoiceEnum provided by user to EthernetConnection
	setChoice(value EthernetConnectionChoiceEnum) EthernetConnection
	// HasChoice checks if Choice has been set in EthernetConnection
	HasChoice() bool
	// PortName returns string, set in EthernetConnection.
	PortName() string
	// SetPortName assigns string provided by user to EthernetConnection
	SetPortName(value string) EthernetConnection
	// HasPortName checks if PortName has been set in EthernetConnection
	HasPortName() bool
	// LagName returns string, set in EthernetConnection.
	LagName() string
	// SetLagName assigns string provided by user to EthernetConnection
	SetLagName(value string) EthernetConnection
	// HasLagName checks if LagName has been set in EthernetConnection
	HasLagName() bool
	// VxlanName returns string, set in EthernetConnection.
	VxlanName() string
	// SetVxlanName assigns string provided by user to EthernetConnection
	SetVxlanName(value string) EthernetConnection
	// HasVxlanName checks if VxlanName has been set in EthernetConnection
	HasVxlanName() bool
}

type EthernetConnectionChoiceEnum string

// Enum of Choice on EthernetConnection
var EthernetConnectionChoice = struct {
	PORT_NAME  EthernetConnectionChoiceEnum
	LAG_NAME   EthernetConnectionChoiceEnum
	VXLAN_NAME EthernetConnectionChoiceEnum
}{
	PORT_NAME:  EthernetConnectionChoiceEnum("port_name"),
	LAG_NAME:   EthernetConnectionChoiceEnum("lag_name"),
	VXLAN_NAME: EthernetConnectionChoiceEnum("vxlan_name"),
}

func (obj *ethernetConnection) Choice() EthernetConnectionChoiceEnum {
	return EthernetConnectionChoiceEnum(obj.obj.Choice.Enum().String())
}

// port_name, lag_name or vxlan_name
// Choice returns a string
func (obj *ethernetConnection) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *ethernetConnection) setChoice(value EthernetConnectionChoiceEnum) EthernetConnection {
	intValue, ok := otg.EthernetConnection_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on EthernetConnectionChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.EthernetConnection_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.VxlanName = nil
	obj.obj.LagName = nil
	obj.obj.PortName = nil
	return obj
}

// Name of the port that the Ethernet interface is configured on.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// PortName returns a string
func (obj *ethernetConnection) PortName() string {

	if obj.obj.PortName == nil {
		obj.setChoice(EthernetConnectionChoice.PORT_NAME)
	}

	return *obj.obj.PortName

}

// Name of the port that the Ethernet interface is configured on.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// PortName returns a string
func (obj *ethernetConnection) HasPortName() bool {
	return obj.obj.PortName != nil
}

// Name of the port that the Ethernet interface is configured on.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// SetPortName sets the string value in the EthernetConnection object
func (obj *ethernetConnection) SetPortName(value string) EthernetConnection {
	obj.setChoice(EthernetConnectionChoice.PORT_NAME)
	obj.obj.PortName = &value
	return obj
}

// Name of the LAG that the Ethernet interface is configured on.
//
// x-constraint:
// - /components/schemas/Lag/properties/name
//
// x-constraint:
// - /components/schemas/Lag/properties/name
//
// LagName returns a string
func (obj *ethernetConnection) LagName() string {

	if obj.obj.LagName == nil {
		obj.setChoice(EthernetConnectionChoice.LAG_NAME)
	}

	return *obj.obj.LagName

}

// Name of the LAG that the Ethernet interface is configured on.
//
// x-constraint:
// - /components/schemas/Lag/properties/name
//
// x-constraint:
// - /components/schemas/Lag/properties/name
//
// LagName returns a string
func (obj *ethernetConnection) HasLagName() bool {
	return obj.obj.LagName != nil
}

// Name of the LAG that the Ethernet interface is configured on.
//
// x-constraint:
// - /components/schemas/Lag/properties/name
//
// x-constraint:
// - /components/schemas/Lag/properties/name
//
// SetLagName sets the string value in the EthernetConnection object
func (obj *ethernetConnection) SetLagName(value string) EthernetConnection {
	obj.setChoice(EthernetConnectionChoice.LAG_NAME)
	obj.obj.LagName = &value
	return obj
}

// Name of the VXLAN instance (or VXLAN tunnel) that this Ethernet interface is connected to.
//
// x-constraint:
// - #/components/schemas/Vxlan.V4Tunnel/properties/name
// - #/components/schemas/Vxlan.V6Tunnel/properties/name
//
// x-constraint:
// - #/components/schemas/Vxlan.V4Tunnel/properties/name
// - #/components/schemas/Vxlan.V6Tunnel/properties/name
//
// VxlanName returns a string
func (obj *ethernetConnection) VxlanName() string {

	if obj.obj.VxlanName == nil {
		obj.setChoice(EthernetConnectionChoice.VXLAN_NAME)
	}

	return *obj.obj.VxlanName

}

// Name of the VXLAN instance (or VXLAN tunnel) that this Ethernet interface is connected to.
//
// x-constraint:
// - #/components/schemas/Vxlan.V4Tunnel/properties/name
// - #/components/schemas/Vxlan.V6Tunnel/properties/name
//
// x-constraint:
// - #/components/schemas/Vxlan.V4Tunnel/properties/name
// - #/components/schemas/Vxlan.V6Tunnel/properties/name
//
// VxlanName returns a string
func (obj *ethernetConnection) HasVxlanName() bool {
	return obj.obj.VxlanName != nil
}

// Name of the VXLAN instance (or VXLAN tunnel) that this Ethernet interface is connected to.
//
// x-constraint:
// - #/components/schemas/Vxlan.V4Tunnel/properties/name
// - #/components/schemas/Vxlan.V6Tunnel/properties/name
//
// x-constraint:
// - #/components/schemas/Vxlan.V4Tunnel/properties/name
// - #/components/schemas/Vxlan.V6Tunnel/properties/name
//
// SetVxlanName sets the string value in the EthernetConnection object
func (obj *ethernetConnection) SetVxlanName(value string) EthernetConnection {
	obj.setChoice(EthernetConnectionChoice.VXLAN_NAME)
	obj.obj.VxlanName = &value
	return obj
}

func (obj *ethernetConnection) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ethernetConnection) setDefault() {
	var choices_set int = 0
	var choice EthernetConnectionChoiceEnum

	if obj.obj.PortName != nil {
		choices_set += 1
		choice = EthernetConnectionChoice.PORT_NAME
	}

	if obj.obj.LagName != nil {
		choices_set += 1
		choice = EthernetConnectionChoice.LAG_NAME
	}

	if obj.obj.VxlanName != nil {
		choices_set += 1
		choice = EthernetConnectionChoice.VXLAN_NAME
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in EthernetConnection")
			}
		} else {
			intVal := otg.EthernetConnection_Choice_Enum_value[string(choice)]
			enumValue := otg.EthernetConnection_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
