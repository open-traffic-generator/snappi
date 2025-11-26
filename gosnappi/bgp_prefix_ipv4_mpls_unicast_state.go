package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpPrefixIpv4MplsUnicastState *****
type bgpPrefixIpv4MplsUnicastState struct {
	validation
	obj                       *otg.BgpPrefixIpv4MplsUnicastState
	marshaller                marshalBgpPrefixIpv4MplsUnicastState
	unMarshaller              unMarshalBgpPrefixIpv4MplsUnicastState
	communitiesHolder         BgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter
	extendedCommunitiesHolder BgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter
	asPathHolder              ResultBgpAsPath
}

func NewBgpPrefixIpv4MplsUnicastState() BgpPrefixIpv4MplsUnicastState {
	obj := bgpPrefixIpv4MplsUnicastState{obj: &otg.BgpPrefixIpv4MplsUnicastState{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpPrefixIpv4MplsUnicastState) msg() *otg.BgpPrefixIpv4MplsUnicastState {
	return obj.obj
}

func (obj *bgpPrefixIpv4MplsUnicastState) setMsg(msg *otg.BgpPrefixIpv4MplsUnicastState) BgpPrefixIpv4MplsUnicastState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpPrefixIpv4MplsUnicastState struct {
	obj *bgpPrefixIpv4MplsUnicastState
}

type marshalBgpPrefixIpv4MplsUnicastState interface {
	// ToProto marshals BgpPrefixIpv4MplsUnicastState to protobuf object *otg.BgpPrefixIpv4MplsUnicastState
	ToProto() (*otg.BgpPrefixIpv4MplsUnicastState, error)
	// ToPbText marshals BgpPrefixIpv4MplsUnicastState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpPrefixIpv4MplsUnicastState to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpPrefixIpv4MplsUnicastState to JSON text
	ToJson() (string, error)
}

type unMarshalbgpPrefixIpv4MplsUnicastState struct {
	obj *bgpPrefixIpv4MplsUnicastState
}

type unMarshalBgpPrefixIpv4MplsUnicastState interface {
	// FromProto unmarshals BgpPrefixIpv4MplsUnicastState from protobuf object *otg.BgpPrefixIpv4MplsUnicastState
	FromProto(msg *otg.BgpPrefixIpv4MplsUnicastState) (BgpPrefixIpv4MplsUnicastState, error)
	// FromPbText unmarshals BgpPrefixIpv4MplsUnicastState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpPrefixIpv4MplsUnicastState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpPrefixIpv4MplsUnicastState from JSON text
	FromJson(value string) error
}

func (obj *bgpPrefixIpv4MplsUnicastState) Marshal() marshalBgpPrefixIpv4MplsUnicastState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpPrefixIpv4MplsUnicastState{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpPrefixIpv4MplsUnicastState) Unmarshal() unMarshalBgpPrefixIpv4MplsUnicastState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpPrefixIpv4MplsUnicastState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpPrefixIpv4MplsUnicastState) ToProto() (*otg.BgpPrefixIpv4MplsUnicastState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpPrefixIpv4MplsUnicastState) FromProto(msg *otg.BgpPrefixIpv4MplsUnicastState) (BgpPrefixIpv4MplsUnicastState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpPrefixIpv4MplsUnicastState) ToPbText() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4MplsUnicastState) FromPbText(value string) error {
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

func (m *marshalbgpPrefixIpv4MplsUnicastState) ToYaml() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4MplsUnicastState) FromYaml(value string) error {
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

func (m *marshalbgpPrefixIpv4MplsUnicastState) ToJson() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4MplsUnicastState) FromJson(value string) error {
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

func (obj *bgpPrefixIpv4MplsUnicastState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv4MplsUnicastState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv4MplsUnicastState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpPrefixIpv4MplsUnicastState) Clone() (BgpPrefixIpv4MplsUnicastState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpPrefixIpv4MplsUnicastState()
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

func (obj *bgpPrefixIpv4MplsUnicastState) setNil() {
	obj.communitiesHolder = nil
	obj.extendedCommunitiesHolder = nil
	obj.asPathHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpPrefixIpv4MplsUnicastState is iPv4 MPLS unicast prefix.
type BgpPrefixIpv4MplsUnicastState interface {
	Validation
	// msg marshals BgpPrefixIpv4MplsUnicastState to protobuf object *otg.BgpPrefixIpv4MplsUnicastState
	// and doesn't set defaults
	msg() *otg.BgpPrefixIpv4MplsUnicastState
	// setMsg unmarshals BgpPrefixIpv4MplsUnicastState from protobuf object *otg.BgpPrefixIpv4MplsUnicastState
	// and doesn't set defaults
	setMsg(*otg.BgpPrefixIpv4MplsUnicastState) BgpPrefixIpv4MplsUnicastState
	// provides marshal interface
	Marshal() marshalBgpPrefixIpv4MplsUnicastState
	// provides unmarshal interface
	Unmarshal() unMarshalBgpPrefixIpv4MplsUnicastState
	// validate validates BgpPrefixIpv4MplsUnicastState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpPrefixIpv4MplsUnicastState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4Address returns string, set in BgpPrefixIpv4MplsUnicastState.
	Ipv4Address() string
	// SetIpv4Address assigns string provided by user to BgpPrefixIpv4MplsUnicastState
	SetIpv4Address(value string) BgpPrefixIpv4MplsUnicastState
	// HasIpv4Address checks if Ipv4Address has been set in BgpPrefixIpv4MplsUnicastState
	HasIpv4Address() bool
	// PrefixLength returns uint32, set in BgpPrefixIpv4MplsUnicastState.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to BgpPrefixIpv4MplsUnicastState
	SetPrefixLength(value uint32) BgpPrefixIpv4MplsUnicastState
	// HasPrefixLength checks if PrefixLength has been set in BgpPrefixIpv4MplsUnicastState
	HasPrefixLength() bool
	// Origin returns BgpPrefixIpv4MplsUnicastStateOriginEnum, set in BgpPrefixIpv4MplsUnicastState
	Origin() BgpPrefixIpv4MplsUnicastStateOriginEnum
	// SetOrigin assigns BgpPrefixIpv4MplsUnicastStateOriginEnum provided by user to BgpPrefixIpv4MplsUnicastState
	SetOrigin(value BgpPrefixIpv4MplsUnicastStateOriginEnum) BgpPrefixIpv4MplsUnicastState
	// HasOrigin checks if Origin has been set in BgpPrefixIpv4MplsUnicastState
	HasOrigin() bool
	// PathId returns uint32, set in BgpPrefixIpv4MplsUnicastState.
	PathId() uint32
	// SetPathId assigns uint32 provided by user to BgpPrefixIpv4MplsUnicastState
	SetPathId(value uint32) BgpPrefixIpv4MplsUnicastState
	// HasPathId checks if PathId has been set in BgpPrefixIpv4MplsUnicastState
	HasPathId() bool
	// Ipv4NextHop returns string, set in BgpPrefixIpv4MplsUnicastState.
	Ipv4NextHop() string
	// SetIpv4NextHop assigns string provided by user to BgpPrefixIpv4MplsUnicastState
	SetIpv4NextHop(value string) BgpPrefixIpv4MplsUnicastState
	// HasIpv4NextHop checks if Ipv4NextHop has been set in BgpPrefixIpv4MplsUnicastState
	HasIpv4NextHop() bool
	// Ipv6NextHop returns string, set in BgpPrefixIpv4MplsUnicastState.
	Ipv6NextHop() string
	// SetIpv6NextHop assigns string provided by user to BgpPrefixIpv4MplsUnicastState
	SetIpv6NextHop(value string) BgpPrefixIpv4MplsUnicastState
	// HasIpv6NextHop checks if Ipv6NextHop has been set in BgpPrefixIpv4MplsUnicastState
	HasIpv6NextHop() bool
	// Labels returns []uint32, set in BgpPrefixIpv4MplsUnicastState.
	Labels() []uint32
	// SetLabels assigns []uint32 provided by user to BgpPrefixIpv4MplsUnicastState
	SetLabels(value []uint32) BgpPrefixIpv4MplsUnicastState
	// Communities returns BgpPrefixIpv4MplsUnicastStateResultBgpCommunityIterIter, set in BgpPrefixIpv4MplsUnicastState
	Communities() BgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter
	// ExtendedCommunities returns BgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIterIter, set in BgpPrefixIpv4MplsUnicastState
	ExtendedCommunities() BgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter
	// AsPath returns ResultBgpAsPath, set in BgpPrefixIpv4MplsUnicastState.
	// ResultBgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed.
	AsPath() ResultBgpAsPath
	// SetAsPath assigns ResultBgpAsPath provided by user to BgpPrefixIpv4MplsUnicastState.
	// ResultBgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed.
	SetAsPath(value ResultBgpAsPath) BgpPrefixIpv4MplsUnicastState
	// HasAsPath checks if AsPath has been set in BgpPrefixIpv4MplsUnicastState
	HasAsPath() bool
	// LocalPreference returns uint32, set in BgpPrefixIpv4MplsUnicastState.
	LocalPreference() uint32
	// SetLocalPreference assigns uint32 provided by user to BgpPrefixIpv4MplsUnicastState
	SetLocalPreference(value uint32) BgpPrefixIpv4MplsUnicastState
	// HasLocalPreference checks if LocalPreference has been set in BgpPrefixIpv4MplsUnicastState
	HasLocalPreference() bool
	// MultiExitDiscriminator returns uint32, set in BgpPrefixIpv4MplsUnicastState.
	MultiExitDiscriminator() uint32
	// SetMultiExitDiscriminator assigns uint32 provided by user to BgpPrefixIpv4MplsUnicastState
	SetMultiExitDiscriminator(value uint32) BgpPrefixIpv4MplsUnicastState
	// HasMultiExitDiscriminator checks if MultiExitDiscriminator has been set in BgpPrefixIpv4MplsUnicastState
	HasMultiExitDiscriminator() bool
	setNil()
}

// An IPv4 unicast address
// Ipv4Address returns a string
func (obj *bgpPrefixIpv4MplsUnicastState) Ipv4Address() string {

	return *obj.obj.Ipv4Address

}

// An IPv4 unicast address
// Ipv4Address returns a string
func (obj *bgpPrefixIpv4MplsUnicastState) HasIpv4Address() bool {
	return obj.obj.Ipv4Address != nil
}

// An IPv4 unicast address
// SetIpv4Address sets the string value in the BgpPrefixIpv4MplsUnicastState object
func (obj *bgpPrefixIpv4MplsUnicastState) SetIpv4Address(value string) BgpPrefixIpv4MplsUnicastState {

	obj.obj.Ipv4Address = &value
	return obj
}

// The length of the prefix.
// PrefixLength returns a uint32
func (obj *bgpPrefixIpv4MplsUnicastState) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// The length of the prefix.
// PrefixLength returns a uint32
func (obj *bgpPrefixIpv4MplsUnicastState) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// The length of the prefix.
// SetPrefixLength sets the uint32 value in the BgpPrefixIpv4MplsUnicastState object
func (obj *bgpPrefixIpv4MplsUnicastState) SetPrefixLength(value uint32) BgpPrefixIpv4MplsUnicastState {

	obj.obj.PrefixLength = &value
	return obj
}

type BgpPrefixIpv4MplsUnicastStateOriginEnum string

// Enum of Origin on BgpPrefixIpv4MplsUnicastState
var BgpPrefixIpv4MplsUnicastStateOrigin = struct {
	IGP        BgpPrefixIpv4MplsUnicastStateOriginEnum
	EGP        BgpPrefixIpv4MplsUnicastStateOriginEnum
	INCOMPLETE BgpPrefixIpv4MplsUnicastStateOriginEnum
}{
	IGP:        BgpPrefixIpv4MplsUnicastStateOriginEnum("igp"),
	EGP:        BgpPrefixIpv4MplsUnicastStateOriginEnum("egp"),
	INCOMPLETE: BgpPrefixIpv4MplsUnicastStateOriginEnum("incomplete"),
}

func (obj *bgpPrefixIpv4MplsUnicastState) Origin() BgpPrefixIpv4MplsUnicastStateOriginEnum {
	return BgpPrefixIpv4MplsUnicastStateOriginEnum(obj.obj.Origin.Enum().String())
}

// The origin of the prefix.
// Origin returns a string
func (obj *bgpPrefixIpv4MplsUnicastState) HasOrigin() bool {
	return obj.obj.Origin != nil
}

func (obj *bgpPrefixIpv4MplsUnicastState) SetOrigin(value BgpPrefixIpv4MplsUnicastStateOriginEnum) BgpPrefixIpv4MplsUnicastState {
	intValue, ok := otg.BgpPrefixIpv4MplsUnicastState_Origin_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpPrefixIpv4MplsUnicastStateOriginEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpPrefixIpv4MplsUnicastState_Origin_Enum(intValue)
	obj.obj.Origin = &enumValue

	return obj
}

// The path id.
// PathId returns a uint32
func (obj *bgpPrefixIpv4MplsUnicastState) PathId() uint32 {

	return *obj.obj.PathId

}

// The path id.
// PathId returns a uint32
func (obj *bgpPrefixIpv4MplsUnicastState) HasPathId() bool {
	return obj.obj.PathId != nil
}

// The path id.
// SetPathId sets the uint32 value in the BgpPrefixIpv4MplsUnicastState object
func (obj *bgpPrefixIpv4MplsUnicastState) SetPathId(value uint32) BgpPrefixIpv4MplsUnicastState {

	obj.obj.PathId = &value
	return obj
}

// The IPv4 address of the egress interface.
// Ipv4NextHop returns a string
func (obj *bgpPrefixIpv4MplsUnicastState) Ipv4NextHop() string {

	return *obj.obj.Ipv4NextHop

}

// The IPv4 address of the egress interface.
// Ipv4NextHop returns a string
func (obj *bgpPrefixIpv4MplsUnicastState) HasIpv4NextHop() bool {
	return obj.obj.Ipv4NextHop != nil
}

// The IPv4 address of the egress interface.
// SetIpv4NextHop sets the string value in the BgpPrefixIpv4MplsUnicastState object
func (obj *bgpPrefixIpv4MplsUnicastState) SetIpv4NextHop(value string) BgpPrefixIpv4MplsUnicastState {

	obj.obj.Ipv4NextHop = &value
	return obj
}

// The IPv6 address of the egress interface.
// Ipv6NextHop returns a string
func (obj *bgpPrefixIpv4MplsUnicastState) Ipv6NextHop() string {

	return *obj.obj.Ipv6NextHop

}

// The IPv6 address of the egress interface.
// Ipv6NextHop returns a string
func (obj *bgpPrefixIpv4MplsUnicastState) HasIpv6NextHop() bool {
	return obj.obj.Ipv6NextHop != nil
}

// The IPv6 address of the egress interface.
// SetIpv6NextHop sets the string value in the BgpPrefixIpv4MplsUnicastState object
func (obj *bgpPrefixIpv4MplsUnicastState) SetIpv6NextHop(value string) BgpPrefixIpv4MplsUnicastState {

	obj.obj.Ipv6NextHop = &value
	return obj
}

// Address Prefix that is bounded to One or More MPLS Labels.
// Labels returns a []uint32
func (obj *bgpPrefixIpv4MplsUnicastState) Labels() []uint32 {
	if obj.obj.Labels == nil {
		obj.obj.Labels = make([]uint32, 0)
	}
	return obj.obj.Labels
}

// Address Prefix that is bounded to One or More MPLS Labels.
// SetLabels sets the []uint32 value in the BgpPrefixIpv4MplsUnicastState object
func (obj *bgpPrefixIpv4MplsUnicastState) SetLabels(value []uint32) BgpPrefixIpv4MplsUnicastState {

	if obj.obj.Labels == nil {
		obj.obj.Labels = make([]uint32, 0)
	}
	obj.obj.Labels = value

	return obj
}

// Optional community attributes.
// Communities returns a []ResultBgpCommunity
func (obj *bgpPrefixIpv4MplsUnicastState) Communities() BgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter {
	if len(obj.obj.Communities) == 0 {
		obj.obj.Communities = []*otg.ResultBgpCommunity{}
	}
	if obj.communitiesHolder == nil {
		obj.communitiesHolder = newBgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter(&obj.obj.Communities).setMsg(obj)
	}
	return obj.communitiesHolder
}

type bgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter struct {
	obj                     *bgpPrefixIpv4MplsUnicastState
	resultBgpCommunitySlice []ResultBgpCommunity
	fieldPtr                *[]*otg.ResultBgpCommunity
}

func newBgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter(ptr *[]*otg.ResultBgpCommunity) BgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter {
	return &bgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter{fieldPtr: ptr}
}

type BgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter interface {
	setMsg(*bgpPrefixIpv4MplsUnicastState) BgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter
	Items() []ResultBgpCommunity
	Add() ResultBgpCommunity
	Append(items ...ResultBgpCommunity) BgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter
	Set(index int, newObj ResultBgpCommunity) BgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter
	Clear() BgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter
	clearHolderSlice() BgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter
	appendHolderSlice(item ResultBgpCommunity) BgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter
}

func (obj *bgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter) setMsg(msg *bgpPrefixIpv4MplsUnicastState) BgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&resultBgpCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter) Items() []ResultBgpCommunity {
	return obj.resultBgpCommunitySlice
}

func (obj *bgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter) Add() ResultBgpCommunity {
	newObj := &otg.ResultBgpCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &resultBgpCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter) Append(items ...ResultBgpCommunity) BgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, item)
	}
	return obj
}

func (obj *bgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter) Set(index int, newObj ResultBgpCommunity) BgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.resultBgpCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter) Clear() BgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ResultBgpCommunity{}
		obj.resultBgpCommunitySlice = []ResultBgpCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter) clearHolderSlice() BgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter {
	if len(obj.resultBgpCommunitySlice) > 0 {
		obj.resultBgpCommunitySlice = []ResultBgpCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter) appendHolderSlice(item ResultBgpCommunity) BgpPrefixIpv4MplsUnicastStateResultBgpCommunityIter {
	obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, item)
	return obj
}

// Optional received Extended Community attributes. Each received Extended Community attribute is available for retrieval in two forms. Support of the 'raw' format in which all 8 bytes (16 hex characters) is always present and available for use. In addition, if supported by the implementation, the Extended Community attribute may also be retrieved in the  'structured' format which is an optional field.
// ExtendedCommunities returns a []ResultExtendedCommunity
func (obj *bgpPrefixIpv4MplsUnicastState) ExtendedCommunities() BgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter {
	if len(obj.obj.ExtendedCommunities) == 0 {
		obj.obj.ExtendedCommunities = []*otg.ResultExtendedCommunity{}
	}
	if obj.extendedCommunitiesHolder == nil {
		obj.extendedCommunitiesHolder = newBgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter(&obj.obj.ExtendedCommunities).setMsg(obj)
	}
	return obj.extendedCommunitiesHolder
}

type bgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter struct {
	obj                          *bgpPrefixIpv4MplsUnicastState
	resultExtendedCommunitySlice []ResultExtendedCommunity
	fieldPtr                     *[]*otg.ResultExtendedCommunity
}

func newBgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter(ptr *[]*otg.ResultExtendedCommunity) BgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter {
	return &bgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter{fieldPtr: ptr}
}

type BgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter interface {
	setMsg(*bgpPrefixIpv4MplsUnicastState) BgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter
	Items() []ResultExtendedCommunity
	Add() ResultExtendedCommunity
	Append(items ...ResultExtendedCommunity) BgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter
	Set(index int, newObj ResultExtendedCommunity) BgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter
	Clear() BgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter
	clearHolderSlice() BgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter
	appendHolderSlice(item ResultExtendedCommunity) BgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter
}

func (obj *bgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter) setMsg(msg *bgpPrefixIpv4MplsUnicastState) BgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&resultExtendedCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter) Items() []ResultExtendedCommunity {
	return obj.resultExtendedCommunitySlice
}

func (obj *bgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter) Add() ResultExtendedCommunity {
	newObj := &otg.ResultExtendedCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &resultExtendedCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter) Append(items ...ResultExtendedCommunity) BgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, item)
	}
	return obj
}

func (obj *bgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter) Set(index int, newObj ResultExtendedCommunity) BgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.resultExtendedCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter) Clear() BgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ResultExtendedCommunity{}
		obj.resultExtendedCommunitySlice = []ResultExtendedCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter) clearHolderSlice() BgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter {
	if len(obj.resultExtendedCommunitySlice) > 0 {
		obj.resultExtendedCommunitySlice = []ResultExtendedCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter) appendHolderSlice(item ResultExtendedCommunity) BgpPrefixIpv4MplsUnicastStateResultExtendedCommunityIter {
	obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, item)
	return obj
}

// description is TBD
// AsPath returns a ResultBgpAsPath
func (obj *bgpPrefixIpv4MplsUnicastState) AsPath() ResultBgpAsPath {
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
func (obj *bgpPrefixIpv4MplsUnicastState) HasAsPath() bool {
	return obj.obj.AsPath != nil
}

// description is TBD
// SetAsPath sets the ResultBgpAsPath value in the BgpPrefixIpv4MplsUnicastState object
func (obj *bgpPrefixIpv4MplsUnicastState) SetAsPath(value ResultBgpAsPath) BgpPrefixIpv4MplsUnicastState {

	obj.asPathHolder = nil
	obj.obj.AsPath = value.msg()

	return obj
}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// LocalPreference returns a uint32
func (obj *bgpPrefixIpv4MplsUnicastState) LocalPreference() uint32 {

	return *obj.obj.LocalPreference

}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// LocalPreference returns a uint32
func (obj *bgpPrefixIpv4MplsUnicastState) HasLocalPreference() bool {
	return obj.obj.LocalPreference != nil
}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// SetLocalPreference sets the uint32 value in the BgpPrefixIpv4MplsUnicastState object
func (obj *bgpPrefixIpv4MplsUnicastState) SetLocalPreference(value uint32) BgpPrefixIpv4MplsUnicastState {

	obj.obj.LocalPreference = &value
	return obj
}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// MultiExitDiscriminator returns a uint32
func (obj *bgpPrefixIpv4MplsUnicastState) MultiExitDiscriminator() uint32 {

	return *obj.obj.MultiExitDiscriminator

}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// MultiExitDiscriminator returns a uint32
func (obj *bgpPrefixIpv4MplsUnicastState) HasMultiExitDiscriminator() bool {
	return obj.obj.MultiExitDiscriminator != nil
}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// SetMultiExitDiscriminator sets the uint32 value in the BgpPrefixIpv4MplsUnicastState object
func (obj *bgpPrefixIpv4MplsUnicastState) SetMultiExitDiscriminator(value uint32) BgpPrefixIpv4MplsUnicastState {

	obj.obj.MultiExitDiscriminator = &value
	return obj
}

func (obj *bgpPrefixIpv4MplsUnicastState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpPrefixIpv4MplsUnicastState.PrefixLength <= 128 but Got %d", *obj.obj.PrefixLength))
		}

	}

	if obj.obj.Ipv4NextHop != nil {

		err := obj.validateIpv4(obj.Ipv4NextHop())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv4MplsUnicastState.Ipv4NextHop"))
		}

	}

	if obj.obj.Ipv6NextHop != nil {

		err := obj.validateIpv6(obj.Ipv6NextHop())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv4MplsUnicastState.Ipv6NextHop"))
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

func (obj *bgpPrefixIpv4MplsUnicastState) setDefault() {

}
