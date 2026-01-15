package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BmpPrefixIpv6UnicastState *****
type bmpPrefixIpv6UnicastState struct {
	validation
	obj                       *otg.BmpPrefixIpv6UnicastState
	marshaller                marshalBmpPrefixIpv6UnicastState
	unMarshaller              unMarshalBmpPrefixIpv6UnicastState
	communitiesHolder         BmpPrefixIpv6UnicastStateResultBgpCommunityIter
	extendedCommunitiesHolder BmpPrefixIpv6UnicastStateResultExtendedCommunityIter
	asPathHolder              ResultBgpAsPath
}

func NewBmpPrefixIpv6UnicastState() BmpPrefixIpv6UnicastState {
	obj := bmpPrefixIpv6UnicastState{obj: &otg.BmpPrefixIpv6UnicastState{}}
	obj.setDefault()
	return &obj
}

func (obj *bmpPrefixIpv6UnicastState) msg() *otg.BmpPrefixIpv6UnicastState {
	return obj.obj
}

func (obj *bmpPrefixIpv6UnicastState) setMsg(msg *otg.BmpPrefixIpv6UnicastState) BmpPrefixIpv6UnicastState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbmpPrefixIpv6UnicastState struct {
	obj *bmpPrefixIpv6UnicastState
}

type marshalBmpPrefixIpv6UnicastState interface {
	// ToProto marshals BmpPrefixIpv6UnicastState to protobuf object *otg.BmpPrefixIpv6UnicastState
	ToProto() (*otg.BmpPrefixIpv6UnicastState, error)
	// ToPbText marshals BmpPrefixIpv6UnicastState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BmpPrefixIpv6UnicastState to YAML text
	ToYaml() (string, error)
	// ToJson marshals BmpPrefixIpv6UnicastState to JSON text
	ToJson() (string, error)
}

type unMarshalbmpPrefixIpv6UnicastState struct {
	obj *bmpPrefixIpv6UnicastState
}

type unMarshalBmpPrefixIpv6UnicastState interface {
	// FromProto unmarshals BmpPrefixIpv6UnicastState from protobuf object *otg.BmpPrefixIpv6UnicastState
	FromProto(msg *otg.BmpPrefixIpv6UnicastState) (BmpPrefixIpv6UnicastState, error)
	// FromPbText unmarshals BmpPrefixIpv6UnicastState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BmpPrefixIpv6UnicastState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BmpPrefixIpv6UnicastState from JSON text
	FromJson(value string) error
}

func (obj *bmpPrefixIpv6UnicastState) Marshal() marshalBmpPrefixIpv6UnicastState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbmpPrefixIpv6UnicastState{obj: obj}
	}
	return obj.marshaller
}

func (obj *bmpPrefixIpv6UnicastState) Unmarshal() unMarshalBmpPrefixIpv6UnicastState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbmpPrefixIpv6UnicastState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbmpPrefixIpv6UnicastState) ToProto() (*otg.BmpPrefixIpv6UnicastState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbmpPrefixIpv6UnicastState) FromProto(msg *otg.BmpPrefixIpv6UnicastState) (BmpPrefixIpv6UnicastState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbmpPrefixIpv6UnicastState) ToPbText() (string, error) {
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

func (m *unMarshalbmpPrefixIpv6UnicastState) FromPbText(value string) error {
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

func (m *marshalbmpPrefixIpv6UnicastState) ToYaml() (string, error) {
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

func (m *unMarshalbmpPrefixIpv6UnicastState) FromYaml(value string) error {
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

func (m *marshalbmpPrefixIpv6UnicastState) ToJson() (string, error) {
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

func (m *unMarshalbmpPrefixIpv6UnicastState) FromJson(value string) error {
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

func (obj *bmpPrefixIpv6UnicastState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bmpPrefixIpv6UnicastState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bmpPrefixIpv6UnicastState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bmpPrefixIpv6UnicastState) Clone() (BmpPrefixIpv6UnicastState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBmpPrefixIpv6UnicastState()
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

func (obj *bmpPrefixIpv6UnicastState) setNil() {
	obj.communitiesHolder = nil
	obj.extendedCommunitiesHolder = nil
	obj.asPathHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BmpPrefixIpv6UnicastState is iPv6 unicast prefix.
type BmpPrefixIpv6UnicastState interface {
	Validation
	// msg marshals BmpPrefixIpv6UnicastState to protobuf object *otg.BmpPrefixIpv6UnicastState
	// and doesn't set defaults
	msg() *otg.BmpPrefixIpv6UnicastState
	// setMsg unmarshals BmpPrefixIpv6UnicastState from protobuf object *otg.BmpPrefixIpv6UnicastState
	// and doesn't set defaults
	setMsg(*otg.BmpPrefixIpv6UnicastState) BmpPrefixIpv6UnicastState
	// provides marshal interface
	Marshal() marshalBmpPrefixIpv6UnicastState
	// provides unmarshal interface
	Unmarshal() unMarshalBmpPrefixIpv6UnicastState
	// validate validates BmpPrefixIpv6UnicastState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BmpPrefixIpv6UnicastState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv6Address returns string, set in BmpPrefixIpv6UnicastState.
	Ipv6Address() string
	// SetIpv6Address assigns string provided by user to BmpPrefixIpv6UnicastState
	SetIpv6Address(value string) BmpPrefixIpv6UnicastState
	// HasIpv6Address checks if Ipv6Address has been set in BmpPrefixIpv6UnicastState
	HasIpv6Address() bool
	// PrefixLength returns uint32, set in BmpPrefixIpv6UnicastState.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to BmpPrefixIpv6UnicastState
	SetPrefixLength(value uint32) BmpPrefixIpv6UnicastState
	// HasPrefixLength checks if PrefixLength has been set in BmpPrefixIpv6UnicastState
	HasPrefixLength() bool
	// Origin returns BmpPrefixIpv6UnicastStateOriginEnum, set in BmpPrefixIpv6UnicastState
	Origin() BmpPrefixIpv6UnicastStateOriginEnum
	// SetOrigin assigns BmpPrefixIpv6UnicastStateOriginEnum provided by user to BmpPrefixIpv6UnicastState
	SetOrigin(value BmpPrefixIpv6UnicastStateOriginEnum) BmpPrefixIpv6UnicastState
	// HasOrigin checks if Origin has been set in BmpPrefixIpv6UnicastState
	HasOrigin() bool
	// PathId returns uint32, set in BmpPrefixIpv6UnicastState.
	PathId() uint32
	// SetPathId assigns uint32 provided by user to BmpPrefixIpv6UnicastState
	SetPathId(value uint32) BmpPrefixIpv6UnicastState
	// HasPathId checks if PathId has been set in BmpPrefixIpv6UnicastState
	HasPathId() bool
	// RouteState returns BmpPrefixIpv6UnicastStateRouteStateEnum, set in BmpPrefixIpv6UnicastState
	RouteState() BmpPrefixIpv6UnicastStateRouteStateEnum
	// SetRouteState assigns BmpPrefixIpv6UnicastStateRouteStateEnum provided by user to BmpPrefixIpv6UnicastState
	SetRouteState(value BmpPrefixIpv6UnicastStateRouteStateEnum) BmpPrefixIpv6UnicastState
	// HasRouteState checks if RouteState has been set in BmpPrefixIpv6UnicastState
	HasRouteState() bool
	// NextHopType returns BmpPrefixIpv6UnicastStateNextHopTypeEnum, set in BmpPrefixIpv6UnicastState
	NextHopType() BmpPrefixIpv6UnicastStateNextHopTypeEnum
	// SetNextHopType assigns BmpPrefixIpv6UnicastStateNextHopTypeEnum provided by user to BmpPrefixIpv6UnicastState
	SetNextHopType(value BmpPrefixIpv6UnicastStateNextHopTypeEnum) BmpPrefixIpv6UnicastState
	// HasNextHopType checks if NextHopType has been set in BmpPrefixIpv6UnicastState
	HasNextHopType() bool
	// Ipv4NextHop returns string, set in BmpPrefixIpv6UnicastState.
	Ipv4NextHop() string
	// SetIpv4NextHop assigns string provided by user to BmpPrefixIpv6UnicastState
	SetIpv4NextHop(value string) BmpPrefixIpv6UnicastState
	// HasIpv4NextHop checks if Ipv4NextHop has been set in BmpPrefixIpv6UnicastState
	HasIpv4NextHop() bool
	// Ipv6NextHop returns string, set in BmpPrefixIpv6UnicastState.
	Ipv6NextHop() string
	// SetIpv6NextHop assigns string provided by user to BmpPrefixIpv6UnicastState
	SetIpv6NextHop(value string) BmpPrefixIpv6UnicastState
	// HasIpv6NextHop checks if Ipv6NextHop has been set in BmpPrefixIpv6UnicastState
	HasIpv6NextHop() bool
	// Communities returns BmpPrefixIpv6UnicastStateResultBgpCommunityIterIter, set in BmpPrefixIpv6UnicastState
	Communities() BmpPrefixIpv6UnicastStateResultBgpCommunityIter
	// ExtendedCommunities returns BmpPrefixIpv6UnicastStateResultExtendedCommunityIterIter, set in BmpPrefixIpv6UnicastState
	ExtendedCommunities() BmpPrefixIpv6UnicastStateResultExtendedCommunityIter
	// AsPath returns ResultBgpAsPath, set in BmpPrefixIpv6UnicastState.
	// ResultBgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed.
	AsPath() ResultBgpAsPath
	// SetAsPath assigns ResultBgpAsPath provided by user to BmpPrefixIpv6UnicastState.
	// ResultBgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed.
	SetAsPath(value ResultBgpAsPath) BmpPrefixIpv6UnicastState
	// HasAsPath checks if AsPath has been set in BmpPrefixIpv6UnicastState
	HasAsPath() bool
	// LocalPreference returns uint32, set in BmpPrefixIpv6UnicastState.
	LocalPreference() uint32
	// SetLocalPreference assigns uint32 provided by user to BmpPrefixIpv6UnicastState
	SetLocalPreference(value uint32) BmpPrefixIpv6UnicastState
	// HasLocalPreference checks if LocalPreference has been set in BmpPrefixIpv6UnicastState
	HasLocalPreference() bool
	// MultiExitDiscriminator returns uint32, set in BmpPrefixIpv6UnicastState.
	MultiExitDiscriminator() uint32
	// SetMultiExitDiscriminator assigns uint32 provided by user to BmpPrefixIpv6UnicastState
	SetMultiExitDiscriminator(value uint32) BmpPrefixIpv6UnicastState
	// HasMultiExitDiscriminator checks if MultiExitDiscriminator has been set in BmpPrefixIpv6UnicastState
	HasMultiExitDiscriminator() bool
	setNil()
}

// An IPv6 unicast address.
// Ipv6Address returns a string
func (obj *bmpPrefixIpv6UnicastState) Ipv6Address() string {

	return *obj.obj.Ipv6Address

}

// An IPv6 unicast address.
// Ipv6Address returns a string
func (obj *bmpPrefixIpv6UnicastState) HasIpv6Address() bool {
	return obj.obj.Ipv6Address != nil
}

// An IPv6 unicast address.
// SetIpv6Address sets the string value in the BmpPrefixIpv6UnicastState object
func (obj *bmpPrefixIpv6UnicastState) SetIpv6Address(value string) BmpPrefixIpv6UnicastState {

	obj.obj.Ipv6Address = &value
	return obj
}

// The length of the IPv6 prefix.
// PrefixLength returns a uint32
func (obj *bmpPrefixIpv6UnicastState) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// The length of the IPv6 prefix.
// PrefixLength returns a uint32
func (obj *bmpPrefixIpv6UnicastState) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// The length of the IPv6 prefix.
// SetPrefixLength sets the uint32 value in the BmpPrefixIpv6UnicastState object
func (obj *bmpPrefixIpv6UnicastState) SetPrefixLength(value uint32) BmpPrefixIpv6UnicastState {

	obj.obj.PrefixLength = &value
	return obj
}

type BmpPrefixIpv6UnicastStateOriginEnum string

// Enum of Origin on BmpPrefixIpv6UnicastState
var BmpPrefixIpv6UnicastStateOrigin = struct {
	IGP        BmpPrefixIpv6UnicastStateOriginEnum
	EGP        BmpPrefixIpv6UnicastStateOriginEnum
	INCOMPLETE BmpPrefixIpv6UnicastStateOriginEnum
}{
	IGP:        BmpPrefixIpv6UnicastStateOriginEnum("igp"),
	EGP:        BmpPrefixIpv6UnicastStateOriginEnum("egp"),
	INCOMPLETE: BmpPrefixIpv6UnicastStateOriginEnum("incomplete"),
}

func (obj *bmpPrefixIpv6UnicastState) Origin() BmpPrefixIpv6UnicastStateOriginEnum {
	return BmpPrefixIpv6UnicastStateOriginEnum(obj.obj.Origin.Enum().String())
}

// The origin of the prefix.
// Origin returns a string
func (obj *bmpPrefixIpv6UnicastState) HasOrigin() bool {
	return obj.obj.Origin != nil
}

func (obj *bmpPrefixIpv6UnicastState) SetOrigin(value BmpPrefixIpv6UnicastStateOriginEnum) BmpPrefixIpv6UnicastState {
	intValue, ok := otg.BmpPrefixIpv6UnicastState_Origin_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BmpPrefixIpv6UnicastStateOriginEnum", string(value)))
		return obj
	}
	enumValue := otg.BmpPrefixIpv6UnicastState_Origin_Enum(intValue)
	obj.obj.Origin = &enumValue

	return obj
}

// The path id.
// PathId returns a uint32
func (obj *bmpPrefixIpv6UnicastState) PathId() uint32 {

	return *obj.obj.PathId

}

// The path id.
// PathId returns a uint32
func (obj *bmpPrefixIpv6UnicastState) HasPathId() bool {
	return obj.obj.PathId != nil
}

// The path id.
// SetPathId sets the uint32 value in the BmpPrefixIpv6UnicastState object
func (obj *bmpPrefixIpv6UnicastState) SetPathId(value uint32) BmpPrefixIpv6UnicastState {

	obj.obj.PathId = &value
	return obj
}

type BmpPrefixIpv6UnicastStateRouteStateEnum string

// Enum of RouteState on BmpPrefixIpv6UnicastState
var BmpPrefixIpv6UnicastStateRouteState = struct {
	ADVERTISED BmpPrefixIpv6UnicastStateRouteStateEnum
	WITHDRAWN  BmpPrefixIpv6UnicastStateRouteStateEnum
}{
	ADVERTISED: BmpPrefixIpv6UnicastStateRouteStateEnum("advertised"),
	WITHDRAWN:  BmpPrefixIpv6UnicastStateRouteStateEnum("withdrawn"),
}

func (obj *bmpPrefixIpv6UnicastState) RouteState() BmpPrefixIpv6UnicastStateRouteStateEnum {
	return BmpPrefixIpv6UnicastStateRouteStateEnum(obj.obj.RouteState.Enum().String())
}

// The state of the route , either advertised or withdrawn. If route is in withdrawn state, only the prefix, prefix_length, origin and path_id (if applicable) may be present.
// RouteState returns a string
func (obj *bmpPrefixIpv6UnicastState) HasRouteState() bool {
	return obj.obj.RouteState != nil
}

func (obj *bmpPrefixIpv6UnicastState) SetRouteState(value BmpPrefixIpv6UnicastStateRouteStateEnum) BmpPrefixIpv6UnicastState {
	intValue, ok := otg.BmpPrefixIpv6UnicastState_RouteState_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BmpPrefixIpv6UnicastStateRouteStateEnum", string(value)))
		return obj
	}
	enumValue := otg.BmpPrefixIpv6UnicastState_RouteState_Enum(intValue)
	obj.obj.RouteState = &enumValue

	return obj
}

type BmpPrefixIpv6UnicastStateNextHopTypeEnum string

// Enum of NextHopType on BmpPrefixIpv6UnicastState
var BmpPrefixIpv6UnicastStateNextHopType = struct {
	IPV4 BmpPrefixIpv6UnicastStateNextHopTypeEnum
	IPV6 BmpPrefixIpv6UnicastStateNextHopTypeEnum
}{
	IPV4: BmpPrefixIpv6UnicastStateNextHopTypeEnum("ipv4"),
	IPV6: BmpPrefixIpv6UnicastStateNextHopTypeEnum("ipv6"),
}

func (obj *bmpPrefixIpv6UnicastState) NextHopType() BmpPrefixIpv6UnicastStateNextHopTypeEnum {
	return BmpPrefixIpv6UnicastStateNextHopTypeEnum(obj.obj.NextHopType.Enum().String())
}

// The next hop type , ipv4 or ipv6 .  If type is ipv4, then the ipv4_next_hop should be included. If type is ipv6, then the ipv6_next_hop should be included.
// NextHopType returns a string
func (obj *bmpPrefixIpv6UnicastState) HasNextHopType() bool {
	return obj.obj.NextHopType != nil
}

func (obj *bmpPrefixIpv6UnicastState) SetNextHopType(value BmpPrefixIpv6UnicastStateNextHopTypeEnum) BmpPrefixIpv6UnicastState {
	intValue, ok := otg.BmpPrefixIpv6UnicastState_NextHopType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BmpPrefixIpv6UnicastStateNextHopTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.BmpPrefixIpv6UnicastState_NextHopType_Enum(intValue)
	obj.obj.NextHopType = &enumValue

	return obj
}

// The IPv4 address of the egress interface.
// Ipv4NextHop returns a string
func (obj *bmpPrefixIpv6UnicastState) Ipv4NextHop() string {

	return *obj.obj.Ipv4NextHop

}

// The IPv4 address of the egress interface.
// Ipv4NextHop returns a string
func (obj *bmpPrefixIpv6UnicastState) HasIpv4NextHop() bool {
	return obj.obj.Ipv4NextHop != nil
}

// The IPv4 address of the egress interface.
// SetIpv4NextHop sets the string value in the BmpPrefixIpv6UnicastState object
func (obj *bmpPrefixIpv6UnicastState) SetIpv4NextHop(value string) BmpPrefixIpv6UnicastState {

	obj.obj.Ipv4NextHop = &value
	return obj
}

// The IPv6 address of the egress interface.
// Ipv6NextHop returns a string
func (obj *bmpPrefixIpv6UnicastState) Ipv6NextHop() string {

	return *obj.obj.Ipv6NextHop

}

// The IPv6 address of the egress interface.
// Ipv6NextHop returns a string
func (obj *bmpPrefixIpv6UnicastState) HasIpv6NextHop() bool {
	return obj.obj.Ipv6NextHop != nil
}

// The IPv6 address of the egress interface.
// SetIpv6NextHop sets the string value in the BmpPrefixIpv6UnicastState object
func (obj *bmpPrefixIpv6UnicastState) SetIpv6NextHop(value string) BmpPrefixIpv6UnicastState {

	obj.obj.Ipv6NextHop = &value
	return obj
}

// Optional community attributes.
// Communities returns a []ResultBgpCommunity
func (obj *bmpPrefixIpv6UnicastState) Communities() BmpPrefixIpv6UnicastStateResultBgpCommunityIter {
	if len(obj.obj.Communities) == 0 {
		obj.obj.Communities = []*otg.ResultBgpCommunity{}
	}
	if obj.communitiesHolder == nil {
		obj.communitiesHolder = newBmpPrefixIpv6UnicastStateResultBgpCommunityIter(&obj.obj.Communities).setMsg(obj)
	}
	return obj.communitiesHolder
}

type bmpPrefixIpv6UnicastStateResultBgpCommunityIter struct {
	obj                     *bmpPrefixIpv6UnicastState
	resultBgpCommunitySlice []ResultBgpCommunity
	fieldPtr                *[]*otg.ResultBgpCommunity
}

func newBmpPrefixIpv6UnicastStateResultBgpCommunityIter(ptr *[]*otg.ResultBgpCommunity) BmpPrefixIpv6UnicastStateResultBgpCommunityIter {
	return &bmpPrefixIpv6UnicastStateResultBgpCommunityIter{fieldPtr: ptr}
}

type BmpPrefixIpv6UnicastStateResultBgpCommunityIter interface {
	setMsg(*bmpPrefixIpv6UnicastState) BmpPrefixIpv6UnicastStateResultBgpCommunityIter
	Items() []ResultBgpCommunity
	Add() ResultBgpCommunity
	Append(items ...ResultBgpCommunity) BmpPrefixIpv6UnicastStateResultBgpCommunityIter
	Set(index int, newObj ResultBgpCommunity) BmpPrefixIpv6UnicastStateResultBgpCommunityIter
	Clear() BmpPrefixIpv6UnicastStateResultBgpCommunityIter
	clearHolderSlice() BmpPrefixIpv6UnicastStateResultBgpCommunityIter
	appendHolderSlice(item ResultBgpCommunity) BmpPrefixIpv6UnicastStateResultBgpCommunityIter
}

func (obj *bmpPrefixIpv6UnicastStateResultBgpCommunityIter) setMsg(msg *bmpPrefixIpv6UnicastState) BmpPrefixIpv6UnicastStateResultBgpCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&resultBgpCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bmpPrefixIpv6UnicastStateResultBgpCommunityIter) Items() []ResultBgpCommunity {
	return obj.resultBgpCommunitySlice
}

func (obj *bmpPrefixIpv6UnicastStateResultBgpCommunityIter) Add() ResultBgpCommunity {
	newObj := &otg.ResultBgpCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &resultBgpCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bmpPrefixIpv6UnicastStateResultBgpCommunityIter) Append(items ...ResultBgpCommunity) BmpPrefixIpv6UnicastStateResultBgpCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, item)
	}
	return obj
}

func (obj *bmpPrefixIpv6UnicastStateResultBgpCommunityIter) Set(index int, newObj ResultBgpCommunity) BmpPrefixIpv6UnicastStateResultBgpCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.resultBgpCommunitySlice[index] = newObj
	return obj
}
func (obj *bmpPrefixIpv6UnicastStateResultBgpCommunityIter) Clear() BmpPrefixIpv6UnicastStateResultBgpCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ResultBgpCommunity{}
		obj.resultBgpCommunitySlice = []ResultBgpCommunity{}
	}
	return obj
}
func (obj *bmpPrefixIpv6UnicastStateResultBgpCommunityIter) clearHolderSlice() BmpPrefixIpv6UnicastStateResultBgpCommunityIter {
	if len(obj.resultBgpCommunitySlice) > 0 {
		obj.resultBgpCommunitySlice = []ResultBgpCommunity{}
	}
	return obj
}
func (obj *bmpPrefixIpv6UnicastStateResultBgpCommunityIter) appendHolderSlice(item ResultBgpCommunity) BmpPrefixIpv6UnicastStateResultBgpCommunityIter {
	obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, item)
	return obj
}

// Optional received Extended Community attributes. Each received Extended Community attribute is available for retrieval in two forms. Support of the 'raw' format in which all 8 bytes (16 hex characters) is always present and available for use. In addition, if supported by the implementation, the Extended Community attribute may also be retrieved in the  'structured' format which is an optional field.
// ExtendedCommunities returns a []ResultExtendedCommunity
func (obj *bmpPrefixIpv6UnicastState) ExtendedCommunities() BmpPrefixIpv6UnicastStateResultExtendedCommunityIter {
	if len(obj.obj.ExtendedCommunities) == 0 {
		obj.obj.ExtendedCommunities = []*otg.ResultExtendedCommunity{}
	}
	if obj.extendedCommunitiesHolder == nil {
		obj.extendedCommunitiesHolder = newBmpPrefixIpv6UnicastStateResultExtendedCommunityIter(&obj.obj.ExtendedCommunities).setMsg(obj)
	}
	return obj.extendedCommunitiesHolder
}

type bmpPrefixIpv6UnicastStateResultExtendedCommunityIter struct {
	obj                          *bmpPrefixIpv6UnicastState
	resultExtendedCommunitySlice []ResultExtendedCommunity
	fieldPtr                     *[]*otg.ResultExtendedCommunity
}

func newBmpPrefixIpv6UnicastStateResultExtendedCommunityIter(ptr *[]*otg.ResultExtendedCommunity) BmpPrefixIpv6UnicastStateResultExtendedCommunityIter {
	return &bmpPrefixIpv6UnicastStateResultExtendedCommunityIter{fieldPtr: ptr}
}

type BmpPrefixIpv6UnicastStateResultExtendedCommunityIter interface {
	setMsg(*bmpPrefixIpv6UnicastState) BmpPrefixIpv6UnicastStateResultExtendedCommunityIter
	Items() []ResultExtendedCommunity
	Add() ResultExtendedCommunity
	Append(items ...ResultExtendedCommunity) BmpPrefixIpv6UnicastStateResultExtendedCommunityIter
	Set(index int, newObj ResultExtendedCommunity) BmpPrefixIpv6UnicastStateResultExtendedCommunityIter
	Clear() BmpPrefixIpv6UnicastStateResultExtendedCommunityIter
	clearHolderSlice() BmpPrefixIpv6UnicastStateResultExtendedCommunityIter
	appendHolderSlice(item ResultExtendedCommunity) BmpPrefixIpv6UnicastStateResultExtendedCommunityIter
}

func (obj *bmpPrefixIpv6UnicastStateResultExtendedCommunityIter) setMsg(msg *bmpPrefixIpv6UnicastState) BmpPrefixIpv6UnicastStateResultExtendedCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&resultExtendedCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bmpPrefixIpv6UnicastStateResultExtendedCommunityIter) Items() []ResultExtendedCommunity {
	return obj.resultExtendedCommunitySlice
}

func (obj *bmpPrefixIpv6UnicastStateResultExtendedCommunityIter) Add() ResultExtendedCommunity {
	newObj := &otg.ResultExtendedCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &resultExtendedCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bmpPrefixIpv6UnicastStateResultExtendedCommunityIter) Append(items ...ResultExtendedCommunity) BmpPrefixIpv6UnicastStateResultExtendedCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, item)
	}
	return obj
}

func (obj *bmpPrefixIpv6UnicastStateResultExtendedCommunityIter) Set(index int, newObj ResultExtendedCommunity) BmpPrefixIpv6UnicastStateResultExtendedCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.resultExtendedCommunitySlice[index] = newObj
	return obj
}
func (obj *bmpPrefixIpv6UnicastStateResultExtendedCommunityIter) Clear() BmpPrefixIpv6UnicastStateResultExtendedCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ResultExtendedCommunity{}
		obj.resultExtendedCommunitySlice = []ResultExtendedCommunity{}
	}
	return obj
}
func (obj *bmpPrefixIpv6UnicastStateResultExtendedCommunityIter) clearHolderSlice() BmpPrefixIpv6UnicastStateResultExtendedCommunityIter {
	if len(obj.resultExtendedCommunitySlice) > 0 {
		obj.resultExtendedCommunitySlice = []ResultExtendedCommunity{}
	}
	return obj
}
func (obj *bmpPrefixIpv6UnicastStateResultExtendedCommunityIter) appendHolderSlice(item ResultExtendedCommunity) BmpPrefixIpv6UnicastStateResultExtendedCommunityIter {
	obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, item)
	return obj
}

// description is TBD
// AsPath returns a ResultBgpAsPath
func (obj *bmpPrefixIpv6UnicastState) AsPath() ResultBgpAsPath {
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
func (obj *bmpPrefixIpv6UnicastState) HasAsPath() bool {
	return obj.obj.AsPath != nil
}

// description is TBD
// SetAsPath sets the ResultBgpAsPath value in the BmpPrefixIpv6UnicastState object
func (obj *bmpPrefixIpv6UnicastState) SetAsPath(value ResultBgpAsPath) BmpPrefixIpv6UnicastState {

	obj.asPathHolder = nil
	obj.obj.AsPath = value.msg()

	return obj
}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// LocalPreference returns a uint32
func (obj *bmpPrefixIpv6UnicastState) LocalPreference() uint32 {

	return *obj.obj.LocalPreference

}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// LocalPreference returns a uint32
func (obj *bmpPrefixIpv6UnicastState) HasLocalPreference() bool {
	return obj.obj.LocalPreference != nil
}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// SetLocalPreference sets the uint32 value in the BmpPrefixIpv6UnicastState object
func (obj *bmpPrefixIpv6UnicastState) SetLocalPreference(value uint32) BmpPrefixIpv6UnicastState {

	obj.obj.LocalPreference = &value
	return obj
}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// MultiExitDiscriminator returns a uint32
func (obj *bmpPrefixIpv6UnicastState) MultiExitDiscriminator() uint32 {

	return *obj.obj.MultiExitDiscriminator

}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// MultiExitDiscriminator returns a uint32
func (obj *bmpPrefixIpv6UnicastState) HasMultiExitDiscriminator() bool {
	return obj.obj.MultiExitDiscriminator != nil
}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// SetMultiExitDiscriminator sets the uint32 value in the BmpPrefixIpv6UnicastState object
func (obj *bmpPrefixIpv6UnicastState) SetMultiExitDiscriminator(value uint32) BmpPrefixIpv6UnicastState {

	obj.obj.MultiExitDiscriminator = &value
	return obj
}

func (obj *bmpPrefixIpv6UnicastState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv6Address != nil {

		err := obj.validateIpv6(obj.Ipv6Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BmpPrefixIpv6UnicastState.Ipv6Address"))
		}

	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BmpPrefixIpv6UnicastState.PrefixLength <= 128 but Got %d", *obj.obj.PrefixLength))
		}

	}

	if obj.obj.Ipv4NextHop != nil {

		err := obj.validateIpv4(obj.Ipv4NextHop())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BmpPrefixIpv6UnicastState.Ipv4NextHop"))
		}

	}

	if obj.obj.Ipv6NextHop != nil {

		err := obj.validateIpv6(obj.Ipv6NextHop())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BmpPrefixIpv6UnicastState.Ipv6NextHop"))
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

func (obj *bmpPrefixIpv6UnicastState) setDefault() {

}
