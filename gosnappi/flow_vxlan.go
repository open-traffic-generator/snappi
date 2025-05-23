package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowVxlan *****
type flowVxlan struct {
	validation
	obj             *otg.FlowVxlan
	marshaller      marshalFlowVxlan
	unMarshaller    unMarshalFlowVxlan
	flagsHolder     PatternFlowVxlanFlags
	reserved0Holder PatternFlowVxlanReserved0
	vniHolder       PatternFlowVxlanVni
	reserved1Holder PatternFlowVxlanReserved1
}

func NewFlowVxlan() FlowVxlan {
	obj := flowVxlan{obj: &otg.FlowVxlan{}}
	obj.setDefault()
	return &obj
}

func (obj *flowVxlan) msg() *otg.FlowVxlan {
	return obj.obj
}

func (obj *flowVxlan) setMsg(msg *otg.FlowVxlan) FlowVxlan {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowVxlan struct {
	obj *flowVxlan
}

type marshalFlowVxlan interface {
	// ToProto marshals FlowVxlan to protobuf object *otg.FlowVxlan
	ToProto() (*otg.FlowVxlan, error)
	// ToPbText marshals FlowVxlan to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowVxlan to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowVxlan to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowVxlan to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowVxlan struct {
	obj *flowVxlan
}

type unMarshalFlowVxlan interface {
	// FromProto unmarshals FlowVxlan from protobuf object *otg.FlowVxlan
	FromProto(msg *otg.FlowVxlan) (FlowVxlan, error)
	// FromPbText unmarshals FlowVxlan from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowVxlan from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowVxlan from JSON text
	FromJson(value string) error
}

func (obj *flowVxlan) Marshal() marshalFlowVxlan {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowVxlan{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowVxlan) Unmarshal() unMarshalFlowVxlan {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowVxlan{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowVxlan) ToProto() (*otg.FlowVxlan, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowVxlan) FromProto(msg *otg.FlowVxlan) (FlowVxlan, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowVxlan) ToPbText() (string, error) {
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

func (m *unMarshalflowVxlan) FromPbText(value string) error {
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

func (m *marshalflowVxlan) ToYaml() (string, error) {
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

func (m *unMarshalflowVxlan) FromYaml(value string) error {
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

func (m *marshalflowVxlan) ToJsonRaw() (string, error) {
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

func (m *marshalflowVxlan) ToJson() (string, error) {
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

func (m *unMarshalflowVxlan) FromJson(value string) error {
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

func (obj *flowVxlan) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowVxlan) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowVxlan) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowVxlan) Clone() (FlowVxlan, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowVxlan()
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

func (obj *flowVxlan) setNil() {
	obj.flagsHolder = nil
	obj.reserved0Holder = nil
	obj.vniHolder = nil
	obj.reserved1Holder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowVxlan is vXLAN packet header
type FlowVxlan interface {
	Validation
	// msg marshals FlowVxlan to protobuf object *otg.FlowVxlan
	// and doesn't set defaults
	msg() *otg.FlowVxlan
	// setMsg unmarshals FlowVxlan from protobuf object *otg.FlowVxlan
	// and doesn't set defaults
	setMsg(*otg.FlowVxlan) FlowVxlan
	// provides marshal interface
	Marshal() marshalFlowVxlan
	// provides unmarshal interface
	Unmarshal() unMarshalFlowVxlan
	// validate validates FlowVxlan
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowVxlan, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns PatternFlowVxlanFlags, set in FlowVxlan.
	// PatternFlowVxlanFlags is flags field with a bit format of RRRRIRRR. The I flag MUST be set to 1 for a valid vxlan network id (VNI).   The other 7 bits (designated "R") are reserved fields and MUST be  set to zero on transmission and ignored on receipt.
	Flags() PatternFlowVxlanFlags
	// SetFlags assigns PatternFlowVxlanFlags provided by user to FlowVxlan.
	// PatternFlowVxlanFlags is flags field with a bit format of RRRRIRRR. The I flag MUST be set to 1 for a valid vxlan network id (VNI).   The other 7 bits (designated "R") are reserved fields and MUST be  set to zero on transmission and ignored on receipt.
	SetFlags(value PatternFlowVxlanFlags) FlowVxlan
	// HasFlags checks if Flags has been set in FlowVxlan
	HasFlags() bool
	// Reserved0 returns PatternFlowVxlanReserved0, set in FlowVxlan.
	// PatternFlowVxlanReserved0 is reserved field
	Reserved0() PatternFlowVxlanReserved0
	// SetReserved0 assigns PatternFlowVxlanReserved0 provided by user to FlowVxlan.
	// PatternFlowVxlanReserved0 is reserved field
	SetReserved0(value PatternFlowVxlanReserved0) FlowVxlan
	// HasReserved0 checks if Reserved0 has been set in FlowVxlan
	HasReserved0() bool
	// Vni returns PatternFlowVxlanVni, set in FlowVxlan.
	// PatternFlowVxlanVni is vXLAN network id
	Vni() PatternFlowVxlanVni
	// SetVni assigns PatternFlowVxlanVni provided by user to FlowVxlan.
	// PatternFlowVxlanVni is vXLAN network id
	SetVni(value PatternFlowVxlanVni) FlowVxlan
	// HasVni checks if Vni has been set in FlowVxlan
	HasVni() bool
	// Reserved1 returns PatternFlowVxlanReserved1, set in FlowVxlan.
	// PatternFlowVxlanReserved1 is reserved field
	Reserved1() PatternFlowVxlanReserved1
	// SetReserved1 assigns PatternFlowVxlanReserved1 provided by user to FlowVxlan.
	// PatternFlowVxlanReserved1 is reserved field
	SetReserved1(value PatternFlowVxlanReserved1) FlowVxlan
	// HasReserved1 checks if Reserved1 has been set in FlowVxlan
	HasReserved1() bool
	setNil()
}

// description is TBD
// Flags returns a PatternFlowVxlanFlags
func (obj *flowVxlan) Flags() PatternFlowVxlanFlags {
	if obj.obj.Flags == nil {
		obj.obj.Flags = NewPatternFlowVxlanFlags().msg()
	}
	if obj.flagsHolder == nil {
		obj.flagsHolder = &patternFlowVxlanFlags{obj: obj.obj.Flags}
	}
	return obj.flagsHolder
}

// description is TBD
// Flags returns a PatternFlowVxlanFlags
func (obj *flowVxlan) HasFlags() bool {
	return obj.obj.Flags != nil
}

// description is TBD
// SetFlags sets the PatternFlowVxlanFlags value in the FlowVxlan object
func (obj *flowVxlan) SetFlags(value PatternFlowVxlanFlags) FlowVxlan {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// description is TBD
// Reserved0 returns a PatternFlowVxlanReserved0
func (obj *flowVxlan) Reserved0() PatternFlowVxlanReserved0 {
	if obj.obj.Reserved0 == nil {
		obj.obj.Reserved0 = NewPatternFlowVxlanReserved0().msg()
	}
	if obj.reserved0Holder == nil {
		obj.reserved0Holder = &patternFlowVxlanReserved0{obj: obj.obj.Reserved0}
	}
	return obj.reserved0Holder
}

// description is TBD
// Reserved0 returns a PatternFlowVxlanReserved0
func (obj *flowVxlan) HasReserved0() bool {
	return obj.obj.Reserved0 != nil
}

// description is TBD
// SetReserved0 sets the PatternFlowVxlanReserved0 value in the FlowVxlan object
func (obj *flowVxlan) SetReserved0(value PatternFlowVxlanReserved0) FlowVxlan {

	obj.reserved0Holder = nil
	obj.obj.Reserved0 = value.msg()

	return obj
}

// description is TBD
// Vni returns a PatternFlowVxlanVni
func (obj *flowVxlan) Vni() PatternFlowVxlanVni {
	if obj.obj.Vni == nil {
		obj.obj.Vni = NewPatternFlowVxlanVni().msg()
	}
	if obj.vniHolder == nil {
		obj.vniHolder = &patternFlowVxlanVni{obj: obj.obj.Vni}
	}
	return obj.vniHolder
}

// description is TBD
// Vni returns a PatternFlowVxlanVni
func (obj *flowVxlan) HasVni() bool {
	return obj.obj.Vni != nil
}

// description is TBD
// SetVni sets the PatternFlowVxlanVni value in the FlowVxlan object
func (obj *flowVxlan) SetVni(value PatternFlowVxlanVni) FlowVxlan {

	obj.vniHolder = nil
	obj.obj.Vni = value.msg()

	return obj
}

// description is TBD
// Reserved1 returns a PatternFlowVxlanReserved1
func (obj *flowVxlan) Reserved1() PatternFlowVxlanReserved1 {
	if obj.obj.Reserved1 == nil {
		obj.obj.Reserved1 = NewPatternFlowVxlanReserved1().msg()
	}
	if obj.reserved1Holder == nil {
		obj.reserved1Holder = &patternFlowVxlanReserved1{obj: obj.obj.Reserved1}
	}
	return obj.reserved1Holder
}

// description is TBD
// Reserved1 returns a PatternFlowVxlanReserved1
func (obj *flowVxlan) HasReserved1() bool {
	return obj.obj.Reserved1 != nil
}

// description is TBD
// SetReserved1 sets the PatternFlowVxlanReserved1 value in the FlowVxlan object
func (obj *flowVxlan) SetReserved1(value PatternFlowVxlanReserved1) FlowVxlan {

	obj.reserved1Holder = nil
	obj.obj.Reserved1 = value.msg()

	return obj
}

func (obj *flowVxlan) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(vObj, set_default)
	}

	if obj.obj.Reserved0 != nil {

		obj.Reserved0().validateObj(vObj, set_default)
	}

	if obj.obj.Vni != nil {

		obj.Vni().validateObj(vObj, set_default)
	}

	if obj.obj.Reserved1 != nil {

		obj.Reserved1().validateObj(vObj, set_default)
	}

}

func (obj *flowVxlan) setDefault() {

}
