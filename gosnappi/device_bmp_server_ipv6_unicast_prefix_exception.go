package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceBmpServerIpv6UnicastPrefixException *****
type deviceBmpServerIpv6UnicastPrefixException struct {
	validation
	obj          *otg.DeviceBmpServerIpv6UnicastPrefixException
	marshaller   marshalDeviceBmpServerIpv6UnicastPrefixException
	unMarshaller unMarshalDeviceBmpServerIpv6UnicastPrefixException
}

func NewDeviceBmpServerIpv6UnicastPrefixException() DeviceBmpServerIpv6UnicastPrefixException {
	obj := deviceBmpServerIpv6UnicastPrefixException{obj: &otg.DeviceBmpServerIpv6UnicastPrefixException{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceBmpServerIpv6UnicastPrefixException) msg() *otg.DeviceBmpServerIpv6UnicastPrefixException {
	return obj.obj
}

func (obj *deviceBmpServerIpv6UnicastPrefixException) setMsg(msg *otg.DeviceBmpServerIpv6UnicastPrefixException) DeviceBmpServerIpv6UnicastPrefixException {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceBmpServerIpv6UnicastPrefixException struct {
	obj *deviceBmpServerIpv6UnicastPrefixException
}

type marshalDeviceBmpServerIpv6UnicastPrefixException interface {
	// ToProto marshals DeviceBmpServerIpv6UnicastPrefixException to protobuf object *otg.DeviceBmpServerIpv6UnicastPrefixException
	ToProto() (*otg.DeviceBmpServerIpv6UnicastPrefixException, error)
	// ToPbText marshals DeviceBmpServerIpv6UnicastPrefixException to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceBmpServerIpv6UnicastPrefixException to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceBmpServerIpv6UnicastPrefixException to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceBmpServerIpv6UnicastPrefixException struct {
	obj *deviceBmpServerIpv6UnicastPrefixException
}

type unMarshalDeviceBmpServerIpv6UnicastPrefixException interface {
	// FromProto unmarshals DeviceBmpServerIpv6UnicastPrefixException from protobuf object *otg.DeviceBmpServerIpv6UnicastPrefixException
	FromProto(msg *otg.DeviceBmpServerIpv6UnicastPrefixException) (DeviceBmpServerIpv6UnicastPrefixException, error)
	// FromPbText unmarshals DeviceBmpServerIpv6UnicastPrefixException from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceBmpServerIpv6UnicastPrefixException from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceBmpServerIpv6UnicastPrefixException from JSON text
	FromJson(value string) error
}

func (obj *deviceBmpServerIpv6UnicastPrefixException) Marshal() marshalDeviceBmpServerIpv6UnicastPrefixException {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceBmpServerIpv6UnicastPrefixException{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceBmpServerIpv6UnicastPrefixException) Unmarshal() unMarshalDeviceBmpServerIpv6UnicastPrefixException {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceBmpServerIpv6UnicastPrefixException{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceBmpServerIpv6UnicastPrefixException) ToProto() (*otg.DeviceBmpServerIpv6UnicastPrefixException, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceBmpServerIpv6UnicastPrefixException) FromProto(msg *otg.DeviceBmpServerIpv6UnicastPrefixException) (DeviceBmpServerIpv6UnicastPrefixException, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceBmpServerIpv6UnicastPrefixException) ToPbText() (string, error) {
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

func (m *unMarshaldeviceBmpServerIpv6UnicastPrefixException) FromPbText(value string) error {
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

func (m *marshaldeviceBmpServerIpv6UnicastPrefixException) ToYaml() (string, error) {
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

func (m *unMarshaldeviceBmpServerIpv6UnicastPrefixException) FromYaml(value string) error {
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

func (m *marshaldeviceBmpServerIpv6UnicastPrefixException) ToJson() (string, error) {
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

func (m *unMarshaldeviceBmpServerIpv6UnicastPrefixException) FromJson(value string) error {
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

func (obj *deviceBmpServerIpv6UnicastPrefixException) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceBmpServerIpv6UnicastPrefixException) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceBmpServerIpv6UnicastPrefixException) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceBmpServerIpv6UnicastPrefixException) Clone() (DeviceBmpServerIpv6UnicastPrefixException, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceBmpServerIpv6UnicastPrefixException()
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

// DeviceBmpServerIpv6UnicastPrefixException is one exception to the specification is specified by a combination of an IPv6 prefix and a prefix length.  If the received prefix masked upto the exception's prefix length matches the prefix specified in the exception, the received prefix is deemed as having matched the specified exception e.g. received prefix 1:1:1:1::/64  and 1:1:1:2::/64 would match specified exception of 1:1:1::/48 but 1:1::/16 or 2:2:2:2::/64 would not.
type DeviceBmpServerIpv6UnicastPrefixException interface {
	Validation
	// msg marshals DeviceBmpServerIpv6UnicastPrefixException to protobuf object *otg.DeviceBmpServerIpv6UnicastPrefixException
	// and doesn't set defaults
	msg() *otg.DeviceBmpServerIpv6UnicastPrefixException
	// setMsg unmarshals DeviceBmpServerIpv6UnicastPrefixException from protobuf object *otg.DeviceBmpServerIpv6UnicastPrefixException
	// and doesn't set defaults
	setMsg(*otg.DeviceBmpServerIpv6UnicastPrefixException) DeviceBmpServerIpv6UnicastPrefixException
	// provides marshal interface
	Marshal() marshalDeviceBmpServerIpv6UnicastPrefixException
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceBmpServerIpv6UnicastPrefixException
	// validate validates DeviceBmpServerIpv6UnicastPrefixException
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceBmpServerIpv6UnicastPrefixException, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv6Prefix returns string, set in DeviceBmpServerIpv6UnicastPrefixException.
	Ipv6Prefix() string
	// SetIpv6Prefix assigns string provided by user to DeviceBmpServerIpv6UnicastPrefixException
	SetIpv6Prefix(value string) DeviceBmpServerIpv6UnicastPrefixException
	// HasIpv6Prefix checks if Ipv6Prefix has been set in DeviceBmpServerIpv6UnicastPrefixException
	HasIpv6Prefix() bool
	// PrefixLength returns uint32, set in DeviceBmpServerIpv6UnicastPrefixException.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to DeviceBmpServerIpv6UnicastPrefixException
	SetPrefixLength(value uint32) DeviceBmpServerIpv6UnicastPrefixException
	// HasPrefixLength checks if PrefixLength has been set in DeviceBmpServerIpv6UnicastPrefixException
	HasPrefixLength() bool
}

// The IPv6 prefix which, combined with the prefix_length, is used to determine if received IPv6 unicast prefix matches this exception or not.
// Ipv6Prefix returns a string
func (obj *deviceBmpServerIpv6UnicastPrefixException) Ipv6Prefix() string {

	return *obj.obj.Ipv6Prefix

}

// The IPv6 prefix which, combined with the prefix_length, is used to determine if received IPv6 unicast prefix matches this exception or not.
// Ipv6Prefix returns a string
func (obj *deviceBmpServerIpv6UnicastPrefixException) HasIpv6Prefix() bool {
	return obj.obj.Ipv6Prefix != nil
}

// The IPv6 prefix which, combined with the prefix_length, is used to determine if received IPv6 unicast prefix matches this exception or not.
// SetIpv6Prefix sets the string value in the DeviceBmpServerIpv6UnicastPrefixException object
func (obj *deviceBmpServerIpv6UnicastPrefixException) SetIpv6Prefix(value string) DeviceBmpServerIpv6UnicastPrefixException {

	obj.obj.Ipv6Prefix = &value
	return obj
}

// The prefix length which, combined with the ipv6_prefix, is used to determine if received IPv6 unicast prefix matches this exception or not.
// PrefixLength returns a uint32
func (obj *deviceBmpServerIpv6UnicastPrefixException) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// The prefix length which, combined with the ipv6_prefix, is used to determine if received IPv6 unicast prefix matches this exception or not.
// PrefixLength returns a uint32
func (obj *deviceBmpServerIpv6UnicastPrefixException) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// The prefix length which, combined with the ipv6_prefix, is used to determine if received IPv6 unicast prefix matches this exception or not.
// SetPrefixLength sets the uint32 value in the DeviceBmpServerIpv6UnicastPrefixException object
func (obj *deviceBmpServerIpv6UnicastPrefixException) SetPrefixLength(value uint32) DeviceBmpServerIpv6UnicastPrefixException {

	obj.obj.PrefixLength = &value
	return obj
}

func (obj *deviceBmpServerIpv6UnicastPrefixException) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv6Prefix != nil {

		err := obj.validateIpv6(obj.Ipv6Prefix())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on DeviceBmpServerIpv6UnicastPrefixException.Ipv6Prefix"))
		}

	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= DeviceBmpServerIpv6UnicastPrefixException.PrefixLength <= 128 but Got %d", *obj.obj.PrefixLength))
		}

	}

}

func (obj *deviceBmpServerIpv6UnicastPrefixException) setDefault() {
	if obj.obj.Ipv6Prefix == nil {
		obj.SetIpv6Prefix("::")
	}
	if obj.obj.PrefixLength == nil {
		obj.SetPrefixLength(48)
	}

}
