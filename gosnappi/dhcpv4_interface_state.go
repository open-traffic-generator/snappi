package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv4InterfaceState *****
type dhcpv4InterfaceState struct {
	validation
	obj          *otg.Dhcpv4InterfaceState
	marshaller   marshalDhcpv4InterfaceState
	unMarshaller unMarshalDhcpv4InterfaceState
}

func NewDhcpv4InterfaceState() Dhcpv4InterfaceState {
	obj := dhcpv4InterfaceState{obj: &otg.Dhcpv4InterfaceState{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv4InterfaceState) msg() *otg.Dhcpv4InterfaceState {
	return obj.obj
}

func (obj *dhcpv4InterfaceState) setMsg(msg *otg.Dhcpv4InterfaceState) Dhcpv4InterfaceState {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv4InterfaceState struct {
	obj *dhcpv4InterfaceState
}

type marshalDhcpv4InterfaceState interface {
	// ToProto marshals Dhcpv4InterfaceState to protobuf object *otg.Dhcpv4InterfaceState
	ToProto() (*otg.Dhcpv4InterfaceState, error)
	// ToPbText marshals Dhcpv4InterfaceState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv4InterfaceState to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv4InterfaceState to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Dhcpv4InterfaceState to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldhcpv4InterfaceState struct {
	obj *dhcpv4InterfaceState
}

type unMarshalDhcpv4InterfaceState interface {
	// FromProto unmarshals Dhcpv4InterfaceState from protobuf object *otg.Dhcpv4InterfaceState
	FromProto(msg *otg.Dhcpv4InterfaceState) (Dhcpv4InterfaceState, error)
	// FromPbText unmarshals Dhcpv4InterfaceState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv4InterfaceState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv4InterfaceState from JSON text
	FromJson(value string) error
}

func (obj *dhcpv4InterfaceState) Marshal() marshalDhcpv4InterfaceState {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv4InterfaceState{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv4InterfaceState) Unmarshal() unMarshalDhcpv4InterfaceState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv4InterfaceState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv4InterfaceState) ToProto() (*otg.Dhcpv4InterfaceState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv4InterfaceState) FromProto(msg *otg.Dhcpv4InterfaceState) (Dhcpv4InterfaceState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv4InterfaceState) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv4InterfaceState) FromPbText(value string) error {
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

func (m *marshaldhcpv4InterfaceState) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv4InterfaceState) FromYaml(value string) error {
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

func (m *marshaldhcpv4InterfaceState) ToJsonRaw() (string, error) {
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

func (m *marshaldhcpv4InterfaceState) ToJson() (string, error) {
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

func (m *unMarshaldhcpv4InterfaceState) FromJson(value string) error {
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

func (obj *dhcpv4InterfaceState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv4InterfaceState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv4InterfaceState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv4InterfaceState) Clone() (Dhcpv4InterfaceState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv4InterfaceState()
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

// Dhcpv4InterfaceState is the IPv4 address associated with this DHCP Client session.
type Dhcpv4InterfaceState interface {
	Validation
	// msg marshals Dhcpv4InterfaceState to protobuf object *otg.Dhcpv4InterfaceState
	// and doesn't set defaults
	msg() *otg.Dhcpv4InterfaceState
	// setMsg unmarshals Dhcpv4InterfaceState from protobuf object *otg.Dhcpv4InterfaceState
	// and doesn't set defaults
	setMsg(*otg.Dhcpv4InterfaceState) Dhcpv4InterfaceState
	// provides marshal interface
	Marshal() marshalDhcpv4InterfaceState
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv4InterfaceState
	// validate validates Dhcpv4InterfaceState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv4InterfaceState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// DhcpClientName returns string, set in Dhcpv4InterfaceState.
	DhcpClientName() string
	// SetDhcpClientName assigns string provided by user to Dhcpv4InterfaceState
	SetDhcpClientName(value string) Dhcpv4InterfaceState
	// HasDhcpClientName checks if DhcpClientName has been set in Dhcpv4InterfaceState
	HasDhcpClientName() bool
	// Ipv4Address returns string, set in Dhcpv4InterfaceState.
	Ipv4Address() string
	// SetIpv4Address assigns string provided by user to Dhcpv4InterfaceState
	SetIpv4Address(value string) Dhcpv4InterfaceState
	// HasIpv4Address checks if Ipv4Address has been set in Dhcpv4InterfaceState
	HasIpv4Address() bool
	// PrefixLength returns uint32, set in Dhcpv4InterfaceState.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to Dhcpv4InterfaceState
	SetPrefixLength(value uint32) Dhcpv4InterfaceState
	// HasPrefixLength checks if PrefixLength has been set in Dhcpv4InterfaceState
	HasPrefixLength() bool
	// GatewayAddress returns string, set in Dhcpv4InterfaceState.
	GatewayAddress() string
	// SetGatewayAddress assigns string provided by user to Dhcpv4InterfaceState
	SetGatewayAddress(value string) Dhcpv4InterfaceState
	// HasGatewayAddress checks if GatewayAddress has been set in Dhcpv4InterfaceState
	HasGatewayAddress() bool
	// LeaseTime returns uint32, set in Dhcpv4InterfaceState.
	LeaseTime() uint32
	// SetLeaseTime assigns uint32 provided by user to Dhcpv4InterfaceState
	SetLeaseTime(value uint32) Dhcpv4InterfaceState
	// HasLeaseTime checks if LeaseTime has been set in Dhcpv4InterfaceState
	HasLeaseTime() bool
	// RenewTime returns uint32, set in Dhcpv4InterfaceState.
	RenewTime() uint32
	// SetRenewTime assigns uint32 provided by user to Dhcpv4InterfaceState
	SetRenewTime(value uint32) Dhcpv4InterfaceState
	// HasRenewTime checks if RenewTime has been set in Dhcpv4InterfaceState
	HasRenewTime() bool
	// RebindTime returns uint32, set in Dhcpv4InterfaceState.
	RebindTime() uint32
	// SetRebindTime assigns uint32 provided by user to Dhcpv4InterfaceState
	SetRebindTime(value uint32) Dhcpv4InterfaceState
	// HasRebindTime checks if RebindTime has been set in Dhcpv4InterfaceState
	HasRebindTime() bool
}

// The name of a DHCPv4 Client.
// DhcpClientName returns a string
func (obj *dhcpv4InterfaceState) DhcpClientName() string {

	return *obj.obj.DhcpClientName

}

// The name of a DHCPv4 Client.
// DhcpClientName returns a string
func (obj *dhcpv4InterfaceState) HasDhcpClientName() bool {
	return obj.obj.DhcpClientName != nil
}

// The name of a DHCPv4 Client.
// SetDhcpClientName sets the string value in the Dhcpv4InterfaceState object
func (obj *dhcpv4InterfaceState) SetDhcpClientName(value string) Dhcpv4InterfaceState {

	obj.obj.DhcpClientName = &value
	return obj
}

// The IPv4 address associated with this DHCP Client session.
// Ipv4Address returns a string
func (obj *dhcpv4InterfaceState) Ipv4Address() string {

	return *obj.obj.Ipv4Address

}

// The IPv4 address associated with this DHCP Client session.
// Ipv4Address returns a string
func (obj *dhcpv4InterfaceState) HasIpv4Address() bool {
	return obj.obj.Ipv4Address != nil
}

// The IPv4 address associated with this DHCP Client session.
// SetIpv4Address sets the string value in the Dhcpv4InterfaceState object
func (obj *dhcpv4InterfaceState) SetIpv4Address(value string) Dhcpv4InterfaceState {

	obj.obj.Ipv4Address = &value
	return obj
}

// The length of the prefix.
// PrefixLength returns a uint32
func (obj *dhcpv4InterfaceState) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// The length of the prefix.
// PrefixLength returns a uint32
func (obj *dhcpv4InterfaceState) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// The length of the prefix.
// SetPrefixLength sets the uint32 value in the Dhcpv4InterfaceState object
func (obj *dhcpv4InterfaceState) SetPrefixLength(value uint32) Dhcpv4InterfaceState {

	obj.obj.PrefixLength = &value
	return obj
}

// The Gateway Ipv4 address associated with this DHCP Client session.
// GatewayAddress returns a string
func (obj *dhcpv4InterfaceState) GatewayAddress() string {

	return *obj.obj.GatewayAddress

}

// The Gateway Ipv4 address associated with this DHCP Client session.
// GatewayAddress returns a string
func (obj *dhcpv4InterfaceState) HasGatewayAddress() bool {
	return obj.obj.GatewayAddress != nil
}

// The Gateway Ipv4 address associated with this DHCP Client session.
// SetGatewayAddress sets the string value in the Dhcpv4InterfaceState object
func (obj *dhcpv4InterfaceState) SetGatewayAddress(value string) Dhcpv4InterfaceState {

	obj.obj.GatewayAddress = &value
	return obj
}

// The duration of the IPv4 address lease, in seconds.
// LeaseTime returns a uint32
func (obj *dhcpv4InterfaceState) LeaseTime() uint32 {

	return *obj.obj.LeaseTime

}

// The duration of the IPv4 address lease, in seconds.
// LeaseTime returns a uint32
func (obj *dhcpv4InterfaceState) HasLeaseTime() bool {
	return obj.obj.LeaseTime != nil
}

// The duration of the IPv4 address lease, in seconds.
// SetLeaseTime sets the uint32 value in the Dhcpv4InterfaceState object
func (obj *dhcpv4InterfaceState) SetLeaseTime(value uint32) Dhcpv4InterfaceState {

	obj.obj.LeaseTime = &value
	return obj
}

// Time in seconds until the DHCPv4 client starts renewing the lease.
// RenewTime returns a uint32
func (obj *dhcpv4InterfaceState) RenewTime() uint32 {

	return *obj.obj.RenewTime

}

// Time in seconds until the DHCPv4 client starts renewing the lease.
// RenewTime returns a uint32
func (obj *dhcpv4InterfaceState) HasRenewTime() bool {
	return obj.obj.RenewTime != nil
}

// Time in seconds until the DHCPv4 client starts renewing the lease.
// SetRenewTime sets the uint32 value in the Dhcpv4InterfaceState object
func (obj *dhcpv4InterfaceState) SetRenewTime(value uint32) Dhcpv4InterfaceState {

	obj.obj.RenewTime = &value
	return obj
}

// Time in seconds until the DHCPv4 client starts rebinding.
// RebindTime returns a uint32
func (obj *dhcpv4InterfaceState) RebindTime() uint32 {

	return *obj.obj.RebindTime

}

// Time in seconds until the DHCPv4 client starts rebinding.
// RebindTime returns a uint32
func (obj *dhcpv4InterfaceState) HasRebindTime() bool {
	return obj.obj.RebindTime != nil
}

// Time in seconds until the DHCPv4 client starts rebinding.
// SetRebindTime sets the uint32 value in the Dhcpv4InterfaceState object
func (obj *dhcpv4InterfaceState) SetRebindTime(value uint32) Dhcpv4InterfaceState {

	obj.obj.RebindTime = &value
	return obj
}

func (obj *dhcpv4InterfaceState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Dhcpv4InterfaceState.PrefixLength <= 32 but Got %d", *obj.obj.PrefixLength))
		}

	}

}

func (obj *dhcpv4InterfaceState) setDefault() {

}
