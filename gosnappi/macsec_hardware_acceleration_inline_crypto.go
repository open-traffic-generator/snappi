package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecHardwareAccelerationInlineCrypto *****
type macsecHardwareAccelerationInlineCrypto struct {
	validation
	obj            *otg.MacsecHardwareAccelerationInlineCrypto
	marshaller     marshalMacsecHardwareAccelerationInlineCrypto
	unMarshaller   unMarshalMacsecHardwareAccelerationInlineCrypto
	typeOfCaHolder MacsecHardwareAccelerationInlineCryptoTypeOfCa
}

func NewMacsecHardwareAccelerationInlineCrypto() MacsecHardwareAccelerationInlineCrypto {
	obj := macsecHardwareAccelerationInlineCrypto{obj: &otg.MacsecHardwareAccelerationInlineCrypto{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecHardwareAccelerationInlineCrypto) msg() *otg.MacsecHardwareAccelerationInlineCrypto {
	return obj.obj
}

func (obj *macsecHardwareAccelerationInlineCrypto) setMsg(msg *otg.MacsecHardwareAccelerationInlineCrypto) MacsecHardwareAccelerationInlineCrypto {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecHardwareAccelerationInlineCrypto struct {
	obj *macsecHardwareAccelerationInlineCrypto
}

type marshalMacsecHardwareAccelerationInlineCrypto interface {
	// ToProto marshals MacsecHardwareAccelerationInlineCrypto to protobuf object *otg.MacsecHardwareAccelerationInlineCrypto
	ToProto() (*otg.MacsecHardwareAccelerationInlineCrypto, error)
	// ToPbText marshals MacsecHardwareAccelerationInlineCrypto to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecHardwareAccelerationInlineCrypto to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecHardwareAccelerationInlineCrypto to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecHardwareAccelerationInlineCrypto struct {
	obj *macsecHardwareAccelerationInlineCrypto
}

type unMarshalMacsecHardwareAccelerationInlineCrypto interface {
	// FromProto unmarshals MacsecHardwareAccelerationInlineCrypto from protobuf object *otg.MacsecHardwareAccelerationInlineCrypto
	FromProto(msg *otg.MacsecHardwareAccelerationInlineCrypto) (MacsecHardwareAccelerationInlineCrypto, error)
	// FromPbText unmarshals MacsecHardwareAccelerationInlineCrypto from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecHardwareAccelerationInlineCrypto from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecHardwareAccelerationInlineCrypto from JSON text
	FromJson(value string) error
}

func (obj *macsecHardwareAccelerationInlineCrypto) Marshal() marshalMacsecHardwareAccelerationInlineCrypto {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecHardwareAccelerationInlineCrypto{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecHardwareAccelerationInlineCrypto) Unmarshal() unMarshalMacsecHardwareAccelerationInlineCrypto {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecHardwareAccelerationInlineCrypto{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecHardwareAccelerationInlineCrypto) ToProto() (*otg.MacsecHardwareAccelerationInlineCrypto, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecHardwareAccelerationInlineCrypto) FromProto(msg *otg.MacsecHardwareAccelerationInlineCrypto) (MacsecHardwareAccelerationInlineCrypto, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecHardwareAccelerationInlineCrypto) ToPbText() (string, error) {
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

func (m *unMarshalmacsecHardwareAccelerationInlineCrypto) FromPbText(value string) error {
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

func (m *marshalmacsecHardwareAccelerationInlineCrypto) ToYaml() (string, error) {
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

func (m *unMarshalmacsecHardwareAccelerationInlineCrypto) FromYaml(value string) error {
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

func (m *marshalmacsecHardwareAccelerationInlineCrypto) ToJson() (string, error) {
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

func (m *unMarshalmacsecHardwareAccelerationInlineCrypto) FromJson(value string) error {
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

func (obj *macsecHardwareAccelerationInlineCrypto) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecHardwareAccelerationInlineCrypto) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecHardwareAccelerationInlineCrypto) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecHardwareAccelerationInlineCrypto) Clone() (MacsecHardwareAccelerationInlineCrypto, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecHardwareAccelerationInlineCrypto()
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

func (obj *macsecHardwareAccelerationInlineCrypto) setNil() {
	obj.typeOfCaHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecHardwareAccelerationInlineCrypto is inline cryto engine configuration. Encryption/ decryption are offloaded to hardware. Also dynamic fields e.g. packet number(PN) and integrity check value(ICV) are updated in MACsec header on transmit.
type MacsecHardwareAccelerationInlineCrypto interface {
	Validation
	// msg marshals MacsecHardwareAccelerationInlineCrypto to protobuf object *otg.MacsecHardwareAccelerationInlineCrypto
	// and doesn't set defaults
	msg() *otg.MacsecHardwareAccelerationInlineCrypto
	// setMsg unmarshals MacsecHardwareAccelerationInlineCrypto from protobuf object *otg.MacsecHardwareAccelerationInlineCrypto
	// and doesn't set defaults
	setMsg(*otg.MacsecHardwareAccelerationInlineCrypto) MacsecHardwareAccelerationInlineCrypto
	// provides marshal interface
	Marshal() marshalMacsecHardwareAccelerationInlineCrypto
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecHardwareAccelerationInlineCrypto
	// validate validates MacsecHardwareAccelerationInlineCrypto
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecHardwareAccelerationInlineCrypto, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RxSectagOffset returns uint32, set in MacsecHardwareAccelerationInlineCrypto.
	RxSectagOffset() uint32
	// SetRxSectagOffset assigns uint32 provided by user to MacsecHardwareAccelerationInlineCrypto
	SetRxSectagOffset(value uint32) MacsecHardwareAccelerationInlineCrypto
	// HasRxSectagOffset checks if RxSectagOffset has been set in MacsecHardwareAccelerationInlineCrypto
	HasRxSectagOffset() bool
	// RxScIdentifyingField returns MacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum, set in MacsecHardwareAccelerationInlineCrypto
	RxScIdentifyingField() MacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum
	// SetRxScIdentifyingField assigns MacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum provided by user to MacsecHardwareAccelerationInlineCrypto
	SetRxScIdentifyingField(value MacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum) MacsecHardwareAccelerationInlineCrypto
	// HasRxScIdentifyingField checks if RxScIdentifyingField has been set in MacsecHardwareAccelerationInlineCrypto
	HasRxScIdentifyingField() bool
	// TypeOfCa returns MacsecHardwareAccelerationInlineCryptoTypeOfCa, set in MacsecHardwareAccelerationInlineCrypto.
	// MacsecHardwareAccelerationInlineCryptoTypeOfCa is type of connectivity association(CA).
	TypeOfCa() MacsecHardwareAccelerationInlineCryptoTypeOfCa
	// SetTypeOfCa assigns MacsecHardwareAccelerationInlineCryptoTypeOfCa provided by user to MacsecHardwareAccelerationInlineCrypto.
	// MacsecHardwareAccelerationInlineCryptoTypeOfCa is type of connectivity association(CA).
	SetTypeOfCa(value MacsecHardwareAccelerationInlineCryptoTypeOfCa) MacsecHardwareAccelerationInlineCrypto
	// HasTypeOfCa checks if TypeOfCa has been set in MacsecHardwareAccelerationInlineCrypto
	HasTypeOfCa() bool
	// MaxCaCount returns MacsecHardwareAccelerationInlineCryptoMaxCaCountEnum, set in MacsecHardwareAccelerationInlineCrypto
	MaxCaCount() MacsecHardwareAccelerationInlineCryptoMaxCaCountEnum
	// SetMaxCaCount assigns MacsecHardwareAccelerationInlineCryptoMaxCaCountEnum provided by user to MacsecHardwareAccelerationInlineCrypto
	SetMaxCaCount(value MacsecHardwareAccelerationInlineCryptoMaxCaCountEnum) MacsecHardwareAccelerationInlineCrypto
	// HasMaxCaCount checks if MaxCaCount has been set in MacsecHardwareAccelerationInlineCrypto
	HasMaxCaCount() bool
	setNil()
}

// Offset of Rx secTAG from the first byte in packet.
// RxSectagOffset returns a uint32
func (obj *macsecHardwareAccelerationInlineCrypto) RxSectagOffset() uint32 {

	return *obj.obj.RxSectagOffset

}

// Offset of Rx secTAG from the first byte in packet.
// RxSectagOffset returns a uint32
func (obj *macsecHardwareAccelerationInlineCrypto) HasRxSectagOffset() bool {
	return obj.obj.RxSectagOffset != nil
}

// Offset of Rx secTAG from the first byte in packet.
// SetRxSectagOffset sets the uint32 value in the MacsecHardwareAccelerationInlineCrypto object
func (obj *macsecHardwareAccelerationInlineCrypto) SetRxSectagOffset(value uint32) MacsecHardwareAccelerationInlineCrypto {

	obj.obj.RxSectagOffset = &value
	return obj
}

type MacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum string

// Enum of RxScIdentifyingField on MacsecHardwareAccelerationInlineCrypto
var MacsecHardwareAccelerationInlineCryptoRxScIdentifyingField = struct {
	SOURCE_MAC   MacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum
	SCI_SYTEM_ID MacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum
	SCI_PORT_ID  MacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum
}{
	SOURCE_MAC:   MacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum("source_mac"),
	SCI_SYTEM_ID: MacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum("sci_sytem_id"),
	SCI_PORT_ID:  MacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum("sci_port_id"),
}

func (obj *macsecHardwareAccelerationInlineCrypto) RxScIdentifyingField() MacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum {
	return MacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum(obj.obj.RxScIdentifyingField.Enum().String())
}

// The field based on which secure channel(SC) will be identified by the receiving port. Supported fields are:- - 1) source MAC - identify SC based on source MAC field. 2) SCI system ID - identify SC bbased on SCI system ID field. 3) SCI port ID - identify based on SCI port ID field.
// RxScIdentifyingField returns a string
func (obj *macsecHardwareAccelerationInlineCrypto) HasRxScIdentifyingField() bool {
	return obj.obj.RxScIdentifyingField != nil
}

func (obj *macsecHardwareAccelerationInlineCrypto) SetRxScIdentifyingField(value MacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum) MacsecHardwareAccelerationInlineCrypto {
	intValue, ok := otg.MacsecHardwareAccelerationInlineCrypto_RxScIdentifyingField_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecHardwareAccelerationInlineCryptoRxScIdentifyingFieldEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecHardwareAccelerationInlineCrypto_RxScIdentifyingField_Enum(intValue)
	obj.obj.RxScIdentifyingField = &enumValue

	return obj
}

// description is TBD
// TypeOfCa returns a MacsecHardwareAccelerationInlineCryptoTypeOfCa
func (obj *macsecHardwareAccelerationInlineCrypto) TypeOfCa() MacsecHardwareAccelerationInlineCryptoTypeOfCa {
	if obj.obj.TypeOfCa == nil {
		obj.obj.TypeOfCa = NewMacsecHardwareAccelerationInlineCryptoTypeOfCa().msg()
	}
	if obj.typeOfCaHolder == nil {
		obj.typeOfCaHolder = &macsecHardwareAccelerationInlineCryptoTypeOfCa{obj: obj.obj.TypeOfCa}
	}
	return obj.typeOfCaHolder
}

// description is TBD
// TypeOfCa returns a MacsecHardwareAccelerationInlineCryptoTypeOfCa
func (obj *macsecHardwareAccelerationInlineCrypto) HasTypeOfCa() bool {
	return obj.obj.TypeOfCa != nil
}

// description is TBD
// SetTypeOfCa sets the MacsecHardwareAccelerationInlineCryptoTypeOfCa value in the MacsecHardwareAccelerationInlineCrypto object
func (obj *macsecHardwareAccelerationInlineCrypto) SetTypeOfCa(value MacsecHardwareAccelerationInlineCryptoTypeOfCa) MacsecHardwareAccelerationInlineCrypto {

	obj.typeOfCaHolder = nil
	obj.obj.TypeOfCa = value.msg()

	return obj
}

type MacsecHardwareAccelerationInlineCryptoMaxCaCountEnum string

// Enum of MaxCaCount on MacsecHardwareAccelerationInlineCrypto
var MacsecHardwareAccelerationInlineCryptoMaxCaCount = struct {
	ONE              MacsecHardwareAccelerationInlineCryptoMaxCaCountEnum
	TWO              MacsecHardwareAccelerationInlineCryptoMaxCaCountEnum
	FOUR             MacsecHardwareAccelerationInlineCryptoMaxCaCountEnum
	EIGHT            MacsecHardwareAccelerationInlineCryptoMaxCaCountEnum
	SIXTEEN          MacsecHardwareAccelerationInlineCryptoMaxCaCountEnum
	THIRTY_TWO       MacsecHardwareAccelerationInlineCryptoMaxCaCountEnum
	SIXTY_FOUR       MacsecHardwareAccelerationInlineCryptoMaxCaCountEnum
	ONE_TWENTY_EIGHT MacsecHardwareAccelerationInlineCryptoMaxCaCountEnum
	TWO_FIFTY_SIX    MacsecHardwareAccelerationInlineCryptoMaxCaCountEnum
}{
	ONE:              MacsecHardwareAccelerationInlineCryptoMaxCaCountEnum("one"),
	TWO:              MacsecHardwareAccelerationInlineCryptoMaxCaCountEnum("two"),
	FOUR:             MacsecHardwareAccelerationInlineCryptoMaxCaCountEnum("four"),
	EIGHT:            MacsecHardwareAccelerationInlineCryptoMaxCaCountEnum("eight"),
	SIXTEEN:          MacsecHardwareAccelerationInlineCryptoMaxCaCountEnum("sixteen"),
	THIRTY_TWO:       MacsecHardwareAccelerationInlineCryptoMaxCaCountEnum("thirty_two"),
	SIXTY_FOUR:       MacsecHardwareAccelerationInlineCryptoMaxCaCountEnum("sixty_four"),
	ONE_TWENTY_EIGHT: MacsecHardwareAccelerationInlineCryptoMaxCaCountEnum("one_twenty_eight"),
	TWO_FIFTY_SIX:    MacsecHardwareAccelerationInlineCryptoMaxCaCountEnum("two_fifty_six"),
}

func (obj *macsecHardwareAccelerationInlineCrypto) MaxCaCount() MacsecHardwareAccelerationInlineCryptoMaxCaCountEnum {
	return MacsecHardwareAccelerationInlineCryptoMaxCaCountEnum(obj.obj.MaxCaCount.Enum().String())
}

// The maximum number of CAs configured on the port. The maximum count supported per port is 256 for Pair-wise CA, each CA having one MACsec device.
// MaxCaCount returns a string
func (obj *macsecHardwareAccelerationInlineCrypto) HasMaxCaCount() bool {
	return obj.obj.MaxCaCount != nil
}

func (obj *macsecHardwareAccelerationInlineCrypto) SetMaxCaCount(value MacsecHardwareAccelerationInlineCryptoMaxCaCountEnum) MacsecHardwareAccelerationInlineCrypto {
	intValue, ok := otg.MacsecHardwareAccelerationInlineCrypto_MaxCaCount_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MacsecHardwareAccelerationInlineCryptoMaxCaCountEnum", string(value)))
		return obj
	}
	enumValue := otg.MacsecHardwareAccelerationInlineCrypto_MaxCaCount_Enum(intValue)
	obj.obj.MaxCaCount = &enumValue

	return obj
}

func (obj *macsecHardwareAccelerationInlineCrypto) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.TypeOfCa != nil {

		obj.TypeOfCa().validateObj(vObj, set_default)
	}

}

func (obj *macsecHardwareAccelerationInlineCrypto) setDefault() {
	if obj.obj.RxSectagOffset == nil {
		obj.SetRxSectagOffset(12)
	}
	if obj.obj.RxScIdentifyingField == nil {
		obj.SetRxScIdentifyingField(MacsecHardwareAccelerationInlineCryptoRxScIdentifyingField.SOURCE_MAC)

	}
	if obj.obj.MaxCaCount == nil {
		obj.SetMaxCaCount(MacsecHardwareAccelerationInlineCryptoMaxCaCount.ONE)

	}

}
