package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BmpServerPeerState *****
type bmpServerPeerState struct {
	validation
	obj                          *otg.BmpServerPeerState
	marshaller                   marshalBmpServerPeerState
	unMarshaller                 unMarshalBmpServerPeerState
	sessionIpv4InformationHolder BmpServerPeerSessionIpv4Information
	sessionIpv6InformationHolder BmpServerPeerSessionIpv6Information
	statsHolder                  BmpServerPeerStats
	prePolicyInRibHolder         BmpServerPeerPrePolicyInRib
	postPolicyInRibHolder        BmpServerPeerPostPolicyInRib
}

func NewBmpServerPeerState() BmpServerPeerState {
	obj := bmpServerPeerState{obj: &otg.BmpServerPeerState{}}
	obj.setDefault()
	return &obj
}

func (obj *bmpServerPeerState) msg() *otg.BmpServerPeerState {
	return obj.obj
}

func (obj *bmpServerPeerState) setMsg(msg *otg.BmpServerPeerState) BmpServerPeerState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbmpServerPeerState struct {
	obj *bmpServerPeerState
}

type marshalBmpServerPeerState interface {
	// ToProto marshals BmpServerPeerState to protobuf object *otg.BmpServerPeerState
	ToProto() (*otg.BmpServerPeerState, error)
	// ToPbText marshals BmpServerPeerState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BmpServerPeerState to YAML text
	ToYaml() (string, error)
	// ToJson marshals BmpServerPeerState to JSON text
	ToJson() (string, error)
}

type unMarshalbmpServerPeerState struct {
	obj *bmpServerPeerState
}

type unMarshalBmpServerPeerState interface {
	// FromProto unmarshals BmpServerPeerState from protobuf object *otg.BmpServerPeerState
	FromProto(msg *otg.BmpServerPeerState) (BmpServerPeerState, error)
	// FromPbText unmarshals BmpServerPeerState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BmpServerPeerState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BmpServerPeerState from JSON text
	FromJson(value string) error
}

func (obj *bmpServerPeerState) Marshal() marshalBmpServerPeerState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbmpServerPeerState{obj: obj}
	}
	return obj.marshaller
}

func (obj *bmpServerPeerState) Unmarshal() unMarshalBmpServerPeerState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbmpServerPeerState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbmpServerPeerState) ToProto() (*otg.BmpServerPeerState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbmpServerPeerState) FromProto(msg *otg.BmpServerPeerState) (BmpServerPeerState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbmpServerPeerState) ToPbText() (string, error) {
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

func (m *unMarshalbmpServerPeerState) FromPbText(value string) error {
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

func (m *marshalbmpServerPeerState) ToYaml() (string, error) {
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

func (m *unMarshalbmpServerPeerState) FromYaml(value string) error {
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

func (m *marshalbmpServerPeerState) ToJson() (string, error) {
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

func (m *unMarshalbmpServerPeerState) FromJson(value string) error {
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

func (obj *bmpServerPeerState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bmpServerPeerState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bmpServerPeerState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bmpServerPeerState) Clone() (BmpServerPeerState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBmpServerPeerState()
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

func (obj *bmpServerPeerState) setNil() {
	obj.sessionIpv4InformationHolder = nil
	obj.sessionIpv6InformationHolder = nil
	obj.statsHolder = nil
	obj.prePolicyInRibHolder = nil
	obj.postPolicyInRibHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BmpServerPeerState is the information returned by a BMP client to the BMP Server related to a specific BGP session on the router.
type BmpServerPeerState interface {
	Validation
	// msg marshals BmpServerPeerState to protobuf object *otg.BmpServerPeerState
	// and doesn't set defaults
	msg() *otg.BmpServerPeerState
	// setMsg unmarshals BmpServerPeerState from protobuf object *otg.BmpServerPeerState
	// and doesn't set defaults
	setMsg(*otg.BmpServerPeerState) BmpServerPeerState
	// provides marshal interface
	Marshal() marshalBmpServerPeerState
	// provides unmarshal interface
	Unmarshal() unMarshalBmpServerPeerState
	// validate validates BmpServerPeerState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BmpServerPeerState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SessionIpType returns BmpServerPeerStateSessionIpTypeEnum, set in BmpServerPeerState
	SessionIpType() BmpServerPeerStateSessionIpTypeEnum
	// SetSessionIpType assigns BmpServerPeerStateSessionIpTypeEnum provided by user to BmpServerPeerState
	SetSessionIpType(value BmpServerPeerStateSessionIpTypeEnum) BmpServerPeerState
	// HasSessionIpType checks if SessionIpType has been set in BmpServerPeerState
	HasSessionIpType() bool
	// SessionIpv4Information returns BmpServerPeerSessionIpv4Information, set in BmpServerPeerState.
	// BmpServerPeerSessionIpv4Information is this object is included if session_ip_type is of type ipv4.
	SessionIpv4Information() BmpServerPeerSessionIpv4Information
	// SetSessionIpv4Information assigns BmpServerPeerSessionIpv4Information provided by user to BmpServerPeerState.
	// BmpServerPeerSessionIpv4Information is this object is included if session_ip_type is of type ipv4.
	SetSessionIpv4Information(value BmpServerPeerSessionIpv4Information) BmpServerPeerState
	// HasSessionIpv4Information checks if SessionIpv4Information has been set in BmpServerPeerState
	HasSessionIpv4Information() bool
	// SessionIpv6Information returns BmpServerPeerSessionIpv6Information, set in BmpServerPeerState.
	// BmpServerPeerSessionIpv6Information is this object is included if session_ip_type is of type ipv6.
	SessionIpv6Information() BmpServerPeerSessionIpv6Information
	// SetSessionIpv6Information assigns BmpServerPeerSessionIpv6Information provided by user to BmpServerPeerState.
	// BmpServerPeerSessionIpv6Information is this object is included if session_ip_type is of type ipv6.
	SetSessionIpv6Information(value BmpServerPeerSessionIpv6Information) BmpServerPeerState
	// HasSessionIpv6Information checks if SessionIpv6Information has been set in BmpServerPeerState
	HasSessionIpv6Information() bool
	// PeerBgpId returns string, set in BmpServerPeerState.
	PeerBgpId() string
	// SetPeerBgpId assigns string provided by user to BmpServerPeerState
	SetPeerBgpId(value string) BmpServerPeerState
	// HasPeerBgpId checks if PeerBgpId has been set in BmpServerPeerState
	HasPeerBgpId() bool
	// Status returns BmpServerPeerStateStatusEnum, set in BmpServerPeerState
	Status() BmpServerPeerStateStatusEnum
	// SetStatus assigns BmpServerPeerStateStatusEnum provided by user to BmpServerPeerState
	SetStatus(value BmpServerPeerStateStatusEnum) BmpServerPeerState
	// HasStatus checks if Status has been set in BmpServerPeerState
	HasStatus() bool
	// Stats returns BmpServerPeerStats, set in BmpServerPeerState.
	// BmpServerPeerStats is the last set of BMP stats sent by the BMP client for this peer as per the message format defined in https://www.rfc-editor.org/rfc/rfc7854.html#section-4.8. It might be absent if no BMP Statistics message have been received yet or peer is still in down state. If the BMP client has sent a subset of the various statistic counters, only those should be set when returning the result.
	Stats() BmpServerPeerStats
	// SetStats assigns BmpServerPeerStats provided by user to BmpServerPeerState.
	// BmpServerPeerStats is the last set of BMP stats sent by the BMP client for this peer as per the message format defined in https://www.rfc-editor.org/rfc/rfc7854.html#section-4.8. It might be absent if no BMP Statistics message have been received yet or peer is still in down state. If the BMP client has sent a subset of the various statistic counters, only those should be set when returning the result.
	SetStats(value BmpServerPeerStats) BmpServerPeerState
	// HasStats checks if Stats has been set in BmpServerPeerState
	HasStats() bool
	// PrePolicyInRib returns BmpServerPeerPrePolicyInRib, set in BmpServerPeerState.
	// BmpServerPeerPrePolicyInRib is the set of routes received from this peer as reported using BMP Route Monitoring messages by the BMP client  which are part of the Adj-RIB-In database *before* inbound policies were applied. This is determined by checking that the L flag is not set in the Per-Peer Header in the received Route Monitoring message. It should also be ensured that the O flag as defined in RFC8671 is not set in the flags. (indicates Adj-RIB-Out)  Note that routes which have been advertised initially and currently in withdrawn state can be included in the returned set of routes by setting the status of that specific route to 'withdrawn'. Routes which have been received from this peer and part of current Pre-Policy In-Rib DB should have the status set to 'advertised'. Note that if prefix_storage.ipv4_unicast.discard or/and prefix_storage.ipv6_unicast.discard is configured or exceptions in the prefix_storage are specified such that all incoming routes are filtered, the corresponding prefix database list will be empty.
	PrePolicyInRib() BmpServerPeerPrePolicyInRib
	// SetPrePolicyInRib assigns BmpServerPeerPrePolicyInRib provided by user to BmpServerPeerState.
	// BmpServerPeerPrePolicyInRib is the set of routes received from this peer as reported using BMP Route Monitoring messages by the BMP client  which are part of the Adj-RIB-In database *before* inbound policies were applied. This is determined by checking that the L flag is not set in the Per-Peer Header in the received Route Monitoring message. It should also be ensured that the O flag as defined in RFC8671 is not set in the flags. (indicates Adj-RIB-Out)  Note that routes which have been advertised initially and currently in withdrawn state can be included in the returned set of routes by setting the status of that specific route to 'withdrawn'. Routes which have been received from this peer and part of current Pre-Policy In-Rib DB should have the status set to 'advertised'. Note that if prefix_storage.ipv4_unicast.discard or/and prefix_storage.ipv6_unicast.discard is configured or exceptions in the prefix_storage are specified such that all incoming routes are filtered, the corresponding prefix database list will be empty.
	SetPrePolicyInRib(value BmpServerPeerPrePolicyInRib) BmpServerPeerState
	// HasPrePolicyInRib checks if PrePolicyInRib has been set in BmpServerPeerState
	HasPrePolicyInRib() bool
	// PostPolicyInRib returns BmpServerPeerPostPolicyInRib, set in BmpServerPeerState.
	// BmpServerPeerPostPolicyInRib is the set of routes received from this peer as reported using BMP Route Monitoring messages by the BMP client  which are part of the Adj-RIB-In database *after* inbound policies were applied. This is determined by checking that the L flag is set in the Per-Peer Header in the received Route Monitoring message. It should also be ensured that the O flag as defined in RFC8671 is not set in the flags. (indicates Adj-RIB-Out)  Note that routes which have been advertised initially and currently in withdrawn state can be included in the returned set of routes by setting the status of that specific route to 'withdrawn'. Routes which have been received from this peer and part of current Post-Policy In-Rib DB should have the status set to 'advertised'. Note that if prefix_storage.ipv4_unicast.discard or/and prefix_storage.ipv6_unicast.discard is configured or exceptions in the prefix_storage are specified such that all incoming routes are filtered, the corresponding prefix database list will be empty. .
	PostPolicyInRib() BmpServerPeerPostPolicyInRib
	// SetPostPolicyInRib assigns BmpServerPeerPostPolicyInRib provided by user to BmpServerPeerState.
	// BmpServerPeerPostPolicyInRib is the set of routes received from this peer as reported using BMP Route Monitoring messages by the BMP client  which are part of the Adj-RIB-In database *after* inbound policies were applied. This is determined by checking that the L flag is set in the Per-Peer Header in the received Route Monitoring message. It should also be ensured that the O flag as defined in RFC8671 is not set in the flags. (indicates Adj-RIB-Out)  Note that routes which have been advertised initially and currently in withdrawn state can be included in the returned set of routes by setting the status of that specific route to 'withdrawn'. Routes which have been received from this peer and part of current Post-Policy In-Rib DB should have the status set to 'advertised'. Note that if prefix_storage.ipv4_unicast.discard or/and prefix_storage.ipv6_unicast.discard is configured or exceptions in the prefix_storage are specified such that all incoming routes are filtered, the corresponding prefix database list will be empty. .
	SetPostPolicyInRib(value BmpServerPeerPostPolicyInRib) BmpServerPeerState
	// HasPostPolicyInRib checks if PostPolicyInRib has been set in BmpServerPeerState
	HasPostPolicyInRib() bool
	setNil()
}

type BmpServerPeerStateSessionIpTypeEnum string

// Enum of SessionIpType on BmpServerPeerState
var BmpServerPeerStateSessionIpType = struct {
	IPV4 BmpServerPeerStateSessionIpTypeEnum
	IPV6 BmpServerPeerStateSessionIpTypeEnum
}{
	IPV4: BmpServerPeerStateSessionIpTypeEnum("ipv4"),
	IPV6: BmpServerPeerStateSessionIpTypeEnum("ipv6"),
}

func (obj *bmpServerPeerState) SessionIpType() BmpServerPeerStateSessionIpTypeEnum {
	return BmpServerPeerStateSessionIpTypeEnum(obj.obj.SessionIpType.Enum().String())
}

// Whether the peer is of type BGPv4 or BGPv6. If the type is ipv4 the session_ipv4_information object should be included to convey the local and remote IPv4 address for the peer. If the type is ipv6 the session_ipv6_information object should be included to convey the local and remote IPv6 address for the peer.
// SessionIpType returns a string
func (obj *bmpServerPeerState) HasSessionIpType() bool {
	return obj.obj.SessionIpType != nil
}

func (obj *bmpServerPeerState) SetSessionIpType(value BmpServerPeerStateSessionIpTypeEnum) BmpServerPeerState {
	intValue, ok := otg.BmpServerPeerState_SessionIpType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BmpServerPeerStateSessionIpTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.BmpServerPeerState_SessionIpType_Enum(intValue)
	obj.obj.SessionIpType = &enumValue

	return obj
}

// This object is included if session_ip_type is of type ipv4.
// SessionIpv4Information returns a BmpServerPeerSessionIpv4Information
func (obj *bmpServerPeerState) SessionIpv4Information() BmpServerPeerSessionIpv4Information {
	if obj.obj.SessionIpv4Information == nil {
		obj.obj.SessionIpv4Information = NewBmpServerPeerSessionIpv4Information().msg()
	}
	if obj.sessionIpv4InformationHolder == nil {
		obj.sessionIpv4InformationHolder = &bmpServerPeerSessionIpv4Information{obj: obj.obj.SessionIpv4Information}
	}
	return obj.sessionIpv4InformationHolder
}

// This object is included if session_ip_type is of type ipv4.
// SessionIpv4Information returns a BmpServerPeerSessionIpv4Information
func (obj *bmpServerPeerState) HasSessionIpv4Information() bool {
	return obj.obj.SessionIpv4Information != nil
}

// This object is included if session_ip_type is of type ipv4.
// SetSessionIpv4Information sets the BmpServerPeerSessionIpv4Information value in the BmpServerPeerState object
func (obj *bmpServerPeerState) SetSessionIpv4Information(value BmpServerPeerSessionIpv4Information) BmpServerPeerState {

	obj.sessionIpv4InformationHolder = nil
	obj.obj.SessionIpv4Information = value.msg()

	return obj
}

// This object is included if session_ip_type is of type ipv6.
// SessionIpv6Information returns a BmpServerPeerSessionIpv6Information
func (obj *bmpServerPeerState) SessionIpv6Information() BmpServerPeerSessionIpv6Information {
	if obj.obj.SessionIpv6Information == nil {
		obj.obj.SessionIpv6Information = NewBmpServerPeerSessionIpv6Information().msg()
	}
	if obj.sessionIpv6InformationHolder == nil {
		obj.sessionIpv6InformationHolder = &bmpServerPeerSessionIpv6Information{obj: obj.obj.SessionIpv6Information}
	}
	return obj.sessionIpv6InformationHolder
}

// This object is included if session_ip_type is of type ipv6.
// SessionIpv6Information returns a BmpServerPeerSessionIpv6Information
func (obj *bmpServerPeerState) HasSessionIpv6Information() bool {
	return obj.obj.SessionIpv6Information != nil
}

// This object is included if session_ip_type is of type ipv6.
// SetSessionIpv6Information sets the BmpServerPeerSessionIpv6Information value in the BmpServerPeerState object
func (obj *bmpServerPeerState) SetSessionIpv6Information(value BmpServerPeerSessionIpv6Information) BmpServerPeerState {

	obj.sessionIpv6InformationHolder = nil
	obj.obj.SessionIpv6Information = value.msg()

	return obj
}

// The BGP Identifier of the peer.
// PeerBgpId returns a string
func (obj *bmpServerPeerState) PeerBgpId() string {

	return *obj.obj.PeerBgpId

}

// The BGP Identifier of the peer.
// PeerBgpId returns a string
func (obj *bmpServerPeerState) HasPeerBgpId() bool {
	return obj.obj.PeerBgpId != nil
}

// The BGP Identifier of the peer.
// SetPeerBgpId sets the string value in the BmpServerPeerState object
func (obj *bmpServerPeerState) SetPeerBgpId(value string) BmpServerPeerState {

	obj.obj.PeerBgpId = &value
	return obj
}

type BmpServerPeerStateStatusEnum string

// Enum of Status on BmpServerPeerState
var BmpServerPeerStateStatus = struct {
	UP   BmpServerPeerStateStatusEnum
	DOWN BmpServerPeerStateStatusEnum
}{
	UP:   BmpServerPeerStateStatusEnum("up"),
	DOWN: BmpServerPeerStateStatusEnum("down"),
}

func (obj *bmpServerPeerState) Status() BmpServerPeerStateStatusEnum {
	return BmpServerPeerStateStatusEnum(obj.obj.Status.Enum().String())
}

// The current state of the BGP session.  Note that per peer stats and rib information might be empty if the peer state is down.
// Status returns a string
func (obj *bmpServerPeerState) HasStatus() bool {
	return obj.obj.Status != nil
}

func (obj *bmpServerPeerState) SetStatus(value BmpServerPeerStateStatusEnum) BmpServerPeerState {
	intValue, ok := otg.BmpServerPeerState_Status_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BmpServerPeerStateStatusEnum", string(value)))
		return obj
	}
	enumValue := otg.BmpServerPeerState_Status_Enum(intValue)
	obj.obj.Status = &enumValue

	return obj
}

// The last set of BMP stats sent by the BMP client for this peer as per the message format defined in https://www.rfc-editor.org/rfc/rfc7854.html#section-4.8. It might be absent if no BMP Statistics message have been received yet or peer is still in down state. If the BMP client has sent a subset of the various statistic counters, only those should be set when returning the result.
// Stats returns a BmpServerPeerStats
func (obj *bmpServerPeerState) Stats() BmpServerPeerStats {
	if obj.obj.Stats == nil {
		obj.obj.Stats = NewBmpServerPeerStats().msg()
	}
	if obj.statsHolder == nil {
		obj.statsHolder = &bmpServerPeerStats{obj: obj.obj.Stats}
	}
	return obj.statsHolder
}

// The last set of BMP stats sent by the BMP client for this peer as per the message format defined in https://www.rfc-editor.org/rfc/rfc7854.html#section-4.8. It might be absent if no BMP Statistics message have been received yet or peer is still in down state. If the BMP client has sent a subset of the various statistic counters, only those should be set when returning the result.
// Stats returns a BmpServerPeerStats
func (obj *bmpServerPeerState) HasStats() bool {
	return obj.obj.Stats != nil
}

// The last set of BMP stats sent by the BMP client for this peer as per the message format defined in https://www.rfc-editor.org/rfc/rfc7854.html#section-4.8. It might be absent if no BMP Statistics message have been received yet or peer is still in down state. If the BMP client has sent a subset of the various statistic counters, only those should be set when returning the result.
// SetStats sets the BmpServerPeerStats value in the BmpServerPeerState object
func (obj *bmpServerPeerState) SetStats(value BmpServerPeerStats) BmpServerPeerState {

	obj.statsHolder = nil
	obj.obj.Stats = value.msg()

	return obj
}

// The set of routes received from this peer as reported using BMP Route Monitoring messages by the BMP client  which are part of the Adj-RIB-In database *before* inbound policies were applied. This is determined by checking that the L flag is not set in the Per-Peer Header in the received Route Monitoring message. It should also be ensured that the O flag as defined in RFC8671 is not set in the flags. (indicates Adj-RIB-Out)  Note that routes which have been advertised initially and currently in withdrawn state can be included in the returned set of routes by setting the status of that specific route to 'withdrawn'. Routes which have been received from this peer and part of current Pre-Policy In-Rib DB should have the status set to 'advertised'.             Note that if prefix_storage.ipv4_unicast.discard or/and prefix_storage.ipv6_unicast.discard is configured or exceptions in the prefix_storage are specified such that all incoming routes are filtered, the prefix database list will be empty.
// PrePolicyInRib returns a BmpServerPeerPrePolicyInRib
func (obj *bmpServerPeerState) PrePolicyInRib() BmpServerPeerPrePolicyInRib {
	if obj.obj.PrePolicyInRib == nil {
		obj.obj.PrePolicyInRib = NewBmpServerPeerPrePolicyInRib().msg()
	}
	if obj.prePolicyInRibHolder == nil {
		obj.prePolicyInRibHolder = &bmpServerPeerPrePolicyInRib{obj: obj.obj.PrePolicyInRib}
	}
	return obj.prePolicyInRibHolder
}

// The set of routes received from this peer as reported using BMP Route Monitoring messages by the BMP client  which are part of the Adj-RIB-In database *before* inbound policies were applied. This is determined by checking that the L flag is not set in the Per-Peer Header in the received Route Monitoring message. It should also be ensured that the O flag as defined in RFC8671 is not set in the flags. (indicates Adj-RIB-Out)  Note that routes which have been advertised initially and currently in withdrawn state can be included in the returned set of routes by setting the status of that specific route to 'withdrawn'. Routes which have been received from this peer and part of current Pre-Policy In-Rib DB should have the status set to 'advertised'.             Note that if prefix_storage.ipv4_unicast.discard or/and prefix_storage.ipv6_unicast.discard is configured or exceptions in the prefix_storage are specified such that all incoming routes are filtered, the prefix database list will be empty.
// PrePolicyInRib returns a BmpServerPeerPrePolicyInRib
func (obj *bmpServerPeerState) HasPrePolicyInRib() bool {
	return obj.obj.PrePolicyInRib != nil
}

// The set of routes received from this peer as reported using BMP Route Monitoring messages by the BMP client  which are part of the Adj-RIB-In database *before* inbound policies were applied. This is determined by checking that the L flag is not set in the Per-Peer Header in the received Route Monitoring message. It should also be ensured that the O flag as defined in RFC8671 is not set in the flags. (indicates Adj-RIB-Out)  Note that routes which have been advertised initially and currently in withdrawn state can be included in the returned set of routes by setting the status of that specific route to 'withdrawn'. Routes which have been received from this peer and part of current Pre-Policy In-Rib DB should have the status set to 'advertised'.             Note that if prefix_storage.ipv4_unicast.discard or/and prefix_storage.ipv6_unicast.discard is configured or exceptions in the prefix_storage are specified such that all incoming routes are filtered, the prefix database list will be empty.
// SetPrePolicyInRib sets the BmpServerPeerPrePolicyInRib value in the BmpServerPeerState object
func (obj *bmpServerPeerState) SetPrePolicyInRib(value BmpServerPeerPrePolicyInRib) BmpServerPeerState {

	obj.prePolicyInRibHolder = nil
	obj.obj.PrePolicyInRib = value.msg()

	return obj
}

// The set of routes received from this peer as reported using BMP Route Monitoring messages by the BMP client  which are part of the Adj-RIB-In database *after* inbound policies were applied. This is determined by checking that the L flag is set in the Per-Peer Header in the received Route Monitoring message. It should also be ensured that the O flag as defined in RFC8671 is not set in the flags. (indicates Adj-RIB-Out)  Note that routes which have been advertised initially and currently in withdrawn state can be included in the returned set of routes by setting the status of that specific route to 'withdrawn'. Routes which have been received from this peer and part of current Post-Policy In-Rib DB should have the status set to 'advertised'. Note that if prefix_storage.ipv4_unicast.discard or/and prefix_storage.ipv6_unicast.discard is configured or exceptions in the prefix_storage are specified such that all incoming routes are filtered, the prefix database list will be empty.
// PostPolicyInRib returns a BmpServerPeerPostPolicyInRib
func (obj *bmpServerPeerState) PostPolicyInRib() BmpServerPeerPostPolicyInRib {
	if obj.obj.PostPolicyInRib == nil {
		obj.obj.PostPolicyInRib = NewBmpServerPeerPostPolicyInRib().msg()
	}
	if obj.postPolicyInRibHolder == nil {
		obj.postPolicyInRibHolder = &bmpServerPeerPostPolicyInRib{obj: obj.obj.PostPolicyInRib}
	}
	return obj.postPolicyInRibHolder
}

// The set of routes received from this peer as reported using BMP Route Monitoring messages by the BMP client  which are part of the Adj-RIB-In database *after* inbound policies were applied. This is determined by checking that the L flag is set in the Per-Peer Header in the received Route Monitoring message. It should also be ensured that the O flag as defined in RFC8671 is not set in the flags. (indicates Adj-RIB-Out)  Note that routes which have been advertised initially and currently in withdrawn state can be included in the returned set of routes by setting the status of that specific route to 'withdrawn'. Routes which have been received from this peer and part of current Post-Policy In-Rib DB should have the status set to 'advertised'. Note that if prefix_storage.ipv4_unicast.discard or/and prefix_storage.ipv6_unicast.discard is configured or exceptions in the prefix_storage are specified such that all incoming routes are filtered, the prefix database list will be empty.
// PostPolicyInRib returns a BmpServerPeerPostPolicyInRib
func (obj *bmpServerPeerState) HasPostPolicyInRib() bool {
	return obj.obj.PostPolicyInRib != nil
}

// The set of routes received from this peer as reported using BMP Route Monitoring messages by the BMP client  which are part of the Adj-RIB-In database *after* inbound policies were applied. This is determined by checking that the L flag is set in the Per-Peer Header in the received Route Monitoring message. It should also be ensured that the O flag as defined in RFC8671 is not set in the flags. (indicates Adj-RIB-Out)  Note that routes which have been advertised initially and currently in withdrawn state can be included in the returned set of routes by setting the status of that specific route to 'withdrawn'. Routes which have been received from this peer and part of current Post-Policy In-Rib DB should have the status set to 'advertised'. Note that if prefix_storage.ipv4_unicast.discard or/and prefix_storage.ipv6_unicast.discard is configured or exceptions in the prefix_storage are specified such that all incoming routes are filtered, the prefix database list will be empty.
// SetPostPolicyInRib sets the BmpServerPeerPostPolicyInRib value in the BmpServerPeerState object
func (obj *bmpServerPeerState) SetPostPolicyInRib(value BmpServerPeerPostPolicyInRib) BmpServerPeerState {

	obj.postPolicyInRibHolder = nil
	obj.obj.PostPolicyInRib = value.msg()

	return obj
}

func (obj *bmpServerPeerState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SessionIpv4Information != nil {

		obj.SessionIpv4Information().validateObj(vObj, set_default)
	}

	if obj.obj.SessionIpv6Information != nil {

		obj.SessionIpv6Information().validateObj(vObj, set_default)
	}

	if obj.obj.PeerBgpId != nil {

		err := obj.validateIpv4(obj.PeerBgpId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BmpServerPeerState.PeerBgpId"))
		}

	}

	if obj.obj.Stats != nil {

		obj.Stats().validateObj(vObj, set_default)
	}

	if obj.obj.PrePolicyInRib != nil {

		obj.PrePolicyInRib().validateObj(vObj, set_default)
	}

	if obj.obj.PostPolicyInRib != nil {

		obj.PostPolicyInRib().validateObj(vObj, set_default)
	}

}

func (obj *bmpServerPeerState) setDefault() {

}
