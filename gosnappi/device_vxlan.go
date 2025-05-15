package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceVxlan *****
type deviceVxlan struct {
	validation
	obj             *otg.DeviceVxlan
	marshaller      marshalDeviceVxlan
	unMarshaller    unMarshalDeviceVxlan
	v4TunnelsHolder DeviceVxlanVxlanV4TunnelIter
	v6TunnelsHolder DeviceVxlanVxlanV6TunnelIter
}

func NewDeviceVxlan() DeviceVxlan {
	obj := deviceVxlan{obj: &otg.DeviceVxlan{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceVxlan) msg() *otg.DeviceVxlan {
	return obj.obj
}

func (obj *deviceVxlan) setMsg(msg *otg.DeviceVxlan) DeviceVxlan {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceVxlan struct {
	obj *deviceVxlan
}

type marshalDeviceVxlan interface {
	// ToProto marshals DeviceVxlan to protobuf object *otg.DeviceVxlan
	ToProto() (*otg.DeviceVxlan, error)
	// ToPbText marshals DeviceVxlan to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceVxlan to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceVxlan to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceVxlan struct {
	obj *deviceVxlan
}

type unMarshalDeviceVxlan interface {
	// FromProto unmarshals DeviceVxlan from protobuf object *otg.DeviceVxlan
	FromProto(msg *otg.DeviceVxlan) (DeviceVxlan, error)
	// FromPbText unmarshals DeviceVxlan from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceVxlan from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceVxlan from JSON text
	FromJson(value string) error
}

func (obj *deviceVxlan) Marshal() marshalDeviceVxlan {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceVxlan{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceVxlan) Unmarshal() unMarshalDeviceVxlan {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceVxlan{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceVxlan) ToProto() (*otg.DeviceVxlan, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceVxlan) FromProto(msg *otg.DeviceVxlan) (DeviceVxlan, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceVxlan) ToPbText() (string, error) {
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

func (m *unMarshaldeviceVxlan) FromPbText(value string) error {
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

func (m *marshaldeviceVxlan) ToYaml() (string, error) {
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

func (m *unMarshaldeviceVxlan) FromYaml(value string) error {
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

func (m *marshaldeviceVxlan) ToJson() (string, error) {
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

func (m *unMarshaldeviceVxlan) FromJson(value string) error {
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

func (obj *deviceVxlan) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceVxlan) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceVxlan) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceVxlan) Clone() (DeviceVxlan, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceVxlan()
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

func (obj *deviceVxlan) setNil() {
	obj.v4TunnelsHolder = nil
	obj.v6TunnelsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceVxlan is description is TBD
type DeviceVxlan interface {
	Validation
	// msg marshals DeviceVxlan to protobuf object *otg.DeviceVxlan
	// and doesn't set defaults
	msg() *otg.DeviceVxlan
	// setMsg unmarshals DeviceVxlan from protobuf object *otg.DeviceVxlan
	// and doesn't set defaults
	setMsg(*otg.DeviceVxlan) DeviceVxlan
	// provides marshal interface
	Marshal() marshalDeviceVxlan
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceVxlan
	// validate validates DeviceVxlan
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceVxlan, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// V4Tunnels returns DeviceVxlanVxlanV4TunnelIterIter, set in DeviceVxlan
	V4Tunnels() DeviceVxlanVxlanV4TunnelIter
	// V6Tunnels returns DeviceVxlanVxlanV6TunnelIterIter, set in DeviceVxlan
	V6Tunnels() DeviceVxlanVxlanV6TunnelIter
	setNil()
}

// IPv4 VXLAN Tunnels
// V4Tunnels returns a []VxlanV4Tunnel
func (obj *deviceVxlan) V4Tunnels() DeviceVxlanVxlanV4TunnelIter {
	if len(obj.obj.V4Tunnels) == 0 {
		obj.obj.V4Tunnels = []*otg.VxlanV4Tunnel{}
	}
	if obj.v4TunnelsHolder == nil {
		obj.v4TunnelsHolder = newDeviceVxlanVxlanV4TunnelIter(&obj.obj.V4Tunnels).setMsg(obj)
	}
	return obj.v4TunnelsHolder
}

type deviceVxlanVxlanV4TunnelIter struct {
	obj                *deviceVxlan
	vxlanV4TunnelSlice []VxlanV4Tunnel
	fieldPtr           *[]*otg.VxlanV4Tunnel
}

func newDeviceVxlanVxlanV4TunnelIter(ptr *[]*otg.VxlanV4Tunnel) DeviceVxlanVxlanV4TunnelIter {
	return &deviceVxlanVxlanV4TunnelIter{fieldPtr: ptr}
}

type DeviceVxlanVxlanV4TunnelIter interface {
	setMsg(*deviceVxlan) DeviceVxlanVxlanV4TunnelIter
	Items() []VxlanV4Tunnel
	Add() VxlanV4Tunnel
	Append(items ...VxlanV4Tunnel) DeviceVxlanVxlanV4TunnelIter
	Set(index int, newObj VxlanV4Tunnel) DeviceVxlanVxlanV4TunnelIter
	Clear() DeviceVxlanVxlanV4TunnelIter
	clearHolderSlice() DeviceVxlanVxlanV4TunnelIter
	appendHolderSlice(item VxlanV4Tunnel) DeviceVxlanVxlanV4TunnelIter
}

func (obj *deviceVxlanVxlanV4TunnelIter) setMsg(msg *deviceVxlan) DeviceVxlanVxlanV4TunnelIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&vxlanV4Tunnel{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceVxlanVxlanV4TunnelIter) Items() []VxlanV4Tunnel {
	return obj.vxlanV4TunnelSlice
}

func (obj *deviceVxlanVxlanV4TunnelIter) Add() VxlanV4Tunnel {
	newObj := &otg.VxlanV4Tunnel{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &vxlanV4Tunnel{obj: newObj}
	newLibObj.setDefault()
	obj.vxlanV4TunnelSlice = append(obj.vxlanV4TunnelSlice, newLibObj)
	return newLibObj
}

func (obj *deviceVxlanVxlanV4TunnelIter) Append(items ...VxlanV4Tunnel) DeviceVxlanVxlanV4TunnelIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.vxlanV4TunnelSlice = append(obj.vxlanV4TunnelSlice, item)
	}
	return obj
}

func (obj *deviceVxlanVxlanV4TunnelIter) Set(index int, newObj VxlanV4Tunnel) DeviceVxlanVxlanV4TunnelIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.vxlanV4TunnelSlice[index] = newObj
	return obj
}
func (obj *deviceVxlanVxlanV4TunnelIter) Clear() DeviceVxlanVxlanV4TunnelIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.VxlanV4Tunnel{}
		obj.vxlanV4TunnelSlice = []VxlanV4Tunnel{}
	}
	return obj
}
func (obj *deviceVxlanVxlanV4TunnelIter) clearHolderSlice() DeviceVxlanVxlanV4TunnelIter {
	if len(obj.vxlanV4TunnelSlice) > 0 {
		obj.vxlanV4TunnelSlice = []VxlanV4Tunnel{}
	}
	return obj
}
func (obj *deviceVxlanVxlanV4TunnelIter) appendHolderSlice(item VxlanV4Tunnel) DeviceVxlanVxlanV4TunnelIter {
	obj.vxlanV4TunnelSlice = append(obj.vxlanV4TunnelSlice, item)
	return obj
}

// IPv6 VXLAN Tunnels
// V6Tunnels returns a []VxlanV6Tunnel
func (obj *deviceVxlan) V6Tunnels() DeviceVxlanVxlanV6TunnelIter {
	if len(obj.obj.V6Tunnels) == 0 {
		obj.obj.V6Tunnels = []*otg.VxlanV6Tunnel{}
	}
	if obj.v6TunnelsHolder == nil {
		obj.v6TunnelsHolder = newDeviceVxlanVxlanV6TunnelIter(&obj.obj.V6Tunnels).setMsg(obj)
	}
	return obj.v6TunnelsHolder
}

type deviceVxlanVxlanV6TunnelIter struct {
	obj                *deviceVxlan
	vxlanV6TunnelSlice []VxlanV6Tunnel
	fieldPtr           *[]*otg.VxlanV6Tunnel
}

func newDeviceVxlanVxlanV6TunnelIter(ptr *[]*otg.VxlanV6Tunnel) DeviceVxlanVxlanV6TunnelIter {
	return &deviceVxlanVxlanV6TunnelIter{fieldPtr: ptr}
}

type DeviceVxlanVxlanV6TunnelIter interface {
	setMsg(*deviceVxlan) DeviceVxlanVxlanV6TunnelIter
	Items() []VxlanV6Tunnel
	Add() VxlanV6Tunnel
	Append(items ...VxlanV6Tunnel) DeviceVxlanVxlanV6TunnelIter
	Set(index int, newObj VxlanV6Tunnel) DeviceVxlanVxlanV6TunnelIter
	Clear() DeviceVxlanVxlanV6TunnelIter
	clearHolderSlice() DeviceVxlanVxlanV6TunnelIter
	appendHolderSlice(item VxlanV6Tunnel) DeviceVxlanVxlanV6TunnelIter
}

func (obj *deviceVxlanVxlanV6TunnelIter) setMsg(msg *deviceVxlan) DeviceVxlanVxlanV6TunnelIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&vxlanV6Tunnel{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceVxlanVxlanV6TunnelIter) Items() []VxlanV6Tunnel {
	return obj.vxlanV6TunnelSlice
}

func (obj *deviceVxlanVxlanV6TunnelIter) Add() VxlanV6Tunnel {
	newObj := &otg.VxlanV6Tunnel{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &vxlanV6Tunnel{obj: newObj}
	newLibObj.setDefault()
	obj.vxlanV6TunnelSlice = append(obj.vxlanV6TunnelSlice, newLibObj)
	return newLibObj
}

func (obj *deviceVxlanVxlanV6TunnelIter) Append(items ...VxlanV6Tunnel) DeviceVxlanVxlanV6TunnelIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.vxlanV6TunnelSlice = append(obj.vxlanV6TunnelSlice, item)
	}
	return obj
}

func (obj *deviceVxlanVxlanV6TunnelIter) Set(index int, newObj VxlanV6Tunnel) DeviceVxlanVxlanV6TunnelIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.vxlanV6TunnelSlice[index] = newObj
	return obj
}
func (obj *deviceVxlanVxlanV6TunnelIter) Clear() DeviceVxlanVxlanV6TunnelIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.VxlanV6Tunnel{}
		obj.vxlanV6TunnelSlice = []VxlanV6Tunnel{}
	}
	return obj
}
func (obj *deviceVxlanVxlanV6TunnelIter) clearHolderSlice() DeviceVxlanVxlanV6TunnelIter {
	if len(obj.vxlanV6TunnelSlice) > 0 {
		obj.vxlanV6TunnelSlice = []VxlanV6Tunnel{}
	}
	return obj
}
func (obj *deviceVxlanVxlanV6TunnelIter) appendHolderSlice(item VxlanV6Tunnel) DeviceVxlanVxlanV6TunnelIter {
	obj.vxlanV6TunnelSlice = append(obj.vxlanV6TunnelSlice, item)
	return obj
}

func (obj *deviceVxlan) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.V4Tunnels) != 0 {

		if set_default {
			obj.V4Tunnels().clearHolderSlice()
			for _, item := range obj.obj.V4Tunnels {
				obj.V4Tunnels().appendHolderSlice(&vxlanV4Tunnel{obj: item})
			}
		}
		for _, item := range obj.V4Tunnels().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.V6Tunnels) != 0 {

		if set_default {
			obj.V6Tunnels().clearHolderSlice()
			for _, item := range obj.obj.V6Tunnels {
				obj.V6Tunnels().appendHolderSlice(&vxlanV6Tunnel{obj: item})
			}
		}
		for _, item := range obj.V6Tunnels().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *deviceVxlan) setDefault() {

}
