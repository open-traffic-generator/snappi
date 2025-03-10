package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Device *****
type device struct {
	validation
	obj                 *otg.Device
	marshaller          marshalDevice
	unMarshaller        unMarshalDevice
	ethernetsHolder     DeviceDeviceEthernetIter
	ipv4LoopbacksHolder DeviceDeviceIpv4LoopbackIter
	ipv6LoopbacksHolder DeviceDeviceIpv6LoopbackIter
	isisHolder          DeviceIsisRouter
	bgpHolder           DeviceBgpRouter
	vxlanHolder         DeviceVxlan
	rsvpHolder          DeviceRsvp
	dhcpServerHolder    DeviceDhcpServer
	ospfv2Holder        DeviceOspfv2Router
	macsecHolder        DeviceMacsec
}

func NewDevice() Device {
	obj := device{obj: &otg.Device{}}
	obj.setDefault()
	return &obj
}

func (obj *device) msg() *otg.Device {
	return obj.obj
}

func (obj *device) setMsg(msg *otg.Device) Device {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldevice struct {
	obj *device
}

type marshalDevice interface {
	// ToProto marshals Device to protobuf object *otg.Device
	ToProto() (*otg.Device, error)
	// ToPbText marshals Device to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Device to YAML text
	ToYaml() (string, error)
	// ToJson marshals Device to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Device to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldevice struct {
	obj *device
}

type unMarshalDevice interface {
	// FromProto unmarshals Device from protobuf object *otg.Device
	FromProto(msg *otg.Device) (Device, error)
	// FromPbText unmarshals Device from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Device from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Device from JSON text
	FromJson(value string) error
}

func (obj *device) Marshal() marshalDevice {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldevice{obj: obj}
	}
	return obj.marshaller
}

func (obj *device) Unmarshal() unMarshalDevice {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldevice{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldevice) ToProto() (*otg.Device, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldevice) FromProto(msg *otg.Device) (Device, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldevice) ToPbText() (string, error) {
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

func (m *unMarshaldevice) FromPbText(value string) error {
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

func (m *marshaldevice) ToYaml() (string, error) {
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

func (m *unMarshaldevice) FromYaml(value string) error {
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

func (m *marshaldevice) ToJsonRaw() (string, error) {
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

func (m *marshaldevice) ToJson() (string, error) {
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

func (m *unMarshaldevice) FromJson(value string) error {
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

func (obj *device) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *device) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *device) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *device) Clone() (Device, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDevice()
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

func (obj *device) setNil() {
	obj.ethernetsHolder = nil
	obj.ipv4LoopbacksHolder = nil
	obj.ipv6LoopbacksHolder = nil
	obj.isisHolder = nil
	obj.bgpHolder = nil
	obj.vxlanHolder = nil
	obj.rsvpHolder = nil
	obj.dhcpServerHolder = nil
	obj.ospfv2Holder = nil
	obj.macsecHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Device is a container for emulated or simulated interfaces, loopback interfaces and protocol configurations.
type Device interface {
	Validation
	// msg marshals Device to protobuf object *otg.Device
	// and doesn't set defaults
	msg() *otg.Device
	// setMsg unmarshals Device from protobuf object *otg.Device
	// and doesn't set defaults
	setMsg(*otg.Device) Device
	// provides marshal interface
	Marshal() marshalDevice
	// provides unmarshal interface
	Unmarshal() unMarshalDevice
	// validate validates Device
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Device, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ethernets returns DeviceDeviceEthernetIterIter, set in Device
	Ethernets() DeviceDeviceEthernetIter
	// Ipv4Loopbacks returns DeviceDeviceIpv4LoopbackIterIter, set in Device
	Ipv4Loopbacks() DeviceDeviceIpv4LoopbackIter
	// Ipv6Loopbacks returns DeviceDeviceIpv6LoopbackIterIter, set in Device
	Ipv6Loopbacks() DeviceDeviceIpv6LoopbackIter
	// Isis returns DeviceIsisRouter, set in Device.
	// DeviceIsisRouter is a container of properties for an ISIS router and its interfaces.
	Isis() DeviceIsisRouter
	// SetIsis assigns DeviceIsisRouter provided by user to Device.
	// DeviceIsisRouter is a container of properties for an ISIS router and its interfaces.
	SetIsis(value DeviceIsisRouter) Device
	// HasIsis checks if Isis has been set in Device
	HasIsis() bool
	// Bgp returns DeviceBgpRouter, set in Device.
	// DeviceBgpRouter is configuration for one or more IPv4 or IPv6 BGP peers.
	Bgp() DeviceBgpRouter
	// SetBgp assigns DeviceBgpRouter provided by user to Device.
	// DeviceBgpRouter is configuration for one or more IPv4 or IPv6 BGP peers.
	SetBgp(value DeviceBgpRouter) Device
	// HasBgp checks if Bgp has been set in Device
	HasBgp() bool
	// Vxlan returns DeviceVxlan, set in Device.
	// DeviceVxlan is description is TBD
	Vxlan() DeviceVxlan
	// SetVxlan assigns DeviceVxlan provided by user to Device.
	// DeviceVxlan is description is TBD
	SetVxlan(value DeviceVxlan) Device
	// HasVxlan checks if Vxlan has been set in Device
	HasVxlan() bool
	// Name returns string, set in Device.
	Name() string
	// SetName assigns string provided by user to Device
	SetName(value string) Device
	// Rsvp returns DeviceRsvp, set in Device.
	// DeviceRsvp is configuration for one or more RSVP interfaces, ingress and egress LSPs. In this model, currently IPv4 RSVP and point-to-point LSPs are supported as per RFC3209 and related specifications.
	Rsvp() DeviceRsvp
	// SetRsvp assigns DeviceRsvp provided by user to Device.
	// DeviceRsvp is configuration for one or more RSVP interfaces, ingress and egress LSPs. In this model, currently IPv4 RSVP and point-to-point LSPs are supported as per RFC3209 and related specifications.
	SetRsvp(value DeviceRsvp) Device
	// HasRsvp checks if Rsvp has been set in Device
	HasRsvp() bool
	// DhcpServer returns DeviceDhcpServer, set in Device.
	// DeviceDhcpServer is configuration for one or more IPv4 or IPv6 DHCP servers.
	DhcpServer() DeviceDhcpServer
	// SetDhcpServer assigns DeviceDhcpServer provided by user to Device.
	// DeviceDhcpServer is configuration for one or more IPv4 or IPv6 DHCP servers.
	SetDhcpServer(value DeviceDhcpServer) Device
	// HasDhcpServer checks if DhcpServer has been set in Device
	HasDhcpServer() bool
	// Ospfv2 returns DeviceOspfv2Router, set in Device.
	// DeviceOspfv2Router is under Review: OSPFv2 is currently under review for pending exploration on use cases.
	//
	// Under Review: OSPFv2 is currently under review for pending exploration on use cases.
	//
	// A container of properties for an OSPFv2 router and its interfaces & Route Ranges.
	Ospfv2() DeviceOspfv2Router
	// SetOspfv2 assigns DeviceOspfv2Router provided by user to Device.
	// DeviceOspfv2Router is under Review: OSPFv2 is currently under review for pending exploration on use cases.
	//
	// Under Review: OSPFv2 is currently under review for pending exploration on use cases.
	//
	// A container of properties for an OSPFv2 router and its interfaces & Route Ranges.
	SetOspfv2(value DeviceOspfv2Router) Device
	// HasOspfv2 checks if Ospfv2 has been set in Device
	HasOspfv2() bool
	// Macsec returns DeviceMacsec, set in Device.
	// DeviceMacsec is a container of properties for a MACsec capable device. Reference https://1.ieee802.org/security/802-1ae/.
	Macsec() DeviceMacsec
	// SetMacsec assigns DeviceMacsec provided by user to Device.
	// DeviceMacsec is a container of properties for a MACsec capable device. Reference https://1.ieee802.org/security/802-1ae/.
	SetMacsec(value DeviceMacsec) Device
	// HasMacsec checks if Macsec has been set in Device
	HasMacsec() bool
	setNil()
}

// Ethernet configuration for one or more emulated or simulated network interfaces.
// Ethernets returns a []DeviceEthernet
func (obj *device) Ethernets() DeviceDeviceEthernetIter {
	if len(obj.obj.Ethernets) == 0 {
		obj.obj.Ethernets = []*otg.DeviceEthernet{}
	}
	if obj.ethernetsHolder == nil {
		obj.ethernetsHolder = newDeviceDeviceEthernetIter(&obj.obj.Ethernets).setMsg(obj)
	}
	return obj.ethernetsHolder
}

type deviceDeviceEthernetIter struct {
	obj                 *device
	deviceEthernetSlice []DeviceEthernet
	fieldPtr            *[]*otg.DeviceEthernet
}

func newDeviceDeviceEthernetIter(ptr *[]*otg.DeviceEthernet) DeviceDeviceEthernetIter {
	return &deviceDeviceEthernetIter{fieldPtr: ptr}
}

type DeviceDeviceEthernetIter interface {
	setMsg(*device) DeviceDeviceEthernetIter
	Items() []DeviceEthernet
	Add() DeviceEthernet
	Append(items ...DeviceEthernet) DeviceDeviceEthernetIter
	Set(index int, newObj DeviceEthernet) DeviceDeviceEthernetIter
	Clear() DeviceDeviceEthernetIter
	clearHolderSlice() DeviceDeviceEthernetIter
	appendHolderSlice(item DeviceEthernet) DeviceDeviceEthernetIter
}

func (obj *deviceDeviceEthernetIter) setMsg(msg *device) DeviceDeviceEthernetIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&deviceEthernet{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceDeviceEthernetIter) Items() []DeviceEthernet {
	return obj.deviceEthernetSlice
}

func (obj *deviceDeviceEthernetIter) Add() DeviceEthernet {
	newObj := &otg.DeviceEthernet{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &deviceEthernet{obj: newObj}
	newLibObj.setDefault()
	obj.deviceEthernetSlice = append(obj.deviceEthernetSlice, newLibObj)
	return newLibObj
}

func (obj *deviceDeviceEthernetIter) Append(items ...DeviceEthernet) DeviceDeviceEthernetIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.deviceEthernetSlice = append(obj.deviceEthernetSlice, item)
	}
	return obj
}

func (obj *deviceDeviceEthernetIter) Set(index int, newObj DeviceEthernet) DeviceDeviceEthernetIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.deviceEthernetSlice[index] = newObj
	return obj
}
func (obj *deviceDeviceEthernetIter) Clear() DeviceDeviceEthernetIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.DeviceEthernet{}
		obj.deviceEthernetSlice = []DeviceEthernet{}
	}
	return obj
}
func (obj *deviceDeviceEthernetIter) clearHolderSlice() DeviceDeviceEthernetIter {
	if len(obj.deviceEthernetSlice) > 0 {
		obj.deviceEthernetSlice = []DeviceEthernet{}
	}
	return obj
}
func (obj *deviceDeviceEthernetIter) appendHolderSlice(item DeviceEthernet) DeviceDeviceEthernetIter {
	obj.deviceEthernetSlice = append(obj.deviceEthernetSlice, item)
	return obj
}

// IPv4 Loopback interface that can be attached to an Ethernet in the same device  or to an Ethernet in another device.
// Ipv4Loopbacks returns a []DeviceIpv4Loopback
func (obj *device) Ipv4Loopbacks() DeviceDeviceIpv4LoopbackIter {
	if len(obj.obj.Ipv4Loopbacks) == 0 {
		obj.obj.Ipv4Loopbacks = []*otg.DeviceIpv4Loopback{}
	}
	if obj.ipv4LoopbacksHolder == nil {
		obj.ipv4LoopbacksHolder = newDeviceDeviceIpv4LoopbackIter(&obj.obj.Ipv4Loopbacks).setMsg(obj)
	}
	return obj.ipv4LoopbacksHolder
}

type deviceDeviceIpv4LoopbackIter struct {
	obj                     *device
	deviceIpv4LoopbackSlice []DeviceIpv4Loopback
	fieldPtr                *[]*otg.DeviceIpv4Loopback
}

func newDeviceDeviceIpv4LoopbackIter(ptr *[]*otg.DeviceIpv4Loopback) DeviceDeviceIpv4LoopbackIter {
	return &deviceDeviceIpv4LoopbackIter{fieldPtr: ptr}
}

type DeviceDeviceIpv4LoopbackIter interface {
	setMsg(*device) DeviceDeviceIpv4LoopbackIter
	Items() []DeviceIpv4Loopback
	Add() DeviceIpv4Loopback
	Append(items ...DeviceIpv4Loopback) DeviceDeviceIpv4LoopbackIter
	Set(index int, newObj DeviceIpv4Loopback) DeviceDeviceIpv4LoopbackIter
	Clear() DeviceDeviceIpv4LoopbackIter
	clearHolderSlice() DeviceDeviceIpv4LoopbackIter
	appendHolderSlice(item DeviceIpv4Loopback) DeviceDeviceIpv4LoopbackIter
}

func (obj *deviceDeviceIpv4LoopbackIter) setMsg(msg *device) DeviceDeviceIpv4LoopbackIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&deviceIpv4Loopback{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceDeviceIpv4LoopbackIter) Items() []DeviceIpv4Loopback {
	return obj.deviceIpv4LoopbackSlice
}

func (obj *deviceDeviceIpv4LoopbackIter) Add() DeviceIpv4Loopback {
	newObj := &otg.DeviceIpv4Loopback{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &deviceIpv4Loopback{obj: newObj}
	newLibObj.setDefault()
	obj.deviceIpv4LoopbackSlice = append(obj.deviceIpv4LoopbackSlice, newLibObj)
	return newLibObj
}

func (obj *deviceDeviceIpv4LoopbackIter) Append(items ...DeviceIpv4Loopback) DeviceDeviceIpv4LoopbackIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.deviceIpv4LoopbackSlice = append(obj.deviceIpv4LoopbackSlice, item)
	}
	return obj
}

func (obj *deviceDeviceIpv4LoopbackIter) Set(index int, newObj DeviceIpv4Loopback) DeviceDeviceIpv4LoopbackIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.deviceIpv4LoopbackSlice[index] = newObj
	return obj
}
func (obj *deviceDeviceIpv4LoopbackIter) Clear() DeviceDeviceIpv4LoopbackIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.DeviceIpv4Loopback{}
		obj.deviceIpv4LoopbackSlice = []DeviceIpv4Loopback{}
	}
	return obj
}
func (obj *deviceDeviceIpv4LoopbackIter) clearHolderSlice() DeviceDeviceIpv4LoopbackIter {
	if len(obj.deviceIpv4LoopbackSlice) > 0 {
		obj.deviceIpv4LoopbackSlice = []DeviceIpv4Loopback{}
	}
	return obj
}
func (obj *deviceDeviceIpv4LoopbackIter) appendHolderSlice(item DeviceIpv4Loopback) DeviceDeviceIpv4LoopbackIter {
	obj.deviceIpv4LoopbackSlice = append(obj.deviceIpv4LoopbackSlice, item)
	return obj
}

// IPv6 Loopback interface that can be attached to an Ethernet in the same device  or to an Ethernet in another device.
// Ipv6Loopbacks returns a []DeviceIpv6Loopback
func (obj *device) Ipv6Loopbacks() DeviceDeviceIpv6LoopbackIter {
	if len(obj.obj.Ipv6Loopbacks) == 0 {
		obj.obj.Ipv6Loopbacks = []*otg.DeviceIpv6Loopback{}
	}
	if obj.ipv6LoopbacksHolder == nil {
		obj.ipv6LoopbacksHolder = newDeviceDeviceIpv6LoopbackIter(&obj.obj.Ipv6Loopbacks).setMsg(obj)
	}
	return obj.ipv6LoopbacksHolder
}

type deviceDeviceIpv6LoopbackIter struct {
	obj                     *device
	deviceIpv6LoopbackSlice []DeviceIpv6Loopback
	fieldPtr                *[]*otg.DeviceIpv6Loopback
}

func newDeviceDeviceIpv6LoopbackIter(ptr *[]*otg.DeviceIpv6Loopback) DeviceDeviceIpv6LoopbackIter {
	return &deviceDeviceIpv6LoopbackIter{fieldPtr: ptr}
}

type DeviceDeviceIpv6LoopbackIter interface {
	setMsg(*device) DeviceDeviceIpv6LoopbackIter
	Items() []DeviceIpv6Loopback
	Add() DeviceIpv6Loopback
	Append(items ...DeviceIpv6Loopback) DeviceDeviceIpv6LoopbackIter
	Set(index int, newObj DeviceIpv6Loopback) DeviceDeviceIpv6LoopbackIter
	Clear() DeviceDeviceIpv6LoopbackIter
	clearHolderSlice() DeviceDeviceIpv6LoopbackIter
	appendHolderSlice(item DeviceIpv6Loopback) DeviceDeviceIpv6LoopbackIter
}

func (obj *deviceDeviceIpv6LoopbackIter) setMsg(msg *device) DeviceDeviceIpv6LoopbackIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&deviceIpv6Loopback{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceDeviceIpv6LoopbackIter) Items() []DeviceIpv6Loopback {
	return obj.deviceIpv6LoopbackSlice
}

func (obj *deviceDeviceIpv6LoopbackIter) Add() DeviceIpv6Loopback {
	newObj := &otg.DeviceIpv6Loopback{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &deviceIpv6Loopback{obj: newObj}
	newLibObj.setDefault()
	obj.deviceIpv6LoopbackSlice = append(obj.deviceIpv6LoopbackSlice, newLibObj)
	return newLibObj
}

func (obj *deviceDeviceIpv6LoopbackIter) Append(items ...DeviceIpv6Loopback) DeviceDeviceIpv6LoopbackIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.deviceIpv6LoopbackSlice = append(obj.deviceIpv6LoopbackSlice, item)
	}
	return obj
}

func (obj *deviceDeviceIpv6LoopbackIter) Set(index int, newObj DeviceIpv6Loopback) DeviceDeviceIpv6LoopbackIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.deviceIpv6LoopbackSlice[index] = newObj
	return obj
}
func (obj *deviceDeviceIpv6LoopbackIter) Clear() DeviceDeviceIpv6LoopbackIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.DeviceIpv6Loopback{}
		obj.deviceIpv6LoopbackSlice = []DeviceIpv6Loopback{}
	}
	return obj
}
func (obj *deviceDeviceIpv6LoopbackIter) clearHolderSlice() DeviceDeviceIpv6LoopbackIter {
	if len(obj.deviceIpv6LoopbackSlice) > 0 {
		obj.deviceIpv6LoopbackSlice = []DeviceIpv6Loopback{}
	}
	return obj
}
func (obj *deviceDeviceIpv6LoopbackIter) appendHolderSlice(item DeviceIpv6Loopback) DeviceDeviceIpv6LoopbackIter {
	obj.deviceIpv6LoopbackSlice = append(obj.deviceIpv6LoopbackSlice, item)
	return obj
}

// The properties of an IS-IS router and its children,  such as IS-IS interfaces and route ranges.
// Isis returns a DeviceIsisRouter
func (obj *device) Isis() DeviceIsisRouter {
	if obj.obj.Isis == nil {
		obj.obj.Isis = NewDeviceIsisRouter().msg()
	}
	if obj.isisHolder == nil {
		obj.isisHolder = &deviceIsisRouter{obj: obj.obj.Isis}
	}
	return obj.isisHolder
}

// The properties of an IS-IS router and its children,  such as IS-IS interfaces and route ranges.
// Isis returns a DeviceIsisRouter
func (obj *device) HasIsis() bool {
	return obj.obj.Isis != nil
}

// The properties of an IS-IS router and its children,  such as IS-IS interfaces and route ranges.
// SetIsis sets the DeviceIsisRouter value in the Device object
func (obj *device) SetIsis(value DeviceIsisRouter) Device {

	obj.isisHolder = nil
	obj.obj.Isis = value.msg()

	return obj
}

// The properties of BGP router and its children,  such as BGPv4, BGPv6 peers and their route ranges.
// Bgp returns a DeviceBgpRouter
func (obj *device) Bgp() DeviceBgpRouter {
	if obj.obj.Bgp == nil {
		obj.obj.Bgp = NewDeviceBgpRouter().msg()
	}
	if obj.bgpHolder == nil {
		obj.bgpHolder = &deviceBgpRouter{obj: obj.obj.Bgp}
	}
	return obj.bgpHolder
}

// The properties of BGP router and its children,  such as BGPv4, BGPv6 peers and their route ranges.
// Bgp returns a DeviceBgpRouter
func (obj *device) HasBgp() bool {
	return obj.obj.Bgp != nil
}

// The properties of BGP router and its children,  such as BGPv4, BGPv6 peers and their route ranges.
// SetBgp sets the DeviceBgpRouter value in the Device object
func (obj *device) SetBgp(value DeviceBgpRouter) Device {

	obj.bgpHolder = nil
	obj.obj.Bgp = value.msg()

	return obj
}

// Configuration of VXLAN tunnel interfaces RFC Ref: https://datatracker.ietf.org/doc/html/rfc7348
// Vxlan returns a DeviceVxlan
func (obj *device) Vxlan() DeviceVxlan {
	if obj.obj.Vxlan == nil {
		obj.obj.Vxlan = NewDeviceVxlan().msg()
	}
	if obj.vxlanHolder == nil {
		obj.vxlanHolder = &deviceVxlan{obj: obj.obj.Vxlan}
	}
	return obj.vxlanHolder
}

// Configuration of VXLAN tunnel interfaces RFC Ref: https://datatracker.ietf.org/doc/html/rfc7348
// Vxlan returns a DeviceVxlan
func (obj *device) HasVxlan() bool {
	return obj.obj.Vxlan != nil
}

// Configuration of VXLAN tunnel interfaces RFC Ref: https://datatracker.ietf.org/doc/html/rfc7348
// SetVxlan sets the DeviceVxlan value in the Device object
func (obj *device) SetVxlan(value DeviceVxlan) Device {

	obj.vxlanHolder = nil
	obj.obj.Vxlan = value.msg()

	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *device) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the Device object
func (obj *device) SetName(value string) Device {

	obj.obj.Name = &value
	return obj
}

// The properties of an RSVP router and its children.
// Rsvp returns a DeviceRsvp
func (obj *device) Rsvp() DeviceRsvp {
	if obj.obj.Rsvp == nil {
		obj.obj.Rsvp = NewDeviceRsvp().msg()
	}
	if obj.rsvpHolder == nil {
		obj.rsvpHolder = &deviceRsvp{obj: obj.obj.Rsvp}
	}
	return obj.rsvpHolder
}

// The properties of an RSVP router and its children.
// Rsvp returns a DeviceRsvp
func (obj *device) HasRsvp() bool {
	return obj.obj.Rsvp != nil
}

// The properties of an RSVP router and its children.
// SetRsvp sets the DeviceRsvp value in the Device object
func (obj *device) SetRsvp(value DeviceRsvp) Device {

	obj.rsvpHolder = nil
	obj.obj.Rsvp = value.msg()

	return obj
}

// The properties of DHCP Server and its children, such as DHCPv4, DHCPv6 servers.
// DhcpServer returns a DeviceDhcpServer
func (obj *device) DhcpServer() DeviceDhcpServer {
	if obj.obj.DhcpServer == nil {
		obj.obj.DhcpServer = NewDeviceDhcpServer().msg()
	}
	if obj.dhcpServerHolder == nil {
		obj.dhcpServerHolder = &deviceDhcpServer{obj: obj.obj.DhcpServer}
	}
	return obj.dhcpServerHolder
}

// The properties of DHCP Server and its children, such as DHCPv4, DHCPv6 servers.
// DhcpServer returns a DeviceDhcpServer
func (obj *device) HasDhcpServer() bool {
	return obj.obj.DhcpServer != nil
}

// The properties of DHCP Server and its children, such as DHCPv4, DHCPv6 servers.
// SetDhcpServer sets the DeviceDhcpServer value in the Device object
func (obj *device) SetDhcpServer(value DeviceDhcpServer) Device {

	obj.dhcpServerHolder = nil
	obj.obj.DhcpServer = value.msg()

	return obj
}

// Configuration for OSPFv2 router.
// Ospfv2 returns a DeviceOspfv2Router
func (obj *device) Ospfv2() DeviceOspfv2Router {
	if obj.obj.Ospfv2 == nil {
		obj.obj.Ospfv2 = NewDeviceOspfv2Router().msg()
	}
	if obj.ospfv2Holder == nil {
		obj.ospfv2Holder = &deviceOspfv2Router{obj: obj.obj.Ospfv2}
	}
	return obj.ospfv2Holder
}

// Configuration for OSPFv2 router.
// Ospfv2 returns a DeviceOspfv2Router
func (obj *device) HasOspfv2() bool {
	return obj.obj.Ospfv2 != nil
}

// Configuration for OSPFv2 router.
// SetOspfv2 sets the DeviceOspfv2Router value in the Device object
func (obj *device) SetOspfv2(value DeviceOspfv2Router) Device {

	obj.ospfv2Holder = nil
	obj.obj.Ospfv2 = value.msg()

	return obj
}

// Configuration of MACsec device.
// Macsec returns a DeviceMacsec
func (obj *device) Macsec() DeviceMacsec {
	if obj.obj.Macsec == nil {
		obj.obj.Macsec = NewDeviceMacsec().msg()
	}
	if obj.macsecHolder == nil {
		obj.macsecHolder = &deviceMacsec{obj: obj.obj.Macsec}
	}
	return obj.macsecHolder
}

// Configuration of MACsec device.
// Macsec returns a DeviceMacsec
func (obj *device) HasMacsec() bool {
	return obj.obj.Macsec != nil
}

// Configuration of MACsec device.
// SetMacsec sets the DeviceMacsec value in the Device object
func (obj *device) SetMacsec(value DeviceMacsec) Device {

	obj.macsecHolder = nil
	obj.obj.Macsec = value.msg()

	return obj
}

func (obj *device) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Ethernets) != 0 {

		if set_default {
			obj.Ethernets().clearHolderSlice()
			for _, item := range obj.obj.Ethernets {
				obj.Ethernets().appendHolderSlice(&deviceEthernet{obj: item})
			}
		}
		for _, item := range obj.Ethernets().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Ipv4Loopbacks) != 0 {

		if set_default {
			obj.Ipv4Loopbacks().clearHolderSlice()
			for _, item := range obj.obj.Ipv4Loopbacks {
				obj.Ipv4Loopbacks().appendHolderSlice(&deviceIpv4Loopback{obj: item})
			}
		}
		for _, item := range obj.Ipv4Loopbacks().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Ipv6Loopbacks) != 0 {

		if set_default {
			obj.Ipv6Loopbacks().clearHolderSlice()
			for _, item := range obj.obj.Ipv6Loopbacks {
				obj.Ipv6Loopbacks().appendHolderSlice(&deviceIpv6Loopback{obj: item})
			}
		}
		for _, item := range obj.Ipv6Loopbacks().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.Isis != nil {

		obj.Isis().validateObj(vObj, set_default)
	}

	if obj.obj.Bgp != nil {

		obj.Bgp().validateObj(vObj, set_default)
	}

	if obj.obj.Vxlan != nil {

		obj.Vxlan().validateObj(vObj, set_default)
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface Device")
	}

	if obj.obj.Rsvp != nil {

		obj.Rsvp().validateObj(vObj, set_default)
	}

	if obj.obj.DhcpServer != nil {

		obj.DhcpServer().validateObj(vObj, set_default)
	}

	if obj.obj.Ospfv2 != nil {

		obj.Ospfv2().validateObj(vObj, set_default)
	}

	if obj.obj.Macsec != nil {

		obj.Macsec().validateObj(vObj, set_default)
	}

}

func (obj *device) setDefault() {

}
