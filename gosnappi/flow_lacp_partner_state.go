package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowLacpPartnerState *****
type flowLacpPartnerState struct {
	validation
	obj                   *otg.FlowLacpPartnerState
	marshaller            marshalFlowLacpPartnerState
	unMarshaller          unMarshalFlowLacpPartnerState
	activityHolder        PatternFlowLacpPartnerStateActivity
	timeoutHolder         PatternFlowLacpPartnerStateTimeout
	aggregationHolder     PatternFlowLacpPartnerStateAggregation
	synchronizationHolder PatternFlowLacpPartnerStateSynchronization
	collectingHolder      PatternFlowLacpPartnerStateCollecting
	distributingHolder    PatternFlowLacpPartnerStateDistributing
	defaultedHolder       PatternFlowLacpPartnerStateDefaulted
	expiredHolder         PatternFlowLacpPartnerStateExpired
}

func NewFlowLacpPartnerState() FlowLacpPartnerState {
	obj := flowLacpPartnerState{obj: &otg.FlowLacpPartnerState{}}
	obj.setDefault()
	return &obj
}

func (obj *flowLacpPartnerState) msg() *otg.FlowLacpPartnerState {
	return obj.obj
}

func (obj *flowLacpPartnerState) setMsg(msg *otg.FlowLacpPartnerState) FlowLacpPartnerState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowLacpPartnerState struct {
	obj *flowLacpPartnerState
}

type marshalFlowLacpPartnerState interface {
	// ToProto marshals FlowLacpPartnerState to protobuf object *otg.FlowLacpPartnerState
	ToProto() (*otg.FlowLacpPartnerState, error)
	// ToPbText marshals FlowLacpPartnerState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowLacpPartnerState to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowLacpPartnerState to JSON text
	ToJson() (string, error)
}

type unMarshalflowLacpPartnerState struct {
	obj *flowLacpPartnerState
}

type unMarshalFlowLacpPartnerState interface {
	// FromProto unmarshals FlowLacpPartnerState from protobuf object *otg.FlowLacpPartnerState
	FromProto(msg *otg.FlowLacpPartnerState) (FlowLacpPartnerState, error)
	// FromPbText unmarshals FlowLacpPartnerState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowLacpPartnerState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowLacpPartnerState from JSON text
	FromJson(value string) error
}

func (obj *flowLacpPartnerState) Marshal() marshalFlowLacpPartnerState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowLacpPartnerState{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowLacpPartnerState) Unmarshal() unMarshalFlowLacpPartnerState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowLacpPartnerState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowLacpPartnerState) ToProto() (*otg.FlowLacpPartnerState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowLacpPartnerState) FromProto(msg *otg.FlowLacpPartnerState) (FlowLacpPartnerState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowLacpPartnerState) ToPbText() (string, error) {
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

func (m *unMarshalflowLacpPartnerState) FromPbText(value string) error {
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

func (m *marshalflowLacpPartnerState) ToYaml() (string, error) {
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

func (m *unMarshalflowLacpPartnerState) FromYaml(value string) error {
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

func (m *marshalflowLacpPartnerState) ToJson() (string, error) {
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

func (m *unMarshalflowLacpPartnerState) FromJson(value string) error {
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

func (obj *flowLacpPartnerState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowLacpPartnerState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowLacpPartnerState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowLacpPartnerState) Clone() (FlowLacpPartnerState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowLacpPartnerState()
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

func (obj *flowLacpPartnerState) setNil() {
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

// FlowLacpPartnerState is this field indicates the Partner's state.
type FlowLacpPartnerState interface {
	Validation
	// msg marshals FlowLacpPartnerState to protobuf object *otg.FlowLacpPartnerState
	// and doesn't set defaults
	msg() *otg.FlowLacpPartnerState
	// setMsg unmarshals FlowLacpPartnerState from protobuf object *otg.FlowLacpPartnerState
	// and doesn't set defaults
	setMsg(*otg.FlowLacpPartnerState) FlowLacpPartnerState
	// provides marshal interface
	Marshal() marshalFlowLacpPartnerState
	// provides unmarshal interface
	Unmarshal() unMarshalFlowLacpPartnerState
	// validate validates FlowLacpPartnerState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowLacpPartnerState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Activity returns PatternFlowLacpPartnerStateActivity, set in FlowLacpPartnerState.
	// PatternFlowLacpPartnerStateActivity is lACP Activity (1=Active, 0=Passive)
	Activity() PatternFlowLacpPartnerStateActivity
	// SetActivity assigns PatternFlowLacpPartnerStateActivity provided by user to FlowLacpPartnerState.
	// PatternFlowLacpPartnerStateActivity is lACP Activity (1=Active, 0=Passive)
	SetActivity(value PatternFlowLacpPartnerStateActivity) FlowLacpPartnerState
	// HasActivity checks if Activity has been set in FlowLacpPartnerState
	HasActivity() bool
	// Timeout returns PatternFlowLacpPartnerStateTimeout, set in FlowLacpPartnerState.
	// PatternFlowLacpPartnerStateTimeout is lACP Timeout (1=Fast timeout, 0=Slow timeout)
	Timeout() PatternFlowLacpPartnerStateTimeout
	// SetTimeout assigns PatternFlowLacpPartnerStateTimeout provided by user to FlowLacpPartnerState.
	// PatternFlowLacpPartnerStateTimeout is lACP Timeout (1=Fast timeout, 0=Slow timeout)
	SetTimeout(value PatternFlowLacpPartnerStateTimeout) FlowLacpPartnerState
	// HasTimeout checks if Timeout has been set in FlowLacpPartnerState
	HasTimeout() bool
	// Aggregation returns PatternFlowLacpPartnerStateAggregation, set in FlowLacpPartnerState.
	// PatternFlowLacpPartnerStateAggregation is aggregation (1=Aggregatable, 0=Individual)
	Aggregation() PatternFlowLacpPartnerStateAggregation
	// SetAggregation assigns PatternFlowLacpPartnerStateAggregation provided by user to FlowLacpPartnerState.
	// PatternFlowLacpPartnerStateAggregation is aggregation (1=Aggregatable, 0=Individual)
	SetAggregation(value PatternFlowLacpPartnerStateAggregation) FlowLacpPartnerState
	// HasAggregation checks if Aggregation has been set in FlowLacpPartnerState
	HasAggregation() bool
	// Synchronization returns PatternFlowLacpPartnerStateSynchronization, set in FlowLacpPartnerState.
	// PatternFlowLacpPartnerStateSynchronization is synchronization (1=In Sync, 0=Out of Sync)
	Synchronization() PatternFlowLacpPartnerStateSynchronization
	// SetSynchronization assigns PatternFlowLacpPartnerStateSynchronization provided by user to FlowLacpPartnerState.
	// PatternFlowLacpPartnerStateSynchronization is synchronization (1=In Sync, 0=Out of Sync)
	SetSynchronization(value PatternFlowLacpPartnerStateSynchronization) FlowLacpPartnerState
	// HasSynchronization checks if Synchronization has been set in FlowLacpPartnerState
	HasSynchronization() bool
	// Collecting returns PatternFlowLacpPartnerStateCollecting, set in FlowLacpPartnerState.
	// PatternFlowLacpPartnerStateCollecting is collecting (1=Enabled, 0=Disabled)
	Collecting() PatternFlowLacpPartnerStateCollecting
	// SetCollecting assigns PatternFlowLacpPartnerStateCollecting provided by user to FlowLacpPartnerState.
	// PatternFlowLacpPartnerStateCollecting is collecting (1=Enabled, 0=Disabled)
	SetCollecting(value PatternFlowLacpPartnerStateCollecting) FlowLacpPartnerState
	// HasCollecting checks if Collecting has been set in FlowLacpPartnerState
	HasCollecting() bool
	// Distributing returns PatternFlowLacpPartnerStateDistributing, set in FlowLacpPartnerState.
	// PatternFlowLacpPartnerStateDistributing is distributing (1=Enabled, 0=Disabled)
	Distributing() PatternFlowLacpPartnerStateDistributing
	// SetDistributing assigns PatternFlowLacpPartnerStateDistributing provided by user to FlowLacpPartnerState.
	// PatternFlowLacpPartnerStateDistributing is distributing (1=Enabled, 0=Disabled)
	SetDistributing(value PatternFlowLacpPartnerStateDistributing) FlowLacpPartnerState
	// HasDistributing checks if Distributing has been set in FlowLacpPartnerState
	HasDistributing() bool
	// Defaulted returns PatternFlowLacpPartnerStateDefaulted, set in FlowLacpPartnerState.
	// PatternFlowLacpPartnerStateDefaulted is defaulted (1=Using defaulted partner info, 0=Using received partner info)
	Defaulted() PatternFlowLacpPartnerStateDefaulted
	// SetDefaulted assigns PatternFlowLacpPartnerStateDefaulted provided by user to FlowLacpPartnerState.
	// PatternFlowLacpPartnerStateDefaulted is defaulted (1=Using defaulted partner info, 0=Using received partner info)
	SetDefaulted(value PatternFlowLacpPartnerStateDefaulted) FlowLacpPartnerState
	// HasDefaulted checks if Defaulted has been set in FlowLacpPartnerState
	HasDefaulted() bool
	// Expired returns PatternFlowLacpPartnerStateExpired, set in FlowLacpPartnerState.
	// PatternFlowLacpPartnerStateExpired is expired (1=Expired, 0=Not Expired)
	Expired() PatternFlowLacpPartnerStateExpired
	// SetExpired assigns PatternFlowLacpPartnerStateExpired provided by user to FlowLacpPartnerState.
	// PatternFlowLacpPartnerStateExpired is expired (1=Expired, 0=Not Expired)
	SetExpired(value PatternFlowLacpPartnerStateExpired) FlowLacpPartnerState
	// HasExpired checks if Expired has been set in FlowLacpPartnerState
	HasExpired() bool
	setNil()
}

// description is TBD
// Activity returns a PatternFlowLacpPartnerStateActivity
func (obj *flowLacpPartnerState) Activity() PatternFlowLacpPartnerStateActivity {
	if obj.obj.Activity == nil {
		obj.obj.Activity = NewPatternFlowLacpPartnerStateActivity().msg()
	}
	if obj.activityHolder == nil {
		obj.activityHolder = &patternFlowLacpPartnerStateActivity{obj: obj.obj.Activity}
	}
	return obj.activityHolder
}

// description is TBD
// Activity returns a PatternFlowLacpPartnerStateActivity
func (obj *flowLacpPartnerState) HasActivity() bool {
	return obj.obj.Activity != nil
}

// description is TBD
// SetActivity sets the PatternFlowLacpPartnerStateActivity value in the FlowLacpPartnerState object
func (obj *flowLacpPartnerState) SetActivity(value PatternFlowLacpPartnerStateActivity) FlowLacpPartnerState {

	obj.activityHolder = nil
	obj.obj.Activity = value.msg()

	return obj
}

// description is TBD
// Timeout returns a PatternFlowLacpPartnerStateTimeout
func (obj *flowLacpPartnerState) Timeout() PatternFlowLacpPartnerStateTimeout {
	if obj.obj.Timeout == nil {
		obj.obj.Timeout = NewPatternFlowLacpPartnerStateTimeout().msg()
	}
	if obj.timeoutHolder == nil {
		obj.timeoutHolder = &patternFlowLacpPartnerStateTimeout{obj: obj.obj.Timeout}
	}
	return obj.timeoutHolder
}

// description is TBD
// Timeout returns a PatternFlowLacpPartnerStateTimeout
func (obj *flowLacpPartnerState) HasTimeout() bool {
	return obj.obj.Timeout != nil
}

// description is TBD
// SetTimeout sets the PatternFlowLacpPartnerStateTimeout value in the FlowLacpPartnerState object
func (obj *flowLacpPartnerState) SetTimeout(value PatternFlowLacpPartnerStateTimeout) FlowLacpPartnerState {

	obj.timeoutHolder = nil
	obj.obj.Timeout = value.msg()

	return obj
}

// description is TBD
// Aggregation returns a PatternFlowLacpPartnerStateAggregation
func (obj *flowLacpPartnerState) Aggregation() PatternFlowLacpPartnerStateAggregation {
	if obj.obj.Aggregation == nil {
		obj.obj.Aggregation = NewPatternFlowLacpPartnerStateAggregation().msg()
	}
	if obj.aggregationHolder == nil {
		obj.aggregationHolder = &patternFlowLacpPartnerStateAggregation{obj: obj.obj.Aggregation}
	}
	return obj.aggregationHolder
}

// description is TBD
// Aggregation returns a PatternFlowLacpPartnerStateAggregation
func (obj *flowLacpPartnerState) HasAggregation() bool {
	return obj.obj.Aggregation != nil
}

// description is TBD
// SetAggregation sets the PatternFlowLacpPartnerStateAggregation value in the FlowLacpPartnerState object
func (obj *flowLacpPartnerState) SetAggregation(value PatternFlowLacpPartnerStateAggregation) FlowLacpPartnerState {

	obj.aggregationHolder = nil
	obj.obj.Aggregation = value.msg()

	return obj
}

// description is TBD
// Synchronization returns a PatternFlowLacpPartnerStateSynchronization
func (obj *flowLacpPartnerState) Synchronization() PatternFlowLacpPartnerStateSynchronization {
	if obj.obj.Synchronization == nil {
		obj.obj.Synchronization = NewPatternFlowLacpPartnerStateSynchronization().msg()
	}
	if obj.synchronizationHolder == nil {
		obj.synchronizationHolder = &patternFlowLacpPartnerStateSynchronization{obj: obj.obj.Synchronization}
	}
	return obj.synchronizationHolder
}

// description is TBD
// Synchronization returns a PatternFlowLacpPartnerStateSynchronization
func (obj *flowLacpPartnerState) HasSynchronization() bool {
	return obj.obj.Synchronization != nil
}

// description is TBD
// SetSynchronization sets the PatternFlowLacpPartnerStateSynchronization value in the FlowLacpPartnerState object
func (obj *flowLacpPartnerState) SetSynchronization(value PatternFlowLacpPartnerStateSynchronization) FlowLacpPartnerState {

	obj.synchronizationHolder = nil
	obj.obj.Synchronization = value.msg()

	return obj
}

// description is TBD
// Collecting returns a PatternFlowLacpPartnerStateCollecting
func (obj *flowLacpPartnerState) Collecting() PatternFlowLacpPartnerStateCollecting {
	if obj.obj.Collecting == nil {
		obj.obj.Collecting = NewPatternFlowLacpPartnerStateCollecting().msg()
	}
	if obj.collectingHolder == nil {
		obj.collectingHolder = &patternFlowLacpPartnerStateCollecting{obj: obj.obj.Collecting}
	}
	return obj.collectingHolder
}

// description is TBD
// Collecting returns a PatternFlowLacpPartnerStateCollecting
func (obj *flowLacpPartnerState) HasCollecting() bool {
	return obj.obj.Collecting != nil
}

// description is TBD
// SetCollecting sets the PatternFlowLacpPartnerStateCollecting value in the FlowLacpPartnerState object
func (obj *flowLacpPartnerState) SetCollecting(value PatternFlowLacpPartnerStateCollecting) FlowLacpPartnerState {

	obj.collectingHolder = nil
	obj.obj.Collecting = value.msg()

	return obj
}

// description is TBD
// Distributing returns a PatternFlowLacpPartnerStateDistributing
func (obj *flowLacpPartnerState) Distributing() PatternFlowLacpPartnerStateDistributing {
	if obj.obj.Distributing == nil {
		obj.obj.Distributing = NewPatternFlowLacpPartnerStateDistributing().msg()
	}
	if obj.distributingHolder == nil {
		obj.distributingHolder = &patternFlowLacpPartnerStateDistributing{obj: obj.obj.Distributing}
	}
	return obj.distributingHolder
}

// description is TBD
// Distributing returns a PatternFlowLacpPartnerStateDistributing
func (obj *flowLacpPartnerState) HasDistributing() bool {
	return obj.obj.Distributing != nil
}

// description is TBD
// SetDistributing sets the PatternFlowLacpPartnerStateDistributing value in the FlowLacpPartnerState object
func (obj *flowLacpPartnerState) SetDistributing(value PatternFlowLacpPartnerStateDistributing) FlowLacpPartnerState {

	obj.distributingHolder = nil
	obj.obj.Distributing = value.msg()

	return obj
}

// description is TBD
// Defaulted returns a PatternFlowLacpPartnerStateDefaulted
func (obj *flowLacpPartnerState) Defaulted() PatternFlowLacpPartnerStateDefaulted {
	if obj.obj.Defaulted == nil {
		obj.obj.Defaulted = NewPatternFlowLacpPartnerStateDefaulted().msg()
	}
	if obj.defaultedHolder == nil {
		obj.defaultedHolder = &patternFlowLacpPartnerStateDefaulted{obj: obj.obj.Defaulted}
	}
	return obj.defaultedHolder
}

// description is TBD
// Defaulted returns a PatternFlowLacpPartnerStateDefaulted
func (obj *flowLacpPartnerState) HasDefaulted() bool {
	return obj.obj.Defaulted != nil
}

// description is TBD
// SetDefaulted sets the PatternFlowLacpPartnerStateDefaulted value in the FlowLacpPartnerState object
func (obj *flowLacpPartnerState) SetDefaulted(value PatternFlowLacpPartnerStateDefaulted) FlowLacpPartnerState {

	obj.defaultedHolder = nil
	obj.obj.Defaulted = value.msg()

	return obj
}

// description is TBD
// Expired returns a PatternFlowLacpPartnerStateExpired
func (obj *flowLacpPartnerState) Expired() PatternFlowLacpPartnerStateExpired {
	if obj.obj.Expired == nil {
		obj.obj.Expired = NewPatternFlowLacpPartnerStateExpired().msg()
	}
	if obj.expiredHolder == nil {
		obj.expiredHolder = &patternFlowLacpPartnerStateExpired{obj: obj.obj.Expired}
	}
	return obj.expiredHolder
}

// description is TBD
// Expired returns a PatternFlowLacpPartnerStateExpired
func (obj *flowLacpPartnerState) HasExpired() bool {
	return obj.obj.Expired != nil
}

// description is TBD
// SetExpired sets the PatternFlowLacpPartnerStateExpired value in the FlowLacpPartnerState object
func (obj *flowLacpPartnerState) SetExpired(value PatternFlowLacpPartnerStateExpired) FlowLacpPartnerState {

	obj.expiredHolder = nil
	obj.obj.Expired = value.msg()

	return obj
}

func (obj *flowLacpPartnerState) validateObj(vObj *validation, set_default bool) {
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

func (obj *flowLacpPartnerState) setDefault() {

}
