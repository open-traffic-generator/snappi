package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption *****
type macsecSecYCryptoEngineTypeStatefulEncryptionDecryption struct {
	validation
	obj                        *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption
	marshaller                 marshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryption
	unMarshaller               unMarshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryption
	initialPnHolder            MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	hardwareAccelerationHolder MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
}

func NewMacsecSecYCryptoEngineTypeStatefulEncryptionDecryption() MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption {
	obj := macsecSecYCryptoEngineTypeStatefulEncryptionDecryption{obj: &otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryption) msg() *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption {
	return obj.obj
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryption) setMsg(msg *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryption struct {
	obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryption
}

type marshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryption interface {
	// ToProto marshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption to protobuf object *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption
	ToProto() (*otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption, error)
	// ToPbText marshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryption struct {
	obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryption
}

type unMarshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryption interface {
	// FromProto unmarshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption from protobuf object *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption
	FromProto(msg *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption) (MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption, error)
	// FromPbText unmarshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption from JSON text
	FromJson(value string) error
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryption) Marshal() marshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryption {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryption{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryption) Unmarshal() unMarshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryption {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryption{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryption) ToProto() (*otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryption) FromProto(msg *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption) (MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryption) ToPbText() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryption) FromPbText(value string) error {
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

func (m *marshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryption) ToYaml() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryption) FromYaml(value string) error {
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

func (m *marshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryption) ToJson() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryption) FromJson(value string) error {
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

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryption) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryption) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryption) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryption) Clone() (MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecSecYCryptoEngineTypeStatefulEncryptionDecryption()
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

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryption) setNil() {
	obj.initialPnHolder = nil
	obj.hardwareAccelerationHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption is the container for stateful encryption and decryption engine settings.
type MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption interface {
	Validation
	// msg marshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption to protobuf object *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption
	// and doesn't set defaults
	msg() *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption
	// setMsg unmarshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption from protobuf object *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption
	// and doesn't set defaults
	setMsg(*otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption
	// provides marshal interface
	Marshal() marshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryption
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryption
	// validate validates MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// InitialPn returns MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn, set in MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption.
	// MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn is initial PN settings.
	InitialPn() MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
	// SetInitialPn assigns MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn provided by user to MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption.
	// MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn is initial PN settings.
	SetInitialPn(value MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption
	// HasInitialPn checks if InitialPn has been set in MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption
	HasInitialPn() bool
	// HardwareAcceleration returns MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration, set in MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption.
	// MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration is hardware acceleration settings.
	HardwareAcceleration() MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
	// SetHardwareAcceleration assigns MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration provided by user to MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption.
	// MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration is hardware acceleration settings.
	SetHardwareAcceleration(value MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption
	// HasHardwareAcceleration checks if HardwareAcceleration has been set in MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption
	HasHardwareAcceleration() bool
	setNil()
}

// description is TBD
// InitialPn returns a MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryption) InitialPn() MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn {
	if obj.obj.InitialPn == nil {
		obj.obj.InitialPn = NewMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn().msg()
	}
	if obj.initialPnHolder == nil {
		obj.initialPnHolder = &macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn{obj: obj.obj.InitialPn}
	}
	return obj.initialPnHolder
}

// description is TBD
// InitialPn returns a MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn
func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryption) HasInitialPn() bool {
	return obj.obj.InitialPn != nil
}

// description is TBD
// SetInitialPn sets the MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn value in the MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption object
func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryption) SetInitialPn(value MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionInitialPn) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption {

	obj.initialPnHolder = nil
	obj.obj.InitialPn = value.msg()

	return obj
}

// description is TBD
// HardwareAcceleration returns a MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryption) HardwareAcceleration() MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration {
	if obj.obj.HardwareAcceleration == nil {
		obj.obj.HardwareAcceleration = NewMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration().msg()
	}
	if obj.hardwareAccelerationHolder == nil {
		obj.hardwareAccelerationHolder = &macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration{obj: obj.obj.HardwareAcceleration}
	}
	return obj.hardwareAccelerationHolder
}

// description is TBD
// HardwareAcceleration returns a MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration
func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryption) HasHardwareAcceleration() bool {
	return obj.obj.HardwareAcceleration != nil
}

// description is TBD
// SetHardwareAcceleration sets the MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration value in the MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption object
func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryption) SetHardwareAcceleration(value MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAcceleration) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryption {

	obj.hardwareAccelerationHolder = nil
	obj.obj.HardwareAcceleration = value.msg()

	return obj
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryption) validateObj(vObj *validation, set_default bool) {
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

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryption) setDefault() {

}
