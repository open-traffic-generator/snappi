package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MkaBasic *****
type mkaBasic struct {
	validation
	obj                         *otg.MkaBasic
	marshaller                  marshalMkaBasic
	unMarshaller                unMarshalMkaBasic
	keySourceHolder             MkaBasicKeySource
	supportedCipherSuitesHolder MkaBasicSupportedCipherSuites
	rekeyModeHolder             MkaBasicRekeyMode
}

func NewMkaBasic() MkaBasic {
	obj := mkaBasic{obj: &otg.MkaBasic{}}
	obj.setDefault()
	return &obj
}

func (obj *mkaBasic) msg() *otg.MkaBasic {
	return obj.obj
}

func (obj *mkaBasic) setMsg(msg *otg.MkaBasic) MkaBasic {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmkaBasic struct {
	obj *mkaBasic
}

type marshalMkaBasic interface {
	// ToProto marshals MkaBasic to protobuf object *otg.MkaBasic
	ToProto() (*otg.MkaBasic, error)
	// ToPbText marshals MkaBasic to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MkaBasic to YAML text
	ToYaml() (string, error)
	// ToJson marshals MkaBasic to JSON text
	ToJson() (string, error)
}

type unMarshalmkaBasic struct {
	obj *mkaBasic
}

type unMarshalMkaBasic interface {
	// FromProto unmarshals MkaBasic from protobuf object *otg.MkaBasic
	FromProto(msg *otg.MkaBasic) (MkaBasic, error)
	// FromPbText unmarshals MkaBasic from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MkaBasic from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MkaBasic from JSON text
	FromJson(value string) error
}

func (obj *mkaBasic) Marshal() marshalMkaBasic {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmkaBasic{obj: obj}
	}
	return obj.marshaller
}

func (obj *mkaBasic) Unmarshal() unMarshalMkaBasic {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmkaBasic{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmkaBasic) ToProto() (*otg.MkaBasic, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmkaBasic) FromProto(msg *otg.MkaBasic) (MkaBasic, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmkaBasic) ToPbText() (string, error) {
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

func (m *unMarshalmkaBasic) FromPbText(value string) error {
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

func (m *marshalmkaBasic) ToYaml() (string, error) {
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

func (m *unMarshalmkaBasic) FromYaml(value string) error {
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

func (m *marshalmkaBasic) ToJson() (string, error) {
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

func (m *unMarshalmkaBasic) FromJson(value string) error {
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

func (obj *mkaBasic) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *mkaBasic) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *mkaBasic) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *mkaBasic) Clone() (MkaBasic, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMkaBasic()
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

func (obj *mkaBasic) setNil() {
	obj.keySourceHolder = nil
	obj.supportedCipherSuitesHolder = nil
	obj.rekeyModeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MkaBasic is a container of basic properties for a KaY.
type MkaBasic interface {
	Validation
	// msg marshals MkaBasic to protobuf object *otg.MkaBasic
	// and doesn't set defaults
	msg() *otg.MkaBasic
	// setMsg unmarshals MkaBasic from protobuf object *otg.MkaBasic
	// and doesn't set defaults
	setMsg(*otg.MkaBasic) MkaBasic
	// provides marshal interface
	Marshal() marshalMkaBasic
	// provides unmarshal interface
	Unmarshal() unMarshalMkaBasic
	// validate validates MkaBasic
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MkaBasic, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// KeyDerivationFunction returns MkaBasicKeyDerivationFunctionEnum, set in MkaBasic
	KeyDerivationFunction() MkaBasicKeyDerivationFunctionEnum
	// SetKeyDerivationFunction assigns MkaBasicKeyDerivationFunctionEnum provided by user to MkaBasic
	SetKeyDerivationFunction(value MkaBasicKeyDerivationFunctionEnum) MkaBasic
	// HasKeyDerivationFunction checks if KeyDerivationFunction has been set in MkaBasic
	HasKeyDerivationFunction() bool
	// KeySource returns MkaBasicKeySource, set in MkaBasic.
	// MkaBasicKeySource is the container for static key settings.
	KeySource() MkaBasicKeySource
	// SetKeySource assigns MkaBasicKeySource provided by user to MkaBasic.
	// MkaBasicKeySource is the container for static key settings.
	SetKeySource(value MkaBasicKeySource) MkaBasic
	// ActorPriority returns string, set in MkaBasic.
	ActorPriority() string
	// SetActorPriority assigns string provided by user to MkaBasic
	SetActorPriority(value string) MkaBasic
	// HasActorPriority checks if ActorPriority has been set in MkaBasic
	HasActorPriority() bool
	// MacsecDesired returns bool, set in MkaBasic.
	MacsecDesired() bool
	// SetMacsecDesired assigns bool provided by user to MkaBasic
	SetMacsecDesired(value bool) MkaBasic
	// HasMacsecDesired checks if MacsecDesired has been set in MkaBasic
	HasMacsecDesired() bool
	// MacsecCapability returns MkaBasicMacsecCapabilityEnum, set in MkaBasic
	MacsecCapability() MkaBasicMacsecCapabilityEnum
	// SetMacsecCapability assigns MkaBasicMacsecCapabilityEnum provided by user to MkaBasic
	SetMacsecCapability(value MkaBasicMacsecCapabilityEnum) MkaBasic
	// HasMacsecCapability checks if MacsecCapability has been set in MkaBasic
	HasMacsecCapability() bool
	// SupportedCipherSuites returns MkaBasicSupportedCipherSuites, set in MkaBasic.
	// MkaBasicSupportedCipherSuites is the container for supported cipher suites.
	SupportedCipherSuites() MkaBasicSupportedCipherSuites
	// SetSupportedCipherSuites assigns MkaBasicSupportedCipherSuites provided by user to MkaBasic.
	// MkaBasicSupportedCipherSuites is the container for supported cipher suites.
	SetSupportedCipherSuites(value MkaBasicSupportedCipherSuites) MkaBasic
	// HasSupportedCipherSuites checks if SupportedCipherSuites has been set in MkaBasic
	HasSupportedCipherSuites() bool
	// MkaVersion returns uint32, set in MkaBasic.
	MkaVersion() uint32
	// SetMkaVersion assigns uint32 provided by user to MkaBasic
	SetMkaVersion(value uint32) MkaBasic
	// HasMkaVersion checks if MkaVersion has been set in MkaBasic
	HasMkaVersion() bool
	// MkaHelloTime returns uint32, set in MkaBasic.
	MkaHelloTime() uint32
	// SetMkaHelloTime assigns uint32 provided by user to MkaBasic
	SetMkaHelloTime(value uint32) MkaBasic
	// HasMkaHelloTime checks if MkaHelloTime has been set in MkaBasic
	HasMkaHelloTime() bool
	// MkaLifeTime returns uint32, set in MkaBasic.
	MkaLifeTime() uint32
	// SetMkaLifeTime assigns uint32 provided by user to MkaBasic
	SetMkaLifeTime(value uint32) MkaBasic
	// HasMkaLifeTime checks if MkaLifeTime has been set in MkaBasic
	HasMkaLifeTime() bool
	// SendIcvIndicatiorInMkpdu returns bool, set in MkaBasic.
	SendIcvIndicatiorInMkpdu() bool
	// SetSendIcvIndicatiorInMkpdu assigns bool provided by user to MkaBasic
	SetSendIcvIndicatiorInMkpdu(value bool) MkaBasic
	// HasSendIcvIndicatiorInMkpdu checks if SendIcvIndicatiorInMkpdu has been set in MkaBasic
	HasSendIcvIndicatiorInMkpdu() bool
	// DelayProtect returns bool, set in MkaBasic.
	DelayProtect() bool
	// SetDelayProtect assigns bool provided by user to MkaBasic
	SetDelayProtect(value bool) MkaBasic
	// HasDelayProtect checks if DelayProtect has been set in MkaBasic
	HasDelayProtect() bool
	// RekeyMode returns MkaBasicRekeyMode, set in MkaBasic.
	// MkaBasicRekeyMode is rekey mode.
	RekeyMode() MkaBasicRekeyMode
	// SetRekeyMode assigns MkaBasicRekeyMode provided by user to MkaBasic.
	// MkaBasicRekeyMode is rekey mode.
	SetRekeyMode(value MkaBasicRekeyMode) MkaBasic
	// HasRekeyMode checks if RekeyMode has been set in MkaBasic
	HasRekeyMode() bool
	setNil()
}

type MkaBasicKeyDerivationFunctionEnum string

// Enum of KeyDerivationFunction on MkaBasic
var MkaBasicKeyDerivationFunction = struct {
	AES_CMAC_128 MkaBasicKeyDerivationFunctionEnum
	AES_CMAC_256 MkaBasicKeyDerivationFunctionEnum
}{
	AES_CMAC_128: MkaBasicKeyDerivationFunctionEnum("aes_cmac_128"),
	AES_CMAC_256: MkaBasicKeyDerivationFunctionEnum("aes_cmac_256"),
}

func (obj *mkaBasic) KeyDerivationFunction() MkaBasicKeyDerivationFunctionEnum {
	return MkaBasicKeyDerivationFunctionEnum(obj.obj.KeyDerivationFunction.Enum().String())
}

// Key Derivation Function.
// KeyDerivationFunction returns a string
func (obj *mkaBasic) HasKeyDerivationFunction() bool {
	return obj.obj.KeyDerivationFunction != nil
}

func (obj *mkaBasic) SetKeyDerivationFunction(value MkaBasicKeyDerivationFunctionEnum) MkaBasic {
	intValue, ok := otg.MkaBasic_KeyDerivationFunction_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MkaBasicKeyDerivationFunctionEnum", string(value)))
		return obj
	}
	enumValue := otg.MkaBasic_KeyDerivationFunction_Enum(intValue)
	obj.obj.KeyDerivationFunction = &enumValue

	return obj
}

// Key source.
// KeySource returns a MkaBasicKeySource
func (obj *mkaBasic) KeySource() MkaBasicKeySource {
	if obj.obj.KeySource == nil {
		obj.obj.KeySource = NewMkaBasicKeySource().msg()
	}
	if obj.keySourceHolder == nil {
		obj.keySourceHolder = &mkaBasicKeySource{obj: obj.obj.KeySource}
	}
	return obj.keySourceHolder
}

// Key source.
// SetKeySource sets the MkaBasicKeySource value in the MkaBasic object
func (obj *mkaBasic) SetKeySource(value MkaBasicKeySource) MkaBasic {

	obj.keySourceHolder = nil
	obj.obj.KeySource = value.msg()

	return obj
}

// Actor priority.
// ActorPriority returns a string
func (obj *mkaBasic) ActorPriority() string {

	return *obj.obj.ActorPriority

}

// Actor priority.
// ActorPriority returns a string
func (obj *mkaBasic) HasActorPriority() bool {
	return obj.obj.ActorPriority != nil
}

// Actor priority.
// SetActorPriority sets the string value in the MkaBasic object
func (obj *mkaBasic) SetActorPriority(value string) MkaBasic {

	obj.obj.ActorPriority = &value
	return obj
}

// Determines whether MACsec is desired or not. It is advertised in periodic Hellos.
// MacsecDesired returns a bool
func (obj *mkaBasic) MacsecDesired() bool {

	return *obj.obj.MacsecDesired

}

// Determines whether MACsec is desired or not. It is advertised in periodic Hellos.
// MacsecDesired returns a bool
func (obj *mkaBasic) HasMacsecDesired() bool {
	return obj.obj.MacsecDesired != nil
}

// Determines whether MACsec is desired or not. It is advertised in periodic Hellos.
// SetMacsecDesired sets the bool value in the MkaBasic object
func (obj *mkaBasic) SetMacsecDesired(value bool) MkaBasic {

	obj.obj.MacsecDesired = &value
	return obj
}

type MkaBasicMacsecCapabilityEnum string

// Enum of MacsecCapability on MkaBasic
var MkaBasicMacsecCapability = struct {
	MACSEC_NOT_IMPLEMENTED                          MkaBasicMacsecCapabilityEnum
	MACSEC_INTEGRITY_WITHOUT_CONFIDENTIALITY        MkaBasicMacsecCapabilityEnum
	MACSEC_INTEGRITY_WITH_NO_CONFIDENTIALITY_OFFSET MkaBasicMacsecCapabilityEnum
	MACSEC_INTEGRITY_WITH_CONFIDENTIALITY_OFFSET    MkaBasicMacsecCapabilityEnum
}{
	MACSEC_NOT_IMPLEMENTED:                          MkaBasicMacsecCapabilityEnum("macsec_not_implemented"),
	MACSEC_INTEGRITY_WITHOUT_CONFIDENTIALITY:        MkaBasicMacsecCapabilityEnum("macsec_integrity_without_confidentiality"),
	MACSEC_INTEGRITY_WITH_NO_CONFIDENTIALITY_OFFSET: MkaBasicMacsecCapabilityEnum("macsec_integrity_with_no_confidentiality_offset"),
	MACSEC_INTEGRITY_WITH_CONFIDENTIALITY_OFFSET:    MkaBasicMacsecCapabilityEnum("macsec_integrity_with_confidentiality_offset"),
}

func (obj *mkaBasic) MacsecCapability() MkaBasicMacsecCapabilityEnum {
	return MkaBasicMacsecCapabilityEnum(obj.obj.MacsecCapability.Enum().String())
}

// MACSec Capability.
// MacsecCapability returns a string
func (obj *mkaBasic) HasMacsecCapability() bool {
	return obj.obj.MacsecCapability != nil
}

func (obj *mkaBasic) SetMacsecCapability(value MkaBasicMacsecCapabilityEnum) MkaBasic {
	intValue, ok := otg.MkaBasic_MacsecCapability_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MkaBasicMacsecCapabilityEnum", string(value)))
		return obj
	}
	enumValue := otg.MkaBasic_MacsecCapability_Enum(intValue)
	obj.obj.MacsecCapability = &enumValue

	return obj
}

// Supported Cipher Suites.
// SupportedCipherSuites returns a MkaBasicSupportedCipherSuites
func (obj *mkaBasic) SupportedCipherSuites() MkaBasicSupportedCipherSuites {
	if obj.obj.SupportedCipherSuites == nil {
		obj.obj.SupportedCipherSuites = NewMkaBasicSupportedCipherSuites().msg()
	}
	if obj.supportedCipherSuitesHolder == nil {
		obj.supportedCipherSuitesHolder = &mkaBasicSupportedCipherSuites{obj: obj.obj.SupportedCipherSuites}
	}
	return obj.supportedCipherSuitesHolder
}

// Supported Cipher Suites.
// SupportedCipherSuites returns a MkaBasicSupportedCipherSuites
func (obj *mkaBasic) HasSupportedCipherSuites() bool {
	return obj.obj.SupportedCipherSuites != nil
}

// Supported Cipher Suites.
// SetSupportedCipherSuites sets the MkaBasicSupportedCipherSuites value in the MkaBasic object
func (obj *mkaBasic) SetSupportedCipherSuites(value MkaBasicSupportedCipherSuites) MkaBasic {

	obj.supportedCipherSuitesHolder = nil
	obj.obj.SupportedCipherSuites = value.msg()

	return obj
}

// MKA Version.
// MkaVersion returns a uint32
func (obj *mkaBasic) MkaVersion() uint32 {

	return *obj.obj.MkaVersion

}

// MKA Version.
// MkaVersion returns a uint32
func (obj *mkaBasic) HasMkaVersion() bool {
	return obj.obj.MkaVersion != nil
}

// MKA Version.
// SetMkaVersion sets the uint32 value in the MkaBasic object
func (obj *mkaBasic) SetMkaVersion(value uint32) MkaBasic {

	obj.obj.MkaVersion = &value
	return obj
}

// MKA Hello Time (msec).
// MkaHelloTime returns a uint32
func (obj *mkaBasic) MkaHelloTime() uint32 {

	return *obj.obj.MkaHelloTime

}

// MKA Hello Time (msec).
// MkaHelloTime returns a uint32
func (obj *mkaBasic) HasMkaHelloTime() bool {
	return obj.obj.MkaHelloTime != nil
}

// MKA Hello Time (msec).
// SetMkaHelloTime sets the uint32 value in the MkaBasic object
func (obj *mkaBasic) SetMkaHelloTime(value uint32) MkaBasic {

	obj.obj.MkaHelloTime = &value
	return obj
}

// MKA Life Time (sec).
// MkaLifeTime returns a uint32
func (obj *mkaBasic) MkaLifeTime() uint32 {

	return *obj.obj.MkaLifeTime

}

// MKA Life Time (sec).
// MkaLifeTime returns a uint32
func (obj *mkaBasic) HasMkaLifeTime() bool {
	return obj.obj.MkaLifeTime != nil
}

// MKA Life Time (sec).
// SetMkaLifeTime sets the uint32 value in the MkaBasic object
func (obj *mkaBasic) SetMkaLifeTime(value uint32) MkaBasic {

	obj.obj.MkaLifeTime = &value
	return obj
}

// Send ICV Indicator in MKPDU.
// SendIcvIndicatiorInMkpdu returns a bool
func (obj *mkaBasic) SendIcvIndicatiorInMkpdu() bool {

	return *obj.obj.SendIcvIndicatiorInMkpdu

}

// Send ICV Indicator in MKPDU.
// SendIcvIndicatiorInMkpdu returns a bool
func (obj *mkaBasic) HasSendIcvIndicatiorInMkpdu() bool {
	return obj.obj.SendIcvIndicatiorInMkpdu != nil
}

// Send ICV Indicator in MKPDU.
// SetSendIcvIndicatiorInMkpdu sets the bool value in the MkaBasic object
func (obj *mkaBasic) SetSendIcvIndicatiorInMkpdu(value bool) MkaBasic {

	obj.obj.SendIcvIndicatiorInMkpdu = &value
	return obj
}

// Delay Protect.
// DelayProtect returns a bool
func (obj *mkaBasic) DelayProtect() bool {

	return *obj.obj.DelayProtect

}

// Delay Protect.
// DelayProtect returns a bool
func (obj *mkaBasic) HasDelayProtect() bool {
	return obj.obj.DelayProtect != nil
}

// Delay Protect.
// SetDelayProtect sets the bool value in the MkaBasic object
func (obj *mkaBasic) SetDelayProtect(value bool) MkaBasic {

	obj.obj.DelayProtect = &value
	return obj
}

// Rekey Mode.
// RekeyMode returns a MkaBasicRekeyMode
func (obj *mkaBasic) RekeyMode() MkaBasicRekeyMode {
	if obj.obj.RekeyMode == nil {
		obj.obj.RekeyMode = NewMkaBasicRekeyMode().msg()
	}
	if obj.rekeyModeHolder == nil {
		obj.rekeyModeHolder = &mkaBasicRekeyMode{obj: obj.obj.RekeyMode}
	}
	return obj.rekeyModeHolder
}

// Rekey Mode.
// RekeyMode returns a MkaBasicRekeyMode
func (obj *mkaBasic) HasRekeyMode() bool {
	return obj.obj.RekeyMode != nil
}

// Rekey Mode.
// SetRekeyMode sets the MkaBasicRekeyMode value in the MkaBasic object
func (obj *mkaBasic) SetRekeyMode(value MkaBasicRekeyMode) MkaBasic {

	obj.rekeyModeHolder = nil
	obj.obj.RekeyMode = value.msg()

	return obj
}

func (obj *mkaBasic) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// KeySource is required
	if obj.obj.KeySource == nil {
		vObj.validationErrors = append(vObj.validationErrors, "KeySource is required field on interface MkaBasic")
	}

	if obj.obj.KeySource != nil {

		obj.KeySource().validateObj(vObj, set_default)
	}

	if obj.obj.ActorPriority != nil {

		if len(*obj.obj.ActorPriority) < 2 || len(*obj.obj.ActorPriority) > 2 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"2 <= length of MkaBasic.ActorPriority <= 2 but Got %d",
					len(*obj.obj.ActorPriority)))
		}

	}

	if obj.obj.SupportedCipherSuites != nil {

		obj.SupportedCipherSuites().validateObj(vObj, set_default)
	}

	if obj.obj.RekeyMode != nil {

		obj.RekeyMode().validateObj(vObj, set_default)
	}

}

func (obj *mkaBasic) setDefault() {
	if obj.obj.KeyDerivationFunction == nil {
		obj.SetKeyDerivationFunction(MkaBasicKeyDerivationFunction.AES_CMAC_128)

	}
	if obj.obj.ActorPriority == nil {
		obj.SetActorPriority("70")
	}
	if obj.obj.MacsecDesired == nil {
		obj.SetMacsecDesired(true)
	}
	if obj.obj.MacsecCapability == nil {
		obj.SetMacsecCapability(MkaBasicMacsecCapability.MACSEC_INTEGRITY_WITH_CONFIDENTIALITY_OFFSET)

	}
	if obj.obj.MkaVersion == nil {
		obj.SetMkaVersion(3)
	}
	if obj.obj.MkaHelloTime == nil {
		obj.SetMkaHelloTime(2000)
	}
	if obj.obj.MkaLifeTime == nil {
		obj.SetMkaLifeTime(6)
	}
	if obj.obj.SendIcvIndicatiorInMkpdu == nil {
		obj.SetSendIcvIndicatiorInMkpdu(true)
	}
	if obj.obj.DelayProtect == nil {
		obj.SetDelayProtect(true)
	}

}
