package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MetricLatency *****
type metricLatency struct {
	validation
	obj          *otg.MetricLatency
	marshaller   marshalMetricLatency
	unMarshaller unMarshalMetricLatency
}

func NewMetricLatency() MetricLatency {
	obj := metricLatency{obj: &otg.MetricLatency{}}
	obj.setDefault()
	return &obj
}

func (obj *metricLatency) msg() *otg.MetricLatency {
	return obj.obj
}

func (obj *metricLatency) setMsg(msg *otg.MetricLatency) MetricLatency {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmetricLatency struct {
	obj *metricLatency
}

type marshalMetricLatency interface {
	// ToProto marshals MetricLatency to protobuf object *otg.MetricLatency
	ToProto() (*otg.MetricLatency, error)
	// ToPbText marshals MetricLatency to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MetricLatency to YAML text
	ToYaml() (string, error)
	// ToJson marshals MetricLatency to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals MetricLatency to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalmetricLatency struct {
	obj *metricLatency
}

type unMarshalMetricLatency interface {
	// FromProto unmarshals MetricLatency from protobuf object *otg.MetricLatency
	FromProto(msg *otg.MetricLatency) (MetricLatency, error)
	// FromPbText unmarshals MetricLatency from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MetricLatency from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MetricLatency from JSON text
	FromJson(value string) error
}

func (obj *metricLatency) Marshal() marshalMetricLatency {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmetricLatency{obj: obj}
	}
	return obj.marshaller
}

func (obj *metricLatency) Unmarshal() unMarshalMetricLatency {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmetricLatency{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmetricLatency) ToProto() (*otg.MetricLatency, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmetricLatency) FromProto(msg *otg.MetricLatency) (MetricLatency, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmetricLatency) ToPbText() (string, error) {
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

func (m *unMarshalmetricLatency) FromPbText(value string) error {
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

func (m *marshalmetricLatency) ToYaml() (string, error) {
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

func (m *unMarshalmetricLatency) FromYaml(value string) error {
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

func (m *marshalmetricLatency) ToJsonRaw() (string, error) {
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

func (m *marshalmetricLatency) ToJson() (string, error) {
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

func (m *unMarshalmetricLatency) FromJson(value string) error {
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

func (obj *metricLatency) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *metricLatency) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *metricLatency) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *metricLatency) Clone() (MetricLatency, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMetricLatency()
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

// MetricLatency is the container for latency metrics.
// The min/max/avg values are dependent on the type of latency measurement
// mode that is configured.
// The container will be empty if the latency has not been configured for
// the flow.
type MetricLatency interface {
	Validation
	// msg marshals MetricLatency to protobuf object *otg.MetricLatency
	// and doesn't set defaults
	msg() *otg.MetricLatency
	// setMsg unmarshals MetricLatency from protobuf object *otg.MetricLatency
	// and doesn't set defaults
	setMsg(*otg.MetricLatency) MetricLatency
	// provides marshal interface
	Marshal() marshalMetricLatency
	// provides unmarshal interface
	Unmarshal() unMarshalMetricLatency
	// validate validates MetricLatency
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MetricLatency, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// MinimumNs returns float64, set in MetricLatency.
	MinimumNs() float64
	// SetMinimumNs assigns float64 provided by user to MetricLatency
	SetMinimumNs(value float64) MetricLatency
	// HasMinimumNs checks if MinimumNs has been set in MetricLatency
	HasMinimumNs() bool
	// MaximumNs returns float64, set in MetricLatency.
	MaximumNs() float64
	// SetMaximumNs assigns float64 provided by user to MetricLatency
	SetMaximumNs(value float64) MetricLatency
	// HasMaximumNs checks if MaximumNs has been set in MetricLatency
	HasMaximumNs() bool
	// AverageNs returns float64, set in MetricLatency.
	AverageNs() float64
	// SetAverageNs assigns float64 provided by user to MetricLatency
	SetAverageNs(value float64) MetricLatency
	// HasAverageNs checks if AverageNs has been set in MetricLatency
	HasAverageNs() bool
}

// Minimum latency in nanoseconds
// MinimumNs returns a float64
func (obj *metricLatency) MinimumNs() float64 {

	return *obj.obj.MinimumNs

}

// Minimum latency in nanoseconds
// MinimumNs returns a float64
func (obj *metricLatency) HasMinimumNs() bool {
	return obj.obj.MinimumNs != nil
}

// Minimum latency in nanoseconds
// SetMinimumNs sets the float64 value in the MetricLatency object
func (obj *metricLatency) SetMinimumNs(value float64) MetricLatency {

	obj.obj.MinimumNs = &value
	return obj
}

// Maximum latency in nanoseconds
// MaximumNs returns a float64
func (obj *metricLatency) MaximumNs() float64 {

	return *obj.obj.MaximumNs

}

// Maximum latency in nanoseconds
// MaximumNs returns a float64
func (obj *metricLatency) HasMaximumNs() bool {
	return obj.obj.MaximumNs != nil
}

// Maximum latency in nanoseconds
// SetMaximumNs sets the float64 value in the MetricLatency object
func (obj *metricLatency) SetMaximumNs(value float64) MetricLatency {

	obj.obj.MaximumNs = &value
	return obj
}

// Average latency in nanoseconds
// AverageNs returns a float64
func (obj *metricLatency) AverageNs() float64 {

	return *obj.obj.AverageNs

}

// Average latency in nanoseconds
// AverageNs returns a float64
func (obj *metricLatency) HasAverageNs() bool {
	return obj.obj.AverageNs != nil
}

// Average latency in nanoseconds
// SetAverageNs sets the float64 value in the MetricLatency object
func (obj *metricLatency) SetAverageNs(value float64) MetricLatency {

	obj.obj.AverageNs = &value
	return obj
}

func (obj *metricLatency) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *metricLatency) setDefault() {

}
