package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIcmpv6NdpRedirectedHeaderOption *****
type flowIcmpv6NdpRedirectedHeaderOption struct {
	validation
	obj             *otg.FlowIcmpv6NdpRedirectedHeaderOption
	marshaller      marshalFlowIcmpv6NdpRedirectedHeaderOption
	unMarshaller    unMarshalFlowIcmpv6NdpRedirectedHeaderOption
	lengthHolder    PatternFlowIcmpv6NdpRedirectedHeaderOptionLength
	reserved1Holder PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1
	reserved2Holder PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
}

func NewFlowIcmpv6NdpRedirectedHeaderOption() FlowIcmpv6NdpRedirectedHeaderOption {
	obj := flowIcmpv6NdpRedirectedHeaderOption{obj: &otg.FlowIcmpv6NdpRedirectedHeaderOption{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIcmpv6NdpRedirectedHeaderOption) msg() *otg.FlowIcmpv6NdpRedirectedHeaderOption {
	return obj.obj
}

func (obj *flowIcmpv6NdpRedirectedHeaderOption) setMsg(msg *otg.FlowIcmpv6NdpRedirectedHeaderOption) FlowIcmpv6NdpRedirectedHeaderOption {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIcmpv6NdpRedirectedHeaderOption struct {
	obj *flowIcmpv6NdpRedirectedHeaderOption
}

type marshalFlowIcmpv6NdpRedirectedHeaderOption interface {
	// ToProto marshals FlowIcmpv6NdpRedirectedHeaderOption to protobuf object *otg.FlowIcmpv6NdpRedirectedHeaderOption
	ToProto() (*otg.FlowIcmpv6NdpRedirectedHeaderOption, error)
	// ToPbText marshals FlowIcmpv6NdpRedirectedHeaderOption to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIcmpv6NdpRedirectedHeaderOption to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIcmpv6NdpRedirectedHeaderOption to JSON text
	ToJson() (string, error)
}

type unMarshalflowIcmpv6NdpRedirectedHeaderOption struct {
	obj *flowIcmpv6NdpRedirectedHeaderOption
}

type unMarshalFlowIcmpv6NdpRedirectedHeaderOption interface {
	// FromProto unmarshals FlowIcmpv6NdpRedirectedHeaderOption from protobuf object *otg.FlowIcmpv6NdpRedirectedHeaderOption
	FromProto(msg *otg.FlowIcmpv6NdpRedirectedHeaderOption) (FlowIcmpv6NdpRedirectedHeaderOption, error)
	// FromPbText unmarshals FlowIcmpv6NdpRedirectedHeaderOption from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIcmpv6NdpRedirectedHeaderOption from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIcmpv6NdpRedirectedHeaderOption from JSON text
	FromJson(value string) error
}

func (obj *flowIcmpv6NdpRedirectedHeaderOption) Marshal() marshalFlowIcmpv6NdpRedirectedHeaderOption {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIcmpv6NdpRedirectedHeaderOption{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIcmpv6NdpRedirectedHeaderOption) Unmarshal() unMarshalFlowIcmpv6NdpRedirectedHeaderOption {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIcmpv6NdpRedirectedHeaderOption{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIcmpv6NdpRedirectedHeaderOption) ToProto() (*otg.FlowIcmpv6NdpRedirectedHeaderOption, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIcmpv6NdpRedirectedHeaderOption) FromProto(msg *otg.FlowIcmpv6NdpRedirectedHeaderOption) (FlowIcmpv6NdpRedirectedHeaderOption, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIcmpv6NdpRedirectedHeaderOption) ToPbText() (string, error) {
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

func (m *unMarshalflowIcmpv6NdpRedirectedHeaderOption) FromPbText(value string) error {
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

func (m *marshalflowIcmpv6NdpRedirectedHeaderOption) ToYaml() (string, error) {
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

func (m *unMarshalflowIcmpv6NdpRedirectedHeaderOption) FromYaml(value string) error {
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

func (m *marshalflowIcmpv6NdpRedirectedHeaderOption) ToJson() (string, error) {
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

func (m *unMarshalflowIcmpv6NdpRedirectedHeaderOption) FromJson(value string) error {
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

func (obj *flowIcmpv6NdpRedirectedHeaderOption) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIcmpv6NdpRedirectedHeaderOption) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIcmpv6NdpRedirectedHeaderOption) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIcmpv6NdpRedirectedHeaderOption) Clone() (FlowIcmpv6NdpRedirectedHeaderOption, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIcmpv6NdpRedirectedHeaderOption()
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

func (obj *flowIcmpv6NdpRedirectedHeaderOption) setNil() {
	obj.lengthHolder = nil
	obj.reserved1Holder = nil
	obj.reserved2Holder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIcmpv6NdpRedirectedHeaderOption is nDP Redirected Header option (Type 4). Defined in RFC 4861 Section 4.6.3. Used in Redirect messages to include as much as possible of the IP packet that triggered the Redirect, without exceeding the minimum IPv6 MTU.
type FlowIcmpv6NdpRedirectedHeaderOption interface {
	Validation
	// msg marshals FlowIcmpv6NdpRedirectedHeaderOption to protobuf object *otg.FlowIcmpv6NdpRedirectedHeaderOption
	// and doesn't set defaults
	msg() *otg.FlowIcmpv6NdpRedirectedHeaderOption
	// setMsg unmarshals FlowIcmpv6NdpRedirectedHeaderOption from protobuf object *otg.FlowIcmpv6NdpRedirectedHeaderOption
	// and doesn't set defaults
	setMsg(*otg.FlowIcmpv6NdpRedirectedHeaderOption) FlowIcmpv6NdpRedirectedHeaderOption
	// provides marshal interface
	Marshal() marshalFlowIcmpv6NdpRedirectedHeaderOption
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIcmpv6NdpRedirectedHeaderOption
	// validate validates FlowIcmpv6NdpRedirectedHeaderOption
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIcmpv6NdpRedirectedHeaderOption, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Length returns PatternFlowIcmpv6NdpRedirectedHeaderOptionLength, set in FlowIcmpv6NdpRedirectedHeaderOption.
	// PatternFlowIcmpv6NdpRedirectedHeaderOptionLength is option length in units of 8 octets, including type/length bytes, reserved bytes and the truncated original packet.
	Length() PatternFlowIcmpv6NdpRedirectedHeaderOptionLength
	// SetLength assigns PatternFlowIcmpv6NdpRedirectedHeaderOptionLength provided by user to FlowIcmpv6NdpRedirectedHeaderOption.
	// PatternFlowIcmpv6NdpRedirectedHeaderOptionLength is option length in units of 8 octets, including type/length bytes, reserved bytes and the truncated original packet.
	SetLength(value PatternFlowIcmpv6NdpRedirectedHeaderOptionLength) FlowIcmpv6NdpRedirectedHeaderOption
	// HasLength checks if Length has been set in FlowIcmpv6NdpRedirectedHeaderOption
	HasLength() bool
	// Reserved1 returns PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1, set in FlowIcmpv6NdpRedirectedHeaderOption.
	// PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1 is reserved. Must be initialized to zero by the sender.
	Reserved1() PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1
	// SetReserved1 assigns PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1 provided by user to FlowIcmpv6NdpRedirectedHeaderOption.
	// PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1 is reserved. Must be initialized to zero by the sender.
	SetReserved1(value PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1) FlowIcmpv6NdpRedirectedHeaderOption
	// HasReserved1 checks if Reserved1 has been set in FlowIcmpv6NdpRedirectedHeaderOption
	HasReserved1() bool
	// Reserved2 returns PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2, set in FlowIcmpv6NdpRedirectedHeaderOption.
	// PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 is reserved. Must be initialized to zero by the sender.
	Reserved2() PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
	// SetReserved2 assigns PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 provided by user to FlowIcmpv6NdpRedirectedHeaderOption.
	// PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 is reserved. Must be initialized to zero by the sender.
	SetReserved2(value PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) FlowIcmpv6NdpRedirectedHeaderOption
	// HasReserved2 checks if Reserved2 has been set in FlowIcmpv6NdpRedirectedHeaderOption
	HasReserved2() bool
	// Data returns string, set in FlowIcmpv6NdpRedirectedHeaderOption.
	Data() string
	// SetData assigns string provided by user to FlowIcmpv6NdpRedirectedHeaderOption
	SetData(value string) FlowIcmpv6NdpRedirectedHeaderOption
	// HasData checks if Data has been set in FlowIcmpv6NdpRedirectedHeaderOption
	HasData() bool
	setNil()
}

// description is TBD
// Length returns a PatternFlowIcmpv6NdpRedirectedHeaderOptionLength
func (obj *flowIcmpv6NdpRedirectedHeaderOption) Length() PatternFlowIcmpv6NdpRedirectedHeaderOptionLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewPatternFlowIcmpv6NdpRedirectedHeaderOptionLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &patternFlowIcmpv6NdpRedirectedHeaderOptionLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// description is TBD
// Length returns a PatternFlowIcmpv6NdpRedirectedHeaderOptionLength
func (obj *flowIcmpv6NdpRedirectedHeaderOption) HasLength() bool {
	return obj.obj.Length != nil
}

// description is TBD
// SetLength sets the PatternFlowIcmpv6NdpRedirectedHeaderOptionLength value in the FlowIcmpv6NdpRedirectedHeaderOption object
func (obj *flowIcmpv6NdpRedirectedHeaderOption) SetLength(value PatternFlowIcmpv6NdpRedirectedHeaderOptionLength) FlowIcmpv6NdpRedirectedHeaderOption {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// description is TBD
// Reserved1 returns a PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1
func (obj *flowIcmpv6NdpRedirectedHeaderOption) Reserved1() PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1 {
	if obj.obj.Reserved1 == nil {
		obj.obj.Reserved1 = NewPatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1().msg()
	}
	if obj.reserved1Holder == nil {
		obj.reserved1Holder = &patternFlowIcmpv6NdpRedirectedHeaderOptionReserved1{obj: obj.obj.Reserved1}
	}
	return obj.reserved1Holder
}

// description is TBD
// Reserved1 returns a PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1
func (obj *flowIcmpv6NdpRedirectedHeaderOption) HasReserved1() bool {
	return obj.obj.Reserved1 != nil
}

// description is TBD
// SetReserved1 sets the PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1 value in the FlowIcmpv6NdpRedirectedHeaderOption object
func (obj *flowIcmpv6NdpRedirectedHeaderOption) SetReserved1(value PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved1) FlowIcmpv6NdpRedirectedHeaderOption {

	obj.reserved1Holder = nil
	obj.obj.Reserved1 = value.msg()

	return obj
}

// description is TBD
// Reserved2 returns a PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
func (obj *flowIcmpv6NdpRedirectedHeaderOption) Reserved2() PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 {
	if obj.obj.Reserved2 == nil {
		obj.obj.Reserved2 = NewPatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2().msg()
	}
	if obj.reserved2Holder == nil {
		obj.reserved2Holder = &patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2{obj: obj.obj.Reserved2}
	}
	return obj.reserved2Holder
}

// description is TBD
// Reserved2 returns a PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
func (obj *flowIcmpv6NdpRedirectedHeaderOption) HasReserved2() bool {
	return obj.obj.Reserved2 != nil
}

// description is TBD
// SetReserved2 sets the PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 value in the FlowIcmpv6NdpRedirectedHeaderOption object
func (obj *flowIcmpv6NdpRedirectedHeaderOption) SetReserved2(value PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) FlowIcmpv6NdpRedirectedHeaderOption {

	obj.reserved2Holder = nil
	obj.obj.Reserved2 = value.msg()

	return obj
}

// As much as possible of the original IP packet that triggered the Redirect, without the Redirect IP packet exceeding the minimum IPv6 MTU. Hex string.
// Data returns a string
func (obj *flowIcmpv6NdpRedirectedHeaderOption) Data() string {

	return *obj.obj.Data

}

// As much as possible of the original IP packet that triggered the Redirect, without the Redirect IP packet exceeding the minimum IPv6 MTU. Hex string.
// Data returns a string
func (obj *flowIcmpv6NdpRedirectedHeaderOption) HasData() bool {
	return obj.obj.Data != nil
}

// As much as possible of the original IP packet that triggered the Redirect, without the Redirect IP packet exceeding the minimum IPv6 MTU. Hex string.
// SetData sets the string value in the FlowIcmpv6NdpRedirectedHeaderOption object
func (obj *flowIcmpv6NdpRedirectedHeaderOption) SetData(value string) FlowIcmpv6NdpRedirectedHeaderOption {

	obj.obj.Data = &value
	return obj
}

func (obj *flowIcmpv6NdpRedirectedHeaderOption) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Length != nil {

		obj.Length().validateObj(vObj, set_default)
	}

	if obj.obj.Reserved1 != nil {

		obj.Reserved1().validateObj(vObj, set_default)
	}

	if obj.obj.Reserved2 != nil {

		obj.Reserved2().validateObj(vObj, set_default)
	}

	if obj.obj.Data != nil {

		err := obj.validateHex(obj.Data())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowIcmpv6NdpRedirectedHeaderOption.Data"))
		}

	}

}

func (obj *flowIcmpv6NdpRedirectedHeaderOption) setDefault() {
	if obj.obj.Data == nil {
		obj.SetData("00")
	}

}
