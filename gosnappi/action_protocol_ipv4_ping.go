package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionProtocolIpv4Ping *****
type actionProtocolIpv4Ping struct {
	validation
	obj            *otg.ActionProtocolIpv4Ping
	marshaller     marshalActionProtocolIpv4Ping
	unMarshaller   unMarshalActionProtocolIpv4Ping
	requestsHolder ActionProtocolIpv4PingActionProtocolIpv4PingRequestIter
}

func NewActionProtocolIpv4Ping() ActionProtocolIpv4Ping {
	obj := actionProtocolIpv4Ping{obj: &otg.ActionProtocolIpv4Ping{}}
	obj.setDefault()
	return &obj
}

func (obj *actionProtocolIpv4Ping) msg() *otg.ActionProtocolIpv4Ping {
	return obj.obj
}

func (obj *actionProtocolIpv4Ping) setMsg(msg *otg.ActionProtocolIpv4Ping) ActionProtocolIpv4Ping {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionProtocolIpv4Ping struct {
	obj *actionProtocolIpv4Ping
}

type marshalActionProtocolIpv4Ping interface {
	// ToProto marshals ActionProtocolIpv4Ping to protobuf object *otg.ActionProtocolIpv4Ping
	ToProto() (*otg.ActionProtocolIpv4Ping, error)
	// ToPbText marshals ActionProtocolIpv4Ping to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionProtocolIpv4Ping to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionProtocolIpv4Ping to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals ActionProtocolIpv4Ping to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalactionProtocolIpv4Ping struct {
	obj *actionProtocolIpv4Ping
}

type unMarshalActionProtocolIpv4Ping interface {
	// FromProto unmarshals ActionProtocolIpv4Ping from protobuf object *otg.ActionProtocolIpv4Ping
	FromProto(msg *otg.ActionProtocolIpv4Ping) (ActionProtocolIpv4Ping, error)
	// FromPbText unmarshals ActionProtocolIpv4Ping from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionProtocolIpv4Ping from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionProtocolIpv4Ping from JSON text
	FromJson(value string) error
}

func (obj *actionProtocolIpv4Ping) Marshal() marshalActionProtocolIpv4Ping {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionProtocolIpv4Ping{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionProtocolIpv4Ping) Unmarshal() unMarshalActionProtocolIpv4Ping {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionProtocolIpv4Ping{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionProtocolIpv4Ping) ToProto() (*otg.ActionProtocolIpv4Ping, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionProtocolIpv4Ping) FromProto(msg *otg.ActionProtocolIpv4Ping) (ActionProtocolIpv4Ping, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionProtocolIpv4Ping) ToPbText() (string, error) {
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

func (m *unMarshalactionProtocolIpv4Ping) FromPbText(value string) error {
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

func (m *marshalactionProtocolIpv4Ping) ToYaml() (string, error) {
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

func (m *unMarshalactionProtocolIpv4Ping) FromYaml(value string) error {
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

func (m *marshalactionProtocolIpv4Ping) ToJsonRaw() (string, error) {
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

func (m *marshalactionProtocolIpv4Ping) ToJson() (string, error) {
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

func (m *unMarshalactionProtocolIpv4Ping) FromJson(value string) error {
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

func (obj *actionProtocolIpv4Ping) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionProtocolIpv4Ping) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionProtocolIpv4Ping) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionProtocolIpv4Ping) Clone() (ActionProtocolIpv4Ping, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionProtocolIpv4Ping()
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

func (obj *actionProtocolIpv4Ping) setNil() {
	obj.requestsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ActionProtocolIpv4Ping is request for initiating ping between multiple source and destination pairs.
type ActionProtocolIpv4Ping interface {
	Validation
	// msg marshals ActionProtocolIpv4Ping to protobuf object *otg.ActionProtocolIpv4Ping
	// and doesn't set defaults
	msg() *otg.ActionProtocolIpv4Ping
	// setMsg unmarshals ActionProtocolIpv4Ping from protobuf object *otg.ActionProtocolIpv4Ping
	// and doesn't set defaults
	setMsg(*otg.ActionProtocolIpv4Ping) ActionProtocolIpv4Ping
	// provides marshal interface
	Marshal() marshalActionProtocolIpv4Ping
	// provides unmarshal interface
	Unmarshal() unMarshalActionProtocolIpv4Ping
	// validate validates ActionProtocolIpv4Ping
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionProtocolIpv4Ping, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Requests returns ActionProtocolIpv4PingActionProtocolIpv4PingRequestIterIter, set in ActionProtocolIpv4Ping
	Requests() ActionProtocolIpv4PingActionProtocolIpv4PingRequestIter
	setNil()
}

// List of IPv4 ping requests.
// Requests returns a []ActionProtocolIpv4PingRequest
func (obj *actionProtocolIpv4Ping) Requests() ActionProtocolIpv4PingActionProtocolIpv4PingRequestIter {
	if len(obj.obj.Requests) == 0 {
		obj.obj.Requests = []*otg.ActionProtocolIpv4PingRequest{}
	}
	if obj.requestsHolder == nil {
		obj.requestsHolder = newActionProtocolIpv4PingActionProtocolIpv4PingRequestIter(&obj.obj.Requests).setMsg(obj)
	}
	return obj.requestsHolder
}

type actionProtocolIpv4PingActionProtocolIpv4PingRequestIter struct {
	obj                                *actionProtocolIpv4Ping
	actionProtocolIpv4PingRequestSlice []ActionProtocolIpv4PingRequest
	fieldPtr                           *[]*otg.ActionProtocolIpv4PingRequest
}

func newActionProtocolIpv4PingActionProtocolIpv4PingRequestIter(ptr *[]*otg.ActionProtocolIpv4PingRequest) ActionProtocolIpv4PingActionProtocolIpv4PingRequestIter {
	return &actionProtocolIpv4PingActionProtocolIpv4PingRequestIter{fieldPtr: ptr}
}

type ActionProtocolIpv4PingActionProtocolIpv4PingRequestIter interface {
	setMsg(*actionProtocolIpv4Ping) ActionProtocolIpv4PingActionProtocolIpv4PingRequestIter
	Items() []ActionProtocolIpv4PingRequest
	Add() ActionProtocolIpv4PingRequest
	Append(items ...ActionProtocolIpv4PingRequest) ActionProtocolIpv4PingActionProtocolIpv4PingRequestIter
	Set(index int, newObj ActionProtocolIpv4PingRequest) ActionProtocolIpv4PingActionProtocolIpv4PingRequestIter
	Clear() ActionProtocolIpv4PingActionProtocolIpv4PingRequestIter
	clearHolderSlice() ActionProtocolIpv4PingActionProtocolIpv4PingRequestIter
	appendHolderSlice(item ActionProtocolIpv4PingRequest) ActionProtocolIpv4PingActionProtocolIpv4PingRequestIter
}

func (obj *actionProtocolIpv4PingActionProtocolIpv4PingRequestIter) setMsg(msg *actionProtocolIpv4Ping) ActionProtocolIpv4PingActionProtocolIpv4PingRequestIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&actionProtocolIpv4PingRequest{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *actionProtocolIpv4PingActionProtocolIpv4PingRequestIter) Items() []ActionProtocolIpv4PingRequest {
	return obj.actionProtocolIpv4PingRequestSlice
}

func (obj *actionProtocolIpv4PingActionProtocolIpv4PingRequestIter) Add() ActionProtocolIpv4PingRequest {
	newObj := &otg.ActionProtocolIpv4PingRequest{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &actionProtocolIpv4PingRequest{obj: newObj}
	newLibObj.setDefault()
	obj.actionProtocolIpv4PingRequestSlice = append(obj.actionProtocolIpv4PingRequestSlice, newLibObj)
	return newLibObj
}

func (obj *actionProtocolIpv4PingActionProtocolIpv4PingRequestIter) Append(items ...ActionProtocolIpv4PingRequest) ActionProtocolIpv4PingActionProtocolIpv4PingRequestIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.actionProtocolIpv4PingRequestSlice = append(obj.actionProtocolIpv4PingRequestSlice, item)
	}
	return obj
}

func (obj *actionProtocolIpv4PingActionProtocolIpv4PingRequestIter) Set(index int, newObj ActionProtocolIpv4PingRequest) ActionProtocolIpv4PingActionProtocolIpv4PingRequestIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.actionProtocolIpv4PingRequestSlice[index] = newObj
	return obj
}
func (obj *actionProtocolIpv4PingActionProtocolIpv4PingRequestIter) Clear() ActionProtocolIpv4PingActionProtocolIpv4PingRequestIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ActionProtocolIpv4PingRequest{}
		obj.actionProtocolIpv4PingRequestSlice = []ActionProtocolIpv4PingRequest{}
	}
	return obj
}
func (obj *actionProtocolIpv4PingActionProtocolIpv4PingRequestIter) clearHolderSlice() ActionProtocolIpv4PingActionProtocolIpv4PingRequestIter {
	if len(obj.actionProtocolIpv4PingRequestSlice) > 0 {
		obj.actionProtocolIpv4PingRequestSlice = []ActionProtocolIpv4PingRequest{}
	}
	return obj
}
func (obj *actionProtocolIpv4PingActionProtocolIpv4PingRequestIter) appendHolderSlice(item ActionProtocolIpv4PingRequest) ActionProtocolIpv4PingActionProtocolIpv4PingRequestIter {
	obj.actionProtocolIpv4PingRequestSlice = append(obj.actionProtocolIpv4PingRequestSlice, item)
	return obj
}

func (obj *actionProtocolIpv4Ping) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Requests) != 0 {

		if set_default {
			obj.Requests().clearHolderSlice()
			for _, item := range obj.obj.Requests {
				obj.Requests().appendHolderSlice(&actionProtocolIpv4PingRequest{obj: item})
			}
		}
		for _, item := range obj.Requests().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *actionProtocolIpv4Ping) setDefault() {

}
