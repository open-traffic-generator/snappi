package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIcmpv6NeighborSolicitation *****
type flowIcmpv6NeighborSolicitation struct {
	validation
	obj                 *otg.FlowIcmpv6NeighborSolicitation
	marshaller          marshalFlowIcmpv6NeighborSolicitation
	unMarshaller        unMarshalFlowIcmpv6NeighborSolicitation
	typeHolder          PatternFlowIcmpv6NeighborSolicitationType
	codeHolder          PatternFlowIcmpv6NeighborSolicitationCode
	checksumHolder      PatternFlowIcmpv6NeighborSolicitationChecksum
	reservedHolder      PatternFlowIcmpv6NeighborSolicitationReserved
	targetAddressHolder PatternFlowIcmpv6NeighborSolicitationTargetAddress
	optionsHolder       FlowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter
}

func NewFlowIcmpv6NeighborSolicitation() FlowIcmpv6NeighborSolicitation {
	obj := flowIcmpv6NeighborSolicitation{obj: &otg.FlowIcmpv6NeighborSolicitation{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIcmpv6NeighborSolicitation) msg() *otg.FlowIcmpv6NeighborSolicitation {
	return obj.obj
}

func (obj *flowIcmpv6NeighborSolicitation) setMsg(msg *otg.FlowIcmpv6NeighborSolicitation) FlowIcmpv6NeighborSolicitation {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIcmpv6NeighborSolicitation struct {
	obj *flowIcmpv6NeighborSolicitation
}

type marshalFlowIcmpv6NeighborSolicitation interface {
	// ToProto marshals FlowIcmpv6NeighborSolicitation to protobuf object *otg.FlowIcmpv6NeighborSolicitation
	ToProto() (*otg.FlowIcmpv6NeighborSolicitation, error)
	// ToPbText marshals FlowIcmpv6NeighborSolicitation to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIcmpv6NeighborSolicitation to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIcmpv6NeighborSolicitation to JSON text
	ToJson() (string, error)
}

type unMarshalflowIcmpv6NeighborSolicitation struct {
	obj *flowIcmpv6NeighborSolicitation
}

type unMarshalFlowIcmpv6NeighborSolicitation interface {
	// FromProto unmarshals FlowIcmpv6NeighborSolicitation from protobuf object *otg.FlowIcmpv6NeighborSolicitation
	FromProto(msg *otg.FlowIcmpv6NeighborSolicitation) (FlowIcmpv6NeighborSolicitation, error)
	// FromPbText unmarshals FlowIcmpv6NeighborSolicitation from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIcmpv6NeighborSolicitation from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIcmpv6NeighborSolicitation from JSON text
	FromJson(value string) error
}

func (obj *flowIcmpv6NeighborSolicitation) Marshal() marshalFlowIcmpv6NeighborSolicitation {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIcmpv6NeighborSolicitation{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIcmpv6NeighborSolicitation) Unmarshal() unMarshalFlowIcmpv6NeighborSolicitation {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIcmpv6NeighborSolicitation{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIcmpv6NeighborSolicitation) ToProto() (*otg.FlowIcmpv6NeighborSolicitation, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIcmpv6NeighborSolicitation) FromProto(msg *otg.FlowIcmpv6NeighborSolicitation) (FlowIcmpv6NeighborSolicitation, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIcmpv6NeighborSolicitation) ToPbText() (string, error) {
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

func (m *unMarshalflowIcmpv6NeighborSolicitation) FromPbText(value string) error {
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

func (m *marshalflowIcmpv6NeighborSolicitation) ToYaml() (string, error) {
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

func (m *unMarshalflowIcmpv6NeighborSolicitation) FromYaml(value string) error {
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

func (m *marshalflowIcmpv6NeighborSolicitation) ToJson() (string, error) {
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

func (m *unMarshalflowIcmpv6NeighborSolicitation) FromJson(value string) error {
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

func (obj *flowIcmpv6NeighborSolicitation) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIcmpv6NeighborSolicitation) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIcmpv6NeighborSolicitation) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIcmpv6NeighborSolicitation) Clone() (FlowIcmpv6NeighborSolicitation, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIcmpv6NeighborSolicitation()
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

func (obj *flowIcmpv6NeighborSolicitation) setNil() {
	obj.typeHolder = nil
	obj.codeHolder = nil
	obj.checksumHolder = nil
	obj.reservedHolder = nil
	obj.targetAddressHolder = nil
	obj.optionsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIcmpv6NeighborSolicitation is packet Header for ICMPv6 NDP Neighbor Solicitation (Type 135, RFC 4861 Section 4.3).
type FlowIcmpv6NeighborSolicitation interface {
	Validation
	// msg marshals FlowIcmpv6NeighborSolicitation to protobuf object *otg.FlowIcmpv6NeighborSolicitation
	// and doesn't set defaults
	msg() *otg.FlowIcmpv6NeighborSolicitation
	// setMsg unmarshals FlowIcmpv6NeighborSolicitation from protobuf object *otg.FlowIcmpv6NeighborSolicitation
	// and doesn't set defaults
	setMsg(*otg.FlowIcmpv6NeighborSolicitation) FlowIcmpv6NeighborSolicitation
	// provides marshal interface
	Marshal() marshalFlowIcmpv6NeighborSolicitation
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIcmpv6NeighborSolicitation
	// validate validates FlowIcmpv6NeighborSolicitation
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIcmpv6NeighborSolicitation, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns PatternFlowIcmpv6NeighborSolicitationType, set in FlowIcmpv6NeighborSolicitation.
	// PatternFlowIcmpv6NeighborSolicitationType is iCMPv6 message type. Fixed at 135 for Neighbor Solicitation.
	Type() PatternFlowIcmpv6NeighborSolicitationType
	// SetType assigns PatternFlowIcmpv6NeighborSolicitationType provided by user to FlowIcmpv6NeighborSolicitation.
	// PatternFlowIcmpv6NeighborSolicitationType is iCMPv6 message type. Fixed at 135 for Neighbor Solicitation.
	SetType(value PatternFlowIcmpv6NeighborSolicitationType) FlowIcmpv6NeighborSolicitation
	// HasType checks if Type has been set in FlowIcmpv6NeighborSolicitation
	HasType() bool
	// Code returns PatternFlowIcmpv6NeighborSolicitationCode, set in FlowIcmpv6NeighborSolicitation.
	// PatternFlowIcmpv6NeighborSolicitationCode is iCMPv6 message code. Must be 0 for Neighbor Solicitation.
	Code() PatternFlowIcmpv6NeighborSolicitationCode
	// SetCode assigns PatternFlowIcmpv6NeighborSolicitationCode provided by user to FlowIcmpv6NeighborSolicitation.
	// PatternFlowIcmpv6NeighborSolicitationCode is iCMPv6 message code. Must be 0 for Neighbor Solicitation.
	SetCode(value PatternFlowIcmpv6NeighborSolicitationCode) FlowIcmpv6NeighborSolicitation
	// HasCode checks if Code has been set in FlowIcmpv6NeighborSolicitation
	HasCode() bool
	// Checksum returns PatternFlowIcmpv6NeighborSolicitationChecksum, set in FlowIcmpv6NeighborSolicitation.
	// PatternFlowIcmpv6NeighborSolicitationChecksum is iCMPv6 checksum
	Checksum() PatternFlowIcmpv6NeighborSolicitationChecksum
	// SetChecksum assigns PatternFlowIcmpv6NeighborSolicitationChecksum provided by user to FlowIcmpv6NeighborSolicitation.
	// PatternFlowIcmpv6NeighborSolicitationChecksum is iCMPv6 checksum
	SetChecksum(value PatternFlowIcmpv6NeighborSolicitationChecksum) FlowIcmpv6NeighborSolicitation
	// HasChecksum checks if Checksum has been set in FlowIcmpv6NeighborSolicitation
	HasChecksum() bool
	// Reserved returns PatternFlowIcmpv6NeighborSolicitationReserved, set in FlowIcmpv6NeighborSolicitation.
	// PatternFlowIcmpv6NeighborSolicitationReserved is reserved field. Must be initialized to zero by the sender and must be ignored by the receiver.
	Reserved() PatternFlowIcmpv6NeighborSolicitationReserved
	// SetReserved assigns PatternFlowIcmpv6NeighborSolicitationReserved provided by user to FlowIcmpv6NeighborSolicitation.
	// PatternFlowIcmpv6NeighborSolicitationReserved is reserved field. Must be initialized to zero by the sender and must be ignored by the receiver.
	SetReserved(value PatternFlowIcmpv6NeighborSolicitationReserved) FlowIcmpv6NeighborSolicitation
	// HasReserved checks if Reserved has been set in FlowIcmpv6NeighborSolicitation
	HasReserved() bool
	// TargetAddress returns PatternFlowIcmpv6NeighborSolicitationTargetAddress, set in FlowIcmpv6NeighborSolicitation.
	// PatternFlowIcmpv6NeighborSolicitationTargetAddress is the IPv6 address of the target of the solicitation. Must not be a multicast address.
	TargetAddress() PatternFlowIcmpv6NeighborSolicitationTargetAddress
	// SetTargetAddress assigns PatternFlowIcmpv6NeighborSolicitationTargetAddress provided by user to FlowIcmpv6NeighborSolicitation.
	// PatternFlowIcmpv6NeighborSolicitationTargetAddress is the IPv6 address of the target of the solicitation. Must not be a multicast address.
	SetTargetAddress(value PatternFlowIcmpv6NeighborSolicitationTargetAddress) FlowIcmpv6NeighborSolicitation
	// HasTargetAddress checks if TargetAddress has been set in FlowIcmpv6NeighborSolicitation
	HasTargetAddress() bool
	// Options returns FlowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIterIter, set in FlowIcmpv6NeighborSolicitation
	Options() FlowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter
	setNil()
}

// description is TBD
// Type returns a PatternFlowIcmpv6NeighborSolicitationType
func (obj *flowIcmpv6NeighborSolicitation) Type() PatternFlowIcmpv6NeighborSolicitationType {
	if obj.obj.Type == nil {
		obj.obj.Type = NewPatternFlowIcmpv6NeighborSolicitationType().msg()
	}
	if obj.typeHolder == nil {
		obj.typeHolder = &patternFlowIcmpv6NeighborSolicitationType{obj: obj.obj.Type}
	}
	return obj.typeHolder
}

// description is TBD
// Type returns a PatternFlowIcmpv6NeighborSolicitationType
func (obj *flowIcmpv6NeighborSolicitation) HasType() bool {
	return obj.obj.Type != nil
}

// description is TBD
// SetType sets the PatternFlowIcmpv6NeighborSolicitationType value in the FlowIcmpv6NeighborSolicitation object
func (obj *flowIcmpv6NeighborSolicitation) SetType(value PatternFlowIcmpv6NeighborSolicitationType) FlowIcmpv6NeighborSolicitation {

	obj.typeHolder = nil
	obj.obj.Type = value.msg()

	return obj
}

// description is TBD
// Code returns a PatternFlowIcmpv6NeighborSolicitationCode
func (obj *flowIcmpv6NeighborSolicitation) Code() PatternFlowIcmpv6NeighborSolicitationCode {
	if obj.obj.Code == nil {
		obj.obj.Code = NewPatternFlowIcmpv6NeighborSolicitationCode().msg()
	}
	if obj.codeHolder == nil {
		obj.codeHolder = &patternFlowIcmpv6NeighborSolicitationCode{obj: obj.obj.Code}
	}
	return obj.codeHolder
}

// description is TBD
// Code returns a PatternFlowIcmpv6NeighborSolicitationCode
func (obj *flowIcmpv6NeighborSolicitation) HasCode() bool {
	return obj.obj.Code != nil
}

// description is TBD
// SetCode sets the PatternFlowIcmpv6NeighborSolicitationCode value in the FlowIcmpv6NeighborSolicitation object
func (obj *flowIcmpv6NeighborSolicitation) SetCode(value PatternFlowIcmpv6NeighborSolicitationCode) FlowIcmpv6NeighborSolicitation {

	obj.codeHolder = nil
	obj.obj.Code = value.msg()

	return obj
}

// description is TBD
// Checksum returns a PatternFlowIcmpv6NeighborSolicitationChecksum
func (obj *flowIcmpv6NeighborSolicitation) Checksum() PatternFlowIcmpv6NeighborSolicitationChecksum {
	if obj.obj.Checksum == nil {
		obj.obj.Checksum = NewPatternFlowIcmpv6NeighborSolicitationChecksum().msg()
	}
	if obj.checksumHolder == nil {
		obj.checksumHolder = &patternFlowIcmpv6NeighborSolicitationChecksum{obj: obj.obj.Checksum}
	}
	return obj.checksumHolder
}

// description is TBD
// Checksum returns a PatternFlowIcmpv6NeighborSolicitationChecksum
func (obj *flowIcmpv6NeighborSolicitation) HasChecksum() bool {
	return obj.obj.Checksum != nil
}

// description is TBD
// SetChecksum sets the PatternFlowIcmpv6NeighborSolicitationChecksum value in the FlowIcmpv6NeighborSolicitation object
func (obj *flowIcmpv6NeighborSolicitation) SetChecksum(value PatternFlowIcmpv6NeighborSolicitationChecksum) FlowIcmpv6NeighborSolicitation {

	obj.checksumHolder = nil
	obj.obj.Checksum = value.msg()

	return obj
}

// description is TBD
// Reserved returns a PatternFlowIcmpv6NeighborSolicitationReserved
func (obj *flowIcmpv6NeighborSolicitation) Reserved() PatternFlowIcmpv6NeighborSolicitationReserved {
	if obj.obj.Reserved == nil {
		obj.obj.Reserved = NewPatternFlowIcmpv6NeighborSolicitationReserved().msg()
	}
	if obj.reservedHolder == nil {
		obj.reservedHolder = &patternFlowIcmpv6NeighborSolicitationReserved{obj: obj.obj.Reserved}
	}
	return obj.reservedHolder
}

// description is TBD
// Reserved returns a PatternFlowIcmpv6NeighborSolicitationReserved
func (obj *flowIcmpv6NeighborSolicitation) HasReserved() bool {
	return obj.obj.Reserved != nil
}

// description is TBD
// SetReserved sets the PatternFlowIcmpv6NeighborSolicitationReserved value in the FlowIcmpv6NeighborSolicitation object
func (obj *flowIcmpv6NeighborSolicitation) SetReserved(value PatternFlowIcmpv6NeighborSolicitationReserved) FlowIcmpv6NeighborSolicitation {

	obj.reservedHolder = nil
	obj.obj.Reserved = value.msg()

	return obj
}

// description is TBD
// TargetAddress returns a PatternFlowIcmpv6NeighborSolicitationTargetAddress
func (obj *flowIcmpv6NeighborSolicitation) TargetAddress() PatternFlowIcmpv6NeighborSolicitationTargetAddress {
	if obj.obj.TargetAddress == nil {
		obj.obj.TargetAddress = NewPatternFlowIcmpv6NeighborSolicitationTargetAddress().msg()
	}
	if obj.targetAddressHolder == nil {
		obj.targetAddressHolder = &patternFlowIcmpv6NeighborSolicitationTargetAddress{obj: obj.obj.TargetAddress}
	}
	return obj.targetAddressHolder
}

// description is TBD
// TargetAddress returns a PatternFlowIcmpv6NeighborSolicitationTargetAddress
func (obj *flowIcmpv6NeighborSolicitation) HasTargetAddress() bool {
	return obj.obj.TargetAddress != nil
}

// description is TBD
// SetTargetAddress sets the PatternFlowIcmpv6NeighborSolicitationTargetAddress value in the FlowIcmpv6NeighborSolicitation object
func (obj *flowIcmpv6NeighborSolicitation) SetTargetAddress(value PatternFlowIcmpv6NeighborSolicitationTargetAddress) FlowIcmpv6NeighborSolicitation {

	obj.targetAddressHolder = nil
	obj.obj.TargetAddress = value.msg()

	return obj
}

// Optional list of NDP options. Source Link-Layer Address SHOULD be included when sent from a unicast address (RFC 4861 Section 4.3).
// Options returns a []FlowIcmpv6NdpOption
func (obj *flowIcmpv6NeighborSolicitation) Options() FlowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter {
	if len(obj.obj.Options) == 0 {
		obj.obj.Options = []*otg.FlowIcmpv6NdpOption{}
	}
	if obj.optionsHolder == nil {
		obj.optionsHolder = newFlowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter(&obj.obj.Options).setMsg(obj)
	}
	return obj.optionsHolder
}

type flowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter struct {
	obj                      *flowIcmpv6NeighborSolicitation
	flowIcmpv6NdpOptionSlice []FlowIcmpv6NdpOption
	fieldPtr                 *[]*otg.FlowIcmpv6NdpOption
}

func newFlowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter(ptr *[]*otg.FlowIcmpv6NdpOption) FlowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter {
	return &flowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter{fieldPtr: ptr}
}

type FlowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter interface {
	setMsg(*flowIcmpv6NeighborSolicitation) FlowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter
	Items() []FlowIcmpv6NdpOption
	Add() FlowIcmpv6NdpOption
	Append(items ...FlowIcmpv6NdpOption) FlowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter
	Set(index int, newObj FlowIcmpv6NdpOption) FlowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter
	Clear() FlowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter
	clearHolderSlice() FlowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter
	appendHolderSlice(item FlowIcmpv6NdpOption) FlowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter
}

func (obj *flowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter) setMsg(msg *flowIcmpv6NeighborSolicitation) FlowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowIcmpv6NdpOption{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter) Items() []FlowIcmpv6NdpOption {
	return obj.flowIcmpv6NdpOptionSlice
}

func (obj *flowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter) Add() FlowIcmpv6NdpOption {
	newObj := &otg.FlowIcmpv6NdpOption{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowIcmpv6NdpOption{obj: newObj}
	newLibObj.setDefault()
	obj.flowIcmpv6NdpOptionSlice = append(obj.flowIcmpv6NdpOptionSlice, newLibObj)
	return newLibObj
}

func (obj *flowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter) Append(items ...FlowIcmpv6NdpOption) FlowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowIcmpv6NdpOptionSlice = append(obj.flowIcmpv6NdpOptionSlice, item)
	}
	return obj
}

func (obj *flowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter) Set(index int, newObj FlowIcmpv6NdpOption) FlowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowIcmpv6NdpOptionSlice[index] = newObj
	return obj
}
func (obj *flowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter) Clear() FlowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowIcmpv6NdpOption{}
		obj.flowIcmpv6NdpOptionSlice = []FlowIcmpv6NdpOption{}
	}
	return obj
}
func (obj *flowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter) clearHolderSlice() FlowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter {
	if len(obj.flowIcmpv6NdpOptionSlice) > 0 {
		obj.flowIcmpv6NdpOptionSlice = []FlowIcmpv6NdpOption{}
	}
	return obj
}
func (obj *flowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter) appendHolderSlice(item FlowIcmpv6NdpOption) FlowIcmpv6NeighborSolicitationFlowIcmpv6NdpOptionIter {
	obj.flowIcmpv6NdpOptionSlice = append(obj.flowIcmpv6NdpOptionSlice, item)
	return obj
}

func (obj *flowIcmpv6NeighborSolicitation) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Type != nil {

		obj.Type().validateObj(vObj, set_default)
	}

	if obj.obj.Code != nil {

		obj.Code().validateObj(vObj, set_default)
	}

	if obj.obj.Checksum != nil {

		obj.Checksum().validateObj(vObj, set_default)
	}

	if obj.obj.Reserved != nil {

		obj.Reserved().validateObj(vObj, set_default)
	}

	if obj.obj.TargetAddress != nil {

		obj.TargetAddress().validateObj(vObj, set_default)
	}

	if len(obj.obj.Options) != 0 {

		if set_default {
			obj.Options().clearHolderSlice()
			for _, item := range obj.obj.Options {
				obj.Options().appendHolderSlice(&flowIcmpv6NdpOption{obj: item})
			}
		}
		for _, item := range obj.Options().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *flowIcmpv6NeighborSolicitation) setDefault() {

}
