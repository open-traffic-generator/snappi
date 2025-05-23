package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LldpNeighborsState *****
type lldpNeighborsState struct {
	validation
	obj                *otg.LldpNeighborsState
	marshaller         marshalLldpNeighborsState
	unMarshaller       unMarshalLldpNeighborsState
	customTlvsHolder   LldpNeighborsStateLldpCustomTLVStateIter
	capabilitiesHolder LldpNeighborsStateLldpCapabilityStateIter
}

func NewLldpNeighborsState() LldpNeighborsState {
	obj := lldpNeighborsState{obj: &otg.LldpNeighborsState{}}
	obj.setDefault()
	return &obj
}

func (obj *lldpNeighborsState) msg() *otg.LldpNeighborsState {
	return obj.obj
}

func (obj *lldpNeighborsState) setMsg(msg *otg.LldpNeighborsState) LldpNeighborsState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshallldpNeighborsState struct {
	obj *lldpNeighborsState
}

type marshalLldpNeighborsState interface {
	// ToProto marshals LldpNeighborsState to protobuf object *otg.LldpNeighborsState
	ToProto() (*otg.LldpNeighborsState, error)
	// ToPbText marshals LldpNeighborsState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LldpNeighborsState to YAML text
	ToYaml() (string, error)
	// ToJson marshals LldpNeighborsState to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals LldpNeighborsState to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshallldpNeighborsState struct {
	obj *lldpNeighborsState
}

type unMarshalLldpNeighborsState interface {
	// FromProto unmarshals LldpNeighborsState from protobuf object *otg.LldpNeighborsState
	FromProto(msg *otg.LldpNeighborsState) (LldpNeighborsState, error)
	// FromPbText unmarshals LldpNeighborsState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LldpNeighborsState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LldpNeighborsState from JSON text
	FromJson(value string) error
}

func (obj *lldpNeighborsState) Marshal() marshalLldpNeighborsState {
	if obj.marshaller == nil {
		obj.marshaller = &marshallldpNeighborsState{obj: obj}
	}
	return obj.marshaller
}

func (obj *lldpNeighborsState) Unmarshal() unMarshalLldpNeighborsState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallldpNeighborsState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallldpNeighborsState) ToProto() (*otg.LldpNeighborsState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallldpNeighborsState) FromProto(msg *otg.LldpNeighborsState) (LldpNeighborsState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallldpNeighborsState) ToPbText() (string, error) {
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

func (m *unMarshallldpNeighborsState) FromPbText(value string) error {
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

func (m *marshallldpNeighborsState) ToYaml() (string, error) {
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

func (m *unMarshallldpNeighborsState) FromYaml(value string) error {
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

func (m *marshallldpNeighborsState) ToJsonRaw() (string, error) {
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

func (m *marshallldpNeighborsState) ToJson() (string, error) {
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

func (m *unMarshallldpNeighborsState) FromJson(value string) error {
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

func (obj *lldpNeighborsState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lldpNeighborsState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lldpNeighborsState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lldpNeighborsState) Clone() (LldpNeighborsState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLldpNeighborsState()
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

func (obj *lldpNeighborsState) setNil() {
	obj.customTlvsHolder = nil
	obj.capabilitiesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// LldpNeighborsState is lLDP neighbor information.
type LldpNeighborsState interface {
	Validation
	// msg marshals LldpNeighborsState to protobuf object *otg.LldpNeighborsState
	// and doesn't set defaults
	msg() *otg.LldpNeighborsState
	// setMsg unmarshals LldpNeighborsState from protobuf object *otg.LldpNeighborsState
	// and doesn't set defaults
	setMsg(*otg.LldpNeighborsState) LldpNeighborsState
	// provides marshal interface
	Marshal() marshalLldpNeighborsState
	// provides unmarshal interface
	Unmarshal() unMarshalLldpNeighborsState
	// validate validates LldpNeighborsState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LldpNeighborsState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LldpName returns string, set in LldpNeighborsState.
	LldpName() string
	// SetLldpName assigns string provided by user to LldpNeighborsState
	SetLldpName(value string) LldpNeighborsState
	// HasLldpName checks if LldpName has been set in LldpNeighborsState
	HasLldpName() bool
	// SystemName returns string, set in LldpNeighborsState.
	SystemName() string
	// SetSystemName assigns string provided by user to LldpNeighborsState
	SetSystemName(value string) LldpNeighborsState
	// HasSystemName checks if SystemName has been set in LldpNeighborsState
	HasSystemName() bool
	// SystemDescription returns string, set in LldpNeighborsState.
	SystemDescription() string
	// SetSystemDescription assigns string provided by user to LldpNeighborsState
	SetSystemDescription(value string) LldpNeighborsState
	// HasSystemDescription checks if SystemDescription has been set in LldpNeighborsState
	HasSystemDescription() bool
	// ChassisId returns string, set in LldpNeighborsState.
	ChassisId() string
	// SetChassisId assigns string provided by user to LldpNeighborsState
	SetChassisId(value string) LldpNeighborsState
	// HasChassisId checks if ChassisId has been set in LldpNeighborsState
	HasChassisId() bool
	// ChassisIdType returns LldpNeighborsStateChassisIdTypeEnum, set in LldpNeighborsState
	ChassisIdType() LldpNeighborsStateChassisIdTypeEnum
	// SetChassisIdType assigns LldpNeighborsStateChassisIdTypeEnum provided by user to LldpNeighborsState
	SetChassisIdType(value LldpNeighborsStateChassisIdTypeEnum) LldpNeighborsState
	// HasChassisIdType checks if ChassisIdType has been set in LldpNeighborsState
	HasChassisIdType() bool
	// NeighborId returns string, set in LldpNeighborsState.
	NeighborId() string
	// SetNeighborId assigns string provided by user to LldpNeighborsState
	SetNeighborId(value string) LldpNeighborsState
	// HasNeighborId checks if NeighborId has been set in LldpNeighborsState
	HasNeighborId() bool
	// Age returns uint32, set in LldpNeighborsState.
	Age() uint32
	// SetAge assigns uint32 provided by user to LldpNeighborsState
	SetAge(value uint32) LldpNeighborsState
	// HasAge checks if Age has been set in LldpNeighborsState
	HasAge() bool
	// LastUpdate returns uint32, set in LldpNeighborsState.
	LastUpdate() uint32
	// SetLastUpdate assigns uint32 provided by user to LldpNeighborsState
	SetLastUpdate(value uint32) LldpNeighborsState
	// HasLastUpdate checks if LastUpdate has been set in LldpNeighborsState
	HasLastUpdate() bool
	// Ttl returns uint32, set in LldpNeighborsState.
	Ttl() uint32
	// SetTtl assigns uint32 provided by user to LldpNeighborsState
	SetTtl(value uint32) LldpNeighborsState
	// HasTtl checks if Ttl has been set in LldpNeighborsState
	HasTtl() bool
	// PortId returns string, set in LldpNeighborsState.
	PortId() string
	// SetPortId assigns string provided by user to LldpNeighborsState
	SetPortId(value string) LldpNeighborsState
	// HasPortId checks if PortId has been set in LldpNeighborsState
	HasPortId() bool
	// PortIdType returns LldpNeighborsStatePortIdTypeEnum, set in LldpNeighborsState
	PortIdType() LldpNeighborsStatePortIdTypeEnum
	// SetPortIdType assigns LldpNeighborsStatePortIdTypeEnum provided by user to LldpNeighborsState
	SetPortIdType(value LldpNeighborsStatePortIdTypeEnum) LldpNeighborsState
	// HasPortIdType checks if PortIdType has been set in LldpNeighborsState
	HasPortIdType() bool
	// PortDescription returns string, set in LldpNeighborsState.
	PortDescription() string
	// SetPortDescription assigns string provided by user to LldpNeighborsState
	SetPortDescription(value string) LldpNeighborsState
	// HasPortDescription checks if PortDescription has been set in LldpNeighborsState
	HasPortDescription() bool
	// ManagementAddress returns string, set in LldpNeighborsState.
	ManagementAddress() string
	// SetManagementAddress assigns string provided by user to LldpNeighborsState
	SetManagementAddress(value string) LldpNeighborsState
	// HasManagementAddress checks if ManagementAddress has been set in LldpNeighborsState
	HasManagementAddress() bool
	// ManagementAddressType returns string, set in LldpNeighborsState.
	ManagementAddressType() string
	// SetManagementAddressType assigns string provided by user to LldpNeighborsState
	SetManagementAddressType(value string) LldpNeighborsState
	// HasManagementAddressType checks if ManagementAddressType has been set in LldpNeighborsState
	HasManagementAddressType() bool
	// CustomTlvs returns LldpNeighborsStateLldpCustomTLVStateIterIter, set in LldpNeighborsState
	CustomTlvs() LldpNeighborsStateLldpCustomTLVStateIter
	// Capabilities returns LldpNeighborsStateLldpCapabilityStateIterIter, set in LldpNeighborsState
	Capabilities() LldpNeighborsStateLldpCapabilityStateIter
	setNil()
}

// The name of the LLDP instance.
// LldpName returns a string
func (obj *lldpNeighborsState) LldpName() string {

	return *obj.obj.LldpName

}

// The name of the LLDP instance.
// LldpName returns a string
func (obj *lldpNeighborsState) HasLldpName() bool {
	return obj.obj.LldpName != nil
}

// The name of the LLDP instance.
// SetLldpName sets the string value in the LldpNeighborsState object
func (obj *lldpNeighborsState) SetLldpName(value string) LldpNeighborsState {

	obj.obj.LldpName = &value
	return obj
}

// The system name field shall contain an alpha-numeric string that  indicates the system's administratively assigned name. The system name  should be the system's fully qualified domain name. If implementations  support IETF RFC 3418, the sysName object should be used for this field.
// SystemName returns a string
func (obj *lldpNeighborsState) SystemName() string {

	return *obj.obj.SystemName

}

// The system name field shall contain an alpha-numeric string that  indicates the system's administratively assigned name. The system name  should be the system's fully qualified domain name. If implementations  support IETF RFC 3418, the sysName object should be used for this field.
// SystemName returns a string
func (obj *lldpNeighborsState) HasSystemName() bool {
	return obj.obj.SystemName != nil
}

// The system name field shall contain an alpha-numeric string that  indicates the system's administratively assigned name. The system name  should be the system's fully qualified domain name. If implementations  support IETF RFC 3418, the sysName object should be used for this field.
// SetSystemName sets the string value in the LldpNeighborsState object
func (obj *lldpNeighborsState) SetSystemName(value string) LldpNeighborsState {

	obj.obj.SystemName = &value
	return obj
}

// The system description field shall contain an alpha-numeric string that  is the textual description of the network entity. The system description  should include the full name and version identification of the system's  hardware type, software operating system, and networking software. If  implementations support IETF RFC 3418, the sysDescr object should be used  for this field.
// SystemDescription returns a string
func (obj *lldpNeighborsState) SystemDescription() string {

	return *obj.obj.SystemDescription

}

// The system description field shall contain an alpha-numeric string that  is the textual description of the network entity. The system description  should include the full name and version identification of the system's  hardware type, software operating system, and networking software. If  implementations support IETF RFC 3418, the sysDescr object should be used  for this field.
// SystemDescription returns a string
func (obj *lldpNeighborsState) HasSystemDescription() bool {
	return obj.obj.SystemDescription != nil
}

// The system description field shall contain an alpha-numeric string that  is the textual description of the network entity. The system description  should include the full name and version identification of the system's  hardware type, software operating system, and networking software. If  implementations support IETF RFC 3418, the sysDescr object should be used  for this field.
// SetSystemDescription sets the string value in the LldpNeighborsState object
func (obj *lldpNeighborsState) SetSystemDescription(value string) LldpNeighborsState {

	obj.obj.SystemDescription = &value
	return obj
}

// The Chassis ID is a mandatory TLV which identifies the chassis component of  the endpoint identifier associated with the transmitting LLDP agent.
// ChassisId returns a string
func (obj *lldpNeighborsState) ChassisId() string {

	return *obj.obj.ChassisId

}

// The Chassis ID is a mandatory TLV which identifies the chassis component of  the endpoint identifier associated with the transmitting LLDP agent.
// ChassisId returns a string
func (obj *lldpNeighborsState) HasChassisId() bool {
	return obj.obj.ChassisId != nil
}

// The Chassis ID is a mandatory TLV which identifies the chassis component of  the endpoint identifier associated with the transmitting LLDP agent.
// SetChassisId sets the string value in the LldpNeighborsState object
func (obj *lldpNeighborsState) SetChassisId(value string) LldpNeighborsState {

	obj.obj.ChassisId = &value
	return obj
}

type LldpNeighborsStateChassisIdTypeEnum string

// Enum of ChassisIdType on LldpNeighborsState
var LldpNeighborsStateChassisIdType = struct {
	PORT_COMPONENT    LldpNeighborsStateChassisIdTypeEnum
	NETWORK_ADDRESS   LldpNeighborsStateChassisIdTypeEnum
	CHASSIS_COMPONENT LldpNeighborsStateChassisIdTypeEnum
	MAC_ADDRESS       LldpNeighborsStateChassisIdTypeEnum
	INTERFACE_NAME    LldpNeighborsStateChassisIdTypeEnum
	LOCAL             LldpNeighborsStateChassisIdTypeEnum
	INTERFACE_ALIAS   LldpNeighborsStateChassisIdTypeEnum
}{
	PORT_COMPONENT:    LldpNeighborsStateChassisIdTypeEnum("port_component"),
	NETWORK_ADDRESS:   LldpNeighborsStateChassisIdTypeEnum("network_address"),
	CHASSIS_COMPONENT: LldpNeighborsStateChassisIdTypeEnum("chassis_component"),
	MAC_ADDRESS:       LldpNeighborsStateChassisIdTypeEnum("mac_address"),
	INTERFACE_NAME:    LldpNeighborsStateChassisIdTypeEnum("interface_name"),
	LOCAL:             LldpNeighborsStateChassisIdTypeEnum("local"),
	INTERFACE_ALIAS:   LldpNeighborsStateChassisIdTypeEnum("interface_alias"),
}

func (obj *lldpNeighborsState) ChassisIdType() LldpNeighborsStateChassisIdTypeEnum {
	return LldpNeighborsStateChassisIdTypeEnum(obj.obj.ChassisIdType.Enum().String())
}

// This field identifies the format and source of the chassis identifier string.  It is an enumerator defined by the LldpChassisIdSubtype object from IEEE 802.1AB  MIB.
// ChassisIdType returns a string
func (obj *lldpNeighborsState) HasChassisIdType() bool {
	return obj.obj.ChassisIdType != nil
}

func (obj *lldpNeighborsState) SetChassisIdType(value LldpNeighborsStateChassisIdTypeEnum) LldpNeighborsState {
	intValue, ok := otg.LldpNeighborsState_ChassisIdType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on LldpNeighborsStateChassisIdTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.LldpNeighborsState_ChassisIdType_Enum(intValue)
	obj.obj.ChassisIdType = &enumValue

	return obj
}

// System generated identifier for the neighbor on the LLDP instance.
// NeighborId returns a string
func (obj *lldpNeighborsState) NeighborId() string {

	return *obj.obj.NeighborId

}

// System generated identifier for the neighbor on the LLDP instance.
// NeighborId returns a string
func (obj *lldpNeighborsState) HasNeighborId() bool {
	return obj.obj.NeighborId != nil
}

// System generated identifier for the neighbor on the LLDP instance.
// SetNeighborId sets the string value in the LldpNeighborsState object
func (obj *lldpNeighborsState) SetNeighborId(value string) LldpNeighborsState {

	obj.obj.NeighborId = &value
	return obj
}

// Age since discovery in seconds.
// Age returns a uint32
func (obj *lldpNeighborsState) Age() uint32 {

	return *obj.obj.Age

}

// Age since discovery in seconds.
// Age returns a uint32
func (obj *lldpNeighborsState) HasAge() bool {
	return obj.obj.Age != nil
}

// Age since discovery in seconds.
// SetAge sets the uint32 value in the LldpNeighborsState object
func (obj *lldpNeighborsState) SetAge(value uint32) LldpNeighborsState {

	obj.obj.Age = &value
	return obj
}

// Seconds since last update received.
// LastUpdate returns a uint32
func (obj *lldpNeighborsState) LastUpdate() uint32 {

	return *obj.obj.LastUpdate

}

// Seconds since last update received.
// LastUpdate returns a uint32
func (obj *lldpNeighborsState) HasLastUpdate() bool {
	return obj.obj.LastUpdate != nil
}

// Seconds since last update received.
// SetLastUpdate sets the uint32 value in the LldpNeighborsState object
func (obj *lldpNeighborsState) SetLastUpdate(value uint32) LldpNeighborsState {

	obj.obj.LastUpdate = &value
	return obj
}

// The time-to-live (TTL) in seconds is a mandatory TLV which indicates how long information from the neighbor  should be considered valid.
// Ttl returns a uint32
func (obj *lldpNeighborsState) Ttl() uint32 {

	return *obj.obj.Ttl

}

// The time-to-live (TTL) in seconds is a mandatory TLV which indicates how long information from the neighbor  should be considered valid.
// Ttl returns a uint32
func (obj *lldpNeighborsState) HasTtl() bool {
	return obj.obj.Ttl != nil
}

// The time-to-live (TTL) in seconds is a mandatory TLV which indicates how long information from the neighbor  should be considered valid.
// SetTtl sets the uint32 value in the LldpNeighborsState object
func (obj *lldpNeighborsState) SetTtl(value uint32) LldpNeighborsState {

	obj.obj.Ttl = &value
	return obj
}

// The Port ID is a mandatory TLV which identifies the port component of the endpoint identifier associated with  the transmitting LLDP agent. If the specified port is an IEEE 802.3 Repeater port, then this TLV is optional.
// PortId returns a string
func (obj *lldpNeighborsState) PortId() string {

	return *obj.obj.PortId

}

// The Port ID is a mandatory TLV which identifies the port component of the endpoint identifier associated with  the transmitting LLDP agent. If the specified port is an IEEE 802.3 Repeater port, then this TLV is optional.
// PortId returns a string
func (obj *lldpNeighborsState) HasPortId() bool {
	return obj.obj.PortId != nil
}

// The Port ID is a mandatory TLV which identifies the port component of the endpoint identifier associated with  the transmitting LLDP agent. If the specified port is an IEEE 802.3 Repeater port, then this TLV is optional.
// SetPortId sets the string value in the LldpNeighborsState object
func (obj *lldpNeighborsState) SetPortId(value string) LldpNeighborsState {

	obj.obj.PortId = &value
	return obj
}

type LldpNeighborsStatePortIdTypeEnum string

// Enum of PortIdType on LldpNeighborsState
var LldpNeighborsStatePortIdType = struct {
	PORT_COMPONENT   LldpNeighborsStatePortIdTypeEnum
	NETWORK_ADDRESS  LldpNeighborsStatePortIdTypeEnum
	AGENT_CIRCUIT_ID LldpNeighborsStatePortIdTypeEnum
	MAC_ADDRESS      LldpNeighborsStatePortIdTypeEnum
	INTERFACE_NAME   LldpNeighborsStatePortIdTypeEnum
	LOCAL            LldpNeighborsStatePortIdTypeEnum
	INTERFACE_ALIAS  LldpNeighborsStatePortIdTypeEnum
}{
	PORT_COMPONENT:   LldpNeighborsStatePortIdTypeEnum("port_component"),
	NETWORK_ADDRESS:  LldpNeighborsStatePortIdTypeEnum("network_address"),
	AGENT_CIRCUIT_ID: LldpNeighborsStatePortIdTypeEnum("agent_circuit_id"),
	MAC_ADDRESS:      LldpNeighborsStatePortIdTypeEnum("mac_address"),
	INTERFACE_NAME:   LldpNeighborsStatePortIdTypeEnum("interface_name"),
	LOCAL:            LldpNeighborsStatePortIdTypeEnum("local"),
	INTERFACE_ALIAS:  LldpNeighborsStatePortIdTypeEnum("interface_alias"),
}

func (obj *lldpNeighborsState) PortIdType() LldpNeighborsStatePortIdTypeEnum {
	return LldpNeighborsStatePortIdTypeEnum(obj.obj.PortIdType.Enum().String())
}

// This field identifies the format and source of the port identifier string. It is an enumerator defined by the PtopoPortIdType object from RFC2922.
// PortIdType returns a string
func (obj *lldpNeighborsState) HasPortIdType() bool {
	return obj.obj.PortIdType != nil
}

func (obj *lldpNeighborsState) SetPortIdType(value LldpNeighborsStatePortIdTypeEnum) LldpNeighborsState {
	intValue, ok := otg.LldpNeighborsState_PortIdType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on LldpNeighborsStatePortIdTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.LldpNeighborsState_PortIdType_Enum(intValue)
	obj.obj.PortIdType = &enumValue

	return obj
}

// The binary string containing the actual port identifier for the port which this LLDP PDU was transmitted. The source  and format of this field is defined by PtopoPortId from RFC2922.
// PortDescription returns a string
func (obj *lldpNeighborsState) PortDescription() string {

	return *obj.obj.PortDescription

}

// The binary string containing the actual port identifier for the port which this LLDP PDU was transmitted. The source  and format of this field is defined by PtopoPortId from RFC2922.
// PortDescription returns a string
func (obj *lldpNeighborsState) HasPortDescription() bool {
	return obj.obj.PortDescription != nil
}

// The binary string containing the actual port identifier for the port which this LLDP PDU was transmitted. The source  and format of this field is defined by PtopoPortId from RFC2922.
// SetPortDescription sets the string value in the LldpNeighborsState object
func (obj *lldpNeighborsState) SetPortDescription(value string) LldpNeighborsState {

	obj.obj.PortDescription = &value
	return obj
}

// The Management Address is a mandatory TLV which identifies a network address associated with the local LLDP agent, which  can be used to reach the agent on the port identified in the Port ID TLV.
// ManagementAddress returns a string
func (obj *lldpNeighborsState) ManagementAddress() string {

	return *obj.obj.ManagementAddress

}

// The Management Address is a mandatory TLV which identifies a network address associated with the local LLDP agent, which  can be used to reach the agent on the port identified in the Port ID TLV.
// ManagementAddress returns a string
func (obj *lldpNeighborsState) HasManagementAddress() bool {
	return obj.obj.ManagementAddress != nil
}

// The Management Address is a mandatory TLV which identifies a network address associated with the local LLDP agent, which  can be used to reach the agent on the port identified in the Port ID TLV.
// SetManagementAddress sets the string value in the LldpNeighborsState object
func (obj *lldpNeighborsState) SetManagementAddress(value string) LldpNeighborsState {

	obj.obj.ManagementAddress = &value
	return obj
}

// The enumerated value for the network address type identified in this TLV. This enumeration is defined in the 'Assigned Numbers'  RFC [RFC3232] and the ianaAddressFamilyNumbers object.
// ManagementAddressType returns a string
func (obj *lldpNeighborsState) ManagementAddressType() string {

	return *obj.obj.ManagementAddressType

}

// The enumerated value for the network address type identified in this TLV. This enumeration is defined in the 'Assigned Numbers'  RFC [RFC3232] and the ianaAddressFamilyNumbers object.
// ManagementAddressType returns a string
func (obj *lldpNeighborsState) HasManagementAddressType() bool {
	return obj.obj.ManagementAddressType != nil
}

// The enumerated value for the network address type identified in this TLV. This enumeration is defined in the 'Assigned Numbers'  RFC [RFC3232] and the ianaAddressFamilyNumbers object.
// SetManagementAddressType sets the string value in the LldpNeighborsState object
func (obj *lldpNeighborsState) SetManagementAddressType(value string) LldpNeighborsState {

	obj.obj.ManagementAddressType = &value
	return obj
}

// description is TBD
// CustomTlvs returns a []LldpCustomTLVState
func (obj *lldpNeighborsState) CustomTlvs() LldpNeighborsStateLldpCustomTLVStateIter {
	if len(obj.obj.CustomTlvs) == 0 {
		obj.obj.CustomTlvs = []*otg.LldpCustomTLVState{}
	}
	if obj.customTlvsHolder == nil {
		obj.customTlvsHolder = newLldpNeighborsStateLldpCustomTLVStateIter(&obj.obj.CustomTlvs).setMsg(obj)
	}
	return obj.customTlvsHolder
}

type lldpNeighborsStateLldpCustomTLVStateIter struct {
	obj                     *lldpNeighborsState
	lldpCustomTLVStateSlice []LldpCustomTLVState
	fieldPtr                *[]*otg.LldpCustomTLVState
}

func newLldpNeighborsStateLldpCustomTLVStateIter(ptr *[]*otg.LldpCustomTLVState) LldpNeighborsStateLldpCustomTLVStateIter {
	return &lldpNeighborsStateLldpCustomTLVStateIter{fieldPtr: ptr}
}

type LldpNeighborsStateLldpCustomTLVStateIter interface {
	setMsg(*lldpNeighborsState) LldpNeighborsStateLldpCustomTLVStateIter
	Items() []LldpCustomTLVState
	Add() LldpCustomTLVState
	Append(items ...LldpCustomTLVState) LldpNeighborsStateLldpCustomTLVStateIter
	Set(index int, newObj LldpCustomTLVState) LldpNeighborsStateLldpCustomTLVStateIter
	Clear() LldpNeighborsStateLldpCustomTLVStateIter
	clearHolderSlice() LldpNeighborsStateLldpCustomTLVStateIter
	appendHolderSlice(item LldpCustomTLVState) LldpNeighborsStateLldpCustomTLVStateIter
}

func (obj *lldpNeighborsStateLldpCustomTLVStateIter) setMsg(msg *lldpNeighborsState) LldpNeighborsStateLldpCustomTLVStateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&lldpCustomTLVState{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *lldpNeighborsStateLldpCustomTLVStateIter) Items() []LldpCustomTLVState {
	return obj.lldpCustomTLVStateSlice
}

func (obj *lldpNeighborsStateLldpCustomTLVStateIter) Add() LldpCustomTLVState {
	newObj := &otg.LldpCustomTLVState{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &lldpCustomTLVState{obj: newObj}
	newLibObj.setDefault()
	obj.lldpCustomTLVStateSlice = append(obj.lldpCustomTLVStateSlice, newLibObj)
	return newLibObj
}

func (obj *lldpNeighborsStateLldpCustomTLVStateIter) Append(items ...LldpCustomTLVState) LldpNeighborsStateLldpCustomTLVStateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.lldpCustomTLVStateSlice = append(obj.lldpCustomTLVStateSlice, item)
	}
	return obj
}

func (obj *lldpNeighborsStateLldpCustomTLVStateIter) Set(index int, newObj LldpCustomTLVState) LldpNeighborsStateLldpCustomTLVStateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.lldpCustomTLVStateSlice[index] = newObj
	return obj
}
func (obj *lldpNeighborsStateLldpCustomTLVStateIter) Clear() LldpNeighborsStateLldpCustomTLVStateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.LldpCustomTLVState{}
		obj.lldpCustomTLVStateSlice = []LldpCustomTLVState{}
	}
	return obj
}
func (obj *lldpNeighborsStateLldpCustomTLVStateIter) clearHolderSlice() LldpNeighborsStateLldpCustomTLVStateIter {
	if len(obj.lldpCustomTLVStateSlice) > 0 {
		obj.lldpCustomTLVStateSlice = []LldpCustomTLVState{}
	}
	return obj
}
func (obj *lldpNeighborsStateLldpCustomTLVStateIter) appendHolderSlice(item LldpCustomTLVState) LldpNeighborsStateLldpCustomTLVStateIter {
	obj.lldpCustomTLVStateSlice = append(obj.lldpCustomTLVStateSlice, item)
	return obj
}

// description is TBD
// Capabilities returns a []LldpCapabilityState
func (obj *lldpNeighborsState) Capabilities() LldpNeighborsStateLldpCapabilityStateIter {
	if len(obj.obj.Capabilities) == 0 {
		obj.obj.Capabilities = []*otg.LldpCapabilityState{}
	}
	if obj.capabilitiesHolder == nil {
		obj.capabilitiesHolder = newLldpNeighborsStateLldpCapabilityStateIter(&obj.obj.Capabilities).setMsg(obj)
	}
	return obj.capabilitiesHolder
}

type lldpNeighborsStateLldpCapabilityStateIter struct {
	obj                      *lldpNeighborsState
	lldpCapabilityStateSlice []LldpCapabilityState
	fieldPtr                 *[]*otg.LldpCapabilityState
}

func newLldpNeighborsStateLldpCapabilityStateIter(ptr *[]*otg.LldpCapabilityState) LldpNeighborsStateLldpCapabilityStateIter {
	return &lldpNeighborsStateLldpCapabilityStateIter{fieldPtr: ptr}
}

type LldpNeighborsStateLldpCapabilityStateIter interface {
	setMsg(*lldpNeighborsState) LldpNeighborsStateLldpCapabilityStateIter
	Items() []LldpCapabilityState
	Add() LldpCapabilityState
	Append(items ...LldpCapabilityState) LldpNeighborsStateLldpCapabilityStateIter
	Set(index int, newObj LldpCapabilityState) LldpNeighborsStateLldpCapabilityStateIter
	Clear() LldpNeighborsStateLldpCapabilityStateIter
	clearHolderSlice() LldpNeighborsStateLldpCapabilityStateIter
	appendHolderSlice(item LldpCapabilityState) LldpNeighborsStateLldpCapabilityStateIter
}

func (obj *lldpNeighborsStateLldpCapabilityStateIter) setMsg(msg *lldpNeighborsState) LldpNeighborsStateLldpCapabilityStateIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&lldpCapabilityState{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *lldpNeighborsStateLldpCapabilityStateIter) Items() []LldpCapabilityState {
	return obj.lldpCapabilityStateSlice
}

func (obj *lldpNeighborsStateLldpCapabilityStateIter) Add() LldpCapabilityState {
	newObj := &otg.LldpCapabilityState{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &lldpCapabilityState{obj: newObj}
	newLibObj.setDefault()
	obj.lldpCapabilityStateSlice = append(obj.lldpCapabilityStateSlice, newLibObj)
	return newLibObj
}

func (obj *lldpNeighborsStateLldpCapabilityStateIter) Append(items ...LldpCapabilityState) LldpNeighborsStateLldpCapabilityStateIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.lldpCapabilityStateSlice = append(obj.lldpCapabilityStateSlice, item)
	}
	return obj
}

func (obj *lldpNeighborsStateLldpCapabilityStateIter) Set(index int, newObj LldpCapabilityState) LldpNeighborsStateLldpCapabilityStateIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.lldpCapabilityStateSlice[index] = newObj
	return obj
}
func (obj *lldpNeighborsStateLldpCapabilityStateIter) Clear() LldpNeighborsStateLldpCapabilityStateIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.LldpCapabilityState{}
		obj.lldpCapabilityStateSlice = []LldpCapabilityState{}
	}
	return obj
}
func (obj *lldpNeighborsStateLldpCapabilityStateIter) clearHolderSlice() LldpNeighborsStateLldpCapabilityStateIter {
	if len(obj.lldpCapabilityStateSlice) > 0 {
		obj.lldpCapabilityStateSlice = []LldpCapabilityState{}
	}
	return obj
}
func (obj *lldpNeighborsStateLldpCapabilityStateIter) appendHolderSlice(item LldpCapabilityState) LldpNeighborsStateLldpCapabilityStateIter {
	obj.lldpCapabilityStateSlice = append(obj.lldpCapabilityStateSlice, item)
	return obj
}

func (obj *lldpNeighborsState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.CustomTlvs) != 0 {

		if set_default {
			obj.CustomTlvs().clearHolderSlice()
			for _, item := range obj.obj.CustomTlvs {
				obj.CustomTlvs().appendHolderSlice(&lldpCustomTLVState{obj: item})
			}
		}
		for _, item := range obj.CustomTlvs().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Capabilities) != 0 {

		if set_default {
			obj.Capabilities().clearHolderSlice()
			for _, item := range obj.obj.Capabilities {
				obj.Capabilities().appendHolderSlice(&lldpCapabilityState{obj: item})
			}
		}
		for _, item := range obj.Capabilities().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *lldpNeighborsState) setDefault() {

}
