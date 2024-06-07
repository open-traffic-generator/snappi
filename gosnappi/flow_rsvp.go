package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRsvp *****
type flowRsvp struct {
	validation
	obj                *otg.FlowRsvp
	marshaller         marshalFlowRsvp
	unMarshaller       unMarshalFlowRsvp
	rsvpChecksumHolder PatternFlowRsvpRsvpChecksum
	timeToLiveHolder   PatternFlowRsvpTimeToLive
	reservedHolder     PatternFlowRsvpReserved
	rsvpLengthHolder   FlowRSVPLength
	messageTypeHolder  FlowRSVPMessage
}

func NewFlowRsvp() FlowRsvp {
	obj := flowRsvp{obj: &otg.FlowRsvp{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRsvp) msg() *otg.FlowRsvp {
	return obj.obj
}

func (obj *flowRsvp) setMsg(msg *otg.FlowRsvp) FlowRsvp {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRsvp struct {
	obj *flowRsvp
}

type marshalFlowRsvp interface {
	// ToProto marshals FlowRsvp to protobuf object *otg.FlowRsvp
	ToProto() (*otg.FlowRsvp, error)
	// ToPbText marshals FlowRsvp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRsvp to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRsvp to JSON text
	ToJson() (string, error)
}

type unMarshalflowRsvp struct {
	obj *flowRsvp
}

type unMarshalFlowRsvp interface {
	// FromProto unmarshals FlowRsvp from protobuf object *otg.FlowRsvp
	FromProto(msg *otg.FlowRsvp) (FlowRsvp, error)
	// FromPbText unmarshals FlowRsvp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRsvp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRsvp from JSON text
	FromJson(value string) error
}

func (obj *flowRsvp) Marshal() marshalFlowRsvp {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRsvp{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRsvp) Unmarshal() unMarshalFlowRsvp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRsvp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRsvp) ToProto() (*otg.FlowRsvp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRsvp) FromProto(msg *otg.FlowRsvp) (FlowRsvp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRsvp) ToPbText() (string, error) {
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

func (m *unMarshalflowRsvp) FromPbText(value string) error {
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

func (m *marshalflowRsvp) ToYaml() (string, error) {
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

func (m *unMarshalflowRsvp) FromYaml(value string) error {
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

func (m *marshalflowRsvp) ToJson() (string, error) {
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

func (m *unMarshalflowRsvp) FromJson(value string) error {
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

func (obj *flowRsvp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRsvp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRsvp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRsvp) Clone() (FlowRsvp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRsvp()
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

func (obj *flowRsvp) setNil() {
	obj.rsvpChecksumHolder = nil
	obj.timeToLiveHolder = nil
	obj.reservedHolder = nil
	obj.rsvpLengthHolder = nil
	obj.messageTypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRsvp is rSVP packet header as defined in RFC2205 and RFC3209. Currently only supported message type is "Path" with mandatory objects and sub-objects.
type FlowRsvp interface {
	Validation
	// msg marshals FlowRsvp to protobuf object *otg.FlowRsvp
	// and doesn't set defaults
	msg() *otg.FlowRsvp
	// setMsg unmarshals FlowRsvp from protobuf object *otg.FlowRsvp
	// and doesn't set defaults
	setMsg(*otg.FlowRsvp) FlowRsvp
	// provides marshal interface
	Marshal() marshalFlowRsvp
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRsvp
	// validate validates FlowRsvp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRsvp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Version returns uint32, set in FlowRsvp.
	Version() uint32
	// SetVersion assigns uint32 provided by user to FlowRsvp
	SetVersion(value uint32) FlowRsvp
	// HasVersion checks if Version has been set in FlowRsvp
	HasVersion() bool
	// Flag returns FlowRsvpFlagEnum, set in FlowRsvp
	Flag() FlowRsvpFlagEnum
	// SetFlag assigns FlowRsvpFlagEnum provided by user to FlowRsvp
	SetFlag(value FlowRsvpFlagEnum) FlowRsvp
	// HasFlag checks if Flag has been set in FlowRsvp
	HasFlag() bool
	// RsvpChecksum returns PatternFlowRsvpRsvpChecksum, set in FlowRsvp.
	// PatternFlowRsvpRsvpChecksum is the one's complement of the one's complement sum of the message, with the checksum field replaced by zero for the purpose of computing the checksum.   An all-zero value means that no checksum was transmitted.
	RsvpChecksum() PatternFlowRsvpRsvpChecksum
	// SetRsvpChecksum assigns PatternFlowRsvpRsvpChecksum provided by user to FlowRsvp.
	// PatternFlowRsvpRsvpChecksum is the one's complement of the one's complement sum of the message, with the checksum field replaced by zero for the purpose of computing the checksum.   An all-zero value means that no checksum was transmitted.
	SetRsvpChecksum(value PatternFlowRsvpRsvpChecksum) FlowRsvp
	// HasRsvpChecksum checks if RsvpChecksum has been set in FlowRsvp
	HasRsvpChecksum() bool
	// TimeToLive returns PatternFlowRsvpTimeToLive, set in FlowRsvp.
	// PatternFlowRsvpTimeToLive is the IP time-to-live(TTL) value with which the message was sent.
	TimeToLive() PatternFlowRsvpTimeToLive
	// SetTimeToLive assigns PatternFlowRsvpTimeToLive provided by user to FlowRsvp.
	// PatternFlowRsvpTimeToLive is the IP time-to-live(TTL) value with which the message was sent.
	SetTimeToLive(value PatternFlowRsvpTimeToLive) FlowRsvp
	// HasTimeToLive checks if TimeToLive has been set in FlowRsvp
	HasTimeToLive() bool
	// Reserved returns PatternFlowRsvpReserved, set in FlowRsvp.
	// PatternFlowRsvpReserved is reserved
	Reserved() PatternFlowRsvpReserved
	// SetReserved assigns PatternFlowRsvpReserved provided by user to FlowRsvp.
	// PatternFlowRsvpReserved is reserved
	SetReserved(value PatternFlowRsvpReserved) FlowRsvp
	// HasReserved checks if Reserved has been set in FlowRsvp
	HasReserved() bool
	// RsvpLength returns FlowRSVPLength, set in FlowRsvp.
	// FlowRSVPLength is description is TBD
	RsvpLength() FlowRSVPLength
	// SetRsvpLength assigns FlowRSVPLength provided by user to FlowRsvp.
	// FlowRSVPLength is description is TBD
	SetRsvpLength(value FlowRSVPLength) FlowRsvp
	// HasRsvpLength checks if RsvpLength has been set in FlowRsvp
	HasRsvpLength() bool
	// MessageType returns FlowRSVPMessage, set in FlowRsvp.
	// FlowRSVPMessage is description is TBD
	MessageType() FlowRSVPMessage
	// SetMessageType assigns FlowRSVPMessage provided by user to FlowRsvp.
	// FlowRSVPMessage is description is TBD
	SetMessageType(value FlowRSVPMessage) FlowRsvp
	// HasMessageType checks if MessageType has been set in FlowRsvp
	HasMessageType() bool
	setNil()
}

// RSVP Protocol Version.
// Version returns a uint32
func (obj *flowRsvp) Version() uint32 {

	return *obj.obj.Version

}

// RSVP Protocol Version.
// Version returns a uint32
func (obj *flowRsvp) HasVersion() bool {
	return obj.obj.Version != nil
}

// RSVP Protocol Version.
// SetVersion sets the uint32 value in the FlowRsvp object
func (obj *flowRsvp) SetVersion(value uint32) FlowRsvp {

	obj.obj.Version = &value
	return obj
}

type FlowRsvpFlagEnum string

// Enum of Flag on FlowRsvp
var FlowRsvpFlag = struct {
	NOT_REFRESH_REDUCTION_CAPABLE FlowRsvpFlagEnum
	REFRESH_REDUCTION_CAPABLE     FlowRsvpFlagEnum
}{
	NOT_REFRESH_REDUCTION_CAPABLE: FlowRsvpFlagEnum("not_refresh_reduction_capable"),
	REFRESH_REDUCTION_CAPABLE:     FlowRsvpFlagEnum("refresh_reduction_capable"),
}

func (obj *flowRsvp) Flag() FlowRsvpFlagEnum {
	return FlowRsvpFlagEnum(obj.obj.Flag.Enum().String())
}

// Flag, 0x01-0x08: Reserved.
// Flag returns a string
func (obj *flowRsvp) HasFlag() bool {
	return obj.obj.Flag != nil
}

func (obj *flowRsvp) SetFlag(value FlowRsvpFlagEnum) FlowRsvp {
	intValue, ok := otg.FlowRsvp_Flag_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowRsvpFlagEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowRsvp_Flag_Enum(intValue)
	obj.obj.Flag = &enumValue

	return obj
}

// description is TBD
// RsvpChecksum returns a PatternFlowRsvpRsvpChecksum
func (obj *flowRsvp) RsvpChecksum() PatternFlowRsvpRsvpChecksum {
	if obj.obj.RsvpChecksum == nil {
		obj.obj.RsvpChecksum = NewPatternFlowRsvpRsvpChecksum().msg()
	}
	if obj.rsvpChecksumHolder == nil {
		obj.rsvpChecksumHolder = &patternFlowRsvpRsvpChecksum{obj: obj.obj.RsvpChecksum}
	}
	return obj.rsvpChecksumHolder
}

// description is TBD
// RsvpChecksum returns a PatternFlowRsvpRsvpChecksum
func (obj *flowRsvp) HasRsvpChecksum() bool {
	return obj.obj.RsvpChecksum != nil
}

// description is TBD
// SetRsvpChecksum sets the PatternFlowRsvpRsvpChecksum value in the FlowRsvp object
func (obj *flowRsvp) SetRsvpChecksum(value PatternFlowRsvpRsvpChecksum) FlowRsvp {

	obj.rsvpChecksumHolder = nil
	obj.obj.RsvpChecksum = value.msg()

	return obj
}

// description is TBD
// TimeToLive returns a PatternFlowRsvpTimeToLive
func (obj *flowRsvp) TimeToLive() PatternFlowRsvpTimeToLive {
	if obj.obj.TimeToLive == nil {
		obj.obj.TimeToLive = NewPatternFlowRsvpTimeToLive().msg()
	}
	if obj.timeToLiveHolder == nil {
		obj.timeToLiveHolder = &patternFlowRsvpTimeToLive{obj: obj.obj.TimeToLive}
	}
	return obj.timeToLiveHolder
}

// description is TBD
// TimeToLive returns a PatternFlowRsvpTimeToLive
func (obj *flowRsvp) HasTimeToLive() bool {
	return obj.obj.TimeToLive != nil
}

// description is TBD
// SetTimeToLive sets the PatternFlowRsvpTimeToLive value in the FlowRsvp object
func (obj *flowRsvp) SetTimeToLive(value PatternFlowRsvpTimeToLive) FlowRsvp {

	obj.timeToLiveHolder = nil
	obj.obj.TimeToLive = value.msg()

	return obj
}

// description is TBD
// Reserved returns a PatternFlowRsvpReserved
func (obj *flowRsvp) Reserved() PatternFlowRsvpReserved {
	if obj.obj.Reserved == nil {
		obj.obj.Reserved = NewPatternFlowRsvpReserved().msg()
	}
	if obj.reservedHolder == nil {
		obj.reservedHolder = &patternFlowRsvpReserved{obj: obj.obj.Reserved}
	}
	return obj.reservedHolder
}

// description is TBD
// Reserved returns a PatternFlowRsvpReserved
func (obj *flowRsvp) HasReserved() bool {
	return obj.obj.Reserved != nil
}

// description is TBD
// SetReserved sets the PatternFlowRsvpReserved value in the FlowRsvp object
func (obj *flowRsvp) SetReserved(value PatternFlowRsvpReserved) FlowRsvp {

	obj.reservedHolder = nil
	obj.obj.Reserved = value.msg()

	return obj
}

// The sum of the lengths of the common header and all objects included in the message.
// RsvpLength returns a FlowRSVPLength
func (obj *flowRsvp) RsvpLength() FlowRSVPLength {
	if obj.obj.RsvpLength == nil {
		obj.obj.RsvpLength = NewFlowRSVPLength().msg()
	}
	if obj.rsvpLengthHolder == nil {
		obj.rsvpLengthHolder = &flowRSVPLength{obj: obj.obj.RsvpLength}
	}
	return obj.rsvpLengthHolder
}

// The sum of the lengths of the common header and all objects included in the message.
// RsvpLength returns a FlowRSVPLength
func (obj *flowRsvp) HasRsvpLength() bool {
	return obj.obj.RsvpLength != nil
}

// The sum of the lengths of the common header and all objects included in the message.
// SetRsvpLength sets the FlowRSVPLength value in the FlowRsvp object
func (obj *flowRsvp) SetRsvpLength(value FlowRSVPLength) FlowRsvp {

	obj.rsvpLengthHolder = nil
	obj.obj.RsvpLength = value.msg()

	return obj
}

// An 8-bit number that identifies the function of the RSVP message. There are aound 20 message types defined in https://www.iana.org/assignments/rsvp-parameters/rsvp-parameters.xhtml#rsvp-parameters-2 . Among these presently supported is "Path"(value: 1) message type.
// MessageType returns a FlowRSVPMessage
func (obj *flowRsvp) MessageType() FlowRSVPMessage {
	if obj.obj.MessageType == nil {
		obj.obj.MessageType = NewFlowRSVPMessage().msg()
	}
	if obj.messageTypeHolder == nil {
		obj.messageTypeHolder = &flowRSVPMessage{obj: obj.obj.MessageType}
	}
	return obj.messageTypeHolder
}

// An 8-bit number that identifies the function of the RSVP message. There are aound 20 message types defined in https://www.iana.org/assignments/rsvp-parameters/rsvp-parameters.xhtml#rsvp-parameters-2 . Among these presently supported is "Path"(value: 1) message type.
// MessageType returns a FlowRSVPMessage
func (obj *flowRsvp) HasMessageType() bool {
	return obj.obj.MessageType != nil
}

// An 8-bit number that identifies the function of the RSVP message. There are aound 20 message types defined in https://www.iana.org/assignments/rsvp-parameters/rsvp-parameters.xhtml#rsvp-parameters-2 . Among these presently supported is "Path"(value: 1) message type.
// SetMessageType sets the FlowRSVPMessage value in the FlowRsvp object
func (obj *flowRsvp) SetMessageType(value FlowRSVPMessage) FlowRsvp {

	obj.messageTypeHolder = nil
	obj.obj.MessageType = value.msg()

	return obj
}

func (obj *flowRsvp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Version != nil {

		if *obj.obj.Version > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowRsvp.Version <= 15 but Got %d", *obj.obj.Version))
		}

	}

	if obj.obj.RsvpChecksum != nil {

		obj.RsvpChecksum().validateObj(vObj, set_default)
	}

	if obj.obj.TimeToLive != nil {

		obj.TimeToLive().validateObj(vObj, set_default)
	}

	if obj.obj.Reserved != nil {

		obj.Reserved().validateObj(vObj, set_default)
	}

	if obj.obj.RsvpLength != nil {

		obj.RsvpLength().validateObj(vObj, set_default)
	}

	if obj.obj.MessageType != nil {

		obj.MessageType().validateObj(vObj, set_default)
	}

}

func (obj *flowRsvp) setDefault() {
	if obj.obj.Version == nil {
		obj.SetVersion(1)
	}
	if obj.obj.Flag == nil {
		obj.SetFlag(FlowRsvpFlag.NOT_REFRESH_REDUCTION_CAPABLE)

	}

}
