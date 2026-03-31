package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv4OptionsTimestampFormatTimestamps *****
type flowIpv4OptionsTimestampFormatTimestamps struct {
	validation
	obj             *otg.FlowIpv4OptionsTimestampFormatTimestamps
	marshaller      marshalFlowIpv4OptionsTimestampFormatTimestamps
	unMarshaller    unMarshalFlowIpv4OptionsTimestampFormatTimestamps
	timestampHolder PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
}

func NewFlowIpv4OptionsTimestampFormatTimestamps() FlowIpv4OptionsTimestampFormatTimestamps {
	obj := flowIpv4OptionsTimestampFormatTimestamps{obj: &otg.FlowIpv4OptionsTimestampFormatTimestamps{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv4OptionsTimestampFormatTimestamps) msg() *otg.FlowIpv4OptionsTimestampFormatTimestamps {
	return obj.obj
}

func (obj *flowIpv4OptionsTimestampFormatTimestamps) setMsg(msg *otg.FlowIpv4OptionsTimestampFormatTimestamps) FlowIpv4OptionsTimestampFormatTimestamps {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv4OptionsTimestampFormatTimestamps struct {
	obj *flowIpv4OptionsTimestampFormatTimestamps
}

type marshalFlowIpv4OptionsTimestampFormatTimestamps interface {
	// ToProto marshals FlowIpv4OptionsTimestampFormatTimestamps to protobuf object *otg.FlowIpv4OptionsTimestampFormatTimestamps
	ToProto() (*otg.FlowIpv4OptionsTimestampFormatTimestamps, error)
	// ToPbText marshals FlowIpv4OptionsTimestampFormatTimestamps to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv4OptionsTimestampFormatTimestamps to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv4OptionsTimestampFormatTimestamps to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv4OptionsTimestampFormatTimestamps struct {
	obj *flowIpv4OptionsTimestampFormatTimestamps
}

type unMarshalFlowIpv4OptionsTimestampFormatTimestamps interface {
	// FromProto unmarshals FlowIpv4OptionsTimestampFormatTimestamps from protobuf object *otg.FlowIpv4OptionsTimestampFormatTimestamps
	FromProto(msg *otg.FlowIpv4OptionsTimestampFormatTimestamps) (FlowIpv4OptionsTimestampFormatTimestamps, error)
	// FromPbText unmarshals FlowIpv4OptionsTimestampFormatTimestamps from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv4OptionsTimestampFormatTimestamps from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv4OptionsTimestampFormatTimestamps from JSON text
	FromJson(value string) error
}

func (obj *flowIpv4OptionsTimestampFormatTimestamps) Marshal() marshalFlowIpv4OptionsTimestampFormatTimestamps {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv4OptionsTimestampFormatTimestamps{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv4OptionsTimestampFormatTimestamps) Unmarshal() unMarshalFlowIpv4OptionsTimestampFormatTimestamps {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv4OptionsTimestampFormatTimestamps{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv4OptionsTimestampFormatTimestamps) ToProto() (*otg.FlowIpv4OptionsTimestampFormatTimestamps, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv4OptionsTimestampFormatTimestamps) FromProto(msg *otg.FlowIpv4OptionsTimestampFormatTimestamps) (FlowIpv4OptionsTimestampFormatTimestamps, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv4OptionsTimestampFormatTimestamps) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv4OptionsTimestampFormatTimestamps) FromPbText(value string) error {
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

func (m *marshalflowIpv4OptionsTimestampFormatTimestamps) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv4OptionsTimestampFormatTimestamps) FromYaml(value string) error {
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

func (m *marshalflowIpv4OptionsTimestampFormatTimestamps) ToJson() (string, error) {
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

func (m *unMarshalflowIpv4OptionsTimestampFormatTimestamps) FromJson(value string) error {
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

func (obj *flowIpv4OptionsTimestampFormatTimestamps) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv4OptionsTimestampFormatTimestamps) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv4OptionsTimestampFormatTimestamps) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv4OptionsTimestampFormatTimestamps) Clone() (FlowIpv4OptionsTimestampFormatTimestamps, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv4OptionsTimestampFormatTimestamps()
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

func (obj *flowIpv4OptionsTimestampFormatTimestamps) setNil() {
	obj.timestampHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv4OptionsTimestampFormatTimestamps is description is TBD
type FlowIpv4OptionsTimestampFormatTimestamps interface {
	Validation
	// msg marshals FlowIpv4OptionsTimestampFormatTimestamps to protobuf object *otg.FlowIpv4OptionsTimestampFormatTimestamps
	// and doesn't set defaults
	msg() *otg.FlowIpv4OptionsTimestampFormatTimestamps
	// setMsg unmarshals FlowIpv4OptionsTimestampFormatTimestamps from protobuf object *otg.FlowIpv4OptionsTimestampFormatTimestamps
	// and doesn't set defaults
	setMsg(*otg.FlowIpv4OptionsTimestampFormatTimestamps) FlowIpv4OptionsTimestampFormatTimestamps
	// provides marshal interface
	Marshal() marshalFlowIpv4OptionsTimestampFormatTimestamps
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv4OptionsTimestampFormatTimestamps
	// validate validates FlowIpv4OptionsTimestampFormatTimestamps
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv4OptionsTimestampFormatTimestamps, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Timestamp returns PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp, set in FlowIpv4OptionsTimestampFormatTimestamps.
	// PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp is a 32-bit value representing the time of packet processing, recorded as the number of milliseconds elapsed since midnight Universal Time (UT).
	Timestamp() PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
	// SetTimestamp assigns PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp provided by user to FlowIpv4OptionsTimestampFormatTimestamps.
	// PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp is a 32-bit value representing the time of packet processing, recorded as the number of milliseconds elapsed since midnight Universal Time (UT).
	SetTimestamp(value PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) FlowIpv4OptionsTimestampFormatTimestamps
	// HasTimestamp checks if Timestamp has been set in FlowIpv4OptionsTimestampFormatTimestamps
	HasTimestamp() bool
	setNil()
}

// description is TBD
// Timestamp returns a PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
func (obj *flowIpv4OptionsTimestampFormatTimestamps) Timestamp() PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp {
	if obj.obj.Timestamp == nil {
		obj.obj.Timestamp = NewPatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp().msg()
	}
	if obj.timestampHolder == nil {
		obj.timestampHolder = &patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp{obj: obj.obj.Timestamp}
	}
	return obj.timestampHolder
}

// description is TBD
// Timestamp returns a PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
func (obj *flowIpv4OptionsTimestampFormatTimestamps) HasTimestamp() bool {
	return obj.obj.Timestamp != nil
}

// description is TBD
// SetTimestamp sets the PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp value in the FlowIpv4OptionsTimestampFormatTimestamps object
func (obj *flowIpv4OptionsTimestampFormatTimestamps) SetTimestamp(value PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) FlowIpv4OptionsTimestampFormatTimestamps {

	obj.timestampHolder = nil
	obj.obj.Timestamp = value.msg()

	return obj
}

func (obj *flowIpv4OptionsTimestampFormatTimestamps) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Timestamp != nil {

		obj.Timestamp().validateObj(vObj, set_default)
	}

}

func (obj *flowIpv4OptionsTimestampFormatTimestamps) setDefault() {

}
