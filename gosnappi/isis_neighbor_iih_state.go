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
	obj                  *otg.IsisNeighborIIHState
	marshaller           marshalIsisNeighborIIHState
	unMarshaller         unMarshalIsisNeighborIIHState
	restartingInfoHolder IsisIIHRestartInfo
	tlvsHolder           IsisIIHTlvs
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
	obj.restartingInfoHolder = nil
	obj.tlvsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisNeighborIIHState is a ISIS Neigbor Information.
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
	// SystemId returns string, set in IsisNeighborIIHState.
	SystemId() string
	// SetSystemId assigns string provided by user to IsisNeighborIIHState
	SetSystemId(value string) IsisNeighborIIHState
	// HasSystemId checks if SystemId has been set in IsisNeighborIIHState
	HasSystemId() bool
	// NetworkType returns IsisNeighborIIHStateNetworkTypeEnum, set in IsisNeighborIIHState
	NetworkType() IsisNeighborIIHStateNetworkTypeEnum
	// SetNetworkType assigns IsisNeighborIIHStateNetworkTypeEnum provided by user to IsisNeighborIIHState
	SetNetworkType(value IsisNeighborIIHStateNetworkTypeEnum) IsisNeighborIIHState
	// HasNetworkType checks if NetworkType has been set in IsisNeighborIIHState
	HasNetworkType() bool
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
	// RestartingInfo returns IsisIIHRestartInfo, set in IsisNeighborIIHState.
	// IsisIIHRestartInfo is this contains the list of TLVs present in a IIH packet.
	RestartingInfo() IsisIIHRestartInfo
	// SetRestartingInfo assigns IsisIIHRestartInfo provided by user to IsisNeighborIIHState.
	// IsisIIHRestartInfo is this contains the list of TLVs present in a IIH packet.
	SetRestartingInfo(value IsisIIHRestartInfo) IsisNeighborIIHState
	// HasRestartingInfo checks if RestartingInfo has been set in IsisNeighborIIHState
	HasRestartingInfo() bool
	// Tlvs returns IsisIIHTlvs, set in IsisNeighborIIHState.
	// IsisIIHTlvs is this contains the list of TLVs present in a IIH packet.
	Tlvs() IsisIIHTlvs
	// SetTlvs assigns IsisIIHTlvs provided by user to IsisNeighborIIHState.
	// IsisIIHTlvs is this contains the list of TLVs present in a IIH packet.
	SetTlvs(value IsisIIHTlvs) IsisNeighborIIHState
	// HasTlvs checks if Tlvs has been set in IsisNeighborIIHState
	HasTlvs() bool
	setNil()
}

// System ID of a neighbor in the format, e.g. '640000000001'.
// SystemId returns a string
func (obj *isisNeighborIIHState) SystemId() string {

	return *obj.obj.SystemId

}

// System ID of a neighbor in the format, e.g. '640000000001'.
// SystemId returns a string
func (obj *isisNeighborIIHState) HasSystemId() bool {
	return obj.obj.SystemId != nil
}

// System ID of a neighbor in the format, e.g. '640000000001'.
// SetSystemId sets the string value in the IsisNeighborIIHState object
func (obj *isisNeighborIIHState) SetSystemId(value string) IsisNeighborIIHState {

	obj.obj.SystemId = &value
	return obj
}

type IsisNeighborIIHStateNetworkTypeEnum string

// Enum of NetworkType on IsisNeighborIIHState
var IsisNeighborIIHStateNetworkType = struct {
	BROADCAST      IsisNeighborIIHStateNetworkTypeEnum
	POINT_TO_POINT IsisNeighborIIHStateNetworkTypeEnum
}{
	BROADCAST:      IsisNeighborIIHStateNetworkTypeEnum("broadcast"),
	POINT_TO_POINT: IsisNeighborIIHStateNetworkTypeEnum("point_to_point"),
}

func (obj *isisNeighborIIHState) NetworkType() IsisNeighborIIHStateNetworkTypeEnum {
	return IsisNeighborIIHStateNetworkTypeEnum(obj.obj.NetworkType.Enum().String())
}

// The type of network link.
// NetworkType returns a string
func (obj *isisNeighborIIHState) HasNetworkType() bool {
	return obj.obj.NetworkType != nil
}

func (obj *isisNeighborIIHState) SetNetworkType(value IsisNeighborIIHStateNetworkTypeEnum) IsisNeighborIIHState {
	intValue, ok := otg.IsisNeighborIIHState_NetworkType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisNeighborIIHStateNetworkTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisNeighborIIHState_NetworkType_Enum(intValue)
	obj.obj.NetworkType = &enumValue

	return obj
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

// This indicates whether Neighbor ISIS router is participating in Level-1 (L1),
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

// Hold timer is in the IIH.
// HoldTimer returns a uint32
func (obj *isisNeighborIIHState) HoldTimer() uint32 {

	return *obj.obj.HoldTimer

}

// Hold timer is in the IIH.
// HoldTimer returns a uint32
func (obj *isisNeighborIIHState) HasHoldTimer() bool {
	return obj.obj.HoldTimer != nil
}

// Hold timer is in the IIH.
// SetHoldTimer sets the uint32 value in the IsisNeighborIIHState object
func (obj *isisNeighborIIHState) SetHoldTimer(value uint32) IsisNeighborIIHState {

	obj.obj.HoldTimer = &value
	return obj
}

// Reference to Restarting Information.
// RestartingInfo returns a IsisIIHRestartInfo
func (obj *isisNeighborIIHState) RestartingInfo() IsisIIHRestartInfo {
	if obj.obj.RestartingInfo == nil {
		obj.obj.RestartingInfo = NewIsisIIHRestartInfo().msg()
	}
	if obj.restartingInfoHolder == nil {
		obj.restartingInfoHolder = &isisIIHRestartInfo{obj: obj.obj.RestartingInfo}
	}
	return obj.restartingInfoHolder
}

// Reference to Restarting Information.
// RestartingInfo returns a IsisIIHRestartInfo
func (obj *isisNeighborIIHState) HasRestartingInfo() bool {
	return obj.obj.RestartingInfo != nil
}

// Reference to Restarting Information.
// SetRestartingInfo sets the IsisIIHRestartInfo value in the IsisNeighborIIHState object
func (obj *isisNeighborIIHState) SetRestartingInfo(value IsisIIHRestartInfo) IsisNeighborIIHState {

	obj.restartingInfoHolder = nil
	obj.obj.RestartingInfo = value.msg()

	return obj
}

// It refers to IIH packet TLVs container.
// Tlvs returns a IsisIIHTlvs
func (obj *isisNeighborIIHState) Tlvs() IsisIIHTlvs {
	if obj.obj.Tlvs == nil {
		obj.obj.Tlvs = NewIsisIIHTlvs().msg()
	}
	if obj.tlvsHolder == nil {
		obj.tlvsHolder = &isisIIHTlvs{obj: obj.obj.Tlvs}
	}
	return obj.tlvsHolder
}

// It refers to IIH packet TLVs container.
// Tlvs returns a IsisIIHTlvs
func (obj *isisNeighborIIHState) HasTlvs() bool {
	return obj.obj.Tlvs != nil
}

// It refers to IIH packet TLVs container.
// SetTlvs sets the IsisIIHTlvs value in the IsisNeighborIIHState object
func (obj *isisNeighborIIHState) SetTlvs(value IsisIIHTlvs) IsisNeighborIIHState {

	obj.tlvsHolder = nil
	obj.obj.Tlvs = value.msg()

	return obj
}

func (obj *isisNeighborIIHState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RestartingInfo != nil {

		obj.RestartingInfo().validateObj(vObj, set_default)
	}

	if obj.obj.Tlvs != nil {

		obj.Tlvs().validateObj(vObj, set_default)
	}

}

func (obj *isisNeighborIIHState) setDefault() {
	if obj.obj.NetworkType == nil {
		obj.SetNetworkType(IsisNeighborIIHStateNetworkType.BROADCAST)

	}

}
