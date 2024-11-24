package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesSegmentRoutingPolicyTypeB *****
type bgpAttributesSegmentRoutingPolicyTypeB struct {
	validation
	obj                        *otg.BgpAttributesSegmentRoutingPolicyTypeB
	marshaller                 marshalBgpAttributesSegmentRoutingPolicyTypeB
	unMarshaller               unMarshalBgpAttributesSegmentRoutingPolicyTypeB
	flagsHolder                BgpAttributesSegmentRoutingPolicyTypeFlags
	srv6EndpointBehaviorHolder BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
}

func NewBgpAttributesSegmentRoutingPolicyTypeB() BgpAttributesSegmentRoutingPolicyTypeB {
	obj := bgpAttributesSegmentRoutingPolicyTypeB{obj: &otg.BgpAttributesSegmentRoutingPolicyTypeB{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeB) msg() *otg.BgpAttributesSegmentRoutingPolicyTypeB {
	return obj.obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeB) setMsg(msg *otg.BgpAttributesSegmentRoutingPolicyTypeB) BgpAttributesSegmentRoutingPolicyTypeB {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesSegmentRoutingPolicyTypeB struct {
	obj *bgpAttributesSegmentRoutingPolicyTypeB
}

type marshalBgpAttributesSegmentRoutingPolicyTypeB interface {
	// ToProto marshals BgpAttributesSegmentRoutingPolicyTypeB to protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeB
	ToProto() (*otg.BgpAttributesSegmentRoutingPolicyTypeB, error)
	// ToPbText marshals BgpAttributesSegmentRoutingPolicyTypeB to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesSegmentRoutingPolicyTypeB to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesSegmentRoutingPolicyTypeB to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpAttributesSegmentRoutingPolicyTypeB to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpAttributesSegmentRoutingPolicyTypeB struct {
	obj *bgpAttributesSegmentRoutingPolicyTypeB
}

type unMarshalBgpAttributesSegmentRoutingPolicyTypeB interface {
	// FromProto unmarshals BgpAttributesSegmentRoutingPolicyTypeB from protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeB
	FromProto(msg *otg.BgpAttributesSegmentRoutingPolicyTypeB) (BgpAttributesSegmentRoutingPolicyTypeB, error)
	// FromPbText unmarshals BgpAttributesSegmentRoutingPolicyTypeB from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesSegmentRoutingPolicyTypeB from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesSegmentRoutingPolicyTypeB from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeB) Marshal() marshalBgpAttributesSegmentRoutingPolicyTypeB {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesSegmentRoutingPolicyTypeB{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeB) Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicyTypeB {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesSegmentRoutingPolicyTypeB{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeB) ToProto() (*otg.BgpAttributesSegmentRoutingPolicyTypeB, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeB) FromProto(msg *otg.BgpAttributesSegmentRoutingPolicyTypeB) (BgpAttributesSegmentRoutingPolicyTypeB, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeB) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeB) FromPbText(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeB) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeB) FromYaml(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeB) ToJsonRaw() (string, error) {
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

func (m *marshalbgpAttributesSegmentRoutingPolicyTypeB) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicyTypeB) FromJson(value string) error {
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

func (obj *bgpAttributesSegmentRoutingPolicyTypeB) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeB) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeB) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeB) Clone() (BgpAttributesSegmentRoutingPolicyTypeB, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesSegmentRoutingPolicyTypeB()
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

func (obj *bgpAttributesSegmentRoutingPolicyTypeB) setNil() {
	obj.flagsHolder = nil
	obj.srv6EndpointBehaviorHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpAttributesSegmentRoutingPolicyTypeB is type B: SID only, in the form of IPv6 address.
// It is encoded as a Segment of Type 13 in the SEGMENT_LIST sub-tlv.
type BgpAttributesSegmentRoutingPolicyTypeB interface {
	Validation
	// msg marshals BgpAttributesSegmentRoutingPolicyTypeB to protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeB
	// and doesn't set defaults
	msg() *otg.BgpAttributesSegmentRoutingPolicyTypeB
	// setMsg unmarshals BgpAttributesSegmentRoutingPolicyTypeB from protobuf object *otg.BgpAttributesSegmentRoutingPolicyTypeB
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesSegmentRoutingPolicyTypeB) BgpAttributesSegmentRoutingPolicyTypeB
	// provides marshal interface
	Marshal() marshalBgpAttributesSegmentRoutingPolicyTypeB
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicyTypeB
	// validate validates BgpAttributesSegmentRoutingPolicyTypeB
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesSegmentRoutingPolicyTypeB, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns BgpAttributesSegmentRoutingPolicyTypeFlags, set in BgpAttributesSegmentRoutingPolicyTypeB.
	// BgpAttributesSegmentRoutingPolicyTypeFlags is flags for each Segment in SEGMENT_LIST sub-tlv.
	// - V-flag. Indicates verification is enabled. See section 5, of https://datatracker.ietf.org/doc/html/rfc9256
	// - A-flag. Indicates presence of SR Algorithm field applicable to Segment Types C, D , I , J and K .
	// - B-Flag. Indicates presence of SRv6 Endpoint Behavior and SID Structure encoding applicable to Segment Types B , I , J and K .
	// - S-Flag: This flag, when set, indicates the presence of the SR-MPLS or SRv6 SID depending on the segment type. (draft-ietf-idr-bgp-sr-segtypes-ext-03 Section 2.10).
	// This flag is applicable for Segment Types C, D, E, F, G, H, I, J, and K.
	Flags() BgpAttributesSegmentRoutingPolicyTypeFlags
	// SetFlags assigns BgpAttributesSegmentRoutingPolicyTypeFlags provided by user to BgpAttributesSegmentRoutingPolicyTypeB.
	// BgpAttributesSegmentRoutingPolicyTypeFlags is flags for each Segment in SEGMENT_LIST sub-tlv.
	// - V-flag. Indicates verification is enabled. See section 5, of https://datatracker.ietf.org/doc/html/rfc9256
	// - A-flag. Indicates presence of SR Algorithm field applicable to Segment Types C, D , I , J and K .
	// - B-Flag. Indicates presence of SRv6 Endpoint Behavior and SID Structure encoding applicable to Segment Types B , I , J and K .
	// - S-Flag: This flag, when set, indicates the presence of the SR-MPLS or SRv6 SID depending on the segment type. (draft-ietf-idr-bgp-sr-segtypes-ext-03 Section 2.10).
	// This flag is applicable for Segment Types C, D, E, F, G, H, I, J, and K.
	SetFlags(value BgpAttributesSegmentRoutingPolicyTypeFlags) BgpAttributesSegmentRoutingPolicyTypeB
	// HasFlags checks if Flags has been set in BgpAttributesSegmentRoutingPolicyTypeB
	HasFlags() bool
	// Srv6Sid returns string, set in BgpAttributesSegmentRoutingPolicyTypeB.
	Srv6Sid() string
	// SetSrv6Sid assigns string provided by user to BgpAttributesSegmentRoutingPolicyTypeB
	SetSrv6Sid(value string) BgpAttributesSegmentRoutingPolicyTypeB
	// HasSrv6Sid checks if Srv6Sid has been set in BgpAttributesSegmentRoutingPolicyTypeB
	HasSrv6Sid() bool
	// Srv6EndpointBehavior returns BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure, set in BgpAttributesSegmentRoutingPolicyTypeB.
	// BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure is configuration for optional SRv6 Endpoint Behavior and SID Structure. Summation of lengths for Locator Block, Locator Node,  Function, and Argument MUST be less than or equal to 128. - This is specified in draft-ietf-idr-sr-policy-safi-02 Section 2.4.4.2.4
	Srv6EndpointBehavior() BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
	// SetSrv6EndpointBehavior assigns BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure provided by user to BgpAttributesSegmentRoutingPolicyTypeB.
	// BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure is configuration for optional SRv6 Endpoint Behavior and SID Structure. Summation of lengths for Locator Block, Locator Node,  Function, and Argument MUST be less than or equal to 128. - This is specified in draft-ietf-idr-sr-policy-safi-02 Section 2.4.4.2.4
	SetSrv6EndpointBehavior(value BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) BgpAttributesSegmentRoutingPolicyTypeB
	// HasSrv6EndpointBehavior checks if Srv6EndpointBehavior has been set in BgpAttributesSegmentRoutingPolicyTypeB
	HasSrv6EndpointBehavior() bool
	setNil()
}

// description is TBD
// Flags returns a BgpAttributesSegmentRoutingPolicyTypeFlags
func (obj *bgpAttributesSegmentRoutingPolicyTypeB) Flags() BgpAttributesSegmentRoutingPolicyTypeFlags {
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
func (obj *bgpAttributesSegmentRoutingPolicyTypeB) HasFlags() bool {
	return obj.obj.Flags != nil
}

// description is TBD
// SetFlags sets the BgpAttributesSegmentRoutingPolicyTypeFlags value in the BgpAttributesSegmentRoutingPolicyTypeB object
func (obj *bgpAttributesSegmentRoutingPolicyTypeB) SetFlags(value BgpAttributesSegmentRoutingPolicyTypeFlags) BgpAttributesSegmentRoutingPolicyTypeB {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// SRv6 SID.
// Srv6Sid returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeB) Srv6Sid() string {

	return *obj.obj.Srv6Sid

}

// SRv6 SID.
// Srv6Sid returns a string
func (obj *bgpAttributesSegmentRoutingPolicyTypeB) HasSrv6Sid() bool {
	return obj.obj.Srv6Sid != nil
}

// SRv6 SID.
// SetSrv6Sid sets the string value in the BgpAttributesSegmentRoutingPolicyTypeB object
func (obj *bgpAttributesSegmentRoutingPolicyTypeB) SetSrv6Sid(value string) BgpAttributesSegmentRoutingPolicyTypeB {

	obj.obj.Srv6Sid = &value
	return obj
}

// description is TBD
// Srv6EndpointBehavior returns a BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure
func (obj *bgpAttributesSegmentRoutingPolicyTypeB) Srv6EndpointBehavior() BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure {
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
func (obj *bgpAttributesSegmentRoutingPolicyTypeB) HasSrv6EndpointBehavior() bool {
	return obj.obj.Srv6EndpointBehavior != nil
}

// description is TBD
// SetSrv6EndpointBehavior sets the BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure value in the BgpAttributesSegmentRoutingPolicyTypeB object
func (obj *bgpAttributesSegmentRoutingPolicyTypeB) SetSrv6EndpointBehavior(value BgpAttributesSegmentRoutingPolicySRv6SIDEndpointBehaviorAndStructure) BgpAttributesSegmentRoutingPolicyTypeB {

	obj.srv6EndpointBehaviorHolder = nil
	obj.obj.Srv6EndpointBehavior = value.msg()

	return obj
}

func (obj *bgpAttributesSegmentRoutingPolicyTypeB) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(vObj, set_default)
	}

	if obj.obj.Srv6Sid != nil {

		err := obj.validateIpv6(obj.Srv6Sid())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpAttributesSegmentRoutingPolicyTypeB.Srv6Sid"))
		}

	}

	if obj.obj.Srv6EndpointBehavior != nil {

		obj.Srv6EndpointBehavior().validateObj(vObj, set_default)
	}

}

func (obj *bgpAttributesSegmentRoutingPolicyTypeB) setDefault() {
	if obj.obj.Srv6Sid == nil {
		obj.SetSrv6Sid("0::0")
	}

}
