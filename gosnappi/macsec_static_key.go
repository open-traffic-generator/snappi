package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecStaticKey *****
type macsecStaticKey struct {
	validation
	obj          *otg.MacsecStaticKey
	marshaller   marshalMacsecStaticKey
	unMarshaller unMarshalMacsecStaticKey
}

func NewMacsecStaticKey() MacsecStaticKey {
	obj := macsecStaticKey{obj: &otg.MacsecStaticKey{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecStaticKey) msg() *otg.MacsecStaticKey {
	return obj.obj
}

func (obj *macsecStaticKey) setMsg(msg *otg.MacsecStaticKey) MacsecStaticKey {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecStaticKey struct {
	obj *macsecStaticKey
}

type marshalMacsecStaticKey interface {
	// ToProto marshals MacsecStaticKey to protobuf object *otg.MacsecStaticKey
	ToProto() (*otg.MacsecStaticKey, error)
	// ToPbText marshals MacsecStaticKey to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecStaticKey to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecStaticKey to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecStaticKey struct {
	obj *macsecStaticKey
}

type unMarshalMacsecStaticKey interface {
	// FromProto unmarshals MacsecStaticKey from protobuf object *otg.MacsecStaticKey
	FromProto(msg *otg.MacsecStaticKey) (MacsecStaticKey, error)
	// FromPbText unmarshals MacsecStaticKey from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecStaticKey from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecStaticKey from JSON text
	FromJson(value string) error
}

func (obj *macsecStaticKey) Marshal() marshalMacsecStaticKey {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecStaticKey{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecStaticKey) Unmarshal() unMarshalMacsecStaticKey {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecStaticKey{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecStaticKey) ToProto() (*otg.MacsecStaticKey, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecStaticKey) FromProto(msg *otg.MacsecStaticKey) (MacsecStaticKey, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecStaticKey) ToPbText() (string, error) {
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

func (m *unMarshalmacsecStaticKey) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalmacsecStaticKey) ToYaml() (string, error) {
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

func (m *unMarshalmacsecStaticKey) FromYaml(value string) error {
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

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalmacsecStaticKey) ToJson() (string, error) {
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

func (m *unMarshalmacsecStaticKey) FromJson(value string) error {
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

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *macsecStaticKey) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecStaticKey) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecStaticKey) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecStaticKey) Clone() (MacsecStaticKey, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecStaticKey()
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

// MacsecStaticKey is a container of static key properties for a secure entity(SecY). This configuration is applicable when no dynamic key management protocol i.e. MACsec key agreement(MKA) is configured. If MKA is configured, any static key configuration is not applicable.
type MacsecStaticKey interface {
	Validation
	// msg marshals MacsecStaticKey to protobuf object *otg.MacsecStaticKey
	// and doesn't set defaults
	msg() *otg.MacsecStaticKey
	// setMsg unmarshals MacsecStaticKey from protobuf object *otg.MacsecStaticKey
	// and doesn't set defaults
	setMsg(*otg.MacsecStaticKey) MacsecStaticKey
	// provides marshal interface
	Marshal() marshalMacsecStaticKey
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecStaticKey
	// validate validates MacsecStaticKey
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecStaticKey, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// CipherSuite returns MacsecStaticKeyCipherSuiteEnum, set in MacsecStaticKey
	CipherSuite() MacsecStaticKeyCipherSuiteEnum
	// SetCipherSuite assigns MacsecStaticKeyCipherSuiteEnum provided by user to MacsecStaticKey
	SetCipherSuite(value MacsecStaticKeyCipherSuiteEnum) MacsecStaticKey
	// HasCipherSuite checks if CipherSuite has been set in MacsecStaticKey
	HasCipherSuite() bool
	// ConfidentialityOffset returns MacsecStaticKeyConfidentialityOffsetEnum, set in MacsecStaticKey
	ConfidentialityOffset() MacsecStaticKeyConfidentialityOffsetEnum
	// SetConfidentialityOffset assigns MacsecStaticKeyConfidentialityOffsetEnum provided by user to MacsecStaticKey
	SetConfidentialityOffset(value MacsecStaticKeyConfidentialityOffsetEnum) MacsecStaticKey
	// HasConfidentialityOffset checks if ConfidentialityOffset has been set in MacsecStaticKey
	HasConfidentialityOffset() bool
}

type MacsecStaticKeyCipherSuiteEnum string

// Enum of CipherSuite on MacsecStaticKey
var MacsecStaticKeyCipherSuite = struct {
	GCM_AES_128     MacsecStaticKeyCipherSuiteEnum
	GCM_AES_256     MacsecStaticKeyCipherSuiteEnum
	GCM_AES_XPN_128 MacsecStaticKeyCipherSuiteEnum
	GCM_AES_XPN_256 MacsecStaticKeyCipherSuiteEnum
}{
	GCM_AES_128:     MacsecStaticKeyCipherSuiteEnum("gcm_aes_128"),
	GCM_AES_256:     MacsecStaticKeyCipherSuiteEnum("gcm_aes_256"),
	GCM_AES_XPN_128: MacsecStaticKeyCipherSuiteEnum("gcm_aes_xpn_128"),
	GCM_AES_XPN_256: MacsecStaticKeyCipherSuiteEnum("gcm_aes_xpn_256"),
}

func (obj *macsecStaticKey) CipherSuite() MacsecStaticKeyCipherSuiteEnum {
	return MacsecStaticKeyCipherSuiteEnum(obj.obj.CipherSuite.Enum().String())
}

// The cipher suite. Choose one from GCM-AES-128 GCM-AES-128 GCM-AES-256 GCM-AES-XPN-128 GCM-AES-XPN-256
// CipherSuite returns a string
func (obj *macsecStaticKey) HasCipherSuite() bool {
	return obj.obj.CipherSuite != nil
}

func (obj *macsecStaticKey) SetCipherSuite(value MacsecStaticKeyCipherSuiteEnum) MacsecStaticKey {
	intValue, ok := otg.MacsecStaticKey_CipherSuite_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecStaticKeyCipherSuiteEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecStaticKey_CipherSuite_Enum(intValue)
	obj.obj.CipherSuite = &enumValue

	return obj
}

type MacsecStaticKeyConfidentialityOffsetEnum string

// Enum of ConfidentialityOffset on MacsecStaticKey
var MacsecStaticKeyConfidentialityOffset = struct {
	ZERO   MacsecStaticKeyConfidentialityOffsetEnum
	THIRTY MacsecStaticKeyConfidentialityOffsetEnum
	FIFTY  MacsecStaticKeyConfidentialityOffsetEnum
}{
	ZERO:   MacsecStaticKeyConfidentialityOffsetEnum("zero"),
	THIRTY: MacsecStaticKeyConfidentialityOffsetEnum("thirty"),
	FIFTY:  MacsecStaticKeyConfidentialityOffsetEnum("fifty"),
}

func (obj *macsecStaticKey) ConfidentialityOffset() MacsecStaticKeyConfidentialityOffsetEnum {
	return MacsecStaticKeyConfidentialityOffsetEnum(obj.obj.ConfidentialityOffset.Enum().String())
}

// Confidentiality offset.
// ConfidentialityOffset returns a string
func (obj *macsecStaticKey) HasConfidentialityOffset() bool {
	return obj.obj.ConfidentialityOffset != nil
}

func (obj *macsecStaticKey) SetConfidentialityOffset(value MacsecStaticKeyConfidentialityOffsetEnum) MacsecStaticKey {
	intValue, ok := otg.MacsecStaticKey_ConfidentialityOffset_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecStaticKeyConfidentialityOffsetEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecStaticKey_ConfidentialityOffset_Enum(intValue)
	obj.obj.ConfidentialityOffset = &enumValue

	return obj
}

func (obj *macsecStaticKey) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *macsecStaticKey) setDefault() {
	if obj.obj.CipherSuite == nil {
		obj.SetCipherSuite(MacsecStaticKeyCipherSuite.GCM_AES_128)

	}
	if obj.obj.ConfidentialityOffset == nil {
		obj.SetConfidentialityOffset(MacsecStaticKeyConfidentialityOffset.ZERO)

	}

}
