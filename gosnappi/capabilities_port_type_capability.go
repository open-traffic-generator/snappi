package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** CapabilitiesPortTypeCapability *****
type capabilitiesPortTypeCapability struct {
	validation
	obj            *otg.CapabilitiesPortTypeCapability
	marshaller     marshalCapabilitiesPortTypeCapability
	unMarshaller   unMarshalCapabilitiesPortTypeCapability
	portTypeHolder CapabilitiesPortType
	pathsHolder    CapabilitiesPortTypeCapabilityCapabilitiesPathIter
}

func NewCapabilitiesPortTypeCapability() CapabilitiesPortTypeCapability {
	obj := capabilitiesPortTypeCapability{obj: &otg.CapabilitiesPortTypeCapability{}}
	obj.setDefault()
	return &obj
}

func (obj *capabilitiesPortTypeCapability) msg() *otg.CapabilitiesPortTypeCapability {
	return obj.obj
}

func (obj *capabilitiesPortTypeCapability) setMsg(msg *otg.CapabilitiesPortTypeCapability) CapabilitiesPortTypeCapability {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalcapabilitiesPortTypeCapability struct {
	obj *capabilitiesPortTypeCapability
}

type marshalCapabilitiesPortTypeCapability interface {
	// ToProto marshals CapabilitiesPortTypeCapability to protobuf object *otg.CapabilitiesPortTypeCapability
	ToProto() (*otg.CapabilitiesPortTypeCapability, error)
	// ToPbText marshals CapabilitiesPortTypeCapability to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals CapabilitiesPortTypeCapability to YAML text
	ToYaml() (string, error)
	// ToJson marshals CapabilitiesPortTypeCapability to JSON text
	ToJson() (string, error)
}

type unMarshalcapabilitiesPortTypeCapability struct {
	obj *capabilitiesPortTypeCapability
}

type unMarshalCapabilitiesPortTypeCapability interface {
	// FromProto unmarshals CapabilitiesPortTypeCapability from protobuf object *otg.CapabilitiesPortTypeCapability
	FromProto(msg *otg.CapabilitiesPortTypeCapability) (CapabilitiesPortTypeCapability, error)
	// FromPbText unmarshals CapabilitiesPortTypeCapability from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals CapabilitiesPortTypeCapability from YAML text
	FromYaml(value string) error
	// FromJson unmarshals CapabilitiesPortTypeCapability from JSON text
	FromJson(value string) error
}

func (obj *capabilitiesPortTypeCapability) Marshal() marshalCapabilitiesPortTypeCapability {
	if obj.marshaller == nil {
		obj.marshaller = &marshalcapabilitiesPortTypeCapability{obj: obj}
	}
	return obj.marshaller
}

func (obj *capabilitiesPortTypeCapability) Unmarshal() unMarshalCapabilitiesPortTypeCapability {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalcapabilitiesPortTypeCapability{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalcapabilitiesPortTypeCapability) ToProto() (*otg.CapabilitiesPortTypeCapability, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalcapabilitiesPortTypeCapability) FromProto(msg *otg.CapabilitiesPortTypeCapability) (CapabilitiesPortTypeCapability, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalcapabilitiesPortTypeCapability) ToPbText() (string, error) {
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

func (m *unMarshalcapabilitiesPortTypeCapability) FromPbText(value string) error {
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

func (m *marshalcapabilitiesPortTypeCapability) ToYaml() (string, error) {
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

func (m *unMarshalcapabilitiesPortTypeCapability) FromYaml(value string) error {
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

func (m *marshalcapabilitiesPortTypeCapability) ToJson() (string, error) {
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

func (m *unMarshalcapabilitiesPortTypeCapability) FromJson(value string) error {
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

func (obj *capabilitiesPortTypeCapability) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *capabilitiesPortTypeCapability) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *capabilitiesPortTypeCapability) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *capabilitiesPortTypeCapability) Clone() (CapabilitiesPortTypeCapability, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewCapabilitiesPortTypeCapability()
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

func (obj *capabilitiesPortTypeCapability) setNil() {
	obj.portTypeHolder = nil
	obj.pathsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// CapabilitiesPortTypeCapability is capabilities supported by one specific port-type supported by the implementation.
// If a specific port_type is requested, only that port_type should be included in the response.
// If no specific port_type is requested, information on all supported port_types in context of current deployment should be returned.
// The same port_type should not be returned more than once in the Capabilities response.
// The array of paths contain the actual list of supported compressed paths of the object model that are supported or not supported by the port_type for which
// the information is being returned , depending on whether the request_type is included_paths or excluded_paths .
type CapabilitiesPortTypeCapability interface {
	Validation
	// msg marshals CapabilitiesPortTypeCapability to protobuf object *otg.CapabilitiesPortTypeCapability
	// and doesn't set defaults
	msg() *otg.CapabilitiesPortTypeCapability
	// setMsg unmarshals CapabilitiesPortTypeCapability from protobuf object *otg.CapabilitiesPortTypeCapability
	// and doesn't set defaults
	setMsg(*otg.CapabilitiesPortTypeCapability) CapabilitiesPortTypeCapability
	// provides marshal interface
	Marshal() marshalCapabilitiesPortTypeCapability
	// provides unmarshal interface
	Unmarshal() unMarshalCapabilitiesPortTypeCapability
	// validate validates CapabilitiesPortTypeCapability
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (CapabilitiesPortTypeCapability, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PortType returns CapabilitiesPortType, set in CapabilitiesPortTypeCapability.
	// CapabilitiesPortType is the information regarding the unique PortType for which the included or excluded capabilities are being returned.
	PortType() CapabilitiesPortType
	// SetPortType assigns CapabilitiesPortType provided by user to CapabilitiesPortTypeCapability.
	// CapabilitiesPortType is the information regarding the unique PortType for which the included or excluded capabilities are being returned.
	SetPortType(value CapabilitiesPortType) CapabilitiesPortTypeCapability
	// HasPortType checks if PortType has been set in CapabilitiesPortTypeCapability
	HasPortType() bool
	// RequestType returns CapabilitiesPortTypeCapabilityRequestTypeEnum, set in CapabilitiesPortTypeCapability
	RequestType() CapabilitiesPortTypeCapabilityRequestTypeEnum
	// SetRequestType assigns CapabilitiesPortTypeCapabilityRequestTypeEnum provided by user to CapabilitiesPortTypeCapability
	SetRequestType(value CapabilitiesPortTypeCapabilityRequestTypeEnum) CapabilitiesPortTypeCapability
	// HasRequestType checks if RequestType has been set in CapabilitiesPortTypeCapability
	HasRequestType() bool
	// StartingRoot returns string, set in CapabilitiesPortTypeCapability.
	StartingRoot() string
	// SetStartingRoot assigns string provided by user to CapabilitiesPortTypeCapability
	SetStartingRoot(value string) CapabilitiesPortTypeCapability
	// HasStartingRoot checks if StartingRoot has been set in CapabilitiesPortTypeCapability
	HasStartingRoot() bool
	// Paths returns CapabilitiesPortTypeCapabilityCapabilitiesPathIterIter, set in CapabilitiesPortTypeCapability
	Paths() CapabilitiesPortTypeCapabilityCapabilitiesPathIter
	setNil()
}

// description is TBD
// PortType returns a CapabilitiesPortType
func (obj *capabilitiesPortTypeCapability) PortType() CapabilitiesPortType {
	if obj.obj.PortType == nil {
		obj.obj.PortType = NewCapabilitiesPortType().msg()
	}
	if obj.portTypeHolder == nil {
		obj.portTypeHolder = &capabilitiesPortType{obj: obj.obj.PortType}
	}
	return obj.portTypeHolder
}

// description is TBD
// PortType returns a CapabilitiesPortType
func (obj *capabilitiesPortTypeCapability) HasPortType() bool {
	return obj.obj.PortType != nil
}

// description is TBD
// SetPortType sets the CapabilitiesPortType value in the CapabilitiesPortTypeCapability object
func (obj *capabilitiesPortTypeCapability) SetPortType(value CapabilitiesPortType) CapabilitiesPortTypeCapability {

	obj.portTypeHolder = nil
	obj.obj.PortType = value.msg()

	return obj
}

type CapabilitiesPortTypeCapabilityRequestTypeEnum string

// Enum of RequestType on CapabilitiesPortTypeCapability
var CapabilitiesPortTypeCapabilityRequestType = struct {
	INCLUDED_PATHS CapabilitiesPortTypeCapabilityRequestTypeEnum
	EXCLUDED_PATHS CapabilitiesPortTypeCapabilityRequestTypeEnum
}{
	INCLUDED_PATHS: CapabilitiesPortTypeCapabilityRequestTypeEnum("included_paths"),
	EXCLUDED_PATHS: CapabilitiesPortTypeCapabilityRequestTypeEnum("excluded_paths"),
}

func (obj *capabilitiesPortTypeCapability) RequestType() CapabilitiesPortTypeCapabilityRequestTypeEnum {
	return CapabilitiesPortTypeCapabilityRequestTypeEnum(obj.obj.RequestType.Enum().String())
}

// If request_type is set in the request, that should be returned in this field . Otherwise , the default value of request_type (included_paths) , should be set.
// RequestType returns a string
func (obj *capabilitiesPortTypeCapability) HasRequestType() bool {
	return obj.obj.RequestType != nil
}

func (obj *capabilitiesPortTypeCapability) SetRequestType(value CapabilitiesPortTypeCapabilityRequestTypeEnum) CapabilitiesPortTypeCapability {
	intValue, ok := otg.CapabilitiesPortTypeCapability_RequestType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on CapabilitiesPortTypeCapabilityRequestTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.CapabilitiesPortTypeCapability_RequestType_Enum(intValue)
	obj.obj.RequestType = &enumValue

	return obj
}

// The starting_root in the request should be copied into this field in the response. If no starting_root is include , this field should be set to '/' , indicating the root  of the entire model hierarchy.
// StartingRoot returns a string
func (obj *capabilitiesPortTypeCapability) StartingRoot() string {

	return *obj.obj.StartingRoot

}

// The starting_root in the request should be copied into this field in the response. If no starting_root is include , this field should be set to '/' , indicating the root  of the entire model hierarchy.
// StartingRoot returns a string
func (obj *capabilitiesPortTypeCapability) HasStartingRoot() bool {
	return obj.obj.StartingRoot != nil
}

// The starting_root in the request should be copied into this field in the response. If no starting_root is include , this field should be set to '/' , indicating the root  of the entire model hierarchy.
// SetStartingRoot sets the string value in the CapabilitiesPortTypeCapability object
func (obj *capabilitiesPortTypeCapability) SetStartingRoot(value string) CapabilitiesPortTypeCapability {

	obj.obj.StartingRoot = &value
	return obj
}

// The complete list of object and attribute paths that provide the user with the full set of included or excluded capabilties for the specific port_type for which the capabilities are being returned.
// Paths returns a []CapabilitiesPath
func (obj *capabilitiesPortTypeCapability) Paths() CapabilitiesPortTypeCapabilityCapabilitiesPathIter {
	if len(obj.obj.Paths) == 0 {
		obj.obj.Paths = []*otg.CapabilitiesPath{}
	}
	if obj.pathsHolder == nil {
		obj.pathsHolder = newCapabilitiesPortTypeCapabilityCapabilitiesPathIter(&obj.obj.Paths).setMsg(obj)
	}
	return obj.pathsHolder
}

type capabilitiesPortTypeCapabilityCapabilitiesPathIter struct {
	obj                   *capabilitiesPortTypeCapability
	capabilitiesPathSlice []CapabilitiesPath
	fieldPtr              *[]*otg.CapabilitiesPath
}

func newCapabilitiesPortTypeCapabilityCapabilitiesPathIter(ptr *[]*otg.CapabilitiesPath) CapabilitiesPortTypeCapabilityCapabilitiesPathIter {
	return &capabilitiesPortTypeCapabilityCapabilitiesPathIter{fieldPtr: ptr}
}

type CapabilitiesPortTypeCapabilityCapabilitiesPathIter interface {
	setMsg(*capabilitiesPortTypeCapability) CapabilitiesPortTypeCapabilityCapabilitiesPathIter
	Items() []CapabilitiesPath
	Add() CapabilitiesPath
	Append(items ...CapabilitiesPath) CapabilitiesPortTypeCapabilityCapabilitiesPathIter
	Set(index int, newObj CapabilitiesPath) CapabilitiesPortTypeCapabilityCapabilitiesPathIter
	Clear() CapabilitiesPortTypeCapabilityCapabilitiesPathIter
	clearHolderSlice() CapabilitiesPortTypeCapabilityCapabilitiesPathIter
	appendHolderSlice(item CapabilitiesPath) CapabilitiesPortTypeCapabilityCapabilitiesPathIter
}

func (obj *capabilitiesPortTypeCapabilityCapabilitiesPathIter) setMsg(msg *capabilitiesPortTypeCapability) CapabilitiesPortTypeCapabilityCapabilitiesPathIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&capabilitiesPath{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *capabilitiesPortTypeCapabilityCapabilitiesPathIter) Items() []CapabilitiesPath {
	return obj.capabilitiesPathSlice
}

func (obj *capabilitiesPortTypeCapabilityCapabilitiesPathIter) Add() CapabilitiesPath {
	newObj := &otg.CapabilitiesPath{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &capabilitiesPath{obj: newObj}
	newLibObj.setDefault()
	obj.capabilitiesPathSlice = append(obj.capabilitiesPathSlice, newLibObj)
	return newLibObj
}

func (obj *capabilitiesPortTypeCapabilityCapabilitiesPathIter) Append(items ...CapabilitiesPath) CapabilitiesPortTypeCapabilityCapabilitiesPathIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.capabilitiesPathSlice = append(obj.capabilitiesPathSlice, item)
	}
	return obj
}

func (obj *capabilitiesPortTypeCapabilityCapabilitiesPathIter) Set(index int, newObj CapabilitiesPath) CapabilitiesPortTypeCapabilityCapabilitiesPathIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.capabilitiesPathSlice[index] = newObj
	return obj
}
func (obj *capabilitiesPortTypeCapabilityCapabilitiesPathIter) Clear() CapabilitiesPortTypeCapabilityCapabilitiesPathIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.CapabilitiesPath{}
		obj.capabilitiesPathSlice = []CapabilitiesPath{}
	}
	return obj
}
func (obj *capabilitiesPortTypeCapabilityCapabilitiesPathIter) clearHolderSlice() CapabilitiesPortTypeCapabilityCapabilitiesPathIter {
	if len(obj.capabilitiesPathSlice) > 0 {
		obj.capabilitiesPathSlice = []CapabilitiesPath{}
	}
	return obj
}
func (obj *capabilitiesPortTypeCapabilityCapabilitiesPathIter) appendHolderSlice(item CapabilitiesPath) CapabilitiesPortTypeCapabilityCapabilitiesPathIter {
	obj.capabilitiesPathSlice = append(obj.capabilitiesPathSlice, item)
	return obj
}

func (obj *capabilitiesPortTypeCapability) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PortType != nil {

		obj.PortType().validateObj(vObj, set_default)
	}

	if len(obj.obj.Paths) != 0 {

		if set_default {
			obj.Paths().clearHolderSlice()
			for _, item := range obj.obj.Paths {
				obj.Paths().appendHolderSlice(&capabilitiesPath{obj: item})
			}
		}
		for _, item := range obj.Paths().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *capabilitiesPortTypeCapability) setDefault() {

}
