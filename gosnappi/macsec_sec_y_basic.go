package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecSecYBasic *****
type macsecSecYBasic struct {
	validation
	obj                 *otg.MacsecSecYBasic
	marshaller          marshalMacsecSecYBasic
	unMarshaller        unMarshalMacsecSecYBasic
	keyGenerationHolder MacsecSecYBasicKeyGeneration
}

func NewMacsecSecYBasic() MacsecSecYBasic {
	obj := macsecSecYBasic{obj: &otg.MacsecSecYBasic{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecSecYBasic) msg() *otg.MacsecSecYBasic {
	return obj.obj
}

func (obj *macsecSecYBasic) setMsg(msg *otg.MacsecSecYBasic) MacsecSecYBasic {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecSecYBasic struct {
	obj *macsecSecYBasic
}

type marshalMacsecSecYBasic interface {
	// ToProto marshals MacsecSecYBasic to protobuf object *otg.MacsecSecYBasic
	ToProto() (*otg.MacsecSecYBasic, error)
	// ToPbText marshals MacsecSecYBasic to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecSecYBasic to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecSecYBasic to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecSecYBasic struct {
	obj *macsecSecYBasic
}

type unMarshalMacsecSecYBasic interface {
	// FromProto unmarshals MacsecSecYBasic from protobuf object *otg.MacsecSecYBasic
	FromProto(msg *otg.MacsecSecYBasic) (MacsecSecYBasic, error)
	// FromPbText unmarshals MacsecSecYBasic from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecSecYBasic from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecSecYBasic from JSON text
	FromJson(value string) error
}

func (obj *macsecSecYBasic) Marshal() marshalMacsecSecYBasic {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecSecYBasic{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecSecYBasic) Unmarshal() unMarshalMacsecSecYBasic {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecSecYBasic{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecSecYBasic) ToProto() (*otg.MacsecSecYBasic, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecSecYBasic) FromProto(msg *otg.MacsecSecYBasic) (MacsecSecYBasic, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecSecYBasic) ToPbText() (string, error) {
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

func (m *unMarshalmacsecSecYBasic) FromPbText(value string) error {
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

func (m *marshalmacsecSecYBasic) ToYaml() (string, error) {
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

func (m *unMarshalmacsecSecYBasic) FromYaml(value string) error {
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

func (m *marshalmacsecSecYBasic) ToJson() (string, error) {
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

func (m *unMarshalmacsecSecYBasic) FromJson(value string) error {
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

func (obj *macsecSecYBasic) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecSecYBasic) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecSecYBasic) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecSecYBasic) Clone() (MacsecSecYBasic, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecSecYBasic()
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

func (obj *macsecSecYBasic) setNil() {
	obj.keyGenerationHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecSecYBasic is a container of basic properties for a SecY.
type MacsecSecYBasic interface {
	Validation
	// msg marshals MacsecSecYBasic to protobuf object *otg.MacsecSecYBasic
	// and doesn't set defaults
	msg() *otg.MacsecSecYBasic
	// setMsg unmarshals MacsecSecYBasic from protobuf object *otg.MacsecSecYBasic
	// and doesn't set defaults
	setMsg(*otg.MacsecSecYBasic) MacsecSecYBasic
	// provides marshal interface
	Marshal() marshalMacsecSecYBasic
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecSecYBasic
	// validate validates MacsecSecYBasic
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecSecYBasic, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// KeyGeneration returns MacsecSecYBasicKeyGeneration, set in MacsecSecYBasic.
	// MacsecSecYBasicKeyGeneration is key Generation.
	KeyGeneration() MacsecSecYBasicKeyGeneration
	// SetKeyGeneration assigns MacsecSecYBasicKeyGeneration provided by user to MacsecSecYBasic.
	// MacsecSecYBasicKeyGeneration is key Generation.
	SetKeyGeneration(value MacsecSecYBasicKeyGeneration) MacsecSecYBasic
	// HasKeyGeneration checks if KeyGeneration has been set in MacsecSecYBasic
	HasKeyGeneration() bool
	setNil()
}

// description is TBD
// KeyGeneration returns a MacsecSecYBasicKeyGeneration
func (obj *macsecSecYBasic) KeyGeneration() MacsecSecYBasicKeyGeneration {
	if obj.obj.KeyGeneration == nil {
		obj.obj.KeyGeneration = NewMacsecSecYBasicKeyGeneration().msg()
	}
	if obj.keyGenerationHolder == nil {
		obj.keyGenerationHolder = &macsecSecYBasicKeyGeneration{obj: obj.obj.KeyGeneration}
	}
	return obj.keyGenerationHolder
}

// description is TBD
// KeyGeneration returns a MacsecSecYBasicKeyGeneration
func (obj *macsecSecYBasic) HasKeyGeneration() bool {
	return obj.obj.KeyGeneration != nil
}

// description is TBD
// SetKeyGeneration sets the MacsecSecYBasicKeyGeneration value in the MacsecSecYBasic object
func (obj *macsecSecYBasic) SetKeyGeneration(value MacsecSecYBasicKeyGeneration) MacsecSecYBasic {

	obj.keyGenerationHolder = nil
	obj.obj.KeyGeneration = value.msg()

	return obj
}

func (obj *macsecSecYBasic) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.KeyGeneration != nil {

		obj.KeyGeneration().validateObj(vObj, set_default)
	}

}

func (obj *macsecSecYBasic) setDefault() {

}
