package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn *****
type macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn struct {
	validation
	obj          *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	marshaller   marshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	unMarshaller unMarshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
}

func NewMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn() MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn {
	obj := macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn{obj: &otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) msg() *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn {
	return obj.obj
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) setMsg(msg *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn struct {
	obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
}

type marshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn interface {
	// ToProto marshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn to protobuf object *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	ToProto() (*otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn, error)
	// ToPbText marshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn struct {
	obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
}

type unMarshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn interface {
	// FromProto unmarshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn from protobuf object *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	FromProto(msg *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) (MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn, error)
	// FromPbText unmarshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn from JSON text
	FromJson(value string) error
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) Marshal() marshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) Unmarshal() unMarshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) ToProto() (*otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) FromProto(msg *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) (MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) ToPbText() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) FromPbText(value string) error {
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

func (m *marshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) ToYaml() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) FromYaml(value string) error {
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

func (m *marshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) ToJson() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) FromJson(value string) error {
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

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) Clone() (MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn()
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

// MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn is initial PN settings.
type MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn interface {
	Validation
	// msg marshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn to protobuf object *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	// and doesn't set defaults
	msg() *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	// setMsg unmarshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn from protobuf object *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	// and doesn't set defaults
	setMsg(*otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	// provides marshal interface
	Marshal() marshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	// validate validates MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Pn returns uint32, set in MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn.
	Pn() uint32
	// SetPn assigns uint32 provided by user to MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	SetPn(value uint32) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	// HasPn checks if Pn has been set in MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	HasPn() bool
}

// Initial Tx PN.
// Pn returns a uint32
func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) Pn() uint32 {

	return *obj.obj.Pn

}

// Initial Tx PN.
// Pn returns a uint32
func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) HasPn() bool {
	return obj.obj.Pn != nil
}

// Initial Tx PN.
// SetPn sets the uint32 value in the MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn object
func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) SetPn(value uint32) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn {

	obj.obj.Pn = &value
	return obj
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Pn != nil {

		if *obj.obj.Pn < 1 || *obj.obj.Pn > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn.Pn <= 4294967295 but Got %d", *obj.obj.Pn))
		}

	}

}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) setDefault() {
	if obj.obj.Pn == nil {
		obj.SetPn(1)
	}

}
