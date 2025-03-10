package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SecureEntityCryptoEngineEncryptOnlyTxSc *****
type secureEntityCryptoEngineEncryptOnlyTxSc struct {
	validation
	obj          *otg.SecureEntityCryptoEngineEncryptOnlyTxSc
	marshaller   marshalSecureEntityCryptoEngineEncryptOnlyTxSc
	unMarshaller unMarshalSecureEntityCryptoEngineEncryptOnlyTxSc
	txPnHolder   SecureEntityCryptoEngineEncryptOnlyTxScTxPn
}

func NewSecureEntityCryptoEngineEncryptOnlyTxSc() SecureEntityCryptoEngineEncryptOnlyTxSc {
	obj := secureEntityCryptoEngineEncryptOnlyTxSc{obj: &otg.SecureEntityCryptoEngineEncryptOnlyTxSc{}}
	obj.setDefault()
	return &obj
}

func (obj *secureEntityCryptoEngineEncryptOnlyTxSc) msg() *otg.SecureEntityCryptoEngineEncryptOnlyTxSc {
	return obj.obj
}

func (obj *secureEntityCryptoEngineEncryptOnlyTxSc) setMsg(msg *otg.SecureEntityCryptoEngineEncryptOnlyTxSc) SecureEntityCryptoEngineEncryptOnlyTxSc {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsecureEntityCryptoEngineEncryptOnlyTxSc struct {
	obj *secureEntityCryptoEngineEncryptOnlyTxSc
}

type marshalSecureEntityCryptoEngineEncryptOnlyTxSc interface {
	// ToProto marshals SecureEntityCryptoEngineEncryptOnlyTxSc to protobuf object *otg.SecureEntityCryptoEngineEncryptOnlyTxSc
	ToProto() (*otg.SecureEntityCryptoEngineEncryptOnlyTxSc, error)
	// ToPbText marshals SecureEntityCryptoEngineEncryptOnlyTxSc to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SecureEntityCryptoEngineEncryptOnlyTxSc to YAML text
	ToYaml() (string, error)
	// ToJson marshals SecureEntityCryptoEngineEncryptOnlyTxSc to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals SecureEntityCryptoEngineEncryptOnlyTxSc to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalsecureEntityCryptoEngineEncryptOnlyTxSc struct {
	obj *secureEntityCryptoEngineEncryptOnlyTxSc
}

type unMarshalSecureEntityCryptoEngineEncryptOnlyTxSc interface {
	// FromProto unmarshals SecureEntityCryptoEngineEncryptOnlyTxSc from protobuf object *otg.SecureEntityCryptoEngineEncryptOnlyTxSc
	FromProto(msg *otg.SecureEntityCryptoEngineEncryptOnlyTxSc) (SecureEntityCryptoEngineEncryptOnlyTxSc, error)
	// FromPbText unmarshals SecureEntityCryptoEngineEncryptOnlyTxSc from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SecureEntityCryptoEngineEncryptOnlyTxSc from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SecureEntityCryptoEngineEncryptOnlyTxSc from JSON text
	FromJson(value string) error
}

func (obj *secureEntityCryptoEngineEncryptOnlyTxSc) Marshal() marshalSecureEntityCryptoEngineEncryptOnlyTxSc {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsecureEntityCryptoEngineEncryptOnlyTxSc{obj: obj}
	}
	return obj.marshaller
}

func (obj *secureEntityCryptoEngineEncryptOnlyTxSc) Unmarshal() unMarshalSecureEntityCryptoEngineEncryptOnlyTxSc {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsecureEntityCryptoEngineEncryptOnlyTxSc{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsecureEntityCryptoEngineEncryptOnlyTxSc) ToProto() (*otg.SecureEntityCryptoEngineEncryptOnlyTxSc, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsecureEntityCryptoEngineEncryptOnlyTxSc) FromProto(msg *otg.SecureEntityCryptoEngineEncryptOnlyTxSc) (SecureEntityCryptoEngineEncryptOnlyTxSc, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsecureEntityCryptoEngineEncryptOnlyTxSc) ToPbText() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptOnlyTxSc) FromPbText(value string) error {
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

func (m *marshalsecureEntityCryptoEngineEncryptOnlyTxSc) ToYaml() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptOnlyTxSc) FromYaml(value string) error {
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

func (m *marshalsecureEntityCryptoEngineEncryptOnlyTxSc) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalsecureEntityCryptoEngineEncryptOnlyTxSc) ToJson() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptOnlyTxSc) FromJson(value string) error {
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

func (obj *secureEntityCryptoEngineEncryptOnlyTxSc) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *secureEntityCryptoEngineEncryptOnlyTxSc) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *secureEntityCryptoEngineEncryptOnlyTxSc) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *secureEntityCryptoEngineEncryptOnlyTxSc) Clone() (SecureEntityCryptoEngineEncryptOnlyTxSc, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSecureEntityCryptoEngineEncryptOnlyTxSc()
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

func (obj *secureEntityCryptoEngineEncryptOnlyTxSc) setNil() {
	obj.txPnHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// SecureEntityCryptoEngineEncryptOnlyTxSc is the container for Tx secure channel configuration.
type SecureEntityCryptoEngineEncryptOnlyTxSc interface {
	Validation
	// msg marshals SecureEntityCryptoEngineEncryptOnlyTxSc to protobuf object *otg.SecureEntityCryptoEngineEncryptOnlyTxSc
	// and doesn't set defaults
	msg() *otg.SecureEntityCryptoEngineEncryptOnlyTxSc
	// setMsg unmarshals SecureEntityCryptoEngineEncryptOnlyTxSc from protobuf object *otg.SecureEntityCryptoEngineEncryptOnlyTxSc
	// and doesn't set defaults
	setMsg(*otg.SecureEntityCryptoEngineEncryptOnlyTxSc) SecureEntityCryptoEngineEncryptOnlyTxSc
	// provides marshal interface
	Marshal() marshalSecureEntityCryptoEngineEncryptOnlyTxSc
	// provides unmarshal interface
	Unmarshal() unMarshalSecureEntityCryptoEngineEncryptOnlyTxSc
	// validate validates SecureEntityCryptoEngineEncryptOnlyTxSc
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SecureEntityCryptoEngineEncryptOnlyTxSc, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// TxPn returns SecureEntityCryptoEngineEncryptOnlyTxScTxPn, set in SecureEntityCryptoEngineEncryptOnlyTxSc.
	// SecureEntityCryptoEngineEncryptOnlyTxScTxPn is tx packet number(PN) configuration.
	TxPn() SecureEntityCryptoEngineEncryptOnlyTxScTxPn
	// SetTxPn assigns SecureEntityCryptoEngineEncryptOnlyTxScTxPn provided by user to SecureEntityCryptoEngineEncryptOnlyTxSc.
	// SecureEntityCryptoEngineEncryptOnlyTxScTxPn is tx packet number(PN) configuration.
	SetTxPn(value SecureEntityCryptoEngineEncryptOnlyTxScTxPn) SecureEntityCryptoEngineEncryptOnlyTxSc
	// HasTxPn checks if TxPn has been set in SecureEntityCryptoEngineEncryptOnlyTxSc
	HasTxPn() bool
	setNil()
}

// description is TBD
// TxPn returns a SecureEntityCryptoEngineEncryptOnlyTxScTxPn
func (obj *secureEntityCryptoEngineEncryptOnlyTxSc) TxPn() SecureEntityCryptoEngineEncryptOnlyTxScTxPn {
	if obj.obj.TxPn == nil {
		obj.obj.TxPn = NewSecureEntityCryptoEngineEncryptOnlyTxScTxPn().msg()
	}
	if obj.txPnHolder == nil {
		obj.txPnHolder = &secureEntityCryptoEngineEncryptOnlyTxScTxPn{obj: obj.obj.TxPn}
	}
	return obj.txPnHolder
}

// description is TBD
// TxPn returns a SecureEntityCryptoEngineEncryptOnlyTxScTxPn
func (obj *secureEntityCryptoEngineEncryptOnlyTxSc) HasTxPn() bool {
	return obj.obj.TxPn != nil
}

// description is TBD
// SetTxPn sets the SecureEntityCryptoEngineEncryptOnlyTxScTxPn value in the SecureEntityCryptoEngineEncryptOnlyTxSc object
func (obj *secureEntityCryptoEngineEncryptOnlyTxSc) SetTxPn(value SecureEntityCryptoEngineEncryptOnlyTxScTxPn) SecureEntityCryptoEngineEncryptOnlyTxSc {

	obj.txPnHolder = nil
	obj.obj.TxPn = value.msg()

	return obj
}

func (obj *secureEntityCryptoEngineEncryptOnlyTxSc) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.TxPn != nil {

		obj.TxPn().validateObj(vObj, set_default)
	}

}

func (obj *secureEntityCryptoEngineEncryptOnlyTxSc) setDefault() {

}
