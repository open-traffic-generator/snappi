package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceBmp *****
type deviceBmp struct {
	validation
	obj                  *otg.DeviceBmp
	marshaller           marshalDeviceBmp
	unMarshaller         unMarshalDeviceBmp
	ipv4InterfacesHolder DeviceBmpDeviceBmpV4InterfaceIter
	ipv6InterfacesHolder DeviceBmpDeviceBmpV6InterfaceIter
}

func NewDeviceBmp() DeviceBmp {
	obj := deviceBmp{obj: &otg.DeviceBmp{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceBmp) msg() *otg.DeviceBmp {
	return obj.obj
}

func (obj *deviceBmp) setMsg(msg *otg.DeviceBmp) DeviceBmp {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceBmp struct {
	obj *deviceBmp
}

type marshalDeviceBmp interface {
	// ToProto marshals DeviceBmp to protobuf object *otg.DeviceBmp
	ToProto() (*otg.DeviceBmp, error)
	// ToPbText marshals DeviceBmp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceBmp to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceBmp to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceBmp struct {
	obj *deviceBmp
}

type unMarshalDeviceBmp interface {
	// FromProto unmarshals DeviceBmp from protobuf object *otg.DeviceBmp
	FromProto(msg *otg.DeviceBmp) (DeviceBmp, error)
	// FromPbText unmarshals DeviceBmp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceBmp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceBmp from JSON text
	FromJson(value string) error
}

func (obj *deviceBmp) Marshal() marshalDeviceBmp {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceBmp{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceBmp) Unmarshal() unMarshalDeviceBmp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceBmp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceBmp) ToProto() (*otg.DeviceBmp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceBmp) FromProto(msg *otg.DeviceBmp) (DeviceBmp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceBmp) ToPbText() (string, error) {
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

func (m *unMarshaldeviceBmp) FromPbText(value string) error {
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

func (m *marshaldeviceBmp) ToYaml() (string, error) {
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

func (m *unMarshaldeviceBmp) FromYaml(value string) error {
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

func (m *marshaldeviceBmp) ToJson() (string, error) {
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

func (m *unMarshaldeviceBmp) FromJson(value string) error {
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

func (obj *deviceBmp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceBmp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceBmp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceBmp) Clone() (DeviceBmp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceBmp()
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

func (obj *deviceBmp) setNil() {
	obj.ipv4InterfacesHolder = nil
	obj.ipv6InterfacesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceBmp is top-level container for BGP Monitoring Protocol (BMP) configuration. BMP, as defined in RFC 7854, provides a mechanism to monitor BGP sessions. This configuration pertains to the device when acting as a BMP Monitor /server, listening for connections from BGP speakers (routers) acting as BMP clients. BMP is unidirectional, meaning the monitoring station only receives information; it doesn't send commands to or control the monitored router.
type DeviceBmp interface {
	Validation
	// msg marshals DeviceBmp to protobuf object *otg.DeviceBmp
	// and doesn't set defaults
	msg() *otg.DeviceBmp
	// setMsg unmarshals DeviceBmp from protobuf object *otg.DeviceBmp
	// and doesn't set defaults
	setMsg(*otg.DeviceBmp) DeviceBmp
	// provides marshal interface
	Marshal() marshalDeviceBmp
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceBmp
	// validate validates DeviceBmp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceBmp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4Interfaces returns DeviceBmpDeviceBmpV4InterfaceIterIter, set in DeviceBmp
	Ipv4Interfaces() DeviceBmpDeviceBmpV4InterfaceIter
	// Ipv6Interfaces returns DeviceBmpDeviceBmpV6InterfaceIterIter, set in DeviceBmp
	Ipv6Interfaces() DeviceBmpDeviceBmpV6InterfaceIter
	setNil()
}

// This contains an array of references to IPv4 interfaces, each of which will have one or more BMP Servers configured, each BMP Server will have a connection to a different monitored router.
// Ipv4Interfaces returns a []DeviceBmpV4Interface
func (obj *deviceBmp) Ipv4Interfaces() DeviceBmpDeviceBmpV4InterfaceIter {
	if len(obj.obj.Ipv4Interfaces) == 0 {
		obj.obj.Ipv4Interfaces = []*otg.DeviceBmpV4Interface{}
	}
	if obj.ipv4InterfacesHolder == nil {
		obj.ipv4InterfacesHolder = newDeviceBmpDeviceBmpV4InterfaceIter(&obj.obj.Ipv4Interfaces).setMsg(obj)
	}
	return obj.ipv4InterfacesHolder
}

type deviceBmpDeviceBmpV4InterfaceIter struct {
	obj                       *deviceBmp
	deviceBmpV4InterfaceSlice []DeviceBmpV4Interface
	fieldPtr                  *[]*otg.DeviceBmpV4Interface
}

func newDeviceBmpDeviceBmpV4InterfaceIter(ptr *[]*otg.DeviceBmpV4Interface) DeviceBmpDeviceBmpV4InterfaceIter {
	return &deviceBmpDeviceBmpV4InterfaceIter{fieldPtr: ptr}
}

type DeviceBmpDeviceBmpV4InterfaceIter interface {
	setMsg(*deviceBmp) DeviceBmpDeviceBmpV4InterfaceIter
	Items() []DeviceBmpV4Interface
	Add() DeviceBmpV4Interface
	Append(items ...DeviceBmpV4Interface) DeviceBmpDeviceBmpV4InterfaceIter
	Set(index int, newObj DeviceBmpV4Interface) DeviceBmpDeviceBmpV4InterfaceIter
	Clear() DeviceBmpDeviceBmpV4InterfaceIter
	clearHolderSlice() DeviceBmpDeviceBmpV4InterfaceIter
	appendHolderSlice(item DeviceBmpV4Interface) DeviceBmpDeviceBmpV4InterfaceIter
}

func (obj *deviceBmpDeviceBmpV4InterfaceIter) setMsg(msg *deviceBmp) DeviceBmpDeviceBmpV4InterfaceIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&deviceBmpV4Interface{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceBmpDeviceBmpV4InterfaceIter) Items() []DeviceBmpV4Interface {
	return obj.deviceBmpV4InterfaceSlice
}

func (obj *deviceBmpDeviceBmpV4InterfaceIter) Add() DeviceBmpV4Interface {
	newObj := &otg.DeviceBmpV4Interface{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &deviceBmpV4Interface{obj: newObj}
	newLibObj.setDefault()
	obj.deviceBmpV4InterfaceSlice = append(obj.deviceBmpV4InterfaceSlice, newLibObj)
	return newLibObj
}

func (obj *deviceBmpDeviceBmpV4InterfaceIter) Append(items ...DeviceBmpV4Interface) DeviceBmpDeviceBmpV4InterfaceIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.deviceBmpV4InterfaceSlice = append(obj.deviceBmpV4InterfaceSlice, item)
	}
	return obj
}

func (obj *deviceBmpDeviceBmpV4InterfaceIter) Set(index int, newObj DeviceBmpV4Interface) DeviceBmpDeviceBmpV4InterfaceIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.deviceBmpV4InterfaceSlice[index] = newObj
	return obj
}
func (obj *deviceBmpDeviceBmpV4InterfaceIter) Clear() DeviceBmpDeviceBmpV4InterfaceIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.DeviceBmpV4Interface{}
		obj.deviceBmpV4InterfaceSlice = []DeviceBmpV4Interface{}
	}
	return obj
}
func (obj *deviceBmpDeviceBmpV4InterfaceIter) clearHolderSlice() DeviceBmpDeviceBmpV4InterfaceIter {
	if len(obj.deviceBmpV4InterfaceSlice) > 0 {
		obj.deviceBmpV4InterfaceSlice = []DeviceBmpV4Interface{}
	}
	return obj
}
func (obj *deviceBmpDeviceBmpV4InterfaceIter) appendHolderSlice(item DeviceBmpV4Interface) DeviceBmpDeviceBmpV4InterfaceIter {
	obj.deviceBmpV4InterfaceSlice = append(obj.deviceBmpV4InterfaceSlice, item)
	return obj
}

// This contains an array of references to IPv6 interfaces, each of which will have one or more BMP Servers configured, each BMP Server will have a connection to a different monitored router.
// Ipv6Interfaces returns a []DeviceBmpV6Interface
func (obj *deviceBmp) Ipv6Interfaces() DeviceBmpDeviceBmpV6InterfaceIter {
	if len(obj.obj.Ipv6Interfaces) == 0 {
		obj.obj.Ipv6Interfaces = []*otg.DeviceBmpV6Interface{}
	}
	if obj.ipv6InterfacesHolder == nil {
		obj.ipv6InterfacesHolder = newDeviceBmpDeviceBmpV6InterfaceIter(&obj.obj.Ipv6Interfaces).setMsg(obj)
	}
	return obj.ipv6InterfacesHolder
}

type deviceBmpDeviceBmpV6InterfaceIter struct {
	obj                       *deviceBmp
	deviceBmpV6InterfaceSlice []DeviceBmpV6Interface
	fieldPtr                  *[]*otg.DeviceBmpV6Interface
}

func newDeviceBmpDeviceBmpV6InterfaceIter(ptr *[]*otg.DeviceBmpV6Interface) DeviceBmpDeviceBmpV6InterfaceIter {
	return &deviceBmpDeviceBmpV6InterfaceIter{fieldPtr: ptr}
}

type DeviceBmpDeviceBmpV6InterfaceIter interface {
	setMsg(*deviceBmp) DeviceBmpDeviceBmpV6InterfaceIter
	Items() []DeviceBmpV6Interface
	Add() DeviceBmpV6Interface
	Append(items ...DeviceBmpV6Interface) DeviceBmpDeviceBmpV6InterfaceIter
	Set(index int, newObj DeviceBmpV6Interface) DeviceBmpDeviceBmpV6InterfaceIter
	Clear() DeviceBmpDeviceBmpV6InterfaceIter
	clearHolderSlice() DeviceBmpDeviceBmpV6InterfaceIter
	appendHolderSlice(item DeviceBmpV6Interface) DeviceBmpDeviceBmpV6InterfaceIter
}

func (obj *deviceBmpDeviceBmpV6InterfaceIter) setMsg(msg *deviceBmp) DeviceBmpDeviceBmpV6InterfaceIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&deviceBmpV6Interface{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceBmpDeviceBmpV6InterfaceIter) Items() []DeviceBmpV6Interface {
	return obj.deviceBmpV6InterfaceSlice
}

func (obj *deviceBmpDeviceBmpV6InterfaceIter) Add() DeviceBmpV6Interface {
	newObj := &otg.DeviceBmpV6Interface{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &deviceBmpV6Interface{obj: newObj}
	newLibObj.setDefault()
	obj.deviceBmpV6InterfaceSlice = append(obj.deviceBmpV6InterfaceSlice, newLibObj)
	return newLibObj
}

func (obj *deviceBmpDeviceBmpV6InterfaceIter) Append(items ...DeviceBmpV6Interface) DeviceBmpDeviceBmpV6InterfaceIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.deviceBmpV6InterfaceSlice = append(obj.deviceBmpV6InterfaceSlice, item)
	}
	return obj
}

func (obj *deviceBmpDeviceBmpV6InterfaceIter) Set(index int, newObj DeviceBmpV6Interface) DeviceBmpDeviceBmpV6InterfaceIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.deviceBmpV6InterfaceSlice[index] = newObj
	return obj
}
func (obj *deviceBmpDeviceBmpV6InterfaceIter) Clear() DeviceBmpDeviceBmpV6InterfaceIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.DeviceBmpV6Interface{}
		obj.deviceBmpV6InterfaceSlice = []DeviceBmpV6Interface{}
	}
	return obj
}
func (obj *deviceBmpDeviceBmpV6InterfaceIter) clearHolderSlice() DeviceBmpDeviceBmpV6InterfaceIter {
	if len(obj.deviceBmpV6InterfaceSlice) > 0 {
		obj.deviceBmpV6InterfaceSlice = []DeviceBmpV6Interface{}
	}
	return obj
}
func (obj *deviceBmpDeviceBmpV6InterfaceIter) appendHolderSlice(item DeviceBmpV6Interface) DeviceBmpDeviceBmpV6InterfaceIter {
	obj.deviceBmpV6InterfaceSlice = append(obj.deviceBmpV6InterfaceSlice, item)
	return obj
}

func (obj *deviceBmp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Ipv4Interfaces) != 0 {

		if set_default {
			obj.Ipv4Interfaces().clearHolderSlice()
			for _, item := range obj.obj.Ipv4Interfaces {
				obj.Ipv4Interfaces().appendHolderSlice(&deviceBmpV4Interface{obj: item})
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
				obj.Ipv6Interfaces().appendHolderSlice(&deviceBmpV6Interface{obj: item})
			}
		}
		for _, item := range obj.Ipv6Interfaces().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *deviceBmp) setDefault() {

}
