package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ServerMetric *****
type dhcpv6ServerMetric struct {
	validation
	obj          *otg.Dhcpv6ServerMetric
	marshaller   marshalDhcpv6ServerMetric
	unMarshaller unMarshalDhcpv6ServerMetric
}

func NewDhcpv6ServerMetric() Dhcpv6ServerMetric {
	obj := dhcpv6ServerMetric{obj: &otg.Dhcpv6ServerMetric{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ServerMetric) msg() *otg.Dhcpv6ServerMetric {
	return obj.obj
}

func (obj *dhcpv6ServerMetric) setMsg(msg *otg.Dhcpv6ServerMetric) Dhcpv6ServerMetric {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ServerMetric struct {
	obj *dhcpv6ServerMetric
}

type marshalDhcpv6ServerMetric interface {
	// ToProto marshals Dhcpv6ServerMetric to protobuf object *otg.Dhcpv6ServerMetric
	ToProto() (*otg.Dhcpv6ServerMetric, error)
	// ToPbText marshals Dhcpv6ServerMetric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ServerMetric to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ServerMetric to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Dhcpv6ServerMetric to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldhcpv6ServerMetric struct {
	obj *dhcpv6ServerMetric
}

type unMarshalDhcpv6ServerMetric interface {
	// FromProto unmarshals Dhcpv6ServerMetric from protobuf object *otg.Dhcpv6ServerMetric
	FromProto(msg *otg.Dhcpv6ServerMetric) (Dhcpv6ServerMetric, error)
	// FromPbText unmarshals Dhcpv6ServerMetric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ServerMetric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ServerMetric from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ServerMetric) Marshal() marshalDhcpv6ServerMetric {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ServerMetric{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ServerMetric) Unmarshal() unMarshalDhcpv6ServerMetric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ServerMetric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ServerMetric) ToProto() (*otg.Dhcpv6ServerMetric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ServerMetric) FromProto(msg *otg.Dhcpv6ServerMetric) (Dhcpv6ServerMetric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ServerMetric) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ServerMetric) FromPbText(value string) error {
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

func (m *marshaldhcpv6ServerMetric) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ServerMetric) FromYaml(value string) error {
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

func (m *marshaldhcpv6ServerMetric) ToJsonRaw() (string, error) {
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

func (m *marshaldhcpv6ServerMetric) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ServerMetric) FromJson(value string) error {
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

func (obj *dhcpv6ServerMetric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ServerMetric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ServerMetric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ServerMetric) Clone() (Dhcpv6ServerMetric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ServerMetric()
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

// Dhcpv6ServerMetric is dHCPv6 per server statistics information.
type Dhcpv6ServerMetric interface {
	Validation
	// msg marshals Dhcpv6ServerMetric to protobuf object *otg.Dhcpv6ServerMetric
	// and doesn't set defaults
	msg() *otg.Dhcpv6ServerMetric
	// setMsg unmarshals Dhcpv6ServerMetric from protobuf object *otg.Dhcpv6ServerMetric
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ServerMetric) Dhcpv6ServerMetric
	// provides marshal interface
	Marshal() marshalDhcpv6ServerMetric
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ServerMetric
	// validate validates Dhcpv6ServerMetric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ServerMetric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in Dhcpv6ServerMetric.
	Name() string
	// SetName assigns string provided by user to Dhcpv6ServerMetric
	SetName(value string) Dhcpv6ServerMetric
	// HasName checks if Name has been set in Dhcpv6ServerMetric
	HasName() bool
	// SolicitsReceived returns uint64, set in Dhcpv6ServerMetric.
	SolicitsReceived() uint64
	// SetSolicitsReceived assigns uint64 provided by user to Dhcpv6ServerMetric
	SetSolicitsReceived(value uint64) Dhcpv6ServerMetric
	// HasSolicitsReceived checks if SolicitsReceived has been set in Dhcpv6ServerMetric
	HasSolicitsReceived() bool
	// SolicitsIgnored returns uint64, set in Dhcpv6ServerMetric.
	SolicitsIgnored() uint64
	// SetSolicitsIgnored assigns uint64 provided by user to Dhcpv6ServerMetric
	SetSolicitsIgnored(value uint64) Dhcpv6ServerMetric
	// HasSolicitsIgnored checks if SolicitsIgnored has been set in Dhcpv6ServerMetric
	HasSolicitsIgnored() bool
	// AdvertisementsSent returns uint64, set in Dhcpv6ServerMetric.
	AdvertisementsSent() uint64
	// SetAdvertisementsSent assigns uint64 provided by user to Dhcpv6ServerMetric
	SetAdvertisementsSent(value uint64) Dhcpv6ServerMetric
	// HasAdvertisementsSent checks if AdvertisementsSent has been set in Dhcpv6ServerMetric
	HasAdvertisementsSent() bool
	// RequestsReceived returns uint64, set in Dhcpv6ServerMetric.
	RequestsReceived() uint64
	// SetRequestsReceived assigns uint64 provided by user to Dhcpv6ServerMetric
	SetRequestsReceived(value uint64) Dhcpv6ServerMetric
	// HasRequestsReceived checks if RequestsReceived has been set in Dhcpv6ServerMetric
	HasRequestsReceived() bool
	// NacksSent returns uint64, set in Dhcpv6ServerMetric.
	NacksSent() uint64
	// SetNacksSent assigns uint64 provided by user to Dhcpv6ServerMetric
	SetNacksSent(value uint64) Dhcpv6ServerMetric
	// HasNacksSent checks if NacksSent has been set in Dhcpv6ServerMetric
	HasNacksSent() bool
	// ConfirmsReceived returns uint64, set in Dhcpv6ServerMetric.
	ConfirmsReceived() uint64
	// SetConfirmsReceived assigns uint64 provided by user to Dhcpv6ServerMetric
	SetConfirmsReceived(value uint64) Dhcpv6ServerMetric
	// HasConfirmsReceived checks if ConfirmsReceived has been set in Dhcpv6ServerMetric
	HasConfirmsReceived() bool
	// RenewalsReceived returns uint64, set in Dhcpv6ServerMetric.
	RenewalsReceived() uint64
	// SetRenewalsReceived assigns uint64 provided by user to Dhcpv6ServerMetric
	SetRenewalsReceived(value uint64) Dhcpv6ServerMetric
	// HasRenewalsReceived checks if RenewalsReceived has been set in Dhcpv6ServerMetric
	HasRenewalsReceived() bool
	// RebindsReceived returns uint64, set in Dhcpv6ServerMetric.
	RebindsReceived() uint64
	// SetRebindsReceived assigns uint64 provided by user to Dhcpv6ServerMetric
	SetRebindsReceived(value uint64) Dhcpv6ServerMetric
	// HasRebindsReceived checks if RebindsReceived has been set in Dhcpv6ServerMetric
	HasRebindsReceived() bool
	// RepliesSent returns uint64, set in Dhcpv6ServerMetric.
	RepliesSent() uint64
	// SetRepliesSent assigns uint64 provided by user to Dhcpv6ServerMetric
	SetRepliesSent(value uint64) Dhcpv6ServerMetric
	// HasRepliesSent checks if RepliesSent has been set in Dhcpv6ServerMetric
	HasRepliesSent() bool
	// ReleasesReceived returns uint64, set in Dhcpv6ServerMetric.
	ReleasesReceived() uint64
	// SetReleasesReceived assigns uint64 provided by user to Dhcpv6ServerMetric
	SetReleasesReceived(value uint64) Dhcpv6ServerMetric
	// HasReleasesReceived checks if ReleasesReceived has been set in Dhcpv6ServerMetric
	HasReleasesReceived() bool
	// DeclinesReceived returns uint64, set in Dhcpv6ServerMetric.
	DeclinesReceived() uint64
	// SetDeclinesReceived assigns uint64 provided by user to Dhcpv6ServerMetric
	SetDeclinesReceived(value uint64) Dhcpv6ServerMetric
	// HasDeclinesReceived checks if DeclinesReceived has been set in Dhcpv6ServerMetric
	HasDeclinesReceived() bool
	// InformationRequestsReceived returns uint64, set in Dhcpv6ServerMetric.
	InformationRequestsReceived() uint64
	// SetInformationRequestsReceived assigns uint64 provided by user to Dhcpv6ServerMetric
	SetInformationRequestsReceived(value uint64) Dhcpv6ServerMetric
	// HasInformationRequestsReceived checks if InformationRequestsReceived has been set in Dhcpv6ServerMetric
	HasInformationRequestsReceived() bool
	// RelayForwardsReceived returns uint64, set in Dhcpv6ServerMetric.
	RelayForwardsReceived() uint64
	// SetRelayForwardsReceived assigns uint64 provided by user to Dhcpv6ServerMetric
	SetRelayForwardsReceived(value uint64) Dhcpv6ServerMetric
	// HasRelayForwardsReceived checks if RelayForwardsReceived has been set in Dhcpv6ServerMetric
	HasRelayForwardsReceived() bool
	// RelayRepliesSent returns uint64, set in Dhcpv6ServerMetric.
	RelayRepliesSent() uint64
	// SetRelayRepliesSent assigns uint64 provided by user to Dhcpv6ServerMetric
	SetRelayRepliesSent(value uint64) Dhcpv6ServerMetric
	// HasRelayRepliesSent checks if RelayRepliesSent has been set in Dhcpv6ServerMetric
	HasRelayRepliesSent() bool
	// ReconfiguresSent returns uint64, set in Dhcpv6ServerMetric.
	ReconfiguresSent() uint64
	// SetReconfiguresSent assigns uint64 provided by user to Dhcpv6ServerMetric
	SetReconfiguresSent(value uint64) Dhcpv6ServerMetric
	// HasReconfiguresSent checks if ReconfiguresSent has been set in Dhcpv6ServerMetric
	HasReconfiguresSent() bool
}

// The name of a configured DHCPv6 Server.
// Name returns a string
func (obj *dhcpv6ServerMetric) Name() string {

	return *obj.obj.Name

}

// The name of a configured DHCPv6 Server.
// Name returns a string
func (obj *dhcpv6ServerMetric) HasName() bool {
	return obj.obj.Name != nil
}

// The name of a configured DHCPv6 Server.
// SetName sets the string value in the Dhcpv6ServerMetric object
func (obj *dhcpv6ServerMetric) SetName(value string) Dhcpv6ServerMetric {

	obj.obj.Name = &value
	return obj
}

// Number of DHCPSOLICIT messages received.
// SolicitsReceived returns a uint64
func (obj *dhcpv6ServerMetric) SolicitsReceived() uint64 {

	return *obj.obj.SolicitsReceived

}

// Number of DHCPSOLICIT messages received.
// SolicitsReceived returns a uint64
func (obj *dhcpv6ServerMetric) HasSolicitsReceived() bool {
	return obj.obj.SolicitsReceived != nil
}

// Number of DHCPSOLICIT messages received.
// SetSolicitsReceived sets the uint64 value in the Dhcpv6ServerMetric object
func (obj *dhcpv6ServerMetric) SetSolicitsReceived(value uint64) Dhcpv6ServerMetric {

	obj.obj.SolicitsReceived = &value
	return obj
}

// Number of DHCPSOLICIT messages ignored.
// SolicitsIgnored returns a uint64
func (obj *dhcpv6ServerMetric) SolicitsIgnored() uint64 {

	return *obj.obj.SolicitsIgnored

}

// Number of DHCPSOLICIT messages ignored.
// SolicitsIgnored returns a uint64
func (obj *dhcpv6ServerMetric) HasSolicitsIgnored() bool {
	return obj.obj.SolicitsIgnored != nil
}

// Number of DHCPSOLICIT messages ignored.
// SetSolicitsIgnored sets the uint64 value in the Dhcpv6ServerMetric object
func (obj *dhcpv6ServerMetric) SetSolicitsIgnored(value uint64) Dhcpv6ServerMetric {

	obj.obj.SolicitsIgnored = &value
	return obj
}

// Number of DHCP Advertise messages sent.
// AdvertisementsSent returns a uint64
func (obj *dhcpv6ServerMetric) AdvertisementsSent() uint64 {

	return *obj.obj.AdvertisementsSent

}

// Number of DHCP Advertise messages sent.
// AdvertisementsSent returns a uint64
func (obj *dhcpv6ServerMetric) HasAdvertisementsSent() bool {
	return obj.obj.AdvertisementsSent != nil
}

// Number of DHCP Advertise messages sent.
// SetAdvertisementsSent sets the uint64 value in the Dhcpv6ServerMetric object
func (obj *dhcpv6ServerMetric) SetAdvertisementsSent(value uint64) Dhcpv6ServerMetric {

	obj.obj.AdvertisementsSent = &value
	return obj
}

// Number of DHCPREQUEST messages received.
// RequestsReceived returns a uint64
func (obj *dhcpv6ServerMetric) RequestsReceived() uint64 {

	return *obj.obj.RequestsReceived

}

// Number of DHCPREQUEST messages received.
// RequestsReceived returns a uint64
func (obj *dhcpv6ServerMetric) HasRequestsReceived() bool {
	return obj.obj.RequestsReceived != nil
}

// Number of DHCPREQUEST messages received.
// SetRequestsReceived sets the uint64 value in the Dhcpv6ServerMetric object
func (obj *dhcpv6ServerMetric) SetRequestsReceived(value uint64) Dhcpv6ServerMetric {

	obj.obj.RequestsReceived = &value
	return obj
}

// Number of naks sent for DHCPREQUEST messages.
// NacksSent returns a uint64
func (obj *dhcpv6ServerMetric) NacksSent() uint64 {

	return *obj.obj.NacksSent

}

// Number of naks sent for DHCPREQUEST messages.
// NacksSent returns a uint64
func (obj *dhcpv6ServerMetric) HasNacksSent() bool {
	return obj.obj.NacksSent != nil
}

// Number of naks sent for DHCPREQUEST messages.
// SetNacksSent sets the uint64 value in the Dhcpv6ServerMetric object
func (obj *dhcpv6ServerMetric) SetNacksSent(value uint64) Dhcpv6ServerMetric {

	obj.obj.NacksSent = &value
	return obj
}

// Number of DHCP Confirm messages received.
// ConfirmsReceived returns a uint64
func (obj *dhcpv6ServerMetric) ConfirmsReceived() uint64 {

	return *obj.obj.ConfirmsReceived

}

// Number of DHCP Confirm messages received.
// ConfirmsReceived returns a uint64
func (obj *dhcpv6ServerMetric) HasConfirmsReceived() bool {
	return obj.obj.ConfirmsReceived != nil
}

// Number of DHCP Confirm messages received.
// SetConfirmsReceived sets the uint64 value in the Dhcpv6ServerMetric object
func (obj *dhcpv6ServerMetric) SetConfirmsReceived(value uint64) Dhcpv6ServerMetric {

	obj.obj.ConfirmsReceived = &value
	return obj
}

// Number of DHCP Renewal messages received.
// RenewalsReceived returns a uint64
func (obj *dhcpv6ServerMetric) RenewalsReceived() uint64 {

	return *obj.obj.RenewalsReceived

}

// Number of DHCP Renewal messages received.
// RenewalsReceived returns a uint64
func (obj *dhcpv6ServerMetric) HasRenewalsReceived() bool {
	return obj.obj.RenewalsReceived != nil
}

// Number of DHCP Renewal messages received.
// SetRenewalsReceived sets the uint64 value in the Dhcpv6ServerMetric object
func (obj *dhcpv6ServerMetric) SetRenewalsReceived(value uint64) Dhcpv6ServerMetric {

	obj.obj.RenewalsReceived = &value
	return obj
}

// Number of DHCP Rebind messages received.
// RebindsReceived returns a uint64
func (obj *dhcpv6ServerMetric) RebindsReceived() uint64 {

	return *obj.obj.RebindsReceived

}

// Number of DHCP Rebind messages received.
// RebindsReceived returns a uint64
func (obj *dhcpv6ServerMetric) HasRebindsReceived() bool {
	return obj.obj.RebindsReceived != nil
}

// Number of DHCP Rebind messages received.
// SetRebindsReceived sets the uint64 value in the Dhcpv6ServerMetric object
func (obj *dhcpv6ServerMetric) SetRebindsReceived(value uint64) Dhcpv6ServerMetric {

	obj.obj.RebindsReceived = &value
	return obj
}

// Number of DHCP Reply messages sent.
// RepliesSent returns a uint64
func (obj *dhcpv6ServerMetric) RepliesSent() uint64 {

	return *obj.obj.RepliesSent

}

// Number of DHCP Reply messages sent.
// RepliesSent returns a uint64
func (obj *dhcpv6ServerMetric) HasRepliesSent() bool {
	return obj.obj.RepliesSent != nil
}

// Number of DHCP Reply messages sent.
// SetRepliesSent sets the uint64 value in the Dhcpv6ServerMetric object
func (obj *dhcpv6ServerMetric) SetRepliesSent(value uint64) Dhcpv6ServerMetric {

	obj.obj.RepliesSent = &value
	return obj
}

// Number of DHCP Release messages received.
// ReleasesReceived returns a uint64
func (obj *dhcpv6ServerMetric) ReleasesReceived() uint64 {

	return *obj.obj.ReleasesReceived

}

// Number of DHCP Release messages received.
// ReleasesReceived returns a uint64
func (obj *dhcpv6ServerMetric) HasReleasesReceived() bool {
	return obj.obj.ReleasesReceived != nil
}

// Number of DHCP Release messages received.
// SetReleasesReceived sets the uint64 value in the Dhcpv6ServerMetric object
func (obj *dhcpv6ServerMetric) SetReleasesReceived(value uint64) Dhcpv6ServerMetric {

	obj.obj.ReleasesReceived = &value
	return obj
}

// Number of DHCP Decline messages received.
// DeclinesReceived returns a uint64
func (obj *dhcpv6ServerMetric) DeclinesReceived() uint64 {

	return *obj.obj.DeclinesReceived

}

// Number of DHCP Decline messages received.
// DeclinesReceived returns a uint64
func (obj *dhcpv6ServerMetric) HasDeclinesReceived() bool {
	return obj.obj.DeclinesReceived != nil
}

// Number of DHCP Decline messages received.
// SetDeclinesReceived sets the uint64 value in the Dhcpv6ServerMetric object
func (obj *dhcpv6ServerMetric) SetDeclinesReceived(value uint64) Dhcpv6ServerMetric {

	obj.obj.DeclinesReceived = &value
	return obj
}

// Number of DHCP Information Request messages received.
// InformationRequestsReceived returns a uint64
func (obj *dhcpv6ServerMetric) InformationRequestsReceived() uint64 {

	return *obj.obj.InformationRequestsReceived

}

// Number of DHCP Information Request messages received.
// InformationRequestsReceived returns a uint64
func (obj *dhcpv6ServerMetric) HasInformationRequestsReceived() bool {
	return obj.obj.InformationRequestsReceived != nil
}

// Number of DHCP Information Request messages received.
// SetInformationRequestsReceived sets the uint64 value in the Dhcpv6ServerMetric object
func (obj *dhcpv6ServerMetric) SetInformationRequestsReceived(value uint64) Dhcpv6ServerMetric {

	obj.obj.InformationRequestsReceived = &value
	return obj
}

// Number of DHCP Relay agent forward messages received.
// RelayForwardsReceived returns a uint64
func (obj *dhcpv6ServerMetric) RelayForwardsReceived() uint64 {

	return *obj.obj.RelayForwardsReceived

}

// Number of DHCP Relay agent forward messages received.
// RelayForwardsReceived returns a uint64
func (obj *dhcpv6ServerMetric) HasRelayForwardsReceived() bool {
	return obj.obj.RelayForwardsReceived != nil
}

// Number of DHCP Relay agent forward messages received.
// SetRelayForwardsReceived sets the uint64 value in the Dhcpv6ServerMetric object
func (obj *dhcpv6ServerMetric) SetRelayForwardsReceived(value uint64) Dhcpv6ServerMetric {

	obj.obj.RelayForwardsReceived = &value
	return obj
}

// Number of DHCP reply messages sent to Relay agent.
// RelayRepliesSent returns a uint64
func (obj *dhcpv6ServerMetric) RelayRepliesSent() uint64 {

	return *obj.obj.RelayRepliesSent

}

// Number of DHCP reply messages sent to Relay agent.
// RelayRepliesSent returns a uint64
func (obj *dhcpv6ServerMetric) HasRelayRepliesSent() bool {
	return obj.obj.RelayRepliesSent != nil
}

// Number of DHCP reply messages sent to Relay agent.
// SetRelayRepliesSent sets the uint64 value in the Dhcpv6ServerMetric object
func (obj *dhcpv6ServerMetric) SetRelayRepliesSent(value uint64) Dhcpv6ServerMetric {

	obj.obj.RelayRepliesSent = &value
	return obj
}

// Number of DHCP Reconfigure messages sent.
// ReconfiguresSent returns a uint64
func (obj *dhcpv6ServerMetric) ReconfiguresSent() uint64 {

	return *obj.obj.ReconfiguresSent

}

// Number of DHCP Reconfigure messages sent.
// ReconfiguresSent returns a uint64
func (obj *dhcpv6ServerMetric) HasReconfiguresSent() bool {
	return obj.obj.ReconfiguresSent != nil
}

// Number of DHCP Reconfigure messages sent.
// SetReconfiguresSent sets the uint64 value in the Dhcpv6ServerMetric object
func (obj *dhcpv6ServerMetric) SetReconfiguresSent(value uint64) Dhcpv6ServerMetric {

	obj.obj.ReconfiguresSent = &value
	return obj
}

func (obj *dhcpv6ServerMetric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *dhcpv6ServerMetric) setDefault() {

}
