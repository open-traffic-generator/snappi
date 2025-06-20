package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceDhcpServer *****
type deviceDhcpServer struct {
	validation
	obj                  *otg.DeviceDhcpServer
	marshaller           marshalDeviceDhcpServer
	unMarshaller         unMarshalDeviceDhcpServer
	ipv4InterfacesHolder DeviceDhcpServerDhcpServerV4Iter
	ipv6InterfacesHolder DeviceDhcpServerDhcpServerV6Iter
}

func NewDeviceDhcpServer() DeviceDhcpServer {
	obj := deviceDhcpServer{obj: &otg.DeviceDhcpServer{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceDhcpServer) msg() *otg.DeviceDhcpServer {
	return obj.obj
}

func (obj *deviceDhcpServer) setMsg(msg *otg.DeviceDhcpServer) DeviceDhcpServer {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceDhcpServer struct {
	obj *deviceDhcpServer
}

type marshalDeviceDhcpServer interface {
	// ToProto marshals DeviceDhcpServer to protobuf object *otg.DeviceDhcpServer
	ToProto() (*otg.DeviceDhcpServer, error)
	// ToPbText marshals DeviceDhcpServer to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceDhcpServer to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceDhcpServer to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceDhcpServer struct {
	obj *deviceDhcpServer
}

type unMarshalDeviceDhcpServer interface {
	// FromProto unmarshals DeviceDhcpServer from protobuf object *otg.DeviceDhcpServer
	FromProto(msg *otg.DeviceDhcpServer) (DeviceDhcpServer, error)
	// FromPbText unmarshals DeviceDhcpServer from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceDhcpServer from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceDhcpServer from JSON text
	FromJson(value string) error
}

func (obj *deviceDhcpServer) Marshal() marshalDeviceDhcpServer {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceDhcpServer{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceDhcpServer) Unmarshal() unMarshalDeviceDhcpServer {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceDhcpServer{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceDhcpServer) ToProto() (*otg.DeviceDhcpServer, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceDhcpServer) FromProto(msg *otg.DeviceDhcpServer) (DeviceDhcpServer, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceDhcpServer) ToPbText() (string, error) {
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

func (m *unMarshaldeviceDhcpServer) FromPbText(value string) error {
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

func (m *marshaldeviceDhcpServer) ToYaml() (string, error) {
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

func (m *unMarshaldeviceDhcpServer) FromYaml(value string) error {
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

func (m *marshaldeviceDhcpServer) ToJson() (string, error) {
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

func (m *unMarshaldeviceDhcpServer) FromJson(value string) error {
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

func (obj *deviceDhcpServer) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceDhcpServer) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceDhcpServer) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceDhcpServer) Clone() (DeviceDhcpServer, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceDhcpServer()
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

func (obj *deviceDhcpServer) setNil() {
	obj.ipv4InterfacesHolder = nil
	obj.ipv6InterfacesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceDhcpServer is configuration for one or more IPv4 or IPv6 DHCP servers.
type DeviceDhcpServer interface {
	Validation
	// msg marshals DeviceDhcpServer to protobuf object *otg.DeviceDhcpServer
	// and doesn't set defaults
	msg() *otg.DeviceDhcpServer
	// setMsg unmarshals DeviceDhcpServer from protobuf object *otg.DeviceDhcpServer
	// and doesn't set defaults
	setMsg(*otg.DeviceDhcpServer) DeviceDhcpServer
	// provides marshal interface
	Marshal() marshalDeviceDhcpServer
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceDhcpServer
	// validate validates DeviceDhcpServer
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceDhcpServer, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4Interfaces returns DeviceDhcpServerDhcpServerV4IterIter, set in DeviceDhcpServer
	Ipv4Interfaces() DeviceDhcpServerDhcpServerV4Iter
	// Ipv6Interfaces returns DeviceDhcpServerDhcpServerV6IterIter, set in DeviceDhcpServer
	Ipv6Interfaces() DeviceDhcpServerDhcpServerV6Iter
	setNil()
}

// This contains an array of references to IPv4 interfaces, each of which will contain one DHCPv4 server.
// Ipv4Interfaces returns a []DhcpServerV4
func (obj *deviceDhcpServer) Ipv4Interfaces() DeviceDhcpServerDhcpServerV4Iter {
	if len(obj.obj.Ipv4Interfaces) == 0 {
		obj.obj.Ipv4Interfaces = []*otg.DhcpServerV4{}
	}
	if obj.ipv4InterfacesHolder == nil {
		obj.ipv4InterfacesHolder = newDeviceDhcpServerDhcpServerV4Iter(&obj.obj.Ipv4Interfaces).setMsg(obj)
	}
	return obj.ipv4InterfacesHolder
}

type deviceDhcpServerDhcpServerV4Iter struct {
	obj               *deviceDhcpServer
	dhcpServerV4Slice []DhcpServerV4
	fieldPtr          *[]*otg.DhcpServerV4
}

func newDeviceDhcpServerDhcpServerV4Iter(ptr *[]*otg.DhcpServerV4) DeviceDhcpServerDhcpServerV4Iter {
	return &deviceDhcpServerDhcpServerV4Iter{fieldPtr: ptr}
}

type DeviceDhcpServerDhcpServerV4Iter interface {
	setMsg(*deviceDhcpServer) DeviceDhcpServerDhcpServerV4Iter
	Items() []DhcpServerV4
	Add() DhcpServerV4
	Append(items ...DhcpServerV4) DeviceDhcpServerDhcpServerV4Iter
	Set(index int, newObj DhcpServerV4) DeviceDhcpServerDhcpServerV4Iter
	Clear() DeviceDhcpServerDhcpServerV4Iter
	clearHolderSlice() DeviceDhcpServerDhcpServerV4Iter
	appendHolderSlice(item DhcpServerV4) DeviceDhcpServerDhcpServerV4Iter
}

func (obj *deviceDhcpServerDhcpServerV4Iter) setMsg(msg *deviceDhcpServer) DeviceDhcpServerDhcpServerV4Iter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&dhcpServerV4{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceDhcpServerDhcpServerV4Iter) Items() []DhcpServerV4 {
	return obj.dhcpServerV4Slice
}

func (obj *deviceDhcpServerDhcpServerV4Iter) Add() DhcpServerV4 {
	newObj := &otg.DhcpServerV4{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &dhcpServerV4{obj: newObj}
	newLibObj.setDefault()
	obj.dhcpServerV4Slice = append(obj.dhcpServerV4Slice, newLibObj)
	return newLibObj
}

func (obj *deviceDhcpServerDhcpServerV4Iter) Append(items ...DhcpServerV4) DeviceDhcpServerDhcpServerV4Iter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.dhcpServerV4Slice = append(obj.dhcpServerV4Slice, item)
	}
	return obj
}

func (obj *deviceDhcpServerDhcpServerV4Iter) Set(index int, newObj DhcpServerV4) DeviceDhcpServerDhcpServerV4Iter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.dhcpServerV4Slice[index] = newObj
	return obj
}
func (obj *deviceDhcpServerDhcpServerV4Iter) Clear() DeviceDhcpServerDhcpServerV4Iter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.DhcpServerV4{}
		obj.dhcpServerV4Slice = []DhcpServerV4{}
	}
	return obj
}
func (obj *deviceDhcpServerDhcpServerV4Iter) clearHolderSlice() DeviceDhcpServerDhcpServerV4Iter {
	if len(obj.dhcpServerV4Slice) > 0 {
		obj.dhcpServerV4Slice = []DhcpServerV4{}
	}
	return obj
}
func (obj *deviceDhcpServerDhcpServerV4Iter) appendHolderSlice(item DhcpServerV4) DeviceDhcpServerDhcpServerV4Iter {
	obj.dhcpServerV4Slice = append(obj.dhcpServerV4Slice, item)
	return obj
}

// This contains an array of references to IPv6 interfaces, each of which will contain one DHCPv6 server.
// Ipv6Interfaces returns a []DhcpServerV6
func (obj *deviceDhcpServer) Ipv6Interfaces() DeviceDhcpServerDhcpServerV6Iter {
	if len(obj.obj.Ipv6Interfaces) == 0 {
		obj.obj.Ipv6Interfaces = []*otg.DhcpServerV6{}
	}
	if obj.ipv6InterfacesHolder == nil {
		obj.ipv6InterfacesHolder = newDeviceDhcpServerDhcpServerV6Iter(&obj.obj.Ipv6Interfaces).setMsg(obj)
	}
	return obj.ipv6InterfacesHolder
}

type deviceDhcpServerDhcpServerV6Iter struct {
	obj               *deviceDhcpServer
	dhcpServerV6Slice []DhcpServerV6
	fieldPtr          *[]*otg.DhcpServerV6
}

func newDeviceDhcpServerDhcpServerV6Iter(ptr *[]*otg.DhcpServerV6) DeviceDhcpServerDhcpServerV6Iter {
	return &deviceDhcpServerDhcpServerV6Iter{fieldPtr: ptr}
}

type DeviceDhcpServerDhcpServerV6Iter interface {
	setMsg(*deviceDhcpServer) DeviceDhcpServerDhcpServerV6Iter
	Items() []DhcpServerV6
	Add() DhcpServerV6
	Append(items ...DhcpServerV6) DeviceDhcpServerDhcpServerV6Iter
	Set(index int, newObj DhcpServerV6) DeviceDhcpServerDhcpServerV6Iter
	Clear() DeviceDhcpServerDhcpServerV6Iter
	clearHolderSlice() DeviceDhcpServerDhcpServerV6Iter
	appendHolderSlice(item DhcpServerV6) DeviceDhcpServerDhcpServerV6Iter
}

func (obj *deviceDhcpServerDhcpServerV6Iter) setMsg(msg *deviceDhcpServer) DeviceDhcpServerDhcpServerV6Iter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&dhcpServerV6{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceDhcpServerDhcpServerV6Iter) Items() []DhcpServerV6 {
	return obj.dhcpServerV6Slice
}

func (obj *deviceDhcpServerDhcpServerV6Iter) Add() DhcpServerV6 {
	newObj := &otg.DhcpServerV6{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &dhcpServerV6{obj: newObj}
	newLibObj.setDefault()
	obj.dhcpServerV6Slice = append(obj.dhcpServerV6Slice, newLibObj)
	return newLibObj
}

func (obj *deviceDhcpServerDhcpServerV6Iter) Append(items ...DhcpServerV6) DeviceDhcpServerDhcpServerV6Iter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.dhcpServerV6Slice = append(obj.dhcpServerV6Slice, item)
	}
	return obj
}

func (obj *deviceDhcpServerDhcpServerV6Iter) Set(index int, newObj DhcpServerV6) DeviceDhcpServerDhcpServerV6Iter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.dhcpServerV6Slice[index] = newObj
	return obj
}
func (obj *deviceDhcpServerDhcpServerV6Iter) Clear() DeviceDhcpServerDhcpServerV6Iter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.DhcpServerV6{}
		obj.dhcpServerV6Slice = []DhcpServerV6{}
	}
	return obj
}
func (obj *deviceDhcpServerDhcpServerV6Iter) clearHolderSlice() DeviceDhcpServerDhcpServerV6Iter {
	if len(obj.dhcpServerV6Slice) > 0 {
		obj.dhcpServerV6Slice = []DhcpServerV6{}
	}
	return obj
}
func (obj *deviceDhcpServerDhcpServerV6Iter) appendHolderSlice(item DhcpServerV6) DeviceDhcpServerDhcpServerV6Iter {
	obj.dhcpServerV6Slice = append(obj.dhcpServerV6Slice, item)
	return obj
}

func (obj *deviceDhcpServer) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Ipv4Interfaces) != 0 {

		if set_default {
			obj.Ipv4Interfaces().clearHolderSlice()
			for _, item := range obj.obj.Ipv4Interfaces {
				obj.Ipv4Interfaces().appendHolderSlice(&dhcpServerV4{obj: item})
			}
		}
		for _, item := range obj.Ipv4Interfaces().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Ipv6Interfaces) != 0 {

		if set_default {
			obj.Ipv6Interfaces().clearHolderSlice()
			for _, item := range obj.obj.Ipv6Interfaces {
				obj.Ipv6Interfaces().appendHolderSlice(&dhcpServerV6{obj: item})
			}
		}
		for _, item := range obj.Ipv6Interfaces().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *deviceDhcpServer) setDefault() {

}
