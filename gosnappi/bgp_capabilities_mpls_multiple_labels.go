package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpCapabilitiesMplsMultipleLabels *****
type bgpCapabilitiesMplsMultipleLabels struct {
	validation
	obj          *otg.BgpCapabilitiesMplsMultipleLabels
	marshaller   marshalBgpCapabilitiesMplsMultipleLabels
	unMarshaller unMarshalBgpCapabilitiesMplsMultipleLabels
}

func NewBgpCapabilitiesMplsMultipleLabels() BgpCapabilitiesMplsMultipleLabels {
	obj := bgpCapabilitiesMplsMultipleLabels{obj: &otg.BgpCapabilitiesMplsMultipleLabels{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpCapabilitiesMplsMultipleLabels) msg() *otg.BgpCapabilitiesMplsMultipleLabels {
	return obj.obj
}

func (obj *bgpCapabilitiesMplsMultipleLabels) setMsg(msg *otg.BgpCapabilitiesMplsMultipleLabels) BgpCapabilitiesMplsMultipleLabels {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpCapabilitiesMplsMultipleLabels struct {
	obj *bgpCapabilitiesMplsMultipleLabels
}

type marshalBgpCapabilitiesMplsMultipleLabels interface {
	// ToProto marshals BgpCapabilitiesMplsMultipleLabels to protobuf object *otg.BgpCapabilitiesMplsMultipleLabels
	ToProto() (*otg.BgpCapabilitiesMplsMultipleLabels, error)
	// ToPbText marshals BgpCapabilitiesMplsMultipleLabels to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpCapabilitiesMplsMultipleLabels to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpCapabilitiesMplsMultipleLabels to JSON text
	ToJson() (string, error)
}

type unMarshalbgpCapabilitiesMplsMultipleLabels struct {
	obj *bgpCapabilitiesMplsMultipleLabels
}

type unMarshalBgpCapabilitiesMplsMultipleLabels interface {
	// FromProto unmarshals BgpCapabilitiesMplsMultipleLabels from protobuf object *otg.BgpCapabilitiesMplsMultipleLabels
	FromProto(msg *otg.BgpCapabilitiesMplsMultipleLabels) (BgpCapabilitiesMplsMultipleLabels, error)
	// FromPbText unmarshals BgpCapabilitiesMplsMultipleLabels from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpCapabilitiesMplsMultipleLabels from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpCapabilitiesMplsMultipleLabels from JSON text
	FromJson(value string) error
}

func (obj *bgpCapabilitiesMplsMultipleLabels) Marshal() marshalBgpCapabilitiesMplsMultipleLabels {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpCapabilitiesMplsMultipleLabels{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpCapabilitiesMplsMultipleLabels) Unmarshal() unMarshalBgpCapabilitiesMplsMultipleLabels {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpCapabilitiesMplsMultipleLabels{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpCapabilitiesMplsMultipleLabels) ToProto() (*otg.BgpCapabilitiesMplsMultipleLabels, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpCapabilitiesMplsMultipleLabels) FromProto(msg *otg.BgpCapabilitiesMplsMultipleLabels) (BgpCapabilitiesMplsMultipleLabels, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpCapabilitiesMplsMultipleLabels) ToPbText() (string, error) {
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

func (m *unMarshalbgpCapabilitiesMplsMultipleLabels) FromPbText(value string) error {
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

func (m *marshalbgpCapabilitiesMplsMultipleLabels) ToYaml() (string, error) {
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

func (m *unMarshalbgpCapabilitiesMplsMultipleLabels) FromYaml(value string) error {
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

func (m *marshalbgpCapabilitiesMplsMultipleLabels) ToJson() (string, error) {
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

func (m *unMarshalbgpCapabilitiesMplsMultipleLabels) FromJson(value string) error {
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

func (obj *bgpCapabilitiesMplsMultipleLabels) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpCapabilitiesMplsMultipleLabels) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpCapabilitiesMplsMultipleLabels) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpCapabilitiesMplsMultipleLabels) Clone() (BgpCapabilitiesMplsMultipleLabels, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpCapabilitiesMplsMultipleLabels()
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

// BgpCapabilitiesMplsMultipleLabels is container for configuring capabilities for multiple carrying Labels Information as per RFC8277. The MPLS multiple Lable capability is advertised in Optional Capability under Multiple Label Capability (code 8). Reference: https://datatracker.ietf.org/doc/html/rfc8277#section-2.1.
type BgpCapabilitiesMplsMultipleLabels interface {
	Validation
	// msg marshals BgpCapabilitiesMplsMultipleLabels to protobuf object *otg.BgpCapabilitiesMplsMultipleLabels
	// and doesn't set defaults
	msg() *otg.BgpCapabilitiesMplsMultipleLabels
	// setMsg unmarshals BgpCapabilitiesMplsMultipleLabels from protobuf object *otg.BgpCapabilitiesMplsMultipleLabels
	// and doesn't set defaults
	setMsg(*otg.BgpCapabilitiesMplsMultipleLabels) BgpCapabilitiesMplsMultipleLabels
	// provides marshal interface
	Marshal() marshalBgpCapabilitiesMplsMultipleLabels
	// provides unmarshal interface
	Unmarshal() unMarshalBgpCapabilitiesMplsMultipleLabels
	// validate validates BgpCapabilitiesMplsMultipleLabels
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpCapabilitiesMplsMultipleLabels, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4LabelCount returns uint32, set in BgpCapabilitiesMplsMultipleLabels.
	Ipv4LabelCount() uint32
	// SetIpv4LabelCount assigns uint32 provided by user to BgpCapabilitiesMplsMultipleLabels
	SetIpv4LabelCount(value uint32) BgpCapabilitiesMplsMultipleLabels
	// HasIpv4LabelCount checks if Ipv4LabelCount has been set in BgpCapabilitiesMplsMultipleLabels
	HasIpv4LabelCount() bool
	// Ipv6LabelCount returns uint32, set in BgpCapabilitiesMplsMultipleLabels.
	Ipv6LabelCount() uint32
	// SetIpv6LabelCount assigns uint32 provided by user to BgpCapabilitiesMplsMultipleLabels
	SetIpv6LabelCount(value uint32) BgpCapabilitiesMplsMultipleLabels
	// HasIpv6LabelCount checks if Ipv6LabelCount has been set in BgpCapabilitiesMplsMultipleLabels
	HasIpv6LabelCount() bool
	// VpnIpv4LabelCount returns uint32, set in BgpCapabilitiesMplsMultipleLabels.
	VpnIpv4LabelCount() uint32
	// SetVpnIpv4LabelCount assigns uint32 provided by user to BgpCapabilitiesMplsMultipleLabels
	SetVpnIpv4LabelCount(value uint32) BgpCapabilitiesMplsMultipleLabels
	// HasVpnIpv4LabelCount checks if VpnIpv4LabelCount has been set in BgpCapabilitiesMplsMultipleLabels
	HasVpnIpv4LabelCount() bool
	// VpnIpv6LabelCount returns uint32, set in BgpCapabilitiesMplsMultipleLabels.
	VpnIpv6LabelCount() uint32
	// SetVpnIpv6LabelCount assigns uint32 provided by user to BgpCapabilitiesMplsMultipleLabels
	SetVpnIpv6LabelCount(value uint32) BgpCapabilitiesMplsMultipleLabels
	// HasVpnIpv6LabelCount checks if VpnIpv6LabelCount has been set in BgpCapabilitiesMplsMultipleLabels
	HasVpnIpv6LabelCount() bool
}

// For the IPv4 address prefix, BGP speak will advertise triple of the form <AFI=1, SAFI=4, Count=n> in BGP Open Message under the Multiprotocol  extensions capability. where n = ipv4_mpls_count.
// Ipv4LabelCount returns a uint32
func (obj *bgpCapabilitiesMplsMultipleLabels) Ipv4LabelCount() uint32 {

	return *obj.obj.Ipv4LabelCount

}

// For the IPv4 address prefix, BGP speak will advertise triple of the form <AFI=1, SAFI=4, Count=n> in BGP Open Message under the Multiprotocol  extensions capability. where n = ipv4_mpls_count.
// Ipv4LabelCount returns a uint32
func (obj *bgpCapabilitiesMplsMultipleLabels) HasIpv4LabelCount() bool {
	return obj.obj.Ipv4LabelCount != nil
}

// For the IPv4 address prefix, BGP speak will advertise triple of the form <AFI=1, SAFI=4, Count=n> in BGP Open Message under the Multiprotocol  extensions capability. where n = ipv4_mpls_count.
// SetIpv4LabelCount sets the uint32 value in the BgpCapabilitiesMplsMultipleLabels object
func (obj *bgpCapabilitiesMplsMultipleLabels) SetIpv4LabelCount(value uint32) BgpCapabilitiesMplsMultipleLabels {

	obj.obj.Ipv4LabelCount = &value
	return obj
}

// For the IPv6 address prefix, BGP speak will advertise triple of the form <AFI=2, SAFI=4, Count=n> in BGP Open Message under the Multiprotocol  extensions capability, where n = ipv6_mpls_count.
// Ipv6LabelCount returns a uint32
func (obj *bgpCapabilitiesMplsMultipleLabels) Ipv6LabelCount() uint32 {

	return *obj.obj.Ipv6LabelCount

}

// For the IPv6 address prefix, BGP speak will advertise triple of the form <AFI=2, SAFI=4, Count=n> in BGP Open Message under the Multiprotocol  extensions capability, where n = ipv6_mpls_count.
// Ipv6LabelCount returns a uint32
func (obj *bgpCapabilitiesMplsMultipleLabels) HasIpv6LabelCount() bool {
	return obj.obj.Ipv6LabelCount != nil
}

// For the IPv6 address prefix, BGP speak will advertise triple of the form <AFI=2, SAFI=4, Count=n> in BGP Open Message under the Multiprotocol  extensions capability, where n = ipv6_mpls_count.
// SetIpv6LabelCount sets the uint32 value in the BgpCapabilitiesMplsMultipleLabels object
func (obj *bgpCapabilitiesMplsMultipleLabels) SetIpv6LabelCount(value uint32) BgpCapabilitiesMplsMultipleLabels {

	obj.obj.Ipv6LabelCount = &value
	return obj
}

// For the VPN-IPv4 address prefix, BGP speak will advertise triple of the form <AFI=1, SAFI=128, Count=n> in BGP Open Message under the Multiprotocol  extensions capability, where n = vpn_ipv4_mpls_count.
// VpnIpv4LabelCount returns a uint32
func (obj *bgpCapabilitiesMplsMultipleLabels) VpnIpv4LabelCount() uint32 {

	return *obj.obj.VpnIpv4LabelCount

}

// For the VPN-IPv4 address prefix, BGP speak will advertise triple of the form <AFI=1, SAFI=128, Count=n> in BGP Open Message under the Multiprotocol  extensions capability, where n = vpn_ipv4_mpls_count.
// VpnIpv4LabelCount returns a uint32
func (obj *bgpCapabilitiesMplsMultipleLabels) HasVpnIpv4LabelCount() bool {
	return obj.obj.VpnIpv4LabelCount != nil
}

// For the VPN-IPv4 address prefix, BGP speak will advertise triple of the form <AFI=1, SAFI=128, Count=n> in BGP Open Message under the Multiprotocol  extensions capability, where n = vpn_ipv4_mpls_count.
// SetVpnIpv4LabelCount sets the uint32 value in the BgpCapabilitiesMplsMultipleLabels object
func (obj *bgpCapabilitiesMplsMultipleLabels) SetVpnIpv4LabelCount(value uint32) BgpCapabilitiesMplsMultipleLabels {

	obj.obj.VpnIpv4LabelCount = &value
	return obj
}

// For the VPN-IPv4 address prefix, BGP speak will advertise triple of the form <AFI=2, SAFI=128, Count=n> in BGP Open Message under the Multiprotocol  extensions capability, where n = vpn_ipv6_mpls_count.
// VpnIpv6LabelCount returns a uint32
func (obj *bgpCapabilitiesMplsMultipleLabels) VpnIpv6LabelCount() uint32 {

	return *obj.obj.VpnIpv6LabelCount

}

// For the VPN-IPv4 address prefix, BGP speak will advertise triple of the form <AFI=2, SAFI=128, Count=n> in BGP Open Message under the Multiprotocol  extensions capability, where n = vpn_ipv6_mpls_count.
// VpnIpv6LabelCount returns a uint32
func (obj *bgpCapabilitiesMplsMultipleLabels) HasVpnIpv6LabelCount() bool {
	return obj.obj.VpnIpv6LabelCount != nil
}

// For the VPN-IPv4 address prefix, BGP speak will advertise triple of the form <AFI=2, SAFI=128, Count=n> in BGP Open Message under the Multiprotocol  extensions capability, where n = vpn_ipv6_mpls_count.
// SetVpnIpv6LabelCount sets the uint32 value in the BgpCapabilitiesMplsMultipleLabels object
func (obj *bgpCapabilitiesMplsMultipleLabels) SetVpnIpv6LabelCount(value uint32) BgpCapabilitiesMplsMultipleLabels {

	obj.obj.VpnIpv6LabelCount = &value
	return obj
}

func (obj *bgpCapabilitiesMplsMultipleLabels) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv4LabelCount != nil {

		if *obj.obj.Ipv4LabelCount < 2 || *obj.obj.Ipv4LabelCount > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("2 <= BgpCapabilitiesMplsMultipleLabels.Ipv4LabelCount <= 255 but Got %d", *obj.obj.Ipv4LabelCount))
		}

	}

	if obj.obj.Ipv6LabelCount != nil {

		if *obj.obj.Ipv6LabelCount < 2 || *obj.obj.Ipv6LabelCount > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("2 <= BgpCapabilitiesMplsMultipleLabels.Ipv6LabelCount <= 255 but Got %d", *obj.obj.Ipv6LabelCount))
		}

	}

	if obj.obj.VpnIpv4LabelCount != nil {

		if *obj.obj.VpnIpv4LabelCount < 2 || *obj.obj.VpnIpv4LabelCount > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("2 <= BgpCapabilitiesMplsMultipleLabels.VpnIpv4LabelCount <= 255 but Got %d", *obj.obj.VpnIpv4LabelCount))
		}

	}

	if obj.obj.VpnIpv6LabelCount != nil {

		if *obj.obj.VpnIpv6LabelCount < 2 || *obj.obj.VpnIpv6LabelCount > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("2 <= BgpCapabilitiesMplsMultipleLabels.VpnIpv6LabelCount <= 255 but Got %d", *obj.obj.VpnIpv6LabelCount))
		}

	}

}

func (obj *bgpCapabilitiesMplsMultipleLabels) setDefault() {
	if obj.obj.Ipv4LabelCount == nil {
		obj.SetIpv4LabelCount(2)
	}
	if obj.obj.Ipv6LabelCount == nil {
		obj.SetIpv6LabelCount(2)
	}
	if obj.obj.VpnIpv4LabelCount == nil {
		obj.SetVpnIpv4LabelCount(2)
	}
	if obj.obj.VpnIpv6LabelCount == nil {
		obj.SetVpnIpv6LabelCount(2)
	}

}
