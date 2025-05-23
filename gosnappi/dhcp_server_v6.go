package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DhcpServerV6 *****
type dhcpServerV6 struct {
	validation
	obj           *otg.DhcpServerV6
	marshaller    marshalDhcpServerV6
	unMarshaller  unMarshalDhcpServerV6
	leasesHolder  DhcpServerV6DhcpV6ServerLeaseIter
	optionsHolder Dhcpv6ServerOptions
}

func NewDhcpServerV6() DhcpServerV6 {
	obj := dhcpServerV6{obj: &otg.DhcpServerV6{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpServerV6) msg() *otg.DhcpServerV6 {
	return obj.obj
}

func (obj *dhcpServerV6) setMsg(msg *otg.DhcpServerV6) DhcpServerV6 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpServerV6 struct {
	obj *dhcpServerV6
}

type marshalDhcpServerV6 interface {
	// ToProto marshals DhcpServerV6 to protobuf object *otg.DhcpServerV6
	ToProto() (*otg.DhcpServerV6, error)
	// ToPbText marshals DhcpServerV6 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DhcpServerV6 to YAML text
	ToYaml() (string, error)
	// ToJson marshals DhcpServerV6 to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals DhcpServerV6 to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldhcpServerV6 struct {
	obj *dhcpServerV6
}

type unMarshalDhcpServerV6 interface {
	// FromProto unmarshals DhcpServerV6 from protobuf object *otg.DhcpServerV6
	FromProto(msg *otg.DhcpServerV6) (DhcpServerV6, error)
	// FromPbText unmarshals DhcpServerV6 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DhcpServerV6 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DhcpServerV6 from JSON text
	FromJson(value string) error
}

func (obj *dhcpServerV6) Marshal() marshalDhcpServerV6 {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpServerV6{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpServerV6) Unmarshal() unMarshalDhcpServerV6 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpServerV6{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpServerV6) ToProto() (*otg.DhcpServerV6, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpServerV6) FromProto(msg *otg.DhcpServerV6) (DhcpServerV6, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpServerV6) ToPbText() (string, error) {
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

func (m *unMarshaldhcpServerV6) FromPbText(value string) error {
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

func (m *marshaldhcpServerV6) ToYaml() (string, error) {
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

func (m *unMarshaldhcpServerV6) FromYaml(value string) error {
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

func (m *marshaldhcpServerV6) ToJsonRaw() (string, error) {
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

func (m *marshaldhcpServerV6) ToJson() (string, error) {
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

func (m *unMarshaldhcpServerV6) FromJson(value string) error {
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

func (obj *dhcpServerV6) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpServerV6) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpServerV6) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpServerV6) Clone() (DhcpServerV6, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpServerV6()
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

func (obj *dhcpServerV6) setNil() {
	obj.leasesHolder = nil
	obj.optionsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DhcpServerV6 is configuration for emulated DHCPv6 Server.
type DhcpServerV6 interface {
	Validation
	// msg marshals DhcpServerV6 to protobuf object *otg.DhcpServerV6
	// and doesn't set defaults
	msg() *otg.DhcpServerV6
	// setMsg unmarshals DhcpServerV6 from protobuf object *otg.DhcpServerV6
	// and doesn't set defaults
	setMsg(*otg.DhcpServerV6) DhcpServerV6
	// provides marshal interface
	Marshal() marshalDhcpServerV6
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpServerV6
	// validate validates DhcpServerV6
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DhcpServerV6, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in DhcpServerV6.
	Name() string
	// SetName assigns string provided by user to DhcpServerV6
	SetName(value string) DhcpServerV6
	// Ipv6Name returns string, set in DhcpServerV6.
	Ipv6Name() string
	// SetIpv6Name assigns string provided by user to DhcpServerV6
	SetIpv6Name(value string) DhcpServerV6
	// RapidCommit returns bool, set in DhcpServerV6.
	RapidCommit() bool
	// SetRapidCommit assigns bool provided by user to DhcpServerV6
	SetRapidCommit(value bool) DhcpServerV6
	// HasRapidCommit checks if RapidCommit has been set in DhcpServerV6
	HasRapidCommit() bool
	// ReconfigureViaRelayAgent returns bool, set in DhcpServerV6.
	ReconfigureViaRelayAgent() bool
	// SetReconfigureViaRelayAgent assigns bool provided by user to DhcpServerV6
	SetReconfigureViaRelayAgent(value bool) DhcpServerV6
	// HasReconfigureViaRelayAgent checks if ReconfigureViaRelayAgent has been set in DhcpServerV6
	HasReconfigureViaRelayAgent() bool
	// Leases returns DhcpServerV6DhcpV6ServerLeaseIterIter, set in DhcpServerV6
	Leases() DhcpServerV6DhcpV6ServerLeaseIter
	// Options returns Dhcpv6ServerOptions, set in DhcpServerV6.
	// Dhcpv6ServerOptions is dHCP server options, these configured options are sent in Dhcp server messages.
	Options() Dhcpv6ServerOptions
	// SetOptions assigns Dhcpv6ServerOptions provided by user to DhcpServerV6.
	// Dhcpv6ServerOptions is dHCP server options, these configured options are sent in Dhcp server messages.
	SetOptions(value Dhcpv6ServerOptions) DhcpServerV6
	// HasOptions checks if Options has been set in DhcpServerV6
	HasOptions() bool
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *dhcpServerV6) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the DhcpServerV6 object
func (obj *dhcpServerV6) SetName(value string) DhcpServerV6 {

	obj.obj.Name = &value
	return obj
}

// The unique name of the IPv6 on which DHCPv6 server will run.
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
//
// Ipv6Name returns a string
func (obj *dhcpServerV6) Ipv6Name() string {

	return *obj.obj.Ipv6Name

}

// The unique name of the IPv6 on which DHCPv6 server will run.
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
//
// SetIpv6Name sets the string value in the DhcpServerV6 object
func (obj *dhcpServerV6) SetIpv6Name(value string) DhcpServerV6 {

	obj.obj.Ipv6Name = &value
	return obj
}

// If Rapid Commit is set, server responds to client initiated Rapid Commit two-message exchanges.
// RapidCommit returns a bool
func (obj *dhcpServerV6) RapidCommit() bool {

	return *obj.obj.RapidCommit

}

// If Rapid Commit is set, server responds to client initiated Rapid Commit two-message exchanges.
// RapidCommit returns a bool
func (obj *dhcpServerV6) HasRapidCommit() bool {
	return obj.obj.RapidCommit != nil
}

// If Rapid Commit is set, server responds to client initiated Rapid Commit two-message exchanges.
// SetRapidCommit sets the bool value in the DhcpServerV6 object
func (obj *dhcpServerV6) SetRapidCommit(value bool) DhcpServerV6 {

	obj.obj.RapidCommit = &value
	return obj
}

// If the server does not have an address to which it can send the Reconfigure message directly to the client, the  server uses a Relay-reply message to send the Reconfigure message to a relay agent that will relay the message to the client.
// ReconfigureViaRelayAgent returns a bool
func (obj *dhcpServerV6) ReconfigureViaRelayAgent() bool {

	return *obj.obj.ReconfigureViaRelayAgent

}

// If the server does not have an address to which it can send the Reconfigure message directly to the client, the  server uses a Relay-reply message to send the Reconfigure message to a relay agent that will relay the message to the client.
// ReconfigureViaRelayAgent returns a bool
func (obj *dhcpServerV6) HasReconfigureViaRelayAgent() bool {
	return obj.obj.ReconfigureViaRelayAgent != nil
}

// If the server does not have an address to which it can send the Reconfigure message directly to the client, the  server uses a Relay-reply message to send the Reconfigure message to a relay agent that will relay the message to the client.
// SetReconfigureViaRelayAgent sets the bool value in the DhcpServerV6 object
func (obj *dhcpServerV6) SetReconfigureViaRelayAgent(value bool) DhcpServerV6 {

	obj.obj.ReconfigureViaRelayAgent = &value
	return obj
}

// Array of DHCP pools configured on a server.
// Leases returns a []DhcpV6ServerLease
func (obj *dhcpServerV6) Leases() DhcpServerV6DhcpV6ServerLeaseIter {
	if len(obj.obj.Leases) == 0 {
		obj.obj.Leases = []*otg.DhcpV6ServerLease{}
	}
	if obj.leasesHolder == nil {
		obj.leasesHolder = newDhcpServerV6DhcpV6ServerLeaseIter(&obj.obj.Leases).setMsg(obj)
	}
	return obj.leasesHolder
}

type dhcpServerV6DhcpV6ServerLeaseIter struct {
	obj                    *dhcpServerV6
	dhcpV6ServerLeaseSlice []DhcpV6ServerLease
	fieldPtr               *[]*otg.DhcpV6ServerLease
}

func newDhcpServerV6DhcpV6ServerLeaseIter(ptr *[]*otg.DhcpV6ServerLease) DhcpServerV6DhcpV6ServerLeaseIter {
	return &dhcpServerV6DhcpV6ServerLeaseIter{fieldPtr: ptr}
}

type DhcpServerV6DhcpV6ServerLeaseIter interface {
	setMsg(*dhcpServerV6) DhcpServerV6DhcpV6ServerLeaseIter
	Items() []DhcpV6ServerLease
	Add() DhcpV6ServerLease
	Append(items ...DhcpV6ServerLease) DhcpServerV6DhcpV6ServerLeaseIter
	Set(index int, newObj DhcpV6ServerLease) DhcpServerV6DhcpV6ServerLeaseIter
	Clear() DhcpServerV6DhcpV6ServerLeaseIter
	clearHolderSlice() DhcpServerV6DhcpV6ServerLeaseIter
	appendHolderSlice(item DhcpV6ServerLease) DhcpServerV6DhcpV6ServerLeaseIter
}

func (obj *dhcpServerV6DhcpV6ServerLeaseIter) setMsg(msg *dhcpServerV6) DhcpServerV6DhcpV6ServerLeaseIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&dhcpV6ServerLease{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *dhcpServerV6DhcpV6ServerLeaseIter) Items() []DhcpV6ServerLease {
	return obj.dhcpV6ServerLeaseSlice
}

func (obj *dhcpServerV6DhcpV6ServerLeaseIter) Add() DhcpV6ServerLease {
	newObj := &otg.DhcpV6ServerLease{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &dhcpV6ServerLease{obj: newObj}
	newLibObj.setDefault()
	obj.dhcpV6ServerLeaseSlice = append(obj.dhcpV6ServerLeaseSlice, newLibObj)
	return newLibObj
}

func (obj *dhcpServerV6DhcpV6ServerLeaseIter) Append(items ...DhcpV6ServerLease) DhcpServerV6DhcpV6ServerLeaseIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.dhcpV6ServerLeaseSlice = append(obj.dhcpV6ServerLeaseSlice, item)
	}
	return obj
}

func (obj *dhcpServerV6DhcpV6ServerLeaseIter) Set(index int, newObj DhcpV6ServerLease) DhcpServerV6DhcpV6ServerLeaseIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.dhcpV6ServerLeaseSlice[index] = newObj
	return obj
}
func (obj *dhcpServerV6DhcpV6ServerLeaseIter) Clear() DhcpServerV6DhcpV6ServerLeaseIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.DhcpV6ServerLease{}
		obj.dhcpV6ServerLeaseSlice = []DhcpV6ServerLease{}
	}
	return obj
}
func (obj *dhcpServerV6DhcpV6ServerLeaseIter) clearHolderSlice() DhcpServerV6DhcpV6ServerLeaseIter {
	if len(obj.dhcpV6ServerLeaseSlice) > 0 {
		obj.dhcpV6ServerLeaseSlice = []DhcpV6ServerLease{}
	}
	return obj
}
func (obj *dhcpServerV6DhcpV6ServerLeaseIter) appendHolderSlice(item DhcpV6ServerLease) DhcpServerV6DhcpV6ServerLeaseIter {
	obj.dhcpV6ServerLeaseSlice = append(obj.dhcpV6ServerLeaseSlice, item)
	return obj
}

// Optional DHCPv4 Server options that are sent in Dhcp server messages.
// Options returns a Dhcpv6ServerOptions
func (obj *dhcpServerV6) Options() Dhcpv6ServerOptions {
	if obj.obj.Options == nil {
		obj.obj.Options = NewDhcpv6ServerOptions().msg()
	}
	if obj.optionsHolder == nil {
		obj.optionsHolder = &dhcpv6ServerOptions{obj: obj.obj.Options}
	}
	return obj.optionsHolder
}

// Optional DHCPv4 Server options that are sent in Dhcp server messages.
// Options returns a Dhcpv6ServerOptions
func (obj *dhcpServerV6) HasOptions() bool {
	return obj.obj.Options != nil
}

// Optional DHCPv4 Server options that are sent in Dhcp server messages.
// SetOptions sets the Dhcpv6ServerOptions value in the DhcpServerV6 object
func (obj *dhcpServerV6) SetOptions(value Dhcpv6ServerOptions) DhcpServerV6 {

	obj.optionsHolder = nil
	obj.obj.Options = value.msg()

	return obj
}

func (obj *dhcpServerV6) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface DhcpServerV6")
	}

	// Ipv6Name is required
	if obj.obj.Ipv6Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Ipv6Name is required field on interface DhcpServerV6")
	}

	if len(obj.obj.Leases) != 0 {

		if set_default {
			obj.Leases().clearHolderSlice()
			for _, item := range obj.obj.Leases {
				obj.Leases().appendHolderSlice(&dhcpV6ServerLease{obj: item})
			}
		}
		for _, item := range obj.Leases().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.Options != nil {

		obj.Options().validateObj(vObj, set_default)
	}

}

func (obj *dhcpServerV6) setDefault() {
	if obj.obj.RapidCommit == nil {
		obj.SetRapidCommit(false)
	}
	if obj.obj.ReconfigureViaRelayAgent == nil {
		obj.SetReconfigureViaRelayAgent(false)
	}

}
