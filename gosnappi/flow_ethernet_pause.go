package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowEthernetPause *****
type flowEthernetPause struct {
	validation
	obj                 *otg.FlowEthernetPause
	marshaller          marshalFlowEthernetPause
	unMarshaller        unMarshalFlowEthernetPause
	dstHolder           PatternFlowEthernetPauseDst
	srcHolder           PatternFlowEthernetPauseSrc
	etherTypeHolder     PatternFlowEthernetPauseEtherType
	controlOpCodeHolder PatternFlowEthernetPauseControlOpCode
	timeHolder          PatternFlowEthernetPauseTime
}

func NewFlowEthernetPause() FlowEthernetPause {
	obj := flowEthernetPause{obj: &otg.FlowEthernetPause{}}
	obj.setDefault()
	return &obj
}

func (obj *flowEthernetPause) msg() *otg.FlowEthernetPause {
	return obj.obj
}

func (obj *flowEthernetPause) setMsg(msg *otg.FlowEthernetPause) FlowEthernetPause {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowEthernetPause struct {
	obj *flowEthernetPause
}

type marshalFlowEthernetPause interface {
	// ToProto marshals FlowEthernetPause to protobuf object *otg.FlowEthernetPause
	ToProto() (*otg.FlowEthernetPause, error)
	// ToPbText marshals FlowEthernetPause to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowEthernetPause to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowEthernetPause to JSON text
	ToJson() (string, error)
}

type unMarshalflowEthernetPause struct {
	obj *flowEthernetPause
}

type unMarshalFlowEthernetPause interface {
	// FromProto unmarshals FlowEthernetPause from protobuf object *otg.FlowEthernetPause
	FromProto(msg *otg.FlowEthernetPause) (FlowEthernetPause, error)
	// FromPbText unmarshals FlowEthernetPause from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowEthernetPause from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowEthernetPause from JSON text
	FromJson(value string) error
}

func (obj *flowEthernetPause) Marshal() marshalFlowEthernetPause {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowEthernetPause{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowEthernetPause) Unmarshal() unMarshalFlowEthernetPause {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowEthernetPause{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowEthernetPause) ToProto() (*otg.FlowEthernetPause, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowEthernetPause) FromProto(msg *otg.FlowEthernetPause) (FlowEthernetPause, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowEthernetPause) ToPbText() (string, error) {
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

func (m *unMarshalflowEthernetPause) FromPbText(value string) error {
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

func (m *marshalflowEthernetPause) ToYaml() (string, error) {
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

func (m *unMarshalflowEthernetPause) FromYaml(value string) error {
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

func (m *marshalflowEthernetPause) ToJson() (string, error) {
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

func (m *unMarshalflowEthernetPause) FromJson(value string) error {
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

func (obj *flowEthernetPause) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowEthernetPause) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowEthernetPause) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowEthernetPause) Clone() (FlowEthernetPause, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowEthernetPause()
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

func (obj *flowEthernetPause) setNil() {
	obj.dstHolder = nil
	obj.srcHolder = nil
	obj.etherTypeHolder = nil
	obj.controlOpCodeHolder = nil
	obj.timeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowEthernetPause is iEEE 802.3x global ethernet pause packet header
type FlowEthernetPause interface {
	Validation
	// msg marshals FlowEthernetPause to protobuf object *otg.FlowEthernetPause
	// and doesn't set defaults
	msg() *otg.FlowEthernetPause
	// setMsg unmarshals FlowEthernetPause from protobuf object *otg.FlowEthernetPause
	// and doesn't set defaults
	setMsg(*otg.FlowEthernetPause) FlowEthernetPause
	// provides marshal interface
	Marshal() marshalFlowEthernetPause
	// provides unmarshal interface
	Unmarshal() unMarshalFlowEthernetPause
	// validate validates FlowEthernetPause
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowEthernetPause, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Dst returns PatternFlowEthernetPauseDst, set in FlowEthernetPause.
	// PatternFlowEthernetPauseDst is destination MAC address
	Dst() PatternFlowEthernetPauseDst
	// SetDst assigns PatternFlowEthernetPauseDst provided by user to FlowEthernetPause.
	// PatternFlowEthernetPauseDst is destination MAC address
	SetDst(value PatternFlowEthernetPauseDst) FlowEthernetPause
	// HasDst checks if Dst has been set in FlowEthernetPause
	HasDst() bool
	// Src returns PatternFlowEthernetPauseSrc, set in FlowEthernetPause.
	// PatternFlowEthernetPauseSrc is source MAC address
	Src() PatternFlowEthernetPauseSrc
	// SetSrc assigns PatternFlowEthernetPauseSrc provided by user to FlowEthernetPause.
	// PatternFlowEthernetPauseSrc is source MAC address
	SetSrc(value PatternFlowEthernetPauseSrc) FlowEthernetPause
	// HasSrc checks if Src has been set in FlowEthernetPause
	HasSrc() bool
	// EtherType returns PatternFlowEthernetPauseEtherType, set in FlowEthernetPause.
	// PatternFlowEthernetPauseEtherType is ethernet type
	EtherType() PatternFlowEthernetPauseEtherType
	// SetEtherType assigns PatternFlowEthernetPauseEtherType provided by user to FlowEthernetPause.
	// PatternFlowEthernetPauseEtherType is ethernet type
	SetEtherType(value PatternFlowEthernetPauseEtherType) FlowEthernetPause
	// HasEtherType checks if EtherType has been set in FlowEthernetPause
	HasEtherType() bool
	// ControlOpCode returns PatternFlowEthernetPauseControlOpCode, set in FlowEthernetPause.
	// PatternFlowEthernetPauseControlOpCode is control operation code
	ControlOpCode() PatternFlowEthernetPauseControlOpCode
	// SetControlOpCode assigns PatternFlowEthernetPauseControlOpCode provided by user to FlowEthernetPause.
	// PatternFlowEthernetPauseControlOpCode is control operation code
	SetControlOpCode(value PatternFlowEthernetPauseControlOpCode) FlowEthernetPause
	// HasControlOpCode checks if ControlOpCode has been set in FlowEthernetPause
	HasControlOpCode() bool
	// Time returns PatternFlowEthernetPauseTime, set in FlowEthernetPause.
	// PatternFlowEthernetPauseTime is time
	Time() PatternFlowEthernetPauseTime
	// SetTime assigns PatternFlowEthernetPauseTime provided by user to FlowEthernetPause.
	// PatternFlowEthernetPauseTime is time
	SetTime(value PatternFlowEthernetPauseTime) FlowEthernetPause
	// HasTime checks if Time has been set in FlowEthernetPause
	HasTime() bool
	setNil()
}

// description is TBD
// Dst returns a PatternFlowEthernetPauseDst
func (obj *flowEthernetPause) Dst() PatternFlowEthernetPauseDst {
	if obj.obj.Dst == nil {
		obj.obj.Dst = NewPatternFlowEthernetPauseDst().msg()
	}
	if obj.dstHolder == nil {
		obj.dstHolder = &patternFlowEthernetPauseDst{obj: obj.obj.Dst}
	}
	return obj.dstHolder
}

// description is TBD
// Dst returns a PatternFlowEthernetPauseDst
func (obj *flowEthernetPause) HasDst() bool {
	return obj.obj.Dst != nil
}

// description is TBD
// SetDst sets the PatternFlowEthernetPauseDst value in the FlowEthernetPause object
func (obj *flowEthernetPause) SetDst(value PatternFlowEthernetPauseDst) FlowEthernetPause {

	obj.dstHolder = nil
	obj.obj.Dst = value.msg()

	return obj
}

// description is TBD
// Src returns a PatternFlowEthernetPauseSrc
func (obj *flowEthernetPause) Src() PatternFlowEthernetPauseSrc {
	if obj.obj.Src == nil {
		obj.obj.Src = NewPatternFlowEthernetPauseSrc().msg()
	}
	if obj.srcHolder == nil {
		obj.srcHolder = &patternFlowEthernetPauseSrc{obj: obj.obj.Src}
	}
	return obj.srcHolder
}

// description is TBD
// Src returns a PatternFlowEthernetPauseSrc
func (obj *flowEthernetPause) HasSrc() bool {
	return obj.obj.Src != nil
}

// description is TBD
// SetSrc sets the PatternFlowEthernetPauseSrc value in the FlowEthernetPause object
func (obj *flowEthernetPause) SetSrc(value PatternFlowEthernetPauseSrc) FlowEthernetPause {

	obj.srcHolder = nil
	obj.obj.Src = value.msg()

	return obj
}

// description is TBD
// EtherType returns a PatternFlowEthernetPauseEtherType
func (obj *flowEthernetPause) EtherType() PatternFlowEthernetPauseEtherType {
	if obj.obj.EtherType == nil {
		obj.obj.EtherType = NewPatternFlowEthernetPauseEtherType().msg()
	}
	if obj.etherTypeHolder == nil {
		obj.etherTypeHolder = &patternFlowEthernetPauseEtherType{obj: obj.obj.EtherType}
	}
	return obj.etherTypeHolder
}

// description is TBD
// EtherType returns a PatternFlowEthernetPauseEtherType
func (obj *flowEthernetPause) HasEtherType() bool {
	return obj.obj.EtherType != nil
}

// description is TBD
// SetEtherType sets the PatternFlowEthernetPauseEtherType value in the FlowEthernetPause object
func (obj *flowEthernetPause) SetEtherType(value PatternFlowEthernetPauseEtherType) FlowEthernetPause {

	obj.etherTypeHolder = nil
	obj.obj.EtherType = value.msg()

	return obj
}

// description is TBD
// ControlOpCode returns a PatternFlowEthernetPauseControlOpCode
func (obj *flowEthernetPause) ControlOpCode() PatternFlowEthernetPauseControlOpCode {
	if obj.obj.ControlOpCode == nil {
		obj.obj.ControlOpCode = NewPatternFlowEthernetPauseControlOpCode().msg()
	}
	if obj.controlOpCodeHolder == nil {
		obj.controlOpCodeHolder = &patternFlowEthernetPauseControlOpCode{obj: obj.obj.ControlOpCode}
	}
	return obj.controlOpCodeHolder
}

// description is TBD
// ControlOpCode returns a PatternFlowEthernetPauseControlOpCode
func (obj *flowEthernetPause) HasControlOpCode() bool {
	return obj.obj.ControlOpCode != nil
}

// description is TBD
// SetControlOpCode sets the PatternFlowEthernetPauseControlOpCode value in the FlowEthernetPause object
func (obj *flowEthernetPause) SetControlOpCode(value PatternFlowEthernetPauseControlOpCode) FlowEthernetPause {

	obj.controlOpCodeHolder = nil
	obj.obj.ControlOpCode = value.msg()

	return obj
}

// description is TBD
// Time returns a PatternFlowEthernetPauseTime
func (obj *flowEthernetPause) Time() PatternFlowEthernetPauseTime {
	if obj.obj.Time == nil {
		obj.obj.Time = NewPatternFlowEthernetPauseTime().msg()
	}
	if obj.timeHolder == nil {
		obj.timeHolder = &patternFlowEthernetPauseTime{obj: obj.obj.Time}
	}
	return obj.timeHolder
}

// description is TBD
// Time returns a PatternFlowEthernetPauseTime
func (obj *flowEthernetPause) HasTime() bool {
	return obj.obj.Time != nil
}

// description is TBD
// SetTime sets the PatternFlowEthernetPauseTime value in the FlowEthernetPause object
func (obj *flowEthernetPause) SetTime(value PatternFlowEthernetPauseTime) FlowEthernetPause {

	obj.timeHolder = nil
	obj.obj.Time = value.msg()

	return obj
}

func (obj *flowEthernetPause) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Dst != nil {

		obj.Dst().validateObj(vObj, set_default)
	}

	if obj.obj.Src != nil {

		obj.Src().validateObj(vObj, set_default)
	}

	if obj.obj.EtherType != nil {

		obj.EtherType().validateObj(vObj, set_default)
	}

	if obj.obj.ControlOpCode != nil {

		obj.ControlOpCode().validateObj(vObj, set_default)
	}

	if obj.obj.Time != nil {

		obj.Time().validateObj(vObj, set_default)
	}

}

func (obj *flowEthernetPause) setDefault() {

}
