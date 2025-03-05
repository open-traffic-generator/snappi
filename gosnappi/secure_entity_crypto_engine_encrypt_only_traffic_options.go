package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SecureEntityCryptoEngineEncryptOnlyTrafficOptions *****
type secureEntityCryptoEngineEncryptOnlyTrafficOptions struct {
	validation
	obj          *otg.SecureEntityCryptoEngineEncryptOnlyTrafficOptions
	marshaller   marshalSecureEntityCryptoEngineEncryptOnlyTrafficOptions
	unMarshaller unMarshalSecureEntityCryptoEngineEncryptOnlyTrafficOptions
}

func NewSecureEntityCryptoEngineEncryptOnlyTrafficOptions() SecureEntityCryptoEngineEncryptOnlyTrafficOptions {
	obj := secureEntityCryptoEngineEncryptOnlyTrafficOptions{obj: &otg.SecureEntityCryptoEngineEncryptOnlyTrafficOptions{}}
	obj.setDefault()
	return &obj
}

func (obj *secureEntityCryptoEngineEncryptOnlyTrafficOptions) msg() *otg.SecureEntityCryptoEngineEncryptOnlyTrafficOptions {
	return obj.obj
}

func (obj *secureEntityCryptoEngineEncryptOnlyTrafficOptions) setMsg(msg *otg.SecureEntityCryptoEngineEncryptOnlyTrafficOptions) SecureEntityCryptoEngineEncryptOnlyTrafficOptions {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsecureEntityCryptoEngineEncryptOnlyTrafficOptions struct {
	obj *secureEntityCryptoEngineEncryptOnlyTrafficOptions
}

type marshalSecureEntityCryptoEngineEncryptOnlyTrafficOptions interface {
	// ToProto marshals SecureEntityCryptoEngineEncryptOnlyTrafficOptions to protobuf object *otg.SecureEntityCryptoEngineEncryptOnlyTrafficOptions
	ToProto() (*otg.SecureEntityCryptoEngineEncryptOnlyTrafficOptions, error)
	// ToPbText marshals SecureEntityCryptoEngineEncryptOnlyTrafficOptions to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SecureEntityCryptoEngineEncryptOnlyTrafficOptions to YAML text
	ToYaml() (string, error)
	// ToJson marshals SecureEntityCryptoEngineEncryptOnlyTrafficOptions to JSON text
	ToJson() (string, error)
}

type unMarshalsecureEntityCryptoEngineEncryptOnlyTrafficOptions struct {
	obj *secureEntityCryptoEngineEncryptOnlyTrafficOptions
}

type unMarshalSecureEntityCryptoEngineEncryptOnlyTrafficOptions interface {
	// FromProto unmarshals SecureEntityCryptoEngineEncryptOnlyTrafficOptions from protobuf object *otg.SecureEntityCryptoEngineEncryptOnlyTrafficOptions
	FromProto(msg *otg.SecureEntityCryptoEngineEncryptOnlyTrafficOptions) (SecureEntityCryptoEngineEncryptOnlyTrafficOptions, error)
	// FromPbText unmarshals SecureEntityCryptoEngineEncryptOnlyTrafficOptions from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SecureEntityCryptoEngineEncryptOnlyTrafficOptions from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SecureEntityCryptoEngineEncryptOnlyTrafficOptions from JSON text
	FromJson(value string) error
}

func (obj *secureEntityCryptoEngineEncryptOnlyTrafficOptions) Marshal() marshalSecureEntityCryptoEngineEncryptOnlyTrafficOptions {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsecureEntityCryptoEngineEncryptOnlyTrafficOptions{obj: obj}
	}
	return obj.marshaller
}

func (obj *secureEntityCryptoEngineEncryptOnlyTrafficOptions) Unmarshal() unMarshalSecureEntityCryptoEngineEncryptOnlyTrafficOptions {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsecureEntityCryptoEngineEncryptOnlyTrafficOptions{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsecureEntityCryptoEngineEncryptOnlyTrafficOptions) ToProto() (*otg.SecureEntityCryptoEngineEncryptOnlyTrafficOptions, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsecureEntityCryptoEngineEncryptOnlyTrafficOptions) FromProto(msg *otg.SecureEntityCryptoEngineEncryptOnlyTrafficOptions) (SecureEntityCryptoEngineEncryptOnlyTrafficOptions, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsecureEntityCryptoEngineEncryptOnlyTrafficOptions) ToPbText() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptOnlyTrafficOptions) FromPbText(value string) error {
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

func (m *marshalsecureEntityCryptoEngineEncryptOnlyTrafficOptions) ToYaml() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptOnlyTrafficOptions) FromYaml(value string) error {
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

func (m *marshalsecureEntityCryptoEngineEncryptOnlyTrafficOptions) ToJson() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptOnlyTrafficOptions) FromJson(value string) error {
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

func (obj *secureEntityCryptoEngineEncryptOnlyTrafficOptions) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *secureEntityCryptoEngineEncryptOnlyTrafficOptions) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *secureEntityCryptoEngineEncryptOnlyTrafficOptions) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *secureEntityCryptoEngineEncryptOnlyTrafficOptions) Clone() (SecureEntityCryptoEngineEncryptOnlyTrafficOptions, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSecureEntityCryptoEngineEncryptOnlyTrafficOptions()
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

// SecureEntityCryptoEngineEncryptOnlyTrafficOptions is encrypt only traffic options.
type SecureEntityCryptoEngineEncryptOnlyTrafficOptions interface {
	Validation
	// msg marshals SecureEntityCryptoEngineEncryptOnlyTrafficOptions to protobuf object *otg.SecureEntityCryptoEngineEncryptOnlyTrafficOptions
	// and doesn't set defaults
	msg() *otg.SecureEntityCryptoEngineEncryptOnlyTrafficOptions
	// setMsg unmarshals SecureEntityCryptoEngineEncryptOnlyTrafficOptions from protobuf object *otg.SecureEntityCryptoEngineEncryptOnlyTrafficOptions
	// and doesn't set defaults
	setMsg(*otg.SecureEntityCryptoEngineEncryptOnlyTrafficOptions) SecureEntityCryptoEngineEncryptOnlyTrafficOptions
	// provides marshal interface
	Marshal() marshalSecureEntityCryptoEngineEncryptOnlyTrafficOptions
	// provides unmarshal interface
	Unmarshal() unMarshalSecureEntityCryptoEngineEncryptOnlyTrafficOptions
	// validate validates SecureEntityCryptoEngineEncryptOnlyTrafficOptions
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SecureEntityCryptoEngineEncryptOnlyTrafficOptions, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SendGratuitousArp returns bool, set in SecureEntityCryptoEngineEncryptOnlyTrafficOptions.
	SendGratuitousArp() bool
	// SetSendGratuitousArp assigns bool provided by user to SecureEntityCryptoEngineEncryptOnlyTrafficOptions
	SetSendGratuitousArp(value bool) SecureEntityCryptoEngineEncryptOnlyTrafficOptions
	// HasSendGratuitousArp checks if SendGratuitousArp has been set in SecureEntityCryptoEngineEncryptOnlyTrafficOptions
	HasSendGratuitousArp() bool
}

// Send gratuitous ARP or not.
// SendGratuitousArp returns a bool
func (obj *secureEntityCryptoEngineEncryptOnlyTrafficOptions) SendGratuitousArp() bool {

	return *obj.obj.SendGratuitousArp

}

// Send gratuitous ARP or not.
// SendGratuitousArp returns a bool
func (obj *secureEntityCryptoEngineEncryptOnlyTrafficOptions) HasSendGratuitousArp() bool {
	return obj.obj.SendGratuitousArp != nil
}

// Send gratuitous ARP or not.
// SetSendGratuitousArp sets the bool value in the SecureEntityCryptoEngineEncryptOnlyTrafficOptions object
func (obj *secureEntityCryptoEngineEncryptOnlyTrafficOptions) SetSendGratuitousArp(value bool) SecureEntityCryptoEngineEncryptOnlyTrafficOptions {

	obj.obj.SendGratuitousArp = &value
	return obj
}

func (obj *secureEntityCryptoEngineEncryptOnlyTrafficOptions) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *secureEntityCryptoEngineEncryptOnlyTrafficOptions) setDefault() {
	if obj.obj.SendGratuitousArp == nil {
		obj.SetSendGratuitousArp(true)
	}

}
