package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn *****
type macsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn struct {
	validation
	obj          *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	marshaller   marshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	unMarshaller unMarshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
}

func NewMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn() MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn {
	obj := macsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn{obj: &otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) msg() *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn {
	return obj.obj
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) setMsg(msg *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn struct {
	obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
}

type marshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn interface {
	// ToProto marshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn to protobuf object *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	ToProto() (*otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn, error)
	// ToPbText marshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn struct {
	obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
}

type unMarshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn interface {
	// FromProto unmarshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn from protobuf object *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	FromProto(msg *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) (MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn, error)
	// FromPbText unmarshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn from JSON text
	FromJson(value string) error
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) Marshal() marshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) Unmarshal() unMarshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) ToProto() (*otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) FromProto(msg *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) (MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) ToPbText() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) FromPbText(value string) error {
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

func (m *marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) ToYaml() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) FromYaml(value string) error {
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

func (m *marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) ToJson() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) FromJson(value string) error {
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

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) Clone() (MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn()
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

// MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn is incrementing PN settings.
type MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn interface {
	Validation
	// msg marshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn to protobuf object *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	// and doesn't set defaults
	msg() *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	// setMsg unmarshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn from protobuf object *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	// and doesn't set defaults
	setMsg(*otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	// provides marshal interface
	Marshal() marshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	// validate validates MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Count returns int32, set in MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn.
	Count() int32
	// SetCount assigns int32 provided by user to MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	SetCount(value int32) MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	// HasCount checks if Count has been set in MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	HasCount() bool
	// First returns int32, set in MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn.
	First() int32
	// SetFirst assigns int32 provided by user to MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	SetFirst(value int32) MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	// HasFirst checks if First has been set in MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	HasFirst() bool
}

// Count of packet numbers in series.
// Count returns a int32
func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) Count() int32 {

	return *obj.obj.Count

}

// Count of packet numbers in series.
// Count returns a int32
func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) HasCount() bool {
	return obj.obj.Count != nil
}

// Count of packet numbers in series.
// SetCount sets the int32 value in the MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn object
func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) SetCount(value int32) MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn {

	obj.obj.Count = &value
	return obj
}

// First PN.
// First returns a int32
func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) First() int32 {

	return *obj.obj.First

}

// First PN.
// First returns a int32
func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) HasFirst() bool {
	return obj.obj.First != nil
}

// First PN.
// SetFirst sets the int32 value in the MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn object
func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) SetFirst(value int32) MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn {

	obj.obj.First = &value
	return obj
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) setDefault() {
	if obj.obj.Count == nil {
		obj.SetCount(100)
	}
	if obj.obj.First == nil {
		obj.SetFirst(10000)
	}

}
