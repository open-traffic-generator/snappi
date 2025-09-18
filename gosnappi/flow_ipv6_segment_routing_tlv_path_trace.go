package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6SegmentRoutingTlvPathTrace *****
type flowIpv6SegmentRoutingTlvPathTrace struct {
	validation
	obj                  *otg.FlowIpv6SegmentRoutingTlvPathTrace
	marshaller           marshalFlowIpv6SegmentRoutingTlvPathTrace
	unMarshaller         unMarshalFlowIpv6SegmentRoutingTlvPathTrace
	interfaceIdHolder    PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
	interfaceLoadHolder  PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
	sessionIdHolder      PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
	sequenceNumberHolder PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
}

func NewFlowIpv6SegmentRoutingTlvPathTrace() FlowIpv6SegmentRoutingTlvPathTrace {
	obj := flowIpv6SegmentRoutingTlvPathTrace{obj: &otg.FlowIpv6SegmentRoutingTlvPathTrace{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6SegmentRoutingTlvPathTrace) msg() *otg.FlowIpv6SegmentRoutingTlvPathTrace {
	return obj.obj
}

func (obj *flowIpv6SegmentRoutingTlvPathTrace) setMsg(msg *otg.FlowIpv6SegmentRoutingTlvPathTrace) FlowIpv6SegmentRoutingTlvPathTrace {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6SegmentRoutingTlvPathTrace struct {
	obj *flowIpv6SegmentRoutingTlvPathTrace
}

type marshalFlowIpv6SegmentRoutingTlvPathTrace interface {
	// ToProto marshals FlowIpv6SegmentRoutingTlvPathTrace to protobuf object *otg.FlowIpv6SegmentRoutingTlvPathTrace
	ToProto() (*otg.FlowIpv6SegmentRoutingTlvPathTrace, error)
	// ToPbText marshals FlowIpv6SegmentRoutingTlvPathTrace to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6SegmentRoutingTlvPathTrace to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6SegmentRoutingTlvPathTrace to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6SegmentRoutingTlvPathTrace struct {
	obj *flowIpv6SegmentRoutingTlvPathTrace
}

type unMarshalFlowIpv6SegmentRoutingTlvPathTrace interface {
	// FromProto unmarshals FlowIpv6SegmentRoutingTlvPathTrace from protobuf object *otg.FlowIpv6SegmentRoutingTlvPathTrace
	FromProto(msg *otg.FlowIpv6SegmentRoutingTlvPathTrace) (FlowIpv6SegmentRoutingTlvPathTrace, error)
	// FromPbText unmarshals FlowIpv6SegmentRoutingTlvPathTrace from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6SegmentRoutingTlvPathTrace from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6SegmentRoutingTlvPathTrace from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6SegmentRoutingTlvPathTrace) Marshal() marshalFlowIpv6SegmentRoutingTlvPathTrace {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6SegmentRoutingTlvPathTrace{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6SegmentRoutingTlvPathTrace) Unmarshal() unMarshalFlowIpv6SegmentRoutingTlvPathTrace {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6SegmentRoutingTlvPathTrace{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6SegmentRoutingTlvPathTrace) ToProto() (*otg.FlowIpv6SegmentRoutingTlvPathTrace, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6SegmentRoutingTlvPathTrace) FromProto(msg *otg.FlowIpv6SegmentRoutingTlvPathTrace) (FlowIpv6SegmentRoutingTlvPathTrace, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6SegmentRoutingTlvPathTrace) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingTlvPathTrace) FromPbText(value string) error {
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

func (m *marshalflowIpv6SegmentRoutingTlvPathTrace) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingTlvPathTrace) FromYaml(value string) error {
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

func (m *marshalflowIpv6SegmentRoutingTlvPathTrace) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingTlvPathTrace) FromJson(value string) error {
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

func (obj *flowIpv6SegmentRoutingTlvPathTrace) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingTlvPathTrace) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingTlvPathTrace) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6SegmentRoutingTlvPathTrace) Clone() (FlowIpv6SegmentRoutingTlvPathTrace, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6SegmentRoutingTlvPathTrace()
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

func (obj *flowIpv6SegmentRoutingTlvPathTrace) setNil() {
	obj.interfaceIdHolder = nil
	obj.interfaceLoadHolder = nil
	obj.sessionIdHolder = nil
	obj.sequenceNumberHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv6SegmentRoutingTlvPathTrace is a TLV used for path tracing and performance monitoring within the SR domain.
type FlowIpv6SegmentRoutingTlvPathTrace interface {
	Validation
	// msg marshals FlowIpv6SegmentRoutingTlvPathTrace to protobuf object *otg.FlowIpv6SegmentRoutingTlvPathTrace
	// and doesn't set defaults
	msg() *otg.FlowIpv6SegmentRoutingTlvPathTrace
	// setMsg unmarshals FlowIpv6SegmentRoutingTlvPathTrace from protobuf object *otg.FlowIpv6SegmentRoutingTlvPathTrace
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6SegmentRoutingTlvPathTrace) FlowIpv6SegmentRoutingTlvPathTrace
	// provides marshal interface
	Marshal() marshalFlowIpv6SegmentRoutingTlvPathTrace
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6SegmentRoutingTlvPathTrace
	// validate validates FlowIpv6SegmentRoutingTlvPathTrace
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6SegmentRoutingTlvPathTrace, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// InterfaceId returns PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId, set in FlowIpv6SegmentRoutingTlvPathTrace.
	// PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId is identifier of the interface where the measurement was taken.
	InterfaceId() PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
	// SetInterfaceId assigns PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId provided by user to FlowIpv6SegmentRoutingTlvPathTrace.
	// PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId is identifier of the interface where the measurement was taken.
	SetInterfaceId(value PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) FlowIpv6SegmentRoutingTlvPathTrace
	// HasInterfaceId checks if InterfaceId has been set in FlowIpv6SegmentRoutingTlvPathTrace
	HasInterfaceId() bool
	// InterfaceLoad returns PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad, set in FlowIpv6SegmentRoutingTlvPathTrace.
	// PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad is load on the interface at the time of measurement.
	InterfaceLoad() PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
	// SetInterfaceLoad assigns PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad provided by user to FlowIpv6SegmentRoutingTlvPathTrace.
	// PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad is load on the interface at the time of measurement.
	SetInterfaceLoad(value PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) FlowIpv6SegmentRoutingTlvPathTrace
	// HasInterfaceLoad checks if InterfaceLoad has been set in FlowIpv6SegmentRoutingTlvPathTrace
	HasInterfaceLoad() bool
	// Timestamp returns string, set in FlowIpv6SegmentRoutingTlvPathTrace.
	Timestamp() string
	// SetTimestamp assigns string provided by user to FlowIpv6SegmentRoutingTlvPathTrace
	SetTimestamp(value string) FlowIpv6SegmentRoutingTlvPathTrace
	// HasTimestamp checks if Timestamp has been set in FlowIpv6SegmentRoutingTlvPathTrace
	HasTimestamp() bool
	// SessionId returns PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId, set in FlowIpv6SegmentRoutingTlvPathTrace.
	// PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId is identifier for the monitoring session.
	SessionId() PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
	// SetSessionId assigns PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId provided by user to FlowIpv6SegmentRoutingTlvPathTrace.
	// PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId is identifier for the monitoring session.
	SetSessionId(value PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId) FlowIpv6SegmentRoutingTlvPathTrace
	// HasSessionId checks if SessionId has been set in FlowIpv6SegmentRoutingTlvPathTrace
	HasSessionId() bool
	// SequenceNumber returns PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber, set in FlowIpv6SegmentRoutingTlvPathTrace.
	// PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber is sequence number for the packet within the monitoring session.
	SequenceNumber() PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
	// SetSequenceNumber assigns PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber provided by user to FlowIpv6SegmentRoutingTlvPathTrace.
	// PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber is sequence number for the packet within the monitoring session.
	SetSequenceNumber(value PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) FlowIpv6SegmentRoutingTlvPathTrace
	// HasSequenceNumber checks if SequenceNumber has been set in FlowIpv6SegmentRoutingTlvPathTrace
	HasSequenceNumber() bool
	setNil()
}

// description is TBD
// InterfaceId returns a PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
func (obj *flowIpv6SegmentRoutingTlvPathTrace) InterfaceId() PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId {
	if obj.obj.InterfaceId == nil {
		obj.obj.InterfaceId = NewPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId().msg()
	}
	if obj.interfaceIdHolder == nil {
		obj.interfaceIdHolder = &patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId{obj: obj.obj.InterfaceId}
	}
	return obj.interfaceIdHolder
}

// description is TBD
// InterfaceId returns a PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
func (obj *flowIpv6SegmentRoutingTlvPathTrace) HasInterfaceId() bool {
	return obj.obj.InterfaceId != nil
}

// description is TBD
// SetInterfaceId sets the PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId value in the FlowIpv6SegmentRoutingTlvPathTrace object
func (obj *flowIpv6SegmentRoutingTlvPathTrace) SetInterfaceId(value PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) FlowIpv6SegmentRoutingTlvPathTrace {

	obj.interfaceIdHolder = nil
	obj.obj.InterfaceId = value.msg()

	return obj
}

// description is TBD
// InterfaceLoad returns a PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
func (obj *flowIpv6SegmentRoutingTlvPathTrace) InterfaceLoad() PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad {
	if obj.obj.InterfaceLoad == nil {
		obj.obj.InterfaceLoad = NewPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad().msg()
	}
	if obj.interfaceLoadHolder == nil {
		obj.interfaceLoadHolder = &patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad{obj: obj.obj.InterfaceLoad}
	}
	return obj.interfaceLoadHolder
}

// description is TBD
// InterfaceLoad returns a PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
func (obj *flowIpv6SegmentRoutingTlvPathTrace) HasInterfaceLoad() bool {
	return obj.obj.InterfaceLoad != nil
}

// description is TBD
// SetInterfaceLoad sets the PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad value in the FlowIpv6SegmentRoutingTlvPathTrace object
func (obj *flowIpv6SegmentRoutingTlvPathTrace) SetInterfaceLoad(value PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) FlowIpv6SegmentRoutingTlvPathTrace {

	obj.interfaceLoadHolder = nil
	obj.obj.InterfaceLoad = value.msg()

	return obj
}

// Timestamp indicating when the measurement was taken.
// Timestamp returns a string
func (obj *flowIpv6SegmentRoutingTlvPathTrace) Timestamp() string {

	return *obj.obj.Timestamp

}

// Timestamp indicating when the measurement was taken.
// Timestamp returns a string
func (obj *flowIpv6SegmentRoutingTlvPathTrace) HasTimestamp() bool {
	return obj.obj.Timestamp != nil
}

// Timestamp indicating when the measurement was taken.
// SetTimestamp sets the string value in the FlowIpv6SegmentRoutingTlvPathTrace object
func (obj *flowIpv6SegmentRoutingTlvPathTrace) SetTimestamp(value string) FlowIpv6SegmentRoutingTlvPathTrace {

	obj.obj.Timestamp = &value
	return obj
}

// description is TBD
// SessionId returns a PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
func (obj *flowIpv6SegmentRoutingTlvPathTrace) SessionId() PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId {
	if obj.obj.SessionId == nil {
		obj.obj.SessionId = NewPatternFlowIpv6SegmentRoutingTlvPathTraceSessionId().msg()
	}
	if obj.sessionIdHolder == nil {
		obj.sessionIdHolder = &patternFlowIpv6SegmentRoutingTlvPathTraceSessionId{obj: obj.obj.SessionId}
	}
	return obj.sessionIdHolder
}

// description is TBD
// SessionId returns a PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
func (obj *flowIpv6SegmentRoutingTlvPathTrace) HasSessionId() bool {
	return obj.obj.SessionId != nil
}

// description is TBD
// SetSessionId sets the PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId value in the FlowIpv6SegmentRoutingTlvPathTrace object
func (obj *flowIpv6SegmentRoutingTlvPathTrace) SetSessionId(value PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId) FlowIpv6SegmentRoutingTlvPathTrace {

	obj.sessionIdHolder = nil
	obj.obj.SessionId = value.msg()

	return obj
}

// description is TBD
// SequenceNumber returns a PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
func (obj *flowIpv6SegmentRoutingTlvPathTrace) SequenceNumber() PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber {
	if obj.obj.SequenceNumber == nil {
		obj.obj.SequenceNumber = NewPatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber().msg()
	}
	if obj.sequenceNumberHolder == nil {
		obj.sequenceNumberHolder = &patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber{obj: obj.obj.SequenceNumber}
	}
	return obj.sequenceNumberHolder
}

// description is TBD
// SequenceNumber returns a PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
func (obj *flowIpv6SegmentRoutingTlvPathTrace) HasSequenceNumber() bool {
	return obj.obj.SequenceNumber != nil
}

// description is TBD
// SetSequenceNumber sets the PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber value in the FlowIpv6SegmentRoutingTlvPathTrace object
func (obj *flowIpv6SegmentRoutingTlvPathTrace) SetSequenceNumber(value PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) FlowIpv6SegmentRoutingTlvPathTrace {

	obj.sequenceNumberHolder = nil
	obj.obj.SequenceNumber = value.msg()

	return obj
}

func (obj *flowIpv6SegmentRoutingTlvPathTrace) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.InterfaceId != nil {

		obj.InterfaceId().validateObj(vObj, set_default)
	}

	if obj.obj.InterfaceLoad != nil {

		obj.InterfaceLoad().validateObj(vObj, set_default)
	}

	if obj.obj.Timestamp != nil {

		err := obj.validateHex(obj.Timestamp())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowIpv6SegmentRoutingTlvPathTrace.Timestamp"))
		}

	}

	if obj.obj.SessionId != nil {

		obj.SessionId().validateObj(vObj, set_default)
	}

	if obj.obj.SequenceNumber != nil {

		obj.SequenceNumber().validateObj(vObj, set_default)
	}

}

func (obj *flowIpv6SegmentRoutingTlvPathTrace) setDefault() {
	if obj.obj.Timestamp == nil {
		obj.SetTimestamp("00")
	}

}
