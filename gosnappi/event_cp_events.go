package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** EventCPEvents *****
type eventCPEvents struct {
	validation
	obj          *otg.EventCPEvents
	marshaller   marshalEventCPEvents
	unMarshaller unMarshalEventCPEvents
}

func NewEventCPEvents() EventCPEvents {
	obj := eventCPEvents{obj: &otg.EventCPEvents{}}
	obj.setDefault()
	return &obj
}

func (obj *eventCPEvents) msg() *otg.EventCPEvents {
	return obj.obj
}

func (obj *eventCPEvents) setMsg(msg *otg.EventCPEvents) EventCPEvents {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaleventCPEvents struct {
	obj *eventCPEvents
}

type marshalEventCPEvents interface {
	// ToProto marshals EventCPEvents to protobuf object *otg.EventCPEvents
	ToProto() (*otg.EventCPEvents, error)
	// ToPbText marshals EventCPEvents to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals EventCPEvents to YAML text
	ToYaml() (string, error)
	// ToJson marshals EventCPEvents to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals EventCPEvents to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaleventCPEvents struct {
	obj *eventCPEvents
}

type unMarshalEventCPEvents interface {
	// FromProto unmarshals EventCPEvents from protobuf object *otg.EventCPEvents
	FromProto(msg *otg.EventCPEvents) (EventCPEvents, error)
	// FromPbText unmarshals EventCPEvents from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals EventCPEvents from YAML text
	FromYaml(value string) error
	// FromJson unmarshals EventCPEvents from JSON text
	FromJson(value string) error
}

func (obj *eventCPEvents) Marshal() marshalEventCPEvents {
	if obj.marshaller == nil {
		obj.marshaller = &marshaleventCPEvents{obj: obj}
	}
	return obj.marshaller
}

func (obj *eventCPEvents) Unmarshal() unMarshalEventCPEvents {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaleventCPEvents{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaleventCPEvents) ToProto() (*otg.EventCPEvents, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaleventCPEvents) FromProto(msg *otg.EventCPEvents) (EventCPEvents, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaleventCPEvents) ToPbText() (string, error) {
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

func (m *unMarshaleventCPEvents) FromPbText(value string) error {
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

func (m *marshaleventCPEvents) ToYaml() (string, error) {
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

func (m *unMarshaleventCPEvents) FromYaml(value string) error {
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

func (m *marshaleventCPEvents) ToJsonRaw() (string, error) {
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

func (m *marshaleventCPEvents) ToJson() (string, error) {
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

func (m *unMarshaleventCPEvents) FromJson(value string) error {
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

func (obj *eventCPEvents) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *eventCPEvents) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *eventCPEvents) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *eventCPEvents) Clone() (EventCPEvents, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewEventCPEvents()
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

// EventCPEvents is the optional container for control plane event configuration.
type EventCPEvents interface {
	Validation
	// msg marshals EventCPEvents to protobuf object *otg.EventCPEvents
	// and doesn't set defaults
	msg() *otg.EventCPEvents
	// setMsg unmarshals EventCPEvents from protobuf object *otg.EventCPEvents
	// and doesn't set defaults
	setMsg(*otg.EventCPEvents) EventCPEvents
	// provides marshal interface
	Marshal() marshalEventCPEvents
	// provides unmarshal interface
	Unmarshal() unMarshalEventCPEvents
	// validate validates EventCPEvents
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (EventCPEvents, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Enable returns bool, set in EventCPEvents.
	Enable() bool
	// SetEnable assigns bool provided by user to EventCPEvents
	SetEnable(value bool) EventCPEvents
	// HasEnable checks if Enable has been set in EventCPEvents
	HasEnable() bool
}

// Setting to true enables start and end time for control_plane events
// associated with started flows to be recorded.
// Enable returns a bool
func (obj *eventCPEvents) Enable() bool {

	return *obj.obj.Enable

}

// Setting to true enables start and end time for control_plane events
// associated with started flows to be recorded.
// Enable returns a bool
func (obj *eventCPEvents) HasEnable() bool {
	return obj.obj.Enable != nil
}

// Setting to true enables start and end time for control_plane events
// associated with started flows to be recorded.
// SetEnable sets the bool value in the EventCPEvents object
func (obj *eventCPEvents) SetEnable(value bool) EventCPEvents {

	obj.obj.Enable = &value
	return obj
}

func (obj *eventCPEvents) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *eventCPEvents) setDefault() {
	if obj.obj.Enable == nil {
		obj.SetEnable(false)
	}

}
