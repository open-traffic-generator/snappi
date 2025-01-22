package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecDeviceEthernetInterface *****
type macsecDeviceEthernetInterface struct {
	validation
	obj          *otg.MacsecDeviceEthernetInterface
	marshaller   marshalMacsecDeviceEthernetInterface
	unMarshaller unMarshalMacsecDeviceEthernetInterface
	secyHolder   Macsec
}

func NewMacsecDeviceEthernetInterface() MacsecDeviceEthernetInterface {
	obj := macsecDeviceEthernetInterface{obj: &otg.MacsecDeviceEthernetInterface{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecDeviceEthernetInterface) msg() *otg.MacsecDeviceEthernetInterface {
	return obj.obj
}

func (obj *macsecDeviceEthernetInterface) setMsg(msg *otg.MacsecDeviceEthernetInterface) MacsecDeviceEthernetInterface {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecDeviceEthernetInterface struct {
	obj *macsecDeviceEthernetInterface
}

type marshalMacsecDeviceEthernetInterface interface {
	// ToProto marshals MacsecDeviceEthernetInterface to protobuf object *otg.MacsecDeviceEthernetInterface
	ToProto() (*otg.MacsecDeviceEthernetInterface, error)
	// ToPbText marshals MacsecDeviceEthernetInterface to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecDeviceEthernetInterface to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecDeviceEthernetInterface to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecDeviceEthernetInterface struct {
	obj *macsecDeviceEthernetInterface
}

type unMarshalMacsecDeviceEthernetInterface interface {
	// FromProto unmarshals MacsecDeviceEthernetInterface from protobuf object *otg.MacsecDeviceEthernetInterface
	FromProto(msg *otg.MacsecDeviceEthernetInterface) (MacsecDeviceEthernetInterface, error)
	// FromPbText unmarshals MacsecDeviceEthernetInterface from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecDeviceEthernetInterface from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecDeviceEthernetInterface from JSON text
	FromJson(value string) error
}

func (obj *macsecDeviceEthernetInterface) Marshal() marshalMacsecDeviceEthernetInterface {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecDeviceEthernetInterface{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecDeviceEthernetInterface) Unmarshal() unMarshalMacsecDeviceEthernetInterface {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecDeviceEthernetInterface{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecDeviceEthernetInterface) ToProto() (*otg.MacsecDeviceEthernetInterface, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecDeviceEthernetInterface) FromProto(msg *otg.MacsecDeviceEthernetInterface) (MacsecDeviceEthernetInterface, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecDeviceEthernetInterface) ToPbText() (string, error) {
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

func (m *unMarshalmacsecDeviceEthernetInterface) FromPbText(value string) error {
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

func (m *marshalmacsecDeviceEthernetInterface) ToYaml() (string, error) {
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

func (m *unMarshalmacsecDeviceEthernetInterface) FromYaml(value string) error {
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

func (m *marshalmacsecDeviceEthernetInterface) ToJson() (string, error) {
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

func (m *unMarshalmacsecDeviceEthernetInterface) FromJson(value string) error {
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

func (obj *macsecDeviceEthernetInterface) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecDeviceEthernetInterface) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecDeviceEthernetInterface) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecDeviceEthernetInterface) Clone() (MacsecDeviceEthernetInterface, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecDeviceEthernetInterface()
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

func (obj *macsecDeviceEthernetInterface) setNil() {
	obj.secyHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecDeviceEthernetInterface is configuration for single MACsec interface.
type MacsecDeviceEthernetInterface interface {
	Validation
	// msg marshals MacsecDeviceEthernetInterface to protobuf object *otg.MacsecDeviceEthernetInterface
	// and doesn't set defaults
	msg() *otg.MacsecDeviceEthernetInterface
	// setMsg unmarshals MacsecDeviceEthernetInterface from protobuf object *otg.MacsecDeviceEthernetInterface
	// and doesn't set defaults
	setMsg(*otg.MacsecDeviceEthernetInterface) MacsecDeviceEthernetInterface
	// provides marshal interface
	Marshal() marshalMacsecDeviceEthernetInterface
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecDeviceEthernetInterface
	// validate validates MacsecDeviceEthernetInterface
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecDeviceEthernetInterface, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EthName returns string, set in MacsecDeviceEthernetInterface.
	EthName() string
	// SetEthName assigns string provided by user to MacsecDeviceEthernetInterface
	SetEthName(value string) MacsecDeviceEthernetInterface
	// Secy returns Macsec, set in MacsecDeviceEthernetInterface.
	// Macsec is configuration of a Secure Entity (SecY).
	Secy() Macsec
	// SetSecy assigns Macsec provided by user to MacsecDeviceEthernetInterface.
	// Macsec is configuration of a Secure Entity (SecY).
	SetSecy(value Macsec) MacsecDeviceEthernetInterface
	setNil()
}

// The unique name of the Ethernet interface on which MACsec is enabled.
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// EthName returns a string
func (obj *macsecDeviceEthernetInterface) EthName() string {

	return *obj.obj.EthName

}

// The unique name of the Ethernet interface on which MACsec is enabled.
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// SetEthName sets the string value in the MacsecDeviceEthernetInterface object
func (obj *macsecDeviceEthernetInterface) SetEthName(value string) MacsecDeviceEthernetInterface {

	obj.obj.EthName = &value
	return obj
}

// This contains the properties of Secure Entity (SecY).
// Secy returns a Macsec
func (obj *macsecDeviceEthernetInterface) Secy() Macsec {
	if obj.obj.Secy == nil {
		obj.obj.Secy = NewMacsec().msg()
	}
	if obj.secyHolder == nil {
		obj.secyHolder = &macsec{obj: obj.obj.Secy}
	}
	return obj.secyHolder
}

// This contains the properties of Secure Entity (SecY).
// SetSecy sets the Macsec value in the MacsecDeviceEthernetInterface object
func (obj *macsecDeviceEthernetInterface) SetSecy(value Macsec) MacsecDeviceEthernetInterface {

	obj.secyHolder = nil
	obj.obj.Secy = value.msg()

	return obj
}

func (obj *macsecDeviceEthernetInterface) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// EthName is required
	if obj.obj.EthName == nil {
		vObj.validationErrors = append(vObj.validationErrors, "EthName is required field on interface MacsecDeviceEthernetInterface")
	}

	// Secy is required
	if obj.obj.Secy == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Secy is required field on interface MacsecDeviceEthernetInterface")
	}

	if obj.obj.Secy != nil {

		obj.Secy().validateObj(vObj, set_default)
	}

}

func (obj *macsecDeviceEthernetInterface) setDefault() {

}
