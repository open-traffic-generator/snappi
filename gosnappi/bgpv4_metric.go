package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Bgpv4Metric *****
type bgpv4Metric struct {
	validation
	obj          *otg.Bgpv4Metric
	marshaller   marshalBgpv4Metric
	unMarshaller unMarshalBgpv4Metric
}

func NewBgpv4Metric() Bgpv4Metric {
	obj := bgpv4Metric{obj: &otg.Bgpv4Metric{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpv4Metric) msg() *otg.Bgpv4Metric {
	return obj.obj
}

func (obj *bgpv4Metric) setMsg(msg *otg.Bgpv4Metric) Bgpv4Metric {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpv4Metric struct {
	obj *bgpv4Metric
}

type marshalBgpv4Metric interface {
	// ToProto marshals Bgpv4Metric to protobuf object *otg.Bgpv4Metric
	ToProto() (*otg.Bgpv4Metric, error)
	// ToPbText marshals Bgpv4Metric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Bgpv4Metric to YAML text
	ToYaml() (string, error)
	// ToJson marshals Bgpv4Metric to JSON text
	ToJson() (string, error)
}

type unMarshalbgpv4Metric struct {
	obj *bgpv4Metric
}

type unMarshalBgpv4Metric interface {
	// FromProto unmarshals Bgpv4Metric from protobuf object *otg.Bgpv4Metric
	FromProto(msg *otg.Bgpv4Metric) (Bgpv4Metric, error)
	// FromPbText unmarshals Bgpv4Metric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Bgpv4Metric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Bgpv4Metric from JSON text
	FromJson(value string) error
}

func (obj *bgpv4Metric) Marshal() marshalBgpv4Metric {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpv4Metric{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpv4Metric) Unmarshal() unMarshalBgpv4Metric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpv4Metric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpv4Metric) ToProto() (*otg.Bgpv4Metric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpv4Metric) FromProto(msg *otg.Bgpv4Metric) (Bgpv4Metric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpv4Metric) ToPbText() (string, error) {
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

func (m *unMarshalbgpv4Metric) FromPbText(value string) error {
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

func (m *marshalbgpv4Metric) ToYaml() (string, error) {
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

func (m *unMarshalbgpv4Metric) FromYaml(value string) error {
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

func (m *marshalbgpv4Metric) ToJson() (string, error) {
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

func (m *unMarshalbgpv4Metric) FromJson(value string) error {
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

func (obj *bgpv4Metric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpv4Metric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpv4Metric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpv4Metric) Clone() (Bgpv4Metric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpv4Metric()
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

// Bgpv4Metric is bGPv4 per peer statistics information.
type Bgpv4Metric interface {
	Validation
	// msg marshals Bgpv4Metric to protobuf object *otg.Bgpv4Metric
	// and doesn't set defaults
	msg() *otg.Bgpv4Metric
	// setMsg unmarshals Bgpv4Metric from protobuf object *otg.Bgpv4Metric
	// and doesn't set defaults
	setMsg(*otg.Bgpv4Metric) Bgpv4Metric
	// provides marshal interface
	Marshal() marshalBgpv4Metric
	// provides unmarshal interface
	Unmarshal() unMarshalBgpv4Metric
	// validate validates Bgpv4Metric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Bgpv4Metric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in Bgpv4Metric.
	Name() string
	// SetName assigns string provided by user to Bgpv4Metric
	SetName(value string) Bgpv4Metric
	// HasName checks if Name has been set in Bgpv4Metric
	HasName() bool
	// SessionState returns Bgpv4MetricSessionStateEnum, set in Bgpv4Metric
	SessionState() Bgpv4MetricSessionStateEnum
	// SetSessionState assigns Bgpv4MetricSessionStateEnum provided by user to Bgpv4Metric
	SetSessionState(value Bgpv4MetricSessionStateEnum) Bgpv4Metric
	// HasSessionState checks if SessionState has been set in Bgpv4Metric
	HasSessionState() bool
	// SessionFlapCount returns uint64, set in Bgpv4Metric.
	SessionFlapCount() uint64
	// SetSessionFlapCount assigns uint64 provided by user to Bgpv4Metric
	SetSessionFlapCount(value uint64) Bgpv4Metric
	// HasSessionFlapCount checks if SessionFlapCount has been set in Bgpv4Metric
	HasSessionFlapCount() bool
	// RoutesAdvertised returns uint64, set in Bgpv4Metric.
	RoutesAdvertised() uint64
	// SetRoutesAdvertised assigns uint64 provided by user to Bgpv4Metric
	SetRoutesAdvertised(value uint64) Bgpv4Metric
	// HasRoutesAdvertised checks if RoutesAdvertised has been set in Bgpv4Metric
	HasRoutesAdvertised() bool
	// RoutesReceived returns uint64, set in Bgpv4Metric.
	RoutesReceived() uint64
	// SetRoutesReceived assigns uint64 provided by user to Bgpv4Metric
	SetRoutesReceived(value uint64) Bgpv4Metric
	// HasRoutesReceived checks if RoutesReceived has been set in Bgpv4Metric
	HasRoutesReceived() bool
	// RouteWithdrawsSent returns uint64, set in Bgpv4Metric.
	RouteWithdrawsSent() uint64
	// SetRouteWithdrawsSent assigns uint64 provided by user to Bgpv4Metric
	SetRouteWithdrawsSent(value uint64) Bgpv4Metric
	// HasRouteWithdrawsSent checks if RouteWithdrawsSent has been set in Bgpv4Metric
	HasRouteWithdrawsSent() bool
	// RouteWithdrawsReceived returns uint64, set in Bgpv4Metric.
	RouteWithdrawsReceived() uint64
	// SetRouteWithdrawsReceived assigns uint64 provided by user to Bgpv4Metric
	SetRouteWithdrawsReceived(value uint64) Bgpv4Metric
	// HasRouteWithdrawsReceived checks if RouteWithdrawsReceived has been set in Bgpv4Metric
	HasRouteWithdrawsReceived() bool
	// UpdatesSent returns uint64, set in Bgpv4Metric.
	UpdatesSent() uint64
	// SetUpdatesSent assigns uint64 provided by user to Bgpv4Metric
	SetUpdatesSent(value uint64) Bgpv4Metric
	// HasUpdatesSent checks if UpdatesSent has been set in Bgpv4Metric
	HasUpdatesSent() bool
	// UpdatesReceived returns uint64, set in Bgpv4Metric.
	UpdatesReceived() uint64
	// SetUpdatesReceived assigns uint64 provided by user to Bgpv4Metric
	SetUpdatesReceived(value uint64) Bgpv4Metric
	// HasUpdatesReceived checks if UpdatesReceived has been set in Bgpv4Metric
	HasUpdatesReceived() bool
	// OpensSent returns uint64, set in Bgpv4Metric.
	OpensSent() uint64
	// SetOpensSent assigns uint64 provided by user to Bgpv4Metric
	SetOpensSent(value uint64) Bgpv4Metric
	// HasOpensSent checks if OpensSent has been set in Bgpv4Metric
	HasOpensSent() bool
	// OpensReceived returns uint64, set in Bgpv4Metric.
	OpensReceived() uint64
	// SetOpensReceived assigns uint64 provided by user to Bgpv4Metric
	SetOpensReceived(value uint64) Bgpv4Metric
	// HasOpensReceived checks if OpensReceived has been set in Bgpv4Metric
	HasOpensReceived() bool
	// KeepalivesSent returns uint64, set in Bgpv4Metric.
	KeepalivesSent() uint64
	// SetKeepalivesSent assigns uint64 provided by user to Bgpv4Metric
	SetKeepalivesSent(value uint64) Bgpv4Metric
	// HasKeepalivesSent checks if KeepalivesSent has been set in Bgpv4Metric
	HasKeepalivesSent() bool
	// KeepalivesReceived returns uint64, set in Bgpv4Metric.
	KeepalivesReceived() uint64
	// SetKeepalivesReceived assigns uint64 provided by user to Bgpv4Metric
	SetKeepalivesReceived(value uint64) Bgpv4Metric
	// HasKeepalivesReceived checks if KeepalivesReceived has been set in Bgpv4Metric
	HasKeepalivesReceived() bool
	// NotificationsSent returns uint64, set in Bgpv4Metric.
	NotificationsSent() uint64
	// SetNotificationsSent assigns uint64 provided by user to Bgpv4Metric
	SetNotificationsSent(value uint64) Bgpv4Metric
	// HasNotificationsSent checks if NotificationsSent has been set in Bgpv4Metric
	HasNotificationsSent() bool
	// NotificationsReceived returns uint64, set in Bgpv4Metric.
	NotificationsReceived() uint64
	// SetNotificationsReceived assigns uint64 provided by user to Bgpv4Metric
	SetNotificationsReceived(value uint64) Bgpv4Metric
	// HasNotificationsReceived checks if NotificationsReceived has been set in Bgpv4Metric
	HasNotificationsReceived() bool
	// FsmState returns Bgpv4MetricFsmStateEnum, set in Bgpv4Metric
	FsmState() Bgpv4MetricFsmStateEnum
	// SetFsmState assigns Bgpv4MetricFsmStateEnum provided by user to Bgpv4Metric
	SetFsmState(value Bgpv4MetricFsmStateEnum) Bgpv4Metric
	// HasFsmState checks if FsmState has been set in Bgpv4Metric
	HasFsmState() bool
	// EndOfRibReceived returns uint64, set in Bgpv4Metric.
	EndOfRibReceived() uint64
	// SetEndOfRibReceived assigns uint64 provided by user to Bgpv4Metric
	SetEndOfRibReceived(value uint64) Bgpv4Metric
	// HasEndOfRibReceived checks if EndOfRibReceived has been set in Bgpv4Metric
	HasEndOfRibReceived() bool
}

// The name of a configured BGPv4 peer.
// Name returns a string
func (obj *bgpv4Metric) Name() string {

	return *obj.obj.Name

}

// The name of a configured BGPv4 peer.
// Name returns a string
func (obj *bgpv4Metric) HasName() bool {
	return obj.obj.Name != nil
}

// The name of a configured BGPv4 peer.
// SetName sets the string value in the Bgpv4Metric object
func (obj *bgpv4Metric) SetName(value string) Bgpv4Metric {

	obj.obj.Name = &value
	return obj
}

type Bgpv4MetricSessionStateEnum string

// Enum of SessionState on Bgpv4Metric
var Bgpv4MetricSessionState = struct {
	UP   Bgpv4MetricSessionStateEnum
	DOWN Bgpv4MetricSessionStateEnum
}{
	UP:   Bgpv4MetricSessionStateEnum("up"),
	DOWN: Bgpv4MetricSessionStateEnum("down"),
}

func (obj *bgpv4Metric) SessionState() Bgpv4MetricSessionStateEnum {
	return Bgpv4MetricSessionStateEnum(obj.obj.SessionState.Enum().String())
}

// Session state as up or down. Up refers to an Established state and Down refers to any other state.
// SessionState returns a string
func (obj *bgpv4Metric) HasSessionState() bool {
	return obj.obj.SessionState != nil
}

func (obj *bgpv4Metric) SetSessionState(value Bgpv4MetricSessionStateEnum) Bgpv4Metric {
	intValue, ok := otg.Bgpv4Metric_SessionState_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Bgpv4MetricSessionStateEnum", string(value)))
		return obj
	}
	enumValue := otg.Bgpv4Metric_SessionState_Enum(intValue)
	obj.obj.SessionState = &enumValue

	return obj
}

// Number of times the session went from Up to Down state.
// SessionFlapCount returns a uint64
func (obj *bgpv4Metric) SessionFlapCount() uint64 {

	return *obj.obj.SessionFlapCount

}

// Number of times the session went from Up to Down state.
// SessionFlapCount returns a uint64
func (obj *bgpv4Metric) HasSessionFlapCount() bool {
	return obj.obj.SessionFlapCount != nil
}

// Number of times the session went from Up to Down state.
// SetSessionFlapCount sets the uint64 value in the Bgpv4Metric object
func (obj *bgpv4Metric) SetSessionFlapCount(value uint64) Bgpv4Metric {

	obj.obj.SessionFlapCount = &value
	return obj
}

// Number of routes advertised.
// RoutesAdvertised returns a uint64
func (obj *bgpv4Metric) RoutesAdvertised() uint64 {

	return *obj.obj.RoutesAdvertised

}

// Number of routes advertised.
// RoutesAdvertised returns a uint64
func (obj *bgpv4Metric) HasRoutesAdvertised() bool {
	return obj.obj.RoutesAdvertised != nil
}

// Number of routes advertised.
// SetRoutesAdvertised sets the uint64 value in the Bgpv4Metric object
func (obj *bgpv4Metric) SetRoutesAdvertised(value uint64) Bgpv4Metric {

	obj.obj.RoutesAdvertised = &value
	return obj
}

// Number of routes received.
// RoutesReceived returns a uint64
func (obj *bgpv4Metric) RoutesReceived() uint64 {

	return *obj.obj.RoutesReceived

}

// Number of routes received.
// RoutesReceived returns a uint64
func (obj *bgpv4Metric) HasRoutesReceived() bool {
	return obj.obj.RoutesReceived != nil
}

// Number of routes received.
// SetRoutesReceived sets the uint64 value in the Bgpv4Metric object
func (obj *bgpv4Metric) SetRoutesReceived(value uint64) Bgpv4Metric {

	obj.obj.RoutesReceived = &value
	return obj
}

// Number of route withdraws sent.
// RouteWithdrawsSent returns a uint64
func (obj *bgpv4Metric) RouteWithdrawsSent() uint64 {

	return *obj.obj.RouteWithdrawsSent

}

// Number of route withdraws sent.
// RouteWithdrawsSent returns a uint64
func (obj *bgpv4Metric) HasRouteWithdrawsSent() bool {
	return obj.obj.RouteWithdrawsSent != nil
}

// Number of route withdraws sent.
// SetRouteWithdrawsSent sets the uint64 value in the Bgpv4Metric object
func (obj *bgpv4Metric) SetRouteWithdrawsSent(value uint64) Bgpv4Metric {

	obj.obj.RouteWithdrawsSent = &value
	return obj
}

// Number of route withdraws received.
// RouteWithdrawsReceived returns a uint64
func (obj *bgpv4Metric) RouteWithdrawsReceived() uint64 {

	return *obj.obj.RouteWithdrawsReceived

}

// Number of route withdraws received.
// RouteWithdrawsReceived returns a uint64
func (obj *bgpv4Metric) HasRouteWithdrawsReceived() bool {
	return obj.obj.RouteWithdrawsReceived != nil
}

// Number of route withdraws received.
// SetRouteWithdrawsReceived sets the uint64 value in the Bgpv4Metric object
func (obj *bgpv4Metric) SetRouteWithdrawsReceived(value uint64) Bgpv4Metric {

	obj.obj.RouteWithdrawsReceived = &value
	return obj
}

// Number of Update messages sent.
// UpdatesSent returns a uint64
func (obj *bgpv4Metric) UpdatesSent() uint64 {

	return *obj.obj.UpdatesSent

}

// Number of Update messages sent.
// UpdatesSent returns a uint64
func (obj *bgpv4Metric) HasUpdatesSent() bool {
	return obj.obj.UpdatesSent != nil
}

// Number of Update messages sent.
// SetUpdatesSent sets the uint64 value in the Bgpv4Metric object
func (obj *bgpv4Metric) SetUpdatesSent(value uint64) Bgpv4Metric {

	obj.obj.UpdatesSent = &value
	return obj
}

// Number of Update messages received.
// UpdatesReceived returns a uint64
func (obj *bgpv4Metric) UpdatesReceived() uint64 {

	return *obj.obj.UpdatesReceived

}

// Number of Update messages received.
// UpdatesReceived returns a uint64
func (obj *bgpv4Metric) HasUpdatesReceived() bool {
	return obj.obj.UpdatesReceived != nil
}

// Number of Update messages received.
// SetUpdatesReceived sets the uint64 value in the Bgpv4Metric object
func (obj *bgpv4Metric) SetUpdatesReceived(value uint64) Bgpv4Metric {

	obj.obj.UpdatesReceived = &value
	return obj
}

// Number of Open messages sent.
// OpensSent returns a uint64
func (obj *bgpv4Metric) OpensSent() uint64 {

	return *obj.obj.OpensSent

}

// Number of Open messages sent.
// OpensSent returns a uint64
func (obj *bgpv4Metric) HasOpensSent() bool {
	return obj.obj.OpensSent != nil
}

// Number of Open messages sent.
// SetOpensSent sets the uint64 value in the Bgpv4Metric object
func (obj *bgpv4Metric) SetOpensSent(value uint64) Bgpv4Metric {

	obj.obj.OpensSent = &value
	return obj
}

// Number of Open messages received.
// OpensReceived returns a uint64
func (obj *bgpv4Metric) OpensReceived() uint64 {

	return *obj.obj.OpensReceived

}

// Number of Open messages received.
// OpensReceived returns a uint64
func (obj *bgpv4Metric) HasOpensReceived() bool {
	return obj.obj.OpensReceived != nil
}

// Number of Open messages received.
// SetOpensReceived sets the uint64 value in the Bgpv4Metric object
func (obj *bgpv4Metric) SetOpensReceived(value uint64) Bgpv4Metric {

	obj.obj.OpensReceived = &value
	return obj
}

// Number of Keepalive messages sent.
// KeepalivesSent returns a uint64
func (obj *bgpv4Metric) KeepalivesSent() uint64 {

	return *obj.obj.KeepalivesSent

}

// Number of Keepalive messages sent.
// KeepalivesSent returns a uint64
func (obj *bgpv4Metric) HasKeepalivesSent() bool {
	return obj.obj.KeepalivesSent != nil
}

// Number of Keepalive messages sent.
// SetKeepalivesSent sets the uint64 value in the Bgpv4Metric object
func (obj *bgpv4Metric) SetKeepalivesSent(value uint64) Bgpv4Metric {

	obj.obj.KeepalivesSent = &value
	return obj
}

// Number of Keepalive messages received.
// KeepalivesReceived returns a uint64
func (obj *bgpv4Metric) KeepalivesReceived() uint64 {

	return *obj.obj.KeepalivesReceived

}

// Number of Keepalive messages received.
// KeepalivesReceived returns a uint64
func (obj *bgpv4Metric) HasKeepalivesReceived() bool {
	return obj.obj.KeepalivesReceived != nil
}

// Number of Keepalive messages received.
// SetKeepalivesReceived sets the uint64 value in the Bgpv4Metric object
func (obj *bgpv4Metric) SetKeepalivesReceived(value uint64) Bgpv4Metric {

	obj.obj.KeepalivesReceived = &value
	return obj
}

// Number of Notification messages sent.
// NotificationsSent returns a uint64
func (obj *bgpv4Metric) NotificationsSent() uint64 {

	return *obj.obj.NotificationsSent

}

// Number of Notification messages sent.
// NotificationsSent returns a uint64
func (obj *bgpv4Metric) HasNotificationsSent() bool {
	return obj.obj.NotificationsSent != nil
}

// Number of Notification messages sent.
// SetNotificationsSent sets the uint64 value in the Bgpv4Metric object
func (obj *bgpv4Metric) SetNotificationsSent(value uint64) Bgpv4Metric {

	obj.obj.NotificationsSent = &value
	return obj
}

// Number of Notification messages received.
// NotificationsReceived returns a uint64
func (obj *bgpv4Metric) NotificationsReceived() uint64 {

	return *obj.obj.NotificationsReceived

}

// Number of Notification messages received.
// NotificationsReceived returns a uint64
func (obj *bgpv4Metric) HasNotificationsReceived() bool {
	return obj.obj.NotificationsReceived != nil
}

// Number of Notification messages received.
// SetNotificationsReceived sets the uint64 value in the Bgpv4Metric object
func (obj *bgpv4Metric) SetNotificationsReceived(value uint64) Bgpv4Metric {

	obj.obj.NotificationsReceived = &value
	return obj
}

type Bgpv4MetricFsmStateEnum string

// Enum of FsmState on Bgpv4Metric
var Bgpv4MetricFsmState = struct {
	IDLE        Bgpv4MetricFsmStateEnum
	CONNECT     Bgpv4MetricFsmStateEnum
	ACTIVE      Bgpv4MetricFsmStateEnum
	OPENSENT    Bgpv4MetricFsmStateEnum
	OPENCONFIRM Bgpv4MetricFsmStateEnum
	ESTABLISHED Bgpv4MetricFsmStateEnum
}{
	IDLE:        Bgpv4MetricFsmStateEnum("idle"),
	CONNECT:     Bgpv4MetricFsmStateEnum("connect"),
	ACTIVE:      Bgpv4MetricFsmStateEnum("active"),
	OPENSENT:    Bgpv4MetricFsmStateEnum("opensent"),
	OPENCONFIRM: Bgpv4MetricFsmStateEnum("openconfirm"),
	ESTABLISHED: Bgpv4MetricFsmStateEnum("established"),
}

func (obj *bgpv4Metric) FsmState() Bgpv4MetricFsmStateEnum {
	return Bgpv4MetricFsmStateEnum(obj.obj.FsmState.Enum().String())
}

// BGP peer FSM (Finite State Machine) state as Idle, Connect, Active, OpenSent, OpenConfirm and Established. In all the states except Established the BGP session is down. Idle refers to the Idle state of the FSM. Connect refers to the state where the session is waiting for the underlying transport session to be established. Active refers to the state where the session is awaiting for a connection from the remote peer. OpenSent refers to the state where the session is in the process of being established. The local system has sent an OPEN message. OpenConfirm refers to the state where the session is in the process of being established. The local system has sent and received an OPEN message and is awaiting a NOTIFICATION or KEEPALIVE message from remote peer. Established refers to the state where the BGP session with the peer is established.
// FsmState returns a string
func (obj *bgpv4Metric) HasFsmState() bool {
	return obj.obj.FsmState != nil
}

func (obj *bgpv4Metric) SetFsmState(value Bgpv4MetricFsmStateEnum) Bgpv4Metric {
	intValue, ok := otg.Bgpv4Metric_FsmState_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Bgpv4MetricFsmStateEnum", string(value)))
		return obj
	}
	enumValue := otg.Bgpv4Metric_FsmState_Enum(intValue)
	obj.obj.FsmState = &enumValue

	return obj
}

// Number of End-of-RIB markers received indicating the completion of the initial routing update for a  particular <AFI, SAFI> address family after the session is established. For the IPv4 unicast address family, the End-of-RIB marker is an UPDATE message with the minimum length. For any other address family, it is an UPDATE message that contains only the MP_UNREACH_NLRI attribute with  no withdrawn routes for that <AFI, SAFI>.
// EndOfRibReceived returns a uint64
func (obj *bgpv4Metric) EndOfRibReceived() uint64 {

	return *obj.obj.EndOfRibReceived

}

// Number of End-of-RIB markers received indicating the completion of the initial routing update for a  particular <AFI, SAFI> address family after the session is established. For the IPv4 unicast address family, the End-of-RIB marker is an UPDATE message with the minimum length. For any other address family, it is an UPDATE message that contains only the MP_UNREACH_NLRI attribute with  no withdrawn routes for that <AFI, SAFI>.
// EndOfRibReceived returns a uint64
func (obj *bgpv4Metric) HasEndOfRibReceived() bool {
	return obj.obj.EndOfRibReceived != nil
}

// Number of End-of-RIB markers received indicating the completion of the initial routing update for a  particular <AFI, SAFI> address family after the session is established. For the IPv4 unicast address family, the End-of-RIB marker is an UPDATE message with the minimum length. For any other address family, it is an UPDATE message that contains only the MP_UNREACH_NLRI attribute with  no withdrawn routes for that <AFI, SAFI>.
// SetEndOfRibReceived sets the uint64 value in the Bgpv4Metric object
func (obj *bgpv4Metric) SetEndOfRibReceived(value uint64) Bgpv4Metric {

	obj.obj.EndOfRibReceived = &value
	return obj
}

func (obj *bgpv4Metric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *bgpv4Metric) setDefault() {

}
