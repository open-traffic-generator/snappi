package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecSecYAdvanceStaticKey *****
type macsecSecYAdvanceStaticKey struct {
	validation
	obj             *otg.MacsecSecYAdvanceStaticKey
	marshaller      marshalMacsecSecYAdvanceStaticKey
	unMarshaller    unMarshalMacsecSecYAdvanceStaticKey
	rekeyModeHolder MacsecSecYAdvanceStaticKeyRekeyMode
}

func NewMacsecSecYAdvanceStaticKey() MacsecSecYAdvanceStaticKey {
	obj := macsecSecYAdvanceStaticKey{obj: &otg.MacsecSecYAdvanceStaticKey{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecSecYAdvanceStaticKey) msg() *otg.MacsecSecYAdvanceStaticKey {
	return obj.obj
}

func (obj *macsecSecYAdvanceStaticKey) setMsg(msg *otg.MacsecSecYAdvanceStaticKey) MacsecSecYAdvanceStaticKey {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecSecYAdvanceStaticKey struct {
	obj *macsecSecYAdvanceStaticKey
}

type marshalMacsecSecYAdvanceStaticKey interface {
	// ToProto marshals MacsecSecYAdvanceStaticKey to protobuf object *otg.MacsecSecYAdvanceStaticKey
	ToProto() (*otg.MacsecSecYAdvanceStaticKey, error)
	// ToPbText marshals MacsecSecYAdvanceStaticKey to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecSecYAdvanceStaticKey to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecSecYAdvanceStaticKey to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecSecYAdvanceStaticKey struct {
	obj *macsecSecYAdvanceStaticKey
}

type unMarshalMacsecSecYAdvanceStaticKey interface {
	// FromProto unmarshals MacsecSecYAdvanceStaticKey from protobuf object *otg.MacsecSecYAdvanceStaticKey
	FromProto(msg *otg.MacsecSecYAdvanceStaticKey) (MacsecSecYAdvanceStaticKey, error)
	// FromPbText unmarshals MacsecSecYAdvanceStaticKey from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecSecYAdvanceStaticKey from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecSecYAdvanceStaticKey from JSON text
	FromJson(value string) error
}

func (obj *macsecSecYAdvanceStaticKey) Marshal() marshalMacsecSecYAdvanceStaticKey {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecSecYAdvanceStaticKey{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecSecYAdvanceStaticKey) Unmarshal() unMarshalMacsecSecYAdvanceStaticKey {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecSecYAdvanceStaticKey{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecSecYAdvanceStaticKey) ToProto() (*otg.MacsecSecYAdvanceStaticKey, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecSecYAdvanceStaticKey) FromProto(msg *otg.MacsecSecYAdvanceStaticKey) (MacsecSecYAdvanceStaticKey, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecSecYAdvanceStaticKey) ToPbText() (string, error) {
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

func (m *unMarshalmacsecSecYAdvanceStaticKey) FromPbText(value string) error {
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

func (m *marshalmacsecSecYAdvanceStaticKey) ToYaml() (string, error) {
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

func (m *unMarshalmacsecSecYAdvanceStaticKey) FromYaml(value string) error {
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

func (m *marshalmacsecSecYAdvanceStaticKey) ToJson() (string, error) {
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

func (m *unMarshalmacsecSecYAdvanceStaticKey) FromJson(value string) error {
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

func (obj *macsecSecYAdvanceStaticKey) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecSecYAdvanceStaticKey) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecSecYAdvanceStaticKey) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecSecYAdvanceStaticKey) Clone() (MacsecSecYAdvanceStaticKey, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecSecYAdvanceStaticKey()
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

func (obj *macsecSecYAdvanceStaticKey) setNil() {
	obj.rekeyModeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecSecYAdvanceStaticKey is static key advance settings.
type MacsecSecYAdvanceStaticKey interface {
	Validation
	// msg marshals MacsecSecYAdvanceStaticKey to protobuf object *otg.MacsecSecYAdvanceStaticKey
	// and doesn't set defaults
	msg() *otg.MacsecSecYAdvanceStaticKey
	// setMsg unmarshals MacsecSecYAdvanceStaticKey from protobuf object *otg.MacsecSecYAdvanceStaticKey
	// and doesn't set defaults
	setMsg(*otg.MacsecSecYAdvanceStaticKey) MacsecSecYAdvanceStaticKey
	// provides marshal interface
	Marshal() marshalMacsecSecYAdvanceStaticKey
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecSecYAdvanceStaticKey
	// validate validates MacsecSecYAdvanceStaticKey
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecSecYAdvanceStaticKey, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RekeyMode returns MacsecSecYAdvanceStaticKeyRekeyMode, set in MacsecSecYAdvanceStaticKey.
	// MacsecSecYAdvanceStaticKeyRekeyMode is rekey mode.
	RekeyMode() MacsecSecYAdvanceStaticKeyRekeyMode
	// SetRekeyMode assigns MacsecSecYAdvanceStaticKeyRekeyMode provided by user to MacsecSecYAdvanceStaticKey.
	// MacsecSecYAdvanceStaticKeyRekeyMode is rekey mode.
	SetRekeyMode(value MacsecSecYAdvanceStaticKeyRekeyMode) MacsecSecYAdvanceStaticKey
	// HasRekeyMode checks if RekeyMode has been set in MacsecSecYAdvanceStaticKey
	HasRekeyMode() bool
	setNil()
}

// description is TBD
// RekeyMode returns a MacsecSecYAdvanceStaticKeyRekeyMode
func (obj *macsecSecYAdvanceStaticKey) RekeyMode() MacsecSecYAdvanceStaticKeyRekeyMode {
	if obj.obj.RekeyMode == nil {
		obj.obj.RekeyMode = NewMacsecSecYAdvanceStaticKeyRekeyMode().msg()
	}
	if obj.rekeyModeHolder == nil {
		obj.rekeyModeHolder = &macsecSecYAdvanceStaticKeyRekeyMode{obj: obj.obj.RekeyMode}
	}
	return obj.rekeyModeHolder
}

// description is TBD
// RekeyMode returns a MacsecSecYAdvanceStaticKeyRekeyMode
func (obj *macsecSecYAdvanceStaticKey) HasRekeyMode() bool {
	return obj.obj.RekeyMode != nil
}

// description is TBD
// SetRekeyMode sets the MacsecSecYAdvanceStaticKeyRekeyMode value in the MacsecSecYAdvanceStaticKey object
func (obj *macsecSecYAdvanceStaticKey) SetRekeyMode(value MacsecSecYAdvanceStaticKeyRekeyMode) MacsecSecYAdvanceStaticKey {

	obj.rekeyModeHolder = nil
	obj.obj.RekeyMode = value.msg()

	return obj
}

func (obj *macsecSecYAdvanceStaticKey) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RekeyMode != nil {

		obj.RekeyMode().validateObj(vObj, set_default)
	}

}

func (obj *macsecSecYAdvanceStaticKey) setDefault() {

}
