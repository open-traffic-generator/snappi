package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathExplicitRouteType1Ipv4Prefix *****
type flowRSVPPathExplicitRouteType1Ipv4Prefix struct {
	validation
	obj               *otg.FlowRSVPPathExplicitRouteType1Ipv4Prefix
	marshaller        marshalFlowRSVPPathExplicitRouteType1Ipv4Prefix
	unMarshaller      unMarshalFlowRSVPPathExplicitRouteType1Ipv4Prefix
	lBitHolder        PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
	lengthHolder      FlowRSVPExplicitRouteLength
	ipv4AddressHolder PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
}

func NewFlowRSVPPathExplicitRouteType1Ipv4Prefix() FlowRSVPPathExplicitRouteType1Ipv4Prefix {
	obj := flowRSVPPathExplicitRouteType1Ipv4Prefix{obj: &otg.FlowRSVPPathExplicitRouteType1Ipv4Prefix{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathExplicitRouteType1Ipv4Prefix) msg() *otg.FlowRSVPPathExplicitRouteType1Ipv4Prefix {
	return obj.obj
}

func (obj *flowRSVPPathExplicitRouteType1Ipv4Prefix) setMsg(msg *otg.FlowRSVPPathExplicitRouteType1Ipv4Prefix) FlowRSVPPathExplicitRouteType1Ipv4Prefix {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathExplicitRouteType1Ipv4Prefix struct {
	obj *flowRSVPPathExplicitRouteType1Ipv4Prefix
}

type marshalFlowRSVPPathExplicitRouteType1Ipv4Prefix interface {
	// ToProto marshals FlowRSVPPathExplicitRouteType1Ipv4Prefix to protobuf object *otg.FlowRSVPPathExplicitRouteType1Ipv4Prefix
	ToProto() (*otg.FlowRSVPPathExplicitRouteType1Ipv4Prefix, error)
	// ToPbText marshals FlowRSVPPathExplicitRouteType1Ipv4Prefix to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathExplicitRouteType1Ipv4Prefix to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathExplicitRouteType1Ipv4Prefix to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowRSVPPathExplicitRouteType1Ipv4Prefix to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowRSVPPathExplicitRouteType1Ipv4Prefix struct {
	obj *flowRSVPPathExplicitRouteType1Ipv4Prefix
}

type unMarshalFlowRSVPPathExplicitRouteType1Ipv4Prefix interface {
	// FromProto unmarshals FlowRSVPPathExplicitRouteType1Ipv4Prefix from protobuf object *otg.FlowRSVPPathExplicitRouteType1Ipv4Prefix
	FromProto(msg *otg.FlowRSVPPathExplicitRouteType1Ipv4Prefix) (FlowRSVPPathExplicitRouteType1Ipv4Prefix, error)
	// FromPbText unmarshals FlowRSVPPathExplicitRouteType1Ipv4Prefix from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathExplicitRouteType1Ipv4Prefix from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathExplicitRouteType1Ipv4Prefix from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathExplicitRouteType1Ipv4Prefix) Marshal() marshalFlowRSVPPathExplicitRouteType1Ipv4Prefix {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathExplicitRouteType1Ipv4Prefix{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathExplicitRouteType1Ipv4Prefix) Unmarshal() unMarshalFlowRSVPPathExplicitRouteType1Ipv4Prefix {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathExplicitRouteType1Ipv4Prefix{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathExplicitRouteType1Ipv4Prefix) ToProto() (*otg.FlowRSVPPathExplicitRouteType1Ipv4Prefix, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathExplicitRouteType1Ipv4Prefix) FromProto(msg *otg.FlowRSVPPathExplicitRouteType1Ipv4Prefix) (FlowRSVPPathExplicitRouteType1Ipv4Prefix, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathExplicitRouteType1Ipv4Prefix) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathExplicitRouteType1Ipv4Prefix) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathExplicitRouteType1Ipv4Prefix) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathExplicitRouteType1Ipv4Prefix) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathExplicitRouteType1Ipv4Prefix) ToJsonRaw() (string, error) {
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

func (m *marshalflowRSVPPathExplicitRouteType1Ipv4Prefix) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathExplicitRouteType1Ipv4Prefix) FromJson(value string) error {
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

func (obj *flowRSVPPathExplicitRouteType1Ipv4Prefix) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathExplicitRouteType1Ipv4Prefix) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathExplicitRouteType1Ipv4Prefix) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathExplicitRouteType1Ipv4Prefix) Clone() (FlowRSVPPathExplicitRouteType1Ipv4Prefix, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathExplicitRouteType1Ipv4Prefix()
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

func (obj *flowRSVPPathExplicitRouteType1Ipv4Prefix) setNil() {
	obj.lBitHolder = nil
	obj.lengthHolder = nil
	obj.ipv4AddressHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathExplicitRouteType1Ipv4Prefix is class = EXPLICIT_ROUTE, Type1 ROUTE_RECORD C-Type = 1 Subobject: IPv4 Prefix, C-Type: 1
type FlowRSVPPathExplicitRouteType1Ipv4Prefix interface {
	Validation
	// msg marshals FlowRSVPPathExplicitRouteType1Ipv4Prefix to protobuf object *otg.FlowRSVPPathExplicitRouteType1Ipv4Prefix
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathExplicitRouteType1Ipv4Prefix
	// setMsg unmarshals FlowRSVPPathExplicitRouteType1Ipv4Prefix from protobuf object *otg.FlowRSVPPathExplicitRouteType1Ipv4Prefix
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathExplicitRouteType1Ipv4Prefix) FlowRSVPPathExplicitRouteType1Ipv4Prefix
	// provides marshal interface
	Marshal() marshalFlowRSVPPathExplicitRouteType1Ipv4Prefix
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathExplicitRouteType1Ipv4Prefix
	// validate validates FlowRSVPPathExplicitRouteType1Ipv4Prefix
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathExplicitRouteType1Ipv4Prefix, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LBit returns PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit, set in FlowRSVPPathExplicitRouteType1Ipv4Prefix.
	// PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit is the L bit is an attribute of the subobject. The L bit is set if the subobject represents a loose hop in the explicit route. If the bit is not set, the subobject represents a strict hop in the explicit route.
	LBit() PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
	// SetLBit assigns PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit provided by user to FlowRSVPPathExplicitRouteType1Ipv4Prefix.
	// PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit is the L bit is an attribute of the subobject. The L bit is set if the subobject represents a loose hop in the explicit route. If the bit is not set, the subobject represents a strict hop in the explicit route.
	SetLBit(value PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) FlowRSVPPathExplicitRouteType1Ipv4Prefix
	// HasLBit checks if LBit has been set in FlowRSVPPathExplicitRouteType1Ipv4Prefix
	HasLBit() bool
	// Length returns FlowRSVPExplicitRouteLength, set in FlowRSVPPathExplicitRouteType1Ipv4Prefix.
	// FlowRSVPExplicitRouteLength is description is TBD
	Length() FlowRSVPExplicitRouteLength
	// SetLength assigns FlowRSVPExplicitRouteLength provided by user to FlowRSVPPathExplicitRouteType1Ipv4Prefix.
	// FlowRSVPExplicitRouteLength is description is TBD
	SetLength(value FlowRSVPExplicitRouteLength) FlowRSVPPathExplicitRouteType1Ipv4Prefix
	// HasLength checks if Length has been set in FlowRSVPPathExplicitRouteType1Ipv4Prefix
	HasLength() bool
	// Ipv4Address returns PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address, set in FlowRSVPPathExplicitRouteType1Ipv4Prefix.
	// PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address is this IPv4 address is treated as a prefix based on the prefix length value below.  Bits beyond the prefix are ignored on receipt and SHOULD be set to zero on transmission.
	Ipv4Address() PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
	// SetIpv4Address assigns PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address provided by user to FlowRSVPPathExplicitRouteType1Ipv4Prefix.
	// PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address is this IPv4 address is treated as a prefix based on the prefix length value below.  Bits beyond the prefix are ignored on receipt and SHOULD be set to zero on transmission.
	SetIpv4Address(value PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) FlowRSVPPathExplicitRouteType1Ipv4Prefix
	// HasIpv4Address checks if Ipv4Address has been set in FlowRSVPPathExplicitRouteType1Ipv4Prefix
	HasIpv4Address() bool
	// Prefix returns uint32, set in FlowRSVPPathExplicitRouteType1Ipv4Prefix.
	Prefix() uint32
	// SetPrefix assigns uint32 provided by user to FlowRSVPPathExplicitRouteType1Ipv4Prefix
	SetPrefix(value uint32) FlowRSVPPathExplicitRouteType1Ipv4Prefix
	// HasPrefix checks if Prefix has been set in FlowRSVPPathExplicitRouteType1Ipv4Prefix
	HasPrefix() bool
	setNil()
}

// description is TBD
// LBit returns a PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
func (obj *flowRSVPPathExplicitRouteType1Ipv4Prefix) LBit() PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit {
	if obj.obj.LBit == nil {
		obj.obj.LBit = NewPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit().msg()
	}
	if obj.lBitHolder == nil {
		obj.lBitHolder = &patternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit{obj: obj.obj.LBit}
	}
	return obj.lBitHolder
}

// description is TBD
// LBit returns a PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit
func (obj *flowRSVPPathExplicitRouteType1Ipv4Prefix) HasLBit() bool {
	return obj.obj.LBit != nil
}

// description is TBD
// SetLBit sets the PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit value in the FlowRSVPPathExplicitRouteType1Ipv4Prefix object
func (obj *flowRSVPPathExplicitRouteType1Ipv4Prefix) SetLBit(value PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixLBit) FlowRSVPPathExplicitRouteType1Ipv4Prefix {

	obj.lBitHolder = nil
	obj.obj.LBit = value.msg()

	return obj
}

// The Length contains the total length of the subobject in bytes,including L,Type and Length fields.   The Length MUST be atleast 4, and MUST be a multiple of 4.
// Length returns a FlowRSVPExplicitRouteLength
func (obj *flowRSVPPathExplicitRouteType1Ipv4Prefix) Length() FlowRSVPExplicitRouteLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewFlowRSVPExplicitRouteLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &flowRSVPExplicitRouteLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// The Length contains the total length of the subobject in bytes,including L,Type and Length fields.   The Length MUST be atleast 4, and MUST be a multiple of 4.
// Length returns a FlowRSVPExplicitRouteLength
func (obj *flowRSVPPathExplicitRouteType1Ipv4Prefix) HasLength() bool {
	return obj.obj.Length != nil
}

// The Length contains the total length of the subobject in bytes,including L,Type and Length fields.   The Length MUST be atleast 4, and MUST be a multiple of 4.
// SetLength sets the FlowRSVPExplicitRouteLength value in the FlowRSVPPathExplicitRouteType1Ipv4Prefix object
func (obj *flowRSVPPathExplicitRouteType1Ipv4Prefix) SetLength(value FlowRSVPExplicitRouteLength) FlowRSVPPathExplicitRouteType1Ipv4Prefix {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// description is TBD
// Ipv4Address returns a PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
func (obj *flowRSVPPathExplicitRouteType1Ipv4Prefix) Ipv4Address() PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address {
	if obj.obj.Ipv4Address == nil {
		obj.obj.Ipv4Address = NewPatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address().msg()
	}
	if obj.ipv4AddressHolder == nil {
		obj.ipv4AddressHolder = &patternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address{obj: obj.obj.Ipv4Address}
	}
	return obj.ipv4AddressHolder
}

// description is TBD
// Ipv4Address returns a PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address
func (obj *flowRSVPPathExplicitRouteType1Ipv4Prefix) HasIpv4Address() bool {
	return obj.obj.Ipv4Address != nil
}

// description is TBD
// SetIpv4Address sets the PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address value in the FlowRSVPPathExplicitRouteType1Ipv4Prefix object
func (obj *flowRSVPPathExplicitRouteType1Ipv4Prefix) SetIpv4Address(value PatternFlowRSVPPathExplicitRouteType1Ipv4PrefixIpv4Address) FlowRSVPPathExplicitRouteType1Ipv4Prefix {

	obj.ipv4AddressHolder = nil
	obj.obj.Ipv4Address = value.msg()

	return obj
}

// The prefix length of the IPv4 address.
// Prefix returns a uint32
func (obj *flowRSVPPathExplicitRouteType1Ipv4Prefix) Prefix() uint32 {

	return *obj.obj.Prefix

}

// The prefix length of the IPv4 address.
// Prefix returns a uint32
func (obj *flowRSVPPathExplicitRouteType1Ipv4Prefix) HasPrefix() bool {
	return obj.obj.Prefix != nil
}

// The prefix length of the IPv4 address.
// SetPrefix sets the uint32 value in the FlowRSVPPathExplicitRouteType1Ipv4Prefix object
func (obj *flowRSVPPathExplicitRouteType1Ipv4Prefix) SetPrefix(value uint32) FlowRSVPPathExplicitRouteType1Ipv4Prefix {

	obj.obj.Prefix = &value
	return obj
}

func (obj *flowRSVPPathExplicitRouteType1Ipv4Prefix) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.LBit != nil {

		obj.LBit().validateObj(vObj, set_default)
	}

	if obj.obj.Length != nil {

		obj.Length().validateObj(vObj, set_default)
	}

	if obj.obj.Ipv4Address != nil {

		obj.Ipv4Address().validateObj(vObj, set_default)
	}

	if obj.obj.Prefix != nil {

		if *obj.obj.Prefix < 1 || *obj.obj.Prefix > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= FlowRSVPPathExplicitRouteType1Ipv4Prefix.Prefix <= 32 but Got %d", *obj.obj.Prefix))
		}

	}

}

func (obj *flowRSVPPathExplicitRouteType1Ipv4Prefix) setDefault() {
	if obj.obj.Prefix == nil {
		obj.SetPrefix(32)
	}

}
