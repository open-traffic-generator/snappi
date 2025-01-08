package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn *****
type macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn struct {
	validation
	obj          *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	marshaller   marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	unMarshaller unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
}

func NewMacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn() MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn {
	obj := macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn{obj: &otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) msg() *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn {
	return obj.obj
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) setMsg(msg *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn struct {
	obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
}

type marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn interface {
	// ToProto marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn to protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	ToProto() (*otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn, error)
	// ToPbText marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn struct {
	obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
}

type unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn interface {
	// FromProto unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn from protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	FromProto(msg *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) (MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn, error)
	// FromPbText unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn from JSON text
	FromJson(value string) error
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) Marshal() marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) Unmarshal() unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) ToProto() (*otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) FromProto(msg *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) (MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) ToPbText() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) FromPbText(value string) error {
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

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) ToYaml() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) FromYaml(value string) error {
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

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) ToJson() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) FromJson(value string) error {
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

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) Clone() (MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn()
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

// MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn is tx PN settings.
// choice:
// description: >-
// Types of Tx PN series.
// type: string
// default: fixed_pn
// x-field-uid: 1
// x-enum:
// fixed_pn:
// x-field-uid: 1
// incrementing_pn:
// x-field-uid: 2
// fixed:
// $ref: '#/components/schemas/Macsec.CryptoEngine.Type.StatelessEncryptionOnly.FixedPn'
// x-field-uid: 2
// incrementing:
// $ref: '#/components/schemas/Macsec.CryptoEngine.Type.StatelessEncryptionOnly.IncrementingPn'
// x-field-uid: 3
type MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn interface {
	Validation
	// msg marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn to protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	// and doesn't set defaults
	msg() *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	// setMsg unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn from protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	// and doesn't set defaults
	setMsg(*otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	// provides marshal interface
	Marshal() marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	// validate validates MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) setDefault() {

}
