package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspSRv6AdvertisedLocatorAsPrefix *****
type isisLspSRv6AdvertisedLocatorAsPrefix struct {
	validation
	obj                    *otg.IsisLspSRv6AdvertisedLocatorAsPrefix
	marshaller             marshalIsisLspSRv6AdvertisedLocatorAsPrefix
	unMarshaller           unMarshalIsisLspSRv6AdvertisedLocatorAsPrefix
	prefixAttributesHolder IsisLspSRv6PrefixAttributes
}

func NewIsisLspSRv6AdvertisedLocatorAsPrefix() IsisLspSRv6AdvertisedLocatorAsPrefix {
	obj := isisLspSRv6AdvertisedLocatorAsPrefix{obj: &otg.IsisLspSRv6AdvertisedLocatorAsPrefix{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspSRv6AdvertisedLocatorAsPrefix) msg() *otg.IsisLspSRv6AdvertisedLocatorAsPrefix {
	return obj.obj
}

func (obj *isisLspSRv6AdvertisedLocatorAsPrefix) setMsg(msg *otg.IsisLspSRv6AdvertisedLocatorAsPrefix) IsisLspSRv6AdvertisedLocatorAsPrefix {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspSRv6AdvertisedLocatorAsPrefix struct {
	obj *isisLspSRv6AdvertisedLocatorAsPrefix
}

type marshalIsisLspSRv6AdvertisedLocatorAsPrefix interface {
	// ToProto marshals IsisLspSRv6AdvertisedLocatorAsPrefix to protobuf object *otg.IsisLspSRv6AdvertisedLocatorAsPrefix
	ToProto() (*otg.IsisLspSRv6AdvertisedLocatorAsPrefix, error)
	// ToPbText marshals IsisLspSRv6AdvertisedLocatorAsPrefix to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspSRv6AdvertisedLocatorAsPrefix to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspSRv6AdvertisedLocatorAsPrefix to JSON text
	ToJson() (string, error)
}

type unMarshalisisLspSRv6AdvertisedLocatorAsPrefix struct {
	obj *isisLspSRv6AdvertisedLocatorAsPrefix
}

type unMarshalIsisLspSRv6AdvertisedLocatorAsPrefix interface {
	// FromProto unmarshals IsisLspSRv6AdvertisedLocatorAsPrefix from protobuf object *otg.IsisLspSRv6AdvertisedLocatorAsPrefix
	FromProto(msg *otg.IsisLspSRv6AdvertisedLocatorAsPrefix) (IsisLspSRv6AdvertisedLocatorAsPrefix, error)
	// FromPbText unmarshals IsisLspSRv6AdvertisedLocatorAsPrefix from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspSRv6AdvertisedLocatorAsPrefix from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspSRv6AdvertisedLocatorAsPrefix from JSON text
	FromJson(value string) error
}

func (obj *isisLspSRv6AdvertisedLocatorAsPrefix) Marshal() marshalIsisLspSRv6AdvertisedLocatorAsPrefix {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspSRv6AdvertisedLocatorAsPrefix{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspSRv6AdvertisedLocatorAsPrefix) Unmarshal() unMarshalIsisLspSRv6AdvertisedLocatorAsPrefix {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspSRv6AdvertisedLocatorAsPrefix{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspSRv6AdvertisedLocatorAsPrefix) ToProto() (*otg.IsisLspSRv6AdvertisedLocatorAsPrefix, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspSRv6AdvertisedLocatorAsPrefix) FromProto(msg *otg.IsisLspSRv6AdvertisedLocatorAsPrefix) (IsisLspSRv6AdvertisedLocatorAsPrefix, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspSRv6AdvertisedLocatorAsPrefix) ToPbText() (string, error) {
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

func (m *unMarshalisisLspSRv6AdvertisedLocatorAsPrefix) FromPbText(value string) error {
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

func (m *marshalisisLspSRv6AdvertisedLocatorAsPrefix) ToYaml() (string, error) {
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

func (m *unMarshalisisLspSRv6AdvertisedLocatorAsPrefix) FromYaml(value string) error {
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

func (m *marshalisisLspSRv6AdvertisedLocatorAsPrefix) ToJson() (string, error) {
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

func (m *unMarshalisisLspSRv6AdvertisedLocatorAsPrefix) FromJson(value string) error {
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

func (obj *isisLspSRv6AdvertisedLocatorAsPrefix) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspSRv6AdvertisedLocatorAsPrefix) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspSRv6AdvertisedLocatorAsPrefix) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspSRv6AdvertisedLocatorAsPrefix) Clone() (IsisLspSRv6AdvertisedLocatorAsPrefix, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspSRv6AdvertisedLocatorAsPrefix()
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

func (obj *isisLspSRv6AdvertisedLocatorAsPrefix) setNil() {
	obj.prefixAttributesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisLspSRv6AdvertisedLocatorAsPrefix is learned state for the secondary IPv6 Reachability prefix advertisement (TLV 236/237) of an SRv6 locator. Present only when the originating router advertised the locator prefix alongside the SRv6 Locator TLV. Reference: RFC 9352 Section 7.1.
type IsisLspSRv6AdvertisedLocatorAsPrefix interface {
	Validation
	// msg marshals IsisLspSRv6AdvertisedLocatorAsPrefix to protobuf object *otg.IsisLspSRv6AdvertisedLocatorAsPrefix
	// and doesn't set defaults
	msg() *otg.IsisLspSRv6AdvertisedLocatorAsPrefix
	// setMsg unmarshals IsisLspSRv6AdvertisedLocatorAsPrefix from protobuf object *otg.IsisLspSRv6AdvertisedLocatorAsPrefix
	// and doesn't set defaults
	setMsg(*otg.IsisLspSRv6AdvertisedLocatorAsPrefix) IsisLspSRv6AdvertisedLocatorAsPrefix
	// provides marshal interface
	Marshal() marshalIsisLspSRv6AdvertisedLocatorAsPrefix
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspSRv6AdvertisedLocatorAsPrefix
	// validate validates IsisLspSRv6AdvertisedLocatorAsPrefix
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspSRv6AdvertisedLocatorAsPrefix, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PrefixAttributes returns IsisLspSRv6PrefixAttributes, set in IsisLspSRv6AdvertisedLocatorAsPrefix.
	// IsisLspSRv6PrefixAttributes is prefix Attribute Flags Sub-TLV (sub-TLV type 4, RFC 7794) learned from the IPv6 Reachability advertisement of an SRv6 locator prefix. Reference: RFC 7794.
	PrefixAttributes() IsisLspSRv6PrefixAttributes
	// SetPrefixAttributes assigns IsisLspSRv6PrefixAttributes provided by user to IsisLspSRv6AdvertisedLocatorAsPrefix.
	// IsisLspSRv6PrefixAttributes is prefix Attribute Flags Sub-TLV (sub-TLV type 4, RFC 7794) learned from the IPv6 Reachability advertisement of an SRv6 locator prefix. Reference: RFC 7794.
	SetPrefixAttributes(value IsisLspSRv6PrefixAttributes) IsisLspSRv6AdvertisedLocatorAsPrefix
	// HasPrefixAttributes checks if PrefixAttributes has been set in IsisLspSRv6AdvertisedLocatorAsPrefix
	HasPrefixAttributes() bool
	setNil()
}

// Prefix Attribute Flags Sub-TLV (sub-TLV type 4, RFC 7794) received in the locator's IPv6 Reachability advertisement. Present only if the originating router included this sub-TLV.
// PrefixAttributes returns a IsisLspSRv6PrefixAttributes
func (obj *isisLspSRv6AdvertisedLocatorAsPrefix) PrefixAttributes() IsisLspSRv6PrefixAttributes {
	if obj.obj.PrefixAttributes == nil {
		obj.obj.PrefixAttributes = NewIsisLspSRv6PrefixAttributes().msg()
	}
	if obj.prefixAttributesHolder == nil {
		obj.prefixAttributesHolder = &isisLspSRv6PrefixAttributes{obj: obj.obj.PrefixAttributes}
	}
	return obj.prefixAttributesHolder
}

// Prefix Attribute Flags Sub-TLV (sub-TLV type 4, RFC 7794) received in the locator's IPv6 Reachability advertisement. Present only if the originating router included this sub-TLV.
// PrefixAttributes returns a IsisLspSRv6PrefixAttributes
func (obj *isisLspSRv6AdvertisedLocatorAsPrefix) HasPrefixAttributes() bool {
	return obj.obj.PrefixAttributes != nil
}

// Prefix Attribute Flags Sub-TLV (sub-TLV type 4, RFC 7794) received in the locator's IPv6 Reachability advertisement. Present only if the originating router included this sub-TLV.
// SetPrefixAttributes sets the IsisLspSRv6PrefixAttributes value in the IsisLspSRv6AdvertisedLocatorAsPrefix object
func (obj *isisLspSRv6AdvertisedLocatorAsPrefix) SetPrefixAttributes(value IsisLspSRv6PrefixAttributes) IsisLspSRv6AdvertisedLocatorAsPrefix {

	obj.prefixAttributesHolder = nil
	obj.obj.PrefixAttributes = value.msg()

	return obj
}

func (obj *isisLspSRv6AdvertisedLocatorAsPrefix) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PrefixAttributes != nil {

		obj.PrefixAttributes().validateObj(vObj, set_default)
	}

}

func (obj *isisLspSRv6AdvertisedLocatorAsPrefix) setDefault() {

}
