package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DhcpV6ServerSecondaryDns *****
type dhcpV6ServerSecondaryDns struct {
	validation
	obj          *otg.DhcpV6ServerSecondaryDns
	marshaller   marshalDhcpV6ServerSecondaryDns
	unMarshaller unMarshalDhcpV6ServerSecondaryDns
}

func NewDhcpV6ServerSecondaryDns() DhcpV6ServerSecondaryDns {
	obj := dhcpV6ServerSecondaryDns{obj: &otg.DhcpV6ServerSecondaryDns{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpV6ServerSecondaryDns) msg() *otg.DhcpV6ServerSecondaryDns {
	return obj.obj
}

func (obj *dhcpV6ServerSecondaryDns) setMsg(msg *otg.DhcpV6ServerSecondaryDns) DhcpV6ServerSecondaryDns {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpV6ServerSecondaryDns struct {
	obj *dhcpV6ServerSecondaryDns
}

type marshalDhcpV6ServerSecondaryDns interface {
	// ToProto marshals DhcpV6ServerSecondaryDns to protobuf object *otg.DhcpV6ServerSecondaryDns
	ToProto() (*otg.DhcpV6ServerSecondaryDns, error)
	// ToPbText marshals DhcpV6ServerSecondaryDns to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DhcpV6ServerSecondaryDns to YAML text
	ToYaml() (string, error)
	// ToJson marshals DhcpV6ServerSecondaryDns to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals DhcpV6ServerSecondaryDns to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldhcpV6ServerSecondaryDns struct {
	obj *dhcpV6ServerSecondaryDns
}

type unMarshalDhcpV6ServerSecondaryDns interface {
	// FromProto unmarshals DhcpV6ServerSecondaryDns from protobuf object *otg.DhcpV6ServerSecondaryDns
	FromProto(msg *otg.DhcpV6ServerSecondaryDns) (DhcpV6ServerSecondaryDns, error)
	// FromPbText unmarshals DhcpV6ServerSecondaryDns from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DhcpV6ServerSecondaryDns from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DhcpV6ServerSecondaryDns from JSON text
	FromJson(value string) error
}

func (obj *dhcpV6ServerSecondaryDns) Marshal() marshalDhcpV6ServerSecondaryDns {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpV6ServerSecondaryDns{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpV6ServerSecondaryDns) Unmarshal() unMarshalDhcpV6ServerSecondaryDns {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpV6ServerSecondaryDns{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpV6ServerSecondaryDns) ToProto() (*otg.DhcpV6ServerSecondaryDns, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpV6ServerSecondaryDns) FromProto(msg *otg.DhcpV6ServerSecondaryDns) (DhcpV6ServerSecondaryDns, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpV6ServerSecondaryDns) ToPbText() (string, error) {
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

func (m *unMarshaldhcpV6ServerSecondaryDns) FromPbText(value string) error {
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

func (m *marshaldhcpV6ServerSecondaryDns) ToYaml() (string, error) {
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

func (m *unMarshaldhcpV6ServerSecondaryDns) FromYaml(value string) error {
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

func (m *marshaldhcpV6ServerSecondaryDns) ToJsonRaw() (string, error) {
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

func (m *marshaldhcpV6ServerSecondaryDns) ToJson() (string, error) {
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

func (m *unMarshaldhcpV6ServerSecondaryDns) FromJson(value string) error {
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

func (obj *dhcpV6ServerSecondaryDns) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpV6ServerSecondaryDns) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpV6ServerSecondaryDns) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpV6ServerSecondaryDns) Clone() (DhcpV6ServerSecondaryDns, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpV6ServerSecondaryDns()
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

// DhcpV6ServerSecondaryDns is advanced Dns configuration for DHCPv6 server.
type DhcpV6ServerSecondaryDns interface {
	Validation
	// msg marshals DhcpV6ServerSecondaryDns to protobuf object *otg.DhcpV6ServerSecondaryDns
	// and doesn't set defaults
	msg() *otg.DhcpV6ServerSecondaryDns
	// setMsg unmarshals DhcpV6ServerSecondaryDns from protobuf object *otg.DhcpV6ServerSecondaryDns
	// and doesn't set defaults
	setMsg(*otg.DhcpV6ServerSecondaryDns) DhcpV6ServerSecondaryDns
	// provides marshal interface
	Marshal() marshalDhcpV6ServerSecondaryDns
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpV6ServerSecondaryDns
	// validate validates DhcpV6ServerSecondaryDns
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DhcpV6ServerSecondaryDns, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ip returns string, set in DhcpV6ServerSecondaryDns.
	Ip() string
	// SetIp assigns string provided by user to DhcpV6ServerSecondaryDns
	SetIp(value string) DhcpV6ServerSecondaryDns
	// HasIp checks if Ip has been set in DhcpV6ServerSecondaryDns
	HasIp() bool
}

// The secondary DNS server address that is offered to DHCP clients that request this information through a TLV option.
// Ip returns a string
func (obj *dhcpV6ServerSecondaryDns) Ip() string {

	return *obj.obj.Ip

}

// The secondary DNS server address that is offered to DHCP clients that request this information through a TLV option.
// Ip returns a string
func (obj *dhcpV6ServerSecondaryDns) HasIp() bool {
	return obj.obj.Ip != nil
}

// The secondary DNS server address that is offered to DHCP clients that request this information through a TLV option.
// SetIp sets the string value in the DhcpV6ServerSecondaryDns object
func (obj *dhcpV6ServerSecondaryDns) SetIp(value string) DhcpV6ServerSecondaryDns {

	obj.obj.Ip = &value
	return obj
}

func (obj *dhcpV6ServerSecondaryDns) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ip != nil {

		err := obj.validateIpv6(obj.Ip())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on DhcpV6ServerSecondaryDns.Ip"))
		}

	}

}

func (obj *dhcpV6ServerSecondaryDns) setDefault() {

}
