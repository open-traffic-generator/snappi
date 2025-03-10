package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SecureEntityStaticKeySak *****
type secureEntityStaticKeySak struct {
	validation
	obj          *otg.SecureEntityStaticKeySak
	marshaller   marshalSecureEntityStaticKeySak
	unMarshaller unMarshalSecureEntityStaticKeySak
}

func NewSecureEntityStaticKeySak() SecureEntityStaticKeySak {
	obj := secureEntityStaticKeySak{obj: &otg.SecureEntityStaticKeySak{}}
	obj.setDefault()
	return &obj
}

func (obj *secureEntityStaticKeySak) msg() *otg.SecureEntityStaticKeySak {
	return obj.obj
}

func (obj *secureEntityStaticKeySak) setMsg(msg *otg.SecureEntityStaticKeySak) SecureEntityStaticKeySak {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsecureEntityStaticKeySak struct {
	obj *secureEntityStaticKeySak
}

type marshalSecureEntityStaticKeySak interface {
	// ToProto marshals SecureEntityStaticKeySak to protobuf object *otg.SecureEntityStaticKeySak
	ToProto() (*otg.SecureEntityStaticKeySak, error)
	// ToPbText marshals SecureEntityStaticKeySak to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SecureEntityStaticKeySak to YAML text
	ToYaml() (string, error)
	// ToJson marshals SecureEntityStaticKeySak to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals SecureEntityStaticKeySak to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalsecureEntityStaticKeySak struct {
	obj *secureEntityStaticKeySak
}

type unMarshalSecureEntityStaticKeySak interface {
	// FromProto unmarshals SecureEntityStaticKeySak from protobuf object *otg.SecureEntityStaticKeySak
	FromProto(msg *otg.SecureEntityStaticKeySak) (SecureEntityStaticKeySak, error)
	// FromPbText unmarshals SecureEntityStaticKeySak from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SecureEntityStaticKeySak from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SecureEntityStaticKeySak from JSON text
	FromJson(value string) error
}

func (obj *secureEntityStaticKeySak) Marshal() marshalSecureEntityStaticKeySak {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsecureEntityStaticKeySak{obj: obj}
	}
	return obj.marshaller
}

func (obj *secureEntityStaticKeySak) Unmarshal() unMarshalSecureEntityStaticKeySak {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsecureEntityStaticKeySak{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsecureEntityStaticKeySak) ToProto() (*otg.SecureEntityStaticKeySak, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsecureEntityStaticKeySak) FromProto(msg *otg.SecureEntityStaticKeySak) (SecureEntityStaticKeySak, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsecureEntityStaticKeySak) ToPbText() (string, error) {
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

func (m *unMarshalsecureEntityStaticKeySak) FromPbText(value string) error {
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

func (m *marshalsecureEntityStaticKeySak) ToYaml() (string, error) {
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

func (m *unMarshalsecureEntityStaticKeySak) FromYaml(value string) error {
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

func (m *marshalsecureEntityStaticKeySak) ToJsonRaw() (string, error) {
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

func (m *marshalsecureEntityStaticKeySak) ToJson() (string, error) {
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

func (m *unMarshalsecureEntityStaticKeySak) FromJson(value string) error {
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

func (obj *secureEntityStaticKeySak) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *secureEntityStaticKeySak) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *secureEntityStaticKeySak) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *secureEntityStaticKeySak) Clone() (SecureEntityStaticKeySak, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSecureEntityStaticKeySak()
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

// SecureEntityStaticKeySak is the container for SAK.
type SecureEntityStaticKeySak interface {
	Validation
	// msg marshals SecureEntityStaticKeySak to protobuf object *otg.SecureEntityStaticKeySak
	// and doesn't set defaults
	msg() *otg.SecureEntityStaticKeySak
	// setMsg unmarshals SecureEntityStaticKeySak from protobuf object *otg.SecureEntityStaticKeySak
	// and doesn't set defaults
	setMsg(*otg.SecureEntityStaticKeySak) SecureEntityStaticKeySak
	// provides marshal interface
	Marshal() marshalSecureEntityStaticKeySak
	// provides unmarshal interface
	Unmarshal() unMarshalSecureEntityStaticKeySak
	// validate validates SecureEntityStaticKeySak
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SecureEntityStaticKeySak, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Sak returns string, set in SecureEntityStaticKeySak.
	Sak() string
	// SetSak assigns string provided by user to SecureEntityStaticKeySak
	SetSak(value string) SecureEntityStaticKeySak
	// HasSak checks if Sak has been set in SecureEntityStaticKeySak
	HasSak() bool
	// Ssci returns string, set in SecureEntityStaticKeySak.
	Ssci() string
	// SetSsci assigns string provided by user to SecureEntityStaticKeySak
	SetSsci(value string) SecureEntityStaticKeySak
	// HasSsci checks if Ssci has been set in SecureEntityStaticKeySak
	HasSsci() bool
	// Salt returns string, set in SecureEntityStaticKeySak.
	Salt() string
	// SetSalt assigns string provided by user to SecureEntityStaticKeySak
	SetSalt(value string) SecureEntityStaticKeySak
	// HasSalt checks if Salt has been set in SecureEntityStaticKeySak
	HasSalt() bool
}

// Secure association key(SAK) bits as hex string. Either 128 bits or 256 bits depending on the chosen cipher suite.
// Sak returns a string
func (obj *secureEntityStaticKeySak) Sak() string {

	return *obj.obj.Sak

}

// Secure association key(SAK) bits as hex string. Either 128 bits or 256 bits depending on the chosen cipher suite.
// Sak returns a string
func (obj *secureEntityStaticKeySak) HasSak() bool {
	return obj.obj.Sak != nil
}

// Secure association key(SAK) bits as hex string. Either 128 bits or 256 bits depending on the chosen cipher suite.
// SetSak sets the string value in the SecureEntityStaticKeySak object
func (obj *secureEntityStaticKeySak) SetSak(value string) SecureEntityStaticKeySak {

	obj.obj.Sak = &value
	return obj
}

// 4 bytes short SCI(SSCI) used in case of XPN cipher suites.
// Ssci returns a string
func (obj *secureEntityStaticKeySak) Ssci() string {

	return *obj.obj.Ssci

}

// 4 bytes short SCI(SSCI) used in case of XPN cipher suites.
// Ssci returns a string
func (obj *secureEntityStaticKeySak) HasSsci() bool {
	return obj.obj.Ssci != nil
}

// 4 bytes short SCI(SSCI) used in case of XPN cipher suites.
// SetSsci sets the string value in the SecureEntityStaticKeySak object
func (obj *secureEntityStaticKeySak) SetSsci(value string) SecureEntityStaticKeySak {

	obj.obj.Ssci = &value
	return obj
}

// 12 bytes salt used in case of XPN cipher suites.
// Salt returns a string
func (obj *secureEntityStaticKeySak) Salt() string {

	return *obj.obj.Salt

}

// 12 bytes salt used in case of XPN cipher suites.
// Salt returns a string
func (obj *secureEntityStaticKeySak) HasSalt() bool {
	return obj.obj.Salt != nil
}

// 12 bytes salt used in case of XPN cipher suites.
// SetSalt sets the string value in the SecureEntityStaticKeySak object
func (obj *secureEntityStaticKeySak) SetSalt(value string) SecureEntityStaticKeySak {

	obj.obj.Salt = &value
	return obj
}

func (obj *secureEntityStaticKeySak) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Sak != nil {

		if len(*obj.obj.Sak) < 1 || len(*obj.obj.Sak) > 64 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"1 <= length of SecureEntityStaticKeySak.Sak <= 64 but Got %d",
					len(*obj.obj.Sak)))
		}

	}

	if obj.obj.Ssci != nil {

		if len(*obj.obj.Ssci) < 1 || len(*obj.obj.Ssci) > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"1 <= length of SecureEntityStaticKeySak.Ssci <= 8 but Got %d",
					len(*obj.obj.Ssci)))
		}

	}

	if obj.obj.Salt != nil {

		if len(*obj.obj.Salt) < 1 || len(*obj.obj.Salt) > 24 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"1 <= length of SecureEntityStaticKeySak.Salt <= 24 but Got %d",
					len(*obj.obj.Salt)))
		}

	}

}

func (obj *secureEntityStaticKeySak) setDefault() {
	if obj.obj.Sak == nil {
		obj.SetSak("F123456789ABCDEF0123456789ABCDEF")
	}
	if obj.obj.Ssci == nil {
		obj.SetSsci("00000001")
	}
	if obj.obj.Salt == nil {
		obj.SetSalt("000000000000000000000001")
	}

}
