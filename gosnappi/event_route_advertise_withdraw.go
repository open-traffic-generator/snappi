package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** EventRouteAdvertiseWithdraw *****
type eventRouteAdvertiseWithdraw struct {
	validation
	obj          *otg.EventRouteAdvertiseWithdraw
	marshaller   marshalEventRouteAdvertiseWithdraw
	unMarshaller unMarshalEventRouteAdvertiseWithdraw
}

func NewEventRouteAdvertiseWithdraw() EventRouteAdvertiseWithdraw {
	obj := eventRouteAdvertiseWithdraw{obj: &otg.EventRouteAdvertiseWithdraw{}}
	obj.setDefault()
	return &obj
}

func (obj *eventRouteAdvertiseWithdraw) msg() *otg.EventRouteAdvertiseWithdraw {
	return obj.obj
}

func (obj *eventRouteAdvertiseWithdraw) setMsg(msg *otg.EventRouteAdvertiseWithdraw) EventRouteAdvertiseWithdraw {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaleventRouteAdvertiseWithdraw struct {
	obj *eventRouteAdvertiseWithdraw
}

type marshalEventRouteAdvertiseWithdraw interface {
	// ToProto marshals EventRouteAdvertiseWithdraw to protobuf object *otg.EventRouteAdvertiseWithdraw
	ToProto() (*otg.EventRouteAdvertiseWithdraw, error)
	// ToPbText marshals EventRouteAdvertiseWithdraw to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals EventRouteAdvertiseWithdraw to YAML text
	ToYaml() (string, error)
	// ToJson marshals EventRouteAdvertiseWithdraw to JSON text
	ToJson() (string, error)
}

type unMarshaleventRouteAdvertiseWithdraw struct {
	obj *eventRouteAdvertiseWithdraw
}

type unMarshalEventRouteAdvertiseWithdraw interface {
	// FromProto unmarshals EventRouteAdvertiseWithdraw from protobuf object *otg.EventRouteAdvertiseWithdraw
	FromProto(msg *otg.EventRouteAdvertiseWithdraw) (EventRouteAdvertiseWithdraw, error)
	// FromPbText unmarshals EventRouteAdvertiseWithdraw from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals EventRouteAdvertiseWithdraw from YAML text
	FromYaml(value string) error
	// FromJson unmarshals EventRouteAdvertiseWithdraw from JSON text
	FromJson(value string) error
}

func (obj *eventRouteAdvertiseWithdraw) Marshal() marshalEventRouteAdvertiseWithdraw {
	if obj.marshaller == nil {
		obj.marshaller = &marshaleventRouteAdvertiseWithdraw{obj: obj}
	}
	return obj.marshaller
}

func (obj *eventRouteAdvertiseWithdraw) Unmarshal() unMarshalEventRouteAdvertiseWithdraw {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaleventRouteAdvertiseWithdraw{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaleventRouteAdvertiseWithdraw) ToProto() (*otg.EventRouteAdvertiseWithdraw, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaleventRouteAdvertiseWithdraw) FromProto(msg *otg.EventRouteAdvertiseWithdraw) (EventRouteAdvertiseWithdraw, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaleventRouteAdvertiseWithdraw) ToPbText() (string, error) {
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

func (m *unMarshaleventRouteAdvertiseWithdraw) FromPbText(value string) error {
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

func (m *marshaleventRouteAdvertiseWithdraw) ToYaml() (string, error) {
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

func (m *unMarshaleventRouteAdvertiseWithdraw) FromYaml(value string) error {
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

func (m *marshaleventRouteAdvertiseWithdraw) ToJson() (string, error) {
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

func (m *unMarshaleventRouteAdvertiseWithdraw) FromJson(value string) error {
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

func (obj *eventRouteAdvertiseWithdraw) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *eventRouteAdvertiseWithdraw) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *eventRouteAdvertiseWithdraw) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *eventRouteAdvertiseWithdraw) Clone() (EventRouteAdvertiseWithdraw, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewEventRouteAdvertiseWithdraw()
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

// EventRouteAdvertiseWithdraw is the optional container for route advertise/withdraw event configuration.
type EventRouteAdvertiseWithdraw interface {
	Validation
	// msg marshals EventRouteAdvertiseWithdraw to protobuf object *otg.EventRouteAdvertiseWithdraw
	// and doesn't set defaults
	msg() *otg.EventRouteAdvertiseWithdraw
	// setMsg unmarshals EventRouteAdvertiseWithdraw from protobuf object *otg.EventRouteAdvertiseWithdraw
	// and doesn't set defaults
	setMsg(*otg.EventRouteAdvertiseWithdraw) EventRouteAdvertiseWithdraw
	// provides marshal interface
	Marshal() marshalEventRouteAdvertiseWithdraw
	// provides unmarshal interface
	Unmarshal() unMarshalEventRouteAdvertiseWithdraw
	// validate validates EventRouteAdvertiseWithdraw
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (EventRouteAdvertiseWithdraw, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Enable returns bool, set in EventRouteAdvertiseWithdraw.
	Enable() bool
	// SetEnable assigns bool provided by user to EventRouteAdvertiseWithdraw
	SetEnable(value bool) EventRouteAdvertiseWithdraw
	// HasEnable checks if Enable has been set in EventRouteAdvertiseWithdraw
	HasEnable() bool
}

// True to enable notifications when a route advertise/withdraw
// event occurs.
// Enable returns a bool
func (obj *eventRouteAdvertiseWithdraw) Enable() bool {

	return *obj.obj.Enable

}

// True to enable notifications when a route advertise/withdraw
// event occurs.
// Enable returns a bool
func (obj *eventRouteAdvertiseWithdraw) HasEnable() bool {
	return obj.obj.Enable != nil
}

// True to enable notifications when a route advertise/withdraw
// event occurs.
// SetEnable sets the bool value in the EventRouteAdvertiseWithdraw object
func (obj *eventRouteAdvertiseWithdraw) SetEnable(value bool) EventRouteAdvertiseWithdraw {

	obj.obj.Enable = &value
	return obj
}

func (obj *eventRouteAdvertiseWithdraw) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *eventRouteAdvertiseWithdraw) setDefault() {
	if obj.obj.Enable == nil {
		obj.SetEnable(false)
	}

}
