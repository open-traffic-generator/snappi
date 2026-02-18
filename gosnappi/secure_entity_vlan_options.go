package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SecureEntityVlanOptions *****
type secureEntityVlanOptions struct {
	validation
	obj          *otg.SecureEntityVlanOptions
	marshaller   marshalSecureEntityVlanOptions
	unMarshaller unMarshalSecureEntityVlanOptions
}

func NewSecureEntityVlanOptions() SecureEntityVlanOptions {
	obj := secureEntityVlanOptions{obj: &otg.SecureEntityVlanOptions{}}
	obj.setDefault()
	return &obj
}

func (obj *secureEntityVlanOptions) msg() *otg.SecureEntityVlanOptions {
	return obj.obj
}

func (obj *secureEntityVlanOptions) setMsg(msg *otg.SecureEntityVlanOptions) SecureEntityVlanOptions {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsecureEntityVlanOptions struct {
	obj *secureEntityVlanOptions
}

type marshalSecureEntityVlanOptions interface {
	// ToProto marshals SecureEntityVlanOptions to protobuf object *otg.SecureEntityVlanOptions
	ToProto() (*otg.SecureEntityVlanOptions, error)
	// ToPbText marshals SecureEntityVlanOptions to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SecureEntityVlanOptions to YAML text
	ToYaml() (string, error)
	// ToJson marshals SecureEntityVlanOptions to JSON text
	ToJson() (string, error)
}

type unMarshalsecureEntityVlanOptions struct {
	obj *secureEntityVlanOptions
}

type unMarshalSecureEntityVlanOptions interface {
	// FromProto unmarshals SecureEntityVlanOptions from protobuf object *otg.SecureEntityVlanOptions
	FromProto(msg *otg.SecureEntityVlanOptions) (SecureEntityVlanOptions, error)
	// FromPbText unmarshals SecureEntityVlanOptions from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SecureEntityVlanOptions from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SecureEntityVlanOptions from JSON text
	FromJson(value string) error
}

func (obj *secureEntityVlanOptions) Marshal() marshalSecureEntityVlanOptions {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsecureEntityVlanOptions{obj: obj}
	}
	return obj.marshaller
}

func (obj *secureEntityVlanOptions) Unmarshal() unMarshalSecureEntityVlanOptions {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsecureEntityVlanOptions{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsecureEntityVlanOptions) ToProto() (*otg.SecureEntityVlanOptions, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsecureEntityVlanOptions) FromProto(msg *otg.SecureEntityVlanOptions) (SecureEntityVlanOptions, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsecureEntityVlanOptions) ToPbText() (string, error) {
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

func (m *unMarshalsecureEntityVlanOptions) FromPbText(value string) error {
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

func (m *marshalsecureEntityVlanOptions) ToYaml() (string, error) {
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

func (m *unMarshalsecureEntityVlanOptions) FromYaml(value string) error {
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

func (m *marshalsecureEntityVlanOptions) ToJson() (string, error) {
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

func (m *unMarshalsecureEntityVlanOptions) FromJson(value string) error {
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

func (obj *secureEntityVlanOptions) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *secureEntityVlanOptions) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *secureEntityVlanOptions) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *secureEntityVlanOptions) Clone() (SecureEntityVlanOptions, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSecureEntityVlanOptions()
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

// SecureEntityVlanOptions is a container for VLAN options of SecY.
type SecureEntityVlanOptions interface {
	Validation
	// msg marshals SecureEntityVlanOptions to protobuf object *otg.SecureEntityVlanOptions
	// and doesn't set defaults
	msg() *otg.SecureEntityVlanOptions
	// setMsg unmarshals SecureEntityVlanOptions from protobuf object *otg.SecureEntityVlanOptions
	// and doesn't set defaults
	setMsg(*otg.SecureEntityVlanOptions) SecureEntityVlanOptions
	// provides marshal interface
	Marshal() marshalSecureEntityVlanOptions
	// provides unmarshal interface
	Unmarshal() unMarshalSecureEntityVlanOptions
	// validate validates SecureEntityVlanOptions
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SecureEntityVlanOptions, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EncryptInterfaceVlans returns bool, set in SecureEntityVlanOptions.
	EncryptInterfaceVlans() bool
	// SetEncryptInterfaceVlans assigns bool provided by user to SecureEntityVlanOptions
	SetEncryptInterfaceVlans(value bool) SecureEntityVlanOptions
	// HasEncryptInterfaceVlans checks if EncryptInterfaceVlans has been set in SecureEntityVlanOptions
	HasEncryptInterfaceVlans() bool
}

// Send interface VLANS as encrypted or not. If it is false, VLANs go in cleartext.
// EncryptInterfaceVlans returns a bool
func (obj *secureEntityVlanOptions) EncryptInterfaceVlans() bool {

	return *obj.obj.EncryptInterfaceVlans

}

// Send interface VLANS as encrypted or not. If it is false, VLANs go in cleartext.
// EncryptInterfaceVlans returns a bool
func (obj *secureEntityVlanOptions) HasEncryptInterfaceVlans() bool {
	return obj.obj.EncryptInterfaceVlans != nil
}

// Send interface VLANS as encrypted or not. If it is false, VLANs go in cleartext.
// SetEncryptInterfaceVlans sets the bool value in the SecureEntityVlanOptions object
func (obj *secureEntityVlanOptions) SetEncryptInterfaceVlans(value bool) SecureEntityVlanOptions {

	obj.obj.EncryptInterfaceVlans = &value
	return obj
}

func (obj *secureEntityVlanOptions) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *secureEntityVlanOptions) setDefault() {
	if obj.obj.EncryptInterfaceVlans == nil {
		obj.SetEncryptInterfaceVlans(true)
	}

}
