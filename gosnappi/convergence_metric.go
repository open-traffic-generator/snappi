package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ConvergenceMetric *****
type convergenceMetric struct {
	validation
	obj          *otg.ConvergenceMetric
	marshaller   marshalConvergenceMetric
	unMarshaller unMarshalConvergenceMetric
	eventsHolder ConvergenceMetricConvergenceEventIter
}

func NewConvergenceMetric() ConvergenceMetric {
	obj := convergenceMetric{obj: &otg.ConvergenceMetric{}}
	obj.setDefault()
	return &obj
}

func (obj *convergenceMetric) msg() *otg.ConvergenceMetric {
	return obj.obj
}

func (obj *convergenceMetric) setMsg(msg *otg.ConvergenceMetric) ConvergenceMetric {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalconvergenceMetric struct {
	obj *convergenceMetric
}

type marshalConvergenceMetric interface {
	// ToProto marshals ConvergenceMetric to protobuf object *otg.ConvergenceMetric
	ToProto() (*otg.ConvergenceMetric, error)
	// ToPbText marshals ConvergenceMetric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ConvergenceMetric to YAML text
	ToYaml() (string, error)
	// ToJson marshals ConvergenceMetric to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals ConvergenceMetric to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalconvergenceMetric struct {
	obj *convergenceMetric
}

type unMarshalConvergenceMetric interface {
	// FromProto unmarshals ConvergenceMetric from protobuf object *otg.ConvergenceMetric
	FromProto(msg *otg.ConvergenceMetric) (ConvergenceMetric, error)
	// FromPbText unmarshals ConvergenceMetric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ConvergenceMetric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ConvergenceMetric from JSON text
	FromJson(value string) error
}

func (obj *convergenceMetric) Marshal() marshalConvergenceMetric {
	if obj.marshaller == nil {
		obj.marshaller = &marshalconvergenceMetric{obj: obj}
	}
	return obj.marshaller
}

func (obj *convergenceMetric) Unmarshal() unMarshalConvergenceMetric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalconvergenceMetric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalconvergenceMetric) ToProto() (*otg.ConvergenceMetric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalconvergenceMetric) FromProto(msg *otg.ConvergenceMetric) (ConvergenceMetric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalconvergenceMetric) ToPbText() (string, error) {
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

func (m *unMarshalconvergenceMetric) FromPbText(value string) error {
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

func (m *marshalconvergenceMetric) ToYaml() (string, error) {
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

func (m *unMarshalconvergenceMetric) FromYaml(value string) error {
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

func (m *marshalconvergenceMetric) ToJsonRaw() (string, error) {
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

func (m *marshalconvergenceMetric) ToJson() (string, error) {
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

func (m *unMarshalconvergenceMetric) FromJson(value string) error {
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

func (obj *convergenceMetric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *convergenceMetric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *convergenceMetric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *convergenceMetric) Clone() (ConvergenceMetric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewConvergenceMetric()
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

func (obj *convergenceMetric) setNil() {
	obj.eventsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ConvergenceMetric is under Review: Convergence metrics is currently under review for pending exploration on use cases.
//
// Under Review: Convergence metrics is currently under review for pending exploration on use cases.
//
// The container for convergence metrics.
type ConvergenceMetric interface {
	Validation
	// msg marshals ConvergenceMetric to protobuf object *otg.ConvergenceMetric
	// and doesn't set defaults
	msg() *otg.ConvergenceMetric
	// setMsg unmarshals ConvergenceMetric from protobuf object *otg.ConvergenceMetric
	// and doesn't set defaults
	setMsg(*otg.ConvergenceMetric) ConvergenceMetric
	// provides marshal interface
	Marshal() marshalConvergenceMetric
	// provides unmarshal interface
	Unmarshal() unMarshalConvergenceMetric
	// validate validates ConvergenceMetric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ConvergenceMetric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in ConvergenceMetric.
	Name() string
	// SetName assigns string provided by user to ConvergenceMetric
	SetName(value string) ConvergenceMetric
	// HasName checks if Name has been set in ConvergenceMetric
	HasName() bool
	// DataPlaneConvergenceUs returns float64, set in ConvergenceMetric.
	DataPlaneConvergenceUs() float64
	// SetDataPlaneConvergenceUs assigns float64 provided by user to ConvergenceMetric
	SetDataPlaneConvergenceUs(value float64) ConvergenceMetric
	// HasDataPlaneConvergenceUs checks if DataPlaneConvergenceUs has been set in ConvergenceMetric
	HasDataPlaneConvergenceUs() bool
	// ControlPlaneDataPlaneConvergenceUs returns float64, set in ConvergenceMetric.
	ControlPlaneDataPlaneConvergenceUs() float64
	// SetControlPlaneDataPlaneConvergenceUs assigns float64 provided by user to ConvergenceMetric
	SetControlPlaneDataPlaneConvergenceUs(value float64) ConvergenceMetric
	// HasControlPlaneDataPlaneConvergenceUs checks if ControlPlaneDataPlaneConvergenceUs has been set in ConvergenceMetric
	HasControlPlaneDataPlaneConvergenceUs() bool
	// Events returns ConvergenceMetricConvergenceEventIterIter, set in ConvergenceMetric
	Events() ConvergenceMetricConvergenceEventIter
	setNil()
}

// The name of a flow.
// Name returns a string
func (obj *convergenceMetric) Name() string {

	return *obj.obj.Name

}

// The name of a flow.
// Name returns a string
func (obj *convergenceMetric) HasName() bool {
	return obj.obj.Name != nil
}

// The name of a flow.
// SetName sets the string value in the ConvergenceMetric object
func (obj *convergenceMetric) SetName(value string) ConvergenceMetric {

	obj.obj.Name = &value
	return obj
}

// The convergence time(microseconds) measured from the data plane perspective only.
// It measures the time w.r.t. last start of the traffic of the affected flow from Below Threshold Timestamp,
// when the rate on Test Port 2 crosses below the Rx Threshold until an acceptable amount of traffic was
// received at time Above Threshold Timestamp, when the rate crosses above the configured
// rx_rate_threshold.
// DataPlaneConvergenceUs returns a float64
func (obj *convergenceMetric) DataPlaneConvergenceUs() float64 {

	return *obj.obj.DataPlaneConvergenceUs

}

// The convergence time(microseconds) measured from the data plane perspective only.
// It measures the time w.r.t. last start of the traffic of the affected flow from Below Threshold Timestamp,
// when the rate on Test Port 2 crosses below the Rx Threshold until an acceptable amount of traffic was
// received at time Above Threshold Timestamp, when the rate crosses above the configured
// rx_rate_threshold.
// DataPlaneConvergenceUs returns a float64
func (obj *convergenceMetric) HasDataPlaneConvergenceUs() bool {
	return obj.obj.DataPlaneConvergenceUs != nil
}

// The convergence time(microseconds) measured from the data plane perspective only.
// It measures the time w.r.t. last start of the traffic of the affected flow from Below Threshold Timestamp,
// when the rate on Test Port 2 crosses below the Rx Threshold until an acceptable amount of traffic was
// received at time Above Threshold Timestamp, when the rate crosses above the configured
// rx_rate_threshold.
// SetDataPlaneConvergenceUs sets the float64 value in the ConvergenceMetric object
func (obj *convergenceMetric) SetDataPlaneConvergenceUs(value float64) ConvergenceMetric {

	obj.obj.DataPlaneConvergenceUs = &value
	return obj
}

// The total convergence time(microseconds), between the event that caused the
// switchover until an acceptable amount of traffic was
// received at time Above Threshold Timestamp, when the rate crosses above the configured
// rx_rate_threshold.
// ControlPlaneDataPlaneConvergenceUs returns a float64
func (obj *convergenceMetric) ControlPlaneDataPlaneConvergenceUs() float64 {

	return *obj.obj.ControlPlaneDataPlaneConvergenceUs

}

// The total convergence time(microseconds), between the event that caused the
// switchover until an acceptable amount of traffic was
// received at time Above Threshold Timestamp, when the rate crosses above the configured
// rx_rate_threshold.
// ControlPlaneDataPlaneConvergenceUs returns a float64
func (obj *convergenceMetric) HasControlPlaneDataPlaneConvergenceUs() bool {
	return obj.obj.ControlPlaneDataPlaneConvergenceUs != nil
}

// The total convergence time(microseconds), between the event that caused the
// switchover until an acceptable amount of traffic was
// received at time Above Threshold Timestamp, when the rate crosses above the configured
// rx_rate_threshold.
// SetControlPlaneDataPlaneConvergenceUs sets the float64 value in the ConvergenceMetric object
func (obj *convergenceMetric) SetControlPlaneDataPlaneConvergenceUs(value float64) ConvergenceMetric {

	obj.obj.ControlPlaneDataPlaneConvergenceUs = &value
	return obj
}

// The events that were used to determine the convergence analytics.
// Events returns a []ConvergenceEvent
func (obj *convergenceMetric) Events() ConvergenceMetricConvergenceEventIter {
	if len(obj.obj.Events) == 0 {
		obj.obj.Events = []*otg.ConvergenceEvent{}
	}
	if obj.eventsHolder == nil {
		obj.eventsHolder = newConvergenceMetricConvergenceEventIter(&obj.obj.Events).setMsg(obj)
	}
	return obj.eventsHolder
}

type convergenceMetricConvergenceEventIter struct {
	obj                   *convergenceMetric
	convergenceEventSlice []ConvergenceEvent
	fieldPtr              *[]*otg.ConvergenceEvent
}

func newConvergenceMetricConvergenceEventIter(ptr *[]*otg.ConvergenceEvent) ConvergenceMetricConvergenceEventIter {
	return &convergenceMetricConvergenceEventIter{fieldPtr: ptr}
}

type ConvergenceMetricConvergenceEventIter interface {
	setMsg(*convergenceMetric) ConvergenceMetricConvergenceEventIter
	Items() []ConvergenceEvent
	Add() ConvergenceEvent
	Append(items ...ConvergenceEvent) ConvergenceMetricConvergenceEventIter
	Set(index int, newObj ConvergenceEvent) ConvergenceMetricConvergenceEventIter
	Clear() ConvergenceMetricConvergenceEventIter
	clearHolderSlice() ConvergenceMetricConvergenceEventIter
	appendHolderSlice(item ConvergenceEvent) ConvergenceMetricConvergenceEventIter
}

func (obj *convergenceMetricConvergenceEventIter) setMsg(msg *convergenceMetric) ConvergenceMetricConvergenceEventIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&convergenceEvent{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *convergenceMetricConvergenceEventIter) Items() []ConvergenceEvent {
	return obj.convergenceEventSlice
}

func (obj *convergenceMetricConvergenceEventIter) Add() ConvergenceEvent {
	newObj := &otg.ConvergenceEvent{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &convergenceEvent{obj: newObj}
	newLibObj.setDefault()
	obj.convergenceEventSlice = append(obj.convergenceEventSlice, newLibObj)
	return newLibObj
}

func (obj *convergenceMetricConvergenceEventIter) Append(items ...ConvergenceEvent) ConvergenceMetricConvergenceEventIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.convergenceEventSlice = append(obj.convergenceEventSlice, item)
	}
	return obj
}

func (obj *convergenceMetricConvergenceEventIter) Set(index int, newObj ConvergenceEvent) ConvergenceMetricConvergenceEventIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.convergenceEventSlice[index] = newObj
	return obj
}
func (obj *convergenceMetricConvergenceEventIter) Clear() ConvergenceMetricConvergenceEventIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ConvergenceEvent{}
		obj.convergenceEventSlice = []ConvergenceEvent{}
	}
	return obj
}
func (obj *convergenceMetricConvergenceEventIter) clearHolderSlice() ConvergenceMetricConvergenceEventIter {
	if len(obj.convergenceEventSlice) > 0 {
		obj.convergenceEventSlice = []ConvergenceEvent{}
	}
	return obj
}
func (obj *convergenceMetricConvergenceEventIter) appendHolderSlice(item ConvergenceEvent) ConvergenceMetricConvergenceEventIter {
	obj.convergenceEventSlice = append(obj.convergenceEventSlice, item)
	return obj
}

func (obj *convergenceMetric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	obj.addWarnings("ConvergenceMetric is under review, Convergence metrics is currently under review for pending exploration on use cases.")

	if len(obj.obj.Events) != 0 {

		if set_default {
			obj.Events().clearHolderSlice()
			for _, item := range obj.obj.Events {
				obj.Events().appendHolderSlice(&convergenceEvent{obj: item})
			}
		}
		for _, item := range obj.Events().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *convergenceMetric) setDefault() {

}
