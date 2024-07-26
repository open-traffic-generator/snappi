package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowLatencyMetrics *****
type flowLatencyMetrics struct {
	validation
	obj          *otg.FlowLatencyMetrics
	marshaller   marshalFlowLatencyMetrics
	unMarshaller unMarshalFlowLatencyMetrics
}

func NewFlowLatencyMetrics() FlowLatencyMetrics {
	obj := flowLatencyMetrics{obj: &otg.FlowLatencyMetrics{}}
	obj.setDefault()
	return &obj
}

func (obj *flowLatencyMetrics) msg() *otg.FlowLatencyMetrics {
	return obj.obj
}

func (obj *flowLatencyMetrics) setMsg(msg *otg.FlowLatencyMetrics) FlowLatencyMetrics {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowLatencyMetrics struct {
	obj *flowLatencyMetrics
}

type marshalFlowLatencyMetrics interface {
	// ToProto marshals FlowLatencyMetrics to protobuf object *otg.FlowLatencyMetrics
	ToProto() (*otg.FlowLatencyMetrics, error)
	// ToPbText marshals FlowLatencyMetrics to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowLatencyMetrics to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowLatencyMetrics to JSON text
	ToJson() (string, error)
}

type unMarshalflowLatencyMetrics struct {
	obj *flowLatencyMetrics
}

type unMarshalFlowLatencyMetrics interface {
	// FromProto unmarshals FlowLatencyMetrics from protobuf object *otg.FlowLatencyMetrics
	FromProto(msg *otg.FlowLatencyMetrics) (FlowLatencyMetrics, error)
	// FromPbText unmarshals FlowLatencyMetrics from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowLatencyMetrics from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowLatencyMetrics from JSON text
	FromJson(value string) error
}

func (obj *flowLatencyMetrics) Marshal() marshalFlowLatencyMetrics {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowLatencyMetrics{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowLatencyMetrics) Unmarshal() unMarshalFlowLatencyMetrics {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowLatencyMetrics{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowLatencyMetrics) ToProto() (*otg.FlowLatencyMetrics, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowLatencyMetrics) FromProto(msg *otg.FlowLatencyMetrics) (FlowLatencyMetrics, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowLatencyMetrics) ToPbText() (string, error) {
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

func (m *unMarshalflowLatencyMetrics) FromPbText(value string) error {
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

func (m *marshalflowLatencyMetrics) ToYaml() (string, error) {
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

func (m *unMarshalflowLatencyMetrics) FromYaml(value string) error {
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

func (m *marshalflowLatencyMetrics) ToJson() (string, error) {
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

func (m *unMarshalflowLatencyMetrics) FromJson(value string) error {
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

func (obj *flowLatencyMetrics) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowLatencyMetrics) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowLatencyMetrics) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowLatencyMetrics) Clone() (FlowLatencyMetrics, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowLatencyMetrics()
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

// FlowLatencyMetrics is the optional container for per flow latency metric configuration.
type FlowLatencyMetrics interface {
	Validation
	// msg marshals FlowLatencyMetrics to protobuf object *otg.FlowLatencyMetrics
	// and doesn't set defaults
	msg() *otg.FlowLatencyMetrics
	// setMsg unmarshals FlowLatencyMetrics from protobuf object *otg.FlowLatencyMetrics
	// and doesn't set defaults
	setMsg(*otg.FlowLatencyMetrics) FlowLatencyMetrics
	// provides marshal interface
	Marshal() marshalFlowLatencyMetrics
	// provides unmarshal interface
	Unmarshal() unMarshalFlowLatencyMetrics
	// validate validates FlowLatencyMetrics
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowLatencyMetrics, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Enable returns bool, set in FlowLatencyMetrics.
	Enable() bool
	// SetEnable assigns bool provided by user to FlowLatencyMetrics
	SetEnable(value bool) FlowLatencyMetrics
	// HasEnable checks if Enable has been set in FlowLatencyMetrics
	HasEnable() bool
	// Mode returns FlowLatencyMetricsModeEnum, set in FlowLatencyMetrics
	Mode() FlowLatencyMetricsModeEnum
	// SetMode assigns FlowLatencyMetricsModeEnum provided by user to FlowLatencyMetrics
	SetMode(value FlowLatencyMetricsModeEnum) FlowLatencyMetrics
	// HasMode checks if Mode has been set in FlowLatencyMetrics
	HasMode() bool
}

// True to enable latency metrics using timestamps.
//
// Enabling this option may affect the resultant packet payload due to
// additional instrumentation data.
// Enable returns a bool
func (obj *flowLatencyMetrics) Enable() bool {

	return *obj.obj.Enable

}

// True to enable latency metrics using timestamps.
//
// Enabling this option may affect the resultant packet payload due to
// additional instrumentation data.
// Enable returns a bool
func (obj *flowLatencyMetrics) HasEnable() bool {
	return obj.obj.Enable != nil
}

// True to enable latency metrics using timestamps.
//
// Enabling this option may affect the resultant packet payload due to
// additional instrumentation data.
// SetEnable sets the bool value in the FlowLatencyMetrics object
func (obj *flowLatencyMetrics) SetEnable(value bool) FlowLatencyMetrics {

	obj.obj.Enable = &value
	return obj
}

type FlowLatencyMetricsModeEnum string

// Enum of Mode on FlowLatencyMetrics
var FlowLatencyMetricsMode = struct {
	STORE_FORWARD FlowLatencyMetricsModeEnum
	CUT_THROUGH   FlowLatencyMetricsModeEnum
}{
	STORE_FORWARD: FlowLatencyMetricsModeEnum("store_forward"),
	CUT_THROUGH:   FlowLatencyMetricsModeEnum("cut_through"),
}

func (obj *flowLatencyMetrics) Mode() FlowLatencyMetricsModeEnum {
	return FlowLatencyMetricsModeEnum(obj.obj.Mode.Enum().String())
}

// Select the type of latency measurement. The different types of
// latency measurements are:
//
// store_forward:
// The time interval starting when the last bit of the frame leaves the
// sending port and ending when the first bit of the frame is seen on
// the receiving port (LIFO).  This is based on the RFC 1242 standard.
//
// cut_through:
// The time interval starting when the first bit of the frame leaves
// the sending port and ending when the first bit of the frame is seen
// on the receiving port (FIFO).  This is based on the RFC 1242
// standard.
// Mode returns a string
func (obj *flowLatencyMetrics) HasMode() bool {
	return obj.obj.Mode != nil
}

func (obj *flowLatencyMetrics) SetMode(value FlowLatencyMetricsModeEnum) FlowLatencyMetrics {
	intValue, ok := otg.FlowLatencyMetrics_Mode_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowLatencyMetricsModeEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowLatencyMetrics_Mode_Enum(intValue)
	obj.obj.Mode = &enumValue

	return obj
}

func (obj *flowLatencyMetrics) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *flowLatencyMetrics) setDefault() {
	if obj.obj.Enable == nil {
		obj.SetEnable(false)
	}
	if obj.obj.Mode == nil {
		obj.SetMode(FlowLatencyMetricsMode.STORE_FORWARD)

	}

}
