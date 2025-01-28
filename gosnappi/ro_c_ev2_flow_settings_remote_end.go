package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RoCEv2FlowSettingsRemoteEnd *****
type roCEv2FlowSettingsRemoteEnd struct {
	validation
	obj          *otg.RoCEv2FlowSettingsRemoteEnd
	marshaller   marshalRoCEv2FlowSettingsRemoteEnd
	unMarshaller unMarshalRoCEv2FlowSettingsRemoteEnd
}

func NewRoCEv2FlowSettingsRemoteEnd() RoCEv2FlowSettingsRemoteEnd {
	obj := roCEv2FlowSettingsRemoteEnd{obj: &otg.RoCEv2FlowSettingsRemoteEnd{}}
	obj.setDefault()
	return &obj
}

func (obj *roCEv2FlowSettingsRemoteEnd) msg() *otg.RoCEv2FlowSettingsRemoteEnd {
	return obj.obj
}

func (obj *roCEv2FlowSettingsRemoteEnd) setMsg(msg *otg.RoCEv2FlowSettingsRemoteEnd) RoCEv2FlowSettingsRemoteEnd {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalroCEv2FlowSettingsRemoteEnd struct {
	obj *roCEv2FlowSettingsRemoteEnd
}

type marshalRoCEv2FlowSettingsRemoteEnd interface {
	// ToProto marshals RoCEv2FlowSettingsRemoteEnd to protobuf object *otg.RoCEv2FlowSettingsRemoteEnd
	ToProto() (*otg.RoCEv2FlowSettingsRemoteEnd, error)
	// ToPbText marshals RoCEv2FlowSettingsRemoteEnd to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RoCEv2FlowSettingsRemoteEnd to YAML text
	ToYaml() (string, error)
	// ToJson marshals RoCEv2FlowSettingsRemoteEnd to JSON text
	ToJson() (string, error)
}

type unMarshalroCEv2FlowSettingsRemoteEnd struct {
	obj *roCEv2FlowSettingsRemoteEnd
}

type unMarshalRoCEv2FlowSettingsRemoteEnd interface {
	// FromProto unmarshals RoCEv2FlowSettingsRemoteEnd from protobuf object *otg.RoCEv2FlowSettingsRemoteEnd
	FromProto(msg *otg.RoCEv2FlowSettingsRemoteEnd) (RoCEv2FlowSettingsRemoteEnd, error)
	// FromPbText unmarshals RoCEv2FlowSettingsRemoteEnd from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RoCEv2FlowSettingsRemoteEnd from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RoCEv2FlowSettingsRemoteEnd from JSON text
	FromJson(value string) error
}

func (obj *roCEv2FlowSettingsRemoteEnd) Marshal() marshalRoCEv2FlowSettingsRemoteEnd {
	if obj.marshaller == nil {
		obj.marshaller = &marshalroCEv2FlowSettingsRemoteEnd{obj: obj}
	}
	return obj.marshaller
}

func (obj *roCEv2FlowSettingsRemoteEnd) Unmarshal() unMarshalRoCEv2FlowSettingsRemoteEnd {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalroCEv2FlowSettingsRemoteEnd{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalroCEv2FlowSettingsRemoteEnd) ToProto() (*otg.RoCEv2FlowSettingsRemoteEnd, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalroCEv2FlowSettingsRemoteEnd) FromProto(msg *otg.RoCEv2FlowSettingsRemoteEnd) (RoCEv2FlowSettingsRemoteEnd, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalroCEv2FlowSettingsRemoteEnd) ToPbText() (string, error) {
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

func (m *unMarshalroCEv2FlowSettingsRemoteEnd) FromPbText(value string) error {
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

func (m *marshalroCEv2FlowSettingsRemoteEnd) ToYaml() (string, error) {
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

func (m *unMarshalroCEv2FlowSettingsRemoteEnd) FromYaml(value string) error {
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

func (m *marshalroCEv2FlowSettingsRemoteEnd) ToJson() (string, error) {
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

func (m *unMarshalroCEv2FlowSettingsRemoteEnd) FromJson(value string) error {
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

func (obj *roCEv2FlowSettingsRemoteEnd) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *roCEv2FlowSettingsRemoteEnd) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *roCEv2FlowSettingsRemoteEnd) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *roCEv2FlowSettingsRemoteEnd) Clone() (RoCEv2FlowSettingsRemoteEnd, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRoCEv2FlowSettingsRemoteEnd()
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

// RoCEv2FlowSettingsRemoteEnd is remote End Flow Settings
type RoCEv2FlowSettingsRemoteEnd interface {
	Validation
	// msg marshals RoCEv2FlowSettingsRemoteEnd to protobuf object *otg.RoCEv2FlowSettingsRemoteEnd
	// and doesn't set defaults
	msg() *otg.RoCEv2FlowSettingsRemoteEnd
	// setMsg unmarshals RoCEv2FlowSettingsRemoteEnd from protobuf object *otg.RoCEv2FlowSettingsRemoteEnd
	// and doesn't set defaults
	setMsg(*otg.RoCEv2FlowSettingsRemoteEnd) RoCEv2FlowSettingsRemoteEnd
	// provides marshal interface
	Marshal() marshalRoCEv2FlowSettingsRemoteEnd
	// provides unmarshal interface
	Unmarshal() unMarshalRoCEv2FlowSettingsRemoteEnd
	// validate validates RoCEv2FlowSettingsRemoteEnd
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RoCEv2FlowSettingsRemoteEnd, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// CustomQp returns bool, set in RoCEv2FlowSettingsRemoteEnd.
	CustomQp() bool
	// SetCustomQp assigns bool provided by user to RoCEv2FlowSettingsRemoteEnd
	SetCustomQp(value bool) RoCEv2FlowSettingsRemoteEnd
	// HasCustomQp checks if CustomQp has been set in RoCEv2FlowSettingsRemoteEnd
	HasCustomQp() bool
	// CustomQpNumber returns uint32, set in RoCEv2FlowSettingsRemoteEnd.
	CustomQpNumber() uint32
	// SetCustomQpNumber assigns uint32 provided by user to RoCEv2FlowSettingsRemoteEnd
	SetCustomQpNumber(value uint32) RoCEv2FlowSettingsRemoteEnd
	// HasCustomQpNumber checks if CustomQpNumber has been set in RoCEv2FlowSettingsRemoteEnd
	HasCustomQpNumber() bool
	// Dscp returns uint32, set in RoCEv2FlowSettingsRemoteEnd.
	Dscp() uint32
	// SetDscp assigns uint32 provided by user to RoCEv2FlowSettingsRemoteEnd
	SetDscp(value uint32) RoCEv2FlowSettingsRemoteEnd
	// HasDscp checks if Dscp has been set in RoCEv2FlowSettingsRemoteEnd
	HasDscp() bool
	// Ecn returns uint32, set in RoCEv2FlowSettingsRemoteEnd.
	Ecn() uint32
	// SetEcn assigns uint32 provided by user to RoCEv2FlowSettingsRemoteEnd
	SetEcn(value uint32) RoCEv2FlowSettingsRemoteEnd
	// HasEcn checks if Ecn has been set in RoCEv2FlowSettingsRemoteEnd
	HasEcn() bool
	// UdpSourcePort returns uint32, set in RoCEv2FlowSettingsRemoteEnd.
	UdpSourcePort() uint32
	// SetUdpSourcePort assigns uint32 provided by user to RoCEv2FlowSettingsRemoteEnd
	SetUdpSourcePort(value uint32) RoCEv2FlowSettingsRemoteEnd
	// HasUdpSourcePort checks if UdpSourcePort has been set in RoCEv2FlowSettingsRemoteEnd
	HasUdpSourcePort() bool
	// ExecuteCommands returns RoCEv2FlowSettingsRemoteEndExecuteCommandsEnum, set in RoCEv2FlowSettingsRemoteEnd
	ExecuteCommands() RoCEv2FlowSettingsRemoteEndExecuteCommandsEnum
	// SetExecuteCommands assigns RoCEv2FlowSettingsRemoteEndExecuteCommandsEnum provided by user to RoCEv2FlowSettingsRemoteEnd
	SetExecuteCommands(value RoCEv2FlowSettingsRemoteEndExecuteCommandsEnum) RoCEv2FlowSettingsRemoteEnd
	// HasExecuteCommands checks if ExecuteCommands has been set in RoCEv2FlowSettingsRemoteEnd
	HasExecuteCommands() bool
	// ImmediateData returns string, set in RoCEv2FlowSettingsRemoteEnd.
	ImmediateData() string
	// SetImmediateData assigns string provided by user to RoCEv2FlowSettingsRemoteEnd
	SetImmediateData(value string) RoCEv2FlowSettingsRemoteEnd
	// HasImmediateData checks if ImmediateData has been set in RoCEv2FlowSettingsRemoteEnd
	HasImmediateData() bool
	// MessageSize returns uint32, set in RoCEv2FlowSettingsRemoteEnd.
	MessageSize() uint32
	// SetMessageSize assigns uint32 provided by user to RoCEv2FlowSettingsRemoteEnd
	SetMessageSize(value uint32) RoCEv2FlowSettingsRemoteEnd
	// HasMessageSize checks if MessageSize has been set in RoCEv2FlowSettingsRemoteEnd
	HasMessageSize() bool
	// MessageSizeUnit returns RoCEv2FlowSettingsRemoteEndMessageSizeUnitEnum, set in RoCEv2FlowSettingsRemoteEnd
	MessageSizeUnit() RoCEv2FlowSettingsRemoteEndMessageSizeUnitEnum
	// SetMessageSizeUnit assigns RoCEv2FlowSettingsRemoteEndMessageSizeUnitEnum provided by user to RoCEv2FlowSettingsRemoteEnd
	SetMessageSizeUnit(value RoCEv2FlowSettingsRemoteEndMessageSizeUnitEnum) RoCEv2FlowSettingsRemoteEnd
	// HasMessageSizeUnit checks if MessageSizeUnit has been set in RoCEv2FlowSettingsRemoteEnd
	HasMessageSizeUnit() bool
	// RemoteQpNumber returns uint32, set in RoCEv2FlowSettingsRemoteEnd.
	RemoteQpNumber() uint32
	// SetRemoteQpNumber assigns uint32 provided by user to RoCEv2FlowSettingsRemoteEnd
	SetRemoteQpNumber(value uint32) RoCEv2FlowSettingsRemoteEnd
	// HasRemoteQpNumber checks if RemoteQpNumber has been set in RoCEv2FlowSettingsRemoteEnd
	HasRemoteQpNumber() bool
	// VirtualAddress returns string, set in RoCEv2FlowSettingsRemoteEnd.
	VirtualAddress() string
	// SetVirtualAddress assigns string provided by user to RoCEv2FlowSettingsRemoteEnd
	SetVirtualAddress(value string) RoCEv2FlowSettingsRemoteEnd
	// HasVirtualAddress checks if VirtualAddress has been set in RoCEv2FlowSettingsRemoteEnd
	HasVirtualAddress() bool
	// RemoteKey returns string, set in RoCEv2FlowSettingsRemoteEnd.
	RemoteKey() string
	// SetRemoteKey assigns string provided by user to RoCEv2FlowSettingsRemoteEnd
	SetRemoteKey(value string) RoCEv2FlowSettingsRemoteEnd
	// HasRemoteKey checks if RemoteKey has been set in RoCEv2FlowSettingsRemoteEnd
	HasRemoteKey() bool
	// InitialPsn returns uint32, set in RoCEv2FlowSettingsRemoteEnd.
	InitialPsn() uint32
	// SetInitialPsn assigns uint32 provided by user to RoCEv2FlowSettingsRemoteEnd
	SetInitialPsn(value uint32) RoCEv2FlowSettingsRemoteEnd
	// HasInitialPsn checks if InitialPsn has been set in RoCEv2FlowSettingsRemoteEnd
	HasInitialPsn() bool
}

// Turn on to define QP number.
// CustomQp returns a bool
func (obj *roCEv2FlowSettingsRemoteEnd) CustomQp() bool {

	return *obj.obj.CustomQp

}

// Turn on to define QP number.
// CustomQp returns a bool
func (obj *roCEv2FlowSettingsRemoteEnd) HasCustomQp() bool {
	return obj.obj.CustomQp != nil
}

// Turn on to define QP number.
// SetCustomQp sets the bool value in the RoCEv2FlowSettingsRemoteEnd object
func (obj *roCEv2FlowSettingsRemoteEnd) SetCustomQp(value bool) RoCEv2FlowSettingsRemoteEnd {

	obj.obj.CustomQp = &value
	return obj
}

// Configure the QP range.
// CustomQpNumber returns a uint32
func (obj *roCEv2FlowSettingsRemoteEnd) CustomQpNumber() uint32 {

	return *obj.obj.CustomQpNumber

}

// Configure the QP range.
// CustomQpNumber returns a uint32
func (obj *roCEv2FlowSettingsRemoteEnd) HasCustomQpNumber() bool {
	return obj.obj.CustomQpNumber != nil
}

// Configure the QP range.
// SetCustomQpNumber sets the uint32 value in the RoCEv2FlowSettingsRemoteEnd object
func (obj *roCEv2FlowSettingsRemoteEnd) SetCustomQpNumber(value uint32) RoCEv2FlowSettingsRemoteEnd {

	obj.obj.CustomQpNumber = &value
	return obj
}

// DSCP value for this flow
// Dscp returns a uint32
func (obj *roCEv2FlowSettingsRemoteEnd) Dscp() uint32 {

	return *obj.obj.Dscp

}

// DSCP value for this flow
// Dscp returns a uint32
func (obj *roCEv2FlowSettingsRemoteEnd) HasDscp() bool {
	return obj.obj.Dscp != nil
}

// DSCP value for this flow
// SetDscp sets the uint32 value in the RoCEv2FlowSettingsRemoteEnd object
func (obj *roCEv2FlowSettingsRemoteEnd) SetDscp(value uint32) RoCEv2FlowSettingsRemoteEnd {

	obj.obj.Dscp = &value
	return obj
}

// This field allows to configure bits of the Traffic Class field in the IPv4 or IPv6 header to encode four different code points.
// Ecn returns a uint32
func (obj *roCEv2FlowSettingsRemoteEnd) Ecn() uint32 {

	return *obj.obj.Ecn

}

// This field allows to configure bits of the Traffic Class field in the IPv4 or IPv6 header to encode four different code points.
// Ecn returns a uint32
func (obj *roCEv2FlowSettingsRemoteEnd) HasEcn() bool {
	return obj.obj.Ecn != nil
}

// This field allows to configure bits of the Traffic Class field in the IPv4 or IPv6 header to encode four different code points.
// SetEcn sets the uint32 value in the RoCEv2FlowSettingsRemoteEnd object
func (obj *roCEv2FlowSettingsRemoteEnd) SetEcn(value uint32) RoCEv2FlowSettingsRemoteEnd {

	obj.obj.Ecn = &value
	return obj
}

// UDP source port number for this flow.
// UdpSourcePort returns a uint32
func (obj *roCEv2FlowSettingsRemoteEnd) UdpSourcePort() uint32 {

	return *obj.obj.UdpSourcePort

}

// UDP source port number for this flow.
// UdpSourcePort returns a uint32
func (obj *roCEv2FlowSettingsRemoteEnd) HasUdpSourcePort() bool {
	return obj.obj.UdpSourcePort != nil
}

// UDP source port number for this flow.
// SetUdpSourcePort sets the uint32 value in the RoCEv2FlowSettingsRemoteEnd object
func (obj *roCEv2FlowSettingsRemoteEnd) SetUdpSourcePort(value uint32) RoCEv2FlowSettingsRemoteEnd {

	obj.obj.UdpSourcePort = &value
	return obj
}

type RoCEv2FlowSettingsRemoteEndExecuteCommandsEnum string

// Enum of ExecuteCommands on RoCEv2FlowSettingsRemoteEnd
var RoCEv2FlowSettingsRemoteEndExecuteCommands = struct {
	NONE      RoCEv2FlowSettingsRemoteEndExecuteCommandsEnum
	RDMAWRITE RoCEv2FlowSettingsRemoteEndExecuteCommandsEnum
}{
	NONE:      RoCEv2FlowSettingsRemoteEndExecuteCommandsEnum("none"),
	RDMAWRITE: RoCEv2FlowSettingsRemoteEndExecuteCommandsEnum("RDMAWrite"),
}

func (obj *roCEv2FlowSettingsRemoteEnd) ExecuteCommands() RoCEv2FlowSettingsRemoteEndExecuteCommandsEnum {
	return RoCEv2FlowSettingsRemoteEndExecuteCommandsEnum(obj.obj.ExecuteCommands.Enum().String())
}

// UDP source port number for this flow.
// ExecuteCommands returns a string
func (obj *roCEv2FlowSettingsRemoteEnd) HasExecuteCommands() bool {
	return obj.obj.ExecuteCommands != nil
}

func (obj *roCEv2FlowSettingsRemoteEnd) SetExecuteCommands(value RoCEv2FlowSettingsRemoteEndExecuteCommandsEnum) RoCEv2FlowSettingsRemoteEnd {
	intValue, ok := otg.RoCEv2FlowSettingsRemoteEnd_ExecuteCommands_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on RoCEv2FlowSettingsRemoteEndExecuteCommandsEnum", string(value)))
		return obj
	}
	enumValue := otg.RoCEv2FlowSettingsRemoteEnd_ExecuteCommands_Enum(intValue)
	obj.obj.ExecuteCommands = &enumValue

	return obj
}

// Immediate Data field required for SEND/WRITE with immediate verb.
// ImmediateData returns a string
func (obj *roCEv2FlowSettingsRemoteEnd) ImmediateData() string {

	return *obj.obj.ImmediateData

}

// Immediate Data field required for SEND/WRITE with immediate verb.
// ImmediateData returns a string
func (obj *roCEv2FlowSettingsRemoteEnd) HasImmediateData() bool {
	return obj.obj.ImmediateData != nil
}

// Immediate Data field required for SEND/WRITE with immediate verb.
// SetImmediateData sets the string value in the RoCEv2FlowSettingsRemoteEnd object
func (obj *roCEv2FlowSettingsRemoteEnd) SetImmediateData(value string) RoCEv2FlowSettingsRemoteEnd {

	obj.obj.ImmediateData = &value
	return obj
}

// The Maximum message size that is allowed to transfer depends on the MTU size and the number of VLANs configured on the interfaces. For any MTU, the maximum number of segmented RDMA WRITE packets for a single WRITE MESSAGE is 65535. That is, a single RDMA WRITE message can be broken into 1 WRITE FIRST, 1 WRITE LAST and (65535-2) WRITE MIDDLE messages. The maximum message size that is allowed to be transferred for a given MTU is constrained by the above conditions. For example, for an MTU size of 1500 bytes, the common header of the RDMA WRITE MIDDLE/LAST will comprise of Ethernet Header + IP Header + UDP Header + BTH Header + iCRC size + Ethernet Trailer size. This works out to be 14+20+8+12+4+4 = 62 bytes. For RDMA WRITE FIRST, we need to add the RETH header size of 16 bytes to the above, which adds up to 78 bytes. Then the maximum message size for 1500 MTU without VLAN becomes: 1500 - WRITE FIRST common header + 65534 * (1500 - WRITE LAST/MIDDLE header size) = 1500 - 78 + 65534 * (1500 - 62) = 94239314 bytes or 89 MB.
// MessageSize returns a uint32
func (obj *roCEv2FlowSettingsRemoteEnd) MessageSize() uint32 {

	return *obj.obj.MessageSize

}

// The Maximum message size that is allowed to transfer depends on the MTU size and the number of VLANs configured on the interfaces. For any MTU, the maximum number of segmented RDMA WRITE packets for a single WRITE MESSAGE is 65535. That is, a single RDMA WRITE message can be broken into 1 WRITE FIRST, 1 WRITE LAST and (65535-2) WRITE MIDDLE messages. The maximum message size that is allowed to be transferred for a given MTU is constrained by the above conditions. For example, for an MTU size of 1500 bytes, the common header of the RDMA WRITE MIDDLE/LAST will comprise of Ethernet Header + IP Header + UDP Header + BTH Header + iCRC size + Ethernet Trailer size. This works out to be 14+20+8+12+4+4 = 62 bytes. For RDMA WRITE FIRST, we need to add the RETH header size of 16 bytes to the above, which adds up to 78 bytes. Then the maximum message size for 1500 MTU without VLAN becomes: 1500 - WRITE FIRST common header + 65534 * (1500 - WRITE LAST/MIDDLE header size) = 1500 - 78 + 65534 * (1500 - 62) = 94239314 bytes or 89 MB.
// MessageSize returns a uint32
func (obj *roCEv2FlowSettingsRemoteEnd) HasMessageSize() bool {
	return obj.obj.MessageSize != nil
}

// The Maximum message size that is allowed to transfer depends on the MTU size and the number of VLANs configured on the interfaces. For any MTU, the maximum number of segmented RDMA WRITE packets for a single WRITE MESSAGE is 65535. That is, a single RDMA WRITE message can be broken into 1 WRITE FIRST, 1 WRITE LAST and (65535-2) WRITE MIDDLE messages. The maximum message size that is allowed to be transferred for a given MTU is constrained by the above conditions. For example, for an MTU size of 1500 bytes, the common header of the RDMA WRITE MIDDLE/LAST will comprise of Ethernet Header + IP Header + UDP Header + BTH Header + iCRC size + Ethernet Trailer size. This works out to be 14+20+8+12+4+4 = 62 bytes. For RDMA WRITE FIRST, we need to add the RETH header size of 16 bytes to the above, which adds up to 78 bytes. Then the maximum message size for 1500 MTU without VLAN becomes: 1500 - WRITE FIRST common header + 65534 * (1500 - WRITE LAST/MIDDLE header size) = 1500 - 78 + 65534 * (1500 - 62) = 94239314 bytes or 89 MB.
// SetMessageSize sets the uint32 value in the RoCEv2FlowSettingsRemoteEnd object
func (obj *roCEv2FlowSettingsRemoteEnd) SetMessageSize(value uint32) RoCEv2FlowSettingsRemoteEnd {

	obj.obj.MessageSize = &value
	return obj
}

type RoCEv2FlowSettingsRemoteEndMessageSizeUnitEnum string

// Enum of MessageSizeUnit on RoCEv2FlowSettingsRemoteEnd
var RoCEv2FlowSettingsRemoteEndMessageSizeUnit = struct {
	BYTE RoCEv2FlowSettingsRemoteEndMessageSizeUnitEnum
	KB   RoCEv2FlowSettingsRemoteEndMessageSizeUnitEnum
	MB   RoCEv2FlowSettingsRemoteEndMessageSizeUnitEnum
}{
	BYTE: RoCEv2FlowSettingsRemoteEndMessageSizeUnitEnum("Byte"),
	KB:   RoCEv2FlowSettingsRemoteEndMessageSizeUnitEnum("KB"),
	MB:   RoCEv2FlowSettingsRemoteEndMessageSizeUnitEnum("MB"),
}

func (obj *roCEv2FlowSettingsRemoteEnd) MessageSizeUnit() RoCEv2FlowSettingsRemoteEndMessageSizeUnitEnum {
	return RoCEv2FlowSettingsRemoteEndMessageSizeUnitEnum(obj.obj.MessageSizeUnit.Enum().String())
}

// Unit of the transfer message size. Available options are Bytes, KB, MB.
// MessageSizeUnit returns a string
func (obj *roCEv2FlowSettingsRemoteEnd) HasMessageSizeUnit() bool {
	return obj.obj.MessageSizeUnit != nil
}

func (obj *roCEv2FlowSettingsRemoteEnd) SetMessageSizeUnit(value RoCEv2FlowSettingsRemoteEndMessageSizeUnitEnum) RoCEv2FlowSettingsRemoteEnd {
	intValue, ok := otg.RoCEv2FlowSettingsRemoteEnd_MessageSizeUnit_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on RoCEv2FlowSettingsRemoteEndMessageSizeUnitEnum", string(value)))
		return obj
	}
	enumValue := otg.RoCEv2FlowSettingsRemoteEnd_MessageSizeUnit_Enum(intValue)
	obj.obj.MessageSizeUnit = &enumValue

	return obj
}

// Remote QP Number.
// RemoteQpNumber returns a uint32
func (obj *roCEv2FlowSettingsRemoteEnd) RemoteQpNumber() uint32 {

	return *obj.obj.RemoteQpNumber

}

// Remote QP Number.
// RemoteQpNumber returns a uint32
func (obj *roCEv2FlowSettingsRemoteEnd) HasRemoteQpNumber() bool {
	return obj.obj.RemoteQpNumber != nil
}

// Remote QP Number.
// SetRemoteQpNumber sets the uint32 value in the RoCEv2FlowSettingsRemoteEnd object
func (obj *roCEv2FlowSettingsRemoteEnd) SetRemoteQpNumber(value uint32) RoCEv2FlowSettingsRemoteEnd {

	obj.obj.RemoteQpNumber = &value
	return obj
}

// Remote virtual address.
// VirtualAddress returns a string
func (obj *roCEv2FlowSettingsRemoteEnd) VirtualAddress() string {

	return *obj.obj.VirtualAddress

}

// Remote virtual address.
// VirtualAddress returns a string
func (obj *roCEv2FlowSettingsRemoteEnd) HasVirtualAddress() bool {
	return obj.obj.VirtualAddress != nil
}

// Remote virtual address.
// SetVirtualAddress sets the string value in the RoCEv2FlowSettingsRemoteEnd object
func (obj *roCEv2FlowSettingsRemoteEnd) SetVirtualAddress(value string) RoCEv2FlowSettingsRemoteEnd {

	obj.obj.VirtualAddress = &value
	return obj
}

// Remote Key.
// RemoteKey returns a string
func (obj *roCEv2FlowSettingsRemoteEnd) RemoteKey() string {

	return *obj.obj.RemoteKey

}

// Remote Key.
// RemoteKey returns a string
func (obj *roCEv2FlowSettingsRemoteEnd) HasRemoteKey() bool {
	return obj.obj.RemoteKey != nil
}

// Remote Key.
// SetRemoteKey sets the string value in the RoCEv2FlowSettingsRemoteEnd object
func (obj *roCEv2FlowSettingsRemoteEnd) SetRemoteKey(value string) RoCEv2FlowSettingsRemoteEnd {

	obj.obj.RemoteKey = &value
	return obj
}

// Initial PSN.
// InitialPsn returns a uint32
func (obj *roCEv2FlowSettingsRemoteEnd) InitialPsn() uint32 {

	return *obj.obj.InitialPsn

}

// Initial PSN.
// InitialPsn returns a uint32
func (obj *roCEv2FlowSettingsRemoteEnd) HasInitialPsn() bool {
	return obj.obj.InitialPsn != nil
}

// Initial PSN.
// SetInitialPsn sets the uint32 value in the RoCEv2FlowSettingsRemoteEnd object
func (obj *roCEv2FlowSettingsRemoteEnd) SetInitialPsn(value uint32) RoCEv2FlowSettingsRemoteEnd {

	obj.obj.InitialPsn = &value
	return obj
}

func (obj *roCEv2FlowSettingsRemoteEnd) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.CustomQpNumber != nil {

		if *obj.obj.CustomQpNumber < 2 || *obj.obj.CustomQpNumber > 33554431 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("2 <= RoCEv2FlowSettingsRemoteEnd.CustomQpNumber <= 33554431 but Got %d", *obj.obj.CustomQpNumber))
		}

	}

	if obj.obj.Dscp != nil {

		if *obj.obj.Dscp > 63 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RoCEv2FlowSettingsRemoteEnd.Dscp <= 63 but Got %d", *obj.obj.Dscp))
		}

	}

	if obj.obj.Ecn != nil {

		if *obj.obj.Ecn > 3 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RoCEv2FlowSettingsRemoteEnd.Ecn <= 3 but Got %d", *obj.obj.Ecn))
		}

	}

	if obj.obj.UdpSourcePort != nil {

		if *obj.obj.UdpSourcePort > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RoCEv2FlowSettingsRemoteEnd.UdpSourcePort <= 65535 but Got %d", *obj.obj.UdpSourcePort))
		}

	}

	if obj.obj.ImmediateData != nil {

		err := obj.validateHex(obj.ImmediateData())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on RoCEv2FlowSettingsRemoteEnd.ImmediateData"))
		}

	}

	if obj.obj.MessageSize != nil {

		if *obj.obj.MessageSize > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RoCEv2FlowSettingsRemoteEnd.MessageSize <= 65535 but Got %d", *obj.obj.MessageSize))
		}

	}

	if obj.obj.RemoteQpNumber != nil {

		if *obj.obj.RemoteQpNumber > 33554431 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RoCEv2FlowSettingsRemoteEnd.RemoteQpNumber <= 33554431 but Got %d", *obj.obj.RemoteQpNumber))
		}

	}

	if obj.obj.VirtualAddress != nil {

		err := obj.validateHex(obj.VirtualAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on RoCEv2FlowSettingsRemoteEnd.VirtualAddress"))
		}

	}

	if obj.obj.RemoteKey != nil {

		err := obj.validateHex(obj.RemoteKey())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on RoCEv2FlowSettingsRemoteEnd.RemoteKey"))
		}

	}

	if obj.obj.InitialPsn != nil {

		if *obj.obj.InitialPsn > 33554431 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RoCEv2FlowSettingsRemoteEnd.InitialPsn <= 33554431 but Got %d", *obj.obj.InitialPsn))
		}

	}

}

func (obj *roCEv2FlowSettingsRemoteEnd) setDefault() {
	if obj.obj.CustomQp == nil {
		obj.SetCustomQp(false)
	}
	if obj.obj.CustomQpNumber == nil {
		obj.SetCustomQpNumber(1)
	}
	if obj.obj.Dscp == nil {
		obj.SetDscp(24)
	}
	if obj.obj.Ecn == nil {
		obj.SetEcn(1)
	}
	if obj.obj.UdpSourcePort == nil {
		obj.SetUdpSourcePort(49152)
	}
	if obj.obj.ImmediateData == nil {
		obj.SetImmediateData("00 00 00 00")
	}
	if obj.obj.MessageSize == nil {
		obj.SetMessageSize(1)
	}
	if obj.obj.RemoteQpNumber == nil {
		obj.SetRemoteQpNumber(0)
	}
	if obj.obj.VirtualAddress == nil {
		obj.SetVirtualAddress("00 00 00 00 00 00 00 00")
	}
	if obj.obj.RemoteKey == nil {
		obj.SetRemoteKey("00 00 00 00")
	}
	if obj.obj.InitialPsn == nil {
		obj.SetInitialPsn(0)
	}

}
