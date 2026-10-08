package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LdpIpv4TargetedPeer *****
type ldpIpv4TargetedPeer struct {
	validation
	obj                  *otg.LdpIpv4TargetedPeer
	marshaller           marshalLdpIpv4TargetedPeer
	unMarshaller         unMarshalLdpIpv4TargetedPeer
	authenticationHolder LdpAuthentication
}

func NewLdpIpv4TargetedPeer() LdpIpv4TargetedPeer {
	obj := ldpIpv4TargetedPeer{obj: &otg.LdpIpv4TargetedPeer{}}
	obj.setDefault()
	return &obj
}

func (obj *ldpIpv4TargetedPeer) msg() *otg.LdpIpv4TargetedPeer {
	return obj.obj
}

func (obj *ldpIpv4TargetedPeer) setMsg(msg *otg.LdpIpv4TargetedPeer) LdpIpv4TargetedPeer {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalldpIpv4TargetedPeer struct {
	obj *ldpIpv4TargetedPeer
}

type marshalLdpIpv4TargetedPeer interface {
	// ToProto marshals LdpIpv4TargetedPeer to protobuf object *otg.LdpIpv4TargetedPeer
	ToProto() (*otg.LdpIpv4TargetedPeer, error)
	// ToPbText marshals LdpIpv4TargetedPeer to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LdpIpv4TargetedPeer to YAML text
	ToYaml() (string, error)
	// ToJson marshals LdpIpv4TargetedPeer to JSON text
	ToJson() (string, error)
}

type unMarshalldpIpv4TargetedPeer struct {
	obj *ldpIpv4TargetedPeer
}

type unMarshalLdpIpv4TargetedPeer interface {
	// FromProto unmarshals LdpIpv4TargetedPeer from protobuf object *otg.LdpIpv4TargetedPeer
	FromProto(msg *otg.LdpIpv4TargetedPeer) (LdpIpv4TargetedPeer, error)
	// FromPbText unmarshals LdpIpv4TargetedPeer from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LdpIpv4TargetedPeer from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LdpIpv4TargetedPeer from JSON text
	FromJson(value string) error
}

func (obj *ldpIpv4TargetedPeer) Marshal() marshalLdpIpv4TargetedPeer {
	if obj.marshaller == nil {
		obj.marshaller = &marshalldpIpv4TargetedPeer{obj: obj}
	}
	return obj.marshaller
}

func (obj *ldpIpv4TargetedPeer) Unmarshal() unMarshalLdpIpv4TargetedPeer {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalldpIpv4TargetedPeer{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalldpIpv4TargetedPeer) ToProto() (*otg.LdpIpv4TargetedPeer, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalldpIpv4TargetedPeer) FromProto(msg *otg.LdpIpv4TargetedPeer) (LdpIpv4TargetedPeer, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalldpIpv4TargetedPeer) ToPbText() (string, error) {
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

func (m *unMarshalldpIpv4TargetedPeer) FromPbText(value string) error {
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

func (m *marshalldpIpv4TargetedPeer) ToYaml() (string, error) {
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

func (m *unMarshalldpIpv4TargetedPeer) FromYaml(value string) error {
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

func (m *marshalldpIpv4TargetedPeer) ToJson() (string, error) {
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

func (m *unMarshalldpIpv4TargetedPeer) FromJson(value string) error {
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

func (obj *ldpIpv4TargetedPeer) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ldpIpv4TargetedPeer) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ldpIpv4TargetedPeer) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ldpIpv4TargetedPeer) Clone() (LdpIpv4TargetedPeer, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLdpIpv4TargetedPeer()
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

func (obj *ldpIpv4TargetedPeer) setNil() {
	obj.authenticationHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// LdpIpv4TargetedPeer is lDP Extended Discovery to one IPv4 peer (RFC 5036 Section 2.4.2). The emulated LSR sends unicast Targeted Hellos (UDP port 646) from the address of ipv4_name to remote_address and forms a session with that LSR. remote_address must be reachable from the DUT, for example through an IGP or a static route.
type LdpIpv4TargetedPeer interface {
	Validation
	// msg marshals LdpIpv4TargetedPeer to protobuf object *otg.LdpIpv4TargetedPeer
	// and doesn't set defaults
	msg() *otg.LdpIpv4TargetedPeer
	// setMsg unmarshals LdpIpv4TargetedPeer from protobuf object *otg.LdpIpv4TargetedPeer
	// and doesn't set defaults
	setMsg(*otg.LdpIpv4TargetedPeer) LdpIpv4TargetedPeer
	// provides marshal interface
	Marshal() marshalLdpIpv4TargetedPeer
	// provides unmarshal interface
	Unmarshal() unMarshalLdpIpv4TargetedPeer
	// validate validates LdpIpv4TargetedPeer
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LdpIpv4TargetedPeer, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in LdpIpv4TargetedPeer.
	Name() string
	// SetName assigns string provided by user to LdpIpv4TargetedPeer
	SetName(value string) LdpIpv4TargetedPeer
	// Ipv4Name returns string, set in LdpIpv4TargetedPeer.
	Ipv4Name() string
	// SetIpv4Name assigns string provided by user to LdpIpv4TargetedPeer
	SetIpv4Name(value string) LdpIpv4TargetedPeer
	// RemoteAddress returns string, set in LdpIpv4TargetedPeer.
	RemoteAddress() string
	// SetRemoteAddress assigns string provided by user to LdpIpv4TargetedPeer
	SetRemoteAddress(value string) LdpIpv4TargetedPeer
	// HelloInterval returns uint32, set in LdpIpv4TargetedPeer.
	HelloInterval() uint32
	// SetHelloInterval assigns uint32 provided by user to LdpIpv4TargetedPeer
	SetHelloInterval(value uint32) LdpIpv4TargetedPeer
	// HasHelloInterval checks if HelloInterval has been set in LdpIpv4TargetedPeer
	HasHelloInterval() bool
	// HelloHoldTime returns uint32, set in LdpIpv4TargetedPeer.
	HelloHoldTime() uint32
	// SetHelloHoldTime assigns uint32 provided by user to LdpIpv4TargetedPeer
	SetHelloHoldTime(value uint32) LdpIpv4TargetedPeer
	// HasHelloHoldTime checks if HelloHoldTime has been set in LdpIpv4TargetedPeer
	HasHelloHoldTime() bool
	// Authentication returns LdpAuthentication, set in LdpIpv4TargetedPeer.
	// LdpAuthentication is lDP session authentication. When this object is present, the TCP MD5 Signature Option (RFC 2385) is used on the LDP session TCP connection (RFC 5036 Section 2.9). Omit the object to disable authentication.
	Authentication() LdpAuthentication
	// SetAuthentication assigns LdpAuthentication provided by user to LdpIpv4TargetedPeer.
	// LdpAuthentication is lDP session authentication. When this object is present, the TCP MD5 Signature Option (RFC 2385) is used on the LDP session TCP connection (RFC 5036 Section 2.9). Omit the object to disable authentication.
	SetAuthentication(value LdpAuthentication) LdpIpv4TargetedPeer
	// HasAuthentication checks if Authentication has been set in LdpIpv4TargetedPeer
	HasAuthentication() bool
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *ldpIpv4TargetedPeer) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the LdpIpv4TargetedPeer object
func (obj *ldpIpv4TargetedPeer) SetName(value string) LdpIpv4TargetedPeer {

	obj.obj.Name = &value
	return obj
}

// The globally unique name of the local IPv4 interface or IPv4 loopback from which Targeted Hellos are sent. It must match the "name" field of an entry in "ipv4_addresses" or "ipv4_loopbacks".
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
// - /components/schemas/Device.Ipv4Loopback/properties/name
//
// Ipv4Name returns a string
func (obj *ldpIpv4TargetedPeer) Ipv4Name() string {

	return *obj.obj.Ipv4Name

}

// The globally unique name of the local IPv4 interface or IPv4 loopback from which Targeted Hellos are sent. It must match the "name" field of an entry in "ipv4_addresses" or "ipv4_loopbacks".
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
// - /components/schemas/Device.Ipv4Loopback/properties/name
//
// SetIpv4Name sets the string value in the LdpIpv4TargetedPeer object
func (obj *ldpIpv4TargetedPeer) SetIpv4Name(value string) LdpIpv4TargetedPeer {

	obj.obj.Ipv4Name = &value
	return obj
}

// The IPv4 address of the targeted peer, to which Targeted Hellos are sent (RFC 5036 Section 2.4.2).
// RemoteAddress returns a string
func (obj *ldpIpv4TargetedPeer) RemoteAddress() string {

	return *obj.obj.RemoteAddress

}

// The IPv4 address of the targeted peer, to which Targeted Hellos are sent (RFC 5036 Section 2.4.2).
// SetRemoteAddress sets the string value in the LdpIpv4TargetedPeer object
func (obj *ldpIpv4TargetedPeer) SetRemoteAddress(value string) LdpIpv4TargetedPeer {

	obj.obj.RemoteAddress = &value
	return obj
}

// The interval in seconds between Targeted Hello messages. It should be at most one third of hello_hold_time (RFC 5036 Section 3.5.2.1).
// HelloInterval returns a uint32
func (obj *ldpIpv4TargetedPeer) HelloInterval() uint32 {

	return *obj.obj.HelloInterval

}

// The interval in seconds between Targeted Hello messages. It should be at most one third of hello_hold_time (RFC 5036 Section 3.5.2.1).
// HelloInterval returns a uint32
func (obj *ldpIpv4TargetedPeer) HasHelloInterval() bool {
	return obj.obj.HelloInterval != nil
}

// The interval in seconds between Targeted Hello messages. It should be at most one third of hello_hold_time (RFC 5036 Section 3.5.2.1).
// SetHelloInterval sets the uint32 value in the LdpIpv4TargetedPeer object
func (obj *ldpIpv4TargetedPeer) SetHelloInterval(value uint32) LdpIpv4TargetedPeer {

	obj.obj.HelloInterval = &value
	return obj
}

// The Hold Time in seconds advertised in the Common Hello Parameters TLV of Targeted Hellos (RFC 5036 Section 3.5.2). The value 65535 means infinite.
// HelloHoldTime returns a uint32
func (obj *ldpIpv4TargetedPeer) HelloHoldTime() uint32 {

	return *obj.obj.HelloHoldTime

}

// The Hold Time in seconds advertised in the Common Hello Parameters TLV of Targeted Hellos (RFC 5036 Section 3.5.2). The value 65535 means infinite.
// HelloHoldTime returns a uint32
func (obj *ldpIpv4TargetedPeer) HasHelloHoldTime() bool {
	return obj.obj.HelloHoldTime != nil
}

// The Hold Time in seconds advertised in the Common Hello Parameters TLV of Targeted Hellos (RFC 5036 Section 3.5.2). The value 65535 means infinite.
// SetHelloHoldTime sets the uint32 value in the LdpIpv4TargetedPeer object
func (obj *ldpIpv4TargetedPeer) SetHelloHoldTime(value uint32) LdpIpv4TargetedPeer {

	obj.obj.HelloHoldTime = &value
	return obj
}

// When present, the TCP MD5 Signature Option is used on the session TCP connection to this peer (RFC 5036 Section 2.9). Omit to disable authentication. If the peer is also reached through an Ldp.Ipv4Interface, both must use the same key.
// Authentication returns a LdpAuthentication
func (obj *ldpIpv4TargetedPeer) Authentication() LdpAuthentication {
	if obj.obj.Authentication == nil {
		obj.obj.Authentication = NewLdpAuthentication().msg()
	}
	if obj.authenticationHolder == nil {
		obj.authenticationHolder = &ldpAuthentication{obj: obj.obj.Authentication}
	}
	return obj.authenticationHolder
}

// When present, the TCP MD5 Signature Option is used on the session TCP connection to this peer (RFC 5036 Section 2.9). Omit to disable authentication. If the peer is also reached through an Ldp.Ipv4Interface, both must use the same key.
// Authentication returns a LdpAuthentication
func (obj *ldpIpv4TargetedPeer) HasAuthentication() bool {
	return obj.obj.Authentication != nil
}

// When present, the TCP MD5 Signature Option is used on the session TCP connection to this peer (RFC 5036 Section 2.9). Omit to disable authentication. If the peer is also reached through an Ldp.Ipv4Interface, both must use the same key.
// SetAuthentication sets the LdpAuthentication value in the LdpIpv4TargetedPeer object
func (obj *ldpIpv4TargetedPeer) SetAuthentication(value LdpAuthentication) LdpIpv4TargetedPeer {

	obj.authenticationHolder = nil
	obj.obj.Authentication = value.msg()

	return obj
}

func (obj *ldpIpv4TargetedPeer) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface LdpIpv4TargetedPeer")
	}

	// Ipv4Name is required
	if obj.obj.Ipv4Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Ipv4Name is required field on interface LdpIpv4TargetedPeer")
	}

	// RemoteAddress is required
	if obj.obj.RemoteAddress == nil {
		vObj.validationErrors = append(vObj.validationErrors, "RemoteAddress is required field on interface LdpIpv4TargetedPeer")
	}
	if obj.obj.RemoteAddress != nil {

		err := obj.validateIpv4(obj.RemoteAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on LdpIpv4TargetedPeer.RemoteAddress"))
		}

	}

	if obj.obj.HelloInterval != nil {

		if *obj.obj.HelloInterval < 1 || *obj.obj.HelloInterval > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= LdpIpv4TargetedPeer.HelloInterval <= 65535 but Got %d", *obj.obj.HelloInterval))
		}

	}

	if obj.obj.HelloHoldTime != nil {

		if *obj.obj.HelloHoldTime < 1 || *obj.obj.HelloHoldTime > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= LdpIpv4TargetedPeer.HelloHoldTime <= 65535 but Got %d", *obj.obj.HelloHoldTime))
		}

	}

	if obj.obj.Authentication != nil {

		obj.Authentication().validateObj(vObj, set_default)
	}

}

func (obj *ldpIpv4TargetedPeer) setDefault() {
	if obj.obj.HelloInterval == nil {
		obj.SetHelloInterval(15)
	}
	if obj.obj.HelloHoldTime == nil {
		obj.SetHelloHoldTime(45)
	}

}
