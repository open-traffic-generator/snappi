package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3ExternalAsLsa *****
type ospfv3ExternalAsLsa struct {
	validation
	obj          *otg.Ospfv3ExternalAsLsa
	marshaller   marshalOspfv3ExternalAsLsa
	unMarshaller unMarshalOspfv3ExternalAsLsa
	headerHolder Ospfv3LsaHeader
}

func NewOspfv3ExternalAsLsa() Ospfv3ExternalAsLsa {
	obj := ospfv3ExternalAsLsa{obj: &otg.Ospfv3ExternalAsLsa{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3ExternalAsLsa) msg() *otg.Ospfv3ExternalAsLsa {
	return obj.obj
}

func (obj *ospfv3ExternalAsLsa) setMsg(msg *otg.Ospfv3ExternalAsLsa) Ospfv3ExternalAsLsa {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3ExternalAsLsa struct {
	obj *ospfv3ExternalAsLsa
}

type marshalOspfv3ExternalAsLsa interface {
	// ToProto marshals Ospfv3ExternalAsLsa to protobuf object *otg.Ospfv3ExternalAsLsa
	ToProto() (*otg.Ospfv3ExternalAsLsa, error)
	// ToPbText marshals Ospfv3ExternalAsLsa to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3ExternalAsLsa to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3ExternalAsLsa to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Ospfv3ExternalAsLsa to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalospfv3ExternalAsLsa struct {
	obj *ospfv3ExternalAsLsa
}

type unMarshalOspfv3ExternalAsLsa interface {
	// FromProto unmarshals Ospfv3ExternalAsLsa from protobuf object *otg.Ospfv3ExternalAsLsa
	FromProto(msg *otg.Ospfv3ExternalAsLsa) (Ospfv3ExternalAsLsa, error)
	// FromPbText unmarshals Ospfv3ExternalAsLsa from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3ExternalAsLsa from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3ExternalAsLsa from JSON text
	FromJson(value string) error
}

func (obj *ospfv3ExternalAsLsa) Marshal() marshalOspfv3ExternalAsLsa {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3ExternalAsLsa{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3ExternalAsLsa) Unmarshal() unMarshalOspfv3ExternalAsLsa {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3ExternalAsLsa{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3ExternalAsLsa) ToProto() (*otg.Ospfv3ExternalAsLsa, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3ExternalAsLsa) FromProto(msg *otg.Ospfv3ExternalAsLsa) (Ospfv3ExternalAsLsa, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3ExternalAsLsa) ToPbText() (string, error) {
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

func (m *unMarshalospfv3ExternalAsLsa) FromPbText(value string) error {
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

func (m *marshalospfv3ExternalAsLsa) ToYaml() (string, error) {
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

func (m *unMarshalospfv3ExternalAsLsa) FromYaml(value string) error {
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

func (m *marshalospfv3ExternalAsLsa) ToJsonRaw() (string, error) {
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

func (m *marshalospfv3ExternalAsLsa) ToJson() (string, error) {
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

func (m *unMarshalospfv3ExternalAsLsa) FromJson(value string) error {
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

func (obj *ospfv3ExternalAsLsa) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3ExternalAsLsa) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3ExternalAsLsa) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3ExternalAsLsa) Clone() (Ospfv3ExternalAsLsa, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3ExternalAsLsa()
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

func (obj *ospfv3ExternalAsLsa) setNil() {
	obj.headerHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv3ExternalAsLsa is contents of OSPFv3 AS-External-LSA - Type 5.
type Ospfv3ExternalAsLsa interface {
	Validation
	// msg marshals Ospfv3ExternalAsLsa to protobuf object *otg.Ospfv3ExternalAsLsa
	// and doesn't set defaults
	msg() *otg.Ospfv3ExternalAsLsa
	// setMsg unmarshals Ospfv3ExternalAsLsa from protobuf object *otg.Ospfv3ExternalAsLsa
	// and doesn't set defaults
	setMsg(*otg.Ospfv3ExternalAsLsa) Ospfv3ExternalAsLsa
	// provides marshal interface
	Marshal() marshalOspfv3ExternalAsLsa
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3ExternalAsLsa
	// validate validates Ospfv3ExternalAsLsa
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3ExternalAsLsa, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Header returns Ospfv3LsaHeader, set in Ospfv3ExternalAsLsa.
	// Ospfv3LsaHeader is attributes in LSA Header.
	Header() Ospfv3LsaHeader
	// SetHeader assigns Ospfv3LsaHeader provided by user to Ospfv3ExternalAsLsa.
	// Ospfv3LsaHeader is attributes in LSA Header.
	SetHeader(value Ospfv3LsaHeader) Ospfv3ExternalAsLsa
	// HasHeader checks if Header has been set in Ospfv3ExternalAsLsa
	HasHeader() bool
	// AddressPrefix returns string, set in Ospfv3ExternalAsLsa.
	AddressPrefix() string
	// SetAddressPrefix assigns string provided by user to Ospfv3ExternalAsLsa
	SetAddressPrefix(value string) Ospfv3ExternalAsLsa
	// HasAddressPrefix checks if AddressPrefix has been set in Ospfv3ExternalAsLsa
	HasAddressPrefix() bool
	// PrefixLength returns uint32, set in Ospfv3ExternalAsLsa.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to Ospfv3ExternalAsLsa
	SetPrefixLength(value uint32) Ospfv3ExternalAsLsa
	// HasPrefixLength checks if PrefixLength has been set in Ospfv3ExternalAsLsa
	HasPrefixLength() bool
	// Metric returns uint32, set in Ospfv3ExternalAsLsa.
	Metric() uint32
	// SetMetric assigns uint32 provided by user to Ospfv3ExternalAsLsa
	SetMetric(value uint32) Ospfv3ExternalAsLsa
	// HasMetric checks if Metric has been set in Ospfv3ExternalAsLsa
	HasMetric() bool
	// ReferencedLsType returns uint32, set in Ospfv3ExternalAsLsa.
	ReferencedLsType() uint32
	// SetReferencedLsType assigns uint32 provided by user to Ospfv3ExternalAsLsa
	SetReferencedLsType(value uint32) Ospfv3ExternalAsLsa
	// HasReferencedLsType checks if ReferencedLsType has been set in Ospfv3ExternalAsLsa
	HasReferencedLsType() bool
	// ForwardingAddress returns string, set in Ospfv3ExternalAsLsa.
	ForwardingAddress() string
	// SetForwardingAddress assigns string provided by user to Ospfv3ExternalAsLsa
	SetForwardingAddress(value string) Ospfv3ExternalAsLsa
	// HasForwardingAddress checks if ForwardingAddress has been set in Ospfv3ExternalAsLsa
	HasForwardingAddress() bool
	// RouteTag returns string, set in Ospfv3ExternalAsLsa.
	RouteTag() string
	// SetRouteTag assigns string provided by user to Ospfv3ExternalAsLsa
	SetRouteTag(value string) Ospfv3ExternalAsLsa
	// HasRouteTag checks if RouteTag has been set in Ospfv3ExternalAsLsa
	HasRouteTag() bool
	setNil()
}

// Contents of the LSA header.
// Header returns a Ospfv3LsaHeader
func (obj *ospfv3ExternalAsLsa) Header() Ospfv3LsaHeader {
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
func (obj *ospfv3ExternalAsLsa) HasHeader() bool {
	return obj.obj.Header != nil
}

// Contents of the LSA header.
// SetHeader sets the Ospfv3LsaHeader value in the Ospfv3ExternalAsLsa object
func (obj *ospfv3ExternalAsLsa) SetHeader(value Ospfv3LsaHeader) Ospfv3ExternalAsLsa {

	obj.headerHolder = nil
	obj.obj.Header = value.msg()

	return obj
}

// The first IPv6 address prefix to be advertised in the LSA.
// AddressPrefix returns a string
func (obj *ospfv3ExternalAsLsa) AddressPrefix() string {

	return *obj.obj.AddressPrefix

}

// The first IPv6 address prefix to be advertised in the LSA.
// AddressPrefix returns a string
func (obj *ospfv3ExternalAsLsa) HasAddressPrefix() bool {
	return obj.obj.AddressPrefix != nil
}

// The first IPv6 address prefix to be advertised in the LSA.
// SetAddressPrefix sets the string value in the Ospfv3ExternalAsLsa object
func (obj *ospfv3ExternalAsLsa) SetAddressPrefix(value string) Ospfv3ExternalAsLsa {

	obj.obj.AddressPrefix = &value
	return obj
}

// The length of the IPv6 address prefix, in bits.
// PrefixLength returns a uint32
func (obj *ospfv3ExternalAsLsa) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// The length of the IPv6 address prefix, in bits.
// PrefixLength returns a uint32
func (obj *ospfv3ExternalAsLsa) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// The length of the IPv6 address prefix, in bits.
// SetPrefixLength sets the uint32 value in the Ospfv3ExternalAsLsa object
func (obj *ospfv3ExternalAsLsa) SetPrefixLength(value uint32) Ospfv3ExternalAsLsa {

	obj.obj.PrefixLength = &value
	return obj
}

// The cost metric value for the route to this destination router.
// Metric returns a uint32
func (obj *ospfv3ExternalAsLsa) Metric() uint32 {

	return *obj.obj.Metric

}

// The cost metric value for the route to this destination router.
// Metric returns a uint32
func (obj *ospfv3ExternalAsLsa) HasMetric() bool {
	return obj.obj.Metric != nil
}

// The cost metric value for the route to this destination router.
// SetMetric sets the uint32 value in the Ospfv3ExternalAsLsa object
func (obj *ospfv3ExternalAsLsa) SetMetric(value uint32) Ospfv3ExternalAsLsa {

	obj.obj.Metric = &value
	return obj
}

// If non-zero, an LSA with this LS type is to be associated with this LSA.
// ReferencedLsType returns a uint32
func (obj *ospfv3ExternalAsLsa) ReferencedLsType() uint32 {

	return *obj.obj.ReferencedLsType

}

// If non-zero, an LSA with this LS type is to be associated with this LSA.
// ReferencedLsType returns a uint32
func (obj *ospfv3ExternalAsLsa) HasReferencedLsType() bool {
	return obj.obj.ReferencedLsType != nil
}

// If non-zero, an LSA with this LS type is to be associated with this LSA.
// SetReferencedLsType sets the uint32 value in the Ospfv3ExternalAsLsa object
func (obj *ospfv3ExternalAsLsa) SetReferencedLsType(value uint32) Ospfv3ExternalAsLsa {

	obj.obj.ReferencedLsType = &value
	return obj
}

// The IPV6 address where traffic for the advertised destination is forwarded.
// ForwardingAddress returns a string
func (obj *ospfv3ExternalAsLsa) ForwardingAddress() string {

	return *obj.obj.ForwardingAddress

}

// The IPV6 address where traffic for the advertised destination is forwarded.
// ForwardingAddress returns a string
func (obj *ospfv3ExternalAsLsa) HasForwardingAddress() bool {
	return obj.obj.ForwardingAddress != nil
}

// The IPV6 address where traffic for the advertised destination is forwarded.
// SetForwardingAddress sets the string value in the Ospfv3ExternalAsLsa object
func (obj *ospfv3ExternalAsLsa) SetForwardingAddress(value string) Ospfv3ExternalAsLsa {

	obj.obj.ForwardingAddress = &value
	return obj
}

// The optional field may be used to communicate additional information between  AS boundary routers.
// RouteTag returns a string
func (obj *ospfv3ExternalAsLsa) RouteTag() string {

	return *obj.obj.RouteTag

}

// The optional field may be used to communicate additional information between  AS boundary routers.
// RouteTag returns a string
func (obj *ospfv3ExternalAsLsa) HasRouteTag() bool {
	return obj.obj.RouteTag != nil
}

// The optional field may be used to communicate additional information between  AS boundary routers.
// SetRouteTag sets the string value in the Ospfv3ExternalAsLsa object
func (obj *ospfv3ExternalAsLsa) SetRouteTag(value string) Ospfv3ExternalAsLsa {

	obj.obj.RouteTag = &value
	return obj
}

func (obj *ospfv3ExternalAsLsa) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Header != nil {

		obj.Header().validateObj(vObj, set_default)
	}

	if obj.obj.AddressPrefix != nil {

		err := obj.validateIpv6(obj.AddressPrefix())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv3ExternalAsLsa.AddressPrefix"))
		}

	}

	if obj.obj.ForwardingAddress != nil {

		err := obj.validateIpv6(obj.ForwardingAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv3ExternalAsLsa.ForwardingAddress"))
		}

	}

	if obj.obj.RouteTag != nil {

		err := obj.validateIpv4(obj.RouteTag())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv3ExternalAsLsa.RouteTag"))
		}

	}

}

func (obj *ospfv3ExternalAsLsa) setDefault() {

}
