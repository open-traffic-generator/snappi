package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpRouteDistinguisher *****
type bgpRouteDistinguisher struct {
	validation
	obj          *otg.BgpRouteDistinguisher
	marshaller   marshalBgpRouteDistinguisher
	unMarshaller unMarshalBgpRouteDistinguisher
}

func NewBgpRouteDistinguisher() BgpRouteDistinguisher {
	obj := bgpRouteDistinguisher{obj: &otg.BgpRouteDistinguisher{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpRouteDistinguisher) msg() *otg.BgpRouteDistinguisher {
	return obj.obj
}

func (obj *bgpRouteDistinguisher) setMsg(msg *otg.BgpRouteDistinguisher) BgpRouteDistinguisher {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpRouteDistinguisher struct {
	obj *bgpRouteDistinguisher
}

type marshalBgpRouteDistinguisher interface {
	// ToProto marshals BgpRouteDistinguisher to protobuf object *otg.BgpRouteDistinguisher
	ToProto() (*otg.BgpRouteDistinguisher, error)
	// ToPbText marshals BgpRouteDistinguisher to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpRouteDistinguisher to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpRouteDistinguisher to JSON text
	ToJson() (string, error)
}

type unMarshalbgpRouteDistinguisher struct {
	obj *bgpRouteDistinguisher
}

type unMarshalBgpRouteDistinguisher interface {
	// FromProto unmarshals BgpRouteDistinguisher from protobuf object *otg.BgpRouteDistinguisher
	FromProto(msg *otg.BgpRouteDistinguisher) (BgpRouteDistinguisher, error)
	// FromPbText unmarshals BgpRouteDistinguisher from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpRouteDistinguisher from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpRouteDistinguisher from JSON text
	FromJson(value string) error
}

func (obj *bgpRouteDistinguisher) Marshal() marshalBgpRouteDistinguisher {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpRouteDistinguisher{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpRouteDistinguisher) Unmarshal() unMarshalBgpRouteDistinguisher {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpRouteDistinguisher{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpRouteDistinguisher) ToProto() (*otg.BgpRouteDistinguisher, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpRouteDistinguisher) FromProto(msg *otg.BgpRouteDistinguisher) (BgpRouteDistinguisher, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpRouteDistinguisher) ToPbText() (string, error) {
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

func (m *unMarshalbgpRouteDistinguisher) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalbgpRouteDistinguisher) ToYaml() (string, error) {
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

func (m *unMarshalbgpRouteDistinguisher) FromYaml(value string) error {
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

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalbgpRouteDistinguisher) ToJson() (string, error) {
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

func (m *unMarshalbgpRouteDistinguisher) FromJson(value string) error {
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

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *bgpRouteDistinguisher) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpRouteDistinguisher) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpRouteDistinguisher) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpRouteDistinguisher) Clone() (BgpRouteDistinguisher, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpRouteDistinguisher()
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

// BgpRouteDistinguisher is bGP Route Distinguisher.
type BgpRouteDistinguisher interface {
	Validation
	// msg marshals BgpRouteDistinguisher to protobuf object *otg.BgpRouteDistinguisher
	// and doesn't set defaults
	msg() *otg.BgpRouteDistinguisher
	// setMsg unmarshals BgpRouteDistinguisher from protobuf object *otg.BgpRouteDistinguisher
	// and doesn't set defaults
	setMsg(*otg.BgpRouteDistinguisher) BgpRouteDistinguisher
	// provides marshal interface
	Marshal() marshalBgpRouteDistinguisher
	// provides unmarshal interface
	Unmarshal() unMarshalBgpRouteDistinguisher
	// validate validates BgpRouteDistinguisher
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpRouteDistinguisher, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RdType returns BgpRouteDistinguisherRdTypeEnum, set in BgpRouteDistinguisher
	RdType() BgpRouteDistinguisherRdTypeEnum
	// SetRdType assigns BgpRouteDistinguisherRdTypeEnum provided by user to BgpRouteDistinguisher
	SetRdType(value BgpRouteDistinguisherRdTypeEnum) BgpRouteDistinguisher
	// HasRdType checks if RdType has been set in BgpRouteDistinguisher
	HasRdType() bool
	// AutoConfigRdIpAddr returns bool, set in BgpRouteDistinguisher.
	AutoConfigRdIpAddr() bool
	// SetAutoConfigRdIpAddr assigns bool provided by user to BgpRouteDistinguisher
	SetAutoConfigRdIpAddr(value bool) BgpRouteDistinguisher
	// HasAutoConfigRdIpAddr checks if AutoConfigRdIpAddr has been set in BgpRouteDistinguisher
	HasAutoConfigRdIpAddr() bool
	// RdValue returns string, set in BgpRouteDistinguisher.
	RdValue() string
	// SetRdValue assigns string provided by user to BgpRouteDistinguisher
	SetRdValue(value string) BgpRouteDistinguisher
	// HasRdValue checks if RdValue has been set in BgpRouteDistinguisher
	HasRdValue() bool
}

type BgpRouteDistinguisherRdTypeEnum string

// Enum of RdType on BgpRouteDistinguisher
var BgpRouteDistinguisherRdType = struct {
	AS_2OCTET    BgpRouteDistinguisherRdTypeEnum
	IPV4_ADDRESS BgpRouteDistinguisherRdTypeEnum
	AS_4OCTET    BgpRouteDistinguisherRdTypeEnum
}{
	AS_2OCTET:    BgpRouteDistinguisherRdTypeEnum("as_2octet"),
	IPV4_ADDRESS: BgpRouteDistinguisherRdTypeEnum("ipv4_address"),
	AS_4OCTET:    BgpRouteDistinguisherRdTypeEnum("as_4octet"),
}

func (obj *bgpRouteDistinguisher) RdType() BgpRouteDistinguisherRdTypeEnum {
	return BgpRouteDistinguisherRdTypeEnum(obj.obj.RdType.Enum().String())
}

// Route Distinguisher Type field of 2 Byte.
// - as_2octet: Two-Octet AS Specific Extended Community (RFC 4360).
// - ipv4_address: IPv4 Address Specific Extended Community (RFC 4360).
// - as_4octet:  4-Octet AS Specific Extended Community (RFC 5668).
// RdType returns a string
func (obj *bgpRouteDistinguisher) HasRdType() bool {
	return obj.obj.RdType != nil
}

func (obj *bgpRouteDistinguisher) SetRdType(value BgpRouteDistinguisherRdTypeEnum) BgpRouteDistinguisher {
	intValue, ok := otg.BgpRouteDistinguisher_RdType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpRouteDistinguisherRdTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpRouteDistinguisher_RdType_Enum(intValue)
	obj.obj.RdType = &enumValue

	return obj
}

// Allow to automatically configure RD IP address  from local ip.
// AutoConfigRdIpAddr returns a bool
func (obj *bgpRouteDistinguisher) AutoConfigRdIpAddr() bool {

	return *obj.obj.AutoConfigRdIpAddr

}

// Allow to automatically configure RD IP address  from local ip.
// AutoConfigRdIpAddr returns a bool
func (obj *bgpRouteDistinguisher) HasAutoConfigRdIpAddr() bool {
	return obj.obj.AutoConfigRdIpAddr != nil
}

// Allow to automatically configure RD IP address  from local ip.
// SetAutoConfigRdIpAddr sets the bool value in the BgpRouteDistinguisher object
func (obj *bgpRouteDistinguisher) SetAutoConfigRdIpAddr(value bool) BgpRouteDistinguisher {

	obj.obj.AutoConfigRdIpAddr = &value
	return obj
}

// Colon separated Extended Community value of 6 Bytes - "AS number: Value".  Example - for the as_2octet or as_4octet "60005:100",  for ipv4_address "1.1.1.1:100"
// RdValue returns a string
func (obj *bgpRouteDistinguisher) RdValue() string {

	return *obj.obj.RdValue

}

// Colon separated Extended Community value of 6 Bytes - "AS number: Value".  Example - for the as_2octet or as_4octet "60005:100",  for ipv4_address "1.1.1.1:100"
// RdValue returns a string
func (obj *bgpRouteDistinguisher) HasRdValue() bool {
	return obj.obj.RdValue != nil
}

// Colon separated Extended Community value of 6 Bytes - "AS number: Value".  Example - for the as_2octet or as_4octet "60005:100",  for ipv4_address "1.1.1.1:100"
// SetRdValue sets the string value in the BgpRouteDistinguisher object
func (obj *bgpRouteDistinguisher) SetRdValue(value string) BgpRouteDistinguisher {

	obj.obj.RdValue = &value
	return obj
}

func (obj *bgpRouteDistinguisher) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *bgpRouteDistinguisher) setDefault() {
	if obj.obj.RdType == nil {
		obj.SetRdType(BgpRouteDistinguisherRdType.AS_2OCTET)

	}
	if obj.obj.AutoConfigRdIpAddr == nil {
		obj.SetAutoConfigRdIpAddr(false)
	}

}
