package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathObjectsClassSession *****
type flowRSVPPathObjectsClassSession struct {
	validation
	obj          *otg.FlowRSVPPathObjectsClassSession
	marshaller   marshalFlowRSVPPathObjectsClassSession
	unMarshaller unMarshalFlowRSVPPathObjectsClassSession
	lengthHolder FlowRSVPObjectLength
	cTypeHolder  FlowRSVPPathObjectsSessionCType
}

func NewFlowRSVPPathObjectsClassSession() FlowRSVPPathObjectsClassSession {
	obj := flowRSVPPathObjectsClassSession{obj: &otg.FlowRSVPPathObjectsClassSession{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathObjectsClassSession) msg() *otg.FlowRSVPPathObjectsClassSession {
	return obj.obj
}

func (obj *flowRSVPPathObjectsClassSession) setMsg(msg *otg.FlowRSVPPathObjectsClassSession) FlowRSVPPathObjectsClassSession {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathObjectsClassSession struct {
	obj *flowRSVPPathObjectsClassSession
}

type marshalFlowRSVPPathObjectsClassSession interface {
	// ToProto marshals FlowRSVPPathObjectsClassSession to protobuf object *otg.FlowRSVPPathObjectsClassSession
	ToProto() (*otg.FlowRSVPPathObjectsClassSession, error)
	// ToPbText marshals FlowRSVPPathObjectsClassSession to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathObjectsClassSession to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathObjectsClassSession to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowRSVPPathObjectsClassSession to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowRSVPPathObjectsClassSession struct {
	obj *flowRSVPPathObjectsClassSession
}

type unMarshalFlowRSVPPathObjectsClassSession interface {
	// FromProto unmarshals FlowRSVPPathObjectsClassSession from protobuf object *otg.FlowRSVPPathObjectsClassSession
	FromProto(msg *otg.FlowRSVPPathObjectsClassSession) (FlowRSVPPathObjectsClassSession, error)
	// FromPbText unmarshals FlowRSVPPathObjectsClassSession from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathObjectsClassSession from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathObjectsClassSession from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathObjectsClassSession) Marshal() marshalFlowRSVPPathObjectsClassSession {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathObjectsClassSession{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathObjectsClassSession) Unmarshal() unMarshalFlowRSVPPathObjectsClassSession {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathObjectsClassSession{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathObjectsClassSession) ToProto() (*otg.FlowRSVPPathObjectsClassSession, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathObjectsClassSession) FromProto(msg *otg.FlowRSVPPathObjectsClassSession) (FlowRSVPPathObjectsClassSession, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathObjectsClassSession) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassSession) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathObjectsClassSession) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassSession) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathObjectsClassSession) ToJsonRaw() (string, error) {
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

func (m *marshalflowRSVPPathObjectsClassSession) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassSession) FromJson(value string) error {
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

func (obj *flowRSVPPathObjectsClassSession) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsClassSession) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsClassSession) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathObjectsClassSession) Clone() (FlowRSVPPathObjectsClassSession, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathObjectsClassSession()
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

func (obj *flowRSVPPathObjectsClassSession) setNil() {
	obj.lengthHolder = nil
	obj.cTypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathObjectsClassSession is c-Type is specific to a class num.
type FlowRSVPPathObjectsClassSession interface {
	Validation
	// msg marshals FlowRSVPPathObjectsClassSession to protobuf object *otg.FlowRSVPPathObjectsClassSession
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathObjectsClassSession
	// setMsg unmarshals FlowRSVPPathObjectsClassSession from protobuf object *otg.FlowRSVPPathObjectsClassSession
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathObjectsClassSession) FlowRSVPPathObjectsClassSession
	// provides marshal interface
	Marshal() marshalFlowRSVPPathObjectsClassSession
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathObjectsClassSession
	// validate validates FlowRSVPPathObjectsClassSession
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathObjectsClassSession, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Length returns FlowRSVPObjectLength, set in FlowRSVPPathObjectsClassSession.
	// FlowRSVPObjectLength is description is TBD
	Length() FlowRSVPObjectLength
	// SetLength assigns FlowRSVPObjectLength provided by user to FlowRSVPPathObjectsClassSession.
	// FlowRSVPObjectLength is description is TBD
	SetLength(value FlowRSVPObjectLength) FlowRSVPPathObjectsClassSession
	// HasLength checks if Length has been set in FlowRSVPPathObjectsClassSession
	HasLength() bool
	// CType returns FlowRSVPPathObjectsSessionCType, set in FlowRSVPPathObjectsClassSession.
	// FlowRSVPPathObjectsSessionCType is the body of an object corresponding to the class number and c-type. Currently supported c-type for SESSION object is LSP Tunnel IPv4 (7).
	CType() FlowRSVPPathObjectsSessionCType
	// SetCType assigns FlowRSVPPathObjectsSessionCType provided by user to FlowRSVPPathObjectsClassSession.
	// FlowRSVPPathObjectsSessionCType is the body of an object corresponding to the class number and c-type. Currently supported c-type for SESSION object is LSP Tunnel IPv4 (7).
	SetCType(value FlowRSVPPathObjectsSessionCType) FlowRSVPPathObjectsClassSession
	// HasCType checks if CType has been set in FlowRSVPPathObjectsClassSession
	HasCType() bool
	setNil()
}

// A 16-bit field containing the total object length in bytes.  Must always be a multiple of 4 or at least 4.
// Length returns a FlowRSVPObjectLength
func (obj *flowRSVPPathObjectsClassSession) Length() FlowRSVPObjectLength {
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
func (obj *flowRSVPPathObjectsClassSession) HasLength() bool {
	return obj.obj.Length != nil
}

// A 16-bit field containing the total object length in bytes.  Must always be a multiple of 4 or at least 4.
// SetLength sets the FlowRSVPObjectLength value in the FlowRSVPPathObjectsClassSession object
func (obj *flowRSVPPathObjectsClassSession) SetLength(value FlowRSVPObjectLength) FlowRSVPPathObjectsClassSession {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// description is TBD
// CType returns a FlowRSVPPathObjectsSessionCType
func (obj *flowRSVPPathObjectsClassSession) CType() FlowRSVPPathObjectsSessionCType {
	if obj.obj.CType == nil {
		obj.obj.CType = NewFlowRSVPPathObjectsSessionCType().msg()
	}
	if obj.cTypeHolder == nil {
		obj.cTypeHolder = &flowRSVPPathObjectsSessionCType{obj: obj.obj.CType}
	}
	return obj.cTypeHolder
}

// description is TBD
// CType returns a FlowRSVPPathObjectsSessionCType
func (obj *flowRSVPPathObjectsClassSession) HasCType() bool {
	return obj.obj.CType != nil
}

// description is TBD
// SetCType sets the FlowRSVPPathObjectsSessionCType value in the FlowRSVPPathObjectsClassSession object
func (obj *flowRSVPPathObjectsClassSession) SetCType(value FlowRSVPPathObjectsSessionCType) FlowRSVPPathObjectsClassSession {

	obj.cTypeHolder = nil
	obj.obj.CType = value.msg()

	return obj
}

func (obj *flowRSVPPathObjectsClassSession) validateObj(vObj *validation, set_default bool) {
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

func (obj *flowRSVPPathObjectsClassSession) setDefault() {

}
