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
	obj             *otg.Dhcpv6InterfaceState
	marshaller      marshalDhcpv6InterfaceState
	unMarshaller    unMarshalDhcpv6InterfaceState
	iapdsHolder     Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter
	addressesHolder Dhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter
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
	obj.setNil()
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
	m.obj.setNil()
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
	m.obj.setNil()
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
	m.obj.setNil()
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

func (obj *dhcpv6InterfaceState) setNil() {
	obj.iapdsHolder = nil
	obj.addressesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
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
	// Iapds returns Dhcpv6InterfaceStateDhcpv6InterfaceIapdIterIter, set in Dhcpv6InterfaceState
	Iapds() Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter
	// Addresses returns Dhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIterIter, set in Dhcpv6InterfaceState
	Addresses() Dhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter
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
	setNil()
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

// The IPv6 IAPD addresses and prefixes associated with this DHCP Client session.
// Iapds returns a []Dhcpv6InterfaceIapd
func (obj *dhcpv6InterfaceState) Iapds() Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter {
	if len(obj.obj.Iapds) == 0 {
		obj.obj.Iapds = []*otg.Dhcpv6InterfaceIapd{}
	}
	if obj.iapdsHolder == nil {
		obj.iapdsHolder = newDhcpv6InterfaceStateDhcpv6InterfaceIapdIter(&obj.obj.Iapds).setMsg(obj)
	}
	return obj.iapdsHolder
}

type dhcpv6InterfaceStateDhcpv6InterfaceIapdIter struct {
	obj                      *dhcpv6InterfaceState
	dhcpv6InterfaceIapdSlice []Dhcpv6InterfaceIapd
	fieldPtr                 *[]*otg.Dhcpv6InterfaceIapd
}

func newDhcpv6InterfaceStateDhcpv6InterfaceIapdIter(ptr *[]*otg.Dhcpv6InterfaceIapd) Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter {
	return &dhcpv6InterfaceStateDhcpv6InterfaceIapdIter{fieldPtr: ptr}
}

type Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter interface {
	setMsg(*dhcpv6InterfaceState) Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter
	Items() []Dhcpv6InterfaceIapd
	Add() Dhcpv6InterfaceIapd
	Append(items ...Dhcpv6InterfaceIapd) Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter
	Set(index int, newObj Dhcpv6InterfaceIapd) Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter
	Clear() Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter
	clearHolderSlice() Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter
	appendHolderSlice(item Dhcpv6InterfaceIapd) Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter
}

func (obj *dhcpv6InterfaceStateDhcpv6InterfaceIapdIter) setMsg(msg *dhcpv6InterfaceState) Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&dhcpv6InterfaceIapd{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *dhcpv6InterfaceStateDhcpv6InterfaceIapdIter) Items() []Dhcpv6InterfaceIapd {
	return obj.dhcpv6InterfaceIapdSlice
}

func (obj *dhcpv6InterfaceStateDhcpv6InterfaceIapdIter) Add() Dhcpv6InterfaceIapd {
	newObj := &otg.Dhcpv6InterfaceIapd{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &dhcpv6InterfaceIapd{obj: newObj}
	newLibObj.setDefault()
	obj.dhcpv6InterfaceIapdSlice = append(obj.dhcpv6InterfaceIapdSlice, newLibObj)
	return newLibObj
}

func (obj *dhcpv6InterfaceStateDhcpv6InterfaceIapdIter) Append(items ...Dhcpv6InterfaceIapd) Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.dhcpv6InterfaceIapdSlice = append(obj.dhcpv6InterfaceIapdSlice, item)
	}
	return obj
}

func (obj *dhcpv6InterfaceStateDhcpv6InterfaceIapdIter) Set(index int, newObj Dhcpv6InterfaceIapd) Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.dhcpv6InterfaceIapdSlice[index] = newObj
	return obj
}
func (obj *dhcpv6InterfaceStateDhcpv6InterfaceIapdIter) Clear() Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Dhcpv6InterfaceIapd{}
		obj.dhcpv6InterfaceIapdSlice = []Dhcpv6InterfaceIapd{}
	}
	return obj
}
func (obj *dhcpv6InterfaceStateDhcpv6InterfaceIapdIter) clearHolderSlice() Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter {
	if len(obj.dhcpv6InterfaceIapdSlice) > 0 {
		obj.dhcpv6InterfaceIapdSlice = []Dhcpv6InterfaceIapd{}
	}
	return obj
}
func (obj *dhcpv6InterfaceStateDhcpv6InterfaceIapdIter) appendHolderSlice(item Dhcpv6InterfaceIapd) Dhcpv6InterfaceStateDhcpv6InterfaceIapdIter {
	obj.dhcpv6InterfaceIapdSlice = append(obj.dhcpv6InterfaceIapdSlice, item)
	return obj
}

// The IPv6 addresses and gateways associated with this DHCP Client session.
// Addresses returns a []Dhcpv6InterfaceAddressInfo
func (obj *dhcpv6InterfaceState) Addresses() Dhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter {
	if len(obj.obj.Addresses) == 0 {
		obj.obj.Addresses = []*otg.Dhcpv6InterfaceAddressInfo{}
	}
	if obj.addressesHolder == nil {
		obj.addressesHolder = newDhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter(&obj.obj.Addresses).setMsg(obj)
	}
	return obj.addressesHolder
}

type dhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter struct {
	obj                             *dhcpv6InterfaceState
	dhcpv6InterfaceAddressInfoSlice []Dhcpv6InterfaceAddressInfo
	fieldPtr                        *[]*otg.Dhcpv6InterfaceAddressInfo
}

func newDhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter(ptr *[]*otg.Dhcpv6InterfaceAddressInfo) Dhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter {
	return &dhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter{fieldPtr: ptr}
}

type Dhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter interface {
	setMsg(*dhcpv6InterfaceState) Dhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter
	Items() []Dhcpv6InterfaceAddressInfo
	Add() Dhcpv6InterfaceAddressInfo
	Append(items ...Dhcpv6InterfaceAddressInfo) Dhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter
	Set(index int, newObj Dhcpv6InterfaceAddressInfo) Dhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter
	Clear() Dhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter
	clearHolderSlice() Dhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter
	appendHolderSlice(item Dhcpv6InterfaceAddressInfo) Dhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter
}

func (obj *dhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter) setMsg(msg *dhcpv6InterfaceState) Dhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&dhcpv6InterfaceAddressInfo{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *dhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter) Items() []Dhcpv6InterfaceAddressInfo {
	return obj.dhcpv6InterfaceAddressInfoSlice
}

func (obj *dhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter) Add() Dhcpv6InterfaceAddressInfo {
	newObj := &otg.Dhcpv6InterfaceAddressInfo{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &dhcpv6InterfaceAddressInfo{obj: newObj}
	newLibObj.setDefault()
	obj.dhcpv6InterfaceAddressInfoSlice = append(obj.dhcpv6InterfaceAddressInfoSlice, newLibObj)
	return newLibObj
}

func (obj *dhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter) Append(items ...Dhcpv6InterfaceAddressInfo) Dhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.dhcpv6InterfaceAddressInfoSlice = append(obj.dhcpv6InterfaceAddressInfoSlice, item)
	}
	return obj
}

func (obj *dhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter) Set(index int, newObj Dhcpv6InterfaceAddressInfo) Dhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.dhcpv6InterfaceAddressInfoSlice[index] = newObj
	return obj
}
func (obj *dhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter) Clear() Dhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Dhcpv6InterfaceAddressInfo{}
		obj.dhcpv6InterfaceAddressInfoSlice = []Dhcpv6InterfaceAddressInfo{}
	}
	return obj
}
func (obj *dhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter) clearHolderSlice() Dhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter {
	if len(obj.dhcpv6InterfaceAddressInfoSlice) > 0 {
		obj.dhcpv6InterfaceAddressInfoSlice = []Dhcpv6InterfaceAddressInfo{}
	}
	return obj
}
func (obj *dhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter) appendHolderSlice(item Dhcpv6InterfaceAddressInfo) Dhcpv6InterfaceStateDhcpv6InterfaceAddressInfoIter {
	obj.dhcpv6InterfaceAddressInfoSlice = append(obj.dhcpv6InterfaceAddressInfoSlice, item)
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

	if len(obj.obj.Iapds) != 0 {

		if set_default {
			obj.Iapds().clearHolderSlice()
			for _, item := range obj.obj.Iapds {
				obj.Iapds().appendHolderSlice(&dhcpv6InterfaceIapd{obj: item})
			}
		}
		for _, item := range obj.Iapds().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Addresses) != 0 {

		if set_default {
			obj.Addresses().clearHolderSlice()
			for _, item := range obj.obj.Addresses {
				obj.Addresses().appendHolderSlice(&dhcpv6InterfaceAddressInfo{obj: item})
			}
		}
		for _, item := range obj.Addresses().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *dhcpv6InterfaceState) setDefault() {

}
