package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Warning *****
type warning struct {
	validation
	obj          *otg.Warning
	marshaller   marshalWarning
	unMarshaller unMarshalWarning
}

func NewWarning() Warning {
	obj := warning{obj: &otg.Warning{}}
	obj.setDefault()
	return &obj
}

func (obj *warning) msg() *otg.Warning {
	return obj.obj
}

func (obj *warning) setMsg(msg *otg.Warning) Warning {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalwarning struct {
	obj *warning
}

type marshalWarning interface {
	// ToProto marshals Warning to protobuf object *otg.Warning
	ToProto() (*otg.Warning, error)
	// ToPbText marshals Warning to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Warning to YAML text
	ToYaml() (string, error)
	// ToJson marshals Warning to JSON text
	ToJson() (string, error)
}

type unMarshalwarning struct {
	obj *warning
}

type unMarshalWarning interface {
	// FromProto unmarshals Warning from protobuf object *otg.Warning
	FromProto(msg *otg.Warning) (Warning, error)
	// FromPbText unmarshals Warning from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Warning from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Warning from JSON text
	FromJson(value string) error
}

func (obj *warning) Marshal() marshalWarning {
	if obj.marshaller == nil {
		obj.marshaller = &marshalwarning{obj: obj}
	}
	return obj.marshaller
}

func (obj *warning) Unmarshal() unMarshalWarning {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalwarning{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalwarning) ToProto() (*otg.Warning, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalwarning) FromProto(msg *otg.Warning) (Warning, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalwarning) ToPbText() (string, error) {
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

func (m *unMarshalwarning) FromPbText(value string) error {
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

func (m *marshalwarning) ToYaml() (string, error) {
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

func (m *unMarshalwarning) FromYaml(value string) error {
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

func (m *marshalwarning) ToJson() (string, error) {
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

func (m *unMarshalwarning) FromJson(value string) error {
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

func (obj *warning) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *warning) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *warning) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *warning) Clone() (Warning, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewWarning()
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

// Warning is a list of warnings that have occurred while executing the request.
type Warning interface {
	Validation
	// msg marshals Warning to protobuf object *otg.Warning
	// and doesn't set defaults
	msg() *otg.Warning
	// setMsg unmarshals Warning from protobuf object *otg.Warning
	// and doesn't set defaults
	setMsg(*otg.Warning) Warning
	// provides marshal interface
	Marshal() marshalWarning
	// provides unmarshal interface
	Unmarshal() unMarshalWarning
	// validate validates Warning
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Warning, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Warnings returns []string, set in Warning.
	Warnings() []string
	// SetWarnings assigns []string provided by user to Warning
	SetWarnings(value []string) Warning
}

// A list of any system specific warnings that have occurred while
// executing the request.
// Warnings returns a []string
func (obj *warning) Warnings() []string {
	if obj.obj.Warnings == nil {
		obj.obj.Warnings = make([]string, 0)
	}
	return obj.obj.Warnings
}

// A list of any system specific warnings that have occurred while
// executing the request.
// SetWarnings sets the []string value in the Warning object
func (obj *warning) SetWarnings(value []string) Warning {

	if obj.obj.Warnings == nil {
		obj.obj.Warnings = make([]string, 0)
	}
	obj.obj.Warnings = value

	return obj
}

func (obj *warning) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *warning) setDefault() {

}
