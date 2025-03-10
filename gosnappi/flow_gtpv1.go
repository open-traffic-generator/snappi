package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowGtpv1 *****
type flowGtpv1 struct {
	validation
	obj                           *otg.FlowGtpv1
	marshaller                    marshalFlowGtpv1
	unMarshaller                  unMarshalFlowGtpv1
	versionHolder                 PatternFlowGtpv1Version
	protocolTypeHolder            PatternFlowGtpv1ProtocolType
	reservedHolder                PatternFlowGtpv1Reserved
	eFlagHolder                   PatternFlowGtpv1EFlag
	sFlagHolder                   PatternFlowGtpv1SFlag
	pnFlagHolder                  PatternFlowGtpv1PnFlag
	messageTypeHolder             PatternFlowGtpv1MessageType
	messageLengthHolder           PatternFlowGtpv1MessageLength
	teidHolder                    PatternFlowGtpv1Teid
	squenceNumberHolder           PatternFlowGtpv1SquenceNumber
	nPduNumberHolder              PatternFlowGtpv1NPduNumber
	nextExtensionHeaderTypeHolder PatternFlowGtpv1NextExtensionHeaderType
	extensionHeadersHolder        FlowGtpv1FlowGtpExtensionIter
}

func NewFlowGtpv1() FlowGtpv1 {
	obj := flowGtpv1{obj: &otg.FlowGtpv1{}}
	obj.setDefault()
	return &obj
}

func (obj *flowGtpv1) msg() *otg.FlowGtpv1 {
	return obj.obj
}

func (obj *flowGtpv1) setMsg(msg *otg.FlowGtpv1) FlowGtpv1 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowGtpv1 struct {
	obj *flowGtpv1
}

type marshalFlowGtpv1 interface {
	// ToProto marshals FlowGtpv1 to protobuf object *otg.FlowGtpv1
	ToProto() (*otg.FlowGtpv1, error)
	// ToPbText marshals FlowGtpv1 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowGtpv1 to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowGtpv1 to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowGtpv1 to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowGtpv1 struct {
	obj *flowGtpv1
}

type unMarshalFlowGtpv1 interface {
	// FromProto unmarshals FlowGtpv1 from protobuf object *otg.FlowGtpv1
	FromProto(msg *otg.FlowGtpv1) (FlowGtpv1, error)
	// FromPbText unmarshals FlowGtpv1 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowGtpv1 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowGtpv1 from JSON text
	FromJson(value string) error
}

func (obj *flowGtpv1) Marshal() marshalFlowGtpv1 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowGtpv1{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowGtpv1) Unmarshal() unMarshalFlowGtpv1 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowGtpv1{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowGtpv1) ToProto() (*otg.FlowGtpv1, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowGtpv1) FromProto(msg *otg.FlowGtpv1) (FlowGtpv1, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowGtpv1) ToPbText() (string, error) {
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

func (m *unMarshalflowGtpv1) FromPbText(value string) error {
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

func (m *marshalflowGtpv1) ToYaml() (string, error) {
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

func (m *unMarshalflowGtpv1) FromYaml(value string) error {
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

func (m *marshalflowGtpv1) ToJsonRaw() (string, error) {
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

func (m *marshalflowGtpv1) ToJson() (string, error) {
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

func (m *unMarshalflowGtpv1) FromJson(value string) error {
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

func (obj *flowGtpv1) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowGtpv1) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowGtpv1) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowGtpv1) Clone() (FlowGtpv1, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowGtpv1()
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

func (obj *flowGtpv1) setNil() {
	obj.versionHolder = nil
	obj.protocolTypeHolder = nil
	obj.reservedHolder = nil
	obj.eFlagHolder = nil
	obj.sFlagHolder = nil
	obj.pnFlagHolder = nil
	obj.messageTypeHolder = nil
	obj.messageLengthHolder = nil
	obj.teidHolder = nil
	obj.squenceNumberHolder = nil
	obj.nPduNumberHolder = nil
	obj.nextExtensionHeaderTypeHolder = nil
	obj.extensionHeadersHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowGtpv1 is gTPv1 packet header
type FlowGtpv1 interface {
	Validation
	// msg marshals FlowGtpv1 to protobuf object *otg.FlowGtpv1
	// and doesn't set defaults
	msg() *otg.FlowGtpv1
	// setMsg unmarshals FlowGtpv1 from protobuf object *otg.FlowGtpv1
	// and doesn't set defaults
	setMsg(*otg.FlowGtpv1) FlowGtpv1
	// provides marshal interface
	Marshal() marshalFlowGtpv1
	// provides unmarshal interface
	Unmarshal() unMarshalFlowGtpv1
	// validate validates FlowGtpv1
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowGtpv1, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Version returns PatternFlowGtpv1Version, set in FlowGtpv1.
	// PatternFlowGtpv1Version is gTPv1 version
	Version() PatternFlowGtpv1Version
	// SetVersion assigns PatternFlowGtpv1Version provided by user to FlowGtpv1.
	// PatternFlowGtpv1Version is gTPv1 version
	SetVersion(value PatternFlowGtpv1Version) FlowGtpv1
	// HasVersion checks if Version has been set in FlowGtpv1
	HasVersion() bool
	// ProtocolType returns PatternFlowGtpv1ProtocolType, set in FlowGtpv1.
	// PatternFlowGtpv1ProtocolType is protocol type, GTP is 1, GTP' is 0
	ProtocolType() PatternFlowGtpv1ProtocolType
	// SetProtocolType assigns PatternFlowGtpv1ProtocolType provided by user to FlowGtpv1.
	// PatternFlowGtpv1ProtocolType is protocol type, GTP is 1, GTP' is 0
	SetProtocolType(value PatternFlowGtpv1ProtocolType) FlowGtpv1
	// HasProtocolType checks if ProtocolType has been set in FlowGtpv1
	HasProtocolType() bool
	// Reserved returns PatternFlowGtpv1Reserved, set in FlowGtpv1.
	// PatternFlowGtpv1Reserved is reserved field
	Reserved() PatternFlowGtpv1Reserved
	// SetReserved assigns PatternFlowGtpv1Reserved provided by user to FlowGtpv1.
	// PatternFlowGtpv1Reserved is reserved field
	SetReserved(value PatternFlowGtpv1Reserved) FlowGtpv1
	// HasReserved checks if Reserved has been set in FlowGtpv1
	HasReserved() bool
	// EFlag returns PatternFlowGtpv1EFlag, set in FlowGtpv1.
	// PatternFlowGtpv1EFlag is extension header field present
	EFlag() PatternFlowGtpv1EFlag
	// SetEFlag assigns PatternFlowGtpv1EFlag provided by user to FlowGtpv1.
	// PatternFlowGtpv1EFlag is extension header field present
	SetEFlag(value PatternFlowGtpv1EFlag) FlowGtpv1
	// HasEFlag checks if EFlag has been set in FlowGtpv1
	HasEFlag() bool
	// SFlag returns PatternFlowGtpv1SFlag, set in FlowGtpv1.
	// PatternFlowGtpv1SFlag is sequence number field present
	SFlag() PatternFlowGtpv1SFlag
	// SetSFlag assigns PatternFlowGtpv1SFlag provided by user to FlowGtpv1.
	// PatternFlowGtpv1SFlag is sequence number field present
	SetSFlag(value PatternFlowGtpv1SFlag) FlowGtpv1
	// HasSFlag checks if SFlag has been set in FlowGtpv1
	HasSFlag() bool
	// PnFlag returns PatternFlowGtpv1PnFlag, set in FlowGtpv1.
	// PatternFlowGtpv1PnFlag is n-PDU field present
	PnFlag() PatternFlowGtpv1PnFlag
	// SetPnFlag assigns PatternFlowGtpv1PnFlag provided by user to FlowGtpv1.
	// PatternFlowGtpv1PnFlag is n-PDU field present
	SetPnFlag(value PatternFlowGtpv1PnFlag) FlowGtpv1
	// HasPnFlag checks if PnFlag has been set in FlowGtpv1
	HasPnFlag() bool
	// MessageType returns PatternFlowGtpv1MessageType, set in FlowGtpv1.
	// PatternFlowGtpv1MessageType is the type of GTP message Different types of messages are defined in 3GPP TS 29.060 section 7.1
	MessageType() PatternFlowGtpv1MessageType
	// SetMessageType assigns PatternFlowGtpv1MessageType provided by user to FlowGtpv1.
	// PatternFlowGtpv1MessageType is the type of GTP message Different types of messages are defined in 3GPP TS 29.060 section 7.1
	SetMessageType(value PatternFlowGtpv1MessageType) FlowGtpv1
	// HasMessageType checks if MessageType has been set in FlowGtpv1
	HasMessageType() bool
	// MessageLength returns PatternFlowGtpv1MessageLength, set in FlowGtpv1.
	// PatternFlowGtpv1MessageLength is the length of the payload (the bytes following the mandatory 8-byte GTP header) in bytes that includes any optional fields
	MessageLength() PatternFlowGtpv1MessageLength
	// SetMessageLength assigns PatternFlowGtpv1MessageLength provided by user to FlowGtpv1.
	// PatternFlowGtpv1MessageLength is the length of the payload (the bytes following the mandatory 8-byte GTP header) in bytes that includes any optional fields
	SetMessageLength(value PatternFlowGtpv1MessageLength) FlowGtpv1
	// HasMessageLength checks if MessageLength has been set in FlowGtpv1
	HasMessageLength() bool
	// Teid returns PatternFlowGtpv1Teid, set in FlowGtpv1.
	// PatternFlowGtpv1Teid is tunnel endpoint identifier (TEID) used to multiplex connections in the same GTP tunnel
	Teid() PatternFlowGtpv1Teid
	// SetTeid assigns PatternFlowGtpv1Teid provided by user to FlowGtpv1.
	// PatternFlowGtpv1Teid is tunnel endpoint identifier (TEID) used to multiplex connections in the same GTP tunnel
	SetTeid(value PatternFlowGtpv1Teid) FlowGtpv1
	// HasTeid checks if Teid has been set in FlowGtpv1
	HasTeid() bool
	// SquenceNumber returns PatternFlowGtpv1SquenceNumber, set in FlowGtpv1.
	// PatternFlowGtpv1SquenceNumber is sequence number. Exists if any of the e_flag, s_flag, or pn_flag bits are on.  Must be interpreted only if the s_flag bit is on.
	SquenceNumber() PatternFlowGtpv1SquenceNumber
	// SetSquenceNumber assigns PatternFlowGtpv1SquenceNumber provided by user to FlowGtpv1.
	// PatternFlowGtpv1SquenceNumber is sequence number. Exists if any of the e_flag, s_flag, or pn_flag bits are on.  Must be interpreted only if the s_flag bit is on.
	SetSquenceNumber(value PatternFlowGtpv1SquenceNumber) FlowGtpv1
	// HasSquenceNumber checks if SquenceNumber has been set in FlowGtpv1
	HasSquenceNumber() bool
	// NPduNumber returns PatternFlowGtpv1NPduNumber, set in FlowGtpv1.
	// PatternFlowGtpv1NPduNumber is n-PDU number. Exists if any of the e_flag, s_flag, or pn_flag bits are on.  Must be interpreted only if the pn_flag bit is on.
	NPduNumber() PatternFlowGtpv1NPduNumber
	// SetNPduNumber assigns PatternFlowGtpv1NPduNumber provided by user to FlowGtpv1.
	// PatternFlowGtpv1NPduNumber is n-PDU number. Exists if any of the e_flag, s_flag, or pn_flag bits are on.  Must be interpreted only if the pn_flag bit is on.
	SetNPduNumber(value PatternFlowGtpv1NPduNumber) FlowGtpv1
	// HasNPduNumber checks if NPduNumber has been set in FlowGtpv1
	HasNPduNumber() bool
	// NextExtensionHeaderType returns PatternFlowGtpv1NextExtensionHeaderType, set in FlowGtpv1.
	// PatternFlowGtpv1NextExtensionHeaderType is next extension header. Exists if any of the e_flag, s_flag, or pn_flag bits are on.  Must be interpreted only if the e_flag bit is on.
	NextExtensionHeaderType() PatternFlowGtpv1NextExtensionHeaderType
	// SetNextExtensionHeaderType assigns PatternFlowGtpv1NextExtensionHeaderType provided by user to FlowGtpv1.
	// PatternFlowGtpv1NextExtensionHeaderType is next extension header. Exists if any of the e_flag, s_flag, or pn_flag bits are on.  Must be interpreted only if the e_flag bit is on.
	SetNextExtensionHeaderType(value PatternFlowGtpv1NextExtensionHeaderType) FlowGtpv1
	// HasNextExtensionHeaderType checks if NextExtensionHeaderType has been set in FlowGtpv1
	HasNextExtensionHeaderType() bool
	// ExtensionHeaders returns FlowGtpv1FlowGtpExtensionIterIter, set in FlowGtpv1
	ExtensionHeaders() FlowGtpv1FlowGtpExtensionIter
	setNil()
}

// description is TBD
// Version returns a PatternFlowGtpv1Version
func (obj *flowGtpv1) Version() PatternFlowGtpv1Version {
	if obj.obj.Version == nil {
		obj.obj.Version = NewPatternFlowGtpv1Version().msg()
	}
	if obj.versionHolder == nil {
		obj.versionHolder = &patternFlowGtpv1Version{obj: obj.obj.Version}
	}
	return obj.versionHolder
}

// description is TBD
// Version returns a PatternFlowGtpv1Version
func (obj *flowGtpv1) HasVersion() bool {
	return obj.obj.Version != nil
}

// description is TBD
// SetVersion sets the PatternFlowGtpv1Version value in the FlowGtpv1 object
func (obj *flowGtpv1) SetVersion(value PatternFlowGtpv1Version) FlowGtpv1 {

	obj.versionHolder = nil
	obj.obj.Version = value.msg()

	return obj
}

// description is TBD
// ProtocolType returns a PatternFlowGtpv1ProtocolType
func (obj *flowGtpv1) ProtocolType() PatternFlowGtpv1ProtocolType {
	if obj.obj.ProtocolType == nil {
		obj.obj.ProtocolType = NewPatternFlowGtpv1ProtocolType().msg()
	}
	if obj.protocolTypeHolder == nil {
		obj.protocolTypeHolder = &patternFlowGtpv1ProtocolType{obj: obj.obj.ProtocolType}
	}
	return obj.protocolTypeHolder
}

// description is TBD
// ProtocolType returns a PatternFlowGtpv1ProtocolType
func (obj *flowGtpv1) HasProtocolType() bool {
	return obj.obj.ProtocolType != nil
}

// description is TBD
// SetProtocolType sets the PatternFlowGtpv1ProtocolType value in the FlowGtpv1 object
func (obj *flowGtpv1) SetProtocolType(value PatternFlowGtpv1ProtocolType) FlowGtpv1 {

	obj.protocolTypeHolder = nil
	obj.obj.ProtocolType = value.msg()

	return obj
}

// description is TBD
// Reserved returns a PatternFlowGtpv1Reserved
func (obj *flowGtpv1) Reserved() PatternFlowGtpv1Reserved {
	if obj.obj.Reserved == nil {
		obj.obj.Reserved = NewPatternFlowGtpv1Reserved().msg()
	}
	if obj.reservedHolder == nil {
		obj.reservedHolder = &patternFlowGtpv1Reserved{obj: obj.obj.Reserved}
	}
	return obj.reservedHolder
}

// description is TBD
// Reserved returns a PatternFlowGtpv1Reserved
func (obj *flowGtpv1) HasReserved() bool {
	return obj.obj.Reserved != nil
}

// description is TBD
// SetReserved sets the PatternFlowGtpv1Reserved value in the FlowGtpv1 object
func (obj *flowGtpv1) SetReserved(value PatternFlowGtpv1Reserved) FlowGtpv1 {

	obj.reservedHolder = nil
	obj.obj.Reserved = value.msg()

	return obj
}

// description is TBD
// EFlag returns a PatternFlowGtpv1EFlag
func (obj *flowGtpv1) EFlag() PatternFlowGtpv1EFlag {
	if obj.obj.EFlag == nil {
		obj.obj.EFlag = NewPatternFlowGtpv1EFlag().msg()
	}
	if obj.eFlagHolder == nil {
		obj.eFlagHolder = &patternFlowGtpv1EFlag{obj: obj.obj.EFlag}
	}
	return obj.eFlagHolder
}

// description is TBD
// EFlag returns a PatternFlowGtpv1EFlag
func (obj *flowGtpv1) HasEFlag() bool {
	return obj.obj.EFlag != nil
}

// description is TBD
// SetEFlag sets the PatternFlowGtpv1EFlag value in the FlowGtpv1 object
func (obj *flowGtpv1) SetEFlag(value PatternFlowGtpv1EFlag) FlowGtpv1 {

	obj.eFlagHolder = nil
	obj.obj.EFlag = value.msg()

	return obj
}

// description is TBD
// SFlag returns a PatternFlowGtpv1SFlag
func (obj *flowGtpv1) SFlag() PatternFlowGtpv1SFlag {
	if obj.obj.SFlag == nil {
		obj.obj.SFlag = NewPatternFlowGtpv1SFlag().msg()
	}
	if obj.sFlagHolder == nil {
		obj.sFlagHolder = &patternFlowGtpv1SFlag{obj: obj.obj.SFlag}
	}
	return obj.sFlagHolder
}

// description is TBD
// SFlag returns a PatternFlowGtpv1SFlag
func (obj *flowGtpv1) HasSFlag() bool {
	return obj.obj.SFlag != nil
}

// description is TBD
// SetSFlag sets the PatternFlowGtpv1SFlag value in the FlowGtpv1 object
func (obj *flowGtpv1) SetSFlag(value PatternFlowGtpv1SFlag) FlowGtpv1 {

	obj.sFlagHolder = nil
	obj.obj.SFlag = value.msg()

	return obj
}

// description is TBD
// PnFlag returns a PatternFlowGtpv1PnFlag
func (obj *flowGtpv1) PnFlag() PatternFlowGtpv1PnFlag {
	if obj.obj.PnFlag == nil {
		obj.obj.PnFlag = NewPatternFlowGtpv1PnFlag().msg()
	}
	if obj.pnFlagHolder == nil {
		obj.pnFlagHolder = &patternFlowGtpv1PnFlag{obj: obj.obj.PnFlag}
	}
	return obj.pnFlagHolder
}

// description is TBD
// PnFlag returns a PatternFlowGtpv1PnFlag
func (obj *flowGtpv1) HasPnFlag() bool {
	return obj.obj.PnFlag != nil
}

// description is TBD
// SetPnFlag sets the PatternFlowGtpv1PnFlag value in the FlowGtpv1 object
func (obj *flowGtpv1) SetPnFlag(value PatternFlowGtpv1PnFlag) FlowGtpv1 {

	obj.pnFlagHolder = nil
	obj.obj.PnFlag = value.msg()

	return obj
}

// description is TBD
// MessageType returns a PatternFlowGtpv1MessageType
func (obj *flowGtpv1) MessageType() PatternFlowGtpv1MessageType {
	if obj.obj.MessageType == nil {
		obj.obj.MessageType = NewPatternFlowGtpv1MessageType().msg()
	}
	if obj.messageTypeHolder == nil {
		obj.messageTypeHolder = &patternFlowGtpv1MessageType{obj: obj.obj.MessageType}
	}
	return obj.messageTypeHolder
}

// description is TBD
// MessageType returns a PatternFlowGtpv1MessageType
func (obj *flowGtpv1) HasMessageType() bool {
	return obj.obj.MessageType != nil
}

// description is TBD
// SetMessageType sets the PatternFlowGtpv1MessageType value in the FlowGtpv1 object
func (obj *flowGtpv1) SetMessageType(value PatternFlowGtpv1MessageType) FlowGtpv1 {

	obj.messageTypeHolder = nil
	obj.obj.MessageType = value.msg()

	return obj
}

// description is TBD
// MessageLength returns a PatternFlowGtpv1MessageLength
func (obj *flowGtpv1) MessageLength() PatternFlowGtpv1MessageLength {
	if obj.obj.MessageLength == nil {
		obj.obj.MessageLength = NewPatternFlowGtpv1MessageLength().msg()
	}
	if obj.messageLengthHolder == nil {
		obj.messageLengthHolder = &patternFlowGtpv1MessageLength{obj: obj.obj.MessageLength}
	}
	return obj.messageLengthHolder
}

// description is TBD
// MessageLength returns a PatternFlowGtpv1MessageLength
func (obj *flowGtpv1) HasMessageLength() bool {
	return obj.obj.MessageLength != nil
}

// description is TBD
// SetMessageLength sets the PatternFlowGtpv1MessageLength value in the FlowGtpv1 object
func (obj *flowGtpv1) SetMessageLength(value PatternFlowGtpv1MessageLength) FlowGtpv1 {

	obj.messageLengthHolder = nil
	obj.obj.MessageLength = value.msg()

	return obj
}

// description is TBD
// Teid returns a PatternFlowGtpv1Teid
func (obj *flowGtpv1) Teid() PatternFlowGtpv1Teid {
	if obj.obj.Teid == nil {
		obj.obj.Teid = NewPatternFlowGtpv1Teid().msg()
	}
	if obj.teidHolder == nil {
		obj.teidHolder = &patternFlowGtpv1Teid{obj: obj.obj.Teid}
	}
	return obj.teidHolder
}

// description is TBD
// Teid returns a PatternFlowGtpv1Teid
func (obj *flowGtpv1) HasTeid() bool {
	return obj.obj.Teid != nil
}

// description is TBD
// SetTeid sets the PatternFlowGtpv1Teid value in the FlowGtpv1 object
func (obj *flowGtpv1) SetTeid(value PatternFlowGtpv1Teid) FlowGtpv1 {

	obj.teidHolder = nil
	obj.obj.Teid = value.msg()

	return obj
}

// description is TBD
// SquenceNumber returns a PatternFlowGtpv1SquenceNumber
func (obj *flowGtpv1) SquenceNumber() PatternFlowGtpv1SquenceNumber {
	if obj.obj.SquenceNumber == nil {
		obj.obj.SquenceNumber = NewPatternFlowGtpv1SquenceNumber().msg()
	}
	if obj.squenceNumberHolder == nil {
		obj.squenceNumberHolder = &patternFlowGtpv1SquenceNumber{obj: obj.obj.SquenceNumber}
	}
	return obj.squenceNumberHolder
}

// description is TBD
// SquenceNumber returns a PatternFlowGtpv1SquenceNumber
func (obj *flowGtpv1) HasSquenceNumber() bool {
	return obj.obj.SquenceNumber != nil
}

// description is TBD
// SetSquenceNumber sets the PatternFlowGtpv1SquenceNumber value in the FlowGtpv1 object
func (obj *flowGtpv1) SetSquenceNumber(value PatternFlowGtpv1SquenceNumber) FlowGtpv1 {

	obj.squenceNumberHolder = nil
	obj.obj.SquenceNumber = value.msg()

	return obj
}

// description is TBD
// NPduNumber returns a PatternFlowGtpv1NPduNumber
func (obj *flowGtpv1) NPduNumber() PatternFlowGtpv1NPduNumber {
	if obj.obj.NPduNumber == nil {
		obj.obj.NPduNumber = NewPatternFlowGtpv1NPduNumber().msg()
	}
	if obj.nPduNumberHolder == nil {
		obj.nPduNumberHolder = &patternFlowGtpv1NPduNumber{obj: obj.obj.NPduNumber}
	}
	return obj.nPduNumberHolder
}

// description is TBD
// NPduNumber returns a PatternFlowGtpv1NPduNumber
func (obj *flowGtpv1) HasNPduNumber() bool {
	return obj.obj.NPduNumber != nil
}

// description is TBD
// SetNPduNumber sets the PatternFlowGtpv1NPduNumber value in the FlowGtpv1 object
func (obj *flowGtpv1) SetNPduNumber(value PatternFlowGtpv1NPduNumber) FlowGtpv1 {

	obj.nPduNumberHolder = nil
	obj.obj.NPduNumber = value.msg()

	return obj
}

// description is TBD
// NextExtensionHeaderType returns a PatternFlowGtpv1NextExtensionHeaderType
func (obj *flowGtpv1) NextExtensionHeaderType() PatternFlowGtpv1NextExtensionHeaderType {
	if obj.obj.NextExtensionHeaderType == nil {
		obj.obj.NextExtensionHeaderType = NewPatternFlowGtpv1NextExtensionHeaderType().msg()
	}
	if obj.nextExtensionHeaderTypeHolder == nil {
		obj.nextExtensionHeaderTypeHolder = &patternFlowGtpv1NextExtensionHeaderType{obj: obj.obj.NextExtensionHeaderType}
	}
	return obj.nextExtensionHeaderTypeHolder
}

// description is TBD
// NextExtensionHeaderType returns a PatternFlowGtpv1NextExtensionHeaderType
func (obj *flowGtpv1) HasNextExtensionHeaderType() bool {
	return obj.obj.NextExtensionHeaderType != nil
}

// description is TBD
// SetNextExtensionHeaderType sets the PatternFlowGtpv1NextExtensionHeaderType value in the FlowGtpv1 object
func (obj *flowGtpv1) SetNextExtensionHeaderType(value PatternFlowGtpv1NextExtensionHeaderType) FlowGtpv1 {

	obj.nextExtensionHeaderTypeHolder = nil
	obj.obj.NextExtensionHeaderType = value.msg()

	return obj
}

// A list of optional extension headers.
// ExtensionHeaders returns a []FlowGtpExtension
func (obj *flowGtpv1) ExtensionHeaders() FlowGtpv1FlowGtpExtensionIter {
	if len(obj.obj.ExtensionHeaders) == 0 {
		obj.obj.ExtensionHeaders = []*otg.FlowGtpExtension{}
	}
	if obj.extensionHeadersHolder == nil {
		obj.extensionHeadersHolder = newFlowGtpv1FlowGtpExtensionIter(&obj.obj.ExtensionHeaders).setMsg(obj)
	}
	return obj.extensionHeadersHolder
}

type flowGtpv1FlowGtpExtensionIter struct {
	obj                   *flowGtpv1
	flowGtpExtensionSlice []FlowGtpExtension
	fieldPtr              *[]*otg.FlowGtpExtension
}

func newFlowGtpv1FlowGtpExtensionIter(ptr *[]*otg.FlowGtpExtension) FlowGtpv1FlowGtpExtensionIter {
	return &flowGtpv1FlowGtpExtensionIter{fieldPtr: ptr}
}

type FlowGtpv1FlowGtpExtensionIter interface {
	setMsg(*flowGtpv1) FlowGtpv1FlowGtpExtensionIter
	Items() []FlowGtpExtension
	Add() FlowGtpExtension
	Append(items ...FlowGtpExtension) FlowGtpv1FlowGtpExtensionIter
	Set(index int, newObj FlowGtpExtension) FlowGtpv1FlowGtpExtensionIter
	Clear() FlowGtpv1FlowGtpExtensionIter
	clearHolderSlice() FlowGtpv1FlowGtpExtensionIter
	appendHolderSlice(item FlowGtpExtension) FlowGtpv1FlowGtpExtensionIter
}

func (obj *flowGtpv1FlowGtpExtensionIter) setMsg(msg *flowGtpv1) FlowGtpv1FlowGtpExtensionIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowGtpExtension{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowGtpv1FlowGtpExtensionIter) Items() []FlowGtpExtension {
	return obj.flowGtpExtensionSlice
}

func (obj *flowGtpv1FlowGtpExtensionIter) Add() FlowGtpExtension {
	newObj := &otg.FlowGtpExtension{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowGtpExtension{obj: newObj}
	newLibObj.setDefault()
	obj.flowGtpExtensionSlice = append(obj.flowGtpExtensionSlice, newLibObj)
	return newLibObj
}

func (obj *flowGtpv1FlowGtpExtensionIter) Append(items ...FlowGtpExtension) FlowGtpv1FlowGtpExtensionIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowGtpExtensionSlice = append(obj.flowGtpExtensionSlice, item)
	}
	return obj
}

func (obj *flowGtpv1FlowGtpExtensionIter) Set(index int, newObj FlowGtpExtension) FlowGtpv1FlowGtpExtensionIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowGtpExtensionSlice[index] = newObj
	return obj
}
func (obj *flowGtpv1FlowGtpExtensionIter) Clear() FlowGtpv1FlowGtpExtensionIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowGtpExtension{}
		obj.flowGtpExtensionSlice = []FlowGtpExtension{}
	}
	return obj
}
func (obj *flowGtpv1FlowGtpExtensionIter) clearHolderSlice() FlowGtpv1FlowGtpExtensionIter {
	if len(obj.flowGtpExtensionSlice) > 0 {
		obj.flowGtpExtensionSlice = []FlowGtpExtension{}
	}
	return obj
}
func (obj *flowGtpv1FlowGtpExtensionIter) appendHolderSlice(item FlowGtpExtension) FlowGtpv1FlowGtpExtensionIter {
	obj.flowGtpExtensionSlice = append(obj.flowGtpExtensionSlice, item)
	return obj
}

func (obj *flowGtpv1) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Version != nil {

		obj.Version().validateObj(vObj, set_default)
	}

	if obj.obj.ProtocolType != nil {

		obj.ProtocolType().validateObj(vObj, set_default)
	}

	if obj.obj.Reserved != nil {

		obj.Reserved().validateObj(vObj, set_default)
	}

	if obj.obj.EFlag != nil {

		obj.EFlag().validateObj(vObj, set_default)
	}

	if obj.obj.SFlag != nil {

		obj.SFlag().validateObj(vObj, set_default)
	}

	if obj.obj.PnFlag != nil {

		obj.PnFlag().validateObj(vObj, set_default)
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

	if obj.obj.SquenceNumber != nil {

		obj.SquenceNumber().validateObj(vObj, set_default)
	}

	if obj.obj.NPduNumber != nil {

		obj.NPduNumber().validateObj(vObj, set_default)
	}

	if obj.obj.NextExtensionHeaderType != nil {

		obj.NextExtensionHeaderType().validateObj(vObj, set_default)
	}

	if len(obj.obj.ExtensionHeaders) != 0 {

		if set_default {
			obj.ExtensionHeaders().clearHolderSlice()
			for _, item := range obj.obj.ExtensionHeaders {
				obj.ExtensionHeaders().appendHolderSlice(&flowGtpExtension{obj: item})
			}
		}
		for _, item := range obj.ExtensionHeaders().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *flowGtpv1) setDefault() {

}
