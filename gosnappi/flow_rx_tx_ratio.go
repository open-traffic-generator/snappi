package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRxTxRatio *****
type flowRxTxRatio struct {
	validation
	obj           *otg.FlowRxTxRatio
	marshaller    marshalFlowRxTxRatio
	unMarshaller  unMarshalFlowRxTxRatio
	rxCountHolder FlowRxTxRatioRxCount
}

func NewFlowRxTxRatio() FlowRxTxRatio {
	obj := flowRxTxRatio{obj: &otg.FlowRxTxRatio{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRxTxRatio) msg() *otg.FlowRxTxRatio {
	return obj.obj
}

func (obj *flowRxTxRatio) setMsg(msg *otg.FlowRxTxRatio) FlowRxTxRatio {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRxTxRatio struct {
	obj *flowRxTxRatio
}

type marshalFlowRxTxRatio interface {
	// ToProto marshals FlowRxTxRatio to protobuf object *otg.FlowRxTxRatio
	ToProto() (*otg.FlowRxTxRatio, error)
	// ToPbText marshals FlowRxTxRatio to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRxTxRatio to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRxTxRatio to JSON text
	ToJson() (string, error)
}

type unMarshalflowRxTxRatio struct {
	obj *flowRxTxRatio
}

type unMarshalFlowRxTxRatio interface {
	// FromProto unmarshals FlowRxTxRatio from protobuf object *otg.FlowRxTxRatio
	FromProto(msg *otg.FlowRxTxRatio) (FlowRxTxRatio, error)
	// FromPbText unmarshals FlowRxTxRatio from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRxTxRatio from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRxTxRatio from JSON text
	FromJson(value string) error
}

func (obj *flowRxTxRatio) Marshal() marshalFlowRxTxRatio {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRxTxRatio{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRxTxRatio) Unmarshal() unMarshalFlowRxTxRatio {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRxTxRatio{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRxTxRatio) ToProto() (*otg.FlowRxTxRatio, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRxTxRatio) FromProto(msg *otg.FlowRxTxRatio) (FlowRxTxRatio, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRxTxRatio) ToPbText() (string, error) {
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

func (m *unMarshalflowRxTxRatio) FromPbText(value string) error {
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

func (m *marshalflowRxTxRatio) ToYaml() (string, error) {
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

func (m *unMarshalflowRxTxRatio) FromYaml(value string) error {
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

func (m *marshalflowRxTxRatio) ToJson() (string, error) {
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

func (m *unMarshalflowRxTxRatio) FromJson(value string) error {
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

func (obj *flowRxTxRatio) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRxTxRatio) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRxTxRatio) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRxTxRatio) Clone() (FlowRxTxRatio, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRxTxRatio()
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

func (obj *flowRxTxRatio) setNil() {
	obj.rxCountHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRxTxRatio is rx Tx ratio is the ratio of expected number of Rx packets across all Rx ports to the Tx packets
// for the configured flow. It is a factor by which the Tx packet count is multiplied to calculate
// the sum of expected Rx packet count, across all Rx ports. This will be used to calculate loss
// percentage of flow at aggregate level.
type FlowRxTxRatio interface {
	Validation
	// msg marshals FlowRxTxRatio to protobuf object *otg.FlowRxTxRatio
	// and doesn't set defaults
	msg() *otg.FlowRxTxRatio
	// setMsg unmarshals FlowRxTxRatio from protobuf object *otg.FlowRxTxRatio
	// and doesn't set defaults
	setMsg(*otg.FlowRxTxRatio) FlowRxTxRatio
	// provides marshal interface
	Marshal() marshalFlowRxTxRatio
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRxTxRatio
	// validate validates FlowRxTxRatio
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRxTxRatio, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowRxTxRatioChoiceEnum, set in FlowRxTxRatio
	Choice() FlowRxTxRatioChoiceEnum
	// setChoice assigns FlowRxTxRatioChoiceEnum provided by user to FlowRxTxRatio
	setChoice(value FlowRxTxRatioChoiceEnum) FlowRxTxRatio
	// HasChoice checks if Choice has been set in FlowRxTxRatio
	HasChoice() bool
	// RxCount returns FlowRxTxRatioRxCount, set in FlowRxTxRatio.
	// FlowRxTxRatioRxCount is this is for cases where one copy of Tx packet is received on all Rx ports and so the sum total of Rx packets
	// received across all Rx ports is a multiple of Rx port count and Tx packets.
	RxCount() FlowRxTxRatioRxCount
	// SetRxCount assigns FlowRxTxRatioRxCount provided by user to FlowRxTxRatio.
	// FlowRxTxRatioRxCount is this is for cases where one copy of Tx packet is received on all Rx ports and so the sum total of Rx packets
	// received across all Rx ports is a multiple of Rx port count and Tx packets.
	SetRxCount(value FlowRxTxRatioRxCount) FlowRxTxRatio
	// HasRxCount checks if RxCount has been set in FlowRxTxRatio
	HasRxCount() bool
	// Value returns float32, set in FlowRxTxRatio.
	Value() float32
	// SetValue assigns float32 provided by user to FlowRxTxRatio
	SetValue(value float32) FlowRxTxRatio
	// HasValue checks if Value has been set in FlowRxTxRatio
	HasValue() bool
	setNil()
}

type FlowRxTxRatioChoiceEnum string

// Enum of Choice on FlowRxTxRatio
var FlowRxTxRatioChoice = struct {
	RX_COUNT FlowRxTxRatioChoiceEnum
	VALUE    FlowRxTxRatioChoiceEnum
}{
	RX_COUNT: FlowRxTxRatioChoiceEnum("rx_count"),
	VALUE:    FlowRxTxRatioChoiceEnum("value"),
}

func (obj *flowRxTxRatio) Choice() FlowRxTxRatioChoiceEnum {
	return FlowRxTxRatioChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowRxTxRatio) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowRxTxRatio) setChoice(value FlowRxTxRatioChoiceEnum) FlowRxTxRatio {
	intValue, ok := otg.FlowRxTxRatio_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowRxTxRatioChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowRxTxRatio_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Value = nil
	obj.obj.RxCount = nil
	obj.rxCountHolder = nil

	if value == FlowRxTxRatioChoice.RX_COUNT {
		obj.obj.RxCount = NewFlowRxTxRatioRxCount().msg()
	}

	if value == FlowRxTxRatioChoice.VALUE {
		defaultValue := float32(1.0)
		obj.obj.Value = &defaultValue
	}

	return obj
}

// description is TBD
// RxCount returns a FlowRxTxRatioRxCount
func (obj *flowRxTxRatio) RxCount() FlowRxTxRatioRxCount {
	if obj.obj.RxCount == nil {
		obj.setChoice(FlowRxTxRatioChoice.RX_COUNT)
	}
	if obj.rxCountHolder == nil {
		obj.rxCountHolder = &flowRxTxRatioRxCount{obj: obj.obj.RxCount}
	}
	return obj.rxCountHolder
}

// description is TBD
// RxCount returns a FlowRxTxRatioRxCount
func (obj *flowRxTxRatio) HasRxCount() bool {
	return obj.obj.RxCount != nil
}

// description is TBD
// SetRxCount sets the FlowRxTxRatioRxCount value in the FlowRxTxRatio object
func (obj *flowRxTxRatio) SetRxCount(value FlowRxTxRatioRxCount) FlowRxTxRatio {
	obj.setChoice(FlowRxTxRatioChoice.RX_COUNT)
	obj.rxCountHolder = nil
	obj.obj.RxCount = value.msg()

	return obj
}

// Should be a positive, non-zero value. The default value of 1, is when the Rx packet count across
// all ports is expected to match the Tx packet count. A custom integer value (>1) can be specified for
// loss calculation for cases when there are multiple destination addresses configured within one flow,
// but DUT is configured to replicate only to a subset of Rx ports. For cases when Tx side generates two
// packets from each source in 1:1 protection mode but only one of the two packets are received by the
// Rx port, we may need to specify a fractional value instead.
// Value returns a float32
func (obj *flowRxTxRatio) Value() float32 {

	if obj.obj.Value == nil {
		obj.setChoice(FlowRxTxRatioChoice.VALUE)
	}

	return *obj.obj.Value

}

// Should be a positive, non-zero value. The default value of 1, is when the Rx packet count across
// all ports is expected to match the Tx packet count. A custom integer value (>1) can be specified for
// loss calculation for cases when there are multiple destination addresses configured within one flow,
// but DUT is configured to replicate only to a subset of Rx ports. For cases when Tx side generates two
// packets from each source in 1:1 protection mode but only one of the two packets are received by the
// Rx port, we may need to specify a fractional value instead.
// Value returns a float32
func (obj *flowRxTxRatio) HasValue() bool {
	return obj.obj.Value != nil
}

// Should be a positive, non-zero value. The default value of 1, is when the Rx packet count across
// all ports is expected to match the Tx packet count. A custom integer value (>1) can be specified for
// loss calculation for cases when there are multiple destination addresses configured within one flow,
// but DUT is configured to replicate only to a subset of Rx ports. For cases when Tx side generates two
// packets from each source in 1:1 protection mode but only one of the two packets are received by the
// Rx port, we may need to specify a fractional value instead.
// SetValue sets the float32 value in the FlowRxTxRatio object
func (obj *flowRxTxRatio) SetValue(value float32) FlowRxTxRatio {
	obj.setChoice(FlowRxTxRatioChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

func (obj *flowRxTxRatio) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RxCount != nil {

		obj.RxCount().validateObj(vObj, set_default)
	}

}

func (obj *flowRxTxRatio) setDefault() {
	var choices_set int = 0
	var choice FlowRxTxRatioChoiceEnum

	if obj.obj.RxCount != nil {
		choices_set += 1
		choice = FlowRxTxRatioChoice.RX_COUNT
	}

	if obj.obj.Value != nil {
		choices_set += 1
		choice = FlowRxTxRatioChoice.VALUE
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowRxTxRatioChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowRxTxRatio")
			}
		} else {
			intVal := otg.FlowRxTxRatio_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowRxTxRatio_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
