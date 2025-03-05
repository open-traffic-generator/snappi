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
	obj                 *otg.RoCEv2FlowSettings
	marshaller          marshalRoCEv2FlowSettings
	unMarshaller        unMarshalRoCEv2FlowSettings
	dcqcnSettingaHolder RoCEv2DCQCN
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
	obj.setNil()
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
	m.obj.setNil()
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
	m.obj.setNil()
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
	m.obj.setNil()
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

func (obj *roCEv2FlowSettings) setNil() {
	obj.dcqcnSettingaHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// RoCEv2FlowSettings is this configuration allows you to configure RDMA flow over the same QP number from same source and destination.
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
	// ConnectionType returns RoCEv2FlowSettingsConnectionTypeEnum, set in RoCEv2FlowSettings
	ConnectionType() RoCEv2FlowSettingsConnectionTypeEnum
	// SetConnectionType assigns RoCEv2FlowSettingsConnectionTypeEnum provided by user to RoCEv2FlowSettings
	SetConnectionType(value RoCEv2FlowSettingsConnectionTypeEnum) RoCEv2FlowSettings
	// HasConnectionType checks if ConnectionType has been set in RoCEv2FlowSettings
	HasConnectionType() bool
	// SourceQpNumber returns uint32, set in RoCEv2FlowSettings.
	SourceQpNumber() uint32
	// SetSourceQpNumber assigns uint32 provided by user to RoCEv2FlowSettings
	SetSourceQpNumber(value uint32) RoCEv2FlowSettings
	// HasSourceQpNumber checks if SourceQpNumber has been set in RoCEv2FlowSettings
	HasSourceQpNumber() bool
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
	// DcqcnSettinga returns RoCEv2DCQCN, set in RoCEv2FlowSettings.
	// RoCEv2DCQCN is roCEv2 DCQCN Settings.
	DcqcnSettinga() RoCEv2DCQCN
	// SetDcqcnSettinga assigns RoCEv2DCQCN provided by user to RoCEv2FlowSettings.
	// RoCEv2DCQCN is roCEv2 DCQCN Settings.
	SetDcqcnSettinga(value RoCEv2DCQCN) RoCEv2FlowSettings
	// HasDcqcnSettinga checks if DcqcnSettinga has been set in RoCEv2FlowSettings
	HasDcqcnSettinga() bool
	setNil()
}

type RoCEv2FlowSettingsConnectionTypeEnum string

// Enum of ConnectionType on RoCEv2FlowSettings
var RoCEv2FlowSettingsConnectionType = struct {
	RELIABLE_CONNECTION          RoCEv2FlowSettingsConnectionTypeEnum
	RELIABLE_DATAGRAM            RoCEv2FlowSettingsConnectionTypeEnum
	EXTENDED_RELIABLE_CONNECTION RoCEv2FlowSettingsConnectionTypeEnum
	UNRELIABLE_DATAGRAM          RoCEv2FlowSettingsConnectionTypeEnum
	UNRELIABLE_CONNECTION        RoCEv2FlowSettingsConnectionTypeEnum
	RAW_IPV6_DATAGRAM            RoCEv2FlowSettingsConnectionTypeEnum
	RAW_ETHERNET_DATAGRAM        RoCEv2FlowSettingsConnectionTypeEnum
}{
	RELIABLE_CONNECTION:          RoCEv2FlowSettingsConnectionTypeEnum("reliable_connection"),
	RELIABLE_DATAGRAM:            RoCEv2FlowSettingsConnectionTypeEnum("reliable_datagram"),
	EXTENDED_RELIABLE_CONNECTION: RoCEv2FlowSettingsConnectionTypeEnum("extended_reliable_connection"),
	UNRELIABLE_DATAGRAM:          RoCEv2FlowSettingsConnectionTypeEnum("unreliable_datagram"),
	UNRELIABLE_CONNECTION:        RoCEv2FlowSettingsConnectionTypeEnum("unreliable_connection"),
	RAW_IPV6_DATAGRAM:            RoCEv2FlowSettingsConnectionTypeEnum("raw_ipv6_datagram"),
	RAW_ETHERNET_DATAGRAM:        RoCEv2FlowSettingsConnectionTypeEnum("raw_ethernet_datagram"),
}

func (obj *roCEv2FlowSettings) ConnectionType() RoCEv2FlowSettingsConnectionTypeEnum {
	return RoCEv2FlowSettingsConnectionTypeEnum(obj.obj.ConnectionType.Enum().String())
}

// There are multiple connection types. Valid values are :  Reliable Connection (RC), Reliable Datagram (RD), Extended Reliable Connection (XRC), Unreliable Datagram (UD),  Unreliable Connection (UC), Raw IPv6 Datagram, Raw Ethertype Datagram.
// ConnectionType returns a string
func (obj *roCEv2FlowSettings) HasConnectionType() bool {
	return obj.obj.ConnectionType != nil
}

func (obj *roCEv2FlowSettings) SetConnectionType(value RoCEv2FlowSettingsConnectionTypeEnum) RoCEv2FlowSettings {
	intValue, ok := otg.RoCEv2FlowSettings_ConnectionType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on RoCEv2FlowSettingsConnectionTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.RoCEv2FlowSettings_ConnectionType_Enum(intValue)
	obj.obj.ConnectionType = &enumValue

	return obj
}

// Configure the QP range.
// SourceQpNumber returns a uint32
func (obj *roCEv2FlowSettings) SourceQpNumber() uint32 {

	return *obj.obj.SourceQpNumber

}

// Configure the QP range.
// SourceQpNumber returns a uint32
func (obj *roCEv2FlowSettings) HasSourceQpNumber() bool {
	return obj.obj.SourceQpNumber != nil
}

// Configure the QP range.
// SetSourceQpNumber sets the uint32 value in the RoCEv2FlowSettings object
func (obj *roCEv2FlowSettings) SetSourceQpNumber(value uint32) RoCEv2FlowSettings {

	obj.obj.SourceQpNumber = &value
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
	READ                 RoCEv2FlowSettingsRocev2VerbEnum
}{
	NONE:                 RoCEv2FlowSettingsRocev2VerbEnum("none"),
	WRITE:                RoCEv2FlowSettingsRocev2VerbEnum("write"),
	WRITE_WITH_IMMEDIATE: RoCEv2FlowSettingsRocev2VerbEnum("write_with_immediate"),
	SEND:                 RoCEv2FlowSettingsRocev2VerbEnum("send"),
	SEND_WITH_IMMEDIATE:  RoCEv2FlowSettingsRocev2VerbEnum("send_with_immediate"),
	READ:                 RoCEv2FlowSettingsRocev2VerbEnum("read"),
}

func (obj *roCEv2FlowSettings) Rocev2Verb() RoCEv2FlowSettingsRocev2VerbEnum {
	return RoCEv2FlowSettingsRocev2VerbEnum(obj.obj.Rocev2Verb.Enum().String())
}

// RoCEv2 Verb, Available options are: none, write, wrtie_with_immediate, send, send_with_immediate and read: The corresponding flow will not take part in traffic.
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

// The Maximum message size that is allowed to transfer depends on the MTU size and the number of VLANs configured on the interfaces.
// MessageSize returns a uint32
func (obj *roCEv2FlowSettings) MessageSize() uint32 {

	return *obj.obj.MessageSize

}

// The Maximum message size that is allowed to transfer depends on the MTU size and the number of VLANs configured on the interfaces.
// MessageSize returns a uint32
func (obj *roCEv2FlowSettings) HasMessageSize() bool {
	return obj.obj.MessageSize != nil
}

// The Maximum message size that is allowed to transfer depends on the MTU size and the number of VLANs configured on the interfaces.
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
	GB   RoCEv2FlowSettingsMessageSizeUnitEnum
}{
	BYTE: RoCEv2FlowSettingsMessageSizeUnitEnum("Byte"),
	KB:   RoCEv2FlowSettingsMessageSizeUnitEnum("KB"),
	MB:   RoCEv2FlowSettingsMessageSizeUnitEnum("MB"),
	GB:   RoCEv2FlowSettingsMessageSizeUnitEnum("GB"),
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

// description is TBD
// DcqcnSettinga returns a RoCEv2DCQCN
func (obj *roCEv2FlowSettings) DcqcnSettinga() RoCEv2DCQCN {
	if obj.obj.DcqcnSettinga == nil {
		obj.obj.DcqcnSettinga = NewRoCEv2DCQCN().msg()
	}
	if obj.dcqcnSettingaHolder == nil {
		obj.dcqcnSettingaHolder = &roCEv2DCQCN{obj: obj.obj.DcqcnSettinga}
	}
	return obj.dcqcnSettingaHolder
}

// description is TBD
// DcqcnSettinga returns a RoCEv2DCQCN
func (obj *roCEv2FlowSettings) HasDcqcnSettinga() bool {
	return obj.obj.DcqcnSettinga != nil
}

// description is TBD
// SetDcqcnSettinga sets the RoCEv2DCQCN value in the RoCEv2FlowSettings object
func (obj *roCEv2FlowSettings) SetDcqcnSettinga(value RoCEv2DCQCN) RoCEv2FlowSettings {

	obj.dcqcnSettingaHolder = nil
	obj.obj.DcqcnSettinga = value.msg()

	return obj
}

func (obj *roCEv2FlowSettings) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SourceQpNumber != nil {

		if *obj.obj.SourceQpNumber < 2 || *obj.obj.SourceQpNumber > 33554431 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("2 <= RoCEv2FlowSettings.SourceQpNumber <= 33554431 but Got %d", *obj.obj.SourceQpNumber))
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

	if obj.obj.DcqcnSettinga != nil {

		obj.DcqcnSettinga().validateObj(vObj, set_default)
	}

}

func (obj *roCEv2FlowSettings) setDefault() {
	if obj.obj.SourceQpNumber == nil {
		obj.SetSourceQpNumber(2)
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
