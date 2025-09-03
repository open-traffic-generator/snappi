package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowLacpActor *****
type flowLacpActor struct {
	validation
	obj                  *otg.FlowLacpActor
	marshaller           marshalFlowLacpActor
	unMarshaller         unMarshalFlowLacpActor
	typeHolder           PatternFlowLacpActorType
	lengthHolder         PatternFlowLacpActorLength
	systemPriorityHolder PatternFlowLacpActorSystemPriority
	systemIdHolder       PatternFlowLacpActorSystemId
	keyHolder            PatternFlowLacpActorKey
	portPriorityHolder   PatternFlowLacpActorPortPriority
	portNumberHolder     PatternFlowLacpActorPortNumber
	actorStateHolder     FlowLacpActorState
	reservedHolder       PatternFlowLacpActorReserved
}

func NewFlowLacpActor() FlowLacpActor {
	obj := flowLacpActor{obj: &otg.FlowLacpActor{}}
	obj.setDefault()
	return &obj
}

func (obj *flowLacpActor) msg() *otg.FlowLacpActor {
	return obj.obj
}

func (obj *flowLacpActor) setMsg(msg *otg.FlowLacpActor) FlowLacpActor {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowLacpActor struct {
	obj *flowLacpActor
}

type marshalFlowLacpActor interface {
	// ToProto marshals FlowLacpActor to protobuf object *otg.FlowLacpActor
	ToProto() (*otg.FlowLacpActor, error)
	// ToPbText marshals FlowLacpActor to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowLacpActor to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowLacpActor to JSON text
	ToJson() (string, error)
}

type unMarshalflowLacpActor struct {
	obj *flowLacpActor
}

type unMarshalFlowLacpActor interface {
	// FromProto unmarshals FlowLacpActor from protobuf object *otg.FlowLacpActor
	FromProto(msg *otg.FlowLacpActor) (FlowLacpActor, error)
	// FromPbText unmarshals FlowLacpActor from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowLacpActor from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowLacpActor from JSON text
	FromJson(value string) error
}

func (obj *flowLacpActor) Marshal() marshalFlowLacpActor {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowLacpActor{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowLacpActor) Unmarshal() unMarshalFlowLacpActor {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowLacpActor{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowLacpActor) ToProto() (*otg.FlowLacpActor, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowLacpActor) FromProto(msg *otg.FlowLacpActor) (FlowLacpActor, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowLacpActor) ToPbText() (string, error) {
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

func (m *unMarshalflowLacpActor) FromPbText(value string) error {
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

func (m *marshalflowLacpActor) ToYaml() (string, error) {
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

func (m *unMarshalflowLacpActor) FromYaml(value string) error {
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

func (m *marshalflowLacpActor) ToJson() (string, error) {
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

func (m *unMarshalflowLacpActor) FromJson(value string) error {
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

func (obj *flowLacpActor) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowLacpActor) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowLacpActor) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowLacpActor) Clone() (FlowLacpActor, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowLacpActor()
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

func (obj *flowLacpActor) setNil() {
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

// FlowLacpActor is information about the local (actor) system.
type FlowLacpActor interface {
	Validation
	// msg marshals FlowLacpActor to protobuf object *otg.FlowLacpActor
	// and doesn't set defaults
	msg() *otg.FlowLacpActor
	// setMsg unmarshals FlowLacpActor from protobuf object *otg.FlowLacpActor
	// and doesn't set defaults
	setMsg(*otg.FlowLacpActor) FlowLacpActor
	// provides marshal interface
	Marshal() marshalFlowLacpActor
	// provides unmarshal interface
	Unmarshal() unMarshalFlowLacpActor
	// validate validates FlowLacpActor
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowLacpActor, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns PatternFlowLacpActorType, set in FlowLacpActor.
	// PatternFlowLacpActorType is tLV Type for Actor Information. The value 0x01 identifies this TLV as containing information about the sending device (the Actor).
	Type() PatternFlowLacpActorType
	// SetType assigns PatternFlowLacpActorType provided by user to FlowLacpActor.
	// PatternFlowLacpActorType is tLV Type for Actor Information. The value 0x01 identifies this TLV as containing information about the sending device (the Actor).
	SetType(value PatternFlowLacpActorType) FlowLacpActor
	// HasType checks if Type has been set in FlowLacpActor
	HasType() bool
	// Length returns PatternFlowLacpActorLength, set in FlowLacpActor.
	// PatternFlowLacpActorLength is the length of the Actor Information TLV payload in bytes.
	Length() PatternFlowLacpActorLength
	// SetLength assigns PatternFlowLacpActorLength provided by user to FlowLacpActor.
	// PatternFlowLacpActorLength is the length of the Actor Information TLV payload in bytes.
	SetLength(value PatternFlowLacpActorLength) FlowLacpActor
	// HasLength checks if Length has been set in FlowLacpActor
	HasLength() bool
	// SystemPriority returns PatternFlowLacpActorSystemPriority, set in FlowLacpActor.
	// PatternFlowLacpActorSystemPriority is the priority assigned to the Actor's system for use in aggregation. A lower numerical value indicates a higher priority. Used to select the active System ID when forming an aggregator.
	SystemPriority() PatternFlowLacpActorSystemPriority
	// SetSystemPriority assigns PatternFlowLacpActorSystemPriority provided by user to FlowLacpActor.
	// PatternFlowLacpActorSystemPriority is the priority assigned to the Actor's system for use in aggregation. A lower numerical value indicates a higher priority. Used to select the active System ID when forming an aggregator.
	SetSystemPriority(value PatternFlowLacpActorSystemPriority) FlowLacpActor
	// HasSystemPriority checks if SystemPriority has been set in FlowLacpActor
	HasSystemPriority() bool
	// SystemId returns PatternFlowLacpActorSystemId, set in FlowLacpActor.
	// PatternFlowLacpActorSystemId is the Actor's System ID, which is a globally unique MAC address assigned to the system containing the Actor.
	SystemId() PatternFlowLacpActorSystemId
	// SetSystemId assigns PatternFlowLacpActorSystemId provided by user to FlowLacpActor.
	// PatternFlowLacpActorSystemId is the Actor's System ID, which is a globally unique MAC address assigned to the system containing the Actor.
	SetSystemId(value PatternFlowLacpActorSystemId) FlowLacpActor
	// HasSystemId checks if SystemId has been set in FlowLacpActor
	HasSystemId() bool
	// Key returns PatternFlowLacpActorKey, set in FlowLacpActor.
	// PatternFlowLacpActorKey is the operational Key value assigned to the Actor's port. The key is generated based on port configuration (e.g., speed, duplex, trunk ID) and is used to identify potential aggregation groups. Only links with the same key can be aggregated together.
	Key() PatternFlowLacpActorKey
	// SetKey assigns PatternFlowLacpActorKey provided by user to FlowLacpActor.
	// PatternFlowLacpActorKey is the operational Key value assigned to the Actor's port. The key is generated based on port configuration (e.g., speed, duplex, trunk ID) and is used to identify potential aggregation groups. Only links with the same key can be aggregated together.
	SetKey(value PatternFlowLacpActorKey) FlowLacpActor
	// HasKey checks if Key has been set in FlowLacpActor
	HasKey() bool
	// PortPriority returns PatternFlowLacpActorPortPriority, set in FlowLacpActor.
	// PatternFlowLacpActorPortPriority is the priority assigned to the Actor's port. A lower numerical value indicates a higher priority. Used to prioritize ports for inclusion in a Link Aggregation Group (LAG) when the group is full.
	PortPriority() PatternFlowLacpActorPortPriority
	// SetPortPriority assigns PatternFlowLacpActorPortPriority provided by user to FlowLacpActor.
	// PatternFlowLacpActorPortPriority is the priority assigned to the Actor's port. A lower numerical value indicates a higher priority. Used to prioritize ports for inclusion in a Link Aggregation Group (LAG) when the group is full.
	SetPortPriority(value PatternFlowLacpActorPortPriority) FlowLacpActor
	// HasPortPriority checks if PortPriority has been set in FlowLacpActor
	HasPortPriority() bool
	// PortNumber returns PatternFlowLacpActorPortNumber, set in FlowLacpActor.
	// PatternFlowLacpActorPortNumber is the port number assigned to the Actor's port. It is a unique identifier for the port within the system.
	PortNumber() PatternFlowLacpActorPortNumber
	// SetPortNumber assigns PatternFlowLacpActorPortNumber provided by user to FlowLacpActor.
	// PatternFlowLacpActorPortNumber is the port number assigned to the Actor's port. It is a unique identifier for the port within the system.
	SetPortNumber(value PatternFlowLacpActorPortNumber) FlowLacpActor
	// HasPortNumber checks if PortNumber has been set in FlowLacpActor
	HasPortNumber() bool
	// ActorState returns FlowLacpActorState, set in FlowLacpActor.
	// FlowLacpActorState is this field indicates the Actor's state.
	ActorState() FlowLacpActorState
	// SetActorState assigns FlowLacpActorState provided by user to FlowLacpActor.
	// FlowLacpActorState is this field indicates the Actor's state.
	SetActorState(value FlowLacpActorState) FlowLacpActor
	// HasActorState checks if ActorState has been set in FlowLacpActor
	HasActorState() bool
	// Reserved returns PatternFlowLacpActorReserved, set in FlowLacpActor.
	// PatternFlowLacpActorReserved is reserved field for future use in the Actor TLV. Should be set to all zeros.
	Reserved() PatternFlowLacpActorReserved
	// SetReserved assigns PatternFlowLacpActorReserved provided by user to FlowLacpActor.
	// PatternFlowLacpActorReserved is reserved field for future use in the Actor TLV. Should be set to all zeros.
	SetReserved(value PatternFlowLacpActorReserved) FlowLacpActor
	// HasReserved checks if Reserved has been set in FlowLacpActor
	HasReserved() bool
	setNil()
}

// description is TBD
// Type returns a PatternFlowLacpActorType
func (obj *flowLacpActor) Type() PatternFlowLacpActorType {
	if obj.obj.Type == nil {
		obj.obj.Type = NewPatternFlowLacpActorType().msg()
	}
	if obj.typeHolder == nil {
		obj.typeHolder = &patternFlowLacpActorType{obj: obj.obj.Type}
	}
	return obj.typeHolder
}

// description is TBD
// Type returns a PatternFlowLacpActorType
func (obj *flowLacpActor) HasType() bool {
	return obj.obj.Type != nil
}

// description is TBD
// SetType sets the PatternFlowLacpActorType value in the FlowLacpActor object
func (obj *flowLacpActor) SetType(value PatternFlowLacpActorType) FlowLacpActor {

	obj.typeHolder = nil
	obj.obj.Type = value.msg()

	return obj
}

// description is TBD
// Length returns a PatternFlowLacpActorLength
func (obj *flowLacpActor) Length() PatternFlowLacpActorLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewPatternFlowLacpActorLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &patternFlowLacpActorLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// description is TBD
// Length returns a PatternFlowLacpActorLength
func (obj *flowLacpActor) HasLength() bool {
	return obj.obj.Length != nil
}

// description is TBD
// SetLength sets the PatternFlowLacpActorLength value in the FlowLacpActor object
func (obj *flowLacpActor) SetLength(value PatternFlowLacpActorLength) FlowLacpActor {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// description is TBD
// SystemPriority returns a PatternFlowLacpActorSystemPriority
func (obj *flowLacpActor) SystemPriority() PatternFlowLacpActorSystemPriority {
	if obj.obj.SystemPriority == nil {
		obj.obj.SystemPriority = NewPatternFlowLacpActorSystemPriority().msg()
	}
	if obj.systemPriorityHolder == nil {
		obj.systemPriorityHolder = &patternFlowLacpActorSystemPriority{obj: obj.obj.SystemPriority}
	}
	return obj.systemPriorityHolder
}

// description is TBD
// SystemPriority returns a PatternFlowLacpActorSystemPriority
func (obj *flowLacpActor) HasSystemPriority() bool {
	return obj.obj.SystemPriority != nil
}

// description is TBD
// SetSystemPriority sets the PatternFlowLacpActorSystemPriority value in the FlowLacpActor object
func (obj *flowLacpActor) SetSystemPriority(value PatternFlowLacpActorSystemPriority) FlowLacpActor {

	obj.systemPriorityHolder = nil
	obj.obj.SystemPriority = value.msg()

	return obj
}

// description is TBD
// SystemId returns a PatternFlowLacpActorSystemId
func (obj *flowLacpActor) SystemId() PatternFlowLacpActorSystemId {
	if obj.obj.SystemId == nil {
		obj.obj.SystemId = NewPatternFlowLacpActorSystemId().msg()
	}
	if obj.systemIdHolder == nil {
		obj.systemIdHolder = &patternFlowLacpActorSystemId{obj: obj.obj.SystemId}
	}
	return obj.systemIdHolder
}

// description is TBD
// SystemId returns a PatternFlowLacpActorSystemId
func (obj *flowLacpActor) HasSystemId() bool {
	return obj.obj.SystemId != nil
}

// description is TBD
// SetSystemId sets the PatternFlowLacpActorSystemId value in the FlowLacpActor object
func (obj *flowLacpActor) SetSystemId(value PatternFlowLacpActorSystemId) FlowLacpActor {

	obj.systemIdHolder = nil
	obj.obj.SystemId = value.msg()

	return obj
}

// description is TBD
// Key returns a PatternFlowLacpActorKey
func (obj *flowLacpActor) Key() PatternFlowLacpActorKey {
	if obj.obj.Key == nil {
		obj.obj.Key = NewPatternFlowLacpActorKey().msg()
	}
	if obj.keyHolder == nil {
		obj.keyHolder = &patternFlowLacpActorKey{obj: obj.obj.Key}
	}
	return obj.keyHolder
}

// description is TBD
// Key returns a PatternFlowLacpActorKey
func (obj *flowLacpActor) HasKey() bool {
	return obj.obj.Key != nil
}

// description is TBD
// SetKey sets the PatternFlowLacpActorKey value in the FlowLacpActor object
func (obj *flowLacpActor) SetKey(value PatternFlowLacpActorKey) FlowLacpActor {

	obj.keyHolder = nil
	obj.obj.Key = value.msg()

	return obj
}

// description is TBD
// PortPriority returns a PatternFlowLacpActorPortPriority
func (obj *flowLacpActor) PortPriority() PatternFlowLacpActorPortPriority {
	if obj.obj.PortPriority == nil {
		obj.obj.PortPriority = NewPatternFlowLacpActorPortPriority().msg()
	}
	if obj.portPriorityHolder == nil {
		obj.portPriorityHolder = &patternFlowLacpActorPortPriority{obj: obj.obj.PortPriority}
	}
	return obj.portPriorityHolder
}

// description is TBD
// PortPriority returns a PatternFlowLacpActorPortPriority
func (obj *flowLacpActor) HasPortPriority() bool {
	return obj.obj.PortPriority != nil
}

// description is TBD
// SetPortPriority sets the PatternFlowLacpActorPortPriority value in the FlowLacpActor object
func (obj *flowLacpActor) SetPortPriority(value PatternFlowLacpActorPortPriority) FlowLacpActor {

	obj.portPriorityHolder = nil
	obj.obj.PortPriority = value.msg()

	return obj
}

// description is TBD
// PortNumber returns a PatternFlowLacpActorPortNumber
func (obj *flowLacpActor) PortNumber() PatternFlowLacpActorPortNumber {
	if obj.obj.PortNumber == nil {
		obj.obj.PortNumber = NewPatternFlowLacpActorPortNumber().msg()
	}
	if obj.portNumberHolder == nil {
		obj.portNumberHolder = &patternFlowLacpActorPortNumber{obj: obj.obj.PortNumber}
	}
	return obj.portNumberHolder
}

// description is TBD
// PortNumber returns a PatternFlowLacpActorPortNumber
func (obj *flowLacpActor) HasPortNumber() bool {
	return obj.obj.PortNumber != nil
}

// description is TBD
// SetPortNumber sets the PatternFlowLacpActorPortNumber value in the FlowLacpActor object
func (obj *flowLacpActor) SetPortNumber(value PatternFlowLacpActorPortNumber) FlowLacpActor {

	obj.portNumberHolder = nil
	obj.obj.PortNumber = value.msg()

	return obj
}

// description is TBD
// ActorState returns a FlowLacpActorState
func (obj *flowLacpActor) ActorState() FlowLacpActorState {
	if obj.obj.ActorState == nil {
		obj.obj.ActorState = NewFlowLacpActorState().msg()
	}
	if obj.actorStateHolder == nil {
		obj.actorStateHolder = &flowLacpActorState{obj: obj.obj.ActorState}
	}
	return obj.actorStateHolder
}

// description is TBD
// ActorState returns a FlowLacpActorState
func (obj *flowLacpActor) HasActorState() bool {
	return obj.obj.ActorState != nil
}

// description is TBD
// SetActorState sets the FlowLacpActorState value in the FlowLacpActor object
func (obj *flowLacpActor) SetActorState(value FlowLacpActorState) FlowLacpActor {

	obj.actorStateHolder = nil
	obj.obj.ActorState = value.msg()

	return obj
}

// description is TBD
// Reserved returns a PatternFlowLacpActorReserved
func (obj *flowLacpActor) Reserved() PatternFlowLacpActorReserved {
	if obj.obj.Reserved == nil {
		obj.obj.Reserved = NewPatternFlowLacpActorReserved().msg()
	}
	if obj.reservedHolder == nil {
		obj.reservedHolder = &patternFlowLacpActorReserved{obj: obj.obj.Reserved}
	}
	return obj.reservedHolder
}

// description is TBD
// Reserved returns a PatternFlowLacpActorReserved
func (obj *flowLacpActor) HasReserved() bool {
	return obj.obj.Reserved != nil
}

// description is TBD
// SetReserved sets the PatternFlowLacpActorReserved value in the FlowLacpActor object
func (obj *flowLacpActor) SetReserved(value PatternFlowLacpActorReserved) FlowLacpActor {

	obj.reservedHolder = nil
	obj.obj.Reserved = value.msg()

	return obj
}

func (obj *flowLacpActor) validateObj(vObj *validation, set_default bool) {
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

func (obj *flowLacpActor) setDefault() {

}
