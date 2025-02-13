package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecTxStaticKey *****
type macsecTxStaticKey struct {
	validation
	obj             *otg.MacsecTxStaticKey
	marshaller      marshalMacsecTxStaticKey
	unMarshaller    unMarshalMacsecTxStaticKey
	rekeyModeHolder MacsecStaticKeyRekeyMode
}

func NewMacsecTxStaticKey() MacsecTxStaticKey {
	obj := macsecTxStaticKey{obj: &otg.MacsecTxStaticKey{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecTxStaticKey) msg() *otg.MacsecTxStaticKey {
	return obj.obj
}

func (obj *macsecTxStaticKey) setMsg(msg *otg.MacsecTxStaticKey) MacsecTxStaticKey {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecTxStaticKey struct {
	obj *macsecTxStaticKey
}

type marshalMacsecTxStaticKey interface {
	// ToProto marshals MacsecTxStaticKey to protobuf object *otg.MacsecTxStaticKey
	ToProto() (*otg.MacsecTxStaticKey, error)
	// ToPbText marshals MacsecTxStaticKey to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecTxStaticKey to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecTxStaticKey to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecTxStaticKey struct {
	obj *macsecTxStaticKey
}

type unMarshalMacsecTxStaticKey interface {
	// FromProto unmarshals MacsecTxStaticKey from protobuf object *otg.MacsecTxStaticKey
	FromProto(msg *otg.MacsecTxStaticKey) (MacsecTxStaticKey, error)
	// FromPbText unmarshals MacsecTxStaticKey from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecTxStaticKey from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecTxStaticKey from JSON text
	FromJson(value string) error
}

func (obj *macsecTxStaticKey) Marshal() marshalMacsecTxStaticKey {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecTxStaticKey{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecTxStaticKey) Unmarshal() unMarshalMacsecTxStaticKey {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecTxStaticKey{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecTxStaticKey) ToProto() (*otg.MacsecTxStaticKey, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecTxStaticKey) FromProto(msg *otg.MacsecTxStaticKey) (MacsecTxStaticKey, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecTxStaticKey) ToPbText() (string, error) {
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

func (m *unMarshalmacsecTxStaticKey) FromPbText(value string) error {
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

func (m *marshalmacsecTxStaticKey) ToYaml() (string, error) {
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

func (m *unMarshalmacsecTxStaticKey) FromYaml(value string) error {
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

func (m *marshalmacsecTxStaticKey) ToJson() (string, error) {
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

func (m *unMarshalmacsecTxStaticKey) FromJson(value string) error {
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

func (obj *macsecTxStaticKey) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecTxStaticKey) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecTxStaticKey) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecTxStaticKey) Clone() (MacsecTxStaticKey, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecTxStaticKey()
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

func (obj *macsecTxStaticKey) setNil() {
	obj.rekeyModeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecTxStaticKey is tx setting for static key.
type MacsecTxStaticKey interface {
	Validation
	// msg marshals MacsecTxStaticKey to protobuf object *otg.MacsecTxStaticKey
	// and doesn't set defaults
	msg() *otg.MacsecTxStaticKey
	// setMsg unmarshals MacsecTxStaticKey from protobuf object *otg.MacsecTxStaticKey
	// and doesn't set defaults
	setMsg(*otg.MacsecTxStaticKey) MacsecTxStaticKey
	// provides marshal interface
	Marshal() marshalMacsecTxStaticKey
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecTxStaticKey
	// validate validates MacsecTxStaticKey
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecTxStaticKey, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RekeyMode returns MacsecStaticKeyRekeyMode, set in MacsecTxStaticKey.
	// MacsecStaticKeyRekeyMode is rekey mode.
	RekeyMode() MacsecStaticKeyRekeyMode
	// SetRekeyMode assigns MacsecStaticKeyRekeyMode provided by user to MacsecTxStaticKey.
	// MacsecStaticKeyRekeyMode is rekey mode.
	SetRekeyMode(value MacsecStaticKeyRekeyMode) MacsecTxStaticKey
	// HasRekeyMode checks if RekeyMode has been set in MacsecTxStaticKey
	HasRekeyMode() bool
	setNil()
}

// description is TBD
// RekeyMode returns a MacsecStaticKeyRekeyMode
func (obj *macsecTxStaticKey) RekeyMode() MacsecStaticKeyRekeyMode {
	if obj.obj.RekeyMode == nil {
		obj.obj.RekeyMode = NewMacsecStaticKeyRekeyMode().msg()
	}
	if obj.rekeyModeHolder == nil {
		obj.rekeyModeHolder = &macsecStaticKeyRekeyMode{obj: obj.obj.RekeyMode}
	}
	return obj.rekeyModeHolder
}

// description is TBD
// RekeyMode returns a MacsecStaticKeyRekeyMode
func (obj *macsecTxStaticKey) HasRekeyMode() bool {
	return obj.obj.RekeyMode != nil
}

// description is TBD
// SetRekeyMode sets the MacsecStaticKeyRekeyMode value in the MacsecTxStaticKey object
func (obj *macsecTxStaticKey) SetRekeyMode(value MacsecStaticKeyRekeyMode) MacsecTxStaticKey {

	obj.rekeyModeHolder = nil
	obj.obj.RekeyMode = value.msg()

	return obj
}

func (obj *macsecTxStaticKey) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RekeyMode != nil {

		obj.RekeyMode().validateObj(vObj, set_default)
	}

}

func (obj *macsecTxStaticKey) setDefault() {

}
