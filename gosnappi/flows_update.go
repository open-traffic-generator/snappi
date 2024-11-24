package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowsUpdate *****
type flowsUpdate struct {
	validation
	obj          *otg.FlowsUpdate
	marshaller   marshalFlowsUpdate
	unMarshaller unMarshalFlowsUpdate
	flowsHolder  FlowsUpdateFlowIter
}

func NewFlowsUpdate() FlowsUpdate {
	obj := flowsUpdate{obj: &otg.FlowsUpdate{}}
	obj.setDefault()
	return &obj
}

func (obj *flowsUpdate) msg() *otg.FlowsUpdate {
	return obj.obj
}

func (obj *flowsUpdate) setMsg(msg *otg.FlowsUpdate) FlowsUpdate {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowsUpdate struct {
	obj *flowsUpdate
}

type marshalFlowsUpdate interface {
	// ToProto marshals FlowsUpdate to protobuf object *otg.FlowsUpdate
	ToProto() (*otg.FlowsUpdate, error)
	// ToPbText marshals FlowsUpdate to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowsUpdate to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowsUpdate to JSON text
	ToJson() (string, error)
}

type unMarshalflowsUpdate struct {
	obj *flowsUpdate
}

type unMarshalFlowsUpdate interface {
	// FromProto unmarshals FlowsUpdate from protobuf object *otg.FlowsUpdate
	FromProto(msg *otg.FlowsUpdate) (FlowsUpdate, error)
	// FromPbText unmarshals FlowsUpdate from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowsUpdate from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowsUpdate from JSON text
	FromJson(value string) error
}

func (obj *flowsUpdate) Marshal() marshalFlowsUpdate {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowsUpdate{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowsUpdate) Unmarshal() unMarshalFlowsUpdate {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowsUpdate{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowsUpdate) ToProto() (*otg.FlowsUpdate, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowsUpdate) FromProto(msg *otg.FlowsUpdate) (FlowsUpdate, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowsUpdate) ToPbText() (string, error) {
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

func (m *unMarshalflowsUpdate) FromPbText(value string) error {
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

func (m *marshalflowsUpdate) ToYaml() (string, error) {
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

func (m *unMarshalflowsUpdate) FromYaml(value string) error {
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

func (m *marshalflowsUpdate) ToJson() (string, error) {
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

func (m *unMarshalflowsUpdate) FromJson(value string) error {
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

func (obj *flowsUpdate) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowsUpdate) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowsUpdate) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowsUpdate) Clone() (FlowsUpdate, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowsUpdate()
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

func (obj *flowsUpdate) setNil() {
	obj.flowsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowsUpdate is a container of flows with associated properties to be updated without affecting the flows current transmit state.
type FlowsUpdate interface {
	Validation
	// msg marshals FlowsUpdate to protobuf object *otg.FlowsUpdate
	// and doesn't set defaults
	msg() *otg.FlowsUpdate
	// setMsg unmarshals FlowsUpdate from protobuf object *otg.FlowsUpdate
	// and doesn't set defaults
	setMsg(*otg.FlowsUpdate) FlowsUpdate
	// provides marshal interface
	Marshal() marshalFlowsUpdate
	// provides unmarshal interface
	Unmarshal() unMarshalFlowsUpdate
	// validate validates FlowsUpdate
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowsUpdate, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PropertyNames returns []FlowsUpdatePropertyNamesEnum, set in FlowsUpdate
	PropertyNames() []FlowsUpdatePropertyNamesEnum
	// SetPropertyNames assigns []FlowsUpdatePropertyNamesEnum provided by user to FlowsUpdate
	SetPropertyNames(value []FlowsUpdatePropertyNamesEnum) FlowsUpdate
	// Flows returns FlowsUpdateFlowIterIter, set in FlowsUpdate
	Flows() FlowsUpdateFlowIter
	setNil()
}

type FlowsUpdatePropertyNamesEnum string

// Enum of PropertyNames on FlowsUpdate
var FlowsUpdatePropertyNames = struct {
	RATE FlowsUpdatePropertyNamesEnum
	SIZE FlowsUpdatePropertyNamesEnum
}{
	RATE: FlowsUpdatePropertyNamesEnum("rate"),
	SIZE: FlowsUpdatePropertyNamesEnum("size"),
}

func (obj *flowsUpdate) PropertyNames() []FlowsUpdatePropertyNamesEnum {
	items := []FlowsUpdatePropertyNamesEnum{}
	for _, item := range obj.obj.PropertyNames {
		items = append(items, FlowsUpdatePropertyNamesEnum(item.String()))
	}
	return items
}

// Flow properties to be updated without affecting the transmit state.
// SetPropertyNames sets the []string value in the FlowsUpdate object
func (obj *flowsUpdate) SetPropertyNames(value []FlowsUpdatePropertyNamesEnum) FlowsUpdate {

	items := []otg.FlowsUpdate_PropertyNames_Enum{}
	for _, item := range value {
		intValue := otg.FlowsUpdate_PropertyNames_Enum_value[string(item)]
		items = append(items, otg.FlowsUpdate_PropertyNames_Enum(intValue))
	}
	obj.obj.PropertyNames = items
	return obj
}

// The list of configured flows for which given property will be updated.
// Flows returns a []Flow
func (obj *flowsUpdate) Flows() FlowsUpdateFlowIter {
	if len(obj.obj.Flows) == 0 {
		obj.obj.Flows = []*otg.Flow{}
	}
	if obj.flowsHolder == nil {
		obj.flowsHolder = newFlowsUpdateFlowIter(&obj.obj.Flows).setMsg(obj)
	}
	return obj.flowsHolder
}

type flowsUpdateFlowIter struct {
	obj       *flowsUpdate
	flowSlice []Flow
	fieldPtr  *[]*otg.Flow
}

func newFlowsUpdateFlowIter(ptr *[]*otg.Flow) FlowsUpdateFlowIter {
	return &flowsUpdateFlowIter{fieldPtr: ptr}
}

type FlowsUpdateFlowIter interface {
	setMsg(*flowsUpdate) FlowsUpdateFlowIter
	Items() []Flow
	Add() Flow
	Append(items ...Flow) FlowsUpdateFlowIter
	Set(index int, newObj Flow) FlowsUpdateFlowIter
	Clear() FlowsUpdateFlowIter
	clearHolderSlice() FlowsUpdateFlowIter
	appendHolderSlice(item Flow) FlowsUpdateFlowIter
}

func (obj *flowsUpdateFlowIter) setMsg(msg *flowsUpdate) FlowsUpdateFlowIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flow{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowsUpdateFlowIter) Items() []Flow {
	return obj.flowSlice
}

func (obj *flowsUpdateFlowIter) Add() Flow {
	newObj := &otg.Flow{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flow{obj: newObj}
	newLibObj.setDefault()
	obj.flowSlice = append(obj.flowSlice, newLibObj)
	return newLibObj
}

func (obj *flowsUpdateFlowIter) Append(items ...Flow) FlowsUpdateFlowIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowSlice = append(obj.flowSlice, item)
	}
	return obj
}

func (obj *flowsUpdateFlowIter) Set(index int, newObj Flow) FlowsUpdateFlowIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowSlice[index] = newObj
	return obj
}
func (obj *flowsUpdateFlowIter) Clear() FlowsUpdateFlowIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Flow{}
		obj.flowSlice = []Flow{}
	}
	return obj
}
func (obj *flowsUpdateFlowIter) clearHolderSlice() FlowsUpdateFlowIter {
	if len(obj.flowSlice) > 0 {
		obj.flowSlice = []Flow{}
	}
	return obj
}
func (obj *flowsUpdateFlowIter) appendHolderSlice(item Flow) FlowsUpdateFlowIter {
	obj.flowSlice = append(obj.flowSlice, item)
	return obj
}

func (obj *flowsUpdate) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Flows) != 0 {

		if set_default {
			obj.Flows().clearHolderSlice()
			for _, item := range obj.obj.Flows {
				obj.Flows().appendHolderSlice(&flow{obj: item})
			}
		}
		for _, item := range obj.Flows().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *flowsUpdate) setDefault() {

}
