package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ServerOptions *****
type dhcpv6ServerOptions struct {
	validation
	obj               *otg.Dhcpv6ServerOptions
	marshaller        marshalDhcpv6ServerOptions
	unMarshaller      unMarshalDhcpv6ServerOptions
	dnsHolder         DhcpV6ServerDns
	vendorInfoHolder  Dhcpv6ServerOptionsVendorInfo
	bootfileUrlHolder Dhcpv6ServerOptionsBootfileUrl
}

func NewDhcpv6ServerOptions() Dhcpv6ServerOptions {
	obj := dhcpv6ServerOptions{obj: &otg.Dhcpv6ServerOptions{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ServerOptions) msg() *otg.Dhcpv6ServerOptions {
	return obj.obj
}

func (obj *dhcpv6ServerOptions) setMsg(msg *otg.Dhcpv6ServerOptions) Dhcpv6ServerOptions {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ServerOptions struct {
	obj *dhcpv6ServerOptions
}

type marshalDhcpv6ServerOptions interface {
	// ToProto marshals Dhcpv6ServerOptions to protobuf object *otg.Dhcpv6ServerOptions
	ToProto() (*otg.Dhcpv6ServerOptions, error)
	// ToPbText marshals Dhcpv6ServerOptions to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ServerOptions to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ServerOptions to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Dhcpv6ServerOptions to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldhcpv6ServerOptions struct {
	obj *dhcpv6ServerOptions
}

type unMarshalDhcpv6ServerOptions interface {
	// FromProto unmarshals Dhcpv6ServerOptions from protobuf object *otg.Dhcpv6ServerOptions
	FromProto(msg *otg.Dhcpv6ServerOptions) (Dhcpv6ServerOptions, error)
	// FromPbText unmarshals Dhcpv6ServerOptions from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ServerOptions from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ServerOptions from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ServerOptions) Marshal() marshalDhcpv6ServerOptions {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ServerOptions{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ServerOptions) Unmarshal() unMarshalDhcpv6ServerOptions {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ServerOptions{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ServerOptions) ToProto() (*otg.Dhcpv6ServerOptions, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ServerOptions) FromProto(msg *otg.Dhcpv6ServerOptions) (Dhcpv6ServerOptions, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ServerOptions) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ServerOptions) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshaldhcpv6ServerOptions) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ServerOptions) FromYaml(value string) error {
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
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshaldhcpv6ServerOptions) ToJsonRaw() (string, error) {
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

func (m *marshaldhcpv6ServerOptions) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ServerOptions) FromJson(value string) error {
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
	m.obj.setNil()
	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *dhcpv6ServerOptions) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ServerOptions) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ServerOptions) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ServerOptions) Clone() (Dhcpv6ServerOptions, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ServerOptions()
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

func (obj *dhcpv6ServerOptions) setNil() {
	obj.dnsHolder = nil
	obj.vendorInfoHolder = nil
	obj.bootfileUrlHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Dhcpv6ServerOptions is dHCP server options, these configured options are sent in Dhcp server messages.
type Dhcpv6ServerOptions interface {
	Validation
	// msg marshals Dhcpv6ServerOptions to protobuf object *otg.Dhcpv6ServerOptions
	// and doesn't set defaults
	msg() *otg.Dhcpv6ServerOptions
	// setMsg unmarshals Dhcpv6ServerOptions from protobuf object *otg.Dhcpv6ServerOptions
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ServerOptions) Dhcpv6ServerOptions
	// provides marshal interface
	Marshal() marshalDhcpv6ServerOptions
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ServerOptions
	// validate validates Dhcpv6ServerOptions
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ServerOptions, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Dns returns DhcpV6ServerDns, set in Dhcpv6ServerOptions.
	// DhcpV6ServerDns is optional Dns configuration for DHCPv6 server.
	Dns() DhcpV6ServerDns
	// SetDns assigns DhcpV6ServerDns provided by user to Dhcpv6ServerOptions.
	// DhcpV6ServerDns is optional Dns configuration for DHCPv6 server.
	SetDns(value DhcpV6ServerDns) Dhcpv6ServerOptions
	// HasDns checks if Dns has been set in Dhcpv6ServerOptions
	HasDns() bool
	// VendorInfo returns Dhcpv6ServerOptionsVendorInfo, set in Dhcpv6ServerOptions.
	// Dhcpv6ServerOptionsVendorInfo is this option is used by servers to exchange vendor-specific information. The option code is 17.
	VendorInfo() Dhcpv6ServerOptionsVendorInfo
	// SetVendorInfo assigns Dhcpv6ServerOptionsVendorInfo provided by user to Dhcpv6ServerOptions.
	// Dhcpv6ServerOptionsVendorInfo is this option is used by servers to exchange vendor-specific information. The option code is 17.
	SetVendorInfo(value Dhcpv6ServerOptionsVendorInfo) Dhcpv6ServerOptions
	// HasVendorInfo checks if VendorInfo has been set in Dhcpv6ServerOptions
	HasVendorInfo() bool
	// BootfileUrl returns Dhcpv6ServerOptionsBootfileUrl, set in Dhcpv6ServerOptions.
	// Dhcpv6ServerOptionsBootfileUrl is the server sends this option to inform the client about a URL to a boot file. This information is required for booting  over the network includes the details about the server on which the boot files can be found, the protocol to be used for  the download (for example,HTTP or TFTP, and the path and name of the boot file on the server. The option code is 59. The URL will contain the network communication protocol, a subdomain, a domain name, and its extension. If the host in the URL  is expressed using an IPv6 address rather than a domain name, the address in the URL then must be enclosed in "[" and "]"  characters, conforming to [RFC3986]. Eg of a boot file url can be "tftp://[xxxx:xxxx:xxxx:xxxx::xxxx]/mboot.efi".
	BootfileUrl() Dhcpv6ServerOptionsBootfileUrl
	// SetBootfileUrl assigns Dhcpv6ServerOptionsBootfileUrl provided by user to Dhcpv6ServerOptions.
	// Dhcpv6ServerOptionsBootfileUrl is the server sends this option to inform the client about a URL to a boot file. This information is required for booting  over the network includes the details about the server on which the boot files can be found, the protocol to be used for  the download (for example,HTTP or TFTP, and the path and name of the boot file on the server. The option code is 59. The URL will contain the network communication protocol, a subdomain, a domain name, and its extension. If the host in the URL  is expressed using an IPv6 address rather than a domain name, the address in the URL then must be enclosed in "[" and "]"  characters, conforming to [RFC3986]. Eg of a boot file url can be "tftp://[xxxx:xxxx:xxxx:xxxx::xxxx]/mboot.efi".
	SetBootfileUrl(value Dhcpv6ServerOptionsBootfileUrl) Dhcpv6ServerOptions
	// HasBootfileUrl checks if BootfileUrl has been set in Dhcpv6ServerOptions
	HasBootfileUrl() bool
	setNil()
}

// Additional DHCP server primary dns and other configuration options.
// Dns returns a DhcpV6ServerDns
func (obj *dhcpv6ServerOptions) Dns() DhcpV6ServerDns {
	if obj.obj.Dns == nil {
		obj.obj.Dns = NewDhcpV6ServerDns().msg()
	}
	if obj.dnsHolder == nil {
		obj.dnsHolder = &dhcpV6ServerDns{obj: obj.obj.Dns}
	}
	return obj.dnsHolder
}

// Additional DHCP server primary dns and other configuration options.
// Dns returns a DhcpV6ServerDns
func (obj *dhcpv6ServerOptions) HasDns() bool {
	return obj.obj.Dns != nil
}

// Additional DHCP server primary dns and other configuration options.
// SetDns sets the DhcpV6ServerDns value in the Dhcpv6ServerOptions object
func (obj *dhcpv6ServerOptions) SetDns(value DhcpV6ServerDns) Dhcpv6ServerOptions {

	obj.dnsHolder = nil
	obj.obj.Dns = value.msg()

	return obj
}

// This option is used by servers to exchange vendor-specific information with clients.
// VendorInfo returns a Dhcpv6ServerOptionsVendorInfo
func (obj *dhcpv6ServerOptions) VendorInfo() Dhcpv6ServerOptionsVendorInfo {
	if obj.obj.VendorInfo == nil {
		obj.obj.VendorInfo = NewDhcpv6ServerOptionsVendorInfo().msg()
	}
	if obj.vendorInfoHolder == nil {
		obj.vendorInfoHolder = &dhcpv6ServerOptionsVendorInfo{obj: obj.obj.VendorInfo}
	}
	return obj.vendorInfoHolder
}

// This option is used by servers to exchange vendor-specific information with clients.
// VendorInfo returns a Dhcpv6ServerOptionsVendorInfo
func (obj *dhcpv6ServerOptions) HasVendorInfo() bool {
	return obj.obj.VendorInfo != nil
}

// This option is used by servers to exchange vendor-specific information with clients.
// SetVendorInfo sets the Dhcpv6ServerOptionsVendorInfo value in the Dhcpv6ServerOptions object
func (obj *dhcpv6ServerOptions) SetVendorInfo(value Dhcpv6ServerOptionsVendorInfo) Dhcpv6ServerOptions {

	obj.vendorInfoHolder = nil
	obj.obj.VendorInfo = value.msg()

	return obj
}

// The server sends this option to inform the client about a URL to a boot file which client will use for
// network boots.
// BootfileUrl returns a Dhcpv6ServerOptionsBootfileUrl
func (obj *dhcpv6ServerOptions) BootfileUrl() Dhcpv6ServerOptionsBootfileUrl {
	if obj.obj.BootfileUrl == nil {
		obj.obj.BootfileUrl = NewDhcpv6ServerOptionsBootfileUrl().msg()
	}
	if obj.bootfileUrlHolder == nil {
		obj.bootfileUrlHolder = &dhcpv6ServerOptionsBootfileUrl{obj: obj.obj.BootfileUrl}
	}
	return obj.bootfileUrlHolder
}

// The server sends this option to inform the client about a URL to a boot file which client will use for
// network boots.
// BootfileUrl returns a Dhcpv6ServerOptionsBootfileUrl
func (obj *dhcpv6ServerOptions) HasBootfileUrl() bool {
	return obj.obj.BootfileUrl != nil
}

// The server sends this option to inform the client about a URL to a boot file which client will use for
// network boots.
// SetBootfileUrl sets the Dhcpv6ServerOptionsBootfileUrl value in the Dhcpv6ServerOptions object
func (obj *dhcpv6ServerOptions) SetBootfileUrl(value Dhcpv6ServerOptionsBootfileUrl) Dhcpv6ServerOptions {

	obj.bootfileUrlHolder = nil
	obj.obj.BootfileUrl = value.msg()

	return obj
}

func (obj *dhcpv6ServerOptions) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Dns != nil {

		obj.Dns().validateObj(vObj, set_default)
	}

	if obj.obj.VendorInfo != nil {

		obj.VendorInfo().validateObj(vObj, set_default)
	}

	if obj.obj.BootfileUrl != nil {

		obj.BootfileUrl().validateObj(vObj, set_default)
	}

}

func (obj *dhcpv6ServerOptions) setDefault() {

}
