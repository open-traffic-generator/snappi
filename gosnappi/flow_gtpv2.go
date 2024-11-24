package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowGtpv2 *****
type flowGtpv2 struct {
	validation
	obj                    *otg.FlowGtpv2
	marshaller             marshalFlowGtpv2
	unMarshaller           unMarshalFlowGtpv2
	versionHolder          PatternFlowGtpv2Version
	piggybackingFlagHolder PatternFlowGtpv2PiggybackingFlag
	teidFlagHolder         PatternFlowGtpv2TeidFlag
	spare1Holder           PatternFlowGtpv2Spare1
	messageTypeHolder      PatternFlowGtpv2MessageType
	messageLengthHolder    PatternFlowGtpv2MessageLength
	teidHolder             PatternFlowGtpv2Teid
	sequenceNumberHolder   PatternFlowGtpv2SequenceNumber
	spare2Holder           PatternFlowGtpv2Spare2
}

func NewFlowGtpv2() FlowGtpv2 {
	obj := flowGtpv2{obj: &otg.FlowGtpv2{}}
	obj.setDefault()
	return &obj
}

func (obj *flowGtpv2) msg() *otg.FlowGtpv2 {
	return obj.obj
}

func (obj *flowGtpv2) setMsg(msg *otg.FlowGtpv2) FlowGtpv2 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowGtpv2 struct {
	obj *flowGtpv2
}

type marshalFlowGtpv2 interface {
	// ToProto marshals FlowGtpv2 to protobuf object *otg.FlowGtpv2
	ToProto() (*otg.FlowGtpv2, error)
	// ToPbText marshals FlowGtpv2 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowGtpv2 to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowGtpv2 to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowGtpv2 to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowGtpv2 struct {
	obj *flowGtpv2
}

type unMarshalFlowGtpv2 interface {
	// FromProto unmarshals FlowGtpv2 from protobuf object *otg.FlowGtpv2
	FromProto(msg *otg.FlowGtpv2) (FlowGtpv2, error)
	// FromPbText unmarshals FlowGtpv2 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowGtpv2 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowGtpv2 from JSON text
	FromJson(value string) error
}

func (obj *flowGtpv2) Marshal() marshalFlowGtpv2 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowGtpv2{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowGtpv2) Unmarshal() unMarshalFlowGtpv2 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowGtpv2{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowGtpv2) ToProto() (*otg.FlowGtpv2, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowGtpv2) FromProto(msg *otg.FlowGtpv2) (FlowGtpv2, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowGtpv2) ToPbText() (string, error) {
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

func (m *unMarshalflowGtpv2) FromPbText(value string) error {
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

func (m *marshalflowGtpv2) ToYaml() (string, error) {
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

func (m *unMarshalflowGtpv2) FromYaml(value string) error {
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

func (m *marshalflowGtpv2) ToJsonRaw() (string, error) {
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

func (m *marshalflowGtpv2) ToJson() (string, error) {
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

func (m *unMarshalflowGtpv2) FromJson(value string) error {
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

func (obj *flowGtpv2) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowGtpv2) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowGtpv2) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowGtpv2) Clone() (FlowGtpv2, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowGtpv2()
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

func (obj *flowGtpv2) setNil() {
	obj.versionHolder = nil
	obj.piggybackingFlagHolder = nil
	obj.teidFlagHolder = nil
	obj.spare1Holder = nil
	obj.messageTypeHolder = nil
	obj.messageLengthHolder = nil
	obj.teidHolder = nil
	obj.sequenceNumberHolder = nil
	obj.spare2Holder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowGtpv2 is gTPv2 packet header
type FlowGtpv2 interface {
	Validation
	// msg marshals FlowGtpv2 to protobuf object *otg.FlowGtpv2
	// and doesn't set defaults
	msg() *otg.FlowGtpv2
	// setMsg unmarshals FlowGtpv2 from protobuf object *otg.FlowGtpv2
	// and doesn't set defaults
	setMsg(*otg.FlowGtpv2) FlowGtpv2
	// provides marshal interface
	Marshal() marshalFlowGtpv2
	// provides unmarshal interface
	Unmarshal() unMarshalFlowGtpv2
	// validate validates FlowGtpv2
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowGtpv2, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Version returns PatternFlowGtpv2Version, set in FlowGtpv2.
	// PatternFlowGtpv2Version is version number
	Version() PatternFlowGtpv2Version
	// SetVersion assigns PatternFlowGtpv2Version provided by user to FlowGtpv2.
	// PatternFlowGtpv2Version is version number
	SetVersion(value PatternFlowGtpv2Version) FlowGtpv2
	// HasVersion checks if Version has been set in FlowGtpv2
	HasVersion() bool
	// PiggybackingFlag returns PatternFlowGtpv2PiggybackingFlag, set in FlowGtpv2.
	// PatternFlowGtpv2PiggybackingFlag is if piggybacking_flag is set to 1 then another GTP-C message with its own header shall be present at the end of the current message
	PiggybackingFlag() PatternFlowGtpv2PiggybackingFlag
	// SetPiggybackingFlag assigns PatternFlowGtpv2PiggybackingFlag provided by user to FlowGtpv2.
	// PatternFlowGtpv2PiggybackingFlag is if piggybacking_flag is set to 1 then another GTP-C message with its own header shall be present at the end of the current message
	SetPiggybackingFlag(value PatternFlowGtpv2PiggybackingFlag) FlowGtpv2
	// HasPiggybackingFlag checks if PiggybackingFlag has been set in FlowGtpv2
	HasPiggybackingFlag() bool
	// TeidFlag returns PatternFlowGtpv2TeidFlag, set in FlowGtpv2.
	// PatternFlowGtpv2TeidFlag is if teid_flag is set to 1 then the TEID field will be present  between the message length and the sequence number. All messages except Echo and Echo reply require TEID to be present
	TeidFlag() PatternFlowGtpv2TeidFlag
	// SetTeidFlag assigns PatternFlowGtpv2TeidFlag provided by user to FlowGtpv2.
	// PatternFlowGtpv2TeidFlag is if teid_flag is set to 1 then the TEID field will be present  between the message length and the sequence number. All messages except Echo and Echo reply require TEID to be present
	SetTeidFlag(value PatternFlowGtpv2TeidFlag) FlowGtpv2
	// HasTeidFlag checks if TeidFlag has been set in FlowGtpv2
	HasTeidFlag() bool
	// Spare1 returns PatternFlowGtpv2Spare1, set in FlowGtpv2.
	// PatternFlowGtpv2Spare1 is a 3-bit reserved field (must be 0).
	Spare1() PatternFlowGtpv2Spare1
	// SetSpare1 assigns PatternFlowGtpv2Spare1 provided by user to FlowGtpv2.
	// PatternFlowGtpv2Spare1 is a 3-bit reserved field (must be 0).
	SetSpare1(value PatternFlowGtpv2Spare1) FlowGtpv2
	// HasSpare1 checks if Spare1 has been set in FlowGtpv2
	HasSpare1() bool
	// MessageType returns PatternFlowGtpv2MessageType, set in FlowGtpv2.
	// PatternFlowGtpv2MessageType is an 8-bit field that indicates the type of GTP message. Different types of messages are defined in 3GPP TS 29.060 section 7.1
	MessageType() PatternFlowGtpv2MessageType
	// SetMessageType assigns PatternFlowGtpv2MessageType provided by user to FlowGtpv2.
	// PatternFlowGtpv2MessageType is an 8-bit field that indicates the type of GTP message. Different types of messages are defined in 3GPP TS 29.060 section 7.1
	SetMessageType(value PatternFlowGtpv2MessageType) FlowGtpv2
	// HasMessageType checks if MessageType has been set in FlowGtpv2
	HasMessageType() bool
	// MessageLength returns PatternFlowGtpv2MessageLength, set in FlowGtpv2.
	// PatternFlowGtpv2MessageLength is a 16-bit field that indicates the length of the payload in bytes, excluding the mandatory GTP-c header (first 4 bytes). Includes the TEID and sequence_number if they are present.
	MessageLength() PatternFlowGtpv2MessageLength
	// SetMessageLength assigns PatternFlowGtpv2MessageLength provided by user to FlowGtpv2.
	// PatternFlowGtpv2MessageLength is a 16-bit field that indicates the length of the payload in bytes, excluding the mandatory GTP-c header (first 4 bytes). Includes the TEID and sequence_number if they are present.
	SetMessageLength(value PatternFlowGtpv2MessageLength) FlowGtpv2
	// HasMessageLength checks if MessageLength has been set in FlowGtpv2
	HasMessageLength() bool
	// Teid returns PatternFlowGtpv2Teid, set in FlowGtpv2.
	// PatternFlowGtpv2Teid is tunnel endpoint identifier. A 32-bit (4-octet) field used to multiplex different connections in the same GTP tunnel. Is present only if the teid_flag is set.
	Teid() PatternFlowGtpv2Teid
	// SetTeid assigns PatternFlowGtpv2Teid provided by user to FlowGtpv2.
	// PatternFlowGtpv2Teid is tunnel endpoint identifier. A 32-bit (4-octet) field used to multiplex different connections in the same GTP tunnel. Is present only if the teid_flag is set.
	SetTeid(value PatternFlowGtpv2Teid) FlowGtpv2
	// HasTeid checks if Teid has been set in FlowGtpv2
	HasTeid() bool
	// SequenceNumber returns PatternFlowGtpv2SequenceNumber, set in FlowGtpv2.
	// PatternFlowGtpv2SequenceNumber is the sequence number
	SequenceNumber() PatternFlowGtpv2SequenceNumber
	// SetSequenceNumber assigns PatternFlowGtpv2SequenceNumber provided by user to FlowGtpv2.
	// PatternFlowGtpv2SequenceNumber is the sequence number
	SetSequenceNumber(value PatternFlowGtpv2SequenceNumber) FlowGtpv2
	// HasSequenceNumber checks if SequenceNumber has been set in FlowGtpv2
	HasSequenceNumber() bool
	// Spare2 returns PatternFlowGtpv2Spare2, set in FlowGtpv2.
	// PatternFlowGtpv2Spare2 is reserved field
	Spare2() PatternFlowGtpv2Spare2
	// SetSpare2 assigns PatternFlowGtpv2Spare2 provided by user to FlowGtpv2.
	// PatternFlowGtpv2Spare2 is reserved field
	SetSpare2(value PatternFlowGtpv2Spare2) FlowGtpv2
	// HasSpare2 checks if Spare2 has been set in FlowGtpv2
	HasSpare2() bool
	setNil()
}

// description is TBD
// Version returns a PatternFlowGtpv2Version
func (obj *flowGtpv2) Version() PatternFlowGtpv2Version {
	if obj.obj.Version == nil {
		obj.obj.Version = NewPatternFlowGtpv2Version().msg()
	}
	if obj.versionHolder == nil {
		obj.versionHolder = &patternFlowGtpv2Version{obj: obj.obj.Version}
	}
	return obj.versionHolder
}

// description is TBD
// Version returns a PatternFlowGtpv2Version
func (obj *flowGtpv2) HasVersion() bool {
	return obj.obj.Version != nil
}

// description is TBD
// SetVersion sets the PatternFlowGtpv2Version value in the FlowGtpv2 object
func (obj *flowGtpv2) SetVersion(value PatternFlowGtpv2Version) FlowGtpv2 {

	obj.versionHolder = nil
	obj.obj.Version = value.msg()

	return obj
}

// description is TBD
// PiggybackingFlag returns a PatternFlowGtpv2PiggybackingFlag
func (obj *flowGtpv2) PiggybackingFlag() PatternFlowGtpv2PiggybackingFlag {
	if obj.obj.PiggybackingFlag == nil {
		obj.obj.PiggybackingFlag = NewPatternFlowGtpv2PiggybackingFlag().msg()
	}
	if obj.piggybackingFlagHolder == nil {
		obj.piggybackingFlagHolder = &patternFlowGtpv2PiggybackingFlag{obj: obj.obj.PiggybackingFlag}
	}
	return obj.piggybackingFlagHolder
}

// description is TBD
// PiggybackingFlag returns a PatternFlowGtpv2PiggybackingFlag
func (obj *flowGtpv2) HasPiggybackingFlag() bool {
	return obj.obj.PiggybackingFlag != nil
}

// description is TBD
// SetPiggybackingFlag sets the PatternFlowGtpv2PiggybackingFlag value in the FlowGtpv2 object
func (obj *flowGtpv2) SetPiggybackingFlag(value PatternFlowGtpv2PiggybackingFlag) FlowGtpv2 {

	obj.piggybackingFlagHolder = nil
	obj.obj.PiggybackingFlag = value.msg()

	return obj
}

// description is TBD
// TeidFlag returns a PatternFlowGtpv2TeidFlag
func (obj *flowGtpv2) TeidFlag() PatternFlowGtpv2TeidFlag {
	if obj.obj.TeidFlag == nil {
		obj.obj.TeidFlag = NewPatternFlowGtpv2TeidFlag().msg()
	}
	if obj.teidFlagHolder == nil {
		obj.teidFlagHolder = &patternFlowGtpv2TeidFlag{obj: obj.obj.TeidFlag}
	}
	return obj.teidFlagHolder
}

// description is TBD
// TeidFlag returns a PatternFlowGtpv2TeidFlag
func (obj *flowGtpv2) HasTeidFlag() bool {
	return obj.obj.TeidFlag != nil
}

// description is TBD
// SetTeidFlag sets the PatternFlowGtpv2TeidFlag value in the FlowGtpv2 object
func (obj *flowGtpv2) SetTeidFlag(value PatternFlowGtpv2TeidFlag) FlowGtpv2 {

	obj.teidFlagHolder = nil
	obj.obj.TeidFlag = value.msg()

	return obj
}

// description is TBD
// Spare1 returns a PatternFlowGtpv2Spare1
func (obj *flowGtpv2) Spare1() PatternFlowGtpv2Spare1 {
	if obj.obj.Spare1 == nil {
		obj.obj.Spare1 = NewPatternFlowGtpv2Spare1().msg()
	}
	if obj.spare1Holder == nil {
		obj.spare1Holder = &patternFlowGtpv2Spare1{obj: obj.obj.Spare1}
	}
	return obj.spare1Holder
}

// description is TBD
// Spare1 returns a PatternFlowGtpv2Spare1
func (obj *flowGtpv2) HasSpare1() bool {
	return obj.obj.Spare1 != nil
}

// description is TBD
// SetSpare1 sets the PatternFlowGtpv2Spare1 value in the FlowGtpv2 object
func (obj *flowGtpv2) SetSpare1(value PatternFlowGtpv2Spare1) FlowGtpv2 {

	obj.spare1Holder = nil
	obj.obj.Spare1 = value.msg()

	return obj
}

// description is TBD
// MessageType returns a PatternFlowGtpv2MessageType
func (obj *flowGtpv2) MessageType() PatternFlowGtpv2MessageType {
	if obj.obj.MessageType == nil {
		obj.obj.MessageType = NewPatternFlowGtpv2MessageType().msg()
	}
	if obj.messageTypeHolder == nil {
		obj.messageTypeHolder = &patternFlowGtpv2MessageType{obj: obj.obj.MessageType}
	}
	return obj.messageTypeHolder
}

// description is TBD
// MessageType returns a PatternFlowGtpv2MessageType
func (obj *flowGtpv2) HasMessageType() bool {
	return obj.obj.MessageType != nil
}

// description is TBD
// SetMessageType sets the PatternFlowGtpv2MessageType value in the FlowGtpv2 object
func (obj *flowGtpv2) SetMessageType(value PatternFlowGtpv2MessageType) FlowGtpv2 {

	obj.messageTypeHolder = nil
	obj.obj.MessageType = value.msg()

	return obj
}

// description is TBD
// MessageLength returns a PatternFlowGtpv2MessageLength
func (obj *flowGtpv2) MessageLength() PatternFlowGtpv2MessageLength {
	if obj.obj.MessageLength == nil {
		obj.obj.MessageLength = NewPatternFlowGtpv2MessageLength().msg()
	}
	if obj.messageLengthHolder == nil {
		obj.messageLengthHolder = &patternFlowGtpv2MessageLength{obj: obj.obj.MessageLength}
	}
	return obj.messageLengthHolder
}

// description is TBD
// MessageLength returns a PatternFlowGtpv2MessageLength
func (obj *flowGtpv2) HasMessageLength() bool {
	return obj.obj.MessageLength != nil
}

// description is TBD
// SetMessageLength sets the PatternFlowGtpv2MessageLength value in the FlowGtpv2 object
func (obj *flowGtpv2) SetMessageLength(value PatternFlowGtpv2MessageLength) FlowGtpv2 {

	obj.messageLengthHolder = nil
	obj.obj.MessageLength = value.msg()

	return obj
}

// description is TBD
// Teid returns a PatternFlowGtpv2Teid
func (obj *flowGtpv2) Teid() PatternFlowGtpv2Teid {
	if obj.obj.Teid == nil {
		obj.obj.Teid = NewPatternFlowGtpv2Teid().msg()
	}
	if obj.teidHolder == nil {
		obj.teidHolder = &patternFlowGtpv2Teid{obj: obj.obj.Teid}
	}
	return obj.teidHolder
}

// description is TBD
// Teid returns a PatternFlowGtpv2Teid
func (obj *flowGtpv2) HasTeid() bool {
	return obj.obj.Teid != nil
}

// description is TBD
// SetTeid sets the PatternFlowGtpv2Teid value in the FlowGtpv2 object
func (obj *flowGtpv2) SetTeid(value PatternFlowGtpv2Teid) FlowGtpv2 {

	obj.teidHolder = nil
	obj.obj.Teid = value.msg()

	return obj
}

// description is TBD
// SequenceNumber returns a PatternFlowGtpv2SequenceNumber
func (obj *flowGtpv2) SequenceNumber() PatternFlowGtpv2SequenceNumber {
	if obj.obj.SequenceNumber == nil {
		obj.obj.SequenceNumber = NewPatternFlowGtpv2SequenceNumber().msg()
	}
	if obj.sequenceNumberHolder == nil {
		obj.sequenceNumberHolder = &patternFlowGtpv2SequenceNumber{obj: obj.obj.SequenceNumber}
	}
	return obj.sequenceNumberHolder
}

// description is TBD
// SequenceNumber returns a PatternFlowGtpv2SequenceNumber
func (obj *flowGtpv2) HasSequenceNumber() bool {
	return obj.obj.SequenceNumber != nil
}

// description is TBD
// SetSequenceNumber sets the PatternFlowGtpv2SequenceNumber value in the FlowGtpv2 object
func (obj *flowGtpv2) SetSequenceNumber(value PatternFlowGtpv2SequenceNumber) FlowGtpv2 {

	obj.sequenceNumberHolder = nil
	obj.obj.SequenceNumber = value.msg()

	return obj
}

// description is TBD
// Spare2 returns a PatternFlowGtpv2Spare2
func (obj *flowGtpv2) Spare2() PatternFlowGtpv2Spare2 {
	if obj.obj.Spare2 == nil {
		obj.obj.Spare2 = NewPatternFlowGtpv2Spare2().msg()
	}
	if obj.spare2Holder == nil {
		obj.spare2Holder = &patternFlowGtpv2Spare2{obj: obj.obj.Spare2}
	}
	return obj.spare2Holder
}

// description is TBD
// Spare2 returns a PatternFlowGtpv2Spare2
func (obj *flowGtpv2) HasSpare2() bool {
	return obj.obj.Spare2 != nil
}

// description is TBD
// SetSpare2 sets the PatternFlowGtpv2Spare2 value in the FlowGtpv2 object
func (obj *flowGtpv2) SetSpare2(value PatternFlowGtpv2Spare2) FlowGtpv2 {

	obj.spare2Holder = nil
	obj.obj.Spare2 = value.msg()

	return obj
}

func (obj *flowGtpv2) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Version != nil {

		obj.Version().validateObj(vObj, set_default)
	}

	if obj.obj.PiggybackingFlag != nil {

		obj.PiggybackingFlag().validateObj(vObj, set_default)
	}

	if obj.obj.TeidFlag != nil {

		obj.TeidFlag().validateObj(vObj, set_default)
	}

	if obj.obj.Spare1 != nil {

		obj.Spare1().validateObj(vObj, set_default)
	}

	if obj.obj.MessageType != nil {

		obj.MessageType().validateObj(vObj, set_default)
	}

	if obj.obj.MessageLength != nil {

		obj.MessageLength().validateObj(vObj, set_default)
	}

	if obj.obj.Teid != nil {

		obj.Teid().validateObj(vObj, set_default)
	}

	if obj.obj.SequenceNumber != nil {

		obj.SequenceNumber().validateObj(vObj, set_default)
	}

	if obj.obj.Spare2 != nil {

		obj.Spare2().validateObj(vObj, set_default)
	}

}

func (obj *flowGtpv2) setDefault() {

}
