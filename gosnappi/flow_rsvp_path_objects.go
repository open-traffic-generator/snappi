package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathObjects *****
type flowRSVPPathObjects struct {
	validation
	obj            *otg.FlowRSVPPathObjects
	marshaller     marshalFlowRSVPPathObjects
	unMarshaller   unMarshalFlowRSVPPathObjects
	classNumHolder FlowRSVPPathObjectsClass
}

func NewFlowRSVPPathObjects() FlowRSVPPathObjects {
	obj := flowRSVPPathObjects{obj: &otg.FlowRSVPPathObjects{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathObjects) msg() *otg.FlowRSVPPathObjects {
	return obj.obj
}

func (obj *flowRSVPPathObjects) setMsg(msg *otg.FlowRSVPPathObjects) FlowRSVPPathObjects {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathObjects struct {
	obj *flowRSVPPathObjects
}

type marshalFlowRSVPPathObjects interface {
	// ToProto marshals FlowRSVPPathObjects to protobuf object *otg.FlowRSVPPathObjects
	ToProto() (*otg.FlowRSVPPathObjects, error)
	// ToPbText marshals FlowRSVPPathObjects to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathObjects to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathObjects to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPPathObjects struct {
	obj *flowRSVPPathObjects
}

type unMarshalFlowRSVPPathObjects interface {
	// FromProto unmarshals FlowRSVPPathObjects from protobuf object *otg.FlowRSVPPathObjects
	FromProto(msg *otg.FlowRSVPPathObjects) (FlowRSVPPathObjects, error)
	// FromPbText unmarshals FlowRSVPPathObjects from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathObjects from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathObjects from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathObjects) Marshal() marshalFlowRSVPPathObjects {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathObjects{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathObjects) Unmarshal() unMarshalFlowRSVPPathObjects {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathObjects{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathObjects) ToProto() (*otg.FlowRSVPPathObjects, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathObjects) FromProto(msg *otg.FlowRSVPPathObjects) (FlowRSVPPathObjects, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathObjects) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathObjects) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathObjects) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathObjects) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathObjects) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathObjects) FromJson(value string) error {
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

func (obj *flowRSVPPathObjects) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjects) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjects) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathObjects) Clone() (FlowRSVPPathObjects, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathObjects()
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

func (obj *flowRSVPPathObjects) setNil() {
	obj.classNumHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathObjects is every RSVP object encapsulated in an RSVP message consists of a 32-bit word header and the object's contents.
type FlowRSVPPathObjects interface {
	Validation
	// msg marshals FlowRSVPPathObjects to protobuf object *otg.FlowRSVPPathObjects
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathObjects
	// setMsg unmarshals FlowRSVPPathObjects from protobuf object *otg.FlowRSVPPathObjects
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathObjects) FlowRSVPPathObjects
	// provides marshal interface
	Marshal() marshalFlowRSVPPathObjects
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathObjects
	// validate validates FlowRSVPPathObjects
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathObjects, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ClassNum returns FlowRSVPPathObjectsClass, set in FlowRSVPPathObjects.
	// FlowRSVPPathObjectsClass is the class number is used to identify the class of an object. Defined in https://www.iana.org/assignments/rsvp-parameters/rsvp-parameters.xhtml#rsvp-parameters-4 . Curently supported class numbers are for "Path" message type. "Path" message: Supported Class numbers and it's value: SESSION: 1, RSVP_HOP: 3, TIME_VALUES: 5, EXPLICIT_ROUTE: 20, LABEL_REQUEST: 19, SESSION_ATTRIBUTE: 207, SENDER_TEMPLATE: 11, SENDER_TSPEC: 12, RECORD_ROUTE: 21, Custom: User defined bytes based on class and c-types not supported in above options.
	ClassNum() FlowRSVPPathObjectsClass
	// SetClassNum assigns FlowRSVPPathObjectsClass provided by user to FlowRSVPPathObjects.
	// FlowRSVPPathObjectsClass is the class number is used to identify the class of an object. Defined in https://www.iana.org/assignments/rsvp-parameters/rsvp-parameters.xhtml#rsvp-parameters-4 . Curently supported class numbers are for "Path" message type. "Path" message: Supported Class numbers and it's value: SESSION: 1, RSVP_HOP: 3, TIME_VALUES: 5, EXPLICIT_ROUTE: 20, LABEL_REQUEST: 19, SESSION_ATTRIBUTE: 207, SENDER_TEMPLATE: 11, SENDER_TSPEC: 12, RECORD_ROUTE: 21, Custom: User defined bytes based on class and c-types not supported in above options.
	SetClassNum(value FlowRSVPPathObjectsClass) FlowRSVPPathObjects
	// HasClassNum checks if ClassNum has been set in FlowRSVPPathObjects
	HasClassNum() bool
	setNil()
}

// description is TBD
// ClassNum returns a FlowRSVPPathObjectsClass
func (obj *flowRSVPPathObjects) ClassNum() FlowRSVPPathObjectsClass {
	if obj.obj.ClassNum == nil {
		obj.obj.ClassNum = NewFlowRSVPPathObjectsClass().msg()
	}
	if obj.classNumHolder == nil {
		obj.classNumHolder = &flowRSVPPathObjectsClass{obj: obj.obj.ClassNum}
	}
	return obj.classNumHolder
}

// description is TBD
// ClassNum returns a FlowRSVPPathObjectsClass
func (obj *flowRSVPPathObjects) HasClassNum() bool {
	return obj.obj.ClassNum != nil
}

// description is TBD
// SetClassNum sets the FlowRSVPPathObjectsClass value in the FlowRSVPPathObjects object
func (obj *flowRSVPPathObjects) SetClassNum(value FlowRSVPPathObjectsClass) FlowRSVPPathObjects {

	obj.classNumHolder = nil
	obj.obj.ClassNum = value.msg()

	return obj
}

func (obj *flowRSVPPathObjects) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ClassNum != nil {

		obj.ClassNum().validateObj(vObj, set_default)
	}

}

func (obj *flowRSVPPathObjects) setDefault() {

}
