package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv4OptionsTimestamp *****
type flowIpv4OptionsTimestamp struct {
	validation
	obj            *otg.FlowIpv4OptionsTimestamp
	marshaller     marshalFlowIpv4OptionsTimestamp
	unMarshaller   unMarshalFlowIpv4OptionsTimestamp
	pointerHolder  FlowIpv4OptionsTimestampPointer
	overflowHolder PatternFlowIpv4OptionsTimestampOverflow
	formatHolder   FlowIpv4OptionsTimestampFormat
}

func NewFlowIpv4OptionsTimestamp() FlowIpv4OptionsTimestamp {
	obj := flowIpv4OptionsTimestamp{obj: &otg.FlowIpv4OptionsTimestamp{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv4OptionsTimestamp) msg() *otg.FlowIpv4OptionsTimestamp {
	return obj.obj
}

func (obj *flowIpv4OptionsTimestamp) setMsg(msg *otg.FlowIpv4OptionsTimestamp) FlowIpv4OptionsTimestamp {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv4OptionsTimestamp struct {
	obj *flowIpv4OptionsTimestamp
}

type marshalFlowIpv4OptionsTimestamp interface {
	// ToProto marshals FlowIpv4OptionsTimestamp to protobuf object *otg.FlowIpv4OptionsTimestamp
	ToProto() (*otg.FlowIpv4OptionsTimestamp, error)
	// ToPbText marshals FlowIpv4OptionsTimestamp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv4OptionsTimestamp to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv4OptionsTimestamp to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv4OptionsTimestamp struct {
	obj *flowIpv4OptionsTimestamp
}

type unMarshalFlowIpv4OptionsTimestamp interface {
	// FromProto unmarshals FlowIpv4OptionsTimestamp from protobuf object *otg.FlowIpv4OptionsTimestamp
	FromProto(msg *otg.FlowIpv4OptionsTimestamp) (FlowIpv4OptionsTimestamp, error)
	// FromPbText unmarshals FlowIpv4OptionsTimestamp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv4OptionsTimestamp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv4OptionsTimestamp from JSON text
	FromJson(value string) error
}

func (obj *flowIpv4OptionsTimestamp) Marshal() marshalFlowIpv4OptionsTimestamp {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv4OptionsTimestamp{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv4OptionsTimestamp) Unmarshal() unMarshalFlowIpv4OptionsTimestamp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv4OptionsTimestamp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv4OptionsTimestamp) ToProto() (*otg.FlowIpv4OptionsTimestamp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv4OptionsTimestamp) FromProto(msg *otg.FlowIpv4OptionsTimestamp) (FlowIpv4OptionsTimestamp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv4OptionsTimestamp) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv4OptionsTimestamp) FromPbText(value string) error {
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

func (m *marshalflowIpv4OptionsTimestamp) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv4OptionsTimestamp) FromYaml(value string) error {
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

func (m *marshalflowIpv4OptionsTimestamp) ToJson() (string, error) {
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

func (m *unMarshalflowIpv4OptionsTimestamp) FromJson(value string) error {
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

func (obj *flowIpv4OptionsTimestamp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv4OptionsTimestamp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv4OptionsTimestamp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv4OptionsTimestamp) Clone() (FlowIpv4OptionsTimestamp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv4OptionsTimestamp()
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

func (obj *flowIpv4OptionsTimestamp) setNil() {
	obj.pointerHolder = nil
	obj.overflowHolder = nil
	obj.formatHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv4OptionsTimestamp is iPv4 timestamp option to be appended to the IPv4 header.
type FlowIpv4OptionsTimestamp interface {
	Validation
	// msg marshals FlowIpv4OptionsTimestamp to protobuf object *otg.FlowIpv4OptionsTimestamp
	// and doesn't set defaults
	msg() *otg.FlowIpv4OptionsTimestamp
	// setMsg unmarshals FlowIpv4OptionsTimestamp from protobuf object *otg.FlowIpv4OptionsTimestamp
	// and doesn't set defaults
	setMsg(*otg.FlowIpv4OptionsTimestamp) FlowIpv4OptionsTimestamp
	// provides marshal interface
	Marshal() marshalFlowIpv4OptionsTimestamp
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv4OptionsTimestamp
	// validate validates FlowIpv4OptionsTimestamp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv4OptionsTimestamp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Pointer returns FlowIpv4OptionsTimestampPointer, set in FlowIpv4OptionsTimestamp.
	// FlowIpv4OptionsTimestampPointer is the attribute pointer indicates the octet offset where the next router should begin recording its data.
	// Choices of input are,
	// 1. auto : The OTG implementation can provide a system generated value for this property. If the implementation is unable to generate a value, the default value must be used.
	// 2. value: User can configure the pointer value.
	Pointer() FlowIpv4OptionsTimestampPointer
	// SetPointer assigns FlowIpv4OptionsTimestampPointer provided by user to FlowIpv4OptionsTimestamp.
	// FlowIpv4OptionsTimestampPointer is the attribute pointer indicates the octet offset where the next router should begin recording its data.
	// Choices of input are,
	// 1. auto : The OTG implementation can provide a system generated value for this property. If the implementation is unable to generate a value, the default value must be used.
	// 2. value: User can configure the pointer value.
	SetPointer(value FlowIpv4OptionsTimestampPointer) FlowIpv4OptionsTimestamp
	// HasPointer checks if Pointer has been set in FlowIpv4OptionsTimestamp
	HasPointer() bool
	// Overflow returns PatternFlowIpv4OptionsTimestampOverflow, set in FlowIpv4OptionsTimestamp.
	// PatternFlowIpv4OptionsTimestampOverflow is a counter that indicates the number of intermediate nodes that were unable to record a timestamp because the options data area was full.
	Overflow() PatternFlowIpv4OptionsTimestampOverflow
	// SetOverflow assigns PatternFlowIpv4OptionsTimestampOverflow provided by user to FlowIpv4OptionsTimestamp.
	// PatternFlowIpv4OptionsTimestampOverflow is a counter that indicates the number of intermediate nodes that were unable to record a timestamp because the options data area was full.
	SetOverflow(value PatternFlowIpv4OptionsTimestampOverflow) FlowIpv4OptionsTimestamp
	// HasOverflow checks if Overflow has been set in FlowIpv4OptionsTimestamp
	HasOverflow() bool
	// Format returns FlowIpv4OptionsTimestampFormat, set in FlowIpv4OptionsTimestamp.
	// FlowIpv4OptionsTimestampFormat is the format field defines the structural layout of the options data area and the specific recording logic routers must follow when appending timestamps and IP addresses.
	Format() FlowIpv4OptionsTimestampFormat
	// SetFormat assigns FlowIpv4OptionsTimestampFormat provided by user to FlowIpv4OptionsTimestamp.
	// FlowIpv4OptionsTimestampFormat is the format field defines the structural layout of the options data area and the specific recording logic routers must follow when appending timestamps and IP addresses.
	SetFormat(value FlowIpv4OptionsTimestampFormat) FlowIpv4OptionsTimestamp
	// HasFormat checks if Format has been set in FlowIpv4OptionsTimestamp
	HasFormat() bool
	setNil()
}

// description is TBD
// Pointer returns a FlowIpv4OptionsTimestampPointer
func (obj *flowIpv4OptionsTimestamp) Pointer() FlowIpv4OptionsTimestampPointer {
	if obj.obj.Pointer == nil {
		obj.obj.Pointer = NewFlowIpv4OptionsTimestampPointer().msg()
	}
	if obj.pointerHolder == nil {
		obj.pointerHolder = &flowIpv4OptionsTimestampPointer{obj: obj.obj.Pointer}
	}
	return obj.pointerHolder
}

// description is TBD
// Pointer returns a FlowIpv4OptionsTimestampPointer
func (obj *flowIpv4OptionsTimestamp) HasPointer() bool {
	return obj.obj.Pointer != nil
}

// description is TBD
// SetPointer sets the FlowIpv4OptionsTimestampPointer value in the FlowIpv4OptionsTimestamp object
func (obj *flowIpv4OptionsTimestamp) SetPointer(value FlowIpv4OptionsTimestampPointer) FlowIpv4OptionsTimestamp {

	obj.pointerHolder = nil
	obj.obj.Pointer = value.msg()

	return obj
}

// A counter that indicates the number of intermediate nodes that were unable to record a timestamp because the options data area was full.
// Overflow returns a PatternFlowIpv4OptionsTimestampOverflow
func (obj *flowIpv4OptionsTimestamp) Overflow() PatternFlowIpv4OptionsTimestampOverflow {
	if obj.obj.Overflow == nil {
		obj.obj.Overflow = NewPatternFlowIpv4OptionsTimestampOverflow().msg()
	}
	if obj.overflowHolder == nil {
		obj.overflowHolder = &patternFlowIpv4OptionsTimestampOverflow{obj: obj.obj.Overflow}
	}
	return obj.overflowHolder
}

// A counter that indicates the number of intermediate nodes that were unable to record a timestamp because the options data area was full.
// Overflow returns a PatternFlowIpv4OptionsTimestampOverflow
func (obj *flowIpv4OptionsTimestamp) HasOverflow() bool {
	return obj.obj.Overflow != nil
}

// A counter that indicates the number of intermediate nodes that were unable to record a timestamp because the options data area was full.
// SetOverflow sets the PatternFlowIpv4OptionsTimestampOverflow value in the FlowIpv4OptionsTimestamp object
func (obj *flowIpv4OptionsTimestamp) SetOverflow(value PatternFlowIpv4OptionsTimestampOverflow) FlowIpv4OptionsTimestamp {

	obj.overflowHolder = nil
	obj.obj.Overflow = value.msg()

	return obj
}

// description is TBD
// Format returns a FlowIpv4OptionsTimestampFormat
func (obj *flowIpv4OptionsTimestamp) Format() FlowIpv4OptionsTimestampFormat {
	if obj.obj.Format == nil {
		obj.obj.Format = NewFlowIpv4OptionsTimestampFormat().msg()
	}
	if obj.formatHolder == nil {
		obj.formatHolder = &flowIpv4OptionsTimestampFormat{obj: obj.obj.Format}
	}
	return obj.formatHolder
}

// description is TBD
// Format returns a FlowIpv4OptionsTimestampFormat
func (obj *flowIpv4OptionsTimestamp) HasFormat() bool {
	return obj.obj.Format != nil
}

// description is TBD
// SetFormat sets the FlowIpv4OptionsTimestampFormat value in the FlowIpv4OptionsTimestamp object
func (obj *flowIpv4OptionsTimestamp) SetFormat(value FlowIpv4OptionsTimestampFormat) FlowIpv4OptionsTimestamp {

	obj.formatHolder = nil
	obj.obj.Format = value.msg()

	return obj
}

func (obj *flowIpv4OptionsTimestamp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Pointer != nil {

		obj.Pointer().validateObj(vObj, set_default)
	}

	if obj.obj.Overflow != nil {

		obj.Overflow().validateObj(vObj, set_default)
	}

	if obj.obj.Format != nil {

		obj.Format().validateObj(vObj, set_default)
	}

}

func (obj *flowIpv4OptionsTimestamp) setDefault() {

}
