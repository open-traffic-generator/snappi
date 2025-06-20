package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Event *****
type event struct {
	validation
	obj            *otg.Event
	marshaller     marshalEvent
	unMarshaller   unMarshalEvent
	cpEventsHolder EventCPEvents
	dpEventsHolder EventDPEvents
}

func NewEvent() Event {
	obj := event{obj: &otg.Event{}}
	obj.setDefault()
	return &obj
}

func (obj *event) msg() *otg.Event {
	return obj.obj
}

func (obj *event) setMsg(msg *otg.Event) Event {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalevent struct {
	obj *event
}

type marshalEvent interface {
	// ToProto marshals Event to protobuf object *otg.Event
	ToProto() (*otg.Event, error)
	// ToPbText marshals Event to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Event to YAML text
	ToYaml() (string, error)
	// ToJson marshals Event to JSON text
	ToJson() (string, error)
}

type unMarshalevent struct {
	obj *event
}

type unMarshalEvent interface {
	// FromProto unmarshals Event from protobuf object *otg.Event
	FromProto(msg *otg.Event) (Event, error)
	// FromPbText unmarshals Event from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Event from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Event from JSON text
	FromJson(value string) error
}

func (obj *event) Marshal() marshalEvent {
	if obj.marshaller == nil {
		obj.marshaller = &marshalevent{obj: obj}
	}
	return obj.marshaller
}

func (obj *event) Unmarshal() unMarshalEvent {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalevent{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalevent) ToProto() (*otg.Event, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalevent) FromProto(msg *otg.Event) (Event, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalevent) ToPbText() (string, error) {
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

func (m *unMarshalevent) FromPbText(value string) error {
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

func (m *marshalevent) ToYaml() (string, error) {
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

func (m *unMarshalevent) FromYaml(value string) error {
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

func (m *marshalevent) ToJson() (string, error) {
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

func (m *unMarshalevent) FromJson(value string) error {
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

func (obj *event) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *event) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *event) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *event) Clone() (Event, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewEvent()
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

func (obj *event) setNil() {
	obj.cpEventsHolder = nil
	obj.dpEventsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Event is under Review: Event configuration is currently under review for pending exploration on use cases.
//
// The optional container for event configuration.
// Both cp_events.enable and dp_events.enable must be explicitly set to true to get
// control_plane_data_plane_convergence_us metric values for convergence metrics.
type Event interface {
	Validation
	// msg marshals Event to protobuf object *otg.Event
	// and doesn't set defaults
	msg() *otg.Event
	// setMsg unmarshals Event from protobuf object *otg.Event
	// and doesn't set defaults
	setMsg(*otg.Event) Event
	// provides marshal interface
	Marshal() marshalEvent
	// provides unmarshal interface
	Unmarshal() unMarshalEvent
	// validate validates Event
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Event, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// CpEvents returns EventCPEvents, set in Event.
	// EventCPEvents is the optional container for control plane event configuration.
	CpEvents() EventCPEvents
	// SetCpEvents assigns EventCPEvents provided by user to Event.
	// EventCPEvents is the optional container for control plane event configuration.
	SetCpEvents(value EventCPEvents) Event
	// HasCpEvents checks if CpEvents has been set in Event
	HasCpEvents() bool
	// DpEvents returns EventDPEvents, set in Event.
	// EventDPEvents is the optional container for data plane event configuration.
	DpEvents() EventDPEvents
	// SetDpEvents assigns EventDPEvents provided by user to Event.
	// EventDPEvents is the optional container for data plane event configuration.
	SetDpEvents(value EventDPEvents) Event
	// HasDpEvents checks if DpEvents has been set in Event
	HasDpEvents() bool
	setNil()
}

// Container for control plane(cp) event configuration.
// CpEvents returns a EventCPEvents
func (obj *event) CpEvents() EventCPEvents {
	if obj.obj.CpEvents == nil {
		obj.obj.CpEvents = NewEventCPEvents().msg()
	}
	if obj.cpEventsHolder == nil {
		obj.cpEventsHolder = &eventCPEvents{obj: obj.obj.CpEvents}
	}
	return obj.cpEventsHolder
}

// Container for control plane(cp) event configuration.
// CpEvents returns a EventCPEvents
func (obj *event) HasCpEvents() bool {
	return obj.obj.CpEvents != nil
}

// Container for control plane(cp) event configuration.
// SetCpEvents sets the EventCPEvents value in the Event object
func (obj *event) SetCpEvents(value EventCPEvents) Event {

	obj.cpEventsHolder = nil
	obj.obj.CpEvents = value.msg()

	return obj
}

// Container for data plane(dp) event configuration.
// Enabling this option may affect the resultant packet payload due to
// additional instrumentation data.
// DpEvents returns a EventDPEvents
func (obj *event) DpEvents() EventDPEvents {
	if obj.obj.DpEvents == nil {
		obj.obj.DpEvents = NewEventDPEvents().msg()
	}
	if obj.dpEventsHolder == nil {
		obj.dpEventsHolder = &eventDPEvents{obj: obj.obj.DpEvents}
	}
	return obj.dpEventsHolder
}

// Container for data plane(dp) event configuration.
// Enabling this option may affect the resultant packet payload due to
// additional instrumentation data.
// DpEvents returns a EventDPEvents
func (obj *event) HasDpEvents() bool {
	return obj.obj.DpEvents != nil
}

// Container for data plane(dp) event configuration.
// Enabling this option may affect the resultant packet payload due to
// additional instrumentation data.
// SetDpEvents sets the EventDPEvents value in the Event object
func (obj *event) SetDpEvents(value EventDPEvents) Event {

	obj.dpEventsHolder = nil
	obj.obj.DpEvents = value.msg()

	return obj
}

func (obj *event) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	obj.addWarnings("Event is under review, Event configuration is currently under review for pending exploration on use cases.")

	if obj.obj.CpEvents != nil {

		obj.CpEvents().validateObj(vObj, set_default)
	}

	if obj.obj.DpEvents != nil {

		obj.DpEvents().validateObj(vObj, set_default)
	}

}

func (obj *event) setDefault() {

}
