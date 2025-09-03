package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowLacpActorState *****
type flowLacpActorState struct {
	validation
	obj                   *otg.FlowLacpActorState
	marshaller            marshalFlowLacpActorState
	unMarshaller          unMarshalFlowLacpActorState
	activityHolder        PatternFlowLacpActorStateActivity
	timeoutHolder         PatternFlowLacpActorStateTimeout
	aggregationHolder     PatternFlowLacpActorStateAggregation
	synchronizationHolder PatternFlowLacpActorStateSynchronization
	collectingHolder      PatternFlowLacpActorStateCollecting
	distributingHolder    PatternFlowLacpActorStateDistributing
	defaultedHolder       PatternFlowLacpActorStateDefaulted
	expiredHolder         PatternFlowLacpActorStateExpired
}

func NewFlowLacpActorState() FlowLacpActorState {
	obj := flowLacpActorState{obj: &otg.FlowLacpActorState{}}
	obj.setDefault()
	return &obj
}

func (obj *flowLacpActorState) msg() *otg.FlowLacpActorState {
	return obj.obj
}

func (obj *flowLacpActorState) setMsg(msg *otg.FlowLacpActorState) FlowLacpActorState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowLacpActorState struct {
	obj *flowLacpActorState
}

type marshalFlowLacpActorState interface {
	// ToProto marshals FlowLacpActorState to protobuf object *otg.FlowLacpActorState
	ToProto() (*otg.FlowLacpActorState, error)
	// ToPbText marshals FlowLacpActorState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowLacpActorState to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowLacpActorState to JSON text
	ToJson() (string, error)
}

type unMarshalflowLacpActorState struct {
	obj *flowLacpActorState
}

type unMarshalFlowLacpActorState interface {
	// FromProto unmarshals FlowLacpActorState from protobuf object *otg.FlowLacpActorState
	FromProto(msg *otg.FlowLacpActorState) (FlowLacpActorState, error)
	// FromPbText unmarshals FlowLacpActorState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowLacpActorState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowLacpActorState from JSON text
	FromJson(value string) error
}

func (obj *flowLacpActorState) Marshal() marshalFlowLacpActorState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowLacpActorState{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowLacpActorState) Unmarshal() unMarshalFlowLacpActorState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowLacpActorState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowLacpActorState) ToProto() (*otg.FlowLacpActorState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowLacpActorState) FromProto(msg *otg.FlowLacpActorState) (FlowLacpActorState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowLacpActorState) ToPbText() (string, error) {
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

func (m *unMarshalflowLacpActorState) FromPbText(value string) error {
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

func (m *marshalflowLacpActorState) ToYaml() (string, error) {
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

func (m *unMarshalflowLacpActorState) FromYaml(value string) error {
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

func (m *marshalflowLacpActorState) ToJson() (string, error) {
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

func (m *unMarshalflowLacpActorState) FromJson(value string) error {
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

func (obj *flowLacpActorState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowLacpActorState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowLacpActorState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowLacpActorState) Clone() (FlowLacpActorState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowLacpActorState()
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

func (obj *flowLacpActorState) setNil() {
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

// FlowLacpActorState is this field indicates the Actor's state.
type FlowLacpActorState interface {
	Validation
	// msg marshals FlowLacpActorState to protobuf object *otg.FlowLacpActorState
	// and doesn't set defaults
	msg() *otg.FlowLacpActorState
	// setMsg unmarshals FlowLacpActorState from protobuf object *otg.FlowLacpActorState
	// and doesn't set defaults
	setMsg(*otg.FlowLacpActorState) FlowLacpActorState
	// provides marshal interface
	Marshal() marshalFlowLacpActorState
	// provides unmarshal interface
	Unmarshal() unMarshalFlowLacpActorState
	// validate validates FlowLacpActorState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowLacpActorState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Activity returns PatternFlowLacpActorStateActivity, set in FlowLacpActorState.
	// PatternFlowLacpActorStateActivity is lACP Activity (1=Active, 0=Passive)
	Activity() PatternFlowLacpActorStateActivity
	// SetActivity assigns PatternFlowLacpActorStateActivity provided by user to FlowLacpActorState.
	// PatternFlowLacpActorStateActivity is lACP Activity (1=Active, 0=Passive)
	SetActivity(value PatternFlowLacpActorStateActivity) FlowLacpActorState
	// HasActivity checks if Activity has been set in FlowLacpActorState
	HasActivity() bool
	// Timeout returns PatternFlowLacpActorStateTimeout, set in FlowLacpActorState.
	// PatternFlowLacpActorStateTimeout is lACP Timeout (1=Fast timeout, 0=Slow timeout)
	Timeout() PatternFlowLacpActorStateTimeout
	// SetTimeout assigns PatternFlowLacpActorStateTimeout provided by user to FlowLacpActorState.
	// PatternFlowLacpActorStateTimeout is lACP Timeout (1=Fast timeout, 0=Slow timeout)
	SetTimeout(value PatternFlowLacpActorStateTimeout) FlowLacpActorState
	// HasTimeout checks if Timeout has been set in FlowLacpActorState
	HasTimeout() bool
	// Aggregation returns PatternFlowLacpActorStateAggregation, set in FlowLacpActorState.
	// PatternFlowLacpActorStateAggregation is aggregation (1=Aggregatable, 0=Individual)
	Aggregation() PatternFlowLacpActorStateAggregation
	// SetAggregation assigns PatternFlowLacpActorStateAggregation provided by user to FlowLacpActorState.
	// PatternFlowLacpActorStateAggregation is aggregation (1=Aggregatable, 0=Individual)
	SetAggregation(value PatternFlowLacpActorStateAggregation) FlowLacpActorState
	// HasAggregation checks if Aggregation has been set in FlowLacpActorState
	HasAggregation() bool
	// Synchronization returns PatternFlowLacpActorStateSynchronization, set in FlowLacpActorState.
	// PatternFlowLacpActorStateSynchronization is synchronization (1=In Sync, 0=Out of Sync)
	Synchronization() PatternFlowLacpActorStateSynchronization
	// SetSynchronization assigns PatternFlowLacpActorStateSynchronization provided by user to FlowLacpActorState.
	// PatternFlowLacpActorStateSynchronization is synchronization (1=In Sync, 0=Out of Sync)
	SetSynchronization(value PatternFlowLacpActorStateSynchronization) FlowLacpActorState
	// HasSynchronization checks if Synchronization has been set in FlowLacpActorState
	HasSynchronization() bool
	// Collecting returns PatternFlowLacpActorStateCollecting, set in FlowLacpActorState.
	// PatternFlowLacpActorStateCollecting is collecting (1=Enabled, 0=Disabled)
	Collecting() PatternFlowLacpActorStateCollecting
	// SetCollecting assigns PatternFlowLacpActorStateCollecting provided by user to FlowLacpActorState.
	// PatternFlowLacpActorStateCollecting is collecting (1=Enabled, 0=Disabled)
	SetCollecting(value PatternFlowLacpActorStateCollecting) FlowLacpActorState
	// HasCollecting checks if Collecting has been set in FlowLacpActorState
	HasCollecting() bool
	// Distributing returns PatternFlowLacpActorStateDistributing, set in FlowLacpActorState.
	// PatternFlowLacpActorStateDistributing is distributing (1=Enabled, 0=Disabled)
	Distributing() PatternFlowLacpActorStateDistributing
	// SetDistributing assigns PatternFlowLacpActorStateDistributing provided by user to FlowLacpActorState.
	// PatternFlowLacpActorStateDistributing is distributing (1=Enabled, 0=Disabled)
	SetDistributing(value PatternFlowLacpActorStateDistributing) FlowLacpActorState
	// HasDistributing checks if Distributing has been set in FlowLacpActorState
	HasDistributing() bool
	// Defaulted returns PatternFlowLacpActorStateDefaulted, set in FlowLacpActorState.
	// PatternFlowLacpActorStateDefaulted is defaulted (1=Using defaulted partner info, 0=Using received partner info)
	Defaulted() PatternFlowLacpActorStateDefaulted
	// SetDefaulted assigns PatternFlowLacpActorStateDefaulted provided by user to FlowLacpActorState.
	// PatternFlowLacpActorStateDefaulted is defaulted (1=Using defaulted partner info, 0=Using received partner info)
	SetDefaulted(value PatternFlowLacpActorStateDefaulted) FlowLacpActorState
	// HasDefaulted checks if Defaulted has been set in FlowLacpActorState
	HasDefaulted() bool
	// Expired returns PatternFlowLacpActorStateExpired, set in FlowLacpActorState.
	// PatternFlowLacpActorStateExpired is expired (1=Expired, 0=Not Expired)
	Expired() PatternFlowLacpActorStateExpired
	// SetExpired assigns PatternFlowLacpActorStateExpired provided by user to FlowLacpActorState.
	// PatternFlowLacpActorStateExpired is expired (1=Expired, 0=Not Expired)
	SetExpired(value PatternFlowLacpActorStateExpired) FlowLacpActorState
	// HasExpired checks if Expired has been set in FlowLacpActorState
	HasExpired() bool
	setNil()
}

// description is TBD
// Activity returns a PatternFlowLacpActorStateActivity
func (obj *flowLacpActorState) Activity() PatternFlowLacpActorStateActivity {
	if obj.obj.Activity == nil {
		obj.obj.Activity = NewPatternFlowLacpActorStateActivity().msg()
	}
	if obj.activityHolder == nil {
		obj.activityHolder = &patternFlowLacpActorStateActivity{obj: obj.obj.Activity}
	}
	return obj.activityHolder
}

// description is TBD
// Activity returns a PatternFlowLacpActorStateActivity
func (obj *flowLacpActorState) HasActivity() bool {
	return obj.obj.Activity != nil
}

// description is TBD
// SetActivity sets the PatternFlowLacpActorStateActivity value in the FlowLacpActorState object
func (obj *flowLacpActorState) SetActivity(value PatternFlowLacpActorStateActivity) FlowLacpActorState {

	obj.activityHolder = nil
	obj.obj.Activity = value.msg()

	return obj
}

// description is TBD
// Timeout returns a PatternFlowLacpActorStateTimeout
func (obj *flowLacpActorState) Timeout() PatternFlowLacpActorStateTimeout {
	if obj.obj.Timeout == nil {
		obj.obj.Timeout = NewPatternFlowLacpActorStateTimeout().msg()
	}
	if obj.timeoutHolder == nil {
		obj.timeoutHolder = &patternFlowLacpActorStateTimeout{obj: obj.obj.Timeout}
	}
	return obj.timeoutHolder
}

// description is TBD
// Timeout returns a PatternFlowLacpActorStateTimeout
func (obj *flowLacpActorState) HasTimeout() bool {
	return obj.obj.Timeout != nil
}

// description is TBD
// SetTimeout sets the PatternFlowLacpActorStateTimeout value in the FlowLacpActorState object
func (obj *flowLacpActorState) SetTimeout(value PatternFlowLacpActorStateTimeout) FlowLacpActorState {

	obj.timeoutHolder = nil
	obj.obj.Timeout = value.msg()

	return obj
}

// description is TBD
// Aggregation returns a PatternFlowLacpActorStateAggregation
func (obj *flowLacpActorState) Aggregation() PatternFlowLacpActorStateAggregation {
	if obj.obj.Aggregation == nil {
		obj.obj.Aggregation = NewPatternFlowLacpActorStateAggregation().msg()
	}
	if obj.aggregationHolder == nil {
		obj.aggregationHolder = &patternFlowLacpActorStateAggregation{obj: obj.obj.Aggregation}
	}
	return obj.aggregationHolder
}

// description is TBD
// Aggregation returns a PatternFlowLacpActorStateAggregation
func (obj *flowLacpActorState) HasAggregation() bool {
	return obj.obj.Aggregation != nil
}

// description is TBD
// SetAggregation sets the PatternFlowLacpActorStateAggregation value in the FlowLacpActorState object
func (obj *flowLacpActorState) SetAggregation(value PatternFlowLacpActorStateAggregation) FlowLacpActorState {

	obj.aggregationHolder = nil
	obj.obj.Aggregation = value.msg()

	return obj
}

// description is TBD
// Synchronization returns a PatternFlowLacpActorStateSynchronization
func (obj *flowLacpActorState) Synchronization() PatternFlowLacpActorStateSynchronization {
	if obj.obj.Synchronization == nil {
		obj.obj.Synchronization = NewPatternFlowLacpActorStateSynchronization().msg()
	}
	if obj.synchronizationHolder == nil {
		obj.synchronizationHolder = &patternFlowLacpActorStateSynchronization{obj: obj.obj.Synchronization}
	}
	return obj.synchronizationHolder
}

// description is TBD
// Synchronization returns a PatternFlowLacpActorStateSynchronization
func (obj *flowLacpActorState) HasSynchronization() bool {
	return obj.obj.Synchronization != nil
}

// description is TBD
// SetSynchronization sets the PatternFlowLacpActorStateSynchronization value in the FlowLacpActorState object
func (obj *flowLacpActorState) SetSynchronization(value PatternFlowLacpActorStateSynchronization) FlowLacpActorState {

	obj.synchronizationHolder = nil
	obj.obj.Synchronization = value.msg()

	return obj
}

// description is TBD
// Collecting returns a PatternFlowLacpActorStateCollecting
func (obj *flowLacpActorState) Collecting() PatternFlowLacpActorStateCollecting {
	if obj.obj.Collecting == nil {
		obj.obj.Collecting = NewPatternFlowLacpActorStateCollecting().msg()
	}
	if obj.collectingHolder == nil {
		obj.collectingHolder = &patternFlowLacpActorStateCollecting{obj: obj.obj.Collecting}
	}
	return obj.collectingHolder
}

// description is TBD
// Collecting returns a PatternFlowLacpActorStateCollecting
func (obj *flowLacpActorState) HasCollecting() bool {
	return obj.obj.Collecting != nil
}

// description is TBD
// SetCollecting sets the PatternFlowLacpActorStateCollecting value in the FlowLacpActorState object
func (obj *flowLacpActorState) SetCollecting(value PatternFlowLacpActorStateCollecting) FlowLacpActorState {

	obj.collectingHolder = nil
	obj.obj.Collecting = value.msg()

	return obj
}

// description is TBD
// Distributing returns a PatternFlowLacpActorStateDistributing
func (obj *flowLacpActorState) Distributing() PatternFlowLacpActorStateDistributing {
	if obj.obj.Distributing == nil {
		obj.obj.Distributing = NewPatternFlowLacpActorStateDistributing().msg()
	}
	if obj.distributingHolder == nil {
		obj.distributingHolder = &patternFlowLacpActorStateDistributing{obj: obj.obj.Distributing}
	}
	return obj.distributingHolder
}

// description is TBD
// Distributing returns a PatternFlowLacpActorStateDistributing
func (obj *flowLacpActorState) HasDistributing() bool {
	return obj.obj.Distributing != nil
}

// description is TBD
// SetDistributing sets the PatternFlowLacpActorStateDistributing value in the FlowLacpActorState object
func (obj *flowLacpActorState) SetDistributing(value PatternFlowLacpActorStateDistributing) FlowLacpActorState {

	obj.distributingHolder = nil
	obj.obj.Distributing = value.msg()

	return obj
}

// description is TBD
// Defaulted returns a PatternFlowLacpActorStateDefaulted
func (obj *flowLacpActorState) Defaulted() PatternFlowLacpActorStateDefaulted {
	if obj.obj.Defaulted == nil {
		obj.obj.Defaulted = NewPatternFlowLacpActorStateDefaulted().msg()
	}
	if obj.defaultedHolder == nil {
		obj.defaultedHolder = &patternFlowLacpActorStateDefaulted{obj: obj.obj.Defaulted}
	}
	return obj.defaultedHolder
}

// description is TBD
// Defaulted returns a PatternFlowLacpActorStateDefaulted
func (obj *flowLacpActorState) HasDefaulted() bool {
	return obj.obj.Defaulted != nil
}

// description is TBD
// SetDefaulted sets the PatternFlowLacpActorStateDefaulted value in the FlowLacpActorState object
func (obj *flowLacpActorState) SetDefaulted(value PatternFlowLacpActorStateDefaulted) FlowLacpActorState {

	obj.defaultedHolder = nil
	obj.obj.Defaulted = value.msg()

	return obj
}

// description is TBD
// Expired returns a PatternFlowLacpActorStateExpired
func (obj *flowLacpActorState) Expired() PatternFlowLacpActorStateExpired {
	if obj.obj.Expired == nil {
		obj.obj.Expired = NewPatternFlowLacpActorStateExpired().msg()
	}
	if obj.expiredHolder == nil {
		obj.expiredHolder = &patternFlowLacpActorStateExpired{obj: obj.obj.Expired}
	}
	return obj.expiredHolder
}

// description is TBD
// Expired returns a PatternFlowLacpActorStateExpired
func (obj *flowLacpActorState) HasExpired() bool {
	return obj.obj.Expired != nil
}

// description is TBD
// SetExpired sets the PatternFlowLacpActorStateExpired value in the FlowLacpActorState object
func (obj *flowLacpActorState) SetExpired(value PatternFlowLacpActorStateExpired) FlowLacpActorState {

	obj.expiredHolder = nil
	obj.obj.Expired = value.msg()

	return obj
}

func (obj *flowLacpActorState) validateObj(vObj *validation, set_default bool) {
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

func (obj *flowLacpActorState) setDefault() {

}
