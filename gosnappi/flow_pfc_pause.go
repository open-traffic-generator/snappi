package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowPfcPause *****
type flowPfcPause struct {
	validation
	obj                     *otg.FlowPfcPause
	marshaller              marshalFlowPfcPause
	unMarshaller            unMarshalFlowPfcPause
	dstHolder               PatternFlowPfcPauseDst
	srcHolder               PatternFlowPfcPauseSrc
	etherTypeHolder         PatternFlowPfcPauseEtherType
	controlOpCodeHolder     PatternFlowPfcPauseControlOpCode
	classEnableVectorHolder PatternFlowPfcPauseClassEnableVector
	pauseClass_0Holder      PatternFlowPfcPausePauseClass0
	pauseClass_1Holder      PatternFlowPfcPausePauseClass1
	pauseClass_2Holder      PatternFlowPfcPausePauseClass2
	pauseClass_3Holder      PatternFlowPfcPausePauseClass3
	pauseClass_4Holder      PatternFlowPfcPausePauseClass4
	pauseClass_5Holder      PatternFlowPfcPausePauseClass5
	pauseClass_6Holder      PatternFlowPfcPausePauseClass6
	pauseClass_7Holder      PatternFlowPfcPausePauseClass7
}

func NewFlowPfcPause() FlowPfcPause {
	obj := flowPfcPause{obj: &otg.FlowPfcPause{}}
	obj.setDefault()
	return &obj
}

func (obj *flowPfcPause) msg() *otg.FlowPfcPause {
	return obj.obj
}

func (obj *flowPfcPause) setMsg(msg *otg.FlowPfcPause) FlowPfcPause {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowPfcPause struct {
	obj *flowPfcPause
}

type marshalFlowPfcPause interface {
	// ToProto marshals FlowPfcPause to protobuf object *otg.FlowPfcPause
	ToProto() (*otg.FlowPfcPause, error)
	// ToPbText marshals FlowPfcPause to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowPfcPause to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowPfcPause to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowPfcPause to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowPfcPause struct {
	obj *flowPfcPause
}

type unMarshalFlowPfcPause interface {
	// FromProto unmarshals FlowPfcPause from protobuf object *otg.FlowPfcPause
	FromProto(msg *otg.FlowPfcPause) (FlowPfcPause, error)
	// FromPbText unmarshals FlowPfcPause from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowPfcPause from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowPfcPause from JSON text
	FromJson(value string) error
}

func (obj *flowPfcPause) Marshal() marshalFlowPfcPause {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowPfcPause{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowPfcPause) Unmarshal() unMarshalFlowPfcPause {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowPfcPause{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowPfcPause) ToProto() (*otg.FlowPfcPause, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowPfcPause) FromProto(msg *otg.FlowPfcPause) (FlowPfcPause, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowPfcPause) ToPbText() (string, error) {
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

func (m *unMarshalflowPfcPause) FromPbText(value string) error {
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

func (m *marshalflowPfcPause) ToYaml() (string, error) {
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

func (m *unMarshalflowPfcPause) FromYaml(value string) error {
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

func (m *marshalflowPfcPause) ToJsonRaw() (string, error) {
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

func (m *marshalflowPfcPause) ToJson() (string, error) {
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

func (m *unMarshalflowPfcPause) FromJson(value string) error {
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

func (obj *flowPfcPause) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowPfcPause) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowPfcPause) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowPfcPause) Clone() (FlowPfcPause, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowPfcPause()
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

func (obj *flowPfcPause) setNil() {
	obj.dstHolder = nil
	obj.srcHolder = nil
	obj.etherTypeHolder = nil
	obj.controlOpCodeHolder = nil
	obj.classEnableVectorHolder = nil
	obj.pauseClass_0Holder = nil
	obj.pauseClass_1Holder = nil
	obj.pauseClass_2Holder = nil
	obj.pauseClass_3Holder = nil
	obj.pauseClass_4Holder = nil
	obj.pauseClass_5Holder = nil
	obj.pauseClass_6Holder = nil
	obj.pauseClass_7Holder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowPfcPause is iEEE 802.1Qbb PFC Pause packet header.
type FlowPfcPause interface {
	Validation
	// msg marshals FlowPfcPause to protobuf object *otg.FlowPfcPause
	// and doesn't set defaults
	msg() *otg.FlowPfcPause
	// setMsg unmarshals FlowPfcPause from protobuf object *otg.FlowPfcPause
	// and doesn't set defaults
	setMsg(*otg.FlowPfcPause) FlowPfcPause
	// provides marshal interface
	Marshal() marshalFlowPfcPause
	// provides unmarshal interface
	Unmarshal() unMarshalFlowPfcPause
	// validate validates FlowPfcPause
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowPfcPause, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Dst returns PatternFlowPfcPauseDst, set in FlowPfcPause.
	// PatternFlowPfcPauseDst is destination MAC address
	Dst() PatternFlowPfcPauseDst
	// SetDst assigns PatternFlowPfcPauseDst provided by user to FlowPfcPause.
	// PatternFlowPfcPauseDst is destination MAC address
	SetDst(value PatternFlowPfcPauseDst) FlowPfcPause
	// HasDst checks if Dst has been set in FlowPfcPause
	HasDst() bool
	// Src returns PatternFlowPfcPauseSrc, set in FlowPfcPause.
	// PatternFlowPfcPauseSrc is source MAC address
	Src() PatternFlowPfcPauseSrc
	// SetSrc assigns PatternFlowPfcPauseSrc provided by user to FlowPfcPause.
	// PatternFlowPfcPauseSrc is source MAC address
	SetSrc(value PatternFlowPfcPauseSrc) FlowPfcPause
	// HasSrc checks if Src has been set in FlowPfcPause
	HasSrc() bool
	// EtherType returns PatternFlowPfcPauseEtherType, set in FlowPfcPause.
	// PatternFlowPfcPauseEtherType is ethernet type
	EtherType() PatternFlowPfcPauseEtherType
	// SetEtherType assigns PatternFlowPfcPauseEtherType provided by user to FlowPfcPause.
	// PatternFlowPfcPauseEtherType is ethernet type
	SetEtherType(value PatternFlowPfcPauseEtherType) FlowPfcPause
	// HasEtherType checks if EtherType has been set in FlowPfcPause
	HasEtherType() bool
	// ControlOpCode returns PatternFlowPfcPauseControlOpCode, set in FlowPfcPause.
	// PatternFlowPfcPauseControlOpCode is control operation code
	ControlOpCode() PatternFlowPfcPauseControlOpCode
	// SetControlOpCode assigns PatternFlowPfcPauseControlOpCode provided by user to FlowPfcPause.
	// PatternFlowPfcPauseControlOpCode is control operation code
	SetControlOpCode(value PatternFlowPfcPauseControlOpCode) FlowPfcPause
	// HasControlOpCode checks if ControlOpCode has been set in FlowPfcPause
	HasControlOpCode() bool
	// ClassEnableVector returns PatternFlowPfcPauseClassEnableVector, set in FlowPfcPause.
	// PatternFlowPfcPauseClassEnableVector is destination
	ClassEnableVector() PatternFlowPfcPauseClassEnableVector
	// SetClassEnableVector assigns PatternFlowPfcPauseClassEnableVector provided by user to FlowPfcPause.
	// PatternFlowPfcPauseClassEnableVector is destination
	SetClassEnableVector(value PatternFlowPfcPauseClassEnableVector) FlowPfcPause
	// HasClassEnableVector checks if ClassEnableVector has been set in FlowPfcPause
	HasClassEnableVector() bool
	// PauseClass0 returns PatternFlowPfcPausePauseClass0, set in FlowPfcPause.
	// PatternFlowPfcPausePauseClass0 is pause class 0
	PauseClass0() PatternFlowPfcPausePauseClass0
	// SetPauseClass0 assigns PatternFlowPfcPausePauseClass0 provided by user to FlowPfcPause.
	// PatternFlowPfcPausePauseClass0 is pause class 0
	SetPauseClass0(value PatternFlowPfcPausePauseClass0) FlowPfcPause
	// HasPauseClass0 checks if PauseClass0 has been set in FlowPfcPause
	HasPauseClass0() bool
	// PauseClass1 returns PatternFlowPfcPausePauseClass1, set in FlowPfcPause.
	// PatternFlowPfcPausePauseClass1 is pause class 1
	PauseClass1() PatternFlowPfcPausePauseClass1
	// SetPauseClass1 assigns PatternFlowPfcPausePauseClass1 provided by user to FlowPfcPause.
	// PatternFlowPfcPausePauseClass1 is pause class 1
	SetPauseClass1(value PatternFlowPfcPausePauseClass1) FlowPfcPause
	// HasPauseClass1 checks if PauseClass1 has been set in FlowPfcPause
	HasPauseClass1() bool
	// PauseClass2 returns PatternFlowPfcPausePauseClass2, set in FlowPfcPause.
	// PatternFlowPfcPausePauseClass2 is pause class 2
	PauseClass2() PatternFlowPfcPausePauseClass2
	// SetPauseClass2 assigns PatternFlowPfcPausePauseClass2 provided by user to FlowPfcPause.
	// PatternFlowPfcPausePauseClass2 is pause class 2
	SetPauseClass2(value PatternFlowPfcPausePauseClass2) FlowPfcPause
	// HasPauseClass2 checks if PauseClass2 has been set in FlowPfcPause
	HasPauseClass2() bool
	// PauseClass3 returns PatternFlowPfcPausePauseClass3, set in FlowPfcPause.
	// PatternFlowPfcPausePauseClass3 is pause class 3
	PauseClass3() PatternFlowPfcPausePauseClass3
	// SetPauseClass3 assigns PatternFlowPfcPausePauseClass3 provided by user to FlowPfcPause.
	// PatternFlowPfcPausePauseClass3 is pause class 3
	SetPauseClass3(value PatternFlowPfcPausePauseClass3) FlowPfcPause
	// HasPauseClass3 checks if PauseClass3 has been set in FlowPfcPause
	HasPauseClass3() bool
	// PauseClass4 returns PatternFlowPfcPausePauseClass4, set in FlowPfcPause.
	// PatternFlowPfcPausePauseClass4 is pause class 4
	PauseClass4() PatternFlowPfcPausePauseClass4
	// SetPauseClass4 assigns PatternFlowPfcPausePauseClass4 provided by user to FlowPfcPause.
	// PatternFlowPfcPausePauseClass4 is pause class 4
	SetPauseClass4(value PatternFlowPfcPausePauseClass4) FlowPfcPause
	// HasPauseClass4 checks if PauseClass4 has been set in FlowPfcPause
	HasPauseClass4() bool
	// PauseClass5 returns PatternFlowPfcPausePauseClass5, set in FlowPfcPause.
	// PatternFlowPfcPausePauseClass5 is pause class 5
	PauseClass5() PatternFlowPfcPausePauseClass5
	// SetPauseClass5 assigns PatternFlowPfcPausePauseClass5 provided by user to FlowPfcPause.
	// PatternFlowPfcPausePauseClass5 is pause class 5
	SetPauseClass5(value PatternFlowPfcPausePauseClass5) FlowPfcPause
	// HasPauseClass5 checks if PauseClass5 has been set in FlowPfcPause
	HasPauseClass5() bool
	// PauseClass6 returns PatternFlowPfcPausePauseClass6, set in FlowPfcPause.
	// PatternFlowPfcPausePauseClass6 is pause class 6
	PauseClass6() PatternFlowPfcPausePauseClass6
	// SetPauseClass6 assigns PatternFlowPfcPausePauseClass6 provided by user to FlowPfcPause.
	// PatternFlowPfcPausePauseClass6 is pause class 6
	SetPauseClass6(value PatternFlowPfcPausePauseClass6) FlowPfcPause
	// HasPauseClass6 checks if PauseClass6 has been set in FlowPfcPause
	HasPauseClass6() bool
	// PauseClass7 returns PatternFlowPfcPausePauseClass7, set in FlowPfcPause.
	// PatternFlowPfcPausePauseClass7 is pause class 7
	PauseClass7() PatternFlowPfcPausePauseClass7
	// SetPauseClass7 assigns PatternFlowPfcPausePauseClass7 provided by user to FlowPfcPause.
	// PatternFlowPfcPausePauseClass7 is pause class 7
	SetPauseClass7(value PatternFlowPfcPausePauseClass7) FlowPfcPause
	// HasPauseClass7 checks if PauseClass7 has been set in FlowPfcPause
	HasPauseClass7() bool
	setNil()
}

// description is TBD
// Dst returns a PatternFlowPfcPauseDst
func (obj *flowPfcPause) Dst() PatternFlowPfcPauseDst {
	if obj.obj.Dst == nil {
		obj.obj.Dst = NewPatternFlowPfcPauseDst().msg()
	}
	if obj.dstHolder == nil {
		obj.dstHolder = &patternFlowPfcPauseDst{obj: obj.obj.Dst}
	}
	return obj.dstHolder
}

// description is TBD
// Dst returns a PatternFlowPfcPauseDst
func (obj *flowPfcPause) HasDst() bool {
	return obj.obj.Dst != nil
}

// description is TBD
// SetDst sets the PatternFlowPfcPauseDst value in the FlowPfcPause object
func (obj *flowPfcPause) SetDst(value PatternFlowPfcPauseDst) FlowPfcPause {

	obj.dstHolder = nil
	obj.obj.Dst = value.msg()

	return obj
}

// description is TBD
// Src returns a PatternFlowPfcPauseSrc
func (obj *flowPfcPause) Src() PatternFlowPfcPauseSrc {
	if obj.obj.Src == nil {
		obj.obj.Src = NewPatternFlowPfcPauseSrc().msg()
	}
	if obj.srcHolder == nil {
		obj.srcHolder = &patternFlowPfcPauseSrc{obj: obj.obj.Src}
	}
	return obj.srcHolder
}

// description is TBD
// Src returns a PatternFlowPfcPauseSrc
func (obj *flowPfcPause) HasSrc() bool {
	return obj.obj.Src != nil
}

// description is TBD
// SetSrc sets the PatternFlowPfcPauseSrc value in the FlowPfcPause object
func (obj *flowPfcPause) SetSrc(value PatternFlowPfcPauseSrc) FlowPfcPause {

	obj.srcHolder = nil
	obj.obj.Src = value.msg()

	return obj
}

// description is TBD
// EtherType returns a PatternFlowPfcPauseEtherType
func (obj *flowPfcPause) EtherType() PatternFlowPfcPauseEtherType {
	if obj.obj.EtherType == nil {
		obj.obj.EtherType = NewPatternFlowPfcPauseEtherType().msg()
	}
	if obj.etherTypeHolder == nil {
		obj.etherTypeHolder = &patternFlowPfcPauseEtherType{obj: obj.obj.EtherType}
	}
	return obj.etherTypeHolder
}

// description is TBD
// EtherType returns a PatternFlowPfcPauseEtherType
func (obj *flowPfcPause) HasEtherType() bool {
	return obj.obj.EtherType != nil
}

// description is TBD
// SetEtherType sets the PatternFlowPfcPauseEtherType value in the FlowPfcPause object
func (obj *flowPfcPause) SetEtherType(value PatternFlowPfcPauseEtherType) FlowPfcPause {

	obj.etherTypeHolder = nil
	obj.obj.EtherType = value.msg()

	return obj
}

// description is TBD
// ControlOpCode returns a PatternFlowPfcPauseControlOpCode
func (obj *flowPfcPause) ControlOpCode() PatternFlowPfcPauseControlOpCode {
	if obj.obj.ControlOpCode == nil {
		obj.obj.ControlOpCode = NewPatternFlowPfcPauseControlOpCode().msg()
	}
	if obj.controlOpCodeHolder == nil {
		obj.controlOpCodeHolder = &patternFlowPfcPauseControlOpCode{obj: obj.obj.ControlOpCode}
	}
	return obj.controlOpCodeHolder
}

// description is TBD
// ControlOpCode returns a PatternFlowPfcPauseControlOpCode
func (obj *flowPfcPause) HasControlOpCode() bool {
	return obj.obj.ControlOpCode != nil
}

// description is TBD
// SetControlOpCode sets the PatternFlowPfcPauseControlOpCode value in the FlowPfcPause object
func (obj *flowPfcPause) SetControlOpCode(value PatternFlowPfcPauseControlOpCode) FlowPfcPause {

	obj.controlOpCodeHolder = nil
	obj.obj.ControlOpCode = value.msg()

	return obj
}

// description is TBD
// ClassEnableVector returns a PatternFlowPfcPauseClassEnableVector
func (obj *flowPfcPause) ClassEnableVector() PatternFlowPfcPauseClassEnableVector {
	if obj.obj.ClassEnableVector == nil {
		obj.obj.ClassEnableVector = NewPatternFlowPfcPauseClassEnableVector().msg()
	}
	if obj.classEnableVectorHolder == nil {
		obj.classEnableVectorHolder = &patternFlowPfcPauseClassEnableVector{obj: obj.obj.ClassEnableVector}
	}
	return obj.classEnableVectorHolder
}

// description is TBD
// ClassEnableVector returns a PatternFlowPfcPauseClassEnableVector
func (obj *flowPfcPause) HasClassEnableVector() bool {
	return obj.obj.ClassEnableVector != nil
}

// description is TBD
// SetClassEnableVector sets the PatternFlowPfcPauseClassEnableVector value in the FlowPfcPause object
func (obj *flowPfcPause) SetClassEnableVector(value PatternFlowPfcPauseClassEnableVector) FlowPfcPause {

	obj.classEnableVectorHolder = nil
	obj.obj.ClassEnableVector = value.msg()

	return obj
}

// description is TBD
// PauseClass0 returns a PatternFlowPfcPausePauseClass0
func (obj *flowPfcPause) PauseClass0() PatternFlowPfcPausePauseClass0 {
	if obj.obj.PauseClass_0 == nil {
		obj.obj.PauseClass_0 = NewPatternFlowPfcPausePauseClass0().msg()
	}
	if obj.pauseClass_0Holder == nil {
		obj.pauseClass_0Holder = &patternFlowPfcPausePauseClass0{obj: obj.obj.PauseClass_0}
	}
	return obj.pauseClass_0Holder
}

// description is TBD
// PauseClass0 returns a PatternFlowPfcPausePauseClass0
func (obj *flowPfcPause) HasPauseClass0() bool {
	return obj.obj.PauseClass_0 != nil
}

// description is TBD
// SetPauseClass0 sets the PatternFlowPfcPausePauseClass0 value in the FlowPfcPause object
func (obj *flowPfcPause) SetPauseClass0(value PatternFlowPfcPausePauseClass0) FlowPfcPause {

	obj.pauseClass_0Holder = nil
	obj.obj.PauseClass_0 = value.msg()

	return obj
}

// description is TBD
// PauseClass1 returns a PatternFlowPfcPausePauseClass1
func (obj *flowPfcPause) PauseClass1() PatternFlowPfcPausePauseClass1 {
	if obj.obj.PauseClass_1 == nil {
		obj.obj.PauseClass_1 = NewPatternFlowPfcPausePauseClass1().msg()
	}
	if obj.pauseClass_1Holder == nil {
		obj.pauseClass_1Holder = &patternFlowPfcPausePauseClass1{obj: obj.obj.PauseClass_1}
	}
	return obj.pauseClass_1Holder
}

// description is TBD
// PauseClass1 returns a PatternFlowPfcPausePauseClass1
func (obj *flowPfcPause) HasPauseClass1() bool {
	return obj.obj.PauseClass_1 != nil
}

// description is TBD
// SetPauseClass1 sets the PatternFlowPfcPausePauseClass1 value in the FlowPfcPause object
func (obj *flowPfcPause) SetPauseClass1(value PatternFlowPfcPausePauseClass1) FlowPfcPause {

	obj.pauseClass_1Holder = nil
	obj.obj.PauseClass_1 = value.msg()

	return obj
}

// description is TBD
// PauseClass2 returns a PatternFlowPfcPausePauseClass2
func (obj *flowPfcPause) PauseClass2() PatternFlowPfcPausePauseClass2 {
	if obj.obj.PauseClass_2 == nil {
		obj.obj.PauseClass_2 = NewPatternFlowPfcPausePauseClass2().msg()
	}
	if obj.pauseClass_2Holder == nil {
		obj.pauseClass_2Holder = &patternFlowPfcPausePauseClass2{obj: obj.obj.PauseClass_2}
	}
	return obj.pauseClass_2Holder
}

// description is TBD
// PauseClass2 returns a PatternFlowPfcPausePauseClass2
func (obj *flowPfcPause) HasPauseClass2() bool {
	return obj.obj.PauseClass_2 != nil
}

// description is TBD
// SetPauseClass2 sets the PatternFlowPfcPausePauseClass2 value in the FlowPfcPause object
func (obj *flowPfcPause) SetPauseClass2(value PatternFlowPfcPausePauseClass2) FlowPfcPause {

	obj.pauseClass_2Holder = nil
	obj.obj.PauseClass_2 = value.msg()

	return obj
}

// description is TBD
// PauseClass3 returns a PatternFlowPfcPausePauseClass3
func (obj *flowPfcPause) PauseClass3() PatternFlowPfcPausePauseClass3 {
	if obj.obj.PauseClass_3 == nil {
		obj.obj.PauseClass_3 = NewPatternFlowPfcPausePauseClass3().msg()
	}
	if obj.pauseClass_3Holder == nil {
		obj.pauseClass_3Holder = &patternFlowPfcPausePauseClass3{obj: obj.obj.PauseClass_3}
	}
	return obj.pauseClass_3Holder
}

// description is TBD
// PauseClass3 returns a PatternFlowPfcPausePauseClass3
func (obj *flowPfcPause) HasPauseClass3() bool {
	return obj.obj.PauseClass_3 != nil
}

// description is TBD
// SetPauseClass3 sets the PatternFlowPfcPausePauseClass3 value in the FlowPfcPause object
func (obj *flowPfcPause) SetPauseClass3(value PatternFlowPfcPausePauseClass3) FlowPfcPause {

	obj.pauseClass_3Holder = nil
	obj.obj.PauseClass_3 = value.msg()

	return obj
}

// description is TBD
// PauseClass4 returns a PatternFlowPfcPausePauseClass4
func (obj *flowPfcPause) PauseClass4() PatternFlowPfcPausePauseClass4 {
	if obj.obj.PauseClass_4 == nil {
		obj.obj.PauseClass_4 = NewPatternFlowPfcPausePauseClass4().msg()
	}
	if obj.pauseClass_4Holder == nil {
		obj.pauseClass_4Holder = &patternFlowPfcPausePauseClass4{obj: obj.obj.PauseClass_4}
	}
	return obj.pauseClass_4Holder
}

// description is TBD
// PauseClass4 returns a PatternFlowPfcPausePauseClass4
func (obj *flowPfcPause) HasPauseClass4() bool {
	return obj.obj.PauseClass_4 != nil
}

// description is TBD
// SetPauseClass4 sets the PatternFlowPfcPausePauseClass4 value in the FlowPfcPause object
func (obj *flowPfcPause) SetPauseClass4(value PatternFlowPfcPausePauseClass4) FlowPfcPause {

	obj.pauseClass_4Holder = nil
	obj.obj.PauseClass_4 = value.msg()

	return obj
}

// description is TBD
// PauseClass5 returns a PatternFlowPfcPausePauseClass5
func (obj *flowPfcPause) PauseClass5() PatternFlowPfcPausePauseClass5 {
	if obj.obj.PauseClass_5 == nil {
		obj.obj.PauseClass_5 = NewPatternFlowPfcPausePauseClass5().msg()
	}
	if obj.pauseClass_5Holder == nil {
		obj.pauseClass_5Holder = &patternFlowPfcPausePauseClass5{obj: obj.obj.PauseClass_5}
	}
	return obj.pauseClass_5Holder
}

// description is TBD
// PauseClass5 returns a PatternFlowPfcPausePauseClass5
func (obj *flowPfcPause) HasPauseClass5() bool {
	return obj.obj.PauseClass_5 != nil
}

// description is TBD
// SetPauseClass5 sets the PatternFlowPfcPausePauseClass5 value in the FlowPfcPause object
func (obj *flowPfcPause) SetPauseClass5(value PatternFlowPfcPausePauseClass5) FlowPfcPause {

	obj.pauseClass_5Holder = nil
	obj.obj.PauseClass_5 = value.msg()

	return obj
}

// description is TBD
// PauseClass6 returns a PatternFlowPfcPausePauseClass6
func (obj *flowPfcPause) PauseClass6() PatternFlowPfcPausePauseClass6 {
	if obj.obj.PauseClass_6 == nil {
		obj.obj.PauseClass_6 = NewPatternFlowPfcPausePauseClass6().msg()
	}
	if obj.pauseClass_6Holder == nil {
		obj.pauseClass_6Holder = &patternFlowPfcPausePauseClass6{obj: obj.obj.PauseClass_6}
	}
	return obj.pauseClass_6Holder
}

// description is TBD
// PauseClass6 returns a PatternFlowPfcPausePauseClass6
func (obj *flowPfcPause) HasPauseClass6() bool {
	return obj.obj.PauseClass_6 != nil
}

// description is TBD
// SetPauseClass6 sets the PatternFlowPfcPausePauseClass6 value in the FlowPfcPause object
func (obj *flowPfcPause) SetPauseClass6(value PatternFlowPfcPausePauseClass6) FlowPfcPause {

	obj.pauseClass_6Holder = nil
	obj.obj.PauseClass_6 = value.msg()

	return obj
}

// description is TBD
// PauseClass7 returns a PatternFlowPfcPausePauseClass7
func (obj *flowPfcPause) PauseClass7() PatternFlowPfcPausePauseClass7 {
	if obj.obj.PauseClass_7 == nil {
		obj.obj.PauseClass_7 = NewPatternFlowPfcPausePauseClass7().msg()
	}
	if obj.pauseClass_7Holder == nil {
		obj.pauseClass_7Holder = &patternFlowPfcPausePauseClass7{obj: obj.obj.PauseClass_7}
	}
	return obj.pauseClass_7Holder
}

// description is TBD
// PauseClass7 returns a PatternFlowPfcPausePauseClass7
func (obj *flowPfcPause) HasPauseClass7() bool {
	return obj.obj.PauseClass_7 != nil
}

// description is TBD
// SetPauseClass7 sets the PatternFlowPfcPausePauseClass7 value in the FlowPfcPause object
func (obj *flowPfcPause) SetPauseClass7(value PatternFlowPfcPausePauseClass7) FlowPfcPause {

	obj.pauseClass_7Holder = nil
	obj.obj.PauseClass_7 = value.msg()

	return obj
}

func (obj *flowPfcPause) validateObj(vObj *validation, set_default bool) {
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

	if obj.obj.ClassEnableVector != nil {

		obj.ClassEnableVector().validateObj(vObj, set_default)
	}

	if obj.obj.PauseClass_0 != nil {

		obj.PauseClass0().validateObj(vObj, set_default)
	}

	if obj.obj.PauseClass_1 != nil {

		obj.PauseClass1().validateObj(vObj, set_default)
	}

	if obj.obj.PauseClass_2 != nil {

		obj.PauseClass2().validateObj(vObj, set_default)
	}

	if obj.obj.PauseClass_3 != nil {

		obj.PauseClass3().validateObj(vObj, set_default)
	}

	if obj.obj.PauseClass_4 != nil {

		obj.PauseClass4().validateObj(vObj, set_default)
	}

	if obj.obj.PauseClass_5 != nil {

		obj.PauseClass5().validateObj(vObj, set_default)
	}

	if obj.obj.PauseClass_6 != nil {

		obj.PauseClass6().validateObj(vObj, set_default)
	}

	if obj.obj.PauseClass_7 != nil {

		obj.PauseClass7().validateObj(vObj, set_default)
	}

}

func (obj *flowPfcPause) setDefault() {

}
