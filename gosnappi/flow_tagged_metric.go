package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowTaggedMetric *****
type flowTaggedMetric struct {
	validation
	obj              *otg.FlowTaggedMetric
	marshaller       marshalFlowTaggedMetric
	unMarshaller     unMarshalFlowTaggedMetric
	tagsHolder       FlowTaggedMetricFlowMetricTagIter
	timestampsHolder MetricTimestamp
	latencyHolder    MetricLatency
}

func NewFlowTaggedMetric() FlowTaggedMetric {
	obj := flowTaggedMetric{obj: &otg.FlowTaggedMetric{}}
	obj.setDefault()
	return &obj
}

func (obj *flowTaggedMetric) msg() *otg.FlowTaggedMetric {
	return obj.obj
}

func (obj *flowTaggedMetric) setMsg(msg *otg.FlowTaggedMetric) FlowTaggedMetric {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowTaggedMetric struct {
	obj *flowTaggedMetric
}

type marshalFlowTaggedMetric interface {
	// ToProto marshals FlowTaggedMetric to protobuf object *otg.FlowTaggedMetric
	ToProto() (*otg.FlowTaggedMetric, error)
	// ToPbText marshals FlowTaggedMetric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowTaggedMetric to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowTaggedMetric to JSON text
	ToJson() (string, error)
}

type unMarshalflowTaggedMetric struct {
	obj *flowTaggedMetric
}

type unMarshalFlowTaggedMetric interface {
	// FromProto unmarshals FlowTaggedMetric from protobuf object *otg.FlowTaggedMetric
	FromProto(msg *otg.FlowTaggedMetric) (FlowTaggedMetric, error)
	// FromPbText unmarshals FlowTaggedMetric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowTaggedMetric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowTaggedMetric from JSON text
	FromJson(value string) error
}

func (obj *flowTaggedMetric) Marshal() marshalFlowTaggedMetric {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowTaggedMetric{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowTaggedMetric) Unmarshal() unMarshalFlowTaggedMetric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowTaggedMetric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowTaggedMetric) ToProto() (*otg.FlowTaggedMetric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowTaggedMetric) FromProto(msg *otg.FlowTaggedMetric) (FlowTaggedMetric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowTaggedMetric) ToPbText() (string, error) {
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

func (m *unMarshalflowTaggedMetric) FromPbText(value string) error {
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

func (m *marshalflowTaggedMetric) ToYaml() (string, error) {
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

func (m *unMarshalflowTaggedMetric) FromYaml(value string) error {
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

func (m *marshalflowTaggedMetric) ToJson() (string, error) {
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

func (m *unMarshalflowTaggedMetric) FromJson(value string) error {
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

func (obj *flowTaggedMetric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowTaggedMetric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowTaggedMetric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowTaggedMetric) Clone() (FlowTaggedMetric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowTaggedMetric()
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

func (obj *flowTaggedMetric) setNil() {
	obj.tagsHolder = nil
	obj.timestampsHolder = nil
	obj.latencyHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowTaggedMetric is metrics for each set of values applicable for configured
// metric tags in ingress or egress packet header fields of corresponding flow.
// The container is keyed by list of tag-value pairs.
type FlowTaggedMetric interface {
	Validation
	// msg marshals FlowTaggedMetric to protobuf object *otg.FlowTaggedMetric
	// and doesn't set defaults
	msg() *otg.FlowTaggedMetric
	// setMsg unmarshals FlowTaggedMetric from protobuf object *otg.FlowTaggedMetric
	// and doesn't set defaults
	setMsg(*otg.FlowTaggedMetric) FlowTaggedMetric
	// provides marshal interface
	Marshal() marshalFlowTaggedMetric
	// provides unmarshal interface
	Unmarshal() unMarshalFlowTaggedMetric
	// validate validates FlowTaggedMetric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowTaggedMetric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Tags returns FlowTaggedMetricFlowMetricTagIterIter, set in FlowTaggedMetric
	Tags() FlowTaggedMetricFlowMetricTagIter
	// FramesTx returns uint64, set in FlowTaggedMetric.
	FramesTx() uint64
	// SetFramesTx assigns uint64 provided by user to FlowTaggedMetric
	SetFramesTx(value uint64) FlowTaggedMetric
	// HasFramesTx checks if FramesTx has been set in FlowTaggedMetric
	HasFramesTx() bool
	// FramesRx returns uint64, set in FlowTaggedMetric.
	FramesRx() uint64
	// SetFramesRx assigns uint64 provided by user to FlowTaggedMetric
	SetFramesRx(value uint64) FlowTaggedMetric
	// HasFramesRx checks if FramesRx has been set in FlowTaggedMetric
	HasFramesRx() bool
	// BytesTx returns uint64, set in FlowTaggedMetric.
	BytesTx() uint64
	// SetBytesTx assigns uint64 provided by user to FlowTaggedMetric
	SetBytesTx(value uint64) FlowTaggedMetric
	// HasBytesTx checks if BytesTx has been set in FlowTaggedMetric
	HasBytesTx() bool
	// BytesRx returns uint64, set in FlowTaggedMetric.
	BytesRx() uint64
	// SetBytesRx assigns uint64 provided by user to FlowTaggedMetric
	SetBytesRx(value uint64) FlowTaggedMetric
	// HasBytesRx checks if BytesRx has been set in FlowTaggedMetric
	HasBytesRx() bool
	// FramesTxRate returns float32, set in FlowTaggedMetric.
	FramesTxRate() float32
	// SetFramesTxRate assigns float32 provided by user to FlowTaggedMetric
	SetFramesTxRate(value float32) FlowTaggedMetric
	// HasFramesTxRate checks if FramesTxRate has been set in FlowTaggedMetric
	HasFramesTxRate() bool
	// FramesRxRate returns float32, set in FlowTaggedMetric.
	FramesRxRate() float32
	// SetFramesRxRate assigns float32 provided by user to FlowTaggedMetric
	SetFramesRxRate(value float32) FlowTaggedMetric
	// HasFramesRxRate checks if FramesRxRate has been set in FlowTaggedMetric
	HasFramesRxRate() bool
	// Loss returns float32, set in FlowTaggedMetric.
	Loss() float32
	// SetLoss assigns float32 provided by user to FlowTaggedMetric
	SetLoss(value float32) FlowTaggedMetric
	// HasLoss checks if Loss has been set in FlowTaggedMetric
	HasLoss() bool
	// Timestamps returns MetricTimestamp, set in FlowTaggedMetric.
	// MetricTimestamp is the container for timestamp metrics.
	// The container will be empty if the timestamp has not been configured for
	// the flow.
	Timestamps() MetricTimestamp
	// SetTimestamps assigns MetricTimestamp provided by user to FlowTaggedMetric.
	// MetricTimestamp is the container for timestamp metrics.
	// The container will be empty if the timestamp has not been configured for
	// the flow.
	SetTimestamps(value MetricTimestamp) FlowTaggedMetric
	// HasTimestamps checks if Timestamps has been set in FlowTaggedMetric
	HasTimestamps() bool
	// Latency returns MetricLatency, set in FlowTaggedMetric.
	// MetricLatency is the container for latency metrics.
	// The min/max/avg values are dependent on the type of latency measurement
	// mode that is configured.
	// The container will be empty if the latency has not been configured for
	// the flow.
	Latency() MetricLatency
	// SetLatency assigns MetricLatency provided by user to FlowTaggedMetric.
	// MetricLatency is the container for latency metrics.
	// The min/max/avg values are dependent on the type of latency measurement
	// mode that is configured.
	// The container will be empty if the latency has not been configured for
	// the flow.
	SetLatency(value MetricLatency) FlowTaggedMetric
	// HasLatency checks if Latency has been set in FlowTaggedMetric
	HasLatency() bool
	// TxL1RateBps returns float32, set in FlowTaggedMetric.
	TxL1RateBps() float32
	// SetTxL1RateBps assigns float32 provided by user to FlowTaggedMetric
	SetTxL1RateBps(value float32) FlowTaggedMetric
	// HasTxL1RateBps checks if TxL1RateBps has been set in FlowTaggedMetric
	HasTxL1RateBps() bool
	// RxL1RateBps returns float32, set in FlowTaggedMetric.
	RxL1RateBps() float32
	// SetRxL1RateBps assigns float32 provided by user to FlowTaggedMetric
	SetRxL1RateBps(value float32) FlowTaggedMetric
	// HasRxL1RateBps checks if RxL1RateBps has been set in FlowTaggedMetric
	HasRxL1RateBps() bool
	// TxRateBytes returns float32, set in FlowTaggedMetric.
	TxRateBytes() float32
	// SetTxRateBytes assigns float32 provided by user to FlowTaggedMetric
	SetTxRateBytes(value float32) FlowTaggedMetric
	// HasTxRateBytes checks if TxRateBytes has been set in FlowTaggedMetric
	HasTxRateBytes() bool
	// RxRateBytes returns float32, set in FlowTaggedMetric.
	RxRateBytes() float32
	// SetRxRateBytes assigns float32 provided by user to FlowTaggedMetric
	SetRxRateBytes(value float32) FlowTaggedMetric
	// HasRxRateBytes checks if RxRateBytes has been set in FlowTaggedMetric
	HasRxRateBytes() bool
	// TxRateBps returns float32, set in FlowTaggedMetric.
	TxRateBps() float32
	// SetTxRateBps assigns float32 provided by user to FlowTaggedMetric
	SetTxRateBps(value float32) FlowTaggedMetric
	// HasTxRateBps checks if TxRateBps has been set in FlowTaggedMetric
	HasTxRateBps() bool
	// RxRateBps returns float32, set in FlowTaggedMetric.
	RxRateBps() float32
	// SetRxRateBps assigns float32 provided by user to FlowTaggedMetric
	SetRxRateBps(value float32) FlowTaggedMetric
	// HasRxRateBps checks if RxRateBps has been set in FlowTaggedMetric
	HasRxRateBps() bool
	// TxRateKbps returns float32, set in FlowTaggedMetric.
	TxRateKbps() float32
	// SetTxRateKbps assigns float32 provided by user to FlowTaggedMetric
	SetTxRateKbps(value float32) FlowTaggedMetric
	// HasTxRateKbps checks if TxRateKbps has been set in FlowTaggedMetric
	HasTxRateKbps() bool
	// RxRateKbps returns float32, set in FlowTaggedMetric.
	RxRateKbps() float32
	// SetRxRateKbps assigns float32 provided by user to FlowTaggedMetric
	SetRxRateKbps(value float32) FlowTaggedMetric
	// HasRxRateKbps checks if RxRateKbps has been set in FlowTaggedMetric
	HasRxRateKbps() bool
	// TxRateMbps returns float32, set in FlowTaggedMetric.
	TxRateMbps() float32
	// SetTxRateMbps assigns float32 provided by user to FlowTaggedMetric
	SetTxRateMbps(value float32) FlowTaggedMetric
	// HasTxRateMbps checks if TxRateMbps has been set in FlowTaggedMetric
	HasTxRateMbps() bool
	// RxRateMbps returns float32, set in FlowTaggedMetric.
	RxRateMbps() float32
	// SetRxRateMbps assigns float32 provided by user to FlowTaggedMetric
	SetRxRateMbps(value float32) FlowTaggedMetric
	// HasRxRateMbps checks if RxRateMbps has been set in FlowTaggedMetric
	HasRxRateMbps() bool
	setNil()
}

// List of tag and value pairs
// Tags returns a []FlowMetricTag
func (obj *flowTaggedMetric) Tags() FlowTaggedMetricFlowMetricTagIter {
	if len(obj.obj.Tags) == 0 {
		obj.obj.Tags = []*otg.FlowMetricTag{}
	}
	if obj.tagsHolder == nil {
		obj.tagsHolder = newFlowTaggedMetricFlowMetricTagIter(&obj.obj.Tags).setMsg(obj)
	}
	return obj.tagsHolder
}

type flowTaggedMetricFlowMetricTagIter struct {
	obj                *flowTaggedMetric
	flowMetricTagSlice []FlowMetricTag
	fieldPtr           *[]*otg.FlowMetricTag
}

func newFlowTaggedMetricFlowMetricTagIter(ptr *[]*otg.FlowMetricTag) FlowTaggedMetricFlowMetricTagIter {
	return &flowTaggedMetricFlowMetricTagIter{fieldPtr: ptr}
}

type FlowTaggedMetricFlowMetricTagIter interface {
	setMsg(*flowTaggedMetric) FlowTaggedMetricFlowMetricTagIter
	Items() []FlowMetricTag
	Add() FlowMetricTag
	Append(items ...FlowMetricTag) FlowTaggedMetricFlowMetricTagIter
	Set(index int, newObj FlowMetricTag) FlowTaggedMetricFlowMetricTagIter
	Clear() FlowTaggedMetricFlowMetricTagIter
	clearHolderSlice() FlowTaggedMetricFlowMetricTagIter
	appendHolderSlice(item FlowMetricTag) FlowTaggedMetricFlowMetricTagIter
}

func (obj *flowTaggedMetricFlowMetricTagIter) setMsg(msg *flowTaggedMetric) FlowTaggedMetricFlowMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowTaggedMetricFlowMetricTagIter) Items() []FlowMetricTag {
	return obj.flowMetricTagSlice
}

func (obj *flowTaggedMetricFlowMetricTagIter) Add() FlowMetricTag {
	newObj := &otg.FlowMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.flowMetricTagSlice = append(obj.flowMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *flowTaggedMetricFlowMetricTagIter) Append(items ...FlowMetricTag) FlowTaggedMetricFlowMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowMetricTagSlice = append(obj.flowMetricTagSlice, item)
	}
	return obj
}

func (obj *flowTaggedMetricFlowMetricTagIter) Set(index int, newObj FlowMetricTag) FlowTaggedMetricFlowMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowMetricTagSlice[index] = newObj
	return obj
}
func (obj *flowTaggedMetricFlowMetricTagIter) Clear() FlowTaggedMetricFlowMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowMetricTag{}
		obj.flowMetricTagSlice = []FlowMetricTag{}
	}
	return obj
}
func (obj *flowTaggedMetricFlowMetricTagIter) clearHolderSlice() FlowTaggedMetricFlowMetricTagIter {
	if len(obj.flowMetricTagSlice) > 0 {
		obj.flowMetricTagSlice = []FlowMetricTag{}
	}
	return obj
}
func (obj *flowTaggedMetricFlowMetricTagIter) appendHolderSlice(item FlowMetricTag) FlowTaggedMetricFlowMetricTagIter {
	obj.flowMetricTagSlice = append(obj.flowMetricTagSlice, item)
	return obj
}

// The current total number of frames transmitted
// FramesTx returns a uint64
func (obj *flowTaggedMetric) FramesTx() uint64 {

	return *obj.obj.FramesTx

}

// The current total number of frames transmitted
// FramesTx returns a uint64
func (obj *flowTaggedMetric) HasFramesTx() bool {
	return obj.obj.FramesTx != nil
}

// The current total number of frames transmitted
// SetFramesTx sets the uint64 value in the FlowTaggedMetric object
func (obj *flowTaggedMetric) SetFramesTx(value uint64) FlowTaggedMetric {

	obj.obj.FramesTx = &value
	return obj
}

// The current total number of valid frames received
// FramesRx returns a uint64
func (obj *flowTaggedMetric) FramesRx() uint64 {

	return *obj.obj.FramesRx

}

// The current total number of valid frames received
// FramesRx returns a uint64
func (obj *flowTaggedMetric) HasFramesRx() bool {
	return obj.obj.FramesRx != nil
}

// The current total number of valid frames received
// SetFramesRx sets the uint64 value in the FlowTaggedMetric object
func (obj *flowTaggedMetric) SetFramesRx(value uint64) FlowTaggedMetric {

	obj.obj.FramesRx = &value
	return obj
}

// The current total number of bytes transmitted
// BytesTx returns a uint64
func (obj *flowTaggedMetric) BytesTx() uint64 {

	return *obj.obj.BytesTx

}

// The current total number of bytes transmitted
// BytesTx returns a uint64
func (obj *flowTaggedMetric) HasBytesTx() bool {
	return obj.obj.BytesTx != nil
}

// The current total number of bytes transmitted
// SetBytesTx sets the uint64 value in the FlowTaggedMetric object
func (obj *flowTaggedMetric) SetBytesTx(value uint64) FlowTaggedMetric {

	obj.obj.BytesTx = &value
	return obj
}

// The current total number of bytes received
// BytesRx returns a uint64
func (obj *flowTaggedMetric) BytesRx() uint64 {

	return *obj.obj.BytesRx

}

// The current total number of bytes received
// BytesRx returns a uint64
func (obj *flowTaggedMetric) HasBytesRx() bool {
	return obj.obj.BytesRx != nil
}

// The current total number of bytes received
// SetBytesRx sets the uint64 value in the FlowTaggedMetric object
func (obj *flowTaggedMetric) SetBytesRx(value uint64) FlowTaggedMetric {

	obj.obj.BytesRx = &value
	return obj
}

// The current rate of frames transmitted
// FramesTxRate returns a float32
func (obj *flowTaggedMetric) FramesTxRate() float32 {

	return *obj.obj.FramesTxRate

}

// The current rate of frames transmitted
// FramesTxRate returns a float32
func (obj *flowTaggedMetric) HasFramesTxRate() bool {
	return obj.obj.FramesTxRate != nil
}

// The current rate of frames transmitted
// SetFramesTxRate sets the float32 value in the FlowTaggedMetric object
func (obj *flowTaggedMetric) SetFramesTxRate(value float32) FlowTaggedMetric {

	obj.obj.FramesTxRate = &value
	return obj
}

// The current rate of valid frames received
// FramesRxRate returns a float32
func (obj *flowTaggedMetric) FramesRxRate() float32 {

	return *obj.obj.FramesRxRate

}

// The current rate of valid frames received
// FramesRxRate returns a float32
func (obj *flowTaggedMetric) HasFramesRxRate() bool {
	return obj.obj.FramesRxRate != nil
}

// The current rate of valid frames received
// SetFramesRxRate sets the float32 value in the FlowTaggedMetric object
func (obj *flowTaggedMetric) SetFramesRxRate(value float32) FlowTaggedMetric {

	obj.obj.FramesRxRate = &value
	return obj
}

// The percentage of lost frames
// Loss returns a float32
func (obj *flowTaggedMetric) Loss() float32 {

	return *obj.obj.Loss

}

// The percentage of lost frames
// Loss returns a float32
func (obj *flowTaggedMetric) HasLoss() bool {
	return obj.obj.Loss != nil
}

// The percentage of lost frames
// SetLoss sets the float32 value in the FlowTaggedMetric object
func (obj *flowTaggedMetric) SetLoss(value float32) FlowTaggedMetric {

	obj.obj.Loss = &value
	return obj
}

// description is TBD
// Timestamps returns a MetricTimestamp
func (obj *flowTaggedMetric) Timestamps() MetricTimestamp {
	if obj.obj.Timestamps == nil {
		obj.obj.Timestamps = NewMetricTimestamp().msg()
	}
	if obj.timestampsHolder == nil {
		obj.timestampsHolder = &metricTimestamp{obj: obj.obj.Timestamps}
	}
	return obj.timestampsHolder
}

// description is TBD
// Timestamps returns a MetricTimestamp
func (obj *flowTaggedMetric) HasTimestamps() bool {
	return obj.obj.Timestamps != nil
}

// description is TBD
// SetTimestamps sets the MetricTimestamp value in the FlowTaggedMetric object
func (obj *flowTaggedMetric) SetTimestamps(value MetricTimestamp) FlowTaggedMetric {

	obj.timestampsHolder = nil
	obj.obj.Timestamps = value.msg()

	return obj
}

// description is TBD
// Latency returns a MetricLatency
func (obj *flowTaggedMetric) Latency() MetricLatency {
	if obj.obj.Latency == nil {
		obj.obj.Latency = NewMetricLatency().msg()
	}
	if obj.latencyHolder == nil {
		obj.latencyHolder = &metricLatency{obj: obj.obj.Latency}
	}
	return obj.latencyHolder
}

// description is TBD
// Latency returns a MetricLatency
func (obj *flowTaggedMetric) HasLatency() bool {
	return obj.obj.Latency != nil
}

// description is TBD
// SetLatency sets the MetricLatency value in the FlowTaggedMetric object
func (obj *flowTaggedMetric) SetLatency(value MetricLatency) FlowTaggedMetric {

	obj.latencyHolder = nil
	obj.obj.Latency = value.msg()

	return obj
}

// The Layer 1 transmission rate in bits per second.
// TxL1RateBps returns a float32
func (obj *flowTaggedMetric) TxL1RateBps() float32 {

	return *obj.obj.TxL1RateBps

}

// The Layer 1 transmission rate in bits per second.
// TxL1RateBps returns a float32
func (obj *flowTaggedMetric) HasTxL1RateBps() bool {
	return obj.obj.TxL1RateBps != nil
}

// The Layer 1 transmission rate in bits per second.
// SetTxL1RateBps sets the float32 value in the FlowTaggedMetric object
func (obj *flowTaggedMetric) SetTxL1RateBps(value float32) FlowTaggedMetric {

	obj.obj.TxL1RateBps = &value
	return obj
}

// The Layer 1 receive rate in bits per second.
// RxL1RateBps returns a float32
func (obj *flowTaggedMetric) RxL1RateBps() float32 {

	return *obj.obj.RxL1RateBps

}

// The Layer 1 receive rate in bits per second.
// RxL1RateBps returns a float32
func (obj *flowTaggedMetric) HasRxL1RateBps() bool {
	return obj.obj.RxL1RateBps != nil
}

// The Layer 1 receive rate in bits per second.
// SetRxL1RateBps sets the float32 value in the FlowTaggedMetric object
func (obj *flowTaggedMetric) SetRxL1RateBps(value float32) FlowTaggedMetric {

	obj.obj.RxL1RateBps = &value
	return obj
}

// The transmission rate in bytes per second.
// TxRateBytes returns a float32
func (obj *flowTaggedMetric) TxRateBytes() float32 {

	return *obj.obj.TxRateBytes

}

// The transmission rate in bytes per second.
// TxRateBytes returns a float32
func (obj *flowTaggedMetric) HasTxRateBytes() bool {
	return obj.obj.TxRateBytes != nil
}

// The transmission rate in bytes per second.
// SetTxRateBytes sets the float32 value in the FlowTaggedMetric object
func (obj *flowTaggedMetric) SetTxRateBytes(value float32) FlowTaggedMetric {

	obj.obj.TxRateBytes = &value
	return obj
}

// The receive rate in bytes per second.
// RxRateBytes returns a float32
func (obj *flowTaggedMetric) RxRateBytes() float32 {

	return *obj.obj.RxRateBytes

}

// The receive rate in bytes per second.
// RxRateBytes returns a float32
func (obj *flowTaggedMetric) HasRxRateBytes() bool {
	return obj.obj.RxRateBytes != nil
}

// The receive rate in bytes per second.
// SetRxRateBytes sets the float32 value in the FlowTaggedMetric object
func (obj *flowTaggedMetric) SetRxRateBytes(value float32) FlowTaggedMetric {

	obj.obj.RxRateBytes = &value
	return obj
}

// The transmission rate in bits per second.
// TxRateBps returns a float32
func (obj *flowTaggedMetric) TxRateBps() float32 {

	return *obj.obj.TxRateBps

}

// The transmission rate in bits per second.
// TxRateBps returns a float32
func (obj *flowTaggedMetric) HasTxRateBps() bool {
	return obj.obj.TxRateBps != nil
}

// The transmission rate in bits per second.
// SetTxRateBps sets the float32 value in the FlowTaggedMetric object
func (obj *flowTaggedMetric) SetTxRateBps(value float32) FlowTaggedMetric {

	obj.obj.TxRateBps = &value
	return obj
}

// The receive rate in bits per second.
// RxRateBps returns a float32
func (obj *flowTaggedMetric) RxRateBps() float32 {

	return *obj.obj.RxRateBps

}

// The receive rate in bits per second.
// RxRateBps returns a float32
func (obj *flowTaggedMetric) HasRxRateBps() bool {
	return obj.obj.RxRateBps != nil
}

// The receive rate in bits per second.
// SetRxRateBps sets the float32 value in the FlowTaggedMetric object
func (obj *flowTaggedMetric) SetRxRateBps(value float32) FlowTaggedMetric {

	obj.obj.RxRateBps = &value
	return obj
}

// The transmission rate in Kilobits per second.
// TxRateKbps returns a float32
func (obj *flowTaggedMetric) TxRateKbps() float32 {

	return *obj.obj.TxRateKbps

}

// The transmission rate in Kilobits per second.
// TxRateKbps returns a float32
func (obj *flowTaggedMetric) HasTxRateKbps() bool {
	return obj.obj.TxRateKbps != nil
}

// The transmission rate in Kilobits per second.
// SetTxRateKbps sets the float32 value in the FlowTaggedMetric object
func (obj *flowTaggedMetric) SetTxRateKbps(value float32) FlowTaggedMetric {

	obj.obj.TxRateKbps = &value
	return obj
}

// The receive rate in Kilobits per second.
// RxRateKbps returns a float32
func (obj *flowTaggedMetric) RxRateKbps() float32 {

	return *obj.obj.RxRateKbps

}

// The receive rate in Kilobits per second.
// RxRateKbps returns a float32
func (obj *flowTaggedMetric) HasRxRateKbps() bool {
	return obj.obj.RxRateKbps != nil
}

// The receive rate in Kilobits per second.
// SetRxRateKbps sets the float32 value in the FlowTaggedMetric object
func (obj *flowTaggedMetric) SetRxRateKbps(value float32) FlowTaggedMetric {

	obj.obj.RxRateKbps = &value
	return obj
}

// The transmission rate in Megabits per second.
// TxRateMbps returns a float32
func (obj *flowTaggedMetric) TxRateMbps() float32 {

	return *obj.obj.TxRateMbps

}

// The transmission rate in Megabits per second.
// TxRateMbps returns a float32
func (obj *flowTaggedMetric) HasTxRateMbps() bool {
	return obj.obj.TxRateMbps != nil
}

// The transmission rate in Megabits per second.
// SetTxRateMbps sets the float32 value in the FlowTaggedMetric object
func (obj *flowTaggedMetric) SetTxRateMbps(value float32) FlowTaggedMetric {

	obj.obj.TxRateMbps = &value
	return obj
}

// The receive rate in Megabits per second.
// RxRateMbps returns a float32
func (obj *flowTaggedMetric) RxRateMbps() float32 {

	return *obj.obj.RxRateMbps

}

// The receive rate in Megabits per second.
// RxRateMbps returns a float32
func (obj *flowTaggedMetric) HasRxRateMbps() bool {
	return obj.obj.RxRateMbps != nil
}

// The receive rate in Megabits per second.
// SetRxRateMbps sets the float32 value in the FlowTaggedMetric object
func (obj *flowTaggedMetric) SetRxRateMbps(value float32) FlowTaggedMetric {

	obj.obj.RxRateMbps = &value
	return obj
}

func (obj *flowTaggedMetric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Tags) != 0 {

		if set_default {
			obj.Tags().clearHolderSlice()
			for _, item := range obj.obj.Tags {
				obj.Tags().appendHolderSlice(&flowMetricTag{obj: item})
			}
		}
		for _, item := range obj.Tags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.Timestamps != nil {

		obj.Timestamps().validateObj(vObj, set_default)
	}

	if obj.obj.Latency != nil {

		obj.Latency().validateObj(vObj, set_default)
	}

}

func (obj *flowTaggedMetric) setDefault() {

}
