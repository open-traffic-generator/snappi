package gosnappi

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** VxlanV6Tunnel *****
type vxlanV6Tunnel struct {
	validation
	obj                     *otg.VxlanV6Tunnel
	marshaller              marshalVxlanV6Tunnel
	unMarshaller            unMarshalVxlanV6Tunnel
	destinationIpModeHolder VxlanV6TunnelDestinationIPMode
}

func NewVxlanV6Tunnel() VxlanV6Tunnel {
	obj := vxlanV6Tunnel{obj: &otg.VxlanV6Tunnel{}}
	obj.setDefault()
	return &obj
}

func (obj *vxlanV6Tunnel) msg() *otg.VxlanV6Tunnel {
	return obj.obj
}

func (obj *vxlanV6Tunnel) setMsg(msg *otg.VxlanV6Tunnel) VxlanV6Tunnel {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalvxlanV6Tunnel struct {
	obj *vxlanV6Tunnel
}

type marshalVxlanV6Tunnel interface {
	// ToProto marshals VxlanV6Tunnel to protobuf object *otg.VxlanV6Tunnel
	ToProto() (*otg.VxlanV6Tunnel, error)
	// ToPbText marshals VxlanV6Tunnel to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals VxlanV6Tunnel to YAML text
	ToYaml() (string, error)
	// ToJson marshals VxlanV6Tunnel to JSON text
	ToJson() (string, error)
}

type unMarshalvxlanV6Tunnel struct {
	obj *vxlanV6Tunnel
}

type unMarshalVxlanV6Tunnel interface {
	// FromProto unmarshals VxlanV6Tunnel from protobuf object *otg.VxlanV6Tunnel
	FromProto(msg *otg.VxlanV6Tunnel) (VxlanV6Tunnel, error)
	// FromPbText unmarshals VxlanV6Tunnel from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals VxlanV6Tunnel from YAML text
	FromYaml(value string) error
	// FromJson unmarshals VxlanV6Tunnel from JSON text
	FromJson(value string) error
}

func (obj *vxlanV6Tunnel) Marshal() marshalVxlanV6Tunnel {
	if obj.marshaller == nil {
		obj.marshaller = &marshalvxlanV6Tunnel{obj: obj}
	}
	return obj.marshaller
}

func (obj *vxlanV6Tunnel) Unmarshal() unMarshalVxlanV6Tunnel {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalvxlanV6Tunnel{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalvxlanV6Tunnel) ToProto() (*otg.VxlanV6Tunnel, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalvxlanV6Tunnel) FromProto(msg *otg.VxlanV6Tunnel) (VxlanV6Tunnel, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalvxlanV6Tunnel) ToPbText() (string, error) {
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

func (m *unMarshalvxlanV6Tunnel) FromPbText(value string) error {
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

func (m *marshalvxlanV6Tunnel) ToYaml() (string, error) {
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

func (m *unMarshalvxlanV6Tunnel) FromYaml(value string) error {
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

func (m *marshalvxlanV6Tunnel) ToJson() (string, error) {
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

func (m *unMarshalvxlanV6Tunnel) FromJson(value string) error {
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

func (obj *vxlanV6Tunnel) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *vxlanV6Tunnel) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *vxlanV6Tunnel) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *vxlanV6Tunnel) Clone() (VxlanV6Tunnel, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewVxlanV6Tunnel()
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

func (obj *vxlanV6Tunnel) setNil() {
	obj.destinationIpModeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// VxlanV6Tunnel is configuration and operational state parameters relating to IPv6 VXLAN tunnel end-point interface.
type VxlanV6Tunnel interface {
	Validation
	// msg marshals VxlanV6Tunnel to protobuf object *otg.VxlanV6Tunnel
	// and doesn't set defaults
	msg() *otg.VxlanV6Tunnel
	// setMsg unmarshals VxlanV6Tunnel from protobuf object *otg.VxlanV6Tunnel
	// and doesn't set defaults
	setMsg(*otg.VxlanV6Tunnel) VxlanV6Tunnel
	// provides marshal interface
	Marshal() marshalVxlanV6Tunnel
	// provides unmarshal interface
	Unmarshal() unMarshalVxlanV6Tunnel
	// validate validates VxlanV6Tunnel
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (VxlanV6Tunnel, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SourceInterface returns string, set in VxlanV6Tunnel.
	SourceInterface() string
	// SetSourceInterface assigns string provided by user to VxlanV6Tunnel
	SetSourceInterface(value string) VxlanV6Tunnel
	// DestinationIpMode returns VxlanV6TunnelDestinationIPMode, set in VxlanV6Tunnel.
	// VxlanV6TunnelDestinationIPMode is communication mode between the VTEPs, either unicast or multicast.
	DestinationIpMode() VxlanV6TunnelDestinationIPMode
	// SetDestinationIpMode assigns VxlanV6TunnelDestinationIPMode provided by user to VxlanV6Tunnel.
	// VxlanV6TunnelDestinationIPMode is communication mode between the VTEPs, either unicast or multicast.
	SetDestinationIpMode(value VxlanV6TunnelDestinationIPMode) VxlanV6Tunnel
	// HasDestinationIpMode checks if DestinationIpMode has been set in VxlanV6Tunnel
	HasDestinationIpMode() bool
	// Vni returns uint32, set in VxlanV6Tunnel.
	Vni() uint32
	// SetVni assigns uint32 provided by user to VxlanV6Tunnel
	SetVni(value uint32) VxlanV6Tunnel
	// Name returns string, set in VxlanV6Tunnel.
	Name() string
	// SetName assigns string provided by user to VxlanV6Tunnel
	SetName(value string) VxlanV6Tunnel
	setNil()
}

// Determines the source interface.
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
// - /components/schemas/Device.Ipv6Loopback/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
// - /components/schemas/Device.Ipv6Loopback/properties/name
//
// SourceInterface returns a string
func (obj *vxlanV6Tunnel) SourceInterface() string {

	return *obj.obj.SourceInterface

}

// Determines the source interface.
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
// - /components/schemas/Device.Ipv6Loopback/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
// - /components/schemas/Device.Ipv6Loopback/properties/name
//
// SetSourceInterface sets the string value in the VxlanV6Tunnel object
func (obj *vxlanV6Tunnel) SetSourceInterface(value string) VxlanV6Tunnel {

	obj.obj.SourceInterface = &value
	return obj
}

// description is TBD
// DestinationIpMode returns a VxlanV6TunnelDestinationIPMode
func (obj *vxlanV6Tunnel) DestinationIpMode() VxlanV6TunnelDestinationIPMode {
	if obj.obj.DestinationIpMode == nil {
		obj.obj.DestinationIpMode = NewVxlanV6TunnelDestinationIPMode().msg()
	}
	if obj.destinationIpModeHolder == nil {
		obj.destinationIpModeHolder = &vxlanV6TunnelDestinationIPMode{obj: obj.obj.DestinationIpMode}
	}
	return obj.destinationIpModeHolder
}

// description is TBD
// DestinationIpMode returns a VxlanV6TunnelDestinationIPMode
func (obj *vxlanV6Tunnel) HasDestinationIpMode() bool {
	return obj.obj.DestinationIpMode != nil
}

// description is TBD
// SetDestinationIpMode sets the VxlanV6TunnelDestinationIPMode value in the VxlanV6Tunnel object
func (obj *vxlanV6Tunnel) SetDestinationIpMode(value VxlanV6TunnelDestinationIPMode) VxlanV6Tunnel {

	obj.destinationIpModeHolder = nil
	obj.obj.DestinationIpMode = value.msg()

	return obj
}

// VXLAN Network Identifier (VNI) to distinguish network instances on the wire
// Vni returns a uint32
func (obj *vxlanV6Tunnel) Vni() uint32 {

	return *obj.obj.Vni

}

// VXLAN Network Identifier (VNI) to distinguish network instances on the wire
// SetVni sets the uint32 value in the VxlanV6Tunnel object
func (obj *vxlanV6Tunnel) SetVni(value uint32) VxlanV6Tunnel {

	obj.obj.Vni = &value
	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *vxlanV6Tunnel) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the VxlanV6Tunnel object
func (obj *vxlanV6Tunnel) SetName(value string) VxlanV6Tunnel {

	obj.obj.Name = &value
	return obj
}

func (obj *vxlanV6Tunnel) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// SourceInterface is required
	if obj.obj.SourceInterface == nil {
		vObj.validationErrors = append(vObj.validationErrors, "SourceInterface is required field on interface VxlanV6Tunnel")
	}

	if obj.obj.DestinationIpMode != nil {

		obj.DestinationIpMode().validateObj(vObj, set_default)
	}

	// Vni is required
	if obj.obj.Vni == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Vni is required field on interface VxlanV6Tunnel")
	}
	if obj.obj.Vni != nil {

		if *obj.obj.Vni < 1 || *obj.obj.Vni > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= VxlanV6Tunnel.Vni <= 16777215 but Got %d", *obj.obj.Vni))
		}

	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface VxlanV6Tunnel")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"VxlanV6Tunnel.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

}

func (obj *vxlanV6Tunnel) setDefault() {

}
