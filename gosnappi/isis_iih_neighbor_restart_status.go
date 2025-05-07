package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisIIHNeighborRestartStatus *****
type isisIIHNeighborRestartStatus struct {
	validation
	obj             *otg.IsisIIHNeighborRestartStatus
	marshaller      marshalIsisIIHNeighborRestartStatus
	unMarshaller    unMarshalIsisIIHNeighborRestartStatus
	succeededHolder IsisIIHNeighborRestartSucceded
}

func NewIsisIIHNeighborRestartStatus() IsisIIHNeighborRestartStatus {
	obj := isisIIHNeighborRestartStatus{obj: &otg.IsisIIHNeighborRestartStatus{}}
	obj.setDefault()
	return &obj
}

func (obj *isisIIHNeighborRestartStatus) msg() *otg.IsisIIHNeighborRestartStatus {
	return obj.obj
}

func (obj *isisIIHNeighborRestartStatus) setMsg(msg *otg.IsisIIHNeighborRestartStatus) IsisIIHNeighborRestartStatus {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisIIHNeighborRestartStatus struct {
	obj *isisIIHNeighborRestartStatus
}

type marshalIsisIIHNeighborRestartStatus interface {
	// ToProto marshals IsisIIHNeighborRestartStatus to protobuf object *otg.IsisIIHNeighborRestartStatus
	ToProto() (*otg.IsisIIHNeighborRestartStatus, error)
	// ToPbText marshals IsisIIHNeighborRestartStatus to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisIIHNeighborRestartStatus to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisIIHNeighborRestartStatus to JSON text
	ToJson() (string, error)
}

type unMarshalisisIIHNeighborRestartStatus struct {
	obj *isisIIHNeighborRestartStatus
}

type unMarshalIsisIIHNeighborRestartStatus interface {
	// FromProto unmarshals IsisIIHNeighborRestartStatus from protobuf object *otg.IsisIIHNeighborRestartStatus
	FromProto(msg *otg.IsisIIHNeighborRestartStatus) (IsisIIHNeighborRestartStatus, error)
	// FromPbText unmarshals IsisIIHNeighborRestartStatus from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisIIHNeighborRestartStatus from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisIIHNeighborRestartStatus from JSON text
	FromJson(value string) error
}

func (obj *isisIIHNeighborRestartStatus) Marshal() marshalIsisIIHNeighborRestartStatus {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisIIHNeighborRestartStatus{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisIIHNeighborRestartStatus) Unmarshal() unMarshalIsisIIHNeighborRestartStatus {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisIIHNeighborRestartStatus{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisIIHNeighborRestartStatus) ToProto() (*otg.IsisIIHNeighborRestartStatus, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisIIHNeighborRestartStatus) FromProto(msg *otg.IsisIIHNeighborRestartStatus) (IsisIIHNeighborRestartStatus, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisIIHNeighborRestartStatus) ToPbText() (string, error) {
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

func (m *unMarshalisisIIHNeighborRestartStatus) FromPbText(value string) error {
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

func (m *marshalisisIIHNeighborRestartStatus) ToYaml() (string, error) {
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

func (m *unMarshalisisIIHNeighborRestartStatus) FromYaml(value string) error {
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

func (m *marshalisisIIHNeighborRestartStatus) ToJson() (string, error) {
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

func (m *unMarshalisisIIHNeighborRestartStatus) FromJson(value string) error {
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

func (obj *isisIIHNeighborRestartStatus) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisIIHNeighborRestartStatus) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisIIHNeighborRestartStatus) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisIIHNeighborRestartStatus) Clone() (IsisIIHNeighborRestartStatus, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisIIHNeighborRestartStatus()
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

func (obj *isisIIHNeighborRestartStatus) setNil() {
	obj.succeededHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisIIHNeighborRestartStatus is this contains the Restarting/Starting/Running state of a neighbor router.
type IsisIIHNeighborRestartStatus interface {
	Validation
	// msg marshals IsisIIHNeighborRestartStatus to protobuf object *otg.IsisIIHNeighborRestartStatus
	// and doesn't set defaults
	msg() *otg.IsisIIHNeighborRestartStatus
	// setMsg unmarshals IsisIIHNeighborRestartStatus from protobuf object *otg.IsisIIHNeighborRestartStatus
	// and doesn't set defaults
	setMsg(*otg.IsisIIHNeighborRestartStatus) IsisIIHNeighborRestartStatus
	// provides marshal interface
	Marshal() marshalIsisIIHNeighborRestartStatus
	// provides unmarshal interface
	Unmarshal() unMarshalIsisIIHNeighborRestartStatus
	// validate validates IsisIIHNeighborRestartStatus
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisIIHNeighborRestartStatus, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// State returns IsisIIHNeighborRestartStatusStateEnum, set in IsisIIHNeighborRestartStatus
	State() IsisIIHNeighborRestartStatusStateEnum
	// SetState assigns IsisIIHNeighborRestartStatusStateEnum provided by user to IsisIIHNeighborRestartStatus
	SetState(value IsisIIHNeighborRestartStatusStateEnum) IsisIIHNeighborRestartStatus
	// HasState checks if State has been set in IsisIIHNeighborRestartStatus
	HasState() bool
	// LastAttemptStatus returns IsisIIHNeighborRestartStatusLastAttemptStatusEnum, set in IsisIIHNeighborRestartStatus
	LastAttemptStatus() IsisIIHNeighborRestartStatusLastAttemptStatusEnum
	// SetLastAttemptStatus assigns IsisIIHNeighborRestartStatusLastAttemptStatusEnum provided by user to IsisIIHNeighborRestartStatus
	SetLastAttemptStatus(value IsisIIHNeighborRestartStatusLastAttemptStatusEnum) IsisIIHNeighborRestartStatus
	// HasLastAttemptStatus checks if LastAttemptStatus has been set in IsisIIHNeighborRestartStatus
	HasLastAttemptStatus() bool
	// Succeeded returns IsisIIHNeighborRestartSucceded, set in IsisIIHNeighborRestartStatus.
	// IsisIIHNeighborRestartSucceded is when neighbor router is Restarting/Starting the Graceful Restart status is successful.
	Succeeded() IsisIIHNeighborRestartSucceded
	// SetSucceeded assigns IsisIIHNeighborRestartSucceded provided by user to IsisIIHNeighborRestartStatus.
	// IsisIIHNeighborRestartSucceded is when neighbor router is Restarting/Starting the Graceful Restart status is successful.
	SetSucceeded(value IsisIIHNeighborRestartSucceded) IsisIIHNeighborRestartStatus
	// HasSucceeded checks if Succeeded has been set in IsisIIHNeighborRestartStatus
	HasSucceeded() bool
	// Failed returns bool, set in IsisIIHNeighborRestartStatus.
	Failed() bool
	// SetFailed assigns bool provided by user to IsisIIHNeighborRestartStatus
	SetFailed(value bool) IsisIIHNeighborRestartStatus
	// HasFailed checks if Failed has been set in IsisIIHNeighborRestartStatus
	HasFailed() bool
	setNil()
}

type IsisIIHNeighborRestartStatusStateEnum string

// Enum of State on IsisIIHNeighborRestartStatus
var IsisIIHNeighborRestartStatusState = struct {
	RUNNING    IsisIIHNeighborRestartStatusStateEnum
	STARTING   IsisIIHNeighborRestartStatusStateEnum
	RESTARTING IsisIIHNeighborRestartStatusStateEnum
}{
	RUNNING:    IsisIIHNeighborRestartStatusStateEnum("running"),
	STARTING:   IsisIIHNeighborRestartStatusStateEnum("starting"),
	RESTARTING: IsisIIHNeighborRestartStatusStateEnum("restarting"),
}

func (obj *isisIIHNeighborRestartStatus) State() IsisIIHNeighborRestartStatusStateEnum {
	return IsisIIHNeighborRestartStatusStateEnum(obj.obj.State.Enum().String())
}

// Current State of Neighbor router when the Restarting Tlv is present in IIH PDU. - starting: Is in Starting state when Restarting Tlv has been received with SA bit set. - running: Is in Running state when Restarting Tlv is not present or Restarting Tlv has been received with SA or RR bits unset. - restarting: Is in Restarting state when Restarting Tlv has been received with RR bits set.
// State returns a string
func (obj *isisIIHNeighborRestartStatus) HasState() bool {
	return obj.obj.State != nil
}

func (obj *isisIIHNeighborRestartStatus) SetState(value IsisIIHNeighborRestartStatusStateEnum) IsisIIHNeighborRestartStatus {
	intValue, ok := otg.IsisIIHNeighborRestartStatus_State_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisIIHNeighborRestartStatusStateEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisIIHNeighborRestartStatus_State_Enum(intValue)
	obj.obj.State = &enumValue

	return obj
}

type IsisIIHNeighborRestartStatusLastAttemptStatusEnum string

// Enum of LastAttemptStatus on IsisIIHNeighborRestartStatus
var IsisIIHNeighborRestartStatusLastAttemptStatus = struct {
	SUCCEEDED IsisIIHNeighborRestartStatusLastAttemptStatusEnum
	FAILED    IsisIIHNeighborRestartStatusLastAttemptStatusEnum
}{
	SUCCEEDED: IsisIIHNeighborRestartStatusLastAttemptStatusEnum("succeeded"),
	FAILED:    IsisIIHNeighborRestartStatusLastAttemptStatusEnum("failed"),
}

func (obj *isisIIHNeighborRestartStatus) LastAttemptStatus() IsisIIHNeighborRestartStatusLastAttemptStatusEnum {
	return IsisIIHNeighborRestartStatusLastAttemptStatusEnum(obj.obj.LastAttemptStatus.Enum().String())
}

// This object holds the information of last Graceful Restart status provied the neigbhor restart happens at least once.
// LastAttemptStatus returns a string
func (obj *isisIIHNeighborRestartStatus) HasLastAttemptStatus() bool {
	return obj.obj.LastAttemptStatus != nil
}

func (obj *isisIIHNeighborRestartStatus) SetLastAttemptStatus(value IsisIIHNeighborRestartStatusLastAttemptStatusEnum) IsisIIHNeighborRestartStatus {
	intValue, ok := otg.IsisIIHNeighborRestartStatus_LastAttemptStatus_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisIIHNeighborRestartStatusLastAttemptStatusEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisIIHNeighborRestartStatus_LastAttemptStatus_Enum(intValue)
	obj.obj.LastAttemptStatus = &enumValue

	return obj
}

// Populated this object when the Graceful Restarting attempt of a neighbor is successful.
// Succeeded returns a IsisIIHNeighborRestartSucceded
func (obj *isisIIHNeighborRestartStatus) Succeeded() IsisIIHNeighborRestartSucceded {
	if obj.obj.Succeeded == nil {
		obj.obj.Succeeded = NewIsisIIHNeighborRestartSucceded().msg()
	}
	if obj.succeededHolder == nil {
		obj.succeededHolder = &isisIIHNeighborRestartSucceded{obj: obj.obj.Succeeded}
	}
	return obj.succeededHolder
}

// Populated this object when the Graceful Restarting attempt of a neighbor is successful.
// Succeeded returns a IsisIIHNeighborRestartSucceded
func (obj *isisIIHNeighborRestartStatus) HasSucceeded() bool {
	return obj.obj.Succeeded != nil
}

// Populated this object when the Graceful Restarting attempt of a neighbor is successful.
// SetSucceeded sets the IsisIIHNeighborRestartSucceded value in the IsisIIHNeighborRestartStatus object
func (obj *isisIIHNeighborRestartStatus) SetSucceeded(value IsisIIHNeighborRestartSucceded) IsisIIHNeighborRestartStatus {

	obj.succeededHolder = nil
	obj.obj.Succeeded = value.msg()

	return obj
}

// This flag is set when the Restarting Router has failed to restart within T3 time.
// Failed returns a bool
func (obj *isisIIHNeighborRestartStatus) Failed() bool {

	return *obj.obj.Failed

}

// This flag is set when the Restarting Router has failed to restart within T3 time.
// Failed returns a bool
func (obj *isisIIHNeighborRestartStatus) HasFailed() bool {
	return obj.obj.Failed != nil
}

// This flag is set when the Restarting Router has failed to restart within T3 time.
// SetFailed sets the bool value in the IsisIIHNeighborRestartStatus object
func (obj *isisIIHNeighborRestartStatus) SetFailed(value bool) IsisIIHNeighborRestartStatus {

	obj.obj.Failed = &value
	return obj
}

func (obj *isisIIHNeighborRestartStatus) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Succeeded != nil {

		obj.Succeeded().validateObj(vObj, set_default)
	}

}

func (obj *isisIIHNeighborRestartStatus) setDefault() {
	if obj.obj.Failed == nil {
		obj.SetFailed(false)
	}

}
