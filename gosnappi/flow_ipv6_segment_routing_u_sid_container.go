package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6SegmentRoutingUSidContainer *****
type flowIpv6SegmentRoutingUSidContainer struct {
	validation
	obj          *otg.FlowIpv6SegmentRoutingUSidContainer
	marshaller   marshalFlowIpv6SegmentRoutingUSidContainer
	unMarshaller unMarshalFlowIpv6SegmentRoutingUSidContainer
	usidsHolder  FlowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter
}

func NewFlowIpv6SegmentRoutingUSidContainer() FlowIpv6SegmentRoutingUSidContainer {
	obj := flowIpv6SegmentRoutingUSidContainer{obj: &otg.FlowIpv6SegmentRoutingUSidContainer{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6SegmentRoutingUSidContainer) msg() *otg.FlowIpv6SegmentRoutingUSidContainer {
	return obj.obj
}

func (obj *flowIpv6SegmentRoutingUSidContainer) setMsg(msg *otg.FlowIpv6SegmentRoutingUSidContainer) FlowIpv6SegmentRoutingUSidContainer {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6SegmentRoutingUSidContainer struct {
	obj *flowIpv6SegmentRoutingUSidContainer
}

type marshalFlowIpv6SegmentRoutingUSidContainer interface {
	// ToProto marshals FlowIpv6SegmentRoutingUSidContainer to protobuf object *otg.FlowIpv6SegmentRoutingUSidContainer
	ToProto() (*otg.FlowIpv6SegmentRoutingUSidContainer, error)
	// ToPbText marshals FlowIpv6SegmentRoutingUSidContainer to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6SegmentRoutingUSidContainer to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6SegmentRoutingUSidContainer to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6SegmentRoutingUSidContainer struct {
	obj *flowIpv6SegmentRoutingUSidContainer
}

type unMarshalFlowIpv6SegmentRoutingUSidContainer interface {
	// FromProto unmarshals FlowIpv6SegmentRoutingUSidContainer from protobuf object *otg.FlowIpv6SegmentRoutingUSidContainer
	FromProto(msg *otg.FlowIpv6SegmentRoutingUSidContainer) (FlowIpv6SegmentRoutingUSidContainer, error)
	// FromPbText unmarshals FlowIpv6SegmentRoutingUSidContainer from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6SegmentRoutingUSidContainer from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6SegmentRoutingUSidContainer from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6SegmentRoutingUSidContainer) Marshal() marshalFlowIpv6SegmentRoutingUSidContainer {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6SegmentRoutingUSidContainer{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6SegmentRoutingUSidContainer) Unmarshal() unMarshalFlowIpv6SegmentRoutingUSidContainer {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6SegmentRoutingUSidContainer{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6SegmentRoutingUSidContainer) ToProto() (*otg.FlowIpv6SegmentRoutingUSidContainer, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6SegmentRoutingUSidContainer) FromProto(msg *otg.FlowIpv6SegmentRoutingUSidContainer) (FlowIpv6SegmentRoutingUSidContainer, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6SegmentRoutingUSidContainer) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingUSidContainer) FromPbText(value string) error {
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

func (m *marshalflowIpv6SegmentRoutingUSidContainer) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingUSidContainer) FromYaml(value string) error {
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

func (m *marshalflowIpv6SegmentRoutingUSidContainer) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingUSidContainer) FromJson(value string) error {
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

func (obj *flowIpv6SegmentRoutingUSidContainer) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingUSidContainer) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingUSidContainer) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6SegmentRoutingUSidContainer) Clone() (FlowIpv6SegmentRoutingUSidContainer, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6SegmentRoutingUSidContainer()
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

func (obj *flowIpv6SegmentRoutingUSidContainer) setNil() {
	obj.usidsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv6SegmentRoutingUSidContainer is a uSID container segment (RFC 9800) - a 128-bit IPv6 address packed with multiple Micro-SIDs (uSIDs). The implementation assembles the 128-bit value by placing the locator block prefix in the high bits, then packing each uSID (locator_node + function) consecutively in units of (locator_node_length + function_length) bits, padding the remainder with zeros. Example using F3216 (lb=32, ln=16, fn=16, arg=0) and locator block fc00::/32:
// usids = [{node: "0001", function: "0001"}, {node: "0002", function: "0001"}]
// => container = fc00:0:1:1:2:1:: (2-hop path, nodes 1 and 2, function 1)
// Reference: RFC 9800 Section 4.
type FlowIpv6SegmentRoutingUSidContainer interface {
	Validation
	// msg marshals FlowIpv6SegmentRoutingUSidContainer to protobuf object *otg.FlowIpv6SegmentRoutingUSidContainer
	// and doesn't set defaults
	msg() *otg.FlowIpv6SegmentRoutingUSidContainer
	// setMsg unmarshals FlowIpv6SegmentRoutingUSidContainer from protobuf object *otg.FlowIpv6SegmentRoutingUSidContainer
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6SegmentRoutingUSidContainer) FlowIpv6SegmentRoutingUSidContainer
	// provides marshal interface
	Marshal() marshalFlowIpv6SegmentRoutingUSidContainer
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6SegmentRoutingUSidContainer
	// validate validates FlowIpv6SegmentRoutingUSidContainer
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6SegmentRoutingUSidContainer, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LocatorBlock returns string, set in FlowIpv6SegmentRoutingUSidContainer.
	LocatorBlock() string
	// SetLocatorBlock assigns string provided by user to FlowIpv6SegmentRoutingUSidContainer
	SetLocatorBlock(value string) FlowIpv6SegmentRoutingUSidContainer
	// LocatorBlockLength returns uint32, set in FlowIpv6SegmentRoutingUSidContainer.
	LocatorBlockLength() uint32
	// SetLocatorBlockLength assigns uint32 provided by user to FlowIpv6SegmentRoutingUSidContainer
	SetLocatorBlockLength(value uint32) FlowIpv6SegmentRoutingUSidContainer
	// HasLocatorBlockLength checks if LocatorBlockLength has been set in FlowIpv6SegmentRoutingUSidContainer
	HasLocatorBlockLength() bool
	// LocatorNodeLength returns uint32, set in FlowIpv6SegmentRoutingUSidContainer.
	LocatorNodeLength() uint32
	// SetLocatorNodeLength assigns uint32 provided by user to FlowIpv6SegmentRoutingUSidContainer
	SetLocatorNodeLength(value uint32) FlowIpv6SegmentRoutingUSidContainer
	// HasLocatorNodeLength checks if LocatorNodeLength has been set in FlowIpv6SegmentRoutingUSidContainer
	HasLocatorNodeLength() bool
	// FunctionLength returns uint32, set in FlowIpv6SegmentRoutingUSidContainer.
	FunctionLength() uint32
	// SetFunctionLength assigns uint32 provided by user to FlowIpv6SegmentRoutingUSidContainer
	SetFunctionLength(value uint32) FlowIpv6SegmentRoutingUSidContainer
	// HasFunctionLength checks if FunctionLength has been set in FlowIpv6SegmentRoutingUSidContainer
	HasFunctionLength() bool
	// Usids returns FlowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIterIter, set in FlowIpv6SegmentRoutingUSidContainer
	Usids() FlowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter
	setNil()
}

// The common locator block IPv6 prefix shared across all uSIDs in this container (e.g., "fc00::" for the fcbb:bbbb::/32 SONiC convention). This prefix occupies the most-significant locator_block_length bits of the 128-bit container.
// LocatorBlock returns a string
func (obj *flowIpv6SegmentRoutingUSidContainer) LocatorBlock() string {

	return *obj.obj.LocatorBlock

}

// The common locator block IPv6 prefix shared across all uSIDs in this container (e.g., "fc00::" for the fcbb:bbbb::/32 SONiC convention). This prefix occupies the most-significant locator_block_length bits of the 128-bit container.
// SetLocatorBlock sets the string value in the FlowIpv6SegmentRoutingUSidContainer object
func (obj *flowIpv6SegmentRoutingUSidContainer) SetLocatorBlock(value string) FlowIpv6SegmentRoutingUSidContainer {

	obj.obj.LocatorBlock = &value
	return obj
}

// Length of the locator block in bits (default 32 for F3216 uSID format). Must match the lb_length in the IsisSRv6.SidStructure used by the emulated devices.
// LocatorBlockLength returns a uint32
func (obj *flowIpv6SegmentRoutingUSidContainer) LocatorBlockLength() uint32 {

	return *obj.obj.LocatorBlockLength

}

// Length of the locator block in bits (default 32 for F3216 uSID format). Must match the lb_length in the IsisSRv6.SidStructure used by the emulated devices.
// LocatorBlockLength returns a uint32
func (obj *flowIpv6SegmentRoutingUSidContainer) HasLocatorBlockLength() bool {
	return obj.obj.LocatorBlockLength != nil
}

// Length of the locator block in bits (default 32 for F3216 uSID format). Must match the lb_length in the IsisSRv6.SidStructure used by the emulated devices.
// SetLocatorBlockLength sets the uint32 value in the FlowIpv6SegmentRoutingUSidContainer object
func (obj *flowIpv6SegmentRoutingUSidContainer) SetLocatorBlockLength(value uint32) FlowIpv6SegmentRoutingUSidContainer {

	obj.obj.LocatorBlockLength = &value
	return obj
}

// Length of the locator node field within each uSID slot, in bits (default 16 for F3216). Each uSID slot is (locator_node_length + function_length) bits wide.
// LocatorNodeLength returns a uint32
func (obj *flowIpv6SegmentRoutingUSidContainer) LocatorNodeLength() uint32 {

	return *obj.obj.LocatorNodeLength

}

// Length of the locator node field within each uSID slot, in bits (default 16 for F3216). Each uSID slot is (locator_node_length + function_length) bits wide.
// LocatorNodeLength returns a uint32
func (obj *flowIpv6SegmentRoutingUSidContainer) HasLocatorNodeLength() bool {
	return obj.obj.LocatorNodeLength != nil
}

// Length of the locator node field within each uSID slot, in bits (default 16 for F3216). Each uSID slot is (locator_node_length + function_length) bits wide.
// SetLocatorNodeLength sets the uint32 value in the FlowIpv6SegmentRoutingUSidContainer object
func (obj *flowIpv6SegmentRoutingUSidContainer) SetLocatorNodeLength(value uint32) FlowIpv6SegmentRoutingUSidContainer {

	obj.obj.LocatorNodeLength = &value
	return obj
}

// Length of the function field within each uSID slot, in bits (default 16 for F3216).
// FunctionLength returns a uint32
func (obj *flowIpv6SegmentRoutingUSidContainer) FunctionLength() uint32 {

	return *obj.obj.FunctionLength

}

// Length of the function field within each uSID slot, in bits (default 16 for F3216).
// FunctionLength returns a uint32
func (obj *flowIpv6SegmentRoutingUSidContainer) HasFunctionLength() bool {
	return obj.obj.FunctionLength != nil
}

// Length of the function field within each uSID slot, in bits (default 16 for F3216).
// SetFunctionLength sets the uint32 value in the FlowIpv6SegmentRoutingUSidContainer object
func (obj *flowIpv6SegmentRoutingUSidContainer) SetFunctionLength(value uint32) FlowIpv6SegmentRoutingUSidContainer {

	obj.obj.FunctionLength = &value
	return obj
}

// Ordered list of uSIDs to pack into this container, from left (most significant) to right. Each uSID entry provides the locator_node and function hex values; the implementation packs them consecutively after the locator block. The number of uSIDs that fit in one container is:
// floor((128 - locator_block_length) /
// (locator_node_length + function_length))
// For F3216: floor((128-32)/(16+16)) = 3 uSIDs per container.
// Usids returns a []FlowIpv6SegmentRoutingUSid
func (obj *flowIpv6SegmentRoutingUSidContainer) Usids() FlowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter {
	if len(obj.obj.Usids) == 0 {
		obj.obj.Usids = []*otg.FlowIpv6SegmentRoutingUSid{}
	}
	if obj.usidsHolder == nil {
		obj.usidsHolder = newFlowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter(&obj.obj.Usids).setMsg(obj)
	}
	return obj.usidsHolder
}

type flowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter struct {
	obj                             *flowIpv6SegmentRoutingUSidContainer
	flowIpv6SegmentRoutingUSidSlice []FlowIpv6SegmentRoutingUSid
	fieldPtr                        *[]*otg.FlowIpv6SegmentRoutingUSid
}

func newFlowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter(ptr *[]*otg.FlowIpv6SegmentRoutingUSid) FlowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter {
	return &flowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter{fieldPtr: ptr}
}

type FlowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter interface {
	setMsg(*flowIpv6SegmentRoutingUSidContainer) FlowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter
	Items() []FlowIpv6SegmentRoutingUSid
	Add() FlowIpv6SegmentRoutingUSid
	Append(items ...FlowIpv6SegmentRoutingUSid) FlowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter
	Set(index int, newObj FlowIpv6SegmentRoutingUSid) FlowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter
	Clear() FlowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter
	clearHolderSlice() FlowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter
	appendHolderSlice(item FlowIpv6SegmentRoutingUSid) FlowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter
}

func (obj *flowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter) setMsg(msg *flowIpv6SegmentRoutingUSidContainer) FlowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowIpv6SegmentRoutingUSid{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter) Items() []FlowIpv6SegmentRoutingUSid {
	return obj.flowIpv6SegmentRoutingUSidSlice
}

func (obj *flowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter) Add() FlowIpv6SegmentRoutingUSid {
	newObj := &otg.FlowIpv6SegmentRoutingUSid{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowIpv6SegmentRoutingUSid{obj: newObj}
	newLibObj.setDefault()
	obj.flowIpv6SegmentRoutingUSidSlice = append(obj.flowIpv6SegmentRoutingUSidSlice, newLibObj)
	return newLibObj
}

func (obj *flowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter) Append(items ...FlowIpv6SegmentRoutingUSid) FlowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowIpv6SegmentRoutingUSidSlice = append(obj.flowIpv6SegmentRoutingUSidSlice, item)
	}
	return obj
}

func (obj *flowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter) Set(index int, newObj FlowIpv6SegmentRoutingUSid) FlowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowIpv6SegmentRoutingUSidSlice[index] = newObj
	return obj
}
func (obj *flowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter) Clear() FlowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowIpv6SegmentRoutingUSid{}
		obj.flowIpv6SegmentRoutingUSidSlice = []FlowIpv6SegmentRoutingUSid{}
	}
	return obj
}
func (obj *flowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter) clearHolderSlice() FlowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter {
	if len(obj.flowIpv6SegmentRoutingUSidSlice) > 0 {
		obj.flowIpv6SegmentRoutingUSidSlice = []FlowIpv6SegmentRoutingUSid{}
	}
	return obj
}
func (obj *flowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter) appendHolderSlice(item FlowIpv6SegmentRoutingUSid) FlowIpv6SegmentRoutingUSidContainerFlowIpv6SegmentRoutingUSidIter {
	obj.flowIpv6SegmentRoutingUSidSlice = append(obj.flowIpv6SegmentRoutingUSidSlice, item)
	return obj
}

func (obj *flowIpv6SegmentRoutingUSidContainer) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// LocatorBlock is required
	if obj.obj.LocatorBlock == nil {
		vObj.validationErrors = append(vObj.validationErrors, "LocatorBlock is required field on interface FlowIpv6SegmentRoutingUSidContainer")
	}
	if obj.obj.LocatorBlock != nil {

		err := obj.validateIpv6(obj.LocatorBlock())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowIpv6SegmentRoutingUSidContainer.LocatorBlock"))
		}

	}

	if obj.obj.LocatorBlockLength != nil {

		if *obj.obj.LocatorBlockLength < 1 || *obj.obj.LocatorBlockLength > 120 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= FlowIpv6SegmentRoutingUSidContainer.LocatorBlockLength <= 120 but Got %d", *obj.obj.LocatorBlockLength))
		}

	}

	if obj.obj.LocatorNodeLength != nil {

		if *obj.obj.LocatorNodeLength > 64 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowIpv6SegmentRoutingUSidContainer.LocatorNodeLength <= 64 but Got %d", *obj.obj.LocatorNodeLength))
		}

	}

	if obj.obj.FunctionLength != nil {

		if *obj.obj.FunctionLength > 64 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowIpv6SegmentRoutingUSidContainer.FunctionLength <= 64 but Got %d", *obj.obj.FunctionLength))
		}

	}

	if len(obj.obj.Usids) != 0 {

		if set_default {
			obj.Usids().clearHolderSlice()
			for _, item := range obj.obj.Usids {
				obj.Usids().appendHolderSlice(&flowIpv6SegmentRoutingUSid{obj: item})
			}
		}
		for _, item := range obj.Usids().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *flowIpv6SegmentRoutingUSidContainer) setDefault() {
	if obj.obj.LocatorBlockLength == nil {
		obj.SetLocatorBlockLength(32)
	}
	if obj.obj.LocatorNodeLength == nil {
		obj.SetLocatorNodeLength(16)
	}
	if obj.obj.FunctionLength == nil {
		obj.SetFunctionLength(16)
	}

}
