package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpCapabilitiesMplsSingleLabel *****
type bgpCapabilitiesMplsSingleLabel struct {
	validation
	obj          *otg.BgpCapabilitiesMplsSingleLabel
	marshaller   marshalBgpCapabilitiesMplsSingleLabel
	unMarshaller unMarshalBgpCapabilitiesMplsSingleLabel
}

func NewBgpCapabilitiesMplsSingleLabel() BgpCapabilitiesMplsSingleLabel {
	obj := bgpCapabilitiesMplsSingleLabel{obj: &otg.BgpCapabilitiesMplsSingleLabel{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpCapabilitiesMplsSingleLabel) msg() *otg.BgpCapabilitiesMplsSingleLabel {
	return obj.obj
}

func (obj *bgpCapabilitiesMplsSingleLabel) setMsg(msg *otg.BgpCapabilitiesMplsSingleLabel) BgpCapabilitiesMplsSingleLabel {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpCapabilitiesMplsSingleLabel struct {
	obj *bgpCapabilitiesMplsSingleLabel
}

type marshalBgpCapabilitiesMplsSingleLabel interface {
	// ToProto marshals BgpCapabilitiesMplsSingleLabel to protobuf object *otg.BgpCapabilitiesMplsSingleLabel
	ToProto() (*otg.BgpCapabilitiesMplsSingleLabel, error)
	// ToPbText marshals BgpCapabilitiesMplsSingleLabel to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpCapabilitiesMplsSingleLabel to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpCapabilitiesMplsSingleLabel to JSON text
	ToJson() (string, error)
}

type unMarshalbgpCapabilitiesMplsSingleLabel struct {
	obj *bgpCapabilitiesMplsSingleLabel
}

type unMarshalBgpCapabilitiesMplsSingleLabel interface {
	// FromProto unmarshals BgpCapabilitiesMplsSingleLabel from protobuf object *otg.BgpCapabilitiesMplsSingleLabel
	FromProto(msg *otg.BgpCapabilitiesMplsSingleLabel) (BgpCapabilitiesMplsSingleLabel, error)
	// FromPbText unmarshals BgpCapabilitiesMplsSingleLabel from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpCapabilitiesMplsSingleLabel from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpCapabilitiesMplsSingleLabel from JSON text
	FromJson(value string) error
}

func (obj *bgpCapabilitiesMplsSingleLabel) Marshal() marshalBgpCapabilitiesMplsSingleLabel {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpCapabilitiesMplsSingleLabel{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpCapabilitiesMplsSingleLabel) Unmarshal() unMarshalBgpCapabilitiesMplsSingleLabel {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpCapabilitiesMplsSingleLabel{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpCapabilitiesMplsSingleLabel) ToProto() (*otg.BgpCapabilitiesMplsSingleLabel, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpCapabilitiesMplsSingleLabel) FromProto(msg *otg.BgpCapabilitiesMplsSingleLabel) (BgpCapabilitiesMplsSingleLabel, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpCapabilitiesMplsSingleLabel) ToPbText() (string, error) {
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

func (m *unMarshalbgpCapabilitiesMplsSingleLabel) FromPbText(value string) error {
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

func (m *marshalbgpCapabilitiesMplsSingleLabel) ToYaml() (string, error) {
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

func (m *unMarshalbgpCapabilitiesMplsSingleLabel) FromYaml(value string) error {
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

func (m *marshalbgpCapabilitiesMplsSingleLabel) ToJson() (string, error) {
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

func (m *unMarshalbgpCapabilitiesMplsSingleLabel) FromJson(value string) error {
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

func (obj *bgpCapabilitiesMplsSingleLabel) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpCapabilitiesMplsSingleLabel) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpCapabilitiesMplsSingleLabel) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpCapabilitiesMplsSingleLabel) Clone() (BgpCapabilitiesMplsSingleLabel, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpCapabilitiesMplsSingleLabel()
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

// BgpCapabilitiesMplsSingleLabel is container for configuring capabilities for carrying single Label Information. The Capabilities Optional Parameter is a triple that includes a one-octet Capability Code,  a one-octet Capability length, and a variable-length Capability Value.
type BgpCapabilitiesMplsSingleLabel interface {
	Validation
	// msg marshals BgpCapabilitiesMplsSingleLabel to protobuf object *otg.BgpCapabilitiesMplsSingleLabel
	// and doesn't set defaults
	msg() *otg.BgpCapabilitiesMplsSingleLabel
	// setMsg unmarshals BgpCapabilitiesMplsSingleLabel from protobuf object *otg.BgpCapabilitiesMplsSingleLabel
	// and doesn't set defaults
	setMsg(*otg.BgpCapabilitiesMplsSingleLabel) BgpCapabilitiesMplsSingleLabel
	// provides marshal interface
	Marshal() marshalBgpCapabilitiesMplsSingleLabel
	// provides unmarshal interface
	Unmarshal() unMarshalBgpCapabilitiesMplsSingleLabel
	// validate validates BgpCapabilitiesMplsSingleLabel
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpCapabilitiesMplsSingleLabel, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4 returns bool, set in BgpCapabilitiesMplsSingleLabel.
	Ipv4() bool
	// SetIpv4 assigns bool provided by user to BgpCapabilitiesMplsSingleLabel
	SetIpv4(value bool) BgpCapabilitiesMplsSingleLabel
	// HasIpv4 checks if Ipv4 has been set in BgpCapabilitiesMplsSingleLabel
	HasIpv4() bool
	// Ipv6 returns bool, set in BgpCapabilitiesMplsSingleLabel.
	Ipv6() bool
	// SetIpv6 assigns bool provided by user to BgpCapabilitiesMplsSingleLabel
	SetIpv6(value bool) BgpCapabilitiesMplsSingleLabel
	// HasIpv6 checks if Ipv6 has been set in BgpCapabilitiesMplsSingleLabel
	HasIpv6() bool
	// VpnIpv4 returns bool, set in BgpCapabilitiesMplsSingleLabel.
	VpnIpv4() bool
	// SetVpnIpv4 assigns bool provided by user to BgpCapabilitiesMplsSingleLabel
	SetVpnIpv4(value bool) BgpCapabilitiesMplsSingleLabel
	// HasVpnIpv4 checks if VpnIpv4 has been set in BgpCapabilitiesMplsSingleLabel
	HasVpnIpv4() bool
	// VpnIpv6 returns bool, set in BgpCapabilitiesMplsSingleLabel.
	VpnIpv6() bool
	// SetVpnIpv6 assigns bool provided by user to BgpCapabilitiesMplsSingleLabel
	SetVpnIpv6(value bool) BgpCapabilitiesMplsSingleLabel
	// HasVpnIpv6 checks if VpnIpv6 has been set in BgpCapabilitiesMplsSingleLabel
	HasVpnIpv6() bool
}

// For the IPv4 address prefix, BGP speaker will advertise (AFI, SAFI) pair capability as (1, 4) in BGP Open Message under the Multiprotocol  extensions capability.
// Ipv4 returns a bool
func (obj *bgpCapabilitiesMplsSingleLabel) Ipv4() bool {

	return *obj.obj.Ipv4

}

// For the IPv4 address prefix, BGP speaker will advertise (AFI, SAFI) pair capability as (1, 4) in BGP Open Message under the Multiprotocol  extensions capability.
// Ipv4 returns a bool
func (obj *bgpCapabilitiesMplsSingleLabel) HasIpv4() bool {
	return obj.obj.Ipv4 != nil
}

// For the IPv4 address prefix, BGP speaker will advertise (AFI, SAFI) pair capability as (1, 4) in BGP Open Message under the Multiprotocol  extensions capability.
// SetIpv4 sets the bool value in the BgpCapabilitiesMplsSingleLabel object
func (obj *bgpCapabilitiesMplsSingleLabel) SetIpv4(value bool) BgpCapabilitiesMplsSingleLabel {

	obj.obj.Ipv4 = &value
	return obj
}

// For the IPv6 address prefix, BGP speaker will advertise (AFI, SAFI) pair capability as (2, 4) in BGP Open Message under the Multiprotocol  extensions capability.
// Ipv6 returns a bool
func (obj *bgpCapabilitiesMplsSingleLabel) Ipv6() bool {

	return *obj.obj.Ipv6

}

// For the IPv6 address prefix, BGP speaker will advertise (AFI, SAFI) pair capability as (2, 4) in BGP Open Message under the Multiprotocol  extensions capability.
// Ipv6 returns a bool
func (obj *bgpCapabilitiesMplsSingleLabel) HasIpv6() bool {
	return obj.obj.Ipv6 != nil
}

// For the IPv6 address prefix, BGP speaker will advertise (AFI, SAFI) pair capability as (2, 4) in BGP Open Message under the Multiprotocol  extensions capability.
// SetIpv6 sets the bool value in the BgpCapabilitiesMplsSingleLabel object
func (obj *bgpCapabilitiesMplsSingleLabel) SetIpv6(value bool) BgpCapabilitiesMplsSingleLabel {

	obj.obj.Ipv6 = &value
	return obj
}

// For the VPN-IPv4 address prefix, BGP speaker will advertise (AFI, SAFI) pair capability as (1, 128) in BGP Open Message under the Multiprotocol  extensions capability.
// VpnIpv4 returns a bool
func (obj *bgpCapabilitiesMplsSingleLabel) VpnIpv4() bool {

	return *obj.obj.VpnIpv4

}

// For the VPN-IPv4 address prefix, BGP speaker will advertise (AFI, SAFI) pair capability as (1, 128) in BGP Open Message under the Multiprotocol  extensions capability.
// VpnIpv4 returns a bool
func (obj *bgpCapabilitiesMplsSingleLabel) HasVpnIpv4() bool {
	return obj.obj.VpnIpv4 != nil
}

// For the VPN-IPv4 address prefix, BGP speaker will advertise (AFI, SAFI) pair capability as (1, 128) in BGP Open Message under the Multiprotocol  extensions capability.
// SetVpnIpv4 sets the bool value in the BgpCapabilitiesMplsSingleLabel object
func (obj *bgpCapabilitiesMplsSingleLabel) SetVpnIpv4(value bool) BgpCapabilitiesMplsSingleLabel {

	obj.obj.VpnIpv4 = &value
	return obj
}

// For the VPN-IPv6 address prefix, BGP speak will advertise (AFI, SAFI) pair capability as (2, 128) in BGP Open Messag under the Multiprotocol  extensions capability.
// VpnIpv6 returns a bool
func (obj *bgpCapabilitiesMplsSingleLabel) VpnIpv6() bool {

	return *obj.obj.VpnIpv6

}

// For the VPN-IPv6 address prefix, BGP speak will advertise (AFI, SAFI) pair capability as (2, 128) in BGP Open Messag under the Multiprotocol  extensions capability.
// VpnIpv6 returns a bool
func (obj *bgpCapabilitiesMplsSingleLabel) HasVpnIpv6() bool {
	return obj.obj.VpnIpv6 != nil
}

// For the VPN-IPv6 address prefix, BGP speak will advertise (AFI, SAFI) pair capability as (2, 128) in BGP Open Messag under the Multiprotocol  extensions capability.
// SetVpnIpv6 sets the bool value in the BgpCapabilitiesMplsSingleLabel object
func (obj *bgpCapabilitiesMplsSingleLabel) SetVpnIpv6(value bool) BgpCapabilitiesMplsSingleLabel {

	obj.obj.VpnIpv6 = &value
	return obj
}

func (obj *bgpCapabilitiesMplsSingleLabel) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *bgpCapabilitiesMplsSingleLabel) setDefault() {
	if obj.obj.Ipv4 == nil {
		obj.SetIpv4(true)
	}
	if obj.obj.Ipv6 == nil {
		obj.SetIpv6(true)
	}
	if obj.obj.VpnIpv4 == nil {
		obj.SetVpnIpv4(true)
	}
	if obj.obj.VpnIpv6 == nil {
		obj.SetVpnIpv6(true)
	}

}
