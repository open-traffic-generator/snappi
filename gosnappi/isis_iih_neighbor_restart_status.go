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
	obj                     *otg.IsisIIHNeighborRestartStatus
	marshaller              marshalIsisIIHNeighborRestartStatus
	unMarshaller            unMarshalIsisIIHNeighborRestartStatus
	lastAttemptStatusHolder IsisIIHNeighborGRLastAttemptStatus
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
	obj.lastAttemptStatusHolder = nil
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
	// LastAttemptStatus returns IsisIIHNeighborGRLastAttemptStatus, set in IsisIIHNeighborRestartStatus.
	// IsisIIHNeighborGRLastAttemptStatus is this object contains the status of the last attempted Graceful Restart status of an ISIS neighbor.
	// - succeeded: Choice is set if the last Graceful was successful.
	// - failed: The last Graceful attempt was unsuccessful.
	// - inprogress: The last Graceful Restart is in progress.
	// - unavailable: Graceful Restart has never been initiated by the neighbor.
	LastAttemptStatus() IsisIIHNeighborGRLastAttemptStatus
	// SetLastAttemptStatus assigns IsisIIHNeighborGRLastAttemptStatus provided by user to IsisIIHNeighborRestartStatus.
	// IsisIIHNeighborGRLastAttemptStatus is this object contains the status of the last attempted Graceful Restart status of an ISIS neighbor.
	// - succeeded: Choice is set if the last Graceful was successful.
	// - failed: The last Graceful attempt was unsuccessful.
	// - inprogress: The last Graceful Restart is in progress.
	// - unavailable: Graceful Restart has never been initiated by the neighbor.
	SetLastAttemptStatus(value IsisIIHNeighborGRLastAttemptStatus) IsisIIHNeighborRestartStatus
	// HasLastAttemptStatus checks if LastAttemptStatus has been set in IsisIIHNeighborRestartStatus
	HasLastAttemptStatus() bool
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

// Current State of Neighbor router.
// - starting: Is in Starting state when Restarting Tlv has been received with SA bit set.
// - running: Is in Running state when Restarting Tlv is not present or Restarting Tlv has been received with SA or RR bits unset.
// - restarting: Is in Restarting state when Restarting Tlv has been received with RR bits set.
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

// This container holds the information of the last Graceful Restart initiated  by the neighbor since the adjacency was established.
// LastAttemptStatus returns a IsisIIHNeighborGRLastAttemptStatus
func (obj *isisIIHNeighborRestartStatus) LastAttemptStatus() IsisIIHNeighborGRLastAttemptStatus {
	if obj.obj.LastAttemptStatus == nil {
		obj.obj.LastAttemptStatus = NewIsisIIHNeighborGRLastAttemptStatus().msg()
	}
	if obj.lastAttemptStatusHolder == nil {
		obj.lastAttemptStatusHolder = &isisIIHNeighborGRLastAttemptStatus{obj: obj.obj.LastAttemptStatus}
	}
	return obj.lastAttemptStatusHolder
}

// This container holds the information of the last Graceful Restart initiated  by the neighbor since the adjacency was established.
// LastAttemptStatus returns a IsisIIHNeighborGRLastAttemptStatus
func (obj *isisIIHNeighborRestartStatus) HasLastAttemptStatus() bool {
	return obj.obj.LastAttemptStatus != nil
}

// This container holds the information of the last Graceful Restart initiated  by the neighbor since the adjacency was established.
// SetLastAttemptStatus sets the IsisIIHNeighborGRLastAttemptStatus value in the IsisIIHNeighborRestartStatus object
func (obj *isisIIHNeighborRestartStatus) SetLastAttemptStatus(value IsisIIHNeighborGRLastAttemptStatus) IsisIIHNeighborRestartStatus {

	obj.lastAttemptStatusHolder = nil
	obj.obj.LastAttemptStatus = value.msg()

	return obj
}

func (obj *isisIIHNeighborRestartStatus) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.LastAttemptStatus != nil {

		obj.LastAttemptStatus().validateObj(vObj, set_default)
	}

}

func (obj *isisIIHNeighborRestartStatus) setDefault() {

}
