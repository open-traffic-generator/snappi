package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn *****
type macsecCryptoEngineStatefulEncryptionDecryptionInitialPn struct {
	validation
	obj          *otg.MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn
	marshaller   marshalMacsecCryptoEngineStatefulEncryptionDecryptionInitialPn
	unMarshaller unMarshalMacsecCryptoEngineStatefulEncryptionDecryptionInitialPn
}

func NewMacsecCryptoEngineStatefulEncryptionDecryptionInitialPn() MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn {
	obj := macsecCryptoEngineStatefulEncryptionDecryptionInitialPn{obj: &otg.MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionInitialPn) msg() *otg.MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn {
	return obj.obj
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionInitialPn) setMsg(msg *otg.MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn) MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecCryptoEngineStatefulEncryptionDecryptionInitialPn struct {
	obj *macsecCryptoEngineStatefulEncryptionDecryptionInitialPn
}

type marshalMacsecCryptoEngineStatefulEncryptionDecryptionInitialPn interface {
	// ToProto marshals MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn to protobuf object *otg.MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn
	ToProto() (*otg.MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn, error)
	// ToPbText marshals MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecCryptoEngineStatefulEncryptionDecryptionInitialPn struct {
	obj *macsecCryptoEngineStatefulEncryptionDecryptionInitialPn
}

type unMarshalMacsecCryptoEngineStatefulEncryptionDecryptionInitialPn interface {
	// FromProto unmarshals MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn from protobuf object *otg.MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn
	FromProto(msg *otg.MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn) (MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn, error)
	// FromPbText unmarshals MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn from JSON text
	FromJson(value string) error
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionInitialPn) Marshal() marshalMacsecCryptoEngineStatefulEncryptionDecryptionInitialPn {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecCryptoEngineStatefulEncryptionDecryptionInitialPn{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionInitialPn) Unmarshal() unMarshalMacsecCryptoEngineStatefulEncryptionDecryptionInitialPn {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecCryptoEngineStatefulEncryptionDecryptionInitialPn{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecCryptoEngineStatefulEncryptionDecryptionInitialPn) ToProto() (*otg.MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecCryptoEngineStatefulEncryptionDecryptionInitialPn) FromProto(msg *otg.MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn) (MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecCryptoEngineStatefulEncryptionDecryptionInitialPn) ToPbText() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineStatefulEncryptionDecryptionInitialPn) FromPbText(value string) error {
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

func (m *marshalmacsecCryptoEngineStatefulEncryptionDecryptionInitialPn) ToYaml() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineStatefulEncryptionDecryptionInitialPn) FromYaml(value string) error {
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

func (m *marshalmacsecCryptoEngineStatefulEncryptionDecryptionInitialPn) ToJson() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineStatefulEncryptionDecryptionInitialPn) FromJson(value string) error {
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

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionInitialPn) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionInitialPn) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionInitialPn) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionInitialPn) Clone() (MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecCryptoEngineStatefulEncryptionDecryptionInitialPn()
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

// MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn is initial PN settings.
type MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn interface {
	Validation
	// msg marshals MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn to protobuf object *otg.MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn
	// and doesn't set defaults
	msg() *otg.MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn
	// setMsg unmarshals MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn from protobuf object *otg.MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn
	// and doesn't set defaults
	setMsg(*otg.MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn) MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn
	// provides marshal interface
	Marshal() marshalMacsecCryptoEngineStatefulEncryptionDecryptionInitialPn
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecCryptoEngineStatefulEncryptionDecryptionInitialPn
	// validate validates MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Pn returns uint32, set in MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn.
	Pn() uint32
	// SetPn assigns uint32 provided by user to MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn
	SetPn(value uint32) MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn
	// HasPn checks if Pn has been set in MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn
	HasPn() bool
}

// Initial Tx PN.
// Pn returns a uint32
func (obj *macsecCryptoEngineStatefulEncryptionDecryptionInitialPn) Pn() uint32 {

	return *obj.obj.Pn

}

// Initial Tx PN.
// Pn returns a uint32
func (obj *macsecCryptoEngineStatefulEncryptionDecryptionInitialPn) HasPn() bool {
	return obj.obj.Pn != nil
}

// Initial Tx PN.
// SetPn sets the uint32 value in the MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn object
func (obj *macsecCryptoEngineStatefulEncryptionDecryptionInitialPn) SetPn(value uint32) MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn {

	obj.obj.Pn = &value
	return obj
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionInitialPn) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Pn != nil {

		if *obj.obj.Pn < 1 || *obj.obj.Pn > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn.Pn <= 4294967295 but Got %d", *obj.obj.Pn))
		}

	}

}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionInitialPn) setDefault() {
	if obj.obj.Pn == nil {
		obj.SetPn(1)
	}

}
