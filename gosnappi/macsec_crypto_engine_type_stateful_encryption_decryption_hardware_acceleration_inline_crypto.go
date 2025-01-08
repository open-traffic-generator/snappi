package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto *****
type macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto struct {
	validation
	obj          *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	marshaller   marshalMacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	unMarshaller unMarshalMacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
}

func NewMacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto() MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {
	obj := macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto{obj: &otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) msg() *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {
	return obj.obj
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) setMsg(msg *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto struct {
	obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
}

type marshalMacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto interface {
	// ToProto marshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto to protobuf object *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	ToProto() (*otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto, error)
	// ToPbText marshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto struct {
	obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
}

type unMarshalMacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto interface {
	// FromProto unmarshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto from protobuf object *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	FromProto(msg *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) (MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto, error)
	// FromPbText unmarshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto from JSON text
	FromJson(value string) error
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) Marshal() marshalMacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) Unmarshal() unMarshalMacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) ToProto() (*otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) FromProto(msg *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) (MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) ToPbText() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) FromPbText(value string) error {
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

func (m *marshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) ToYaml() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) FromYaml(value string) error {
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

func (m *marshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) ToJson() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) FromJson(value string) error {
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

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) Clone() (MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto()
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

// MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto is inline cryto engine settings.
type MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto interface {
	Validation
	// msg marshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto to protobuf object *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// and doesn't set defaults
	msg() *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// setMsg unmarshals MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto from protobuf object *otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// and doesn't set defaults
	setMsg(*otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// provides marshal interface
	Marshal() marshalMacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// validate validates MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RxSectagOffset returns int32, set in MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto.
	RxSectagOffset() int32
	// SetRxSectagOffset assigns int32 provided by user to MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	SetRxSectagOffset(value int32) MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// HasRxSectagOffset checks if RxSectagOffset has been set in MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	HasRxSectagOffset() bool
	// TypeOfCa returns MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum, set in MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	TypeOfCa() MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum
	// SetTypeOfCa assigns MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum provided by user to MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	SetTypeOfCa(value MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum) MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// HasTypeOfCa checks if TypeOfCa has been set in MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	HasTypeOfCa() bool
	// MaxCaCount returns int32, set in MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto.
	MaxCaCount() int32
	// SetMaxCaCount assigns int32 provided by user to MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	SetMaxCaCount(value int32) MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// HasMaxCaCount checks if MaxCaCount has been set in MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	HasMaxCaCount() bool
	// MaxDutTxScPerCa returns int32, set in MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto.
	MaxDutTxScPerCa() int32
	// SetMaxDutTxScPerCa assigns int32 provided by user to MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	SetMaxDutTxScPerCa(value int32) MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// HasMaxDutTxScPerCa checks if MaxDutTxScPerCa has been set in MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	HasMaxDutTxScPerCa() bool
	// MaxDevicePerCa returns int32, set in MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto.
	MaxDevicePerCa() int32
	// SetMaxDevicePerCa assigns int32 provided by user to MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	SetMaxDevicePerCa(value int32) MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// HasMaxDevicePerCa checks if MaxDevicePerCa has been set in MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	HasMaxDevicePerCa() bool
	// RxScIdentifyingField returns MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum, set in MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	RxScIdentifyingField() MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum
	// SetRxScIdentifyingField assigns MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum provided by user to MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	SetRxScIdentifyingField(value MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum) MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	// HasRxScIdentifyingField checks if RxScIdentifyingField has been set in MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
	HasRxScIdentifyingField() bool
}

// Offset of Rx secTAG from the first byte in packet.
// RxSectagOffset returns a int32
func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) RxSectagOffset() int32 {

	return *obj.obj.RxSectagOffset

}

// Offset of Rx secTAG from the first byte in packet.
// RxSectagOffset returns a int32
func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) HasRxSectagOffset() bool {
	return obj.obj.RxSectagOffset != nil
}

// Offset of Rx secTAG from the first byte in packet.
// SetRxSectagOffset sets the int32 value in the MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto object
func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) SetRxSectagOffset(value int32) MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {

	obj.obj.RxSectagOffset = &value
	return obj
}

type MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum string

// Enum of TypeOfCa on MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
var MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCa = struct {
	PAIRWISE_CA           MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum
	GROUP_CA_SINGLE_DUT   MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum
	GROUP_CA_MULTIPE_DUTS MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum
}{
	PAIRWISE_CA:           MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum("pairwise_ca"),
	GROUP_CA_SINGLE_DUT:   MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum("group_ca_single_dut"),
	GROUP_CA_MULTIPE_DUTS: MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum("group_ca_multipe_duts"),
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) TypeOfCa() MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum {
	return MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum(obj.obj.TypeOfCa.Enum().String())
}

// Type of CA.
// TypeOfCa returns a string
func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) HasTypeOfCa() bool {
	return obj.obj.TypeOfCa != nil
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) SetTypeOfCa(value MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum) MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {
	intValue, ok := otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto_TypeOfCa_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoTypeOfCaEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto_TypeOfCa_Enum(intValue)
	obj.obj.TypeOfCa = &enumValue

	return obj
}

// Maximum CA count.
// MaxCaCount returns a int32
func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) MaxCaCount() int32 {

	return *obj.obj.MaxCaCount

}

// Maximum CA count.
// MaxCaCount returns a int32
func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) HasMaxCaCount() bool {
	return obj.obj.MaxCaCount != nil
}

// Maximum CA count.
// SetMaxCaCount sets the int32 value in the MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto object
func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) SetMaxCaCount(value int32) MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {

	obj.obj.MaxCaCount = &value
	return obj
}

// Maximum DUT Tx SC per CA.
// MaxDutTxScPerCa returns a int32
func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) MaxDutTxScPerCa() int32 {

	return *obj.obj.MaxDutTxScPerCa

}

// Maximum DUT Tx SC per CA.
// MaxDutTxScPerCa returns a int32
func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) HasMaxDutTxScPerCa() bool {
	return obj.obj.MaxDutTxScPerCa != nil
}

// Maximum DUT Tx SC per CA.
// SetMaxDutTxScPerCa sets the int32 value in the MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto object
func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) SetMaxDutTxScPerCa(value int32) MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {

	obj.obj.MaxDutTxScPerCa = &value
	return obj
}

// Maximum devices per CA.
// MaxDevicePerCa returns a int32
func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) MaxDevicePerCa() int32 {

	return *obj.obj.MaxDevicePerCa

}

// Maximum devices per CA.
// MaxDevicePerCa returns a int32
func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) HasMaxDevicePerCa() bool {
	return obj.obj.MaxDevicePerCa != nil
}

// Maximum devices per CA.
// SetMaxDevicePerCa sets the int32 value in the MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto object
func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) SetMaxDevicePerCa(value int32) MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {

	obj.obj.MaxDevicePerCa = &value
	return obj
}

type MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum string

// Enum of RxScIdentifyingField on MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto
var MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingField = struct {
	SOURCE_MAC   MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum
	SCI_SYTEM_ID MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum
	SCI_PORT_ID  MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum
}{
	SOURCE_MAC:   MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum("source_mac"),
	SCI_SYTEM_ID: MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum("sci_sytem_id"),
	SCI_PORT_ID:  MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum("sci_port_id"),
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) RxScIdentifyingField() MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum {
	return MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum(obj.obj.RxScIdentifyingField.Enum().String())
}

// Rx SC identifying_field.
// RxScIdentifyingField returns a string
func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) HasRxScIdentifyingField() bool {
	return obj.obj.RxScIdentifyingField != nil
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) SetRxScIdentifyingField(value MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum) MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto {
	intValue, ok := otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto_RxScIdentifyingField_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto_RxScIdentifyingField_Enum(intValue)
	obj.obj.RxScIdentifyingField = &enumValue

	return obj
}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *macsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCrypto) setDefault() {
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
		obj.SetRxScIdentifyingField(MacsecCryptoEngineTypeStatefulEncryptionDecryptionHardwareAccelerationInlineCryptoRxScIdentifyingField.SOURCE_MAC)

	}

}
