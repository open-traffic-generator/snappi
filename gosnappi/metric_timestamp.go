package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MetricTimestamp *****
type metricTimestamp struct {
	validation
	obj          *otg.MetricTimestamp
	marshaller   marshalMetricTimestamp
	unMarshaller unMarshalMetricTimestamp
}

func NewMetricTimestamp() MetricTimestamp {
	obj := metricTimestamp{obj: &otg.MetricTimestamp{}}
	obj.setDefault()
	return &obj
}

func (obj *metricTimestamp) msg() *otg.MetricTimestamp {
	return obj.obj
}

func (obj *metricTimestamp) setMsg(msg *otg.MetricTimestamp) MetricTimestamp {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmetricTimestamp struct {
	obj *metricTimestamp
}

type marshalMetricTimestamp interface {
	// ToProto marshals MetricTimestamp to protobuf object *otg.MetricTimestamp
	ToProto() (*otg.MetricTimestamp, error)
	// ToPbText marshals MetricTimestamp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MetricTimestamp to YAML text
	ToYaml() (string, error)
	// ToJson marshals MetricTimestamp to JSON text
	ToJson() (string, error)
}

type unMarshalmetricTimestamp struct {
	obj *metricTimestamp
}

type unMarshalMetricTimestamp interface {
	// FromProto unmarshals MetricTimestamp from protobuf object *otg.MetricTimestamp
	FromProto(msg *otg.MetricTimestamp) (MetricTimestamp, error)
	// FromPbText unmarshals MetricTimestamp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MetricTimestamp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MetricTimestamp from JSON text
	FromJson(value string) error
}

func (obj *metricTimestamp) Marshal() marshalMetricTimestamp {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmetricTimestamp{obj: obj}
	}
	return obj.marshaller
}

func (obj *metricTimestamp) Unmarshal() unMarshalMetricTimestamp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmetricTimestamp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmetricTimestamp) ToProto() (*otg.MetricTimestamp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmetricTimestamp) FromProto(msg *otg.MetricTimestamp) (MetricTimestamp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmetricTimestamp) ToPbText() (string, error) {
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

func (m *unMarshalmetricTimestamp) FromPbText(value string) error {
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

func (m *marshalmetricTimestamp) ToYaml() (string, error) {
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

func (m *unMarshalmetricTimestamp) FromYaml(value string) error {
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

func (m *marshalmetricTimestamp) ToJson() (string, error) {
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

func (m *unMarshalmetricTimestamp) FromJson(value string) error {
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

func (obj *metricTimestamp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *metricTimestamp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *metricTimestamp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *metricTimestamp) Clone() (MetricTimestamp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMetricTimestamp()
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

// MetricTimestamp is the container for timestamp metrics.
// The container will be empty if the timestamp has not been configured for
// the flow.
type MetricTimestamp interface {
	Validation
	// msg marshals MetricTimestamp to protobuf object *otg.MetricTimestamp
	// and doesn't set defaults
	msg() *otg.MetricTimestamp
	// setMsg unmarshals MetricTimestamp from protobuf object *otg.MetricTimestamp
	// and doesn't set defaults
	setMsg(*otg.MetricTimestamp) MetricTimestamp
	// provides marshal interface
	Marshal() marshalMetricTimestamp
	// provides unmarshal interface
	Unmarshal() unMarshalMetricTimestamp
	// validate validates MetricTimestamp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MetricTimestamp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// FirstTimestampNs returns float64, set in MetricTimestamp.
	FirstTimestampNs() float64
	// SetFirstTimestampNs assigns float64 provided by user to MetricTimestamp
	SetFirstTimestampNs(value float64) MetricTimestamp
	// HasFirstTimestampNs checks if FirstTimestampNs has been set in MetricTimestamp
	HasFirstTimestampNs() bool
	// LastTimestampNs returns float64, set in MetricTimestamp.
	LastTimestampNs() float64
	// SetLastTimestampNs assigns float64 provided by user to MetricTimestamp
	SetLastTimestampNs(value float64) MetricTimestamp
	// HasLastTimestampNs checks if LastTimestampNs has been set in MetricTimestamp
	HasLastTimestampNs() bool
}

// First timestamp in nanoseconds
// FirstTimestampNs returns a float64
func (obj *metricTimestamp) FirstTimestampNs() float64 {

	return *obj.obj.FirstTimestampNs

}

// First timestamp in nanoseconds
// FirstTimestampNs returns a float64
func (obj *metricTimestamp) HasFirstTimestampNs() bool {
	return obj.obj.FirstTimestampNs != nil
}

// First timestamp in nanoseconds
// SetFirstTimestampNs sets the float64 value in the MetricTimestamp object
func (obj *metricTimestamp) SetFirstTimestampNs(value float64) MetricTimestamp {

	obj.obj.FirstTimestampNs = &value
	return obj
}

// Last timestamp in nanoseconds
// LastTimestampNs returns a float64
func (obj *metricTimestamp) LastTimestampNs() float64 {

	return *obj.obj.LastTimestampNs

}

// Last timestamp in nanoseconds
// LastTimestampNs returns a float64
func (obj *metricTimestamp) HasLastTimestampNs() bool {
	return obj.obj.LastTimestampNs != nil
}

// Last timestamp in nanoseconds
// SetLastTimestampNs sets the float64 value in the MetricTimestamp object
func (obj *metricTimestamp) SetLastTimestampNs(value float64) MetricTimestamp {

	obj.obj.LastTimestampNs = &value
	return obj
}

func (obj *metricTimestamp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *metricTimestamp) setDefault() {

}
