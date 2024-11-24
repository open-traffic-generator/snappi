package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** EventRxRateThreshold *****
type eventRxRateThreshold struct {
	validation
	obj          *otg.EventRxRateThreshold
	marshaller   marshalEventRxRateThreshold
	unMarshaller unMarshalEventRxRateThreshold
}

func NewEventRxRateThreshold() EventRxRateThreshold {
	obj := eventRxRateThreshold{obj: &otg.EventRxRateThreshold{}}
	obj.setDefault()
	return &obj
}

func (obj *eventRxRateThreshold) msg() *otg.EventRxRateThreshold {
	return obj.obj
}

func (obj *eventRxRateThreshold) setMsg(msg *otg.EventRxRateThreshold) EventRxRateThreshold {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaleventRxRateThreshold struct {
	obj *eventRxRateThreshold
}

type marshalEventRxRateThreshold interface {
	// ToProto marshals EventRxRateThreshold to protobuf object *otg.EventRxRateThreshold
	ToProto() (*otg.EventRxRateThreshold, error)
	// ToPbText marshals EventRxRateThreshold to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals EventRxRateThreshold to YAML text
	ToYaml() (string, error)
	// ToJson marshals EventRxRateThreshold to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals EventRxRateThreshold to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaleventRxRateThreshold struct {
	obj *eventRxRateThreshold
}

type unMarshalEventRxRateThreshold interface {
	// FromProto unmarshals EventRxRateThreshold from protobuf object *otg.EventRxRateThreshold
	FromProto(msg *otg.EventRxRateThreshold) (EventRxRateThreshold, error)
	// FromPbText unmarshals EventRxRateThreshold from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals EventRxRateThreshold from YAML text
	FromYaml(value string) error
	// FromJson unmarshals EventRxRateThreshold from JSON text
	FromJson(value string) error
}

func (obj *eventRxRateThreshold) Marshal() marshalEventRxRateThreshold {
	if obj.marshaller == nil {
		obj.marshaller = &marshaleventRxRateThreshold{obj: obj}
	}
	return obj.marshaller
}

func (obj *eventRxRateThreshold) Unmarshal() unMarshalEventRxRateThreshold {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaleventRxRateThreshold{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaleventRxRateThreshold) ToProto() (*otg.EventRxRateThreshold, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaleventRxRateThreshold) FromProto(msg *otg.EventRxRateThreshold) (EventRxRateThreshold, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaleventRxRateThreshold) ToPbText() (string, error) {
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

func (m *unMarshaleventRxRateThreshold) FromPbText(value string) error {
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

func (m *marshaleventRxRateThreshold) ToYaml() (string, error) {
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

func (m *unMarshaleventRxRateThreshold) FromYaml(value string) error {
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

func (m *marshaleventRxRateThreshold) ToJsonRaw() (string, error) {
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

func (m *marshaleventRxRateThreshold) ToJson() (string, error) {
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

func (m *unMarshaleventRxRateThreshold) FromJson(value string) error {
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

func (obj *eventRxRateThreshold) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *eventRxRateThreshold) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *eventRxRateThreshold) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *eventRxRateThreshold) Clone() (EventRxRateThreshold, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewEventRxRateThreshold()
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

// EventRxRateThreshold is the optional container for rx rate threshold event configuration.
type EventRxRateThreshold interface {
	Validation
	// msg marshals EventRxRateThreshold to protobuf object *otg.EventRxRateThreshold
	// and doesn't set defaults
	msg() *otg.EventRxRateThreshold
	// setMsg unmarshals EventRxRateThreshold from protobuf object *otg.EventRxRateThreshold
	// and doesn't set defaults
	setMsg(*otg.EventRxRateThreshold) EventRxRateThreshold
	// provides marshal interface
	Marshal() marshalEventRxRateThreshold
	// provides unmarshal interface
	Unmarshal() unMarshalEventRxRateThreshold
	// validate validates EventRxRateThreshold
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (EventRxRateThreshold, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Enable returns bool, set in EventRxRateThreshold.
	Enable() bool
	// SetEnable assigns bool provided by user to EventRxRateThreshold
	SetEnable(value bool) EventRxRateThreshold
	// HasEnable checks if Enable has been set in EventRxRateThreshold
	HasEnable() bool
	// Threshold returns float32, set in EventRxRateThreshold.
	Threshold() float32
	// SetThreshold assigns float32 provided by user to EventRxRateThreshold
	SetThreshold(value float32) EventRxRateThreshold
	// HasThreshold checks if Threshold has been set in EventRxRateThreshold
	HasThreshold() bool
}

// True to enable the rx_rate_threshold event.
// Enabling this option may affect the resultant packet payload due to
// additional instrumentation data.
// Enable returns a bool
func (obj *eventRxRateThreshold) Enable() bool {

	return *obj.obj.Enable

}

// True to enable the rx_rate_threshold event.
// Enabling this option may affect the resultant packet payload due to
// additional instrumentation data.
// Enable returns a bool
func (obj *eventRxRateThreshold) HasEnable() bool {
	return obj.obj.Enable != nil
}

// True to enable the rx_rate_threshold event.
// Enabling this option may affect the resultant packet payload due to
// additional instrumentation data.
// SetEnable sets the bool value in the EventRxRateThreshold object
func (obj *eventRxRateThreshold) SetEnable(value bool) EventRxRateThreshold {

	obj.obj.Enable = &value
	return obj
}

// True to enable notifications when the rx rate of a flow passes above
// or below the threshold value.
// Threshold returns a float32
func (obj *eventRxRateThreshold) Threshold() float32 {

	return *obj.obj.Threshold

}

// True to enable notifications when the rx rate of a flow passes above
// or below the threshold value.
// Threshold returns a float32
func (obj *eventRxRateThreshold) HasThreshold() bool {
	return obj.obj.Threshold != nil
}

// True to enable notifications when the rx rate of a flow passes above
// or below the threshold value.
// SetThreshold sets the float32 value in the EventRxRateThreshold object
func (obj *eventRxRateThreshold) SetThreshold(value float32) EventRxRateThreshold {

	obj.obj.Threshold = &value
	return obj
}

func (obj *eventRxRateThreshold) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Threshold != nil {

		if *obj.obj.Threshold < 0 || *obj.obj.Threshold > 100 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= EventRxRateThreshold.Threshold <= 100 but Got %f", *obj.obj.Threshold))
		}

	}

}

func (obj *eventRxRateThreshold) setDefault() {
	if obj.obj.Enable == nil {
		obj.SetEnable(false)
	}
	if obj.obj.Threshold == nil {
		obj.SetThreshold(95)
	}

}
