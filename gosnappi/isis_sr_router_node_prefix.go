package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisSRRouterNodePrefix *****
type isisSRRouterNodePrefix struct {
	validation
	obj                    *otg.IsisSRRouterNodePrefix
	marshaller             marshalIsisSRRouterNodePrefix
	unMarshaller           unMarshalIsisSRRouterNodePrefix
	customNodePrefixHolder IsisSRIP
	prefixSidParamsHolder  IsisSRPrefixSID
}

func NewIsisSRRouterNodePrefix() IsisSRRouterNodePrefix {
	obj := isisSRRouterNodePrefix{obj: &otg.IsisSRRouterNodePrefix{}}
	obj.setDefault()
	return &obj
}

func (obj *isisSRRouterNodePrefix) msg() *otg.IsisSRRouterNodePrefix {
	return obj.obj
}

func (obj *isisSRRouterNodePrefix) setMsg(msg *otg.IsisSRRouterNodePrefix) IsisSRRouterNodePrefix {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisSRRouterNodePrefix struct {
	obj *isisSRRouterNodePrefix
}

type marshalIsisSRRouterNodePrefix interface {
	// ToProto marshals IsisSRRouterNodePrefix to protobuf object *otg.IsisSRRouterNodePrefix
	ToProto() (*otg.IsisSRRouterNodePrefix, error)
	// ToPbText marshals IsisSRRouterNodePrefix to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisSRRouterNodePrefix to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisSRRouterNodePrefix to JSON text
	ToJson() (string, error)
}

type unMarshalisisSRRouterNodePrefix struct {
	obj *isisSRRouterNodePrefix
}

type unMarshalIsisSRRouterNodePrefix interface {
	// FromProto unmarshals IsisSRRouterNodePrefix from protobuf object *otg.IsisSRRouterNodePrefix
	FromProto(msg *otg.IsisSRRouterNodePrefix) (IsisSRRouterNodePrefix, error)
	// FromPbText unmarshals IsisSRRouterNodePrefix from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisSRRouterNodePrefix from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisSRRouterNodePrefix from JSON text
	FromJson(value string) error
}

func (obj *isisSRRouterNodePrefix) Marshal() marshalIsisSRRouterNodePrefix {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisSRRouterNodePrefix{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisSRRouterNodePrefix) Unmarshal() unMarshalIsisSRRouterNodePrefix {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisSRRouterNodePrefix{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisSRRouterNodePrefix) ToProto() (*otg.IsisSRRouterNodePrefix, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisSRRouterNodePrefix) FromProto(msg *otg.IsisSRRouterNodePrefix) (IsisSRRouterNodePrefix, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisSRRouterNodePrefix) ToPbText() (string, error) {
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

func (m *unMarshalisisSRRouterNodePrefix) FromPbText(value string) error {
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

func (m *marshalisisSRRouterNodePrefix) ToYaml() (string, error) {
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

func (m *unMarshalisisSRRouterNodePrefix) FromYaml(value string) error {
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

func (m *marshalisisSRRouterNodePrefix) ToJson() (string, error) {
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

func (m *unMarshalisisSRRouterNodePrefix) FromJson(value string) error {
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

func (obj *isisSRRouterNodePrefix) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisSRRouterNodePrefix) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisSRRouterNodePrefix) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisSRRouterNodePrefix) Clone() (IsisSRRouterNodePrefix, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisSRRouterNodePrefix()
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

func (obj *isisSRRouterNodePrefix) setNil() {
	obj.customNodePrefixHolder = nil
	obj.prefixSidParamsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisSRRouterNodePrefix is sID/Index/Label: As defined in Section 2.1.1.1. Reference: https://datatracker.ietf.org/doc/html/rfc8667#VANDLFLAGS.
type IsisSRRouterNodePrefix interface {
	Validation
	// msg marshals IsisSRRouterNodePrefix to protobuf object *otg.IsisSRRouterNodePrefix
	// and doesn't set defaults
	msg() *otg.IsisSRRouterNodePrefix
	// setMsg unmarshals IsisSRRouterNodePrefix from protobuf object *otg.IsisSRRouterNodePrefix
	// and doesn't set defaults
	setMsg(*otg.IsisSRRouterNodePrefix) IsisSRRouterNodePrefix
	// provides marshal interface
	Marshal() marshalIsisSRRouterNodePrefix
	// provides unmarshal interface
	Unmarshal() unMarshalIsisSRRouterNodePrefix
	// validate validates IsisSRRouterNodePrefix
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisSRRouterNodePrefix, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns IsisSRRouterNodePrefixChoiceEnum, set in IsisSRRouterNodePrefix
	Choice() IsisSRRouterNodePrefixChoiceEnum
	// setChoice assigns IsisSRRouterNodePrefixChoiceEnum provided by user to IsisSRRouterNodePrefix
	setChoice(value IsisSRRouterNodePrefixChoiceEnum) IsisSRRouterNodePrefix
	// HasChoice checks if Choice has been set in IsisSRRouterNodePrefix
	HasChoice() bool
	// getter for Ipv4TeRouterId to set choice.
	Ipv4TeRouterId()
	// getter for InterfaceIp to set choice.
	InterfaceIp()
	// CustomNodePrefix returns IsisSRIP, set in IsisSRRouterNodePrefix.
	// IsisSRIP is tBD
	CustomNodePrefix() IsisSRIP
	// SetCustomNodePrefix assigns IsisSRIP provided by user to IsisSRRouterNodePrefix.
	// IsisSRIP is tBD
	SetCustomNodePrefix(value IsisSRIP) IsisSRRouterNodePrefix
	// HasCustomNodePrefix checks if CustomNodePrefix has been set in IsisSRRouterNodePrefix
	HasCustomNodePrefix() bool
	// Redistribution returns IsisSRRouterNodePrefixRedistributionEnum, set in IsisSRRouterNodePrefix
	Redistribution() IsisSRRouterNodePrefixRedistributionEnum
	// SetRedistribution assigns IsisSRRouterNodePrefixRedistributionEnum provided by user to IsisSRRouterNodePrefix
	SetRedistribution(value IsisSRRouterNodePrefixRedistributionEnum) IsisSRRouterNodePrefix
	// HasRedistribution checks if Redistribution has been set in IsisSRRouterNodePrefix
	HasRedistribution() bool
	// PrefixAttrEnabled returns bool, set in IsisSRRouterNodePrefix.
	PrefixAttrEnabled() bool
	// SetPrefixAttrEnabled assigns bool provided by user to IsisSRRouterNodePrefix
	SetPrefixAttrEnabled(value bool) IsisSRRouterNodePrefix
	// HasPrefixAttrEnabled checks if PrefixAttrEnabled has been set in IsisSRRouterNodePrefix
	HasPrefixAttrEnabled() bool
	// XFlag returns bool, set in IsisSRRouterNodePrefix.
	XFlag() bool
	// SetXFlag assigns bool provided by user to IsisSRRouterNodePrefix
	SetXFlag(value bool) IsisSRRouterNodePrefix
	// HasXFlag checks if XFlag has been set in IsisSRRouterNodePrefix
	HasXFlag() bool
	// RFlag returns bool, set in IsisSRRouterNodePrefix.
	RFlag() bool
	// SetRFlag assigns bool provided by user to IsisSRRouterNodePrefix
	SetRFlag(value bool) IsisSRRouterNodePrefix
	// HasRFlag checks if RFlag has been set in IsisSRRouterNodePrefix
	HasRFlag() bool
	// NFlag returns bool, set in IsisSRRouterNodePrefix.
	NFlag() bool
	// SetNFlag assigns bool provided by user to IsisSRRouterNodePrefix
	SetNFlag(value bool) IsisSRRouterNodePrefix
	// HasNFlag checks if NFlag has been set in IsisSRRouterNodePrefix
	HasNFlag() bool
	// PrefixSidParams returns IsisSRPrefixSID, set in IsisSRRouterNodePrefix.
	// IsisSRPrefixSID is this contains the properties of IS-IS Prefix-SID and its attributes for  the extended Ipv4 and Ipv6 reachability. Refernce: https://datatracker.ietf.org/doc/html/rfc8667#name-prefix-segment-identifier-p.
	PrefixSidParams() IsisSRPrefixSID
	// SetPrefixSidParams assigns IsisSRPrefixSID provided by user to IsisSRRouterNodePrefix.
	// IsisSRPrefixSID is this contains the properties of IS-IS Prefix-SID and its attributes for  the extended Ipv4 and Ipv6 reachability. Refernce: https://datatracker.ietf.org/doc/html/rfc8667#name-prefix-segment-identifier-p.
	SetPrefixSidParams(value IsisSRPrefixSID) IsisSRRouterNodePrefix
	// HasPrefixSidParams checks if PrefixSidParams has been set in IsisSRRouterNodePrefix
	HasPrefixSidParams() bool
	setNil()
}

type IsisSRRouterNodePrefixChoiceEnum string

// Enum of Choice on IsisSRRouterNodePrefix
var IsisSRRouterNodePrefixChoice = struct {
	IPV4_TE_ROUTER_ID  IsisSRRouterNodePrefixChoiceEnum
	INTERFACE_IP       IsisSRRouterNodePrefixChoiceEnum
	CUSTOM_NODE_PREFIX IsisSRRouterNodePrefixChoiceEnum
}{
	IPV4_TE_ROUTER_ID:  IsisSRRouterNodePrefixChoiceEnum("ipv4_te_router_id"),
	INTERFACE_IP:       IsisSRRouterNodePrefixChoiceEnum("interface_ip"),
	CUSTOM_NODE_PREFIX: IsisSRRouterNodePrefixChoiceEnum("custom_node_prefix"),
}

func (obj *isisSRRouterNodePrefix) Choice() IsisSRRouterNodePrefixChoiceEnum {
	return IsisSRRouterNodePrefixChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for Ipv4TeRouterId to set choice
func (obj *isisSRRouterNodePrefix) Ipv4TeRouterId() {
	obj.setChoice(IsisSRRouterNodePrefixChoice.IPV4_TE_ROUTER_ID)
}

// getter for InterfaceIp to set choice
func (obj *isisSRRouterNodePrefix) InterfaceIp() {
	obj.setChoice(IsisSRRouterNodePrefixChoice.INTERFACE_IP)
}

// Node Prefix Segment Identifier (Prefix-SID).
// - ipv4_te_router_id: IPv4 Traffic Engineering(TE) router id to be used as Router Capability ID.
// - interface_ip: When underlined IPv4 interface address of the ISIS interface to be assigned as Router Capability ID.
// - custom: When, Router Capability ID needs to be configured different from above two.
// Choice returns a string
func (obj *isisSRRouterNodePrefix) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *isisSRRouterNodePrefix) setChoice(value IsisSRRouterNodePrefixChoiceEnum) IsisSRRouterNodePrefix {
	intValue, ok := otg.IsisSRRouterNodePrefix_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisSRRouterNodePrefixChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisSRRouterNodePrefix_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.CustomNodePrefix = nil
	obj.customNodePrefixHolder = nil

	if value == IsisSRRouterNodePrefixChoice.CUSTOM_NODE_PREFIX {
		obj.obj.CustomNodePrefix = NewIsisSRIP().msg()
	}

	return obj
}

// The node prefix taken in IPv4 format.
// CustomNodePrefix returns a IsisSRIP
func (obj *isisSRRouterNodePrefix) CustomNodePrefix() IsisSRIP {
	if obj.obj.CustomNodePrefix == nil {
		obj.setChoice(IsisSRRouterNodePrefixChoice.CUSTOM_NODE_PREFIX)
	}
	if obj.customNodePrefixHolder == nil {
		obj.customNodePrefixHolder = &isisSRIP{obj: obj.obj.CustomNodePrefix}
	}
	return obj.customNodePrefixHolder
}

// The node prefix taken in IPv4 format.
// CustomNodePrefix returns a IsisSRIP
func (obj *isisSRRouterNodePrefix) HasCustomNodePrefix() bool {
	return obj.obj.CustomNodePrefix != nil
}

// The node prefix taken in IPv4 format.
// SetCustomNodePrefix sets the IsisSRIP value in the IsisSRRouterNodePrefix object
func (obj *isisSRRouterNodePrefix) SetCustomNodePrefix(value IsisSRIP) IsisSRRouterNodePrefix {
	obj.setChoice(IsisSRRouterNodePrefixChoice.CUSTOM_NODE_PREFIX)
	obj.customNodePrefixHolder = nil
	obj.obj.CustomNodePrefix = value.msg()

	return obj
}

type IsisSRRouterNodePrefixRedistributionEnum string

// Enum of Redistribution on IsisSRRouterNodePrefix
var IsisSRRouterNodePrefixRedistribution = struct {
	UP   IsisSRRouterNodePrefixRedistributionEnum
	DOWN IsisSRRouterNodePrefixRedistributionEnum
}{
	UP:   IsisSRRouterNodePrefixRedistributionEnum("up"),
	DOWN: IsisSRRouterNodePrefixRedistributionEnum("down"),
}

func (obj *isisSRRouterNodePrefix) Redistribution() IsisSRRouterNodePrefixRedistributionEnum {
	return IsisSRRouterNodePrefixRedistributionEnum(obj.obj.Redistribution.Enum().String())
}

// Defines the Up/Down (Redistribution) bit defined for TLVs 128 and 130 by RFC 2966.
// It is used for domain-wide advertisement of prefix information.
// If up - used when a prefix is initially advertised within the ISIS L3 hierarchy, and for all other prefixes in L1 and L2 LSPs. (default)
// When down - used when an L1/L2 router advertises L2 prefixes in L1 LSPs.
// The prefixes are being advertised from a higher level (L2) down to a lower level (L1).
// Redistribution returns a string
func (obj *isisSRRouterNodePrefix) HasRedistribution() bool {
	return obj.obj.Redistribution != nil
}

func (obj *isisSRRouterNodePrefix) SetRedistribution(value IsisSRRouterNodePrefixRedistributionEnum) IsisSRRouterNodePrefix {
	intValue, ok := otg.IsisSRRouterNodePrefix_Redistribution_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisSRRouterNodePrefixRedistributionEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisSRRouterNodePrefix_Redistribution_Enum(intValue)
	obj.obj.Redistribution = &enumValue

	return obj
}

// Specifies whether the sub-TLV for IPv4/IPv6 Extended Reachability Attribute Flags
// will be advertised or not.
// PrefixAttrEnabled returns a bool
func (obj *isisSRRouterNodePrefix) PrefixAttrEnabled() bool {

	return *obj.obj.PrefixAttrEnabled

}

// Specifies whether the sub-TLV for IPv4/IPv6 Extended Reachability Attribute Flags
// will be advertised or not.
// PrefixAttrEnabled returns a bool
func (obj *isisSRRouterNodePrefix) HasPrefixAttrEnabled() bool {
	return obj.obj.PrefixAttrEnabled != nil
}

// Specifies whether the sub-TLV for IPv4/IPv6 Extended Reachability Attribute Flags
// will be advertised or not.
// SetPrefixAttrEnabled sets the bool value in the IsisSRRouterNodePrefix object
func (obj *isisSRRouterNodePrefix) SetPrefixAttrEnabled(value bool) IsisSRRouterNodePrefix {

	obj.obj.PrefixAttrEnabled = &value
	return obj
}

// External Prefix Flag (Bit 0)
// XFlag returns a bool
func (obj *isisSRRouterNodePrefix) XFlag() bool {

	return *obj.obj.XFlag

}

// External Prefix Flag (Bit 0)
// XFlag returns a bool
func (obj *isisSRRouterNodePrefix) HasXFlag() bool {
	return obj.obj.XFlag != nil
}

// External Prefix Flag (Bit 0)
// SetXFlag sets the bool value in the IsisSRRouterNodePrefix object
func (obj *isisSRRouterNodePrefix) SetXFlag(value bool) IsisSRRouterNodePrefix {

	obj.obj.XFlag = &value
	return obj
}

// Re-advertisement Flag (Bit 1)
// RFlag returns a bool
func (obj *isisSRRouterNodePrefix) RFlag() bool {

	return *obj.obj.RFlag

}

// Re-advertisement Flag (Bit 1)
// RFlag returns a bool
func (obj *isisSRRouterNodePrefix) HasRFlag() bool {
	return obj.obj.RFlag != nil
}

// Re-advertisement Flag (Bit 1)
// SetRFlag sets the bool value in the IsisSRRouterNodePrefix object
func (obj *isisSRRouterNodePrefix) SetRFlag(value bool) IsisSRRouterNodePrefix {

	obj.obj.RFlag = &value
	return obj
}

// Node Flag (Bit 2)
// NFlag returns a bool
func (obj *isisSRRouterNodePrefix) NFlag() bool {

	return *obj.obj.NFlag

}

// Node Flag (Bit 2)
// NFlag returns a bool
func (obj *isisSRRouterNodePrefix) HasNFlag() bool {
	return obj.obj.NFlag != nil
}

// Node Flag (Bit 2)
// SetNFlag sets the bool value in the IsisSRRouterNodePrefix object
func (obj *isisSRRouterNodePrefix) SetNFlag(value bool) IsisSRRouterNodePrefix {

	obj.obj.NFlag = &value
	return obj
}

// description is TBD
// PrefixSidParams returns a IsisSRPrefixSID
func (obj *isisSRRouterNodePrefix) PrefixSidParams() IsisSRPrefixSID {
	if obj.obj.PrefixSidParams == nil {
		obj.obj.PrefixSidParams = NewIsisSRPrefixSID().msg()
	}
	if obj.prefixSidParamsHolder == nil {
		obj.prefixSidParamsHolder = &isisSRPrefixSID{obj: obj.obj.PrefixSidParams}
	}
	return obj.prefixSidParamsHolder
}

// description is TBD
// PrefixSidParams returns a IsisSRPrefixSID
func (obj *isisSRRouterNodePrefix) HasPrefixSidParams() bool {
	return obj.obj.PrefixSidParams != nil
}

// description is TBD
// SetPrefixSidParams sets the IsisSRPrefixSID value in the IsisSRRouterNodePrefix object
func (obj *isisSRRouterNodePrefix) SetPrefixSidParams(value IsisSRPrefixSID) IsisSRRouterNodePrefix {

	obj.prefixSidParamsHolder = nil
	obj.obj.PrefixSidParams = value.msg()

	return obj
}

func (obj *isisSRRouterNodePrefix) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.CustomNodePrefix != nil {

		obj.CustomNodePrefix().validateObj(vObj, set_default)
	}

	if obj.obj.PrefixSidParams != nil {

		obj.PrefixSidParams().validateObj(vObj, set_default)
	}

}

func (obj *isisSRRouterNodePrefix) setDefault() {
	var choices_set int = 0
	var choice IsisSRRouterNodePrefixChoiceEnum

	if obj.obj.CustomNodePrefix != nil {
		choices_set += 1
		choice = IsisSRRouterNodePrefixChoice.CUSTOM_NODE_PREFIX
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(IsisSRRouterNodePrefixChoice.IPV4_TE_ROUTER_ID)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in IsisSRRouterNodePrefix")
			}
		} else {
			intVal := otg.IsisSRRouterNodePrefix_Choice_Enum_value[string(choice)]
			enumValue := otg.IsisSRRouterNodePrefix_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.Redistribution == nil {
		obj.SetRedistribution(IsisSRRouterNodePrefixRedistribution.UP)

	}
	if obj.obj.PrefixAttrEnabled == nil {
		obj.SetPrefixAttrEnabled(false)
	}
	if obj.obj.XFlag == nil {
		obj.SetXFlag(false)
	}
	if obj.obj.RFlag == nil {
		obj.SetRFlag(false)
	}
	if obj.obj.NFlag == nil {
		obj.SetNFlag(false)
	}

}
