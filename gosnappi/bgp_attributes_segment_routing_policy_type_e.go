package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesSegmentRoutingPolicyTypeE *****
type bgpAttributesSegmentRoutingPolicyTypeE struct {
	validation
	obj             *otg.BgpAttributesSegmentRoutingPolicyTypeE
	marshaller      marshalBgpAttributesSegmentRoutingPolicyTypeE
	unMarshaller    unMarshalBgpAttributesSegmentRoutingPolicyTypeE
	flagsHolder     BgpAttributesSegmentRoutingPolicyTypeFlags
	srMplsSidHolder BgpAttributesSidMpls
}

func NewBgpAttributesSegmentRoutingPolicyTypeE() BgpAttributesSegmentRoutingPolicyTypeE {
	obj := bgpAttributesSegmentRoutingPolicyTypeE{obj: &otg.BgpAttributesSegmentRoutingPolicyTypeE{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeE) msg() *otg.BgpAttributesSegmentRoutingPolicyTypeE {
	return obj.obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeE) setMsg(msg *otg.BgpAttributesSegmentRoutingPolicyTypeE) BgpAttributesSegmentRoutingPolicyTypeE {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesSegmentRoutingPolicyTypeE struct {
	obj *bgpAttributesSegmentRoutingPolicyTypeE
}

type marshalBgpAttributesSegmentRoutingPolicyTypeE interface {
	// ToProto marshals BgpAttributesSegmentRoutingPolicyTypeE to protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeE
	ToProto() (*otg.BgpAttributesSegmentRoutingPolicyTypeE, error)
	// ToPbText marshals BgpAttributesSegmentRoutingPolicyTypeE to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesSegmentRoutingPolicyTypeE to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesSegmentRoutingPolicyTypeE to JSON text
	ToJson() (string, error)
}

type unMarshalbgpAttributesSegmentRoutingPolicyTypeE struct {
	obj *bgpAttributesSegmentRoutingPolicyTypeE
}

type unMarshalBgpAttributesSegmentRoutingPolicyTypeE interface {
	// FromProto unmarshals BgpAttributesSegmentRoutingPolicyTypeE from protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeE
	FromProto(msg *otg.BgpAttributesSegmentRoutingPolicyTypeE) (BgpAttributesSegmentRoutingPolicyTypeE, error)
	// FromPbText unmarshals BgpAttributesSegmentRoutingPolicyTypeE from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesSegmentRoutingPolicyTypeE from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesSegmentRoutingPolicyTypeE from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeE) Marshal() marshalBgpAttributesSegmentRoutingPolicyTypeE {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesSegmentRoutingPolicyTypeE{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeE) Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicyTypeE {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesSegmentRoutingPolicyTypeE{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeE) ToProto() (*otg.BgpAttributesSegmentRoutingPolicyTypeE, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeE) FromProto(msg *otg.BgpAttributesSegmentRoutingPolicyTypeE) (BgpAttributesSegmentRoutingPolicyTypeE, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeE) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeE) FromPbText(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeE) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeE) FromYaml(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeE) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeE) FromJson(value string) error {
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

func (obj *bgpAttributesSegmentRoutingPolicyTypeE) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeE) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeE) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeE) Clone() (BgpAttributesSegmentRoutingPolicyTypeE, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesSegmentRoutingPolicyTypeE()
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

func (obj *bgpAttributesSegmentRoutingPolicyTypeE) setNil() {
	obj.flagsHolder = nil
	obj.srMplsSidHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpAttributesSegmentRoutingPolicyTypeE is type E: IPv4 Address and Local Interface ID with optional SID
// It is encoded as a Segment of Type 5 in the SEGMENT_LIST sub-tlv.
type BgpAttributesSegmentRoutingPolicyTypeE interface {
	Validation
	// msg marshals BgpAttributesSegmentRoutingPolicyTypeE to protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeE
	// and doesn't set defaults
	msg() *otg.BgpAttributesSegmentRoutingPolicyTypeE
	// setMsg unmarshals BgpAttributesSegmentRoutingPolicyTypeE from protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeE
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesSegmentRoutingPolicyTypeE) BgpAttributesSegmentRoutingPolicyTypeE
	// provides marshal interface
	Marshal() marshalBgpAttributesSegmentRoutingPolicyTypeE
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicyTypeE
	// validate validates BgpAttributesSegmentRoutingPolicyTypeE
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesSegmentRoutingPolicyTypeE, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns BgpAttributesSegmentRoutingPolicyTypeFlags, set in BgpAttributesSegmentRoutingPolicyTypeE.
	// BgpAttributesSegmentRoutingPolicyTypeFlags is flags for each Segment in SEGMENT_LIST sub-tlv.
	// - V-flag. Indicates verification is enabled. See section 5, of https://datatracker.ietf.org/doc/html/rfc9256
	// - A-flag. Indicates presence of SR Algorithm field applicable to Segment Types C, D , I , J and K .
	// - B-Flag. Indicates presence of SRv6 Endpoint Behavior and SID Structure encoding applicable to Segment Types B , I , J and K .
	// - S-Flag: This flag, when set, indicates the presence of the SR-MPLS or SRv6 SID depending on the segment type. (draft-ietf-idr-bgp-sr-segtypes-ext-03 Section 2.10).
	// This flag is applicable for Segment Types C, D, E, F, G, H, I, J, and K.
	Flags() BgpAttributesSegmentRoutingPolicyTypeFlags
	// SetFlags assigns BgpAttributesSegmentRoutingPolicyTypeFlags provided by user to BgpAttributesSegmentRoutingPolicyTypeE.
	// BgpAttributesSegmentRoutingPolicyTypeFlags is flags for each Segment in SEGMENT_LIST sub-tlv.
	// - V-flag. Indicates verification is enabled. See section 5, of https://datatracker.ietf.org/doc/html/rfc9256
	// - A-flag. Indicates presence of SR Algorithm field applicable to Segment Types C, D , I , J and K .
	// - B-Flag. Indicates presence of SRv6 Endpoint Behavior and SID Structure encoding applicable to Segment Types B , I , J and K .
	// - S-Flag: This flag, when set, indicates the presence of the SR-MPLS or SRv6 SID depending on the segment type. (draft-ietf-idr-bgp-sr-segtypes-ext-03 Section 2.10).
	// This flag is applicable for Segment Types C, D, E, F, G, H, I, J, and K.
	SetFlags(value BgpAttributesSegmentRoutingPolicyTypeFlags) BgpAttributesSegmentRoutingPolicyTypeE
	// HasFlags checks if Flags has been set in BgpAttributesSegmentRoutingPolicyTypeE
	HasFlags() bool
	// LocalInterfaceId returns uint32, set in BgpAttributesSegmentRoutingPolicyTypeE.
	LocalInterfaceId() uint32
	// SetLocalInterfaceId assigns uint32 provided by user to BgpAttributesSegmentRoutingPolicyTypeE
	SetLocalInterfaceId(value uint32) BgpAttributesSegmentRoutingPolicyTypeE
	// HasLocalInterfaceId checks if LocalInterfaceId has been set in BgpAttributesSegmentRoutingPolicyTypeE
	HasLocalInterfaceId() bool
	// Ipv4NodeAddress returns string, set in BgpAttributesSegmentRoutingPolicyTypeE.
	Ipv4NodeAddress() string
	// SetIpv4NodeAddress assigns string provided by user to BgpAttributesSegmentRoutingPolicyTypeE
	SetIpv4NodeAddress(value string) BgpAttributesSegmentRoutingPolicyTypeE
	// HasIpv4NodeAddress checks if Ipv4NodeAddress has been set in BgpAttributesSegmentRoutingPolicyTypeE
	HasIpv4NodeAddress() bool
	// SrMplsSid returns BgpAttributesSidMpls, set in BgpAttributesSegmentRoutingPolicyTypeE.
	// BgpAttributesSidMpls is this carries a 20 bit Multi Protocol Label Switching alongwith 3 bits traffic class, 1 bit indicating presence
	// or absence of Bottom-Of-Stack and 8 bits carrying the Time to Live value.
	SrMplsSid() BgpAttributesSidMpls
	// SetSrMplsSid assigns BgpAttributesSidMpls provided by user to BgpAttributesSegmentRoutingPolicyTypeE.
	// BgpAttributesSidMpls is this carries a 20 bit Multi Protocol Label Switching alongwith 3 bits traffic class, 1 bit indicating presence
	// or absence of Bottom-Of-Stack and 8 bits carrying the Time to Live value.
	SetSrMplsSid(value BgpAttributesSidMpls) BgpAttributesSegmentRoutingPolicyTypeE
	// HasSrMplsSid checks if SrMplsSid has been set in BgpAttributesSegmentRoutingPolicyTypeE
	HasSrMplsSid() bool
	setNil()
}

// description is TBD
// Flags returns a BgpAttributesSegmentRoutingPolicyTypeFlags
func (obj *bgpAttributesSegmentRoutingPolicyTypeE) Flags() BgpAttributesSegmentRoutingPolicyTypeFlags {
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
func (obj *bgpAttributesSegmentRoutingPolicyTypeE) HasFlags() bool {
	return obj.obj.Flags != nil
}

// description is TBD
// SetFlags sets the BgpAttributesSegmentRoutingPolicyTypeFlags value in the BgpAttributesSegmentRoutingPolicyTypeE object
func (obj *bgpAttributesSegmentRoutingPolicyTypeE) SetFlags(value BgpAttributesSegmentRoutingPolicyTypeFlags) BgpAttributesSegmentRoutingPolicyTypeE {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// The Interface Index as defined in [RFC8664].
// LocalInterfaceId returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicyTypeE) LocalInterfaceId() uint32 {

	return *obj.obj.LocalInterfaceId

}

// The Interface Index as defined in [RFC8664].
// LocalInterfaceId returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicyTypeE) HasLocalInterfaceId() bool {
	return obj.obj.LocalInterfaceId != nil
}

// The Interface Index as defined in [RFC8664].
// SetLocalInterfaceId sets the uint32 value in the BgpAttributesSegmentRoutingPolicyTypeE object
func (obj *bgpAttributesSegmentRoutingPolicyTypeE) SetLocalInterfaceId(value uint32) BgpAttributesSegmentRoutingPolicyTypeE {

	obj.obj.LocalInterfaceId = &value
	return obj
}

// IPv4 address representing a node.
// Ipv4NodeAddress returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeE) Ipv4NodeAddress() string {

	return *obj.obj.Ipv4NodeAddress

}

// IPv4 address representing a node.
// Ipv4NodeAddress returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeE) HasIpv4NodeAddress() bool {
	return obj.obj.Ipv4NodeAddress != nil
}

// IPv4 address representing a node.
// SetIpv4NodeAddress sets the string value in the BgpAttributesSegmentRoutingPolicyTypeE object
func (obj *bgpAttributesSegmentRoutingPolicyTypeE) SetIpv4NodeAddress(value string) BgpAttributesSegmentRoutingPolicyTypeE {

	obj.obj.Ipv4NodeAddress = &value
	return obj
}

// Optional SR-MPLS SID.
// SrMplsSid returns a BgpAttributesSidMpls
func (obj *bgpAttributesSegmentRoutingPolicyTypeE) SrMplsSid() BgpAttributesSidMpls {
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
func (obj *bgpAttributesSegmentRoutingPolicyTypeE) HasSrMplsSid() bool {
	return obj.obj.SrMplsSid != nil
}

// Optional SR-MPLS SID.
// SetSrMplsSid sets the BgpAttributesSidMpls value in the BgpAttributesSegmentRoutingPolicyTypeE object
func (obj *bgpAttributesSegmentRoutingPolicyTypeE) SetSrMplsSid(value BgpAttributesSidMpls) BgpAttributesSegmentRoutingPolicyTypeE {

	obj.srMplsSidHolder = nil
	obj.obj.SrMplsSid = value.msg()

	return obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeE) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(vObj, set_default)
	}

	if obj.obj.Ipv4NodeAddress != nil {

		err := obj.validateIpv4(obj.Ipv4NodeAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpAttributesSegmentRoutingPolicyTypeE.Ipv4NodeAddress"))
		}

	}

	if obj.obj.SrMplsSid != nil {

		obj.SrMplsSid().validateObj(vObj, set_default)
	}

}

func (obj *bgpAttributesSegmentRoutingPolicyTypeE) setDefault() {
	if obj.obj.LocalInterfaceId == nil {
		obj.SetLocalInterfaceId(0)
	}
	if obj.obj.Ipv4NodeAddress == nil {
		obj.SetIpv4NodeAddress("0.0.0.0")
	}

}
