package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisSRCapabilityFlags *****
type isisSRCapabilityFlags struct {
	validation
	obj          *otg.IsisSRCapabilityFlags
	marshaller   marshalIsisSRCapabilityFlags
	unMarshaller unMarshalIsisSRCapabilityFlags
}

func NewIsisSRCapabilityFlags() IsisSRCapabilityFlags {
	obj := isisSRCapabilityFlags{obj: &otg.IsisSRCapabilityFlags{}}
	obj.setDefault()
	return &obj
}

func (obj *isisSRCapabilityFlags) msg() *otg.IsisSRCapabilityFlags {
	return obj.obj
}

func (obj *isisSRCapabilityFlags) setMsg(msg *otg.IsisSRCapabilityFlags) IsisSRCapabilityFlags {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisSRCapabilityFlags struct {
	obj *isisSRCapabilityFlags
}

type marshalIsisSRCapabilityFlags interface {
	// ToProto marshals IsisSRCapabilityFlags to protobuf object *otg.IsisSRCapabilityFlags
	ToProto() (*otg.IsisSRCapabilityFlags, error)
	// ToPbText marshals IsisSRCapabilityFlags to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisSRCapabilityFlags to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisSRCapabilityFlags to JSON text
	ToJson() (string, error)
}

type unMarshalisisSRCapabilityFlags struct {
	obj *isisSRCapabilityFlags
}

type unMarshalIsisSRCapabilityFlags interface {
	// FromProto unmarshals IsisSRCapabilityFlags from protobuf object *otg.IsisSRCapabilityFlags
	FromProto(msg *otg.IsisSRCapabilityFlags) (IsisSRCapabilityFlags, error)
	// FromPbText unmarshals IsisSRCapabilityFlags from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisSRCapabilityFlags from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisSRCapabilityFlags from JSON text
	FromJson(value string) error
}

func (obj *isisSRCapabilityFlags) Marshal() marshalIsisSRCapabilityFlags {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisSRCapabilityFlags{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisSRCapabilityFlags) Unmarshal() unMarshalIsisSRCapabilityFlags {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisSRCapabilityFlags{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisSRCapabilityFlags) ToProto() (*otg.IsisSRCapabilityFlags, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisSRCapabilityFlags) FromProto(msg *otg.IsisSRCapabilityFlags) (IsisSRCapabilityFlags, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisSRCapabilityFlags) ToPbText() (string, error) {
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

func (m *unMarshalisisSRCapabilityFlags) FromPbText(value string) error {
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

func (m *marshalisisSRCapabilityFlags) ToYaml() (string, error) {
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

func (m *unMarshalisisSRCapabilityFlags) FromYaml(value string) error {
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

func (m *marshalisisSRCapabilityFlags) ToJson() (string, error) {
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

func (m *unMarshalisisSRCapabilityFlags) FromJson(value string) error {
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

func (obj *isisSRCapabilityFlags) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisSRCapabilityFlags) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisSRCapabilityFlags) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisSRCapabilityFlags) Clone() (IsisSRCapabilityFlags, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisSRCapabilityFlags()
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

// IsisSRCapabilityFlags is container for the configuration of IS-IS SR-CAPABILITY flags.
// The Router Capability TLV specifies flags that control its advertisement.
// The SR-Capabilities sub-TLV MUST be propagated throughout the level and MUST NOT be advertised across level boundaries.
// Therefore, Router Capability TLV distribution flags are set accordingly, i.e.,
// the S-Flag in the Router Capability TLV [RFC7981] MUST be unset.
// Reference: https://datatracker.ietf.org/doc/html/rfc8667#section-3.1-5.1.1.6.1.
type IsisSRCapabilityFlags interface {
	Validation
	// msg marshals IsisSRCapabilityFlags to protobuf object *otg.IsisSRCapabilityFlags
	// and doesn't set defaults
	msg() *otg.IsisSRCapabilityFlags
	// setMsg unmarshals IsisSRCapabilityFlags from protobuf object *otg.IsisSRCapabilityFlags
	// and doesn't set defaults
	setMsg(*otg.IsisSRCapabilityFlags) IsisSRCapabilityFlags
	// provides marshal interface
	Marshal() marshalIsisSRCapabilityFlags
	// provides unmarshal interface
	Unmarshal() unMarshalIsisSRCapabilityFlags
	// validate validates IsisSRCapabilityFlags
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisSRCapabilityFlags, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4Mpls returns bool, set in IsisSRCapabilityFlags.
	Ipv4Mpls() bool
	// SetIpv4Mpls assigns bool provided by user to IsisSRCapabilityFlags
	SetIpv4Mpls(value bool) IsisSRCapabilityFlags
	// HasIpv4Mpls checks if Ipv4Mpls has been set in IsisSRCapabilityFlags
	HasIpv4Mpls() bool
	// Ipv6Mpls returns bool, set in IsisSRCapabilityFlags.
	Ipv6Mpls() bool
	// SetIpv6Mpls assigns bool provided by user to IsisSRCapabilityFlags
	SetIpv6Mpls(value bool) IsisSRCapabilityFlags
	// HasIpv6Mpls checks if Ipv6Mpls has been set in IsisSRCapabilityFlags
	HasIpv6Mpls() bool
}

// I-Flag for the MPLS IPv4 Flag. If set, then the router is capable of processing SR-MPLS-encapsulated IPv4 packets on all interfaces.
// Ipv4Mpls returns a bool
func (obj *isisSRCapabilityFlags) Ipv4Mpls() bool {

	return *obj.obj.Ipv4Mpls

}

// I-Flag for the MPLS IPv4 Flag. If set, then the router is capable of processing SR-MPLS-encapsulated IPv4 packets on all interfaces.
// Ipv4Mpls returns a bool
func (obj *isisSRCapabilityFlags) HasIpv4Mpls() bool {
	return obj.obj.Ipv4Mpls != nil
}

// I-Flag for the MPLS IPv4 Flag. If set, then the router is capable of processing SR-MPLS-encapsulated IPv4 packets on all interfaces.
// SetIpv4Mpls sets the bool value in the IsisSRCapabilityFlags object
func (obj *isisSRCapabilityFlags) SetIpv4Mpls(value bool) IsisSRCapabilityFlags {

	obj.obj.Ipv4Mpls = &value
	return obj
}

// V-Flag for the MPLS IPv6 Flag. If set, then the router is capable of processing SR-MPLS-encapsulated IPv6 packets on all interfaces.
// Ipv6Mpls returns a bool
func (obj *isisSRCapabilityFlags) Ipv6Mpls() bool {

	return *obj.obj.Ipv6Mpls

}

// V-Flag for the MPLS IPv6 Flag. If set, then the router is capable of processing SR-MPLS-encapsulated IPv6 packets on all interfaces.
// Ipv6Mpls returns a bool
func (obj *isisSRCapabilityFlags) HasIpv6Mpls() bool {
	return obj.obj.Ipv6Mpls != nil
}

// V-Flag for the MPLS IPv6 Flag. If set, then the router is capable of processing SR-MPLS-encapsulated IPv6 packets on all interfaces.
// SetIpv6Mpls sets the bool value in the IsisSRCapabilityFlags object
func (obj *isisSRCapabilityFlags) SetIpv6Mpls(value bool) IsisSRCapabilityFlags {

	obj.obj.Ipv6Mpls = &value
	return obj
}

func (obj *isisSRCapabilityFlags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisSRCapabilityFlags) setDefault() {
	if obj.obj.Ipv4Mpls == nil {
		obj.SetIpv4Mpls(true)
	}
	if obj.obj.Ipv6Mpls == nil {
		obj.SetIpv6Mpls(true)
	}

}
