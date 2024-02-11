package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceBgpRouter *****
type deviceBgpRouter struct {
	validation
	obj                  *otg.DeviceBgpRouter
	marshaller           marshalDeviceBgpRouter
	unMarshaller         unMarshalDeviceBgpRouter
	ipv4InterfacesHolder DeviceBgpRouterBgpV4InterfaceIter
	ipv6InterfacesHolder DeviceBgpRouterBgpV6InterfaceIter
}

func NewDeviceBgpRouter() DeviceBgpRouter {
	obj := deviceBgpRouter{obj: &otg.DeviceBgpRouter{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceBgpRouter) msg() *otg.DeviceBgpRouter {
	return obj.obj
}

func (obj *deviceBgpRouter) setMsg(msg *otg.DeviceBgpRouter) DeviceBgpRouter {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceBgpRouter struct {
	obj *deviceBgpRouter
}

type marshalDeviceBgpRouter interface {
	// ToProto marshals DeviceBgpRouter to protobuf object *otg.DeviceBgpRouter
	ToProto() (*otg.DeviceBgpRouter, error)
	// ToPbText marshals DeviceBgpRouter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceBgpRouter to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceBgpRouter to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceBgpRouter struct {
	obj *deviceBgpRouter
}

type unMarshalDeviceBgpRouter interface {
	// FromProto unmarshals DeviceBgpRouter from protobuf object *otg.DeviceBgpRouter
	FromProto(msg *otg.DeviceBgpRouter) (DeviceBgpRouter, error)
	// FromPbText unmarshals DeviceBgpRouter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceBgpRouter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceBgpRouter from JSON text
	FromJson(value string) error
}

func (obj *deviceBgpRouter) Marshal() marshalDeviceBgpRouter {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceBgpRouter{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceBgpRouter) Unmarshal() unMarshalDeviceBgpRouter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceBgpRouter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceBgpRouter) ToProto() (*otg.DeviceBgpRouter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceBgpRouter) FromProto(msg *otg.DeviceBgpRouter) (DeviceBgpRouter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceBgpRouter) ToPbText() (string, error) {
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

func (m *unMarshaldeviceBgpRouter) FromPbText(value string) error {
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

func (m *marshaldeviceBgpRouter) ToYaml() (string, error) {
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

func (m *unMarshaldeviceBgpRouter) FromYaml(value string) error {
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

func (m *marshaldeviceBgpRouter) ToJson() (string, error) {
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

func (m *unMarshaldeviceBgpRouter) FromJson(value string) error {
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

func (obj *deviceBgpRouter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceBgpRouter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceBgpRouter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceBgpRouter) Clone() (DeviceBgpRouter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceBgpRouter()
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

func (obj *deviceBgpRouter) setNil() {
	obj.ipv4InterfacesHolder = nil
	obj.ipv6InterfacesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceBgpRouter is configuration for one or more IPv4 or IPv6 BGP peers.
type DeviceBgpRouter interface {
	Validation
	// msg marshals DeviceBgpRouter to protobuf object *otg.DeviceBgpRouter
	// and doesn't set defaults
	msg() *otg.DeviceBgpRouter
	// setMsg unmarshals DeviceBgpRouter from protobuf object *otg.DeviceBgpRouter
	// and doesn't set defaults
	setMsg(*otg.DeviceBgpRouter) DeviceBgpRouter
	// provides marshal interface
	Marshal() marshalDeviceBgpRouter
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceBgpRouter
	// validate validates DeviceBgpRouter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceBgpRouter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterId returns string, set in DeviceBgpRouter.
	RouterId() string
	// SetRouterId assigns string provided by user to DeviceBgpRouter
	SetRouterId(value string) DeviceBgpRouter
	// Ipv4Interfaces returns DeviceBgpRouterBgpV4InterfaceIterIter, set in DeviceBgpRouter
	Ipv4Interfaces() DeviceBgpRouterBgpV4InterfaceIter
	// Ipv6Interfaces returns DeviceBgpRouterBgpV6InterfaceIterIter, set in DeviceBgpRouter
	Ipv6Interfaces() DeviceBgpRouterBgpV6InterfaceIter
	setNil()
}

// The BGP router ID is a unique identifier used by BGP. It is a 32-bit value that is often represented by an IPv4 address.
// RouterId returns a string
func (obj *deviceBgpRouter) RouterId() string {

	return *obj.obj.RouterId

}

// The BGP router ID is a unique identifier used by BGP. It is a 32-bit value that is often represented by an IPv4 address.
// SetRouterId sets the string value in the DeviceBgpRouter object
func (obj *deviceBgpRouter) SetRouterId(value string) DeviceBgpRouter {

	obj.obj.RouterId = &value
	return obj
}

// This contains an array of references to IPv4 interfaces,  each of which will have list of peers to different destinations.
// Ipv4Interfaces returns a []BgpV4Interface
func (obj *deviceBgpRouter) Ipv4Interfaces() DeviceBgpRouterBgpV4InterfaceIter {
	if len(obj.obj.Ipv4Interfaces) == 0 {
		obj.obj.Ipv4Interfaces = []*otg.BgpV4Interface{}
	}
	if obj.ipv4InterfacesHolder == nil {
		obj.ipv4InterfacesHolder = newDeviceBgpRouterBgpV4InterfaceIter(&obj.obj.Ipv4Interfaces).setMsg(obj)
	}
	return obj.ipv4InterfacesHolder
}

type deviceBgpRouterBgpV4InterfaceIter struct {
	obj                 *deviceBgpRouter
	bgpV4InterfaceSlice []BgpV4Interface
	fieldPtr            *[]*otg.BgpV4Interface
}

func newDeviceBgpRouterBgpV4InterfaceIter(ptr *[]*otg.BgpV4Interface) DeviceBgpRouterBgpV4InterfaceIter {
	return &deviceBgpRouterBgpV4InterfaceIter{fieldPtr: ptr}
}

type DeviceBgpRouterBgpV4InterfaceIter interface {
	setMsg(*deviceBgpRouter) DeviceBgpRouterBgpV4InterfaceIter
	Items() []BgpV4Interface
	Add() BgpV4Interface
	Append(items ...BgpV4Interface) DeviceBgpRouterBgpV4InterfaceIter
	Set(index int, newObj BgpV4Interface) DeviceBgpRouterBgpV4InterfaceIter
	Clear() DeviceBgpRouterBgpV4InterfaceIter
	clearHolderSlice() DeviceBgpRouterBgpV4InterfaceIter
	appendHolderSlice(item BgpV4Interface) DeviceBgpRouterBgpV4InterfaceIter
}

func (obj *deviceBgpRouterBgpV4InterfaceIter) setMsg(msg *deviceBgpRouter) DeviceBgpRouterBgpV4InterfaceIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpV4Interface{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceBgpRouterBgpV4InterfaceIter) Items() []BgpV4Interface {
	return obj.bgpV4InterfaceSlice
}

func (obj *deviceBgpRouterBgpV4InterfaceIter) Add() BgpV4Interface {
	newObj := &otg.BgpV4Interface{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpV4Interface{obj: newObj}
	newLibObj.setDefault()
	obj.bgpV4InterfaceSlice = append(obj.bgpV4InterfaceSlice, newLibObj)
	return newLibObj
}

func (obj *deviceBgpRouterBgpV4InterfaceIter) Append(items ...BgpV4Interface) DeviceBgpRouterBgpV4InterfaceIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpV4InterfaceSlice = append(obj.bgpV4InterfaceSlice, item)
	}
	return obj
}

func (obj *deviceBgpRouterBgpV4InterfaceIter) Set(index int, newObj BgpV4Interface) DeviceBgpRouterBgpV4InterfaceIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpV4InterfaceSlice[index] = newObj
	return obj
}
func (obj *deviceBgpRouterBgpV4InterfaceIter) Clear() DeviceBgpRouterBgpV4InterfaceIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpV4Interface{}
		obj.bgpV4InterfaceSlice = []BgpV4Interface{}
	}
	return obj
}
func (obj *deviceBgpRouterBgpV4InterfaceIter) clearHolderSlice() DeviceBgpRouterBgpV4InterfaceIter {
	if len(obj.bgpV4InterfaceSlice) > 0 {
		obj.bgpV4InterfaceSlice = []BgpV4Interface{}
	}
	return obj
}
func (obj *deviceBgpRouterBgpV4InterfaceIter) appendHolderSlice(item BgpV4Interface) DeviceBgpRouterBgpV4InterfaceIter {
	obj.bgpV4InterfaceSlice = append(obj.bgpV4InterfaceSlice, item)
	return obj
}

// This contains an array of references to IPv6 interfaces,  each of which will have list of peers to different destinations.
// Ipv6Interfaces returns a []BgpV6Interface
func (obj *deviceBgpRouter) Ipv6Interfaces() DeviceBgpRouterBgpV6InterfaceIter {
	if len(obj.obj.Ipv6Interfaces) == 0 {
		obj.obj.Ipv6Interfaces = []*otg.BgpV6Interface{}
	}
	if obj.ipv6InterfacesHolder == nil {
		obj.ipv6InterfacesHolder = newDeviceBgpRouterBgpV6InterfaceIter(&obj.obj.Ipv6Interfaces).setMsg(obj)
	}
	return obj.ipv6InterfacesHolder
}

type deviceBgpRouterBgpV6InterfaceIter struct {
	obj                 *deviceBgpRouter
	bgpV6InterfaceSlice []BgpV6Interface
	fieldPtr            *[]*otg.BgpV6Interface
}

func newDeviceBgpRouterBgpV6InterfaceIter(ptr *[]*otg.BgpV6Interface) DeviceBgpRouterBgpV6InterfaceIter {
	return &deviceBgpRouterBgpV6InterfaceIter{fieldPtr: ptr}
}

type DeviceBgpRouterBgpV6InterfaceIter interface {
	setMsg(*deviceBgpRouter) DeviceBgpRouterBgpV6InterfaceIter
	Items() []BgpV6Interface
	Add() BgpV6Interface
	Append(items ...BgpV6Interface) DeviceBgpRouterBgpV6InterfaceIter
	Set(index int, newObj BgpV6Interface) DeviceBgpRouterBgpV6InterfaceIter
	Clear() DeviceBgpRouterBgpV6InterfaceIter
	clearHolderSlice() DeviceBgpRouterBgpV6InterfaceIter
	appendHolderSlice(item BgpV6Interface) DeviceBgpRouterBgpV6InterfaceIter
}

func (obj *deviceBgpRouterBgpV6InterfaceIter) setMsg(msg *deviceBgpRouter) DeviceBgpRouterBgpV6InterfaceIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpV6Interface{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceBgpRouterBgpV6InterfaceIter) Items() []BgpV6Interface {
	return obj.bgpV6InterfaceSlice
}

func (obj *deviceBgpRouterBgpV6InterfaceIter) Add() BgpV6Interface {
	newObj := &otg.BgpV6Interface{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpV6Interface{obj: newObj}
	newLibObj.setDefault()
	obj.bgpV6InterfaceSlice = append(obj.bgpV6InterfaceSlice, newLibObj)
	return newLibObj
}

func (obj *deviceBgpRouterBgpV6InterfaceIter) Append(items ...BgpV6Interface) DeviceBgpRouterBgpV6InterfaceIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpV6InterfaceSlice = append(obj.bgpV6InterfaceSlice, item)
	}
	return obj
}

func (obj *deviceBgpRouterBgpV6InterfaceIter) Set(index int, newObj BgpV6Interface) DeviceBgpRouterBgpV6InterfaceIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpV6InterfaceSlice[index] = newObj
	return obj
}
func (obj *deviceBgpRouterBgpV6InterfaceIter) Clear() DeviceBgpRouterBgpV6InterfaceIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpV6Interface{}
		obj.bgpV6InterfaceSlice = []BgpV6Interface{}
	}
	return obj
}
func (obj *deviceBgpRouterBgpV6InterfaceIter) clearHolderSlice() DeviceBgpRouterBgpV6InterfaceIter {
	if len(obj.bgpV6InterfaceSlice) > 0 {
		obj.bgpV6InterfaceSlice = []BgpV6Interface{}
	}
	return obj
}
func (obj *deviceBgpRouterBgpV6InterfaceIter) appendHolderSlice(item BgpV6Interface) DeviceBgpRouterBgpV6InterfaceIter {
	obj.bgpV6InterfaceSlice = append(obj.bgpV6InterfaceSlice, item)
	return obj
}

func (obj *deviceBgpRouter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// RouterId is required
	if obj.obj.RouterId == nil {
		vObj.validationErrors = append(vObj.validationErrors, "RouterId is required field on interface DeviceBgpRouter")
	}
	if obj.obj.RouterId != nil {

		err := obj.validateIpv4(obj.RouterId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on DeviceBgpRouter.RouterId"))
		}

	}

	if len(obj.obj.Ipv4Interfaces) != 0 {

		if set_default {
			obj.Ipv4Interfaces().clearHolderSlice()
			for _, item := range obj.obj.Ipv4Interfaces {
				obj.Ipv4Interfaces().appendHolderSlice(&bgpV4Interface{obj: item})
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
				obj.Ipv6Interfaces().appendHolderSlice(&bgpV6Interface{obj: item})
			}
		}
		for _, item := range obj.Ipv6Interfaces().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *deviceBgpRouter) setDefault() {

}
