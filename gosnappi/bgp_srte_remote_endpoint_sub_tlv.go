package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpSrteRemoteEndpointSubTlv *****
type bgpSrteRemoteEndpointSubTlv struct {
	validation
	obj          *otg.BgpSrteRemoteEndpointSubTlv
	marshaller   marshalBgpSrteRemoteEndpointSubTlv
	unMarshaller unMarshalBgpSrteRemoteEndpointSubTlv
}

func NewBgpSrteRemoteEndpointSubTlv() BgpSrteRemoteEndpointSubTlv {
	obj := bgpSrteRemoteEndpointSubTlv{obj: &otg.BgpSrteRemoteEndpointSubTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpSrteRemoteEndpointSubTlv) msg() *otg.BgpSrteRemoteEndpointSubTlv {
	return obj.obj
}

func (obj *bgpSrteRemoteEndpointSubTlv) setMsg(msg *otg.BgpSrteRemoteEndpointSubTlv) BgpSrteRemoteEndpointSubTlv {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpSrteRemoteEndpointSubTlv struct {
	obj *bgpSrteRemoteEndpointSubTlv
}

type marshalBgpSrteRemoteEndpointSubTlv interface {
	// ToProto marshals BgpSrteRemoteEndpointSubTlv to protobuf object *otg.BgpSrteRemoteEndpointSubTlv
	ToProto() (*otg.BgpSrteRemoteEndpointSubTlv, error)
	// ToPbText marshals BgpSrteRemoteEndpointSubTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpSrteRemoteEndpointSubTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpSrteRemoteEndpointSubTlv to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpSrteRemoteEndpointSubTlv to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpSrteRemoteEndpointSubTlv struct {
	obj *bgpSrteRemoteEndpointSubTlv
}

type unMarshalBgpSrteRemoteEndpointSubTlv interface {
	// FromProto unmarshals BgpSrteRemoteEndpointSubTlv from protobuf object *otg.BgpSrteRemoteEndpointSubTlv
	FromProto(msg *otg.BgpSrteRemoteEndpointSubTlv) (BgpSrteRemoteEndpointSubTlv, error)
	// FromPbText unmarshals BgpSrteRemoteEndpointSubTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpSrteRemoteEndpointSubTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpSrteRemoteEndpointSubTlv from JSON text
	FromJson(value string) error
}

func (obj *bgpSrteRemoteEndpointSubTlv) Marshal() marshalBgpSrteRemoteEndpointSubTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpSrteRemoteEndpointSubTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpSrteRemoteEndpointSubTlv) Unmarshal() unMarshalBgpSrteRemoteEndpointSubTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpSrteRemoteEndpointSubTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpSrteRemoteEndpointSubTlv) ToProto() (*otg.BgpSrteRemoteEndpointSubTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpSrteRemoteEndpointSubTlv) FromProto(msg *otg.BgpSrteRemoteEndpointSubTlv) (BgpSrteRemoteEndpointSubTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpSrteRemoteEndpointSubTlv) ToPbText() (string, error) {
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

func (m *unMarshalbgpSrteRemoteEndpointSubTlv) FromPbText(value string) error {
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

func (m *marshalbgpSrteRemoteEndpointSubTlv) ToYaml() (string, error) {
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

func (m *unMarshalbgpSrteRemoteEndpointSubTlv) FromYaml(value string) error {
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

func (m *marshalbgpSrteRemoteEndpointSubTlv) ToJsonRaw() (string, error) {
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

func (m *marshalbgpSrteRemoteEndpointSubTlv) ToJson() (string, error) {
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

func (m *unMarshalbgpSrteRemoteEndpointSubTlv) FromJson(value string) error {
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

func (obj *bgpSrteRemoteEndpointSubTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpSrteRemoteEndpointSubTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpSrteRemoteEndpointSubTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpSrteRemoteEndpointSubTlv) Clone() (BgpSrteRemoteEndpointSubTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpSrteRemoteEndpointSubTlv()
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

// BgpSrteRemoteEndpointSubTlv is configuration for the BGP remote endpoint sub TLV.
type BgpSrteRemoteEndpointSubTlv interface {
	Validation
	// msg marshals BgpSrteRemoteEndpointSubTlv to protobuf object *otg.BgpSrteRemoteEndpointSubTlv
	// and doesn't set defaults
	msg() *otg.BgpSrteRemoteEndpointSubTlv
	// setMsg unmarshals BgpSrteRemoteEndpointSubTlv from protobuf object *otg.BgpSrteRemoteEndpointSubTlv
	// and doesn't set defaults
	setMsg(*otg.BgpSrteRemoteEndpointSubTlv) BgpSrteRemoteEndpointSubTlv
	// provides marshal interface
	Marshal() marshalBgpSrteRemoteEndpointSubTlv
	// provides unmarshal interface
	Unmarshal() unMarshalBgpSrteRemoteEndpointSubTlv
	// validate validates BgpSrteRemoteEndpointSubTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpSrteRemoteEndpointSubTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// AsNumber returns uint32, set in BgpSrteRemoteEndpointSubTlv.
	AsNumber() uint32
	// SetAsNumber assigns uint32 provided by user to BgpSrteRemoteEndpointSubTlv
	SetAsNumber(value uint32) BgpSrteRemoteEndpointSubTlv
	// HasAsNumber checks if AsNumber has been set in BgpSrteRemoteEndpointSubTlv
	HasAsNumber() bool
	// AddressFamily returns BgpSrteRemoteEndpointSubTlvAddressFamilyEnum, set in BgpSrteRemoteEndpointSubTlv
	AddressFamily() BgpSrteRemoteEndpointSubTlvAddressFamilyEnum
	// SetAddressFamily assigns BgpSrteRemoteEndpointSubTlvAddressFamilyEnum provided by user to BgpSrteRemoteEndpointSubTlv
	SetAddressFamily(value BgpSrteRemoteEndpointSubTlvAddressFamilyEnum) BgpSrteRemoteEndpointSubTlv
	// HasAddressFamily checks if AddressFamily has been set in BgpSrteRemoteEndpointSubTlv
	HasAddressFamily() bool
	// Ipv4Address returns string, set in BgpSrteRemoteEndpointSubTlv.
	Ipv4Address() string
	// SetIpv4Address assigns string provided by user to BgpSrteRemoteEndpointSubTlv
	SetIpv4Address(value string) BgpSrteRemoteEndpointSubTlv
	// HasIpv4Address checks if Ipv4Address has been set in BgpSrteRemoteEndpointSubTlv
	HasIpv4Address() bool
	// Ipv6Address returns string, set in BgpSrteRemoteEndpointSubTlv.
	Ipv6Address() string
	// SetIpv6Address assigns string provided by user to BgpSrteRemoteEndpointSubTlv
	SetIpv6Address(value string) BgpSrteRemoteEndpointSubTlv
	// HasIpv6Address checks if Ipv6Address has been set in BgpSrteRemoteEndpointSubTlv
	HasIpv6Address() bool
}

// Autonomous system (AS) number
// AsNumber returns a uint32
func (obj *bgpSrteRemoteEndpointSubTlv) AsNumber() uint32 {

	return *obj.obj.AsNumber

}

// Autonomous system (AS) number
// AsNumber returns a uint32
func (obj *bgpSrteRemoteEndpointSubTlv) HasAsNumber() bool {
	return obj.obj.AsNumber != nil
}

// Autonomous system (AS) number
// SetAsNumber sets the uint32 value in the BgpSrteRemoteEndpointSubTlv object
func (obj *bgpSrteRemoteEndpointSubTlv) SetAsNumber(value uint32) BgpSrteRemoteEndpointSubTlv {

	obj.obj.AsNumber = &value
	return obj
}

type BgpSrteRemoteEndpointSubTlvAddressFamilyEnum string

// Enum of AddressFamily on BgpSrteRemoteEndpointSubTlv
var BgpSrteRemoteEndpointSubTlvAddressFamily = struct {
	IPV4 BgpSrteRemoteEndpointSubTlvAddressFamilyEnum
	IPV6 BgpSrteRemoteEndpointSubTlvAddressFamilyEnum
}{
	IPV4: BgpSrteRemoteEndpointSubTlvAddressFamilyEnum("ipv4"),
	IPV6: BgpSrteRemoteEndpointSubTlvAddressFamilyEnum("ipv6"),
}

func (obj *bgpSrteRemoteEndpointSubTlv) AddressFamily() BgpSrteRemoteEndpointSubTlvAddressFamilyEnum {
	return BgpSrteRemoteEndpointSubTlvAddressFamilyEnum(obj.obj.AddressFamily.Enum().String())
}

// Determines the address type
// AddressFamily returns a string
func (obj *bgpSrteRemoteEndpointSubTlv) HasAddressFamily() bool {
	return obj.obj.AddressFamily != nil
}

func (obj *bgpSrteRemoteEndpointSubTlv) SetAddressFamily(value BgpSrteRemoteEndpointSubTlvAddressFamilyEnum) BgpSrteRemoteEndpointSubTlv {
	intValue, ok := otg.BgpSrteRemoteEndpointSubTlv_AddressFamily_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpSrteRemoteEndpointSubTlvAddressFamilyEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpSrteRemoteEndpointSubTlv_AddressFamily_Enum(intValue)
	obj.obj.AddressFamily = &enumValue

	return obj
}

// The IPv4 address of the Remote Endpoint.
// Ipv4Address returns a string
func (obj *bgpSrteRemoteEndpointSubTlv) Ipv4Address() string {

	return *obj.obj.Ipv4Address

}

// The IPv4 address of the Remote Endpoint.
// Ipv4Address returns a string
func (obj *bgpSrteRemoteEndpointSubTlv) HasIpv4Address() bool {
	return obj.obj.Ipv4Address != nil
}

// The IPv4 address of the Remote Endpoint.
// SetIpv4Address sets the string value in the BgpSrteRemoteEndpointSubTlv object
func (obj *bgpSrteRemoteEndpointSubTlv) SetIpv4Address(value string) BgpSrteRemoteEndpointSubTlv {

	obj.obj.Ipv4Address = &value
	return obj
}

// The IPv6 address of the Remote Endpoint.
// Ipv6Address returns a string
func (obj *bgpSrteRemoteEndpointSubTlv) Ipv6Address() string {

	return *obj.obj.Ipv6Address

}

// The IPv6 address of the Remote Endpoint.
// Ipv6Address returns a string
func (obj *bgpSrteRemoteEndpointSubTlv) HasIpv6Address() bool {
	return obj.obj.Ipv6Address != nil
}

// The IPv6 address of the Remote Endpoint.
// SetIpv6Address sets the string value in the BgpSrteRemoteEndpointSubTlv object
func (obj *bgpSrteRemoteEndpointSubTlv) SetIpv6Address(value string) BgpSrteRemoteEndpointSubTlv {

	obj.obj.Ipv6Address = &value
	return obj
}

func (obj *bgpSrteRemoteEndpointSubTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv4Address != nil {

		err := obj.validateIpv4(obj.Ipv4Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteRemoteEndpointSubTlv.Ipv4Address"))
		}

	}

	if obj.obj.Ipv6Address != nil {

		err := obj.validateIpv6(obj.Ipv6Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteRemoteEndpointSubTlv.Ipv6Address"))
		}

	}

}

func (obj *bgpSrteRemoteEndpointSubTlv) setDefault() {
	if obj.obj.AsNumber == nil {
		obj.SetAsNumber(0)
	}
	if obj.obj.AddressFamily == nil {
		obj.SetAddressFamily(BgpSrteRemoteEndpointSubTlvAddressFamily.IPV4)

	}
	if obj.obj.Ipv4Address == nil {
		obj.SetIpv4Address("0.0.0.0")
	}
	if obj.obj.Ipv6Address == nil {
		obj.SetIpv6Address("::0")
	}

}
