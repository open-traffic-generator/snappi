package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowMetricsRequest *****
type flowMetricsRequest struct {
	validation
	obj                 *otg.FlowMetricsRequest
	marshaller          marshalFlowMetricsRequest
	unMarshaller        unMarshalFlowMetricsRequest
	taggedMetricsHolder FlowTaggedMetricsFilter
}

func NewFlowMetricsRequest() FlowMetricsRequest {
	obj := flowMetricsRequest{obj: &otg.FlowMetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *flowMetricsRequest) msg() *otg.FlowMetricsRequest {
	return obj.obj
}

func (obj *flowMetricsRequest) setMsg(msg *otg.FlowMetricsRequest) FlowMetricsRequest {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowMetricsRequest struct {
	obj *flowMetricsRequest
}

type marshalFlowMetricsRequest interface {
	// ToProto marshals FlowMetricsRequest to protobuf object *otg.FlowMetricsRequest
	ToProto() (*otg.FlowMetricsRequest, error)
	// ToPbText marshals FlowMetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowMetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowMetricsRequest to JSON text
	ToJson() (string, error)
}

type unMarshalflowMetricsRequest struct {
	obj *flowMetricsRequest
}

type unMarshalFlowMetricsRequest interface {
	// FromProto unmarshals FlowMetricsRequest from protobuf object *otg.FlowMetricsRequest
	FromProto(msg *otg.FlowMetricsRequest) (FlowMetricsRequest, error)
	// FromPbText unmarshals FlowMetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowMetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowMetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *flowMetricsRequest) Marshal() marshalFlowMetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowMetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowMetricsRequest) Unmarshal() unMarshalFlowMetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowMetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowMetricsRequest) ToProto() (*otg.FlowMetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowMetricsRequest) FromProto(msg *otg.FlowMetricsRequest) (FlowMetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowMetricsRequest) ToPbText() (string, error) {
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

func (m *unMarshalflowMetricsRequest) FromPbText(value string) error {
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

func (m *marshalflowMetricsRequest) ToYaml() (string, error) {
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

func (m *unMarshalflowMetricsRequest) FromYaml(value string) error {
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

func (m *marshalflowMetricsRequest) ToJson() (string, error) {
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

func (m *unMarshalflowMetricsRequest) FromJson(value string) error {
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

func (obj *flowMetricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowMetricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowMetricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowMetricsRequest) Clone() (FlowMetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowMetricsRequest()
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

func (obj *flowMetricsRequest) setNil() {
	obj.taggedMetricsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowMetricsRequest is the container for a flow metric request.
type FlowMetricsRequest interface {
	Validation
	// msg marshals FlowMetricsRequest to protobuf object *otg.FlowMetricsRequest
	// and doesn't set defaults
	msg() *otg.FlowMetricsRequest
	// setMsg unmarshals FlowMetricsRequest from protobuf object *otg.FlowMetricsRequest
	// and doesn't set defaults
	setMsg(*otg.FlowMetricsRequest) FlowMetricsRequest
	// provides marshal interface
	Marshal() marshalFlowMetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalFlowMetricsRequest
	// validate validates FlowMetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowMetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// FlowNames returns []string, set in FlowMetricsRequest.
	FlowNames() []string
	// SetFlowNames assigns []string provided by user to FlowMetricsRequest
	SetFlowNames(value []string) FlowMetricsRequest
	// MetricNames returns []FlowMetricsRequestMetricNamesEnum, set in FlowMetricsRequest
	MetricNames() []FlowMetricsRequestMetricNamesEnum
	// SetMetricNames assigns []FlowMetricsRequestMetricNamesEnum provided by user to FlowMetricsRequest
	SetMetricNames(value []FlowMetricsRequestMetricNamesEnum) FlowMetricsRequest
	// TaggedMetrics returns FlowTaggedMetricsFilter, set in FlowMetricsRequest.
	// FlowTaggedMetricsFilter is filter for tagged metrics
	TaggedMetrics() FlowTaggedMetricsFilter
	// SetTaggedMetrics assigns FlowTaggedMetricsFilter provided by user to FlowMetricsRequest.
	// FlowTaggedMetricsFilter is filter for tagged metrics
	SetTaggedMetrics(value FlowTaggedMetricsFilter) FlowMetricsRequest
	// HasTaggedMetrics checks if TaggedMetrics has been set in FlowMetricsRequest
	HasTaggedMetrics() bool
	setNil()
}

// Flow metrics will be retrieved for these flow names.
// If no flow names are specified then all flows will be returned.
//
// x-constraint:
// - /components/schemas/Flow/properties/name
//
// FlowNames returns a []string
func (obj *flowMetricsRequest) FlowNames() []string {
	if obj.obj.FlowNames == nil {
		obj.obj.FlowNames = make([]string, 0)
	}
	return obj.obj.FlowNames
}

// Flow metrics will be retrieved for these flow names.
// If no flow names are specified then all flows will be returned.
//
// x-constraint:
// - /components/schemas/Flow/properties/name
//
// SetFlowNames sets the []string value in the FlowMetricsRequest object
func (obj *flowMetricsRequest) SetFlowNames(value []string) FlowMetricsRequest {

	if obj.obj.FlowNames == nil {
		obj.obj.FlowNames = make([]string, 0)
	}
	obj.obj.FlowNames = value

	return obj
}

type FlowMetricsRequestMetricNamesEnum string

// Enum of MetricNames on FlowMetricsRequest
var FlowMetricsRequestMetricNames = struct {
	TRANSMIT       FlowMetricsRequestMetricNamesEnum
	FRAMES_TX      FlowMetricsRequestMetricNamesEnum
	FRAMES_RX      FlowMetricsRequestMetricNamesEnum
	BYTES_TX       FlowMetricsRequestMetricNamesEnum
	BYTES_RX       FlowMetricsRequestMetricNamesEnum
	FRAMES_TX_RATE FlowMetricsRequestMetricNamesEnum
	FRAMES_RX_RATE FlowMetricsRequestMetricNamesEnum
	TX_L1_RATE_BPS FlowMetricsRequestMetricNamesEnum
	RX_L1_RATE_BPS FlowMetricsRequestMetricNamesEnum
	TX_RATE_BYTES  FlowMetricsRequestMetricNamesEnum
	RX_RATE_BYTES  FlowMetricsRequestMetricNamesEnum
	TX_RATE_BPS    FlowMetricsRequestMetricNamesEnum
	RX_RATE_BPS    FlowMetricsRequestMetricNamesEnum
	TX_RATE_KBPS   FlowMetricsRequestMetricNamesEnum
	RX_RATE_KBPS   FlowMetricsRequestMetricNamesEnum
	TX_RATE_MBPS   FlowMetricsRequestMetricNamesEnum
	RX_RATE_MBPS   FlowMetricsRequestMetricNamesEnum
}{
	TRANSMIT:       FlowMetricsRequestMetricNamesEnum("transmit"),
	FRAMES_TX:      FlowMetricsRequestMetricNamesEnum("frames_tx"),
	FRAMES_RX:      FlowMetricsRequestMetricNamesEnum("frames_rx"),
	BYTES_TX:       FlowMetricsRequestMetricNamesEnum("bytes_tx"),
	BYTES_RX:       FlowMetricsRequestMetricNamesEnum("bytes_rx"),
	FRAMES_TX_RATE: FlowMetricsRequestMetricNamesEnum("frames_tx_rate"),
	FRAMES_RX_RATE: FlowMetricsRequestMetricNamesEnum("frames_rx_rate"),
	TX_L1_RATE_BPS: FlowMetricsRequestMetricNamesEnum("tx_l1_rate_bps"),
	RX_L1_RATE_BPS: FlowMetricsRequestMetricNamesEnum("rx_l1_rate_bps"),
	TX_RATE_BYTES:  FlowMetricsRequestMetricNamesEnum("tx_rate_bytes"),
	RX_RATE_BYTES:  FlowMetricsRequestMetricNamesEnum("rx_rate_bytes"),
	TX_RATE_BPS:    FlowMetricsRequestMetricNamesEnum("tx_rate_bps"),
	RX_RATE_BPS:    FlowMetricsRequestMetricNamesEnum("rx_rate_bps"),
	TX_RATE_KBPS:   FlowMetricsRequestMetricNamesEnum("tx_rate_kbps"),
	RX_RATE_KBPS:   FlowMetricsRequestMetricNamesEnum("rx_rate_kbps"),
	TX_RATE_MBPS:   FlowMetricsRequestMetricNamesEnum("tx_rate_mbps"),
	RX_RATE_MBPS:   FlowMetricsRequestMetricNamesEnum("rx_rate_mbps"),
}

func (obj *flowMetricsRequest) MetricNames() []FlowMetricsRequestMetricNamesEnum {
	items := []FlowMetricsRequestMetricNamesEnum{}
	for _, item := range obj.obj.MetricNames {
		items = append(items, FlowMetricsRequestMetricNamesEnum(item.String()))
	}
	return items
}

// The list of metric names that the returned result set will contain. If the list is empty then all metrics will be returned.
// SetMetricNames sets the []string value in the FlowMetricsRequest object
func (obj *flowMetricsRequest) SetMetricNames(value []FlowMetricsRequestMetricNamesEnum) FlowMetricsRequest {

	items := []otg.FlowMetricsRequest_MetricNames_Enum{}
	for _, item := range value {
		intValue := otg.FlowMetricsRequest_MetricNames_Enum_value[string(item)]
		items = append(items, otg.FlowMetricsRequest_MetricNames_Enum(intValue))
	}
	obj.obj.MetricNames = items
	return obj
}

// description is TBD
// TaggedMetrics returns a FlowTaggedMetricsFilter
func (obj *flowMetricsRequest) TaggedMetrics() FlowTaggedMetricsFilter {
	if obj.obj.TaggedMetrics == nil {
		obj.obj.TaggedMetrics = NewFlowTaggedMetricsFilter().msg()
	}
	if obj.taggedMetricsHolder == nil {
		obj.taggedMetricsHolder = &flowTaggedMetricsFilter{obj: obj.obj.TaggedMetrics}
	}
	return obj.taggedMetricsHolder
}

// description is TBD
// TaggedMetrics returns a FlowTaggedMetricsFilter
func (obj *flowMetricsRequest) HasTaggedMetrics() bool {
	return obj.obj.TaggedMetrics != nil
}

// description is TBD
// SetTaggedMetrics sets the FlowTaggedMetricsFilter value in the FlowMetricsRequest object
func (obj *flowMetricsRequest) SetTaggedMetrics(value FlowTaggedMetricsFilter) FlowMetricsRequest {

	obj.taggedMetricsHolder = nil
	obj.obj.TaggedMetrics = value.msg()

	return obj
}

func (obj *flowMetricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.TaggedMetrics != nil {

		obj.TaggedMetrics().validateObj(vObj, set_default)
	}

}

func (obj *flowMetricsRequest) setDefault() {

}
