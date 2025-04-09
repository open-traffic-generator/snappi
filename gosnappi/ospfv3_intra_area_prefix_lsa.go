package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3IntraAreaPrefixLsa *****
type ospfv3IntraAreaPrefixLsa struct {
	validation
	obj          *otg.Ospfv3IntraAreaPrefixLsa
	marshaller   marshalOspfv3IntraAreaPrefixLsa
	unMarshaller unMarshalOspfv3IntraAreaPrefixLsa
	headerHolder Ospfv3LsaHeader
}

func NewOspfv3IntraAreaPrefixLsa() Ospfv3IntraAreaPrefixLsa {
	obj := ospfv3IntraAreaPrefixLsa{obj: &otg.Ospfv3IntraAreaPrefixLsa{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3IntraAreaPrefixLsa) msg() *otg.Ospfv3IntraAreaPrefixLsa {
	return obj.obj
}

func (obj *ospfv3IntraAreaPrefixLsa) setMsg(msg *otg.Ospfv3IntraAreaPrefixLsa) Ospfv3IntraAreaPrefixLsa {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3IntraAreaPrefixLsa struct {
	obj *ospfv3IntraAreaPrefixLsa
}

type marshalOspfv3IntraAreaPrefixLsa interface {
	// ToProto marshals Ospfv3IntraAreaPrefixLsa to protobuf object *otg.Ospfv3IntraAreaPrefixLsa
	ToProto() (*otg.Ospfv3IntraAreaPrefixLsa, error)
	// ToPbText marshals Ospfv3IntraAreaPrefixLsa to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3IntraAreaPrefixLsa to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3IntraAreaPrefixLsa to JSON text
	ToJson() (string, error)
}

type unMarshalospfv3IntraAreaPrefixLsa struct {
	obj *ospfv3IntraAreaPrefixLsa
}

type unMarshalOspfv3IntraAreaPrefixLsa interface {
	// FromProto unmarshals Ospfv3IntraAreaPrefixLsa from protobuf object *otg.Ospfv3IntraAreaPrefixLsa
	FromProto(msg *otg.Ospfv3IntraAreaPrefixLsa) (Ospfv3IntraAreaPrefixLsa, error)
	// FromPbText unmarshals Ospfv3IntraAreaPrefixLsa from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3IntraAreaPrefixLsa from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3IntraAreaPrefixLsa from JSON text
	FromJson(value string) error
}

func (obj *ospfv3IntraAreaPrefixLsa) Marshal() marshalOspfv3IntraAreaPrefixLsa {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3IntraAreaPrefixLsa{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3IntraAreaPrefixLsa) Unmarshal() unMarshalOspfv3IntraAreaPrefixLsa {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3IntraAreaPrefixLsa{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3IntraAreaPrefixLsa) ToProto() (*otg.Ospfv3IntraAreaPrefixLsa, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3IntraAreaPrefixLsa) FromProto(msg *otg.Ospfv3IntraAreaPrefixLsa) (Ospfv3IntraAreaPrefixLsa, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3IntraAreaPrefixLsa) ToPbText() (string, error) {
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

func (m *unMarshalospfv3IntraAreaPrefixLsa) FromPbText(value string) error {
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

func (m *marshalospfv3IntraAreaPrefixLsa) ToYaml() (string, error) {
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

func (m *unMarshalospfv3IntraAreaPrefixLsa) FromYaml(value string) error {
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

func (m *marshalospfv3IntraAreaPrefixLsa) ToJson() (string, error) {
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

func (m *unMarshalospfv3IntraAreaPrefixLsa) FromJson(value string) error {
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

func (obj *ospfv3IntraAreaPrefixLsa) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3IntraAreaPrefixLsa) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3IntraAreaPrefixLsa) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3IntraAreaPrefixLsa) Clone() (Ospfv3IntraAreaPrefixLsa, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3IntraAreaPrefixLsa()
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

func (obj *ospfv3IntraAreaPrefixLsa) setNil() {
	obj.headerHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv3IntraAreaPrefixLsa is contents of OSPFv3 Intra-Area-Prefix LSA - Type 9.
type Ospfv3IntraAreaPrefixLsa interface {
	Validation
	// msg marshals Ospfv3IntraAreaPrefixLsa to protobuf object *otg.Ospfv3IntraAreaPrefixLsa
	// and doesn't set defaults
	msg() *otg.Ospfv3IntraAreaPrefixLsa
	// setMsg unmarshals Ospfv3IntraAreaPrefixLsa from protobuf object *otg.Ospfv3IntraAreaPrefixLsa
	// and doesn't set defaults
	setMsg(*otg.Ospfv3IntraAreaPrefixLsa) Ospfv3IntraAreaPrefixLsa
	// provides marshal interface
	Marshal() marshalOspfv3IntraAreaPrefixLsa
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3IntraAreaPrefixLsa
	// validate validates Ospfv3IntraAreaPrefixLsa
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3IntraAreaPrefixLsa, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Header returns Ospfv3LsaHeader, set in Ospfv3IntraAreaPrefixLsa.
	// Ospfv3LsaHeader is attributes in LSA Header.
	Header() Ospfv3LsaHeader
	// SetHeader assigns Ospfv3LsaHeader provided by user to Ospfv3IntraAreaPrefixLsa.
	// Ospfv3LsaHeader is attributes in LSA Header.
	SetHeader(value Ospfv3LsaHeader) Ospfv3IntraAreaPrefixLsa
	// HasHeader checks if Header has been set in Ospfv3IntraAreaPrefixLsa
	HasHeader() bool
	// AddressPrefix returns string, set in Ospfv3IntraAreaPrefixLsa.
	AddressPrefix() string
	// SetAddressPrefix assigns string provided by user to Ospfv3IntraAreaPrefixLsa
	SetAddressPrefix(value string) Ospfv3IntraAreaPrefixLsa
	// HasAddressPrefix checks if AddressPrefix has been set in Ospfv3IntraAreaPrefixLsa
	HasAddressPrefix() bool
	// PrefixLength returns uint32, set in Ospfv3IntraAreaPrefixLsa.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to Ospfv3IntraAreaPrefixLsa
	SetPrefixLength(value uint32) Ospfv3IntraAreaPrefixLsa
	// HasPrefixLength checks if PrefixLength has been set in Ospfv3IntraAreaPrefixLsa
	HasPrefixLength() bool
	// Metric returns uint32, set in Ospfv3IntraAreaPrefixLsa.
	Metric() uint32
	// SetMetric assigns uint32 provided by user to Ospfv3IntraAreaPrefixLsa
	SetMetric(value uint32) Ospfv3IntraAreaPrefixLsa
	// HasMetric checks if Metric has been set in Ospfv3IntraAreaPrefixLsa
	HasMetric() bool
	setNil()
}

// Contents of the LSA header.
// Header returns a Ospfv3LsaHeader
func (obj *ospfv3IntraAreaPrefixLsa) Header() Ospfv3LsaHeader {
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
func (obj *ospfv3IntraAreaPrefixLsa) HasHeader() bool {
	return obj.obj.Header != nil
}

// Contents of the LSA header.
// SetHeader sets the Ospfv3LsaHeader value in the Ospfv3IntraAreaPrefixLsa object
func (obj *ospfv3IntraAreaPrefixLsa) SetHeader(value Ospfv3LsaHeader) Ospfv3IntraAreaPrefixLsa {

	obj.headerHolder = nil
	obj.obj.Header = value.msg()

	return obj
}

// The first IPv6 address prefix to be advertised in the LSA.
// AddressPrefix returns a string
func (obj *ospfv3IntraAreaPrefixLsa) AddressPrefix() string {

	return *obj.obj.AddressPrefix

}

// The first IPv6 address prefix to be advertised in the LSA.
// AddressPrefix returns a string
func (obj *ospfv3IntraAreaPrefixLsa) HasAddressPrefix() bool {
	return obj.obj.AddressPrefix != nil
}

// The first IPv6 address prefix to be advertised in the LSA.
// SetAddressPrefix sets the string value in the Ospfv3IntraAreaPrefixLsa object
func (obj *ospfv3IntraAreaPrefixLsa) SetAddressPrefix(value string) Ospfv3IntraAreaPrefixLsa {

	obj.obj.AddressPrefix = &value
	return obj
}

// The length of the IPv6 address prefix, in bits.
// PrefixLength returns a uint32
func (obj *ospfv3IntraAreaPrefixLsa) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// The length of the IPv6 address prefix, in bits.
// PrefixLength returns a uint32
func (obj *ospfv3IntraAreaPrefixLsa) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// The length of the IPv6 address prefix, in bits.
// SetPrefixLength sets the uint32 value in the Ospfv3IntraAreaPrefixLsa object
func (obj *ospfv3IntraAreaPrefixLsa) SetPrefixLength(value uint32) Ospfv3IntraAreaPrefixLsa {

	obj.obj.PrefixLength = &value
	return obj
}

// The cost metric value for the route to this destination router.
// Metric returns a uint32
func (obj *ospfv3IntraAreaPrefixLsa) Metric() uint32 {

	return *obj.obj.Metric

}

// The cost metric value for the route to this destination router.
// Metric returns a uint32
func (obj *ospfv3IntraAreaPrefixLsa) HasMetric() bool {
	return obj.obj.Metric != nil
}

// The cost metric value for the route to this destination router.
// SetMetric sets the uint32 value in the Ospfv3IntraAreaPrefixLsa object
func (obj *ospfv3IntraAreaPrefixLsa) SetMetric(value uint32) Ospfv3IntraAreaPrefixLsa {

	obj.obj.Metric = &value
	return obj
}

func (obj *ospfv3IntraAreaPrefixLsa) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Header != nil {

		obj.Header().validateObj(vObj, set_default)
	}

	if obj.obj.AddressPrefix != nil {

		err := obj.validateIpv6(obj.AddressPrefix())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv3IntraAreaPrefixLsa.AddressPrefix"))
		}

	}

}

func (obj *ospfv3IntraAreaPrefixLsa) setDefault() {

}
