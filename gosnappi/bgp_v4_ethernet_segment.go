package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpV4EthernetSegment *****
type bgpV4EthernetSegment struct {
	validation
	obj                  *otg.BgpV4EthernetSegment
	marshaller           marshalBgpV4EthernetSegment
	unMarshaller         unMarshalBgpV4EthernetSegment
	dfElectionHolder     BgpEthernetSegmentDfElection
	evisHolder           BgpV4EthernetSegmentBgpV4EvpnEvisIter
	advancedHolder       BgpRouteAdvanced
	communitiesHolder    BgpV4EthernetSegmentBgpCommunityIter
	extCommunitiesHolder BgpV4EthernetSegmentBgpExtCommunityIter
	asPathHolder         BgpAsPath
}

func NewBgpV4EthernetSegment() BgpV4EthernetSegment {
	obj := bgpV4EthernetSegment{obj: &otg.BgpV4EthernetSegment{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpV4EthernetSegment) msg() *otg.BgpV4EthernetSegment {
	return obj.obj
}

func (obj *bgpV4EthernetSegment) setMsg(msg *otg.BgpV4EthernetSegment) BgpV4EthernetSegment {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpV4EthernetSegment struct {
	obj *bgpV4EthernetSegment
}

type marshalBgpV4EthernetSegment interface {
	// ToProto marshals BgpV4EthernetSegment to protobuf object *otg.BgpV4EthernetSegment
	ToProto() (*otg.BgpV4EthernetSegment, error)
	// ToPbText marshals BgpV4EthernetSegment to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpV4EthernetSegment to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpV4EthernetSegment to JSON text
	ToJson() (string, error)
}

type unMarshalbgpV4EthernetSegment struct {
	obj *bgpV4EthernetSegment
}

type unMarshalBgpV4EthernetSegment interface {
	// FromProto unmarshals BgpV4EthernetSegment from protobuf object *otg.BgpV4EthernetSegment
	FromProto(msg *otg.BgpV4EthernetSegment) (BgpV4EthernetSegment, error)
	// FromPbText unmarshals BgpV4EthernetSegment from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpV4EthernetSegment from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpV4EthernetSegment from JSON text
	FromJson(value string) error
}

func (obj *bgpV4EthernetSegment) Marshal() marshalBgpV4EthernetSegment {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpV4EthernetSegment{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpV4EthernetSegment) Unmarshal() unMarshalBgpV4EthernetSegment {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpV4EthernetSegment{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpV4EthernetSegment) ToProto() (*otg.BgpV4EthernetSegment, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpV4EthernetSegment) FromProto(msg *otg.BgpV4EthernetSegment) (BgpV4EthernetSegment, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpV4EthernetSegment) ToPbText() (string, error) {
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

func (m *unMarshalbgpV4EthernetSegment) FromPbText(value string) error {
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

func (m *marshalbgpV4EthernetSegment) ToYaml() (string, error) {
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

func (m *unMarshalbgpV4EthernetSegment) FromYaml(value string) error {
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

func (m *marshalbgpV4EthernetSegment) ToJson() (string, error) {
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

func (m *unMarshalbgpV4EthernetSegment) FromJson(value string) error {
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

func (obj *bgpV4EthernetSegment) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpV4EthernetSegment) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpV4EthernetSegment) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpV4EthernetSegment) Clone() (BgpV4EthernetSegment, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpV4EthernetSegment()
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

func (obj *bgpV4EthernetSegment) setNil() {
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

// BgpV4EthernetSegment is configuration for BGP Ethernet Segment ranges. Advertises following routes -
//
// Type 4 - Ethernet Segment Route
type BgpV4EthernetSegment interface {
	Validation
	// msg marshals BgpV4EthernetSegment to protobuf object *otg.BgpV4EthernetSegment
	// and doesn't set defaults
	msg() *otg.BgpV4EthernetSegment
	// setMsg unmarshals BgpV4EthernetSegment from protobuf object *otg.BgpV4EthernetSegment
	// and doesn't set defaults
	setMsg(*otg.BgpV4EthernetSegment) BgpV4EthernetSegment
	// provides marshal interface
	Marshal() marshalBgpV4EthernetSegment
	// provides unmarshal interface
	Unmarshal() unMarshalBgpV4EthernetSegment
	// validate validates BgpV4EthernetSegment
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpV4EthernetSegment, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// DfElection returns BgpEthernetSegmentDfElection, set in BgpV4EthernetSegment.
	// BgpEthernetSegmentDfElection is configuration for Designated Forwarder (DF) election among the Provider Edge (PE) routers on the same Ethernet Segment.
	DfElection() BgpEthernetSegmentDfElection
	// SetDfElection assigns BgpEthernetSegmentDfElection provided by user to BgpV4EthernetSegment.
	// BgpEthernetSegmentDfElection is configuration for Designated Forwarder (DF) election among the Provider Edge (PE) routers on the same Ethernet Segment.
	SetDfElection(value BgpEthernetSegmentDfElection) BgpV4EthernetSegment
	// HasDfElection checks if DfElection has been set in BgpV4EthernetSegment
	HasDfElection() bool
	// Evis returns BgpV4EthernetSegmentBgpV4EvpnEvisIterIter, set in BgpV4EthernetSegment
	Evis() BgpV4EthernetSegmentBgpV4EvpnEvisIter
	// Esi returns string, set in BgpV4EthernetSegment.
	Esi() string
	// SetEsi assigns string provided by user to BgpV4EthernetSegment
	SetEsi(value string) BgpV4EthernetSegment
	// HasEsi checks if Esi has been set in BgpV4EthernetSegment
	HasEsi() bool
	// ActiveMode returns BgpV4EthernetSegmentActiveModeEnum, set in BgpV4EthernetSegment
	ActiveMode() BgpV4EthernetSegmentActiveModeEnum
	// SetActiveMode assigns BgpV4EthernetSegmentActiveModeEnum provided by user to BgpV4EthernetSegment
	SetActiveMode(value BgpV4EthernetSegmentActiveModeEnum) BgpV4EthernetSegment
	// HasActiveMode checks if ActiveMode has been set in BgpV4EthernetSegment
	HasActiveMode() bool
	// EsiLabel returns uint32, set in BgpV4EthernetSegment.
	EsiLabel() uint32
	// SetEsiLabel assigns uint32 provided by user to BgpV4EthernetSegment
	SetEsiLabel(value uint32) BgpV4EthernetSegment
	// HasEsiLabel checks if EsiLabel has been set in BgpV4EthernetSegment
	HasEsiLabel() bool
	// Advanced returns BgpRouteAdvanced, set in BgpV4EthernetSegment.
	// BgpRouteAdvanced is configuration for advanced BGP route range settings.
	Advanced() BgpRouteAdvanced
	// SetAdvanced assigns BgpRouteAdvanced provided by user to BgpV4EthernetSegment.
	// BgpRouteAdvanced is configuration for advanced BGP route range settings.
	SetAdvanced(value BgpRouteAdvanced) BgpV4EthernetSegment
	// HasAdvanced checks if Advanced has been set in BgpV4EthernetSegment
	HasAdvanced() bool
	// Communities returns BgpV4EthernetSegmentBgpCommunityIterIter, set in BgpV4EthernetSegment
	Communities() BgpV4EthernetSegmentBgpCommunityIter
	// ExtCommunities returns BgpV4EthernetSegmentBgpExtCommunityIterIter, set in BgpV4EthernetSegment
	ExtCommunities() BgpV4EthernetSegmentBgpExtCommunityIter
	// AsPath returns BgpAsPath, set in BgpV4EthernetSegment.
	// BgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed. This contains the configuration of how to include the Local AS in the AS path attribute of the MP REACH NLRI. It also contains optional configuration of additional AS Path Segments that can be included in the AS Path attribute. The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that a routing information passes through to reach the destination.
	AsPath() BgpAsPath
	// SetAsPath assigns BgpAsPath provided by user to BgpV4EthernetSegment.
	// BgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed. This contains the configuration of how to include the Local AS in the AS path attribute of the MP REACH NLRI. It also contains optional configuration of additional AS Path Segments that can be included in the AS Path attribute. The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that a routing information passes through to reach the destination.
	SetAsPath(value BgpAsPath) BgpV4EthernetSegment
	// HasAsPath checks if AsPath has been set in BgpV4EthernetSegment
	HasAsPath() bool
	setNil()
}

// Designated Forwarder (DF) election configuration.
// DfElection returns a BgpEthernetSegmentDfElection
func (obj *bgpV4EthernetSegment) DfElection() BgpEthernetSegmentDfElection {
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
func (obj *bgpV4EthernetSegment) HasDfElection() bool {
	return obj.obj.DfElection != nil
}

// Designated Forwarder (DF) election configuration.
// SetDfElection sets the BgpEthernetSegmentDfElection value in the BgpV4EthernetSegment object
func (obj *bgpV4EthernetSegment) SetDfElection(value BgpEthernetSegmentDfElection) BgpV4EthernetSegment {

	obj.dfElectionHolder = nil
	obj.obj.DfElection = value.msg()

	return obj
}

// This contains the list of EVIs.
// Evis returns a []BgpV4EvpnEvis
func (obj *bgpV4EthernetSegment) Evis() BgpV4EthernetSegmentBgpV4EvpnEvisIter {
	if len(obj.obj.Evis) == 0 {
		obj.obj.Evis = []*otg.BgpV4EvpnEvis{}
	}
	if obj.evisHolder == nil {
		obj.evisHolder = newBgpV4EthernetSegmentBgpV4EvpnEvisIter(&obj.obj.Evis).setMsg(obj)
	}
	return obj.evisHolder
}

type bgpV4EthernetSegmentBgpV4EvpnEvisIter struct {
	obj                *bgpV4EthernetSegment
	bgpV4EvpnEvisSlice []BgpV4EvpnEvis
	fieldPtr           *[]*otg.BgpV4EvpnEvis
}

func newBgpV4EthernetSegmentBgpV4EvpnEvisIter(ptr *[]*otg.BgpV4EvpnEvis) BgpV4EthernetSegmentBgpV4EvpnEvisIter {
	return &bgpV4EthernetSegmentBgpV4EvpnEvisIter{fieldPtr: ptr}
}

type BgpV4EthernetSegmentBgpV4EvpnEvisIter interface {
	setMsg(*bgpV4EthernetSegment) BgpV4EthernetSegmentBgpV4EvpnEvisIter
	Items() []BgpV4EvpnEvis
	Add() BgpV4EvpnEvis
	Append(items ...BgpV4EvpnEvis) BgpV4EthernetSegmentBgpV4EvpnEvisIter
	Set(index int, newObj BgpV4EvpnEvis) BgpV4EthernetSegmentBgpV4EvpnEvisIter
	Clear() BgpV4EthernetSegmentBgpV4EvpnEvisIter
	clearHolderSlice() BgpV4EthernetSegmentBgpV4EvpnEvisIter
	appendHolderSlice(item BgpV4EvpnEvis) BgpV4EthernetSegmentBgpV4EvpnEvisIter
}

func (obj *bgpV4EthernetSegmentBgpV4EvpnEvisIter) setMsg(msg *bgpV4EthernetSegment) BgpV4EthernetSegmentBgpV4EvpnEvisIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpV4EvpnEvis{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV4EthernetSegmentBgpV4EvpnEvisIter) Items() []BgpV4EvpnEvis {
	return obj.bgpV4EvpnEvisSlice
}

func (obj *bgpV4EthernetSegmentBgpV4EvpnEvisIter) Add() BgpV4EvpnEvis {
	newObj := &otg.BgpV4EvpnEvis{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpV4EvpnEvis{obj: newObj}
	newLibObj.setDefault()
	obj.bgpV4EvpnEvisSlice = append(obj.bgpV4EvpnEvisSlice, newLibObj)
	return newLibObj
}

func (obj *bgpV4EthernetSegmentBgpV4EvpnEvisIter) Append(items ...BgpV4EvpnEvis) BgpV4EthernetSegmentBgpV4EvpnEvisIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpV4EvpnEvisSlice = append(obj.bgpV4EvpnEvisSlice, item)
	}
	return obj
}

func (obj *bgpV4EthernetSegmentBgpV4EvpnEvisIter) Set(index int, newObj BgpV4EvpnEvis) BgpV4EthernetSegmentBgpV4EvpnEvisIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpV4EvpnEvisSlice[index] = newObj
	return obj
}
func (obj *bgpV4EthernetSegmentBgpV4EvpnEvisIter) Clear() BgpV4EthernetSegmentBgpV4EvpnEvisIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpV4EvpnEvis{}
		obj.bgpV4EvpnEvisSlice = []BgpV4EvpnEvis{}
	}
	return obj
}
func (obj *bgpV4EthernetSegmentBgpV4EvpnEvisIter) clearHolderSlice() BgpV4EthernetSegmentBgpV4EvpnEvisIter {
	if len(obj.bgpV4EvpnEvisSlice) > 0 {
		obj.bgpV4EvpnEvisSlice = []BgpV4EvpnEvis{}
	}
	return obj
}
func (obj *bgpV4EthernetSegmentBgpV4EvpnEvisIter) appendHolderSlice(item BgpV4EvpnEvis) BgpV4EthernetSegmentBgpV4EvpnEvisIter {
	obj.bgpV4EvpnEvisSlice = append(obj.bgpV4EvpnEvisSlice, item)
	return obj
}

// 10-octet Ethernet Segment Identifier (ESI) Example - For multi-home scenario nonZero ESI is '10000000000000000000' .
// Esi returns a string
func (obj *bgpV4EthernetSegment) Esi() string {

	return *obj.obj.Esi

}

// 10-octet Ethernet Segment Identifier (ESI) Example - For multi-home scenario nonZero ESI is '10000000000000000000' .
// Esi returns a string
func (obj *bgpV4EthernetSegment) HasEsi() bool {
	return obj.obj.Esi != nil
}

// 10-octet Ethernet Segment Identifier (ESI) Example - For multi-home scenario nonZero ESI is '10000000000000000000' .
// SetEsi sets the string value in the BgpV4EthernetSegment object
func (obj *bgpV4EthernetSegment) SetEsi(value string) BgpV4EthernetSegment {

	obj.obj.Esi = &value
	return obj
}

type BgpV4EthernetSegmentActiveModeEnum string

// Enum of ActiveMode on BgpV4EthernetSegment
var BgpV4EthernetSegmentActiveMode = struct {
	SINGLE_ACTIVE BgpV4EthernetSegmentActiveModeEnum
	ALL_ACTIVE    BgpV4EthernetSegmentActiveModeEnum
}{
	SINGLE_ACTIVE: BgpV4EthernetSegmentActiveModeEnum("single_active"),
	ALL_ACTIVE:    BgpV4EthernetSegmentActiveModeEnum("all_active"),
}

func (obj *bgpV4EthernetSegment) ActiveMode() BgpV4EthernetSegmentActiveModeEnum {
	return BgpV4EthernetSegmentActiveModeEnum(obj.obj.ActiveMode.Enum().String())
}

// Single Active or All Active mode Redundancy mode selection for Multi-home.
// ActiveMode returns a string
func (obj *bgpV4EthernetSegment) HasActiveMode() bool {
	return obj.obj.ActiveMode != nil
}

func (obj *bgpV4EthernetSegment) SetActiveMode(value BgpV4EthernetSegmentActiveModeEnum) BgpV4EthernetSegment {
	intValue, ok := otg.BgpV4EthernetSegment_ActiveMode_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpV4EthernetSegmentActiveModeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpV4EthernetSegment_ActiveMode_Enum(intValue)
	obj.obj.ActiveMode = &enumValue

	return obj
}

// The label value to be advertised as ESI Label in ESI Label Extended Community. This is included in Ethernet Auto-discovery per ES Routes advertised by a router.
// EsiLabel returns a uint32
func (obj *bgpV4EthernetSegment) EsiLabel() uint32 {

	return *obj.obj.EsiLabel

}

// The label value to be advertised as ESI Label in ESI Label Extended Community. This is included in Ethernet Auto-discovery per ES Routes advertised by a router.
// EsiLabel returns a uint32
func (obj *bgpV4EthernetSegment) HasEsiLabel() bool {
	return obj.obj.EsiLabel != nil
}

// The label value to be advertised as ESI Label in ESI Label Extended Community. This is included in Ethernet Auto-discovery per ES Routes advertised by a router.
// SetEsiLabel sets the uint32 value in the BgpV4EthernetSegment object
func (obj *bgpV4EthernetSegment) SetEsiLabel(value uint32) BgpV4EthernetSegment {

	obj.obj.EsiLabel = &value
	return obj
}

// description is TBD
// Advanced returns a BgpRouteAdvanced
func (obj *bgpV4EthernetSegment) Advanced() BgpRouteAdvanced {
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
func (obj *bgpV4EthernetSegment) HasAdvanced() bool {
	return obj.obj.Advanced != nil
}

// description is TBD
// SetAdvanced sets the BgpRouteAdvanced value in the BgpV4EthernetSegment object
func (obj *bgpV4EthernetSegment) SetAdvanced(value BgpRouteAdvanced) BgpV4EthernetSegment {

	obj.advancedHolder = nil
	obj.obj.Advanced = value.msg()

	return obj
}

// Optional community settings.
// Communities returns a []BgpCommunity
func (obj *bgpV4EthernetSegment) Communities() BgpV4EthernetSegmentBgpCommunityIter {
	if len(obj.obj.Communities) == 0 {
		obj.obj.Communities = []*otg.BgpCommunity{}
	}
	if obj.communitiesHolder == nil {
		obj.communitiesHolder = newBgpV4EthernetSegmentBgpCommunityIter(&obj.obj.Communities).setMsg(obj)
	}
	return obj.communitiesHolder
}

type bgpV4EthernetSegmentBgpCommunityIter struct {
	obj               *bgpV4EthernetSegment
	bgpCommunitySlice []BgpCommunity
	fieldPtr          *[]*otg.BgpCommunity
}

func newBgpV4EthernetSegmentBgpCommunityIter(ptr *[]*otg.BgpCommunity) BgpV4EthernetSegmentBgpCommunityIter {
	return &bgpV4EthernetSegmentBgpCommunityIter{fieldPtr: ptr}
}

type BgpV4EthernetSegmentBgpCommunityIter interface {
	setMsg(*bgpV4EthernetSegment) BgpV4EthernetSegmentBgpCommunityIter
	Items() []BgpCommunity
	Add() BgpCommunity
	Append(items ...BgpCommunity) BgpV4EthernetSegmentBgpCommunityIter
	Set(index int, newObj BgpCommunity) BgpV4EthernetSegmentBgpCommunityIter
	Clear() BgpV4EthernetSegmentBgpCommunityIter
	clearHolderSlice() BgpV4EthernetSegmentBgpCommunityIter
	appendHolderSlice(item BgpCommunity) BgpV4EthernetSegmentBgpCommunityIter
}

func (obj *bgpV4EthernetSegmentBgpCommunityIter) setMsg(msg *bgpV4EthernetSegment) BgpV4EthernetSegmentBgpCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV4EthernetSegmentBgpCommunityIter) Items() []BgpCommunity {
	return obj.bgpCommunitySlice
}

func (obj *bgpV4EthernetSegmentBgpCommunityIter) Add() BgpCommunity {
	newObj := &otg.BgpCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpV4EthernetSegmentBgpCommunityIter) Append(items ...BgpCommunity) BgpV4EthernetSegmentBgpCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, item)
	}
	return obj
}

func (obj *bgpV4EthernetSegmentBgpCommunityIter) Set(index int, newObj BgpCommunity) BgpV4EthernetSegmentBgpCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpV4EthernetSegmentBgpCommunityIter) Clear() BgpV4EthernetSegmentBgpCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpCommunity{}
		obj.bgpCommunitySlice = []BgpCommunity{}
	}
	return obj
}
func (obj *bgpV4EthernetSegmentBgpCommunityIter) clearHolderSlice() BgpV4EthernetSegmentBgpCommunityIter {
	if len(obj.bgpCommunitySlice) > 0 {
		obj.bgpCommunitySlice = []BgpCommunity{}
	}
	return obj
}
func (obj *bgpV4EthernetSegmentBgpCommunityIter) appendHolderSlice(item BgpCommunity) BgpV4EthernetSegmentBgpCommunityIter {
	obj.bgpCommunitySlice = append(obj.bgpCommunitySlice, item)
	return obj
}

// Optional Extended Community settings. The Extended Communities Attribute is a transitive optional BGP attribute, with the Type Code 16. Community and Extended Communities  attributes are utilized to trigger routing decisions, such as acceptance, rejection,  preference, or redistribution. An extended community is an 8-Bytes value. It is divided into two main parts. The first 2 Bytes of the community encode a type and sub-type fields and the last 6 Bytes carry a unique set of data in a format defined by the type and sub-type field. Extended communities provide a larger  range for grouping or categorizing communities. When type is administrator_as_2octet or administrator_as_4octet, the valid sub types are route target and origin. The valid value for  administrator_as_2octet and administrator_as_4octet type is either two byte AS followed by four byte local administrator id or four byte AS followed by two  byte local administrator id.  When type is administrator_ipv4_address the valid sub types are route target and origin. The valid value for  administrator_ipv4_address is a four byte IPv4 address followed by a two byte local administrator id.  When type is opaque, valid sub types are color and encapsulation. When sub type is color, first two bytes of the value field contain flags and last four bytes  contains the value of the color. When sub type is encapsulation the first four bytes of value field are reserved and last two bytes carries the tunnel type from  IANA's "ETHER TYPES" registry e.g IPv4 (protocol type = 0x0800), IPv6 (protocol type = 0x86dd), and MPLS (protocol type = 0x8847). When type is administrator_as_2octet_link_bandwidth the valid sub type is extended_bandwidth. The first two bytes of the value field contains the AS number and the last four bytes contains the bandwidth in IEEE floating point format.  When type is evpn the valid subtype is mac_address. In the value field the low-order bit of the first byte(Flags) is defined as the "Sticky/static" flag and may be set to 1, indicating the MAC address is static and cannot move. The second byte is reserved and the  last four bytes contain the sequence number which is used to ensure that PEs retain the correct MAC/IP Advertisement route when multiple updates  occur for the same MAC address.
// ExtCommunities returns a []BgpExtCommunity
func (obj *bgpV4EthernetSegment) ExtCommunities() BgpV4EthernetSegmentBgpExtCommunityIter {
	if len(obj.obj.ExtCommunities) == 0 {
		obj.obj.ExtCommunities = []*otg.BgpExtCommunity{}
	}
	if obj.extCommunitiesHolder == nil {
		obj.extCommunitiesHolder = newBgpV4EthernetSegmentBgpExtCommunityIter(&obj.obj.ExtCommunities).setMsg(obj)
	}
	return obj.extCommunitiesHolder
}

type bgpV4EthernetSegmentBgpExtCommunityIter struct {
	obj                  *bgpV4EthernetSegment
	bgpExtCommunitySlice []BgpExtCommunity
	fieldPtr             *[]*otg.BgpExtCommunity
}

func newBgpV4EthernetSegmentBgpExtCommunityIter(ptr *[]*otg.BgpExtCommunity) BgpV4EthernetSegmentBgpExtCommunityIter {
	return &bgpV4EthernetSegmentBgpExtCommunityIter{fieldPtr: ptr}
}

type BgpV4EthernetSegmentBgpExtCommunityIter interface {
	setMsg(*bgpV4EthernetSegment) BgpV4EthernetSegmentBgpExtCommunityIter
	Items() []BgpExtCommunity
	Add() BgpExtCommunity
	Append(items ...BgpExtCommunity) BgpV4EthernetSegmentBgpExtCommunityIter
	Set(index int, newObj BgpExtCommunity) BgpV4EthernetSegmentBgpExtCommunityIter
	Clear() BgpV4EthernetSegmentBgpExtCommunityIter
	clearHolderSlice() BgpV4EthernetSegmentBgpExtCommunityIter
	appendHolderSlice(item BgpExtCommunity) BgpV4EthernetSegmentBgpExtCommunityIter
}

func (obj *bgpV4EthernetSegmentBgpExtCommunityIter) setMsg(msg *bgpV4EthernetSegment) BgpV4EthernetSegmentBgpExtCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpExtCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpV4EthernetSegmentBgpExtCommunityIter) Items() []BgpExtCommunity {
	return obj.bgpExtCommunitySlice
}

func (obj *bgpV4EthernetSegmentBgpExtCommunityIter) Add() BgpExtCommunity {
	newObj := &otg.BgpExtCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpExtCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.bgpExtCommunitySlice = append(obj.bgpExtCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpV4EthernetSegmentBgpExtCommunityIter) Append(items ...BgpExtCommunity) BgpV4EthernetSegmentBgpExtCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpExtCommunitySlice = append(obj.bgpExtCommunitySlice, item)
	}
	return obj
}

func (obj *bgpV4EthernetSegmentBgpExtCommunityIter) Set(index int, newObj BgpExtCommunity) BgpV4EthernetSegmentBgpExtCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpExtCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpV4EthernetSegmentBgpExtCommunityIter) Clear() BgpV4EthernetSegmentBgpExtCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpExtCommunity{}
		obj.bgpExtCommunitySlice = []BgpExtCommunity{}
	}
	return obj
}
func (obj *bgpV4EthernetSegmentBgpExtCommunityIter) clearHolderSlice() BgpV4EthernetSegmentBgpExtCommunityIter {
	if len(obj.bgpExtCommunitySlice) > 0 {
		obj.bgpExtCommunitySlice = []BgpExtCommunity{}
	}
	return obj
}
func (obj *bgpV4EthernetSegmentBgpExtCommunityIter) appendHolderSlice(item BgpExtCommunity) BgpV4EthernetSegmentBgpExtCommunityIter {
	obj.bgpExtCommunitySlice = append(obj.bgpExtCommunitySlice, item)
	return obj
}

// Optional AS PATH settings.
// AsPath returns a BgpAsPath
func (obj *bgpV4EthernetSegment) AsPath() BgpAsPath {
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
func (obj *bgpV4EthernetSegment) HasAsPath() bool {
	return obj.obj.AsPath != nil
}

// Optional AS PATH settings.
// SetAsPath sets the BgpAsPath value in the BgpV4EthernetSegment object
func (obj *bgpV4EthernetSegment) SetAsPath(value BgpAsPath) BgpV4EthernetSegment {

	obj.asPathHolder = nil
	obj.obj.AsPath = value.msg()

	return obj
}

func (obj *bgpV4EthernetSegment) validateObj(vObj *validation, set_default bool) {
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
				obj.Evis().appendHolderSlice(&bgpV4EvpnEvis{obj: item})
			}
		}
		for _, item := range obj.Evis().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.Esi != nil {

		err := obj.validateHex(obj.Esi())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpV4EthernetSegment.Esi"))
		}

	}

	if obj.obj.EsiLabel != nil {

		if *obj.obj.EsiLabel > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpV4EthernetSegment.EsiLabel <= 16777215 but Got %d", *obj.obj.EsiLabel))
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

func (obj *bgpV4EthernetSegment) setDefault() {
	if obj.obj.Esi == nil {
		obj.SetEsi("00000000000000000000")
	}
	if obj.obj.ActiveMode == nil {
		obj.SetActiveMode(BgpV4EthernetSegmentActiveMode.ALL_ACTIVE)

	}
	if obj.obj.EsiLabel == nil {
		obj.SetEsiLabel(0)
	}

}
