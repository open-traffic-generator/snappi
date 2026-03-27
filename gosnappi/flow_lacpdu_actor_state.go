package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowLacpduActorState *****
type flowLacpduActorState struct {
	validation
	obj                   *otg.FlowLacpduActorState
	marshaller            marshalFlowLacpduActorState
	unMarshaller          unMarshalFlowLacpduActorState
	activityHolder        PatternFlowLacpduActorStateActivity
	timeoutHolder         PatternFlowLacpduActorStateTimeout
	aggregationHolder     PatternFlowLacpduActorStateAggregation
	synchronizationHolder PatternFlowLacpduActorStateSynchronization
	collectingHolder      PatternFlowLacpduActorStateCollecting
	distributingHolder    PatternFlowLacpduActorStateDistributing
	defaultedHolder       PatternFlowLacpduActorStateDefaulted
	expiredHolder         PatternFlowLacpduActorStateExpired
}

func NewFlowLacpduActorState() FlowLacpduActorState {
	obj := flowLacpduActorState{obj: &otg.FlowLacpduActorState{}}
	obj.setDefault()
	return &obj
}

func (obj *flowLacpduActorState) msg() *otg.FlowLacpduActorState {
	return obj.obj
}

func (obj *flowLacpduActorState) setMsg(msg *otg.FlowLacpduActorState) FlowLacpduActorState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowLacpduActorState struct {
	obj *flowLacpduActorState
}

type marshalFlowLacpduActorState interface {
	// ToProto marshals FlowLacpduActorState to protobuf object *otg.FlowLacpduActorState
	ToProto() (*otg.FlowLacpduActorState, error)
	// ToPbText marshals FlowLacpduActorState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowLacpduActorState to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowLacpduActorState to JSON text
	ToJson() (string, error)
}

type unMarshalflowLacpduActorState struct {
	obj *flowLacpduActorState
}

type unMarshalFlowLacpduActorState interface {
	// FromProto unmarshals FlowLacpduActorState from protobuf object *otg.FlowLacpduActorState
	FromProto(msg *otg.FlowLacpduActorState) (FlowLacpduActorState, error)
	// FromPbText unmarshals FlowLacpduActorState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowLacpduActorState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowLacpduActorState from JSON text
	FromJson(value string) error
}

func (obj *flowLacpduActorState) Marshal() marshalFlowLacpduActorState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowLacpduActorState{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowLacpduActorState) Unmarshal() unMarshalFlowLacpduActorState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowLacpduActorState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowLacpduActorState) ToProto() (*otg.FlowLacpduActorState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowLacpduActorState) FromProto(msg *otg.FlowLacpduActorState) (FlowLacpduActorState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowLacpduActorState) ToPbText() (string, error) {
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

func (m *unMarshalflowLacpduActorState) FromPbText(value string) error {
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

func (m *marshalflowLacpduActorState) ToYaml() (string, error) {
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

func (m *unMarshalflowLacpduActorState) FromYaml(value string) error {
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

func (m *marshalflowLacpduActorState) ToJson() (string, error) {
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

func (m *unMarshalflowLacpduActorState) FromJson(value string) error {
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

func (obj *flowLacpduActorState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowLacpduActorState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowLacpduActorState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowLacpduActorState) Clone() (FlowLacpduActorState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowLacpduActorState()
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

func (obj *flowLacpduActorState) setNil() {
	obj.activityHolder = nil
	obj.timeoutHolder = nil
	obj.aggregationHolder = nil
	obj.synchronizationHolder = nil
	obj.collectingHolder = nil
	obj.distributingHolder = nil
	obj.defaultedHolder = nil
	obj.expiredHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowLacpduActorState is this field indicates the Actor's state.
type FlowLacpduActorState interface {
	Validation
	// msg marshals FlowLacpduActorState to protobuf object *otg.FlowLacpduActorState
	// and doesn't set defaults
	msg() *otg.FlowLacpduActorState
	// setMsg unmarshals FlowLacpduActorState from protobuf object *otg.FlowLacpduActorState
	// and doesn't set defaults
	setMsg(*otg.FlowLacpduActorState) FlowLacpduActorState
	// provides marshal interface
	Marshal() marshalFlowLacpduActorState
	// provides unmarshal interface
	Unmarshal() unMarshalFlowLacpduActorState
	// validate validates FlowLacpduActorState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowLacpduActorState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Activity returns PatternFlowLacpduActorStateActivity, set in FlowLacpduActorState.
	// PatternFlowLacpduActorStateActivity is lACP Activity (1=Active, 0=Passive)
	Activity() PatternFlowLacpduActorStateActivity
	// SetActivity assigns PatternFlowLacpduActorStateActivity provided by user to FlowLacpduActorState.
	// PatternFlowLacpduActorStateActivity is lACP Activity (1=Active, 0=Passive)
	SetActivity(value PatternFlowLacpduActorStateActivity) FlowLacpduActorState
	// HasActivity checks if Activity has been set in FlowLacpduActorState
	HasActivity() bool
	// Timeout returns PatternFlowLacpduActorStateTimeout, set in FlowLacpduActorState.
	// PatternFlowLacpduActorStateTimeout is lACP Timeout (1=Fast timeout, 0=Slow timeout)
	Timeout() PatternFlowLacpduActorStateTimeout
	// SetTimeout assigns PatternFlowLacpduActorStateTimeout provided by user to FlowLacpduActorState.
	// PatternFlowLacpduActorStateTimeout is lACP Timeout (1=Fast timeout, 0=Slow timeout)
	SetTimeout(value PatternFlowLacpduActorStateTimeout) FlowLacpduActorState
	// HasTimeout checks if Timeout has been set in FlowLacpduActorState
	HasTimeout() bool
	// Aggregation returns PatternFlowLacpduActorStateAggregation, set in FlowLacpduActorState.
	// PatternFlowLacpduActorStateAggregation is aggregation (1=Aggregatable, 0=Individual)
	Aggregation() PatternFlowLacpduActorStateAggregation
	// SetAggregation assigns PatternFlowLacpduActorStateAggregation provided by user to FlowLacpduActorState.
	// PatternFlowLacpduActorStateAggregation is aggregation (1=Aggregatable, 0=Individual)
	SetAggregation(value PatternFlowLacpduActorStateAggregation) FlowLacpduActorState
	// HasAggregation checks if Aggregation has been set in FlowLacpduActorState
	HasAggregation() bool
	// Synchronization returns PatternFlowLacpduActorStateSynchronization, set in FlowLacpduActorState.
	// PatternFlowLacpduActorStateSynchronization is synchronization (1=In Sync, 0=Out of Sync)
	Synchronization() PatternFlowLacpduActorStateSynchronization
	// SetSynchronization assigns PatternFlowLacpduActorStateSynchronization provided by user to FlowLacpduActorState.
	// PatternFlowLacpduActorStateSynchronization is synchronization (1=In Sync, 0=Out of Sync)
	SetSynchronization(value PatternFlowLacpduActorStateSynchronization) FlowLacpduActorState
	// HasSynchronization checks if Synchronization has been set in FlowLacpduActorState
	HasSynchronization() bool
	// Collecting returns PatternFlowLacpduActorStateCollecting, set in FlowLacpduActorState.
	// PatternFlowLacpduActorStateCollecting is collecting (1=Enabled, 0=Disabled)
	Collecting() PatternFlowLacpduActorStateCollecting
	// SetCollecting assigns PatternFlowLacpduActorStateCollecting provided by user to FlowLacpduActorState.
	// PatternFlowLacpduActorStateCollecting is collecting (1=Enabled, 0=Disabled)
	SetCollecting(value PatternFlowLacpduActorStateCollecting) FlowLacpduActorState
	// HasCollecting checks if Collecting has been set in FlowLacpduActorState
	HasCollecting() bool
	// Distributing returns PatternFlowLacpduActorStateDistributing, set in FlowLacpduActorState.
	// PatternFlowLacpduActorStateDistributing is distributing (1=Enabled, 0=Disabled)
	Distributing() PatternFlowLacpduActorStateDistributing
	// SetDistributing assigns PatternFlowLacpduActorStateDistributing provided by user to FlowLacpduActorState.
	// PatternFlowLacpduActorStateDistributing is distributing (1=Enabled, 0=Disabled)
	SetDistributing(value PatternFlowLacpduActorStateDistributing) FlowLacpduActorState
	// HasDistributing checks if Distributing has been set in FlowLacpduActorState
	HasDistributing() bool
	// Defaulted returns PatternFlowLacpduActorStateDefaulted, set in FlowLacpduActorState.
	// PatternFlowLacpduActorStateDefaulted is defaulted (1=Using defaulted partner info, 0=Using received partner info)
	Defaulted() PatternFlowLacpduActorStateDefaulted
	// SetDefaulted assigns PatternFlowLacpduActorStateDefaulted provided by user to FlowLacpduActorState.
	// PatternFlowLacpduActorStateDefaulted is defaulted (1=Using defaulted partner info, 0=Using received partner info)
	SetDefaulted(value PatternFlowLacpduActorStateDefaulted) FlowLacpduActorState
	// HasDefaulted checks if Defaulted has been set in FlowLacpduActorState
	HasDefaulted() bool
	// Expired returns PatternFlowLacpduActorStateExpired, set in FlowLacpduActorState.
	// PatternFlowLacpduActorStateExpired is expired (1=Expired, 0=Not Expired)
	Expired() PatternFlowLacpduActorStateExpired
	// SetExpired assigns PatternFlowLacpduActorStateExpired provided by user to FlowLacpduActorState.
	// PatternFlowLacpduActorStateExpired is expired (1=Expired, 0=Not Expired)
	SetExpired(value PatternFlowLacpduActorStateExpired) FlowLacpduActorState
	// HasExpired checks if Expired has been set in FlowLacpduActorState
	HasExpired() bool
	setNil()
}

// description is TBD
// Activity returns a PatternFlowLacpduActorStateActivity
func (obj *flowLacpduActorState) Activity() PatternFlowLacpduActorStateActivity {
	if obj.obj.Activity == nil {
		obj.obj.Activity = NewPatternFlowLacpduActorStateActivity().msg()
	}
	if obj.activityHolder == nil {
		obj.activityHolder = &patternFlowLacpduActorStateActivity{obj: obj.obj.Activity}
	}
	return obj.activityHolder
}

// description is TBD
// Activity returns a PatternFlowLacpduActorStateActivity
func (obj *flowLacpduActorState) HasActivity() bool {
	return obj.obj.Activity != nil
}

// description is TBD
// SetActivity sets the PatternFlowLacpduActorStateActivity value in the FlowLacpduActorState object
func (obj *flowLacpduActorState) SetActivity(value PatternFlowLacpduActorStateActivity) FlowLacpduActorState {

	obj.activityHolder = nil
	obj.obj.Activity = value.msg()

	return obj
}

// description is TBD
// Timeout returns a PatternFlowLacpduActorStateTimeout
func (obj *flowLacpduActorState) Timeout() PatternFlowLacpduActorStateTimeout {
	if obj.obj.Timeout == nil {
		obj.obj.Timeout = NewPatternFlowLacpduActorStateTimeout().msg()
	}
	if obj.timeoutHolder == nil {
		obj.timeoutHolder = &patternFlowLacpduActorStateTimeout{obj: obj.obj.Timeout}
	}
	return obj.timeoutHolder
}

// description is TBD
// Timeout returns a PatternFlowLacpduActorStateTimeout
func (obj *flowLacpduActorState) HasTimeout() bool {
	return obj.obj.Timeout != nil
}

// description is TBD
// SetTimeout sets the PatternFlowLacpduActorStateTimeout value in the FlowLacpduActorState object
func (obj *flowLacpduActorState) SetTimeout(value PatternFlowLacpduActorStateTimeout) FlowLacpduActorState {

	obj.timeoutHolder = nil
	obj.obj.Timeout = value.msg()

	return obj
}

// description is TBD
// Aggregation returns a PatternFlowLacpduActorStateAggregation
func (obj *flowLacpduActorState) Aggregation() PatternFlowLacpduActorStateAggregation {
	if obj.obj.Aggregation == nil {
		obj.obj.Aggregation = NewPatternFlowLacpduActorStateAggregation().msg()
	}
	if obj.aggregationHolder == nil {
		obj.aggregationHolder = &patternFlowLacpduActorStateAggregation{obj: obj.obj.Aggregation}
	}
	return obj.aggregationHolder
}

// description is TBD
// Aggregation returns a PatternFlowLacpduActorStateAggregation
func (obj *flowLacpduActorState) HasAggregation() bool {
	return obj.obj.Aggregation != nil
}

// description is TBD
// SetAggregation sets the PatternFlowLacpduActorStateAggregation value in the FlowLacpduActorState object
func (obj *flowLacpduActorState) SetAggregation(value PatternFlowLacpduActorStateAggregation) FlowLacpduActorState {

	obj.aggregationHolder = nil
	obj.obj.Aggregation = value.msg()

	return obj
}

// description is TBD
// Synchronization returns a PatternFlowLacpduActorStateSynchronization
func (obj *flowLacpduActorState) Synchronization() PatternFlowLacpduActorStateSynchronization {
	if obj.obj.Synchronization == nil {
		obj.obj.Synchronization = NewPatternFlowLacpduActorStateSynchronization().msg()
	}
	if obj.synchronizationHolder == nil {
		obj.synchronizationHolder = &patternFlowLacpduActorStateSynchronization{obj: obj.obj.Synchronization}
	}
	return obj.synchronizationHolder
}

// description is TBD
// Synchronization returns a PatternFlowLacpduActorStateSynchronization
func (obj *flowLacpduActorState) HasSynchronization() bool {
	return obj.obj.Synchronization != nil
}

// description is TBD
// SetSynchronization sets the PatternFlowLacpduActorStateSynchronization value in the FlowLacpduActorState object
func (obj *flowLacpduActorState) SetSynchronization(value PatternFlowLacpduActorStateSynchronization) FlowLacpduActorState {

	obj.synchronizationHolder = nil
	obj.obj.Synchronization = value.msg()

	return obj
}

// description is TBD
// Collecting returns a PatternFlowLacpduActorStateCollecting
func (obj *flowLacpduActorState) Collecting() PatternFlowLacpduActorStateCollecting {
	if obj.obj.Collecting == nil {
		obj.obj.Collecting = NewPatternFlowLacpduActorStateCollecting().msg()
	}
	if obj.collectingHolder == nil {
		obj.collectingHolder = &patternFlowLacpduActorStateCollecting{obj: obj.obj.Collecting}
	}
	return obj.collectingHolder
}

// description is TBD
// Collecting returns a PatternFlowLacpduActorStateCollecting
func (obj *flowLacpduActorState) HasCollecting() bool {
	return obj.obj.Collecting != nil
}

// description is TBD
// SetCollecting sets the PatternFlowLacpduActorStateCollecting value in the FlowLacpduActorState object
func (obj *flowLacpduActorState) SetCollecting(value PatternFlowLacpduActorStateCollecting) FlowLacpduActorState {

	obj.collectingHolder = nil
	obj.obj.Collecting = value.msg()

	return obj
}

// description is TBD
// Distributing returns a PatternFlowLacpduActorStateDistributing
func (obj *flowLacpduActorState) Distributing() PatternFlowLacpduActorStateDistributing {
	if obj.obj.Distributing == nil {
		obj.obj.Distributing = NewPatternFlowLacpduActorStateDistributing().msg()
	}
	if obj.distributingHolder == nil {
		obj.distributingHolder = &patternFlowLacpduActorStateDistributing{obj: obj.obj.Distributing}
	}
	return obj.distributingHolder
}

// description is TBD
// Distributing returns a PatternFlowLacpduActorStateDistributing
func (obj *flowLacpduActorState) HasDistributing() bool {
	return obj.obj.Distributing != nil
}

// description is TBD
// SetDistributing sets the PatternFlowLacpduActorStateDistributing value in the FlowLacpduActorState object
func (obj *flowLacpduActorState) SetDistributing(value PatternFlowLacpduActorStateDistributing) FlowLacpduActorState {

	obj.distributingHolder = nil
	obj.obj.Distributing = value.msg()

	return obj
}

// description is TBD
// Defaulted returns a PatternFlowLacpduActorStateDefaulted
func (obj *flowLacpduActorState) Defaulted() PatternFlowLacpduActorStateDefaulted {
	if obj.obj.Defaulted == nil {
		obj.obj.Defaulted = NewPatternFlowLacpduActorStateDefaulted().msg()
	}
	if obj.defaultedHolder == nil {
		obj.defaultedHolder = &patternFlowLacpduActorStateDefaulted{obj: obj.obj.Defaulted}
	}
	return obj.defaultedHolder
}

// description is TBD
// Defaulted returns a PatternFlowLacpduActorStateDefaulted
func (obj *flowLacpduActorState) HasDefaulted() bool {
	return obj.obj.Defaulted != nil
}

// description is TBD
// SetDefaulted sets the PatternFlowLacpduActorStateDefaulted value in the FlowLacpduActorState object
func (obj *flowLacpduActorState) SetDefaulted(value PatternFlowLacpduActorStateDefaulted) FlowLacpduActorState {

	obj.defaultedHolder = nil
	obj.obj.Defaulted = value.msg()

	return obj
}

// description is TBD
// Expired returns a PatternFlowLacpduActorStateExpired
func (obj *flowLacpduActorState) Expired() PatternFlowLacpduActorStateExpired {
	if obj.obj.Expired == nil {
		obj.obj.Expired = NewPatternFlowLacpduActorStateExpired().msg()
	}
	if obj.expiredHolder == nil {
		obj.expiredHolder = &patternFlowLacpduActorStateExpired{obj: obj.obj.Expired}
	}
	return obj.expiredHolder
}

// description is TBD
// Expired returns a PatternFlowLacpduActorStateExpired
func (obj *flowLacpduActorState) HasExpired() bool {
	return obj.obj.Expired != nil
}

// description is TBD
// SetExpired sets the PatternFlowLacpduActorStateExpired value in the FlowLacpduActorState object
func (obj *flowLacpduActorState) SetExpired(value PatternFlowLacpduActorStateExpired) FlowLacpduActorState {

	obj.expiredHolder = nil
	obj.obj.Expired = value.msg()

	return obj
}

func (obj *flowLacpduActorState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Activity != nil {

		obj.Activity().validateObj(vObj, set_default)
	}

	if obj.obj.Timeout != nil {

		obj.Timeout().validateObj(vObj, set_default)
	}

	if obj.obj.Aggregation != nil {

		obj.Aggregation().validateObj(vObj, set_default)
	}

	if obj.obj.Synchronization != nil {

		obj.Synchronization().validateObj(vObj, set_default)
	}

	if obj.obj.Collecting != nil {

		obj.Collecting().validateObj(vObj, set_default)
	}

	if obj.obj.Distributing != nil {

		obj.Distributing().validateObj(vObj, set_default)
	}

	if obj.obj.Defaulted != nil {

		obj.Defaulted().validateObj(vObj, set_default)
	}

	if obj.obj.Expired != nil {

		obj.Expired().validateObj(vObj, set_default)
	}

}

func (obj *flowLacpduActorState) setDefault() {

}
