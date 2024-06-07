package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathObjectsClassRsvpHop *****
type flowRSVPPathObjectsClassRsvpHop struct {
	validation
	obj          *otg.FlowRSVPPathObjectsClassRsvpHop
	marshaller   marshalFlowRSVPPathObjectsClassRsvpHop
	unMarshaller unMarshalFlowRSVPPathObjectsClassRsvpHop
	lengthHolder FlowRSVPObjectLength
	cTypeHolder  FlowRSVPPathObjectsRsvpHopCType
}

func NewFlowRSVPPathObjectsClassRsvpHop() FlowRSVPPathObjectsClassRsvpHop {
	obj := flowRSVPPathObjectsClassRsvpHop{obj: &otg.FlowRSVPPathObjectsClassRsvpHop{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathObjectsClassRsvpHop) msg() *otg.FlowRSVPPathObjectsClassRsvpHop {
	return obj.obj
}

func (obj *flowRSVPPathObjectsClassRsvpHop) setMsg(msg *otg.FlowRSVPPathObjectsClassRsvpHop) FlowRSVPPathObjectsClassRsvpHop {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathObjectsClassRsvpHop struct {
	obj *flowRSVPPathObjectsClassRsvpHop
}

type marshalFlowRSVPPathObjectsClassRsvpHop interface {
	// ToProto marshals FlowRSVPPathObjectsClassRsvpHop to protobuf object *otg.FlowRSVPPathObjectsClassRsvpHop
	ToProto() (*otg.FlowRSVPPathObjectsClassRsvpHop, error)
	// ToPbText marshals FlowRSVPPathObjectsClassRsvpHop to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathObjectsClassRsvpHop to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathObjectsClassRsvpHop to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPPathObjectsClassRsvpHop struct {
	obj *flowRSVPPathObjectsClassRsvpHop
}

type unMarshalFlowRSVPPathObjectsClassRsvpHop interface {
	// FromProto unmarshals FlowRSVPPathObjectsClassRsvpHop from protobuf object *otg.FlowRSVPPathObjectsClassRsvpHop
	FromProto(msg *otg.FlowRSVPPathObjectsClassRsvpHop) (FlowRSVPPathObjectsClassRsvpHop, error)
	// FromPbText unmarshals FlowRSVPPathObjectsClassRsvpHop from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathObjectsClassRsvpHop from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathObjectsClassRsvpHop from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathObjectsClassRsvpHop) Marshal() marshalFlowRSVPPathObjectsClassRsvpHop {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathObjectsClassRsvpHop{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathObjectsClassRsvpHop) Unmarshal() unMarshalFlowRSVPPathObjectsClassRsvpHop {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathObjectsClassRsvpHop{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathObjectsClassRsvpHop) ToProto() (*otg.FlowRSVPPathObjectsClassRsvpHop, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathObjectsClassRsvpHop) FromProto(msg *otg.FlowRSVPPathObjectsClassRsvpHop) (FlowRSVPPathObjectsClassRsvpHop, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathObjectsClassRsvpHop) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassRsvpHop) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathObjectsClassRsvpHop) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassRsvpHop) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathObjectsClassRsvpHop) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsClassRsvpHop) FromJson(value string) error {
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

func (obj *flowRSVPPathObjectsClassRsvpHop) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsClassRsvpHop) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsClassRsvpHop) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathObjectsClassRsvpHop) Clone() (FlowRSVPPathObjectsClassRsvpHop, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathObjectsClassRsvpHop()
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

func (obj *flowRSVPPathObjectsClassRsvpHop) setNil() {
	obj.lengthHolder = nil
	obj.cTypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathObjectsClassRsvpHop is c-Type is specific to a class num.
type FlowRSVPPathObjectsClassRsvpHop interface {
	Validation
	// msg marshals FlowRSVPPathObjectsClassRsvpHop to protobuf object *otg.FlowRSVPPathObjectsClassRsvpHop
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathObjectsClassRsvpHop
	// setMsg unmarshals FlowRSVPPathObjectsClassRsvpHop from protobuf object *otg.FlowRSVPPathObjectsClassRsvpHop
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathObjectsClassRsvpHop) FlowRSVPPathObjectsClassRsvpHop
	// provides marshal interface
	Marshal() marshalFlowRSVPPathObjectsClassRsvpHop
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathObjectsClassRsvpHop
	// validate validates FlowRSVPPathObjectsClassRsvpHop
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathObjectsClassRsvpHop, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Length returns FlowRSVPObjectLength, set in FlowRSVPPathObjectsClassRsvpHop.
	// FlowRSVPObjectLength is description is TBD
	Length() FlowRSVPObjectLength
	// SetLength assigns FlowRSVPObjectLength provided by user to FlowRSVPPathObjectsClassRsvpHop.
	// FlowRSVPObjectLength is description is TBD
	SetLength(value FlowRSVPObjectLength) FlowRSVPPathObjectsClassRsvpHop
	// HasLength checks if Length has been set in FlowRSVPPathObjectsClassRsvpHop
	HasLength() bool
	// CType returns FlowRSVPPathObjectsRsvpHopCType, set in FlowRSVPPathObjectsClassRsvpHop.
	// FlowRSVPPathObjectsRsvpHopCType is object for RSVP_HOP class. Currently supported c-type is IPv4 (1).
	CType() FlowRSVPPathObjectsRsvpHopCType
	// SetCType assigns FlowRSVPPathObjectsRsvpHopCType provided by user to FlowRSVPPathObjectsClassRsvpHop.
	// FlowRSVPPathObjectsRsvpHopCType is object for RSVP_HOP class. Currently supported c-type is IPv4 (1).
	SetCType(value FlowRSVPPathObjectsRsvpHopCType) FlowRSVPPathObjectsClassRsvpHop
	// HasCType checks if CType has been set in FlowRSVPPathObjectsClassRsvpHop
	HasCType() bool
	setNil()
}

// A 16-bit field containing the total object length in bytes.  Must always be a multiple of 4 or at least 4.
// Length returns a FlowRSVPObjectLength
func (obj *flowRSVPPathObjectsClassRsvpHop) Length() FlowRSVPObjectLength {
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
func (obj *flowRSVPPathObjectsClassRsvpHop) HasLength() bool {
	return obj.obj.Length != nil
}

// A 16-bit field containing the total object length in bytes.  Must always be a multiple of 4 or at least 4.
// SetLength sets the FlowRSVPObjectLength value in the FlowRSVPPathObjectsClassRsvpHop object
func (obj *flowRSVPPathObjectsClassRsvpHop) SetLength(value FlowRSVPObjectLength) FlowRSVPPathObjectsClassRsvpHop {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// description is TBD
// CType returns a FlowRSVPPathObjectsRsvpHopCType
func (obj *flowRSVPPathObjectsClassRsvpHop) CType() FlowRSVPPathObjectsRsvpHopCType {
	if obj.obj.CType == nil {
		obj.obj.CType = NewFlowRSVPPathObjectsRsvpHopCType().msg()
	}
	if obj.cTypeHolder == nil {
		obj.cTypeHolder = &flowRSVPPathObjectsRsvpHopCType{obj: obj.obj.CType}
	}
	return obj.cTypeHolder
}

// description is TBD
// CType returns a FlowRSVPPathObjectsRsvpHopCType
func (obj *flowRSVPPathObjectsClassRsvpHop) HasCType() bool {
	return obj.obj.CType != nil
}

// description is TBD
// SetCType sets the FlowRSVPPathObjectsRsvpHopCType value in the FlowRSVPPathObjectsClassRsvpHop object
func (obj *flowRSVPPathObjectsClassRsvpHop) SetCType(value FlowRSVPPathObjectsRsvpHopCType) FlowRSVPPathObjectsClassRsvpHop {

	obj.cTypeHolder = nil
	obj.obj.CType = value.msg()

	return obj
}

func (obj *flowRSVPPathObjectsClassRsvpHop) validateObj(vObj *validation, set_default bool) {
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

func (obj *flowRSVPPathObjectsClassRsvpHop) setDefault() {

}
