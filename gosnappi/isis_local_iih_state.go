package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLocalIIHState *****
type isisLocalIIHState struct {
	validation
	obj                    *otg.IsisLocalIIHState
	marshaller             marshalIsisLocalIIHState
	unMarshaller           unMarshalIsisLocalIIHState
	restartingStatusHolder IsisIIHLocalRestartStatus
}

func NewIsisLocalIIHState() IsisLocalIIHState {
	obj := isisLocalIIHState{obj: &otg.IsisLocalIIHState{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLocalIIHState) msg() *otg.IsisLocalIIHState {
	return obj.obj
}

func (obj *isisLocalIIHState) setMsg(msg *otg.IsisLocalIIHState) IsisLocalIIHState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLocalIIHState struct {
	obj *isisLocalIIHState
}

type marshalIsisLocalIIHState interface {
	// ToProto marshals IsisLocalIIHState to protobuf object *otg.IsisLocalIIHState
	ToProto() (*otg.IsisLocalIIHState, error)
	// ToPbText marshals IsisLocalIIHState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLocalIIHState to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLocalIIHState to JSON text
	ToJson() (string, error)
}

type unMarshalisisLocalIIHState struct {
	obj *isisLocalIIHState
}

type unMarshalIsisLocalIIHState interface {
	// FromProto unmarshals IsisLocalIIHState from protobuf object *otg.IsisLocalIIHState
	FromProto(msg *otg.IsisLocalIIHState) (IsisLocalIIHState, error)
	// FromPbText unmarshals IsisLocalIIHState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLocalIIHState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLocalIIHState from JSON text
	FromJson(value string) error
}

func (obj *isisLocalIIHState) Marshal() marshalIsisLocalIIHState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLocalIIHState{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLocalIIHState) Unmarshal() unMarshalIsisLocalIIHState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLocalIIHState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLocalIIHState) ToProto() (*otg.IsisLocalIIHState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLocalIIHState) FromProto(msg *otg.IsisLocalIIHState) (IsisLocalIIHState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLocalIIHState) ToPbText() (string, error) {
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

func (m *unMarshalisisLocalIIHState) FromPbText(value string) error {
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

func (m *marshalisisLocalIIHState) ToYaml() (string, error) {
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

func (m *unMarshalisisLocalIIHState) FromYaml(value string) error {
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

func (m *marshalisisLocalIIHState) ToJson() (string, error) {
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

func (m *unMarshalisisLocalIIHState) FromJson(value string) error {
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

func (obj *isisLocalIIHState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLocalIIHState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLocalIIHState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLocalIIHState) Clone() (IsisLocalIIHState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLocalIIHState()
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

func (obj *isisLocalIIHState) setNil() {
	obj.restartingStatusHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisLocalIIHState is information for a local adjacency.
type IsisLocalIIHState interface {
	Validation
	// msg marshals IsisLocalIIHState to protobuf object *otg.IsisLocalIIHState
	// and doesn't set defaults
	msg() *otg.IsisLocalIIHState
	// setMsg unmarshals IsisLocalIIHState from protobuf object *otg.IsisLocalIIHState
	// and doesn't set defaults
	setMsg(*otg.IsisLocalIIHState) IsisLocalIIHState
	// provides marshal interface
	Marshal() marshalIsisLocalIIHState
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLocalIIHState
	// validate validates IsisLocalIIHState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLocalIIHState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LevelType returns IsisLocalIIHStateLevelTypeEnum, set in IsisLocalIIHState
	LevelType() IsisLocalIIHStateLevelTypeEnum
	// SetLevelType assigns IsisLocalIIHStateLevelTypeEnum provided by user to IsisLocalIIHState
	SetLevelType(value IsisLocalIIHStateLevelTypeEnum) IsisLocalIIHState
	// HasLevelType checks if LevelType has been set in IsisLocalIIHState
	HasLevelType() bool
	// HoldTimer returns uint32, set in IsisLocalIIHState.
	HoldTimer() uint32
	// SetHoldTimer assigns uint32 provided by user to IsisLocalIIHState
	SetHoldTimer(value uint32) IsisLocalIIHState
	// HasHoldTimer checks if HoldTimer has been set in IsisLocalIIHState
	HasHoldTimer() bool
	// RestartingStatus returns IsisIIHLocalRestartStatus, set in IsisLocalIIHState.
	// IsisIIHLocalRestartStatus is this contains the Restarting/Starting/Running state of this router.
	RestartingStatus() IsisIIHLocalRestartStatus
	// SetRestartingStatus assigns IsisIIHLocalRestartStatus provided by user to IsisLocalIIHState.
	// IsisIIHLocalRestartStatus is this contains the Restarting/Starting/Running state of this router.
	SetRestartingStatus(value IsisIIHLocalRestartStatus) IsisLocalIIHState
	// HasRestartingStatus checks if RestartingStatus has been set in IsisLocalIIHState
	HasRestartingStatus() bool
	setNil()
}

type IsisLocalIIHStateLevelTypeEnum string

// Enum of LevelType on IsisLocalIIHState
var IsisLocalIIHStateLevelType = struct {
	LEVEL_1   IsisLocalIIHStateLevelTypeEnum
	LEVEL_2   IsisLocalIIHStateLevelTypeEnum
	LEVEL_1_2 IsisLocalIIHStateLevelTypeEnum
}{
	LEVEL_1:   IsisLocalIIHStateLevelTypeEnum("level_1"),
	LEVEL_2:   IsisLocalIIHStateLevelTypeEnum("level_2"),
	LEVEL_1_2: IsisLocalIIHStateLevelTypeEnum("level_1_2"),
}

func (obj *isisLocalIIHState) LevelType() IsisLocalIIHStateLevelTypeEnum {
	return IsisLocalIIHStateLevelTypeEnum(obj.obj.LevelType.Enum().String())
}

// This indicates whether this IS-IS router is participating in Level-1 (L1),
// Level-2 (L2) or both L1 and L2 domains on this interface.
// LevelType returns a string
func (obj *isisLocalIIHState) HasLevelType() bool {
	return obj.obj.LevelType != nil
}

func (obj *isisLocalIIHState) SetLevelType(value IsisLocalIIHStateLevelTypeEnum) IsisLocalIIHState {
	intValue, ok := otg.IsisLocalIIHState_LevelType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisLocalIIHStateLevelTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisLocalIIHState_LevelType_Enum(intValue)
	obj.obj.LevelType = &enumValue

	return obj
}

// Hold timer being sent in the IIH PDU.
// HoldTimer returns a uint32
func (obj *isisLocalIIHState) HoldTimer() uint32 {

	return *obj.obj.HoldTimer

}

// Hold timer being sent in the IIH PDU.
// HoldTimer returns a uint32
func (obj *isisLocalIIHState) HasHoldTimer() bool {
	return obj.obj.HoldTimer != nil
}

// Hold timer being sent in the IIH PDU.
// SetHoldTimer sets the uint32 value in the IsisLocalIIHState object
func (obj *isisLocalIIHState) SetHoldTimer(value uint32) IsisLocalIIHState {

	obj.obj.HoldTimer = &value
	return obj
}

// Reference to Restarting Information.
// RestartingStatus returns a IsisIIHLocalRestartStatus
func (obj *isisLocalIIHState) RestartingStatus() IsisIIHLocalRestartStatus {
	if obj.obj.RestartingStatus == nil {
		obj.obj.RestartingStatus = NewIsisIIHLocalRestartStatus().msg()
	}
	if obj.restartingStatusHolder == nil {
		obj.restartingStatusHolder = &isisIIHLocalRestartStatus{obj: obj.obj.RestartingStatus}
	}
	return obj.restartingStatusHolder
}

// Reference to Restarting Information.
// RestartingStatus returns a IsisIIHLocalRestartStatus
func (obj *isisLocalIIHState) HasRestartingStatus() bool {
	return obj.obj.RestartingStatus != nil
}

// Reference to Restarting Information.
// SetRestartingStatus sets the IsisIIHLocalRestartStatus value in the IsisLocalIIHState object
func (obj *isisLocalIIHState) SetRestartingStatus(value IsisIIHLocalRestartStatus) IsisLocalIIHState {

	obj.restartingStatusHolder = nil
	obj.obj.RestartingStatus = value.msg()

	return obj
}

func (obj *isisLocalIIHState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RestartingStatus != nil {

		obj.RestartingStatus().validateObj(vObj, set_default)
	}

}

func (obj *isisLocalIIHState) setDefault() {

}
