package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Bgpv6Metric *****
type bgpv6Metric struct {
	validation
	obj          *otg.Bgpv6Metric
	marshaller   marshalBgpv6Metric
	unMarshaller unMarshalBgpv6Metric
}

func NewBgpv6Metric() Bgpv6Metric {
	obj := bgpv6Metric{obj: &otg.Bgpv6Metric{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpv6Metric) msg() *otg.Bgpv6Metric {
	return obj.obj
}

func (obj *bgpv6Metric) setMsg(msg *otg.Bgpv6Metric) Bgpv6Metric {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpv6Metric struct {
	obj *bgpv6Metric
}

type marshalBgpv6Metric interface {
	// ToProto marshals Bgpv6Metric to protobuf object *otg.Bgpv6Metric
	ToProto() (*otg.Bgpv6Metric, error)
	// ToPbText marshals Bgpv6Metric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Bgpv6Metric to YAML text
	ToYaml() (string, error)
	// ToJson marshals Bgpv6Metric to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Bgpv6Metric to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpv6Metric struct {
	obj *bgpv6Metric
}

type unMarshalBgpv6Metric interface {
	// FromProto unmarshals Bgpv6Metric from protobuf object *otg.Bgpv6Metric
	FromProto(msg *otg.Bgpv6Metric) (Bgpv6Metric, error)
	// FromPbText unmarshals Bgpv6Metric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Bgpv6Metric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Bgpv6Metric from JSON text
	FromJson(value string) error
}

func (obj *bgpv6Metric) Marshal() marshalBgpv6Metric {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpv6Metric{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpv6Metric) Unmarshal() unMarshalBgpv6Metric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpv6Metric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpv6Metric) ToProto() (*otg.Bgpv6Metric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpv6Metric) FromProto(msg *otg.Bgpv6Metric) (Bgpv6Metric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpv6Metric) ToPbText() (string, error) {
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

func (m *unMarshalbgpv6Metric) FromPbText(value string) error {
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

func (m *marshalbgpv6Metric) ToYaml() (string, error) {
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

func (m *unMarshalbgpv6Metric) FromYaml(value string) error {
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

func (m *marshalbgpv6Metric) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalbgpv6Metric) ToJson() (string, error) {
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

func (m *unMarshalbgpv6Metric) FromJson(value string) error {
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

func (obj *bgpv6Metric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpv6Metric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpv6Metric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpv6Metric) Clone() (Bgpv6Metric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpv6Metric()
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

// Bgpv6Metric is bGPv6 per peer statistics information.
type Bgpv6Metric interface {
	Validation
	// msg marshals Bgpv6Metric to protobuf object *otg.Bgpv6Metric
	// and doesn't set defaults
	msg() *otg.Bgpv6Metric
	// setMsg unmarshals Bgpv6Metric from protobuf object *otg.Bgpv6Metric
	// and doesn't set defaults
	setMsg(*otg.Bgpv6Metric) Bgpv6Metric
	// provides marshal interface
	Marshal() marshalBgpv6Metric
	// provides unmarshal interface
	Unmarshal() unMarshalBgpv6Metric
	// validate validates Bgpv6Metric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Bgpv6Metric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in Bgpv6Metric.
	Name() string
	// SetName assigns string provided by user to Bgpv6Metric
	SetName(value string) Bgpv6Metric
	// HasName checks if Name has been set in Bgpv6Metric
	HasName() bool
	// SessionState returns Bgpv6MetricSessionStateEnum, set in Bgpv6Metric
	SessionState() Bgpv6MetricSessionStateEnum
	// SetSessionState assigns Bgpv6MetricSessionStateEnum provided by user to Bgpv6Metric
	SetSessionState(value Bgpv6MetricSessionStateEnum) Bgpv6Metric
	// HasSessionState checks if SessionState has been set in Bgpv6Metric
	HasSessionState() bool
	// SessionFlapCount returns uint64, set in Bgpv6Metric.
	SessionFlapCount() uint64
	// SetSessionFlapCount assigns uint64 provided by user to Bgpv6Metric
	SetSessionFlapCount(value uint64) Bgpv6Metric
	// HasSessionFlapCount checks if SessionFlapCount has been set in Bgpv6Metric
	HasSessionFlapCount() bool
	// RoutesAdvertised returns uint64, set in Bgpv6Metric.
	RoutesAdvertised() uint64
	// SetRoutesAdvertised assigns uint64 provided by user to Bgpv6Metric
	SetRoutesAdvertised(value uint64) Bgpv6Metric
	// HasRoutesAdvertised checks if RoutesAdvertised has been set in Bgpv6Metric
	HasRoutesAdvertised() bool
	// RoutesReceived returns uint64, set in Bgpv6Metric.
	RoutesReceived() uint64
	// SetRoutesReceived assigns uint64 provided by user to Bgpv6Metric
	SetRoutesReceived(value uint64) Bgpv6Metric
	// HasRoutesReceived checks if RoutesReceived has been set in Bgpv6Metric
	HasRoutesReceived() bool
	// RouteWithdrawsSent returns uint64, set in Bgpv6Metric.
	RouteWithdrawsSent() uint64
	// SetRouteWithdrawsSent assigns uint64 provided by user to Bgpv6Metric
	SetRouteWithdrawsSent(value uint64) Bgpv6Metric
	// HasRouteWithdrawsSent checks if RouteWithdrawsSent has been set in Bgpv6Metric
	HasRouteWithdrawsSent() bool
	// RouteWithdrawsReceived returns uint64, set in Bgpv6Metric.
	RouteWithdrawsReceived() uint64
	// SetRouteWithdrawsReceived assigns uint64 provided by user to Bgpv6Metric
	SetRouteWithdrawsReceived(value uint64) Bgpv6Metric
	// HasRouteWithdrawsReceived checks if RouteWithdrawsReceived has been set in Bgpv6Metric
	HasRouteWithdrawsReceived() bool
	// UpdatesSent returns uint64, set in Bgpv6Metric.
	UpdatesSent() uint64
	// SetUpdatesSent assigns uint64 provided by user to Bgpv6Metric
	SetUpdatesSent(value uint64) Bgpv6Metric
	// HasUpdatesSent checks if UpdatesSent has been set in Bgpv6Metric
	HasUpdatesSent() bool
	// UpdatesReceived returns uint64, set in Bgpv6Metric.
	UpdatesReceived() uint64
	// SetUpdatesReceived assigns uint64 provided by user to Bgpv6Metric
	SetUpdatesReceived(value uint64) Bgpv6Metric
	// HasUpdatesReceived checks if UpdatesReceived has been set in Bgpv6Metric
	HasUpdatesReceived() bool
	// OpensSent returns uint64, set in Bgpv6Metric.
	OpensSent() uint64
	// SetOpensSent assigns uint64 provided by user to Bgpv6Metric
	SetOpensSent(value uint64) Bgpv6Metric
	// HasOpensSent checks if OpensSent has been set in Bgpv6Metric
	HasOpensSent() bool
	// OpensReceived returns uint64, set in Bgpv6Metric.
	OpensReceived() uint64
	// SetOpensReceived assigns uint64 provided by user to Bgpv6Metric
	SetOpensReceived(value uint64) Bgpv6Metric
	// HasOpensReceived checks if OpensReceived has been set in Bgpv6Metric
	HasOpensReceived() bool
	// KeepalivesSent returns uint64, set in Bgpv6Metric.
	KeepalivesSent() uint64
	// SetKeepalivesSent assigns uint64 provided by user to Bgpv6Metric
	SetKeepalivesSent(value uint64) Bgpv6Metric
	// HasKeepalivesSent checks if KeepalivesSent has been set in Bgpv6Metric
	HasKeepalivesSent() bool
	// KeepalivesReceived returns uint64, set in Bgpv6Metric.
	KeepalivesReceived() uint64
	// SetKeepalivesReceived assigns uint64 provided by user to Bgpv6Metric
	SetKeepalivesReceived(value uint64) Bgpv6Metric
	// HasKeepalivesReceived checks if KeepalivesReceived has been set in Bgpv6Metric
	HasKeepalivesReceived() bool
	// NotificationsSent returns uint64, set in Bgpv6Metric.
	NotificationsSent() uint64
	// SetNotificationsSent assigns uint64 provided by user to Bgpv6Metric
	SetNotificationsSent(value uint64) Bgpv6Metric
	// HasNotificationsSent checks if NotificationsSent has been set in Bgpv6Metric
	HasNotificationsSent() bool
	// NotificationsReceived returns uint64, set in Bgpv6Metric.
	NotificationsReceived() uint64
	// SetNotificationsReceived assigns uint64 provided by user to Bgpv6Metric
	SetNotificationsReceived(value uint64) Bgpv6Metric
	// HasNotificationsReceived checks if NotificationsReceived has been set in Bgpv6Metric
	HasNotificationsReceived() bool
	// FsmState returns Bgpv6MetricFsmStateEnum, set in Bgpv6Metric
	FsmState() Bgpv6MetricFsmStateEnum
	// SetFsmState assigns Bgpv6MetricFsmStateEnum provided by user to Bgpv6Metric
	SetFsmState(value Bgpv6MetricFsmStateEnum) Bgpv6Metric
	// HasFsmState checks if FsmState has been set in Bgpv6Metric
	HasFsmState() bool
	// EndOfRibReceived returns uint64, set in Bgpv6Metric.
	EndOfRibReceived() uint64
	// SetEndOfRibReceived assigns uint64 provided by user to Bgpv6Metric
	SetEndOfRibReceived(value uint64) Bgpv6Metric
	// HasEndOfRibReceived checks if EndOfRibReceived has been set in Bgpv6Metric
	HasEndOfRibReceived() bool
}

// The name of a configured BGPv6 peer.
// Name returns a string
func (obj *bgpv6Metric) Name() string {

	return *obj.obj.Name

}

// The name of a configured BGPv6 peer.
// Name returns a string
func (obj *bgpv6Metric) HasName() bool {
	return obj.obj.Name != nil
}

// The name of a configured BGPv6 peer.
// SetName sets the string value in the Bgpv6Metric object
func (obj *bgpv6Metric) SetName(value string) Bgpv6Metric {

	obj.obj.Name = &value
	return obj
}

type Bgpv6MetricSessionStateEnum string

// Enum of SessionState on Bgpv6Metric
var Bgpv6MetricSessionState = struct {
	UP   Bgpv6MetricSessionStateEnum
	DOWN Bgpv6MetricSessionStateEnum
}{
	UP:   Bgpv6MetricSessionStateEnum("up"),
	DOWN: Bgpv6MetricSessionStateEnum("down"),
}

func (obj *bgpv6Metric) SessionState() Bgpv6MetricSessionStateEnum {
	return Bgpv6MetricSessionStateEnum(obj.obj.SessionState.Enum().String())
}

// Session state as up or down. Up refers to an Established state and Down refers to any other state.
// SessionState returns a string
func (obj *bgpv6Metric) HasSessionState() bool {
	return obj.obj.SessionState != nil
}

func (obj *bgpv6Metric) SetSessionState(value Bgpv6MetricSessionStateEnum) Bgpv6Metric {
	intValue, ok := otg.Bgpv6Metric_SessionState_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Bgpv6MetricSessionStateEnum", string(value)))
		return obj
	}
	enumValue := otg.Bgpv6Metric_SessionState_Enum(intValue)
	obj.obj.SessionState = &enumValue

	return obj
}

// Number of times the session went from Up to Down state.
// SessionFlapCount returns a uint64
func (obj *bgpv6Metric) SessionFlapCount() uint64 {

	return *obj.obj.SessionFlapCount

}

// Number of times the session went from Up to Down state.
// SessionFlapCount returns a uint64
func (obj *bgpv6Metric) HasSessionFlapCount() bool {
	return obj.obj.SessionFlapCount != nil
}

// Number of times the session went from Up to Down state.
// SetSessionFlapCount sets the uint64 value in the Bgpv6Metric object
func (obj *bgpv6Metric) SetSessionFlapCount(value uint64) Bgpv6Metric {

	obj.obj.SessionFlapCount = &value
	return obj
}

// Number of routes advertised.
// RoutesAdvertised returns a uint64
func (obj *bgpv6Metric) RoutesAdvertised() uint64 {

	return *obj.obj.RoutesAdvertised

}

// Number of routes advertised.
// RoutesAdvertised returns a uint64
func (obj *bgpv6Metric) HasRoutesAdvertised() bool {
	return obj.obj.RoutesAdvertised != nil
}

// Number of routes advertised.
// SetRoutesAdvertised sets the uint64 value in the Bgpv6Metric object
func (obj *bgpv6Metric) SetRoutesAdvertised(value uint64) Bgpv6Metric {

	obj.obj.RoutesAdvertised = &value
	return obj
}

// Number of routes received.
// RoutesReceived returns a uint64
func (obj *bgpv6Metric) RoutesReceived() uint64 {

	return *obj.obj.RoutesReceived

}

// Number of routes received.
// RoutesReceived returns a uint64
func (obj *bgpv6Metric) HasRoutesReceived() bool {
	return obj.obj.RoutesReceived != nil
}

// Number of routes received.
// SetRoutesReceived sets the uint64 value in the Bgpv6Metric object
func (obj *bgpv6Metric) SetRoutesReceived(value uint64) Bgpv6Metric {

	obj.obj.RoutesReceived = &value
	return obj
}

// Number of route withdraws sent.
// RouteWithdrawsSent returns a uint64
func (obj *bgpv6Metric) RouteWithdrawsSent() uint64 {

	return *obj.obj.RouteWithdrawsSent

}

// Number of route withdraws sent.
// RouteWithdrawsSent returns a uint64
func (obj *bgpv6Metric) HasRouteWithdrawsSent() bool {
	return obj.obj.RouteWithdrawsSent != nil
}

// Number of route withdraws sent.
// SetRouteWithdrawsSent sets the uint64 value in the Bgpv6Metric object
func (obj *bgpv6Metric) SetRouteWithdrawsSent(value uint64) Bgpv6Metric {

	obj.obj.RouteWithdrawsSent = &value
	return obj
}

// Number of route withdraws received.
// RouteWithdrawsReceived returns a uint64
func (obj *bgpv6Metric) RouteWithdrawsReceived() uint64 {

	return *obj.obj.RouteWithdrawsReceived

}

// Number of route withdraws received.
// RouteWithdrawsReceived returns a uint64
func (obj *bgpv6Metric) HasRouteWithdrawsReceived() bool {
	return obj.obj.RouteWithdrawsReceived != nil
}

// Number of route withdraws received.
// SetRouteWithdrawsReceived sets the uint64 value in the Bgpv6Metric object
func (obj *bgpv6Metric) SetRouteWithdrawsReceived(value uint64) Bgpv6Metric {

	obj.obj.RouteWithdrawsReceived = &value
	return obj
}

// Number of Update messages sent.
// UpdatesSent returns a uint64
func (obj *bgpv6Metric) UpdatesSent() uint64 {

	return *obj.obj.UpdatesSent

}

// Number of Update messages sent.
// UpdatesSent returns a uint64
func (obj *bgpv6Metric) HasUpdatesSent() bool {
	return obj.obj.UpdatesSent != nil
}

// Number of Update messages sent.
// SetUpdatesSent sets the uint64 value in the Bgpv6Metric object
func (obj *bgpv6Metric) SetUpdatesSent(value uint64) Bgpv6Metric {

	obj.obj.UpdatesSent = &value
	return obj
}

// Number of Update messages received.
// UpdatesReceived returns a uint64
func (obj *bgpv6Metric) UpdatesReceived() uint64 {

	return *obj.obj.UpdatesReceived

}

// Number of Update messages received.
// UpdatesReceived returns a uint64
func (obj *bgpv6Metric) HasUpdatesReceived() bool {
	return obj.obj.UpdatesReceived != nil
}

// Number of Update messages received.
// SetUpdatesReceived sets the uint64 value in the Bgpv6Metric object
func (obj *bgpv6Metric) SetUpdatesReceived(value uint64) Bgpv6Metric {

	obj.obj.UpdatesReceived = &value
	return obj
}

// Number of Open messages sent.
// OpensSent returns a uint64
func (obj *bgpv6Metric) OpensSent() uint64 {

	return *obj.obj.OpensSent

}

// Number of Open messages sent.
// OpensSent returns a uint64
func (obj *bgpv6Metric) HasOpensSent() bool {
	return obj.obj.OpensSent != nil
}

// Number of Open messages sent.
// SetOpensSent sets the uint64 value in the Bgpv6Metric object
func (obj *bgpv6Metric) SetOpensSent(value uint64) Bgpv6Metric {

	obj.obj.OpensSent = &value
	return obj
}

// Number of Open messages received.
// OpensReceived returns a uint64
func (obj *bgpv6Metric) OpensReceived() uint64 {

	return *obj.obj.OpensReceived

}

// Number of Open messages received.
// OpensReceived returns a uint64
func (obj *bgpv6Metric) HasOpensReceived() bool {
	return obj.obj.OpensReceived != nil
}

// Number of Open messages received.
// SetOpensReceived sets the uint64 value in the Bgpv6Metric object
func (obj *bgpv6Metric) SetOpensReceived(value uint64) Bgpv6Metric {

	obj.obj.OpensReceived = &value
	return obj
}

// Number of Keepalive messages sent.
// KeepalivesSent returns a uint64
func (obj *bgpv6Metric) KeepalivesSent() uint64 {

	return *obj.obj.KeepalivesSent

}

// Number of Keepalive messages sent.
// KeepalivesSent returns a uint64
func (obj *bgpv6Metric) HasKeepalivesSent() bool {
	return obj.obj.KeepalivesSent != nil
}

// Number of Keepalive messages sent.
// SetKeepalivesSent sets the uint64 value in the Bgpv6Metric object
func (obj *bgpv6Metric) SetKeepalivesSent(value uint64) Bgpv6Metric {

	obj.obj.KeepalivesSent = &value
	return obj
}

// Number of Keepalive messages received.
// KeepalivesReceived returns a uint64
func (obj *bgpv6Metric) KeepalivesReceived() uint64 {

	return *obj.obj.KeepalivesReceived

}

// Number of Keepalive messages received.
// KeepalivesReceived returns a uint64
func (obj *bgpv6Metric) HasKeepalivesReceived() bool {
	return obj.obj.KeepalivesReceived != nil
}

// Number of Keepalive messages received.
// SetKeepalivesReceived sets the uint64 value in the Bgpv6Metric object
func (obj *bgpv6Metric) SetKeepalivesReceived(value uint64) Bgpv6Metric {

	obj.obj.KeepalivesReceived = &value
	return obj
}

// Number of Notification messages sent.
// NotificationsSent returns a uint64
func (obj *bgpv6Metric) NotificationsSent() uint64 {

	return *obj.obj.NotificationsSent

}

// Number of Notification messages sent.
// NotificationsSent returns a uint64
func (obj *bgpv6Metric) HasNotificationsSent() bool {
	return obj.obj.NotificationsSent != nil
}

// Number of Notification messages sent.
// SetNotificationsSent sets the uint64 value in the Bgpv6Metric object
func (obj *bgpv6Metric) SetNotificationsSent(value uint64) Bgpv6Metric {

	obj.obj.NotificationsSent = &value
	return obj
}

// Number of Notification messages received.
// NotificationsReceived returns a uint64
func (obj *bgpv6Metric) NotificationsReceived() uint64 {

	return *obj.obj.NotificationsReceived

}

// Number of Notification messages received.
// NotificationsReceived returns a uint64
func (obj *bgpv6Metric) HasNotificationsReceived() bool {
	return obj.obj.NotificationsReceived != nil
}

// Number of Notification messages received.
// SetNotificationsReceived sets the uint64 value in the Bgpv6Metric object
func (obj *bgpv6Metric) SetNotificationsReceived(value uint64) Bgpv6Metric {

	obj.obj.NotificationsReceived = &value
	return obj
}

type Bgpv6MetricFsmStateEnum string

// Enum of FsmState on Bgpv6Metric
var Bgpv6MetricFsmState = struct {
	IDLE        Bgpv6MetricFsmStateEnum
	CONNECT     Bgpv6MetricFsmStateEnum
	ACTIVE      Bgpv6MetricFsmStateEnum
	OPENSENT    Bgpv6MetricFsmStateEnum
	OPENCONFIRM Bgpv6MetricFsmStateEnum
	ESTABLISHED Bgpv6MetricFsmStateEnum
}{
	IDLE:        Bgpv6MetricFsmStateEnum("idle"),
	CONNECT:     Bgpv6MetricFsmStateEnum("connect"),
	ACTIVE:      Bgpv6MetricFsmStateEnum("active"),
	OPENSENT:    Bgpv6MetricFsmStateEnum("opensent"),
	OPENCONFIRM: Bgpv6MetricFsmStateEnum("openconfirm"),
	ESTABLISHED: Bgpv6MetricFsmStateEnum("established"),
}

func (obj *bgpv6Metric) FsmState() Bgpv6MetricFsmStateEnum {
	return Bgpv6MetricFsmStateEnum(obj.obj.FsmState.Enum().String())
}

// BGP peer FSM (Finite State Machine) state as Idle, Connect, Active, OpenSent, OpenConfirm and Established. In all the states except Established the BGP session is down. Idle refers to the Idle state of the FSM. Connect refers to the state where the session is waiting for the underlying transport session to be established. Active refers to the state where the session is awaiting for a connection from the remote peer. OpenSent refers to the state where the session is in the process of being established. The local system has sent an OPEN message. OpenConfirm refers to the state where the session is in the process of being established. The local system has sent and received an OPEN message and is awaiting a NOTIFICATION or KEEPALIVE message from remote peer. Established refers to the state where the BGP session with the peer is established.
// FsmState returns a string
func (obj *bgpv6Metric) HasFsmState() bool {
	return obj.obj.FsmState != nil
}

func (obj *bgpv6Metric) SetFsmState(value Bgpv6MetricFsmStateEnum) Bgpv6Metric {
	intValue, ok := otg.Bgpv6Metric_FsmState_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Bgpv6MetricFsmStateEnum", string(value)))
		return obj
	}
	enumValue := otg.Bgpv6Metric_FsmState_Enum(intValue)
	obj.obj.FsmState = &enumValue

	return obj
}

// Number of End-of-RIB markers received indicating the completion of the initial routing update for a  particular <AFI, SAFI> address family after the session is established. For the IPv4 unicast address family, the End-of-RIB marker is an UPDATE message with the minimum length. For any other address family, it is an UPDATE message that contains only the MP_UNREACH_NLRI attribute with  no withdrawn routes for that <AFI, SAFI>.
// EndOfRibReceived returns a uint64
func (obj *bgpv6Metric) EndOfRibReceived() uint64 {

	return *obj.obj.EndOfRibReceived

}

// Number of End-of-RIB markers received indicating the completion of the initial routing update for a  particular <AFI, SAFI> address family after the session is established. For the IPv4 unicast address family, the End-of-RIB marker is an UPDATE message with the minimum length. For any other address family, it is an UPDATE message that contains only the MP_UNREACH_NLRI attribute with  no withdrawn routes for that <AFI, SAFI>.
// EndOfRibReceived returns a uint64
func (obj *bgpv6Metric) HasEndOfRibReceived() bool {
	return obj.obj.EndOfRibReceived != nil
}

// Number of End-of-RIB markers received indicating the completion of the initial routing update for a  particular <AFI, SAFI> address family after the session is established. For the IPv4 unicast address family, the End-of-RIB marker is an UPDATE message with the minimum length. For any other address family, it is an UPDATE message that contains only the MP_UNREACH_NLRI attribute with  no withdrawn routes for that <AFI, SAFI>.
// SetEndOfRibReceived sets the uint64 value in the Bgpv6Metric object
func (obj *bgpv6Metric) SetEndOfRibReceived(value uint64) Bgpv6Metric {

	obj.obj.EndOfRibReceived = &value
	return obj
}

func (obj *bgpv6Metric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *bgpv6Metric) setDefault() {

}
