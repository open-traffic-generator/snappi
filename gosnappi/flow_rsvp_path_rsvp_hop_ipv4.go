package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathRsvpHopIpv4 *****
type flowRSVPPathRsvpHopIpv4 struct {
	validation
	obj                          *otg.FlowRSVPPathRsvpHopIpv4
	marshaller                   marshalFlowRSVPPathRsvpHopIpv4
	unMarshaller                 unMarshalFlowRSVPPathRsvpHopIpv4
	ipv4AddressHolder            PatternFlowRSVPPathRsvpHopIpv4Ipv4Address
	logicalInterfaceHandleHolder PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
}

func NewFlowRSVPPathRsvpHopIpv4() FlowRSVPPathRsvpHopIpv4 {
	obj := flowRSVPPathRsvpHopIpv4{obj: &otg.FlowRSVPPathRsvpHopIpv4{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathRsvpHopIpv4) msg() *otg.FlowRSVPPathRsvpHopIpv4 {
	return obj.obj
}

func (obj *flowRSVPPathRsvpHopIpv4) setMsg(msg *otg.FlowRSVPPathRsvpHopIpv4) FlowRSVPPathRsvpHopIpv4 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathRsvpHopIpv4 struct {
	obj *flowRSVPPathRsvpHopIpv4
}

type marshalFlowRSVPPathRsvpHopIpv4 interface {
	// ToProto marshals FlowRSVPPathRsvpHopIpv4 to protobuf object *otg.FlowRSVPPathRsvpHopIpv4
	ToProto() (*otg.FlowRSVPPathRsvpHopIpv4, error)
	// ToPbText marshals FlowRSVPPathRsvpHopIpv4 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathRsvpHopIpv4 to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathRsvpHopIpv4 to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowRSVPPathRsvpHopIpv4 to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowRSVPPathRsvpHopIpv4 struct {
	obj *flowRSVPPathRsvpHopIpv4
}

type unMarshalFlowRSVPPathRsvpHopIpv4 interface {
	// FromProto unmarshals FlowRSVPPathRsvpHopIpv4 from protobuf object *otg.FlowRSVPPathRsvpHopIpv4
	FromProto(msg *otg.FlowRSVPPathRsvpHopIpv4) (FlowRSVPPathRsvpHopIpv4, error)
	// FromPbText unmarshals FlowRSVPPathRsvpHopIpv4 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathRsvpHopIpv4 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathRsvpHopIpv4 from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathRsvpHopIpv4) Marshal() marshalFlowRSVPPathRsvpHopIpv4 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathRsvpHopIpv4{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathRsvpHopIpv4) Unmarshal() unMarshalFlowRSVPPathRsvpHopIpv4 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathRsvpHopIpv4{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathRsvpHopIpv4) ToProto() (*otg.FlowRSVPPathRsvpHopIpv4, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathRsvpHopIpv4) FromProto(msg *otg.FlowRSVPPathRsvpHopIpv4) (FlowRSVPPathRsvpHopIpv4, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathRsvpHopIpv4) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathRsvpHopIpv4) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathRsvpHopIpv4) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathRsvpHopIpv4) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathRsvpHopIpv4) ToJsonRaw() (string, error) {
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

func (m *marshalflowRSVPPathRsvpHopIpv4) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathRsvpHopIpv4) FromJson(value string) error {
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

func (obj *flowRSVPPathRsvpHopIpv4) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathRsvpHopIpv4) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathRsvpHopIpv4) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathRsvpHopIpv4) Clone() (FlowRSVPPathRsvpHopIpv4, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathRsvpHopIpv4()
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

func (obj *flowRSVPPathRsvpHopIpv4) setNil() {
	obj.ipv4AddressHolder = nil
	obj.logicalInterfaceHandleHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathRsvpHopIpv4 is iPv4 RSVP_HOP object: Class = 3, C-Type = 1
type FlowRSVPPathRsvpHopIpv4 interface {
	Validation
	// msg marshals FlowRSVPPathRsvpHopIpv4 to protobuf object *otg.FlowRSVPPathRsvpHopIpv4
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathRsvpHopIpv4
	// setMsg unmarshals FlowRSVPPathRsvpHopIpv4 from protobuf object *otg.FlowRSVPPathRsvpHopIpv4
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathRsvpHopIpv4) FlowRSVPPathRsvpHopIpv4
	// provides marshal interface
	Marshal() marshalFlowRSVPPathRsvpHopIpv4
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathRsvpHopIpv4
	// validate validates FlowRSVPPathRsvpHopIpv4
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathRsvpHopIpv4, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4Address returns PatternFlowRSVPPathRsvpHopIpv4Ipv4Address, set in FlowRSVPPathRsvpHopIpv4.
	// PatternFlowRSVPPathRsvpHopIpv4Ipv4Address is the IPv4 address of the interface through which the last RSVP-knowledgeable hop forwarded this message.
	Ipv4Address() PatternFlowRSVPPathRsvpHopIpv4Ipv4Address
	// SetIpv4Address assigns PatternFlowRSVPPathRsvpHopIpv4Ipv4Address provided by user to FlowRSVPPathRsvpHopIpv4.
	// PatternFlowRSVPPathRsvpHopIpv4Ipv4Address is the IPv4 address of the interface through which the last RSVP-knowledgeable hop forwarded this message.
	SetIpv4Address(value PatternFlowRSVPPathRsvpHopIpv4Ipv4Address) FlowRSVPPathRsvpHopIpv4
	// HasIpv4Address checks if Ipv4Address has been set in FlowRSVPPathRsvpHopIpv4
	HasIpv4Address() bool
	// LogicalInterfaceHandle returns PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle, set in FlowRSVPPathRsvpHopIpv4.
	// PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle is logical Interface Handle (LIH) is used to distinguish logical outgoing interfaces. A node receiving an LIH in a Path message saves its value and returns it in the HOP objects of subsequent Resv messages sent to the node that originated the LIH. The LIH should be identically zero if there is no logical interface handle.
	LogicalInterfaceHandle() PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
	// SetLogicalInterfaceHandle assigns PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle provided by user to FlowRSVPPathRsvpHopIpv4.
	// PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle is logical Interface Handle (LIH) is used to distinguish logical outgoing interfaces. A node receiving an LIH in a Path message saves its value and returns it in the HOP objects of subsequent Resv messages sent to the node that originated the LIH. The LIH should be identically zero if there is no logical interface handle.
	SetLogicalInterfaceHandle(value PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) FlowRSVPPathRsvpHopIpv4
	// HasLogicalInterfaceHandle checks if LogicalInterfaceHandle has been set in FlowRSVPPathRsvpHopIpv4
	HasLogicalInterfaceHandle() bool
	setNil()
}

// description is TBD
// Ipv4Address returns a PatternFlowRSVPPathRsvpHopIpv4Ipv4Address
func (obj *flowRSVPPathRsvpHopIpv4) Ipv4Address() PatternFlowRSVPPathRsvpHopIpv4Ipv4Address {
	if obj.obj.Ipv4Address == nil {
		obj.obj.Ipv4Address = NewPatternFlowRSVPPathRsvpHopIpv4Ipv4Address().msg()
	}
	if obj.ipv4AddressHolder == nil {
		obj.ipv4AddressHolder = &patternFlowRSVPPathRsvpHopIpv4Ipv4Address{obj: obj.obj.Ipv4Address}
	}
	return obj.ipv4AddressHolder
}

// description is TBD
// Ipv4Address returns a PatternFlowRSVPPathRsvpHopIpv4Ipv4Address
func (obj *flowRSVPPathRsvpHopIpv4) HasIpv4Address() bool {
	return obj.obj.Ipv4Address != nil
}

// description is TBD
// SetIpv4Address sets the PatternFlowRSVPPathRsvpHopIpv4Ipv4Address value in the FlowRSVPPathRsvpHopIpv4 object
func (obj *flowRSVPPathRsvpHopIpv4) SetIpv4Address(value PatternFlowRSVPPathRsvpHopIpv4Ipv4Address) FlowRSVPPathRsvpHopIpv4 {

	obj.ipv4AddressHolder = nil
	obj.obj.Ipv4Address = value.msg()

	return obj
}

// description is TBD
// LogicalInterfaceHandle returns a PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
func (obj *flowRSVPPathRsvpHopIpv4) LogicalInterfaceHandle() PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle {
	if obj.obj.LogicalInterfaceHandle == nil {
		obj.obj.LogicalInterfaceHandle = NewPatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle().msg()
	}
	if obj.logicalInterfaceHandleHolder == nil {
		obj.logicalInterfaceHandleHolder = &patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle{obj: obj.obj.LogicalInterfaceHandle}
	}
	return obj.logicalInterfaceHandleHolder
}

// description is TBD
// LogicalInterfaceHandle returns a PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
func (obj *flowRSVPPathRsvpHopIpv4) HasLogicalInterfaceHandle() bool {
	return obj.obj.LogicalInterfaceHandle != nil
}

// description is TBD
// SetLogicalInterfaceHandle sets the PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle value in the FlowRSVPPathRsvpHopIpv4 object
func (obj *flowRSVPPathRsvpHopIpv4) SetLogicalInterfaceHandle(value PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) FlowRSVPPathRsvpHopIpv4 {

	obj.logicalInterfaceHandleHolder = nil
	obj.obj.LogicalInterfaceHandle = value.msg()

	return obj
}

func (obj *flowRSVPPathRsvpHopIpv4) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv4Address != nil {

		obj.Ipv4Address().validateObj(vObj, set_default)
	}

	if obj.obj.LogicalInterfaceHandle != nil {

		obj.LogicalInterfaceHandle().validateObj(vObj, set_default)
	}

}

func (obj *flowRSVPPathRsvpHopIpv4) setDefault() {

}
