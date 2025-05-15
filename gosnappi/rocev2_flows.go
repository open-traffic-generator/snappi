package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2Flows *****
type rocev2Flows struct {
	validation
	obj           *otg.Rocev2Flows
	marshaller    marshalRocev2Flows
	unMarshaller  unMarshalRocev2Flows
	txPortsHolder Rocev2FlowsRocev2TxPortsIter
}

func NewRocev2Flows() Rocev2Flows {
	obj := rocev2Flows{obj: &otg.Rocev2Flows{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2Flows) msg() *otg.Rocev2Flows {
	return obj.obj
}

func (obj *rocev2Flows) setMsg(msg *otg.Rocev2Flows) Rocev2Flows {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2Flows struct {
	obj *rocev2Flows
}

type marshalRocev2Flows interface {
	// ToProto marshals Rocev2Flows to protobuf object *otg.Rocev2Flows
	ToProto() (*otg.Rocev2Flows, error)
	// ToPbText marshals Rocev2Flows to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2Flows to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2Flows to JSON text
	ToJson() (string, error)
}

type unMarshalrocev2Flows struct {
	obj *rocev2Flows
}

type unMarshalRocev2Flows interface {
	// FromProto unmarshals Rocev2Flows from protobuf object *otg.Rocev2Flows
	FromProto(msg *otg.Rocev2Flows) (Rocev2Flows, error)
	// FromPbText unmarshals Rocev2Flows from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2Flows from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2Flows from JSON text
	FromJson(value string) error
}

func (obj *rocev2Flows) Marshal() marshalRocev2Flows {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2Flows{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2Flows) Unmarshal() unMarshalRocev2Flows {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2Flows{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2Flows) ToProto() (*otg.Rocev2Flows, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2Flows) FromProto(msg *otg.Rocev2Flows) (Rocev2Flows, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2Flows) ToPbText() (string, error) {
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

func (m *unMarshalrocev2Flows) FromPbText(value string) error {
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

func (m *marshalrocev2Flows) ToYaml() (string, error) {
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

func (m *unMarshalrocev2Flows) FromYaml(value string) error {
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

func (m *marshalrocev2Flows) ToJson() (string, error) {
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

func (m *unMarshalrocev2Flows) FromJson(value string) error {
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

func (obj *rocev2Flows) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2Flows) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2Flows) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2Flows) Clone() (Rocev2Flows, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2Flows()
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

func (obj *rocev2Flows) setNil() {
	obj.txPortsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Rocev2Flows is roCEv2 traffic flow configuration.
type Rocev2Flows interface {
	Validation
	// msg marshals Rocev2Flows to protobuf object *otg.Rocev2Flows
	// and doesn't set defaults
	msg() *otg.Rocev2Flows
	// setMsg unmarshals Rocev2Flows from protobuf object *otg.Rocev2Flows
	// and doesn't set defaults
	setMsg(*otg.Rocev2Flows) Rocev2Flows
	// provides marshal interface
	Marshal() marshalRocev2Flows
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2Flows
	// validate validates Rocev2Flows
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2Flows, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// TxPorts returns Rocev2FlowsRocev2TxPortsIterIter, set in Rocev2Flows
	TxPorts() Rocev2FlowsRocev2TxPortsIter
	setNil()
}

// Specifies the list of transmit (TX) ports used for sending RoCEv2 traffic.
// TxPorts returns a []Rocev2TxPorts
func (obj *rocev2Flows) TxPorts() Rocev2FlowsRocev2TxPortsIter {
	if len(obj.obj.TxPorts) == 0 {
		obj.obj.TxPorts = []*otg.Rocev2TxPorts{}
	}
	if obj.txPortsHolder == nil {
		obj.txPortsHolder = newRocev2FlowsRocev2TxPortsIter(&obj.obj.TxPorts).setMsg(obj)
	}
	return obj.txPortsHolder
}

type rocev2FlowsRocev2TxPortsIter struct {
	obj                *rocev2Flows
	rocev2TxPortsSlice []Rocev2TxPorts
	fieldPtr           *[]*otg.Rocev2TxPorts
}

func newRocev2FlowsRocev2TxPortsIter(ptr *[]*otg.Rocev2TxPorts) Rocev2FlowsRocev2TxPortsIter {
	return &rocev2FlowsRocev2TxPortsIter{fieldPtr: ptr}
}

type Rocev2FlowsRocev2TxPortsIter interface {
	setMsg(*rocev2Flows) Rocev2FlowsRocev2TxPortsIter
	Items() []Rocev2TxPorts
	Add() Rocev2TxPorts
	Append(items ...Rocev2TxPorts) Rocev2FlowsRocev2TxPortsIter
	Set(index int, newObj Rocev2TxPorts) Rocev2FlowsRocev2TxPortsIter
	Clear() Rocev2FlowsRocev2TxPortsIter
	clearHolderSlice() Rocev2FlowsRocev2TxPortsIter
	appendHolderSlice(item Rocev2TxPorts) Rocev2FlowsRocev2TxPortsIter
}

func (obj *rocev2FlowsRocev2TxPortsIter) setMsg(msg *rocev2Flows) Rocev2FlowsRocev2TxPortsIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&rocev2TxPorts{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *rocev2FlowsRocev2TxPortsIter) Items() []Rocev2TxPorts {
	return obj.rocev2TxPortsSlice
}

func (obj *rocev2FlowsRocev2TxPortsIter) Add() Rocev2TxPorts {
	newObj := &otg.Rocev2TxPorts{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &rocev2TxPorts{obj: newObj}
	newLibObj.setDefault()
	obj.rocev2TxPortsSlice = append(obj.rocev2TxPortsSlice, newLibObj)
	return newLibObj
}

func (obj *rocev2FlowsRocev2TxPortsIter) Append(items ...Rocev2TxPorts) Rocev2FlowsRocev2TxPortsIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.rocev2TxPortsSlice = append(obj.rocev2TxPortsSlice, item)
	}
	return obj
}

func (obj *rocev2FlowsRocev2TxPortsIter) Set(index int, newObj Rocev2TxPorts) Rocev2FlowsRocev2TxPortsIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.rocev2TxPortsSlice[index] = newObj
	return obj
}
func (obj *rocev2FlowsRocev2TxPortsIter) Clear() Rocev2FlowsRocev2TxPortsIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Rocev2TxPorts{}
		obj.rocev2TxPortsSlice = []Rocev2TxPorts{}
	}
	return obj
}
func (obj *rocev2FlowsRocev2TxPortsIter) clearHolderSlice() Rocev2FlowsRocev2TxPortsIter {
	if len(obj.rocev2TxPortsSlice) > 0 {
		obj.rocev2TxPortsSlice = []Rocev2TxPorts{}
	}
	return obj
}
func (obj *rocev2FlowsRocev2TxPortsIter) appendHolderSlice(item Rocev2TxPorts) Rocev2FlowsRocev2TxPortsIter {
	obj.rocev2TxPortsSlice = append(obj.rocev2TxPortsSlice, item)
	return obj
}

func (obj *rocev2Flows) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.TxPorts) != 0 {

		if set_default {
			obj.TxPorts().clearHolderSlice()
			for _, item := range obj.obj.TxPorts {
				obj.TxPorts().appendHolderSlice(&rocev2TxPorts{obj: item})
			}
		}
		for _, item := range obj.TxPorts().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *rocev2Flows) setDefault() {

}
