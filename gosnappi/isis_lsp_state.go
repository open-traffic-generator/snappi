package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspState *****
type isisLspState struct {
	validation
	obj          *otg.IsisLspState
	marshaller   marshalIsisLspState
	unMarshaller unMarshalIsisLspState
	flagsHolder  IsisLspFlags
	tlvsHolder   IsisLspTlvs
}

func NewIsisLspState() IsisLspState {
	obj := isisLspState{obj: &otg.IsisLspState{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspState) msg() *otg.IsisLspState {
	return obj.obj
}

func (obj *isisLspState) setMsg(msg *otg.IsisLspState) IsisLspState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspState struct {
	obj *isisLspState
}

type marshalIsisLspState interface {
	// ToProto marshals IsisLspState to protobuf object *otg.IsisLspState
	ToProto() (*otg.IsisLspState, error)
	// ToPbText marshals IsisLspState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspState to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspState to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals IsisLspState to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalisisLspState struct {
	obj *isisLspState
}

type unMarshalIsisLspState interface {
	// FromProto unmarshals IsisLspState from protobuf object *otg.IsisLspState
	FromProto(msg *otg.IsisLspState) (IsisLspState, error)
	// FromPbText unmarshals IsisLspState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspState from JSON text
	FromJson(value string) error
}

func (obj *isisLspState) Marshal() marshalIsisLspState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspState{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspState) Unmarshal() unMarshalIsisLspState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspState) ToProto() (*otg.IsisLspState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspState) FromProto(msg *otg.IsisLspState) (IsisLspState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspState) ToPbText() (string, error) {
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

func (m *unMarshalisisLspState) FromPbText(value string) error {
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

func (m *marshalisisLspState) ToYaml() (string, error) {
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

func (m *unMarshalisisLspState) FromYaml(value string) error {
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

func (m *marshalisisLspState) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalisisLspState) ToJson() (string, error) {
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

func (m *unMarshalisisLspState) FromJson(value string) error {
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

func (obj *isisLspState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspState) Clone() (IsisLspState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspState()
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

func (obj *isisLspState) setNil() {
	obj.flagsHolder = nil
	obj.tlvsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisLspState is iSIS LSP.
type IsisLspState interface {
	Validation
	// msg marshals IsisLspState to protobuf object *otg.IsisLspState
	// and doesn't set defaults
	msg() *otg.IsisLspState
	// setMsg unmarshals IsisLspState from protobuf object *otg.IsisLspState
	// and doesn't set defaults
	setMsg(*otg.IsisLspState) IsisLspState
	// provides marshal interface
	Marshal() marshalIsisLspState
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspState
	// validate validates IsisLspState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LspId returns string, set in IsisLspState.
	LspId() string
	// SetLspId assigns string provided by user to IsisLspState
	SetLspId(value string) IsisLspState
	// PduType returns IsisLspStatePduTypeEnum, set in IsisLspState
	PduType() IsisLspStatePduTypeEnum
	// SetPduType assigns IsisLspStatePduTypeEnum provided by user to IsisLspState
	SetPduType(value IsisLspStatePduTypeEnum) IsisLspState
	// HasPduType checks if PduType has been set in IsisLspState
	HasPduType() bool
	// RemainingLifetime returns uint32, set in IsisLspState.
	RemainingLifetime() uint32
	// SetRemainingLifetime assigns uint32 provided by user to IsisLspState
	SetRemainingLifetime(value uint32) IsisLspState
	// HasRemainingLifetime checks if RemainingLifetime has been set in IsisLspState
	HasRemainingLifetime() bool
	// SequenceNumber returns uint64, set in IsisLspState.
	SequenceNumber() uint64
	// SetSequenceNumber assigns uint64 provided by user to IsisLspState
	SetSequenceNumber(value uint64) IsisLspState
	// HasSequenceNumber checks if SequenceNumber has been set in IsisLspState
	HasSequenceNumber() bool
	// PduLength returns uint32, set in IsisLspState.
	PduLength() uint32
	// SetPduLength assigns uint32 provided by user to IsisLspState
	SetPduLength(value uint32) IsisLspState
	// HasPduLength checks if PduLength has been set in IsisLspState
	HasPduLength() bool
	// Flags returns IsisLspFlags, set in IsisLspState.
	// IsisLspFlags is lSP Type flags.
	Flags() IsisLspFlags
	// SetFlags assigns IsisLspFlags provided by user to IsisLspState.
	// IsisLspFlags is lSP Type flags.
	SetFlags(value IsisLspFlags) IsisLspState
	// HasFlags checks if Flags has been set in IsisLspState
	HasFlags() bool
	// IsType returns uint32, set in IsisLspState.
	IsType() uint32
	// SetIsType assigns uint32 provided by user to IsisLspState
	SetIsType(value uint32) IsisLspState
	// HasIsType checks if IsType has been set in IsisLspState
	HasIsType() bool
	// Tlvs returns IsisLspTlvs, set in IsisLspState.
	// IsisLspTlvs is this contains the list of TLVs present in one LSP.
	Tlvs() IsisLspTlvs
	// SetTlvs assigns IsisLspTlvs provided by user to IsisLspState.
	// IsisLspTlvs is this contains the list of TLVs present in one LSP.
	SetTlvs(value IsisLspTlvs) IsisLspState
	// HasTlvs checks if Tlvs has been set in IsisLspState
	HasTlvs() bool
	setNil()
}

// LSP ID in the format, e.g. '640000000001-00-00'. LSP ID consists of the System ID of a neighbor, the Pseudonode ID, and the LSP number. The last two bytes represent Pseudonode ID and LSP number respectively. A pseudonode is a logical representation of the LAN which is generated by a Designated Intermediate System (DIS) on a LAN segment. If one LSP exceeds the maximum LSP size then it is sent in another LSP with the LSP number incremented by one. A router's learned LSP gets refreshed by 'remaining_lifetime'. Then the sequence number is incremented by 1.
// LspId returns a string
func (obj *isisLspState) LspId() string {

	return *obj.obj.LspId

}

// LSP ID in the format, e.g. '640000000001-00-00'. LSP ID consists of the System ID of a neighbor, the Pseudonode ID, and the LSP number. The last two bytes represent Pseudonode ID and LSP number respectively. A pseudonode is a logical representation of the LAN which is generated by a Designated Intermediate System (DIS) on a LAN segment. If one LSP exceeds the maximum LSP size then it is sent in another LSP with the LSP number incremented by one. A router's learned LSP gets refreshed by 'remaining_lifetime'. Then the sequence number is incremented by 1.
// SetLspId sets the string value in the IsisLspState object
func (obj *isisLspState) SetLspId(value string) IsisLspState {

	obj.obj.LspId = &value
	return obj
}

type IsisLspStatePduTypeEnum string

// Enum of PduType on IsisLspState
var IsisLspStatePduType = struct {
	LEVEL_1 IsisLspStatePduTypeEnum
	LEVEL_2 IsisLspStatePduTypeEnum
}{
	LEVEL_1: IsisLspStatePduTypeEnum("level_1"),
	LEVEL_2: IsisLspStatePduTypeEnum("level_2"),
}

func (obj *isisLspState) PduType() IsisLspStatePduTypeEnum {
	return IsisLspStatePduTypeEnum(obj.obj.PduType.Enum().String())
}

// Link State PDU type.
// PduType returns a string
func (obj *isisLspState) HasPduType() bool {
	return obj.obj.PduType != nil
}

func (obj *isisLspState) SetPduType(value IsisLspStatePduTypeEnum) IsisLspState {
	intValue, ok := otg.IsisLspState_PduType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisLspStatePduTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisLspState_PduType_Enum(intValue)
	obj.obj.PduType = &enumValue

	return obj
}

// Remaining lifetime in seconds before LSP expires.
// RemainingLifetime returns a uint32
func (obj *isisLspState) RemainingLifetime() uint32 {

	return *obj.obj.RemainingLifetime

}

// Remaining lifetime in seconds before LSP expires.
// RemainingLifetime returns a uint32
func (obj *isisLspState) HasRemainingLifetime() bool {
	return obj.obj.RemainingLifetime != nil
}

// Remaining lifetime in seconds before LSP expires.
// SetRemainingLifetime sets the uint32 value in the IsisLspState object
func (obj *isisLspState) SetRemainingLifetime(value uint32) IsisLspState {

	obj.obj.RemainingLifetime = &value
	return obj
}

// Sequence number of the LSP.
// SequenceNumber returns a uint64
func (obj *isisLspState) SequenceNumber() uint64 {

	return *obj.obj.SequenceNumber

}

// Sequence number of the LSP.
// SequenceNumber returns a uint64
func (obj *isisLspState) HasSequenceNumber() bool {
	return obj.obj.SequenceNumber != nil
}

// Sequence number of the LSP.
// SetSequenceNumber sets the uint64 value in the IsisLspState object
func (obj *isisLspState) SetSequenceNumber(value uint64) IsisLspState {

	obj.obj.SequenceNumber = &value
	return obj
}

// Total length of the LSP.
// PduLength returns a uint32
func (obj *isisLspState) PduLength() uint32 {

	return *obj.obj.PduLength

}

// Total length of the LSP.
// PduLength returns a uint32
func (obj *isisLspState) HasPduLength() bool {
	return obj.obj.PduLength != nil
}

// Total length of the LSP.
// SetPduLength sets the uint32 value in the IsisLspState object
func (obj *isisLspState) SetPduLength(value uint32) IsisLspState {

	obj.obj.PduLength = &value
	return obj
}

// LSP Type-Block flags.
// Flags returns a IsisLspFlags
func (obj *isisLspState) Flags() IsisLspFlags {
	if obj.obj.Flags == nil {
		obj.obj.Flags = NewIsisLspFlags().msg()
	}
	if obj.flagsHolder == nil {
		obj.flagsHolder = &isisLspFlags{obj: obj.obj.Flags}
	}
	return obj.flagsHolder
}

// LSP Type-Block flags.
// Flags returns a IsisLspFlags
func (obj *isisLspState) HasFlags() bool {
	return obj.obj.Flags != nil
}

// LSP Type-Block flags.
// SetFlags sets the IsisLspFlags value in the IsisLspState object
func (obj *isisLspState) SetFlags(value IsisLspFlags) IsisLspState {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// IS Type - bits 1 and 2 indicate the type of Intermediate System.
// 1 - ( i.e. bit 1 set) Level 1 Intermediate system.
// 2 - Unused value.
// 3 - (i.e. bits 1 and 2 set) Level 2 Intermediate system.
// IsType returns a uint32
func (obj *isisLspState) IsType() uint32 {

	return *obj.obj.IsType

}

// IS Type - bits 1 and 2 indicate the type of Intermediate System.
// 1 - ( i.e. bit 1 set) Level 1 Intermediate system.
// 2 - Unused value.
// 3 - (i.e. bits 1 and 2 set) Level 2 Intermediate system.
// IsType returns a uint32
func (obj *isisLspState) HasIsType() bool {
	return obj.obj.IsType != nil
}

// IS Type - bits 1 and 2 indicate the type of Intermediate System.
// 1 - ( i.e. bit 1 set) Level 1 Intermediate system.
// 2 - Unused value.
// 3 - (i.e. bits 1 and 2 set) Level 2 Intermediate system.
// SetIsType sets the uint32 value in the IsisLspState object
func (obj *isisLspState) SetIsType(value uint32) IsisLspState {

	obj.obj.IsType = &value
	return obj
}

// It refers to Link State PDU State TLVs container.
// Tlvs returns a IsisLspTlvs
func (obj *isisLspState) Tlvs() IsisLspTlvs {
	if obj.obj.Tlvs == nil {
		obj.obj.Tlvs = NewIsisLspTlvs().msg()
	}
	if obj.tlvsHolder == nil {
		obj.tlvsHolder = &isisLspTlvs{obj: obj.obj.Tlvs}
	}
	return obj.tlvsHolder
}

// It refers to Link State PDU State TLVs container.
// Tlvs returns a IsisLspTlvs
func (obj *isisLspState) HasTlvs() bool {
	return obj.obj.Tlvs != nil
}

// It refers to Link State PDU State TLVs container.
// SetTlvs sets the IsisLspTlvs value in the IsisLspState object
func (obj *isisLspState) SetTlvs(value IsisLspTlvs) IsisLspState {

	obj.tlvsHolder = nil
	obj.obj.Tlvs = value.msg()

	return obj
}

func (obj *isisLspState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// LspId is required
	if obj.obj.LspId == nil {
		vObj.validationErrors = append(vObj.validationErrors, "LspId is required field on interface IsisLspState")
	}

	if obj.obj.PduLength != nil {

		if *obj.obj.PduLength > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisLspState.PduLength <= 65535 but Got %d", *obj.obj.PduLength))
		}

	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(vObj, set_default)
	}

	if obj.obj.IsType != nil {

		if *obj.obj.IsType > 3 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisLspState.IsType <= 3 but Got %d", *obj.obj.IsType))
		}

	}

	if obj.obj.Tlvs != nil {

		obj.Tlvs().validateObj(vObj, set_default)
	}

}

func (obj *isisLspState) setDefault() {

}
