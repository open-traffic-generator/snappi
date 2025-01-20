package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecSecYBasicKeyGenerationStatic *****
type macsecSecYBasicKeyGenerationStatic struct {
	validation
	obj          *otg.MacsecSecYBasicKeyGenerationStatic
	marshaller   marshalMacsecSecYBasicKeyGenerationStatic
	unMarshaller unMarshalMacsecSecYBasicKeyGenerationStatic
}

func NewMacsecSecYBasicKeyGenerationStatic() MacsecSecYBasicKeyGenerationStatic {
	obj := macsecSecYBasicKeyGenerationStatic{obj: &otg.MacsecSecYBasicKeyGenerationStatic{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecSecYBasicKeyGenerationStatic) msg() *otg.MacsecSecYBasicKeyGenerationStatic {
	return obj.obj
}

func (obj *macsecSecYBasicKeyGenerationStatic) setMsg(msg *otg.MacsecSecYBasicKeyGenerationStatic) MacsecSecYBasicKeyGenerationStatic {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecSecYBasicKeyGenerationStatic struct {
	obj *macsecSecYBasicKeyGenerationStatic
}

type marshalMacsecSecYBasicKeyGenerationStatic interface {
	// ToProto marshals MacsecSecYBasicKeyGenerationStatic to protobuf object *otg.MacsecSecYBasicKeyGenerationStatic
	ToProto() (*otg.MacsecSecYBasicKeyGenerationStatic, error)
	// ToPbText marshals MacsecSecYBasicKeyGenerationStatic to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecSecYBasicKeyGenerationStatic to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecSecYBasicKeyGenerationStatic to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecSecYBasicKeyGenerationStatic struct {
	obj *macsecSecYBasicKeyGenerationStatic
}

type unMarshalMacsecSecYBasicKeyGenerationStatic interface {
	// FromProto unmarshals MacsecSecYBasicKeyGenerationStatic from protobuf object *otg.MacsecSecYBasicKeyGenerationStatic
	FromProto(msg *otg.MacsecSecYBasicKeyGenerationStatic) (MacsecSecYBasicKeyGenerationStatic, error)
	// FromPbText unmarshals MacsecSecYBasicKeyGenerationStatic from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecSecYBasicKeyGenerationStatic from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecSecYBasicKeyGenerationStatic from JSON text
	FromJson(value string) error
}

func (obj *macsecSecYBasicKeyGenerationStatic) Marshal() marshalMacsecSecYBasicKeyGenerationStatic {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecSecYBasicKeyGenerationStatic{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecSecYBasicKeyGenerationStatic) Unmarshal() unMarshalMacsecSecYBasicKeyGenerationStatic {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecSecYBasicKeyGenerationStatic{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecSecYBasicKeyGenerationStatic) ToProto() (*otg.MacsecSecYBasicKeyGenerationStatic, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecSecYBasicKeyGenerationStatic) FromProto(msg *otg.MacsecSecYBasicKeyGenerationStatic) (MacsecSecYBasicKeyGenerationStatic, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecSecYBasicKeyGenerationStatic) ToPbText() (string, error) {
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

func (m *unMarshalmacsecSecYBasicKeyGenerationStatic) FromPbText(value string) error {
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

func (m *marshalmacsecSecYBasicKeyGenerationStatic) ToYaml() (string, error) {
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

func (m *unMarshalmacsecSecYBasicKeyGenerationStatic) FromYaml(value string) error {
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

func (m *marshalmacsecSecYBasicKeyGenerationStatic) ToJson() (string, error) {
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

func (m *unMarshalmacsecSecYBasicKeyGenerationStatic) FromJson(value string) error {
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

func (obj *macsecSecYBasicKeyGenerationStatic) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecSecYBasicKeyGenerationStatic) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecSecYBasicKeyGenerationStatic) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecSecYBasicKeyGenerationStatic) Clone() (MacsecSecYBasicKeyGenerationStatic, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecSecYBasicKeyGenerationStatic()
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

// MacsecSecYBasicKeyGenerationStatic is the container for static key settings.
type MacsecSecYBasicKeyGenerationStatic interface {
	Validation
	// msg marshals MacsecSecYBasicKeyGenerationStatic to protobuf object *otg.MacsecSecYBasicKeyGenerationStatic
	// and doesn't set defaults
	msg() *otg.MacsecSecYBasicKeyGenerationStatic
	// setMsg unmarshals MacsecSecYBasicKeyGenerationStatic from protobuf object *otg.MacsecSecYBasicKeyGenerationStatic
	// and doesn't set defaults
	setMsg(*otg.MacsecSecYBasicKeyGenerationStatic) MacsecSecYBasicKeyGenerationStatic
	// provides marshal interface
	Marshal() marshalMacsecSecYBasicKeyGenerationStatic
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecSecYBasicKeyGenerationStatic
	// validate validates MacsecSecYBasicKeyGenerationStatic
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecSecYBasicKeyGenerationStatic, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// CipherSuite returns MacsecSecYBasicKeyGenerationStaticCipherSuiteEnum, set in MacsecSecYBasicKeyGenerationStatic
	CipherSuite() MacsecSecYBasicKeyGenerationStaticCipherSuiteEnum
	// SetCipherSuite assigns MacsecSecYBasicKeyGenerationStaticCipherSuiteEnum provided by user to MacsecSecYBasicKeyGenerationStatic
	SetCipherSuite(value MacsecSecYBasicKeyGenerationStaticCipherSuiteEnum) MacsecSecYBasicKeyGenerationStatic
	// HasCipherSuite checks if CipherSuite has been set in MacsecSecYBasicKeyGenerationStatic
	HasCipherSuite() bool
	// ConfidentialityOffset returns MacsecSecYBasicKeyGenerationStaticConfidentialityOffsetEnum, set in MacsecSecYBasicKeyGenerationStatic
	ConfidentialityOffset() MacsecSecYBasicKeyGenerationStaticConfidentialityOffsetEnum
	// SetConfidentialityOffset assigns MacsecSecYBasicKeyGenerationStaticConfidentialityOffsetEnum provided by user to MacsecSecYBasicKeyGenerationStatic
	SetConfidentialityOffset(value MacsecSecYBasicKeyGenerationStaticConfidentialityOffsetEnum) MacsecSecYBasicKeyGenerationStatic
	// HasConfidentialityOffset checks if ConfidentialityOffset has been set in MacsecSecYBasicKeyGenerationStatic
	HasConfidentialityOffset() bool
}

type MacsecSecYBasicKeyGenerationStaticCipherSuiteEnum string

// Enum of CipherSuite on MacsecSecYBasicKeyGenerationStatic
var MacsecSecYBasicKeyGenerationStaticCipherSuite = struct {
	GCM_AES_128     MacsecSecYBasicKeyGenerationStaticCipherSuiteEnum
	GCM_AES_256     MacsecSecYBasicKeyGenerationStaticCipherSuiteEnum
	GCM_AES_XPN_128 MacsecSecYBasicKeyGenerationStaticCipherSuiteEnum
	GCM_AES_XPN_256 MacsecSecYBasicKeyGenerationStaticCipherSuiteEnum
}{
	GCM_AES_128:     MacsecSecYBasicKeyGenerationStaticCipherSuiteEnum("gcm_aes_128"),
	GCM_AES_256:     MacsecSecYBasicKeyGenerationStaticCipherSuiteEnum("gcm_aes_256"),
	GCM_AES_XPN_128: MacsecSecYBasicKeyGenerationStaticCipherSuiteEnum("gcm_aes_xpn_128"),
	GCM_AES_XPN_256: MacsecSecYBasicKeyGenerationStaticCipherSuiteEnum("gcm_aes_xpn_256"),
}

func (obj *macsecSecYBasicKeyGenerationStatic) CipherSuite() MacsecSecYBasicKeyGenerationStaticCipherSuiteEnum {
	return MacsecSecYBasicKeyGenerationStaticCipherSuiteEnum(obj.obj.CipherSuite.Enum().String())
}

// The cipher suite. Choose one from GCM-AES-128 GCM-AES-128 GCM-AES-256 GCM-AES-XPN-128 GCM-AES-XPN-256
// CipherSuite returns a string
func (obj *macsecSecYBasicKeyGenerationStatic) HasCipherSuite() bool {
	return obj.obj.CipherSuite != nil
}

func (obj *macsecSecYBasicKeyGenerationStatic) SetCipherSuite(value MacsecSecYBasicKeyGenerationStaticCipherSuiteEnum) MacsecSecYBasicKeyGenerationStatic {
	intValue, ok := otg.MacsecSecYBasicKeyGenerationStatic_CipherSuite_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecSecYBasicKeyGenerationStaticCipherSuiteEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecSecYBasicKeyGenerationStatic_CipherSuite_Enum(intValue)
	obj.obj.CipherSuite = &enumValue

	return obj
}

type MacsecSecYBasicKeyGenerationStaticConfidentialityOffsetEnum string

// Enum of ConfidentialityOffset on MacsecSecYBasicKeyGenerationStatic
var MacsecSecYBasicKeyGenerationStaticConfidentialityOffset = struct {
	ZERO   MacsecSecYBasicKeyGenerationStaticConfidentialityOffsetEnum
	THIRTY MacsecSecYBasicKeyGenerationStaticConfidentialityOffsetEnum
	FIFTY  MacsecSecYBasicKeyGenerationStaticConfidentialityOffsetEnum
}{
	ZERO:   MacsecSecYBasicKeyGenerationStaticConfidentialityOffsetEnum("zero"),
	THIRTY: MacsecSecYBasicKeyGenerationStaticConfidentialityOffsetEnum("thirty"),
	FIFTY:  MacsecSecYBasicKeyGenerationStaticConfidentialityOffsetEnum("fifty"),
}

func (obj *macsecSecYBasicKeyGenerationStatic) ConfidentialityOffset() MacsecSecYBasicKeyGenerationStaticConfidentialityOffsetEnum {
	return MacsecSecYBasicKeyGenerationStaticConfidentialityOffsetEnum(obj.obj.ConfidentialityOffset.Enum().String())
}

// Confidentiality offset.
// ConfidentialityOffset returns a string
func (obj *macsecSecYBasicKeyGenerationStatic) HasConfidentialityOffset() bool {
	return obj.obj.ConfidentialityOffset != nil
}

func (obj *macsecSecYBasicKeyGenerationStatic) SetConfidentialityOffset(value MacsecSecYBasicKeyGenerationStaticConfidentialityOffsetEnum) MacsecSecYBasicKeyGenerationStatic {
	intValue, ok := otg.MacsecSecYBasicKeyGenerationStatic_ConfidentialityOffset_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecSecYBasicKeyGenerationStaticConfidentialityOffsetEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecSecYBasicKeyGenerationStatic_ConfidentialityOffset_Enum(intValue)
	obj.obj.ConfidentialityOffset = &enumValue

	return obj
}

func (obj *macsecSecYBasicKeyGenerationStatic) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *macsecSecYBasicKeyGenerationStatic) setDefault() {
	if obj.obj.CipherSuite == nil {
		obj.SetCipherSuite(MacsecSecYBasicKeyGenerationStaticCipherSuite.GCM_AES_128)

	}
	if obj.obj.ConfidentialityOffset == nil {
		obj.SetConfidentialityOffset(MacsecSecYBasicKeyGenerationStaticConfidentialityOffset.ZERO)

	}

}
