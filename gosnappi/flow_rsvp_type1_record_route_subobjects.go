package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPType1RecordRouteSubobjects *****
type flowRSVPType1RecordRouteSubobjects struct {
	validation
	obj          *otg.FlowRSVPType1RecordRouteSubobjects
	marshaller   marshalFlowRSVPType1RecordRouteSubobjects
	unMarshaller unMarshalFlowRSVPType1RecordRouteSubobjects
	typeHolder   FlowRSVPPathObjectsRecordRouteSubObjectType
}

func NewFlowRSVPType1RecordRouteSubobjects() FlowRSVPType1RecordRouteSubobjects {
	obj := flowRSVPType1RecordRouteSubobjects{obj: &otg.FlowRSVPType1RecordRouteSubobjects{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPType1RecordRouteSubobjects) msg() *otg.FlowRSVPType1RecordRouteSubobjects {
	return obj.obj
}

func (obj *flowRSVPType1RecordRouteSubobjects) setMsg(msg *otg.FlowRSVPType1RecordRouteSubobjects) FlowRSVPType1RecordRouteSubobjects {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPType1RecordRouteSubobjects struct {
	obj *flowRSVPType1RecordRouteSubobjects
}

type marshalFlowRSVPType1RecordRouteSubobjects interface {
	// ToProto marshals FlowRSVPType1RecordRouteSubobjects to protobuf object *otg.FlowRSVPType1RecordRouteSubobjects
	ToProto() (*otg.FlowRSVPType1RecordRouteSubobjects, error)
	// ToPbText marshals FlowRSVPType1RecordRouteSubobjects to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPType1RecordRouteSubobjects to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPType1RecordRouteSubobjects to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowRSVPType1RecordRouteSubobjects to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowRSVPType1RecordRouteSubobjects struct {
	obj *flowRSVPType1RecordRouteSubobjects
}

type unMarshalFlowRSVPType1RecordRouteSubobjects interface {
	// FromProto unmarshals FlowRSVPType1RecordRouteSubobjects from protobuf object *otg.FlowRSVPType1RecordRouteSubobjects
	FromProto(msg *otg.FlowRSVPType1RecordRouteSubobjects) (FlowRSVPType1RecordRouteSubobjects, error)
	// FromPbText unmarshals FlowRSVPType1RecordRouteSubobjects from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPType1RecordRouteSubobjects from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPType1RecordRouteSubobjects from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPType1RecordRouteSubobjects) Marshal() marshalFlowRSVPType1RecordRouteSubobjects {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPType1RecordRouteSubobjects{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPType1RecordRouteSubobjects) Unmarshal() unMarshalFlowRSVPType1RecordRouteSubobjects {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPType1RecordRouteSubobjects{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPType1RecordRouteSubobjects) ToProto() (*otg.FlowRSVPType1RecordRouteSubobjects, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPType1RecordRouteSubobjects) FromProto(msg *otg.FlowRSVPType1RecordRouteSubobjects) (FlowRSVPType1RecordRouteSubobjects, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPType1RecordRouteSubobjects) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPType1RecordRouteSubobjects) FromPbText(value string) error {
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

func (m *marshalflowRSVPType1RecordRouteSubobjects) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPType1RecordRouteSubobjects) FromYaml(value string) error {
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

func (m *marshalflowRSVPType1RecordRouteSubobjects) ToJsonRaw() (string, error) {
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

func (m *marshalflowRSVPType1RecordRouteSubobjects) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPType1RecordRouteSubobjects) FromJson(value string) error {
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

func (obj *flowRSVPType1RecordRouteSubobjects) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPType1RecordRouteSubobjects) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPType1RecordRouteSubobjects) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPType1RecordRouteSubobjects) Clone() (FlowRSVPType1RecordRouteSubobjects, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPType1RecordRouteSubobjects()
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

func (obj *flowRSVPType1RecordRouteSubobjects) setNil() {
	obj.typeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPType1RecordRouteSubobjects is type is specific to a subobject.
type FlowRSVPType1RecordRouteSubobjects interface {
	Validation
	// msg marshals FlowRSVPType1RecordRouteSubobjects to protobuf object *otg.FlowRSVPType1RecordRouteSubobjects
	// and doesn't set defaults
	msg() *otg.FlowRSVPType1RecordRouteSubobjects
	// setMsg unmarshals FlowRSVPType1RecordRouteSubobjects from protobuf object *otg.FlowRSVPType1RecordRouteSubobjects
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPType1RecordRouteSubobjects) FlowRSVPType1RecordRouteSubobjects
	// provides marshal interface
	Marshal() marshalFlowRSVPType1RecordRouteSubobjects
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPType1RecordRouteSubobjects
	// validate validates FlowRSVPType1RecordRouteSubobjects
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPType1RecordRouteSubobjects, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns FlowRSVPPathObjectsRecordRouteSubObjectType, set in FlowRSVPType1RecordRouteSubobjects.
	// FlowRSVPPathObjectsRecordRouteSubObjectType is currently supported subobjects are IPv4 address(1) and Label(3).
	Type() FlowRSVPPathObjectsRecordRouteSubObjectType
	// SetType assigns FlowRSVPPathObjectsRecordRouteSubObjectType provided by user to FlowRSVPType1RecordRouteSubobjects.
	// FlowRSVPPathObjectsRecordRouteSubObjectType is currently supported subobjects are IPv4 address(1) and Label(3).
	SetType(value FlowRSVPPathObjectsRecordRouteSubObjectType) FlowRSVPType1RecordRouteSubobjects
	// HasType checks if Type has been set in FlowRSVPType1RecordRouteSubobjects
	HasType() bool
	setNil()
}

// description is TBD
// Type returns a FlowRSVPPathObjectsRecordRouteSubObjectType
func (obj *flowRSVPType1RecordRouteSubobjects) Type() FlowRSVPPathObjectsRecordRouteSubObjectType {
	if obj.obj.Type == nil {
		obj.obj.Type = NewFlowRSVPPathObjectsRecordRouteSubObjectType().msg()
	}
	if obj.typeHolder == nil {
		obj.typeHolder = &flowRSVPPathObjectsRecordRouteSubObjectType{obj: obj.obj.Type}
	}
	return obj.typeHolder
}

// description is TBD
// Type returns a FlowRSVPPathObjectsRecordRouteSubObjectType
func (obj *flowRSVPType1RecordRouteSubobjects) HasType() bool {
	return obj.obj.Type != nil
}

// description is TBD
// SetType sets the FlowRSVPPathObjectsRecordRouteSubObjectType value in the FlowRSVPType1RecordRouteSubobjects object
func (obj *flowRSVPType1RecordRouteSubobjects) SetType(value FlowRSVPPathObjectsRecordRouteSubObjectType) FlowRSVPType1RecordRouteSubobjects {

	obj.typeHolder = nil
	obj.obj.Type = value.msg()

	return obj
}

func (obj *flowRSVPType1RecordRouteSubobjects) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Type != nil {

		obj.Type().validateObj(vObj, set_default)
	}

}

func (obj *flowRSVPType1RecordRouteSubobjects) setDefault() {

}
