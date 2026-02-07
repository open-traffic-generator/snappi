package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SecureEntityCryptoEngineEncryptDecrypt *****
type secureEntityCryptoEngineEncryptDecrypt struct {
	validation
	obj                        *otg.SecureEntityCryptoEngineEncryptDecrypt
	marshaller                 marshalSecureEntityCryptoEngineEncryptDecrypt
	unMarshaller               unMarshalSecureEntityCryptoEngineEncryptDecrypt
	txPnHolder                 SecureEntityCryptoEngineEncryptDecryptTxPn
	hardwareAccelerationHolder SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
}

func NewSecureEntityCryptoEngineEncryptDecrypt() SecureEntityCryptoEngineEncryptDecrypt {
	obj := secureEntityCryptoEngineEncryptDecrypt{obj: &otg.SecureEntityCryptoEngineEncryptDecrypt{}}
	obj.setDefault()
	return &obj
}

func (obj *secureEntityCryptoEngineEncryptDecrypt) msg() *otg.SecureEntityCryptoEngineEncryptDecrypt {
	return obj.obj
}

func (obj *secureEntityCryptoEngineEncryptDecrypt) setMsg(msg *otg.SecureEntityCryptoEngineEncryptDecrypt) SecureEntityCryptoEngineEncryptDecrypt {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsecureEntityCryptoEngineEncryptDecrypt struct {
	obj *secureEntityCryptoEngineEncryptDecrypt
}

type marshalSecureEntityCryptoEngineEncryptDecrypt interface {
	// ToProto marshals SecureEntityCryptoEngineEncryptDecrypt to protobuf object *otg.SecureEntityCryptoEngineEncryptDecrypt
	ToProto() (*otg.SecureEntityCryptoEngineEncryptDecrypt, error)
	// ToPbText marshals SecureEntityCryptoEngineEncryptDecrypt to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SecureEntityCryptoEngineEncryptDecrypt to YAML text
	ToYaml() (string, error)
	// ToJson marshals SecureEntityCryptoEngineEncryptDecrypt to JSON text
	ToJson() (string, error)
}

type unMarshalsecureEntityCryptoEngineEncryptDecrypt struct {
	obj *secureEntityCryptoEngineEncryptDecrypt
}

type unMarshalSecureEntityCryptoEngineEncryptDecrypt interface {
	// FromProto unmarshals SecureEntityCryptoEngineEncryptDecrypt from protobuf object *otg.SecureEntityCryptoEngineEncryptDecrypt
	FromProto(msg *otg.SecureEntityCryptoEngineEncryptDecrypt) (SecureEntityCryptoEngineEncryptDecrypt, error)
	// FromPbText unmarshals SecureEntityCryptoEngineEncryptDecrypt from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SecureEntityCryptoEngineEncryptDecrypt from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SecureEntityCryptoEngineEncryptDecrypt from JSON text
	FromJson(value string) error
}

func (obj *secureEntityCryptoEngineEncryptDecrypt) Marshal() marshalSecureEntityCryptoEngineEncryptDecrypt {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsecureEntityCryptoEngineEncryptDecrypt{obj: obj}
	}
	return obj.marshaller
}

func (obj *secureEntityCryptoEngineEncryptDecrypt) Unmarshal() unMarshalSecureEntityCryptoEngineEncryptDecrypt {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsecureEntityCryptoEngineEncryptDecrypt{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsecureEntityCryptoEngineEncryptDecrypt) ToProto() (*otg.SecureEntityCryptoEngineEncryptDecrypt, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsecureEntityCryptoEngineEncryptDecrypt) FromProto(msg *otg.SecureEntityCryptoEngineEncryptDecrypt) (SecureEntityCryptoEngineEncryptDecrypt, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsecureEntityCryptoEngineEncryptDecrypt) ToPbText() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptDecrypt) FromPbText(value string) error {
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

func (m *marshalsecureEntityCryptoEngineEncryptDecrypt) ToYaml() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptDecrypt) FromYaml(value string) error {
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

func (m *marshalsecureEntityCryptoEngineEncryptDecrypt) ToJson() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptDecrypt) FromJson(value string) error {
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

func (obj *secureEntityCryptoEngineEncryptDecrypt) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *secureEntityCryptoEngineEncryptDecrypt) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *secureEntityCryptoEngineEncryptDecrypt) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *secureEntityCryptoEngineEncryptDecrypt) Clone() (SecureEntityCryptoEngineEncryptDecrypt, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSecureEntityCryptoEngineEncryptDecrypt()
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

func (obj *secureEntityCryptoEngineEncryptDecrypt) setNil() {
	obj.txPnHolder = nil
	obj.hardwareAccelerationHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// SecureEntityCryptoEngineEncryptDecrypt is the container for encrypt and decrypt engine configuration.
type SecureEntityCryptoEngineEncryptDecrypt interface {
	Validation
	// msg marshals SecureEntityCryptoEngineEncryptDecrypt to protobuf object *otg.SecureEntityCryptoEngineEncryptDecrypt
	// and doesn't set defaults
	msg() *otg.SecureEntityCryptoEngineEncryptDecrypt
	// setMsg unmarshals SecureEntityCryptoEngineEncryptDecrypt from protobuf object *otg.SecureEntityCryptoEngineEncryptDecrypt
	// and doesn't set defaults
	setMsg(*otg.SecureEntityCryptoEngineEncryptDecrypt) SecureEntityCryptoEngineEncryptDecrypt
	// provides marshal interface
	Marshal() marshalSecureEntityCryptoEngineEncryptDecrypt
	// provides unmarshal interface
	Unmarshal() unMarshalSecureEntityCryptoEngineEncryptDecrypt
	// validate validates SecureEntityCryptoEngineEncryptDecrypt
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SecureEntityCryptoEngineEncryptDecrypt, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// TxPn returns SecureEntityCryptoEngineEncryptDecryptTxPn, set in SecureEntityCryptoEngineEncryptDecrypt.
	// SecureEntityCryptoEngineEncryptDecryptTxPn is tx packet number(PN) configuration.
	TxPn() SecureEntityCryptoEngineEncryptDecryptTxPn
	// SetTxPn assigns SecureEntityCryptoEngineEncryptDecryptTxPn provided by user to SecureEntityCryptoEngineEncryptDecrypt.
	// SecureEntityCryptoEngineEncryptDecryptTxPn is tx packet number(PN) configuration.
	SetTxPn(value SecureEntityCryptoEngineEncryptDecryptTxPn) SecureEntityCryptoEngineEncryptDecrypt
	// HasTxPn checks if TxPn has been set in SecureEntityCryptoEngineEncryptDecrypt
	HasTxPn() bool
	// HardwareAcceleration returns SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration, set in SecureEntityCryptoEngineEncryptDecrypt.
	// SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration is hardware acceleration configuration for offloading MACsec processing to hardware.
	HardwareAcceleration() SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	// SetHardwareAcceleration assigns SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration provided by user to SecureEntityCryptoEngineEncryptDecrypt.
	// SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration is hardware acceleration configuration for offloading MACsec processing to hardware.
	SetHardwareAcceleration(value SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) SecureEntityCryptoEngineEncryptDecrypt
	// HasHardwareAcceleration checks if HardwareAcceleration has been set in SecureEntityCryptoEngineEncryptDecrypt
	HasHardwareAcceleration() bool
	setNil()
}

// description is TBD
// TxPn returns a SecureEntityCryptoEngineEncryptDecryptTxPn
func (obj *secureEntityCryptoEngineEncryptDecrypt) TxPn() SecureEntityCryptoEngineEncryptDecryptTxPn {
	if obj.obj.TxPn == nil {
		obj.obj.TxPn = NewSecureEntityCryptoEngineEncryptDecryptTxPn().msg()
	}
	if obj.txPnHolder == nil {
		obj.txPnHolder = &secureEntityCryptoEngineEncryptDecryptTxPn{obj: obj.obj.TxPn}
	}
	return obj.txPnHolder
}

// description is TBD
// TxPn returns a SecureEntityCryptoEngineEncryptDecryptTxPn
func (obj *secureEntityCryptoEngineEncryptDecrypt) HasTxPn() bool {
	return obj.obj.TxPn != nil
}

// description is TBD
// SetTxPn sets the SecureEntityCryptoEngineEncryptDecryptTxPn value in the SecureEntityCryptoEngineEncryptDecrypt object
func (obj *secureEntityCryptoEngineEncryptDecrypt) SetTxPn(value SecureEntityCryptoEngineEncryptDecryptTxPn) SecureEntityCryptoEngineEncryptDecrypt {

	obj.txPnHolder = nil
	obj.obj.TxPn = value.msg()

	return obj
}

// description is TBD
// HardwareAcceleration returns a SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
func (obj *secureEntityCryptoEngineEncryptDecrypt) HardwareAcceleration() SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration {
	if obj.obj.HardwareAcceleration == nil {
		obj.obj.HardwareAcceleration = NewSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration().msg()
	}
	if obj.hardwareAccelerationHolder == nil {
		obj.hardwareAccelerationHolder = &secureEntityCryptoEngineEncryptDecryptHardwareAcceleration{obj: obj.obj.HardwareAcceleration}
	}
	return obj.hardwareAccelerationHolder
}

// description is TBD
// HardwareAcceleration returns a SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
func (obj *secureEntityCryptoEngineEncryptDecrypt) HasHardwareAcceleration() bool {
	return obj.obj.HardwareAcceleration != nil
}

// description is TBD
// SetHardwareAcceleration sets the SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration value in the SecureEntityCryptoEngineEncryptDecrypt object
func (obj *secureEntityCryptoEngineEncryptDecrypt) SetHardwareAcceleration(value SecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) SecureEntityCryptoEngineEncryptDecrypt {

	obj.hardwareAccelerationHolder = nil
	obj.obj.HardwareAcceleration = value.msg()

	return obj
}

func (obj *secureEntityCryptoEngineEncryptDecrypt) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.TxPn != nil {

		obj.TxPn().validateObj(vObj, set_default)
	}

	if obj.obj.HardwareAcceleration != nil {

		obj.HardwareAcceleration().validateObj(vObj, set_default)
	}

}

func (obj *secureEntityCryptoEngineEncryptDecrypt) setDefault() {

}
