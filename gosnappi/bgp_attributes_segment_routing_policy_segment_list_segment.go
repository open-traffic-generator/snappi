package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesSegmentRoutingPolicySegmentListSegment *****
type bgpAttributesSegmentRoutingPolicySegmentListSegment struct {
	validation
	obj          *otg.BgpAttributesSegmentRoutingPolicySegmentListSegment
	marshaller   marshalBgpAttributesSegmentRoutingPolicySegmentListSegment
	unMarshaller unMarshalBgpAttributesSegmentRoutingPolicySegmentListSegment
	typeAHolder  BgpAttributesSegmentRoutingPolicyTypeA
	typeBHolder  BgpAttributesSegmentRoutingPolicyTypeB
	typeCHolder  BgpAttributesSegmentRoutingPolicyTypeC
	typeDHolder  BgpAttributesSegmentRoutingPolicyTypeD
	typeEHolder  BgpAttributesSegmentRoutingPolicyTypeE
	typeFHolder  BgpAttributesSegmentRoutingPolicyTypeF
	typeGHolder  BgpAttributesSegmentRoutingPolicyTypeG
	typeHHolder  BgpAttributesSegmentRoutingPolicyTypeH
	typeIHolder  BgpAttributesSegmentRoutingPolicyTypeI
	typeJHolder  BgpAttributesSegmentRoutingPolicyTypeJ
	typeKHolder  BgpAttributesSegmentRoutingPolicyTypeK
}

func NewBgpAttributesSegmentRoutingPolicySegmentListSegment() BgpAttributesSegmentRoutingPolicySegmentListSegment {
	obj := bgpAttributesSegmentRoutingPolicySegmentListSegment{obj: &otg.BgpAttributesSegmentRoutingPolicySegmentListSegment{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) msg() *otg.BgpAttributesSegmentRoutingPolicySegmentListSegment {
	return obj.obj
}

func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) setMsg(msg *otg.BgpAttributesSegmentRoutingPolicySegmentListSegment) BgpAttributesSegmentRoutingPolicySegmentListSegment {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesSegmentRoutingPolicySegmentListSegment struct {
	obj *bgpAttributesSegmentRoutingPolicySegmentListSegment
}

type marshalBgpAttributesSegmentRoutingPolicySegmentListSegment interface {
	// ToProto marshals BgpAttributesSegmentRoutingPolicySegmentListSegment to protobuf object *otg.BgpAttributesSegmentRoutingPolicySegmentListSegment
	ToProto() (*otg.BgpAttributesSegmentRoutingPolicySegmentListSegment, error)
	// ToPbText marshals BgpAttributesSegmentRoutingPolicySegmentListSegment to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesSegmentRoutingPolicySegmentListSegment to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesSegmentRoutingPolicySegmentListSegment to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpAttributesSegmentRoutingPolicySegmentListSegment to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpAttributesSegmentRoutingPolicySegmentListSegment struct {
	obj *bgpAttributesSegmentRoutingPolicySegmentListSegment
}

type unMarshalBgpAttributesSegmentRoutingPolicySegmentListSegment interface {
	// FromProto unmarshals BgpAttributesSegmentRoutingPolicySegmentListSegment from protobuf object *otg.BgpAttributesSegmentRoutingPolicySegmentListSegment
	FromProto(msg *otg.BgpAttributesSegmentRoutingPolicySegmentListSegment) (BgpAttributesSegmentRoutingPolicySegmentListSegment, error)
	// FromPbText unmarshals BgpAttributesSegmentRoutingPolicySegmentListSegment from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesSegmentRoutingPolicySegmentListSegment from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesSegmentRoutingPolicySegmentListSegment from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) Marshal() marshalBgpAttributesSegmentRoutingPolicySegmentListSegment {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesSegmentRoutingPolicySegmentListSegment{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicySegmentListSegment {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesSegmentRoutingPolicySegmentListSegment{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesSegmentRoutingPolicySegmentListSegment) ToProto() (*otg.BgpAttributesSegmentRoutingPolicySegmentListSegment, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesSegmentRoutingPolicySegmentListSegment) FromProto(msg *otg.BgpAttributesSegmentRoutingPolicySegmentListSegment) (BgpAttributesSegmentRoutingPolicySegmentListSegment, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesSegmentRoutingPolicySegmentListSegment) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicySegmentListSegment) FromPbText(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicySegmentListSegment) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicySegmentListSegment) FromYaml(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicySegmentListSegment) ToJsonRaw() (string, error) {
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

func (m *marshalbgpAttributesSegmentRoutingPolicySegmentListSegment) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicySegmentListSegment) FromJson(value string) error {
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

func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) Clone() (BgpAttributesSegmentRoutingPolicySegmentListSegment, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesSegmentRoutingPolicySegmentListSegment()
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

func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) setNil() {
	obj.typeAHolder = nil
	obj.typeBHolder = nil
	obj.typeCHolder = nil
	obj.typeDHolder = nil
	obj.typeEHolder = nil
	obj.typeFHolder = nil
	obj.typeGHolder = nil
	obj.typeHHolder = nil
	obj.typeIHolder = nil
	obj.typeJHolder = nil
	obj.typeKHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpAttributesSegmentRoutingPolicySegmentListSegment is a Segment sub-TLV describes a single segment in a segment list  i.e., a single
// element of the explicit path. The Segment sub-TLVs are optional.
// Segment Types A and B are defined as described in 2.4.4.2.
// Segment Types C upto K are defined as described in in draft-ietf-idr-bgp-sr-segtypes-ext-03  .
type BgpAttributesSegmentRoutingPolicySegmentListSegment interface {
	Validation
	// msg marshals BgpAttributesSegmentRoutingPolicySegmentListSegment to protobuf object *otg.BgpAttributesSegmentRoutingPolicySegmentListSegment
	// and doesn't set defaults
	msg() *otg.BgpAttributesSegmentRoutingPolicySegmentListSegment
	// setMsg unmarshals BgpAttributesSegmentRoutingPolicySegmentListSegment from protobuf object *otg.BgpAttributesSegmentRoutingPolicySegmentListSegment
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesSegmentRoutingPolicySegmentListSegment) BgpAttributesSegmentRoutingPolicySegmentListSegment
	// provides marshal interface
	Marshal() marshalBgpAttributesSegmentRoutingPolicySegmentListSegment
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicySegmentListSegment
	// validate validates BgpAttributesSegmentRoutingPolicySegmentListSegment
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesSegmentRoutingPolicySegmentListSegment, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum, set in BgpAttributesSegmentRoutingPolicySegmentListSegment
	Choice() BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum
	// setChoice assigns BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum provided by user to BgpAttributesSegmentRoutingPolicySegmentListSegment
	setChoice(value BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum) BgpAttributesSegmentRoutingPolicySegmentListSegment
	// TypeA returns BgpAttributesSegmentRoutingPolicyTypeA, set in BgpAttributesSegmentRoutingPolicySegmentListSegment.
	// BgpAttributesSegmentRoutingPolicyTypeA is type A: SID only, in the form of MPLS Label.
	// It is encoded as a Segment of Type 1 in the SEGMENT_LIST sub-tlv.
	TypeA() BgpAttributesSegmentRoutingPolicyTypeA
	// SetTypeA assigns BgpAttributesSegmentRoutingPolicyTypeA provided by user to BgpAttributesSegmentRoutingPolicySegmentListSegment.
	// BgpAttributesSegmentRoutingPolicyTypeA is type A: SID only, in the form of MPLS Label.
	// It is encoded as a Segment of Type 1 in the SEGMENT_LIST sub-tlv.
	SetTypeA(value BgpAttributesSegmentRoutingPolicyTypeA) BgpAttributesSegmentRoutingPolicySegmentListSegment
	// HasTypeA checks if TypeA has been set in BgpAttributesSegmentRoutingPolicySegmentListSegment
	HasTypeA() bool
	// TypeB returns BgpAttributesSegmentRoutingPolicyTypeB, set in BgpAttributesSegmentRoutingPolicySegmentListSegment.
	// BgpAttributesSegmentRoutingPolicyTypeB is type B: SID only, in the form of IPv6 address.
	// It is encoded as a Segment of Type 13 in the SEGMENT_LIST sub-tlv.
	TypeB() BgpAttributesSegmentRoutingPolicyTypeB
	// SetTypeB assigns BgpAttributesSegmentRoutingPolicyTypeB provided by user to BgpAttributesSegmentRoutingPolicySegmentListSegment.
	// BgpAttributesSegmentRoutingPolicyTypeB is type B: SID only, in the form of IPv6 address.
	// It is encoded as a Segment of Type 13 in the SEGMENT_LIST sub-tlv.
	SetTypeB(value BgpAttributesSegmentRoutingPolicyTypeB) BgpAttributesSegmentRoutingPolicySegmentListSegment
	// HasTypeB checks if TypeB has been set in BgpAttributesSegmentRoutingPolicySegmentListSegment
	HasTypeB() bool
	// TypeC returns BgpAttributesSegmentRoutingPolicyTypeC, set in BgpAttributesSegmentRoutingPolicySegmentListSegment.
	// BgpAttributesSegmentRoutingPolicyTypeC is type C: IPv4 Node Address with optional SID.
	// It is encoded as a Segment of Type 3 in the SEGMENT_LIST sub-tlv.
	TypeC() BgpAttributesSegmentRoutingPolicyTypeC
	// SetTypeC assigns BgpAttributesSegmentRoutingPolicyTypeC provided by user to BgpAttributesSegmentRoutingPolicySegmentListSegment.
	// BgpAttributesSegmentRoutingPolicyTypeC is type C: IPv4 Node Address with optional SID.
	// It is encoded as a Segment of Type 3 in the SEGMENT_LIST sub-tlv.
	SetTypeC(value BgpAttributesSegmentRoutingPolicyTypeC) BgpAttributesSegmentRoutingPolicySegmentListSegment
	// HasTypeC checks if TypeC has been set in BgpAttributesSegmentRoutingPolicySegmentListSegment
	HasTypeC() bool
	// TypeD returns BgpAttributesSegmentRoutingPolicyTypeD, set in BgpAttributesSegmentRoutingPolicySegmentListSegment.
	// BgpAttributesSegmentRoutingPolicyTypeD is type D: IPv6 Node Address with optional SID for SR MPLS.
	// It is encoded as a Segment of Type 4 in the SEGMENT_LIST sub-tlv.
	TypeD() BgpAttributesSegmentRoutingPolicyTypeD
	// SetTypeD assigns BgpAttributesSegmentRoutingPolicyTypeD provided by user to BgpAttributesSegmentRoutingPolicySegmentListSegment.
	// BgpAttributesSegmentRoutingPolicyTypeD is type D: IPv6 Node Address with optional SID for SR MPLS.
	// It is encoded as a Segment of Type 4 in the SEGMENT_LIST sub-tlv.
	SetTypeD(value BgpAttributesSegmentRoutingPolicyTypeD) BgpAttributesSegmentRoutingPolicySegmentListSegment
	// HasTypeD checks if TypeD has been set in BgpAttributesSegmentRoutingPolicySegmentListSegment
	HasTypeD() bool
	// TypeE returns BgpAttributesSegmentRoutingPolicyTypeE, set in BgpAttributesSegmentRoutingPolicySegmentListSegment.
	// BgpAttributesSegmentRoutingPolicyTypeE is type E: IPv4 Address and Local Interface ID with optional SID
	// It is encoded as a Segment of Type 5 in the SEGMENT_LIST sub-tlv.
	TypeE() BgpAttributesSegmentRoutingPolicyTypeE
	// SetTypeE assigns BgpAttributesSegmentRoutingPolicyTypeE provided by user to BgpAttributesSegmentRoutingPolicySegmentListSegment.
	// BgpAttributesSegmentRoutingPolicyTypeE is type E: IPv4 Address and Local Interface ID with optional SID
	// It is encoded as a Segment of Type 5 in the SEGMENT_LIST sub-tlv.
	SetTypeE(value BgpAttributesSegmentRoutingPolicyTypeE) BgpAttributesSegmentRoutingPolicySegmentListSegment
	// HasTypeE checks if TypeE has been set in BgpAttributesSegmentRoutingPolicySegmentListSegment
	HasTypeE() bool
	// TypeF returns BgpAttributesSegmentRoutingPolicyTypeF, set in BgpAttributesSegmentRoutingPolicySegmentListSegment.
	// BgpAttributesSegmentRoutingPolicyTypeF is type F: IPv4 Local and Remote addresses with optional SR-MPLS SID.
	// It is encoded as a Segment of Type 6 in the SEGMENT_LIST sub-tlv.
	TypeF() BgpAttributesSegmentRoutingPolicyTypeF
	// SetTypeF assigns BgpAttributesSegmentRoutingPolicyTypeF provided by user to BgpAttributesSegmentRoutingPolicySegmentListSegment.
	// BgpAttributesSegmentRoutingPolicyTypeF is type F: IPv4 Local and Remote addresses with optional SR-MPLS SID.
	// It is encoded as a Segment of Type 6 in the SEGMENT_LIST sub-tlv.
	SetTypeF(value BgpAttributesSegmentRoutingPolicyTypeF) BgpAttributesSegmentRoutingPolicySegmentListSegment
	// HasTypeF checks if TypeF has been set in BgpAttributesSegmentRoutingPolicySegmentListSegment
	HasTypeF() bool
	// TypeG returns BgpAttributesSegmentRoutingPolicyTypeG, set in BgpAttributesSegmentRoutingPolicySegmentListSegment.
	// BgpAttributesSegmentRoutingPolicyTypeG is type G: IPv6 Address, Interface ID for local and remote pair with optional SID for SR MPLS.
	// It is encoded as a Segment of Type 7 in the SEGMENT_LIST sub-tlv.
	TypeG() BgpAttributesSegmentRoutingPolicyTypeG
	// SetTypeG assigns BgpAttributesSegmentRoutingPolicyTypeG provided by user to BgpAttributesSegmentRoutingPolicySegmentListSegment.
	// BgpAttributesSegmentRoutingPolicyTypeG is type G: IPv6 Address, Interface ID for local and remote pair with optional SID for SR MPLS.
	// It is encoded as a Segment of Type 7 in the SEGMENT_LIST sub-tlv.
	SetTypeG(value BgpAttributesSegmentRoutingPolicyTypeG) BgpAttributesSegmentRoutingPolicySegmentListSegment
	// HasTypeG checks if TypeG has been set in BgpAttributesSegmentRoutingPolicySegmentListSegment
	HasTypeG() bool
	// TypeH returns BgpAttributesSegmentRoutingPolicyTypeH, set in BgpAttributesSegmentRoutingPolicySegmentListSegment.
	// BgpAttributesSegmentRoutingPolicyTypeH is type H: IPv6 Local and Remote addresses with optional SID for SR MPLS.
	// It is encoded as a Segment of Type 8 in the SEGMENT_LIST sub-tlv.
	TypeH() BgpAttributesSegmentRoutingPolicyTypeH
	// SetTypeH assigns BgpAttributesSegmentRoutingPolicyTypeH provided by user to BgpAttributesSegmentRoutingPolicySegmentListSegment.
	// BgpAttributesSegmentRoutingPolicyTypeH is type H: IPv6 Local and Remote addresses with optional SID for SR MPLS.
	// It is encoded as a Segment of Type 8 in the SEGMENT_LIST sub-tlv.
	SetTypeH(value BgpAttributesSegmentRoutingPolicyTypeH) BgpAttributesSegmentRoutingPolicySegmentListSegment
	// HasTypeH checks if TypeH has been set in BgpAttributesSegmentRoutingPolicySegmentListSegment
	HasTypeH() bool
	// TypeI returns BgpAttributesSegmentRoutingPolicyTypeI, set in BgpAttributesSegmentRoutingPolicySegmentListSegment.
	// BgpAttributesSegmentRoutingPolicyTypeI is type I: IPv6 Node Address with optional SR Algorithm and optional SRv6 SID.
	// It is encoded as a Segment of Type 14 in the SEGMENT_LIST sub-tlv.
	TypeI() BgpAttributesSegmentRoutingPolicyTypeI
	// SetTypeI assigns BgpAttributesSegmentRoutingPolicyTypeI provided by user to BgpAttributesSegmentRoutingPolicySegmentListSegment.
	// BgpAttributesSegmentRoutingPolicyTypeI is type I: IPv6 Node Address with optional SR Algorithm and optional SRv6 SID.
	// It is encoded as a Segment of Type 14 in the SEGMENT_LIST sub-tlv.
	SetTypeI(value BgpAttributesSegmentRoutingPolicyTypeI) BgpAttributesSegmentRoutingPolicySegmentListSegment
	// HasTypeI checks if TypeI has been set in BgpAttributesSegmentRoutingPolicySegmentListSegment
	HasTypeI() bool
	// TypeJ returns BgpAttributesSegmentRoutingPolicyTypeJ, set in BgpAttributesSegmentRoutingPolicySegmentListSegment.
	// BgpAttributesSegmentRoutingPolicyTypeJ is type J: IPv6 Address, Interface ID for local and remote pair for SRv6 with optional SID.
	// It is encoded as a Segment of Type 15 in the SEGMENT_LIST sub-tlv.
	TypeJ() BgpAttributesSegmentRoutingPolicyTypeJ
	// SetTypeJ assigns BgpAttributesSegmentRoutingPolicyTypeJ provided by user to BgpAttributesSegmentRoutingPolicySegmentListSegment.
	// BgpAttributesSegmentRoutingPolicyTypeJ is type J: IPv6 Address, Interface ID for local and remote pair for SRv6 with optional SID.
	// It is encoded as a Segment of Type 15 in the SEGMENT_LIST sub-tlv.
	SetTypeJ(value BgpAttributesSegmentRoutingPolicyTypeJ) BgpAttributesSegmentRoutingPolicySegmentListSegment
	// HasTypeJ checks if TypeJ has been set in BgpAttributesSegmentRoutingPolicySegmentListSegment
	HasTypeJ() bool
	// TypeK returns BgpAttributesSegmentRoutingPolicyTypeK, set in BgpAttributesSegmentRoutingPolicySegmentListSegment.
	// BgpAttributesSegmentRoutingPolicyTypeK is type K: IPv6 Local and Remote addresses for SRv6 with optional SID.
	// It is encoded as a Segment of Type 16 in the SEGMENT_LIST sub-tlv.
	TypeK() BgpAttributesSegmentRoutingPolicyTypeK
	// SetTypeK assigns BgpAttributesSegmentRoutingPolicyTypeK provided by user to BgpAttributesSegmentRoutingPolicySegmentListSegment.
	// BgpAttributesSegmentRoutingPolicyTypeK is type K: IPv6 Local and Remote addresses for SRv6 with optional SID.
	// It is encoded as a Segment of Type 16 in the SEGMENT_LIST sub-tlv.
	SetTypeK(value BgpAttributesSegmentRoutingPolicyTypeK) BgpAttributesSegmentRoutingPolicySegmentListSegment
	// HasTypeK checks if TypeK has been set in BgpAttributesSegmentRoutingPolicySegmentListSegment
	HasTypeK() bool
	setNil()
}

type BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum string

// Enum of Choice on BgpAttributesSegmentRoutingPolicySegmentListSegment
var BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice = struct {
	TYPE_A BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum
	TYPE_B BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum
	TYPE_C BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum
	TYPE_D BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum
	TYPE_E BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum
	TYPE_F BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum
	TYPE_G BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum
	TYPE_H BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum
	TYPE_I BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum
	TYPE_J BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum
	TYPE_K BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum
}{
	TYPE_A: BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum("type_a"),
	TYPE_B: BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum("type_b"),
	TYPE_C: BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum("type_c"),
	TYPE_D: BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum("type_d"),
	TYPE_E: BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum("type_e"),
	TYPE_F: BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum("type_f"),
	TYPE_G: BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum("type_g"),
	TYPE_H: BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum("type_h"),
	TYPE_I: BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum("type_i"),
	TYPE_J: BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum("type_j"),
	TYPE_K: BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum("type_k"),
}

func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) Choice() BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum {
	return BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) setChoice(value BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum) BgpAttributesSegmentRoutingPolicySegmentListSegment {
	intValue, ok := otg.BgpAttributesSegmentRoutingPolicySegmentListSegment_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpAttributesSegmentRoutingPolicySegmentListSegment_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.TypeK = nil
	obj.typeKHolder = nil
	obj.obj.TypeJ = nil
	obj.typeJHolder = nil
	obj.obj.TypeI = nil
	obj.typeIHolder = nil
	obj.obj.TypeH = nil
	obj.typeHHolder = nil
	obj.obj.TypeG = nil
	obj.typeGHolder = nil
	obj.obj.TypeF = nil
	obj.typeFHolder = nil
	obj.obj.TypeE = nil
	obj.typeEHolder = nil
	obj.obj.TypeD = nil
	obj.typeDHolder = nil
	obj.obj.TypeC = nil
	obj.typeCHolder = nil
	obj.obj.TypeB = nil
	obj.typeBHolder = nil
	obj.obj.TypeA = nil
	obj.typeAHolder = nil

	if value == BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_A {
		obj.obj.TypeA = NewBgpAttributesSegmentRoutingPolicyTypeA().msg()
	}

	if value == BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_B {
		obj.obj.TypeB = NewBgpAttributesSegmentRoutingPolicyTypeB().msg()
	}

	if value == BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_C {
		obj.obj.TypeC = NewBgpAttributesSegmentRoutingPolicyTypeC().msg()
	}

	if value == BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_D {
		obj.obj.TypeD = NewBgpAttributesSegmentRoutingPolicyTypeD().msg()
	}

	if value == BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_E {
		obj.obj.TypeE = NewBgpAttributesSegmentRoutingPolicyTypeE().msg()
	}

	if value == BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_F {
		obj.obj.TypeF = NewBgpAttributesSegmentRoutingPolicyTypeF().msg()
	}

	if value == BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_G {
		obj.obj.TypeG = NewBgpAttributesSegmentRoutingPolicyTypeG().msg()
	}

	if value == BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_H {
		obj.obj.TypeH = NewBgpAttributesSegmentRoutingPolicyTypeH().msg()
	}

	if value == BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_I {
		obj.obj.TypeI = NewBgpAttributesSegmentRoutingPolicyTypeI().msg()
	}

	if value == BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_J {
		obj.obj.TypeJ = NewBgpAttributesSegmentRoutingPolicyTypeJ().msg()
	}

	if value == BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_K {
		obj.obj.TypeK = NewBgpAttributesSegmentRoutingPolicyTypeK().msg()
	}

	return obj
}

// description is TBD
// TypeA returns a BgpAttributesSegmentRoutingPolicyTypeA
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) TypeA() BgpAttributesSegmentRoutingPolicyTypeA {
	if obj.obj.TypeA == nil {
		obj.setChoice(BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_A)
	}
	if obj.typeAHolder == nil {
		obj.typeAHolder = &bgpAttributesSegmentRoutingPolicyTypeA{obj: obj.obj.TypeA}
	}
	return obj.typeAHolder
}

// description is TBD
// TypeA returns a BgpAttributesSegmentRoutingPolicyTypeA
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) HasTypeA() bool {
	return obj.obj.TypeA != nil
}

// description is TBD
// SetTypeA sets the BgpAttributesSegmentRoutingPolicyTypeA value in the BgpAttributesSegmentRoutingPolicySegmentListSegment object
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) SetTypeA(value BgpAttributesSegmentRoutingPolicyTypeA) BgpAttributesSegmentRoutingPolicySegmentListSegment {
	obj.setChoice(BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_A)
	obj.typeAHolder = nil
	obj.obj.TypeA = value.msg()

	return obj
}

// description is TBD
// TypeB returns a BgpAttributesSegmentRoutingPolicyTypeB
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) TypeB() BgpAttributesSegmentRoutingPolicyTypeB {
	if obj.obj.TypeB == nil {
		obj.setChoice(BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_B)
	}
	if obj.typeBHolder == nil {
		obj.typeBHolder = &bgpAttributesSegmentRoutingPolicyTypeB{obj: obj.obj.TypeB}
	}
	return obj.typeBHolder
}

// description is TBD
// TypeB returns a BgpAttributesSegmentRoutingPolicyTypeB
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) HasTypeB() bool {
	return obj.obj.TypeB != nil
}

// description is TBD
// SetTypeB sets the BgpAttributesSegmentRoutingPolicyTypeB value in the BgpAttributesSegmentRoutingPolicySegmentListSegment object
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) SetTypeB(value BgpAttributesSegmentRoutingPolicyTypeB) BgpAttributesSegmentRoutingPolicySegmentListSegment {
	obj.setChoice(BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_B)
	obj.typeBHolder = nil
	obj.obj.TypeB = value.msg()

	return obj
}

// description is TBD
// TypeC returns a BgpAttributesSegmentRoutingPolicyTypeC
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) TypeC() BgpAttributesSegmentRoutingPolicyTypeC {
	if obj.obj.TypeC == nil {
		obj.setChoice(BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_C)
	}
	if obj.typeCHolder == nil {
		obj.typeCHolder = &bgpAttributesSegmentRoutingPolicyTypeC{obj: obj.obj.TypeC}
	}
	return obj.typeCHolder
}

// description is TBD
// TypeC returns a BgpAttributesSegmentRoutingPolicyTypeC
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) HasTypeC() bool {
	return obj.obj.TypeC != nil
}

// description is TBD
// SetTypeC sets the BgpAttributesSegmentRoutingPolicyTypeC value in the BgpAttributesSegmentRoutingPolicySegmentListSegment object
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) SetTypeC(value BgpAttributesSegmentRoutingPolicyTypeC) BgpAttributesSegmentRoutingPolicySegmentListSegment {
	obj.setChoice(BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_C)
	obj.typeCHolder = nil
	obj.obj.TypeC = value.msg()

	return obj
}

// description is TBD
// TypeD returns a BgpAttributesSegmentRoutingPolicyTypeD
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) TypeD() BgpAttributesSegmentRoutingPolicyTypeD {
	if obj.obj.TypeD == nil {
		obj.setChoice(BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_D)
	}
	if obj.typeDHolder == nil {
		obj.typeDHolder = &bgpAttributesSegmentRoutingPolicyTypeD{obj: obj.obj.TypeD}
	}
	return obj.typeDHolder
}

// description is TBD
// TypeD returns a BgpAttributesSegmentRoutingPolicyTypeD
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) HasTypeD() bool {
	return obj.obj.TypeD != nil
}

// description is TBD
// SetTypeD sets the BgpAttributesSegmentRoutingPolicyTypeD value in the BgpAttributesSegmentRoutingPolicySegmentListSegment object
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) SetTypeD(value BgpAttributesSegmentRoutingPolicyTypeD) BgpAttributesSegmentRoutingPolicySegmentListSegment {
	obj.setChoice(BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_D)
	obj.typeDHolder = nil
	obj.obj.TypeD = value.msg()

	return obj
}

// description is TBD
// TypeE returns a BgpAttributesSegmentRoutingPolicyTypeE
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) TypeE() BgpAttributesSegmentRoutingPolicyTypeE {
	if obj.obj.TypeE == nil {
		obj.setChoice(BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_E)
	}
	if obj.typeEHolder == nil {
		obj.typeEHolder = &bgpAttributesSegmentRoutingPolicyTypeE{obj: obj.obj.TypeE}
	}
	return obj.typeEHolder
}

// description is TBD
// TypeE returns a BgpAttributesSegmentRoutingPolicyTypeE
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) HasTypeE() bool {
	return obj.obj.TypeE != nil
}

// description is TBD
// SetTypeE sets the BgpAttributesSegmentRoutingPolicyTypeE value in the BgpAttributesSegmentRoutingPolicySegmentListSegment object
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) SetTypeE(value BgpAttributesSegmentRoutingPolicyTypeE) BgpAttributesSegmentRoutingPolicySegmentListSegment {
	obj.setChoice(BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_E)
	obj.typeEHolder = nil
	obj.obj.TypeE = value.msg()

	return obj
}

// description is TBD
// TypeF returns a BgpAttributesSegmentRoutingPolicyTypeF
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) TypeF() BgpAttributesSegmentRoutingPolicyTypeF {
	if obj.obj.TypeF == nil {
		obj.setChoice(BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_F)
	}
	if obj.typeFHolder == nil {
		obj.typeFHolder = &bgpAttributesSegmentRoutingPolicyTypeF{obj: obj.obj.TypeF}
	}
	return obj.typeFHolder
}

// description is TBD
// TypeF returns a BgpAttributesSegmentRoutingPolicyTypeF
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) HasTypeF() bool {
	return obj.obj.TypeF != nil
}

// description is TBD
// SetTypeF sets the BgpAttributesSegmentRoutingPolicyTypeF value in the BgpAttributesSegmentRoutingPolicySegmentListSegment object
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) SetTypeF(value BgpAttributesSegmentRoutingPolicyTypeF) BgpAttributesSegmentRoutingPolicySegmentListSegment {
	obj.setChoice(BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_F)
	obj.typeFHolder = nil
	obj.obj.TypeF = value.msg()

	return obj
}

// description is TBD
// TypeG returns a BgpAttributesSegmentRoutingPolicyTypeG
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) TypeG() BgpAttributesSegmentRoutingPolicyTypeG {
	if obj.obj.TypeG == nil {
		obj.setChoice(BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_G)
	}
	if obj.typeGHolder == nil {
		obj.typeGHolder = &bgpAttributesSegmentRoutingPolicyTypeG{obj: obj.obj.TypeG}
	}
	return obj.typeGHolder
}

// description is TBD
// TypeG returns a BgpAttributesSegmentRoutingPolicyTypeG
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) HasTypeG() bool {
	return obj.obj.TypeG != nil
}

// description is TBD
// SetTypeG sets the BgpAttributesSegmentRoutingPolicyTypeG value in the BgpAttributesSegmentRoutingPolicySegmentListSegment object
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) SetTypeG(value BgpAttributesSegmentRoutingPolicyTypeG) BgpAttributesSegmentRoutingPolicySegmentListSegment {
	obj.setChoice(BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_G)
	obj.typeGHolder = nil
	obj.obj.TypeG = value.msg()

	return obj
}

// description is TBD
// TypeH returns a BgpAttributesSegmentRoutingPolicyTypeH
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) TypeH() BgpAttributesSegmentRoutingPolicyTypeH {
	if obj.obj.TypeH == nil {
		obj.setChoice(BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_H)
	}
	if obj.typeHHolder == nil {
		obj.typeHHolder = &bgpAttributesSegmentRoutingPolicyTypeH{obj: obj.obj.TypeH}
	}
	return obj.typeHHolder
}

// description is TBD
// TypeH returns a BgpAttributesSegmentRoutingPolicyTypeH
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) HasTypeH() bool {
	return obj.obj.TypeH != nil
}

// description is TBD
// SetTypeH sets the BgpAttributesSegmentRoutingPolicyTypeH value in the BgpAttributesSegmentRoutingPolicySegmentListSegment object
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) SetTypeH(value BgpAttributesSegmentRoutingPolicyTypeH) BgpAttributesSegmentRoutingPolicySegmentListSegment {
	obj.setChoice(BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_H)
	obj.typeHHolder = nil
	obj.obj.TypeH = value.msg()

	return obj
}

// description is TBD
// TypeI returns a BgpAttributesSegmentRoutingPolicyTypeI
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) TypeI() BgpAttributesSegmentRoutingPolicyTypeI {
	if obj.obj.TypeI == nil {
		obj.setChoice(BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_I)
	}
	if obj.typeIHolder == nil {
		obj.typeIHolder = &bgpAttributesSegmentRoutingPolicyTypeI{obj: obj.obj.TypeI}
	}
	return obj.typeIHolder
}

// description is TBD
// TypeI returns a BgpAttributesSegmentRoutingPolicyTypeI
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) HasTypeI() bool {
	return obj.obj.TypeI != nil
}

// description is TBD
// SetTypeI sets the BgpAttributesSegmentRoutingPolicyTypeI value in the BgpAttributesSegmentRoutingPolicySegmentListSegment object
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) SetTypeI(value BgpAttributesSegmentRoutingPolicyTypeI) BgpAttributesSegmentRoutingPolicySegmentListSegment {
	obj.setChoice(BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_I)
	obj.typeIHolder = nil
	obj.obj.TypeI = value.msg()

	return obj
}

// description is TBD
// TypeJ returns a BgpAttributesSegmentRoutingPolicyTypeJ
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) TypeJ() BgpAttributesSegmentRoutingPolicyTypeJ {
	if obj.obj.TypeJ == nil {
		obj.setChoice(BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_J)
	}
	if obj.typeJHolder == nil {
		obj.typeJHolder = &bgpAttributesSegmentRoutingPolicyTypeJ{obj: obj.obj.TypeJ}
	}
	return obj.typeJHolder
}

// description is TBD
// TypeJ returns a BgpAttributesSegmentRoutingPolicyTypeJ
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) HasTypeJ() bool {
	return obj.obj.TypeJ != nil
}

// description is TBD
// SetTypeJ sets the BgpAttributesSegmentRoutingPolicyTypeJ value in the BgpAttributesSegmentRoutingPolicySegmentListSegment object
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) SetTypeJ(value BgpAttributesSegmentRoutingPolicyTypeJ) BgpAttributesSegmentRoutingPolicySegmentListSegment {
	obj.setChoice(BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_J)
	obj.typeJHolder = nil
	obj.obj.TypeJ = value.msg()

	return obj
}

// description is TBD
// TypeK returns a BgpAttributesSegmentRoutingPolicyTypeK
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) TypeK() BgpAttributesSegmentRoutingPolicyTypeK {
	if obj.obj.TypeK == nil {
		obj.setChoice(BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_K)
	}
	if obj.typeKHolder == nil {
		obj.typeKHolder = &bgpAttributesSegmentRoutingPolicyTypeK{obj: obj.obj.TypeK}
	}
	return obj.typeKHolder
}

// description is TBD
// TypeK returns a BgpAttributesSegmentRoutingPolicyTypeK
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) HasTypeK() bool {
	return obj.obj.TypeK != nil
}

// description is TBD
// SetTypeK sets the BgpAttributesSegmentRoutingPolicyTypeK value in the BgpAttributesSegmentRoutingPolicySegmentListSegment object
func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) SetTypeK(value BgpAttributesSegmentRoutingPolicyTypeK) BgpAttributesSegmentRoutingPolicySegmentListSegment {
	obj.setChoice(BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_K)
	obj.typeKHolder = nil
	obj.obj.TypeK = value.msg()

	return obj
}

func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface BgpAttributesSegmentRoutingPolicySegmentListSegment")
	}

	if obj.obj.TypeA != nil {

		obj.TypeA().validateObj(vObj, set_default)
	}

	if obj.obj.TypeB != nil {

		obj.TypeB().validateObj(vObj, set_default)
	}

	if obj.obj.TypeC != nil {

		obj.TypeC().validateObj(vObj, set_default)
	}

	if obj.obj.TypeD != nil {

		obj.TypeD().validateObj(vObj, set_default)
	}

	if obj.obj.TypeE != nil {

		obj.TypeE().validateObj(vObj, set_default)
	}

	if obj.obj.TypeF != nil {

		obj.TypeF().validateObj(vObj, set_default)
	}

	if obj.obj.TypeG != nil {

		obj.TypeG().validateObj(vObj, set_default)
	}

	if obj.obj.TypeH != nil {

		obj.TypeH().validateObj(vObj, set_default)
	}

	if obj.obj.TypeI != nil {

		obj.TypeI().validateObj(vObj, set_default)
	}

	if obj.obj.TypeJ != nil {

		obj.TypeJ().validateObj(vObj, set_default)
	}

	if obj.obj.TypeK != nil {

		obj.TypeK().validateObj(vObj, set_default)
	}

}

func (obj *bgpAttributesSegmentRoutingPolicySegmentListSegment) setDefault() {
	var choices_set int = 0
	var choice BgpAttributesSegmentRoutingPolicySegmentListSegmentChoiceEnum

	if obj.obj.TypeA != nil {
		choices_set += 1
		choice = BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_A
	}

	if obj.obj.TypeB != nil {
		choices_set += 1
		choice = BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_B
	}

	if obj.obj.TypeC != nil {
		choices_set += 1
		choice = BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_C
	}

	if obj.obj.TypeD != nil {
		choices_set += 1
		choice = BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_D
	}

	if obj.obj.TypeE != nil {
		choices_set += 1
		choice = BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_E
	}

	if obj.obj.TypeF != nil {
		choices_set += 1
		choice = BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_F
	}

	if obj.obj.TypeG != nil {
		choices_set += 1
		choice = BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_G
	}

	if obj.obj.TypeH != nil {
		choices_set += 1
		choice = BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_H
	}

	if obj.obj.TypeI != nil {
		choices_set += 1
		choice = BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_I
	}

	if obj.obj.TypeJ != nil {
		choices_set += 1
		choice = BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_J
	}

	if obj.obj.TypeK != nil {
		choices_set += 1
		choice = BgpAttributesSegmentRoutingPolicySegmentListSegmentChoice.TYPE_K
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in BgpAttributesSegmentRoutingPolicySegmentListSegment")
			}
		} else {
			intVal := otg.BgpAttributesSegmentRoutingPolicySegmentListSegment_Choice_Enum_value[string(choice)]
			enumValue := otg.BgpAttributesSegmentRoutingPolicySegmentListSegment_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
