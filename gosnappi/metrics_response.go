package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MetricsResponse *****
type metricsResponse struct {
	validation
	obj                *otg.MetricsResponse
	marshaller         marshalMetricsResponse
	unMarshaller       unMarshalMetricsResponse
	portMetricsHolder  MetricsResponsePortMetricIter
	flowMetricsHolder  MetricsResponseFlowMetricIter
	bgpv4MetricsHolder MetricsResponseBgpv4MetricIter
	bgpv6MetricsHolder MetricsResponseBgpv6MetricIter
	isisMetricsHolder  MetricsResponseIsisMetricIter
	lagMetricsHolder   MetricsResponseLagMetricIter
	lacpMetricsHolder  MetricsResponseLacpMetricIter
	lldpMetricsHolder  MetricsResponseLldpMetricIter
	rsvpMetricsHolder  MetricsResponseRsvpMetricIter
}

func NewMetricsResponse() MetricsResponse {
	obj := metricsResponse{obj: &otg.MetricsResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *metricsResponse) msg() *otg.MetricsResponse {
	return obj.obj
}

func (obj *metricsResponse) setMsg(msg *otg.MetricsResponse) MetricsResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmetricsResponse struct {
	obj *metricsResponse
}

type marshalMetricsResponse interface {
	// ToProto marshals MetricsResponse to protobuf object *otg.MetricsResponse
	ToProto() (*otg.MetricsResponse, error)
	// ToPbText marshals MetricsResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MetricsResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals MetricsResponse to JSON text
	ToJson() (string, error)
}

type unMarshalmetricsResponse struct {
	obj *metricsResponse
}

type unMarshalMetricsResponse interface {
	// FromProto unmarshals MetricsResponse from protobuf object *otg.MetricsResponse
	FromProto(msg *otg.MetricsResponse) (MetricsResponse, error)
	// FromPbText unmarshals MetricsResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MetricsResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MetricsResponse from JSON text
	FromJson(value string) error
}

func (obj *metricsResponse) Marshal() marshalMetricsResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmetricsResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *metricsResponse) Unmarshal() unMarshalMetricsResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmetricsResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmetricsResponse) ToProto() (*otg.MetricsResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmetricsResponse) FromProto(msg *otg.MetricsResponse) (MetricsResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmetricsResponse) ToPbText() (string, error) {
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

func (m *unMarshalmetricsResponse) FromPbText(value string) error {
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

func (m *marshalmetricsResponse) ToYaml() (string, error) {
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

func (m *unMarshalmetricsResponse) FromYaml(value string) error {
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

func (m *marshalmetricsResponse) ToJson() (string, error) {
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

func (m *unMarshalmetricsResponse) FromJson(value string) error {
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

func (obj *metricsResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *metricsResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *metricsResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *metricsResponse) Clone() (MetricsResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMetricsResponse()
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

func (obj *metricsResponse) setNil() {
	obj.portMetricsHolder = nil
	obj.flowMetricsHolder = nil
	obj.bgpv4MetricsHolder = nil
	obj.bgpv6MetricsHolder = nil
	obj.isisMetricsHolder = nil
	obj.lagMetricsHolder = nil
	obj.lacpMetricsHolder = nil
	obj.lldpMetricsHolder = nil
	obj.rsvpMetricsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MetricsResponse is response containing chosen traffic generator metrics.
type MetricsResponse interface {
	Validation
	// msg marshals MetricsResponse to protobuf object *otg.MetricsResponse
	// and doesn't set defaults
	msg() *otg.MetricsResponse
	// setMsg unmarshals MetricsResponse from protobuf object *otg.MetricsResponse
	// and doesn't set defaults
	setMsg(*otg.MetricsResponse) MetricsResponse
	// provides marshal interface
	Marshal() marshalMetricsResponse
	// provides unmarshal interface
	Unmarshal() unMarshalMetricsResponse
	// validate validates MetricsResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MetricsResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MetricsResponseChoiceEnum, set in MetricsResponse
	Choice() MetricsResponseChoiceEnum
	// setChoice assigns MetricsResponseChoiceEnum provided by user to MetricsResponse
	setChoice(value MetricsResponseChoiceEnum) MetricsResponse
	// HasChoice checks if Choice has been set in MetricsResponse
	HasChoice() bool
	// PortMetrics returns MetricsResponsePortMetricIterIter, set in MetricsResponse
	PortMetrics() MetricsResponsePortMetricIter
	// FlowMetrics returns MetricsResponseFlowMetricIterIter, set in MetricsResponse
	FlowMetrics() MetricsResponseFlowMetricIter
	// Bgpv4Metrics returns MetricsResponseBgpv4MetricIterIter, set in MetricsResponse
	Bgpv4Metrics() MetricsResponseBgpv4MetricIter
	// Bgpv6Metrics returns MetricsResponseBgpv6MetricIterIter, set in MetricsResponse
	Bgpv6Metrics() MetricsResponseBgpv6MetricIter
	// IsisMetrics returns MetricsResponseIsisMetricIterIter, set in MetricsResponse
	IsisMetrics() MetricsResponseIsisMetricIter
	// LagMetrics returns MetricsResponseLagMetricIterIter, set in MetricsResponse
	LagMetrics() MetricsResponseLagMetricIter
	// LacpMetrics returns MetricsResponseLacpMetricIterIter, set in MetricsResponse
	LacpMetrics() MetricsResponseLacpMetricIter
	// LldpMetrics returns MetricsResponseLldpMetricIterIter, set in MetricsResponse
	LldpMetrics() MetricsResponseLldpMetricIter
	// RsvpMetrics returns MetricsResponseRsvpMetricIterIter, set in MetricsResponse
	RsvpMetrics() MetricsResponseRsvpMetricIter
	setNil()
}

type MetricsResponseChoiceEnum string

// Enum of Choice on MetricsResponse
var MetricsResponseChoice = struct {
	FLOW_METRICS  MetricsResponseChoiceEnum
	PORT_METRICS  MetricsResponseChoiceEnum
	BGPV4_METRICS MetricsResponseChoiceEnum
	BGPV6_METRICS MetricsResponseChoiceEnum
	ISIS_METRICS  MetricsResponseChoiceEnum
	LAG_METRICS   MetricsResponseChoiceEnum
	LACP_METRICS  MetricsResponseChoiceEnum
	LLDP_METRICS  MetricsResponseChoiceEnum
	RSVP_METRICS  MetricsResponseChoiceEnum
}{
	FLOW_METRICS:  MetricsResponseChoiceEnum("flow_metrics"),
	PORT_METRICS:  MetricsResponseChoiceEnum("port_metrics"),
	BGPV4_METRICS: MetricsResponseChoiceEnum("bgpv4_metrics"),
	BGPV6_METRICS: MetricsResponseChoiceEnum("bgpv6_metrics"),
	ISIS_METRICS:  MetricsResponseChoiceEnum("isis_metrics"),
	LAG_METRICS:   MetricsResponseChoiceEnum("lag_metrics"),
	LACP_METRICS:  MetricsResponseChoiceEnum("lacp_metrics"),
	LLDP_METRICS:  MetricsResponseChoiceEnum("lldp_metrics"),
	RSVP_METRICS:  MetricsResponseChoiceEnum("rsvp_metrics"),
}

func (obj *metricsResponse) Choice() MetricsResponseChoiceEnum {
	return MetricsResponseChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *metricsResponse) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *metricsResponse) setChoice(value MetricsResponseChoiceEnum) MetricsResponse {
	intValue, ok := otg.MetricsResponse_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MetricsResponseChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.MetricsResponse_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.RsvpMetrics = nil
	obj.rsvpMetricsHolder = nil
	obj.obj.LldpMetrics = nil
	obj.lldpMetricsHolder = nil
	obj.obj.LacpMetrics = nil
	obj.lacpMetricsHolder = nil
	obj.obj.LagMetrics = nil
	obj.lagMetricsHolder = nil
	obj.obj.IsisMetrics = nil
	obj.isisMetricsHolder = nil
	obj.obj.Bgpv6Metrics = nil
	obj.bgpv6MetricsHolder = nil
	obj.obj.Bgpv4Metrics = nil
	obj.bgpv4MetricsHolder = nil
	obj.obj.FlowMetrics = nil
	obj.flowMetricsHolder = nil
	obj.obj.PortMetrics = nil
	obj.portMetricsHolder = nil

	if value == MetricsResponseChoice.PORT_METRICS {
		obj.obj.PortMetrics = []*otg.PortMetric{}
	}

	if value == MetricsResponseChoice.FLOW_METRICS {
		obj.obj.FlowMetrics = []*otg.FlowMetric{}
	}

	if value == MetricsResponseChoice.BGPV4_METRICS {
		obj.obj.Bgpv4Metrics = []*otg.Bgpv4Metric{}
	}

	if value == MetricsResponseChoice.BGPV6_METRICS {
		obj.obj.Bgpv6Metrics = []*otg.Bgpv6Metric{}
	}

	if value == MetricsResponseChoice.ISIS_METRICS {
		obj.obj.IsisMetrics = []*otg.IsisMetric{}
	}

	if value == MetricsResponseChoice.LAG_METRICS {
		obj.obj.LagMetrics = []*otg.LagMetric{}
	}

	if value == MetricsResponseChoice.LACP_METRICS {
		obj.obj.LacpMetrics = []*otg.LacpMetric{}
	}

	if value == MetricsResponseChoice.LLDP_METRICS {
		obj.obj.LldpMetrics = []*otg.LldpMetric{}
	}

	if value == MetricsResponseChoice.RSVP_METRICS {
		obj.obj.RsvpMetrics = []*otg.RsvpMetric{}
	}

	return obj
}

// description is TBD
// PortMetrics returns a []PortMetric
func (obj *metricsResponse) PortMetrics() MetricsResponsePortMetricIter {
	if len(obj.obj.PortMetrics) == 0 {
		obj.setChoice(MetricsResponseChoice.PORT_METRICS)
	}
	if obj.portMetricsHolder == nil {
		obj.portMetricsHolder = newMetricsResponsePortMetricIter(&obj.obj.PortMetrics).setMsg(obj)
	}
	return obj.portMetricsHolder
}

type metricsResponsePortMetricIter struct {
	obj             *metricsResponse
	portMetricSlice []PortMetric
	fieldPtr        *[]*otg.PortMetric
}

func newMetricsResponsePortMetricIter(ptr *[]*otg.PortMetric) MetricsResponsePortMetricIter {
	return &metricsResponsePortMetricIter{fieldPtr: ptr}
}

type MetricsResponsePortMetricIter interface {
	setMsg(*metricsResponse) MetricsResponsePortMetricIter
	Items() []PortMetric
	Add() PortMetric
	Append(items ...PortMetric) MetricsResponsePortMetricIter
	Set(index int, newObj PortMetric) MetricsResponsePortMetricIter
	Clear() MetricsResponsePortMetricIter
	clearHolderSlice() MetricsResponsePortMetricIter
	appendHolderSlice(item PortMetric) MetricsResponsePortMetricIter
}

func (obj *metricsResponsePortMetricIter) setMsg(msg *metricsResponse) MetricsResponsePortMetricIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&portMetric{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *metricsResponsePortMetricIter) Items() []PortMetric {
	return obj.portMetricSlice
}

func (obj *metricsResponsePortMetricIter) Add() PortMetric {
	newObj := &otg.PortMetric{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &portMetric{obj: newObj}
	newLibObj.setDefault()
	obj.portMetricSlice = append(obj.portMetricSlice, newLibObj)
	return newLibObj
}

func (obj *metricsResponsePortMetricIter) Append(items ...PortMetric) MetricsResponsePortMetricIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.portMetricSlice = append(obj.portMetricSlice, item)
	}
	return obj
}

func (obj *metricsResponsePortMetricIter) Set(index int, newObj PortMetric) MetricsResponsePortMetricIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.portMetricSlice[index] = newObj
	return obj
}
func (obj *metricsResponsePortMetricIter) Clear() MetricsResponsePortMetricIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PortMetric{}
		obj.portMetricSlice = []PortMetric{}
	}
	return obj
}
func (obj *metricsResponsePortMetricIter) clearHolderSlice() MetricsResponsePortMetricIter {
	if len(obj.portMetricSlice) > 0 {
		obj.portMetricSlice = []PortMetric{}
	}
	return obj
}
func (obj *metricsResponsePortMetricIter) appendHolderSlice(item PortMetric) MetricsResponsePortMetricIter {
	obj.portMetricSlice = append(obj.portMetricSlice, item)
	return obj
}

// description is TBD
// FlowMetrics returns a []FlowMetric
func (obj *metricsResponse) FlowMetrics() MetricsResponseFlowMetricIter {
	if len(obj.obj.FlowMetrics) == 0 {
		obj.setChoice(MetricsResponseChoice.FLOW_METRICS)
	}
	if obj.flowMetricsHolder == nil {
		obj.flowMetricsHolder = newMetricsResponseFlowMetricIter(&obj.obj.FlowMetrics).setMsg(obj)
	}
	return obj.flowMetricsHolder
}

type metricsResponseFlowMetricIter struct {
	obj             *metricsResponse
	flowMetricSlice []FlowMetric
	fieldPtr        *[]*otg.FlowMetric
}

func newMetricsResponseFlowMetricIter(ptr *[]*otg.FlowMetric) MetricsResponseFlowMetricIter {
	return &metricsResponseFlowMetricIter{fieldPtr: ptr}
}

type MetricsResponseFlowMetricIter interface {
	setMsg(*metricsResponse) MetricsResponseFlowMetricIter
	Items() []FlowMetric
	Add() FlowMetric
	Append(items ...FlowMetric) MetricsResponseFlowMetricIter
	Set(index int, newObj FlowMetric) MetricsResponseFlowMetricIter
	Clear() MetricsResponseFlowMetricIter
	clearHolderSlice() MetricsResponseFlowMetricIter
	appendHolderSlice(item FlowMetric) MetricsResponseFlowMetricIter
}

func (obj *metricsResponseFlowMetricIter) setMsg(msg *metricsResponse) MetricsResponseFlowMetricIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowMetric{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *metricsResponseFlowMetricIter) Items() []FlowMetric {
	return obj.flowMetricSlice
}

func (obj *metricsResponseFlowMetricIter) Add() FlowMetric {
	newObj := &otg.FlowMetric{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowMetric{obj: newObj}
	newLibObj.setDefault()
	obj.flowMetricSlice = append(obj.flowMetricSlice, newLibObj)
	return newLibObj
}

func (obj *metricsResponseFlowMetricIter) Append(items ...FlowMetric) MetricsResponseFlowMetricIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowMetricSlice = append(obj.flowMetricSlice, item)
	}
	return obj
}

func (obj *metricsResponseFlowMetricIter) Set(index int, newObj FlowMetric) MetricsResponseFlowMetricIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowMetricSlice[index] = newObj
	return obj
}
func (obj *metricsResponseFlowMetricIter) Clear() MetricsResponseFlowMetricIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowMetric{}
		obj.flowMetricSlice = []FlowMetric{}
	}
	return obj
}
func (obj *metricsResponseFlowMetricIter) clearHolderSlice() MetricsResponseFlowMetricIter {
	if len(obj.flowMetricSlice) > 0 {
		obj.flowMetricSlice = []FlowMetric{}
	}
	return obj
}
func (obj *metricsResponseFlowMetricIter) appendHolderSlice(item FlowMetric) MetricsResponseFlowMetricIter {
	obj.flowMetricSlice = append(obj.flowMetricSlice, item)
	return obj
}

// description is TBD
// Bgpv4Metrics returns a []Bgpv4Metric
func (obj *metricsResponse) Bgpv4Metrics() MetricsResponseBgpv4MetricIter {
	if len(obj.obj.Bgpv4Metrics) == 0 {
		obj.setChoice(MetricsResponseChoice.BGPV4_METRICS)
	}
	if obj.bgpv4MetricsHolder == nil {
		obj.bgpv4MetricsHolder = newMetricsResponseBgpv4MetricIter(&obj.obj.Bgpv4Metrics).setMsg(obj)
	}
	return obj.bgpv4MetricsHolder
}

type metricsResponseBgpv4MetricIter struct {
	obj              *metricsResponse
	bgpv4MetricSlice []Bgpv4Metric
	fieldPtr         *[]*otg.Bgpv4Metric
}

func newMetricsResponseBgpv4MetricIter(ptr *[]*otg.Bgpv4Metric) MetricsResponseBgpv4MetricIter {
	return &metricsResponseBgpv4MetricIter{fieldPtr: ptr}
}

type MetricsResponseBgpv4MetricIter interface {
	setMsg(*metricsResponse) MetricsResponseBgpv4MetricIter
	Items() []Bgpv4Metric
	Add() Bgpv4Metric
	Append(items ...Bgpv4Metric) MetricsResponseBgpv4MetricIter
	Set(index int, newObj Bgpv4Metric) MetricsResponseBgpv4MetricIter
	Clear() MetricsResponseBgpv4MetricIter
	clearHolderSlice() MetricsResponseBgpv4MetricIter
	appendHolderSlice(item Bgpv4Metric) MetricsResponseBgpv4MetricIter
}

func (obj *metricsResponseBgpv4MetricIter) setMsg(msg *metricsResponse) MetricsResponseBgpv4MetricIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpv4Metric{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *metricsResponseBgpv4MetricIter) Items() []Bgpv4Metric {
	return obj.bgpv4MetricSlice
}

func (obj *metricsResponseBgpv4MetricIter) Add() Bgpv4Metric {
	newObj := &otg.Bgpv4Metric{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpv4Metric{obj: newObj}
	newLibObj.setDefault()
	obj.bgpv4MetricSlice = append(obj.bgpv4MetricSlice, newLibObj)
	return newLibObj
}

func (obj *metricsResponseBgpv4MetricIter) Append(items ...Bgpv4Metric) MetricsResponseBgpv4MetricIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpv4MetricSlice = append(obj.bgpv4MetricSlice, item)
	}
	return obj
}

func (obj *metricsResponseBgpv4MetricIter) Set(index int, newObj Bgpv4Metric) MetricsResponseBgpv4MetricIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpv4MetricSlice[index] = newObj
	return obj
}
func (obj *metricsResponseBgpv4MetricIter) Clear() MetricsResponseBgpv4MetricIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Bgpv4Metric{}
		obj.bgpv4MetricSlice = []Bgpv4Metric{}
	}
	return obj
}
func (obj *metricsResponseBgpv4MetricIter) clearHolderSlice() MetricsResponseBgpv4MetricIter {
	if len(obj.bgpv4MetricSlice) > 0 {
		obj.bgpv4MetricSlice = []Bgpv4Metric{}
	}
	return obj
}
func (obj *metricsResponseBgpv4MetricIter) appendHolderSlice(item Bgpv4Metric) MetricsResponseBgpv4MetricIter {
	obj.bgpv4MetricSlice = append(obj.bgpv4MetricSlice, item)
	return obj
}

// description is TBD
// Bgpv6Metrics returns a []Bgpv6Metric
func (obj *metricsResponse) Bgpv6Metrics() MetricsResponseBgpv6MetricIter {
	if len(obj.obj.Bgpv6Metrics) == 0 {
		obj.setChoice(MetricsResponseChoice.BGPV6_METRICS)
	}
	if obj.bgpv6MetricsHolder == nil {
		obj.bgpv6MetricsHolder = newMetricsResponseBgpv6MetricIter(&obj.obj.Bgpv6Metrics).setMsg(obj)
	}
	return obj.bgpv6MetricsHolder
}

type metricsResponseBgpv6MetricIter struct {
	obj              *metricsResponse
	bgpv6MetricSlice []Bgpv6Metric
	fieldPtr         *[]*otg.Bgpv6Metric
}

func newMetricsResponseBgpv6MetricIter(ptr *[]*otg.Bgpv6Metric) MetricsResponseBgpv6MetricIter {
	return &metricsResponseBgpv6MetricIter{fieldPtr: ptr}
}

type MetricsResponseBgpv6MetricIter interface {
	setMsg(*metricsResponse) MetricsResponseBgpv6MetricIter
	Items() []Bgpv6Metric
	Add() Bgpv6Metric
	Append(items ...Bgpv6Metric) MetricsResponseBgpv6MetricIter
	Set(index int, newObj Bgpv6Metric) MetricsResponseBgpv6MetricIter
	Clear() MetricsResponseBgpv6MetricIter
	clearHolderSlice() MetricsResponseBgpv6MetricIter
	appendHolderSlice(item Bgpv6Metric) MetricsResponseBgpv6MetricIter
}

func (obj *metricsResponseBgpv6MetricIter) setMsg(msg *metricsResponse) MetricsResponseBgpv6MetricIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpv6Metric{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *metricsResponseBgpv6MetricIter) Items() []Bgpv6Metric {
	return obj.bgpv6MetricSlice
}

func (obj *metricsResponseBgpv6MetricIter) Add() Bgpv6Metric {
	newObj := &otg.Bgpv6Metric{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpv6Metric{obj: newObj}
	newLibObj.setDefault()
	obj.bgpv6MetricSlice = append(obj.bgpv6MetricSlice, newLibObj)
	return newLibObj
}

func (obj *metricsResponseBgpv6MetricIter) Append(items ...Bgpv6Metric) MetricsResponseBgpv6MetricIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpv6MetricSlice = append(obj.bgpv6MetricSlice, item)
	}
	return obj
}

func (obj *metricsResponseBgpv6MetricIter) Set(index int, newObj Bgpv6Metric) MetricsResponseBgpv6MetricIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpv6MetricSlice[index] = newObj
	return obj
}
func (obj *metricsResponseBgpv6MetricIter) Clear() MetricsResponseBgpv6MetricIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Bgpv6Metric{}
		obj.bgpv6MetricSlice = []Bgpv6Metric{}
	}
	return obj
}
func (obj *metricsResponseBgpv6MetricIter) clearHolderSlice() MetricsResponseBgpv6MetricIter {
	if len(obj.bgpv6MetricSlice) > 0 {
		obj.bgpv6MetricSlice = []Bgpv6Metric{}
	}
	return obj
}
func (obj *metricsResponseBgpv6MetricIter) appendHolderSlice(item Bgpv6Metric) MetricsResponseBgpv6MetricIter {
	obj.bgpv6MetricSlice = append(obj.bgpv6MetricSlice, item)
	return obj
}

// description is TBD
// IsisMetrics returns a []IsisMetric
func (obj *metricsResponse) IsisMetrics() MetricsResponseIsisMetricIter {
	if len(obj.obj.IsisMetrics) == 0 {
		obj.setChoice(MetricsResponseChoice.ISIS_METRICS)
	}
	if obj.isisMetricsHolder == nil {
		obj.isisMetricsHolder = newMetricsResponseIsisMetricIter(&obj.obj.IsisMetrics).setMsg(obj)
	}
	return obj.isisMetricsHolder
}

type metricsResponseIsisMetricIter struct {
	obj             *metricsResponse
	isisMetricSlice []IsisMetric
	fieldPtr        *[]*otg.IsisMetric
}

func newMetricsResponseIsisMetricIter(ptr *[]*otg.IsisMetric) MetricsResponseIsisMetricIter {
	return &metricsResponseIsisMetricIter{fieldPtr: ptr}
}

type MetricsResponseIsisMetricIter interface {
	setMsg(*metricsResponse) MetricsResponseIsisMetricIter
	Items() []IsisMetric
	Add() IsisMetric
	Append(items ...IsisMetric) MetricsResponseIsisMetricIter
	Set(index int, newObj IsisMetric) MetricsResponseIsisMetricIter
	Clear() MetricsResponseIsisMetricIter
	clearHolderSlice() MetricsResponseIsisMetricIter
	appendHolderSlice(item IsisMetric) MetricsResponseIsisMetricIter
}

func (obj *metricsResponseIsisMetricIter) setMsg(msg *metricsResponse) MetricsResponseIsisMetricIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisMetric{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *metricsResponseIsisMetricIter) Items() []IsisMetric {
	return obj.isisMetricSlice
}

func (obj *metricsResponseIsisMetricIter) Add() IsisMetric {
	newObj := &otg.IsisMetric{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisMetric{obj: newObj}
	newLibObj.setDefault()
	obj.isisMetricSlice = append(obj.isisMetricSlice, newLibObj)
	return newLibObj
}

func (obj *metricsResponseIsisMetricIter) Append(items ...IsisMetric) MetricsResponseIsisMetricIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisMetricSlice = append(obj.isisMetricSlice, item)
	}
	return obj
}

func (obj *metricsResponseIsisMetricIter) Set(index int, newObj IsisMetric) MetricsResponseIsisMetricIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisMetricSlice[index] = newObj
	return obj
}
func (obj *metricsResponseIsisMetricIter) Clear() MetricsResponseIsisMetricIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisMetric{}
		obj.isisMetricSlice = []IsisMetric{}
	}
	return obj
}
func (obj *metricsResponseIsisMetricIter) clearHolderSlice() MetricsResponseIsisMetricIter {
	if len(obj.isisMetricSlice) > 0 {
		obj.isisMetricSlice = []IsisMetric{}
	}
	return obj
}
func (obj *metricsResponseIsisMetricIter) appendHolderSlice(item IsisMetric) MetricsResponseIsisMetricIter {
	obj.isisMetricSlice = append(obj.isisMetricSlice, item)
	return obj
}

// description is TBD
// LagMetrics returns a []LagMetric
func (obj *metricsResponse) LagMetrics() MetricsResponseLagMetricIter {
	if len(obj.obj.LagMetrics) == 0 {
		obj.setChoice(MetricsResponseChoice.LAG_METRICS)
	}
	if obj.lagMetricsHolder == nil {
		obj.lagMetricsHolder = newMetricsResponseLagMetricIter(&obj.obj.LagMetrics).setMsg(obj)
	}
	return obj.lagMetricsHolder
}

type metricsResponseLagMetricIter struct {
	obj            *metricsResponse
	lagMetricSlice []LagMetric
	fieldPtr       *[]*otg.LagMetric
}

func newMetricsResponseLagMetricIter(ptr *[]*otg.LagMetric) MetricsResponseLagMetricIter {
	return &metricsResponseLagMetricIter{fieldPtr: ptr}
}

type MetricsResponseLagMetricIter interface {
	setMsg(*metricsResponse) MetricsResponseLagMetricIter
	Items() []LagMetric
	Add() LagMetric
	Append(items ...LagMetric) MetricsResponseLagMetricIter
	Set(index int, newObj LagMetric) MetricsResponseLagMetricIter
	Clear() MetricsResponseLagMetricIter
	clearHolderSlice() MetricsResponseLagMetricIter
	appendHolderSlice(item LagMetric) MetricsResponseLagMetricIter
}

func (obj *metricsResponseLagMetricIter) setMsg(msg *metricsResponse) MetricsResponseLagMetricIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&lagMetric{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *metricsResponseLagMetricIter) Items() []LagMetric {
	return obj.lagMetricSlice
}

func (obj *metricsResponseLagMetricIter) Add() LagMetric {
	newObj := &otg.LagMetric{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &lagMetric{obj: newObj}
	newLibObj.setDefault()
	obj.lagMetricSlice = append(obj.lagMetricSlice, newLibObj)
	return newLibObj
}

func (obj *metricsResponseLagMetricIter) Append(items ...LagMetric) MetricsResponseLagMetricIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.lagMetricSlice = append(obj.lagMetricSlice, item)
	}
	return obj
}

func (obj *metricsResponseLagMetricIter) Set(index int, newObj LagMetric) MetricsResponseLagMetricIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.lagMetricSlice[index] = newObj
	return obj
}
func (obj *metricsResponseLagMetricIter) Clear() MetricsResponseLagMetricIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.LagMetric{}
		obj.lagMetricSlice = []LagMetric{}
	}
	return obj
}
func (obj *metricsResponseLagMetricIter) clearHolderSlice() MetricsResponseLagMetricIter {
	if len(obj.lagMetricSlice) > 0 {
		obj.lagMetricSlice = []LagMetric{}
	}
	return obj
}
func (obj *metricsResponseLagMetricIter) appendHolderSlice(item LagMetric) MetricsResponseLagMetricIter {
	obj.lagMetricSlice = append(obj.lagMetricSlice, item)
	return obj
}

// description is TBD
// LacpMetrics returns a []LacpMetric
func (obj *metricsResponse) LacpMetrics() MetricsResponseLacpMetricIter {
	if len(obj.obj.LacpMetrics) == 0 {
		obj.setChoice(MetricsResponseChoice.LACP_METRICS)
	}
	if obj.lacpMetricsHolder == nil {
		obj.lacpMetricsHolder = newMetricsResponseLacpMetricIter(&obj.obj.LacpMetrics).setMsg(obj)
	}
	return obj.lacpMetricsHolder
}

type metricsResponseLacpMetricIter struct {
	obj             *metricsResponse
	lacpMetricSlice []LacpMetric
	fieldPtr        *[]*otg.LacpMetric
}

func newMetricsResponseLacpMetricIter(ptr *[]*otg.LacpMetric) MetricsResponseLacpMetricIter {
	return &metricsResponseLacpMetricIter{fieldPtr: ptr}
}

type MetricsResponseLacpMetricIter interface {
	setMsg(*metricsResponse) MetricsResponseLacpMetricIter
	Items() []LacpMetric
	Add() LacpMetric
	Append(items ...LacpMetric) MetricsResponseLacpMetricIter
	Set(index int, newObj LacpMetric) MetricsResponseLacpMetricIter
	Clear() MetricsResponseLacpMetricIter
	clearHolderSlice() MetricsResponseLacpMetricIter
	appendHolderSlice(item LacpMetric) MetricsResponseLacpMetricIter
}

func (obj *metricsResponseLacpMetricIter) setMsg(msg *metricsResponse) MetricsResponseLacpMetricIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&lacpMetric{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *metricsResponseLacpMetricIter) Items() []LacpMetric {
	return obj.lacpMetricSlice
}

func (obj *metricsResponseLacpMetricIter) Add() LacpMetric {
	newObj := &otg.LacpMetric{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &lacpMetric{obj: newObj}
	newLibObj.setDefault()
	obj.lacpMetricSlice = append(obj.lacpMetricSlice, newLibObj)
	return newLibObj
}

func (obj *metricsResponseLacpMetricIter) Append(items ...LacpMetric) MetricsResponseLacpMetricIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.lacpMetricSlice = append(obj.lacpMetricSlice, item)
	}
	return obj
}

func (obj *metricsResponseLacpMetricIter) Set(index int, newObj LacpMetric) MetricsResponseLacpMetricIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.lacpMetricSlice[index] = newObj
	return obj
}
func (obj *metricsResponseLacpMetricIter) Clear() MetricsResponseLacpMetricIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.LacpMetric{}
		obj.lacpMetricSlice = []LacpMetric{}
	}
	return obj
}
func (obj *metricsResponseLacpMetricIter) clearHolderSlice() MetricsResponseLacpMetricIter {
	if len(obj.lacpMetricSlice) > 0 {
		obj.lacpMetricSlice = []LacpMetric{}
	}
	return obj
}
func (obj *metricsResponseLacpMetricIter) appendHolderSlice(item LacpMetric) MetricsResponseLacpMetricIter {
	obj.lacpMetricSlice = append(obj.lacpMetricSlice, item)
	return obj
}

// description is TBD
// LldpMetrics returns a []LldpMetric
func (obj *metricsResponse) LldpMetrics() MetricsResponseLldpMetricIter {
	if len(obj.obj.LldpMetrics) == 0 {
		obj.setChoice(MetricsResponseChoice.LLDP_METRICS)
	}
	if obj.lldpMetricsHolder == nil {
		obj.lldpMetricsHolder = newMetricsResponseLldpMetricIter(&obj.obj.LldpMetrics).setMsg(obj)
	}
	return obj.lldpMetricsHolder
}

type metricsResponseLldpMetricIter struct {
	obj             *metricsResponse
	lldpMetricSlice []LldpMetric
	fieldPtr        *[]*otg.LldpMetric
}

func newMetricsResponseLldpMetricIter(ptr *[]*otg.LldpMetric) MetricsResponseLldpMetricIter {
	return &metricsResponseLldpMetricIter{fieldPtr: ptr}
}

type MetricsResponseLldpMetricIter interface {
	setMsg(*metricsResponse) MetricsResponseLldpMetricIter
	Items() []LldpMetric
	Add() LldpMetric
	Append(items ...LldpMetric) MetricsResponseLldpMetricIter
	Set(index int, newObj LldpMetric) MetricsResponseLldpMetricIter
	Clear() MetricsResponseLldpMetricIter
	clearHolderSlice() MetricsResponseLldpMetricIter
	appendHolderSlice(item LldpMetric) MetricsResponseLldpMetricIter
}

func (obj *metricsResponseLldpMetricIter) setMsg(msg *metricsResponse) MetricsResponseLldpMetricIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&lldpMetric{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *metricsResponseLldpMetricIter) Items() []LldpMetric {
	return obj.lldpMetricSlice
}

func (obj *metricsResponseLldpMetricIter) Add() LldpMetric {
	newObj := &otg.LldpMetric{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &lldpMetric{obj: newObj}
	newLibObj.setDefault()
	obj.lldpMetricSlice = append(obj.lldpMetricSlice, newLibObj)
	return newLibObj
}

func (obj *metricsResponseLldpMetricIter) Append(items ...LldpMetric) MetricsResponseLldpMetricIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.lldpMetricSlice = append(obj.lldpMetricSlice, item)
	}
	return obj
}

func (obj *metricsResponseLldpMetricIter) Set(index int, newObj LldpMetric) MetricsResponseLldpMetricIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.lldpMetricSlice[index] = newObj
	return obj
}
func (obj *metricsResponseLldpMetricIter) Clear() MetricsResponseLldpMetricIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.LldpMetric{}
		obj.lldpMetricSlice = []LldpMetric{}
	}
	return obj
}
func (obj *metricsResponseLldpMetricIter) clearHolderSlice() MetricsResponseLldpMetricIter {
	if len(obj.lldpMetricSlice) > 0 {
		obj.lldpMetricSlice = []LldpMetric{}
	}
	return obj
}
func (obj *metricsResponseLldpMetricIter) appendHolderSlice(item LldpMetric) MetricsResponseLldpMetricIter {
	obj.lldpMetricSlice = append(obj.lldpMetricSlice, item)
	return obj
}

// description is TBD
// RsvpMetrics returns a []RsvpMetric
func (obj *metricsResponse) RsvpMetrics() MetricsResponseRsvpMetricIter {
	if len(obj.obj.RsvpMetrics) == 0 {
		obj.setChoice(MetricsResponseChoice.RSVP_METRICS)
	}
	if obj.rsvpMetricsHolder == nil {
		obj.rsvpMetricsHolder = newMetricsResponseRsvpMetricIter(&obj.obj.RsvpMetrics).setMsg(obj)
	}
	return obj.rsvpMetricsHolder
}

type metricsResponseRsvpMetricIter struct {
	obj             *metricsResponse
	rsvpMetricSlice []RsvpMetric
	fieldPtr        *[]*otg.RsvpMetric
}

func newMetricsResponseRsvpMetricIter(ptr *[]*otg.RsvpMetric) MetricsResponseRsvpMetricIter {
	return &metricsResponseRsvpMetricIter{fieldPtr: ptr}
}

type MetricsResponseRsvpMetricIter interface {
	setMsg(*metricsResponse) MetricsResponseRsvpMetricIter
	Items() []RsvpMetric
	Add() RsvpMetric
	Append(items ...RsvpMetric) MetricsResponseRsvpMetricIter
	Set(index int, newObj RsvpMetric) MetricsResponseRsvpMetricIter
	Clear() MetricsResponseRsvpMetricIter
	clearHolderSlice() MetricsResponseRsvpMetricIter
	appendHolderSlice(item RsvpMetric) MetricsResponseRsvpMetricIter
}

func (obj *metricsResponseRsvpMetricIter) setMsg(msg *metricsResponse) MetricsResponseRsvpMetricIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&rsvpMetric{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *metricsResponseRsvpMetricIter) Items() []RsvpMetric {
	return obj.rsvpMetricSlice
}

func (obj *metricsResponseRsvpMetricIter) Add() RsvpMetric {
	newObj := &otg.RsvpMetric{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &rsvpMetric{obj: newObj}
	newLibObj.setDefault()
	obj.rsvpMetricSlice = append(obj.rsvpMetricSlice, newLibObj)
	return newLibObj
}

func (obj *metricsResponseRsvpMetricIter) Append(items ...RsvpMetric) MetricsResponseRsvpMetricIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.rsvpMetricSlice = append(obj.rsvpMetricSlice, item)
	}
	return obj
}

func (obj *metricsResponseRsvpMetricIter) Set(index int, newObj RsvpMetric) MetricsResponseRsvpMetricIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.rsvpMetricSlice[index] = newObj
	return obj
}
func (obj *metricsResponseRsvpMetricIter) Clear() MetricsResponseRsvpMetricIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.RsvpMetric{}
		obj.rsvpMetricSlice = []RsvpMetric{}
	}
	return obj
}
func (obj *metricsResponseRsvpMetricIter) clearHolderSlice() MetricsResponseRsvpMetricIter {
	if len(obj.rsvpMetricSlice) > 0 {
		obj.rsvpMetricSlice = []RsvpMetric{}
	}
	return obj
}
func (obj *metricsResponseRsvpMetricIter) appendHolderSlice(item RsvpMetric) MetricsResponseRsvpMetricIter {
	obj.rsvpMetricSlice = append(obj.rsvpMetricSlice, item)
	return obj
}

func (obj *metricsResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.PortMetrics) != 0 {

		if set_default {
			obj.PortMetrics().clearHolderSlice()
			for _, item := range obj.obj.PortMetrics {
				obj.PortMetrics().appendHolderSlice(&portMetric{obj: item})
			}
		}
		for _, item := range obj.PortMetrics().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.FlowMetrics) != 0 {

		if set_default {
			obj.FlowMetrics().clearHolderSlice()
			for _, item := range obj.obj.FlowMetrics {
				obj.FlowMetrics().appendHolderSlice(&flowMetric{obj: item})
			}
		}
		for _, item := range obj.FlowMetrics().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Bgpv4Metrics) != 0 {

		if set_default {
			obj.Bgpv4Metrics().clearHolderSlice()
			for _, item := range obj.obj.Bgpv4Metrics {
				obj.Bgpv4Metrics().appendHolderSlice(&bgpv4Metric{obj: item})
			}
		}
		for _, item := range obj.Bgpv4Metrics().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Bgpv6Metrics) != 0 {

		if set_default {
			obj.Bgpv6Metrics().clearHolderSlice()
			for _, item := range obj.obj.Bgpv6Metrics {
				obj.Bgpv6Metrics().appendHolderSlice(&bgpv6Metric{obj: item})
			}
		}
		for _, item := range obj.Bgpv6Metrics().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.IsisMetrics) != 0 {

		if set_default {
			obj.IsisMetrics().clearHolderSlice()
			for _, item := range obj.obj.IsisMetrics {
				obj.IsisMetrics().appendHolderSlice(&isisMetric{obj: item})
			}
		}
		for _, item := range obj.IsisMetrics().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.LagMetrics) != 0 {

		if set_default {
			obj.LagMetrics().clearHolderSlice()
			for _, item := range obj.obj.LagMetrics {
				obj.LagMetrics().appendHolderSlice(&lagMetric{obj: item})
			}
		}
		for _, item := range obj.LagMetrics().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.LacpMetrics) != 0 {

		if set_default {
			obj.LacpMetrics().clearHolderSlice()
			for _, item := range obj.obj.LacpMetrics {
				obj.LacpMetrics().appendHolderSlice(&lacpMetric{obj: item})
			}
		}
		for _, item := range obj.LacpMetrics().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.LldpMetrics) != 0 {

		if set_default {
			obj.LldpMetrics().clearHolderSlice()
			for _, item := range obj.obj.LldpMetrics {
				obj.LldpMetrics().appendHolderSlice(&lldpMetric{obj: item})
			}
		}
		for _, item := range obj.LldpMetrics().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.RsvpMetrics) != 0 {

		if set_default {
			obj.RsvpMetrics().clearHolderSlice()
			for _, item := range obj.obj.RsvpMetrics {
				obj.RsvpMetrics().appendHolderSlice(&rsvpMetric{obj: item})
			}
		}
		for _, item := range obj.RsvpMetrics().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *metricsResponse) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(MetricsResponseChoice.PORT_METRICS)

	}

}