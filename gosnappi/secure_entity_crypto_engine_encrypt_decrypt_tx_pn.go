package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SecureEntityCryptoEngineEncryptDecryptTxPn *****
type secureEntityCryptoEngineEncryptDecryptTxPn struct {
	validation
	obj          *otg.SecureEntityCryptoEngineEncryptDecryptTxPn
	marshaller   marshalSecureEntityCryptoEngineEncryptDecryptTxPn
	unMarshaller unMarshalSecureEntityCryptoEngineEncryptDecryptTxPn
}

func NewSecureEntityCryptoEngineEncryptDecryptTxPn() SecureEntityCryptoEngineEncryptDecryptTxPn {
	obj := secureEntityCryptoEngineEncryptDecryptTxPn{obj: &otg.SecureEntityCryptoEngineEncryptDecryptTxPn{}}
	obj.setDefault()
	return &obj
}

func (obj *secureEntityCryptoEngineEncryptDecryptTxPn) msg() *otg.SecureEntityCryptoEngineEncryptDecryptTxPn {
	return obj.obj
}

func (obj *secureEntityCryptoEngineEncryptDecryptTxPn) setMsg(msg *otg.SecureEntityCryptoEngineEncryptDecryptTxPn) SecureEntityCryptoEngineEncryptDecryptTxPn {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsecureEntityCryptoEngineEncryptDecryptTxPn struct {
	obj *secureEntityCryptoEngineEncryptDecryptTxPn
}

type marshalSecureEntityCryptoEngineEncryptDecryptTxPn interface {
	// ToProto marshals SecureEntityCryptoEngineEncryptDecryptTxPn to protobuf object *otg.SecureEntityCryptoEngineEncryptDecryptTxPn
	ToProto() (*otg.SecureEntityCryptoEngineEncryptDecryptTxPn, error)
	// ToPbText marshals SecureEntityCryptoEngineEncryptDecryptTxPn to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SecureEntityCryptoEngineEncryptDecryptTxPn to YAML text
	ToYaml() (string, error)
	// ToJson marshals SecureEntityCryptoEngineEncryptDecryptTxPn to JSON text
	ToJson() (string, error)
}

type unMarshalsecureEntityCryptoEngineEncryptDecryptTxPn struct {
	obj *secureEntityCryptoEngineEncryptDecryptTxPn
}

type unMarshalSecureEntityCryptoEngineEncryptDecryptTxPn interface {
	// FromProto unmarshals SecureEntityCryptoEngineEncryptDecryptTxPn from protobuf object *otg.SecureEntityCryptoEngineEncryptDecryptTxPn
	FromProto(msg *otg.SecureEntityCryptoEngineEncryptDecryptTxPn) (SecureEntityCryptoEngineEncryptDecryptTxPn, error)
	// FromPbText unmarshals SecureEntityCryptoEngineEncryptDecryptTxPn from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SecureEntityCryptoEngineEncryptDecryptTxPn from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SecureEntityCryptoEngineEncryptDecryptTxPn from JSON text
	FromJson(value string) error
}

func (obj *secureEntityCryptoEngineEncryptDecryptTxPn) Marshal() marshalSecureEntityCryptoEngineEncryptDecryptTxPn {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsecureEntityCryptoEngineEncryptDecryptTxPn{obj: obj}
	}
	return obj.marshaller
}

func (obj *secureEntityCryptoEngineEncryptDecryptTxPn) Unmarshal() unMarshalSecureEntityCryptoEngineEncryptDecryptTxPn {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsecureEntityCryptoEngineEncryptDecryptTxPn{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsecureEntityCryptoEngineEncryptDecryptTxPn) ToProto() (*otg.SecureEntityCryptoEngineEncryptDecryptTxPn, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsecureEntityCryptoEngineEncryptDecryptTxPn) FromProto(msg *otg.SecureEntityCryptoEngineEncryptDecryptTxPn) (SecureEntityCryptoEngineEncryptDecryptTxPn, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsecureEntityCryptoEngineEncryptDecryptTxPn) ToPbText() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptDecryptTxPn) FromPbText(value string) error {
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

func (m *marshalsecureEntityCryptoEngineEncryptDecryptTxPn) ToYaml() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptDecryptTxPn) FromYaml(value string) error {
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

func (m *marshalsecureEntityCryptoEngineEncryptDecryptTxPn) ToJson() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptDecryptTxPn) FromJson(value string) error {
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

func (obj *secureEntityCryptoEngineEncryptDecryptTxPn) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *secureEntityCryptoEngineEncryptDecryptTxPn) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *secureEntityCryptoEngineEncryptDecryptTxPn) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *secureEntityCryptoEngineEncryptDecryptTxPn) Clone() (SecureEntityCryptoEngineEncryptDecryptTxPn, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSecureEntityCryptoEngineEncryptDecryptTxPn()
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

// SecureEntityCryptoEngineEncryptDecryptTxPn is tx packet number(PN) configuration.
type SecureEntityCryptoEngineEncryptDecryptTxPn interface {
	Validation
	// msg marshals SecureEntityCryptoEngineEncryptDecryptTxPn to protobuf object *otg.SecureEntityCryptoEngineEncryptDecryptTxPn
	// and doesn't set defaults
	msg() *otg.SecureEntityCryptoEngineEncryptDecryptTxPn
	// setMsg unmarshals SecureEntityCryptoEngineEncryptDecryptTxPn from protobuf object *otg.SecureEntityCryptoEngineEncryptDecryptTxPn
	// and doesn't set defaults
	setMsg(*otg.SecureEntityCryptoEngineEncryptDecryptTxPn) SecureEntityCryptoEngineEncryptDecryptTxPn
	// provides marshal interface
	Marshal() marshalSecureEntityCryptoEngineEncryptDecryptTxPn
	// provides unmarshal interface
	Unmarshal() unMarshalSecureEntityCryptoEngineEncryptDecryptTxPn
	// validate validates SecureEntityCryptoEngineEncryptDecryptTxPn
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SecureEntityCryptoEngineEncryptDecryptTxPn, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// StartingPn returns uint32, set in SecureEntityCryptoEngineEncryptDecryptTxPn.
	StartingPn() uint32
	// SetStartingPn assigns uint32 provided by user to SecureEntityCryptoEngineEncryptDecryptTxPn
	SetStartingPn(value uint32) SecureEntityCryptoEngineEncryptDecryptTxPn
	// HasStartingPn checks if StartingPn has been set in SecureEntityCryptoEngineEncryptDecryptTxPn
	HasStartingPn() bool
	// StartingXpn returns string, set in SecureEntityCryptoEngineEncryptDecryptTxPn.
	StartingXpn() string
	// SetStartingXpn assigns string provided by user to SecureEntityCryptoEngineEncryptDecryptTxPn
	SetStartingXpn(value string) SecureEntityCryptoEngineEncryptDecryptTxPn
	// HasStartingXpn checks if StartingXpn has been set in SecureEntityCryptoEngineEncryptDecryptTxPn
	HasStartingXpn() bool
}

// The starting packet number(PN).
// StartingPn returns a uint32
func (obj *secureEntityCryptoEngineEncryptDecryptTxPn) StartingPn() uint32 {

	return *obj.obj.StartingPn

}

// The starting packet number(PN).
// StartingPn returns a uint32
func (obj *secureEntityCryptoEngineEncryptDecryptTxPn) HasStartingPn() bool {
	return obj.obj.StartingPn != nil
}

// The starting packet number(PN).
// SetStartingPn sets the uint32 value in the SecureEntityCryptoEngineEncryptDecryptTxPn object
func (obj *secureEntityCryptoEngineEncryptDecryptTxPn) SetStartingPn(value uint32) SecureEntityCryptoEngineEncryptDecryptTxPn {

	obj.obj.StartingPn = &value
	return obj
}

// The starting extended packet number(XPN).
// StartingXpn returns a string
func (obj *secureEntityCryptoEngineEncryptDecryptTxPn) StartingXpn() string {

	return *obj.obj.StartingXpn

}

// The starting extended packet number(XPN).
// StartingXpn returns a string
func (obj *secureEntityCryptoEngineEncryptDecryptTxPn) HasStartingXpn() bool {
	return obj.obj.StartingXpn != nil
}

// The starting extended packet number(XPN).
// SetStartingXpn sets the string value in the SecureEntityCryptoEngineEncryptDecryptTxPn object
func (obj *secureEntityCryptoEngineEncryptDecryptTxPn) SetStartingXpn(value string) SecureEntityCryptoEngineEncryptDecryptTxPn {

	obj.obj.StartingXpn = &value
	return obj
}

func (obj *secureEntityCryptoEngineEncryptDecryptTxPn) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.StartingPn != nil {

		if *obj.obj.StartingPn < 1 || *obj.obj.StartingPn > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= SecureEntityCryptoEngineEncryptDecryptTxPn.StartingPn <= 4294967295 but Got %d", *obj.obj.StartingPn))
		}

	}

	if obj.obj.StartingXpn != nil {

		if len(*obj.obj.StartingXpn) < 1 || len(*obj.obj.StartingXpn) > 18 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"1 <= length of SecureEntityCryptoEngineEncryptDecryptTxPn.StartingXpn <= 18 but Got %d",
					len(*obj.obj.StartingXpn)))
		}

	}

}

func (obj *secureEntityCryptoEngineEncryptDecryptTxPn) setDefault() {
	if obj.obj.StartingPn == nil {
		obj.SetStartingPn(1)
	}
	if obj.obj.StartingXpn == nil {
		obj.SetStartingXpn("0x0000000000000001")
	}

}
