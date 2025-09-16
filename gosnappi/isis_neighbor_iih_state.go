package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisNeighborIIHState *****
type isisNeighborIIHState struct {
	validation
	obj                    *otg.IsisNeighborIIHState
	marshaller             marshalIsisNeighborIIHState
	unMarshaller           unMarshalIsisNeighborIIHState
	restartingStatusHolder IsisIIHNeighborRestartStatus
	tlvsHolder             IsisIIHNeighborTlvs
}

func NewIsisNeighborIIHState() IsisNeighborIIHState {
	obj := isisNeighborIIHState{obj: &otg.IsisNeighborIIHState{}}
	obj.setDefault()
	return &obj
}

func (obj *isisNeighborIIHState) msg() *otg.IsisNeighborIIHState {
	return obj.obj
}

func (obj *isisNeighborIIHState) setMsg(msg *otg.IsisNeighborIIHState) IsisNeighborIIHState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisNeighborIIHState struct {
	obj *isisNeighborIIHState
}

type marshalIsisNeighborIIHState interface {
	// ToProto marshals IsisNeighborIIHState to protobuf object *otg.IsisNeighborIIHState
	ToProto() (*otg.IsisNeighborIIHState, error)
	// ToPbText marshals IsisNeighborIIHState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisNeighborIIHState to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisNeighborIIHState to JSON text
	ToJson() (string, error)
}

type unMarshalisisNeighborIIHState struct {
	obj *isisNeighborIIHState
}

type unMarshalIsisNeighborIIHState interface {
	// FromProto unmarshals IsisNeighborIIHState from protobuf object *otg.IsisNeighborIIHState
	FromProto(msg *otg.IsisNeighborIIHState) (IsisNeighborIIHState, error)
	// FromPbText unmarshals IsisNeighborIIHState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisNeighborIIHState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisNeighborIIHState from JSON text
	FromJson(value string) error
}

func (obj *isisNeighborIIHState) Marshal() marshalIsisNeighborIIHState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisNeighborIIHState{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisNeighborIIHState) Unmarshal() unMarshalIsisNeighborIIHState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisNeighborIIHState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisNeighborIIHState) ToProto() (*otg.IsisNeighborIIHState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisNeighborIIHState) FromProto(msg *otg.IsisNeighborIIHState) (IsisNeighborIIHState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisNeighborIIHState) ToPbText() (string, error) {
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

func (m *unMarshalisisNeighborIIHState) FromPbText(value string) error {
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

func (m *marshalisisNeighborIIHState) ToYaml() (string, error) {
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

func (m *unMarshalisisNeighborIIHState) FromYaml(value string) error {
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

func (m *marshalisisNeighborIIHState) ToJson() (string, error) {
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

func (m *unMarshalisisNeighborIIHState) FromJson(value string) error {
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

func (obj *isisNeighborIIHState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisNeighborIIHState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisNeighborIIHState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisNeighborIIHState) Clone() (IsisNeighborIIHState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisNeighborIIHState()
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

func (obj *isisNeighborIIHState) setNil() {
	obj.restartingStatusHolder = nil
	obj.tlvsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisNeighborIIHState is information for neighbor adjacency State.
type IsisNeighborIIHState interface {
	Validation
	// msg marshals IsisNeighborIIHState to protobuf object *otg.IsisNeighborIIHState
	// and doesn't set defaults
	msg() *otg.IsisNeighborIIHState
	// setMsg unmarshals IsisNeighborIIHState from protobuf object *otg.IsisNeighborIIHState
	// and doesn't set defaults
	setMsg(*otg.IsisNeighborIIHState) IsisNeighborIIHState
	// provides marshal interface
	Marshal() marshalIsisNeighborIIHState
	// provides unmarshal interface
	Unmarshal() unMarshalIsisNeighborIIHState
	// validate validates IsisNeighborIIHState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisNeighborIIHState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LevelType returns IsisNeighborIIHStateLevelTypeEnum, set in IsisNeighborIIHState
	LevelType() IsisNeighborIIHStateLevelTypeEnum
	// SetLevelType assigns IsisNeighborIIHStateLevelTypeEnum provided by user to IsisNeighborIIHState
	SetLevelType(value IsisNeighborIIHStateLevelTypeEnum) IsisNeighborIIHState
	// HasLevelType checks if LevelType has been set in IsisNeighborIIHState
	HasLevelType() bool
	// HoldTimer returns uint32, set in IsisNeighborIIHState.
	HoldTimer() uint32
	// SetHoldTimer assigns uint32 provided by user to IsisNeighborIIHState
	SetHoldTimer(value uint32) IsisNeighborIIHState
	// HasHoldTimer checks if HoldTimer has been set in IsisNeighborIIHState
	HasHoldTimer() bool
	// RestartingStatus returns IsisIIHNeighborRestartStatus, set in IsisNeighborIIHState.
	// IsisIIHNeighborRestartStatus is this contains the Restarting/Starting/Running state of a neighbor router.
	RestartingStatus() IsisIIHNeighborRestartStatus
	// SetRestartingStatus assigns IsisIIHNeighborRestartStatus provided by user to IsisNeighborIIHState.
	// IsisIIHNeighborRestartStatus is this contains the Restarting/Starting/Running state of a neighbor router.
	SetRestartingStatus(value IsisIIHNeighborRestartStatus) IsisNeighborIIHState
	// HasRestartingStatus checks if RestartingStatus has been set in IsisNeighborIIHState
	HasRestartingStatus() bool
	// Tlvs returns IsisIIHNeighborTlvs, set in IsisNeighborIIHState.
	// IsisIIHNeighborTlvs is this contains the list of TLVs present in a IIH PDU received from a neighbor IS-IS router.
	Tlvs() IsisIIHNeighborTlvs
	// SetTlvs assigns IsisIIHNeighborTlvs provided by user to IsisNeighborIIHState.
	// IsisIIHNeighborTlvs is this contains the list of TLVs present in a IIH PDU received from a neighbor IS-IS router.
	SetTlvs(value IsisIIHNeighborTlvs) IsisNeighborIIHState
	// HasTlvs checks if Tlvs has been set in IsisNeighborIIHState
	HasTlvs() bool
	setNil()
}

type IsisNeighborIIHStateLevelTypeEnum string

// Enum of LevelType on IsisNeighborIIHState
var IsisNeighborIIHStateLevelType = struct {
	LEVEL_1   IsisNeighborIIHStateLevelTypeEnum
	LEVEL_2   IsisNeighborIIHStateLevelTypeEnum
	LEVEL_1_2 IsisNeighborIIHStateLevelTypeEnum
}{
	LEVEL_1:   IsisNeighborIIHStateLevelTypeEnum("level_1"),
	LEVEL_2:   IsisNeighborIIHStateLevelTypeEnum("level_2"),
	LEVEL_1_2: IsisNeighborIIHStateLevelTypeEnum("level_1_2"),
}

func (obj *isisNeighborIIHState) LevelType() IsisNeighborIIHStateLevelTypeEnum {
	return IsisNeighborIIHStateLevelTypeEnum(obj.obj.LevelType.Enum().String())
}

// This indicates whether Neighbor IS-IS router is participating in Level-1 (L1),
// Level-2 (L2) or both L1 and L2 domains on this interface.
// LevelType returns a string
func (obj *isisNeighborIIHState) HasLevelType() bool {
	return obj.obj.LevelType != nil
}

func (obj *isisNeighborIIHState) SetLevelType(value IsisNeighborIIHStateLevelTypeEnum) IsisNeighborIIHState {
	intValue, ok := otg.IsisNeighborIIHState_LevelType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisNeighborIIHStateLevelTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisNeighborIIHState_LevelType_Enum(intValue)
	obj.obj.LevelType = &enumValue

	return obj
}

// Hold timer received in the IIH PDU sent by the neighbor..
// HoldTimer returns a uint32
func (obj *isisNeighborIIHState) HoldTimer() uint32 {

	return *obj.obj.HoldTimer

}

// Hold timer received in the IIH PDU sent by the neighbor..
// HoldTimer returns a uint32
func (obj *isisNeighborIIHState) HasHoldTimer() bool {
	return obj.obj.HoldTimer != nil
}

// Hold timer received in the IIH PDU sent by the neighbor..
// SetHoldTimer sets the uint32 value in the IsisNeighborIIHState object
func (obj *isisNeighborIIHState) SetHoldTimer(value uint32) IsisNeighborIIHState {

	obj.obj.HoldTimer = &value
	return obj
}

// Reference to Restarting Information.
// RestartingStatus returns a IsisIIHNeighborRestartStatus
func (obj *isisNeighborIIHState) RestartingStatus() IsisIIHNeighborRestartStatus {
	if obj.obj.RestartingStatus == nil {
		obj.obj.RestartingStatus = NewIsisIIHNeighborRestartStatus().msg()
	}
	if obj.restartingStatusHolder == nil {
		obj.restartingStatusHolder = &isisIIHNeighborRestartStatus{obj: obj.obj.RestartingStatus}
	}
	return obj.restartingStatusHolder
}

// Reference to Restarting Information.
// RestartingStatus returns a IsisIIHNeighborRestartStatus
func (obj *isisNeighborIIHState) HasRestartingStatus() bool {
	return obj.obj.RestartingStatus != nil
}

// Reference to Restarting Information.
// SetRestartingStatus sets the IsisIIHNeighborRestartStatus value in the IsisNeighborIIHState object
func (obj *isisNeighborIIHState) SetRestartingStatus(value IsisIIHNeighborRestartStatus) IsisNeighborIIHState {

	obj.restartingStatusHolder = nil
	obj.obj.RestartingStatus = value.msg()

	return obj
}

// It refers to IIH PDU TLVs container.
// Tlvs returns a IsisIIHNeighborTlvs
func (obj *isisNeighborIIHState) Tlvs() IsisIIHNeighborTlvs {
	if obj.obj.Tlvs == nil {
		obj.obj.Tlvs = NewIsisIIHNeighborTlvs().msg()
	}
	if obj.tlvsHolder == nil {
		obj.tlvsHolder = &isisIIHNeighborTlvs{obj: obj.obj.Tlvs}
	}
	return obj.tlvsHolder
}

// It refers to IIH PDU TLVs container.
// Tlvs returns a IsisIIHNeighborTlvs
func (obj *isisNeighborIIHState) HasTlvs() bool {
	return obj.obj.Tlvs != nil
}

// It refers to IIH PDU TLVs container.
// SetTlvs sets the IsisIIHNeighborTlvs value in the IsisNeighborIIHState object
func (obj *isisNeighborIIHState) SetTlvs(value IsisIIHNeighborTlvs) IsisNeighborIIHState {

	obj.tlvsHolder = nil
	obj.obj.Tlvs = value.msg()

	return obj
}

func (obj *isisNeighborIIHState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RestartingStatus != nil {

		obj.RestartingStatus().validateObj(vObj, set_default)
	}

	if obj.obj.Tlvs != nil {

		obj.Tlvs().validateObj(vObj, set_default)
	}

}

func (obj *isisNeighborIIHState) setDefault() {

}
