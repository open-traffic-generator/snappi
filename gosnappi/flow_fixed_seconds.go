package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowFixedSeconds *****
type flowFixedSeconds struct {
	validation
	obj          *otg.FlowFixedSeconds
	marshaller   marshalFlowFixedSeconds
	unMarshaller unMarshalFlowFixedSeconds
	delayHolder  FlowDelay
}

func NewFlowFixedSeconds() FlowFixedSeconds {
	obj := flowFixedSeconds{obj: &otg.FlowFixedSeconds{}}
	obj.setDefault()
	return &obj
}

func (obj *flowFixedSeconds) msg() *otg.FlowFixedSeconds {
	return obj.obj
}

func (obj *flowFixedSeconds) setMsg(msg *otg.FlowFixedSeconds) FlowFixedSeconds {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowFixedSeconds struct {
	obj *flowFixedSeconds
}

type marshalFlowFixedSeconds interface {
	// ToProto marshals FlowFixedSeconds to protobuf object *otg.FlowFixedSeconds
	ToProto() (*otg.FlowFixedSeconds, error)
	// ToPbText marshals FlowFixedSeconds to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowFixedSeconds to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowFixedSeconds to JSON text
	ToJson() (string, error)
}

type unMarshalflowFixedSeconds struct {
	obj *flowFixedSeconds
}

type unMarshalFlowFixedSeconds interface {
	// FromProto unmarshals FlowFixedSeconds from protobuf object *otg.FlowFixedSeconds
	FromProto(msg *otg.FlowFixedSeconds) (FlowFixedSeconds, error)
	// FromPbText unmarshals FlowFixedSeconds from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowFixedSeconds from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowFixedSeconds from JSON text
	FromJson(value string) error
}

func (obj *flowFixedSeconds) Marshal() marshalFlowFixedSeconds {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowFixedSeconds{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowFixedSeconds) Unmarshal() unMarshalFlowFixedSeconds {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowFixedSeconds{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowFixedSeconds) ToProto() (*otg.FlowFixedSeconds, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowFixedSeconds) FromProto(msg *otg.FlowFixedSeconds) (FlowFixedSeconds, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowFixedSeconds) ToPbText() (string, error) {
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

func (m *unMarshalflowFixedSeconds) FromPbText(value string) error {
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

func (m *marshalflowFixedSeconds) ToYaml() (string, error) {
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

func (m *unMarshalflowFixedSeconds) FromYaml(value string) error {
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

func (m *marshalflowFixedSeconds) ToJson() (string, error) {
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

func (m *unMarshalflowFixedSeconds) FromJson(value string) error {
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

func (obj *flowFixedSeconds) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowFixedSeconds) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowFixedSeconds) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowFixedSeconds) Clone() (FlowFixedSeconds, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowFixedSeconds()
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

func (obj *flowFixedSeconds) setNil() {
	obj.delayHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowFixedSeconds is transmit for a fixed number of seconds after which the flow will stop.
type FlowFixedSeconds interface {
	Validation
	// msg marshals FlowFixedSeconds to protobuf object *otg.FlowFixedSeconds
	// and doesn't set defaults
	msg() *otg.FlowFixedSeconds
	// setMsg unmarshals FlowFixedSeconds from protobuf object *otg.FlowFixedSeconds
	// and doesn't set defaults
	setMsg(*otg.FlowFixedSeconds) FlowFixedSeconds
	// provides marshal interface
	Marshal() marshalFlowFixedSeconds
	// provides unmarshal interface
	Unmarshal() unMarshalFlowFixedSeconds
	// validate validates FlowFixedSeconds
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowFixedSeconds, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Seconds returns float32, set in FlowFixedSeconds.
	Seconds() float32
	// SetSeconds assigns float32 provided by user to FlowFixedSeconds
	SetSeconds(value float32) FlowFixedSeconds
	// HasSeconds checks if Seconds has been set in FlowFixedSeconds
	HasSeconds() bool
	// Gap returns uint32, set in FlowFixedSeconds.
	Gap() uint32
	// SetGap assigns uint32 provided by user to FlowFixedSeconds
	SetGap(value uint32) FlowFixedSeconds
	// HasGap checks if Gap has been set in FlowFixedSeconds
	HasGap() bool
	// Delay returns FlowDelay, set in FlowFixedSeconds.
	// FlowDelay is the optional container to specify the delay before starting
	// transmission of packets.
	Delay() FlowDelay
	// SetDelay assigns FlowDelay provided by user to FlowFixedSeconds.
	// FlowDelay is the optional container to specify the delay before starting
	// transmission of packets.
	SetDelay(value FlowDelay) FlowFixedSeconds
	// HasDelay checks if Delay has been set in FlowFixedSeconds
	HasDelay() bool
	setNil()
}

// Stop transmit of the flow after this number of seconds.
// Seconds returns a float32
func (obj *flowFixedSeconds) Seconds() float32 {

	return *obj.obj.Seconds

}

// Stop transmit of the flow after this number of seconds.
// Seconds returns a float32
func (obj *flowFixedSeconds) HasSeconds() bool {
	return obj.obj.Seconds != nil
}

// Stop transmit of the flow after this number of seconds.
// SetSeconds sets the float32 value in the FlowFixedSeconds object
func (obj *flowFixedSeconds) SetSeconds(value float32) FlowFixedSeconds {

	obj.obj.Seconds = &value
	return obj
}

// The minimum gap between packets expressed as bytes.
// Gap returns a uint32
func (obj *flowFixedSeconds) Gap() uint32 {

	return *obj.obj.Gap

}

// The minimum gap between packets expressed as bytes.
// Gap returns a uint32
func (obj *flowFixedSeconds) HasGap() bool {
	return obj.obj.Gap != nil
}

// The minimum gap between packets expressed as bytes.
// SetGap sets the uint32 value in the FlowFixedSeconds object
func (obj *flowFixedSeconds) SetGap(value uint32) FlowFixedSeconds {

	obj.obj.Gap = &value
	return obj
}

// description is TBD
// Delay returns a FlowDelay
func (obj *flowFixedSeconds) Delay() FlowDelay {
	if obj.obj.Delay == nil {
		obj.obj.Delay = NewFlowDelay().msg()
	}
	if obj.delayHolder == nil {
		obj.delayHolder = &flowDelay{obj: obj.obj.Delay}
	}
	return obj.delayHolder
}

// description is TBD
// Delay returns a FlowDelay
func (obj *flowFixedSeconds) HasDelay() bool {
	return obj.obj.Delay != nil
}

// description is TBD
// SetDelay sets the FlowDelay value in the FlowFixedSeconds object
func (obj *flowFixedSeconds) SetDelay(value FlowDelay) FlowFixedSeconds {

	obj.delayHolder = nil
	obj.obj.Delay = value.msg()

	return obj
}

func (obj *flowFixedSeconds) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Seconds != nil {

		if *obj.obj.Seconds < 0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowFixedSeconds.Seconds <= max(float32) but Got %f", *obj.obj.Seconds))
		}

	}

	if obj.obj.Delay != nil {

		obj.Delay().validateObj(vObj, set_default)
	}

}

func (obj *flowFixedSeconds) setDefault() {
	if obj.obj.Seconds == nil {
		obj.SetSeconds(1)
	}
	if obj.obj.Gap == nil {
		obj.SetGap(12)
	}

}
