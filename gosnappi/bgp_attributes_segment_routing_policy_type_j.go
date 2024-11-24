package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesSegmentRoutingPolicyTypeJ *****
type bgpAttributesSegmentRoutingPolicyTypeJ struct {
	validation
	obj                        *otg.BgpAttributesSegmentRoutingPolicyTypeJ
	marshaller                 marshalBgpAttributesSegmentRoutingPolicyTypeJ
	unMarshaller               unMarshalBgpAttributesSegmentRoutingPolicyTypeJ
	flagsHolder                BgpAttributesSegmentRoutingPolicyTypeFlags
	srv6SidHolder              BgpAttributesSidSrv6
	srv6EndpointBehaviorHolder BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
}

func NewBgpAttributesSegmentRoutingPolicyTypeJ() BgpAttributesSegmentRoutingPolicyTypeJ {
	obj := bgpAttributesSegmentRoutingPolicyTypeJ{obj: &otg.BgpAttributesSegmentRoutingPolicyTypeJ{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) msg() *otg.BgpAttributesSegmentRoutingPolicyTypeJ {
	return obj.obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) setMsg(msg *otg.BgpAttributesSegmentRoutingPolicyTypeJ) BgpAttributesSegmentRoutingPolicyTypeJ {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesSegmentRoutingPolicyTypeJ struct {
	obj *bgpAttributesSegmentRoutingPolicyTypeJ
}

type marshalBgpAttributesSegmentRoutingPolicyTypeJ interface {
	// ToProto marshals BgpAttributesSegmentRoutingPolicyTypeJ to protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeJ
	ToProto() (*otg.BgpAttributesSegmentRoutingPolicyTypeJ, error)
	// ToPbText marshals BgpAttributesSegmentRoutingPolicyTypeJ to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesSegmentRoutingPolicyTypeJ to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesSegmentRoutingPolicyTypeJ to JSON text
	ToJson() (string, error)
}

type unMarshalbgpAttributesSegmentRoutingPolicyTypeJ struct {
	obj *bgpAttributesSegmentRoutingPolicyTypeJ
}

type unMarshalBgpAttributesSegmentRoutingPolicyTypeJ interface {
	// FromProto unmarshals BgpAttributesSegmentRoutingPolicyTypeJ from protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeJ
	FromProto(msg *otg.BgpAttributesSegmentRoutingPolicyTypeJ) (BgpAttributesSegmentRoutingPolicyTypeJ, error)
	// FromPbText unmarshals BgpAttributesSegmentRoutingPolicyTypeJ from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesSegmentRoutingPolicyTypeJ from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesSegmentRoutingPolicyTypeJ from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) Marshal() marshalBgpAttributesSegmentRoutingPolicyTypeJ {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesSegmentRoutingPolicyTypeJ{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicyTypeJ {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesSegmentRoutingPolicyTypeJ{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeJ) ToProto() (*otg.BgpAttributesSegmentRoutingPolicyTypeJ, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeJ) FromProto(msg *otg.BgpAttributesSegmentRoutingPolicyTypeJ) (BgpAttributesSegmentRoutingPolicyTypeJ, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeJ) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeJ) FromPbText(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeJ) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeJ) FromYaml(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeJ) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeJ) FromJson(value string) error {
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

func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) Clone() (BgpAttributesSegmentRoutingPolicyTypeJ, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesSegmentRoutingPolicyTypeJ()
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

func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) setNil() {
	obj.flagsHolder = nil
	obj.srv6SidHolder = nil
	obj.srv6EndpointBehaviorHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpAttributesSegmentRoutingPolicyTypeJ is type J: IPv6 Address, Interface ID for local and remote pair for SRv6 with optional SID.
// It is encoded as a Segment of Type 15 in the SEGMENT_LIST sub-tlv.
type BgpAttributesSegmentRoutingPolicyTypeJ interface {
	Validation
	// msg marshals BgpAttributesSegmentRoutingPolicyTypeJ to protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeJ
	// and doesn't set defaults
	msg() *otg.BgpAttributesSegmentRoutingPolicyTypeJ
	// setMsg unmarshals BgpAttributesSegmentRoutingPolicyTypeJ from protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeJ
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesSegmentRoutingPolicyTypeJ) BgpAttributesSegmentRoutingPolicyTypeJ
	// provides marshal interface
	Marshal() marshalBgpAttributesSegmentRoutingPolicyTypeJ
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicyTypeJ
	// validate validates BgpAttributesSegmentRoutingPolicyTypeJ
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesSegmentRoutingPolicyTypeJ, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns BgpAttributesSegmentRoutingPolicyTypeFlags, set in BgpAttributesSegmentRoutingPolicyTypeJ.
	// BgpAttributesSegmentRoutingPolicyTypeFlags is flags for each Segment in SEGMENT_LIST sub-tlv.
	// - V-flag. Indicates verification is enabled. See section 5, of https://datatracker.ietf.org/doc/html/rfc9256
	// - A-flag. Indicates presence of SR Algorithm field applicable to Segment Types C, D , I , J and K .
	// - B-Flag. Indicates presence of SRv6 Endpoint Behavior and SID Structure encoding applicable to Segment Types B , I , J and K .
	// - S-Flag: This flag, when set, indicates the presence of the SR-MPLS or SRv6 SID depending on the segment type. (draft-ietf-idr-bgp-sr-segtypes-ext-03 Section 2.10).
	// This flag is applicable for Segment Types C, D, E, F, G, H, I, J, and K.
	Flags() BgpAttributesSegmentRoutingPolicyTypeFlags
	// SetFlags assigns BgpAttributesSegmentRoutingPolicyTypeFlags provided by user to BgpAttributesSegmentRoutingPolicyTypeJ.
	// BgpAttributesSegmentRoutingPolicyTypeFlags is flags for each Segment in SEGMENT_LIST sub-tlv.
	// - V-flag. Indicates verification is enabled. See section 5, of https://datatracker.ietf.org/doc/html/rfc9256
	// - A-flag. Indicates presence of SR Algorithm field applicable to Segment Types C, D , I , J and K .
	// - B-Flag. Indicates presence of SRv6 Endpoint Behavior and SID Structure encoding applicable to Segment Types B , I , J and K .
	// - S-Flag: This flag, when set, indicates the presence of the SR-MPLS or SRv6 SID depending on the segment type. (draft-ietf-idr-bgp-sr-segtypes-ext-03 Section 2.10).
	// This flag is applicable for Segment Types C, D, E, F, G, H, I, J, and K.
	SetFlags(value BgpAttributesSegmentRoutingPolicyTypeFlags) BgpAttributesSegmentRoutingPolicyTypeJ
	// HasFlags checks if Flags has been set in BgpAttributesSegmentRoutingPolicyTypeJ
	HasFlags() bool
	// SrAlgorithm returns uint32, set in BgpAttributesSegmentRoutingPolicyTypeJ.
	SrAlgorithm() uint32
	// SetSrAlgorithm assigns uint32 provided by user to BgpAttributesSegmentRoutingPolicyTypeJ
	SetSrAlgorithm(value uint32) BgpAttributesSegmentRoutingPolicyTypeJ
	// HasSrAlgorithm checks if SrAlgorithm has been set in BgpAttributesSegmentRoutingPolicyTypeJ
	HasSrAlgorithm() bool
	// LocalInterfaceId returns uint32, set in BgpAttributesSegmentRoutingPolicyTypeJ.
	LocalInterfaceId() uint32
	// SetLocalInterfaceId assigns uint32 provided by user to BgpAttributesSegmentRoutingPolicyTypeJ
	SetLocalInterfaceId(value uint32) BgpAttributesSegmentRoutingPolicyTypeJ
	// HasLocalInterfaceId checks if LocalInterfaceId has been set in BgpAttributesSegmentRoutingPolicyTypeJ
	HasLocalInterfaceId() bool
	// LocalIpv6NodeAddress returns string, set in BgpAttributesSegmentRoutingPolicyTypeJ.
	LocalIpv6NodeAddress() string
	// SetLocalIpv6NodeAddress assigns string provided by user to BgpAttributesSegmentRoutingPolicyTypeJ
	SetLocalIpv6NodeAddress(value string) BgpAttributesSegmentRoutingPolicyTypeJ
	// HasLocalIpv6NodeAddress checks if LocalIpv6NodeAddress has been set in BgpAttributesSegmentRoutingPolicyTypeJ
	HasLocalIpv6NodeAddress() bool
	// RemoteInterfaceId returns uint32, set in BgpAttributesSegmentRoutingPolicyTypeJ.
	RemoteInterfaceId() uint32
	// SetRemoteInterfaceId assigns uint32 provided by user to BgpAttributesSegmentRoutingPolicyTypeJ
	SetRemoteInterfaceId(value uint32) BgpAttributesSegmentRoutingPolicyTypeJ
	// HasRemoteInterfaceId checks if RemoteInterfaceId has been set in BgpAttributesSegmentRoutingPolicyTypeJ
	HasRemoteInterfaceId() bool
	// RemoteIpv6NodeAddress returns string, set in BgpAttributesSegmentRoutingPolicyTypeJ.
	RemoteIpv6NodeAddress() string
	// SetRemoteIpv6NodeAddress assigns string provided by user to BgpAttributesSegmentRoutingPolicyTypeJ
	SetRemoteIpv6NodeAddress(value string) BgpAttributesSegmentRoutingPolicyTypeJ
	// HasRemoteIpv6NodeAddress checks if RemoteIpv6NodeAddress has been set in BgpAttributesSegmentRoutingPolicyTypeJ
	HasRemoteIpv6NodeAddress() bool
	// Srv6Sid returns BgpAttributesSidSrv6, set in BgpAttributesSegmentRoutingPolicyTypeJ.
	// BgpAttributesSidSrv6 is an IPv6 address denoting a SRv6 SID.
	Srv6Sid() BgpAttributesSidSrv6
	// SetSrv6Sid assigns BgpAttributesSidSrv6 provided by user to BgpAttributesSegmentRoutingPolicyTypeJ.
	// BgpAttributesSidSrv6 is an IPv6 address denoting a SRv6 SID.
	SetSrv6Sid(value BgpAttributesSidSrv6) BgpAttributesSegmentRoutingPolicyTypeJ
	// HasSrv6Sid checks if Srv6Sid has been set in BgpAttributesSegmentRoutingPolicyTypeJ
	HasSrv6Sid() bool
	// Srv6EndpointBehavior returns BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure, set in BgpAttributesSegmentRoutingPolicyTypeJ.
	// BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure is configuration for optional SRv6 Endpoint Behavior and SID Structure. Summation of lengths for Locator Block, Locator Node,  Function, and Argument MUST be less than or equal to 128. - This is specified in draft-ietf-idr-sr-policy-safi-02 Section 2.4.4.2.4
	Srv6EndpointBehavior() BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	// SetSrv6EndpointBehavior assigns BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure provided by user to BgpAttributesSegmentRoutingPolicyTypeJ.
	// BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure is configuration for optional SRv6 Endpoint Behavior and SID Structure. Summation of lengths for Locator Block, Locator Node,  Function, and Argument MUST be less than or equal to 128. - This is specified in draft-ietf-idr-sr-policy-safi-02 Section 2.4.4.2.4
	SetSrv6EndpointBehavior(value BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) BgpAttributesSegmentRoutingPolicyTypeJ
	// HasSrv6EndpointBehavior checks if Srv6EndpointBehavior has been set in BgpAttributesSegmentRoutingPolicyTypeJ
	HasSrv6EndpointBehavior() bool
	setNil()
}

// description is TBD
// Flags returns a BgpAttributesSegmentRoutingPolicyTypeFlags
func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) Flags() BgpAttributesSegmentRoutingPolicyTypeFlags {
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
func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) HasFlags() bool {
	return obj.obj.Flags != nil
}

// description is TBD
// SetFlags sets the BgpAttributesSegmentRoutingPolicyTypeFlags value in the BgpAttributesSegmentRoutingPolicyTypeJ object
func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) SetFlags(value BgpAttributesSegmentRoutingPolicyTypeFlags) BgpAttributesSegmentRoutingPolicyTypeJ {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// SR Algorithm identifier when A-Flag in on. If A-flag is not enabled, it should be set to 0 on transmission and ignored on receipt.
// SrAlgorithm returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) SrAlgorithm() uint32 {

	return *obj.obj.SrAlgorithm

}

// SR Algorithm identifier when A-Flag in on. If A-flag is not enabled, it should be set to 0 on transmission and ignored on receipt.
// SrAlgorithm returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) HasSrAlgorithm() bool {
	return obj.obj.SrAlgorithm != nil
}

// SR Algorithm identifier when A-Flag in on. If A-flag is not enabled, it should be set to 0 on transmission and ignored on receipt.
// SetSrAlgorithm sets the uint32 value in the BgpAttributesSegmentRoutingPolicyTypeJ object
func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) SetSrAlgorithm(value uint32) BgpAttributesSegmentRoutingPolicyTypeJ {

	obj.obj.SrAlgorithm = &value
	return obj
}

// The local Interface Index as defined in [RFC8664].
// LocalInterfaceId returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) LocalInterfaceId() uint32 {

	return *obj.obj.LocalInterfaceId

}

// The local Interface Index as defined in [RFC8664].
// LocalInterfaceId returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) HasLocalInterfaceId() bool {
	return obj.obj.LocalInterfaceId != nil
}

// The local Interface Index as defined in [RFC8664].
// SetLocalInterfaceId sets the uint32 value in the BgpAttributesSegmentRoutingPolicyTypeJ object
func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) SetLocalInterfaceId(value uint32) BgpAttributesSegmentRoutingPolicyTypeJ {

	obj.obj.LocalInterfaceId = &value
	return obj
}

// The IPv6 address representing the local node.
// LocalIpv6NodeAddress returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) LocalIpv6NodeAddress() string {

	return *obj.obj.LocalIpv6NodeAddress

}

// The IPv6 address representing the local node.
// LocalIpv6NodeAddress returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) HasLocalIpv6NodeAddress() bool {
	return obj.obj.LocalIpv6NodeAddress != nil
}

// The IPv6 address representing the local node.
// SetLocalIpv6NodeAddress sets the string value in the BgpAttributesSegmentRoutingPolicyTypeJ object
func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) SetLocalIpv6NodeAddress(value string) BgpAttributesSegmentRoutingPolicyTypeJ {

	obj.obj.LocalIpv6NodeAddress = &value
	return obj
}

// The remote Interface Index as defined in [RFC8664]. The value MAY be set to zero when the local node address and interface identifiers are sufficient to describe the link.
// RemoteInterfaceId returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) RemoteInterfaceId() uint32 {

	return *obj.obj.RemoteInterfaceId

}

// The remote Interface Index as defined in [RFC8664]. The value MAY be set to zero when the local node address and interface identifiers are sufficient to describe the link.
// RemoteInterfaceId returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) HasRemoteInterfaceId() bool {
	return obj.obj.RemoteInterfaceId != nil
}

// The remote Interface Index as defined in [RFC8664]. The value MAY be set to zero when the local node address and interface identifiers are sufficient to describe the link.
// SetRemoteInterfaceId sets the uint32 value in the BgpAttributesSegmentRoutingPolicyTypeJ object
func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) SetRemoteInterfaceId(value uint32) BgpAttributesSegmentRoutingPolicyTypeJ {

	obj.obj.RemoteInterfaceId = &value
	return obj
}

// IPv6 address representing the remote node. The value MAY be set to zero when the local node address and interface identifiers are sufficient to describe the link.
// RemoteIpv6NodeAddress returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) RemoteIpv6NodeAddress() string {

	return *obj.obj.RemoteIpv6NodeAddress

}

// IPv6 address representing the remote node. The value MAY be set to zero when the local node address and interface identifiers are sufficient to describe the link.
// RemoteIpv6NodeAddress returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) HasRemoteIpv6NodeAddress() bool {
	return obj.obj.RemoteIpv6NodeAddress != nil
}

// IPv6 address representing the remote node. The value MAY be set to zero when the local node address and interface identifiers are sufficient to describe the link.
// SetRemoteIpv6NodeAddress sets the string value in the BgpAttributesSegmentRoutingPolicyTypeJ object
func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) SetRemoteIpv6NodeAddress(value string) BgpAttributesSegmentRoutingPolicyTypeJ {

	obj.obj.RemoteIpv6NodeAddress = &value
	return obj
}

// description is TBD
// Srv6Sid returns a BgpAttributesSidSrv6
func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) Srv6Sid() BgpAttributesSidSrv6 {
	if obj.obj.Srv6Sid == nil {
		obj.obj.Srv6Sid = NewBgpAttributesSidSrv6().msg()
	}
	if obj.srv6SidHolder == nil {
		obj.srv6SidHolder = &bgpAttributesSidSrv6{obj: obj.obj.Srv6Sid}
	}
	return obj.srv6SidHolder
}

// description is TBD
// Srv6Sid returns a BgpAttributesSidSrv6
func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) HasSrv6Sid() bool {
	return obj.obj.Srv6Sid != nil
}

// description is TBD
// SetSrv6Sid sets the BgpAttributesSidSrv6 value in the BgpAttributesSegmentRoutingPolicyTypeJ object
func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) SetSrv6Sid(value BgpAttributesSidSrv6) BgpAttributesSegmentRoutingPolicyTypeJ {

	obj.srv6SidHolder = nil
	obj.obj.Srv6Sid = value.msg()

	return obj
}

// description is TBD
// Srv6EndpointBehavior returns a BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) Srv6EndpointBehavior() BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure {
	if obj.obj.Srv6EndpointBehavior == nil {
		obj.obj.Srv6EndpointBehavior = NewBgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure().msg()
	}
	if obj.srv6EndpointBehaviorHolder == nil {
		obj.srv6EndpointBehaviorHolder = &bgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure{obj: obj.obj.Srv6EndpointBehavior}
	}
	return obj.srv6EndpointBehaviorHolder
}

// description is TBD
// Srv6EndpointBehavior returns a BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) HasSrv6EndpointBehavior() bool {
	return obj.obj.Srv6EndpointBehavior != nil
}

// description is TBD
// SetSrv6EndpointBehavior sets the BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure value in the BgpAttributesSegmentRoutingPolicyTypeJ object
func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) SetSrv6EndpointBehavior(value BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) BgpAttributesSegmentRoutingPolicyTypeJ {

	obj.srv6EndpointBehaviorHolder = nil
	obj.obj.Srv6EndpointBehavior = value.msg()

	return obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) validateObj(vObj *validation, set_default bool) {
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
				fmt.Sprintf("0 <= BgpAttributesSegmentRoutingPolicyTypeJ.SrAlgorithm <= 255 but Got %d", *obj.obj.SrAlgorithm))
		}

	}

	if obj.obj.LocalIpv6NodeAddress != nil {

		err := obj.validateIpv6(obj.LocalIpv6NodeAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpAttributesSegmentRoutingPolicyTypeJ.LocalIpv6NodeAddress"))
		}

	}

	if obj.obj.RemoteIpv6NodeAddress != nil {

		err := obj.validateIpv6(obj.RemoteIpv6NodeAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpAttributesSegmentRoutingPolicyTypeJ.RemoteIpv6NodeAddress"))
		}

	}

	if obj.obj.Srv6Sid != nil {

		obj.Srv6Sid().validateObj(vObj, set_default)
	}

	if obj.obj.Srv6EndpointBehavior != nil {

		obj.Srv6EndpointBehavior().validateObj(vObj, set_default)
	}

}

func (obj *bgpAttributesSegmentRoutingPolicyTypeJ) setDefault() {
	if obj.obj.SrAlgorithm == nil {
		obj.SetSrAlgorithm(0)
	}
	if obj.obj.LocalInterfaceId == nil {
		obj.SetLocalInterfaceId(0)
	}
	if obj.obj.LocalIpv6NodeAddress == nil {
		obj.SetLocalIpv6NodeAddress("0::0")
	}
	if obj.obj.RemoteInterfaceId == nil {
		obj.SetRemoteInterfaceId(0)
	}
	if obj.obj.RemoteIpv6NodeAddress == nil {
		obj.SetRemoteIpv6NodeAddress("0::0")
	}

}
