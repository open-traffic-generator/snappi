package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpPrefixIpv4UnicastState *****
type bgpPrefixIpv4UnicastState struct {
	validation
	obj                       *otg.BgpPrefixIpv4UnicastState
	marshaller                marshalBgpPrefixIpv4UnicastState
	unMarshaller              unMarshalBgpPrefixIpv4UnicastState
	communitiesHolder         BgpPrefixIpv4UnicastStateResultBgpCommunityIter
	extendedCommunitiesHolder BgpPrefixIpv4UnicastStateResultExtendedCommunityIter
	asPathHolder              ResultBgpAsPath
}

func NewBgpPrefixIpv4UnicastState() BgpPrefixIpv4UnicastState {
	obj := bgpPrefixIpv4UnicastState{obj: &otg.BgpPrefixIpv4UnicastState{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpPrefixIpv4UnicastState) msg() *otg.BgpPrefixIpv4UnicastState {
	return obj.obj
}

func (obj *bgpPrefixIpv4UnicastState) setMsg(msg *otg.BgpPrefixIpv4UnicastState) BgpPrefixIpv4UnicastState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpPrefixIpv4UnicastState struct {
	obj *bgpPrefixIpv4UnicastState
}

type marshalBgpPrefixIpv4UnicastState interface {
	// ToProto marshals BgpPrefixIpv4UnicastState to protobuf object *otg.BgpPrefixIpv4UnicastState
	ToProto() (*otg.BgpPrefixIpv4UnicastState, error)
	// ToPbText marshals BgpPrefixIpv4UnicastState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpPrefixIpv4UnicastState to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpPrefixIpv4UnicastState to JSON text
	ToJson() (string, error)
}

type unMarshalbgpPrefixIpv4UnicastState struct {
	obj *bgpPrefixIpv4UnicastState
}

type unMarshalBgpPrefixIpv4UnicastState interface {
	// FromProto unmarshals BgpPrefixIpv4UnicastState from protobuf object *otg.BgpPrefixIpv4UnicastState
	FromProto(msg *otg.BgpPrefixIpv4UnicastState) (BgpPrefixIpv4UnicastState, error)
	// FromPbText unmarshals BgpPrefixIpv4UnicastState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpPrefixIpv4UnicastState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpPrefixIpv4UnicastState from JSON text
	FromJson(value string) error
}

func (obj *bgpPrefixIpv4UnicastState) Marshal() marshalBgpPrefixIpv4UnicastState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpPrefixIpv4UnicastState{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpPrefixIpv4UnicastState) Unmarshal() unMarshalBgpPrefixIpv4UnicastState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpPrefixIpv4UnicastState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpPrefixIpv4UnicastState) ToProto() (*otg.BgpPrefixIpv4UnicastState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpPrefixIpv4UnicastState) FromProto(msg *otg.BgpPrefixIpv4UnicastState) (BgpPrefixIpv4UnicastState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpPrefixIpv4UnicastState) ToPbText() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4UnicastState) FromPbText(value string) error {
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

func (m *marshalbgpPrefixIpv4UnicastState) ToYaml() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4UnicastState) FromYaml(value string) error {
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

func (m *marshalbgpPrefixIpv4UnicastState) ToJson() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4UnicastState) FromJson(value string) error {
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

func (obj *bgpPrefixIpv4UnicastState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv4UnicastState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv4UnicastState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpPrefixIpv4UnicastState) Clone() (BgpPrefixIpv4UnicastState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpPrefixIpv4UnicastState()
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

func (obj *bgpPrefixIpv4UnicastState) setNil() {
	obj.communitiesHolder = nil
	obj.extendedCommunitiesHolder = nil
	obj.asPathHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpPrefixIpv4UnicastState is iPv4 unicast prefix.
type BgpPrefixIpv4UnicastState interface {
	Validation
	// msg marshals BgpPrefixIpv4UnicastState to protobuf object *otg.BgpPrefixIpv4UnicastState
	// and doesn't set defaults
	msg() *otg.BgpPrefixIpv4UnicastState
	// setMsg unmarshals BgpPrefixIpv4UnicastState from protobuf object *otg.BgpPrefixIpv4UnicastState
	// and doesn't set defaults
	setMsg(*otg.BgpPrefixIpv4UnicastState) BgpPrefixIpv4UnicastState
	// provides marshal interface
	Marshal() marshalBgpPrefixIpv4UnicastState
	// provides unmarshal interface
	Unmarshal() unMarshalBgpPrefixIpv4UnicastState
	// validate validates BgpPrefixIpv4UnicastState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpPrefixIpv4UnicastState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4Address returns string, set in BgpPrefixIpv4UnicastState.
	Ipv4Address() string
	// SetIpv4Address assigns string provided by user to BgpPrefixIpv4UnicastState
	SetIpv4Address(value string) BgpPrefixIpv4UnicastState
	// HasIpv4Address checks if Ipv4Address has been set in BgpPrefixIpv4UnicastState
	HasIpv4Address() bool
	// PrefixLength returns uint32, set in BgpPrefixIpv4UnicastState.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to BgpPrefixIpv4UnicastState
	SetPrefixLength(value uint32) BgpPrefixIpv4UnicastState
	// HasPrefixLength checks if PrefixLength has been set in BgpPrefixIpv4UnicastState
	HasPrefixLength() bool
	// Origin returns BgpPrefixIpv4UnicastStateOriginEnum, set in BgpPrefixIpv4UnicastState
	Origin() BgpPrefixIpv4UnicastStateOriginEnum
	// SetOrigin assigns BgpPrefixIpv4UnicastStateOriginEnum provided by user to BgpPrefixIpv4UnicastState
	SetOrigin(value BgpPrefixIpv4UnicastStateOriginEnum) BgpPrefixIpv4UnicastState
	// HasOrigin checks if Origin has been set in BgpPrefixIpv4UnicastState
	HasOrigin() bool
	// PathId returns uint32, set in BgpPrefixIpv4UnicastState.
	PathId() uint32
	// SetPathId assigns uint32 provided by user to BgpPrefixIpv4UnicastState
	SetPathId(value uint32) BgpPrefixIpv4UnicastState
	// HasPathId checks if PathId has been set in BgpPrefixIpv4UnicastState
	HasPathId() bool
	// Ipv4NextHop returns string, set in BgpPrefixIpv4UnicastState.
	Ipv4NextHop() string
	// SetIpv4NextHop assigns string provided by user to BgpPrefixIpv4UnicastState
	SetIpv4NextHop(value string) BgpPrefixIpv4UnicastState
	// HasIpv4NextHop checks if Ipv4NextHop has been set in BgpPrefixIpv4UnicastState
	HasIpv4NextHop() bool
	// Ipv6NextHop returns string, set in BgpPrefixIpv4UnicastState.
	Ipv6NextHop() string
	// SetIpv6NextHop assigns string provided by user to BgpPrefixIpv4UnicastState
	SetIpv6NextHop(value string) BgpPrefixIpv4UnicastState
	// HasIpv6NextHop checks if Ipv6NextHop has been set in BgpPrefixIpv4UnicastState
	HasIpv6NextHop() bool
	// Communities returns BgpPrefixIpv4UnicastStateResultBgpCommunityIterIter, set in BgpPrefixIpv4UnicastState
	Communities() BgpPrefixIpv4UnicastStateResultBgpCommunityIter
	// ExtendedCommunities returns BgpPrefixIpv4UnicastStateResultExtendedCommunityIterIter, set in BgpPrefixIpv4UnicastState
	ExtendedCommunities() BgpPrefixIpv4UnicastStateResultExtendedCommunityIter
	// AsPath returns ResultBgpAsPath, set in BgpPrefixIpv4UnicastState.
	// ResultBgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed.
	AsPath() ResultBgpAsPath
	// SetAsPath assigns ResultBgpAsPath provided by user to BgpPrefixIpv4UnicastState.
	// ResultBgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed.
	SetAsPath(value ResultBgpAsPath) BgpPrefixIpv4UnicastState
	// HasAsPath checks if AsPath has been set in BgpPrefixIpv4UnicastState
	HasAsPath() bool
	// LocalPreference returns uint32, set in BgpPrefixIpv4UnicastState.
	LocalPreference() uint32
	// SetLocalPreference assigns uint32 provided by user to BgpPrefixIpv4UnicastState
	SetLocalPreference(value uint32) BgpPrefixIpv4UnicastState
	// HasLocalPreference checks if LocalPreference has been set in BgpPrefixIpv4UnicastState
	HasLocalPreference() bool
	// MultiExitDiscriminator returns uint32, set in BgpPrefixIpv4UnicastState.
	MultiExitDiscriminator() uint32
	// SetMultiExitDiscriminator assigns uint32 provided by user to BgpPrefixIpv4UnicastState
	SetMultiExitDiscriminator(value uint32) BgpPrefixIpv4UnicastState
	// HasMultiExitDiscriminator checks if MultiExitDiscriminator has been set in BgpPrefixIpv4UnicastState
	HasMultiExitDiscriminator() bool
	setNil()
}

// An IPv4 unicast address
// Ipv4Address returns a string
func (obj *bgpPrefixIpv4UnicastState) Ipv4Address() string {

	return *obj.obj.Ipv4Address

}

// An IPv4 unicast address
// Ipv4Address returns a string
func (obj *bgpPrefixIpv4UnicastState) HasIpv4Address() bool {
	return obj.obj.Ipv4Address != nil
}

// An IPv4 unicast address
// SetIpv4Address sets the string value in the BgpPrefixIpv4UnicastState object
func (obj *bgpPrefixIpv4UnicastState) SetIpv4Address(value string) BgpPrefixIpv4UnicastState {

	obj.obj.Ipv4Address = &value
	return obj
}

// The length of the prefix.
// PrefixLength returns a uint32
func (obj *bgpPrefixIpv4UnicastState) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// The length of the prefix.
// PrefixLength returns a uint32
func (obj *bgpPrefixIpv4UnicastState) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// The length of the prefix.
// SetPrefixLength sets the uint32 value in the BgpPrefixIpv4UnicastState object
func (obj *bgpPrefixIpv4UnicastState) SetPrefixLength(value uint32) BgpPrefixIpv4UnicastState {

	obj.obj.PrefixLength = &value
	return obj
}

type BgpPrefixIpv4UnicastStateOriginEnum string

// Enum of Origin on BgpPrefixIpv4UnicastState
var BgpPrefixIpv4UnicastStateOrigin = struct {
	IGP        BgpPrefixIpv4UnicastStateOriginEnum
	EGP        BgpPrefixIpv4UnicastStateOriginEnum
	INCOMPLETE BgpPrefixIpv4UnicastStateOriginEnum
}{
	IGP:        BgpPrefixIpv4UnicastStateOriginEnum("igp"),
	EGP:        BgpPrefixIpv4UnicastStateOriginEnum("egp"),
	INCOMPLETE: BgpPrefixIpv4UnicastStateOriginEnum("incomplete"),
}

func (obj *bgpPrefixIpv4UnicastState) Origin() BgpPrefixIpv4UnicastStateOriginEnum {
	return BgpPrefixIpv4UnicastStateOriginEnum(obj.obj.Origin.Enum().String())
}

// The origin of the prefix.
// Origin returns a string
func (obj *bgpPrefixIpv4UnicastState) HasOrigin() bool {
	return obj.obj.Origin != nil
}

func (obj *bgpPrefixIpv4UnicastState) SetOrigin(value BgpPrefixIpv4UnicastStateOriginEnum) BgpPrefixIpv4UnicastState {
	intValue, ok := otg.BgpPrefixIpv4UnicastState_Origin_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpPrefixIpv4UnicastStateOriginEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpPrefixIpv4UnicastState_Origin_Enum(intValue)
	obj.obj.Origin = &enumValue

	return obj
}

// The path id.
// PathId returns a uint32
func (obj *bgpPrefixIpv4UnicastState) PathId() uint32 {

	return *obj.obj.PathId

}

// The path id.
// PathId returns a uint32
func (obj *bgpPrefixIpv4UnicastState) HasPathId() bool {
	return obj.obj.PathId != nil
}

// The path id.
// SetPathId sets the uint32 value in the BgpPrefixIpv4UnicastState object
func (obj *bgpPrefixIpv4UnicastState) SetPathId(value uint32) BgpPrefixIpv4UnicastState {

	obj.obj.PathId = &value
	return obj
}

// The IPv4 address of the egress interface.
// Ipv4NextHop returns a string
func (obj *bgpPrefixIpv4UnicastState) Ipv4NextHop() string {

	return *obj.obj.Ipv4NextHop

}

// The IPv4 address of the egress interface.
// Ipv4NextHop returns a string
func (obj *bgpPrefixIpv4UnicastState) HasIpv4NextHop() bool {
	return obj.obj.Ipv4NextHop != nil
}

// The IPv4 address of the egress interface.
// SetIpv4NextHop sets the string value in the BgpPrefixIpv4UnicastState object
func (obj *bgpPrefixIpv4UnicastState) SetIpv4NextHop(value string) BgpPrefixIpv4UnicastState {

	obj.obj.Ipv4NextHop = &value
	return obj
}

// The IPv6 address of the egress interface.
// Ipv6NextHop returns a string
func (obj *bgpPrefixIpv4UnicastState) Ipv6NextHop() string {

	return *obj.obj.Ipv6NextHop

}

// The IPv6 address of the egress interface.
// Ipv6NextHop returns a string
func (obj *bgpPrefixIpv4UnicastState) HasIpv6NextHop() bool {
	return obj.obj.Ipv6NextHop != nil
}

// The IPv6 address of the egress interface.
// SetIpv6NextHop sets the string value in the BgpPrefixIpv4UnicastState object
func (obj *bgpPrefixIpv4UnicastState) SetIpv6NextHop(value string) BgpPrefixIpv4UnicastState {

	obj.obj.Ipv6NextHop = &value
	return obj
}

// Optional community attributes.
// Communities returns a []ResultBgpCommunity
func (obj *bgpPrefixIpv4UnicastState) Communities() BgpPrefixIpv4UnicastStateResultBgpCommunityIter {
	if len(obj.obj.Communities) == 0 {
		obj.obj.Communities = []*otg.ResultBgpCommunity{}
	}
	if obj.communitiesHolder == nil {
		obj.communitiesHolder = newBgpPrefixIpv4UnicastStateResultBgpCommunityIter(&obj.obj.Communities).setMsg(obj)
	}
	return obj.communitiesHolder
}

type bgpPrefixIpv4UnicastStateResultBgpCommunityIter struct {
	obj                     *bgpPrefixIpv4UnicastState
	resultBgpCommunitySlice []ResultBgpCommunity
	fieldPtr                *[]*otg.ResultBgpCommunity
}

func newBgpPrefixIpv4UnicastStateResultBgpCommunityIter(ptr *[]*otg.ResultBgpCommunity) BgpPrefixIpv4UnicastStateResultBgpCommunityIter {
	return &bgpPrefixIpv4UnicastStateResultBgpCommunityIter{fieldPtr: ptr}
}

type BgpPrefixIpv4UnicastStateResultBgpCommunityIter interface {
	setMsg(*bgpPrefixIpv4UnicastState) BgpPrefixIpv4UnicastStateResultBgpCommunityIter
	Items() []ResultBgpCommunity
	Add() ResultBgpCommunity
	Append(items ...ResultBgpCommunity) BgpPrefixIpv4UnicastStateResultBgpCommunityIter
	Set(index int, newObj ResultBgpCommunity) BgpPrefixIpv4UnicastStateResultBgpCommunityIter
	Clear() BgpPrefixIpv4UnicastStateResultBgpCommunityIter
	clearHolderSlice() BgpPrefixIpv4UnicastStateResultBgpCommunityIter
	appendHolderSlice(item ResultBgpCommunity) BgpPrefixIpv4UnicastStateResultBgpCommunityIter
}

func (obj *bgpPrefixIpv4UnicastStateResultBgpCommunityIter) setMsg(msg *bgpPrefixIpv4UnicastState) BgpPrefixIpv4UnicastStateResultBgpCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&resultBgpCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixIpv4UnicastStateResultBgpCommunityIter) Items() []ResultBgpCommunity {
	return obj.resultBgpCommunitySlice
}

func (obj *bgpPrefixIpv4UnicastStateResultBgpCommunityIter) Add() ResultBgpCommunity {
	newObj := &otg.ResultBgpCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &resultBgpCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixIpv4UnicastStateResultBgpCommunityIter) Append(items ...ResultBgpCommunity) BgpPrefixIpv4UnicastStateResultBgpCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, item)
	}
	return obj
}

func (obj *bgpPrefixIpv4UnicastStateResultBgpCommunityIter) Set(index int, newObj ResultBgpCommunity) BgpPrefixIpv4UnicastStateResultBgpCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.resultBgpCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpPrefixIpv4UnicastStateResultBgpCommunityIter) Clear() BgpPrefixIpv4UnicastStateResultBgpCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ResultBgpCommunity{}
		obj.resultBgpCommunitySlice = []ResultBgpCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv4UnicastStateResultBgpCommunityIter) clearHolderSlice() BgpPrefixIpv4UnicastStateResultBgpCommunityIter {
	if len(obj.resultBgpCommunitySlice) > 0 {
		obj.resultBgpCommunitySlice = []ResultBgpCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv4UnicastStateResultBgpCommunityIter) appendHolderSlice(item ResultBgpCommunity) BgpPrefixIpv4UnicastStateResultBgpCommunityIter {
	obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, item)
	return obj
}

// Optional received Extended Community attributes. Each received Extended Community attribute is available for retrieval in two forms. Support of the 'raw' format in which all 8 bytes (16 hex characters) is always present and available for use. In addition, if supported by the implementation, the Extended Community attribute may also be retrieved in the  'structured' format which is an optional field.
// ExtendedCommunities returns a []ResultExtendedCommunity
func (obj *bgpPrefixIpv4UnicastState) ExtendedCommunities() BgpPrefixIpv4UnicastStateResultExtendedCommunityIter {
	if len(obj.obj.ExtendedCommunities) == 0 {
		obj.obj.ExtendedCommunities = []*otg.ResultExtendedCommunity{}
	}
	if obj.extendedCommunitiesHolder == nil {
		obj.extendedCommunitiesHolder = newBgpPrefixIpv4UnicastStateResultExtendedCommunityIter(&obj.obj.ExtendedCommunities).setMsg(obj)
	}
	return obj.extendedCommunitiesHolder
}

type bgpPrefixIpv4UnicastStateResultExtendedCommunityIter struct {
	obj                          *bgpPrefixIpv4UnicastState
	resultExtendedCommunitySlice []ResultExtendedCommunity
	fieldPtr                     *[]*otg.ResultExtendedCommunity
}

func newBgpPrefixIpv4UnicastStateResultExtendedCommunityIter(ptr *[]*otg.ResultExtendedCommunity) BgpPrefixIpv4UnicastStateResultExtendedCommunityIter {
	return &bgpPrefixIpv4UnicastStateResultExtendedCommunityIter{fieldPtr: ptr}
}

type BgpPrefixIpv4UnicastStateResultExtendedCommunityIter interface {
	setMsg(*bgpPrefixIpv4UnicastState) BgpPrefixIpv4UnicastStateResultExtendedCommunityIter
	Items() []ResultExtendedCommunity
	Add() ResultExtendedCommunity
	Append(items ...ResultExtendedCommunity) BgpPrefixIpv4UnicastStateResultExtendedCommunityIter
	Set(index int, newObj ResultExtendedCommunity) BgpPrefixIpv4UnicastStateResultExtendedCommunityIter
	Clear() BgpPrefixIpv4UnicastStateResultExtendedCommunityIter
	clearHolderSlice() BgpPrefixIpv4UnicastStateResultExtendedCommunityIter
	appendHolderSlice(item ResultExtendedCommunity) BgpPrefixIpv4UnicastStateResultExtendedCommunityIter
}

func (obj *bgpPrefixIpv4UnicastStateResultExtendedCommunityIter) setMsg(msg *bgpPrefixIpv4UnicastState) BgpPrefixIpv4UnicastStateResultExtendedCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&resultExtendedCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixIpv4UnicastStateResultExtendedCommunityIter) Items() []ResultExtendedCommunity {
	return obj.resultExtendedCommunitySlice
}

func (obj *bgpPrefixIpv4UnicastStateResultExtendedCommunityIter) Add() ResultExtendedCommunity {
	newObj := &otg.ResultExtendedCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &resultExtendedCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixIpv4UnicastStateResultExtendedCommunityIter) Append(items ...ResultExtendedCommunity) BgpPrefixIpv4UnicastStateResultExtendedCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, item)
	}
	return obj
}

func (obj *bgpPrefixIpv4UnicastStateResultExtendedCommunityIter) Set(index int, newObj ResultExtendedCommunity) BgpPrefixIpv4UnicastStateResultExtendedCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.resultExtendedCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpPrefixIpv4UnicastStateResultExtendedCommunityIter) Clear() BgpPrefixIpv4UnicastStateResultExtendedCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ResultExtendedCommunity{}
		obj.resultExtendedCommunitySlice = []ResultExtendedCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv4UnicastStateResultExtendedCommunityIter) clearHolderSlice() BgpPrefixIpv4UnicastStateResultExtendedCommunityIter {
	if len(obj.resultExtendedCommunitySlice) > 0 {
		obj.resultExtendedCommunitySlice = []ResultExtendedCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv4UnicastStateResultExtendedCommunityIter) appendHolderSlice(item ResultExtendedCommunity) BgpPrefixIpv4UnicastStateResultExtendedCommunityIter {
	obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, item)
	return obj
}

// description is TBD
// AsPath returns a ResultBgpAsPath
func (obj *bgpPrefixIpv4UnicastState) AsPath() ResultBgpAsPath {
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
func (obj *bgpPrefixIpv4UnicastState) HasAsPath() bool {
	return obj.obj.AsPath != nil
}

// description is TBD
// SetAsPath sets the ResultBgpAsPath value in the BgpPrefixIpv4UnicastState object
func (obj *bgpPrefixIpv4UnicastState) SetAsPath(value ResultBgpAsPath) BgpPrefixIpv4UnicastState {

	obj.asPathHolder = nil
	obj.obj.AsPath = value.msg()

	return obj
}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// LocalPreference returns a uint32
func (obj *bgpPrefixIpv4UnicastState) LocalPreference() uint32 {

	return *obj.obj.LocalPreference

}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// LocalPreference returns a uint32
func (obj *bgpPrefixIpv4UnicastState) HasLocalPreference() bool {
	return obj.obj.LocalPreference != nil
}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// SetLocalPreference sets the uint32 value in the BgpPrefixIpv4UnicastState object
func (obj *bgpPrefixIpv4UnicastState) SetLocalPreference(value uint32) BgpPrefixIpv4UnicastState {

	obj.obj.LocalPreference = &value
	return obj
}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// MultiExitDiscriminator returns a uint32
func (obj *bgpPrefixIpv4UnicastState) MultiExitDiscriminator() uint32 {

	return *obj.obj.MultiExitDiscriminator

}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// MultiExitDiscriminator returns a uint32
func (obj *bgpPrefixIpv4UnicastState) HasMultiExitDiscriminator() bool {
	return obj.obj.MultiExitDiscriminator != nil
}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// SetMultiExitDiscriminator sets the uint32 value in the BgpPrefixIpv4UnicastState object
func (obj *bgpPrefixIpv4UnicastState) SetMultiExitDiscriminator(value uint32) BgpPrefixIpv4UnicastState {

	obj.obj.MultiExitDiscriminator = &value
	return obj
}

func (obj *bgpPrefixIpv4UnicastState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpPrefixIpv4UnicastState.PrefixLength <= 128 but Got %d", *obj.obj.PrefixLength))
		}

	}

	if obj.obj.Ipv4NextHop != nil {

		err := obj.validateIpv4(obj.Ipv4NextHop())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv4UnicastState.Ipv4NextHop"))
		}

	}

	if obj.obj.Ipv6NextHop != nil {

		err := obj.validateIpv6(obj.Ipv6NextHop())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv4UnicastState.Ipv6NextHop"))
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

func (obj *bgpPrefixIpv4UnicastState) setDefault() {

}
