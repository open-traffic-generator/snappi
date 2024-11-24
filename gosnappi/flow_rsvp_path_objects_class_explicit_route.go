package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathObjectsClassExplicitRoute *****
type flowRSVPPathObjectsClassExplicitRoute struct {
	validation
	obj          *otg.FlowRSVPPathObjectsClassExplicitRoute
	marshaller   marshalFlowRSVPPathObjectsClassExplicitRoute
	unMarshaller unMarshalFlowRSVPPathObjectsClassExplicitRoute
	lengthHolder FlowRSVPObjectLength
	cTypeHolder  FlowRSVPPathObjectsClassExplicitRouteCType
}

func NewFlowRSVPPathObjectsClassExplicitRoute() FlowRSVPPathObjectsClassExplicitRoute {
	obj := flowRSVPPathObjectsClassExplicitRoute{obj: &otg.FlowRSVPPathObjectsClassExplicitRoute{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathObjectsClassExplicitRoute) msg() *otg.FlowRSVPPathObjectsClassExplicitRoute {
	return obj.obj
}

func (obj *flowRSVPPathObjectsClassExplicitRoute) setMsg(msg *otg.FlowRSVPPathObjectsClassExplicitRoute) FlowRSVPPathObjectsClassExplicitRoute {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathObjectsClassExplicitRoute struct {
	obj *flowRSVPPathObjectsClassExplicitRoute
}

type marshalFlowRSVPPathObjectsClassExplicitRoute interface {
	// ToProto marshals FlowRSVPPathObjectsClassExplicitRoute to protobuf object *otg.FlowRSVPPathObjectsClassExplicitRoute
	ToProto() (*otg.FlowRSVPPathObjectsClassExplicitRoute, error)
	// ToPbText marshals FlowRSVPPathObjectsClassExplicitRoute to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathObjectsClassExplicitRoute to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathObjectsClassExplicitRoute to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPPathObjectsClassExplicitRoute struct {
	obj *flowRSVPPathObjectsClassExplicitRoute
}

type unMarshalFlowRSVPPathObjectsClassExplicitRoute interface {
	// FromProto unmarshals FlowRSVPPathObjectsClassExplicitRoute from protobuf object *otg.FlowRSVPPathObjectsClassExplicitRoute
	FromProto(msg *otg.FlowRSVPPathObjectsClassExplicitRoute) (FlowRSVPPathObjectsClassExplicitRoute, error)
	// FromPbText unmarshals FlowRSVPPathObjectsClassExplicitRoute from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathObjectsClassExplicitRoute from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathObjectsClassExplicitRoute from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathObjectsClassExplicitRoute) Marshal() marshalFlowRSVPPathObjectsClassExplicitRoute {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathObjectsClassExplicitRoute{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathObjectsClassExplicitRoute) Unmarshal() unMarshalFlowRSVPPathObjectsClassExplicitRoute {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathObjectsClassExplicitRoute{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathObjectsClassExplicitRoute) ToProto() (*otg.FlowRSVPPathObjectsClassExplicitRoute, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathObjectsClassExplicitRoute) FromProto(msg *otg.FlowRSVPPathObjectsClassExplicitRoute) (FlowRSVPPathObjectsClassExplicitRoute, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathObjectsClassExplicitRoute) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassExplicitRoute) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathObjectsClassExplicitRoute) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassExplicitRoute) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathObjectsClassExplicitRoute) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassExplicitRoute) FromJson(value string) error {
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

func (obj *flowRSVPPathObjectsClassExplicitRoute) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsClassExplicitRoute) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsClassExplicitRoute) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathObjectsClassExplicitRoute) Clone() (FlowRSVPPathObjectsClassExplicitRoute, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathObjectsClassExplicitRoute()
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

func (obj *flowRSVPPathObjectsClassExplicitRoute) setNil() {
	obj.lengthHolder = nil
	obj.cTypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathObjectsClassExplicitRoute is c-Type is specific to a class num.
type FlowRSVPPathObjectsClassExplicitRoute interface {
	Validation
	// msg marshals FlowRSVPPathObjectsClassExplicitRoute to protobuf object *otg.FlowRSVPPathObjectsClassExplicitRoute
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathObjectsClassExplicitRoute
	// setMsg unmarshals FlowRSVPPathObjectsClassExplicitRoute from protobuf object *otg.FlowRSVPPathObjectsClassExplicitRoute
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathObjectsClassExplicitRoute) FlowRSVPPathObjectsClassExplicitRoute
	// provides marshal interface
	Marshal() marshalFlowRSVPPathObjectsClassExplicitRoute
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathObjectsClassExplicitRoute
	// validate validates FlowRSVPPathObjectsClassExplicitRoute
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathObjectsClassExplicitRoute, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Length returns FlowRSVPObjectLength, set in FlowRSVPPathObjectsClassExplicitRoute.
	// FlowRSVPObjectLength is description is TBD
	Length() FlowRSVPObjectLength
	// SetLength assigns FlowRSVPObjectLength provided by user to FlowRSVPPathObjectsClassExplicitRoute.
	// FlowRSVPObjectLength is description is TBD
	SetLength(value FlowRSVPObjectLength) FlowRSVPPathObjectsClassExplicitRoute
	// HasLength checks if Length has been set in FlowRSVPPathObjectsClassExplicitRoute
	HasLength() bool
	// CType returns FlowRSVPPathObjectsClassExplicitRouteCType, set in FlowRSVPPathObjectsClassExplicitRoute.
	// FlowRSVPPathObjectsClassExplicitRouteCType is object for EXPLICIT_ROUTE class and c-type is Type 1 Explicit Route (1).
	CType() FlowRSVPPathObjectsClassExplicitRouteCType
	// SetCType assigns FlowRSVPPathObjectsClassExplicitRouteCType provided by user to FlowRSVPPathObjectsClassExplicitRoute.
	// FlowRSVPPathObjectsClassExplicitRouteCType is object for EXPLICIT_ROUTE class and c-type is Type 1 Explicit Route (1).
	SetCType(value FlowRSVPPathObjectsClassExplicitRouteCType) FlowRSVPPathObjectsClassExplicitRoute
	// HasCType checks if CType has been set in FlowRSVPPathObjectsClassExplicitRoute
	HasCType() bool
	setNil()
}

// A 16-bit field containing the total object length in bytes.  Must always be a multiple of 4 or at least 4.
// Length returns a FlowRSVPObjectLength
func (obj *flowRSVPPathObjectsClassExplicitRoute) Length() FlowRSVPObjectLength {
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
func (obj *flowRSVPPathObjectsClassExplicitRoute) HasLength() bool {
	return obj.obj.Length != nil
}

// A 16-bit field containing the total object length in bytes.  Must always be a multiple of 4 or at least 4.
// SetLength sets the FlowRSVPObjectLength value in the FlowRSVPPathObjectsClassExplicitRoute object
func (obj *flowRSVPPathObjectsClassExplicitRoute) SetLength(value FlowRSVPObjectLength) FlowRSVPPathObjectsClassExplicitRoute {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// description is TBD
// CType returns a FlowRSVPPathObjectsClassExplicitRouteCType
func (obj *flowRSVPPathObjectsClassExplicitRoute) CType() FlowRSVPPathObjectsClassExplicitRouteCType {
	if obj.obj.CType == nil {
		obj.obj.CType = NewFlowRSVPPathObjectsClassExplicitRouteCType().msg()
	}
	if obj.cTypeHolder == nil {
		obj.cTypeHolder = &flowRSVPPathObjectsClassExplicitRouteCType{obj: obj.obj.CType}
	}
	return obj.cTypeHolder
}

// description is TBD
// CType returns a FlowRSVPPathObjectsClassExplicitRouteCType
func (obj *flowRSVPPathObjectsClassExplicitRoute) HasCType() bool {
	return obj.obj.CType != nil
}

// description is TBD
// SetCType sets the FlowRSVPPathObjectsClassExplicitRouteCType value in the FlowRSVPPathObjectsClassExplicitRoute object
func (obj *flowRSVPPathObjectsClassExplicitRoute) SetCType(value FlowRSVPPathObjectsClassExplicitRouteCType) FlowRSVPPathObjectsClassExplicitRoute {

	obj.cTypeHolder = nil
	obj.obj.CType = value.msg()

	return obj
}

func (obj *flowRSVPPathObjectsClassExplicitRoute) validateObj(vObj *validation, set_default bool) {
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

func (obj *flowRSVPPathObjectsClassExplicitRoute) setDefault() {

}
