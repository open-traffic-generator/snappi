package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesSegmentRoutingPolicyTypeC *****
type bgpAttributesSegmentRoutingPolicyTypeC struct {
	validation
	obj             *otg.BgpAttributesSegmentRoutingPolicyTypeC
	marshaller      marshalBgpAttributesSegmentRoutingPolicyTypeC
	unMarshaller    unMarshalBgpAttributesSegmentRoutingPolicyTypeC
	flagsHolder     BgpAttributesSegmentRoutingPolicyTypeFlags
	srMplsSidHolder BgpAttributesSidMpls
}

func NewBgpAttributesSegmentRoutingPolicyTypeC() BgpAttributesSegmentRoutingPolicyTypeC {
	obj := bgpAttributesSegmentRoutingPolicyTypeC{obj: &otg.BgpAttributesSegmentRoutingPolicyTypeC{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeC) msg() *otg.BgpAttributesSegmentRoutingPolicyTypeC {
	return obj.obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeC) setMsg(msg *otg.BgpAttributesSegmentRoutingPolicyTypeC) BgpAttributesSegmentRoutingPolicyTypeC {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesSegmentRoutingPolicyTypeC struct {
	obj *bgpAttributesSegmentRoutingPolicyTypeC
}

type marshalBgpAttributesSegmentRoutingPolicyTypeC interface {
	// ToProto marshals BgpAttributesSegmentRoutingPolicyTypeC to protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeC
	ToProto() (*otg.BgpAttributesSegmentRoutingPolicyTypeC, error)
	// ToPbText marshals BgpAttributesSegmentRoutingPolicyTypeC to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesSegmentRoutingPolicyTypeC to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesSegmentRoutingPolicyTypeC to JSON text
	ToJson() (string, error)
}

type unMarshalbgpAttributesSegmentRoutingPolicyTypeC struct {
	obj *bgpAttributesSegmentRoutingPolicyTypeC
}

type unMarshalBgpAttributesSegmentRoutingPolicyTypeC interface {
	// FromProto unmarshals BgpAttributesSegmentRoutingPolicyTypeC from protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeC
	FromProto(msg *otg.BgpAttributesSegmentRoutingPolicyTypeC) (BgpAttributesSegmentRoutingPolicyTypeC, error)
	// FromPbText unmarshals BgpAttributesSegmentRoutingPolicyTypeC from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesSegmentRoutingPolicyTypeC from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesSegmentRoutingPolicyTypeC from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeC) Marshal() marshalBgpAttributesSegmentRoutingPolicyTypeC {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesSegmentRoutingPolicyTypeC{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeC) Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicyTypeC {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesSegmentRoutingPolicyTypeC{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeC) ToProto() (*otg.BgpAttributesSegmentRoutingPolicyTypeC, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeC) FromProto(msg *otg.BgpAttributesSegmentRoutingPolicyTypeC) (BgpAttributesSegmentRoutingPolicyTypeC, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeC) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeC) FromPbText(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeC) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeC) FromYaml(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeC) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeC) FromJson(value string) error {
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

func (obj *bgpAttributesSegmentRoutingPolicyTypeC) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeC) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeC) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeC) Clone() (BgpAttributesSegmentRoutingPolicyTypeC, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesSegmentRoutingPolicyTypeC()
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

func (obj *bgpAttributesSegmentRoutingPolicyTypeC) setNil() {
	obj.flagsHolder = nil
	obj.srMplsSidHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpAttributesSegmentRoutingPolicyTypeC is type C: IPv4 Node Address with optional SID.
// It is encoded as a Segment of Type 3 in the SEGMENT_LIST sub-tlv.
type BgpAttributesSegmentRoutingPolicyTypeC interface {
	Validation
	// msg marshals BgpAttributesSegmentRoutingPolicyTypeC to protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeC
	// and doesn't set defaults
	msg() *otg.BgpAttributesSegmentRoutingPolicyTypeC
	// setMsg unmarshals BgpAttributesSegmentRoutingPolicyTypeC from protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeC
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesSegmentRoutingPolicyTypeC) BgpAttributesSegmentRoutingPolicyTypeC
	// provides marshal interface
	Marshal() marshalBgpAttributesSegmentRoutingPolicyTypeC
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicyTypeC
	// validate validates BgpAttributesSegmentRoutingPolicyTypeC
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesSegmentRoutingPolicyTypeC, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns BgpAttributesSegmentRoutingPolicyTypeFlags, set in BgpAttributesSegmentRoutingPolicyTypeC.
	// BgpAttributesSegmentRoutingPolicyTypeFlags is flags for each Segment in SEGMENT_LIST sub-tlv.
	// - V-flag. Indicates verification is enabled. See section 5, of https://datatracker.ietf.org/doc/html/rfc9256
	// - A-flag. Indicates presence of SR Algorithm field applicable to Segment Types C, D , I , J and K .
	// - B-Flag. Indicates presence of SRv6 Endpoint Behavior and SID Structure encoding applicable to Segment Types B , I , J and K .
	// - S-Flag: This flag, when set, indicates the presence of the SR-MPLS or SRv6 SID depending on the segment type. (draft-ietf-idr-bgp-sr-segtypes-ext-03 Section 2.10).
	// This flag is applicable for Segment Types C, D, E, F, G, H, I, J, and K.
	Flags() BgpAttributesSegmentRoutingPolicyTypeFlags
	// SetFlags assigns BgpAttributesSegmentRoutingPolicyTypeFlags provided by user to BgpAttributesSegmentRoutingPolicyTypeC.
	// BgpAttributesSegmentRoutingPolicyTypeFlags is flags for each Segment in SEGMENT_LIST sub-tlv.
	// - V-flag. Indicates verification is enabled. See section 5, of https://datatracker.ietf.org/doc/html/rfc9256
	// - A-flag. Indicates presence of SR Algorithm field applicable to Segment Types C, D , I , J and K .
	// - B-Flag. Indicates presence of SRv6 Endpoint Behavior and SID Structure encoding applicable to Segment Types B , I , J and K .
	// - S-Flag: This flag, when set, indicates the presence of the SR-MPLS or SRv6 SID depending on the segment type. (draft-ietf-idr-bgp-sr-segtypes-ext-03 Section 2.10).
	// This flag is applicable for Segment Types C, D, E, F, G, H, I, J, and K.
	SetFlags(value BgpAttributesSegmentRoutingPolicyTypeFlags) BgpAttributesSegmentRoutingPolicyTypeC
	// HasFlags checks if Flags has been set in BgpAttributesSegmentRoutingPolicyTypeC
	HasFlags() bool
	// SrAlgorithm returns uint32, set in BgpAttributesSegmentRoutingPolicyTypeC.
	SrAlgorithm() uint32
	// SetSrAlgorithm assigns uint32 provided by user to BgpAttributesSegmentRoutingPolicyTypeC
	SetSrAlgorithm(value uint32) BgpAttributesSegmentRoutingPolicyTypeC
	// HasSrAlgorithm checks if SrAlgorithm has been set in BgpAttributesSegmentRoutingPolicyTypeC
	HasSrAlgorithm() bool
	// Ipv4NodeAddress returns string, set in BgpAttributesSegmentRoutingPolicyTypeC.
	Ipv4NodeAddress() string
	// SetIpv4NodeAddress assigns string provided by user to BgpAttributesSegmentRoutingPolicyTypeC
	SetIpv4NodeAddress(value string) BgpAttributesSegmentRoutingPolicyTypeC
	// HasIpv4NodeAddress checks if Ipv4NodeAddress has been set in BgpAttributesSegmentRoutingPolicyTypeC
	HasIpv4NodeAddress() bool
	// SrMplsSid returns BgpAttributesSidMpls, set in BgpAttributesSegmentRoutingPolicyTypeC.
	// BgpAttributesSidMpls is this carries a 20 bit Multi Protocol Label Switching alongwith 3 bits traffic class, 1 bit indicating presence
	// or absence of Bottom-Of-Stack and 8 bits carrying the Time to Live value.
	SrMplsSid() BgpAttributesSidMpls
	// SetSrMplsSid assigns BgpAttributesSidMpls provided by user to BgpAttributesSegmentRoutingPolicyTypeC.
	// BgpAttributesSidMpls is this carries a 20 bit Multi Protocol Label Switching alongwith 3 bits traffic class, 1 bit indicating presence
	// or absence of Bottom-Of-Stack and 8 bits carrying the Time to Live value.
	SetSrMplsSid(value BgpAttributesSidMpls) BgpAttributesSegmentRoutingPolicyTypeC
	// HasSrMplsSid checks if SrMplsSid has been set in BgpAttributesSegmentRoutingPolicyTypeC
	HasSrMplsSid() bool
	setNil()
}

// description is TBD
// Flags returns a BgpAttributesSegmentRoutingPolicyTypeFlags
func (obj *bgpAttributesSegmentRoutingPolicyTypeC) Flags() BgpAttributesSegmentRoutingPolicyTypeFlags {
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
func (obj *bgpAttributesSegmentRoutingPolicyTypeC) HasFlags() bool {
	return obj.obj.Flags != nil
}

// description is TBD
// SetFlags sets the BgpAttributesSegmentRoutingPolicyTypeFlags value in the BgpAttributesSegmentRoutingPolicyTypeC object
func (obj *bgpAttributesSegmentRoutingPolicyTypeC) SetFlags(value BgpAttributesSegmentRoutingPolicyTypeFlags) BgpAttributesSegmentRoutingPolicyTypeC {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// SR Algorithm identifier when A-Flag in on. If A-flag is not enabled, it should be set to 0 on transmission and ignored on receipt.
// SrAlgorithm returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicyTypeC) SrAlgorithm() uint32 {

	return *obj.obj.SrAlgorithm

}

// SR Algorithm identifier when A-Flag in on. If A-flag is not enabled, it should be set to 0 on transmission and ignored on receipt.
// SrAlgorithm returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicyTypeC) HasSrAlgorithm() bool {
	return obj.obj.SrAlgorithm != nil
}

// SR Algorithm identifier when A-Flag in on. If A-flag is not enabled, it should be set to 0 on transmission and ignored on receipt.
// SetSrAlgorithm sets the uint32 value in the BgpAttributesSegmentRoutingPolicyTypeC object
func (obj *bgpAttributesSegmentRoutingPolicyTypeC) SetSrAlgorithm(value uint32) BgpAttributesSegmentRoutingPolicyTypeC {

	obj.obj.SrAlgorithm = &value
	return obj
}

// IPv4 address representing a node.
// Ipv4NodeAddress returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeC) Ipv4NodeAddress() string {

	return *obj.obj.Ipv4NodeAddress

}

// IPv4 address representing a node.
// Ipv4NodeAddress returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeC) HasIpv4NodeAddress() bool {
	return obj.obj.Ipv4NodeAddress != nil
}

// IPv4 address representing a node.
// SetIpv4NodeAddress sets the string value in the BgpAttributesSegmentRoutingPolicyTypeC object
func (obj *bgpAttributesSegmentRoutingPolicyTypeC) SetIpv4NodeAddress(value string) BgpAttributesSegmentRoutingPolicyTypeC {

	obj.obj.Ipv4NodeAddress = &value
	return obj
}

// Optional SR-MPLS SID.
// SrMplsSid returns a BgpAttributesSidMpls
func (obj *bgpAttributesSegmentRoutingPolicyTypeC) SrMplsSid() BgpAttributesSidMpls {
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
func (obj *bgpAttributesSegmentRoutingPolicyTypeC) HasSrMplsSid() bool {
	return obj.obj.SrMplsSid != nil
}

// Optional SR-MPLS SID.
// SetSrMplsSid sets the BgpAttributesSidMpls value in the BgpAttributesSegmentRoutingPolicyTypeC object
func (obj *bgpAttributesSegmentRoutingPolicyTypeC) SetSrMplsSid(value BgpAttributesSidMpls) BgpAttributesSegmentRoutingPolicyTypeC {

	obj.srMplsSidHolder = nil
	obj.obj.SrMplsSid = value.msg()

	return obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeC) validateObj(vObj *validation, set_default bool) {
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
				fmt.Sprintf("0 <= BgpAttributesSegmentRoutingPolicyTypeC.SrAlgorithm <= 255 but Got %d", *obj.obj.SrAlgorithm))
		}

	}

	if obj.obj.Ipv4NodeAddress != nil {

		err := obj.validateIpv4(obj.Ipv4NodeAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpAttributesSegmentRoutingPolicyTypeC.Ipv4NodeAddress"))
		}

	}

	if obj.obj.SrMplsSid != nil {

		obj.SrMplsSid().validateObj(vObj, set_default)
	}

}

func (obj *bgpAttributesSegmentRoutingPolicyTypeC) setDefault() {
	if obj.obj.SrAlgorithm == nil {
		obj.SetSrAlgorithm(0)
	}
	if obj.obj.Ipv4NodeAddress == nil {
		obj.SetIpv4NodeAddress("0.0.0.0")
	}

}
