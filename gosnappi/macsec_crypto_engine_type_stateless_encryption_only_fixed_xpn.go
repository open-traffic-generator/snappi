package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn *****
type macsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn struct {
	validation
	obj          *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn
	marshaller   marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn
	unMarshaller unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn
}

func NewMacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn() MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn {
	obj := macsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn{obj: &otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn) msg() *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn {
	return obj.obj
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn) setMsg(msg *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn) MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn struct {
	obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn
}

type marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn interface {
	// ToProto marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn to protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn
	ToProto() (*otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn, error)
	// ToPbText marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn struct {
	obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn
}

type unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn interface {
	// FromProto unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn from protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn
	FromProto(msg *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn) (MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn, error)
	// FromPbText unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn from JSON text
	FromJson(value string) error
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn) Marshal() marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn) Unmarshal() unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn) ToProto() (*otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn) FromProto(msg *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn) (MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn) ToPbText() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn) FromPbText(value string) error {
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

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn) ToYaml() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn) FromYaml(value string) error {
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

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn) ToJson() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn) FromJson(value string) error {
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

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn) Clone() (MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn()
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

// MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn is fixed XPN settings.
type MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn interface {
	Validation
	// msg marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn to protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn
	// and doesn't set defaults
	msg() *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn
	// setMsg unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn from protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn
	// and doesn't set defaults
	setMsg(*otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn) MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn
	// provides marshal interface
	Marshal() marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn
	// validate validates MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Xpn returns int32, set in MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn.
	Xpn() int32
	// SetXpn assigns int32 provided by user to MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn
	SetXpn(value int32) MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn
	// HasXpn checks if Xpn has been set in MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn
	HasXpn() bool
}

// Fixed Tx XPN. 8 bytes PN with which all packets will be sent out.
// Xpn returns a int32
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn) Xpn() int32 {

	return *obj.obj.Xpn

}

// Fixed Tx XPN. 8 bytes PN with which all packets will be sent out.
// Xpn returns a int32
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn) HasXpn() bool {
	return obj.obj.Xpn != nil
}

// Fixed Tx XPN. 8 bytes PN with which all packets will be sent out.
// SetXpn sets the int32 value in the MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn object
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn) SetXpn(value int32) MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn {

	obj.obj.Xpn = &value
	return obj
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Xpn != nil {

		if *obj.obj.Xpn < 1 || *obj.obj.Xpn > 2147483647 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= MacsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn.Xpn <= 2147483647 but Got %d", *obj.obj.Xpn))
		}

	}

}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyFixedXpn) setDefault() {
	if obj.obj.Xpn == nil {
		obj.SetXpn(6)
	}

}
