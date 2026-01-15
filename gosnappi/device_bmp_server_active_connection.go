package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceBmpServerActiveConnection *****
type deviceBmpServerActiveConnection struct {
	validation
	obj          *otg.DeviceBmpServerActiveConnection
	marshaller   marshalDeviceBmpServerActiveConnection
	unMarshaller unMarshalDeviceBmpServerActiveConnection
}

func NewDeviceBmpServerActiveConnection() DeviceBmpServerActiveConnection {
	obj := deviceBmpServerActiveConnection{obj: &otg.DeviceBmpServerActiveConnection{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceBmpServerActiveConnection) msg() *otg.DeviceBmpServerActiveConnection {
	return obj.obj
}

func (obj *deviceBmpServerActiveConnection) setMsg(msg *otg.DeviceBmpServerActiveConnection) DeviceBmpServerActiveConnection {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceBmpServerActiveConnection struct {
	obj *deviceBmpServerActiveConnection
}

type marshalDeviceBmpServerActiveConnection interface {
	// ToProto marshals DeviceBmpServerActiveConnection to protobuf object *otg.DeviceBmpServerActiveConnection
	ToProto() (*otg.DeviceBmpServerActiveConnection, error)
	// ToPbText marshals DeviceBmpServerActiveConnection to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceBmpServerActiveConnection to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceBmpServerActiveConnection to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceBmpServerActiveConnection struct {
	obj *deviceBmpServerActiveConnection
}

type unMarshalDeviceBmpServerActiveConnection interface {
	// FromProto unmarshals DeviceBmpServerActiveConnection from protobuf object *otg.DeviceBmpServerActiveConnection
	FromProto(msg *otg.DeviceBmpServerActiveConnection) (DeviceBmpServerActiveConnection, error)
	// FromPbText unmarshals DeviceBmpServerActiveConnection from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceBmpServerActiveConnection from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceBmpServerActiveConnection from JSON text
	FromJson(value string) error
}

func (obj *deviceBmpServerActiveConnection) Marshal() marshalDeviceBmpServerActiveConnection {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceBmpServerActiveConnection{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceBmpServerActiveConnection) Unmarshal() unMarshalDeviceBmpServerActiveConnection {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceBmpServerActiveConnection{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceBmpServerActiveConnection) ToProto() (*otg.DeviceBmpServerActiveConnection, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceBmpServerActiveConnection) FromProto(msg *otg.DeviceBmpServerActiveConnection) (DeviceBmpServerActiveConnection, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceBmpServerActiveConnection) ToPbText() (string, error) {
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

func (m *unMarshaldeviceBmpServerActiveConnection) FromPbText(value string) error {
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

func (m *marshaldeviceBmpServerActiveConnection) ToYaml() (string, error) {
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

func (m *unMarshaldeviceBmpServerActiveConnection) FromYaml(value string) error {
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

func (m *marshaldeviceBmpServerActiveConnection) ToJson() (string, error) {
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

func (m *unMarshaldeviceBmpServerActiveConnection) FromJson(value string) error {
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

func (obj *deviceBmpServerActiveConnection) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceBmpServerActiveConnection) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceBmpServerActiveConnection) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceBmpServerActiveConnection) Clone() (DeviceBmpServerActiveConnection, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceBmpServerActiveConnection()
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

// DeviceBmpServerActiveConnection is container of information when BMP Server is configured in active mode. This means that BMP Server will initiate the TCP connection to the remote BMP client .         Note that in this case it is required to configure the BMP client in passive mode for the BMP session to not be rejected from both ends.
type DeviceBmpServerActiveConnection interface {
	Validation
	// msg marshals DeviceBmpServerActiveConnection to protobuf object *otg.DeviceBmpServerActiveConnection
	// and doesn't set defaults
	msg() *otg.DeviceBmpServerActiveConnection
	// setMsg unmarshals DeviceBmpServerActiveConnection from protobuf object *otg.DeviceBmpServerActiveConnection
	// and doesn't set defaults
	setMsg(*otg.DeviceBmpServerActiveConnection) DeviceBmpServerActiveConnection
	// provides marshal interface
	Marshal() marshalDeviceBmpServerActiveConnection
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceBmpServerActiveConnection
	// validate validates DeviceBmpServerActiveConnection
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceBmpServerActiveConnection, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RemotePort returns uint32, set in DeviceBmpServerActiveConnection.
	RemotePort() uint32
	// SetRemotePort assigns uint32 provided by user to DeviceBmpServerActiveConnection
	SetRemotePort(value uint32) DeviceBmpServerActiveConnection
	// HasRemotePort checks if RemotePort has been set in DeviceBmpServerActiveConnection
	HasRemotePort() bool
}

// The TCP port number on which to initiate BMP connection to the remote BMP client router.
// RemotePort returns a uint32
func (obj *deviceBmpServerActiveConnection) RemotePort() uint32 {

	return *obj.obj.RemotePort

}

// The TCP port number on which to initiate BMP connection to the remote BMP client router.
// RemotePort returns a uint32
func (obj *deviceBmpServerActiveConnection) HasRemotePort() bool {
	return obj.obj.RemotePort != nil
}

// The TCP port number on which to initiate BMP connection to the remote BMP client router.
// SetRemotePort sets the uint32 value in the DeviceBmpServerActiveConnection object
func (obj *deviceBmpServerActiveConnection) SetRemotePort(value uint32) DeviceBmpServerActiveConnection {

	obj.obj.RemotePort = &value
	return obj
}

func (obj *deviceBmpServerActiveConnection) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RemotePort != nil {

		if *obj.obj.RemotePort > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= DeviceBmpServerActiveConnection.RemotePort <= 65535 but Got %d", *obj.obj.RemotePort))
		}

	}

}

func (obj *deviceBmpServerActiveConnection) setDefault() {
	if obj.obj.RemotePort == nil {
		obj.SetRemotePort(11019)
	}

}
