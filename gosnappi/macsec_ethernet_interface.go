package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecEthernetInterface *****
type macsecEthernetInterface struct {
	validation
	obj          *otg.MacsecEthernetInterface
	marshaller   marshalMacsecEthernetInterface
	unMarshaller unMarshalMacsecEthernetInterface
	secyHolder   MacsecSecY
}

func NewMacsecEthernetInterface() MacsecEthernetInterface {
	obj := macsecEthernetInterface{obj: &otg.MacsecEthernetInterface{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecEthernetInterface) msg() *otg.MacsecEthernetInterface {
	return obj.obj
}

func (obj *macsecEthernetInterface) setMsg(msg *otg.MacsecEthernetInterface) MacsecEthernetInterface {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecEthernetInterface struct {
	obj *macsecEthernetInterface
}

type marshalMacsecEthernetInterface interface {
	// ToProto marshals MacsecEthernetInterface to protobuf object *otg.MacsecEthernetInterface
	ToProto() (*otg.MacsecEthernetInterface, error)
	// ToPbText marshals MacsecEthernetInterface to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecEthernetInterface to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecEthernetInterface to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecEthernetInterface struct {
	obj *macsecEthernetInterface
}

type unMarshalMacsecEthernetInterface interface {
	// FromProto unmarshals MacsecEthernetInterface from protobuf object *otg.MacsecEthernetInterface
	FromProto(msg *otg.MacsecEthernetInterface) (MacsecEthernetInterface, error)
	// FromPbText unmarshals MacsecEthernetInterface from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecEthernetInterface from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecEthernetInterface from JSON text
	FromJson(value string) error
}

func (obj *macsecEthernetInterface) Marshal() marshalMacsecEthernetInterface {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecEthernetInterface{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecEthernetInterface) Unmarshal() unMarshalMacsecEthernetInterface {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecEthernetInterface{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecEthernetInterface) ToProto() (*otg.MacsecEthernetInterface, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecEthernetInterface) FromProto(msg *otg.MacsecEthernetInterface) (MacsecEthernetInterface, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecEthernetInterface) ToPbText() (string, error) {
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

func (m *unMarshalmacsecEthernetInterface) FromPbText(value string) error {
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

func (m *marshalmacsecEthernetInterface) ToYaml() (string, error) {
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

func (m *unMarshalmacsecEthernetInterface) FromYaml(value string) error {
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

func (m *marshalmacsecEthernetInterface) ToJson() (string, error) {
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

func (m *unMarshalmacsecEthernetInterface) FromJson(value string) error {
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

func (obj *macsecEthernetInterface) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecEthernetInterface) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecEthernetInterface) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecEthernetInterface) Clone() (MacsecEthernetInterface, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecEthernetInterface()
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

func (obj *macsecEthernetInterface) setNil() {
	obj.secyHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecEthernetInterface is configuration for single MACsec interface.
type MacsecEthernetInterface interface {
	Validation
	// msg marshals MacsecEthernetInterface to protobuf object *otg.MacsecEthernetInterface
	// and doesn't set defaults
	msg() *otg.MacsecEthernetInterface
	// setMsg unmarshals MacsecEthernetInterface from protobuf object *otg.MacsecEthernetInterface
	// and doesn't set defaults
	setMsg(*otg.MacsecEthernetInterface) MacsecEthernetInterface
	// provides marshal interface
	Marshal() marshalMacsecEthernetInterface
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecEthernetInterface
	// validate validates MacsecEthernetInterface
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecEthernetInterface, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EthName returns string, set in MacsecEthernetInterface.
	EthName() string
	// SetEthName assigns string provided by user to MacsecEthernetInterface
	SetEthName(value string) MacsecEthernetInterface
	// Secy returns MacsecSecY, set in MacsecEthernetInterface.
	// MacsecSecY is configuration for a Secure Entity (SecY).
	Secy() MacsecSecY
	// SetSecy assigns MacsecSecY provided by user to MacsecEthernetInterface.
	// MacsecSecY is configuration for a Secure Entity (SecY).
	SetSecy(value MacsecSecY) MacsecEthernetInterface
	setNil()
}

// The unique name of the Ethernet interface on which MACsec is enabled.
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// EthName returns a string
func (obj *macsecEthernetInterface) EthName() string {

	return *obj.obj.EthName

}

// The unique name of the Ethernet interface on which MACsec is enabled.
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// SetEthName sets the string value in the MacsecEthernetInterface object
func (obj *macsecEthernetInterface) SetEthName(value string) MacsecEthernetInterface {

	obj.obj.EthName = &value
	return obj
}

// This contains the properties of Secure Entity (SecY).
// Secy returns a MacsecSecY
func (obj *macsecEthernetInterface) Secy() MacsecSecY {
	if obj.obj.Secy == nil {
		obj.obj.Secy = NewMacsecSecY().msg()
	}
	if obj.secyHolder == nil {
		obj.secyHolder = &macsecSecY{obj: obj.obj.Secy}
	}
	return obj.secyHolder
}

// This contains the properties of Secure Entity (SecY).
// SetSecy sets the MacsecSecY value in the MacsecEthernetInterface object
func (obj *macsecEthernetInterface) SetSecy(value MacsecSecY) MacsecEthernetInterface {

	obj.secyHolder = nil
	obj.obj.Secy = value.msg()

	return obj
}

func (obj *macsecEthernetInterface) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// EthName is required
	if obj.obj.EthName == nil {
		vObj.validationErrors = append(vObj.validationErrors, "EthName is required field on interface MacsecEthernetInterface")
	}

	// Secy is required
	if obj.obj.Secy == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Secy is required field on interface MacsecEthernetInterface")
	}

	if obj.obj.Secy != nil {

		obj.Secy().validateObj(vObj, set_default)
	}

}

func (obj *macsecEthernetInterface) setDefault() {

}
