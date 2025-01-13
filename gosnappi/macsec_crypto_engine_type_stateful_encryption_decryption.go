package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecCryptoEngineTypeStatefulEncryptionDecryption *****
type macsecCryptoEngineTypeStatefulEncryptionDecryption struct {
	validation
	obj                        *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryption
	marshaller                 marshalMacsecCryptoEngineTypeStatefulEncryptionDecryption
	unMarshaller               unMarshalMacsecCryptoEngineTypeStatefulEncryptionDecryption
	initialPnHolder            MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	hardwareAccelerationHolder MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
}

func NewMacsecCryptoEngineTypeStatefulEncryptionDecryption() MacsecCryptoEngineTypeStatefulEncryptionDecryption {
	obj := macsecCryptoEngineTypeStatefulEncryptionDecryption{obj: &otg.MacsecCryptoEngineTypeStatefulEncryptionDecryption{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryption) msg() *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryption {
	return obj.obj
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryption) setMsg(msg *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryption) MacsecCryptoEngineTypeStatefulEncryptionDecryption {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecCryptoEngineTypeStatefulEncryptionDecryption struct {
	obj *macsecCryptoEngineTypeStatefulEncryptionDecryption
}

type marshalMacsecCryptoEngineTypeStatefulEncryptionDecryption interface {
	// ToProto marshals MacsecCryptoEngineTypeStatefulEncryptionDecryption to protobuf object *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryption
	ToProto() (*otg.MacsecCryptoEngineTypeStatefulEncryptionDecryption, error)
	// ToPbText marshals MacsecCryptoEngineTypeStatefulEncryptionDecryption to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecCryptoEngineTypeStatefulEncryptionDecryption to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecCryptoEngineTypeStatefulEncryptionDecryption to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecCryptoEngineTypeStatefulEncryptionDecryption struct {
	obj *macsecCryptoEngineTypeStatefulEncryptionDecryption
}

type unMarshalMacsecCryptoEngineTypeStatefulEncryptionDecryption interface {
	// FromProto unmarshals MacsecCryptoEngineTypeStatefulEncryptionDecryption from protobuf object *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryption
	FromProto(msg *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryption) (MacsecCryptoEngineTypeStatefulEncryptionDecryption, error)
	// FromPbText unmarshals MacsecCryptoEngineTypeStatefulEncryptionDecryption from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecCryptoEngineTypeStatefulEncryptionDecryption from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecCryptoEngineTypeStatefulEncryptionDecryption from JSON text
	FromJson(value string) error
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryption) Marshal() marshalMacsecCryptoEngineTypeStatefulEncryptionDecryption {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecCryptoEngineTypeStatefulEncryptionDecryption{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryption) Unmarshal() unMarshalMacsecCryptoEngineTypeStatefulEncryptionDecryption {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecCryptoEngineTypeStatefulEncryptionDecryption{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecCryptoEngineTypeStatefulEncryptionDecryption) ToProto() (*otg.MacsecCryptoEngineTypeStatefulEncryptionDecryption, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecCryptoEngineTypeStatefulEncryptionDecryption) FromProto(msg *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryption) (MacsecCryptoEngineTypeStatefulEncryptionDecryption, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecCryptoEngineTypeStatefulEncryptionDecryption) ToPbText() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatefulEncryptionDecryption) FromPbText(value string) error {
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

func (m *marshalmacsecCryptoEngineTypeStatefulEncryptionDecryption) ToYaml() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatefulEncryptionDecryption) FromYaml(value string) error {
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

func (m *marshalmacsecCryptoEngineTypeStatefulEncryptionDecryption) ToJson() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatefulEncryptionDecryption) FromJson(value string) error {
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

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryption) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryption) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryption) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryption) Clone() (MacsecCryptoEngineTypeStatefulEncryptionDecryption, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecCryptoEngineTypeStatefulEncryptionDecryption()
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

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryption) setNil() {
	obj.initialPnHolder = nil
	obj.hardwareAccelerationHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecCryptoEngineTypeStatefulEncryptionDecryption is the container for stateful encryption and decryption engine settings.
type MacsecCryptoEngineTypeStatefulEncryptionDecryption interface {
	Validation
	// msg marshals MacsecCryptoEngineTypeStatefulEncryptionDecryption to protobuf object *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryption
	// and doesn't set defaults
	msg() *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryption
	// setMsg unmarshals MacsecCryptoEngineTypeStatefulEncryptionDecryption from protobuf object *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryption
	// and doesn't set defaults
	setMsg(*otg.MacsecCryptoEngineTypeStatefulEncryptionDecryption) MacsecCryptoEngineTypeStatefulEncryptionDecryption
	// provides marshal interface
	Marshal() marshalMacsecCryptoEngineTypeStatefulEncryptionDecryption
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecCryptoEngineTypeStatefulEncryptionDecryption
	// validate validates MacsecCryptoEngineTypeStatefulEncryptionDecryption
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecCryptoEngineTypeStatefulEncryptionDecryption, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// InitialPn returns MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn, set in MacsecCryptoEngineTypeStatefulEncryptionDecryption.
	// MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn is initial PN settings.
	InitialPn() MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	// SetInitialPn assigns MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn provided by user to MacsecCryptoEngineTypeStatefulEncryptionDecryption.
	// MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn is initial PN settings.
	SetInitialPn(value MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) MacsecCryptoEngineTypeStatefulEncryptionDecryption
	// HasInitialPn checks if InitialPn has been set in MacsecCryptoEngineTypeStatefulEncryptionDecryption
	HasInitialPn() bool
	// HardwareAcceleration returns MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration, set in MacsecCryptoEngineTypeStatefulEncryptionDecryption.
	// MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration is hardware acceleration settings.
	HardwareAcceleration() MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	// SetHardwareAcceleration assigns MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration provided by user to MacsecCryptoEngineTypeStatefulEncryptionDecryption.
	// MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration is hardware acceleration settings.
	SetHardwareAcceleration(value MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) MacsecCryptoEngineTypeStatefulEncryptionDecryption
	// HasHardwareAcceleration checks if HardwareAcceleration has been set in MacsecCryptoEngineTypeStatefulEncryptionDecryption
	HasHardwareAcceleration() bool
	setNil()
}

// description is TBD
// InitialPn returns a MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryption) InitialPn() MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn {
	if obj.obj.InitialPn == nil {
		obj.obj.InitialPn = NewMacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn().msg()
	}
	if obj.initialPnHolder == nil {
		obj.initialPnHolder = &macsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn{obj: obj.obj.InitialPn}
	}
	return obj.initialPnHolder
}

// description is TBD
// InitialPn returns a MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryption) HasInitialPn() bool {
	return obj.obj.InitialPn != nil
}

// description is TBD
// SetInitialPn sets the MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn value in the MacsecCryptoEngineTypeStatefulEncryptionDecryption object
func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryption) SetInitialPn(value MacsecCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) MacsecCryptoEngineTypeStatefulEncryptionDecryption {

	obj.initialPnHolder = nil
	obj.obj.InitialPn = value.msg()

	return obj
}

// description is TBD
// HardwareAcceleration returns a MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryption) HardwareAcceleration() MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration {
	if obj.obj.HardwareAcceleration == nil {
		obj.obj.HardwareAcceleration = NewMacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration().msg()
	}
	if obj.hardwareAccelerationHolder == nil {
		obj.hardwareAccelerationHolder = &macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration{obj: obj.obj.HardwareAcceleration}
	}
	return obj.hardwareAccelerationHolder
}

// description is TBD
// HardwareAcceleration returns a MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryption) HasHardwareAcceleration() bool {
	return obj.obj.HardwareAcceleration != nil
}

// description is TBD
// SetHardwareAcceleration sets the MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration value in the MacsecCryptoEngineTypeStatefulEncryptionDecryption object
func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryption) SetHardwareAcceleration(value MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) MacsecCryptoEngineTypeStatefulEncryptionDecryption {

	obj.hardwareAccelerationHolder = nil
	obj.obj.HardwareAcceleration = value.msg()

	return obj
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryption) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.InitialPn != nil {

		obj.InitialPn().validateObj(vObj, set_default)
	}

	if obj.obj.HardwareAcceleration != nil {

		obj.HardwareAcceleration().validateObj(vObj, set_default)
	}

}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryption) setDefault() {

}
