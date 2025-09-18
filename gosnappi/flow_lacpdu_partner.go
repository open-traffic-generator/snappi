package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowLacpduPartner *****
type flowLacpduPartner struct {
	validation
	obj                  *otg.FlowLacpduPartner
	marshaller           marshalFlowLacpduPartner
	unMarshaller         unMarshalFlowLacpduPartner
	typeHolder           PatternFlowLacpduPartnerType
	lengthHolder         PatternFlowLacpduPartnerLength
	systemPriorityHolder PatternFlowLacpduPartnerSystemPriority
	systemIdHolder       PatternFlowLacpduPartnerSystemId
	keyHolder            PatternFlowLacpduPartnerKey
	portPriorityHolder   PatternFlowLacpduPartnerPortPriority
	portNumberHolder     PatternFlowLacpduPartnerPortNumber
	partnerStateHolder   FlowLacpduPartnerState
	reservedHolder       PatternFlowLacpduPartnerReserved
}

func NewFlowLacpduPartner() FlowLacpduPartner {
	obj := flowLacpduPartner{obj: &otg.FlowLacpduPartner{}}
	obj.setDefault()
	return &obj
}

func (obj *flowLacpduPartner) msg() *otg.FlowLacpduPartner {
	return obj.obj
}

func (obj *flowLacpduPartner) setMsg(msg *otg.FlowLacpduPartner) FlowLacpduPartner {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowLacpduPartner struct {
	obj *flowLacpduPartner
}

type marshalFlowLacpduPartner interface {
	// ToProto marshals FlowLacpduPartner to protobuf object *otg.FlowLacpduPartner
	ToProto() (*otg.FlowLacpduPartner, error)
	// ToPbText marshals FlowLacpduPartner to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowLacpduPartner to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowLacpduPartner to JSON text
	ToJson() (string, error)
}

type unMarshalflowLacpduPartner struct {
	obj *flowLacpduPartner
}

type unMarshalFlowLacpduPartner interface {
	// FromProto unmarshals FlowLacpduPartner from protobuf object *otg.FlowLacpduPartner
	FromProto(msg *otg.FlowLacpduPartner) (FlowLacpduPartner, error)
	// FromPbText unmarshals FlowLacpduPartner from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowLacpduPartner from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowLacpduPartner from JSON text
	FromJson(value string) error
}

func (obj *flowLacpduPartner) Marshal() marshalFlowLacpduPartner {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowLacpduPartner{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowLacpduPartner) Unmarshal() unMarshalFlowLacpduPartner {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowLacpduPartner{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowLacpduPartner) ToProto() (*otg.FlowLacpduPartner, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowLacpduPartner) FromProto(msg *otg.FlowLacpduPartner) (FlowLacpduPartner, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowLacpduPartner) ToPbText() (string, error) {
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

func (m *unMarshalflowLacpduPartner) FromPbText(value string) error {
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

func (m *marshalflowLacpduPartner) ToYaml() (string, error) {
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

func (m *unMarshalflowLacpduPartner) FromYaml(value string) error {
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

func (m *marshalflowLacpduPartner) ToJson() (string, error) {
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

func (m *unMarshalflowLacpduPartner) FromJson(value string) error {
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

func (obj *flowLacpduPartner) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowLacpduPartner) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowLacpduPartner) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowLacpduPartner) Clone() (FlowLacpduPartner, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowLacpduPartner()
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

func (obj *flowLacpduPartner) setNil() {
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

// FlowLacpduPartner is information about the remote (partner) system.
type FlowLacpduPartner interface {
	Validation
	// msg marshals FlowLacpduPartner to protobuf object *otg.FlowLacpduPartner
	// and doesn't set defaults
	msg() *otg.FlowLacpduPartner
	// setMsg unmarshals FlowLacpduPartner from protobuf object *otg.FlowLacpduPartner
	// and doesn't set defaults
	setMsg(*otg.FlowLacpduPartner) FlowLacpduPartner
	// provides marshal interface
	Marshal() marshalFlowLacpduPartner
	// provides unmarshal interface
	Unmarshal() unMarshalFlowLacpduPartner
	// validate validates FlowLacpduPartner
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowLacpduPartner, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns PatternFlowLacpduPartnerType, set in FlowLacpduPartner.
	// PatternFlowLacpduPartnerType is tLV Type for Partner Information. The value 0x02 identifies this TLV as containing information about the remote device (the Partner), as understood by the Actor.
	Type() PatternFlowLacpduPartnerType
	// SetType assigns PatternFlowLacpduPartnerType provided by user to FlowLacpduPartner.
	// PatternFlowLacpduPartnerType is tLV Type for Partner Information. The value 0x02 identifies this TLV as containing information about the remote device (the Partner), as understood by the Actor.
	SetType(value PatternFlowLacpduPartnerType) FlowLacpduPartner
	// HasType checks if Type has been set in FlowLacpduPartner
	HasType() bool
	// Length returns PatternFlowLacpduPartnerLength, set in FlowLacpduPartner.
	// PatternFlowLacpduPartnerLength is the length of the Partner Information TLV payload in bytes.
	Length() PatternFlowLacpduPartnerLength
	// SetLength assigns PatternFlowLacpduPartnerLength provided by user to FlowLacpduPartner.
	// PatternFlowLacpduPartnerLength is the length of the Partner Information TLV payload in bytes.
	SetLength(value PatternFlowLacpduPartnerLength) FlowLacpduPartner
	// HasLength checks if Length has been set in FlowLacpduPartner
	HasLength() bool
	// SystemPriority returns PatternFlowLacpduPartnerSystemPriority, set in FlowLacpduPartner.
	// PatternFlowLacpduPartnerSystemPriority is the priority of the Partner's system, as received by the Actor.
	SystemPriority() PatternFlowLacpduPartnerSystemPriority
	// SetSystemPriority assigns PatternFlowLacpduPartnerSystemPriority provided by user to FlowLacpduPartner.
	// PatternFlowLacpduPartnerSystemPriority is the priority of the Partner's system, as received by the Actor.
	SetSystemPriority(value PatternFlowLacpduPartnerSystemPriority) FlowLacpduPartner
	// HasSystemPriority checks if SystemPriority has been set in FlowLacpduPartner
	HasSystemPriority() bool
	// SystemId returns PatternFlowLacpduPartnerSystemId, set in FlowLacpduPartner.
	// PatternFlowLacpduPartnerSystemId is the Partner's System ID (MAC address), as received by the Actor.
	SystemId() PatternFlowLacpduPartnerSystemId
	// SetSystemId assigns PatternFlowLacpduPartnerSystemId provided by user to FlowLacpduPartner.
	// PatternFlowLacpduPartnerSystemId is the Partner's System ID (MAC address), as received by the Actor.
	SetSystemId(value PatternFlowLacpduPartnerSystemId) FlowLacpduPartner
	// HasSystemId checks if SystemId has been set in FlowLacpduPartner
	HasSystemId() bool
	// Key returns PatternFlowLacpduPartnerKey, set in FlowLacpduPartner.
	// PatternFlowLacpduPartnerKey is the operational Key value of the Partner's port, as received by the Actor.
	Key() PatternFlowLacpduPartnerKey
	// SetKey assigns PatternFlowLacpduPartnerKey provided by user to FlowLacpduPartner.
	// PatternFlowLacpduPartnerKey is the operational Key value of the Partner's port, as received by the Actor.
	SetKey(value PatternFlowLacpduPartnerKey) FlowLacpduPartner
	// HasKey checks if Key has been set in FlowLacpduPartner
	HasKey() bool
	// PortPriority returns PatternFlowLacpduPartnerPortPriority, set in FlowLacpduPartner.
	// PatternFlowLacpduPartnerPortPriority is the priority of the Partner's port, as received by the Actor.
	PortPriority() PatternFlowLacpduPartnerPortPriority
	// SetPortPriority assigns PatternFlowLacpduPartnerPortPriority provided by user to FlowLacpduPartner.
	// PatternFlowLacpduPartnerPortPriority is the priority of the Partner's port, as received by the Actor.
	SetPortPriority(value PatternFlowLacpduPartnerPortPriority) FlowLacpduPartner
	// HasPortPriority checks if PortPriority has been set in FlowLacpduPartner
	HasPortPriority() bool
	// PortNumber returns PatternFlowLacpduPartnerPortNumber, set in FlowLacpduPartner.
	// PatternFlowLacpduPartnerPortNumber is the port number of the Partner's port, as received by the Actor.
	PortNumber() PatternFlowLacpduPartnerPortNumber
	// SetPortNumber assigns PatternFlowLacpduPartnerPortNumber provided by user to FlowLacpduPartner.
	// PatternFlowLacpduPartnerPortNumber is the port number of the Partner's port, as received by the Actor.
	SetPortNumber(value PatternFlowLacpduPartnerPortNumber) FlowLacpduPartner
	// HasPortNumber checks if PortNumber has been set in FlowLacpduPartner
	HasPortNumber() bool
	// PartnerState returns FlowLacpduPartnerState, set in FlowLacpduPartner.
	// FlowLacpduPartnerState is this field indicates the Partner's state.
	PartnerState() FlowLacpduPartnerState
	// SetPartnerState assigns FlowLacpduPartnerState provided by user to FlowLacpduPartner.
	// FlowLacpduPartnerState is this field indicates the Partner's state.
	SetPartnerState(value FlowLacpduPartnerState) FlowLacpduPartner
	// HasPartnerState checks if PartnerState has been set in FlowLacpduPartner
	HasPartnerState() bool
	// Reserved returns PatternFlowLacpduPartnerReserved, set in FlowLacpduPartner.
	// PatternFlowLacpduPartnerReserved is reserved field for future use in the Partner TLV. Should be set to all zeros.
	Reserved() PatternFlowLacpduPartnerReserved
	// SetReserved assigns PatternFlowLacpduPartnerReserved provided by user to FlowLacpduPartner.
	// PatternFlowLacpduPartnerReserved is reserved field for future use in the Partner TLV. Should be set to all zeros.
	SetReserved(value PatternFlowLacpduPartnerReserved) FlowLacpduPartner
	// HasReserved checks if Reserved has been set in FlowLacpduPartner
	HasReserved() bool
	setNil()
}

// description is TBD
// Type returns a PatternFlowLacpduPartnerType
func (obj *flowLacpduPartner) Type() PatternFlowLacpduPartnerType {
	if obj.obj.Type == nil {
		obj.obj.Type = NewPatternFlowLacpduPartnerType().msg()
	}
	if obj.typeHolder == nil {
		obj.typeHolder = &patternFlowLacpduPartnerType{obj: obj.obj.Type}
	}
	return obj.typeHolder
}

// description is TBD
// Type returns a PatternFlowLacpduPartnerType
func (obj *flowLacpduPartner) HasType() bool {
	return obj.obj.Type != nil
}

// description is TBD
// SetType sets the PatternFlowLacpduPartnerType value in the FlowLacpduPartner object
func (obj *flowLacpduPartner) SetType(value PatternFlowLacpduPartnerType) FlowLacpduPartner {

	obj.typeHolder = nil
	obj.obj.Type = value.msg()

	return obj
}

// description is TBD
// Length returns a PatternFlowLacpduPartnerLength
func (obj *flowLacpduPartner) Length() PatternFlowLacpduPartnerLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewPatternFlowLacpduPartnerLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &patternFlowLacpduPartnerLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// description is TBD
// Length returns a PatternFlowLacpduPartnerLength
func (obj *flowLacpduPartner) HasLength() bool {
	return obj.obj.Length != nil
}

// description is TBD
// SetLength sets the PatternFlowLacpduPartnerLength value in the FlowLacpduPartner object
func (obj *flowLacpduPartner) SetLength(value PatternFlowLacpduPartnerLength) FlowLacpduPartner {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// description is TBD
// SystemPriority returns a PatternFlowLacpduPartnerSystemPriority
func (obj *flowLacpduPartner) SystemPriority() PatternFlowLacpduPartnerSystemPriority {
	if obj.obj.SystemPriority == nil {
		obj.obj.SystemPriority = NewPatternFlowLacpduPartnerSystemPriority().msg()
	}
	if obj.systemPriorityHolder == nil {
		obj.systemPriorityHolder = &patternFlowLacpduPartnerSystemPriority{obj: obj.obj.SystemPriority}
	}
	return obj.systemPriorityHolder
}

// description is TBD
// SystemPriority returns a PatternFlowLacpduPartnerSystemPriority
func (obj *flowLacpduPartner) HasSystemPriority() bool {
	return obj.obj.SystemPriority != nil
}

// description is TBD
// SetSystemPriority sets the PatternFlowLacpduPartnerSystemPriority value in the FlowLacpduPartner object
func (obj *flowLacpduPartner) SetSystemPriority(value PatternFlowLacpduPartnerSystemPriority) FlowLacpduPartner {

	obj.systemPriorityHolder = nil
	obj.obj.SystemPriority = value.msg()

	return obj
}

// description is TBD
// SystemId returns a PatternFlowLacpduPartnerSystemId
func (obj *flowLacpduPartner) SystemId() PatternFlowLacpduPartnerSystemId {
	if obj.obj.SystemId == nil {
		obj.obj.SystemId = NewPatternFlowLacpduPartnerSystemId().msg()
	}
	if obj.systemIdHolder == nil {
		obj.systemIdHolder = &patternFlowLacpduPartnerSystemId{obj: obj.obj.SystemId}
	}
	return obj.systemIdHolder
}

// description is TBD
// SystemId returns a PatternFlowLacpduPartnerSystemId
func (obj *flowLacpduPartner) HasSystemId() bool {
	return obj.obj.SystemId != nil
}

// description is TBD
// SetSystemId sets the PatternFlowLacpduPartnerSystemId value in the FlowLacpduPartner object
func (obj *flowLacpduPartner) SetSystemId(value PatternFlowLacpduPartnerSystemId) FlowLacpduPartner {

	obj.systemIdHolder = nil
	obj.obj.SystemId = value.msg()

	return obj
}

// description is TBD
// Key returns a PatternFlowLacpduPartnerKey
func (obj *flowLacpduPartner) Key() PatternFlowLacpduPartnerKey {
	if obj.obj.Key == nil {
		obj.obj.Key = NewPatternFlowLacpduPartnerKey().msg()
	}
	if obj.keyHolder == nil {
		obj.keyHolder = &patternFlowLacpduPartnerKey{obj: obj.obj.Key}
	}
	return obj.keyHolder
}

// description is TBD
// Key returns a PatternFlowLacpduPartnerKey
func (obj *flowLacpduPartner) HasKey() bool {
	return obj.obj.Key != nil
}

// description is TBD
// SetKey sets the PatternFlowLacpduPartnerKey value in the FlowLacpduPartner object
func (obj *flowLacpduPartner) SetKey(value PatternFlowLacpduPartnerKey) FlowLacpduPartner {

	obj.keyHolder = nil
	obj.obj.Key = value.msg()

	return obj
}

// description is TBD
// PortPriority returns a PatternFlowLacpduPartnerPortPriority
func (obj *flowLacpduPartner) PortPriority() PatternFlowLacpduPartnerPortPriority {
	if obj.obj.PortPriority == nil {
		obj.obj.PortPriority = NewPatternFlowLacpduPartnerPortPriority().msg()
	}
	if obj.portPriorityHolder == nil {
		obj.portPriorityHolder = &patternFlowLacpduPartnerPortPriority{obj: obj.obj.PortPriority}
	}
	return obj.portPriorityHolder
}

// description is TBD
// PortPriority returns a PatternFlowLacpduPartnerPortPriority
func (obj *flowLacpduPartner) HasPortPriority() bool {
	return obj.obj.PortPriority != nil
}

// description is TBD
// SetPortPriority sets the PatternFlowLacpduPartnerPortPriority value in the FlowLacpduPartner object
func (obj *flowLacpduPartner) SetPortPriority(value PatternFlowLacpduPartnerPortPriority) FlowLacpduPartner {

	obj.portPriorityHolder = nil
	obj.obj.PortPriority = value.msg()

	return obj
}

// description is TBD
// PortNumber returns a PatternFlowLacpduPartnerPortNumber
func (obj *flowLacpduPartner) PortNumber() PatternFlowLacpduPartnerPortNumber {
	if obj.obj.PortNumber == nil {
		obj.obj.PortNumber = NewPatternFlowLacpduPartnerPortNumber().msg()
	}
	if obj.portNumberHolder == nil {
		obj.portNumberHolder = &patternFlowLacpduPartnerPortNumber{obj: obj.obj.PortNumber}
	}
	return obj.portNumberHolder
}

// description is TBD
// PortNumber returns a PatternFlowLacpduPartnerPortNumber
func (obj *flowLacpduPartner) HasPortNumber() bool {
	return obj.obj.PortNumber != nil
}

// description is TBD
// SetPortNumber sets the PatternFlowLacpduPartnerPortNumber value in the FlowLacpduPartner object
func (obj *flowLacpduPartner) SetPortNumber(value PatternFlowLacpduPartnerPortNumber) FlowLacpduPartner {

	obj.portNumberHolder = nil
	obj.obj.PortNumber = value.msg()

	return obj
}

// description is TBD
// PartnerState returns a FlowLacpduPartnerState
func (obj *flowLacpduPartner) PartnerState() FlowLacpduPartnerState {
	if obj.obj.PartnerState == nil {
		obj.obj.PartnerState = NewFlowLacpduPartnerState().msg()
	}
	if obj.partnerStateHolder == nil {
		obj.partnerStateHolder = &flowLacpduPartnerState{obj: obj.obj.PartnerState}
	}
	return obj.partnerStateHolder
}

// description is TBD
// PartnerState returns a FlowLacpduPartnerState
func (obj *flowLacpduPartner) HasPartnerState() bool {
	return obj.obj.PartnerState != nil
}

// description is TBD
// SetPartnerState sets the FlowLacpduPartnerState value in the FlowLacpduPartner object
func (obj *flowLacpduPartner) SetPartnerState(value FlowLacpduPartnerState) FlowLacpduPartner {

	obj.partnerStateHolder = nil
	obj.obj.PartnerState = value.msg()

	return obj
}

// description is TBD
// Reserved returns a PatternFlowLacpduPartnerReserved
func (obj *flowLacpduPartner) Reserved() PatternFlowLacpduPartnerReserved {
	if obj.obj.Reserved == nil {
		obj.obj.Reserved = NewPatternFlowLacpduPartnerReserved().msg()
	}
	if obj.reservedHolder == nil {
		obj.reservedHolder = &patternFlowLacpduPartnerReserved{obj: obj.obj.Reserved}
	}
	return obj.reservedHolder
}

// description is TBD
// Reserved returns a PatternFlowLacpduPartnerReserved
func (obj *flowLacpduPartner) HasReserved() bool {
	return obj.obj.Reserved != nil
}

// description is TBD
// SetReserved sets the PatternFlowLacpduPartnerReserved value in the FlowLacpduPartner object
func (obj *flowLacpduPartner) SetReserved(value PatternFlowLacpduPartnerReserved) FlowLacpduPartner {

	obj.reservedHolder = nil
	obj.obj.Reserved = value.msg()

	return obj
}

func (obj *flowLacpduPartner) validateObj(vObj *validation, set_default bool) {
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

func (obj *flowLacpduPartner) setDefault() {

}
