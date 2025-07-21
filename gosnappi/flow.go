package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Flow *****
type flow struct {
	validation
	obj                *otg.Flow
	marshaller         marshalFlow
	unMarshaller       unMarshalFlow
	txRxHolder         FlowTxRx
	packetHolder       FlowFlowHeaderIter
	egressPacketHolder FlowFlowHeaderIter
	sizeHolder         FlowSize
	rateHolder         FlowRate
	durationHolder     FlowDuration
	metricsHolder      FlowMetrics
}

func NewFlow() Flow {
	obj := flow{obj: &otg.Flow{}}
	obj.setDefault()
	return &obj
}

func (obj *flow) msg() *otg.Flow {
	return obj.obj
}

func (obj *flow) setMsg(msg *otg.Flow) Flow {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflow struct {
	obj *flow
}

type marshalFlow interface {
	// ToProto marshals Flow to protobuf object *otg.Flow
	ToProto() (*otg.Flow, error)
	// ToPbText marshals Flow to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Flow to YAML text
	ToYaml() (string, error)
	// ToJson marshals Flow to JSON text
	ToJson() (string, error)
}

type unMarshalflow struct {
	obj *flow
}

type unMarshalFlow interface {
	// FromProto unmarshals Flow from protobuf object *otg.Flow
	FromProto(msg *otg.Flow) (Flow, error)
	// FromPbText unmarshals Flow from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Flow from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Flow from JSON text
	FromJson(value string) error
}

func (obj *flow) Marshal() marshalFlow {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflow{obj: obj}
	}
	return obj.marshaller
}

func (obj *flow) Unmarshal() unMarshalFlow {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflow{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflow) ToProto() (*otg.Flow, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflow) FromProto(msg *otg.Flow) (Flow, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflow) ToPbText() (string, error) {
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

func (m *unMarshalflow) FromPbText(value string) error {
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

func (m *marshalflow) ToYaml() (string, error) {
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

func (m *unMarshalflow) FromYaml(value string) error {
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

func (m *marshalflow) ToJson() (string, error) {
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

func (m *unMarshalflow) FromJson(value string) error {
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

func (obj *flow) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flow) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flow) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flow) Clone() (Flow, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlow()
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

func (obj *flow) setNil() {
	obj.txRxHolder = nil
	obj.packetHolder = nil
	obj.egressPacketHolder = nil
	obj.sizeHolder = nil
	obj.rateHolder = nil
	obj.durationHolder = nil
	obj.metricsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Flow is a high level data plane traffic flow.
type Flow interface {
	Validation
	// msg marshals Flow to protobuf object *otg.Flow
	// and doesn't set defaults
	msg() *otg.Flow
	// setMsg unmarshals Flow from protobuf object *otg.Flow
	// and doesn't set defaults
	setMsg(*otg.Flow) Flow
	// provides marshal interface
	Marshal() marshalFlow
	// provides unmarshal interface
	Unmarshal() unMarshalFlow
	// validate validates Flow
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Flow, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// TxRx returns FlowTxRx, set in Flow.
	// FlowTxRx is a container for different types of transmit and receive
	// endpoint containers.
	TxRx() FlowTxRx
	// SetTxRx assigns FlowTxRx provided by user to Flow.
	// FlowTxRx is a container for different types of transmit and receive
	// endpoint containers.
	SetTxRx(value FlowTxRx) Flow
	// Packet returns FlowFlowHeaderIterIter, set in Flow
	Packet() FlowFlowHeaderIter
	// EgressPacket returns FlowFlowHeaderIterIter, set in Flow
	EgressPacket() FlowFlowHeaderIter
	// Size returns FlowSize, set in Flow.
	// FlowSize is the frame size which overrides the total length of the packet
	Size() FlowSize
	// SetSize assigns FlowSize provided by user to Flow.
	// FlowSize is the frame size which overrides the total length of the packet
	SetSize(value FlowSize) Flow
	// HasSize checks if Size has been set in Flow
	HasSize() bool
	// Rate returns FlowRate, set in Flow.
	// FlowRate is the rate of packet transmission
	Rate() FlowRate
	// SetRate assigns FlowRate provided by user to Flow.
	// FlowRate is the rate of packet transmission
	SetRate(value FlowRate) Flow
	// HasRate checks if Rate has been set in Flow
	HasRate() bool
	// Duration returns FlowDuration, set in Flow.
	// FlowDuration is a container for different transmit durations.
	Duration() FlowDuration
	// SetDuration assigns FlowDuration provided by user to Flow.
	// FlowDuration is a container for different transmit durations.
	SetDuration(value FlowDuration) Flow
	// HasDuration checks if Duration has been set in Flow
	HasDuration() bool
	// Metrics returns FlowMetrics, set in Flow.
	// FlowMetrics is the optional container for configuring flow metrics.
	Metrics() FlowMetrics
	// SetMetrics assigns FlowMetrics provided by user to Flow.
	// FlowMetrics is the optional container for configuring flow metrics.
	SetMetrics(value FlowMetrics) Flow
	// HasMetrics checks if Metrics has been set in Flow
	HasMetrics() bool
	// Name returns string, set in Flow.
	Name() string
	// SetName assigns string provided by user to Flow
	SetName(value string) Flow
	setNil()
}

// The transmit and receive endpoints.
// TxRx returns a FlowTxRx
func (obj *flow) TxRx() FlowTxRx {
	if obj.obj.TxRx == nil {
		obj.obj.TxRx = NewFlowTxRx().msg()
	}
	if obj.txRxHolder == nil {
		obj.txRxHolder = &flowTxRx{obj: obj.obj.TxRx}
	}
	return obj.txRxHolder
}

// The transmit and receive endpoints.
// SetTxRx sets the FlowTxRx value in the Flow object
func (obj *flow) SetTxRx(value FlowTxRx) Flow {

	obj.txRxHolder = nil
	obj.obj.TxRx = value.msg()

	return obj
}

// The list of protocol headers defining the shape of all
// intended packets in corresponding flow as it is transmitted
// by traffic-generator port.
//
// The order of protocol headers assigned to the list is the
// order they will appear on the wire.
//
// In the case of an empty list the keyword/value of minItems: 1
// indicates that an implementation MUST provide at least one
// Flow.Header object.
//
// The default value for the Flow.Header choice property is ethernet
// which will result in an implementation by default providing at least
// one ethernet packet header.
// Packet returns a []FlowHeader
func (obj *flow) Packet() FlowFlowHeaderIter {
	if len(obj.obj.Packet) == 0 {
		obj.obj.Packet = []*otg.FlowHeader{}
	}
	if obj.packetHolder == nil {
		obj.packetHolder = newFlowFlowHeaderIter(&obj.obj.Packet).setMsg(obj)
	}
	return obj.packetHolder
}

type flowFlowHeaderIter struct {
	obj             *flow
	flowHeaderSlice []FlowHeader
	fieldPtr        *[]*otg.FlowHeader
}

func newFlowFlowHeaderIter(ptr *[]*otg.FlowHeader) FlowFlowHeaderIter {
	return &flowFlowHeaderIter{fieldPtr: ptr}
}

type FlowFlowHeaderIter interface {
	setMsg(*flow) FlowFlowHeaderIter
	Items() []FlowHeader
	Add() FlowHeader
	Append(items ...FlowHeader) FlowFlowHeaderIter
	Set(index int, newObj FlowHeader) FlowFlowHeaderIter
	Clear() FlowFlowHeaderIter
	clearHolderSlice() FlowFlowHeaderIter
	appendHolderSlice(item FlowHeader) FlowFlowHeaderIter
}

func (obj *flowFlowHeaderIter) setMsg(msg *flow) FlowFlowHeaderIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowHeader{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowFlowHeaderIter) Items() []FlowHeader {
	return obj.flowHeaderSlice
}

func (obj *flowFlowHeaderIter) Add() FlowHeader {
	newObj := &otg.FlowHeader{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowHeader{obj: newObj}
	newLibObj.setDefault()
	obj.flowHeaderSlice = append(obj.flowHeaderSlice, newLibObj)
	return newLibObj
}

func (obj *flowFlowHeaderIter) Append(items ...FlowHeader) FlowFlowHeaderIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowHeaderSlice = append(obj.flowHeaderSlice, item)
	}
	return obj
}

func (obj *flowFlowHeaderIter) Set(index int, newObj FlowHeader) FlowFlowHeaderIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowHeaderSlice[index] = newObj
	return obj
}
func (obj *flowFlowHeaderIter) Clear() FlowFlowHeaderIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowHeader{}
		obj.flowHeaderSlice = []FlowHeader{}
	}
	return obj
}
func (obj *flowFlowHeaderIter) clearHolderSlice() FlowFlowHeaderIter {
	if len(obj.flowHeaderSlice) > 0 {
		obj.flowHeaderSlice = []FlowHeader{}
	}
	return obj
}
func (obj *flowFlowHeaderIter) appendHolderSlice(item FlowHeader) FlowFlowHeaderIter {
	obj.flowHeaderSlice = append(obj.flowHeaderSlice, item)
	return obj
}

// Under Review: The packet header schema for egress tracking currently exposes unwanted fields. The query structure for tagged metrics inside flows metrics requires documenting expected response format.
//
// Under Review: The packet header schema for egress tracking currently exposes unwanted fields. The query structure for tagged metrics inside flows metrics requires documenting expected response format.
//
// The list of protocol headers defining the shape of all
// intended packets in corresponding flow as it is received
// by traffic-generator port.
//
// For all protocol headers, only the `metric_tags` property is configurable.
// EgressPacket returns a []FlowHeader
func (obj *flow) EgressPacket() FlowFlowHeaderIter {
	if len(obj.obj.EgressPacket) == 0 {
		obj.obj.EgressPacket = []*otg.FlowHeader{}
	}
	if obj.egressPacketHolder == nil {
		obj.egressPacketHolder = newFlowFlowHeaderIter(&obj.obj.EgressPacket).setMsg(obj)
	}
	return obj.egressPacketHolder
}

// The size of the packets.
// Size returns a FlowSize
func (obj *flow) Size() FlowSize {
	if obj.obj.Size == nil {
		obj.obj.Size = NewFlowSize().msg()
	}
	if obj.sizeHolder == nil {
		obj.sizeHolder = &flowSize{obj: obj.obj.Size}
	}
	return obj.sizeHolder
}

// The size of the packets.
// Size returns a FlowSize
func (obj *flow) HasSize() bool {
	return obj.obj.Size != nil
}

// The size of the packets.
// SetSize sets the FlowSize value in the Flow object
func (obj *flow) SetSize(value FlowSize) Flow {

	obj.sizeHolder = nil
	obj.obj.Size = value.msg()

	return obj
}

// The transmit rate of the packets.
// Rate returns a FlowRate
func (obj *flow) Rate() FlowRate {
	if obj.obj.Rate == nil {
		obj.obj.Rate = NewFlowRate().msg()
	}
	if obj.rateHolder == nil {
		obj.rateHolder = &flowRate{obj: obj.obj.Rate}
	}
	return obj.rateHolder
}

// The transmit rate of the packets.
// Rate returns a FlowRate
func (obj *flow) HasRate() bool {
	return obj.obj.Rate != nil
}

// The transmit rate of the packets.
// SetRate sets the FlowRate value in the Flow object
func (obj *flow) SetRate(value FlowRate) Flow {

	obj.rateHolder = nil
	obj.obj.Rate = value.msg()

	return obj
}

// The transmit duration of the packets.
// Duration returns a FlowDuration
func (obj *flow) Duration() FlowDuration {
	if obj.obj.Duration == nil {
		obj.obj.Duration = NewFlowDuration().msg()
	}
	if obj.durationHolder == nil {
		obj.durationHolder = &flowDuration{obj: obj.obj.Duration}
	}
	return obj.durationHolder
}

// The transmit duration of the packets.
// Duration returns a FlowDuration
func (obj *flow) HasDuration() bool {
	return obj.obj.Duration != nil
}

// The transmit duration of the packets.
// SetDuration sets the FlowDuration value in the Flow object
func (obj *flow) SetDuration(value FlowDuration) Flow {

	obj.durationHolder = nil
	obj.obj.Duration = value.msg()

	return obj
}

// Flow metrics.
// Metrics returns a FlowMetrics
func (obj *flow) Metrics() FlowMetrics {
	if obj.obj.Metrics == nil {
		obj.obj.Metrics = NewFlowMetrics().msg()
	}
	if obj.metricsHolder == nil {
		obj.metricsHolder = &flowMetrics{obj: obj.obj.Metrics}
	}
	return obj.metricsHolder
}

// Flow metrics.
// Metrics returns a FlowMetrics
func (obj *flow) HasMetrics() bool {
	return obj.obj.Metrics != nil
}

// Flow metrics.
// SetMetrics sets the FlowMetrics value in the Flow object
func (obj *flow) SetMetrics(value FlowMetrics) Flow {

	obj.metricsHolder = nil
	obj.obj.Metrics = value.msg()

	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *flow) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the Flow object
func (obj *flow) SetName(value string) Flow {

	obj.obj.Name = &value
	return obj
}

func (obj *flow) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// TxRx is required
	if obj.obj.TxRx == nil {
		vObj.validationErrors = append(vObj.validationErrors, "TxRx is required field on interface Flow")
	}

	if obj.obj.TxRx != nil {

		obj.TxRx().validateObj(vObj, set_default)
	}

	if len(obj.obj.Packet) != 0 {

		if set_default {
			obj.Packet().clearHolderSlice()
			for _, item := range obj.obj.Packet {
				obj.Packet().appendHolderSlice(&flowHeader{obj: item})
			}
		}
		for _, item := range obj.Packet().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.EgressPacket) != 0 {

		if set_default {
			obj.EgressPacket().clearHolderSlice()
			for _, item := range obj.obj.EgressPacket {
				obj.EgressPacket().appendHolderSlice(&flowHeader{obj: item})
			}
		}
		for _, item := range obj.EgressPacket().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.Size != nil {

		obj.Size().validateObj(vObj, set_default)
	}

	if obj.obj.Rate != nil {

		obj.Rate().validateObj(vObj, set_default)
	}

	if obj.obj.Duration != nil {

		obj.Duration().validateObj(vObj, set_default)
	}

	if obj.obj.Metrics != nil {

		obj.Metrics().validateObj(vObj, set_default)
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface Flow")
	}
}

func (obj *flow) setDefault() {

}
