package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** VxlanV6TunnelDestinationIPModeMulticast *****
type vxlanV6TunnelDestinationIPModeMulticast struct {
	validation
	obj          *otg.VxlanV6TunnelDestinationIPModeMulticast
	marshaller   marshalVxlanV6TunnelDestinationIPModeMulticast
	unMarshaller unMarshalVxlanV6TunnelDestinationIPModeMulticast
}

func NewVxlanV6TunnelDestinationIPModeMulticast() VxlanV6TunnelDestinationIPModeMulticast {
	obj := vxlanV6TunnelDestinationIPModeMulticast{obj: &otg.VxlanV6TunnelDestinationIPModeMulticast{}}
	obj.setDefault()
	return &obj
}

func (obj *vxlanV6TunnelDestinationIPModeMulticast) msg() *otg.VxlanV6TunnelDestinationIPModeMulticast {
	return obj.obj
}

func (obj *vxlanV6TunnelDestinationIPModeMulticast) setMsg(msg *otg.VxlanV6TunnelDestinationIPModeMulticast) VxlanV6TunnelDestinationIPModeMulticast {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalvxlanV6TunnelDestinationIPModeMulticast struct {
	obj *vxlanV6TunnelDestinationIPModeMulticast
}

type marshalVxlanV6TunnelDestinationIPModeMulticast interface {
	// ToProto marshals VxlanV6TunnelDestinationIPModeMulticast to protobuf object *otg.VxlanV6TunnelDestinationIPModeMulticast
	ToProto() (*otg.VxlanV6TunnelDestinationIPModeMulticast, error)
	// ToPbText marshals VxlanV6TunnelDestinationIPModeMulticast to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals VxlanV6TunnelDestinationIPModeMulticast to YAML text
	ToYaml() (string, error)
	// ToJson marshals VxlanV6TunnelDestinationIPModeMulticast to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals VxlanV6TunnelDestinationIPModeMulticast to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalvxlanV6TunnelDestinationIPModeMulticast struct {
	obj *vxlanV6TunnelDestinationIPModeMulticast
}

type unMarshalVxlanV6TunnelDestinationIPModeMulticast interface {
	// FromProto unmarshals VxlanV6TunnelDestinationIPModeMulticast from protobuf object *otg.VxlanV6TunnelDestinationIPModeMulticast
	FromProto(msg *otg.VxlanV6TunnelDestinationIPModeMulticast) (VxlanV6TunnelDestinationIPModeMulticast, error)
	// FromPbText unmarshals VxlanV6TunnelDestinationIPModeMulticast from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals VxlanV6TunnelDestinationIPModeMulticast from YAML text
	FromYaml(value string) error
	// FromJson unmarshals VxlanV6TunnelDestinationIPModeMulticast from JSON text
	FromJson(value string) error
}

func (obj *vxlanV6TunnelDestinationIPModeMulticast) Marshal() marshalVxlanV6TunnelDestinationIPModeMulticast {
	if obj.marshaller == nil {
		obj.marshaller = &marshalvxlanV6TunnelDestinationIPModeMulticast{obj: obj}
	}
	return obj.marshaller
}

func (obj *vxlanV6TunnelDestinationIPModeMulticast) Unmarshal() unMarshalVxlanV6TunnelDestinationIPModeMulticast {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalvxlanV6TunnelDestinationIPModeMulticast{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalvxlanV6TunnelDestinationIPModeMulticast) ToProto() (*otg.VxlanV6TunnelDestinationIPModeMulticast, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalvxlanV6TunnelDestinationIPModeMulticast) FromProto(msg *otg.VxlanV6TunnelDestinationIPModeMulticast) (VxlanV6TunnelDestinationIPModeMulticast, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalvxlanV6TunnelDestinationIPModeMulticast) ToPbText() (string, error) {
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

func (m *unMarshalvxlanV6TunnelDestinationIPModeMulticast) FromPbText(value string) error {
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

func (m *marshalvxlanV6TunnelDestinationIPModeMulticast) ToYaml() (string, error) {
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

func (m *unMarshalvxlanV6TunnelDestinationIPModeMulticast) FromYaml(value string) error {
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

func (m *marshalvxlanV6TunnelDestinationIPModeMulticast) ToJsonRaw() (string, error) {
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

func (m *marshalvxlanV6TunnelDestinationIPModeMulticast) ToJson() (string, error) {
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

func (m *unMarshalvxlanV6TunnelDestinationIPModeMulticast) FromJson(value string) error {
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

func (obj *vxlanV6TunnelDestinationIPModeMulticast) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *vxlanV6TunnelDestinationIPModeMulticast) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *vxlanV6TunnelDestinationIPModeMulticast) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *vxlanV6TunnelDestinationIPModeMulticast) Clone() (VxlanV6TunnelDestinationIPModeMulticast, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewVxlanV6TunnelDestinationIPModeMulticast()
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

// VxlanV6TunnelDestinationIPModeMulticast is multicast Group address for member VNI(VXLAN Network Identifier)
type VxlanV6TunnelDestinationIPModeMulticast interface {
	Validation
	// msg marshals VxlanV6TunnelDestinationIPModeMulticast to protobuf object *otg.VxlanV6TunnelDestinationIPModeMulticast
	// and doesn't set defaults
	msg() *otg.VxlanV6TunnelDestinationIPModeMulticast
	// setMsg unmarshals VxlanV6TunnelDestinationIPModeMulticast from protobuf object *otg.VxlanV6TunnelDestinationIPModeMulticast
	// and doesn't set defaults
	setMsg(*otg.VxlanV6TunnelDestinationIPModeMulticast) VxlanV6TunnelDestinationIPModeMulticast
	// provides marshal interface
	Marshal() marshalVxlanV6TunnelDestinationIPModeMulticast
	// provides unmarshal interface
	Unmarshal() unMarshalVxlanV6TunnelDestinationIPModeMulticast
	// validate validates VxlanV6TunnelDestinationIPModeMulticast
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (VxlanV6TunnelDestinationIPModeMulticast, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Address returns string, set in VxlanV6TunnelDestinationIPModeMulticast.
	Address() string
	// SetAddress assigns string provided by user to VxlanV6TunnelDestinationIPModeMulticast
	SetAddress(value string) VxlanV6TunnelDestinationIPModeMulticast
	// HasAddress checks if Address has been set in VxlanV6TunnelDestinationIPModeMulticast
	HasAddress() bool
}

// IPv6 Multicast address
// Address returns a string
func (obj *vxlanV6TunnelDestinationIPModeMulticast) Address() string {

	return *obj.obj.Address

}

// IPv6 Multicast address
// Address returns a string
func (obj *vxlanV6TunnelDestinationIPModeMulticast) HasAddress() bool {
	return obj.obj.Address != nil
}

// IPv6 Multicast address
// SetAddress sets the string value in the VxlanV6TunnelDestinationIPModeMulticast object
func (obj *vxlanV6TunnelDestinationIPModeMulticast) SetAddress(value string) VxlanV6TunnelDestinationIPModeMulticast {

	obj.obj.Address = &value
	return obj
}

func (obj *vxlanV6TunnelDestinationIPModeMulticast) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Address != nil {

		err := obj.validateIpv6(obj.Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on VxlanV6TunnelDestinationIPModeMulticast.Address"))
		}

	}

}

func (obj *vxlanV6TunnelDestinationIPModeMulticast) setDefault() {

}
