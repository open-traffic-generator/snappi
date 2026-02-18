package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SecureEntityCryptoEngineEncryptOnlyIncrementingPn *****
type secureEntityCryptoEngineEncryptOnlyIncrementingPn struct {
	validation
	obj          *otg.SecureEntityCryptoEngineEncryptOnlyIncrementingPn
	marshaller   marshalSecureEntityCryptoEngineEncryptOnlyIncrementingPn
	unMarshaller unMarshalSecureEntityCryptoEngineEncryptOnlyIncrementingPn
}

func NewSecureEntityCryptoEngineEncryptOnlyIncrementingPn() SecureEntityCryptoEngineEncryptOnlyIncrementingPn {
	obj := secureEntityCryptoEngineEncryptOnlyIncrementingPn{obj: &otg.SecureEntityCryptoEngineEncryptOnlyIncrementingPn{}}
	obj.setDefault()
	return &obj
}

func (obj *secureEntityCryptoEngineEncryptOnlyIncrementingPn) msg() *otg.SecureEntityCryptoEngineEncryptOnlyIncrementingPn {
	return obj.obj
}

func (obj *secureEntityCryptoEngineEncryptOnlyIncrementingPn) setMsg(msg *otg.SecureEntityCryptoEngineEncryptOnlyIncrementingPn) SecureEntityCryptoEngineEncryptOnlyIncrementingPn {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsecureEntityCryptoEngineEncryptOnlyIncrementingPn struct {
	obj *secureEntityCryptoEngineEncryptOnlyIncrementingPn
}

type marshalSecureEntityCryptoEngineEncryptOnlyIncrementingPn interface {
	// ToProto marshals SecureEntityCryptoEngineEncryptOnlyIncrementingPn to protobuf object *otg.SecureEntityCryptoEngineEncryptOnlyIncrementingPn
	ToProto() (*otg.SecureEntityCryptoEngineEncryptOnlyIncrementingPn, error)
	// ToPbText marshals SecureEntityCryptoEngineEncryptOnlyIncrementingPn to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SecureEntityCryptoEngineEncryptOnlyIncrementingPn to YAML text
	ToYaml() (string, error)
	// ToJson marshals SecureEntityCryptoEngineEncryptOnlyIncrementingPn to JSON text
	ToJson() (string, error)
}

type unMarshalsecureEntityCryptoEngineEncryptOnlyIncrementingPn struct {
	obj *secureEntityCryptoEngineEncryptOnlyIncrementingPn
}

type unMarshalSecureEntityCryptoEngineEncryptOnlyIncrementingPn interface {
	// FromProto unmarshals SecureEntityCryptoEngineEncryptOnlyIncrementingPn from protobuf object *otg.SecureEntityCryptoEngineEncryptOnlyIncrementingPn
	FromProto(msg *otg.SecureEntityCryptoEngineEncryptOnlyIncrementingPn) (SecureEntityCryptoEngineEncryptOnlyIncrementingPn, error)
	// FromPbText unmarshals SecureEntityCryptoEngineEncryptOnlyIncrementingPn from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SecureEntityCryptoEngineEncryptOnlyIncrementingPn from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SecureEntityCryptoEngineEncryptOnlyIncrementingPn from JSON text
	FromJson(value string) error
}

func (obj *secureEntityCryptoEngineEncryptOnlyIncrementingPn) Marshal() marshalSecureEntityCryptoEngineEncryptOnlyIncrementingPn {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsecureEntityCryptoEngineEncryptOnlyIncrementingPn{obj: obj}
	}
	return obj.marshaller
}

func (obj *secureEntityCryptoEngineEncryptOnlyIncrementingPn) Unmarshal() unMarshalSecureEntityCryptoEngineEncryptOnlyIncrementingPn {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsecureEntityCryptoEngineEncryptOnlyIncrementingPn{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsecureEntityCryptoEngineEncryptOnlyIncrementingPn) ToProto() (*otg.SecureEntityCryptoEngineEncryptOnlyIncrementingPn, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsecureEntityCryptoEngineEncryptOnlyIncrementingPn) FromProto(msg *otg.SecureEntityCryptoEngineEncryptOnlyIncrementingPn) (SecureEntityCryptoEngineEncryptOnlyIncrementingPn, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsecureEntityCryptoEngineEncryptOnlyIncrementingPn) ToPbText() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptOnlyIncrementingPn) FromPbText(value string) error {
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

func (m *marshalsecureEntityCryptoEngineEncryptOnlyIncrementingPn) ToYaml() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptOnlyIncrementingPn) FromYaml(value string) error {
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

func (m *marshalsecureEntityCryptoEngineEncryptOnlyIncrementingPn) ToJson() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptOnlyIncrementingPn) FromJson(value string) error {
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

func (obj *secureEntityCryptoEngineEncryptOnlyIncrementingPn) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *secureEntityCryptoEngineEncryptOnlyIncrementingPn) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *secureEntityCryptoEngineEncryptOnlyIncrementingPn) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *secureEntityCryptoEngineEncryptOnlyIncrementingPn) Clone() (SecureEntityCryptoEngineEncryptOnlyIncrementingPn, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSecureEntityCryptoEngineEncryptOnlyIncrementingPn()
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

// SecureEntityCryptoEngineEncryptOnlyIncrementingPn is incrementing packet number(PN) configuration.
type SecureEntityCryptoEngineEncryptOnlyIncrementingPn interface {
	Validation
	// msg marshals SecureEntityCryptoEngineEncryptOnlyIncrementingPn to protobuf object *otg.SecureEntityCryptoEngineEncryptOnlyIncrementingPn
	// and doesn't set defaults
	msg() *otg.SecureEntityCryptoEngineEncryptOnlyIncrementingPn
	// setMsg unmarshals SecureEntityCryptoEngineEncryptOnlyIncrementingPn from protobuf object *otg.SecureEntityCryptoEngineEncryptOnlyIncrementingPn
	// and doesn't set defaults
	setMsg(*otg.SecureEntityCryptoEngineEncryptOnlyIncrementingPn) SecureEntityCryptoEngineEncryptOnlyIncrementingPn
	// provides marshal interface
	Marshal() marshalSecureEntityCryptoEngineEncryptOnlyIncrementingPn
	// provides unmarshal interface
	Unmarshal() unMarshalSecureEntityCryptoEngineEncryptOnlyIncrementingPn
	// validate validates SecureEntityCryptoEngineEncryptOnlyIncrementingPn
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SecureEntityCryptoEngineEncryptOnlyIncrementingPn, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Count returns uint32, set in SecureEntityCryptoEngineEncryptOnlyIncrementingPn.
	Count() uint32
	// SetCount assigns uint32 provided by user to SecureEntityCryptoEngineEncryptOnlyIncrementingPn
	SetCount(value uint32) SecureEntityCryptoEngineEncryptOnlyIncrementingPn
	// HasCount checks if Count has been set in SecureEntityCryptoEngineEncryptOnlyIncrementingPn
	HasCount() bool
	// StartingPn returns uint32, set in SecureEntityCryptoEngineEncryptOnlyIncrementingPn.
	StartingPn() uint32
	// SetStartingPn assigns uint32 provided by user to SecureEntityCryptoEngineEncryptOnlyIncrementingPn
	SetStartingPn(value uint32) SecureEntityCryptoEngineEncryptOnlyIncrementingPn
	// HasStartingPn checks if StartingPn has been set in SecureEntityCryptoEngineEncryptOnlyIncrementingPn
	HasStartingPn() bool
	// StartingXpn returns string, set in SecureEntityCryptoEngineEncryptOnlyIncrementingPn.
	StartingXpn() string
	// SetStartingXpn assigns string provided by user to SecureEntityCryptoEngineEncryptOnlyIncrementingPn
	SetStartingXpn(value string) SecureEntityCryptoEngineEncryptOnlyIncrementingPn
	// HasStartingXpn checks if StartingXpn has been set in SecureEntityCryptoEngineEncryptOnlyIncrementingPn
	HasStartingXpn() bool
}

// Count of packet numbers in series.
// Count returns a uint32
func (obj *secureEntityCryptoEngineEncryptOnlyIncrementingPn) Count() uint32 {

	return *obj.obj.Count

}

// Count of packet numbers in series.
// Count returns a uint32
func (obj *secureEntityCryptoEngineEncryptOnlyIncrementingPn) HasCount() bool {
	return obj.obj.Count != nil
}

// Count of packet numbers in series.
// SetCount sets the uint32 value in the SecureEntityCryptoEngineEncryptOnlyIncrementingPn object
func (obj *secureEntityCryptoEngineEncryptOnlyIncrementingPn) SetCount(value uint32) SecureEntityCryptoEngineEncryptOnlyIncrementingPn {

	obj.obj.Count = &value
	return obj
}

// The starting packet number(PN).
// StartingPn returns a uint32
func (obj *secureEntityCryptoEngineEncryptOnlyIncrementingPn) StartingPn() uint32 {

	return *obj.obj.StartingPn

}

// The starting packet number(PN).
// StartingPn returns a uint32
func (obj *secureEntityCryptoEngineEncryptOnlyIncrementingPn) HasStartingPn() bool {
	return obj.obj.StartingPn != nil
}

// The starting packet number(PN).
// SetStartingPn sets the uint32 value in the SecureEntityCryptoEngineEncryptOnlyIncrementingPn object
func (obj *secureEntityCryptoEngineEncryptOnlyIncrementingPn) SetStartingPn(value uint32) SecureEntityCryptoEngineEncryptOnlyIncrementingPn {

	obj.obj.StartingPn = &value
	return obj
}

// The starting extended packet number(XPN).
// StartingXpn returns a string
func (obj *secureEntityCryptoEngineEncryptOnlyIncrementingPn) StartingXpn() string {

	return *obj.obj.StartingXpn

}

// The starting extended packet number(XPN).
// StartingXpn returns a string
func (obj *secureEntityCryptoEngineEncryptOnlyIncrementingPn) HasStartingXpn() bool {
	return obj.obj.StartingXpn != nil
}

// The starting extended packet number(XPN).
// SetStartingXpn sets the string value in the SecureEntityCryptoEngineEncryptOnlyIncrementingPn object
func (obj *secureEntityCryptoEngineEncryptOnlyIncrementingPn) SetStartingXpn(value string) SecureEntityCryptoEngineEncryptOnlyIncrementingPn {

	obj.obj.StartingXpn = &value
	return obj
}

func (obj *secureEntityCryptoEngineEncryptOnlyIncrementingPn) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Count != nil {

		if *obj.obj.Count < 2 || *obj.obj.Count > 1000000 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("2 <= SecureEntityCryptoEngineEncryptOnlyIncrementingPn.Count <= 1000000 but Got %d", *obj.obj.Count))
		}

	}

	if obj.obj.StartingPn != nil {

		if *obj.obj.StartingPn < 1 || *obj.obj.StartingPn > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= SecureEntityCryptoEngineEncryptOnlyIncrementingPn.StartingPn <= 4294967295 but Got %d", *obj.obj.StartingPn))
		}

	}

	if obj.obj.StartingXpn != nil {

		if len(*obj.obj.StartingXpn) < 1 || len(*obj.obj.StartingXpn) > 18 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"1 <= length of SecureEntityCryptoEngineEncryptOnlyIncrementingPn.StartingXpn <= 18 but Got %d",
					len(*obj.obj.StartingXpn)))
		}

	}

}

func (obj *secureEntityCryptoEngineEncryptOnlyIncrementingPn) setDefault() {
	if obj.obj.Count == nil {
		obj.SetCount(100)
	}
	if obj.obj.StartingPn == nil {
		obj.SetStartingPn(10000)
	}
	if obj.obj.StartingXpn == nil {
		obj.SetStartingXpn("0x0000000000000001")
	}

}
