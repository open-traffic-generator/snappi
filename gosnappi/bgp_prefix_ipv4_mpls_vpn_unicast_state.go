package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpPrefixIpv4MplsVpnUnicastState *****
type bgpPrefixIpv4MplsVpnUnicastState struct {
	validation
	obj                       *otg.BgpPrefixIpv4MplsVpnUnicastState
	marshaller                marshalBgpPrefixIpv4MplsVpnUnicastState
	unMarshaller              unMarshalBgpPrefixIpv4MplsVpnUnicastState
	communitiesHolder         BgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter
	extendedCommunitiesHolder BgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter
	asPathHolder              ResultBgpAsPath
}

func NewBgpPrefixIpv4MplsVpnUnicastState() BgpPrefixIpv4MplsVpnUnicastState {
	obj := bgpPrefixIpv4MplsVpnUnicastState{obj: &otg.BgpPrefixIpv4MplsVpnUnicastState{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastState) msg() *otg.BgpPrefixIpv4MplsVpnUnicastState {
	return obj.obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastState) setMsg(msg *otg.BgpPrefixIpv4MplsVpnUnicastState) BgpPrefixIpv4MplsVpnUnicastState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpPrefixIpv4MplsVpnUnicastState struct {
	obj *bgpPrefixIpv4MplsVpnUnicastState
}

type marshalBgpPrefixIpv4MplsVpnUnicastState interface {
	// ToProto marshals BgpPrefixIpv4MplsVpnUnicastState to protobuf object *otg.BgpPrefixIpv4MplsVpnUnicastState
	ToProto() (*otg.BgpPrefixIpv4MplsVpnUnicastState, error)
	// ToPbText marshals BgpPrefixIpv4MplsVpnUnicastState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpPrefixIpv4MplsVpnUnicastState to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpPrefixIpv4MplsVpnUnicastState to JSON text
	ToJson() (string, error)
}

type unMarshalbgpPrefixIpv4MplsVpnUnicastState struct {
	obj *bgpPrefixIpv4MplsVpnUnicastState
}

type unMarshalBgpPrefixIpv4MplsVpnUnicastState interface {
	// FromProto unmarshals BgpPrefixIpv4MplsVpnUnicastState from protobuf object *otg.BgpPrefixIpv4MplsVpnUnicastState
	FromProto(msg *otg.BgpPrefixIpv4MplsVpnUnicastState) (BgpPrefixIpv4MplsVpnUnicastState, error)
	// FromPbText unmarshals BgpPrefixIpv4MplsVpnUnicastState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpPrefixIpv4MplsVpnUnicastState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpPrefixIpv4MplsVpnUnicastState from JSON text
	FromJson(value string) error
}

func (obj *bgpPrefixIpv4MplsVpnUnicastState) Marshal() marshalBgpPrefixIpv4MplsVpnUnicastState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpPrefixIpv4MplsVpnUnicastState{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpPrefixIpv4MplsVpnUnicastState) Unmarshal() unMarshalBgpPrefixIpv4MplsVpnUnicastState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpPrefixIpv4MplsVpnUnicastState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpPrefixIpv4MplsVpnUnicastState) ToProto() (*otg.BgpPrefixIpv4MplsVpnUnicastState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpPrefixIpv4MplsVpnUnicastState) FromProto(msg *otg.BgpPrefixIpv4MplsVpnUnicastState) (BgpPrefixIpv4MplsVpnUnicastState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpPrefixIpv4MplsVpnUnicastState) ToPbText() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4MplsVpnUnicastState) FromPbText(value string) error {
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

func (m *marshalbgpPrefixIpv4MplsVpnUnicastState) ToYaml() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4MplsVpnUnicastState) FromYaml(value string) error {
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

func (m *marshalbgpPrefixIpv4MplsVpnUnicastState) ToJson() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4MplsVpnUnicastState) FromJson(value string) error {
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

func (obj *bgpPrefixIpv4MplsVpnUnicastState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv4MplsVpnUnicastState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv4MplsVpnUnicastState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpPrefixIpv4MplsVpnUnicastState) Clone() (BgpPrefixIpv4MplsVpnUnicastState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpPrefixIpv4MplsVpnUnicastState()
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

func (obj *bgpPrefixIpv4MplsVpnUnicastState) setNil() {
	obj.communitiesHolder = nil
	obj.extendedCommunitiesHolder = nil
	obj.asPathHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpPrefixIpv4MplsVpnUnicastState is bGP/MPLS L3VPN (RFC 4364) VPN-IPv4 learned prefix, received under the VPN-IPv4 AFI/SAFI (AFI 1, SAFI 128).
type BgpPrefixIpv4MplsVpnUnicastState interface {
	Validation
	// msg marshals BgpPrefixIpv4MplsVpnUnicastState to protobuf object *otg.BgpPrefixIpv4MplsVpnUnicastState
	// and doesn't set defaults
	msg() *otg.BgpPrefixIpv4MplsVpnUnicastState
	// setMsg unmarshals BgpPrefixIpv4MplsVpnUnicastState from protobuf object *otg.BgpPrefixIpv4MplsVpnUnicastState
	// and doesn't set defaults
	setMsg(*otg.BgpPrefixIpv4MplsVpnUnicastState) BgpPrefixIpv4MplsVpnUnicastState
	// provides marshal interface
	Marshal() marshalBgpPrefixIpv4MplsVpnUnicastState
	// provides unmarshal interface
	Unmarshal() unMarshalBgpPrefixIpv4MplsVpnUnicastState
	// validate validates BgpPrefixIpv4MplsVpnUnicastState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpPrefixIpv4MplsVpnUnicastState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouteDistinguisher returns string, set in BgpPrefixIpv4MplsVpnUnicastState.
	RouteDistinguisher() string
	// SetRouteDistinguisher assigns string provided by user to BgpPrefixIpv4MplsVpnUnicastState
	SetRouteDistinguisher(value string) BgpPrefixIpv4MplsVpnUnicastState
	// HasRouteDistinguisher checks if RouteDistinguisher has been set in BgpPrefixIpv4MplsVpnUnicastState
	HasRouteDistinguisher() bool
	// Ipv4Address returns string, set in BgpPrefixIpv4MplsVpnUnicastState.
	Ipv4Address() string
	// SetIpv4Address assigns string provided by user to BgpPrefixIpv4MplsVpnUnicastState
	SetIpv4Address(value string) BgpPrefixIpv4MplsVpnUnicastState
	// HasIpv4Address checks if Ipv4Address has been set in BgpPrefixIpv4MplsVpnUnicastState
	HasIpv4Address() bool
	// PrefixLength returns uint32, set in BgpPrefixIpv4MplsVpnUnicastState.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to BgpPrefixIpv4MplsVpnUnicastState
	SetPrefixLength(value uint32) BgpPrefixIpv4MplsVpnUnicastState
	// HasPrefixLength checks if PrefixLength has been set in BgpPrefixIpv4MplsVpnUnicastState
	HasPrefixLength() bool
	// Origin returns BgpPrefixIpv4MplsVpnUnicastStateOriginEnum, set in BgpPrefixIpv4MplsVpnUnicastState
	Origin() BgpPrefixIpv4MplsVpnUnicastStateOriginEnum
	// SetOrigin assigns BgpPrefixIpv4MplsVpnUnicastStateOriginEnum provided by user to BgpPrefixIpv4MplsVpnUnicastState
	SetOrigin(value BgpPrefixIpv4MplsVpnUnicastStateOriginEnum) BgpPrefixIpv4MplsVpnUnicastState
	// HasOrigin checks if Origin has been set in BgpPrefixIpv4MplsVpnUnicastState
	HasOrigin() bool
	// PathId returns uint32, set in BgpPrefixIpv4MplsVpnUnicastState.
	PathId() uint32
	// SetPathId assigns uint32 provided by user to BgpPrefixIpv4MplsVpnUnicastState
	SetPathId(value uint32) BgpPrefixIpv4MplsVpnUnicastState
	// HasPathId checks if PathId has been set in BgpPrefixIpv4MplsVpnUnicastState
	HasPathId() bool
	// Ipv4NextHop returns string, set in BgpPrefixIpv4MplsVpnUnicastState.
	Ipv4NextHop() string
	// SetIpv4NextHop assigns string provided by user to BgpPrefixIpv4MplsVpnUnicastState
	SetIpv4NextHop(value string) BgpPrefixIpv4MplsVpnUnicastState
	// HasIpv4NextHop checks if Ipv4NextHop has been set in BgpPrefixIpv4MplsVpnUnicastState
	HasIpv4NextHop() bool
	// Ipv6NextHop returns string, set in BgpPrefixIpv4MplsVpnUnicastState.
	Ipv6NextHop() string
	// SetIpv6NextHop assigns string provided by user to BgpPrefixIpv4MplsVpnUnicastState
	SetIpv6NextHop(value string) BgpPrefixIpv4MplsVpnUnicastState
	// HasIpv6NextHop checks if Ipv6NextHop has been set in BgpPrefixIpv4MplsVpnUnicastState
	HasIpv6NextHop() bool
	// Labels returns []uint32, set in BgpPrefixIpv4MplsVpnUnicastState.
	Labels() []uint32
	// SetLabels assigns []uint32 provided by user to BgpPrefixIpv4MplsVpnUnicastState
	SetLabels(value []uint32) BgpPrefixIpv4MplsVpnUnicastState
	// Communities returns BgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIterIter, set in BgpPrefixIpv4MplsVpnUnicastState
	Communities() BgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter
	// ExtendedCommunities returns BgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIterIter, set in BgpPrefixIpv4MplsVpnUnicastState
	ExtendedCommunities() BgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter
	// AsPath returns ResultBgpAsPath, set in BgpPrefixIpv4MplsVpnUnicastState.
	// ResultBgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed.
	AsPath() ResultBgpAsPath
	// SetAsPath assigns ResultBgpAsPath provided by user to BgpPrefixIpv4MplsVpnUnicastState.
	// ResultBgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed.
	SetAsPath(value ResultBgpAsPath) BgpPrefixIpv4MplsVpnUnicastState
	// HasAsPath checks if AsPath has been set in BgpPrefixIpv4MplsVpnUnicastState
	HasAsPath() bool
	// LocalPreference returns uint32, set in BgpPrefixIpv4MplsVpnUnicastState.
	LocalPreference() uint32
	// SetLocalPreference assigns uint32 provided by user to BgpPrefixIpv4MplsVpnUnicastState
	SetLocalPreference(value uint32) BgpPrefixIpv4MplsVpnUnicastState
	// HasLocalPreference checks if LocalPreference has been set in BgpPrefixIpv4MplsVpnUnicastState
	HasLocalPreference() bool
	// MultiExitDiscriminator returns uint32, set in BgpPrefixIpv4MplsVpnUnicastState.
	MultiExitDiscriminator() uint32
	// SetMultiExitDiscriminator assigns uint32 provided by user to BgpPrefixIpv4MplsVpnUnicastState
	SetMultiExitDiscriminator(value uint32) BgpPrefixIpv4MplsVpnUnicastState
	// HasMultiExitDiscriminator checks if MultiExitDiscriminator has been set in BgpPrefixIpv4MplsVpnUnicastState
	HasMultiExitDiscriminator() bool
	setNil()
}

// The Route Distinguisher (RFC 4364 Section 4.1) received as part of the VPN-IPv4 NLRI, formatted as a colon separated value, for example "60005:100" or "1.1.1.1:100".
// RouteDistinguisher returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastState) RouteDistinguisher() string {

	return *obj.obj.RouteDistinguisher

}

// The Route Distinguisher (RFC 4364 Section 4.1) received as part of the VPN-IPv4 NLRI, formatted as a colon separated value, for example "60005:100" or "1.1.1.1:100".
// RouteDistinguisher returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastState) HasRouteDistinguisher() bool {
	return obj.obj.RouteDistinguisher != nil
}

// The Route Distinguisher (RFC 4364 Section 4.1) received as part of the VPN-IPv4 NLRI, formatted as a colon separated value, for example "60005:100" or "1.1.1.1:100".
// SetRouteDistinguisher sets the string value in the BgpPrefixIpv4MplsVpnUnicastState object
func (obj *bgpPrefixIpv4MplsVpnUnicastState) SetRouteDistinguisher(value string) BgpPrefixIpv4MplsVpnUnicastState {

	obj.obj.RouteDistinguisher = &value
	return obj
}

// An IPv4 unicast address.
// Ipv4Address returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastState) Ipv4Address() string {

	return *obj.obj.Ipv4Address

}

// An IPv4 unicast address.
// Ipv4Address returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastState) HasIpv4Address() bool {
	return obj.obj.Ipv4Address != nil
}

// An IPv4 unicast address.
// SetIpv4Address sets the string value in the BgpPrefixIpv4MplsVpnUnicastState object
func (obj *bgpPrefixIpv4MplsVpnUnicastState) SetIpv4Address(value string) BgpPrefixIpv4MplsVpnUnicastState {

	obj.obj.Ipv4Address = &value
	return obj
}

// description is TBD
// PrefixLength returns a uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastState) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// description is TBD
// PrefixLength returns a uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastState) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// description is TBD
// SetPrefixLength sets the uint32 value in the BgpPrefixIpv4MplsVpnUnicastState object
func (obj *bgpPrefixIpv4MplsVpnUnicastState) SetPrefixLength(value uint32) BgpPrefixIpv4MplsVpnUnicastState {

	obj.obj.PrefixLength = &value
	return obj
}

type BgpPrefixIpv4MplsVpnUnicastStateOriginEnum string

// Enum of Origin on BgpPrefixIpv4MplsVpnUnicastState
var BgpPrefixIpv4MplsVpnUnicastStateOrigin = struct {
	IGP        BgpPrefixIpv4MplsVpnUnicastStateOriginEnum
	EGP        BgpPrefixIpv4MplsVpnUnicastStateOriginEnum
	INCOMPLETE BgpPrefixIpv4MplsVpnUnicastStateOriginEnum
}{
	IGP:        BgpPrefixIpv4MplsVpnUnicastStateOriginEnum("igp"),
	EGP:        BgpPrefixIpv4MplsVpnUnicastStateOriginEnum("egp"),
	INCOMPLETE: BgpPrefixIpv4MplsVpnUnicastStateOriginEnum("incomplete"),
}

func (obj *bgpPrefixIpv4MplsVpnUnicastState) Origin() BgpPrefixIpv4MplsVpnUnicastStateOriginEnum {
	return BgpPrefixIpv4MplsVpnUnicastStateOriginEnum(obj.obj.Origin.Enum().String())
}

// The origin of the prefix.
// Origin returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastState) HasOrigin() bool {
	return obj.obj.Origin != nil
}

func (obj *bgpPrefixIpv4MplsVpnUnicastState) SetOrigin(value BgpPrefixIpv4MplsVpnUnicastStateOriginEnum) BgpPrefixIpv4MplsVpnUnicastState {
	intValue, ok := otg.BgpPrefixIpv4MplsVpnUnicastState_Origin_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpPrefixIpv4MplsVpnUnicastStateOriginEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpPrefixIpv4MplsVpnUnicastState_Origin_Enum(intValue)
	obj.obj.Origin = &enumValue

	return obj
}

// The path id.
// PathId returns a uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastState) PathId() uint32 {

	return *obj.obj.PathId

}

// The path id.
// PathId returns a uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastState) HasPathId() bool {
	return obj.obj.PathId != nil
}

// The path id.
// SetPathId sets the uint32 value in the BgpPrefixIpv4MplsVpnUnicastState object
func (obj *bgpPrefixIpv4MplsVpnUnicastState) SetPathId(value uint32) BgpPrefixIpv4MplsVpnUnicastState {

	obj.obj.PathId = &value
	return obj
}

// The IPv4 address of the egress interface.
// Ipv4NextHop returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastState) Ipv4NextHop() string {

	return *obj.obj.Ipv4NextHop

}

// The IPv4 address of the egress interface.
// Ipv4NextHop returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastState) HasIpv4NextHop() bool {
	return obj.obj.Ipv4NextHop != nil
}

// The IPv4 address of the egress interface.
// SetIpv4NextHop sets the string value in the BgpPrefixIpv4MplsVpnUnicastState object
func (obj *bgpPrefixIpv4MplsVpnUnicastState) SetIpv4NextHop(value string) BgpPrefixIpv4MplsVpnUnicastState {

	obj.obj.Ipv4NextHop = &value
	return obj
}

// The IPv6 address of the egress interface.
// Ipv6NextHop returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastState) Ipv6NextHop() string {

	return *obj.obj.Ipv6NextHop

}

// The IPv6 address of the egress interface.
// Ipv6NextHop returns a string
func (obj *bgpPrefixIpv4MplsVpnUnicastState) HasIpv6NextHop() bool {
	return obj.obj.Ipv6NextHop != nil
}

// The IPv6 address of the egress interface.
// SetIpv6NextHop sets the string value in the BgpPrefixIpv4MplsVpnUnicastState object
func (obj *bgpPrefixIpv4MplsVpnUnicastState) SetIpv6NextHop(value string) BgpPrefixIpv4MplsVpnUnicastState {

	obj.obj.Ipv6NextHop = &value
	return obj
}

// One or more MPLS VPN Label 24 bit values bound to this VPN-IPv4 prefix (RFC 4364 Section 3).
// Labels returns a []uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastState) Labels() []uint32 {
	if obj.obj.Labels == nil {
		obj.obj.Labels = make([]uint32, 0)
	}
	return obj.obj.Labels
}

// One or more MPLS VPN Label 24 bit values bound to this VPN-IPv4 prefix (RFC 4364 Section 3).
// SetLabels sets the []uint32 value in the BgpPrefixIpv4MplsVpnUnicastState object
func (obj *bgpPrefixIpv4MplsVpnUnicastState) SetLabels(value []uint32) BgpPrefixIpv4MplsVpnUnicastState {

	if obj.obj.Labels == nil {
		obj.obj.Labels = make([]uint32, 0)
	}
	obj.obj.Labels = value

	return obj
}

// Optional community attributes.
// Communities returns a []ResultBgpCommunity
func (obj *bgpPrefixIpv4MplsVpnUnicastState) Communities() BgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter {
	if len(obj.obj.Communities) == 0 {
		obj.obj.Communities = []*otg.ResultBgpCommunity{}
	}
	if obj.communitiesHolder == nil {
		obj.communitiesHolder = newBgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter(&obj.obj.Communities).setMsg(obj)
	}
	return obj.communitiesHolder
}

type bgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter struct {
	obj                     *bgpPrefixIpv4MplsVpnUnicastState
	resultBgpCommunitySlice []ResultBgpCommunity
	fieldPtr                *[]*otg.ResultBgpCommunity
}

func newBgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter(ptr *[]*otg.ResultBgpCommunity) BgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter {
	return &bgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter{fieldPtr: ptr}
}

type BgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter interface {
	setMsg(*bgpPrefixIpv4MplsVpnUnicastState) BgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter
	Items() []ResultBgpCommunity
	Add() ResultBgpCommunity
	Append(items ...ResultBgpCommunity) BgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter
	Set(index int, newObj ResultBgpCommunity) BgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter
	Clear() BgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter
	clearHolderSlice() BgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter
	appendHolderSlice(item ResultBgpCommunity) BgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter
}

func (obj *bgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter) setMsg(msg *bgpPrefixIpv4MplsVpnUnicastState) BgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&resultBgpCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter) Items() []ResultBgpCommunity {
	return obj.resultBgpCommunitySlice
}

func (obj *bgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter) Add() ResultBgpCommunity {
	newObj := &otg.ResultBgpCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &resultBgpCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter) Append(items ...ResultBgpCommunity) BgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, item)
	}
	return obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter) Set(index int, newObj ResultBgpCommunity) BgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.resultBgpCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter) Clear() BgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ResultBgpCommunity{}
		obj.resultBgpCommunitySlice = []ResultBgpCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter) clearHolderSlice() BgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter {
	if len(obj.resultBgpCommunitySlice) > 0 {
		obj.resultBgpCommunitySlice = []ResultBgpCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter) appendHolderSlice(item ResultBgpCommunity) BgpPrefixIpv4MplsVpnUnicastStateResultBgpCommunityIter {
	obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, item)
	return obj
}

// Optional received Extended Community attributes, including the Route Target(s) (RFC 4360) attached to this VPN-IPv4 route. Each received Extended Community attribute is available for retrieval in two forms. Support of the 'raw' format in which all 8 bytes (16 hex characters) is always present and available for use. In addition, if supported by the implementation, the Extended Community attribute may also be retrieved in the 'structured' format which is an optional field.
// ExtendedCommunities returns a []ResultExtendedCommunity
func (obj *bgpPrefixIpv4MplsVpnUnicastState) ExtendedCommunities() BgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter {
	if len(obj.obj.ExtendedCommunities) == 0 {
		obj.obj.ExtendedCommunities = []*otg.ResultExtendedCommunity{}
	}
	if obj.extendedCommunitiesHolder == nil {
		obj.extendedCommunitiesHolder = newBgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter(&obj.obj.ExtendedCommunities).setMsg(obj)
	}
	return obj.extendedCommunitiesHolder
}

type bgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter struct {
	obj                          *bgpPrefixIpv4MplsVpnUnicastState
	resultExtendedCommunitySlice []ResultExtendedCommunity
	fieldPtr                     *[]*otg.ResultExtendedCommunity
}

func newBgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter(ptr *[]*otg.ResultExtendedCommunity) BgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter {
	return &bgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter{fieldPtr: ptr}
}

type BgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter interface {
	setMsg(*bgpPrefixIpv4MplsVpnUnicastState) BgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter
	Items() []ResultExtendedCommunity
	Add() ResultExtendedCommunity
	Append(items ...ResultExtendedCommunity) BgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter
	Set(index int, newObj ResultExtendedCommunity) BgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter
	Clear() BgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter
	clearHolderSlice() BgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter
	appendHolderSlice(item ResultExtendedCommunity) BgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter
}

func (obj *bgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter) setMsg(msg *bgpPrefixIpv4MplsVpnUnicastState) BgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&resultExtendedCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter) Items() []ResultExtendedCommunity {
	return obj.resultExtendedCommunitySlice
}

func (obj *bgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter) Add() ResultExtendedCommunity {
	newObj := &otg.ResultExtendedCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &resultExtendedCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter) Append(items ...ResultExtendedCommunity) BgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, item)
	}
	return obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter) Set(index int, newObj ResultExtendedCommunity) BgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.resultExtendedCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter) Clear() BgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ResultExtendedCommunity{}
		obj.resultExtendedCommunitySlice = []ResultExtendedCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter) clearHolderSlice() BgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter {
	if len(obj.resultExtendedCommunitySlice) > 0 {
		obj.resultExtendedCommunitySlice = []ResultExtendedCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter) appendHolderSlice(item ResultExtendedCommunity) BgpPrefixIpv4MplsVpnUnicastStateResultExtendedCommunityIter {
	obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, item)
	return obj
}

// description is TBD
// AsPath returns a ResultBgpAsPath
func (obj *bgpPrefixIpv4MplsVpnUnicastState) AsPath() ResultBgpAsPath {
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
func (obj *bgpPrefixIpv4MplsVpnUnicastState) HasAsPath() bool {
	return obj.obj.AsPath != nil
}

// description is TBD
// SetAsPath sets the ResultBgpAsPath value in the BgpPrefixIpv4MplsVpnUnicastState object
func (obj *bgpPrefixIpv4MplsVpnUnicastState) SetAsPath(value ResultBgpAsPath) BgpPrefixIpv4MplsVpnUnicastState {

	obj.asPathHolder = nil
	obj.obj.AsPath = value.msg()

	return obj
}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// LocalPreference returns a uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastState) LocalPreference() uint32 {

	return *obj.obj.LocalPreference

}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// LocalPreference returns a uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastState) HasLocalPreference() bool {
	return obj.obj.LocalPreference != nil
}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// SetLocalPreference sets the uint32 value in the BgpPrefixIpv4MplsVpnUnicastState object
func (obj *bgpPrefixIpv4MplsVpnUnicastState) SetLocalPreference(value uint32) BgpPrefixIpv4MplsVpnUnicastState {

	obj.obj.LocalPreference = &value
	return obj
}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// MultiExitDiscriminator returns a uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastState) MultiExitDiscriminator() uint32 {

	return *obj.obj.MultiExitDiscriminator

}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// MultiExitDiscriminator returns a uint32
func (obj *bgpPrefixIpv4MplsVpnUnicastState) HasMultiExitDiscriminator() bool {
	return obj.obj.MultiExitDiscriminator != nil
}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// SetMultiExitDiscriminator sets the uint32 value in the BgpPrefixIpv4MplsVpnUnicastState object
func (obj *bgpPrefixIpv4MplsVpnUnicastState) SetMultiExitDiscriminator(value uint32) BgpPrefixIpv4MplsVpnUnicastState {

	obj.obj.MultiExitDiscriminator = &value
	return obj
}

func (obj *bgpPrefixIpv4MplsVpnUnicastState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpPrefixIpv4MplsVpnUnicastState.PrefixLength <= 32 but Got %d", *obj.obj.PrefixLength))
		}

	}

	if obj.obj.Ipv4NextHop != nil {

		err := obj.validateIpv4(obj.Ipv4NextHop())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv4MplsVpnUnicastState.Ipv4NextHop"))
		}

	}

	if obj.obj.Ipv6NextHop != nil {

		err := obj.validateIpv6(obj.Ipv6NextHop())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv4MplsVpnUnicastState.Ipv6NextHop"))
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

func (obj *bgpPrefixIpv4MplsVpnUnicastState) setDefault() {

}
