package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UpdateProtocolConfigIsisInterfaceProperty *****
type updateProtocolConfigIsisInterfaceProperty struct {
	validation
	obj          *otg.UpdateProtocolConfigIsisInterfaceProperty
	marshaller   marshalUpdateProtocolConfigIsisInterfaceProperty
	unMarshaller unMarshalUpdateProtocolConfigIsisInterfaceProperty
	metricHolder UpdateProtocolConfigIsisInterfaceMetric
}

func NewUpdateProtocolConfigIsisInterfaceProperty() UpdateProtocolConfigIsisInterfaceProperty {
	obj := updateProtocolConfigIsisInterfaceProperty{obj: &otg.UpdateProtocolConfigIsisInterfaceProperty{}}
	obj.setDefault()
	return &obj
}

func (obj *updateProtocolConfigIsisInterfaceProperty) msg() *otg.UpdateProtocolConfigIsisInterfaceProperty {
	return obj.obj
}

func (obj *updateProtocolConfigIsisInterfaceProperty) setMsg(msg *otg.UpdateProtocolConfigIsisInterfaceProperty) UpdateProtocolConfigIsisInterfaceProperty {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalupdateProtocolConfigIsisInterfaceProperty struct {
	obj *updateProtocolConfigIsisInterfaceProperty
}

type marshalUpdateProtocolConfigIsisInterfaceProperty interface {
	// ToProto marshals UpdateProtocolConfigIsisInterfaceProperty to protobuf object *otg.UpdateProtocolConfigIsisInterfaceProperty
	ToProto() (*otg.UpdateProtocolConfigIsisInterfaceProperty, error)
	// ToPbText marshals UpdateProtocolConfigIsisInterfaceProperty to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UpdateProtocolConfigIsisInterfaceProperty to YAML text
	ToYaml() (string, error)
	// ToJson marshals UpdateProtocolConfigIsisInterfaceProperty to JSON text
	ToJson() (string, error)
}

type unMarshalupdateProtocolConfigIsisInterfaceProperty struct {
	obj *updateProtocolConfigIsisInterfaceProperty
}

type unMarshalUpdateProtocolConfigIsisInterfaceProperty interface {
	// FromProto unmarshals UpdateProtocolConfigIsisInterfaceProperty from protobuf object *otg.UpdateProtocolConfigIsisInterfaceProperty
	FromProto(msg *otg.UpdateProtocolConfigIsisInterfaceProperty) (UpdateProtocolConfigIsisInterfaceProperty, error)
	// FromPbText unmarshals UpdateProtocolConfigIsisInterfaceProperty from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UpdateProtocolConfigIsisInterfaceProperty from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UpdateProtocolConfigIsisInterfaceProperty from JSON text
	FromJson(value string) error
}

func (obj *updateProtocolConfigIsisInterfaceProperty) Marshal() marshalUpdateProtocolConfigIsisInterfaceProperty {
	if obj.marshaller == nil {
		obj.marshaller = &marshalupdateProtocolConfigIsisInterfaceProperty{obj: obj}
	}
	return obj.marshaller
}

func (obj *updateProtocolConfigIsisInterfaceProperty) Unmarshal() unMarshalUpdateProtocolConfigIsisInterfaceProperty {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalupdateProtocolConfigIsisInterfaceProperty{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalupdateProtocolConfigIsisInterfaceProperty) ToProto() (*otg.UpdateProtocolConfigIsisInterfaceProperty, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalupdateProtocolConfigIsisInterfaceProperty) FromProto(msg *otg.UpdateProtocolConfigIsisInterfaceProperty) (UpdateProtocolConfigIsisInterfaceProperty, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalupdateProtocolConfigIsisInterfaceProperty) ToPbText() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisInterfaceProperty) FromPbText(value string) error {
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

func (m *marshalupdateProtocolConfigIsisInterfaceProperty) ToYaml() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisInterfaceProperty) FromYaml(value string) error {
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

func (m *marshalupdateProtocolConfigIsisInterfaceProperty) ToJson() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisInterfaceProperty) FromJson(value string) error {
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

func (obj *updateProtocolConfigIsisInterfaceProperty) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *updateProtocolConfigIsisInterfaceProperty) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *updateProtocolConfigIsisInterfaceProperty) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *updateProtocolConfigIsisInterfaceProperty) Clone() (UpdateProtocolConfigIsisInterfaceProperty, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUpdateProtocolConfigIsisInterfaceProperty()
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

func (obj *updateProtocolConfigIsisInterfaceProperty) setNil() {
	obj.metricHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// UpdateProtocolConfigIsisInterfaceProperty is a single IS-IS interface property update selected by choice.
type UpdateProtocolConfigIsisInterfaceProperty interface {
	Validation
	// msg marshals UpdateProtocolConfigIsisInterfaceProperty to protobuf object *otg.UpdateProtocolConfigIsisInterfaceProperty
	// and doesn't set defaults
	msg() *otg.UpdateProtocolConfigIsisInterfaceProperty
	// setMsg unmarshals UpdateProtocolConfigIsisInterfaceProperty from protobuf object *otg.UpdateProtocolConfigIsisInterfaceProperty
	// and doesn't set defaults
	setMsg(*otg.UpdateProtocolConfigIsisInterfaceProperty) UpdateProtocolConfigIsisInterfaceProperty
	// provides marshal interface
	Marshal() marshalUpdateProtocolConfigIsisInterfaceProperty
	// provides unmarshal interface
	Unmarshal() unMarshalUpdateProtocolConfigIsisInterfaceProperty
	// validate validates UpdateProtocolConfigIsisInterfaceProperty
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UpdateProtocolConfigIsisInterfaceProperty, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns UpdateProtocolConfigIsisInterfacePropertyChoiceEnum, set in UpdateProtocolConfigIsisInterfaceProperty
	Choice() UpdateProtocolConfigIsisInterfacePropertyChoiceEnum
	// setChoice assigns UpdateProtocolConfigIsisInterfacePropertyChoiceEnum provided by user to UpdateProtocolConfigIsisInterfaceProperty
	setChoice(value UpdateProtocolConfigIsisInterfacePropertyChoiceEnum) UpdateProtocolConfigIsisInterfaceProperty
	// HasChoice checks if Choice has been set in UpdateProtocolConfigIsisInterfaceProperty
	HasChoice() bool
	// Metric returns UpdateProtocolConfigIsisInterfaceMetric, set in UpdateProtocolConfigIsisInterfaceProperty.
	// UpdateProtocolConfigIsisInterfaceMetric is the metric cost to be applied to the IS-IS interface.
	Metric() UpdateProtocolConfigIsisInterfaceMetric
	// SetMetric assigns UpdateProtocolConfigIsisInterfaceMetric provided by user to UpdateProtocolConfigIsisInterfaceProperty.
	// UpdateProtocolConfigIsisInterfaceMetric is the metric cost to be applied to the IS-IS interface.
	SetMetric(value UpdateProtocolConfigIsisInterfaceMetric) UpdateProtocolConfigIsisInterfaceProperty
	// HasMetric checks if Metric has been set in UpdateProtocolConfigIsisInterfaceProperty
	HasMetric() bool
	setNil()
}

type UpdateProtocolConfigIsisInterfacePropertyChoiceEnum string

// Enum of Choice on UpdateProtocolConfigIsisInterfaceProperty
var UpdateProtocolConfigIsisInterfacePropertyChoice = struct {
	METRIC UpdateProtocolConfigIsisInterfacePropertyChoiceEnum
}{
	METRIC: UpdateProtocolConfigIsisInterfacePropertyChoiceEnum("metric"),
}

func (obj *updateProtocolConfigIsisInterfaceProperty) Choice() UpdateProtocolConfigIsisInterfacePropertyChoiceEnum {
	return UpdateProtocolConfigIsisInterfacePropertyChoiceEnum(obj.obj.Choice.Enum().String())
}

// The attribute to be updated on the interface.
// Choice returns a string
func (obj *updateProtocolConfigIsisInterfaceProperty) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *updateProtocolConfigIsisInterfaceProperty) setChoice(value UpdateProtocolConfigIsisInterfacePropertyChoiceEnum) UpdateProtocolConfigIsisInterfaceProperty {
	intValue, ok := otg.UpdateProtocolConfigIsisInterfaceProperty_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on UpdateProtocolConfigIsisInterfacePropertyChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.UpdateProtocolConfigIsisInterfaceProperty_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Metric = nil
	obj.metricHolder = nil

	if value == UpdateProtocolConfigIsisInterfacePropertyChoice.METRIC {
		obj.obj.Metric = NewUpdateProtocolConfigIsisInterfaceMetric().msg()
	}

	return obj
}

// The metric value to be updated on the IS-IS interface. Updating the metric on an emulated interface will flap the IS-IS session. Updating the metric on a simulated interface will maintain the IS-IS session in Up state and trigger an LSP update with the new metric value sent to neighbors.
// Metric returns a UpdateProtocolConfigIsisInterfaceMetric
func (obj *updateProtocolConfigIsisInterfaceProperty) Metric() UpdateProtocolConfigIsisInterfaceMetric {
	if obj.obj.Metric == nil {
		obj.setChoice(UpdateProtocolConfigIsisInterfacePropertyChoice.METRIC)
	}
	if obj.metricHolder == nil {
		obj.metricHolder = &updateProtocolConfigIsisInterfaceMetric{obj: obj.obj.Metric}
	}
	return obj.metricHolder
}

// The metric value to be updated on the IS-IS interface. Updating the metric on an emulated interface will flap the IS-IS session. Updating the metric on a simulated interface will maintain the IS-IS session in Up state and trigger an LSP update with the new metric value sent to neighbors.
// Metric returns a UpdateProtocolConfigIsisInterfaceMetric
func (obj *updateProtocolConfigIsisInterfaceProperty) HasMetric() bool {
	return obj.obj.Metric != nil
}

// The metric value to be updated on the IS-IS interface. Updating the metric on an emulated interface will flap the IS-IS session. Updating the metric on a simulated interface will maintain the IS-IS session in Up state and trigger an LSP update with the new metric value sent to neighbors.
// SetMetric sets the UpdateProtocolConfigIsisInterfaceMetric value in the UpdateProtocolConfigIsisInterfaceProperty object
func (obj *updateProtocolConfigIsisInterfaceProperty) SetMetric(value UpdateProtocolConfigIsisInterfaceMetric) UpdateProtocolConfigIsisInterfaceProperty {
	obj.setChoice(UpdateProtocolConfigIsisInterfacePropertyChoice.METRIC)
	obj.metricHolder = nil
	obj.obj.Metric = value.msg()

	return obj
}

func (obj *updateProtocolConfigIsisInterfaceProperty) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Metric != nil {

		obj.Metric().validateObj(vObj, set_default)
	}

}

func (obj *updateProtocolConfigIsisInterfaceProperty) setDefault() {
	var choices_set int = 0
	var choice UpdateProtocolConfigIsisInterfacePropertyChoiceEnum

	if obj.obj.Metric != nil {
		choices_set += 1
		choice = UpdateProtocolConfigIsisInterfacePropertyChoice.METRIC
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in UpdateProtocolConfigIsisInterfaceProperty")
			}
		} else {
			intVal := otg.UpdateProtocolConfigIsisInterfaceProperty_Choice_Enum_value[string(choice)]
			enumValue := otg.UpdateProtocolConfigIsisInterfaceProperty_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
