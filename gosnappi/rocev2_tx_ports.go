package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2TxPorts *****
type rocev2TxPorts struct {
	validation
	obj                *otg.Rocev2TxPorts
	marshaller         marshalRocev2TxPorts
	unMarshaller       unMarshalRocev2TxPorts
	transmitTypeHolder Rocev2TransmitType
}

func NewRocev2TxPorts() Rocev2TxPorts {
	obj := rocev2TxPorts{obj: &otg.Rocev2TxPorts{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2TxPorts) msg() *otg.Rocev2TxPorts {
	return obj.obj
}

func (obj *rocev2TxPorts) setMsg(msg *otg.Rocev2TxPorts) Rocev2TxPorts {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2TxPorts struct {
	obj *rocev2TxPorts
}

type marshalRocev2TxPorts interface {
	// ToProto marshals Rocev2TxPorts to protobuf object *otg.Rocev2TxPorts
	ToProto() (*otg.Rocev2TxPorts, error)
	// ToPbText marshals Rocev2TxPorts to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2TxPorts to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2TxPorts to JSON text
	ToJson() (string, error)
}

type unMarshalrocev2TxPorts struct {
	obj *rocev2TxPorts
}

type unMarshalRocev2TxPorts interface {
	// FromProto unmarshals Rocev2TxPorts from protobuf object *otg.Rocev2TxPorts
	FromProto(msg *otg.Rocev2TxPorts) (Rocev2TxPorts, error)
	// FromPbText unmarshals Rocev2TxPorts from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2TxPorts from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2TxPorts from JSON text
	FromJson(value string) error
}

func (obj *rocev2TxPorts) Marshal() marshalRocev2TxPorts {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2TxPorts{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2TxPorts) Unmarshal() unMarshalRocev2TxPorts {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2TxPorts{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2TxPorts) ToProto() (*otg.Rocev2TxPorts, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2TxPorts) FromProto(msg *otg.Rocev2TxPorts) (Rocev2TxPorts, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2TxPorts) ToPbText() (string, error) {
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

func (m *unMarshalrocev2TxPorts) FromPbText(value string) error {
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

func (m *marshalrocev2TxPorts) ToYaml() (string, error) {
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

func (m *unMarshalrocev2TxPorts) FromYaml(value string) error {
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

func (m *marshalrocev2TxPorts) ToJson() (string, error) {
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

func (m *unMarshalrocev2TxPorts) FromJson(value string) error {
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

func (obj *rocev2TxPorts) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2TxPorts) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2TxPorts) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2TxPorts) Clone() (Rocev2TxPorts, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2TxPorts()
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

func (obj *rocev2TxPorts) setNil() {
	obj.transmitTypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Rocev2TxPorts is defines the transmit (TX) port settings for RoCEv2 traffic, including the port name and transmission type configuration.
type Rocev2TxPorts interface {
	Validation
	// msg marshals Rocev2TxPorts to protobuf object *otg.Rocev2TxPorts
	// and doesn't set defaults
	msg() *otg.Rocev2TxPorts
	// setMsg unmarshals Rocev2TxPorts from protobuf object *otg.Rocev2TxPorts
	// and doesn't set defaults
	setMsg(*otg.Rocev2TxPorts) Rocev2TxPorts
	// provides marshal interface
	Marshal() marshalRocev2TxPorts
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2TxPorts
	// validate validates Rocev2TxPorts
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2TxPorts, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PortName returns string, set in Rocev2TxPorts.
	PortName() string
	// SetPortName assigns string provided by user to Rocev2TxPorts
	SetPortName(value string) Rocev2TxPorts
	// HasPortName checks if PortName has been set in Rocev2TxPorts
	HasPortName() bool
	// TransmitType returns Rocev2TransmitType, set in Rocev2TxPorts.
	// Rocev2TransmitType is roCEv2 flows can be configured to run in continuous mode or fixed iteration.
	TransmitType() Rocev2TransmitType
	// SetTransmitType assigns Rocev2TransmitType provided by user to Rocev2TxPorts.
	// Rocev2TransmitType is roCEv2 flows can be configured to run in continuous mode or fixed iteration.
	SetTransmitType(value Rocev2TransmitType) Rocev2TxPorts
	// HasTransmitType checks if TransmitType has been set in Rocev2TxPorts
	HasTransmitType() bool
	setNil()
}

// The name of port for which this settings will be applied to.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// PortName returns a string
func (obj *rocev2TxPorts) PortName() string {

	return *obj.obj.PortName

}

// The name of port for which this settings will be applied to.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// PortName returns a string
func (obj *rocev2TxPorts) HasPortName() bool {
	return obj.obj.PortName != nil
}

// The name of port for which this settings will be applied to.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// SetPortName sets the string value in the Rocev2TxPorts object
func (obj *rocev2TxPorts) SetPortName(value string) Rocev2TxPorts {

	obj.obj.PortName = &value
	return obj
}

// description is TBD
// TransmitType returns a Rocev2TransmitType
func (obj *rocev2TxPorts) TransmitType() Rocev2TransmitType {
	if obj.obj.TransmitType == nil {
		obj.obj.TransmitType = NewRocev2TransmitType().msg()
	}
	if obj.transmitTypeHolder == nil {
		obj.transmitTypeHolder = &rocev2TransmitType{obj: obj.obj.TransmitType}
	}
	return obj.transmitTypeHolder
}

// description is TBD
// TransmitType returns a Rocev2TransmitType
func (obj *rocev2TxPorts) HasTransmitType() bool {
	return obj.obj.TransmitType != nil
}

// description is TBD
// SetTransmitType sets the Rocev2TransmitType value in the Rocev2TxPorts object
func (obj *rocev2TxPorts) SetTransmitType(value Rocev2TransmitType) Rocev2TxPorts {

	obj.transmitTypeHolder = nil
	obj.obj.TransmitType = value.msg()

	return obj
}

func (obj *rocev2TxPorts) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.TransmitType != nil {

		obj.TransmitType().validateObj(vObj, set_default)
	}

}

func (obj *rocev2TxPorts) setDefault() {

}
