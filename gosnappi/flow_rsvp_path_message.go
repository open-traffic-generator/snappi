package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathMessage *****
type flowRSVPPathMessage struct {
	validation
	obj           *otg.FlowRSVPPathMessage
	marshaller    marshalFlowRSVPPathMessage
	unMarshaller  unMarshalFlowRSVPPathMessage
	objectsHolder FlowRSVPPathMessageFlowRSVPPathObjectsIter
}

func NewFlowRSVPPathMessage() FlowRSVPPathMessage {
	obj := flowRSVPPathMessage{obj: &otg.FlowRSVPPathMessage{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathMessage) msg() *otg.FlowRSVPPathMessage {
	return obj.obj
}

func (obj *flowRSVPPathMessage) setMsg(msg *otg.FlowRSVPPathMessage) FlowRSVPPathMessage {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathMessage struct {
	obj *flowRSVPPathMessage
}

type marshalFlowRSVPPathMessage interface {
	// ToProto marshals FlowRSVPPathMessage to protobuf object *otg.FlowRSVPPathMessage
	ToProto() (*otg.FlowRSVPPathMessage, error)
	// ToPbText marshals FlowRSVPPathMessage to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathMessage to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathMessage to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowRSVPPathMessage to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowRSVPPathMessage struct {
	obj *flowRSVPPathMessage
}

type unMarshalFlowRSVPPathMessage interface {
	// FromProto unmarshals FlowRSVPPathMessage from protobuf object *otg.FlowRSVPPathMessage
	FromProto(msg *otg.FlowRSVPPathMessage) (FlowRSVPPathMessage, error)
	// FromPbText unmarshals FlowRSVPPathMessage from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathMessage from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathMessage from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathMessage) Marshal() marshalFlowRSVPPathMessage {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathMessage{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathMessage) Unmarshal() unMarshalFlowRSVPPathMessage {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathMessage{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathMessage) ToProto() (*otg.FlowRSVPPathMessage, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathMessage) FromProto(msg *otg.FlowRSVPPathMessage) (FlowRSVPPathMessage, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathMessage) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathMessage) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathMessage) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathMessage) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathMessage) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalflowRSVPPathMessage) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathMessage) FromJson(value string) error {
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

func (obj *flowRSVPPathMessage) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathMessage) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathMessage) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathMessage) Clone() (FlowRSVPPathMessage, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathMessage()
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

func (obj *flowRSVPPathMessage) setNil() {
	obj.objectsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathMessage is "Path" message requires the following list of objects in order as defined in https://www.rfc-editor.org/rfc/rfc3209.html#page-15: 1. SESSION 2. RSVP_HOP 3. TIME_VALUES 4. EXPLICIT_ROUTE [optional] 5. LABEL_REQUEST 6. SESSION_ATTRIBUTE [optional] 7. SENDER_TEMPLATE 8. SENDER_TSPEC 9. RECORD_ROUTE [optional]
type FlowRSVPPathMessage interface {
	Validation
	// msg marshals FlowRSVPPathMessage to protobuf object *otg.FlowRSVPPathMessage
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathMessage
	// setMsg unmarshals FlowRSVPPathMessage from protobuf object *otg.FlowRSVPPathMessage
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathMessage) FlowRSVPPathMessage
	// provides marshal interface
	Marshal() marshalFlowRSVPPathMessage
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathMessage
	// validate validates FlowRSVPPathMessage
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathMessage, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Objects returns FlowRSVPPathMessageFlowRSVPPathObjectsIterIter, set in FlowRSVPPathMessage
	Objects() FlowRSVPPathMessageFlowRSVPPathObjectsIter
	setNil()
}

// "Path" message requires atleast SESSION, RSVP_HOP, TIME_VALUES, LABEL_REQUEST, SENDER_TEMPLATE and SENDER_TSPEC objects in order.
// Objects returns a []FlowRSVPPathObjects
func (obj *flowRSVPPathMessage) Objects() FlowRSVPPathMessageFlowRSVPPathObjectsIter {
	if len(obj.obj.Objects) == 0 {
		obj.obj.Objects = []*otg.FlowRSVPPathObjects{}
	}
	if obj.objectsHolder == nil {
		obj.objectsHolder = newFlowRSVPPathMessageFlowRSVPPathObjectsIter(&obj.obj.Objects).setMsg(obj)
	}
	return obj.objectsHolder
}

type flowRSVPPathMessageFlowRSVPPathObjectsIter struct {
	obj                      *flowRSVPPathMessage
	flowRSVPPathObjectsSlice []FlowRSVPPathObjects
	fieldPtr                 *[]*otg.FlowRSVPPathObjects
}

func newFlowRSVPPathMessageFlowRSVPPathObjectsIter(ptr *[]*otg.FlowRSVPPathObjects) FlowRSVPPathMessageFlowRSVPPathObjectsIter {
	return &flowRSVPPathMessageFlowRSVPPathObjectsIter{fieldPtr: ptr}
}

type FlowRSVPPathMessageFlowRSVPPathObjectsIter interface {
	setMsg(*flowRSVPPathMessage) FlowRSVPPathMessageFlowRSVPPathObjectsIter
	Items() []FlowRSVPPathObjects
	Add() FlowRSVPPathObjects
	Append(items ...FlowRSVPPathObjects) FlowRSVPPathMessageFlowRSVPPathObjectsIter
	Set(index int, newObj FlowRSVPPathObjects) FlowRSVPPathMessageFlowRSVPPathObjectsIter
	Clear() FlowRSVPPathMessageFlowRSVPPathObjectsIter
	clearHolderSlice() FlowRSVPPathMessageFlowRSVPPathObjectsIter
	appendHolderSlice(item FlowRSVPPathObjects) FlowRSVPPathMessageFlowRSVPPathObjectsIter
}

func (obj *flowRSVPPathMessageFlowRSVPPathObjectsIter) setMsg(msg *flowRSVPPathMessage) FlowRSVPPathMessageFlowRSVPPathObjectsIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowRSVPPathObjects{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowRSVPPathMessageFlowRSVPPathObjectsIter) Items() []FlowRSVPPathObjects {
	return obj.flowRSVPPathObjectsSlice
}

func (obj *flowRSVPPathMessageFlowRSVPPathObjectsIter) Add() FlowRSVPPathObjects {
	newObj := &otg.FlowRSVPPathObjects{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowRSVPPathObjects{obj: newObj}
	newLibObj.setDefault()
	obj.flowRSVPPathObjectsSlice = append(obj.flowRSVPPathObjectsSlice, newLibObj)
	return newLibObj
}

func (obj *flowRSVPPathMessageFlowRSVPPathObjectsIter) Append(items ...FlowRSVPPathObjects) FlowRSVPPathMessageFlowRSVPPathObjectsIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowRSVPPathObjectsSlice = append(obj.flowRSVPPathObjectsSlice, item)
	}
	return obj
}

func (obj *flowRSVPPathMessageFlowRSVPPathObjectsIter) Set(index int, newObj FlowRSVPPathObjects) FlowRSVPPathMessageFlowRSVPPathObjectsIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowRSVPPathObjectsSlice[index] = newObj
	return obj
}
func (obj *flowRSVPPathMessageFlowRSVPPathObjectsIter) Clear() FlowRSVPPathMessageFlowRSVPPathObjectsIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowRSVPPathObjects{}
		obj.flowRSVPPathObjectsSlice = []FlowRSVPPathObjects{}
	}
	return obj
}
func (obj *flowRSVPPathMessageFlowRSVPPathObjectsIter) clearHolderSlice() FlowRSVPPathMessageFlowRSVPPathObjectsIter {
	if len(obj.flowRSVPPathObjectsSlice) > 0 {
		obj.flowRSVPPathObjectsSlice = []FlowRSVPPathObjects{}
	}
	return obj
}
func (obj *flowRSVPPathMessageFlowRSVPPathObjectsIter) appendHolderSlice(item FlowRSVPPathObjects) FlowRSVPPathMessageFlowRSVPPathObjectsIter {
	obj.flowRSVPPathObjectsSlice = append(obj.flowRSVPPathObjectsSlice, item)
	return obj
}

func (obj *flowRSVPPathMessage) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Objects) != 0 {

		if set_default {
			obj.Objects().clearHolderSlice()
			for _, item := range obj.obj.Objects {
				obj.Objects().appendHolderSlice(&flowRSVPPathObjects{obj: item})
			}
		}
		for _, item := range obj.Objects().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *flowRSVPPathMessage) setDefault() {

}
