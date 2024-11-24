package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowArp *****
type flowArp struct {
	validation
	obj                      *otg.FlowArp
	marshaller               marshalFlowArp
	unMarshaller             unMarshalFlowArp
	hardwareTypeHolder       PatternFlowArpHardwareType
	protocolTypeHolder       PatternFlowArpProtocolType
	hardwareLengthHolder     PatternFlowArpHardwareLength
	protocolLengthHolder     PatternFlowArpProtocolLength
	operationHolder          PatternFlowArpOperation
	senderHardwareAddrHolder PatternFlowArpSenderHardwareAddr
	senderProtocolAddrHolder PatternFlowArpSenderProtocolAddr
	targetHardwareAddrHolder PatternFlowArpTargetHardwareAddr
	targetProtocolAddrHolder PatternFlowArpTargetProtocolAddr
}

func NewFlowArp() FlowArp {
	obj := flowArp{obj: &otg.FlowArp{}}
	obj.setDefault()
	return &obj
}

func (obj *flowArp) msg() *otg.FlowArp {
	return obj.obj
}

func (obj *flowArp) setMsg(msg *otg.FlowArp) FlowArp {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowArp struct {
	obj *flowArp
}

type marshalFlowArp interface {
	// ToProto marshals FlowArp to protobuf object *otg.FlowArp
	ToProto() (*otg.FlowArp, error)
	// ToPbText marshals FlowArp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowArp to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowArp to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowArp to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowArp struct {
	obj *flowArp
}

type unMarshalFlowArp interface {
	// FromProto unmarshals FlowArp from protobuf object *otg.FlowArp
	FromProto(msg *otg.FlowArp) (FlowArp, error)
	// FromPbText unmarshals FlowArp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowArp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowArp from JSON text
	FromJson(value string) error
}

func (obj *flowArp) Marshal() marshalFlowArp {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowArp{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowArp) Unmarshal() unMarshalFlowArp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowArp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowArp) ToProto() (*otg.FlowArp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowArp) FromProto(msg *otg.FlowArp) (FlowArp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowArp) ToPbText() (string, error) {
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

func (m *unMarshalflowArp) FromPbText(value string) error {
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

func (m *marshalflowArp) ToYaml() (string, error) {
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

func (m *unMarshalflowArp) FromYaml(value string) error {
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

func (m *marshalflowArp) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalflowArp) ToJson() (string, error) {
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

func (m *unMarshalflowArp) FromJson(value string) error {
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

func (obj *flowArp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowArp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowArp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowArp) Clone() (FlowArp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowArp()
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

func (obj *flowArp) setNil() {
	obj.hardwareTypeHolder = nil
	obj.protocolTypeHolder = nil
	obj.hardwareLengthHolder = nil
	obj.protocolLengthHolder = nil
	obj.operationHolder = nil
	obj.senderHardwareAddrHolder = nil
	obj.senderProtocolAddrHolder = nil
	obj.targetHardwareAddrHolder = nil
	obj.targetProtocolAddrHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowArp is aRP packet header
type FlowArp interface {
	Validation
	// msg marshals FlowArp to protobuf object *otg.FlowArp
	// and doesn't set defaults
	msg() *otg.FlowArp
	// setMsg unmarshals FlowArp from protobuf object *otg.FlowArp
	// and doesn't set defaults
	setMsg(*otg.FlowArp) FlowArp
	// provides marshal interface
	Marshal() marshalFlowArp
	// provides unmarshal interface
	Unmarshal() unMarshalFlowArp
	// validate validates FlowArp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowArp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// HardwareType returns PatternFlowArpHardwareType, set in FlowArp.
	// PatternFlowArpHardwareType is network link protocol type
	HardwareType() PatternFlowArpHardwareType
	// SetHardwareType assigns PatternFlowArpHardwareType provided by user to FlowArp.
	// PatternFlowArpHardwareType is network link protocol type
	SetHardwareType(value PatternFlowArpHardwareType) FlowArp
	// HasHardwareType checks if HardwareType has been set in FlowArp
	HasHardwareType() bool
	// ProtocolType returns PatternFlowArpProtocolType, set in FlowArp.
	// PatternFlowArpProtocolType is the internetwork protocol for which the ARP request is intended
	ProtocolType() PatternFlowArpProtocolType
	// SetProtocolType assigns PatternFlowArpProtocolType provided by user to FlowArp.
	// PatternFlowArpProtocolType is the internetwork protocol for which the ARP request is intended
	SetProtocolType(value PatternFlowArpProtocolType) FlowArp
	// HasProtocolType checks if ProtocolType has been set in FlowArp
	HasProtocolType() bool
	// HardwareLength returns PatternFlowArpHardwareLength, set in FlowArp.
	// PatternFlowArpHardwareLength is length (in octets) of a hardware address
	HardwareLength() PatternFlowArpHardwareLength
	// SetHardwareLength assigns PatternFlowArpHardwareLength provided by user to FlowArp.
	// PatternFlowArpHardwareLength is length (in octets) of a hardware address
	SetHardwareLength(value PatternFlowArpHardwareLength) FlowArp
	// HasHardwareLength checks if HardwareLength has been set in FlowArp
	HasHardwareLength() bool
	// ProtocolLength returns PatternFlowArpProtocolLength, set in FlowArp.
	// PatternFlowArpProtocolLength is length (in octets) of internetwork addresses
	ProtocolLength() PatternFlowArpProtocolLength
	// SetProtocolLength assigns PatternFlowArpProtocolLength provided by user to FlowArp.
	// PatternFlowArpProtocolLength is length (in octets) of internetwork addresses
	SetProtocolLength(value PatternFlowArpProtocolLength) FlowArp
	// HasProtocolLength checks if ProtocolLength has been set in FlowArp
	HasProtocolLength() bool
	// Operation returns PatternFlowArpOperation, set in FlowArp.
	// PatternFlowArpOperation is the operation that the sender is performing
	Operation() PatternFlowArpOperation
	// SetOperation assigns PatternFlowArpOperation provided by user to FlowArp.
	// PatternFlowArpOperation is the operation that the sender is performing
	SetOperation(value PatternFlowArpOperation) FlowArp
	// HasOperation checks if Operation has been set in FlowArp
	HasOperation() bool
	// SenderHardwareAddr returns PatternFlowArpSenderHardwareAddr, set in FlowArp.
	// PatternFlowArpSenderHardwareAddr is media address of the sender
	SenderHardwareAddr() PatternFlowArpSenderHardwareAddr
	// SetSenderHardwareAddr assigns PatternFlowArpSenderHardwareAddr provided by user to FlowArp.
	// PatternFlowArpSenderHardwareAddr is media address of the sender
	SetSenderHardwareAddr(value PatternFlowArpSenderHardwareAddr) FlowArp
	// HasSenderHardwareAddr checks if SenderHardwareAddr has been set in FlowArp
	HasSenderHardwareAddr() bool
	// SenderProtocolAddr returns PatternFlowArpSenderProtocolAddr, set in FlowArp.
	// PatternFlowArpSenderProtocolAddr is internetwork address of the sender
	SenderProtocolAddr() PatternFlowArpSenderProtocolAddr
	// SetSenderProtocolAddr assigns PatternFlowArpSenderProtocolAddr provided by user to FlowArp.
	// PatternFlowArpSenderProtocolAddr is internetwork address of the sender
	SetSenderProtocolAddr(value PatternFlowArpSenderProtocolAddr) FlowArp
	// HasSenderProtocolAddr checks if SenderProtocolAddr has been set in FlowArp
	HasSenderProtocolAddr() bool
	// TargetHardwareAddr returns PatternFlowArpTargetHardwareAddr, set in FlowArp.
	// PatternFlowArpTargetHardwareAddr is media address of the target
	TargetHardwareAddr() PatternFlowArpTargetHardwareAddr
	// SetTargetHardwareAddr assigns PatternFlowArpTargetHardwareAddr provided by user to FlowArp.
	// PatternFlowArpTargetHardwareAddr is media address of the target
	SetTargetHardwareAddr(value PatternFlowArpTargetHardwareAddr) FlowArp
	// HasTargetHardwareAddr checks if TargetHardwareAddr has been set in FlowArp
	HasTargetHardwareAddr() bool
	// TargetProtocolAddr returns PatternFlowArpTargetProtocolAddr, set in FlowArp.
	// PatternFlowArpTargetProtocolAddr is internetwork address of the target
	TargetProtocolAddr() PatternFlowArpTargetProtocolAddr
	// SetTargetProtocolAddr assigns PatternFlowArpTargetProtocolAddr provided by user to FlowArp.
	// PatternFlowArpTargetProtocolAddr is internetwork address of the target
	SetTargetProtocolAddr(value PatternFlowArpTargetProtocolAddr) FlowArp
	// HasTargetProtocolAddr checks if TargetProtocolAddr has been set in FlowArp
	HasTargetProtocolAddr() bool
	setNil()
}

// description is TBD
// HardwareType returns a PatternFlowArpHardwareType
func (obj *flowArp) HardwareType() PatternFlowArpHardwareType {
	if obj.obj.HardwareType == nil {
		obj.obj.HardwareType = NewPatternFlowArpHardwareType().msg()
	}
	if obj.hardwareTypeHolder == nil {
		obj.hardwareTypeHolder = &patternFlowArpHardwareType{obj: obj.obj.HardwareType}
	}
	return obj.hardwareTypeHolder
}

// description is TBD
// HardwareType returns a PatternFlowArpHardwareType
func (obj *flowArp) HasHardwareType() bool {
	return obj.obj.HardwareType != nil
}

// description is TBD
// SetHardwareType sets the PatternFlowArpHardwareType value in the FlowArp object
func (obj *flowArp) SetHardwareType(value PatternFlowArpHardwareType) FlowArp {

	obj.hardwareTypeHolder = nil
	obj.obj.HardwareType = value.msg()

	return obj
}

// description is TBD
// ProtocolType returns a PatternFlowArpProtocolType
func (obj *flowArp) ProtocolType() PatternFlowArpProtocolType {
	if obj.obj.ProtocolType == nil {
		obj.obj.ProtocolType = NewPatternFlowArpProtocolType().msg()
	}
	if obj.protocolTypeHolder == nil {
		obj.protocolTypeHolder = &patternFlowArpProtocolType{obj: obj.obj.ProtocolType}
	}
	return obj.protocolTypeHolder
}

// description is TBD
// ProtocolType returns a PatternFlowArpProtocolType
func (obj *flowArp) HasProtocolType() bool {
	return obj.obj.ProtocolType != nil
}

// description is TBD
// SetProtocolType sets the PatternFlowArpProtocolType value in the FlowArp object
func (obj *flowArp) SetProtocolType(value PatternFlowArpProtocolType) FlowArp {

	obj.protocolTypeHolder = nil
	obj.obj.ProtocolType = value.msg()

	return obj
}

// description is TBD
// HardwareLength returns a PatternFlowArpHardwareLength
func (obj *flowArp) HardwareLength() PatternFlowArpHardwareLength {
	if obj.obj.HardwareLength == nil {
		obj.obj.HardwareLength = NewPatternFlowArpHardwareLength().msg()
	}
	if obj.hardwareLengthHolder == nil {
		obj.hardwareLengthHolder = &patternFlowArpHardwareLength{obj: obj.obj.HardwareLength}
	}
	return obj.hardwareLengthHolder
}

// description is TBD
// HardwareLength returns a PatternFlowArpHardwareLength
func (obj *flowArp) HasHardwareLength() bool {
	return obj.obj.HardwareLength != nil
}

// description is TBD
// SetHardwareLength sets the PatternFlowArpHardwareLength value in the FlowArp object
func (obj *flowArp) SetHardwareLength(value PatternFlowArpHardwareLength) FlowArp {

	obj.hardwareLengthHolder = nil
	obj.obj.HardwareLength = value.msg()

	return obj
}

// description is TBD
// ProtocolLength returns a PatternFlowArpProtocolLength
func (obj *flowArp) ProtocolLength() PatternFlowArpProtocolLength {
	if obj.obj.ProtocolLength == nil {
		obj.obj.ProtocolLength = NewPatternFlowArpProtocolLength().msg()
	}
	if obj.protocolLengthHolder == nil {
		obj.protocolLengthHolder = &patternFlowArpProtocolLength{obj: obj.obj.ProtocolLength}
	}
	return obj.protocolLengthHolder
}

// description is TBD
// ProtocolLength returns a PatternFlowArpProtocolLength
func (obj *flowArp) HasProtocolLength() bool {
	return obj.obj.ProtocolLength != nil
}

// description is TBD
// SetProtocolLength sets the PatternFlowArpProtocolLength value in the FlowArp object
func (obj *flowArp) SetProtocolLength(value PatternFlowArpProtocolLength) FlowArp {

	obj.protocolLengthHolder = nil
	obj.obj.ProtocolLength = value.msg()

	return obj
}

// description is TBD
// Operation returns a PatternFlowArpOperation
func (obj *flowArp) Operation() PatternFlowArpOperation {
	if obj.obj.Operation == nil {
		obj.obj.Operation = NewPatternFlowArpOperation().msg()
	}
	if obj.operationHolder == nil {
		obj.operationHolder = &patternFlowArpOperation{obj: obj.obj.Operation}
	}
	return obj.operationHolder
}

// description is TBD
// Operation returns a PatternFlowArpOperation
func (obj *flowArp) HasOperation() bool {
	return obj.obj.Operation != nil
}

// description is TBD
// SetOperation sets the PatternFlowArpOperation value in the FlowArp object
func (obj *flowArp) SetOperation(value PatternFlowArpOperation) FlowArp {

	obj.operationHolder = nil
	obj.obj.Operation = value.msg()

	return obj
}

// description is TBD
// SenderHardwareAddr returns a PatternFlowArpSenderHardwareAddr
func (obj *flowArp) SenderHardwareAddr() PatternFlowArpSenderHardwareAddr {
	if obj.obj.SenderHardwareAddr == nil {
		obj.obj.SenderHardwareAddr = NewPatternFlowArpSenderHardwareAddr().msg()
	}
	if obj.senderHardwareAddrHolder == nil {
		obj.senderHardwareAddrHolder = &patternFlowArpSenderHardwareAddr{obj: obj.obj.SenderHardwareAddr}
	}
	return obj.senderHardwareAddrHolder
}

// description is TBD
// SenderHardwareAddr returns a PatternFlowArpSenderHardwareAddr
func (obj *flowArp) HasSenderHardwareAddr() bool {
	return obj.obj.SenderHardwareAddr != nil
}

// description is TBD
// SetSenderHardwareAddr sets the PatternFlowArpSenderHardwareAddr value in the FlowArp object
func (obj *flowArp) SetSenderHardwareAddr(value PatternFlowArpSenderHardwareAddr) FlowArp {

	obj.senderHardwareAddrHolder = nil
	obj.obj.SenderHardwareAddr = value.msg()

	return obj
}

// description is TBD
// SenderProtocolAddr returns a PatternFlowArpSenderProtocolAddr
func (obj *flowArp) SenderProtocolAddr() PatternFlowArpSenderProtocolAddr {
	if obj.obj.SenderProtocolAddr == nil {
		obj.obj.SenderProtocolAddr = NewPatternFlowArpSenderProtocolAddr().msg()
	}
	if obj.senderProtocolAddrHolder == nil {
		obj.senderProtocolAddrHolder = &patternFlowArpSenderProtocolAddr{obj: obj.obj.SenderProtocolAddr}
	}
	return obj.senderProtocolAddrHolder
}

// description is TBD
// SenderProtocolAddr returns a PatternFlowArpSenderProtocolAddr
func (obj *flowArp) HasSenderProtocolAddr() bool {
	return obj.obj.SenderProtocolAddr != nil
}

// description is TBD
// SetSenderProtocolAddr sets the PatternFlowArpSenderProtocolAddr value in the FlowArp object
func (obj *flowArp) SetSenderProtocolAddr(value PatternFlowArpSenderProtocolAddr) FlowArp {

	obj.senderProtocolAddrHolder = nil
	obj.obj.SenderProtocolAddr = value.msg()

	return obj
}

// description is TBD
// TargetHardwareAddr returns a PatternFlowArpTargetHardwareAddr
func (obj *flowArp) TargetHardwareAddr() PatternFlowArpTargetHardwareAddr {
	if obj.obj.TargetHardwareAddr == nil {
		obj.obj.TargetHardwareAddr = NewPatternFlowArpTargetHardwareAddr().msg()
	}
	if obj.targetHardwareAddrHolder == nil {
		obj.targetHardwareAddrHolder = &patternFlowArpTargetHardwareAddr{obj: obj.obj.TargetHardwareAddr}
	}
	return obj.targetHardwareAddrHolder
}

// description is TBD
// TargetHardwareAddr returns a PatternFlowArpTargetHardwareAddr
func (obj *flowArp) HasTargetHardwareAddr() bool {
	return obj.obj.TargetHardwareAddr != nil
}

// description is TBD
// SetTargetHardwareAddr sets the PatternFlowArpTargetHardwareAddr value in the FlowArp object
func (obj *flowArp) SetTargetHardwareAddr(value PatternFlowArpTargetHardwareAddr) FlowArp {

	obj.targetHardwareAddrHolder = nil
	obj.obj.TargetHardwareAddr = value.msg()

	return obj
}

// description is TBD
// TargetProtocolAddr returns a PatternFlowArpTargetProtocolAddr
func (obj *flowArp) TargetProtocolAddr() PatternFlowArpTargetProtocolAddr {
	if obj.obj.TargetProtocolAddr == nil {
		obj.obj.TargetProtocolAddr = NewPatternFlowArpTargetProtocolAddr().msg()
	}
	if obj.targetProtocolAddrHolder == nil {
		obj.targetProtocolAddrHolder = &patternFlowArpTargetProtocolAddr{obj: obj.obj.TargetProtocolAddr}
	}
	return obj.targetProtocolAddrHolder
}

// description is TBD
// TargetProtocolAddr returns a PatternFlowArpTargetProtocolAddr
func (obj *flowArp) HasTargetProtocolAddr() bool {
	return obj.obj.TargetProtocolAddr != nil
}

// description is TBD
// SetTargetProtocolAddr sets the PatternFlowArpTargetProtocolAddr value in the FlowArp object
func (obj *flowArp) SetTargetProtocolAddr(value PatternFlowArpTargetProtocolAddr) FlowArp {

	obj.targetProtocolAddrHolder = nil
	obj.obj.TargetProtocolAddr = value.msg()

	return obj
}

func (obj *flowArp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.HardwareType != nil {

		obj.HardwareType().validateObj(vObj, set_default)
	}

	if obj.obj.ProtocolType != nil {

		obj.ProtocolType().validateObj(vObj, set_default)
	}

	if obj.obj.HardwareLength != nil {

		obj.HardwareLength().validateObj(vObj, set_default)
	}

	if obj.obj.ProtocolLength != nil {

		obj.ProtocolLength().validateObj(vObj, set_default)
	}

	if obj.obj.Operation != nil {

		obj.Operation().validateObj(vObj, set_default)
	}

	if obj.obj.SenderHardwareAddr != nil {

		obj.SenderHardwareAddr().validateObj(vObj, set_default)
	}

	if obj.obj.SenderProtocolAddr != nil {

		obj.SenderProtocolAddr().validateObj(vObj, set_default)
	}

	if obj.obj.TargetHardwareAddr != nil {

		obj.TargetHardwareAddr().validateObj(vObj, set_default)
	}

	if obj.obj.TargetProtocolAddr != nil {

		obj.TargetProtocolAddr().validateObj(vObj, set_default)
	}

}

func (obj *flowArp) setDefault() {

}
