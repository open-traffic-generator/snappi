package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Layer1Ieee8023X *****
type layer1Ieee8023X struct {
	validation
	obj          *otg.Layer1Ieee8023X
	marshaller   marshalLayer1Ieee8023X
	unMarshaller unMarshalLayer1Ieee8023X
}

func NewLayer1Ieee8023X() Layer1Ieee8023X {
	obj := layer1Ieee8023X{obj: &otg.Layer1Ieee8023X{}}
	obj.setDefault()
	return &obj
}

func (obj *layer1Ieee8023X) msg() *otg.Layer1Ieee8023X {
	return obj.obj
}

func (obj *layer1Ieee8023X) setMsg(msg *otg.Layer1Ieee8023X) Layer1Ieee8023X {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallayer1Ieee8023X struct {
	obj *layer1Ieee8023X
}

type marshalLayer1Ieee8023X interface {
	// ToProto marshals Layer1Ieee8023X to protobuf object *otg.Layer1Ieee8023X
	ToProto() (*otg.Layer1Ieee8023X, error)
	// ToPbText marshals Layer1Ieee8023X to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Layer1Ieee8023X to YAML text
	ToYaml() (string, error)
	// ToJson marshals Layer1Ieee8023X to JSON text
	ToJson() (string, error)
}

type unMarshallayer1Ieee8023X struct {
	obj *layer1Ieee8023X
}

type unMarshalLayer1Ieee8023X interface {
	// FromProto unmarshals Layer1Ieee8023X from protobuf object *otg.Layer1Ieee8023X
	FromProto(msg *otg.Layer1Ieee8023X) (Layer1Ieee8023X, error)
	// FromPbText unmarshals Layer1Ieee8023X from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Layer1Ieee8023X from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Layer1Ieee8023X from JSON text
	FromJson(value string) error
}

func (obj *layer1Ieee8023X) Marshal() marshalLayer1Ieee8023X {
	if obj.marshaller == nil {
		obj.marshaller = &marshallayer1Ieee8023X{obj: obj}
	}
	return obj.marshaller
}

func (obj *layer1Ieee8023X) Unmarshal() unMarshalLayer1Ieee8023X {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallayer1Ieee8023X{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallayer1Ieee8023X) ToProto() (*otg.Layer1Ieee8023X, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallayer1Ieee8023X) FromProto(msg *otg.Layer1Ieee8023X) (Layer1Ieee8023X, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallayer1Ieee8023X) ToPbText() (string, error) {
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

func (m *unMarshallayer1Ieee8023X) FromPbText(value string) error {
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

func (m *marshallayer1Ieee8023X) ToYaml() (string, error) {
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

func (m *unMarshallayer1Ieee8023X) FromYaml(value string) error {
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

func (m *marshallayer1Ieee8023X) ToJson() (string, error) {
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

func (m *unMarshallayer1Ieee8023X) FromJson(value string) error {
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

func (obj *layer1Ieee8023X) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *layer1Ieee8023X) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *layer1Ieee8023X) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *layer1Ieee8023X) Clone() (Layer1Ieee8023X, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLayer1Ieee8023X()
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

// Layer1Ieee8023X is a container for ieee 802.3x rx pause settings
type Layer1Ieee8023X interface {
	Validation
	// msg marshals Layer1Ieee8023X to protobuf object *otg.Layer1Ieee8023X
	// and doesn't set defaults
	msg() *otg.Layer1Ieee8023X
	// setMsg unmarshals Layer1Ieee8023X from protobuf object *otg.Layer1Ieee8023X
	// and doesn't set defaults
	setMsg(*otg.Layer1Ieee8023X) Layer1Ieee8023X
	// provides marshal interface
	Marshal() marshalLayer1Ieee8023X
	// provides unmarshal interface
	Unmarshal() unMarshalLayer1Ieee8023X
	// validate validates Layer1Ieee8023X
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Layer1Ieee8023X, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
}

func (obj *layer1Ieee8023X) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *layer1Ieee8023X) setDefault() {

}
