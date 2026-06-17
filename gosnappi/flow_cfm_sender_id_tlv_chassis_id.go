package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfmSenderIdTlvChassisId *****
type flowCfmSenderIdTlvChassisId struct {
	validation
	obj                           *otg.FlowCfmSenderIdTlvChassisId
	marshaller                    marshalFlowCfmSenderIdTlvChassisId
	unMarshaller                  unMarshalFlowCfmSenderIdTlvChassisId
	lengthHolder                  FlowCfmLength
	networkAddressHolder          FlowCfmNetwork
	managementAddressDomainHolder FlowCfmSenderIdTlvMADomain
}

func NewFlowCfmSenderIdTlvChassisId() FlowCfmSenderIdTlvChassisId {
	obj := flowCfmSenderIdTlvChassisId{obj: &otg.FlowCfmSenderIdTlvChassisId{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfmSenderIdTlvChassisId) msg() *otg.FlowCfmSenderIdTlvChassisId {
	return obj.obj
}

func (obj *flowCfmSenderIdTlvChassisId) setMsg(msg *otg.FlowCfmSenderIdTlvChassisId) FlowCfmSenderIdTlvChassisId {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfmSenderIdTlvChassisId struct {
	obj *flowCfmSenderIdTlvChassisId
}

type marshalFlowCfmSenderIdTlvChassisId interface {
	// ToProto marshals FlowCfmSenderIdTlvChassisId to protobuf object *otg.FlowCfmSenderIdTlvChassisId
	ToProto() (*otg.FlowCfmSenderIdTlvChassisId, error)
	// ToPbText marshals FlowCfmSenderIdTlvChassisId to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfmSenderIdTlvChassisId to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfmSenderIdTlvChassisId to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfmSenderIdTlvChassisId struct {
	obj *flowCfmSenderIdTlvChassisId
}

type unMarshalFlowCfmSenderIdTlvChassisId interface {
	// FromProto unmarshals FlowCfmSenderIdTlvChassisId from protobuf object *otg.FlowCfmSenderIdTlvChassisId
	FromProto(msg *otg.FlowCfmSenderIdTlvChassisId) (FlowCfmSenderIdTlvChassisId, error)
	// FromPbText unmarshals FlowCfmSenderIdTlvChassisId from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfmSenderIdTlvChassisId from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfmSenderIdTlvChassisId from JSON text
	FromJson(value string) error
}

func (obj *flowCfmSenderIdTlvChassisId) Marshal() marshalFlowCfmSenderIdTlvChassisId {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfmSenderIdTlvChassisId{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfmSenderIdTlvChassisId) Unmarshal() unMarshalFlowCfmSenderIdTlvChassisId {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfmSenderIdTlvChassisId{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfmSenderIdTlvChassisId) ToProto() (*otg.FlowCfmSenderIdTlvChassisId, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfmSenderIdTlvChassisId) FromProto(msg *otg.FlowCfmSenderIdTlvChassisId) (FlowCfmSenderIdTlvChassisId, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfmSenderIdTlvChassisId) ToPbText() (string, error) {
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

func (m *unMarshalflowCfmSenderIdTlvChassisId) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalflowCfmSenderIdTlvChassisId) ToYaml() (string, error) {
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

func (m *unMarshalflowCfmSenderIdTlvChassisId) FromYaml(value string) error {
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
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalflowCfmSenderIdTlvChassisId) ToJson() (string, error) {
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

func (m *unMarshalflowCfmSenderIdTlvChassisId) FromJson(value string) error {
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
	m.obj.setNil()
	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *flowCfmSenderIdTlvChassisId) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfmSenderIdTlvChassisId) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfmSenderIdTlvChassisId) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfmSenderIdTlvChassisId) Clone() (FlowCfmSenderIdTlvChassisId, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfmSenderIdTlvChassisId()
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

func (obj *flowCfmSenderIdTlvChassisId) setNil() {
	obj.lengthHolder = nil
	obj.networkAddressHolder = nil
	obj.managementAddressDomainHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowCfmSenderIdTlvChassisId is sender ID TLV with Chassis ID
type FlowCfmSenderIdTlvChassisId interface {
	Validation
	// msg marshals FlowCfmSenderIdTlvChassisId to protobuf object *otg.FlowCfmSenderIdTlvChassisId
	// and doesn't set defaults
	msg() *otg.FlowCfmSenderIdTlvChassisId
	// setMsg unmarshals FlowCfmSenderIdTlvChassisId from protobuf object *otg.FlowCfmSenderIdTlvChassisId
	// and doesn't set defaults
	setMsg(*otg.FlowCfmSenderIdTlvChassisId) FlowCfmSenderIdTlvChassisId
	// provides marshal interface
	Marshal() marshalFlowCfmSenderIdTlvChassisId
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfmSenderIdTlvChassisId
	// validate validates FlowCfmSenderIdTlvChassisId
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfmSenderIdTlvChassisId, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowCfmSenderIdTlvChassisIdChoiceEnum, set in FlowCfmSenderIdTlvChassisId
	Choice() FlowCfmSenderIdTlvChassisIdChoiceEnum
	// setChoice assigns FlowCfmSenderIdTlvChassisIdChoiceEnum provided by user to FlowCfmSenderIdTlvChassisId
	setChoice(value FlowCfmSenderIdTlvChassisIdChoiceEnum) FlowCfmSenderIdTlvChassisId
	// HasChoice checks if Choice has been set in FlowCfmSenderIdTlvChassisId
	HasChoice() bool
	// Length returns FlowCfmLength, set in FlowCfmSenderIdTlvChassisId.
	// FlowCfmLength is length field with auto-compute or explicit-value choice. When auto, the implementation derives the value from the associated data field.
	Length() FlowCfmLength
	// SetLength assigns FlowCfmLength provided by user to FlowCfmSenderIdTlvChassisId.
	// FlowCfmLength is length field with auto-compute or explicit-value choice. When auto, the implementation derives the value from the associated data field.
	SetLength(value FlowCfmLength) FlowCfmSenderIdTlvChassisId
	// HasLength checks if Length has been set in FlowCfmSenderIdTlvChassisId
	HasLength() bool
	// ChassisComponent returns string, set in FlowCfmSenderIdTlvChassisId.
	ChassisComponent() string
	// SetChassisComponent assigns string provided by user to FlowCfmSenderIdTlvChassisId
	SetChassisComponent(value string) FlowCfmSenderIdTlvChassisId
	// HasChassisComponent checks if ChassisComponent has been set in FlowCfmSenderIdTlvChassisId
	HasChassisComponent() bool
	// InterfaceAlias returns string, set in FlowCfmSenderIdTlvChassisId.
	InterfaceAlias() string
	// SetInterfaceAlias assigns string provided by user to FlowCfmSenderIdTlvChassisId
	SetInterfaceAlias(value string) FlowCfmSenderIdTlvChassisId
	// HasInterfaceAlias checks if InterfaceAlias has been set in FlowCfmSenderIdTlvChassisId
	HasInterfaceAlias() bool
	// PortComponent returns string, set in FlowCfmSenderIdTlvChassisId.
	PortComponent() string
	// SetPortComponent assigns string provided by user to FlowCfmSenderIdTlvChassisId
	SetPortComponent(value string) FlowCfmSenderIdTlvChassisId
	// HasPortComponent checks if PortComponent has been set in FlowCfmSenderIdTlvChassisId
	HasPortComponent() bool
	// MacAddress returns string, set in FlowCfmSenderIdTlvChassisId.
	MacAddress() string
	// SetMacAddress assigns string provided by user to FlowCfmSenderIdTlvChassisId
	SetMacAddress(value string) FlowCfmSenderIdTlvChassisId
	// HasMacAddress checks if MacAddress has been set in FlowCfmSenderIdTlvChassisId
	HasMacAddress() bool
	// NetworkAddress returns FlowCfmNetwork, set in FlowCfmSenderIdTlvChassisId.
	// FlowCfmNetwork is typed network address used as the Sender ID TLV Chassis ID Subtype 5 (IEEE 802.1AB-2016 Table 8-3). On the wire, the choice value encodes the leading IANA address family octet (1 = IPv4, 2 = IPv6), followed by the corresponding IP address bytes.
	NetworkAddress() FlowCfmNetwork
	// SetNetworkAddress assigns FlowCfmNetwork provided by user to FlowCfmSenderIdTlvChassisId.
	// FlowCfmNetwork is typed network address used as the Sender ID TLV Chassis ID Subtype 5 (IEEE 802.1AB-2016 Table 8-3). On the wire, the choice value encodes the leading IANA address family octet (1 = IPv4, 2 = IPv6), followed by the corresponding IP address bytes.
	SetNetworkAddress(value FlowCfmNetwork) FlowCfmSenderIdTlvChassisId
	// HasNetworkAddress checks if NetworkAddress has been set in FlowCfmSenderIdTlvChassisId
	HasNetworkAddress() bool
	// InterfaceName returns string, set in FlowCfmSenderIdTlvChassisId.
	InterfaceName() string
	// SetInterfaceName assigns string provided by user to FlowCfmSenderIdTlvChassisId
	SetInterfaceName(value string) FlowCfmSenderIdTlvChassisId
	// HasInterfaceName checks if InterfaceName has been set in FlowCfmSenderIdTlvChassisId
	HasInterfaceName() bool
	// LocallyAssigned returns string, set in FlowCfmSenderIdTlvChassisId.
	LocallyAssigned() string
	// SetLocallyAssigned assigns string provided by user to FlowCfmSenderIdTlvChassisId
	SetLocallyAssigned(value string) FlowCfmSenderIdTlvChassisId
	// HasLocallyAssigned checks if LocallyAssigned has been set in FlowCfmSenderIdTlvChassisId
	HasLocallyAssigned() bool
	// ManagementAddressDomain returns FlowCfmSenderIdTlvMADomain, set in FlowCfmSenderIdTlvChassisId.
	// FlowCfmSenderIdTlvMADomain is sender ID TLV with Management Address Domain
	ManagementAddressDomain() FlowCfmSenderIdTlvMADomain
	// SetManagementAddressDomain assigns FlowCfmSenderIdTlvMADomain provided by user to FlowCfmSenderIdTlvChassisId.
	// FlowCfmSenderIdTlvMADomain is sender ID TLV with Management Address Domain
	SetManagementAddressDomain(value FlowCfmSenderIdTlvMADomain) FlowCfmSenderIdTlvChassisId
	// HasManagementAddressDomain checks if ManagementAddressDomain has been set in FlowCfmSenderIdTlvChassisId
	HasManagementAddressDomain() bool
	setNil()
}

type FlowCfmSenderIdTlvChassisIdChoiceEnum string

// Enum of Choice on FlowCfmSenderIdTlvChassisId
var FlowCfmSenderIdTlvChassisIdChoice = struct {
	CHASSIS_COMPONENT FlowCfmSenderIdTlvChassisIdChoiceEnum
	INTERFACE_ALIAS   FlowCfmSenderIdTlvChassisIdChoiceEnum
	PORT_COMPONENT    FlowCfmSenderIdTlvChassisIdChoiceEnum
	MAC_ADDRESS       FlowCfmSenderIdTlvChassisIdChoiceEnum
	NETWORK_ADDRESS   FlowCfmSenderIdTlvChassisIdChoiceEnum
	INTERFACE_NAME    FlowCfmSenderIdTlvChassisIdChoiceEnum
	LOCALLY_ASSIGNED  FlowCfmSenderIdTlvChassisIdChoiceEnum
}{
	CHASSIS_COMPONENT: FlowCfmSenderIdTlvChassisIdChoiceEnum("chassis_component"),
	INTERFACE_ALIAS:   FlowCfmSenderIdTlvChassisIdChoiceEnum("interface_alias"),
	PORT_COMPONENT:    FlowCfmSenderIdTlvChassisIdChoiceEnum("port_component"),
	MAC_ADDRESS:       FlowCfmSenderIdTlvChassisIdChoiceEnum("mac_address"),
	NETWORK_ADDRESS:   FlowCfmSenderIdTlvChassisIdChoiceEnum("network_address"),
	INTERFACE_NAME:    FlowCfmSenderIdTlvChassisIdChoiceEnum("interface_name"),
	LOCALLY_ASSIGNED:  FlowCfmSenderIdTlvChassisIdChoiceEnum("locally_assigned"),
}

func (obj *flowCfmSenderIdTlvChassisId) Choice() FlowCfmSenderIdTlvChassisIdChoiceEnum {
	return FlowCfmSenderIdTlvChassisIdChoiceEnum(obj.obj.Choice.Enum().String())
}

// Select the LLDP Chassis ID Subtype (IEEE 802.1AB-2016 Table 8-3).
// Choice returns a string
func (obj *flowCfmSenderIdTlvChassisId) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowCfmSenderIdTlvChassisId) setChoice(value FlowCfmSenderIdTlvChassisIdChoiceEnum) FlowCfmSenderIdTlvChassisId {
	intValue, ok := otg.FlowCfmSenderIdTlvChassisId_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowCfmSenderIdTlvChassisIdChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowCfmSenderIdTlvChassisId_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.LocallyAssigned = nil
	obj.obj.InterfaceName = nil
	obj.obj.NetworkAddress = nil
	obj.networkAddressHolder = nil
	obj.obj.MacAddress = nil
	obj.obj.PortComponent = nil
	obj.obj.InterfaceAlias = nil
	obj.obj.ChassisComponent = nil

	if value == FlowCfmSenderIdTlvChassisIdChoice.CHASSIS_COMPONENT {
		defaultValue := ""
		obj.obj.ChassisComponent = &defaultValue
	}

	if value == FlowCfmSenderIdTlvChassisIdChoice.INTERFACE_ALIAS {
		defaultValue := ""
		obj.obj.InterfaceAlias = &defaultValue
	}

	if value == FlowCfmSenderIdTlvChassisIdChoice.PORT_COMPONENT {
		defaultValue := ""
		obj.obj.PortComponent = &defaultValue
	}

	if value == FlowCfmSenderIdTlvChassisIdChoice.MAC_ADDRESS {
		defaultValue := "00:00:00:00:00:00"
		obj.obj.MacAddress = &defaultValue
	}

	if value == FlowCfmSenderIdTlvChassisIdChoice.NETWORK_ADDRESS {
		obj.obj.NetworkAddress = NewFlowCfmNetwork().msg()
	}

	if value == FlowCfmSenderIdTlvChassisIdChoice.INTERFACE_NAME {
		defaultValue := ""
		obj.obj.InterfaceName = &defaultValue
	}

	if value == FlowCfmSenderIdTlvChassisIdChoice.LOCALLY_ASSIGNED {
		defaultValue := ""
		obj.obj.LocallyAssigned = &defaultValue
	}

	return obj
}

// description is TBD
// Length returns a FlowCfmLength
func (obj *flowCfmSenderIdTlvChassisId) Length() FlowCfmLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewFlowCfmLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &flowCfmLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// description is TBD
// Length returns a FlowCfmLength
func (obj *flowCfmSenderIdTlvChassisId) HasLength() bool {
	return obj.obj.Length != nil
}

// description is TBD
// SetLength sets the FlowCfmLength value in the FlowCfmSenderIdTlvChassisId object
func (obj *flowCfmSenderIdTlvChassisId) SetLength(value FlowCfmLength) FlowCfmSenderIdTlvChassisId {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// Identifies the hardware component FRU/shelf.
// ChassisComponent returns a string
func (obj *flowCfmSenderIdTlvChassisId) ChassisComponent() string {

	if obj.obj.ChassisComponent == nil {
		obj.setChoice(FlowCfmSenderIdTlvChassisIdChoice.CHASSIS_COMPONENT)
	}

	return *obj.obj.ChassisComponent

}

// Identifies the hardware component FRU/shelf.
// ChassisComponent returns a string
func (obj *flowCfmSenderIdTlvChassisId) HasChassisComponent() bool {
	return obj.obj.ChassisComponent != nil
}

// Identifies the hardware component FRU/shelf.
// SetChassisComponent sets the string value in the FlowCfmSenderIdTlvChassisId object
func (obj *flowCfmSenderIdTlvChassisId) SetChassisComponent(value string) FlowCfmSenderIdTlvChassisId {
	obj.setChoice(FlowCfmSenderIdTlvChassisIdChoice.CHASSIS_COMPONENT)
	obj.obj.ChassisComponent = &value
	return obj
}

// Identifies the interface by its manager-assigned name.
// InterfaceAlias returns a string
func (obj *flowCfmSenderIdTlvChassisId) InterfaceAlias() string {

	if obj.obj.InterfaceAlias == nil {
		obj.setChoice(FlowCfmSenderIdTlvChassisIdChoice.INTERFACE_ALIAS)
	}

	return *obj.obj.InterfaceAlias

}

// Identifies the interface by its manager-assigned name.
// InterfaceAlias returns a string
func (obj *flowCfmSenderIdTlvChassisId) HasInterfaceAlias() bool {
	return obj.obj.InterfaceAlias != nil
}

// Identifies the interface by its manager-assigned name.
// SetInterfaceAlias sets the string value in the FlowCfmSenderIdTlvChassisId object
func (obj *flowCfmSenderIdTlvChassisId) SetInterfaceAlias(value string) FlowCfmSenderIdTlvChassisId {
	obj.setChoice(FlowCfmSenderIdTlvChassisIdChoice.INTERFACE_ALIAS)
	obj.obj.InterfaceAlias = &value
	return obj
}

// Identifies a port-class physical component of the chassis.
// PortComponent returns a string
func (obj *flowCfmSenderIdTlvChassisId) PortComponent() string {

	if obj.obj.PortComponent == nil {
		obj.setChoice(FlowCfmSenderIdTlvChassisIdChoice.PORT_COMPONENT)
	}

	return *obj.obj.PortComponent

}

// Identifies a port-class physical component of the chassis.
// PortComponent returns a string
func (obj *flowCfmSenderIdTlvChassisId) HasPortComponent() bool {
	return obj.obj.PortComponent != nil
}

// Identifies a port-class physical component of the chassis.
// SetPortComponent sets the string value in the FlowCfmSenderIdTlvChassisId object
func (obj *flowCfmSenderIdTlvChassisId) SetPortComponent(value string) FlowCfmSenderIdTlvChassisId {
	obj.setChoice(FlowCfmSenderIdTlvChassisIdChoice.PORT_COMPONENT)
	obj.obj.PortComponent = &value
	return obj
}

// MAC address identifying the chassis.
// MacAddress returns a string
func (obj *flowCfmSenderIdTlvChassisId) MacAddress() string {

	if obj.obj.MacAddress == nil {
		obj.setChoice(FlowCfmSenderIdTlvChassisIdChoice.MAC_ADDRESS)
	}

	return *obj.obj.MacAddress

}

// MAC address identifying the chassis.
// MacAddress returns a string
func (obj *flowCfmSenderIdTlvChassisId) HasMacAddress() bool {
	return obj.obj.MacAddress != nil
}

// MAC address identifying the chassis.
// SetMacAddress sets the string value in the FlowCfmSenderIdTlvChassisId object
func (obj *flowCfmSenderIdTlvChassisId) SetMacAddress(value string) FlowCfmSenderIdTlvChassisId {
	obj.setChoice(FlowCfmSenderIdTlvChassisIdChoice.MAC_ADDRESS)
	obj.obj.MacAddress = &value
	return obj
}

// Network address identifying the chassis (IEEE 802.1AB-2016 Chassis ID Subtype 5).
// NetworkAddress returns a FlowCfmNetwork
func (obj *flowCfmSenderIdTlvChassisId) NetworkAddress() FlowCfmNetwork {
	if obj.obj.NetworkAddress == nil {
		obj.setChoice(FlowCfmSenderIdTlvChassisIdChoice.NETWORK_ADDRESS)
	}
	if obj.networkAddressHolder == nil {
		obj.networkAddressHolder = &flowCfmNetwork{obj: obj.obj.NetworkAddress}
	}
	return obj.networkAddressHolder
}

// Network address identifying the chassis (IEEE 802.1AB-2016 Chassis ID Subtype 5).
// NetworkAddress returns a FlowCfmNetwork
func (obj *flowCfmSenderIdTlvChassisId) HasNetworkAddress() bool {
	return obj.obj.NetworkAddress != nil
}

// Network address identifying the chassis (IEEE 802.1AB-2016 Chassis ID Subtype 5).
// SetNetworkAddress sets the FlowCfmNetwork value in the FlowCfmSenderIdTlvChassisId object
func (obj *flowCfmSenderIdTlvChassisId) SetNetworkAddress(value FlowCfmNetwork) FlowCfmSenderIdTlvChassisId {
	obj.setChoice(FlowCfmSenderIdTlvChassisIdChoice.NETWORK_ADDRESS)
	obj.networkAddressHolder = nil
	obj.obj.NetworkAddress = value.msg()

	return obj
}

// Identifies the interface by a system-assigned name.
// InterfaceName returns a string
func (obj *flowCfmSenderIdTlvChassisId) InterfaceName() string {

	if obj.obj.InterfaceName == nil {
		obj.setChoice(FlowCfmSenderIdTlvChassisIdChoice.INTERFACE_NAME)
	}

	return *obj.obj.InterfaceName

}

// Identifies the interface by a system-assigned name.
// InterfaceName returns a string
func (obj *flowCfmSenderIdTlvChassisId) HasInterfaceName() bool {
	return obj.obj.InterfaceName != nil
}

// Identifies the interface by a system-assigned name.
// SetInterfaceName sets the string value in the FlowCfmSenderIdTlvChassisId object
func (obj *flowCfmSenderIdTlvChassisId) SetInterfaceName(value string) FlowCfmSenderIdTlvChassisId {
	obj.setChoice(FlowCfmSenderIdTlvChassisIdChoice.INTERFACE_NAME)
	obj.obj.InterfaceName = &value
	return obj
}

// String assigned by the local administrator.
// LocallyAssigned returns a string
func (obj *flowCfmSenderIdTlvChassisId) LocallyAssigned() string {

	if obj.obj.LocallyAssigned == nil {
		obj.setChoice(FlowCfmSenderIdTlvChassisIdChoice.LOCALLY_ASSIGNED)
	}

	return *obj.obj.LocallyAssigned

}

// String assigned by the local administrator.
// LocallyAssigned returns a string
func (obj *flowCfmSenderIdTlvChassisId) HasLocallyAssigned() bool {
	return obj.obj.LocallyAssigned != nil
}

// String assigned by the local administrator.
// SetLocallyAssigned sets the string value in the FlowCfmSenderIdTlvChassisId object
func (obj *flowCfmSenderIdTlvChassisId) SetLocallyAssigned(value string) FlowCfmSenderIdTlvChassisId {
	obj.setChoice(FlowCfmSenderIdTlvChassisIdChoice.LOCALLY_ASSIGNED)
	obj.obj.LocallyAssigned = &value
	return obj
}

// Management Address Domain and address of the sender. When absent, the Management Address Domain Length field is set to 0 in the Sender ID TLV (IEEE 802.1Q-2018 Section 21.5.5.1).
// ManagementAddressDomain returns a FlowCfmSenderIdTlvMADomain
func (obj *flowCfmSenderIdTlvChassisId) ManagementAddressDomain() FlowCfmSenderIdTlvMADomain {
	if obj.obj.ManagementAddressDomain == nil {
		obj.obj.ManagementAddressDomain = NewFlowCfmSenderIdTlvMADomain().msg()
	}
	if obj.managementAddressDomainHolder == nil {
		obj.managementAddressDomainHolder = &flowCfmSenderIdTlvMADomain{obj: obj.obj.ManagementAddressDomain}
	}
	return obj.managementAddressDomainHolder
}

// Management Address Domain and address of the sender. When absent, the Management Address Domain Length field is set to 0 in the Sender ID TLV (IEEE 802.1Q-2018 Section 21.5.5.1).
// ManagementAddressDomain returns a FlowCfmSenderIdTlvMADomain
func (obj *flowCfmSenderIdTlvChassisId) HasManagementAddressDomain() bool {
	return obj.obj.ManagementAddressDomain != nil
}

// Management Address Domain and address of the sender. When absent, the Management Address Domain Length field is set to 0 in the Sender ID TLV (IEEE 802.1Q-2018 Section 21.5.5.1).
// SetManagementAddressDomain sets the FlowCfmSenderIdTlvMADomain value in the FlowCfmSenderIdTlvChassisId object
func (obj *flowCfmSenderIdTlvChassisId) SetManagementAddressDomain(value FlowCfmSenderIdTlvMADomain) FlowCfmSenderIdTlvChassisId {

	obj.managementAddressDomainHolder = nil
	obj.obj.ManagementAddressDomain = value.msg()

	return obj
}

func (obj *flowCfmSenderIdTlvChassisId) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Length != nil {

		obj.Length().validateObj(vObj, set_default)
	}

	if obj.obj.MacAddress != nil {

		err := obj.validateMac(obj.MacAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowCfmSenderIdTlvChassisId.MacAddress"))
		}

	}

	if obj.obj.NetworkAddress != nil {

		obj.NetworkAddress().validateObj(vObj, set_default)
	}

	if obj.obj.ManagementAddressDomain != nil {

		obj.ManagementAddressDomain().validateObj(vObj, set_default)
	}

}

func (obj *flowCfmSenderIdTlvChassisId) setDefault() {
	var choices_set int = 0
	var choice FlowCfmSenderIdTlvChassisIdChoiceEnum

	if obj.obj.ChassisComponent != nil {
		choices_set += 1
		choice = FlowCfmSenderIdTlvChassisIdChoice.CHASSIS_COMPONENT
	}

	if obj.obj.InterfaceAlias != nil {
		choices_set += 1
		choice = FlowCfmSenderIdTlvChassisIdChoice.INTERFACE_ALIAS
	}

	if obj.obj.PortComponent != nil {
		choices_set += 1
		choice = FlowCfmSenderIdTlvChassisIdChoice.PORT_COMPONENT
	}

	if obj.obj.MacAddress != nil {
		choices_set += 1
		choice = FlowCfmSenderIdTlvChassisIdChoice.MAC_ADDRESS
	}

	if obj.obj.NetworkAddress != nil {
		choices_set += 1
		choice = FlowCfmSenderIdTlvChassisIdChoice.NETWORK_ADDRESS
	}

	if obj.obj.InterfaceName != nil {
		choices_set += 1
		choice = FlowCfmSenderIdTlvChassisIdChoice.INTERFACE_NAME
	}

	if obj.obj.LocallyAssigned != nil {
		choices_set += 1
		choice = FlowCfmSenderIdTlvChassisIdChoice.LOCALLY_ASSIGNED
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowCfmSenderIdTlvChassisIdChoice.MAC_ADDRESS)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowCfmSenderIdTlvChassisId")
			}
		} else {
			intVal := otg.FlowCfmSenderIdTlvChassisId_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowCfmSenderIdTlvChassisId_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
