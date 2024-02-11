package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceEthernet *****
type deviceEthernet struct {
	validation
	obj                 *otg.DeviceEthernet
	marshaller          marshalDeviceEthernet
	unMarshaller        unMarshalDeviceEthernet
	connectionHolder    EthernetConnection
	ipv4AddressesHolder DeviceEthernetDeviceIpv4Iter
	ipv6AddressesHolder DeviceEthernetDeviceIpv6Iter
	vlansHolder         DeviceEthernetDeviceVlanIter
}

func NewDeviceEthernet() DeviceEthernet {
	obj := deviceEthernet{obj: &otg.DeviceEthernet{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceEthernet) msg() *otg.DeviceEthernet {
	return obj.obj
}

func (obj *deviceEthernet) setMsg(msg *otg.DeviceEthernet) DeviceEthernet {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceEthernet struct {
	obj *deviceEthernet
}

type marshalDeviceEthernet interface {
	// ToProto marshals DeviceEthernet to protobuf object *otg.DeviceEthernet
	ToProto() (*otg.DeviceEthernet, error)
	// ToPbText marshals DeviceEthernet to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceEthernet to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceEthernet to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceEthernet struct {
	obj *deviceEthernet
}

type unMarshalDeviceEthernet interface {
	// FromProto unmarshals DeviceEthernet from protobuf object *otg.DeviceEthernet
	FromProto(msg *otg.DeviceEthernet) (DeviceEthernet, error)
	// FromPbText unmarshals DeviceEthernet from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceEthernet from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceEthernet from JSON text
	FromJson(value string) error
}

func (obj *deviceEthernet) Marshal() marshalDeviceEthernet {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceEthernet{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceEthernet) Unmarshal() unMarshalDeviceEthernet {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceEthernet{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceEthernet) ToProto() (*otg.DeviceEthernet, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceEthernet) FromProto(msg *otg.DeviceEthernet) (DeviceEthernet, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceEthernet) ToPbText() (string, error) {
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

func (m *unMarshaldeviceEthernet) FromPbText(value string) error {
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

func (m *marshaldeviceEthernet) ToYaml() (string, error) {
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

func (m *unMarshaldeviceEthernet) FromYaml(value string) error {
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

func (m *marshaldeviceEthernet) ToJson() (string, error) {
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

func (m *unMarshaldeviceEthernet) FromJson(value string) error {
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

func (obj *deviceEthernet) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceEthernet) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceEthernet) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceEthernet) Clone() (DeviceEthernet, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceEthernet()
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

func (obj *deviceEthernet) setNil() {
	obj.connectionHolder = nil
	obj.ipv4AddressesHolder = nil
	obj.ipv6AddressesHolder = nil
	obj.vlansHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceEthernet is an Ethernet interface with IPv4 and IPv6 addresses.
type DeviceEthernet interface {
	Validation
	// msg marshals DeviceEthernet to protobuf object *otg.DeviceEthernet
	// and doesn't set defaults
	msg() *otg.DeviceEthernet
	// setMsg unmarshals DeviceEthernet from protobuf object *otg.DeviceEthernet
	// and doesn't set defaults
	setMsg(*otg.DeviceEthernet) DeviceEthernet
	// provides marshal interface
	Marshal() marshalDeviceEthernet
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceEthernet
	// validate validates DeviceEthernet
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceEthernet, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Connection returns EthernetConnection, set in DeviceEthernet.
	// EthernetConnection is ethernet interface connection to a port, LAG or VXLAN tunnel.
	Connection() EthernetConnection
	// SetConnection assigns EthernetConnection provided by user to DeviceEthernet.
	// EthernetConnection is ethernet interface connection to a port, LAG or VXLAN tunnel.
	SetConnection(value EthernetConnection) DeviceEthernet
	// HasConnection checks if Connection has been set in DeviceEthernet
	HasConnection() bool
	// Ipv4Addresses returns DeviceEthernetDeviceIpv4IterIter, set in DeviceEthernet
	Ipv4Addresses() DeviceEthernetDeviceIpv4Iter
	// Ipv6Addresses returns DeviceEthernetDeviceIpv6IterIter, set in DeviceEthernet
	Ipv6Addresses() DeviceEthernetDeviceIpv6Iter
	// Mac returns string, set in DeviceEthernet.
	Mac() string
	// SetMac assigns string provided by user to DeviceEthernet
	SetMac(value string) DeviceEthernet
	// Mtu returns uint32, set in DeviceEthernet.
	Mtu() uint32
	// SetMtu assigns uint32 provided by user to DeviceEthernet
	SetMtu(value uint32) DeviceEthernet
	// HasMtu checks if Mtu has been set in DeviceEthernet
	HasMtu() bool
	// Vlans returns DeviceEthernetDeviceVlanIterIter, set in DeviceEthernet
	Vlans() DeviceEthernetDeviceVlanIter
	// Name returns string, set in DeviceEthernet.
	Name() string
	// SetName assigns string provided by user to DeviceEthernet
	SetName(value string) DeviceEthernet
	setNil()
}

// Device connection to physical, LAG or another device.
// Connection returns a EthernetConnection
func (obj *deviceEthernet) Connection() EthernetConnection {
	if obj.obj.Connection == nil {
		obj.obj.Connection = NewEthernetConnection().msg()
	}
	if obj.connectionHolder == nil {
		obj.connectionHolder = &ethernetConnection{obj: obj.obj.Connection}
	}
	return obj.connectionHolder
}

// Device connection to physical, LAG or another device.
// Connection returns a EthernetConnection
func (obj *deviceEthernet) HasConnection() bool {
	return obj.obj.Connection != nil
}

// Device connection to physical, LAG or another device.
// SetConnection sets the EthernetConnection value in the DeviceEthernet object
func (obj *deviceEthernet) SetConnection(value EthernetConnection) DeviceEthernet {

	obj.connectionHolder = nil
	obj.obj.Connection = value.msg()

	return obj
}

// List of IPv4 addresses and their gateways.
// Ipv4Addresses returns a []DeviceIpv4
func (obj *deviceEthernet) Ipv4Addresses() DeviceEthernetDeviceIpv4Iter {
	if len(obj.obj.Ipv4Addresses) == 0 {
		obj.obj.Ipv4Addresses = []*otg.DeviceIpv4{}
	}
	if obj.ipv4AddressesHolder == nil {
		obj.ipv4AddressesHolder = newDeviceEthernetDeviceIpv4Iter(&obj.obj.Ipv4Addresses).setMsg(obj)
	}
	return obj.ipv4AddressesHolder
}

type deviceEthernetDeviceIpv4Iter struct {
	obj             *deviceEthernet
	deviceIpv4Slice []DeviceIpv4
	fieldPtr        *[]*otg.DeviceIpv4
}

func newDeviceEthernetDeviceIpv4Iter(ptr *[]*otg.DeviceIpv4) DeviceEthernetDeviceIpv4Iter {
	return &deviceEthernetDeviceIpv4Iter{fieldPtr: ptr}
}

type DeviceEthernetDeviceIpv4Iter interface {
	setMsg(*deviceEthernet) DeviceEthernetDeviceIpv4Iter
	Items() []DeviceIpv4
	Add() DeviceIpv4
	Append(items ...DeviceIpv4) DeviceEthernetDeviceIpv4Iter
	Set(index int, newObj DeviceIpv4) DeviceEthernetDeviceIpv4Iter
	Clear() DeviceEthernetDeviceIpv4Iter
	clearHolderSlice() DeviceEthernetDeviceIpv4Iter
	appendHolderSlice(item DeviceIpv4) DeviceEthernetDeviceIpv4Iter
}

func (obj *deviceEthernetDeviceIpv4Iter) setMsg(msg *deviceEthernet) DeviceEthernetDeviceIpv4Iter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&deviceIpv4{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceEthernetDeviceIpv4Iter) Items() []DeviceIpv4 {
	return obj.deviceIpv4Slice
}

func (obj *deviceEthernetDeviceIpv4Iter) Add() DeviceIpv4 {
	newObj := &otg.DeviceIpv4{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &deviceIpv4{obj: newObj}
	newLibObj.setDefault()
	obj.deviceIpv4Slice = append(obj.deviceIpv4Slice, newLibObj)
	return newLibObj
}

func (obj *deviceEthernetDeviceIpv4Iter) Append(items ...DeviceIpv4) DeviceEthernetDeviceIpv4Iter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.deviceIpv4Slice = append(obj.deviceIpv4Slice, item)
	}
	return obj
}

func (obj *deviceEthernetDeviceIpv4Iter) Set(index int, newObj DeviceIpv4) DeviceEthernetDeviceIpv4Iter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.deviceIpv4Slice[index] = newObj
	return obj
}
func (obj *deviceEthernetDeviceIpv4Iter) Clear() DeviceEthernetDeviceIpv4Iter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.DeviceIpv4{}
		obj.deviceIpv4Slice = []DeviceIpv4{}
	}
	return obj
}
func (obj *deviceEthernetDeviceIpv4Iter) clearHolderSlice() DeviceEthernetDeviceIpv4Iter {
	if len(obj.deviceIpv4Slice) > 0 {
		obj.deviceIpv4Slice = []DeviceIpv4{}
	}
	return obj
}
func (obj *deviceEthernetDeviceIpv4Iter) appendHolderSlice(item DeviceIpv4) DeviceEthernetDeviceIpv4Iter {
	obj.deviceIpv4Slice = append(obj.deviceIpv4Slice, item)
	return obj
}

// List of global IPv6 addresses and their gateways.
// The Link Local IPv6 address will be automatically generated.
// Ipv6Addresses returns a []DeviceIpv6
func (obj *deviceEthernet) Ipv6Addresses() DeviceEthernetDeviceIpv6Iter {
	if len(obj.obj.Ipv6Addresses) == 0 {
		obj.obj.Ipv6Addresses = []*otg.DeviceIpv6{}
	}
	if obj.ipv6AddressesHolder == nil {
		obj.ipv6AddressesHolder = newDeviceEthernetDeviceIpv6Iter(&obj.obj.Ipv6Addresses).setMsg(obj)
	}
	return obj.ipv6AddressesHolder
}

type deviceEthernetDeviceIpv6Iter struct {
	obj             *deviceEthernet
	deviceIpv6Slice []DeviceIpv6
	fieldPtr        *[]*otg.DeviceIpv6
}

func newDeviceEthernetDeviceIpv6Iter(ptr *[]*otg.DeviceIpv6) DeviceEthernetDeviceIpv6Iter {
	return &deviceEthernetDeviceIpv6Iter{fieldPtr: ptr}
}

type DeviceEthernetDeviceIpv6Iter interface {
	setMsg(*deviceEthernet) DeviceEthernetDeviceIpv6Iter
	Items() []DeviceIpv6
	Add() DeviceIpv6
	Append(items ...DeviceIpv6) DeviceEthernetDeviceIpv6Iter
	Set(index int, newObj DeviceIpv6) DeviceEthernetDeviceIpv6Iter
	Clear() DeviceEthernetDeviceIpv6Iter
	clearHolderSlice() DeviceEthernetDeviceIpv6Iter
	appendHolderSlice(item DeviceIpv6) DeviceEthernetDeviceIpv6Iter
}

func (obj *deviceEthernetDeviceIpv6Iter) setMsg(msg *deviceEthernet) DeviceEthernetDeviceIpv6Iter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&deviceIpv6{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceEthernetDeviceIpv6Iter) Items() []DeviceIpv6 {
	return obj.deviceIpv6Slice
}

func (obj *deviceEthernetDeviceIpv6Iter) Add() DeviceIpv6 {
	newObj := &otg.DeviceIpv6{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &deviceIpv6{obj: newObj}
	newLibObj.setDefault()
	obj.deviceIpv6Slice = append(obj.deviceIpv6Slice, newLibObj)
	return newLibObj
}

func (obj *deviceEthernetDeviceIpv6Iter) Append(items ...DeviceIpv6) DeviceEthernetDeviceIpv6Iter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.deviceIpv6Slice = append(obj.deviceIpv6Slice, item)
	}
	return obj
}

func (obj *deviceEthernetDeviceIpv6Iter) Set(index int, newObj DeviceIpv6) DeviceEthernetDeviceIpv6Iter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.deviceIpv6Slice[index] = newObj
	return obj
}
func (obj *deviceEthernetDeviceIpv6Iter) Clear() DeviceEthernetDeviceIpv6Iter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.DeviceIpv6{}
		obj.deviceIpv6Slice = []DeviceIpv6{}
	}
	return obj
}
func (obj *deviceEthernetDeviceIpv6Iter) clearHolderSlice() DeviceEthernetDeviceIpv6Iter {
	if len(obj.deviceIpv6Slice) > 0 {
		obj.deviceIpv6Slice = []DeviceIpv6{}
	}
	return obj
}
func (obj *deviceEthernetDeviceIpv6Iter) appendHolderSlice(item DeviceIpv6) DeviceEthernetDeviceIpv6Iter {
	obj.deviceIpv6Slice = append(obj.deviceIpv6Slice, item)
	return obj
}

// Media Access Control address.
// Mac returns a string
func (obj *deviceEthernet) Mac() string {

	return *obj.obj.Mac

}

// Media Access Control address.
// SetMac sets the string value in the DeviceEthernet object
func (obj *deviceEthernet) SetMac(value string) DeviceEthernet {

	obj.obj.Mac = &value
	return obj
}

// Maximum Transmission Unit.
// Mtu returns a uint32
func (obj *deviceEthernet) Mtu() uint32 {

	return *obj.obj.Mtu

}

// Maximum Transmission Unit.
// Mtu returns a uint32
func (obj *deviceEthernet) HasMtu() bool {
	return obj.obj.Mtu != nil
}

// Maximum Transmission Unit.
// SetMtu sets the uint32 value in the DeviceEthernet object
func (obj *deviceEthernet) SetMtu(value uint32) DeviceEthernet {

	obj.obj.Mtu = &value
	return obj
}

// List of VLANs
// Vlans returns a []DeviceVlan
func (obj *deviceEthernet) Vlans() DeviceEthernetDeviceVlanIter {
	if len(obj.obj.Vlans) == 0 {
		obj.obj.Vlans = []*otg.DeviceVlan{}
	}
	if obj.vlansHolder == nil {
		obj.vlansHolder = newDeviceEthernetDeviceVlanIter(&obj.obj.Vlans).setMsg(obj)
	}
	return obj.vlansHolder
}

type deviceEthernetDeviceVlanIter struct {
	obj             *deviceEthernet
	deviceVlanSlice []DeviceVlan
	fieldPtr        *[]*otg.DeviceVlan
}

func newDeviceEthernetDeviceVlanIter(ptr *[]*otg.DeviceVlan) DeviceEthernetDeviceVlanIter {
	return &deviceEthernetDeviceVlanIter{fieldPtr: ptr}
}

type DeviceEthernetDeviceVlanIter interface {
	setMsg(*deviceEthernet) DeviceEthernetDeviceVlanIter
	Items() []DeviceVlan
	Add() DeviceVlan
	Append(items ...DeviceVlan) DeviceEthernetDeviceVlanIter
	Set(index int, newObj DeviceVlan) DeviceEthernetDeviceVlanIter
	Clear() DeviceEthernetDeviceVlanIter
	clearHolderSlice() DeviceEthernetDeviceVlanIter
	appendHolderSlice(item DeviceVlan) DeviceEthernetDeviceVlanIter
}

func (obj *deviceEthernetDeviceVlanIter) setMsg(msg *deviceEthernet) DeviceEthernetDeviceVlanIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&deviceVlan{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceEthernetDeviceVlanIter) Items() []DeviceVlan {
	return obj.deviceVlanSlice
}

func (obj *deviceEthernetDeviceVlanIter) Add() DeviceVlan {
	newObj := &otg.DeviceVlan{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &deviceVlan{obj: newObj}
	newLibObj.setDefault()
	obj.deviceVlanSlice = append(obj.deviceVlanSlice, newLibObj)
	return newLibObj
}

func (obj *deviceEthernetDeviceVlanIter) Append(items ...DeviceVlan) DeviceEthernetDeviceVlanIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.deviceVlanSlice = append(obj.deviceVlanSlice, item)
	}
	return obj
}

func (obj *deviceEthernetDeviceVlanIter) Set(index int, newObj DeviceVlan) DeviceEthernetDeviceVlanIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.deviceVlanSlice[index] = newObj
	return obj
}
func (obj *deviceEthernetDeviceVlanIter) Clear() DeviceEthernetDeviceVlanIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.DeviceVlan{}
		obj.deviceVlanSlice = []DeviceVlan{}
	}
	return obj
}
func (obj *deviceEthernetDeviceVlanIter) clearHolderSlice() DeviceEthernetDeviceVlanIter {
	if len(obj.deviceVlanSlice) > 0 {
		obj.deviceVlanSlice = []DeviceVlan{}
	}
	return obj
}
func (obj *deviceEthernetDeviceVlanIter) appendHolderSlice(item DeviceVlan) DeviceEthernetDeviceVlanIter {
	obj.deviceVlanSlice = append(obj.deviceVlanSlice, item)
	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *deviceEthernet) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the DeviceEthernet object
func (obj *deviceEthernet) SetName(value string) DeviceEthernet {

	obj.obj.Name = &value
	return obj
}

func (obj *deviceEthernet) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Connection != nil {

		obj.Connection().validateObj(vObj, set_default)
	}

	if len(obj.obj.Ipv4Addresses) != 0 {

		if set_default {
			obj.Ipv4Addresses().clearHolderSlice()
			for _, item := range obj.obj.Ipv4Addresses {
				obj.Ipv4Addresses().appendHolderSlice(&deviceIpv4{obj: item})
			}
		}
		for _, item := range obj.Ipv4Addresses().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Ipv6Addresses) != 0 {

		if set_default {
			obj.Ipv6Addresses().clearHolderSlice()
			for _, item := range obj.obj.Ipv6Addresses {
				obj.Ipv6Addresses().appendHolderSlice(&deviceIpv6{obj: item})
			}
		}
		for _, item := range obj.Ipv6Addresses().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	// Mac is required
	if obj.obj.Mac == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Mac is required field on interface DeviceEthernet")
	}
	if obj.obj.Mac != nil {

		err := obj.validateMac(obj.Mac())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on DeviceEthernet.Mac"))
		}

	}

	if obj.obj.Mtu != nil {

		if *obj.obj.Mtu > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= DeviceEthernet.Mtu <= 65535 but Got %d", *obj.obj.Mtu))
		}

	}

	if len(obj.obj.Vlans) != 0 {

		if set_default {
			obj.Vlans().clearHolderSlice()
			for _, item := range obj.obj.Vlans {
				obj.Vlans().appendHolderSlice(&deviceVlan{obj: item})
			}
		}
		for _, item := range obj.Vlans().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface DeviceEthernet")
	}
}

func (obj *deviceEthernet) setDefault() {
	if obj.obj.Mtu == nil {
		obj.SetMtu(1500)
	}

}
