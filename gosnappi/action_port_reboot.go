package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionPortReboot *****
type actionPortReboot struct {
	validation
	obj          *otg.ActionPortReboot
	marshaller   marshalActionPortReboot
	unMarshaller unMarshalActionPortReboot
}

func NewActionPortReboot() ActionPortReboot {
	obj := actionPortReboot{obj: &otg.ActionPortReboot{}}
	obj.setDefault()
	return &obj
}

func (obj *actionPortReboot) msg() *otg.ActionPortReboot {
	return obj.obj
}

func (obj *actionPortReboot) setMsg(msg *otg.ActionPortReboot) ActionPortReboot {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionPortReboot struct {
	obj *actionPortReboot
}

type marshalActionPortReboot interface {
	// ToProto marshals ActionPortReboot to protobuf object *otg.ActionPortReboot
	ToProto() (*otg.ActionPortReboot, error)
	// ToPbText marshals ActionPortReboot to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionPortReboot to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionPortReboot to JSON text
	ToJson() (string, error)
}

type unMarshalactionPortReboot struct {
	obj *actionPortReboot
}

type unMarshalActionPortReboot interface {
	// FromProto unmarshals ActionPortReboot from protobuf object *otg.ActionPortReboot
	FromProto(msg *otg.ActionPortReboot) (ActionPortReboot, error)
	// FromPbText unmarshals ActionPortReboot from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionPortReboot from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionPortReboot from JSON text
	FromJson(value string) error
}

func (obj *actionPortReboot) Marshal() marshalActionPortReboot {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionPortReboot{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionPortReboot) Unmarshal() unMarshalActionPortReboot {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionPortReboot{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionPortReboot) ToProto() (*otg.ActionPortReboot, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionPortReboot) FromProto(msg *otg.ActionPortReboot) (ActionPortReboot, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionPortReboot) ToPbText() (string, error) {
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

func (m *unMarshalactionPortReboot) FromPbText(value string) error {
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

func (m *marshalactionPortReboot) ToYaml() (string, error) {
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

func (m *unMarshalactionPortReboot) FromYaml(value string) error {
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

func (m *marshalactionPortReboot) ToJson() (string, error) {
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

func (m *unMarshalactionPortReboot) FromJson(value string) error {
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

func (obj *actionPortReboot) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionPortReboot) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionPortReboot) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionPortReboot) Clone() (ActionPortReboot, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionPortReboot()
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

// ActionPortReboot is resets the configured ports to initialized state.
type ActionPortReboot interface {
	Validation
	// msg marshals ActionPortReboot to protobuf object *otg.ActionPortReboot
	// and doesn't set defaults
	msg() *otg.ActionPortReboot
	// setMsg unmarshals ActionPortReboot from protobuf object *otg.ActionPortReboot
	// and doesn't set defaults
	setMsg(*otg.ActionPortReboot) ActionPortReboot
	// provides marshal interface
	Marshal() marshalActionPortReboot
	// provides unmarshal interface
	Unmarshal() unMarshalActionPortReboot
	// validate validates ActionPortReboot
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionPortReboot, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PortNames returns []string, set in ActionPortReboot.
	PortNames() []string
	// SetPortNames assigns []string provided by user to ActionPortReboot
	SetPortNames(value []string) ActionPortReboot
}

// The names of target ports. An empty or null list will target all ports.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// PortNames returns a []string
func (obj *actionPortReboot) PortNames() []string {
	if obj.obj.PortNames == nil {
		obj.obj.PortNames = make([]string, 0)
	}
	return obj.obj.PortNames
}

// The names of target ports. An empty or null list will target all ports.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// SetPortNames sets the []string value in the ActionPortReboot object
func (obj *actionPortReboot) SetPortNames(value []string) ActionPortReboot {

	if obj.obj.PortNames == nil {
		obj.obj.PortNames = make([]string, 0)
	}
	obj.obj.PortNames = value

	return obj
}

func (obj *actionPortReboot) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *actionPortReboot) setDefault() {

}
