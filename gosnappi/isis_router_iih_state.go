package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisRouterIIHState *****
type isisRouterIIHState struct {
	validation
	obj                  *otg.IsisRouterIIHState
	marshaller           marshalIsisRouterIIHState
	unMarshaller         unMarshalIsisRouterIIHState
	restartingInfoHolder IsisIIHRestartInfo
	tlvsHolder           IsisIIHTlvs
}

func NewIsisRouterIIHState() IsisRouterIIHState {
	obj := isisRouterIIHState{obj: &otg.IsisRouterIIHState{}}
	obj.setDefault()
	return &obj
}

func (obj *isisRouterIIHState) msg() *otg.IsisRouterIIHState {
	return obj.obj
}

func (obj *isisRouterIIHState) setMsg(msg *otg.IsisRouterIIHState) IsisRouterIIHState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisRouterIIHState struct {
	obj *isisRouterIIHState
}

type marshalIsisRouterIIHState interface {
	// ToProto marshals IsisRouterIIHState to protobuf object *otg.IsisRouterIIHState
	ToProto() (*otg.IsisRouterIIHState, error)
	// ToPbText marshals IsisRouterIIHState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisRouterIIHState to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisRouterIIHState to JSON text
	ToJson() (string, error)
}

type unMarshalisisRouterIIHState struct {
	obj *isisRouterIIHState
}

type unMarshalIsisRouterIIHState interface {
	// FromProto unmarshals IsisRouterIIHState from protobuf object *otg.IsisRouterIIHState
	FromProto(msg *otg.IsisRouterIIHState) (IsisRouterIIHState, error)
	// FromPbText unmarshals IsisRouterIIHState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisRouterIIHState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisRouterIIHState from JSON text
	FromJson(value string) error
}

func (obj *isisRouterIIHState) Marshal() marshalIsisRouterIIHState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisRouterIIHState{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisRouterIIHState) Unmarshal() unMarshalIsisRouterIIHState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisRouterIIHState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisRouterIIHState) ToProto() (*otg.IsisRouterIIHState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisRouterIIHState) FromProto(msg *otg.IsisRouterIIHState) (IsisRouterIIHState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisRouterIIHState) ToPbText() (string, error) {
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

func (m *unMarshalisisRouterIIHState) FromPbText(value string) error {
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

func (m *marshalisisRouterIIHState) ToYaml() (string, error) {
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

func (m *unMarshalisisRouterIIHState) FromYaml(value string) error {
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

func (m *marshalisisRouterIIHState) ToJson() (string, error) {
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

func (m *unMarshalisisRouterIIHState) FromJson(value string) error {
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

func (obj *isisRouterIIHState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisRouterIIHState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisRouterIIHState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisRouterIIHState) Clone() (IsisRouterIIHState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisRouterIIHState()
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

func (obj *isisRouterIIHState) setNil() {
	obj.restartingInfoHolder = nil
	obj.tlvsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisRouterIIHState is a ISIS Router IIH Information.
type IsisRouterIIHState interface {
	Validation
	// msg marshals IsisRouterIIHState to protobuf object *otg.IsisRouterIIHState
	// and doesn't set defaults
	msg() *otg.IsisRouterIIHState
	// setMsg unmarshals IsisRouterIIHState from protobuf object *otg.IsisRouterIIHState
	// and doesn't set defaults
	setMsg(*otg.IsisRouterIIHState) IsisRouterIIHState
	// provides marshal interface
	Marshal() marshalIsisRouterIIHState
	// provides unmarshal interface
	Unmarshal() unMarshalIsisRouterIIHState
	// validate validates IsisRouterIIHState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisRouterIIHState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LevelType returns IsisRouterIIHStateLevelTypeEnum, set in IsisRouterIIHState
	LevelType() IsisRouterIIHStateLevelTypeEnum
	// SetLevelType assigns IsisRouterIIHStateLevelTypeEnum provided by user to IsisRouterIIHState
	SetLevelType(value IsisRouterIIHStateLevelTypeEnum) IsisRouterIIHState
	// HasLevelType checks if LevelType has been set in IsisRouterIIHState
	HasLevelType() bool
	// HoldTimer returns uint32, set in IsisRouterIIHState.
	HoldTimer() uint32
	// SetHoldTimer assigns uint32 provided by user to IsisRouterIIHState
	SetHoldTimer(value uint32) IsisRouterIIHState
	// HasHoldTimer checks if HoldTimer has been set in IsisRouterIIHState
	HasHoldTimer() bool
	// RestartingInfo returns IsisIIHRestartInfo, set in IsisRouterIIHState.
	// IsisIIHRestartInfo is this contains the list of TLVs present in a IIH packet.
	RestartingInfo() IsisIIHRestartInfo
	// SetRestartingInfo assigns IsisIIHRestartInfo provided by user to IsisRouterIIHState.
	// IsisIIHRestartInfo is this contains the list of TLVs present in a IIH packet.
	SetRestartingInfo(value IsisIIHRestartInfo) IsisRouterIIHState
	// HasRestartingInfo checks if RestartingInfo has been set in IsisRouterIIHState
	HasRestartingInfo() bool
	// Tlvs returns IsisIIHTlvs, set in IsisRouterIIHState.
	// IsisIIHTlvs is this contains the list of TLVs present in a IIH packet.
	Tlvs() IsisIIHTlvs
	// SetTlvs assigns IsisIIHTlvs provided by user to IsisRouterIIHState.
	// IsisIIHTlvs is this contains the list of TLVs present in a IIH packet.
	SetTlvs(value IsisIIHTlvs) IsisRouterIIHState
	// HasTlvs checks if Tlvs has been set in IsisRouterIIHState
	HasTlvs() bool
	setNil()
}

type IsisRouterIIHStateLevelTypeEnum string

// Enum of LevelType on IsisRouterIIHState
var IsisRouterIIHStateLevelType = struct {
	LEVEL_1   IsisRouterIIHStateLevelTypeEnum
	LEVEL_2   IsisRouterIIHStateLevelTypeEnum
	LEVEL_1_2 IsisRouterIIHStateLevelTypeEnum
}{
	LEVEL_1:   IsisRouterIIHStateLevelTypeEnum("level_1"),
	LEVEL_2:   IsisRouterIIHStateLevelTypeEnum("level_2"),
	LEVEL_1_2: IsisRouterIIHStateLevelTypeEnum("level_1_2"),
}

func (obj *isisRouterIIHState) LevelType() IsisRouterIIHStateLevelTypeEnum {
	return IsisRouterIIHStateLevelTypeEnum(obj.obj.LevelType.Enum().String())
}

// This indicates whether this router is participating in Level-1 (L1),  Level-2 (L2) or both L1 and L2 domains on this interface.
// LevelType returns a string
func (obj *isisRouterIIHState) HasLevelType() bool {
	return obj.obj.LevelType != nil
}

func (obj *isisRouterIIHState) SetLevelType(value IsisRouterIIHStateLevelTypeEnum) IsisRouterIIHState {
	intValue, ok := otg.IsisRouterIIHState_LevelType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisRouterIIHStateLevelTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisRouterIIHState_LevelType_Enum(intValue)
	obj.obj.LevelType = &enumValue

	return obj
}

// Hold timer is in the IIH
// HoldTimer returns a uint32
func (obj *isisRouterIIHState) HoldTimer() uint32 {

	return *obj.obj.HoldTimer

}

// Hold timer is in the IIH
// HoldTimer returns a uint32
func (obj *isisRouterIIHState) HasHoldTimer() bool {
	return obj.obj.HoldTimer != nil
}

// Hold timer is in the IIH
// SetHoldTimer sets the uint32 value in the IsisRouterIIHState object
func (obj *isisRouterIIHState) SetHoldTimer(value uint32) IsisRouterIIHState {

	obj.obj.HoldTimer = &value
	return obj
}

// Reference to Restarting Information.
// RestartingInfo returns a IsisIIHRestartInfo
func (obj *isisRouterIIHState) RestartingInfo() IsisIIHRestartInfo {
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
func (obj *isisRouterIIHState) HasRestartingInfo() bool {
	return obj.obj.RestartingInfo != nil
}

// Reference to Restarting Information.
// SetRestartingInfo sets the IsisIIHRestartInfo value in the IsisRouterIIHState object
func (obj *isisRouterIIHState) SetRestartingInfo(value IsisIIHRestartInfo) IsisRouterIIHState {

	obj.restartingInfoHolder = nil
	obj.obj.RestartingInfo = value.msg()

	return obj
}

// It refers to IIH packet TLVs container.
// Tlvs returns a IsisIIHTlvs
func (obj *isisRouterIIHState) Tlvs() IsisIIHTlvs {
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
func (obj *isisRouterIIHState) HasTlvs() bool {
	return obj.obj.Tlvs != nil
}

// It refers to IIH packet TLVs container.
// SetTlvs sets the IsisIIHTlvs value in the IsisRouterIIHState object
func (obj *isisRouterIIHState) SetTlvs(value IsisIIHTlvs) IsisRouterIIHState {

	obj.tlvsHolder = nil
	obj.obj.Tlvs = value.msg()

	return obj
}

func (obj *isisRouterIIHState) validateObj(vObj *validation, set_default bool) {
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

func (obj *isisRouterIIHState) setDefault() {

}
