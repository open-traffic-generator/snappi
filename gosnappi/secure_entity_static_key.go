package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SecureEntityStaticKey *****
type secureEntityStaticKey struct {
	validation
	obj          *otg.SecureEntityStaticKey
	marshaller   marshalSecureEntityStaticKey
	unMarshaller unMarshalSecureEntityStaticKey
	txHolder     SecureEntityStaticKeyTx
	rxHolder     SecureEntityStaticKeyRx
}

func NewSecureEntityStaticKey() SecureEntityStaticKey {
	obj := secureEntityStaticKey{obj: &otg.SecureEntityStaticKey{}}
	obj.setDefault()
	return &obj
}

func (obj *secureEntityStaticKey) msg() *otg.SecureEntityStaticKey {
	return obj.obj
}

func (obj *secureEntityStaticKey) setMsg(msg *otg.SecureEntityStaticKey) SecureEntityStaticKey {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsecureEntityStaticKey struct {
	obj *secureEntityStaticKey
}

type marshalSecureEntityStaticKey interface {
	// ToProto marshals SecureEntityStaticKey to protobuf object *otg.SecureEntityStaticKey
	ToProto() (*otg.SecureEntityStaticKey, error)
	// ToPbText marshals SecureEntityStaticKey to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SecureEntityStaticKey to YAML text
	ToYaml() (string, error)
	// ToJson marshals SecureEntityStaticKey to JSON text
	ToJson() (string, error)
}

type unMarshalsecureEntityStaticKey struct {
	obj *secureEntityStaticKey
}

type unMarshalSecureEntityStaticKey interface {
	// FromProto unmarshals SecureEntityStaticKey from protobuf object *otg.SecureEntityStaticKey
	FromProto(msg *otg.SecureEntityStaticKey) (SecureEntityStaticKey, error)
	// FromPbText unmarshals SecureEntityStaticKey from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SecureEntityStaticKey from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SecureEntityStaticKey from JSON text
	FromJson(value string) error
}

func (obj *secureEntityStaticKey) Marshal() marshalSecureEntityStaticKey {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsecureEntityStaticKey{obj: obj}
	}
	return obj.marshaller
}

func (obj *secureEntityStaticKey) Unmarshal() unMarshalSecureEntityStaticKey {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsecureEntityStaticKey{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsecureEntityStaticKey) ToProto() (*otg.SecureEntityStaticKey, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsecureEntityStaticKey) FromProto(msg *otg.SecureEntityStaticKey) (SecureEntityStaticKey, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsecureEntityStaticKey) ToPbText() (string, error) {
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

func (m *unMarshalsecureEntityStaticKey) FromPbText(value string) error {
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

func (m *marshalsecureEntityStaticKey) ToYaml() (string, error) {
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

func (m *unMarshalsecureEntityStaticKey) FromYaml(value string) error {
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

func (m *marshalsecureEntityStaticKey) ToJson() (string, error) {
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

func (m *unMarshalsecureEntityStaticKey) FromJson(value string) error {
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

func (obj *secureEntityStaticKey) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *secureEntityStaticKey) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *secureEntityStaticKey) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *secureEntityStaticKey) Clone() (SecureEntityStaticKey, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSecureEntityStaticKey()
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

func (obj *secureEntityStaticKey) setNil() {
	obj.txHolder = nil
	obj.rxHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// SecureEntityStaticKey is a container of static key properties for a secure entity(SecY). This configuration is applicable when no dynamic key management protocol i.e. MACsec key agreement(MKA) is configured. If MKA is configured, any static key configuration is not applicable.
type SecureEntityStaticKey interface {
	Validation
	// msg marshals SecureEntityStaticKey to protobuf object *otg.SecureEntityStaticKey
	// and doesn't set defaults
	msg() *otg.SecureEntityStaticKey
	// setMsg unmarshals SecureEntityStaticKey from protobuf object *otg.SecureEntityStaticKey
	// and doesn't set defaults
	setMsg(*otg.SecureEntityStaticKey) SecureEntityStaticKey
	// provides marshal interface
	Marshal() marshalSecureEntityStaticKey
	// provides unmarshal interface
	Unmarshal() unMarshalSecureEntityStaticKey
	// validate validates SecureEntityStaticKey
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SecureEntityStaticKey, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// CipherSuite returns SecureEntityStaticKeyCipherSuiteEnum, set in SecureEntityStaticKey
	CipherSuite() SecureEntityStaticKeyCipherSuiteEnum
	// SetCipherSuite assigns SecureEntityStaticKeyCipherSuiteEnum provided by user to SecureEntityStaticKey
	SetCipherSuite(value SecureEntityStaticKeyCipherSuiteEnum) SecureEntityStaticKey
	// HasCipherSuite checks if CipherSuite has been set in SecureEntityStaticKey
	HasCipherSuite() bool
	// Confidentiality returns bool, set in SecureEntityStaticKey.
	Confidentiality() bool
	// SetConfidentiality assigns bool provided by user to SecureEntityStaticKey
	SetConfidentiality(value bool) SecureEntityStaticKey
	// HasConfidentiality checks if Confidentiality has been set in SecureEntityStaticKey
	HasConfidentiality() bool
	// ConfidentialityOffset returns SecureEntityStaticKeyConfidentialityOffsetEnum, set in SecureEntityStaticKey
	ConfidentialityOffset() SecureEntityStaticKeyConfidentialityOffsetEnum
	// SetConfidentialityOffset assigns SecureEntityStaticKeyConfidentialityOffsetEnum provided by user to SecureEntityStaticKey
	SetConfidentialityOffset(value SecureEntityStaticKeyConfidentialityOffsetEnum) SecureEntityStaticKey
	// HasConfidentialityOffset checks if ConfidentialityOffset has been set in SecureEntityStaticKey
	HasConfidentialityOffset() bool
	// Tx returns SecureEntityStaticKeyTx, set in SecureEntityStaticKey.
	// SecureEntityStaticKeyTx is a container of static key Tx properties.
	Tx() SecureEntityStaticKeyTx
	// SetTx assigns SecureEntityStaticKeyTx provided by user to SecureEntityStaticKey.
	// SecureEntityStaticKeyTx is a container of static key Tx properties.
	SetTx(value SecureEntityStaticKeyTx) SecureEntityStaticKey
	// HasTx checks if Tx has been set in SecureEntityStaticKey
	HasTx() bool
	// Rx returns SecureEntityStaticKeyRx, set in SecureEntityStaticKey.
	// SecureEntityStaticKeyRx is a container of static key Rx properties.
	Rx() SecureEntityStaticKeyRx
	// SetRx assigns SecureEntityStaticKeyRx provided by user to SecureEntityStaticKey.
	// SecureEntityStaticKeyRx is a container of static key Rx properties.
	SetRx(value SecureEntityStaticKeyRx) SecureEntityStaticKey
	// HasRx checks if Rx has been set in SecureEntityStaticKey
	HasRx() bool
	setNil()
}

type SecureEntityStaticKeyCipherSuiteEnum string

// Enum of CipherSuite on SecureEntityStaticKey
var SecureEntityStaticKeyCipherSuite = struct {
	GCM_AES_128     SecureEntityStaticKeyCipherSuiteEnum
	GCM_AES_256     SecureEntityStaticKeyCipherSuiteEnum
	GCM_AES_XPN_128 SecureEntityStaticKeyCipherSuiteEnum
	GCM_AES_XPN_256 SecureEntityStaticKeyCipherSuiteEnum
}{
	GCM_AES_128:     SecureEntityStaticKeyCipherSuiteEnum("gcm_aes_128"),
	GCM_AES_256:     SecureEntityStaticKeyCipherSuiteEnum("gcm_aes_256"),
	GCM_AES_XPN_128: SecureEntityStaticKeyCipherSuiteEnum("gcm_aes_xpn_128"),
	GCM_AES_XPN_256: SecureEntityStaticKeyCipherSuiteEnum("gcm_aes_xpn_256"),
}

func (obj *secureEntityStaticKey) CipherSuite() SecureEntityStaticKeyCipherSuiteEnum {
	return SecureEntityStaticKeyCipherSuiteEnum(obj.obj.CipherSuite.Enum().String())
}

// The cipher suite. Choose one from GCM-AES-128 GCM-AES-128 GCM-AES-256 GCM-AES-XPN-128 GCM-AES-XPN-256
// CipherSuite returns a string
func (obj *secureEntityStaticKey) HasCipherSuite() bool {
	return obj.obj.CipherSuite != nil
}

func (obj *secureEntityStaticKey) SetCipherSuite(value SecureEntityStaticKeyCipherSuiteEnum) SecureEntityStaticKey {
	intValue, ok := otg.SecureEntityStaticKey_CipherSuite_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on SecureEntityStaticKeyCipherSuiteEnum", string(value)))
		return obj
	}
	enumValue := otg.SecureEntityStaticKey_CipherSuite_Enum(intValue)
	obj.obj.CipherSuite = &enumValue

	return obj
}

// Encrypt or not.
// Confidentiality returns a bool
func (obj *secureEntityStaticKey) Confidentiality() bool {

	return *obj.obj.Confidentiality

}

// Encrypt or not.
// Confidentiality returns a bool
func (obj *secureEntityStaticKey) HasConfidentiality() bool {
	return obj.obj.Confidentiality != nil
}

// Encrypt or not.
// SetConfidentiality sets the bool value in the SecureEntityStaticKey object
func (obj *secureEntityStaticKey) SetConfidentiality(value bool) SecureEntityStaticKey {

	obj.obj.Confidentiality = &value
	return obj
}

type SecureEntityStaticKeyConfidentialityOffsetEnum string

// Enum of ConfidentialityOffset on SecureEntityStaticKey
var SecureEntityStaticKeyConfidentialityOffset = struct {
	ZERO   SecureEntityStaticKeyConfidentialityOffsetEnum
	THIRTY SecureEntityStaticKeyConfidentialityOffsetEnum
	FIFTY  SecureEntityStaticKeyConfidentialityOffsetEnum
}{
	ZERO:   SecureEntityStaticKeyConfidentialityOffsetEnum("zero"),
	THIRTY: SecureEntityStaticKeyConfidentialityOffsetEnum("thirty"),
	FIFTY:  SecureEntityStaticKeyConfidentialityOffsetEnum("fifty"),
}

func (obj *secureEntityStaticKey) ConfidentialityOffset() SecureEntityStaticKeyConfidentialityOffsetEnum {
	return SecureEntityStaticKeyConfidentialityOffsetEnum(obj.obj.ConfidentialityOffset.Enum().String())
}

// Confidentiality offset.
// ConfidentialityOffset returns a string
func (obj *secureEntityStaticKey) HasConfidentialityOffset() bool {
	return obj.obj.ConfidentialityOffset != nil
}

func (obj *secureEntityStaticKey) SetConfidentialityOffset(value SecureEntityStaticKeyConfidentialityOffsetEnum) SecureEntityStaticKey {
	intValue, ok := otg.SecureEntityStaticKey_ConfidentialityOffset_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on SecureEntityStaticKeyConfidentialityOffsetEnum", string(value)))
		return obj
	}
	enumValue := otg.SecureEntityStaticKey_ConfidentialityOffset_Enum(intValue)
	obj.obj.ConfidentialityOffset = &enumValue

	return obj
}

// Tx properties of SecY.
// Tx returns a SecureEntityStaticKeyTx
func (obj *secureEntityStaticKey) Tx() SecureEntityStaticKeyTx {
	if obj.obj.Tx == nil {
		obj.obj.Tx = NewSecureEntityStaticKeyTx().msg()
	}
	if obj.txHolder == nil {
		obj.txHolder = &secureEntityStaticKeyTx{obj: obj.obj.Tx}
	}
	return obj.txHolder
}

// Tx properties of SecY.
// Tx returns a SecureEntityStaticKeyTx
func (obj *secureEntityStaticKey) HasTx() bool {
	return obj.obj.Tx != nil
}

// Tx properties of SecY.
// SetTx sets the SecureEntityStaticKeyTx value in the SecureEntityStaticKey object
func (obj *secureEntityStaticKey) SetTx(value SecureEntityStaticKeyTx) SecureEntityStaticKey {

	obj.txHolder = nil
	obj.obj.Tx = value.msg()

	return obj
}

// Rx properties of SecY.
// Rx returns a SecureEntityStaticKeyRx
func (obj *secureEntityStaticKey) Rx() SecureEntityStaticKeyRx {
	if obj.obj.Rx == nil {
		obj.obj.Rx = NewSecureEntityStaticKeyRx().msg()
	}
	if obj.rxHolder == nil {
		obj.rxHolder = &secureEntityStaticKeyRx{obj: obj.obj.Rx}
	}
	return obj.rxHolder
}

// Rx properties of SecY.
// Rx returns a SecureEntityStaticKeyRx
func (obj *secureEntityStaticKey) HasRx() bool {
	return obj.obj.Rx != nil
}

// Rx properties of SecY.
// SetRx sets the SecureEntityStaticKeyRx value in the SecureEntityStaticKey object
func (obj *secureEntityStaticKey) SetRx(value SecureEntityStaticKeyRx) SecureEntityStaticKey {

	obj.rxHolder = nil
	obj.obj.Rx = value.msg()

	return obj
}

func (obj *secureEntityStaticKey) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Tx != nil {

		obj.Tx().validateObj(vObj, set_default)
	}

	if obj.obj.Rx != nil {

		obj.Rx().validateObj(vObj, set_default)
	}

}

func (obj *secureEntityStaticKey) setDefault() {
	if obj.obj.CipherSuite == nil {
		obj.SetCipherSuite(SecureEntityStaticKeyCipherSuite.GCM_AES_128)

	}
	if obj.obj.Confidentiality == nil {
		obj.SetConfidentiality(true)
	}
	if obj.obj.ConfidentialityOffset == nil {
		obj.SetConfidentialityOffset(SecureEntityStaticKeyConfidentialityOffset.ZERO)

	}

}
