package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecCryptoEngineStatelessEncryptionOnlyFixedPn *****
type macsecCryptoEngineStatelessEncryptionOnlyFixedPn struct {
	validation
	obj          *otg.MacsecCryptoEngineStatelessEncryptionOnlyFixedPn
	marshaller   marshalMacsecCryptoEngineStatelessEncryptionOnlyFixedPn
	unMarshaller unMarshalMacsecCryptoEngineStatelessEncryptionOnlyFixedPn
}

func NewMacsecCryptoEngineStatelessEncryptionOnlyFixedPn() MacsecCryptoEngineStatelessEncryptionOnlyFixedPn {
	obj := macsecCryptoEngineStatelessEncryptionOnlyFixedPn{obj: &otg.MacsecCryptoEngineStatelessEncryptionOnlyFixedPn{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyFixedPn) msg() *otg.MacsecCryptoEngineStatelessEncryptionOnlyFixedPn {
	return obj.obj
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyFixedPn) setMsg(msg *otg.MacsecCryptoEngineStatelessEncryptionOnlyFixedPn) MacsecCryptoEngineStatelessEncryptionOnlyFixedPn {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecCryptoEngineStatelessEncryptionOnlyFixedPn struct {
	obj *macsecCryptoEngineStatelessEncryptionOnlyFixedPn
}

type marshalMacsecCryptoEngineStatelessEncryptionOnlyFixedPn interface {
	// ToProto marshals MacsecCryptoEngineStatelessEncryptionOnlyFixedPn to protobuf object *otg.MacsecCryptoEngineStatelessEncryptionOnlyFixedPn
	ToProto() (*otg.MacsecCryptoEngineStatelessEncryptionOnlyFixedPn, error)
	// ToPbText marshals MacsecCryptoEngineStatelessEncryptionOnlyFixedPn to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecCryptoEngineStatelessEncryptionOnlyFixedPn to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecCryptoEngineStatelessEncryptionOnlyFixedPn to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecCryptoEngineStatelessEncryptionOnlyFixedPn struct {
	obj *macsecCryptoEngineStatelessEncryptionOnlyFixedPn
}

type unMarshalMacsecCryptoEngineStatelessEncryptionOnlyFixedPn interface {
	// FromProto unmarshals MacsecCryptoEngineStatelessEncryptionOnlyFixedPn from protobuf object *otg.MacsecCryptoEngineStatelessEncryptionOnlyFixedPn
	FromProto(msg *otg.MacsecCryptoEngineStatelessEncryptionOnlyFixedPn) (MacsecCryptoEngineStatelessEncryptionOnlyFixedPn, error)
	// FromPbText unmarshals MacsecCryptoEngineStatelessEncryptionOnlyFixedPn from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecCryptoEngineStatelessEncryptionOnlyFixedPn from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecCryptoEngineStatelessEncryptionOnlyFixedPn from JSON text
	FromJson(value string) error
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyFixedPn) Marshal() marshalMacsecCryptoEngineStatelessEncryptionOnlyFixedPn {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecCryptoEngineStatelessEncryptionOnlyFixedPn{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyFixedPn) Unmarshal() unMarshalMacsecCryptoEngineStatelessEncryptionOnlyFixedPn {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecCryptoEngineStatelessEncryptionOnlyFixedPn{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecCryptoEngineStatelessEncryptionOnlyFixedPn) ToProto() (*otg.MacsecCryptoEngineStatelessEncryptionOnlyFixedPn, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecCryptoEngineStatelessEncryptionOnlyFixedPn) FromProto(msg *otg.MacsecCryptoEngineStatelessEncryptionOnlyFixedPn) (MacsecCryptoEngineStatelessEncryptionOnlyFixedPn, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecCryptoEngineStatelessEncryptionOnlyFixedPn) ToPbText() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineStatelessEncryptionOnlyFixedPn) FromPbText(value string) error {
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

func (m *marshalmacsecCryptoEngineStatelessEncryptionOnlyFixedPn) ToYaml() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineStatelessEncryptionOnlyFixedPn) FromYaml(value string) error {
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

func (m *marshalmacsecCryptoEngineStatelessEncryptionOnlyFixedPn) ToJson() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineStatelessEncryptionOnlyFixedPn) FromJson(value string) error {
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

func (obj *macsecCryptoEngineStatelessEncryptionOnlyFixedPn) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyFixedPn) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyFixedPn) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyFixedPn) Clone() (MacsecCryptoEngineStatelessEncryptionOnlyFixedPn, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecCryptoEngineStatelessEncryptionOnlyFixedPn()
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

// MacsecCryptoEngineStatelessEncryptionOnlyFixedPn is fixed packet number(PN) configuration.
type MacsecCryptoEngineStatelessEncryptionOnlyFixedPn interface {
	Validation
	// msg marshals MacsecCryptoEngineStatelessEncryptionOnlyFixedPn to protobuf object *otg.MacsecCryptoEngineStatelessEncryptionOnlyFixedPn
	// and doesn't set defaults
	msg() *otg.MacsecCryptoEngineStatelessEncryptionOnlyFixedPn
	// setMsg unmarshals MacsecCryptoEngineStatelessEncryptionOnlyFixedPn from protobuf object *otg.MacsecCryptoEngineStatelessEncryptionOnlyFixedPn
	// and doesn't set defaults
	setMsg(*otg.MacsecCryptoEngineStatelessEncryptionOnlyFixedPn) MacsecCryptoEngineStatelessEncryptionOnlyFixedPn
	// provides marshal interface
	Marshal() marshalMacsecCryptoEngineStatelessEncryptionOnlyFixedPn
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecCryptoEngineStatelessEncryptionOnlyFixedPn
	// validate validates MacsecCryptoEngineStatelessEncryptionOnlyFixedPn
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecCryptoEngineStatelessEncryptionOnlyFixedPn, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Pn returns uint32, set in MacsecCryptoEngineStatelessEncryptionOnlyFixedPn.
	Pn() uint32
	// SetPn assigns uint32 provided by user to MacsecCryptoEngineStatelessEncryptionOnlyFixedPn
	SetPn(value uint32) MacsecCryptoEngineStatelessEncryptionOnlyFixedPn
	// HasPn checks if Pn has been set in MacsecCryptoEngineStatelessEncryptionOnlyFixedPn
	HasPn() bool
	// Xpn returns string, set in MacsecCryptoEngineStatelessEncryptionOnlyFixedPn.
	Xpn() string
	// SetXpn assigns string provided by user to MacsecCryptoEngineStatelessEncryptionOnlyFixedPn
	SetXpn(value string) MacsecCryptoEngineStatelessEncryptionOnlyFixedPn
	// HasXpn checks if Xpn has been set in MacsecCryptoEngineStatelessEncryptionOnlyFixedPn
	HasXpn() bool
}

// Fixed Tx packet number(PN). 4 bytes PN with which all packets will be sent out.
// Pn returns a uint32
func (obj *macsecCryptoEngineStatelessEncryptionOnlyFixedPn) Pn() uint32 {

	return *obj.obj.Pn

}

// Fixed Tx packet number(PN). 4 bytes PN with which all packets will be sent out.
// Pn returns a uint32
func (obj *macsecCryptoEngineStatelessEncryptionOnlyFixedPn) HasPn() bool {
	return obj.obj.Pn != nil
}

// Fixed Tx packet number(PN). 4 bytes PN with which all packets will be sent out.
// SetPn sets the uint32 value in the MacsecCryptoEngineStatelessEncryptionOnlyFixedPn object
func (obj *macsecCryptoEngineStatelessEncryptionOnlyFixedPn) SetPn(value uint32) MacsecCryptoEngineStatelessEncryptionOnlyFixedPn {

	obj.obj.Pn = &value
	return obj
}

// Fixed Tx extended packet number(XPN). 8 bytes XPN with which all packets will be sent out.
// Xpn returns a string
func (obj *macsecCryptoEngineStatelessEncryptionOnlyFixedPn) Xpn() string {

	return *obj.obj.Xpn

}

// Fixed Tx extended packet number(XPN). 8 bytes XPN with which all packets will be sent out.
// Xpn returns a string
func (obj *macsecCryptoEngineStatelessEncryptionOnlyFixedPn) HasXpn() bool {
	return obj.obj.Xpn != nil
}

// Fixed Tx extended packet number(XPN). 8 bytes XPN with which all packets will be sent out.
// SetXpn sets the string value in the MacsecCryptoEngineStatelessEncryptionOnlyFixedPn object
func (obj *macsecCryptoEngineStatelessEncryptionOnlyFixedPn) SetXpn(value string) MacsecCryptoEngineStatelessEncryptionOnlyFixedPn {

	obj.obj.Xpn = &value
	return obj
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyFixedPn) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Pn != nil {

		if *obj.obj.Pn < 4294967295 || *obj.obj.Pn > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("4294967295 <= MacsecCryptoEngineStatelessEncryptionOnlyFixedPn.Pn <= 4294967295 but Got %d", *obj.obj.Pn))
		}

	}

	if obj.obj.Xpn != nil {

		if len(*obj.obj.Xpn) < 1 || len(*obj.obj.Xpn) > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"1 <= length of MacsecCryptoEngineStatelessEncryptionOnlyFixedPn.Xpn <= 8 but Got %d",
					len(*obj.obj.Xpn)))
		}

	}

}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyFixedPn) setDefault() {
	if obj.obj.Pn == nil {
		obj.SetPn(6)
	}
	if obj.obj.Xpn == nil {
		obj.SetXpn("0x0000000000000006")
	}

}
