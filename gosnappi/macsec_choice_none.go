package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecChoiceNone *****
type macsecChoiceNone struct {
	validation
	obj          *otg.MacsecChoiceNone
	marshaller   marshalMacsecChoiceNone
	unMarshaller unMarshalMacsecChoiceNone
}

func NewMacsecChoiceNone() MacsecChoiceNone {
	obj := macsecChoiceNone{obj: &otg.MacsecChoiceNone{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecChoiceNone) msg() *otg.MacsecChoiceNone {
	return obj.obj
}

func (obj *macsecChoiceNone) setMsg(msg *otg.MacsecChoiceNone) MacsecChoiceNone {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecChoiceNone struct {
	obj *macsecChoiceNone
}

type marshalMacsecChoiceNone interface {
	// ToProto marshals MacsecChoiceNone to protobuf object *otg.MacsecChoiceNone
	ToProto() (*otg.MacsecChoiceNone, error)
	// ToPbText marshals MacsecChoiceNone to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecChoiceNone to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecChoiceNone to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecChoiceNone struct {
	obj *macsecChoiceNone
}

type unMarshalMacsecChoiceNone interface {
	// FromProto unmarshals MacsecChoiceNone from protobuf object *otg.MacsecChoiceNone
	FromProto(msg *otg.MacsecChoiceNone) (MacsecChoiceNone, error)
	// FromPbText unmarshals MacsecChoiceNone from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecChoiceNone from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecChoiceNone from JSON text
	FromJson(value string) error
}

func (obj *macsecChoiceNone) Marshal() marshalMacsecChoiceNone {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecChoiceNone{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecChoiceNone) Unmarshal() unMarshalMacsecChoiceNone {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecChoiceNone{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecChoiceNone) ToProto() (*otg.MacsecChoiceNone, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecChoiceNone) FromProto(msg *otg.MacsecChoiceNone) (MacsecChoiceNone, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecChoiceNone) ToPbText() (string, error) {
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

func (m *unMarshalmacsecChoiceNone) FromPbText(value string) error {
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

func (m *marshalmacsecChoiceNone) ToYaml() (string, error) {
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

func (m *unMarshalmacsecChoiceNone) FromYaml(value string) error {
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

func (m *marshalmacsecChoiceNone) ToJson() (string, error) {
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

func (m *unMarshalmacsecChoiceNone) FromJson(value string) error {
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

func (obj *macsecChoiceNone) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecChoiceNone) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecChoiceNone) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecChoiceNone) Clone() (MacsecChoiceNone, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecChoiceNone()
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

// MacsecChoiceNone is a empty container that is used to indicate a null choice.
type MacsecChoiceNone interface {
	Validation
	// msg marshals MacsecChoiceNone to protobuf object *otg.MacsecChoiceNone
	// and doesn't set defaults
	msg() *otg.MacsecChoiceNone
	// setMsg unmarshals MacsecChoiceNone from protobuf object *otg.MacsecChoiceNone
	// and doesn't set defaults
	setMsg(*otg.MacsecChoiceNone) MacsecChoiceNone
	// provides marshal interface
	Marshal() marshalMacsecChoiceNone
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecChoiceNone
	// validate validates MacsecChoiceNone
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecChoiceNone, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
}

func (obj *macsecChoiceNone) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *macsecChoiceNone) setDefault() {

}
