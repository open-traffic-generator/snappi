package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LdpMetric *****
type ldpMetric struct {
	validation
	obj          *otg.LdpMetric
	marshaller   marshalLdpMetric
	unMarshaller unMarshalLdpMetric
}

func NewLdpMetric() LdpMetric {
	obj := ldpMetric{obj: &otg.LdpMetric{}}
	obj.setDefault()
	return &obj
}

func (obj *ldpMetric) msg() *otg.LdpMetric {
	return obj.obj
}

func (obj *ldpMetric) setMsg(msg *otg.LdpMetric) LdpMetric {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalldpMetric struct {
	obj *ldpMetric
}

type marshalLdpMetric interface {
	// ToProto marshals LdpMetric to protobuf object *otg.LdpMetric
	ToProto() (*otg.LdpMetric, error)
	// ToPbText marshals LdpMetric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LdpMetric to YAML text
	ToYaml() (string, error)
	// ToJson marshals LdpMetric to JSON text
	ToJson() (string, error)
}

type unMarshalldpMetric struct {
	obj *ldpMetric
}

type unMarshalLdpMetric interface {
	// FromProto unmarshals LdpMetric from protobuf object *otg.LdpMetric
	FromProto(msg *otg.LdpMetric) (LdpMetric, error)
	// FromPbText unmarshals LdpMetric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LdpMetric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LdpMetric from JSON text
	FromJson(value string) error
}

func (obj *ldpMetric) Marshal() marshalLdpMetric {
	if obj.marshaller == nil {
		obj.marshaller = &marshalldpMetric{obj: obj}
	}
	return obj.marshaller
}

func (obj *ldpMetric) Unmarshal() unMarshalLdpMetric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalldpMetric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalldpMetric) ToProto() (*otg.LdpMetric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalldpMetric) FromProto(msg *otg.LdpMetric) (LdpMetric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalldpMetric) ToPbText() (string, error) {
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

func (m *unMarshalldpMetric) FromPbText(value string) error {
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

func (m *marshalldpMetric) ToYaml() (string, error) {
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

func (m *unMarshalldpMetric) FromYaml(value string) error {
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

func (m *marshalldpMetric) ToJson() (string, error) {
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

func (m *unMarshalldpMetric) FromJson(value string) error {
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

func (obj *ldpMetric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ldpMetric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ldpMetric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ldpMetric) Clone() (LdpMetric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLdpMetric()
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

// LdpMetric is lDP per router statistics information. A counter that is not reported by the implementation is absent.
type LdpMetric interface {
	Validation
	// msg marshals LdpMetric to protobuf object *otg.LdpMetric
	// and doesn't set defaults
	msg() *otg.LdpMetric
	// setMsg unmarshals LdpMetric from protobuf object *otg.LdpMetric
	// and doesn't set defaults
	setMsg(*otg.LdpMetric) LdpMetric
	// provides marshal interface
	Marshal() marshalLdpMetric
	// provides unmarshal interface
	Unmarshal() unMarshalLdpMetric
	// validate validates LdpMetric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LdpMetric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in LdpMetric.
	Name() string
	// SetName assigns string provided by user to LdpMetric
	SetName(value string) LdpMetric
	// HasName checks if Name has been set in LdpMetric
	HasName() bool
	// SessionsUp returns uint64, set in LdpMetric.
	SessionsUp() uint64
	// SetSessionsUp assigns uint64 provided by user to LdpMetric
	SetSessionsUp(value uint64) LdpMetric
	// HasSessionsUp checks if SessionsUp has been set in LdpMetric
	HasSessionsUp() bool
	// SessionsFlap returns uint64, set in LdpMetric.
	SessionsFlap() uint64
	// SetSessionsFlap assigns uint64 provided by user to LdpMetric
	SetSessionsFlap(value uint64) LdpMetric
	// HasSessionsFlap checks if SessionsFlap has been set in LdpMetric
	HasSessionsFlap() bool
	// SessionsNonExistent returns uint64, set in LdpMetric.
	SessionsNonExistent() uint64
	// SetSessionsNonExistent assigns uint64 provided by user to LdpMetric
	SetSessionsNonExistent(value uint64) LdpMetric
	// HasSessionsNonExistent checks if SessionsNonExistent has been set in LdpMetric
	HasSessionsNonExistent() bool
	// SessionsInitialized returns uint64, set in LdpMetric.
	SessionsInitialized() uint64
	// SetSessionsInitialized assigns uint64 provided by user to LdpMetric
	SetSessionsInitialized(value uint64) LdpMetric
	// HasSessionsInitialized checks if SessionsInitialized has been set in LdpMetric
	HasSessionsInitialized() bool
	// SessionsOpenReceived returns uint64, set in LdpMetric.
	SessionsOpenReceived() uint64
	// SetSessionsOpenReceived assigns uint64 provided by user to LdpMetric
	SetSessionsOpenReceived(value uint64) LdpMetric
	// HasSessionsOpenReceived checks if SessionsOpenReceived has been set in LdpMetric
	HasSessionsOpenReceived() bool
	// SessionsOpenSent returns uint64, set in LdpMetric.
	SessionsOpenSent() uint64
	// SetSessionsOpenSent assigns uint64 provided by user to LdpMetric
	SetSessionsOpenSent(value uint64) LdpMetric
	// HasSessionsOpenSent checks if SessionsOpenSent has been set in LdpMetric
	HasSessionsOpenSent() bool
	// SessionsOperational returns uint64, set in LdpMetric.
	SessionsOperational() uint64
	// SetSessionsOperational assigns uint64 provided by user to LdpMetric
	SetSessionsOperational(value uint64) LdpMetric
	// HasSessionsOperational checks if SessionsOperational has been set in LdpMetric
	HasSessionsOperational() bool
	// NotificationsSent returns uint64, set in LdpMetric.
	NotificationsSent() uint64
	// SetNotificationsSent assigns uint64 provided by user to LdpMetric
	SetNotificationsSent(value uint64) LdpMetric
	// HasNotificationsSent checks if NotificationsSent has been set in LdpMetric
	HasNotificationsSent() bool
	// NotificationsReceived returns uint64, set in LdpMetric.
	NotificationsReceived() uint64
	// SetNotificationsReceived assigns uint64 provided by user to LdpMetric
	SetNotificationsReceived(value uint64) LdpMetric
	// HasNotificationsReceived checks if NotificationsReceived has been set in LdpMetric
	HasNotificationsReceived() bool
	// LabelMappingsSent returns uint64, set in LdpMetric.
	LabelMappingsSent() uint64
	// SetLabelMappingsSent assigns uint64 provided by user to LdpMetric
	SetLabelMappingsSent(value uint64) LdpMetric
	// HasLabelMappingsSent checks if LabelMappingsSent has been set in LdpMetric
	HasLabelMappingsSent() bool
	// LabelMappingsReceived returns uint64, set in LdpMetric.
	LabelMappingsReceived() uint64
	// SetLabelMappingsReceived assigns uint64 provided by user to LdpMetric
	SetLabelMappingsReceived(value uint64) LdpMetric
	// HasLabelMappingsReceived checks if LabelMappingsReceived has been set in LdpMetric
	HasLabelMappingsReceived() bool
	// LabelRequestsSent returns uint64, set in LdpMetric.
	LabelRequestsSent() uint64
	// SetLabelRequestsSent assigns uint64 provided by user to LdpMetric
	SetLabelRequestsSent(value uint64) LdpMetric
	// HasLabelRequestsSent checks if LabelRequestsSent has been set in LdpMetric
	HasLabelRequestsSent() bool
	// LabelRequestsReceived returns uint64, set in LdpMetric.
	LabelRequestsReceived() uint64
	// SetLabelRequestsReceived assigns uint64 provided by user to LdpMetric
	SetLabelRequestsReceived(value uint64) LdpMetric
	// HasLabelRequestsReceived checks if LabelRequestsReceived has been set in LdpMetric
	HasLabelRequestsReceived() bool
	// LabelWithdrawsSent returns uint64, set in LdpMetric.
	LabelWithdrawsSent() uint64
	// SetLabelWithdrawsSent assigns uint64 provided by user to LdpMetric
	SetLabelWithdrawsSent(value uint64) LdpMetric
	// HasLabelWithdrawsSent checks if LabelWithdrawsSent has been set in LdpMetric
	HasLabelWithdrawsSent() bool
	// LabelWithdrawsReceived returns uint64, set in LdpMetric.
	LabelWithdrawsReceived() uint64
	// SetLabelWithdrawsReceived assigns uint64 provided by user to LdpMetric
	SetLabelWithdrawsReceived(value uint64) LdpMetric
	// HasLabelWithdrawsReceived checks if LabelWithdrawsReceived has been set in LdpMetric
	HasLabelWithdrawsReceived() bool
	// LabelReleasesSent returns uint64, set in LdpMetric.
	LabelReleasesSent() uint64
	// SetLabelReleasesSent assigns uint64 provided by user to LdpMetric
	SetLabelReleasesSent(value uint64) LdpMetric
	// HasLabelReleasesSent checks if LabelReleasesSent has been set in LdpMetric
	HasLabelReleasesSent() bool
	// LabelReleasesReceived returns uint64, set in LdpMetric.
	LabelReleasesReceived() uint64
	// SetLabelReleasesReceived assigns uint64 provided by user to LdpMetric
	SetLabelReleasesReceived(value uint64) LdpMetric
	// HasLabelReleasesReceived checks if LabelReleasesReceived has been set in LdpMetric
	HasLabelReleasesReceived() bool
	// LabelAbortRequestsSent returns uint64, set in LdpMetric.
	LabelAbortRequestsSent() uint64
	// SetLabelAbortRequestsSent assigns uint64 provided by user to LdpMetric
	SetLabelAbortRequestsSent(value uint64) LdpMetric
	// HasLabelAbortRequestsSent checks if LabelAbortRequestsSent has been set in LdpMetric
	HasLabelAbortRequestsSent() bool
	// LabelAbortRequestsReceived returns uint64, set in LdpMetric.
	LabelAbortRequestsReceived() uint64
	// SetLabelAbortRequestsReceived assigns uint64 provided by user to LdpMetric
	SetLabelAbortRequestsReceived(value uint64) LdpMetric
	// HasLabelAbortRequestsReceived checks if LabelAbortRequestsReceived has been set in LdpMetric
	HasLabelAbortRequestsReceived() bool
	// IngressLspsUp returns uint64, set in LdpMetric.
	IngressLspsUp() uint64
	// SetIngressLspsUp assigns uint64 provided by user to LdpMetric
	SetIngressLspsUp(value uint64) LdpMetric
	// HasIngressLspsUp checks if IngressLspsUp has been set in LdpMetric
	HasIngressLspsUp() bool
	// EgressLspsUp returns uint64, set in LdpMetric.
	EgressLspsUp() uint64
	// SetEgressLspsUp assigns uint64 provided by user to LdpMetric
	SetEgressLspsUp(value uint64) LdpMetric
	// HasEgressLspsUp checks if EgressLspsUp has been set in LdpMetric
	HasEgressLspsUp() bool
}

// The name of a configured LDP router.
// Name returns a string
func (obj *ldpMetric) Name() string {

	return *obj.obj.Name

}

// The name of a configured LDP router.
// Name returns a string
func (obj *ldpMetric) HasName() bool {
	return obj.obj.Name != nil
}

// The name of a configured LDP router.
// SetName sets the string value in the LdpMetric object
func (obj *ldpMetric) SetName(value string) LdpMetric {

	obj.obj.Name = &value
	return obj
}

// The number of LDP sessions of this router that are up, that is in the OPERATIONAL state (RFC 5036 Section 2.5.4). Sessions formed through Basic (link) discovery (RFC 5036 Section 2.4.1) and through Extended (targeted) discovery (RFC 5036 Section 2.4.2) are both counted.
// SessionsUp returns a uint64
func (obj *ldpMetric) SessionsUp() uint64 {

	return *obj.obj.SessionsUp

}

// The number of LDP sessions of this router that are up, that is in the OPERATIONAL state (RFC 5036 Section 2.5.4). Sessions formed through Basic (link) discovery (RFC 5036 Section 2.4.1) and through Extended (targeted) discovery (RFC 5036 Section 2.4.2) are both counted.
// SessionsUp returns a uint64
func (obj *ldpMetric) HasSessionsUp() bool {
	return obj.obj.SessionsUp != nil
}

// The number of LDP sessions of this router that are up, that is in the OPERATIONAL state (RFC 5036 Section 2.5.4). Sessions formed through Basic (link) discovery (RFC 5036 Section 2.4.1) and through Extended (targeted) discovery (RFC 5036 Section 2.4.2) are both counted.
// SetSessionsUp sets the uint64 value in the LdpMetric object
func (obj *ldpMetric) SetSessionsUp(value uint64) LdpMetric {

	obj.obj.SessionsUp = &value
	return obj
}

// The number of times an LDP session of this router left the OPERATIONAL state (RFC 5036 Section 2.5.4).
// SessionsFlap returns a uint64
func (obj *ldpMetric) SessionsFlap() uint64 {

	return *obj.obj.SessionsFlap

}

// The number of times an LDP session of this router left the OPERATIONAL state (RFC 5036 Section 2.5.4).
// SessionsFlap returns a uint64
func (obj *ldpMetric) HasSessionsFlap() bool {
	return obj.obj.SessionsFlap != nil
}

// The number of times an LDP session of this router left the OPERATIONAL state (RFC 5036 Section 2.5.4).
// SetSessionsFlap sets the uint64 value in the LdpMetric object
func (obj *ldpMetric) SetSessionsFlap(value uint64) LdpMetric {

	obj.obj.SessionsFlap = &value
	return obj
}

// The number of LDP sessions of this router currently in the NON EXISTENT state of the session initialization state machine (RFC 5036 Section 2.5.4).
// SessionsNonExistent returns a uint64
func (obj *ldpMetric) SessionsNonExistent() uint64 {

	return *obj.obj.SessionsNonExistent

}

// The number of LDP sessions of this router currently in the NON EXISTENT state of the session initialization state machine (RFC 5036 Section 2.5.4).
// SessionsNonExistent returns a uint64
func (obj *ldpMetric) HasSessionsNonExistent() bool {
	return obj.obj.SessionsNonExistent != nil
}

// The number of LDP sessions of this router currently in the NON EXISTENT state of the session initialization state machine (RFC 5036 Section 2.5.4).
// SetSessionsNonExistent sets the uint64 value in the LdpMetric object
func (obj *ldpMetric) SetSessionsNonExistent(value uint64) LdpMetric {

	obj.obj.SessionsNonExistent = &value
	return obj
}

// The number of LDP sessions of this router currently in the INITIALIZED state of the session initialization state machine (RFC 5036 Section 2.5.4).
// SessionsInitialized returns a uint64
func (obj *ldpMetric) SessionsInitialized() uint64 {

	return *obj.obj.SessionsInitialized

}

// The number of LDP sessions of this router currently in the INITIALIZED state of the session initialization state machine (RFC 5036 Section 2.5.4).
// SessionsInitialized returns a uint64
func (obj *ldpMetric) HasSessionsInitialized() bool {
	return obj.obj.SessionsInitialized != nil
}

// The number of LDP sessions of this router currently in the INITIALIZED state of the session initialization state machine (RFC 5036 Section 2.5.4).
// SetSessionsInitialized sets the uint64 value in the LdpMetric object
func (obj *ldpMetric) SetSessionsInitialized(value uint64) LdpMetric {

	obj.obj.SessionsInitialized = &value
	return obj
}

// The number of LDP sessions of this router currently in the OPENREC state of the session initialization state machine (RFC 5036 Section 2.5.4).
// SessionsOpenReceived returns a uint64
func (obj *ldpMetric) SessionsOpenReceived() uint64 {

	return *obj.obj.SessionsOpenReceived

}

// The number of LDP sessions of this router currently in the OPENREC state of the session initialization state machine (RFC 5036 Section 2.5.4).
// SessionsOpenReceived returns a uint64
func (obj *ldpMetric) HasSessionsOpenReceived() bool {
	return obj.obj.SessionsOpenReceived != nil
}

// The number of LDP sessions of this router currently in the OPENREC state of the session initialization state machine (RFC 5036 Section 2.5.4).
// SetSessionsOpenReceived sets the uint64 value in the LdpMetric object
func (obj *ldpMetric) SetSessionsOpenReceived(value uint64) LdpMetric {

	obj.obj.SessionsOpenReceived = &value
	return obj
}

// The number of LDP sessions of this router currently in the OPENSENT state of the session initialization state machine (RFC 5036 Section 2.5.4).
// SessionsOpenSent returns a uint64
func (obj *ldpMetric) SessionsOpenSent() uint64 {

	return *obj.obj.SessionsOpenSent

}

// The number of LDP sessions of this router currently in the OPENSENT state of the session initialization state machine (RFC 5036 Section 2.5.4).
// SessionsOpenSent returns a uint64
func (obj *ldpMetric) HasSessionsOpenSent() bool {
	return obj.obj.SessionsOpenSent != nil
}

// The number of LDP sessions of this router currently in the OPENSENT state of the session initialization state machine (RFC 5036 Section 2.5.4).
// SetSessionsOpenSent sets the uint64 value in the LdpMetric object
func (obj *ldpMetric) SetSessionsOpenSent(value uint64) LdpMetric {

	obj.obj.SessionsOpenSent = &value
	return obj
}

// The number of LDP sessions of this router currently in the OPERATIONAL state of the session initialization state machine (RFC 5036 Section 2.5.4).
// SessionsOperational returns a uint64
func (obj *ldpMetric) SessionsOperational() uint64 {

	return *obj.obj.SessionsOperational

}

// The number of LDP sessions of this router currently in the OPERATIONAL state of the session initialization state machine (RFC 5036 Section 2.5.4).
// SessionsOperational returns a uint64
func (obj *ldpMetric) HasSessionsOperational() bool {
	return obj.obj.SessionsOperational != nil
}

// The number of LDP sessions of this router currently in the OPERATIONAL state of the session initialization state machine (RFC 5036 Section 2.5.4).
// SetSessionsOperational sets the uint64 value in the LdpMetric object
func (obj *ldpMetric) SetSessionsOperational(value uint64) LdpMetric {

	obj.obj.SessionsOperational = &value
	return obj
}

// The number of Notification messages (type 0x0001, RFC 5036 Section 3.5.1) sent.
// NotificationsSent returns a uint64
func (obj *ldpMetric) NotificationsSent() uint64 {

	return *obj.obj.NotificationsSent

}

// The number of Notification messages (type 0x0001, RFC 5036 Section 3.5.1) sent.
// NotificationsSent returns a uint64
func (obj *ldpMetric) HasNotificationsSent() bool {
	return obj.obj.NotificationsSent != nil
}

// The number of Notification messages (type 0x0001, RFC 5036 Section 3.5.1) sent.
// SetNotificationsSent sets the uint64 value in the LdpMetric object
func (obj *ldpMetric) SetNotificationsSent(value uint64) LdpMetric {

	obj.obj.NotificationsSent = &value
	return obj
}

// The number of Notification messages (type 0x0001, RFC 5036 Section 3.5.1) received.
// NotificationsReceived returns a uint64
func (obj *ldpMetric) NotificationsReceived() uint64 {

	return *obj.obj.NotificationsReceived

}

// The number of Notification messages (type 0x0001, RFC 5036 Section 3.5.1) received.
// NotificationsReceived returns a uint64
func (obj *ldpMetric) HasNotificationsReceived() bool {
	return obj.obj.NotificationsReceived != nil
}

// The number of Notification messages (type 0x0001, RFC 5036 Section 3.5.1) received.
// SetNotificationsReceived sets the uint64 value in the LdpMetric object
func (obj *ldpMetric) SetNotificationsReceived(value uint64) LdpMetric {

	obj.obj.NotificationsReceived = &value
	return obj
}

// The number of Label Mapping messages (type 0x0400, RFC 5036 Section 3.5.7) sent.
// LabelMappingsSent returns a uint64
func (obj *ldpMetric) LabelMappingsSent() uint64 {

	return *obj.obj.LabelMappingsSent

}

// The number of Label Mapping messages (type 0x0400, RFC 5036 Section 3.5.7) sent.
// LabelMappingsSent returns a uint64
func (obj *ldpMetric) HasLabelMappingsSent() bool {
	return obj.obj.LabelMappingsSent != nil
}

// The number of Label Mapping messages (type 0x0400, RFC 5036 Section 3.5.7) sent.
// SetLabelMappingsSent sets the uint64 value in the LdpMetric object
func (obj *ldpMetric) SetLabelMappingsSent(value uint64) LdpMetric {

	obj.obj.LabelMappingsSent = &value
	return obj
}

// The number of Label Mapping messages (type 0x0400, RFC 5036 Section 3.5.7) received.
// LabelMappingsReceived returns a uint64
func (obj *ldpMetric) LabelMappingsReceived() uint64 {

	return *obj.obj.LabelMappingsReceived

}

// The number of Label Mapping messages (type 0x0400, RFC 5036 Section 3.5.7) received.
// LabelMappingsReceived returns a uint64
func (obj *ldpMetric) HasLabelMappingsReceived() bool {
	return obj.obj.LabelMappingsReceived != nil
}

// The number of Label Mapping messages (type 0x0400, RFC 5036 Section 3.5.7) received.
// SetLabelMappingsReceived sets the uint64 value in the LdpMetric object
func (obj *ldpMetric) SetLabelMappingsReceived(value uint64) LdpMetric {

	obj.obj.LabelMappingsReceived = &value
	return obj
}

// The number of Label Request messages (type 0x0401, RFC 5036 Section 3.5.8) sent.
// LabelRequestsSent returns a uint64
func (obj *ldpMetric) LabelRequestsSent() uint64 {

	return *obj.obj.LabelRequestsSent

}

// The number of Label Request messages (type 0x0401, RFC 5036 Section 3.5.8) sent.
// LabelRequestsSent returns a uint64
func (obj *ldpMetric) HasLabelRequestsSent() bool {
	return obj.obj.LabelRequestsSent != nil
}

// The number of Label Request messages (type 0x0401, RFC 5036 Section 3.5.8) sent.
// SetLabelRequestsSent sets the uint64 value in the LdpMetric object
func (obj *ldpMetric) SetLabelRequestsSent(value uint64) LdpMetric {

	obj.obj.LabelRequestsSent = &value
	return obj
}

// The number of Label Request messages (type 0x0401, RFC 5036 Section 3.5.8) received.
// LabelRequestsReceived returns a uint64
func (obj *ldpMetric) LabelRequestsReceived() uint64 {

	return *obj.obj.LabelRequestsReceived

}

// The number of Label Request messages (type 0x0401, RFC 5036 Section 3.5.8) received.
// LabelRequestsReceived returns a uint64
func (obj *ldpMetric) HasLabelRequestsReceived() bool {
	return obj.obj.LabelRequestsReceived != nil
}

// The number of Label Request messages (type 0x0401, RFC 5036 Section 3.5.8) received.
// SetLabelRequestsReceived sets the uint64 value in the LdpMetric object
func (obj *ldpMetric) SetLabelRequestsReceived(value uint64) LdpMetric {

	obj.obj.LabelRequestsReceived = &value
	return obj
}

// The number of Label Withdraw messages (type 0x0402, RFC 5036 Section 3.5.10) sent.
// LabelWithdrawsSent returns a uint64
func (obj *ldpMetric) LabelWithdrawsSent() uint64 {

	return *obj.obj.LabelWithdrawsSent

}

// The number of Label Withdraw messages (type 0x0402, RFC 5036 Section 3.5.10) sent.
// LabelWithdrawsSent returns a uint64
func (obj *ldpMetric) HasLabelWithdrawsSent() bool {
	return obj.obj.LabelWithdrawsSent != nil
}

// The number of Label Withdraw messages (type 0x0402, RFC 5036 Section 3.5.10) sent.
// SetLabelWithdrawsSent sets the uint64 value in the LdpMetric object
func (obj *ldpMetric) SetLabelWithdrawsSent(value uint64) LdpMetric {

	obj.obj.LabelWithdrawsSent = &value
	return obj
}

// The number of Label Withdraw messages (type 0x0402, RFC 5036 Section 3.5.10) received.
// LabelWithdrawsReceived returns a uint64
func (obj *ldpMetric) LabelWithdrawsReceived() uint64 {

	return *obj.obj.LabelWithdrawsReceived

}

// The number of Label Withdraw messages (type 0x0402, RFC 5036 Section 3.5.10) received.
// LabelWithdrawsReceived returns a uint64
func (obj *ldpMetric) HasLabelWithdrawsReceived() bool {
	return obj.obj.LabelWithdrawsReceived != nil
}

// The number of Label Withdraw messages (type 0x0402, RFC 5036 Section 3.5.10) received.
// SetLabelWithdrawsReceived sets the uint64 value in the LdpMetric object
func (obj *ldpMetric) SetLabelWithdrawsReceived(value uint64) LdpMetric {

	obj.obj.LabelWithdrawsReceived = &value
	return obj
}

// The number of Label Release messages (type 0x0403, RFC 5036 Section 3.5.11) sent.
// LabelReleasesSent returns a uint64
func (obj *ldpMetric) LabelReleasesSent() uint64 {

	return *obj.obj.LabelReleasesSent

}

// The number of Label Release messages (type 0x0403, RFC 5036 Section 3.5.11) sent.
// LabelReleasesSent returns a uint64
func (obj *ldpMetric) HasLabelReleasesSent() bool {
	return obj.obj.LabelReleasesSent != nil
}

// The number of Label Release messages (type 0x0403, RFC 5036 Section 3.5.11) sent.
// SetLabelReleasesSent sets the uint64 value in the LdpMetric object
func (obj *ldpMetric) SetLabelReleasesSent(value uint64) LdpMetric {

	obj.obj.LabelReleasesSent = &value
	return obj
}

// The number of Label Release messages (type 0x0403, RFC 5036 Section 3.5.11) received.
// LabelReleasesReceived returns a uint64
func (obj *ldpMetric) LabelReleasesReceived() uint64 {

	return *obj.obj.LabelReleasesReceived

}

// The number of Label Release messages (type 0x0403, RFC 5036 Section 3.5.11) received.
// LabelReleasesReceived returns a uint64
func (obj *ldpMetric) HasLabelReleasesReceived() bool {
	return obj.obj.LabelReleasesReceived != nil
}

// The number of Label Release messages (type 0x0403, RFC 5036 Section 3.5.11) received.
// SetLabelReleasesReceived sets the uint64 value in the LdpMetric object
func (obj *ldpMetric) SetLabelReleasesReceived(value uint64) LdpMetric {

	obj.obj.LabelReleasesReceived = &value
	return obj
}

// The number of Label Abort Request messages (type 0x0404, RFC 5036 Section 3.5.9) sent.
// LabelAbortRequestsSent returns a uint64
func (obj *ldpMetric) LabelAbortRequestsSent() uint64 {

	return *obj.obj.LabelAbortRequestsSent

}

// The number of Label Abort Request messages (type 0x0404, RFC 5036 Section 3.5.9) sent.
// LabelAbortRequestsSent returns a uint64
func (obj *ldpMetric) HasLabelAbortRequestsSent() bool {
	return obj.obj.LabelAbortRequestsSent != nil
}

// The number of Label Abort Request messages (type 0x0404, RFC 5036 Section 3.5.9) sent.
// SetLabelAbortRequestsSent sets the uint64 value in the LdpMetric object
func (obj *ldpMetric) SetLabelAbortRequestsSent(value uint64) LdpMetric {

	obj.obj.LabelAbortRequestsSent = &value
	return obj
}

// The number of Label Abort Request messages (type 0x0404, RFC 5036 Section 3.5.9) received.
// LabelAbortRequestsReceived returns a uint64
func (obj *ldpMetric) LabelAbortRequestsReceived() uint64 {

	return *obj.obj.LabelAbortRequestsReceived

}

// The number of Label Abort Request messages (type 0x0404, RFC 5036 Section 3.5.9) received.
// LabelAbortRequestsReceived returns a uint64
func (obj *ldpMetric) HasLabelAbortRequestsReceived() bool {
	return obj.obj.LabelAbortRequestsReceived != nil
}

// The number of Label Abort Request messages (type 0x0404, RFC 5036 Section 3.5.9) received.
// SetLabelAbortRequestsReceived sets the uint64 value in the LdpMetric object
func (obj *ldpMetric) SetLabelAbortRequestsReceived(value uint64) LdpMetric {

	obj.obj.LabelAbortRequestsReceived = &value
	return obj
}

// The number of IPv4 Prefix FECs for which this LSR has received a label from a peer, that is LSPs for which this LSR is the ingress.
// IngressLspsUp returns a uint64
func (obj *ldpMetric) IngressLspsUp() uint64 {

	return *obj.obj.IngressLspsUp

}

// The number of IPv4 Prefix FECs for which this LSR has received a label from a peer, that is LSPs for which this LSR is the ingress.
// IngressLspsUp returns a uint64
func (obj *ldpMetric) HasIngressLspsUp() bool {
	return obj.obj.IngressLspsUp != nil
}

// The number of IPv4 Prefix FECs for which this LSR has received a label from a peer, that is LSPs for which this LSR is the ingress.
// SetIngressLspsUp sets the uint64 value in the LdpMetric object
func (obj *ldpMetric) SetIngressLspsUp(value uint64) LdpMetric {

	obj.obj.IngressLspsUp = &value
	return obj
}

// The number of IPv4 Prefix FECs advertised by this LSR for which a session to a peer is OPERATIONAL, that is LSPs for which this LSR is the egress.
// EgressLspsUp returns a uint64
func (obj *ldpMetric) EgressLspsUp() uint64 {

	return *obj.obj.EgressLspsUp

}

// The number of IPv4 Prefix FECs advertised by this LSR for which a session to a peer is OPERATIONAL, that is LSPs for which this LSR is the egress.
// EgressLspsUp returns a uint64
func (obj *ldpMetric) HasEgressLspsUp() bool {
	return obj.obj.EgressLspsUp != nil
}

// The number of IPv4 Prefix FECs advertised by this LSR for which a session to a peer is OPERATIONAL, that is LSPs for which this LSR is the egress.
// SetEgressLspsUp sets the uint64 value in the LdpMetric object
func (obj *ldpMetric) SetEgressLspsUp(value uint64) LdpMetric {

	obj.obj.EgressLspsUp = &value
	return obj
}

func (obj *ldpMetric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ldpMetric) setDefault() {

}
