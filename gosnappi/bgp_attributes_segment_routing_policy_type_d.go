package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesSegmentRoutingPolicyTypeD *****
type bgpAttributesSegmentRoutingPolicyTypeD struct {
	validation
	obj             *otg.BgpAttributesSegmentRoutingPolicyTypeD
	marshaller      marshalBgpAttributesSegmentRoutingPolicyTypeD
	unMarshaller    unMarshalBgpAttributesSegmentRoutingPolicyTypeD
	flagsHolder     BgpAttributesSegmentRoutingPolicyTypeFlags
	srMplsSidHolder BgpAttributesSidMpls
}

func NewBgpAttributesSegmentRoutingPolicyTypeD() BgpAttributesSegmentRoutingPolicyTypeD {
	obj := bgpAttributesSegmentRoutingPolicyTypeD{obj: &otg.BgpAttributesSegmentRoutingPolicyTypeD{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeD) msg() *otg.BgpAttributesSegmentRoutingPolicyTypeD {
	return obj.obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeD) setMsg(msg *otg.BgpAttributesSegmentRoutingPolicyTypeD) BgpAttributesSegmentRoutingPolicyTypeD {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesSegmentRoutingPolicyTypeD struct {
	obj *bgpAttributesSegmentRoutingPolicyTypeD
}

type marshalBgpAttributesSegmentRoutingPolicyTypeD interface {
	// ToProto marshals BgpAttributesSegmentRoutingPolicyTypeD to protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeD
	ToProto() (*otg.BgpAttributesSegmentRoutingPolicyTypeD, error)
	// ToPbText marshals BgpAttributesSegmentRoutingPolicyTypeD to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesSegmentRoutingPolicyTypeD to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesSegmentRoutingPolicyTypeD to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpAttributesSegmentRoutingPolicyTypeD to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpAttributesSegmentRoutingPolicyTypeD struct {
	obj *bgpAttributesSegmentRoutingPolicyTypeD
}

type unMarshalBgpAttributesSegmentRoutingPolicyTypeD interface {
	// FromProto unmarshals BgpAttributesSegmentRoutingPolicyTypeD from protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeD
	FromProto(msg *otg.BgpAttributesSegmentRoutingPolicyTypeD) (BgpAttributesSegmentRoutingPolicyTypeD, error)
	// FromPbText unmarshals BgpAttributesSegmentRoutingPolicyTypeD from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesSegmentRoutingPolicyTypeD from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesSegmentRoutingPolicyTypeD from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeD) Marshal() marshalBgpAttributesSegmentRoutingPolicyTypeD {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesSegmentRoutingPolicyTypeD{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeD) Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicyTypeD {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesSegmentRoutingPolicyTypeD{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeD) ToProto() (*otg.BgpAttributesSegmentRoutingPolicyTypeD, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeD) FromProto(msg *otg.BgpAttributesSegmentRoutingPolicyTypeD) (BgpAttributesSegmentRoutingPolicyTypeD, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeD) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeD) FromPbText(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeD) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeD) FromYaml(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeD) ToJsonRaw() (string, error) {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeD) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeD) FromJson(value string) error {
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

func (obj *bgpAttributesSegmentRoutingPolicyTypeD) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeD) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeD) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeD) Clone() (BgpAttributesSegmentRoutingPolicyTypeD, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesSegmentRoutingPolicyTypeD()
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

func (obj *bgpAttributesSegmentRoutingPolicyTypeD) setNil() {
	obj.flagsHolder = nil
	obj.srMplsSidHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpAttributesSegmentRoutingPolicyTypeD is type D: IPv6 Node Address with optional SID for SR MPLS.
// It is encoded as a Segment of Type 4 in the SEGMENT_LIST sub-tlv.
type BgpAttributesSegmentRoutingPolicyTypeD interface {
	Validation
	// msg marshals BgpAttributesSegmentRoutingPolicyTypeD to protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeD
	// and doesn't set defaults
	msg() *otg.BgpAttributesSegmentRoutingPolicyTypeD
	// setMsg unmarshals BgpAttributesSegmentRoutingPolicyTypeD from protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeD
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesSegmentRoutingPolicyTypeD) BgpAttributesSegmentRoutingPolicyTypeD
	// provides marshal interface
	Marshal() marshalBgpAttributesSegmentRoutingPolicyTypeD
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicyTypeD
	// validate validates BgpAttributesSegmentRoutingPolicyTypeD
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesSegmentRoutingPolicyTypeD, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns BgpAttributesSegmentRoutingPolicyTypeFlags, set in BgpAttributesSegmentRoutingPolicyTypeD.
	// BgpAttributesSegmentRoutingPolicyTypeFlags is flags for each Segment in SEGMENT_LIST sub-tlv.
	// - V-flag. Indicates verification is enabled. See section 5, of https://datatracker.ietf.org/doc/html/rfc9256
	// - A-flag. Indicates presence of SR Algorithm field applicable to Segment Types C, D , I , J and K .
	// - B-Flag. Indicates presence of SRv6 Endpoint Behavior and SID Structure encoding applicable to Segment Types B , I , J and K .
	// - S-Flag: This flag, when set, indicates the presence of the SR-MPLS or SRv6 SID depending on the segment type. (draft-ietf-idr-bgp-sr-segtypes-ext-03 Section 2.10).
	// This flag is applicable for Segment Types C, D, E, F, G, H, I, J, and K.
	Flags() BgpAttributesSegmentRoutingPolicyTypeFlags
	// SetFlags assigns BgpAttributesSegmentRoutingPolicyTypeFlags provided by user to BgpAttributesSegmentRoutingPolicyTypeD.
	// BgpAttributesSegmentRoutingPolicyTypeFlags is flags for each Segment in SEGMENT_LIST sub-tlv.
	// - V-flag. Indicates verification is enabled. See section 5, of https://datatracker.ietf.org/doc/html/rfc9256
	// - A-flag. Indicates presence of SR Algorithm field applicable to Segment Types C, D , I , J and K .
	// - B-Flag. Indicates presence of SRv6 Endpoint Behavior and SID Structure encoding applicable to Segment Types B , I , J and K .
	// - S-Flag: This flag, when set, indicates the presence of the SR-MPLS or SRv6 SID depending on the segment type. (draft-ietf-idr-bgp-sr-segtypes-ext-03 Section 2.10).
	// This flag is applicable for Segment Types C, D, E, F, G, H, I, J, and K.
	SetFlags(value BgpAttributesSegmentRoutingPolicyTypeFlags) BgpAttributesSegmentRoutingPolicyTypeD
	// HasFlags checks if Flags has been set in BgpAttributesSegmentRoutingPolicyTypeD
	HasFlags() bool
	// SrAlgorithm returns uint32, set in BgpAttributesSegmentRoutingPolicyTypeD.
	SrAlgorithm() uint32
	// SetSrAlgorithm assigns uint32 provided by user to BgpAttributesSegmentRoutingPolicyTypeD
	SetSrAlgorithm(value uint32) BgpAttributesSegmentRoutingPolicyTypeD
	// HasSrAlgorithm checks if SrAlgorithm has been set in BgpAttributesSegmentRoutingPolicyTypeD
	HasSrAlgorithm() bool
	// Ipv6NodeAddress returns string, set in BgpAttributesSegmentRoutingPolicyTypeD.
	Ipv6NodeAddress() string
	// SetIpv6NodeAddress assigns string provided by user to BgpAttributesSegmentRoutingPolicyTypeD
	SetIpv6NodeAddress(value string) BgpAttributesSegmentRoutingPolicyTypeD
	// HasIpv6NodeAddress checks if Ipv6NodeAddress has been set in BgpAttributesSegmentRoutingPolicyTypeD
	HasIpv6NodeAddress() bool
	// SrMplsSid returns BgpAttributesSidMpls, set in BgpAttributesSegmentRoutingPolicyTypeD.
	// BgpAttributesSidMpls is this carries a 20 bit Multi Protocol Label Switching alongwith 3 bits traffic class, 1 bit indicating presence
	// or absence of Bottom-Of-Stack and 8 bits carrying the Time to Live value.
	SrMplsSid() BgpAttributesSidMpls
	// SetSrMplsSid assigns BgpAttributesSidMpls provided by user to BgpAttributesSegmentRoutingPolicyTypeD.
	// BgpAttributesSidMpls is this carries a 20 bit Multi Protocol Label Switching alongwith 3 bits traffic class, 1 bit indicating presence
	// or absence of Bottom-Of-Stack and 8 bits carrying the Time to Live value.
	SetSrMplsSid(value BgpAttributesSidMpls) BgpAttributesSegmentRoutingPolicyTypeD
	// HasSrMplsSid checks if SrMplsSid has been set in BgpAttributesSegmentRoutingPolicyTypeD
	HasSrMplsSid() bool
	setNil()
}

// description is TBD
// Flags returns a BgpAttributesSegmentRoutingPolicyTypeFlags
func (obj *bgpAttributesSegmentRoutingPolicyTypeD) Flags() BgpAttributesSegmentRoutingPolicyTypeFlags {
	if obj.obj.Flags == nil {
		obj.obj.Flags = NewBgpAttributesSegmentRoutingPolicyTypeFlags().msg()
	}
	if obj.flagsHolder == nil {
		obj.flagsHolder = &bgpAttributesSegmentRoutingPolicyTypeFlags{obj: obj.obj.Flags}
	}
	return obj.flagsHolder
}

// description is TBD
// Flags returns a BgpAttributesSegmentRoutingPolicyTypeFlags
func (obj *bgpAttributesSegmentRoutingPolicyTypeD) HasFlags() bool {
	return obj.obj.Flags != nil
}

// description is TBD
// SetFlags sets the BgpAttributesSegmentRoutingPolicyTypeFlags value in the BgpAttributesSegmentRoutingPolicyTypeD object
func (obj *bgpAttributesSegmentRoutingPolicyTypeD) SetFlags(value BgpAttributesSegmentRoutingPolicyTypeFlags) BgpAttributesSegmentRoutingPolicyTypeD {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// SR Algorithm identifier when A-Flag in on. If A-flag is not enabled, it should be set to 0 on transmission and ignored on receipt.
// SrAlgorithm returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicyTypeD) SrAlgorithm() uint32 {

	return *obj.obj.SrAlgorithm

}

// SR Algorithm identifier when A-Flag in on. If A-flag is not enabled, it should be set to 0 on transmission and ignored on receipt.
// SrAlgorithm returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicyTypeD) HasSrAlgorithm() bool {
	return obj.obj.SrAlgorithm != nil
}

// SR Algorithm identifier when A-Flag in on. If A-flag is not enabled, it should be set to 0 on transmission and ignored on receipt.
// SetSrAlgorithm sets the uint32 value in the BgpAttributesSegmentRoutingPolicyTypeD object
func (obj *bgpAttributesSegmentRoutingPolicyTypeD) SetSrAlgorithm(value uint32) BgpAttributesSegmentRoutingPolicyTypeD {

	obj.obj.SrAlgorithm = &value
	return obj
}

// IPv6 address representing a node.
// Ipv6NodeAddress returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeD) Ipv6NodeAddress() string {

	return *obj.obj.Ipv6NodeAddress

}

// IPv6 address representing a node.
// Ipv6NodeAddress returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeD) HasIpv6NodeAddress() bool {
	return obj.obj.Ipv6NodeAddress != nil
}

// IPv6 address representing a node.
// SetIpv6NodeAddress sets the string value in the BgpAttributesSegmentRoutingPolicyTypeD object
func (obj *bgpAttributesSegmentRoutingPolicyTypeD) SetIpv6NodeAddress(value string) BgpAttributesSegmentRoutingPolicyTypeD {

	obj.obj.Ipv6NodeAddress = &value
	return obj
}

// Optional SR-MPLS SID.
// SrMplsSid returns a BgpAttributesSidMpls
func (obj *bgpAttributesSegmentRoutingPolicyTypeD) SrMplsSid() BgpAttributesSidMpls {
	if obj.obj.SrMplsSid == nil {
		obj.obj.SrMplsSid = NewBgpAttributesSidMpls().msg()
	}
	if obj.srMplsSidHolder == nil {
		obj.srMplsSidHolder = &bgpAttributesSidMpls{obj: obj.obj.SrMplsSid}
	}
	return obj.srMplsSidHolder
}

// Optional SR-MPLS SID.
// SrMplsSid returns a BgpAttributesSidMpls
func (obj *bgpAttributesSegmentRoutingPolicyTypeD) HasSrMplsSid() bool {
	return obj.obj.SrMplsSid != nil
}

// Optional SR-MPLS SID.
// SetSrMplsSid sets the BgpAttributesSidMpls value in the BgpAttributesSegmentRoutingPolicyTypeD object
func (obj *bgpAttributesSegmentRoutingPolicyTypeD) SetSrMplsSid(value BgpAttributesSidMpls) BgpAttributesSegmentRoutingPolicyTypeD {

	obj.srMplsSidHolder = nil
	obj.obj.SrMplsSid = value.msg()

	return obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeD) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(vObj, set_default)
	}

	if obj.obj.SrAlgorithm != nil {

		if *obj.obj.SrAlgorithm > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpAttributesSegmentRoutingPolicyTypeD.SrAlgorithm <= 255 but Got %d", *obj.obj.SrAlgorithm))
		}

	}

	if obj.obj.Ipv6NodeAddress != nil {

		err := obj.validateIpv6(obj.Ipv6NodeAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpAttributesSegmentRoutingPolicyTypeD.Ipv6NodeAddress"))
		}

	}

	if obj.obj.SrMplsSid != nil {

		obj.SrMplsSid().validateObj(vObj, set_default)
	}

}

func (obj *bgpAttributesSegmentRoutingPolicyTypeD) setDefault() {
	if obj.obj.SrAlgorithm == nil {
		obj.SetSrAlgorithm(0)
	}
	if obj.obj.Ipv6NodeAddress == nil {
		obj.SetIpv6NodeAddress("0::0")
	}

}
