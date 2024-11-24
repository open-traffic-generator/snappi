package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionProtocolIpv6Ping *****
type actionProtocolIpv6Ping struct {
	validation
	obj            *otg.ActionProtocolIpv6Ping
	marshaller     marshalActionProtocolIpv6Ping
	unMarshaller   unMarshalActionProtocolIpv6Ping
	requestsHolder ActionProtocolIpv6PingActionProtocolIpv6PingRequestIter
}

func NewActionProtocolIpv6Ping() ActionProtocolIpv6Ping {
	obj := actionProtocolIpv6Ping{obj: &otg.ActionProtocolIpv6Ping{}}
	obj.setDefault()
	return &obj
}

func (obj *actionProtocolIpv6Ping) msg() *otg.ActionProtocolIpv6Ping {
	return obj.obj
}

func (obj *actionProtocolIpv6Ping) setMsg(msg *otg.ActionProtocolIpv6Ping) ActionProtocolIpv6Ping {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionProtocolIpv6Ping struct {
	obj *actionProtocolIpv6Ping
}

type marshalActionProtocolIpv6Ping interface {
	// ToProto marshals ActionProtocolIpv6Ping to protobuf object *otg.ActionProtocolIpv6Ping
	ToProto() (*otg.ActionProtocolIpv6Ping, error)
	// ToPbText marshals ActionProtocolIpv6Ping to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionProtocolIpv6Ping to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionProtocolIpv6Ping to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals ActionProtocolIpv6Ping to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalactionProtocolIpv6Ping struct {
	obj *actionProtocolIpv6Ping
}

type unMarshalActionProtocolIpv6Ping interface {
	// FromProto unmarshals ActionProtocolIpv6Ping from protobuf object *otg.ActionProtocolIpv6Ping
	FromProto(msg *otg.ActionProtocolIpv6Ping) (ActionProtocolIpv6Ping, error)
	// FromPbText unmarshals ActionProtocolIpv6Ping from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionProtocolIpv6Ping from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionProtocolIpv6Ping from JSON text
	FromJson(value string) error
}

func (obj *actionProtocolIpv6Ping) Marshal() marshalActionProtocolIpv6Ping {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionProtocolIpv6Ping{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionProtocolIpv6Ping) Unmarshal() unMarshalActionProtocolIpv6Ping {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionProtocolIpv6Ping{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionProtocolIpv6Ping) ToProto() (*otg.ActionProtocolIpv6Ping, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionProtocolIpv6Ping) FromProto(msg *otg.ActionProtocolIpv6Ping) (ActionProtocolIpv6Ping, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionProtocolIpv6Ping) ToPbText() (string, error) {
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

func (m *unMarshalactionProtocolIpv6Ping) FromPbText(value string) error {
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

func (m *marshalactionProtocolIpv6Ping) ToYaml() (string, error) {
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

func (m *unMarshalactionProtocolIpv6Ping) FromYaml(value string) error {
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

func (m *marshalactionProtocolIpv6Ping) ToJsonRaw() (string, error) {
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

func (m *marshalactionProtocolIpv6Ping) ToJson() (string, error) {
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

func (m *unMarshalactionProtocolIpv6Ping) FromJson(value string) error {
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

func (obj *actionProtocolIpv6Ping) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionProtocolIpv6Ping) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionProtocolIpv6Ping) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionProtocolIpv6Ping) Clone() (ActionProtocolIpv6Ping, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionProtocolIpv6Ping()
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

func (obj *actionProtocolIpv6Ping) setNil() {
	obj.requestsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ActionProtocolIpv6Ping is request for initiating ping between multiple source and destination pairs.
type ActionProtocolIpv6Ping interface {
	Validation
	// msg marshals ActionProtocolIpv6Ping to protobuf object *otg.ActionProtocolIpv6Ping
	// and doesn't set defaults
	msg() *otg.ActionProtocolIpv6Ping
	// setMsg unmarshals ActionProtocolIpv6Ping from protobuf object *otg.ActionProtocolIpv6Ping
	// and doesn't set defaults
	setMsg(*otg.ActionProtocolIpv6Ping) ActionProtocolIpv6Ping
	// provides marshal interface
	Marshal() marshalActionProtocolIpv6Ping
	// provides unmarshal interface
	Unmarshal() unMarshalActionProtocolIpv6Ping
	// validate validates ActionProtocolIpv6Ping
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionProtocolIpv6Ping, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Requests returns ActionProtocolIpv6PingActionProtocolIpv6PingRequestIterIter, set in ActionProtocolIpv6Ping
	Requests() ActionProtocolIpv6PingActionProtocolIpv6PingRequestIter
	setNil()
}

// List of IPv6 ping requests.
// Requests returns a []ActionProtocolIpv6PingRequest
func (obj *actionProtocolIpv6Ping) Requests() ActionProtocolIpv6PingActionProtocolIpv6PingRequestIter {
	if len(obj.obj.Requests) == 0 {
		obj.obj.Requests = []*otg.ActionProtocolIpv6PingRequest{}
	}
	if obj.requestsHolder == nil {
		obj.requestsHolder = newActionProtocolIpv6PingActionProtocolIpv6PingRequestIter(&obj.obj.Requests).setMsg(obj)
	}
	return obj.requestsHolder
}

type actionProtocolIpv6PingActionProtocolIpv6PingRequestIter struct {
	obj                                *actionProtocolIpv6Ping
	actionProtocolIpv6PingRequestSlice []ActionProtocolIpv6PingRequest
	fieldPtr                           *[]*otg.ActionProtocolIpv6PingRequest
}

func newActionProtocolIpv6PingActionProtocolIpv6PingRequestIter(ptr *[]*otg.ActionProtocolIpv6PingRequest) ActionProtocolIpv6PingActionProtocolIpv6PingRequestIter {
	return &actionProtocolIpv6PingActionProtocolIpv6PingRequestIter{fieldPtr: ptr}
}

type ActionProtocolIpv6PingActionProtocolIpv6PingRequestIter interface {
	setMsg(*actionProtocolIpv6Ping) ActionProtocolIpv6PingActionProtocolIpv6PingRequestIter
	Items() []ActionProtocolIpv6PingRequest
	Add() ActionProtocolIpv6PingRequest
	Append(items ...ActionProtocolIpv6PingRequest) ActionProtocolIpv6PingActionProtocolIpv6PingRequestIter
	Set(index int, newObj ActionProtocolIpv6PingRequest) ActionProtocolIpv6PingActionProtocolIpv6PingRequestIter
	Clear() ActionProtocolIpv6PingActionProtocolIpv6PingRequestIter
	clearHolderSlice() ActionProtocolIpv6PingActionProtocolIpv6PingRequestIter
	appendHolderSlice(item ActionProtocolIpv6PingRequest) ActionProtocolIpv6PingActionProtocolIpv6PingRequestIter
}

func (obj *actionProtocolIpv6PingActionProtocolIpv6PingRequestIter) setMsg(msg *actionProtocolIpv6Ping) ActionProtocolIpv6PingActionProtocolIpv6PingRequestIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&actionProtocolIpv6PingRequest{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *actionProtocolIpv6PingActionProtocolIpv6PingRequestIter) Items() []ActionProtocolIpv6PingRequest {
	return obj.actionProtocolIpv6PingRequestSlice
}

func (obj *actionProtocolIpv6PingActionProtocolIpv6PingRequestIter) Add() ActionProtocolIpv6PingRequest {
	newObj := &otg.ActionProtocolIpv6PingRequest{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &actionProtocolIpv6PingRequest{obj: newObj}
	newLibObj.setDefault()
	obj.actionProtocolIpv6PingRequestSlice = append(obj.actionProtocolIpv6PingRequestSlice, newLibObj)
	return newLibObj
}

func (obj *actionProtocolIpv6PingActionProtocolIpv6PingRequestIter) Append(items ...ActionProtocolIpv6PingRequest) ActionProtocolIpv6PingActionProtocolIpv6PingRequestIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.actionProtocolIpv6PingRequestSlice = append(obj.actionProtocolIpv6PingRequestSlice, item)
	}
	return obj
}

func (obj *actionProtocolIpv6PingActionProtocolIpv6PingRequestIter) Set(index int, newObj ActionProtocolIpv6PingRequest) ActionProtocolIpv6PingActionProtocolIpv6PingRequestIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.actionProtocolIpv6PingRequestSlice[index] = newObj
	return obj
}
func (obj *actionProtocolIpv6PingActionProtocolIpv6PingRequestIter) Clear() ActionProtocolIpv6PingActionProtocolIpv6PingRequestIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ActionProtocolIpv6PingRequest{}
		obj.actionProtocolIpv6PingRequestSlice = []ActionProtocolIpv6PingRequest{}
	}
	return obj
}
func (obj *actionProtocolIpv6PingActionProtocolIpv6PingRequestIter) clearHolderSlice() ActionProtocolIpv6PingActionProtocolIpv6PingRequestIter {
	if len(obj.actionProtocolIpv6PingRequestSlice) > 0 {
		obj.actionProtocolIpv6PingRequestSlice = []ActionProtocolIpv6PingRequest{}
	}
	return obj
}
func (obj *actionProtocolIpv6PingActionProtocolIpv6PingRequestIter) appendHolderSlice(item ActionProtocolIpv6PingRequest) ActionProtocolIpv6PingActionProtocolIpv6PingRequestIter {
	obj.actionProtocolIpv6PingRequestSlice = append(obj.actionProtocolIpv6PingRequestSlice, item)
	return obj
}

func (obj *actionProtocolIpv6Ping) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Requests) != 0 {

		if set_default {
			obj.Requests().clearHolderSlice()
			for _, item := range obj.obj.Requests {
				obj.Requests().appendHolderSlice(&actionProtocolIpv6PingRequest{obj: item})
			}
		}
		for _, item := range obj.Requests().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *actionProtocolIpv6Ping) setDefault() {

}
