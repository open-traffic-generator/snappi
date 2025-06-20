package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathObjectsClassRecordRoute *****
type flowRSVPPathObjectsClassRecordRoute struct {
	validation
	obj          *otg.FlowRSVPPathObjectsClassRecordRoute
	marshaller   marshalFlowRSVPPathObjectsClassRecordRoute
	unMarshaller unMarshalFlowRSVPPathObjectsClassRecordRoute
	lengthHolder FlowRSVPObjectLength
	cTypeHolder  FlowRSVPPathObjectsRecordRouteCType
}

func NewFlowRSVPPathObjectsClassRecordRoute() FlowRSVPPathObjectsClassRecordRoute {
	obj := flowRSVPPathObjectsClassRecordRoute{obj: &otg.FlowRSVPPathObjectsClassRecordRoute{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathObjectsClassRecordRoute) msg() *otg.FlowRSVPPathObjectsClassRecordRoute {
	return obj.obj
}

func (obj *flowRSVPPathObjectsClassRecordRoute) setMsg(msg *otg.FlowRSVPPathObjectsClassRecordRoute) FlowRSVPPathObjectsClassRecordRoute {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathObjectsClassRecordRoute struct {
	obj *flowRSVPPathObjectsClassRecordRoute
}

type marshalFlowRSVPPathObjectsClassRecordRoute interface {
	// ToProto marshals FlowRSVPPathObjectsClassRecordRoute to protobuf object *otg.FlowRSVPPathObjectsClassRecordRoute
	ToProto() (*otg.FlowRSVPPathObjectsClassRecordRoute, error)
	// ToPbText marshals FlowRSVPPathObjectsClassRecordRoute to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathObjectsClassRecordRoute to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathObjectsClassRecordRoute to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPPathObjectsClassRecordRoute struct {
	obj *flowRSVPPathObjectsClassRecordRoute
}

type unMarshalFlowRSVPPathObjectsClassRecordRoute interface {
	// FromProto unmarshals FlowRSVPPathObjectsClassRecordRoute from protobuf object *otg.FlowRSVPPathObjectsClassRecordRoute
	FromProto(msg *otg.FlowRSVPPathObjectsClassRecordRoute) (FlowRSVPPathObjectsClassRecordRoute, error)
	// FromPbText unmarshals FlowRSVPPathObjectsClassRecordRoute from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathObjectsClassRecordRoute from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathObjectsClassRecordRoute from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathObjectsClassRecordRoute) Marshal() marshalFlowRSVPPathObjectsClassRecordRoute {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathObjectsClassRecordRoute{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathObjectsClassRecordRoute) Unmarshal() unMarshalFlowRSVPPathObjectsClassRecordRoute {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathObjectsClassRecordRoute{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathObjectsClassRecordRoute) ToProto() (*otg.FlowRSVPPathObjectsClassRecordRoute, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathObjectsClassRecordRoute) FromProto(msg *otg.FlowRSVPPathObjectsClassRecordRoute) (FlowRSVPPathObjectsClassRecordRoute, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathObjectsClassRecordRoute) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassRecordRoute) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathObjectsClassRecordRoute) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassRecordRoute) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathObjectsClassRecordRoute) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassRecordRoute) FromJson(value string) error {
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

func (obj *flowRSVPPathObjectsClassRecordRoute) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsClassRecordRoute) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsClassRecordRoute) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathObjectsClassRecordRoute) Clone() (FlowRSVPPathObjectsClassRecordRoute, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathObjectsClassRecordRoute()
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

func (obj *flowRSVPPathObjectsClassRecordRoute) setNil() {
	obj.lengthHolder = nil
	obj.cTypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathObjectsClassRecordRoute is c-Type is specific to a class num.
type FlowRSVPPathObjectsClassRecordRoute interface {
	Validation
	// msg marshals FlowRSVPPathObjectsClassRecordRoute to protobuf object *otg.FlowRSVPPathObjectsClassRecordRoute
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathObjectsClassRecordRoute
	// setMsg unmarshals FlowRSVPPathObjectsClassRecordRoute from protobuf object *otg.FlowRSVPPathObjectsClassRecordRoute
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathObjectsClassRecordRoute) FlowRSVPPathObjectsClassRecordRoute
	// provides marshal interface
	Marshal() marshalFlowRSVPPathObjectsClassRecordRoute
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathObjectsClassRecordRoute
	// validate validates FlowRSVPPathObjectsClassRecordRoute
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathObjectsClassRecordRoute, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Length returns FlowRSVPObjectLength, set in FlowRSVPPathObjectsClassRecordRoute.
	// FlowRSVPObjectLength is description is TBD
	Length() FlowRSVPObjectLength
	// SetLength assigns FlowRSVPObjectLength provided by user to FlowRSVPPathObjectsClassRecordRoute.
	// FlowRSVPObjectLength is description is TBD
	SetLength(value FlowRSVPObjectLength) FlowRSVPPathObjectsClassRecordRoute
	// HasLength checks if Length has been set in FlowRSVPPathObjectsClassRecordRoute
	HasLength() bool
	// CType returns FlowRSVPPathObjectsRecordRouteCType, set in FlowRSVPPathObjectsClassRecordRoute.
	// FlowRSVPPathObjectsRecordRouteCType is object for RECORD_ROUTE class. c-type is Type 1 Route Record (1).
	CType() FlowRSVPPathObjectsRecordRouteCType
	// SetCType assigns FlowRSVPPathObjectsRecordRouteCType provided by user to FlowRSVPPathObjectsClassRecordRoute.
	// FlowRSVPPathObjectsRecordRouteCType is object for RECORD_ROUTE class. c-type is Type 1 Route Record (1).
	SetCType(value FlowRSVPPathObjectsRecordRouteCType) FlowRSVPPathObjectsClassRecordRoute
	// HasCType checks if CType has been set in FlowRSVPPathObjectsClassRecordRoute
	HasCType() bool
	setNil()
}

// A 16-bit field containing the total object length in bytes.  Must always be a multiple of 4 or at least 4.
// Length returns a FlowRSVPObjectLength
func (obj *flowRSVPPathObjectsClassRecordRoute) Length() FlowRSVPObjectLength {
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
func (obj *flowRSVPPathObjectsClassRecordRoute) HasLength() bool {
	return obj.obj.Length != nil
}

// A 16-bit field containing the total object length in bytes.  Must always be a multiple of 4 or at least 4.
// SetLength sets the FlowRSVPObjectLength value in the FlowRSVPPathObjectsClassRecordRoute object
func (obj *flowRSVPPathObjectsClassRecordRoute) SetLength(value FlowRSVPObjectLength) FlowRSVPPathObjectsClassRecordRoute {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// description is TBD
// CType returns a FlowRSVPPathObjectsRecordRouteCType
func (obj *flowRSVPPathObjectsClassRecordRoute) CType() FlowRSVPPathObjectsRecordRouteCType {
	if obj.obj.CType == nil {
		obj.obj.CType = NewFlowRSVPPathObjectsRecordRouteCType().msg()
	}
	if obj.cTypeHolder == nil {
		obj.cTypeHolder = &flowRSVPPathObjectsRecordRouteCType{obj: obj.obj.CType}
	}
	return obj.cTypeHolder
}

// description is TBD
// CType returns a FlowRSVPPathObjectsRecordRouteCType
func (obj *flowRSVPPathObjectsClassRecordRoute) HasCType() bool {
	return obj.obj.CType != nil
}

// description is TBD
// SetCType sets the FlowRSVPPathObjectsRecordRouteCType value in the FlowRSVPPathObjectsClassRecordRoute object
func (obj *flowRSVPPathObjectsClassRecordRoute) SetCType(value FlowRSVPPathObjectsRecordRouteCType) FlowRSVPPathObjectsClassRecordRoute {

	obj.cTypeHolder = nil
	obj.obj.CType = value.msg()

	return obj
}

func (obj *flowRSVPPathObjectsClassRecordRoute) validateObj(vObj *validation, set_default bool) {
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

func (obj *flowRSVPPathObjectsClassRecordRoute) setDefault() {

}
