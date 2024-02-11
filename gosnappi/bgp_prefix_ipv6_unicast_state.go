package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpPrefixIpv6UnicastState *****
type bgpPrefixIpv6UnicastState struct {
	validation
	obj               *otg.BgpPrefixIpv6UnicastState
	marshaller        marshalBgpPrefixIpv6UnicastState
	unMarshaller      unMarshalBgpPrefixIpv6UnicastState
	communitiesHolder BgpPrefixIpv6UnicastStateResultBgpCommunityIter
	asPathHolder      ResultBgpAsPath
}

func NewBgpPrefixIpv6UnicastState() BgpPrefixIpv6UnicastState {
	obj := bgpPrefixIpv6UnicastState{obj: &otg.BgpPrefixIpv6UnicastState{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpPrefixIpv6UnicastState) msg() *otg.BgpPrefixIpv6UnicastState {
	return obj.obj
}

func (obj *bgpPrefixIpv6UnicastState) setMsg(msg *otg.BgpPrefixIpv6UnicastState) BgpPrefixIpv6UnicastState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpPrefixIpv6UnicastState struct {
	obj *bgpPrefixIpv6UnicastState
}

type marshalBgpPrefixIpv6UnicastState interface {
	// ToProto marshals BgpPrefixIpv6UnicastState to protobuf object *otg.BgpPrefixIpv6UnicastState
	ToProto() (*otg.BgpPrefixIpv6UnicastState, error)
	// ToPbText marshals BgpPrefixIpv6UnicastState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpPrefixIpv6UnicastState to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpPrefixIpv6UnicastState to JSON text
	ToJson() (string, error)
}

type unMarshalbgpPrefixIpv6UnicastState struct {
	obj *bgpPrefixIpv6UnicastState
}

type unMarshalBgpPrefixIpv6UnicastState interface {
	// FromProto unmarshals BgpPrefixIpv6UnicastState from protobuf object *otg.BgpPrefixIpv6UnicastState
	FromProto(msg *otg.BgpPrefixIpv6UnicastState) (BgpPrefixIpv6UnicastState, error)
	// FromPbText unmarshals BgpPrefixIpv6UnicastState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpPrefixIpv6UnicastState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpPrefixIpv6UnicastState from JSON text
	FromJson(value string) error
}

func (obj *bgpPrefixIpv6UnicastState) Marshal() marshalBgpPrefixIpv6UnicastState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpPrefixIpv6UnicastState{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpPrefixIpv6UnicastState) Unmarshal() unMarshalBgpPrefixIpv6UnicastState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpPrefixIpv6UnicastState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpPrefixIpv6UnicastState) ToProto() (*otg.BgpPrefixIpv6UnicastState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpPrefixIpv6UnicastState) FromProto(msg *otg.BgpPrefixIpv6UnicastState) (BgpPrefixIpv6UnicastState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpPrefixIpv6UnicastState) ToPbText() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6UnicastState) FromPbText(value string) error {
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

func (m *marshalbgpPrefixIpv6UnicastState) ToYaml() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6UnicastState) FromYaml(value string) error {
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

func (m *marshalbgpPrefixIpv6UnicastState) ToJson() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6UnicastState) FromJson(value string) error {
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

func (obj *bgpPrefixIpv6UnicastState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv6UnicastState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv6UnicastState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpPrefixIpv6UnicastState) Clone() (BgpPrefixIpv6UnicastState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpPrefixIpv6UnicastState()
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

func (obj *bgpPrefixIpv6UnicastState) setNil() {
	obj.communitiesHolder = nil
	obj.asPathHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpPrefixIpv6UnicastState is iPv6 unicast prefix.
type BgpPrefixIpv6UnicastState interface {
	Validation
	// msg marshals BgpPrefixIpv6UnicastState to protobuf object *otg.BgpPrefixIpv6UnicastState
	// and doesn't set defaults
	msg() *otg.BgpPrefixIpv6UnicastState
	// setMsg unmarshals BgpPrefixIpv6UnicastState from protobuf object *otg.BgpPrefixIpv6UnicastState
	// and doesn't set defaults
	setMsg(*otg.BgpPrefixIpv6UnicastState) BgpPrefixIpv6UnicastState
	// provides marshal interface
	Marshal() marshalBgpPrefixIpv6UnicastState
	// provides unmarshal interface
	Unmarshal() unMarshalBgpPrefixIpv6UnicastState
	// validate validates BgpPrefixIpv6UnicastState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpPrefixIpv6UnicastState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv6Address returns string, set in BgpPrefixIpv6UnicastState.
	Ipv6Address() string
	// SetIpv6Address assigns string provided by user to BgpPrefixIpv6UnicastState
	SetIpv6Address(value string) BgpPrefixIpv6UnicastState
	// HasIpv6Address checks if Ipv6Address has been set in BgpPrefixIpv6UnicastState
	HasIpv6Address() bool
	// PrefixLength returns uint32, set in BgpPrefixIpv6UnicastState.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to BgpPrefixIpv6UnicastState
	SetPrefixLength(value uint32) BgpPrefixIpv6UnicastState
	// HasPrefixLength checks if PrefixLength has been set in BgpPrefixIpv6UnicastState
	HasPrefixLength() bool
	// Origin returns BgpPrefixIpv6UnicastStateOriginEnum, set in BgpPrefixIpv6UnicastState
	Origin() BgpPrefixIpv6UnicastStateOriginEnum
	// SetOrigin assigns BgpPrefixIpv6UnicastStateOriginEnum provided by user to BgpPrefixIpv6UnicastState
	SetOrigin(value BgpPrefixIpv6UnicastStateOriginEnum) BgpPrefixIpv6UnicastState
	// HasOrigin checks if Origin has been set in BgpPrefixIpv6UnicastState
	HasOrigin() bool
	// PathId returns uint32, set in BgpPrefixIpv6UnicastState.
	PathId() uint32
	// SetPathId assigns uint32 provided by user to BgpPrefixIpv6UnicastState
	SetPathId(value uint32) BgpPrefixIpv6UnicastState
	// HasPathId checks if PathId has been set in BgpPrefixIpv6UnicastState
	HasPathId() bool
	// Ipv4NextHop returns string, set in BgpPrefixIpv6UnicastState.
	Ipv4NextHop() string
	// SetIpv4NextHop assigns string provided by user to BgpPrefixIpv6UnicastState
	SetIpv4NextHop(value string) BgpPrefixIpv6UnicastState
	// HasIpv4NextHop checks if Ipv4NextHop has been set in BgpPrefixIpv6UnicastState
	HasIpv4NextHop() bool
	// Ipv6NextHop returns string, set in BgpPrefixIpv6UnicastState.
	Ipv6NextHop() string
	// SetIpv6NextHop assigns string provided by user to BgpPrefixIpv6UnicastState
	SetIpv6NextHop(value string) BgpPrefixIpv6UnicastState
	// HasIpv6NextHop checks if Ipv6NextHop has been set in BgpPrefixIpv6UnicastState
	HasIpv6NextHop() bool
	// Communities returns BgpPrefixIpv6UnicastStateResultBgpCommunityIterIter, set in BgpPrefixIpv6UnicastState
	Communities() BgpPrefixIpv6UnicastStateResultBgpCommunityIter
	// AsPath returns ResultBgpAsPath, set in BgpPrefixIpv6UnicastState.
	// ResultBgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed.
	AsPath() ResultBgpAsPath
	// SetAsPath assigns ResultBgpAsPath provided by user to BgpPrefixIpv6UnicastState.
	// ResultBgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed.
	SetAsPath(value ResultBgpAsPath) BgpPrefixIpv6UnicastState
	// HasAsPath checks if AsPath has been set in BgpPrefixIpv6UnicastState
	HasAsPath() bool
	// LocalPreference returns uint32, set in BgpPrefixIpv6UnicastState.
	LocalPreference() uint32
	// SetLocalPreference assigns uint32 provided by user to BgpPrefixIpv6UnicastState
	SetLocalPreference(value uint32) BgpPrefixIpv6UnicastState
	// HasLocalPreference checks if LocalPreference has been set in BgpPrefixIpv6UnicastState
	HasLocalPreference() bool
	// MultiExitDiscriminator returns uint32, set in BgpPrefixIpv6UnicastState.
	MultiExitDiscriminator() uint32
	// SetMultiExitDiscriminator assigns uint32 provided by user to BgpPrefixIpv6UnicastState
	SetMultiExitDiscriminator(value uint32) BgpPrefixIpv6UnicastState
	// HasMultiExitDiscriminator checks if MultiExitDiscriminator has been set in BgpPrefixIpv6UnicastState
	HasMultiExitDiscriminator() bool
	setNil()
}

// An IPv6 unicast address
// Ipv6Address returns a string
func (obj *bgpPrefixIpv6UnicastState) Ipv6Address() string {

	return *obj.obj.Ipv6Address

}

// An IPv6 unicast address
// Ipv6Address returns a string
func (obj *bgpPrefixIpv6UnicastState) HasIpv6Address() bool {
	return obj.obj.Ipv6Address != nil
}

// An IPv6 unicast address
// SetIpv6Address sets the string value in the BgpPrefixIpv6UnicastState object
func (obj *bgpPrefixIpv6UnicastState) SetIpv6Address(value string) BgpPrefixIpv6UnicastState {

	obj.obj.Ipv6Address = &value
	return obj
}

// The length of the prefix.
// PrefixLength returns a uint32
func (obj *bgpPrefixIpv6UnicastState) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// The length of the prefix.
// PrefixLength returns a uint32
func (obj *bgpPrefixIpv6UnicastState) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// The length of the prefix.
// SetPrefixLength sets the uint32 value in the BgpPrefixIpv6UnicastState object
func (obj *bgpPrefixIpv6UnicastState) SetPrefixLength(value uint32) BgpPrefixIpv6UnicastState {

	obj.obj.PrefixLength = &value
	return obj
}

type BgpPrefixIpv6UnicastStateOriginEnum string

// Enum of Origin on BgpPrefixIpv6UnicastState
var BgpPrefixIpv6UnicastStateOrigin = struct {
	IGP        BgpPrefixIpv6UnicastStateOriginEnum
	EGP        BgpPrefixIpv6UnicastStateOriginEnum
	INCOMPLETE BgpPrefixIpv6UnicastStateOriginEnum
}{
	IGP:        BgpPrefixIpv6UnicastStateOriginEnum("igp"),
	EGP:        BgpPrefixIpv6UnicastStateOriginEnum("egp"),
	INCOMPLETE: BgpPrefixIpv6UnicastStateOriginEnum("incomplete"),
}

func (obj *bgpPrefixIpv6UnicastState) Origin() BgpPrefixIpv6UnicastStateOriginEnum {
	return BgpPrefixIpv6UnicastStateOriginEnum(obj.obj.Origin.Enum().String())
}

// The origin of the prefix.
// Origin returns a string
func (obj *bgpPrefixIpv6UnicastState) HasOrigin() bool {
	return obj.obj.Origin != nil
}

func (obj *bgpPrefixIpv6UnicastState) SetOrigin(value BgpPrefixIpv6UnicastStateOriginEnum) BgpPrefixIpv6UnicastState {
	intValue, ok := otg.BgpPrefixIpv6UnicastState_Origin_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpPrefixIpv6UnicastStateOriginEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpPrefixIpv6UnicastState_Origin_Enum(intValue)
	obj.obj.Origin = &enumValue

	return obj
}

// The path id.
// PathId returns a uint32
func (obj *bgpPrefixIpv6UnicastState) PathId() uint32 {

	return *obj.obj.PathId

}

// The path id.
// PathId returns a uint32
func (obj *bgpPrefixIpv6UnicastState) HasPathId() bool {
	return obj.obj.PathId != nil
}

// The path id.
// SetPathId sets the uint32 value in the BgpPrefixIpv6UnicastState object
func (obj *bgpPrefixIpv6UnicastState) SetPathId(value uint32) BgpPrefixIpv6UnicastState {

	obj.obj.PathId = &value
	return obj
}

// The IPv4 address of the egress interface.
// Ipv4NextHop returns a string
func (obj *bgpPrefixIpv6UnicastState) Ipv4NextHop() string {

	return *obj.obj.Ipv4NextHop

}

// The IPv4 address of the egress interface.
// Ipv4NextHop returns a string
func (obj *bgpPrefixIpv6UnicastState) HasIpv4NextHop() bool {
	return obj.obj.Ipv4NextHop != nil
}

// The IPv4 address of the egress interface.
// SetIpv4NextHop sets the string value in the BgpPrefixIpv6UnicastState object
func (obj *bgpPrefixIpv6UnicastState) SetIpv4NextHop(value string) BgpPrefixIpv6UnicastState {

	obj.obj.Ipv4NextHop = &value
	return obj
}

// The IPv6 address of the egress interface.
// Ipv6NextHop returns a string
func (obj *bgpPrefixIpv6UnicastState) Ipv6NextHop() string {

	return *obj.obj.Ipv6NextHop

}

// The IPv6 address of the egress interface.
// Ipv6NextHop returns a string
func (obj *bgpPrefixIpv6UnicastState) HasIpv6NextHop() bool {
	return obj.obj.Ipv6NextHop != nil
}

// The IPv6 address of the egress interface.
// SetIpv6NextHop sets the string value in the BgpPrefixIpv6UnicastState object
func (obj *bgpPrefixIpv6UnicastState) SetIpv6NextHop(value string) BgpPrefixIpv6UnicastState {

	obj.obj.Ipv6NextHop = &value
	return obj
}

// Optional community attributes.
// Communities returns a []ResultBgpCommunity
func (obj *bgpPrefixIpv6UnicastState) Communities() BgpPrefixIpv6UnicastStateResultBgpCommunityIter {
	if len(obj.obj.Communities) == 0 {
		obj.obj.Communities = []*otg.ResultBgpCommunity{}
	}
	if obj.communitiesHolder == nil {
		obj.communitiesHolder = newBgpPrefixIpv6UnicastStateResultBgpCommunityIter(&obj.obj.Communities).setMsg(obj)
	}
	return obj.communitiesHolder
}

type bgpPrefixIpv6UnicastStateResultBgpCommunityIter struct {
	obj                     *bgpPrefixIpv6UnicastState
	resultBgpCommunitySlice []ResultBgpCommunity
	fieldPtr                *[]*otg.ResultBgpCommunity
}

func newBgpPrefixIpv6UnicastStateResultBgpCommunityIter(ptr *[]*otg.ResultBgpCommunity) BgpPrefixIpv6UnicastStateResultBgpCommunityIter {
	return &bgpPrefixIpv6UnicastStateResultBgpCommunityIter{fieldPtr: ptr}
}

type BgpPrefixIpv6UnicastStateResultBgpCommunityIter interface {
	setMsg(*bgpPrefixIpv6UnicastState) BgpPrefixIpv6UnicastStateResultBgpCommunityIter
	Items() []ResultBgpCommunity
	Add() ResultBgpCommunity
	Append(items ...ResultBgpCommunity) BgpPrefixIpv6UnicastStateResultBgpCommunityIter
	Set(index int, newObj ResultBgpCommunity) BgpPrefixIpv6UnicastStateResultBgpCommunityIter
	Clear() BgpPrefixIpv6UnicastStateResultBgpCommunityIter
	clearHolderSlice() BgpPrefixIpv6UnicastStateResultBgpCommunityIter
	appendHolderSlice(item ResultBgpCommunity) BgpPrefixIpv6UnicastStateResultBgpCommunityIter
}

func (obj *bgpPrefixIpv6UnicastStateResultBgpCommunityIter) setMsg(msg *bgpPrefixIpv6UnicastState) BgpPrefixIpv6UnicastStateResultBgpCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&resultBgpCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixIpv6UnicastStateResultBgpCommunityIter) Items() []ResultBgpCommunity {
	return obj.resultBgpCommunitySlice
}

func (obj *bgpPrefixIpv6UnicastStateResultBgpCommunityIter) Add() ResultBgpCommunity {
	newObj := &otg.ResultBgpCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &resultBgpCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixIpv6UnicastStateResultBgpCommunityIter) Append(items ...ResultBgpCommunity) BgpPrefixIpv6UnicastStateResultBgpCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, item)
	}
	return obj
}

func (obj *bgpPrefixIpv6UnicastStateResultBgpCommunityIter) Set(index int, newObj ResultBgpCommunity) BgpPrefixIpv6UnicastStateResultBgpCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.resultBgpCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpPrefixIpv6UnicastStateResultBgpCommunityIter) Clear() BgpPrefixIpv6UnicastStateResultBgpCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ResultBgpCommunity{}
		obj.resultBgpCommunitySlice = []ResultBgpCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv6UnicastStateResultBgpCommunityIter) clearHolderSlice() BgpPrefixIpv6UnicastStateResultBgpCommunityIter {
	if len(obj.resultBgpCommunitySlice) > 0 {
		obj.resultBgpCommunitySlice = []ResultBgpCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv6UnicastStateResultBgpCommunityIter) appendHolderSlice(item ResultBgpCommunity) BgpPrefixIpv6UnicastStateResultBgpCommunityIter {
	obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, item)
	return obj
}

// description is TBD
// AsPath returns a ResultBgpAsPath
func (obj *bgpPrefixIpv6UnicastState) AsPath() ResultBgpAsPath {
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
func (obj *bgpPrefixIpv6UnicastState) HasAsPath() bool {
	return obj.obj.AsPath != nil
}

// description is TBD
// SetAsPath sets the ResultBgpAsPath value in the BgpPrefixIpv6UnicastState object
func (obj *bgpPrefixIpv6UnicastState) SetAsPath(value ResultBgpAsPath) BgpPrefixIpv6UnicastState {

	obj.asPathHolder = nil
	obj.obj.AsPath = value.msg()

	return obj
}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// LocalPreference returns a uint32
func (obj *bgpPrefixIpv6UnicastState) LocalPreference() uint32 {

	return *obj.obj.LocalPreference

}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// LocalPreference returns a uint32
func (obj *bgpPrefixIpv6UnicastState) HasLocalPreference() bool {
	return obj.obj.LocalPreference != nil
}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// SetLocalPreference sets the uint32 value in the BgpPrefixIpv6UnicastState object
func (obj *bgpPrefixIpv6UnicastState) SetLocalPreference(value uint32) BgpPrefixIpv6UnicastState {

	obj.obj.LocalPreference = &value
	return obj
}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// MultiExitDiscriminator returns a uint32
func (obj *bgpPrefixIpv6UnicastState) MultiExitDiscriminator() uint32 {

	return *obj.obj.MultiExitDiscriminator

}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// MultiExitDiscriminator returns a uint32
func (obj *bgpPrefixIpv6UnicastState) HasMultiExitDiscriminator() bool {
	return obj.obj.MultiExitDiscriminator != nil
}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// SetMultiExitDiscriminator sets the uint32 value in the BgpPrefixIpv6UnicastState object
func (obj *bgpPrefixIpv6UnicastState) SetMultiExitDiscriminator(value uint32) BgpPrefixIpv6UnicastState {

	obj.obj.MultiExitDiscriminator = &value
	return obj
}

func (obj *bgpPrefixIpv6UnicastState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpPrefixIpv6UnicastState.PrefixLength <= 128 but Got %d", *obj.obj.PrefixLength))
		}

	}

	if obj.obj.Ipv4NextHop != nil {

		err := obj.validateIpv4(obj.Ipv4NextHop())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv6UnicastState.Ipv4NextHop"))
		}

	}

	if obj.obj.Ipv6NextHop != nil {

		err := obj.validateIpv6(obj.Ipv6NextHop())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv6UnicastState.Ipv6NextHop"))
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

	if obj.obj.AsPath != nil {

		obj.AsPath().validateObj(vObj, set_default)
	}

}

func (obj *bgpPrefixIpv6UnicastState) setDefault() {

}
