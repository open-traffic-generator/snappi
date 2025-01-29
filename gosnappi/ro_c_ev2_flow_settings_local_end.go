package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RoCEv2FlowSettingsLocalEnd *****
type roCEv2FlowSettingsLocalEnd struct {
	validation
	obj          *otg.RoCEv2FlowSettingsLocalEnd
	marshaller   marshalRoCEv2FlowSettingsLocalEnd
	unMarshaller unMarshalRoCEv2FlowSettingsLocalEnd
}

func NewRoCEv2FlowSettingsLocalEnd() RoCEv2FlowSettingsLocalEnd {
	obj := roCEv2FlowSettingsLocalEnd{obj: &otg.RoCEv2FlowSettingsLocalEnd{}}
	obj.setDefault()
	return &obj
}

func (obj *roCEv2FlowSettingsLocalEnd) msg() *otg.RoCEv2FlowSettingsLocalEnd {
	return obj.obj
}

func (obj *roCEv2FlowSettingsLocalEnd) setMsg(msg *otg.RoCEv2FlowSettingsLocalEnd) RoCEv2FlowSettingsLocalEnd {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalroCEv2FlowSettingsLocalEnd struct {
	obj *roCEv2FlowSettingsLocalEnd
}

type marshalRoCEv2FlowSettingsLocalEnd interface {
	// ToProto marshals RoCEv2FlowSettingsLocalEnd to protobuf object *otg.RoCEv2FlowSettingsLocalEnd
	ToProto() (*otg.RoCEv2FlowSettingsLocalEnd, error)
	// ToPbText marshals RoCEv2FlowSettingsLocalEnd to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RoCEv2FlowSettingsLocalEnd to YAML text
	ToYaml() (string, error)
	// ToJson marshals RoCEv2FlowSettingsLocalEnd to JSON text
	ToJson() (string, error)
}

type unMarshalroCEv2FlowSettingsLocalEnd struct {
	obj *roCEv2FlowSettingsLocalEnd
}

type unMarshalRoCEv2FlowSettingsLocalEnd interface {
	// FromProto unmarshals RoCEv2FlowSettingsLocalEnd from protobuf object *otg.RoCEv2FlowSettingsLocalEnd
	FromProto(msg *otg.RoCEv2FlowSettingsLocalEnd) (RoCEv2FlowSettingsLocalEnd, error)
	// FromPbText unmarshals RoCEv2FlowSettingsLocalEnd from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RoCEv2FlowSettingsLocalEnd from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RoCEv2FlowSettingsLocalEnd from JSON text
	FromJson(value string) error
}

func (obj *roCEv2FlowSettingsLocalEnd) Marshal() marshalRoCEv2FlowSettingsLocalEnd {
	if obj.marshaller == nil {
		obj.marshaller = &marshalroCEv2FlowSettingsLocalEnd{obj: obj}
	}
	return obj.marshaller
}

func (obj *roCEv2FlowSettingsLocalEnd) Unmarshal() unMarshalRoCEv2FlowSettingsLocalEnd {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalroCEv2FlowSettingsLocalEnd{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalroCEv2FlowSettingsLocalEnd) ToProto() (*otg.RoCEv2FlowSettingsLocalEnd, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalroCEv2FlowSettingsLocalEnd) FromProto(msg *otg.RoCEv2FlowSettingsLocalEnd) (RoCEv2FlowSettingsLocalEnd, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalroCEv2FlowSettingsLocalEnd) ToPbText() (string, error) {
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

func (m *unMarshalroCEv2FlowSettingsLocalEnd) FromPbText(value string) error {
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

func (m *marshalroCEv2FlowSettingsLocalEnd) ToYaml() (string, error) {
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

func (m *unMarshalroCEv2FlowSettingsLocalEnd) FromYaml(value string) error {
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

func (m *marshalroCEv2FlowSettingsLocalEnd) ToJson() (string, error) {
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

func (m *unMarshalroCEv2FlowSettingsLocalEnd) FromJson(value string) error {
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

func (obj *roCEv2FlowSettingsLocalEnd) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *roCEv2FlowSettingsLocalEnd) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *roCEv2FlowSettingsLocalEnd) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *roCEv2FlowSettingsLocalEnd) Clone() (RoCEv2FlowSettingsLocalEnd, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRoCEv2FlowSettingsLocalEnd()
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

// RoCEv2FlowSettingsLocalEnd is local End Flow Settings
type RoCEv2FlowSettingsLocalEnd interface {
	Validation
	// msg marshals RoCEv2FlowSettingsLocalEnd to protobuf object *otg.RoCEv2FlowSettingsLocalEnd
	// and doesn't set defaults
	msg() *otg.RoCEv2FlowSettingsLocalEnd
	// setMsg unmarshals RoCEv2FlowSettingsLocalEnd from protobuf object *otg.RoCEv2FlowSettingsLocalEnd
	// and doesn't set defaults
	setMsg(*otg.RoCEv2FlowSettingsLocalEnd) RoCEv2FlowSettingsLocalEnd
	// provides marshal interface
	Marshal() marshalRoCEv2FlowSettingsLocalEnd
	// provides unmarshal interface
	Unmarshal() unMarshalRoCEv2FlowSettingsLocalEnd
	// validate validates RoCEv2FlowSettingsLocalEnd
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RoCEv2FlowSettingsLocalEnd, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// CustomQp returns bool, set in RoCEv2FlowSettingsLocalEnd.
	CustomQp() bool
	// SetCustomQp assigns bool provided by user to RoCEv2FlowSettingsLocalEnd
	SetCustomQp(value bool) RoCEv2FlowSettingsLocalEnd
	// HasCustomQp checks if CustomQp has been set in RoCEv2FlowSettingsLocalEnd
	HasCustomQp() bool
	// CustomQpNumber returns uint32, set in RoCEv2FlowSettingsLocalEnd.
	CustomQpNumber() uint32
	// SetCustomQpNumber assigns uint32 provided by user to RoCEv2FlowSettingsLocalEnd
	SetCustomQpNumber(value uint32) RoCEv2FlowSettingsLocalEnd
	// HasCustomQpNumber checks if CustomQpNumber has been set in RoCEv2FlowSettingsLocalEnd
	HasCustomQpNumber() bool
	// Dscp returns uint32, set in RoCEv2FlowSettingsLocalEnd.
	Dscp() uint32
	// SetDscp assigns uint32 provided by user to RoCEv2FlowSettingsLocalEnd
	SetDscp(value uint32) RoCEv2FlowSettingsLocalEnd
	// HasDscp checks if Dscp has been set in RoCEv2FlowSettingsLocalEnd
	HasDscp() bool
	// Ecn returns uint32, set in RoCEv2FlowSettingsLocalEnd.
	Ecn() uint32
	// SetEcn assigns uint32 provided by user to RoCEv2FlowSettingsLocalEnd
	SetEcn(value uint32) RoCEv2FlowSettingsLocalEnd
	// HasEcn checks if Ecn has been set in RoCEv2FlowSettingsLocalEnd
	HasEcn() bool
	// UdpSourcePort returns uint32, set in RoCEv2FlowSettingsLocalEnd.
	UdpSourcePort() uint32
	// SetUdpSourcePort assigns uint32 provided by user to RoCEv2FlowSettingsLocalEnd
	SetUdpSourcePort(value uint32) RoCEv2FlowSettingsLocalEnd
	// HasUdpSourcePort checks if UdpSourcePort has been set in RoCEv2FlowSettingsLocalEnd
	HasUdpSourcePort() bool
	// ExecuteCommands returns RoCEv2FlowSettingsLocalEndExecuteCommandsEnum, set in RoCEv2FlowSettingsLocalEnd
	ExecuteCommands() RoCEv2FlowSettingsLocalEndExecuteCommandsEnum
	// SetExecuteCommands assigns RoCEv2FlowSettingsLocalEndExecuteCommandsEnum provided by user to RoCEv2FlowSettingsLocalEnd
	SetExecuteCommands(value RoCEv2FlowSettingsLocalEndExecuteCommandsEnum) RoCEv2FlowSettingsLocalEnd
	// HasExecuteCommands checks if ExecuteCommands has been set in RoCEv2FlowSettingsLocalEnd
	HasExecuteCommands() bool
	// ImmediateData returns string, set in RoCEv2FlowSettingsLocalEnd.
	ImmediateData() string
	// SetImmediateData assigns string provided by user to RoCEv2FlowSettingsLocalEnd
	SetImmediateData(value string) RoCEv2FlowSettingsLocalEnd
	// HasImmediateData checks if ImmediateData has been set in RoCEv2FlowSettingsLocalEnd
	HasImmediateData() bool
	// MessageSize returns uint32, set in RoCEv2FlowSettingsLocalEnd.
	MessageSize() uint32
	// SetMessageSize assigns uint32 provided by user to RoCEv2FlowSettingsLocalEnd
	SetMessageSize(value uint32) RoCEv2FlowSettingsLocalEnd
	// HasMessageSize checks if MessageSize has been set in RoCEv2FlowSettingsLocalEnd
	HasMessageSize() bool
	// MessageSizeUnit returns RoCEv2FlowSettingsLocalEndMessageSizeUnitEnum, set in RoCEv2FlowSettingsLocalEnd
	MessageSizeUnit() RoCEv2FlowSettingsLocalEndMessageSizeUnitEnum
	// SetMessageSizeUnit assigns RoCEv2FlowSettingsLocalEndMessageSizeUnitEnum provided by user to RoCEv2FlowSettingsLocalEnd
	SetMessageSizeUnit(value RoCEv2FlowSettingsLocalEndMessageSizeUnitEnum) RoCEv2FlowSettingsLocalEnd
	// HasMessageSizeUnit checks if MessageSizeUnit has been set in RoCEv2FlowSettingsLocalEnd
	HasMessageSizeUnit() bool
	// InitialPsn returns uint32, set in RoCEv2FlowSettingsLocalEnd.
	InitialPsn() uint32
	// SetInitialPsn assigns uint32 provided by user to RoCEv2FlowSettingsLocalEnd
	SetInitialPsn(value uint32) RoCEv2FlowSettingsLocalEnd
	// HasInitialPsn checks if InitialPsn has been set in RoCEv2FlowSettingsLocalEnd
	HasInitialPsn() bool
}

// Turn on to define QP number.
// CustomQp returns a bool
func (obj *roCEv2FlowSettingsLocalEnd) CustomQp() bool {

	return *obj.obj.CustomQp

}

// Turn on to define QP number.
// CustomQp returns a bool
func (obj *roCEv2FlowSettingsLocalEnd) HasCustomQp() bool {
	return obj.obj.CustomQp != nil
}

// Turn on to define QP number.
// SetCustomQp sets the bool value in the RoCEv2FlowSettingsLocalEnd object
func (obj *roCEv2FlowSettingsLocalEnd) SetCustomQp(value bool) RoCEv2FlowSettingsLocalEnd {

	obj.obj.CustomQp = &value
	return obj
}

// Configure the QP range.
// CustomQpNumber returns a uint32
func (obj *roCEv2FlowSettingsLocalEnd) CustomQpNumber() uint32 {

	return *obj.obj.CustomQpNumber

}

// Configure the QP range.
// CustomQpNumber returns a uint32
func (obj *roCEv2FlowSettingsLocalEnd) HasCustomQpNumber() bool {
	return obj.obj.CustomQpNumber != nil
}

// Configure the QP range.
// SetCustomQpNumber sets the uint32 value in the RoCEv2FlowSettingsLocalEnd object
func (obj *roCEv2FlowSettingsLocalEnd) SetCustomQpNumber(value uint32) RoCEv2FlowSettingsLocalEnd {

	obj.obj.CustomQpNumber = &value
	return obj
}

// DSCP value for this flow
// Dscp returns a uint32
func (obj *roCEv2FlowSettingsLocalEnd) Dscp() uint32 {

	return *obj.obj.Dscp

}

// DSCP value for this flow
// Dscp returns a uint32
func (obj *roCEv2FlowSettingsLocalEnd) HasDscp() bool {
	return obj.obj.Dscp != nil
}

// DSCP value for this flow
// SetDscp sets the uint32 value in the RoCEv2FlowSettingsLocalEnd object
func (obj *roCEv2FlowSettingsLocalEnd) SetDscp(value uint32) RoCEv2FlowSettingsLocalEnd {

	obj.obj.Dscp = &value
	return obj
}

// This field allows to configure bits of the Traffic Class field in the IPv4 or IPv6 header to encode four different code points.
// Ecn returns a uint32
func (obj *roCEv2FlowSettingsLocalEnd) Ecn() uint32 {

	return *obj.obj.Ecn

}

// This field allows to configure bits of the Traffic Class field in the IPv4 or IPv6 header to encode four different code points.
// Ecn returns a uint32
func (obj *roCEv2FlowSettingsLocalEnd) HasEcn() bool {
	return obj.obj.Ecn != nil
}

// This field allows to configure bits of the Traffic Class field in the IPv4 or IPv6 header to encode four different code points.
// SetEcn sets the uint32 value in the RoCEv2FlowSettingsLocalEnd object
func (obj *roCEv2FlowSettingsLocalEnd) SetEcn(value uint32) RoCEv2FlowSettingsLocalEnd {

	obj.obj.Ecn = &value
	return obj
}

// UDP source port number for this flow.
// UdpSourcePort returns a uint32
func (obj *roCEv2FlowSettingsLocalEnd) UdpSourcePort() uint32 {

	return *obj.obj.UdpSourcePort

}

// UDP source port number for this flow.
// UdpSourcePort returns a uint32
func (obj *roCEv2FlowSettingsLocalEnd) HasUdpSourcePort() bool {
	return obj.obj.UdpSourcePort != nil
}

// UDP source port number for this flow.
// SetUdpSourcePort sets the uint32 value in the RoCEv2FlowSettingsLocalEnd object
func (obj *roCEv2FlowSettingsLocalEnd) SetUdpSourcePort(value uint32) RoCEv2FlowSettingsLocalEnd {

	obj.obj.UdpSourcePort = &value
	return obj
}

type RoCEv2FlowSettingsLocalEndExecuteCommandsEnum string

// Enum of ExecuteCommands on RoCEv2FlowSettingsLocalEnd
var RoCEv2FlowSettingsLocalEndExecuteCommands = struct {
	NONE      RoCEv2FlowSettingsLocalEndExecuteCommandsEnum
	RDMAWRITE RoCEv2FlowSettingsLocalEndExecuteCommandsEnum
}{
	NONE:      RoCEv2FlowSettingsLocalEndExecuteCommandsEnum("none"),
	RDMAWRITE: RoCEv2FlowSettingsLocalEndExecuteCommandsEnum("RDMAWrite"),
}

func (obj *roCEv2FlowSettingsLocalEnd) ExecuteCommands() RoCEv2FlowSettingsLocalEndExecuteCommandsEnum {
	return RoCEv2FlowSettingsLocalEndExecuteCommandsEnum(obj.obj.ExecuteCommands.Enum().String())
}

// UDP source port number for this flow.
// ExecuteCommands returns a string
func (obj *roCEv2FlowSettingsLocalEnd) HasExecuteCommands() bool {
	return obj.obj.ExecuteCommands != nil
}

func (obj *roCEv2FlowSettingsLocalEnd) SetExecuteCommands(value RoCEv2FlowSettingsLocalEndExecuteCommandsEnum) RoCEv2FlowSettingsLocalEnd {
	intValue, ok := otg.RoCEv2FlowSettingsLocalEnd_ExecuteCommands_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on RoCEv2FlowSettingsLocalEndExecuteCommandsEnum", string(value)))
		return obj
	}
	enumValue := otg.RoCEv2FlowSettingsLocalEnd_ExecuteCommands_Enum(intValue)
	obj.obj.ExecuteCommands = &enumValue

	return obj
}

// Immediate Data field required for SEND/WRITE with immediate verb.
// ImmediateData returns a string
func (obj *roCEv2FlowSettingsLocalEnd) ImmediateData() string {

	return *obj.obj.ImmediateData

}

// Immediate Data field required for SEND/WRITE with immediate verb.
// ImmediateData returns a string
func (obj *roCEv2FlowSettingsLocalEnd) HasImmediateData() bool {
	return obj.obj.ImmediateData != nil
}

// Immediate Data field required for SEND/WRITE with immediate verb.
// SetImmediateData sets the string value in the RoCEv2FlowSettingsLocalEnd object
func (obj *roCEv2FlowSettingsLocalEnd) SetImmediateData(value string) RoCEv2FlowSettingsLocalEnd {

	obj.obj.ImmediateData = &value
	return obj
}

// The Maximum message size that is allowed to transfer depends on the MTU size and the number of VLANs configured on the interfaces. For any MTU, the maximum number of segmented RDMA WRITE packets for a single WRITE MESSAGE is 65535. That is, a single RDMA WRITE message can be broken into 1 WRITE FIRST, 1 WRITE LAST and (65535-2) WRITE MIDDLE messages. The maximum message size that is allowed to be transferred for a given MTU is constrained by the above conditions. For example, for an MTU size of 1500 bytes, the common header of the RDMA WRITE MIDDLE/LAST will comprise of Ethernet Header + IP Header + UDP Header + BTH Header + iCRC size + Ethernet Trailer size. This works out to be 14+20+8+12+4+4 = 62 bytes. For RDMA WRITE FIRST, we need to add the RETH header size of 16 bytes to the above, which adds up to 78 bytes. Then the maximum message size for 1500 MTU without VLAN becomes: 1500 - WRITE FIRST common header + 65534 * (1500 - WRITE LAST/MIDDLE header size) = 1500 - 78 + 65534 * (1500 - 62) = 94239314 bytes or 89 MB.
// MessageSize returns a uint32
func (obj *roCEv2FlowSettingsLocalEnd) MessageSize() uint32 {

	return *obj.obj.MessageSize

}

// The Maximum message size that is allowed to transfer depends on the MTU size and the number of VLANs configured on the interfaces. For any MTU, the maximum number of segmented RDMA WRITE packets for a single WRITE MESSAGE is 65535. That is, a single RDMA WRITE message can be broken into 1 WRITE FIRST, 1 WRITE LAST and (65535-2) WRITE MIDDLE messages. The maximum message size that is allowed to be transferred for a given MTU is constrained by the above conditions. For example, for an MTU size of 1500 bytes, the common header of the RDMA WRITE MIDDLE/LAST will comprise of Ethernet Header + IP Header + UDP Header + BTH Header + iCRC size + Ethernet Trailer size. This works out to be 14+20+8+12+4+4 = 62 bytes. For RDMA WRITE FIRST, we need to add the RETH header size of 16 bytes to the above, which adds up to 78 bytes. Then the maximum message size for 1500 MTU without VLAN becomes: 1500 - WRITE FIRST common header + 65534 * (1500 - WRITE LAST/MIDDLE header size) = 1500 - 78 + 65534 * (1500 - 62) = 94239314 bytes or 89 MB.
// MessageSize returns a uint32
func (obj *roCEv2FlowSettingsLocalEnd) HasMessageSize() bool {
	return obj.obj.MessageSize != nil
}

// The Maximum message size that is allowed to transfer depends on the MTU size and the number of VLANs configured on the interfaces. For any MTU, the maximum number of segmented RDMA WRITE packets for a single WRITE MESSAGE is 65535. That is, a single RDMA WRITE message can be broken into 1 WRITE FIRST, 1 WRITE LAST and (65535-2) WRITE MIDDLE messages. The maximum message size that is allowed to be transferred for a given MTU is constrained by the above conditions. For example, for an MTU size of 1500 bytes, the common header of the RDMA WRITE MIDDLE/LAST will comprise of Ethernet Header + IP Header + UDP Header + BTH Header + iCRC size + Ethernet Trailer size. This works out to be 14+20+8+12+4+4 = 62 bytes. For RDMA WRITE FIRST, we need to add the RETH header size of 16 bytes to the above, which adds up to 78 bytes. Then the maximum message size for 1500 MTU without VLAN becomes: 1500 - WRITE FIRST common header + 65534 * (1500 - WRITE LAST/MIDDLE header size) = 1500 - 78 + 65534 * (1500 - 62) = 94239314 bytes or 89 MB.
// SetMessageSize sets the uint32 value in the RoCEv2FlowSettingsLocalEnd object
func (obj *roCEv2FlowSettingsLocalEnd) SetMessageSize(value uint32) RoCEv2FlowSettingsLocalEnd {

	obj.obj.MessageSize = &value
	return obj
}

type RoCEv2FlowSettingsLocalEndMessageSizeUnitEnum string

// Enum of MessageSizeUnit on RoCEv2FlowSettingsLocalEnd
var RoCEv2FlowSettingsLocalEndMessageSizeUnit = struct {
	BYTE RoCEv2FlowSettingsLocalEndMessageSizeUnitEnum
	KB   RoCEv2FlowSettingsLocalEndMessageSizeUnitEnum
	MB   RoCEv2FlowSettingsLocalEndMessageSizeUnitEnum
}{
	BYTE: RoCEv2FlowSettingsLocalEndMessageSizeUnitEnum("Byte"),
	KB:   RoCEv2FlowSettingsLocalEndMessageSizeUnitEnum("KB"),
	MB:   RoCEv2FlowSettingsLocalEndMessageSizeUnitEnum("MB"),
}

func (obj *roCEv2FlowSettingsLocalEnd) MessageSizeUnit() RoCEv2FlowSettingsLocalEndMessageSizeUnitEnum {
	return RoCEv2FlowSettingsLocalEndMessageSizeUnitEnum(obj.obj.MessageSizeUnit.Enum().String())
}

// Unit of the transfer message size. Available options are Bytes, KB, MB.
// MessageSizeUnit returns a string
func (obj *roCEv2FlowSettingsLocalEnd) HasMessageSizeUnit() bool {
	return obj.obj.MessageSizeUnit != nil
}

func (obj *roCEv2FlowSettingsLocalEnd) SetMessageSizeUnit(value RoCEv2FlowSettingsLocalEndMessageSizeUnitEnum) RoCEv2FlowSettingsLocalEnd {
	intValue, ok := otg.RoCEv2FlowSettingsLocalEnd_MessageSizeUnit_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on RoCEv2FlowSettingsLocalEndMessageSizeUnitEnum", string(value)))
		return obj
	}
	enumValue := otg.RoCEv2FlowSettingsLocalEnd_MessageSizeUnit_Enum(intValue)
	obj.obj.MessageSizeUnit = &enumValue

	return obj
}

// Initial PSN
// InitialPsn returns a uint32
func (obj *roCEv2FlowSettingsLocalEnd) InitialPsn() uint32 {

	return *obj.obj.InitialPsn

}

// Initial PSN
// InitialPsn returns a uint32
func (obj *roCEv2FlowSettingsLocalEnd) HasInitialPsn() bool {
	return obj.obj.InitialPsn != nil
}

// Initial PSN
// SetInitialPsn sets the uint32 value in the RoCEv2FlowSettingsLocalEnd object
func (obj *roCEv2FlowSettingsLocalEnd) SetInitialPsn(value uint32) RoCEv2FlowSettingsLocalEnd {

	obj.obj.InitialPsn = &value
	return obj
}

func (obj *roCEv2FlowSettingsLocalEnd) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.CustomQpNumber != nil {

		if *obj.obj.CustomQpNumber < 2 || *obj.obj.CustomQpNumber > 33554431 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("2 <= RoCEv2FlowSettingsLocalEnd.CustomQpNumber <= 33554431 but Got %d", *obj.obj.CustomQpNumber))
		}

	}

	if obj.obj.Dscp != nil {

		if *obj.obj.Dscp > 63 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RoCEv2FlowSettingsLocalEnd.Dscp <= 63 but Got %d", *obj.obj.Dscp))
		}

	}

	if obj.obj.Ecn != nil {

		if *obj.obj.Ecn > 3 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RoCEv2FlowSettingsLocalEnd.Ecn <= 3 but Got %d", *obj.obj.Ecn))
		}

	}

	if obj.obj.UdpSourcePort != nil {

		if *obj.obj.UdpSourcePort > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RoCEv2FlowSettingsLocalEnd.UdpSourcePort <= 65535 but Got %d", *obj.obj.UdpSourcePort))
		}

	}

	if obj.obj.ImmediateData != nil {

		err := obj.validateHex(obj.ImmediateData())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on RoCEv2FlowSettingsLocalEnd.ImmediateData"))
		}

	}

	if obj.obj.MessageSize != nil {

		if *obj.obj.MessageSize > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RoCEv2FlowSettingsLocalEnd.MessageSize <= 65535 but Got %d", *obj.obj.MessageSize))
		}

	}

	if obj.obj.InitialPsn != nil {

		if *obj.obj.InitialPsn > 33554431 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RoCEv2FlowSettingsLocalEnd.InitialPsn <= 33554431 but Got %d", *obj.obj.InitialPsn))
		}

	}

}

func (obj *roCEv2FlowSettingsLocalEnd) setDefault() {
	if obj.obj.CustomQp == nil {
		obj.SetCustomQp(false)
	}
	if obj.obj.CustomQpNumber == nil {
		obj.SetCustomQpNumber(2)
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
	if obj.obj.InitialPsn == nil {
		obj.SetInitialPsn(24)
	}

}
