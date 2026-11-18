package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIcmpv6NdpMtuOption *****
type flowIcmpv6NdpMtuOption struct {
	validation
	obj            *otg.FlowIcmpv6NdpMtuOption
	marshaller     marshalFlowIcmpv6NdpMtuOption
	unMarshaller   unMarshalFlowIcmpv6NdpMtuOption
	lengthHolder   PatternFlowIcmpv6NdpMtuOptionLength
	reservedHolder PatternFlowIcmpv6NdpMtuOptionReserved
	mtuHolder      PatternFlowIcmpv6NdpMtuOptionMtu
}

func NewFlowIcmpv6NdpMtuOption() FlowIcmpv6NdpMtuOption {
	obj := flowIcmpv6NdpMtuOption{obj: &otg.FlowIcmpv6NdpMtuOption{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIcmpv6NdpMtuOption) msg() *otg.FlowIcmpv6NdpMtuOption {
	return obj.obj
}

func (obj *flowIcmpv6NdpMtuOption) setMsg(msg *otg.FlowIcmpv6NdpMtuOption) FlowIcmpv6NdpMtuOption {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIcmpv6NdpMtuOption struct {
	obj *flowIcmpv6NdpMtuOption
}

type marshalFlowIcmpv6NdpMtuOption interface {
	// ToProto marshals FlowIcmpv6NdpMtuOption to protobuf object *otg.FlowIcmpv6NdpMtuOption
	ToProto() (*otg.FlowIcmpv6NdpMtuOption, error)
	// ToPbText marshals FlowIcmpv6NdpMtuOption to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIcmpv6NdpMtuOption to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIcmpv6NdpMtuOption to JSON text
	ToJson() (string, error)
}

type unMarshalflowIcmpv6NdpMtuOption struct {
	obj *flowIcmpv6NdpMtuOption
}

type unMarshalFlowIcmpv6NdpMtuOption interface {
	// FromProto unmarshals FlowIcmpv6NdpMtuOption from protobuf object *otg.FlowIcmpv6NdpMtuOption
	FromProto(msg *otg.FlowIcmpv6NdpMtuOption) (FlowIcmpv6NdpMtuOption, error)
	// FromPbText unmarshals FlowIcmpv6NdpMtuOption from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIcmpv6NdpMtuOption from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIcmpv6NdpMtuOption from JSON text
	FromJson(value string) error
}

func (obj *flowIcmpv6NdpMtuOption) Marshal() marshalFlowIcmpv6NdpMtuOption {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIcmpv6NdpMtuOption{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIcmpv6NdpMtuOption) Unmarshal() unMarshalFlowIcmpv6NdpMtuOption {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIcmpv6NdpMtuOption{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIcmpv6NdpMtuOption) ToProto() (*otg.FlowIcmpv6NdpMtuOption, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIcmpv6NdpMtuOption) FromProto(msg *otg.FlowIcmpv6NdpMtuOption) (FlowIcmpv6NdpMtuOption, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIcmpv6NdpMtuOption) ToPbText() (string, error) {
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

func (m *unMarshalflowIcmpv6NdpMtuOption) FromPbText(value string) error {
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

func (m *marshalflowIcmpv6NdpMtuOption) ToYaml() (string, error) {
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

func (m *unMarshalflowIcmpv6NdpMtuOption) FromYaml(value string) error {
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

func (m *marshalflowIcmpv6NdpMtuOption) ToJson() (string, error) {
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

func (m *unMarshalflowIcmpv6NdpMtuOption) FromJson(value string) error {
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

func (obj *flowIcmpv6NdpMtuOption) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIcmpv6NdpMtuOption) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIcmpv6NdpMtuOption) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIcmpv6NdpMtuOption) Clone() (FlowIcmpv6NdpMtuOption, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIcmpv6NdpMtuOption()
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

func (obj *flowIcmpv6NdpMtuOption) setNil() {
	obj.lengthHolder = nil
	obj.reservedHolder = nil
	obj.mtuHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIcmpv6NdpMtuOption is nDP MTU option (Type 5). Defined in RFC 4861 Section 4.6.4. Used in Router Advertisement messages to ensure that all nodes on a link use the same MTU value in those cases where the link MTU is not well known.
type FlowIcmpv6NdpMtuOption interface {
	Validation
	// msg marshals FlowIcmpv6NdpMtuOption to protobuf object *otg.FlowIcmpv6NdpMtuOption
	// and doesn't set defaults
	msg() *otg.FlowIcmpv6NdpMtuOption
	// setMsg unmarshals FlowIcmpv6NdpMtuOption from protobuf object *otg.FlowIcmpv6NdpMtuOption
	// and doesn't set defaults
	setMsg(*otg.FlowIcmpv6NdpMtuOption) FlowIcmpv6NdpMtuOption
	// provides marshal interface
	Marshal() marshalFlowIcmpv6NdpMtuOption
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIcmpv6NdpMtuOption
	// validate validates FlowIcmpv6NdpMtuOption
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIcmpv6NdpMtuOption, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Length returns PatternFlowIcmpv6NdpMtuOptionLength, set in FlowIcmpv6NdpMtuOption.
	// PatternFlowIcmpv6NdpMtuOptionLength is option length in units of 8 octets. Fixed at 1 for the MTU option.
	Length() PatternFlowIcmpv6NdpMtuOptionLength
	// SetLength assigns PatternFlowIcmpv6NdpMtuOptionLength provided by user to FlowIcmpv6NdpMtuOption.
	// PatternFlowIcmpv6NdpMtuOptionLength is option length in units of 8 octets. Fixed at 1 for the MTU option.
	SetLength(value PatternFlowIcmpv6NdpMtuOptionLength) FlowIcmpv6NdpMtuOption
	// HasLength checks if Length has been set in FlowIcmpv6NdpMtuOption
	HasLength() bool
	// Reserved returns PatternFlowIcmpv6NdpMtuOptionReserved, set in FlowIcmpv6NdpMtuOption.
	// PatternFlowIcmpv6NdpMtuOptionReserved is reserved. Must be initialized to zero by the sender.
	Reserved() PatternFlowIcmpv6NdpMtuOptionReserved
	// SetReserved assigns PatternFlowIcmpv6NdpMtuOptionReserved provided by user to FlowIcmpv6NdpMtuOption.
	// PatternFlowIcmpv6NdpMtuOptionReserved is reserved. Must be initialized to zero by the sender.
	SetReserved(value PatternFlowIcmpv6NdpMtuOptionReserved) FlowIcmpv6NdpMtuOption
	// HasReserved checks if Reserved has been set in FlowIcmpv6NdpMtuOption
	HasReserved() bool
	// Mtu returns PatternFlowIcmpv6NdpMtuOptionMtu, set in FlowIcmpv6NdpMtuOption.
	// PatternFlowIcmpv6NdpMtuOptionMtu is recommended MTU for the link, in octets.
	Mtu() PatternFlowIcmpv6NdpMtuOptionMtu
	// SetMtu assigns PatternFlowIcmpv6NdpMtuOptionMtu provided by user to FlowIcmpv6NdpMtuOption.
	// PatternFlowIcmpv6NdpMtuOptionMtu is recommended MTU for the link, in octets.
	SetMtu(value PatternFlowIcmpv6NdpMtuOptionMtu) FlowIcmpv6NdpMtuOption
	// HasMtu checks if Mtu has been set in FlowIcmpv6NdpMtuOption
	HasMtu() bool
	setNil()
}

// description is TBD
// Length returns a PatternFlowIcmpv6NdpMtuOptionLength
func (obj *flowIcmpv6NdpMtuOption) Length() PatternFlowIcmpv6NdpMtuOptionLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewPatternFlowIcmpv6NdpMtuOptionLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &patternFlowIcmpv6NdpMtuOptionLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// description is TBD
// Length returns a PatternFlowIcmpv6NdpMtuOptionLength
func (obj *flowIcmpv6NdpMtuOption) HasLength() bool {
	return obj.obj.Length != nil
}

// description is TBD
// SetLength sets the PatternFlowIcmpv6NdpMtuOptionLength value in the FlowIcmpv6NdpMtuOption object
func (obj *flowIcmpv6NdpMtuOption) SetLength(value PatternFlowIcmpv6NdpMtuOptionLength) FlowIcmpv6NdpMtuOption {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// description is TBD
// Reserved returns a PatternFlowIcmpv6NdpMtuOptionReserved
func (obj *flowIcmpv6NdpMtuOption) Reserved() PatternFlowIcmpv6NdpMtuOptionReserved {
	if obj.obj.Reserved == nil {
		obj.obj.Reserved = NewPatternFlowIcmpv6NdpMtuOptionReserved().msg()
	}
	if obj.reservedHolder == nil {
		obj.reservedHolder = &patternFlowIcmpv6NdpMtuOptionReserved{obj: obj.obj.Reserved}
	}
	return obj.reservedHolder
}

// description is TBD
// Reserved returns a PatternFlowIcmpv6NdpMtuOptionReserved
func (obj *flowIcmpv6NdpMtuOption) HasReserved() bool {
	return obj.obj.Reserved != nil
}

// description is TBD
// SetReserved sets the PatternFlowIcmpv6NdpMtuOptionReserved value in the FlowIcmpv6NdpMtuOption object
func (obj *flowIcmpv6NdpMtuOption) SetReserved(value PatternFlowIcmpv6NdpMtuOptionReserved) FlowIcmpv6NdpMtuOption {

	obj.reservedHolder = nil
	obj.obj.Reserved = value.msg()

	return obj
}

// description is TBD
// Mtu returns a PatternFlowIcmpv6NdpMtuOptionMtu
func (obj *flowIcmpv6NdpMtuOption) Mtu() PatternFlowIcmpv6NdpMtuOptionMtu {
	if obj.obj.Mtu == nil {
		obj.obj.Mtu = NewPatternFlowIcmpv6NdpMtuOptionMtu().msg()
	}
	if obj.mtuHolder == nil {
		obj.mtuHolder = &patternFlowIcmpv6NdpMtuOptionMtu{obj: obj.obj.Mtu}
	}
	return obj.mtuHolder
}

// description is TBD
// Mtu returns a PatternFlowIcmpv6NdpMtuOptionMtu
func (obj *flowIcmpv6NdpMtuOption) HasMtu() bool {
	return obj.obj.Mtu != nil
}

// description is TBD
// SetMtu sets the PatternFlowIcmpv6NdpMtuOptionMtu value in the FlowIcmpv6NdpMtuOption object
func (obj *flowIcmpv6NdpMtuOption) SetMtu(value PatternFlowIcmpv6NdpMtuOptionMtu) FlowIcmpv6NdpMtuOption {

	obj.mtuHolder = nil
	obj.obj.Mtu = value.msg()

	return obj
}

func (obj *flowIcmpv6NdpMtuOption) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Length != nil {

		obj.Length().validateObj(vObj, set_default)
	}

	if obj.obj.Reserved != nil {

		obj.Reserved().validateObj(vObj, set_default)
	}

	if obj.obj.Mtu != nil {

		obj.Mtu().validateObj(vObj, set_default)
	}

}

func (obj *flowIcmpv6NdpMtuOption) setDefault() {

}
