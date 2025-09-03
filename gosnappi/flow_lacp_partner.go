package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowLacpPartner *****
type flowLacpPartner struct {
	validation
	obj                  *otg.FlowLacpPartner
	marshaller           marshalFlowLacpPartner
	unMarshaller         unMarshalFlowLacpPartner
	typeHolder           PatternFlowLacpPartnerType
	lengthHolder         PatternFlowLacpPartnerLength
	systemPriorityHolder PatternFlowLacpPartnerSystemPriority
	systemIdHolder       PatternFlowLacpPartnerSystemId
	keyHolder            PatternFlowLacpPartnerKey
	portPriorityHolder   PatternFlowLacpPartnerPortPriority
	portNumberHolder     PatternFlowLacpPartnerPortNumber
	partnerStateHolder   FlowLacpPartnerState
	reservedHolder       PatternFlowLacpPartnerReserved
}

func NewFlowLacpPartner() FlowLacpPartner {
	obj := flowLacpPartner{obj: &otg.FlowLacpPartner{}}
	obj.setDefault()
	return &obj
}

func (obj *flowLacpPartner) msg() *otg.FlowLacpPartner {
	return obj.obj
}

func (obj *flowLacpPartner) setMsg(msg *otg.FlowLacpPartner) FlowLacpPartner {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowLacpPartner struct {
	obj *flowLacpPartner
}

type marshalFlowLacpPartner interface {
	// ToProto marshals FlowLacpPartner to protobuf object *otg.FlowLacpPartner
	ToProto() (*otg.FlowLacpPartner, error)
	// ToPbText marshals FlowLacpPartner to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowLacpPartner to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowLacpPartner to JSON text
	ToJson() (string, error)
}

type unMarshalflowLacpPartner struct {
	obj *flowLacpPartner
}

type unMarshalFlowLacpPartner interface {
	// FromProto unmarshals FlowLacpPartner from protobuf object *otg.FlowLacpPartner
	FromProto(msg *otg.FlowLacpPartner) (FlowLacpPartner, error)
	// FromPbText unmarshals FlowLacpPartner from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowLacpPartner from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowLacpPartner from JSON text
	FromJson(value string) error
}

func (obj *flowLacpPartner) Marshal() marshalFlowLacpPartner {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowLacpPartner{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowLacpPartner) Unmarshal() unMarshalFlowLacpPartner {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowLacpPartner{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowLacpPartner) ToProto() (*otg.FlowLacpPartner, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowLacpPartner) FromProto(msg *otg.FlowLacpPartner) (FlowLacpPartner, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowLacpPartner) ToPbText() (string, error) {
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

func (m *unMarshalflowLacpPartner) FromPbText(value string) error {
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

func (m *marshalflowLacpPartner) ToYaml() (string, error) {
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

func (m *unMarshalflowLacpPartner) FromYaml(value string) error {
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

func (m *marshalflowLacpPartner) ToJson() (string, error) {
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

func (m *unMarshalflowLacpPartner) FromJson(value string) error {
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

func (obj *flowLacpPartner) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowLacpPartner) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowLacpPartner) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowLacpPartner) Clone() (FlowLacpPartner, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowLacpPartner()
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

func (obj *flowLacpPartner) setNil() {
	obj.typeHolder = nil
	obj.lengthHolder = nil
	obj.systemPriorityHolder = nil
	obj.systemIdHolder = nil
	obj.keyHolder = nil
	obj.portPriorityHolder = nil
	obj.portNumberHolder = nil
	obj.partnerStateHolder = nil
	obj.reservedHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowLacpPartner is information about the remote (partner) system.
type FlowLacpPartner interface {
	Validation
	// msg marshals FlowLacpPartner to protobuf object *otg.FlowLacpPartner
	// and doesn't set defaults
	msg() *otg.FlowLacpPartner
	// setMsg unmarshals FlowLacpPartner from protobuf object *otg.FlowLacpPartner
	// and doesn't set defaults
	setMsg(*otg.FlowLacpPartner) FlowLacpPartner
	// provides marshal interface
	Marshal() marshalFlowLacpPartner
	// provides unmarshal interface
	Unmarshal() unMarshalFlowLacpPartner
	// validate validates FlowLacpPartner
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowLacpPartner, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns PatternFlowLacpPartnerType, set in FlowLacpPartner.
	// PatternFlowLacpPartnerType is tLV Type for Partner Information. The value 0x02 identifies this TLV as containing information about the remote device (the Partner), as understood by the Actor.
	Type() PatternFlowLacpPartnerType
	// SetType assigns PatternFlowLacpPartnerType provided by user to FlowLacpPartner.
	// PatternFlowLacpPartnerType is tLV Type for Partner Information. The value 0x02 identifies this TLV as containing information about the remote device (the Partner), as understood by the Actor.
	SetType(value PatternFlowLacpPartnerType) FlowLacpPartner
	// HasType checks if Type has been set in FlowLacpPartner
	HasType() bool
	// Length returns PatternFlowLacpPartnerLength, set in FlowLacpPartner.
	// PatternFlowLacpPartnerLength is the length of the Partner Information TLV payload in bytes.
	Length() PatternFlowLacpPartnerLength
	// SetLength assigns PatternFlowLacpPartnerLength provided by user to FlowLacpPartner.
	// PatternFlowLacpPartnerLength is the length of the Partner Information TLV payload in bytes.
	SetLength(value PatternFlowLacpPartnerLength) FlowLacpPartner
	// HasLength checks if Length has been set in FlowLacpPartner
	HasLength() bool
	// SystemPriority returns PatternFlowLacpPartnerSystemPriority, set in FlowLacpPartner.
	// PatternFlowLacpPartnerSystemPriority is the priority of the Partner's system, as received by the Actor.
	SystemPriority() PatternFlowLacpPartnerSystemPriority
	// SetSystemPriority assigns PatternFlowLacpPartnerSystemPriority provided by user to FlowLacpPartner.
	// PatternFlowLacpPartnerSystemPriority is the priority of the Partner's system, as received by the Actor.
	SetSystemPriority(value PatternFlowLacpPartnerSystemPriority) FlowLacpPartner
	// HasSystemPriority checks if SystemPriority has been set in FlowLacpPartner
	HasSystemPriority() bool
	// SystemId returns PatternFlowLacpPartnerSystemId, set in FlowLacpPartner.
	// PatternFlowLacpPartnerSystemId is the Partner's System ID (MAC address), as received by the Actor.
	SystemId() PatternFlowLacpPartnerSystemId
	// SetSystemId assigns PatternFlowLacpPartnerSystemId provided by user to FlowLacpPartner.
	// PatternFlowLacpPartnerSystemId is the Partner's System ID (MAC address), as received by the Actor.
	SetSystemId(value PatternFlowLacpPartnerSystemId) FlowLacpPartner
	// HasSystemId checks if SystemId has been set in FlowLacpPartner
	HasSystemId() bool
	// Key returns PatternFlowLacpPartnerKey, set in FlowLacpPartner.
	// PatternFlowLacpPartnerKey is the operational Key value of the Partner's port, as received by the Actor.
	Key() PatternFlowLacpPartnerKey
	// SetKey assigns PatternFlowLacpPartnerKey provided by user to FlowLacpPartner.
	// PatternFlowLacpPartnerKey is the operational Key value of the Partner's port, as received by the Actor.
	SetKey(value PatternFlowLacpPartnerKey) FlowLacpPartner
	// HasKey checks if Key has been set in FlowLacpPartner
	HasKey() bool
	// PortPriority returns PatternFlowLacpPartnerPortPriority, set in FlowLacpPartner.
	// PatternFlowLacpPartnerPortPriority is the priority of the Partner's port, as received by the Actor.
	PortPriority() PatternFlowLacpPartnerPortPriority
	// SetPortPriority assigns PatternFlowLacpPartnerPortPriority provided by user to FlowLacpPartner.
	// PatternFlowLacpPartnerPortPriority is the priority of the Partner's port, as received by the Actor.
	SetPortPriority(value PatternFlowLacpPartnerPortPriority) FlowLacpPartner
	// HasPortPriority checks if PortPriority has been set in FlowLacpPartner
	HasPortPriority() bool
	// PortNumber returns PatternFlowLacpPartnerPortNumber, set in FlowLacpPartner.
	// PatternFlowLacpPartnerPortNumber is the port number of the Partner's port, as received by the Actor.
	PortNumber() PatternFlowLacpPartnerPortNumber
	// SetPortNumber assigns PatternFlowLacpPartnerPortNumber provided by user to FlowLacpPartner.
	// PatternFlowLacpPartnerPortNumber is the port number of the Partner's port, as received by the Actor.
	SetPortNumber(value PatternFlowLacpPartnerPortNumber) FlowLacpPartner
	// HasPortNumber checks if PortNumber has been set in FlowLacpPartner
	HasPortNumber() bool
	// PartnerState returns FlowLacpPartnerState, set in FlowLacpPartner.
	// FlowLacpPartnerState is this field indicates the Partner's state.
	PartnerState() FlowLacpPartnerState
	// SetPartnerState assigns FlowLacpPartnerState provided by user to FlowLacpPartner.
	// FlowLacpPartnerState is this field indicates the Partner's state.
	SetPartnerState(value FlowLacpPartnerState) FlowLacpPartner
	// HasPartnerState checks if PartnerState has been set in FlowLacpPartner
	HasPartnerState() bool
	// Reserved returns PatternFlowLacpPartnerReserved, set in FlowLacpPartner.
	// PatternFlowLacpPartnerReserved is reserved field for future use in the Partner TLV. Should be set to all zeros.
	Reserved() PatternFlowLacpPartnerReserved
	// SetReserved assigns PatternFlowLacpPartnerReserved provided by user to FlowLacpPartner.
	// PatternFlowLacpPartnerReserved is reserved field for future use in the Partner TLV. Should be set to all zeros.
	SetReserved(value PatternFlowLacpPartnerReserved) FlowLacpPartner
	// HasReserved checks if Reserved has been set in FlowLacpPartner
	HasReserved() bool
	setNil()
}

// description is TBD
// Type returns a PatternFlowLacpPartnerType
func (obj *flowLacpPartner) Type() PatternFlowLacpPartnerType {
	if obj.obj.Type == nil {
		obj.obj.Type = NewPatternFlowLacpPartnerType().msg()
	}
	if obj.typeHolder == nil {
		obj.typeHolder = &patternFlowLacpPartnerType{obj: obj.obj.Type}
	}
	return obj.typeHolder
}

// description is TBD
// Type returns a PatternFlowLacpPartnerType
func (obj *flowLacpPartner) HasType() bool {
	return obj.obj.Type != nil
}

// description is TBD
// SetType sets the PatternFlowLacpPartnerType value in the FlowLacpPartner object
func (obj *flowLacpPartner) SetType(value PatternFlowLacpPartnerType) FlowLacpPartner {

	obj.typeHolder = nil
	obj.obj.Type = value.msg()

	return obj
}

// description is TBD
// Length returns a PatternFlowLacpPartnerLength
func (obj *flowLacpPartner) Length() PatternFlowLacpPartnerLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewPatternFlowLacpPartnerLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &patternFlowLacpPartnerLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// description is TBD
// Length returns a PatternFlowLacpPartnerLength
func (obj *flowLacpPartner) HasLength() bool {
	return obj.obj.Length != nil
}

// description is TBD
// SetLength sets the PatternFlowLacpPartnerLength value in the FlowLacpPartner object
func (obj *flowLacpPartner) SetLength(value PatternFlowLacpPartnerLength) FlowLacpPartner {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// description is TBD
// SystemPriority returns a PatternFlowLacpPartnerSystemPriority
func (obj *flowLacpPartner) SystemPriority() PatternFlowLacpPartnerSystemPriority {
	if obj.obj.SystemPriority == nil {
		obj.obj.SystemPriority = NewPatternFlowLacpPartnerSystemPriority().msg()
	}
	if obj.systemPriorityHolder == nil {
		obj.systemPriorityHolder = &patternFlowLacpPartnerSystemPriority{obj: obj.obj.SystemPriority}
	}
	return obj.systemPriorityHolder
}

// description is TBD
// SystemPriority returns a PatternFlowLacpPartnerSystemPriority
func (obj *flowLacpPartner) HasSystemPriority() bool {
	return obj.obj.SystemPriority != nil
}

// description is TBD
// SetSystemPriority sets the PatternFlowLacpPartnerSystemPriority value in the FlowLacpPartner object
func (obj *flowLacpPartner) SetSystemPriority(value PatternFlowLacpPartnerSystemPriority) FlowLacpPartner {

	obj.systemPriorityHolder = nil
	obj.obj.SystemPriority = value.msg()

	return obj
}

// description is TBD
// SystemId returns a PatternFlowLacpPartnerSystemId
func (obj *flowLacpPartner) SystemId() PatternFlowLacpPartnerSystemId {
	if obj.obj.SystemId == nil {
		obj.obj.SystemId = NewPatternFlowLacpPartnerSystemId().msg()
	}
	if obj.systemIdHolder == nil {
		obj.systemIdHolder = &patternFlowLacpPartnerSystemId{obj: obj.obj.SystemId}
	}
	return obj.systemIdHolder
}

// description is TBD
// SystemId returns a PatternFlowLacpPartnerSystemId
func (obj *flowLacpPartner) HasSystemId() bool {
	return obj.obj.SystemId != nil
}

// description is TBD
// SetSystemId sets the PatternFlowLacpPartnerSystemId value in the FlowLacpPartner object
func (obj *flowLacpPartner) SetSystemId(value PatternFlowLacpPartnerSystemId) FlowLacpPartner {

	obj.systemIdHolder = nil
	obj.obj.SystemId = value.msg()

	return obj
}

// description is TBD
// Key returns a PatternFlowLacpPartnerKey
func (obj *flowLacpPartner) Key() PatternFlowLacpPartnerKey {
	if obj.obj.Key == nil {
		obj.obj.Key = NewPatternFlowLacpPartnerKey().msg()
	}
	if obj.keyHolder == nil {
		obj.keyHolder = &patternFlowLacpPartnerKey{obj: obj.obj.Key}
	}
	return obj.keyHolder
}

// description is TBD
// Key returns a PatternFlowLacpPartnerKey
func (obj *flowLacpPartner) HasKey() bool {
	return obj.obj.Key != nil
}

// description is TBD
// SetKey sets the PatternFlowLacpPartnerKey value in the FlowLacpPartner object
func (obj *flowLacpPartner) SetKey(value PatternFlowLacpPartnerKey) FlowLacpPartner {

	obj.keyHolder = nil
	obj.obj.Key = value.msg()

	return obj
}

// description is TBD
// PortPriority returns a PatternFlowLacpPartnerPortPriority
func (obj *flowLacpPartner) PortPriority() PatternFlowLacpPartnerPortPriority {
	if obj.obj.PortPriority == nil {
		obj.obj.PortPriority = NewPatternFlowLacpPartnerPortPriority().msg()
	}
	if obj.portPriorityHolder == nil {
		obj.portPriorityHolder = &patternFlowLacpPartnerPortPriority{obj: obj.obj.PortPriority}
	}
	return obj.portPriorityHolder
}

// description is TBD
// PortPriority returns a PatternFlowLacpPartnerPortPriority
func (obj *flowLacpPartner) HasPortPriority() bool {
	return obj.obj.PortPriority != nil
}

// description is TBD
// SetPortPriority sets the PatternFlowLacpPartnerPortPriority value in the FlowLacpPartner object
func (obj *flowLacpPartner) SetPortPriority(value PatternFlowLacpPartnerPortPriority) FlowLacpPartner {

	obj.portPriorityHolder = nil
	obj.obj.PortPriority = value.msg()

	return obj
}

// description is TBD
// PortNumber returns a PatternFlowLacpPartnerPortNumber
func (obj *flowLacpPartner) PortNumber() PatternFlowLacpPartnerPortNumber {
	if obj.obj.PortNumber == nil {
		obj.obj.PortNumber = NewPatternFlowLacpPartnerPortNumber().msg()
	}
	if obj.portNumberHolder == nil {
		obj.portNumberHolder = &patternFlowLacpPartnerPortNumber{obj: obj.obj.PortNumber}
	}
	return obj.portNumberHolder
}

// description is TBD
// PortNumber returns a PatternFlowLacpPartnerPortNumber
func (obj *flowLacpPartner) HasPortNumber() bool {
	return obj.obj.PortNumber != nil
}

// description is TBD
// SetPortNumber sets the PatternFlowLacpPartnerPortNumber value in the FlowLacpPartner object
func (obj *flowLacpPartner) SetPortNumber(value PatternFlowLacpPartnerPortNumber) FlowLacpPartner {

	obj.portNumberHolder = nil
	obj.obj.PortNumber = value.msg()

	return obj
}

// description is TBD
// PartnerState returns a FlowLacpPartnerState
func (obj *flowLacpPartner) PartnerState() FlowLacpPartnerState {
	if obj.obj.PartnerState == nil {
		obj.obj.PartnerState = NewFlowLacpPartnerState().msg()
	}
	if obj.partnerStateHolder == nil {
		obj.partnerStateHolder = &flowLacpPartnerState{obj: obj.obj.PartnerState}
	}
	return obj.partnerStateHolder
}

// description is TBD
// PartnerState returns a FlowLacpPartnerState
func (obj *flowLacpPartner) HasPartnerState() bool {
	return obj.obj.PartnerState != nil
}

// description is TBD
// SetPartnerState sets the FlowLacpPartnerState value in the FlowLacpPartner object
func (obj *flowLacpPartner) SetPartnerState(value FlowLacpPartnerState) FlowLacpPartner {

	obj.partnerStateHolder = nil
	obj.obj.PartnerState = value.msg()

	return obj
}

// description is TBD
// Reserved returns a PatternFlowLacpPartnerReserved
func (obj *flowLacpPartner) Reserved() PatternFlowLacpPartnerReserved {
	if obj.obj.Reserved == nil {
		obj.obj.Reserved = NewPatternFlowLacpPartnerReserved().msg()
	}
	if obj.reservedHolder == nil {
		obj.reservedHolder = &patternFlowLacpPartnerReserved{obj: obj.obj.Reserved}
	}
	return obj.reservedHolder
}

// description is TBD
// Reserved returns a PatternFlowLacpPartnerReserved
func (obj *flowLacpPartner) HasReserved() bool {
	return obj.obj.Reserved != nil
}

// description is TBD
// SetReserved sets the PatternFlowLacpPartnerReserved value in the FlowLacpPartner object
func (obj *flowLacpPartner) SetReserved(value PatternFlowLacpPartnerReserved) FlowLacpPartner {

	obj.reservedHolder = nil
	obj.obj.Reserved = value.msg()

	return obj
}

func (obj *flowLacpPartner) validateObj(vObj *validation, set_default bool) {
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

	if obj.obj.PartnerState != nil {

		obj.PartnerState().validateObj(vObj, set_default)
	}

	if obj.obj.Reserved != nil {

		obj.Reserved().validateObj(vObj, set_default)
	}

}

func (obj *flowLacpPartner) setDefault() {

}
