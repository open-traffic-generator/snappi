package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceMacsecEthernetInterface *****
type deviceMacsecEthernetInterface struct {
	validation
	obj                *otg.DeviceMacsecEthernetInterface
	marshaller         marshalDeviceMacsecEthernetInterface
	unMarshaller       unMarshalDeviceMacsecEthernetInterface
	secureEntityHolder SecureEntity
}

func NewDeviceMacsecEthernetInterface() DeviceMacsecEthernetInterface {
	obj := deviceMacsecEthernetInterface{obj: &otg.DeviceMacsecEthernetInterface{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceMacsecEthernetInterface) msg() *otg.DeviceMacsecEthernetInterface {
	return obj.obj
}

func (obj *deviceMacsecEthernetInterface) setMsg(msg *otg.DeviceMacsecEthernetInterface) DeviceMacsecEthernetInterface {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceMacsecEthernetInterface struct {
	obj *deviceMacsecEthernetInterface
}

type marshalDeviceMacsecEthernetInterface interface {
	// ToProto marshals DeviceMacsecEthernetInterface to protobuf object *otg.DeviceMacsecEthernetInterface
	ToProto() (*otg.DeviceMacsecEthernetInterface, error)
	// ToPbText marshals DeviceMacsecEthernetInterface to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceMacsecEthernetInterface to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceMacsecEthernetInterface to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceMacsecEthernetInterface struct {
	obj *deviceMacsecEthernetInterface
}

type unMarshalDeviceMacsecEthernetInterface interface {
	// FromProto unmarshals DeviceMacsecEthernetInterface from protobuf object *otg.DeviceMacsecEthernetInterface
	FromProto(msg *otg.DeviceMacsecEthernetInterface) (DeviceMacsecEthernetInterface, error)
	// FromPbText unmarshals DeviceMacsecEthernetInterface from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceMacsecEthernetInterface from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceMacsecEthernetInterface from JSON text
	FromJson(value string) error
}

func (obj *deviceMacsecEthernetInterface) Marshal() marshalDeviceMacsecEthernetInterface {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceMacsecEthernetInterface{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceMacsecEthernetInterface) Unmarshal() unMarshalDeviceMacsecEthernetInterface {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceMacsecEthernetInterface{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceMacsecEthernetInterface) ToProto() (*otg.DeviceMacsecEthernetInterface, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceMacsecEthernetInterface) FromProto(msg *otg.DeviceMacsecEthernetInterface) (DeviceMacsecEthernetInterface, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceMacsecEthernetInterface) ToPbText() (string, error) {
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

func (m *unMarshaldeviceMacsecEthernetInterface) FromPbText(value string) error {
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

func (m *marshaldeviceMacsecEthernetInterface) ToYaml() (string, error) {
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

func (m *unMarshaldeviceMacsecEthernetInterface) FromYaml(value string) error {
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

func (m *marshaldeviceMacsecEthernetInterface) ToJson() (string, error) {
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

func (m *unMarshaldeviceMacsecEthernetInterface) FromJson(value string) error {
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

func (obj *deviceMacsecEthernetInterface) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceMacsecEthernetInterface) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceMacsecEthernetInterface) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceMacsecEthernetInterface) Clone() (DeviceMacsecEthernetInterface, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceMacsecEthernetInterface()
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

func (obj *deviceMacsecEthernetInterface) setNil() {
	obj.secureEntityHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceMacsecEthernetInterface is configuration for single MACsec interface.
type DeviceMacsecEthernetInterface interface {
	Validation
	// msg marshals DeviceMacsecEthernetInterface to protobuf object *otg.DeviceMacsecEthernetInterface
	// and doesn't set defaults
	msg() *otg.DeviceMacsecEthernetInterface
	// setMsg unmarshals DeviceMacsecEthernetInterface from protobuf object *otg.DeviceMacsecEthernetInterface
	// and doesn't set defaults
	setMsg(*otg.DeviceMacsecEthernetInterface) DeviceMacsecEthernetInterface
	// provides marshal interface
	Marshal() marshalDeviceMacsecEthernetInterface
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceMacsecEthernetInterface
	// validate validates DeviceMacsecEthernetInterface
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceMacsecEthernetInterface, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EthName returns string, set in DeviceMacsecEthernetInterface.
	EthName() string
	// SetEthName assigns string provided by user to DeviceMacsecEthernetInterface
	SetEthName(value string) DeviceMacsecEthernetInterface
	// SecureEntity returns SecureEntity, set in DeviceMacsecEthernetInterface.
	// SecureEntity is configuration of a Secure Entity (SecY).
	SecureEntity() SecureEntity
	// SetSecureEntity assigns SecureEntity provided by user to DeviceMacsecEthernetInterface.
	// SecureEntity is configuration of a Secure Entity (SecY).
	SetSecureEntity(value SecureEntity) DeviceMacsecEthernetInterface
	setNil()
}

// The unique name of the Ethernet interface on which MACsec is enabled.
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// EthName returns a string
func (obj *deviceMacsecEthernetInterface) EthName() string {

	return *obj.obj.EthName

}

// The unique name of the Ethernet interface on which MACsec is enabled.
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// SetEthName sets the string value in the DeviceMacsecEthernetInterface object
func (obj *deviceMacsecEthernetInterface) SetEthName(value string) DeviceMacsecEthernetInterface {

	obj.obj.EthName = &value
	return obj
}

// This contains the properties of Secure Entity (SecY).
// SecureEntity returns a SecureEntity
func (obj *deviceMacsecEthernetInterface) SecureEntity() SecureEntity {
	if obj.obj.SecureEntity == nil {
		obj.obj.SecureEntity = NewSecureEntity().msg()
	}
	if obj.secureEntityHolder == nil {
		obj.secureEntityHolder = &secureEntity{obj: obj.obj.SecureEntity}
	}
	return obj.secureEntityHolder
}

// This contains the properties of Secure Entity (SecY).
// SetSecureEntity sets the SecureEntity value in the DeviceMacsecEthernetInterface object
func (obj *deviceMacsecEthernetInterface) SetSecureEntity(value SecureEntity) DeviceMacsecEthernetInterface {

	obj.secureEntityHolder = nil
	obj.obj.SecureEntity = value.msg()

	return obj
}

func (obj *deviceMacsecEthernetInterface) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// EthName is required
	if obj.obj.EthName == nil {
		vObj.validationErrors = append(vObj.validationErrors, "EthName is required field on interface DeviceMacsecEthernetInterface")
	}

	// SecureEntity is required
	if obj.obj.SecureEntity == nil {
		vObj.validationErrors = append(vObj.validationErrors, "SecureEntity is required field on interface DeviceMacsecEthernetInterface")
	}

	if obj.obj.SecureEntity != nil {

		obj.SecureEntity().validateObj(vObj, set_default)
	}

}

func (obj *deviceMacsecEthernetInterface) setDefault() {

}
