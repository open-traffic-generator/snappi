package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecBasicKeyGenerationMka *****
type macsecBasicKeyGenerationMka struct {
	validation
	obj          *otg.MacsecBasicKeyGenerationMka
	marshaller   marshalMacsecBasicKeyGenerationMka
	unMarshaller unMarshalMacsecBasicKeyGenerationMka
}

func NewMacsecBasicKeyGenerationMka() MacsecBasicKeyGenerationMka {
	obj := macsecBasicKeyGenerationMka{obj: &otg.MacsecBasicKeyGenerationMka{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecBasicKeyGenerationMka) msg() *otg.MacsecBasicKeyGenerationMka {
	return obj.obj
}

func (obj *macsecBasicKeyGenerationMka) setMsg(msg *otg.MacsecBasicKeyGenerationMka) MacsecBasicKeyGenerationMka {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecBasicKeyGenerationMka struct {
	obj *macsecBasicKeyGenerationMka
}

type marshalMacsecBasicKeyGenerationMka interface {
	// ToProto marshals MacsecBasicKeyGenerationMka to protobuf object *otg.MacsecBasicKeyGenerationMka
	ToProto() (*otg.MacsecBasicKeyGenerationMka, error)
	// ToPbText marshals MacsecBasicKeyGenerationMka to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecBasicKeyGenerationMka to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecBasicKeyGenerationMka to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecBasicKeyGenerationMka struct {
	obj *macsecBasicKeyGenerationMka
}

type unMarshalMacsecBasicKeyGenerationMka interface {
	// FromProto unmarshals MacsecBasicKeyGenerationMka from protobuf object *otg.MacsecBasicKeyGenerationMka
	FromProto(msg *otg.MacsecBasicKeyGenerationMka) (MacsecBasicKeyGenerationMka, error)
	// FromPbText unmarshals MacsecBasicKeyGenerationMka from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecBasicKeyGenerationMka from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecBasicKeyGenerationMka from JSON text
	FromJson(value string) error
}

func (obj *macsecBasicKeyGenerationMka) Marshal() marshalMacsecBasicKeyGenerationMka {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecBasicKeyGenerationMka{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecBasicKeyGenerationMka) Unmarshal() unMarshalMacsecBasicKeyGenerationMka {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecBasicKeyGenerationMka{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecBasicKeyGenerationMka) ToProto() (*otg.MacsecBasicKeyGenerationMka, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecBasicKeyGenerationMka) FromProto(msg *otg.MacsecBasicKeyGenerationMka) (MacsecBasicKeyGenerationMka, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecBasicKeyGenerationMka) ToPbText() (string, error) {
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

func (m *unMarshalmacsecBasicKeyGenerationMka) FromPbText(value string) error {
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

func (m *marshalmacsecBasicKeyGenerationMka) ToYaml() (string, error) {
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

func (m *unMarshalmacsecBasicKeyGenerationMka) FromYaml(value string) error {
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

func (m *marshalmacsecBasicKeyGenerationMka) ToJson() (string, error) {
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

func (m *unMarshalmacsecBasicKeyGenerationMka) FromJson(value string) error {
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

func (obj *macsecBasicKeyGenerationMka) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecBasicKeyGenerationMka) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecBasicKeyGenerationMka) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecBasicKeyGenerationMka) Clone() (MacsecBasicKeyGenerationMka, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecBasicKeyGenerationMka()
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

// MacsecBasicKeyGenerationMka is the container for MKA settings.
type MacsecBasicKeyGenerationMka interface {
	Validation
	// msg marshals MacsecBasicKeyGenerationMka to protobuf object *otg.MacsecBasicKeyGenerationMka
	// and doesn't set defaults
	msg() *otg.MacsecBasicKeyGenerationMka
	// setMsg unmarshals MacsecBasicKeyGenerationMka from protobuf object *otg.MacsecBasicKeyGenerationMka
	// and doesn't set defaults
	setMsg(*otg.MacsecBasicKeyGenerationMka) MacsecBasicKeyGenerationMka
	// provides marshal interface
	Marshal() marshalMacsecBasicKeyGenerationMka
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecBasicKeyGenerationMka
	// validate validates MacsecBasicKeyGenerationMka
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecBasicKeyGenerationMka, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// MkaName returns string, set in MacsecBasicKeyGenerationMka.
	MkaName() string
	// SetMkaName assigns string provided by user to MacsecBasicKeyGenerationMka
	SetMkaName(value string) MacsecBasicKeyGenerationMka
	// HasMkaName checks if MkaName has been set in MacsecBasicKeyGenerationMka
	HasMkaName() bool
}

// The unique name of the MKA instance controlling MACsec.
// MkaName returns a string
func (obj *macsecBasicKeyGenerationMka) MkaName() string {

	return *obj.obj.MkaName

}

// The unique name of the MKA instance controlling MACsec.
// MkaName returns a string
func (obj *macsecBasicKeyGenerationMka) HasMkaName() bool {
	return obj.obj.MkaName != nil
}

// The unique name of the MKA instance controlling MACsec.
// SetMkaName sets the string value in the MacsecBasicKeyGenerationMka object
func (obj *macsecBasicKeyGenerationMka) SetMkaName(value string) MacsecBasicKeyGenerationMka {

	obj.obj.MkaName = &value
	return obj
}

func (obj *macsecBasicKeyGenerationMka) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *macsecBasicKeyGenerationMka) setDefault() {

}
