package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn *****
type lagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn struct {
	validation
	obj          *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn
	marshaller   marshalLagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn
	unMarshaller unMarshalLagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn
}

func NewLagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn() LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn {
	obj := lagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn{obj: &otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn{}}
	obj.setDefault()
	return &obj
}

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) msg() *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn {
	return obj.obj
}

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) setMsg(msg *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn struct {
	obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn
}

type marshalLagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn interface {
	// ToProto marshals LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn to protobuf object *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn
	ToProto() (*otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn, error)
	// ToPbText marshals LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn to YAML text
	ToYaml() (string, error)
	// ToJson marshals LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn to JSON text
	ToJson() (string, error)
}

type unMarshallagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn struct {
	obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn
}

type unMarshalLagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn interface {
	// FromProto unmarshals LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn from protobuf object *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn
	FromProto(msg *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) (LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn, error)
	// FromPbText unmarshals LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn from JSON text
	FromJson(value string) error
}

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) Marshal() marshalLagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn {
	if obj.marshaller == nil {
		obj.marshaller = &marshallagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn{obj: obj}
	}
	return obj.marshaller
}

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) Unmarshal() unMarshalLagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) ToProto() (*otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) FromProto(msg *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) (LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) ToPbText() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) FromPbText(value string) error {
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

func (m *marshallagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) ToYaml() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) FromYaml(value string) error {
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

func (m *marshallagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) ToJson() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) FromJson(value string) error {
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

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) Clone() (LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn()
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

// LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn is tx packet number(PN) configuration.
type LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn interface {
	Validation
	// msg marshals LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn to protobuf object *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn
	// and doesn't set defaults
	msg() *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn
	// setMsg unmarshals LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn from protobuf object *otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn
	// and doesn't set defaults
	setMsg(*otg.LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn
	// provides marshal interface
	Marshal() marshalLagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn
	// provides unmarshal interface
	Unmarshal() unMarshalLagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn
	// validate validates LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// StartingPn returns uint32, set in LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn.
	StartingPn() uint32
	// SetStartingPn assigns uint32 provided by user to LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn
	SetStartingPn(value uint32) LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn
	// HasStartingPn checks if StartingPn has been set in LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn
	HasStartingPn() bool
	// StartingXpn returns string, set in LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn.
	StartingXpn() string
	// SetStartingXpn assigns string provided by user to LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn
	SetStartingXpn(value string) LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn
	// HasStartingXpn checks if StartingXpn has been set in LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn
	HasStartingXpn() bool
}

// The starting packet number(PN).
// StartingPn returns a uint32
func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) StartingPn() uint32 {

	return *obj.obj.StartingPn

}

// The starting packet number(PN).
// StartingPn returns a uint32
func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) HasStartingPn() bool {
	return obj.obj.StartingPn != nil
}

// The starting packet number(PN).
// SetStartingPn sets the uint32 value in the LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn object
func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) SetStartingPn(value uint32) LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn {

	obj.obj.StartingPn = &value
	return obj
}

// The starting extended packet number(XPN).
// StartingXpn returns a string
func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) StartingXpn() string {

	return *obj.obj.StartingXpn

}

// The starting extended packet number(XPN).
// StartingXpn returns a string
func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) HasStartingXpn() bool {
	return obj.obj.StartingXpn != nil
}

// The starting extended packet number(XPN).
// SetStartingXpn sets the string value in the LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn object
func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) SetStartingXpn(value string) LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn {

	obj.obj.StartingXpn = &value
	return obj
}

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.StartingPn != nil {

		if *obj.obj.StartingPn < 1 || *obj.obj.StartingPn > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn.StartingPn <= 4294967295 but Got %d", *obj.obj.StartingPn))
		}

	}

	if obj.obj.StartingXpn != nil {

		if len(*obj.obj.StartingXpn) < 1 || len(*obj.obj.StartingXpn) > 16 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"1 <= length of LagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn.StartingXpn <= 16 but Got %d",
					len(*obj.obj.StartingXpn)))
		}

	}

}

func (obj *lagPortMacsecSecureEntityCryptoEngineEncryptDecryptTxPn) setDefault() {
	if obj.obj.StartingPn == nil {
		obj.SetStartingPn(1)
	}
	if obj.obj.StartingXpn == nil {
		obj.SetStartingXpn("01")
	}

}
