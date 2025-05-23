package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** VxlanV4TunnelDestinationIPModeMulticast *****
type vxlanV4TunnelDestinationIPModeMulticast struct {
	validation
	obj          *otg.VxlanV4TunnelDestinationIPModeMulticast
	marshaller   marshalVxlanV4TunnelDestinationIPModeMulticast
	unMarshaller unMarshalVxlanV4TunnelDestinationIPModeMulticast
}

func NewVxlanV4TunnelDestinationIPModeMulticast() VxlanV4TunnelDestinationIPModeMulticast {
	obj := vxlanV4TunnelDestinationIPModeMulticast{obj: &otg.VxlanV4TunnelDestinationIPModeMulticast{}}
	obj.setDefault()
	return &obj
}

func (obj *vxlanV4TunnelDestinationIPModeMulticast) msg() *otg.VxlanV4TunnelDestinationIPModeMulticast {
	return obj.obj
}

func (obj *vxlanV4TunnelDestinationIPModeMulticast) setMsg(msg *otg.VxlanV4TunnelDestinationIPModeMulticast) VxlanV4TunnelDestinationIPModeMulticast {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalvxlanV4TunnelDestinationIPModeMulticast struct {
	obj *vxlanV4TunnelDestinationIPModeMulticast
}

type marshalVxlanV4TunnelDestinationIPModeMulticast interface {
	// ToProto marshals VxlanV4TunnelDestinationIPModeMulticast to protobuf object *otg.VxlanV4TunnelDestinationIPModeMulticast
	ToProto() (*otg.VxlanV4TunnelDestinationIPModeMulticast, error)
	// ToPbText marshals VxlanV4TunnelDestinationIPModeMulticast to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals VxlanV4TunnelDestinationIPModeMulticast to YAML text
	ToYaml() (string, error)
	// ToJson marshals VxlanV4TunnelDestinationIPModeMulticast to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals VxlanV4TunnelDestinationIPModeMulticast to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalvxlanV4TunnelDestinationIPModeMulticast struct {
	obj *vxlanV4TunnelDestinationIPModeMulticast
}

type unMarshalVxlanV4TunnelDestinationIPModeMulticast interface {
	// FromProto unmarshals VxlanV4TunnelDestinationIPModeMulticast from protobuf object *otg.VxlanV4TunnelDestinationIPModeMulticast
	FromProto(msg *otg.VxlanV4TunnelDestinationIPModeMulticast) (VxlanV4TunnelDestinationIPModeMulticast, error)
	// FromPbText unmarshals VxlanV4TunnelDestinationIPModeMulticast from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals VxlanV4TunnelDestinationIPModeMulticast from YAML text
	FromYaml(value string) error
	// FromJson unmarshals VxlanV4TunnelDestinationIPModeMulticast from JSON text
	FromJson(value string) error
}

func (obj *vxlanV4TunnelDestinationIPModeMulticast) Marshal() marshalVxlanV4TunnelDestinationIPModeMulticast {
	if obj.marshaller == nil {
		obj.marshaller = &marshalvxlanV4TunnelDestinationIPModeMulticast{obj: obj}
	}
	return obj.marshaller
}

func (obj *vxlanV4TunnelDestinationIPModeMulticast) Unmarshal() unMarshalVxlanV4TunnelDestinationIPModeMulticast {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalvxlanV4TunnelDestinationIPModeMulticast{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalvxlanV4TunnelDestinationIPModeMulticast) ToProto() (*otg.VxlanV4TunnelDestinationIPModeMulticast, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalvxlanV4TunnelDestinationIPModeMulticast) FromProto(msg *otg.VxlanV4TunnelDestinationIPModeMulticast) (VxlanV4TunnelDestinationIPModeMulticast, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalvxlanV4TunnelDestinationIPModeMulticast) ToPbText() (string, error) {
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

func (m *unMarshalvxlanV4TunnelDestinationIPModeMulticast) FromPbText(value string) error {
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

func (m *marshalvxlanV4TunnelDestinationIPModeMulticast) ToYaml() (string, error) {
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

func (m *unMarshalvxlanV4TunnelDestinationIPModeMulticast) FromYaml(value string) error {
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

func (m *marshalvxlanV4TunnelDestinationIPModeMulticast) ToJsonRaw() (string, error) {
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

func (m *marshalvxlanV4TunnelDestinationIPModeMulticast) ToJson() (string, error) {
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

func (m *unMarshalvxlanV4TunnelDestinationIPModeMulticast) FromJson(value string) error {
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

func (obj *vxlanV4TunnelDestinationIPModeMulticast) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *vxlanV4TunnelDestinationIPModeMulticast) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *vxlanV4TunnelDestinationIPModeMulticast) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *vxlanV4TunnelDestinationIPModeMulticast) Clone() (VxlanV4TunnelDestinationIPModeMulticast, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewVxlanV4TunnelDestinationIPModeMulticast()
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

// VxlanV4TunnelDestinationIPModeMulticast is multicast Group address for member VNI(VXLAN Network Identifier)
type VxlanV4TunnelDestinationIPModeMulticast interface {
	Validation
	// msg marshals VxlanV4TunnelDestinationIPModeMulticast to protobuf object *otg.VxlanV4TunnelDestinationIPModeMulticast
	// and doesn't set defaults
	msg() *otg.VxlanV4TunnelDestinationIPModeMulticast
	// setMsg unmarshals VxlanV4TunnelDestinationIPModeMulticast from protobuf object *otg.VxlanV4TunnelDestinationIPModeMulticast
	// and doesn't set defaults
	setMsg(*otg.VxlanV4TunnelDestinationIPModeMulticast) VxlanV4TunnelDestinationIPModeMulticast
	// provides marshal interface
	Marshal() marshalVxlanV4TunnelDestinationIPModeMulticast
	// provides unmarshal interface
	Unmarshal() unMarshalVxlanV4TunnelDestinationIPModeMulticast
	// validate validates VxlanV4TunnelDestinationIPModeMulticast
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (VxlanV4TunnelDestinationIPModeMulticast, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Address returns string, set in VxlanV4TunnelDestinationIPModeMulticast.
	Address() string
	// SetAddress assigns string provided by user to VxlanV4TunnelDestinationIPModeMulticast
	SetAddress(value string) VxlanV4TunnelDestinationIPModeMulticast
	// HasAddress checks if Address has been set in VxlanV4TunnelDestinationIPModeMulticast
	HasAddress() bool
}

// IPv4 Multicast address
// Address returns a string
func (obj *vxlanV4TunnelDestinationIPModeMulticast) Address() string {

	return *obj.obj.Address

}

// IPv4 Multicast address
// Address returns a string
func (obj *vxlanV4TunnelDestinationIPModeMulticast) HasAddress() bool {
	return obj.obj.Address != nil
}

// IPv4 Multicast address
// SetAddress sets the string value in the VxlanV4TunnelDestinationIPModeMulticast object
func (obj *vxlanV4TunnelDestinationIPModeMulticast) SetAddress(value string) VxlanV4TunnelDestinationIPModeMulticast {

	obj.obj.Address = &value
	return obj
}

func (obj *vxlanV4TunnelDestinationIPModeMulticast) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Address != nil {

		err := obj.validateIpv4(obj.Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on VxlanV4TunnelDestinationIPModeMulticast.Address"))
		}

	}

}

func (obj *vxlanV4TunnelDestinationIPModeMulticast) setDefault() {

}
