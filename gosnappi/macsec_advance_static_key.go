package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecAdvanceStaticKey *****
type macsecAdvanceStaticKey struct {
	validation
	obj             *otg.MacsecAdvanceStaticKey
	marshaller      marshalMacsecAdvanceStaticKey
	unMarshaller    unMarshalMacsecAdvanceStaticKey
	rekeyModeHolder MacsecAdvanceStaticKeyRekeyMode
}

func NewMacsecAdvanceStaticKey() MacsecAdvanceStaticKey {
	obj := macsecAdvanceStaticKey{obj: &otg.MacsecAdvanceStaticKey{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecAdvanceStaticKey) msg() *otg.MacsecAdvanceStaticKey {
	return obj.obj
}

func (obj *macsecAdvanceStaticKey) setMsg(msg *otg.MacsecAdvanceStaticKey) MacsecAdvanceStaticKey {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecAdvanceStaticKey struct {
	obj *macsecAdvanceStaticKey
}

type marshalMacsecAdvanceStaticKey interface {
	// ToProto marshals MacsecAdvanceStaticKey to protobuf object *otg.MacsecAdvanceStaticKey
	ToProto() (*otg.MacsecAdvanceStaticKey, error)
	// ToPbText marshals MacsecAdvanceStaticKey to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecAdvanceStaticKey to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecAdvanceStaticKey to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecAdvanceStaticKey struct {
	obj *macsecAdvanceStaticKey
}

type unMarshalMacsecAdvanceStaticKey interface {
	// FromProto unmarshals MacsecAdvanceStaticKey from protobuf object *otg.MacsecAdvanceStaticKey
	FromProto(msg *otg.MacsecAdvanceStaticKey) (MacsecAdvanceStaticKey, error)
	// FromPbText unmarshals MacsecAdvanceStaticKey from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecAdvanceStaticKey from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecAdvanceStaticKey from JSON text
	FromJson(value string) error
}

func (obj *macsecAdvanceStaticKey) Marshal() marshalMacsecAdvanceStaticKey {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecAdvanceStaticKey{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecAdvanceStaticKey) Unmarshal() unMarshalMacsecAdvanceStaticKey {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecAdvanceStaticKey{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecAdvanceStaticKey) ToProto() (*otg.MacsecAdvanceStaticKey, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecAdvanceStaticKey) FromProto(msg *otg.MacsecAdvanceStaticKey) (MacsecAdvanceStaticKey, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecAdvanceStaticKey) ToPbText() (string, error) {
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

func (m *unMarshalmacsecAdvanceStaticKey) FromPbText(value string) error {
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

func (m *marshalmacsecAdvanceStaticKey) ToYaml() (string, error) {
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

func (m *unMarshalmacsecAdvanceStaticKey) FromYaml(value string) error {
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

func (m *marshalmacsecAdvanceStaticKey) ToJson() (string, error) {
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

func (m *unMarshalmacsecAdvanceStaticKey) FromJson(value string) error {
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

func (obj *macsecAdvanceStaticKey) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecAdvanceStaticKey) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecAdvanceStaticKey) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecAdvanceStaticKey) Clone() (MacsecAdvanceStaticKey, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecAdvanceStaticKey()
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

func (obj *macsecAdvanceStaticKey) setNil() {
	obj.rekeyModeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecAdvanceStaticKey is static key advance settings.
type MacsecAdvanceStaticKey interface {
	Validation
	// msg marshals MacsecAdvanceStaticKey to protobuf object *otg.MacsecAdvanceStaticKey
	// and doesn't set defaults
	msg() *otg.MacsecAdvanceStaticKey
	// setMsg unmarshals MacsecAdvanceStaticKey from protobuf object *otg.MacsecAdvanceStaticKey
	// and doesn't set defaults
	setMsg(*otg.MacsecAdvanceStaticKey) MacsecAdvanceStaticKey
	// provides marshal interface
	Marshal() marshalMacsecAdvanceStaticKey
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecAdvanceStaticKey
	// validate validates MacsecAdvanceStaticKey
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecAdvanceStaticKey, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RekeyMode returns MacsecAdvanceStaticKeyRekeyMode, set in MacsecAdvanceStaticKey.
	// MacsecAdvanceStaticKeyRekeyMode is rekey mode.
	RekeyMode() MacsecAdvanceStaticKeyRekeyMode
	// SetRekeyMode assigns MacsecAdvanceStaticKeyRekeyMode provided by user to MacsecAdvanceStaticKey.
	// MacsecAdvanceStaticKeyRekeyMode is rekey mode.
	SetRekeyMode(value MacsecAdvanceStaticKeyRekeyMode) MacsecAdvanceStaticKey
	// HasRekeyMode checks if RekeyMode has been set in MacsecAdvanceStaticKey
	HasRekeyMode() bool
	setNil()
}

// description is TBD
// RekeyMode returns a MacsecAdvanceStaticKeyRekeyMode
func (obj *macsecAdvanceStaticKey) RekeyMode() MacsecAdvanceStaticKeyRekeyMode {
	if obj.obj.RekeyMode == nil {
		obj.obj.RekeyMode = NewMacsecAdvanceStaticKeyRekeyMode().msg()
	}
	if obj.rekeyModeHolder == nil {
		obj.rekeyModeHolder = &macsecAdvanceStaticKeyRekeyMode{obj: obj.obj.RekeyMode}
	}
	return obj.rekeyModeHolder
}

// description is TBD
// RekeyMode returns a MacsecAdvanceStaticKeyRekeyMode
func (obj *macsecAdvanceStaticKey) HasRekeyMode() bool {
	return obj.obj.RekeyMode != nil
}

// description is TBD
// SetRekeyMode sets the MacsecAdvanceStaticKeyRekeyMode value in the MacsecAdvanceStaticKey object
func (obj *macsecAdvanceStaticKey) SetRekeyMode(value MacsecAdvanceStaticKeyRekeyMode) MacsecAdvanceStaticKey {

	obj.rekeyModeHolder = nil
	obj.obj.RekeyMode = value.msg()

	return obj
}

func (obj *macsecAdvanceStaticKey) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RekeyMode != nil {

		obj.RekeyMode().validateObj(vObj, set_default)
	}

}

func (obj *macsecAdvanceStaticKey) setDefault() {

}
