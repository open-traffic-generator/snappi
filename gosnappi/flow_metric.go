package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowMetric *****
type flowMetric struct {
	validation
	obj                 *otg.FlowMetric
	marshaller          marshalFlowMetric
	unMarshaller        unMarshalFlowMetric
	timestampsHolder    MetricTimestamp
	latencyHolder       MetricLatency
	taggedMetricsHolder FlowMetricFlowTaggedMetricIter
}

func NewFlowMetric() FlowMetric {
	obj := flowMetric{obj: &otg.FlowMetric{}}
	obj.setDefault()
	return &obj
}

func (obj *flowMetric) msg() *otg.FlowMetric {
	return obj.obj
}

func (obj *flowMetric) setMsg(msg *otg.FlowMetric) FlowMetric {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowMetric struct {
	obj *flowMetric
}

type marshalFlowMetric interface {
	// ToProto marshals FlowMetric to protobuf object *otg.FlowMetric
	ToProto() (*otg.FlowMetric, error)
	// ToPbText marshals FlowMetric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowMetric to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowMetric to JSON text
	ToJson() (string, error)
}

type unMarshalflowMetric struct {
	obj *flowMetric
}

type unMarshalFlowMetric interface {
	// FromProto unmarshals FlowMetric from protobuf object *otg.FlowMetric
	FromProto(msg *otg.FlowMetric) (FlowMetric, error)
	// FromPbText unmarshals FlowMetric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowMetric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowMetric from JSON text
	FromJson(value string) error
}

func (obj *flowMetric) Marshal() marshalFlowMetric {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowMetric{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowMetric) Unmarshal() unMarshalFlowMetric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowMetric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowMetric) ToProto() (*otg.FlowMetric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowMetric) FromProto(msg *otg.FlowMetric) (FlowMetric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowMetric) ToPbText() (string, error) {
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

func (m *unMarshalflowMetric) FromPbText(value string) error {
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

func (m *marshalflowMetric) ToYaml() (string, error) {
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

func (m *unMarshalflowMetric) FromYaml(value string) error {
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

func (m *marshalflowMetric) ToJson() (string, error) {
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

func (m *unMarshalflowMetric) FromJson(value string) error {
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

func (obj *flowMetric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowMetric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowMetric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowMetric) Clone() (FlowMetric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowMetric()
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

func (obj *flowMetric) setNil() {
	obj.timestampsHolder = nil
	obj.latencyHolder = nil
	obj.taggedMetricsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowMetric is a container for flow metrics.
// The container is keyed by the name, port_tx and port_rx.
type FlowMetric interface {
	Validation
	// msg marshals FlowMetric to protobuf object *otg.FlowMetric
	// and doesn't set defaults
	msg() *otg.FlowMetric
	// setMsg unmarshals FlowMetric from protobuf object *otg.FlowMetric
	// and doesn't set defaults
	setMsg(*otg.FlowMetric) FlowMetric
	// provides marshal interface
	Marshal() marshalFlowMetric
	// provides unmarshal interface
	Unmarshal() unMarshalFlowMetric
	// validate validates FlowMetric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowMetric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in FlowMetric.
	Name() string
	// SetName assigns string provided by user to FlowMetric
	SetName(value string) FlowMetric
	// HasName checks if Name has been set in FlowMetric
	HasName() bool
	// PortTx returns string, set in FlowMetric.
	PortTx() string
	// SetPortTx assigns string provided by user to FlowMetric
	SetPortTx(value string) FlowMetric
	// HasPortTx checks if PortTx has been set in FlowMetric
	HasPortTx() bool
	// PortRx returns string, set in FlowMetric.
	PortRx() string
	// SetPortRx assigns string provided by user to FlowMetric
	SetPortRx(value string) FlowMetric
	// HasPortRx checks if PortRx has been set in FlowMetric
	HasPortRx() bool
	// Transmit returns FlowMetricTransmitEnum, set in FlowMetric
	Transmit() FlowMetricTransmitEnum
	// SetTransmit assigns FlowMetricTransmitEnum provided by user to FlowMetric
	SetTransmit(value FlowMetricTransmitEnum) FlowMetric
	// HasTransmit checks if Transmit has been set in FlowMetric
	HasTransmit() bool
	// FramesTx returns uint64, set in FlowMetric.
	FramesTx() uint64
	// SetFramesTx assigns uint64 provided by user to FlowMetric
	SetFramesTx(value uint64) FlowMetric
	// HasFramesTx checks if FramesTx has been set in FlowMetric
	HasFramesTx() bool
	// FramesRx returns uint64, set in FlowMetric.
	FramesRx() uint64
	// SetFramesRx assigns uint64 provided by user to FlowMetric
	SetFramesRx(value uint64) FlowMetric
	// HasFramesRx checks if FramesRx has been set in FlowMetric
	HasFramesRx() bool
	// BytesTx returns uint64, set in FlowMetric.
	BytesTx() uint64
	// SetBytesTx assigns uint64 provided by user to FlowMetric
	SetBytesTx(value uint64) FlowMetric
	// HasBytesTx checks if BytesTx has been set in FlowMetric
	HasBytesTx() bool
	// BytesRx returns uint64, set in FlowMetric.
	BytesRx() uint64
	// SetBytesRx assigns uint64 provided by user to FlowMetric
	SetBytesRx(value uint64) FlowMetric
	// HasBytesRx checks if BytesRx has been set in FlowMetric
	HasBytesRx() bool
	// FramesTxRate returns float32, set in FlowMetric.
	FramesTxRate() float32
	// SetFramesTxRate assigns float32 provided by user to FlowMetric
	SetFramesTxRate(value float32) FlowMetric
	// HasFramesTxRate checks if FramesTxRate has been set in FlowMetric
	HasFramesTxRate() bool
	// FramesRxRate returns float32, set in FlowMetric.
	FramesRxRate() float32
	// SetFramesRxRate assigns float32 provided by user to FlowMetric
	SetFramesRxRate(value float32) FlowMetric
	// HasFramesRxRate checks if FramesRxRate has been set in FlowMetric
	HasFramesRxRate() bool
	// Loss returns float32, set in FlowMetric.
	Loss() float32
	// SetLoss assigns float32 provided by user to FlowMetric
	SetLoss(value float32) FlowMetric
	// HasLoss checks if Loss has been set in FlowMetric
	HasLoss() bool
	// Timestamps returns MetricTimestamp, set in FlowMetric.
	// MetricTimestamp is the container for timestamp metrics.
	// The container will be empty if the timestamp has not been configured for
	// the flow.
	Timestamps() MetricTimestamp
	// SetTimestamps assigns MetricTimestamp provided by user to FlowMetric.
	// MetricTimestamp is the container for timestamp metrics.
	// The container will be empty if the timestamp has not been configured for
	// the flow.
	SetTimestamps(value MetricTimestamp) FlowMetric
	// HasTimestamps checks if Timestamps has been set in FlowMetric
	HasTimestamps() bool
	// Latency returns MetricLatency, set in FlowMetric.
	// MetricLatency is the container for latency metrics.
	// The min/max/avg values are dependent on the type of latency measurement
	// mode that is configured.
	// The container will be empty if the latency has not been configured for
	// the flow.
	Latency() MetricLatency
	// SetLatency assigns MetricLatency provided by user to FlowMetric.
	// MetricLatency is the container for latency metrics.
	// The min/max/avg values are dependent on the type of latency measurement
	// mode that is configured.
	// The container will be empty if the latency has not been configured for
	// the flow.
	SetLatency(value MetricLatency) FlowMetric
	// HasLatency checks if Latency has been set in FlowMetric
	HasLatency() bool
	// TaggedMetrics returns FlowMetricFlowTaggedMetricIterIter, set in FlowMetric
	TaggedMetrics() FlowMetricFlowTaggedMetricIter
	setNil()
}

// The name of the flow
// Name returns a string
func (obj *flowMetric) Name() string {

	return *obj.obj.Name

}

// The name of the flow
// Name returns a string
func (obj *flowMetric) HasName() bool {
	return obj.obj.Name != nil
}

// The name of the flow
// SetName sets the string value in the FlowMetric object
func (obj *flowMetric) SetName(value string) FlowMetric {

	obj.obj.Name = &value
	return obj
}

// The name of the transmit port
// PortTx returns a string
func (obj *flowMetric) PortTx() string {

	return *obj.obj.PortTx

}

// The name of the transmit port
// PortTx returns a string
func (obj *flowMetric) HasPortTx() bool {
	return obj.obj.PortTx != nil
}

// The name of the transmit port
// SetPortTx sets the string value in the FlowMetric object
func (obj *flowMetric) SetPortTx(value string) FlowMetric {

	obj.obj.PortTx = &value
	return obj
}

// The name of the receive port
// PortRx returns a string
func (obj *flowMetric) PortRx() string {

	return *obj.obj.PortRx

}

// The name of the receive port
// PortRx returns a string
func (obj *flowMetric) HasPortRx() bool {
	return obj.obj.PortRx != nil
}

// The name of the receive port
// SetPortRx sets the string value in the FlowMetric object
func (obj *flowMetric) SetPortRx(value string) FlowMetric {

	obj.obj.PortRx = &value
	return obj
}

type FlowMetricTransmitEnum string

// Enum of Transmit on FlowMetric
var FlowMetricTransmit = struct {
	STARTED FlowMetricTransmitEnum
	STOPPED FlowMetricTransmitEnum
	PAUSED  FlowMetricTransmitEnum
}{
	STARTED: FlowMetricTransmitEnum("started"),
	STOPPED: FlowMetricTransmitEnum("stopped"),
	PAUSED:  FlowMetricTransmitEnum("paused"),
}

func (obj *flowMetric) Transmit() FlowMetricTransmitEnum {
	return FlowMetricTransmitEnum(obj.obj.Transmit.Enum().String())
}

// The transmit state of the flow.
// Transmit returns a string
func (obj *flowMetric) HasTransmit() bool {
	return obj.obj.Transmit != nil
}

func (obj *flowMetric) SetTransmit(value FlowMetricTransmitEnum) FlowMetric {
	intValue, ok := otg.FlowMetric_Transmit_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowMetricTransmitEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowMetric_Transmit_Enum(intValue)
	obj.obj.Transmit = &enumValue

	return obj
}

// The current total number of frames transmitted
// FramesTx returns a uint64
func (obj *flowMetric) FramesTx() uint64 {

	return *obj.obj.FramesTx

}

// The current total number of frames transmitted
// FramesTx returns a uint64
func (obj *flowMetric) HasFramesTx() bool {
	return obj.obj.FramesTx != nil
}

// The current total number of frames transmitted
// SetFramesTx sets the uint64 value in the FlowMetric object
func (obj *flowMetric) SetFramesTx(value uint64) FlowMetric {

	obj.obj.FramesTx = &value
	return obj
}

// The current total number of valid frames received
// FramesRx returns a uint64
func (obj *flowMetric) FramesRx() uint64 {

	return *obj.obj.FramesRx

}

// The current total number of valid frames received
// FramesRx returns a uint64
func (obj *flowMetric) HasFramesRx() bool {
	return obj.obj.FramesRx != nil
}

// The current total number of valid frames received
// SetFramesRx sets the uint64 value in the FlowMetric object
func (obj *flowMetric) SetFramesRx(value uint64) FlowMetric {

	obj.obj.FramesRx = &value
	return obj
}

// The current total number of bytes transmitted
// BytesTx returns a uint64
func (obj *flowMetric) BytesTx() uint64 {

	return *obj.obj.BytesTx

}

// The current total number of bytes transmitted
// BytesTx returns a uint64
func (obj *flowMetric) HasBytesTx() bool {
	return obj.obj.BytesTx != nil
}

// The current total number of bytes transmitted
// SetBytesTx sets the uint64 value in the FlowMetric object
func (obj *flowMetric) SetBytesTx(value uint64) FlowMetric {

	obj.obj.BytesTx = &value
	return obj
}

// The current total number of bytes received
// BytesRx returns a uint64
func (obj *flowMetric) BytesRx() uint64 {

	return *obj.obj.BytesRx

}

// The current total number of bytes received
// BytesRx returns a uint64
func (obj *flowMetric) HasBytesRx() bool {
	return obj.obj.BytesRx != nil
}

// The current total number of bytes received
// SetBytesRx sets the uint64 value in the FlowMetric object
func (obj *flowMetric) SetBytesRx(value uint64) FlowMetric {

	obj.obj.BytesRx = &value
	return obj
}

// The current rate of frames transmitted
// FramesTxRate returns a float32
func (obj *flowMetric) FramesTxRate() float32 {

	return *obj.obj.FramesTxRate

}

// The current rate of frames transmitted
// FramesTxRate returns a float32
func (obj *flowMetric) HasFramesTxRate() bool {
	return obj.obj.FramesTxRate != nil
}

// The current rate of frames transmitted
// SetFramesTxRate sets the float32 value in the FlowMetric object
func (obj *flowMetric) SetFramesTxRate(value float32) FlowMetric {

	obj.obj.FramesTxRate = &value
	return obj
}

// The current rate of valid frames received
// FramesRxRate returns a float32
func (obj *flowMetric) FramesRxRate() float32 {

	return *obj.obj.FramesRxRate

}

// The current rate of valid frames received
// FramesRxRate returns a float32
func (obj *flowMetric) HasFramesRxRate() bool {
	return obj.obj.FramesRxRate != nil
}

// The current rate of valid frames received
// SetFramesRxRate sets the float32 value in the FlowMetric object
func (obj *flowMetric) SetFramesRxRate(value float32) FlowMetric {

	obj.obj.FramesRxRate = &value
	return obj
}

// The percentage of lost frames
// Loss returns a float32
func (obj *flowMetric) Loss() float32 {

	return *obj.obj.Loss

}

// The percentage of lost frames
// Loss returns a float32
func (obj *flowMetric) HasLoss() bool {
	return obj.obj.Loss != nil
}

// The percentage of lost frames
// SetLoss sets the float32 value in the FlowMetric object
func (obj *flowMetric) SetLoss(value float32) FlowMetric {

	obj.obj.Loss = &value
	return obj
}

// description is TBD
// Timestamps returns a MetricTimestamp
func (obj *flowMetric) Timestamps() MetricTimestamp {
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
func (obj *flowMetric) HasTimestamps() bool {
	return obj.obj.Timestamps != nil
}

// description is TBD
// SetTimestamps sets the MetricTimestamp value in the FlowMetric object
func (obj *flowMetric) SetTimestamps(value MetricTimestamp) FlowMetric {

	obj.timestampsHolder = nil
	obj.obj.Timestamps = value.msg()

	return obj
}

// description is TBD
// Latency returns a MetricLatency
func (obj *flowMetric) Latency() MetricLatency {
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
func (obj *flowMetric) HasLatency() bool {
	return obj.obj.Latency != nil
}

// description is TBD
// SetLatency sets the MetricLatency value in the FlowMetric object
func (obj *flowMetric) SetLatency(value MetricLatency) FlowMetric {

	obj.latencyHolder = nil
	obj.obj.Latency = value.msg()

	return obj
}

// List of metrics corresponding to a set of values applicable
// for configured metric tags in ingress or egress packet header fields of corresponding flow.
// The container is keyed by list of tag-value pairs.
// TaggedMetrics returns a []FlowTaggedMetric
func (obj *flowMetric) TaggedMetrics() FlowMetricFlowTaggedMetricIter {
	if len(obj.obj.TaggedMetrics) == 0 {
		obj.obj.TaggedMetrics = []*otg.FlowTaggedMetric{}
	}
	if obj.taggedMetricsHolder == nil {
		obj.taggedMetricsHolder = newFlowMetricFlowTaggedMetricIter(&obj.obj.TaggedMetrics).setMsg(obj)
	}
	return obj.taggedMetricsHolder
}

type flowMetricFlowTaggedMetricIter struct {
	obj                   *flowMetric
	flowTaggedMetricSlice []FlowTaggedMetric
	fieldPtr              *[]*otg.FlowTaggedMetric
}

func newFlowMetricFlowTaggedMetricIter(ptr *[]*otg.FlowTaggedMetric) FlowMetricFlowTaggedMetricIter {
	return &flowMetricFlowTaggedMetricIter{fieldPtr: ptr}
}

type FlowMetricFlowTaggedMetricIter interface {
	setMsg(*flowMetric) FlowMetricFlowTaggedMetricIter
	Items() []FlowTaggedMetric
	Add() FlowTaggedMetric
	Append(items ...FlowTaggedMetric) FlowMetricFlowTaggedMetricIter
	Set(index int, newObj FlowTaggedMetric) FlowMetricFlowTaggedMetricIter
	Clear() FlowMetricFlowTaggedMetricIter
	clearHolderSlice() FlowMetricFlowTaggedMetricIter
	appendHolderSlice(item FlowTaggedMetric) FlowMetricFlowTaggedMetricIter
}

func (obj *flowMetricFlowTaggedMetricIter) setMsg(msg *flowMetric) FlowMetricFlowTaggedMetricIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowTaggedMetric{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowMetricFlowTaggedMetricIter) Items() []FlowTaggedMetric {
	return obj.flowTaggedMetricSlice
}

func (obj *flowMetricFlowTaggedMetricIter) Add() FlowTaggedMetric {
	newObj := &otg.FlowTaggedMetric{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowTaggedMetric{obj: newObj}
	newLibObj.setDefault()
	obj.flowTaggedMetricSlice = append(obj.flowTaggedMetricSlice, newLibObj)
	return newLibObj
}

func (obj *flowMetricFlowTaggedMetricIter) Append(items ...FlowTaggedMetric) FlowMetricFlowTaggedMetricIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowTaggedMetricSlice = append(obj.flowTaggedMetricSlice, item)
	}
	return obj
}

func (obj *flowMetricFlowTaggedMetricIter) Set(index int, newObj FlowTaggedMetric) FlowMetricFlowTaggedMetricIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowTaggedMetricSlice[index] = newObj
	return obj
}
func (obj *flowMetricFlowTaggedMetricIter) Clear() FlowMetricFlowTaggedMetricIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowTaggedMetric{}
		obj.flowTaggedMetricSlice = []FlowTaggedMetric{}
	}
	return obj
}
func (obj *flowMetricFlowTaggedMetricIter) clearHolderSlice() FlowMetricFlowTaggedMetricIter {
	if len(obj.flowTaggedMetricSlice) > 0 {
		obj.flowTaggedMetricSlice = []FlowTaggedMetric{}
	}
	return obj
}
func (obj *flowMetricFlowTaggedMetricIter) appendHolderSlice(item FlowTaggedMetric) FlowMetricFlowTaggedMetricIter {
	obj.flowTaggedMetricSlice = append(obj.flowTaggedMetricSlice, item)
	return obj
}

func (obj *flowMetric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Timestamps != nil {

		obj.Timestamps().validateObj(vObj, set_default)
	}

	if obj.obj.Latency != nil {

		obj.Latency().validateObj(vObj, set_default)
	}

	if len(obj.obj.TaggedMetrics) != 0 {

		if set_default {
			obj.TaggedMetrics().clearHolderSlice()
			for _, item := range obj.obj.TaggedMetrics {
				obj.TaggedMetrics().appendHolderSlice(&flowTaggedMetric{obj: item})
			}
		}
		for _, item := range obj.TaggedMetrics().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *flowMetric) setDefault() {

}
