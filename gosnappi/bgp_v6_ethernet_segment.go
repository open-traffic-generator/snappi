package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpV6EthernetSegment *****
type bgpV6EthernetSegment struct {
	validation
	obj                  *otg.BgpV6EthernetSegment
	marshaller           marshalBgpV6EthernetSegment
	unMarshaller         unMarshalBgpV6EthernetSegment
	dfElectionHolder     BgpEthernetSegmentDfElection
	evisHolder           BgpV6EthernetSegmentBgpV6EvpnEvisIter
	advancedHolder       BgpRouteAdvanced
	communitiesHolder    BgpV6EthernetSegmentBgpCommunityIter
	extCommunitiesHolder BgpV6EthernetSegmentBgpExtCommunityIter
	asPathHolder         BgpAsPath
}

func NewBgpV6EthernetSegment() BgpV6EthernetSegment {
	obj := bgpV6EthernetSegment{obj: &otg.BgpV6EthernetSegment{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpV6EthernetSegment) msg() *otg.BgpV6EthernetSegment {
	return obj.obj
}

func (obj *bgpV6EthernetSegment) setMsg(msg *otg.BgpV6EthernetSegment) BgpV6EthernetSegment {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpV6EthernetSegment struct {
	obj *bgpV6EthernetSegment
}

type marshalBgpV6EthernetSegment interface {
	// ToProto marshals BgpV6EthernetSegment to protobuf object *otg.BgpV6EthernetSegment
	ToProto() (*otg.BgpV6EthernetSegment, error)
	// ToPbText marshals BgpV6EthernetSegment to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpV6EthernetSegment to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpV6EthernetSegment to JSON text
	ToJson() (string, error)
}

type unMarshalbgpV6EthernetSegment struct {
	obj *bgpV6EthernetSegment
}

type unMarshalBgpV6EthernetSegment interface {
	// FromProto unmarshals BgpV6EthernetSegment from protobuf object *otg.BgpV6EthernetSegment
	FromProto(msg *otg.BgpV6EthernetSegment) (BgpV6EthernetSegment, error)
	// FromPbText unmarshals BgpV6EthernetSegment from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpV6EthernetSegment from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpV6EthernetSegment from JSON text
	FromJson(value string) error
}

func (obj *bgpV6EthernetSegment) Marshal() marshalBgpV6EthernetSegment {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpV6EthernetSegment{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpV6EthernetSegment) Unmarshal() unMarshalBgpV6EthernetSegment {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpV6EthernetSegment{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpV6EthernetSegment) ToProto() (*otg.BgpV6EthernetSegment, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpV6EthernetSegment) FromProto(msg *otg.BgpV6EthernetSegment) (BgpV6EthernetSegment, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpV6EthernetSegment) ToPbText() (string, error) {
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

func (m *unMarshalbgpV6EthernetSegment) FromPbText(value string) error {
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

func (m *marshalbgpV6EthernetSegment) ToYaml() (string, error) {
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

func (m *unMarshalbgpV6EthernetSegment) FromYaml(value string) error {
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

func (m *marshalbgpV6EthernetSegment) ToJson() (string, error) {
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

func (m *unMarshalbgpV6EthernetSegment) FromJson(value string) error {
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

func (obj *bgpV6EthernetSegment) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpV6EthernetSegment) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpV6EthernetSegment) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpV6EthernetSegment) Clone() (BgpV6EthernetSegment, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpV6EthernetSegment()
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

func (obj *bgpV6EthernetSegment) setNil() {
	obj.dfElectionHolder = nil
	obj.evisHolder = nil
	obj.advancedHolder = nil
	obj.communitiesHolder = nil
	obj.extCommunitiesHolder = nil
	obj.asPathHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpV6EthernetSegment is configuration for BGP Ethernet Segment ranges. Advertises following routes -
//
// Type 4 - Ethernet Segment Route
type BgpV6EthernetSegment interface {
	Validation
	// msg marshals BgpV6EthernetSegment to protobuf object *otg.BgpV6EthernetSegment
	// and doesn't set defaults
	msg() *otg.BgpV6EthernetSegment
	// setMsg unmarshals BgpV6EthernetSegment from protobuf object *otg.BgpV6EthernetSegment
	// and doesn't set defaults
	setMsg(*otg.BgpV6EthernetSegment) BgpV6EthernetSegment
	// provides marshal interface
	Marshal() marshalBgpV6EthernetSegment
	// provides unmarshal interface
	Unmarshal() unMarshalBgpV6EthernetSegment
	// validate validates BgpV6EthernetSegment
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpV6EthernetSegment, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// DfElection returns BgpEthernetSegmentDfElection, set in BgpV6EthernetSegment.
	// BgpEthernetSegmentDfElection is configuration for Designated Forwarder (DF) election among the Provider Edge (PE) routers on the same Ethernet Segment.
	DfElection() BgpEthernetSegmentDfElection
	// SetDfElection assigns BgpEthernetSegmentDfElection provided by user to BgpV6EthernetSegment.
	// BgpEthernetSegmentDfElection is configuration for Designated Forwarder (DF) election among the Provider Edge (PE) routers on the same Ethernet Segment.
	SetDfElection(value BgpEthernetSegmentDfElection) BgpV6EthernetSegment
	// HasDfElection checks if DfElection has been set in BgpV6EthernetSegment
	HasDfElection() bool
	// Evis returns BgpV6EthernetSegmentBgpV6EvpnEvisIterIter, set in BgpV6EthernetSegment
	Evis() BgpV6EthernetSegmentBgpV6EvpnEvisIter
	// Esi returns string, set in BgpV6EthernetSegment.
	Esi() string
	// SetEsi assigns string provided by user to BgpV6EthernetSegment
	SetEsi(value string) BgpV6EthernetSegment
	// HasEsi checks if Esi has been set in BgpV6EthernetSegment
	HasEsi() bool
	// ActiveMode returns BgpV6EthernetSegmentActiveModeEnum, set in BgpV6EthernetSegment
	ActiveMode() BgpV6EthernetSegmentActiveModeEnum
	// SetActiveMode assigns BgpV6EthernetSegmentActiveModeEnum provided by user to BgpV6EthernetSegment
	SetActiveMode(value BgpV6EthernetSegmentActiveModeEnum) BgpV6EthernetSegment
	// HasActiveMode checks if ActiveMode has been set in BgpV6EthernetSegment
	HasActiveMode() bool
	// EsiLabel returns uint32, set in BgpV6EthernetSegment.
	EsiLabel() uint32
	// SetEsiLabel assigns uint32 provided by user to BgpV6EthernetSegment
	SetEsiLabel(value uint32) BgpV6EthernetSegment
	// HasEsiLabel checks if EsiLabel has been set in BgpV6EthernetSegment
	HasEsiLabel() bool
	// Advanced returns BgpRouteAdvanced, set in BgpV6EthernetSegment.
	// BgpRouteAdvanced is configuration for advanced BGP route range settings.
	Advanced() BgpRouteAdvanced
	// SetAdvanced assigns BgpRouteAdvanced provided by user to BgpV6EthernetSegment.
	// BgpRouteAdvanced is configuration for advanced BGP route range settings.
	SetAdvanced(value BgpRouteAdvanced) BgpV6EthernetSegment
	// HasAdvanced checks if Advanced has been set in BgpV6EthernetSegment
	HasAdvanced() bool
	// Communities returns BgpV6EthernetSegmentBgpCommunityIterIter, set in BgpV6EthernetSegment
	Communities() BgpV6EthernetSegmentBgpCommunityIter
	// ExtCommunities returns BgpV6EthernetSegmentBgpExtCommunityIterIter, set in BgpV6EthernetSegment
	ExtCommunities() BgpV6EthernetSegmentBgpExtCommunityIter
	// AsPath returns BgpAsPath, set in BgpV6EthernetSegment.
	// BgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed. This contains the configuration of how to include the Local AS in the AS path attribute of the MP REACH NLRI. It also contains optional configuration of additional AS Path Segments that can be included in the AS Path attribute. The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that a routing information passes through to reach the destination.
	AsPath() BgpAsPath
	// SetAsPath assigns BgpAsPath provided by user to BgpV6EthernetSegment.
	// BgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed. This contains the configuration of how to include the Local AS in the AS path attribute of the MP REACH NLRI. It also contains optional configuration of additional AS Path Segments that can be included in the AS Path attribute. The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that a routing information passes through to reach the destination.
	SetAsPath(value BgpAsPath) BgpV6EthernetSegment
	// HasAsPath checks if AsPath has been set in BgpV6EthernetSegment
	HasAsPath() bool
	setNil()
}

// Designated Forwarder (DF) election configuration.
// DfElection returns a BgpEthernetSegmentDfElection
func (obj *bgpV6EthernetSegment) DfElection() BgpEthernetSegmentDfElection {
	if obj.obj.DfElection == nil {
		obj.obj.DfElection = NewBgpEthernetSegmentDfElection().msg()
	}
	if obj.dfElectionHolder == nil {
		obj.dfElectionHolder = &bgpEthernetSegmentDfElection{obj: obj.obj.DfElection}
	}
	return obj.dfElectionHolder
}

// Designated Forwarder (DF) election configuration.
// DfElection returns a BgpEthernetSegmentDfElection
func (obj *bgpV6EthernetSegment) HasDfElection() bool {
	return obj.obj.DfElection != nil
}

// Designated Forwarder (DF) election configuration.
// SetDfElection sets the BgpEthernetSegmentDfElection value in the BgpV6EthernetSegment object
func (obj *bgpV6EthernetSegment) SetDfElection(value BgpEthernetSegmentDfElection) BgpV6EthernetSegment {

	obj.dfElectionHolder = nil
	obj.obj.DfElection = value.msg()

	return obj
}

// This contains the list of EVIs.
// Evis returns a []BgpV6EvpnEvis
func (obj *bgpV6EthernetSegment) Evis() BgpV6EthernetSegmentBgpV6EvpnEvisIter {
	if len(obj.obj.Evis) == 0 {
		obj.obj.Evis = []*otg.BgpV6EvpnEvis{}
	}
	if obj.evisHolder == nil {
		obj.evisHolder = newBgpV6EthernetSegmentBgpV6EvpnEvisIter(&obj.obj.Evis).setMsg(obj)
	}
	return obj.evisHolder
}

type bgpV6EthernetSegmentBgpV6EvpnEvisIter struct {
	obj                *bgpV6EthernetSegment
	bgpV6EvpnEvisSlice []BgpV6EvpnEvis
	fieldPtr           *[]*otg.BgpV6EvpnEvis
}

func newBgpV6EthernetSegmentBgpV6EvpnEvisIter(ptr *[]*otg.BgpV6EvpnEvis) BgpV6EthernetSegmentBgpV6EvpnEvisIter {
	return &bgpV6EthernetSegmentBgpV6EvpnEvisIter{fieldPtr: ptr}
}

type BgpV6EthernetSegmentBgpV6EvpnEvisIter interface {
	setMsg(*bgpV6EthernetSegment) BgpV6EthernetSegmentBgpV6EvpnEvisIter
	Items() []BgpV6EvpnEvis
	Add() BgpV6EvpnEvis
	Append(items ...BgpV6EvpnEvis) BgpV6EthernetSegmentBgpV6EvpnEvisIter
	Set(index int, newObj BgpV6EvpnEvis) BgpV6EthernetSegmentBgpV6EvpnEvisIter
	Clear() BgpV6EthernetSegmentBgpV6EvpnEvisIter
	clearHolderSlice() BgpV6EthernetSegmentBgpV6EvpnEvisIter
	appendHolderSlice(item BgpV6EvpnEvis) BgpV6EthernetSegmentBgpV6EvpnEvisIter
}

func (obj *bgpV6EthernetSegmentBgpV6EvpnEvisIter) setMsg(msg *bgpV6EthernetSegment) BgpV6EthernetSegmentBgpV6EvpnEvisIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpV6EvpnEvis{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV6EthernetSegmentBgpV6EvpnEvisIter) Items() []BgpV6EvpnEvis {
	return obj.bgpV6EvpnEvisSlice
}

func (obj *bgpV6EthernetSegmentBgpV6EvpnEvisIter) Add() BgpV6EvpnEvis {
	newObj := &otg.BgpV6EvpnEvis{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpV6EvpnEvis{obj: newObj}
	newLibObj.setDefault()
	obj.bgpV6EvpnEvisSlice = append(obj.bgpV6EvpnEvisSlice, newLibObj)
	return newLibObj
}

func (obj *bgpV6EthernetSegmentBgpV6EvpnEvisIter) Append(items ...BgpV6EvpnEvis) BgpV6EthernetSegmentBgpV6EvpnEvisIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpV6EvpnEvisSlice = append(obj.bgpV6EvpnEvisSlice, item)
	}
	return obj
}

func (obj *bgpV6EthernetSegmentBgpV6EvpnEvisIter) Set(index int, newObj BgpV6EvpnEvis) BgpV6EthernetSegmentBgpV6EvpnEvisIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpV6EvpnEvisSlice[index] = newObj
	return obj
}
func (obj *bgpV6EthernetSegmentBgpV6EvpnEvisIter) Clear() BgpV6EthernetSegmentBgpV6EvpnEvisIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpV6EvpnEvis{}
		obj.bgpV6EvpnEvisSlice = []BgpV6EvpnEvis{}
	}
	return obj
}
func (obj *bgpV6EthernetSegmentBgpV6EvpnEvisIter) clearHolderSlice() BgpV6EthernetSegmentBgpV6EvpnEvisIter {
	if len(obj.bgpV6EvpnEvisSlice) > 0 {
		obj.bgpV6EvpnEvisSlice = []BgpV6EvpnEvis{}
	}
	return obj
}
func (obj *bgpV6EthernetSegmentBgpV6EvpnEvisIter) appendHolderSlice(item BgpV6EvpnEvis) BgpV6EthernetSegmentBgpV6EvpnEvisIter {
	obj.bgpV6EvpnEvisSlice = append(obj.bgpV6EvpnEvisSlice, item)
	return obj
}

// 10-octet Ethernet Segment Identifier (ESI) Example - For multi-home scenario nonZero ESI is '10000000000000000000' .
// Esi returns a string
func (obj *bgpV6EthernetSegment) Esi() string {

	return *obj.obj.Esi

}

// 10-octet Ethernet Segment Identifier (ESI) Example - For multi-home scenario nonZero ESI is '10000000000000000000' .
// Esi returns a string
func (obj *bgpV6EthernetSegment) HasEsi() bool {
	return obj.obj.Esi != nil
}

// 10-octet Ethernet Segment Identifier (ESI) Example - For multi-home scenario nonZero ESI is '10000000000000000000' .
// SetEsi sets the string value in the BgpV6EthernetSegment object
func (obj *bgpV6EthernetSegment) SetEsi(value string) BgpV6EthernetSegment {

	obj.obj.Esi = &value
	return obj
}

type BgpV6EthernetSegmentActiveModeEnum string

// Enum of ActiveMode on BgpV6EthernetSegment
var BgpV6EthernetSegmentActiveMode = struct {
	SINGLE_ACTIVE BgpV6EthernetSegmentActiveModeEnum
	ALL_ACTIVE    BgpV6EthernetSegmentActiveModeEnum
}{
	SINGLE_ACTIVE: BgpV6EthernetSegmentActiveModeEnum("single_active"),
	ALL_ACTIVE:    BgpV6EthernetSegmentActiveModeEnum("all_active"),
}

func (obj *bgpV6EthernetSegment) ActiveMode() BgpV6EthernetSegmentActiveModeEnum {
	return BgpV6EthernetSegmentActiveModeEnum(obj.obj.ActiveMode.Enum().String())
}

// Single Active or All Active mode Redundancy mode selection for Multi-home.
// ActiveMode returns a string
func (obj *bgpV6EthernetSegment) HasActiveMode() bool {
	return obj.obj.ActiveMode != nil
}

func (obj *bgpV6EthernetSegment) SetActiveMode(value BgpV6EthernetSegmentActiveModeEnum) BgpV6EthernetSegment {
	intValue, ok := otg.BgpV6EthernetSegment_ActiveMode_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpV6EthernetSegmentActiveModeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpV6EthernetSegment_ActiveMode_Enum(intValue)
	obj.obj.ActiveMode = &enumValue

	return obj
}

// The label value to be advertised as ESI Label in ESI Label Extended Community. This is included in Ethernet Auto-discovery per ES Routes advertised by a router.
// EsiLabel returns a uint32
func (obj *bgpV6EthernetSegment) EsiLabel() uint32 {

	return *obj.obj.EsiLabel

}

// The label value to be advertised as ESI Label in ESI Label Extended Community. This is included in Ethernet Auto-discovery per ES Routes advertised by a router.
// EsiLabel returns a uint32
func (obj *bgpV6EthernetSegment) HasEsiLabel() bool {
	return obj.obj.EsiLabel != nil
}

// The label value to be advertised as ESI Label in ESI Label Extended Community. This is included in Ethernet Auto-discovery per ES Routes advertised by a router.
// SetEsiLabel sets the uint32 value in the BgpV6EthernetSegment object
func (obj *bgpV6EthernetSegment) SetEsiLabel(value uint32) BgpV6EthernetSegment {

	obj.obj.EsiLabel = &value
	return obj
}

// description is TBD
// Advanced returns a BgpRouteAdvanced
func (obj *bgpV6EthernetSegment) Advanced() BgpRouteAdvanced {
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
func (obj *bgpV6EthernetSegment) HasAdvanced() bool {
	return obj.obj.Advanced != nil
}

// description is TBD
// SetAdvanced sets the BgpRouteAdvanced value in the BgpV6EthernetSegment object
func (obj *bgpV6EthernetSegment) SetAdvanced(value BgpRouteAdvanced) BgpV6EthernetSegment {

	obj.advancedHolder = nil
	obj.obj.Advanced = value.msg()

	return obj
}

// Optional community settings.
// Communities returns a []BgpCommunity
func (obj *bgpV6EthernetSegment) Communities() BgpV6EthernetSegmentBgpCommunityIter {
	if len(obj.obj.Communities) == 0 {
		obj.obj.Communities = []*otg.BgpCommunity{}
	}
	if obj.communitiesHolder == nil {
		obj.communitiesHolder = newBgpV6EthernetSegmentBgpCommunityIter(&obj.obj.Communities).setMsg(obj)
	}
	return obj.communitiesHolder
}

type bgpV6EthernetSegmentBgpCommunityIter struct {
	obj               *bgpV6EthernetSegment
	bgpCommunitySlice []BgpCommunity
	fieldPtr          *[]*otg.BgpCommunity
}

func newBgpV6EthernetSegmentBgpCommunityIter(ptr *[]*otg.BgpCommunity) BgpV6EthernetSegmentBgpCommunityIter {
	return &bgpV6EthernetSegmentBgpCommunityIter{fieldPtr: ptr}
}

type BgpV6EthernetSegmentBgpCommunityIter interface {
	setMsg(*bgpV6EthernetSegment) BgpV6EthernetSegmentBgpCommunityIter
	Items() []BgpCommunity
	Add() BgpCommunity
	Append(items ...BgpCommunity) BgpV6EthernetSegmentBgpCommunityIter
	Set(index int, newObj BgpCommunity) BgpV6EthernetSegmentBgpCommunityIter
	Clear() BgpV6EthernetSegmentBgpCommunityIter
	clearHolderSlice() BgpV6EthernetSegmentBgpCommunityIter
	appendHolderSlice(item BgpCommunity) BgpV6EthernetSegmentBgpCommunityIter
}

func (obj *bgpV6EthernetSegmentBgpCommunityIter) setMsg(msg *bgpV6EthernetSegment) BgpV6EthernetSegmentBgpCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV6EthernetSegmentBgpCommunityIter) Items() []BgpCommunity {
	return obj.bgpCommunitySlice
}

func (obj *bgpV6EthernetSegmentBgpCommunityIter) Add() BgpCommunity {
	newObj := &otg.BgpCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpV6EthernetSegmentBgpCommunityIter) Append(items ...BgpCommunity) BgpV6EthernetSegmentBgpCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, item)
	}
	return obj
}

func (obj *bgpV6EthernetSegmentBgpCommunityIter) Set(index int, newObj BgpCommunity) BgpV6EthernetSegmentBgpCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpV6EthernetSegmentBgpCommunityIter) Clear() BgpV6EthernetSegmentBgpCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpCommunity{}
		obj.bgpCommunitySlice = []BgpCommunity{}
	}
	return obj
}
func (obj *bgpV6EthernetSegmentBgpCommunityIter) clearHolderSlice() BgpV6EthernetSegmentBgpCommunityIter {
	if len(obj.bgpCommunitySlice) > 0 {
		obj.bgpCommunitySlice = []BgpCommunity{}
	}
	return obj
}
func (obj *bgpV6EthernetSegmentBgpCommunityIter) appendHolderSlice(item BgpCommunity) BgpV6EthernetSegmentBgpCommunityIter {
	obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, item)
	return obj
}

// Optional Extended Community settings. The Extended Communities Attribute is a transitive optional BGP attribute, with the Type Code 16. Community and Extended Communities  attributes are utilized to trigger routing decisions, such as acceptance, rejection,  preference, or redistribution. An extended community is an 8-Bytes value. It is divided into two main parts. The first 2 Bytes of the community encode a type and sub-type fields and the last 6 Bytes carry a unique set of data in a format defined by the type and sub-type field. Extended communities provide a larger  range for grouping or categorizing communities. When type is administrator_as_2octet or administrator_as_4octet, the valid sub types are route target and origin. The valid value for  administrator_as_2octet and administrator_as_4octet type is either two byte AS followed by four byte local administrator id or four byte AS followed by two  byte local administrator id.  When type is administrator_ipv4_address the valid sub types are route target and origin. The valid value for  administrator_ipv4_address is a four byte IPv4 address followed by a two byte local administrator id.  When type is opaque, valid sub types are color and encapsulation. When sub type is color, first two bytes of the value field contain flags and last four bytes  contains the value of the color. When sub type is encapsulation the first four bytes of value field are reserved and last two bytes carries the tunnel type from  IANA's "ETHER TYPES" registry e.g IPv4 (protocol type = 0x0800), IPv6 (protocol type = 0x86dd), and MPLS (protocol type = 0x8847). When type is administrator_as_2octet_link_bandwidth the valid sub type is extended_bandwidth. The first two bytes of the value field contains the AS number and the last four bytes contains the bandwidth in IEEE floating point format.  When type is evpn the valid subtype is mac_address. In the value field the low-order bit of the first byte(Flags) is defined as the "Sticky/static" flag and may be set to 1, indicating the MAC address is static and cannot move. The second byte is reserved and the  last four bytes contain the sequence number which is used to ensure that PEs retain the correct MAC/IP Advertisement route when multiple updates  occur for the same MAC address.
// ExtCommunities returns a []BgpExtCommunity
func (obj *bgpV6EthernetSegment) ExtCommunities() BgpV6EthernetSegmentBgpExtCommunityIter {
	if len(obj.obj.ExtCommunities) == 0 {
		obj.obj.ExtCommunities = []*otg.BgpExtCommunity{}
	}
	if obj.extCommunitiesHolder == nil {
		obj.extCommunitiesHolder = newBgpV6EthernetSegmentBgpExtCommunityIter(&obj.obj.ExtCommunities).setMsg(obj)
	}
	return obj.extCommunitiesHolder
}

type bgpV6EthernetSegmentBgpExtCommunityIter struct {
	obj                  *bgpV6EthernetSegment
	bgpExtCommunitySlice []BgpExtCommunity
	fieldPtr             *[]*otg.BgpExtCommunity
}

func newBgpV6EthernetSegmentBgpExtCommunityIter(ptr *[]*otg.BgpExtCommunity) BgpV6EthernetSegmentBgpExtCommunityIter {
	return &bgpV6EthernetSegmentBgpExtCommunityIter{fieldPtr: ptr}
}

type BgpV6EthernetSegmentBgpExtCommunityIter interface {
	setMsg(*bgpV6EthernetSegment) BgpV6EthernetSegmentBgpExtCommunityIter
	Items() []BgpExtCommunity
	Add() BgpExtCommunity
	Append(items ...BgpExtCommunity) BgpV6EthernetSegmentBgpExtCommunityIter
	Set(index int, newObj BgpExtCommunity) BgpV6EthernetSegmentBgpExtCommunityIter
	Clear() BgpV6EthernetSegmentBgpExtCommunityIter
	clearHolderSlice() BgpV6EthernetSegmentBgpExtCommunityIter
	appendHolderSlice(item BgpExtCommunity) BgpV6EthernetSegmentBgpExtCommunityIter
}

func (obj *bgpV6EthernetSegmentBgpExtCommunityIter) setMsg(msg *bgpV6EthernetSegment) BgpV6EthernetSegmentBgpExtCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpExtCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV6EthernetSegmentBgpExtCommunityIter) Items() []BgpExtCommunity {
	return obj.bgpExtCommunitySlice
}

func (obj *bgpV6EthernetSegmentBgpExtCommunityIter) Add() BgpExtCommunity {
	newObj := &otg.BgpExtCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpExtCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.bgpExtCommunitySlice = append(obj.bgpExtCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpV6EthernetSegmentBgpExtCommunityIter) Append(items ...BgpExtCommunity) BgpV6EthernetSegmentBgpExtCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpExtCommunitySlice = append(obj.bgpExtCommunitySlice, item)
	}
	return obj
}

func (obj *bgpV6EthernetSegmentBgpExtCommunityIter) Set(index int, newObj BgpExtCommunity) BgpV6EthernetSegmentBgpExtCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpExtCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpV6EthernetSegmentBgpExtCommunityIter) Clear() BgpV6EthernetSegmentBgpExtCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpExtCommunity{}
		obj.bgpExtCommunitySlice = []BgpExtCommunity{}
	}
	return obj
}
func (obj *bgpV6EthernetSegmentBgpExtCommunityIter) clearHolderSlice() BgpV6EthernetSegmentBgpExtCommunityIter {
	if len(obj.bgpExtCommunitySlice) > 0 {
		obj.bgpExtCommunitySlice = []BgpExtCommunity{}
	}
	return obj
}
func (obj *bgpV6EthernetSegmentBgpExtCommunityIter) appendHolderSlice(item BgpExtCommunity) BgpV6EthernetSegmentBgpExtCommunityIter {
	obj.bgpExtCommunitySlice = append(obj.bgpExtCommunitySlice, item)
	return obj
}

// Optional AS PATH settings.
// AsPath returns a BgpAsPath
func (obj *bgpV6EthernetSegment) AsPath() BgpAsPath {
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
func (obj *bgpV6EthernetSegment) HasAsPath() bool {
	return obj.obj.AsPath != nil
}

// Optional AS PATH settings.
// SetAsPath sets the BgpAsPath value in the BgpV6EthernetSegment object
func (obj *bgpV6EthernetSegment) SetAsPath(value BgpAsPath) BgpV6EthernetSegment {

	obj.asPathHolder = nil
	obj.obj.AsPath = value.msg()

	return obj
}

func (obj *bgpV6EthernetSegment) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.DfElection != nil {

		obj.DfElection().validateObj(vObj, set_default)
	}

	if len(obj.obj.Evis) != 0 {

		if set_default {
			obj.Evis().clearHolderSlice()
			for _, item := range obj.obj.Evis {
				obj.Evis().appendHolderSlice(&bgpV6EvpnEvis{obj: item})
			}
		}
		for _, item := range obj.Evis().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.Esi != nil {

		err := obj.validateHex(obj.Esi())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpV6EthernetSegment.Esi"))
		}

	}

	if obj.obj.EsiLabel != nil {

		if *obj.obj.EsiLabel > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpV6EthernetSegment.EsiLabel <= 16777215 but Got %d", *obj.obj.EsiLabel))
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

func (obj *bgpV6EthernetSegment) setDefault() {
	if obj.obj.Esi == nil {
		obj.SetEsi("00000000000000000000")
	}
	if obj.obj.ActiveMode == nil {
		obj.SetActiveMode(BgpV6EthernetSegmentActiveMode.ALL_ACTIVE)

	}
	if obj.obj.EsiLabel == nil {
		obj.SetEsiLabel(0)
	}

}
