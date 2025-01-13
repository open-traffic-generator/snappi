package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn *****
type macsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn struct {
	validation
	obj          *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	marshaller   marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	unMarshaller unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn
}

func NewMacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn() MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn {
	obj := macsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn{obj: &otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn) msg() *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn {
	return obj.obj
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn) setMsg(msg *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn) MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn struct {
	obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn
}

type marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn interface {
	// ToProto marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn to protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	ToProto() (*otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn, error)
	// ToPbText marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn struct {
	obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn
}

type unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn interface {
	// FromProto unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn from protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	FromProto(msg *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn) (MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn, error)
	// FromPbText unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn from JSON text
	FromJson(value string) error
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn) Marshal() marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn) Unmarshal() unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn) ToProto() (*otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn) FromProto(msg *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn) (MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn) ToPbText() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn) FromPbText(value string) error {
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

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn) ToYaml() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn) FromYaml(value string) error {
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

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn) ToJson() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn) FromJson(value string) error {
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

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn) Clone() (MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn()
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

// MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn is fixed PN settings.
type MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn interface {
	Validation
	// msg marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn to protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	// and doesn't set defaults
	msg() *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	// setMsg unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn from protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	// and doesn't set defaults
	setMsg(*otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn) MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	// provides marshal interface
	Marshal() marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	// validate validates MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Pn returns int32, set in MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn.
	Pn() int32
	// SetPn assigns int32 provided by user to MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	SetPn(value int32) MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	// HasPn checks if Pn has been set in MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	HasPn() bool
}

// Fixed Tx PN.
// Pn returns a int32
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn) Pn() int32 {

	return *obj.obj.Pn

}

// Fixed Tx PN.
// Pn returns a int32
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn) HasPn() bool {
	return obj.obj.Pn != nil
}

// Fixed Tx PN.
// SetPn sets the int32 value in the MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn object
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn) SetPn(value int32) MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn {

	obj.obj.Pn = &value
	return obj
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedPn) setDefault() {
	if obj.obj.Pn == nil {
		obj.SetPn(6)
	}

}
