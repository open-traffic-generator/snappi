package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathLabelRequestWithoutLabelRange *****
type flowRSVPPathLabelRequestWithoutLabelRange struct {
	validation
	obj            *otg.FlowRSVPPathLabelRequestWithoutLabelRange
	marshaller     marshalFlowRSVPPathLabelRequestWithoutLabelRange
	unMarshaller   unMarshalFlowRSVPPathLabelRequestWithoutLabelRange
	reservedHolder PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
	l3PidHolder    PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
}

func NewFlowRSVPPathLabelRequestWithoutLabelRange() FlowRSVPPathLabelRequestWithoutLabelRange {
	obj := flowRSVPPathLabelRequestWithoutLabelRange{obj: &otg.FlowRSVPPathLabelRequestWithoutLabelRange{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathLabelRequestWithoutLabelRange) msg() *otg.FlowRSVPPathLabelRequestWithoutLabelRange {
	return obj.obj
}

func (obj *flowRSVPPathLabelRequestWithoutLabelRange) setMsg(msg *otg.FlowRSVPPathLabelRequestWithoutLabelRange) FlowRSVPPathLabelRequestWithoutLabelRange {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathLabelRequestWithoutLabelRange struct {
	obj *flowRSVPPathLabelRequestWithoutLabelRange
}

type marshalFlowRSVPPathLabelRequestWithoutLabelRange interface {
	// ToProto marshals FlowRSVPPathLabelRequestWithoutLabelRange to protobuf object *otg.FlowRSVPPathLabelRequestWithoutLabelRange
	ToProto() (*otg.FlowRSVPPathLabelRequestWithoutLabelRange, error)
	// ToPbText marshals FlowRSVPPathLabelRequestWithoutLabelRange to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathLabelRequestWithoutLabelRange to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathLabelRequestWithoutLabelRange to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowRSVPPathLabelRequestWithoutLabelRange to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowRSVPPathLabelRequestWithoutLabelRange struct {
	obj *flowRSVPPathLabelRequestWithoutLabelRange
}

type unMarshalFlowRSVPPathLabelRequestWithoutLabelRange interface {
	// FromProto unmarshals FlowRSVPPathLabelRequestWithoutLabelRange from protobuf object *otg.FlowRSVPPathLabelRequestWithoutLabelRange
	FromProto(msg *otg.FlowRSVPPathLabelRequestWithoutLabelRange) (FlowRSVPPathLabelRequestWithoutLabelRange, error)
	// FromPbText unmarshals FlowRSVPPathLabelRequestWithoutLabelRange from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathLabelRequestWithoutLabelRange from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathLabelRequestWithoutLabelRange from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathLabelRequestWithoutLabelRange) Marshal() marshalFlowRSVPPathLabelRequestWithoutLabelRange {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathLabelRequestWithoutLabelRange{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathLabelRequestWithoutLabelRange) Unmarshal() unMarshalFlowRSVPPathLabelRequestWithoutLabelRange {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathLabelRequestWithoutLabelRange{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathLabelRequestWithoutLabelRange) ToProto() (*otg.FlowRSVPPathLabelRequestWithoutLabelRange, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathLabelRequestWithoutLabelRange) FromProto(msg *otg.FlowRSVPPathLabelRequestWithoutLabelRange) (FlowRSVPPathLabelRequestWithoutLabelRange, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathLabelRequestWithoutLabelRange) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathLabelRequestWithoutLabelRange) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathLabelRequestWithoutLabelRange) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathLabelRequestWithoutLabelRange) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathLabelRequestWithoutLabelRange) ToJsonRaw() (string, error) {
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

func (m *marshalflowRSVPPathLabelRequestWithoutLabelRange) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathLabelRequestWithoutLabelRange) FromJson(value string) error {
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

func (obj *flowRSVPPathLabelRequestWithoutLabelRange) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathLabelRequestWithoutLabelRange) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathLabelRequestWithoutLabelRange) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathLabelRequestWithoutLabelRange) Clone() (FlowRSVPPathLabelRequestWithoutLabelRange, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathLabelRequestWithoutLabelRange()
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

func (obj *flowRSVPPathLabelRequestWithoutLabelRange) setNil() {
	obj.reservedHolder = nil
	obj.l3PidHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathLabelRequestWithoutLabelRange is class = LABEL_REQUEST, Without Label Range C-Type = 1
type FlowRSVPPathLabelRequestWithoutLabelRange interface {
	Validation
	// msg marshals FlowRSVPPathLabelRequestWithoutLabelRange to protobuf object *otg.FlowRSVPPathLabelRequestWithoutLabelRange
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathLabelRequestWithoutLabelRange
	// setMsg unmarshals FlowRSVPPathLabelRequestWithoutLabelRange from protobuf object *otg.FlowRSVPPathLabelRequestWithoutLabelRange
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathLabelRequestWithoutLabelRange) FlowRSVPPathLabelRequestWithoutLabelRange
	// provides marshal interface
	Marshal() marshalFlowRSVPPathLabelRequestWithoutLabelRange
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathLabelRequestWithoutLabelRange
	// validate validates FlowRSVPPathLabelRequestWithoutLabelRange
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathLabelRequestWithoutLabelRange, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Reserved returns PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved, set in FlowRSVPPathLabelRequestWithoutLabelRange.
	// PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved is this field is reserved.   It MUST be set to zero on transmission and MUST be ignored on receipt.
	Reserved() PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
	// SetReserved assigns PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved provided by user to FlowRSVPPathLabelRequestWithoutLabelRange.
	// PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved is this field is reserved.   It MUST be set to zero on transmission and MUST be ignored on receipt.
	SetReserved(value PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) FlowRSVPPathLabelRequestWithoutLabelRange
	// HasReserved checks if Reserved has been set in FlowRSVPPathLabelRequestWithoutLabelRange
	HasReserved() bool
	// L3Pid returns PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid, set in FlowRSVPPathLabelRequestWithoutLabelRange.
	L3Pid() PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
	// SetL3Pid assigns PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid provided by user to FlowRSVPPathLabelRequestWithoutLabelRange.
	SetL3Pid(value PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) FlowRSVPPathLabelRequestWithoutLabelRange
	// HasL3Pid checks if L3Pid has been set in FlowRSVPPathLabelRequestWithoutLabelRange
	HasL3Pid() bool
	setNil()
}

// description is TBD
// Reserved returns a PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
func (obj *flowRSVPPathLabelRequestWithoutLabelRange) Reserved() PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved {
	if obj.obj.Reserved == nil {
		obj.obj.Reserved = NewPatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved().msg()
	}
	if obj.reservedHolder == nil {
		obj.reservedHolder = &patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved{obj: obj.obj.Reserved}
	}
	return obj.reservedHolder
}

// description is TBD
// Reserved returns a PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
func (obj *flowRSVPPathLabelRequestWithoutLabelRange) HasReserved() bool {
	return obj.obj.Reserved != nil
}

// description is TBD
// SetReserved sets the PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved value in the FlowRSVPPathLabelRequestWithoutLabelRange object
func (obj *flowRSVPPathLabelRequestWithoutLabelRange) SetReserved(value PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) FlowRSVPPathLabelRequestWithoutLabelRange {

	obj.reservedHolder = nil
	obj.obj.Reserved = value.msg()

	return obj
}

// description is TBD
// L3Pid returns a PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
func (obj *flowRSVPPathLabelRequestWithoutLabelRange) L3Pid() PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid {
	if obj.obj.L3Pid == nil {
		obj.obj.L3Pid = NewPatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid().msg()
	}
	if obj.l3PidHolder == nil {
		obj.l3PidHolder = &patternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid{obj: obj.obj.L3Pid}
	}
	return obj.l3PidHolder
}

// description is TBD
// L3Pid returns a PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid
func (obj *flowRSVPPathLabelRequestWithoutLabelRange) HasL3Pid() bool {
	return obj.obj.L3Pid != nil
}

// description is TBD
// SetL3Pid sets the PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid value in the FlowRSVPPathLabelRequestWithoutLabelRange object
func (obj *flowRSVPPathLabelRequestWithoutLabelRange) SetL3Pid(value PatternFlowRSVPPathLabelRequestWithoutLabelRangeL3Pid) FlowRSVPPathLabelRequestWithoutLabelRange {

	obj.l3PidHolder = nil
	obj.obj.L3Pid = value.msg()

	return obj
}

func (obj *flowRSVPPathLabelRequestWithoutLabelRange) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Reserved != nil {

		obj.Reserved().validateObj(vObj, set_default)
	}

	if obj.obj.L3Pid != nil {

		obj.L3Pid().validateObj(vObj, set_default)
	}

}

func (obj *flowRSVPPathLabelRequestWithoutLabelRange) setDefault() {

}
