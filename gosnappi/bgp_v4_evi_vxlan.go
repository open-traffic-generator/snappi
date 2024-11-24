package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpV4EviVxlan *****
type bgpV4EviVxlan struct {
	validation
	obj                       *otg.BgpV4EviVxlan
	marshaller                marshalBgpV4EviVxlan
	unMarshaller              unMarshalBgpV4EviVxlan
	broadcastDomainsHolder    BgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter
	routeDistinguisherHolder  BgpRouteDistinguisher
	routeTargetExportHolder   BgpV4EviVxlanBgpRouteTargetIter
	routeTargetImportHolder   BgpV4EviVxlanBgpRouteTargetIter
	l3RouteTargetExportHolder BgpV4EviVxlanBgpRouteTargetIter
	l3RouteTargetImportHolder BgpV4EviVxlanBgpRouteTargetIter
	advancedHolder            BgpRouteAdvanced
	communitiesHolder         BgpV4EviVxlanBgpCommunityIter
	extCommunitiesHolder      BgpV4EviVxlanBgpExtCommunityIter
	asPathHolder              BgpAsPath
}

func NewBgpV4EviVxlan() BgpV4EviVxlan {
	obj := bgpV4EviVxlan{obj: &otg.BgpV4EviVxlan{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpV4EviVxlan) msg() *otg.BgpV4EviVxlan {
	return obj.obj
}

func (obj *bgpV4EviVxlan) setMsg(msg *otg.BgpV4EviVxlan) BgpV4EviVxlan {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpV4EviVxlan struct {
	obj *bgpV4EviVxlan
}

type marshalBgpV4EviVxlan interface {
	// ToProto marshals BgpV4EviVxlan to protobuf object *otg.BgpV4EviVxlan
	ToProto() (*otg.BgpV4EviVxlan, error)
	// ToPbText marshals BgpV4EviVxlan to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpV4EviVxlan to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpV4EviVxlan to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpV4EviVxlan to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpV4EviVxlan struct {
	obj *bgpV4EviVxlan
}

type unMarshalBgpV4EviVxlan interface {
	// FromProto unmarshals BgpV4EviVxlan from protobuf object *otg.BgpV4EviVxlan
	FromProto(msg *otg.BgpV4EviVxlan) (BgpV4EviVxlan, error)
	// FromPbText unmarshals BgpV4EviVxlan from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpV4EviVxlan from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpV4EviVxlan from JSON text
	FromJson(value string) error
}

func (obj *bgpV4EviVxlan) Marshal() marshalBgpV4EviVxlan {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpV4EviVxlan{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpV4EviVxlan) Unmarshal() unMarshalBgpV4EviVxlan {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpV4EviVxlan{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpV4EviVxlan) ToProto() (*otg.BgpV4EviVxlan, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpV4EviVxlan) FromProto(msg *otg.BgpV4EviVxlan) (BgpV4EviVxlan, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpV4EviVxlan) ToPbText() (string, error) {
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

func (m *unMarshalbgpV4EviVxlan) FromPbText(value string) error {
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

func (m *marshalbgpV4EviVxlan) ToYaml() (string, error) {
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

func (m *unMarshalbgpV4EviVxlan) FromYaml(value string) error {
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

func (m *marshalbgpV4EviVxlan) ToJsonRaw() (string, error) {
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

func (m *marshalbgpV4EviVxlan) ToJson() (string, error) {
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

func (m *unMarshalbgpV4EviVxlan) FromJson(value string) error {
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

func (obj *bgpV4EviVxlan) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpV4EviVxlan) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpV4EviVxlan) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpV4EviVxlan) Clone() (BgpV4EviVxlan, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpV4EviVxlan()
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

func (obj *bgpV4EviVxlan) setNil() {
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

// BgpV4EviVxlan is configuration for BGP EVPN EVI. Advertises following routes -
//
// # Type 3 - Inclusive Multicast Ethernet Tag Route
//
// Type 1 -  Ethernet Auto-discovery Route (Per EVI)
//
// Type 1 -  Ethernet Auto-discovery Route (Per ES)
type BgpV4EviVxlan interface {
	Validation
	// msg marshals BgpV4EviVxlan to protobuf object *otg.BgpV4EviVxlan
	// and doesn't set defaults
	msg() *otg.BgpV4EviVxlan
	// setMsg unmarshals BgpV4EviVxlan from protobuf object *otg.BgpV4EviVxlan
	// and doesn't set defaults
	setMsg(*otg.BgpV4EviVxlan) BgpV4EviVxlan
	// provides marshal interface
	Marshal() marshalBgpV4EviVxlan
	// provides unmarshal interface
	Unmarshal() unMarshalBgpV4EviVxlan
	// validate validates BgpV4EviVxlan
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpV4EviVxlan, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// BroadcastDomains returns BgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIterIter, set in BgpV4EviVxlan
	BroadcastDomains() BgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter
	// ReplicationType returns BgpV4EviVxlanReplicationTypeEnum, set in BgpV4EviVxlan
	ReplicationType() BgpV4EviVxlanReplicationTypeEnum
	// SetReplicationType assigns BgpV4EviVxlanReplicationTypeEnum provided by user to BgpV4EviVxlan
	SetReplicationType(value BgpV4EviVxlanReplicationTypeEnum) BgpV4EviVxlan
	// HasReplicationType checks if ReplicationType has been set in BgpV4EviVxlan
	HasReplicationType() bool
	// PmsiLabel returns uint32, set in BgpV4EviVxlan.
	PmsiLabel() uint32
	// SetPmsiLabel assigns uint32 provided by user to BgpV4EviVxlan
	SetPmsiLabel(value uint32) BgpV4EviVxlan
	// HasPmsiLabel checks if PmsiLabel has been set in BgpV4EviVxlan
	HasPmsiLabel() bool
	// AdLabel returns uint32, set in BgpV4EviVxlan.
	AdLabel() uint32
	// SetAdLabel assigns uint32 provided by user to BgpV4EviVxlan
	SetAdLabel(value uint32) BgpV4EviVxlan
	// HasAdLabel checks if AdLabel has been set in BgpV4EviVxlan
	HasAdLabel() bool
	// RouteDistinguisher returns BgpRouteDistinguisher, set in BgpV4EviVxlan.
	// BgpRouteDistinguisher is bGP Route Distinguisher.
	RouteDistinguisher() BgpRouteDistinguisher
	// SetRouteDistinguisher assigns BgpRouteDistinguisher provided by user to BgpV4EviVxlan.
	// BgpRouteDistinguisher is bGP Route Distinguisher.
	SetRouteDistinguisher(value BgpRouteDistinguisher) BgpV4EviVxlan
	// HasRouteDistinguisher checks if RouteDistinguisher has been set in BgpV4EviVxlan
	HasRouteDistinguisher() bool
	// RouteTargetExport returns BgpV4EviVxlanBgpRouteTargetIterIter, set in BgpV4EviVxlan
	RouteTargetExport() BgpV4EviVxlanBgpRouteTargetIter
	// RouteTargetImport returns BgpV4EviVxlanBgpRouteTargetIterIter, set in BgpV4EviVxlan
	RouteTargetImport() BgpV4EviVxlanBgpRouteTargetIter
	// L3RouteTargetExport returns BgpV4EviVxlanBgpRouteTargetIterIter, set in BgpV4EviVxlan
	L3RouteTargetExport() BgpV4EviVxlanBgpRouteTargetIter
	// L3RouteTargetImport returns BgpV4EviVxlanBgpRouteTargetIterIter, set in BgpV4EviVxlan
	L3RouteTargetImport() BgpV4EviVxlanBgpRouteTargetIter
	// Advanced returns BgpRouteAdvanced, set in BgpV4EviVxlan.
	// BgpRouteAdvanced is configuration for advanced BGP route range settings.
	Advanced() BgpRouteAdvanced
	// SetAdvanced assigns BgpRouteAdvanced provided by user to BgpV4EviVxlan.
	// BgpRouteAdvanced is configuration for advanced BGP route range settings.
	SetAdvanced(value BgpRouteAdvanced) BgpV4EviVxlan
	// HasAdvanced checks if Advanced has been set in BgpV4EviVxlan
	HasAdvanced() bool
	// Communities returns BgpV4EviVxlanBgpCommunityIterIter, set in BgpV4EviVxlan
	Communities() BgpV4EviVxlanBgpCommunityIter
	// ExtCommunities returns BgpV4EviVxlanBgpExtCommunityIterIter, set in BgpV4EviVxlan
	ExtCommunities() BgpV4EviVxlanBgpExtCommunityIter
	// AsPath returns BgpAsPath, set in BgpV4EviVxlan.
	// BgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed. This contains the configuration of how to include the Local AS in the AS path attribute of the MP REACH NLRI. It also contains optional configuration of additional AS Path Segments that can be included in the AS Path attribute. The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that a routing information passes through to reach the destination.
	AsPath() BgpAsPath
	// SetAsPath assigns BgpAsPath provided by user to BgpV4EviVxlan.
	// BgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed. This contains the configuration of how to include the Local AS in the AS path attribute of the MP REACH NLRI. It also contains optional configuration of additional AS Path Segments that can be included in the AS Path attribute. The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that a routing information passes through to reach the destination.
	SetAsPath(value BgpAsPath) BgpV4EviVxlan
	// HasAsPath checks if AsPath has been set in BgpV4EviVxlan
	HasAsPath() bool
	setNil()
}

// This contains the list of Broadcast Domains to be configured per EVI.
// BroadcastDomains returns a []BgpV4EviVxlanBroadcastDomain
func (obj *bgpV4EviVxlan) BroadcastDomains() BgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter {
	if len(obj.obj.BroadcastDomains) == 0 {
		obj.obj.BroadcastDomains = []*otg.BgpV4EviVxlanBroadcastDomain{}
	}
	if obj.broadcastDomainsHolder == nil {
		obj.broadcastDomainsHolder = newBgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter(&obj.obj.BroadcastDomains).setMsg(obj)
	}
	return obj.broadcastDomainsHolder
}

type bgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter struct {
	obj                               *bgpV4EviVxlan
	bgpV4EviVxlanBroadcastDomainSlice []BgpV4EviVxlanBroadcastDomain
	fieldPtr                          *[]*otg.BgpV4EviVxlanBroadcastDomain
}

func newBgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter(ptr *[]*otg.BgpV4EviVxlanBroadcastDomain) BgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter {
	return &bgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter{fieldPtr: ptr}
}

type BgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter interface {
	setMsg(*bgpV4EviVxlan) BgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter
	Items() []BgpV4EviVxlanBroadcastDomain
	Add() BgpV4EviVxlanBroadcastDomain
	Append(items ...BgpV4EviVxlanBroadcastDomain) BgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter
	Set(index int, newObj BgpV4EviVxlanBroadcastDomain) BgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter
	Clear() BgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter
	clearHolderSlice() BgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter
	appendHolderSlice(item BgpV4EviVxlanBroadcastDomain) BgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter
}

func (obj *bgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter) setMsg(msg *bgpV4EviVxlan) BgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpV4EviVxlanBroadcastDomain{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter) Items() []BgpV4EviVxlanBroadcastDomain {
	return obj.bgpV4EviVxlanBroadcastDomainSlice
}

func (obj *bgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter) Add() BgpV4EviVxlanBroadcastDomain {
	newObj := &otg.BgpV4EviVxlanBroadcastDomain{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpV4EviVxlanBroadcastDomain{obj: newObj}
	newLibObj.setDefault()
	obj.bgpV4EviVxlanBroadcastDomainSlice = append(obj.bgpV4EviVxlanBroadcastDomainSlice, newLibObj)
	return newLibObj
}

func (obj *bgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter) Append(items ...BgpV4EviVxlanBroadcastDomain) BgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpV4EviVxlanBroadcastDomainSlice = append(obj.bgpV4EviVxlanBroadcastDomainSlice, item)
	}
	return obj
}

func (obj *bgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter) Set(index int, newObj BgpV4EviVxlanBroadcastDomain) BgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpV4EviVxlanBroadcastDomainSlice[index] = newObj
	return obj
}
func (obj *bgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter) Clear() BgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpV4EviVxlanBroadcastDomain{}
		obj.bgpV4EviVxlanBroadcastDomainSlice = []BgpV4EviVxlanBroadcastDomain{}
	}
	return obj
}
func (obj *bgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter) clearHolderSlice() BgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter {
	if len(obj.bgpV4EviVxlanBroadcastDomainSlice) > 0 {
		obj.bgpV4EviVxlanBroadcastDomainSlice = []BgpV4EviVxlanBroadcastDomain{}
	}
	return obj
}
func (obj *bgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter) appendHolderSlice(item BgpV4EviVxlanBroadcastDomain) BgpV4EviVxlanBgpV4EviVxlanBroadcastDomainIter {
	obj.bgpV4EviVxlanBroadcastDomainSlice = append(obj.bgpV4EviVxlanBroadcastDomainSlice, item)
	return obj
}

type BgpV4EviVxlanReplicationTypeEnum string

// Enum of ReplicationType on BgpV4EviVxlan
var BgpV4EviVxlanReplicationType = struct {
	INGRESS_REPLICATION BgpV4EviVxlanReplicationTypeEnum
}{
	INGRESS_REPLICATION: BgpV4EviVxlanReplicationTypeEnum("ingress_replication"),
}

func (obj *bgpV4EviVxlan) ReplicationType() BgpV4EviVxlanReplicationTypeEnum {
	return BgpV4EviVxlanReplicationTypeEnum(obj.obj.ReplicationType.Enum().String())
}

// This model only supports Ingress Replication
// ReplicationType returns a string
func (obj *bgpV4EviVxlan) HasReplicationType() bool {
	return obj.obj.ReplicationType != nil
}

func (obj *bgpV4EviVxlan) SetReplicationType(value BgpV4EviVxlanReplicationTypeEnum) BgpV4EviVxlan {
	intValue, ok := otg.BgpV4EviVxlan_ReplicationType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpV4EviVxlanReplicationTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpV4EviVxlan_ReplicationType_Enum(intValue)
	obj.obj.ReplicationType = &enumValue

	return obj
}

// Downstream assigned VNI to be carried as Part of P-Multicast Service Interface Tunnel attribute (PMSI Tunnel Attribute) in Type 3 Inclusive Multicast Ethernet Tag Route.
// PmsiLabel returns a uint32
func (obj *bgpV4EviVxlan) PmsiLabel() uint32 {

	return *obj.obj.PmsiLabel

}

// Downstream assigned VNI to be carried as Part of P-Multicast Service Interface Tunnel attribute (PMSI Tunnel Attribute) in Type 3 Inclusive Multicast Ethernet Tag Route.
// PmsiLabel returns a uint32
func (obj *bgpV4EviVxlan) HasPmsiLabel() bool {
	return obj.obj.PmsiLabel != nil
}

// Downstream assigned VNI to be carried as Part of P-Multicast Service Interface Tunnel attribute (PMSI Tunnel Attribute) in Type 3 Inclusive Multicast Ethernet Tag Route.
// SetPmsiLabel sets the uint32 value in the BgpV4EviVxlan object
func (obj *bgpV4EviVxlan) SetPmsiLabel(value uint32) BgpV4EviVxlan {

	obj.obj.PmsiLabel = &value
	return obj
}

// The Auto-discovery Route label (AD label) value, which gets advertised in the Ethernet Auto-discovery Route per <EVI, ESI>
// AdLabel returns a uint32
func (obj *bgpV4EviVxlan) AdLabel() uint32 {

	return *obj.obj.AdLabel

}

// The Auto-discovery Route label (AD label) value, which gets advertised in the Ethernet Auto-discovery Route per <EVI, ESI>
// AdLabel returns a uint32
func (obj *bgpV4EviVxlan) HasAdLabel() bool {
	return obj.obj.AdLabel != nil
}

// The Auto-discovery Route label (AD label) value, which gets advertised in the Ethernet Auto-discovery Route per <EVI, ESI>
// SetAdLabel sets the uint32 value in the BgpV4EviVxlan object
func (obj *bgpV4EviVxlan) SetAdLabel(value uint32) BgpV4EviVxlan {

	obj.obj.AdLabel = &value
	return obj
}

// Colon separated Extended Community value of 6 Bytes - "AS number: Value" identifying an EVI.            Example - for the as_2octet "60005:100".
// RouteDistinguisher returns a BgpRouteDistinguisher
func (obj *bgpV4EviVxlan) RouteDistinguisher() BgpRouteDistinguisher {
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
func (obj *bgpV4EviVxlan) HasRouteDistinguisher() bool {
	return obj.obj.RouteDistinguisher != nil
}

// Colon separated Extended Community value of 6 Bytes - "AS number: Value" identifying an EVI.            Example - for the as_2octet "60005:100".
// SetRouteDistinguisher sets the BgpRouteDistinguisher value in the BgpV4EviVxlan object
func (obj *bgpV4EviVxlan) SetRouteDistinguisher(value BgpRouteDistinguisher) BgpV4EviVxlan {

	obj.routeDistinguisherHolder = nil
	obj.obj.RouteDistinguisher = value.msg()

	return obj
}

// List of Layer 2 Virtual Network Identifier (L2VNI) export targets associated with this EVI.
// RouteTargetExport returns a []BgpRouteTarget
func (obj *bgpV4EviVxlan) RouteTargetExport() BgpV4EviVxlanBgpRouteTargetIter {
	if len(obj.obj.RouteTargetExport) == 0 {
		obj.obj.RouteTargetExport = []*otg.BgpRouteTarget{}
	}
	if obj.routeTargetExportHolder == nil {
		obj.routeTargetExportHolder = newBgpV4EviVxlanBgpRouteTargetIter(&obj.obj.RouteTargetExport).setMsg(obj)
	}
	return obj.routeTargetExportHolder
}

type bgpV4EviVxlanBgpRouteTargetIter struct {
	obj                 *bgpV4EviVxlan
	bgpRouteTargetSlice []BgpRouteTarget
	fieldPtr            *[]*otg.BgpRouteTarget
}

func newBgpV4EviVxlanBgpRouteTargetIter(ptr *[]*otg.BgpRouteTarget) BgpV4EviVxlanBgpRouteTargetIter {
	return &bgpV4EviVxlanBgpRouteTargetIter{fieldPtr: ptr}
}

type BgpV4EviVxlanBgpRouteTargetIter interface {
	setMsg(*bgpV4EviVxlan) BgpV4EviVxlanBgpRouteTargetIter
	Items() []BgpRouteTarget
	Add() BgpRouteTarget
	Append(items ...BgpRouteTarget) BgpV4EviVxlanBgpRouteTargetIter
	Set(index int, newObj BgpRouteTarget) BgpV4EviVxlanBgpRouteTargetIter
	Clear() BgpV4EviVxlanBgpRouteTargetIter
	clearHolderSlice() BgpV4EviVxlanBgpRouteTargetIter
	appendHolderSlice(item BgpRouteTarget) BgpV4EviVxlanBgpRouteTargetIter
}

func (obj *bgpV4EviVxlanBgpRouteTargetIter) setMsg(msg *bgpV4EviVxlan) BgpV4EviVxlanBgpRouteTargetIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpRouteTarget{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV4EviVxlanBgpRouteTargetIter) Items() []BgpRouteTarget {
	return obj.bgpRouteTargetSlice
}

func (obj *bgpV4EviVxlanBgpRouteTargetIter) Add() BgpRouteTarget {
	newObj := &otg.BgpRouteTarget{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpRouteTarget{obj: newObj}
	newLibObj.setDefault()
	obj.bgpRouteTargetSlice = append(obj.bgpRouteTargetSlice, newLibObj)
	return newLibObj
}

func (obj *bgpV4EviVxlanBgpRouteTargetIter) Append(items ...BgpRouteTarget) BgpV4EviVxlanBgpRouteTargetIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpRouteTargetSlice = append(obj.bgpRouteTargetSlice, item)
	}
	return obj
}

func (obj *bgpV4EviVxlanBgpRouteTargetIter) Set(index int, newObj BgpRouteTarget) BgpV4EviVxlanBgpRouteTargetIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpRouteTargetSlice[index] = newObj
	return obj
}
func (obj *bgpV4EviVxlanBgpRouteTargetIter) Clear() BgpV4EviVxlanBgpRouteTargetIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpRouteTarget{}
		obj.bgpRouteTargetSlice = []BgpRouteTarget{}
	}
	return obj
}
func (obj *bgpV4EviVxlanBgpRouteTargetIter) clearHolderSlice() BgpV4EviVxlanBgpRouteTargetIter {
	if len(obj.bgpRouteTargetSlice) > 0 {
		obj.bgpRouteTargetSlice = []BgpRouteTarget{}
	}
	return obj
}
func (obj *bgpV4EviVxlanBgpRouteTargetIter) appendHolderSlice(item BgpRouteTarget) BgpV4EviVxlanBgpRouteTargetIter {
	obj.bgpRouteTargetSlice = append(obj.bgpRouteTargetSlice, item)
	return obj
}

// List of L2VNI import targets associated with this EVI.
// RouteTargetImport returns a []BgpRouteTarget
func (obj *bgpV4EviVxlan) RouteTargetImport() BgpV4EviVxlanBgpRouteTargetIter {
	if len(obj.obj.RouteTargetImport) == 0 {
		obj.obj.RouteTargetImport = []*otg.BgpRouteTarget{}
	}
	if obj.routeTargetImportHolder == nil {
		obj.routeTargetImportHolder = newBgpV4EviVxlanBgpRouteTargetIter(&obj.obj.RouteTargetImport).setMsg(obj)
	}
	return obj.routeTargetImportHolder
}

// List of Layer 3 Virtual Network Identifier (L3VNI) Export Route Targets.
// L3RouteTargetExport returns a []BgpRouteTarget
func (obj *bgpV4EviVxlan) L3RouteTargetExport() BgpV4EviVxlanBgpRouteTargetIter {
	if len(obj.obj.L3RouteTargetExport) == 0 {
		obj.obj.L3RouteTargetExport = []*otg.BgpRouteTarget{}
	}
	if obj.l3RouteTargetExportHolder == nil {
		obj.l3RouteTargetExportHolder = newBgpV4EviVxlanBgpRouteTargetIter(&obj.obj.L3RouteTargetExport).setMsg(obj)
	}
	return obj.l3RouteTargetExportHolder
}

// List of L3VNI Import Route Targets.
// L3RouteTargetImport returns a []BgpRouteTarget
func (obj *bgpV4EviVxlan) L3RouteTargetImport() BgpV4EviVxlanBgpRouteTargetIter {
	if len(obj.obj.L3RouteTargetImport) == 0 {
		obj.obj.L3RouteTargetImport = []*otg.BgpRouteTarget{}
	}
	if obj.l3RouteTargetImportHolder == nil {
		obj.l3RouteTargetImportHolder = newBgpV4EviVxlanBgpRouteTargetIter(&obj.obj.L3RouteTargetImport).setMsg(obj)
	}
	return obj.l3RouteTargetImportHolder
}

// description is TBD
// Advanced returns a BgpRouteAdvanced
func (obj *bgpV4EviVxlan) Advanced() BgpRouteAdvanced {
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
func (obj *bgpV4EviVxlan) HasAdvanced() bool {
	return obj.obj.Advanced != nil
}

// description is TBD
// SetAdvanced sets the BgpRouteAdvanced value in the BgpV4EviVxlan object
func (obj *bgpV4EviVxlan) SetAdvanced(value BgpRouteAdvanced) BgpV4EviVxlan {

	obj.advancedHolder = nil
	obj.obj.Advanced = value.msg()

	return obj
}

// Optional community settings.
// Communities returns a []BgpCommunity
func (obj *bgpV4EviVxlan) Communities() BgpV4EviVxlanBgpCommunityIter {
	if len(obj.obj.Communities) == 0 {
		obj.obj.Communities = []*otg.BgpCommunity{}
	}
	if obj.communitiesHolder == nil {
		obj.communitiesHolder = newBgpV4EviVxlanBgpCommunityIter(&obj.obj.Communities).setMsg(obj)
	}
	return obj.communitiesHolder
}

type bgpV4EviVxlanBgpCommunityIter struct {
	obj               *bgpV4EviVxlan
	bgpCommunitySlice []BgpCommunity
	fieldPtr          *[]*otg.BgpCommunity
}

func newBgpV4EviVxlanBgpCommunityIter(ptr *[]*otg.BgpCommunity) BgpV4EviVxlanBgpCommunityIter {
	return &bgpV4EviVxlanBgpCommunityIter{fieldPtr: ptr}
}

type BgpV4EviVxlanBgpCommunityIter interface {
	setMsg(*bgpV4EviVxlan) BgpV4EviVxlanBgpCommunityIter
	Items() []BgpCommunity
	Add() BgpCommunity
	Append(items ...BgpCommunity) BgpV4EviVxlanBgpCommunityIter
	Set(index int, newObj BgpCommunity) BgpV4EviVxlanBgpCommunityIter
	Clear() BgpV4EviVxlanBgpCommunityIter
	clearHolderSlice() BgpV4EviVxlanBgpCommunityIter
	appendHolderSlice(item BgpCommunity) BgpV4EviVxlanBgpCommunityIter
}

func (obj *bgpV4EviVxlanBgpCommunityIter) setMsg(msg *bgpV4EviVxlan) BgpV4EviVxlanBgpCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV4EviVxlanBgpCommunityIter) Items() []BgpCommunity {
	return obj.bgpCommunitySlice
}

func (obj *bgpV4EviVxlanBgpCommunityIter) Add() BgpCommunity {
	newObj := &otg.BgpCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpV4EviVxlanBgpCommunityIter) Append(items ...BgpCommunity) BgpV4EviVxlanBgpCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, item)
	}
	return obj
}

func (obj *bgpV4EviVxlanBgpCommunityIter) Set(index int, newObj BgpCommunity) BgpV4EviVxlanBgpCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpV4EviVxlanBgpCommunityIter) Clear() BgpV4EviVxlanBgpCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpCommunity{}
		obj.bgpCommunitySlice = []BgpCommunity{}
	}
	return obj
}
func (obj *bgpV4EviVxlanBgpCommunityIter) clearHolderSlice() BgpV4EviVxlanBgpCommunityIter {
	if len(obj.bgpCommunitySlice) > 0 {
		obj.bgpCommunitySlice = []BgpCommunity{}
	}
	return obj
}
func (obj *bgpV4EviVxlanBgpCommunityIter) appendHolderSlice(item BgpCommunity) BgpV4EviVxlanBgpCommunityIter {
	obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, item)
	return obj
}

// Optional Extended Community settings. The Extended Communities Attribute is a transitive optional BGP attribute, with the Type Code 16. Community and Extended Communities  attributes are utilized to trigger routing decisions, such as acceptance, rejection,  preference, or redistribution. An extended community is an 8-Bytes value. It is divided into two main parts. The first 2 Bytes of the community encode a type and sub-type fields and the last 6 Bytes carry a unique set of data in a format defined by the type and sub-type field. Extended communities provide a larger  range for grouping or categorizing communities. When type is administrator_as_2octet or administrator_as_4octet, the valid sub types are route target and origin. The valid value for  administrator_as_2octet and administrator_as_4octet type is either two byte AS followed by four byte local administrator id or four byte AS followed by two  byte local administrator id.  When type is administrator_ipv4_address the valid sub types are route target and origin. The valid value for  administrator_ipv4_address is a four byte IPv4 address followed by a two byte local administrator id.  When type is opaque, valid sub types are color and encapsulation. When sub type is color, first two bytes of the value field contain flags and last four bytes  contains the value of the color. When sub type is encapsulation the first four bytes of value field are reserved and last two bytes carries the tunnel type from  IANA's "ETHER TYPES" registry e.g IPv4 (protocol type = 0x0800), IPv6 (protocol type = 0x86dd), and MPLS (protocol type = 0x8847). When type is administrator_as_2octet_link_bandwidth the valid sub type is extended_bandwidth. The first two bytes of the value field contains the AS number and the last four bytes contains the bandwidth in IEEE floating point format.  When type is evpn the valid subtype is mac_address. In the value field the low-order bit of the first byte(Flags) is defined as the "Sticky/static" flag and may be set to 1, indicating the MAC address is static and cannot move. The second byte is reserved and the  last four bytes contain the sequence number which is used to ensure that PEs retain the correct MAC/IP Advertisement route when multiple updates  occur for the same MAC address.
// ExtCommunities returns a []BgpExtCommunity
func (obj *bgpV4EviVxlan) ExtCommunities() BgpV4EviVxlanBgpExtCommunityIter {
	if len(obj.obj.ExtCommunities) == 0 {
		obj.obj.ExtCommunities = []*otg.BgpExtCommunity{}
	}
	if obj.extCommunitiesHolder == nil {
		obj.extCommunitiesHolder = newBgpV4EviVxlanBgpExtCommunityIter(&obj.obj.ExtCommunities).setMsg(obj)
	}
	return obj.extCommunitiesHolder
}

type bgpV4EviVxlanBgpExtCommunityIter struct {
	obj                  *bgpV4EviVxlan
	bgpExtCommunitySlice []BgpExtCommunity
	fieldPtr             *[]*otg.BgpExtCommunity
}

func newBgpV4EviVxlanBgpExtCommunityIter(ptr *[]*otg.BgpExtCommunity) BgpV4EviVxlanBgpExtCommunityIter {
	return &bgpV4EviVxlanBgpExtCommunityIter{fieldPtr: ptr}
}

type BgpV4EviVxlanBgpExtCommunityIter interface {
	setMsg(*bgpV4EviVxlan) BgpV4EviVxlanBgpExtCommunityIter
	Items() []BgpExtCommunity
	Add() BgpExtCommunity
	Append(items ...BgpExtCommunity) BgpV4EviVxlanBgpExtCommunityIter
	Set(index int, newObj BgpExtCommunity) BgpV4EviVxlanBgpExtCommunityIter
	Clear() BgpV4EviVxlanBgpExtCommunityIter
	clearHolderSlice() BgpV4EviVxlanBgpExtCommunityIter
	appendHolderSlice(item BgpExtCommunity) BgpV4EviVxlanBgpExtCommunityIter
}

func (obj *bgpV4EviVxlanBgpExtCommunityIter) setMsg(msg *bgpV4EviVxlan) BgpV4EviVxlanBgpExtCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpExtCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV4EviVxlanBgpExtCommunityIter) Items() []BgpExtCommunity {
	return obj.bgpExtCommunitySlice
}

func (obj *bgpV4EviVxlanBgpExtCommunityIter) Add() BgpExtCommunity {
	newObj := &otg.BgpExtCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpExtCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.bgpExtCommunitySlice = append(obj.bgpExtCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpV4EviVxlanBgpExtCommunityIter) Append(items ...BgpExtCommunity) BgpV4EviVxlanBgpExtCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpExtCommunitySlice = append(obj.bgpExtCommunitySlice, item)
	}
	return obj
}

func (obj *bgpV4EviVxlanBgpExtCommunityIter) Set(index int, newObj BgpExtCommunity) BgpV4EviVxlanBgpExtCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpExtCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpV4EviVxlanBgpExtCommunityIter) Clear() BgpV4EviVxlanBgpExtCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpExtCommunity{}
		obj.bgpExtCommunitySlice = []BgpExtCommunity{}
	}
	return obj
}
func (obj *bgpV4EviVxlanBgpExtCommunityIter) clearHolderSlice() BgpV4EviVxlanBgpExtCommunityIter {
	if len(obj.bgpExtCommunitySlice) > 0 {
		obj.bgpExtCommunitySlice = []BgpExtCommunity{}
	}
	return obj
}
func (obj *bgpV4EviVxlanBgpExtCommunityIter) appendHolderSlice(item BgpExtCommunity) BgpV4EviVxlanBgpExtCommunityIter {
	obj.bgpExtCommunitySlice = append(obj.bgpExtCommunitySlice, item)
	return obj
}

// Optional AS PATH settings.
// AsPath returns a BgpAsPath
func (obj *bgpV4EviVxlan) AsPath() BgpAsPath {
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
func (obj *bgpV4EviVxlan) HasAsPath() bool {
	return obj.obj.AsPath != nil
}

// Optional AS PATH settings.
// SetAsPath sets the BgpAsPath value in the BgpV4EviVxlan object
func (obj *bgpV4EviVxlan) SetAsPath(value BgpAsPath) BgpV4EviVxlan {

	obj.asPathHolder = nil
	obj.obj.AsPath = value.msg()

	return obj
}

func (obj *bgpV4EviVxlan) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.BroadcastDomains) != 0 {

		if set_default {
			obj.BroadcastDomains().clearHolderSlice()
			for _, item := range obj.obj.BroadcastDomains {
				obj.BroadcastDomains().appendHolderSlice(&bgpV4EviVxlanBroadcastDomain{obj: item})
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
				fmt.Sprintf("0 <= BgpV4EviVxlan.PmsiLabel <= 16777215 but Got %d", *obj.obj.PmsiLabel))
		}

	}

	if obj.obj.AdLabel != nil {

		if *obj.obj.AdLabel > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpV4EviVxlan.AdLabel <= 16777215 but Got %d", *obj.obj.AdLabel))
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

func (obj *bgpV4EviVxlan) setDefault() {
	if obj.obj.ReplicationType == nil {
		obj.SetReplicationType(BgpV4EviVxlanReplicationType.INGRESS_REPLICATION)

	}
	if obj.obj.PmsiLabel == nil {
		obj.SetPmsiLabel(16)
	}
	if obj.obj.AdLabel == nil {
		obj.SetAdLabel(0)
	}

}
