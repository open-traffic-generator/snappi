package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6InterfaceIapd *****
type dhcpv6InterfaceIapd struct {
	validation
	obj          *otg.Dhcpv6InterfaceIapd
	marshaller   marshalDhcpv6InterfaceIapd
	unMarshaller unMarshalDhcpv6InterfaceIapd
}

func NewDhcpv6InterfaceIapd() Dhcpv6InterfaceIapd {
	obj := dhcpv6InterfaceIapd{obj: &otg.Dhcpv6InterfaceIapd{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6InterfaceIapd) msg() *otg.Dhcpv6InterfaceIapd {
	return obj.obj
}

func (obj *dhcpv6InterfaceIapd) setMsg(msg *otg.Dhcpv6InterfaceIapd) Dhcpv6InterfaceIapd {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6InterfaceIapd struct {
	obj *dhcpv6InterfaceIapd
}

type marshalDhcpv6InterfaceIapd interface {
	// ToProto marshals Dhcpv6InterfaceIapd to protobuf object *otg.Dhcpv6InterfaceIapd
	ToProto() (*otg.Dhcpv6InterfaceIapd, error)
	// ToPbText marshals Dhcpv6InterfaceIapd to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6InterfaceIapd to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6InterfaceIapd to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Dhcpv6InterfaceIapd to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldhcpv6InterfaceIapd struct {
	obj *dhcpv6InterfaceIapd
}

type unMarshalDhcpv6InterfaceIapd interface {
	// FromProto unmarshals Dhcpv6InterfaceIapd from protobuf object *otg.Dhcpv6InterfaceIapd
	FromProto(msg *otg.Dhcpv6InterfaceIapd) (Dhcpv6InterfaceIapd, error)
	// FromPbText unmarshals Dhcpv6InterfaceIapd from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6InterfaceIapd from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6InterfaceIapd from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6InterfaceIapd) Marshal() marshalDhcpv6InterfaceIapd {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6InterfaceIapd{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6InterfaceIapd) Unmarshal() unMarshalDhcpv6InterfaceIapd {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6InterfaceIapd{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6InterfaceIapd) ToProto() (*otg.Dhcpv6InterfaceIapd, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6InterfaceIapd) FromProto(msg *otg.Dhcpv6InterfaceIapd) (Dhcpv6InterfaceIapd, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6InterfaceIapd) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6InterfaceIapd) FromPbText(value string) error {
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

func (m *marshaldhcpv6InterfaceIapd) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6InterfaceIapd) FromYaml(value string) error {
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

func (m *marshaldhcpv6InterfaceIapd) ToJsonRaw() (string, error) {
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

func (m *marshaldhcpv6InterfaceIapd) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6InterfaceIapd) FromJson(value string) error {
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

func (obj *dhcpv6InterfaceIapd) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6InterfaceIapd) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6InterfaceIapd) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6InterfaceIapd) Clone() (Dhcpv6InterfaceIapd, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6InterfaceIapd()
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

// Dhcpv6InterfaceIapd is the IPv6 IAPD address and prefix length associated with this DHCP Client session.
type Dhcpv6InterfaceIapd interface {
	Validation
	// msg marshals Dhcpv6InterfaceIapd to protobuf object *otg.Dhcpv6InterfaceIapd
	// and doesn't set defaults
	msg() *otg.Dhcpv6InterfaceIapd
	// setMsg unmarshals Dhcpv6InterfaceIapd from protobuf object *otg.Dhcpv6InterfaceIapd
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6InterfaceIapd) Dhcpv6InterfaceIapd
	// provides marshal interface
	Marshal() marshalDhcpv6InterfaceIapd
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6InterfaceIapd
	// validate validates Dhcpv6InterfaceIapd
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6InterfaceIapd, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Address returns string, set in Dhcpv6InterfaceIapd.
	Address() string
	// SetAddress assigns string provided by user to Dhcpv6InterfaceIapd
	SetAddress(value string) Dhcpv6InterfaceIapd
	// HasAddress checks if Address has been set in Dhcpv6InterfaceIapd
	HasAddress() bool
	// PrefixLength returns uint32, set in Dhcpv6InterfaceIapd.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to Dhcpv6InterfaceIapd
	SetPrefixLength(value uint32) Dhcpv6InterfaceIapd
	// HasPrefixLength checks if PrefixLength has been set in Dhcpv6InterfaceIapd
	HasPrefixLength() bool
	// LeaseTime returns uint32, set in Dhcpv6InterfaceIapd.
	LeaseTime() uint32
	// SetLeaseTime assigns uint32 provided by user to Dhcpv6InterfaceIapd
	SetLeaseTime(value uint32) Dhcpv6InterfaceIapd
	// HasLeaseTime checks if LeaseTime has been set in Dhcpv6InterfaceIapd
	HasLeaseTime() bool
}

// The IAPD address associated with this DHCPv6 Client session.
// Address returns a string
func (obj *dhcpv6InterfaceIapd) Address() string {

	return *obj.obj.Address

}

// The IAPD address associated with this DHCPv6 Client session.
// Address returns a string
func (obj *dhcpv6InterfaceIapd) HasAddress() bool {
	return obj.obj.Address != nil
}

// The IAPD address associated with this DHCPv6 Client session.
// SetAddress sets the string value in the Dhcpv6InterfaceIapd object
func (obj *dhcpv6InterfaceIapd) SetAddress(value string) Dhcpv6InterfaceIapd {

	obj.obj.Address = &value
	return obj
}

// The prefix length of the IAPD address associated with this DHCPv6 Client session.
// PrefixLength returns a uint32
func (obj *dhcpv6InterfaceIapd) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// The prefix length of the IAPD address associated with this DHCPv6 Client session.
// PrefixLength returns a uint32
func (obj *dhcpv6InterfaceIapd) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// The prefix length of the IAPD address associated with this DHCPv6 Client session.
// SetPrefixLength sets the uint32 value in the Dhcpv6InterfaceIapd object
func (obj *dhcpv6InterfaceIapd) SetPrefixLength(value uint32) Dhcpv6InterfaceIapd {

	obj.obj.PrefixLength = &value
	return obj
}

// The duration of the IPv6 address lease, in seconds.
// LeaseTime returns a uint32
func (obj *dhcpv6InterfaceIapd) LeaseTime() uint32 {

	return *obj.obj.LeaseTime

}

// The duration of the IPv6 address lease, in seconds.
// LeaseTime returns a uint32
func (obj *dhcpv6InterfaceIapd) HasLeaseTime() bool {
	return obj.obj.LeaseTime != nil
}

// The duration of the IPv6 address lease, in seconds.
// SetLeaseTime sets the uint32 value in the Dhcpv6InterfaceIapd object
func (obj *dhcpv6InterfaceIapd) SetLeaseTime(value uint32) Dhcpv6InterfaceIapd {

	obj.obj.LeaseTime = &value
	return obj
}

func (obj *dhcpv6InterfaceIapd) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Address != nil {

		err := obj.validateIpv6(obj.Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Dhcpv6InterfaceIapd.Address"))
		}

	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Dhcpv6InterfaceIapd.PrefixLength <= 128 but Got %d", *obj.obj.PrefixLength))
		}

	}

}

func (obj *dhcpv6InterfaceIapd) setDefault() {

}
