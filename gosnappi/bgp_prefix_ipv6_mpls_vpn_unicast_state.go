package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpPrefixIpv6MplsVpnUnicastState *****
type bgpPrefixIpv6MplsVpnUnicastState struct {
	validation
	obj                       *otg.BgpPrefixIpv6MplsVpnUnicastState
	marshaller                marshalBgpPrefixIpv6MplsVpnUnicastState
	unMarshaller              unMarshalBgpPrefixIpv6MplsVpnUnicastState
	communitiesHolder         BgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter
	extendedCommunitiesHolder BgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter
	asPathHolder              ResultBgpAsPath
}

func NewBgpPrefixIpv6MplsVpnUnicastState() BgpPrefixIpv6MplsVpnUnicastState {
	obj := bgpPrefixIpv6MplsVpnUnicastState{obj: &otg.BgpPrefixIpv6MplsVpnUnicastState{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastState) msg() *otg.BgpPrefixIpv6MplsVpnUnicastState {
	return obj.obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastState) setMsg(msg *otg.BgpPrefixIpv6MplsVpnUnicastState) BgpPrefixIpv6MplsVpnUnicastState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpPrefixIpv6MplsVpnUnicastState struct {
	obj *bgpPrefixIpv6MplsVpnUnicastState
}

type marshalBgpPrefixIpv6MplsVpnUnicastState interface {
	// ToProto marshals BgpPrefixIpv6MplsVpnUnicastState to protobuf object *otg.BgpPrefixIpv6MplsVpnUnicastState
	ToProto() (*otg.BgpPrefixIpv6MplsVpnUnicastState, error)
	// ToPbText marshals BgpPrefixIpv6MplsVpnUnicastState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpPrefixIpv6MplsVpnUnicastState to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpPrefixIpv6MplsVpnUnicastState to JSON text
	ToJson() (string, error)
}

type unMarshalbgpPrefixIpv6MplsVpnUnicastState struct {
	obj *bgpPrefixIpv6MplsVpnUnicastState
}

type unMarshalBgpPrefixIpv6MplsVpnUnicastState interface {
	// FromProto unmarshals BgpPrefixIpv6MplsVpnUnicastState from protobuf object *otg.BgpPrefixIpv6MplsVpnUnicastState
	FromProto(msg *otg.BgpPrefixIpv6MplsVpnUnicastState) (BgpPrefixIpv6MplsVpnUnicastState, error)
	// FromPbText unmarshals BgpPrefixIpv6MplsVpnUnicastState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpPrefixIpv6MplsVpnUnicastState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpPrefixIpv6MplsVpnUnicastState from JSON text
	FromJson(value string) error
}

func (obj *bgpPrefixIpv6MplsVpnUnicastState) Marshal() marshalBgpPrefixIpv6MplsVpnUnicastState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpPrefixIpv6MplsVpnUnicastState{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpPrefixIpv6MplsVpnUnicastState) Unmarshal() unMarshalBgpPrefixIpv6MplsVpnUnicastState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpPrefixIpv6MplsVpnUnicastState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpPrefixIpv6MplsVpnUnicastState) ToProto() (*otg.BgpPrefixIpv6MplsVpnUnicastState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpPrefixIpv6MplsVpnUnicastState) FromProto(msg *otg.BgpPrefixIpv6MplsVpnUnicastState) (BgpPrefixIpv6MplsVpnUnicastState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpPrefixIpv6MplsVpnUnicastState) ToPbText() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6MplsVpnUnicastState) FromPbText(value string) error {
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

func (m *marshalbgpPrefixIpv6MplsVpnUnicastState) ToYaml() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6MplsVpnUnicastState) FromYaml(value string) error {
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

func (m *marshalbgpPrefixIpv6MplsVpnUnicastState) ToJson() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6MplsVpnUnicastState) FromJson(value string) error {
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

func (obj *bgpPrefixIpv6MplsVpnUnicastState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv6MplsVpnUnicastState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv6MplsVpnUnicastState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpPrefixIpv6MplsVpnUnicastState) Clone() (BgpPrefixIpv6MplsVpnUnicastState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpPrefixIpv6MplsVpnUnicastState()
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

func (obj *bgpPrefixIpv6MplsVpnUnicastState) setNil() {
	obj.communitiesHolder = nil
	obj.extendedCommunitiesHolder = nil
	obj.asPathHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpPrefixIpv6MplsVpnUnicastState is bGP/MPLS L3VPN (RFC 4659) VPN-IPv6 (6VPE) learned prefix, received under the VPN-IPv6 AFI/SAFI (AFI 2, SAFI 128).
type BgpPrefixIpv6MplsVpnUnicastState interface {
	Validation
	// msg marshals BgpPrefixIpv6MplsVpnUnicastState to protobuf object *otg.BgpPrefixIpv6MplsVpnUnicastState
	// and doesn't set defaults
	msg() *otg.BgpPrefixIpv6MplsVpnUnicastState
	// setMsg unmarshals BgpPrefixIpv6MplsVpnUnicastState from protobuf object *otg.BgpPrefixIpv6MplsVpnUnicastState
	// and doesn't set defaults
	setMsg(*otg.BgpPrefixIpv6MplsVpnUnicastState) BgpPrefixIpv6MplsVpnUnicastState
	// provides marshal interface
	Marshal() marshalBgpPrefixIpv6MplsVpnUnicastState
	// provides unmarshal interface
	Unmarshal() unMarshalBgpPrefixIpv6MplsVpnUnicastState
	// validate validates BgpPrefixIpv6MplsVpnUnicastState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpPrefixIpv6MplsVpnUnicastState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouteDistinguisher returns string, set in BgpPrefixIpv6MplsVpnUnicastState.
	RouteDistinguisher() string
	// SetRouteDistinguisher assigns string provided by user to BgpPrefixIpv6MplsVpnUnicastState
	SetRouteDistinguisher(value string) BgpPrefixIpv6MplsVpnUnicastState
	// HasRouteDistinguisher checks if RouteDistinguisher has been set in BgpPrefixIpv6MplsVpnUnicastState
	HasRouteDistinguisher() bool
	// Ipv6Address returns string, set in BgpPrefixIpv6MplsVpnUnicastState.
	Ipv6Address() string
	// SetIpv6Address assigns string provided by user to BgpPrefixIpv6MplsVpnUnicastState
	SetIpv6Address(value string) BgpPrefixIpv6MplsVpnUnicastState
	// HasIpv6Address checks if Ipv6Address has been set in BgpPrefixIpv6MplsVpnUnicastState
	HasIpv6Address() bool
	// PrefixLength returns uint32, set in BgpPrefixIpv6MplsVpnUnicastState.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to BgpPrefixIpv6MplsVpnUnicastState
	SetPrefixLength(value uint32) BgpPrefixIpv6MplsVpnUnicastState
	// HasPrefixLength checks if PrefixLength has been set in BgpPrefixIpv6MplsVpnUnicastState
	HasPrefixLength() bool
	// Origin returns BgpPrefixIpv6MplsVpnUnicastStateOriginEnum, set in BgpPrefixIpv6MplsVpnUnicastState
	Origin() BgpPrefixIpv6MplsVpnUnicastStateOriginEnum
	// SetOrigin assigns BgpPrefixIpv6MplsVpnUnicastStateOriginEnum provided by user to BgpPrefixIpv6MplsVpnUnicastState
	SetOrigin(value BgpPrefixIpv6MplsVpnUnicastStateOriginEnum) BgpPrefixIpv6MplsVpnUnicastState
	// HasOrigin checks if Origin has been set in BgpPrefixIpv6MplsVpnUnicastState
	HasOrigin() bool
	// PathId returns uint32, set in BgpPrefixIpv6MplsVpnUnicastState.
	PathId() uint32
	// SetPathId assigns uint32 provided by user to BgpPrefixIpv6MplsVpnUnicastState
	SetPathId(value uint32) BgpPrefixIpv6MplsVpnUnicastState
	// HasPathId checks if PathId has been set in BgpPrefixIpv6MplsVpnUnicastState
	HasPathId() bool
	// Ipv4NextHop returns string, set in BgpPrefixIpv6MplsVpnUnicastState.
	Ipv4NextHop() string
	// SetIpv4NextHop assigns string provided by user to BgpPrefixIpv6MplsVpnUnicastState
	SetIpv4NextHop(value string) BgpPrefixIpv6MplsVpnUnicastState
	// HasIpv4NextHop checks if Ipv4NextHop has been set in BgpPrefixIpv6MplsVpnUnicastState
	HasIpv4NextHop() bool
	// Ipv6NextHop returns string, set in BgpPrefixIpv6MplsVpnUnicastState.
	Ipv6NextHop() string
	// SetIpv6NextHop assigns string provided by user to BgpPrefixIpv6MplsVpnUnicastState
	SetIpv6NextHop(value string) BgpPrefixIpv6MplsVpnUnicastState
	// HasIpv6NextHop checks if Ipv6NextHop has been set in BgpPrefixIpv6MplsVpnUnicastState
	HasIpv6NextHop() bool
	// Labels returns []uint32, set in BgpPrefixIpv6MplsVpnUnicastState.
	Labels() []uint32
	// SetLabels assigns []uint32 provided by user to BgpPrefixIpv6MplsVpnUnicastState
	SetLabels(value []uint32) BgpPrefixIpv6MplsVpnUnicastState
	// Communities returns BgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIterIter, set in BgpPrefixIpv6MplsVpnUnicastState
	Communities() BgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter
	// ExtendedCommunities returns BgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIterIter, set in BgpPrefixIpv6MplsVpnUnicastState
	ExtendedCommunities() BgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter
	// AsPath returns ResultBgpAsPath, set in BgpPrefixIpv6MplsVpnUnicastState.
	// ResultBgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed.
	AsPath() ResultBgpAsPath
	// SetAsPath assigns ResultBgpAsPath provided by user to BgpPrefixIpv6MplsVpnUnicastState.
	// ResultBgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed.
	SetAsPath(value ResultBgpAsPath) BgpPrefixIpv6MplsVpnUnicastState
	// HasAsPath checks if AsPath has been set in BgpPrefixIpv6MplsVpnUnicastState
	HasAsPath() bool
	// LocalPreference returns uint32, set in BgpPrefixIpv6MplsVpnUnicastState.
	LocalPreference() uint32
	// SetLocalPreference assigns uint32 provided by user to BgpPrefixIpv6MplsVpnUnicastState
	SetLocalPreference(value uint32) BgpPrefixIpv6MplsVpnUnicastState
	// HasLocalPreference checks if LocalPreference has been set in BgpPrefixIpv6MplsVpnUnicastState
	HasLocalPreference() bool
	// MultiExitDiscriminator returns uint32, set in BgpPrefixIpv6MplsVpnUnicastState.
	MultiExitDiscriminator() uint32
	// SetMultiExitDiscriminator assigns uint32 provided by user to BgpPrefixIpv6MplsVpnUnicastState
	SetMultiExitDiscriminator(value uint32) BgpPrefixIpv6MplsVpnUnicastState
	// HasMultiExitDiscriminator checks if MultiExitDiscriminator has been set in BgpPrefixIpv6MplsVpnUnicastState
	HasMultiExitDiscriminator() bool
	setNil()
}

// The Route Distinguisher (RFC 4364 Section 4.1) received as part of the VPN-IPv6 NLRI, formatted as a colon separated value, for example "60005:100" or "1.1.1.1:100".
// RouteDistinguisher returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastState) RouteDistinguisher() string {

	return *obj.obj.RouteDistinguisher

}

// The Route Distinguisher (RFC 4364 Section 4.1) received as part of the VPN-IPv6 NLRI, formatted as a colon separated value, for example "60005:100" or "1.1.1.1:100".
// RouteDistinguisher returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastState) HasRouteDistinguisher() bool {
	return obj.obj.RouteDistinguisher != nil
}

// The Route Distinguisher (RFC 4364 Section 4.1) received as part of the VPN-IPv6 NLRI, formatted as a colon separated value, for example "60005:100" or "1.1.1.1:100".
// SetRouteDistinguisher sets the string value in the BgpPrefixIpv6MplsVpnUnicastState object
func (obj *bgpPrefixIpv6MplsVpnUnicastState) SetRouteDistinguisher(value string) BgpPrefixIpv6MplsVpnUnicastState {

	obj.obj.RouteDistinguisher = &value
	return obj
}

// An IPv6 unicast address.
// Ipv6Address returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastState) Ipv6Address() string {

	return *obj.obj.Ipv6Address

}

// An IPv6 unicast address.
// Ipv6Address returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastState) HasIpv6Address() bool {
	return obj.obj.Ipv6Address != nil
}

// An IPv6 unicast address.
// SetIpv6Address sets the string value in the BgpPrefixIpv6MplsVpnUnicastState object
func (obj *bgpPrefixIpv6MplsVpnUnicastState) SetIpv6Address(value string) BgpPrefixIpv6MplsVpnUnicastState {

	obj.obj.Ipv6Address = &value
	return obj
}

// description is TBD
// PrefixLength returns a uint32
func (obj *bgpPrefixIpv6MplsVpnUnicastState) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// description is TBD
// PrefixLength returns a uint32
func (obj *bgpPrefixIpv6MplsVpnUnicastState) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// description is TBD
// SetPrefixLength sets the uint32 value in the BgpPrefixIpv6MplsVpnUnicastState object
func (obj *bgpPrefixIpv6MplsVpnUnicastState) SetPrefixLength(value uint32) BgpPrefixIpv6MplsVpnUnicastState {

	obj.obj.PrefixLength = &value
	return obj
}

type BgpPrefixIpv6MplsVpnUnicastStateOriginEnum string

// Enum of Origin on BgpPrefixIpv6MplsVpnUnicastState
var BgpPrefixIpv6MplsVpnUnicastStateOrigin = struct {
	IGP        BgpPrefixIpv6MplsVpnUnicastStateOriginEnum
	EGP        BgpPrefixIpv6MplsVpnUnicastStateOriginEnum
	INCOMPLETE BgpPrefixIpv6MplsVpnUnicastStateOriginEnum
}{
	IGP:        BgpPrefixIpv6MplsVpnUnicastStateOriginEnum("igp"),
	EGP:        BgpPrefixIpv6MplsVpnUnicastStateOriginEnum("egp"),
	INCOMPLETE: BgpPrefixIpv6MplsVpnUnicastStateOriginEnum("incomplete"),
}

func (obj *bgpPrefixIpv6MplsVpnUnicastState) Origin() BgpPrefixIpv6MplsVpnUnicastStateOriginEnum {
	return BgpPrefixIpv6MplsVpnUnicastStateOriginEnum(obj.obj.Origin.Enum().String())
}

// The origin of the prefix.
// Origin returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastState) HasOrigin() bool {
	return obj.obj.Origin != nil
}

func (obj *bgpPrefixIpv6MplsVpnUnicastState) SetOrigin(value BgpPrefixIpv6MplsVpnUnicastStateOriginEnum) BgpPrefixIpv6MplsVpnUnicastState {
	intValue, ok := otg.BgpPrefixIpv6MplsVpnUnicastState_Origin_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpPrefixIpv6MplsVpnUnicastStateOriginEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpPrefixIpv6MplsVpnUnicastState_Origin_Enum(intValue)
	obj.obj.Origin = &enumValue

	return obj
}

// The path id.
// PathId returns a uint32
func (obj *bgpPrefixIpv6MplsVpnUnicastState) PathId() uint32 {

	return *obj.obj.PathId

}

// The path id.
// PathId returns a uint32
func (obj *bgpPrefixIpv6MplsVpnUnicastState) HasPathId() bool {
	return obj.obj.PathId != nil
}

// The path id.
// SetPathId sets the uint32 value in the BgpPrefixIpv6MplsVpnUnicastState object
func (obj *bgpPrefixIpv6MplsVpnUnicastState) SetPathId(value uint32) BgpPrefixIpv6MplsVpnUnicastState {

	obj.obj.PathId = &value
	return obj
}

// The IPv4 address of the egress interface.
// Ipv4NextHop returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastState) Ipv4NextHop() string {

	return *obj.obj.Ipv4NextHop

}

// The IPv4 address of the egress interface.
// Ipv4NextHop returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastState) HasIpv4NextHop() bool {
	return obj.obj.Ipv4NextHop != nil
}

// The IPv4 address of the egress interface.
// SetIpv4NextHop sets the string value in the BgpPrefixIpv6MplsVpnUnicastState object
func (obj *bgpPrefixIpv6MplsVpnUnicastState) SetIpv4NextHop(value string) BgpPrefixIpv6MplsVpnUnicastState {

	obj.obj.Ipv4NextHop = &value
	return obj
}

// The IPv6 address of the egress interface.
// Ipv6NextHop returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastState) Ipv6NextHop() string {

	return *obj.obj.Ipv6NextHop

}

// The IPv6 address of the egress interface.
// Ipv6NextHop returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastState) HasIpv6NextHop() bool {
	return obj.obj.Ipv6NextHop != nil
}

// The IPv6 address of the egress interface.
// SetIpv6NextHop sets the string value in the BgpPrefixIpv6MplsVpnUnicastState object
func (obj *bgpPrefixIpv6MplsVpnUnicastState) SetIpv6NextHop(value string) BgpPrefixIpv6MplsVpnUnicastState {

	obj.obj.Ipv6NextHop = &value
	return obj
}

// One or more MPLS VPN Label 24 bit values bound to this VPN-IPv6 prefix (RFC 4364 Section 3, RFC 4659).
// Labels returns a []uint32
func (obj *bgpPrefixIpv6MplsVpnUnicastState) Labels() []uint32 {
	if obj.obj.Labels == nil {
		obj.obj.Labels = make([]uint32, 0)
	}
	return obj.obj.Labels
}

// One or more MPLS VPN Label 24 bit values bound to this VPN-IPv6 prefix (RFC 4364 Section 3, RFC 4659).
// SetLabels sets the []uint32 value in the BgpPrefixIpv6MplsVpnUnicastState object
func (obj *bgpPrefixIpv6MplsVpnUnicastState) SetLabels(value []uint32) BgpPrefixIpv6MplsVpnUnicastState {

	if obj.obj.Labels == nil {
		obj.obj.Labels = make([]uint32, 0)
	}
	obj.obj.Labels = value

	return obj
}

// Optional community attributes.
// Communities returns a []ResultBgpCommunity
func (obj *bgpPrefixIpv6MplsVpnUnicastState) Communities() BgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter {
	if len(obj.obj.Communities) == 0 {
		obj.obj.Communities = []*otg.ResultBgpCommunity{}
	}
	if obj.communitiesHolder == nil {
		obj.communitiesHolder = newBgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter(&obj.obj.Communities).setMsg(obj)
	}
	return obj.communitiesHolder
}

type bgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter struct {
	obj                     *bgpPrefixIpv6MplsVpnUnicastState
	resultBgpCommunitySlice []ResultBgpCommunity
	fieldPtr                *[]*otg.ResultBgpCommunity
}

func newBgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter(ptr *[]*otg.ResultBgpCommunity) BgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter {
	return &bgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter{fieldPtr: ptr}
}

type BgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter interface {
	setMsg(*bgpPrefixIpv6MplsVpnUnicastState) BgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter
	Items() []ResultBgpCommunity
	Add() ResultBgpCommunity
	Append(items ...ResultBgpCommunity) BgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter
	Set(index int, newObj ResultBgpCommunity) BgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter
	Clear() BgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter
	clearHolderSlice() BgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter
	appendHolderSlice(item ResultBgpCommunity) BgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter
}

func (obj *bgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter) setMsg(msg *bgpPrefixIpv6MplsVpnUnicastState) BgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&resultBgpCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter) Items() []ResultBgpCommunity {
	return obj.resultBgpCommunitySlice
}

func (obj *bgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter) Add() ResultBgpCommunity {
	newObj := &otg.ResultBgpCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &resultBgpCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter) Append(items ...ResultBgpCommunity) BgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, item)
	}
	return obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter) Set(index int, newObj ResultBgpCommunity) BgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.resultBgpCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter) Clear() BgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ResultBgpCommunity{}
		obj.resultBgpCommunitySlice = []ResultBgpCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter) clearHolderSlice() BgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter {
	if len(obj.resultBgpCommunitySlice) > 0 {
		obj.resultBgpCommunitySlice = []ResultBgpCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter) appendHolderSlice(item ResultBgpCommunity) BgpPrefixIpv6MplsVpnUnicastStateResultBgpCommunityIter {
	obj.resultBgpCommunitySlice = append(obj.resultBgpCommunitySlice, item)
	return obj
}

// Optional received Extended Community attributes, including the Route Target(s) (RFC 4360) attached to this VPN-IPv6 route. Each received Extended Community attribute is available for retrieval in two forms. Support of the 'raw' format in which all 8 bytes (16 hex characters) is always present and available for use. In addition, if supported by the implementation, the Extended Community attribute may also be retrieved in the 'structured' format which is an optional field.
// ExtendedCommunities returns a []ResultExtendedCommunity
func (obj *bgpPrefixIpv6MplsVpnUnicastState) ExtendedCommunities() BgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter {
	if len(obj.obj.ExtendedCommunities) == 0 {
		obj.obj.ExtendedCommunities = []*otg.ResultExtendedCommunity{}
	}
	if obj.extendedCommunitiesHolder == nil {
		obj.extendedCommunitiesHolder = newBgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter(&obj.obj.ExtendedCommunities).setMsg(obj)
	}
	return obj.extendedCommunitiesHolder
}

type bgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter struct {
	obj                          *bgpPrefixIpv6MplsVpnUnicastState
	resultExtendedCommunitySlice []ResultExtendedCommunity
	fieldPtr                     *[]*otg.ResultExtendedCommunity
}

func newBgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter(ptr *[]*otg.ResultExtendedCommunity) BgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter {
	return &bgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter{fieldPtr: ptr}
}

type BgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter interface {
	setMsg(*bgpPrefixIpv6MplsVpnUnicastState) BgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter
	Items() []ResultExtendedCommunity
	Add() ResultExtendedCommunity
	Append(items ...ResultExtendedCommunity) BgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter
	Set(index int, newObj ResultExtendedCommunity) BgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter
	Clear() BgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter
	clearHolderSlice() BgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter
	appendHolderSlice(item ResultExtendedCommunity) BgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter
}

func (obj *bgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter) setMsg(msg *bgpPrefixIpv6MplsVpnUnicastState) BgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&resultExtendedCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter) Items() []ResultExtendedCommunity {
	return obj.resultExtendedCommunitySlice
}

func (obj *bgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter) Add() ResultExtendedCommunity {
	newObj := &otg.ResultExtendedCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &resultExtendedCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter) Append(items ...ResultExtendedCommunity) BgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, item)
	}
	return obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter) Set(index int, newObj ResultExtendedCommunity) BgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.resultExtendedCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter) Clear() BgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ResultExtendedCommunity{}
		obj.resultExtendedCommunitySlice = []ResultExtendedCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter) clearHolderSlice() BgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter {
	if len(obj.resultExtendedCommunitySlice) > 0 {
		obj.resultExtendedCommunitySlice = []ResultExtendedCommunity{}
	}
	return obj
}
func (obj *bgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter) appendHolderSlice(item ResultExtendedCommunity) BgpPrefixIpv6MplsVpnUnicastStateResultExtendedCommunityIter {
	obj.resultExtendedCommunitySlice = append(obj.resultExtendedCommunitySlice, item)
	return obj
}

// description is TBD
// AsPath returns a ResultBgpAsPath
func (obj *bgpPrefixIpv6MplsVpnUnicastState) AsPath() ResultBgpAsPath {
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
func (obj *bgpPrefixIpv6MplsVpnUnicastState) HasAsPath() bool {
	return obj.obj.AsPath != nil
}

// description is TBD
// SetAsPath sets the ResultBgpAsPath value in the BgpPrefixIpv6MplsVpnUnicastState object
func (obj *bgpPrefixIpv6MplsVpnUnicastState) SetAsPath(value ResultBgpAsPath) BgpPrefixIpv6MplsVpnUnicastState {

	obj.asPathHolder = nil
	obj.obj.AsPath = value.msg()

	return obj
}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// LocalPreference returns a uint32
func (obj *bgpPrefixIpv6MplsVpnUnicastState) LocalPreference() uint32 {

	return *obj.obj.LocalPreference

}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// LocalPreference returns a uint32
func (obj *bgpPrefixIpv6MplsVpnUnicastState) HasLocalPreference() bool {
	return obj.obj.LocalPreference != nil
}

// The local preference is a well-known attribute and the value is used for route selection. The route with the highest local preference value is preferred.
// SetLocalPreference sets the uint32 value in the BgpPrefixIpv6MplsVpnUnicastState object
func (obj *bgpPrefixIpv6MplsVpnUnicastState) SetLocalPreference(value uint32) BgpPrefixIpv6MplsVpnUnicastState {

	obj.obj.LocalPreference = &value
	return obj
}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// MultiExitDiscriminator returns a uint32
func (obj *bgpPrefixIpv6MplsVpnUnicastState) MultiExitDiscriminator() uint32 {

	return *obj.obj.MultiExitDiscriminator

}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// MultiExitDiscriminator returns a uint32
func (obj *bgpPrefixIpv6MplsVpnUnicastState) HasMultiExitDiscriminator() bool {
	return obj.obj.MultiExitDiscriminator != nil
}

// The multi exit discriminator (MED) is an optional non-transitive attribute and the value is used for route selection. The route with the lowest MED value is preferred.
// SetMultiExitDiscriminator sets the uint32 value in the BgpPrefixIpv6MplsVpnUnicastState object
func (obj *bgpPrefixIpv6MplsVpnUnicastState) SetMultiExitDiscriminator(value uint32) BgpPrefixIpv6MplsVpnUnicastState {

	obj.obj.MultiExitDiscriminator = &value
	return obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpPrefixIpv6MplsVpnUnicastState.PrefixLength <= 128 but Got %d", *obj.obj.PrefixLength))
		}

	}

	if obj.obj.Ipv4NextHop != nil {

		err := obj.validateIpv4(obj.Ipv4NextHop())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv6MplsVpnUnicastState.Ipv4NextHop"))
		}

	}

	if obj.obj.Ipv6NextHop != nil {

		err := obj.validateIpv6(obj.Ipv6NextHop())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv6MplsVpnUnicastState.Ipv6NextHop"))
		}

	}

	if obj.obj.Labels != nil {

		for _, item := range obj.obj.Labels {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= BgpPrefixIpv6MplsVpnUnicastState.Labels <= 255 but Got %d", item))
			}

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

func (obj *bgpPrefixIpv6MplsVpnUnicastState) setDefault() {

}
