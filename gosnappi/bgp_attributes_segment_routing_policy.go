package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesSegmentRoutingPolicy *****
type bgpAttributesSegmentRoutingPolicy struct {
	validation
	obj                                *otg.BgpAttributesSegmentRoutingPolicy
	marshaller                         marshalBgpAttributesSegmentRoutingPolicy
	unMarshaller                       unMarshalBgpAttributesSegmentRoutingPolicy
	bindingSegmentIdentifierHolder     BgpAttributesBsid
	srv6BindingSegmentIdentifierHolder BgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter
	preferenceHolder                   BgpAttributesSrPolicyPreference
	priorityHolder                     BgpAttributesSrPolicyPriority
	policyNameHolder                   BgpAttributesSrPolicyPolicyName
	policyCandidateNameHolder          BgpAttributesSrPolicyPolicyCandidateName
	explicitNullLabelPolicyHolder      BgpAttributesSrPolicyExplicitNullPolicy
	segmentListHolder                  BgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter
}

func NewBgpAttributesSegmentRoutingPolicy() BgpAttributesSegmentRoutingPolicy {
	obj := bgpAttributesSegmentRoutingPolicy{obj: &otg.BgpAttributesSegmentRoutingPolicy{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesSegmentRoutingPolicy) msg() *otg.BgpAttributesSegmentRoutingPolicy {
	return obj.obj
}

func (obj *bgpAttributesSegmentRoutingPolicy) setMsg(msg *otg.BgpAttributesSegmentRoutingPolicy) BgpAttributesSegmentRoutingPolicy {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesSegmentRoutingPolicy struct {
	obj *bgpAttributesSegmentRoutingPolicy
}

type marshalBgpAttributesSegmentRoutingPolicy interface {
	// ToProto marshals BgpAttributesSegmentRoutingPolicy to protobuf object *otg.BgpAttributesSegmentRoutingPolicy
	ToProto() (*otg.BgpAttributesSegmentRoutingPolicy, error)
	// ToPbText marshals BgpAttributesSegmentRoutingPolicy to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesSegmentRoutingPolicy to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesSegmentRoutingPolicy to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpAttributesSegmentRoutingPolicy to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpAttributesSegmentRoutingPolicy struct {
	obj *bgpAttributesSegmentRoutingPolicy
}

type unMarshalBgpAttributesSegmentRoutingPolicy interface {
	// FromProto unmarshals BgpAttributesSegmentRoutingPolicy from protobuf object *otg.BgpAttributesSegmentRoutingPolicy
	FromProto(msg *otg.BgpAttributesSegmentRoutingPolicy) (BgpAttributesSegmentRoutingPolicy, error)
	// FromPbText unmarshals BgpAttributesSegmentRoutingPolicy from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesSegmentRoutingPolicy from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesSegmentRoutingPolicy from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesSegmentRoutingPolicy) Marshal() marshalBgpAttributesSegmentRoutingPolicy {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesSegmentRoutingPolicy{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesSegmentRoutingPolicy) Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicy {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesSegmentRoutingPolicy{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesSegmentRoutingPolicy) ToProto() (*otg.BgpAttributesSegmentRoutingPolicy, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesSegmentRoutingPolicy) FromProto(msg *otg.BgpAttributesSegmentRoutingPolicy) (BgpAttributesSegmentRoutingPolicy, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesSegmentRoutingPolicy) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicy) FromPbText(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicy) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicy) FromYaml(value string) error {
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

func (m *marshalbgpAttributesSegmentRoutingPolicy) ToJsonRaw() (string, error) {
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

func (m *marshalbgpAttributesSegmentRoutingPolicy) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesSegmentRoutingPolicy) FromJson(value string) error {
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

func (obj *bgpAttributesSegmentRoutingPolicy) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicy) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesSegmentRoutingPolicy) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesSegmentRoutingPolicy) Clone() (BgpAttributesSegmentRoutingPolicy, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesSegmentRoutingPolicy()
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

func (obj *bgpAttributesSegmentRoutingPolicy) setNil() {
	obj.bindingSegmentIdentifierHolder = nil
	obj.srv6BindingSegmentIdentifierHolder = nil
	obj.preferenceHolder = nil
	obj.priorityHolder = nil
	obj.policyNameHolder = nil
	obj.policyCandidateNameHolder = nil
	obj.explicitNullLabelPolicyHolder = nil
	obj.segmentListHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpAttributesSegmentRoutingPolicy is optional Segment Routing Policy information as defined in draft-ietf-idr-sr-policy-safi-02.
// This information is carried in TUNNEL_ENCAPSULATION attribute with type set to  SR Policy (15).
type BgpAttributesSegmentRoutingPolicy interface {
	Validation
	// msg marshals BgpAttributesSegmentRoutingPolicy to protobuf object *otg.BgpAttributesSegmentRoutingPolicy
	// and doesn't set defaults
	msg() *otg.BgpAttributesSegmentRoutingPolicy
	// setMsg unmarshals BgpAttributesSegmentRoutingPolicy from protobuf object *otg.BgpAttributesSegmentRoutingPolicy
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesSegmentRoutingPolicy) BgpAttributesSegmentRoutingPolicy
	// provides marshal interface
	Marshal() marshalBgpAttributesSegmentRoutingPolicy
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesSegmentRoutingPolicy
	// validate validates BgpAttributesSegmentRoutingPolicy
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesSegmentRoutingPolicy, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// BindingSegmentIdentifier returns BgpAttributesBsid, set in BgpAttributesSegmentRoutingPolicy.
	// BgpAttributesBsid is the Binding Segment Identifier is an optional sub-tlv of type 13 that can be sent with a SR Policy
	// Tunnel Encapsulation attribute.
	// When the active candidate path has a specified Binding Segment Identifier, the SR Policy uses that
	// BSID if this value (label in MPLS, IPv6 address in SRv6) is available.
	// - The format of the sub-TLV is defined in draft-ietf-idr-sr-policy-safi-02 Section 2.4.2 .
	// - It is recommended that if SRv6 Binding SID is desired to be signalled, the SRv6 Binding SID sub-TLV that enables
	// the specification of the SRv6 Endpoint Behavior should be used.
	BindingSegmentIdentifier() BgpAttributesBsid
	// SetBindingSegmentIdentifier assigns BgpAttributesBsid provided by user to BgpAttributesSegmentRoutingPolicy.
	// BgpAttributesBsid is the Binding Segment Identifier is an optional sub-tlv of type 13 that can be sent with a SR Policy
	// Tunnel Encapsulation attribute.
	// When the active candidate path has a specified Binding Segment Identifier, the SR Policy uses that
	// BSID if this value (label in MPLS, IPv6 address in SRv6) is available.
	// - The format of the sub-TLV is defined in draft-ietf-idr-sr-policy-safi-02 Section 2.4.2 .
	// - It is recommended that if SRv6 Binding SID is desired to be signalled, the SRv6 Binding SID sub-TLV that enables
	// the specification of the SRv6 Endpoint Behavior should be used.
	SetBindingSegmentIdentifier(value BgpAttributesBsid) BgpAttributesSegmentRoutingPolicy
	// HasBindingSegmentIdentifier checks if BindingSegmentIdentifier has been set in BgpAttributesSegmentRoutingPolicy
	HasBindingSegmentIdentifier() bool
	// Srv6BindingSegmentIdentifier returns BgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIterIter, set in BgpAttributesSegmentRoutingPolicy
	Srv6BindingSegmentIdentifier() BgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter
	// Preference returns BgpAttributesSrPolicyPreference, set in BgpAttributesSegmentRoutingPolicy.
	// BgpAttributesSrPolicyPreference is optional Preference sub-tlv (Type 12) is used to select the best candidate path for an SR Policy.
	// It is defined in Section 2.4.1 of draft-ietf-idr-sr-policy-safi-02 .
	Preference() BgpAttributesSrPolicyPreference
	// SetPreference assigns BgpAttributesSrPolicyPreference provided by user to BgpAttributesSegmentRoutingPolicy.
	// BgpAttributesSrPolicyPreference is optional Preference sub-tlv (Type 12) is used to select the best candidate path for an SR Policy.
	// It is defined in Section 2.4.1 of draft-ietf-idr-sr-policy-safi-02 .
	SetPreference(value BgpAttributesSrPolicyPreference) BgpAttributesSegmentRoutingPolicy
	// HasPreference checks if Preference has been set in BgpAttributesSegmentRoutingPolicy
	HasPreference() bool
	// Priority returns BgpAttributesSrPolicyPriority, set in BgpAttributesSegmentRoutingPolicy.
	// BgpAttributesSrPolicyPriority is optional Priority sub-tlv (Type 15) used to select the order in which policies should be re-computed.
	// - It is defined in Section 2.4.6 of draft-ietf-idr-sr-policy-safi-02 .
	Priority() BgpAttributesSrPolicyPriority
	// SetPriority assigns BgpAttributesSrPolicyPriority provided by user to BgpAttributesSegmentRoutingPolicy.
	// BgpAttributesSrPolicyPriority is optional Priority sub-tlv (Type 15) used to select the order in which policies should be re-computed.
	// - It is defined in Section 2.4.6 of draft-ietf-idr-sr-policy-safi-02 .
	SetPriority(value BgpAttributesSrPolicyPriority) BgpAttributesSegmentRoutingPolicy
	// HasPriority checks if Priority has been set in BgpAttributesSegmentRoutingPolicy
	HasPriority() bool
	// PolicyName returns BgpAttributesSrPolicyPolicyName, set in BgpAttributesSegmentRoutingPolicy.
	// BgpAttributesSrPolicyPolicyName is optional Policy Name sub-tlv (Type 130) which carries the symbolic name for the SR Policy for which the
	// candidate path is being advertised for debugging.
	// - It is defined in Section 2.4.8 of draft-ietf-idr-sr-policy-safi-02 .
	PolicyName() BgpAttributesSrPolicyPolicyName
	// SetPolicyName assigns BgpAttributesSrPolicyPolicyName provided by user to BgpAttributesSegmentRoutingPolicy.
	// BgpAttributesSrPolicyPolicyName is optional Policy Name sub-tlv (Type 130) which carries the symbolic name for the SR Policy for which the
	// candidate path is being advertised for debugging.
	// - It is defined in Section 2.4.8 of draft-ietf-idr-sr-policy-safi-02 .
	SetPolicyName(value BgpAttributesSrPolicyPolicyName) BgpAttributesSegmentRoutingPolicy
	// HasPolicyName checks if PolicyName has been set in BgpAttributesSegmentRoutingPolicy
	HasPolicyName() bool
	// PolicyCandidateName returns BgpAttributesSrPolicyPolicyCandidateName, set in BgpAttributesSegmentRoutingPolicy.
	// BgpAttributesSrPolicyPolicyCandidateName is optional Policy Candidate Path Name sub-tlv (Type 129) which carries the symbolic name for the SR Policy candidate path
	// for debugging.
	// - It is defined in Section 2.4.7 of draft-ietf-idr-sr-policy-safi-02 .
	PolicyCandidateName() BgpAttributesSrPolicyPolicyCandidateName
	// SetPolicyCandidateName assigns BgpAttributesSrPolicyPolicyCandidateName provided by user to BgpAttributesSegmentRoutingPolicy.
	// BgpAttributesSrPolicyPolicyCandidateName is optional Policy Candidate Path Name sub-tlv (Type 129) which carries the symbolic name for the SR Policy candidate path
	// for debugging.
	// - It is defined in Section 2.4.7 of draft-ietf-idr-sr-policy-safi-02 .
	SetPolicyCandidateName(value BgpAttributesSrPolicyPolicyCandidateName) BgpAttributesSegmentRoutingPolicy
	// HasPolicyCandidateName checks if PolicyCandidateName has been set in BgpAttributesSegmentRoutingPolicy
	HasPolicyCandidateName() bool
	// ExplicitNullLabelPolicy returns BgpAttributesSrPolicyExplicitNullPolicy, set in BgpAttributesSegmentRoutingPolicy.
	// BgpAttributesSrPolicyExplicitNullPolicy is this is an optional sub-tlv (Type 14) which indicates whether an Explicit NULL Label must be pushed on an unlabeled IP
	// packet before other labels for IPv4 or IPv6 flows.
	// - It is defined in Section 2.4.5 of draft-ietf-idr-sr-policy-safi-02.
	ExplicitNullLabelPolicy() BgpAttributesSrPolicyExplicitNullPolicy
	// SetExplicitNullLabelPolicy assigns BgpAttributesSrPolicyExplicitNullPolicy provided by user to BgpAttributesSegmentRoutingPolicy.
	// BgpAttributesSrPolicyExplicitNullPolicy is this is an optional sub-tlv (Type 14) which indicates whether an Explicit NULL Label must be pushed on an unlabeled IP
	// packet before other labels for IPv4 or IPv6 flows.
	// - It is defined in Section 2.4.5 of draft-ietf-idr-sr-policy-safi-02.
	SetExplicitNullLabelPolicy(value BgpAttributesSrPolicyExplicitNullPolicy) BgpAttributesSegmentRoutingPolicy
	// HasExplicitNullLabelPolicy checks if ExplicitNullLabelPolicy has been set in BgpAttributesSegmentRoutingPolicy
	HasExplicitNullLabelPolicy() bool
	// SegmentList returns BgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIterIter, set in BgpAttributesSegmentRoutingPolicy
	SegmentList() BgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter
	setNil()
}

// description is TBD
// BindingSegmentIdentifier returns a BgpAttributesBsid
func (obj *bgpAttributesSegmentRoutingPolicy) BindingSegmentIdentifier() BgpAttributesBsid {
	if obj.obj.BindingSegmentIdentifier == nil {
		obj.obj.BindingSegmentIdentifier = NewBgpAttributesBsid().msg()
	}
	if obj.bindingSegmentIdentifierHolder == nil {
		obj.bindingSegmentIdentifierHolder = &bgpAttributesBsid{obj: obj.obj.BindingSegmentIdentifier}
	}
	return obj.bindingSegmentIdentifierHolder
}

// description is TBD
// BindingSegmentIdentifier returns a BgpAttributesBsid
func (obj *bgpAttributesSegmentRoutingPolicy) HasBindingSegmentIdentifier() bool {
	return obj.obj.BindingSegmentIdentifier != nil
}

// description is TBD
// SetBindingSegmentIdentifier sets the BgpAttributesBsid value in the BgpAttributesSegmentRoutingPolicy object
func (obj *bgpAttributesSegmentRoutingPolicy) SetBindingSegmentIdentifier(value BgpAttributesBsid) BgpAttributesSegmentRoutingPolicy {

	obj.bindingSegmentIdentifierHolder = nil
	obj.obj.BindingSegmentIdentifier = value.msg()

	return obj
}

// The SRv6 Binding SID sub-TLV is an optional sub-TLV of type 20 that is used to signal the SRv6 Binding SID
// related information of an SR Policy candidate path.
// - More than one SRv6 Binding SID sub-TLVs MAY be signaled in the same SR Policy encoding to indicate one
// or more SRv6 SIDs, each with potentially different SRv6 Endpoint Behaviors to be instantiated.
// - The format of the sub-TLV is defined in draft-ietf-idr-sr-policy-safi-02 Section 2.4.3 .
// Srv6BindingSegmentIdentifier returns a []BgpAttributesSrv6Bsid
func (obj *bgpAttributesSegmentRoutingPolicy) Srv6BindingSegmentIdentifier() BgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter {
	if len(obj.obj.Srv6BindingSegmentIdentifier) == 0 {
		obj.obj.Srv6BindingSegmentIdentifier = []*otg.BgpAttributesSrv6Bsid{}
	}
	if obj.srv6BindingSegmentIdentifierHolder == nil {
		obj.srv6BindingSegmentIdentifierHolder = newBgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter(&obj.obj.Srv6BindingSegmentIdentifier).setMsg(obj)
	}
	return obj.srv6BindingSegmentIdentifierHolder
}

type bgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter struct {
	obj                        *bgpAttributesSegmentRoutingPolicy
	bgpAttributesSrv6BsidSlice []BgpAttributesSrv6Bsid
	fieldPtr                   *[]*otg.BgpAttributesSrv6Bsid
}

func newBgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter(ptr *[]*otg.BgpAttributesSrv6Bsid) BgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter {
	return &bgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter{fieldPtr: ptr}
}

type BgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter interface {
	setMsg(*bgpAttributesSegmentRoutingPolicy) BgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter
	Items() []BgpAttributesSrv6Bsid
	Add() BgpAttributesSrv6Bsid
	Append(items ...BgpAttributesSrv6Bsid) BgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter
	Set(index int, newObj BgpAttributesSrv6Bsid) BgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter
	Clear() BgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter
	clearHolderSlice() BgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter
	appendHolderSlice(item BgpAttributesSrv6Bsid) BgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter
}

func (obj *bgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter) setMsg(msg *bgpAttributesSegmentRoutingPolicy) BgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpAttributesSrv6Bsid{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter) Items() []BgpAttributesSrv6Bsid {
	return obj.bgpAttributesSrv6BsidSlice
}

func (obj *bgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter) Add() BgpAttributesSrv6Bsid {
	newObj := &otg.BgpAttributesSrv6Bsid{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpAttributesSrv6Bsid{obj: newObj}
	newLibObj.setDefault()
	obj.bgpAttributesSrv6BsidSlice = append(obj.bgpAttributesSrv6BsidSlice, newLibObj)
	return newLibObj
}

func (obj *bgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter) Append(items ...BgpAttributesSrv6Bsid) BgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpAttributesSrv6BsidSlice = append(obj.bgpAttributesSrv6BsidSlice, item)
	}
	return obj
}

func (obj *bgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter) Set(index int, newObj BgpAttributesSrv6Bsid) BgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpAttributesSrv6BsidSlice[index] = newObj
	return obj
}
func (obj *bgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter) Clear() BgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpAttributesSrv6Bsid{}
		obj.bgpAttributesSrv6BsidSlice = []BgpAttributesSrv6Bsid{}
	}
	return obj
}
func (obj *bgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter) clearHolderSlice() BgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter {
	if len(obj.bgpAttributesSrv6BsidSlice) > 0 {
		obj.bgpAttributesSrv6BsidSlice = []BgpAttributesSrv6Bsid{}
	}
	return obj
}
func (obj *bgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter) appendHolderSlice(item BgpAttributesSrv6Bsid) BgpAttributesSegmentRoutingPolicyBgpAttributesSrv6BsidIter {
	obj.bgpAttributesSrv6BsidSlice = append(obj.bgpAttributesSrv6BsidSlice, item)
	return obj
}

// description is TBD
// Preference returns a BgpAttributesSrPolicyPreference
func (obj *bgpAttributesSegmentRoutingPolicy) Preference() BgpAttributesSrPolicyPreference {
	if obj.obj.Preference == nil {
		obj.obj.Preference = NewBgpAttributesSrPolicyPreference().msg()
	}
	if obj.preferenceHolder == nil {
		obj.preferenceHolder = &bgpAttributesSrPolicyPreference{obj: obj.obj.Preference}
	}
	return obj.preferenceHolder
}

// description is TBD
// Preference returns a BgpAttributesSrPolicyPreference
func (obj *bgpAttributesSegmentRoutingPolicy) HasPreference() bool {
	return obj.obj.Preference != nil
}

// description is TBD
// SetPreference sets the BgpAttributesSrPolicyPreference value in the BgpAttributesSegmentRoutingPolicy object
func (obj *bgpAttributesSegmentRoutingPolicy) SetPreference(value BgpAttributesSrPolicyPreference) BgpAttributesSegmentRoutingPolicy {

	obj.preferenceHolder = nil
	obj.obj.Preference = value.msg()

	return obj
}

// description is TBD
// Priority returns a BgpAttributesSrPolicyPriority
func (obj *bgpAttributesSegmentRoutingPolicy) Priority() BgpAttributesSrPolicyPriority {
	if obj.obj.Priority == nil {
		obj.obj.Priority = NewBgpAttributesSrPolicyPriority().msg()
	}
	if obj.priorityHolder == nil {
		obj.priorityHolder = &bgpAttributesSrPolicyPriority{obj: obj.obj.Priority}
	}
	return obj.priorityHolder
}

// description is TBD
// Priority returns a BgpAttributesSrPolicyPriority
func (obj *bgpAttributesSegmentRoutingPolicy) HasPriority() bool {
	return obj.obj.Priority != nil
}

// description is TBD
// SetPriority sets the BgpAttributesSrPolicyPriority value in the BgpAttributesSegmentRoutingPolicy object
func (obj *bgpAttributesSegmentRoutingPolicy) SetPriority(value BgpAttributesSrPolicyPriority) BgpAttributesSegmentRoutingPolicy {

	obj.priorityHolder = nil
	obj.obj.Priority = value.msg()

	return obj
}

// description is TBD
// PolicyName returns a BgpAttributesSrPolicyPolicyName
func (obj *bgpAttributesSegmentRoutingPolicy) PolicyName() BgpAttributesSrPolicyPolicyName {
	if obj.obj.PolicyName == nil {
		obj.obj.PolicyName = NewBgpAttributesSrPolicyPolicyName().msg()
	}
	if obj.policyNameHolder == nil {
		obj.policyNameHolder = &bgpAttributesSrPolicyPolicyName{obj: obj.obj.PolicyName}
	}
	return obj.policyNameHolder
}

// description is TBD
// PolicyName returns a BgpAttributesSrPolicyPolicyName
func (obj *bgpAttributesSegmentRoutingPolicy) HasPolicyName() bool {
	return obj.obj.PolicyName != nil
}

// description is TBD
// SetPolicyName sets the BgpAttributesSrPolicyPolicyName value in the BgpAttributesSegmentRoutingPolicy object
func (obj *bgpAttributesSegmentRoutingPolicy) SetPolicyName(value BgpAttributesSrPolicyPolicyName) BgpAttributesSegmentRoutingPolicy {

	obj.policyNameHolder = nil
	obj.obj.PolicyName = value.msg()

	return obj
}

// description is TBD
// PolicyCandidateName returns a BgpAttributesSrPolicyPolicyCandidateName
func (obj *bgpAttributesSegmentRoutingPolicy) PolicyCandidateName() BgpAttributesSrPolicyPolicyCandidateName {
	if obj.obj.PolicyCandidateName == nil {
		obj.obj.PolicyCandidateName = NewBgpAttributesSrPolicyPolicyCandidateName().msg()
	}
	if obj.policyCandidateNameHolder == nil {
		obj.policyCandidateNameHolder = &bgpAttributesSrPolicyPolicyCandidateName{obj: obj.obj.PolicyCandidateName}
	}
	return obj.policyCandidateNameHolder
}

// description is TBD
// PolicyCandidateName returns a BgpAttributesSrPolicyPolicyCandidateName
func (obj *bgpAttributesSegmentRoutingPolicy) HasPolicyCandidateName() bool {
	return obj.obj.PolicyCandidateName != nil
}

// description is TBD
// SetPolicyCandidateName sets the BgpAttributesSrPolicyPolicyCandidateName value in the BgpAttributesSegmentRoutingPolicy object
func (obj *bgpAttributesSegmentRoutingPolicy) SetPolicyCandidateName(value BgpAttributesSrPolicyPolicyCandidateName) BgpAttributesSegmentRoutingPolicy {

	obj.policyCandidateNameHolder = nil
	obj.obj.PolicyCandidateName = value.msg()

	return obj
}

// description is TBD
// ExplicitNullLabelPolicy returns a BgpAttributesSrPolicyExplicitNullPolicy
func (obj *bgpAttributesSegmentRoutingPolicy) ExplicitNullLabelPolicy() BgpAttributesSrPolicyExplicitNullPolicy {
	if obj.obj.ExplicitNullLabelPolicy == nil {
		obj.obj.ExplicitNullLabelPolicy = NewBgpAttributesSrPolicyExplicitNullPolicy().msg()
	}
	if obj.explicitNullLabelPolicyHolder == nil {
		obj.explicitNullLabelPolicyHolder = &bgpAttributesSrPolicyExplicitNullPolicy{obj: obj.obj.ExplicitNullLabelPolicy}
	}
	return obj.explicitNullLabelPolicyHolder
}

// description is TBD
// ExplicitNullLabelPolicy returns a BgpAttributesSrPolicyExplicitNullPolicy
func (obj *bgpAttributesSegmentRoutingPolicy) HasExplicitNullLabelPolicy() bool {
	return obj.obj.ExplicitNullLabelPolicy != nil
}

// description is TBD
// SetExplicitNullLabelPolicy sets the BgpAttributesSrPolicyExplicitNullPolicy value in the BgpAttributesSegmentRoutingPolicy object
func (obj *bgpAttributesSegmentRoutingPolicy) SetExplicitNullLabelPolicy(value BgpAttributesSrPolicyExplicitNullPolicy) BgpAttributesSegmentRoutingPolicy {

	obj.explicitNullLabelPolicyHolder = nil
	obj.obj.ExplicitNullLabelPolicy = value.msg()

	return obj
}

// description is TBD
// SegmentList returns a []BgpAttributesSrPolicySegmentList
func (obj *bgpAttributesSegmentRoutingPolicy) SegmentList() BgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter {
	if len(obj.obj.SegmentList) == 0 {
		obj.obj.SegmentList = []*otg.BgpAttributesSrPolicySegmentList{}
	}
	if obj.segmentListHolder == nil {
		obj.segmentListHolder = newBgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter(&obj.obj.SegmentList).setMsg(obj)
	}
	return obj.segmentListHolder
}

type bgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter struct {
	obj                                   *bgpAttributesSegmentRoutingPolicy
	bgpAttributesSrPolicySegmentListSlice []BgpAttributesSrPolicySegmentList
	fieldPtr                              *[]*otg.BgpAttributesSrPolicySegmentList
}

func newBgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter(ptr *[]*otg.BgpAttributesSrPolicySegmentList) BgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter {
	return &bgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter{fieldPtr: ptr}
}

type BgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter interface {
	setMsg(*bgpAttributesSegmentRoutingPolicy) BgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter
	Items() []BgpAttributesSrPolicySegmentList
	Add() BgpAttributesSrPolicySegmentList
	Append(items ...BgpAttributesSrPolicySegmentList) BgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter
	Set(index int, newObj BgpAttributesSrPolicySegmentList) BgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter
	Clear() BgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter
	clearHolderSlice() BgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter
	appendHolderSlice(item BgpAttributesSrPolicySegmentList) BgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter
}

func (obj *bgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter) setMsg(msg *bgpAttributesSegmentRoutingPolicy) BgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpAttributesSrPolicySegmentList{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter) Items() []BgpAttributesSrPolicySegmentList {
	return obj.bgpAttributesSrPolicySegmentListSlice
}

func (obj *bgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter) Add() BgpAttributesSrPolicySegmentList {
	newObj := &otg.BgpAttributesSrPolicySegmentList{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpAttributesSrPolicySegmentList{obj: newObj}
	newLibObj.setDefault()
	obj.bgpAttributesSrPolicySegmentListSlice = append(obj.bgpAttributesSrPolicySegmentListSlice, newLibObj)
	return newLibObj
}

func (obj *bgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter) Append(items ...BgpAttributesSrPolicySegmentList) BgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpAttributesSrPolicySegmentListSlice = append(obj.bgpAttributesSrPolicySegmentListSlice, item)
	}
	return obj
}

func (obj *bgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter) Set(index int, newObj BgpAttributesSrPolicySegmentList) BgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpAttributesSrPolicySegmentListSlice[index] = newObj
	return obj
}
func (obj *bgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter) Clear() BgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpAttributesSrPolicySegmentList{}
		obj.bgpAttributesSrPolicySegmentListSlice = []BgpAttributesSrPolicySegmentList{}
	}
	return obj
}
func (obj *bgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter) clearHolderSlice() BgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter {
	if len(obj.bgpAttributesSrPolicySegmentListSlice) > 0 {
		obj.bgpAttributesSrPolicySegmentListSlice = []BgpAttributesSrPolicySegmentList{}
	}
	return obj
}
func (obj *bgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter) appendHolderSlice(item BgpAttributesSrPolicySegmentList) BgpAttributesSegmentRoutingPolicyBgpAttributesSrPolicySegmentListIter {
	obj.bgpAttributesSrPolicySegmentListSlice = append(obj.bgpAttributesSrPolicySegmentListSlice, item)
	return obj
}

func (obj *bgpAttributesSegmentRoutingPolicy) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.BindingSegmentIdentifier != nil {

		obj.BindingSegmentIdentifier().validateObj(vObj, set_default)
	}

	if len(obj.obj.Srv6BindingSegmentIdentifier) != 0 {

		if set_default {
			obj.Srv6BindingSegmentIdentifier().clearHolderSlice()
			for _, item := range obj.obj.Srv6BindingSegmentIdentifier {
				obj.Srv6BindingSegmentIdentifier().appendHolderSlice(&bgpAttributesSrv6Bsid{obj: item})
			}
		}
		for _, item := range obj.Srv6BindingSegmentIdentifier().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.Preference != nil {

		obj.Preference().validateObj(vObj, set_default)
	}

	if obj.obj.Priority != nil {

		obj.Priority().validateObj(vObj, set_default)
	}

	if obj.obj.PolicyName != nil {

		obj.PolicyName().validateObj(vObj, set_default)
	}

	if obj.obj.PolicyCandidateName != nil {

		obj.PolicyCandidateName().validateObj(vObj, set_default)
	}

	if obj.obj.ExplicitNullLabelPolicy != nil {

		obj.ExplicitNullLabelPolicy().validateObj(vObj, set_default)
	}

	if len(obj.obj.SegmentList) != 0 {

		if set_default {
			obj.SegmentList().clearHolderSlice()
			for _, item := range obj.obj.SegmentList {
				obj.SegmentList().appendHolderSlice(&bgpAttributesSrPolicySegmentList{obj: item})
			}
		}
		for _, item := range obj.SegmentList().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *bgpAttributesSegmentRoutingPolicy) setDefault() {

}
