package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ClientMetric *****
type dhcpv6ClientMetric struct {
	validation
	obj          *otg.Dhcpv6ClientMetric
	marshaller   marshalDhcpv6ClientMetric
	unMarshaller unMarshalDhcpv6ClientMetric
}

func NewDhcpv6ClientMetric() Dhcpv6ClientMetric {
	obj := dhcpv6ClientMetric{obj: &otg.Dhcpv6ClientMetric{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ClientMetric) msg() *otg.Dhcpv6ClientMetric {
	return obj.obj
}

func (obj *dhcpv6ClientMetric) setMsg(msg *otg.Dhcpv6ClientMetric) Dhcpv6ClientMetric {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ClientMetric struct {
	obj *dhcpv6ClientMetric
}

type marshalDhcpv6ClientMetric interface {
	// ToProto marshals Dhcpv6ClientMetric to protobuf object *otg.Dhcpv6ClientMetric
	ToProto() (*otg.Dhcpv6ClientMetric, error)
	// ToPbText marshals Dhcpv6ClientMetric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ClientMetric to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ClientMetric to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Dhcpv6ClientMetric to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldhcpv6ClientMetric struct {
	obj *dhcpv6ClientMetric
}

type unMarshalDhcpv6ClientMetric interface {
	// FromProto unmarshals Dhcpv6ClientMetric from protobuf object *otg.Dhcpv6ClientMetric
	FromProto(msg *otg.Dhcpv6ClientMetric) (Dhcpv6ClientMetric, error)
	// FromPbText unmarshals Dhcpv6ClientMetric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ClientMetric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ClientMetric from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ClientMetric) Marshal() marshalDhcpv6ClientMetric {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ClientMetric{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ClientMetric) Unmarshal() unMarshalDhcpv6ClientMetric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ClientMetric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ClientMetric) ToProto() (*otg.Dhcpv6ClientMetric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ClientMetric) FromProto(msg *otg.Dhcpv6ClientMetric) (Dhcpv6ClientMetric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ClientMetric) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ClientMetric) FromPbText(value string) error {
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

func (m *marshaldhcpv6ClientMetric) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ClientMetric) FromYaml(value string) error {
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

func (m *marshaldhcpv6ClientMetric) ToJsonRaw() (string, error) {
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

func (m *marshaldhcpv6ClientMetric) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ClientMetric) FromJson(value string) error {
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

func (obj *dhcpv6ClientMetric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ClientMetric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ClientMetric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ClientMetric) Clone() (Dhcpv6ClientMetric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ClientMetric()
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

// Dhcpv6ClientMetric is dHCPv6 per peer statistics information.
type Dhcpv6ClientMetric interface {
	Validation
	// msg marshals Dhcpv6ClientMetric to protobuf object *otg.Dhcpv6ClientMetric
	// and doesn't set defaults
	msg() *otg.Dhcpv6ClientMetric
	// setMsg unmarshals Dhcpv6ClientMetric from protobuf object *otg.Dhcpv6ClientMetric
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ClientMetric) Dhcpv6ClientMetric
	// provides marshal interface
	Marshal() marshalDhcpv6ClientMetric
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ClientMetric
	// validate validates Dhcpv6ClientMetric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ClientMetric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in Dhcpv6ClientMetric.
	Name() string
	// SetName assigns string provided by user to Dhcpv6ClientMetric
	SetName(value string) Dhcpv6ClientMetric
	// HasName checks if Name has been set in Dhcpv6ClientMetric
	HasName() bool
	// SolicitsSent returns uint64, set in Dhcpv6ClientMetric.
	SolicitsSent() uint64
	// SetSolicitsSent assigns uint64 provided by user to Dhcpv6ClientMetric
	SetSolicitsSent(value uint64) Dhcpv6ClientMetric
	// HasSolicitsSent checks if SolicitsSent has been set in Dhcpv6ClientMetric
	HasSolicitsSent() bool
	// AdvertisementsReceived returns uint64, set in Dhcpv6ClientMetric.
	AdvertisementsReceived() uint64
	// SetAdvertisementsReceived assigns uint64 provided by user to Dhcpv6ClientMetric
	SetAdvertisementsReceived(value uint64) Dhcpv6ClientMetric
	// HasAdvertisementsReceived checks if AdvertisementsReceived has been set in Dhcpv6ClientMetric
	HasAdvertisementsReceived() bool
	// AdvertisementsIgnored returns uint64, set in Dhcpv6ClientMetric.
	AdvertisementsIgnored() uint64
	// SetAdvertisementsIgnored assigns uint64 provided by user to Dhcpv6ClientMetric
	SetAdvertisementsIgnored(value uint64) Dhcpv6ClientMetric
	// HasAdvertisementsIgnored checks if AdvertisementsIgnored has been set in Dhcpv6ClientMetric
	HasAdvertisementsIgnored() bool
	// RequestsSent returns uint64, set in Dhcpv6ClientMetric.
	RequestsSent() uint64
	// SetRequestsSent assigns uint64 provided by user to Dhcpv6ClientMetric
	SetRequestsSent(value uint64) Dhcpv6ClientMetric
	// HasRequestsSent checks if RequestsSent has been set in Dhcpv6ClientMetric
	HasRequestsSent() bool
	// NacksReceived returns uint64, set in Dhcpv6ClientMetric.
	NacksReceived() uint64
	// SetNacksReceived assigns uint64 provided by user to Dhcpv6ClientMetric
	SetNacksReceived(value uint64) Dhcpv6ClientMetric
	// HasNacksReceived checks if NacksReceived has been set in Dhcpv6ClientMetric
	HasNacksReceived() bool
	// RepliesReceived returns uint64, set in Dhcpv6ClientMetric.
	RepliesReceived() uint64
	// SetRepliesReceived assigns uint64 provided by user to Dhcpv6ClientMetric
	SetRepliesReceived(value uint64) Dhcpv6ClientMetric
	// HasRepliesReceived checks if RepliesReceived has been set in Dhcpv6ClientMetric
	HasRepliesReceived() bool
	// InformationRequestsSent returns uint64, set in Dhcpv6ClientMetric.
	InformationRequestsSent() uint64
	// SetInformationRequestsSent assigns uint64 provided by user to Dhcpv6ClientMetric
	SetInformationRequestsSent(value uint64) Dhcpv6ClientMetric
	// HasInformationRequestsSent checks if InformationRequestsSent has been set in Dhcpv6ClientMetric
	HasInformationRequestsSent() bool
	// RenewsSent returns uint64, set in Dhcpv6ClientMetric.
	RenewsSent() uint64
	// SetRenewsSent assigns uint64 provided by user to Dhcpv6ClientMetric
	SetRenewsSent(value uint64) Dhcpv6ClientMetric
	// HasRenewsSent checks if RenewsSent has been set in Dhcpv6ClientMetric
	HasRenewsSent() bool
	// RebindsSent returns uint64, set in Dhcpv6ClientMetric.
	RebindsSent() uint64
	// SetRebindsSent assigns uint64 provided by user to Dhcpv6ClientMetric
	SetRebindsSent(value uint64) Dhcpv6ClientMetric
	// HasRebindsSent checks if RebindsSent has been set in Dhcpv6ClientMetric
	HasRebindsSent() bool
	// ReleasesSent returns uint64, set in Dhcpv6ClientMetric.
	ReleasesSent() uint64
	// SetReleasesSent assigns uint64 provided by user to Dhcpv6ClientMetric
	SetReleasesSent(value uint64) Dhcpv6ClientMetric
	// HasReleasesSent checks if ReleasesSent has been set in Dhcpv6ClientMetric
	HasReleasesSent() bool
	// ReconfiguresReceived returns uint64, set in Dhcpv6ClientMetric.
	ReconfiguresReceived() uint64
	// SetReconfiguresReceived assigns uint64 provided by user to Dhcpv6ClientMetric
	SetReconfiguresReceived(value uint64) Dhcpv6ClientMetric
	// HasReconfiguresReceived checks if ReconfiguresReceived has been set in Dhcpv6ClientMetric
	HasReconfiguresReceived() bool
	// RapidCommitSolicitsSent returns uint64, set in Dhcpv6ClientMetric.
	RapidCommitSolicitsSent() uint64
	// SetRapidCommitSolicitsSent assigns uint64 provided by user to Dhcpv6ClientMetric
	SetRapidCommitSolicitsSent(value uint64) Dhcpv6ClientMetric
	// HasRapidCommitSolicitsSent checks if RapidCommitSolicitsSent has been set in Dhcpv6ClientMetric
	HasRapidCommitSolicitsSent() bool
	// RapidCommitRepliesReceived returns uint64, set in Dhcpv6ClientMetric.
	RapidCommitRepliesReceived() uint64
	// SetRapidCommitRepliesReceived assigns uint64 provided by user to Dhcpv6ClientMetric
	SetRapidCommitRepliesReceived(value uint64) Dhcpv6ClientMetric
	// HasRapidCommitRepliesReceived checks if RapidCommitRepliesReceived has been set in Dhcpv6ClientMetric
	HasRapidCommitRepliesReceived() bool
}

// The name of a configured DHCPv6 client.
// Name returns a string
func (obj *dhcpv6ClientMetric) Name() string {

	return *obj.obj.Name

}

// The name of a configured DHCPv6 client.
// Name returns a string
func (obj *dhcpv6ClientMetric) HasName() bool {
	return obj.obj.Name != nil
}

// The name of a configured DHCPv6 client.
// SetName sets the string value in the Dhcpv6ClientMetric object
func (obj *dhcpv6ClientMetric) SetName(value string) Dhcpv6ClientMetric {

	obj.obj.Name = &value
	return obj
}

// Number of DHCPSOLICIT messages sent.
// SolicitsSent returns a uint64
func (obj *dhcpv6ClientMetric) SolicitsSent() uint64 {

	return *obj.obj.SolicitsSent

}

// Number of DHCPSOLICIT messages sent.
// SolicitsSent returns a uint64
func (obj *dhcpv6ClientMetric) HasSolicitsSent() bool {
	return obj.obj.SolicitsSent != nil
}

// Number of DHCPSOLICIT messages sent.
// SetSolicitsSent sets the uint64 value in the Dhcpv6ClientMetric object
func (obj *dhcpv6ClientMetric) SetSolicitsSent(value uint64) Dhcpv6ClientMetric {

	obj.obj.SolicitsSent = &value
	return obj
}

// Number of DHCPADVERTISE messages received.
// AdvertisementsReceived returns a uint64
func (obj *dhcpv6ClientMetric) AdvertisementsReceived() uint64 {

	return *obj.obj.AdvertisementsReceived

}

// Number of DHCPADVERTISE messages received.
// AdvertisementsReceived returns a uint64
func (obj *dhcpv6ClientMetric) HasAdvertisementsReceived() bool {
	return obj.obj.AdvertisementsReceived != nil
}

// Number of DHCPADVERTISE messages received.
// SetAdvertisementsReceived sets the uint64 value in the Dhcpv6ClientMetric object
func (obj *dhcpv6ClientMetric) SetAdvertisementsReceived(value uint64) Dhcpv6ClientMetric {

	obj.obj.AdvertisementsReceived = &value
	return obj
}

// Number of DHCPADVERTISE messages ignored.
// AdvertisementsIgnored returns a uint64
func (obj *dhcpv6ClientMetric) AdvertisementsIgnored() uint64 {

	return *obj.obj.AdvertisementsIgnored

}

// Number of DHCPADVERTISE messages ignored.
// AdvertisementsIgnored returns a uint64
func (obj *dhcpv6ClientMetric) HasAdvertisementsIgnored() bool {
	return obj.obj.AdvertisementsIgnored != nil
}

// Number of DHCPADVERTISE messages ignored.
// SetAdvertisementsIgnored sets the uint64 value in the Dhcpv6ClientMetric object
func (obj *dhcpv6ClientMetric) SetAdvertisementsIgnored(value uint64) Dhcpv6ClientMetric {

	obj.obj.AdvertisementsIgnored = &value
	return obj
}

// Number of DHCPREQUEST messages sent.
// RequestsSent returns a uint64
func (obj *dhcpv6ClientMetric) RequestsSent() uint64 {

	return *obj.obj.RequestsSent

}

// Number of DHCPREQUEST messages sent.
// RequestsSent returns a uint64
func (obj *dhcpv6ClientMetric) HasRequestsSent() bool {
	return obj.obj.RequestsSent != nil
}

// Number of DHCPREQUEST messages sent.
// SetRequestsSent sets the uint64 value in the Dhcpv6ClientMetric object
func (obj *dhcpv6ClientMetric) SetRequestsSent(value uint64) Dhcpv6ClientMetric {

	obj.obj.RequestsSent = &value
	return obj
}

// Number of negative lease DHCPNACK messages received.
// NacksReceived returns a uint64
func (obj *dhcpv6ClientMetric) NacksReceived() uint64 {

	return *obj.obj.NacksReceived

}

// Number of negative lease DHCPNACK messages received.
// NacksReceived returns a uint64
func (obj *dhcpv6ClientMetric) HasNacksReceived() bool {
	return obj.obj.NacksReceived != nil
}

// Number of negative lease DHCPNACK messages received.
// SetNacksReceived sets the uint64 value in the Dhcpv6ClientMetric object
func (obj *dhcpv6ClientMetric) SetNacksReceived(value uint64) Dhcpv6ClientMetric {

	obj.obj.NacksReceived = &value
	return obj
}

// Number of DHCPOFFER messages received.
// RepliesReceived returns a uint64
func (obj *dhcpv6ClientMetric) RepliesReceived() uint64 {

	return *obj.obj.RepliesReceived

}

// Number of DHCPOFFER messages received.
// RepliesReceived returns a uint64
func (obj *dhcpv6ClientMetric) HasRepliesReceived() bool {
	return obj.obj.RepliesReceived != nil
}

// Number of DHCPOFFER messages received.
// SetRepliesReceived sets the uint64 value in the Dhcpv6ClientMetric object
func (obj *dhcpv6ClientMetric) SetRepliesReceived(value uint64) Dhcpv6ClientMetric {

	obj.obj.RepliesReceived = &value
	return obj
}

// Number of DHCP Inform requests sent.
// InformationRequestsSent returns a uint64
func (obj *dhcpv6ClientMetric) InformationRequestsSent() uint64 {

	return *obj.obj.InformationRequestsSent

}

// Number of DHCP Inform requests sent.
// InformationRequestsSent returns a uint64
func (obj *dhcpv6ClientMetric) HasInformationRequestsSent() bool {
	return obj.obj.InformationRequestsSent != nil
}

// Number of DHCP Inform requests sent.
// SetInformationRequestsSent sets the uint64 value in the Dhcpv6ClientMetric object
func (obj *dhcpv6ClientMetric) SetInformationRequestsSent(value uint64) Dhcpv6ClientMetric {

	obj.obj.InformationRequestsSent = &value
	return obj
}

// Number of DHCP renew messages sent.
// RenewsSent returns a uint64
func (obj *dhcpv6ClientMetric) RenewsSent() uint64 {

	return *obj.obj.RenewsSent

}

// Number of DHCP renew messages sent.
// RenewsSent returns a uint64
func (obj *dhcpv6ClientMetric) HasRenewsSent() bool {
	return obj.obj.RenewsSent != nil
}

// Number of DHCP renew messages sent.
// SetRenewsSent sets the uint64 value in the Dhcpv6ClientMetric object
func (obj *dhcpv6ClientMetric) SetRenewsSent(value uint64) Dhcpv6ClientMetric {

	obj.obj.RenewsSent = &value
	return obj
}

// Number of DHCP rebind messages sent.
// RebindsSent returns a uint64
func (obj *dhcpv6ClientMetric) RebindsSent() uint64 {

	return *obj.obj.RebindsSent

}

// Number of DHCP rebind messages sent.
// RebindsSent returns a uint64
func (obj *dhcpv6ClientMetric) HasRebindsSent() bool {
	return obj.obj.RebindsSent != nil
}

// Number of DHCP rebind messages sent.
// SetRebindsSent sets the uint64 value in the Dhcpv6ClientMetric object
func (obj *dhcpv6ClientMetric) SetRebindsSent(value uint64) Dhcpv6ClientMetric {

	obj.obj.RebindsSent = &value
	return obj
}

// Number of DHCP Release messages sent.
// ReleasesSent returns a uint64
func (obj *dhcpv6ClientMetric) ReleasesSent() uint64 {

	return *obj.obj.ReleasesSent

}

// Number of DHCP Release messages sent.
// ReleasesSent returns a uint64
func (obj *dhcpv6ClientMetric) HasReleasesSent() bool {
	return obj.obj.ReleasesSent != nil
}

// Number of DHCP Release messages sent.
// SetReleasesSent sets the uint64 value in the Dhcpv6ClientMetric object
func (obj *dhcpv6ClientMetric) SetReleasesSent(value uint64) Dhcpv6ClientMetric {

	obj.obj.ReleasesSent = &value
	return obj
}

// Number of DHCP Reconfigure messages received.
// ReconfiguresReceived returns a uint64
func (obj *dhcpv6ClientMetric) ReconfiguresReceived() uint64 {

	return *obj.obj.ReconfiguresReceived

}

// Number of DHCP Reconfigure messages received.
// ReconfiguresReceived returns a uint64
func (obj *dhcpv6ClientMetric) HasReconfiguresReceived() bool {
	return obj.obj.ReconfiguresReceived != nil
}

// Number of DHCP Reconfigure messages received.
// SetReconfiguresReceived sets the uint64 value in the Dhcpv6ClientMetric object
func (obj *dhcpv6ClientMetric) SetReconfiguresReceived(value uint64) Dhcpv6ClientMetric {

	obj.obj.ReconfiguresReceived = &value
	return obj
}

// Number of rapid commit DHCPSOLICIT messages sent.
// RapidCommitSolicitsSent returns a uint64
func (obj *dhcpv6ClientMetric) RapidCommitSolicitsSent() uint64 {

	return *obj.obj.RapidCommitSolicitsSent

}

// Number of rapid commit DHCPSOLICIT messages sent.
// RapidCommitSolicitsSent returns a uint64
func (obj *dhcpv6ClientMetric) HasRapidCommitSolicitsSent() bool {
	return obj.obj.RapidCommitSolicitsSent != nil
}

// Number of rapid commit DHCPSOLICIT messages sent.
// SetRapidCommitSolicitsSent sets the uint64 value in the Dhcpv6ClientMetric object
func (obj *dhcpv6ClientMetric) SetRapidCommitSolicitsSent(value uint64) Dhcpv6ClientMetric {

	obj.obj.RapidCommitSolicitsSent = &value
	return obj
}

// Number of rapid commit DHCP Reply messages received.
// RapidCommitRepliesReceived returns a uint64
func (obj *dhcpv6ClientMetric) RapidCommitRepliesReceived() uint64 {

	return *obj.obj.RapidCommitRepliesReceived

}

// Number of rapid commit DHCP Reply messages received.
// RapidCommitRepliesReceived returns a uint64
func (obj *dhcpv6ClientMetric) HasRapidCommitRepliesReceived() bool {
	return obj.obj.RapidCommitRepliesReceived != nil
}

// Number of rapid commit DHCP Reply messages received.
// SetRapidCommitRepliesReceived sets the uint64 value in the Dhcpv6ClientMetric object
func (obj *dhcpv6ClientMetric) SetRapidCommitRepliesReceived(value uint64) Dhcpv6ClientMetric {

	obj.obj.RapidCommitRepliesReceived = &value
	return obj
}

func (obj *dhcpv6ClientMetric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *dhcpv6ClientMetric) setDefault() {

}
