package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowDuration *****
type flowDuration struct {
	validation
	obj                *otg.FlowDuration
	marshaller         marshalFlowDuration
	unMarshaller       unMarshalFlowDuration
	fixedPacketsHolder FlowFixedPackets
	fixedSecondsHolder FlowFixedSeconds
	burstHolder        FlowBurst
	continuousHolder   FlowContinuous
}

func NewFlowDuration() FlowDuration {
	obj := flowDuration{obj: &otg.FlowDuration{}}
	obj.setDefault()
	return &obj
}

func (obj *flowDuration) msg() *otg.FlowDuration {
	return obj.obj
}

func (obj *flowDuration) setMsg(msg *otg.FlowDuration) FlowDuration {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowDuration struct {
	obj *flowDuration
}

type marshalFlowDuration interface {
	// ToProto marshals FlowDuration to protobuf object *otg.FlowDuration
	ToProto() (*otg.FlowDuration, error)
	// ToPbText marshals FlowDuration to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowDuration to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowDuration to JSON text
	ToJson() (string, error)
}

type unMarshalflowDuration struct {
	obj *flowDuration
}

type unMarshalFlowDuration interface {
	// FromProto unmarshals FlowDuration from protobuf object *otg.FlowDuration
	FromProto(msg *otg.FlowDuration) (FlowDuration, error)
	// FromPbText unmarshals FlowDuration from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowDuration from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowDuration from JSON text
	FromJson(value string) error
}

func (obj *flowDuration) Marshal() marshalFlowDuration {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowDuration{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowDuration) Unmarshal() unMarshalFlowDuration {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowDuration{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowDuration) ToProto() (*otg.FlowDuration, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowDuration) FromProto(msg *otg.FlowDuration) (FlowDuration, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowDuration) ToPbText() (string, error) {
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

func (m *unMarshalflowDuration) FromPbText(value string) error {
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

func (m *marshalflowDuration) ToYaml() (string, error) {
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

func (m *unMarshalflowDuration) FromYaml(value string) error {
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

func (m *marshalflowDuration) ToJson() (string, error) {
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

func (m *unMarshalflowDuration) FromJson(value string) error {
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

func (obj *flowDuration) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowDuration) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowDuration) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowDuration) Clone() (FlowDuration, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowDuration()
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

func (obj *flowDuration) setNil() {
	obj.fixedPacketsHolder = nil
	obj.fixedSecondsHolder = nil
	obj.burstHolder = nil
	obj.continuousHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowDuration is a container for different transmit durations.
type FlowDuration interface {
	Validation
	// msg marshals FlowDuration to protobuf object *otg.FlowDuration
	// and doesn't set defaults
	msg() *otg.FlowDuration
	// setMsg unmarshals FlowDuration from protobuf object *otg.FlowDuration
	// and doesn't set defaults
	setMsg(*otg.FlowDuration) FlowDuration
	// provides marshal interface
	Marshal() marshalFlowDuration
	// provides unmarshal interface
	Unmarshal() unMarshalFlowDuration
	// validate validates FlowDuration
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowDuration, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowDurationChoiceEnum, set in FlowDuration
	Choice() FlowDurationChoiceEnum
	// setChoice assigns FlowDurationChoiceEnum provided by user to FlowDuration
	setChoice(value FlowDurationChoiceEnum) FlowDuration
	// HasChoice checks if Choice has been set in FlowDuration
	HasChoice() bool
	// FixedPackets returns FlowFixedPackets, set in FlowDuration.
	// FlowFixedPackets is transmit a fixed number of packets after which the flow will stop.
	FixedPackets() FlowFixedPackets
	// SetFixedPackets assigns FlowFixedPackets provided by user to FlowDuration.
	// FlowFixedPackets is transmit a fixed number of packets after which the flow will stop.
	SetFixedPackets(value FlowFixedPackets) FlowDuration
	// HasFixedPackets checks if FixedPackets has been set in FlowDuration
	HasFixedPackets() bool
	// FixedSeconds returns FlowFixedSeconds, set in FlowDuration.
	// FlowFixedSeconds is transmit for a fixed number of seconds after which the flow will stop.
	FixedSeconds() FlowFixedSeconds
	// SetFixedSeconds assigns FlowFixedSeconds provided by user to FlowDuration.
	// FlowFixedSeconds is transmit for a fixed number of seconds after which the flow will stop.
	SetFixedSeconds(value FlowFixedSeconds) FlowDuration
	// HasFixedSeconds checks if FixedSeconds has been set in FlowDuration
	HasFixedSeconds() bool
	// Burst returns FlowBurst, set in FlowDuration.
	// FlowBurst is transmits continuous or fixed burst of packets.
	// For continuous burst of packets, it will not automatically stop.
	// For fixed burst of packets, it will stop after transmitting fixed number of bursts.
	Burst() FlowBurst
	// SetBurst assigns FlowBurst provided by user to FlowDuration.
	// FlowBurst is transmits continuous or fixed burst of packets.
	// For continuous burst of packets, it will not automatically stop.
	// For fixed burst of packets, it will stop after transmitting fixed number of bursts.
	SetBurst(value FlowBurst) FlowDuration
	// HasBurst checks if Burst has been set in FlowDuration
	HasBurst() bool
	// Continuous returns FlowContinuous, set in FlowDuration.
	// FlowContinuous is transmit will be continuous and will not stop automatically.
	Continuous() FlowContinuous
	// SetContinuous assigns FlowContinuous provided by user to FlowDuration.
	// FlowContinuous is transmit will be continuous and will not stop automatically.
	SetContinuous(value FlowContinuous) FlowDuration
	// HasContinuous checks if Continuous has been set in FlowDuration
	HasContinuous() bool
	setNil()
}

type FlowDurationChoiceEnum string

// Enum of Choice on FlowDuration
var FlowDurationChoice = struct {
	FIXED_PACKETS FlowDurationChoiceEnum
	FIXED_SECONDS FlowDurationChoiceEnum
	BURST         FlowDurationChoiceEnum
	CONTINUOUS    FlowDurationChoiceEnum
}{
	FIXED_PACKETS: FlowDurationChoiceEnum("fixed_packets"),
	FIXED_SECONDS: FlowDurationChoiceEnum("fixed_seconds"),
	BURST:         FlowDurationChoiceEnum("burst"),
	CONTINUOUS:    FlowDurationChoiceEnum("continuous"),
}

func (obj *flowDuration) Choice() FlowDurationChoiceEnum {
	return FlowDurationChoiceEnum(obj.obj.Choice.Enum().String())
}

// A choice used to determine the type of duration.
// Choice returns a string
func (obj *flowDuration) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowDuration) setChoice(value FlowDurationChoiceEnum) FlowDuration {
	intValue, ok := otg.FlowDuration_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowDurationChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowDuration_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Continuous = nil
	obj.continuousHolder = nil
	obj.obj.Burst = nil
	obj.burstHolder = nil
	obj.obj.FixedSeconds = nil
	obj.fixedSecondsHolder = nil
	obj.obj.FixedPackets = nil
	obj.fixedPacketsHolder = nil

	if value == FlowDurationChoice.FIXED_PACKETS {
		obj.obj.FixedPackets = NewFlowFixedPackets().msg()
	}

	if value == FlowDurationChoice.FIXED_SECONDS {
		obj.obj.FixedSeconds = NewFlowFixedSeconds().msg()
	}

	if value == FlowDurationChoice.BURST {
		obj.obj.Burst = NewFlowBurst().msg()
	}

	if value == FlowDurationChoice.CONTINUOUS {
		obj.obj.Continuous = NewFlowContinuous().msg()
	}

	return obj
}

// description is TBD
// FixedPackets returns a FlowFixedPackets
func (obj *flowDuration) FixedPackets() FlowFixedPackets {
	if obj.obj.FixedPackets == nil {
		obj.setChoice(FlowDurationChoice.FIXED_PACKETS)
	}
	if obj.fixedPacketsHolder == nil {
		obj.fixedPacketsHolder = &flowFixedPackets{obj: obj.obj.FixedPackets}
	}
	return obj.fixedPacketsHolder
}

// description is TBD
// FixedPackets returns a FlowFixedPackets
func (obj *flowDuration) HasFixedPackets() bool {
	return obj.obj.FixedPackets != nil
}

// description is TBD
// SetFixedPackets sets the FlowFixedPackets value in the FlowDuration object
func (obj *flowDuration) SetFixedPackets(value FlowFixedPackets) FlowDuration {
	obj.setChoice(FlowDurationChoice.FIXED_PACKETS)
	obj.fixedPacketsHolder = nil
	obj.obj.FixedPackets = value.msg()

	return obj
}

// description is TBD
// FixedSeconds returns a FlowFixedSeconds
func (obj *flowDuration) FixedSeconds() FlowFixedSeconds {
	if obj.obj.FixedSeconds == nil {
		obj.setChoice(FlowDurationChoice.FIXED_SECONDS)
	}
	if obj.fixedSecondsHolder == nil {
		obj.fixedSecondsHolder = &flowFixedSeconds{obj: obj.obj.FixedSeconds}
	}
	return obj.fixedSecondsHolder
}

// description is TBD
// FixedSeconds returns a FlowFixedSeconds
func (obj *flowDuration) HasFixedSeconds() bool {
	return obj.obj.FixedSeconds != nil
}

// description is TBD
// SetFixedSeconds sets the FlowFixedSeconds value in the FlowDuration object
func (obj *flowDuration) SetFixedSeconds(value FlowFixedSeconds) FlowDuration {
	obj.setChoice(FlowDurationChoice.FIXED_SECONDS)
	obj.fixedSecondsHolder = nil
	obj.obj.FixedSeconds = value.msg()

	return obj
}

// description is TBD
// Burst returns a FlowBurst
func (obj *flowDuration) Burst() FlowBurst {
	if obj.obj.Burst == nil {
		obj.setChoice(FlowDurationChoice.BURST)
	}
	if obj.burstHolder == nil {
		obj.burstHolder = &flowBurst{obj: obj.obj.Burst}
	}
	return obj.burstHolder
}

// description is TBD
// Burst returns a FlowBurst
func (obj *flowDuration) HasBurst() bool {
	return obj.obj.Burst != nil
}

// description is TBD
// SetBurst sets the FlowBurst value in the FlowDuration object
func (obj *flowDuration) SetBurst(value FlowBurst) FlowDuration {
	obj.setChoice(FlowDurationChoice.BURST)
	obj.burstHolder = nil
	obj.obj.Burst = value.msg()

	return obj
}

// description is TBD
// Continuous returns a FlowContinuous
func (obj *flowDuration) Continuous() FlowContinuous {
	if obj.obj.Continuous == nil {
		obj.setChoice(FlowDurationChoice.CONTINUOUS)
	}
	if obj.continuousHolder == nil {
		obj.continuousHolder = &flowContinuous{obj: obj.obj.Continuous}
	}
	return obj.continuousHolder
}

// description is TBD
// Continuous returns a FlowContinuous
func (obj *flowDuration) HasContinuous() bool {
	return obj.obj.Continuous != nil
}

// description is TBD
// SetContinuous sets the FlowContinuous value in the FlowDuration object
func (obj *flowDuration) SetContinuous(value FlowContinuous) FlowDuration {
	obj.setChoice(FlowDurationChoice.CONTINUOUS)
	obj.continuousHolder = nil
	obj.obj.Continuous = value.msg()

	return obj
}

func (obj *flowDuration) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.FixedPackets != nil {

		obj.FixedPackets().validateObj(vObj, set_default)
	}

	if obj.obj.FixedSeconds != nil {

		obj.FixedSeconds().validateObj(vObj, set_default)
	}

	if obj.obj.Burst != nil {

		obj.Burst().validateObj(vObj, set_default)
	}

	if obj.obj.Continuous != nil {

		obj.Continuous().validateObj(vObj, set_default)
	}

}

func (obj *flowDuration) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(FlowDurationChoice.CONTINUOUS)

	}

}
