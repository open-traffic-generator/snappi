package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowTcp *****
type flowTcp struct {
	validation
	obj              *otg.FlowTcp
	marshaller       marshalFlowTcp
	unMarshaller     unMarshalFlowTcp
	srcPortHolder    PatternFlowTcpSrcPort
	dstPortHolder    PatternFlowTcpDstPort
	seqNumHolder     PatternFlowTcpSeqNum
	ackNumHolder     PatternFlowTcpAckNum
	dataOffsetHolder PatternFlowTcpDataOffset
	ecnNsHolder      PatternFlowTcpEcnNs
	ecnCwrHolder     PatternFlowTcpEcnCwr
	ecnEchoHolder    PatternFlowTcpEcnEcho
	ctlUrgHolder     PatternFlowTcpCtlUrg
	ctlAckHolder     PatternFlowTcpCtlAck
	ctlPshHolder     PatternFlowTcpCtlPsh
	ctlRstHolder     PatternFlowTcpCtlRst
	ctlSynHolder     PatternFlowTcpCtlSyn
	ctlFinHolder     PatternFlowTcpCtlFin
	windowHolder     PatternFlowTcpWindow
	checksumHolder   PatternFlowTcpChecksum
}

func NewFlowTcp() FlowTcp {
	obj := flowTcp{obj: &otg.FlowTcp{}}
	obj.setDefault()
	return &obj
}

func (obj *flowTcp) msg() *otg.FlowTcp {
	return obj.obj
}

func (obj *flowTcp) setMsg(msg *otg.FlowTcp) FlowTcp {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowTcp struct {
	obj *flowTcp
}

type marshalFlowTcp interface {
	// ToProto marshals FlowTcp to protobuf object *otg.FlowTcp
	ToProto() (*otg.FlowTcp, error)
	// ToPbText marshals FlowTcp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowTcp to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowTcp to JSON text
	ToJson() (string, error)
}

type unMarshalflowTcp struct {
	obj *flowTcp
}

type unMarshalFlowTcp interface {
	// FromProto unmarshals FlowTcp from protobuf object *otg.FlowTcp
	FromProto(msg *otg.FlowTcp) (FlowTcp, error)
	// FromPbText unmarshals FlowTcp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowTcp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowTcp from JSON text
	FromJson(value string) error
}

func (obj *flowTcp) Marshal() marshalFlowTcp {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowTcp{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowTcp) Unmarshal() unMarshalFlowTcp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowTcp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowTcp) ToProto() (*otg.FlowTcp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowTcp) FromProto(msg *otg.FlowTcp) (FlowTcp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowTcp) ToPbText() (string, error) {
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

func (m *unMarshalflowTcp) FromPbText(value string) error {
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

func (m *marshalflowTcp) ToYaml() (string, error) {
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

func (m *unMarshalflowTcp) FromYaml(value string) error {
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

func (m *marshalflowTcp) ToJson() (string, error) {
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

func (m *unMarshalflowTcp) FromJson(value string) error {
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

func (obj *flowTcp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowTcp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowTcp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowTcp) Clone() (FlowTcp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowTcp()
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

func (obj *flowTcp) setNil() {
	obj.srcPortHolder = nil
	obj.dstPortHolder = nil
	obj.seqNumHolder = nil
	obj.ackNumHolder = nil
	obj.dataOffsetHolder = nil
	obj.ecnNsHolder = nil
	obj.ecnCwrHolder = nil
	obj.ecnEchoHolder = nil
	obj.ctlUrgHolder = nil
	obj.ctlAckHolder = nil
	obj.ctlPshHolder = nil
	obj.ctlRstHolder = nil
	obj.ctlSynHolder = nil
	obj.ctlFinHolder = nil
	obj.windowHolder = nil
	obj.checksumHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowTcp is tCP packet header
type FlowTcp interface {
	Validation
	// msg marshals FlowTcp to protobuf object *otg.FlowTcp
	// and doesn't set defaults
	msg() *otg.FlowTcp
	// setMsg unmarshals FlowTcp from protobuf object *otg.FlowTcp
	// and doesn't set defaults
	setMsg(*otg.FlowTcp) FlowTcp
	// provides marshal interface
	Marshal() marshalFlowTcp
	// provides unmarshal interface
	Unmarshal() unMarshalFlowTcp
	// validate validates FlowTcp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowTcp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SrcPort returns PatternFlowTcpSrcPort, set in FlowTcp.
	// PatternFlowTcpSrcPort is source port
	SrcPort() PatternFlowTcpSrcPort
	// SetSrcPort assigns PatternFlowTcpSrcPort provided by user to FlowTcp.
	// PatternFlowTcpSrcPort is source port
	SetSrcPort(value PatternFlowTcpSrcPort) FlowTcp
	// HasSrcPort checks if SrcPort has been set in FlowTcp
	HasSrcPort() bool
	// DstPort returns PatternFlowTcpDstPort, set in FlowTcp.
	// PatternFlowTcpDstPort is destination port
	DstPort() PatternFlowTcpDstPort
	// SetDstPort assigns PatternFlowTcpDstPort provided by user to FlowTcp.
	// PatternFlowTcpDstPort is destination port
	SetDstPort(value PatternFlowTcpDstPort) FlowTcp
	// HasDstPort checks if DstPort has been set in FlowTcp
	HasDstPort() bool
	// SeqNum returns PatternFlowTcpSeqNum, set in FlowTcp.
	// PatternFlowTcpSeqNum is sequence number
	SeqNum() PatternFlowTcpSeqNum
	// SetSeqNum assigns PatternFlowTcpSeqNum provided by user to FlowTcp.
	// PatternFlowTcpSeqNum is sequence number
	SetSeqNum(value PatternFlowTcpSeqNum) FlowTcp
	// HasSeqNum checks if SeqNum has been set in FlowTcp
	HasSeqNum() bool
	// AckNum returns PatternFlowTcpAckNum, set in FlowTcp.
	// PatternFlowTcpAckNum is acknowledgement number
	AckNum() PatternFlowTcpAckNum
	// SetAckNum assigns PatternFlowTcpAckNum provided by user to FlowTcp.
	// PatternFlowTcpAckNum is acknowledgement number
	SetAckNum(value PatternFlowTcpAckNum) FlowTcp
	// HasAckNum checks if AckNum has been set in FlowTcp
	HasAckNum() bool
	// DataOffset returns PatternFlowTcpDataOffset, set in FlowTcp.
	// PatternFlowTcpDataOffset is the number of 32 bit words in the TCP header. This indicates where the data begins.
	DataOffset() PatternFlowTcpDataOffset
	// SetDataOffset assigns PatternFlowTcpDataOffset provided by user to FlowTcp.
	// PatternFlowTcpDataOffset is the number of 32 bit words in the TCP header. This indicates where the data begins.
	SetDataOffset(value PatternFlowTcpDataOffset) FlowTcp
	// HasDataOffset checks if DataOffset has been set in FlowTcp
	HasDataOffset() bool
	// EcnNs returns PatternFlowTcpEcnNs, set in FlowTcp.
	// PatternFlowTcpEcnNs is explicit congestion notification, concealment protection.
	EcnNs() PatternFlowTcpEcnNs
	// SetEcnNs assigns PatternFlowTcpEcnNs provided by user to FlowTcp.
	// PatternFlowTcpEcnNs is explicit congestion notification, concealment protection.
	SetEcnNs(value PatternFlowTcpEcnNs) FlowTcp
	// HasEcnNs checks if EcnNs has been set in FlowTcp
	HasEcnNs() bool
	// EcnCwr returns PatternFlowTcpEcnCwr, set in FlowTcp.
	// PatternFlowTcpEcnCwr is explicit congestion notification, congestion window reduced.
	EcnCwr() PatternFlowTcpEcnCwr
	// SetEcnCwr assigns PatternFlowTcpEcnCwr provided by user to FlowTcp.
	// PatternFlowTcpEcnCwr is explicit congestion notification, congestion window reduced.
	SetEcnCwr(value PatternFlowTcpEcnCwr) FlowTcp
	// HasEcnCwr checks if EcnCwr has been set in FlowTcp
	HasEcnCwr() bool
	// EcnEcho returns PatternFlowTcpEcnEcho, set in FlowTcp.
	// PatternFlowTcpEcnEcho is explicit congestion notification, echo. 1 indicates the peer is ecn capable. 0 indicates that a packet with ipv4.ecn = 11 in the ip header was  received during normal transmission.
	EcnEcho() PatternFlowTcpEcnEcho
	// SetEcnEcho assigns PatternFlowTcpEcnEcho provided by user to FlowTcp.
	// PatternFlowTcpEcnEcho is explicit congestion notification, echo. 1 indicates the peer is ecn capable. 0 indicates that a packet with ipv4.ecn = 11 in the ip header was  received during normal transmission.
	SetEcnEcho(value PatternFlowTcpEcnEcho) FlowTcp
	// HasEcnEcho checks if EcnEcho has been set in FlowTcp
	HasEcnEcho() bool
	// CtlUrg returns PatternFlowTcpCtlUrg, set in FlowTcp.
	// PatternFlowTcpCtlUrg is a value of 1 indicates that the urgent pointer field is significant.
	CtlUrg() PatternFlowTcpCtlUrg
	// SetCtlUrg assigns PatternFlowTcpCtlUrg provided by user to FlowTcp.
	// PatternFlowTcpCtlUrg is a value of 1 indicates that the urgent pointer field is significant.
	SetCtlUrg(value PatternFlowTcpCtlUrg) FlowTcp
	// HasCtlUrg checks if CtlUrg has been set in FlowTcp
	HasCtlUrg() bool
	// CtlAck returns PatternFlowTcpCtlAck, set in FlowTcp.
	// PatternFlowTcpCtlAck is a value of 1 indicates that the ackknowledgment field is significant.
	CtlAck() PatternFlowTcpCtlAck
	// SetCtlAck assigns PatternFlowTcpCtlAck provided by user to FlowTcp.
	// PatternFlowTcpCtlAck is a value of 1 indicates that the ackknowledgment field is significant.
	SetCtlAck(value PatternFlowTcpCtlAck) FlowTcp
	// HasCtlAck checks if CtlAck has been set in FlowTcp
	HasCtlAck() bool
	// CtlPsh returns PatternFlowTcpCtlPsh, set in FlowTcp.
	// PatternFlowTcpCtlPsh is asks to push the buffered data to the receiving application.
	CtlPsh() PatternFlowTcpCtlPsh
	// SetCtlPsh assigns PatternFlowTcpCtlPsh provided by user to FlowTcp.
	// PatternFlowTcpCtlPsh is asks to push the buffered data to the receiving application.
	SetCtlPsh(value PatternFlowTcpCtlPsh) FlowTcp
	// HasCtlPsh checks if CtlPsh has been set in FlowTcp
	HasCtlPsh() bool
	// CtlRst returns PatternFlowTcpCtlRst, set in FlowTcp.
	// PatternFlowTcpCtlRst is reset the connection.
	CtlRst() PatternFlowTcpCtlRst
	// SetCtlRst assigns PatternFlowTcpCtlRst provided by user to FlowTcp.
	// PatternFlowTcpCtlRst is reset the connection.
	SetCtlRst(value PatternFlowTcpCtlRst) FlowTcp
	// HasCtlRst checks if CtlRst has been set in FlowTcp
	HasCtlRst() bool
	// CtlSyn returns PatternFlowTcpCtlSyn, set in FlowTcp.
	// PatternFlowTcpCtlSyn is synchronize sequenece numbers.
	CtlSyn() PatternFlowTcpCtlSyn
	// SetCtlSyn assigns PatternFlowTcpCtlSyn provided by user to FlowTcp.
	// PatternFlowTcpCtlSyn is synchronize sequenece numbers.
	SetCtlSyn(value PatternFlowTcpCtlSyn) FlowTcp
	// HasCtlSyn checks if CtlSyn has been set in FlowTcp
	HasCtlSyn() bool
	// CtlFin returns PatternFlowTcpCtlFin, set in FlowTcp.
	// PatternFlowTcpCtlFin is last packet from the sender.
	CtlFin() PatternFlowTcpCtlFin
	// SetCtlFin assigns PatternFlowTcpCtlFin provided by user to FlowTcp.
	// PatternFlowTcpCtlFin is last packet from the sender.
	SetCtlFin(value PatternFlowTcpCtlFin) FlowTcp
	// HasCtlFin checks if CtlFin has been set in FlowTcp
	HasCtlFin() bool
	// Window returns PatternFlowTcpWindow, set in FlowTcp.
	// PatternFlowTcpWindow is tcp connection window.
	Window() PatternFlowTcpWindow
	// SetWindow assigns PatternFlowTcpWindow provided by user to FlowTcp.
	// PatternFlowTcpWindow is tcp connection window.
	SetWindow(value PatternFlowTcpWindow) FlowTcp
	// HasWindow checks if Window has been set in FlowTcp
	HasWindow() bool
	// Checksum returns PatternFlowTcpChecksum, set in FlowTcp.
	// PatternFlowTcpChecksum is the one's complement of the one's complement sum of all 16 bit words in header and text.  An all-zero value means that no checksum will be transmitted.   While computing the checksum, the checksum field itself is replaced with zeros.
	Checksum() PatternFlowTcpChecksum
	// SetChecksum assigns PatternFlowTcpChecksum provided by user to FlowTcp.
	// PatternFlowTcpChecksum is the one's complement of the one's complement sum of all 16 bit words in header and text.  An all-zero value means that no checksum will be transmitted.   While computing the checksum, the checksum field itself is replaced with zeros.
	SetChecksum(value PatternFlowTcpChecksum) FlowTcp
	// HasChecksum checks if Checksum has been set in FlowTcp
	HasChecksum() bool
	setNil()
}

// description is TBD
// SrcPort returns a PatternFlowTcpSrcPort
func (obj *flowTcp) SrcPort() PatternFlowTcpSrcPort {
	if obj.obj.SrcPort == nil {
		obj.obj.SrcPort = NewPatternFlowTcpSrcPort().msg()
	}
	if obj.srcPortHolder == nil {
		obj.srcPortHolder = &patternFlowTcpSrcPort{obj: obj.obj.SrcPort}
	}
	return obj.srcPortHolder
}

// description is TBD
// SrcPort returns a PatternFlowTcpSrcPort
func (obj *flowTcp) HasSrcPort() bool {
	return obj.obj.SrcPort != nil
}

// description is TBD
// SetSrcPort sets the PatternFlowTcpSrcPort value in the FlowTcp object
func (obj *flowTcp) SetSrcPort(value PatternFlowTcpSrcPort) FlowTcp {

	obj.srcPortHolder = nil
	obj.obj.SrcPort = value.msg()

	return obj
}

// description is TBD
// DstPort returns a PatternFlowTcpDstPort
func (obj *flowTcp) DstPort() PatternFlowTcpDstPort {
	if obj.obj.DstPort == nil {
		obj.obj.DstPort = NewPatternFlowTcpDstPort().msg()
	}
	if obj.dstPortHolder == nil {
		obj.dstPortHolder = &patternFlowTcpDstPort{obj: obj.obj.DstPort}
	}
	return obj.dstPortHolder
}

// description is TBD
// DstPort returns a PatternFlowTcpDstPort
func (obj *flowTcp) HasDstPort() bool {
	return obj.obj.DstPort != nil
}

// description is TBD
// SetDstPort sets the PatternFlowTcpDstPort value in the FlowTcp object
func (obj *flowTcp) SetDstPort(value PatternFlowTcpDstPort) FlowTcp {

	obj.dstPortHolder = nil
	obj.obj.DstPort = value.msg()

	return obj
}

// description is TBD
// SeqNum returns a PatternFlowTcpSeqNum
func (obj *flowTcp) SeqNum() PatternFlowTcpSeqNum {
	if obj.obj.SeqNum == nil {
		obj.obj.SeqNum = NewPatternFlowTcpSeqNum().msg()
	}
	if obj.seqNumHolder == nil {
		obj.seqNumHolder = &patternFlowTcpSeqNum{obj: obj.obj.SeqNum}
	}
	return obj.seqNumHolder
}

// description is TBD
// SeqNum returns a PatternFlowTcpSeqNum
func (obj *flowTcp) HasSeqNum() bool {
	return obj.obj.SeqNum != nil
}

// description is TBD
// SetSeqNum sets the PatternFlowTcpSeqNum value in the FlowTcp object
func (obj *flowTcp) SetSeqNum(value PatternFlowTcpSeqNum) FlowTcp {

	obj.seqNumHolder = nil
	obj.obj.SeqNum = value.msg()

	return obj
}

// description is TBD
// AckNum returns a PatternFlowTcpAckNum
func (obj *flowTcp) AckNum() PatternFlowTcpAckNum {
	if obj.obj.AckNum == nil {
		obj.obj.AckNum = NewPatternFlowTcpAckNum().msg()
	}
	if obj.ackNumHolder == nil {
		obj.ackNumHolder = &patternFlowTcpAckNum{obj: obj.obj.AckNum}
	}
	return obj.ackNumHolder
}

// description is TBD
// AckNum returns a PatternFlowTcpAckNum
func (obj *flowTcp) HasAckNum() bool {
	return obj.obj.AckNum != nil
}

// description is TBD
// SetAckNum sets the PatternFlowTcpAckNum value in the FlowTcp object
func (obj *flowTcp) SetAckNum(value PatternFlowTcpAckNum) FlowTcp {

	obj.ackNumHolder = nil
	obj.obj.AckNum = value.msg()

	return obj
}

// description is TBD
// DataOffset returns a PatternFlowTcpDataOffset
func (obj *flowTcp) DataOffset() PatternFlowTcpDataOffset {
	if obj.obj.DataOffset == nil {
		obj.obj.DataOffset = NewPatternFlowTcpDataOffset().msg()
	}
	if obj.dataOffsetHolder == nil {
		obj.dataOffsetHolder = &patternFlowTcpDataOffset{obj: obj.obj.DataOffset}
	}
	return obj.dataOffsetHolder
}

// description is TBD
// DataOffset returns a PatternFlowTcpDataOffset
func (obj *flowTcp) HasDataOffset() bool {
	return obj.obj.DataOffset != nil
}

// description is TBD
// SetDataOffset sets the PatternFlowTcpDataOffset value in the FlowTcp object
func (obj *flowTcp) SetDataOffset(value PatternFlowTcpDataOffset) FlowTcp {

	obj.dataOffsetHolder = nil
	obj.obj.DataOffset = value.msg()

	return obj
}

// description is TBD
// EcnNs returns a PatternFlowTcpEcnNs
func (obj *flowTcp) EcnNs() PatternFlowTcpEcnNs {
	if obj.obj.EcnNs == nil {
		obj.obj.EcnNs = NewPatternFlowTcpEcnNs().msg()
	}
	if obj.ecnNsHolder == nil {
		obj.ecnNsHolder = &patternFlowTcpEcnNs{obj: obj.obj.EcnNs}
	}
	return obj.ecnNsHolder
}

// description is TBD
// EcnNs returns a PatternFlowTcpEcnNs
func (obj *flowTcp) HasEcnNs() bool {
	return obj.obj.EcnNs != nil
}

// description is TBD
// SetEcnNs sets the PatternFlowTcpEcnNs value in the FlowTcp object
func (obj *flowTcp) SetEcnNs(value PatternFlowTcpEcnNs) FlowTcp {

	obj.ecnNsHolder = nil
	obj.obj.EcnNs = value.msg()

	return obj
}

// description is TBD
// EcnCwr returns a PatternFlowTcpEcnCwr
func (obj *flowTcp) EcnCwr() PatternFlowTcpEcnCwr {
	if obj.obj.EcnCwr == nil {
		obj.obj.EcnCwr = NewPatternFlowTcpEcnCwr().msg()
	}
	if obj.ecnCwrHolder == nil {
		obj.ecnCwrHolder = &patternFlowTcpEcnCwr{obj: obj.obj.EcnCwr}
	}
	return obj.ecnCwrHolder
}

// description is TBD
// EcnCwr returns a PatternFlowTcpEcnCwr
func (obj *flowTcp) HasEcnCwr() bool {
	return obj.obj.EcnCwr != nil
}

// description is TBD
// SetEcnCwr sets the PatternFlowTcpEcnCwr value in the FlowTcp object
func (obj *flowTcp) SetEcnCwr(value PatternFlowTcpEcnCwr) FlowTcp {

	obj.ecnCwrHolder = nil
	obj.obj.EcnCwr = value.msg()

	return obj
}

// description is TBD
// EcnEcho returns a PatternFlowTcpEcnEcho
func (obj *flowTcp) EcnEcho() PatternFlowTcpEcnEcho {
	if obj.obj.EcnEcho == nil {
		obj.obj.EcnEcho = NewPatternFlowTcpEcnEcho().msg()
	}
	if obj.ecnEchoHolder == nil {
		obj.ecnEchoHolder = &patternFlowTcpEcnEcho{obj: obj.obj.EcnEcho}
	}
	return obj.ecnEchoHolder
}

// description is TBD
// EcnEcho returns a PatternFlowTcpEcnEcho
func (obj *flowTcp) HasEcnEcho() bool {
	return obj.obj.EcnEcho != nil
}

// description is TBD
// SetEcnEcho sets the PatternFlowTcpEcnEcho value in the FlowTcp object
func (obj *flowTcp) SetEcnEcho(value PatternFlowTcpEcnEcho) FlowTcp {

	obj.ecnEchoHolder = nil
	obj.obj.EcnEcho = value.msg()

	return obj
}

// description is TBD
// CtlUrg returns a PatternFlowTcpCtlUrg
func (obj *flowTcp) CtlUrg() PatternFlowTcpCtlUrg {
	if obj.obj.CtlUrg == nil {
		obj.obj.CtlUrg = NewPatternFlowTcpCtlUrg().msg()
	}
	if obj.ctlUrgHolder == nil {
		obj.ctlUrgHolder = &patternFlowTcpCtlUrg{obj: obj.obj.CtlUrg}
	}
	return obj.ctlUrgHolder
}

// description is TBD
// CtlUrg returns a PatternFlowTcpCtlUrg
func (obj *flowTcp) HasCtlUrg() bool {
	return obj.obj.CtlUrg != nil
}

// description is TBD
// SetCtlUrg sets the PatternFlowTcpCtlUrg value in the FlowTcp object
func (obj *flowTcp) SetCtlUrg(value PatternFlowTcpCtlUrg) FlowTcp {

	obj.ctlUrgHolder = nil
	obj.obj.CtlUrg = value.msg()

	return obj
}

// description is TBD
// CtlAck returns a PatternFlowTcpCtlAck
func (obj *flowTcp) CtlAck() PatternFlowTcpCtlAck {
	if obj.obj.CtlAck == nil {
		obj.obj.CtlAck = NewPatternFlowTcpCtlAck().msg()
	}
	if obj.ctlAckHolder == nil {
		obj.ctlAckHolder = &patternFlowTcpCtlAck{obj: obj.obj.CtlAck}
	}
	return obj.ctlAckHolder
}

// description is TBD
// CtlAck returns a PatternFlowTcpCtlAck
func (obj *flowTcp) HasCtlAck() bool {
	return obj.obj.CtlAck != nil
}

// description is TBD
// SetCtlAck sets the PatternFlowTcpCtlAck value in the FlowTcp object
func (obj *flowTcp) SetCtlAck(value PatternFlowTcpCtlAck) FlowTcp {

	obj.ctlAckHolder = nil
	obj.obj.CtlAck = value.msg()

	return obj
}

// description is TBD
// CtlPsh returns a PatternFlowTcpCtlPsh
func (obj *flowTcp) CtlPsh() PatternFlowTcpCtlPsh {
	if obj.obj.CtlPsh == nil {
		obj.obj.CtlPsh = NewPatternFlowTcpCtlPsh().msg()
	}
	if obj.ctlPshHolder == nil {
		obj.ctlPshHolder = &patternFlowTcpCtlPsh{obj: obj.obj.CtlPsh}
	}
	return obj.ctlPshHolder
}

// description is TBD
// CtlPsh returns a PatternFlowTcpCtlPsh
func (obj *flowTcp) HasCtlPsh() bool {
	return obj.obj.CtlPsh != nil
}

// description is TBD
// SetCtlPsh sets the PatternFlowTcpCtlPsh value in the FlowTcp object
func (obj *flowTcp) SetCtlPsh(value PatternFlowTcpCtlPsh) FlowTcp {

	obj.ctlPshHolder = nil
	obj.obj.CtlPsh = value.msg()

	return obj
}

// description is TBD
// CtlRst returns a PatternFlowTcpCtlRst
func (obj *flowTcp) CtlRst() PatternFlowTcpCtlRst {
	if obj.obj.CtlRst == nil {
		obj.obj.CtlRst = NewPatternFlowTcpCtlRst().msg()
	}
	if obj.ctlRstHolder == nil {
		obj.ctlRstHolder = &patternFlowTcpCtlRst{obj: obj.obj.CtlRst}
	}
	return obj.ctlRstHolder
}

// description is TBD
// CtlRst returns a PatternFlowTcpCtlRst
func (obj *flowTcp) HasCtlRst() bool {
	return obj.obj.CtlRst != nil
}

// description is TBD
// SetCtlRst sets the PatternFlowTcpCtlRst value in the FlowTcp object
func (obj *flowTcp) SetCtlRst(value PatternFlowTcpCtlRst) FlowTcp {

	obj.ctlRstHolder = nil
	obj.obj.CtlRst = value.msg()

	return obj
}

// description is TBD
// CtlSyn returns a PatternFlowTcpCtlSyn
func (obj *flowTcp) CtlSyn() PatternFlowTcpCtlSyn {
	if obj.obj.CtlSyn == nil {
		obj.obj.CtlSyn = NewPatternFlowTcpCtlSyn().msg()
	}
	if obj.ctlSynHolder == nil {
		obj.ctlSynHolder = &patternFlowTcpCtlSyn{obj: obj.obj.CtlSyn}
	}
	return obj.ctlSynHolder
}

// description is TBD
// CtlSyn returns a PatternFlowTcpCtlSyn
func (obj *flowTcp) HasCtlSyn() bool {
	return obj.obj.CtlSyn != nil
}

// description is TBD
// SetCtlSyn sets the PatternFlowTcpCtlSyn value in the FlowTcp object
func (obj *flowTcp) SetCtlSyn(value PatternFlowTcpCtlSyn) FlowTcp {

	obj.ctlSynHolder = nil
	obj.obj.CtlSyn = value.msg()

	return obj
}

// description is TBD
// CtlFin returns a PatternFlowTcpCtlFin
func (obj *flowTcp) CtlFin() PatternFlowTcpCtlFin {
	if obj.obj.CtlFin == nil {
		obj.obj.CtlFin = NewPatternFlowTcpCtlFin().msg()
	}
	if obj.ctlFinHolder == nil {
		obj.ctlFinHolder = &patternFlowTcpCtlFin{obj: obj.obj.CtlFin}
	}
	return obj.ctlFinHolder
}

// description is TBD
// CtlFin returns a PatternFlowTcpCtlFin
func (obj *flowTcp) HasCtlFin() bool {
	return obj.obj.CtlFin != nil
}

// description is TBD
// SetCtlFin sets the PatternFlowTcpCtlFin value in the FlowTcp object
func (obj *flowTcp) SetCtlFin(value PatternFlowTcpCtlFin) FlowTcp {

	obj.ctlFinHolder = nil
	obj.obj.CtlFin = value.msg()

	return obj
}

// description is TBD
// Window returns a PatternFlowTcpWindow
func (obj *flowTcp) Window() PatternFlowTcpWindow {
	if obj.obj.Window == nil {
		obj.obj.Window = NewPatternFlowTcpWindow().msg()
	}
	if obj.windowHolder == nil {
		obj.windowHolder = &patternFlowTcpWindow{obj: obj.obj.Window}
	}
	return obj.windowHolder
}

// description is TBD
// Window returns a PatternFlowTcpWindow
func (obj *flowTcp) HasWindow() bool {
	return obj.obj.Window != nil
}

// description is TBD
// SetWindow sets the PatternFlowTcpWindow value in the FlowTcp object
func (obj *flowTcp) SetWindow(value PatternFlowTcpWindow) FlowTcp {

	obj.windowHolder = nil
	obj.obj.Window = value.msg()

	return obj
}

// description is TBD
// Checksum returns a PatternFlowTcpChecksum
func (obj *flowTcp) Checksum() PatternFlowTcpChecksum {
	if obj.obj.Checksum == nil {
		obj.obj.Checksum = NewPatternFlowTcpChecksum().msg()
	}
	if obj.checksumHolder == nil {
		obj.checksumHolder = &patternFlowTcpChecksum{obj: obj.obj.Checksum}
	}
	return obj.checksumHolder
}

// description is TBD
// Checksum returns a PatternFlowTcpChecksum
func (obj *flowTcp) HasChecksum() bool {
	return obj.obj.Checksum != nil
}

// description is TBD
// SetChecksum sets the PatternFlowTcpChecksum value in the FlowTcp object
func (obj *flowTcp) SetChecksum(value PatternFlowTcpChecksum) FlowTcp {

	obj.checksumHolder = nil
	obj.obj.Checksum = value.msg()

	return obj
}

func (obj *flowTcp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SrcPort != nil {

		obj.SrcPort().validateObj(vObj, set_default)
	}

	if obj.obj.DstPort != nil {

		obj.DstPort().validateObj(vObj, set_default)
	}

	if obj.obj.SeqNum != nil {

		obj.SeqNum().validateObj(vObj, set_default)
	}

	if obj.obj.AckNum != nil {

		obj.AckNum().validateObj(vObj, set_default)
	}

	if obj.obj.DataOffset != nil {

		obj.DataOffset().validateObj(vObj, set_default)
	}

	if obj.obj.EcnNs != nil {

		obj.EcnNs().validateObj(vObj, set_default)
	}

	if obj.obj.EcnCwr != nil {

		obj.EcnCwr().validateObj(vObj, set_default)
	}

	if obj.obj.EcnEcho != nil {

		obj.EcnEcho().validateObj(vObj, set_default)
	}

	if obj.obj.CtlUrg != nil {

		obj.CtlUrg().validateObj(vObj, set_default)
	}

	if obj.obj.CtlAck != nil {

		obj.CtlAck().validateObj(vObj, set_default)
	}

	if obj.obj.CtlPsh != nil {

		obj.CtlPsh().validateObj(vObj, set_default)
	}

	if obj.obj.CtlRst != nil {

		obj.CtlRst().validateObj(vObj, set_default)
	}

	if obj.obj.CtlSyn != nil {

		obj.CtlSyn().validateObj(vObj, set_default)
	}

	if obj.obj.CtlFin != nil {

		obj.CtlFin().validateObj(vObj, set_default)
	}

	if obj.obj.Window != nil {

		obj.Window().validateObj(vObj, set_default)
	}

	if obj.obj.Checksum != nil {

		obj.Checksum().validateObj(vObj, set_default)
	}

}

func (obj *flowTcp) setDefault() {

}
