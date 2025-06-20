package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowSnmpv2CPDU *****
type flowSnmpv2CPDU struct {
	validation
	obj                    *otg.FlowSnmpv2CPDU
	marshaller             marshalFlowSnmpv2CPDU
	unMarshaller           unMarshalFlowSnmpv2CPDU
	requestIdHolder        PatternFlowSnmpv2CPDURequestId
	errorIndexHolder       PatternFlowSnmpv2CPDUErrorIndex
	variableBindingsHolder FlowSnmpv2CPDUFlowSnmpv2CVariableBindingIter
}

func NewFlowSnmpv2CPDU() FlowSnmpv2CPDU {
	obj := flowSnmpv2CPDU{obj: &otg.FlowSnmpv2CPDU{}}
	obj.setDefault()
	return &obj
}

func (obj *flowSnmpv2CPDU) msg() *otg.FlowSnmpv2CPDU {
	return obj.obj
}

func (obj *flowSnmpv2CPDU) setMsg(msg *otg.FlowSnmpv2CPDU) FlowSnmpv2CPDU {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowSnmpv2CPDU struct {
	obj *flowSnmpv2CPDU
}

type marshalFlowSnmpv2CPDU interface {
	// ToProto marshals FlowSnmpv2CPDU to protobuf object *otg.FlowSnmpv2CPDU
	ToProto() (*otg.FlowSnmpv2CPDU, error)
	// ToPbText marshals FlowSnmpv2CPDU to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowSnmpv2CPDU to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowSnmpv2CPDU to JSON text
	ToJson() (string, error)
}

type unMarshalflowSnmpv2CPDU struct {
	obj *flowSnmpv2CPDU
}

type unMarshalFlowSnmpv2CPDU interface {
	// FromProto unmarshals FlowSnmpv2CPDU from protobuf object *otg.FlowSnmpv2CPDU
	FromProto(msg *otg.FlowSnmpv2CPDU) (FlowSnmpv2CPDU, error)
	// FromPbText unmarshals FlowSnmpv2CPDU from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowSnmpv2CPDU from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowSnmpv2CPDU from JSON text
	FromJson(value string) error
}

func (obj *flowSnmpv2CPDU) Marshal() marshalFlowSnmpv2CPDU {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowSnmpv2CPDU{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowSnmpv2CPDU) Unmarshal() unMarshalFlowSnmpv2CPDU {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowSnmpv2CPDU{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowSnmpv2CPDU) ToProto() (*otg.FlowSnmpv2CPDU, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowSnmpv2CPDU) FromProto(msg *otg.FlowSnmpv2CPDU) (FlowSnmpv2CPDU, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowSnmpv2CPDU) ToPbText() (string, error) {
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

func (m *unMarshalflowSnmpv2CPDU) FromPbText(value string) error {
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

func (m *marshalflowSnmpv2CPDU) ToYaml() (string, error) {
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

func (m *unMarshalflowSnmpv2CPDU) FromYaml(value string) error {
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

func (m *marshalflowSnmpv2CPDU) ToJson() (string, error) {
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

func (m *unMarshalflowSnmpv2CPDU) FromJson(value string) error {
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

func (obj *flowSnmpv2CPDU) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowSnmpv2CPDU) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowSnmpv2CPDU) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowSnmpv2CPDU) Clone() (FlowSnmpv2CPDU, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowSnmpv2CPDU()
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

func (obj *flowSnmpv2CPDU) setNil() {
	obj.requestIdHolder = nil
	obj.errorIndexHolder = nil
	obj.variableBindingsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowSnmpv2CPDU is this contains the body of the SNMPv2C PDU.
type FlowSnmpv2CPDU interface {
	Validation
	// msg marshals FlowSnmpv2CPDU to protobuf object *otg.FlowSnmpv2CPDU
	// and doesn't set defaults
	msg() *otg.FlowSnmpv2CPDU
	// setMsg unmarshals FlowSnmpv2CPDU from protobuf object *otg.FlowSnmpv2CPDU
	// and doesn't set defaults
	setMsg(*otg.FlowSnmpv2CPDU) FlowSnmpv2CPDU
	// provides marshal interface
	Marshal() marshalFlowSnmpv2CPDU
	// provides unmarshal interface
	Unmarshal() unMarshalFlowSnmpv2CPDU
	// validate validates FlowSnmpv2CPDU
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowSnmpv2CPDU, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RequestId returns PatternFlowSnmpv2CPDURequestId, set in FlowSnmpv2CPDU.
	RequestId() PatternFlowSnmpv2CPDURequestId
	// SetRequestId assigns PatternFlowSnmpv2CPDURequestId provided by user to FlowSnmpv2CPDU.
	SetRequestId(value PatternFlowSnmpv2CPDURequestId) FlowSnmpv2CPDU
	// HasRequestId checks if RequestId has been set in FlowSnmpv2CPDU
	HasRequestId() bool
	// ErrorStatus returns FlowSnmpv2CPDUErrorStatusEnum, set in FlowSnmpv2CPDU
	ErrorStatus() FlowSnmpv2CPDUErrorStatusEnum
	// SetErrorStatus assigns FlowSnmpv2CPDUErrorStatusEnum provided by user to FlowSnmpv2CPDU
	SetErrorStatus(value FlowSnmpv2CPDUErrorStatusEnum) FlowSnmpv2CPDU
	// HasErrorStatus checks if ErrorStatus has been set in FlowSnmpv2CPDU
	HasErrorStatus() bool
	// ErrorIndex returns PatternFlowSnmpv2CPDUErrorIndex, set in FlowSnmpv2CPDU.
	ErrorIndex() PatternFlowSnmpv2CPDUErrorIndex
	// SetErrorIndex assigns PatternFlowSnmpv2CPDUErrorIndex provided by user to FlowSnmpv2CPDU.
	SetErrorIndex(value PatternFlowSnmpv2CPDUErrorIndex) FlowSnmpv2CPDU
	// HasErrorIndex checks if ErrorIndex has been set in FlowSnmpv2CPDU
	HasErrorIndex() bool
	// VariableBindings returns FlowSnmpv2CPDUFlowSnmpv2CVariableBindingIterIter, set in FlowSnmpv2CPDU
	VariableBindings() FlowSnmpv2CPDUFlowSnmpv2CVariableBindingIter
	setNil()
}

// description is TBD
// RequestId returns a PatternFlowSnmpv2CPDURequestId
func (obj *flowSnmpv2CPDU) RequestId() PatternFlowSnmpv2CPDURequestId {
	if obj.obj.RequestId == nil {
		obj.obj.RequestId = NewPatternFlowSnmpv2CPDURequestId().msg()
	}
	if obj.requestIdHolder == nil {
		obj.requestIdHolder = &patternFlowSnmpv2CPDURequestId{obj: obj.obj.RequestId}
	}
	return obj.requestIdHolder
}

// description is TBD
// RequestId returns a PatternFlowSnmpv2CPDURequestId
func (obj *flowSnmpv2CPDU) HasRequestId() bool {
	return obj.obj.RequestId != nil
}

// description is TBD
// SetRequestId sets the PatternFlowSnmpv2CPDURequestId value in the FlowSnmpv2CPDU object
func (obj *flowSnmpv2CPDU) SetRequestId(value PatternFlowSnmpv2CPDURequestId) FlowSnmpv2CPDU {

	obj.requestIdHolder = nil
	obj.obj.RequestId = value.msg()

	return obj
}

type FlowSnmpv2CPDUErrorStatusEnum string

// Enum of ErrorStatus on FlowSnmpv2CPDU
var FlowSnmpv2CPDUErrorStatus = struct {
	NO_ERROR             FlowSnmpv2CPDUErrorStatusEnum
	TOO_BIG              FlowSnmpv2CPDUErrorStatusEnum
	NO_SUCH_NAME         FlowSnmpv2CPDUErrorStatusEnum
	BAD_VALUE            FlowSnmpv2CPDUErrorStatusEnum
	READ_ONLY            FlowSnmpv2CPDUErrorStatusEnum
	GEN_ERR              FlowSnmpv2CPDUErrorStatusEnum
	NO_ACCESS            FlowSnmpv2CPDUErrorStatusEnum
	WRONG_TYPE           FlowSnmpv2CPDUErrorStatusEnum
	WRONG_LENGTH         FlowSnmpv2CPDUErrorStatusEnum
	WRONG_ENCODING       FlowSnmpv2CPDUErrorStatusEnum
	WRONG_VALUE          FlowSnmpv2CPDUErrorStatusEnum
	NO_CREATION          FlowSnmpv2CPDUErrorStatusEnum
	INCONSISTENT_VALUE   FlowSnmpv2CPDUErrorStatusEnum
	RESOURCE_UNAVAILABLE FlowSnmpv2CPDUErrorStatusEnum
	COMMIT_FAILED        FlowSnmpv2CPDUErrorStatusEnum
	UNDO_FAILED          FlowSnmpv2CPDUErrorStatusEnum
	AUTHORIZATION_ERROR  FlowSnmpv2CPDUErrorStatusEnum
	NOT_WRITABLE         FlowSnmpv2CPDUErrorStatusEnum
	INCONSISTENT_NAME    FlowSnmpv2CPDUErrorStatusEnum
}{
	NO_ERROR:             FlowSnmpv2CPDUErrorStatusEnum("no_error"),
	TOO_BIG:              FlowSnmpv2CPDUErrorStatusEnum("too_big"),
	NO_SUCH_NAME:         FlowSnmpv2CPDUErrorStatusEnum("no_such_name"),
	BAD_VALUE:            FlowSnmpv2CPDUErrorStatusEnum("bad_value"),
	READ_ONLY:            FlowSnmpv2CPDUErrorStatusEnum("read_only"),
	GEN_ERR:              FlowSnmpv2CPDUErrorStatusEnum("gen_err"),
	NO_ACCESS:            FlowSnmpv2CPDUErrorStatusEnum("no_access"),
	WRONG_TYPE:           FlowSnmpv2CPDUErrorStatusEnum("wrong_type"),
	WRONG_LENGTH:         FlowSnmpv2CPDUErrorStatusEnum("wrong_length"),
	WRONG_ENCODING:       FlowSnmpv2CPDUErrorStatusEnum("wrong_encoding"),
	WRONG_VALUE:          FlowSnmpv2CPDUErrorStatusEnum("wrong_value"),
	NO_CREATION:          FlowSnmpv2CPDUErrorStatusEnum("no_creation"),
	INCONSISTENT_VALUE:   FlowSnmpv2CPDUErrorStatusEnum("inconsistent_value"),
	RESOURCE_UNAVAILABLE: FlowSnmpv2CPDUErrorStatusEnum("resource_unavailable"),
	COMMIT_FAILED:        FlowSnmpv2CPDUErrorStatusEnum("commit_failed"),
	UNDO_FAILED:          FlowSnmpv2CPDUErrorStatusEnum("undo_failed"),
	AUTHORIZATION_ERROR:  FlowSnmpv2CPDUErrorStatusEnum("authorization_error"),
	NOT_WRITABLE:         FlowSnmpv2CPDUErrorStatusEnum("not_writable"),
	INCONSISTENT_NAME:    FlowSnmpv2CPDUErrorStatusEnum("inconsistent_name"),
}

func (obj *flowSnmpv2CPDU) ErrorStatus() FlowSnmpv2CPDUErrorStatusEnum {
	return FlowSnmpv2CPDUErrorStatusEnum(obj.obj.ErrorStatus.Enum().String())
}

// The SNMP agent places an error code in this field in the response message if an error occurred processing the request.
// ErrorStatus returns a string
func (obj *flowSnmpv2CPDU) HasErrorStatus() bool {
	return obj.obj.ErrorStatus != nil
}

func (obj *flowSnmpv2CPDU) SetErrorStatus(value FlowSnmpv2CPDUErrorStatusEnum) FlowSnmpv2CPDU {
	intValue, ok := otg.FlowSnmpv2CPDU_ErrorStatus_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowSnmpv2CPDUErrorStatusEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowSnmpv2CPDU_ErrorStatus_Enum(intValue)
	obj.obj.ErrorStatus = &enumValue

	return obj
}

// description is TBD
// ErrorIndex returns a PatternFlowSnmpv2CPDUErrorIndex
func (obj *flowSnmpv2CPDU) ErrorIndex() PatternFlowSnmpv2CPDUErrorIndex {
	if obj.obj.ErrorIndex == nil {
		obj.obj.ErrorIndex = NewPatternFlowSnmpv2CPDUErrorIndex().msg()
	}
	if obj.errorIndexHolder == nil {
		obj.errorIndexHolder = &patternFlowSnmpv2CPDUErrorIndex{obj: obj.obj.ErrorIndex}
	}
	return obj.errorIndexHolder
}

// description is TBD
// ErrorIndex returns a PatternFlowSnmpv2CPDUErrorIndex
func (obj *flowSnmpv2CPDU) HasErrorIndex() bool {
	return obj.obj.ErrorIndex != nil
}

// description is TBD
// SetErrorIndex sets the PatternFlowSnmpv2CPDUErrorIndex value in the FlowSnmpv2CPDU object
func (obj *flowSnmpv2CPDU) SetErrorIndex(value PatternFlowSnmpv2CPDUErrorIndex) FlowSnmpv2CPDU {

	obj.errorIndexHolder = nil
	obj.obj.ErrorIndex = value.msg()

	return obj
}

// A Sequence of variable_bindings.
// VariableBindings returns a []FlowSnmpv2CVariableBinding
func (obj *flowSnmpv2CPDU) VariableBindings() FlowSnmpv2CPDUFlowSnmpv2CVariableBindingIter {
	if len(obj.obj.VariableBindings) == 0 {
		obj.obj.VariableBindings = []*otg.FlowSnmpv2CVariableBinding{}
	}
	if obj.variableBindingsHolder == nil {
		obj.variableBindingsHolder = newFlowSnmpv2CPDUFlowSnmpv2CVariableBindingIter(&obj.obj.VariableBindings).setMsg(obj)
	}
	return obj.variableBindingsHolder
}

type flowSnmpv2CPDUFlowSnmpv2CVariableBindingIter struct {
	obj                             *flowSnmpv2CPDU
	flowSnmpv2CVariableBindingSlice []FlowSnmpv2CVariableBinding
	fieldPtr                        *[]*otg.FlowSnmpv2CVariableBinding
}

func newFlowSnmpv2CPDUFlowSnmpv2CVariableBindingIter(ptr *[]*otg.FlowSnmpv2CVariableBinding) FlowSnmpv2CPDUFlowSnmpv2CVariableBindingIter {
	return &flowSnmpv2CPDUFlowSnmpv2CVariableBindingIter{fieldPtr: ptr}
}

type FlowSnmpv2CPDUFlowSnmpv2CVariableBindingIter interface {
	setMsg(*flowSnmpv2CPDU) FlowSnmpv2CPDUFlowSnmpv2CVariableBindingIter
	Items() []FlowSnmpv2CVariableBinding
	Add() FlowSnmpv2CVariableBinding
	Append(items ...FlowSnmpv2CVariableBinding) FlowSnmpv2CPDUFlowSnmpv2CVariableBindingIter
	Set(index int, newObj FlowSnmpv2CVariableBinding) FlowSnmpv2CPDUFlowSnmpv2CVariableBindingIter
	Clear() FlowSnmpv2CPDUFlowSnmpv2CVariableBindingIter
	clearHolderSlice() FlowSnmpv2CPDUFlowSnmpv2CVariableBindingIter
	appendHolderSlice(item FlowSnmpv2CVariableBinding) FlowSnmpv2CPDUFlowSnmpv2CVariableBindingIter
}

func (obj *flowSnmpv2CPDUFlowSnmpv2CVariableBindingIter) setMsg(msg *flowSnmpv2CPDU) FlowSnmpv2CPDUFlowSnmpv2CVariableBindingIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowSnmpv2CVariableBinding{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowSnmpv2CPDUFlowSnmpv2CVariableBindingIter) Items() []FlowSnmpv2CVariableBinding {
	return obj.flowSnmpv2CVariableBindingSlice
}

func (obj *flowSnmpv2CPDUFlowSnmpv2CVariableBindingIter) Add() FlowSnmpv2CVariableBinding {
	newObj := &otg.FlowSnmpv2CVariableBinding{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowSnmpv2CVariableBinding{obj: newObj}
	newLibObj.setDefault()
	obj.flowSnmpv2CVariableBindingSlice = append(obj.flowSnmpv2CVariableBindingSlice, newLibObj)
	return newLibObj
}

func (obj *flowSnmpv2CPDUFlowSnmpv2CVariableBindingIter) Append(items ...FlowSnmpv2CVariableBinding) FlowSnmpv2CPDUFlowSnmpv2CVariableBindingIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowSnmpv2CVariableBindingSlice = append(obj.flowSnmpv2CVariableBindingSlice, item)
	}
	return obj
}

func (obj *flowSnmpv2CPDUFlowSnmpv2CVariableBindingIter) Set(index int, newObj FlowSnmpv2CVariableBinding) FlowSnmpv2CPDUFlowSnmpv2CVariableBindingIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowSnmpv2CVariableBindingSlice[index] = newObj
	return obj
}
func (obj *flowSnmpv2CPDUFlowSnmpv2CVariableBindingIter) Clear() FlowSnmpv2CPDUFlowSnmpv2CVariableBindingIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowSnmpv2CVariableBinding{}
		obj.flowSnmpv2CVariableBindingSlice = []FlowSnmpv2CVariableBinding{}
	}
	return obj
}
func (obj *flowSnmpv2CPDUFlowSnmpv2CVariableBindingIter) clearHolderSlice() FlowSnmpv2CPDUFlowSnmpv2CVariableBindingIter {
	if len(obj.flowSnmpv2CVariableBindingSlice) > 0 {
		obj.flowSnmpv2CVariableBindingSlice = []FlowSnmpv2CVariableBinding{}
	}
	return obj
}
func (obj *flowSnmpv2CPDUFlowSnmpv2CVariableBindingIter) appendHolderSlice(item FlowSnmpv2CVariableBinding) FlowSnmpv2CPDUFlowSnmpv2CVariableBindingIter {
	obj.flowSnmpv2CVariableBindingSlice = append(obj.flowSnmpv2CVariableBindingSlice, item)
	return obj
}

func (obj *flowSnmpv2CPDU) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RequestId != nil {

		obj.RequestId().validateObj(vObj, set_default)
	}

	if obj.obj.ErrorIndex != nil {

		obj.ErrorIndex().validateObj(vObj, set_default)
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

func (obj *flowSnmpv2CPDU) setDefault() {
	if obj.obj.ErrorStatus == nil {
		obj.SetErrorStatus(FlowSnmpv2CPDUErrorStatus.NO_ERROR)

	}

}
