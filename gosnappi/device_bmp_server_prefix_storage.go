package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceBmpServerPrefixStorage *****
type deviceBmpServerPrefixStorage struct {
	validation
	obj               *otg.DeviceBmpServerPrefixStorage
	marshaller        marshalDeviceBmpServerPrefixStorage
	unMarshaller      unMarshalDeviceBmpServerPrefixStorage
	ipv4UnicastHolder DeviceBmpServerIpv4UnicastPrefixStorage
	ipv6UnicastHolder DeviceBmpServerIpv6UnicastPrefixStorage
}

func NewDeviceBmpServerPrefixStorage() DeviceBmpServerPrefixStorage {
	obj := deviceBmpServerPrefixStorage{obj: &otg.DeviceBmpServerPrefixStorage{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceBmpServerPrefixStorage) msg() *otg.DeviceBmpServerPrefixStorage {
	return obj.obj
}

func (obj *deviceBmpServerPrefixStorage) setMsg(msg *otg.DeviceBmpServerPrefixStorage) DeviceBmpServerPrefixStorage {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceBmpServerPrefixStorage struct {
	obj *deviceBmpServerPrefixStorage
}

type marshalDeviceBmpServerPrefixStorage interface {
	// ToProto marshals DeviceBmpServerPrefixStorage to protobuf object *otg.DeviceBmpServerPrefixStorage
	ToProto() (*otg.DeviceBmpServerPrefixStorage, error)
	// ToPbText marshals DeviceBmpServerPrefixStorage to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceBmpServerPrefixStorage to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceBmpServerPrefixStorage to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceBmpServerPrefixStorage struct {
	obj *deviceBmpServerPrefixStorage
}

type unMarshalDeviceBmpServerPrefixStorage interface {
	// FromProto unmarshals DeviceBmpServerPrefixStorage from protobuf object *otg.DeviceBmpServerPrefixStorage
	FromProto(msg *otg.DeviceBmpServerPrefixStorage) (DeviceBmpServerPrefixStorage, error)
	// FromPbText unmarshals DeviceBmpServerPrefixStorage from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceBmpServerPrefixStorage from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceBmpServerPrefixStorage from JSON text
	FromJson(value string) error
}

func (obj *deviceBmpServerPrefixStorage) Marshal() marshalDeviceBmpServerPrefixStorage {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceBmpServerPrefixStorage{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceBmpServerPrefixStorage) Unmarshal() unMarshalDeviceBmpServerPrefixStorage {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceBmpServerPrefixStorage{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceBmpServerPrefixStorage) ToProto() (*otg.DeviceBmpServerPrefixStorage, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceBmpServerPrefixStorage) FromProto(msg *otg.DeviceBmpServerPrefixStorage) (DeviceBmpServerPrefixStorage, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceBmpServerPrefixStorage) ToPbText() (string, error) {
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

func (m *unMarshaldeviceBmpServerPrefixStorage) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshaldeviceBmpServerPrefixStorage) ToYaml() (string, error) {
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

func (m *unMarshaldeviceBmpServerPrefixStorage) FromYaml(value string) error {
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
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshaldeviceBmpServerPrefixStorage) ToJson() (string, error) {
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

func (m *unMarshaldeviceBmpServerPrefixStorage) FromJson(value string) error {
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
	m.obj.setNil()
	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *deviceBmpServerPrefixStorage) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceBmpServerPrefixStorage) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceBmpServerPrefixStorage) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceBmpServerPrefixStorage) Clone() (DeviceBmpServerPrefixStorage, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceBmpServerPrefixStorage()
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

func (obj *deviceBmpServerPrefixStorage) setNil() {
	obj.ipv4UnicastHolder = nil
	obj.ipv6UnicastHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceBmpServerPrefixStorage is optional object containing information about whether IPv4 and IPv6 unicast prefixes learned from the BMP client should be stored or not for future retrieval using get_states and  exceptions to the configured choice for the same. If the object is not included, by default IPv4 and IPv6 Unicast Prefixes are not stored by the BMP server and only received metrics  are incremented on receipt of IPv4 / IPv6 Unicast routes via BMP Monitor messages.
type DeviceBmpServerPrefixStorage interface {
	Validation
	// msg marshals DeviceBmpServerPrefixStorage to protobuf object *otg.DeviceBmpServerPrefixStorage
	// and doesn't set defaults
	msg() *otg.DeviceBmpServerPrefixStorage
	// setMsg unmarshals DeviceBmpServerPrefixStorage from protobuf object *otg.DeviceBmpServerPrefixStorage
	// and doesn't set defaults
	setMsg(*otg.DeviceBmpServerPrefixStorage) DeviceBmpServerPrefixStorage
	// provides marshal interface
	Marshal() marshalDeviceBmpServerPrefixStorage
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceBmpServerPrefixStorage
	// validate validates DeviceBmpServerPrefixStorage
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceBmpServerPrefixStorage, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4Unicast returns DeviceBmpServerIpv4UnicastPrefixStorage, set in DeviceBmpServerPrefixStorage.
	// DeviceBmpServerIpv4UnicastPrefixStorage is optional object containing information about whether IPv4 unicast prefixes learned from the BMP client should be stored or not for future retrieval using get_states and  exceptions to the configured choice for the same. If the object is not included, by default IPv4 unicast prefixes are not stored by the BMP server.
	Ipv4Unicast() DeviceBmpServerIpv4UnicastPrefixStorage
	// SetIpv4Unicast assigns DeviceBmpServerIpv4UnicastPrefixStorage provided by user to DeviceBmpServerPrefixStorage.
	// DeviceBmpServerIpv4UnicastPrefixStorage is optional object containing information about whether IPv4 unicast prefixes learned from the BMP client should be stored or not for future retrieval using get_states and  exceptions to the configured choice for the same. If the object is not included, by default IPv4 unicast prefixes are not stored by the BMP server.
	SetIpv4Unicast(value DeviceBmpServerIpv4UnicastPrefixStorage) DeviceBmpServerPrefixStorage
	// HasIpv4Unicast checks if Ipv4Unicast has been set in DeviceBmpServerPrefixStorage
	HasIpv4Unicast() bool
	// Ipv6Unicast returns DeviceBmpServerIpv6UnicastPrefixStorage, set in DeviceBmpServerPrefixStorage.
	// DeviceBmpServerIpv6UnicastPrefixStorage is optional object containing information about whether IPv6 unicast prefixes learned from the BMP client should be stored or not for future retrieval using get_states and  exceptions to the configured choice for the same. If the object is not included, by default IPv6 unicast prefixes are not stored by the BMP server.
	Ipv6Unicast() DeviceBmpServerIpv6UnicastPrefixStorage
	// SetIpv6Unicast assigns DeviceBmpServerIpv6UnicastPrefixStorage provided by user to DeviceBmpServerPrefixStorage.
	// DeviceBmpServerIpv6UnicastPrefixStorage is optional object containing information about whether IPv6 unicast prefixes learned from the BMP client should be stored or not for future retrieval using get_states and  exceptions to the configured choice for the same. If the object is not included, by default IPv6 unicast prefixes are not stored by the BMP server.
	SetIpv6Unicast(value DeviceBmpServerIpv6UnicastPrefixStorage) DeviceBmpServerPrefixStorage
	// HasIpv6Unicast checks if Ipv6Unicast has been set in DeviceBmpServerPrefixStorage
	HasIpv6Unicast() bool
	setNil()
}

// Optional object containing information about whether IPv4 unicast prefixes learned from the BMP client should be stored or not for future retrieval using get_states and  exceptions to the configured choice for the same. If the object is not included, by default IPv4 unicast prefixes are not stored by the BMP server.
// Ipv4Unicast returns a DeviceBmpServerIpv4UnicastPrefixStorage
func (obj *deviceBmpServerPrefixStorage) Ipv4Unicast() DeviceBmpServerIpv4UnicastPrefixStorage {
	if obj.obj.Ipv4Unicast == nil {
		obj.obj.Ipv4Unicast = NewDeviceBmpServerIpv4UnicastPrefixStorage().msg()
	}
	if obj.ipv4UnicastHolder == nil {
		obj.ipv4UnicastHolder = &deviceBmpServerIpv4UnicastPrefixStorage{obj: obj.obj.Ipv4Unicast}
	}
	return obj.ipv4UnicastHolder
}

// Optional object containing information about whether IPv4 unicast prefixes learned from the BMP client should be stored or not for future retrieval using get_states and  exceptions to the configured choice for the same. If the object is not included, by default IPv4 unicast prefixes are not stored by the BMP server.
// Ipv4Unicast returns a DeviceBmpServerIpv4UnicastPrefixStorage
func (obj *deviceBmpServerPrefixStorage) HasIpv4Unicast() bool {
	return obj.obj.Ipv4Unicast != nil
}

// Optional object containing information about whether IPv4 unicast prefixes learned from the BMP client should be stored or not for future retrieval using get_states and  exceptions to the configured choice for the same. If the object is not included, by default IPv4 unicast prefixes are not stored by the BMP server.
// SetIpv4Unicast sets the DeviceBmpServerIpv4UnicastPrefixStorage value in the DeviceBmpServerPrefixStorage object
func (obj *deviceBmpServerPrefixStorage) SetIpv4Unicast(value DeviceBmpServerIpv4UnicastPrefixStorage) DeviceBmpServerPrefixStorage {

	obj.ipv4UnicastHolder = nil
	obj.obj.Ipv4Unicast = value.msg()

	return obj
}

// Optional object containing information about whether IPv6 unicast prefixes learned from the BMP client should be stored or not for future retrieval using get_states and  exceptions to the configured choice for the same. If the object is not included, by default IPv6 unicast prefixes are not stored by the BMP server.
// Ipv6Unicast returns a DeviceBmpServerIpv6UnicastPrefixStorage
func (obj *deviceBmpServerPrefixStorage) Ipv6Unicast() DeviceBmpServerIpv6UnicastPrefixStorage {
	if obj.obj.Ipv6Unicast == nil {
		obj.obj.Ipv6Unicast = NewDeviceBmpServerIpv6UnicastPrefixStorage().msg()
	}
	if obj.ipv6UnicastHolder == nil {
		obj.ipv6UnicastHolder = &deviceBmpServerIpv6UnicastPrefixStorage{obj: obj.obj.Ipv6Unicast}
	}
	return obj.ipv6UnicastHolder
}

// Optional object containing information about whether IPv6 unicast prefixes learned from the BMP client should be stored or not for future retrieval using get_states and  exceptions to the configured choice for the same. If the object is not included, by default IPv6 unicast prefixes are not stored by the BMP server.
// Ipv6Unicast returns a DeviceBmpServerIpv6UnicastPrefixStorage
func (obj *deviceBmpServerPrefixStorage) HasIpv6Unicast() bool {
	return obj.obj.Ipv6Unicast != nil
}

// Optional object containing information about whether IPv6 unicast prefixes learned from the BMP client should be stored or not for future retrieval using get_states and  exceptions to the configured choice for the same. If the object is not included, by default IPv6 unicast prefixes are not stored by the BMP server.
// SetIpv6Unicast sets the DeviceBmpServerIpv6UnicastPrefixStorage value in the DeviceBmpServerPrefixStorage object
func (obj *deviceBmpServerPrefixStorage) SetIpv6Unicast(value DeviceBmpServerIpv6UnicastPrefixStorage) DeviceBmpServerPrefixStorage {

	obj.ipv6UnicastHolder = nil
	obj.obj.Ipv6Unicast = value.msg()

	return obj
}

func (obj *deviceBmpServerPrefixStorage) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv4Unicast != nil {

		obj.Ipv4Unicast().validateObj(vObj, set_default)
	}

	if obj.obj.Ipv6Unicast != nil {

		obj.Ipv6Unicast().validateObj(vObj, set_default)
	}

}

func (obj *deviceBmpServerPrefixStorage) setDefault() {

}
