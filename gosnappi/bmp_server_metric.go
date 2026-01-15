package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BmpServerMetric *****
type bmpServerMetric struct {
	validation
	obj          *otg.BmpServerMetric
	marshaller   marshalBmpServerMetric
	unMarshaller unMarshalBmpServerMetric
}

func NewBmpServerMetric() BmpServerMetric {
	obj := bmpServerMetric{obj: &otg.BmpServerMetric{}}
	obj.setDefault()
	return &obj
}

func (obj *bmpServerMetric) msg() *otg.BmpServerMetric {
	return obj.obj
}

func (obj *bmpServerMetric) setMsg(msg *otg.BmpServerMetric) BmpServerMetric {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbmpServerMetric struct {
	obj *bmpServerMetric
}

type marshalBmpServerMetric interface {
	// ToProto marshals BmpServerMetric to protobuf object *otg.BmpServerMetric
	ToProto() (*otg.BmpServerMetric, error)
	// ToPbText marshals BmpServerMetric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BmpServerMetric to YAML text
	ToYaml() (string, error)
	// ToJson marshals BmpServerMetric to JSON text
	ToJson() (string, error)
}

type unMarshalbmpServerMetric struct {
	obj *bmpServerMetric
}

type unMarshalBmpServerMetric interface {
	// FromProto unmarshals BmpServerMetric from protobuf object *otg.BmpServerMetric
	FromProto(msg *otg.BmpServerMetric) (BmpServerMetric, error)
	// FromPbText unmarshals BmpServerMetric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BmpServerMetric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BmpServerMetric from JSON text
	FromJson(value string) error
}

func (obj *bmpServerMetric) Marshal() marshalBmpServerMetric {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbmpServerMetric{obj: obj}
	}
	return obj.marshaller
}

func (obj *bmpServerMetric) Unmarshal() unMarshalBmpServerMetric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbmpServerMetric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbmpServerMetric) ToProto() (*otg.BmpServerMetric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbmpServerMetric) FromProto(msg *otg.BmpServerMetric) (BmpServerMetric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbmpServerMetric) ToPbText() (string, error) {
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

func (m *unMarshalbmpServerMetric) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalbmpServerMetric) ToYaml() (string, error) {
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

func (m *unMarshalbmpServerMetric) FromYaml(value string) error {
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

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalbmpServerMetric) ToJson() (string, error) {
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

func (m *unMarshalbmpServerMetric) FromJson(value string) error {
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

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *bmpServerMetric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bmpServerMetric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bmpServerMetric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bmpServerMetric) Clone() (BmpServerMetric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBmpServerMetric()
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

// BmpServerMetric is bMP Server statistics information.
type BmpServerMetric interface {
	Validation
	// msg marshals BmpServerMetric to protobuf object *otg.BmpServerMetric
	// and doesn't set defaults
	msg() *otg.BmpServerMetric
	// setMsg unmarshals BmpServerMetric from protobuf object *otg.BmpServerMetric
	// and doesn't set defaults
	setMsg(*otg.BmpServerMetric) BmpServerMetric
	// provides marshal interface
	Marshal() marshalBmpServerMetric
	// provides unmarshal interface
	Unmarshal() unMarshalBmpServerMetric
	// validate validates BmpServerMetric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BmpServerMetric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in BmpServerMetric.
	Name() string
	// SetName assigns string provided by user to BmpServerMetric
	SetName(value string) BmpServerMetric
	// HasName checks if Name has been set in BmpServerMetric
	HasName() bool
	// SessionState returns BmpServerMetricSessionStateEnum, set in BmpServerMetric
	SessionState() BmpServerMetricSessionStateEnum
	// SetSessionState assigns BmpServerMetricSessionStateEnum provided by user to BmpServerMetric
	SetSessionState(value BmpServerMetricSessionStateEnum) BmpServerMetric
	// HasSessionState checks if SessionState has been set in BmpServerMetric
	HasSessionState() bool
	// FlapCount returns uint64, set in BmpServerMetric.
	FlapCount() uint64
	// SetFlapCount assigns uint64 provided by user to BmpServerMetric
	SetFlapCount(value uint64) BmpServerMetric
	// HasFlapCount checks if FlapCount has been set in BmpServerMetric
	HasFlapCount() bool
	// RouteMonitoringMessagesReceived returns uint64, set in BmpServerMetric.
	RouteMonitoringMessagesReceived() uint64
	// SetRouteMonitoringMessagesReceived assigns uint64 provided by user to BmpServerMetric
	SetRouteMonitoringMessagesReceived(value uint64) BmpServerMetric
	// HasRouteMonitoringMessagesReceived checks if RouteMonitoringMessagesReceived has been set in BmpServerMetric
	HasRouteMonitoringMessagesReceived() bool
	// StatisticsMessagesReceived returns uint64, set in BmpServerMetric.
	StatisticsMessagesReceived() uint64
	// SetStatisticsMessagesReceived assigns uint64 provided by user to BmpServerMetric
	SetStatisticsMessagesReceived(value uint64) BmpServerMetric
	// HasStatisticsMessagesReceived checks if StatisticsMessagesReceived has been set in BmpServerMetric
	HasStatisticsMessagesReceived() bool
	// PeerDownMessagesReceived returns uint64, set in BmpServerMetric.
	PeerDownMessagesReceived() uint64
	// SetPeerDownMessagesReceived assigns uint64 provided by user to BmpServerMetric
	SetPeerDownMessagesReceived(value uint64) BmpServerMetric
	// HasPeerDownMessagesReceived checks if PeerDownMessagesReceived has been set in BmpServerMetric
	HasPeerDownMessagesReceived() bool
	// PeerUpMessagesReceived returns uint64, set in BmpServerMetric.
	PeerUpMessagesReceived() uint64
	// SetPeerUpMessagesReceived assigns uint64 provided by user to BmpServerMetric
	SetPeerUpMessagesReceived(value uint64) BmpServerMetric
	// HasPeerUpMessagesReceived checks if PeerUpMessagesReceived has been set in BmpServerMetric
	HasPeerUpMessagesReceived() bool
	// InitiationMessagesReceived returns uint64, set in BmpServerMetric.
	InitiationMessagesReceived() uint64
	// SetInitiationMessagesReceived assigns uint64 provided by user to BmpServerMetric
	SetInitiationMessagesReceived(value uint64) BmpServerMetric
	// HasInitiationMessagesReceived checks if InitiationMessagesReceived has been set in BmpServerMetric
	HasInitiationMessagesReceived() bool
	// RouteMirroringMessagesReceived returns uint64, set in BmpServerMetric.
	RouteMirroringMessagesReceived() uint64
	// SetRouteMirroringMessagesReceived assigns uint64 provided by user to BmpServerMetric
	SetRouteMirroringMessagesReceived(value uint64) BmpServerMetric
	// HasRouteMirroringMessagesReceived checks if RouteMirroringMessagesReceived has been set in BmpServerMetric
	HasRouteMirroringMessagesReceived() bool
	// TerminationMessagesReceived returns uint64, set in BmpServerMetric.
	TerminationMessagesReceived() uint64
	// SetTerminationMessagesReceived assigns uint64 provided by user to BmpServerMetric
	SetTerminationMessagesReceived(value uint64) BmpServerMetric
	// HasTerminationMessagesReceived checks if TerminationMessagesReceived has been set in BmpServerMetric
	HasTerminationMessagesReceived() bool
	// PrePolicyIpv4UnicastRoutesReceived returns uint64, set in BmpServerMetric.
	PrePolicyIpv4UnicastRoutesReceived() uint64
	// SetPrePolicyIpv4UnicastRoutesReceived assigns uint64 provided by user to BmpServerMetric
	SetPrePolicyIpv4UnicastRoutesReceived(value uint64) BmpServerMetric
	// HasPrePolicyIpv4UnicastRoutesReceived checks if PrePolicyIpv4UnicastRoutesReceived has been set in BmpServerMetric
	HasPrePolicyIpv4UnicastRoutesReceived() bool
	// PostPolicyIpv4UnicastRoutesReceived returns uint64, set in BmpServerMetric.
	PostPolicyIpv4UnicastRoutesReceived() uint64
	// SetPostPolicyIpv4UnicastRoutesReceived assigns uint64 provided by user to BmpServerMetric
	SetPostPolicyIpv4UnicastRoutesReceived(value uint64) BmpServerMetric
	// HasPostPolicyIpv4UnicastRoutesReceived checks if PostPolicyIpv4UnicastRoutesReceived has been set in BmpServerMetric
	HasPostPolicyIpv4UnicastRoutesReceived() bool
	// PrePolicyIpv6UnicastRoutesReceived returns uint64, set in BmpServerMetric.
	PrePolicyIpv6UnicastRoutesReceived() uint64
	// SetPrePolicyIpv6UnicastRoutesReceived assigns uint64 provided by user to BmpServerMetric
	SetPrePolicyIpv6UnicastRoutesReceived(value uint64) BmpServerMetric
	// HasPrePolicyIpv6UnicastRoutesReceived checks if PrePolicyIpv6UnicastRoutesReceived has been set in BmpServerMetric
	HasPrePolicyIpv6UnicastRoutesReceived() bool
	// PostPolicyIpv6UnicastRoutesReceived returns uint64, set in BmpServerMetric.
	PostPolicyIpv6UnicastRoutesReceived() uint64
	// SetPostPolicyIpv6UnicastRoutesReceived assigns uint64 provided by user to BmpServerMetric
	SetPostPolicyIpv6UnicastRoutesReceived(value uint64) BmpServerMetric
	// HasPostPolicyIpv6UnicastRoutesReceived checks if PostPolicyIpv6UnicastRoutesReceived has been set in BmpServerMetric
	HasPostPolicyIpv6UnicastRoutesReceived() bool
}

// The name of the configured BMP server.
// Name returns a string
func (obj *bmpServerMetric) Name() string {

	return *obj.obj.Name

}

// The name of the configured BMP server.
// Name returns a string
func (obj *bmpServerMetric) HasName() bool {
	return obj.obj.Name != nil
}

// The name of the configured BMP server.
// SetName sets the string value in the BmpServerMetric object
func (obj *bmpServerMetric) SetName(value string) BmpServerMetric {

	obj.obj.Name = &value
	return obj
}

type BmpServerMetricSessionStateEnum string

// Enum of SessionState on BmpServerMetric
var BmpServerMetricSessionState = struct {
	UP   BmpServerMetricSessionStateEnum
	DOWN BmpServerMetricSessionStateEnum
}{
	UP:   BmpServerMetricSessionStateEnum("up"),
	DOWN: BmpServerMetricSessionStateEnum("down"),
}

func (obj *bmpServerMetric) SessionState() BmpServerMetricSessionStateEnum {
	return BmpServerMetricSessionStateEnum(obj.obj.SessionState.Enum().String())
}

// Session state as up or down. The session goes to 'up' state when a BMP monitored client router sends a Initiation message to the BMP Server. The session goes to 'down' state when a BMP monitored client router sends a Termination message to the BMP Server or the TCP connection is lost. On initial start, the state as reported as 'down'.
// SessionState returns a string
func (obj *bmpServerMetric) HasSessionState() bool {
	return obj.obj.SessionState != nil
}

func (obj *bmpServerMetric) SetSessionState(value BmpServerMetricSessionStateEnum) BmpServerMetric {
	intValue, ok := otg.BmpServerMetric_SessionState_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BmpServerMetricSessionStateEnum", string(value)))
		return obj
	}
	enumValue := otg.BmpServerMetric_SessionState_Enum(intValue)
	obj.obj.SessionState = &enumValue

	return obj
}

// Number of times the BMP session with the BMP client went from Up to Down state.
// FlapCount returns a uint64
func (obj *bmpServerMetric) FlapCount() uint64 {

	return *obj.obj.FlapCount

}

// Number of times the BMP session with the BMP client went from Up to Down state.
// FlapCount returns a uint64
func (obj *bmpServerMetric) HasFlapCount() bool {
	return obj.obj.FlapCount != nil
}

// Number of times the BMP session with the BMP client went from Up to Down state.
// SetFlapCount sets the uint64 value in the BmpServerMetric object
func (obj *bmpServerMetric) SetFlapCount(value uint64) BmpServerMetric {

	obj.obj.FlapCount = &value
	return obj
}

// Number of BMP Route Monitoring Messages received from the BMP client from the time the BMP server was started.
// RouteMonitoringMessagesReceived returns a uint64
func (obj *bmpServerMetric) RouteMonitoringMessagesReceived() uint64 {

	return *obj.obj.RouteMonitoringMessagesReceived

}

// Number of BMP Route Monitoring Messages received from the BMP client from the time the BMP server was started.
// RouteMonitoringMessagesReceived returns a uint64
func (obj *bmpServerMetric) HasRouteMonitoringMessagesReceived() bool {
	return obj.obj.RouteMonitoringMessagesReceived != nil
}

// Number of BMP Route Monitoring Messages received from the BMP client from the time the BMP server was started.
// SetRouteMonitoringMessagesReceived sets the uint64 value in the BmpServerMetric object
func (obj *bmpServerMetric) SetRouteMonitoringMessagesReceived(value uint64) BmpServerMetric {

	obj.obj.RouteMonitoringMessagesReceived = &value
	return obj
}

// Number of BMP Route Statistics Messages received from the BMP client from the time the BMP server was started.
// StatisticsMessagesReceived returns a uint64
func (obj *bmpServerMetric) StatisticsMessagesReceived() uint64 {

	return *obj.obj.StatisticsMessagesReceived

}

// Number of BMP Route Statistics Messages received from the BMP client from the time the BMP server was started.
// StatisticsMessagesReceived returns a uint64
func (obj *bmpServerMetric) HasStatisticsMessagesReceived() bool {
	return obj.obj.StatisticsMessagesReceived != nil
}

// Number of BMP Route Statistics Messages received from the BMP client from the time the BMP server was started.
// SetStatisticsMessagesReceived sets the uint64 value in the BmpServerMetric object
func (obj *bmpServerMetric) SetStatisticsMessagesReceived(value uint64) BmpServerMetric {

	obj.obj.StatisticsMessagesReceived = &value
	return obj
}

// Number of BMP Peer Down Messages received from the BMP client from the time the BMP server was started.
// PeerDownMessagesReceived returns a uint64
func (obj *bmpServerMetric) PeerDownMessagesReceived() uint64 {

	return *obj.obj.PeerDownMessagesReceived

}

// Number of BMP Peer Down Messages received from the BMP client from the time the BMP server was started.
// PeerDownMessagesReceived returns a uint64
func (obj *bmpServerMetric) HasPeerDownMessagesReceived() bool {
	return obj.obj.PeerDownMessagesReceived != nil
}

// Number of BMP Peer Down Messages received from the BMP client from the time the BMP server was started.
// SetPeerDownMessagesReceived sets the uint64 value in the BmpServerMetric object
func (obj *bmpServerMetric) SetPeerDownMessagesReceived(value uint64) BmpServerMetric {

	obj.obj.PeerDownMessagesReceived = &value
	return obj
}

// Number of BMP Peer Up Messages received from the BMP client from the time the BMP server was started.
// PeerUpMessagesReceived returns a uint64
func (obj *bmpServerMetric) PeerUpMessagesReceived() uint64 {

	return *obj.obj.PeerUpMessagesReceived

}

// Number of BMP Peer Up Messages received from the BMP client from the time the BMP server was started.
// PeerUpMessagesReceived returns a uint64
func (obj *bmpServerMetric) HasPeerUpMessagesReceived() bool {
	return obj.obj.PeerUpMessagesReceived != nil
}

// Number of BMP Peer Up Messages received from the BMP client from the time the BMP server was started.
// SetPeerUpMessagesReceived sets the uint64 value in the BmpServerMetric object
func (obj *bmpServerMetric) SetPeerUpMessagesReceived(value uint64) BmpServerMetric {

	obj.obj.PeerUpMessagesReceived = &value
	return obj
}

// Number of BMP Initiation Messages received from the BMP client from the time the BMP server was started .
// InitiationMessagesReceived returns a uint64
func (obj *bmpServerMetric) InitiationMessagesReceived() uint64 {

	return *obj.obj.InitiationMessagesReceived

}

// Number of BMP Initiation Messages received from the BMP client from the time the BMP server was started .
// InitiationMessagesReceived returns a uint64
func (obj *bmpServerMetric) HasInitiationMessagesReceived() bool {
	return obj.obj.InitiationMessagesReceived != nil
}

// Number of BMP Initiation Messages received from the BMP client from the time the BMP server was started .
// SetInitiationMessagesReceived sets the uint64 value in the BmpServerMetric object
func (obj *bmpServerMetric) SetInitiationMessagesReceived(value uint64) BmpServerMetric {

	obj.obj.InitiationMessagesReceived = &value
	return obj
}

// Number of BMP Route Mirroring Messages received from the BMP client from the time the BMP server was started.
// RouteMirroringMessagesReceived returns a uint64
func (obj *bmpServerMetric) RouteMirroringMessagesReceived() uint64 {

	return *obj.obj.RouteMirroringMessagesReceived

}

// Number of BMP Route Mirroring Messages received from the BMP client from the time the BMP server was started.
// RouteMirroringMessagesReceived returns a uint64
func (obj *bmpServerMetric) HasRouteMirroringMessagesReceived() bool {
	return obj.obj.RouteMirroringMessagesReceived != nil
}

// Number of BMP Route Mirroring Messages received from the BMP client from the time the BMP server was started.
// SetRouteMirroringMessagesReceived sets the uint64 value in the BmpServerMetric object
func (obj *bmpServerMetric) SetRouteMirroringMessagesReceived(value uint64) BmpServerMetric {

	obj.obj.RouteMirroringMessagesReceived = &value
	return obj
}

// Number of BMP Termination Messages received from the BMP client from the time the BMP server was started.
// TerminationMessagesReceived returns a uint64
func (obj *bmpServerMetric) TerminationMessagesReceived() uint64 {

	return *obj.obj.TerminationMessagesReceived

}

// Number of BMP Termination Messages received from the BMP client from the time the BMP server was started.
// TerminationMessagesReceived returns a uint64
func (obj *bmpServerMetric) HasTerminationMessagesReceived() bool {
	return obj.obj.TerminationMessagesReceived != nil
}

// Number of BMP Termination Messages received from the BMP client from the time the BMP server was started.
// SetTerminationMessagesReceived sets the uint64 value in the BmpServerMetric object
func (obj *bmpServerMetric) SetTerminationMessagesReceived(value uint64) BmpServerMetric {

	obj.obj.TerminationMessagesReceived = &value
	return obj
}

// Number of IPv4 Unicast prefixes received within BMP Monitor messages with L flag unset (indicating absence of application on in-bound policy on the advertised prefix)  from the BMP client from the time the BMP server was started.
// PrePolicyIpv4UnicastRoutesReceived returns a uint64
func (obj *bmpServerMetric) PrePolicyIpv4UnicastRoutesReceived() uint64 {

	return *obj.obj.PrePolicyIpv4UnicastRoutesReceived

}

// Number of IPv4 Unicast prefixes received within BMP Monitor messages with L flag unset (indicating absence of application on in-bound policy on the advertised prefix)  from the BMP client from the time the BMP server was started.
// PrePolicyIpv4UnicastRoutesReceived returns a uint64
func (obj *bmpServerMetric) HasPrePolicyIpv4UnicastRoutesReceived() bool {
	return obj.obj.PrePolicyIpv4UnicastRoutesReceived != nil
}

// Number of IPv4 Unicast prefixes received within BMP Monitor messages with L flag unset (indicating absence of application on in-bound policy on the advertised prefix)  from the BMP client from the time the BMP server was started.
// SetPrePolicyIpv4UnicastRoutesReceived sets the uint64 value in the BmpServerMetric object
func (obj *bmpServerMetric) SetPrePolicyIpv4UnicastRoutesReceived(value uint64) BmpServerMetric {

	obj.obj.PrePolicyIpv4UnicastRoutesReceived = &value
	return obj
}

// Number of IPv4 Unicast prefixes received within BMP Monitor messages with L flag set (indicating application of in-bound policy on the advertised prefix)  from the BMP client from the time the BMP server was started.
// PostPolicyIpv4UnicastRoutesReceived returns a uint64
func (obj *bmpServerMetric) PostPolicyIpv4UnicastRoutesReceived() uint64 {

	return *obj.obj.PostPolicyIpv4UnicastRoutesReceived

}

// Number of IPv4 Unicast prefixes received within BMP Monitor messages with L flag set (indicating application of in-bound policy on the advertised prefix)  from the BMP client from the time the BMP server was started.
// PostPolicyIpv4UnicastRoutesReceived returns a uint64
func (obj *bmpServerMetric) HasPostPolicyIpv4UnicastRoutesReceived() bool {
	return obj.obj.PostPolicyIpv4UnicastRoutesReceived != nil
}

// Number of IPv4 Unicast prefixes received within BMP Monitor messages with L flag set (indicating application of in-bound policy on the advertised prefix)  from the BMP client from the time the BMP server was started.
// SetPostPolicyIpv4UnicastRoutesReceived sets the uint64 value in the BmpServerMetric object
func (obj *bmpServerMetric) SetPostPolicyIpv4UnicastRoutesReceived(value uint64) BmpServerMetric {

	obj.obj.PostPolicyIpv4UnicastRoutesReceived = &value
	return obj
}

// Number of IPv6 Unicast prefixes received within BMP Monitor messages with L flag unset (indicating absence of application on in-bound policy on the advertised prefix)  from the BMP client from the time the BMP server was started.
// PrePolicyIpv6UnicastRoutesReceived returns a uint64
func (obj *bmpServerMetric) PrePolicyIpv6UnicastRoutesReceived() uint64 {

	return *obj.obj.PrePolicyIpv6UnicastRoutesReceived

}

// Number of IPv6 Unicast prefixes received within BMP Monitor messages with L flag unset (indicating absence of application on in-bound policy on the advertised prefix)  from the BMP client from the time the BMP server was started.
// PrePolicyIpv6UnicastRoutesReceived returns a uint64
func (obj *bmpServerMetric) HasPrePolicyIpv6UnicastRoutesReceived() bool {
	return obj.obj.PrePolicyIpv6UnicastRoutesReceived != nil
}

// Number of IPv6 Unicast prefixes received within BMP Monitor messages with L flag unset (indicating absence of application on in-bound policy on the advertised prefix)  from the BMP client from the time the BMP server was started.
// SetPrePolicyIpv6UnicastRoutesReceived sets the uint64 value in the BmpServerMetric object
func (obj *bmpServerMetric) SetPrePolicyIpv6UnicastRoutesReceived(value uint64) BmpServerMetric {

	obj.obj.PrePolicyIpv6UnicastRoutesReceived = &value
	return obj
}

// Number of IPv6 Unicast prefixes received within BMP Monitor messages with L flag set (indicating application of in-bound policy on the advertised prefix)  from the BMP client from the time the BMP server was started.
// PostPolicyIpv6UnicastRoutesReceived returns a uint64
func (obj *bmpServerMetric) PostPolicyIpv6UnicastRoutesReceived() uint64 {

	return *obj.obj.PostPolicyIpv6UnicastRoutesReceived

}

// Number of IPv6 Unicast prefixes received within BMP Monitor messages with L flag set (indicating application of in-bound policy on the advertised prefix)  from the BMP client from the time the BMP server was started.
// PostPolicyIpv6UnicastRoutesReceived returns a uint64
func (obj *bmpServerMetric) HasPostPolicyIpv6UnicastRoutesReceived() bool {
	return obj.obj.PostPolicyIpv6UnicastRoutesReceived != nil
}

// Number of IPv6 Unicast prefixes received within BMP Monitor messages with L flag set (indicating application of in-bound policy on the advertised prefix)  from the BMP client from the time the BMP server was started.
// SetPostPolicyIpv6UnicastRoutesReceived sets the uint64 value in the BmpServerMetric object
func (obj *bmpServerMetric) SetPostPolicyIpv6UnicastRoutesReceived(value uint64) BmpServerMetric {

	obj.obj.PostPolicyIpv6UnicastRoutesReceived = &value
	return obj
}

func (obj *bmpServerMetric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *bmpServerMetric) setDefault() {

}
