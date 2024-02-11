package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRate *****
type flowRate struct {
	validation
	obj          *otg.FlowRate
	marshaller   marshalFlowRate
	unMarshaller unMarshalFlowRate
}

func NewFlowRate() FlowRate {
	obj := flowRate{obj: &otg.FlowRate{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRate) msg() *otg.FlowRate {
	return obj.obj
}

func (obj *flowRate) setMsg(msg *otg.FlowRate) FlowRate {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRate struct {
	obj *flowRate
}

type marshalFlowRate interface {
	// ToProto marshals FlowRate to protobuf object *otg.FlowRate
	ToProto() (*otg.FlowRate, error)
	// ToPbText marshals FlowRate to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRate to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRate to JSON text
	ToJson() (string, error)
}

type unMarshalflowRate struct {
	obj *flowRate
}

type unMarshalFlowRate interface {
	// FromProto unmarshals FlowRate from protobuf object *otg.FlowRate
	FromProto(msg *otg.FlowRate) (FlowRate, error)
	// FromPbText unmarshals FlowRate from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRate from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRate from JSON text
	FromJson(value string) error
}

func (obj *flowRate) Marshal() marshalFlowRate {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRate{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRate) Unmarshal() unMarshalFlowRate {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRate{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRate) ToProto() (*otg.FlowRate, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRate) FromProto(msg *otg.FlowRate) (FlowRate, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRate) ToPbText() (string, error) {
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

func (m *unMarshalflowRate) FromPbText(value string) error {
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

func (m *marshalflowRate) ToYaml() (string, error) {
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

func (m *unMarshalflowRate) FromYaml(value string) error {
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

func (m *marshalflowRate) ToJson() (string, error) {
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

func (m *unMarshalflowRate) FromJson(value string) error {
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

func (obj *flowRate) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRate) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRate) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRate) Clone() (FlowRate, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRate()
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

// FlowRate is the rate of packet transmission
type FlowRate interface {
	Validation
	// msg marshals FlowRate to protobuf object *otg.FlowRate
	// and doesn't set defaults
	msg() *otg.FlowRate
	// setMsg unmarshals FlowRate from protobuf object *otg.FlowRate
	// and doesn't set defaults
	setMsg(*otg.FlowRate) FlowRate
	// provides marshal interface
	Marshal() marshalFlowRate
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRate
	// validate validates FlowRate
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRate, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowRateChoiceEnum, set in FlowRate
	Choice() FlowRateChoiceEnum
	// setChoice assigns FlowRateChoiceEnum provided by user to FlowRate
	setChoice(value FlowRateChoiceEnum) FlowRate
	// HasChoice checks if Choice has been set in FlowRate
	HasChoice() bool
	// Pps returns uint64, set in FlowRate.
	Pps() uint64
	// SetPps assigns uint64 provided by user to FlowRate
	SetPps(value uint64) FlowRate
	// HasPps checks if Pps has been set in FlowRate
	HasPps() bool
	// Bps returns uint64, set in FlowRate.
	Bps() uint64
	// SetBps assigns uint64 provided by user to FlowRate
	SetBps(value uint64) FlowRate
	// HasBps checks if Bps has been set in FlowRate
	HasBps() bool
	// Kbps returns uint64, set in FlowRate.
	Kbps() uint64
	// SetKbps assigns uint64 provided by user to FlowRate
	SetKbps(value uint64) FlowRate
	// HasKbps checks if Kbps has been set in FlowRate
	HasKbps() bool
	// Mbps returns uint64, set in FlowRate.
	Mbps() uint64
	// SetMbps assigns uint64 provided by user to FlowRate
	SetMbps(value uint64) FlowRate
	// HasMbps checks if Mbps has been set in FlowRate
	HasMbps() bool
	// Gbps returns uint32, set in FlowRate.
	Gbps() uint32
	// SetGbps assigns uint32 provided by user to FlowRate
	SetGbps(value uint32) FlowRate
	// HasGbps checks if Gbps has been set in FlowRate
	HasGbps() bool
	// Percentage returns float32, set in FlowRate.
	Percentage() float32
	// SetPercentage assigns float32 provided by user to FlowRate
	SetPercentage(value float32) FlowRate
	// HasPercentage checks if Percentage has been set in FlowRate
	HasPercentage() bool
}

type FlowRateChoiceEnum string

// Enum of Choice on FlowRate
var FlowRateChoice = struct {
	PPS        FlowRateChoiceEnum
	BPS        FlowRateChoiceEnum
	KBPS       FlowRateChoiceEnum
	MBPS       FlowRateChoiceEnum
	GBPS       FlowRateChoiceEnum
	PERCENTAGE FlowRateChoiceEnum
}{
	PPS:        FlowRateChoiceEnum("pps"),
	BPS:        FlowRateChoiceEnum("bps"),
	KBPS:       FlowRateChoiceEnum("kbps"),
	MBPS:       FlowRateChoiceEnum("mbps"),
	GBPS:       FlowRateChoiceEnum("gbps"),
	PERCENTAGE: FlowRateChoiceEnum("percentage"),
}

func (obj *flowRate) Choice() FlowRateChoiceEnum {
	return FlowRateChoiceEnum(obj.obj.Choice.Enum().String())
}

// The available types of flow rate.
// Choice returns a string
func (obj *flowRate) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowRate) setChoice(value FlowRateChoiceEnum) FlowRate {
	intValue, ok := otg.FlowRate_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowRateChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowRate_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Percentage = nil
	obj.obj.Gbps = nil
	obj.obj.Mbps = nil
	obj.obj.Kbps = nil
	obj.obj.Bps = nil
	obj.obj.Pps = nil

	if value == FlowRateChoice.PPS {
		defaultValue := uint64(1000)
		obj.obj.Pps = &defaultValue
	}

	if value == FlowRateChoice.BPS {
		defaultValue := uint64(1000000000)
		obj.obj.Bps = &defaultValue
	}

	if value == FlowRateChoice.KBPS {
		defaultValue := uint64(1000000)
		obj.obj.Kbps = &defaultValue
	}

	if value == FlowRateChoice.MBPS {
		defaultValue := uint64(1000)
		obj.obj.Mbps = &defaultValue
	}

	if value == FlowRateChoice.GBPS {
		defaultValue := uint32(1)
		obj.obj.Gbps = &defaultValue
	}

	if value == FlowRateChoice.PERCENTAGE {
		defaultValue := float32(100)
		obj.obj.Percentage = &defaultValue
	}

	return obj
}

// Packets per second.
// Pps returns a uint64
func (obj *flowRate) Pps() uint64 {

	if obj.obj.Pps == nil {
		obj.setChoice(FlowRateChoice.PPS)
	}

	return *obj.obj.Pps

}

// Packets per second.
// Pps returns a uint64
func (obj *flowRate) HasPps() bool {
	return obj.obj.Pps != nil
}

// Packets per second.
// SetPps sets the uint64 value in the FlowRate object
func (obj *flowRate) SetPps(value uint64) FlowRate {
	obj.setChoice(FlowRateChoice.PPS)
	obj.obj.Pps = &value
	return obj
}

// Bits per second.
// Bps returns a uint64
func (obj *flowRate) Bps() uint64 {

	if obj.obj.Bps == nil {
		obj.setChoice(FlowRateChoice.BPS)
	}

	return *obj.obj.Bps

}

// Bits per second.
// Bps returns a uint64
func (obj *flowRate) HasBps() bool {
	return obj.obj.Bps != nil
}

// Bits per second.
// SetBps sets the uint64 value in the FlowRate object
func (obj *flowRate) SetBps(value uint64) FlowRate {
	obj.setChoice(FlowRateChoice.BPS)
	obj.obj.Bps = &value
	return obj
}

// Kilobits per second.
// Kbps returns a uint64
func (obj *flowRate) Kbps() uint64 {

	if obj.obj.Kbps == nil {
		obj.setChoice(FlowRateChoice.KBPS)
	}

	return *obj.obj.Kbps

}

// Kilobits per second.
// Kbps returns a uint64
func (obj *flowRate) HasKbps() bool {
	return obj.obj.Kbps != nil
}

// Kilobits per second.
// SetKbps sets the uint64 value in the FlowRate object
func (obj *flowRate) SetKbps(value uint64) FlowRate {
	obj.setChoice(FlowRateChoice.KBPS)
	obj.obj.Kbps = &value
	return obj
}

// Megabits per second.
// Mbps returns a uint64
func (obj *flowRate) Mbps() uint64 {

	if obj.obj.Mbps == nil {
		obj.setChoice(FlowRateChoice.MBPS)
	}

	return *obj.obj.Mbps

}

// Megabits per second.
// Mbps returns a uint64
func (obj *flowRate) HasMbps() bool {
	return obj.obj.Mbps != nil
}

// Megabits per second.
// SetMbps sets the uint64 value in the FlowRate object
func (obj *flowRate) SetMbps(value uint64) FlowRate {
	obj.setChoice(FlowRateChoice.MBPS)
	obj.obj.Mbps = &value
	return obj
}

// Gigabits per second.
// Gbps returns a uint32
func (obj *flowRate) Gbps() uint32 {

	if obj.obj.Gbps == nil {
		obj.setChoice(FlowRateChoice.GBPS)
	}

	return *obj.obj.Gbps

}

// Gigabits per second.
// Gbps returns a uint32
func (obj *flowRate) HasGbps() bool {
	return obj.obj.Gbps != nil
}

// Gigabits per second.
// SetGbps sets the uint32 value in the FlowRate object
func (obj *flowRate) SetGbps(value uint32) FlowRate {
	obj.setChoice(FlowRateChoice.GBPS)
	obj.obj.Gbps = &value
	return obj
}

// The percentage of a port location's available bandwidth.
// Percentage returns a float32
func (obj *flowRate) Percentage() float32 {

	if obj.obj.Percentage == nil {
		obj.setChoice(FlowRateChoice.PERCENTAGE)
	}

	return *obj.obj.Percentage

}

// The percentage of a port location's available bandwidth.
// Percentage returns a float32
func (obj *flowRate) HasPercentage() bool {
	return obj.obj.Percentage != nil
}

// The percentage of a port location's available bandwidth.
// SetPercentage sets the float32 value in the FlowRate object
func (obj *flowRate) SetPercentage(value float32) FlowRate {
	obj.setChoice(FlowRateChoice.PERCENTAGE)
	obj.obj.Percentage = &value
	return obj
}

func (obj *flowRate) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Pps != nil {

		if *obj.obj.Pps < 1 || *obj.obj.Pps > 18446744073709551615 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= FlowRate.Pps <= 18446744073709551615 but Got %d", *obj.obj.Pps))
		}

	}

	if obj.obj.Bps != nil {

		if *obj.obj.Bps < 672 || *obj.obj.Bps > 18446744073709551615 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("672 <= FlowRate.Bps <= 18446744073709551615 but Got %d", *obj.obj.Bps))
		}

	}

	if obj.obj.Kbps != nil {

		if *obj.obj.Kbps < 1 || *obj.obj.Kbps > 18446744073709551615 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= FlowRate.Kbps <= 18446744073709551615 but Got %d", *obj.obj.Kbps))
		}

	}

	if obj.obj.Mbps != nil {

		if *obj.obj.Mbps < 1 || *obj.obj.Mbps > 18446744073709551615 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= FlowRate.Mbps <= 18446744073709551615 but Got %d", *obj.obj.Mbps))
		}

	}

	if obj.obj.Gbps != nil {

		if *obj.obj.Gbps < 1 || *obj.obj.Gbps > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= FlowRate.Gbps <= 4294967295 but Got %d", *obj.obj.Gbps))
		}

	}

	if obj.obj.Percentage != nil {

		if *obj.obj.Percentage < 0 || *obj.obj.Percentage > 100 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowRate.Percentage <= 100 but Got %f", *obj.obj.Percentage))
		}

	}

}

func (obj *flowRate) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(FlowRateChoice.PPS)

	}

}
