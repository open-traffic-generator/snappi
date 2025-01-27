package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** VxlanV4TunnelDestinationIPModeUnicastVtep *****
type vxlanV4TunnelDestinationIPModeUnicastVtep struct {
	validation
	obj                       *otg.VxlanV4TunnelDestinationIPModeUnicastVtep
	marshaller                marshalVxlanV4TunnelDestinationIPModeUnicastVtep
	unMarshaller              unMarshalVxlanV4TunnelDestinationIPModeUnicastVtep
	arpSuppressionCacheHolder VxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter
}

func NewVxlanV4TunnelDestinationIPModeUnicastVtep() VxlanV4TunnelDestinationIPModeUnicastVtep {
	obj := vxlanV4TunnelDestinationIPModeUnicastVtep{obj: &otg.VxlanV4TunnelDestinationIPModeUnicastVtep{}}
	obj.setDefault()
	return &obj
}

func (obj *vxlanV4TunnelDestinationIPModeUnicastVtep) msg() *otg.VxlanV4TunnelDestinationIPModeUnicastVtep {
	return obj.obj
}

func (obj *vxlanV4TunnelDestinationIPModeUnicastVtep) setMsg(msg *otg.VxlanV4TunnelDestinationIPModeUnicastVtep) VxlanV4TunnelDestinationIPModeUnicastVtep {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalvxlanV4TunnelDestinationIPModeUnicastVtep struct {
	obj *vxlanV4TunnelDestinationIPModeUnicastVtep
}

type marshalVxlanV4TunnelDestinationIPModeUnicastVtep interface {
	// ToProto marshals VxlanV4TunnelDestinationIPModeUnicastVtep to protobuf object *otg.VxlanV4TunnelDestinationIPModeUnicastVtep
	ToProto() (*otg.VxlanV4TunnelDestinationIPModeUnicastVtep, error)
	// ToPbText marshals VxlanV4TunnelDestinationIPModeUnicastVtep to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals VxlanV4TunnelDestinationIPModeUnicastVtep to YAML text
	ToYaml() (string, error)
	// ToJson marshals VxlanV4TunnelDestinationIPModeUnicastVtep to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals VxlanV4TunnelDestinationIPModeUnicastVtep to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalvxlanV4TunnelDestinationIPModeUnicastVtep struct {
	obj *vxlanV4TunnelDestinationIPModeUnicastVtep
}

type unMarshalVxlanV4TunnelDestinationIPModeUnicastVtep interface {
	// FromProto unmarshals VxlanV4TunnelDestinationIPModeUnicastVtep from protobuf object *otg.VxlanV4TunnelDestinationIPModeUnicastVtep
	FromProto(msg *otg.VxlanV4TunnelDestinationIPModeUnicastVtep) (VxlanV4TunnelDestinationIPModeUnicastVtep, error)
	// FromPbText unmarshals VxlanV4TunnelDestinationIPModeUnicastVtep from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals VxlanV4TunnelDestinationIPModeUnicastVtep from YAML text
	FromYaml(value string) error
	// FromJson unmarshals VxlanV4TunnelDestinationIPModeUnicastVtep from JSON text
	FromJson(value string) error
}

func (obj *vxlanV4TunnelDestinationIPModeUnicastVtep) Marshal() marshalVxlanV4TunnelDestinationIPModeUnicastVtep {
	if obj.marshaller == nil {
		obj.marshaller = &marshalvxlanV4TunnelDestinationIPModeUnicastVtep{obj: obj}
	}
	return obj.marshaller
}

func (obj *vxlanV4TunnelDestinationIPModeUnicastVtep) Unmarshal() unMarshalVxlanV4TunnelDestinationIPModeUnicastVtep {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalvxlanV4TunnelDestinationIPModeUnicastVtep{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalvxlanV4TunnelDestinationIPModeUnicastVtep) ToProto() (*otg.VxlanV4TunnelDestinationIPModeUnicastVtep, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalvxlanV4TunnelDestinationIPModeUnicastVtep) FromProto(msg *otg.VxlanV4TunnelDestinationIPModeUnicastVtep) (VxlanV4TunnelDestinationIPModeUnicastVtep, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalvxlanV4TunnelDestinationIPModeUnicastVtep) ToPbText() (string, error) {
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

func (m *unMarshalvxlanV4TunnelDestinationIPModeUnicastVtep) FromPbText(value string) error {
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

func (m *marshalvxlanV4TunnelDestinationIPModeUnicastVtep) ToYaml() (string, error) {
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

func (m *unMarshalvxlanV4TunnelDestinationIPModeUnicastVtep) FromYaml(value string) error {
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

func (m *marshalvxlanV4TunnelDestinationIPModeUnicastVtep) ToJsonRaw() (string, error) {
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

func (m *marshalvxlanV4TunnelDestinationIPModeUnicastVtep) ToJson() (string, error) {
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

func (m *unMarshalvxlanV4TunnelDestinationIPModeUnicastVtep) FromJson(value string) error {
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

func (obj *vxlanV4TunnelDestinationIPModeUnicastVtep) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *vxlanV4TunnelDestinationIPModeUnicastVtep) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *vxlanV4TunnelDestinationIPModeUnicastVtep) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *vxlanV4TunnelDestinationIPModeUnicastVtep) Clone() (VxlanV4TunnelDestinationIPModeUnicastVtep, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewVxlanV4TunnelDestinationIPModeUnicastVtep()
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

func (obj *vxlanV4TunnelDestinationIPModeUnicastVtep) setNil() {
	obj.arpSuppressionCacheHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// VxlanV4TunnelDestinationIPModeUnicastVtep is vTEP (VXLAN Tunnel End Point (VTEP)) parameters
type VxlanV4TunnelDestinationIPModeUnicastVtep interface {
	Validation
	// msg marshals VxlanV4TunnelDestinationIPModeUnicastVtep to protobuf object *otg.VxlanV4TunnelDestinationIPModeUnicastVtep
	// and doesn't set defaults
	msg() *otg.VxlanV4TunnelDestinationIPModeUnicastVtep
	// setMsg unmarshals VxlanV4TunnelDestinationIPModeUnicastVtep from protobuf object *otg.VxlanV4TunnelDestinationIPModeUnicastVtep
	// and doesn't set defaults
	setMsg(*otg.VxlanV4TunnelDestinationIPModeUnicastVtep) VxlanV4TunnelDestinationIPModeUnicastVtep
	// provides marshal interface
	Marshal() marshalVxlanV4TunnelDestinationIPModeUnicastVtep
	// provides unmarshal interface
	Unmarshal() unMarshalVxlanV4TunnelDestinationIPModeUnicastVtep
	// validate validates VxlanV4TunnelDestinationIPModeUnicastVtep
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (VxlanV4TunnelDestinationIPModeUnicastVtep, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RemoteVtepAddress returns string, set in VxlanV4TunnelDestinationIPModeUnicastVtep.
	RemoteVtepAddress() string
	// SetRemoteVtepAddress assigns string provided by user to VxlanV4TunnelDestinationIPModeUnicastVtep
	SetRemoteVtepAddress(value string) VxlanV4TunnelDestinationIPModeUnicastVtep
	// HasRemoteVtepAddress checks if RemoteVtepAddress has been set in VxlanV4TunnelDestinationIPModeUnicastVtep
	HasRemoteVtepAddress() bool
	// ArpSuppressionCache returns VxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIterIter, set in VxlanV4TunnelDestinationIPModeUnicastVtep
	ArpSuppressionCache() VxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter
	setNil()
}

// Remote VXLAN Tunnel End Point address
// RemoteVtepAddress returns a string
func (obj *vxlanV4TunnelDestinationIPModeUnicastVtep) RemoteVtepAddress() string {

	return *obj.obj.RemoteVtepAddress

}

// Remote VXLAN Tunnel End Point address
// RemoteVtepAddress returns a string
func (obj *vxlanV4TunnelDestinationIPModeUnicastVtep) HasRemoteVtepAddress() bool {
	return obj.obj.RemoteVtepAddress != nil
}

// Remote VXLAN Tunnel End Point address
// SetRemoteVtepAddress sets the string value in the VxlanV4TunnelDestinationIPModeUnicastVtep object
func (obj *vxlanV4TunnelDestinationIPModeUnicastVtep) SetRemoteVtepAddress(value string) VxlanV4TunnelDestinationIPModeUnicastVtep {

	obj.obj.RemoteVtepAddress = &value
	return obj
}

// Each VTEP maintains an ARP suppression cache table for known IP hosts and their associated MAC addresses in the VNI segment.  When an end host in the VNI sends an ARP request for another end-host IP address,  its local VTEP intercepts the ARP request and checks for the ARP-resolved IP address in its ARP suppression cache table.  If it finds a match, the local VTEP sends an ARP response on behalf of the remote end host.
// ArpSuppressionCache returns a []VxlanTunnelDestinationIPModeUnicastArpSuppressionCache
func (obj *vxlanV4TunnelDestinationIPModeUnicastVtep) ArpSuppressionCache() VxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter {
	if len(obj.obj.ArpSuppressionCache) == 0 {
		obj.obj.ArpSuppressionCache = []*otg.VxlanTunnelDestinationIPModeUnicastArpSuppressionCache{}
	}
	if obj.arpSuppressionCacheHolder == nil {
		obj.arpSuppressionCacheHolder = newVxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter(&obj.obj.ArpSuppressionCache).setMsg(obj)
	}
	return obj.arpSuppressionCacheHolder
}

type vxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter struct {
	obj                                                         *vxlanV4TunnelDestinationIPModeUnicastVtep
	vxlanTunnelDestinationIPModeUnicastArpSuppressionCacheSlice []VxlanTunnelDestinationIPModeUnicastArpSuppressionCache
	fieldPtr                                                    *[]*otg.VxlanTunnelDestinationIPModeUnicastArpSuppressionCache
}

func newVxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter(ptr *[]*otg.VxlanTunnelDestinationIPModeUnicastArpSuppressionCache) VxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter {
	return &vxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter{fieldPtr: ptr}
}

type VxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter interface {
	setMsg(*vxlanV4TunnelDestinationIPModeUnicastVtep) VxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter
	Items() []VxlanTunnelDestinationIPModeUnicastArpSuppressionCache
	Add() VxlanTunnelDestinationIPModeUnicastArpSuppressionCache
	Append(items ...VxlanTunnelDestinationIPModeUnicastArpSuppressionCache) VxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter
	Set(index int, newObj VxlanTunnelDestinationIPModeUnicastArpSuppressionCache) VxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter
	Clear() VxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter
	clearHolderSlice() VxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter
	appendHolderSlice(item VxlanTunnelDestinationIPModeUnicastArpSuppressionCache) VxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter
}

func (obj *vxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter) setMsg(msg *vxlanV4TunnelDestinationIPModeUnicastVtep) VxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&vxlanTunnelDestinationIPModeUnicastArpSuppressionCache{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *vxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter) Items() []VxlanTunnelDestinationIPModeUnicastArpSuppressionCache {
	return obj.vxlanTunnelDestinationIPModeUnicastArpSuppressionCacheSlice
}

func (obj *vxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter) Add() VxlanTunnelDestinationIPModeUnicastArpSuppressionCache {
	newObj := &otg.VxlanTunnelDestinationIPModeUnicastArpSuppressionCache{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &vxlanTunnelDestinationIPModeUnicastArpSuppressionCache{obj: newObj}
	newLibObj.setDefault()
	obj.vxlanTunnelDestinationIPModeUnicastArpSuppressionCacheSlice = append(obj.vxlanTunnelDestinationIPModeUnicastArpSuppressionCacheSlice, newLibObj)
	return newLibObj
}

func (obj *vxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter) Append(items ...VxlanTunnelDestinationIPModeUnicastArpSuppressionCache) VxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.vxlanTunnelDestinationIPModeUnicastArpSuppressionCacheSlice = append(obj.vxlanTunnelDestinationIPModeUnicastArpSuppressionCacheSlice, item)
	}
	return obj
}

func (obj *vxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter) Set(index int, newObj VxlanTunnelDestinationIPModeUnicastArpSuppressionCache) VxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.vxlanTunnelDestinationIPModeUnicastArpSuppressionCacheSlice[index] = newObj
	return obj
}
func (obj *vxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter) Clear() VxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.VxlanTunnelDestinationIPModeUnicastArpSuppressionCache{}
		obj.vxlanTunnelDestinationIPModeUnicastArpSuppressionCacheSlice = []VxlanTunnelDestinationIPModeUnicastArpSuppressionCache{}
	}
	return obj
}
func (obj *vxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter) clearHolderSlice() VxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter {
	if len(obj.vxlanTunnelDestinationIPModeUnicastArpSuppressionCacheSlice) > 0 {
		obj.vxlanTunnelDestinationIPModeUnicastArpSuppressionCacheSlice = []VxlanTunnelDestinationIPModeUnicastArpSuppressionCache{}
	}
	return obj
}
func (obj *vxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter) appendHolderSlice(item VxlanTunnelDestinationIPModeUnicastArpSuppressionCache) VxlanV4TunnelDestinationIPModeUnicastVtepVxlanTunnelDestinationIPModeUnicastArpSuppressionCacheIter {
	obj.vxlanTunnelDestinationIPModeUnicastArpSuppressionCacheSlice = append(obj.vxlanTunnelDestinationIPModeUnicastArpSuppressionCacheSlice, item)
	return obj
}

func (obj *vxlanV4TunnelDestinationIPModeUnicastVtep) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RemoteVtepAddress != nil {

		err := obj.validateIpv4(obj.RemoteVtepAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on VxlanV4TunnelDestinationIPModeUnicastVtep.RemoteVtepAddress"))
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

func (obj *vxlanV4TunnelDestinationIPModeUnicastVtep) setDefault() {

}
