package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UpdateProtocolConfigOspfv2InterfaceAttribute *****
type updateProtocolConfigOspfv2InterfaceAttribute struct {
	validation
	obj          *otg.UpdateProtocolConfigOspfv2InterfaceAttribute
	marshaller   marshalUpdateProtocolConfigOspfv2InterfaceAttribute
	unMarshaller unMarshalUpdateProtocolConfigOspfv2InterfaceAttribute
}

func NewUpdateProtocolConfigOspfv2InterfaceAttribute() UpdateProtocolConfigOspfv2InterfaceAttribute {
	obj := updateProtocolConfigOspfv2InterfaceAttribute{obj: &otg.UpdateProtocolConfigOspfv2InterfaceAttribute{}}
	obj.setDefault()
	return &obj
}

func (obj *updateProtocolConfigOspfv2InterfaceAttribute) msg() *otg.UpdateProtocolConfigOspfv2InterfaceAttribute {
	return obj.obj
}

func (obj *updateProtocolConfigOspfv2InterfaceAttribute) setMsg(msg *otg.UpdateProtocolConfigOspfv2InterfaceAttribute) UpdateProtocolConfigOspfv2InterfaceAttribute {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalupdateProtocolConfigOspfv2InterfaceAttribute struct {
	obj *updateProtocolConfigOspfv2InterfaceAttribute
}

type marshalUpdateProtocolConfigOspfv2InterfaceAttribute interface {
	// ToProto marshals UpdateProtocolConfigOspfv2InterfaceAttribute to protobuf object *otg.UpdateProtocolConfigOspfv2InterfaceAttribute
	ToProto() (*otg.UpdateProtocolConfigOspfv2InterfaceAttribute, error)
	// ToPbText marshals UpdateProtocolConfigOspfv2InterfaceAttribute to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UpdateProtocolConfigOspfv2InterfaceAttribute to YAML text
	ToYaml() (string, error)
	// ToJson marshals UpdateProtocolConfigOspfv2InterfaceAttribute to JSON text
	ToJson() (string, error)
}

type unMarshalupdateProtocolConfigOspfv2InterfaceAttribute struct {
	obj *updateProtocolConfigOspfv2InterfaceAttribute
}

type unMarshalUpdateProtocolConfigOspfv2InterfaceAttribute interface {
	// FromProto unmarshals UpdateProtocolConfigOspfv2InterfaceAttribute from protobuf object *otg.UpdateProtocolConfigOspfv2InterfaceAttribute
	FromProto(msg *otg.UpdateProtocolConfigOspfv2InterfaceAttribute) (UpdateProtocolConfigOspfv2InterfaceAttribute, error)
	// FromPbText unmarshals UpdateProtocolConfigOspfv2InterfaceAttribute from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UpdateProtocolConfigOspfv2InterfaceAttribute from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UpdateProtocolConfigOspfv2InterfaceAttribute from JSON text
	FromJson(value string) error
}

func (obj *updateProtocolConfigOspfv2InterfaceAttribute) Marshal() marshalUpdateProtocolConfigOspfv2InterfaceAttribute {
	if obj.marshaller == nil {
		obj.marshaller = &marshalupdateProtocolConfigOspfv2InterfaceAttribute{obj: obj}
	}
	return obj.marshaller
}

func (obj *updateProtocolConfigOspfv2InterfaceAttribute) Unmarshal() unMarshalUpdateProtocolConfigOspfv2InterfaceAttribute {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalupdateProtocolConfigOspfv2InterfaceAttribute{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalupdateProtocolConfigOspfv2InterfaceAttribute) ToProto() (*otg.UpdateProtocolConfigOspfv2InterfaceAttribute, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalupdateProtocolConfigOspfv2InterfaceAttribute) FromProto(msg *otg.UpdateProtocolConfigOspfv2InterfaceAttribute) (UpdateProtocolConfigOspfv2InterfaceAttribute, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalupdateProtocolConfigOspfv2InterfaceAttribute) ToPbText() (string, error) {
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

func (m *unMarshalupdateProtocolConfigOspfv2InterfaceAttribute) FromPbText(value string) error {
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

func (m *marshalupdateProtocolConfigOspfv2InterfaceAttribute) ToYaml() (string, error) {
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

func (m *unMarshalupdateProtocolConfigOspfv2InterfaceAttribute) FromYaml(value string) error {
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

func (m *marshalupdateProtocolConfigOspfv2InterfaceAttribute) ToJson() (string, error) {
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

func (m *unMarshalupdateProtocolConfigOspfv2InterfaceAttribute) FromJson(value string) error {
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

func (obj *updateProtocolConfigOspfv2InterfaceAttribute) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *updateProtocolConfigOspfv2InterfaceAttribute) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *updateProtocolConfigOspfv2InterfaceAttribute) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *updateProtocolConfigOspfv2InterfaceAttribute) Clone() (UpdateProtocolConfigOspfv2InterfaceAttribute, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUpdateProtocolConfigOspfv2InterfaceAttribute()
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

// UpdateProtocolConfigOspfv2InterfaceAttribute is a single OSPFv2 interface attribute update. The choice field identifies which attribute is being changed.
type UpdateProtocolConfigOspfv2InterfaceAttribute interface {
	Validation
	// msg marshals UpdateProtocolConfigOspfv2InterfaceAttribute to protobuf object *otg.UpdateProtocolConfigOspfv2InterfaceAttribute
	// and doesn't set defaults
	msg() *otg.UpdateProtocolConfigOspfv2InterfaceAttribute
	// setMsg unmarshals UpdateProtocolConfigOspfv2InterfaceAttribute from protobuf object *otg.UpdateProtocolConfigOspfv2InterfaceAttribute
	// and doesn't set defaults
	setMsg(*otg.UpdateProtocolConfigOspfv2InterfaceAttribute) UpdateProtocolConfigOspfv2InterfaceAttribute
	// provides marshal interface
	Marshal() marshalUpdateProtocolConfigOspfv2InterfaceAttribute
	// provides unmarshal interface
	Unmarshal() unMarshalUpdateProtocolConfigOspfv2InterfaceAttribute
	// validate validates UpdateProtocolConfigOspfv2InterfaceAttribute
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UpdateProtocolConfigOspfv2InterfaceAttribute, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns UpdateProtocolConfigOspfv2InterfaceAttributeChoiceEnum, set in UpdateProtocolConfigOspfv2InterfaceAttribute
	Choice() UpdateProtocolConfigOspfv2InterfaceAttributeChoiceEnum
	// setChoice assigns UpdateProtocolConfigOspfv2InterfaceAttributeChoiceEnum provided by user to UpdateProtocolConfigOspfv2InterfaceAttribute
	setChoice(value UpdateProtocolConfigOspfv2InterfaceAttributeChoiceEnum) UpdateProtocolConfigOspfv2InterfaceAttribute
	// HasChoice checks if Choice has been set in UpdateProtocolConfigOspfv2InterfaceAttribute
	HasChoice() bool
	// RoutingMetric returns uint32, set in UpdateProtocolConfigOspfv2InterfaceAttribute.
	RoutingMetric() uint32
	// SetRoutingMetric assigns uint32 provided by user to UpdateProtocolConfigOspfv2InterfaceAttribute
	SetRoutingMetric(value uint32) UpdateProtocolConfigOspfv2InterfaceAttribute
	// HasRoutingMetric checks if RoutingMetric has been set in UpdateProtocolConfigOspfv2InterfaceAttribute
	HasRoutingMetric() bool
}

type UpdateProtocolConfigOspfv2InterfaceAttributeChoiceEnum string

// Enum of Choice on UpdateProtocolConfigOspfv2InterfaceAttribute
var UpdateProtocolConfigOspfv2InterfaceAttributeChoice = struct {
	ROUTING_METRIC UpdateProtocolConfigOspfv2InterfaceAttributeChoiceEnum
}{
	ROUTING_METRIC: UpdateProtocolConfigOspfv2InterfaceAttributeChoiceEnum("routing_metric"),
}

func (obj *updateProtocolConfigOspfv2InterfaceAttribute) Choice() UpdateProtocolConfigOspfv2InterfaceAttributeChoiceEnum {
	return UpdateProtocolConfigOspfv2InterfaceAttributeChoiceEnum(obj.obj.Choice.Enum().String())
}

// The interface attribute to be updated.
// Choice returns a string
func (obj *updateProtocolConfigOspfv2InterfaceAttribute) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *updateProtocolConfigOspfv2InterfaceAttribute) setChoice(value UpdateProtocolConfigOspfv2InterfaceAttributeChoiceEnum) UpdateProtocolConfigOspfv2InterfaceAttribute {
	intValue, ok := otg.UpdateProtocolConfigOspfv2InterfaceAttribute_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on UpdateProtocolConfigOspfv2InterfaceAttributeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.UpdateProtocolConfigOspfv2InterfaceAttribute_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.RoutingMetric = nil

	if value == UpdateProtocolConfigOspfv2InterfaceAttributeChoice.ROUTING_METRIC {
		defaultValue := uint32(10)
		obj.obj.RoutingMetric = &defaultValue
	}

	return obj
}

// The routing metric (interface output cost) to be updated on the OSPFv2
// interface. It is advertised as the 16-bit TOS 0 metric of the corresponding
// link in the Router-LSA (RFC 2328 Appendix A.4.2); the interface output cost
// has a valid range of 1 to 65535 (RFC 2328 Section C.3). This applies equally
// to emulated and simulated OSPFv2 interfaces, since a simulated router
// originates its own Router-LSA carrying the same 16-bit link metric.
//
// On a simulated OSPFv2 interface (no real adjacency with the DUT; the interface
// belongs to a simulated router inside a simulated topology), updating the metric
// is a true on-the-fly operation: the OSPFv2 session remains in the Up state and
// the simulated router re-originates its Router-LSA (with an incremented LS
// Sequence Number) carrying the new metric, which neighbors receive without any
// session interruption.
//
// On an emulated OSPFv2 interface (real adjacency with the DUT), a true
// on-the-fly metric update may not be supported. In that case the implementation
// should return a warning, disable the session, re-enable it with the updated
// metric, and reflect the new value in the Router-LSA once the session comes
// back up.
// RoutingMetric returns a uint32
func (obj *updateProtocolConfigOspfv2InterfaceAttribute) RoutingMetric() uint32 {

	if obj.obj.RoutingMetric == nil {
		obj.setChoice(UpdateProtocolConfigOspfv2InterfaceAttributeChoice.ROUTING_METRIC)
	}

	return *obj.obj.RoutingMetric

}

// The routing metric (interface output cost) to be updated on the OSPFv2
// interface. It is advertised as the 16-bit TOS 0 metric of the corresponding
// link in the Router-LSA (RFC 2328 Appendix A.4.2); the interface output cost
// has a valid range of 1 to 65535 (RFC 2328 Section C.3). This applies equally
// to emulated and simulated OSPFv2 interfaces, since a simulated router
// originates its own Router-LSA carrying the same 16-bit link metric.
//
// On a simulated OSPFv2 interface (no real adjacency with the DUT; the interface
// belongs to a simulated router inside a simulated topology), updating the metric
// is a true on-the-fly operation: the OSPFv2 session remains in the Up state and
// the simulated router re-originates its Router-LSA (with an incremented LS
// Sequence Number) carrying the new metric, which neighbors receive without any
// session interruption.
//
// On an emulated OSPFv2 interface (real adjacency with the DUT), a true
// on-the-fly metric update may not be supported. In that case the implementation
// should return a warning, disable the session, re-enable it with the updated
// metric, and reflect the new value in the Router-LSA once the session comes
// back up.
// RoutingMetric returns a uint32
func (obj *updateProtocolConfigOspfv2InterfaceAttribute) HasRoutingMetric() bool {
	return obj.obj.RoutingMetric != nil
}

// The routing metric (interface output cost) to be updated on the OSPFv2
// interface. It is advertised as the 16-bit TOS 0 metric of the corresponding
// link in the Router-LSA (RFC 2328 Appendix A.4.2); the interface output cost
// has a valid range of 1 to 65535 (RFC 2328 Section C.3). This applies equally
// to emulated and simulated OSPFv2 interfaces, since a simulated router
// originates its own Router-LSA carrying the same 16-bit link metric.
//
// On a simulated OSPFv2 interface (no real adjacency with the DUT; the interface
// belongs to a simulated router inside a simulated topology), updating the metric
// is a true on-the-fly operation: the OSPFv2 session remains in the Up state and
// the simulated router re-originates its Router-LSA (with an incremented LS
// Sequence Number) carrying the new metric, which neighbors receive without any
// session interruption.
//
// On an emulated OSPFv2 interface (real adjacency with the DUT), a true
// on-the-fly metric update may not be supported. In that case the implementation
// should return a warning, disable the session, re-enable it with the updated
// metric, and reflect the new value in the Router-LSA once the session comes
// back up.
// SetRoutingMetric sets the uint32 value in the UpdateProtocolConfigOspfv2InterfaceAttribute object
func (obj *updateProtocolConfigOspfv2InterfaceAttribute) SetRoutingMetric(value uint32) UpdateProtocolConfigOspfv2InterfaceAttribute {
	obj.setChoice(UpdateProtocolConfigOspfv2InterfaceAttributeChoice.ROUTING_METRIC)
	obj.obj.RoutingMetric = &value
	return obj
}

func (obj *updateProtocolConfigOspfv2InterfaceAttribute) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RoutingMetric != nil {

		if *obj.obj.RoutingMetric > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= UpdateProtocolConfigOspfv2InterfaceAttribute.RoutingMetric <= 65535 but Got %d", *obj.obj.RoutingMetric))
		}

	}

}

func (obj *updateProtocolConfigOspfv2InterfaceAttribute) setDefault() {
	var choices_set int = 0
	var choice UpdateProtocolConfigOspfv2InterfaceAttributeChoiceEnum

	if obj.obj.RoutingMetric != nil {
		choices_set += 1
		choice = UpdateProtocolConfigOspfv2InterfaceAttributeChoice.ROUTING_METRIC
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in UpdateProtocolConfigOspfv2InterfaceAttribute")
			}
		} else {
			intVal := otg.UpdateProtocolConfigOspfv2InterfaceAttribute_Choice_Enum_value[string(choice)]
			enumValue := otg.UpdateProtocolConfigOspfv2InterfaceAttribute_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.RoutingMetric == nil && choice == UpdateProtocolConfigOspfv2InterfaceAttributeChoice.ROUTING_METRIC {
		obj.SetRoutingMetric(10)
	}

}
