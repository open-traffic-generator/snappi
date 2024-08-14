package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ServerLeaseState *****
type dhcpv6ServerLeaseState struct {
	validation
	obj          *otg.Dhcpv6ServerLeaseState
	marshaller   marshalDhcpv6ServerLeaseState
	unMarshaller unMarshalDhcpv6ServerLeaseState
}

func NewDhcpv6ServerLeaseState() Dhcpv6ServerLeaseState {
	obj := dhcpv6ServerLeaseState{obj: &otg.Dhcpv6ServerLeaseState{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ServerLeaseState) msg() *otg.Dhcpv6ServerLeaseState {
	return obj.obj
}

func (obj *dhcpv6ServerLeaseState) setMsg(msg *otg.Dhcpv6ServerLeaseState) Dhcpv6ServerLeaseState {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ServerLeaseState struct {
	obj *dhcpv6ServerLeaseState
}

type marshalDhcpv6ServerLeaseState interface {
	// ToProto marshals Dhcpv6ServerLeaseState to protobuf object *otg.Dhcpv6ServerLeaseState
	ToProto() (*otg.Dhcpv6ServerLeaseState, error)
	// ToPbText marshals Dhcpv6ServerLeaseState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ServerLeaseState to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ServerLeaseState to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpv6ServerLeaseState struct {
	obj *dhcpv6ServerLeaseState
}

type unMarshalDhcpv6ServerLeaseState interface {
	// FromProto unmarshals Dhcpv6ServerLeaseState from protobuf object *otg.Dhcpv6ServerLeaseState
	FromProto(msg *otg.Dhcpv6ServerLeaseState) (Dhcpv6ServerLeaseState, error)
	// FromPbText unmarshals Dhcpv6ServerLeaseState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ServerLeaseState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ServerLeaseState from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ServerLeaseState) Marshal() marshalDhcpv6ServerLeaseState {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ServerLeaseState{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ServerLeaseState) Unmarshal() unMarshalDhcpv6ServerLeaseState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ServerLeaseState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ServerLeaseState) ToProto() (*otg.Dhcpv6ServerLeaseState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ServerLeaseState) FromProto(msg *otg.Dhcpv6ServerLeaseState) (Dhcpv6ServerLeaseState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ServerLeaseState) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ServerLeaseState) FromPbText(value string) error {
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

func (m *marshaldhcpv6ServerLeaseState) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ServerLeaseState) FromYaml(value string) error {
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

func (m *marshaldhcpv6ServerLeaseState) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ServerLeaseState) FromJson(value string) error {
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

func (obj *dhcpv6ServerLeaseState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ServerLeaseState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ServerLeaseState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ServerLeaseState) Clone() (Dhcpv6ServerLeaseState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ServerLeaseState()
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

// Dhcpv6ServerLeaseState is iPv6 unicast prefix.
type Dhcpv6ServerLeaseState interface {
	Validation
	// msg marshals Dhcpv6ServerLeaseState to protobuf object *otg.Dhcpv6ServerLeaseState
	// and doesn't set defaults
	msg() *otg.Dhcpv6ServerLeaseState
	// setMsg unmarshals Dhcpv6ServerLeaseState from protobuf object *otg.Dhcpv6ServerLeaseState
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ServerLeaseState) Dhcpv6ServerLeaseState
	// provides marshal interface
	Marshal() marshalDhcpv6ServerLeaseState
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ServerLeaseState
	// validate validates Dhcpv6ServerLeaseState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ServerLeaseState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Address returns string, set in Dhcpv6ServerLeaseState.
	Address() string
	// SetAddress assigns string provided by user to Dhcpv6ServerLeaseState
	SetAddress(value string) Dhcpv6ServerLeaseState
	// HasAddress checks if Address has been set in Dhcpv6ServerLeaseState
	HasAddress() bool
	// ValidTime returns uint32, set in Dhcpv6ServerLeaseState.
	ValidTime() uint32
	// SetValidTime assigns uint32 provided by user to Dhcpv6ServerLeaseState
	SetValidTime(value uint32) Dhcpv6ServerLeaseState
	// HasValidTime checks if ValidTime has been set in Dhcpv6ServerLeaseState
	HasValidTime() bool
	// PreferredTime returns uint32, set in Dhcpv6ServerLeaseState.
	PreferredTime() uint32
	// SetPreferredTime assigns uint32 provided by user to Dhcpv6ServerLeaseState
	SetPreferredTime(value uint32) Dhcpv6ServerLeaseState
	// HasPreferredTime checks if PreferredTime has been set in Dhcpv6ServerLeaseState
	HasPreferredTime() bool
	// RenewTime returns uint32, set in Dhcpv6ServerLeaseState.
	RenewTime() uint32
	// SetRenewTime assigns uint32 provided by user to Dhcpv6ServerLeaseState
	SetRenewTime(value uint32) Dhcpv6ServerLeaseState
	// HasRenewTime checks if RenewTime has been set in Dhcpv6ServerLeaseState
	HasRenewTime() bool
	// RebindTime returns uint32, set in Dhcpv6ServerLeaseState.
	RebindTime() uint32
	// SetRebindTime assigns uint32 provided by user to Dhcpv6ServerLeaseState
	SetRebindTime(value uint32) Dhcpv6ServerLeaseState
	// HasRebindTime checks if RebindTime has been set in Dhcpv6ServerLeaseState
	HasRebindTime() bool
	// ClientId returns string, set in Dhcpv6ServerLeaseState.
	ClientId() string
	// SetClientId assigns string provided by user to Dhcpv6ServerLeaseState
	SetClientId(value string) Dhcpv6ServerLeaseState
	// HasClientId checks if ClientId has been set in Dhcpv6ServerLeaseState
	HasClientId() bool
	// RemoteId returns string, set in Dhcpv6ServerLeaseState.
	RemoteId() string
	// SetRemoteId assigns string provided by user to Dhcpv6ServerLeaseState
	SetRemoteId(value string) Dhcpv6ServerLeaseState
	// HasRemoteId checks if RemoteId has been set in Dhcpv6ServerLeaseState
	HasRemoteId() bool
	// InterfaceId returns string, set in Dhcpv6ServerLeaseState.
	InterfaceId() string
	// SetInterfaceId assigns string provided by user to Dhcpv6ServerLeaseState
	SetInterfaceId(value string) Dhcpv6ServerLeaseState
	// HasInterfaceId checks if InterfaceId has been set in Dhcpv6ServerLeaseState
	HasInterfaceId() bool
}

// The IPv6 address associated with this lease.
// Address returns a string
func (obj *dhcpv6ServerLeaseState) Address() string {

	return *obj.obj.Address

}

// The IPv6 address associated with this lease.
// Address returns a string
func (obj *dhcpv6ServerLeaseState) HasAddress() bool {
	return obj.obj.Address != nil
}

// The IPv6 address associated with this lease.
// SetAddress sets the string value in the Dhcpv6ServerLeaseState object
func (obj *dhcpv6ServerLeaseState) SetAddress(value string) Dhcpv6ServerLeaseState {

	obj.obj.Address = &value
	return obj
}

// The time in seconds, IP address lease will expire.
// ValidTime returns a uint32
func (obj *dhcpv6ServerLeaseState) ValidTime() uint32 {

	return *obj.obj.ValidTime

}

// The time in seconds, IP address lease will expire.
// ValidTime returns a uint32
func (obj *dhcpv6ServerLeaseState) HasValidTime() bool {
	return obj.obj.ValidTime != nil
}

// The time in seconds, IP address lease will expire.
// SetValidTime sets the uint32 value in the Dhcpv6ServerLeaseState object
func (obj *dhcpv6ServerLeaseState) SetValidTime(value uint32) Dhcpv6ServerLeaseState {

	obj.obj.ValidTime = &value
	return obj
}

// The time in seconds, elapsed time since address has been renewed.
// PreferredTime returns a uint32
func (obj *dhcpv6ServerLeaseState) PreferredTime() uint32 {

	return *obj.obj.PreferredTime

}

// The time in seconds, elapsed time since address has been renewed.
// PreferredTime returns a uint32
func (obj *dhcpv6ServerLeaseState) HasPreferredTime() bool {
	return obj.obj.PreferredTime != nil
}

// The time in seconds, elapsed time since address has been renewed.
// SetPreferredTime sets the uint32 value in the Dhcpv6ServerLeaseState object
func (obj *dhcpv6ServerLeaseState) SetPreferredTime(value uint32) Dhcpv6ServerLeaseState {

	obj.obj.PreferredTime = &value
	return obj
}

// Time in seconds until the DHCPv6 client starts renewing the lease.
// RenewTime returns a uint32
func (obj *dhcpv6ServerLeaseState) RenewTime() uint32 {

	return *obj.obj.RenewTime

}

// Time in seconds until the DHCPv6 client starts renewing the lease.
// RenewTime returns a uint32
func (obj *dhcpv6ServerLeaseState) HasRenewTime() bool {
	return obj.obj.RenewTime != nil
}

// Time in seconds until the DHCPv6 client starts renewing the lease.
// SetRenewTime sets the uint32 value in the Dhcpv6ServerLeaseState object
func (obj *dhcpv6ServerLeaseState) SetRenewTime(value uint32) Dhcpv6ServerLeaseState {

	obj.obj.RenewTime = &value
	return obj
}

// Time in seconds until the DHCPv6 client starts rebinding.
// RebindTime returns a uint32
func (obj *dhcpv6ServerLeaseState) RebindTime() uint32 {

	return *obj.obj.RebindTime

}

// Time in seconds until the DHCPv6 client starts rebinding.
// RebindTime returns a uint32
func (obj *dhcpv6ServerLeaseState) HasRebindTime() bool {
	return obj.obj.RebindTime != nil
}

// Time in seconds until the DHCPv6 client starts rebinding.
// SetRebindTime sets the uint32 value in the Dhcpv6ServerLeaseState object
func (obj *dhcpv6ServerLeaseState) SetRebindTime(value uint32) Dhcpv6ServerLeaseState {

	obj.obj.RebindTime = &value
	return obj
}

// The ID of the DHCPv6 client holding this lease.
// ClientId returns a string
func (obj *dhcpv6ServerLeaseState) ClientId() string {

	return *obj.obj.ClientId

}

// The ID of the DHCPv6 client holding this lease.
// ClientId returns a string
func (obj *dhcpv6ServerLeaseState) HasClientId() bool {
	return obj.obj.ClientId != nil
}

// The ID of the DHCPv6 client holding this lease.
// SetClientId sets the string value in the Dhcpv6ServerLeaseState object
func (obj *dhcpv6ServerLeaseState) SetClientId(value string) Dhcpv6ServerLeaseState {

	obj.obj.ClientId = &value
	return obj
}

// The Remote ID option found in the last request message.
// RemoteId returns a string
func (obj *dhcpv6ServerLeaseState) RemoteId() string {

	return *obj.obj.RemoteId

}

// The Remote ID option found in the last request message.
// RemoteId returns a string
func (obj *dhcpv6ServerLeaseState) HasRemoteId() bool {
	return obj.obj.RemoteId != nil
}

// The Remote ID option found in the last request message.
// SetRemoteId sets the string value in the Dhcpv6ServerLeaseState object
func (obj *dhcpv6ServerLeaseState) SetRemoteId(value string) Dhcpv6ServerLeaseState {

	obj.obj.RemoteId = &value
	return obj
}

// The Interface ID option found in the last request message.
// InterfaceId returns a string
func (obj *dhcpv6ServerLeaseState) InterfaceId() string {

	return *obj.obj.InterfaceId

}

// The Interface ID option found in the last request message.
// InterfaceId returns a string
func (obj *dhcpv6ServerLeaseState) HasInterfaceId() bool {
	return obj.obj.InterfaceId != nil
}

// The Interface ID option found in the last request message.
// SetInterfaceId sets the string value in the Dhcpv6ServerLeaseState object
func (obj *dhcpv6ServerLeaseState) SetInterfaceId(value string) Dhcpv6ServerLeaseState {

	obj.obj.InterfaceId = &value
	return obj
}

func (obj *dhcpv6ServerLeaseState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *dhcpv6ServerLeaseState) setDefault() {

}
