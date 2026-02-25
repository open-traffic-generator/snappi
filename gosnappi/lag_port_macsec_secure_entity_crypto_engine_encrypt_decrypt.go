package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt *****
type lagPortMacsecSecureEntityCryptoEngineEncryptDecrypt struct {
	validation
	obj                        *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt
	marshaller                 marshalLagPortMacsecSecureEntityCryptoEngineEncryptDecrypt
	unMarshaller               unMarshalLagPortMacsecSecureEntityCryptoEngineEncryptDecrypt
	txPnHolder                 LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn
	hardwareAccelerationHolder LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
}

func NewLagPortMacsecSecureEntityCryptoEngineEncryptDecrypt() LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt {
	obj := lagPortMacsecSecureEntityCryptoEngineEncryptDecrypt{obj: &otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt{}}
	obj.setDefault()
	return &obj
}

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) msg() *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt {
	return obj.obj
}

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) setMsg(msg *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshallagPortMacsecSecureEntityCryptoEngineEncryptDecrypt struct {
	obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecrypt
}

type marshalLagPortMacsecSecureEntityCryptoEngineEncryptDecrypt interface {
	// ToProto marshals LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt to protobuf object *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt
	ToProto() (*otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt, error)
	// ToPbText marshals LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt to YAML text
	ToYaml() (string, error)
	// ToJson marshals LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt to JSON text
	ToJson() (string, error)
}

type unMarshallagPortMacsecSecureEntityCryptoEngineEncryptDecrypt struct {
	obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecrypt
}

type unMarshalLagPortMacsecSecureEntityCryptoEngineEncryptDecrypt interface {
	// FromProto unmarshals LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt from protobuf object *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt
	FromProto(msg *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) (LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt, error)
	// FromPbText unmarshals LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt from JSON text
	FromJson(value string) error
}

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) Marshal() marshalLagPortMacsecSecureEntityCryptoEngineEncryptDecrypt {
	if obj.marshaller == nil {
		obj.marshaller = &marshallagPortMacsecSecureEntityCryptoEngineEncryptDecrypt{obj: obj}
	}
	return obj.marshaller
}

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) Unmarshal() unMarshalLagPortMacsecSecureEntityCryptoEngineEncryptDecrypt {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallagPortMacsecSecureEntityCryptoEngineEncryptDecrypt{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) ToProto() (*otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) FromProto(msg *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) (LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) ToPbText() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) FromPbText(value string) error {
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

func (m *marshallagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) ToYaml() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) FromYaml(value string) error {
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

func (m *marshallagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) ToJson() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) FromJson(value string) error {
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

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) Clone() (LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLagPortMacsecSecureEntityCryptoEngineEncryptDecrypt()
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

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) setNil() {
	obj.txPnHolder = nil
	obj.hardwareAccelerationHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt is the container for configuration of crypto engine of encrypt and decrypt type. Such engine can both encrypt transmitted packets and decrypt packets on arrival. It can have hardware acceleration for faster encryption/ decryption. As both encryption and decryption are possible, stateful (e.g. TCP) traffic can be sent/ received.
type LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt interface {
	Validation
	// msg marshals LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt to protobuf object *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt
	// and doesn't set defaults
	msg() *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt
	// setMsg unmarshals LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt from protobuf object *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt
	// and doesn't set defaults
	setMsg(*otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt
	// provides marshal interface
	Marshal() marshalLagPortMacsecSecureEntityCryptoEngineEncryptDecrypt
	// provides unmarshal interface
	Unmarshal() unMarshalLagPortMacsecSecureEntityCryptoEngineEncryptDecrypt
	// validate validates LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// TxPn returns LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn, set in LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt.
	// LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn is tx packet number(PN) configuration.
	TxPn() LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn
	// SetTxPn assigns LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn provided by user to LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt.
	// LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn is tx packet number(PN) configuration.
	SetTxPn(value LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt
	// HasTxPn checks if TxPn has been set in LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt
	HasTxPn() bool
	// HardwareAcceleration returns LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration, set in LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt.
	// LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration is hardware acceleration configuration for offloading MACsec processing to hardware.
	HardwareAcceleration() LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
	// SetHardwareAcceleration assigns LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration provided by user to LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt.
	// LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration is hardware acceleration configuration for offloading MACsec processing to hardware.
	SetHardwareAcceleration(value LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt
	// HasHardwareAcceleration checks if HardwareAcceleration has been set in LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt
	HasHardwareAcceleration() bool
	setNil()
}

// description is TBD
// TxPn returns a LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn
func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) TxPn() LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn {
	if obj.obj.TxPn == nil {
		obj.obj.TxPn = NewLagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn().msg()
	}
	if obj.txPnHolder == nil {
		obj.txPnHolder = &lagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn{obj: obj.obj.TxPn}
	}
	return obj.txPnHolder
}

// description is TBD
// TxPn returns a LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn
func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) HasTxPn() bool {
	return obj.obj.TxPn != nil
}

// description is TBD
// SetTxPn sets the LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn value in the LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt object
func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) SetTxPn(value LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt {

	obj.txPnHolder = nil
	obj.obj.TxPn = value.msg()

	return obj
}

// description is TBD
// HardwareAcceleration returns a LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) HardwareAcceleration() LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration {
	if obj.obj.HardwareAcceleration == nil {
		obj.obj.HardwareAcceleration = NewLagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration().msg()
	}
	if obj.hardwareAccelerationHolder == nil {
		obj.hardwareAccelerationHolder = &lagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration{obj: obj.obj.HardwareAcceleration}
	}
	return obj.hardwareAccelerationHolder
}

// description is TBD
// HardwareAcceleration returns a LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration
func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) HasHardwareAcceleration() bool {
	return obj.obj.HardwareAcceleration != nil
}

// description is TBD
// SetHardwareAcceleration sets the LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration value in the LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt object
func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) SetHardwareAcceleration(value LagPortMacsecSecureEntityCryptoEngineEncryptDecryptHardwareAcceleration) LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt {

	obj.hardwareAccelerationHolder = nil
	obj.obj.HardwareAcceleration = value.msg()

	return obj
}

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) validateObj(vObj *validation, set_default bool) {
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

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) setDefault() {

}
