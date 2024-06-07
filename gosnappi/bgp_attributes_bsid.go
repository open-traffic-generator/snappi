package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesBsid *****
type bgpAttributesBsid struct {
	validation
	obj          *otg.BgpAttributesBsid
	marshaller   marshalBgpAttributesBsid
	unMarshaller unMarshalBgpAttributesBsid
	mplsHolder   BgpAttributesBsidMpls
	srv6Holder   BgpAttributesBsidSrv6
}

func NewBgpAttributesBsid() BgpAttributesBsid {
	obj := bgpAttributesBsid{obj: &otg.BgpAttributesBsid{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesBsid) msg() *otg.BgpAttributesBsid {
	return obj.obj
}

func (obj *bgpAttributesBsid) setMsg(msg *otg.BgpAttributesBsid) BgpAttributesBsid {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesBsid struct {
	obj *bgpAttributesBsid
}

type marshalBgpAttributesBsid interface {
	// ToProto marshals BgpAttributesBsid to protobuf object *otg.BgpAttributesBsid
	ToProto() (*otg.BgpAttributesBsid, error)
	// ToPbText marshals BgpAttributesBsid to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesBsid to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesBsid to JSON text
	ToJson() (string, error)
}

type unMarshalbgpAttributesBsid struct {
	obj *bgpAttributesBsid
}

type unMarshalBgpAttributesBsid interface {
	// FromProto unmarshals BgpAttributesBsid from protobuf object *otg.BgpAttributesBsid
	FromProto(msg *otg.BgpAttributesBsid) (BgpAttributesBsid, error)
	// FromPbText unmarshals BgpAttributesBsid from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesBsid from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesBsid from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesBsid) Marshal() marshalBgpAttributesBsid {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesBsid{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesBsid) Unmarshal() unMarshalBgpAttributesBsid {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesBsid{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesBsid) ToProto() (*otg.BgpAttributesBsid, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesBsid) FromProto(msg *otg.BgpAttributesBsid) (BgpAttributesBsid, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesBsid) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesBsid) FromPbText(value string) error {
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

func (m *marshalbgpAttributesBsid) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesBsid) FromYaml(value string) error {
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

func (m *marshalbgpAttributesBsid) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesBsid) FromJson(value string) error {
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

func (obj *bgpAttributesBsid) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesBsid) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesBsid) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesBsid) Clone() (BgpAttributesBsid, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesBsid()
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

func (obj *bgpAttributesBsid) setNil() {
	obj.mplsHolder = nil
	obj.srv6Holder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpAttributesBsid is the Binding Segment Identifier is an optional sub-tlv of type 13 that can be sent with a SR Policy
// Tunnel Encapsulation attribute.
// When the active candidate path has a specified Binding Segment Identifier, the SR Policy uses that
// BSID if this value (label in MPLS, IPv6 address in SRv6) is available.
// - The format of the sub-TLV is defined in draft-ietf-idr-sr-policy-safi-02 Section 2.4.2 .
// - It is recommended that if SRv6 Binding SID is desired to be signalled, the SRv6 Binding SID sub-TLV that enables
// the specification of the SRv6 Endpoint Behavior should be used.
type BgpAttributesBsid interface {
	Validation
	// msg marshals BgpAttributesBsid to protobuf object *otg.BgpAttributesBsid
	// and doesn't set defaults
	msg() *otg.BgpAttributesBsid
	// setMsg unmarshals BgpAttributesBsid from protobuf object *otg.BgpAttributesBsid
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesBsid) BgpAttributesBsid
	// provides marshal interface
	Marshal() marshalBgpAttributesBsid
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesBsid
	// validate validates BgpAttributesBsid
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesBsid, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns BgpAttributesBsidChoiceEnum, set in BgpAttributesBsid
	Choice() BgpAttributesBsidChoiceEnum
	// setChoice assigns BgpAttributesBsidChoiceEnum provided by user to BgpAttributesBsid
	setChoice(value BgpAttributesBsidChoiceEnum) BgpAttributesBsid
	// Mpls returns BgpAttributesBsidMpls, set in BgpAttributesBsid.
	// BgpAttributesBsidMpls is when the active candidate path has a specified Binding Segment Identifier, the SR Policy uses that BSID defined
	// as a MPLS label.The format of the sub-TLV is defined in draft-ietf-idr-sr-policy-safi-02  Section 2.4.2 .
	Mpls() BgpAttributesBsidMpls
	// SetMpls assigns BgpAttributesBsidMpls provided by user to BgpAttributesBsid.
	// BgpAttributesBsidMpls is when the active candidate path has a specified Binding Segment Identifier, the SR Policy uses that BSID defined
	// as a MPLS label.The format of the sub-TLV is defined in draft-ietf-idr-sr-policy-safi-02  Section 2.4.2 .
	SetMpls(value BgpAttributesBsidMpls) BgpAttributesBsid
	// HasMpls checks if Mpls has been set in BgpAttributesBsid
	HasMpls() bool
	// Srv6 returns BgpAttributesBsidSrv6, set in BgpAttributesBsid.
	// BgpAttributesBsidSrv6 is when the active candidate path has a specified Binding Segment Identifier, the SR Policy uses that BSID defined
	// as an IPv6 Address.The format of the sub-TLV is defined in draft-ietf-idr-sr-policy-safi-02 Section 2.4.2 .
	Srv6() BgpAttributesBsidSrv6
	// SetSrv6 assigns BgpAttributesBsidSrv6 provided by user to BgpAttributesBsid.
	// BgpAttributesBsidSrv6 is when the active candidate path has a specified Binding Segment Identifier, the SR Policy uses that BSID defined
	// as an IPv6 Address.The format of the sub-TLV is defined in draft-ietf-idr-sr-policy-safi-02 Section 2.4.2 .
	SetSrv6(value BgpAttributesBsidSrv6) BgpAttributesBsid
	// HasSrv6 checks if Srv6 has been set in BgpAttributesBsid
	HasSrv6() bool
	setNil()
}

type BgpAttributesBsidChoiceEnum string

// Enum of Choice on BgpAttributesBsid
var BgpAttributesBsidChoice = struct {
	MPLS BgpAttributesBsidChoiceEnum
	SRV6 BgpAttributesBsidChoiceEnum
}{
	MPLS: BgpAttributesBsidChoiceEnum("mpls"),
	SRV6: BgpAttributesBsidChoiceEnum("srv6"),
}

func (obj *bgpAttributesBsid) Choice() BgpAttributesBsidChoiceEnum {
	return BgpAttributesBsidChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *bgpAttributesBsid) setChoice(value BgpAttributesBsidChoiceEnum) BgpAttributesBsid {
	intValue, ok := otg.BgpAttributesBsid_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpAttributesBsidChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpAttributesBsid_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Srv6 = nil
	obj.srv6Holder = nil
	obj.obj.Mpls = nil
	obj.mplsHolder = nil

	if value == BgpAttributesBsidChoice.MPLS {
		obj.obj.Mpls = NewBgpAttributesBsidMpls().msg()
	}

	if value == BgpAttributesBsidChoice.SRV6 {
		obj.obj.Srv6 = NewBgpAttributesBsidSrv6().msg()
	}

	return obj
}

// description is TBD
// Mpls returns a BgpAttributesBsidMpls
func (obj *bgpAttributesBsid) Mpls() BgpAttributesBsidMpls {
	if obj.obj.Mpls == nil {
		obj.setChoice(BgpAttributesBsidChoice.MPLS)
	}
	if obj.mplsHolder == nil {
		obj.mplsHolder = &bgpAttributesBsidMpls{obj: obj.obj.Mpls}
	}
	return obj.mplsHolder
}

// description is TBD
// Mpls returns a BgpAttributesBsidMpls
func (obj *bgpAttributesBsid) HasMpls() bool {
	return obj.obj.Mpls != nil
}

// description is TBD
// SetMpls sets the BgpAttributesBsidMpls value in the BgpAttributesBsid object
func (obj *bgpAttributesBsid) SetMpls(value BgpAttributesBsidMpls) BgpAttributesBsid {
	obj.setChoice(BgpAttributesBsidChoice.MPLS)
	obj.mplsHolder = nil
	obj.obj.Mpls = value.msg()

	return obj
}

// description is TBD
// Srv6 returns a BgpAttributesBsidSrv6
func (obj *bgpAttributesBsid) Srv6() BgpAttributesBsidSrv6 {
	if obj.obj.Srv6 == nil {
		obj.setChoice(BgpAttributesBsidChoice.SRV6)
	}
	if obj.srv6Holder == nil {
		obj.srv6Holder = &bgpAttributesBsidSrv6{obj: obj.obj.Srv6}
	}
	return obj.srv6Holder
}

// description is TBD
// Srv6 returns a BgpAttributesBsidSrv6
func (obj *bgpAttributesBsid) HasSrv6() bool {
	return obj.obj.Srv6 != nil
}

// description is TBD
// SetSrv6 sets the BgpAttributesBsidSrv6 value in the BgpAttributesBsid object
func (obj *bgpAttributesBsid) SetSrv6(value BgpAttributesBsidSrv6) BgpAttributesBsid {
	obj.setChoice(BgpAttributesBsidChoice.SRV6)
	obj.srv6Holder = nil
	obj.obj.Srv6 = value.msg()

	return obj
}

func (obj *bgpAttributesBsid) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface BgpAttributesBsid")
	}

	if obj.obj.Mpls != nil {

		obj.Mpls().validateObj(vObj, set_default)
	}

	if obj.obj.Srv6 != nil {

		obj.Srv6().validateObj(vObj, set_default)
	}

}

func (obj *bgpAttributesBsid) setDefault() {
	var choices_set int = 0
	var choice BgpAttributesBsidChoiceEnum

	if obj.obj.Mpls != nil {
		choices_set += 1
		choice = BgpAttributesBsidChoice.MPLS
	}

	if obj.obj.Srv6 != nil {
		choices_set += 1
		choice = BgpAttributesBsidChoice.SRV6
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in BgpAttributesBsid")
			}
		} else {
			intVal := otg.BgpAttributesBsid_Choice_Enum_value[string(choice)]
			enumValue := otg.BgpAttributesBsid_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
