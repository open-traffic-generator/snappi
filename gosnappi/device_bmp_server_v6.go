package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceBmpServerV6 *****
type deviceBmpServerV6 struct {
	validation
	obj                 *otg.DeviceBmpServerV6
	marshaller          marshalDeviceBmpServerV6
	unMarshaller        unMarshalDeviceBmpServerV6
	connectionHolder    DeviceBmpServerConnection
	prefixStorageHolder DeviceBmpServerPrefixStorage
}

func NewDeviceBmpServerV6() DeviceBmpServerV6 {
	obj := deviceBmpServerV6{obj: &otg.DeviceBmpServerV6{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceBmpServerV6) msg() *otg.DeviceBmpServerV6 {
	return obj.obj
}

func (obj *deviceBmpServerV6) setMsg(msg *otg.DeviceBmpServerV6) DeviceBmpServerV6 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceBmpServerV6 struct {
	obj *deviceBmpServerV6
}

type marshalDeviceBmpServerV6 interface {
	// ToProto marshals DeviceBmpServerV6 to protobuf object *otg.DeviceBmpServerV6
	ToProto() (*otg.DeviceBmpServerV6, error)
	// ToPbText marshals DeviceBmpServerV6 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceBmpServerV6 to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceBmpServerV6 to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceBmpServerV6 struct {
	obj *deviceBmpServerV6
}

type unMarshalDeviceBmpServerV6 interface {
	// FromProto unmarshals DeviceBmpServerV6 from protobuf object *otg.DeviceBmpServerV6
	FromProto(msg *otg.DeviceBmpServerV6) (DeviceBmpServerV6, error)
	// FromPbText unmarshals DeviceBmpServerV6 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceBmpServerV6 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceBmpServerV6 from JSON text
	FromJson(value string) error
}

func (obj *deviceBmpServerV6) Marshal() marshalDeviceBmpServerV6 {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceBmpServerV6{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceBmpServerV6) Unmarshal() unMarshalDeviceBmpServerV6 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceBmpServerV6{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceBmpServerV6) ToProto() (*otg.DeviceBmpServerV6, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceBmpServerV6) FromProto(msg *otg.DeviceBmpServerV6) (DeviceBmpServerV6, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceBmpServerV6) ToPbText() (string, error) {
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

func (m *unMarshaldeviceBmpServerV6) FromPbText(value string) error {
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

func (m *marshaldeviceBmpServerV6) ToYaml() (string, error) {
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

func (m *unMarshaldeviceBmpServerV6) FromYaml(value string) error {
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

func (m *marshaldeviceBmpServerV6) ToJson() (string, error) {
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

func (m *unMarshaldeviceBmpServerV6) FromJson(value string) error {
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

func (obj *deviceBmpServerV6) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceBmpServerV6) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceBmpServerV6) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceBmpServerV6) Clone() (DeviceBmpServerV6, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceBmpServerV6()
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

func (obj *deviceBmpServerV6) setNil() {
	obj.connectionHolder = nil
	obj.prefixStorageHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceBmpServerV6 is configuration for a BMP Server connected to a specific IPv6 BMP client.
type DeviceBmpServerV6 interface {
	Validation
	// msg marshals DeviceBmpServerV6 to protobuf object *otg.DeviceBmpServerV6
	// and doesn't set defaults
	msg() *otg.DeviceBmpServerV6
	// setMsg unmarshals DeviceBmpServerV6 from protobuf object *otg.DeviceBmpServerV6
	// and doesn't set defaults
	setMsg(*otg.DeviceBmpServerV6) DeviceBmpServerV6
	// provides marshal interface
	Marshal() marshalDeviceBmpServerV6
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceBmpServerV6
	// validate validates DeviceBmpServerV6
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceBmpServerV6, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in DeviceBmpServerV6.
	Name() string
	// SetName assigns string provided by user to DeviceBmpServerV6
	SetName(value string) DeviceBmpServerV6
	// ClientIp returns string, set in DeviceBmpServerV6.
	ClientIp() string
	// SetClientIp assigns string provided by user to DeviceBmpServerV6
	SetClientIp(value string) DeviceBmpServerV6
	// Connection returns DeviceBmpServerConnection, set in DeviceBmpServerV6.
	// DeviceBmpServerConnection is container of information about whether the BMP Server should operate in passive or active mode and corresponding information depending on the mode.
	Connection() DeviceBmpServerConnection
	// SetConnection assigns DeviceBmpServerConnection provided by user to DeviceBmpServerV6.
	// DeviceBmpServerConnection is container of information about whether the BMP Server should operate in passive or active mode and corresponding information depending on the mode.
	SetConnection(value DeviceBmpServerConnection) DeviceBmpServerV6
	// HasConnection checks if Connection has been set in DeviceBmpServerV6
	HasConnection() bool
	// PrefixStorage returns DeviceBmpServerPrefixStorage, set in DeviceBmpServerV6.
	// DeviceBmpServerPrefixStorage is optional object containing information about whether IPv4 and IPv6 unicast prefixes learned from the BMP client should be stored or not for future retrieval using get_states and  exceptions to the configured choice for the same. If the object is not included, by default IPv4 and IPv6 Unicast Prefixes are not stored by the BMP server and only received metrics  are incremented on receipt of IPv4 / IPv6 Unicast routes via BMP Monitor messages.
	PrefixStorage() DeviceBmpServerPrefixStorage
	// SetPrefixStorage assigns DeviceBmpServerPrefixStorage provided by user to DeviceBmpServerV6.
	// DeviceBmpServerPrefixStorage is optional object containing information about whether IPv4 and IPv6 unicast prefixes learned from the BMP client should be stored or not for future retrieval using get_states and  exceptions to the configured choice for the same. If the object is not included, by default IPv4 and IPv6 Unicast Prefixes are not stored by the BMP server and only received metrics  are incremented on receipt of IPv4 / IPv6 Unicast routes via BMP Monitor messages.
	SetPrefixStorage(value DeviceBmpServerPrefixStorage) DeviceBmpServerV6
	// HasPrefixStorage checks if PrefixStorage has been set in DeviceBmpServerV6
	HasPrefixStorage() bool
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *deviceBmpServerV6) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the DeviceBmpServerV6 object
func (obj *deviceBmpServerV6) SetName(value string) DeviceBmpServerV6 {

	obj.obj.Name = &value
	return obj
}

// The IPv6 address of the BMP client from which connections will be accepted from or initiated with by this BMP Server.
// ClientIp returns a string
func (obj *deviceBmpServerV6) ClientIp() string {

	return *obj.obj.ClientIp

}

// The IPv6 address of the BMP client from which connections will be accepted from or initiated with by this BMP Server.
// SetClientIp sets the string value in the DeviceBmpServerV6 object
func (obj *deviceBmpServerV6) SetClientIp(value string) DeviceBmpServerV6 {

	obj.obj.ClientIp = &value
	return obj
}

// Optional object containing information about whether the BMP Server should operate in passive or active mode and corresponding information depending on the mode.
// Connection returns a DeviceBmpServerConnection
func (obj *deviceBmpServerV6) Connection() DeviceBmpServerConnection {
	if obj.obj.Connection == nil {
		obj.obj.Connection = NewDeviceBmpServerConnection().msg()
	}
	if obj.connectionHolder == nil {
		obj.connectionHolder = &deviceBmpServerConnection{obj: obj.obj.Connection}
	}
	return obj.connectionHolder
}

// Optional object containing information about whether the BMP Server should operate in passive or active mode and corresponding information depending on the mode.
// Connection returns a DeviceBmpServerConnection
func (obj *deviceBmpServerV6) HasConnection() bool {
	return obj.obj.Connection != nil
}

// Optional object containing information about whether the BMP Server should operate in passive or active mode and corresponding information depending on the mode.
// SetConnection sets the DeviceBmpServerConnection value in the DeviceBmpServerV6 object
func (obj *deviceBmpServerV6) SetConnection(value DeviceBmpServerConnection) DeviceBmpServerV6 {

	obj.connectionHolder = nil
	obj.obj.Connection = value.msg()

	return obj
}

// Optional object containing information about whether IPv4 and IPv6 unicast prefixes learned from the BMP client should be stored or not for future retrieval using get_states and  exceptions to the configured choice for the same.  If the object is not included, by default IPv4 and IPv6 unicast prefixes are not stored by the BMP server and only received metrics are incremented on receipt of IPv4 / IPv6 unicast prefixes within  BMP Monitor messages.
// PrefixStorage returns a DeviceBmpServerPrefixStorage
func (obj *deviceBmpServerV6) PrefixStorage() DeviceBmpServerPrefixStorage {
	if obj.obj.PrefixStorage == nil {
		obj.obj.PrefixStorage = NewDeviceBmpServerPrefixStorage().msg()
	}
	if obj.prefixStorageHolder == nil {
		obj.prefixStorageHolder = &deviceBmpServerPrefixStorage{obj: obj.obj.PrefixStorage}
	}
	return obj.prefixStorageHolder
}

// Optional object containing information about whether IPv4 and IPv6 unicast prefixes learned from the BMP client should be stored or not for future retrieval using get_states and  exceptions to the configured choice for the same.  If the object is not included, by default IPv4 and IPv6 unicast prefixes are not stored by the BMP server and only received metrics are incremented on receipt of IPv4 / IPv6 unicast prefixes within  BMP Monitor messages.
// PrefixStorage returns a DeviceBmpServerPrefixStorage
func (obj *deviceBmpServerV6) HasPrefixStorage() bool {
	return obj.obj.PrefixStorage != nil
}

// Optional object containing information about whether IPv4 and IPv6 unicast prefixes learned from the BMP client should be stored or not for future retrieval using get_states and  exceptions to the configured choice for the same.  If the object is not included, by default IPv4 and IPv6 unicast prefixes are not stored by the BMP server and only received metrics are incremented on receipt of IPv4 / IPv6 unicast prefixes within  BMP Monitor messages.
// SetPrefixStorage sets the DeviceBmpServerPrefixStorage value in the DeviceBmpServerV6 object
func (obj *deviceBmpServerV6) SetPrefixStorage(value DeviceBmpServerPrefixStorage) DeviceBmpServerV6 {

	obj.prefixStorageHolder = nil
	obj.obj.PrefixStorage = value.msg()

	return obj
}

func (obj *deviceBmpServerV6) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface DeviceBmpServerV6")
	}

	// ClientIp is required
	if obj.obj.ClientIp == nil {
		vObj.validationErrors = append(vObj.validationErrors, "ClientIp is required field on interface DeviceBmpServerV6")
	}
	if obj.obj.ClientIp != nil {

		err := obj.validateIpv6(obj.ClientIp())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on DeviceBmpServerV6.ClientIp"))
		}

	}

	if obj.obj.Connection != nil {

		obj.Connection().validateObj(vObj, set_default)
	}

	if obj.obj.PrefixStorage != nil {

		obj.PrefixStorage().validateObj(vObj, set_default)
	}

}

func (obj *deviceBmpServerV6) setDefault() {

}
