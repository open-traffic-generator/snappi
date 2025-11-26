package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BmpPrefixIpv4UnicastState *****
type bmpPrefixIpv4UnicastState struct {
	validation
	obj                       *otg.BmpPrefixIpv4UnicastState
	marshaller                marshalBmpPrefixIpv4UnicastState
	unMarshaller              unMarshalBmpPrefixIpv4UnicastState
	communitiesHolder         BmpPrefixIpv4UnicastStateResultBgpCommunityIter
	extendedCommunitiesHolder BmpPrefixIpv4UnicastStateResultExtendedCommunityIter
	asPathHolder              ResultBgpAsPath
}

func NewBmpPrefixIpv4UnicastState() BmpPrefixIpv4UnicastState {
	obj := bmpPrefixIpv4UnicastState{obj: &otg.BmpPrefixIpv4UnicastState{}}
	obj.setDefault()
	return &obj
}

func (obj *bmpPrefixIpv4UnicastState) msg() *otg.BmpPrefixIpv4UnicastState {
	return obj.obj
}

func (obj *bmpPrefixIpv4UnicastState) setMsg(msg *otg.BmpPrefixIpv4UnicastState) BmpPrefixIpv4UnicastState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbmpPrefixIpv4UnicastState struct {
	obj *bmpPrefixIpv4UnicastState
}

type marshalBmpPrefixIpv4UnicastState interface {
	// ToProto marshals BmpPrefixIpv4UnicastState to protobuf object *otg.BmpPrefixIpv4UnicastState
	ToProto() (*otg.BmpPrefixIpv4UnicastState, error)
	// ToPbText marshals BmpPrefixIpv4UnicastState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BmpPrefixIpv4UnicastState to YAML text
	ToYaml() (string, error)
	// ToJson marshals BmpPrefixIpv4UnicastState to JSON text
	ToJson() (string, error)
}

type unMarshalbmpPrefixIpv4UnicastState struct {
	obj *bmpPrefixIpv4UnicastState
}

type unMarshalBmpPrefixIpv4UnicastState interface {
	// FromProto unmarshals BmpPrefixIpv4UnicastState from protobuf object *otg.BmpPrefixIpv4UnicastState
	FromProto(msg *otg.BmpPrefixIpv4UnicastState) (BmpPrefixIpv4UnicastState, error)
	// FromPbText unmarshals BmpPrefixIpv4UnicastState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BmpPrefixIpv4UnicastState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BmpPrefixIpv4UnicastState from JSON text
	FromJson(value string) error
}

func (obj *bmpPrefixIpv4UnicastState) Marshal() marshalBmpPrefixIpv4UnicastState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbmpPrefixIpv4UnicastState{obj: obj}
	}
	return obj.marshaller
}

func (obj *bmpPrefixIpv4UnicastState) Unmarshal() unMarshalBmpPrefixIpv4UnicastState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbmpPrefixIpv4UnicastState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbmpPrefixIpv4UnicastState) ToProto() (*otg.BmpPrefixIpv4UnicastState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbmpPrefixIpv4UnicastState) FromProto(msg *otg.BmpPrefixIpv4UnicastState) (BmpPrefixIpv4UnicastState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbmpPrefixIpv4UnicastState) ToPbText() (string, error) {
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

func (m *unMarshalbmpPrefixIpv4UnicastState) FromPbText(value string) error {
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

func (m *marshalbmpPrefixIpv4UnicastState) ToYaml() (string, error) {
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

func (m *unMarshalbmpPrefixIpv4UnicastState) FromYaml(value string) error {
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

func (m *marshalbmpPrefixIpv4UnicastState) ToJson() (string, error) {
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

func (m *unMarshalbmpPrefixIpv4UnicastState) FromJson(value string) error {
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

func (obj *bmpPrefixIpv4UnicastState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bmpPrefixIpv4UnicastState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bmpPrefixIpv4UnicastState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bmpPrefixIpv4UnicastState) Clone() (BmpPrefixIpv4UnicastState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBmpPrefixIpv4UnicastState()
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

func (obj *bmpPrefixIpv4UnicastState) setNil() {
	obj.communitiesHolder = nil
	obj.extendedCommunitiesHolder = nil
	obj.asPathHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BmpPrefixIpv4UnicastState is iPv4 unicast prefix.
type BmpPrefixIpv4UnicastState interface {
	Validation
	// msg marshals BmpPrefixIpv4UnicastState to protobuf object *otg.BmpPrefixIpv4UnicastState
	// and doesn't set defaults
	msg() *otg.BmpPrefixIpv4UnicastState
	// setMsg unmarshals BmpPrefixIpv4UnicastState from protobuf object *otg.BmpPrefixIpv4UnicastState
	// and doesn't set defaults
	setMsg(*otg.BmpPrefixIpv4UnicastState) BmpPrefixIpv4UnicastState
	// provides marshal interface
	Marshal() marshalBmpPrefixIpv4UnicastState
	// provides unmarshal interface
	Unmarshal() unMarshalBmpPrefixIpv4UnicastState
	// validate validates BmpPrefixIpv4UnicastState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BmpPrefixIpv4UnicastState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4Address returns string, set in BmpPrefixIpv4UnicastState.
	Ipv4Address() string
	// SetIpv4Address assigns string provided by user to BmpPrefixIpv4UnicastState
	SetIpv4Address(value string) BmpPrefixIpv4UnicastState
	// HasIpv4Address checks if Ipv4Address has been set in BmpPrefixIpv4UnicastState
	HasIpv4Address() bool
	// PrefixLength returns uint32, set in BmpPrefixIpv4UnicastState.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to BmpPrefixIpv4UnicastState
	SetPrefixLength(value uint32) BmpPrefixIpv4UnicastState
	// HasPrefixLength checks if PrefixLength has been set in BmpPrefixIpv4UnicastState
	HasPrefixLength() bool
	// Origin returns BmpPrefixIpv4UnicastStateOriginEnum, set in BmpPrefixIpv4UnicastState
	Origin() BmpPrefixIpv4UnicastStateOriginEnum
	// SetOrigin assigns BmpPrefixIpv4UnicastStateOriginEnum provided by user to BmpPrefixIpv4UnicastState
	SetOrigin(value BmpPrefixIpv4UnicastStateOriginEnum) BmpPrefixIpv4UnicastState
	// HasOrigin checks if Origin has been set in BmpPrefixIpv4UnicastState
	HasOrigin() bool
	// PathId returns uint32, set in BmpPrefixIpv4UnicastState.
	PathId() uint32
	// SetPathId assigns uint32 provided by user to BmpPrefixIpv4UnicastState
	SetPathId(value uint32) BmpPrefixIpv4UnicastState
	// HasPathId checks if PathId has been set in BmpPrefixIpv4UnicastState
	HasPathId() bool
	// RouteState returns BmpPrefixIpv4UnicastStateRouteStateEnum, set in BmpPrefixIpv4UnicastState
	RouteState() BmpPrefixIpv4UnicastStateRouteStateEnum
	// SetRouteState assigns BmpPrefixIpv4UnicastStateRouteStateEnum provided by user to BmpPrefixIpv4UnicastState
	SetRouteState(value BmpPrefixIpv4UnicastStateRouteStateEnum) BmpPrefixIpv4UnicastState
	// HasRouteState checks if RouteState has been set in BmpPrefixIpv4UnicastState
	HasRouteState() bool
	// NextHopType returns BmpPrefixIpv4UnicastStateNextHopTypeEnum, set in BmpPrefixIpv4UnicastState
	NextHopType() BmpPrefixIpv4UnicastStateNextHopTypeEnum
	// SetNextHopType assigns BmpPrefixIpv4UnicastStateNextHopTypeEnum provided by user to BmpPrefixIpv4UnicastState
	SetNextHopType(value BmpPrefixIpv4UnicastStateNextHopTypeEnum) BmpPrefixIpv4UnicastState
	// HasNextHopType checks if NextHopType has been set in BmpPrefixIpv4UnicastState
	HasNextHopType() bool
	// Ipv4NextHop returns string, set in BmpPrefixIpv4UnicastState.
	Ipv4NextHop() string
	// SetIpv4NextHop assigns string provided by user to BmpPrefixIpv4UnicastState
	SetIpv4NextHop(value string) BmpPrefixIpv4UnicastState
	// HasIpv4NextHop checks if Ipv4NextHop has been set in BmpPrefixIpv4UnicastState
	HasIpv4NextHop() bool
	// Ipv6NextHop returns string, set in BmpPrefixIpv4UnicastState.
	Ipv6NextHop() string
	// SetIpv6NextHop assigns string provided by user to BmpPrefixIpv4UnicastState
	SetIpv6NextHop(value string) BmpPrefixIpv4UnicastState
	// HasIpv6NextHop checks if Ipv6NextHop has been set in BmpPrefixIpv4UnicastState
	HasIpv6NextHop() bool
	// Communities returns BmpPrefixIpv4UnicastStateResultBgpCommunityIterIter, set in BmpPrefixIpv4UnicastState
	Communities() BmpPrefixIpv4UnicastStateResultBgpCommunityIter
	// ExtendedCommunities returns BmpPrefixIpv4UnicastStateResultExtendedCommunityIterIter, set in BmpPrefixIpv4UnicastState
	ExtendedCommunities() BmpPrefixIpv4UnicastStateResultExtendedCommunityIter
	// AsPath returns ResultBgpAsPath, set in BmpPrefixIpv4UnicastState.
	// ResultBgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed.
	AsPath() ResultBgpAsPath
	// SetAsPath assigns ResultBgpAsPath provided by user to BmpPrefixIpv4UnicastState.
	// ResultBgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed.
	SetAsPath(value ResultBgpAsPath) BmpPrefixIpv4UnicastState
	// HasAsPath checks if AsPath has been set in BmpPrefixIpv4UnicastState
	HasAsPath() bool
	// LocalPreference returns uint32, set in BmpPrefixIpv4UnicastState.
	LocalPreference() uint32
	// SetLocalPreference assigns uint32 provided by user to BmpPrefixIpv4UnicastState
	SetLocalPreference(value uint32) BmpPrefixIpv4UnicastState
	// HasLocalPreference checks if LocalPreference has been set in BmpPrefixIpv4UnicastState
	HasLocalPreference() bool
	// MultiExitDiscriminator returns uint32, set in BmpPrefixIpv4UnicastState.
	MultiExitDiscriminator() uint32
	// SetMultiExitDiscriminator assigns uint32 provided by user to BmpPrefixIpv4UnicastState
	SetMultiExitDiscriminator(value uint32) BmpPrefixIpv4UnicastState
	// HasMultiExitDiscriminator checks if MultiExitDiscriminator has been set in BmpPrefixIpv4UnicastState
	HasMultiExitDiscriminator() bool
	setNil()
}

// An IPv4 unicast address.
// Ipv4Address returns a string
func (obj *bmpPrefixIpv4UnicastState) Ipv4Address() string {

	return *obj.obj.Ipv4Address

}

// An IPv4 unicast address.
// Ipv4Address returns a string
func (obj *bmpPrefixIpv4UnicastState) HasIpv4Address() bool {
	return obj.obj.Ipv4Address != nil
}

// An IPv4 unicast address.
// SetIpv4Address sets the string value in the BmpPrefixIpv4UnicastState object
func (obj *bmpPrefixIpv4UnicastState) SetIpv4Address(value string) BmpPrefixIpv4UnicastState {

	obj.obj.Ipv4Address = &value
	return obj
}

// The length of the IPv4 prefix.
// PrefixLength returns a uint32
func (obj *bmpPrefixIpv4UnicastState) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// The length of the IPv4 prefix.
// PrefixLength returns a uint32
func (obj *bmpPrefixIpv4UnicastState) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// The length of the IPv4 prefix.
// SetPrefixLength sets the uint32 value in the BmpPrefixIpv4UnicastState object
func (obj *bmpPrefixIpv4UnicastState) SetPrefixLength(value uint32) BmpPrefixIpv4UnicastState {

	obj.obj.PrefixLength = &value
	return obj
}

type BmpPrefixIpv4UnicastStateOriginEnum string

// Enum of Origin on BmpPrefixIpv4UnicastState
var BmpPrefixIpv4UnicastStateOrigin = struct {
	IGP        BmpPrefixIpv4UnicastStateOriginEnum
	EGP        BmpPrefixIpv4UnicastStateOriginEnum
	INCOMPLETE BmpPrefixIpv4UnicastStateOriginEnum
}{
	IGP:        BmpPrefixIpv4UnicastStateOriginEnum("igp"),
	EGP:        BmpPrefixIpv4UnicastStateOriginEnum("egp"),
	INCOMPLETE: BmpPrefixIpv4UnicastStateOriginEnum("incomplete"),
}

func (obj *bmpPrefixIpv4UnicastState) Origin() BmpPrefixIpv4UnicastStateOriginEnum {
	return BmpPrefixIpv4UnicastStateOriginEnum(obj.obj.Origin.Enum().String())
}

// The origin of the prefix.
// Origin returns a string
func (obj *bmpPrefixIpv4UnicastState) HasOrigin() bool {
	return obj.obj.Origin != nil
}

func (obj *bmpPrefixIpv4UnicastState) SetOrigin(value BmpPrefixIpv4UnicastStateOriginEnum) BmpPrefixIpv4UnicastState {
	intValue, ok := otg.BmpPrefixIpv4UnicastState_Origin_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BmpPrefixIpv4UnicastStateOriginEnum", string(value)))
		return obj
	}
	enumValue := otg.BmpPrefixIpv4UnicastState_Origin_Enum(intValue)
	obj.obj.Origin = &enumValue

	return obj
}

// The path id.
// PathId returns a uint32
func (obj *bmpPrefixIpv4UnicastState) PathId() uint32 {

	return *obj.obj.PathId

}

// The path id.
// PathId returns a uint32
func (obj *bmpPrefixIpv4UnicastState) HasPathId() bool {
	return obj.obj.PathId != nil
}

// The path id.
// SetPathId sets the uint32 value in the BmpPrefixIpv4UnicastState object
func (obj *bmpPrefixIpv4UnicastState) SetPathId(value uint32) BmpPrefixIpv4UnicastState {

	obj.obj.PathId = &value
	return obj
}

type BmpPrefixIpv4UnicastStateRouteStateEnum string

// Enum of RouteState on BmpPrefixIpv4UnicastState
var BmpPrefixIpv4UnicastStateRouteState = struct {
	ADVERTISED BmpPrefixIpv4UnicastStateRouteStateEnum
	WITHDRAWN  BmpPrefixIpv4UnicastStateRouteStateEnum
}{
	ADVERTISED: BmpPrefixIpv4UnicastStateRouteStateEnum("advertised"),
	WITHDRAWN:  BmpPrefixIpv4UnicastStateRouteStateEnum("withdrawn"),
}

func (obj *bmpPrefixIpv4UnicastState) RouteState() BmpPrefixIpv4UnicastStateRouteStateEnum {
	return BmpPrefixIpv4UnicastStateRouteStateEnum(obj.obj.RouteState.Enum().String())
}

// The state of the route , either advertised or withdrawn. If route is in withdrawn state, only the prefix, prefix_length, origin and path_id (if applicable) may be present.
// RouteState returns a string
func (obj *bmpPrefixIpv4UnicastState) HasRouteState() bool {
	return obj.obj.RouteState != nil
}

func (obj *bmpPrefixIpv4UnicastState) SetRouteState(value BmpPrefixIpv4UnicastStateRouteStateEnum) BmpPrefixIpv4UnicastState {
	intValue, ok := otg.BmpPrefixIpv4UnicastState_RouteState_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BmpPrefixIpv4UnicastStateRouteStateEnum", string(value)))
		return obj
	}
	enumValue := otg.BmpPrefixIpv4UnicastState_RouteState_Enum(intValue)
	obj.obj.RouteState = &enumValue

	return obj
}

type BmpPrefixIpv4UnicastStateNextHopTypeEnum string

// Enum of NextHopType on BmpPrefixIpv4UnicastState
var BmpPrefixIpv4UnicastStateNextHopType = struct {
	IPV4 BmpPrefixIpv4UnicastStateNextHopTypeEnum
	IPV6 BmpPrefixIpv4UnicastStateNextHopTypeEnum
}{
	IPV4: BmpPrefixIpv4UnicastStateNextHopTypeEnum("ipv4"),
	IPV6: BmpPrefixIpv4UnicastStateNextHopTypeEnum("ipv6"),
}

func (obj *bmpPrefixIpv4UnicastState) NextHopType() BmpPrefixIpv4UnicastStateNextHopTypeEnum {
	return BmpPrefixIpv4UnicastStateNextHopTypeEnum(obj.obj.NextHopType.Enum().String())
}

// The next hop type , ipv4 or ipv6 .  If type is ipv4, then the ipv4_next_hop should be included. If type is ipv6, then the ipv6_next_hop should be included.
// NextHopType returns a string
func (obj *bmpPrefixIpv4UnicastState) HasNextHopType() bool {
	return obj.obj.NextHopType != nil
}

func (obj *bmpPrefixIpv4UnicastState) SetNextHopType(value BmpPrefixIpv4UnicastStateNextHopTypeEnum) BmpPrefixIpv4UnicastState {
	intValue, ok := otg.BmpPrefixIpv4UnicastState_NextHopType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BmpPrefixIpv4UnicastStateNextHopTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.BmpPrefixIpv4UnicastState_NextHopType_Enum(intValue)
	obj.obj.NextHopType = &enumValue

	return obj
}

// The IPv4 address of the egress interface.
// Ipv4NextHop returns a string
func (obj *bmpPrefixIpv4UnicastState) Ipv4NextHop() string {

	return *obj.obj.Ipv4NextHop

}

// The IPv4 address of the egress interface.
// Ipv4NextHop returns a string
func (obj *bmpPrefixIpv4UnicastState) HasIpv4NextHop() bool {
	return obj.obj.Ipv4NextHop != nil
}

// The IPv4 address of the egress interface.
// SetIpv4NextHop sets the string value in the BmpPrefixIpv4UnicastState object
func (obj *bmpPrefixIpv4UnicastState) SetIpv4NextHop(value string) BmpPrefixIpv4UnicastState {

	obj.obj.Ipv4NextHop = &value
	return obj
}

// The IPv6 address of the egress interface.
// Ipv6NextHop returns a string
func (obj *bmpPrefixIpv4UnicastState) Ipv6NextHop() string {

	return *obj.obj.Ipv6NextHop

}

// The IPv6 address of the egress interface.
// Ipv6NextHop returns a string
func (obj *bmpPrefixIpv4UnicastState) HasIpv6NextHop() bool {
	return obj.obj.Ipv6NextHop != nil
}

// The IPv6 address of the egress interface.
// SetIpv6NextHop sets the string value in the BmpPrefixIpv4UnicastState object
func (obj *bmpPrefixIpv4UnicastState) SetIpv6NextHop(value string) BmpPrefixIpv4UnicastState {

	obj.obj.Ipv6NextHop = &value
	return obj
}

// Optional community attributes.
// Communities returns a []ResultBgpCommunity
func (obj *bmpPrefixIpv4UnicastState) Communities() BmpPrefixIpv4UnicastStateResultBgpCommunityIter {
	if len(obj.obj.Communities) == 0 {
		obj.obj.Communities = []*otg.ResultBgpCommunity{}
	}
	if obj.communitiesHolder == nil {
		obj.communitiesHolder = newBmpPrefixIpv4UnicastStateResultBgpCommunityIter(&obj.obj.Communities).setMsg(obj)
	}
	return obj.communitiesHolder
}

type bmpPrefixIpv4UnicastStateResultBgpCommunityIter struct {
	obj                     *bmpPrefixIpv4UnicastState
	resultBgpCommunitySlice []ResultBgpCommunity
	fieldPtr                *[]*otg.ResultBgpCommunity
}

func newBmpPrefixIpv4UnicastStateResultBgpCommunityIter(ptr *[]*otg.ResultBgpCommunity) BmpPrefixIpv4UnicastStateResultBgpCommunityIter {
	return &bmpPrefixIpv4UnicastStateResultBgpCommunityIter{fieldPtr: ptr}
}

type BmpPrefixIpv4UnicastStateResultBgpCommunityIter interface {
	setMsg(*bmpPrefixIpv4UnicastState) BmpPrefixIpv4UnicastStateResultBgpCommunityIter
	Items() []ResultBgpCommunity
	Add() ResultBgpCommunity
	Append(items ...ResultBgpCommunity) BmpPrefixIpv4UnicastStateResultBgpCommunityIter
	Set(index int, newObj ResultBgpCommunity) BmpPrefixIpv4UnicastStateResultBgpCommunityIter
	Clear() BmpPrefixIpv4UnicastStateResultBgpCommunityIter
	clearHolderSlice() BmpPrefixIpv4UnicastStateResultBgpCommunityIter
	appendHolderSlice(item ResultBgpCommunity) BmpPrefixIpv4UnicastStateResultBgpCommunityIter
}

func (obj *bmpPrefixIpv4UnicastStateResultBgpCommunityIter) setMsg(msg *bmpPrefixIpv4UnicastState) BmpPrefixIpv4UnicastStateResultBgpCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&resultBgpCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bmpPrefixIpv4UnicastStateResultBgpCommunityIter) Items() []ResultBgpCommunity {
	return obj.resultBgpCommunitySlice
}

func (obj *bmpPrefixIpv4UnicastStateResultBgpCommunityIter) Add() ResultBgpCommunity {
	newObj := &otg.ResultBgpCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &resultBgpCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bmpPrefixIpv4UnicastStateResultBgpCommunityIter) Append(items ...ResultBgpCommunity) BmpPrefixIpv4UnicastStateResultBgpCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, item)
	}
	return obj
}

func (obj *bmpPrefixIpv4UnicastStateResultBgpCommunityIter) Set(index int, newObj ResultBgpCommunity) BmpPrefixIpv4UnicastStateResultBgpCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.resultBgpCommunitySlice[index] = newObj
	return obj
}
func (obj *bmpPrefixIpv4UnicastStateResultBgpCommunityIter) Clear() BmpPrefixIpv4UnicastStateResultBgpCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ResultBgpCommunity{}
		obj.resultBgpCommunitySlice = []ResultBgpCommunity{}
	}
	return obj
}
func (obj *bmpPrefixIpv4UnicastStateResultBgpCommunityIter) clearHolderSlice() BmpPrefixIpv4UnicastStateResultBgpCommunityIter {
	if len(obj.resultBgpCommunitySlice) > 0 {
		obj.resultBgpCommunitySlice = []ResultBgpCommunity{}
	}
	return obj
}
func (obj *bmpPrefixIpv4UnicastStateResultBgpCommunityIter) appendHolderSlice(item ResultBgpCommunity) BmpPrefixIpv4UnicastStateResultBgpCommunityIter {
	obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, item)
	return obj
}

// Optional received Extended Community attributes. Each received Extended Community attribute is available for retrieval in two forms. Support of the 'raw' format in which all 8 bytes (16 hex characters) is always present and available for use. In addition, if supported by the implementation, the Extended Community attribute may also be retrieved in the  'structured' format which is an optional field.
// ExtendedCommunities returns a []ResultExtendedCommunity
func (obj *bmpPrefixIpv4UnicastState) ExtendedCommunities() BmpPrefixIpv4UnicastStateResultExtendedCommunityIter {
	if len(obj.obj.ExtendedCommunities) == 0 {
		obj.obj.ExtendedCommunities = []*otg.ResultExtendedCommunity{}
	}
	if obj.extendedCommunitiesHolder == nil {
		obj.extendedCommunitiesHolder = newBmpPrefixIpv4UnicastStateResultExtendedCommunityIter(&obj.obj.ExtendedCommunities).setMsg(obj)
	}
	return obj.extendedCommunitiesHolder
}

type bmpPrefixIpv4UnicastStateResultExtendedCommunityIter struct {
	obj                          *bmpPrefixIpv4UnicastState
	resultExtendedCommunitySlice []ResultExtendedCommunity
	fieldPtr                     *[]*otg.ResultExtendedCommunity
}

func newBmpPrefixIpv4UnicastStateResultExtendedCommunityIter(ptr *[]*otg.ResultExtendedCommunity) BmpPrefixIpv4UnicastStateResultExtendedCommunityIter {
	return &bmpPrefixIpv4UnicastStateResultExtendedCommunityIter{fieldPtr: ptr}
}

type BmpPrefixIpv4UnicastStateResultExtendedCommunityIter interface {
	setMsg(*bmpPrefixIpv4UnicastState) BmpPrefixIpv4UnicastStateResultExtendedCommunityIter
	Items() []ResultExtendedCommunity
	Add() ResultExtendedCommunity
	Append(items ...ResultExtendedCommunity) BmpPrefixIpv4UnicastStateResultExtendedCommunityIter
	Set(index int, newObj ResultExtendedCommunity) BmpPrefixIpv4UnicastStateResultExtendedCommunityIter
	Clear() BmpPrefixIpv4UnicastStateResultExtendedCommunityIter
	clearHolderSlice() BmpPrefixIpv4UnicastStateResultExtendedCommunityIter
	appendHolderSlice(item ResultExtendedCommunity) BmpPrefixIpv4UnicastStateResultExtendedCommunityIter
}

func (obj *bmpPrefixIpv4UnicastStateResultExtendedCommunityIter) setMsg(msg *bmpPrefixIpv4UnicastState) BmpPrefixIpv4UnicastStateResultExtendedCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&resultExtendedCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bmpPrefixIpv4UnicastStateResultExtendedCommunityIter) Items() []ResultExtendedCommunity {
	return obj.resultExtendedCommunitySlice
}

func (obj *bmpPrefixIpv4UnicastStateResultExtendedCommunityIter) Add() ResultExtendedCommunity {
	newObj := &otg.ResultExtendedCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &resultExtendedCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bmpPrefixIpv4UnicastStateResultExtendedCommunityIter) Append(items ...ResultExtendedCommunity) BmpPrefixIpv4UnicastStateResultExtendedCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, item)
	}
	return obj
}

func (obj *bmpPrefixIpv4UnicastStateResultExtendedCommunityIter) Set(index int, newObj ResultExtendedCommunity) BmpPrefixIpv4UnicastStateResultExtendedCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.resultExtendedCommunitySlice[index] = newObj
	return obj
}
func (obj *bmpPrefixIpv4UnicastStateResultExtendedCommunityIter) Clear() BmpPrefixIpv4UnicastStateResultExtendedCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ResultExtendedCommunity{}
		obj.resultExtendedCommunitySlice = []ResultExtendedCommunity{}
	}
	return obj
}
func (obj *bmpPrefixIpv4UnicastStateResultExtendedCommunityIter) clearHolderSlice() BmpPrefixIpv4UnicastStateResultExtendedCommunityIter {
	if len(obj.resultExtendedCommunitySlice) > 0 {
		obj.resultExtendedCommunitySlice = []ResultExtendedCommunity{}
	}
	return obj
}
func (obj *bmpPrefixIpv4UnicastStateResultExtendedCommunityIter) appendHolderSlice(item ResultExtendedCommunity) BmpPrefixIpv4UnicastStateResultExtendedCommunityIter {
	obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, item)
	return obj
}

// description is TBD
// AsPath returns a ResultBgpAsPath
func (obj *bmpPrefixIpv4UnicastState) AsPath() ResultBgpAsPath {
	if obj.obj.AsPath == nil {
		obj.obj.AsPath = NewResultBgpAsPath().msg()
	}
	if obj.asPathHolder == nil {
		obj.asPathHolder = &resultBgpAsPath{obj: obj.obj.AsPath}
	}
	return obj.asPathHolder
}

// description is TBD
// AsPath returns a ResultBgpAsPath
func (obj *bmpPrefixIpv4UnicastState) HasAsPath() bool {
	return obj.obj.AsPath != nil
}

// description is TBD
// SetAsPath sets the ResultBgpAsPath value in the BmpPrefixIpv4UnicastState object
func (obj *bmpPrefixIpv4UnicastState) SetAsPath(value ResultBgpAsPath) BmpPrefixIpv4UnicastState {

	obj.asPathHolder = nil
	obj.obj.AsPath = value.msg()

	return obj
}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// LocalPreference returns a uint32
func (obj *bmpPrefixIpv4UnicastState) LocalPreference() uint32 {

	return *obj.obj.LocalPreference

}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// LocalPreference returns a uint32
func (obj *bmpPrefixIpv4UnicastState) HasLocalPreference() bool {
	return obj.obj.LocalPreference != nil
}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// SetLocalPreference sets the uint32 value in the BmpPrefixIpv4UnicastState object
func (obj *bmpPrefixIpv4UnicastState) SetLocalPreference(value uint32) BmpPrefixIpv4UnicastState {

	obj.obj.LocalPreference = &value
	return obj
}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// MultiExitDiscriminator returns a uint32
func (obj *bmpPrefixIpv4UnicastState) MultiExitDiscriminator() uint32 {

	return *obj.obj.MultiExitDiscriminator

}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// MultiExitDiscriminator returns a uint32
func (obj *bmpPrefixIpv4UnicastState) HasMultiExitDiscriminator() bool {
	return obj.obj.MultiExitDiscriminator != nil
}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// SetMultiExitDiscriminator sets the uint32 value in the BmpPrefixIpv4UnicastState object
func (obj *bmpPrefixIpv4UnicastState) SetMultiExitDiscriminator(value uint32) BmpPrefixIpv4UnicastState {

	obj.obj.MultiExitDiscriminator = &value
	return obj
}

func (obj *bmpPrefixIpv4UnicastState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv4Address != nil {

		err := obj.validateIpv4(obj.Ipv4Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BmpPrefixIpv4UnicastState.Ipv4Address"))
		}

	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BmpPrefixIpv4UnicastState.PrefixLength <= 32 but Got %d", *obj.obj.PrefixLength))
		}

	}

	if obj.obj.Ipv4NextHop != nil {

		err := obj.validateIpv4(obj.Ipv4NextHop())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BmpPrefixIpv4UnicastState.Ipv4NextHop"))
		}

	}

	if obj.obj.Ipv6NextHop != nil {

		err := obj.validateIpv6(obj.Ipv6NextHop())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BmpPrefixIpv4UnicastState.Ipv6NextHop"))
		}

	}

	if len(obj.obj.Communities) != 0 {

		if set_default {
			obj.Communities().clearHolderSlice()
			for _, item := range obj.obj.Communities {
				obj.Communities().appendHolderSlice(&resultBgpCommunity{obj: item})
			}
		}
		for _, item := range obj.Communities().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.ExtendedCommunities) != 0 {

		if set_default {
			obj.ExtendedCommunities().clearHolderSlice()
			for _, item := range obj.obj.ExtendedCommunities {
				obj.ExtendedCommunities().appendHolderSlice(&resultExtendedCommunity{obj: item})
			}
		}
		for _, item := range obj.ExtendedCommunities().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.AsPath != nil {

		obj.AsPath().validateObj(vObj, set_default)
	}

}

func (obj *bmpPrefixIpv4UnicastState) setDefault() {

}
