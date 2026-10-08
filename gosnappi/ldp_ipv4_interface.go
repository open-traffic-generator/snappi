package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LdpIpv4Interface *****
type ldpIpv4Interface struct {
	validation
	obj                  *otg.LdpIpv4Interface
	marshaller           marshalLdpIpv4Interface
	unMarshaller         unMarshalLdpIpv4Interface
	authenticationHolder LdpAuthentication
}

func NewLdpIpv4Interface() LdpIpv4Interface {
	obj := ldpIpv4Interface{obj: &otg.LdpIpv4Interface{}}
	obj.setDefault()
	return &obj
}

func (obj *ldpIpv4Interface) msg() *otg.LdpIpv4Interface {
	return obj.obj
}

func (obj *ldpIpv4Interface) setMsg(msg *otg.LdpIpv4Interface) LdpIpv4Interface {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalldpIpv4Interface struct {
	obj *ldpIpv4Interface
}

type marshalLdpIpv4Interface interface {
	// ToProto marshals LdpIpv4Interface to protobuf object *otg.LdpIpv4Interface
	ToProto() (*otg.LdpIpv4Interface, error)
	// ToPbText marshals LdpIpv4Interface to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LdpIpv4Interface to YAML text
	ToYaml() (string, error)
	// ToJson marshals LdpIpv4Interface to JSON text
	ToJson() (string, error)
}

type unMarshalldpIpv4Interface struct {
	obj *ldpIpv4Interface
}

type unMarshalLdpIpv4Interface interface {
	// FromProto unmarshals LdpIpv4Interface from protobuf object *otg.LdpIpv4Interface
	FromProto(msg *otg.LdpIpv4Interface) (LdpIpv4Interface, error)
	// FromPbText unmarshals LdpIpv4Interface from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LdpIpv4Interface from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LdpIpv4Interface from JSON text
	FromJson(value string) error
}

func (obj *ldpIpv4Interface) Marshal() marshalLdpIpv4Interface {
	if obj.marshaller == nil {
		obj.marshaller = &marshalldpIpv4Interface{obj: obj}
	}
	return obj.marshaller
}

func (obj *ldpIpv4Interface) Unmarshal() unMarshalLdpIpv4Interface {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalldpIpv4Interface{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalldpIpv4Interface) ToProto() (*otg.LdpIpv4Interface, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalldpIpv4Interface) FromProto(msg *otg.LdpIpv4Interface) (LdpIpv4Interface, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalldpIpv4Interface) ToPbText() (string, error) {
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

func (m *unMarshalldpIpv4Interface) FromPbText(value string) error {
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

func (m *marshalldpIpv4Interface) ToYaml() (string, error) {
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

func (m *unMarshalldpIpv4Interface) FromYaml(value string) error {
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

func (m *marshalldpIpv4Interface) ToJson() (string, error) {
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

func (m *unMarshalldpIpv4Interface) FromJson(value string) error {
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

func (obj *ldpIpv4Interface) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ldpIpv4Interface) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ldpIpv4Interface) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ldpIpv4Interface) Clone() (LdpIpv4Interface, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLdpIpv4Interface()
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

func (obj *ldpIpv4Interface) setNil() {
	obj.authenticationHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// LdpIpv4Interface is lDP Basic Discovery on an IPv4 interface (RFC 5036 Section 2.4.1). The emulated LSR sends LDP Link Hellos (UDP port 646, destination 224.0.0.2) on the referenced IPv4 interface and forms a session with each LSR discovered on the link.
type LdpIpv4Interface interface {
	Validation
	// msg marshals LdpIpv4Interface to protobuf object *otg.LdpIpv4Interface
	// and doesn't set defaults
	msg() *otg.LdpIpv4Interface
	// setMsg unmarshals LdpIpv4Interface from protobuf object *otg.LdpIpv4Interface
	// and doesn't set defaults
	setMsg(*otg.LdpIpv4Interface) LdpIpv4Interface
	// provides marshal interface
	Marshal() marshalLdpIpv4Interface
	// provides unmarshal interface
	Unmarshal() unMarshalLdpIpv4Interface
	// validate validates LdpIpv4Interface
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LdpIpv4Interface, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in LdpIpv4Interface.
	Name() string
	// SetName assigns string provided by user to LdpIpv4Interface
	SetName(value string) LdpIpv4Interface
	// Ipv4Name returns string, set in LdpIpv4Interface.
	Ipv4Name() string
	// SetIpv4Name assigns string provided by user to LdpIpv4Interface
	SetIpv4Name(value string) LdpIpv4Interface
	// HelloInterval returns uint32, set in LdpIpv4Interface.
	HelloInterval() uint32
	// SetHelloInterval assigns uint32 provided by user to LdpIpv4Interface
	SetHelloInterval(value uint32) LdpIpv4Interface
	// HasHelloInterval checks if HelloInterval has been set in LdpIpv4Interface
	HasHelloInterval() bool
	// HelloHoldTime returns uint32, set in LdpIpv4Interface.
	HelloHoldTime() uint32
	// SetHelloHoldTime assigns uint32 provided by user to LdpIpv4Interface
	SetHelloHoldTime(value uint32) LdpIpv4Interface
	// HasHelloHoldTime checks if HelloHoldTime has been set in LdpIpv4Interface
	HasHelloHoldTime() bool
	// Authentication returns LdpAuthentication, set in LdpIpv4Interface.
	// LdpAuthentication is lDP session authentication. When this object is present, the TCP MD5 Signature Option (RFC 2385) is used on the LDP session TCP connection (RFC 5036 Section 2.9). Omit the object to disable authentication.
	Authentication() LdpAuthentication
	// SetAuthentication assigns LdpAuthentication provided by user to LdpIpv4Interface.
	// LdpAuthentication is lDP session authentication. When this object is present, the TCP MD5 Signature Option (RFC 2385) is used on the LDP session TCP connection (RFC 5036 Section 2.9). Omit the object to disable authentication.
	SetAuthentication(value LdpAuthentication) LdpIpv4Interface
	// HasAuthentication checks if Authentication has been set in LdpIpv4Interface
	HasAuthentication() bool
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *ldpIpv4Interface) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the LdpIpv4Interface object
func (obj *ldpIpv4Interface) SetName(value string) LdpIpv4Interface {

	obj.obj.Name = &value
	return obj
}

// The globally unique name of the IPv4 interface on which Link Hellos are sent. It must match the "name" field of an entry in "ipv4_addresses". A loopback cannot be used, because Link Hellos are sent on a link.
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
//
// Ipv4Name returns a string
func (obj *ldpIpv4Interface) Ipv4Name() string {

	return *obj.obj.Ipv4Name

}

// The globally unique name of the IPv4 interface on which Link Hellos are sent. It must match the "name" field of an entry in "ipv4_addresses". A loopback cannot be used, because Link Hellos are sent on a link.
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
//
// SetIpv4Name sets the string value in the LdpIpv4Interface object
func (obj *ldpIpv4Interface) SetIpv4Name(value string) LdpIpv4Interface {

	obj.obj.Ipv4Name = &value
	return obj
}

// The interval in seconds between Link Hello messages. RFC 5036 Section 3.5.2.1 asks that Hellos are sent well within the hold time; hello_interval should be at most one third of hello_hold_time.
// HelloInterval returns a uint32
func (obj *ldpIpv4Interface) HelloInterval() uint32 {

	return *obj.obj.HelloInterval

}

// The interval in seconds between Link Hello messages. RFC 5036 Section 3.5.2.1 asks that Hellos are sent well within the hold time; hello_interval should be at most one third of hello_hold_time.
// HelloInterval returns a uint32
func (obj *ldpIpv4Interface) HasHelloInterval() bool {
	return obj.obj.HelloInterval != nil
}

// The interval in seconds between Link Hello messages. RFC 5036 Section 3.5.2.1 asks that Hellos are sent well within the hold time; hello_interval should be at most one third of hello_hold_time.
// SetHelloInterval sets the uint32 value in the LdpIpv4Interface object
func (obj *ldpIpv4Interface) SetHelloInterval(value uint32) LdpIpv4Interface {

	obj.obj.HelloInterval = &value
	return obj
}

// The Hold Time in seconds advertised in the Common Hello Parameters TLV of Link Hellos (RFC 5036 Section 3.5.2). The hello adjacency is removed when no Hello is received within the negotiated hold time, which is the smaller of the two advertised values. The value 65535 means infinite.
// HelloHoldTime returns a uint32
func (obj *ldpIpv4Interface) HelloHoldTime() uint32 {

	return *obj.obj.HelloHoldTime

}

// The Hold Time in seconds advertised in the Common Hello Parameters TLV of Link Hellos (RFC 5036 Section 3.5.2). The hello adjacency is removed when no Hello is received within the negotiated hold time, which is the smaller of the two advertised values. The value 65535 means infinite.
// HelloHoldTime returns a uint32
func (obj *ldpIpv4Interface) HasHelloHoldTime() bool {
	return obj.obj.HelloHoldTime != nil
}

// The Hold Time in seconds advertised in the Common Hello Parameters TLV of Link Hellos (RFC 5036 Section 3.5.2). The hello adjacency is removed when no Hello is received within the negotiated hold time, which is the smaller of the two advertised values. The value 65535 means infinite.
// SetHelloHoldTime sets the uint32 value in the LdpIpv4Interface object
func (obj *ldpIpv4Interface) SetHelloHoldTime(value uint32) LdpIpv4Interface {

	obj.obj.HelloHoldTime = &value
	return obj
}

// When present, the TCP MD5 Signature Option is used on the session TCP connection of every session discovered on this interface (RFC 5036 Section 2.9). Omit to disable authentication. A session is a single TCP connection, so all adjacencies that lead to the same peer LSR must use the same key.
// Authentication returns a LdpAuthentication
func (obj *ldpIpv4Interface) Authentication() LdpAuthentication {
	if obj.obj.Authentication == nil {
		obj.obj.Authentication = NewLdpAuthentication().msg()
	}
	if obj.authenticationHolder == nil {
		obj.authenticationHolder = &ldpAuthentication{obj: obj.obj.Authentication}
	}
	return obj.authenticationHolder
}

// When present, the TCP MD5 Signature Option is used on the session TCP connection of every session discovered on this interface (RFC 5036 Section 2.9). Omit to disable authentication. A session is a single TCP connection, so all adjacencies that lead to the same peer LSR must use the same key.
// Authentication returns a LdpAuthentication
func (obj *ldpIpv4Interface) HasAuthentication() bool {
	return obj.obj.Authentication != nil
}

// When present, the TCP MD5 Signature Option is used on the session TCP connection of every session discovered on this interface (RFC 5036 Section 2.9). Omit to disable authentication. A session is a single TCP connection, so all adjacencies that lead to the same peer LSR must use the same key.
// SetAuthentication sets the LdpAuthentication value in the LdpIpv4Interface object
func (obj *ldpIpv4Interface) SetAuthentication(value LdpAuthentication) LdpIpv4Interface {

	obj.authenticationHolder = nil
	obj.obj.Authentication = value.msg()

	return obj
}

func (obj *ldpIpv4Interface) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface LdpIpv4Interface")
	}

	// Ipv4Name is required
	if obj.obj.Ipv4Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Ipv4Name is required field on interface LdpIpv4Interface")
	}

	if obj.obj.HelloInterval != nil {

		if *obj.obj.HelloInterval < 1 || *obj.obj.HelloInterval > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= LdpIpv4Interface.HelloInterval <= 65535 but Got %d", *obj.obj.HelloInterval))
		}

	}

	if obj.obj.HelloHoldTime != nil {

		if *obj.obj.HelloHoldTime < 1 || *obj.obj.HelloHoldTime > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= LdpIpv4Interface.HelloHoldTime <= 65535 but Got %d", *obj.obj.HelloHoldTime))
		}

	}

	if obj.obj.Authentication != nil {

		obj.Authentication().validateObj(vObj, set_default)
	}

}

func (obj *ldpIpv4Interface) setDefault() {
	if obj.obj.HelloInterval == nil {
		obj.SetHelloInterval(5)
	}
	if obj.obj.HelloHoldTime == nil {
		obj.SetHelloHoldTime(15)
	}

}
