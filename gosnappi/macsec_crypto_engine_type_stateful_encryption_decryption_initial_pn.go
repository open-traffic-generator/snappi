package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn *****
type macsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn struct {
	validation
	obj          *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	marshaller   marshalMacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	unMarshaller unMarshalMacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
}

func NewMacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn() MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn {
	obj := macsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn{obj: &otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) msg() *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn {
	return obj.obj
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) setMsg(msg *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn struct {
	obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
}

type marshalMacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn interface {
	// ToProto marshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn to protobuf object *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	ToProto() (*otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn, error)
	// ToPbText marshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn struct {
	obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
}

type unMarshalMacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn interface {
	// FromProto unmarshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn from protobuf object *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	FromProto(msg *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) (MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn, error)
	// FromPbText unmarshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn from JSON text
	FromJson(value string) error
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) Marshal() marshalMacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) Unmarshal() unMarshalMacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) ToProto() (*otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) FromProto(msg *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) (MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) ToPbText() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) FromPbText(value string) error {
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

func (m *marshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) ToYaml() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) FromYaml(value string) error {
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

func (m *marshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) ToJson() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) FromJson(value string) error {
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

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) Clone() (MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn()
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

// MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn is initial PN settings.
type MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn interface {
	Validation
	// msg marshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn to protobuf object *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	// and doesn't set defaults
	msg() *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	// setMsg unmarshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn from protobuf object *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	// and doesn't set defaults
	setMsg(*otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	// provides marshal interface
	Marshal() marshalMacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	// validate validates MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Pn returns int32, set in MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn.
	Pn() int32
	// SetPn assigns int32 provided by user to MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	SetPn(value int32) MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	// HasPn checks if Pn has been set in MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	HasPn() bool
}

// Initial Tx PN.
// Pn returns a int32
func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) Pn() int32 {

	return *obj.obj.Pn

}

// Initial Tx PN.
// Pn returns a int32
func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) HasPn() bool {
	return obj.obj.Pn != nil
}

// Initial Tx PN.
// SetPn sets the int32 value in the MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn object
func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) SetPn(value int32) MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn {

	obj.obj.Pn = &value
	return obj
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) setDefault() {
	if obj.obj.Pn == nil {
		obj.SetPn(1)
	}

}
