package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv4OptionsTimestampFormatTimestamp *****
type flowIpv4OptionsTimestampFormatTimestamp struct {
	validation
	obj             *otg.FlowIpv4OptionsTimestampFormatTimestamp
	marshaller      marshalFlowIpv4OptionsTimestampFormatTimestamp
	unMarshaller    unMarshalFlowIpv4OptionsTimestampFormatTimestamp
	timestampHolder PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
}

func NewFlowIpv4OptionsTimestampFormatTimestamp() FlowIpv4OptionsTimestampFormatTimestamp {
	obj := flowIpv4OptionsTimestampFormatTimestamp{obj: &otg.FlowIpv4OptionsTimestampFormatTimestamp{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv4OptionsTimestampFormatTimestamp) msg() *otg.FlowIpv4OptionsTimestampFormatTimestamp {
	return obj.obj
}

func (obj *flowIpv4OptionsTimestampFormatTimestamp) setMsg(msg *otg.FlowIpv4OptionsTimestampFormatTimestamp) FlowIpv4OptionsTimestampFormatTimestamp {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv4OptionsTimestampFormatTimestamp struct {
	obj *flowIpv4OptionsTimestampFormatTimestamp
}

type marshalFlowIpv4OptionsTimestampFormatTimestamp interface {
	// ToProto marshals FlowIpv4OptionsTimestampFormatTimestamp to protobuf object *otg.FlowIpv4OptionsTimestampFormatTimestamp
	ToProto() (*otg.FlowIpv4OptionsTimestampFormatTimestamp, error)
	// ToPbText marshals FlowIpv4OptionsTimestampFormatTimestamp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv4OptionsTimestampFormatTimestamp to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv4OptionsTimestampFormatTimestamp to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv4OptionsTimestampFormatTimestamp struct {
	obj *flowIpv4OptionsTimestampFormatTimestamp
}

type unMarshalFlowIpv4OptionsTimestampFormatTimestamp interface {
	// FromProto unmarshals FlowIpv4OptionsTimestampFormatTimestamp from protobuf object *otg.FlowIpv4OptionsTimestampFormatTimestamp
	FromProto(msg *otg.FlowIpv4OptionsTimestampFormatTimestamp) (FlowIpv4OptionsTimestampFormatTimestamp, error)
	// FromPbText unmarshals FlowIpv4OptionsTimestampFormatTimestamp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv4OptionsTimestampFormatTimestamp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv4OptionsTimestampFormatTimestamp from JSON text
	FromJson(value string) error
}

func (obj *flowIpv4OptionsTimestampFormatTimestamp) Marshal() marshalFlowIpv4OptionsTimestampFormatTimestamp {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv4OptionsTimestampFormatTimestamp{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv4OptionsTimestampFormatTimestamp) Unmarshal() unMarshalFlowIpv4OptionsTimestampFormatTimestamp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv4OptionsTimestampFormatTimestamp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv4OptionsTimestampFormatTimestamp) ToProto() (*otg.FlowIpv4OptionsTimestampFormatTimestamp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv4OptionsTimestampFormatTimestamp) FromProto(msg *otg.FlowIpv4OptionsTimestampFormatTimestamp) (FlowIpv4OptionsTimestampFormatTimestamp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv4OptionsTimestampFormatTimestamp) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv4OptionsTimestampFormatTimestamp) FromPbText(value string) error {
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

func (m *marshalflowIpv4OptionsTimestampFormatTimestamp) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv4OptionsTimestampFormatTimestamp) FromYaml(value string) error {
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

func (m *marshalflowIpv4OptionsTimestampFormatTimestamp) ToJson() (string, error) {
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

func (m *unMarshalflowIpv4OptionsTimestampFormatTimestamp) FromJson(value string) error {
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

func (obj *flowIpv4OptionsTimestampFormatTimestamp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv4OptionsTimestampFormatTimestamp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv4OptionsTimestampFormatTimestamp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv4OptionsTimestampFormatTimestamp) Clone() (FlowIpv4OptionsTimestampFormatTimestamp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv4OptionsTimestampFormatTimestamp()
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

func (obj *flowIpv4OptionsTimestampFormatTimestamp) setNil() {
	obj.timestampHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv4OptionsTimestampFormatTimestamp is description is TBD
type FlowIpv4OptionsTimestampFormatTimestamp interface {
	Validation
	// msg marshals FlowIpv4OptionsTimestampFormatTimestamp to protobuf object *otg.FlowIpv4OptionsTimestampFormatTimestamp
	// and doesn't set defaults
	msg() *otg.FlowIpv4OptionsTimestampFormatTimestamp
	// setMsg unmarshals FlowIpv4OptionsTimestampFormatTimestamp from protobuf object *otg.FlowIpv4OptionsTimestampFormatTimestamp
	// and doesn't set defaults
	setMsg(*otg.FlowIpv4OptionsTimestampFormatTimestamp) FlowIpv4OptionsTimestampFormatTimestamp
	// provides marshal interface
	Marshal() marshalFlowIpv4OptionsTimestampFormatTimestamp
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv4OptionsTimestampFormatTimestamp
	// validate validates FlowIpv4OptionsTimestampFormatTimestamp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv4OptionsTimestampFormatTimestamp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Timestamp returns PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp, set in FlowIpv4OptionsTimestampFormatTimestamp.
	// PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp is a 32-bit value representing the time of packet processing, recorded as the number of milliseconds elapsed since midnight Universal Time (UT).
	Timestamp() PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
	// SetTimestamp assigns PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp provided by user to FlowIpv4OptionsTimestampFormatTimestamp.
	// PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp is a 32-bit value representing the time of packet processing, recorded as the number of milliseconds elapsed since midnight Universal Time (UT).
	SetTimestamp(value PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp) FlowIpv4OptionsTimestampFormatTimestamp
	// HasTimestamp checks if Timestamp has been set in FlowIpv4OptionsTimestampFormatTimestamp
	HasTimestamp() bool
	setNil()
}

// description is TBD
// Timestamp returns a PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
func (obj *flowIpv4OptionsTimestampFormatTimestamp) Timestamp() PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp {
	if obj.obj.Timestamp == nil {
		obj.obj.Timestamp = NewPatternFlowIpv4OptionsTimestampFormatTimestampTimestamp().msg()
	}
	if obj.timestampHolder == nil {
		obj.timestampHolder = &patternFlowIpv4OptionsTimestampFormatTimestampTimestamp{obj: obj.obj.Timestamp}
	}
	return obj.timestampHolder
}

// description is TBD
// Timestamp returns a PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
func (obj *flowIpv4OptionsTimestampFormatTimestamp) HasTimestamp() bool {
	return obj.obj.Timestamp != nil
}

// description is TBD
// SetTimestamp sets the PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp value in the FlowIpv4OptionsTimestampFormatTimestamp object
func (obj *flowIpv4OptionsTimestampFormatTimestamp) SetTimestamp(value PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp) FlowIpv4OptionsTimestampFormatTimestamp {

	obj.timestampHolder = nil
	obj.obj.Timestamp = value.msg()

	return obj
}

func (obj *flowIpv4OptionsTimestampFormatTimestamp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Timestamp != nil {

		obj.Timestamp().validateObj(vObj, set_default)
	}

}

func (obj *flowIpv4OptionsTimestampFormatTimestamp) setDefault() {

}
