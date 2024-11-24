package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowUdp *****
type flowUdp struct {
	validation
	obj            *otg.FlowUdp
	marshaller     marshalFlowUdp
	unMarshaller   unMarshalFlowUdp
	srcPortHolder  PatternFlowUdpSrcPort
	dstPortHolder  PatternFlowUdpDstPort
	lengthHolder   PatternFlowUdpLength
	checksumHolder PatternFlowUdpChecksum
}

func NewFlowUdp() FlowUdp {
	obj := flowUdp{obj: &otg.FlowUdp{}}
	obj.setDefault()
	return &obj
}

func (obj *flowUdp) msg() *otg.FlowUdp {
	return obj.obj
}

func (obj *flowUdp) setMsg(msg *otg.FlowUdp) FlowUdp {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowUdp struct {
	obj *flowUdp
}

type marshalFlowUdp interface {
	// ToProto marshals FlowUdp to protobuf object *otg.FlowUdp
	ToProto() (*otg.FlowUdp, error)
	// ToPbText marshals FlowUdp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowUdp to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowUdp to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowUdp to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowUdp struct {
	obj *flowUdp
}

type unMarshalFlowUdp interface {
	// FromProto unmarshals FlowUdp from protobuf object *otg.FlowUdp
	FromProto(msg *otg.FlowUdp) (FlowUdp, error)
	// FromPbText unmarshals FlowUdp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowUdp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowUdp from JSON text
	FromJson(value string) error
}

func (obj *flowUdp) Marshal() marshalFlowUdp {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowUdp{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowUdp) Unmarshal() unMarshalFlowUdp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowUdp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowUdp) ToProto() (*otg.FlowUdp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowUdp) FromProto(msg *otg.FlowUdp) (FlowUdp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowUdp) ToPbText() (string, error) {
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

func (m *unMarshalflowUdp) FromPbText(value string) error {
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

func (m *marshalflowUdp) ToYaml() (string, error) {
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

func (m *unMarshalflowUdp) FromYaml(value string) error {
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

func (m *marshalflowUdp) ToJsonRaw() (string, error) {
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

func (m *marshalflowUdp) ToJson() (string, error) {
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

func (m *unMarshalflowUdp) FromJson(value string) error {
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

func (obj *flowUdp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowUdp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowUdp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowUdp) Clone() (FlowUdp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowUdp()
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

func (obj *flowUdp) setNil() {
	obj.srcPortHolder = nil
	obj.dstPortHolder = nil
	obj.lengthHolder = nil
	obj.checksumHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowUdp is uDP packet header
type FlowUdp interface {
	Validation
	// msg marshals FlowUdp to protobuf object *otg.FlowUdp
	// and doesn't set defaults
	msg() *otg.FlowUdp
	// setMsg unmarshals FlowUdp from protobuf object *otg.FlowUdp
	// and doesn't set defaults
	setMsg(*otg.FlowUdp) FlowUdp
	// provides marshal interface
	Marshal() marshalFlowUdp
	// provides unmarshal interface
	Unmarshal() unMarshalFlowUdp
	// validate validates FlowUdp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowUdp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SrcPort returns PatternFlowUdpSrcPort, set in FlowUdp.
	// PatternFlowUdpSrcPort is source port
	SrcPort() PatternFlowUdpSrcPort
	// SetSrcPort assigns PatternFlowUdpSrcPort provided by user to FlowUdp.
	// PatternFlowUdpSrcPort is source port
	SetSrcPort(value PatternFlowUdpSrcPort) FlowUdp
	// HasSrcPort checks if SrcPort has been set in FlowUdp
	HasSrcPort() bool
	// DstPort returns PatternFlowUdpDstPort, set in FlowUdp.
	// PatternFlowUdpDstPort is destination port
	DstPort() PatternFlowUdpDstPort
	// SetDstPort assigns PatternFlowUdpDstPort provided by user to FlowUdp.
	// PatternFlowUdpDstPort is destination port
	SetDstPort(value PatternFlowUdpDstPort) FlowUdp
	// HasDstPort checks if DstPort has been set in FlowUdp
	HasDstPort() bool
	// Length returns PatternFlowUdpLength, set in FlowUdp.
	// PatternFlowUdpLength is length
	Length() PatternFlowUdpLength
	// SetLength assigns PatternFlowUdpLength provided by user to FlowUdp.
	// PatternFlowUdpLength is length
	SetLength(value PatternFlowUdpLength) FlowUdp
	// HasLength checks if Length has been set in FlowUdp
	HasLength() bool
	// Checksum returns PatternFlowUdpChecksum, set in FlowUdp.
	// PatternFlowUdpChecksum is uDP checksum
	Checksum() PatternFlowUdpChecksum
	// SetChecksum assigns PatternFlowUdpChecksum provided by user to FlowUdp.
	// PatternFlowUdpChecksum is uDP checksum
	SetChecksum(value PatternFlowUdpChecksum) FlowUdp
	// HasChecksum checks if Checksum has been set in FlowUdp
	HasChecksum() bool
	setNil()
}

// description is TBD
// SrcPort returns a PatternFlowUdpSrcPort
func (obj *flowUdp) SrcPort() PatternFlowUdpSrcPort {
	if obj.obj.SrcPort == nil {
		obj.obj.SrcPort = NewPatternFlowUdpSrcPort().msg()
	}
	if obj.srcPortHolder == nil {
		obj.srcPortHolder = &patternFlowUdpSrcPort{obj: obj.obj.SrcPort}
	}
	return obj.srcPortHolder
}

// description is TBD
// SrcPort returns a PatternFlowUdpSrcPort
func (obj *flowUdp) HasSrcPort() bool {
	return obj.obj.SrcPort != nil
}

// description is TBD
// SetSrcPort sets the PatternFlowUdpSrcPort value in the FlowUdp object
func (obj *flowUdp) SetSrcPort(value PatternFlowUdpSrcPort) FlowUdp {

	obj.srcPortHolder = nil
	obj.obj.SrcPort = value.msg()

	return obj
}

// description is TBD
// DstPort returns a PatternFlowUdpDstPort
func (obj *flowUdp) DstPort() PatternFlowUdpDstPort {
	if obj.obj.DstPort == nil {
		obj.obj.DstPort = NewPatternFlowUdpDstPort().msg()
	}
	if obj.dstPortHolder == nil {
		obj.dstPortHolder = &patternFlowUdpDstPort{obj: obj.obj.DstPort}
	}
	return obj.dstPortHolder
}

// description is TBD
// DstPort returns a PatternFlowUdpDstPort
func (obj *flowUdp) HasDstPort() bool {
	return obj.obj.DstPort != nil
}

// description is TBD
// SetDstPort sets the PatternFlowUdpDstPort value in the FlowUdp object
func (obj *flowUdp) SetDstPort(value PatternFlowUdpDstPort) FlowUdp {

	obj.dstPortHolder = nil
	obj.obj.DstPort = value.msg()

	return obj
}

// description is TBD
// Length returns a PatternFlowUdpLength
func (obj *flowUdp) Length() PatternFlowUdpLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewPatternFlowUdpLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &patternFlowUdpLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// description is TBD
// Length returns a PatternFlowUdpLength
func (obj *flowUdp) HasLength() bool {
	return obj.obj.Length != nil
}

// description is TBD
// SetLength sets the PatternFlowUdpLength value in the FlowUdp object
func (obj *flowUdp) SetLength(value PatternFlowUdpLength) FlowUdp {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// description is TBD
// Checksum returns a PatternFlowUdpChecksum
func (obj *flowUdp) Checksum() PatternFlowUdpChecksum {
	if obj.obj.Checksum == nil {
		obj.obj.Checksum = NewPatternFlowUdpChecksum().msg()
	}
	if obj.checksumHolder == nil {
		obj.checksumHolder = &patternFlowUdpChecksum{obj: obj.obj.Checksum}
	}
	return obj.checksumHolder
}

// description is TBD
// Checksum returns a PatternFlowUdpChecksum
func (obj *flowUdp) HasChecksum() bool {
	return obj.obj.Checksum != nil
}

// description is TBD
// SetChecksum sets the PatternFlowUdpChecksum value in the FlowUdp object
func (obj *flowUdp) SetChecksum(value PatternFlowUdpChecksum) FlowUdp {

	obj.checksumHolder = nil
	obj.obj.Checksum = value.msg()

	return obj
}

func (obj *flowUdp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SrcPort != nil {

		obj.SrcPort().validateObj(vObj, set_default)
	}

	if obj.obj.DstPort != nil {

		obj.DstPort().validateObj(vObj, set_default)
	}

	if obj.obj.Length != nil {

		obj.Length().validateObj(vObj, set_default)
	}

	if obj.obj.Checksum != nil {

		obj.Checksum().validateObj(vObj, set_default)
	}

}

func (obj *flowUdp) setDefault() {

}
