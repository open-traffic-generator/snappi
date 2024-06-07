package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** VxlanV6TunnelDestinationIPModeUnicastVtep *****
type vxlanV6TunnelDestinationIPModeUnicastVtep struct {
	validation
	obj                       *otg.VxlanV6TunnelDestinationIPModeUnicastVtep
	marshaller                marshalVxlanV6TunnelDestinationIPModeUnicastVtep
	unMarshaller              unMarshalVxlanV6TunnelDestinationIPModeUnicastVtep
	arpSuppressionCacheHolder VxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter
}

func NewVxlanV6TunnelDestinationIPModeUnicastVtep() VxlanV6TunnelDestinationIPModeUnicastVtep {
	obj := vxlanV6TunnelDestinationIPModeUnicastVtep{obj: &otg.VxlanV6TunnelDestinationIPModeUnicastVtep{}}
	obj.setDefault()
	return &obj
}

func (obj *vxlanV6TunnelDestinationIPModeUnicastVtep) msg() *otg.VxlanV6TunnelDestinationIPModeUnicastVtep {
	return obj.obj
}

func (obj *vxlanV6TunnelDestinationIPModeUnicastVtep) setMsg(msg *otg.VxlanV6TunnelDestinationIPModeUnicastVtep) VxlanV6TunnelDestinationIPModeUnicastVtep {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalvxlanV6TunnelDestinationIPModeUnicastVtep struct {
	obj *vxlanV6TunnelDestinationIPModeUnicastVtep
}

type marshalVxlanV6TunnelDestinationIPModeUnicastVtep interface {
	// ToProto marshals VxlanV6TunnelDestinationIPModeUnicastVtep to protobuf object *otg.VxlanV6TunnelDestinationIPModeUnicastVtep
	ToProto() (*otg.VxlanV6TunnelDestinationIPModeUnicastVtep, error)
	// ToPbText marshals VxlanV6TunnelDestinationIPModeUnicastVtep to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals VxlanV6TunnelDestinationIPModeUnicastVtep to YAML text
	ToYaml() (string, error)
	// ToJson marshals VxlanV6TunnelDestinationIPModeUnicastVtep to JSON text
	ToJson() (string, error)
}

type unMarshalvxlanV6TunnelDestinationIPModeUnicastVtep struct {
	obj *vxlanV6TunnelDestinationIPModeUnicastVtep
}

type unMarshalVxlanV6TunnelDestinationIPModeUnicastVtep interface {
	// FromProto unmarshals VxlanV6TunnelDestinationIPModeUnicastVtep from protobuf object *otg.VxlanV6TunnelDestinationIPModeUnicastVtep
	FromProto(msg *otg.VxlanV6TunnelDestinationIPModeUnicastVtep) (VxlanV6TunnelDestinationIPModeUnicastVtep, error)
	// FromPbText unmarshals VxlanV6TunnelDestinationIPModeUnicastVtep from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals VxlanV6TunnelDestinationIPModeUnicastVtep from YAML text
	FromYaml(value string) error
	// FromJson unmarshals VxlanV6TunnelDestinationIPModeUnicastVtep from JSON text
	FromJson(value string) error
}

func (obj *vxlanV6TunnelDestinationIPModeUnicastVtep) Marshal() marshalVxlanV6TunnelDestinationIPModeUnicastVtep {
	if obj.marshaller == nil {
		obj.marshaller = &marshalvxlanV6TunnelDestinationIPModeUnicastVtep{obj: obj}
	}
	return obj.marshaller
}

func (obj *vxlanV6TunnelDestinationIPModeUnicastVtep) Unmarshal() unMarshalVxlanV6TunnelDestinationIPModeUnicastVtep {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalvxlanV6TunnelDestinationIPModeUnicastVtep{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalvxlanV6TunnelDestinationIPModeUnicastVtep) ToProto() (*otg.VxlanV6TunnelDestinationIPModeUnicastVtep, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalvxlanV6TunnelDestinationIPModeUnicastVtep) FromProto(msg *otg.VxlanV6TunnelDestinationIPModeUnicastVtep) (VxlanV6TunnelDestinationIPModeUnicastVtep, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalvxlanV6TunnelDestinationIPModeUnicastVtep) ToPbText() (string, error) {
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

func (m *unMarshalvxlanV6TunnelDestinationIPModeUnicastVtep) FromPbText(value string) error {
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

func (m *marshalvxlanV6TunnelDestinationIPModeUnicastVtep) ToYaml() (string, error) {
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

func (m *unMarshalvxlanV6TunnelDestinationIPModeUnicastVtep) FromYaml(value string) error {
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

func (m *marshalvxlanV6TunnelDestinationIPModeUnicastVtep) ToJson() (string, error) {
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

func (m *unMarshalvxlanV6TunnelDestinationIPModeUnicastVtep) FromJson(value string) error {
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

func (obj *vxlanV6TunnelDestinationIPModeUnicastVtep) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *vxlanV6TunnelDestinationIPModeUnicastVtep) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *vxlanV6TunnelDestinationIPModeUnicastVtep) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *vxlanV6TunnelDestinationIPModeUnicastVtep) Clone() (VxlanV6TunnelDestinationIPModeUnicastVtep, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewVxlanV6TunnelDestinationIPModeUnicastVtep()
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

func (obj *vxlanV6TunnelDestinationIPModeUnicastVtep) setNil() {
	obj.arpSuppressionCacheHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// VxlanV6TunnelDestinationIPModeUnicastVtep is vTEP (VXLAN Tunnel End Point (VTEP)) parameters
type VxlanV6TunnelDestinationIPModeUnicastVtep interface {
	Validation
	// msg marshals VxlanV6TunnelDestinationIPModeUnicastVtep to protobuf object *otg.VxlanV6TunnelDestinationIPModeUnicastVtep
	// and doesn't set defaults
	msg() *otg.VxlanV6TunnelDestinationIPModeUnicastVtep
	// setMsg unmarshals VxlanV6TunnelDestinationIPModeUnicastVtep from protobuf object *otg.VxlanV6TunnelDestinationIPModeUnicastVtep
	// and doesn't set defaults
	setMsg(*otg.VxlanV6TunnelDestinationIPModeUnicastVtep) VxlanV6TunnelDestinationIPModeUnicastVtep
	// provides marshal interface
	Marshal() marshalVxlanV6TunnelDestinationIPModeUnicastVtep
	// provides unmarshal interface
	Unmarshal() unMarshalVxlanV6TunnelDestinationIPModeUnicastVtep
	// validate validates VxlanV6TunnelDestinationIPModeUnicastVtep
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (VxlanV6TunnelDestinationIPModeUnicastVtep, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RemoteVtepAddress returns string, set in VxlanV6TunnelDestinationIPModeUnicastVtep.
	RemoteVtepAddress() string
	// SetRemoteVtepAddress assigns string provided by user to VxlanV6TunnelDestinationIPModeUnicastVtep
	SetRemoteVtepAddress(value string) VxlanV6TunnelDestinationIPModeUnicastVtep
	// HasRemoteVtepAddress checks if RemoteVtepAddress has been set in VxlanV6TunnelDestinationIPModeUnicastVtep
	HasRemoteVtepAddress() bool
	// ArpSuppressionCache returns VxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIterIter, set in VxlanV6TunnelDestinationIPModeUnicastVtep
	ArpSuppressionCache() VxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter
	setNil()
}

// Remote VXLAN Tunnel End Point address
// RemoteVtepAddress returns a string
func (obj *vxlanV6TunnelDestinationIPModeUnicastVtep) RemoteVtepAddress() string {

	return *obj.obj.RemoteVtepAddress

}

// Remote VXLAN Tunnel End Point address
// RemoteVtepAddress returns a string
func (obj *vxlanV6TunnelDestinationIPModeUnicastVtep) HasRemoteVtepAddress() bool {
	return obj.obj.RemoteVtepAddress != nil
}

// Remote VXLAN Tunnel End Point address
// SetRemoteVtepAddress sets the string value in the VxlanV6TunnelDestinationIPModeUnicastVtep object
func (obj *vxlanV6TunnelDestinationIPModeUnicastVtep) SetRemoteVtepAddress(value string) VxlanV6TunnelDestinationIPModeUnicastVtep {

	obj.obj.RemoteVtepAddress = &value
	return obj
}

// Each VTEP maintains an ARP suppression cache table for known IP hosts and their associated MAC addresses in the VNI segment.  When an end host in the VNI sends an ARP request for another end-host IP address,  its local VTEP intercepts the ARP request and checks for the ARP-resolved IP address in its ARP suppression cache table.  If it finds a match, the local VTEP sends an ARP response on behalf of the remote end host.
// ArpSuppressionCache returns a []VxlanTunnelDestinationIPModeUnicastArpSuppressionCache
func (obj *vxlanV6TunnelDestinationIPModeUnicastVtep) ArpSuppressionCache() VxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter {
	if len(obj.obj.ArpSuppressionCache) == 0 {
		obj.obj.ArpSuppressionCache = []*otg.VxlanTunnelDestinationIPModeUnicastArpSuppressionCache{}
	}
	if obj.arpSuppressionCacheHolder == nil {
		obj.arpSuppressionCacheHolder = newVxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter(&obj.obj.ArpSuppressionCache).setMsg(obj)
	}
	return obj.arpSuppressionCacheHolder
}

type vxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter struct {
	obj                                                         *vxlanV6TunnelDestinationIPModeUnicastVtep
	vxlanTunnelDestinationIPModeUnicastArpSuppressionCacheSlice []VxlanTunnelDestinationIPModeUnicastArpSuppressionCache
	fieldPtr                                                    *[]*otg.VxlanTunnelDestinationIPModeUnicastArpSuppressionCache
}

func newVxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter(ptr *[]*otg.VxlanTunnelDestinationIPModeUnicastArpSuppressionCache) VxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter {
	return &vxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter{fieldPtr: ptr}
}

type VxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter interface {
	setMsg(*vxlanV6TunnelDestinationIPModeUnicastVtep) VxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter
	Items() []VxlanTunnelDestinationIPModeUnicastArpSuppressionCache
	Add() VxlanTunnelDestinationIPModeUnicastArpSuppressionCache
	Append(items ...VxlanTunnelDestinationIPModeUnicastArpSuppressionCache) VxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter
	Set(index int, newObj VxlanTunnelDestinationIPModeUnicastArpSuppressionCache) VxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter
	Clear() VxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter
	clearHolderSlice() VxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter
	appendHolderSlice(item VxlanTunnelDestinationIPModeUnicastArpSuppressionCache) VxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter
}

func (obj *vxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter) setMsg(msg *vxlanV6TunnelDestinationIPModeUnicastVtep) VxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&vxlanTunnelDestinationIPModeUnicastArpSuppressionCache{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *vxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter) Items() []VxlanTunnelDestinationIPModeUnicastArpSuppressionCache {
	return obj.vxlanTunnelDestinationIPModeUnicastArpSuppressionCacheSlice
}

func (obj *vxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter) Add() VxlanTunnelDestinationIPModeUnicastArpSuppressionCache {
	newObj := &otg.VxlanTunnelDestinationIPModeUnicastArpSuppressionCache{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &vxlanTunnelDestinationIPModeUnicastArpSuppressionCache{obj: newObj}
	newLibObj.setDefault()
	obj.vxlanTunnelDestinationIPModeUnicastArpSuppressionCacheSlice = append(obj.vxlanTunnelDestinationIPModeUnicastArpSuppressionCacheSlice, newLibObj)
	return newLibObj
}

func (obj *vxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter) Append(items ...VxlanTunnelDestinationIPModeUnicastArpSuppressionCache) VxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.vxlanTunnelDestinationIPModeUnicastArpSuppressionCacheSlice = append(obj.vxlanTunnelDestinationIPModeUnicastArpSuppressionCacheSlice, item)
	}
	return obj
}

func (obj *vxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter) Set(index int, newObj VxlanTunnelDestinationIPModeUnicastArpSuppressionCache) VxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.vxlanTunnelDestinationIPModeUnicastArpSuppressionCacheSlice[index] = newObj
	return obj
}
func (obj *vxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter) Clear() VxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.VxlanTunnelDestinationIPModeUnicastArpSuppressionCache{}
		obj.vxlanTunnelDestinationIPModeUnicastArpSuppressionCacheSlice = []VxlanTunnelDestinationIPModeUnicastArpSuppressionCache{}
	}
	return obj
}
func (obj *vxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter) clearHolderSlice() VxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter {
	if len(obj.vxlanTunnelDestinationIPModeUnicastArpSuppressionCacheSlice) > 0 {
		obj.vxlanTunnelDestinationIPModeUnicastArpSuppressionCacheSlice = []VxlanTunnelDestinationIPModeUnicastArpSuppressionCache{}
	}
	return obj
}
func (obj *vxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter) appendHolderSlice(item VxlanTunnelDestinationIPModeUnicastArpSuppressionCache) VxlanV6TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter {
	obj.vxlanTunnelDestinationIPModeUnicastArpSuppressionCacheSlice = append(obj.vxlanTunnelDestinationIPModeUnicastArpSuppressionCacheSlice, item)
	return obj
}

func (obj *vxlanV6TunnelDestinationIPModeUnicastVtep) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RemoteVtepAddress != nil {

		err := obj.validateIpv6(obj.RemoteVtepAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on VxlanV6TunnelDestinationIPModeUnicastVtep.RemoteVtepAddress"))
		}

	}

	if len(obj.obj.ArpSuppressionCache) != 0 {

		if set_default {
			obj.ArpSuppressionCache().clearHolderSlice()
			for _, item := range obj.obj.ArpSuppressionCache {
				obj.ArpSuppressionCache().appendHolderSlice(&vxlanTunnelDestinationIPModeUnicastArpSuppressionCache{obj: item})
			}
		}
		for _, item := range obj.ArpSuppressionCache().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *vxlanV6TunnelDestinationIPModeUnicastVtep) setDefault() {

}
