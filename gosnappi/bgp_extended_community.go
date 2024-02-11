package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpExtendedCommunity *****
type bgpExtendedCommunity struct {
	validation
	obj                              *otg.BgpExtendedCommunity
	marshaller                       marshalBgpExtendedCommunity
	unMarshaller                     unMarshalBgpExtendedCommunity
	transitive_2OctetAsTypeHolder    BgpExtendedCommunityTransitive2OctetAsType
	transitiveIpv4AddressTypeHolder  BgpExtendedCommunityTransitiveIpv4AddressType
	transitive_4OctetAsTypeHolder    BgpExtendedCommunityTransitive4OctetAsType
	transitiveOpaqueTypeHolder       BgpExtendedCommunityTransitiveOpaqueType
	transitiveEvpnTypeHolder         BgpExtendedCommunityTransitiveEvpnType
	nonTransitive_2OctetAsTypeHolder BgpExtendedCommunityNonTransitive2OctetAsType
	customHolder                     BgpExtendedCommunityCustomType
}

func NewBgpExtendedCommunity() BgpExtendedCommunity {
	obj := bgpExtendedCommunity{obj: &otg.BgpExtendedCommunity{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpExtendedCommunity) msg() *otg.BgpExtendedCommunity {
	return obj.obj
}

func (obj *bgpExtendedCommunity) setMsg(msg *otg.BgpExtendedCommunity) BgpExtendedCommunity {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpExtendedCommunity struct {
	obj *bgpExtendedCommunity
}

type marshalBgpExtendedCommunity interface {
	// ToProto marshals BgpExtendedCommunity to protobuf object *otg.BgpExtendedCommunity
	ToProto() (*otg.BgpExtendedCommunity, error)
	// ToPbText marshals BgpExtendedCommunity to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpExtendedCommunity to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpExtendedCommunity to JSON text
	ToJson() (string, error)
}

type unMarshalbgpExtendedCommunity struct {
	obj *bgpExtendedCommunity
}

type unMarshalBgpExtendedCommunity interface {
	// FromProto unmarshals BgpExtendedCommunity from protobuf object *otg.BgpExtendedCommunity
	FromProto(msg *otg.BgpExtendedCommunity) (BgpExtendedCommunity, error)
	// FromPbText unmarshals BgpExtendedCommunity from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpExtendedCommunity from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpExtendedCommunity from JSON text
	FromJson(value string) error
}

func (obj *bgpExtendedCommunity) Marshal() marshalBgpExtendedCommunity {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpExtendedCommunity{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpExtendedCommunity) Unmarshal() unMarshalBgpExtendedCommunity {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpExtendedCommunity{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpExtendedCommunity) ToProto() (*otg.BgpExtendedCommunity, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpExtendedCommunity) FromProto(msg *otg.BgpExtendedCommunity) (BgpExtendedCommunity, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpExtendedCommunity) ToPbText() (string, error) {
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

func (m *unMarshalbgpExtendedCommunity) FromPbText(value string) error {
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

func (m *marshalbgpExtendedCommunity) ToYaml() (string, error) {
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

func (m *unMarshalbgpExtendedCommunity) FromYaml(value string) error {
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

func (m *marshalbgpExtendedCommunity) ToJson() (string, error) {
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

func (m *unMarshalbgpExtendedCommunity) FromJson(value string) error {
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

func (obj *bgpExtendedCommunity) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunity) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunity) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpExtendedCommunity) Clone() (BgpExtendedCommunity, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpExtendedCommunity()
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

func (obj *bgpExtendedCommunity) setNil() {
	obj.transitive_2OctetAsTypeHolder = nil
	obj.transitiveIpv4AddressTypeHolder = nil
	obj.transitive_4OctetAsTypeHolder = nil
	obj.transitiveOpaqueTypeHolder = nil
	obj.transitiveEvpnTypeHolder = nil
	obj.nonTransitive_2OctetAsTypeHolder = nil
	obj.customHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpExtendedCommunity is the Extended Communities Attribute is a optional BGP attribute,defined in RFC4360 with the Type Code 16.  Community and Extended Communities  attributes are utilized to trigger routing decisions, such as acceptance, rejection,  preference, or redistribution.  An extended community is an 8-Bytes value.It is divided into two main parts. The first 2 Bytes of the community  encode a type and optonal sub-type field. The last 6 bytes (or 7 bytes for types without a sub-type) carry a unique set of data in a format defined by the type and optional sub-type field.  Extended communities provide a larger  range for grouping or categorizing communities.
type BgpExtendedCommunity interface {
	Validation
	// msg marshals BgpExtendedCommunity to protobuf object *otg.BgpExtendedCommunity
	// and doesn't set defaults
	msg() *otg.BgpExtendedCommunity
	// setMsg unmarshals BgpExtendedCommunity from protobuf object *otg.BgpExtendedCommunity
	// and doesn't set defaults
	setMsg(*otg.BgpExtendedCommunity) BgpExtendedCommunity
	// provides marshal interface
	Marshal() marshalBgpExtendedCommunity
	// provides unmarshal interface
	Unmarshal() unMarshalBgpExtendedCommunity
	// validate validates BgpExtendedCommunity
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpExtendedCommunity, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns BgpExtendedCommunityChoiceEnum, set in BgpExtendedCommunity
	Choice() BgpExtendedCommunityChoiceEnum
	// setChoice assigns BgpExtendedCommunityChoiceEnum provided by user to BgpExtendedCommunity
	setChoice(value BgpExtendedCommunityChoiceEnum) BgpExtendedCommunity
	// HasChoice checks if Choice has been set in BgpExtendedCommunity
	HasChoice() bool
	// Transitive2OctetAsType returns BgpExtendedCommunityTransitive2OctetAsType, set in BgpExtendedCommunity.
	// BgpExtendedCommunityTransitive2OctetAsType is the Transitive Two-Octet AS-Specific Extended Community is sent as type 0x00 .
	Transitive2OctetAsType() BgpExtendedCommunityTransitive2OctetAsType
	// SetTransitive2OctetAsType assigns BgpExtendedCommunityTransitive2OctetAsType provided by user to BgpExtendedCommunity.
	// BgpExtendedCommunityTransitive2OctetAsType is the Transitive Two-Octet AS-Specific Extended Community is sent as type 0x00 .
	SetTransitive2OctetAsType(value BgpExtendedCommunityTransitive2OctetAsType) BgpExtendedCommunity
	// HasTransitive2OctetAsType checks if Transitive2OctetAsType has been set in BgpExtendedCommunity
	HasTransitive2OctetAsType() bool
	// TransitiveIpv4AddressType returns BgpExtendedCommunityTransitiveIpv4AddressType, set in BgpExtendedCommunity.
	// BgpExtendedCommunityTransitiveIpv4AddressType is the Transitive IPv4 Address Specific Extended Community is sent as type 0x01.
	TransitiveIpv4AddressType() BgpExtendedCommunityTransitiveIpv4AddressType
	// SetTransitiveIpv4AddressType assigns BgpExtendedCommunityTransitiveIpv4AddressType provided by user to BgpExtendedCommunity.
	// BgpExtendedCommunityTransitiveIpv4AddressType is the Transitive IPv4 Address Specific Extended Community is sent as type 0x01.
	SetTransitiveIpv4AddressType(value BgpExtendedCommunityTransitiveIpv4AddressType) BgpExtendedCommunity
	// HasTransitiveIpv4AddressType checks if TransitiveIpv4AddressType has been set in BgpExtendedCommunity
	HasTransitiveIpv4AddressType() bool
	// Transitive4OctetAsType returns BgpExtendedCommunityTransitive4OctetAsType, set in BgpExtendedCommunity.
	// BgpExtendedCommunityTransitive4OctetAsType is the Transitive Four-Octet AS-Specific Extended Community is sent as type 0x02. It is defined in RFC 5668.
	Transitive4OctetAsType() BgpExtendedCommunityTransitive4OctetAsType
	// SetTransitive4OctetAsType assigns BgpExtendedCommunityTransitive4OctetAsType provided by user to BgpExtendedCommunity.
	// BgpExtendedCommunityTransitive4OctetAsType is the Transitive Four-Octet AS-Specific Extended Community is sent as type 0x02. It is defined in RFC 5668.
	SetTransitive4OctetAsType(value BgpExtendedCommunityTransitive4OctetAsType) BgpExtendedCommunity
	// HasTransitive4OctetAsType checks if Transitive4OctetAsType has been set in BgpExtendedCommunity
	HasTransitive4OctetAsType() bool
	// TransitiveOpaqueType returns BgpExtendedCommunityTransitiveOpaqueType, set in BgpExtendedCommunity.
	// BgpExtendedCommunityTransitiveOpaqueType is the Transitive Opaque Extended Community is sent as type 0x03.
	TransitiveOpaqueType() BgpExtendedCommunityTransitiveOpaqueType
	// SetTransitiveOpaqueType assigns BgpExtendedCommunityTransitiveOpaqueType provided by user to BgpExtendedCommunity.
	// BgpExtendedCommunityTransitiveOpaqueType is the Transitive Opaque Extended Community is sent as type 0x03.
	SetTransitiveOpaqueType(value BgpExtendedCommunityTransitiveOpaqueType) BgpExtendedCommunity
	// HasTransitiveOpaqueType checks if TransitiveOpaqueType has been set in BgpExtendedCommunity
	HasTransitiveOpaqueType() bool
	// TransitiveEvpnType returns BgpExtendedCommunityTransitiveEvpnType, set in BgpExtendedCommunity.
	// BgpExtendedCommunityTransitiveEvpnType is the Transitive EVPN Extended Community is sent as type 0x06 .
	TransitiveEvpnType() BgpExtendedCommunityTransitiveEvpnType
	// SetTransitiveEvpnType assigns BgpExtendedCommunityTransitiveEvpnType provided by user to BgpExtendedCommunity.
	// BgpExtendedCommunityTransitiveEvpnType is the Transitive EVPN Extended Community is sent as type 0x06 .
	SetTransitiveEvpnType(value BgpExtendedCommunityTransitiveEvpnType) BgpExtendedCommunity
	// HasTransitiveEvpnType checks if TransitiveEvpnType has been set in BgpExtendedCommunity
	HasTransitiveEvpnType() bool
	// NonTransitive2OctetAsType returns BgpExtendedCommunityNonTransitive2OctetAsType, set in BgpExtendedCommunity.
	// BgpExtendedCommunityNonTransitive2OctetAsType is the Non-Transitive Two-Octet AS-Specific Extended Community is sent as type 0x40.
	NonTransitive2OctetAsType() BgpExtendedCommunityNonTransitive2OctetAsType
	// SetNonTransitive2OctetAsType assigns BgpExtendedCommunityNonTransitive2OctetAsType provided by user to BgpExtendedCommunity.
	// BgpExtendedCommunityNonTransitive2OctetAsType is the Non-Transitive Two-Octet AS-Specific Extended Community is sent as type 0x40.
	SetNonTransitive2OctetAsType(value BgpExtendedCommunityNonTransitive2OctetAsType) BgpExtendedCommunity
	// HasNonTransitive2OctetAsType checks if NonTransitive2OctetAsType has been set in BgpExtendedCommunity
	HasNonTransitive2OctetAsType() bool
	// Custom returns BgpExtendedCommunityCustomType, set in BgpExtendedCommunity.
	// BgpExtendedCommunityCustomType is add a custom Extended Community with a combination of types , sub-types and values not explicitly specified above or not defined yet.
	Custom() BgpExtendedCommunityCustomType
	// SetCustom assigns BgpExtendedCommunityCustomType provided by user to BgpExtendedCommunity.
	// BgpExtendedCommunityCustomType is add a custom Extended Community with a combination of types , sub-types and values not explicitly specified above or not defined yet.
	SetCustom(value BgpExtendedCommunityCustomType) BgpExtendedCommunity
	// HasCustom checks if Custom has been set in BgpExtendedCommunity
	HasCustom() bool
	setNil()
}

type BgpExtendedCommunityChoiceEnum string

// Enum of Choice on BgpExtendedCommunity
var BgpExtendedCommunityChoice = struct {
	TRANSITIVE_2OCTET_AS_TYPE     BgpExtendedCommunityChoiceEnum
	TRANSITIVE_IPV4_ADDRESS_TYPE  BgpExtendedCommunityChoiceEnum
	TRANSITIVE_4OCTET_AS_TYPE     BgpExtendedCommunityChoiceEnum
	TRANSITIVE_OPAQUE_TYPE        BgpExtendedCommunityChoiceEnum
	TRANSITIVE_EVPN_TYPE          BgpExtendedCommunityChoiceEnum
	NON_TRANSITIVE_2OCTET_AS_TYPE BgpExtendedCommunityChoiceEnum
	CUSTOM                        BgpExtendedCommunityChoiceEnum
}{
	TRANSITIVE_2OCTET_AS_TYPE:     BgpExtendedCommunityChoiceEnum("transitive_2octet_as_type"),
	TRANSITIVE_IPV4_ADDRESS_TYPE:  BgpExtendedCommunityChoiceEnum("transitive_ipv4_address_type"),
	TRANSITIVE_4OCTET_AS_TYPE:     BgpExtendedCommunityChoiceEnum("transitive_4octet_as_type"),
	TRANSITIVE_OPAQUE_TYPE:        BgpExtendedCommunityChoiceEnum("transitive_opaque_type"),
	TRANSITIVE_EVPN_TYPE:          BgpExtendedCommunityChoiceEnum("transitive_evpn_type"),
	NON_TRANSITIVE_2OCTET_AS_TYPE: BgpExtendedCommunityChoiceEnum("non_transitive_2octet_as_type"),
	CUSTOM:                        BgpExtendedCommunityChoiceEnum("custom"),
}

func (obj *bgpExtendedCommunity) Choice() BgpExtendedCommunityChoiceEnum {
	return BgpExtendedCommunityChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *bgpExtendedCommunity) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *bgpExtendedCommunity) setChoice(value BgpExtendedCommunityChoiceEnum) BgpExtendedCommunity {
	intValue, ok := otg.BgpExtendedCommunity_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpExtendedCommunityChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpExtendedCommunity_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Custom = nil
	obj.customHolder = nil
	obj.obj.NonTransitive_2OctetAsType = nil
	obj.nonTransitive_2OctetAsTypeHolder = nil
	obj.obj.TransitiveEvpnType = nil
	obj.transitiveEvpnTypeHolder = nil
	obj.obj.TransitiveOpaqueType = nil
	obj.transitiveOpaqueTypeHolder = nil
	obj.obj.Transitive_4OctetAsType = nil
	obj.transitive_4OctetAsTypeHolder = nil
	obj.obj.TransitiveIpv4AddressType = nil
	obj.transitiveIpv4AddressTypeHolder = nil
	obj.obj.Transitive_2OctetAsType = nil
	obj.transitive_2OctetAsTypeHolder = nil

	if value == BgpExtendedCommunityChoice.TRANSITIVE_2OCTET_AS_TYPE {
		obj.obj.Transitive_2OctetAsType = NewBgpExtendedCommunityTransitive2OctetAsType().msg()
	}

	if value == BgpExtendedCommunityChoice.TRANSITIVE_IPV4_ADDRESS_TYPE {
		obj.obj.TransitiveIpv4AddressType = NewBgpExtendedCommunityTransitiveIpv4AddressType().msg()
	}

	if value == BgpExtendedCommunityChoice.TRANSITIVE_4OCTET_AS_TYPE {
		obj.obj.Transitive_4OctetAsType = NewBgpExtendedCommunityTransitive4OctetAsType().msg()
	}

	if value == BgpExtendedCommunityChoice.TRANSITIVE_OPAQUE_TYPE {
		obj.obj.TransitiveOpaqueType = NewBgpExtendedCommunityTransitiveOpaqueType().msg()
	}

	if value == BgpExtendedCommunityChoice.TRANSITIVE_EVPN_TYPE {
		obj.obj.TransitiveEvpnType = NewBgpExtendedCommunityTransitiveEvpnType().msg()
	}

	if value == BgpExtendedCommunityChoice.NON_TRANSITIVE_2OCTET_AS_TYPE {
		obj.obj.NonTransitive_2OctetAsType = NewBgpExtendedCommunityNonTransitive2OctetAsType().msg()
	}

	if value == BgpExtendedCommunityChoice.CUSTOM {
		obj.obj.Custom = NewBgpExtendedCommunityCustomType().msg()
	}

	return obj
}

// description is TBD
// Transitive2OctetAsType returns a BgpExtendedCommunityTransitive2OctetAsType
func (obj *bgpExtendedCommunity) Transitive2OctetAsType() BgpExtendedCommunityTransitive2OctetAsType {
	if obj.obj.Transitive_2OctetAsType == nil {
		obj.setChoice(BgpExtendedCommunityChoice.TRANSITIVE_2OCTET_AS_TYPE)
	}
	if obj.transitive_2OctetAsTypeHolder == nil {
		obj.transitive_2OctetAsTypeHolder = &bgpExtendedCommunityTransitive2OctetAsType{obj: obj.obj.Transitive_2OctetAsType}
	}
	return obj.transitive_2OctetAsTypeHolder
}

// description is TBD
// Transitive2OctetAsType returns a BgpExtendedCommunityTransitive2OctetAsType
func (obj *bgpExtendedCommunity) HasTransitive2OctetAsType() bool {
	return obj.obj.Transitive_2OctetAsType != nil
}

// description is TBD
// SetTransitive2OctetAsType sets the BgpExtendedCommunityTransitive2OctetAsType value in the BgpExtendedCommunity object
func (obj *bgpExtendedCommunity) SetTransitive2OctetAsType(value BgpExtendedCommunityTransitive2OctetAsType) BgpExtendedCommunity {
	obj.setChoice(BgpExtendedCommunityChoice.TRANSITIVE_2OCTET_AS_TYPE)
	obj.transitive_2OctetAsTypeHolder = nil
	obj.obj.Transitive_2OctetAsType = value.msg()

	return obj
}

// description is TBD
// TransitiveIpv4AddressType returns a BgpExtendedCommunityTransitiveIpv4AddressType
func (obj *bgpExtendedCommunity) TransitiveIpv4AddressType() BgpExtendedCommunityTransitiveIpv4AddressType {
	if obj.obj.TransitiveIpv4AddressType == nil {
		obj.setChoice(BgpExtendedCommunityChoice.TRANSITIVE_IPV4_ADDRESS_TYPE)
	}
	if obj.transitiveIpv4AddressTypeHolder == nil {
		obj.transitiveIpv4AddressTypeHolder = &bgpExtendedCommunityTransitiveIpv4AddressType{obj: obj.obj.TransitiveIpv4AddressType}
	}
	return obj.transitiveIpv4AddressTypeHolder
}

// description is TBD
// TransitiveIpv4AddressType returns a BgpExtendedCommunityTransitiveIpv4AddressType
func (obj *bgpExtendedCommunity) HasTransitiveIpv4AddressType() bool {
	return obj.obj.TransitiveIpv4AddressType != nil
}

// description is TBD
// SetTransitiveIpv4AddressType sets the BgpExtendedCommunityTransitiveIpv4AddressType value in the BgpExtendedCommunity object
func (obj *bgpExtendedCommunity) SetTransitiveIpv4AddressType(value BgpExtendedCommunityTransitiveIpv4AddressType) BgpExtendedCommunity {
	obj.setChoice(BgpExtendedCommunityChoice.TRANSITIVE_IPV4_ADDRESS_TYPE)
	obj.transitiveIpv4AddressTypeHolder = nil
	obj.obj.TransitiveIpv4AddressType = value.msg()

	return obj
}

// description is TBD
// Transitive4OctetAsType returns a BgpExtendedCommunityTransitive4OctetAsType
func (obj *bgpExtendedCommunity) Transitive4OctetAsType() BgpExtendedCommunityTransitive4OctetAsType {
	if obj.obj.Transitive_4OctetAsType == nil {
		obj.setChoice(BgpExtendedCommunityChoice.TRANSITIVE_4OCTET_AS_TYPE)
	}
	if obj.transitive_4OctetAsTypeHolder == nil {
		obj.transitive_4OctetAsTypeHolder = &bgpExtendedCommunityTransitive4OctetAsType{obj: obj.obj.Transitive_4OctetAsType}
	}
	return obj.transitive_4OctetAsTypeHolder
}

// description is TBD
// Transitive4OctetAsType returns a BgpExtendedCommunityTransitive4OctetAsType
func (obj *bgpExtendedCommunity) HasTransitive4OctetAsType() bool {
	return obj.obj.Transitive_4OctetAsType != nil
}

// description is TBD
// SetTransitive4OctetAsType sets the BgpExtendedCommunityTransitive4OctetAsType value in the BgpExtendedCommunity object
func (obj *bgpExtendedCommunity) SetTransitive4OctetAsType(value BgpExtendedCommunityTransitive4OctetAsType) BgpExtendedCommunity {
	obj.setChoice(BgpExtendedCommunityChoice.TRANSITIVE_4OCTET_AS_TYPE)
	obj.transitive_4OctetAsTypeHolder = nil
	obj.obj.Transitive_4OctetAsType = value.msg()

	return obj
}

// description is TBD
// TransitiveOpaqueType returns a BgpExtendedCommunityTransitiveOpaqueType
func (obj *bgpExtendedCommunity) TransitiveOpaqueType() BgpExtendedCommunityTransitiveOpaqueType {
	if obj.obj.TransitiveOpaqueType == nil {
		obj.setChoice(BgpExtendedCommunityChoice.TRANSITIVE_OPAQUE_TYPE)
	}
	if obj.transitiveOpaqueTypeHolder == nil {
		obj.transitiveOpaqueTypeHolder = &bgpExtendedCommunityTransitiveOpaqueType{obj: obj.obj.TransitiveOpaqueType}
	}
	return obj.transitiveOpaqueTypeHolder
}

// description is TBD
// TransitiveOpaqueType returns a BgpExtendedCommunityTransitiveOpaqueType
func (obj *bgpExtendedCommunity) HasTransitiveOpaqueType() bool {
	return obj.obj.TransitiveOpaqueType != nil
}

// description is TBD
// SetTransitiveOpaqueType sets the BgpExtendedCommunityTransitiveOpaqueType value in the BgpExtendedCommunity object
func (obj *bgpExtendedCommunity) SetTransitiveOpaqueType(value BgpExtendedCommunityTransitiveOpaqueType) BgpExtendedCommunity {
	obj.setChoice(BgpExtendedCommunityChoice.TRANSITIVE_OPAQUE_TYPE)
	obj.transitiveOpaqueTypeHolder = nil
	obj.obj.TransitiveOpaqueType = value.msg()

	return obj
}

// description is TBD
// TransitiveEvpnType returns a BgpExtendedCommunityTransitiveEvpnType
func (obj *bgpExtendedCommunity) TransitiveEvpnType() BgpExtendedCommunityTransitiveEvpnType {
	if obj.obj.TransitiveEvpnType == nil {
		obj.setChoice(BgpExtendedCommunityChoice.TRANSITIVE_EVPN_TYPE)
	}
	if obj.transitiveEvpnTypeHolder == nil {
		obj.transitiveEvpnTypeHolder = &bgpExtendedCommunityTransitiveEvpnType{obj: obj.obj.TransitiveEvpnType}
	}
	return obj.transitiveEvpnTypeHolder
}

// description is TBD
// TransitiveEvpnType returns a BgpExtendedCommunityTransitiveEvpnType
func (obj *bgpExtendedCommunity) HasTransitiveEvpnType() bool {
	return obj.obj.TransitiveEvpnType != nil
}

// description is TBD
// SetTransitiveEvpnType sets the BgpExtendedCommunityTransitiveEvpnType value in the BgpExtendedCommunity object
func (obj *bgpExtendedCommunity) SetTransitiveEvpnType(value BgpExtendedCommunityTransitiveEvpnType) BgpExtendedCommunity {
	obj.setChoice(BgpExtendedCommunityChoice.TRANSITIVE_EVPN_TYPE)
	obj.transitiveEvpnTypeHolder = nil
	obj.obj.TransitiveEvpnType = value.msg()

	return obj
}

// description is TBD
// NonTransitive2OctetAsType returns a BgpExtendedCommunityNonTransitive2OctetAsType
func (obj *bgpExtendedCommunity) NonTransitive2OctetAsType() BgpExtendedCommunityNonTransitive2OctetAsType {
	if obj.obj.NonTransitive_2OctetAsType == nil {
		obj.setChoice(BgpExtendedCommunityChoice.NON_TRANSITIVE_2OCTET_AS_TYPE)
	}
	if obj.nonTransitive_2OctetAsTypeHolder == nil {
		obj.nonTransitive_2OctetAsTypeHolder = &bgpExtendedCommunityNonTransitive2OctetAsType{obj: obj.obj.NonTransitive_2OctetAsType}
	}
	return obj.nonTransitive_2OctetAsTypeHolder
}

// description is TBD
// NonTransitive2OctetAsType returns a BgpExtendedCommunityNonTransitive2OctetAsType
func (obj *bgpExtendedCommunity) HasNonTransitive2OctetAsType() bool {
	return obj.obj.NonTransitive_2OctetAsType != nil
}

// description is TBD
// SetNonTransitive2OctetAsType sets the BgpExtendedCommunityNonTransitive2OctetAsType value in the BgpExtendedCommunity object
func (obj *bgpExtendedCommunity) SetNonTransitive2OctetAsType(value BgpExtendedCommunityNonTransitive2OctetAsType) BgpExtendedCommunity {
	obj.setChoice(BgpExtendedCommunityChoice.NON_TRANSITIVE_2OCTET_AS_TYPE)
	obj.nonTransitive_2OctetAsTypeHolder = nil
	obj.obj.NonTransitive_2OctetAsType = value.msg()

	return obj
}

// description is TBD
// Custom returns a BgpExtendedCommunityCustomType
func (obj *bgpExtendedCommunity) Custom() BgpExtendedCommunityCustomType {
	if obj.obj.Custom == nil {
		obj.setChoice(BgpExtendedCommunityChoice.CUSTOM)
	}
	if obj.customHolder == nil {
		obj.customHolder = &bgpExtendedCommunityCustomType{obj: obj.obj.Custom}
	}
	return obj.customHolder
}

// description is TBD
// Custom returns a BgpExtendedCommunityCustomType
func (obj *bgpExtendedCommunity) HasCustom() bool {
	return obj.obj.Custom != nil
}

// description is TBD
// SetCustom sets the BgpExtendedCommunityCustomType value in the BgpExtendedCommunity object
func (obj *bgpExtendedCommunity) SetCustom(value BgpExtendedCommunityCustomType) BgpExtendedCommunity {
	obj.setChoice(BgpExtendedCommunityChoice.CUSTOM)
	obj.customHolder = nil
	obj.obj.Custom = value.msg()

	return obj
}

func (obj *bgpExtendedCommunity) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Transitive_2OctetAsType != nil {

		obj.Transitive2OctetAsType().validateObj(vObj, set_default)
	}

	if obj.obj.TransitiveIpv4AddressType != nil {

		obj.TransitiveIpv4AddressType().validateObj(vObj, set_default)
	}

	if obj.obj.Transitive_4OctetAsType != nil {

		obj.Transitive4OctetAsType().validateObj(vObj, set_default)
	}

	if obj.obj.TransitiveOpaqueType != nil {

		obj.TransitiveOpaqueType().validateObj(vObj, set_default)
	}

	if obj.obj.TransitiveEvpnType != nil {

		obj.TransitiveEvpnType().validateObj(vObj, set_default)
	}

	if obj.obj.NonTransitive_2OctetAsType != nil {

		obj.NonTransitive2OctetAsType().validateObj(vObj, set_default)
	}

	if obj.obj.Custom != nil {

		obj.Custom().validateObj(vObj, set_default)
	}

}

func (obj *bgpExtendedCommunity) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(BgpExtendedCommunityChoice.TRANSITIVE_2OCTET_AS_TYPE)

	}

}
