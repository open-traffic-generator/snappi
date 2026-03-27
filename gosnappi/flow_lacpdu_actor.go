package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowLacpduActor *****
type flowLacpduActor struct {
	validation
	obj                  *otg.FlowLacpduActor
	marshaller           marshalFlowLacpduActor
	unMarshaller         unMarshalFlowLacpduActor
	typeHolder           PatternFlowLacpduActorType
	lengthHolder         PatternFlowLacpduActorLength
	systemPriorityHolder PatternFlowLacpduActorSystemPriority
	systemIdHolder       PatternFlowLacpduActorSystemId
	keyHolder            PatternFlowLacpduActorKey
	portPriorityHolder   PatternFlowLacpduActorPortPriority
	portNumberHolder     PatternFlowLacpduActorPortNumber
	actorStateHolder     FlowLacpduActorState
	reservedHolder       PatternFlowLacpduActorReserved
}

func NewFlowLacpduActor() FlowLacpduActor {
	obj := flowLacpduActor{obj: &otg.FlowLacpduActor{}}
	obj.setDefault()
	return &obj
}

func (obj *flowLacpduActor) msg() *otg.FlowLacpduActor {
	return obj.obj
}

func (obj *flowLacpduActor) setMsg(msg *otg.FlowLacpduActor) FlowLacpduActor {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowLacpduActor struct {
	obj *flowLacpduActor
}

type marshalFlowLacpduActor interface {
	// ToProto marshals FlowLacpduActor to protobuf object *otg.FlowLacpduActor
	ToProto() (*otg.FlowLacpduActor, error)
	// ToPbText marshals FlowLacpduActor to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowLacpduActor to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowLacpduActor to JSON text
	ToJson() (string, error)
}

type unMarshalflowLacpduActor struct {
	obj *flowLacpduActor
}

type unMarshalFlowLacpduActor interface {
	// FromProto unmarshals FlowLacpduActor from protobuf object *otg.FlowLacpduActor
	FromProto(msg *otg.FlowLacpduActor) (FlowLacpduActor, error)
	// FromPbText unmarshals FlowLacpduActor from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowLacpduActor from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowLacpduActor from JSON text
	FromJson(value string) error
}

func (obj *flowLacpduActor) Marshal() marshalFlowLacpduActor {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowLacpduActor{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowLacpduActor) Unmarshal() unMarshalFlowLacpduActor {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowLacpduActor{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowLacpduActor) ToProto() (*otg.FlowLacpduActor, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowLacpduActor) FromProto(msg *otg.FlowLacpduActor) (FlowLacpduActor, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowLacpduActor) ToPbText() (string, error) {
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

func (m *unMarshalflowLacpduActor) FromPbText(value string) error {
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

func (m *marshalflowLacpduActor) ToYaml() (string, error) {
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

func (m *unMarshalflowLacpduActor) FromYaml(value string) error {
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

func (m *marshalflowLacpduActor) ToJson() (string, error) {
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

func (m *unMarshalflowLacpduActor) FromJson(value string) error {
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

func (obj *flowLacpduActor) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowLacpduActor) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowLacpduActor) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowLacpduActor) Clone() (FlowLacpduActor, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowLacpduActor()
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

func (obj *flowLacpduActor) setNil() {
	obj.typeHolder = nil
	obj.lengthHolder = nil
	obj.systemPriorityHolder = nil
	obj.systemIdHolder = nil
	obj.keyHolder = nil
	obj.portPriorityHolder = nil
	obj.portNumberHolder = nil
	obj.actorStateHolder = nil
	obj.reservedHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowLacpduActor is information about the local (actor) system.
type FlowLacpduActor interface {
	Validation
	// msg marshals FlowLacpduActor to protobuf object *otg.FlowLacpduActor
	// and doesn't set defaults
	msg() *otg.FlowLacpduActor
	// setMsg unmarshals FlowLacpduActor from protobuf object *otg.FlowLacpduActor
	// and doesn't set defaults
	setMsg(*otg.FlowLacpduActor) FlowLacpduActor
	// provides marshal interface
	Marshal() marshalFlowLacpduActor
	// provides unmarshal interface
	Unmarshal() unMarshalFlowLacpduActor
	// validate validates FlowLacpduActor
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowLacpduActor, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns PatternFlowLacpduActorType, set in FlowLacpduActor.
	// PatternFlowLacpduActorType is tLV Type for Actor Information. The value 0x01 identifies this TLV as containing information about the sending device (the Actor).
	Type() PatternFlowLacpduActorType
	// SetType assigns PatternFlowLacpduActorType provided by user to FlowLacpduActor.
	// PatternFlowLacpduActorType is tLV Type for Actor Information. The value 0x01 identifies this TLV as containing information about the sending device (the Actor).
	SetType(value PatternFlowLacpduActorType) FlowLacpduActor
	// HasType checks if Type has been set in FlowLacpduActor
	HasType() bool
	// Length returns PatternFlowLacpduActorLength, set in FlowLacpduActor.
	// PatternFlowLacpduActorLength is the length of the Actor Information TLV payload in bytes.
	Length() PatternFlowLacpduActorLength
	// SetLength assigns PatternFlowLacpduActorLength provided by user to FlowLacpduActor.
	// PatternFlowLacpduActorLength is the length of the Actor Information TLV payload in bytes.
	SetLength(value PatternFlowLacpduActorLength) FlowLacpduActor
	// HasLength checks if Length has been set in FlowLacpduActor
	HasLength() bool
	// SystemPriority returns PatternFlowLacpduActorSystemPriority, set in FlowLacpduActor.
	// PatternFlowLacpduActorSystemPriority is the priority assigned to the Actor's system for use in aggregation. A lower numerical value indicates a higher priority. Used to select the active System ID when forming an aggregator.
	SystemPriority() PatternFlowLacpduActorSystemPriority
	// SetSystemPriority assigns PatternFlowLacpduActorSystemPriority provided by user to FlowLacpduActor.
	// PatternFlowLacpduActorSystemPriority is the priority assigned to the Actor's system for use in aggregation. A lower numerical value indicates a higher priority. Used to select the active System ID when forming an aggregator.
	SetSystemPriority(value PatternFlowLacpduActorSystemPriority) FlowLacpduActor
	// HasSystemPriority checks if SystemPriority has been set in FlowLacpduActor
	HasSystemPriority() bool
	// SystemId returns PatternFlowLacpduActorSystemId, set in FlowLacpduActor.
	// PatternFlowLacpduActorSystemId is the Actor's System ID, which is a globally unique MAC address assigned to the system containing the Actor.
	SystemId() PatternFlowLacpduActorSystemId
	// SetSystemId assigns PatternFlowLacpduActorSystemId provided by user to FlowLacpduActor.
	// PatternFlowLacpduActorSystemId is the Actor's System ID, which is a globally unique MAC address assigned to the system containing the Actor.
	SetSystemId(value PatternFlowLacpduActorSystemId) FlowLacpduActor
	// HasSystemId checks if SystemId has been set in FlowLacpduActor
	HasSystemId() bool
	// Key returns PatternFlowLacpduActorKey, set in FlowLacpduActor.
	// PatternFlowLacpduActorKey is the operational Key value assigned to the Actor's port. The key is generated based on port configuration (e.g., speed, duplex, trunk ID) and is used to identify potential aggregation groups. Only links with the same key can be aggregated together.
	Key() PatternFlowLacpduActorKey
	// SetKey assigns PatternFlowLacpduActorKey provided by user to FlowLacpduActor.
	// PatternFlowLacpduActorKey is the operational Key value assigned to the Actor's port. The key is generated based on port configuration (e.g., speed, duplex, trunk ID) and is used to identify potential aggregation groups. Only links with the same key can be aggregated together.
	SetKey(value PatternFlowLacpduActorKey) FlowLacpduActor
	// HasKey checks if Key has been set in FlowLacpduActor
	HasKey() bool
	// PortPriority returns PatternFlowLacpduActorPortPriority, set in FlowLacpduActor.
	// PatternFlowLacpduActorPortPriority is the priority assigned to the Actor's port. A lower numerical value indicates a higher priority. Used to prioritize ports for inclusion in a Link Aggregation Group (LAG) when the group is full.
	PortPriority() PatternFlowLacpduActorPortPriority
	// SetPortPriority assigns PatternFlowLacpduActorPortPriority provided by user to FlowLacpduActor.
	// PatternFlowLacpduActorPortPriority is the priority assigned to the Actor's port. A lower numerical value indicates a higher priority. Used to prioritize ports for inclusion in a Link Aggregation Group (LAG) when the group is full.
	SetPortPriority(value PatternFlowLacpduActorPortPriority) FlowLacpduActor
	// HasPortPriority checks if PortPriority has been set in FlowLacpduActor
	HasPortPriority() bool
	// PortNumber returns PatternFlowLacpduActorPortNumber, set in FlowLacpduActor.
	// PatternFlowLacpduActorPortNumber is the port number assigned to the Actor's port. It is a unique identifier for the port within the system.
	PortNumber() PatternFlowLacpduActorPortNumber
	// SetPortNumber assigns PatternFlowLacpduActorPortNumber provided by user to FlowLacpduActor.
	// PatternFlowLacpduActorPortNumber is the port number assigned to the Actor's port. It is a unique identifier for the port within the system.
	SetPortNumber(value PatternFlowLacpduActorPortNumber) FlowLacpduActor
	// HasPortNumber checks if PortNumber has been set in FlowLacpduActor
	HasPortNumber() bool
	// ActorState returns FlowLacpduActorState, set in FlowLacpduActor.
	// FlowLacpduActorState is this field indicates the Actor's state.
	ActorState() FlowLacpduActorState
	// SetActorState assigns FlowLacpduActorState provided by user to FlowLacpduActor.
	// FlowLacpduActorState is this field indicates the Actor's state.
	SetActorState(value FlowLacpduActorState) FlowLacpduActor
	// HasActorState checks if ActorState has been set in FlowLacpduActor
	HasActorState() bool
	// Reserved returns PatternFlowLacpduActorReserved, set in FlowLacpduActor.
	// PatternFlowLacpduActorReserved is reserved field for future use in the Actor TLV. Should be set to all zeros.
	Reserved() PatternFlowLacpduActorReserved
	// SetReserved assigns PatternFlowLacpduActorReserved provided by user to FlowLacpduActor.
	// PatternFlowLacpduActorReserved is reserved field for future use in the Actor TLV. Should be set to all zeros.
	SetReserved(value PatternFlowLacpduActorReserved) FlowLacpduActor
	// HasReserved checks if Reserved has been set in FlowLacpduActor
	HasReserved() bool
	setNil()
}

// description is TBD
// Type returns a PatternFlowLacpduActorType
func (obj *flowLacpduActor) Type() PatternFlowLacpduActorType {
	if obj.obj.Type == nil {
		obj.obj.Type = NewPatternFlowLacpduActorType().msg()
	}
	if obj.typeHolder == nil {
		obj.typeHolder = &patternFlowLacpduActorType{obj: obj.obj.Type}
	}
	return obj.typeHolder
}

// description is TBD
// Type returns a PatternFlowLacpduActorType
func (obj *flowLacpduActor) HasType() bool {
	return obj.obj.Type != nil
}

// description is TBD
// SetType sets the PatternFlowLacpduActorType value in the FlowLacpduActor object
func (obj *flowLacpduActor) SetType(value PatternFlowLacpduActorType) FlowLacpduActor {

	obj.typeHolder = nil
	obj.obj.Type = value.msg()

	return obj
}

// description is TBD
// Length returns a PatternFlowLacpduActorLength
func (obj *flowLacpduActor) Length() PatternFlowLacpduActorLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewPatternFlowLacpduActorLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &patternFlowLacpduActorLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// description is TBD
// Length returns a PatternFlowLacpduActorLength
func (obj *flowLacpduActor) HasLength() bool {
	return obj.obj.Length != nil
}

// description is TBD
// SetLength sets the PatternFlowLacpduActorLength value in the FlowLacpduActor object
func (obj *flowLacpduActor) SetLength(value PatternFlowLacpduActorLength) FlowLacpduActor {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// description is TBD
// SystemPriority returns a PatternFlowLacpduActorSystemPriority
func (obj *flowLacpduActor) SystemPriority() PatternFlowLacpduActorSystemPriority {
	if obj.obj.SystemPriority == nil {
		obj.obj.SystemPriority = NewPatternFlowLacpduActorSystemPriority().msg()
	}
	if obj.systemPriorityHolder == nil {
		obj.systemPriorityHolder = &patternFlowLacpduActorSystemPriority{obj: obj.obj.SystemPriority}
	}
	return obj.systemPriorityHolder
}

// description is TBD
// SystemPriority returns a PatternFlowLacpduActorSystemPriority
func (obj *flowLacpduActor) HasSystemPriority() bool {
	return obj.obj.SystemPriority != nil
}

// description is TBD
// SetSystemPriority sets the PatternFlowLacpduActorSystemPriority value in the FlowLacpduActor object
func (obj *flowLacpduActor) SetSystemPriority(value PatternFlowLacpduActorSystemPriority) FlowLacpduActor {

	obj.systemPriorityHolder = nil
	obj.obj.SystemPriority = value.msg()

	return obj
}

// description is TBD
// SystemId returns a PatternFlowLacpduActorSystemId
func (obj *flowLacpduActor) SystemId() PatternFlowLacpduActorSystemId {
	if obj.obj.SystemId == nil {
		obj.obj.SystemId = NewPatternFlowLacpduActorSystemId().msg()
	}
	if obj.systemIdHolder == nil {
		obj.systemIdHolder = &patternFlowLacpduActorSystemId{obj: obj.obj.SystemId}
	}
	return obj.systemIdHolder
}

// description is TBD
// SystemId returns a PatternFlowLacpduActorSystemId
func (obj *flowLacpduActor) HasSystemId() bool {
	return obj.obj.SystemId != nil
}

// description is TBD
// SetSystemId sets the PatternFlowLacpduActorSystemId value in the FlowLacpduActor object
func (obj *flowLacpduActor) SetSystemId(value PatternFlowLacpduActorSystemId) FlowLacpduActor {

	obj.systemIdHolder = nil
	obj.obj.SystemId = value.msg()

	return obj
}

// description is TBD
// Key returns a PatternFlowLacpduActorKey
func (obj *flowLacpduActor) Key() PatternFlowLacpduActorKey {
	if obj.obj.Key == nil {
		obj.obj.Key = NewPatternFlowLacpduActorKey().msg()
	}
	if obj.keyHolder == nil {
		obj.keyHolder = &patternFlowLacpduActorKey{obj: obj.obj.Key}
	}
	return obj.keyHolder
}

// description is TBD
// Key returns a PatternFlowLacpduActorKey
func (obj *flowLacpduActor) HasKey() bool {
	return obj.obj.Key != nil
}

// description is TBD
// SetKey sets the PatternFlowLacpduActorKey value in the FlowLacpduActor object
func (obj *flowLacpduActor) SetKey(value PatternFlowLacpduActorKey) FlowLacpduActor {

	obj.keyHolder = nil
	obj.obj.Key = value.msg()

	return obj
}

// description is TBD
// PortPriority returns a PatternFlowLacpduActorPortPriority
func (obj *flowLacpduActor) PortPriority() PatternFlowLacpduActorPortPriority {
	if obj.obj.PortPriority == nil {
		obj.obj.PortPriority = NewPatternFlowLacpduActorPortPriority().msg()
	}
	if obj.portPriorityHolder == nil {
		obj.portPriorityHolder = &patternFlowLacpduActorPortPriority{obj: obj.obj.PortPriority}
	}
	return obj.portPriorityHolder
}

// description is TBD
// PortPriority returns a PatternFlowLacpduActorPortPriority
func (obj *flowLacpduActor) HasPortPriority() bool {
	return obj.obj.PortPriority != nil
}

// description is TBD
// SetPortPriority sets the PatternFlowLacpduActorPortPriority value in the FlowLacpduActor object
func (obj *flowLacpduActor) SetPortPriority(value PatternFlowLacpduActorPortPriority) FlowLacpduActor {

	obj.portPriorityHolder = nil
	obj.obj.PortPriority = value.msg()

	return obj
}

// description is TBD
// PortNumber returns a PatternFlowLacpduActorPortNumber
func (obj *flowLacpduActor) PortNumber() PatternFlowLacpduActorPortNumber {
	if obj.obj.PortNumber == nil {
		obj.obj.PortNumber = NewPatternFlowLacpduActorPortNumber().msg()
	}
	if obj.portNumberHolder == nil {
		obj.portNumberHolder = &patternFlowLacpduActorPortNumber{obj: obj.obj.PortNumber}
	}
	return obj.portNumberHolder
}

// description is TBD
// PortNumber returns a PatternFlowLacpduActorPortNumber
func (obj *flowLacpduActor) HasPortNumber() bool {
	return obj.obj.PortNumber != nil
}

// description is TBD
// SetPortNumber sets the PatternFlowLacpduActorPortNumber value in the FlowLacpduActor object
func (obj *flowLacpduActor) SetPortNumber(value PatternFlowLacpduActorPortNumber) FlowLacpduActor {

	obj.portNumberHolder = nil
	obj.obj.PortNumber = value.msg()

	return obj
}

// description is TBD
// ActorState returns a FlowLacpduActorState
func (obj *flowLacpduActor) ActorState() FlowLacpduActorState {
	if obj.obj.ActorState == nil {
		obj.obj.ActorState = NewFlowLacpduActorState().msg()
	}
	if obj.actorStateHolder == nil {
		obj.actorStateHolder = &flowLacpduActorState{obj: obj.obj.ActorState}
	}
	return obj.actorStateHolder
}

// description is TBD
// ActorState returns a FlowLacpduActorState
func (obj *flowLacpduActor) HasActorState() bool {
	return obj.obj.ActorState != nil
}

// description is TBD
// SetActorState sets the FlowLacpduActorState value in the FlowLacpduActor object
func (obj *flowLacpduActor) SetActorState(value FlowLacpduActorState) FlowLacpduActor {

	obj.actorStateHolder = nil
	obj.obj.ActorState = value.msg()

	return obj
}

// description is TBD
// Reserved returns a PatternFlowLacpduActorReserved
func (obj *flowLacpduActor) Reserved() PatternFlowLacpduActorReserved {
	if obj.obj.Reserved == nil {
		obj.obj.Reserved = NewPatternFlowLacpduActorReserved().msg()
	}
	if obj.reservedHolder == nil {
		obj.reservedHolder = &patternFlowLacpduActorReserved{obj: obj.obj.Reserved}
	}
	return obj.reservedHolder
}

// description is TBD
// Reserved returns a PatternFlowLacpduActorReserved
func (obj *flowLacpduActor) HasReserved() bool {
	return obj.obj.Reserved != nil
}

// description is TBD
// SetReserved sets the PatternFlowLacpduActorReserved value in the FlowLacpduActor object
func (obj *flowLacpduActor) SetReserved(value PatternFlowLacpduActorReserved) FlowLacpduActor {

	obj.reservedHolder = nil
	obj.obj.Reserved = value.msg()

	return obj
}

func (obj *flowLacpduActor) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Type != nil {

		obj.Type().validateObj(vObj, set_default)
	}

	if obj.obj.Length != nil {

		obj.Length().validateObj(vObj, set_default)
	}

	if obj.obj.SystemPriority != nil {

		obj.SystemPriority().validateObj(vObj, set_default)
	}

	if obj.obj.SystemId != nil {

		obj.SystemId().validateObj(vObj, set_default)
	}

	if obj.obj.Key != nil {

		obj.Key().validateObj(vObj, set_default)
	}

	if obj.obj.PortPriority != nil {

		obj.PortPriority().validateObj(vObj, set_default)
	}

	if obj.obj.PortNumber != nil {

		obj.PortNumber().validateObj(vObj, set_default)
	}

	if obj.obj.ActorState != nil {

		obj.ActorState().validateObj(vObj, set_default)
	}

	if obj.obj.Reserved != nil {

		obj.Reserved().validateObj(vObj, set_default)
	}

}

func (obj *flowLacpduActor) setDefault() {

}
