package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceBmpServerIpv4UnicastPrefixException *****
type deviceBmpServerIpv4UnicastPrefixException struct {
	validation
	obj          *otg.DeviceBmpServerIpv4UnicastPrefixException
	marshaller   marshalDeviceBmpServerIpv4UnicastPrefixException
	unMarshaller unMarshalDeviceBmpServerIpv4UnicastPrefixException
}

func NewDeviceBmpServerIpv4UnicastPrefixException() DeviceBmpServerIpv4UnicastPrefixException {
	obj := deviceBmpServerIpv4UnicastPrefixException{obj: &otg.DeviceBmpServerIpv4UnicastPrefixException{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceBmpServerIpv4UnicastPrefixException) msg() *otg.DeviceBmpServerIpv4UnicastPrefixException {
	return obj.obj
}

func (obj *deviceBmpServerIpv4UnicastPrefixException) setMsg(msg *otg.DeviceBmpServerIpv4UnicastPrefixException) DeviceBmpServerIpv4UnicastPrefixException {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceBmpServerIpv4UnicastPrefixException struct {
	obj *deviceBmpServerIpv4UnicastPrefixException
}

type marshalDeviceBmpServerIpv4UnicastPrefixException interface {
	// ToProto marshals DeviceBmpServerIpv4UnicastPrefixException to protobuf object *otg.DeviceBmpServerIpv4UnicastPrefixException
	ToProto() (*otg.DeviceBmpServerIpv4UnicastPrefixException, error)
	// ToPbText marshals DeviceBmpServerIpv4UnicastPrefixException to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceBmpServerIpv4UnicastPrefixException to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceBmpServerIpv4UnicastPrefixException to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceBmpServerIpv4UnicastPrefixException struct {
	obj *deviceBmpServerIpv4UnicastPrefixException
}

type unMarshalDeviceBmpServerIpv4UnicastPrefixException interface {
	// FromProto unmarshals DeviceBmpServerIpv4UnicastPrefixException from protobuf object *otg.DeviceBmpServerIpv4UnicastPrefixException
	FromProto(msg *otg.DeviceBmpServerIpv4UnicastPrefixException) (DeviceBmpServerIpv4UnicastPrefixException, error)
	// FromPbText unmarshals DeviceBmpServerIpv4UnicastPrefixException from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceBmpServerIpv4UnicastPrefixException from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceBmpServerIpv4UnicastPrefixException from JSON text
	FromJson(value string) error
}

func (obj *deviceBmpServerIpv4UnicastPrefixException) Marshal() marshalDeviceBmpServerIpv4UnicastPrefixException {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceBmpServerIpv4UnicastPrefixException{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceBmpServerIpv4UnicastPrefixException) Unmarshal() unMarshalDeviceBmpServerIpv4UnicastPrefixException {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceBmpServerIpv4UnicastPrefixException{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceBmpServerIpv4UnicastPrefixException) ToProto() (*otg.DeviceBmpServerIpv4UnicastPrefixException, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceBmpServerIpv4UnicastPrefixException) FromProto(msg *otg.DeviceBmpServerIpv4UnicastPrefixException) (DeviceBmpServerIpv4UnicastPrefixException, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceBmpServerIpv4UnicastPrefixException) ToPbText() (string, error) {
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

func (m *unMarshaldeviceBmpServerIpv4UnicastPrefixException) FromPbText(value string) error {
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

func (m *marshaldeviceBmpServerIpv4UnicastPrefixException) ToYaml() (string, error) {
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

func (m *unMarshaldeviceBmpServerIpv4UnicastPrefixException) FromYaml(value string) error {
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

func (m *marshaldeviceBmpServerIpv4UnicastPrefixException) ToJson() (string, error) {
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

func (m *unMarshaldeviceBmpServerIpv4UnicastPrefixException) FromJson(value string) error {
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

func (obj *deviceBmpServerIpv4UnicastPrefixException) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceBmpServerIpv4UnicastPrefixException) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceBmpServerIpv4UnicastPrefixException) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceBmpServerIpv4UnicastPrefixException) Clone() (DeviceBmpServerIpv4UnicastPrefixException, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceBmpServerIpv4UnicastPrefixException()
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

// DeviceBmpServerIpv4UnicastPrefixException is one exception to the specification is specified by a combination of an IPv4 prefix and a prefix length.  If the received prefix masked upto the exception's prefix length matches the prefix specified in the exception, the received prefix is deemed as having matched the specified exception e.g. received prefix 172.16.1.0/24 and 172.16.2.0/24 would match specified exception of 172.16.0.0/16 but 172.0.0.0/8 or 192.16.2.0/24 would not.
type DeviceBmpServerIpv4UnicastPrefixException interface {
	Validation
	// msg marshals DeviceBmpServerIpv4UnicastPrefixException to protobuf object *otg.DeviceBmpServerIpv4UnicastPrefixException
	// and doesn't set defaults
	msg() *otg.DeviceBmpServerIpv4UnicastPrefixException
	// setMsg unmarshals DeviceBmpServerIpv4UnicastPrefixException from protobuf object *otg.DeviceBmpServerIpv4UnicastPrefixException
	// and doesn't set defaults
	setMsg(*otg.DeviceBmpServerIpv4UnicastPrefixException) DeviceBmpServerIpv4UnicastPrefixException
	// provides marshal interface
	Marshal() marshalDeviceBmpServerIpv4UnicastPrefixException
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceBmpServerIpv4UnicastPrefixException
	// validate validates DeviceBmpServerIpv4UnicastPrefixException
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceBmpServerIpv4UnicastPrefixException, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4Prefix returns string, set in DeviceBmpServerIpv4UnicastPrefixException.
	Ipv4Prefix() string
	// SetIpv4Prefix assigns string provided by user to DeviceBmpServerIpv4UnicastPrefixException
	SetIpv4Prefix(value string) DeviceBmpServerIpv4UnicastPrefixException
	// HasIpv4Prefix checks if Ipv4Prefix has been set in DeviceBmpServerIpv4UnicastPrefixException
	HasIpv4Prefix() bool
	// PrefixLength returns uint32, set in DeviceBmpServerIpv4UnicastPrefixException.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to DeviceBmpServerIpv4UnicastPrefixException
	SetPrefixLength(value uint32) DeviceBmpServerIpv4UnicastPrefixException
	// HasPrefixLength checks if PrefixLength has been set in DeviceBmpServerIpv4UnicastPrefixException
	HasPrefixLength() bool
}

// The IPv4 prefix which, combined with the prefix_length, is used to determine if received IPv4 unicast prefix matches this exception or not.
// Ipv4Prefix returns a string
func (obj *deviceBmpServerIpv4UnicastPrefixException) Ipv4Prefix() string {

	return *obj.obj.Ipv4Prefix

}

// The IPv4 prefix which, combined with the prefix_length, is used to determine if received IPv4 unicast prefix matches this exception or not.
// Ipv4Prefix returns a string
func (obj *deviceBmpServerIpv4UnicastPrefixException) HasIpv4Prefix() bool {
	return obj.obj.Ipv4Prefix != nil
}

// The IPv4 prefix which, combined with the prefix_length, is used to determine if received IPv4 unicast prefix matches this exception or not.
// SetIpv4Prefix sets the string value in the DeviceBmpServerIpv4UnicastPrefixException object
func (obj *deviceBmpServerIpv4UnicastPrefixException) SetIpv4Prefix(value string) DeviceBmpServerIpv4UnicastPrefixException {

	obj.obj.Ipv4Prefix = &value
	return obj
}

// The prefix length which, combined with the ipv4_prefix, is used to determine if received IPv4 unicast prefix matches this exception or not.
// PrefixLength returns a uint32
func (obj *deviceBmpServerIpv4UnicastPrefixException) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// The prefix length which, combined with the ipv4_prefix, is used to determine if received IPv4 unicast prefix matches this exception or not.
// PrefixLength returns a uint32
func (obj *deviceBmpServerIpv4UnicastPrefixException) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// The prefix length which, combined with the ipv4_prefix, is used to determine if received IPv4 unicast prefix matches this exception or not.
// SetPrefixLength sets the uint32 value in the DeviceBmpServerIpv4UnicastPrefixException object
func (obj *deviceBmpServerIpv4UnicastPrefixException) SetPrefixLength(value uint32) DeviceBmpServerIpv4UnicastPrefixException {

	obj.obj.PrefixLength = &value
	return obj
}

func (obj *deviceBmpServerIpv4UnicastPrefixException) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv4Prefix != nil {

		err := obj.validateIpv4(obj.Ipv4Prefix())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on DeviceBmpServerIpv4UnicastPrefixException.Ipv4Prefix"))
		}

	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= DeviceBmpServerIpv4UnicastPrefixException.PrefixLength <= 32 but Got %d", *obj.obj.PrefixLength))
		}

	}

}

func (obj *deviceBmpServerIpv4UnicastPrefixException) setDefault() {
	if obj.obj.Ipv4Prefix == nil {
		obj.SetIpv4Prefix("0.0.0.0")
	}
	if obj.obj.PrefixLength == nil {
		obj.SetPrefixLength(16)
	}

}
