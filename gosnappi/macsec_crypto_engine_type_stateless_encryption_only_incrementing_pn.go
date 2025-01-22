package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn *****
type macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn struct {
	validation
	obj          *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	marshaller   marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	unMarshaller unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
}

func NewMacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn() MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn {
	obj := macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn{obj: &otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) msg() *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn {
	return obj.obj
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) setMsg(msg *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn struct {
	obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
}

type marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn interface {
	// ToProto marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn to protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	ToProto() (*otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn, error)
	// ToPbText marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn struct {
	obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
}

type unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn interface {
	// FromProto unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn from protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	FromProto(msg *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) (MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn, error)
	// FromPbText unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn from JSON text
	FromJson(value string) error
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) Marshal() marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) Unmarshal() unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) ToProto() (*otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) FromProto(msg *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) (MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) ToPbText() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) FromPbText(value string) error {
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

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) ToYaml() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) FromYaml(value string) error {
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

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) ToJson() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) FromJson(value string) error {
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

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) Clone() (MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn()
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

// MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn is incrementing PN settings.
type MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn interface {
	Validation
	// msg marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn to protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	// and doesn't set defaults
	msg() *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	// setMsg unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn from protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	// and doesn't set defaults
	setMsg(*otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	// provides marshal interface
	Marshal() marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	// validate validates MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Count returns uint32, set in MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn.
	Count() uint32
	// SetCount assigns uint32 provided by user to MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	SetCount(value uint32) MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	// HasCount checks if Count has been set in MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	HasCount() bool
	// First returns uint32, set in MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn.
	First() uint32
	// SetFirst assigns uint32 provided by user to MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	SetFirst(value uint32) MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	// HasFirst checks if First has been set in MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn
	HasFirst() bool
}

// Count of packet numbers in series.
// Count returns a uint32
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) Count() uint32 {

	return *obj.obj.Count

}

// Count of packet numbers in series.
// Count returns a uint32
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) HasCount() bool {
	return obj.obj.Count != nil
}

// Count of packet numbers in series.
// SetCount sets the uint32 value in the MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn object
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) SetCount(value uint32) MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn {

	obj.obj.Count = &value
	return obj
}

// First PN.
// First returns a uint32
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) First() uint32 {

	return *obj.obj.First

}

// First PN.
// First returns a uint32
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) HasFirst() bool {
	return obj.obj.First != nil
}

// First PN.
// SetFirst sets the uint32 value in the MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn object
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) SetFirst(value uint32) MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn {

	obj.obj.First = &value
	return obj
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Count != nil {

		if *obj.obj.Count < 2 || *obj.obj.Count > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("2 <= MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn.Count <= 4294967295 but Got %d", *obj.obj.Count))
		}

	}

	if obj.obj.First != nil {

		if *obj.obj.First < 1 || *obj.obj.First > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn.First <= 4294967295 but Got %d", *obj.obj.First))
		}

	}

}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingPn) setDefault() {
	if obj.obj.Count == nil {
		obj.SetCount(100)
	}
	if obj.obj.First == nil {
		obj.SetFirst(10000)
	}

}
