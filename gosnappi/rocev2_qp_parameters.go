package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2QPParameters *****
type rocev2QPParameters struct {
	validation
	obj          *otg.Rocev2QPParameters
	marshaller   marshalRocev2QPParameters
	unMarshaller unMarshalRocev2QPParameters
}

func NewRocev2QPParameters() Rocev2QPParameters {
	obj := rocev2QPParameters{obj: &otg.Rocev2QPParameters{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2QPParameters) msg() *otg.Rocev2QPParameters {
	return obj.obj
}

func (obj *rocev2QPParameters) setMsg(msg *otg.Rocev2QPParameters) Rocev2QPParameters {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2QPParameters struct {
	obj *rocev2QPParameters
}

type marshalRocev2QPParameters interface {
	// ToProto marshals Rocev2QPParameters to protobuf object *otg.Rocev2QPParameters
	ToProto() (*otg.Rocev2QPParameters, error)
	// ToPbText marshals Rocev2QPParameters to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2QPParameters to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2QPParameters to JSON text
	ToJson() (string, error)
}

type unMarshalrocev2QPParameters struct {
	obj *rocev2QPParameters
}

type unMarshalRocev2QPParameters interface {
	// FromProto unmarshals Rocev2QPParameters from protobuf object *otg.Rocev2QPParameters
	FromProto(msg *otg.Rocev2QPParameters) (Rocev2QPParameters, error)
	// FromPbText unmarshals Rocev2QPParameters from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2QPParameters from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2QPParameters from JSON text
	FromJson(value string) error
}

func (obj *rocev2QPParameters) Marshal() marshalRocev2QPParameters {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2QPParameters{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2QPParameters) Unmarshal() unMarshalRocev2QPParameters {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2QPParameters{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2QPParameters) ToProto() (*otg.Rocev2QPParameters, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2QPParameters) FromProto(msg *otg.Rocev2QPParameters) (Rocev2QPParameters, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2QPParameters) ToPbText() (string, error) {
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

func (m *unMarshalrocev2QPParameters) FromPbText(value string) error {
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

func (m *marshalrocev2QPParameters) ToYaml() (string, error) {
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

func (m *unMarshalrocev2QPParameters) FromYaml(value string) error {
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

func (m *marshalrocev2QPParameters) ToJson() (string, error) {
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

func (m *unMarshalrocev2QPParameters) FromJson(value string) error {
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

func (obj *rocev2QPParameters) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2QPParameters) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2QPParameters) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2QPParameters) Clone() (Rocev2QPParameters, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2QPParameters()
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

// Rocev2QPParameters is defines the parameters for configuring a RoCEv2 QP.
type Rocev2QPParameters interface {
	Validation
	// msg marshals Rocev2QPParameters to protobuf object *otg.Rocev2QPParameters
	// and doesn't set defaults
	msg() *otg.Rocev2QPParameters
	// setMsg unmarshals Rocev2QPParameters from protobuf object *otg.Rocev2QPParameters
	// and doesn't set defaults
	setMsg(*otg.Rocev2QPParameters) Rocev2QPParameters
	// provides marshal interface
	Marshal() marshalRocev2QPParameters
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2QPParameters
	// validate validates Rocev2QPParameters
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2QPParameters, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SourceQpNumber returns uint32, set in Rocev2QPParameters.
	SourceQpNumber() uint32
	// SetSourceQpNumber assigns uint32 provided by user to Rocev2QPParameters
	SetSourceQpNumber(value uint32) Rocev2QPParameters
	// HasSourceQpNumber checks if SourceQpNumber has been set in Rocev2QPParameters
	HasSourceQpNumber() bool
	// Dscp returns uint32, set in Rocev2QPParameters.
	Dscp() uint32
	// SetDscp assigns uint32 provided by user to Rocev2QPParameters
	SetDscp(value uint32) Rocev2QPParameters
	// HasDscp checks if Dscp has been set in Rocev2QPParameters
	HasDscp() bool
	// Ecn returns Rocev2QPParametersEcnEnum, set in Rocev2QPParameters
	Ecn() Rocev2QPParametersEcnEnum
	// SetEcn assigns Rocev2QPParametersEcnEnum provided by user to Rocev2QPParameters
	SetEcn(value Rocev2QPParametersEcnEnum) Rocev2QPParameters
	// HasEcn checks if Ecn has been set in Rocev2QPParameters
	HasEcn() bool
	// UdpSourcePort returns uint32, set in Rocev2QPParameters.
	UdpSourcePort() uint32
	// SetUdpSourcePort assigns uint32 provided by user to Rocev2QPParameters
	SetUdpSourcePort(value uint32) Rocev2QPParameters
	// HasUdpSourcePort checks if UdpSourcePort has been set in Rocev2QPParameters
	HasUdpSourcePort() bool
	// InitialPsn returns uint64, set in Rocev2QPParameters.
	InitialPsn() uint64
	// SetInitialPsn assigns uint64 provided by user to Rocev2QPParameters
	SetInitialPsn(value uint64) Rocev2QPParameters
	// HasInitialPsn checks if InitialPsn has been set in Rocev2QPParameters
	HasInitialPsn() bool
	// VirtualAddress returns string, set in Rocev2QPParameters.
	VirtualAddress() string
	// SetVirtualAddress assigns string provided by user to Rocev2QPParameters
	SetVirtualAddress(value string) Rocev2QPParameters
	// HasVirtualAddress checks if VirtualAddress has been set in Rocev2QPParameters
	HasVirtualAddress() bool
	// RemoteKey returns string, set in Rocev2QPParameters.
	RemoteKey() string
	// SetRemoteKey assigns string provided by user to Rocev2QPParameters
	SetRemoteKey(value string) Rocev2QPParameters
	// HasRemoteKey checks if RemoteKey has been set in Rocev2QPParameters
	HasRemoteKey() bool
}

// Configure Source QP number which initiates the RDMA operation.
// SourceQpNumber returns a uint32
func (obj *rocev2QPParameters) SourceQpNumber() uint32 {

	return *obj.obj.SourceQpNumber

}

// Configure Source QP number which initiates the RDMA operation.
// SourceQpNumber returns a uint32
func (obj *rocev2QPParameters) HasSourceQpNumber() bool {
	return obj.obj.SourceQpNumber != nil
}

// Configure Source QP number which initiates the RDMA operation.
// SetSourceQpNumber sets the uint32 value in the Rocev2QPParameters object
func (obj *rocev2QPParameters) SetSourceQpNumber(value uint32) Rocev2QPParameters {

	obj.obj.SourceQpNumber = &value
	return obj
}

// DSCP value for the RDMA data packets.
// Dscp returns a uint32
func (obj *rocev2QPParameters) Dscp() uint32 {

	return *obj.obj.Dscp

}

// DSCP value for the RDMA data packets.
// Dscp returns a uint32
func (obj *rocev2QPParameters) HasDscp() bool {
	return obj.obj.Dscp != nil
}

// DSCP value for the RDMA data packets.
// SetDscp sets the uint32 value in the Rocev2QPParameters object
func (obj *rocev2QPParameters) SetDscp(value uint32) Rocev2QPParameters {

	obj.obj.Dscp = &value
	return obj
}

type Rocev2QPParametersEcnEnum string

// Enum of Ecn on Rocev2QPParameters
var Rocev2QPParametersEcn = struct {
	NON_ECT Rocev2QPParametersEcnEnum
	ECT_1   Rocev2QPParametersEcnEnum
	ECT_0   Rocev2QPParametersEcnEnum
	CE      Rocev2QPParametersEcnEnum
}{
	NON_ECT: Rocev2QPParametersEcnEnum("non_ect"),
	ECT_1:   Rocev2QPParametersEcnEnum("ect_1"),
	ECT_0:   Rocev2QPParametersEcnEnum("ect_0"),
	CE:      Rocev2QPParametersEcnEnum("ce"),
}

func (obj *rocev2QPParameters) Ecn() Rocev2QPParametersEcnEnum {
	return Rocev2QPParametersEcnEnum(obj.obj.Ecn.Enum().String())
}

// This field allows to configure bits of the Traffic Class field in the IPv4 or IPv6 header to encode four different code points. Those are non_ect, ect_1, ect_0 and ce. non_ect quivalent is 00, ect_1 represent 01, ect_0 represent 10 and ce means 11.
// Ecn returns a string
func (obj *rocev2QPParameters) HasEcn() bool {
	return obj.obj.Ecn != nil
}

func (obj *rocev2QPParameters) SetEcn(value Rocev2QPParametersEcnEnum) Rocev2QPParameters {
	intValue, ok := otg.Rocev2QPParameters_Ecn_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Rocev2QPParametersEcnEnum", string(value)))
		return obj
	}
	enumValue := otg.Rocev2QPParameters_Ecn_Enum(intValue)
	obj.obj.Ecn = &enumValue

	return obj
}

// UDP source port number for this QP.
// UdpSourcePort returns a uint32
func (obj *rocev2QPParameters) UdpSourcePort() uint32 {

	return *obj.obj.UdpSourcePort

}

// UDP source port number for this QP.
// UdpSourcePort returns a uint32
func (obj *rocev2QPParameters) HasUdpSourcePort() bool {
	return obj.obj.UdpSourcePort != nil
}

// UDP source port number for this QP.
// SetUdpSourcePort sets the uint32 value in the Rocev2QPParameters object
func (obj *rocev2QPParameters) SetUdpSourcePort(value uint32) Rocev2QPParameters {

	obj.obj.UdpSourcePort = &value
	return obj
}

// Initial packet sequence number of the data transfer packet generated for this QP.
// InitialPsn returns a uint64
func (obj *rocev2QPParameters) InitialPsn() uint64 {

	return *obj.obj.InitialPsn

}

// Initial packet sequence number of the data transfer packet generated for this QP.
// InitialPsn returns a uint64
func (obj *rocev2QPParameters) HasInitialPsn() bool {
	return obj.obj.InitialPsn != nil
}

// Initial packet sequence number of the data transfer packet generated for this QP.
// SetInitialPsn sets the uint64 value in the Rocev2QPParameters object
func (obj *rocev2QPParameters) SetInitialPsn(value uint64) Rocev2QPParameters {

	obj.obj.InitialPsn = &value
	return obj
}

// Virtual Address where the data transfer from the remote QP will write to.
// VirtualAddress returns a string
func (obj *rocev2QPParameters) VirtualAddress() string {

	return *obj.obj.VirtualAddress

}

// Virtual Address where the data transfer from the remote QP will write to.
// VirtualAddress returns a string
func (obj *rocev2QPParameters) HasVirtualAddress() bool {
	return obj.obj.VirtualAddress != nil
}

// Virtual Address where the data transfer from the remote QP will write to.
// SetVirtualAddress sets the string value in the Rocev2QPParameters object
func (obj *rocev2QPParameters) SetVirtualAddress(value string) Rocev2QPParameters {

	obj.obj.VirtualAddress = &value
	return obj
}

// Remote Key linked to the QP's virtual address.
// RemoteKey returns a string
func (obj *rocev2QPParameters) RemoteKey() string {

	return *obj.obj.RemoteKey

}

// Remote Key linked to the QP's virtual address.
// RemoteKey returns a string
func (obj *rocev2QPParameters) HasRemoteKey() bool {
	return obj.obj.RemoteKey != nil
}

// Remote Key linked to the QP's virtual address.
// SetRemoteKey sets the string value in the Rocev2QPParameters object
func (obj *rocev2QPParameters) SetRemoteKey(value string) Rocev2QPParameters {

	obj.obj.RemoteKey = &value
	return obj
}

func (obj *rocev2QPParameters) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SourceQpNumber != nil {

		if *obj.obj.SourceQpNumber < 2 || *obj.obj.SourceQpNumber > 33554431 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("2 <= Rocev2QPParameters.SourceQpNumber <= 33554431 but Got %d", *obj.obj.SourceQpNumber))
		}

	}

	if obj.obj.Dscp != nil {

		if *obj.obj.Dscp > 63 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2QPParameters.Dscp <= 63 but Got %d", *obj.obj.Dscp))
		}

	}

	if obj.obj.UdpSourcePort != nil {

		if *obj.obj.UdpSourcePort > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2QPParameters.UdpSourcePort <= 65535 but Got %d", *obj.obj.UdpSourcePort))
		}

	}

	if obj.obj.InitialPsn != nil {

		if *obj.obj.InitialPsn > 33554431 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2QPParameters.InitialPsn <= 33554431 but Got %d", *obj.obj.InitialPsn))
		}

	}

	if obj.obj.VirtualAddress != nil {

		if len(*obj.obj.VirtualAddress) > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"None <= length of Rocev2QPParameters.VirtualAddress <= 16 but Got %d",
					len(*obj.obj.VirtualAddress)))
		}

	}

	if obj.obj.RemoteKey != nil {

		if len(*obj.obj.RemoteKey) > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"None <= length of Rocev2QPParameters.RemoteKey <= 8 but Got %d",
					len(*obj.obj.RemoteKey)))
		}

	}

}

func (obj *rocev2QPParameters) setDefault() {
	if obj.obj.SourceQpNumber == nil {
		obj.SetSourceQpNumber(2)
	}
	if obj.obj.Dscp == nil {
		obj.SetDscp(24)
	}
	if obj.obj.Ecn == nil {
		obj.SetEcn(Rocev2QPParametersEcn.ECT_1)

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
