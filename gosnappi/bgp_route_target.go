package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpRouteTarget *****
type bgpRouteTarget struct {
	validation
	obj          *otg.BgpRouteTarget
	marshaller   marshalBgpRouteTarget
	unMarshaller unMarshalBgpRouteTarget
}

func NewBgpRouteTarget() BgpRouteTarget {
	obj := bgpRouteTarget{obj: &otg.BgpRouteTarget{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpRouteTarget) msg() *otg.BgpRouteTarget {
	return obj.obj
}

func (obj *bgpRouteTarget) setMsg(msg *otg.BgpRouteTarget) BgpRouteTarget {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpRouteTarget struct {
	obj *bgpRouteTarget
}

type marshalBgpRouteTarget interface {
	// ToProto marshals BgpRouteTarget to protobuf object *otg.BgpRouteTarget
	ToProto() (*otg.BgpRouteTarget, error)
	// ToPbText marshals BgpRouteTarget to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpRouteTarget to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpRouteTarget to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpRouteTarget to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpRouteTarget struct {
	obj *bgpRouteTarget
}

type unMarshalBgpRouteTarget interface {
	// FromProto unmarshals BgpRouteTarget from protobuf object *otg.BgpRouteTarget
	FromProto(msg *otg.BgpRouteTarget) (BgpRouteTarget, error)
	// FromPbText unmarshals BgpRouteTarget from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpRouteTarget from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpRouteTarget from JSON text
	FromJson(value string) error
}

func (obj *bgpRouteTarget) Marshal() marshalBgpRouteTarget {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpRouteTarget{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpRouteTarget) Unmarshal() unMarshalBgpRouteTarget {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpRouteTarget{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpRouteTarget) ToProto() (*otg.BgpRouteTarget, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpRouteTarget) FromProto(msg *otg.BgpRouteTarget) (BgpRouteTarget, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpRouteTarget) ToPbText() (string, error) {
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

func (m *unMarshalbgpRouteTarget) FromPbText(value string) error {
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

func (m *marshalbgpRouteTarget) ToYaml() (string, error) {
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

func (m *unMarshalbgpRouteTarget) FromYaml(value string) error {
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

func (m *marshalbgpRouteTarget) ToJsonRaw() (string, error) {
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

func (m *marshalbgpRouteTarget) ToJson() (string, error) {
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

func (m *unMarshalbgpRouteTarget) FromJson(value string) error {
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

func (obj *bgpRouteTarget) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpRouteTarget) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpRouteTarget) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpRouteTarget) Clone() (BgpRouteTarget, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpRouteTarget()
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

// BgpRouteTarget is bGP Route Target.
type BgpRouteTarget interface {
	Validation
	// msg marshals BgpRouteTarget to protobuf object *otg.BgpRouteTarget
	// and doesn't set defaults
	msg() *otg.BgpRouteTarget
	// setMsg unmarshals BgpRouteTarget from protobuf object *otg.BgpRouteTarget
	// and doesn't set defaults
	setMsg(*otg.BgpRouteTarget) BgpRouteTarget
	// provides marshal interface
	Marshal() marshalBgpRouteTarget
	// provides unmarshal interface
	Unmarshal() unMarshalBgpRouteTarget
	// validate validates BgpRouteTarget
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpRouteTarget, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RtType returns BgpRouteTargetRtTypeEnum, set in BgpRouteTarget
	RtType() BgpRouteTargetRtTypeEnum
	// SetRtType assigns BgpRouteTargetRtTypeEnum provided by user to BgpRouteTarget
	SetRtType(value BgpRouteTargetRtTypeEnum) BgpRouteTarget
	// HasRtType checks if RtType has been set in BgpRouteTarget
	HasRtType() bool
	// RtValue returns string, set in BgpRouteTarget.
	RtValue() string
	// SetRtValue assigns string provided by user to BgpRouteTarget
	SetRtValue(value string) BgpRouteTarget
	// HasRtValue checks if RtValue has been set in BgpRouteTarget
	HasRtValue() bool
}

type BgpRouteTargetRtTypeEnum string

// Enum of RtType on BgpRouteTarget
var BgpRouteTargetRtType = struct {
	AS_2OCTET    BgpRouteTargetRtTypeEnum
	IPV4_ADDRESS BgpRouteTargetRtTypeEnum
	AS_4OCTET    BgpRouteTargetRtTypeEnum
}{
	AS_2OCTET:    BgpRouteTargetRtTypeEnum("as_2octet"),
	IPV4_ADDRESS: BgpRouteTargetRtTypeEnum("ipv4_address"),
	AS_4OCTET:    BgpRouteTargetRtTypeEnum("as_4octet"),
}

func (obj *bgpRouteTarget) RtType() BgpRouteTargetRtTypeEnum {
	return BgpRouteTargetRtTypeEnum(obj.obj.RtType.Enum().String())
}

// Extended Community Type field of 2 Byte.
// - as_2octet: Two-Octet AS Specific Extended Community (RFC 4360).
// - ipv4_address: IPv4 Address Specific Extended Community (RFC 4360).
// - as_4octet:  4-Octet AS Specific Extended Community (RFC 5668).
// RtType returns a string
func (obj *bgpRouteTarget) HasRtType() bool {
	return obj.obj.RtType != nil
}

func (obj *bgpRouteTarget) SetRtType(value BgpRouteTargetRtTypeEnum) BgpRouteTarget {
	intValue, ok := otg.BgpRouteTarget_RtType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpRouteTargetRtTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpRouteTarget_RtType_Enum(intValue)
	obj.obj.RtType = &enumValue

	return obj
}

// Colon separated Extended Community value of 6 Bytes - AS number: Assigned Number.   Example - for the as_2octet or as_4octet "60005:100",  for ipv4_address "1.1.1.1:100"
// RtValue returns a string
func (obj *bgpRouteTarget) RtValue() string {

	return *obj.obj.RtValue

}

// Colon separated Extended Community value of 6 Bytes - AS number: Assigned Number.   Example - for the as_2octet or as_4octet "60005:100",  for ipv4_address "1.1.1.1:100"
// RtValue returns a string
func (obj *bgpRouteTarget) HasRtValue() bool {
	return obj.obj.RtValue != nil
}

// Colon separated Extended Community value of 6 Bytes - AS number: Assigned Number.   Example - for the as_2octet or as_4octet "60005:100",  for ipv4_address "1.1.1.1:100"
// SetRtValue sets the string value in the BgpRouteTarget object
func (obj *bgpRouteTarget) SetRtValue(value string) BgpRouteTarget {

	obj.obj.RtValue = &value
	return obj
}

func (obj *bgpRouteTarget) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *bgpRouteTarget) setDefault() {

}
