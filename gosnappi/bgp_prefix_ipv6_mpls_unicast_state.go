package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpPrefixIpv6MplsUnicastState *****
type bgpPrefixIpv6MplsUnicastState struct {
	validation
	obj                       *otg.BgpPrefixIpv6MplsUnicastState
	marshaller                marshalBgpPrefixIpv6MplsUnicastState
	unMarshaller              unMarshalBgpPrefixIpv6MplsUnicastState
	communitiesHolder         BgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter
	extendedCommunitiesHolder BgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter
	asPathHolder              ResultBgpAsPath
}

func NewBgpPrefixIpv6MplsUnicastState() BgpPrefixIpv6MplsUnicastState {
	obj := bgpPrefixIpv6MplsUnicastState{obj: &otg.BgpPrefixIpv6MplsUnicastState{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpPrefixIpv6MplsUnicastState) msg() *otg.BgpPrefixIpv6MplsUnicastState {
	return obj.obj
}

func (obj *bgpPrefixIpv6MplsUnicastState) setMsg(msg *otg.BgpPrefixIpv6MplsUnicastState) BgpPrefixIpv6MplsUnicastState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpPrefixIpv6MplsUnicastState struct {
	obj *bgpPrefixIpv6MplsUnicastState
}

type marshalBgpPrefixIpv6MplsUnicastState interface {
	// ToProto marshals BgpPrefixIpv6MplsUnicastState to protobuf object *otg.BgpPrefixIpv6MplsUnicastState
	ToProto() (*otg.BgpPrefixIpv6MplsUnicastState, error)
	// ToPbText marshals BgpPrefixIpv6MplsUnicastState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpPrefixIpv6MplsUnicastState to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpPrefixIpv6MplsUnicastState to JSON text
	ToJson() (string, error)
}

type unMarshalbgpPrefixIpv6MplsUnicastState struct {
	obj *bgpPrefixIpv6MplsUnicastState
}

type unMarshalBgpPrefixIpv6MplsUnicastState interface {
	// FromProto unmarshals BgpPrefixIpv6MplsUnicastState from protobuf object *otg.BgpPrefixIpv6MplsUnicastState
	FromProto(msg *otg.BgpPrefixIpv6MplsUnicastState) (BgpPrefixIpv6MplsUnicastState, error)
	// FromPbText unmarshals BgpPrefixIpv6MplsUnicastState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpPrefixIpv6MplsUnicastState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpPrefixIpv6MplsUnicastState from JSON text
	FromJson(value string) error
}

func (obj *bgpPrefixIpv6MplsUnicastState) Marshal() marshalBgpPrefixIpv6MplsUnicastState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpPrefixIpv6MplsUnicastState{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpPrefixIpv6MplsUnicastState) Unmarshal() unMarshalBgpPrefixIpv6MplsUnicastState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpPrefixIpv6MplsUnicastState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpPrefixIpv6MplsUnicastState) ToProto() (*otg.BgpPrefixIpv6MplsUnicastState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpPrefixIpv6MplsUnicastState) FromProto(msg *otg.BgpPrefixIpv6MplsUnicastState) (BgpPrefixIpv6MplsUnicastState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpPrefixIpv6MplsUnicastState) ToPbText() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6MplsUnicastState) FromPbText(value string) error {
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

func (m *marshalbgpPrefixIpv6MplsUnicastState) ToYaml() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6MplsUnicastState) FromYaml(value string) error {
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

func (m *marshalbgpPrefixIpv6MplsUnicastState) ToJson() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6MplsUnicastState) FromJson(value string) error {
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

func (obj *bgpPrefixIpv6MplsUnicastState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv6MplsUnicastState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv6MplsUnicastState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpPrefixIpv6MplsUnicastState) Clone() (BgpPrefixIpv6MplsUnicastState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpPrefixIpv6MplsUnicastState()
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

func (obj *bgpPrefixIpv6MplsUnicastState) setNil() {
	obj.communitiesHolder = nil
	obj.extendedCommunitiesHolder = nil
	obj.asPathHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpPrefixIpv6MplsUnicastState is iPv6 MPLS unicast prefix.
type BgpPrefixIpv6MplsUnicastState interface {
	Validation
	// msg marshals BgpPrefixIpv6MplsUnicastState to protobuf object *otg.BgpPrefixIpv6MplsUnicastState
	// and doesn't set defaults
	msg() *otg.BgpPrefixIpv6MplsUnicastState
	// setMsg unmarshals BgpPrefixIpv6MplsUnicastState from protobuf object *otg.BgpPrefixIpv6MplsUnicastState
	// and doesn't set defaults
	setMsg(*otg.BgpPrefixIpv6MplsUnicastState) BgpPrefixIpv6MplsUnicastState
	// provides marshal interface
	Marshal() marshalBgpPrefixIpv6MplsUnicastState
	// provides unmarshal interface
	Unmarshal() unMarshalBgpPrefixIpv6MplsUnicastState
	// validate validates BgpPrefixIpv6MplsUnicastState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpPrefixIpv6MplsUnicastState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv6Address returns string, set in BgpPrefixIpv6MplsUnicastState.
	Ipv6Address() string
	// SetIpv6Address assigns string provided by user to BgpPrefixIpv6MplsUnicastState
	SetIpv6Address(value string) BgpPrefixIpv6MplsUnicastState
	// HasIpv6Address checks if Ipv6Address has been set in BgpPrefixIpv6MplsUnicastState
	HasIpv6Address() bool
	// PrefixLength returns uint32, set in BgpPrefixIpv6MplsUnicastState.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to BgpPrefixIpv6MplsUnicastState
	SetPrefixLength(value uint32) BgpPrefixIpv6MplsUnicastState
	// HasPrefixLength checks if PrefixLength has been set in BgpPrefixIpv6MplsUnicastState
	HasPrefixLength() bool
	// Origin returns BgpPrefixIpv6MplsUnicastStateOriginEnum, set in BgpPrefixIpv6MplsUnicastState
	Origin() BgpPrefixIpv6MplsUnicastStateOriginEnum
	// SetOrigin assigns BgpPrefixIpv6MplsUnicastStateOriginEnum provided by user to BgpPrefixIpv6MplsUnicastState
	SetOrigin(value BgpPrefixIpv6MplsUnicastStateOriginEnum) BgpPrefixIpv6MplsUnicastState
	// HasOrigin checks if Origin has been set in BgpPrefixIpv6MplsUnicastState
	HasOrigin() bool
	// PathId returns uint32, set in BgpPrefixIpv6MplsUnicastState.
	PathId() uint32
	// SetPathId assigns uint32 provided by user to BgpPrefixIpv6MplsUnicastState
	SetPathId(value uint32) BgpPrefixIpv6MplsUnicastState
	// HasPathId checks if PathId has been set in BgpPrefixIpv6MplsUnicastState
	HasPathId() bool
	// Ipv4NextHop returns string, set in BgpPrefixIpv6MplsUnicastState.
	Ipv4NextHop() string
	// SetIpv4NextHop assigns string provided by user to BgpPrefixIpv6MplsUnicastState
	SetIpv4NextHop(value string) BgpPrefixIpv6MplsUnicastState
	// HasIpv4NextHop checks if Ipv4NextHop has been set in BgpPrefixIpv6MplsUnicastState
	HasIpv4NextHop() bool
	// Ipv6NextHop returns string, set in BgpPrefixIpv6MplsUnicastState.
	Ipv6NextHop() string
	// SetIpv6NextHop assigns string provided by user to BgpPrefixIpv6MplsUnicastState
	SetIpv6NextHop(value string) BgpPrefixIpv6MplsUnicastState
	// HasIpv6NextHop checks if Ipv6NextHop has been set in BgpPrefixIpv6MplsUnicastState
	HasIpv6NextHop() bool
	// Labels returns []uint32, set in BgpPrefixIpv6MplsUnicastState.
	Labels() []uint32
	// SetLabels assigns []uint32 provided by user to BgpPrefixIpv6MplsUnicastState
	SetLabels(value []uint32) BgpPrefixIpv6MplsUnicastState
	// Communities returns BgpPrefixIpv6MplsUnicastStateResultBgpCommunityIterIter, set in BgpPrefixIpv6MplsUnicastState
	Communities() BgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter
	// ExtendedCommunities returns BgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIterIter, set in BgpPrefixIpv6MplsUnicastState
	ExtendedCommunities() BgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter
	// AsPath returns ResultBgpAsPath, set in BgpPrefixIpv6MplsUnicastState.
	// ResultBgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed.
	AsPath() ResultBgpAsPath
	// SetAsPath assigns ResultBgpAsPath provided by user to BgpPrefixIpv6MplsUnicastState.
	// ResultBgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed.
	SetAsPath(value ResultBgpAsPath) BgpPrefixIpv6MplsUnicastState
	// HasAsPath checks if AsPath has been set in BgpPrefixIpv6MplsUnicastState
	HasAsPath() bool
	// LocalPreference returns uint32, set in BgpPrefixIpv6MplsUnicastState.
	LocalPreference() uint32
	// SetLocalPreference assigns uint32 provided by user to BgpPrefixIpv6MplsUnicastState
	SetLocalPreference(value uint32) BgpPrefixIpv6MplsUnicastState
	// HasLocalPreference checks if LocalPreference has been set in BgpPrefixIpv6MplsUnicastState
	HasLocalPreference() bool
	// MultiExitDiscriminator returns uint32, set in BgpPrefixIpv6MplsUnicastState.
	MultiExitDiscriminator() uint32
	// SetMultiExitDiscriminator assigns uint32 provided by user to BgpPrefixIpv6MplsUnicastState
	SetMultiExitDiscriminator(value uint32) BgpPrefixIpv6MplsUnicastState
	// HasMultiExitDiscriminator checks if MultiExitDiscriminator has been set in BgpPrefixIpv6MplsUnicastState
	HasMultiExitDiscriminator() bool
	setNil()
}

// An IPv6 unicast address
// Ipv6Address returns a string
func (obj *bgpPrefixIpv6MplsUnicastState) Ipv6Address() string {

	return *obj.obj.Ipv6Address

}

// An IPv6 unicast address
// Ipv6Address returns a string
func (obj *bgpPrefixIpv6MplsUnicastState) HasIpv6Address() bool {
	return obj.obj.Ipv6Address != nil
}

// An IPv6 unicast address
// SetIpv6Address sets the string value in the BgpPrefixIpv6MplsUnicastState object
func (obj *bgpPrefixIpv6MplsUnicastState) SetIpv6Address(value string) BgpPrefixIpv6MplsUnicastState {

	obj.obj.Ipv6Address = &value
	return obj
}

// description is TBD
// PrefixLength returns a uint32
func (obj *bgpPrefixIpv6MplsUnicastState) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// description is TBD
// PrefixLength returns a uint32
func (obj *bgpPrefixIpv6MplsUnicastState) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// description is TBD
// SetPrefixLength sets the uint32 value in the BgpPrefixIpv6MplsUnicastState object
func (obj *bgpPrefixIpv6MplsUnicastState) SetPrefixLength(value uint32) BgpPrefixIpv6MplsUnicastState {

	obj.obj.PrefixLength = &value
	return obj
}

type BgpPrefixIpv6MplsUnicastStateOriginEnum string

// Enum of Origin on BgpPrefixIpv6MplsUnicastState
var BgpPrefixIpv6MplsUnicastStateOrigin = struct {
	IGP        BgpPrefixIpv6MplsUnicastStateOriginEnum
	EGP        BgpPrefixIpv6MplsUnicastStateOriginEnum
	INCOMPLETE BgpPrefixIpv6MplsUnicastStateOriginEnum
}{
	IGP:        BgpPrefixIpv6MplsUnicastStateOriginEnum("igp"),
	EGP:        BgpPrefixIpv6MplsUnicastStateOriginEnum("egp"),
	INCOMPLETE: BgpPrefixIpv6MplsUnicastStateOriginEnum("incomplete"),
}

func (obj *bgpPrefixIpv6MplsUnicastState) Origin() BgpPrefixIpv6MplsUnicastStateOriginEnum {
	return BgpPrefixIpv6MplsUnicastStateOriginEnum(obj.obj.Origin.Enum().String())
}

// The origin of the prefix.
// Origin returns a string
func (obj *bgpPrefixIpv6MplsUnicastState) HasOrigin() bool {
	return obj.obj.Origin != nil
}

func (obj *bgpPrefixIpv6MplsUnicastState) SetOrigin(value BgpPrefixIpv6MplsUnicastStateOriginEnum) BgpPrefixIpv6MplsUnicastState {
	intValue, ok := otg.BgpPrefixIpv6MplsUnicastState_Origin_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpPrefixIpv6MplsUnicastStateOriginEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpPrefixIpv6MplsUnicastState_Origin_Enum(intValue)
	obj.obj.Origin = &enumValue

	return obj
}

// The path id.
// PathId returns a uint32
func (obj *bgpPrefixIpv6MplsUnicastState) PathId() uint32 {

	return *obj.obj.PathId

}

// The path id.
// PathId returns a uint32
func (obj *bgpPrefixIpv6MplsUnicastState) HasPathId() bool {
	return obj.obj.PathId != nil
}

// The path id.
// SetPathId sets the uint32 value in the BgpPrefixIpv6MplsUnicastState object
func (obj *bgpPrefixIpv6MplsUnicastState) SetPathId(value uint32) BgpPrefixIpv6MplsUnicastState {

	obj.obj.PathId = &value
	return obj
}

// The IPv4 address of the egress interface.
// Ipv4NextHop returns a string
func (obj *bgpPrefixIpv6MplsUnicastState) Ipv4NextHop() string {

	return *obj.obj.Ipv4NextHop

}

// The IPv4 address of the egress interface.
// Ipv4NextHop returns a string
func (obj *bgpPrefixIpv6MplsUnicastState) HasIpv4NextHop() bool {
	return obj.obj.Ipv4NextHop != nil
}

// The IPv4 address of the egress interface.
// SetIpv4NextHop sets the string value in the BgpPrefixIpv6MplsUnicastState object
func (obj *bgpPrefixIpv6MplsUnicastState) SetIpv4NextHop(value string) BgpPrefixIpv6MplsUnicastState {

	obj.obj.Ipv4NextHop = &value
	return obj
}

// The IPv6 address of the egress interface.
// Ipv6NextHop returns a string
func (obj *bgpPrefixIpv6MplsUnicastState) Ipv6NextHop() string {

	return *obj.obj.Ipv6NextHop

}

// The IPv6 address of the egress interface.
// Ipv6NextHop returns a string
func (obj *bgpPrefixIpv6MplsUnicastState) HasIpv6NextHop() bool {
	return obj.obj.Ipv6NextHop != nil
}

// The IPv6 address of the egress interface.
// SetIpv6NextHop sets the string value in the BgpPrefixIpv6MplsUnicastState object
func (obj *bgpPrefixIpv6MplsUnicastState) SetIpv6NextHop(value string) BgpPrefixIpv6MplsUnicastState {

	obj.obj.Ipv6NextHop = &value
	return obj
}

// One or more MPLS Label 24 bit values bound to this address prefix.
// Labels returns a []uint32
func (obj *bgpPrefixIpv6MplsUnicastState) Labels() []uint32 {
	if obj.obj.Labels == nil {
		obj.obj.Labels = make([]uint32, 0)
	}
	return obj.obj.Labels
}

// One or more MPLS Label 24 bit values bound to this address prefix.
// SetLabels sets the []uint32 value in the BgpPrefixIpv6MplsUnicastState object
func (obj *bgpPrefixIpv6MplsUnicastState) SetLabels(value []uint32) BgpPrefixIpv6MplsUnicastState {

	if obj.obj.Labels == nil {
		obj.obj.Labels = make([]uint32, 0)
	}
	obj.obj.Labels = value

	return obj
}

// Optional community attributes.
// Communities returns a []ResultBgpCommunity
func (obj *bgpPrefixIpv6MplsUnicastState) Communities() BgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter {
	if len(obj.obj.Communities) == 0 {
		obj.obj.Communities = []*otg.ResultBgpCommunity{}
	}
	if obj.communitiesHolder == nil {
		obj.communitiesHolder = newBgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter(&obj.obj.Communities).setMsg(obj)
	}
	return obj.communitiesHolder
}

type bgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter struct {
	obj                     *bgpPrefixIpv6MplsUnicastState
	resultBgpCommunitySlice []ResultBgpCommunity
	fieldPtr                *[]*otg.ResultBgpCommunity
}

func newBgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter(ptr *[]*otg.ResultBgpCommunity) BgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter {
	return &bgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter{fieldPtr: ptr}
}

type BgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter interface {
	setMsg(*bgpPrefixIpv6MplsUnicastState) BgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter
	Items() []ResultBgpCommunity
	Add() ResultBgpCommunity
	Append(items ...ResultBgpCommunity) BgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter
	Set(index int, newObj ResultBgpCommunity) BgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter
	Clear() BgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter
	clearHolderSlice() BgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter
	appendHolderSlice(item ResultBgpCommunity) BgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter
}

func (obj *bgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter) setMsg(msg *bgpPrefixIpv6MplsUnicastState) BgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&resultBgpCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter) Items() []ResultBgpCommunity {
	return obj.resultBgpCommunitySlice
}

func (obj *bgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter) Add() ResultBgpCommunity {
	newObj := &otg.ResultBgpCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &resultBgpCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter) Append(items ...ResultBgpCommunity) BgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, item)
	}
	return obj
}

func (obj *bgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter) Set(index int, newObj ResultBgpCommunity) BgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.resultBgpCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter) Clear() BgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ResultBgpCommunity{}
		obj.resultBgpCommunitySlice = []ResultBgpCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter) clearHolderSlice() BgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter {
	if len(obj.resultBgpCommunitySlice) > 0 {
		obj.resultBgpCommunitySlice = []ResultBgpCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter) appendHolderSlice(item ResultBgpCommunity) BgpPrefixIpv6MplsUnicastStateResultBgpCommunityIter {
	obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, item)
	return obj
}

// Optional received Extended Community attributes. Each received Extended Community attribute is available for retrieval in two forms. Support of the 'raw' format in which all 8 bytes (16 hex characters) is always present and available for use. In addition, if supported by the implementation, the Extended Community attribute may also be retrieved in the  'structured' format which is an optional field.
// ExtendedCommunities returns a []ResultExtendedCommunity
func (obj *bgpPrefixIpv6MplsUnicastState) ExtendedCommunities() BgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter {
	if len(obj.obj.ExtendedCommunities) == 0 {
		obj.obj.ExtendedCommunities = []*otg.ResultExtendedCommunity{}
	}
	if obj.extendedCommunitiesHolder == nil {
		obj.extendedCommunitiesHolder = newBgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter(&obj.obj.ExtendedCommunities).setMsg(obj)
	}
	return obj.extendedCommunitiesHolder
}

type bgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter struct {
	obj                          *bgpPrefixIpv6MplsUnicastState
	resultExtendedCommunitySlice []ResultExtendedCommunity
	fieldPtr                     *[]*otg.ResultExtendedCommunity
}

func newBgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter(ptr *[]*otg.ResultExtendedCommunity) BgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter {
	return &bgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter{fieldPtr: ptr}
}

type BgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter interface {
	setMsg(*bgpPrefixIpv6MplsUnicastState) BgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter
	Items() []ResultExtendedCommunity
	Add() ResultExtendedCommunity
	Append(items ...ResultExtendedCommunity) BgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter
	Set(index int, newObj ResultExtendedCommunity) BgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter
	Clear() BgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter
	clearHolderSlice() BgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter
	appendHolderSlice(item ResultExtendedCommunity) BgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter
}

func (obj *bgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter) setMsg(msg *bgpPrefixIpv6MplsUnicastState) BgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&resultExtendedCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter) Items() []ResultExtendedCommunity {
	return obj.resultExtendedCommunitySlice
}

func (obj *bgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter) Add() ResultExtendedCommunity {
	newObj := &otg.ResultExtendedCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &resultExtendedCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter) Append(items ...ResultExtendedCommunity) BgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, item)
	}
	return obj
}

func (obj *bgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter) Set(index int, newObj ResultExtendedCommunity) BgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.resultExtendedCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter) Clear() BgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ResultExtendedCommunity{}
		obj.resultExtendedCommunitySlice = []ResultExtendedCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter) clearHolderSlice() BgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter {
	if len(obj.resultExtendedCommunitySlice) > 0 {
		obj.resultExtendedCommunitySlice = []ResultExtendedCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter) appendHolderSlice(item ResultExtendedCommunity) BgpPrefixIpv6MplsUnicastStateResultExtendedCommunityIter {
	obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, item)
	return obj
}

// description is TBD
// AsPath returns a ResultBgpAsPath
func (obj *bgpPrefixIpv6MplsUnicastState) AsPath() ResultBgpAsPath {
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
func (obj *bgpPrefixIpv6MplsUnicastState) HasAsPath() bool {
	return obj.obj.AsPath != nil
}

// description is TBD
// SetAsPath sets the ResultBgpAsPath value in the BgpPrefixIpv6MplsUnicastState object
func (obj *bgpPrefixIpv6MplsUnicastState) SetAsPath(value ResultBgpAsPath) BgpPrefixIpv6MplsUnicastState {

	obj.asPathHolder = nil
	obj.obj.AsPath = value.msg()

	return obj
}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// LocalPreference returns a uint32
func (obj *bgpPrefixIpv6MplsUnicastState) LocalPreference() uint32 {

	return *obj.obj.LocalPreference

}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// LocalPreference returns a uint32
func (obj *bgpPrefixIpv6MplsUnicastState) HasLocalPreference() bool {
	return obj.obj.LocalPreference != nil
}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// SetLocalPreference sets the uint32 value in the BgpPrefixIpv6MplsUnicastState object
func (obj *bgpPrefixIpv6MplsUnicastState) SetLocalPreference(value uint32) BgpPrefixIpv6MplsUnicastState {

	obj.obj.LocalPreference = &value
	return obj
}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// MultiExitDiscriminator returns a uint32
func (obj *bgpPrefixIpv6MplsUnicastState) MultiExitDiscriminator() uint32 {

	return *obj.obj.MultiExitDiscriminator

}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// MultiExitDiscriminator returns a uint32
func (obj *bgpPrefixIpv6MplsUnicastState) HasMultiExitDiscriminator() bool {
	return obj.obj.MultiExitDiscriminator != nil
}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// SetMultiExitDiscriminator sets the uint32 value in the BgpPrefixIpv6MplsUnicastState object
func (obj *bgpPrefixIpv6MplsUnicastState) SetMultiExitDiscriminator(value uint32) BgpPrefixIpv6MplsUnicastState {

	obj.obj.MultiExitDiscriminator = &value
	return obj
}

func (obj *bgpPrefixIpv6MplsUnicastState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpPrefixIpv6MplsUnicastState.PrefixLength <= 128 but Got %d", *obj.obj.PrefixLength))
		}

	}

	if obj.obj.Ipv4NextHop != nil {

		err := obj.validateIpv4(obj.Ipv4NextHop())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv6MplsUnicastState.Ipv4NextHop"))
		}

	}

	if obj.obj.Ipv6NextHop != nil {

		err := obj.validateIpv6(obj.Ipv6NextHop())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv6MplsUnicastState.Ipv6NextHop"))
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

func (obj *bgpPrefixIpv6MplsUnicastState) setDefault() {

}
