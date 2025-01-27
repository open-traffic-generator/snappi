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
	obj                          *otg.Event
	marshaller                   marshalEvent
	unMarshaller                 unMarshalEvent
	linkHolder                   EventLink
	rxRateThresholdHolder        EventRxRateThreshold
	routeAdvertiseWithdrawHolder EventRouteAdvertiseWithdraw
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
	// ToJsonRaw marshals Event to raw JSON text
	ToJsonRaw() (string, error)
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

func (m *marshalevent) ToJsonRaw() (string, error) {
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
	obj.linkHolder = nil
	obj.rxRateThresholdHolder = nil
	obj.routeAdvertiseWithdrawHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Event is the optional container for event configuration.
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
	// Enable returns bool, set in Event.
	Enable() bool
	// SetEnable assigns bool provided by user to Event
	SetEnable(value bool) Event
	// HasEnable checks if Enable has been set in Event
	HasEnable() bool
	// Link returns EventLink, set in Event.
	// EventLink is the optional container for link up/down event configuration.
	Link() EventLink
	// SetLink assigns EventLink provided by user to Event.
	// EventLink is the optional container for link up/down event configuration.
	SetLink(value EventLink) Event
	// HasLink checks if Link has been set in Event
	HasLink() bool
	// RxRateThreshold returns EventRxRateThreshold, set in Event.
	// EventRxRateThreshold is the optional container for rx rate threshold event configuration.
	RxRateThreshold() EventRxRateThreshold
	// SetRxRateThreshold assigns EventRxRateThreshold provided by user to Event.
	// EventRxRateThreshold is the optional container for rx rate threshold event configuration.
	SetRxRateThreshold(value EventRxRateThreshold) Event
	// HasRxRateThreshold checks if RxRateThreshold has been set in Event
	HasRxRateThreshold() bool
	// RouteAdvertiseWithdraw returns EventRouteAdvertiseWithdraw, set in Event.
	// EventRouteAdvertiseWithdraw is the optional container for route advertise/withdraw event configuration.
	RouteAdvertiseWithdraw() EventRouteAdvertiseWithdraw
	// SetRouteAdvertiseWithdraw assigns EventRouteAdvertiseWithdraw provided by user to Event.
	// EventRouteAdvertiseWithdraw is the optional container for route advertise/withdraw event configuration.
	SetRouteAdvertiseWithdraw(value EventRouteAdvertiseWithdraw) Event
	// HasRouteAdvertiseWithdraw checks if RouteAdvertiseWithdraw has been set in Event
	HasRouteAdvertiseWithdraw() bool
	setNil()
}

// True to enable all events.
// Enabling this option may affect the resultant packet payload due to
// additional instrumentation data.
// Enable returns a bool
func (obj *event) Enable() bool {

	return *obj.obj.Enable

}

// True to enable all events.
// Enabling this option may affect the resultant packet payload due to
// additional instrumentation data.
// Enable returns a bool
func (obj *event) HasEnable() bool {
	return obj.obj.Enable != nil
}

// True to enable all events.
// Enabling this option may affect the resultant packet payload due to
// additional instrumentation data.
// SetEnable sets the bool value in the Event object
func (obj *event) SetEnable(value bool) Event {

	obj.obj.Enable = &value
	return obj
}

// description is TBD
// Link returns a EventLink
func (obj *event) Link() EventLink {
	if obj.obj.Link == nil {
		obj.obj.Link = NewEventLink().msg()
	}
	if obj.linkHolder == nil {
		obj.linkHolder = &eventLink{obj: obj.obj.Link}
	}
	return obj.linkHolder
}

// description is TBD
// Link returns a EventLink
func (obj *event) HasLink() bool {
	return obj.obj.Link != nil
}

// description is TBD
// SetLink sets the EventLink value in the Event object
func (obj *event) SetLink(value EventLink) Event {

	obj.linkHolder = nil
	obj.obj.Link = value.msg()

	return obj
}

// description is TBD
// RxRateThreshold returns a EventRxRateThreshold
func (obj *event) RxRateThreshold() EventRxRateThreshold {
	if obj.obj.RxRateThreshold == nil {
		obj.obj.RxRateThreshold = NewEventRxRateThreshold().msg()
	}
	if obj.rxRateThresholdHolder == nil {
		obj.rxRateThresholdHolder = &eventRxRateThreshold{obj: obj.obj.RxRateThreshold}
	}
	return obj.rxRateThresholdHolder
}

// description is TBD
// RxRateThreshold returns a EventRxRateThreshold
func (obj *event) HasRxRateThreshold() bool {
	return obj.obj.RxRateThreshold != nil
}

// description is TBD
// SetRxRateThreshold sets the EventRxRateThreshold value in the Event object
func (obj *event) SetRxRateThreshold(value EventRxRateThreshold) Event {

	obj.rxRateThresholdHolder = nil
	obj.obj.RxRateThreshold = value.msg()

	return obj
}

// description is TBD
// RouteAdvertiseWithdraw returns a EventRouteAdvertiseWithdraw
func (obj *event) RouteAdvertiseWithdraw() EventRouteAdvertiseWithdraw {
	if obj.obj.RouteAdvertiseWithdraw == nil {
		obj.obj.RouteAdvertiseWithdraw = NewEventRouteAdvertiseWithdraw().msg()
	}
	if obj.routeAdvertiseWithdrawHolder == nil {
		obj.routeAdvertiseWithdrawHolder = &eventRouteAdvertiseWithdraw{obj: obj.obj.RouteAdvertiseWithdraw}
	}
	return obj.routeAdvertiseWithdrawHolder
}

// description is TBD
// RouteAdvertiseWithdraw returns a EventRouteAdvertiseWithdraw
func (obj *event) HasRouteAdvertiseWithdraw() bool {
	return obj.obj.RouteAdvertiseWithdraw != nil
}

// description is TBD
// SetRouteAdvertiseWithdraw sets the EventRouteAdvertiseWithdraw value in the Event object
func (obj *event) SetRouteAdvertiseWithdraw(value EventRouteAdvertiseWithdraw) Event {

	obj.routeAdvertiseWithdrawHolder = nil
	obj.obj.RouteAdvertiseWithdraw = value.msg()

	return obj
}

func (obj *event) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Link != nil {

		obj.Link().validateObj(vObj, set_default)
	}

	if obj.obj.RxRateThreshold != nil {

		obj.RxRateThreshold().validateObj(vObj, set_default)
	}

	if obj.obj.RouteAdvertiseWithdraw != nil {

		obj.RouteAdvertiseWithdraw().validateObj(vObj, set_default)
	}

}

func (obj *event) setDefault() {
	if obj.obj.Enable == nil {
		obj.SetEnable(false)
	}

}
