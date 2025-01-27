package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceDhcpv6ClientDuidValue *****
type deviceDhcpv6ClientDuidValue struct {
	validation
	obj          *otg.DeviceDhcpv6ClientDuidValue
	marshaller   marshalDeviceDhcpv6ClientDuidValue
	unMarshaller unMarshalDeviceDhcpv6ClientDuidValue
}

func NewDeviceDhcpv6ClientDuidValue() DeviceDhcpv6ClientDuidValue {
	obj := deviceDhcpv6ClientDuidValue{obj: &otg.DeviceDhcpv6ClientDuidValue{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceDhcpv6ClientDuidValue) msg() *otg.DeviceDhcpv6ClientDuidValue {
	return obj.obj
}

func (obj *deviceDhcpv6ClientDuidValue) setMsg(msg *otg.DeviceDhcpv6ClientDuidValue) DeviceDhcpv6ClientDuidValue {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceDhcpv6ClientDuidValue struct {
	obj *deviceDhcpv6ClientDuidValue
}

type marshalDeviceDhcpv6ClientDuidValue interface {
	// ToProto marshals DeviceDhcpv6ClientDuidValue to protobuf object *otg.DeviceDhcpv6ClientDuidValue
	ToProto() (*otg.DeviceDhcpv6ClientDuidValue, error)
	// ToPbText marshals DeviceDhcpv6ClientDuidValue to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceDhcpv6ClientDuidValue to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceDhcpv6ClientDuidValue to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals DeviceDhcpv6ClientDuidValue to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldeviceDhcpv6ClientDuidValue struct {
	obj *deviceDhcpv6ClientDuidValue
}

type unMarshalDeviceDhcpv6ClientDuidValue interface {
	// FromProto unmarshals DeviceDhcpv6ClientDuidValue from protobuf object *otg.DeviceDhcpv6ClientDuidValue
	FromProto(msg *otg.DeviceDhcpv6ClientDuidValue) (DeviceDhcpv6ClientDuidValue, error)
	// FromPbText unmarshals DeviceDhcpv6ClientDuidValue from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceDhcpv6ClientDuidValue from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceDhcpv6ClientDuidValue from JSON text
	FromJson(value string) error
}

func (obj *deviceDhcpv6ClientDuidValue) Marshal() marshalDeviceDhcpv6ClientDuidValue {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceDhcpv6ClientDuidValue{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceDhcpv6ClientDuidValue) Unmarshal() unMarshalDeviceDhcpv6ClientDuidValue {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceDhcpv6ClientDuidValue{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceDhcpv6ClientDuidValue) ToProto() (*otg.DeviceDhcpv6ClientDuidValue, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceDhcpv6ClientDuidValue) FromProto(msg *otg.DeviceDhcpv6ClientDuidValue) (DeviceDhcpv6ClientDuidValue, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceDhcpv6ClientDuidValue) ToPbText() (string, error) {
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

func (m *unMarshaldeviceDhcpv6ClientDuidValue) FromPbText(value string) error {
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

func (m *marshaldeviceDhcpv6ClientDuidValue) ToYaml() (string, error) {
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

func (m *unMarshaldeviceDhcpv6ClientDuidValue) FromYaml(value string) error {
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

func (m *marshaldeviceDhcpv6ClientDuidValue) ToJsonRaw() (string, error) {
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

func (m *marshaldeviceDhcpv6ClientDuidValue) ToJson() (string, error) {
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

func (m *unMarshaldeviceDhcpv6ClientDuidValue) FromJson(value string) error {
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

func (obj *deviceDhcpv6ClientDuidValue) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceDhcpv6ClientDuidValue) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceDhcpv6ClientDuidValue) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceDhcpv6ClientDuidValue) Clone() (DeviceDhcpv6ClientDuidValue, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceDhcpv6ClientDuidValue()
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

// DeviceDhcpv6ClientDuidValue is the container for the DUID-EN. This consists of the 4-octet vendor's registered Private Enterprise Number  as maintained by IANA [IANA-PEN] followed by a unique identifier assigned by the vendor.
type DeviceDhcpv6ClientDuidValue interface {
	Validation
	// msg marshals DeviceDhcpv6ClientDuidValue to protobuf object *otg.DeviceDhcpv6ClientDuidValue
	// and doesn't set defaults
	msg() *otg.DeviceDhcpv6ClientDuidValue
	// setMsg unmarshals DeviceDhcpv6ClientDuidValue from protobuf object *otg.DeviceDhcpv6ClientDuidValue
	// and doesn't set defaults
	setMsg(*otg.DeviceDhcpv6ClientDuidValue) DeviceDhcpv6ClientDuidValue
	// provides marshal interface
	Marshal() marshalDeviceDhcpv6ClientDuidValue
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceDhcpv6ClientDuidValue
	// validate validates DeviceDhcpv6ClientDuidValue
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceDhcpv6ClientDuidValue, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EnterpriseId returns uint32, set in DeviceDhcpv6ClientDuidValue.
	EnterpriseId() uint32
	// SetEnterpriseId assigns uint32 provided by user to DeviceDhcpv6ClientDuidValue
	SetEnterpriseId(value uint32) DeviceDhcpv6ClientDuidValue
	// HasEnterpriseId checks if EnterpriseId has been set in DeviceDhcpv6ClientDuidValue
	HasEnterpriseId() bool
	// VendorId returns uint32, set in DeviceDhcpv6ClientDuidValue.
	VendorId() uint32
	// SetVendorId assigns uint32 provided by user to DeviceDhcpv6ClientDuidValue
	SetVendorId(value uint32) DeviceDhcpv6ClientDuidValue
	// HasVendorId checks if VendorId has been set in DeviceDhcpv6ClientDuidValue
	HasVendorId() bool
}

// 4-octet vendor's registered Private Enterprise Number as maintained by IANA [IANA-PEN].
// EnterpriseId returns a uint32
func (obj *deviceDhcpv6ClientDuidValue) EnterpriseId() uint32 {

	return *obj.obj.EnterpriseId

}

// 4-octet vendor's registered Private Enterprise Number as maintained by IANA [IANA-PEN].
// EnterpriseId returns a uint32
func (obj *deviceDhcpv6ClientDuidValue) HasEnterpriseId() bool {
	return obj.obj.EnterpriseId != nil
}

// 4-octet vendor's registered Private Enterprise Number as maintained by IANA [IANA-PEN].
// SetEnterpriseId sets the uint32 value in the DeviceDhcpv6ClientDuidValue object
func (obj *deviceDhcpv6ClientDuidValue) SetEnterpriseId(value uint32) DeviceDhcpv6ClientDuidValue {

	obj.obj.EnterpriseId = &value
	return obj
}

// Unique identifier assigned by the vendor.
// VendorId returns a uint32
func (obj *deviceDhcpv6ClientDuidValue) VendorId() uint32 {

	return *obj.obj.VendorId

}

// Unique identifier assigned by the vendor.
// VendorId returns a uint32
func (obj *deviceDhcpv6ClientDuidValue) HasVendorId() bool {
	return obj.obj.VendorId != nil
}

// Unique identifier assigned by the vendor.
// SetVendorId sets the uint32 value in the DeviceDhcpv6ClientDuidValue object
func (obj *deviceDhcpv6ClientDuidValue) SetVendorId(value uint32) DeviceDhcpv6ClientDuidValue {

	obj.obj.VendorId = &value
	return obj
}

func (obj *deviceDhcpv6ClientDuidValue) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.EnterpriseId != nil {

		if *obj.obj.EnterpriseId < 1 || *obj.obj.EnterpriseId > 2147483647 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= DeviceDhcpv6ClientDuidValue.EnterpriseId <= 2147483647 but Got %d", *obj.obj.EnterpriseId))
		}

	}

	if obj.obj.VendorId != nil {

		if *obj.obj.VendorId < 1 || *obj.obj.VendorId > 2147483647 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= DeviceDhcpv6ClientDuidValue.VendorId <= 2147483647 but Got %d", *obj.obj.VendorId))
		}

	}

}

func (obj *deviceDhcpv6ClientDuidValue) setDefault() {
	if obj.obj.EnterpriseId == nil {
		obj.SetEnterpriseId(10)
	}
	if obj.obj.VendorId == nil {
		obj.SetVendorId(10)
	}

}
