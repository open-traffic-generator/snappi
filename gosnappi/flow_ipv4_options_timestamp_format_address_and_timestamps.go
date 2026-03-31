package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv4OptionsTimestampFormatAddressAndTimestamps *****
type flowIpv4OptionsTimestampFormatAddressAndTimestamps struct {
	validation
	obj             *otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamps
	marshaller      marshalFlowIpv4OptionsTimestampFormatAddressAndTimestamps
	unMarshaller    unMarshalFlowIpv4OptionsTimestampFormatAddressAndTimestamps
	addressHolder   PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
	timestampHolder PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
}

func NewFlowIpv4OptionsTimestampFormatAddressAndTimestamps() FlowIpv4OptionsTimestampFormatAddressAndTimestamps {
	obj := flowIpv4OptionsTimestampFormatAddressAndTimestamps{obj: &otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamps{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamps) msg() *otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamps {
	return obj.obj
}

func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamps) setMsg(msg *otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamps) FlowIpv4OptionsTimestampFormatAddressAndTimestamps {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv4OptionsTimestampFormatAddressAndTimestamps struct {
	obj *flowIpv4OptionsTimestampFormatAddressAndTimestamps
}

type marshalFlowIpv4OptionsTimestampFormatAddressAndTimestamps interface {
	// ToProto marshals FlowIpv4OptionsTimestampFormatAddressAndTimestamps to protobuf object *otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamps
	ToProto() (*otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamps, error)
	// ToPbText marshals FlowIpv4OptionsTimestampFormatAddressAndTimestamps to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv4OptionsTimestampFormatAddressAndTimestamps to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv4OptionsTimestampFormatAddressAndTimestamps to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv4OptionsTimestampFormatAddressAndTimestamps struct {
	obj *flowIpv4OptionsTimestampFormatAddressAndTimestamps
}

type unMarshalFlowIpv4OptionsTimestampFormatAddressAndTimestamps interface {
	// FromProto unmarshals FlowIpv4OptionsTimestampFormatAddressAndTimestamps from protobuf object *otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamps
	FromProto(msg *otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamps) (FlowIpv4OptionsTimestampFormatAddressAndTimestamps, error)
	// FromPbText unmarshals FlowIpv4OptionsTimestampFormatAddressAndTimestamps from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv4OptionsTimestampFormatAddressAndTimestamps from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv4OptionsTimestampFormatAddressAndTimestamps from JSON text
	FromJson(value string) error
}

func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamps) Marshal() marshalFlowIpv4OptionsTimestampFormatAddressAndTimestamps {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv4OptionsTimestampFormatAddressAndTimestamps{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamps) Unmarshal() unMarshalFlowIpv4OptionsTimestampFormatAddressAndTimestamps {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv4OptionsTimestampFormatAddressAndTimestamps{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv4OptionsTimestampFormatAddressAndTimestamps) ToProto() (*otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamps, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv4OptionsTimestampFormatAddressAndTimestamps) FromProto(msg *otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamps) (FlowIpv4OptionsTimestampFormatAddressAndTimestamps, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv4OptionsTimestampFormatAddressAndTimestamps) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv4OptionsTimestampFormatAddressAndTimestamps) FromPbText(value string) error {
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

func (m *marshalflowIpv4OptionsTimestampFormatAddressAndTimestamps) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv4OptionsTimestampFormatAddressAndTimestamps) FromYaml(value string) error {
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

func (m *marshalflowIpv4OptionsTimestampFormatAddressAndTimestamps) ToJson() (string, error) {
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

func (m *unMarshalflowIpv4OptionsTimestampFormatAddressAndTimestamps) FromJson(value string) error {
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

func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamps) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamps) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamps) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamps) Clone() (FlowIpv4OptionsTimestampFormatAddressAndTimestamps, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv4OptionsTimestampFormatAddressAndTimestamps()
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

func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamps) setNil() {
	obj.addressHolder = nil
	obj.timestampHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv4OptionsTimestampFormatAddressAndTimestamps is description is TBD
type FlowIpv4OptionsTimestampFormatAddressAndTimestamps interface {
	Validation
	// msg marshals FlowIpv4OptionsTimestampFormatAddressAndTimestamps to protobuf object *otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamps
	// and doesn't set defaults
	msg() *otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamps
	// setMsg unmarshals FlowIpv4OptionsTimestampFormatAddressAndTimestamps from protobuf object *otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamps
	// and doesn't set defaults
	setMsg(*otg.FlowIpv4OptionsTimestampFormatAddressAndTimestamps) FlowIpv4OptionsTimestampFormatAddressAndTimestamps
	// provides marshal interface
	Marshal() marshalFlowIpv4OptionsTimestampFormatAddressAndTimestamps
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv4OptionsTimestampFormatAddressAndTimestamps
	// validate validates FlowIpv4OptionsTimestampFormatAddressAndTimestamps
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv4OptionsTimestampFormatAddressAndTimestamps, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Address returns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress, set in FlowIpv4OptionsTimestampFormatAddressAndTimestamps.
	// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress is iPv4 address of the router interface that handled the packet, serving as a network identifier for the corresponding timestamp entry.
	Address() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
	// SetAddress assigns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress provided by user to FlowIpv4OptionsTimestampFormatAddressAndTimestamps.
	// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress is iPv4 address of the router interface that handled the packet, serving as a network identifier for the corresponding timestamp entry.
	SetAddress(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) FlowIpv4OptionsTimestampFormatAddressAndTimestamps
	// HasAddress checks if Address has been set in FlowIpv4OptionsTimestampFormatAddressAndTimestamps
	HasAddress() bool
	// Timestamp returns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp, set in FlowIpv4OptionsTimestampFormatAddressAndTimestamps.
	// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp is a 32-bit value representing the time of packet processing, recorded as the number of milliseconds elapsed since midnight Universal Time (UT).
	Timestamp() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
	// SetTimestamp assigns PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp provided by user to FlowIpv4OptionsTimestampFormatAddressAndTimestamps.
	// PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp is a 32-bit value representing the time of packet processing, recorded as the number of milliseconds elapsed since midnight Universal Time (UT).
	SetTimestamp(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) FlowIpv4OptionsTimestampFormatAddressAndTimestamps
	// HasTimestamp checks if Timestamp has been set in FlowIpv4OptionsTimestampFormatAddressAndTimestamps
	HasTimestamp() bool
	setNil()
}

// description is TBD
// Address returns a PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamps) Address() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress {
	if obj.obj.Address == nil {
		obj.obj.Address = NewPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress().msg()
	}
	if obj.addressHolder == nil {
		obj.addressHolder = &patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress{obj: obj.obj.Address}
	}
	return obj.addressHolder
}

// description is TBD
// Address returns a PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress
func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamps) HasAddress() bool {
	return obj.obj.Address != nil
}

// description is TBD
// SetAddress sets the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress value in the FlowIpv4OptionsTimestampFormatAddressAndTimestamps object
func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamps) SetAddress(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsAddress) FlowIpv4OptionsTimestampFormatAddressAndTimestamps {

	obj.addressHolder = nil
	obj.obj.Address = value.msg()

	return obj
}

// description is TBD
// Timestamp returns a PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamps) Timestamp() PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp {
	if obj.obj.Timestamp == nil {
		obj.obj.Timestamp = NewPatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp().msg()
	}
	if obj.timestampHolder == nil {
		obj.timestampHolder = &patternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp{obj: obj.obj.Timestamp}
	}
	return obj.timestampHolder
}

// description is TBD
// Timestamp returns a PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp
func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamps) HasTimestamp() bool {
	return obj.obj.Timestamp != nil
}

// description is TBD
// SetTimestamp sets the PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp value in the FlowIpv4OptionsTimestampFormatAddressAndTimestamps object
func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamps) SetTimestamp(value PatternFlowIpv4OptionsTimestampFormatAddressAndTimestampsTimestamp) FlowIpv4OptionsTimestampFormatAddressAndTimestamps {

	obj.timestampHolder = nil
	obj.obj.Timestamp = value.msg()

	return obj
}

func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamps) validateObj(vObj *validation, set_default bool) {
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

func (obj *flowIpv4OptionsTimestampFormatAddressAndTimestamps) setDefault() {

}
