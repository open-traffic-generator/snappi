package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPType1ExplicitRouteSubobjects *****
type flowRSVPType1ExplicitRouteSubobjects struct {
	validation
	obj          *otg.FlowRSVPType1ExplicitRouteSubobjects
	marshaller   marshalFlowRSVPType1ExplicitRouteSubobjects
	unMarshaller unMarshalFlowRSVPType1ExplicitRouteSubobjects
	typeHolder   FlowRSVPType1ExplicitRouteSubobjectsType
}

func NewFlowRSVPType1ExplicitRouteSubobjects() FlowRSVPType1ExplicitRouteSubobjects {
	obj := flowRSVPType1ExplicitRouteSubobjects{obj: &otg.FlowRSVPType1ExplicitRouteSubobjects{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPType1ExplicitRouteSubobjects) msg() *otg.FlowRSVPType1ExplicitRouteSubobjects {
	return obj.obj
}

func (obj *flowRSVPType1ExplicitRouteSubobjects) setMsg(msg *otg.FlowRSVPType1ExplicitRouteSubobjects) FlowRSVPType1ExplicitRouteSubobjects {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPType1ExplicitRouteSubobjects struct {
	obj *flowRSVPType1ExplicitRouteSubobjects
}

type marshalFlowRSVPType1ExplicitRouteSubobjects interface {
	// ToProto marshals FlowRSVPType1ExplicitRouteSubobjects to protobuf object *otg.FlowRSVPType1ExplicitRouteSubobjects
	ToProto() (*otg.FlowRSVPType1ExplicitRouteSubobjects, error)
	// ToPbText marshals FlowRSVPType1ExplicitRouteSubobjects to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPType1ExplicitRouteSubobjects to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPType1ExplicitRouteSubobjects to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPType1ExplicitRouteSubobjects struct {
	obj *flowRSVPType1ExplicitRouteSubobjects
}

type unMarshalFlowRSVPType1ExplicitRouteSubobjects interface {
	// FromProto unmarshals FlowRSVPType1ExplicitRouteSubobjects from protobuf object *otg.FlowRSVPType1ExplicitRouteSubobjects
	FromProto(msg *otg.FlowRSVPType1ExplicitRouteSubobjects) (FlowRSVPType1ExplicitRouteSubobjects, error)
	// FromPbText unmarshals FlowRSVPType1ExplicitRouteSubobjects from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPType1ExplicitRouteSubobjects from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPType1ExplicitRouteSubobjects from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPType1ExplicitRouteSubobjects) Marshal() marshalFlowRSVPType1ExplicitRouteSubobjects {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPType1ExplicitRouteSubobjects{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPType1ExplicitRouteSubobjects) Unmarshal() unMarshalFlowRSVPType1ExplicitRouteSubobjects {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPType1ExplicitRouteSubobjects{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPType1ExplicitRouteSubobjects) ToProto() (*otg.FlowRSVPType1ExplicitRouteSubobjects, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPType1ExplicitRouteSubobjects) FromProto(msg *otg.FlowRSVPType1ExplicitRouteSubobjects) (FlowRSVPType1ExplicitRouteSubobjects, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPType1ExplicitRouteSubobjects) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPType1ExplicitRouteSubobjects) FromPbText(value string) error {
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

func (m *marshalflowRSVPType1ExplicitRouteSubobjects) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPType1ExplicitRouteSubobjects) FromYaml(value string) error {
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

func (m *marshalflowRSVPType1ExplicitRouteSubobjects) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPType1ExplicitRouteSubobjects) FromJson(value string) error {
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

func (obj *flowRSVPType1ExplicitRouteSubobjects) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPType1ExplicitRouteSubobjects) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPType1ExplicitRouteSubobjects) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPType1ExplicitRouteSubobjects) Clone() (FlowRSVPType1ExplicitRouteSubobjects, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPType1ExplicitRouteSubobjects()
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

func (obj *flowRSVPType1ExplicitRouteSubobjects) setNil() {
	obj.typeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPType1ExplicitRouteSubobjects is type is specific to a subobject.
type FlowRSVPType1ExplicitRouteSubobjects interface {
	Validation
	// msg marshals FlowRSVPType1ExplicitRouteSubobjects to protobuf object *otg.FlowRSVPType1ExplicitRouteSubobjects
	// and doesn't set defaults
	msg() *otg.FlowRSVPType1ExplicitRouteSubobjects
	// setMsg unmarshals FlowRSVPType1ExplicitRouteSubobjects from protobuf object *otg.FlowRSVPType1ExplicitRouteSubobjects
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPType1ExplicitRouteSubobjects) FlowRSVPType1ExplicitRouteSubobjects
	// provides marshal interface
	Marshal() marshalFlowRSVPType1ExplicitRouteSubobjects
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPType1ExplicitRouteSubobjects
	// validate validates FlowRSVPType1ExplicitRouteSubobjects
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPType1ExplicitRouteSubobjects, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns FlowRSVPType1ExplicitRouteSubobjectsType, set in FlowRSVPType1ExplicitRouteSubobjects.
	// FlowRSVPType1ExplicitRouteSubobjectsType is currently supported subobjects are IPv4 address(1) and Autonomous system number(32).
	Type() FlowRSVPType1ExplicitRouteSubobjectsType
	// SetType assigns FlowRSVPType1ExplicitRouteSubobjectsType provided by user to FlowRSVPType1ExplicitRouteSubobjects.
	// FlowRSVPType1ExplicitRouteSubobjectsType is currently supported subobjects are IPv4 address(1) and Autonomous system number(32).
	SetType(value FlowRSVPType1ExplicitRouteSubobjectsType) FlowRSVPType1ExplicitRouteSubobjects
	// HasType checks if Type has been set in FlowRSVPType1ExplicitRouteSubobjects
	HasType() bool
	setNil()
}

// description is TBD
// Type returns a FlowRSVPType1ExplicitRouteSubobjectsType
func (obj *flowRSVPType1ExplicitRouteSubobjects) Type() FlowRSVPType1ExplicitRouteSubobjectsType {
	if obj.obj.Type == nil {
		obj.obj.Type = NewFlowRSVPType1ExplicitRouteSubobjectsType().msg()
	}
	if obj.typeHolder == nil {
		obj.typeHolder = &flowRSVPType1ExplicitRouteSubobjectsType{obj: obj.obj.Type}
	}
	return obj.typeHolder
}

// description is TBD
// Type returns a FlowRSVPType1ExplicitRouteSubobjectsType
func (obj *flowRSVPType1ExplicitRouteSubobjects) HasType() bool {
	return obj.obj.Type != nil
}

// description is TBD
// SetType sets the FlowRSVPType1ExplicitRouteSubobjectsType value in the FlowRSVPType1ExplicitRouteSubobjects object
func (obj *flowRSVPType1ExplicitRouteSubobjects) SetType(value FlowRSVPType1ExplicitRouteSubobjectsType) FlowRSVPType1ExplicitRouteSubobjects {

	obj.typeHolder = nil
	obj.obj.Type = value.msg()

	return obj
}

func (obj *flowRSVPType1ExplicitRouteSubobjects) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Type != nil {

		obj.Type().validateObj(vObj, set_default)
	}

}

func (obj *flowRSVPType1ExplicitRouteSubobjects) setDefault() {

}
