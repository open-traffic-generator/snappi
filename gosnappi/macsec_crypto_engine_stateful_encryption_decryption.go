package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecCryptoEngineStatefulEncryptionDecryption *****
type macsecCryptoEngineStatefulEncryptionDecryption struct {
	validation
	obj                        *otg.MacsecCryptoEngineStatefulEncryptionDecryption
	marshaller                 marshalMacsecCryptoEngineStatefulEncryptionDecryption
	unMarshaller               unMarshalMacsecCryptoEngineStatefulEncryptionDecryption
	initialPnHolder            MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn
	hardwareAccelerationHolder MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration
}

func NewMacsecCryptoEngineStatefulEncryptionDecryption() MacsecCryptoEngineStatefulEncryptionDecryption {
	obj := macsecCryptoEngineStatefulEncryptionDecryption{obj: &otg.MacsecCryptoEngineStatefulEncryptionDecryption{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryption) msg() *otg.MacsecCryptoEngineStatefulEncryptionDecryption {
	return obj.obj
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryption) setMsg(msg *otg.MacsecCryptoEngineStatefulEncryptionDecryption) MacsecCryptoEngineStatefulEncryptionDecryption {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecCryptoEngineStatefulEncryptionDecryption struct {
	obj *macsecCryptoEngineStatefulEncryptionDecryption
}

type marshalMacsecCryptoEngineStatefulEncryptionDecryption interface {
	// ToProto marshals MacsecCryptoEngineStatefulEncryptionDecryption to protobuf object *otg.MacsecCryptoEngineStatefulEncryptionDecryption
	ToProto() (*otg.MacsecCryptoEngineStatefulEncryptionDecryption, error)
	// ToPbText marshals MacsecCryptoEngineStatefulEncryptionDecryption to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecCryptoEngineStatefulEncryptionDecryption to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecCryptoEngineStatefulEncryptionDecryption to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecCryptoEngineStatefulEncryptionDecryption struct {
	obj *macsecCryptoEngineStatefulEncryptionDecryption
}

type unMarshalMacsecCryptoEngineStatefulEncryptionDecryption interface {
	// FromProto unmarshals MacsecCryptoEngineStatefulEncryptionDecryption from protobuf object *otg.MacsecCryptoEngineStatefulEncryptionDecryption
	FromProto(msg *otg.MacsecCryptoEngineStatefulEncryptionDecryption) (MacsecCryptoEngineStatefulEncryptionDecryption, error)
	// FromPbText unmarshals MacsecCryptoEngineStatefulEncryptionDecryption from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecCryptoEngineStatefulEncryptionDecryption from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecCryptoEngineStatefulEncryptionDecryption from JSON text
	FromJson(value string) error
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryption) Marshal() marshalMacsecCryptoEngineStatefulEncryptionDecryption {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecCryptoEngineStatefulEncryptionDecryption{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryption) Unmarshal() unMarshalMacsecCryptoEngineStatefulEncryptionDecryption {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecCryptoEngineStatefulEncryptionDecryption{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecCryptoEngineStatefulEncryptionDecryption) ToProto() (*otg.MacsecCryptoEngineStatefulEncryptionDecryption, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecCryptoEngineStatefulEncryptionDecryption) FromProto(msg *otg.MacsecCryptoEngineStatefulEncryptionDecryption) (MacsecCryptoEngineStatefulEncryptionDecryption, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecCryptoEngineStatefulEncryptionDecryption) ToPbText() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineStatefulEncryptionDecryption) FromPbText(value string) error {
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

func (m *marshalmacsecCryptoEngineStatefulEncryptionDecryption) ToYaml() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineStatefulEncryptionDecryption) FromYaml(value string) error {
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

func (m *marshalmacsecCryptoEngineStatefulEncryptionDecryption) ToJson() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineStatefulEncryptionDecryption) FromJson(value string) error {
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

func (obj *macsecCryptoEngineStatefulEncryptionDecryption) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryption) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryption) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryption) Clone() (MacsecCryptoEngineStatefulEncryptionDecryption, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecCryptoEngineStatefulEncryptionDecryption()
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

func (obj *macsecCryptoEngineStatefulEncryptionDecryption) setNil() {
	obj.initialPnHolder = nil
	obj.hardwareAccelerationHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecCryptoEngineStatefulEncryptionDecryption is the container for stateful encryption and decryption engine configuration.
type MacsecCryptoEngineStatefulEncryptionDecryption interface {
	Validation
	// msg marshals MacsecCryptoEngineStatefulEncryptionDecryption to protobuf object *otg.MacsecCryptoEngineStatefulEncryptionDecryption
	// and doesn't set defaults
	msg() *otg.MacsecCryptoEngineStatefulEncryptionDecryption
	// setMsg unmarshals MacsecCryptoEngineStatefulEncryptionDecryption from protobuf object *otg.MacsecCryptoEngineStatefulEncryptionDecryption
	// and doesn't set defaults
	setMsg(*otg.MacsecCryptoEngineStatefulEncryptionDecryption) MacsecCryptoEngineStatefulEncryptionDecryption
	// provides marshal interface
	Marshal() marshalMacsecCryptoEngineStatefulEncryptionDecryption
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecCryptoEngineStatefulEncryptionDecryption
	// validate validates MacsecCryptoEngineStatefulEncryptionDecryption
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecCryptoEngineStatefulEncryptionDecryption, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// InitialPn returns MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn, set in MacsecCryptoEngineStatefulEncryptionDecryption.
	// MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn is initial packet number(PN) configuration.
	InitialPn() MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn
	// SetInitialPn assigns MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn provided by user to MacsecCryptoEngineStatefulEncryptionDecryption.
	// MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn is initial packet number(PN) configuration.
	SetInitialPn(value MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn) MacsecCryptoEngineStatefulEncryptionDecryption
	// HasInitialPn checks if InitialPn has been set in MacsecCryptoEngineStatefulEncryptionDecryption
	HasInitialPn() bool
	// HardwareAcceleration returns MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration, set in MacsecCryptoEngineStatefulEncryptionDecryption.
	// MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration is hardware acceleration configuration for offloading MACsec processing to hardware.
	HardwareAcceleration() MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration
	// SetHardwareAcceleration assigns MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration provided by user to MacsecCryptoEngineStatefulEncryptionDecryption.
	// MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration is hardware acceleration configuration for offloading MACsec processing to hardware.
	SetHardwareAcceleration(value MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) MacsecCryptoEngineStatefulEncryptionDecryption
	// HasHardwareAcceleration checks if HardwareAcceleration has been set in MacsecCryptoEngineStatefulEncryptionDecryption
	HasHardwareAcceleration() bool
	setNil()
}

// description is TBD
// InitialPn returns a MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn
func (obj *macsecCryptoEngineStatefulEncryptionDecryption) InitialPn() MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn {
	if obj.obj.InitialPn == nil {
		obj.obj.InitialPn = NewMacsecCryptoEngineStatefulEncryptionDecryptionInitialPn().msg()
	}
	if obj.initialPnHolder == nil {
		obj.initialPnHolder = &macsecCryptoEngineStatefulEncryptionDecryptionInitialPn{obj: obj.obj.InitialPn}
	}
	return obj.initialPnHolder
}

// description is TBD
// InitialPn returns a MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn
func (obj *macsecCryptoEngineStatefulEncryptionDecryption) HasInitialPn() bool {
	return obj.obj.InitialPn != nil
}

// description is TBD
// SetInitialPn sets the MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn value in the MacsecCryptoEngineStatefulEncryptionDecryption object
func (obj *macsecCryptoEngineStatefulEncryptionDecryption) SetInitialPn(value MacsecCryptoEngineStatefulEncryptionDecryptionInitialPn) MacsecCryptoEngineStatefulEncryptionDecryption {

	obj.initialPnHolder = nil
	obj.obj.InitialPn = value.msg()

	return obj
}

// description is TBD
// HardwareAcceleration returns a MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration
func (obj *macsecCryptoEngineStatefulEncryptionDecryption) HardwareAcceleration() MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration {
	if obj.obj.HardwareAcceleration == nil {
		obj.obj.HardwareAcceleration = NewMacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration().msg()
	}
	if obj.hardwareAccelerationHolder == nil {
		obj.hardwareAccelerationHolder = &macsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration{obj: obj.obj.HardwareAcceleration}
	}
	return obj.hardwareAccelerationHolder
}

// description is TBD
// HardwareAcceleration returns a MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration
func (obj *macsecCryptoEngineStatefulEncryptionDecryption) HasHardwareAcceleration() bool {
	return obj.obj.HardwareAcceleration != nil
}

// description is TBD
// SetHardwareAcceleration sets the MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration value in the MacsecCryptoEngineStatefulEncryptionDecryption object
func (obj *macsecCryptoEngineStatefulEncryptionDecryption) SetHardwareAcceleration(value MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAcceleration) MacsecCryptoEngineStatefulEncryptionDecryption {

	obj.hardwareAccelerationHolder = nil
	obj.obj.HardwareAcceleration = value.msg()

	return obj
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryption) validateObj(vObj *validation, set_default bool) {
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

func (obj *macsecCryptoEngineStatefulEncryptionDecryption) setDefault() {

}
