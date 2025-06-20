package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv4ClientMetric *****
type dhcpv4ClientMetric struct {
	validation
	obj          *otg.Dhcpv4ClientMetric
	marshaller   marshalDhcpv4ClientMetric
	unMarshaller unMarshalDhcpv4ClientMetric
}

func NewDhcpv4ClientMetric() Dhcpv4ClientMetric {
	obj := dhcpv4ClientMetric{obj: &otg.Dhcpv4ClientMetric{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv4ClientMetric) msg() *otg.Dhcpv4ClientMetric {
	return obj.obj
}

func (obj *dhcpv4ClientMetric) setMsg(msg *otg.Dhcpv4ClientMetric) Dhcpv4ClientMetric {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv4ClientMetric struct {
	obj *dhcpv4ClientMetric
}

type marshalDhcpv4ClientMetric interface {
	// ToProto marshals Dhcpv4ClientMetric to protobuf object *otg.Dhcpv4ClientMetric
	ToProto() (*otg.Dhcpv4ClientMetric, error)
	// ToPbText marshals Dhcpv4ClientMetric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv4ClientMetric to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv4ClientMetric to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpv4ClientMetric struct {
	obj *dhcpv4ClientMetric
}

type unMarshalDhcpv4ClientMetric interface {
	// FromProto unmarshals Dhcpv4ClientMetric from protobuf object *otg.Dhcpv4ClientMetric
	FromProto(msg *otg.Dhcpv4ClientMetric) (Dhcpv4ClientMetric, error)
	// FromPbText unmarshals Dhcpv4ClientMetric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv4ClientMetric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv4ClientMetric from JSON text
	FromJson(value string) error
}

func (obj *dhcpv4ClientMetric) Marshal() marshalDhcpv4ClientMetric {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv4ClientMetric{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv4ClientMetric) Unmarshal() unMarshalDhcpv4ClientMetric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv4ClientMetric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv4ClientMetric) ToProto() (*otg.Dhcpv4ClientMetric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv4ClientMetric) FromProto(msg *otg.Dhcpv4ClientMetric) (Dhcpv4ClientMetric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv4ClientMetric) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv4ClientMetric) FromPbText(value string) error {
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

func (m *marshaldhcpv4ClientMetric) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv4ClientMetric) FromYaml(value string) error {
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

func (m *marshaldhcpv4ClientMetric) ToJson() (string, error) {
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

func (m *unMarshaldhcpv4ClientMetric) FromJson(value string) error {
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

func (obj *dhcpv4ClientMetric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv4ClientMetric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv4ClientMetric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv4ClientMetric) Clone() (Dhcpv4ClientMetric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv4ClientMetric()
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

// Dhcpv4ClientMetric is dHCPv4 per peer statistics information.
type Dhcpv4ClientMetric interface {
	Validation
	// msg marshals Dhcpv4ClientMetric to protobuf object *otg.Dhcpv4ClientMetric
	// and doesn't set defaults
	msg() *otg.Dhcpv4ClientMetric
	// setMsg unmarshals Dhcpv4ClientMetric from protobuf object *otg.Dhcpv4ClientMetric
	// and doesn't set defaults
	setMsg(*otg.Dhcpv4ClientMetric) Dhcpv4ClientMetric
	// provides marshal interface
	Marshal() marshalDhcpv4ClientMetric
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv4ClientMetric
	// validate validates Dhcpv4ClientMetric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv4ClientMetric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in Dhcpv4ClientMetric.
	Name() string
	// SetName assigns string provided by user to Dhcpv4ClientMetric
	SetName(value string) Dhcpv4ClientMetric
	// HasName checks if Name has been set in Dhcpv4ClientMetric
	HasName() bool
	// DiscoversSent returns uint64, set in Dhcpv4ClientMetric.
	DiscoversSent() uint64
	// SetDiscoversSent assigns uint64 provided by user to Dhcpv4ClientMetric
	SetDiscoversSent(value uint64) Dhcpv4ClientMetric
	// HasDiscoversSent checks if DiscoversSent has been set in Dhcpv4ClientMetric
	HasDiscoversSent() bool
	// OffersReceived returns uint64, set in Dhcpv4ClientMetric.
	OffersReceived() uint64
	// SetOffersReceived assigns uint64 provided by user to Dhcpv4ClientMetric
	SetOffersReceived(value uint64) Dhcpv4ClientMetric
	// HasOffersReceived checks if OffersReceived has been set in Dhcpv4ClientMetric
	HasOffersReceived() bool
	// RequestsSent returns uint64, set in Dhcpv4ClientMetric.
	RequestsSent() uint64
	// SetRequestsSent assigns uint64 provided by user to Dhcpv4ClientMetric
	SetRequestsSent(value uint64) Dhcpv4ClientMetric
	// HasRequestsSent checks if RequestsSent has been set in Dhcpv4ClientMetric
	HasRequestsSent() bool
	// AcksReceived returns uint64, set in Dhcpv4ClientMetric.
	AcksReceived() uint64
	// SetAcksReceived assigns uint64 provided by user to Dhcpv4ClientMetric
	SetAcksReceived(value uint64) Dhcpv4ClientMetric
	// HasAcksReceived checks if AcksReceived has been set in Dhcpv4ClientMetric
	HasAcksReceived() bool
	// NacksReceived returns uint64, set in Dhcpv4ClientMetric.
	NacksReceived() uint64
	// SetNacksReceived assigns uint64 provided by user to Dhcpv4ClientMetric
	SetNacksReceived(value uint64) Dhcpv4ClientMetric
	// HasNacksReceived checks if NacksReceived has been set in Dhcpv4ClientMetric
	HasNacksReceived() bool
	// ReleasesSent returns uint64, set in Dhcpv4ClientMetric.
	ReleasesSent() uint64
	// SetReleasesSent assigns uint64 provided by user to Dhcpv4ClientMetric
	SetReleasesSent(value uint64) Dhcpv4ClientMetric
	// HasReleasesSent checks if ReleasesSent has been set in Dhcpv4ClientMetric
	HasReleasesSent() bool
	// DeclinesSent returns uint64, set in Dhcpv4ClientMetric.
	DeclinesSent() uint64
	// SetDeclinesSent assigns uint64 provided by user to Dhcpv4ClientMetric
	SetDeclinesSent(value uint64) Dhcpv4ClientMetric
	// HasDeclinesSent checks if DeclinesSent has been set in Dhcpv4ClientMetric
	HasDeclinesSent() bool
}

// The name of a configured DHCPv4 client.
// Name returns a string
func (obj *dhcpv4ClientMetric) Name() string {

	return *obj.obj.Name

}

// The name of a configured DHCPv4 client.
// Name returns a string
func (obj *dhcpv4ClientMetric) HasName() bool {
	return obj.obj.Name != nil
}

// The name of a configured DHCPv4 client.
// SetName sets the string value in the Dhcpv4ClientMetric object
func (obj *dhcpv4ClientMetric) SetName(value string) Dhcpv4ClientMetric {

	obj.obj.Name = &value
	return obj
}

// Number of DHCPDISCOVER messages sent.
// DiscoversSent returns a uint64
func (obj *dhcpv4ClientMetric) DiscoversSent() uint64 {

	return *obj.obj.DiscoversSent

}

// Number of DHCPDISCOVER messages sent.
// DiscoversSent returns a uint64
func (obj *dhcpv4ClientMetric) HasDiscoversSent() bool {
	return obj.obj.DiscoversSent != nil
}

// Number of DHCPDISCOVER messages sent.
// SetDiscoversSent sets the uint64 value in the Dhcpv4ClientMetric object
func (obj *dhcpv4ClientMetric) SetDiscoversSent(value uint64) Dhcpv4ClientMetric {

	obj.obj.DiscoversSent = &value
	return obj
}

// Number of DHCPOFFER messages received.
// OffersReceived returns a uint64
func (obj *dhcpv4ClientMetric) OffersReceived() uint64 {

	return *obj.obj.OffersReceived

}

// Number of DHCPOFFER messages received.
// OffersReceived returns a uint64
func (obj *dhcpv4ClientMetric) HasOffersReceived() bool {
	return obj.obj.OffersReceived != nil
}

// Number of DHCPOFFER messages received.
// SetOffersReceived sets the uint64 value in the Dhcpv4ClientMetric object
func (obj *dhcpv4ClientMetric) SetOffersReceived(value uint64) Dhcpv4ClientMetric {

	obj.obj.OffersReceived = &value
	return obj
}

// Number of DHCPREQUEST messages sent.
// RequestsSent returns a uint64
func (obj *dhcpv4ClientMetric) RequestsSent() uint64 {

	return *obj.obj.RequestsSent

}

// Number of DHCPREQUEST messages sent.
// RequestsSent returns a uint64
func (obj *dhcpv4ClientMetric) HasRequestsSent() bool {
	return obj.obj.RequestsSent != nil
}

// Number of DHCPREQUEST messages sent.
// SetRequestsSent sets the uint64 value in the Dhcpv4ClientMetric object
func (obj *dhcpv4ClientMetric) SetRequestsSent(value uint64) Dhcpv4ClientMetric {

	obj.obj.RequestsSent = &value
	return obj
}

// Number of lease DHCPACK messages received.
// AcksReceived returns a uint64
func (obj *dhcpv4ClientMetric) AcksReceived() uint64 {

	return *obj.obj.AcksReceived

}

// Number of lease DHCPACK messages received.
// AcksReceived returns a uint64
func (obj *dhcpv4ClientMetric) HasAcksReceived() bool {
	return obj.obj.AcksReceived != nil
}

// Number of lease DHCPACK messages received.
// SetAcksReceived sets the uint64 value in the Dhcpv4ClientMetric object
func (obj *dhcpv4ClientMetric) SetAcksReceived(value uint64) Dhcpv4ClientMetric {

	obj.obj.AcksReceived = &value
	return obj
}

// Number of negative lease DHCPNACK messages received.
// NacksReceived returns a uint64
func (obj *dhcpv4ClientMetric) NacksReceived() uint64 {

	return *obj.obj.NacksReceived

}

// Number of negative lease DHCPNACK messages received.
// NacksReceived returns a uint64
func (obj *dhcpv4ClientMetric) HasNacksReceived() bool {
	return obj.obj.NacksReceived != nil
}

// Number of negative lease DHCPNACK messages received.
// SetNacksReceived sets the uint64 value in the Dhcpv4ClientMetric object
func (obj *dhcpv4ClientMetric) SetNacksReceived(value uint64) Dhcpv4ClientMetric {

	obj.obj.NacksReceived = &value
	return obj
}

// Number of DHCPRELEASE messages sent.
// ReleasesSent returns a uint64
func (obj *dhcpv4ClientMetric) ReleasesSent() uint64 {

	return *obj.obj.ReleasesSent

}

// Number of DHCPRELEASE messages sent.
// ReleasesSent returns a uint64
func (obj *dhcpv4ClientMetric) HasReleasesSent() bool {
	return obj.obj.ReleasesSent != nil
}

// Number of DHCPRELEASE messages sent.
// SetReleasesSent sets the uint64 value in the Dhcpv4ClientMetric object
func (obj *dhcpv4ClientMetric) SetReleasesSent(value uint64) Dhcpv4ClientMetric {

	obj.obj.ReleasesSent = &value
	return obj
}

// Number of DHCPDECLINE messages sent.
// DeclinesSent returns a uint64
func (obj *dhcpv4ClientMetric) DeclinesSent() uint64 {

	return *obj.obj.DeclinesSent

}

// Number of DHCPDECLINE messages sent.
// DeclinesSent returns a uint64
func (obj *dhcpv4ClientMetric) HasDeclinesSent() bool {
	return obj.obj.DeclinesSent != nil
}

// Number of DHCPDECLINE messages sent.
// SetDeclinesSent sets the uint64 value in the Dhcpv4ClientMetric object
func (obj *dhcpv4ClientMetric) SetDeclinesSent(value uint64) Dhcpv4ClientMetric {

	obj.obj.DeclinesSent = &value
	return obj
}

func (obj *dhcpv4ClientMetric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *dhcpv4ClientMetric) setDefault() {

}
