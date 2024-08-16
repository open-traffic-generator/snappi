package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6InterfaceState *****
type dhcpv6InterfaceState struct {
	validation
	obj          *otg.Dhcpv6InterfaceState
	marshaller   marshalDhcpv6InterfaceState
	unMarshaller unMarshalDhcpv6InterfaceState
}

func NewDhcpv6InterfaceState() Dhcpv6InterfaceState {
	obj := dhcpv6InterfaceState{obj: &otg.Dhcpv6InterfaceState{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6InterfaceState) msg() *otg.Dhcpv6InterfaceState {
	return obj.obj
}

func (obj *dhcpv6InterfaceState) setMsg(msg *otg.Dhcpv6InterfaceState) Dhcpv6InterfaceState {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6InterfaceState struct {
	obj *dhcpv6InterfaceState
}

type marshalDhcpv6InterfaceState interface {
	// ToProto marshals Dhcpv6InterfaceState to protobuf object *otg.Dhcpv6InterfaceState
	ToProto() (*otg.Dhcpv6InterfaceState, error)
	// ToPbText marshals Dhcpv6InterfaceState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6InterfaceState to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6InterfaceState to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpv6InterfaceState struct {
	obj *dhcpv6InterfaceState
}

type unMarshalDhcpv6InterfaceState interface {
	// FromProto unmarshals Dhcpv6InterfaceState from protobuf object *otg.Dhcpv6InterfaceState
	FromProto(msg *otg.Dhcpv6InterfaceState) (Dhcpv6InterfaceState, error)
	// FromPbText unmarshals Dhcpv6InterfaceState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6InterfaceState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6InterfaceState from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6InterfaceState) Marshal() marshalDhcpv6InterfaceState {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6InterfaceState{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6InterfaceState) Unmarshal() unMarshalDhcpv6InterfaceState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6InterfaceState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6InterfaceState) ToProto() (*otg.Dhcpv6InterfaceState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6InterfaceState) FromProto(msg *otg.Dhcpv6InterfaceState) (Dhcpv6InterfaceState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6InterfaceState) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6InterfaceState) FromPbText(value string) error {
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

func (m *marshaldhcpv6InterfaceState) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6InterfaceState) FromYaml(value string) error {
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

func (m *marshaldhcpv6InterfaceState) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6InterfaceState) FromJson(value string) error {
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

func (obj *dhcpv6InterfaceState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6InterfaceState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6InterfaceState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6InterfaceState) Clone() (Dhcpv6InterfaceState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6InterfaceState()
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

// Dhcpv6InterfaceState is the IPv6 address associated with this DHCP Client session.
type Dhcpv6InterfaceState interface {
	Validation
	// msg marshals Dhcpv6InterfaceState to protobuf object *otg.Dhcpv6InterfaceState
	// and doesn't set defaults
	msg() *otg.Dhcpv6InterfaceState
	// setMsg unmarshals Dhcpv6InterfaceState from protobuf object *otg.Dhcpv6InterfaceState
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6InterfaceState) Dhcpv6InterfaceState
	// provides marshal interface
	Marshal() marshalDhcpv6InterfaceState
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6InterfaceState
	// validate validates Dhcpv6InterfaceState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6InterfaceState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// DhcpClientName returns string, set in Dhcpv6InterfaceState.
	DhcpClientName() string
	// SetDhcpClientName assigns string provided by user to Dhcpv6InterfaceState
	SetDhcpClientName(value string) Dhcpv6InterfaceState
	// HasDhcpClientName checks if DhcpClientName has been set in Dhcpv6InterfaceState
	HasDhcpClientName() bool
	// Ipv6IapdAddress returns string, set in Dhcpv6InterfaceState.
	Ipv6IapdAddress() string
	// SetIpv6IapdAddress assigns string provided by user to Dhcpv6InterfaceState
	SetIpv6IapdAddress(value string) Dhcpv6InterfaceState
	// HasIpv6IapdAddress checks if Ipv6IapdAddress has been set in Dhcpv6InterfaceState
	HasIpv6IapdAddress() bool
	// IapdPrefixLength returns uint32, set in Dhcpv6InterfaceState.
	IapdPrefixLength() uint32
	// SetIapdPrefixLength assigns uint32 provided by user to Dhcpv6InterfaceState
	SetIapdPrefixLength(value uint32) Dhcpv6InterfaceState
	// HasIapdPrefixLength checks if IapdPrefixLength has been set in Dhcpv6InterfaceState
	HasIapdPrefixLength() bool
	// Ipv6Address returns string, set in Dhcpv6InterfaceState.
	Ipv6Address() string
	// SetIpv6Address assigns string provided by user to Dhcpv6InterfaceState
	SetIpv6Address(value string) Dhcpv6InterfaceState
	// HasIpv6Address checks if Ipv6Address has been set in Dhcpv6InterfaceState
	HasIpv6Address() bool
	// GatewayAddress returns string, set in Dhcpv6InterfaceState.
	GatewayAddress() string
	// SetGatewayAddress assigns string provided by user to Dhcpv6InterfaceState
	SetGatewayAddress(value string) Dhcpv6InterfaceState
	// HasGatewayAddress checks if GatewayAddress has been set in Dhcpv6InterfaceState
	HasGatewayAddress() bool
	// LeaseTime returns uint32, set in Dhcpv6InterfaceState.
	LeaseTime() uint32
	// SetLeaseTime assigns uint32 provided by user to Dhcpv6InterfaceState
	SetLeaseTime(value uint32) Dhcpv6InterfaceState
	// HasLeaseTime checks if LeaseTime has been set in Dhcpv6InterfaceState
	HasLeaseTime() bool
	// RenewTime returns uint32, set in Dhcpv6InterfaceState.
	RenewTime() uint32
	// SetRenewTime assigns uint32 provided by user to Dhcpv6InterfaceState
	SetRenewTime(value uint32) Dhcpv6InterfaceState
	// HasRenewTime checks if RenewTime has been set in Dhcpv6InterfaceState
	HasRenewTime() bool
	// RebindTime returns uint32, set in Dhcpv6InterfaceState.
	RebindTime() uint32
	// SetRebindTime assigns uint32 provided by user to Dhcpv6InterfaceState
	SetRebindTime(value uint32) Dhcpv6InterfaceState
	// HasRebindTime checks if RebindTime has been set in Dhcpv6InterfaceState
	HasRebindTime() bool
}

// The name of a DHCPv6 Client.
// DhcpClientName returns a string
func (obj *dhcpv6InterfaceState) DhcpClientName() string {

	return *obj.obj.DhcpClientName

}

// The name of a DHCPv6 Client.
// DhcpClientName returns a string
func (obj *dhcpv6InterfaceState) HasDhcpClientName() bool {
	return obj.obj.DhcpClientName != nil
}

// The name of a DHCPv6 Client.
// SetDhcpClientName sets the string value in the Dhcpv6InterfaceState object
func (obj *dhcpv6InterfaceState) SetDhcpClientName(value string) Dhcpv6InterfaceState {

	obj.obj.DhcpClientName = &value
	return obj
}

// The IPv6 IAPD address associated with this DHCP Client session.
// Ipv6IapdAddress returns a string
func (obj *dhcpv6InterfaceState) Ipv6IapdAddress() string {

	return *obj.obj.Ipv6IapdAddress

}

// The IPv6 IAPD address associated with this DHCP Client session.
// Ipv6IapdAddress returns a string
func (obj *dhcpv6InterfaceState) HasIpv6IapdAddress() bool {
	return obj.obj.Ipv6IapdAddress != nil
}

// The IPv6 IAPD address associated with this DHCP Client session.
// SetIpv6IapdAddress sets the string value in the Dhcpv6InterfaceState object
func (obj *dhcpv6InterfaceState) SetIpv6IapdAddress(value string) Dhcpv6InterfaceState {

	obj.obj.Ipv6IapdAddress = &value
	return obj
}

// The prefix length of the IPv6 IAPD address associated with this DHCP Client session.
// IapdPrefixLength returns a uint32
func (obj *dhcpv6InterfaceState) IapdPrefixLength() uint32 {

	return *obj.obj.IapdPrefixLength

}

// The prefix length of the IPv6 IAPD address associated with this DHCP Client session.
// IapdPrefixLength returns a uint32
func (obj *dhcpv6InterfaceState) HasIapdPrefixLength() bool {
	return obj.obj.IapdPrefixLength != nil
}

// The prefix length of the IPv6 IAPD address associated with this DHCP Client session.
// SetIapdPrefixLength sets the uint32 value in the Dhcpv6InterfaceState object
func (obj *dhcpv6InterfaceState) SetIapdPrefixLength(value uint32) Dhcpv6InterfaceState {

	obj.obj.IapdPrefixLength = &value
	return obj
}

// The IPv6 address associated with this DHCP Client session.
// Ipv6Address returns a string
func (obj *dhcpv6InterfaceState) Ipv6Address() string {

	return *obj.obj.Ipv6Address

}

// The IPv6 address associated with this DHCP Client session.
// Ipv6Address returns a string
func (obj *dhcpv6InterfaceState) HasIpv6Address() bool {
	return obj.obj.Ipv6Address != nil
}

// The IPv6 address associated with this DHCP Client session.
// SetIpv6Address sets the string value in the Dhcpv6InterfaceState object
func (obj *dhcpv6InterfaceState) SetIpv6Address(value string) Dhcpv6InterfaceState {

	obj.obj.Ipv6Address = &value
	return obj
}

// The Gateway IPV6 address associated with this DHCP Client session.
// GatewayAddress returns a string
func (obj *dhcpv6InterfaceState) GatewayAddress() string {

	return *obj.obj.GatewayAddress

}

// The Gateway IPV6 address associated with this DHCP Client session.
// GatewayAddress returns a string
func (obj *dhcpv6InterfaceState) HasGatewayAddress() bool {
	return obj.obj.GatewayAddress != nil
}

// The Gateway IPV6 address associated with this DHCP Client session.
// SetGatewayAddress sets the string value in the Dhcpv6InterfaceState object
func (obj *dhcpv6InterfaceState) SetGatewayAddress(value string) Dhcpv6InterfaceState {

	obj.obj.GatewayAddress = &value
	return obj
}

// The duration of the IPv6 address lease, in seconds.
// LeaseTime returns a uint32
func (obj *dhcpv6InterfaceState) LeaseTime() uint32 {

	return *obj.obj.LeaseTime

}

// The duration of the IPv6 address lease, in seconds.
// LeaseTime returns a uint32
func (obj *dhcpv6InterfaceState) HasLeaseTime() bool {
	return obj.obj.LeaseTime != nil
}

// The duration of the IPv6 address lease, in seconds.
// SetLeaseTime sets the uint32 value in the Dhcpv6InterfaceState object
func (obj *dhcpv6InterfaceState) SetLeaseTime(value uint32) Dhcpv6InterfaceState {

	obj.obj.LeaseTime = &value
	return obj
}

// Time in seconds until the DHCPv6 client starts renewing the lease.
// RenewTime returns a uint32
func (obj *dhcpv6InterfaceState) RenewTime() uint32 {

	return *obj.obj.RenewTime

}

// Time in seconds until the DHCPv6 client starts renewing the lease.
// RenewTime returns a uint32
func (obj *dhcpv6InterfaceState) HasRenewTime() bool {
	return obj.obj.RenewTime != nil
}

// Time in seconds until the DHCPv6 client starts renewing the lease.
// SetRenewTime sets the uint32 value in the Dhcpv6InterfaceState object
func (obj *dhcpv6InterfaceState) SetRenewTime(value uint32) Dhcpv6InterfaceState {

	obj.obj.RenewTime = &value
	return obj
}

// Time in seconds until the DHCPv6 client starts rebinding.
// RebindTime returns a uint32
func (obj *dhcpv6InterfaceState) RebindTime() uint32 {

	return *obj.obj.RebindTime

}

// Time in seconds until the DHCPv6 client starts rebinding.
// RebindTime returns a uint32
func (obj *dhcpv6InterfaceState) HasRebindTime() bool {
	return obj.obj.RebindTime != nil
}

// Time in seconds until the DHCPv6 client starts rebinding.
// SetRebindTime sets the uint32 value in the Dhcpv6InterfaceState object
func (obj *dhcpv6InterfaceState) SetRebindTime(value uint32) Dhcpv6InterfaceState {

	obj.obj.RebindTime = &value
	return obj
}

func (obj *dhcpv6InterfaceState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.IapdPrefixLength != nil {

		if *obj.obj.IapdPrefixLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Dhcpv6InterfaceState.IapdPrefixLength <= 128 but Got %d", *obj.obj.IapdPrefixLength))
		}

	}

}

func (obj *dhcpv6InterfaceState) setDefault() {

}
