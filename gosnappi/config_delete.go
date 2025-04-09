package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ConfigDelete *****
type configDelete struct {
	validation
	obj          *otg.ConfigDelete
	marshaller   marshalConfigDelete
	unMarshaller unMarshalConfigDelete
}

func NewConfigDelete() ConfigDelete {
	obj := configDelete{obj: &otg.ConfigDelete{}}
	obj.setDefault()
	return &obj
}

func (obj *configDelete) msg() *otg.ConfigDelete {
	return obj.obj
}

func (obj *configDelete) setMsg(msg *otg.ConfigDelete) ConfigDelete {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalconfigDelete struct {
	obj *configDelete
}

type marshalConfigDelete interface {
	// ToProto marshals ConfigDelete to protobuf object *otg.ConfigDelete
	ToProto() (*otg.ConfigDelete, error)
	// ToPbText marshals ConfigDelete to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ConfigDelete to YAML text
	ToYaml() (string, error)
	// ToJson marshals ConfigDelete to JSON text
	ToJson() (string, error)
}

type unMarshalconfigDelete struct {
	obj *configDelete
}

type unMarshalConfigDelete interface {
	// FromProto unmarshals ConfigDelete from protobuf object *otg.ConfigDelete
	FromProto(msg *otg.ConfigDelete) (ConfigDelete, error)
	// FromPbText unmarshals ConfigDelete from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ConfigDelete from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ConfigDelete from JSON text
	FromJson(value string) error
}

func (obj *configDelete) Marshal() marshalConfigDelete {
	if obj.marshaller == nil {
		obj.marshaller = &marshalconfigDelete{obj: obj}
	}
	return obj.marshaller
}

func (obj *configDelete) Unmarshal() unMarshalConfigDelete {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalconfigDelete{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalconfigDelete) ToProto() (*otg.ConfigDelete, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalconfigDelete) FromProto(msg *otg.ConfigDelete) (ConfigDelete, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalconfigDelete) ToPbText() (string, error) {
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

func (m *unMarshalconfigDelete) FromPbText(value string) error {
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

func (m *marshalconfigDelete) ToYaml() (string, error) {
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

func (m *unMarshalconfigDelete) FromYaml(value string) error {
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

func (m *marshalconfigDelete) ToJson() (string, error) {
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

func (m *unMarshalconfigDelete) FromJson(value string) error {
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

func (obj *configDelete) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *configDelete) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *configDelete) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *configDelete) Clone() (ConfigDelete, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewConfigDelete()
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

// ConfigDelete is a container for all configuration resources of various types to be deleted.
type ConfigDelete interface {
	Validation
	// msg marshals ConfigDelete to protobuf object *otg.ConfigDelete
	// and doesn't set defaults
	msg() *otg.ConfigDelete
	// setMsg unmarshals ConfigDelete from protobuf object *otg.ConfigDelete
	// and doesn't set defaults
	setMsg(*otg.ConfigDelete) ConfigDelete
	// provides marshal interface
	Marshal() marshalConfigDelete
	// provides unmarshal interface
	Unmarshal() unMarshalConfigDelete
	// validate validates ConfigDelete
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ConfigDelete, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
}

func (obj *configDelete) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *configDelete) setDefault() {

}
