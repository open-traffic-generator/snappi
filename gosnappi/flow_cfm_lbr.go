package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfmLbr *****
type flowCfmLbr struct {
	validation
	obj                 *otg.FlowCfmLbr
	marshaller          marshalFlowCfmLbr
	unMarshaller        unMarshalFlowCfmLbr
	flagsHolder         PatternFlowCfmLbrFlags
	transactionIdHolder PatternFlowCfmLbrTransactionId
	tlvsHolder          FlowCfmLbrFlowCfmTlvsIter
}

func NewFlowCfmLbr() FlowCfmLbr {
	obj := flowCfmLbr{obj: &otg.FlowCfmLbr{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfmLbr) msg() *otg.FlowCfmLbr {
	return obj.obj
}

func (obj *flowCfmLbr) setMsg(msg *otg.FlowCfmLbr) FlowCfmLbr {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfmLbr struct {
	obj *flowCfmLbr
}

type marshalFlowCfmLbr interface {
	// ToProto marshals FlowCfmLbr to protobuf object *otg.FlowCfmLbr
	ToProto() (*otg.FlowCfmLbr, error)
	// ToPbText marshals FlowCfmLbr to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfmLbr to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfmLbr to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfmLbr struct {
	obj *flowCfmLbr
}

type unMarshalFlowCfmLbr interface {
	// FromProto unmarshals FlowCfmLbr from protobuf object *otg.FlowCfmLbr
	FromProto(msg *otg.FlowCfmLbr) (FlowCfmLbr, error)
	// FromPbText unmarshals FlowCfmLbr from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfmLbr from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfmLbr from JSON text
	FromJson(value string) error
}

func (obj *flowCfmLbr) Marshal() marshalFlowCfmLbr {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfmLbr{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfmLbr) Unmarshal() unMarshalFlowCfmLbr {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfmLbr{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfmLbr) ToProto() (*otg.FlowCfmLbr, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfmLbr) FromProto(msg *otg.FlowCfmLbr) (FlowCfmLbr, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfmLbr) ToPbText() (string, error) {
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

func (m *unMarshalflowCfmLbr) FromPbText(value string) error {
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

func (m *marshalflowCfmLbr) ToYaml() (string, error) {
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

func (m *unMarshalflowCfmLbr) FromYaml(value string) error {
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

func (m *marshalflowCfmLbr) ToJson() (string, error) {
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

func (m *unMarshalflowCfmLbr) FromJson(value string) error {
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

func (obj *flowCfmLbr) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfmLbr) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfmLbr) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfmLbr) Clone() (FlowCfmLbr, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfmLbr()
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

func (obj *flowCfmLbr) setNil() {
	obj.flagsHolder = nil
	obj.transactionIdHolder = nil
	obj.tlvsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowCfmLbr is loopback Reply (LBR) messages are unicast frames sent by a Maintenance Association End Point (MEP) or Intermediate Point (MIP) to a source MEP in response to a Loopback Message (LBM).
type FlowCfmLbr interface {
	Validation
	// msg marshals FlowCfmLbr to protobuf object *otg.FlowCfmLbr
	// and doesn't set defaults
	msg() *otg.FlowCfmLbr
	// setMsg unmarshals FlowCfmLbr from protobuf object *otg.FlowCfmLbr
	// and doesn't set defaults
	setMsg(*otg.FlowCfmLbr) FlowCfmLbr
	// provides marshal interface
	Marshal() marshalFlowCfmLbr
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfmLbr
	// validate validates FlowCfmLbr
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfmLbr, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns PatternFlowCfmLbrFlags, set in FlowCfmLbr.
	// PatternFlowCfmLbrFlags is defines operational flags for LBR.
	Flags() PatternFlowCfmLbrFlags
	// SetFlags assigns PatternFlowCfmLbrFlags provided by user to FlowCfmLbr.
	// PatternFlowCfmLbrFlags is defines operational flags for LBR.
	SetFlags(value PatternFlowCfmLbrFlags) FlowCfmLbr
	// HasFlags checks if Flags has been set in FlowCfmLbr
	HasFlags() bool
	// TransactionId returns PatternFlowCfmLbrTransactionId, set in FlowCfmLbr.
	// PatternFlowCfmLbrTransactionId is transaction id or sequence number copied from the LBM to match the reply to the specific request.
	TransactionId() PatternFlowCfmLbrTransactionId
	// SetTransactionId assigns PatternFlowCfmLbrTransactionId provided by user to FlowCfmLbr.
	// PatternFlowCfmLbrTransactionId is transaction id or sequence number copied from the LBM to match the reply to the specific request.
	SetTransactionId(value PatternFlowCfmLbrTransactionId) FlowCfmLbr
	// HasTransactionId checks if TransactionId has been set in FlowCfmLbr
	HasTransactionId() bool
	// Tlvs returns FlowCfmLbrFlowCfmTlvsIterIter, set in FlowCfmLbr
	Tlvs() FlowCfmLbrFlowCfmTlvsIter
	setNil()
}

// description is TBD
// Flags returns a PatternFlowCfmLbrFlags
func (obj *flowCfmLbr) Flags() PatternFlowCfmLbrFlags {
	if obj.obj.Flags == nil {
		obj.obj.Flags = NewPatternFlowCfmLbrFlags().msg()
	}
	if obj.flagsHolder == nil {
		obj.flagsHolder = &patternFlowCfmLbrFlags{obj: obj.obj.Flags}
	}
	return obj.flagsHolder
}

// description is TBD
// Flags returns a PatternFlowCfmLbrFlags
func (obj *flowCfmLbr) HasFlags() bool {
	return obj.obj.Flags != nil
}

// description is TBD
// SetFlags sets the PatternFlowCfmLbrFlags value in the FlowCfmLbr object
func (obj *flowCfmLbr) SetFlags(value PatternFlowCfmLbrFlags) FlowCfmLbr {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// description is TBD
// TransactionId returns a PatternFlowCfmLbrTransactionId
func (obj *flowCfmLbr) TransactionId() PatternFlowCfmLbrTransactionId {
	if obj.obj.TransactionId == nil {
		obj.obj.TransactionId = NewPatternFlowCfmLbrTransactionId().msg()
	}
	if obj.transactionIdHolder == nil {
		obj.transactionIdHolder = &patternFlowCfmLbrTransactionId{obj: obj.obj.TransactionId}
	}
	return obj.transactionIdHolder
}

// description is TBD
// TransactionId returns a PatternFlowCfmLbrTransactionId
func (obj *flowCfmLbr) HasTransactionId() bool {
	return obj.obj.TransactionId != nil
}

// description is TBD
// SetTransactionId sets the PatternFlowCfmLbrTransactionId value in the FlowCfmLbr object
func (obj *flowCfmLbr) SetTransactionId(value PatternFlowCfmLbrTransactionId) FlowCfmLbr {

	obj.transactionIdHolder = nil
	obj.obj.TransactionId = value.msg()

	return obj
}

// description is TBD
// Tlvs returns a []FlowCfmTlvs
func (obj *flowCfmLbr) Tlvs() FlowCfmLbrFlowCfmTlvsIter {
	if len(obj.obj.Tlvs) == 0 {
		obj.obj.Tlvs = []*otg.FlowCfmTlvs{}
	}
	if obj.tlvsHolder == nil {
		obj.tlvsHolder = newFlowCfmLbrFlowCfmTlvsIter(&obj.obj.Tlvs).setMsg(obj)
	}
	return obj.tlvsHolder
}

type flowCfmLbrFlowCfmTlvsIter struct {
	obj              *flowCfmLbr
	flowCfmTlvsSlice []FlowCfmTlvs
	fieldPtr         *[]*otg.FlowCfmTlvs
}

func newFlowCfmLbrFlowCfmTlvsIter(ptr *[]*otg.FlowCfmTlvs) FlowCfmLbrFlowCfmTlvsIter {
	return &flowCfmLbrFlowCfmTlvsIter{fieldPtr: ptr}
}

type FlowCfmLbrFlowCfmTlvsIter interface {
	setMsg(*flowCfmLbr) FlowCfmLbrFlowCfmTlvsIter
	Items() []FlowCfmTlvs
	Add() FlowCfmTlvs
	Append(items ...FlowCfmTlvs) FlowCfmLbrFlowCfmTlvsIter
	Set(index int, newObj FlowCfmTlvs) FlowCfmLbrFlowCfmTlvsIter
	Clear() FlowCfmLbrFlowCfmTlvsIter
	clearHolderSlice() FlowCfmLbrFlowCfmTlvsIter
	appendHolderSlice(item FlowCfmTlvs) FlowCfmLbrFlowCfmTlvsIter
}

func (obj *flowCfmLbrFlowCfmTlvsIter) setMsg(msg *flowCfmLbr) FlowCfmLbrFlowCfmTlvsIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowCfmTlvs{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowCfmLbrFlowCfmTlvsIter) Items() []FlowCfmTlvs {
	return obj.flowCfmTlvsSlice
}

func (obj *flowCfmLbrFlowCfmTlvsIter) Add() FlowCfmTlvs {
	newObj := &otg.FlowCfmTlvs{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowCfmTlvs{obj: newObj}
	newLibObj.setDefault()
	obj.flowCfmTlvsSlice = append(obj.flowCfmTlvsSlice, newLibObj)
	return newLibObj
}

func (obj *flowCfmLbrFlowCfmTlvsIter) Append(items ...FlowCfmTlvs) FlowCfmLbrFlowCfmTlvsIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowCfmTlvsSlice = append(obj.flowCfmTlvsSlice, item)
	}
	return obj
}

func (obj *flowCfmLbrFlowCfmTlvsIter) Set(index int, newObj FlowCfmTlvs) FlowCfmLbrFlowCfmTlvsIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowCfmTlvsSlice[index] = newObj
	return obj
}
func (obj *flowCfmLbrFlowCfmTlvsIter) Clear() FlowCfmLbrFlowCfmTlvsIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowCfmTlvs{}
		obj.flowCfmTlvsSlice = []FlowCfmTlvs{}
	}
	return obj
}
func (obj *flowCfmLbrFlowCfmTlvsIter) clearHolderSlice() FlowCfmLbrFlowCfmTlvsIter {
	if len(obj.flowCfmTlvsSlice) > 0 {
		obj.flowCfmTlvsSlice = []FlowCfmTlvs{}
	}
	return obj
}
func (obj *flowCfmLbrFlowCfmTlvsIter) appendHolderSlice(item FlowCfmTlvs) FlowCfmLbrFlowCfmTlvsIter {
	obj.flowCfmTlvsSlice = append(obj.flowCfmTlvsSlice, item)
	return obj
}

func (obj *flowCfmLbr) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(vObj, set_default)
	}

	if obj.obj.TransactionId != nil {

		obj.TransactionId().validateObj(vObj, set_default)
	}

	if len(obj.obj.Tlvs) != 0 {

		if set_default {
			obj.Tlvs().clearHolderSlice()
			for _, item := range obj.obj.Tlvs {
				obj.Tlvs().appendHolderSlice(&flowCfmTlvs{obj: item})
			}
		}
		for _, item := range obj.Tlvs().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *flowCfmLbr) setDefault() {

}
