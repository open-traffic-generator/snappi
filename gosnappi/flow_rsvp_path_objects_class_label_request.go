package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathObjectsClassLabelRequest *****
type flowRSVPPathObjectsClassLabelRequest struct {
	validation
	obj          *otg.FlowRSVPPathObjectsClassLabelRequest
	marshaller   marshalFlowRSVPPathObjectsClassLabelRequest
	unMarshaller unMarshalFlowRSVPPathObjectsClassLabelRequest
	lengthHolder FlowRSVPObjectLength
	cTypeHolder  FlowRSVPPathObjectsLabelRequestCType
}

func NewFlowRSVPPathObjectsClassLabelRequest() FlowRSVPPathObjectsClassLabelRequest {
	obj := flowRSVPPathObjectsClassLabelRequest{obj: &otg.FlowRSVPPathObjectsClassLabelRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathObjectsClassLabelRequest) msg() *otg.FlowRSVPPathObjectsClassLabelRequest {
	return obj.obj
}

func (obj *flowRSVPPathObjectsClassLabelRequest) setMsg(msg *otg.FlowRSVPPathObjectsClassLabelRequest) FlowRSVPPathObjectsClassLabelRequest {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathObjectsClassLabelRequest struct {
	obj *flowRSVPPathObjectsClassLabelRequest
}

type marshalFlowRSVPPathObjectsClassLabelRequest interface {
	// ToProto marshals FlowRSVPPathObjectsClassLabelRequest to protobuf object *otg.FlowRSVPPathObjectsClassLabelRequest
	ToProto() (*otg.FlowRSVPPathObjectsClassLabelRequest, error)
	// ToPbText marshals FlowRSVPPathObjectsClassLabelRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathObjectsClassLabelRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathObjectsClassLabelRequest to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowRSVPPathObjectsClassLabelRequest to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowRSVPPathObjectsClassLabelRequest struct {
	obj *flowRSVPPathObjectsClassLabelRequest
}

type unMarshalFlowRSVPPathObjectsClassLabelRequest interface {
	// FromProto unmarshals FlowRSVPPathObjectsClassLabelRequest from protobuf object *otg.FlowRSVPPathObjectsClassLabelRequest
	FromProto(msg *otg.FlowRSVPPathObjectsClassLabelRequest) (FlowRSVPPathObjectsClassLabelRequest, error)
	// FromPbText unmarshals FlowRSVPPathObjectsClassLabelRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathObjectsClassLabelRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathObjectsClassLabelRequest from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathObjectsClassLabelRequest) Marshal() marshalFlowRSVPPathObjectsClassLabelRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathObjectsClassLabelRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathObjectsClassLabelRequest) Unmarshal() unMarshalFlowRSVPPathObjectsClassLabelRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathObjectsClassLabelRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathObjectsClassLabelRequest) ToProto() (*otg.FlowRSVPPathObjectsClassLabelRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathObjectsClassLabelRequest) FromProto(msg *otg.FlowRSVPPathObjectsClassLabelRequest) (FlowRSVPPathObjectsClassLabelRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathObjectsClassLabelRequest) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassLabelRequest) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathObjectsClassLabelRequest) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassLabelRequest) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathObjectsClassLabelRequest) ToJsonRaw() (string, error) {
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

func (m *marshalflowRSVPPathObjectsClassLabelRequest) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassLabelRequest) FromJson(value string) error {
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

func (obj *flowRSVPPathObjectsClassLabelRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsClassLabelRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsClassLabelRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathObjectsClassLabelRequest) Clone() (FlowRSVPPathObjectsClassLabelRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathObjectsClassLabelRequest()
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

func (obj *flowRSVPPathObjectsClassLabelRequest) setNil() {
	obj.lengthHolder = nil
	obj.cTypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathObjectsClassLabelRequest is c-Type is specific to a class num.
type FlowRSVPPathObjectsClassLabelRequest interface {
	Validation
	// msg marshals FlowRSVPPathObjectsClassLabelRequest to protobuf object *otg.FlowRSVPPathObjectsClassLabelRequest
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathObjectsClassLabelRequest
	// setMsg unmarshals FlowRSVPPathObjectsClassLabelRequest from protobuf object *otg.FlowRSVPPathObjectsClassLabelRequest
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathObjectsClassLabelRequest) FlowRSVPPathObjectsClassLabelRequest
	// provides marshal interface
	Marshal() marshalFlowRSVPPathObjectsClassLabelRequest
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathObjectsClassLabelRequest
	// validate validates FlowRSVPPathObjectsClassLabelRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathObjectsClassLabelRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Length returns FlowRSVPObjectLength, set in FlowRSVPPathObjectsClassLabelRequest.
	// FlowRSVPObjectLength is description is TBD
	Length() FlowRSVPObjectLength
	// SetLength assigns FlowRSVPObjectLength provided by user to FlowRSVPPathObjectsClassLabelRequest.
	// FlowRSVPObjectLength is description is TBD
	SetLength(value FlowRSVPObjectLength) FlowRSVPPathObjectsClassLabelRequest
	// HasLength checks if Length has been set in FlowRSVPPathObjectsClassLabelRequest
	HasLength() bool
	// CType returns FlowRSVPPathObjectsLabelRequestCType, set in FlowRSVPPathObjectsClassLabelRequest.
	// FlowRSVPPathObjectsLabelRequestCType is object for LABEL_REQUEST class. Currently supported c-type is Without Label Range (1).
	CType() FlowRSVPPathObjectsLabelRequestCType
	// SetCType assigns FlowRSVPPathObjectsLabelRequestCType provided by user to FlowRSVPPathObjectsClassLabelRequest.
	// FlowRSVPPathObjectsLabelRequestCType is object for LABEL_REQUEST class. Currently supported c-type is Without Label Range (1).
	SetCType(value FlowRSVPPathObjectsLabelRequestCType) FlowRSVPPathObjectsClassLabelRequest
	// HasCType checks if CType has been set in FlowRSVPPathObjectsClassLabelRequest
	HasCType() bool
	setNil()
}

// A 16-bit field containing the total object length in bytes.  Must always be a multiple of 4 or at least 4.
// Length returns a FlowRSVPObjectLength
func (obj *flowRSVPPathObjectsClassLabelRequest) Length() FlowRSVPObjectLength {
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
func (obj *flowRSVPPathObjectsClassLabelRequest) HasLength() bool {
	return obj.obj.Length != nil
}

// A 16-bit field containing the total object length in bytes.  Must always be a multiple of 4 or at least 4.
// SetLength sets the FlowRSVPObjectLength value in the FlowRSVPPathObjectsClassLabelRequest object
func (obj *flowRSVPPathObjectsClassLabelRequest) SetLength(value FlowRSVPObjectLength) FlowRSVPPathObjectsClassLabelRequest {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// description is TBD
// CType returns a FlowRSVPPathObjectsLabelRequestCType
func (obj *flowRSVPPathObjectsClassLabelRequest) CType() FlowRSVPPathObjectsLabelRequestCType {
	if obj.obj.CType == nil {
		obj.obj.CType = NewFlowRSVPPathObjectsLabelRequestCType().msg()
	}
	if obj.cTypeHolder == nil {
		obj.cTypeHolder = &flowRSVPPathObjectsLabelRequestCType{obj: obj.obj.CType}
	}
	return obj.cTypeHolder
}

// description is TBD
// CType returns a FlowRSVPPathObjectsLabelRequestCType
func (obj *flowRSVPPathObjectsClassLabelRequest) HasCType() bool {
	return obj.obj.CType != nil
}

// description is TBD
// SetCType sets the FlowRSVPPathObjectsLabelRequestCType value in the FlowRSVPPathObjectsClassLabelRequest object
func (obj *flowRSVPPathObjectsClassLabelRequest) SetCType(value FlowRSVPPathObjectsLabelRequestCType) FlowRSVPPathObjectsClassLabelRequest {

	obj.cTypeHolder = nil
	obj.obj.CType = value.msg()

	return obj
}

func (obj *flowRSVPPathObjectsClassLabelRequest) validateObj(vObj *validation, set_default bool) {
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

func (obj *flowRSVPPathObjectsClassLabelRequest) setDefault() {

}
