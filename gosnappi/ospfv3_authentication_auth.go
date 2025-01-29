package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3AuthenticationAuth *****
type ospfv3AuthenticationAuth struct {
	validation
	obj          *otg.Ospfv3AuthenticationAuth
	marshaller   marshalOspfv3AuthenticationAuth
	unMarshaller unMarshalOspfv3AuthenticationAuth
	algoHolder   Ospfv3AuthenticationAlgo
}

func NewOspfv3AuthenticationAuth() Ospfv3AuthenticationAuth {
	obj := ospfv3AuthenticationAuth{obj: &otg.Ospfv3AuthenticationAuth{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3AuthenticationAuth) msg() *otg.Ospfv3AuthenticationAuth {
	return obj.obj
}

func (obj *ospfv3AuthenticationAuth) setMsg(msg *otg.Ospfv3AuthenticationAuth) Ospfv3AuthenticationAuth {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3AuthenticationAuth struct {
	obj *ospfv3AuthenticationAuth
}

type marshalOspfv3AuthenticationAuth interface {
	// ToProto marshals Ospfv3AuthenticationAuth to protobuf object *otg.Ospfv3AuthenticationAuth
	ToProto() (*otg.Ospfv3AuthenticationAuth, error)
	// ToPbText marshals Ospfv3AuthenticationAuth to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3AuthenticationAuth to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3AuthenticationAuth to JSON text
	ToJson() (string, error)
}

type unMarshalospfv3AuthenticationAuth struct {
	obj *ospfv3AuthenticationAuth
}

type unMarshalOspfv3AuthenticationAuth interface {
	// FromProto unmarshals Ospfv3AuthenticationAuth from protobuf object *otg.Ospfv3AuthenticationAuth
	FromProto(msg *otg.Ospfv3AuthenticationAuth) (Ospfv3AuthenticationAuth, error)
	// FromPbText unmarshals Ospfv3AuthenticationAuth from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3AuthenticationAuth from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3AuthenticationAuth from JSON text
	FromJson(value string) error
}

func (obj *ospfv3AuthenticationAuth) Marshal() marshalOspfv3AuthenticationAuth {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3AuthenticationAuth{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3AuthenticationAuth) Unmarshal() unMarshalOspfv3AuthenticationAuth {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3AuthenticationAuth{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3AuthenticationAuth) ToProto() (*otg.Ospfv3AuthenticationAuth, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3AuthenticationAuth) FromProto(msg *otg.Ospfv3AuthenticationAuth) (Ospfv3AuthenticationAuth, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3AuthenticationAuth) ToPbText() (string, error) {
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

func (m *unMarshalospfv3AuthenticationAuth) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalospfv3AuthenticationAuth) ToYaml() (string, error) {
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

func (m *unMarshalospfv3AuthenticationAuth) FromYaml(value string) error {
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
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalospfv3AuthenticationAuth) ToJson() (string, error) {
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

func (m *unMarshalospfv3AuthenticationAuth) FromJson(value string) error {
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
	m.obj.setNil()
	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *ospfv3AuthenticationAuth) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3AuthenticationAuth) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3AuthenticationAuth) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3AuthenticationAuth) Clone() (Ospfv3AuthenticationAuth, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3AuthenticationAuth()
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

func (obj *ospfv3AuthenticationAuth) setNil() {
	obj.algoHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv3AuthenticationAuth is container of OSPFv3 interface Cryptographic authentication.
type Ospfv3AuthenticationAuth interface {
	Validation
	// msg marshals Ospfv3AuthenticationAuth to protobuf object *otg.Ospfv3AuthenticationAuth
	// and doesn't set defaults
	msg() *otg.Ospfv3AuthenticationAuth
	// setMsg unmarshals Ospfv3AuthenticationAuth from protobuf object *otg.Ospfv3AuthenticationAuth
	// and doesn't set defaults
	setMsg(*otg.Ospfv3AuthenticationAuth) Ospfv3AuthenticationAuth
	// provides marshal interface
	Marshal() marshalOspfv3AuthenticationAuth
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3AuthenticationAuth
	// validate validates Ospfv3AuthenticationAuth
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3AuthenticationAuth, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Algo returns Ospfv3AuthenticationAlgo, set in Ospfv3AuthenticationAuth.
	// Ospfv3AuthenticationAlgo is the authentication algorithm used with the OSPFv3 SA.
	// - HMAC-SHA-1
	// - HMAC-SHA-256
	// - HMAC-SHA-384
	// - HMAC-SHA-512
	Algo() Ospfv3AuthenticationAlgo
	// SetAlgo assigns Ospfv3AuthenticationAlgo provided by user to Ospfv3AuthenticationAuth.
	// Ospfv3AuthenticationAlgo is the authentication algorithm used with the OSPFv3 SA.
	// - HMAC-SHA-1
	// - HMAC-SHA-256
	// - HMAC-SHA-384
	// - HMAC-SHA-512
	SetAlgo(value Ospfv3AuthenticationAlgo) Ospfv3AuthenticationAuth
	// HasAlgo checks if Algo has been set in Ospfv3AuthenticationAuth
	HasAlgo() bool
	// KeyId returns uint32, set in Ospfv3AuthenticationAuth.
	KeyId() uint32
	// SetKeyId assigns uint32 provided by user to Ospfv3AuthenticationAuth
	SetKeyId(value uint32) Ospfv3AuthenticationAuth
	// HasKeyId checks if KeyId has been set in Ospfv3AuthenticationAuth
	HasKeyId() bool
	// Key returns string, set in Ospfv3AuthenticationAuth.
	Key() string
	// SetKey assigns string provided by user to Ospfv3AuthenticationAuth
	SetKey(value string) Ospfv3AuthenticationAuth
	// HasKey checks if Key has been set in Ospfv3AuthenticationAuth
	HasKey() bool
	setNil()
}

// The authentication algorithm used with the OSPFv3 SA.
// Algo returns a Ospfv3AuthenticationAlgo
func (obj *ospfv3AuthenticationAuth) Algo() Ospfv3AuthenticationAlgo {
	if obj.obj.Algo == nil {
		obj.obj.Algo = NewOspfv3AuthenticationAlgo().msg()
	}
	if obj.algoHolder == nil {
		obj.algoHolder = &ospfv3AuthenticationAlgo{obj: obj.obj.Algo}
	}
	return obj.algoHolder
}

// The authentication algorithm used with the OSPFv3 SA.
// Algo returns a Ospfv3AuthenticationAlgo
func (obj *ospfv3AuthenticationAuth) HasAlgo() bool {
	return obj.obj.Algo != nil
}

// The authentication algorithm used with the OSPFv3 SA.
// SetAlgo sets the Ospfv3AuthenticationAlgo value in the Ospfv3AuthenticationAuth object
func (obj *ospfv3AuthenticationAuth) SetAlgo(value Ospfv3AuthenticationAlgo) Ospfv3AuthenticationAuth {

	obj.algoHolder = nil
	obj.obj.Algo = value.msg()

	return obj
}

// The unique OSPFv3 Authentication Key Identifier per-interface.
// KeyId returns a uint32
func (obj *ospfv3AuthenticationAuth) KeyId() uint32 {

	return *obj.obj.KeyId

}

// The unique OSPFv3 Authentication Key Identifier per-interface.
// KeyId returns a uint32
func (obj *ospfv3AuthenticationAuth) HasKeyId() bool {
	return obj.obj.KeyId != nil
}

// The unique OSPFv3 Authentication Key Identifier per-interface.
// SetKeyId sets the uint32 value in the Ospfv3AuthenticationAuth object
func (obj *ospfv3AuthenticationAuth) SetKeyId(value uint32) Ospfv3AuthenticationAuth {

	obj.obj.KeyId = &value
	return obj
}

// An alphanumeric Cryptographic Authentication Key associated with this OSPFv3 SA.
// Key returns a string
func (obj *ospfv3AuthenticationAuth) Key() string {

	return *obj.obj.Key

}

// An alphanumeric Cryptographic Authentication Key associated with this OSPFv3 SA.
// Key returns a string
func (obj *ospfv3AuthenticationAuth) HasKey() bool {
	return obj.obj.Key != nil
}

// An alphanumeric Cryptographic Authentication Key associated with this OSPFv3 SA.
// SetKey sets the string value in the Ospfv3AuthenticationAuth object
func (obj *ospfv3AuthenticationAuth) SetKey(value string) Ospfv3AuthenticationAuth {

	obj.obj.Key = &value
	return obj
}

func (obj *ospfv3AuthenticationAuth) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Algo != nil {

		obj.Algo().validateObj(vObj, set_default)
	}

	if obj.obj.KeyId != nil {

		if *obj.obj.KeyId > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Ospfv3AuthenticationAuth.KeyId <= 65535 but Got %d", *obj.obj.KeyId))
		}

	}

	if obj.obj.Key != nil {

		if len(*obj.obj.Key) < 1 || len(*obj.obj.Key) > 64 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"1 <= length of Ospfv3AuthenticationAuth.Key <= 64 but Got %d",
					len(*obj.obj.Key)))
		}

	}

}

func (obj *ospfv3AuthenticationAuth) setDefault() {
	if obj.obj.KeyId == nil {
		obj.SetKeyId(1)
	}

}
