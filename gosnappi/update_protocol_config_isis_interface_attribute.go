package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UpdateProtocolConfigIsisInterfaceAttribute *****
type updateProtocolConfigIsisInterfaceAttribute struct {
	validation
	obj          *otg.UpdateProtocolConfigIsisInterfaceAttribute
	marshaller   marshalUpdateProtocolConfigIsisInterfaceAttribute
	unMarshaller unMarshalUpdateProtocolConfigIsisInterfaceAttribute
}

func NewUpdateProtocolConfigIsisInterfaceAttribute() UpdateProtocolConfigIsisInterfaceAttribute {
	obj := updateProtocolConfigIsisInterfaceAttribute{obj: &otg.UpdateProtocolConfigIsisInterfaceAttribute{}}
	obj.setDefault()
	return &obj
}

func (obj *updateProtocolConfigIsisInterfaceAttribute) msg() *otg.UpdateProtocolConfigIsisInterfaceAttribute {
	return obj.obj
}

func (obj *updateProtocolConfigIsisInterfaceAttribute) setMsg(msg *otg.UpdateProtocolConfigIsisInterfaceAttribute) UpdateProtocolConfigIsisInterfaceAttribute {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalupdateProtocolConfigIsisInterfaceAttribute struct {
	obj *updateProtocolConfigIsisInterfaceAttribute
}

type marshalUpdateProtocolConfigIsisInterfaceAttribute interface {
	// ToProto marshals UpdateProtocolConfigIsisInterfaceAttribute to protobuf object *otg.UpdateProtocolConfigIsisInterfaceAttribute
	ToProto() (*otg.UpdateProtocolConfigIsisInterfaceAttribute, error)
	// ToPbText marshals UpdateProtocolConfigIsisInterfaceAttribute to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UpdateProtocolConfigIsisInterfaceAttribute to YAML text
	ToYaml() (string, error)
	// ToJson marshals UpdateProtocolConfigIsisInterfaceAttribute to JSON text
	ToJson() (string, error)
}

type unMarshalupdateProtocolConfigIsisInterfaceAttribute struct {
	obj *updateProtocolConfigIsisInterfaceAttribute
}

type unMarshalUpdateProtocolConfigIsisInterfaceAttribute interface {
	// FromProto unmarshals UpdateProtocolConfigIsisInterfaceAttribute from protobuf object *otg.UpdateProtocolConfigIsisInterfaceAttribute
	FromProto(msg *otg.UpdateProtocolConfigIsisInterfaceAttribute) (UpdateProtocolConfigIsisInterfaceAttribute, error)
	// FromPbText unmarshals UpdateProtocolConfigIsisInterfaceAttribute from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UpdateProtocolConfigIsisInterfaceAttribute from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UpdateProtocolConfigIsisInterfaceAttribute from JSON text
	FromJson(value string) error
}

func (obj *updateProtocolConfigIsisInterfaceAttribute) Marshal() marshalUpdateProtocolConfigIsisInterfaceAttribute {
	if obj.marshaller == nil {
		obj.marshaller = &marshalupdateProtocolConfigIsisInterfaceAttribute{obj: obj}
	}
	return obj.marshaller
}

func (obj *updateProtocolConfigIsisInterfaceAttribute) Unmarshal() unMarshalUpdateProtocolConfigIsisInterfaceAttribute {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalupdateProtocolConfigIsisInterfaceAttribute{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalupdateProtocolConfigIsisInterfaceAttribute) ToProto() (*otg.UpdateProtocolConfigIsisInterfaceAttribute, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalupdateProtocolConfigIsisInterfaceAttribute) FromProto(msg *otg.UpdateProtocolConfigIsisInterfaceAttribute) (UpdateProtocolConfigIsisInterfaceAttribute, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalupdateProtocolConfigIsisInterfaceAttribute) ToPbText() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisInterfaceAttribute) FromPbText(value string) error {
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

func (m *marshalupdateProtocolConfigIsisInterfaceAttribute) ToYaml() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisInterfaceAttribute) FromYaml(value string) error {
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

func (m *marshalupdateProtocolConfigIsisInterfaceAttribute) ToJson() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisInterfaceAttribute) FromJson(value string) error {
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

func (obj *updateProtocolConfigIsisInterfaceAttribute) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *updateProtocolConfigIsisInterfaceAttribute) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *updateProtocolConfigIsisInterfaceAttribute) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *updateProtocolConfigIsisInterfaceAttribute) Clone() (UpdateProtocolConfigIsisInterfaceAttribute, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUpdateProtocolConfigIsisInterfaceAttribute()
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

// UpdateProtocolConfigIsisInterfaceAttribute is a single IS-IS interface attribute update. The choice field identifies which attribute is being changed.
type UpdateProtocolConfigIsisInterfaceAttribute interface {
	Validation
	// msg marshals UpdateProtocolConfigIsisInterfaceAttribute to protobuf object *otg.UpdateProtocolConfigIsisInterfaceAttribute
	// and doesn't set defaults
	msg() *otg.UpdateProtocolConfigIsisInterfaceAttribute
	// setMsg unmarshals UpdateProtocolConfigIsisInterfaceAttribute from protobuf object *otg.UpdateProtocolConfigIsisInterfaceAttribute
	// and doesn't set defaults
	setMsg(*otg.UpdateProtocolConfigIsisInterfaceAttribute) UpdateProtocolConfigIsisInterfaceAttribute
	// provides marshal interface
	Marshal() marshalUpdateProtocolConfigIsisInterfaceAttribute
	// provides unmarshal interface
	Unmarshal() unMarshalUpdateProtocolConfigIsisInterfaceAttribute
	// validate validates UpdateProtocolConfigIsisInterfaceAttribute
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UpdateProtocolConfigIsisInterfaceAttribute, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns UpdateProtocolConfigIsisInterfaceAttributeChoiceEnum, set in UpdateProtocolConfigIsisInterfaceAttribute
	Choice() UpdateProtocolConfigIsisInterfaceAttributeChoiceEnum
	// setChoice assigns UpdateProtocolConfigIsisInterfaceAttributeChoiceEnum provided by user to UpdateProtocolConfigIsisInterfaceAttribute
	setChoice(value UpdateProtocolConfigIsisInterfaceAttributeChoiceEnum) UpdateProtocolConfigIsisInterfaceAttribute
	// HasChoice checks if Choice has been set in UpdateProtocolConfigIsisInterfaceAttribute
	HasChoice() bool
	// Metric returns uint32, set in UpdateProtocolConfigIsisInterfaceAttribute.
	Metric() uint32
	// SetMetric assigns uint32 provided by user to UpdateProtocolConfigIsisInterfaceAttribute
	SetMetric(value uint32) UpdateProtocolConfigIsisInterfaceAttribute
	// HasMetric checks if Metric has been set in UpdateProtocolConfigIsisInterfaceAttribute
	HasMetric() bool
}

type UpdateProtocolConfigIsisInterfaceAttributeChoiceEnum string

// Enum of Choice on UpdateProtocolConfigIsisInterfaceAttribute
var UpdateProtocolConfigIsisInterfaceAttributeChoice = struct {
	METRIC UpdateProtocolConfigIsisInterfaceAttributeChoiceEnum
}{
	METRIC: UpdateProtocolConfigIsisInterfaceAttributeChoiceEnum("metric"),
}

func (obj *updateProtocolConfigIsisInterfaceAttribute) Choice() UpdateProtocolConfigIsisInterfaceAttributeChoiceEnum {
	return UpdateProtocolConfigIsisInterfaceAttributeChoiceEnum(obj.obj.Choice.Enum().String())
}

// The interface attribute to be updated.
// Choice returns a string
func (obj *updateProtocolConfigIsisInterfaceAttribute) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *updateProtocolConfigIsisInterfaceAttribute) setChoice(value UpdateProtocolConfigIsisInterfaceAttributeChoiceEnum) UpdateProtocolConfigIsisInterfaceAttribute {
	intValue, ok := otg.UpdateProtocolConfigIsisInterfaceAttribute_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on UpdateProtocolConfigIsisInterfaceAttributeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.UpdateProtocolConfigIsisInterfaceAttribute_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Metric = nil

	if value == UpdateProtocolConfigIsisInterfaceAttributeChoice.METRIC {
		defaultValue := uint32(10)
		obj.obj.Metric = &defaultValue
	}

	return obj
}

// The metric value to be updated on the IS-IS interface. The maximum value of 16777215 (2^24-1) applies only when wide metric is enabled (RFC 5305). With narrow metrics the maximum is 63; values above 63 will be rejected. Updating the metric on an emulated interface will flap the IS-IS session. Updating the metric on a simulated interface will maintain the IS-IS session in Up state and trigger an LSP update with the new metric value sent to neighbors.
// Metric returns a uint32
func (obj *updateProtocolConfigIsisInterfaceAttribute) Metric() uint32 {

	if obj.obj.Metric == nil {
		obj.setChoice(UpdateProtocolConfigIsisInterfaceAttributeChoice.METRIC)
	}

	return *obj.obj.Metric

}

// The metric value to be updated on the IS-IS interface. The maximum value of 16777215 (2^24-1) applies only when wide metric is enabled (RFC 5305). With narrow metrics the maximum is 63; values above 63 will be rejected. Updating the metric on an emulated interface will flap the IS-IS session. Updating the metric on a simulated interface will maintain the IS-IS session in Up state and trigger an LSP update with the new metric value sent to neighbors.
// Metric returns a uint32
func (obj *updateProtocolConfigIsisInterfaceAttribute) HasMetric() bool {
	return obj.obj.Metric != nil
}

// The metric value to be updated on the IS-IS interface. The maximum value of 16777215 (2^24-1) applies only when wide metric is enabled (RFC 5305). With narrow metrics the maximum is 63; values above 63 will be rejected. Updating the metric on an emulated interface will flap the IS-IS session. Updating the metric on a simulated interface will maintain the IS-IS session in Up state and trigger an LSP update with the new metric value sent to neighbors.
// SetMetric sets the uint32 value in the UpdateProtocolConfigIsisInterfaceAttribute object
func (obj *updateProtocolConfigIsisInterfaceAttribute) SetMetric(value uint32) UpdateProtocolConfigIsisInterfaceAttribute {
	obj.setChoice(UpdateProtocolConfigIsisInterfaceAttributeChoice.METRIC)
	obj.obj.Metric = &value
	return obj
}

func (obj *updateProtocolConfigIsisInterfaceAttribute) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Metric != nil {

		if *obj.obj.Metric > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= UpdateProtocolConfigIsisInterfaceAttribute.Metric <= 16777215 but Got %d", *obj.obj.Metric))
		}

	}

}

func (obj *updateProtocolConfigIsisInterfaceAttribute) setDefault() {
	var choices_set int = 0
	var choice UpdateProtocolConfigIsisInterfaceAttributeChoiceEnum

	if obj.obj.Metric != nil {
		choices_set += 1
		choice = UpdateProtocolConfigIsisInterfaceAttributeChoice.METRIC
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in UpdateProtocolConfigIsisInterfaceAttribute")
			}
		} else {
			intVal := otg.UpdateProtocolConfigIsisInterfaceAttribute_Choice_Enum_value[string(choice)]
			enumValue := otg.UpdateProtocolConfigIsisInterfaceAttribute_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.Metric == nil && choice == UpdateProtocolConfigIsisInterfaceAttributeChoice.METRIC {
		obj.SetMetric(10)
	}

}
