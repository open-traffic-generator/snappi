package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionResponseProtocolIpv6Ping *****
type actionResponseProtocolIpv6Ping struct {
	validation
	obj             *otg.ActionResponseProtocolIpv6Ping
	marshaller      marshalActionResponseProtocolIpv6Ping
	unMarshaller    unMarshalActionResponseProtocolIpv6Ping
	responsesHolder ActionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter
}

func NewActionResponseProtocolIpv6Ping() ActionResponseProtocolIpv6Ping {
	obj := actionResponseProtocolIpv6Ping{obj: &otg.ActionResponseProtocolIpv6Ping{}}
	obj.setDefault()
	return &obj
}

func (obj *actionResponseProtocolIpv6Ping) msg() *otg.ActionResponseProtocolIpv6Ping {
	return obj.obj
}

func (obj *actionResponseProtocolIpv6Ping) setMsg(msg *otg.ActionResponseProtocolIpv6Ping) ActionResponseProtocolIpv6Ping {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionResponseProtocolIpv6Ping struct {
	obj *actionResponseProtocolIpv6Ping
}

type marshalActionResponseProtocolIpv6Ping interface {
	// ToProto marshals ActionResponseProtocolIpv6Ping to protobuf object *otg.ActionResponseProtocolIpv6Ping
	ToProto() (*otg.ActionResponseProtocolIpv6Ping, error)
	// ToPbText marshals ActionResponseProtocolIpv6Ping to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionResponseProtocolIpv6Ping to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionResponseProtocolIpv6Ping to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals ActionResponseProtocolIpv6Ping to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalactionResponseProtocolIpv6Ping struct {
	obj *actionResponseProtocolIpv6Ping
}

type unMarshalActionResponseProtocolIpv6Ping interface {
	// FromProto unmarshals ActionResponseProtocolIpv6Ping from protobuf object *otg.ActionResponseProtocolIpv6Ping
	FromProto(msg *otg.ActionResponseProtocolIpv6Ping) (ActionResponseProtocolIpv6Ping, error)
	// FromPbText unmarshals ActionResponseProtocolIpv6Ping from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionResponseProtocolIpv6Ping from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionResponseProtocolIpv6Ping from JSON text
	FromJson(value string) error
}

func (obj *actionResponseProtocolIpv6Ping) Marshal() marshalActionResponseProtocolIpv6Ping {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionResponseProtocolIpv6Ping{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionResponseProtocolIpv6Ping) Unmarshal() unMarshalActionResponseProtocolIpv6Ping {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionResponseProtocolIpv6Ping{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionResponseProtocolIpv6Ping) ToProto() (*otg.ActionResponseProtocolIpv6Ping, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionResponseProtocolIpv6Ping) FromProto(msg *otg.ActionResponseProtocolIpv6Ping) (ActionResponseProtocolIpv6Ping, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionResponseProtocolIpv6Ping) ToPbText() (string, error) {
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

func (m *unMarshalactionResponseProtocolIpv6Ping) FromPbText(value string) error {
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

func (m *marshalactionResponseProtocolIpv6Ping) ToYaml() (string, error) {
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

func (m *unMarshalactionResponseProtocolIpv6Ping) FromYaml(value string) error {
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

func (m *marshalactionResponseProtocolIpv6Ping) ToJsonRaw() (string, error) {
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

func (m *marshalactionResponseProtocolIpv6Ping) ToJson() (string, error) {
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

func (m *unMarshalactionResponseProtocolIpv6Ping) FromJson(value string) error {
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

func (obj *actionResponseProtocolIpv6Ping) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionResponseProtocolIpv6Ping) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionResponseProtocolIpv6Ping) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionResponseProtocolIpv6Ping) Clone() (ActionResponseProtocolIpv6Ping, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionResponseProtocolIpv6Ping()
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

func (obj *actionResponseProtocolIpv6Ping) setNil() {
	obj.responsesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ActionResponseProtocolIpv6Ping is response for ping initiated between multiple source and destination pairs.
type ActionResponseProtocolIpv6Ping interface {
	Validation
	// msg marshals ActionResponseProtocolIpv6Ping to protobuf object *otg.ActionResponseProtocolIpv6Ping
	// and doesn't set defaults
	msg() *otg.ActionResponseProtocolIpv6Ping
	// setMsg unmarshals ActionResponseProtocolIpv6Ping from protobuf object *otg.ActionResponseProtocolIpv6Ping
	// and doesn't set defaults
	setMsg(*otg.ActionResponseProtocolIpv6Ping) ActionResponseProtocolIpv6Ping
	// provides marshal interface
	Marshal() marshalActionResponseProtocolIpv6Ping
	// provides unmarshal interface
	Unmarshal() unMarshalActionResponseProtocolIpv6Ping
	// validate validates ActionResponseProtocolIpv6Ping
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionResponseProtocolIpv6Ping, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Responses returns ActionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIterIter, set in ActionResponseProtocolIpv6Ping
	Responses() ActionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter
	setNil()
}

// List of responses for IPv6 ping responses.
// Responses returns a []ActionResponseProtocolIpv6PingResponse
func (obj *actionResponseProtocolIpv6Ping) Responses() ActionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter {
	if len(obj.obj.Responses) == 0 {
		obj.obj.Responses = []*otg.ActionResponseProtocolIpv6PingResponse{}
	}
	if obj.responsesHolder == nil {
		obj.responsesHolder = newActionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter(&obj.obj.Responses).setMsg(obj)
	}
	return obj.responsesHolder
}

type actionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter struct {
	obj                                         *actionResponseProtocolIpv6Ping
	actionResponseProtocolIpv6PingResponseSlice []ActionResponseProtocolIpv6PingResponse
	fieldPtr                                    *[]*otg.ActionResponseProtocolIpv6PingResponse
}

func newActionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter(ptr *[]*otg.ActionResponseProtocolIpv6PingResponse) ActionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter {
	return &actionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter{fieldPtr: ptr}
}

type ActionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter interface {
	setMsg(*actionResponseProtocolIpv6Ping) ActionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter
	Items() []ActionResponseProtocolIpv6PingResponse
	Add() ActionResponseProtocolIpv6PingResponse
	Append(items ...ActionResponseProtocolIpv6PingResponse) ActionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter
	Set(index int, newObj ActionResponseProtocolIpv6PingResponse) ActionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter
	Clear() ActionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter
	clearHolderSlice() ActionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter
	appendHolderSlice(item ActionResponseProtocolIpv6PingResponse) ActionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter
}

func (obj *actionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter) setMsg(msg *actionResponseProtocolIpv6Ping) ActionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&actionResponseProtocolIpv6PingResponse{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *actionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter) Items() []ActionResponseProtocolIpv6PingResponse {
	return obj.actionResponseProtocolIpv6PingResponseSlice
}

func (obj *actionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter) Add() ActionResponseProtocolIpv6PingResponse {
	newObj := &otg.ActionResponseProtocolIpv6PingResponse{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &actionResponseProtocolIpv6PingResponse{obj: newObj}
	newLibObj.setDefault()
	obj.actionResponseProtocolIpv6PingResponseSlice = append(obj.actionResponseProtocolIpv6PingResponseSlice, newLibObj)
	return newLibObj
}

func (obj *actionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter) Append(items ...ActionResponseProtocolIpv6PingResponse) ActionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.actionResponseProtocolIpv6PingResponseSlice = append(obj.actionResponseProtocolIpv6PingResponseSlice, item)
	}
	return obj
}

func (obj *actionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter) Set(index int, newObj ActionResponseProtocolIpv6PingResponse) ActionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.actionResponseProtocolIpv6PingResponseSlice[index] = newObj
	return obj
}
func (obj *actionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter) Clear() ActionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ActionResponseProtocolIpv6PingResponse{}
		obj.actionResponseProtocolIpv6PingResponseSlice = []ActionResponseProtocolIpv6PingResponse{}
	}
	return obj
}
func (obj *actionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter) clearHolderSlice() ActionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter {
	if len(obj.actionResponseProtocolIpv6PingResponseSlice) > 0 {
		obj.actionResponseProtocolIpv6PingResponseSlice = []ActionResponseProtocolIpv6PingResponse{}
	}
	return obj
}
func (obj *actionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter) appendHolderSlice(item ActionResponseProtocolIpv6PingResponse) ActionResponseProtocolIpv6PingActionResponseProtocolIpv6PingResponseIter {
	obj.actionResponseProtocolIpv6PingResponseSlice = append(obj.actionResponseProtocolIpv6PingResponseSlice, item)
	return obj
}

func (obj *actionResponseProtocolIpv6Ping) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Responses) != 0 {

		if set_default {
			obj.Responses().clearHolderSlice()
			for _, item := range obj.obj.Responses {
				obj.Responses().appendHolderSlice(&actionResponseProtocolIpv6PingResponse{obj: item})
			}
		}
		for _, item := range obj.Responses().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *actionResponseProtocolIpv6Ping) setDefault() {

}
