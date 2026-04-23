package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfm *****
type flowCfm struct {
	validation
	obj           *otg.FlowCfm
	marshaller    marshalFlowCfm
	unMarshaller  unMarshalFlowCfm
	mdLevelHolder PatternFlowCfmMdLevel
	versionHolder PatternFlowCfmVersion
	opCodeHolder  FlowCfmSelectOpCode
}

func NewFlowCfm() FlowCfm {
	obj := flowCfm{obj: &otg.FlowCfm{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfm) msg() *otg.FlowCfm {
	return obj.obj
}

func (obj *flowCfm) setMsg(msg *otg.FlowCfm) FlowCfm {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfm struct {
	obj *flowCfm
}

type marshalFlowCfm interface {
	// ToProto marshals FlowCfm to protobuf object *otg.FlowCfm
	ToProto() (*otg.FlowCfm, error)
	// ToPbText marshals FlowCfm to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfm to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfm to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfm struct {
	obj *flowCfm
}

type unMarshalFlowCfm interface {
	// FromProto unmarshals FlowCfm from protobuf object *otg.FlowCfm
	FromProto(msg *otg.FlowCfm) (FlowCfm, error)
	// FromPbText unmarshals FlowCfm from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfm from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfm from JSON text
	FromJson(value string) error
}

func (obj *flowCfm) Marshal() marshalFlowCfm {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfm{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfm) Unmarshal() unMarshalFlowCfm {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfm{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfm) ToProto() (*otg.FlowCfm, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfm) FromProto(msg *otg.FlowCfm) (FlowCfm, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfm) ToPbText() (string, error) {
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

func (m *unMarshalflowCfm) FromPbText(value string) error {
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

func (m *marshalflowCfm) ToYaml() (string, error) {
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

func (m *unMarshalflowCfm) FromYaml(value string) error {
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

func (m *marshalflowCfm) ToJson() (string, error) {
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

func (m *unMarshalflowCfm) FromJson(value string) error {
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

func (obj *flowCfm) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfm) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfm) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfm) Clone() (FlowCfm, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfm()
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

func (obj *flowCfm) setNil() {
	obj.mdLevelHolder = nil
	obj.versionHolder = nil
	obj.opCodeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowCfm is defines the fields of a Connectivity Fault Management (CFM) Packet Data Unit (PDU) as specified by IEEE 802.1.
type FlowCfm interface {
	Validation
	// msg marshals FlowCfm to protobuf object *otg.FlowCfm
	// and doesn't set defaults
	msg() *otg.FlowCfm
	// setMsg unmarshals FlowCfm from protobuf object *otg.FlowCfm
	// and doesn't set defaults
	setMsg(*otg.FlowCfm) FlowCfm
	// provides marshal interface
	Marshal() marshalFlowCfm
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfm
	// validate validates FlowCfm
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfm, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// MdLevel returns PatternFlowCfmMdLevel, set in FlowCfm.
	// PatternFlowCfmMdLevel is maintenance Domain Level
	MdLevel() PatternFlowCfmMdLevel
	// SetMdLevel assigns PatternFlowCfmMdLevel provided by user to FlowCfm.
	// PatternFlowCfmMdLevel is maintenance Domain Level
	SetMdLevel(value PatternFlowCfmMdLevel) FlowCfm
	// HasMdLevel checks if MdLevel has been set in FlowCfm
	HasMdLevel() bool
	// Version returns PatternFlowCfmVersion, set in FlowCfm.
	// PatternFlowCfmVersion is indicates the version as specified in 8021Q
	Version() PatternFlowCfmVersion
	// SetVersion assigns PatternFlowCfmVersion provided by user to FlowCfm.
	// PatternFlowCfmVersion is indicates the version as specified in 8021Q
	SetVersion(value PatternFlowCfmVersion) FlowCfm
	// HasVersion checks if Version has been set in FlowCfm
	HasVersion() bool
	// OpCode returns FlowCfmSelectOpCode, set in FlowCfm.
	// FlowCfmSelectOpCode is select message header. Currently only Continuity Check Message (CCM) and Loobback Reply (LBR) are supported.
	OpCode() FlowCfmSelectOpCode
	// SetOpCode assigns FlowCfmSelectOpCode provided by user to FlowCfm.
	// FlowCfmSelectOpCode is select message header. Currently only Continuity Check Message (CCM) and Loobback Reply (LBR) are supported.
	SetOpCode(value FlowCfmSelectOpCode) FlowCfm
	// HasOpCode checks if OpCode has been set in FlowCfm
	HasOpCode() bool
	setNil()
}

// description is TBD
// MdLevel returns a PatternFlowCfmMdLevel
func (obj *flowCfm) MdLevel() PatternFlowCfmMdLevel {
	if obj.obj.MdLevel == nil {
		obj.obj.MdLevel = NewPatternFlowCfmMdLevel().msg()
	}
	if obj.mdLevelHolder == nil {
		obj.mdLevelHolder = &patternFlowCfmMdLevel{obj: obj.obj.MdLevel}
	}
	return obj.mdLevelHolder
}

// description is TBD
// MdLevel returns a PatternFlowCfmMdLevel
func (obj *flowCfm) HasMdLevel() bool {
	return obj.obj.MdLevel != nil
}

// description is TBD
// SetMdLevel sets the PatternFlowCfmMdLevel value in the FlowCfm object
func (obj *flowCfm) SetMdLevel(value PatternFlowCfmMdLevel) FlowCfm {

	obj.mdLevelHolder = nil
	obj.obj.MdLevel = value.msg()

	return obj
}

// description is TBD
// Version returns a PatternFlowCfmVersion
func (obj *flowCfm) Version() PatternFlowCfmVersion {
	if obj.obj.Version == nil {
		obj.obj.Version = NewPatternFlowCfmVersion().msg()
	}
	if obj.versionHolder == nil {
		obj.versionHolder = &patternFlowCfmVersion{obj: obj.obj.Version}
	}
	return obj.versionHolder
}

// description is TBD
// Version returns a PatternFlowCfmVersion
func (obj *flowCfm) HasVersion() bool {
	return obj.obj.Version != nil
}

// description is TBD
// SetVersion sets the PatternFlowCfmVersion value in the FlowCfm object
func (obj *flowCfm) SetVersion(value PatternFlowCfmVersion) FlowCfm {

	obj.versionHolder = nil
	obj.obj.Version = value.msg()

	return obj
}

// description is TBD
// OpCode returns a FlowCfmSelectOpCode
func (obj *flowCfm) OpCode() FlowCfmSelectOpCode {
	if obj.obj.OpCode == nil {
		obj.obj.OpCode = NewFlowCfmSelectOpCode().msg()
	}
	if obj.opCodeHolder == nil {
		obj.opCodeHolder = &flowCfmSelectOpCode{obj: obj.obj.OpCode}
	}
	return obj.opCodeHolder
}

// description is TBD
// OpCode returns a FlowCfmSelectOpCode
func (obj *flowCfm) HasOpCode() bool {
	return obj.obj.OpCode != nil
}

// description is TBD
// SetOpCode sets the FlowCfmSelectOpCode value in the FlowCfm object
func (obj *flowCfm) SetOpCode(value FlowCfmSelectOpCode) FlowCfm {

	obj.opCodeHolder = nil
	obj.obj.OpCode = value.msg()

	return obj
}

func (obj *flowCfm) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MdLevel != nil {

		obj.MdLevel().validateObj(vObj, set_default)
	}

	if obj.obj.Version != nil {

		obj.Version().validateObj(vObj, set_default)
	}

	if obj.obj.OpCode != nil {

		obj.OpCode().validateObj(vObj, set_default)
	}

}

func (obj *flowCfm) setDefault() {

}
