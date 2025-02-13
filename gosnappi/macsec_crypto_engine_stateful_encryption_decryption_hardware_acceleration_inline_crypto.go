package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto *****
type macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto struct {
	validation
	obj          *otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	marshaller   marshalMacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	unMarshaller unMarshalMacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
}

func NewMacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto() MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {
	obj := macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto{obj: &otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) msg() *otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {
	return obj.obj
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) setMsg(msg *otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto struct {
	obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
}

type marshalMacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto interface {
	// ToProto marshals MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto to protobuf object *otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	ToProto() (*otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto, error)
	// ToPbText marshals MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto struct {
	obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
}

type unMarshalMacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto interface {
	// FromProto unmarshals MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto from protobuf object *otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	FromProto(msg *otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) (MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto, error)
	// FromPbText unmarshals MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto from JSON text
	FromJson(value string) error
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) Marshal() marshalMacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) Unmarshal() unMarshalMacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) ToProto() (*otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) FromProto(msg *otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) (MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) ToPbText() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) FromPbText(value string) error {
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

func (m *marshalmacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) ToYaml() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) FromYaml(value string) error {
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

func (m *marshalmacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) ToJson() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) FromJson(value string) error {
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

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) Clone() (MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto()
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

// MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto is inline cryto engine settings.
type MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto interface {
	Validation
	// msg marshals MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto to protobuf object *otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// and doesn't set defaults
	msg() *otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// setMsg unmarshals MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto from protobuf object *otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// and doesn't set defaults
	setMsg(*otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// provides marshal interface
	Marshal() marshalMacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// validate validates MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RxSectagOffset returns uint32, set in MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto.
	RxSectagOffset() uint32
	// SetRxSectagOffset assigns uint32 provided by user to MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	SetRxSectagOffset(value uint32) MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// HasRxSectagOffset checks if RxSectagOffset has been set in MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	HasRxSectagOffset() bool
	// TypeOfCa returns MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum, set in MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	TypeOfCa() MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum
	// SetTypeOfCa assigns MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum provided by user to MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	SetTypeOfCa(value MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum) MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// HasTypeOfCa checks if TypeOfCa has been set in MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	HasTypeOfCa() bool
	// MaxCaCount returns uint32, set in MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto.
	MaxCaCount() uint32
	// SetMaxCaCount assigns uint32 provided by user to MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	SetMaxCaCount(value uint32) MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// HasMaxCaCount checks if MaxCaCount has been set in MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	HasMaxCaCount() bool
	// MaxDutTxScPerCa returns uint32, set in MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto.
	MaxDutTxScPerCa() uint32
	// SetMaxDutTxScPerCa assigns uint32 provided by user to MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	SetMaxDutTxScPerCa(value uint32) MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// HasMaxDutTxScPerCa checks if MaxDutTxScPerCa has been set in MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	HasMaxDutTxScPerCa() bool
	// MaxDevicePerCa returns uint32, set in MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto.
	MaxDevicePerCa() uint32
	// SetMaxDevicePerCa assigns uint32 provided by user to MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	SetMaxDevicePerCa(value uint32) MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// HasMaxDevicePerCa checks if MaxDevicePerCa has been set in MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	HasMaxDevicePerCa() bool
	// RxScIdentifyingField returns MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum, set in MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	RxScIdentifyingField() MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum
	// SetRxScIdentifyingField assigns MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum provided by user to MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	SetRxScIdentifyingField(value MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum) MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// HasRxScIdentifyingField checks if RxScIdentifyingField has been set in MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	HasRxScIdentifyingField() bool
}

// Offset of Rx secTAG from the first byte in packet.
// RxSectagOffset returns a uint32
func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) RxSectagOffset() uint32 {

	return *obj.obj.RxSectagOffset

}

// Offset of Rx secTAG from the first byte in packet.
// RxSectagOffset returns a uint32
func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) HasRxSectagOffset() bool {
	return obj.obj.RxSectagOffset != nil
}

// Offset of Rx secTAG from the first byte in packet.
// SetRxSectagOffset sets the uint32 value in the MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto object
func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) SetRxSectagOffset(value uint32) MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {

	obj.obj.RxSectagOffset = &value
	return obj
}

type MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum string

// Enum of TypeOfCa on MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
var MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCa = struct {
	PAIRWISE_CA           MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum
	GROUP_CA_SINGLE_DUT   MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum
	GROUP_CA_MULTIPE_DUTS MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum
}{
	PAIRWISE_CA:           MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum("pairwise_ca"),
	GROUP_CA_SINGLE_DUT:   MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum("group_ca_single_dut"),
	GROUP_CA_MULTIPE_DUTS: MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum("group_ca_multipe_duts"),
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) TypeOfCa() MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum {
	return MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum(obj.obj.TypeOfCa.Enum().String())
}

// Type of CA.
// TypeOfCa returns a string
func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) HasTypeOfCa() bool {
	return obj.obj.TypeOfCa != nil
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) SetTypeOfCa(value MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum) MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {
	intValue, ok := otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto_TypeOfCa_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto_TypeOfCa_Enum(intValue)
	obj.obj.TypeOfCa = &enumValue

	return obj
}

// Maximum CA count.
// MaxCaCount returns a uint32
func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) MaxCaCount() uint32 {

	return *obj.obj.MaxCaCount

}

// Maximum CA count.
// MaxCaCount returns a uint32
func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) HasMaxCaCount() bool {
	return obj.obj.MaxCaCount != nil
}

// Maximum CA count.
// SetMaxCaCount sets the uint32 value in the MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto object
func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) SetMaxCaCount(value uint32) MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {

	obj.obj.MaxCaCount = &value
	return obj
}

// Maximum DUT Tx SC per CA.
// MaxDutTxScPerCa returns a uint32
func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) MaxDutTxScPerCa() uint32 {

	return *obj.obj.MaxDutTxScPerCa

}

// Maximum DUT Tx SC per CA.
// MaxDutTxScPerCa returns a uint32
func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) HasMaxDutTxScPerCa() bool {
	return obj.obj.MaxDutTxScPerCa != nil
}

// Maximum DUT Tx SC per CA.
// SetMaxDutTxScPerCa sets the uint32 value in the MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto object
func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) SetMaxDutTxScPerCa(value uint32) MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {

	obj.obj.MaxDutTxScPerCa = &value
	return obj
}

// Maximum devices per CA.
// MaxDevicePerCa returns a uint32
func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) MaxDevicePerCa() uint32 {

	return *obj.obj.MaxDevicePerCa

}

// Maximum devices per CA.
// MaxDevicePerCa returns a uint32
func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) HasMaxDevicePerCa() bool {
	return obj.obj.MaxDevicePerCa != nil
}

// Maximum devices per CA.
// SetMaxDevicePerCa sets the uint32 value in the MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto object
func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) SetMaxDevicePerCa(value uint32) MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {

	obj.obj.MaxDevicePerCa = &value
	return obj
}

type MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum string

// Enum of RxScIdentifyingField on MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
var MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingField = struct {
	SOURCE_MAC   MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum
	SCI_SYTEM_ID MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum
	SCI_PORT_ID  MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum
}{
	SOURCE_MAC:   MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum("source_mac"),
	SCI_SYTEM_ID: MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum("sci_sytem_id"),
	SCI_PORT_ID:  MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum("sci_port_id"),
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) RxScIdentifyingField() MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum {
	return MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum(obj.obj.RxScIdentifyingField.Enum().String())
}

// Rx SC identifying_field.
// RxScIdentifyingField returns a string
func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) HasRxScIdentifyingField() bool {
	return obj.obj.RxScIdentifyingField != nil
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) SetRxScIdentifyingField(value MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum) MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {
	intValue, ok := otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto_RxScIdentifyingField_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto_RxScIdentifyingField_Enum(intValue)
	obj.obj.RxScIdentifyingField = &enumValue

	return obj
}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *macsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) setDefault() {
	if obj.obj.RxSectagOffset == nil {
		obj.SetRxSectagOffset(12)
	}
	if obj.obj.MaxCaCount == nil {
		obj.SetMaxCaCount(256)
	}
	if obj.obj.MaxDutTxScPerCa == nil {
		obj.SetMaxDutTxScPerCa(1)
	}
	if obj.obj.MaxDevicePerCa == nil {
		obj.SetMaxDevicePerCa(256)
	}
	if obj.obj.RxScIdentifyingField == nil {
		obj.SetRxScIdentifyingField(MacsecCryptoEngineStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingField.SOURCE_MAC)

	}

}
