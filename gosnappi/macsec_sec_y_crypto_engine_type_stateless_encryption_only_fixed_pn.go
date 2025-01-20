package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn *****
type macsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn struct {
	validation
	obj          *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	marshaller   marshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	unMarshaller unMarshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn
}

func NewMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn() MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn {
	obj := macsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn{obj: &otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn) msg() *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn {
	return obj.obj
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn) setMsg(msg *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn) MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn struct {
	obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn
}

type marshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn interface {
	// ToProto marshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn to protobuf object *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	ToProto() (*otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn, error)
	// ToPbText marshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn struct {
	obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn
}

type unMarshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn interface {
	// FromProto unmarshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn from protobuf object *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	FromProto(msg *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn) (MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn, error)
	// FromPbText unmarshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn from JSON text
	FromJson(value string) error
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn) Marshal() marshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn) Unmarshal() unMarshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn) ToProto() (*otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn) FromProto(msg *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn) (MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn) ToPbText() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn) FromPbText(value string) error {
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

func (m *marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn) ToYaml() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn) FromYaml(value string) error {
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

func (m *marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn) ToJson() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn) FromJson(value string) error {
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

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn) Clone() (MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn()
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

// MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn is fixed PN settings.
type MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn interface {
	Validation
	// msg marshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn to protobuf object *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	// and doesn't set defaults
	msg() *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	// setMsg unmarshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn from protobuf object *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	// and doesn't set defaults
	setMsg(*otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn) MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	// provides marshal interface
	Marshal() marshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	// validate validates MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Pn returns uint32, set in MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn.
	Pn() uint32
	// SetPn assigns uint32 provided by user to MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	SetPn(value uint32) MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	// HasPn checks if Pn has been set in MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn
	HasPn() bool
}

// Fixed Tx PN.
// Pn returns a uint32
func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn) Pn() uint32 {

	return *obj.obj.Pn

}

// Fixed Tx PN.
// Pn returns a uint32
func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn) HasPn() bool {
	return obj.obj.Pn != nil
}

// Fixed Tx PN.
// SetPn sets the uint32 value in the MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn object
func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn) SetPn(value uint32) MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn {

	obj.obj.Pn = &value
	return obj
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Pn != nil {

		if *obj.obj.Pn < 1 || *obj.obj.Pn > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn.Pn <= 4294967295 but Got %d", *obj.obj.Pn))
		}

	}

}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyFixedPn) setDefault() {
	if obj.obj.Pn == nil {
		obj.SetPn(6)
	}

}
