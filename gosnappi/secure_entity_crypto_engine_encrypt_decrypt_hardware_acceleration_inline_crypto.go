package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto *****
type secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto struct {
	validation
	obj          *otg.SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	marshaller   marshalSecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	unMarshaller unMarshalSecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
}

func NewSecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto() SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto {
	obj := secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto{obj: &otg.SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto{}}
	obj.setDefault()
	return &obj
}

func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) msg() *otg.SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto {
	return obj.obj
}

func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) setMsg(msg *otg.SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto struct {
	obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
}

type marshalSecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto interface {
	// ToProto marshals SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto to protobuf object *otg.SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	ToProto() (*otg.SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto, error)
	// ToPbText marshals SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto to YAML text
	ToYaml() (string, error)
	// ToJson marshals SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto to JSON text
	ToJson() (string, error)
}

type unMarshalsecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto struct {
	obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
}

type unMarshalSecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto interface {
	// FromProto unmarshals SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto from protobuf object *otg.SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	FromProto(msg *otg.SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) (SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto, error)
	// FromPbText unmarshals SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto from JSON text
	FromJson(value string) error
}

func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) Marshal() marshalSecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto{obj: obj}
	}
	return obj.marshaller
}

func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) Unmarshal() unMarshalSecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) ToProto() (*otg.SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) FromProto(msg *otg.SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) (SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) ToPbText() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) FromPbText(value string) error {
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

func (m *marshalsecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) ToYaml() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) FromYaml(value string) error {
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

func (m *marshalsecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) ToJson() (string, error) {
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

func (m *unMarshalsecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) FromJson(value string) error {
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

func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) Clone() (SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto()
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

// SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto is inline crypto engine configuration. Encryption/ decryption are offloaded to hardware. Also dynamic fields e.g. packet number(PN) and integrity check value(ICV) are updated in MACsec header on transmit.
type SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto interface {
	Validation
	// msg marshals SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto to protobuf object *otg.SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	// and doesn't set defaults
	msg() *otg.SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	// setMsg unmarshals SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto from protobuf object *otg.SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	// and doesn't set defaults
	setMsg(*otg.SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	// provides marshal interface
	Marshal() marshalSecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	// provides unmarshal interface
	Unmarshal() unMarshalSecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	// validate validates SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RxSectagOffset returns uint32, set in SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto.
	RxSectagOffset() uint32
	// SetRxSectagOffset assigns uint32 provided by user to SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	SetRxSectagOffset(value uint32) SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	// HasRxSectagOffset checks if RxSectagOffset has been set in SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	HasRxSectagOffset() bool
	// TypeOfCa returns SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoTypeOfCaEnum, set in SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	TypeOfCa() SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoTypeOfCaEnum
	// SetTypeOfCa assigns SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoTypeOfCaEnum provided by user to SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	SetTypeOfCa(value SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoTypeOfCaEnum) SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	// HasTypeOfCa checks if TypeOfCa has been set in SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	HasTypeOfCa() bool
	// MaxCaCount returns uint32, set in SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto.
	MaxCaCount() uint32
	// SetMaxCaCount assigns uint32 provided by user to SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	SetMaxCaCount(value uint32) SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	// HasMaxCaCount checks if MaxCaCount has been set in SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	HasMaxCaCount() bool
	// MaxDutTxScPerCa returns uint32, set in SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto.
	MaxDutTxScPerCa() uint32
	// SetMaxDutTxScPerCa assigns uint32 provided by user to SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	SetMaxDutTxScPerCa(value uint32) SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	// HasMaxDutTxScPerCa checks if MaxDutTxScPerCa has been set in SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	HasMaxDutTxScPerCa() bool
	// MaxDevicePerCa returns uint32, set in SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto.
	MaxDevicePerCa() uint32
	// SetMaxDevicePerCa assigns uint32 provided by user to SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	SetMaxDevicePerCa(value uint32) SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	// HasMaxDevicePerCa checks if MaxDevicePerCa has been set in SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	HasMaxDevicePerCa() bool
	// RxScIdentifyingField returns SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum, set in SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	RxScIdentifyingField() SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum
	// SetRxScIdentifyingField assigns SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum provided by user to SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	SetRxScIdentifyingField(value SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum) SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	// HasRxScIdentifyingField checks if RxScIdentifyingField has been set in SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
	HasRxScIdentifyingField() bool
}

// Offset of Rx secTAG from the first byte in packet.
// RxSectagOffset returns a uint32
func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) RxSectagOffset() uint32 {

	return *obj.obj.RxSectagOffset

}

// Offset of Rx secTAG from the first byte in packet.
// RxSectagOffset returns a uint32
func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) HasRxSectagOffset() bool {
	return obj.obj.RxSectagOffset != nil
}

// Offset of Rx secTAG from the first byte in packet.
// SetRxSectagOffset sets the uint32 value in the SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto object
func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) SetRxSectagOffset(value uint32) SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto {

	obj.obj.RxSectagOffset = &value
	return obj
}

type SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoTypeOfCaEnum string

// Enum of TypeOfCa on SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
var SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoTypeOfCa = struct {
	PAIRWISE_CA           SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoTypeOfCaEnum
	GROUP_CA_SINGLE_DUT   SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoTypeOfCaEnum
	GROUP_CA_MULTIPE_DUTS SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoTypeOfCaEnum
}{
	PAIRWISE_CA:           SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoTypeOfCaEnum("pairwise_ca"),
	GROUP_CA_SINGLE_DUT:   SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoTypeOfCaEnum("group_ca_single_dut"),
	GROUP_CA_MULTIPE_DUTS: SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoTypeOfCaEnum("group_ca_multipe_duts"),
}

func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) TypeOfCa() SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoTypeOfCaEnum {
	return SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoTypeOfCaEnum(obj.obj.TypeOfCa.Enum().String())
}

// Type of connectivity association(CA).
// TypeOfCa returns a string
func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) HasTypeOfCa() bool {
	return obj.obj.TypeOfCa != nil
}

func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) SetTypeOfCa(value SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoTypeOfCaEnum) SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto {
	intValue, ok := otg.SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto_TypeOfCa_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoTypeOfCaEnum", string(value)))
		return obj
	}
	enumValue := otg.SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto_TypeOfCa_Enum(intValue)
	obj.obj.TypeOfCa = &enumValue

	return obj
}

// The maximum number of CAs configured on the port. The maximum count supported per port is 256 for Pair-wise CA, each CA having one MACsec device.
// MaxCaCount returns a uint32
func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) MaxCaCount() uint32 {

	return *obj.obj.MaxCaCount

}

// The maximum number of CAs configured on the port. The maximum count supported per port is 256 for Pair-wise CA, each CA having one MACsec device.
// MaxCaCount returns a uint32
func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) HasMaxCaCount() bool {
	return obj.obj.MaxCaCount != nil
}

// The maximum number of CAs configured on the port. The maximum count supported per port is 256 for Pair-wise CA, each CA having one MACsec device.
// SetMaxCaCount sets the uint32 value in the SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto object
func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) SetMaxCaCount(value uint32) SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto {

	obj.obj.MaxCaCount = &value
	return obj
}

// The maximum number of DUT transmit SCs that can be supported per CA. The count should be number of Tx SCs supported by the DUT per CA, multiplied by number of DUTs in the CA in case of group CA with multiple DUTs scenario.
// MaxDutTxScPerCa returns a uint32
func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) MaxDutTxScPerCa() uint32 {

	return *obj.obj.MaxDutTxScPerCa

}

// The maximum number of DUT transmit SCs that can be supported per CA. The count should be number of Tx SCs supported by the DUT per CA, multiplied by number of DUTs in the CA in case of group CA with multiple DUTs scenario.
// MaxDutTxScPerCa returns a uint32
func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) HasMaxDutTxScPerCa() bool {
	return obj.obj.MaxDutTxScPerCa != nil
}

// The maximum number of DUT transmit SCs that can be supported per CA. The count should be number of Tx SCs supported by the DUT per CA, multiplied by number of DUTs in the CA in case of group CA with multiple DUTs scenario.
// SetMaxDutTxScPerCa sets the uint32 value in the SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto object
func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) SetMaxDutTxScPerCa(value uint32) SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto {

	obj.obj.MaxDutTxScPerCa = &value
	return obj
}

// The maximum number of MACsec devices at test port that can be supported on each CA. This number is calculated automatically based on the values configured for Max CA Count and Max DUT Tx SC Per CA. Number of MACsec devices at test port should be configured accordingly.
// MaxDevicePerCa returns a uint32
func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) MaxDevicePerCa() uint32 {

	return *obj.obj.MaxDevicePerCa

}

// The maximum number of MACsec devices at test port that can be supported on each CA. This number is calculated automatically based on the values configured for Max CA Count and Max DUT Tx SC Per CA. Number of MACsec devices at test port should be configured accordingly.
// MaxDevicePerCa returns a uint32
func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) HasMaxDevicePerCa() bool {
	return obj.obj.MaxDevicePerCa != nil
}

// The maximum number of MACsec devices at test port that can be supported on each CA. This number is calculated automatically based on the values configured for Max CA Count and Max DUT Tx SC Per CA. Number of MACsec devices at test port should be configured accordingly.
// SetMaxDevicePerCa sets the uint32 value in the SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto object
func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) SetMaxDevicePerCa(value uint32) SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto {

	obj.obj.MaxDevicePerCa = &value
	return obj
}

type SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum string

// Enum of RxScIdentifyingField on SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto
var SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoRxScIdentifyingField = struct {
	SOURCE_MAC   SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum
	SCI_SYTEM_ID SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum
	SCI_PORT_ID  SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum
}{
	SOURCE_MAC:   SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum("source_mac"),
	SCI_SYTEM_ID: SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum("sci_sytem_id"),
	SCI_PORT_ID:  SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum("sci_port_id"),
}

func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) RxScIdentifyingField() SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum {
	return SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum(obj.obj.RxScIdentifyingField.Enum().String())
}

// The field based on which secure channel(SC) will be identified by the receiving port. Supported fields are:- - 1) source MAC - identify SC based on source MAC field. 2) SCI system ID - identify SC bbased on SCI system ID field. 3) SCI port ID - identify based on SCI port ID field.
// RxScIdentifyingField returns a string
func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) HasRxScIdentifyingField() bool {
	return obj.obj.RxScIdentifyingField != nil
}

func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) SetRxScIdentifyingField(value SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum) SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto {
	intValue, ok := otg.SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto_RxScIdentifyingField_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum", string(value)))
		return obj
	}
	enumValue := otg.SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto_RxScIdentifyingField_Enum(intValue)
	obj.obj.RxScIdentifyingField = &enumValue

	return obj
}

func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MaxCaCount != nil {

		if *obj.obj.MaxCaCount < 1 || *obj.obj.MaxCaCount > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto.MaxCaCount <= 256 but Got %d", *obj.obj.MaxCaCount))
		}

	}

	if obj.obj.MaxDutTxScPerCa != nil {

		if *obj.obj.MaxDutTxScPerCa < 1 || *obj.obj.MaxDutTxScPerCa > 256 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto.MaxDutTxScPerCa <= 256 but Got %d", *obj.obj.MaxDutTxScPerCa))
		}

	}

	if obj.obj.MaxDevicePerCa != nil {

		if *obj.obj.MaxDevicePerCa < 1 || *obj.obj.MaxDevicePerCa > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto.MaxDevicePerCa <= 4294967295 but Got %d", *obj.obj.MaxDevicePerCa))
		}

	}

}

func (obj *secureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCrypto) setDefault() {
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
		obj.SetRxScIdentifyingField(SecureEntityCryptoEngineEncryptDecryptHardwareAccelerationInlineCryptoRxScIdentifyingField.SOURCE_MAC)

	}

}
