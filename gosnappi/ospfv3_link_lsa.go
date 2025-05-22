package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3LinkLsa *****
type ospfv3LinkLsa struct {
	validation
	obj          *otg.Ospfv3LinkLsa
	marshaller   marshalOspfv3LinkLsa
	unMarshaller unMarshalOspfv3LinkLsa
	headerHolder Ospfv3LsaHeader
}

func NewOspfv3LinkLsa() Ospfv3LinkLsa {
	obj := ospfv3LinkLsa{obj: &otg.Ospfv3LinkLsa{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3LinkLsa) msg() *otg.Ospfv3LinkLsa {
	return obj.obj
}

func (obj *ospfv3LinkLsa) setMsg(msg *otg.Ospfv3LinkLsa) Ospfv3LinkLsa {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3LinkLsa struct {
	obj *ospfv3LinkLsa
}

type marshalOspfv3LinkLsa interface {
	// ToProto marshals Ospfv3LinkLsa to protobuf object *otg.Ospfv3LinkLsa
	ToProto() (*otg.Ospfv3LinkLsa, error)
	// ToPbText marshals Ospfv3LinkLsa to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3LinkLsa to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3LinkLsa to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Ospfv3LinkLsa to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalospfv3LinkLsa struct {
	obj *ospfv3LinkLsa
}

type unMarshalOspfv3LinkLsa interface {
	// FromProto unmarshals Ospfv3LinkLsa from protobuf object *otg.Ospfv3LinkLsa
	FromProto(msg *otg.Ospfv3LinkLsa) (Ospfv3LinkLsa, error)
	// FromPbText unmarshals Ospfv3LinkLsa from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3LinkLsa from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3LinkLsa from JSON text
	FromJson(value string) error
}

func (obj *ospfv3LinkLsa) Marshal() marshalOspfv3LinkLsa {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3LinkLsa{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3LinkLsa) Unmarshal() unMarshalOspfv3LinkLsa {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3LinkLsa{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3LinkLsa) ToProto() (*otg.Ospfv3LinkLsa, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3LinkLsa) FromProto(msg *otg.Ospfv3LinkLsa) (Ospfv3LinkLsa, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3LinkLsa) ToPbText() (string, error) {
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

func (m *unMarshalospfv3LinkLsa) FromPbText(value string) error {
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

func (m *marshalospfv3LinkLsa) ToYaml() (string, error) {
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

func (m *unMarshalospfv3LinkLsa) FromYaml(value string) error {
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

func (m *marshalospfv3LinkLsa) ToJsonRaw() (string, error) {
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

func (m *marshalospfv3LinkLsa) ToJson() (string, error) {
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

func (m *unMarshalospfv3LinkLsa) FromJson(value string) error {
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

func (obj *ospfv3LinkLsa) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3LinkLsa) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3LinkLsa) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3LinkLsa) Clone() (Ospfv3LinkLsa, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3LinkLsa()
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

func (obj *ospfv3LinkLsa) setNil() {
	obj.headerHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv3LinkLsa is contents of OSPFv3 Link LSA - Type 8.
type Ospfv3LinkLsa interface {
	Validation
	// msg marshals Ospfv3LinkLsa to protobuf object *otg.Ospfv3LinkLsa
	// and doesn't set defaults
	msg() *otg.Ospfv3LinkLsa
	// setMsg unmarshals Ospfv3LinkLsa from protobuf object *otg.Ospfv3LinkLsa
	// and doesn't set defaults
	setMsg(*otg.Ospfv3LinkLsa) Ospfv3LinkLsa
	// provides marshal interface
	Marshal() marshalOspfv3LinkLsa
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3LinkLsa
	// validate validates Ospfv3LinkLsa
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3LinkLsa, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Header returns Ospfv3LsaHeader, set in Ospfv3LinkLsa.
	// Ospfv3LsaHeader is attributes in LSA Header.
	Header() Ospfv3LsaHeader
	// SetHeader assigns Ospfv3LsaHeader provided by user to Ospfv3LinkLsa.
	// Ospfv3LsaHeader is attributes in LSA Header.
	SetHeader(value Ospfv3LsaHeader) Ospfv3LinkLsa
	// HasHeader checks if Header has been set in Ospfv3LinkLsa
	HasHeader() bool
	// AddressPrefix returns string, set in Ospfv3LinkLsa.
	AddressPrefix() string
	// SetAddressPrefix assigns string provided by user to Ospfv3LinkLsa
	SetAddressPrefix(value string) Ospfv3LinkLsa
	// HasAddressPrefix checks if AddressPrefix has been set in Ospfv3LinkLsa
	HasAddressPrefix() bool
	// PrefixLength returns uint32, set in Ospfv3LinkLsa.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to Ospfv3LinkLsa
	SetPrefixLength(value uint32) Ospfv3LinkLsa
	// HasPrefixLength checks if PrefixLength has been set in Ospfv3LinkLsa
	HasPrefixLength() bool
	// LinkLocalAddress returns string, set in Ospfv3LinkLsa.
	LinkLocalAddress() string
	// SetLinkLocalAddress assigns string provided by user to Ospfv3LinkLsa
	SetLinkLocalAddress(value string) Ospfv3LinkLsa
	// HasLinkLocalAddress checks if LinkLocalAddress has been set in Ospfv3LinkLsa
	HasLinkLocalAddress() bool
	setNil()
}

// Contents of the LSA header.
// Header returns a Ospfv3LsaHeader
func (obj *ospfv3LinkLsa) Header() Ospfv3LsaHeader {
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
func (obj *ospfv3LinkLsa) HasHeader() bool {
	return obj.obj.Header != nil
}

// Contents of the LSA header.
// SetHeader sets the Ospfv3LsaHeader value in the Ospfv3LinkLsa object
func (obj *ospfv3LinkLsa) SetHeader(value Ospfv3LsaHeader) Ospfv3LinkLsa {

	obj.headerHolder = nil
	obj.obj.Header = value.msg()

	return obj
}

// The first IPv6 address prefix to be advertised in the LSA.
// AddressPrefix returns a string
func (obj *ospfv3LinkLsa) AddressPrefix() string {

	return *obj.obj.AddressPrefix

}

// The first IPv6 address prefix to be advertised in the LSA.
// AddressPrefix returns a string
func (obj *ospfv3LinkLsa) HasAddressPrefix() bool {
	return obj.obj.AddressPrefix != nil
}

// The first IPv6 address prefix to be advertised in the LSA.
// SetAddressPrefix sets the string value in the Ospfv3LinkLsa object
func (obj *ospfv3LinkLsa) SetAddressPrefix(value string) Ospfv3LinkLsa {

	obj.obj.AddressPrefix = &value
	return obj
}

// The length of the IPv6 address prefix, in bits.
// PrefixLength returns a uint32
func (obj *ospfv3LinkLsa) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// The length of the IPv6 address prefix, in bits.
// PrefixLength returns a uint32
func (obj *ospfv3LinkLsa) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// The length of the IPv6 address prefix, in bits.
// SetPrefixLength sets the uint32 value in the Ospfv3LinkLsa object
func (obj *ospfv3LinkLsa) SetPrefixLength(value uint32) Ospfv3LinkLsa {

	obj.obj.PrefixLength = &value
	return obj
}

// The IPV6 Link Local address for the originating router's interface attached to this link.
// LinkLocalAddress returns a string
func (obj *ospfv3LinkLsa) LinkLocalAddress() string {

	return *obj.obj.LinkLocalAddress

}

// The IPV6 Link Local address for the originating router's interface attached to this link.
// LinkLocalAddress returns a string
func (obj *ospfv3LinkLsa) HasLinkLocalAddress() bool {
	return obj.obj.LinkLocalAddress != nil
}

// The IPV6 Link Local address for the originating router's interface attached to this link.
// SetLinkLocalAddress sets the string value in the Ospfv3LinkLsa object
func (obj *ospfv3LinkLsa) SetLinkLocalAddress(value string) Ospfv3LinkLsa {

	obj.obj.LinkLocalAddress = &value
	return obj
}

func (obj *ospfv3LinkLsa) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Header != nil {

		obj.Header().validateObj(vObj, set_default)
	}

	if obj.obj.AddressPrefix != nil {

		err := obj.validateIpv6(obj.AddressPrefix())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv3LinkLsa.AddressPrefix"))
		}

	}

	if obj.obj.LinkLocalAddress != nil {

		err := obj.validateIpv6(obj.LinkLocalAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv3LinkLsa.LinkLocalAddress"))
		}

	}

}

func (obj *ospfv3LinkLsa) setDefault() {

}
