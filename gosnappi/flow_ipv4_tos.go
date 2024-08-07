package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv4Tos *****
type flowIpv4Tos struct {
	validation
	obj               *otg.FlowIpv4Tos
	marshaller        marshalFlowIpv4Tos
	unMarshaller      unMarshalFlowIpv4Tos
	precedenceHolder  PatternFlowIpv4TosPrecedence
	delayHolder       PatternFlowIpv4TosDelay
	throughputHolder  PatternFlowIpv4TosThroughput
	reliabilityHolder PatternFlowIpv4TosReliability
	monetaryHolder    PatternFlowIpv4TosMonetary
	unusedHolder      PatternFlowIpv4TosUnused
}

func NewFlowIpv4Tos() FlowIpv4Tos {
	obj := flowIpv4Tos{obj: &otg.FlowIpv4Tos{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv4Tos) msg() *otg.FlowIpv4Tos {
	return obj.obj
}

func (obj *flowIpv4Tos) setMsg(msg *otg.FlowIpv4Tos) FlowIpv4Tos {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv4Tos struct {
	obj *flowIpv4Tos
}

type marshalFlowIpv4Tos interface {
	// ToProto marshals FlowIpv4Tos to protobuf object *otg.FlowIpv4Tos
	ToProto() (*otg.FlowIpv4Tos, error)
	// ToPbText marshals FlowIpv4Tos to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv4Tos to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv4Tos to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv4Tos struct {
	obj *flowIpv4Tos
}

type unMarshalFlowIpv4Tos interface {
	// FromProto unmarshals FlowIpv4Tos from protobuf object *otg.FlowIpv4Tos
	FromProto(msg *otg.FlowIpv4Tos) (FlowIpv4Tos, error)
	// FromPbText unmarshals FlowIpv4Tos from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv4Tos from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv4Tos from JSON text
	FromJson(value string) error
}

func (obj *flowIpv4Tos) Marshal() marshalFlowIpv4Tos {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv4Tos{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv4Tos) Unmarshal() unMarshalFlowIpv4Tos {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv4Tos{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv4Tos) ToProto() (*otg.FlowIpv4Tos, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv4Tos) FromProto(msg *otg.FlowIpv4Tos) (FlowIpv4Tos, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv4Tos) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv4Tos) FromPbText(value string) error {
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

func (m *marshalflowIpv4Tos) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv4Tos) FromYaml(value string) error {
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

func (m *marshalflowIpv4Tos) ToJson() (string, error) {
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

func (m *unMarshalflowIpv4Tos) FromJson(value string) error {
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

func (obj *flowIpv4Tos) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv4Tos) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv4Tos) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv4Tos) Clone() (FlowIpv4Tos, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv4Tos()
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

func (obj *flowIpv4Tos) setNil() {
	obj.precedenceHolder = nil
	obj.delayHolder = nil
	obj.throughputHolder = nil
	obj.reliabilityHolder = nil
	obj.monetaryHolder = nil
	obj.unusedHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv4Tos is type of service (TOS) packet field.
type FlowIpv4Tos interface {
	Validation
	// msg marshals FlowIpv4Tos to protobuf object *otg.FlowIpv4Tos
	// and doesn't set defaults
	msg() *otg.FlowIpv4Tos
	// setMsg unmarshals FlowIpv4Tos from protobuf object *otg.FlowIpv4Tos
	// and doesn't set defaults
	setMsg(*otg.FlowIpv4Tos) FlowIpv4Tos
	// provides marshal interface
	Marshal() marshalFlowIpv4Tos
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv4Tos
	// validate validates FlowIpv4Tos
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv4Tos, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Precedence returns PatternFlowIpv4TosPrecedence, set in FlowIpv4Tos.
	// PatternFlowIpv4TosPrecedence is precedence
	Precedence() PatternFlowIpv4TosPrecedence
	// SetPrecedence assigns PatternFlowIpv4TosPrecedence provided by user to FlowIpv4Tos.
	// PatternFlowIpv4TosPrecedence is precedence
	SetPrecedence(value PatternFlowIpv4TosPrecedence) FlowIpv4Tos
	// HasPrecedence checks if Precedence has been set in FlowIpv4Tos
	HasPrecedence() bool
	// Delay returns PatternFlowIpv4TosDelay, set in FlowIpv4Tos.
	// PatternFlowIpv4TosDelay is delay
	Delay() PatternFlowIpv4TosDelay
	// SetDelay assigns PatternFlowIpv4TosDelay provided by user to FlowIpv4Tos.
	// PatternFlowIpv4TosDelay is delay
	SetDelay(value PatternFlowIpv4TosDelay) FlowIpv4Tos
	// HasDelay checks if Delay has been set in FlowIpv4Tos
	HasDelay() bool
	// Throughput returns PatternFlowIpv4TosThroughput, set in FlowIpv4Tos.
	// PatternFlowIpv4TosThroughput is throughput
	Throughput() PatternFlowIpv4TosThroughput
	// SetThroughput assigns PatternFlowIpv4TosThroughput provided by user to FlowIpv4Tos.
	// PatternFlowIpv4TosThroughput is throughput
	SetThroughput(value PatternFlowIpv4TosThroughput) FlowIpv4Tos
	// HasThroughput checks if Throughput has been set in FlowIpv4Tos
	HasThroughput() bool
	// Reliability returns PatternFlowIpv4TosReliability, set in FlowIpv4Tos.
	// PatternFlowIpv4TosReliability is reliability
	Reliability() PatternFlowIpv4TosReliability
	// SetReliability assigns PatternFlowIpv4TosReliability provided by user to FlowIpv4Tos.
	// PatternFlowIpv4TosReliability is reliability
	SetReliability(value PatternFlowIpv4TosReliability) FlowIpv4Tos
	// HasReliability checks if Reliability has been set in FlowIpv4Tos
	HasReliability() bool
	// Monetary returns PatternFlowIpv4TosMonetary, set in FlowIpv4Tos.
	// PatternFlowIpv4TosMonetary is monetary
	Monetary() PatternFlowIpv4TosMonetary
	// SetMonetary assigns PatternFlowIpv4TosMonetary provided by user to FlowIpv4Tos.
	// PatternFlowIpv4TosMonetary is monetary
	SetMonetary(value PatternFlowIpv4TosMonetary) FlowIpv4Tos
	// HasMonetary checks if Monetary has been set in FlowIpv4Tos
	HasMonetary() bool
	// Unused returns PatternFlowIpv4TosUnused, set in FlowIpv4Tos.
	// PatternFlowIpv4TosUnused is unused
	Unused() PatternFlowIpv4TosUnused
	// SetUnused assigns PatternFlowIpv4TosUnused provided by user to FlowIpv4Tos.
	// PatternFlowIpv4TosUnused is unused
	SetUnused(value PatternFlowIpv4TosUnused) FlowIpv4Tos
	// HasUnused checks if Unused has been set in FlowIpv4Tos
	HasUnused() bool
	setNil()
}

// description is TBD
// Precedence returns a PatternFlowIpv4TosPrecedence
func (obj *flowIpv4Tos) Precedence() PatternFlowIpv4TosPrecedence {
	if obj.obj.Precedence == nil {
		obj.obj.Precedence = NewPatternFlowIpv4TosPrecedence().msg()
	}
	if obj.precedenceHolder == nil {
		obj.precedenceHolder = &patternFlowIpv4TosPrecedence{obj: obj.obj.Precedence}
	}
	return obj.precedenceHolder
}

// description is TBD
// Precedence returns a PatternFlowIpv4TosPrecedence
func (obj *flowIpv4Tos) HasPrecedence() bool {
	return obj.obj.Precedence != nil
}

// description is TBD
// SetPrecedence sets the PatternFlowIpv4TosPrecedence value in the FlowIpv4Tos object
func (obj *flowIpv4Tos) SetPrecedence(value PatternFlowIpv4TosPrecedence) FlowIpv4Tos {

	obj.precedenceHolder = nil
	obj.obj.Precedence = value.msg()

	return obj
}

// description is TBD
// Delay returns a PatternFlowIpv4TosDelay
func (obj *flowIpv4Tos) Delay() PatternFlowIpv4TosDelay {
	if obj.obj.Delay == nil {
		obj.obj.Delay = NewPatternFlowIpv4TosDelay().msg()
	}
	if obj.delayHolder == nil {
		obj.delayHolder = &patternFlowIpv4TosDelay{obj: obj.obj.Delay}
	}
	return obj.delayHolder
}

// description is TBD
// Delay returns a PatternFlowIpv4TosDelay
func (obj *flowIpv4Tos) HasDelay() bool {
	return obj.obj.Delay != nil
}

// description is TBD
// SetDelay sets the PatternFlowIpv4TosDelay value in the FlowIpv4Tos object
func (obj *flowIpv4Tos) SetDelay(value PatternFlowIpv4TosDelay) FlowIpv4Tos {

	obj.delayHolder = nil
	obj.obj.Delay = value.msg()

	return obj
}

// description is TBD
// Throughput returns a PatternFlowIpv4TosThroughput
func (obj *flowIpv4Tos) Throughput() PatternFlowIpv4TosThroughput {
	if obj.obj.Throughput == nil {
		obj.obj.Throughput = NewPatternFlowIpv4TosThroughput().msg()
	}
	if obj.throughputHolder == nil {
		obj.throughputHolder = &patternFlowIpv4TosThroughput{obj: obj.obj.Throughput}
	}
	return obj.throughputHolder
}

// description is TBD
// Throughput returns a PatternFlowIpv4TosThroughput
func (obj *flowIpv4Tos) HasThroughput() bool {
	return obj.obj.Throughput != nil
}

// description is TBD
// SetThroughput sets the PatternFlowIpv4TosThroughput value in the FlowIpv4Tos object
func (obj *flowIpv4Tos) SetThroughput(value PatternFlowIpv4TosThroughput) FlowIpv4Tos {

	obj.throughputHolder = nil
	obj.obj.Throughput = value.msg()

	return obj
}

// description is TBD
// Reliability returns a PatternFlowIpv4TosReliability
func (obj *flowIpv4Tos) Reliability() PatternFlowIpv4TosReliability {
	if obj.obj.Reliability == nil {
		obj.obj.Reliability = NewPatternFlowIpv4TosReliability().msg()
	}
	if obj.reliabilityHolder == nil {
		obj.reliabilityHolder = &patternFlowIpv4TosReliability{obj: obj.obj.Reliability}
	}
	return obj.reliabilityHolder
}

// description is TBD
// Reliability returns a PatternFlowIpv4TosReliability
func (obj *flowIpv4Tos) HasReliability() bool {
	return obj.obj.Reliability != nil
}

// description is TBD
// SetReliability sets the PatternFlowIpv4TosReliability value in the FlowIpv4Tos object
func (obj *flowIpv4Tos) SetReliability(value PatternFlowIpv4TosReliability) FlowIpv4Tos {

	obj.reliabilityHolder = nil
	obj.obj.Reliability = value.msg()

	return obj
}

// description is TBD
// Monetary returns a PatternFlowIpv4TosMonetary
func (obj *flowIpv4Tos) Monetary() PatternFlowIpv4TosMonetary {
	if obj.obj.Monetary == nil {
		obj.obj.Monetary = NewPatternFlowIpv4TosMonetary().msg()
	}
	if obj.monetaryHolder == nil {
		obj.monetaryHolder = &patternFlowIpv4TosMonetary{obj: obj.obj.Monetary}
	}
	return obj.monetaryHolder
}

// description is TBD
// Monetary returns a PatternFlowIpv4TosMonetary
func (obj *flowIpv4Tos) HasMonetary() bool {
	return obj.obj.Monetary != nil
}

// description is TBD
// SetMonetary sets the PatternFlowIpv4TosMonetary value in the FlowIpv4Tos object
func (obj *flowIpv4Tos) SetMonetary(value PatternFlowIpv4TosMonetary) FlowIpv4Tos {

	obj.monetaryHolder = nil
	obj.obj.Monetary = value.msg()

	return obj
}

// description is TBD
// Unused returns a PatternFlowIpv4TosUnused
func (obj *flowIpv4Tos) Unused() PatternFlowIpv4TosUnused {
	if obj.obj.Unused == nil {
		obj.obj.Unused = NewPatternFlowIpv4TosUnused().msg()
	}
	if obj.unusedHolder == nil {
		obj.unusedHolder = &patternFlowIpv4TosUnused{obj: obj.obj.Unused}
	}
	return obj.unusedHolder
}

// description is TBD
// Unused returns a PatternFlowIpv4TosUnused
func (obj *flowIpv4Tos) HasUnused() bool {
	return obj.obj.Unused != nil
}

// description is TBD
// SetUnused sets the PatternFlowIpv4TosUnused value in the FlowIpv4Tos object
func (obj *flowIpv4Tos) SetUnused(value PatternFlowIpv4TosUnused) FlowIpv4Tos {

	obj.unusedHolder = nil
	obj.obj.Unused = value.msg()

	return obj
}

func (obj *flowIpv4Tos) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Precedence != nil {

		obj.Precedence().validateObj(vObj, set_default)
	}

	if obj.obj.Delay != nil {

		obj.Delay().validateObj(vObj, set_default)
	}

	if obj.obj.Throughput != nil {

		obj.Throughput().validateObj(vObj, set_default)
	}

	if obj.obj.Reliability != nil {

		obj.Reliability().validateObj(vObj, set_default)
	}

	if obj.obj.Monetary != nil {

		obj.Monetary().validateObj(vObj, set_default)
	}

	if obj.obj.Unused != nil {

		obj.Unused().validateObj(vObj, set_default)
	}

}

func (obj *flowIpv4Tos) setDefault() {

}
