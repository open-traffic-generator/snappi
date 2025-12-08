package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** CapabilitiesResponse *****
type capabilitiesResponse struct {
	validation
	obj                           *otg.CapabilitiesResponse
	marshaller                    marshalCapabilitiesResponse
	unMarshaller                  unMarshalCapabilitiesResponse
	perPortTypeCapabilitiesHolder CapabilitiesResponseCapabilitiesPortTypeCapabilityIter
}

func NewCapabilitiesResponse() CapabilitiesResponse {
	obj := capabilitiesResponse{obj: &otg.CapabilitiesResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *capabilitiesResponse) msg() *otg.CapabilitiesResponse {
	return obj.obj
}

func (obj *capabilitiesResponse) setMsg(msg *otg.CapabilitiesResponse) CapabilitiesResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalcapabilitiesResponse struct {
	obj *capabilitiesResponse
}

type marshalCapabilitiesResponse interface {
	// ToProto marshals CapabilitiesResponse to protobuf object *otg.CapabilitiesResponse
	ToProto() (*otg.CapabilitiesResponse, error)
	// ToPbText marshals CapabilitiesResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals CapabilitiesResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals CapabilitiesResponse to JSON text
	ToJson() (string, error)
}

type unMarshalcapabilitiesResponse struct {
	obj *capabilitiesResponse
}

type unMarshalCapabilitiesResponse interface {
	// FromProto unmarshals CapabilitiesResponse from protobuf object *otg.CapabilitiesResponse
	FromProto(msg *otg.CapabilitiesResponse) (CapabilitiesResponse, error)
	// FromPbText unmarshals CapabilitiesResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals CapabilitiesResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals CapabilitiesResponse from JSON text
	FromJson(value string) error
}

func (obj *capabilitiesResponse) Marshal() marshalCapabilitiesResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalcapabilitiesResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *capabilitiesResponse) Unmarshal() unMarshalCapabilitiesResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalcapabilitiesResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalcapabilitiesResponse) ToProto() (*otg.CapabilitiesResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalcapabilitiesResponse) FromProto(msg *otg.CapabilitiesResponse) (CapabilitiesResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalcapabilitiesResponse) ToPbText() (string, error) {
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

func (m *unMarshalcapabilitiesResponse) FromPbText(value string) error {
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

func (m *marshalcapabilitiesResponse) ToYaml() (string, error) {
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

func (m *unMarshalcapabilitiesResponse) FromYaml(value string) error {
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

func (m *marshalcapabilitiesResponse) ToJson() (string, error) {
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

func (m *unMarshalcapabilitiesResponse) FromJson(value string) error {
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

func (obj *capabilitiesResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *capabilitiesResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *capabilitiesResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *capabilitiesResponse) Clone() (CapabilitiesResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewCapabilitiesResponse()
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

func (obj *capabilitiesResponse) setNil() {
	obj.perPortTypeCapabilitiesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// CapabilitiesResponse is container for Capabiities Response from the traffic generator implementation.
type CapabilitiesResponse interface {
	Validation
	// msg marshals CapabilitiesResponse to protobuf object *otg.CapabilitiesResponse
	// and doesn't set defaults
	msg() *otg.CapabilitiesResponse
	// setMsg unmarshals CapabilitiesResponse from protobuf object *otg.CapabilitiesResponse
	// and doesn't set defaults
	setMsg(*otg.CapabilitiesResponse) CapabilitiesResponse
	// provides marshal interface
	Marshal() marshalCapabilitiesResponse
	// provides unmarshal interface
	Unmarshal() unMarshalCapabilitiesResponse
	// validate validates CapabilitiesResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (CapabilitiesResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PerPortTypeCapabilities returns CapabilitiesResponseCapabilitiesPortTypeCapabilityIterIter, set in CapabilitiesResponse
	PerPortTypeCapabilities() CapabilitiesResponseCapabilitiesPortTypeCapabilityIter
	setNil()
}

// Implementations should try to compress the returned capabilities information in terms of the following:
// 1) If the implementations support multiple types of port_types with different capabilities but is able to deduce that the current deployment can only support one or a sub-set of these port-types,
// it should return the capabilties only of the sub-set of port_types supported by the current deployment of the implementation to which the request is made, unless specifically requested for.
// 2) If the request_type is of included_paths , it should include only the top-level nodes under which all supported objects and attributes are supported.
// However, if only some of the attributes or child objects are supported under a node in the hierarchy , it should explicitly include all those child objects , specifically excluding the unsupported ones.
// 3) If the request_type is of excluded_paths , it should include only the top-level nodes under which all supported objects and attributes are unsupported.
// However, if only some of the attributes or child objects are unsupported under a node in the hierarchy , it should explicitly include all those child objects , specifically excluding the supported ones.
// PerPortTypeCapabilities returns a []CapabilitiesPortTypeCapability
func (obj *capabilitiesResponse) PerPortTypeCapabilities() CapabilitiesResponseCapabilitiesPortTypeCapabilityIter {
	if len(obj.obj.PerPortTypeCapabilities) == 0 {
		obj.obj.PerPortTypeCapabilities = []*otg.CapabilitiesPortTypeCapability{}
	}
	if obj.perPortTypeCapabilitiesHolder == nil {
		obj.perPortTypeCapabilitiesHolder = newCapabilitiesResponseCapabilitiesPortTypeCapabilityIter(&obj.obj.PerPortTypeCapabilities).setMsg(obj)
	}
	return obj.perPortTypeCapabilitiesHolder
}

type capabilitiesResponseCapabilitiesPortTypeCapabilityIter struct {
	obj                                 *capabilitiesResponse
	capabilitiesPortTypeCapabilitySlice []CapabilitiesPortTypeCapability
	fieldPtr                            *[]*otg.CapabilitiesPortTypeCapability
}

func newCapabilitiesResponseCapabilitiesPortTypeCapabilityIter(ptr *[]*otg.CapabilitiesPortTypeCapability) CapabilitiesResponseCapabilitiesPortTypeCapabilityIter {
	return &capabilitiesResponseCapabilitiesPortTypeCapabilityIter{fieldPtr: ptr}
}

type CapabilitiesResponseCapabilitiesPortTypeCapabilityIter interface {
	setMsg(*capabilitiesResponse) CapabilitiesResponseCapabilitiesPortTypeCapabilityIter
	Items() []CapabilitiesPortTypeCapability
	Add() CapabilitiesPortTypeCapability
	Append(items ...CapabilitiesPortTypeCapability) CapabilitiesResponseCapabilitiesPortTypeCapabilityIter
	Set(index int, newObj CapabilitiesPortTypeCapability) CapabilitiesResponseCapabilitiesPortTypeCapabilityIter
	Clear() CapabilitiesResponseCapabilitiesPortTypeCapabilityIter
	clearHolderSlice() CapabilitiesResponseCapabilitiesPortTypeCapabilityIter
	appendHolderSlice(item CapabilitiesPortTypeCapability) CapabilitiesResponseCapabilitiesPortTypeCapabilityIter
}

func (obj *capabilitiesResponseCapabilitiesPortTypeCapabilityIter) setMsg(msg *capabilitiesResponse) CapabilitiesResponseCapabilitiesPortTypeCapabilityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&capabilitiesPortTypeCapability{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *capabilitiesResponseCapabilitiesPortTypeCapabilityIter) Items() []CapabilitiesPortTypeCapability {
	return obj.capabilitiesPortTypeCapabilitySlice
}

func (obj *capabilitiesResponseCapabilitiesPortTypeCapabilityIter) Add() CapabilitiesPortTypeCapability {
	newObj := &otg.CapabilitiesPortTypeCapability{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &capabilitiesPortTypeCapability{obj: newObj}
	newLibObj.setDefault()
	obj.capabilitiesPortTypeCapabilitySlice = append(obj.capabilitiesPortTypeCapabilitySlice, newLibObj)
	return newLibObj
}

func (obj *capabilitiesResponseCapabilitiesPortTypeCapabilityIter) Append(items ...CapabilitiesPortTypeCapability) CapabilitiesResponseCapabilitiesPortTypeCapabilityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.capabilitiesPortTypeCapabilitySlice = append(obj.capabilitiesPortTypeCapabilitySlice, item)
	}
	return obj
}

func (obj *capabilitiesResponseCapabilitiesPortTypeCapabilityIter) Set(index int, newObj CapabilitiesPortTypeCapability) CapabilitiesResponseCapabilitiesPortTypeCapabilityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.capabilitiesPortTypeCapabilitySlice[index] = newObj
	return obj
}
func (obj *capabilitiesResponseCapabilitiesPortTypeCapabilityIter) Clear() CapabilitiesResponseCapabilitiesPortTypeCapabilityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.CapabilitiesPortTypeCapability{}
		obj.capabilitiesPortTypeCapabilitySlice = []CapabilitiesPortTypeCapability{}
	}
	return obj
}
func (obj *capabilitiesResponseCapabilitiesPortTypeCapabilityIter) clearHolderSlice() CapabilitiesResponseCapabilitiesPortTypeCapabilityIter {
	if len(obj.capabilitiesPortTypeCapabilitySlice) > 0 {
		obj.capabilitiesPortTypeCapabilitySlice = []CapabilitiesPortTypeCapability{}
	}
	return obj
}
func (obj *capabilitiesResponseCapabilitiesPortTypeCapabilityIter) appendHolderSlice(item CapabilitiesPortTypeCapability) CapabilitiesResponseCapabilitiesPortTypeCapabilityIter {
	obj.capabilitiesPortTypeCapabilitySlice = append(obj.capabilitiesPortTypeCapabilitySlice, item)
	return obj
}

func (obj *capabilitiesResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.PerPortTypeCapabilities) != 0 {

		if set_default {
			obj.PerPortTypeCapabilities().clearHolderSlice()
			for _, item := range obj.obj.PerPortTypeCapabilities {
				obj.PerPortTypeCapabilities().appendHolderSlice(&capabilitiesPortTypeCapability{obj: item})
			}
		}
		for _, item := range obj.PerPortTypeCapabilities().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *capabilitiesResponse) setDefault() {

}
