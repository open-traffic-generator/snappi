package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathObjectsClassSenderTspec *****
type flowRSVPPathObjectsClassSenderTspec struct {
	validation
	obj          *otg.FlowRSVPPathObjectsClassSenderTspec
	marshaller   marshalFlowRSVPPathObjectsClassSenderTspec
	unMarshaller unMarshalFlowRSVPPathObjectsClassSenderTspec
	lengthHolder FlowRSVPObjectLength
	cTypeHolder  FlowRSVPPathObjectsSenderTspecCType
}

func NewFlowRSVPPathObjectsClassSenderTspec() FlowRSVPPathObjectsClassSenderTspec {
	obj := flowRSVPPathObjectsClassSenderTspec{obj: &otg.FlowRSVPPathObjectsClassSenderTspec{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathObjectsClassSenderTspec) msg() *otg.FlowRSVPPathObjectsClassSenderTspec {
	return obj.obj
}

func (obj *flowRSVPPathObjectsClassSenderTspec) setMsg(msg *otg.FlowRSVPPathObjectsClassSenderTspec) FlowRSVPPathObjectsClassSenderTspec {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathObjectsClassSenderTspec struct {
	obj *flowRSVPPathObjectsClassSenderTspec
}

type marshalFlowRSVPPathObjectsClassSenderTspec interface {
	// ToProto marshals FlowRSVPPathObjectsClassSenderTspec to protobuf object *otg.FlowRSVPPathObjectsClassSenderTspec
	ToProto() (*otg.FlowRSVPPathObjectsClassSenderTspec, error)
	// ToPbText marshals FlowRSVPPathObjectsClassSenderTspec to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathObjectsClassSenderTspec to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathObjectsClassSenderTspec to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowRSVPPathObjectsClassSenderTspec to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowRSVPPathObjectsClassSenderTspec struct {
	obj *flowRSVPPathObjectsClassSenderTspec
}

type unMarshalFlowRSVPPathObjectsClassSenderTspec interface {
	// FromProto unmarshals FlowRSVPPathObjectsClassSenderTspec from protobuf object *otg.FlowRSVPPathObjectsClassSenderTspec
	FromProto(msg *otg.FlowRSVPPathObjectsClassSenderTspec) (FlowRSVPPathObjectsClassSenderTspec, error)
	// FromPbText unmarshals FlowRSVPPathObjectsClassSenderTspec from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathObjectsClassSenderTspec from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathObjectsClassSenderTspec from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathObjectsClassSenderTspec) Marshal() marshalFlowRSVPPathObjectsClassSenderTspec {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathObjectsClassSenderTspec{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathObjectsClassSenderTspec) Unmarshal() unMarshalFlowRSVPPathObjectsClassSenderTspec {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathObjectsClassSenderTspec{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathObjectsClassSenderTspec) ToProto() (*otg.FlowRSVPPathObjectsClassSenderTspec, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathObjectsClassSenderTspec) FromProto(msg *otg.FlowRSVPPathObjectsClassSenderTspec) (FlowRSVPPathObjectsClassSenderTspec, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathObjectsClassSenderTspec) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassSenderTspec) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathObjectsClassSenderTspec) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassSenderTspec) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathObjectsClassSenderTspec) ToJsonRaw() (string, error) {
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

func (m *marshalflowRSVPPathObjectsClassSenderTspec) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassSenderTspec) FromJson(value string) error {
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

func (obj *flowRSVPPathObjectsClassSenderTspec) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsClassSenderTspec) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsClassSenderTspec) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathObjectsClassSenderTspec) Clone() (FlowRSVPPathObjectsClassSenderTspec, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathObjectsClassSenderTspec()
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

func (obj *flowRSVPPathObjectsClassSenderTspec) setNil() {
	obj.lengthHolder = nil
	obj.cTypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathObjectsClassSenderTspec is c-Type is specific to a class num.
type FlowRSVPPathObjectsClassSenderTspec interface {
	Validation
	// msg marshals FlowRSVPPathObjectsClassSenderTspec to protobuf object *otg.FlowRSVPPathObjectsClassSenderTspec
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathObjectsClassSenderTspec
	// setMsg unmarshals FlowRSVPPathObjectsClassSenderTspec from protobuf object *otg.FlowRSVPPathObjectsClassSenderTspec
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathObjectsClassSenderTspec) FlowRSVPPathObjectsClassSenderTspec
	// provides marshal interface
	Marshal() marshalFlowRSVPPathObjectsClassSenderTspec
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathObjectsClassSenderTspec
	// validate validates FlowRSVPPathObjectsClassSenderTspec
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathObjectsClassSenderTspec, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Length returns FlowRSVPObjectLength, set in FlowRSVPPathObjectsClassSenderTspec.
	// FlowRSVPObjectLength is description is TBD
	Length() FlowRSVPObjectLength
	// SetLength assigns FlowRSVPObjectLength provided by user to FlowRSVPPathObjectsClassSenderTspec.
	// FlowRSVPObjectLength is description is TBD
	SetLength(value FlowRSVPObjectLength) FlowRSVPPathObjectsClassSenderTspec
	// HasLength checks if Length has been set in FlowRSVPPathObjectsClassSenderTspec
	HasLength() bool
	// CType returns FlowRSVPPathObjectsSenderTspecCType, set in FlowRSVPPathObjectsClassSenderTspec.
	// FlowRSVPPathObjectsSenderTspecCType is object for SENDER_TSPEC class. Currently supported c-type is int-serv (2).
	CType() FlowRSVPPathObjectsSenderTspecCType
	// SetCType assigns FlowRSVPPathObjectsSenderTspecCType provided by user to FlowRSVPPathObjectsClassSenderTspec.
	// FlowRSVPPathObjectsSenderTspecCType is object for SENDER_TSPEC class. Currently supported c-type is int-serv (2).
	SetCType(value FlowRSVPPathObjectsSenderTspecCType) FlowRSVPPathObjectsClassSenderTspec
	// HasCType checks if CType has been set in FlowRSVPPathObjectsClassSenderTspec
	HasCType() bool
	setNil()
}

// A 16-bit field containing the total object length in bytes.  Must always be a multiple of 4 or at least 4.
// Length returns a FlowRSVPObjectLength
func (obj *flowRSVPPathObjectsClassSenderTspec) Length() FlowRSVPObjectLength {
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
func (obj *flowRSVPPathObjectsClassSenderTspec) HasLength() bool {
	return obj.obj.Length != nil
}

// A 16-bit field containing the total object length in bytes.  Must always be a multiple of 4 or at least 4.
// SetLength sets the FlowRSVPObjectLength value in the FlowRSVPPathObjectsClassSenderTspec object
func (obj *flowRSVPPathObjectsClassSenderTspec) SetLength(value FlowRSVPObjectLength) FlowRSVPPathObjectsClassSenderTspec {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// description is TBD
// CType returns a FlowRSVPPathObjectsSenderTspecCType
func (obj *flowRSVPPathObjectsClassSenderTspec) CType() FlowRSVPPathObjectsSenderTspecCType {
	if obj.obj.CType == nil {
		obj.obj.CType = NewFlowRSVPPathObjectsSenderTspecCType().msg()
	}
	if obj.cTypeHolder == nil {
		obj.cTypeHolder = &flowRSVPPathObjectsSenderTspecCType{obj: obj.obj.CType}
	}
	return obj.cTypeHolder
}

// description is TBD
// CType returns a FlowRSVPPathObjectsSenderTspecCType
func (obj *flowRSVPPathObjectsClassSenderTspec) HasCType() bool {
	return obj.obj.CType != nil
}

// description is TBD
// SetCType sets the FlowRSVPPathObjectsSenderTspecCType value in the FlowRSVPPathObjectsClassSenderTspec object
func (obj *flowRSVPPathObjectsClassSenderTspec) SetCType(value FlowRSVPPathObjectsSenderTspecCType) FlowRSVPPathObjectsClassSenderTspec {

	obj.cTypeHolder = nil
	obj.obj.CType = value.msg()

	return obj
}

func (obj *flowRSVPPathObjectsClassSenderTspec) validateObj(vObj *validation, set_default bool) {
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

func (obj *flowRSVPPathObjectsClassSenderTspec) setDefault() {

}
