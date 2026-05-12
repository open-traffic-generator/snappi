package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6SegmentRoutinguSidFlags *****
type flowIpv6SegmentRoutinguSidFlags struct {
	validation
	obj          *otg.FlowIpv6SegmentRoutinguSidFlags
	marshaller   marshalFlowIpv6SegmentRoutinguSidFlags
	unMarshaller unMarshalFlowIpv6SegmentRoutinguSidFlags
	oamHolder    PatternFlowIpv6SegmentRoutinguSidFlagsOam
}

func NewFlowIpv6SegmentRoutinguSidFlags() FlowIpv6SegmentRoutinguSidFlags {
	obj := flowIpv6SegmentRoutinguSidFlags{obj: &otg.FlowIpv6SegmentRoutinguSidFlags{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6SegmentRoutinguSidFlags) msg() *otg.FlowIpv6SegmentRoutinguSidFlags {
	return obj.obj
}

func (obj *flowIpv6SegmentRoutinguSidFlags) setMsg(msg *otg.FlowIpv6SegmentRoutinguSidFlags) FlowIpv6SegmentRoutinguSidFlags {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6SegmentRoutinguSidFlags struct {
	obj *flowIpv6SegmentRoutinguSidFlags
}

type marshalFlowIpv6SegmentRoutinguSidFlags interface {
	// ToProto marshals FlowIpv6SegmentRoutinguSidFlags to protobuf object *otg.FlowIpv6SegmentRoutinguSidFlags
	ToProto() (*otg.FlowIpv6SegmentRoutinguSidFlags, error)
	// ToPbText marshals FlowIpv6SegmentRoutinguSidFlags to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6SegmentRoutinguSidFlags to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6SegmentRoutinguSidFlags to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6SegmentRoutinguSidFlags struct {
	obj *flowIpv6SegmentRoutinguSidFlags
}

type unMarshalFlowIpv6SegmentRoutinguSidFlags interface {
	// FromProto unmarshals FlowIpv6SegmentRoutinguSidFlags from protobuf object *otg.FlowIpv6SegmentRoutinguSidFlags
	FromProto(msg *otg.FlowIpv6SegmentRoutinguSidFlags) (FlowIpv6SegmentRoutinguSidFlags, error)
	// FromPbText unmarshals FlowIpv6SegmentRoutinguSidFlags from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6SegmentRoutinguSidFlags from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6SegmentRoutinguSidFlags from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6SegmentRoutinguSidFlags) Marshal() marshalFlowIpv6SegmentRoutinguSidFlags {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6SegmentRoutinguSidFlags{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6SegmentRoutinguSidFlags) Unmarshal() unMarshalFlowIpv6SegmentRoutinguSidFlags {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6SegmentRoutinguSidFlags{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6SegmentRoutinguSidFlags) ToProto() (*otg.FlowIpv6SegmentRoutinguSidFlags, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6SegmentRoutinguSidFlags) FromProto(msg *otg.FlowIpv6SegmentRoutinguSidFlags) (FlowIpv6SegmentRoutinguSidFlags, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6SegmentRoutinguSidFlags) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutinguSidFlags) FromPbText(value string) error {
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

func (m *marshalflowIpv6SegmentRoutinguSidFlags) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutinguSidFlags) FromYaml(value string) error {
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

func (m *marshalflowIpv6SegmentRoutinguSidFlags) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutinguSidFlags) FromJson(value string) error {
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

func (obj *flowIpv6SegmentRoutinguSidFlags) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutinguSidFlags) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutinguSidFlags) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6SegmentRoutinguSidFlags) Clone() (FlowIpv6SegmentRoutinguSidFlags, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6SegmentRoutinguSidFlags()
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

func (obj *flowIpv6SegmentRoutinguSidFlags) setNil() {
	obj.oamHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv6SegmentRoutinguSidFlags is sRH Flags field (RFC 8754 Section 2.1). An 8-bit field; RFC 8754 marks all bits as reserved.
type FlowIpv6SegmentRoutinguSidFlags interface {
	Validation
	// msg marshals FlowIpv6SegmentRoutinguSidFlags to protobuf object *otg.FlowIpv6SegmentRoutinguSidFlags
	// and doesn't set defaults
	msg() *otg.FlowIpv6SegmentRoutinguSidFlags
	// setMsg unmarshals FlowIpv6SegmentRoutinguSidFlags from protobuf object *otg.FlowIpv6SegmentRoutinguSidFlags
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6SegmentRoutinguSidFlags) FlowIpv6SegmentRoutinguSidFlags
	// provides marshal interface
	Marshal() marshalFlowIpv6SegmentRoutinguSidFlags
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6SegmentRoutinguSidFlags
	// validate validates FlowIpv6SegmentRoutinguSidFlags
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6SegmentRoutinguSidFlags, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Oam returns PatternFlowIpv6SegmentRoutinguSidFlagsOam, set in FlowIpv6SegmentRoutinguSidFlags.
	// PatternFlowIpv6SegmentRoutinguSidFlagsOam is (RFC 9259, section 2) OAM flag or O-flag is set at bit-2. Indicates if the packet is an Operations, Administration, and Maintenance (OAM) packet.
	Oam() PatternFlowIpv6SegmentRoutinguSidFlagsOam
	// SetOam assigns PatternFlowIpv6SegmentRoutinguSidFlagsOam provided by user to FlowIpv6SegmentRoutinguSidFlags.
	// PatternFlowIpv6SegmentRoutinguSidFlagsOam is (RFC 9259, section 2) OAM flag or O-flag is set at bit-2. Indicates if the packet is an Operations, Administration, and Maintenance (OAM) packet.
	SetOam(value PatternFlowIpv6SegmentRoutinguSidFlagsOam) FlowIpv6SegmentRoutinguSidFlags
	// HasOam checks if Oam has been set in FlowIpv6SegmentRoutinguSidFlags
	HasOam() bool
	setNil()
}

// description is TBD
// Oam returns a PatternFlowIpv6SegmentRoutinguSidFlagsOam
func (obj *flowIpv6SegmentRoutinguSidFlags) Oam() PatternFlowIpv6SegmentRoutinguSidFlagsOam {
	if obj.obj.Oam == nil {
		obj.obj.Oam = NewPatternFlowIpv6SegmentRoutinguSidFlagsOam().msg()
	}
	if obj.oamHolder == nil {
		obj.oamHolder = &patternFlowIpv6SegmentRoutinguSidFlagsOam{obj: obj.obj.Oam}
	}
	return obj.oamHolder
}

// description is TBD
// Oam returns a PatternFlowIpv6SegmentRoutinguSidFlagsOam
func (obj *flowIpv6SegmentRoutinguSidFlags) HasOam() bool {
	return obj.obj.Oam != nil
}

// description is TBD
// SetOam sets the PatternFlowIpv6SegmentRoutinguSidFlagsOam value in the FlowIpv6SegmentRoutinguSidFlags object
func (obj *flowIpv6SegmentRoutinguSidFlags) SetOam(value PatternFlowIpv6SegmentRoutinguSidFlagsOam) FlowIpv6SegmentRoutinguSidFlags {

	obj.oamHolder = nil
	obj.obj.Oam = value.msg()

	return obj
}

func (obj *flowIpv6SegmentRoutinguSidFlags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Oam != nil {

		obj.Oam().validateObj(vObj, set_default)
	}

}

func (obj *flowIpv6SegmentRoutinguSidFlags) setDefault() {

}
