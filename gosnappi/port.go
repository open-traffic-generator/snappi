package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Port *****
type port struct {
	validation
	obj          *otg.Port
	marshaller   marshalPort
	unMarshaller unMarshalPort
}

func NewPort() Port {
	obj := port{obj: &otg.Port{}}
	obj.setDefault()
	return &obj
}

func (obj *port) msg() *otg.Port {
	return obj.obj
}

func (obj *port) setMsg(msg *otg.Port) Port {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalport struct {
	obj *port
}

type marshalPort interface {
	// ToProto marshals Port to protobuf object *otg.Port
	ToProto() (*otg.Port, error)
	// ToPbText marshals Port to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Port to YAML text
	ToYaml() (string, error)
	// ToJson marshals Port to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Port to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalport struct {
	obj *port
}

type unMarshalPort interface {
	// FromProto unmarshals Port from protobuf object *otg.Port
	FromProto(msg *otg.Port) (Port, error)
	// FromPbText unmarshals Port from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Port from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Port from JSON text
	FromJson(value string) error
}

func (obj *port) Marshal() marshalPort {
	if obj.marshaller == nil {
		obj.marshaller = &marshalport{obj: obj}
	}
	return obj.marshaller
}

func (obj *port) Unmarshal() unMarshalPort {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalport{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalport) ToProto() (*otg.Port, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalport) FromProto(msg *otg.Port) (Port, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalport) ToPbText() (string, error) {
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

func (m *unMarshalport) FromPbText(value string) error {
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

func (m *marshalport) ToYaml() (string, error) {
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

func (m *unMarshalport) FromYaml(value string) error {
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

func (m *marshalport) ToJsonRaw() (string, error) {
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

func (m *marshalport) ToJson() (string, error) {
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

func (m *unMarshalport) FromJson(value string) error {
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

func (obj *port) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *port) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *port) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *port) Clone() (Port, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPort()
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

// Port is an abstract test port.
type Port interface {
	Validation
	// msg marshals Port to protobuf object *otg.Port
	// and doesn't set defaults
	msg() *otg.Port
	// setMsg unmarshals Port from protobuf object *otg.Port
	// and doesn't set defaults
	setMsg(*otg.Port) Port
	// provides marshal interface
	Marshal() marshalPort
	// provides unmarshal interface
	Unmarshal() unMarshalPort
	// validate validates Port
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Port, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Location returns string, set in Port.
	Location() string
	// SetLocation assigns string provided by user to Port
	SetLocation(value string) Port
	// HasLocation checks if Location has been set in Port
	HasLocation() bool
	// Name returns string, set in Port.
	Name() string
	// SetName assigns string provided by user to Port
	SetName(value string) Port
}

// The location of a test port.  It is the endpoint where packets will emit from.
// Test port locations can be the following:
// - physical appliance with multiple ports
// - physical chassis with multiple cards and ports
// - local interface
// - virtual machine, docker container, kubernetes cluster
//
// The test port location format is implementation specific. Use the /results/capabilities API to determine what formats an  implementation supports for the location property.
// Get the configured location state by using the /results/port API.
// Location returns a string
func (obj *port) Location() string {

	return *obj.obj.Location

}

// The location of a test port.  It is the endpoint where packets will emit from.
// Test port locations can be the following:
// - physical appliance with multiple ports
// - physical chassis with multiple cards and ports
// - local interface
// - virtual machine, docker container, kubernetes cluster
//
// The test port location format is implementation specific. Use the /results/capabilities API to determine what formats an  implementation supports for the location property.
// Get the configured location state by using the /results/port API.
// Location returns a string
func (obj *port) HasLocation() bool {
	return obj.obj.Location != nil
}

// The location of a test port.  It is the endpoint where packets will emit from.
// Test port locations can be the following:
// - physical appliance with multiple ports
// - physical chassis with multiple cards and ports
// - local interface
// - virtual machine, docker container, kubernetes cluster
//
// The test port location format is implementation specific. Use the /results/capabilities API to determine what formats an  implementation supports for the location property.
// Get the configured location state by using the /results/port API.
// SetLocation sets the string value in the Port object
func (obj *port) SetLocation(value string) Port {

	obj.obj.Location = &value
	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *port) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the Port object
func (obj *port) SetName(value string) Port {

	obj.obj.Name = &value
	return obj
}

func (obj *port) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface Port")
	}
}

func (obj *port) setDefault() {

}
