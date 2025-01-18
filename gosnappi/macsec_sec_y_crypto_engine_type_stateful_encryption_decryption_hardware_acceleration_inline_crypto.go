package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto *****
type macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto struct {
	validation
	obj          *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	marshaller   marshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	unMarshaller unMarshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
}

func NewMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto() MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {
	obj := macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto{obj: &otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) msg() *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {
	return obj.obj
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) setMsg(msg *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto struct {
	obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
}

type marshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto interface {
	// ToProto marshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto to protobuf object *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	ToProto() (*otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto, error)
	// ToPbText marshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto struct {
	obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
}

type unMarshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto interface {
	// FromProto unmarshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto from protobuf object *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	FromProto(msg *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) (MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto, error)
	// FromPbText unmarshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto from JSON text
	FromJson(value string) error
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) Marshal() marshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) Unmarshal() unMarshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) ToProto() (*otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) FromProto(msg *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) (MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) ToPbText() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) FromPbText(value string) error {
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

func (m *marshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) ToYaml() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) FromYaml(value string) error {
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

func (m *marshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) ToJson() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) FromJson(value string) error {
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

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) Clone() (MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto()
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

// MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto is inline cryto engine settings.
type MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto interface {
	Validation
	// msg marshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto to protobuf object *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// and doesn't set defaults
	msg() *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// setMsg unmarshals MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto from protobuf object *otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// and doesn't set defaults
	setMsg(*otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// provides marshal interface
	Marshal() marshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// validate validates MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RxSectagOffset returns int32, set in MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto.
	RxSectagOffset() int32
	// SetRxSectagOffset assigns int32 provided by user to MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	SetRxSectagOffset(value int32) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// HasRxSectagOffset checks if RxSectagOffset has been set in MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	HasRxSectagOffset() bool
	// TypeOfCa returns MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum, set in MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	TypeOfCa() MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum
	// SetTypeOfCa assigns MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum provided by user to MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	SetTypeOfCa(value MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// HasTypeOfCa checks if TypeOfCa has been set in MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	HasTypeOfCa() bool
	// MaxCaCount returns int32, set in MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto.
	MaxCaCount() int32
	// SetMaxCaCount assigns int32 provided by user to MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	SetMaxCaCount(value int32) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// HasMaxCaCount checks if MaxCaCount has been set in MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	HasMaxCaCount() bool
	// MaxDutTxScPerCa returns int32, set in MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto.
	MaxDutTxScPerCa() int32
	// SetMaxDutTxScPerCa assigns int32 provided by user to MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	SetMaxDutTxScPerCa(value int32) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// HasMaxDutTxScPerCa checks if MaxDutTxScPerCa has been set in MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	HasMaxDutTxScPerCa() bool
	// MaxDevicePerCa returns int32, set in MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto.
	MaxDevicePerCa() int32
	// SetMaxDevicePerCa assigns int32 provided by user to MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	SetMaxDevicePerCa(value int32) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// HasMaxDevicePerCa checks if MaxDevicePerCa has been set in MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	HasMaxDevicePerCa() bool
	// RxScIdentifyingField returns MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum, set in MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	RxScIdentifyingField() MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum
	// SetRxScIdentifyingField assigns MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum provided by user to MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	SetRxScIdentifyingField(value MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// HasRxScIdentifyingField checks if RxScIdentifyingField has been set in MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	HasRxScIdentifyingField() bool
}

// Offset of Rx secTAG from the first byte in packet.
// RxSectagOffset returns a int32
func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) RxSectagOffset() int32 {

	return *obj.obj.RxSectagOffset

}

// Offset of Rx secTAG from the first byte in packet.
// RxSectagOffset returns a int32
func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) HasRxSectagOffset() bool {
	return obj.obj.RxSectagOffset != nil
}

// Offset of Rx secTAG from the first byte in packet.
// SetRxSectagOffset sets the int32 value in the MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto object
func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) SetRxSectagOffset(value int32) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {

	obj.obj.RxSectagOffset = &value
	return obj
}

type MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum string

// Enum of TypeOfCa on MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
var MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCa = struct {
	PAIRWISE_CA           MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum
	GROUP_CA_SINGLE_DUT   MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum
	GROUP_CA_MULTIPE_DUTS MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum
}{
	PAIRWISE_CA:           MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum("pairwise_ca"),
	GROUP_CA_SINGLE_DUT:   MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum("group_ca_single_dut"),
	GROUP_CA_MULTIPE_DUTS: MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum("group_ca_multipe_duts"),
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) TypeOfCa() MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum {
	return MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum(obj.obj.TypeOfCa.Enum().String())
}

// Type of CA.
// TypeOfCa returns a string
func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) HasTypeOfCa() bool {
	return obj.obj.TypeOfCa != nil
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) SetTypeOfCa(value MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {
	intValue, ok := otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto_TypeOfCa_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto_TypeOfCa_Enum(intValue)
	obj.obj.TypeOfCa = &enumValue

	return obj
}

// Maximum CA count.
// MaxCaCount returns a int32
func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) MaxCaCount() int32 {

	return *obj.obj.MaxCaCount

}

// Maximum CA count.
// MaxCaCount returns a int32
func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) HasMaxCaCount() bool {
	return obj.obj.MaxCaCount != nil
}

// Maximum CA count.
// SetMaxCaCount sets the int32 value in the MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto object
func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) SetMaxCaCount(value int32) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {

	obj.obj.MaxCaCount = &value
	return obj
}

// Maximum DUT Tx SC per CA.
// MaxDutTxScPerCa returns a int32
func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) MaxDutTxScPerCa() int32 {

	return *obj.obj.MaxDutTxScPerCa

}

// Maximum DUT Tx SC per CA.
// MaxDutTxScPerCa returns a int32
func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) HasMaxDutTxScPerCa() bool {
	return obj.obj.MaxDutTxScPerCa != nil
}

// Maximum DUT Tx SC per CA.
// SetMaxDutTxScPerCa sets the int32 value in the MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto object
func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) SetMaxDutTxScPerCa(value int32) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {

	obj.obj.MaxDutTxScPerCa = &value
	return obj
}

// Maximum devices per CA.
// MaxDevicePerCa returns a int32
func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) MaxDevicePerCa() int32 {

	return *obj.obj.MaxDevicePerCa

}

// Maximum devices per CA.
// MaxDevicePerCa returns a int32
func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) HasMaxDevicePerCa() bool {
	return obj.obj.MaxDevicePerCa != nil
}

// Maximum devices per CA.
// SetMaxDevicePerCa sets the int32 value in the MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto object
func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) SetMaxDevicePerCa(value int32) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {

	obj.obj.MaxDevicePerCa = &value
	return obj
}

type MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum string

// Enum of RxScIdentifyingField on MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
var MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingField = struct {
	SOURCE_MAC   MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum
	SCI_SYTEM_ID MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum
	SCI_PORT_ID  MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum
}{
	SOURCE_MAC:   MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum("source_mac"),
	SCI_SYTEM_ID: MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum("sci_sytem_id"),
	SCI_PORT_ID:  MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum("sci_port_id"),
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) RxScIdentifyingField() MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum {
	return MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum(obj.obj.RxScIdentifyingField.Enum().String())
}

// Rx SC identifying_field.
// RxScIdentifyingField returns a string
func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) HasRxScIdentifyingField() bool {
	return obj.obj.RxScIdentifyingField != nil
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) SetRxScIdentifyingField(value MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum) MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {
	intValue, ok := otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto_RxScIdentifyingField_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto_RxScIdentifyingField_Enum(intValue)
	obj.obj.RxScIdentifyingField = &enumValue

	return obj
}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *macsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) setDefault() {
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
		obj.SetRxScIdentifyingField(MacsecSecYCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingField.SOURCE_MAC)

	}

}
