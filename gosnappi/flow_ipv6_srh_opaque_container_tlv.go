package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6SRHOpaqueContainerTlv *****
type flowIpv6SRHOpaqueContainerTlv struct {
	validation
	obj          *otg.FlowIpv6SRHOpaqueContainerTlv
	marshaller   marshalFlowIpv6SRHOpaqueContainerTlv
	unMarshaller unMarshalFlowIpv6SRHOpaqueContainerTlv
	typeHolder   PatternFlowIpv6SRHOpaqueContainerTlvType
	lengthHolder PatternFlowIpv6SRHOpaqueContainerTlvLength
}

func NewFlowIpv6SRHOpaqueContainerTlv() FlowIpv6SRHOpaqueContainerTlv {
	obj := flowIpv6SRHOpaqueContainerTlv{obj: &otg.FlowIpv6SRHOpaqueContainerTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6SRHOpaqueContainerTlv) msg() *otg.FlowIpv6SRHOpaqueContainerTlv {
	return obj.obj
}

func (obj *flowIpv6SRHOpaqueContainerTlv) setMsg(msg *otg.FlowIpv6SRHOpaqueContainerTlv) FlowIpv6SRHOpaqueContainerTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6SRHOpaqueContainerTlv struct {
	obj *flowIpv6SRHOpaqueContainerTlv
}

type marshalFlowIpv6SRHOpaqueContainerTlv interface {
	// ToProto marshals FlowIpv6SRHOpaqueContainerTlv to protobuf object *otg.FlowIpv6SRHOpaqueContainerTlv
	ToProto() (*otg.FlowIpv6SRHOpaqueContainerTlv, error)
	// ToPbText marshals FlowIpv6SRHOpaqueContainerTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6SRHOpaqueContainerTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6SRHOpaqueContainerTlv to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6SRHOpaqueContainerTlv struct {
	obj *flowIpv6SRHOpaqueContainerTlv
}

type unMarshalFlowIpv6SRHOpaqueContainerTlv interface {
	// FromProto unmarshals FlowIpv6SRHOpaqueContainerTlv from protobuf object *otg.FlowIpv6SRHOpaqueContainerTlv
	FromProto(msg *otg.FlowIpv6SRHOpaqueContainerTlv) (FlowIpv6SRHOpaqueContainerTlv, error)
	// FromPbText unmarshals FlowIpv6SRHOpaqueContainerTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6SRHOpaqueContainerTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6SRHOpaqueContainerTlv from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6SRHOpaqueContainerTlv) Marshal() marshalFlowIpv6SRHOpaqueContainerTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6SRHOpaqueContainerTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6SRHOpaqueContainerTlv) Unmarshal() unMarshalFlowIpv6SRHOpaqueContainerTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6SRHOpaqueContainerTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6SRHOpaqueContainerTlv) ToProto() (*otg.FlowIpv6SRHOpaqueContainerTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6SRHOpaqueContainerTlv) FromProto(msg *otg.FlowIpv6SRHOpaqueContainerTlv) (FlowIpv6SRHOpaqueContainerTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6SRHOpaqueContainerTlv) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6SRHOpaqueContainerTlv) FromPbText(value string) error {
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

func (m *marshalflowIpv6SRHOpaqueContainerTlv) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6SRHOpaqueContainerTlv) FromYaml(value string) error {
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

func (m *marshalflowIpv6SRHOpaqueContainerTlv) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6SRHOpaqueContainerTlv) FromJson(value string) error {
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

func (obj *flowIpv6SRHOpaqueContainerTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6SRHOpaqueContainerTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6SRHOpaqueContainerTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6SRHOpaqueContainerTlv) Clone() (FlowIpv6SRHOpaqueContainerTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6SRHOpaqueContainerTlv()
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

func (obj *flowIpv6SRHOpaqueContainerTlv) setNil() {
	obj.typeHolder = nil
	obj.lengthHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv6SRHOpaqueContainerTlv is sRH Opaque Container TLV (type 3, RFC 8754 Section 2.1). Carries implementation-specific or application-defined opaque data in the SRH TLV section. Transit nodes do not interpret the contents. When present this TLV is included in the SRH TLV section; omit the object to suppress it. Reference: RFC 8754 Section 2.1.
type FlowIpv6SRHOpaqueContainerTlv interface {
	Validation
	// msg marshals FlowIpv6SRHOpaqueContainerTlv to protobuf object *otg.FlowIpv6SRHOpaqueContainerTlv
	// and doesn't set defaults
	msg() *otg.FlowIpv6SRHOpaqueContainerTlv
	// setMsg unmarshals FlowIpv6SRHOpaqueContainerTlv from protobuf object *otg.FlowIpv6SRHOpaqueContainerTlv
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6SRHOpaqueContainerTlv) FlowIpv6SRHOpaqueContainerTlv
	// provides marshal interface
	Marshal() marshalFlowIpv6SRHOpaqueContainerTlv
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6SRHOpaqueContainerTlv
	// validate validates FlowIpv6SRHOpaqueContainerTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6SRHOpaqueContainerTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns PatternFlowIpv6SRHOpaqueContainerTlvType, set in FlowIpv6SRHOpaqueContainerTlv.
	// PatternFlowIpv6SRHOpaqueContainerTlvType is tLV type code (RFC 8754 Section 2.1). The canonical value is 3. Set to a different value only for negative or conformance testing.
	Type() PatternFlowIpv6SRHOpaqueContainerTlvType
	// SetType assigns PatternFlowIpv6SRHOpaqueContainerTlvType provided by user to FlowIpv6SRHOpaqueContainerTlv.
	// PatternFlowIpv6SRHOpaqueContainerTlvType is tLV type code (RFC 8754 Section 2.1). The canonical value is 3. Set to a different value only for negative or conformance testing.
	SetType(value PatternFlowIpv6SRHOpaqueContainerTlvType) FlowIpv6SRHOpaqueContainerTlv
	// HasType checks if Type has been set in FlowIpv6SRHOpaqueContainerTlv
	HasType() bool
	// Length returns PatternFlowIpv6SRHOpaqueContainerTlvLength, set in FlowIpv6SRHOpaqueContainerTlv.
	// PatternFlowIpv6SRHOpaqueContainerTlvLength is length of the TLV Value field in bytes (RFC 8754 Section 2.1). When auto is assigned the implementation derives this from the length of the value string.
	Length() PatternFlowIpv6SRHOpaqueContainerTlvLength
	// SetLength assigns PatternFlowIpv6SRHOpaqueContainerTlvLength provided by user to FlowIpv6SRHOpaqueContainerTlv.
	// PatternFlowIpv6SRHOpaqueContainerTlvLength is length of the TLV Value field in bytes (RFC 8754 Section 2.1). When auto is assigned the implementation derives this from the length of the value string.
	SetLength(value PatternFlowIpv6SRHOpaqueContainerTlvLength) FlowIpv6SRHOpaqueContainerTlv
	// HasLength checks if Length has been set in FlowIpv6SRHOpaqueContainerTlv
	HasLength() bool
	// Value returns string, set in FlowIpv6SRHOpaqueContainerTlv.
	Value() string
	// SetValue assigns string provided by user to FlowIpv6SRHOpaqueContainerTlv
	SetValue(value string) FlowIpv6SRHOpaqueContainerTlv
	// HasValue checks if Value has been set in FlowIpv6SRHOpaqueContainerTlv
	HasValue() bool
	setNil()
}

// description is TBD
// Type returns a PatternFlowIpv6SRHOpaqueContainerTlvType
func (obj *flowIpv6SRHOpaqueContainerTlv) Type() PatternFlowIpv6SRHOpaqueContainerTlvType {
	if obj.obj.Type == nil {
		obj.obj.Type = NewPatternFlowIpv6SRHOpaqueContainerTlvType().msg()
	}
	if obj.typeHolder == nil {
		obj.typeHolder = &patternFlowIpv6SRHOpaqueContainerTlvType{obj: obj.obj.Type}
	}
	return obj.typeHolder
}

// description is TBD
// Type returns a PatternFlowIpv6SRHOpaqueContainerTlvType
func (obj *flowIpv6SRHOpaqueContainerTlv) HasType() bool {
	return obj.obj.Type != nil
}

// description is TBD
// SetType sets the PatternFlowIpv6SRHOpaqueContainerTlvType value in the FlowIpv6SRHOpaqueContainerTlv object
func (obj *flowIpv6SRHOpaqueContainerTlv) SetType(value PatternFlowIpv6SRHOpaqueContainerTlvType) FlowIpv6SRHOpaqueContainerTlv {

	obj.typeHolder = nil
	obj.obj.Type = value.msg()

	return obj
}

// description is TBD
// Length returns a PatternFlowIpv6SRHOpaqueContainerTlvLength
func (obj *flowIpv6SRHOpaqueContainerTlv) Length() PatternFlowIpv6SRHOpaqueContainerTlvLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewPatternFlowIpv6SRHOpaqueContainerTlvLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &patternFlowIpv6SRHOpaqueContainerTlvLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// description is TBD
// Length returns a PatternFlowIpv6SRHOpaqueContainerTlvLength
func (obj *flowIpv6SRHOpaqueContainerTlv) HasLength() bool {
	return obj.obj.Length != nil
}

// description is TBD
// SetLength sets the PatternFlowIpv6SRHOpaqueContainerTlvLength value in the FlowIpv6SRHOpaqueContainerTlv object
func (obj *flowIpv6SRHOpaqueContainerTlv) SetLength(value PatternFlowIpv6SRHOpaqueContainerTlvLength) FlowIpv6SRHOpaqueContainerTlv {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// Opaque data as a hex string (RFC 8754 Section 2.1). The content is implementation-defined and not interpreted by transit nodes. Example: "deadbeef01020304".
// Value returns a string
func (obj *flowIpv6SRHOpaqueContainerTlv) Value() string {

	return *obj.obj.Value

}

// Opaque data as a hex string (RFC 8754 Section 2.1). The content is implementation-defined and not interpreted by transit nodes. Example: "deadbeef01020304".
// Value returns a string
func (obj *flowIpv6SRHOpaqueContainerTlv) HasValue() bool {
	return obj.obj.Value != nil
}

// Opaque data as a hex string (RFC 8754 Section 2.1). The content is implementation-defined and not interpreted by transit nodes. Example: "deadbeef01020304".
// SetValue sets the string value in the FlowIpv6SRHOpaqueContainerTlv object
func (obj *flowIpv6SRHOpaqueContainerTlv) SetValue(value string) FlowIpv6SRHOpaqueContainerTlv {

	obj.obj.Value = &value
	return obj
}

func (obj *flowIpv6SRHOpaqueContainerTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Type != nil {

		obj.Type().validateObj(vObj, set_default)
	}

	if obj.obj.Length != nil {

		obj.Length().validateObj(vObj, set_default)
	}

	if obj.obj.Value != nil {

		err := obj.validateHex(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowIpv6SRHOpaqueContainerTlv.Value"))
		}

	}

}

func (obj *flowIpv6SRHOpaqueContainerTlv) setDefault() {

}
