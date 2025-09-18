package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowLacpduPartnerState *****
type flowLacpduPartnerState struct {
	validation
	obj                   *otg.FlowLacpduPartnerState
	marshaller            marshalFlowLacpduPartnerState
	unMarshaller          unMarshalFlowLacpduPartnerState
	activityHolder        PatternFlowLacpduPartnerStateActivity
	timeoutHolder         PatternFlowLacpduPartnerStateTimeout
	aggregationHolder     PatternFlowLacpduPartnerStateAggregation
	synchronizationHolder PatternFlowLacpduPartnerStateSynchronization
	collectingHolder      PatternFlowLacpduPartnerStateCollecting
	distributingHolder    PatternFlowLacpduPartnerStateDistributing
	defaultedHolder       PatternFlowLacpduPartnerStateDefaulted
	expiredHolder         PatternFlowLacpduPartnerStateExpired
}

func NewFlowLacpduPartnerState() FlowLacpduPartnerState {
	obj := flowLacpduPartnerState{obj: &otg.FlowLacpduPartnerState{}}
	obj.setDefault()
	return &obj
}

func (obj *flowLacpduPartnerState) msg() *otg.FlowLacpduPartnerState {
	return obj.obj
}

func (obj *flowLacpduPartnerState) setMsg(msg *otg.FlowLacpduPartnerState) FlowLacpduPartnerState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowLacpduPartnerState struct {
	obj *flowLacpduPartnerState
}

type marshalFlowLacpduPartnerState interface {
	// ToProto marshals FlowLacpduPartnerState to protobuf object *otg.FlowLacpduPartnerState
	ToProto() (*otg.FlowLacpduPartnerState, error)
	// ToPbText marshals FlowLacpduPartnerState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowLacpduPartnerState to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowLacpduPartnerState to JSON text
	ToJson() (string, error)
}

type unMarshalflowLacpduPartnerState struct {
	obj *flowLacpduPartnerState
}

type unMarshalFlowLacpduPartnerState interface {
	// FromProto unmarshals FlowLacpduPartnerState from protobuf object *otg.FlowLacpduPartnerState
	FromProto(msg *otg.FlowLacpduPartnerState) (FlowLacpduPartnerState, error)
	// FromPbText unmarshals FlowLacpduPartnerState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowLacpduPartnerState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowLacpduPartnerState from JSON text
	FromJson(value string) error
}

func (obj *flowLacpduPartnerState) Marshal() marshalFlowLacpduPartnerState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowLacpduPartnerState{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowLacpduPartnerState) Unmarshal() unMarshalFlowLacpduPartnerState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowLacpduPartnerState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowLacpduPartnerState) ToProto() (*otg.FlowLacpduPartnerState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowLacpduPartnerState) FromProto(msg *otg.FlowLacpduPartnerState) (FlowLacpduPartnerState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowLacpduPartnerState) ToPbText() (string, error) {
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

func (m *unMarshalflowLacpduPartnerState) FromPbText(value string) error {
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

func (m *marshalflowLacpduPartnerState) ToYaml() (string, error) {
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

func (m *unMarshalflowLacpduPartnerState) FromYaml(value string) error {
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

func (m *marshalflowLacpduPartnerState) ToJson() (string, error) {
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

func (m *unMarshalflowLacpduPartnerState) FromJson(value string) error {
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

func (obj *flowLacpduPartnerState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowLacpduPartnerState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowLacpduPartnerState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowLacpduPartnerState) Clone() (FlowLacpduPartnerState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowLacpduPartnerState()
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

func (obj *flowLacpduPartnerState) setNil() {
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

// FlowLacpduPartnerState is this field indicates the Partner's state.
type FlowLacpduPartnerState interface {
	Validation
	// msg marshals FlowLacpduPartnerState to protobuf object *otg.FlowLacpduPartnerState
	// and doesn't set defaults
	msg() *otg.FlowLacpduPartnerState
	// setMsg unmarshals FlowLacpduPartnerState from protobuf object *otg.FlowLacpduPartnerState
	// and doesn't set defaults
	setMsg(*otg.FlowLacpduPartnerState) FlowLacpduPartnerState
	// provides marshal interface
	Marshal() marshalFlowLacpduPartnerState
	// provides unmarshal interface
	Unmarshal() unMarshalFlowLacpduPartnerState
	// validate validates FlowLacpduPartnerState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowLacpduPartnerState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Activity returns PatternFlowLacpduPartnerStateActivity, set in FlowLacpduPartnerState.
	// PatternFlowLacpduPartnerStateActivity is lACP Activity (1=Active, 0=Passive)
	Activity() PatternFlowLacpduPartnerStateActivity
	// SetActivity assigns PatternFlowLacpduPartnerStateActivity provided by user to FlowLacpduPartnerState.
	// PatternFlowLacpduPartnerStateActivity is lACP Activity (1=Active, 0=Passive)
	SetActivity(value PatternFlowLacpduPartnerStateActivity) FlowLacpduPartnerState
	// HasActivity checks if Activity has been set in FlowLacpduPartnerState
	HasActivity() bool
	// Timeout returns PatternFlowLacpduPartnerStateTimeout, set in FlowLacpduPartnerState.
	// PatternFlowLacpduPartnerStateTimeout is lACP Timeout (1=Fast timeout, 0=Slow timeout)
	Timeout() PatternFlowLacpduPartnerStateTimeout
	// SetTimeout assigns PatternFlowLacpduPartnerStateTimeout provided by user to FlowLacpduPartnerState.
	// PatternFlowLacpduPartnerStateTimeout is lACP Timeout (1=Fast timeout, 0=Slow timeout)
	SetTimeout(value PatternFlowLacpduPartnerStateTimeout) FlowLacpduPartnerState
	// HasTimeout checks if Timeout has been set in FlowLacpduPartnerState
	HasTimeout() bool
	// Aggregation returns PatternFlowLacpduPartnerStateAggregation, set in FlowLacpduPartnerState.
	// PatternFlowLacpduPartnerStateAggregation is aggregation (1=Aggregatable, 0=Individual)
	Aggregation() PatternFlowLacpduPartnerStateAggregation
	// SetAggregation assigns PatternFlowLacpduPartnerStateAggregation provided by user to FlowLacpduPartnerState.
	// PatternFlowLacpduPartnerStateAggregation is aggregation (1=Aggregatable, 0=Individual)
	SetAggregation(value PatternFlowLacpduPartnerStateAggregation) FlowLacpduPartnerState
	// HasAggregation checks if Aggregation has been set in FlowLacpduPartnerState
	HasAggregation() bool
	// Synchronization returns PatternFlowLacpduPartnerStateSynchronization, set in FlowLacpduPartnerState.
	// PatternFlowLacpduPartnerStateSynchronization is synchronization (1=In Sync, 0=Out of Sync)
	Synchronization() PatternFlowLacpduPartnerStateSynchronization
	// SetSynchronization assigns PatternFlowLacpduPartnerStateSynchronization provided by user to FlowLacpduPartnerState.
	// PatternFlowLacpduPartnerStateSynchronization is synchronization (1=In Sync, 0=Out of Sync)
	SetSynchronization(value PatternFlowLacpduPartnerStateSynchronization) FlowLacpduPartnerState
	// HasSynchronization checks if Synchronization has been set in FlowLacpduPartnerState
	HasSynchronization() bool
	// Collecting returns PatternFlowLacpduPartnerStateCollecting, set in FlowLacpduPartnerState.
	// PatternFlowLacpduPartnerStateCollecting is collecting (1=Enabled, 0=Disabled)
	Collecting() PatternFlowLacpduPartnerStateCollecting
	// SetCollecting assigns PatternFlowLacpduPartnerStateCollecting provided by user to FlowLacpduPartnerState.
	// PatternFlowLacpduPartnerStateCollecting is collecting (1=Enabled, 0=Disabled)
	SetCollecting(value PatternFlowLacpduPartnerStateCollecting) FlowLacpduPartnerState
	// HasCollecting checks if Collecting has been set in FlowLacpduPartnerState
	HasCollecting() bool
	// Distributing returns PatternFlowLacpduPartnerStateDistributing, set in FlowLacpduPartnerState.
	// PatternFlowLacpduPartnerStateDistributing is distributing (1=Enabled, 0=Disabled)
	Distributing() PatternFlowLacpduPartnerStateDistributing
	// SetDistributing assigns PatternFlowLacpduPartnerStateDistributing provided by user to FlowLacpduPartnerState.
	// PatternFlowLacpduPartnerStateDistributing is distributing (1=Enabled, 0=Disabled)
	SetDistributing(value PatternFlowLacpduPartnerStateDistributing) FlowLacpduPartnerState
	// HasDistributing checks if Distributing has been set in FlowLacpduPartnerState
	HasDistributing() bool
	// Defaulted returns PatternFlowLacpduPartnerStateDefaulted, set in FlowLacpduPartnerState.
	// PatternFlowLacpduPartnerStateDefaulted is defaulted (1=Using defaulted partner info, 0=Using received partner info)
	Defaulted() PatternFlowLacpduPartnerStateDefaulted
	// SetDefaulted assigns PatternFlowLacpduPartnerStateDefaulted provided by user to FlowLacpduPartnerState.
	// PatternFlowLacpduPartnerStateDefaulted is defaulted (1=Using defaulted partner info, 0=Using received partner info)
	SetDefaulted(value PatternFlowLacpduPartnerStateDefaulted) FlowLacpduPartnerState
	// HasDefaulted checks if Defaulted has been set in FlowLacpduPartnerState
	HasDefaulted() bool
	// Expired returns PatternFlowLacpduPartnerStateExpired, set in FlowLacpduPartnerState.
	// PatternFlowLacpduPartnerStateExpired is expired (1=Expired, 0=Not Expired)
	Expired() PatternFlowLacpduPartnerStateExpired
	// SetExpired assigns PatternFlowLacpduPartnerStateExpired provided by user to FlowLacpduPartnerState.
	// PatternFlowLacpduPartnerStateExpired is expired (1=Expired, 0=Not Expired)
	SetExpired(value PatternFlowLacpduPartnerStateExpired) FlowLacpduPartnerState
	// HasExpired checks if Expired has been set in FlowLacpduPartnerState
	HasExpired() bool
	setNil()
}

// description is TBD
// Activity returns a PatternFlowLacpduPartnerStateActivity
func (obj *flowLacpduPartnerState) Activity() PatternFlowLacpduPartnerStateActivity {
	if obj.obj.Activity == nil {
		obj.obj.Activity = NewPatternFlowLacpduPartnerStateActivity().msg()
	}
	if obj.activityHolder == nil {
		obj.activityHolder = &patternFlowLacpduPartnerStateActivity{obj: obj.obj.Activity}
	}
	return obj.activityHolder
}

// description is TBD
// Activity returns a PatternFlowLacpduPartnerStateActivity
func (obj *flowLacpduPartnerState) HasActivity() bool {
	return obj.obj.Activity != nil
}

// description is TBD
// SetActivity sets the PatternFlowLacpduPartnerStateActivity value in the FlowLacpduPartnerState object
func (obj *flowLacpduPartnerState) SetActivity(value PatternFlowLacpduPartnerStateActivity) FlowLacpduPartnerState {

	obj.activityHolder = nil
	obj.obj.Activity = value.msg()

	return obj
}

// description is TBD
// Timeout returns a PatternFlowLacpduPartnerStateTimeout
func (obj *flowLacpduPartnerState) Timeout() PatternFlowLacpduPartnerStateTimeout {
	if obj.obj.Timeout == nil {
		obj.obj.Timeout = NewPatternFlowLacpduPartnerStateTimeout().msg()
	}
	if obj.timeoutHolder == nil {
		obj.timeoutHolder = &patternFlowLacpduPartnerStateTimeout{obj: obj.obj.Timeout}
	}
	return obj.timeoutHolder
}

// description is TBD
// Timeout returns a PatternFlowLacpduPartnerStateTimeout
func (obj *flowLacpduPartnerState) HasTimeout() bool {
	return obj.obj.Timeout != nil
}

// description is TBD
// SetTimeout sets the PatternFlowLacpduPartnerStateTimeout value in the FlowLacpduPartnerState object
func (obj *flowLacpduPartnerState) SetTimeout(value PatternFlowLacpduPartnerStateTimeout) FlowLacpduPartnerState {

	obj.timeoutHolder = nil
	obj.obj.Timeout = value.msg()

	return obj
}

// description is TBD
// Aggregation returns a PatternFlowLacpduPartnerStateAggregation
func (obj *flowLacpduPartnerState) Aggregation() PatternFlowLacpduPartnerStateAggregation {
	if obj.obj.Aggregation == nil {
		obj.obj.Aggregation = NewPatternFlowLacpduPartnerStateAggregation().msg()
	}
	if obj.aggregationHolder == nil {
		obj.aggregationHolder = &patternFlowLacpduPartnerStateAggregation{obj: obj.obj.Aggregation}
	}
	return obj.aggregationHolder
}

// description is TBD
// Aggregation returns a PatternFlowLacpduPartnerStateAggregation
func (obj *flowLacpduPartnerState) HasAggregation() bool {
	return obj.obj.Aggregation != nil
}

// description is TBD
// SetAggregation sets the PatternFlowLacpduPartnerStateAggregation value in the FlowLacpduPartnerState object
func (obj *flowLacpduPartnerState) SetAggregation(value PatternFlowLacpduPartnerStateAggregation) FlowLacpduPartnerState {

	obj.aggregationHolder = nil
	obj.obj.Aggregation = value.msg()

	return obj
}

// description is TBD
// Synchronization returns a PatternFlowLacpduPartnerStateSynchronization
func (obj *flowLacpduPartnerState) Synchronization() PatternFlowLacpduPartnerStateSynchronization {
	if obj.obj.Synchronization == nil {
		obj.obj.Synchronization = NewPatternFlowLacpduPartnerStateSynchronization().msg()
	}
	if obj.synchronizationHolder == nil {
		obj.synchronizationHolder = &patternFlowLacpduPartnerStateSynchronization{obj: obj.obj.Synchronization}
	}
	return obj.synchronizationHolder
}

// description is TBD
// Synchronization returns a PatternFlowLacpduPartnerStateSynchronization
func (obj *flowLacpduPartnerState) HasSynchronization() bool {
	return obj.obj.Synchronization != nil
}

// description is TBD
// SetSynchronization sets the PatternFlowLacpduPartnerStateSynchronization value in the FlowLacpduPartnerState object
func (obj *flowLacpduPartnerState) SetSynchronization(value PatternFlowLacpduPartnerStateSynchronization) FlowLacpduPartnerState {

	obj.synchronizationHolder = nil
	obj.obj.Synchronization = value.msg()

	return obj
}

// description is TBD
// Collecting returns a PatternFlowLacpduPartnerStateCollecting
func (obj *flowLacpduPartnerState) Collecting() PatternFlowLacpduPartnerStateCollecting {
	if obj.obj.Collecting == nil {
		obj.obj.Collecting = NewPatternFlowLacpduPartnerStateCollecting().msg()
	}
	if obj.collectingHolder == nil {
		obj.collectingHolder = &patternFlowLacpduPartnerStateCollecting{obj: obj.obj.Collecting}
	}
	return obj.collectingHolder
}

// description is TBD
// Collecting returns a PatternFlowLacpduPartnerStateCollecting
func (obj *flowLacpduPartnerState) HasCollecting() bool {
	return obj.obj.Collecting != nil
}

// description is TBD
// SetCollecting sets the PatternFlowLacpduPartnerStateCollecting value in the FlowLacpduPartnerState object
func (obj *flowLacpduPartnerState) SetCollecting(value PatternFlowLacpduPartnerStateCollecting) FlowLacpduPartnerState {

	obj.collectingHolder = nil
	obj.obj.Collecting = value.msg()

	return obj
}

// description is TBD
// Distributing returns a PatternFlowLacpduPartnerStateDistributing
func (obj *flowLacpduPartnerState) Distributing() PatternFlowLacpduPartnerStateDistributing {
	if obj.obj.Distributing == nil {
		obj.obj.Distributing = NewPatternFlowLacpduPartnerStateDistributing().msg()
	}
	if obj.distributingHolder == nil {
		obj.distributingHolder = &patternFlowLacpduPartnerStateDistributing{obj: obj.obj.Distributing}
	}
	return obj.distributingHolder
}

// description is TBD
// Distributing returns a PatternFlowLacpduPartnerStateDistributing
func (obj *flowLacpduPartnerState) HasDistributing() bool {
	return obj.obj.Distributing != nil
}

// description is TBD
// SetDistributing sets the PatternFlowLacpduPartnerStateDistributing value in the FlowLacpduPartnerState object
func (obj *flowLacpduPartnerState) SetDistributing(value PatternFlowLacpduPartnerStateDistributing) FlowLacpduPartnerState {

	obj.distributingHolder = nil
	obj.obj.Distributing = value.msg()

	return obj
}

// description is TBD
// Defaulted returns a PatternFlowLacpduPartnerStateDefaulted
func (obj *flowLacpduPartnerState) Defaulted() PatternFlowLacpduPartnerStateDefaulted {
	if obj.obj.Defaulted == nil {
		obj.obj.Defaulted = NewPatternFlowLacpduPartnerStateDefaulted().msg()
	}
	if obj.defaultedHolder == nil {
		obj.defaultedHolder = &patternFlowLacpduPartnerStateDefaulted{obj: obj.obj.Defaulted}
	}
	return obj.defaultedHolder
}

// description is TBD
// Defaulted returns a PatternFlowLacpduPartnerStateDefaulted
func (obj *flowLacpduPartnerState) HasDefaulted() bool {
	return obj.obj.Defaulted != nil
}

// description is TBD
// SetDefaulted sets the PatternFlowLacpduPartnerStateDefaulted value in the FlowLacpduPartnerState object
func (obj *flowLacpduPartnerState) SetDefaulted(value PatternFlowLacpduPartnerStateDefaulted) FlowLacpduPartnerState {

	obj.defaultedHolder = nil
	obj.obj.Defaulted = value.msg()

	return obj
}

// description is TBD
// Expired returns a PatternFlowLacpduPartnerStateExpired
func (obj *flowLacpduPartnerState) Expired() PatternFlowLacpduPartnerStateExpired {
	if obj.obj.Expired == nil {
		obj.obj.Expired = NewPatternFlowLacpduPartnerStateExpired().msg()
	}
	if obj.expiredHolder == nil {
		obj.expiredHolder = &patternFlowLacpduPartnerStateExpired{obj: obj.obj.Expired}
	}
	return obj.expiredHolder
}

// description is TBD
// Expired returns a PatternFlowLacpduPartnerStateExpired
func (obj *flowLacpduPartnerState) HasExpired() bool {
	return obj.obj.Expired != nil
}

// description is TBD
// SetExpired sets the PatternFlowLacpduPartnerStateExpired value in the FlowLacpduPartnerState object
func (obj *flowLacpduPartnerState) SetExpired(value PatternFlowLacpduPartnerStateExpired) FlowLacpduPartnerState {

	obj.expiredHolder = nil
	obj.obj.Expired = value.msg()

	return obj
}

func (obj *flowLacpduPartnerState) validateObj(vObj *validation, set_default bool) {
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

func (obj *flowLacpduPartnerState) setDefault() {

}
