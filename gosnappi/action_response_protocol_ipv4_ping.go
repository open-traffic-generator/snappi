package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionResponseProtocolIpv4Ping *****
type actionResponseProtocolIpv4Ping struct {
	validation
	obj             *otg.ActionResponseProtocolIpv4Ping
	marshaller      marshalActionResponseProtocolIpv4Ping
	unMarshaller    unMarshalActionResponseProtocolIpv4Ping
	responsesHolder ActionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter
}

func NewActionResponseProtocolIpv4Ping() ActionResponseProtocolIpv4Ping {
	obj := actionResponseProtocolIpv4Ping{obj: &otg.ActionResponseProtocolIpv4Ping{}}
	obj.setDefault()
	return &obj
}

func (obj *actionResponseProtocolIpv4Ping) msg() *otg.ActionResponseProtocolIpv4Ping {
	return obj.obj
}

func (obj *actionResponseProtocolIpv4Ping) setMsg(msg *otg.ActionResponseProtocolIpv4Ping) ActionResponseProtocolIpv4Ping {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionResponseProtocolIpv4Ping struct {
	obj *actionResponseProtocolIpv4Ping
}

type marshalActionResponseProtocolIpv4Ping interface {
	// ToProto marshals ActionResponseProtocolIpv4Ping to protobuf object *otg.ActionResponseProtocolIpv4Ping
	ToProto() (*otg.ActionResponseProtocolIpv4Ping, error)
	// ToPbText marshals ActionResponseProtocolIpv4Ping to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionResponseProtocolIpv4Ping to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionResponseProtocolIpv4Ping to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals ActionResponseProtocolIpv4Ping to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalactionResponseProtocolIpv4Ping struct {
	obj *actionResponseProtocolIpv4Ping
}

type unMarshalActionResponseProtocolIpv4Ping interface {
	// FromProto unmarshals ActionResponseProtocolIpv4Ping from protobuf object *otg.ActionResponseProtocolIpv4Ping
	FromProto(msg *otg.ActionResponseProtocolIpv4Ping) (ActionResponseProtocolIpv4Ping, error)
	// FromPbText unmarshals ActionResponseProtocolIpv4Ping from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionResponseProtocolIpv4Ping from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionResponseProtocolIpv4Ping from JSON text
	FromJson(value string) error
}

func (obj *actionResponseProtocolIpv4Ping) Marshal() marshalActionResponseProtocolIpv4Ping {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionResponseProtocolIpv4Ping{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionResponseProtocolIpv4Ping) Unmarshal() unMarshalActionResponseProtocolIpv4Ping {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionResponseProtocolIpv4Ping{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionResponseProtocolIpv4Ping) ToProto() (*otg.ActionResponseProtocolIpv4Ping, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionResponseProtocolIpv4Ping) FromProto(msg *otg.ActionResponseProtocolIpv4Ping) (ActionResponseProtocolIpv4Ping, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionResponseProtocolIpv4Ping) ToPbText() (string, error) {
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

func (m *unMarshalactionResponseProtocolIpv4Ping) FromPbText(value string) error {
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

func (m *marshalactionResponseProtocolIpv4Ping) ToYaml() (string, error) {
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

func (m *unMarshalactionResponseProtocolIpv4Ping) FromYaml(value string) error {
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

func (m *marshalactionResponseProtocolIpv4Ping) ToJsonRaw() (string, error) {
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

func (m *marshalactionResponseProtocolIpv4Ping) ToJson() (string, error) {
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

func (m *unMarshalactionResponseProtocolIpv4Ping) FromJson(value string) error {
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

func (obj *actionResponseProtocolIpv4Ping) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionResponseProtocolIpv4Ping) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionResponseProtocolIpv4Ping) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionResponseProtocolIpv4Ping) Clone() (ActionResponseProtocolIpv4Ping, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionResponseProtocolIpv4Ping()
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

func (obj *actionResponseProtocolIpv4Ping) setNil() {
	obj.responsesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ActionResponseProtocolIpv4Ping is response for ping initiated between multiple source and destination pairs.
type ActionResponseProtocolIpv4Ping interface {
	Validation
	// msg marshals ActionResponseProtocolIpv4Ping to protobuf object *otg.ActionResponseProtocolIpv4Ping
	// and doesn't set defaults
	msg() *otg.ActionResponseProtocolIpv4Ping
	// setMsg unmarshals ActionResponseProtocolIpv4Ping from protobuf object *otg.ActionResponseProtocolIpv4Ping
	// and doesn't set defaults
	setMsg(*otg.ActionResponseProtocolIpv4Ping) ActionResponseProtocolIpv4Ping
	// provides marshal interface
	Marshal() marshalActionResponseProtocolIpv4Ping
	// provides unmarshal interface
	Unmarshal() unMarshalActionResponseProtocolIpv4Ping
	// validate validates ActionResponseProtocolIpv4Ping
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionResponseProtocolIpv4Ping, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Responses returns ActionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIterIter, set in ActionResponseProtocolIpv4Ping
	Responses() ActionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter
	setNil()
}

// List of responses for IPv4 ping responses.
// Responses returns a []ActionResponseProtocolIpv4PingResponse
func (obj *actionResponseProtocolIpv4Ping) Responses() ActionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter {
	if len(obj.obj.Responses) == 0 {
		obj.obj.Responses = []*otg.ActionResponseProtocolIpv4PingResponse{}
	}
	if obj.responsesHolder == nil {
		obj.responsesHolder = newActionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter(&obj.obj.Responses).setMsg(obj)
	}
	return obj.responsesHolder
}

type actionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter struct {
	obj                                         *actionResponseProtocolIpv4Ping
	actionResponseProtocolIpv4PingResponseSlice []ActionResponseProtocolIpv4PingResponse
	fieldPtr                                    *[]*otg.ActionResponseProtocolIpv4PingResponse
}

func newActionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter(ptr *[]*otg.ActionResponseProtocolIpv4PingResponse) ActionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter {
	return &actionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter{fieldPtr: ptr}
}

type ActionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter interface {
	setMsg(*actionResponseProtocolIpv4Ping) ActionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter
	Items() []ActionResponseProtocolIpv4PingResponse
	Add() ActionResponseProtocolIpv4PingResponse
	Append(items ...ActionResponseProtocolIpv4PingResponse) ActionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter
	Set(index int, newObj ActionResponseProtocolIpv4PingResponse) ActionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter
	Clear() ActionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter
	clearHolderSlice() ActionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter
	appendHolderSlice(item ActionResponseProtocolIpv4PingResponse) ActionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter
}

func (obj *actionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter) setMsg(msg *actionResponseProtocolIpv4Ping) ActionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&actionResponseProtocolIpv4PingResponse{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *actionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter) Items() []ActionResponseProtocolIpv4PingResponse {
	return obj.actionResponseProtocolIpv4PingResponseSlice
}

func (obj *actionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter) Add() ActionResponseProtocolIpv4PingResponse {
	newObj := &otg.ActionResponseProtocolIpv4PingResponse{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &actionResponseProtocolIpv4PingResponse{obj: newObj}
	newLibObj.setDefault()
	obj.actionResponseProtocolIpv4PingResponseSlice = append(obj.actionResponseProtocolIpv4PingResponseSlice, newLibObj)
	return newLibObj
}

func (obj *actionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter) Append(items ...ActionResponseProtocolIpv4PingResponse) ActionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.actionResponseProtocolIpv4PingResponseSlice = append(obj.actionResponseProtocolIpv4PingResponseSlice, item)
	}
	return obj
}

func (obj *actionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter) Set(index int, newObj ActionResponseProtocolIpv4PingResponse) ActionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.actionResponseProtocolIpv4PingResponseSlice[index] = newObj
	return obj
}
func (obj *actionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter) Clear() ActionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ActionResponseProtocolIpv4PingResponse{}
		obj.actionResponseProtocolIpv4PingResponseSlice = []ActionResponseProtocolIpv4PingResponse{}
	}
	return obj
}
func (obj *actionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter) clearHolderSlice() ActionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter {
	if len(obj.actionResponseProtocolIpv4PingResponseSlice) > 0 {
		obj.actionResponseProtocolIpv4PingResponseSlice = []ActionResponseProtocolIpv4PingResponse{}
	}
	return obj
}
func (obj *actionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter) appendHolderSlice(item ActionResponseProtocolIpv4PingResponse) ActionResponseProtocolIpv4PingActionResponseProtocolIpv4PingResponseIter {
	obj.actionResponseProtocolIpv4PingResponseSlice = append(obj.actionResponseProtocolIpv4PingResponseSlice, item)
	return obj
}

func (obj *actionResponseProtocolIpv4Ping) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Responses) != 0 {

		if set_default {
			obj.Responses().clearHolderSlice()
			for _, item := range obj.obj.Responses {
				obj.Responses().appendHolderSlice(&actionResponseProtocolIpv4PingResponse{obj: item})
			}
		}
		for _, item := range obj.Responses().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *actionResponseProtocolIpv4Ping) setDefault() {

}
