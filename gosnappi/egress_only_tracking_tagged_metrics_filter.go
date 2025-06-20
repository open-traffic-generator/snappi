package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** EgressOnlyTrackingTaggedMetricsFilter *****
type egressOnlyTrackingTaggedMetricsFilter struct {
	validation
	obj          *otg.EgressOnlyTrackingTaggedMetricsFilter
	marshaller   marshalEgressOnlyTrackingTaggedMetricsFilter
	unMarshaller unMarshalEgressOnlyTrackingTaggedMetricsFilter
}

func NewEgressOnlyTrackingTaggedMetricsFilter() EgressOnlyTrackingTaggedMetricsFilter {
	obj := egressOnlyTrackingTaggedMetricsFilter{obj: &otg.EgressOnlyTrackingTaggedMetricsFilter{}}
	obj.setDefault()
	return &obj
}

func (obj *egressOnlyTrackingTaggedMetricsFilter) msg() *otg.EgressOnlyTrackingTaggedMetricsFilter {
	return obj.obj
}

func (obj *egressOnlyTrackingTaggedMetricsFilter) setMsg(msg *otg.EgressOnlyTrackingTaggedMetricsFilter) EgressOnlyTrackingTaggedMetricsFilter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalegressOnlyTrackingTaggedMetricsFilter struct {
	obj *egressOnlyTrackingTaggedMetricsFilter
}

type marshalEgressOnlyTrackingTaggedMetricsFilter interface {
	// ToProto marshals EgressOnlyTrackingTaggedMetricsFilter to protobuf object *otg.EgressOnlyTrackingTaggedMetricsFilter
	ToProto() (*otg.EgressOnlyTrackingTaggedMetricsFilter, error)
	// ToPbText marshals EgressOnlyTrackingTaggedMetricsFilter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals EgressOnlyTrackingTaggedMetricsFilter to YAML text
	ToYaml() (string, error)
	// ToJson marshals EgressOnlyTrackingTaggedMetricsFilter to JSON text
	ToJson() (string, error)
}

type unMarshalegressOnlyTrackingTaggedMetricsFilter struct {
	obj *egressOnlyTrackingTaggedMetricsFilter
}

type unMarshalEgressOnlyTrackingTaggedMetricsFilter interface {
	// FromProto unmarshals EgressOnlyTrackingTaggedMetricsFilter from protobuf object *otg.EgressOnlyTrackingTaggedMetricsFilter
	FromProto(msg *otg.EgressOnlyTrackingTaggedMetricsFilter) (EgressOnlyTrackingTaggedMetricsFilter, error)
	// FromPbText unmarshals EgressOnlyTrackingTaggedMetricsFilter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals EgressOnlyTrackingTaggedMetricsFilter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals EgressOnlyTrackingTaggedMetricsFilter from JSON text
	FromJson(value string) error
}

func (obj *egressOnlyTrackingTaggedMetricsFilter) Marshal() marshalEgressOnlyTrackingTaggedMetricsFilter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalegressOnlyTrackingTaggedMetricsFilter{obj: obj}
	}
	return obj.marshaller
}

func (obj *egressOnlyTrackingTaggedMetricsFilter) Unmarshal() unMarshalEgressOnlyTrackingTaggedMetricsFilter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalegressOnlyTrackingTaggedMetricsFilter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalegressOnlyTrackingTaggedMetricsFilter) ToProto() (*otg.EgressOnlyTrackingTaggedMetricsFilter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalegressOnlyTrackingTaggedMetricsFilter) FromProto(msg *otg.EgressOnlyTrackingTaggedMetricsFilter) (EgressOnlyTrackingTaggedMetricsFilter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalegressOnlyTrackingTaggedMetricsFilter) ToPbText() (string, error) {
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

func (m *unMarshalegressOnlyTrackingTaggedMetricsFilter) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalegressOnlyTrackingTaggedMetricsFilter) ToYaml() (string, error) {
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

func (m *unMarshalegressOnlyTrackingTaggedMetricsFilter) FromYaml(value string) error {
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

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalegressOnlyTrackingTaggedMetricsFilter) ToJson() (string, error) {
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

func (m *unMarshalegressOnlyTrackingTaggedMetricsFilter) FromJson(value string) error {
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

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *egressOnlyTrackingTaggedMetricsFilter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *egressOnlyTrackingTaggedMetricsFilter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *egressOnlyTrackingTaggedMetricsFilter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *egressOnlyTrackingTaggedMetricsFilter) Clone() (EgressOnlyTrackingTaggedMetricsFilter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewEgressOnlyTrackingTaggedMetricsFilter()
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

// EgressOnlyTrackingTaggedMetricsFilter is filter for tagged metrics
type EgressOnlyTrackingTaggedMetricsFilter interface {
	Validation
	// msg marshals EgressOnlyTrackingTaggedMetricsFilter to protobuf object *otg.EgressOnlyTrackingTaggedMetricsFilter
	// and doesn't set defaults
	msg() *otg.EgressOnlyTrackingTaggedMetricsFilter
	// setMsg unmarshals EgressOnlyTrackingTaggedMetricsFilter from protobuf object *otg.EgressOnlyTrackingTaggedMetricsFilter
	// and doesn't set defaults
	setMsg(*otg.EgressOnlyTrackingTaggedMetricsFilter) EgressOnlyTrackingTaggedMetricsFilter
	// provides marshal interface
	Marshal() marshalEgressOnlyTrackingTaggedMetricsFilter
	// provides unmarshal interface
	Unmarshal() unMarshalEgressOnlyTrackingTaggedMetricsFilter
	// validate validates EgressOnlyTrackingTaggedMetricsFilter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (EgressOnlyTrackingTaggedMetricsFilter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// IncludeEmptyMetrics returns bool, set in EgressOnlyTrackingTaggedMetricsFilter.
	IncludeEmptyMetrics() bool
	// SetIncludeEmptyMetrics assigns bool provided by user to EgressOnlyTrackingTaggedMetricsFilter
	SetIncludeEmptyMetrics(value bool) EgressOnlyTrackingTaggedMetricsFilter
	// HasIncludeEmptyMetrics checks if IncludeEmptyMetrics has been set in EgressOnlyTrackingTaggedMetricsFilter
	HasIncludeEmptyMetrics() bool
	// MetricNames returns []EgressOnlyTrackingTaggedMetricsFilterMetricNamesEnum, set in EgressOnlyTrackingTaggedMetricsFilter
	MetricNames() []EgressOnlyTrackingTaggedMetricsFilterMetricNamesEnum
	// SetMetricNames assigns []EgressOnlyTrackingTaggedMetricsFilterMetricNamesEnum provided by user to EgressOnlyTrackingTaggedMetricsFilter
	SetMetricNames(value []EgressOnlyTrackingTaggedMetricsFilterMetricNamesEnum) EgressOnlyTrackingTaggedMetricsFilter
}

// Controls inclusion/exclusion of tagged metrics where each underlying attribute has zero value or absent value.
// IncludeEmptyMetrics returns a bool
func (obj *egressOnlyTrackingTaggedMetricsFilter) IncludeEmptyMetrics() bool {

	return *obj.obj.IncludeEmptyMetrics

}

// Controls inclusion/exclusion of tagged metrics where each underlying attribute has zero value or absent value.
// IncludeEmptyMetrics returns a bool
func (obj *egressOnlyTrackingTaggedMetricsFilter) HasIncludeEmptyMetrics() bool {
	return obj.obj.IncludeEmptyMetrics != nil
}

// Controls inclusion/exclusion of tagged metrics where each underlying attribute has zero value or absent value.
// SetIncludeEmptyMetrics sets the bool value in the EgressOnlyTrackingTaggedMetricsFilter object
func (obj *egressOnlyTrackingTaggedMetricsFilter) SetIncludeEmptyMetrics(value bool) EgressOnlyTrackingTaggedMetricsFilter {

	obj.obj.IncludeEmptyMetrics = &value
	return obj
}

type EgressOnlyTrackingTaggedMetricsFilterMetricNamesEnum string

// Enum of MetricNames on EgressOnlyTrackingTaggedMetricsFilter
var EgressOnlyTrackingTaggedMetricsFilterMetricNames = struct {
	FRAMES_RX      EgressOnlyTrackingTaggedMetricsFilterMetricNamesEnum
	BYTES_RX       EgressOnlyTrackingTaggedMetricsFilterMetricNamesEnum
	FRAMES_RX_RATE EgressOnlyTrackingTaggedMetricsFilterMetricNamesEnum
	RX_L1_RATE_BPS EgressOnlyTrackingTaggedMetricsFilterMetricNamesEnum
	RX_RATE_BYTES  EgressOnlyTrackingTaggedMetricsFilterMetricNamesEnum
	RX_RATE_BPS    EgressOnlyTrackingTaggedMetricsFilterMetricNamesEnum
	RX_RATE_KBPS   EgressOnlyTrackingTaggedMetricsFilterMetricNamesEnum
	RX_RATE_MBPS   EgressOnlyTrackingTaggedMetricsFilterMetricNamesEnum
	TX_METRICS     EgressOnlyTrackingTaggedMetricsFilterMetricNamesEnum
}{
	FRAMES_RX:      EgressOnlyTrackingTaggedMetricsFilterMetricNamesEnum("frames_rx"),
	BYTES_RX:       EgressOnlyTrackingTaggedMetricsFilterMetricNamesEnum("bytes_rx"),
	FRAMES_RX_RATE: EgressOnlyTrackingTaggedMetricsFilterMetricNamesEnum("frames_rx_rate"),
	RX_L1_RATE_BPS: EgressOnlyTrackingTaggedMetricsFilterMetricNamesEnum("rx_l1_rate_bps"),
	RX_RATE_BYTES:  EgressOnlyTrackingTaggedMetricsFilterMetricNamesEnum("rx_rate_bytes"),
	RX_RATE_BPS:    EgressOnlyTrackingTaggedMetricsFilterMetricNamesEnum("rx_rate_bps"),
	RX_RATE_KBPS:   EgressOnlyTrackingTaggedMetricsFilterMetricNamesEnum("rx_rate_kbps"),
	RX_RATE_MBPS:   EgressOnlyTrackingTaggedMetricsFilterMetricNamesEnum("rx_rate_mbps"),
	TX_METRICS:     EgressOnlyTrackingTaggedMetricsFilterMetricNamesEnum("tx_metrics"),
}

func (obj *egressOnlyTrackingTaggedMetricsFilter) MetricNames() []EgressOnlyTrackingTaggedMetricsFilterMetricNamesEnum {
	items := []EgressOnlyTrackingTaggedMetricsFilterMetricNamesEnum{}
	for _, item := range obj.obj.MetricNames {
		items = append(items, EgressOnlyTrackingTaggedMetricsFilterMetricNamesEnum(item.String()))
	}
	return items
}

// The list of metric names that the returned result set will contain. If the list is empty then all metrics will be returned. Note: tx_metrics is optional, it is applicable where implementation is able to retrieve transmitter information. In order to get Tx metrics, tx_metric must be added in metric_names and all supported Tx metrics will be returned as listed in metric response.
// SetMetricNames sets the []string value in the EgressOnlyTrackingTaggedMetricsFilter object
func (obj *egressOnlyTrackingTaggedMetricsFilter) SetMetricNames(value []EgressOnlyTrackingTaggedMetricsFilterMetricNamesEnum) EgressOnlyTrackingTaggedMetricsFilter {

	items := []otg.EgressOnlyTrackingTaggedMetricsFilter_MetricNames_Enum{}
	for _, item := range value {
		intValue := otg.EgressOnlyTrackingTaggedMetricsFilter_MetricNames_Enum_value[string(item)]
		items = append(items, otg.EgressOnlyTrackingTaggedMetricsFilter_MetricNames_Enum(intValue))
	}
	obj.obj.MetricNames = items
	return obj
}

func (obj *egressOnlyTrackingTaggedMetricsFilter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *egressOnlyTrackingTaggedMetricsFilter) setDefault() {
	if obj.obj.IncludeEmptyMetrics == nil {
		obj.SetIncludeEmptyMetrics(false)
	}

}
