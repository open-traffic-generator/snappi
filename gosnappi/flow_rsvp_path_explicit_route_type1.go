package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathExplicitRouteType1 *****
type flowRSVPPathExplicitRouteType1 struct {
	validation
	obj              *otg.FlowRSVPPathExplicitRouteType1
	marshaller       marshalFlowRSVPPathExplicitRouteType1
	unMarshaller     unMarshalFlowRSVPPathExplicitRouteType1
	subobjectsHolder FlowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter
}

func NewFlowRSVPPathExplicitRouteType1() FlowRSVPPathExplicitRouteType1 {
	obj := flowRSVPPathExplicitRouteType1{obj: &otg.FlowRSVPPathExplicitRouteType1{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathExplicitRouteType1) msg() *otg.FlowRSVPPathExplicitRouteType1 {
	return obj.obj
}

func (obj *flowRSVPPathExplicitRouteType1) setMsg(msg *otg.FlowRSVPPathExplicitRouteType1) FlowRSVPPathExplicitRouteType1 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathExplicitRouteType1 struct {
	obj *flowRSVPPathExplicitRouteType1
}

type marshalFlowRSVPPathExplicitRouteType1 interface {
	// ToProto marshals FlowRSVPPathExplicitRouteType1 to protobuf object *otg.FlowRSVPPathExplicitRouteType1
	ToProto() (*otg.FlowRSVPPathExplicitRouteType1, error)
	// ToPbText marshals FlowRSVPPathExplicitRouteType1 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathExplicitRouteType1 to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathExplicitRouteType1 to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPPathExplicitRouteType1 struct {
	obj *flowRSVPPathExplicitRouteType1
}

type unMarshalFlowRSVPPathExplicitRouteType1 interface {
	// FromProto unmarshals FlowRSVPPathExplicitRouteType1 from protobuf object *otg.FlowRSVPPathExplicitRouteType1
	FromProto(msg *otg.FlowRSVPPathExplicitRouteType1) (FlowRSVPPathExplicitRouteType1, error)
	// FromPbText unmarshals FlowRSVPPathExplicitRouteType1 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathExplicitRouteType1 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathExplicitRouteType1 from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathExplicitRouteType1) Marshal() marshalFlowRSVPPathExplicitRouteType1 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathExplicitRouteType1{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathExplicitRouteType1) Unmarshal() unMarshalFlowRSVPPathExplicitRouteType1 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathExplicitRouteType1{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathExplicitRouteType1) ToProto() (*otg.FlowRSVPPathExplicitRouteType1, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathExplicitRouteType1) FromProto(msg *otg.FlowRSVPPathExplicitRouteType1) (FlowRSVPPathExplicitRouteType1, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathExplicitRouteType1) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathExplicitRouteType1) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathExplicitRouteType1) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathExplicitRouteType1) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathExplicitRouteType1) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathExplicitRouteType1) FromJson(value string) error {
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

func (obj *flowRSVPPathExplicitRouteType1) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathExplicitRouteType1) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathExplicitRouteType1) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathExplicitRouteType1) Clone() (FlowRSVPPathExplicitRouteType1, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathExplicitRouteType1()
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

func (obj *flowRSVPPathExplicitRouteType1) setNil() {
	obj.subobjectsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathExplicitRouteType1 is type1 Explicit Route has subobjects. Currently supported subobjects are IPv4 prefix and Autonomous system number.
type FlowRSVPPathExplicitRouteType1 interface {
	Validation
	// msg marshals FlowRSVPPathExplicitRouteType1 to protobuf object *otg.FlowRSVPPathExplicitRouteType1
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathExplicitRouteType1
	// setMsg unmarshals FlowRSVPPathExplicitRouteType1 from protobuf object *otg.FlowRSVPPathExplicitRouteType1
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathExplicitRouteType1) FlowRSVPPathExplicitRouteType1
	// provides marshal interface
	Marshal() marshalFlowRSVPPathExplicitRouteType1
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathExplicitRouteType1
	// validate validates FlowRSVPPathExplicitRouteType1
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathExplicitRouteType1, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Subobjects returns FlowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIterIter, set in FlowRSVPPathExplicitRouteType1
	Subobjects() FlowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter
	setNil()
}

// description is TBD
// Subobjects returns a []FlowRSVPType1ExplicitRouteSubobjects
func (obj *flowRSVPPathExplicitRouteType1) Subobjects() FlowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter {
	if len(obj.obj.Subobjects) == 0 {
		obj.obj.Subobjects = []*otg.FlowRSVPType1ExplicitRouteSubobjects{}
	}
	if obj.subobjectsHolder == nil {
		obj.subobjectsHolder = newFlowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter(&obj.obj.Subobjects).setMsg(obj)
	}
	return obj.subobjectsHolder
}

type flowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter struct {
	obj                                       *flowRSVPPathExplicitRouteType1
	flowRSVPType1ExplicitRouteSubobjectsSlice []FlowRSVPType1ExplicitRouteSubobjects
	fieldPtr                                  *[]*otg.FlowRSVPType1ExplicitRouteSubobjects
}

func newFlowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter(ptr *[]*otg.FlowRSVPType1ExplicitRouteSubobjects) FlowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter {
	return &flowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter{fieldPtr: ptr}
}

type FlowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter interface {
	setMsg(*flowRSVPPathExplicitRouteType1) FlowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter
	Items() []FlowRSVPType1ExplicitRouteSubobjects
	Add() FlowRSVPType1ExplicitRouteSubobjects
	Append(items ...FlowRSVPType1ExplicitRouteSubobjects) FlowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter
	Set(index int, newObj FlowRSVPType1ExplicitRouteSubobjects) FlowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter
	Clear() FlowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter
	clearHolderSlice() FlowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter
	appendHolderSlice(item FlowRSVPType1ExplicitRouteSubobjects) FlowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter
}

func (obj *flowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter) setMsg(msg *flowRSVPPathExplicitRouteType1) FlowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowRSVPType1ExplicitRouteSubobjects{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter) Items() []FlowRSVPType1ExplicitRouteSubobjects {
	return obj.flowRSVPType1ExplicitRouteSubobjectsSlice
}

func (obj *flowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter) Add() FlowRSVPType1ExplicitRouteSubobjects {
	newObj := &otg.FlowRSVPType1ExplicitRouteSubobjects{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowRSVPType1ExplicitRouteSubobjects{obj: newObj}
	newLibObj.setDefault()
	obj.flowRSVPType1ExplicitRouteSubobjectsSlice = append(obj.flowRSVPType1ExplicitRouteSubobjectsSlice, newLibObj)
	return newLibObj
}

func (obj *flowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter) Append(items ...FlowRSVPType1ExplicitRouteSubobjects) FlowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowRSVPType1ExplicitRouteSubobjectsSlice = append(obj.flowRSVPType1ExplicitRouteSubobjectsSlice, item)
	}
	return obj
}

func (obj *flowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter) Set(index int, newObj FlowRSVPType1ExplicitRouteSubobjects) FlowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowRSVPType1ExplicitRouteSubobjectsSlice[index] = newObj
	return obj
}
func (obj *flowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter) Clear() FlowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowRSVPType1ExplicitRouteSubobjects{}
		obj.flowRSVPType1ExplicitRouteSubobjectsSlice = []FlowRSVPType1ExplicitRouteSubobjects{}
	}
	return obj
}
func (obj *flowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter) clearHolderSlice() FlowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter {
	if len(obj.flowRSVPType1ExplicitRouteSubobjectsSlice) > 0 {
		obj.flowRSVPType1ExplicitRouteSubobjectsSlice = []FlowRSVPType1ExplicitRouteSubobjects{}
	}
	return obj
}
func (obj *flowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter) appendHolderSlice(item FlowRSVPType1ExplicitRouteSubobjects) FlowRSVPPathExplicitRouteType1FlowRSVPType1ExplicitRouteSubobjectsIter {
	obj.flowRSVPType1ExplicitRouteSubobjectsSlice = append(obj.flowRSVPType1ExplicitRouteSubobjectsSlice, item)
	return obj
}

func (obj *flowRSVPPathExplicitRouteType1) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Subobjects) != 0 {

		if set_default {
			obj.Subobjects().clearHolderSlice()
			for _, item := range obj.obj.Subobjects {
				obj.Subobjects().appendHolderSlice(&flowRSVPType1ExplicitRouteSubobjects{obj: item})
			}
		}
		for _, item := range obj.Subobjects().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *flowRSVPPathExplicitRouteType1) setDefault() {

}
