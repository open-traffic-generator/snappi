package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3NssaLsa *****
type ospfv3NssaLsa struct {
	validation
	obj          *otg.Ospfv3NssaLsa
	marshaller   marshalOspfv3NssaLsa
	unMarshaller unMarshalOspfv3NssaLsa
	headerHolder Ospfv3LsaHeader
}

func NewOspfv3NssaLsa() Ospfv3NssaLsa {
	obj := ospfv3NssaLsa{obj: &otg.Ospfv3NssaLsa{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3NssaLsa) msg() *otg.Ospfv3NssaLsa {
	return obj.obj
}

func (obj *ospfv3NssaLsa) setMsg(msg *otg.Ospfv3NssaLsa) Ospfv3NssaLsa {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3NssaLsa struct {
	obj *ospfv3NssaLsa
}

type marshalOspfv3NssaLsa interface {
	// ToProto marshals Ospfv3NssaLsa to protobuf object *otg.Ospfv3NssaLsa
	ToProto() (*otg.Ospfv3NssaLsa, error)
	// ToPbText marshals Ospfv3NssaLsa to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3NssaLsa to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3NssaLsa to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Ospfv3NssaLsa to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalospfv3NssaLsa struct {
	obj *ospfv3NssaLsa
}

type unMarshalOspfv3NssaLsa interface {
	// FromProto unmarshals Ospfv3NssaLsa from protobuf object *otg.Ospfv3NssaLsa
	FromProto(msg *otg.Ospfv3NssaLsa) (Ospfv3NssaLsa, error)
	// FromPbText unmarshals Ospfv3NssaLsa from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3NssaLsa from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3NssaLsa from JSON text
	FromJson(value string) error
}

func (obj *ospfv3NssaLsa) Marshal() marshalOspfv3NssaLsa {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3NssaLsa{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3NssaLsa) Unmarshal() unMarshalOspfv3NssaLsa {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3NssaLsa{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3NssaLsa) ToProto() (*otg.Ospfv3NssaLsa, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3NssaLsa) FromProto(msg *otg.Ospfv3NssaLsa) (Ospfv3NssaLsa, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3NssaLsa) ToPbText() (string, error) {
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

func (m *unMarshalospfv3NssaLsa) FromPbText(value string) error {
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

func (m *marshalospfv3NssaLsa) ToYaml() (string, error) {
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

func (m *unMarshalospfv3NssaLsa) FromYaml(value string) error {
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

func (m *marshalospfv3NssaLsa) ToJsonRaw() (string, error) {
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

func (m *marshalospfv3NssaLsa) ToJson() (string, error) {
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

func (m *unMarshalospfv3NssaLsa) FromJson(value string) error {
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

func (obj *ospfv3NssaLsa) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3NssaLsa) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3NssaLsa) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3NssaLsa) Clone() (Ospfv3NssaLsa, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3NssaLsa()
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

func (obj *ospfv3NssaLsa) setNil() {
	obj.headerHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv3NssaLsa is contents of OSPFv3 NSSA LSA - Type 7.
type Ospfv3NssaLsa interface {
	Validation
	// msg marshals Ospfv3NssaLsa to protobuf object *otg.Ospfv3NssaLsa
	// and doesn't set defaults
	msg() *otg.Ospfv3NssaLsa
	// setMsg unmarshals Ospfv3NssaLsa from protobuf object *otg.Ospfv3NssaLsa
	// and doesn't set defaults
	setMsg(*otg.Ospfv3NssaLsa) Ospfv3NssaLsa
	// provides marshal interface
	Marshal() marshalOspfv3NssaLsa
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3NssaLsa
	// validate validates Ospfv3NssaLsa
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3NssaLsa, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Header returns Ospfv3LsaHeader, set in Ospfv3NssaLsa.
	// Ospfv3LsaHeader is attributes in LSA Header.
	Header() Ospfv3LsaHeader
	// SetHeader assigns Ospfv3LsaHeader provided by user to Ospfv3NssaLsa.
	// Ospfv3LsaHeader is attributes in LSA Header.
	SetHeader(value Ospfv3LsaHeader) Ospfv3NssaLsa
	// HasHeader checks if Header has been set in Ospfv3NssaLsa
	HasHeader() bool
	// AddressPrefix returns string, set in Ospfv3NssaLsa.
	AddressPrefix() string
	// SetAddressPrefix assigns string provided by user to Ospfv3NssaLsa
	SetAddressPrefix(value string) Ospfv3NssaLsa
	// HasAddressPrefix checks if AddressPrefix has been set in Ospfv3NssaLsa
	HasAddressPrefix() bool
	// PrefixLength returns uint32, set in Ospfv3NssaLsa.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to Ospfv3NssaLsa
	SetPrefixLength(value uint32) Ospfv3NssaLsa
	// HasPrefixLength checks if PrefixLength has been set in Ospfv3NssaLsa
	HasPrefixLength() bool
	// Metric returns uint32, set in Ospfv3NssaLsa.
	Metric() uint32
	// SetMetric assigns uint32 provided by user to Ospfv3NssaLsa
	SetMetric(value uint32) Ospfv3NssaLsa
	// HasMetric checks if Metric has been set in Ospfv3NssaLsa
	HasMetric() bool
	// ForwardingAddress returns string, set in Ospfv3NssaLsa.
	ForwardingAddress() string
	// SetForwardingAddress assigns string provided by user to Ospfv3NssaLsa
	SetForwardingAddress(value string) Ospfv3NssaLsa
	// HasForwardingAddress checks if ForwardingAddress has been set in Ospfv3NssaLsa
	HasForwardingAddress() bool
	// RouteTag returns string, set in Ospfv3NssaLsa.
	RouteTag() string
	// SetRouteTag assigns string provided by user to Ospfv3NssaLsa
	SetRouteTag(value string) Ospfv3NssaLsa
	// HasRouteTag checks if RouteTag has been set in Ospfv3NssaLsa
	HasRouteTag() bool
	setNil()
}

// Contents of the LSA header.
// Header returns a Ospfv3LsaHeader
func (obj *ospfv3NssaLsa) Header() Ospfv3LsaHeader {
	if obj.obj.Header == nil {
		obj.obj.Header = NewOspfv3LsaHeader().msg()
	}
	if obj.headerHolder == nil {
		obj.headerHolder = &ospfv3LsaHeader{obj: obj.obj.Header}
	}
	return obj.headerHolder
}

// Contents of the LSA header.
// Header returns a Ospfv3LsaHeader
func (obj *ospfv3NssaLsa) HasHeader() bool {
	return obj.obj.Header != nil
}

// Contents of the LSA header.
// SetHeader sets the Ospfv3LsaHeader value in the Ospfv3NssaLsa object
func (obj *ospfv3NssaLsa) SetHeader(value Ospfv3LsaHeader) Ospfv3NssaLsa {

	obj.headerHolder = nil
	obj.obj.Header = value.msg()

	return obj
}

// The first IPv6 address prefix to be advertised in the LSA.
// AddressPrefix returns a string
func (obj *ospfv3NssaLsa) AddressPrefix() string {

	return *obj.obj.AddressPrefix

}

// The first IPv6 address prefix to be advertised in the LSA.
// AddressPrefix returns a string
func (obj *ospfv3NssaLsa) HasAddressPrefix() bool {
	return obj.obj.AddressPrefix != nil
}

// The first IPv6 address prefix to be advertised in the LSA.
// SetAddressPrefix sets the string value in the Ospfv3NssaLsa object
func (obj *ospfv3NssaLsa) SetAddressPrefix(value string) Ospfv3NssaLsa {

	obj.obj.AddressPrefix = &value
	return obj
}

// The length of the IPv6 address prefix, in bits.
// PrefixLength returns a uint32
func (obj *ospfv3NssaLsa) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// The length of the IPv6 address prefix, in bits.
// PrefixLength returns a uint32
func (obj *ospfv3NssaLsa) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// The length of the IPv6 address prefix, in bits.
// SetPrefixLength sets the uint32 value in the Ospfv3NssaLsa object
func (obj *ospfv3NssaLsa) SetPrefixLength(value uint32) Ospfv3NssaLsa {

	obj.obj.PrefixLength = &value
	return obj
}

// The cost metric value for the route to this destination router.
// Metric returns a uint32
func (obj *ospfv3NssaLsa) Metric() uint32 {

	return *obj.obj.Metric

}

// The cost metric value for the route to this destination router.
// Metric returns a uint32
func (obj *ospfv3NssaLsa) HasMetric() bool {
	return obj.obj.Metric != nil
}

// The cost metric value for the route to this destination router.
// SetMetric sets the uint32 value in the Ospfv3NssaLsa object
func (obj *ospfv3NssaLsa) SetMetric(value uint32) Ospfv3NssaLsa {

	obj.obj.Metric = &value
	return obj
}

// The IPV6 address where traffic for the advertised destination is forwarded.
// ForwardingAddress returns a string
func (obj *ospfv3NssaLsa) ForwardingAddress() string {

	return *obj.obj.ForwardingAddress

}

// The IPV6 address where traffic for the advertised destination is forwarded.
// ForwardingAddress returns a string
func (obj *ospfv3NssaLsa) HasForwardingAddress() bool {
	return obj.obj.ForwardingAddress != nil
}

// The IPV6 address where traffic for the advertised destination is forwarded.
// SetForwardingAddress sets the string value in the Ospfv3NssaLsa object
func (obj *ospfv3NssaLsa) SetForwardingAddress(value string) Ospfv3NssaLsa {

	obj.obj.ForwardingAddress = &value
	return obj
}

// The optional field may be used to communicate additional information between  AS boundary routers.
// RouteTag returns a string
func (obj *ospfv3NssaLsa) RouteTag() string {

	return *obj.obj.RouteTag

}

// The optional field may be used to communicate additional information between  AS boundary routers.
// RouteTag returns a string
func (obj *ospfv3NssaLsa) HasRouteTag() bool {
	return obj.obj.RouteTag != nil
}

// The optional field may be used to communicate additional information between  AS boundary routers.
// SetRouteTag sets the string value in the Ospfv3NssaLsa object
func (obj *ospfv3NssaLsa) SetRouteTag(value string) Ospfv3NssaLsa {

	obj.obj.RouteTag = &value
	return obj
}

func (obj *ospfv3NssaLsa) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Header != nil {

		obj.Header().validateObj(vObj, set_default)
	}

	if obj.obj.AddressPrefix != nil {

		err := obj.validateIpv6(obj.AddressPrefix())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv3NssaLsa.AddressPrefix"))
		}

	}

	if obj.obj.ForwardingAddress != nil {

		err := obj.validateIpv6(obj.ForwardingAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv3NssaLsa.ForwardingAddress"))
		}

	}

	if obj.obj.RouteTag != nil {

		err := obj.validateIpv4(obj.RouteTag())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv3NssaLsa.RouteTag"))
		}

	}

}

func (obj *ospfv3NssaLsa) setDefault() {

}
