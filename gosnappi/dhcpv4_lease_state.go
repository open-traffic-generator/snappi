package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv4LeaseState *****
type dhcpv4LeaseState struct {
	validation
	obj          *otg.Dhcpv4LeaseState
	marshaller   marshalDhcpv4LeaseState
	unMarshaller unMarshalDhcpv4LeaseState
}

func NewDhcpv4LeaseState() Dhcpv4LeaseState {
	obj := dhcpv4LeaseState{obj: &otg.Dhcpv4LeaseState{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv4LeaseState) msg() *otg.Dhcpv4LeaseState {
	return obj.obj
}

func (obj *dhcpv4LeaseState) setMsg(msg *otg.Dhcpv4LeaseState) Dhcpv4LeaseState {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv4LeaseState struct {
	obj *dhcpv4LeaseState
}

type marshalDhcpv4LeaseState interface {
	// ToProto marshals Dhcpv4LeaseState to protobuf object *otg.Dhcpv4LeaseState
	ToProto() (*otg.Dhcpv4LeaseState, error)
	// ToPbText marshals Dhcpv4LeaseState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv4LeaseState to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv4LeaseState to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Dhcpv4LeaseState to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldhcpv4LeaseState struct {
	obj *dhcpv4LeaseState
}

type unMarshalDhcpv4LeaseState interface {
	// FromProto unmarshals Dhcpv4LeaseState from protobuf object *otg.Dhcpv4LeaseState
	FromProto(msg *otg.Dhcpv4LeaseState) (Dhcpv4LeaseState, error)
	// FromPbText unmarshals Dhcpv4LeaseState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv4LeaseState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv4LeaseState from JSON text
	FromJson(value string) error
}

func (obj *dhcpv4LeaseState) Marshal() marshalDhcpv4LeaseState {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv4LeaseState{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv4LeaseState) Unmarshal() unMarshalDhcpv4LeaseState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv4LeaseState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv4LeaseState) ToProto() (*otg.Dhcpv4LeaseState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv4LeaseState) FromProto(msg *otg.Dhcpv4LeaseState) (Dhcpv4LeaseState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv4LeaseState) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv4LeaseState) FromPbText(value string) error {
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

func (m *marshaldhcpv4LeaseState) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv4LeaseState) FromYaml(value string) error {
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

func (m *marshaldhcpv4LeaseState) ToJsonRaw() (string, error) {
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

func (m *marshaldhcpv4LeaseState) ToJson() (string, error) {
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

func (m *unMarshaldhcpv4LeaseState) FromJson(value string) error {
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

func (obj *dhcpv4LeaseState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv4LeaseState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv4LeaseState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv4LeaseState) Clone() (Dhcpv4LeaseState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv4LeaseState()
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

// Dhcpv4LeaseState is iPv4 address lease state.
type Dhcpv4LeaseState interface {
	Validation
	// msg marshals Dhcpv4LeaseState to protobuf object *otg.Dhcpv4LeaseState
	// and doesn't set defaults
	msg() *otg.Dhcpv4LeaseState
	// setMsg unmarshals Dhcpv4LeaseState from protobuf object *otg.Dhcpv4LeaseState
	// and doesn't set defaults
	setMsg(*otg.Dhcpv4LeaseState) Dhcpv4LeaseState
	// provides marshal interface
	Marshal() marshalDhcpv4LeaseState
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv4LeaseState
	// validate validates Dhcpv4LeaseState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv4LeaseState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Address returns string, set in Dhcpv4LeaseState.
	Address() string
	// SetAddress assigns string provided by user to Dhcpv4LeaseState
	SetAddress(value string) Dhcpv4LeaseState
	// HasAddress checks if Address has been set in Dhcpv4LeaseState
	HasAddress() bool
	// ValidTime returns uint32, set in Dhcpv4LeaseState.
	ValidTime() uint32
	// SetValidTime assigns uint32 provided by user to Dhcpv4LeaseState
	SetValidTime(value uint32) Dhcpv4LeaseState
	// HasValidTime checks if ValidTime has been set in Dhcpv4LeaseState
	HasValidTime() bool
	// PreferredTime returns uint32, set in Dhcpv4LeaseState.
	PreferredTime() uint32
	// SetPreferredTime assigns uint32 provided by user to Dhcpv4LeaseState
	SetPreferredTime(value uint32) Dhcpv4LeaseState
	// HasPreferredTime checks if PreferredTime has been set in Dhcpv4LeaseState
	HasPreferredTime() bool
	// RenewTime returns uint32, set in Dhcpv4LeaseState.
	RenewTime() uint32
	// SetRenewTime assigns uint32 provided by user to Dhcpv4LeaseState
	SetRenewTime(value uint32) Dhcpv4LeaseState
	// HasRenewTime checks if RenewTime has been set in Dhcpv4LeaseState
	HasRenewTime() bool
	// RebindTime returns uint32, set in Dhcpv4LeaseState.
	RebindTime() uint32
	// SetRebindTime assigns uint32 provided by user to Dhcpv4LeaseState
	SetRebindTime(value uint32) Dhcpv4LeaseState
	// HasRebindTime checks if RebindTime has been set in Dhcpv4LeaseState
	HasRebindTime() bool
	// ClientId returns string, set in Dhcpv4LeaseState.
	ClientId() string
	// SetClientId assigns string provided by user to Dhcpv4LeaseState
	SetClientId(value string) Dhcpv4LeaseState
	// HasClientId checks if ClientId has been set in Dhcpv4LeaseState
	HasClientId() bool
	// CircuitId returns string, set in Dhcpv4LeaseState.
	CircuitId() string
	// SetCircuitId assigns string provided by user to Dhcpv4LeaseState
	SetCircuitId(value string) Dhcpv4LeaseState
	// HasCircuitId checks if CircuitId has been set in Dhcpv4LeaseState
	HasCircuitId() bool
	// RemoteId returns string, set in Dhcpv4LeaseState.
	RemoteId() string
	// SetRemoteId assigns string provided by user to Dhcpv4LeaseState
	SetRemoteId(value string) Dhcpv4LeaseState
	// HasRemoteId checks if RemoteId has been set in Dhcpv4LeaseState
	HasRemoteId() bool
}

// The IPv4 address associated with this lease.
// Address returns a string
func (obj *dhcpv4LeaseState) Address() string {

	return *obj.obj.Address

}

// The IPv4 address associated with this lease.
// Address returns a string
func (obj *dhcpv4LeaseState) HasAddress() bool {
	return obj.obj.Address != nil
}

// The IPv4 address associated with this lease.
// SetAddress sets the string value in the Dhcpv4LeaseState object
func (obj *dhcpv4LeaseState) SetAddress(value string) Dhcpv4LeaseState {

	obj.obj.Address = &value
	return obj
}

// The time in seconds after which the IPv4 address lease will expire.
// ValidTime returns a uint32
func (obj *dhcpv4LeaseState) ValidTime() uint32 {

	return *obj.obj.ValidTime

}

// The time in seconds after which the IPv4 address lease will expire.
// ValidTime returns a uint32
func (obj *dhcpv4LeaseState) HasValidTime() bool {
	return obj.obj.ValidTime != nil
}

// The time in seconds after which the IPv4 address lease will expire.
// SetValidTime sets the uint32 value in the Dhcpv4LeaseState object
func (obj *dhcpv4LeaseState) SetValidTime(value uint32) Dhcpv4LeaseState {

	obj.obj.ValidTime = &value
	return obj
}

// The elapsed time in seconds since the address has been renewed.
// PreferredTime returns a uint32
func (obj *dhcpv4LeaseState) PreferredTime() uint32 {

	return *obj.obj.PreferredTime

}

// The elapsed time in seconds since the address has been renewed.
// PreferredTime returns a uint32
func (obj *dhcpv4LeaseState) HasPreferredTime() bool {
	return obj.obj.PreferredTime != nil
}

// The elapsed time in seconds since the address has been renewed.
// SetPreferredTime sets the uint32 value in the Dhcpv4LeaseState object
func (obj *dhcpv4LeaseState) SetPreferredTime(value uint32) Dhcpv4LeaseState {

	obj.obj.PreferredTime = &value
	return obj
}

// Time in seconds until the DHCPv4 client starts renewing the lease.
// RenewTime returns a uint32
func (obj *dhcpv4LeaseState) RenewTime() uint32 {

	return *obj.obj.RenewTime

}

// Time in seconds until the DHCPv4 client starts renewing the lease.
// RenewTime returns a uint32
func (obj *dhcpv4LeaseState) HasRenewTime() bool {
	return obj.obj.RenewTime != nil
}

// Time in seconds until the DHCPv4 client starts renewing the lease.
// SetRenewTime sets the uint32 value in the Dhcpv4LeaseState object
func (obj *dhcpv4LeaseState) SetRenewTime(value uint32) Dhcpv4LeaseState {

	obj.obj.RenewTime = &value
	return obj
}

// Time in seconds until the DHCPv4 client starts rebinding.
// RebindTime returns a uint32
func (obj *dhcpv4LeaseState) RebindTime() uint32 {

	return *obj.obj.RebindTime

}

// Time in seconds until the DHCPv4 client starts rebinding.
// RebindTime returns a uint32
func (obj *dhcpv4LeaseState) HasRebindTime() bool {
	return obj.obj.RebindTime != nil
}

// Time in seconds until the DHCPv4 client starts rebinding.
// SetRebindTime sets the uint32 value in the Dhcpv4LeaseState object
func (obj *dhcpv4LeaseState) SetRebindTime(value uint32) Dhcpv4LeaseState {

	obj.obj.RebindTime = &value
	return obj
}

// The ID of the DHCPv4 client holding this lease.
// ClientId returns a string
func (obj *dhcpv4LeaseState) ClientId() string {

	return *obj.obj.ClientId

}

// The ID of the DHCPv4 client holding this lease.
// ClientId returns a string
func (obj *dhcpv4LeaseState) HasClientId() bool {
	return obj.obj.ClientId != nil
}

// The ID of the DHCPv4 client holding this lease.
// SetClientId sets the string value in the Dhcpv4LeaseState object
func (obj *dhcpv4LeaseState) SetClientId(value string) Dhcpv4LeaseState {

	obj.obj.ClientId = &value
	return obj
}

// The Circuit ID option found in the last request message.
// CircuitId returns a string
func (obj *dhcpv4LeaseState) CircuitId() string {

	return *obj.obj.CircuitId

}

// The Circuit ID option found in the last request message.
// CircuitId returns a string
func (obj *dhcpv4LeaseState) HasCircuitId() bool {
	return obj.obj.CircuitId != nil
}

// The Circuit ID option found in the last request message.
// SetCircuitId sets the string value in the Dhcpv4LeaseState object
func (obj *dhcpv4LeaseState) SetCircuitId(value string) Dhcpv4LeaseState {

	obj.obj.CircuitId = &value
	return obj
}

// The Remote ID option found in the last request message.
// RemoteId returns a string
func (obj *dhcpv4LeaseState) RemoteId() string {

	return *obj.obj.RemoteId

}

// The Remote ID option found in the last request message.
// RemoteId returns a string
func (obj *dhcpv4LeaseState) HasRemoteId() bool {
	return obj.obj.RemoteId != nil
}

// The Remote ID option found in the last request message.
// SetRemoteId sets the string value in the Dhcpv4LeaseState object
func (obj *dhcpv4LeaseState) SetRemoteId(value string) Dhcpv4LeaseState {

	obj.obj.RemoteId = &value
	return obj
}

func (obj *dhcpv4LeaseState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *dhcpv4LeaseState) setDefault() {

}
