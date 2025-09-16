package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathRecordRouteType1 *****
type flowRSVPPathRecordRouteType1 struct {
	validation
	obj              *otg.FlowRSVPPathRecordRouteType1
	marshaller       marshalFlowRSVPPathRecordRouteType1
	unMarshaller     unMarshalFlowRSVPPathRecordRouteType1
	subobjectsHolder FlowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter
}

func NewFlowRSVPPathRecordRouteType1() FlowRSVPPathRecordRouteType1 {
	obj := flowRSVPPathRecordRouteType1{obj: &otg.FlowRSVPPathRecordRouteType1{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathRecordRouteType1) msg() *otg.FlowRSVPPathRecordRouteType1 {
	return obj.obj
}

func (obj *flowRSVPPathRecordRouteType1) setMsg(msg *otg.FlowRSVPPathRecordRouteType1) FlowRSVPPathRecordRouteType1 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathRecordRouteType1 struct {
	obj *flowRSVPPathRecordRouteType1
}

type marshalFlowRSVPPathRecordRouteType1 interface {
	// ToProto marshals FlowRSVPPathRecordRouteType1 to protobuf object *otg.FlowRSVPPathRecordRouteType1
	ToProto() (*otg.FlowRSVPPathRecordRouteType1, error)
	// ToPbText marshals FlowRSVPPathRecordRouteType1 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathRecordRouteType1 to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathRecordRouteType1 to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPPathRecordRouteType1 struct {
	obj *flowRSVPPathRecordRouteType1
}

type unMarshalFlowRSVPPathRecordRouteType1 interface {
	// FromProto unmarshals FlowRSVPPathRecordRouteType1 from protobuf object *otg.FlowRSVPPathRecordRouteType1
	FromProto(msg *otg.FlowRSVPPathRecordRouteType1) (FlowRSVPPathRecordRouteType1, error)
	// FromPbText unmarshals FlowRSVPPathRecordRouteType1 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathRecordRouteType1 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathRecordRouteType1 from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathRecordRouteType1) Marshal() marshalFlowRSVPPathRecordRouteType1 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathRecordRouteType1{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathRecordRouteType1) Unmarshal() unMarshalFlowRSVPPathRecordRouteType1 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathRecordRouteType1{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathRecordRouteType1) ToProto() (*otg.FlowRSVPPathRecordRouteType1, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathRecordRouteType1) FromProto(msg *otg.FlowRSVPPathRecordRouteType1) (FlowRSVPPathRecordRouteType1, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathRecordRouteType1) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathRecordRouteType1) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathRecordRouteType1) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathRecordRouteType1) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathRecordRouteType1) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathRecordRouteType1) FromJson(value string) error {
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

func (obj *flowRSVPPathRecordRouteType1) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathRecordRouteType1) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathRecordRouteType1) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathRecordRouteType1) Clone() (FlowRSVPPathRecordRouteType1, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathRecordRouteType1()
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

func (obj *flowRSVPPathRecordRouteType1) setNil() {
	obj.subobjectsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathRecordRouteType1 is type1 record route has list of subobjects. Currently supported subobjects are IPv4 address(1) and Label(3).
type FlowRSVPPathRecordRouteType1 interface {
	Validation
	// msg marshals FlowRSVPPathRecordRouteType1 to protobuf object *otg.FlowRSVPPathRecordRouteType1
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathRecordRouteType1
	// setMsg unmarshals FlowRSVPPathRecordRouteType1 from protobuf object *otg.FlowRSVPPathRecordRouteType1
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathRecordRouteType1) FlowRSVPPathRecordRouteType1
	// provides marshal interface
	Marshal() marshalFlowRSVPPathRecordRouteType1
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathRecordRouteType1
	// validate validates FlowRSVPPathRecordRouteType1
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathRecordRouteType1, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Subobjects returns FlowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIterIter, set in FlowRSVPPathRecordRouteType1
	Subobjects() FlowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter
	setNil()
}

// description is TBD
// Subobjects returns a []FlowRSVPType1RecordRouteSubobjects
func (obj *flowRSVPPathRecordRouteType1) Subobjects() FlowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter {
	if len(obj.obj.Subobjects) == 0 {
		obj.obj.Subobjects = []*otg.FlowRSVPType1RecordRouteSubobjects{}
	}
	if obj.subobjectsHolder == nil {
		obj.subobjectsHolder = newFlowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter(&obj.obj.Subobjects).setMsg(obj)
	}
	return obj.subobjectsHolder
}

type flowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter struct {
	obj                                     *flowRSVPPathRecordRouteType1
	flowRSVPType1RecordRouteSubobjectsSlice []FlowRSVPType1RecordRouteSubobjects
	fieldPtr                                *[]*otg.FlowRSVPType1RecordRouteSubobjects
}

func newFlowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter(ptr *[]*otg.FlowRSVPType1RecordRouteSubobjects) FlowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter {
	return &flowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter{fieldPtr: ptr}
}

type FlowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter interface {
	setMsg(*flowRSVPPathRecordRouteType1) FlowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter
	Items() []FlowRSVPType1RecordRouteSubobjects
	Add() FlowRSVPType1RecordRouteSubobjects
	Append(items ...FlowRSVPType1RecordRouteSubobjects) FlowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter
	Set(index int, newObj FlowRSVPType1RecordRouteSubobjects) FlowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter
	Clear() FlowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter
	clearHolderSlice() FlowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter
	appendHolderSlice(item FlowRSVPType1RecordRouteSubobjects) FlowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter
}

func (obj *flowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter) setMsg(msg *flowRSVPPathRecordRouteType1) FlowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowRSVPType1RecordRouteSubobjects{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter) Items() []FlowRSVPType1RecordRouteSubobjects {
	return obj.flowRSVPType1RecordRouteSubobjectsSlice
}

func (obj *flowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter) Add() FlowRSVPType1RecordRouteSubobjects {
	newObj := &otg.FlowRSVPType1RecordRouteSubobjects{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowRSVPType1RecordRouteSubobjects{obj: newObj}
	newLibObj.setDefault()
	obj.flowRSVPType1RecordRouteSubobjectsSlice = append(obj.flowRSVPType1RecordRouteSubobjectsSlice, newLibObj)
	return newLibObj
}

func (obj *flowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter) Append(items ...FlowRSVPType1RecordRouteSubobjects) FlowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowRSVPType1RecordRouteSubobjectsSlice = append(obj.flowRSVPType1RecordRouteSubobjectsSlice, item)
	}
	return obj
}

func (obj *flowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter) Set(index int, newObj FlowRSVPType1RecordRouteSubobjects) FlowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowRSVPType1RecordRouteSubobjectsSlice[index] = newObj
	return obj
}
func (obj *flowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter) Clear() FlowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowRSVPType1RecordRouteSubobjects{}
		obj.flowRSVPType1RecordRouteSubobjectsSlice = []FlowRSVPType1RecordRouteSubobjects{}
	}
	return obj
}
func (obj *flowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter) clearHolderSlice() FlowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter {
	if len(obj.flowRSVPType1RecordRouteSubobjectsSlice) > 0 {
		obj.flowRSVPType1RecordRouteSubobjectsSlice = []FlowRSVPType1RecordRouteSubobjects{}
	}
	return obj
}
func (obj *flowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter) appendHolderSlice(item FlowRSVPType1RecordRouteSubobjects) FlowRSVPPathRecordRouteType1FlowRSVPType1RecordRouteSubobjectsIter {
	obj.flowRSVPType1RecordRouteSubobjectsSlice = append(obj.flowRSVPType1RecordRouteSubobjectsSlice, item)
	return obj
}

func (obj *flowRSVPPathRecordRouteType1) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Subobjects) != 0 {

		if set_default {
			obj.Subobjects().clearHolderSlice()
			for _, item := range obj.obj.Subobjects {
				obj.Subobjects().appendHolderSlice(&flowRSVPType1RecordRouteSubobjects{obj: item})
			}
		}
		for _, item := range obj.Subobjects().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *flowRSVPPathRecordRouteType1) setDefault() {

}
