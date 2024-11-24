package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowTaggedMetricsFilter *****
type flowTaggedMetricsFilter struct {
	validation
	obj           *otg.FlowTaggedMetricsFilter
	marshaller    marshalFlowTaggedMetricsFilter
	unMarshaller  unMarshalFlowTaggedMetricsFilter
	filtersHolder FlowTaggedMetricsFilterFlowMetricTagFilterIter
}

func NewFlowTaggedMetricsFilter() FlowTaggedMetricsFilter {
	obj := flowTaggedMetricsFilter{obj: &otg.FlowTaggedMetricsFilter{}}
	obj.setDefault()
	return &obj
}

func (obj *flowTaggedMetricsFilter) msg() *otg.FlowTaggedMetricsFilter {
	return obj.obj
}

func (obj *flowTaggedMetricsFilter) setMsg(msg *otg.FlowTaggedMetricsFilter) FlowTaggedMetricsFilter {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowTaggedMetricsFilter struct {
	obj *flowTaggedMetricsFilter
}

type marshalFlowTaggedMetricsFilter interface {
	// ToProto marshals FlowTaggedMetricsFilter to protobuf object *otg.FlowTaggedMetricsFilter
	ToProto() (*otg.FlowTaggedMetricsFilter, error)
	// ToPbText marshals FlowTaggedMetricsFilter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowTaggedMetricsFilter to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowTaggedMetricsFilter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowTaggedMetricsFilter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowTaggedMetricsFilter struct {
	obj *flowTaggedMetricsFilter
}

type unMarshalFlowTaggedMetricsFilter interface {
	// FromProto unmarshals FlowTaggedMetricsFilter from protobuf object *otg.FlowTaggedMetricsFilter
	FromProto(msg *otg.FlowTaggedMetricsFilter) (FlowTaggedMetricsFilter, error)
	// FromPbText unmarshals FlowTaggedMetricsFilter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowTaggedMetricsFilter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowTaggedMetricsFilter from JSON text
	FromJson(value string) error
}

func (obj *flowTaggedMetricsFilter) Marshal() marshalFlowTaggedMetricsFilter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowTaggedMetricsFilter{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowTaggedMetricsFilter) Unmarshal() unMarshalFlowTaggedMetricsFilter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowTaggedMetricsFilter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowTaggedMetricsFilter) ToProto() (*otg.FlowTaggedMetricsFilter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowTaggedMetricsFilter) FromProto(msg *otg.FlowTaggedMetricsFilter) (FlowTaggedMetricsFilter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowTaggedMetricsFilter) ToPbText() (string, error) {
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

func (m *unMarshalflowTaggedMetricsFilter) FromPbText(value string) error {
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

func (m *marshalflowTaggedMetricsFilter) ToYaml() (string, error) {
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

func (m *unMarshalflowTaggedMetricsFilter) FromYaml(value string) error {
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

func (m *marshalflowTaggedMetricsFilter) ToJsonRaw() (string, error) {
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

func (m *marshalflowTaggedMetricsFilter) ToJson() (string, error) {
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

func (m *unMarshalflowTaggedMetricsFilter) FromJson(value string) error {
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

func (obj *flowTaggedMetricsFilter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowTaggedMetricsFilter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowTaggedMetricsFilter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowTaggedMetricsFilter) Clone() (FlowTaggedMetricsFilter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowTaggedMetricsFilter()
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

func (obj *flowTaggedMetricsFilter) setNil() {
	obj.filtersHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowTaggedMetricsFilter is filter for tagged metrics
type FlowTaggedMetricsFilter interface {
	Validation
	// msg marshals FlowTaggedMetricsFilter to protobuf object *otg.FlowTaggedMetricsFilter
	// and doesn't set defaults
	msg() *otg.FlowTaggedMetricsFilter
	// setMsg unmarshals FlowTaggedMetricsFilter from protobuf object *otg.FlowTaggedMetricsFilter
	// and doesn't set defaults
	setMsg(*otg.FlowTaggedMetricsFilter) FlowTaggedMetricsFilter
	// provides marshal interface
	Marshal() marshalFlowTaggedMetricsFilter
	// provides unmarshal interface
	Unmarshal() unMarshalFlowTaggedMetricsFilter
	// validate validates FlowTaggedMetricsFilter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowTaggedMetricsFilter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Include returns bool, set in FlowTaggedMetricsFilter.
	Include() bool
	// SetInclude assigns bool provided by user to FlowTaggedMetricsFilter
	SetInclude(value bool) FlowTaggedMetricsFilter
	// HasInclude checks if Include has been set in FlowTaggedMetricsFilter
	HasInclude() bool
	// IncludeEmptyMetrics returns bool, set in FlowTaggedMetricsFilter.
	IncludeEmptyMetrics() bool
	// SetIncludeEmptyMetrics assigns bool provided by user to FlowTaggedMetricsFilter
	SetIncludeEmptyMetrics(value bool) FlowTaggedMetricsFilter
	// HasIncludeEmptyMetrics checks if IncludeEmptyMetrics has been set in FlowTaggedMetricsFilter
	HasIncludeEmptyMetrics() bool
	// MetricNames returns []FlowTaggedMetricsFilterMetricNamesEnum, set in FlowTaggedMetricsFilter
	MetricNames() []FlowTaggedMetricsFilterMetricNamesEnum
	// SetMetricNames assigns []FlowTaggedMetricsFilterMetricNamesEnum provided by user to FlowTaggedMetricsFilter
	SetMetricNames(value []FlowTaggedMetricsFilterMetricNamesEnum) FlowTaggedMetricsFilter
	// Filters returns FlowTaggedMetricsFilterFlowMetricTagFilterIterIter, set in FlowTaggedMetricsFilter
	Filters() FlowTaggedMetricsFilterFlowMetricTagFilterIter
	setNil()
}

// Controls inclusion/exclusion of tagged metrics when fetching flow metrics.
// Include returns a bool
func (obj *flowTaggedMetricsFilter) Include() bool {

	return *obj.obj.Include

}

// Controls inclusion/exclusion of tagged metrics when fetching flow metrics.
// Include returns a bool
func (obj *flowTaggedMetricsFilter) HasInclude() bool {
	return obj.obj.Include != nil
}

// Controls inclusion/exclusion of tagged metrics when fetching flow metrics.
// SetInclude sets the bool value in the FlowTaggedMetricsFilter object
func (obj *flowTaggedMetricsFilter) SetInclude(value bool) FlowTaggedMetricsFilter {

	obj.obj.Include = &value
	return obj
}

// Controls inclusion/exclusion of tagged metrics where each underlying attributes has zero value or absent value.
// IncludeEmptyMetrics returns a bool
func (obj *flowTaggedMetricsFilter) IncludeEmptyMetrics() bool {

	return *obj.obj.IncludeEmptyMetrics

}

// Controls inclusion/exclusion of tagged metrics where each underlying attributes has zero value or absent value.
// IncludeEmptyMetrics returns a bool
func (obj *flowTaggedMetricsFilter) HasIncludeEmptyMetrics() bool {
	return obj.obj.IncludeEmptyMetrics != nil
}

// Controls inclusion/exclusion of tagged metrics where each underlying attributes has zero value or absent value.
// SetIncludeEmptyMetrics sets the bool value in the FlowTaggedMetricsFilter object
func (obj *flowTaggedMetricsFilter) SetIncludeEmptyMetrics(value bool) FlowTaggedMetricsFilter {

	obj.obj.IncludeEmptyMetrics = &value
	return obj
}

type FlowTaggedMetricsFilterMetricNamesEnum string

// Enum of MetricNames on FlowTaggedMetricsFilter
var FlowTaggedMetricsFilterMetricNames = struct {
	FRAMES_TX      FlowTaggedMetricsFilterMetricNamesEnum
	FRAMES_RX      FlowTaggedMetricsFilterMetricNamesEnum
	BYTES_TX       FlowTaggedMetricsFilterMetricNamesEnum
	BYTES_RX       FlowTaggedMetricsFilterMetricNamesEnum
	FRAMES_TX_RATE FlowTaggedMetricsFilterMetricNamesEnum
	FRAMES_RX_RATE FlowTaggedMetricsFilterMetricNamesEnum
}{
	FRAMES_TX:      FlowTaggedMetricsFilterMetricNamesEnum("frames_tx"),
	FRAMES_RX:      FlowTaggedMetricsFilterMetricNamesEnum("frames_rx"),
	BYTES_TX:       FlowTaggedMetricsFilterMetricNamesEnum("bytes_tx"),
	BYTES_RX:       FlowTaggedMetricsFilterMetricNamesEnum("bytes_rx"),
	FRAMES_TX_RATE: FlowTaggedMetricsFilterMetricNamesEnum("frames_tx_rate"),
	FRAMES_RX_RATE: FlowTaggedMetricsFilterMetricNamesEnum("frames_rx_rate"),
}

func (obj *flowTaggedMetricsFilter) MetricNames() []FlowTaggedMetricsFilterMetricNamesEnum {
	items := []FlowTaggedMetricsFilterMetricNamesEnum{}
	for _, item := range obj.obj.MetricNames {
		items = append(items, FlowTaggedMetricsFilterMetricNamesEnum(item.String()))
	}
	return items
}

// The list of metric names that the returned result set will contain. If the list is empty then all metrics will be returned.
// SetMetricNames sets the []string value in the FlowTaggedMetricsFilter object
func (obj *flowTaggedMetricsFilter) SetMetricNames(value []FlowTaggedMetricsFilterMetricNamesEnum) FlowTaggedMetricsFilter {

	items := []otg.FlowTaggedMetricsFilter_MetricNames_Enum{}
	for _, item := range value {
		intValue := otg.FlowTaggedMetricsFilter_MetricNames_Enum_value[string(item)]
		items = append(items, otg.FlowTaggedMetricsFilter_MetricNames_Enum(intValue))
	}
	obj.obj.MetricNames = items
	return obj
}

// List of filters to selectively fetch tagged metrics with certain tag and corresponding value.
// Filters returns a []FlowMetricTagFilter
func (obj *flowTaggedMetricsFilter) Filters() FlowTaggedMetricsFilterFlowMetricTagFilterIter {
	if len(obj.obj.Filters) == 0 {
		obj.obj.Filters = []*otg.FlowMetricTagFilter{}
	}
	if obj.filtersHolder == nil {
		obj.filtersHolder = newFlowTaggedMetricsFilterFlowMetricTagFilterIter(&obj.obj.Filters).setMsg(obj)
	}
	return obj.filtersHolder
}

type flowTaggedMetricsFilterFlowMetricTagFilterIter struct {
	obj                      *flowTaggedMetricsFilter
	flowMetricTagFilterSlice []FlowMetricTagFilter
	fieldPtr                 *[]*otg.FlowMetricTagFilter
}

func newFlowTaggedMetricsFilterFlowMetricTagFilterIter(ptr *[]*otg.FlowMetricTagFilter) FlowTaggedMetricsFilterFlowMetricTagFilterIter {
	return &flowTaggedMetricsFilterFlowMetricTagFilterIter{fieldPtr: ptr}
}

type FlowTaggedMetricsFilterFlowMetricTagFilterIter interface {
	setMsg(*flowTaggedMetricsFilter) FlowTaggedMetricsFilterFlowMetricTagFilterIter
	Items() []FlowMetricTagFilter
	Add() FlowMetricTagFilter
	Append(items ...FlowMetricTagFilter) FlowTaggedMetricsFilterFlowMetricTagFilterIter
	Set(index int, newObj FlowMetricTagFilter) FlowTaggedMetricsFilterFlowMetricTagFilterIter
	Clear() FlowTaggedMetricsFilterFlowMetricTagFilterIter
	clearHolderSlice() FlowTaggedMetricsFilterFlowMetricTagFilterIter
	appendHolderSlice(item FlowMetricTagFilter) FlowTaggedMetricsFilterFlowMetricTagFilterIter
}

func (obj *flowTaggedMetricsFilterFlowMetricTagFilterIter) setMsg(msg *flowTaggedMetricsFilter) FlowTaggedMetricsFilterFlowMetricTagFilterIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowMetricTagFilter{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowTaggedMetricsFilterFlowMetricTagFilterIter) Items() []FlowMetricTagFilter {
	return obj.flowMetricTagFilterSlice
}

func (obj *flowTaggedMetricsFilterFlowMetricTagFilterIter) Add() FlowMetricTagFilter {
	newObj := &otg.FlowMetricTagFilter{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowMetricTagFilter{obj: newObj}
	newLibObj.setDefault()
	obj.flowMetricTagFilterSlice = append(obj.flowMetricTagFilterSlice, newLibObj)
	return newLibObj
}

func (obj *flowTaggedMetricsFilterFlowMetricTagFilterIter) Append(items ...FlowMetricTagFilter) FlowTaggedMetricsFilterFlowMetricTagFilterIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowMetricTagFilterSlice = append(obj.flowMetricTagFilterSlice, item)
	}
	return obj
}

func (obj *flowTaggedMetricsFilterFlowMetricTagFilterIter) Set(index int, newObj FlowMetricTagFilter) FlowTaggedMetricsFilterFlowMetricTagFilterIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowMetricTagFilterSlice[index] = newObj
	return obj
}
func (obj *flowTaggedMetricsFilterFlowMetricTagFilterIter) Clear() FlowTaggedMetricsFilterFlowMetricTagFilterIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowMetricTagFilter{}
		obj.flowMetricTagFilterSlice = []FlowMetricTagFilter{}
	}
	return obj
}
func (obj *flowTaggedMetricsFilterFlowMetricTagFilterIter) clearHolderSlice() FlowTaggedMetricsFilterFlowMetricTagFilterIter {
	if len(obj.flowMetricTagFilterSlice) > 0 {
		obj.flowMetricTagFilterSlice = []FlowMetricTagFilter{}
	}
	return obj
}
func (obj *flowTaggedMetricsFilterFlowMetricTagFilterIter) appendHolderSlice(item FlowMetricTagFilter) FlowTaggedMetricsFilterFlowMetricTagFilterIter {
	obj.flowMetricTagFilterSlice = append(obj.flowMetricTagFilterSlice, item)
	return obj
}

func (obj *flowTaggedMetricsFilter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Filters) != 0 {

		if set_default {
			obj.Filters().clearHolderSlice()
			for _, item := range obj.obj.Filters {
				obj.Filters().appendHolderSlice(&flowMetricTagFilter{obj: item})
			}
		}
		for _, item := range obj.Filters().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *flowTaggedMetricsFilter) setDefault() {
	if obj.obj.Include == nil {
		obj.SetInclude(true)
	}
	if obj.obj.IncludeEmptyMetrics == nil {
		obj.SetIncludeEmptyMetrics(false)
	}

}
