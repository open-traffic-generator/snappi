package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3Advanced *****
type ospfv3Advanced struct {
	validation
	obj          *otg.Ospfv3Advanced
	marshaller   marshalOspfv3Advanced
	unMarshaller unMarshalOspfv3Advanced
}

func NewOspfv3Advanced() Ospfv3Advanced {
	obj := ospfv3Advanced{obj: &otg.Ospfv3Advanced{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3Advanced) msg() *otg.Ospfv3Advanced {
	return obj.obj
}

func (obj *ospfv3Advanced) setMsg(msg *otg.Ospfv3Advanced) Ospfv3Advanced {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3Advanced struct {
	obj *ospfv3Advanced
}

type marshalOspfv3Advanced interface {
	// ToProto marshals Ospfv3Advanced to protobuf object *otg.Ospfv3Advanced
	ToProto() (*otg.Ospfv3Advanced, error)
	// ToPbText marshals Ospfv3Advanced to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3Advanced to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3Advanced to JSON text
	ToJson() (string, error)
}

type unMarshalospfv3Advanced struct {
	obj *ospfv3Advanced
}

type unMarshalOspfv3Advanced interface {
	// FromProto unmarshals Ospfv3Advanced from protobuf object *otg.Ospfv3Advanced
	FromProto(msg *otg.Ospfv3Advanced) (Ospfv3Advanced, error)
	// FromPbText unmarshals Ospfv3Advanced from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3Advanced from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3Advanced from JSON text
	FromJson(value string) error
}

func (obj *ospfv3Advanced) Marshal() marshalOspfv3Advanced {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3Advanced{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3Advanced) Unmarshal() unMarshalOspfv3Advanced {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3Advanced{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3Advanced) ToProto() (*otg.Ospfv3Advanced, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3Advanced) FromProto(msg *otg.Ospfv3Advanced) (Ospfv3Advanced, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3Advanced) ToPbText() (string, error) {
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

func (m *unMarshalospfv3Advanced) FromPbText(value string) error {
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

func (m *marshalospfv3Advanced) ToYaml() (string, error) {
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

func (m *unMarshalospfv3Advanced) FromYaml(value string) error {
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

func (m *marshalospfv3Advanced) ToJson() (string, error) {
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

func (m *unMarshalospfv3Advanced) FromJson(value string) error {
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

func (obj *ospfv3Advanced) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3Advanced) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3Advanced) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3Advanced) Clone() (Ospfv3Advanced, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3Advanced()
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

// Ospfv3Advanced is the OSPFv3 router optional attributes.
type Ospfv3Advanced interface {
	Validation
	// msg marshals Ospfv3Advanced to protobuf object *otg.Ospfv3Advanced
	// and doesn't set defaults
	msg() *otg.Ospfv3Advanced
	// setMsg unmarshals Ospfv3Advanced from protobuf object *otg.Ospfv3Advanced
	// and doesn't set defaults
	setMsg(*otg.Ospfv3Advanced) Ospfv3Advanced
	// provides marshal interface
	Marshal() marshalOspfv3Advanced
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3Advanced
	// validate validates Ospfv3Advanced
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3Advanced, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LsaBBit returns bool, set in Ospfv3Advanced.
	LsaBBit() bool
	// SetLsaBBit assigns bool provided by user to Ospfv3Advanced
	SetLsaBBit(value bool) Ospfv3Advanced
	// HasLsaBBit checks if LsaBBit has been set in Ospfv3Advanced
	HasLsaBBit() bool
	// LsaEBit returns bool, set in Ospfv3Advanced.
	LsaEBit() bool
	// SetLsaEBit assigns bool provided by user to Ospfv3Advanced
	SetLsaEBit(value bool) Ospfv3Advanced
	// HasLsaEBit checks if LsaEBit has been set in Ospfv3Advanced
	HasLsaEBit() bool
}

// Set to indicate that the router acts as an Area Border Router.
// LsaBBit returns a bool
func (obj *ospfv3Advanced) LsaBBit() bool {

	return *obj.obj.LsaBBit

}

// Set to indicate that the router acts as an Area Border Router.
// LsaBBit returns a bool
func (obj *ospfv3Advanced) HasLsaBBit() bool {
	return obj.obj.LsaBBit != nil
}

// Set to indicate that the router acts as an Area Border Router.
// SetLsaBBit sets the bool value in the Ospfv3Advanced object
func (obj *ospfv3Advanced) SetLsaBBit(value bool) Ospfv3Advanced {

	obj.obj.LsaBBit = &value
	return obj
}

// Set to indicate that the router acts as an AS Boundary Router.
// LsaEBit returns a bool
func (obj *ospfv3Advanced) LsaEBit() bool {

	return *obj.obj.LsaEBit

}

// Set to indicate that the router acts as an AS Boundary Router.
// LsaEBit returns a bool
func (obj *ospfv3Advanced) HasLsaEBit() bool {
	return obj.obj.LsaEBit != nil
}

// Set to indicate that the router acts as an AS Boundary Router.
// SetLsaEBit sets the bool value in the Ospfv3Advanced object
func (obj *ospfv3Advanced) SetLsaEBit(value bool) Ospfv3Advanced {

	obj.obj.LsaEBit = &value
	return obj
}

func (obj *ospfv3Advanced) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ospfv3Advanced) setDefault() {
	if obj.obj.LsaBBit == nil {
		obj.SetLsaBBit(false)
	}
	if obj.obj.LsaEBit == nil {
		obj.SetLsaEBit(false)
	}

}
