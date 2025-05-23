package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** VxlanTunnelDestinationIPModeUnicastArpSuppressionCache *****
type vxlanTunnelDestinationIPModeUnicastArpSuppressionCache struct {
	validation
	obj          *otg.VxlanTunnelDestinationIPModeUnicastArpSuppressionCache
	marshaller   marshalVxlanTunnelDestinationIPModeUnicastArpSuppressionCache
	unMarshaller unMarshalVxlanTunnelDestinationIPModeUnicastArpSuppressionCache
}

func NewVxlanTunnelDestinationIPModeUnicastArpSuppressionCache() VxlanTunnelDestinationIPModeUnicastArpSuppressionCache {
	obj := vxlanTunnelDestinationIPModeUnicastArpSuppressionCache{obj: &otg.VxlanTunnelDestinationIPModeUnicastArpSuppressionCache{}}
	obj.setDefault()
	return &obj
}

func (obj *vxlanTunnelDestinationIPModeUnicastArpSuppressionCache) msg() *otg.VxlanTunnelDestinationIPModeUnicastArpSuppressionCache {
	return obj.obj
}

func (obj *vxlanTunnelDestinationIPModeUnicastArpSuppressionCache) setMsg(msg *otg.VxlanTunnelDestinationIPModeUnicastArpSuppressionCache) VxlanTunnelDestinationIPModeUnicastArpSuppressionCache {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalvxlanTunnelDestinationIPModeUnicastArpSuppressionCache struct {
	obj *vxlanTunnelDestinationIPModeUnicastArpSuppressionCache
}

type marshalVxlanTunnelDestinationIPModeUnicastArpSuppressionCache interface {
	// ToProto marshals VxlanTunnelDestinationIPModeUnicastArpSuppressionCache to protobuf object *otg.VxlanTunnelDestinationIPModeUnicastArpSuppressionCache
	ToProto() (*otg.VxlanTunnelDestinationIPModeUnicastArpSuppressionCache, error)
	// ToPbText marshals VxlanTunnelDestinationIPModeUnicastArpSuppressionCache to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals VxlanTunnelDestinationIPModeUnicastArpSuppressionCache to YAML text
	ToYaml() (string, error)
	// ToJson marshals VxlanTunnelDestinationIPModeUnicastArpSuppressionCache to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals VxlanTunnelDestinationIPModeUnicastArpSuppressionCache to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalvxlanTunnelDestinationIPModeUnicastArpSuppressionCache struct {
	obj *vxlanTunnelDestinationIPModeUnicastArpSuppressionCache
}

type unMarshalVxlanTunnelDestinationIPModeUnicastArpSuppressionCache interface {
	// FromProto unmarshals VxlanTunnelDestinationIPModeUnicastArpSuppressionCache from protobuf object *otg.VxlanTunnelDestinationIPModeUnicastArpSuppressionCache
	FromProto(msg *otg.VxlanTunnelDestinationIPModeUnicastArpSuppressionCache) (VxlanTunnelDestinationIPModeUnicastArpSuppressionCache, error)
	// FromPbText unmarshals VxlanTunnelDestinationIPModeUnicastArpSuppressionCache from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals VxlanTunnelDestinationIPModeUnicastArpSuppressionCache from YAML text
	FromYaml(value string) error
	// FromJson unmarshals VxlanTunnelDestinationIPModeUnicastArpSuppressionCache from JSON text
	FromJson(value string) error
}

func (obj *vxlanTunnelDestinationIPModeUnicastArpSuppressionCache) Marshal() marshalVxlanTunnelDestinationIPModeUnicastArpSuppressionCache {
	if obj.marshaller == nil {
		obj.marshaller = &marshalvxlanTunnelDestinationIPModeUnicastArpSuppressionCache{obj: obj}
	}
	return obj.marshaller
}

func (obj *vxlanTunnelDestinationIPModeUnicastArpSuppressionCache) Unmarshal() unMarshalVxlanTunnelDestinationIPModeUnicastArpSuppressionCache {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalvxlanTunnelDestinationIPModeUnicastArpSuppressionCache{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalvxlanTunnelDestinationIPModeUnicastArpSuppressionCache) ToProto() (*otg.VxlanTunnelDestinationIPModeUnicastArpSuppressionCache, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalvxlanTunnelDestinationIPModeUnicastArpSuppressionCache) FromProto(msg *otg.VxlanTunnelDestinationIPModeUnicastArpSuppressionCache) (VxlanTunnelDestinationIPModeUnicastArpSuppressionCache, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalvxlanTunnelDestinationIPModeUnicastArpSuppressionCache) ToPbText() (string, error) {
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

func (m *unMarshalvxlanTunnelDestinationIPModeUnicastArpSuppressionCache) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalvxlanTunnelDestinationIPModeUnicastArpSuppressionCache) ToYaml() (string, error) {
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

func (m *unMarshalvxlanTunnelDestinationIPModeUnicastArpSuppressionCache) FromYaml(value string) error {
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

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalvxlanTunnelDestinationIPModeUnicastArpSuppressionCache) ToJsonRaw() (string, error) {
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

func (m *marshalvxlanTunnelDestinationIPModeUnicastArpSuppressionCache) ToJson() (string, error) {
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

func (m *unMarshalvxlanTunnelDestinationIPModeUnicastArpSuppressionCache) FromJson(value string) error {
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

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *vxlanTunnelDestinationIPModeUnicastArpSuppressionCache) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *vxlanTunnelDestinationIPModeUnicastArpSuppressionCache) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *vxlanTunnelDestinationIPModeUnicastArpSuppressionCache) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *vxlanTunnelDestinationIPModeUnicastArpSuppressionCache) Clone() (VxlanTunnelDestinationIPModeUnicastArpSuppressionCache, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewVxlanTunnelDestinationIPModeUnicastArpSuppressionCache()
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

// VxlanTunnelDestinationIPModeUnicastArpSuppressionCache is each VTEP maintains an ARP suppression cache table for known IP hosts and their associated MAC addresses in the VNI segment.  When an end host in the VNI sends an ARP request for another end-host IP address,  its local VTEP intercepts the ARP request and checks for the ARP-resolved IP address in its ARP suppression cache table.  If it finds a match, the local VTEP sends an ARP response on behalf of the remote end host.
type VxlanTunnelDestinationIPModeUnicastArpSuppressionCache interface {
	Validation
	// msg marshals VxlanTunnelDestinationIPModeUnicastArpSuppressionCache to protobuf object *otg.VxlanTunnelDestinationIPModeUnicastArpSuppressionCache
	// and doesn't set defaults
	msg() *otg.VxlanTunnelDestinationIPModeUnicastArpSuppressionCache
	// setMsg unmarshals VxlanTunnelDestinationIPModeUnicastArpSuppressionCache from protobuf object *otg.VxlanTunnelDestinationIPModeUnicastArpSuppressionCache
	// and doesn't set defaults
	setMsg(*otg.VxlanTunnelDestinationIPModeUnicastArpSuppressionCache) VxlanTunnelDestinationIPModeUnicastArpSuppressionCache
	// provides marshal interface
	Marshal() marshalVxlanTunnelDestinationIPModeUnicastArpSuppressionCache
	// provides unmarshal interface
	Unmarshal() unMarshalVxlanTunnelDestinationIPModeUnicastArpSuppressionCache
	// validate validates VxlanTunnelDestinationIPModeUnicastArpSuppressionCache
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (VxlanTunnelDestinationIPModeUnicastArpSuppressionCache, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RemoteVmMac returns string, set in VxlanTunnelDestinationIPModeUnicastArpSuppressionCache.
	RemoteVmMac() string
	// SetRemoteVmMac assigns string provided by user to VxlanTunnelDestinationIPModeUnicastArpSuppressionCache
	SetRemoteVmMac(value string) VxlanTunnelDestinationIPModeUnicastArpSuppressionCache
	// HasRemoteVmMac checks if RemoteVmMac has been set in VxlanTunnelDestinationIPModeUnicastArpSuppressionCache
	HasRemoteVmMac() bool
	// RemoteVmIpv4 returns string, set in VxlanTunnelDestinationIPModeUnicastArpSuppressionCache.
	RemoteVmIpv4() string
	// SetRemoteVmIpv4 assigns string provided by user to VxlanTunnelDestinationIPModeUnicastArpSuppressionCache
	SetRemoteVmIpv4(value string) VxlanTunnelDestinationIPModeUnicastArpSuppressionCache
	// HasRemoteVmIpv4 checks if RemoteVmIpv4 has been set in VxlanTunnelDestinationIPModeUnicastArpSuppressionCache
	HasRemoteVmIpv4() bool
}

// Remote VM MAC address bound to Remote VM IPv4 address
// RemoteVmMac returns a string
func (obj *vxlanTunnelDestinationIPModeUnicastArpSuppressionCache) RemoteVmMac() string {

	return *obj.obj.RemoteVmMac

}

// Remote VM MAC address bound to Remote VM IPv4 address
// RemoteVmMac returns a string
func (obj *vxlanTunnelDestinationIPModeUnicastArpSuppressionCache) HasRemoteVmMac() bool {
	return obj.obj.RemoteVmMac != nil
}

// Remote VM MAC address bound to Remote VM IPv4 address
// SetRemoteVmMac sets the string value in the VxlanTunnelDestinationIPModeUnicastArpSuppressionCache object
func (obj *vxlanTunnelDestinationIPModeUnicastArpSuppressionCache) SetRemoteVmMac(value string) VxlanTunnelDestinationIPModeUnicastArpSuppressionCache {

	obj.obj.RemoteVmMac = &value
	return obj
}

// Remote VM IPv4 address
// RemoteVmIpv4 returns a string
func (obj *vxlanTunnelDestinationIPModeUnicastArpSuppressionCache) RemoteVmIpv4() string {

	return *obj.obj.RemoteVmIpv4

}

// Remote VM IPv4 address
// RemoteVmIpv4 returns a string
func (obj *vxlanTunnelDestinationIPModeUnicastArpSuppressionCache) HasRemoteVmIpv4() bool {
	return obj.obj.RemoteVmIpv4 != nil
}

// Remote VM IPv4 address
// SetRemoteVmIpv4 sets the string value in the VxlanTunnelDestinationIPModeUnicastArpSuppressionCache object
func (obj *vxlanTunnelDestinationIPModeUnicastArpSuppressionCache) SetRemoteVmIpv4(value string) VxlanTunnelDestinationIPModeUnicastArpSuppressionCache {

	obj.obj.RemoteVmIpv4 = &value
	return obj
}

func (obj *vxlanTunnelDestinationIPModeUnicastArpSuppressionCache) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RemoteVmMac != nil {

		err := obj.validateMac(obj.RemoteVmMac())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on VxlanTunnelDestinationIPModeUnicastArpSuppressionCache.RemoteVmMac"))
		}

	}

	if obj.obj.RemoteVmIpv4 != nil {

		err := obj.validateIpv4(obj.RemoteVmIpv4())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on VxlanTunnelDestinationIPModeUnicastArpSuppressionCache.RemoteVmIpv4"))
		}

	}

}

func (obj *vxlanTunnelDestinationIPModeUnicastArpSuppressionCache) setDefault() {

}
