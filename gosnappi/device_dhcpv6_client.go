package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceDhcpv6Client *****
type deviceDhcpv6Client struct {
	validation
	obj                    *otg.DeviceDhcpv6Client
	marshaller             marshalDeviceDhcpv6Client
	unMarshaller           unMarshalDeviceDhcpv6Client
	serverIdentifierHolder Dhcpv6ClientOptionsServerIdentifier
	iaTypeHolder           DeviceDhcpv6ClientIaType
	duidTypeHolder         DeviceDhcpv6ClientDuidType
	optionsHolder          DeviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter
	vendorClassHolder      Dhcpv6ClientOptionsVendorClass
	vendorInfoHolder       Dhcpv6ClientOptionsVendorInfo
	fqdnHolder             Dhcpv6ClientOptionsFqdn
}

func NewDeviceDhcpv6Client() DeviceDhcpv6Client {
	obj := deviceDhcpv6Client{obj: &otg.DeviceDhcpv6Client{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceDhcpv6Client) msg() *otg.DeviceDhcpv6Client {
	return obj.obj
}

func (obj *deviceDhcpv6Client) setMsg(msg *otg.DeviceDhcpv6Client) DeviceDhcpv6Client {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceDhcpv6Client struct {
	obj *deviceDhcpv6Client
}

type marshalDeviceDhcpv6Client interface {
	// ToProto marshals DeviceDhcpv6Client to protobuf object *otg.DeviceDhcpv6Client
	ToProto() (*otg.DeviceDhcpv6Client, error)
	// ToPbText marshals DeviceDhcpv6Client to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceDhcpv6Client to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceDhcpv6Client to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceDhcpv6Client struct {
	obj *deviceDhcpv6Client
}

type unMarshalDeviceDhcpv6Client interface {
	// FromProto unmarshals DeviceDhcpv6Client from protobuf object *otg.DeviceDhcpv6Client
	FromProto(msg *otg.DeviceDhcpv6Client) (DeviceDhcpv6Client, error)
	// FromPbText unmarshals DeviceDhcpv6Client from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceDhcpv6Client from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceDhcpv6Client from JSON text
	FromJson(value string) error
}

func (obj *deviceDhcpv6Client) Marshal() marshalDeviceDhcpv6Client {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceDhcpv6Client{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceDhcpv6Client) Unmarshal() unMarshalDeviceDhcpv6Client {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceDhcpv6Client{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceDhcpv6Client) ToProto() (*otg.DeviceDhcpv6Client, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceDhcpv6Client) FromProto(msg *otg.DeviceDhcpv6Client) (DeviceDhcpv6Client, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceDhcpv6Client) ToPbText() (string, error) {
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

func (m *unMarshaldeviceDhcpv6Client) FromPbText(value string) error {
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

func (m *marshaldeviceDhcpv6Client) ToYaml() (string, error) {
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

func (m *unMarshaldeviceDhcpv6Client) FromYaml(value string) error {
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

func (m *marshaldeviceDhcpv6Client) ToJson() (string, error) {
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

func (m *unMarshaldeviceDhcpv6Client) FromJson(value string) error {
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

func (obj *deviceDhcpv6Client) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceDhcpv6Client) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceDhcpv6Client) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceDhcpv6Client) Clone() (DeviceDhcpv6Client, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceDhcpv6Client()
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

func (obj *deviceDhcpv6Client) setNil() {
	obj.serverIdentifierHolder = nil
	obj.iaTypeHolder = nil
	obj.duidTypeHolder = nil
	obj.optionsHolder = nil
	obj.vendorClassHolder = nil
	obj.vendorInfoHolder = nil
	obj.fqdnHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceDhcpv6Client is configuration for emulated DHCPv6 Client on a single Interface. If the DHCPv6 Client receives one or more  DHCPv6 ADVERTISE messages from one or more servers then the client chooses one server from which to request  configuration parameters, based on the configuration parameters offered by the server in the DHCPv6 ADVERTISE  messages. If all configuration parameters match then the first server will be chosen. https://www.rfc-editor.org/rfc/rfc8415.html
type DeviceDhcpv6Client interface {
	Validation
	// msg marshals DeviceDhcpv6Client to protobuf object *otg.DeviceDhcpv6Client
	// and doesn't set defaults
	msg() *otg.DeviceDhcpv6Client
	// setMsg unmarshals DeviceDhcpv6Client from protobuf object *otg.DeviceDhcpv6Client
	// and doesn't set defaults
	setMsg(*otg.DeviceDhcpv6Client) DeviceDhcpv6Client
	// provides marshal interface
	Marshal() marshalDeviceDhcpv6Client
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceDhcpv6Client
	// validate validates DeviceDhcpv6Client
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceDhcpv6Client, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in DeviceDhcpv6Client.
	Name() string
	// SetName assigns string provided by user to DeviceDhcpv6Client
	SetName(value string) DeviceDhcpv6Client
	// RapidCommit returns bool, set in DeviceDhcpv6Client.
	RapidCommit() bool
	// SetRapidCommit assigns bool provided by user to DeviceDhcpv6Client
	SetRapidCommit(value bool) DeviceDhcpv6Client
	// HasRapidCommit checks if RapidCommit has been set in DeviceDhcpv6Client
	HasRapidCommit() bool
	// ServerIdentifier returns Dhcpv6ClientOptionsServerIdentifier, set in DeviceDhcpv6Client.
	// Dhcpv6ClientOptionsServerIdentifier is description is TBD
	ServerIdentifier() Dhcpv6ClientOptionsServerIdentifier
	// SetServerIdentifier assigns Dhcpv6ClientOptionsServerIdentifier provided by user to DeviceDhcpv6Client.
	// Dhcpv6ClientOptionsServerIdentifier is description is TBD
	SetServerIdentifier(value Dhcpv6ClientOptionsServerIdentifier) DeviceDhcpv6Client
	// HasServerIdentifier checks if ServerIdentifier has been set in DeviceDhcpv6Client
	HasServerIdentifier() bool
	// IaType returns DeviceDhcpv6ClientIaType, set in DeviceDhcpv6Client.
	IaType() DeviceDhcpv6ClientIaType
	// SetIaType assigns DeviceDhcpv6ClientIaType provided by user to DeviceDhcpv6Client.
	SetIaType(value DeviceDhcpv6ClientIaType) DeviceDhcpv6Client
	// HasIaType checks if IaType has been set in DeviceDhcpv6Client
	HasIaType() bool
	// DuidType returns DeviceDhcpv6ClientDuidType, set in DeviceDhcpv6Client.
	DuidType() DeviceDhcpv6ClientDuidType
	// SetDuidType assigns DeviceDhcpv6ClientDuidType provided by user to DeviceDhcpv6Client.
	SetDuidType(value DeviceDhcpv6ClientDuidType) DeviceDhcpv6Client
	// HasDuidType checks if DuidType has been set in DeviceDhcpv6Client
	HasDuidType() bool
	// Options returns DeviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIterIter, set in DeviceDhcpv6Client
	Options() DeviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter
	// VendorClass returns Dhcpv6ClientOptionsVendorClass, set in DeviceDhcpv6Client.
	// Dhcpv6ClientOptionsVendorClass is this option is used by a client to identify the vendor that manufactured the hardware on which the client is running. The option code is 16.
	VendorClass() Dhcpv6ClientOptionsVendorClass
	// SetVendorClass assigns Dhcpv6ClientOptionsVendorClass provided by user to DeviceDhcpv6Client.
	// Dhcpv6ClientOptionsVendorClass is this option is used by a client to identify the vendor that manufactured the hardware on which the client is running. The option code is 16.
	SetVendorClass(value Dhcpv6ClientOptionsVendorClass) DeviceDhcpv6Client
	// HasVendorClass checks if VendorClass has been set in DeviceDhcpv6Client
	HasVendorClass() bool
	// VendorInfo returns Dhcpv6ClientOptionsVendorInfo, set in DeviceDhcpv6Client.
	// Dhcpv6ClientOptionsVendorInfo is this option is used by clients to exchange vendor-specific information. The option code is 17.
	VendorInfo() Dhcpv6ClientOptionsVendorInfo
	// SetVendorInfo assigns Dhcpv6ClientOptionsVendorInfo provided by user to DeviceDhcpv6Client.
	// Dhcpv6ClientOptionsVendorInfo is this option is used by clients to exchange vendor-specific information. The option code is 17.
	SetVendorInfo(value Dhcpv6ClientOptionsVendorInfo) DeviceDhcpv6Client
	// HasVendorInfo checks if VendorInfo has been set in DeviceDhcpv6Client
	HasVendorInfo() bool
	// Fqdn returns Dhcpv6ClientOptionsFqdn, set in DeviceDhcpv6Client.
	// Dhcpv6ClientOptionsFqdn is dHCPv6 server needs to know the FQDN of the client for the addresses for the client's IA_NA bindings in order to update the IPv6-address-to-FQDN mapping. This option allows the client to convey its FQDN to the server. The Client  FQDN option also contains Flags that DHCPv6 clients and servers use to negotiate who does which updates. The option code is 39.
	Fqdn() Dhcpv6ClientOptionsFqdn
	// SetFqdn assigns Dhcpv6ClientOptionsFqdn provided by user to DeviceDhcpv6Client.
	// Dhcpv6ClientOptionsFqdn is dHCPv6 server needs to know the FQDN of the client for the addresses for the client's IA_NA bindings in order to update the IPv6-address-to-FQDN mapping. This option allows the client to convey its FQDN to the server. The Client  FQDN option also contains Flags that DHCPv6 clients and servers use to negotiate who does which updates. The option code is 39.
	SetFqdn(value Dhcpv6ClientOptionsFqdn) DeviceDhcpv6Client
	// HasFqdn checks if Fqdn has been set in DeviceDhcpv6Client
	HasFqdn() bool
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *deviceDhcpv6Client) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the DeviceDhcpv6Client object
func (obj *deviceDhcpv6Client) SetName(value string) DeviceDhcpv6Client {

	obj.obj.Name = &value
	return obj
}

// If Rapid Commit is set, client initiates Rapid Commit two-message exchange by including Rapid Commit option  in Solicit message.
// RapidCommit returns a bool
func (obj *deviceDhcpv6Client) RapidCommit() bool {

	return *obj.obj.RapidCommit

}

// If Rapid Commit is set, client initiates Rapid Commit two-message exchange by including Rapid Commit option  in Solicit message.
// RapidCommit returns a bool
func (obj *deviceDhcpv6Client) HasRapidCommit() bool {
	return obj.obj.RapidCommit != nil
}

// If Rapid Commit is set, client initiates Rapid Commit two-message exchange by including Rapid Commit option  in Solicit message.
// SetRapidCommit sets the bool value in the DeviceDhcpv6Client object
func (obj *deviceDhcpv6Client) SetRapidCommit(value bool) DeviceDhcpv6Client {

	obj.obj.RapidCommit = &value
	return obj
}

// A client uses multicast to reach all servers or an individual server. An individual server is indicated by
// specifying that server's DUID in a Server Identifier option in the client's message (all servers will receive
// this message but only the indicated server will respond). All servers are indicated by not supplying this option.
// ServerIdentifier returns a Dhcpv6ClientOptionsServerIdentifier
func (obj *deviceDhcpv6Client) ServerIdentifier() Dhcpv6ClientOptionsServerIdentifier {
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
func (obj *deviceDhcpv6Client) HasServerIdentifier() bool {
	return obj.obj.ServerIdentifier != nil
}

// A client uses multicast to reach all servers or an individual server. An individual server is indicated by
// specifying that server's DUID in a Server Identifier option in the client's message (all servers will receive
// this message but only the indicated server will respond). All servers are indicated by not supplying this option.
// SetServerIdentifier sets the Dhcpv6ClientOptionsServerIdentifier value in the DeviceDhcpv6Client object
func (obj *deviceDhcpv6Client) SetServerIdentifier(value Dhcpv6ClientOptionsServerIdentifier) DeviceDhcpv6Client {

	obj.serverIdentifierHolder = nil
	obj.obj.ServerIdentifier = value.msg()

	return obj
}

// Each IA has an associated IAID. Differnet IA options represent different types of IPv6 addresses and parameters
// accepted by DHCPv6 clients each used in different context by an IPv6 node.
// IaType returns a DeviceDhcpv6ClientIaType
func (obj *deviceDhcpv6Client) IaType() DeviceDhcpv6ClientIaType {
	if obj.obj.IaType == nil {
		obj.obj.IaType = NewDeviceDhcpv6ClientIaType().msg()
	}
	if obj.iaTypeHolder == nil {
		obj.iaTypeHolder = &deviceDhcpv6ClientIaType{obj: obj.obj.IaType}
	}
	return obj.iaTypeHolder
}

// Each IA has an associated IAID. Differnet IA options represent different types of IPv6 addresses and parameters
// accepted by DHCPv6 clients each used in different context by an IPv6 node.
// IaType returns a DeviceDhcpv6ClientIaType
func (obj *deviceDhcpv6Client) HasIaType() bool {
	return obj.obj.IaType != nil
}

// Each IA has an associated IAID. Differnet IA options represent different types of IPv6 addresses and parameters
// accepted by DHCPv6 clients each used in different context by an IPv6 node.
// SetIaType sets the DeviceDhcpv6ClientIaType value in the DeviceDhcpv6Client object
func (obj *deviceDhcpv6Client) SetIaType(value DeviceDhcpv6ClientIaType) DeviceDhcpv6Client {

	obj.iaTypeHolder = nil
	obj.obj.IaType = value.msg()

	return obj
}

// Each DHCP client and server has a DUID. DHCP clients and servers use DUIDs to identify each other.
// DuidType returns a DeviceDhcpv6ClientDuidType
func (obj *deviceDhcpv6Client) DuidType() DeviceDhcpv6ClientDuidType {
	if obj.obj.DuidType == nil {
		obj.obj.DuidType = NewDeviceDhcpv6ClientDuidType().msg()
	}
	if obj.duidTypeHolder == nil {
		obj.duidTypeHolder = &deviceDhcpv6ClientDuidType{obj: obj.obj.DuidType}
	}
	return obj.duidTypeHolder
}

// Each DHCP client and server has a DUID. DHCP clients and servers use DUIDs to identify each other.
// DuidType returns a DeviceDhcpv6ClientDuidType
func (obj *deviceDhcpv6Client) HasDuidType() bool {
	return obj.obj.DuidType != nil
}

// Each DHCP client and server has a DUID. DHCP clients and servers use DUIDs to identify each other.
// SetDuidType sets the DeviceDhcpv6ClientDuidType value in the DeviceDhcpv6Client object
func (obj *deviceDhcpv6Client) SetDuidType(value DeviceDhcpv6ClientDuidType) DeviceDhcpv6Client {

	obj.duidTypeHolder = nil
	obj.obj.DuidType = value.msg()

	return obj
}

// List of options in a message between a client and a server.
// Options returns a []Dhcpv6ClientOptionsOptionsRequest
func (obj *deviceDhcpv6Client) Options() DeviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter {
	if len(obj.obj.Options) == 0 {
		obj.obj.Options = []*otg.Dhcpv6ClientOptionsOptionsRequest{}
	}
	if obj.optionsHolder == nil {
		obj.optionsHolder = newDeviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter(&obj.obj.Options).setMsg(obj)
	}
	return obj.optionsHolder
}

type deviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter struct {
	obj                                    *deviceDhcpv6Client
	dhcpv6ClientOptionsOptionsRequestSlice []Dhcpv6ClientOptionsOptionsRequest
	fieldPtr                               *[]*otg.Dhcpv6ClientOptionsOptionsRequest
}

func newDeviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter(ptr *[]*otg.Dhcpv6ClientOptionsOptionsRequest) DeviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter {
	return &deviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter{fieldPtr: ptr}
}

type DeviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter interface {
	setMsg(*deviceDhcpv6Client) DeviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter
	Items() []Dhcpv6ClientOptionsOptionsRequest
	Add() Dhcpv6ClientOptionsOptionsRequest
	Append(items ...Dhcpv6ClientOptionsOptionsRequest) DeviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter
	Set(index int, newObj Dhcpv6ClientOptionsOptionsRequest) DeviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter
	Clear() DeviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter
	clearHolderSlice() DeviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter
	appendHolderSlice(item Dhcpv6ClientOptionsOptionsRequest) DeviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter
}

func (obj *deviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter) setMsg(msg *deviceDhcpv6Client) DeviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&dhcpv6ClientOptionsOptionsRequest{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter) Items() []Dhcpv6ClientOptionsOptionsRequest {
	return obj.dhcpv6ClientOptionsOptionsRequestSlice
}

func (obj *deviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter) Add() Dhcpv6ClientOptionsOptionsRequest {
	newObj := &otg.Dhcpv6ClientOptionsOptionsRequest{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &dhcpv6ClientOptionsOptionsRequest{obj: newObj}
	newLibObj.setDefault()
	obj.dhcpv6ClientOptionsOptionsRequestSlice = append(obj.dhcpv6ClientOptionsOptionsRequestSlice, newLibObj)
	return newLibObj
}

func (obj *deviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter) Append(items ...Dhcpv6ClientOptionsOptionsRequest) DeviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.dhcpv6ClientOptionsOptionsRequestSlice = append(obj.dhcpv6ClientOptionsOptionsRequestSlice, item)
	}
	return obj
}

func (obj *deviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter) Set(index int, newObj Dhcpv6ClientOptionsOptionsRequest) DeviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.dhcpv6ClientOptionsOptionsRequestSlice[index] = newObj
	return obj
}
func (obj *deviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter) Clear() DeviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Dhcpv6ClientOptionsOptionsRequest{}
		obj.dhcpv6ClientOptionsOptionsRequestSlice = []Dhcpv6ClientOptionsOptionsRequest{}
	}
	return obj
}
func (obj *deviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter) clearHolderSlice() DeviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter {
	if len(obj.dhcpv6ClientOptionsOptionsRequestSlice) > 0 {
		obj.dhcpv6ClientOptionsOptionsRequestSlice = []Dhcpv6ClientOptionsOptionsRequest{}
	}
	return obj
}
func (obj *deviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter) appendHolderSlice(item Dhcpv6ClientOptionsOptionsRequest) DeviceDhcpv6ClientDhcpv6ClientOptionsOptionsRequestIter {
	obj.dhcpv6ClientOptionsOptionsRequestSlice = append(obj.dhcpv6ClientOptionsOptionsRequestSlice, item)
	return obj
}

// The vendor class option is used by a client to identify the vendor that manufactured the hardware on which
// the client is running. The information contained in the data area of this option is contained in one or more
// opaque fields that identify details of the hardware configuration.
// VendorClass returns a Dhcpv6ClientOptionsVendorClass
func (obj *deviceDhcpv6Client) VendorClass() Dhcpv6ClientOptionsVendorClass {
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
func (obj *deviceDhcpv6Client) HasVendorClass() bool {
	return obj.obj.VendorClass != nil
}

// The vendor class option is used by a client to identify the vendor that manufactured the hardware on which
// the client is running. The information contained in the data area of this option is contained in one or more
// opaque fields that identify details of the hardware configuration.
// SetVendorClass sets the Dhcpv6ClientOptionsVendorClass value in the DeviceDhcpv6Client object
func (obj *deviceDhcpv6Client) SetVendorClass(value Dhcpv6ClientOptionsVendorClass) DeviceDhcpv6Client {

	obj.vendorClassHolder = nil
	obj.obj.VendorClass = value.msg()

	return obj
}

// This option is used by clients to exchange vendor-specific information with servers.
// VendorInfo returns a Dhcpv6ClientOptionsVendorInfo
func (obj *deviceDhcpv6Client) VendorInfo() Dhcpv6ClientOptionsVendorInfo {
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
func (obj *deviceDhcpv6Client) HasVendorInfo() bool {
	return obj.obj.VendorInfo != nil
}

// This option is used by clients to exchange vendor-specific information with servers.
// SetVendorInfo sets the Dhcpv6ClientOptionsVendorInfo value in the DeviceDhcpv6Client object
func (obj *deviceDhcpv6Client) SetVendorInfo(value Dhcpv6ClientOptionsVendorInfo) DeviceDhcpv6Client {

	obj.vendorInfoHolder = nil
	obj.obj.VendorInfo = value.msg()

	return obj
}

// DHCPv6 server needs to know the FQDN of the client for the addresses for the client's IA_NA bindings in order to
// update the IPv6-address-to-FQDN mapping. This option allows the client to convey its FQDN to the server. The Client
// FQDN option also contains Flags that DHCPv6 clients and servers use to negotiate who does which update.
// Fqdn returns a Dhcpv6ClientOptionsFqdn
func (obj *deviceDhcpv6Client) Fqdn() Dhcpv6ClientOptionsFqdn {
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
func (obj *deviceDhcpv6Client) HasFqdn() bool {
	return obj.obj.Fqdn != nil
}

// DHCPv6 server needs to know the FQDN of the client for the addresses for the client's IA_NA bindings in order to
// update the IPv6-address-to-FQDN mapping. This option allows the client to convey its FQDN to the server. The Client
// FQDN option also contains Flags that DHCPv6 clients and servers use to negotiate who does which update.
// SetFqdn sets the Dhcpv6ClientOptionsFqdn value in the DeviceDhcpv6Client object
func (obj *deviceDhcpv6Client) SetFqdn(value Dhcpv6ClientOptionsFqdn) DeviceDhcpv6Client {

	obj.fqdnHolder = nil
	obj.obj.Fqdn = value.msg()

	return obj
}

func (obj *deviceDhcpv6Client) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface DeviceDhcpv6Client")
	}

	if obj.obj.ServerIdentifier != nil {

		obj.ServerIdentifier().validateObj(vObj, set_default)
	}

	if obj.obj.IaType != nil {

		obj.IaType().validateObj(vObj, set_default)
	}

	if obj.obj.DuidType != nil {

		obj.DuidType().validateObj(vObj, set_default)
	}

	if len(obj.obj.Options) != 0 {

		if set_default {
			obj.Options().clearHolderSlice()
			for _, item := range obj.obj.Options {
				obj.Options().appendHolderSlice(&dhcpv6ClientOptionsOptionsRequest{obj: item})
			}
		}
		for _, item := range obj.Options().Items() {
			item.validateObj(vObj, set_default)
		}

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

func (obj *deviceDhcpv6Client) setDefault() {
	if obj.obj.RapidCommit == nil {
		obj.SetRapidCommit(false)
	}

}
