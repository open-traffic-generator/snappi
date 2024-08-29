package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6InterfaceIa *****
type dhcpv6InterfaceIa struct {
	validation
	obj          *otg.Dhcpv6InterfaceIa
	marshaller   marshalDhcpv6InterfaceIa
	unMarshaller unMarshalDhcpv6InterfaceIa
}

func NewDhcpv6InterfaceIa() Dhcpv6InterfaceIa {
	obj := dhcpv6InterfaceIa{obj: &otg.Dhcpv6InterfaceIa{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6InterfaceIa) msg() *otg.Dhcpv6InterfaceIa {
	return obj.obj
}

func (obj *dhcpv6InterfaceIa) setMsg(msg *otg.Dhcpv6InterfaceIa) Dhcpv6InterfaceIa {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6InterfaceIa struct {
	obj *dhcpv6InterfaceIa
}

type marshalDhcpv6InterfaceIa interface {
	// ToProto marshals Dhcpv6InterfaceIa to protobuf object *otg.Dhcpv6InterfaceIa
	ToProto() (*otg.Dhcpv6InterfaceIa, error)
	// ToPbText marshals Dhcpv6InterfaceIa to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6InterfaceIa to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6InterfaceIa to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpv6InterfaceIa struct {
	obj *dhcpv6InterfaceIa
}

type unMarshalDhcpv6InterfaceIa interface {
	// FromProto unmarshals Dhcpv6InterfaceIa from protobuf object *otg.Dhcpv6InterfaceIa
	FromProto(msg *otg.Dhcpv6InterfaceIa) (Dhcpv6InterfaceIa, error)
	// FromPbText unmarshals Dhcpv6InterfaceIa from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6InterfaceIa from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6InterfaceIa from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6InterfaceIa) Marshal() marshalDhcpv6InterfaceIa {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6InterfaceIa{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6InterfaceIa) Unmarshal() unMarshalDhcpv6InterfaceIa {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6InterfaceIa{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6InterfaceIa) ToProto() (*otg.Dhcpv6InterfaceIa, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6InterfaceIa) FromProto(msg *otg.Dhcpv6InterfaceIa) (Dhcpv6InterfaceIa, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6InterfaceIa) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6InterfaceIa) FromPbText(value string) error {
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

func (m *marshaldhcpv6InterfaceIa) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6InterfaceIa) FromYaml(value string) error {
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

func (m *marshaldhcpv6InterfaceIa) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6InterfaceIa) FromJson(value string) error {
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

func (obj *dhcpv6InterfaceIa) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6InterfaceIa) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6InterfaceIa) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6InterfaceIa) Clone() (Dhcpv6InterfaceIa, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6InterfaceIa()
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

// Dhcpv6InterfaceIa is the IPv6 IATA/IANA address and gateway associated with this DHCP Client session.
type Dhcpv6InterfaceIa interface {
	Validation
	// msg marshals Dhcpv6InterfaceIa to protobuf object *otg.Dhcpv6InterfaceIa
	// and doesn't set defaults
	msg() *otg.Dhcpv6InterfaceIa
	// setMsg unmarshals Dhcpv6InterfaceIa from protobuf object *otg.Dhcpv6InterfaceIa
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6InterfaceIa) Dhcpv6InterfaceIa
	// provides marshal interface
	Marshal() marshalDhcpv6InterfaceIa
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6InterfaceIa
	// validate validates Dhcpv6InterfaceIa
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6InterfaceIa, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Address returns string, set in Dhcpv6InterfaceIa.
	Address() string
	// SetAddress assigns string provided by user to Dhcpv6InterfaceIa
	SetAddress(value string) Dhcpv6InterfaceIa
	// HasAddress checks if Address has been set in Dhcpv6InterfaceIa
	HasAddress() bool
	// Gateway returns string, set in Dhcpv6InterfaceIa.
	Gateway() string
	// SetGateway assigns string provided by user to Dhcpv6InterfaceIa
	SetGateway(value string) Dhcpv6InterfaceIa
	// HasGateway checks if Gateway has been set in Dhcpv6InterfaceIa
	HasGateway() bool
	// LeaseTime returns uint32, set in Dhcpv6InterfaceIa.
	LeaseTime() uint32
	// SetLeaseTime assigns uint32 provided by user to Dhcpv6InterfaceIa
	SetLeaseTime(value uint32) Dhcpv6InterfaceIa
	// HasLeaseTime checks if LeaseTime has been set in Dhcpv6InterfaceIa
	HasLeaseTime() bool
}

// The address associated with this DHCPv6 Client session.
// Address returns a string
func (obj *dhcpv6InterfaceIa) Address() string {

	return *obj.obj.Address

}

// The address associated with this DHCPv6 Client session.
// Address returns a string
func (obj *dhcpv6InterfaceIa) HasAddress() bool {
	return obj.obj.Address != nil
}

// The address associated with this DHCPv6 Client session.
// SetAddress sets the string value in the Dhcpv6InterfaceIa object
func (obj *dhcpv6InterfaceIa) SetAddress(value string) Dhcpv6InterfaceIa {

	obj.obj.Address = &value
	return obj
}

// The Gateway address associated with this DHCPv6 Client session.
// Gateway returns a string
func (obj *dhcpv6InterfaceIa) Gateway() string {

	return *obj.obj.Gateway

}

// The Gateway address associated with this DHCPv6 Client session.
// Gateway returns a string
func (obj *dhcpv6InterfaceIa) HasGateway() bool {
	return obj.obj.Gateway != nil
}

// The Gateway address associated with this DHCPv6 Client session.
// SetGateway sets the string value in the Dhcpv6InterfaceIa object
func (obj *dhcpv6InterfaceIa) SetGateway(value string) Dhcpv6InterfaceIa {

	obj.obj.Gateway = &value
	return obj
}

// The duration of the IPv6 address lease, in seconds.
// LeaseTime returns a uint32
func (obj *dhcpv6InterfaceIa) LeaseTime() uint32 {

	return *obj.obj.LeaseTime

}

// The duration of the IPv6 address lease, in seconds.
// LeaseTime returns a uint32
func (obj *dhcpv6InterfaceIa) HasLeaseTime() bool {
	return obj.obj.LeaseTime != nil
}

// The duration of the IPv6 address lease, in seconds.
// SetLeaseTime sets the uint32 value in the Dhcpv6InterfaceIa object
func (obj *dhcpv6InterfaceIa) SetLeaseTime(value uint32) Dhcpv6InterfaceIa {

	obj.obj.LeaseTime = &value
	return obj
}

func (obj *dhcpv6InterfaceIa) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Address != nil {

		err := obj.validateIpv6(obj.Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Dhcpv6InterfaceIa.Address"))
		}

	}

	if obj.obj.Gateway != nil {

		err := obj.validateIpv6(obj.Gateway())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Dhcpv6InterfaceIa.Gateway"))
		}

	}

}

func (obj *dhcpv6InterfaceIa) setDefault() {

}
