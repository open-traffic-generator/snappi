package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesSegmentRoutingPolicyTypeI *****
type bgpAttributesSegmentRoutingPolicyTypeI struct {
	validation
	obj                        *otg.BgpAttributesSegmentRoutingPolicyTypeI
	marshaller                 marshalBgpAttributesSegmentRoutingPolicyTypeI
	unMarshaller               unMarshalBgpAttributesSegmentRoutingPolicyTypeI
	flagsHolder                BgpAttributesSegmentRoutingPolicyTypeFlags
	srv6SidHolder              BgpAttributesSidSrv6
	srv6EndpointBehaviorHolder BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
}

func NewBgpAttributesSegmentRoutingPolicyTypeI() BgpAttributesSegmentRoutingPolicyTypeI {
	obj := bgpAttributesSegmentRoutingPolicyTypeI{obj: &otg.BgpAttributesSegmentRoutingPolicyTypeI{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeI) msg() *otg.BgpAttributesSegmentRoutingPolicyTypeI {
	return obj.obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeI) setMsg(msg *otg.BgpAttributesSegmentRoutingPolicyTypeI) BgpAttributesSegmentRoutingPolicyTypeI {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesSegmentRoutingPolicyTypeI struct {
	obj *bgpAttributesSegmentRoutingPolicyTypeI
}

type marshalBgpAttributesSegmentRoutingPolicyTypeI interface {
	// ToProto marshals BgpAttributesSegmentRoutingPolicyTypeI to protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeI
	ToProto() (*otg.BgpAttributesSegmentRoutingPolicyTypeI, error)
	// ToPbText marshals BgpAttributesSegmentRoutingPolicyTypeI to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesSegmentRoutingPolicyTypeI to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesSegmentRoutingPolicyTypeI to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpAttributesSegmentRoutingPolicyTypeI to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpAttributesSegmentRoutingPolicyTypeI struct {
	obj *bgpAttributesSegmentRoutingPolicyTypeI
}

type unMarshalBgpAttributesSegmentRoutingPolicyTypeI interface {
	// FromProto unmarshals BgpAttributesSegmentRoutingPolicyTypeI from protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeI
	FromProto(msg *otg.BgpAttributesSegmentRoutingPolicyTypeI) (BgpAttributesSegmentRoutingPolicyTypeI, error)
	// FromPbText unmarshals BgpAttributesSegmentRoutingPolicyTypeI from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesSegmentRoutingPolicyTypeI from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesSegmentRoutingPolicyTypeI from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeI) Marshal() marshalBgpAttributesSegmentRoutingPolicyTypeI {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesSegmentRoutingPolicyTypeI{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeI) Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicyTypeI {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesSegmentRoutingPolicyTypeI{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeI) ToProto() (*otg.BgpAttributesSegmentRoutingPolicyTypeI, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeI) FromProto(msg *otg.BgpAttributesSegmentRoutingPolicyTypeI) (BgpAttributesSegmentRoutingPolicyTypeI, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeI) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeI) FromPbText(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeI) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeI) FromYaml(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeI) ToJsonRaw() (string, error) {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeI) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeI) FromJson(value string) error {
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

func (obj *bgpAttributesSegmentRoutingPolicyTypeI) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeI) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeI) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeI) Clone() (BgpAttributesSegmentRoutingPolicyTypeI, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesSegmentRoutingPolicyTypeI()
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

func (obj *bgpAttributesSegmentRoutingPolicyTypeI) setNil() {
	obj.flagsHolder = nil
	obj.srv6SidHolder = nil
	obj.srv6EndpointBehaviorHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpAttributesSegmentRoutingPolicyTypeI is type I: IPv6 Node Address with optional SR Algorithm and optional SRv6 SID.
// It is encoded as a Segment of Type 14 in the SEGMENT_LIST sub-tlv.
type BgpAttributesSegmentRoutingPolicyTypeI interface {
	Validation
	// msg marshals BgpAttributesSegmentRoutingPolicyTypeI to protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeI
	// and doesn't set defaults
	msg() *otg.BgpAttributesSegmentRoutingPolicyTypeI
	// setMsg unmarshals BgpAttributesSegmentRoutingPolicyTypeI from protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeI
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesSegmentRoutingPolicyTypeI) BgpAttributesSegmentRoutingPolicyTypeI
	// provides marshal interface
	Marshal() marshalBgpAttributesSegmentRoutingPolicyTypeI
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicyTypeI
	// validate validates BgpAttributesSegmentRoutingPolicyTypeI
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesSegmentRoutingPolicyTypeI, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns BgpAttributesSegmentRoutingPolicyTypeFlags, set in BgpAttributesSegmentRoutingPolicyTypeI.
	// BgpAttributesSegmentRoutingPolicyTypeFlags is flags for each Segment in SEGMENT_LIST sub-tlv.
	// - V-flag. Indicates verification is enabled. See section 5, of https://datatracker.ietf.org/doc/html/rfc9256
	// - A-flag. Indicates presence of SR Algorithm field applicable to Segment Types C, D , I , J and K .
	// - B-Flag. Indicates presence of SRv6 Endpoint Behavior and SID Structure encoding applicable to Segment Types B , I , J and K .
	// - S-Flag: This flag, when set, indicates the presence of the SR-MPLS or SRv6 SID depending on the segment type. (draft-ietf-idr-bgp-sr-segtypes-ext-03 Section 2.10).
	// This flag is applicable for Segment Types C, D, E, F, G, H, I, J, and K.
	Flags() BgpAttributesSegmentRoutingPolicyTypeFlags
	// SetFlags assigns BgpAttributesSegmentRoutingPolicyTypeFlags provided by user to BgpAttributesSegmentRoutingPolicyTypeI.
	// BgpAttributesSegmentRoutingPolicyTypeFlags is flags for each Segment in SEGMENT_LIST sub-tlv.
	// - V-flag. Indicates verification is enabled. See section 5, of https://datatracker.ietf.org/doc/html/rfc9256
	// - A-flag. Indicates presence of SR Algorithm field applicable to Segment Types C, D , I , J and K .
	// - B-Flag. Indicates presence of SRv6 Endpoint Behavior and SID Structure encoding applicable to Segment Types B , I , J and K .
	// - S-Flag: This flag, when set, indicates the presence of the SR-MPLS or SRv6 SID depending on the segment type. (draft-ietf-idr-bgp-sr-segtypes-ext-03 Section 2.10).
	// This flag is applicable for Segment Types C, D, E, F, G, H, I, J, and K.
	SetFlags(value BgpAttributesSegmentRoutingPolicyTypeFlags) BgpAttributesSegmentRoutingPolicyTypeI
	// HasFlags checks if Flags has been set in BgpAttributesSegmentRoutingPolicyTypeI
	HasFlags() bool
	// SrAlgorithm returns uint32, set in BgpAttributesSegmentRoutingPolicyTypeI.
	SrAlgorithm() uint32
	// SetSrAlgorithm assigns uint32 provided by user to BgpAttributesSegmentRoutingPolicyTypeI
	SetSrAlgorithm(value uint32) BgpAttributesSegmentRoutingPolicyTypeI
	// HasSrAlgorithm checks if SrAlgorithm has been set in BgpAttributesSegmentRoutingPolicyTypeI
	HasSrAlgorithm() bool
	// Ipv6NodeAddress returns string, set in BgpAttributesSegmentRoutingPolicyTypeI.
	Ipv6NodeAddress() string
	// SetIpv6NodeAddress assigns string provided by user to BgpAttributesSegmentRoutingPolicyTypeI
	SetIpv6NodeAddress(value string) BgpAttributesSegmentRoutingPolicyTypeI
	// HasIpv6NodeAddress checks if Ipv6NodeAddress has been set in BgpAttributesSegmentRoutingPolicyTypeI
	HasIpv6NodeAddress() bool
	// Srv6Sid returns BgpAttributesSidSrv6, set in BgpAttributesSegmentRoutingPolicyTypeI.
	// BgpAttributesSidSrv6 is an IPv6 address denoting a SRv6 SID.
	Srv6Sid() BgpAttributesSidSrv6
	// SetSrv6Sid assigns BgpAttributesSidSrv6 provided by user to BgpAttributesSegmentRoutingPolicyTypeI.
	// BgpAttributesSidSrv6 is an IPv6 address denoting a SRv6 SID.
	SetSrv6Sid(value BgpAttributesSidSrv6) BgpAttributesSegmentRoutingPolicyTypeI
	// HasSrv6Sid checks if Srv6Sid has been set in BgpAttributesSegmentRoutingPolicyTypeI
	HasSrv6Sid() bool
	// Srv6EndpointBehavior returns BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure, set in BgpAttributesSegmentRoutingPolicyTypeI.
	// BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure is configuration for optional SRv6 Endpoint Behavior and SID Structure. Summation of lengths for Locator Block, Locator Node,  Function, and Argument MUST be less than or equal to 128. - This is specified in draft-ietf-idr-sr-policy-safi-02 Section 2.4.4.2.4
	Srv6EndpointBehavior() BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	// SetSrv6EndpointBehavior assigns BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure provided by user to BgpAttributesSegmentRoutingPolicyTypeI.
	// BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure is configuration for optional SRv6 Endpoint Behavior and SID Structure. Summation of lengths for Locator Block, Locator Node,  Function, and Argument MUST be less than or equal to 128. - This is specified in draft-ietf-idr-sr-policy-safi-02 Section 2.4.4.2.4
	SetSrv6EndpointBehavior(value BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) BgpAttributesSegmentRoutingPolicyTypeI
	// HasSrv6EndpointBehavior checks if Srv6EndpointBehavior has been set in BgpAttributesSegmentRoutingPolicyTypeI
	HasSrv6EndpointBehavior() bool
	setNil()
}

// description is TBD
// Flags returns a BgpAttributesSegmentRoutingPolicyTypeFlags
func (obj *bgpAttributesSegmentRoutingPolicyTypeI) Flags() BgpAttributesSegmentRoutingPolicyTypeFlags {
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
func (obj *bgpAttributesSegmentRoutingPolicyTypeI) HasFlags() bool {
	return obj.obj.Flags != nil
}

// description is TBD
// SetFlags sets the BgpAttributesSegmentRoutingPolicyTypeFlags value in the BgpAttributesSegmentRoutingPolicyTypeI object
func (obj *bgpAttributesSegmentRoutingPolicyTypeI) SetFlags(value BgpAttributesSegmentRoutingPolicyTypeFlags) BgpAttributesSegmentRoutingPolicyTypeI {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// SR Algorithm identifier when A-Flag in on. If A-flag is not enabled, it should be set to 0 on transmission and ignored on receipt.
// SrAlgorithm returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicyTypeI) SrAlgorithm() uint32 {

	return *obj.obj.SrAlgorithm

}

// SR Algorithm identifier when A-Flag in on. If A-flag is not enabled, it should be set to 0 on transmission and ignored on receipt.
// SrAlgorithm returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicyTypeI) HasSrAlgorithm() bool {
	return obj.obj.SrAlgorithm != nil
}

// SR Algorithm identifier when A-Flag in on. If A-flag is not enabled, it should be set to 0 on transmission and ignored on receipt.
// SetSrAlgorithm sets the uint32 value in the BgpAttributesSegmentRoutingPolicyTypeI object
func (obj *bgpAttributesSegmentRoutingPolicyTypeI) SetSrAlgorithm(value uint32) BgpAttributesSegmentRoutingPolicyTypeI {

	obj.obj.SrAlgorithm = &value
	return obj
}

// IPv6 address representing a node.
// Ipv6NodeAddress returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeI) Ipv6NodeAddress() string {

	return *obj.obj.Ipv6NodeAddress

}

// IPv6 address representing a node.
// Ipv6NodeAddress returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeI) HasIpv6NodeAddress() bool {
	return obj.obj.Ipv6NodeAddress != nil
}

// IPv6 address representing a node.
// SetIpv6NodeAddress sets the string value in the BgpAttributesSegmentRoutingPolicyTypeI object
func (obj *bgpAttributesSegmentRoutingPolicyTypeI) SetIpv6NodeAddress(value string) BgpAttributesSegmentRoutingPolicyTypeI {

	obj.obj.Ipv6NodeAddress = &value
	return obj
}

// description is TBD
// Srv6Sid returns a BgpAttributesSidSrv6
func (obj *bgpAttributesSegmentRoutingPolicyTypeI) Srv6Sid() BgpAttributesSidSrv6 {
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
func (obj *bgpAttributesSegmentRoutingPolicyTypeI) HasSrv6Sid() bool {
	return obj.obj.Srv6Sid != nil
}

// description is TBD
// SetSrv6Sid sets the BgpAttributesSidSrv6 value in the BgpAttributesSegmentRoutingPolicyTypeI object
func (obj *bgpAttributesSegmentRoutingPolicyTypeI) SetSrv6Sid(value BgpAttributesSidSrv6) BgpAttributesSegmentRoutingPolicyTypeI {

	obj.srv6SidHolder = nil
	obj.obj.Srv6Sid = value.msg()

	return obj
}

// description is TBD
// Srv6EndpointBehavior returns a BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
func (obj *bgpAttributesSegmentRoutingPolicyTypeI) Srv6EndpointBehavior() BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure {
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
func (obj *bgpAttributesSegmentRoutingPolicyTypeI) HasSrv6EndpointBehavior() bool {
	return obj.obj.Srv6EndpointBehavior != nil
}

// description is TBD
// SetSrv6EndpointBehavior sets the BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure value in the BgpAttributesSegmentRoutingPolicyTypeI object
func (obj *bgpAttributesSegmentRoutingPolicyTypeI) SetSrv6EndpointBehavior(value BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) BgpAttributesSegmentRoutingPolicyTypeI {

	obj.srv6EndpointBehaviorHolder = nil
	obj.obj.Srv6EndpointBehavior = value.msg()

	return obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeI) validateObj(vObj *validation, set_default bool) {
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
				fmt.Sprintf("0 <= BgpAttributesSegmentRoutingPolicyTypeI.SrAlgorithm <= 255 but Got %d", *obj.obj.SrAlgorithm))
		}

	}

	if obj.obj.Ipv6NodeAddress != nil {

		err := obj.validateIpv6(obj.Ipv6NodeAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpAttributesSegmentRoutingPolicyTypeI.Ipv6NodeAddress"))
		}

	}

	if obj.obj.Srv6Sid != nil {

		obj.Srv6Sid().validateObj(vObj, set_default)
	}

	if obj.obj.Srv6EndpointBehavior != nil {

		obj.Srv6EndpointBehavior().validateObj(vObj, set_default)
	}

}

func (obj *bgpAttributesSegmentRoutingPolicyTypeI) setDefault() {
	if obj.obj.SrAlgorithm == nil {
		obj.SetSrAlgorithm(0)
	}
	if obj.obj.Ipv6NodeAddress == nil {
		obj.SetIpv6NodeAddress("0::0")
	}

}
