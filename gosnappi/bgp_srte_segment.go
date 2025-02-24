package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpSrteSegment *****
type bgpSrteSegment struct {
	validation
	obj          *otg.BgpSrteSegment
	marshaller   marshalBgpSrteSegment
	unMarshaller unMarshalBgpSrteSegment
	typeAHolder  BgpSrteSegmentATypeSubTlv
	typeBHolder  BgpSrteSegmentBTypeSubTlv
	typeCHolder  BgpSrteSegmentCTypeSubTlv
	typeDHolder  BgpSrteSegmentDTypeSubTlv
	typeEHolder  BgpSrteSegmentETypeSubTlv
	typeFHolder  BgpSrteSegmentFTypeSubTlv
	typeGHolder  BgpSrteSegmentGTypeSubTlv
	typeHHolder  BgpSrteSegmentHTypeSubTlv
	typeIHolder  BgpSrteSegmentITypeSubTlv
	typeJHolder  BgpSrteSegmentJTypeSubTlv
	typeKHolder  BgpSrteSegmentKTypeSubTlv
}

func NewBgpSrteSegment() BgpSrteSegment {
	obj := bgpSrteSegment{obj: &otg.BgpSrteSegment{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpSrteSegment) msg() *otg.BgpSrteSegment {
	return obj.obj
}

func (obj *bgpSrteSegment) setMsg(msg *otg.BgpSrteSegment) BgpSrteSegment {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpSrteSegment struct {
	obj *bgpSrteSegment
}

type marshalBgpSrteSegment interface {
	// ToProto marshals BgpSrteSegment to protobuf object *otg.BgpSrteSegment
	ToProto() (*otg.BgpSrteSegment, error)
	// ToPbText marshals BgpSrteSegment to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpSrteSegment to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpSrteSegment to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpSrteSegment to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpSrteSegment struct {
	obj *bgpSrteSegment
}

type unMarshalBgpSrteSegment interface {
	// FromProto unmarshals BgpSrteSegment from protobuf object *otg.BgpSrteSegment
	FromProto(msg *otg.BgpSrteSegment) (BgpSrteSegment, error)
	// FromPbText unmarshals BgpSrteSegment from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpSrteSegment from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpSrteSegment from JSON text
	FromJson(value string) error
}

func (obj *bgpSrteSegment) Marshal() marshalBgpSrteSegment {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpSrteSegment{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpSrteSegment) Unmarshal() unMarshalBgpSrteSegment {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpSrteSegment{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpSrteSegment) ToProto() (*otg.BgpSrteSegment, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpSrteSegment) FromProto(msg *otg.BgpSrteSegment) (BgpSrteSegment, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpSrteSegment) ToPbText() (string, error) {
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

func (m *unMarshalbgpSrteSegment) FromPbText(value string) error {
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

func (m *marshalbgpSrteSegment) ToYaml() (string, error) {
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

func (m *unMarshalbgpSrteSegment) FromYaml(value string) error {
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

func (m *marshalbgpSrteSegment) ToJsonRaw() (string, error) {
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

func (m *marshalbgpSrteSegment) ToJson() (string, error) {
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

func (m *unMarshalbgpSrteSegment) FromJson(value string) error {
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

func (obj *bgpSrteSegment) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpSrteSegment) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpSrteSegment) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpSrteSegment) Clone() (BgpSrteSegment, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpSrteSegment()
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

func (obj *bgpSrteSegment) setNil() {
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

// BgpSrteSegment is a Segment sub-TLV describes a single segment in a segment list  i.e., a single element of the explicit path. The Segment sub-TLVs are optional.
type BgpSrteSegment interface {
	Validation
	// msg marshals BgpSrteSegment to protobuf object *otg.BgpSrteSegment
	// and doesn't set defaults
	msg() *otg.BgpSrteSegment
	// setMsg unmarshals BgpSrteSegment from protobuf object *otg.BgpSrteSegment
	// and doesn't set defaults
	setMsg(*otg.BgpSrteSegment) BgpSrteSegment
	// provides marshal interface
	Marshal() marshalBgpSrteSegment
	// provides unmarshal interface
	Unmarshal() unMarshalBgpSrteSegment
	// validate validates BgpSrteSegment
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpSrteSegment, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SegmentType returns BgpSrteSegmentSegmentTypeEnum, set in BgpSrteSegment
	SegmentType() BgpSrteSegmentSegmentTypeEnum
	// SetSegmentType assigns BgpSrteSegmentSegmentTypeEnum provided by user to BgpSrteSegment
	SetSegmentType(value BgpSrteSegmentSegmentTypeEnum) BgpSrteSegment
	// TypeA returns BgpSrteSegmentATypeSubTlv, set in BgpSrteSegment.
	// BgpSrteSegmentATypeSubTlv is type  A: SID only, in the form of MPLS Label.
	TypeA() BgpSrteSegmentATypeSubTlv
	// SetTypeA assigns BgpSrteSegmentATypeSubTlv provided by user to BgpSrteSegment.
	// BgpSrteSegmentATypeSubTlv is type  A: SID only, in the form of MPLS Label.
	SetTypeA(value BgpSrteSegmentATypeSubTlv) BgpSrteSegment
	// HasTypeA checks if TypeA has been set in BgpSrteSegment
	HasTypeA() bool
	// TypeB returns BgpSrteSegmentBTypeSubTlv, set in BgpSrteSegment.
	// BgpSrteSegmentBTypeSubTlv is type  B: SID only, in the form of IPv6 address.
	TypeB() BgpSrteSegmentBTypeSubTlv
	// SetTypeB assigns BgpSrteSegmentBTypeSubTlv provided by user to BgpSrteSegment.
	// BgpSrteSegmentBTypeSubTlv is type  B: SID only, in the form of IPv6 address.
	SetTypeB(value BgpSrteSegmentBTypeSubTlv) BgpSrteSegment
	// HasTypeB checks if TypeB has been set in BgpSrteSegment
	HasTypeB() bool
	// TypeC returns BgpSrteSegmentCTypeSubTlv, set in BgpSrteSegment.
	// BgpSrteSegmentCTypeSubTlv is type C: IPv4 Node Address with optional SID.
	TypeC() BgpSrteSegmentCTypeSubTlv
	// SetTypeC assigns BgpSrteSegmentCTypeSubTlv provided by user to BgpSrteSegment.
	// BgpSrteSegmentCTypeSubTlv is type C: IPv4 Node Address with optional SID.
	SetTypeC(value BgpSrteSegmentCTypeSubTlv) BgpSrteSegment
	// HasTypeC checks if TypeC has been set in BgpSrteSegment
	HasTypeC() bool
	// TypeD returns BgpSrteSegmentDTypeSubTlv, set in BgpSrteSegment.
	// BgpSrteSegmentDTypeSubTlv is type D: IPv6 Node Address with optional SID for SR MPLS.
	TypeD() BgpSrteSegmentDTypeSubTlv
	// SetTypeD assigns BgpSrteSegmentDTypeSubTlv provided by user to BgpSrteSegment.
	// BgpSrteSegmentDTypeSubTlv is type D: IPv6 Node Address with optional SID for SR MPLS.
	SetTypeD(value BgpSrteSegmentDTypeSubTlv) BgpSrteSegment
	// HasTypeD checks if TypeD has been set in BgpSrteSegment
	HasTypeD() bool
	// TypeE returns BgpSrteSegmentETypeSubTlv, set in BgpSrteSegment.
	// BgpSrteSegmentETypeSubTlv is type E: IPv4 Address and Local Interface ID with optional SID
	TypeE() BgpSrteSegmentETypeSubTlv
	// SetTypeE assigns BgpSrteSegmentETypeSubTlv provided by user to BgpSrteSegment.
	// BgpSrteSegmentETypeSubTlv is type E: IPv4 Address and Local Interface ID with optional SID
	SetTypeE(value BgpSrteSegmentETypeSubTlv) BgpSrteSegment
	// HasTypeE checks if TypeE has been set in BgpSrteSegment
	HasTypeE() bool
	// TypeF returns BgpSrteSegmentFTypeSubTlv, set in BgpSrteSegment.
	// BgpSrteSegmentFTypeSubTlv is type F: IPv4 Local and Remote addresses with optional SID.
	TypeF() BgpSrteSegmentFTypeSubTlv
	// SetTypeF assigns BgpSrteSegmentFTypeSubTlv provided by user to BgpSrteSegment.
	// BgpSrteSegmentFTypeSubTlv is type F: IPv4 Local and Remote addresses with optional SID.
	SetTypeF(value BgpSrteSegmentFTypeSubTlv) BgpSrteSegment
	// HasTypeF checks if TypeF has been set in BgpSrteSegment
	HasTypeF() bool
	// TypeG returns BgpSrteSegmentGTypeSubTlv, set in BgpSrteSegment.
	// BgpSrteSegmentGTypeSubTlv is type G: IPv6 Address, Interface ID for local and remote pair with optional SID for SR MPLS.
	TypeG() BgpSrteSegmentGTypeSubTlv
	// SetTypeG assigns BgpSrteSegmentGTypeSubTlv provided by user to BgpSrteSegment.
	// BgpSrteSegmentGTypeSubTlv is type G: IPv6 Address, Interface ID for local and remote pair with optional SID for SR MPLS.
	SetTypeG(value BgpSrteSegmentGTypeSubTlv) BgpSrteSegment
	// HasTypeG checks if TypeG has been set in BgpSrteSegment
	HasTypeG() bool
	// TypeH returns BgpSrteSegmentHTypeSubTlv, set in BgpSrteSegment.
	// BgpSrteSegmentHTypeSubTlv is type H: IPv6 Local and Remote addresses with optional SID for SR MPLS.
	TypeH() BgpSrteSegmentHTypeSubTlv
	// SetTypeH assigns BgpSrteSegmentHTypeSubTlv provided by user to BgpSrteSegment.
	// BgpSrteSegmentHTypeSubTlv is type H: IPv6 Local and Remote addresses with optional SID for SR MPLS.
	SetTypeH(value BgpSrteSegmentHTypeSubTlv) BgpSrteSegment
	// HasTypeH checks if TypeH has been set in BgpSrteSegment
	HasTypeH() bool
	// TypeI returns BgpSrteSegmentITypeSubTlv, set in BgpSrteSegment.
	// BgpSrteSegmentITypeSubTlv is type I: IPv6 Node Address with optional SRv6 SID.
	TypeI() BgpSrteSegmentITypeSubTlv
	// SetTypeI assigns BgpSrteSegmentITypeSubTlv provided by user to BgpSrteSegment.
	// BgpSrteSegmentITypeSubTlv is type I: IPv6 Node Address with optional SRv6 SID.
	SetTypeI(value BgpSrteSegmentITypeSubTlv) BgpSrteSegment
	// HasTypeI checks if TypeI has been set in BgpSrteSegment
	HasTypeI() bool
	// TypeJ returns BgpSrteSegmentJTypeSubTlv, set in BgpSrteSegment.
	// BgpSrteSegmentJTypeSubTlv is type J: IPv6 Address, Interface ID for local and remote pair for SRv6 with optional SID.
	TypeJ() BgpSrteSegmentJTypeSubTlv
	// SetTypeJ assigns BgpSrteSegmentJTypeSubTlv provided by user to BgpSrteSegment.
	// BgpSrteSegmentJTypeSubTlv is type J: IPv6 Address, Interface ID for local and remote pair for SRv6 with optional SID.
	SetTypeJ(value BgpSrteSegmentJTypeSubTlv) BgpSrteSegment
	// HasTypeJ checks if TypeJ has been set in BgpSrteSegment
	HasTypeJ() bool
	// TypeK returns BgpSrteSegmentKTypeSubTlv, set in BgpSrteSegment.
	// BgpSrteSegmentKTypeSubTlv is type K: IPv6 Local and Remote addresses for SRv6 with optional SID.
	TypeK() BgpSrteSegmentKTypeSubTlv
	// SetTypeK assigns BgpSrteSegmentKTypeSubTlv provided by user to BgpSrteSegment.
	// BgpSrteSegmentKTypeSubTlv is type K: IPv6 Local and Remote addresses for SRv6 with optional SID.
	SetTypeK(value BgpSrteSegmentKTypeSubTlv) BgpSrteSegment
	// HasTypeK checks if TypeK has been set in BgpSrteSegment
	HasTypeK() bool
	// Name returns string, set in BgpSrteSegment.
	Name() string
	// SetName assigns string provided by user to BgpSrteSegment
	SetName(value string) BgpSrteSegment
	// Active returns bool, set in BgpSrteSegment.
	Active() bool
	// SetActive assigns bool provided by user to BgpSrteSegment
	SetActive(value bool) BgpSrteSegment
	// HasActive checks if Active has been set in BgpSrteSegment
	HasActive() bool
	setNil()
}

type BgpSrteSegmentSegmentTypeEnum string

// Enum of SegmentType on BgpSrteSegment
var BgpSrteSegmentSegmentType = struct {
	TYPE_A BgpSrteSegmentSegmentTypeEnum
	TYPE_B BgpSrteSegmentSegmentTypeEnum
	TYPE_C BgpSrteSegmentSegmentTypeEnum
	TYPE_D BgpSrteSegmentSegmentTypeEnum
	TYPE_E BgpSrteSegmentSegmentTypeEnum
	TYPE_F BgpSrteSegmentSegmentTypeEnum
	TYPE_G BgpSrteSegmentSegmentTypeEnum
	TYPE_H BgpSrteSegmentSegmentTypeEnum
	TYPE_I BgpSrteSegmentSegmentTypeEnum
	TYPE_J BgpSrteSegmentSegmentTypeEnum
	TYPE_K BgpSrteSegmentSegmentTypeEnum
}{
	TYPE_A: BgpSrteSegmentSegmentTypeEnum("type_a"),
	TYPE_B: BgpSrteSegmentSegmentTypeEnum("type_b"),
	TYPE_C: BgpSrteSegmentSegmentTypeEnum("type_c"),
	TYPE_D: BgpSrteSegmentSegmentTypeEnum("type_d"),
	TYPE_E: BgpSrteSegmentSegmentTypeEnum("type_e"),
	TYPE_F: BgpSrteSegmentSegmentTypeEnum("type_f"),
	TYPE_G: BgpSrteSegmentSegmentTypeEnum("type_g"),
	TYPE_H: BgpSrteSegmentSegmentTypeEnum("type_h"),
	TYPE_I: BgpSrteSegmentSegmentTypeEnum("type_i"),
	TYPE_J: BgpSrteSegmentSegmentTypeEnum("type_j"),
	TYPE_K: BgpSrteSegmentSegmentTypeEnum("type_k"),
}

func (obj *bgpSrteSegment) SegmentType() BgpSrteSegmentSegmentTypeEnum {
	return BgpSrteSegmentSegmentTypeEnum(obj.obj.SegmentType.Enum().String())
}

func (obj *bgpSrteSegment) SetSegmentType(value BgpSrteSegmentSegmentTypeEnum) BgpSrteSegment {
	intValue, ok := otg.BgpSrteSegment_SegmentType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpSrteSegmentSegmentTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpSrteSegment_SegmentType_Enum(intValue)
	obj.obj.SegmentType = &enumValue

	return obj
}

// description is TBD
// TypeA returns a BgpSrteSegmentATypeSubTlv
func (obj *bgpSrteSegment) TypeA() BgpSrteSegmentATypeSubTlv {
	if obj.obj.TypeA == nil {
		obj.obj.TypeA = NewBgpSrteSegmentATypeSubTlv().msg()
	}
	if obj.typeAHolder == nil {
		obj.typeAHolder = &bgpSrteSegmentATypeSubTlv{obj: obj.obj.TypeA}
	}
	return obj.typeAHolder
}

// description is TBD
// TypeA returns a BgpSrteSegmentATypeSubTlv
func (obj *bgpSrteSegment) HasTypeA() bool {
	return obj.obj.TypeA != nil
}

// description is TBD
// SetTypeA sets the BgpSrteSegmentATypeSubTlv value in the BgpSrteSegment object
func (obj *bgpSrteSegment) SetTypeA(value BgpSrteSegmentATypeSubTlv) BgpSrteSegment {

	obj.typeAHolder = nil
	obj.obj.TypeA = value.msg()

	return obj
}

// description is TBD
// TypeB returns a BgpSrteSegmentBTypeSubTlv
func (obj *bgpSrteSegment) TypeB() BgpSrteSegmentBTypeSubTlv {
	if obj.obj.TypeB == nil {
		obj.obj.TypeB = NewBgpSrteSegmentBTypeSubTlv().msg()
	}
	if obj.typeBHolder == nil {
		obj.typeBHolder = &bgpSrteSegmentBTypeSubTlv{obj: obj.obj.TypeB}
	}
	return obj.typeBHolder
}

// description is TBD
// TypeB returns a BgpSrteSegmentBTypeSubTlv
func (obj *bgpSrteSegment) HasTypeB() bool {
	return obj.obj.TypeB != nil
}

// description is TBD
// SetTypeB sets the BgpSrteSegmentBTypeSubTlv value in the BgpSrteSegment object
func (obj *bgpSrteSegment) SetTypeB(value BgpSrteSegmentBTypeSubTlv) BgpSrteSegment {

	obj.typeBHolder = nil
	obj.obj.TypeB = value.msg()

	return obj
}

// description is TBD
// TypeC returns a BgpSrteSegmentCTypeSubTlv
func (obj *bgpSrteSegment) TypeC() BgpSrteSegmentCTypeSubTlv {
	if obj.obj.TypeC == nil {
		obj.obj.TypeC = NewBgpSrteSegmentCTypeSubTlv().msg()
	}
	if obj.typeCHolder == nil {
		obj.typeCHolder = &bgpSrteSegmentCTypeSubTlv{obj: obj.obj.TypeC}
	}
	return obj.typeCHolder
}

// description is TBD
// TypeC returns a BgpSrteSegmentCTypeSubTlv
func (obj *bgpSrteSegment) HasTypeC() bool {
	return obj.obj.TypeC != nil
}

// description is TBD
// SetTypeC sets the BgpSrteSegmentCTypeSubTlv value in the BgpSrteSegment object
func (obj *bgpSrteSegment) SetTypeC(value BgpSrteSegmentCTypeSubTlv) BgpSrteSegment {

	obj.typeCHolder = nil
	obj.obj.TypeC = value.msg()

	return obj
}

// description is TBD
// TypeD returns a BgpSrteSegmentDTypeSubTlv
func (obj *bgpSrteSegment) TypeD() BgpSrteSegmentDTypeSubTlv {
	if obj.obj.TypeD == nil {
		obj.obj.TypeD = NewBgpSrteSegmentDTypeSubTlv().msg()
	}
	if obj.typeDHolder == nil {
		obj.typeDHolder = &bgpSrteSegmentDTypeSubTlv{obj: obj.obj.TypeD}
	}
	return obj.typeDHolder
}

// description is TBD
// TypeD returns a BgpSrteSegmentDTypeSubTlv
func (obj *bgpSrteSegment) HasTypeD() bool {
	return obj.obj.TypeD != nil
}

// description is TBD
// SetTypeD sets the BgpSrteSegmentDTypeSubTlv value in the BgpSrteSegment object
func (obj *bgpSrteSegment) SetTypeD(value BgpSrteSegmentDTypeSubTlv) BgpSrteSegment {

	obj.typeDHolder = nil
	obj.obj.TypeD = value.msg()

	return obj
}

// description is TBD
// TypeE returns a BgpSrteSegmentETypeSubTlv
func (obj *bgpSrteSegment) TypeE() BgpSrteSegmentETypeSubTlv {
	if obj.obj.TypeE == nil {
		obj.obj.TypeE = NewBgpSrteSegmentETypeSubTlv().msg()
	}
	if obj.typeEHolder == nil {
		obj.typeEHolder = &bgpSrteSegmentETypeSubTlv{obj: obj.obj.TypeE}
	}
	return obj.typeEHolder
}

// description is TBD
// TypeE returns a BgpSrteSegmentETypeSubTlv
func (obj *bgpSrteSegment) HasTypeE() bool {
	return obj.obj.TypeE != nil
}

// description is TBD
// SetTypeE sets the BgpSrteSegmentETypeSubTlv value in the BgpSrteSegment object
func (obj *bgpSrteSegment) SetTypeE(value BgpSrteSegmentETypeSubTlv) BgpSrteSegment {

	obj.typeEHolder = nil
	obj.obj.TypeE = value.msg()

	return obj
}

// description is TBD
// TypeF returns a BgpSrteSegmentFTypeSubTlv
func (obj *bgpSrteSegment) TypeF() BgpSrteSegmentFTypeSubTlv {
	if obj.obj.TypeF == nil {
		obj.obj.TypeF = NewBgpSrteSegmentFTypeSubTlv().msg()
	}
	if obj.typeFHolder == nil {
		obj.typeFHolder = &bgpSrteSegmentFTypeSubTlv{obj: obj.obj.TypeF}
	}
	return obj.typeFHolder
}

// description is TBD
// TypeF returns a BgpSrteSegmentFTypeSubTlv
func (obj *bgpSrteSegment) HasTypeF() bool {
	return obj.obj.TypeF != nil
}

// description is TBD
// SetTypeF sets the BgpSrteSegmentFTypeSubTlv value in the BgpSrteSegment object
func (obj *bgpSrteSegment) SetTypeF(value BgpSrteSegmentFTypeSubTlv) BgpSrteSegment {

	obj.typeFHolder = nil
	obj.obj.TypeF = value.msg()

	return obj
}

// description is TBD
// TypeG returns a BgpSrteSegmentGTypeSubTlv
func (obj *bgpSrteSegment) TypeG() BgpSrteSegmentGTypeSubTlv {
	if obj.obj.TypeG == nil {
		obj.obj.TypeG = NewBgpSrteSegmentGTypeSubTlv().msg()
	}
	if obj.typeGHolder == nil {
		obj.typeGHolder = &bgpSrteSegmentGTypeSubTlv{obj: obj.obj.TypeG}
	}
	return obj.typeGHolder
}

// description is TBD
// TypeG returns a BgpSrteSegmentGTypeSubTlv
func (obj *bgpSrteSegment) HasTypeG() bool {
	return obj.obj.TypeG != nil
}

// description is TBD
// SetTypeG sets the BgpSrteSegmentGTypeSubTlv value in the BgpSrteSegment object
func (obj *bgpSrteSegment) SetTypeG(value BgpSrteSegmentGTypeSubTlv) BgpSrteSegment {

	obj.typeGHolder = nil
	obj.obj.TypeG = value.msg()

	return obj
}

// description is TBD
// TypeH returns a BgpSrteSegmentHTypeSubTlv
func (obj *bgpSrteSegment) TypeH() BgpSrteSegmentHTypeSubTlv {
	if obj.obj.TypeH == nil {
		obj.obj.TypeH = NewBgpSrteSegmentHTypeSubTlv().msg()
	}
	if obj.typeHHolder == nil {
		obj.typeHHolder = &bgpSrteSegmentHTypeSubTlv{obj: obj.obj.TypeH}
	}
	return obj.typeHHolder
}

// description is TBD
// TypeH returns a BgpSrteSegmentHTypeSubTlv
func (obj *bgpSrteSegment) HasTypeH() bool {
	return obj.obj.TypeH != nil
}

// description is TBD
// SetTypeH sets the BgpSrteSegmentHTypeSubTlv value in the BgpSrteSegment object
func (obj *bgpSrteSegment) SetTypeH(value BgpSrteSegmentHTypeSubTlv) BgpSrteSegment {

	obj.typeHHolder = nil
	obj.obj.TypeH = value.msg()

	return obj
}

// description is TBD
// TypeI returns a BgpSrteSegmentITypeSubTlv
func (obj *bgpSrteSegment) TypeI() BgpSrteSegmentITypeSubTlv {
	if obj.obj.TypeI == nil {
		obj.obj.TypeI = NewBgpSrteSegmentITypeSubTlv().msg()
	}
	if obj.typeIHolder == nil {
		obj.typeIHolder = &bgpSrteSegmentITypeSubTlv{obj: obj.obj.TypeI}
	}
	return obj.typeIHolder
}

// description is TBD
// TypeI returns a BgpSrteSegmentITypeSubTlv
func (obj *bgpSrteSegment) HasTypeI() bool {
	return obj.obj.TypeI != nil
}

// description is TBD
// SetTypeI sets the BgpSrteSegmentITypeSubTlv value in the BgpSrteSegment object
func (obj *bgpSrteSegment) SetTypeI(value BgpSrteSegmentITypeSubTlv) BgpSrteSegment {

	obj.typeIHolder = nil
	obj.obj.TypeI = value.msg()

	return obj
}

// description is TBD
// TypeJ returns a BgpSrteSegmentJTypeSubTlv
func (obj *bgpSrteSegment) TypeJ() BgpSrteSegmentJTypeSubTlv {
	if obj.obj.TypeJ == nil {
		obj.obj.TypeJ = NewBgpSrteSegmentJTypeSubTlv().msg()
	}
	if obj.typeJHolder == nil {
		obj.typeJHolder = &bgpSrteSegmentJTypeSubTlv{obj: obj.obj.TypeJ}
	}
	return obj.typeJHolder
}

// description is TBD
// TypeJ returns a BgpSrteSegmentJTypeSubTlv
func (obj *bgpSrteSegment) HasTypeJ() bool {
	return obj.obj.TypeJ != nil
}

// description is TBD
// SetTypeJ sets the BgpSrteSegmentJTypeSubTlv value in the BgpSrteSegment object
func (obj *bgpSrteSegment) SetTypeJ(value BgpSrteSegmentJTypeSubTlv) BgpSrteSegment {

	obj.typeJHolder = nil
	obj.obj.TypeJ = value.msg()

	return obj
}

// description is TBD
// TypeK returns a BgpSrteSegmentKTypeSubTlv
func (obj *bgpSrteSegment) TypeK() BgpSrteSegmentKTypeSubTlv {
	if obj.obj.TypeK == nil {
		obj.obj.TypeK = NewBgpSrteSegmentKTypeSubTlv().msg()
	}
	if obj.typeKHolder == nil {
		obj.typeKHolder = &bgpSrteSegmentKTypeSubTlv{obj: obj.obj.TypeK}
	}
	return obj.typeKHolder
}

// description is TBD
// TypeK returns a BgpSrteSegmentKTypeSubTlv
func (obj *bgpSrteSegment) HasTypeK() bool {
	return obj.obj.TypeK != nil
}

// description is TBD
// SetTypeK sets the BgpSrteSegmentKTypeSubTlv value in the BgpSrteSegment object
func (obj *bgpSrteSegment) SetTypeK(value BgpSrteSegmentKTypeSubTlv) BgpSrteSegment {

	obj.typeKHolder = nil
	obj.obj.TypeK = value.msg()

	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *bgpSrteSegment) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the BgpSrteSegment object
func (obj *bgpSrteSegment) SetName(value string) BgpSrteSegment {

	obj.obj.Name = &value
	return obj
}

// If enabled means that this part of the configuration including any active 'children' nodes will be advertised to peer.  If disabled, this means that though config is present, it is not taking any part of the test but can be activated at run-time to advertise just this part of the configuration to the peer.
// Active returns a bool
func (obj *bgpSrteSegment) Active() bool {

	return *obj.obj.Active

}

// If enabled means that this part of the configuration including any active 'children' nodes will be advertised to peer.  If disabled, this means that though config is present, it is not taking any part of the test but can be activated at run-time to advertise just this part of the configuration to the peer.
// Active returns a bool
func (obj *bgpSrteSegment) HasActive() bool {
	return obj.obj.Active != nil
}

// If enabled means that this part of the configuration including any active 'children' nodes will be advertised to peer.  If disabled, this means that though config is present, it is not taking any part of the test but can be activated at run-time to advertise just this part of the configuration to the peer.
// SetActive sets the bool value in the BgpSrteSegment object
func (obj *bgpSrteSegment) SetActive(value bool) BgpSrteSegment {

	obj.obj.Active = &value
	return obj
}

func (obj *bgpSrteSegment) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// SegmentType is required
	if obj.obj.SegmentType == nil {
		vObj.validationErrors = append(vObj.validationErrors, "SegmentType is required field on interface BgpSrteSegment")
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

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface BgpSrteSegment")
	}
}

func (obj *bgpSrteSegment) setDefault() {
	if obj.obj.Active == nil {
		obj.SetActive(true)
	}

}
