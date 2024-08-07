package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowMetrics *****
type flowMetrics struct {
	validation
	obj                        *otg.FlowMetrics
	marshaller                 marshalFlowMetrics
	unMarshaller               unMarshalFlowMetrics
	rxTxRatioHolder            FlowRxTxRatio
	latencyHolder              FlowLatencyMetrics
	predefinedMetricTagsHolder FlowPredefinedTags
}

func NewFlowMetrics() FlowMetrics {
	obj := flowMetrics{obj: &otg.FlowMetrics{}}
	obj.setDefault()
	return &obj
}

func (obj *flowMetrics) msg() *otg.FlowMetrics {
	return obj.obj
}

func (obj *flowMetrics) setMsg(msg *otg.FlowMetrics) FlowMetrics {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowMetrics struct {
	obj *flowMetrics
}

type marshalFlowMetrics interface {
	// ToProto marshals FlowMetrics to protobuf object *otg.FlowMetrics
	ToProto() (*otg.FlowMetrics, error)
	// ToPbText marshals FlowMetrics to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowMetrics to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowMetrics to JSON text
	ToJson() (string, error)
}

type unMarshalflowMetrics struct {
	obj *flowMetrics
}

type unMarshalFlowMetrics interface {
	// FromProto unmarshals FlowMetrics from protobuf object *otg.FlowMetrics
	FromProto(msg *otg.FlowMetrics) (FlowMetrics, error)
	// FromPbText unmarshals FlowMetrics from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowMetrics from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowMetrics from JSON text
	FromJson(value string) error
}

func (obj *flowMetrics) Marshal() marshalFlowMetrics {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowMetrics{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowMetrics) Unmarshal() unMarshalFlowMetrics {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowMetrics{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowMetrics) ToProto() (*otg.FlowMetrics, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowMetrics) FromProto(msg *otg.FlowMetrics) (FlowMetrics, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowMetrics) ToPbText() (string, error) {
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

func (m *unMarshalflowMetrics) FromPbText(value string) error {
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

func (m *marshalflowMetrics) ToYaml() (string, error) {
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

func (m *unMarshalflowMetrics) FromYaml(value string) error {
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

func (m *marshalflowMetrics) ToJson() (string, error) {
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

func (m *unMarshalflowMetrics) FromJson(value string) error {
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

func (obj *flowMetrics) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowMetrics) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowMetrics) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowMetrics) Clone() (FlowMetrics, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowMetrics()
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

func (obj *flowMetrics) setNil() {
	obj.rxTxRatioHolder = nil
	obj.latencyHolder = nil
	obj.predefinedMetricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowMetrics is the optional container for configuring flow metrics.
type FlowMetrics interface {
	Validation
	// msg marshals FlowMetrics to protobuf object *otg.FlowMetrics
	// and doesn't set defaults
	msg() *otg.FlowMetrics
	// setMsg unmarshals FlowMetrics from protobuf object *otg.FlowMetrics
	// and doesn't set defaults
	setMsg(*otg.FlowMetrics) FlowMetrics
	// provides marshal interface
	Marshal() marshalFlowMetrics
	// provides unmarshal interface
	Unmarshal() unMarshalFlowMetrics
	// validate validates FlowMetrics
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowMetrics, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Enable returns bool, set in FlowMetrics.
	Enable() bool
	// SetEnable assigns bool provided by user to FlowMetrics
	SetEnable(value bool) FlowMetrics
	// HasEnable checks if Enable has been set in FlowMetrics
	HasEnable() bool
	// Loss returns bool, set in FlowMetrics.
	Loss() bool
	// SetLoss assigns bool provided by user to FlowMetrics
	SetLoss(value bool) FlowMetrics
	// HasLoss checks if Loss has been set in FlowMetrics
	HasLoss() bool
	// RxTxRatio returns FlowRxTxRatio, set in FlowMetrics.
	// FlowRxTxRatio is rx Tx ratio is the ratio of expected number of Rx packets across all Rx ports to the Tx packets
	// for the configured flow. It is a factor by which the Tx packet count is multiplied to calculate
	// the sum of expected Rx packet count, across all Rx ports. This will be used to calculate loss
	// percentage of flow at aggregate level.
	RxTxRatio() FlowRxTxRatio
	// SetRxTxRatio assigns FlowRxTxRatio provided by user to FlowMetrics.
	// FlowRxTxRatio is rx Tx ratio is the ratio of expected number of Rx packets across all Rx ports to the Tx packets
	// for the configured flow. It is a factor by which the Tx packet count is multiplied to calculate
	// the sum of expected Rx packet count, across all Rx ports. This will be used to calculate loss
	// percentage of flow at aggregate level.
	SetRxTxRatio(value FlowRxTxRatio) FlowMetrics
	// HasRxTxRatio checks if RxTxRatio has been set in FlowMetrics
	HasRxTxRatio() bool
	// Timestamps returns bool, set in FlowMetrics.
	Timestamps() bool
	// SetTimestamps assigns bool provided by user to FlowMetrics
	SetTimestamps(value bool) FlowMetrics
	// HasTimestamps checks if Timestamps has been set in FlowMetrics
	HasTimestamps() bool
	// Latency returns FlowLatencyMetrics, set in FlowMetrics.
	// FlowLatencyMetrics is the optional container for per flow latency metric configuration.
	Latency() FlowLatencyMetrics
	// SetLatency assigns FlowLatencyMetrics provided by user to FlowMetrics.
	// FlowLatencyMetrics is the optional container for per flow latency metric configuration.
	SetLatency(value FlowLatencyMetrics) FlowMetrics
	// HasLatency checks if Latency has been set in FlowMetrics
	HasLatency() bool
	// PredefinedMetricTags returns FlowPredefinedTags, set in FlowMetrics.
	// FlowPredefinedTags is list of predefined flow tracking options, outside packet fields, that can be enabled.
	PredefinedMetricTags() FlowPredefinedTags
	// SetPredefinedMetricTags assigns FlowPredefinedTags provided by user to FlowMetrics.
	// FlowPredefinedTags is list of predefined flow tracking options, outside packet fields, that can be enabled.
	SetPredefinedMetricTags(value FlowPredefinedTags) FlowMetrics
	// HasPredefinedMetricTags checks if PredefinedMetricTags has been set in FlowMetrics
	HasPredefinedMetricTags() bool
	setNil()
}

// Enables flow metrics.
// Enabling this option may affect the resultant packet payload due to
// additional instrumentation data.
// Enable returns a bool
func (obj *flowMetrics) Enable() bool {

	return *obj.obj.Enable

}

// Enables flow metrics.
// Enabling this option may affect the resultant packet payload due to
// additional instrumentation data.
// Enable returns a bool
func (obj *flowMetrics) HasEnable() bool {
	return obj.obj.Enable != nil
}

// Enables flow metrics.
// Enabling this option may affect the resultant packet payload due to
// additional instrumentation data.
// SetEnable sets the bool value in the FlowMetrics object
func (obj *flowMetrics) SetEnable(value bool) FlowMetrics {

	obj.obj.Enable = &value
	return obj
}

// Enables additional flow metric loss calculation.
// Loss returns a bool
func (obj *flowMetrics) Loss() bool {

	return *obj.obj.Loss

}

// Enables additional flow metric loss calculation.
// Loss returns a bool
func (obj *flowMetrics) HasLoss() bool {
	return obj.obj.Loss != nil
}

// Enables additional flow metric loss calculation.
// SetLoss sets the bool value in the FlowMetrics object
func (obj *flowMetrics) SetLoss(value bool) FlowMetrics {

	obj.obj.Loss = &value
	return obj
}

// Rx Tx ratio.
// RxTxRatio returns a FlowRxTxRatio
func (obj *flowMetrics) RxTxRatio() FlowRxTxRatio {
	if obj.obj.RxTxRatio == nil {
		obj.obj.RxTxRatio = NewFlowRxTxRatio().msg()
	}
	if obj.rxTxRatioHolder == nil {
		obj.rxTxRatioHolder = &flowRxTxRatio{obj: obj.obj.RxTxRatio}
	}
	return obj.rxTxRatioHolder
}

// Rx Tx ratio.
// RxTxRatio returns a FlowRxTxRatio
func (obj *flowMetrics) HasRxTxRatio() bool {
	return obj.obj.RxTxRatio != nil
}

// Rx Tx ratio.
// SetRxTxRatio sets the FlowRxTxRatio value in the FlowMetrics object
func (obj *flowMetrics) SetRxTxRatio(value FlowRxTxRatio) FlowMetrics {

	obj.rxTxRatioHolder = nil
	obj.obj.RxTxRatio = value.msg()

	return obj
}

// Enables additional flow metric first and last timestamps.
// Timestamps returns a bool
func (obj *flowMetrics) Timestamps() bool {

	return *obj.obj.Timestamps

}

// Enables additional flow metric first and last timestamps.
// Timestamps returns a bool
func (obj *flowMetrics) HasTimestamps() bool {
	return obj.obj.Timestamps != nil
}

// Enables additional flow metric first and last timestamps.
// SetTimestamps sets the bool value in the FlowMetrics object
func (obj *flowMetrics) SetTimestamps(value bool) FlowMetrics {

	obj.obj.Timestamps = &value
	return obj
}

// Latency metrics.
// Latency returns a FlowLatencyMetrics
func (obj *flowMetrics) Latency() FlowLatencyMetrics {
	if obj.obj.Latency == nil {
		obj.obj.Latency = NewFlowLatencyMetrics().msg()
	}
	if obj.latencyHolder == nil {
		obj.latencyHolder = &flowLatencyMetrics{obj: obj.obj.Latency}
	}
	return obj.latencyHolder
}

// Latency metrics.
// Latency returns a FlowLatencyMetrics
func (obj *flowMetrics) HasLatency() bool {
	return obj.obj.Latency != nil
}

// Latency metrics.
// SetLatency sets the FlowLatencyMetrics value in the FlowMetrics object
func (obj *flowMetrics) SetLatency(value FlowLatencyMetrics) FlowMetrics {

	obj.latencyHolder = nil
	obj.obj.Latency = value.msg()

	return obj
}

// Predefined metric tags
// PredefinedMetricTags returns a FlowPredefinedTags
func (obj *flowMetrics) PredefinedMetricTags() FlowPredefinedTags {
	if obj.obj.PredefinedMetricTags == nil {
		obj.obj.PredefinedMetricTags = NewFlowPredefinedTags().msg()
	}
	if obj.predefinedMetricTagsHolder == nil {
		obj.predefinedMetricTagsHolder = &flowPredefinedTags{obj: obj.obj.PredefinedMetricTags}
	}
	return obj.predefinedMetricTagsHolder
}

// Predefined metric tags
// PredefinedMetricTags returns a FlowPredefinedTags
func (obj *flowMetrics) HasPredefinedMetricTags() bool {
	return obj.obj.PredefinedMetricTags != nil
}

// Predefined metric tags
// SetPredefinedMetricTags sets the FlowPredefinedTags value in the FlowMetrics object
func (obj *flowMetrics) SetPredefinedMetricTags(value FlowPredefinedTags) FlowMetrics {

	obj.predefinedMetricTagsHolder = nil
	obj.obj.PredefinedMetricTags = value.msg()

	return obj
}

func (obj *flowMetrics) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RxTxRatio != nil {

		obj.RxTxRatio().validateObj(vObj, set_default)
	}

	if obj.obj.Latency != nil {

		obj.Latency().validateObj(vObj, set_default)
	}

	if obj.obj.PredefinedMetricTags != nil {

		obj.PredefinedMetricTags().validateObj(vObj, set_default)
	}

}

func (obj *flowMetrics) setDefault() {
	if obj.obj.Enable == nil {
		obj.SetEnable(false)
	}
	if obj.obj.Loss == nil {
		obj.SetLoss(false)
	}
	if obj.obj.Timestamps == nil {
		obj.SetTimestamps(false)
	}

}
