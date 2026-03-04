package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv4OptionsTimestampFormatAddressAndTimestamp *****
type flowIpv4OptionsTimestampFormatAddressAndTimestamp struct {
	validation
	obj             *otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamp
	marshaller      marshalFlowIpv4OptionsTimestampFormatAddressAndTimestamp
	unMarshaller    unMarshalFlowIpv4OptionsTimestampFormatAddressAndTimestamp
	addressHolder   PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
	timestampHolder PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
}

func NewFlowIpv4OptionsTimestampFormatAddressAndTimestamp() FlowIpv4OptionsTimestampFormatAddressAndTimestamp {
	obj := flowIpv4OptionsTimestampFormatAddressAndTimestamp{obj: &otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamp{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamp) msg() *otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamp {
	return obj.obj
}

func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamp) setMsg(msg *otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamp) FlowIpv4OptionsTimestampFormatAddressAndTimestamp {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv4OptionsTimestampFormatAddressAndTimestamp struct {
	obj *flowIpv4OptionsTimestampFormatAddressAndTimestamp
}

type marshalFlowIpv4OptionsTimestampFormatAddressAndTimestamp interface {
	// ToProto marshals FlowIpv4OptionsTimestampFormatAddressAndTimestamp to protobuf object *otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamp
	ToProto() (*otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamp, error)
	// ToPbText marshals FlowIpv4OptionsTimestampFormatAddressAndTimestamp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv4OptionsTimestampFormatAddressAndTimestamp to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv4OptionsTimestampFormatAddressAndTimestamp to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv4OptionsTimestampFormatAddressAndTimestamp struct {
	obj *flowIpv4OptionsTimestampFormatAddressAndTimestamp
}

type unMarshalFlowIpv4OptionsTimestampFormatAddressAndTimestamp interface {
	// FromProto unmarshals FlowIpv4OptionsTimestampFormatAddressAndTimestamp from protobuf object *otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamp
	FromProto(msg *otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamp) (FlowIpv4OptionsTimestampFormatAddressAndTimestamp, error)
	// FromPbText unmarshals FlowIpv4OptionsTimestampFormatAddressAndTimestamp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv4OptionsTimestampFormatAddressAndTimestamp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv4OptionsTimestampFormatAddressAndTimestamp from JSON text
	FromJson(value string) error
}

func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamp) Marshal() marshalFlowIpv4OptionsTimestampFormatAddressAndTimestamp {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv4OptionsTimestampFormatAddressAndTimestamp{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamp) Unmarshal() unMarshalFlowIpv4OptionsTimestampFormatAddressAndTimestamp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv4OptionsTimestampFormatAddressAndTimestamp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv4OptionsTimestampFormatAddressAndTimestamp) ToProto() (*otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv4OptionsTimestampFormatAddressAndTimestamp) FromProto(msg *otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamp) (FlowIpv4OptionsTimestampFormatAddressAndTimestamp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv4OptionsTimestampFormatAddressAndTimestamp) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv4OptionsTimestampFormatAddressAndTimestamp) FromPbText(value string) error {
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

func (m *marshalflowIpv4OptionsTimestampFormatAddressAndTimestamp) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv4OptionsTimestampFormatAddressAndTimestamp) FromYaml(value string) error {
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

func (m *marshalflowIpv4OptionsTimestampFormatAddressAndTimestamp) ToJson() (string, error) {
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

func (m *unMarshalflowIpv4OptionsTimestampFormatAddressAndTimestamp) FromJson(value string) error {
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

func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamp) Clone() (FlowIpv4OptionsTimestampFormatAddressAndTimestamp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv4OptionsTimestampFormatAddressAndTimestamp()
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

func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamp) setNil() {
	obj.addressHolder = nil
	obj.timestampHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv4OptionsTimestampFormatAddressAndTimestamp is description is TBD
type FlowIpv4OptionsTimestampFormatAddressAndTimestamp interface {
	Validation
	// msg marshals FlowIpv4OptionsTimestampFormatAddressAndTimestamp to protobuf object *otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamp
	// and doesn't set defaults
	msg() *otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamp
	// setMsg unmarshals FlowIpv4OptionsTimestampFormatAddressAndTimestamp from protobuf object *otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamp
	// and doesn't set defaults
	setMsg(*otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamp) FlowIpv4OptionsTimestampFormatAddressAndTimestamp
	// provides marshal interface
	Marshal() marshalFlowIpv4OptionsTimestampFormatAddressAndTimestamp
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv4OptionsTimestampFormatAddressAndTimestamp
	// validate validates FlowIpv4OptionsTimestampFormatAddressAndTimestamp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv4OptionsTimestampFormatAddressAndTimestamp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Address returns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress, set in FlowIpv4OptionsTimestampFormatAddressAndTimestamp.
	// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress is iPv4 address of the router interface that handled the packet, serving as a network identifier for the corresponding timestamp entry.
	Address() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
	// SetAddress assigns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress provided by user to FlowIpv4OptionsTimestampFormatAddressAndTimestamp.
	// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress is iPv4 address of the router interface that handled the packet, serving as a network identifier for the corresponding timestamp entry.
	SetAddress(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) FlowIpv4OptionsTimestampFormatAddressAndTimestamp
	// HasAddress checks if Address has been set in FlowIpv4OptionsTimestampFormatAddressAndTimestamp
	HasAddress() bool
	// Timestamp returns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp, set in FlowIpv4OptionsTimestampFormatAddressAndTimestamp.
	// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp is a 32-bit value representing the time of packet processing, recorded as the number of milliseconds elapsed since midnight Universal Time (UT).
	Timestamp() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
	// SetTimestamp assigns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp provided by user to FlowIpv4OptionsTimestampFormatAddressAndTimestamp.
	// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp is a 32-bit value representing the time of packet processing, recorded as the number of milliseconds elapsed since midnight Universal Time (UT).
	SetTimestamp(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) FlowIpv4OptionsTimestampFormatAddressAndTimestamp
	// HasTimestamp checks if Timestamp has been set in FlowIpv4OptionsTimestampFormatAddressAndTimestamp
	HasTimestamp() bool
	setNil()
}

// description is TBD
// Address returns a PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamp) Address() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress {
	if obj.obj.Address == nil {
		obj.obj.Address = NewPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress().msg()
	}
	if obj.addressHolder == nil {
		obj.addressHolder = &patternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress{obj: obj.obj.Address}
	}
	return obj.addressHolder
}

// description is TBD
// Address returns a PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress
func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamp) HasAddress() bool {
	return obj.obj.Address != nil
}

// description is TBD
// SetAddress sets the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress value in the FlowIpv4OptionsTimestampFormatAddressAndTimestamp object
func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamp) SetAddress(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampAddress) FlowIpv4OptionsTimestampFormatAddressAndTimestamp {

	obj.addressHolder = nil
	obj.obj.Address = value.msg()

	return obj
}

// description is TBD
// Timestamp returns a PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamp) Timestamp() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp {
	if obj.obj.Timestamp == nil {
		obj.obj.Timestamp = NewPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp().msg()
	}
	if obj.timestampHolder == nil {
		obj.timestampHolder = &patternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp{obj: obj.obj.Timestamp}
	}
	return obj.timestampHolder
}

// description is TBD
// Timestamp returns a PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp
func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamp) HasTimestamp() bool {
	return obj.obj.Timestamp != nil
}

// description is TBD
// SetTimestamp sets the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp value in the FlowIpv4OptionsTimestampFormatAddressAndTimestamp object
func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamp) SetTimestamp(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampTimestamp) FlowIpv4OptionsTimestampFormatAddressAndTimestamp {

	obj.timestampHolder = nil
	obj.obj.Timestamp = value.msg()

	return obj
}

func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Address != nil {

		obj.Address().validateObj(vObj, set_default)
	}

	if obj.obj.Timestamp != nil {

		obj.Timestamp().validateObj(vObj, set_default)
	}

}

func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamp) setDefault() {

}
