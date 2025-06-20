package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathObjectsClassTimeValues *****
type flowRSVPPathObjectsClassTimeValues struct {
	validation
	obj          *otg.FlowRSVPPathObjectsClassTimeValues
	marshaller   marshalFlowRSVPPathObjectsClassTimeValues
	unMarshaller unMarshalFlowRSVPPathObjectsClassTimeValues
	lengthHolder FlowRSVPObjectLength
	cTypeHolder  FlowRSVPPathObjectsTimeValuesCType
}

func NewFlowRSVPPathObjectsClassTimeValues() FlowRSVPPathObjectsClassTimeValues {
	obj := flowRSVPPathObjectsClassTimeValues{obj: &otg.FlowRSVPPathObjectsClassTimeValues{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathObjectsClassTimeValues) msg() *otg.FlowRSVPPathObjectsClassTimeValues {
	return obj.obj
}

func (obj *flowRSVPPathObjectsClassTimeValues) setMsg(msg *otg.FlowRSVPPathObjectsClassTimeValues) FlowRSVPPathObjectsClassTimeValues {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathObjectsClassTimeValues struct {
	obj *flowRSVPPathObjectsClassTimeValues
}

type marshalFlowRSVPPathObjectsClassTimeValues interface {
	// ToProto marshals FlowRSVPPathObjectsClassTimeValues to protobuf object *otg.FlowRSVPPathObjectsClassTimeValues
	ToProto() (*otg.FlowRSVPPathObjectsClassTimeValues, error)
	// ToPbText marshals FlowRSVPPathObjectsClassTimeValues to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathObjectsClassTimeValues to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathObjectsClassTimeValues to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPPathObjectsClassTimeValues struct {
	obj *flowRSVPPathObjectsClassTimeValues
}

type unMarshalFlowRSVPPathObjectsClassTimeValues interface {
	// FromProto unmarshals FlowRSVPPathObjectsClassTimeValues from protobuf object *otg.FlowRSVPPathObjectsClassTimeValues
	FromProto(msg *otg.FlowRSVPPathObjectsClassTimeValues) (FlowRSVPPathObjectsClassTimeValues, error)
	// FromPbText unmarshals FlowRSVPPathObjectsClassTimeValues from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathObjectsClassTimeValues from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathObjectsClassTimeValues from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathObjectsClassTimeValues) Marshal() marshalFlowRSVPPathObjectsClassTimeValues {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathObjectsClassTimeValues{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathObjectsClassTimeValues) Unmarshal() unMarshalFlowRSVPPathObjectsClassTimeValues {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathObjectsClassTimeValues{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathObjectsClassTimeValues) ToProto() (*otg.FlowRSVPPathObjectsClassTimeValues, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathObjectsClassTimeValues) FromProto(msg *otg.FlowRSVPPathObjectsClassTimeValues) (FlowRSVPPathObjectsClassTimeValues, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathObjectsClassTimeValues) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassTimeValues) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathObjectsClassTimeValues) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassTimeValues) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathObjectsClassTimeValues) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassTimeValues) FromJson(value string) error {
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

func (obj *flowRSVPPathObjectsClassTimeValues) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsClassTimeValues) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsClassTimeValues) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathObjectsClassTimeValues) Clone() (FlowRSVPPathObjectsClassTimeValues, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathObjectsClassTimeValues()
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

func (obj *flowRSVPPathObjectsClassTimeValues) setNil() {
	obj.lengthHolder = nil
	obj.cTypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathObjectsClassTimeValues is c-Type is specific to a class num.
type FlowRSVPPathObjectsClassTimeValues interface {
	Validation
	// msg marshals FlowRSVPPathObjectsClassTimeValues to protobuf object *otg.FlowRSVPPathObjectsClassTimeValues
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathObjectsClassTimeValues
	// setMsg unmarshals FlowRSVPPathObjectsClassTimeValues from protobuf object *otg.FlowRSVPPathObjectsClassTimeValues
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathObjectsClassTimeValues) FlowRSVPPathObjectsClassTimeValues
	// provides marshal interface
	Marshal() marshalFlowRSVPPathObjectsClassTimeValues
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathObjectsClassTimeValues
	// validate validates FlowRSVPPathObjectsClassTimeValues
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathObjectsClassTimeValues, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Length returns FlowRSVPObjectLength, set in FlowRSVPPathObjectsClassTimeValues.
	// FlowRSVPObjectLength is description is TBD
	Length() FlowRSVPObjectLength
	// SetLength assigns FlowRSVPObjectLength provided by user to FlowRSVPPathObjectsClassTimeValues.
	// FlowRSVPObjectLength is description is TBD
	SetLength(value FlowRSVPObjectLength) FlowRSVPPathObjectsClassTimeValues
	// HasLength checks if Length has been set in FlowRSVPPathObjectsClassTimeValues
	HasLength() bool
	// CType returns FlowRSVPPathObjectsTimeValuesCType, set in FlowRSVPPathObjectsClassTimeValues.
	// FlowRSVPPathObjectsTimeValuesCType is object for TIME_VALUES class. Currently supported c-type is Type 1 Time Value (1).
	CType() FlowRSVPPathObjectsTimeValuesCType
	// SetCType assigns FlowRSVPPathObjectsTimeValuesCType provided by user to FlowRSVPPathObjectsClassTimeValues.
	// FlowRSVPPathObjectsTimeValuesCType is object for TIME_VALUES class. Currently supported c-type is Type 1 Time Value (1).
	SetCType(value FlowRSVPPathObjectsTimeValuesCType) FlowRSVPPathObjectsClassTimeValues
	// HasCType checks if CType has been set in FlowRSVPPathObjectsClassTimeValues
	HasCType() bool
	setNil()
}

// A 16-bit field containing the total object length in bytes.  Must always be a multiple of 4 or at least 4.
// Length returns a FlowRSVPObjectLength
func (obj *flowRSVPPathObjectsClassTimeValues) Length() FlowRSVPObjectLength {
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
func (obj *flowRSVPPathObjectsClassTimeValues) HasLength() bool {
	return obj.obj.Length != nil
}

// A 16-bit field containing the total object length in bytes.  Must always be a multiple of 4 or at least 4.
// SetLength sets the FlowRSVPObjectLength value in the FlowRSVPPathObjectsClassTimeValues object
func (obj *flowRSVPPathObjectsClassTimeValues) SetLength(value FlowRSVPObjectLength) FlowRSVPPathObjectsClassTimeValues {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// description is TBD
// CType returns a FlowRSVPPathObjectsTimeValuesCType
func (obj *flowRSVPPathObjectsClassTimeValues) CType() FlowRSVPPathObjectsTimeValuesCType {
	if obj.obj.CType == nil {
		obj.obj.CType = NewFlowRSVPPathObjectsTimeValuesCType().msg()
	}
	if obj.cTypeHolder == nil {
		obj.cTypeHolder = &flowRSVPPathObjectsTimeValuesCType{obj: obj.obj.CType}
	}
	return obj.cTypeHolder
}

// description is TBD
// CType returns a FlowRSVPPathObjectsTimeValuesCType
func (obj *flowRSVPPathObjectsClassTimeValues) HasCType() bool {
	return obj.obj.CType != nil
}

// description is TBD
// SetCType sets the FlowRSVPPathObjectsTimeValuesCType value in the FlowRSVPPathObjectsClassTimeValues object
func (obj *flowRSVPPathObjectsClassTimeValues) SetCType(value FlowRSVPPathObjectsTimeValuesCType) FlowRSVPPathObjectsClassTimeValues {

	obj.cTypeHolder = nil
	obj.obj.CType = value.msg()

	return obj
}

func (obj *flowRSVPPathObjectsClassTimeValues) validateObj(vObj *validation, set_default bool) {
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

func (obj *flowRSVPPathObjectsClassTimeValues) setDefault() {

}
