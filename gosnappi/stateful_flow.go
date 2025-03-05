package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StatefulFlow *****
type statefulFlow struct {
	validation
	obj               *otg.StatefulFlow
	marshaller        marshalStatefulFlow
	unMarshaller      unMarshalStatefulFlow
	rocev2FlowsHolder StatefulFlowRoCEv2FlowIter
}

func NewStatefulFlow() StatefulFlow {
	obj := statefulFlow{obj: &otg.StatefulFlow{}}
	obj.setDefault()
	return &obj
}

func (obj *statefulFlow) msg() *otg.StatefulFlow {
	return obj.obj
}

func (obj *statefulFlow) setMsg(msg *otg.StatefulFlow) StatefulFlow {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstatefulFlow struct {
	obj *statefulFlow
}

type marshalStatefulFlow interface {
	// ToProto marshals StatefulFlow to protobuf object *otg.StatefulFlow
	ToProto() (*otg.StatefulFlow, error)
	// ToPbText marshals StatefulFlow to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StatefulFlow to YAML text
	ToYaml() (string, error)
	// ToJson marshals StatefulFlow to JSON text
	ToJson() (string, error)
}

type unMarshalstatefulFlow struct {
	obj *statefulFlow
}

type unMarshalStatefulFlow interface {
	// FromProto unmarshals StatefulFlow from protobuf object *otg.StatefulFlow
	FromProto(msg *otg.StatefulFlow) (StatefulFlow, error)
	// FromPbText unmarshals StatefulFlow from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StatefulFlow from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StatefulFlow from JSON text
	FromJson(value string) error
}

func (obj *statefulFlow) Marshal() marshalStatefulFlow {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstatefulFlow{obj: obj}
	}
	return obj.marshaller
}

func (obj *statefulFlow) Unmarshal() unMarshalStatefulFlow {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstatefulFlow{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstatefulFlow) ToProto() (*otg.StatefulFlow, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstatefulFlow) FromProto(msg *otg.StatefulFlow) (StatefulFlow, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstatefulFlow) ToPbText() (string, error) {
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

func (m *unMarshalstatefulFlow) FromPbText(value string) error {
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

func (m *marshalstatefulFlow) ToYaml() (string, error) {
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

func (m *unMarshalstatefulFlow) FromYaml(value string) error {
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

func (m *marshalstatefulFlow) ToJson() (string, error) {
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

func (m *unMarshalstatefulFlow) FromJson(value string) error {
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

func (obj *statefulFlow) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *statefulFlow) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *statefulFlow) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *statefulFlow) Clone() (StatefulFlow, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStatefulFlow()
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

func (obj *statefulFlow) setNil() {
	obj.rocev2FlowsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// StatefulFlow is a high level data plane traffic flow.
type StatefulFlow interface {
	Validation
	// msg marshals StatefulFlow to protobuf object *otg.StatefulFlow
	// and doesn't set defaults
	msg() *otg.StatefulFlow
	// setMsg unmarshals StatefulFlow from protobuf object *otg.StatefulFlow
	// and doesn't set defaults
	setMsg(*otg.StatefulFlow) StatefulFlow
	// provides marshal interface
	Marshal() marshalStatefulFlow
	// provides unmarshal interface
	Unmarshal() unMarshalStatefulFlow
	// validate validates StatefulFlow
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StatefulFlow, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns StatefulFlowChoiceEnum, set in StatefulFlow
	Choice() StatefulFlowChoiceEnum
	// setChoice assigns StatefulFlowChoiceEnum provided by user to StatefulFlow
	setChoice(value StatefulFlowChoiceEnum) StatefulFlow
	// HasChoice checks if Choice has been set in StatefulFlow
	HasChoice() bool
	// getter for Rocev2 to set choice.
	Rocev2()
	// Rocev2Flows returns StatefulFlowRoCEv2FlowIterIter, set in StatefulFlow
	Rocev2Flows() StatefulFlowRoCEv2FlowIter
	setNil()
}

type StatefulFlowChoiceEnum string

// Enum of Choice on StatefulFlow
var StatefulFlowChoice = struct {
	ROCEV2 StatefulFlowChoiceEnum
}{
	ROCEV2: StatefulFlowChoiceEnum("rocev2"),
}

func (obj *statefulFlow) Choice() StatefulFlowChoiceEnum {
	return StatefulFlowChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for Rocev2 to set choice
func (obj *statefulFlow) Rocev2() {
	obj.setChoice(StatefulFlowChoice.ROCEV2)
}

// Stateful traffic configuration.
// Choice returns a string
func (obj *statefulFlow) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *statefulFlow) setChoice(value StatefulFlowChoiceEnum) StatefulFlow {
	intValue, ok := otg.StatefulFlow_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StatefulFlowChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.StatefulFlow_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

// RoCEv2 Flow Groups.
// Rocev2Flows returns a []RoCEv2Flow
func (obj *statefulFlow) Rocev2Flows() StatefulFlowRoCEv2FlowIter {
	if len(obj.obj.Rocev2Flows) == 0 {
		obj.obj.Rocev2Flows = []*otg.RoCEv2Flow{}
	}
	if obj.rocev2FlowsHolder == nil {
		obj.rocev2FlowsHolder = newStatefulFlowRoCEv2FlowIter(&obj.obj.Rocev2Flows).setMsg(obj)
	}
	return obj.rocev2FlowsHolder
}

type statefulFlowRoCEv2FlowIter struct {
	obj             *statefulFlow
	roCEv2FlowSlice []RoCEv2Flow
	fieldPtr        *[]*otg.RoCEv2Flow
}

func newStatefulFlowRoCEv2FlowIter(ptr *[]*otg.RoCEv2Flow) StatefulFlowRoCEv2FlowIter {
	return &statefulFlowRoCEv2FlowIter{fieldPtr: ptr}
}

type StatefulFlowRoCEv2FlowIter interface {
	setMsg(*statefulFlow) StatefulFlowRoCEv2FlowIter
	Items() []RoCEv2Flow
	Add() RoCEv2Flow
	Append(items ...RoCEv2Flow) StatefulFlowRoCEv2FlowIter
	Set(index int, newObj RoCEv2Flow) StatefulFlowRoCEv2FlowIter
	Clear() StatefulFlowRoCEv2FlowIter
	clearHolderSlice() StatefulFlowRoCEv2FlowIter
	appendHolderSlice(item RoCEv2Flow) StatefulFlowRoCEv2FlowIter
}

func (obj *statefulFlowRoCEv2FlowIter) setMsg(msg *statefulFlow) StatefulFlowRoCEv2FlowIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&roCEv2Flow{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *statefulFlowRoCEv2FlowIter) Items() []RoCEv2Flow {
	return obj.roCEv2FlowSlice
}

func (obj *statefulFlowRoCEv2FlowIter) Add() RoCEv2Flow {
	newObj := &otg.RoCEv2Flow{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &roCEv2Flow{obj: newObj}
	newLibObj.setDefault()
	obj.roCEv2FlowSlice = append(obj.roCEv2FlowSlice, newLibObj)
	return newLibObj
}

func (obj *statefulFlowRoCEv2FlowIter) Append(items ...RoCEv2Flow) StatefulFlowRoCEv2FlowIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.roCEv2FlowSlice = append(obj.roCEv2FlowSlice, item)
	}
	return obj
}

func (obj *statefulFlowRoCEv2FlowIter) Set(index int, newObj RoCEv2Flow) StatefulFlowRoCEv2FlowIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.roCEv2FlowSlice[index] = newObj
	return obj
}
func (obj *statefulFlowRoCEv2FlowIter) Clear() StatefulFlowRoCEv2FlowIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.RoCEv2Flow{}
		obj.roCEv2FlowSlice = []RoCEv2Flow{}
	}
	return obj
}
func (obj *statefulFlowRoCEv2FlowIter) clearHolderSlice() StatefulFlowRoCEv2FlowIter {
	if len(obj.roCEv2FlowSlice) > 0 {
		obj.roCEv2FlowSlice = []RoCEv2Flow{}
	}
	return obj
}
func (obj *statefulFlowRoCEv2FlowIter) appendHolderSlice(item RoCEv2Flow) StatefulFlowRoCEv2FlowIter {
	obj.roCEv2FlowSlice = append(obj.roCEv2FlowSlice, item)
	return obj
}

func (obj *statefulFlow) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Rocev2Flows) != 0 {

		if set_default {
			obj.Rocev2Flows().clearHolderSlice()
			for _, item := range obj.obj.Rocev2Flows {
				obj.Rocev2Flows().appendHolderSlice(&roCEv2Flow{obj: item})
			}
		}
		for _, item := range obj.Rocev2Flows().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *statefulFlow) setDefault() {
	var choices_set int = 0
	var choice StatefulFlowChoiceEnum
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in StatefulFlow")
			}
		} else {
			intVal := otg.StatefulFlow_Choice_Enum_value[string(choice)]
			enumValue := otg.StatefulFlow_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
