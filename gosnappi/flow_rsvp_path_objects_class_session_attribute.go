package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathObjectsClassSessionAttribute *****
type flowRSVPPathObjectsClassSessionAttribute struct {
	validation
	obj          *otg.FlowRSVPPathObjectsClassSessionAttribute
	marshaller   marshalFlowRSVPPathObjectsClassSessionAttribute
	unMarshaller unMarshalFlowRSVPPathObjectsClassSessionAttribute
	lengthHolder FlowRSVPObjectLength
	cTypeHolder  FlowRSVPPathObjectsSessionAttributeCType
}

func NewFlowRSVPPathObjectsClassSessionAttribute() FlowRSVPPathObjectsClassSessionAttribute {
	obj := flowRSVPPathObjectsClassSessionAttribute{obj: &otg.FlowRSVPPathObjectsClassSessionAttribute{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathObjectsClassSessionAttribute) msg() *otg.FlowRSVPPathObjectsClassSessionAttribute {
	return obj.obj
}

func (obj *flowRSVPPathObjectsClassSessionAttribute) setMsg(msg *otg.FlowRSVPPathObjectsClassSessionAttribute) FlowRSVPPathObjectsClassSessionAttribute {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathObjectsClassSessionAttribute struct {
	obj *flowRSVPPathObjectsClassSessionAttribute
}

type marshalFlowRSVPPathObjectsClassSessionAttribute interface {
	// ToProto marshals FlowRSVPPathObjectsClassSessionAttribute to protobuf object *otg.FlowRSVPPathObjectsClassSessionAttribute
	ToProto() (*otg.FlowRSVPPathObjectsClassSessionAttribute, error)
	// ToPbText marshals FlowRSVPPathObjectsClassSessionAttribute to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathObjectsClassSessionAttribute to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathObjectsClassSessionAttribute to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPPathObjectsClassSessionAttribute struct {
	obj *flowRSVPPathObjectsClassSessionAttribute
}

type unMarshalFlowRSVPPathObjectsClassSessionAttribute interface {
	// FromProto unmarshals FlowRSVPPathObjectsClassSessionAttribute from protobuf object *otg.FlowRSVPPathObjectsClassSessionAttribute
	FromProto(msg *otg.FlowRSVPPathObjectsClassSessionAttribute) (FlowRSVPPathObjectsClassSessionAttribute, error)
	// FromPbText unmarshals FlowRSVPPathObjectsClassSessionAttribute from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathObjectsClassSessionAttribute from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathObjectsClassSessionAttribute from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathObjectsClassSessionAttribute) Marshal() marshalFlowRSVPPathObjectsClassSessionAttribute {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathObjectsClassSessionAttribute{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathObjectsClassSessionAttribute) Unmarshal() unMarshalFlowRSVPPathObjectsClassSessionAttribute {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathObjectsClassSessionAttribute{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathObjectsClassSessionAttribute) ToProto() (*otg.FlowRSVPPathObjectsClassSessionAttribute, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathObjectsClassSessionAttribute) FromProto(msg *otg.FlowRSVPPathObjectsClassSessionAttribute) (FlowRSVPPathObjectsClassSessionAttribute, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathObjectsClassSessionAttribute) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassSessionAttribute) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathObjectsClassSessionAttribute) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassSessionAttribute) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathObjectsClassSessionAttribute) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassSessionAttribute) FromJson(value string) error {
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

func (obj *flowRSVPPathObjectsClassSessionAttribute) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsClassSessionAttribute) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsClassSessionAttribute) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathObjectsClassSessionAttribute) Clone() (FlowRSVPPathObjectsClassSessionAttribute, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathObjectsClassSessionAttribute()
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

func (obj *flowRSVPPathObjectsClassSessionAttribute) setNil() {
	obj.lengthHolder = nil
	obj.cTypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathObjectsClassSessionAttribute is c-Type is specific to a class num.
type FlowRSVPPathObjectsClassSessionAttribute interface {
	Validation
	// msg marshals FlowRSVPPathObjectsClassSessionAttribute to protobuf object *otg.FlowRSVPPathObjectsClassSessionAttribute
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathObjectsClassSessionAttribute
	// setMsg unmarshals FlowRSVPPathObjectsClassSessionAttribute from protobuf object *otg.FlowRSVPPathObjectsClassSessionAttribute
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathObjectsClassSessionAttribute) FlowRSVPPathObjectsClassSessionAttribute
	// provides marshal interface
	Marshal() marshalFlowRSVPPathObjectsClassSessionAttribute
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathObjectsClassSessionAttribute
	// validate validates FlowRSVPPathObjectsClassSessionAttribute
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathObjectsClassSessionAttribute, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Length returns FlowRSVPObjectLength, set in FlowRSVPPathObjectsClassSessionAttribute.
	// FlowRSVPObjectLength is description is TBD
	Length() FlowRSVPObjectLength
	// SetLength assigns FlowRSVPObjectLength provided by user to FlowRSVPPathObjectsClassSessionAttribute.
	// FlowRSVPObjectLength is description is TBD
	SetLength(value FlowRSVPObjectLength) FlowRSVPPathObjectsClassSessionAttribute
	// HasLength checks if Length has been set in FlowRSVPPathObjectsClassSessionAttribute
	HasLength() bool
	// CType returns FlowRSVPPathObjectsSessionAttributeCType, set in FlowRSVPPathObjectsClassSessionAttribute.
	// FlowRSVPPathObjectsSessionAttributeCType is object for SESSION_ATTRIBUTE class. Currently supported c-type is LSP_Tunnel_RA (1) and LSP_Tunnel (7).
	CType() FlowRSVPPathObjectsSessionAttributeCType
	// SetCType assigns FlowRSVPPathObjectsSessionAttributeCType provided by user to FlowRSVPPathObjectsClassSessionAttribute.
	// FlowRSVPPathObjectsSessionAttributeCType is object for SESSION_ATTRIBUTE class. Currently supported c-type is LSP_Tunnel_RA (1) and LSP_Tunnel (7).
	SetCType(value FlowRSVPPathObjectsSessionAttributeCType) FlowRSVPPathObjectsClassSessionAttribute
	// HasCType checks if CType has been set in FlowRSVPPathObjectsClassSessionAttribute
	HasCType() bool
	setNil()
}

// A 16-bit field containing the total object length in bytes.  Must always be a multiple of 4 or at least 4.
// Length returns a FlowRSVPObjectLength
func (obj *flowRSVPPathObjectsClassSessionAttribute) Length() FlowRSVPObjectLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewFlowRSVPObjectLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &flowRSVPObjectLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// A 16-bit field containing the total object length in bytes.  Must always be a multiple of 4 or at least 4.
// Length returns a FlowRSVPObjectLength
func (obj *flowRSVPPathObjectsClassSessionAttribute) HasLength() bool {
	return obj.obj.Length != nil
}

// A 16-bit field containing the total object length in bytes.  Must always be a multiple of 4 or at least 4.
// SetLength sets the FlowRSVPObjectLength value in the FlowRSVPPathObjectsClassSessionAttribute object
func (obj *flowRSVPPathObjectsClassSessionAttribute) SetLength(value FlowRSVPObjectLength) FlowRSVPPathObjectsClassSessionAttribute {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// description is TBD
// CType returns a FlowRSVPPathObjectsSessionAttributeCType
func (obj *flowRSVPPathObjectsClassSessionAttribute) CType() FlowRSVPPathObjectsSessionAttributeCType {
	if obj.obj.CType == nil {
		obj.obj.CType = NewFlowRSVPPathObjectsSessionAttributeCType().msg()
	}
	if obj.cTypeHolder == nil {
		obj.cTypeHolder = &flowRSVPPathObjectsSessionAttributeCType{obj: obj.obj.CType}
	}
	return obj.cTypeHolder
}

// description is TBD
// CType returns a FlowRSVPPathObjectsSessionAttributeCType
func (obj *flowRSVPPathObjectsClassSessionAttribute) HasCType() bool {
	return obj.obj.CType != nil
}

// description is TBD
// SetCType sets the FlowRSVPPathObjectsSessionAttributeCType value in the FlowRSVPPathObjectsClassSessionAttribute object
func (obj *flowRSVPPathObjectsClassSessionAttribute) SetCType(value FlowRSVPPathObjectsSessionAttributeCType) FlowRSVPPathObjectsClassSessionAttribute {

	obj.cTypeHolder = nil
	obj.obj.CType = value.msg()

	return obj
}

func (obj *flowRSVPPathObjectsClassSessionAttribute) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Length != nil {

		obj.Length().validateObj(vObj, set_default)
	}

	if obj.obj.CType != nil {

		obj.CType().validateObj(vObj, set_default)
	}

}

func (obj *flowRSVPPathObjectsClassSessionAttribute) setDefault() {

}
