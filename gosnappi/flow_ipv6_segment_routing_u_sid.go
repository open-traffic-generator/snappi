package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6SegmentRoutingUSid *****
type flowIpv6SegmentRoutingUSid struct {
	validation
	obj          *otg.FlowIpv6SegmentRoutingUSid
	marshaller   marshalFlowIpv6SegmentRoutingUSid
	unMarshaller unMarshalFlowIpv6SegmentRoutingUSid
}

func NewFlowIpv6SegmentRoutingUSid() FlowIpv6SegmentRoutingUSid {
	obj := flowIpv6SegmentRoutingUSid{obj: &otg.FlowIpv6SegmentRoutingUSid{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6SegmentRoutingUSid) msg() *otg.FlowIpv6SegmentRoutingUSid {
	return obj.obj
}

func (obj *flowIpv6SegmentRoutingUSid) setMsg(msg *otg.FlowIpv6SegmentRoutingUSid) FlowIpv6SegmentRoutingUSid {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6SegmentRoutingUSid struct {
	obj *flowIpv6SegmentRoutingUSid
}

type marshalFlowIpv6SegmentRoutingUSid interface {
	// ToProto marshals FlowIpv6SegmentRoutingUSid to protobuf object *otg.FlowIpv6SegmentRoutingUSid
	ToProto() (*otg.FlowIpv6SegmentRoutingUSid, error)
	// ToPbText marshals FlowIpv6SegmentRoutingUSid to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6SegmentRoutingUSid to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6SegmentRoutingUSid to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6SegmentRoutingUSid struct {
	obj *flowIpv6SegmentRoutingUSid
}

type unMarshalFlowIpv6SegmentRoutingUSid interface {
	// FromProto unmarshals FlowIpv6SegmentRoutingUSid from protobuf object *otg.FlowIpv6SegmentRoutingUSid
	FromProto(msg *otg.FlowIpv6SegmentRoutingUSid) (FlowIpv6SegmentRoutingUSid, error)
	// FromPbText unmarshals FlowIpv6SegmentRoutingUSid from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6SegmentRoutingUSid from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6SegmentRoutingUSid from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6SegmentRoutingUSid) Marshal() marshalFlowIpv6SegmentRoutingUSid {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6SegmentRoutingUSid{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6SegmentRoutingUSid) Unmarshal() unMarshalFlowIpv6SegmentRoutingUSid {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6SegmentRoutingUSid{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6SegmentRoutingUSid) ToProto() (*otg.FlowIpv6SegmentRoutingUSid, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6SegmentRoutingUSid) FromProto(msg *otg.FlowIpv6SegmentRoutingUSid) (FlowIpv6SegmentRoutingUSid, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6SegmentRoutingUSid) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingUSid) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalflowIpv6SegmentRoutingUSid) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingUSid) FromYaml(value string) error {
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

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalflowIpv6SegmentRoutingUSid) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingUSid) FromJson(value string) error {
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

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *flowIpv6SegmentRoutingUSid) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingUSid) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingUSid) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6SegmentRoutingUSid) Clone() (FlowIpv6SegmentRoutingUSid, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6SegmentRoutingUSid()
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

// FlowIpv6SegmentRoutingUSid is one Micro-SID (uSID) entry within a uSID container (RFC 9800). Specifies the locator node and function hex values that are packed into consecutive bits of the 128-bit container value.
type FlowIpv6SegmentRoutingUSid interface {
	Validation
	// msg marshals FlowIpv6SegmentRoutingUSid to protobuf object *otg.FlowIpv6SegmentRoutingUSid
	// and doesn't set defaults
	msg() *otg.FlowIpv6SegmentRoutingUSid
	// setMsg unmarshals FlowIpv6SegmentRoutingUSid from protobuf object *otg.FlowIpv6SegmentRoutingUSid
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6SegmentRoutingUSid) FlowIpv6SegmentRoutingUSid
	// provides marshal interface
	Marshal() marshalFlowIpv6SegmentRoutingUSid
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6SegmentRoutingUSid
	// validate validates FlowIpv6SegmentRoutingUSid
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6SegmentRoutingUSid, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LocatorNode returns string, set in FlowIpv6SegmentRoutingUSid.
	LocatorNode() string
	// SetLocatorNode assigns string provided by user to FlowIpv6SegmentRoutingUSid
	SetLocatorNode(value string) FlowIpv6SegmentRoutingUSid
	// HasLocatorNode checks if LocatorNode has been set in FlowIpv6SegmentRoutingUSid
	HasLocatorNode() bool
	// Function returns string, set in FlowIpv6SegmentRoutingUSid.
	Function() string
	// SetFunction assigns string provided by user to FlowIpv6SegmentRoutingUSid
	SetFunction(value string) FlowIpv6SegmentRoutingUSid
	// HasFunction checks if Function has been set in FlowIpv6SegmentRoutingUSid
	HasFunction() bool
}

// Locator node identifier as a hex string, occupying locator_node_length bits in the container. For F3216 (locator_node_length=16), this is a 4-nibble string. Example: "0001" identifies node 1.
// LocatorNode returns a string
func (obj *flowIpv6SegmentRoutingUSid) LocatorNode() string {

	return *obj.obj.LocatorNode

}

// Locator node identifier as a hex string, occupying locator_node_length bits in the container. For F3216 (locator_node_length=16), this is a 4-nibble string. Example: "0001" identifies node 1.
// LocatorNode returns a string
func (obj *flowIpv6SegmentRoutingUSid) HasLocatorNode() bool {
	return obj.obj.LocatorNode != nil
}

// Locator node identifier as a hex string, occupying locator_node_length bits in the container. For F3216 (locator_node_length=16), this is a 4-nibble string. Example: "0001" identifies node 1.
// SetLocatorNode sets the string value in the FlowIpv6SegmentRoutingUSid object
func (obj *flowIpv6SegmentRoutingUSid) SetLocatorNode(value string) FlowIpv6SegmentRoutingUSid {

	obj.obj.LocatorNode = &value
	return obj
}

// Function identifier as a hex string, occupying function_length bits in the container. For F3216 (function_length=16), this is a 4-nibble string. Example: "0001" for End, "00c8" for End.X.
// Function returns a string
func (obj *flowIpv6SegmentRoutingUSid) Function() string {

	return *obj.obj.Function

}

// Function identifier as a hex string, occupying function_length bits in the container. For F3216 (function_length=16), this is a 4-nibble string. Example: "0001" for End, "00c8" for End.X.
// Function returns a string
func (obj *flowIpv6SegmentRoutingUSid) HasFunction() bool {
	return obj.obj.Function != nil
}

// Function identifier as a hex string, occupying function_length bits in the container. For F3216 (function_length=16), this is a 4-nibble string. Example: "0001" for End, "00c8" for End.X.
// SetFunction sets the string value in the FlowIpv6SegmentRoutingUSid object
func (obj *flowIpv6SegmentRoutingUSid) SetFunction(value string) FlowIpv6SegmentRoutingUSid {

	obj.obj.Function = &value
	return obj
}

func (obj *flowIpv6SegmentRoutingUSid) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.LocatorNode != nil {

		err := obj.validateHex(obj.LocatorNode())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowIpv6SegmentRoutingUSid.LocatorNode"))
		}

	}

	if obj.obj.Function != nil {

		err := obj.validateHex(obj.Function())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowIpv6SegmentRoutingUSid.Function"))
		}

	}

}

func (obj *flowIpv6SegmentRoutingUSid) setDefault() {

}
