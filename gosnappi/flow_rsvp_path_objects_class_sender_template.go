package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathObjectsClassSenderTemplate *****
type flowRSVPPathObjectsClassSenderTemplate struct {
	validation
	obj          *otg.FlowRSVPPathObjectsClassSenderTemplate
	marshaller   marshalFlowRSVPPathObjectsClassSenderTemplate
	unMarshaller unMarshalFlowRSVPPathObjectsClassSenderTemplate
	lengthHolder FlowRSVPObjectLength
	cTypeHolder  FlowRSVPPathObjectsSenderTemplateCType
}

func NewFlowRSVPPathObjectsClassSenderTemplate() FlowRSVPPathObjectsClassSenderTemplate {
	obj := flowRSVPPathObjectsClassSenderTemplate{obj: &otg.FlowRSVPPathObjectsClassSenderTemplate{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathObjectsClassSenderTemplate) msg() *otg.FlowRSVPPathObjectsClassSenderTemplate {
	return obj.obj
}

func (obj *flowRSVPPathObjectsClassSenderTemplate) setMsg(msg *otg.FlowRSVPPathObjectsClassSenderTemplate) FlowRSVPPathObjectsClassSenderTemplate {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathObjectsClassSenderTemplate struct {
	obj *flowRSVPPathObjectsClassSenderTemplate
}

type marshalFlowRSVPPathObjectsClassSenderTemplate interface {
	// ToProto marshals FlowRSVPPathObjectsClassSenderTemplate to protobuf object *otg.FlowRSVPPathObjectsClassSenderTemplate
	ToProto() (*otg.FlowRSVPPathObjectsClassSenderTemplate, error)
	// ToPbText marshals FlowRSVPPathObjectsClassSenderTemplate to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathObjectsClassSenderTemplate to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathObjectsClassSenderTemplate to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPPathObjectsClassSenderTemplate struct {
	obj *flowRSVPPathObjectsClassSenderTemplate
}

type unMarshalFlowRSVPPathObjectsClassSenderTemplate interface {
	// FromProto unmarshals FlowRSVPPathObjectsClassSenderTemplate from protobuf object *otg.FlowRSVPPathObjectsClassSenderTemplate
	FromProto(msg *otg.FlowRSVPPathObjectsClassSenderTemplate) (FlowRSVPPathObjectsClassSenderTemplate, error)
	// FromPbText unmarshals FlowRSVPPathObjectsClassSenderTemplate from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathObjectsClassSenderTemplate from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathObjectsClassSenderTemplate from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathObjectsClassSenderTemplate) Marshal() marshalFlowRSVPPathObjectsClassSenderTemplate {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathObjectsClassSenderTemplate{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathObjectsClassSenderTemplate) Unmarshal() unMarshalFlowRSVPPathObjectsClassSenderTemplate {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathObjectsClassSenderTemplate{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathObjectsClassSenderTemplate) ToProto() (*otg.FlowRSVPPathObjectsClassSenderTemplate, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathObjectsClassSenderTemplate) FromProto(msg *otg.FlowRSVPPathObjectsClassSenderTemplate) (FlowRSVPPathObjectsClassSenderTemplate, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathObjectsClassSenderTemplate) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassSenderTemplate) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathObjectsClassSenderTemplate) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassSenderTemplate) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathObjectsClassSenderTemplate) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassSenderTemplate) FromJson(value string) error {
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

func (obj *flowRSVPPathObjectsClassSenderTemplate) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsClassSenderTemplate) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsClassSenderTemplate) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathObjectsClassSenderTemplate) Clone() (FlowRSVPPathObjectsClassSenderTemplate, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathObjectsClassSenderTemplate()
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

func (obj *flowRSVPPathObjectsClassSenderTemplate) setNil() {
	obj.lengthHolder = nil
	obj.cTypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathObjectsClassSenderTemplate is c-Type is specific to a class num.
type FlowRSVPPathObjectsClassSenderTemplate interface {
	Validation
	// msg marshals FlowRSVPPathObjectsClassSenderTemplate to protobuf object *otg.FlowRSVPPathObjectsClassSenderTemplate
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathObjectsClassSenderTemplate
	// setMsg unmarshals FlowRSVPPathObjectsClassSenderTemplate from protobuf object *otg.FlowRSVPPathObjectsClassSenderTemplate
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathObjectsClassSenderTemplate) FlowRSVPPathObjectsClassSenderTemplate
	// provides marshal interface
	Marshal() marshalFlowRSVPPathObjectsClassSenderTemplate
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathObjectsClassSenderTemplate
	// validate validates FlowRSVPPathObjectsClassSenderTemplate
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathObjectsClassSenderTemplate, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Length returns FlowRSVPObjectLength, set in FlowRSVPPathObjectsClassSenderTemplate.
	// FlowRSVPObjectLength is description is TBD
	Length() FlowRSVPObjectLength
	// SetLength assigns FlowRSVPObjectLength provided by user to FlowRSVPPathObjectsClassSenderTemplate.
	// FlowRSVPObjectLength is description is TBD
	SetLength(value FlowRSVPObjectLength) FlowRSVPPathObjectsClassSenderTemplate
	// HasLength checks if Length has been set in FlowRSVPPathObjectsClassSenderTemplate
	HasLength() bool
	// CType returns FlowRSVPPathObjectsSenderTemplateCType, set in FlowRSVPPathObjectsClassSenderTemplate.
	// FlowRSVPPathObjectsSenderTemplateCType is object for SENDER_TEMPLATE class. Currently supported c-type is LSP Tunnel IPv4 (7).
	CType() FlowRSVPPathObjectsSenderTemplateCType
	// SetCType assigns FlowRSVPPathObjectsSenderTemplateCType provided by user to FlowRSVPPathObjectsClassSenderTemplate.
	// FlowRSVPPathObjectsSenderTemplateCType is object for SENDER_TEMPLATE class. Currently supported c-type is LSP Tunnel IPv4 (7).
	SetCType(value FlowRSVPPathObjectsSenderTemplateCType) FlowRSVPPathObjectsClassSenderTemplate
	// HasCType checks if CType has been set in FlowRSVPPathObjectsClassSenderTemplate
	HasCType() bool
	setNil()
}

// A 16-bit field containing the total object length in bytes.  Must always be a multiple of 4 or at least 4.
// Length returns a FlowRSVPObjectLength
func (obj *flowRSVPPathObjectsClassSenderTemplate) Length() FlowRSVPObjectLength {
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
func (obj *flowRSVPPathObjectsClassSenderTemplate) HasLength() bool {
	return obj.obj.Length != nil
}

// A 16-bit field containing the total object length in bytes.  Must always be a multiple of 4 or at least 4.
// SetLength sets the FlowRSVPObjectLength value in the FlowRSVPPathObjectsClassSenderTemplate object
func (obj *flowRSVPPathObjectsClassSenderTemplate) SetLength(value FlowRSVPObjectLength) FlowRSVPPathObjectsClassSenderTemplate {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// description is TBD
// CType returns a FlowRSVPPathObjectsSenderTemplateCType
func (obj *flowRSVPPathObjectsClassSenderTemplate) CType() FlowRSVPPathObjectsSenderTemplateCType {
	if obj.obj.CType == nil {
		obj.obj.CType = NewFlowRSVPPathObjectsSenderTemplateCType().msg()
	}
	if obj.cTypeHolder == nil {
		obj.cTypeHolder = &flowRSVPPathObjectsSenderTemplateCType{obj: obj.obj.CType}
	}
	return obj.cTypeHolder
}

// description is TBD
// CType returns a FlowRSVPPathObjectsSenderTemplateCType
func (obj *flowRSVPPathObjectsClassSenderTemplate) HasCType() bool {
	return obj.obj.CType != nil
}

// description is TBD
// SetCType sets the FlowRSVPPathObjectsSenderTemplateCType value in the FlowRSVPPathObjectsClassSenderTemplate object
func (obj *flowRSVPPathObjectsClassSenderTemplate) SetCType(value FlowRSVPPathObjectsSenderTemplateCType) FlowRSVPPathObjectsClassSenderTemplate {

	obj.cTypeHolder = nil
	obj.obj.CType = value.msg()

	return obj
}

func (obj *flowRSVPPathObjectsClassSenderTemplate) validateObj(vObj *validation, set_default bool) {
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

func (obj *flowRSVPPathObjectsClassSenderTemplate) setDefault() {

}
