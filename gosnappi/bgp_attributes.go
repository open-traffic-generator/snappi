package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributes *****
type bgpAttributes struct {
	validation
	obj                          *otg.BgpAttributes
	marshaller                   marshalBgpAttributes
	unMarshaller                 unMarshalBgpAttributes
	otherAttributesHolder        BgpAttributesBgpAttributesOtherAttributeIter
	asPathHolder                 BgpAttributesAsPath
	as4PathHolder                BgpAttributesAs4Path
	nextHopHolder                BgpAttributesNextHop
	multiExitDiscriminatorHolder BgpAttributesMultiExitDiscriminator
	localPreferenceHolder        BgpAttributesLocalPreference
	aggregatorHolder             BgpAttributesAggregator
	as4AggregatorHolder          BgpAttributesAs4Aggregator
	communityHolder              BgpAttributesBgpAttributesCommunityIter
	originatorIdHolder           BgpAttributesOriginatorId
	extendedCommunitiesHolder    BgpAttributesBgpExtendedCommunityIter
	mpReachHolder                BgpAttributesMpReachNlri
	mpUnreachHolder              BgpAttributesMpUnreachNlri
}

func NewBgpAttributes() BgpAttributes {
	obj := bgpAttributes{obj: &otg.BgpAttributes{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributes) msg() *otg.BgpAttributes {
	return obj.obj
}

func (obj *bgpAttributes) setMsg(msg *otg.BgpAttributes) BgpAttributes {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributes struct {
	obj *bgpAttributes
}

type marshalBgpAttributes interface {
	// ToProto marshals BgpAttributes to protobuf object *otg.BgpAttributes
	ToProto() (*otg.BgpAttributes, error)
	// ToPbText marshals BgpAttributes to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributes to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributes to JSON text
	ToJson() (string, error)
}

type unMarshalbgpAttributes struct {
	obj *bgpAttributes
}

type unMarshalBgpAttributes interface {
	// FromProto unmarshals BgpAttributes from protobuf object *otg.BgpAttributes
	FromProto(msg *otg.BgpAttributes) (BgpAttributes, error)
	// FromPbText unmarshals BgpAttributes from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributes from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributes from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributes) Marshal() marshalBgpAttributes {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributes{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributes) Unmarshal() unMarshalBgpAttributes {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributes{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributes) ToProto() (*otg.BgpAttributes, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributes) FromProto(msg *otg.BgpAttributes) (BgpAttributes, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributes) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributes) FromPbText(value string) error {
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

func (m *marshalbgpAttributes) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributes) FromYaml(value string) error {
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

func (m *marshalbgpAttributes) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributes) FromJson(value string) error {
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

func (obj *bgpAttributes) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributes) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributes) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributes) Clone() (BgpAttributes, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributes()
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

func (obj *bgpAttributes) setNil() {
	obj.otherAttributesHolder = nil
	obj.asPathHolder = nil
	obj.as4PathHolder = nil
	obj.nextHopHolder = nil
	obj.multiExitDiscriminatorHolder = nil
	obj.localPreferenceHolder = nil
	obj.aggregatorHolder = nil
	obj.as4AggregatorHolder = nil
	obj.communityHolder = nil
	obj.originatorIdHolder = nil
	obj.extendedCommunitiesHolder = nil
	obj.mpReachHolder = nil
	obj.mpUnreachHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpAttributes is attributes carried in the Update packet alongwith the reach/unreach prefixes.
type BgpAttributes interface {
	Validation
	// msg marshals BgpAttributes to protobuf object *otg.BgpAttributes
	// and doesn't set defaults
	msg() *otg.BgpAttributes
	// setMsg unmarshals BgpAttributes from protobuf object *otg.BgpAttributes
	// and doesn't set defaults
	setMsg(*otg.BgpAttributes) BgpAttributes
	// provides marshal interface
	Marshal() marshalBgpAttributes
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributes
	// validate validates BgpAttributes
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributes, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// OtherAttributes returns BgpAttributesBgpAttributesOtherAttributeIterIter, set in BgpAttributes
	OtherAttributes() BgpAttributesBgpAttributesOtherAttributeIter
	// Origin returns BgpAttributesOriginEnum, set in BgpAttributes
	Origin() BgpAttributesOriginEnum
	// SetOrigin assigns BgpAttributesOriginEnum provided by user to BgpAttributes
	SetOrigin(value BgpAttributesOriginEnum) BgpAttributes
	// HasOrigin checks if Origin has been set in BgpAttributes
	HasOrigin() bool
	// AsPath returns BgpAttributesAsPath, set in BgpAttributes.
	// BgpAttributesAsPath is the AS_PATH attribute identifies the autonomous systems through  which routing information
	// carried in this UPDATE message has passed.
	// This contains the configuration of how to include the Local AS in the AS path
	// attribute of the MP REACH NLRI. It also contains optional configuration of
	// additional AS Path Segments that can be included in the AS Path attribute.
	// The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that
	// a routing information passes through to reach the destination.
	// There are two modes in which AS numbers can be encoded in the AS Path Segments
	// - When the AS Path is being exchanged between old and new BGP speakers or between two old BGP speakers , the AS numbers are encoded as 2 byte values.
	// - When the AS Path is being exchanged between two new BGP speakers supporting 4 byte AS , the AS numbers are encoded as 4 byte values.
	AsPath() BgpAttributesAsPath
	// SetAsPath assigns BgpAttributesAsPath provided by user to BgpAttributes.
	// BgpAttributesAsPath is the AS_PATH attribute identifies the autonomous systems through  which routing information
	// carried in this UPDATE message has passed.
	// This contains the configuration of how to include the Local AS in the AS path
	// attribute of the MP REACH NLRI. It also contains optional configuration of
	// additional AS Path Segments that can be included in the AS Path attribute.
	// The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that
	// a routing information passes through to reach the destination.
	// There are two modes in which AS numbers can be encoded in the AS Path Segments
	// - When the AS Path is being exchanged between old and new BGP speakers or between two old BGP speakers , the AS numbers are encoded as 2 byte values.
	// - When the AS Path is being exchanged between two new BGP speakers supporting 4 byte AS , the AS numbers are encoded as 4 byte values.
	SetAsPath(value BgpAttributesAsPath) BgpAttributes
	// HasAsPath checks if AsPath has been set in BgpAttributes
	HasAsPath() bool
	// As4Path returns BgpAttributesAs4Path, set in BgpAttributes.
	// BgpAttributesAs4Path is the AS4_PATH attribute identifies the autonomous systems through  which routing information
	// carried in this UPDATE message has passed.
	// This contains the configuration of how to include the Local AS in the AS path
	// attribute of the MP REACH NLRI. It also contains optional configuration of
	// additional AS Path Segments that can be included in the AS Path attribute.
	// The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that
	// a routing information passes through to reach the destination.
	// AS4_PATH is only exchanged in two scenarios:
	// - When an old BGP speaker has to forward a received AS4_PATH containing 4 byte AS numbers to new BGP speaker.
	// - When a new BGP speaker is connected to an old BGP speaker and has to propagate 4 byte AS numbers via the old BGP speaker.
	// Its usage is described in RFC4893.
	As4Path() BgpAttributesAs4Path
	// SetAs4Path assigns BgpAttributesAs4Path provided by user to BgpAttributes.
	// BgpAttributesAs4Path is the AS4_PATH attribute identifies the autonomous systems through  which routing information
	// carried in this UPDATE message has passed.
	// This contains the configuration of how to include the Local AS in the AS path
	// attribute of the MP REACH NLRI. It also contains optional configuration of
	// additional AS Path Segments that can be included in the AS Path attribute.
	// The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that
	// a routing information passes through to reach the destination.
	// AS4_PATH is only exchanged in two scenarios:
	// - When an old BGP speaker has to forward a received AS4_PATH containing 4 byte AS numbers to new BGP speaker.
	// - When a new BGP speaker is connected to an old BGP speaker and has to propagate 4 byte AS numbers via the old BGP speaker.
	// Its usage is described in RFC4893.
	SetAs4Path(value BgpAttributesAs4Path) BgpAttributes
	// HasAs4Path checks if As4Path has been set in BgpAttributes
	HasAs4Path() bool
	// NextHop returns BgpAttributesNextHop, set in BgpAttributes.
	// BgpAttributesNextHop is next hop to be sent inside MP_REACH NLRI or as the NEXT_HOP attribute if advertised as traditional NLRI.
	NextHop() BgpAttributesNextHop
	// SetNextHop assigns BgpAttributesNextHop provided by user to BgpAttributes.
	// BgpAttributesNextHop is next hop to be sent inside MP_REACH NLRI or as the NEXT_HOP attribute if advertised as traditional NLRI.
	SetNextHop(value BgpAttributesNextHop) BgpAttributes
	// HasNextHop checks if NextHop has been set in BgpAttributes
	HasNextHop() bool
	// MultiExitDiscriminator returns BgpAttributesMultiExitDiscriminator, set in BgpAttributes.
	// BgpAttributesMultiExitDiscriminator is optional MULTI_EXIT_DISCRIMINATOR attribute sent to the peer to help in the route selection process.
	MultiExitDiscriminator() BgpAttributesMultiExitDiscriminator
	// SetMultiExitDiscriminator assigns BgpAttributesMultiExitDiscriminator provided by user to BgpAttributes.
	// BgpAttributesMultiExitDiscriminator is optional MULTI_EXIT_DISCRIMINATOR attribute sent to the peer to help in the route selection process.
	SetMultiExitDiscriminator(value BgpAttributesMultiExitDiscriminator) BgpAttributes
	// HasMultiExitDiscriminator checks if MultiExitDiscriminator has been set in BgpAttributes
	HasMultiExitDiscriminator() bool
	// LocalPreference returns BgpAttributesLocalPreference, set in BgpAttributes.
	// BgpAttributesLocalPreference is optional LOCAL_PREFERENCE attribute sent to the peer to indicate the degree of preference
	// for externally learned routes.This should be included only for internal peers.It is
	// used for the selection of the path for the traffic leaving the AS.The route with the
	// highest local preference value is preferred.
	LocalPreference() BgpAttributesLocalPreference
	// SetLocalPreference assigns BgpAttributesLocalPreference provided by user to BgpAttributes.
	// BgpAttributesLocalPreference is optional LOCAL_PREFERENCE attribute sent to the peer to indicate the degree of preference
	// for externally learned routes.This should be included only for internal peers.It is
	// used for the selection of the path for the traffic leaving the AS.The route with the
	// highest local preference value is preferred.
	SetLocalPreference(value BgpAttributesLocalPreference) BgpAttributes
	// HasLocalPreference checks if LocalPreference has been set in BgpAttributes
	HasLocalPreference() bool
	// IncludeAtomicAggregator returns bool, set in BgpAttributes.
	IncludeAtomicAggregator() bool
	// SetIncludeAtomicAggregator assigns bool provided by user to BgpAttributes
	SetIncludeAtomicAggregator(value bool) BgpAttributes
	// HasIncludeAtomicAggregator checks if IncludeAtomicAggregator has been set in BgpAttributes
	HasIncludeAtomicAggregator() bool
	// Aggregator returns BgpAttributesAggregator, set in BgpAttributes.
	// BgpAttributesAggregator is optional AGGREGATOR attribute which maybe be added by a BGP speaker which performs route aggregation.
	// When AGGREGATOR attribute is being sent to a new BGP speaker , the AS number is encoded as a 4 byte value.
	// When AGGREGATOR attribute is being exchanged between a new and an old BGP speaker or between two old BGP speakers,
	// the AS number is encoded as a 2 byte value.
	// It contain the AS number and IP address of the speaker performing the aggregation.
	Aggregator() BgpAttributesAggregator
	// SetAggregator assigns BgpAttributesAggregator provided by user to BgpAttributes.
	// BgpAttributesAggregator is optional AGGREGATOR attribute which maybe be added by a BGP speaker which performs route aggregation.
	// When AGGREGATOR attribute is being sent to a new BGP speaker , the AS number is encoded as a 4 byte value.
	// When AGGREGATOR attribute is being exchanged between a new and an old BGP speaker or between two old BGP speakers,
	// the AS number is encoded as a 2 byte value.
	// It contain the AS number and IP address of the speaker performing the aggregation.
	SetAggregator(value BgpAttributesAggregator) BgpAttributes
	// HasAggregator checks if Aggregator has been set in BgpAttributes
	HasAggregator() bool
	// As4Aggregator returns BgpAttributesAs4Aggregator, set in BgpAttributes.
	// BgpAttributesAs4Aggregator is optional AS4_AGGREGATOR attribute which maybe be added by a BGP speaker in one of two cases:
	// - If it is a new BGP speaker speaking to an old BGP speaker and needs to send a 4 byte value for the AS number of the BGP route aggregator.
	// - If it is a old BGP speaker speaking to a new BGP speaker and has to transparently forward a received AS4_AGGREGATOR from some other peer.
	// Its usage is described in RFC4893.
	As4Aggregator() BgpAttributesAs4Aggregator
	// SetAs4Aggregator assigns BgpAttributesAs4Aggregator provided by user to BgpAttributes.
	// BgpAttributesAs4Aggregator is optional AS4_AGGREGATOR attribute which maybe be added by a BGP speaker in one of two cases:
	// - If it is a new BGP speaker speaking to an old BGP speaker and needs to send a 4 byte value for the AS number of the BGP route aggregator.
	// - If it is a old BGP speaker speaking to a new BGP speaker and has to transparently forward a received AS4_AGGREGATOR from some other peer.
	// Its usage is described in RFC4893.
	SetAs4Aggregator(value BgpAttributesAs4Aggregator) BgpAttributes
	// HasAs4Aggregator checks if As4Aggregator has been set in BgpAttributes
	HasAs4Aggregator() bool
	// Community returns BgpAttributesBgpAttributesCommunityIterIter, set in BgpAttributes
	Community() BgpAttributesBgpAttributesCommunityIter
	// OriginatorId returns BgpAttributesOriginatorId, set in BgpAttributes.
	// BgpAttributesOriginatorId is optional ORIGINATOR_ID attribute (type code 9) carries the Router Id of the route's originator in the local AS.
	OriginatorId() BgpAttributesOriginatorId
	// SetOriginatorId assigns BgpAttributesOriginatorId provided by user to BgpAttributes.
	// BgpAttributesOriginatorId is optional ORIGINATOR_ID attribute (type code 9) carries the Router Id of the route's originator in the local AS.
	SetOriginatorId(value BgpAttributesOriginatorId) BgpAttributes
	// HasOriginatorId checks if OriginatorId has been set in BgpAttributes
	HasOriginatorId() bool
	// ClusterIds returns []string, set in BgpAttributes.
	ClusterIds() []string
	// SetClusterIds assigns []string provided by user to BgpAttributes
	SetClusterIds(value []string) BgpAttributes
	// ExtendedCommunities returns BgpAttributesBgpExtendedCommunityIterIter, set in BgpAttributes
	ExtendedCommunities() BgpAttributesBgpExtendedCommunityIter
	// MpReach returns BgpAttributesMpReachNlri, set in BgpAttributes.
	// BgpAttributesMpReachNlri is the MP_REACH attribute is an optional attribute which can be included in the attributes of a BGP Update message as defined in https://datatracker.ietf.org/doc/html/rfc4760#section-3.
	// The following AFI / SAFI combinations are supported:
	// - IPv4 Unicast with AFI as 1 and SAFI as 1
	// - IPv6 Unicast with AFI as 2 and SAFI as 1
	MpReach() BgpAttributesMpReachNlri
	// SetMpReach assigns BgpAttributesMpReachNlri provided by user to BgpAttributes.
	// BgpAttributesMpReachNlri is the MP_REACH attribute is an optional attribute which can be included in the attributes of a BGP Update message as defined in https://datatracker.ietf.org/doc/html/rfc4760#section-3.
	// The following AFI / SAFI combinations are supported:
	// - IPv4 Unicast with AFI as 1 and SAFI as 1
	// - IPv6 Unicast with AFI as 2 and SAFI as 1
	SetMpReach(value BgpAttributesMpReachNlri) BgpAttributes
	// HasMpReach checks if MpReach has been set in BgpAttributes
	HasMpReach() bool
	// MpUnreach returns BgpAttributesMpUnreachNlri, set in BgpAttributes.
	// BgpAttributesMpUnreachNlri is the MP_UNREACH attribute is an optional attribute which can be included in the attributes of a BGP Update message as defined in https://datatracker.ietf.org/doc/html/rfc4760#section-3.
	// The following AFI / SAFI combinations are supported:
	// - IPv4 Unicast with AFI as 1 and SAFI as 1
	// - IPv6 Unicast with AFI as 2 and SAFI as 1
	MpUnreach() BgpAttributesMpUnreachNlri
	// SetMpUnreach assigns BgpAttributesMpUnreachNlri provided by user to BgpAttributes.
	// BgpAttributesMpUnreachNlri is the MP_UNREACH attribute is an optional attribute which can be included in the attributes of a BGP Update message as defined in https://datatracker.ietf.org/doc/html/rfc4760#section-3.
	// The following AFI / SAFI combinations are supported:
	// - IPv4 Unicast with AFI as 1 and SAFI as 1
	// - IPv6 Unicast with AFI as 2 and SAFI as 1
	SetMpUnreach(value BgpAttributesMpUnreachNlri) BgpAttributes
	// HasMpUnreach checks if MpUnreach has been set in BgpAttributes
	HasMpUnreach() bool
	setNil()
}

// Any attributes not present in the list of configurable attributes should be added to the list of unknown attributes.
// OtherAttributes returns a []BgpAttributesOtherAttribute
func (obj *bgpAttributes) OtherAttributes() BgpAttributesBgpAttributesOtherAttributeIter {
	if len(obj.obj.OtherAttributes) == 0 {
		obj.obj.OtherAttributes = []*otg.BgpAttributesOtherAttribute{}
	}
	if obj.otherAttributesHolder == nil {
		obj.otherAttributesHolder = newBgpAttributesBgpAttributesOtherAttributeIter(&obj.obj.OtherAttributes).setMsg(obj)
	}
	return obj.otherAttributesHolder
}

type bgpAttributesBgpAttributesOtherAttributeIter struct {
	obj                              *bgpAttributes
	bgpAttributesOtherAttributeSlice []BgpAttributesOtherAttribute
	fieldPtr                         *[]*otg.BgpAttributesOtherAttribute
}

func newBgpAttributesBgpAttributesOtherAttributeIter(ptr *[]*otg.BgpAttributesOtherAttribute) BgpAttributesBgpAttributesOtherAttributeIter {
	return &bgpAttributesBgpAttributesOtherAttributeIter{fieldPtr: ptr}
}

type BgpAttributesBgpAttributesOtherAttributeIter interface {
	setMsg(*bgpAttributes) BgpAttributesBgpAttributesOtherAttributeIter
	Items() []BgpAttributesOtherAttribute
	Add() BgpAttributesOtherAttribute
	Append(items ...BgpAttributesOtherAttribute) BgpAttributesBgpAttributesOtherAttributeIter
	Set(index int, newObj BgpAttributesOtherAttribute) BgpAttributesBgpAttributesOtherAttributeIter
	Clear() BgpAttributesBgpAttributesOtherAttributeIter
	clearHolderSlice() BgpAttributesBgpAttributesOtherAttributeIter
	appendHolderSlice(item BgpAttributesOtherAttribute) BgpAttributesBgpAttributesOtherAttributeIter
}

func (obj *bgpAttributesBgpAttributesOtherAttributeIter) setMsg(msg *bgpAttributes) BgpAttributesBgpAttributesOtherAttributeIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpAttributesOtherAttribute{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpAttributesBgpAttributesOtherAttributeIter) Items() []BgpAttributesOtherAttribute {
	return obj.bgpAttributesOtherAttributeSlice
}

func (obj *bgpAttributesBgpAttributesOtherAttributeIter) Add() BgpAttributesOtherAttribute {
	newObj := &otg.BgpAttributesOtherAttribute{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpAttributesOtherAttribute{obj: newObj}
	newLibObj.setDefault()
	obj.bgpAttributesOtherAttributeSlice = append(obj.bgpAttributesOtherAttributeSlice, newLibObj)
	return newLibObj
}

func (obj *bgpAttributesBgpAttributesOtherAttributeIter) Append(items ...BgpAttributesOtherAttribute) BgpAttributesBgpAttributesOtherAttributeIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpAttributesOtherAttributeSlice = append(obj.bgpAttributesOtherAttributeSlice, item)
	}
	return obj
}

func (obj *bgpAttributesBgpAttributesOtherAttributeIter) Set(index int, newObj BgpAttributesOtherAttribute) BgpAttributesBgpAttributesOtherAttributeIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpAttributesOtherAttributeSlice[index] = newObj
	return obj
}
func (obj *bgpAttributesBgpAttributesOtherAttributeIter) Clear() BgpAttributesBgpAttributesOtherAttributeIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpAttributesOtherAttribute{}
		obj.bgpAttributesOtherAttributeSlice = []BgpAttributesOtherAttribute{}
	}
	return obj
}
func (obj *bgpAttributesBgpAttributesOtherAttributeIter) clearHolderSlice() BgpAttributesBgpAttributesOtherAttributeIter {
	if len(obj.bgpAttributesOtherAttributeSlice) > 0 {
		obj.bgpAttributesOtherAttributeSlice = []BgpAttributesOtherAttribute{}
	}
	return obj
}
func (obj *bgpAttributesBgpAttributesOtherAttributeIter) appendHolderSlice(item BgpAttributesOtherAttribute) BgpAttributesBgpAttributesOtherAttributeIter {
	obj.bgpAttributesOtherAttributeSlice = append(obj.bgpAttributesOtherAttributeSlice, item)
	return obj
}

type BgpAttributesOriginEnum string

// Enum of Origin on BgpAttributes
var BgpAttributesOrigin = struct {
	IGP        BgpAttributesOriginEnum
	EGP        BgpAttributesOriginEnum
	INCOMPLETE BgpAttributesOriginEnum
}{
	IGP:        BgpAttributesOriginEnum("igp"),
	EGP:        BgpAttributesOriginEnum("egp"),
	INCOMPLETE: BgpAttributesOriginEnum("incomplete"),
}

func (obj *bgpAttributes) Origin() BgpAttributesOriginEnum {
	return BgpAttributesOriginEnum(obj.obj.Origin.Enum().String())
}

// The ORIGIN attribute is a mandatory attribute which can take three values:
// the prefix originates from an interior routing protocol 'igp', it originates from 'egp'
// or the origin is 'incomplete',if the prefix is learned through other means.
// Origin returns a string
func (obj *bgpAttributes) HasOrigin() bool {
	return obj.obj.Origin != nil
}

func (obj *bgpAttributes) SetOrigin(value BgpAttributesOriginEnum) BgpAttributes {
	intValue, ok := otg.BgpAttributes_Origin_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpAttributesOriginEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpAttributes_Origin_Enum(intValue)
	obj.obj.Origin = &enumValue

	return obj
}

// AS_PATH attribute to be included in the Update.
// AsPath returns a BgpAttributesAsPath
func (obj *bgpAttributes) AsPath() BgpAttributesAsPath {
	if obj.obj.AsPath == nil {
		obj.obj.AsPath = NewBgpAttributesAsPath().msg()
	}
	if obj.asPathHolder == nil {
		obj.asPathHolder = &bgpAttributesAsPath{obj: obj.obj.AsPath}
	}
	return obj.asPathHolder
}

// AS_PATH attribute to be included in the Update.
// AsPath returns a BgpAttributesAsPath
func (obj *bgpAttributes) HasAsPath() bool {
	return obj.obj.AsPath != nil
}

// AS_PATH attribute to be included in the Update.
// SetAsPath sets the BgpAttributesAsPath value in the BgpAttributes object
func (obj *bgpAttributes) SetAsPath(value BgpAttributesAsPath) BgpAttributes {

	obj.asPathHolder = nil
	obj.obj.AsPath = value.msg()

	return obj
}

// AS4_PATH attribute to be included in the Update.
// As4Path returns a BgpAttributesAs4Path
func (obj *bgpAttributes) As4Path() BgpAttributesAs4Path {
	if obj.obj.As4Path == nil {
		obj.obj.As4Path = NewBgpAttributesAs4Path().msg()
	}
	if obj.as4PathHolder == nil {
		obj.as4PathHolder = &bgpAttributesAs4Path{obj: obj.obj.As4Path}
	}
	return obj.as4PathHolder
}

// AS4_PATH attribute to be included in the Update.
// As4Path returns a BgpAttributesAs4Path
func (obj *bgpAttributes) HasAs4Path() bool {
	return obj.obj.As4Path != nil
}

// AS4_PATH attribute to be included in the Update.
// SetAs4Path sets the BgpAttributesAs4Path value in the BgpAttributes object
func (obj *bgpAttributes) SetAs4Path(value BgpAttributesAs4Path) BgpAttributes {

	obj.as4PathHolder = nil
	obj.obj.As4Path = value.msg()

	return obj
}

// description is TBD
// NextHop returns a BgpAttributesNextHop
func (obj *bgpAttributes) NextHop() BgpAttributesNextHop {
	if obj.obj.NextHop == nil {
		obj.obj.NextHop = NewBgpAttributesNextHop().msg()
	}
	if obj.nextHopHolder == nil {
		obj.nextHopHolder = &bgpAttributesNextHop{obj: obj.obj.NextHop}
	}
	return obj.nextHopHolder
}

// description is TBD
// NextHop returns a BgpAttributesNextHop
func (obj *bgpAttributes) HasNextHop() bool {
	return obj.obj.NextHop != nil
}

// description is TBD
// SetNextHop sets the BgpAttributesNextHop value in the BgpAttributes object
func (obj *bgpAttributes) SetNextHop(value BgpAttributesNextHop) BgpAttributes {

	obj.nextHopHolder = nil
	obj.obj.NextHop = value.msg()

	return obj
}

// description is TBD
// MultiExitDiscriminator returns a BgpAttributesMultiExitDiscriminator
func (obj *bgpAttributes) MultiExitDiscriminator() BgpAttributesMultiExitDiscriminator {
	if obj.obj.MultiExitDiscriminator == nil {
		obj.obj.MultiExitDiscriminator = NewBgpAttributesMultiExitDiscriminator().msg()
	}
	if obj.multiExitDiscriminatorHolder == nil {
		obj.multiExitDiscriminatorHolder = &bgpAttributesMultiExitDiscriminator{obj: obj.obj.MultiExitDiscriminator}
	}
	return obj.multiExitDiscriminatorHolder
}

// description is TBD
// MultiExitDiscriminator returns a BgpAttributesMultiExitDiscriminator
func (obj *bgpAttributes) HasMultiExitDiscriminator() bool {
	return obj.obj.MultiExitDiscriminator != nil
}

// description is TBD
// SetMultiExitDiscriminator sets the BgpAttributesMultiExitDiscriminator value in the BgpAttributes object
func (obj *bgpAttributes) SetMultiExitDiscriminator(value BgpAttributesMultiExitDiscriminator) BgpAttributes {

	obj.multiExitDiscriminatorHolder = nil
	obj.obj.MultiExitDiscriminator = value.msg()

	return obj
}

// description is TBD
// LocalPreference returns a BgpAttributesLocalPreference
func (obj *bgpAttributes) LocalPreference() BgpAttributesLocalPreference {
	if obj.obj.LocalPreference == nil {
		obj.obj.LocalPreference = NewBgpAttributesLocalPreference().msg()
	}
	if obj.localPreferenceHolder == nil {
		obj.localPreferenceHolder = &bgpAttributesLocalPreference{obj: obj.obj.LocalPreference}
	}
	return obj.localPreferenceHolder
}

// description is TBD
// LocalPreference returns a BgpAttributesLocalPreference
func (obj *bgpAttributes) HasLocalPreference() bool {
	return obj.obj.LocalPreference != nil
}

// description is TBD
// SetLocalPreference sets the BgpAttributesLocalPreference value in the BgpAttributes object
func (obj *bgpAttributes) SetLocalPreference(value BgpAttributesLocalPreference) BgpAttributes {

	obj.localPreferenceHolder = nil
	obj.obj.LocalPreference = value.msg()

	return obj
}

// If enabled, it indicates that the ATOMIC_AGGREGATOR attribute should be included in the Update.
// Presence of this attribute Indicates that this route might not be getting sent on a fully optimized path
// since some intermediate BGP speaker has aggregated the route.
// IncludeAtomicAggregator returns a bool
func (obj *bgpAttributes) IncludeAtomicAggregator() bool {

	return *obj.obj.IncludeAtomicAggregator

}

// If enabled, it indicates that the ATOMIC_AGGREGATOR attribute should be included in the Update.
// Presence of this attribute Indicates that this route might not be getting sent on a fully optimized path
// since some intermediate BGP speaker has aggregated the route.
// IncludeAtomicAggregator returns a bool
func (obj *bgpAttributes) HasIncludeAtomicAggregator() bool {
	return obj.obj.IncludeAtomicAggregator != nil
}

// If enabled, it indicates that the ATOMIC_AGGREGATOR attribute should be included in the Update.
// Presence of this attribute Indicates that this route might not be getting sent on a fully optimized path
// since some intermediate BGP speaker has aggregated the route.
// SetIncludeAtomicAggregator sets the bool value in the BgpAttributes object
func (obj *bgpAttributes) SetIncludeAtomicAggregator(value bool) BgpAttributes {

	obj.obj.IncludeAtomicAggregator = &value
	return obj
}

// description is TBD
// Aggregator returns a BgpAttributesAggregator
func (obj *bgpAttributes) Aggregator() BgpAttributesAggregator {
	if obj.obj.Aggregator == nil {
		obj.obj.Aggregator = NewBgpAttributesAggregator().msg()
	}
	if obj.aggregatorHolder == nil {
		obj.aggregatorHolder = &bgpAttributesAggregator{obj: obj.obj.Aggregator}
	}
	return obj.aggregatorHolder
}

// description is TBD
// Aggregator returns a BgpAttributesAggregator
func (obj *bgpAttributes) HasAggregator() bool {
	return obj.obj.Aggregator != nil
}

// description is TBD
// SetAggregator sets the BgpAttributesAggregator value in the BgpAttributes object
func (obj *bgpAttributes) SetAggregator(value BgpAttributesAggregator) BgpAttributes {

	obj.aggregatorHolder = nil
	obj.obj.Aggregator = value.msg()

	return obj
}

// description is TBD
// As4Aggregator returns a BgpAttributesAs4Aggregator
func (obj *bgpAttributes) As4Aggregator() BgpAttributesAs4Aggregator {
	if obj.obj.As4Aggregator == nil {
		obj.obj.As4Aggregator = NewBgpAttributesAs4Aggregator().msg()
	}
	if obj.as4AggregatorHolder == nil {
		obj.as4AggregatorHolder = &bgpAttributesAs4Aggregator{obj: obj.obj.As4Aggregator}
	}
	return obj.as4AggregatorHolder
}

// description is TBD
// As4Aggregator returns a BgpAttributesAs4Aggregator
func (obj *bgpAttributes) HasAs4Aggregator() bool {
	return obj.obj.As4Aggregator != nil
}

// description is TBD
// SetAs4Aggregator sets the BgpAttributesAs4Aggregator value in the BgpAttributes object
func (obj *bgpAttributes) SetAs4Aggregator(value BgpAttributesAs4Aggregator) BgpAttributes {

	obj.as4AggregatorHolder = nil
	obj.obj.As4Aggregator = value.msg()

	return obj
}

// description is TBD
// Community returns a []BgpAttributesCommunity
func (obj *bgpAttributes) Community() BgpAttributesBgpAttributesCommunityIter {
	if len(obj.obj.Community) == 0 {
		obj.obj.Community = []*otg.BgpAttributesCommunity{}
	}
	if obj.communityHolder == nil {
		obj.communityHolder = newBgpAttributesBgpAttributesCommunityIter(&obj.obj.Community).setMsg(obj)
	}
	return obj.communityHolder
}

type bgpAttributesBgpAttributesCommunityIter struct {
	obj                         *bgpAttributes
	bgpAttributesCommunitySlice []BgpAttributesCommunity
	fieldPtr                    *[]*otg.BgpAttributesCommunity
}

func newBgpAttributesBgpAttributesCommunityIter(ptr *[]*otg.BgpAttributesCommunity) BgpAttributesBgpAttributesCommunityIter {
	return &bgpAttributesBgpAttributesCommunityIter{fieldPtr: ptr}
}

type BgpAttributesBgpAttributesCommunityIter interface {
	setMsg(*bgpAttributes) BgpAttributesBgpAttributesCommunityIter
	Items() []BgpAttributesCommunity
	Add() BgpAttributesCommunity
	Append(items ...BgpAttributesCommunity) BgpAttributesBgpAttributesCommunityIter
	Set(index int, newObj BgpAttributesCommunity) BgpAttributesBgpAttributesCommunityIter
	Clear() BgpAttributesBgpAttributesCommunityIter
	clearHolderSlice() BgpAttributesBgpAttributesCommunityIter
	appendHolderSlice(item BgpAttributesCommunity) BgpAttributesBgpAttributesCommunityIter
}

func (obj *bgpAttributesBgpAttributesCommunityIter) setMsg(msg *bgpAttributes) BgpAttributesBgpAttributesCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpAttributesCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpAttributesBgpAttributesCommunityIter) Items() []BgpAttributesCommunity {
	return obj.bgpAttributesCommunitySlice
}

func (obj *bgpAttributesBgpAttributesCommunityIter) Add() BgpAttributesCommunity {
	newObj := &otg.BgpAttributesCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpAttributesCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.bgpAttributesCommunitySlice = append(obj.bgpAttributesCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpAttributesBgpAttributesCommunityIter) Append(items ...BgpAttributesCommunity) BgpAttributesBgpAttributesCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpAttributesCommunitySlice = append(obj.bgpAttributesCommunitySlice, item)
	}
	return obj
}

func (obj *bgpAttributesBgpAttributesCommunityIter) Set(index int, newObj BgpAttributesCommunity) BgpAttributesBgpAttributesCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpAttributesCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpAttributesBgpAttributesCommunityIter) Clear() BgpAttributesBgpAttributesCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpAttributesCommunity{}
		obj.bgpAttributesCommunitySlice = []BgpAttributesCommunity{}
	}
	return obj
}
func (obj *bgpAttributesBgpAttributesCommunityIter) clearHolderSlice() BgpAttributesBgpAttributesCommunityIter {
	if len(obj.bgpAttributesCommunitySlice) > 0 {
		obj.bgpAttributesCommunitySlice = []BgpAttributesCommunity{}
	}
	return obj
}
func (obj *bgpAttributesBgpAttributesCommunityIter) appendHolderSlice(item BgpAttributesCommunity) BgpAttributesBgpAttributesCommunityIter {
	obj.bgpAttributesCommunitySlice = append(obj.bgpAttributesCommunitySlice, item)
	return obj
}

// description is TBD
// OriginatorId returns a BgpAttributesOriginatorId
func (obj *bgpAttributes) OriginatorId() BgpAttributesOriginatorId {
	if obj.obj.OriginatorId == nil {
		obj.obj.OriginatorId = NewBgpAttributesOriginatorId().msg()
	}
	if obj.originatorIdHolder == nil {
		obj.originatorIdHolder = &bgpAttributesOriginatorId{obj: obj.obj.OriginatorId}
	}
	return obj.originatorIdHolder
}

// description is TBD
// OriginatorId returns a BgpAttributesOriginatorId
func (obj *bgpAttributes) HasOriginatorId() bool {
	return obj.obj.OriginatorId != nil
}

// description is TBD
// SetOriginatorId sets the BgpAttributesOriginatorId value in the BgpAttributes object
func (obj *bgpAttributes) SetOriginatorId(value BgpAttributesOriginatorId) BgpAttributes {

	obj.originatorIdHolder = nil
	obj.obj.OriginatorId = value.msg()

	return obj
}

// When a Route Reflector reflects a route, it prepends the local CLUSTER_ID to the CLUSTER_LIST as defined in RFC4456.
// ClusterIds returns a []string
func (obj *bgpAttributes) ClusterIds() []string {
	if obj.obj.ClusterIds == nil {
		obj.obj.ClusterIds = make([]string, 0)
	}
	return obj.obj.ClusterIds
}

// When a Route Reflector reflects a route, it prepends the local CLUSTER_ID to the CLUSTER_LIST as defined in RFC4456.
// SetClusterIds sets the []string value in the BgpAttributes object
func (obj *bgpAttributes) SetClusterIds(value []string) BgpAttributes {

	if obj.obj.ClusterIds == nil {
		obj.obj.ClusterIds = make([]string, 0)
	}
	obj.obj.ClusterIds = value

	return obj
}

// Optional EXTENDED_COMMUNITY attribute settings.
// The EXTENDED_COMMUNITY Attribute is a transitive optional BGP attribute, with the Type Code 16. Community and Extended Communities  attributes
// are utilized to trigger routing decisions, such as acceptance, rejection,  preference, or redistribution. An extended community is an eight byte value.
// It is divided into two main parts. The first two bytes of the community encode a type and sub-type fields and the last six bytes carry a unique set
// of data in a format defined by the type and sub-type field. Extended communities provide a larger range for grouping or categorizing communities.
// ExtendedCommunities returns a []BgpExtendedCommunity
func (obj *bgpAttributes) ExtendedCommunities() BgpAttributesBgpExtendedCommunityIter {
	if len(obj.obj.ExtendedCommunities) == 0 {
		obj.obj.ExtendedCommunities = []*otg.BgpExtendedCommunity{}
	}
	if obj.extendedCommunitiesHolder == nil {
		obj.extendedCommunitiesHolder = newBgpAttributesBgpExtendedCommunityIter(&obj.obj.ExtendedCommunities).setMsg(obj)
	}
	return obj.extendedCommunitiesHolder
}

type bgpAttributesBgpExtendedCommunityIter struct {
	obj                       *bgpAttributes
	bgpExtendedCommunitySlice []BgpExtendedCommunity
	fieldPtr                  *[]*otg.BgpExtendedCommunity
}

func newBgpAttributesBgpExtendedCommunityIter(ptr *[]*otg.BgpExtendedCommunity) BgpAttributesBgpExtendedCommunityIter {
	return &bgpAttributesBgpExtendedCommunityIter{fieldPtr: ptr}
}

type BgpAttributesBgpExtendedCommunityIter interface {
	setMsg(*bgpAttributes) BgpAttributesBgpExtendedCommunityIter
	Items() []BgpExtendedCommunity
	Add() BgpExtendedCommunity
	Append(items ...BgpExtendedCommunity) BgpAttributesBgpExtendedCommunityIter
	Set(index int, newObj BgpExtendedCommunity) BgpAttributesBgpExtendedCommunityIter
	Clear() BgpAttributesBgpExtendedCommunityIter
	clearHolderSlice() BgpAttributesBgpExtendedCommunityIter
	appendHolderSlice(item BgpExtendedCommunity) BgpAttributesBgpExtendedCommunityIter
}

func (obj *bgpAttributesBgpExtendedCommunityIter) setMsg(msg *bgpAttributes) BgpAttributesBgpExtendedCommunityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpExtendedCommunity{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpAttributesBgpExtendedCommunityIter) Items() []BgpExtendedCommunity {
	return obj.bgpExtendedCommunitySlice
}

func (obj *bgpAttributesBgpExtendedCommunityIter) Add() BgpExtendedCommunity {
	newObj := &otg.BgpExtendedCommunity{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpExtendedCommunity{obj: newObj}
	newLibObj.setDefault()
	obj.bgpExtendedCommunitySlice = append(obj.bgpExtendedCommunitySlice, newLibObj)
	return newLibObj
}

func (obj *bgpAttributesBgpExtendedCommunityIter) Append(items ...BgpExtendedCommunity) BgpAttributesBgpExtendedCommunityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpExtendedCommunitySlice = append(obj.bgpExtendedCommunitySlice, item)
	}
	return obj
}

func (obj *bgpAttributesBgpExtendedCommunityIter) Set(index int, newObj BgpExtendedCommunity) BgpAttributesBgpExtendedCommunityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpExtendedCommunitySlice[index] = newObj
	return obj
}
func (obj *bgpAttributesBgpExtendedCommunityIter) Clear() BgpAttributesBgpExtendedCommunityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpExtendedCommunity{}
		obj.bgpExtendedCommunitySlice = []BgpExtendedCommunity{}
	}
	return obj
}
func (obj *bgpAttributesBgpExtendedCommunityIter) clearHolderSlice() BgpAttributesBgpExtendedCommunityIter {
	if len(obj.bgpExtendedCommunitySlice) > 0 {
		obj.bgpExtendedCommunitySlice = []BgpExtendedCommunity{}
	}
	return obj
}
func (obj *bgpAttributesBgpExtendedCommunityIter) appendHolderSlice(item BgpExtendedCommunity) BgpAttributesBgpExtendedCommunityIter {
	obj.bgpExtendedCommunitySlice = append(obj.bgpExtendedCommunitySlice, item)
	return obj
}

// description is TBD
// MpReach returns a BgpAttributesMpReachNlri
func (obj *bgpAttributes) MpReach() BgpAttributesMpReachNlri {
	if obj.obj.MpReach == nil {
		obj.obj.MpReach = NewBgpAttributesMpReachNlri().msg()
	}
	if obj.mpReachHolder == nil {
		obj.mpReachHolder = &bgpAttributesMpReachNlri{obj: obj.obj.MpReach}
	}
	return obj.mpReachHolder
}

// description is TBD
// MpReach returns a BgpAttributesMpReachNlri
func (obj *bgpAttributes) HasMpReach() bool {
	return obj.obj.MpReach != nil
}

// description is TBD
// SetMpReach sets the BgpAttributesMpReachNlri value in the BgpAttributes object
func (obj *bgpAttributes) SetMpReach(value BgpAttributesMpReachNlri) BgpAttributes {

	obj.mpReachHolder = nil
	obj.obj.MpReach = value.msg()

	return obj
}

// description is TBD
// MpUnreach returns a BgpAttributesMpUnreachNlri
func (obj *bgpAttributes) MpUnreach() BgpAttributesMpUnreachNlri {
	if obj.obj.MpUnreach == nil {
		obj.obj.MpUnreach = NewBgpAttributesMpUnreachNlri().msg()
	}
	if obj.mpUnreachHolder == nil {
		obj.mpUnreachHolder = &bgpAttributesMpUnreachNlri{obj: obj.obj.MpUnreach}
	}
	return obj.mpUnreachHolder
}

// description is TBD
// MpUnreach returns a BgpAttributesMpUnreachNlri
func (obj *bgpAttributes) HasMpUnreach() bool {
	return obj.obj.MpUnreach != nil
}

// description is TBD
// SetMpUnreach sets the BgpAttributesMpUnreachNlri value in the BgpAttributes object
func (obj *bgpAttributes) SetMpUnreach(value BgpAttributesMpUnreachNlri) BgpAttributes {

	obj.mpUnreachHolder = nil
	obj.obj.MpUnreach = value.msg()

	return obj
}

func (obj *bgpAttributes) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.OtherAttributes) != 0 {

		if set_default {
			obj.OtherAttributes().clearHolderSlice()
			for _, item := range obj.obj.OtherAttributes {
				obj.OtherAttributes().appendHolderSlice(&bgpAttributesOtherAttribute{obj: item})
			}
		}
		for _, item := range obj.OtherAttributes().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.AsPath != nil {

		obj.AsPath().validateObj(vObj, set_default)
	}

	if obj.obj.As4Path != nil {

		obj.As4Path().validateObj(vObj, set_default)
	}

	if obj.obj.NextHop != nil {

		obj.NextHop().validateObj(vObj, set_default)
	}

	if obj.obj.MultiExitDiscriminator != nil {

		obj.MultiExitDiscriminator().validateObj(vObj, set_default)
	}

	if obj.obj.LocalPreference != nil {

		obj.LocalPreference().validateObj(vObj, set_default)
	}

	if obj.obj.Aggregator != nil {

		obj.Aggregator().validateObj(vObj, set_default)
	}

	if obj.obj.As4Aggregator != nil {

		obj.As4Aggregator().validateObj(vObj, set_default)
	}

	if len(obj.obj.Community) != 0 {

		if set_default {
			obj.Community().clearHolderSlice()
			for _, item := range obj.obj.Community {
				obj.Community().appendHolderSlice(&bgpAttributesCommunity{obj: item})
			}
		}
		for _, item := range obj.Community().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.OriginatorId != nil {

		obj.OriginatorId().validateObj(vObj, set_default)
	}

	if obj.obj.ClusterIds != nil {

		err := obj.validateIpv4Slice(obj.ClusterIds())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpAttributes.ClusterIds"))
		}

	}

	if len(obj.obj.ExtendedCommunities) != 0 {

		if set_default {
			obj.ExtendedCommunities().clearHolderSlice()
			for _, item := range obj.obj.ExtendedCommunities {
				obj.ExtendedCommunities().appendHolderSlice(&bgpExtendedCommunity{obj: item})
			}
		}
		for _, item := range obj.ExtendedCommunities().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.MpReach != nil {

		obj.MpReach().validateObj(vObj, set_default)
	}

	if obj.obj.MpUnreach != nil {

		obj.MpUnreach().validateObj(vObj, set_default)
	}

}

func (obj *bgpAttributes) setDefault() {
	if obj.obj.Origin == nil {
		obj.SetOrigin(BgpAttributesOrigin.INCOMPLETE)

	}
	if obj.obj.IncludeAtomicAggregator == nil {
		obj.SetIncludeAtomicAggregator(false)
	}

}
