package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowMpls *****
type flowMpls struct {
	validation
	obj                 *otg.FlowMpls
	marshaller          marshalFlowMpls
	unMarshaller        unMarshalFlowMpls
	labelHolder         PatternFlowMplsLabel
	trafficClassHolder  PatternFlowMplsTrafficClass
	bottomOfStackHolder PatternFlowMplsBottomOfStack
	timeToLiveHolder    PatternFlowMplsTimeToLive
}

func NewFlowMpls() FlowMpls {
	obj := flowMpls{obj: &otg.FlowMpls{}}
	obj.setDefault()
	return &obj
}

func (obj *flowMpls) msg() *otg.FlowMpls {
	return obj.obj
}

func (obj *flowMpls) setMsg(msg *otg.FlowMpls) FlowMpls {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowMpls struct {
	obj *flowMpls
}

type marshalFlowMpls interface {
	// ToProto marshals FlowMpls to protobuf object *otg.FlowMpls
	ToProto() (*otg.FlowMpls, error)
	// ToPbText marshals FlowMpls to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowMpls to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowMpls to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowMpls to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowMpls struct {
	obj *flowMpls
}

type unMarshalFlowMpls interface {
	// FromProto unmarshals FlowMpls from protobuf object *otg.FlowMpls
	FromProto(msg *otg.FlowMpls) (FlowMpls, error)
	// FromPbText unmarshals FlowMpls from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowMpls from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowMpls from JSON text
	FromJson(value string) error
}

func (obj *flowMpls) Marshal() marshalFlowMpls {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowMpls{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowMpls) Unmarshal() unMarshalFlowMpls {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowMpls{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowMpls) ToProto() (*otg.FlowMpls, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowMpls) FromProto(msg *otg.FlowMpls) (FlowMpls, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowMpls) ToPbText() (string, error) {
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

func (m *unMarshalflowMpls) FromPbText(value string) error {
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

func (m *marshalflowMpls) ToYaml() (string, error) {
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

func (m *unMarshalflowMpls) FromYaml(value string) error {
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

func (m *marshalflowMpls) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalflowMpls) ToJson() (string, error) {
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

func (m *unMarshalflowMpls) FromJson(value string) error {
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

func (obj *flowMpls) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowMpls) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowMpls) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowMpls) Clone() (FlowMpls, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowMpls()
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

func (obj *flowMpls) setNil() {
	obj.labelHolder = nil
	obj.trafficClassHolder = nil
	obj.bottomOfStackHolder = nil
	obj.timeToLiveHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowMpls is mPLS packet header; When configuring multiple such headers, the count shall not exceed 20.
type FlowMpls interface {
	Validation
	// msg marshals FlowMpls to protobuf object *otg.FlowMpls
	// and doesn't set defaults
	msg() *otg.FlowMpls
	// setMsg unmarshals FlowMpls from protobuf object *otg.FlowMpls
	// and doesn't set defaults
	setMsg(*otg.FlowMpls) FlowMpls
	// provides marshal interface
	Marshal() marshalFlowMpls
	// provides unmarshal interface
	Unmarshal() unMarshalFlowMpls
	// validate validates FlowMpls
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowMpls, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Label returns PatternFlowMplsLabel, set in FlowMpls.
	// PatternFlowMplsLabel is label of routers
	Label() PatternFlowMplsLabel
	// SetLabel assigns PatternFlowMplsLabel provided by user to FlowMpls.
	// PatternFlowMplsLabel is label of routers
	SetLabel(value PatternFlowMplsLabel) FlowMpls
	// HasLabel checks if Label has been set in FlowMpls
	HasLabel() bool
	// TrafficClass returns PatternFlowMplsTrafficClass, set in FlowMpls.
	// PatternFlowMplsTrafficClass is traffic class
	TrafficClass() PatternFlowMplsTrafficClass
	// SetTrafficClass assigns PatternFlowMplsTrafficClass provided by user to FlowMpls.
	// PatternFlowMplsTrafficClass is traffic class
	SetTrafficClass(value PatternFlowMplsTrafficClass) FlowMpls
	// HasTrafficClass checks if TrafficClass has been set in FlowMpls
	HasTrafficClass() bool
	// BottomOfStack returns PatternFlowMplsBottomOfStack, set in FlowMpls.
	// PatternFlowMplsBottomOfStack is bottom of stack
	BottomOfStack() PatternFlowMplsBottomOfStack
	// SetBottomOfStack assigns PatternFlowMplsBottomOfStack provided by user to FlowMpls.
	// PatternFlowMplsBottomOfStack is bottom of stack
	SetBottomOfStack(value PatternFlowMplsBottomOfStack) FlowMpls
	// HasBottomOfStack checks if BottomOfStack has been set in FlowMpls
	HasBottomOfStack() bool
	// TimeToLive returns PatternFlowMplsTimeToLive, set in FlowMpls.
	// PatternFlowMplsTimeToLive is time to live
	TimeToLive() PatternFlowMplsTimeToLive
	// SetTimeToLive assigns PatternFlowMplsTimeToLive provided by user to FlowMpls.
	// PatternFlowMplsTimeToLive is time to live
	SetTimeToLive(value PatternFlowMplsTimeToLive) FlowMpls
	// HasTimeToLive checks if TimeToLive has been set in FlowMpls
	HasTimeToLive() bool
	setNil()
}

// description is TBD
// Label returns a PatternFlowMplsLabel
func (obj *flowMpls) Label() PatternFlowMplsLabel {
	if obj.obj.Label == nil {
		obj.obj.Label = NewPatternFlowMplsLabel().msg()
	}
	if obj.labelHolder == nil {
		obj.labelHolder = &patternFlowMplsLabel{obj: obj.obj.Label}
	}
	return obj.labelHolder
}

// description is TBD
// Label returns a PatternFlowMplsLabel
func (obj *flowMpls) HasLabel() bool {
	return obj.obj.Label != nil
}

// description is TBD
// SetLabel sets the PatternFlowMplsLabel value in the FlowMpls object
func (obj *flowMpls) SetLabel(value PatternFlowMplsLabel) FlowMpls {

	obj.labelHolder = nil
	obj.obj.Label = value.msg()

	return obj
}

// description is TBD
// TrafficClass returns a PatternFlowMplsTrafficClass
func (obj *flowMpls) TrafficClass() PatternFlowMplsTrafficClass {
	if obj.obj.TrafficClass == nil {
		obj.obj.TrafficClass = NewPatternFlowMplsTrafficClass().msg()
	}
	if obj.trafficClassHolder == nil {
		obj.trafficClassHolder = &patternFlowMplsTrafficClass{obj: obj.obj.TrafficClass}
	}
	return obj.trafficClassHolder
}

// description is TBD
// TrafficClass returns a PatternFlowMplsTrafficClass
func (obj *flowMpls) HasTrafficClass() bool {
	return obj.obj.TrafficClass != nil
}

// description is TBD
// SetTrafficClass sets the PatternFlowMplsTrafficClass value in the FlowMpls object
func (obj *flowMpls) SetTrafficClass(value PatternFlowMplsTrafficClass) FlowMpls {

	obj.trafficClassHolder = nil
	obj.obj.TrafficClass = value.msg()

	return obj
}

// description is TBD
// BottomOfStack returns a PatternFlowMplsBottomOfStack
func (obj *flowMpls) BottomOfStack() PatternFlowMplsBottomOfStack {
	if obj.obj.BottomOfStack == nil {
		obj.obj.BottomOfStack = NewPatternFlowMplsBottomOfStack().msg()
	}
	if obj.bottomOfStackHolder == nil {
		obj.bottomOfStackHolder = &patternFlowMplsBottomOfStack{obj: obj.obj.BottomOfStack}
	}
	return obj.bottomOfStackHolder
}

// description is TBD
// BottomOfStack returns a PatternFlowMplsBottomOfStack
func (obj *flowMpls) HasBottomOfStack() bool {
	return obj.obj.BottomOfStack != nil
}

// description is TBD
// SetBottomOfStack sets the PatternFlowMplsBottomOfStack value in the FlowMpls object
func (obj *flowMpls) SetBottomOfStack(value PatternFlowMplsBottomOfStack) FlowMpls {

	obj.bottomOfStackHolder = nil
	obj.obj.BottomOfStack = value.msg()

	return obj
}

// description is TBD
// TimeToLive returns a PatternFlowMplsTimeToLive
func (obj *flowMpls) TimeToLive() PatternFlowMplsTimeToLive {
	if obj.obj.TimeToLive == nil {
		obj.obj.TimeToLive = NewPatternFlowMplsTimeToLive().msg()
	}
	if obj.timeToLiveHolder == nil {
		obj.timeToLiveHolder = &patternFlowMplsTimeToLive{obj: obj.obj.TimeToLive}
	}
	return obj.timeToLiveHolder
}

// description is TBD
// TimeToLive returns a PatternFlowMplsTimeToLive
func (obj *flowMpls) HasTimeToLive() bool {
	return obj.obj.TimeToLive != nil
}

// description is TBD
// SetTimeToLive sets the PatternFlowMplsTimeToLive value in the FlowMpls object
func (obj *flowMpls) SetTimeToLive(value PatternFlowMplsTimeToLive) FlowMpls {

	obj.timeToLiveHolder = nil
	obj.obj.TimeToLive = value.msg()

	return obj
}

func (obj *flowMpls) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Label != nil {

		obj.Label().validateObj(vObj, set_default)
	}

	if obj.obj.TrafficClass != nil {

		obj.TrafficClass().validateObj(vObj, set_default)
	}

	if obj.obj.BottomOfStack != nil {

		obj.BottomOfStack().validateObj(vObj, set_default)
	}

	if obj.obj.TimeToLive != nil {

		obj.TimeToLive().validateObj(vObj, set_default)
	}

}

func (obj *flowMpls) setDefault() {

}
