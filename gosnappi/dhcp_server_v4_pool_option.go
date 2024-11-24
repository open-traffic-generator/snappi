package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DhcpServerV4PoolOption *****
type dhcpServerV4PoolOption struct {
	validation
	obj          *otg.DhcpServerV4PoolOption
	marshaller   marshalDhcpServerV4PoolOption
	unMarshaller unMarshalDhcpServerV4PoolOption
}

func NewDhcpServerV4PoolOption() DhcpServerV4PoolOption {
	obj := dhcpServerV4PoolOption{obj: &otg.DhcpServerV4PoolOption{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpServerV4PoolOption) msg() *otg.DhcpServerV4PoolOption {
	return obj.obj
}

func (obj *dhcpServerV4PoolOption) setMsg(msg *otg.DhcpServerV4PoolOption) DhcpServerV4PoolOption {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpServerV4PoolOption struct {
	obj *dhcpServerV4PoolOption
}

type marshalDhcpServerV4PoolOption interface {
	// ToProto marshals DhcpServerV4PoolOption to protobuf object *otg.DhcpServerV4PoolOption
	ToProto() (*otg.DhcpServerV4PoolOption, error)
	// ToPbText marshals DhcpServerV4PoolOption to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DhcpServerV4PoolOption to YAML text
	ToYaml() (string, error)
	// ToJson marshals DhcpServerV4PoolOption to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpServerV4PoolOption struct {
	obj *dhcpServerV4PoolOption
}

type unMarshalDhcpServerV4PoolOption interface {
	// FromProto unmarshals DhcpServerV4PoolOption from protobuf object *otg.DhcpServerV4PoolOption
	FromProto(msg *otg.DhcpServerV4PoolOption) (DhcpServerV4PoolOption, error)
	// FromPbText unmarshals DhcpServerV4PoolOption from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DhcpServerV4PoolOption from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DhcpServerV4PoolOption from JSON text
	FromJson(value string) error
}

func (obj *dhcpServerV4PoolOption) Marshal() marshalDhcpServerV4PoolOption {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpServerV4PoolOption{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpServerV4PoolOption) Unmarshal() unMarshalDhcpServerV4PoolOption {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpServerV4PoolOption{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpServerV4PoolOption) ToProto() (*otg.DhcpServerV4PoolOption, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpServerV4PoolOption) FromProto(msg *otg.DhcpServerV4PoolOption) (DhcpServerV4PoolOption, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpServerV4PoolOption) ToPbText() (string, error) {
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

func (m *unMarshaldhcpServerV4PoolOption) FromPbText(value string) error {
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

func (m *marshaldhcpServerV4PoolOption) ToYaml() (string, error) {
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

func (m *unMarshaldhcpServerV4PoolOption) FromYaml(value string) error {
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

func (m *marshaldhcpServerV4PoolOption) ToJson() (string, error) {
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

func (m *unMarshaldhcpServerV4PoolOption) FromJson(value string) error {
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

func (obj *dhcpServerV4PoolOption) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpServerV4PoolOption) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpServerV4PoolOption) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpServerV4PoolOption) Clone() (DhcpServerV4PoolOption, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpServerV4PoolOption()
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

// DhcpServerV4PoolOption is optional configuration for DHCPv4 address pool for the lease.
type DhcpServerV4PoolOption interface {
	Validation
	// msg marshals DhcpServerV4PoolOption to protobuf object *otg.DhcpServerV4PoolOption
	// and doesn't set defaults
	msg() *otg.DhcpServerV4PoolOption
	// setMsg unmarshals DhcpServerV4PoolOption from protobuf object *otg.DhcpServerV4PoolOption
	// and doesn't set defaults
	setMsg(*otg.DhcpServerV4PoolOption) DhcpServerV4PoolOption
	// provides marshal interface
	Marshal() marshalDhcpServerV4PoolOption
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpServerV4PoolOption
	// validate validates DhcpServerV4PoolOption
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DhcpServerV4PoolOption, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterAddress returns string, set in DhcpServerV4PoolOption.
	RouterAddress() string
	// SetRouterAddress assigns string provided by user to DhcpServerV4PoolOption
	SetRouterAddress(value string) DhcpServerV4PoolOption
	// HasRouterAddress checks if RouterAddress has been set in DhcpServerV4PoolOption
	HasRouterAddress() bool
	// PrimaryDnsServer returns string, set in DhcpServerV4PoolOption.
	PrimaryDnsServer() string
	// SetPrimaryDnsServer assigns string provided by user to DhcpServerV4PoolOption
	SetPrimaryDnsServer(value string) DhcpServerV4PoolOption
	// HasPrimaryDnsServer checks if PrimaryDnsServer has been set in DhcpServerV4PoolOption
	HasPrimaryDnsServer() bool
	// SecondaryDnsServer returns string, set in DhcpServerV4PoolOption.
	SecondaryDnsServer() string
	// SetSecondaryDnsServer assigns string provided by user to DhcpServerV4PoolOption
	SetSecondaryDnsServer(value string) DhcpServerV4PoolOption
	// HasSecondaryDnsServer checks if SecondaryDnsServer has been set in DhcpServerV4PoolOption
	HasSecondaryDnsServer() bool
	// EchoRelayWithTlv82 returns bool, set in DhcpServerV4PoolOption.
	EchoRelayWithTlv82() bool
	// SetEchoRelayWithTlv82 assigns bool provided by user to DhcpServerV4PoolOption
	SetEchoRelayWithTlv82(value bool) DhcpServerV4PoolOption
	// HasEchoRelayWithTlv82 checks if EchoRelayWithTlv82 has been set in DhcpServerV4PoolOption
	HasEchoRelayWithTlv82() bool
}

// The Router address advertised by the DHCPv4 server in Offer and Ack messages.
// RouterAddress returns a string
func (obj *dhcpServerV4PoolOption) RouterAddress() string {

	return *obj.obj.RouterAddress

}

// The Router address advertised by the DHCPv4 server in Offer and Ack messages.
// RouterAddress returns a string
func (obj *dhcpServerV4PoolOption) HasRouterAddress() bool {
	return obj.obj.RouterAddress != nil
}

// The Router address advertised by the DHCPv4 server in Offer and Ack messages.
// SetRouterAddress sets the string value in the DhcpServerV4PoolOption object
func (obj *dhcpServerV4PoolOption) SetRouterAddress(value string) DhcpServerV4PoolOption {

	obj.obj.RouterAddress = &value
	return obj
}

// The primary DNS server address that is offered to DHCP clients that request this information through a TLV option.
// PrimaryDnsServer returns a string
func (obj *dhcpServerV4PoolOption) PrimaryDnsServer() string {

	return *obj.obj.PrimaryDnsServer

}

// The primary DNS server address that is offered to DHCP clients that request this information through a TLV option.
// PrimaryDnsServer returns a string
func (obj *dhcpServerV4PoolOption) HasPrimaryDnsServer() bool {
	return obj.obj.PrimaryDnsServer != nil
}

// The primary DNS server address that is offered to DHCP clients that request this information through a TLV option.
// SetPrimaryDnsServer sets the string value in the DhcpServerV4PoolOption object
func (obj *dhcpServerV4PoolOption) SetPrimaryDnsServer(value string) DhcpServerV4PoolOption {

	obj.obj.PrimaryDnsServer = &value
	return obj
}

// The primary DNS server address that is offered to DHCP clients that request this information through a TLV option.
// SecondaryDnsServer returns a string
func (obj *dhcpServerV4PoolOption) SecondaryDnsServer() string {

	return *obj.obj.SecondaryDnsServer

}

// The primary DNS server address that is offered to DHCP clients that request this information through a TLV option.
// SecondaryDnsServer returns a string
func (obj *dhcpServerV4PoolOption) HasSecondaryDnsServer() bool {
	return obj.obj.SecondaryDnsServer != nil
}

// The primary DNS server address that is offered to DHCP clients that request this information through a TLV option.
// SetSecondaryDnsServer sets the string value in the DhcpServerV4PoolOption object
func (obj *dhcpServerV4PoolOption) SetSecondaryDnsServer(value string) DhcpServerV4PoolOption {

	obj.obj.SecondaryDnsServer = &value
	return obj
}

// If selected, the DHCP server includes in its replies the TLV information for the DHCPv4 Relay Agent Option 82 and the corresponding sub-TLVs that it receives from a DHCP relay agent, otherwise it replies without including this TLV.
// EchoRelayWithTlv82 returns a bool
func (obj *dhcpServerV4PoolOption) EchoRelayWithTlv82() bool {

	return *obj.obj.EchoRelayWithTlv_82

}

// If selected, the DHCP server includes in its replies the TLV information for the DHCPv4 Relay Agent Option 82 and the corresponding sub-TLVs that it receives from a DHCP relay agent, otherwise it replies without including this TLV.
// EchoRelayWithTlv82 returns a bool
func (obj *dhcpServerV4PoolOption) HasEchoRelayWithTlv82() bool {
	return obj.obj.EchoRelayWithTlv_82 != nil
}

// If selected, the DHCP server includes in its replies the TLV information for the DHCPv4 Relay Agent Option 82 and the corresponding sub-TLVs that it receives from a DHCP relay agent, otherwise it replies without including this TLV.
// SetEchoRelayWithTlv82 sets the bool value in the DhcpServerV4PoolOption object
func (obj *dhcpServerV4PoolOption) SetEchoRelayWithTlv82(value bool) DhcpServerV4PoolOption {

	obj.obj.EchoRelayWithTlv_82 = &value
	return obj
}

func (obj *dhcpServerV4PoolOption) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RouterAddress != nil {

		err := obj.validateIpv4(obj.RouterAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on DhcpServerV4PoolOption.RouterAddress"))
		}

	}

	if obj.obj.PrimaryDnsServer != nil {

		err := obj.validateIpv4(obj.PrimaryDnsServer())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on DhcpServerV4PoolOption.PrimaryDnsServer"))
		}

	}

	if obj.obj.SecondaryDnsServer != nil {

		err := obj.validateIpv4(obj.SecondaryDnsServer())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on DhcpServerV4PoolOption.SecondaryDnsServer"))
		}

	}

}

func (obj *dhcpServerV4PoolOption) setDefault() {
	if obj.obj.RouterAddress == nil {
		obj.SetRouterAddress("0.0.0.0")
	}
	if obj.obj.PrimaryDnsServer == nil {
		obj.SetPrimaryDnsServer("0.0.0.0")
	}
	if obj.obj.SecondaryDnsServer == nil {
		obj.SetSecondaryDnsServer("0.0.0.0")
	}
	if obj.obj.EchoRelayWithTlv_82 == nil {
		obj.SetEchoRelayWithTlv82(true)
	}

}
