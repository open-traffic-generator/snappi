package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceBmpServerPassiveConnection *****
type deviceBmpServerPassiveConnection struct {
	validation
	obj          *otg.DeviceBmpServerPassiveConnection
	marshaller   marshalDeviceBmpServerPassiveConnection
	unMarshaller unMarshalDeviceBmpServerPassiveConnection
}

func NewDeviceBmpServerPassiveConnection() DeviceBmpServerPassiveConnection {
	obj := deviceBmpServerPassiveConnection{obj: &otg.DeviceBmpServerPassiveConnection{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceBmpServerPassiveConnection) msg() *otg.DeviceBmpServerPassiveConnection {
	return obj.obj
}

func (obj *deviceBmpServerPassiveConnection) setMsg(msg *otg.DeviceBmpServerPassiveConnection) DeviceBmpServerPassiveConnection {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceBmpServerPassiveConnection struct {
	obj *deviceBmpServerPassiveConnection
}

type marshalDeviceBmpServerPassiveConnection interface {
	// ToProto marshals DeviceBmpServerPassiveConnection to protobuf object *otg.DeviceBmpServerPassiveConnection
	ToProto() (*otg.DeviceBmpServerPassiveConnection, error)
	// ToPbText marshals DeviceBmpServerPassiveConnection to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceBmpServerPassiveConnection to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceBmpServerPassiveConnection to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceBmpServerPassiveConnection struct {
	obj *deviceBmpServerPassiveConnection
}

type unMarshalDeviceBmpServerPassiveConnection interface {
	// FromProto unmarshals DeviceBmpServerPassiveConnection from protobuf object *otg.DeviceBmpServerPassiveConnection
	FromProto(msg *otg.DeviceBmpServerPassiveConnection) (DeviceBmpServerPassiveConnection, error)
	// FromPbText unmarshals DeviceBmpServerPassiveConnection from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceBmpServerPassiveConnection from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceBmpServerPassiveConnection from JSON text
	FromJson(value string) error
}

func (obj *deviceBmpServerPassiveConnection) Marshal() marshalDeviceBmpServerPassiveConnection {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceBmpServerPassiveConnection{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceBmpServerPassiveConnection) Unmarshal() unMarshalDeviceBmpServerPassiveConnection {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceBmpServerPassiveConnection{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceBmpServerPassiveConnection) ToProto() (*otg.DeviceBmpServerPassiveConnection, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceBmpServerPassiveConnection) FromProto(msg *otg.DeviceBmpServerPassiveConnection) (DeviceBmpServerPassiveConnection, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceBmpServerPassiveConnection) ToPbText() (string, error) {
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

func (m *unMarshaldeviceBmpServerPassiveConnection) FromPbText(value string) error {
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

func (m *marshaldeviceBmpServerPassiveConnection) ToYaml() (string, error) {
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

func (m *unMarshaldeviceBmpServerPassiveConnection) FromYaml(value string) error {
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

func (m *marshaldeviceBmpServerPassiveConnection) ToJson() (string, error) {
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

func (m *unMarshaldeviceBmpServerPassiveConnection) FromJson(value string) error {
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

func (obj *deviceBmpServerPassiveConnection) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceBmpServerPassiveConnection) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceBmpServerPassiveConnection) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceBmpServerPassiveConnection) Clone() (DeviceBmpServerPassiveConnection, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceBmpServerPassiveConnection()
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

// DeviceBmpServerPassiveConnection is container of information when BMP Server is configured in passive mode. This means that BMP Server will not initiate the TCP connection but wait for the BMP client it is configured to accept connection from  to initiate the connection.  Note that in this case it is required to configure the BMP client in active mode otherwise BMP connection will not be intiated by either end.
type DeviceBmpServerPassiveConnection interface {
	Validation
	// msg marshals DeviceBmpServerPassiveConnection to protobuf object *otg.DeviceBmpServerPassiveConnection
	// and doesn't set defaults
	msg() *otg.DeviceBmpServerPassiveConnection
	// setMsg unmarshals DeviceBmpServerPassiveConnection from protobuf object *otg.DeviceBmpServerPassiveConnection
	// and doesn't set defaults
	setMsg(*otg.DeviceBmpServerPassiveConnection) DeviceBmpServerPassiveConnection
	// provides marshal interface
	Marshal() marshalDeviceBmpServerPassiveConnection
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceBmpServerPassiveConnection
	// validate validates DeviceBmpServerPassiveConnection
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceBmpServerPassiveConnection, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ListenPort returns uint32, set in DeviceBmpServerPassiveConnection.
	ListenPort() uint32
	// SetListenPort assigns uint32 provided by user to DeviceBmpServerPassiveConnection
	SetListenPort(value uint32) DeviceBmpServerPassiveConnection
	// HasListenPort checks if ListenPort has been set in DeviceBmpServerPassiveConnection
	HasListenPort() bool
}

// The TCP port number on which to listen for TCP connections from the remote BMP client router.
// ListenPort returns a uint32
func (obj *deviceBmpServerPassiveConnection) ListenPort() uint32 {

	return *obj.obj.ListenPort

}

// The TCP port number on which to listen for TCP connections from the remote BMP client router.
// ListenPort returns a uint32
func (obj *deviceBmpServerPassiveConnection) HasListenPort() bool {
	return obj.obj.ListenPort != nil
}

// The TCP port number on which to listen for TCP connections from the remote BMP client router.
// SetListenPort sets the uint32 value in the DeviceBmpServerPassiveConnection object
func (obj *deviceBmpServerPassiveConnection) SetListenPort(value uint32) DeviceBmpServerPassiveConnection {

	obj.obj.ListenPort = &value
	return obj
}

func (obj *deviceBmpServerPassiveConnection) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ListenPort != nil {

		if *obj.obj.ListenPort > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= DeviceBmpServerPassiveConnection.ListenPort <= 65535 but Got %d", *obj.obj.ListenPort))
		}

	}

}

func (obj *deviceBmpServerPassiveConnection) setDefault() {
	if obj.obj.ListenPort == nil {
		obj.SetListenPort(11019)
	}

}
