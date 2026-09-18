package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2OpaqueLsaExtendedPrefix *****
type ospfv2OpaqueLsaExtendedPrefix struct {
	validation
	obj             *otg.Ospfv2OpaqueLsaExtendedPrefix
	marshaller      marshalOspfv2OpaqueLsaExtendedPrefix
	unMarshaller    unMarshalOspfv2OpaqueLsaExtendedPrefix
	prefixSidHolder Ospfv2LsaPrefixSid
}

func NewOspfv2OpaqueLsaExtendedPrefix() Ospfv2OpaqueLsaExtendedPrefix {
	obj := ospfv2OpaqueLsaExtendedPrefix{obj: &otg.Ospfv2OpaqueLsaExtendedPrefix{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2OpaqueLsaExtendedPrefix) msg() *otg.Ospfv2OpaqueLsaExtendedPrefix {
	return obj.obj
}

func (obj *ospfv2OpaqueLsaExtendedPrefix) setMsg(msg *otg.Ospfv2OpaqueLsaExtendedPrefix) Ospfv2OpaqueLsaExtendedPrefix {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2OpaqueLsaExtendedPrefix struct {
	obj *ospfv2OpaqueLsaExtendedPrefix
}

type marshalOspfv2OpaqueLsaExtendedPrefix interface {
	// ToProto marshals Ospfv2OpaqueLsaExtendedPrefix to protobuf object *otg.Ospfv2OpaqueLsaExtendedPrefix
	ToProto() (*otg.Ospfv2OpaqueLsaExtendedPrefix, error)
	// ToPbText marshals Ospfv2OpaqueLsaExtendedPrefix to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2OpaqueLsaExtendedPrefix to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2OpaqueLsaExtendedPrefix to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2OpaqueLsaExtendedPrefix struct {
	obj *ospfv2OpaqueLsaExtendedPrefix
}

type unMarshalOspfv2OpaqueLsaExtendedPrefix interface {
	// FromProto unmarshals Ospfv2OpaqueLsaExtendedPrefix from protobuf object *otg.Ospfv2OpaqueLsaExtendedPrefix
	FromProto(msg *otg.Ospfv2OpaqueLsaExtendedPrefix) (Ospfv2OpaqueLsaExtendedPrefix, error)
	// FromPbText unmarshals Ospfv2OpaqueLsaExtendedPrefix from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2OpaqueLsaExtendedPrefix from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2OpaqueLsaExtendedPrefix from JSON text
	FromJson(value string) error
}

func (obj *ospfv2OpaqueLsaExtendedPrefix) Marshal() marshalOspfv2OpaqueLsaExtendedPrefix {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2OpaqueLsaExtendedPrefix{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2OpaqueLsaExtendedPrefix) Unmarshal() unMarshalOspfv2OpaqueLsaExtendedPrefix {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2OpaqueLsaExtendedPrefix{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2OpaqueLsaExtendedPrefix) ToProto() (*otg.Ospfv2OpaqueLsaExtendedPrefix, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2OpaqueLsaExtendedPrefix) FromProto(msg *otg.Ospfv2OpaqueLsaExtendedPrefix) (Ospfv2OpaqueLsaExtendedPrefix, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2OpaqueLsaExtendedPrefix) ToPbText() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaExtendedPrefix) FromPbText(value string) error {
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

func (m *marshalospfv2OpaqueLsaExtendedPrefix) ToYaml() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaExtendedPrefix) FromYaml(value string) error {
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

func (m *marshalospfv2OpaqueLsaExtendedPrefix) ToJson() (string, error) {
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

func (m *unMarshalospfv2OpaqueLsaExtendedPrefix) FromJson(value string) error {
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

func (obj *ospfv2OpaqueLsaExtendedPrefix) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2OpaqueLsaExtendedPrefix) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2OpaqueLsaExtendedPrefix) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2OpaqueLsaExtendedPrefix) Clone() (Ospfv2OpaqueLsaExtendedPrefix, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2OpaqueLsaExtendedPrefix()
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

func (obj *ospfv2OpaqueLsaExtendedPrefix) setNil() {
	obj.prefixSidHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2OpaqueLsaExtendedPrefix is a decoded OSPFv2 Extended Prefix TLV of an Extended Prefix Opaque LSA, TLV type 1
// (RFC 7684 Section 2.1).
// The TLV names the prefix it applies to in its own body, so no reference to another LSA
// is needed to identify it. Correlate it to the LSA that advertises that prefix by
// matching prefix against router_lsas[].links[].id of a stub link for an intra-area
// prefix, or against header.lsa_id of network_summary_lsas[] for an inter-area prefix,
// external_as_lsas[] for an AS-external prefix or nssa_lsas[] for an NSSA-external
// prefix. RFC 2328 Sections 12.4.3 and 12.4.4 define the Link State ID of a Summary-LSA
// and of an AS-External-LSA as the prefix itself, and RFC 3101 gives the NSSA-LSA the
// same format, so that correlation is a direct value match rather than a new lookup.
type Ospfv2OpaqueLsaExtendedPrefix interface {
	Validation
	// msg marshals Ospfv2OpaqueLsaExtendedPrefix to protobuf object *otg.Ospfv2OpaqueLsaExtendedPrefix
	// and doesn't set defaults
	msg() *otg.Ospfv2OpaqueLsaExtendedPrefix
	// setMsg unmarshals Ospfv2OpaqueLsaExtendedPrefix from protobuf object *otg.Ospfv2OpaqueLsaExtendedPrefix
	// and doesn't set defaults
	setMsg(*otg.Ospfv2OpaqueLsaExtendedPrefix) Ospfv2OpaqueLsaExtendedPrefix
	// provides marshal interface
	Marshal() marshalOspfv2OpaqueLsaExtendedPrefix
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2OpaqueLsaExtendedPrefix
	// validate validates Ospfv2OpaqueLsaExtendedPrefix
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2OpaqueLsaExtendedPrefix, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Prefix returns string, set in Ospfv2OpaqueLsaExtendedPrefix.
	Prefix() string
	// SetPrefix assigns string provided by user to Ospfv2OpaqueLsaExtendedPrefix
	SetPrefix(value string) Ospfv2OpaqueLsaExtendedPrefix
	// HasPrefix checks if Prefix has been set in Ospfv2OpaqueLsaExtendedPrefix
	HasPrefix() bool
	// PrefixLength returns uint32, set in Ospfv2OpaqueLsaExtendedPrefix.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to Ospfv2OpaqueLsaExtendedPrefix
	SetPrefixLength(value uint32) Ospfv2OpaqueLsaExtendedPrefix
	// HasPrefixLength checks if PrefixLength has been set in Ospfv2OpaqueLsaExtendedPrefix
	HasPrefixLength() bool
	// PrefixSid returns Ospfv2LsaPrefixSid, set in Ospfv2OpaqueLsaExtendedPrefix.
	// Ospfv2LsaPrefixSid is the learned OSPFv2 Prefix-SID and its attributes, decoded from the Prefix-SID sub-TLV of
	// the Extended Prefix Opaque LSA (RFC 8665).
	PrefixSid() Ospfv2LsaPrefixSid
	// SetPrefixSid assigns Ospfv2LsaPrefixSid provided by user to Ospfv2OpaqueLsaExtendedPrefix.
	// Ospfv2LsaPrefixSid is the learned OSPFv2 Prefix-SID and its attributes, decoded from the Prefix-SID sub-TLV of
	// the Extended Prefix Opaque LSA (RFC 8665).
	SetPrefixSid(value Ospfv2LsaPrefixSid) Ospfv2OpaqueLsaExtendedPrefix
	// HasPrefixSid checks if PrefixSid has been set in Ospfv2OpaqueLsaExtendedPrefix
	HasPrefixSid() bool
	setNil()
}

// The IPv4 address prefix the TLV applies to, decoded from the Address Prefix field
// (RFC 7684 Section 2.1).
// Prefix returns a string
func (obj *ospfv2OpaqueLsaExtendedPrefix) Prefix() string {

	return *obj.obj.Prefix

}

// The IPv4 address prefix the TLV applies to, decoded from the Address Prefix field
// (RFC 7684 Section 2.1).
// Prefix returns a string
func (obj *ospfv2OpaqueLsaExtendedPrefix) HasPrefix() bool {
	return obj.obj.Prefix != nil
}

// The IPv4 address prefix the TLV applies to, decoded from the Address Prefix field
// (RFC 7684 Section 2.1).
// SetPrefix sets the string value in the Ospfv2OpaqueLsaExtendedPrefix object
func (obj *ospfv2OpaqueLsaExtendedPrefix) SetPrefix(value string) Ospfv2OpaqueLsaExtendedPrefix {

	obj.obj.Prefix = &value
	return obj
}

// The length in bits of prefix, decoded from the Prefix Length field
// (RFC 7684 Section 2.1).
// PrefixLength returns a uint32
func (obj *ospfv2OpaqueLsaExtendedPrefix) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// The length in bits of prefix, decoded from the Prefix Length field
// (RFC 7684 Section 2.1).
// PrefixLength returns a uint32
func (obj *ospfv2OpaqueLsaExtendedPrefix) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// The length in bits of prefix, decoded from the Prefix Length field
// (RFC 7684 Section 2.1).
// SetPrefixLength sets the uint32 value in the Ospfv2OpaqueLsaExtendedPrefix object
func (obj *ospfv2OpaqueLsaExtendedPrefix) SetPrefixLength(value uint32) Ospfv2OpaqueLsaExtendedPrefix {

	obj.obj.PrefixLength = &value
	return obj
}

// The Prefix-SID advertised for this prefix, decoded from the Prefix-SID sub-TLV,
// sub-type 2 (RFC 8665 Section 5).
// PrefixSid returns a Ospfv2LsaPrefixSid
func (obj *ospfv2OpaqueLsaExtendedPrefix) PrefixSid() Ospfv2LsaPrefixSid {
	if obj.obj.PrefixSid == nil {
		obj.obj.PrefixSid = NewOspfv2LsaPrefixSid().msg()
	}
	if obj.prefixSidHolder == nil {
		obj.prefixSidHolder = &ospfv2LsaPrefixSid{obj: obj.obj.PrefixSid}
	}
	return obj.prefixSidHolder
}

// The Prefix-SID advertised for this prefix, decoded from the Prefix-SID sub-TLV,
// sub-type 2 (RFC 8665 Section 5).
// PrefixSid returns a Ospfv2LsaPrefixSid
func (obj *ospfv2OpaqueLsaExtendedPrefix) HasPrefixSid() bool {
	return obj.obj.PrefixSid != nil
}

// The Prefix-SID advertised for this prefix, decoded from the Prefix-SID sub-TLV,
// sub-type 2 (RFC 8665 Section 5).
// SetPrefixSid sets the Ospfv2LsaPrefixSid value in the Ospfv2OpaqueLsaExtendedPrefix object
func (obj *ospfv2OpaqueLsaExtendedPrefix) SetPrefixSid(value Ospfv2LsaPrefixSid) Ospfv2OpaqueLsaExtendedPrefix {

	obj.prefixSidHolder = nil
	obj.obj.PrefixSid = value.msg()

	return obj
}

func (obj *ospfv2OpaqueLsaExtendedPrefix) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Prefix != nil {

		err := obj.validateIpv4(obj.Prefix())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv2OpaqueLsaExtendedPrefix.Prefix"))
		}

	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Ospfv2OpaqueLsaExtendedPrefix.PrefixLength <= 32 but Got %d", *obj.obj.PrefixLength))
		}

	}

	if obj.obj.PrefixSid != nil {

		obj.PrefixSid().validateObj(vObj, set_default)
	}

}

func (obj *ospfv2OpaqueLsaExtendedPrefix) setDefault() {

}
