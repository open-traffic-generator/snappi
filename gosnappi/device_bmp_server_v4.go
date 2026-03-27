package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceBmpServerV4 *****
type deviceBmpServerV4 struct {
	validation
	obj                 *otg.DeviceBmpServerV4
	marshaller          marshalDeviceBmpServerV4
	unMarshaller        unMarshalDeviceBmpServerV4
	connectionHolder    DeviceBmpServerConnection
	prefixStorageHolder DeviceBmpServerPrefixStorage
}

func NewDeviceBmpServerV4() DeviceBmpServerV4 {
	obj := deviceBmpServerV4{obj: &otg.DeviceBmpServerV4{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceBmpServerV4) msg() *otg.DeviceBmpServerV4 {
	return obj.obj
}

func (obj *deviceBmpServerV4) setMsg(msg *otg.DeviceBmpServerV4) DeviceBmpServerV4 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceBmpServerV4 struct {
	obj *deviceBmpServerV4
}

type marshalDeviceBmpServerV4 interface {
	// ToProto marshals DeviceBmpServerV4 to protobuf object *otg.DeviceBmpServerV4
	ToProto() (*otg.DeviceBmpServerV4, error)
	// ToPbText marshals DeviceBmpServerV4 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceBmpServerV4 to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceBmpServerV4 to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceBmpServerV4 struct {
	obj *deviceBmpServerV4
}

type unMarshalDeviceBmpServerV4 interface {
	// FromProto unmarshals DeviceBmpServerV4 from protobuf object *otg.DeviceBmpServerV4
	FromProto(msg *otg.DeviceBmpServerV4) (DeviceBmpServerV4, error)
	// FromPbText unmarshals DeviceBmpServerV4 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceBmpServerV4 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceBmpServerV4 from JSON text
	FromJson(value string) error
}

func (obj *deviceBmpServerV4) Marshal() marshalDeviceBmpServerV4 {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceBmpServerV4{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceBmpServerV4) Unmarshal() unMarshalDeviceBmpServerV4 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceBmpServerV4{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceBmpServerV4) ToProto() (*otg.DeviceBmpServerV4, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceBmpServerV4) FromProto(msg *otg.DeviceBmpServerV4) (DeviceBmpServerV4, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceBmpServerV4) ToPbText() (string, error) {
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

func (m *unMarshaldeviceBmpServerV4) FromPbText(value string) error {
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

func (m *marshaldeviceBmpServerV4) ToYaml() (string, error) {
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

func (m *unMarshaldeviceBmpServerV4) FromYaml(value string) error {
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

func (m *marshaldeviceBmpServerV4) ToJson() (string, error) {
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

func (m *unMarshaldeviceBmpServerV4) FromJson(value string) error {
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

func (obj *deviceBmpServerV4) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceBmpServerV4) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceBmpServerV4) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceBmpServerV4) Clone() (DeviceBmpServerV4, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceBmpServerV4()
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

func (obj *deviceBmpServerV4) setNil() {
	obj.connectionHolder = nil
	obj.prefixStorageHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceBmpServerV4 is configuration for a BMP Server for a specific IPv4 BMP client.
type DeviceBmpServerV4 interface {
	Validation
	// msg marshals DeviceBmpServerV4 to protobuf object *otg.DeviceBmpServerV4
	// and doesn't set defaults
	msg() *otg.DeviceBmpServerV4
	// setMsg unmarshals DeviceBmpServerV4 from protobuf object *otg.DeviceBmpServerV4
	// and doesn't set defaults
	setMsg(*otg.DeviceBmpServerV4) DeviceBmpServerV4
	// provides marshal interface
	Marshal() marshalDeviceBmpServerV4
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceBmpServerV4
	// validate validates DeviceBmpServerV4
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceBmpServerV4, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in DeviceBmpServerV4.
	Name() string
	// SetName assigns string provided by user to DeviceBmpServerV4
	SetName(value string) DeviceBmpServerV4
	// ClientIp returns string, set in DeviceBmpServerV4.
	ClientIp() string
	// SetClientIp assigns string provided by user to DeviceBmpServerV4
	SetClientIp(value string) DeviceBmpServerV4
	// Connection returns DeviceBmpServerConnection, set in DeviceBmpServerV4.
	// DeviceBmpServerConnection is container of information about whether the BMP Server should operate in passive or active mode and corresponding information depending on the mode.
	Connection() DeviceBmpServerConnection
	// SetConnection assigns DeviceBmpServerConnection provided by user to DeviceBmpServerV4.
	// DeviceBmpServerConnection is container of information about whether the BMP Server should operate in passive or active mode and corresponding information depending on the mode.
	SetConnection(value DeviceBmpServerConnection) DeviceBmpServerV4
	// HasConnection checks if Connection has been set in DeviceBmpServerV4
	HasConnection() bool
	// PrefixStorage returns DeviceBmpServerPrefixStorage, set in DeviceBmpServerV4.
	// DeviceBmpServerPrefixStorage is optional object containing information about whether IPv4 and IPv6 unicast prefixes learned from the BMP client should be stored or not for future retrieval using get_states and  exceptions to the configured choice for the same. If the object is not included, by default IPv4 and IPv6 Unicast Prefixes are not stored by the BMP server and only received metrics  are incremented on receipt of IPv4 / IPv6 Unicast routes via BMP Monitor messages.
	PrefixStorage() DeviceBmpServerPrefixStorage
	// SetPrefixStorage assigns DeviceBmpServerPrefixStorage provided by user to DeviceBmpServerV4.
	// DeviceBmpServerPrefixStorage is optional object containing information about whether IPv4 and IPv6 unicast prefixes learned from the BMP client should be stored or not for future retrieval using get_states and  exceptions to the configured choice for the same. If the object is not included, by default IPv4 and IPv6 Unicast Prefixes are not stored by the BMP server and only received metrics  are incremented on receipt of IPv4 / IPv6 Unicast routes via BMP Monitor messages.
	SetPrefixStorage(value DeviceBmpServerPrefixStorage) DeviceBmpServerV4
	// HasPrefixStorage checks if PrefixStorage has been set in DeviceBmpServerV4
	HasPrefixStorage() bool
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *deviceBmpServerV4) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the DeviceBmpServerV4 object
func (obj *deviceBmpServerV4) SetName(value string) DeviceBmpServerV4 {

	obj.obj.Name = &value
	return obj
}

// The IPv4 address of the BMP client from which connections will be accepted from or initiated with by this BMP Server.
// ClientIp returns a string
func (obj *deviceBmpServerV4) ClientIp() string {

	return *obj.obj.ClientIp

}

// The IPv4 address of the BMP client from which connections will be accepted from or initiated with by this BMP Server.
// SetClientIp sets the string value in the DeviceBmpServerV4 object
func (obj *deviceBmpServerV4) SetClientIp(value string) DeviceBmpServerV4 {

	obj.obj.ClientIp = &value
	return obj
}

// Optional object containing information about whether the BMP Server should operate in passive or active mode and corresponding information depending on the mode.
// Connection returns a DeviceBmpServerConnection
func (obj *deviceBmpServerV4) Connection() DeviceBmpServerConnection {
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
func (obj *deviceBmpServerV4) HasConnection() bool {
	return obj.obj.Connection != nil
}

// Optional object containing information about whether the BMP Server should operate in passive or active mode and corresponding information depending on the mode.
// SetConnection sets the DeviceBmpServerConnection value in the DeviceBmpServerV4 object
func (obj *deviceBmpServerV4) SetConnection(value DeviceBmpServerConnection) DeviceBmpServerV4 {

	obj.connectionHolder = nil
	obj.obj.Connection = value.msg()

	return obj
}

// Optional object containing information about whether IPv4 and IPv6 unicast prefixes learned from the BMP client should be stored or not for future retrieval using get_states and  exceptions to the configured choice for the same.  If the object is not included, by default IPv4 and IPv6 unicast prefixes are not stored by the BMP server and only received metrics are incremented on receipt of IPv4 / IPv6 unicast prefixes within  BMP Monitor messages.
// PrefixStorage returns a DeviceBmpServerPrefixStorage
func (obj *deviceBmpServerV4) PrefixStorage() DeviceBmpServerPrefixStorage {
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
func (obj *deviceBmpServerV4) HasPrefixStorage() bool {
	return obj.obj.PrefixStorage != nil
}

// Optional object containing information about whether IPv4 and IPv6 unicast prefixes learned from the BMP client should be stored or not for future retrieval using get_states and  exceptions to the configured choice for the same.  If the object is not included, by default IPv4 and IPv6 unicast prefixes are not stored by the BMP server and only received metrics are incremented on receipt of IPv4 / IPv6 unicast prefixes within  BMP Monitor messages.
// SetPrefixStorage sets the DeviceBmpServerPrefixStorage value in the DeviceBmpServerV4 object
func (obj *deviceBmpServerV4) SetPrefixStorage(value DeviceBmpServerPrefixStorage) DeviceBmpServerV4 {

	obj.prefixStorageHolder = nil
	obj.obj.PrefixStorage = value.msg()

	return obj
}

func (obj *deviceBmpServerV4) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface DeviceBmpServerV4")
	}

	// ClientIp is required
	if obj.obj.ClientIp == nil {
		vObj.validationErrors = append(vObj.validationErrors, "ClientIp is required field on interface DeviceBmpServerV4")
	}
	if obj.obj.ClientIp != nil {

		err := obj.validateIpv4(obj.ClientIp())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on DeviceBmpServerV4.ClientIp"))
		}

	}

	if obj.obj.Connection != nil {

		obj.Connection().validateObj(vObj, set_default)
	}

	if obj.obj.PrefixStorage != nil {

		obj.PrefixStorage().validateObj(vObj, set_default)
	}

}

func (obj *deviceBmpServerV4) setDefault() {

}
