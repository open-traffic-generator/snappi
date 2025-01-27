package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesSegmentRoutingPolicyTypeK *****
type bgpAttributesSegmentRoutingPolicyTypeK struct {
	validation
	obj                        *otg.BgpAttributesSegmentRoutingPolicyTypeK
	marshaller                 marshalBgpAttributesSegmentRoutingPolicyTypeK
	unMarshaller               unMarshalBgpAttributesSegmentRoutingPolicyTypeK
	flagsHolder                BgpAttributesSegmentRoutingPolicyTypeFlags
	srv6SidHolder              BgpAttributesSidSrv6
	srv6EndpointBehaviorHolder BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
}

func NewBgpAttributesSegmentRoutingPolicyTypeK() BgpAttributesSegmentRoutingPolicyTypeK {
	obj := bgpAttributesSegmentRoutingPolicyTypeK{obj: &otg.BgpAttributesSegmentRoutingPolicyTypeK{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeK) msg() *otg.BgpAttributesSegmentRoutingPolicyTypeK {
	return obj.obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeK) setMsg(msg *otg.BgpAttributesSegmentRoutingPolicyTypeK) BgpAttributesSegmentRoutingPolicyTypeK {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesSegmentRoutingPolicyTypeK struct {
	obj *bgpAttributesSegmentRoutingPolicyTypeK
}

type marshalBgpAttributesSegmentRoutingPolicyTypeK interface {
	// ToProto marshals BgpAttributesSegmentRoutingPolicyTypeK to protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeK
	ToProto() (*otg.BgpAttributesSegmentRoutingPolicyTypeK, error)
	// ToPbText marshals BgpAttributesSegmentRoutingPolicyTypeK to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesSegmentRoutingPolicyTypeK to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesSegmentRoutingPolicyTypeK to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpAttributesSegmentRoutingPolicyTypeK to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpAttributesSegmentRoutingPolicyTypeK struct {
	obj *bgpAttributesSegmentRoutingPolicyTypeK
}

type unMarshalBgpAttributesSegmentRoutingPolicyTypeK interface {
	// FromProto unmarshals BgpAttributesSegmentRoutingPolicyTypeK from protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeK
	FromProto(msg *otg.BgpAttributesSegmentRoutingPolicyTypeK) (BgpAttributesSegmentRoutingPolicyTypeK, error)
	// FromPbText unmarshals BgpAttributesSegmentRoutingPolicyTypeK from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesSegmentRoutingPolicyTypeK from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesSegmentRoutingPolicyTypeK from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeK) Marshal() marshalBgpAttributesSegmentRoutingPolicyTypeK {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesSegmentRoutingPolicyTypeK{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeK) Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicyTypeK {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesSegmentRoutingPolicyTypeK{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeK) ToProto() (*otg.BgpAttributesSegmentRoutingPolicyTypeK, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeK) FromProto(msg *otg.BgpAttributesSegmentRoutingPolicyTypeK) (BgpAttributesSegmentRoutingPolicyTypeK, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeK) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeK) FromPbText(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeK) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeK) FromYaml(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeK) ToJsonRaw() (string, error) {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeK) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeK) FromJson(value string) error {
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

func (obj *bgpAttributesSegmentRoutingPolicyTypeK) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeK) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeK) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeK) Clone() (BgpAttributesSegmentRoutingPolicyTypeK, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesSegmentRoutingPolicyTypeK()
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

func (obj *bgpAttributesSegmentRoutingPolicyTypeK) setNil() {
	obj.flagsHolder = nil
	obj.srv6SidHolder = nil
	obj.srv6EndpointBehaviorHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpAttributesSegmentRoutingPolicyTypeK is type K: IPv6 Local and Remote addresses for SRv6 with optional SID.
// It is encoded as a Segment of Type 16 in the SEGMENT_LIST sub-tlv.
type BgpAttributesSegmentRoutingPolicyTypeK interface {
	Validation
	// msg marshals BgpAttributesSegmentRoutingPolicyTypeK to protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeK
	// and doesn't set defaults
	msg() *otg.BgpAttributesSegmentRoutingPolicyTypeK
	// setMsg unmarshals BgpAttributesSegmentRoutingPolicyTypeK from protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeK
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesSegmentRoutingPolicyTypeK) BgpAttributesSegmentRoutingPolicyTypeK
	// provides marshal interface
	Marshal() marshalBgpAttributesSegmentRoutingPolicyTypeK
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicyTypeK
	// validate validates BgpAttributesSegmentRoutingPolicyTypeK
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesSegmentRoutingPolicyTypeK, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns BgpAttributesSegmentRoutingPolicyTypeFlags, set in BgpAttributesSegmentRoutingPolicyTypeK.
	// BgpAttributesSegmentRoutingPolicyTypeFlags is flags for each Segment in SEGMENT_LIST sub-tlv.
	// - V-flag. Indicates verification is enabled. See section 5, of https://datatracker.ietf.org/doc/html/rfc9256
	// - A-flag. Indicates presence of SR Algorithm field applicable to Segment Types C, D , I , J and K .
	// - B-Flag. Indicates presence of SRv6 Endpoint Behavior and SID Structure encoding applicable to Segment Types B , I , J and K .
	// - S-Flag: This flag, when set, indicates the presence of the SR-MPLS or SRv6 SID depending on the segment type. (draft-ietf-idr-bgp-sr-segtypes-ext-03 Section 2.10).
	// This flag is applicable for Segment Types C, D, E, F, G, H, I, J, and K.
	Flags() BgpAttributesSegmentRoutingPolicyTypeFlags
	// SetFlags assigns BgpAttributesSegmentRoutingPolicyTypeFlags provided by user to BgpAttributesSegmentRoutingPolicyTypeK.
	// BgpAttributesSegmentRoutingPolicyTypeFlags is flags for each Segment in SEGMENT_LIST sub-tlv.
	// - V-flag. Indicates verification is enabled. See section 5, of https://datatracker.ietf.org/doc/html/rfc9256
	// - A-flag. Indicates presence of SR Algorithm field applicable to Segment Types C, D , I , J and K .
	// - B-Flag. Indicates presence of SRv6 Endpoint Behavior and SID Structure encoding applicable to Segment Types B , I , J and K .
	// - S-Flag: This flag, when set, indicates the presence of the SR-MPLS or SRv6 SID depending on the segment type. (draft-ietf-idr-bgp-sr-segtypes-ext-03 Section 2.10).
	// This flag is applicable for Segment Types C, D, E, F, G, H, I, J, and K.
	SetFlags(value BgpAttributesSegmentRoutingPolicyTypeFlags) BgpAttributesSegmentRoutingPolicyTypeK
	// HasFlags checks if Flags has been set in BgpAttributesSegmentRoutingPolicyTypeK
	HasFlags() bool
	// SrAlgorithm returns uint32, set in BgpAttributesSegmentRoutingPolicyTypeK.
	SrAlgorithm() uint32
	// SetSrAlgorithm assigns uint32 provided by user to BgpAttributesSegmentRoutingPolicyTypeK
	SetSrAlgorithm(value uint32) BgpAttributesSegmentRoutingPolicyTypeK
	// HasSrAlgorithm checks if SrAlgorithm has been set in BgpAttributesSegmentRoutingPolicyTypeK
	HasSrAlgorithm() bool
	// LocalIpv6Address returns string, set in BgpAttributesSegmentRoutingPolicyTypeK.
	LocalIpv6Address() string
	// SetLocalIpv6Address assigns string provided by user to BgpAttributesSegmentRoutingPolicyTypeK
	SetLocalIpv6Address(value string) BgpAttributesSegmentRoutingPolicyTypeK
	// HasLocalIpv6Address checks if LocalIpv6Address has been set in BgpAttributesSegmentRoutingPolicyTypeK
	HasLocalIpv6Address() bool
	// RemoteIpv6Address returns string, set in BgpAttributesSegmentRoutingPolicyTypeK.
	RemoteIpv6Address() string
	// SetRemoteIpv6Address assigns string provided by user to BgpAttributesSegmentRoutingPolicyTypeK
	SetRemoteIpv6Address(value string) BgpAttributesSegmentRoutingPolicyTypeK
	// HasRemoteIpv6Address checks if RemoteIpv6Address has been set in BgpAttributesSegmentRoutingPolicyTypeK
	HasRemoteIpv6Address() bool
	// Srv6Sid returns BgpAttributesSidSrv6, set in BgpAttributesSegmentRoutingPolicyTypeK.
	// BgpAttributesSidSrv6 is an IPv6 address denoting a SRv6 SID.
	Srv6Sid() BgpAttributesSidSrv6
	// SetSrv6Sid assigns BgpAttributesSidSrv6 provided by user to BgpAttributesSegmentRoutingPolicyTypeK.
	// BgpAttributesSidSrv6 is an IPv6 address denoting a SRv6 SID.
	SetSrv6Sid(value BgpAttributesSidSrv6) BgpAttributesSegmentRoutingPolicyTypeK
	// HasSrv6Sid checks if Srv6Sid has been set in BgpAttributesSegmentRoutingPolicyTypeK
	HasSrv6Sid() bool
	// Srv6EndpointBehavior returns BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure, set in BgpAttributesSegmentRoutingPolicyTypeK.
	// BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure is configuration for optional SRv6 Endpoint Behavior and SID Structure. Summation of lengths for Locator Block, Locator Node,  Function, and Argument MUST be less than or equal to 128. - This is specified in draft-ietf-idr-sr-policy-safi-02 Section 2.4.4.2.4
	Srv6EndpointBehavior() BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	// SetSrv6EndpointBehavior assigns BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure provided by user to BgpAttributesSegmentRoutingPolicyTypeK.
	// BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure is configuration for optional SRv6 Endpoint Behavior and SID Structure. Summation of lengths for Locator Block, Locator Node,  Function, and Argument MUST be less than or equal to 128. - This is specified in draft-ietf-idr-sr-policy-safi-02 Section 2.4.4.2.4
	SetSrv6EndpointBehavior(value BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) BgpAttributesSegmentRoutingPolicyTypeK
	// HasSrv6EndpointBehavior checks if Srv6EndpointBehavior has been set in BgpAttributesSegmentRoutingPolicyTypeK
	HasSrv6EndpointBehavior() bool
	setNil()
}

// description is TBD
// Flags returns a BgpAttributesSegmentRoutingPolicyTypeFlags
func (obj *bgpAttributesSegmentRoutingPolicyTypeK) Flags() BgpAttributesSegmentRoutingPolicyTypeFlags {
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
func (obj *bgpAttributesSegmentRoutingPolicyTypeK) HasFlags() bool {
	return obj.obj.Flags != nil
}

// description is TBD
// SetFlags sets the BgpAttributesSegmentRoutingPolicyTypeFlags value in the BgpAttributesSegmentRoutingPolicyTypeK object
func (obj *bgpAttributesSegmentRoutingPolicyTypeK) SetFlags(value BgpAttributesSegmentRoutingPolicyTypeFlags) BgpAttributesSegmentRoutingPolicyTypeK {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// SR Algorithm identifier when A-Flag in on. If A-flag is not enabled, it should be set to 0 on transmission and ignored on receipt.
// SrAlgorithm returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicyTypeK) SrAlgorithm() uint32 {

	return *obj.obj.SrAlgorithm

}

// SR Algorithm identifier when A-Flag in on. If A-flag is not enabled, it should be set to 0 on transmission and ignored on receipt.
// SrAlgorithm returns a uint32
func (obj *bgpAttributesSegmentRoutingPolicyTypeK) HasSrAlgorithm() bool {
	return obj.obj.SrAlgorithm != nil
}

// SR Algorithm identifier when A-Flag in on. If A-flag is not enabled, it should be set to 0 on transmission and ignored on receipt.
// SetSrAlgorithm sets the uint32 value in the BgpAttributesSegmentRoutingPolicyTypeK object
func (obj *bgpAttributesSegmentRoutingPolicyTypeK) SetSrAlgorithm(value uint32) BgpAttributesSegmentRoutingPolicyTypeK {

	obj.obj.SrAlgorithm = &value
	return obj
}

// Local IPv6 Address.
// LocalIpv6Address returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeK) LocalIpv6Address() string {

	return *obj.obj.LocalIpv6Address

}

// Local IPv6 Address.
// LocalIpv6Address returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeK) HasLocalIpv6Address() bool {
	return obj.obj.LocalIpv6Address != nil
}

// Local IPv6 Address.
// SetLocalIpv6Address sets the string value in the BgpAttributesSegmentRoutingPolicyTypeK object
func (obj *bgpAttributesSegmentRoutingPolicyTypeK) SetLocalIpv6Address(value string) BgpAttributesSegmentRoutingPolicyTypeK {

	obj.obj.LocalIpv6Address = &value
	return obj
}

// Remote IPv6 Address.
// RemoteIpv6Address returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeK) RemoteIpv6Address() string {

	return *obj.obj.RemoteIpv6Address

}

// Remote IPv6 Address.
// RemoteIpv6Address returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeK) HasRemoteIpv6Address() bool {
	return obj.obj.RemoteIpv6Address != nil
}

// Remote IPv6 Address.
// SetRemoteIpv6Address sets the string value in the BgpAttributesSegmentRoutingPolicyTypeK object
func (obj *bgpAttributesSegmentRoutingPolicyTypeK) SetRemoteIpv6Address(value string) BgpAttributesSegmentRoutingPolicyTypeK {

	obj.obj.RemoteIpv6Address = &value
	return obj
}

// description is TBD
// Srv6Sid returns a BgpAttributesSidSrv6
func (obj *bgpAttributesSegmentRoutingPolicyTypeK) Srv6Sid() BgpAttributesSidSrv6 {
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
func (obj *bgpAttributesSegmentRoutingPolicyTypeK) HasSrv6Sid() bool {
	return obj.obj.Srv6Sid != nil
}

// description is TBD
// SetSrv6Sid sets the BgpAttributesSidSrv6 value in the BgpAttributesSegmentRoutingPolicyTypeK object
func (obj *bgpAttributesSegmentRoutingPolicyTypeK) SetSrv6Sid(value BgpAttributesSidSrv6) BgpAttributesSegmentRoutingPolicyTypeK {

	obj.srv6SidHolder = nil
	obj.obj.Srv6Sid = value.msg()

	return obj
}

// description is TBD
// Srv6EndpointBehavior returns a BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
func (obj *bgpAttributesSegmentRoutingPolicyTypeK) Srv6EndpointBehavior() BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure {
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
func (obj *bgpAttributesSegmentRoutingPolicyTypeK) HasSrv6EndpointBehavior() bool {
	return obj.obj.Srv6EndpointBehavior != nil
}

// description is TBD
// SetSrv6EndpointBehavior sets the BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure value in the BgpAttributesSegmentRoutingPolicyTypeK object
func (obj *bgpAttributesSegmentRoutingPolicyTypeK) SetSrv6EndpointBehavior(value BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) BgpAttributesSegmentRoutingPolicyTypeK {

	obj.srv6EndpointBehaviorHolder = nil
	obj.obj.Srv6EndpointBehavior = value.msg()

	return obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeK) validateObj(vObj *validation, set_default bool) {
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
				fmt.Sprintf("0 <= BgpAttributesSegmentRoutingPolicyTypeK.SrAlgorithm <= 255 but Got %d", *obj.obj.SrAlgorithm))
		}

	}

	if obj.obj.LocalIpv6Address != nil {

		err := obj.validateIpv6(obj.LocalIpv6Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpAttributesSegmentRoutingPolicyTypeK.LocalIpv6Address"))
		}

	}

	if obj.obj.RemoteIpv6Address != nil {

		err := obj.validateIpv6(obj.RemoteIpv6Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpAttributesSegmentRoutingPolicyTypeK.RemoteIpv6Address"))
		}

	}

	if obj.obj.Srv6Sid != nil {

		obj.Srv6Sid().validateObj(vObj, set_default)
	}

	if obj.obj.Srv6EndpointBehavior != nil {

		obj.Srv6EndpointBehavior().validateObj(vObj, set_default)
	}

}

func (obj *bgpAttributesSegmentRoutingPolicyTypeK) setDefault() {
	if obj.obj.SrAlgorithm == nil {
		obj.SetSrAlgorithm(0)
	}
	if obj.obj.LocalIpv6Address == nil {
		obj.SetLocalIpv6Address("0::0")
	}
	if obj.obj.RemoteIpv6Address == nil {
		obj.SetRemoteIpv6Address("0::0")
	}

}
