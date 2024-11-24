package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceDhcpv6ClientOptions *****
type deviceDhcpv6ClientOptions struct {
	validation
	obj                    *otg.DeviceDhcpv6ClientOptions
	marshaller             marshalDeviceDhcpv6ClientOptions
	unMarshaller           unMarshalDeviceDhcpv6ClientOptions
	serverIdentifierHolder Dhcpv6ClientOptionsServerIdentifier
	vendorClassHolder      Dhcpv6ClientOptionsVendorClass
	vendorInfoHolder       Dhcpv6ClientOptionsVendorInfo
	fqdnHolder             Dhcpv6ClientOptionsFqdn
}

func NewDeviceDhcpv6ClientOptions() DeviceDhcpv6ClientOptions {
	obj := deviceDhcpv6ClientOptions{obj: &otg.DeviceDhcpv6ClientOptions{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceDhcpv6ClientOptions) msg() *otg.DeviceDhcpv6ClientOptions {
	return obj.obj
}

func (obj *deviceDhcpv6ClientOptions) setMsg(msg *otg.DeviceDhcpv6ClientOptions) DeviceDhcpv6ClientOptions {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceDhcpv6ClientOptions struct {
	obj *deviceDhcpv6ClientOptions
}

type marshalDeviceDhcpv6ClientOptions interface {
	// ToProto marshals DeviceDhcpv6ClientOptions to protobuf object *otg.DeviceDhcpv6ClientOptions
	ToProto() (*otg.DeviceDhcpv6ClientOptions, error)
	// ToPbText marshals DeviceDhcpv6ClientOptions to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceDhcpv6ClientOptions to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceDhcpv6ClientOptions to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceDhcpv6ClientOptions struct {
	obj *deviceDhcpv6ClientOptions
}

type unMarshalDeviceDhcpv6ClientOptions interface {
	// FromProto unmarshals DeviceDhcpv6ClientOptions from protobuf object *otg.DeviceDhcpv6ClientOptions
	FromProto(msg *otg.DeviceDhcpv6ClientOptions) (DeviceDhcpv6ClientOptions, error)
	// FromPbText unmarshals DeviceDhcpv6ClientOptions from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceDhcpv6ClientOptions from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceDhcpv6ClientOptions from JSON text
	FromJson(value string) error
}

func (obj *deviceDhcpv6ClientOptions) Marshal() marshalDeviceDhcpv6ClientOptions {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceDhcpv6ClientOptions{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceDhcpv6ClientOptions) Unmarshal() unMarshalDeviceDhcpv6ClientOptions {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceDhcpv6ClientOptions{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceDhcpv6ClientOptions) ToProto() (*otg.DeviceDhcpv6ClientOptions, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceDhcpv6ClientOptions) FromProto(msg *otg.DeviceDhcpv6ClientOptions) (DeviceDhcpv6ClientOptions, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceDhcpv6ClientOptions) ToPbText() (string, error) {
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

func (m *unMarshaldeviceDhcpv6ClientOptions) FromPbText(value string) error {
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

func (m *marshaldeviceDhcpv6ClientOptions) ToYaml() (string, error) {
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

func (m *unMarshaldeviceDhcpv6ClientOptions) FromYaml(value string) error {
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

func (m *marshaldeviceDhcpv6ClientOptions) ToJson() (string, error) {
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

func (m *unMarshaldeviceDhcpv6ClientOptions) FromJson(value string) error {
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

func (obj *deviceDhcpv6ClientOptions) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceDhcpv6ClientOptions) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceDhcpv6ClientOptions) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceDhcpv6ClientOptions) Clone() (DeviceDhcpv6ClientOptions, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceDhcpv6ClientOptions()
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

func (obj *deviceDhcpv6ClientOptions) setNil() {
	obj.serverIdentifierHolder = nil
	obj.vendorClassHolder = nil
	obj.vendorInfoHolder = nil
	obj.fqdnHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceDhcpv6ClientOptions is dHCP client options, these configured options are sent in Dhcp client messages.
type DeviceDhcpv6ClientOptions interface {
	Validation
	// msg marshals DeviceDhcpv6ClientOptions to protobuf object *otg.DeviceDhcpv6ClientOptions
	// and doesn't set defaults
	msg() *otg.DeviceDhcpv6ClientOptions
	// setMsg unmarshals DeviceDhcpv6ClientOptions from protobuf object *otg.DeviceDhcpv6ClientOptions
	// and doesn't set defaults
	setMsg(*otg.DeviceDhcpv6ClientOptions) DeviceDhcpv6ClientOptions
	// provides marshal interface
	Marshal() marshalDeviceDhcpv6ClientOptions
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceDhcpv6ClientOptions
	// validate validates DeviceDhcpv6ClientOptions
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceDhcpv6ClientOptions, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ServerIdentifier returns Dhcpv6ClientOptionsServerIdentifier, set in DeviceDhcpv6ClientOptions.
	// Dhcpv6ClientOptionsServerIdentifier is description is TBD
	ServerIdentifier() Dhcpv6ClientOptionsServerIdentifier
	// SetServerIdentifier assigns Dhcpv6ClientOptionsServerIdentifier provided by user to DeviceDhcpv6ClientOptions.
	// Dhcpv6ClientOptionsServerIdentifier is description is TBD
	SetServerIdentifier(value Dhcpv6ClientOptionsServerIdentifier) DeviceDhcpv6ClientOptions
	// HasServerIdentifier checks if ServerIdentifier has been set in DeviceDhcpv6ClientOptions
	HasServerIdentifier() bool
	// VendorClass returns Dhcpv6ClientOptionsVendorClass, set in DeviceDhcpv6ClientOptions.
	// Dhcpv6ClientOptionsVendorClass is this option is used by a client to identify the vendor that manufactured the hardware on which the client is running. The option code is 16.
	VendorClass() Dhcpv6ClientOptionsVendorClass
	// SetVendorClass assigns Dhcpv6ClientOptionsVendorClass provided by user to DeviceDhcpv6ClientOptions.
	// Dhcpv6ClientOptionsVendorClass is this option is used by a client to identify the vendor that manufactured the hardware on which the client is running. The option code is 16.
	SetVendorClass(value Dhcpv6ClientOptionsVendorClass) DeviceDhcpv6ClientOptions
	// HasVendorClass checks if VendorClass has been set in DeviceDhcpv6ClientOptions
	HasVendorClass() bool
	// VendorInfo returns Dhcpv6ClientOptionsVendorInfo, set in DeviceDhcpv6ClientOptions.
	// Dhcpv6ClientOptionsVendorInfo is this option is used by clients to exchange vendor-specific information. The option code is 17.
	VendorInfo() Dhcpv6ClientOptionsVendorInfo
	// SetVendorInfo assigns Dhcpv6ClientOptionsVendorInfo provided by user to DeviceDhcpv6ClientOptions.
	// Dhcpv6ClientOptionsVendorInfo is this option is used by clients to exchange vendor-specific information. The option code is 17.
	SetVendorInfo(value Dhcpv6ClientOptionsVendorInfo) DeviceDhcpv6ClientOptions
	// HasVendorInfo checks if VendorInfo has been set in DeviceDhcpv6ClientOptions
	HasVendorInfo() bool
	// Fqdn returns Dhcpv6ClientOptionsFqdn, set in DeviceDhcpv6ClientOptions.
	// Dhcpv6ClientOptionsFqdn is dHCPv6 server needs to know the FQDN of the client for the addresses for the client's IA_NA bindings in order to update the IPv6-address-to-FQDN mapping. This option allows the client to convey its FQDN to the server. The Client  FQDN option also contains Flags that DHCPv6 clients and servers use to negotiate who does which updates. The option code is 39.
	Fqdn() Dhcpv6ClientOptionsFqdn
	// SetFqdn assigns Dhcpv6ClientOptionsFqdn provided by user to DeviceDhcpv6ClientOptions.
	// Dhcpv6ClientOptionsFqdn is dHCPv6 server needs to know the FQDN of the client for the addresses for the client's IA_NA bindings in order to update the IPv6-address-to-FQDN mapping. This option allows the client to convey its FQDN to the server. The Client  FQDN option also contains Flags that DHCPv6 clients and servers use to negotiate who does which updates. The option code is 39.
	SetFqdn(value Dhcpv6ClientOptionsFqdn) DeviceDhcpv6ClientOptions
	// HasFqdn checks if Fqdn has been set in DeviceDhcpv6ClientOptions
	HasFqdn() bool
	setNil()
}

// A client uses multicast to reach all servers or an individual server. An individual server is indicated by
// specifying that server's DUID in a Server Identifier option in the client's message (all servers will receive
// this message but only the indicated server will respond). All servers are indicated by not supplying this option.
// ServerIdentifier returns a Dhcpv6ClientOptionsServerIdentifier
func (obj *deviceDhcpv6ClientOptions) ServerIdentifier() Dhcpv6ClientOptionsServerIdentifier {
	if obj.obj.ServerIdentifier == nil {
		obj.obj.ServerIdentifier = NewDhcpv6ClientOptionsServerIdentifier().msg()
	}
	if obj.serverIdentifierHolder == nil {
		obj.serverIdentifierHolder = &dhcpv6ClientOptionsServerIdentifier{obj: obj.obj.ServerIdentifier}
	}
	return obj.serverIdentifierHolder
}

// A client uses multicast to reach all servers or an individual server. An individual server is indicated by
// specifying that server's DUID in a Server Identifier option in the client's message (all servers will receive
// this message but only the indicated server will respond). All servers are indicated by not supplying this option.
// ServerIdentifier returns a Dhcpv6ClientOptionsServerIdentifier
func (obj *deviceDhcpv6ClientOptions) HasServerIdentifier() bool {
	return obj.obj.ServerIdentifier != nil
}

// A client uses multicast to reach all servers or an individual server. An individual server is indicated by
// specifying that server's DUID in a Server Identifier option in the client's message (all servers will receive
// this message but only the indicated server will respond). All servers are indicated by not supplying this option.
// SetServerIdentifier sets the Dhcpv6ClientOptionsServerIdentifier value in the DeviceDhcpv6ClientOptions object
func (obj *deviceDhcpv6ClientOptions) SetServerIdentifier(value Dhcpv6ClientOptionsServerIdentifier) DeviceDhcpv6ClientOptions {

	obj.serverIdentifierHolder = nil
	obj.obj.ServerIdentifier = value.msg()

	return obj
}

// The vendor class option is used by a client to identify the vendor that manufactured the hardware on which
// the client is running. The information contained in the data area of this option is contained in one or more
// opaque fields that identify details of the hardware configuration.
// VendorClass returns a Dhcpv6ClientOptionsVendorClass
func (obj *deviceDhcpv6ClientOptions) VendorClass() Dhcpv6ClientOptionsVendorClass {
	if obj.obj.VendorClass == nil {
		obj.obj.VendorClass = NewDhcpv6ClientOptionsVendorClass().msg()
	}
	if obj.vendorClassHolder == nil {
		obj.vendorClassHolder = &dhcpv6ClientOptionsVendorClass{obj: obj.obj.VendorClass}
	}
	return obj.vendorClassHolder
}

// The vendor class option is used by a client to identify the vendor that manufactured the hardware on which
// the client is running. The information contained in the data area of this option is contained in one or more
// opaque fields that identify details of the hardware configuration.
// VendorClass returns a Dhcpv6ClientOptionsVendorClass
func (obj *deviceDhcpv6ClientOptions) HasVendorClass() bool {
	return obj.obj.VendorClass != nil
}

// The vendor class option is used by a client to identify the vendor that manufactured the hardware on which
// the client is running. The information contained in the data area of this option is contained in one or more
// opaque fields that identify details of the hardware configuration.
// SetVendorClass sets the Dhcpv6ClientOptionsVendorClass value in the DeviceDhcpv6ClientOptions object
func (obj *deviceDhcpv6ClientOptions) SetVendorClass(value Dhcpv6ClientOptionsVendorClass) DeviceDhcpv6ClientOptions {

	obj.vendorClassHolder = nil
	obj.obj.VendorClass = value.msg()

	return obj
}

// This option is used by clients to exchange vendor-specific information with servers.
// VendorInfo returns a Dhcpv6ClientOptionsVendorInfo
func (obj *deviceDhcpv6ClientOptions) VendorInfo() Dhcpv6ClientOptionsVendorInfo {
	if obj.obj.VendorInfo == nil {
		obj.obj.VendorInfo = NewDhcpv6ClientOptionsVendorInfo().msg()
	}
	if obj.vendorInfoHolder == nil {
		obj.vendorInfoHolder = &dhcpv6ClientOptionsVendorInfo{obj: obj.obj.VendorInfo}
	}
	return obj.vendorInfoHolder
}

// This option is used by clients to exchange vendor-specific information with servers.
// VendorInfo returns a Dhcpv6ClientOptionsVendorInfo
func (obj *deviceDhcpv6ClientOptions) HasVendorInfo() bool {
	return obj.obj.VendorInfo != nil
}

// This option is used by clients to exchange vendor-specific information with servers.
// SetVendorInfo sets the Dhcpv6ClientOptionsVendorInfo value in the DeviceDhcpv6ClientOptions object
func (obj *deviceDhcpv6ClientOptions) SetVendorInfo(value Dhcpv6ClientOptionsVendorInfo) DeviceDhcpv6ClientOptions {

	obj.vendorInfoHolder = nil
	obj.obj.VendorInfo = value.msg()

	return obj
}

// DHCPv6 server needs to know the FQDN of the client for the addresses for the client's IA_NA bindings in order to
// update the IPv6-address-to-FQDN mapping. This option allows the client to convey its FQDN to the server. The Client
// FQDN option also contains Flags that DHCPv6 clients and servers use to negotiate who does which update.
// Fqdn returns a Dhcpv6ClientOptionsFqdn
func (obj *deviceDhcpv6ClientOptions) Fqdn() Dhcpv6ClientOptionsFqdn {
	if obj.obj.Fqdn == nil {
		obj.obj.Fqdn = NewDhcpv6ClientOptionsFqdn().msg()
	}
	if obj.fqdnHolder == nil {
		obj.fqdnHolder = &dhcpv6ClientOptionsFqdn{obj: obj.obj.Fqdn}
	}
	return obj.fqdnHolder
}

// DHCPv6 server needs to know the FQDN of the client for the addresses for the client's IA_NA bindings in order to
// update the IPv6-address-to-FQDN mapping. This option allows the client to convey its FQDN to the server. The Client
// FQDN option also contains Flags that DHCPv6 clients and servers use to negotiate who does which update.
// Fqdn returns a Dhcpv6ClientOptionsFqdn
func (obj *deviceDhcpv6ClientOptions) HasFqdn() bool {
	return obj.obj.Fqdn != nil
}

// DHCPv6 server needs to know the FQDN of the client for the addresses for the client's IA_NA bindings in order to
// update the IPv6-address-to-FQDN mapping. This option allows the client to convey its FQDN to the server. The Client
// FQDN option also contains Flags that DHCPv6 clients and servers use to negotiate who does which update.
// SetFqdn sets the Dhcpv6ClientOptionsFqdn value in the DeviceDhcpv6ClientOptions object
func (obj *deviceDhcpv6ClientOptions) SetFqdn(value Dhcpv6ClientOptionsFqdn) DeviceDhcpv6ClientOptions {

	obj.fqdnHolder = nil
	obj.obj.Fqdn = value.msg()

	return obj
}

func (obj *deviceDhcpv6ClientOptions) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ServerIdentifier != nil {

		obj.ServerIdentifier().validateObj(vObj, set_default)
	}

	if obj.obj.VendorClass != nil {

		obj.VendorClass().validateObj(vObj, set_default)
	}

	if obj.obj.VendorInfo != nil {

		obj.VendorInfo().validateObj(vObj, set_default)
	}

	if obj.obj.Fqdn != nil {

		obj.Fqdn().validateObj(vObj, set_default)
	}

}

func (obj *deviceDhcpv6ClientOptions) setDefault() {

}
