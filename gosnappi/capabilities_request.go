package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** CapabilitiesRequest *****
type capabilitiesRequest struct {
	validation
	obj            *otg.CapabilitiesRequest
	marshaller     marshalCapabilitiesRequest
	unMarshaller   unMarshalCapabilitiesRequest
	portTypeHolder CapabilitiesPortType
}

func NewCapabilitiesRequest() CapabilitiesRequest {
	obj := capabilitiesRequest{obj: &otg.CapabilitiesRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *capabilitiesRequest) msg() *otg.CapabilitiesRequest {
	return obj.obj
}

func (obj *capabilitiesRequest) setMsg(msg *otg.CapabilitiesRequest) CapabilitiesRequest {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalcapabilitiesRequest struct {
	obj *capabilitiesRequest
}

type marshalCapabilitiesRequest interface {
	// ToProto marshals CapabilitiesRequest to protobuf object *otg.CapabilitiesRequest
	ToProto() (*otg.CapabilitiesRequest, error)
	// ToPbText marshals CapabilitiesRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals CapabilitiesRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals CapabilitiesRequest to JSON text
	ToJson() (string, error)
}

type unMarshalcapabilitiesRequest struct {
	obj *capabilitiesRequest
}

type unMarshalCapabilitiesRequest interface {
	// FromProto unmarshals CapabilitiesRequest from protobuf object *otg.CapabilitiesRequest
	FromProto(msg *otg.CapabilitiesRequest) (CapabilitiesRequest, error)
	// FromPbText unmarshals CapabilitiesRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals CapabilitiesRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals CapabilitiesRequest from JSON text
	FromJson(value string) error
}

func (obj *capabilitiesRequest) Marshal() marshalCapabilitiesRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalcapabilitiesRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *capabilitiesRequest) Unmarshal() unMarshalCapabilitiesRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalcapabilitiesRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalcapabilitiesRequest) ToProto() (*otg.CapabilitiesRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalcapabilitiesRequest) FromProto(msg *otg.CapabilitiesRequest) (CapabilitiesRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalcapabilitiesRequest) ToPbText() (string, error) {
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

func (m *unMarshalcapabilitiesRequest) FromPbText(value string) error {
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

func (m *marshalcapabilitiesRequest) ToYaml() (string, error) {
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

func (m *unMarshalcapabilitiesRequest) FromYaml(value string) error {
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

func (m *marshalcapabilitiesRequest) ToJson() (string, error) {
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

func (m *unMarshalcapabilitiesRequest) FromJson(value string) error {
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

func (obj *capabilitiesRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *capabilitiesRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *capabilitiesRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *capabilitiesRequest) Clone() (CapabilitiesRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewCapabilitiesRequest()
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

func (obj *capabilitiesRequest) setNil() {
	obj.portTypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// CapabilitiesRequest is request to traffic generator for supported capabilities.
// An implementation supporting capabilities API should support ability to return meanigful information
// even when no attributes are explicitly exposed in the capabilities request.
type CapabilitiesRequest interface {
	Validation
	// msg marshals CapabilitiesRequest to protobuf object *otg.CapabilitiesRequest
	// and doesn't set defaults
	msg() *otg.CapabilitiesRequest
	// setMsg unmarshals CapabilitiesRequest from protobuf object *otg.CapabilitiesRequest
	// and doesn't set defaults
	setMsg(*otg.CapabilitiesRequest) CapabilitiesRequest
	// provides marshal interface
	Marshal() marshalCapabilitiesRequest
	// provides unmarshal interface
	Unmarshal() unMarshalCapabilitiesRequest
	// validate validates CapabilitiesRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (CapabilitiesRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PortType returns CapabilitiesPortType, set in CapabilitiesRequest.
	// CapabilitiesPortType is the information regarding the unique PortType for which the included or excluded capabilities are being returned.
	PortType() CapabilitiesPortType
	// SetPortType assigns CapabilitiesPortType provided by user to CapabilitiesRequest.
	// CapabilitiesPortType is the information regarding the unique PortType for which the included or excluded capabilities are being returned.
	SetPortType(value CapabilitiesPortType) CapabilitiesRequest
	// HasPortType checks if PortType has been set in CapabilitiesRequest
	HasPortType() bool
	// RequestType returns CapabilitiesRequestRequestTypeEnum, set in CapabilitiesRequest
	RequestType() CapabilitiesRequestRequestTypeEnum
	// SetRequestType assigns CapabilitiesRequestRequestTypeEnum provided by user to CapabilitiesRequest
	SetRequestType(value CapabilitiesRequestRequestTypeEnum) CapabilitiesRequest
	// HasRequestType checks if RequestType has been set in CapabilitiesRequest
	HasRequestType() bool
	// StartingRoot returns string, set in CapabilitiesRequest.
	StartingRoot() string
	// SetStartingRoot assigns string provided by user to CapabilitiesRequest
	SetStartingRoot(value string) CapabilitiesRequest
	// HasStartingRoot checks if StartingRoot has been set in CapabilitiesRequest
	HasStartingRoot() bool
	setNil()
}

// description is TBD
// PortType returns a CapabilitiesPortType
func (obj *capabilitiesRequest) PortType() CapabilitiesPortType {
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
func (obj *capabilitiesRequest) HasPortType() bool {
	return obj.obj.PortType != nil
}

// description is TBD
// SetPortType sets the CapabilitiesPortType value in the CapabilitiesRequest object
func (obj *capabilitiesRequest) SetPortType(value CapabilitiesPortType) CapabilitiesRequest {

	obj.portTypeHolder = nil
	obj.obj.PortType = value.msg()

	return obj
}

type CapabilitiesRequestRequestTypeEnum string

// Enum of RequestType on CapabilitiesRequest
var CapabilitiesRequestRequestType = struct {
	INCLUDED_PATHS CapabilitiesRequestRequestTypeEnum
	EXCLUDED_PATHS CapabilitiesRequestRequestTypeEnum
}{
	INCLUDED_PATHS: CapabilitiesRequestRequestTypeEnum("included_paths"),
	EXCLUDED_PATHS: CapabilitiesRequestRequestTypeEnum("excluded_paths"),
}

func (obj *capabilitiesRequest) RequestType() CapabilitiesRequestRequestTypeEnum {
	return CapabilitiesRequestRequestTypeEnum(obj.obj.RequestType.Enum().String())
}

// If included_paths is set , this indicates that the user is interested in know about the objects and attributes
// supported by the implementation on its endpoints in an heirarchial manner starting from the root or user specified sub-root of the full object heirarchy.
// If excluded_paths is set , this indicates that the user is interested in know about the objects and attributes
// specifically NOT supported by the implementation on its endpoints in an heirarchial manner starting from the root or user specified sub-root of the full object heirarchy.
// Note that if the implementation itself is supporting an older version of models, it is possible that for excluded_paths request, it will not be able to return objects and attributes
// added in later versions of the model at minimum.
// RequestType returns a string
func (obj *capabilitiesRequest) HasRequestType() bool {
	return obj.obj.RequestType != nil
}

func (obj *capabilitiesRequest) SetRequestType(value CapabilitiesRequestRequestTypeEnum) CapabilitiesRequest {
	intValue, ok := otg.CapabilitiesRequest_RequestType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on CapabilitiesRequestRequestTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.CapabilitiesRequest_RequestType_Enum(intValue)
	obj.obj.RequestType = &enumValue

	return obj
}

// The starting point of the hierarchy from where user is requesting for included/excluded capabilities.By default, the included/excluded objects/actions and attributes of the entire hierarchy is
// provided. The starting_root should always be of the generic format `/<node1>/<node2>/..../<node-n>[/<attribute[/choice/<enum value of choice attribute >]]` where node-x is a valid child object of node-(x-1) and
// optional `attribute` is a child attribute of node-n i.e. it can also be used to request for providing included or excluded information about a specific attribute or value of a leaf choice attribute.
// e.g. a) To know related included or excluded attributes for BGP only , the starting_root  should be set as /set_config/devices/bgp ;
// b) To know whether IPv4 ping is supported or not , the starting_root should be set as /set_control_action/protocol/ipv4/ping
// StartingRoot returns a string
func (obj *capabilitiesRequest) StartingRoot() string {

	return *obj.obj.StartingRoot

}

// The starting point of the hierarchy from where user is requesting for included/excluded capabilities.By default, the included/excluded objects/actions and attributes of the entire hierarchy is
// provided. The starting_root should always be of the generic format `/<node1>/<node2>/..../<node-n>[/<attribute[/choice/<enum value of choice attribute >]]` where node-x is a valid child object of node-(x-1) and
// optional `attribute` is a child attribute of node-n i.e. it can also be used to request for providing included or excluded information about a specific attribute or value of a leaf choice attribute.
// e.g. a) To know related included or excluded attributes for BGP only , the starting_root  should be set as /set_config/devices/bgp ;
// b) To know whether IPv4 ping is supported or not , the starting_root should be set as /set_control_action/protocol/ipv4/ping
// StartingRoot returns a string
func (obj *capabilitiesRequest) HasStartingRoot() bool {
	return obj.obj.StartingRoot != nil
}

// The starting point of the hierarchy from where user is requesting for included/excluded capabilities.By default, the included/excluded objects/actions and attributes of the entire hierarchy is
// provided. The starting_root should always be of the generic format `/<node1>/<node2>/..../<node-n>[/<attribute[/choice/<enum value of choice attribute >]]` where node-x is a valid child object of node-(x-1) and
// optional `attribute` is a child attribute of node-n i.e. it can also be used to request for providing included or excluded information about a specific attribute or value of a leaf choice attribute.
// e.g. a) To know related included or excluded attributes for BGP only , the starting_root  should be set as /set_config/devices/bgp ;
// b) To know whether IPv4 ping is supported or not , the starting_root should be set as /set_control_action/protocol/ipv4/ping
// SetStartingRoot sets the string value in the CapabilitiesRequest object
func (obj *capabilitiesRequest) SetStartingRoot(value string) CapabilitiesRequest {

	obj.obj.StartingRoot = &value
	return obj
}

func (obj *capabilitiesRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PortType != nil {

		obj.PortType().validateObj(vObj, set_default)
	}

}

func (obj *capabilitiesRequest) setDefault() {
	if obj.obj.RequestType == nil {
		obj.SetRequestType(CapabilitiesRequestRequestType.INCLUDED_PATHS)

	}
	if obj.obj.StartingRoot == nil {
		obj.SetStartingRoot("/")
	}

}
