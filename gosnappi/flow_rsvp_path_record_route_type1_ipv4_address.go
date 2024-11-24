package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathRecordRouteType1Ipv4Address *****
type flowRSVPPathRecordRouteType1Ipv4Address struct {
	validation
	obj                *otg.FlowRSVPPathRecordRouteType1Ipv4Address
	marshaller         marshalFlowRSVPPathRecordRouteType1Ipv4Address
	unMarshaller       unMarshalFlowRSVPPathRecordRouteType1Ipv4Address
	lengthHolder       FlowRSVPRouteRecordLength
	ipv4AddressHolder  PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
	prefixLengthHolder PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
	flagsHolder        FlowRSVPRecordRouteIPv4Flag
}

func NewFlowRSVPPathRecordRouteType1Ipv4Address() FlowRSVPPathRecordRouteType1Ipv4Address {
	obj := flowRSVPPathRecordRouteType1Ipv4Address{obj: &otg.FlowRSVPPathRecordRouteType1Ipv4Address{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathRecordRouteType1Ipv4Address) msg() *otg.FlowRSVPPathRecordRouteType1Ipv4Address {
	return obj.obj
}

func (obj *flowRSVPPathRecordRouteType1Ipv4Address) setMsg(msg *otg.FlowRSVPPathRecordRouteType1Ipv4Address) FlowRSVPPathRecordRouteType1Ipv4Address {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathRecordRouteType1Ipv4Address struct {
	obj *flowRSVPPathRecordRouteType1Ipv4Address
}

type marshalFlowRSVPPathRecordRouteType1Ipv4Address interface {
	// ToProto marshals FlowRSVPPathRecordRouteType1Ipv4Address to protobuf object *otg.FlowRSVPPathRecordRouteType1Ipv4Address
	ToProto() (*otg.FlowRSVPPathRecordRouteType1Ipv4Address, error)
	// ToPbText marshals FlowRSVPPathRecordRouteType1Ipv4Address to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathRecordRouteType1Ipv4Address to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathRecordRouteType1Ipv4Address to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowRSVPPathRecordRouteType1Ipv4Address to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowRSVPPathRecordRouteType1Ipv4Address struct {
	obj *flowRSVPPathRecordRouteType1Ipv4Address
}

type unMarshalFlowRSVPPathRecordRouteType1Ipv4Address interface {
	// FromProto unmarshals FlowRSVPPathRecordRouteType1Ipv4Address from protobuf object *otg.FlowRSVPPathRecordRouteType1Ipv4Address
	FromProto(msg *otg.FlowRSVPPathRecordRouteType1Ipv4Address) (FlowRSVPPathRecordRouteType1Ipv4Address, error)
	// FromPbText unmarshals FlowRSVPPathRecordRouteType1Ipv4Address from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathRecordRouteType1Ipv4Address from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathRecordRouteType1Ipv4Address from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathRecordRouteType1Ipv4Address) Marshal() marshalFlowRSVPPathRecordRouteType1Ipv4Address {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathRecordRouteType1Ipv4Address{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathRecordRouteType1Ipv4Address) Unmarshal() unMarshalFlowRSVPPathRecordRouteType1Ipv4Address {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathRecordRouteType1Ipv4Address{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathRecordRouteType1Ipv4Address) ToProto() (*otg.FlowRSVPPathRecordRouteType1Ipv4Address, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathRecordRouteType1Ipv4Address) FromProto(msg *otg.FlowRSVPPathRecordRouteType1Ipv4Address) (FlowRSVPPathRecordRouteType1Ipv4Address, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathRecordRouteType1Ipv4Address) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathRecordRouteType1Ipv4Address) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathRecordRouteType1Ipv4Address) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathRecordRouteType1Ipv4Address) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathRecordRouteType1Ipv4Address) ToJsonRaw() (string, error) {
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

func (m *marshalflowRSVPPathRecordRouteType1Ipv4Address) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathRecordRouteType1Ipv4Address) FromJson(value string) error {
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

func (obj *flowRSVPPathRecordRouteType1Ipv4Address) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathRecordRouteType1Ipv4Address) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathRecordRouteType1Ipv4Address) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathRecordRouteType1Ipv4Address) Clone() (FlowRSVPPathRecordRouteType1Ipv4Address, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathRecordRouteType1Ipv4Address()
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

func (obj *flowRSVPPathRecordRouteType1Ipv4Address) setNil() {
	obj.lengthHolder = nil
	obj.ipv4AddressHolder = nil
	obj.prefixLengthHolder = nil
	obj.flagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathRecordRouteType1Ipv4Address is class = RECORD_ROUTE, Type1 ROUTE_RECORD C-Type = 1 Subobject: IPv4 Address, C-Type: 1
type FlowRSVPPathRecordRouteType1Ipv4Address interface {
	Validation
	// msg marshals FlowRSVPPathRecordRouteType1Ipv4Address to protobuf object *otg.FlowRSVPPathRecordRouteType1Ipv4Address
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathRecordRouteType1Ipv4Address
	// setMsg unmarshals FlowRSVPPathRecordRouteType1Ipv4Address from protobuf object *otg.FlowRSVPPathRecordRouteType1Ipv4Address
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathRecordRouteType1Ipv4Address) FlowRSVPPathRecordRouteType1Ipv4Address
	// provides marshal interface
	Marshal() marshalFlowRSVPPathRecordRouteType1Ipv4Address
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathRecordRouteType1Ipv4Address
	// validate validates FlowRSVPPathRecordRouteType1Ipv4Address
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathRecordRouteType1Ipv4Address, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Length returns FlowRSVPRouteRecordLength, set in FlowRSVPPathRecordRouteType1Ipv4Address.
	// FlowRSVPRouteRecordLength is description is TBD
	Length() FlowRSVPRouteRecordLength
	// SetLength assigns FlowRSVPRouteRecordLength provided by user to FlowRSVPPathRecordRouteType1Ipv4Address.
	// FlowRSVPRouteRecordLength is description is TBD
	SetLength(value FlowRSVPRouteRecordLength) FlowRSVPPathRecordRouteType1Ipv4Address
	// HasLength checks if Length has been set in FlowRSVPPathRecordRouteType1Ipv4Address
	HasLength() bool
	// Ipv4Address returns PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address, set in FlowRSVPPathRecordRouteType1Ipv4Address.
	// PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address is a 32-bit unicast, host address.  Any network-reachable interface address is allowed here. Illegal addresses, such as certain loopback addresses, SHOULD NOT be used.
	Ipv4Address() PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
	// SetIpv4Address assigns PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address provided by user to FlowRSVPPathRecordRouteType1Ipv4Address.
	// PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address is a 32-bit unicast, host address.  Any network-reachable interface address is allowed here. Illegal addresses, such as certain loopback addresses, SHOULD NOT be used.
	SetIpv4Address(value PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) FlowRSVPPathRecordRouteType1Ipv4Address
	// HasIpv4Address checks if Ipv4Address has been set in FlowRSVPPathRecordRouteType1Ipv4Address
	HasIpv4Address() bool
	// PrefixLength returns PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength, set in FlowRSVPPathRecordRouteType1Ipv4Address.
	// PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength is prefix-length of IPv4 address.
	PrefixLength() PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
	// SetPrefixLength assigns PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength provided by user to FlowRSVPPathRecordRouteType1Ipv4Address.
	// PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength is prefix-length of IPv4 address.
	SetPrefixLength(value PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) FlowRSVPPathRecordRouteType1Ipv4Address
	// HasPrefixLength checks if PrefixLength has been set in FlowRSVPPathRecordRouteType1Ipv4Address
	HasPrefixLength() bool
	// Flags returns FlowRSVPRecordRouteIPv4Flag, set in FlowRSVPPathRecordRouteType1Ipv4Address.
	// FlowRSVPRecordRouteIPv4Flag is description is TBD
	Flags() FlowRSVPRecordRouteIPv4Flag
	// SetFlags assigns FlowRSVPRecordRouteIPv4Flag provided by user to FlowRSVPPathRecordRouteType1Ipv4Address.
	// FlowRSVPRecordRouteIPv4Flag is description is TBD
	SetFlags(value FlowRSVPRecordRouteIPv4Flag) FlowRSVPPathRecordRouteType1Ipv4Address
	// HasFlags checks if Flags has been set in FlowRSVPPathRecordRouteType1Ipv4Address
	HasFlags() bool
	setNil()
}

// The Length contains the total length of the subobject in bytes, including the Type and Length fields.   The Length MUST be atleast 4, and MUST be a multiple of 4.
// Length returns a FlowRSVPRouteRecordLength
func (obj *flowRSVPPathRecordRouteType1Ipv4Address) Length() FlowRSVPRouteRecordLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewFlowRSVPRouteRecordLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &flowRSVPRouteRecordLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// The Length contains the total length of the subobject in bytes, including the Type and Length fields.   The Length MUST be atleast 4, and MUST be a multiple of 4.
// Length returns a FlowRSVPRouteRecordLength
func (obj *flowRSVPPathRecordRouteType1Ipv4Address) HasLength() bool {
	return obj.obj.Length != nil
}

// The Length contains the total length of the subobject in bytes, including the Type and Length fields.   The Length MUST be atleast 4, and MUST be a multiple of 4.
// SetLength sets the FlowRSVPRouteRecordLength value in the FlowRSVPPathRecordRouteType1Ipv4Address object
func (obj *flowRSVPPathRecordRouteType1Ipv4Address) SetLength(value FlowRSVPRouteRecordLength) FlowRSVPPathRecordRouteType1Ipv4Address {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// description is TBD
// Ipv4Address returns a PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
func (obj *flowRSVPPathRecordRouteType1Ipv4Address) Ipv4Address() PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address {
	if obj.obj.Ipv4Address == nil {
		obj.obj.Ipv4Address = NewPatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address().msg()
	}
	if obj.ipv4AddressHolder == nil {
		obj.ipv4AddressHolder = &patternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address{obj: obj.obj.Ipv4Address}
	}
	return obj.ipv4AddressHolder
}

// description is TBD
// Ipv4Address returns a PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address
func (obj *flowRSVPPathRecordRouteType1Ipv4Address) HasIpv4Address() bool {
	return obj.obj.Ipv4Address != nil
}

// description is TBD
// SetIpv4Address sets the PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address value in the FlowRSVPPathRecordRouteType1Ipv4Address object
func (obj *flowRSVPPathRecordRouteType1Ipv4Address) SetIpv4Address(value PatternFlowRSVPPathRecordRouteType1Ipv4AddressIpv4Address) FlowRSVPPathRecordRouteType1Ipv4Address {

	obj.ipv4AddressHolder = nil
	obj.obj.Ipv4Address = value.msg()

	return obj
}

// description is TBD
// PrefixLength returns a PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
func (obj *flowRSVPPathRecordRouteType1Ipv4Address) PrefixLength() PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength {
	if obj.obj.PrefixLength == nil {
		obj.obj.PrefixLength = NewPatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength().msg()
	}
	if obj.prefixLengthHolder == nil {
		obj.prefixLengthHolder = &patternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength{obj: obj.obj.PrefixLength}
	}
	return obj.prefixLengthHolder
}

// description is TBD
// PrefixLength returns a PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength
func (obj *flowRSVPPathRecordRouteType1Ipv4Address) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// description is TBD
// SetPrefixLength sets the PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength value in the FlowRSVPPathRecordRouteType1Ipv4Address object
func (obj *flowRSVPPathRecordRouteType1Ipv4Address) SetPrefixLength(value PatternFlowRSVPPathRecordRouteType1Ipv4AddressPrefixLength) FlowRSVPPathRecordRouteType1Ipv4Address {

	obj.prefixLengthHolder = nil
	obj.obj.PrefixLength = value.msg()

	return obj
}

// 0x01  local_protection_available, 0x02  local_protection_in_use
// Flags returns a FlowRSVPRecordRouteIPv4Flag
func (obj *flowRSVPPathRecordRouteType1Ipv4Address) Flags() FlowRSVPRecordRouteIPv4Flag {
	if obj.obj.Flags == nil {
		obj.obj.Flags = NewFlowRSVPRecordRouteIPv4Flag().msg()
	}
	if obj.flagsHolder == nil {
		obj.flagsHolder = &flowRSVPRecordRouteIPv4Flag{obj: obj.obj.Flags}
	}
	return obj.flagsHolder
}

// 0x01  local_protection_available, 0x02  local_protection_in_use
// Flags returns a FlowRSVPRecordRouteIPv4Flag
func (obj *flowRSVPPathRecordRouteType1Ipv4Address) HasFlags() bool {
	return obj.obj.Flags != nil
}

// 0x01  local_protection_available, 0x02  local_protection_in_use
// SetFlags sets the FlowRSVPRecordRouteIPv4Flag value in the FlowRSVPPathRecordRouteType1Ipv4Address object
func (obj *flowRSVPPathRecordRouteType1Ipv4Address) SetFlags(value FlowRSVPRecordRouteIPv4Flag) FlowRSVPPathRecordRouteType1Ipv4Address {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

func (obj *flowRSVPPathRecordRouteType1Ipv4Address) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Length != nil {

		obj.Length().validateObj(vObj, set_default)
	}

	if obj.obj.Ipv4Address != nil {

		obj.Ipv4Address().validateObj(vObj, set_default)
	}

	if obj.obj.PrefixLength != nil {

		obj.PrefixLength().validateObj(vObj, set_default)
	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(vObj, set_default)
	}

}

func (obj *flowRSVPPathRecordRouteType1Ipv4Address) setDefault() {

}
