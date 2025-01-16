package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecBasicKeyGenerationStatic *****
type macsecBasicKeyGenerationStatic struct {
	validation
	obj          *otg.MacsecBasicKeyGenerationStatic
	marshaller   marshalMacsecBasicKeyGenerationStatic
	unMarshaller unMarshalMacsecBasicKeyGenerationStatic
}

func NewMacsecBasicKeyGenerationStatic() MacsecBasicKeyGenerationStatic {
	obj := macsecBasicKeyGenerationStatic{obj: &otg.MacsecBasicKeyGenerationStatic{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecBasicKeyGenerationStatic) msg() *otg.MacsecBasicKeyGenerationStatic {
	return obj.obj
}

func (obj *macsecBasicKeyGenerationStatic) setMsg(msg *otg.MacsecBasicKeyGenerationStatic) MacsecBasicKeyGenerationStatic {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecBasicKeyGenerationStatic struct {
	obj *macsecBasicKeyGenerationStatic
}

type marshalMacsecBasicKeyGenerationStatic interface {
	// ToProto marshals MacsecBasicKeyGenerationStatic to protobuf object *otg.MacsecBasicKeyGenerationStatic
	ToProto() (*otg.MacsecBasicKeyGenerationStatic, error)
	// ToPbText marshals MacsecBasicKeyGenerationStatic to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecBasicKeyGenerationStatic to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecBasicKeyGenerationStatic to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecBasicKeyGenerationStatic struct {
	obj *macsecBasicKeyGenerationStatic
}

type unMarshalMacsecBasicKeyGenerationStatic interface {
	// FromProto unmarshals MacsecBasicKeyGenerationStatic from protobuf object *otg.MacsecBasicKeyGenerationStatic
	FromProto(msg *otg.MacsecBasicKeyGenerationStatic) (MacsecBasicKeyGenerationStatic, error)
	// FromPbText unmarshals MacsecBasicKeyGenerationStatic from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecBasicKeyGenerationStatic from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecBasicKeyGenerationStatic from JSON text
	FromJson(value string) error
}

func (obj *macsecBasicKeyGenerationStatic) Marshal() marshalMacsecBasicKeyGenerationStatic {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecBasicKeyGenerationStatic{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecBasicKeyGenerationStatic) Unmarshal() unMarshalMacsecBasicKeyGenerationStatic {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecBasicKeyGenerationStatic{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecBasicKeyGenerationStatic) ToProto() (*otg.MacsecBasicKeyGenerationStatic, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecBasicKeyGenerationStatic) FromProto(msg *otg.MacsecBasicKeyGenerationStatic) (MacsecBasicKeyGenerationStatic, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecBasicKeyGenerationStatic) ToPbText() (string, error) {
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

func (m *unMarshalmacsecBasicKeyGenerationStatic) FromPbText(value string) error {
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

func (m *marshalmacsecBasicKeyGenerationStatic) ToYaml() (string, error) {
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

func (m *unMarshalmacsecBasicKeyGenerationStatic) FromYaml(value string) error {
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

func (m *marshalmacsecBasicKeyGenerationStatic) ToJson() (string, error) {
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

func (m *unMarshalmacsecBasicKeyGenerationStatic) FromJson(value string) error {
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

func (obj *macsecBasicKeyGenerationStatic) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecBasicKeyGenerationStatic) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecBasicKeyGenerationStatic) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecBasicKeyGenerationStatic) Clone() (MacsecBasicKeyGenerationStatic, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecBasicKeyGenerationStatic()
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

// MacsecBasicKeyGenerationStatic is the container for static key settings.
type MacsecBasicKeyGenerationStatic interface {
	Validation
	// msg marshals MacsecBasicKeyGenerationStatic to protobuf object *otg.MacsecBasicKeyGenerationStatic
	// and doesn't set defaults
	msg() *otg.MacsecBasicKeyGenerationStatic
	// setMsg unmarshals MacsecBasicKeyGenerationStatic from protobuf object *otg.MacsecBasicKeyGenerationStatic
	// and doesn't set defaults
	setMsg(*otg.MacsecBasicKeyGenerationStatic) MacsecBasicKeyGenerationStatic
	// provides marshal interface
	Marshal() marshalMacsecBasicKeyGenerationStatic
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecBasicKeyGenerationStatic
	// validate validates MacsecBasicKeyGenerationStatic
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecBasicKeyGenerationStatic, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// CipherSuite returns MacsecBasicKeyGenerationStaticCipherSuiteEnum, set in MacsecBasicKeyGenerationStatic
	CipherSuite() MacsecBasicKeyGenerationStaticCipherSuiteEnum
	// SetCipherSuite assigns MacsecBasicKeyGenerationStaticCipherSuiteEnum provided by user to MacsecBasicKeyGenerationStatic
	SetCipherSuite(value MacsecBasicKeyGenerationStaticCipherSuiteEnum) MacsecBasicKeyGenerationStatic
	// HasCipherSuite checks if CipherSuite has been set in MacsecBasicKeyGenerationStatic
	HasCipherSuite() bool
	// ConfidentialtyOffset returns int32, set in MacsecBasicKeyGenerationStatic.
	ConfidentialtyOffset() int32
	// SetConfidentialtyOffset assigns int32 provided by user to MacsecBasicKeyGenerationStatic
	SetConfidentialtyOffset(value int32) MacsecBasicKeyGenerationStatic
	// HasConfidentialtyOffset checks if ConfidentialtyOffset has been set in MacsecBasicKeyGenerationStatic
	HasConfidentialtyOffset() bool
}

type MacsecBasicKeyGenerationStaticCipherSuiteEnum string

// Enum of CipherSuite on MacsecBasicKeyGenerationStatic
var MacsecBasicKeyGenerationStaticCipherSuite = struct {
	GCM_AES_128     MacsecBasicKeyGenerationStaticCipherSuiteEnum
	GCM_AES_256     MacsecBasicKeyGenerationStaticCipherSuiteEnum
	GCM_AES_XPN_128 MacsecBasicKeyGenerationStaticCipherSuiteEnum
	GCM_AES_XPN_256 MacsecBasicKeyGenerationStaticCipherSuiteEnum
}{
	GCM_AES_128:     MacsecBasicKeyGenerationStaticCipherSuiteEnum("gcm_aes_128"),
	GCM_AES_256:     MacsecBasicKeyGenerationStaticCipherSuiteEnum("gcm_aes_256"),
	GCM_AES_XPN_128: MacsecBasicKeyGenerationStaticCipherSuiteEnum("gcm_aes_xpn_128"),
	GCM_AES_XPN_256: MacsecBasicKeyGenerationStaticCipherSuiteEnum("gcm_aes_xpn_256"),
}

func (obj *macsecBasicKeyGenerationStatic) CipherSuite() MacsecBasicKeyGenerationStaticCipherSuiteEnum {
	return MacsecBasicKeyGenerationStaticCipherSuiteEnum(obj.obj.CipherSuite.Enum().String())
}

// The cipher suite. Choose one from GCM-AES-128 GCM-AES-128 GCM-AES-256 GCM-AES-XPN-128 GCM-AES-XPN-256
// CipherSuite returns a string
func (obj *macsecBasicKeyGenerationStatic) HasCipherSuite() bool {
	return obj.obj.CipherSuite != nil
}

func (obj *macsecBasicKeyGenerationStatic) SetCipherSuite(value MacsecBasicKeyGenerationStaticCipherSuiteEnum) MacsecBasicKeyGenerationStatic {
	intValue, ok := otg.MacsecBasicKeyGenerationStatic_CipherSuite_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecBasicKeyGenerationStaticCipherSuiteEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecBasicKeyGenerationStatic_CipherSuite_Enum(intValue)
	obj.obj.CipherSuite = &enumValue

	return obj
}

// Confidentiality offset.
// ConfidentialtyOffset returns a int32
func (obj *macsecBasicKeyGenerationStatic) ConfidentialtyOffset() int32 {

	return *obj.obj.ConfidentialtyOffset

}

// Confidentiality offset.
// ConfidentialtyOffset returns a int32
func (obj *macsecBasicKeyGenerationStatic) HasConfidentialtyOffset() bool {
	return obj.obj.ConfidentialtyOffset != nil
}

// Confidentiality offset.
// SetConfidentialtyOffset sets the int32 value in the MacsecBasicKeyGenerationStatic object
func (obj *macsecBasicKeyGenerationStatic) SetConfidentialtyOffset(value int32) MacsecBasicKeyGenerationStatic {

	obj.obj.ConfidentialtyOffset = &value
	return obj
}

func (obj *macsecBasicKeyGenerationStatic) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *macsecBasicKeyGenerationStatic) setDefault() {
	if obj.obj.CipherSuite == nil {
		obj.SetCipherSuite(MacsecBasicKeyGenerationStaticCipherSuite.GCM_AES_128)

	}
	if obj.obj.ConfidentialtyOffset == nil {
		obj.SetConfidentialtyOffset(0)
	}

}
