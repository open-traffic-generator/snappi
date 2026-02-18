package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SecureEntityCryptoEngineEncryptOnlyFixedPn *****
type secureEntityCryptoEngineEncryptOnlyFixedPn struct {
	validation
	obj          *otg.SecureEntityCryptoEngineEncryptOnlyFixedPn
	marshaller   marshalSecureEntityCryptoEngineEncryptOnlyFixedPn
	unMarshaller unMarshalSecureEntityCryptoEngineEncryptOnlyFixedPn
}

func NewSecureEntityCryptoEngineEncryptOnlyFixedPn() SecureEntityCryptoEngineEncryptOnlyFixedPn {
	obj := secureEntityCryptoEngineEncryptOnlyFixedPn{obj: &otg.SecureEntityCryptoEngineEncryptOnlyFixedPn{}}
	obj.setDefault()
	return &obj
}

func (obj *secureEntityCryptoEngineEncryptOnlyFixedPn) msg() *otg.SecureEntityCryptoEngineEncryptOnlyFixedPn {
	return obj.obj
}

func (obj *secureEntityCryptoEngineEncryptOnlyFixedPn) setMsg(msg *otg.SecureEntityCryptoEngineEncryptOnlyFixedPn) SecureEntityCryptoEngineEncryptOnlyFixedPn {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsecureEntityCryptoEngineEncryptOnlyFixedPn struct {
	obj *secureEntityCryptoEngineEncryptOnlyFixedPn
}

type marshalSecureEntityCryptoEngineEncryptOnlyFixedPn interface {
	// ToProto marshals SecureEntityCryptoEngineEncryptOnlyFixedPn to protobuf object *otg.SecureEntityCryptoEngineEncryptOnlyFixedPn
	ToProto() (*otg.SecureEntityCryptoEngineEncryptOnlyFixedPn, error)
	// ToPbText marshals SecureEntityCryptoEngineEncryptOnlyFixedPn to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SecureEntityCryptoEngineEncryptOnlyFixedPn to YAML text
	ToYaml() (string, error)
	// ToJson marshals SecureEntityCryptoEngineEncryptOnlyFixedPn to JSON text
	ToJson() (string, error)
}

type unMarshalsecureEntityCryptoEngineEncryptOnlyFixedPn struct {
	obj *secureEntityCryptoEngineEncryptOnlyFixedPn
}

type unMarshalSecureEntityCryptoEngineEncryptOnlyFixedPn interface {
	// FromProto unmarshals SecureEntityCryptoEngineEncryptOnlyFixedPn from protobuf object *otg.SecureEntityCryptoEngineEncryptOnlyFixedPn
	FromProto(msg *otg.SecureEntityCryptoEngineEncryptOnlyFixedPn) (SecureEntityCryptoEngineEncryptOnlyFixedPn, error)
	// FromPbText unmarshals SecureEntityCryptoEngineEncryptOnlyFixedPn from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SecureEntityCryptoEngineEncryptOnlyFixedPn from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SecureEntityCryptoEngineEncryptOnlyFixedPn from JSON text
	FromJson(value string) error
}

func (obj *secureEntityCryptoEngineEncryptOnlyFixedPn) Marshal() marshalSecureEntityCryptoEngineEncryptOnlyFixedPn {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsecureEntityCryptoEngineEncryptOnlyFixedPn{obj: obj}
	}
	return obj.marshaller
}

func (obj *secureEntityCryptoEngineEncryptOnlyFixedPn) Unmarshal() unMarshalSecureEntityCryptoEngineEncryptOnlyFixedPn {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsecureEntityCryptoEngineEncryptOnlyFixedPn{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsecureEntityCryptoEngineEncryptOnlyFixedPn) ToProto() (*otg.SecureEntityCryptoEngineEncryptOnlyFixedPn, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsecureEntityCryptoEngineEncryptOnlyFixedPn) FromProto(msg *otg.SecureEntityCryptoEngineEncryptOnlyFixedPn) (SecureEntityCryptoEngineEncryptOnlyFixedPn, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsecureEntityCryptoEngineEncryptOnlyFixedPn) ToPbText() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptOnlyFixedPn) FromPbText(value string) error {
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

func (m *marshalsecureEntityCryptoEngineEncryptOnlyFixedPn) ToYaml() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptOnlyFixedPn) FromYaml(value string) error {
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

func (m *marshalsecureEntityCryptoEngineEncryptOnlyFixedPn) ToJson() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptOnlyFixedPn) FromJson(value string) error {
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

func (obj *secureEntityCryptoEngineEncryptOnlyFixedPn) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *secureEntityCryptoEngineEncryptOnlyFixedPn) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *secureEntityCryptoEngineEncryptOnlyFixedPn) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *secureEntityCryptoEngineEncryptOnlyFixedPn) Clone() (SecureEntityCryptoEngineEncryptOnlyFixedPn, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSecureEntityCryptoEngineEncryptOnlyFixedPn()
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

// SecureEntityCryptoEngineEncryptOnlyFixedPn is fixed packet number(PN) configuration.
type SecureEntityCryptoEngineEncryptOnlyFixedPn interface {
	Validation
	// msg marshals SecureEntityCryptoEngineEncryptOnlyFixedPn to protobuf object *otg.SecureEntityCryptoEngineEncryptOnlyFixedPn
	// and doesn't set defaults
	msg() *otg.SecureEntityCryptoEngineEncryptOnlyFixedPn
	// setMsg unmarshals SecureEntityCryptoEngineEncryptOnlyFixedPn from protobuf object *otg.SecureEntityCryptoEngineEncryptOnlyFixedPn
	// and doesn't set defaults
	setMsg(*otg.SecureEntityCryptoEngineEncryptOnlyFixedPn) SecureEntityCryptoEngineEncryptOnlyFixedPn
	// provides marshal interface
	Marshal() marshalSecureEntityCryptoEngineEncryptOnlyFixedPn
	// provides unmarshal interface
	Unmarshal() unMarshalSecureEntityCryptoEngineEncryptOnlyFixedPn
	// validate validates SecureEntityCryptoEngineEncryptOnlyFixedPn
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SecureEntityCryptoEngineEncryptOnlyFixedPn, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Pn returns uint32, set in SecureEntityCryptoEngineEncryptOnlyFixedPn.
	Pn() uint32
	// SetPn assigns uint32 provided by user to SecureEntityCryptoEngineEncryptOnlyFixedPn
	SetPn(value uint32) SecureEntityCryptoEngineEncryptOnlyFixedPn
	// HasPn checks if Pn has been set in SecureEntityCryptoEngineEncryptOnlyFixedPn
	HasPn() bool
	// Xpn returns string, set in SecureEntityCryptoEngineEncryptOnlyFixedPn.
	Xpn() string
	// SetXpn assigns string provided by user to SecureEntityCryptoEngineEncryptOnlyFixedPn
	SetXpn(value string) SecureEntityCryptoEngineEncryptOnlyFixedPn
	// HasXpn checks if Xpn has been set in SecureEntityCryptoEngineEncryptOnlyFixedPn
	HasXpn() bool
}

// Fixed Tx packet number(PN). 4 bytes PN with which all packets will be sent out.
// Pn returns a uint32
func (obj *secureEntityCryptoEngineEncryptOnlyFixedPn) Pn() uint32 {

	return *obj.obj.Pn

}

// Fixed Tx packet number(PN). 4 bytes PN with which all packets will be sent out.
// Pn returns a uint32
func (obj *secureEntityCryptoEngineEncryptOnlyFixedPn) HasPn() bool {
	return obj.obj.Pn != nil
}

// Fixed Tx packet number(PN). 4 bytes PN with which all packets will be sent out.
// SetPn sets the uint32 value in the SecureEntityCryptoEngineEncryptOnlyFixedPn object
func (obj *secureEntityCryptoEngineEncryptOnlyFixedPn) SetPn(value uint32) SecureEntityCryptoEngineEncryptOnlyFixedPn {

	obj.obj.Pn = &value
	return obj
}

// Fixed Tx extended packet number(XPN). 8 bytes XPN with which all packets will be sent out.
// Xpn returns a string
func (obj *secureEntityCryptoEngineEncryptOnlyFixedPn) Xpn() string {

	return *obj.obj.Xpn

}

// Fixed Tx extended packet number(XPN). 8 bytes XPN with which all packets will be sent out.
// Xpn returns a string
func (obj *secureEntityCryptoEngineEncryptOnlyFixedPn) HasXpn() bool {
	return obj.obj.Xpn != nil
}

// Fixed Tx extended packet number(XPN). 8 bytes XPN with which all packets will be sent out.
// SetXpn sets the string value in the SecureEntityCryptoEngineEncryptOnlyFixedPn object
func (obj *secureEntityCryptoEngineEncryptOnlyFixedPn) SetXpn(value string) SecureEntityCryptoEngineEncryptOnlyFixedPn {

	obj.obj.Xpn = &value
	return obj
}

func (obj *secureEntityCryptoEngineEncryptOnlyFixedPn) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Pn != nil {

		if *obj.obj.Pn < 1 || *obj.obj.Pn > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= SecureEntityCryptoEngineEncryptOnlyFixedPn.Pn <= 4294967295 but Got %d", *obj.obj.Pn))
		}

	}

	if obj.obj.Xpn != nil {

		if len(*obj.obj.Xpn) < 1 || len(*obj.obj.Xpn) > 18 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"1 <= length of SecureEntityCryptoEngineEncryptOnlyFixedPn.Xpn <= 18 but Got %d",
					len(*obj.obj.Xpn)))
		}

	}

}

func (obj *secureEntityCryptoEngineEncryptOnlyFixedPn) setDefault() {
	if obj.obj.Pn == nil {
		obj.SetPn(6)
	}
	if obj.obj.Xpn == nil {
		obj.SetXpn("0x0000000000000006")
	}

}
