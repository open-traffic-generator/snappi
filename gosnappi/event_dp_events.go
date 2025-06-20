package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** EventDPEvents *****
type eventDPEvents struct {
	validation
	obj          *otg.EventDPEvents
	marshaller   marshalEventDPEvents
	unMarshaller unMarshalEventDPEvents
}

func NewEventDPEvents() EventDPEvents {
	obj := eventDPEvents{obj: &otg.EventDPEvents{}}
	obj.setDefault()
	return &obj
}

func (obj *eventDPEvents) msg() *otg.EventDPEvents {
	return obj.obj
}

func (obj *eventDPEvents) setMsg(msg *otg.EventDPEvents) EventDPEvents {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaleventDPEvents struct {
	obj *eventDPEvents
}

type marshalEventDPEvents interface {
	// ToProto marshals EventDPEvents to protobuf object *otg.EventDPEvents
	ToProto() (*otg.EventDPEvents, error)
	// ToPbText marshals EventDPEvents to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals EventDPEvents to YAML text
	ToYaml() (string, error)
	// ToJson marshals EventDPEvents to JSON text
	ToJson() (string, error)
}

type unMarshaleventDPEvents struct {
	obj *eventDPEvents
}

type unMarshalEventDPEvents interface {
	// FromProto unmarshals EventDPEvents from protobuf object *otg.EventDPEvents
	FromProto(msg *otg.EventDPEvents) (EventDPEvents, error)
	// FromPbText unmarshals EventDPEvents from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals EventDPEvents from YAML text
	FromYaml(value string) error
	// FromJson unmarshals EventDPEvents from JSON text
	FromJson(value string) error
}

func (obj *eventDPEvents) Marshal() marshalEventDPEvents {
	if obj.marshaller == nil {
		obj.marshaller = &marshaleventDPEvents{obj: obj}
	}
	return obj.marshaller
}

func (obj *eventDPEvents) Unmarshal() unMarshalEventDPEvents {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaleventDPEvents{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaleventDPEvents) ToProto() (*otg.EventDPEvents, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaleventDPEvents) FromProto(msg *otg.EventDPEvents) (EventDPEvents, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaleventDPEvents) ToPbText() (string, error) {
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

func (m *unMarshaleventDPEvents) FromPbText(value string) error {
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

func (m *marshaleventDPEvents) ToYaml() (string, error) {
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

func (m *unMarshaleventDPEvents) FromYaml(value string) error {
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

func (m *marshaleventDPEvents) ToJson() (string, error) {
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

func (m *unMarshaleventDPEvents) FromJson(value string) error {
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

func (obj *eventDPEvents) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *eventDPEvents) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *eventDPEvents) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *eventDPEvents) Clone() (EventDPEvents, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewEventDPEvents()
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

// EventDPEvents is the optional container for data plane event configuration.
type EventDPEvents interface {
	Validation
	// msg marshals EventDPEvents to protobuf object *otg.EventDPEvents
	// and doesn't set defaults
	msg() *otg.EventDPEvents
	// setMsg unmarshals EventDPEvents from protobuf object *otg.EventDPEvents
	// and doesn't set defaults
	setMsg(*otg.EventDPEvents) EventDPEvents
	// provides marshal interface
	Marshal() marshalEventDPEvents
	// provides unmarshal interface
	Unmarshal() unMarshalEventDPEvents
	// validate validates EventDPEvents
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (EventDPEvents, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Enable returns bool, set in EventDPEvents.
	Enable() bool
	// SetEnable assigns bool provided by user to EventDPEvents
	SetEnable(value bool) EventDPEvents
	// HasEnable checks if Enable has been set in EventDPEvents
	HasEnable() bool
	// RxRateThreshold returns float32, set in EventDPEvents.
	RxRateThreshold() float32
	// SetRxRateThreshold assigns float32 provided by user to EventDPEvents
	SetRxRateThreshold(value float32) EventDPEvents
	// HasRxRateThreshold checks if RxRateThreshold has been set in EventDPEvents
	HasRxRateThreshold() bool
}

// Setting to true enables flow_rx_rate_above_threshold and flow_rx_rate_below_threshold
// timestamps to be recorded when data packets switch between multiple rx_ports on the receive ports for affected flows.
// Enable returns a bool
func (obj *eventDPEvents) Enable() bool {

	return *obj.obj.Enable

}

// Setting to true enables flow_rx_rate_above_threshold and flow_rx_rate_below_threshold
// timestamps to be recorded when data packets switch between multiple rx_ports on the receive ports for affected flows.
// Enable returns a bool
func (obj *eventDPEvents) HasEnable() bool {
	return obj.obj.Enable != nil
}

// Setting to true enables flow_rx_rate_above_threshold and flow_rx_rate_below_threshold
// timestamps to be recorded when data packets switch between multiple rx_ports on the receive ports for affected flows.
// SetEnable sets the bool value in the EventDPEvents object
func (obj *eventDPEvents) SetEnable(value bool) EventDPEvents {

	obj.obj.Enable = &value
	return obj
}

// Setting to true enables timestamps to be recorded when the rx rate of a flow goes above
// or below the threshold value.
// RxRateThreshold returns a float32
func (obj *eventDPEvents) RxRateThreshold() float32 {

	return *obj.obj.RxRateThreshold

}

// Setting to true enables timestamps to be recorded when the rx rate of a flow goes above
// or below the threshold value.
// RxRateThreshold returns a float32
func (obj *eventDPEvents) HasRxRateThreshold() bool {
	return obj.obj.RxRateThreshold != nil
}

// Setting to true enables timestamps to be recorded when the rx rate of a flow goes above
// or below the threshold value.
// SetRxRateThreshold sets the float32 value in the EventDPEvents object
func (obj *eventDPEvents) SetRxRateThreshold(value float32) EventDPEvents {

	obj.obj.RxRateThreshold = &value
	return obj
}

func (obj *eventDPEvents) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RxRateThreshold != nil {

		if *obj.obj.RxRateThreshold < 0 || *obj.obj.RxRateThreshold > 100 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= EventDPEvents.RxRateThreshold <= 100 but Got %f", *obj.obj.RxRateThreshold))
		}

	}

}

func (obj *eventDPEvents) setDefault() {
	if obj.obj.Enable == nil {
		obj.SetEnable(false)
	}
	if obj.obj.RxRateThreshold == nil {
		obj.SetRxRateThreshold(95)
	}

}
