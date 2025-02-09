package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn *****
type macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn struct {
	validation
	obj          *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn
	marshaller   marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn
	unMarshaller unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn
}

func NewMacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn() MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn {
	obj := macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn{obj: &otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) msg() *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn {
	return obj.obj
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) setMsg(msg *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn struct {
	obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn
}

type marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn interface {
	// ToProto marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn to protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn
	ToProto() (*otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn, error)
	// ToPbText marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn struct {
	obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn
}

type unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn interface {
	// FromProto unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn from protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn
	FromProto(msg *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) (MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn, error)
	// FromPbText unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn from JSON text
	FromJson(value string) error
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) Marshal() marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) Unmarshal() unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) ToProto() (*otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) FromProto(msg *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) (MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) ToPbText() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) FromPbText(value string) error {
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

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) ToYaml() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) FromYaml(value string) error {
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

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) ToJson() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) FromJson(value string) error {
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

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) Clone() (MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn()
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

// MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn is incrementing XPN settings.
type MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn interface {
	Validation
	// msg marshals MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn to protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn
	// and doesn't set defaults
	msg() *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn
	// setMsg unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn from protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn
	// and doesn't set defaults
	setMsg(*otg.MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn
	// provides marshal interface
	Marshal() marshalMacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn
	// validate validates MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Count returns uint32, set in MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn.
	Count() uint32
	// SetCount assigns uint32 provided by user to MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn
	SetCount(value uint32) MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn
	// HasCount checks if Count has been set in MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn
	HasCount() bool
	// First returns int32, set in MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn.
	First() int32
	// SetFirst assigns int32 provided by user to MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn
	SetFirst(value int32) MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn
	// HasFirst checks if First has been set in MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn
	HasFirst() bool
}

// Count of packet numbers in series.
// Count returns a uint32
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) Count() uint32 {

	return *obj.obj.Count

}

// Count of packet numbers in series.
// Count returns a uint32
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) HasCount() bool {
	return obj.obj.Count != nil
}

// Count of packet numbers in series.
// SetCount sets the uint32 value in the MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn object
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) SetCount(value uint32) MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn {

	obj.obj.Count = &value
	return obj
}

// First XPN.
// First returns a int32
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) First() int32 {

	return *obj.obj.First

}

// First XPN.
// First returns a int32
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) HasFirst() bool {
	return obj.obj.First != nil
}

// First XPN.
// SetFirst sets the int32 value in the MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn object
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) SetFirst(value int32) MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn {

	obj.obj.First = &value
	return obj
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Count != nil {

		if *obj.obj.Count < 2 || *obj.obj.Count > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("2 <= MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn.Count <= 4294967295 but Got %d", *obj.obj.Count))
		}

	}

	if obj.obj.First != nil {

		if *obj.obj.First < 1 || *obj.obj.First > 2147483647 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= MacsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn.First <= 2147483647 but Got %d", *obj.obj.First))
		}

	}

}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnlyIncrementingXpn) setDefault() {
	if obj.obj.Count == nil {
		obj.SetCount(100)
	}
	if obj.obj.First == nil {
		obj.SetFirst(65536)
	}

}
