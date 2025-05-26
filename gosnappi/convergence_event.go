package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ConvergenceEvent *****
type convergenceEvent struct {
	validation
	obj          *otg.ConvergenceEvent
	marshaller   marshalConvergenceEvent
	unMarshaller unMarshalConvergenceEvent
}

func NewConvergenceEvent() ConvergenceEvent {
	obj := convergenceEvent{obj: &otg.ConvergenceEvent{}}
	obj.setDefault()
	return &obj
}

func (obj *convergenceEvent) msg() *otg.ConvergenceEvent {
	return obj.obj
}

func (obj *convergenceEvent) setMsg(msg *otg.ConvergenceEvent) ConvergenceEvent {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalconvergenceEvent struct {
	obj *convergenceEvent
}

type marshalConvergenceEvent interface {
	// ToProto marshals ConvergenceEvent to protobuf object *otg.ConvergenceEvent
	ToProto() (*otg.ConvergenceEvent, error)
	// ToPbText marshals ConvergenceEvent to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ConvergenceEvent to YAML text
	ToYaml() (string, error)
	// ToJson marshals ConvergenceEvent to JSON text
	ToJson() (string, error)
}

type unMarshalconvergenceEvent struct {
	obj *convergenceEvent
}

type unMarshalConvergenceEvent interface {
	// FromProto unmarshals ConvergenceEvent from protobuf object *otg.ConvergenceEvent
	FromProto(msg *otg.ConvergenceEvent) (ConvergenceEvent, error)
	// FromPbText unmarshals ConvergenceEvent from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ConvergenceEvent from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ConvergenceEvent from JSON text
	FromJson(value string) error
}

func (obj *convergenceEvent) Marshal() marshalConvergenceEvent {
	if obj.marshaller == nil {
		obj.marshaller = &marshalconvergenceEvent{obj: obj}
	}
	return obj.marshaller
}

func (obj *convergenceEvent) Unmarshal() unMarshalConvergenceEvent {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalconvergenceEvent{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalconvergenceEvent) ToProto() (*otg.ConvergenceEvent, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalconvergenceEvent) FromProto(msg *otg.ConvergenceEvent) (ConvergenceEvent, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalconvergenceEvent) ToPbText() (string, error) {
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

func (m *unMarshalconvergenceEvent) FromPbText(value string) error {
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

func (m *marshalconvergenceEvent) ToYaml() (string, error) {
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

func (m *unMarshalconvergenceEvent) FromYaml(value string) error {
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

func (m *marshalconvergenceEvent) ToJson() (string, error) {
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

func (m *unMarshalconvergenceEvent) FromJson(value string) error {
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

func (obj *convergenceEvent) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *convergenceEvent) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *convergenceEvent) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *convergenceEvent) Clone() (ConvergenceEvent, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewConvergenceEvent()
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

// ConvergenceEvent is a container for an event that has occurred in the system affecting the convergence time recorded for the flow.
type ConvergenceEvent interface {
	Validation
	// msg marshals ConvergenceEvent to protobuf object *otg.ConvergenceEvent
	// and doesn't set defaults
	msg() *otg.ConvergenceEvent
	// setMsg unmarshals ConvergenceEvent from protobuf object *otg.ConvergenceEvent
	// and doesn't set defaults
	setMsg(*otg.ConvergenceEvent) ConvergenceEvent
	// provides marshal interface
	Marshal() marshalConvergenceEvent
	// provides unmarshal interface
	Unmarshal() unMarshalConvergenceEvent
	// validate validates ConvergenceEvent
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ConvergenceEvent, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns ConvergenceEventTypeEnum, set in ConvergenceEvent
	Type() ConvergenceEventTypeEnum
	// SetType assigns ConvergenceEventTypeEnum provided by user to ConvergenceEvent
	SetType(value ConvergenceEventTypeEnum) ConvergenceEvent
	// HasType checks if Type has been set in ConvergenceEvent
	HasType() bool
	// Source returns string, set in ConvergenceEvent.
	Source() string
	// SetSource assigns string provided by user to ConvergenceEvent
	SetSource(value string) ConvergenceEvent
	// HasSource checks if Source has been set in ConvergenceEvent
	HasSource() bool
	// BeginTimestampNs returns float64, set in ConvergenceEvent.
	BeginTimestampNs() float64
	// SetBeginTimestampNs assigns float64 provided by user to ConvergenceEvent
	SetBeginTimestampNs(value float64) ConvergenceEvent
	// HasBeginTimestampNs checks if BeginTimestampNs has been set in ConvergenceEvent
	HasBeginTimestampNs() bool
	// EndTimestampNs returns float64, set in ConvergenceEvent.
	EndTimestampNs() float64
	// SetEndTimestampNs assigns float64 provided by user to ConvergenceEvent
	SetEndTimestampNs(value float64) ConvergenceEvent
	// HasEndTimestampNs checks if EndTimestampNs has been set in ConvergenceEvent
	HasEndTimestampNs() bool
}

type ConvergenceEventTypeEnum string

// Enum of Type on ConvergenceEvent
var ConvergenceEventType = struct {
	LINK_DOWN                    ConvergenceEventTypeEnum
	LINK_UP                      ConvergenceEventTypeEnum
	ROUTE_WITHDRAW               ConvergenceEventTypeEnum
	ROUTE_ADVERTISE              ConvergenceEventTypeEnum
	FLOW_RX_RATE_ABOVE_THRESHOLD ConvergenceEventTypeEnum
	FLOW_RX_RATE_BELOW_THRESHOLD ConvergenceEventTypeEnum
}{
	LINK_DOWN:                    ConvergenceEventTypeEnum("link_down"),
	LINK_UP:                      ConvergenceEventTypeEnum("link_up"),
	ROUTE_WITHDRAW:               ConvergenceEventTypeEnum("route_withdraw"),
	ROUTE_ADVERTISE:              ConvergenceEventTypeEnum("route_advertise"),
	FLOW_RX_RATE_ABOVE_THRESHOLD: ConvergenceEventTypeEnum("flow_rx_rate_above_threshold"),
	FLOW_RX_RATE_BELOW_THRESHOLD: ConvergenceEventTypeEnum("flow_rx_rate_below_threshold"),
}

func (obj *convergenceEvent) Type() ConvergenceEventTypeEnum {
	return ConvergenceEventTypeEnum(obj.obj.Type.Enum().String())
}

// The type of control plane or data plane event that occurred.
// Type returns a string
func (obj *convergenceEvent) HasType() bool {
	return obj.obj.Type != nil
}

func (obj *convergenceEvent) SetType(value ConvergenceEventTypeEnum) ConvergenceEvent {
	intValue, ok := otg.ConvergenceEvent_Type_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ConvergenceEventTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.ConvergenceEvent_Type_Enum(intValue)
	obj.obj.Type = &enumValue

	return obj
}

// The source of the event.
// The source MUST be the value of one of the x-constraint paths,
// which means the source MUST be a unique name in the configuration.
//
// x-constraint:
// - /components/schemas/Port/properties/name
// - /components/schemas/Flow/properties/name
// - /components/schemas/Device.Bgpv4Route/properties/name
// - /components/schemas/Device.Bgpv6Route/properties/name
//
// x-constraint:
// - /components/schemas/Port/properties/name
// - /components/schemas/Flow/properties/name
// - /components/schemas/Device.Bgpv4Route/properties/name
// - /components/schemas/Device.Bgpv6Route/properties/name
//
// Source returns a string
func (obj *convergenceEvent) Source() string {

	return *obj.obj.Source

}

// The source of the event.
// The source MUST be the value of one of the x-constraint paths,
// which means the source MUST be a unique name in the configuration.
//
// x-constraint:
// - /components/schemas/Port/properties/name
// - /components/schemas/Flow/properties/name
// - /components/schemas/Device.Bgpv4Route/properties/name
// - /components/schemas/Device.Bgpv6Route/properties/name
//
// x-constraint:
// - /components/schemas/Port/properties/name
// - /components/schemas/Flow/properties/name
// - /components/schemas/Device.Bgpv4Route/properties/name
// - /components/schemas/Device.Bgpv6Route/properties/name
//
// Source returns a string
func (obj *convergenceEvent) HasSource() bool {
	return obj.obj.Source != nil
}

// The source of the event.
// The source MUST be the value of one of the x-constraint paths,
// which means the source MUST be a unique name in the configuration.
//
// x-constraint:
// - /components/schemas/Port/properties/name
// - /components/schemas/Flow/properties/name
// - /components/schemas/Device.Bgpv4Route/properties/name
// - /components/schemas/Device.Bgpv6Route/properties/name
//
// x-constraint:
// - /components/schemas/Port/properties/name
// - /components/schemas/Flow/properties/name
// - /components/schemas/Device.Bgpv4Route/properties/name
// - /components/schemas/Device.Bgpv6Route/properties/name
//
// SetSource sets the string value in the ConvergenceEvent object
func (obj *convergenceEvent) SetSource(value string) ConvergenceEvent {

	obj.obj.Source = &value
	return obj
}

// The timestamp(nanoseconds) of the starting event that triggers convergence.
// BeginTimestampNs returns a float64
func (obj *convergenceEvent) BeginTimestampNs() float64 {

	return *obj.obj.BeginTimestampNs

}

// The timestamp(nanoseconds) of the starting event that triggers convergence.
// BeginTimestampNs returns a float64
func (obj *convergenceEvent) HasBeginTimestampNs() bool {
	return obj.obj.BeginTimestampNs != nil
}

// The timestamp(nanoseconds) of the starting event that triggers convergence.
// SetBeginTimestampNs sets the float64 value in the ConvergenceEvent object
func (obj *convergenceEvent) SetBeginTimestampNs(value float64) ConvergenceEvent {

	obj.obj.BeginTimestampNs = &value
	return obj
}

// The timestamp(nanoseconds) of the end event that triggers convergence.
// EndTimestampNs returns a float64
func (obj *convergenceEvent) EndTimestampNs() float64 {

	return *obj.obj.EndTimestampNs

}

// The timestamp(nanoseconds) of the end event that triggers convergence.
// EndTimestampNs returns a float64
func (obj *convergenceEvent) HasEndTimestampNs() bool {
	return obj.obj.EndTimestampNs != nil
}

// The timestamp(nanoseconds) of the end event that triggers convergence.
// SetEndTimestampNs sets the float64 value in the ConvergenceEvent object
func (obj *convergenceEvent) SetEndTimestampNs(value float64) ConvergenceEvent {

	obj.obj.EndTimestampNs = &value
	return obj
}

func (obj *convergenceEvent) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *convergenceEvent) setDefault() {

}
