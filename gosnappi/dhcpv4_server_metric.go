package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv4ServerMetric *****
type dhcpv4ServerMetric struct {
	validation
	obj          *otg.Dhcpv4ServerMetric
	marshaller   marshalDhcpv4ServerMetric
	unMarshaller unMarshalDhcpv4ServerMetric
}

func NewDhcpv4ServerMetric() Dhcpv4ServerMetric {
	obj := dhcpv4ServerMetric{obj: &otg.Dhcpv4ServerMetric{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv4ServerMetric) msg() *otg.Dhcpv4ServerMetric {
	return obj.obj
}

func (obj *dhcpv4ServerMetric) setMsg(msg *otg.Dhcpv4ServerMetric) Dhcpv4ServerMetric {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv4ServerMetric struct {
	obj *dhcpv4ServerMetric
}

type marshalDhcpv4ServerMetric interface {
	// ToProto marshals Dhcpv4ServerMetric to protobuf object *otg.Dhcpv4ServerMetric
	ToProto() (*otg.Dhcpv4ServerMetric, error)
	// ToPbText marshals Dhcpv4ServerMetric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv4ServerMetric to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv4ServerMetric to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Dhcpv4ServerMetric to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldhcpv4ServerMetric struct {
	obj *dhcpv4ServerMetric
}

type unMarshalDhcpv4ServerMetric interface {
	// FromProto unmarshals Dhcpv4ServerMetric from protobuf object *otg.Dhcpv4ServerMetric
	FromProto(msg *otg.Dhcpv4ServerMetric) (Dhcpv4ServerMetric, error)
	// FromPbText unmarshals Dhcpv4ServerMetric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv4ServerMetric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv4ServerMetric from JSON text
	FromJson(value string) error
}

func (obj *dhcpv4ServerMetric) Marshal() marshalDhcpv4ServerMetric {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv4ServerMetric{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv4ServerMetric) Unmarshal() unMarshalDhcpv4ServerMetric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv4ServerMetric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv4ServerMetric) ToProto() (*otg.Dhcpv4ServerMetric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv4ServerMetric) FromProto(msg *otg.Dhcpv4ServerMetric) (Dhcpv4ServerMetric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv4ServerMetric) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv4ServerMetric) FromPbText(value string) error {
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

func (m *marshaldhcpv4ServerMetric) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv4ServerMetric) FromYaml(value string) error {
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

func (m *marshaldhcpv4ServerMetric) ToJsonRaw() (string, error) {
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

func (m *marshaldhcpv4ServerMetric) ToJson() (string, error) {
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

func (m *unMarshaldhcpv4ServerMetric) FromJson(value string) error {
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

func (obj *dhcpv4ServerMetric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv4ServerMetric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv4ServerMetric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv4ServerMetric) Clone() (Dhcpv4ServerMetric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv4ServerMetric()
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

// Dhcpv4ServerMetric is dHCPv4 per peer statistics information.
type Dhcpv4ServerMetric interface {
	Validation
	// msg marshals Dhcpv4ServerMetric to protobuf object *otg.Dhcpv4ServerMetric
	// and doesn't set defaults
	msg() *otg.Dhcpv4ServerMetric
	// setMsg unmarshals Dhcpv4ServerMetric from protobuf object *otg.Dhcpv4ServerMetric
	// and doesn't set defaults
	setMsg(*otg.Dhcpv4ServerMetric) Dhcpv4ServerMetric
	// provides marshal interface
	Marshal() marshalDhcpv4ServerMetric
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv4ServerMetric
	// validate validates Dhcpv4ServerMetric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv4ServerMetric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in Dhcpv4ServerMetric.
	Name() string
	// SetName assigns string provided by user to Dhcpv4ServerMetric
	SetName(value string) Dhcpv4ServerMetric
	// HasName checks if Name has been set in Dhcpv4ServerMetric
	HasName() bool
	// DiscoversReceived returns uint64, set in Dhcpv4ServerMetric.
	DiscoversReceived() uint64
	// SetDiscoversReceived assigns uint64 provided by user to Dhcpv4ServerMetric
	SetDiscoversReceived(value uint64) Dhcpv4ServerMetric
	// HasDiscoversReceived checks if DiscoversReceived has been set in Dhcpv4ServerMetric
	HasDiscoversReceived() bool
	// OffersSent returns uint64, set in Dhcpv4ServerMetric.
	OffersSent() uint64
	// SetOffersSent assigns uint64 provided by user to Dhcpv4ServerMetric
	SetOffersSent(value uint64) Dhcpv4ServerMetric
	// HasOffersSent checks if OffersSent has been set in Dhcpv4ServerMetric
	HasOffersSent() bool
	// RequestsReceived returns uint64, set in Dhcpv4ServerMetric.
	RequestsReceived() uint64
	// SetRequestsReceived assigns uint64 provided by user to Dhcpv4ServerMetric
	SetRequestsReceived(value uint64) Dhcpv4ServerMetric
	// HasRequestsReceived checks if RequestsReceived has been set in Dhcpv4ServerMetric
	HasRequestsReceived() bool
	// AcksSent returns uint64, set in Dhcpv4ServerMetric.
	AcksSent() uint64
	// SetAcksSent assigns uint64 provided by user to Dhcpv4ServerMetric
	SetAcksSent(value uint64) Dhcpv4ServerMetric
	// HasAcksSent checks if AcksSent has been set in Dhcpv4ServerMetric
	HasAcksSent() bool
	// NacksSent returns uint64, set in Dhcpv4ServerMetric.
	NacksSent() uint64
	// SetNacksSent assigns uint64 provided by user to Dhcpv4ServerMetric
	SetNacksSent(value uint64) Dhcpv4ServerMetric
	// HasNacksSent checks if NacksSent has been set in Dhcpv4ServerMetric
	HasNacksSent() bool
	// ReleasesReceived returns uint64, set in Dhcpv4ServerMetric.
	ReleasesReceived() uint64
	// SetReleasesReceived assigns uint64 provided by user to Dhcpv4ServerMetric
	SetReleasesReceived(value uint64) Dhcpv4ServerMetric
	// HasReleasesReceived checks if ReleasesReceived has been set in Dhcpv4ServerMetric
	HasReleasesReceived() bool
	// DeclinesReceived returns uint64, set in Dhcpv4ServerMetric.
	DeclinesReceived() uint64
	// SetDeclinesReceived assigns uint64 provided by user to Dhcpv4ServerMetric
	SetDeclinesReceived(value uint64) Dhcpv4ServerMetric
	// HasDeclinesReceived checks if DeclinesReceived has been set in Dhcpv4ServerMetric
	HasDeclinesReceived() bool
}

// The name of a configured DHCPv4 Server.
// Name returns a string
func (obj *dhcpv4ServerMetric) Name() string {

	return *obj.obj.Name

}

// The name of a configured DHCPv4 Server.
// Name returns a string
func (obj *dhcpv4ServerMetric) HasName() bool {
	return obj.obj.Name != nil
}

// The name of a configured DHCPv4 Server.
// SetName sets the string value in the Dhcpv4ServerMetric object
func (obj *dhcpv4ServerMetric) SetName(value string) Dhcpv4ServerMetric {

	obj.obj.Name = &value
	return obj
}

// Number of DHCPDISCOVER messages received.
// DiscoversReceived returns a uint64
func (obj *dhcpv4ServerMetric) DiscoversReceived() uint64 {

	return *obj.obj.DiscoversReceived

}

// Number of DHCPDISCOVER messages received.
// DiscoversReceived returns a uint64
func (obj *dhcpv4ServerMetric) HasDiscoversReceived() bool {
	return obj.obj.DiscoversReceived != nil
}

// Number of DHCPDISCOVER messages received.
// SetDiscoversReceived sets the uint64 value in the Dhcpv4ServerMetric object
func (obj *dhcpv4ServerMetric) SetDiscoversReceived(value uint64) Dhcpv4ServerMetric {

	obj.obj.DiscoversReceived = &value
	return obj
}

// Number of DHCPOFFER messages sent.
// OffersSent returns a uint64
func (obj *dhcpv4ServerMetric) OffersSent() uint64 {

	return *obj.obj.OffersSent

}

// Number of DHCPOFFER messages sent.
// OffersSent returns a uint64
func (obj *dhcpv4ServerMetric) HasOffersSent() bool {
	return obj.obj.OffersSent != nil
}

// Number of DHCPOFFER messages sent.
// SetOffersSent sets the uint64 value in the Dhcpv4ServerMetric object
func (obj *dhcpv4ServerMetric) SetOffersSent(value uint64) Dhcpv4ServerMetric {

	obj.obj.OffersSent = &value
	return obj
}

// Number of DHCPOFFER messages received.
// RequestsReceived returns a uint64
func (obj *dhcpv4ServerMetric) RequestsReceived() uint64 {

	return *obj.obj.RequestsReceived

}

// Number of DHCPOFFER messages received.
// RequestsReceived returns a uint64
func (obj *dhcpv4ServerMetric) HasRequestsReceived() bool {
	return obj.obj.RequestsReceived != nil
}

// Number of DHCPOFFER messages received.
// SetRequestsReceived sets the uint64 value in the Dhcpv4ServerMetric object
func (obj *dhcpv4ServerMetric) SetRequestsReceived(value uint64) Dhcpv4ServerMetric {

	obj.obj.RequestsReceived = &value
	return obj
}

// Number of lease DHCPACK messages sent.
// AcksSent returns a uint64
func (obj *dhcpv4ServerMetric) AcksSent() uint64 {

	return *obj.obj.AcksSent

}

// Number of lease DHCPACK messages sent.
// AcksSent returns a uint64
func (obj *dhcpv4ServerMetric) HasAcksSent() bool {
	return obj.obj.AcksSent != nil
}

// Number of lease DHCPACK messages sent.
// SetAcksSent sets the uint64 value in the Dhcpv4ServerMetric object
func (obj *dhcpv4ServerMetric) SetAcksSent(value uint64) Dhcpv4ServerMetric {

	obj.obj.AcksSent = &value
	return obj
}

// Number of negative lease DHCPNACK messages sent.
// NacksSent returns a uint64
func (obj *dhcpv4ServerMetric) NacksSent() uint64 {

	return *obj.obj.NacksSent

}

// Number of negative lease DHCPNACK messages sent.
// NacksSent returns a uint64
func (obj *dhcpv4ServerMetric) HasNacksSent() bool {
	return obj.obj.NacksSent != nil
}

// Number of negative lease DHCPNACK messages sent.
// SetNacksSent sets the uint64 value in the Dhcpv4ServerMetric object
func (obj *dhcpv4ServerMetric) SetNacksSent(value uint64) Dhcpv4ServerMetric {

	obj.obj.NacksSent = &value
	return obj
}

// Number of DHCPRELEASE messages received.
// ReleasesReceived returns a uint64
func (obj *dhcpv4ServerMetric) ReleasesReceived() uint64 {

	return *obj.obj.ReleasesReceived

}

// Number of DHCPRELEASE messages received.
// ReleasesReceived returns a uint64
func (obj *dhcpv4ServerMetric) HasReleasesReceived() bool {
	return obj.obj.ReleasesReceived != nil
}

// Number of DHCPRELEASE messages received.
// SetReleasesReceived sets the uint64 value in the Dhcpv4ServerMetric object
func (obj *dhcpv4ServerMetric) SetReleasesReceived(value uint64) Dhcpv4ServerMetric {

	obj.obj.ReleasesReceived = &value
	return obj
}

// Number of DHCPDECLINE messages received.
// DeclinesReceived returns a uint64
func (obj *dhcpv4ServerMetric) DeclinesReceived() uint64 {

	return *obj.obj.DeclinesReceived

}

// Number of DHCPDECLINE messages received.
// DeclinesReceived returns a uint64
func (obj *dhcpv4ServerMetric) HasDeclinesReceived() bool {
	return obj.obj.DeclinesReceived != nil
}

// Number of DHCPDECLINE messages received.
// SetDeclinesReceived sets the uint64 value in the Dhcpv4ServerMetric object
func (obj *dhcpv4ServerMetric) SetDeclinesReceived(value uint64) Dhcpv4ServerMetric {

	obj.obj.DeclinesReceived = &value
	return obj
}

func (obj *dhcpv4ServerMetric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *dhcpv4ServerMetric) setDefault() {

}
