package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowVlan *****
type flowVlan struct {
	validation
	obj            *otg.FlowVlan
	marshaller     marshalFlowVlan
	unMarshaller   unMarshalFlowVlan
	priorityHolder PatternFlowVlanPriority
	cfiHolder      PatternFlowVlanCfi
	idHolder       PatternFlowVlanId
	tpidHolder     PatternFlowVlanTpid
}

func NewFlowVlan() FlowVlan {
	obj := flowVlan{obj: &otg.FlowVlan{}}
	obj.setDefault()
	return &obj
}

func (obj *flowVlan) msg() *otg.FlowVlan {
	return obj.obj
}

func (obj *flowVlan) setMsg(msg *otg.FlowVlan) FlowVlan {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowVlan struct {
	obj *flowVlan
}

type marshalFlowVlan interface {
	// ToProto marshals FlowVlan to protobuf object *otg.FlowVlan
	ToProto() (*otg.FlowVlan, error)
	// ToPbText marshals FlowVlan to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowVlan to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowVlan to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowVlan to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowVlan struct {
	obj *flowVlan
}

type unMarshalFlowVlan interface {
	// FromProto unmarshals FlowVlan from protobuf object *otg.FlowVlan
	FromProto(msg *otg.FlowVlan) (FlowVlan, error)
	// FromPbText unmarshals FlowVlan from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowVlan from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowVlan from JSON text
	FromJson(value string) error
}

func (obj *flowVlan) Marshal() marshalFlowVlan {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowVlan{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowVlan) Unmarshal() unMarshalFlowVlan {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowVlan{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowVlan) ToProto() (*otg.FlowVlan, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowVlan) FromProto(msg *otg.FlowVlan) (FlowVlan, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowVlan) ToPbText() (string, error) {
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

func (m *unMarshalflowVlan) FromPbText(value string) error {
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

func (m *marshalflowVlan) ToYaml() (string, error) {
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

func (m *unMarshalflowVlan) FromYaml(value string) error {
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

func (m *marshalflowVlan) ToJsonRaw() (string, error) {
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

func (m *marshalflowVlan) ToJson() (string, error) {
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

func (m *unMarshalflowVlan) FromJson(value string) error {
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

func (obj *flowVlan) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowVlan) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowVlan) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowVlan) Clone() (FlowVlan, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowVlan()
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

func (obj *flowVlan) setNil() {
	obj.priorityHolder = nil
	obj.cfiHolder = nil
	obj.idHolder = nil
	obj.tpidHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowVlan is vLAN packet header
type FlowVlan interface {
	Validation
	// msg marshals FlowVlan to protobuf object *otg.FlowVlan
	// and doesn't set defaults
	msg() *otg.FlowVlan
	// setMsg unmarshals FlowVlan from protobuf object *otg.FlowVlan
	// and doesn't set defaults
	setMsg(*otg.FlowVlan) FlowVlan
	// provides marshal interface
	Marshal() marshalFlowVlan
	// provides unmarshal interface
	Unmarshal() unMarshalFlowVlan
	// validate validates FlowVlan
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowVlan, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Priority returns PatternFlowVlanPriority, set in FlowVlan.
	// PatternFlowVlanPriority is priority code point
	Priority() PatternFlowVlanPriority
	// SetPriority assigns PatternFlowVlanPriority provided by user to FlowVlan.
	// PatternFlowVlanPriority is priority code point
	SetPriority(value PatternFlowVlanPriority) FlowVlan
	// HasPriority checks if Priority has been set in FlowVlan
	HasPriority() bool
	// Cfi returns PatternFlowVlanCfi, set in FlowVlan.
	// PatternFlowVlanCfi is canonical format indicator or drop elegible indicator
	Cfi() PatternFlowVlanCfi
	// SetCfi assigns PatternFlowVlanCfi provided by user to FlowVlan.
	// PatternFlowVlanCfi is canonical format indicator or drop elegible indicator
	SetCfi(value PatternFlowVlanCfi) FlowVlan
	// HasCfi checks if Cfi has been set in FlowVlan
	HasCfi() bool
	// Id returns PatternFlowVlanId, set in FlowVlan.
	// PatternFlowVlanId is vlan identifier
	Id() PatternFlowVlanId
	// SetId assigns PatternFlowVlanId provided by user to FlowVlan.
	// PatternFlowVlanId is vlan identifier
	SetId(value PatternFlowVlanId) FlowVlan
	// HasId checks if Id has been set in FlowVlan
	HasId() bool
	// Tpid returns PatternFlowVlanTpid, set in FlowVlan.
	// PatternFlowVlanTpid is protocol identifier
	Tpid() PatternFlowVlanTpid
	// SetTpid assigns PatternFlowVlanTpid provided by user to FlowVlan.
	// PatternFlowVlanTpid is protocol identifier
	SetTpid(value PatternFlowVlanTpid) FlowVlan
	// HasTpid checks if Tpid has been set in FlowVlan
	HasTpid() bool
	setNil()
}

// description is TBD
// Priority returns a PatternFlowVlanPriority
func (obj *flowVlan) Priority() PatternFlowVlanPriority {
	if obj.obj.Priority == nil {
		obj.obj.Priority = NewPatternFlowVlanPriority().msg()
	}
	if obj.priorityHolder == nil {
		obj.priorityHolder = &patternFlowVlanPriority{obj: obj.obj.Priority}
	}
	return obj.priorityHolder
}

// description is TBD
// Priority returns a PatternFlowVlanPriority
func (obj *flowVlan) HasPriority() bool {
	return obj.obj.Priority != nil
}

// description is TBD
// SetPriority sets the PatternFlowVlanPriority value in the FlowVlan object
func (obj *flowVlan) SetPriority(value PatternFlowVlanPriority) FlowVlan {

	obj.priorityHolder = nil
	obj.obj.Priority = value.msg()

	return obj
}

// description is TBD
// Cfi returns a PatternFlowVlanCfi
func (obj *flowVlan) Cfi() PatternFlowVlanCfi {
	if obj.obj.Cfi == nil {
		obj.obj.Cfi = NewPatternFlowVlanCfi().msg()
	}
	if obj.cfiHolder == nil {
		obj.cfiHolder = &patternFlowVlanCfi{obj: obj.obj.Cfi}
	}
	return obj.cfiHolder
}

// description is TBD
// Cfi returns a PatternFlowVlanCfi
func (obj *flowVlan) HasCfi() bool {
	return obj.obj.Cfi != nil
}

// description is TBD
// SetCfi sets the PatternFlowVlanCfi value in the FlowVlan object
func (obj *flowVlan) SetCfi(value PatternFlowVlanCfi) FlowVlan {

	obj.cfiHolder = nil
	obj.obj.Cfi = value.msg()

	return obj
}

// description is TBD
// Id returns a PatternFlowVlanId
func (obj *flowVlan) Id() PatternFlowVlanId {
	if obj.obj.Id == nil {
		obj.obj.Id = NewPatternFlowVlanId().msg()
	}
	if obj.idHolder == nil {
		obj.idHolder = &patternFlowVlanId{obj: obj.obj.Id}
	}
	return obj.idHolder
}

// description is TBD
// Id returns a PatternFlowVlanId
func (obj *flowVlan) HasId() bool {
	return obj.obj.Id != nil
}

// description is TBD
// SetId sets the PatternFlowVlanId value in the FlowVlan object
func (obj *flowVlan) SetId(value PatternFlowVlanId) FlowVlan {

	obj.idHolder = nil
	obj.obj.Id = value.msg()

	return obj
}

// description is TBD
// Tpid returns a PatternFlowVlanTpid
func (obj *flowVlan) Tpid() PatternFlowVlanTpid {
	if obj.obj.Tpid == nil {
		obj.obj.Tpid = NewPatternFlowVlanTpid().msg()
	}
	if obj.tpidHolder == nil {
		obj.tpidHolder = &patternFlowVlanTpid{obj: obj.obj.Tpid}
	}
	return obj.tpidHolder
}

// description is TBD
// Tpid returns a PatternFlowVlanTpid
func (obj *flowVlan) HasTpid() bool {
	return obj.obj.Tpid != nil
}

// description is TBD
// SetTpid sets the PatternFlowVlanTpid value in the FlowVlan object
func (obj *flowVlan) SetTpid(value PatternFlowVlanTpid) FlowVlan {

	obj.tpidHolder = nil
	obj.obj.Tpid = value.msg()

	return obj
}

func (obj *flowVlan) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Priority != nil {

		obj.Priority().validateObj(vObj, set_default)
	}

	if obj.obj.Cfi != nil {

		obj.Cfi().validateObj(vObj, set_default)
	}

	if obj.obj.Id != nil {

		obj.Id().validateObj(vObj, set_default)
	}

	if obj.obj.Tpid != nil {

		obj.Tpid().validateObj(vObj, set_default)
	}

}

func (obj *flowVlan) setDefault() {

}
