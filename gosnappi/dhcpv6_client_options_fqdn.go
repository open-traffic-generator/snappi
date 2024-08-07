package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ClientOptionsFqdn *****
type dhcpv6ClientOptionsFqdn struct {
	validation
	obj                          *otg.Dhcpv6ClientOptionsFqdn
	marshaller                   marshalDhcpv6ClientOptionsFqdn
	unMarshaller                 unMarshalDhcpv6ClientOptionsFqdn
	associatedDhcpMessagesHolder Dhcpv6ClientOptionsIncludedMessages
}

func NewDhcpv6ClientOptionsFqdn() Dhcpv6ClientOptionsFqdn {
	obj := dhcpv6ClientOptionsFqdn{obj: &otg.Dhcpv6ClientOptionsFqdn{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ClientOptionsFqdn) msg() *otg.Dhcpv6ClientOptionsFqdn {
	return obj.obj
}

func (obj *dhcpv6ClientOptionsFqdn) setMsg(msg *otg.Dhcpv6ClientOptionsFqdn) Dhcpv6ClientOptionsFqdn {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ClientOptionsFqdn struct {
	obj *dhcpv6ClientOptionsFqdn
}

type marshalDhcpv6ClientOptionsFqdn interface {
	// ToProto marshals Dhcpv6ClientOptionsFqdn to protobuf object *otg.Dhcpv6ClientOptionsFqdn
	ToProto() (*otg.Dhcpv6ClientOptionsFqdn, error)
	// ToPbText marshals Dhcpv6ClientOptionsFqdn to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ClientOptionsFqdn to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ClientOptionsFqdn to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpv6ClientOptionsFqdn struct {
	obj *dhcpv6ClientOptionsFqdn
}

type unMarshalDhcpv6ClientOptionsFqdn interface {
	// FromProto unmarshals Dhcpv6ClientOptionsFqdn from protobuf object *otg.Dhcpv6ClientOptionsFqdn
	FromProto(msg *otg.Dhcpv6ClientOptionsFqdn) (Dhcpv6ClientOptionsFqdn, error)
	// FromPbText unmarshals Dhcpv6ClientOptionsFqdn from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ClientOptionsFqdn from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ClientOptionsFqdn from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ClientOptionsFqdn) Marshal() marshalDhcpv6ClientOptionsFqdn {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ClientOptionsFqdn{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ClientOptionsFqdn) Unmarshal() unMarshalDhcpv6ClientOptionsFqdn {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ClientOptionsFqdn{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ClientOptionsFqdn) ToProto() (*otg.Dhcpv6ClientOptionsFqdn, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ClientOptionsFqdn) FromProto(msg *otg.Dhcpv6ClientOptionsFqdn) (Dhcpv6ClientOptionsFqdn, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ClientOptionsFqdn) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsFqdn) FromPbText(value string) error {
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

func (m *marshaldhcpv6ClientOptionsFqdn) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsFqdn) FromYaml(value string) error {
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

func (m *marshaldhcpv6ClientOptionsFqdn) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsFqdn) FromJson(value string) error {
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

func (obj *dhcpv6ClientOptionsFqdn) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsFqdn) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsFqdn) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ClientOptionsFqdn) Clone() (Dhcpv6ClientOptionsFqdn, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ClientOptionsFqdn()
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

func (obj *dhcpv6ClientOptionsFqdn) setNil() {
	obj.associatedDhcpMessagesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Dhcpv6ClientOptionsFqdn is dHCPv6 server needs to know the FQDN of the client for the addresses for the client's IA_NA bindings in order to update the IPv6-address-to-FQDN mapping. This option allows the client to convey its FQDN to the server. The Client  FQDN option also contains Flags that DHCPv6 clients and servers use to negotiate who does which updates. The option code is 39.
type Dhcpv6ClientOptionsFqdn interface {
	Validation
	// msg marshals Dhcpv6ClientOptionsFqdn to protobuf object *otg.Dhcpv6ClientOptionsFqdn
	// and doesn't set defaults
	msg() *otg.Dhcpv6ClientOptionsFqdn
	// setMsg unmarshals Dhcpv6ClientOptionsFqdn from protobuf object *otg.Dhcpv6ClientOptionsFqdn
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ClientOptionsFqdn) Dhcpv6ClientOptionsFqdn
	// provides marshal interface
	Marshal() marshalDhcpv6ClientOptionsFqdn
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ClientOptionsFqdn
	// validate validates Dhcpv6ClientOptionsFqdn
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ClientOptionsFqdn, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// FlagS returns bool, set in Dhcpv6ClientOptionsFqdn.
	FlagS() bool
	// SetFlagS assigns bool provided by user to Dhcpv6ClientOptionsFqdn
	SetFlagS(value bool) Dhcpv6ClientOptionsFqdn
	// HasFlagS checks if FlagS has been set in Dhcpv6ClientOptionsFqdn
	HasFlagS() bool
	// FlagO returns bool, set in Dhcpv6ClientOptionsFqdn.
	FlagO() bool
	// SetFlagO assigns bool provided by user to Dhcpv6ClientOptionsFqdn
	SetFlagO(value bool) Dhcpv6ClientOptionsFqdn
	// HasFlagO checks if FlagO has been set in Dhcpv6ClientOptionsFqdn
	HasFlagO() bool
	// FlagN returns bool, set in Dhcpv6ClientOptionsFqdn.
	FlagN() bool
	// SetFlagN assigns bool provided by user to Dhcpv6ClientOptionsFqdn
	SetFlagN(value bool) Dhcpv6ClientOptionsFqdn
	// HasFlagN checks if FlagN has been set in Dhcpv6ClientOptionsFqdn
	HasFlagN() bool
	// DomainName returns string, set in Dhcpv6ClientOptionsFqdn.
	DomainName() string
	// SetDomainName assigns string provided by user to Dhcpv6ClientOptionsFqdn
	SetDomainName(value string) Dhcpv6ClientOptionsFqdn
	// AssociatedDhcpMessages returns Dhcpv6ClientOptionsIncludedMessages, set in Dhcpv6ClientOptionsFqdn.
	// Dhcpv6ClientOptionsIncludedMessages is the dhcpv6 client messages where the option will be included. If all is selected the selected option will be added in the all the Dhcpv6 client messages, else based on the selection in particular Dhcpv6 client messages the option will be included.
	AssociatedDhcpMessages() Dhcpv6ClientOptionsIncludedMessages
	// SetAssociatedDhcpMessages assigns Dhcpv6ClientOptionsIncludedMessages provided by user to Dhcpv6ClientOptionsFqdn.
	// Dhcpv6ClientOptionsIncludedMessages is the dhcpv6 client messages where the option will be included. If all is selected the selected option will be added in the all the Dhcpv6 client messages, else based on the selection in particular Dhcpv6 client messages the option will be included.
	SetAssociatedDhcpMessages(value Dhcpv6ClientOptionsIncludedMessages) Dhcpv6ClientOptionsFqdn
	// HasAssociatedDhcpMessages checks if AssociatedDhcpMessages has been set in Dhcpv6ClientOptionsFqdn
	HasAssociatedDhcpMessages() bool
	setNil()
}

// The "S" bit indicates whether the server should or should not perform the AAAA RR (FQDN-to-address) DNS updates.  A client sets the bit to 0 to indicate that the server should not perform the updates and 1 to indicate that the  server should perform the updates. The state of the bit in the reply from the server indicates the action to be  taken by the server. If it is 1, the server has taken responsibility for AAAA RR updates for the FQDN.
// FlagS returns a bool
func (obj *dhcpv6ClientOptionsFqdn) FlagS() bool {

	return *obj.obj.FlagS

}

// The "S" bit indicates whether the server should or should not perform the AAAA RR (FQDN-to-address) DNS updates.  A client sets the bit to 0 to indicate that the server should not perform the updates and 1 to indicate that the  server should perform the updates. The state of the bit in the reply from the server indicates the action to be  taken by the server. If it is 1, the server has taken responsibility for AAAA RR updates for the FQDN.
// FlagS returns a bool
func (obj *dhcpv6ClientOptionsFqdn) HasFlagS() bool {
	return obj.obj.FlagS != nil
}

// The "S" bit indicates whether the server should or should not perform the AAAA RR (FQDN-to-address) DNS updates.  A client sets the bit to 0 to indicate that the server should not perform the updates and 1 to indicate that the  server should perform the updates. The state of the bit in the reply from the server indicates the action to be  taken by the server. If it is 1, the server has taken responsibility for AAAA RR updates for the FQDN.
// SetFlagS sets the bool value in the Dhcpv6ClientOptionsFqdn object
func (obj *dhcpv6ClientOptionsFqdn) SetFlagS(value bool) Dhcpv6ClientOptionsFqdn {

	obj.obj.FlagS = &value
	return obj
}

// The "O" bit indicates whether the server has overridden the client's preference for the "S" bit. A client must set this  bit to 0. A server must set this bit to 1 if the "S" bit in its reply to the client does not match the "S" bit received  from the client.
// FlagO returns a bool
func (obj *dhcpv6ClientOptionsFqdn) FlagO() bool {

	return *obj.obj.FlagO

}

// The "O" bit indicates whether the server has overridden the client's preference for the "S" bit. A client must set this  bit to 0. A server must set this bit to 1 if the "S" bit in its reply to the client does not match the "S" bit received  from the client.
// FlagO returns a bool
func (obj *dhcpv6ClientOptionsFqdn) HasFlagO() bool {
	return obj.obj.FlagO != nil
}

// The "O" bit indicates whether the server has overridden the client's preference for the "S" bit. A client must set this  bit to 0. A server must set this bit to 1 if the "S" bit in its reply to the client does not match the "S" bit received  from the client.
// SetFlagO sets the bool value in the Dhcpv6ClientOptionsFqdn object
func (obj *dhcpv6ClientOptionsFqdn) SetFlagO(value bool) Dhcpv6ClientOptionsFqdn {

	obj.obj.FlagO = &value
	return obj
}

// The "N" bit indicates whether the server should not perform any DNS updates. A client sets this bit to 0 to request that  the server should perform updates (the PTR RR and possibly the AAAA RR based on the "S" bit) or to 1 to request that the  server should not perform any DNS updates. A server sets the "N" bit to indicate whether the server shall (0) or shall not (1)  perform DNS updates. If the "N" bit is 1, the "S" bit MUST be 0.
// FlagN returns a bool
func (obj *dhcpv6ClientOptionsFqdn) FlagN() bool {

	return *obj.obj.FlagN

}

// The "N" bit indicates whether the server should not perform any DNS updates. A client sets this bit to 0 to request that  the server should perform updates (the PTR RR and possibly the AAAA RR based on the "S" bit) or to 1 to request that the  server should not perform any DNS updates. A server sets the "N" bit to indicate whether the server shall (0) or shall not (1)  perform DNS updates. If the "N" bit is 1, the "S" bit MUST be 0.
// FlagN returns a bool
func (obj *dhcpv6ClientOptionsFqdn) HasFlagN() bool {
	return obj.obj.FlagN != nil
}

// The "N" bit indicates whether the server should not perform any DNS updates. A client sets this bit to 0 to request that  the server should perform updates (the PTR RR and possibly the AAAA RR based on the "S" bit) or to 1 to request that the  server should not perform any DNS updates. A server sets the "N" bit to indicate whether the server shall (0) or shall not (1)  perform DNS updates. If the "N" bit is 1, the "S" bit MUST be 0.
// SetFlagN sets the bool value in the Dhcpv6ClientOptionsFqdn object
func (obj *dhcpv6ClientOptionsFqdn) SetFlagN(value bool) Dhcpv6ClientOptionsFqdn {

	obj.obj.FlagN = &value
	return obj
}

// The Domain Name part of the option carries all or part of the FQDN of a DHCPv6 client. A client MAY also leave the  Domain Name field empty if it desires the server to provide a name. A fully qualified domain name (FQDN) is the complete  address of an internet host or computer. It provides its exact location within the domain name system (DNS) by specifying  the hostname, domain name and top-level domain (TLD). An FQDN isn't the same as a URL but rather is a part of it that fully  identifies the server to which the request is addressed. An FQDN doesn't carry the TCP/IP protocol information, such as Hypertext  Transfer Protocol (HTTP) or Hypertext Transfer Protocol Secure (HTTPS), which is always used at the beginning of a URL. Therefore,  adding the prefix http:// or https:// to the FQDN turns it into a full URL. One example can be microsoft.com.
// DomainName returns a string
func (obj *dhcpv6ClientOptionsFqdn) DomainName() string {

	return *obj.obj.DomainName

}

// The Domain Name part of the option carries all or part of the FQDN of a DHCPv6 client. A client MAY also leave the  Domain Name field empty if it desires the server to provide a name. A fully qualified domain name (FQDN) is the complete  address of an internet host or computer. It provides its exact location within the domain name system (DNS) by specifying  the hostname, domain name and top-level domain (TLD). An FQDN isn't the same as a URL but rather is a part of it that fully  identifies the server to which the request is addressed. An FQDN doesn't carry the TCP/IP protocol information, such as Hypertext  Transfer Protocol (HTTP) or Hypertext Transfer Protocol Secure (HTTPS), which is always used at the beginning of a URL. Therefore,  adding the prefix http:// or https:// to the FQDN turns it into a full URL. One example can be microsoft.com.
// SetDomainName sets the string value in the Dhcpv6ClientOptionsFqdn object
func (obj *dhcpv6ClientOptionsFqdn) SetDomainName(value string) Dhcpv6ClientOptionsFqdn {

	obj.obj.DomainName = &value
	return obj
}

// The list of dhcpv6 client messages where this option is included.
// AssociatedDhcpMessages returns a Dhcpv6ClientOptionsIncludedMessages
func (obj *dhcpv6ClientOptionsFqdn) AssociatedDhcpMessages() Dhcpv6ClientOptionsIncludedMessages {
	if obj.obj.AssociatedDhcpMessages == nil {
		obj.obj.AssociatedDhcpMessages = NewDhcpv6ClientOptionsIncludedMessages().msg()
	}
	if obj.associatedDhcpMessagesHolder == nil {
		obj.associatedDhcpMessagesHolder = &dhcpv6ClientOptionsIncludedMessages{obj: obj.obj.AssociatedDhcpMessages}
	}
	return obj.associatedDhcpMessagesHolder
}

// The list of dhcpv6 client messages where this option is included.
// AssociatedDhcpMessages returns a Dhcpv6ClientOptionsIncludedMessages
func (obj *dhcpv6ClientOptionsFqdn) HasAssociatedDhcpMessages() bool {
	return obj.obj.AssociatedDhcpMessages != nil
}

// The list of dhcpv6 client messages where this option is included.
// SetAssociatedDhcpMessages sets the Dhcpv6ClientOptionsIncludedMessages value in the Dhcpv6ClientOptionsFqdn object
func (obj *dhcpv6ClientOptionsFqdn) SetAssociatedDhcpMessages(value Dhcpv6ClientOptionsIncludedMessages) Dhcpv6ClientOptionsFqdn {

	obj.associatedDhcpMessagesHolder = nil
	obj.obj.AssociatedDhcpMessages = value.msg()

	return obj
}

func (obj *dhcpv6ClientOptionsFqdn) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// DomainName is required
	if obj.obj.DomainName == nil {
		vObj.validationErrors = append(vObj.validationErrors, "DomainName is required field on interface Dhcpv6ClientOptionsFqdn")
	}

	if obj.obj.AssociatedDhcpMessages != nil {

		obj.AssociatedDhcpMessages().validateObj(vObj, set_default)
	}

}

func (obj *dhcpv6ClientOptionsFqdn) setDefault() {
	if obj.obj.FlagS == nil {
		obj.SetFlagS(true)
	}
	if obj.obj.FlagO == nil {
		obj.SetFlagO(false)
	}
	if obj.obj.FlagN == nil {
		obj.SetFlagN(false)
	}

}
