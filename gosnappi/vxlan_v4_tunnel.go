package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** VxlanV4Tunnel *****
type vxlanV4Tunnel struct {
	validation
	obj                     *otg.VxlanV4Tunnel
	marshaller              marshalVxlanV4Tunnel
	unMarshaller            unMarshalVxlanV4Tunnel
	destinationIpModeHolder VxlanV4TunnelDestinationIPMode
}

func NewVxlanV4Tunnel() VxlanV4Tunnel {
	obj := vxlanV4Tunnel{obj: &otg.VxlanV4Tunnel{}}
	obj.setDefault()
	return &obj
}

func (obj *vxlanV4Tunnel) msg() *otg.VxlanV4Tunnel {
	return obj.obj
}

func (obj *vxlanV4Tunnel) setMsg(msg *otg.VxlanV4Tunnel) VxlanV4Tunnel {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalvxlanV4Tunnel struct {
	obj *vxlanV4Tunnel
}

type marshalVxlanV4Tunnel interface {
	// ToProto marshals VxlanV4Tunnel to protobuf object *otg.VxlanV4Tunnel
	ToProto() (*otg.VxlanV4Tunnel, error)
	// ToPbText marshals VxlanV4Tunnel to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals VxlanV4Tunnel to YAML text
	ToYaml() (string, error)
	// ToJson marshals VxlanV4Tunnel to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals VxlanV4Tunnel to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalvxlanV4Tunnel struct {
	obj *vxlanV4Tunnel
}

type unMarshalVxlanV4Tunnel interface {
	// FromProto unmarshals VxlanV4Tunnel from protobuf object *otg.VxlanV4Tunnel
	FromProto(msg *otg.VxlanV4Tunnel) (VxlanV4Tunnel, error)
	// FromPbText unmarshals VxlanV4Tunnel from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals VxlanV4Tunnel from YAML text
	FromYaml(value string) error
	// FromJson unmarshals VxlanV4Tunnel from JSON text
	FromJson(value string) error
}

func (obj *vxlanV4Tunnel) Marshal() marshalVxlanV4Tunnel {
	if obj.marshaller == nil {
		obj.marshaller = &marshalvxlanV4Tunnel{obj: obj}
	}
	return obj.marshaller
}

func (obj *vxlanV4Tunnel) Unmarshal() unMarshalVxlanV4Tunnel {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalvxlanV4Tunnel{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalvxlanV4Tunnel) ToProto() (*otg.VxlanV4Tunnel, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalvxlanV4Tunnel) FromProto(msg *otg.VxlanV4Tunnel) (VxlanV4Tunnel, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalvxlanV4Tunnel) ToPbText() (string, error) {
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

func (m *unMarshalvxlanV4Tunnel) FromPbText(value string) error {
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

func (m *marshalvxlanV4Tunnel) ToYaml() (string, error) {
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

func (m *unMarshalvxlanV4Tunnel) FromYaml(value string) error {
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

func (m *marshalvxlanV4Tunnel) ToJsonRaw() (string, error) {
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

func (m *marshalvxlanV4Tunnel) ToJson() (string, error) {
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

func (m *unMarshalvxlanV4Tunnel) FromJson(value string) error {
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

func (obj *vxlanV4Tunnel) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *vxlanV4Tunnel) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *vxlanV4Tunnel) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *vxlanV4Tunnel) Clone() (VxlanV4Tunnel, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewVxlanV4Tunnel()
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

func (obj *vxlanV4Tunnel) setNil() {
	obj.destinationIpModeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// VxlanV4Tunnel is configuration and operational state parameters relating to IPv4 VXLAN tunnel end-point interface.
type VxlanV4Tunnel interface {
	Validation
	// msg marshals VxlanV4Tunnel to protobuf object *otg.VxlanV4Tunnel
	// and doesn't set defaults
	msg() *otg.VxlanV4Tunnel
	// setMsg unmarshals VxlanV4Tunnel from protobuf object *otg.VxlanV4Tunnel
	// and doesn't set defaults
	setMsg(*otg.VxlanV4Tunnel) VxlanV4Tunnel
	// provides marshal interface
	Marshal() marshalVxlanV4Tunnel
	// provides unmarshal interface
	Unmarshal() unMarshalVxlanV4Tunnel
	// validate validates VxlanV4Tunnel
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (VxlanV4Tunnel, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SourceInterface returns string, set in VxlanV4Tunnel.
	SourceInterface() string
	// SetSourceInterface assigns string provided by user to VxlanV4Tunnel
	SetSourceInterface(value string) VxlanV4Tunnel
	// DestinationIpMode returns VxlanV4TunnelDestinationIPMode, set in VxlanV4Tunnel.
	// VxlanV4TunnelDestinationIPMode is communication mode between the VTEPs, either unicast or multicast.
	DestinationIpMode() VxlanV4TunnelDestinationIPMode
	// SetDestinationIpMode assigns VxlanV4TunnelDestinationIPMode provided by user to VxlanV4Tunnel.
	// VxlanV4TunnelDestinationIPMode is communication mode between the VTEPs, either unicast or multicast.
	SetDestinationIpMode(value VxlanV4TunnelDestinationIPMode) VxlanV4Tunnel
	// HasDestinationIpMode checks if DestinationIpMode has been set in VxlanV4Tunnel
	HasDestinationIpMode() bool
	// Vni returns uint32, set in VxlanV4Tunnel.
	Vni() uint32
	// SetVni assigns uint32 provided by user to VxlanV4Tunnel
	SetVni(value uint32) VxlanV4Tunnel
	// Name returns string, set in VxlanV4Tunnel.
	Name() string
	// SetName assigns string provided by user to VxlanV4Tunnel
	SetName(value string) VxlanV4Tunnel
	setNil()
}

// Determines the source interface.
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
// - /components/schemas/Device.Ipv4Loopback/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
// - /components/schemas/Device.Ipv4Loopback/properties/name
//
// SourceInterface returns a string
func (obj *vxlanV4Tunnel) SourceInterface() string {

	return *obj.obj.SourceInterface

}

// Determines the source interface.
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
// - /components/schemas/Device.Ipv4Loopback/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
// - /components/schemas/Device.Ipv4Loopback/properties/name
//
// SetSourceInterface sets the string value in the VxlanV4Tunnel object
func (obj *vxlanV4Tunnel) SetSourceInterface(value string) VxlanV4Tunnel {

	obj.obj.SourceInterface = &value
	return obj
}

// description is TBD
// DestinationIpMode returns a VxlanV4TunnelDestinationIPMode
func (obj *vxlanV4Tunnel) DestinationIpMode() VxlanV4TunnelDestinationIPMode {
	if obj.obj.DestinationIpMode == nil {
		obj.obj.DestinationIpMode = NewVxlanV4TunnelDestinationIPMode().msg()
	}
	if obj.destinationIpModeHolder == nil {
		obj.destinationIpModeHolder = &vxlanV4TunnelDestinationIPMode{obj: obj.obj.DestinationIpMode}
	}
	return obj.destinationIpModeHolder
}

// description is TBD
// DestinationIpMode returns a VxlanV4TunnelDestinationIPMode
func (obj *vxlanV4Tunnel) HasDestinationIpMode() bool {
	return obj.obj.DestinationIpMode != nil
}

// description is TBD
// SetDestinationIpMode sets the VxlanV4TunnelDestinationIPMode value in the VxlanV4Tunnel object
func (obj *vxlanV4Tunnel) SetDestinationIpMode(value VxlanV4TunnelDestinationIPMode) VxlanV4Tunnel {

	obj.destinationIpModeHolder = nil
	obj.obj.DestinationIpMode = value.msg()

	return obj
}

// VXLAN Network Identifier (VNI) to distinguish network instances on the wire
// Vni returns a uint32
func (obj *vxlanV4Tunnel) Vni() uint32 {

	return *obj.obj.Vni

}

// VXLAN Network Identifier (VNI) to distinguish network instances on the wire
// SetVni sets the uint32 value in the VxlanV4Tunnel object
func (obj *vxlanV4Tunnel) SetVni(value uint32) VxlanV4Tunnel {

	obj.obj.Vni = &value
	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *vxlanV4Tunnel) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the VxlanV4Tunnel object
func (obj *vxlanV4Tunnel) SetName(value string) VxlanV4Tunnel {

	obj.obj.Name = &value
	return obj
}

func (obj *vxlanV4Tunnel) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// SourceInterface is required
	if obj.obj.SourceInterface == nil {
		vObj.validationErrors = append(vObj.validationErrors, "SourceInterface is required field on interface VxlanV4Tunnel")
	}

	if obj.obj.DestinationIpMode != nil {

		obj.DestinationIpMode().validateObj(vObj, set_default)
	}

	// Vni is required
	if obj.obj.Vni == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Vni is required field on interface VxlanV4Tunnel")
	}
	if obj.obj.Vni != nil {

		if *obj.obj.Vni < 1 || *obj.obj.Vni > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= VxlanV4Tunnel.Vni <= 16777215 but Got %d", *obj.obj.Vni))
		}

	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface VxlanV4Tunnel")
	}
}

func (obj *vxlanV4Tunnel) setDefault() {

}
