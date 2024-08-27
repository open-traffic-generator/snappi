package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6InterfaceAddressInfo *****
type dhcpv6InterfaceAddressInfo struct {
	validation
	obj          *otg.Dhcpv6InterfaceAddressInfo
	marshaller   marshalDhcpv6InterfaceAddressInfo
	unMarshaller unMarshalDhcpv6InterfaceAddressInfo
}

func NewDhcpv6InterfaceAddressInfo() Dhcpv6InterfaceAddressInfo {
	obj := dhcpv6InterfaceAddressInfo{obj: &otg.Dhcpv6InterfaceAddressInfo{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6InterfaceAddressInfo) msg() *otg.Dhcpv6InterfaceAddressInfo {
	return obj.obj
}

func (obj *dhcpv6InterfaceAddressInfo) setMsg(msg *otg.Dhcpv6InterfaceAddressInfo) Dhcpv6InterfaceAddressInfo {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6InterfaceAddressInfo struct {
	obj *dhcpv6InterfaceAddressInfo
}

type marshalDhcpv6InterfaceAddressInfo interface {
	// ToProto marshals Dhcpv6InterfaceAddressInfo to protobuf object *otg.Dhcpv6InterfaceAddressInfo
	ToProto() (*otg.Dhcpv6InterfaceAddressInfo, error)
	// ToPbText marshals Dhcpv6InterfaceAddressInfo to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6InterfaceAddressInfo to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6InterfaceAddressInfo to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpv6InterfaceAddressInfo struct {
	obj *dhcpv6InterfaceAddressInfo
}

type unMarshalDhcpv6InterfaceAddressInfo interface {
	// FromProto unmarshals Dhcpv6InterfaceAddressInfo from protobuf object *otg.Dhcpv6InterfaceAddressInfo
	FromProto(msg *otg.Dhcpv6InterfaceAddressInfo) (Dhcpv6InterfaceAddressInfo, error)
	// FromPbText unmarshals Dhcpv6InterfaceAddressInfo from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6InterfaceAddressInfo from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6InterfaceAddressInfo from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6InterfaceAddressInfo) Marshal() marshalDhcpv6InterfaceAddressInfo {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6InterfaceAddressInfo{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6InterfaceAddressInfo) Unmarshal() unMarshalDhcpv6InterfaceAddressInfo {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6InterfaceAddressInfo{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6InterfaceAddressInfo) ToProto() (*otg.Dhcpv6InterfaceAddressInfo, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6InterfaceAddressInfo) FromProto(msg *otg.Dhcpv6InterfaceAddressInfo) (Dhcpv6InterfaceAddressInfo, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6InterfaceAddressInfo) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6InterfaceAddressInfo) FromPbText(value string) error {
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

func (m *marshaldhcpv6InterfaceAddressInfo) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6InterfaceAddressInfo) FromYaml(value string) error {
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

func (m *marshaldhcpv6InterfaceAddressInfo) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6InterfaceAddressInfo) FromJson(value string) error {
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

func (obj *dhcpv6InterfaceAddressInfo) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6InterfaceAddressInfo) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6InterfaceAddressInfo) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6InterfaceAddressInfo) Clone() (Dhcpv6InterfaceAddressInfo, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6InterfaceAddressInfo()
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

// Dhcpv6InterfaceAddressInfo is the IPv6 address and gateway associated with this DHCP Client session.
type Dhcpv6InterfaceAddressInfo interface {
	Validation
	// msg marshals Dhcpv6InterfaceAddressInfo to protobuf object *otg.Dhcpv6InterfaceAddressInfo
	// and doesn't set defaults
	msg() *otg.Dhcpv6InterfaceAddressInfo
	// setMsg unmarshals Dhcpv6InterfaceAddressInfo from protobuf object *otg.Dhcpv6InterfaceAddressInfo
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6InterfaceAddressInfo) Dhcpv6InterfaceAddressInfo
	// provides marshal interface
	Marshal() marshalDhcpv6InterfaceAddressInfo
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6InterfaceAddressInfo
	// validate validates Dhcpv6InterfaceAddressInfo
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6InterfaceAddressInfo, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Address returns string, set in Dhcpv6InterfaceAddressInfo.
	Address() string
	// SetAddress assigns string provided by user to Dhcpv6InterfaceAddressInfo
	SetAddress(value string) Dhcpv6InterfaceAddressInfo
	// HasAddress checks if Address has been set in Dhcpv6InterfaceAddressInfo
	HasAddress() bool
	// Gateway returns uint32, set in Dhcpv6InterfaceAddressInfo.
	Gateway() uint32
	// SetGateway assigns uint32 provided by user to Dhcpv6InterfaceAddressInfo
	SetGateway(value uint32) Dhcpv6InterfaceAddressInfo
	// HasGateway checks if Gateway has been set in Dhcpv6InterfaceAddressInfo
	HasGateway() bool
}

// The address associated with this DHCPv6 Client session.
// Address returns a string
func (obj *dhcpv6InterfaceAddressInfo) Address() string {

	return *obj.obj.Address

}

// The address associated with this DHCPv6 Client session.
// Address returns a string
func (obj *dhcpv6InterfaceAddressInfo) HasAddress() bool {
	return obj.obj.Address != nil
}

// The address associated with this DHCPv6 Client session.
// SetAddress sets the string value in the Dhcpv6InterfaceAddressInfo object
func (obj *dhcpv6InterfaceAddressInfo) SetAddress(value string) Dhcpv6InterfaceAddressInfo {

	obj.obj.Address = &value
	return obj
}

// The Gateway address associated with this DHCPv6 Client session.
// Gateway returns a uint32
func (obj *dhcpv6InterfaceAddressInfo) Gateway() uint32 {

	return *obj.obj.Gateway

}

// The Gateway address associated with this DHCPv6 Client session.
// Gateway returns a uint32
func (obj *dhcpv6InterfaceAddressInfo) HasGateway() bool {
	return obj.obj.Gateway != nil
}

// The Gateway address associated with this DHCPv6 Client session.
// SetGateway sets the uint32 value in the Dhcpv6InterfaceAddressInfo object
func (obj *dhcpv6InterfaceAddressInfo) SetGateway(value uint32) Dhcpv6InterfaceAddressInfo {

	obj.obj.Gateway = &value
	return obj
}

func (obj *dhcpv6InterfaceAddressInfo) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Address != nil {

		err := obj.validateIpv6(obj.Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Dhcpv6InterfaceAddressInfo.Address"))
		}

	}

	if obj.obj.Gateway != nil {

		if *obj.obj.Gateway > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Dhcpv6InterfaceAddressInfo.Gateway <= 128 but Got %d", *obj.obj.Gateway))
		}

	}

}

func (obj *dhcpv6InterfaceAddressInfo) setDefault() {

}
