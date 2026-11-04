package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6SRHPadTlv *****
type flowIpv6SRHPadTlv struct {
	validation
	obj          *otg.FlowIpv6SRHPadTlv
	marshaller   marshalFlowIpv6SRHPadTlv
	unMarshaller unMarshalFlowIpv6SRHPadTlv
	typeHolder   PatternFlowIpv6SRHPadTlvType
	lengthHolder PatternFlowIpv6SRHPadTlvLength
}

func NewFlowIpv6SRHPadTlv() FlowIpv6SRHPadTlv {
	obj := flowIpv6SRHPadTlv{obj: &otg.FlowIpv6SRHPadTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6SRHPadTlv) msg() *otg.FlowIpv6SRHPadTlv {
	return obj.obj
}

func (obj *flowIpv6SRHPadTlv) setMsg(msg *otg.FlowIpv6SRHPadTlv) FlowIpv6SRHPadTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6SRHPadTlv struct {
	obj *flowIpv6SRHPadTlv
}

type marshalFlowIpv6SRHPadTlv interface {
	// ToProto marshals FlowIpv6SRHPadTlv to protobuf object *otg.FlowIpv6SRHPadTlv
	ToProto() (*otg.FlowIpv6SRHPadTlv, error)
	// ToPbText marshals FlowIpv6SRHPadTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6SRHPadTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6SRHPadTlv to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6SRHPadTlv struct {
	obj *flowIpv6SRHPadTlv
}

type unMarshalFlowIpv6SRHPadTlv interface {
	// FromProto unmarshals FlowIpv6SRHPadTlv from protobuf object *otg.FlowIpv6SRHPadTlv
	FromProto(msg *otg.FlowIpv6SRHPadTlv) (FlowIpv6SRHPadTlv, error)
	// FromPbText unmarshals FlowIpv6SRHPadTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6SRHPadTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6SRHPadTlv from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6SRHPadTlv) Marshal() marshalFlowIpv6SRHPadTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6SRHPadTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6SRHPadTlv) Unmarshal() unMarshalFlowIpv6SRHPadTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6SRHPadTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6SRHPadTlv) ToProto() (*otg.FlowIpv6SRHPadTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6SRHPadTlv) FromProto(msg *otg.FlowIpv6SRHPadTlv) (FlowIpv6SRHPadTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6SRHPadTlv) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6SRHPadTlv) FromPbText(value string) error {
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

func (m *marshalflowIpv6SRHPadTlv) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6SRHPadTlv) FromYaml(value string) error {
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

func (m *marshalflowIpv6SRHPadTlv) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6SRHPadTlv) FromJson(value string) error {
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

func (obj *flowIpv6SRHPadTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6SRHPadTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6SRHPadTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6SRHPadTlv) Clone() (FlowIpv6SRHPadTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6SRHPadTlv()
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

func (obj *flowIpv6SRHPadTlv) setNil() {
	obj.typeHolder = nil
	obj.lengthHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv6SRHPadTlv is sRH Padding TLV (type 4, RFC 8754 Section 2.1). Used to align the SRH TLV block to an 8-byte boundary. The padding bytes are set to zero and are skipped by transit nodes. When present this TLV is included in the SRH TLV section; omit the object to suppress it. Reference: RFC 8754 Section 2.1.
type FlowIpv6SRHPadTlv interface {
	Validation
	// msg marshals FlowIpv6SRHPadTlv to protobuf object *otg.FlowIpv6SRHPadTlv
	// and doesn't set defaults
	msg() *otg.FlowIpv6SRHPadTlv
	// setMsg unmarshals FlowIpv6SRHPadTlv from protobuf object *otg.FlowIpv6SRHPadTlv
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6SRHPadTlv) FlowIpv6SRHPadTlv
	// provides marshal interface
	Marshal() marshalFlowIpv6SRHPadTlv
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6SRHPadTlv
	// validate validates FlowIpv6SRHPadTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6SRHPadTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns PatternFlowIpv6SRHPadTlvType, set in FlowIpv6SRHPadTlv.
	// PatternFlowIpv6SRHPadTlvType is tLV type code (RFC 8754 Section 2.1). The canonical value is 4. Set to a different value only for negative or conformance testing.
	Type() PatternFlowIpv6SRHPadTlvType
	// SetType assigns PatternFlowIpv6SRHPadTlvType provided by user to FlowIpv6SRHPadTlv.
	// PatternFlowIpv6SRHPadTlvType is tLV type code (RFC 8754 Section 2.1). The canonical value is 4. Set to a different value only for negative or conformance testing.
	SetType(value PatternFlowIpv6SRHPadTlvType) FlowIpv6SRHPadTlv
	// HasType checks if Type has been set in FlowIpv6SRHPadTlv
	HasType() bool
	// Length returns PatternFlowIpv6SRHPadTlvLength, set in FlowIpv6SRHPadTlv.
	// PatternFlowIpv6SRHPadTlvLength is number of padding bytes following the Type and Length fields (RFC 8754 Section 2.1). When auto is assigned the implementation computes the value needed to align the TLV section to an 8-byte boundary.
	Length() PatternFlowIpv6SRHPadTlvLength
	// SetLength assigns PatternFlowIpv6SRHPadTlvLength provided by user to FlowIpv6SRHPadTlv.
	// PatternFlowIpv6SRHPadTlvLength is number of padding bytes following the Type and Length fields (RFC 8754 Section 2.1). When auto is assigned the implementation computes the value needed to align the TLV section to an 8-byte boundary.
	SetLength(value PatternFlowIpv6SRHPadTlvLength) FlowIpv6SRHPadTlv
	// HasLength checks if Length has been set in FlowIpv6SRHPadTlv
	HasLength() bool
	setNil()
}

// description is TBD
// Type returns a PatternFlowIpv6SRHPadTlvType
func (obj *flowIpv6SRHPadTlv) Type() PatternFlowIpv6SRHPadTlvType {
	if obj.obj.Type == nil {
		obj.obj.Type = NewPatternFlowIpv6SRHPadTlvType().msg()
	}
	if obj.typeHolder == nil {
		obj.typeHolder = &patternFlowIpv6SRHPadTlvType{obj: obj.obj.Type}
	}
	return obj.typeHolder
}

// description is TBD
// Type returns a PatternFlowIpv6SRHPadTlvType
func (obj *flowIpv6SRHPadTlv) HasType() bool {
	return obj.obj.Type != nil
}

// description is TBD
// SetType sets the PatternFlowIpv6SRHPadTlvType value in the FlowIpv6SRHPadTlv object
func (obj *flowIpv6SRHPadTlv) SetType(value PatternFlowIpv6SRHPadTlvType) FlowIpv6SRHPadTlv {

	obj.typeHolder = nil
	obj.obj.Type = value.msg()

	return obj
}

// description is TBD
// Length returns a PatternFlowIpv6SRHPadTlvLength
func (obj *flowIpv6SRHPadTlv) Length() PatternFlowIpv6SRHPadTlvLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewPatternFlowIpv6SRHPadTlvLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &patternFlowIpv6SRHPadTlvLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// description is TBD
// Length returns a PatternFlowIpv6SRHPadTlvLength
func (obj *flowIpv6SRHPadTlv) HasLength() bool {
	return obj.obj.Length != nil
}

// description is TBD
// SetLength sets the PatternFlowIpv6SRHPadTlvLength value in the FlowIpv6SRHPadTlv object
func (obj *flowIpv6SRHPadTlv) SetLength(value PatternFlowIpv6SRHPadTlvLength) FlowIpv6SRHPadTlv {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

func (obj *flowIpv6SRHPadTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Type != nil {

		obj.Type().validateObj(vObj, set_default)
	}

	if obj.obj.Length != nil {

		obj.Length().validateObj(vObj, set_default)
	}

}

func (obj *flowIpv6SRHPadTlv) setDefault() {

}
