package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SecureEntityCryptoEngineEncryptDecryptTxSc *****
type secureEntityCryptoEngineEncryptDecryptTxSc struct {
	validation
	obj          *otg.SecureEntityCryptoEngineEncryptDecryptTxSc
	marshaller   marshalSecureEntityCryptoEngineEncryptDecryptTxSc
	unMarshaller unMarshalSecureEntityCryptoEngineEncryptDecryptTxSc
	txPnHolder   SecureEntityCryptoEngineEncryptDecryptTxScTxPn
}

func NewSecureEntityCryptoEngineEncryptDecryptTxSc() SecureEntityCryptoEngineEncryptDecryptTxSc {
	obj := secureEntityCryptoEngineEncryptDecryptTxSc{obj: &otg.SecureEntityCryptoEngineEncryptDecryptTxSc{}}
	obj.setDefault()
	return &obj
}

func (obj *secureEntityCryptoEngineEncryptDecryptTxSc) msg() *otg.SecureEntityCryptoEngineEncryptDecryptTxSc {
	return obj.obj
}

func (obj *secureEntityCryptoEngineEncryptDecryptTxSc) setMsg(msg *otg.SecureEntityCryptoEngineEncryptDecryptTxSc) SecureEntityCryptoEngineEncryptDecryptTxSc {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsecureEntityCryptoEngineEncryptDecryptTxSc struct {
	obj *secureEntityCryptoEngineEncryptDecryptTxSc
}

type marshalSecureEntityCryptoEngineEncryptDecryptTxSc interface {
	// ToProto marshals SecureEntityCryptoEngineEncryptDecryptTxSc to protobuf object *otg.SecureEntityCryptoEngineEncryptDecryptTxSc
	ToProto() (*otg.SecureEntityCryptoEngineEncryptDecryptTxSc, error)
	// ToPbText marshals SecureEntityCryptoEngineEncryptDecryptTxSc to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SecureEntityCryptoEngineEncryptDecryptTxSc to YAML text
	ToYaml() (string, error)
	// ToJson marshals SecureEntityCryptoEngineEncryptDecryptTxSc to JSON text
	ToJson() (string, error)
}

type unMarshalsecureEntityCryptoEngineEncryptDecryptTxSc struct {
	obj *secureEntityCryptoEngineEncryptDecryptTxSc
}

type unMarshalSecureEntityCryptoEngineEncryptDecryptTxSc interface {
	// FromProto unmarshals SecureEntityCryptoEngineEncryptDecryptTxSc from protobuf object *otg.SecureEntityCryptoEngineEncryptDecryptTxSc
	FromProto(msg *otg.SecureEntityCryptoEngineEncryptDecryptTxSc) (SecureEntityCryptoEngineEncryptDecryptTxSc, error)
	// FromPbText unmarshals SecureEntityCryptoEngineEncryptDecryptTxSc from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SecureEntityCryptoEngineEncryptDecryptTxSc from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SecureEntityCryptoEngineEncryptDecryptTxSc from JSON text
	FromJson(value string) error
}

func (obj *secureEntityCryptoEngineEncryptDecryptTxSc) Marshal() marshalSecureEntityCryptoEngineEncryptDecryptTxSc {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsecureEntityCryptoEngineEncryptDecryptTxSc{obj: obj}
	}
	return obj.marshaller
}

func (obj *secureEntityCryptoEngineEncryptDecryptTxSc) Unmarshal() unMarshalSecureEntityCryptoEngineEncryptDecryptTxSc {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsecureEntityCryptoEngineEncryptDecryptTxSc{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsecureEntityCryptoEngineEncryptDecryptTxSc) ToProto() (*otg.SecureEntityCryptoEngineEncryptDecryptTxSc, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsecureEntityCryptoEngineEncryptDecryptTxSc) FromProto(msg *otg.SecureEntityCryptoEngineEncryptDecryptTxSc) (SecureEntityCryptoEngineEncryptDecryptTxSc, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsecureEntityCryptoEngineEncryptDecryptTxSc) ToPbText() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptDecryptTxSc) FromPbText(value string) error {
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

func (m *marshalsecureEntityCryptoEngineEncryptDecryptTxSc) ToYaml() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptDecryptTxSc) FromYaml(value string) error {
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

func (m *marshalsecureEntityCryptoEngineEncryptDecryptTxSc) ToJson() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptDecryptTxSc) FromJson(value string) error {
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

func (obj *secureEntityCryptoEngineEncryptDecryptTxSc) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *secureEntityCryptoEngineEncryptDecryptTxSc) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *secureEntityCryptoEngineEncryptDecryptTxSc) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *secureEntityCryptoEngineEncryptDecryptTxSc) Clone() (SecureEntityCryptoEngineEncryptDecryptTxSc, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSecureEntityCryptoEngineEncryptDecryptTxSc()
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

func (obj *secureEntityCryptoEngineEncryptDecryptTxSc) setNil() {
	obj.txPnHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// SecureEntityCryptoEngineEncryptDecryptTxSc is secure channel configuration.
type SecureEntityCryptoEngineEncryptDecryptTxSc interface {
	Validation
	// msg marshals SecureEntityCryptoEngineEncryptDecryptTxSc to protobuf object *otg.SecureEntityCryptoEngineEncryptDecryptTxSc
	// and doesn't set defaults
	msg() *otg.SecureEntityCryptoEngineEncryptDecryptTxSc
	// setMsg unmarshals SecureEntityCryptoEngineEncryptDecryptTxSc from protobuf object *otg.SecureEntityCryptoEngineEncryptDecryptTxSc
	// and doesn't set defaults
	setMsg(*otg.SecureEntityCryptoEngineEncryptDecryptTxSc) SecureEntityCryptoEngineEncryptDecryptTxSc
	// provides marshal interface
	Marshal() marshalSecureEntityCryptoEngineEncryptDecryptTxSc
	// provides unmarshal interface
	Unmarshal() unMarshalSecureEntityCryptoEngineEncryptDecryptTxSc
	// validate validates SecureEntityCryptoEngineEncryptDecryptTxSc
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SecureEntityCryptoEngineEncryptDecryptTxSc, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// TxPn returns SecureEntityCryptoEngineEncryptDecryptTxScTxPn, set in SecureEntityCryptoEngineEncryptDecryptTxSc.
	// SecureEntityCryptoEngineEncryptDecryptTxScTxPn is starting packet number(PN) configuration.
	TxPn() SecureEntityCryptoEngineEncryptDecryptTxScTxPn
	// SetTxPn assigns SecureEntityCryptoEngineEncryptDecryptTxScTxPn provided by user to SecureEntityCryptoEngineEncryptDecryptTxSc.
	// SecureEntityCryptoEngineEncryptDecryptTxScTxPn is starting packet number(PN) configuration.
	SetTxPn(value SecureEntityCryptoEngineEncryptDecryptTxScTxPn) SecureEntityCryptoEngineEncryptDecryptTxSc
	// HasTxPn checks if TxPn has been set in SecureEntityCryptoEngineEncryptDecryptTxSc
	HasTxPn() bool
	setNil()
}

// description is TBD
// TxPn returns a SecureEntityCryptoEngineEncryptDecryptTxScTxPn
func (obj *secureEntityCryptoEngineEncryptDecryptTxSc) TxPn() SecureEntityCryptoEngineEncryptDecryptTxScTxPn {
	if obj.obj.TxPn == nil {
		obj.obj.TxPn = NewSecureEntityCryptoEngineEncryptDecryptTxScTxPn().msg()
	}
	if obj.txPnHolder == nil {
		obj.txPnHolder = &secureEntityCryptoEngineEncryptDecryptTxScTxPn{obj: obj.obj.TxPn}
	}
	return obj.txPnHolder
}

// description is TBD
// TxPn returns a SecureEntityCryptoEngineEncryptDecryptTxScTxPn
func (obj *secureEntityCryptoEngineEncryptDecryptTxSc) HasTxPn() bool {
	return obj.obj.TxPn != nil
}

// description is TBD
// SetTxPn sets the SecureEntityCryptoEngineEncryptDecryptTxScTxPn value in the SecureEntityCryptoEngineEncryptDecryptTxSc object
func (obj *secureEntityCryptoEngineEncryptDecryptTxSc) SetTxPn(value SecureEntityCryptoEngineEncryptDecryptTxScTxPn) SecureEntityCryptoEngineEncryptDecryptTxSc {

	obj.txPnHolder = nil
	obj.obj.TxPn = value.msg()

	return obj
}

func (obj *secureEntityCryptoEngineEncryptDecryptTxSc) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.TxPn != nil {

		obj.TxPn().validateObj(vObj, set_default)
	}

}

func (obj *secureEntityCryptoEngineEncryptDecryptTxSc) setDefault() {

}
