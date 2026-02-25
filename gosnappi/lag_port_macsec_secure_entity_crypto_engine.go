package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LagPortMacsecSecureEntityCryptoEngine *****
type lagPortMacsecSecureEntityCryptoEngine struct {
	validation
	obj                  *otg.LagPortMacsecSecureEntityCryptoEngine
	marshaller           marshalLagPortMacsecSecureEntityCryptoEngine
	unMarshaller         unMarshalLagPortMacsecSecureEntityCryptoEngine
	encryptDecryptHolder LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt
}

func NewLagPortMacsecSecureEntityCryptoEngine() LagPortMacsecSecureEntityCryptoEngine {
	obj := lagPortMacsecSecureEntityCryptoEngine{obj: &otg.LagPortMacsecSecureEntityCryptoEngine{}}
	obj.setDefault()
	return &obj
}

func (obj *lagPortMacsecSecureEntityCryptoEngine) msg() *otg.LagPortMacsecSecureEntityCryptoEngine {
	return obj.obj
}

func (obj *lagPortMacsecSecureEntityCryptoEngine) setMsg(msg *otg.LagPortMacsecSecureEntityCryptoEngine) LagPortMacsecSecureEntityCryptoEngine {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshallagPortMacsecSecureEntityCryptoEngine struct {
	obj *lagPortMacsecSecureEntityCryptoEngine
}

type marshalLagPortMacsecSecureEntityCryptoEngine interface {
	// ToProto marshals LagPortMacsecSecureEntityCryptoEngine to protobuf object *otg.LagPortMacsecSecureEntityCryptoEngine
	ToProto() (*otg.LagPortMacsecSecureEntityCryptoEngine, error)
	// ToPbText marshals LagPortMacsecSecureEntityCryptoEngine to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LagPortMacsecSecureEntityCryptoEngine to YAML text
	ToYaml() (string, error)
	// ToJson marshals LagPortMacsecSecureEntityCryptoEngine to JSON text
	ToJson() (string, error)
}

type unMarshallagPortMacsecSecureEntityCryptoEngine struct {
	obj *lagPortMacsecSecureEntityCryptoEngine
}

type unMarshalLagPortMacsecSecureEntityCryptoEngine interface {
	// FromProto unmarshals LagPortMacsecSecureEntityCryptoEngine from protobuf object *otg.LagPortMacsecSecureEntityCryptoEngine
	FromProto(msg *otg.LagPortMacsecSecureEntityCryptoEngine) (LagPortMacsecSecureEntityCryptoEngine, error)
	// FromPbText unmarshals LagPortMacsecSecureEntityCryptoEngine from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LagPortMacsecSecureEntityCryptoEngine from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LagPortMacsecSecureEntityCryptoEngine from JSON text
	FromJson(value string) error
}

func (obj *lagPortMacsecSecureEntityCryptoEngine) Marshal() marshalLagPortMacsecSecureEntityCryptoEngine {
	if obj.marshaller == nil {
		obj.marshaller = &marshallagPortMacsecSecureEntityCryptoEngine{obj: obj}
	}
	return obj.marshaller
}

func (obj *lagPortMacsecSecureEntityCryptoEngine) Unmarshal() unMarshalLagPortMacsecSecureEntityCryptoEngine {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallagPortMacsecSecureEntityCryptoEngine{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallagPortMacsecSecureEntityCryptoEngine) ToProto() (*otg.LagPortMacsecSecureEntityCryptoEngine, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallagPortMacsecSecureEntityCryptoEngine) FromProto(msg *otg.LagPortMacsecSecureEntityCryptoEngine) (LagPortMacsecSecureEntityCryptoEngine, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallagPortMacsecSecureEntityCryptoEngine) ToPbText() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntityCryptoEngine) FromPbText(value string) error {
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

func (m *marshallagPortMacsecSecureEntityCryptoEngine) ToYaml() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntityCryptoEngine) FromYaml(value string) error {
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

func (m *marshallagPortMacsecSecureEntityCryptoEngine) ToJson() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntityCryptoEngine) FromJson(value string) error {
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

func (obj *lagPortMacsecSecureEntityCryptoEngine) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lagPortMacsecSecureEntityCryptoEngine) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lagPortMacsecSecureEntityCryptoEngine) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lagPortMacsecSecureEntityCryptoEngine) Clone() (LagPortMacsecSecureEntityCryptoEngine, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLagPortMacsecSecureEntityCryptoEngine()
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

func (obj *lagPortMacsecSecureEntityCryptoEngine) setNil() {
	obj.encryptDecryptHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// LagPortMacsecSecureEntityCryptoEngine is a container of crypto engine properties of a SecY.
type LagPortMacsecSecureEntityCryptoEngine interface {
	Validation
	// msg marshals LagPortMacsecSecureEntityCryptoEngine to protobuf object *otg.LagPortMacsecSecureEntityCryptoEngine
	// and doesn't set defaults
	msg() *otg.LagPortMacsecSecureEntityCryptoEngine
	// setMsg unmarshals LagPortMacsecSecureEntityCryptoEngine from protobuf object *otg.LagPortMacsecSecureEntityCryptoEngine
	// and doesn't set defaults
	setMsg(*otg.LagPortMacsecSecureEntityCryptoEngine) LagPortMacsecSecureEntityCryptoEngine
	// provides marshal interface
	Marshal() marshalLagPortMacsecSecureEntityCryptoEngine
	// provides unmarshal interface
	Unmarshal() unMarshalLagPortMacsecSecureEntityCryptoEngine
	// validate validates LagPortMacsecSecureEntityCryptoEngine
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LagPortMacsecSecureEntityCryptoEngine, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EncryptDecrypt returns LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt, set in LagPortMacsecSecureEntityCryptoEngine.
	// LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt is the container for configuration of crypto engine of encrypt and decrypt type. Such engine can both encrypt transmitted packets and decrypt packets on arrival. It can have hardware acceleration for faster encryption/ decryption. As both encryption and decryption are possible, stateful (e.g. TCP) traffic can be sent/ received.
	EncryptDecrypt() LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt
	// SetEncryptDecrypt assigns LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt provided by user to LagPortMacsecSecureEntityCryptoEngine.
	// LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt is the container for configuration of crypto engine of encrypt and decrypt type. Such engine can both encrypt transmitted packets and decrypt packets on arrival. It can have hardware acceleration for faster encryption/ decryption. As both encryption and decryption are possible, stateful (e.g. TCP) traffic can be sent/ received.
	SetEncryptDecrypt(value LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) LagPortMacsecSecureEntityCryptoEngine
	// HasEncryptDecrypt checks if EncryptDecrypt has been set in LagPortMacsecSecureEntityCryptoEngine
	HasEncryptDecrypt() bool
	setNil()
}

// description is TBD
// EncryptDecrypt returns a LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt
func (obj *lagPortMacsecSecureEntityCryptoEngine) EncryptDecrypt() LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt {
	if obj.obj.EncryptDecrypt == nil {
		obj.obj.EncryptDecrypt = NewLagPortMacsecSecureEntityCryptoEngineEncryptDecrypt().msg()
	}
	if obj.encryptDecryptHolder == nil {
		obj.encryptDecryptHolder = &lagPortMacsecSecureEntityCryptoEngineEncryptDecrypt{obj: obj.obj.EncryptDecrypt}
	}
	return obj.encryptDecryptHolder
}

// description is TBD
// EncryptDecrypt returns a LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt
func (obj *lagPortMacsecSecureEntityCryptoEngine) HasEncryptDecrypt() bool {
	return obj.obj.EncryptDecrypt != nil
}

// description is TBD
// SetEncryptDecrypt sets the LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt value in the LagPortMacsecSecureEntityCryptoEngine object
func (obj *lagPortMacsecSecureEntityCryptoEngine) SetEncryptDecrypt(value LagPortMacsecSecureEntityCryptoEngineEncryptDecrypt) LagPortMacsecSecureEntityCryptoEngine {

	obj.encryptDecryptHolder = nil
	obj.obj.EncryptDecrypt = value.msg()

	return obj
}

func (obj *lagPortMacsecSecureEntityCryptoEngine) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.EncryptDecrypt != nil {

		obj.EncryptDecrypt().validateObj(vObj, set_default)
	}

}

func (obj *lagPortMacsecSecureEntityCryptoEngine) setDefault() {

}
