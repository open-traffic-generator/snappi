package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpV6EviVxlan *****
type bgpV6EviVxlan struct {
	validation
	obj                       *otg.BgpV6EviVxlan
	marshaller                marshalBgpV6EviVxlan
	unMarshaller              unMarshalBgpV6EviVxlan
	broadcastDomainsHolder    BgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter
	routeDistinguisherHolder  BgpRouteDistinguisher
	routeTargetExportHolder   BgpV6EviVxlanBgpRouteTargetIter
	routeTargetImportHolder   BgpV6EviVxlanBgpRouteTargetIter
	l3RouteTargetExportHolder BgpV6EviVxlanBgpRouteTargetIter
	l3RouteTargetImportHolder BgpV6EviVxlanBgpRouteTargetIter
	advancedHolder            BgpRouteAdvanced
	communitiesHolder         BgpV6EviVxlanBgpCommunityIter
	extCommunitiesHolder      BgpV6EviVxlanBgpExtCommunityIter
	asPathHolder              BgpAsPath
}

func NewBgpV6EviVxlan() BgpV6EviVxlan {
	obj := bgpV6EviVxlan{obj: &otg.BgpV6EviVxlan{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpV6EviVxlan) msg() *otg.BgpV6EviVxlan {
	return obj.obj
}

func (obj *bgpV6EviVxlan) setMsg(msg *otg.BgpV6EviVxlan) BgpV6EviVxlan {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpV6EviVxlan struct {
	obj *bgpV6EviVxlan
}

type marshalBgpV6EviVxlan interface {
	// ToProto marshals BgpV6EviVxlan to protobuf object *otg.BgpV6EviVxlan
	ToProto() (*otg.BgpV6EviVxlan, error)
	// ToPbText marshals BgpV6EviVxlan to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpV6EviVxlan to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpV6EviVxlan to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpV6EviVxlan to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpV6EviVxlan struct {
	obj *bgpV6EviVxlan
}

type unMarshalBgpV6EviVxlan interface {
	// FromProto unmarshals BgpV6EviVxlan from protobuf object *otg.BgpV6EviVxlan
	FromProto(msg *otg.BgpV6EviVxlan) (BgpV6EviVxlan, error)
	// FromPbText unmarshals BgpV6EviVxlan from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpV6EviVxlan from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpV6EviVxlan from JSON text
	FromJson(value string) error
}

func (obj *bgpV6EviVxlan) Marshal() marshalBgpV6EviVxlan {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpV6EviVxlan{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpV6EviVxlan) Unmarshal() unMarshalBgpV6EviVxlan {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpV6EviVxlan{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpV6EviVxlan) ToProto() (*otg.BgpV6EviVxlan, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpV6EviVxlan) FromProto(msg *otg.BgpV6EviVxlan) (BgpV6EviVxlan, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpV6EviVxlan) ToPbText() (string, error) {
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

func (m *unMarshalbgpV6EviVxlan) FromPbText(value string) error {
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

func (m *marshalbgpV6EviVxlan) ToYaml() (string, error) {
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

func (m *unMarshalbgpV6EviVxlan) FromYaml(value string) error {
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

func (m *marshalbgpV6EviVxlan) ToJsonRaw() (string, error) {
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

func (m *marshalbgpV6EviVxlan) ToJson() (string, error) {
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

func (m *unMarshalbgpV6EviVxlan) FromJson(value string) error {
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

func (obj *bgpV6EviVxlan) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpV6EviVxlan) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpV6EviVxlan) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpV6EviVxlan) Clone() (BgpV6EviVxlan, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpV6EviVxlan()
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

func (obj *bgpV6EviVxlan) setNil() {
	obj.broadcastDomainsHolder = nil
	obj.routeDistinguisherHolder = nil
	obj.routeTargetExportHolder = nil
	obj.routeTargetImportHolder = nil
	obj.l3RouteTargetExportHolder = nil
	obj.l3RouteTargetImportHolder = nil
	obj.advancedHolder = nil
	obj.communitiesHolder = nil
	obj.extCommunitiesHolder = nil
	obj.asPathHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpV6EviVxlan is configuration for BGP EVPN EVI. Advertises following routes -
//
// # Type 3 - Inclusive Multicast Ethernet Tag Route
//
// Type 1 -  Ethernet Auto-discovery Route (Per EVI)
//
// Type 1 -  Ethernet Auto-discovery Route (Per ES)
type BgpV6EviVxlan interface {
	Validation
	// msg marshals BgpV6EviVxlan to protobuf object *otg.BgpV6EviVxlan
	// and doesn't set defaults
	msg() *otg.BgpV6EviVxlan
	// setMsg unmarshals BgpV6EviVxlan from protobuf object *otg.BgpV6EviVxlan
	// and doesn't set defaults
	setMsg(*otg.BgpV6EviVxlan) BgpV6EviVxlan
	// provides marshal interface
	Marshal() marshalBgpV6EviVxlan
	// provides unmarshal interface
	Unmarshal() unMarshalBgpV6EviVxlan
	// validate validates BgpV6EviVxlan
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpV6EviVxlan, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// BroadcastDomains returns BgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIterIter, set in BgpV6EviVxlan
	BroadcastDomains() BgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter
	// ReplicationType returns BgpV6EviVxlanReplicationTypeEnum, set in BgpV6EviVxlan
	ReplicationType() BgpV6EviVxlanReplicationTypeEnum
	// SetReplicationType assigns BgpV6EviVxlanReplicationTypeEnum provided by user to BgpV6EviVxlan
	SetReplicationType(value BgpV6EviVxlanReplicationTypeEnum) BgpV6EviVxlan
	// HasReplicationType checks if ReplicationType has been set in BgpV6EviVxlan
	HasReplicationType() bool
	// PmsiLabel returns uint32, set in BgpV6EviVxlan.
	PmsiLabel() uint32
	// SetPmsiLabel assigns uint32 provided by user to BgpV6EviVxlan
	SetPmsiLabel(value uint32) BgpV6EviVxlan
	// HasPmsiLabel checks if PmsiLabel has been set in BgpV6EviVxlan
	HasPmsiLabel() bool
	// AdLabel returns uint32, set in BgpV6EviVxlan.
	AdLabel() uint32
	// SetAdLabel assigns uint32 provided by user to BgpV6EviVxlan
	SetAdLabel(value uint32) BgpV6EviVxlan
	// HasAdLabel checks if AdLabel has been set in BgpV6EviVxlan
	HasAdLabel() bool
	// RouteDistinguisher returns BgpRouteDistinguisher, set in BgpV6EviVxlan.
	// BgpRouteDistinguisher is bGP Route Distinguisher.
	RouteDistinguisher() BgpRouteDistinguisher
	// SetRouteDistinguisher assigns BgpRouteDistinguisher provided by user to BgpV6EviVxlan.
	// BgpRouteDistinguisher is bGP Route Distinguisher.
	SetRouteDistinguisher(value BgpRouteDistinguisher) BgpV6EviVxlan
	// HasRouteDistinguisher checks if RouteDistinguisher has been set in BgpV6EviVxlan
	HasRouteDistinguisher() bool
	// RouteTargetExport returns BgpV6EviVxlanBgpRouteTargetIterIter, set in BgpV6EviVxlan
	RouteTargetExport() BgpV6EviVxlanBgpRouteTargetIter
	// RouteTargetImport returns BgpV6EviVxlanBgpRouteTargetIterIter, set in BgpV6EviVxlan
	RouteTargetImport() BgpV6EviVxlanBgpRouteTargetIter
	// L3RouteTargetExport returns BgpV6EviVxlanBgpRouteTargetIterIter, set in BgpV6EviVxlan
	L3RouteTargetExport() BgpV6EviVxlanBgpRouteTargetIter
	// L3RouteTargetImport returns BgpV6EviVxlanBgpRouteTargetIterIter, set in BgpV6EviVxlan
	L3RouteTargetImport() BgpV6EviVxlanBgpRouteTargetIter
	// Advanced returns BgpRouteAdvanced, set in BgpV6EviVxlan.
	// BgpRouteAdvanced is configuration for advanced BGP route range settings.
	Advanced() BgpRouteAdvanced
	// SetAdvanced assigns BgpRouteAdvanced provided by user to BgpV6EviVxlan.
	// BgpRouteAdvanced is configuration for advanced BGP route range settings.
	SetAdvanced(value BgpRouteAdvanced) BgpV6EviVxlan
	// HasAdvanced checks if Advanced has been set in BgpV6EviVxlan
	HasAdvanced() bool
	// Communities returns BgpV6EviVxlanBgpCommunityIterIter, set in BgpV6EviVxlan
	Communities() BgpV6EviVxlanBgpCommunityIter
	// ExtCommunities returns BgpV6EviVxlanBgpExtCommunityIterIter, set in BgpV6EviVxlan
	ExtCommunities() BgpV6EviVxlanBgpExtCommunityIter
	// AsPath returns BgpAsPath, set in BgpV6EviVxlan.
	// BgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed. This contains the configuration of how to include the Local AS in the AS path attribute of the MP REACH NLRI. It also contains optional configuration of additional AS Path Segments that can be included in the AS Path attribute. The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that a routing information passes through to reach the destination.
	AsPath() BgpAsPath
	// SetAsPath assigns BgpAsPath provided by user to BgpV6EviVxlan.
	// BgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed. This contains the configuration of how to include the Local AS in the AS path attribute of the MP REACH NLRI. It also contains optional configuration of additional AS Path Segments that can be included in the AS Path attribute. The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that a routing information passes through to reach the destination.
	SetAsPath(value BgpAsPath) BgpV6EviVxlan
	// HasAsPath checks if AsPath has been set in BgpV6EviVxlan
	HasAsPath() bool
	setNil()
}

// This contains the list of Broadcast Domains to be configured per EVI.
// BroadcastDomains returns a []BgpV6EviVxlanBroadcastDomain
func (obj *bgpV6EviVxlan) BroadcastDomains() BgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter {
	if len(obj.obj.BroadcastDomains) == 0 {
		obj.obj.BroadcastDomains = []*otg.BgpV6EviVxlanBroadcastDomain{}
	}
	if obj.broadcastDomainsHolder == nil {
		obj.broadcastDomainsHolder = newBgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter(&obj.obj.BroadcastDomains).setMsg(obj)
	}
	return obj.broadcastDomainsHolder
}

type bgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter struct {
	obj                               *bgpV6EviVxlan
	bgpV6EviVxlanBroadcastDomainSlice []BgpV6EviVxlanBroadcastDomain
	fieldPtr                          *[]*otg.BgpV6EviVxlanBroadcastDomain
}

func newBgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter(ptr *[]*otg.BgpV6EviVxlanBroadcastDomain) BgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter {
	return &bgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter{fieldPtr: ptr}
}

type BgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter interface {
	setMsg(*bgpV6EviVxlan) BgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter
	Items() []BgpV6EviVxlanBroadcastDomain
	Add() BgpV6EviVxlanBroadcastDomain
	Append(items ...BgpV6EviVxlanBroadcastDomain) BgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter
	Set(index int, newObj BgpV6EviVxlanBroadcastDomain) BgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter
	Clear() BgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter
	clearHolderSlice() BgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter
	appendHolderSlice(item BgpV6EviVxlanBroadcastDomain) BgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter
}

func (obj *bgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter) setMsg(msg *bgpV6EviVxlan) BgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpV6EviVxlanBroadcastDomain{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter) Items() []BgpV6EviVxlanBroadcastDomain {
	return obj.bgpV6EviVxlanBroadcastDomainSlice
}

func (obj *bgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter) Add() BgpV6EviVxlanBroadcastDomain {
	newObj := &otg.BgpV6EviVxlanBroadcastDomain{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpV6EviVxlanBroadcastDomain{obj: newObj}
	newLibObj.setDefault()
	obj.bgpV6EviVxlanBroadcastDomainSlice = append(obj.bgpV6EviVxlanBroadcastDomainSlice, newLibObj)
	return newLibObj
}

func (obj *bgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter) Append(items ...BgpV6EviVxlanBroadcastDomain) BgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpV6EviVxlanBroadcastDomainSlice = append(obj.bgpV6EviVxlanBroadcastDomainSlice, item)
	}
	return obj
}

func (obj *bgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter) Set(index int, newObj BgpV6EviVxlanBroadcastDomain) BgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpV6EviVxlanBroadcastDomainSlice[index] = newObj
	return obj
}
func (obj *bgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter) Clear() BgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpV6EviVxlanBroadcastDomain{}
		obj.bgpV6EviVxlanBroadcastDomainSlice = []BgpV6EviVxlanBroadcastDomain{}
	}
	return obj
}
func (obj *bgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter) clearHolderSlice() BgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter {
	if len(obj.bgpV6EviVxlanBroadcastDomainSlice) > 0 {
		obj.bgpV6EviVxlanBroadcastDomainSlice = []BgpV6EviVxlanBroadcastDomain{}
	}
	return obj
}
func (obj *bgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter) appendHolderSlice(item BgpV6EviVxlanBroadcastDomain) BgpV6EviVxlanBgpV6EviVxlanBroadcastDomainIter {
	obj.bgpV6EviVxlanBroadcastDomainSlice = append(obj.bgpV6EviVxlanBroadcastDomainSlice, item)
	return obj
}

type BgpV6EviVxlanReplicationTypeEnum string

// Enum of ReplicationType on BgpV6EviVxlan
var BgpV6EviVxlanReplicationType = struct {
	INGRESS_REPLICATION BgpV6EviVxlanReplicationTypeEnum
}{
	INGRESS_REPLICATION: BgpV6EviVxlanReplicationTypeEnum("ingress_replication"),
}

func (obj *bgpV6EviVxlan) ReplicationType() BgpV6EviVxlanReplicationTypeEnum {
	return BgpV6EviVxlanReplicationTypeEnum(obj.obj.ReplicationType.Enum().String())
}

// This model only supports Ingress Replication
// ReplicationType returns a string
func (obj *bgpV6EviVxlan) HasReplicationType() bool {
	return obj.obj.ReplicationType != nil
}

func (obj *bgpV6EviVxlan) SetReplicationType(value BgpV6EviVxlanReplicationTypeEnum) BgpV6EviVxlan {
	intValue, ok := otg.BgpV6EviVxlan_ReplicationType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpV6EviVxlanReplicationTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpV6EviVxlan_ReplicationType_Enum(intValue)
	obj.obj.ReplicationType = &enumValue

	return obj
}

// Downstream assigned VNI to be carried as Part of P-Multicast Service Interface Tunnel attribute (PMSI Tunnel Attribute) in Type 3 Inclusive Multicast Ethernet Tag Route.
// PmsiLabel returns a uint32
func (obj *bgpV6EviVxlan) PmsiLabel() uint32 {

	return *obj.obj.PmsiLabel

}

// Downstream assigned VNI to be carried as Part of P-Multicast Service Interface Tunnel attribute (PMSI Tunnel Attribute) in Type 3 Inclusive Multicast Ethernet Tag Route.
// PmsiLabel returns a uint32
func (obj *bgpV6EviVxlan) HasPmsiLabel() bool {
	return obj.obj.PmsiLabel != nil
}

// Downstream assigned VNI to be carried as Part of P-Multicast Service Interface Tunnel attribute (PMSI Tunnel Attribute) in Type 3 Inclusive Multicast Ethernet Tag Route.
// SetPmsiLabel sets the uint32 value in the BgpV6EviVxlan object
func (obj *bgpV6EviVxlan) SetPmsiLabel(value uint32) BgpV6EviVxlan {

	obj.obj.PmsiLabel = &value
	return obj
}

// The Auto-discovery Route label (AD label) value, which gets advertised in the Ethernet Auto-discovery Route per <EVI, ESI>
// AdLabel returns a uint32
func (obj *bgpV6EviVxlan) AdLabel() uint32 {

	return *obj.obj.AdLabel

}

// The Auto-discovery Route label (AD label) value, which gets advertised in the Ethernet Auto-discovery Route per <EVI, ESI>
// AdLabel returns a uint32
func (obj *bgpV6EviVxlan) HasAdLabel() bool {
	return obj.obj.AdLabel != nil
}

// The Auto-discovery Route label (AD label) value, which gets advertised in the Ethernet Auto-discovery Route per <EVI, ESI>
// SetAdLabel sets the uint32 value in the BgpV6EviVxlan object
func (obj *bgpV6EviVxlan) SetAdLabel(value uint32) BgpV6EviVxlan {

	obj.obj.AdLabel = &value
	return obj
}

// Colon separated Extended Community value of 6 Bytes - "AS number: Value" identifying an EVI.            Example - for the as_2octet "60005:100".
// RouteDistinguisher returns a BgpRouteDistinguisher
func (obj *bgpV6EviVxlan) RouteDistinguisher() BgpRouteDistinguisher {
	if obj.obj.RouteDistinguisher == nil {
		obj.obj.RouteDistinguisher = NewBgpRouteDistinguisher().msg()
	}
	if obj.routeDistinguisherHolder == nil {
		obj.routeDistinguisherHolder = &bgpRouteDistinguisher{obj: obj.obj.RouteDistinguisher}
	}
	return obj.routeDistinguisherHolder
}

// Colon separated Extended Community value of 6 Bytes - "AS number: Value" identifying an EVI.            Example - for the as_2octet "60005:100".
// RouteDistinguisher returns a BgpRouteDistinguisher
func (obj *bgpV6EviVxlan) HasRouteDistinguisher() bool {
	return obj.obj.RouteDistinguisher != nil
}

// Colon separated Extended Community value of 6 Bytes - "AS number: Value" identifying an EVI.            Example - for the as_2octet "60005:100".
// SetRouteDistinguisher sets the BgpRouteDistinguisher value in the BgpV6EviVxlan object
func (obj *bgpV6EviVxlan) SetRouteDistinguisher(value BgpRouteDistinguisher) BgpV6EviVxlan {

	obj.routeDistinguisherHolder = nil
	obj.obj.RouteDistinguisher = value.msg()

	return obj
}

// List of Layer 2 Virtual Network Identifier (L2VNI) export targets associated with this EVI.
// RouteTargetExport returns a []BgpRouteTarget
func (obj *bgpV6EviVxlan) RouteTargetExport() BgpV6EviVxlanBgpRouteTargetIter {
	if len(obj.obj.RouteTargetExport) == 0 {
		obj.obj.RouteTargetExport = []*otg.BgpRouteTarget{}
	}
	if obj.routeTargetExportHolder == nil {
		obj.routeTargetExportHolder = newBgpV6EviVxlanBgpRouteTargetIter(&obj.obj.RouteTargetExport).setMsg(obj)
	}
	return obj.routeTargetExportHolder
}

type bgpV6EviVxlanBgpRouteTargetIter struct {
	obj                 *bgpV6EviVxlan
	bgpRouteTargetSlice []BgpRouteTarget
	fieldPtr            *[]*otg.BgpRouteTarget
}

func newBgpV6EviVxlanBgpRouteTargetIter(ptr *[]*otg.BgpRouteTarget) BgpV6EviVxlanBgpRouteTargetIter {
	return &bgpV6EviVxlanBgpRouteTargetIter{fieldPtr: ptr}
}

type BgpV6EviVxlanBgpRouteTargetIter interface {
	setMsg(*bgpV6EviVxlan) BgpV6EviVxlanBgpRouteTargetIter
	Items() []BgpRouteTarget
	Add() BgpRouteTarget
	Append(items ...BgpRouteTarget) BgpV6EviVxlanBgpRouteTargetIter
	Set(index int, newObj BgpRouteTarget) BgpV6EviVxlanBgpRouteTargetIter
	Clear() BgpV6EviVxlanBgpRouteTargetIter
	clearHolderSlice() BgpV6EviVxlanBgpRouteTargetIter
	appendHolderSlice(item BgpRouteTarget) BgpV6EviVxlanBgpRouteTargetIter
}

func (obj *bgpV6EviVxlanBgpRouteTargetIter) setMsg(msg *bgpV6EviVxlan) BgpV6EviVxlanBgpRouteTargetIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpRouteTarget{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV6EviVxlanBgpRouteTargetIter) Items() []BgpRouteTarget {
	return obj.bgpRouteTargetSlice
}

func (obj *bgpV6EviVxlanBgpRouteTargetIter) Add() BgpRouteTarget {
	newObj := &otg.BgpRouteTarget{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpRouteTarget{obj: newObj}
	newLibObj.setDefault()
	obj.bgpRouteTargetSlice = append(obj.bgpRouteTargetSlice, newLibObj)
	return newLibObj
}

func (obj *bgpV6EviVxlanBgpRouteTargetIter) Append(items ...BgpRouteTarget) BgpV6EviVxlanBgpRouteTargetIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpRouteTargetSlice = append(obj.bgpRouteTargetSlice, item)
	}
	return obj
}

func (obj *bgpV6EviVxlanBgpRouteTargetIter) Set(index int, newObj BgpRouteTarget) BgpV6EviVxlanBgpRouteTargetIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpRouteTargetSlice[index] = newObj
	return obj
}
func (obj *bgpV6EviVxlanBgpRouteTargetIter) Clear() BgpV6EviVxlanBgpRouteTargetIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpRouteTarget{}
		obj.bgpRouteTargetSlice = []BgpRouteTarget{}
	}
	return obj
}
func (obj *bgpV6EviVxlanBgpRouteTargetIter) clearHolderSlice() BgpV6EviVxlanBgpRouteTargetIter {
	if len(obj.bgpRouteTargetSlice) > 0 {
		obj.bgpRouteTargetSlice = []BgpRouteTarget{}
	}
	return obj
}
func (obj *bgpV6EviVxlanBgpRouteTargetIter) appendHolderSlice(item BgpRouteTarget) BgpV6EviVxlanBgpRouteTargetIter {
	obj.bgpRouteTargetSlice = append(obj.bgpRouteTargetSlice, item)
	return obj
}

// List of L2VNI import targets associated with this EVI.
// RouteTargetImport returns a []BgpRouteTarget
func (obj *bgpV6EviVxlan) RouteTargetImport() BgpV6EviVxlanBgpRouteTargetIter {
	if len(obj.obj.RouteTargetImport) == 0 {
		obj.obj.RouteTargetImport = []*otg.BgpRouteTarget{}
	}
	if obj.routeTargetImportHolder == nil {
		obj.routeTargetImportHolder = newBgpV6EviVxlanBgpRouteTargetIter(&obj.obj.RouteTargetImport).setMsg(obj)
	}
	return obj.routeTargetImportHolder
}

// List of Layer 3 Virtual Network Identifier (L3VNI) Export Route Targets.
// L3RouteTargetExport returns a []BgpRouteTarget
func (obj *bgpV6EviVxlan) L3RouteTargetExport() BgpV6EviVxlanBgpRouteTargetIter {
	if len(obj.obj.L3RouteTargetExport) == 0 {
		obj.obj.L3RouteTargetExport = []*otg.BgpRouteTarget{}
	}
	if obj.l3RouteTargetExportHolder == nil {
		obj.l3RouteTargetExportHolder = newBgpV6EviVxlanBgpRouteTargetIter(&obj.obj.L3RouteTargetExport).setMsg(obj)
	}
	return obj.l3RouteTargetExportHolder
}

// List of L3VNI Import Route Targets.
// L3RouteTargetImport returns a []BgpRouteTarget
func (obj *bgpV6EviVxlan) L3RouteTargetImport() BgpV6EviVxlanBgpRouteTargetIter {
	if len(obj.obj.L3RouteTargetImport) == 0 {
		obj.obj.L3RouteTargetImport = []*otg.BgpRouteTarget{}
	}
	if obj.l3RouteTargetImportHolder == nil {
		obj.l3RouteTargetImportHolder = newBgpV6EviVxlanBgpRouteTargetIter(&obj.obj.L3RouteTargetImport).setMsg(obj)
	}
	return obj.l3RouteTargetImportHolder
}

// description is TBD
// Advanced returns a BgpRouteAdvanced
func (obj *bgpV6EviVxlan) Advanced() BgpRouteAdvanced {
	if obj.obj.Advanced == nil {
		obj.obj.Advanced = NewBgpRouteAdvanced().msg()
	}
	if obj.advancedHolder == nil {
		obj.advancedHolder = &bgpRouteAdvanced{obj: obj.obj.Advanced}
	}
	return obj.advancedHolder
}

// description is TBD
// Advanced returns a BgpRouteAdvanced
func (obj *bgpV6EviVxlan) HasAdvanced() bool {
	return obj.obj.Advanced != nil
}

// description is TBD
// SetAdvanced sets the BgpRouteAdvanced value in the BgpV6EviVxlan object
func (obj *bgpV6EviVxlan) SetAdvanced(value BgpRouteAdvanced) BgpV6EviVxlan {

	obj.advancedHolder = nil
	obj.obj.Advanced = value.msg()

	return obj
}

// Optional community settings.
// Communities returns a []BgpCommunity
func (obj *bgpV6EviVxlan) Communities() BgpV6EviVxlanBgpCommunityIter {
	if len(obj.obj.Communities) == 0 {
		obj.obj.Communities = []*otg.BgpCommunity{}
	}
	if obj.communitiesHolder == nil {
		obj.communitiesHolder = newBgpV6EviVxlanBgpCommunityIter(&obj.obj.Communities).setMsg(obj)
	}
	return obj.communitiesHolder
}

type bgpV6EviVxlanBgpCommunityIter struct {
	obj               *bgpV6EviVxlan
	bgpCommunitySlice []BgpCommunity
	fieldPtr          *[]*otg.BgpCommunity
}

func newBgpV6EviVxlanBgpCommunityIter(ptr *[]*otg.BgpCommunity) BgpV6EviVxlanBgpCommunityIter {
	return &bgpV6EviVxlanBgpCommunityIter{fieldPtr: ptr}
}

type BgpV6EviVxlanBgpCommunityIter interface {
	setMsg(*bgpV6EviVxlan) BgpV6EviVxlanBgpCommunityIter
	Items() []BgpCommunity
	Add() BgpCommunity
	Append(items ...BgpCommunity) BgpV6EviVxlanBgpCommunityIter
	Set(index int, newObj BgpCommunity) BgpV6EviVxlanBgpCommunityIter
	Clear() BgpV6EviVxlanBgpCommunityIter
	clearHolderSlice() BgpV6EviVxlanBgpCommunityIter
	appendHolderSlice(item BgpCommunity) BgpV6EviVxlanBgpCommunityIter
}

func (obj *bgpV6EviVxlanBgpCommunityIter) setMsg(msg *bgpV6EviVxlan) BgpV6EviVxlanBgpCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV6EviVxlanBgpCommunityIter) Items() []BgpCommunity {
	return obj.bgpCommunitySlice
}

func (obj *bgpV6EviVxlanBgpCommunityIter) Add() BgpCommunity {
	newObj := &otg.BgpCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpV6EviVxlanBgpCommunityIter) Append(items ...BgpCommunity) BgpV6EviVxlanBgpCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, item)
	}
	return obj
}

func (obj *bgpV6EviVxlanBgpCommunityIter) Set(index int, newObj BgpCommunity) BgpV6EviVxlanBgpCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpV6EviVxlanBgpCommunityIter) Clear() BgpV6EviVxlanBgpCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpCommunity{}
		obj.bgpCommunitySlice = []BgpCommunity{}
	}
	return obj
}
func (obj *bgpV6EviVxlanBgpCommunityIter) clearHolderSlice() BgpV6EviVxlanBgpCommunityIter {
	if len(obj.bgpCommunitySlice) > 0 {
		obj.bgpCommunitySlice = []BgpCommunity{}
	}
	return obj
}
func (obj *bgpV6EviVxlanBgpCommunityIter) appendHolderSlice(item BgpCommunity) BgpV6EviVxlanBgpCommunityIter {
	obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, item)
	return obj
}

// Optional Extended Community settings. The Extended Communities Attribute is a transitive optional BGP attribute, with the Type Code 16. Community and Extended Communities  attributes are utilized to trigger routing decisions, such as acceptance, rejection,  preference, or redistribution. An extended community is an 8-Bytes value. It is divided into two main parts. The first 2 Bytes of the community encode a type and sub-type fields and the last 6 Bytes carry a unique set of data in a format defined by the type and sub-type field. Extended communities provide a larger  range for grouping or categorizing communities. When type is administrator_as_2octet or administrator_as_4octet, the valid sub types are route target and origin. The valid value for  administrator_as_2octet and administrator_as_4octet type is either two byte AS followed by four byte local administrator id or four byte AS followed by two  byte local administrator id.  When type is administrator_ipv4_address the valid sub types are route target and origin. The valid value for  administrator_ipv4_address is a four byte IPv4 address followed by a two byte local administrator id.  When type is opaque, valid sub types are color and encapsulation. When sub type is color, first two bytes of the value field contain flags and last four bytes  contains the value of the color. When sub type is encapsulation the first four bytes of value field are reserved and last two bytes carries the tunnel type from  IANA's "ETHER TYPES" registry e.g IPv4 (protocol type = 0x0800), IPv6 (protocol type = 0x86dd), and MPLS (protocol type = 0x8847). When type is administrator_as_2octet_link_bandwidth the valid sub type is extended_bandwidth. The first two bytes of the value field contains the AS number and the last four bytes contains the bandwidth in IEEE floating point format.  When type is evpn the valid subtype is mac_address. In the value field the low-order bit of the first byte(Flags) is defined as the "Sticky/static" flag and may be set to 1, indicating the MAC address is static and cannot move. The second byte is reserved and the  last four bytes contain the sequence number which is used to ensure that PEs retain the correct MAC/IP Advertisement route when multiple updates  occur for the same MAC address.
// ExtCommunities returns a []BgpExtCommunity
func (obj *bgpV6EviVxlan) ExtCommunities() BgpV6EviVxlanBgpExtCommunityIter {
	if len(obj.obj.ExtCommunities) == 0 {
		obj.obj.ExtCommunities = []*otg.BgpExtCommunity{}
	}
	if obj.extCommunitiesHolder == nil {
		obj.extCommunitiesHolder = newBgpV6EviVxlanBgpExtCommunityIter(&obj.obj.ExtCommunities).setMsg(obj)
	}
	return obj.extCommunitiesHolder
}

type bgpV6EviVxlanBgpExtCommunityIter struct {
	obj                  *bgpV6EviVxlan
	bgpExtCommunitySlice []BgpExtCommunity
	fieldPtr             *[]*otg.BgpExtCommunity
}

func newBgpV6EviVxlanBgpExtCommunityIter(ptr *[]*otg.BgpExtCommunity) BgpV6EviVxlanBgpExtCommunityIter {
	return &bgpV6EviVxlanBgpExtCommunityIter{fieldPtr: ptr}
}

type BgpV6EviVxlanBgpExtCommunityIter interface {
	setMsg(*bgpV6EviVxlan) BgpV6EviVxlanBgpExtCommunityIter
	Items() []BgpExtCommunity
	Add() BgpExtCommunity
	Append(items ...BgpExtCommunity) BgpV6EviVxlanBgpExtCommunityIter
	Set(index int, newObj BgpExtCommunity) BgpV6EviVxlanBgpExtCommunityIter
	Clear() BgpV6EviVxlanBgpExtCommunityIter
	clearHolderSlice() BgpV6EviVxlanBgpExtCommunityIter
	appendHolderSlice(item BgpExtCommunity) BgpV6EviVxlanBgpExtCommunityIter
}

func (obj *bgpV6EviVxlanBgpExtCommunityIter) setMsg(msg *bgpV6EviVxlan) BgpV6EviVxlanBgpExtCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpExtCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV6EviVxlanBgpExtCommunityIter) Items() []BgpExtCommunity {
	return obj.bgpExtCommunitySlice
}

func (obj *bgpV6EviVxlanBgpExtCommunityIter) Add() BgpExtCommunity {
	newObj := &otg.BgpExtCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpExtCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.bgpExtCommunitySlice = append(obj.bgpExtCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpV6EviVxlanBgpExtCommunityIter) Append(items ...BgpExtCommunity) BgpV6EviVxlanBgpExtCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpExtCommunitySlice = append(obj.bgpExtCommunitySlice, item)
	}
	return obj
}

func (obj *bgpV6EviVxlanBgpExtCommunityIter) Set(index int, newObj BgpExtCommunity) BgpV6EviVxlanBgpExtCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpExtCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpV6EviVxlanBgpExtCommunityIter) Clear() BgpV6EviVxlanBgpExtCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpExtCommunity{}
		obj.bgpExtCommunitySlice = []BgpExtCommunity{}
	}
	return obj
}
func (obj *bgpV6EviVxlanBgpExtCommunityIter) clearHolderSlice() BgpV6EviVxlanBgpExtCommunityIter {
	if len(obj.bgpExtCommunitySlice) > 0 {
		obj.bgpExtCommunitySlice = []BgpExtCommunity{}
	}
	return obj
}
func (obj *bgpV6EviVxlanBgpExtCommunityIter) appendHolderSlice(item BgpExtCommunity) BgpV6EviVxlanBgpExtCommunityIter {
	obj.bgpExtCommunitySlice = append(obj.bgpExtCommunitySlice, item)
	return obj
}

// Optional AS PATH settings.
// AsPath returns a BgpAsPath
func (obj *bgpV6EviVxlan) AsPath() BgpAsPath {
	if obj.obj.AsPath == nil {
		obj.obj.AsPath = NewBgpAsPath().msg()
	}
	if obj.asPathHolder == nil {
		obj.asPathHolder = &bgpAsPath{obj: obj.obj.AsPath}
	}
	return obj.asPathHolder
}

// Optional AS PATH settings.
// AsPath returns a BgpAsPath
func (obj *bgpV6EviVxlan) HasAsPath() bool {
	return obj.obj.AsPath != nil
}

// Optional AS PATH settings.
// SetAsPath sets the BgpAsPath value in the BgpV6EviVxlan object
func (obj *bgpV6EviVxlan) SetAsPath(value BgpAsPath) BgpV6EviVxlan {

	obj.asPathHolder = nil
	obj.obj.AsPath = value.msg()

	return obj
}

func (obj *bgpV6EviVxlan) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.BroadcastDomains) != 0 {

		if set_default {
			obj.BroadcastDomains().clearHolderSlice()
			for _, item := range obj.obj.BroadcastDomains {
				obj.BroadcastDomains().appendHolderSlice(&bgpV6EviVxlanBroadcastDomain{obj: item})
			}
		}
		for _, item := range obj.BroadcastDomains().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.PmsiLabel != nil {

		if *obj.obj.PmsiLabel > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpV6EviVxlan.PmsiLabel <= 16777215 but Got %d", *obj.obj.PmsiLabel))
		}

	}

	if obj.obj.AdLabel != nil {

		if *obj.obj.AdLabel > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpV6EviVxlan.AdLabel <= 16777215 but Got %d", *obj.obj.AdLabel))
		}

	}

	if obj.obj.RouteDistinguisher != nil {

		obj.RouteDistinguisher().validateObj(vObj, set_default)
	}

	if len(obj.obj.RouteTargetExport) != 0 {

		if set_default {
			obj.RouteTargetExport().clearHolderSlice()
			for _, item := range obj.obj.RouteTargetExport {
				obj.RouteTargetExport().appendHolderSlice(&bgpRouteTarget{obj: item})
			}
		}
		for _, item := range obj.RouteTargetExport().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.RouteTargetImport) != 0 {

		if set_default {
			obj.RouteTargetImport().clearHolderSlice()
			for _, item := range obj.obj.RouteTargetImport {
				obj.RouteTargetImport().appendHolderSlice(&bgpRouteTarget{obj: item})
			}
		}
		for _, item := range obj.RouteTargetImport().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.L3RouteTargetExport) != 0 {

		if set_default {
			obj.L3RouteTargetExport().clearHolderSlice()
			for _, item := range obj.obj.L3RouteTargetExport {
				obj.L3RouteTargetExport().appendHolderSlice(&bgpRouteTarget{obj: item})
			}
		}
		for _, item := range obj.L3RouteTargetExport().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.L3RouteTargetImport) != 0 {

		if set_default {
			obj.L3RouteTargetImport().clearHolderSlice()
			for _, item := range obj.obj.L3RouteTargetImport {
				obj.L3RouteTargetImport().appendHolderSlice(&bgpRouteTarget{obj: item})
			}
		}
		for _, item := range obj.L3RouteTargetImport().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.Advanced != nil {

		obj.Advanced().validateObj(vObj, set_default)
	}

	if len(obj.obj.Communities) != 0 {

		if set_default {
			obj.Communities().clearHolderSlice()
			for _, item := range obj.obj.Communities {
				obj.Communities().appendHolderSlice(&bgpCommunity{obj: item})
			}
		}
		for _, item := range obj.Communities().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.ExtCommunities) != 0 {

		if set_default {
			obj.ExtCommunities().clearHolderSlice()
			for _, item := range obj.obj.ExtCommunities {
				obj.ExtCommunities().appendHolderSlice(&bgpExtCommunity{obj: item})
			}
		}
		for _, item := range obj.ExtCommunities().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.AsPath != nil {

		obj.AsPath().validateObj(vObj, set_default)
	}

}

func (obj *bgpV6EviVxlan) setDefault() {
	if obj.obj.ReplicationType == nil {
		obj.SetReplicationType(BgpV6EviVxlanReplicationType.INGRESS_REPLICATION)

	}
	if obj.obj.PmsiLabel == nil {
		obj.SetPmsiLabel(16)
	}
	if obj.obj.AdLabel == nil {
		obj.SetAdLabel(0)
	}

}
