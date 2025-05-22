package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3Capabilities *****
type ospfv3Capabilities struct {
	validation
	obj          *otg.Ospfv3Capabilities
	marshaller   marshalOspfv3Capabilities
	unMarshaller unMarshalOspfv3Capabilities
}

func NewOspfv3Capabilities() Ospfv3Capabilities {
	obj := ospfv3Capabilities{obj: &otg.Ospfv3Capabilities{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3Capabilities) msg() *otg.Ospfv3Capabilities {
	return obj.obj
}

func (obj *ospfv3Capabilities) setMsg(msg *otg.Ospfv3Capabilities) Ospfv3Capabilities {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3Capabilities struct {
	obj *ospfv3Capabilities
}

type marshalOspfv3Capabilities interface {
	// ToProto marshals Ospfv3Capabilities to protobuf object *otg.Ospfv3Capabilities
	ToProto() (*otg.Ospfv3Capabilities, error)
	// ToPbText marshals Ospfv3Capabilities to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3Capabilities to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3Capabilities to JSON text
	ToJson() (string, error)
}

type unMarshalospfv3Capabilities struct {
	obj *ospfv3Capabilities
}

type unMarshalOspfv3Capabilities interface {
	// FromProto unmarshals Ospfv3Capabilities from protobuf object *otg.Ospfv3Capabilities
	FromProto(msg *otg.Ospfv3Capabilities) (Ospfv3Capabilities, error)
	// FromPbText unmarshals Ospfv3Capabilities from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3Capabilities from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3Capabilities from JSON text
	FromJson(value string) error
}

func (obj *ospfv3Capabilities) Marshal() marshalOspfv3Capabilities {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3Capabilities{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3Capabilities) Unmarshal() unMarshalOspfv3Capabilities {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3Capabilities{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3Capabilities) ToProto() (*otg.Ospfv3Capabilities, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3Capabilities) FromProto(msg *otg.Ospfv3Capabilities) (Ospfv3Capabilities, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3Capabilities) ToPbText() (string, error) {
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

func (m *unMarshalospfv3Capabilities) FromPbText(value string) error {
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

func (m *marshalospfv3Capabilities) ToYaml() (string, error) {
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

func (m *unMarshalospfv3Capabilities) FromYaml(value string) error {
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

func (m *marshalospfv3Capabilities) ToJson() (string, error) {
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

func (m *unMarshalospfv3Capabilities) FromJson(value string) error {
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

func (obj *ospfv3Capabilities) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3Capabilities) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3Capabilities) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3Capabilities) Clone() (Ospfv3Capabilities, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3Capabilities()
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

// Ospfv3Capabilities is the OSPFv3 router optional attributes.
type Ospfv3Capabilities interface {
	Validation
	// msg marshals Ospfv3Capabilities to protobuf object *otg.Ospfv3Capabilities
	// and doesn't set defaults
	msg() *otg.Ospfv3Capabilities
	// setMsg unmarshals Ospfv3Capabilities from protobuf object *otg.Ospfv3Capabilities
	// and doesn't set defaults
	setMsg(*otg.Ospfv3Capabilities) Ospfv3Capabilities
	// provides marshal interface
	Marshal() marshalOspfv3Capabilities
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3Capabilities
	// validate validates Ospfv3Capabilities
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3Capabilities, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LsaBBit returns bool, set in Ospfv3Capabilities.
	LsaBBit() bool
	// SetLsaBBit assigns bool provided by user to Ospfv3Capabilities
	SetLsaBBit(value bool) Ospfv3Capabilities
	// HasLsaBBit checks if LsaBBit has been set in Ospfv3Capabilities
	HasLsaBBit() bool
	// LsaEBit returns bool, set in Ospfv3Capabilities.
	LsaEBit() bool
	// SetLsaEBit assigns bool provided by user to Ospfv3Capabilities
	SetLsaEBit(value bool) Ospfv3Capabilities
	// HasLsaEBit checks if LsaEBit has been set in Ospfv3Capabilities
	HasLsaEBit() bool
}

// Set to indicate that the router acts as an Area Border Router.
// LsaBBit returns a bool
func (obj *ospfv3Capabilities) LsaBBit() bool {

	return *obj.obj.LsaBBit

}

// Set to indicate that the router acts as an Area Border Router.
// LsaBBit returns a bool
func (obj *ospfv3Capabilities) HasLsaBBit() bool {
	return obj.obj.LsaBBit != nil
}

// Set to indicate that the router acts as an Area Border Router.
// SetLsaBBit sets the bool value in the Ospfv3Capabilities object
func (obj *ospfv3Capabilities) SetLsaBBit(value bool) Ospfv3Capabilities {

	obj.obj.LsaBBit = &value
	return obj
}

// Set to indicate that the router acts as an AS Boundary Router.
// LsaEBit returns a bool
func (obj *ospfv3Capabilities) LsaEBit() bool {

	return *obj.obj.LsaEBit

}

// Set to indicate that the router acts as an AS Boundary Router.
// LsaEBit returns a bool
func (obj *ospfv3Capabilities) HasLsaEBit() bool {
	return obj.obj.LsaEBit != nil
}

// Set to indicate that the router acts as an AS Boundary Router.
// SetLsaEBit sets the bool value in the Ospfv3Capabilities object
func (obj *ospfv3Capabilities) SetLsaEBit(value bool) Ospfv3Capabilities {

	obj.obj.LsaEBit = &value
	return obj
}

func (obj *ospfv3Capabilities) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ospfv3Capabilities) setDefault() {
	if obj.obj.LsaBBit == nil {
		obj.SetLsaBBit(false)
	}
	if obj.obj.LsaEBit == nil {
		obj.SetLsaEBit(false)
	}

}
