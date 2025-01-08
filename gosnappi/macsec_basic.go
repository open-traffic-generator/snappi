package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecBasic *****
type macsecBasic struct {
	validation
	obj                 *otg.MacsecBasic
	marshaller          marshalMacsecBasic
	unMarshaller        unMarshalMacsecBasic
	keyGenerationHolder MacsecBasicKeyGeneration
}

func NewMacsecBasic() MacsecBasic {
	obj := macsecBasic{obj: &otg.MacsecBasic{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecBasic) msg() *otg.MacsecBasic {
	return obj.obj
}

func (obj *macsecBasic) setMsg(msg *otg.MacsecBasic) MacsecBasic {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecBasic struct {
	obj *macsecBasic
}

type marshalMacsecBasic interface {
	// ToProto marshals MacsecBasic to protobuf object *otg.MacsecBasic
	ToProto() (*otg.MacsecBasic, error)
	// ToPbText marshals MacsecBasic to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecBasic to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecBasic to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecBasic struct {
	obj *macsecBasic
}

type unMarshalMacsecBasic interface {
	// FromProto unmarshals MacsecBasic from protobuf object *otg.MacsecBasic
	FromProto(msg *otg.MacsecBasic) (MacsecBasic, error)
	// FromPbText unmarshals MacsecBasic from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecBasic from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecBasic from JSON text
	FromJson(value string) error
}

func (obj *macsecBasic) Marshal() marshalMacsecBasic {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecBasic{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecBasic) Unmarshal() unMarshalMacsecBasic {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecBasic{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecBasic) ToProto() (*otg.MacsecBasic, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecBasic) FromProto(msg *otg.MacsecBasic) (MacsecBasic, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecBasic) ToPbText() (string, error) {
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

func (m *unMarshalmacsecBasic) FromPbText(value string) error {
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

func (m *marshalmacsecBasic) ToYaml() (string, error) {
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

func (m *unMarshalmacsecBasic) FromYaml(value string) error {
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

func (m *marshalmacsecBasic) ToJson() (string, error) {
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

func (m *unMarshalmacsecBasic) FromJson(value string) error {
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

func (obj *macsecBasic) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecBasic) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecBasic) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecBasic) Clone() (MacsecBasic, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecBasic()
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

func (obj *macsecBasic) setNil() {
	obj.keyGenerationHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecBasic is a container of basic properties for a MACsec interface.
type MacsecBasic interface {
	Validation
	// msg marshals MacsecBasic to protobuf object *otg.MacsecBasic
	// and doesn't set defaults
	msg() *otg.MacsecBasic
	// setMsg unmarshals MacsecBasic from protobuf object *otg.MacsecBasic
	// and doesn't set defaults
	setMsg(*otg.MacsecBasic) MacsecBasic
	// provides marshal interface
	Marshal() marshalMacsecBasic
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecBasic
	// validate validates MacsecBasic
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecBasic, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// KeyGeneration returns MacsecBasicKeyGeneration, set in MacsecBasic.
	// MacsecBasicKeyGeneration is key Generation.
	KeyGeneration() MacsecBasicKeyGeneration
	// SetKeyGeneration assigns MacsecBasicKeyGeneration provided by user to MacsecBasic.
	// MacsecBasicKeyGeneration is key Generation.
	SetKeyGeneration(value MacsecBasicKeyGeneration) MacsecBasic
	// HasKeyGeneration checks if KeyGeneration has been set in MacsecBasic
	HasKeyGeneration() bool
	setNil()
}

// description is TBD
// KeyGeneration returns a MacsecBasicKeyGeneration
func (obj *macsecBasic) KeyGeneration() MacsecBasicKeyGeneration {
	if obj.obj.KeyGeneration == nil {
		obj.obj.KeyGeneration = NewMacsecBasicKeyGeneration().msg()
	}
	if obj.keyGenerationHolder == nil {
		obj.keyGenerationHolder = &macsecBasicKeyGeneration{obj: obj.obj.KeyGeneration}
	}
	return obj.keyGenerationHolder
}

// description is TBD
// KeyGeneration returns a MacsecBasicKeyGeneration
func (obj *macsecBasic) HasKeyGeneration() bool {
	return obj.obj.KeyGeneration != nil
}

// description is TBD
// SetKeyGeneration sets the MacsecBasicKeyGeneration value in the MacsecBasic object
func (obj *macsecBasic) SetKeyGeneration(value MacsecBasicKeyGeneration) MacsecBasic {

	obj.keyGenerationHolder = nil
	obj.obj.KeyGeneration = value.msg()

	return obj
}

func (obj *macsecBasic) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.KeyGeneration != nil {

		obj.KeyGeneration().validateObj(vObj, set_default)
	}

}

func (obj *macsecBasic) setDefault() {

}
