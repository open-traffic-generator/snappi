package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2Flow *****
type rocev2Flow struct {
	validation
	obj           *otg.Rocev2Flow
	marshaller    marshalRocev2Flow
	unMarshaller  unMarshalRocev2Flow
	txPortsHolder Rocev2FlowRocev2TxPortsIter
}

func NewRocev2Flow() Rocev2Flow {
	obj := rocev2Flow{obj: &otg.Rocev2Flow{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2Flow) msg() *otg.Rocev2Flow {
	return obj.obj
}

func (obj *rocev2Flow) setMsg(msg *otg.Rocev2Flow) Rocev2Flow {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2Flow struct {
	obj *rocev2Flow
}

type marshalRocev2Flow interface {
	// ToProto marshals Rocev2Flow to protobuf object *otg.Rocev2Flow
	ToProto() (*otg.Rocev2Flow, error)
	// ToPbText marshals Rocev2Flow to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2Flow to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2Flow to JSON text
	ToJson() (string, error)
}

type unMarshalrocev2Flow struct {
	obj *rocev2Flow
}

type unMarshalRocev2Flow interface {
	// FromProto unmarshals Rocev2Flow from protobuf object *otg.Rocev2Flow
	FromProto(msg *otg.Rocev2Flow) (Rocev2Flow, error)
	// FromPbText unmarshals Rocev2Flow from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2Flow from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2Flow from JSON text
	FromJson(value string) error
}

func (obj *rocev2Flow) Marshal() marshalRocev2Flow {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2Flow{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2Flow) Unmarshal() unMarshalRocev2Flow {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2Flow{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2Flow) ToProto() (*otg.Rocev2Flow, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2Flow) FromProto(msg *otg.Rocev2Flow) (Rocev2Flow, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2Flow) ToPbText() (string, error) {
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

func (m *unMarshalrocev2Flow) FromPbText(value string) error {
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

func (m *marshalrocev2Flow) ToYaml() (string, error) {
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

func (m *unMarshalrocev2Flow) FromYaml(value string) error {
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

func (m *marshalrocev2Flow) ToJson() (string, error) {
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

func (m *unMarshalrocev2Flow) FromJson(value string) error {
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

func (obj *rocev2Flow) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2Flow) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2Flow) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2Flow) Clone() (Rocev2Flow, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2Flow()
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

func (obj *rocev2Flow) setNil() {
	obj.txPortsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Rocev2Flow is rocev2 traffic flow configuration.
type Rocev2Flow interface {
	Validation
	// msg marshals Rocev2Flow to protobuf object *otg.Rocev2Flow
	// and doesn't set defaults
	msg() *otg.Rocev2Flow
	// setMsg unmarshals Rocev2Flow from protobuf object *otg.Rocev2Flow
	// and doesn't set defaults
	setMsg(*otg.Rocev2Flow) Rocev2Flow
	// provides marshal interface
	Marshal() marshalRocev2Flow
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2Flow
	// validate validates Rocev2Flow
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2Flow, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// TxPorts returns Rocev2FlowRocev2TxPortsIterIter, set in Rocev2Flow
	TxPorts() Rocev2FlowRocev2TxPortsIter
	setNil()
}

// ****Must Add proper description here****
// TxPorts returns a []Rocev2TxPorts
func (obj *rocev2Flow) TxPorts() Rocev2FlowRocev2TxPortsIter {
	if len(obj.obj.TxPorts) == 0 {
		obj.obj.TxPorts = []*otg.Rocev2TxPorts{}
	}
	if obj.txPortsHolder == nil {
		obj.txPortsHolder = newRocev2FlowRocev2TxPortsIter(&obj.obj.TxPorts).setMsg(obj)
	}
	return obj.txPortsHolder
}

type rocev2FlowRocev2TxPortsIter struct {
	obj                *rocev2Flow
	rocev2TxPortsSlice []Rocev2TxPorts
	fieldPtr           *[]*otg.Rocev2TxPorts
}

func newRocev2FlowRocev2TxPortsIter(ptr *[]*otg.Rocev2TxPorts) Rocev2FlowRocev2TxPortsIter {
	return &rocev2FlowRocev2TxPortsIter{fieldPtr: ptr}
}

type Rocev2FlowRocev2TxPortsIter interface {
	setMsg(*rocev2Flow) Rocev2FlowRocev2TxPortsIter
	Items() []Rocev2TxPorts
	Add() Rocev2TxPorts
	Append(items ...Rocev2TxPorts) Rocev2FlowRocev2TxPortsIter
	Set(index int, newObj Rocev2TxPorts) Rocev2FlowRocev2TxPortsIter
	Clear() Rocev2FlowRocev2TxPortsIter
	clearHolderSlice() Rocev2FlowRocev2TxPortsIter
	appendHolderSlice(item Rocev2TxPorts) Rocev2FlowRocev2TxPortsIter
}

func (obj *rocev2FlowRocev2TxPortsIter) setMsg(msg *rocev2Flow) Rocev2FlowRocev2TxPortsIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&rocev2TxPorts{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *rocev2FlowRocev2TxPortsIter) Items() []Rocev2TxPorts {
	return obj.rocev2TxPortsSlice
}

func (obj *rocev2FlowRocev2TxPortsIter) Add() Rocev2TxPorts {
	newObj := &otg.Rocev2TxPorts{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &rocev2TxPorts{obj: newObj}
	newLibObj.setDefault()
	obj.rocev2TxPortsSlice = append(obj.rocev2TxPortsSlice, newLibObj)
	return newLibObj
}

func (obj *rocev2FlowRocev2TxPortsIter) Append(items ...Rocev2TxPorts) Rocev2FlowRocev2TxPortsIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.rocev2TxPortsSlice = append(obj.rocev2TxPortsSlice, item)
	}
	return obj
}

func (obj *rocev2FlowRocev2TxPortsIter) Set(index int, newObj Rocev2TxPorts) Rocev2FlowRocev2TxPortsIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.rocev2TxPortsSlice[index] = newObj
	return obj
}
func (obj *rocev2FlowRocev2TxPortsIter) Clear() Rocev2FlowRocev2TxPortsIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Rocev2TxPorts{}
		obj.rocev2TxPortsSlice = []Rocev2TxPorts{}
	}
	return obj
}
func (obj *rocev2FlowRocev2TxPortsIter) clearHolderSlice() Rocev2FlowRocev2TxPortsIter {
	if len(obj.rocev2TxPortsSlice) > 0 {
		obj.rocev2TxPortsSlice = []Rocev2TxPorts{}
	}
	return obj
}
func (obj *rocev2FlowRocev2TxPortsIter) appendHolderSlice(item Rocev2TxPorts) Rocev2FlowRocev2TxPortsIter {
	obj.rocev2TxPortsSlice = append(obj.rocev2TxPortsSlice, item)
	return obj
}

func (obj *rocev2Flow) validateObj(vObj *validation, set_default bool) {
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

func (obj *rocev2Flow) setDefault() {

}
