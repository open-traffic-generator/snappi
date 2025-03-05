package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SecureEntityCryptoEngineEncryptDecryptTxScTxPn *****
type secureEntityCryptoEngineEncryptDecryptTxScTxPn struct {
	validation
	obj          *otg.SecureEntityCryptoEngineEncryptDecryptTxScTxPn
	marshaller   marshalSecureEntityCryptoEngineEncryptDecryptTxScTxPn
	unMarshaller unMarshalSecureEntityCryptoEngineEncryptDecryptTxScTxPn
}

func NewSecureEntityCryptoEngineEncryptDecryptTxScTxPn() SecureEntityCryptoEngineEncryptDecryptTxScTxPn {
	obj := secureEntityCryptoEngineEncryptDecryptTxScTxPn{obj: &otg.SecureEntityCryptoEngineEncryptDecryptTxScTxPn{}}
	obj.setDefault()
	return &obj
}

func (obj *secureEntityCryptoEngineEncryptDecryptTxScTxPn) msg() *otg.SecureEntityCryptoEngineEncryptDecryptTxScTxPn {
	return obj.obj
}

func (obj *secureEntityCryptoEngineEncryptDecryptTxScTxPn) setMsg(msg *otg.SecureEntityCryptoEngineEncryptDecryptTxScTxPn) SecureEntityCryptoEngineEncryptDecryptTxScTxPn {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsecureEntityCryptoEngineEncryptDecryptTxScTxPn struct {
	obj *secureEntityCryptoEngineEncryptDecryptTxScTxPn
}

type marshalSecureEntityCryptoEngineEncryptDecryptTxScTxPn interface {
	// ToProto marshals SecureEntityCryptoEngineEncryptDecryptTxScTxPn to protobuf object *otg.SecureEntityCryptoEngineEncryptDecryptTxScTxPn
	ToProto() (*otg.SecureEntityCryptoEngineEncryptDecryptTxScTxPn, error)
	// ToPbText marshals SecureEntityCryptoEngineEncryptDecryptTxScTxPn to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SecureEntityCryptoEngineEncryptDecryptTxScTxPn to YAML text
	ToYaml() (string, error)
	// ToJson marshals SecureEntityCryptoEngineEncryptDecryptTxScTxPn to JSON text
	ToJson() (string, error)
}

type unMarshalsecureEntityCryptoEngineEncryptDecryptTxScTxPn struct {
	obj *secureEntityCryptoEngineEncryptDecryptTxScTxPn
}

type unMarshalSecureEntityCryptoEngineEncryptDecryptTxScTxPn interface {
	// FromProto unmarshals SecureEntityCryptoEngineEncryptDecryptTxScTxPn from protobuf object *otg.SecureEntityCryptoEngineEncryptDecryptTxScTxPn
	FromProto(msg *otg.SecureEntityCryptoEngineEncryptDecryptTxScTxPn) (SecureEntityCryptoEngineEncryptDecryptTxScTxPn, error)
	// FromPbText unmarshals SecureEntityCryptoEngineEncryptDecryptTxScTxPn from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SecureEntityCryptoEngineEncryptDecryptTxScTxPn from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SecureEntityCryptoEngineEncryptDecryptTxScTxPn from JSON text
	FromJson(value string) error
}

func (obj *secureEntityCryptoEngineEncryptDecryptTxScTxPn) Marshal() marshalSecureEntityCryptoEngineEncryptDecryptTxScTxPn {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsecureEntityCryptoEngineEncryptDecryptTxScTxPn{obj: obj}
	}
	return obj.marshaller
}

func (obj *secureEntityCryptoEngineEncryptDecryptTxScTxPn) Unmarshal() unMarshalSecureEntityCryptoEngineEncryptDecryptTxScTxPn {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsecureEntityCryptoEngineEncryptDecryptTxScTxPn{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsecureEntityCryptoEngineEncryptDecryptTxScTxPn) ToProto() (*otg.SecureEntityCryptoEngineEncryptDecryptTxScTxPn, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsecureEntityCryptoEngineEncryptDecryptTxScTxPn) FromProto(msg *otg.SecureEntityCryptoEngineEncryptDecryptTxScTxPn) (SecureEntityCryptoEngineEncryptDecryptTxScTxPn, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsecureEntityCryptoEngineEncryptDecryptTxScTxPn) ToPbText() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptDecryptTxScTxPn) FromPbText(value string) error {
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

func (m *marshalsecureEntityCryptoEngineEncryptDecryptTxScTxPn) ToYaml() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptDecryptTxScTxPn) FromYaml(value string) error {
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

func (m *marshalsecureEntityCryptoEngineEncryptDecryptTxScTxPn) ToJson() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptDecryptTxScTxPn) FromJson(value string) error {
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

func (obj *secureEntityCryptoEngineEncryptDecryptTxScTxPn) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *secureEntityCryptoEngineEncryptDecryptTxScTxPn) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *secureEntityCryptoEngineEncryptDecryptTxScTxPn) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *secureEntityCryptoEngineEncryptDecryptTxScTxPn) Clone() (SecureEntityCryptoEngineEncryptDecryptTxScTxPn, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSecureEntityCryptoEngineEncryptDecryptTxScTxPn()
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

// SecureEntityCryptoEngineEncryptDecryptTxScTxPn is starting packet number(PN) configuration.
type SecureEntityCryptoEngineEncryptDecryptTxScTxPn interface {
	Validation
	// msg marshals SecureEntityCryptoEngineEncryptDecryptTxScTxPn to protobuf object *otg.SecureEntityCryptoEngineEncryptDecryptTxScTxPn
	// and doesn't set defaults
	msg() *otg.SecureEntityCryptoEngineEncryptDecryptTxScTxPn
	// setMsg unmarshals SecureEntityCryptoEngineEncryptDecryptTxScTxPn from protobuf object *otg.SecureEntityCryptoEngineEncryptDecryptTxScTxPn
	// and doesn't set defaults
	setMsg(*otg.SecureEntityCryptoEngineEncryptDecryptTxScTxPn) SecureEntityCryptoEngineEncryptDecryptTxScTxPn
	// provides marshal interface
	Marshal() marshalSecureEntityCryptoEngineEncryptDecryptTxScTxPn
	// provides unmarshal interface
	Unmarshal() unMarshalSecureEntityCryptoEngineEncryptDecryptTxScTxPn
	// validate validates SecureEntityCryptoEngineEncryptDecryptTxScTxPn
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SecureEntityCryptoEngineEncryptDecryptTxScTxPn, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// StartingPn returns uint32, set in SecureEntityCryptoEngineEncryptDecryptTxScTxPn.
	StartingPn() uint32
	// SetStartingPn assigns uint32 provided by user to SecureEntityCryptoEngineEncryptDecryptTxScTxPn
	SetStartingPn(value uint32) SecureEntityCryptoEngineEncryptDecryptTxScTxPn
	// HasStartingPn checks if StartingPn has been set in SecureEntityCryptoEngineEncryptDecryptTxScTxPn
	HasStartingPn() bool
	// StartingXpn returns string, set in SecureEntityCryptoEngineEncryptDecryptTxScTxPn.
	StartingXpn() string
	// SetStartingXpn assigns string provided by user to SecureEntityCryptoEngineEncryptDecryptTxScTxPn
	SetStartingXpn(value string) SecureEntityCryptoEngineEncryptDecryptTxScTxPn
	// HasStartingXpn checks if StartingXpn has been set in SecureEntityCryptoEngineEncryptDecryptTxScTxPn
	HasStartingXpn() bool
}

// Initial Tx packet number(PN).
// StartingPn returns a uint32
func (obj *secureEntityCryptoEngineEncryptDecryptTxScTxPn) StartingPn() uint32 {

	return *obj.obj.StartingPn

}

// Initial Tx packet number(PN).
// StartingPn returns a uint32
func (obj *secureEntityCryptoEngineEncryptDecryptTxScTxPn) HasStartingPn() bool {
	return obj.obj.StartingPn != nil
}

// Initial Tx packet number(PN).
// SetStartingPn sets the uint32 value in the SecureEntityCryptoEngineEncryptDecryptTxScTxPn object
func (obj *secureEntityCryptoEngineEncryptDecryptTxScTxPn) SetStartingPn(value uint32) SecureEntityCryptoEngineEncryptDecryptTxScTxPn {

	obj.obj.StartingPn = &value
	return obj
}

// The starting extended packet number(XPN).
// StartingXpn returns a string
func (obj *secureEntityCryptoEngineEncryptDecryptTxScTxPn) StartingXpn() string {

	return *obj.obj.StartingXpn

}

// The starting extended packet number(XPN).
// StartingXpn returns a string
func (obj *secureEntityCryptoEngineEncryptDecryptTxScTxPn) HasStartingXpn() bool {
	return obj.obj.StartingXpn != nil
}

// The starting extended packet number(XPN).
// SetStartingXpn sets the string value in the SecureEntityCryptoEngineEncryptDecryptTxScTxPn object
func (obj *secureEntityCryptoEngineEncryptDecryptTxScTxPn) SetStartingXpn(value string) SecureEntityCryptoEngineEncryptDecryptTxScTxPn {

	obj.obj.StartingXpn = &value
	return obj
}

func (obj *secureEntityCryptoEngineEncryptDecryptTxScTxPn) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.StartingPn != nil {

		if *obj.obj.StartingPn < 1 || *obj.obj.StartingPn > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= SecureEntityCryptoEngineEncryptDecryptTxScTxPn.StartingPn <= 4294967295 but Got %d", *obj.obj.StartingPn))
		}

	}

	if obj.obj.StartingXpn != nil {

		if len(*obj.obj.StartingXpn) < 1 || len(*obj.obj.StartingXpn) > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"1 <= length of SecureEntityCryptoEngineEncryptDecryptTxScTxPn.StartingXpn <= 16 but Got %d",
					len(*obj.obj.StartingXpn)))
		}

	}

}

func (obj *secureEntityCryptoEngineEncryptDecryptTxScTxPn) setDefault() {
	if obj.obj.StartingPn == nil {
		obj.SetStartingPn(1)
	}
	if obj.obj.StartingXpn == nil {
		obj.SetStartingXpn("0x0000000000000001")
	}

}
