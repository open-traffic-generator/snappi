package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** CapabilitiesPortType *****
type capabilitiesPortType struct {
	validation
	obj          *otg.CapabilitiesPortType
	marshaller   marshalCapabilitiesPortType
	unMarshaller unMarshalCapabilitiesPortType
}

func NewCapabilitiesPortType() CapabilitiesPortType {
	obj := capabilitiesPortType{obj: &otg.CapabilitiesPortType{}}
	obj.setDefault()
	return &obj
}

func (obj *capabilitiesPortType) msg() *otg.CapabilitiesPortType {
	return obj.obj
}

func (obj *capabilitiesPortType) setMsg(msg *otg.CapabilitiesPortType) CapabilitiesPortType {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalcapabilitiesPortType struct {
	obj *capabilitiesPortType
}

type marshalCapabilitiesPortType interface {
	// ToProto marshals CapabilitiesPortType to protobuf object *otg.CapabilitiesPortType
	ToProto() (*otg.CapabilitiesPortType, error)
	// ToPbText marshals CapabilitiesPortType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals CapabilitiesPortType to YAML text
	ToYaml() (string, error)
	// ToJson marshals CapabilitiesPortType to JSON text
	ToJson() (string, error)
}

type unMarshalcapabilitiesPortType struct {
	obj *capabilitiesPortType
}

type unMarshalCapabilitiesPortType interface {
	// FromProto unmarshals CapabilitiesPortType from protobuf object *otg.CapabilitiesPortType
	FromProto(msg *otg.CapabilitiesPortType) (CapabilitiesPortType, error)
	// FromPbText unmarshals CapabilitiesPortType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals CapabilitiesPortType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals CapabilitiesPortType from JSON text
	FromJson(value string) error
}

func (obj *capabilitiesPortType) Marshal() marshalCapabilitiesPortType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalcapabilitiesPortType{obj: obj}
	}
	return obj.marshaller
}

func (obj *capabilitiesPortType) Unmarshal() unMarshalCapabilitiesPortType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalcapabilitiesPortType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalcapabilitiesPortType) ToProto() (*otg.CapabilitiesPortType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalcapabilitiesPortType) FromProto(msg *otg.CapabilitiesPortType) (CapabilitiesPortType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalcapabilitiesPortType) ToPbText() (string, error) {
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

func (m *unMarshalcapabilitiesPortType) FromPbText(value string) error {
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

func (m *marshalcapabilitiesPortType) ToYaml() (string, error) {
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

func (m *unMarshalcapabilitiesPortType) FromYaml(value string) error {
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

func (m *marshalcapabilitiesPortType) ToJson() (string, error) {
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

func (m *unMarshalcapabilitiesPortType) FromJson(value string) error {
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

func (obj *capabilitiesPortType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *capabilitiesPortType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *capabilitiesPortType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *capabilitiesPortType) Clone() (CapabilitiesPortType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewCapabilitiesPortType()
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

// CapabilitiesPortType is the information regarding the unique PortType for which the included or excluded capabilities are being returned.
type CapabilitiesPortType interface {
	Validation
	// msg marshals CapabilitiesPortType to protobuf object *otg.CapabilitiesPortType
	// and doesn't set defaults
	msg() *otg.CapabilitiesPortType
	// setMsg unmarshals CapabilitiesPortType from protobuf object *otg.CapabilitiesPortType
	// and doesn't set defaults
	setMsg(*otg.CapabilitiesPortType) CapabilitiesPortType
	// provides marshal interface
	Marshal() marshalCapabilitiesPortType
	// provides unmarshal interface
	Unmarshal() unMarshalCapabilitiesPortType
	// validate validates CapabilitiesPortType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (CapabilitiesPortType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Value returns string, set in CapabilitiesPortType.
	Value() string
	// SetValue assigns string provided by user to CapabilitiesPortType
	SetValue(value string) CapabilitiesPortType
	// HasValue checks if Value has been set in CapabilitiesPortType
	HasValue() bool
}

// An unique string in the context of a specific open-traffic-generator that denotes a specific type of group of test ports with same set of capabilties. e.g. 'ixia-c' could be the port_type.value for the reference ixia-c implementation. This name should be distinct and easily distinguishable by user of the test implementation  adhering to this open-traffic-generator model.
// Value returns a string
func (obj *capabilitiesPortType) Value() string {

	return *obj.obj.Value

}

// An unique string in the context of a specific open-traffic-generator that denotes a specific type of group of test ports with same set of capabilties. e.g. 'ixia-c' could be the port_type.value for the reference ixia-c implementation. This name should be distinct and easily distinguishable by user of the test implementation  adhering to this open-traffic-generator model.
// Value returns a string
func (obj *capabilitiesPortType) HasValue() bool {
	return obj.obj.Value != nil
}

// An unique string in the context of a specific open-traffic-generator that denotes a specific type of group of test ports with same set of capabilties. e.g. 'ixia-c' could be the port_type.value for the reference ixia-c implementation. This name should be distinct and easily distinguishable by user of the test implementation  adhering to this open-traffic-generator model.
// SetValue sets the string value in the CapabilitiesPortType object
func (obj *capabilitiesPortType) SetValue(value string) CapabilitiesPortType {

	obj.obj.Value = &value
	return obj
}

func (obj *capabilitiesPortType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *capabilitiesPortType) setDefault() {

}
