package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpSrteBindingSubTlv *****
type bgpSrteBindingSubTlv struct {
	validation
	obj          *otg.BgpSrteBindingSubTlv
	marshaller   marshalBgpSrteBindingSubTlv
	unMarshaller unMarshalBgpSrteBindingSubTlv
}

func NewBgpSrteBindingSubTlv() BgpSrteBindingSubTlv {
	obj := bgpSrteBindingSubTlv{obj: &otg.BgpSrteBindingSubTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpSrteBindingSubTlv) msg() *otg.BgpSrteBindingSubTlv {
	return obj.obj
}

func (obj *bgpSrteBindingSubTlv) setMsg(msg *otg.BgpSrteBindingSubTlv) BgpSrteBindingSubTlv {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpSrteBindingSubTlv struct {
	obj *bgpSrteBindingSubTlv
}

type marshalBgpSrteBindingSubTlv interface {
	// ToProto marshals BgpSrteBindingSubTlv to protobuf object *otg.BgpSrteBindingSubTlv
	ToProto() (*otg.BgpSrteBindingSubTlv, error)
	// ToPbText marshals BgpSrteBindingSubTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpSrteBindingSubTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpSrteBindingSubTlv to JSON text
	ToJson() (string, error)
}

type unMarshalbgpSrteBindingSubTlv struct {
	obj *bgpSrteBindingSubTlv
}

type unMarshalBgpSrteBindingSubTlv interface {
	// FromProto unmarshals BgpSrteBindingSubTlv from protobuf object *otg.BgpSrteBindingSubTlv
	FromProto(msg *otg.BgpSrteBindingSubTlv) (BgpSrteBindingSubTlv, error)
	// FromPbText unmarshals BgpSrteBindingSubTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpSrteBindingSubTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpSrteBindingSubTlv from JSON text
	FromJson(value string) error
}

func (obj *bgpSrteBindingSubTlv) Marshal() marshalBgpSrteBindingSubTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpSrteBindingSubTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpSrteBindingSubTlv) Unmarshal() unMarshalBgpSrteBindingSubTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpSrteBindingSubTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpSrteBindingSubTlv) ToProto() (*otg.BgpSrteBindingSubTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpSrteBindingSubTlv) FromProto(msg *otg.BgpSrteBindingSubTlv) (BgpSrteBindingSubTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpSrteBindingSubTlv) ToPbText() (string, error) {
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

func (m *unMarshalbgpSrteBindingSubTlv) FromPbText(value string) error {
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

func (m *marshalbgpSrteBindingSubTlv) ToYaml() (string, error) {
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

func (m *unMarshalbgpSrteBindingSubTlv) FromYaml(value string) error {
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

func (m *marshalbgpSrteBindingSubTlv) ToJson() (string, error) {
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

func (m *unMarshalbgpSrteBindingSubTlv) FromJson(value string) error {
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

func (obj *bgpSrteBindingSubTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpSrteBindingSubTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpSrteBindingSubTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpSrteBindingSubTlv) Clone() (BgpSrteBindingSubTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpSrteBindingSubTlv()
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

// BgpSrteBindingSubTlv is configuration for the binding SID sub-TLV.  This is used to signal the binding SID related information  of the SR Policy candidate path.
type BgpSrteBindingSubTlv interface {
	Validation
	// msg marshals BgpSrteBindingSubTlv to protobuf object *otg.BgpSrteBindingSubTlv
	// and doesn't set defaults
	msg() *otg.BgpSrteBindingSubTlv
	// setMsg unmarshals BgpSrteBindingSubTlv from protobuf object *otg.BgpSrteBindingSubTlv
	// and doesn't set defaults
	setMsg(*otg.BgpSrteBindingSubTlv) BgpSrteBindingSubTlv
	// provides marshal interface
	Marshal() marshalBgpSrteBindingSubTlv
	// provides unmarshal interface
	Unmarshal() unMarshalBgpSrteBindingSubTlv
	// validate validates BgpSrteBindingSubTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpSrteBindingSubTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// BindingSidType returns BgpSrteBindingSubTlvBindingSidTypeEnum, set in BgpSrteBindingSubTlv
	BindingSidType() BgpSrteBindingSubTlvBindingSidTypeEnum
	// SetBindingSidType assigns BgpSrteBindingSubTlvBindingSidTypeEnum provided by user to BgpSrteBindingSubTlv
	SetBindingSidType(value BgpSrteBindingSubTlvBindingSidTypeEnum) BgpSrteBindingSubTlv
	// HasBindingSidType checks if BindingSidType has been set in BgpSrteBindingSubTlv
	HasBindingSidType() bool
	// FourOctetSid returns uint32, set in BgpSrteBindingSubTlv.
	FourOctetSid() uint32
	// SetFourOctetSid assigns uint32 provided by user to BgpSrteBindingSubTlv
	SetFourOctetSid(value uint32) BgpSrteBindingSubTlv
	// HasFourOctetSid checks if FourOctetSid has been set in BgpSrteBindingSubTlv
	HasFourOctetSid() bool
	// Ipv6Sid returns string, set in BgpSrteBindingSubTlv.
	Ipv6Sid() string
	// SetIpv6Sid assigns string provided by user to BgpSrteBindingSubTlv
	SetIpv6Sid(value string) BgpSrteBindingSubTlv
	// HasIpv6Sid checks if Ipv6Sid has been set in BgpSrteBindingSubTlv
	HasIpv6Sid() bool
	// SFlag returns bool, set in BgpSrteBindingSubTlv.
	SFlag() bool
	// SetSFlag assigns bool provided by user to BgpSrteBindingSubTlv
	SetSFlag(value bool) BgpSrteBindingSubTlv
	// HasSFlag checks if SFlag has been set in BgpSrteBindingSubTlv
	HasSFlag() bool
	// IFlag returns bool, set in BgpSrteBindingSubTlv.
	IFlag() bool
	// SetIFlag assigns bool provided by user to BgpSrteBindingSubTlv
	SetIFlag(value bool) BgpSrteBindingSubTlv
	// HasIFlag checks if IFlag has been set in BgpSrteBindingSubTlv
	HasIFlag() bool
}

type BgpSrteBindingSubTlvBindingSidTypeEnum string

// Enum of BindingSidType on BgpSrteBindingSubTlv
var BgpSrteBindingSubTlvBindingSidType = struct {
	NO_BINDING     BgpSrteBindingSubTlvBindingSidTypeEnum
	FOUR_OCTET_SID BgpSrteBindingSubTlvBindingSidTypeEnum
	IPV6_SID       BgpSrteBindingSubTlvBindingSidTypeEnum
}{
	NO_BINDING:     BgpSrteBindingSubTlvBindingSidTypeEnum("no_binding"),
	FOUR_OCTET_SID: BgpSrteBindingSubTlvBindingSidTypeEnum("four_octet_sid"),
	IPV6_SID:       BgpSrteBindingSubTlvBindingSidTypeEnum("ipv6_sid"),
}

func (obj *bgpSrteBindingSubTlv) BindingSidType() BgpSrteBindingSubTlvBindingSidTypeEnum {
	return BgpSrteBindingSubTlvBindingSidTypeEnum(obj.obj.BindingSidType.Enum().String())
}

// Type of the binding SID.  Supported types are "No Binding SID" or "Four Octets Sid" or "IPv6 SID".
// BindingSidType returns a string
func (obj *bgpSrteBindingSubTlv) HasBindingSidType() bool {
	return obj.obj.BindingSidType != nil
}

func (obj *bgpSrteBindingSubTlv) SetBindingSidType(value BgpSrteBindingSubTlvBindingSidTypeEnum) BgpSrteBindingSubTlv {
	intValue, ok := otg.BgpSrteBindingSubTlv_BindingSidType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpSrteBindingSubTlvBindingSidTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpSrteBindingSubTlv_BindingSidType_Enum(intValue)
	obj.obj.BindingSidType = &enumValue

	return obj
}

// Binding SID is encoded in 4 octets.
// FourOctetSid returns a uint32
func (obj *bgpSrteBindingSubTlv) FourOctetSid() uint32 {

	return *obj.obj.FourOctetSid

}

// Binding SID is encoded in 4 octets.
// FourOctetSid returns a uint32
func (obj *bgpSrteBindingSubTlv) HasFourOctetSid() bool {
	return obj.obj.FourOctetSid != nil
}

// Binding SID is encoded in 4 octets.
// SetFourOctetSid sets the uint32 value in the BgpSrteBindingSubTlv object
func (obj *bgpSrteBindingSubTlv) SetFourOctetSid(value uint32) BgpSrteBindingSubTlv {

	obj.obj.FourOctetSid = &value
	return obj
}

// IPv6 SID value.
// Ipv6Sid returns a string
func (obj *bgpSrteBindingSubTlv) Ipv6Sid() string {

	return *obj.obj.Ipv6Sid

}

// IPv6 SID value.
// Ipv6Sid returns a string
func (obj *bgpSrteBindingSubTlv) HasIpv6Sid() bool {
	return obj.obj.Ipv6Sid != nil
}

// IPv6 SID value.
// SetIpv6Sid sets the string value in the BgpSrteBindingSubTlv object
func (obj *bgpSrteBindingSubTlv) SetIpv6Sid(value string) BgpSrteBindingSubTlv {

	obj.obj.Ipv6Sid = &value
	return obj
}

// S-Flag encodes the "Specified-BSID-only" behavior.
// SFlag returns a bool
func (obj *bgpSrteBindingSubTlv) SFlag() bool {

	return *obj.obj.SFlag

}

// S-Flag encodes the "Specified-BSID-only" behavior.
// SFlag returns a bool
func (obj *bgpSrteBindingSubTlv) HasSFlag() bool {
	return obj.obj.SFlag != nil
}

// S-Flag encodes the "Specified-BSID-only" behavior.
// SetSFlag sets the bool value in the BgpSrteBindingSubTlv object
func (obj *bgpSrteBindingSubTlv) SetSFlag(value bool) BgpSrteBindingSubTlv {

	obj.obj.SFlag = &value
	return obj
}

// I-Flag encodes the "Drop Upon Invalid" behavior.
// IFlag returns a bool
func (obj *bgpSrteBindingSubTlv) IFlag() bool {

	return *obj.obj.IFlag

}

// I-Flag encodes the "Drop Upon Invalid" behavior.
// IFlag returns a bool
func (obj *bgpSrteBindingSubTlv) HasIFlag() bool {
	return obj.obj.IFlag != nil
}

// I-Flag encodes the "Drop Upon Invalid" behavior.
// SetIFlag sets the bool value in the BgpSrteBindingSubTlv object
func (obj *bgpSrteBindingSubTlv) SetIFlag(value bool) BgpSrteBindingSubTlv {

	obj.obj.IFlag = &value
	return obj
}

func (obj *bgpSrteBindingSubTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv6Sid != nil {

		err := obj.validateIpv6(obj.Ipv6Sid())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteBindingSubTlv.Ipv6Sid"))
		}

	}

}

func (obj *bgpSrteBindingSubTlv) setDefault() {
	if obj.obj.BindingSidType == nil {
		obj.SetBindingSidType(BgpSrteBindingSubTlvBindingSidType.NO_BINDING)

	}
	if obj.obj.SFlag == nil {
		obj.SetSFlag(false)
	}
	if obj.obj.IFlag == nil {
		obj.SetIFlag(false)
	}

}
