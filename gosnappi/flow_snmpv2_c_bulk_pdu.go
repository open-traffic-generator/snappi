package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowSnmpv2CBulkPDU *****
type flowSnmpv2CBulkPDU struct {
	validation
	obj                    *otg.FlowSnmpv2CBulkPDU
	marshaller             marshalFlowSnmpv2CBulkPDU
	unMarshaller           unMarshalFlowSnmpv2CBulkPDU
	requestIdHolder        PatternFlowSnmpv2CBulkPDURequestId
	nonRepeatersHolder     PatternFlowSnmpv2CBulkPDUNonRepeaters
	maxRepetitionsHolder   PatternFlowSnmpv2CBulkPDUMaxRepetitions
	variableBindingsHolder FlowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter
}

func NewFlowSnmpv2CBulkPDU() FlowSnmpv2CBulkPDU {
	obj := flowSnmpv2CBulkPDU{obj: &otg.FlowSnmpv2CBulkPDU{}}
	obj.setDefault()
	return &obj
}

func (obj *flowSnmpv2CBulkPDU) msg() *otg.FlowSnmpv2CBulkPDU {
	return obj.obj
}

func (obj *flowSnmpv2CBulkPDU) setMsg(msg *otg.FlowSnmpv2CBulkPDU) FlowSnmpv2CBulkPDU {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowSnmpv2CBulkPDU struct {
	obj *flowSnmpv2CBulkPDU
}

type marshalFlowSnmpv2CBulkPDU interface {
	// ToProto marshals FlowSnmpv2CBulkPDU to protobuf object *otg.FlowSnmpv2CBulkPDU
	ToProto() (*otg.FlowSnmpv2CBulkPDU, error)
	// ToPbText marshals FlowSnmpv2CBulkPDU to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowSnmpv2CBulkPDU to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowSnmpv2CBulkPDU to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowSnmpv2CBulkPDU to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowSnmpv2CBulkPDU struct {
	obj *flowSnmpv2CBulkPDU
}

type unMarshalFlowSnmpv2CBulkPDU interface {
	// FromProto unmarshals FlowSnmpv2CBulkPDU from protobuf object *otg.FlowSnmpv2CBulkPDU
	FromProto(msg *otg.FlowSnmpv2CBulkPDU) (FlowSnmpv2CBulkPDU, error)
	// FromPbText unmarshals FlowSnmpv2CBulkPDU from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowSnmpv2CBulkPDU from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowSnmpv2CBulkPDU from JSON text
	FromJson(value string) error
}

func (obj *flowSnmpv2CBulkPDU) Marshal() marshalFlowSnmpv2CBulkPDU {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowSnmpv2CBulkPDU{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowSnmpv2CBulkPDU) Unmarshal() unMarshalFlowSnmpv2CBulkPDU {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowSnmpv2CBulkPDU{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowSnmpv2CBulkPDU) ToProto() (*otg.FlowSnmpv2CBulkPDU, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowSnmpv2CBulkPDU) FromProto(msg *otg.FlowSnmpv2CBulkPDU) (FlowSnmpv2CBulkPDU, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowSnmpv2CBulkPDU) ToPbText() (string, error) {
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

func (m *unMarshalflowSnmpv2CBulkPDU) FromPbText(value string) error {
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

func (m *marshalflowSnmpv2CBulkPDU) ToYaml() (string, error) {
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

func (m *unMarshalflowSnmpv2CBulkPDU) FromYaml(value string) error {
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

func (m *marshalflowSnmpv2CBulkPDU) ToJsonRaw() (string, error) {
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

func (m *marshalflowSnmpv2CBulkPDU) ToJson() (string, error) {
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

func (m *unMarshalflowSnmpv2CBulkPDU) FromJson(value string) error {
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

func (obj *flowSnmpv2CBulkPDU) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowSnmpv2CBulkPDU) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowSnmpv2CBulkPDU) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowSnmpv2CBulkPDU) Clone() (FlowSnmpv2CBulkPDU, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowSnmpv2CBulkPDU()
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

func (obj *flowSnmpv2CBulkPDU) setNil() {
	obj.requestIdHolder = nil
	obj.nonRepeatersHolder = nil
	obj.maxRepetitionsHolder = nil
	obj.variableBindingsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowSnmpv2CBulkPDU is the purpose of the GetBulkRequest-PDU is to request the transfer of a potentially large amount of data, including, but not limited to, the efficient and rapid retrieval of large tables.
type FlowSnmpv2CBulkPDU interface {
	Validation
	// msg marshals FlowSnmpv2CBulkPDU to protobuf object *otg.FlowSnmpv2CBulkPDU
	// and doesn't set defaults
	msg() *otg.FlowSnmpv2CBulkPDU
	// setMsg unmarshals FlowSnmpv2CBulkPDU from protobuf object *otg.FlowSnmpv2CBulkPDU
	// and doesn't set defaults
	setMsg(*otg.FlowSnmpv2CBulkPDU) FlowSnmpv2CBulkPDU
	// provides marshal interface
	Marshal() marshalFlowSnmpv2CBulkPDU
	// provides unmarshal interface
	Unmarshal() unMarshalFlowSnmpv2CBulkPDU
	// validate validates FlowSnmpv2CBulkPDU
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowSnmpv2CBulkPDU, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RequestId returns PatternFlowSnmpv2CBulkPDURequestId, set in FlowSnmpv2CBulkPDU.
	RequestId() PatternFlowSnmpv2CBulkPDURequestId
	// SetRequestId assigns PatternFlowSnmpv2CBulkPDURequestId provided by user to FlowSnmpv2CBulkPDU.
	SetRequestId(value PatternFlowSnmpv2CBulkPDURequestId) FlowSnmpv2CBulkPDU
	// HasRequestId checks if RequestId has been set in FlowSnmpv2CBulkPDU
	HasRequestId() bool
	// NonRepeaters returns PatternFlowSnmpv2CBulkPDUNonRepeaters, set in FlowSnmpv2CBulkPDU.
	NonRepeaters() PatternFlowSnmpv2CBulkPDUNonRepeaters
	// SetNonRepeaters assigns PatternFlowSnmpv2CBulkPDUNonRepeaters provided by user to FlowSnmpv2CBulkPDU.
	SetNonRepeaters(value PatternFlowSnmpv2CBulkPDUNonRepeaters) FlowSnmpv2CBulkPDU
	// HasNonRepeaters checks if NonRepeaters has been set in FlowSnmpv2CBulkPDU
	HasNonRepeaters() bool
	// MaxRepetitions returns PatternFlowSnmpv2CBulkPDUMaxRepetitions, set in FlowSnmpv2CBulkPDU.
	MaxRepetitions() PatternFlowSnmpv2CBulkPDUMaxRepetitions
	// SetMaxRepetitions assigns PatternFlowSnmpv2CBulkPDUMaxRepetitions provided by user to FlowSnmpv2CBulkPDU.
	SetMaxRepetitions(value PatternFlowSnmpv2CBulkPDUMaxRepetitions) FlowSnmpv2CBulkPDU
	// HasMaxRepetitions checks if MaxRepetitions has been set in FlowSnmpv2CBulkPDU
	HasMaxRepetitions() bool
	// VariableBindings returns FlowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIterIter, set in FlowSnmpv2CBulkPDU
	VariableBindings() FlowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter
	setNil()
}

// description is TBD
// RequestId returns a PatternFlowSnmpv2CBulkPDURequestId
func (obj *flowSnmpv2CBulkPDU) RequestId() PatternFlowSnmpv2CBulkPDURequestId {
	if obj.obj.RequestId == nil {
		obj.obj.RequestId = NewPatternFlowSnmpv2CBulkPDURequestId().msg()
	}
	if obj.requestIdHolder == nil {
		obj.requestIdHolder = &patternFlowSnmpv2CBulkPDURequestId{obj: obj.obj.RequestId}
	}
	return obj.requestIdHolder
}

// description is TBD
// RequestId returns a PatternFlowSnmpv2CBulkPDURequestId
func (obj *flowSnmpv2CBulkPDU) HasRequestId() bool {
	return obj.obj.RequestId != nil
}

// description is TBD
// SetRequestId sets the PatternFlowSnmpv2CBulkPDURequestId value in the FlowSnmpv2CBulkPDU object
func (obj *flowSnmpv2CBulkPDU) SetRequestId(value PatternFlowSnmpv2CBulkPDURequestId) FlowSnmpv2CBulkPDU {

	obj.requestIdHolder = nil
	obj.obj.RequestId = value.msg()

	return obj
}

// description is TBD
// NonRepeaters returns a PatternFlowSnmpv2CBulkPDUNonRepeaters
func (obj *flowSnmpv2CBulkPDU) NonRepeaters() PatternFlowSnmpv2CBulkPDUNonRepeaters {
	if obj.obj.NonRepeaters == nil {
		obj.obj.NonRepeaters = NewPatternFlowSnmpv2CBulkPDUNonRepeaters().msg()
	}
	if obj.nonRepeatersHolder == nil {
		obj.nonRepeatersHolder = &patternFlowSnmpv2CBulkPDUNonRepeaters{obj: obj.obj.NonRepeaters}
	}
	return obj.nonRepeatersHolder
}

// description is TBD
// NonRepeaters returns a PatternFlowSnmpv2CBulkPDUNonRepeaters
func (obj *flowSnmpv2CBulkPDU) HasNonRepeaters() bool {
	return obj.obj.NonRepeaters != nil
}

// description is TBD
// SetNonRepeaters sets the PatternFlowSnmpv2CBulkPDUNonRepeaters value in the FlowSnmpv2CBulkPDU object
func (obj *flowSnmpv2CBulkPDU) SetNonRepeaters(value PatternFlowSnmpv2CBulkPDUNonRepeaters) FlowSnmpv2CBulkPDU {

	obj.nonRepeatersHolder = nil
	obj.obj.NonRepeaters = value.msg()

	return obj
}

// description is TBD
// MaxRepetitions returns a PatternFlowSnmpv2CBulkPDUMaxRepetitions
func (obj *flowSnmpv2CBulkPDU) MaxRepetitions() PatternFlowSnmpv2CBulkPDUMaxRepetitions {
	if obj.obj.MaxRepetitions == nil {
		obj.obj.MaxRepetitions = NewPatternFlowSnmpv2CBulkPDUMaxRepetitions().msg()
	}
	if obj.maxRepetitionsHolder == nil {
		obj.maxRepetitionsHolder = &patternFlowSnmpv2CBulkPDUMaxRepetitions{obj: obj.obj.MaxRepetitions}
	}
	return obj.maxRepetitionsHolder
}

// description is TBD
// MaxRepetitions returns a PatternFlowSnmpv2CBulkPDUMaxRepetitions
func (obj *flowSnmpv2CBulkPDU) HasMaxRepetitions() bool {
	return obj.obj.MaxRepetitions != nil
}

// description is TBD
// SetMaxRepetitions sets the PatternFlowSnmpv2CBulkPDUMaxRepetitions value in the FlowSnmpv2CBulkPDU object
func (obj *flowSnmpv2CBulkPDU) SetMaxRepetitions(value PatternFlowSnmpv2CBulkPDUMaxRepetitions) FlowSnmpv2CBulkPDU {

	obj.maxRepetitionsHolder = nil
	obj.obj.MaxRepetitions = value.msg()

	return obj
}

// A Sequence of variable_bindings.
// VariableBindings returns a []FlowSnmpv2CVariableBinding
func (obj *flowSnmpv2CBulkPDU) VariableBindings() FlowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter {
	if len(obj.obj.VariableBindings) == 0 {
		obj.obj.VariableBindings = []*otg.FlowSnmpv2CVariableBinding{}
	}
	if obj.variableBindingsHolder == nil {
		obj.variableBindingsHolder = newFlowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter(&obj.obj.VariableBindings).setMsg(obj)
	}
	return obj.variableBindingsHolder
}

type flowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter struct {
	obj                             *flowSnmpv2CBulkPDU
	flowSnmpv2CVariableBindingSlice []FlowSnmpv2CVariableBinding
	fieldPtr                        *[]*otg.FlowSnmpv2CVariableBinding
}

func newFlowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter(ptr *[]*otg.FlowSnmpv2CVariableBinding) FlowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter {
	return &flowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter{fieldPtr: ptr}
}

type FlowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter interface {
	setMsg(*flowSnmpv2CBulkPDU) FlowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter
	Items() []FlowSnmpv2CVariableBinding
	Add() FlowSnmpv2CVariableBinding
	Append(items ...FlowSnmpv2CVariableBinding) FlowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter
	Set(index int, newObj FlowSnmpv2CVariableBinding) FlowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter
	Clear() FlowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter
	clearHolderSlice() FlowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter
	appendHolderSlice(item FlowSnmpv2CVariableBinding) FlowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter
}

func (obj *flowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter) setMsg(msg *flowSnmpv2CBulkPDU) FlowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowSnmpv2CVariableBinding{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter) Items() []FlowSnmpv2CVariableBinding {
	return obj.flowSnmpv2CVariableBindingSlice
}

func (obj *flowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter) Add() FlowSnmpv2CVariableBinding {
	newObj := &otg.FlowSnmpv2CVariableBinding{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowSnmpv2CVariableBinding{obj: newObj}
	newLibObj.setDefault()
	obj.flowSnmpv2CVariableBindingSlice = append(obj.flowSnmpv2CVariableBindingSlice, newLibObj)
	return newLibObj
}

func (obj *flowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter) Append(items ...FlowSnmpv2CVariableBinding) FlowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowSnmpv2CVariableBindingSlice = append(obj.flowSnmpv2CVariableBindingSlice, item)
	}
	return obj
}

func (obj *flowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter) Set(index int, newObj FlowSnmpv2CVariableBinding) FlowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowSnmpv2CVariableBindingSlice[index] = newObj
	return obj
}
func (obj *flowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter) Clear() FlowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowSnmpv2CVariableBinding{}
		obj.flowSnmpv2CVariableBindingSlice = []FlowSnmpv2CVariableBinding{}
	}
	return obj
}
func (obj *flowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter) clearHolderSlice() FlowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter {
	if len(obj.flowSnmpv2CVariableBindingSlice) > 0 {
		obj.flowSnmpv2CVariableBindingSlice = []FlowSnmpv2CVariableBinding{}
	}
	return obj
}
func (obj *flowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter) appendHolderSlice(item FlowSnmpv2CVariableBinding) FlowSnmpv2CBulkPDUFlowSnmpv2CVariableBindingIter {
	obj.flowSnmpv2CVariableBindingSlice = append(obj.flowSnmpv2CVariableBindingSlice, item)
	return obj
}

func (obj *flowSnmpv2CBulkPDU) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RequestId != nil {

		obj.RequestId().validateObj(vObj, set_default)
	}

	if obj.obj.NonRepeaters != nil {

		obj.NonRepeaters().validateObj(vObj, set_default)
	}

	if obj.obj.MaxRepetitions != nil {

		obj.MaxRepetitions().validateObj(vObj, set_default)
	}

	if len(obj.obj.VariableBindings) != 0 {

		if set_default {
			obj.VariableBindings().clearHolderSlice()
			for _, item := range obj.obj.VariableBindings {
				obj.VariableBindings().appendHolderSlice(&flowSnmpv2CVariableBinding{obj: item})
			}
		}
		for _, item := range obj.VariableBindings().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *flowSnmpv2CBulkPDU) setDefault() {

}
