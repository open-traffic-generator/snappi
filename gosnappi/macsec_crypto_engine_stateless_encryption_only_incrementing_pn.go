package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn *****
type macsecCryptoEngineStatelessEncryptionOnlyIncrementingPn struct {
	validation
	obj          *otg.MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn
	marshaller   marshalMacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn
	unMarshaller unMarshalMacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn
}

func NewMacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn() MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn {
	obj := macsecCryptoEngineStatelessEncryptionOnlyIncrementingPn{obj: &otg.MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) msg() *otg.MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn {
	return obj.obj
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) setMsg(msg *otg.MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn struct {
	obj *macsecCryptoEngineStatelessEncryptionOnlyIncrementingPn
}

type marshalMacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn interface {
	// ToProto marshals MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn to protobuf object *otg.MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn
	ToProto() (*otg.MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn, error)
	// ToPbText marshals MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn struct {
	obj *macsecCryptoEngineStatelessEncryptionOnlyIncrementingPn
}

type unMarshalMacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn interface {
	// FromProto unmarshals MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn from protobuf object *otg.MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn
	FromProto(msg *otg.MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) (MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn, error)
	// FromPbText unmarshals MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn from JSON text
	FromJson(value string) error
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) Marshal() marshalMacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) Unmarshal() unMarshalMacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) ToProto() (*otg.MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) FromProto(msg *otg.MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) (MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) ToPbText() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) FromPbText(value string) error {
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

func (m *marshalmacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) ToYaml() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) FromYaml(value string) error {
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

func (m *marshalmacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) ToJson() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) FromJson(value string) error {
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

func (obj *macsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) Clone() (MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn()
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

// MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn is incrementing PN settings.
type MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn interface {
	Validation
	// msg marshals MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn to protobuf object *otg.MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn
	// and doesn't set defaults
	msg() *otg.MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn
	// setMsg unmarshals MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn from protobuf object *otg.MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn
	// and doesn't set defaults
	setMsg(*otg.MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn
	// provides marshal interface
	Marshal() marshalMacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn
	// validate validates MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Count returns uint32, set in MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn.
	Count() uint32
	// SetCount assigns uint32 provided by user to MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn
	SetCount(value uint32) MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn
	// HasCount checks if Count has been set in MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn
	HasCount() bool
	// FirstPn returns uint32, set in MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn.
	FirstPn() uint32
	// SetFirstPn assigns uint32 provided by user to MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn
	SetFirstPn(value uint32) MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn
	// HasFirstPn checks if FirstPn has been set in MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn
	HasFirstPn() bool
	// FirstXpn returns int32, set in MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn.
	FirstXpn() int32
	// SetFirstXpn assigns int32 provided by user to MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn
	SetFirstXpn(value int32) MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn
	// HasFirstXpn checks if FirstXpn has been set in MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn
	HasFirstXpn() bool
}

// Count of packet numbers in series.
// Count returns a uint32
func (obj *macsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) Count() uint32 {

	return *obj.obj.Count

}

// Count of packet numbers in series.
// Count returns a uint32
func (obj *macsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) HasCount() bool {
	return obj.obj.Count != nil
}

// Count of packet numbers in series.
// SetCount sets the uint32 value in the MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn object
func (obj *macsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) SetCount(value uint32) MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn {

	obj.obj.Count = &value
	return obj
}

// The first PN.
// FirstPn returns a uint32
func (obj *macsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) FirstPn() uint32 {

	return *obj.obj.FirstPn

}

// The first PN.
// FirstPn returns a uint32
func (obj *macsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) HasFirstPn() bool {
	return obj.obj.FirstPn != nil
}

// The first PN.
// SetFirstPn sets the uint32 value in the MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn object
func (obj *macsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) SetFirstPn(value uint32) MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn {

	obj.obj.FirstPn = &value
	return obj
}

// The first XPN.
// FirstXpn returns a int32
func (obj *macsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) FirstXpn() int32 {

	return *obj.obj.FirstXpn

}

// The first XPN.
// FirstXpn returns a int32
func (obj *macsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) HasFirstXpn() bool {
	return obj.obj.FirstXpn != nil
}

// The first XPN.
// SetFirstXpn sets the int32 value in the MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn object
func (obj *macsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) SetFirstXpn(value int32) MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn {

	obj.obj.FirstXpn = &value
	return obj
}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Count != nil {

		if *obj.obj.Count < 2 || *obj.obj.Count > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("2 <= MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn.Count <= 4294967295 but Got %d", *obj.obj.Count))
		}

	}

	if obj.obj.FirstPn != nil {

		if *obj.obj.FirstPn < 1 || *obj.obj.FirstPn > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn.FirstPn <= 4294967295 but Got %d", *obj.obj.FirstPn))
		}

	}

	if obj.obj.FirstXpn != nil {

		if *obj.obj.FirstXpn < 1 || *obj.obj.FirstXpn > 2147483647 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= MacsecCryptoEngineStatelessEncryptionOnlyIncrementingPn.FirstXpn <= 2147483647 but Got %d", *obj.obj.FirstXpn))
		}

	}

}

func (obj *macsecCryptoEngineStatelessEncryptionOnlyIncrementingPn) setDefault() {
	if obj.obj.Count == nil {
		obj.SetCount(100)
	}
	if obj.obj.FirstPn == nil {
		obj.SetFirstPn(10000)
	}
	if obj.obj.FirstXpn == nil {
		obj.SetFirstXpn(65536)
	}

}
