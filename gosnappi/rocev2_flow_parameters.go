package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2FlowParameters *****
type rocev2FlowParameters struct {
	validation
	obj          *otg.Rocev2FlowParameters
	marshaller   marshalRocev2FlowParameters
	unMarshaller unMarshalRocev2FlowParameters
}

func NewRocev2FlowParameters() Rocev2FlowParameters {
	obj := rocev2FlowParameters{obj: &otg.Rocev2FlowParameters{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2FlowParameters) msg() *otg.Rocev2FlowParameters {
	return obj.obj
}

func (obj *rocev2FlowParameters) setMsg(msg *otg.Rocev2FlowParameters) Rocev2FlowParameters {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2FlowParameters struct {
	obj *rocev2FlowParameters
}

type marshalRocev2FlowParameters interface {
	// ToProto marshals Rocev2FlowParameters to protobuf object *otg.Rocev2FlowParameters
	ToProto() (*otg.Rocev2FlowParameters, error)
	// ToPbText marshals Rocev2FlowParameters to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2FlowParameters to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2FlowParameters to JSON text
	ToJson() (string, error)
}

type unMarshalrocev2FlowParameters struct {
	obj *rocev2FlowParameters
}

type unMarshalRocev2FlowParameters interface {
	// FromProto unmarshals Rocev2FlowParameters from protobuf object *otg.Rocev2FlowParameters
	FromProto(msg *otg.Rocev2FlowParameters) (Rocev2FlowParameters, error)
	// FromPbText unmarshals Rocev2FlowParameters from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2FlowParameters from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2FlowParameters from JSON text
	FromJson(value string) error
}

func (obj *rocev2FlowParameters) Marshal() marshalRocev2FlowParameters {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2FlowParameters{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2FlowParameters) Unmarshal() unMarshalRocev2FlowParameters {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2FlowParameters{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2FlowParameters) ToProto() (*otg.Rocev2FlowParameters, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2FlowParameters) FromProto(msg *otg.Rocev2FlowParameters) (Rocev2FlowParameters, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2FlowParameters) ToPbText() (string, error) {
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

func (m *unMarshalrocev2FlowParameters) FromPbText(value string) error {
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

func (m *marshalrocev2FlowParameters) ToYaml() (string, error) {
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

func (m *unMarshalrocev2FlowParameters) FromYaml(value string) error {
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

func (m *marshalrocev2FlowParameters) ToJson() (string, error) {
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

func (m *unMarshalrocev2FlowParameters) FromJson(value string) error {
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

func (obj *rocev2FlowParameters) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2FlowParameters) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2FlowParameters) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2FlowParameters) Clone() (Rocev2FlowParameters, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2FlowParameters()
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

// Rocev2FlowParameters is ****Description here****
type Rocev2FlowParameters interface {
	Validation
	// msg marshals Rocev2FlowParameters to protobuf object *otg.Rocev2FlowParameters
	// and doesn't set defaults
	msg() *otg.Rocev2FlowParameters
	// setMsg unmarshals Rocev2FlowParameters from protobuf object *otg.Rocev2FlowParameters
	// and doesn't set defaults
	setMsg(*otg.Rocev2FlowParameters) Rocev2FlowParameters
	// provides marshal interface
	Marshal() marshalRocev2FlowParameters
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2FlowParameters
	// validate validates Rocev2FlowParameters
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2FlowParameters, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SourceQpNumber returns uint32, set in Rocev2FlowParameters.
	SourceQpNumber() uint32
	// SetSourceQpNumber assigns uint32 provided by user to Rocev2FlowParameters
	SetSourceQpNumber(value uint32) Rocev2FlowParameters
	// HasSourceQpNumber checks if SourceQpNumber has been set in Rocev2FlowParameters
	HasSourceQpNumber() bool
	// Dscp returns uint32, set in Rocev2FlowParameters.
	Dscp() uint32
	// SetDscp assigns uint32 provided by user to Rocev2FlowParameters
	SetDscp(value uint32) Rocev2FlowParameters
	// HasDscp checks if Dscp has been set in Rocev2FlowParameters
	HasDscp() bool
	// Ecn returns uint32, set in Rocev2FlowParameters.
	Ecn() uint32
	// SetEcn assigns uint32 provided by user to Rocev2FlowParameters
	SetEcn(value uint32) Rocev2FlowParameters
	// HasEcn checks if Ecn has been set in Rocev2FlowParameters
	HasEcn() bool
	// UdpSourcePort returns uint32, set in Rocev2FlowParameters.
	UdpSourcePort() uint32
	// SetUdpSourcePort assigns uint32 provided by user to Rocev2FlowParameters
	SetUdpSourcePort(value uint32) Rocev2FlowParameters
	// HasUdpSourcePort checks if UdpSourcePort has been set in Rocev2FlowParameters
	HasUdpSourcePort() bool
	// InitialPsn returns uint64, set in Rocev2FlowParameters.
	InitialPsn() uint64
	// SetInitialPsn assigns uint64 provided by user to Rocev2FlowParameters
	SetInitialPsn(value uint64) Rocev2FlowParameters
	// HasInitialPsn checks if InitialPsn has been set in Rocev2FlowParameters
	HasInitialPsn() bool
	// VirtualAddress returns string, set in Rocev2FlowParameters.
	VirtualAddress() string
	// SetVirtualAddress assigns string provided by user to Rocev2FlowParameters
	SetVirtualAddress(value string) Rocev2FlowParameters
	// HasVirtualAddress checks if VirtualAddress has been set in Rocev2FlowParameters
	HasVirtualAddress() bool
	// RemoteKey returns string, set in Rocev2FlowParameters.
	RemoteKey() string
	// SetRemoteKey assigns string provided by user to Rocev2FlowParameters
	SetRemoteKey(value string) Rocev2FlowParameters
	// HasRemoteKey checks if RemoteKey has been set in Rocev2FlowParameters
	HasRemoteKey() bool
}

// Configure Source QP.
// SourceQpNumber returns a uint32
func (obj *rocev2FlowParameters) SourceQpNumber() uint32 {

	return *obj.obj.SourceQpNumber

}

// Configure Source QP.
// SourceQpNumber returns a uint32
func (obj *rocev2FlowParameters) HasSourceQpNumber() bool {
	return obj.obj.SourceQpNumber != nil
}

// Configure Source QP.
// SetSourceQpNumber sets the uint32 value in the Rocev2FlowParameters object
func (obj *rocev2FlowParameters) SetSourceQpNumber(value uint32) Rocev2FlowParameters {

	obj.obj.SourceQpNumber = &value
	return obj
}

// DSCP value for this flow
// Dscp returns a uint32
func (obj *rocev2FlowParameters) Dscp() uint32 {

	return *obj.obj.Dscp

}

// DSCP value for this flow
// Dscp returns a uint32
func (obj *rocev2FlowParameters) HasDscp() bool {
	return obj.obj.Dscp != nil
}

// DSCP value for this flow
// SetDscp sets the uint32 value in the Rocev2FlowParameters object
func (obj *rocev2FlowParameters) SetDscp(value uint32) Rocev2FlowParameters {

	obj.obj.Dscp = &value
	return obj
}

// This field allows to configure bits of the Traffic Class field in the IPv4 or IPv6 header to encode four different code points.
// Ecn returns a uint32
func (obj *rocev2FlowParameters) Ecn() uint32 {

	return *obj.obj.Ecn

}

// This field allows to configure bits of the Traffic Class field in the IPv4 or IPv6 header to encode four different code points.
// Ecn returns a uint32
func (obj *rocev2FlowParameters) HasEcn() bool {
	return obj.obj.Ecn != nil
}

// This field allows to configure bits of the Traffic Class field in the IPv4 or IPv6 header to encode four different code points.
// SetEcn sets the uint32 value in the Rocev2FlowParameters object
func (obj *rocev2FlowParameters) SetEcn(value uint32) Rocev2FlowParameters {

	obj.obj.Ecn = &value
	return obj
}

// UDP source port number for this flow.
// UdpSourcePort returns a uint32
func (obj *rocev2FlowParameters) UdpSourcePort() uint32 {

	return *obj.obj.UdpSourcePort

}

// UDP source port number for this flow.
// UdpSourcePort returns a uint32
func (obj *rocev2FlowParameters) HasUdpSourcePort() bool {
	return obj.obj.UdpSourcePort != nil
}

// UDP source port number for this flow.
// SetUdpSourcePort sets the uint32 value in the Rocev2FlowParameters object
func (obj *rocev2FlowParameters) SetUdpSourcePort(value uint32) Rocev2FlowParameters {

	obj.obj.UdpSourcePort = &value
	return obj
}

// Initial PSN.
// InitialPsn returns a uint64
func (obj *rocev2FlowParameters) InitialPsn() uint64 {

	return *obj.obj.InitialPsn

}

// Initial PSN.
// InitialPsn returns a uint64
func (obj *rocev2FlowParameters) HasInitialPsn() bool {
	return obj.obj.InitialPsn != nil
}

// Initial PSN.
// SetInitialPsn sets the uint64 value in the Rocev2FlowParameters object
func (obj *rocev2FlowParameters) SetInitialPsn(value uint64) Rocev2FlowParameters {

	obj.obj.InitialPsn = &value
	return obj
}

// Virutal Address.
// VirtualAddress returns a string
func (obj *rocev2FlowParameters) VirtualAddress() string {

	return *obj.obj.VirtualAddress

}

// Virutal Address.
// VirtualAddress returns a string
func (obj *rocev2FlowParameters) HasVirtualAddress() bool {
	return obj.obj.VirtualAddress != nil
}

// Virutal Address.
// SetVirtualAddress sets the string value in the Rocev2FlowParameters object
func (obj *rocev2FlowParameters) SetVirtualAddress(value string) Rocev2FlowParameters {

	obj.obj.VirtualAddress = &value
	return obj
}

// Remote Key.
// RemoteKey returns a string
func (obj *rocev2FlowParameters) RemoteKey() string {

	return *obj.obj.RemoteKey

}

// Remote Key.
// RemoteKey returns a string
func (obj *rocev2FlowParameters) HasRemoteKey() bool {
	return obj.obj.RemoteKey != nil
}

// Remote Key.
// SetRemoteKey sets the string value in the Rocev2FlowParameters object
func (obj *rocev2FlowParameters) SetRemoteKey(value string) Rocev2FlowParameters {

	obj.obj.RemoteKey = &value
	return obj
}

func (obj *rocev2FlowParameters) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SourceQpNumber != nil {

		if *obj.obj.SourceQpNumber < 2 || *obj.obj.SourceQpNumber > 33554431 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("2 <= Rocev2FlowParameters.SourceQpNumber <= 33554431 but Got %d", *obj.obj.SourceQpNumber))
		}

	}

	if obj.obj.Dscp != nil {

		if *obj.obj.Dscp > 63 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2FlowParameters.Dscp <= 63 but Got %d", *obj.obj.Dscp))
		}

	}

	if obj.obj.Ecn != nil {

		if *obj.obj.Ecn > 3 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2FlowParameters.Ecn <= 3 but Got %d", *obj.obj.Ecn))
		}

	}

	if obj.obj.UdpSourcePort != nil {

		if *obj.obj.UdpSourcePort > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2FlowParameters.UdpSourcePort <= 65535 but Got %d", *obj.obj.UdpSourcePort))
		}

	}

	if obj.obj.InitialPsn != nil {

		if *obj.obj.InitialPsn > 33554431 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2FlowParameters.InitialPsn <= 33554431 but Got %d", *obj.obj.InitialPsn))
		}

	}

	if obj.obj.VirtualAddress != nil {

		if len(*obj.obj.VirtualAddress) > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"None <= length of Rocev2FlowParameters.VirtualAddress <= 16 but Got %d",
					len(*obj.obj.VirtualAddress)))
		}

	}

	if obj.obj.RemoteKey != nil {

		if len(*obj.obj.RemoteKey) > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"None <= length of Rocev2FlowParameters.RemoteKey <= 8 but Got %d",
					len(*obj.obj.RemoteKey)))
		}

	}

}

func (obj *rocev2FlowParameters) setDefault() {
	if obj.obj.SourceQpNumber == nil {
		obj.SetSourceQpNumber(2)
	}
	if obj.obj.Dscp == nil {
		obj.SetDscp(24)
	}
	if obj.obj.Ecn == nil {
		obj.SetEcn(1)
	}
	if obj.obj.UdpSourcePort == nil {
		obj.SetUdpSourcePort(49152)
	}
	if obj.obj.InitialPsn == nil {
		obj.SetInitialPsn(0)
	}
	if obj.obj.VirtualAddress == nil {
		obj.SetVirtualAddress("0000000000000000")
	}
	if obj.obj.RemoteKey == nil {
		obj.SetRemoteKey("00000000")
	}

}
