package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UpdateProtocolConfigIsisInterfaceUpdateGroup *****
type updateProtocolConfigIsisInterfaceUpdateGroup struct {
	validation
	obj                             *otg.UpdateProtocolConfigIsisInterfaceUpdateGroup
	marshaller                      marshalUpdateProtocolConfigIsisInterfaceUpdateGroup
	unMarshaller                    unMarshalUpdateProtocolConfigIsisInterfaceUpdateGroup
	trafficEngineeringMetricsHolder UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
}

func NewUpdateProtocolConfigIsisInterfaceUpdateGroup() UpdateProtocolConfigIsisInterfaceUpdateGroup {
	obj := updateProtocolConfigIsisInterfaceUpdateGroup{obj: &otg.UpdateProtocolConfigIsisInterfaceUpdateGroup{}}
	obj.setDefault()
	return &obj
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) msg() *otg.UpdateProtocolConfigIsisInterfaceUpdateGroup {
	return obj.obj
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) setMsg(msg *otg.UpdateProtocolConfigIsisInterfaceUpdateGroup) UpdateProtocolConfigIsisInterfaceUpdateGroup {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalupdateProtocolConfigIsisInterfaceUpdateGroup struct {
	obj *updateProtocolConfigIsisInterfaceUpdateGroup
}

type marshalUpdateProtocolConfigIsisInterfaceUpdateGroup interface {
	// ToProto marshals UpdateProtocolConfigIsisInterfaceUpdateGroup to protobuf object *otg.UpdateProtocolConfigIsisInterfaceUpdateGroup
	ToProto() (*otg.UpdateProtocolConfigIsisInterfaceUpdateGroup, error)
	// ToPbText marshals UpdateProtocolConfigIsisInterfaceUpdateGroup to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UpdateProtocolConfigIsisInterfaceUpdateGroup to YAML text
	ToYaml() (string, error)
	// ToJson marshals UpdateProtocolConfigIsisInterfaceUpdateGroup to JSON text
	ToJson() (string, error)
}

type unMarshalupdateProtocolConfigIsisInterfaceUpdateGroup struct {
	obj *updateProtocolConfigIsisInterfaceUpdateGroup
}

type unMarshalUpdateProtocolConfigIsisInterfaceUpdateGroup interface {
	// FromProto unmarshals UpdateProtocolConfigIsisInterfaceUpdateGroup from protobuf object *otg.UpdateProtocolConfigIsisInterfaceUpdateGroup
	FromProto(msg *otg.UpdateProtocolConfigIsisInterfaceUpdateGroup) (UpdateProtocolConfigIsisInterfaceUpdateGroup, error)
	// FromPbText unmarshals UpdateProtocolConfigIsisInterfaceUpdateGroup from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UpdateProtocolConfigIsisInterfaceUpdateGroup from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UpdateProtocolConfigIsisInterfaceUpdateGroup from JSON text
	FromJson(value string) error
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) Marshal() marshalUpdateProtocolConfigIsisInterfaceUpdateGroup {
	if obj.marshaller == nil {
		obj.marshaller = &marshalupdateProtocolConfigIsisInterfaceUpdateGroup{obj: obj}
	}
	return obj.marshaller
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) Unmarshal() unMarshalUpdateProtocolConfigIsisInterfaceUpdateGroup {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalupdateProtocolConfigIsisInterfaceUpdateGroup{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalupdateProtocolConfigIsisInterfaceUpdateGroup) ToProto() (*otg.UpdateProtocolConfigIsisInterfaceUpdateGroup, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalupdateProtocolConfigIsisInterfaceUpdateGroup) FromProto(msg *otg.UpdateProtocolConfigIsisInterfaceUpdateGroup) (UpdateProtocolConfigIsisInterfaceUpdateGroup, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalupdateProtocolConfigIsisInterfaceUpdateGroup) ToPbText() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisInterfaceUpdateGroup) FromPbText(value string) error {
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

func (m *marshalupdateProtocolConfigIsisInterfaceUpdateGroup) ToYaml() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisInterfaceUpdateGroup) FromYaml(value string) error {
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

func (m *marshalupdateProtocolConfigIsisInterfaceUpdateGroup) ToJson() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisInterfaceUpdateGroup) FromJson(value string) error {
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

func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) Clone() (UpdateProtocolConfigIsisInterfaceUpdateGroup, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUpdateProtocolConfigIsisInterfaceUpdateGroup()
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

func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) setNil() {
	obj.trafficEngineeringMetricsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// UpdateProtocolConfigIsisInterfaceUpdateGroup is an update group that applies a single attribute update to one or more IS-IS interfaces. All interfaces listed in names receive the same update selected by choice. Use multiple update groups in a single request to apply different attribute updates or different values to different subsets of interfaces.
// If the session is up but true on-the-fly update is not supported for the attribute (e.g. metric change on an emulated interface): a warning is returned indicating that the session will be disabled and re-enabled, and the updated attribute will be reflected once the session comes back up.
type UpdateProtocolConfigIsisInterfaceUpdateGroup interface {
	Validation
	// msg marshals UpdateProtocolConfigIsisInterfaceUpdateGroup to protobuf object *otg.UpdateProtocolConfigIsisInterfaceUpdateGroup
	// and doesn't set defaults
	msg() *otg.UpdateProtocolConfigIsisInterfaceUpdateGroup
	// setMsg unmarshals UpdateProtocolConfigIsisInterfaceUpdateGroup from protobuf object *otg.UpdateProtocolConfigIsisInterfaceUpdateGroup
	// and doesn't set defaults
	setMsg(*otg.UpdateProtocolConfigIsisInterfaceUpdateGroup) UpdateProtocolConfigIsisInterfaceUpdateGroup
	// provides marshal interface
	Marshal() marshalUpdateProtocolConfigIsisInterfaceUpdateGroup
	// provides unmarshal interface
	Unmarshal() unMarshalUpdateProtocolConfigIsisInterfaceUpdateGroup
	// validate validates UpdateProtocolConfigIsisInterfaceUpdateGroup
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UpdateProtocolConfigIsisInterfaceUpdateGroup, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Names returns []string, set in UpdateProtocolConfigIsisInterfaceUpdateGroup.
	Names() []string
	// SetNames assigns []string provided by user to UpdateProtocolConfigIsisInterfaceUpdateGroup
	SetNames(value []string) UpdateProtocolConfigIsisInterfaceUpdateGroup
	// Choice returns UpdateProtocolConfigIsisInterfaceUpdateGroupChoiceEnum, set in UpdateProtocolConfigIsisInterfaceUpdateGroup
	Choice() UpdateProtocolConfigIsisInterfaceUpdateGroupChoiceEnum
	// setChoice assigns UpdateProtocolConfigIsisInterfaceUpdateGroupChoiceEnum provided by user to UpdateProtocolConfigIsisInterfaceUpdateGroup
	setChoice(value UpdateProtocolConfigIsisInterfaceUpdateGroupChoiceEnum) UpdateProtocolConfigIsisInterfaceUpdateGroup
	// HasChoice checks if Choice has been set in UpdateProtocolConfigIsisInterfaceUpdateGroup
	HasChoice() bool
	// Metric returns uint32, set in UpdateProtocolConfigIsisInterfaceUpdateGroup.
	Metric() uint32
	// SetMetric assigns uint32 provided by user to UpdateProtocolConfigIsisInterfaceUpdateGroup
	SetMetric(value uint32) UpdateProtocolConfigIsisInterfaceUpdateGroup
	// HasMetric checks if Metric has been set in UpdateProtocolConfigIsisInterfaceUpdateGroup
	HasMetric() bool
	// Priority returns uint32, set in UpdateProtocolConfigIsisInterfaceUpdateGroup.
	Priority() uint32
	// SetPriority assigns uint32 provided by user to UpdateProtocolConfigIsisInterfaceUpdateGroup
	SetPriority(value uint32) UpdateProtocolConfigIsisInterfaceUpdateGroup
	// HasPriority checks if Priority has been set in UpdateProtocolConfigIsisInterfaceUpdateGroup
	HasPriority() bool
	// TrafficEngineeringMetrics returns UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics, set in UpdateProtocolConfigIsisInterfaceUpdateGroup.
	// UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics is the traffic engineering metric values to be updated on the IS-IS interface.
	TrafficEngineeringMetrics() UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
	// SetTrafficEngineeringMetrics assigns UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics provided by user to UpdateProtocolConfigIsisInterfaceUpdateGroup.
	// UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics is the traffic engineering metric values to be updated on the IS-IS interface.
	SetTrafficEngineeringMetrics(value UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) UpdateProtocolConfigIsisInterfaceUpdateGroup
	// HasTrafficEngineeringMetrics checks if TrafficEngineeringMetrics has been set in UpdateProtocolConfigIsisInterfaceUpdateGroup
	HasTrafficEngineeringMetrics() bool
	setNil()
}

// The names of the IS-IS interfaces to which the attribute update will be applied.
// Names returns a []string
func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) Names() []string {
	if obj.obj.Names == nil {
		obj.obj.Names = make([]string, 0)
	}
	return obj.obj.Names
}

// The names of the IS-IS interfaces to which the attribute update will be applied.
// SetNames sets the []string value in the UpdateProtocolConfigIsisInterfaceUpdateGroup object
func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) SetNames(value []string) UpdateProtocolConfigIsisInterfaceUpdateGroup {

	if obj.obj.Names == nil {
		obj.obj.Names = make([]string, 0)
	}
	obj.obj.Names = value

	return obj
}

type UpdateProtocolConfigIsisInterfaceUpdateGroupChoiceEnum string

// Enum of Choice on UpdateProtocolConfigIsisInterfaceUpdateGroup
var UpdateProtocolConfigIsisInterfaceUpdateGroupChoice = struct {
	METRIC                      UpdateProtocolConfigIsisInterfaceUpdateGroupChoiceEnum
	PRIORITY                    UpdateProtocolConfigIsisInterfaceUpdateGroupChoiceEnum
	TRAFFIC_ENGINEERING_METRICS UpdateProtocolConfigIsisInterfaceUpdateGroupChoiceEnum
}{
	METRIC:                      UpdateProtocolConfigIsisInterfaceUpdateGroupChoiceEnum("metric"),
	PRIORITY:                    UpdateProtocolConfigIsisInterfaceUpdateGroupChoiceEnum("priority"),
	TRAFFIC_ENGINEERING_METRICS: UpdateProtocolConfigIsisInterfaceUpdateGroupChoiceEnum("traffic_engineering_metrics"),
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) Choice() UpdateProtocolConfigIsisInterfaceUpdateGroupChoiceEnum {
	return UpdateProtocolConfigIsisInterfaceUpdateGroupChoiceEnum(obj.obj.Choice.Enum().String())
}

// The attribute to be updated on the interface.
// Choice returns a string
func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) setChoice(value UpdateProtocolConfigIsisInterfaceUpdateGroupChoiceEnum) UpdateProtocolConfigIsisInterfaceUpdateGroup {
	intValue, ok := otg.UpdateProtocolConfigIsisInterfaceUpdateGroup_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on UpdateProtocolConfigIsisInterfaceUpdateGroupChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.UpdateProtocolConfigIsisInterfaceUpdateGroup_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.TrafficEngineeringMetrics = nil
	obj.trafficEngineeringMetricsHolder = nil
	obj.obj.Priority = nil
	obj.obj.Metric = nil

	if value == UpdateProtocolConfigIsisInterfaceUpdateGroupChoice.METRIC {
		defaultValue := uint32(10)
		obj.obj.Metric = &defaultValue
	}

	if value == UpdateProtocolConfigIsisInterfaceUpdateGroupChoice.PRIORITY {
		defaultValue := uint32(64)
		obj.obj.Priority = &defaultValue
	}

	if value == UpdateProtocolConfigIsisInterfaceUpdateGroupChoice.TRAFFIC_ENGINEERING_METRICS {
		obj.obj.TrafficEngineeringMetrics = NewUpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics().msg()
	}

	return obj
}

// The metric value to be updated on the IS-IS interface. Updating the metric on an emulated interface will flap the IS-IS session. Updating the metric on a simulated interface will maintain the IS-IS session in Up state and trigger an LSP update with the new metric value sent to neighbors.
// Metric returns a uint32
func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) Metric() uint32 {

	if obj.obj.Metric == nil {
		obj.setChoice(UpdateProtocolConfigIsisInterfaceUpdateGroupChoice.METRIC)
	}

	return *obj.obj.Metric

}

// The metric value to be updated on the IS-IS interface. Updating the metric on an emulated interface will flap the IS-IS session. Updating the metric on a simulated interface will maintain the IS-IS session in Up state and trigger an LSP update with the new metric value sent to neighbors.
// Metric returns a uint32
func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) HasMetric() bool {
	return obj.obj.Metric != nil
}

// The metric value to be updated on the IS-IS interface. Updating the metric on an emulated interface will flap the IS-IS session. Updating the metric on a simulated interface will maintain the IS-IS session in Up state and trigger an LSP update with the new metric value sent to neighbors.
// SetMetric sets the uint32 value in the UpdateProtocolConfigIsisInterfaceUpdateGroup object
func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) SetMetric(value uint32) UpdateProtocolConfigIsisInterfaceUpdateGroup {
	obj.setChoice(UpdateProtocolConfigIsisInterfaceUpdateGroupChoice.METRIC)
	obj.obj.Metric = &value
	return obj
}

// The priority value to be updated on the IS-IS interface.
// Priority returns a uint32
func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) Priority() uint32 {

	if obj.obj.Priority == nil {
		obj.setChoice(UpdateProtocolConfigIsisInterfaceUpdateGroupChoice.PRIORITY)
	}

	return *obj.obj.Priority

}

// The priority value to be updated on the IS-IS interface.
// Priority returns a uint32
func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) HasPriority() bool {
	return obj.obj.Priority != nil
}

// The priority value to be updated on the IS-IS interface.
// SetPriority sets the uint32 value in the UpdateProtocolConfigIsisInterfaceUpdateGroup object
func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) SetPriority(value uint32) UpdateProtocolConfigIsisInterfaceUpdateGroup {
	obj.setChoice(UpdateProtocolConfigIsisInterfaceUpdateGroupChoice.PRIORITY)
	obj.obj.Priority = &value
	return obj
}

// The traffic engineering metric values to be updated on the IS-IS interface.
// TrafficEngineeringMetrics returns a UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) TrafficEngineeringMetrics() UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics {
	if obj.obj.TrafficEngineeringMetrics == nil {
		obj.setChoice(UpdateProtocolConfigIsisInterfaceUpdateGroupChoice.TRAFFIC_ENGINEERING_METRICS)
	}
	if obj.trafficEngineeringMetricsHolder == nil {
		obj.trafficEngineeringMetricsHolder = &updateProtocolConfigIsisInterfaceTrafficEngineeringMetrics{obj: obj.obj.TrafficEngineeringMetrics}
	}
	return obj.trafficEngineeringMetricsHolder
}

// The traffic engineering metric values to be updated on the IS-IS interface.
// TrafficEngineeringMetrics returns a UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics
func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) HasTrafficEngineeringMetrics() bool {
	return obj.obj.TrafficEngineeringMetrics != nil
}

// The traffic engineering metric values to be updated on the IS-IS interface.
// SetTrafficEngineeringMetrics sets the UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics value in the UpdateProtocolConfigIsisInterfaceUpdateGroup object
func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) SetTrafficEngineeringMetrics(value UpdateProtocolConfigIsisInterfaceTrafficEngineeringMetrics) UpdateProtocolConfigIsisInterfaceUpdateGroup {
	obj.setChoice(UpdateProtocolConfigIsisInterfaceUpdateGroupChoice.TRAFFIC_ENGINEERING_METRICS)
	obj.trafficEngineeringMetricsHolder = nil
	obj.obj.TrafficEngineeringMetrics = value.msg()

	return obj
}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Metric != nil {

		if *obj.obj.Metric > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= UpdateProtocolConfigIsisInterfaceUpdateGroup.Metric <= 16777215 but Got %d", *obj.obj.Metric))
		}

	}

	if obj.obj.Priority != nil {

		if *obj.obj.Priority > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= UpdateProtocolConfigIsisInterfaceUpdateGroup.Priority <= 255 but Got %d", *obj.obj.Priority))
		}

	}

	if obj.obj.TrafficEngineeringMetrics != nil {

		obj.TrafficEngineeringMetrics().validateObj(vObj, set_default)
	}

}

func (obj *updateProtocolConfigIsisInterfaceUpdateGroup) setDefault() {
	var choices_set int = 0
	var choice UpdateProtocolConfigIsisInterfaceUpdateGroupChoiceEnum

	if obj.obj.Metric != nil {
		choices_set += 1
		choice = UpdateProtocolConfigIsisInterfaceUpdateGroupChoice.METRIC
	}

	if obj.obj.Priority != nil {
		choices_set += 1
		choice = UpdateProtocolConfigIsisInterfaceUpdateGroupChoice.PRIORITY
	}

	if obj.obj.TrafficEngineeringMetrics != nil {
		choices_set += 1
		choice = UpdateProtocolConfigIsisInterfaceUpdateGroupChoice.TRAFFIC_ENGINEERING_METRICS
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in UpdateProtocolConfigIsisInterfaceUpdateGroup")
			}
		} else {
			intVal := otg.UpdateProtocolConfigIsisInterfaceUpdateGroup_Choice_Enum_value[string(choice)]
			enumValue := otg.UpdateProtocolConfigIsisInterfaceUpdateGroup_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.Metric == nil && choice == UpdateProtocolConfigIsisInterfaceUpdateGroupChoice.METRIC {
		obj.SetMetric(10)
	}
	if obj.obj.Priority == nil && choice == UpdateProtocolConfigIsisInterfaceUpdateGroupChoice.PRIORITY {
		obj.SetPriority(64)
	}

}
