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
	obj          *otg.StatefulFlow
	marshaller   marshalStatefulFlow
	unMarshaller unMarshalStatefulFlow
	rocev2Holder StatefulFlowRocev2FlowIter
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
	obj.rocev2Holder = nil
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
	// Rocev2 returns StatefulFlowRocev2FlowIterIter, set in StatefulFlow
	Rocev2() StatefulFlowRocev2FlowIter
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
	obj.obj.Rocev2 = nil
	obj.rocev2Holder = nil

	if value == StatefulFlowChoice.ROCEV2 {
		obj.obj.Rocev2 = []*otg.Rocev2Flow{}
	}

	return obj
}

// Rocev2 Flow Groups.
// Rocev2 returns a []Rocev2Flow
func (obj *statefulFlow) Rocev2() StatefulFlowRocev2FlowIter {
	if len(obj.obj.Rocev2) == 0 {
		obj.setChoice(StatefulFlowChoice.ROCEV2)
	}
	if obj.rocev2Holder == nil {
		obj.rocev2Holder = newStatefulFlowRocev2FlowIter(&obj.obj.Rocev2).setMsg(obj)
	}
	return obj.rocev2Holder
}

type statefulFlowRocev2FlowIter struct {
	obj             *statefulFlow
	rocev2FlowSlice []Rocev2Flow
	fieldPtr        *[]*otg.Rocev2Flow
}

func newStatefulFlowRocev2FlowIter(ptr *[]*otg.Rocev2Flow) StatefulFlowRocev2FlowIter {
	return &statefulFlowRocev2FlowIter{fieldPtr: ptr}
}

type StatefulFlowRocev2FlowIter interface {
	setMsg(*statefulFlow) StatefulFlowRocev2FlowIter
	Items() []Rocev2Flow
	Add() Rocev2Flow
	Append(items ...Rocev2Flow) StatefulFlowRocev2FlowIter
	Set(index int, newObj Rocev2Flow) StatefulFlowRocev2FlowIter
	Clear() StatefulFlowRocev2FlowIter
	clearHolderSlice() StatefulFlowRocev2FlowIter
	appendHolderSlice(item Rocev2Flow) StatefulFlowRocev2FlowIter
}

func (obj *statefulFlowRocev2FlowIter) setMsg(msg *statefulFlow) StatefulFlowRocev2FlowIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&rocev2Flow{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *statefulFlowRocev2FlowIter) Items() []Rocev2Flow {
	return obj.rocev2FlowSlice
}

func (obj *statefulFlowRocev2FlowIter) Add() Rocev2Flow {
	newObj := &otg.Rocev2Flow{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &rocev2Flow{obj: newObj}
	newLibObj.setDefault()
	obj.rocev2FlowSlice = append(obj.rocev2FlowSlice, newLibObj)
	return newLibObj
}

func (obj *statefulFlowRocev2FlowIter) Append(items ...Rocev2Flow) StatefulFlowRocev2FlowIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.rocev2FlowSlice = append(obj.rocev2FlowSlice, item)
	}
	return obj
}

func (obj *statefulFlowRocev2FlowIter) Set(index int, newObj Rocev2Flow) StatefulFlowRocev2FlowIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.rocev2FlowSlice[index] = newObj
	return obj
}
func (obj *statefulFlowRocev2FlowIter) Clear() StatefulFlowRocev2FlowIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Rocev2Flow{}
		obj.rocev2FlowSlice = []Rocev2Flow{}
	}
	return obj
}
func (obj *statefulFlowRocev2FlowIter) clearHolderSlice() StatefulFlowRocev2FlowIter {
	if len(obj.rocev2FlowSlice) > 0 {
		obj.rocev2FlowSlice = []Rocev2Flow{}
	}
	return obj
}
func (obj *statefulFlowRocev2FlowIter) appendHolderSlice(item Rocev2Flow) StatefulFlowRocev2FlowIter {
	obj.rocev2FlowSlice = append(obj.rocev2FlowSlice, item)
	return obj
}

func (obj *statefulFlow) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Rocev2) != 0 {

		if set_default {
			obj.Rocev2().clearHolderSlice()
			for _, item := range obj.obj.Rocev2 {
				obj.Rocev2().appendHolderSlice(&rocev2Flow{obj: item})
			}
		}
		for _, item := range obj.Rocev2().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *statefulFlow) setDefault() {
	var choices_set int = 0
	var choice StatefulFlowChoiceEnum

	if len(obj.obj.Rocev2) > 0 {
		choices_set += 1
		choice = StatefulFlowChoice.ROCEV2
	}
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
