package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RoCEv2FlowSettings *****
type roCEv2FlowSettings struct {
	validation
	obj          *otg.RoCEv2FlowSettings
	marshaller   marshalRoCEv2FlowSettings
	unMarshaller unMarshalRoCEv2FlowSettings
}

func NewRoCEv2FlowSettings() RoCEv2FlowSettings {
	obj := roCEv2FlowSettings{obj: &otg.RoCEv2FlowSettings{}}
	obj.setDefault()
	return &obj
}

func (obj *roCEv2FlowSettings) msg() *otg.RoCEv2FlowSettings {
	return obj.obj
}

func (obj *roCEv2FlowSettings) setMsg(msg *otg.RoCEv2FlowSettings) RoCEv2FlowSettings {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalroCEv2FlowSettings struct {
	obj *roCEv2FlowSettings
}

type marshalRoCEv2FlowSettings interface {
	// ToProto marshals RoCEv2FlowSettings to protobuf object *otg.RoCEv2FlowSettings
	ToProto() (*otg.RoCEv2FlowSettings, error)
	// ToPbText marshals RoCEv2FlowSettings to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RoCEv2FlowSettings to YAML text
	ToYaml() (string, error)
	// ToJson marshals RoCEv2FlowSettings to JSON text
	ToJson() (string, error)
}

type unMarshalroCEv2FlowSettings struct {
	obj *roCEv2FlowSettings
}

type unMarshalRoCEv2FlowSettings interface {
	// FromProto unmarshals RoCEv2FlowSettings from protobuf object *otg.RoCEv2FlowSettings
	FromProto(msg *otg.RoCEv2FlowSettings) (RoCEv2FlowSettings, error)
	// FromPbText unmarshals RoCEv2FlowSettings from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RoCEv2FlowSettings from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RoCEv2FlowSettings from JSON text
	FromJson(value string) error
}

func (obj *roCEv2FlowSettings) Marshal() marshalRoCEv2FlowSettings {
	if obj.marshaller == nil {
		obj.marshaller = &marshalroCEv2FlowSettings{obj: obj}
	}
	return obj.marshaller
}

func (obj *roCEv2FlowSettings) Unmarshal() unMarshalRoCEv2FlowSettings {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalroCEv2FlowSettings{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalroCEv2FlowSettings) ToProto() (*otg.RoCEv2FlowSettings, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalroCEv2FlowSettings) FromProto(msg *otg.RoCEv2FlowSettings) (RoCEv2FlowSettings, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalroCEv2FlowSettings) ToPbText() (string, error) {
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

func (m *unMarshalroCEv2FlowSettings) FromPbText(value string) error {
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

func (m *marshalroCEv2FlowSettings) ToYaml() (string, error) {
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

func (m *unMarshalroCEv2FlowSettings) FromYaml(value string) error {
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

func (m *marshalroCEv2FlowSettings) ToJson() (string, error) {
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

func (m *unMarshalroCEv2FlowSettings) FromJson(value string) error {
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

func (obj *roCEv2FlowSettings) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *roCEv2FlowSettings) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *roCEv2FlowSettings) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *roCEv2FlowSettings) Clone() (RoCEv2FlowSettings, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRoCEv2FlowSettings()
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

// RoCEv2FlowSettings is this section has two views, Local End and Remote End.
// Both views have same configurations. However, the remote and local peer IP addresses are interchanged.
// This configuration allows you to configure RDMA flow over the same QP number from same source and destination.
// Default value for commands at Remote peer is set to None.
// So that by default, this remote peer does not initiate any traffic flow.
type RoCEv2FlowSettings interface {
	Validation
	// msg marshals RoCEv2FlowSettings to protobuf object *otg.RoCEv2FlowSettings
	// and doesn't set defaults
	msg() *otg.RoCEv2FlowSettings
	// setMsg unmarshals RoCEv2FlowSettings from protobuf object *otg.RoCEv2FlowSettings
	// and doesn't set defaults
	setMsg(*otg.RoCEv2FlowSettings) RoCEv2FlowSettings
	// provides marshal interface
	Marshal() marshalRoCEv2FlowSettings
	// provides unmarshal interface
	Unmarshal() unMarshalRoCEv2FlowSettings
	// validate validates RoCEv2FlowSettings
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RoCEv2FlowSettings, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// CustomQp returns bool, set in RoCEv2FlowSettings.
	CustomQp() bool
	// SetCustomQp assigns bool provided by user to RoCEv2FlowSettings
	SetCustomQp(value bool) RoCEv2FlowSettings
	// HasCustomQp checks if CustomQp has been set in RoCEv2FlowSettings
	HasCustomQp() bool
	// CustomQpNumber returns uint32, set in RoCEv2FlowSettings.
	CustomQpNumber() uint32
	// SetCustomQpNumber assigns uint32 provided by user to RoCEv2FlowSettings
	SetCustomQpNumber(value uint32) RoCEv2FlowSettings
	// HasCustomQpNumber checks if CustomQpNumber has been set in RoCEv2FlowSettings
	HasCustomQpNumber() bool
	// Dscp returns uint32, set in RoCEv2FlowSettings.
	Dscp() uint32
	// SetDscp assigns uint32 provided by user to RoCEv2FlowSettings
	SetDscp(value uint32) RoCEv2FlowSettings
	// HasDscp checks if Dscp has been set in RoCEv2FlowSettings
	HasDscp() bool
	// Ecn returns uint32, set in RoCEv2FlowSettings.
	Ecn() uint32
	// SetEcn assigns uint32 provided by user to RoCEv2FlowSettings
	SetEcn(value uint32) RoCEv2FlowSettings
	// HasEcn checks if Ecn has been set in RoCEv2FlowSettings
	HasEcn() bool
	// UdpSourcePort returns uint32, set in RoCEv2FlowSettings.
	UdpSourcePort() uint32
	// SetUdpSourcePort assigns uint32 provided by user to RoCEv2FlowSettings
	SetUdpSourcePort(value uint32) RoCEv2FlowSettings
	// HasUdpSourcePort checks if UdpSourcePort has been set in RoCEv2FlowSettings
	HasUdpSourcePort() bool
	// Rocev2Verb returns RoCEv2FlowSettingsRocev2VerbEnum, set in RoCEv2FlowSettings
	Rocev2Verb() RoCEv2FlowSettingsRocev2VerbEnum
	// SetRocev2Verb assigns RoCEv2FlowSettingsRocev2VerbEnum provided by user to RoCEv2FlowSettings
	SetRocev2Verb(value RoCEv2FlowSettingsRocev2VerbEnum) RoCEv2FlowSettings
	// HasRocev2Verb checks if Rocev2Verb has been set in RoCEv2FlowSettings
	HasRocev2Verb() bool
	// ImmediateData returns string, set in RoCEv2FlowSettings.
	ImmediateData() string
	// SetImmediateData assigns string provided by user to RoCEv2FlowSettings
	SetImmediateData(value string) RoCEv2FlowSettings
	// HasImmediateData checks if ImmediateData has been set in RoCEv2FlowSettings
	HasImmediateData() bool
	// MessageSize returns uint32, set in RoCEv2FlowSettings.
	MessageSize() uint32
	// SetMessageSize assigns uint32 provided by user to RoCEv2FlowSettings
	SetMessageSize(value uint32) RoCEv2FlowSettings
	// HasMessageSize checks if MessageSize has been set in RoCEv2FlowSettings
	HasMessageSize() bool
	// MessageSizeUnit returns RoCEv2FlowSettingsMessageSizeUnitEnum, set in RoCEv2FlowSettings
	MessageSizeUnit() RoCEv2FlowSettingsMessageSizeUnitEnum
	// SetMessageSizeUnit assigns RoCEv2FlowSettingsMessageSizeUnitEnum provided by user to RoCEv2FlowSettings
	SetMessageSizeUnit(value RoCEv2FlowSettingsMessageSizeUnitEnum) RoCEv2FlowSettings
	// HasMessageSizeUnit checks if MessageSizeUnit has been set in RoCEv2FlowSettings
	HasMessageSizeUnit() bool
}

// Turn on to define QP number.
// CustomQp returns a bool
func (obj *roCEv2FlowSettings) CustomQp() bool {

	return *obj.obj.CustomQp

}

// Turn on to define QP number.
// CustomQp returns a bool
func (obj *roCEv2FlowSettings) HasCustomQp() bool {
	return obj.obj.CustomQp != nil
}

// Turn on to define QP number.
// SetCustomQp sets the bool value in the RoCEv2FlowSettings object
func (obj *roCEv2FlowSettings) SetCustomQp(value bool) RoCEv2FlowSettings {

	obj.obj.CustomQp = &value
	return obj
}

// Configure the QP range.
// CustomQpNumber returns a uint32
func (obj *roCEv2FlowSettings) CustomQpNumber() uint32 {

	return *obj.obj.CustomQpNumber

}

// Configure the QP range.
// CustomQpNumber returns a uint32
func (obj *roCEv2FlowSettings) HasCustomQpNumber() bool {
	return obj.obj.CustomQpNumber != nil
}

// Configure the QP range.
// SetCustomQpNumber sets the uint32 value in the RoCEv2FlowSettings object
func (obj *roCEv2FlowSettings) SetCustomQpNumber(value uint32) RoCEv2FlowSettings {

	obj.obj.CustomQpNumber = &value
	return obj
}

// DSCP value for this flow
// Dscp returns a uint32
func (obj *roCEv2FlowSettings) Dscp() uint32 {

	return *obj.obj.Dscp

}

// DSCP value for this flow
// Dscp returns a uint32
func (obj *roCEv2FlowSettings) HasDscp() bool {
	return obj.obj.Dscp != nil
}

// DSCP value for this flow
// SetDscp sets the uint32 value in the RoCEv2FlowSettings object
func (obj *roCEv2FlowSettings) SetDscp(value uint32) RoCEv2FlowSettings {

	obj.obj.Dscp = &value
	return obj
}

// This field allows to configure bits of the Traffic Class field in the IPv4 or IPv6 header to encode four different code points.
// Ecn returns a uint32
func (obj *roCEv2FlowSettings) Ecn() uint32 {

	return *obj.obj.Ecn

}

// This field allows to configure bits of the Traffic Class field in the IPv4 or IPv6 header to encode four different code points.
// Ecn returns a uint32
func (obj *roCEv2FlowSettings) HasEcn() bool {
	return obj.obj.Ecn != nil
}

// This field allows to configure bits of the Traffic Class field in the IPv4 or IPv6 header to encode four different code points.
// SetEcn sets the uint32 value in the RoCEv2FlowSettings object
func (obj *roCEv2FlowSettings) SetEcn(value uint32) RoCEv2FlowSettings {

	obj.obj.Ecn = &value
	return obj
}

// UDP source port number for this flow.
// UdpSourcePort returns a uint32
func (obj *roCEv2FlowSettings) UdpSourcePort() uint32 {

	return *obj.obj.UdpSourcePort

}

// UDP source port number for this flow.
// UdpSourcePort returns a uint32
func (obj *roCEv2FlowSettings) HasUdpSourcePort() bool {
	return obj.obj.UdpSourcePort != nil
}

// UDP source port number for this flow.
// SetUdpSourcePort sets the uint32 value in the RoCEv2FlowSettings object
func (obj *roCEv2FlowSettings) SetUdpSourcePort(value uint32) RoCEv2FlowSettings {

	obj.obj.UdpSourcePort = &value
	return obj
}

type RoCEv2FlowSettingsRocev2VerbEnum string

// Enum of Rocev2Verb on RoCEv2FlowSettings
var RoCEv2FlowSettingsRocev2Verb = struct {
	NONE                 RoCEv2FlowSettingsRocev2VerbEnum
	WRITE                RoCEv2FlowSettingsRocev2VerbEnum
	WRITE_WITH_IMMEDIATE RoCEv2FlowSettingsRocev2VerbEnum
	SEND                 RoCEv2FlowSettingsRocev2VerbEnum
	SEND_WITH_IMMEDIATE  RoCEv2FlowSettingsRocev2VerbEnum
}{
	NONE:                 RoCEv2FlowSettingsRocev2VerbEnum("none"),
	WRITE:                RoCEv2FlowSettingsRocev2VerbEnum("write"),
	WRITE_WITH_IMMEDIATE: RoCEv2FlowSettingsRocev2VerbEnum("write_with_immediate"),
	SEND:                 RoCEv2FlowSettingsRocev2VerbEnum("send"),
	SEND_WITH_IMMEDIATE:  RoCEv2FlowSettingsRocev2VerbEnum("send_with_immediate"),
}

func (obj *roCEv2FlowSettings) Rocev2Verb() RoCEv2FlowSettingsRocev2VerbEnum {
	return RoCEv2FlowSettingsRocev2VerbEnum(obj.obj.Rocev2Verb.Enum().String())
}

// RoCEv2 Verb, Available options are: RDMA WRITE None: The corresponding flow will not take part in traffic..
// Rocev2Verb returns a string
func (obj *roCEv2FlowSettings) HasRocev2Verb() bool {
	return obj.obj.Rocev2Verb != nil
}

func (obj *roCEv2FlowSettings) SetRocev2Verb(value RoCEv2FlowSettingsRocev2VerbEnum) RoCEv2FlowSettings {
	intValue, ok := otg.RoCEv2FlowSettings_Rocev2Verb_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on RoCEv2FlowSettingsRocev2VerbEnum", string(value)))
		return obj
	}
	enumValue := otg.RoCEv2FlowSettings_Rocev2Verb_Enum(intValue)
	obj.obj.Rocev2Verb = &enumValue

	return obj
}

// Immediate Data field required for SEND/WRITE with immediate verb.
// ImmediateData returns a string
func (obj *roCEv2FlowSettings) ImmediateData() string {

	return *obj.obj.ImmediateData

}

// Immediate Data field required for SEND/WRITE with immediate verb.
// ImmediateData returns a string
func (obj *roCEv2FlowSettings) HasImmediateData() bool {
	return obj.obj.ImmediateData != nil
}

// Immediate Data field required for SEND/WRITE with immediate verb.
// SetImmediateData sets the string value in the RoCEv2FlowSettings object
func (obj *roCEv2FlowSettings) SetImmediateData(value string) RoCEv2FlowSettings {

	obj.obj.ImmediateData = &value
	return obj
}

// The Maximum message size that is allowed to transfer depends on the MTU size and the number of VLANs configured on the interfaces. For any MTU, the maximum number of segmented RDMA WRITE packets for a single WRITE MESSAGE is 65535. That is, a single RDMA WRITE message can be broken into 1 WRITE FIRST, 1 WRITE LAST and (65535-2) WRITE MIDDLE messages. The maximum message size that is allowed to be transferred for a given MTU is constrained by the above conditions. For example, for an MTU size of 1500 bytes, the common header of the RDMA WRITE MIDDLE/LAST will comprise of Ethernet Header + IP Header + UDP Header + BTH Header + iCRC size + Ethernet Trailer size. This works out to be 14+20+8+12+4+4 = 62 bytes. For RDMA WRITE FIRST, we need to add the RETH header size of 16 bytes to the above, which adds up to 78 bytes. Then the maximum message size for 1500 MTU without VLAN becomes: 1500 - WRITE FIRST common header + 65534 * (1500 - WRITE LAST/MIDDLE header size) = 1500 - 78 + 65534 * (1500 - 62) = 94239314 bytes or 89 MB.
// MessageSize returns a uint32
func (obj *roCEv2FlowSettings) MessageSize() uint32 {

	return *obj.obj.MessageSize

}

// The Maximum message size that is allowed to transfer depends on the MTU size and the number of VLANs configured on the interfaces. For any MTU, the maximum number of segmented RDMA WRITE packets for a single WRITE MESSAGE is 65535. That is, a single RDMA WRITE message can be broken into 1 WRITE FIRST, 1 WRITE LAST and (65535-2) WRITE MIDDLE messages. The maximum message size that is allowed to be transferred for a given MTU is constrained by the above conditions. For example, for an MTU size of 1500 bytes, the common header of the RDMA WRITE MIDDLE/LAST will comprise of Ethernet Header + IP Header + UDP Header + BTH Header + iCRC size + Ethernet Trailer size. This works out to be 14+20+8+12+4+4 = 62 bytes. For RDMA WRITE FIRST, we need to add the RETH header size of 16 bytes to the above, which adds up to 78 bytes. Then the maximum message size for 1500 MTU without VLAN becomes: 1500 - WRITE FIRST common header + 65534 * (1500 - WRITE LAST/MIDDLE header size) = 1500 - 78 + 65534 * (1500 - 62) = 94239314 bytes or 89 MB.
// MessageSize returns a uint32
func (obj *roCEv2FlowSettings) HasMessageSize() bool {
	return obj.obj.MessageSize != nil
}

// The Maximum message size that is allowed to transfer depends on the MTU size and the number of VLANs configured on the interfaces. For any MTU, the maximum number of segmented RDMA WRITE packets for a single WRITE MESSAGE is 65535. That is, a single RDMA WRITE message can be broken into 1 WRITE FIRST, 1 WRITE LAST and (65535-2) WRITE MIDDLE messages. The maximum message size that is allowed to be transferred for a given MTU is constrained by the above conditions. For example, for an MTU size of 1500 bytes, the common header of the RDMA WRITE MIDDLE/LAST will comprise of Ethernet Header + IP Header + UDP Header + BTH Header + iCRC size + Ethernet Trailer size. This works out to be 14+20+8+12+4+4 = 62 bytes. For RDMA WRITE FIRST, we need to add the RETH header size of 16 bytes to the above, which adds up to 78 bytes. Then the maximum message size for 1500 MTU without VLAN becomes: 1500 - WRITE FIRST common header + 65534 * (1500 - WRITE LAST/MIDDLE header size) = 1500 - 78 + 65534 * (1500 - 62) = 94239314 bytes or 89 MB.
// SetMessageSize sets the uint32 value in the RoCEv2FlowSettings object
func (obj *roCEv2FlowSettings) SetMessageSize(value uint32) RoCEv2FlowSettings {

	obj.obj.MessageSize = &value
	return obj
}

type RoCEv2FlowSettingsMessageSizeUnitEnum string

// Enum of MessageSizeUnit on RoCEv2FlowSettings
var RoCEv2FlowSettingsMessageSizeUnit = struct {
	BYTE RoCEv2FlowSettingsMessageSizeUnitEnum
	KB   RoCEv2FlowSettingsMessageSizeUnitEnum
	MB   RoCEv2FlowSettingsMessageSizeUnitEnum
}{
	BYTE: RoCEv2FlowSettingsMessageSizeUnitEnum("Byte"),
	KB:   RoCEv2FlowSettingsMessageSizeUnitEnum("KB"),
	MB:   RoCEv2FlowSettingsMessageSizeUnitEnum("MB"),
}

func (obj *roCEv2FlowSettings) MessageSizeUnit() RoCEv2FlowSettingsMessageSizeUnitEnum {
	return RoCEv2FlowSettingsMessageSizeUnitEnum(obj.obj.MessageSizeUnit.Enum().String())
}

// Unit of the transfer message size. Available options are Bytes, KB, MB.
// MessageSizeUnit returns a string
func (obj *roCEv2FlowSettings) HasMessageSizeUnit() bool {
	return obj.obj.MessageSizeUnit != nil
}

func (obj *roCEv2FlowSettings) SetMessageSizeUnit(value RoCEv2FlowSettingsMessageSizeUnitEnum) RoCEv2FlowSettings {
	intValue, ok := otg.RoCEv2FlowSettings_MessageSizeUnit_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on RoCEv2FlowSettingsMessageSizeUnitEnum", string(value)))
		return obj
	}
	enumValue := otg.RoCEv2FlowSettings_MessageSizeUnit_Enum(intValue)
	obj.obj.MessageSizeUnit = &enumValue

	return obj
}

func (obj *roCEv2FlowSettings) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.CustomQpNumber != nil {

		if *obj.obj.CustomQpNumber < 2 || *obj.obj.CustomQpNumber > 33554431 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("2 <= RoCEv2FlowSettings.CustomQpNumber <= 33554431 but Got %d", *obj.obj.CustomQpNumber))
		}

	}

	if obj.obj.Dscp != nil {

		if *obj.obj.Dscp > 63 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RoCEv2FlowSettings.Dscp <= 63 but Got %d", *obj.obj.Dscp))
		}

	}

	if obj.obj.Ecn != nil {

		if *obj.obj.Ecn > 3 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RoCEv2FlowSettings.Ecn <= 3 but Got %d", *obj.obj.Ecn))
		}

	}

	if obj.obj.UdpSourcePort != nil {

		if *obj.obj.UdpSourcePort > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RoCEv2FlowSettings.UdpSourcePort <= 65535 but Got %d", *obj.obj.UdpSourcePort))
		}

	}

	if obj.obj.ImmediateData != nil {

		err := obj.validateHex(obj.ImmediateData())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on RoCEv2FlowSettings.ImmediateData"))
		}

	}

	if obj.obj.MessageSize != nil {

		if *obj.obj.MessageSize > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RoCEv2FlowSettings.MessageSize <= 65535 but Got %d", *obj.obj.MessageSize))
		}

	}

}

func (obj *roCEv2FlowSettings) setDefault() {
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
	if obj.obj.Rocev2Verb == nil {
		obj.SetRocev2Verb(RoCEv2FlowSettingsRocev2Verb.WRITE)

	}
	if obj.obj.ImmediateData == nil {
		obj.SetImmediateData("00000000")
	}
	if obj.obj.MessageSize == nil {
		obj.SetMessageSize(1)
	}
	if obj.obj.MessageSizeUnit == nil {
		obj.SetMessageSizeUnit(RoCEv2FlowSettingsMessageSizeUnit.MB)

	}

}
