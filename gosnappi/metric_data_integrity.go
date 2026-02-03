package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MetricDataIntegrity *****
type metricDataIntegrity struct {
	validation
	obj          *otg.MetricDataIntegrity
	marshaller   marshalMetricDataIntegrity
	unMarshaller unMarshalMetricDataIntegrity
}

func NewMetricDataIntegrity() MetricDataIntegrity {
	obj := metricDataIntegrity{obj: &otg.MetricDataIntegrity{}}
	obj.setDefault()
	return &obj
}

func (obj *metricDataIntegrity) msg() *otg.MetricDataIntegrity {
	return obj.obj
}

func (obj *metricDataIntegrity) setMsg(msg *otg.MetricDataIntegrity) MetricDataIntegrity {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmetricDataIntegrity struct {
	obj *metricDataIntegrity
}

type marshalMetricDataIntegrity interface {
	// ToProto marshals MetricDataIntegrity to protobuf object *otg.MetricDataIntegrity
	ToProto() (*otg.MetricDataIntegrity, error)
	// ToPbText marshals MetricDataIntegrity to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MetricDataIntegrity to YAML text
	ToYaml() (string, error)
	// ToJson marshals MetricDataIntegrity to JSON text
	ToJson() (string, error)
}

type unMarshalmetricDataIntegrity struct {
	obj *metricDataIntegrity
}

type unMarshalMetricDataIntegrity interface {
	// FromProto unmarshals MetricDataIntegrity from protobuf object *otg.MetricDataIntegrity
	FromProto(msg *otg.MetricDataIntegrity) (MetricDataIntegrity, error)
	// FromPbText unmarshals MetricDataIntegrity from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MetricDataIntegrity from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MetricDataIntegrity from JSON text
	FromJson(value string) error
}

func (obj *metricDataIntegrity) Marshal() marshalMetricDataIntegrity {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmetricDataIntegrity{obj: obj}
	}
	return obj.marshaller
}

func (obj *metricDataIntegrity) Unmarshal() unMarshalMetricDataIntegrity {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmetricDataIntegrity{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmetricDataIntegrity) ToProto() (*otg.MetricDataIntegrity, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmetricDataIntegrity) FromProto(msg *otg.MetricDataIntegrity) (MetricDataIntegrity, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmetricDataIntegrity) ToPbText() (string, error) {
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

func (m *unMarshalmetricDataIntegrity) FromPbText(value string) error {
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

func (m *marshalmetricDataIntegrity) ToYaml() (string, error) {
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

func (m *unMarshalmetricDataIntegrity) FromYaml(value string) error {
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

func (m *marshalmetricDataIntegrity) ToJson() (string, error) {
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

func (m *unMarshalmetricDataIntegrity) FromJson(value string) error {
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

func (obj *metricDataIntegrity) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *metricDataIntegrity) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *metricDataIntegrity) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *metricDataIntegrity) Clone() (MetricDataIntegrity, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMetricDataIntegrity()
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

// MetricDataIntegrity is the container for data integrity metrics.
// The container will be empty if the global option for data integrity check has not been enabled.
type MetricDataIntegrity interface {
	Validation
	// msg marshals MetricDataIntegrity to protobuf object *otg.MetricDataIntegrity
	// and doesn't set defaults
	msg() *otg.MetricDataIntegrity
	// setMsg unmarshals MetricDataIntegrity from protobuf object *otg.MetricDataIntegrity
	// and doesn't set defaults
	setMsg(*otg.MetricDataIntegrity) MetricDataIntegrity
	// provides marshal interface
	Marshal() marshalMetricDataIntegrity
	// provides unmarshal interface
	Unmarshal() unMarshalMetricDataIntegrity
	// validate validates MetricDataIntegrity
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MetricDataIntegrity, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// FramesRx returns uint64, set in MetricDataIntegrity.
	FramesRx() uint64
	// SetFramesRx assigns uint64 provided by user to MetricDataIntegrity
	SetFramesRx(value uint64) MetricDataIntegrity
	// HasFramesRx checks if FramesRx has been set in MetricDataIntegrity
	HasFramesRx() bool
	// Errors returns uint64, set in MetricDataIntegrity.
	Errors() uint64
	// SetErrors assigns uint64 provided by user to MetricDataIntegrity
	SetErrors(value uint64) MetricDataIntegrity
	// HasErrors checks if Errors has been set in MetricDataIntegrity
	HasErrors() bool
}

// The current total number of valid frames received with payload intact
// FramesRx returns a uint64
func (obj *metricDataIntegrity) FramesRx() uint64 {

	return *obj.obj.FramesRx

}

// The current total number of valid frames received with payload intact
// FramesRx returns a uint64
func (obj *metricDataIntegrity) HasFramesRx() bool {
	return obj.obj.FramesRx != nil
}

// The current total number of valid frames received with payload intact
// SetFramesRx sets the uint64 value in the MetricDataIntegrity object
func (obj *metricDataIntegrity) SetFramesRx(value uint64) MetricDataIntegrity {

	obj.obj.FramesRx = &value
	return obj
}

// The current total number of valid frames received with payload modified
// Errors returns a uint64
func (obj *metricDataIntegrity) Errors() uint64 {

	return *obj.obj.Errors

}

// The current total number of valid frames received with payload modified
// Errors returns a uint64
func (obj *metricDataIntegrity) HasErrors() bool {
	return obj.obj.Errors != nil
}

// The current total number of valid frames received with payload modified
// SetErrors sets the uint64 value in the MetricDataIntegrity object
func (obj *metricDataIntegrity) SetErrors(value uint64) MetricDataIntegrity {

	obj.obj.Errors = &value
	return obj
}

func (obj *metricDataIntegrity) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *metricDataIntegrity) setDefault() {

}
