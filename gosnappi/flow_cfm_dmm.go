package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfmDmm *****
type flowCfmDmm struct {
	validation
	obj                          *otg.FlowCfmDmm
	marshaller                   marshalFlowCfmDmm
	unMarshaller                 unMarshalFlowCfmDmm
	txTimestampFHolder           PatternFlowCfmDmmTxTimestampF
	reservedForDmmReceiverHolder PatternFlowCfmDmmReservedForDmmReceiver
	reservedForDmrHolder         PatternFlowCfmDmmReservedForDmr
	reservedForDmrReceiverHolder PatternFlowCfmDmmReservedForDmrReceiver
	tlvsHolder                   FlowCfmDmmFlowCfmTlvsIter
}

func NewFlowCfmDmm() FlowCfmDmm {
	obj := flowCfmDmm{obj: &otg.FlowCfmDmm{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfmDmm) msg() *otg.FlowCfmDmm {
	return obj.obj
}

func (obj *flowCfmDmm) setMsg(msg *otg.FlowCfmDmm) FlowCfmDmm {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfmDmm struct {
	obj *flowCfmDmm
}

type marshalFlowCfmDmm interface {
	// ToProto marshals FlowCfmDmm to protobuf object *otg.FlowCfmDmm
	ToProto() (*otg.FlowCfmDmm, error)
	// ToPbText marshals FlowCfmDmm to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfmDmm to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfmDmm to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfmDmm struct {
	obj *flowCfmDmm
}

type unMarshalFlowCfmDmm interface {
	// FromProto unmarshals FlowCfmDmm from protobuf object *otg.FlowCfmDmm
	FromProto(msg *otg.FlowCfmDmm) (FlowCfmDmm, error)
	// FromPbText unmarshals FlowCfmDmm from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfmDmm from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfmDmm from JSON text
	FromJson(value string) error
}

func (obj *flowCfmDmm) Marshal() marshalFlowCfmDmm {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfmDmm{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfmDmm) Unmarshal() unMarshalFlowCfmDmm {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfmDmm{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfmDmm) ToProto() (*otg.FlowCfmDmm, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfmDmm) FromProto(msg *otg.FlowCfmDmm) (FlowCfmDmm, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfmDmm) ToPbText() (string, error) {
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

func (m *unMarshalflowCfmDmm) FromPbText(value string) error {
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

func (m *marshalflowCfmDmm) ToYaml() (string, error) {
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

func (m *unMarshalflowCfmDmm) FromYaml(value string) error {
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

func (m *marshalflowCfmDmm) ToJson() (string, error) {
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

func (m *unMarshalflowCfmDmm) FromJson(value string) error {
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

func (obj *flowCfmDmm) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfmDmm) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfmDmm) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfmDmm) Clone() (FlowCfmDmm, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfmDmm()
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

func (obj *flowCfmDmm) setNil() {
	obj.txTimestampFHolder = nil
	obj.reservedForDmmReceiverHolder = nil
	obj.reservedForDmrHolder = nil
	obj.reservedForDmrReceiverHolder = nil
	obj.tlvsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowCfmDmm is delay Measurement Message (DMM) is a diagnostic frame sent by a Maintenance End Point (MEP) to a peer MEP to measure one-way or two-way frame delay, as defined in ITU-T Y.1731 Section 8.2. DMM flag bits are reserved and set to zero.
type FlowCfmDmm interface {
	Validation
	// msg marshals FlowCfmDmm to protobuf object *otg.FlowCfmDmm
	// and doesn't set defaults
	msg() *otg.FlowCfmDmm
	// setMsg unmarshals FlowCfmDmm from protobuf object *otg.FlowCfmDmm
	// and doesn't set defaults
	setMsg(*otg.FlowCfmDmm) FlowCfmDmm
	// provides marshal interface
	Marshal() marshalFlowCfmDmm
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfmDmm
	// validate validates FlowCfmDmm
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfmDmm, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// TxTimestampF returns PatternFlowCfmDmmTxTimestampF, set in FlowCfmDmm.
	// PatternFlowCfmDmmTxTimestampF is the timestamp at which the DMM frame is transmitted by the originating MEP (ITU-T Y.1731 Section 8.2).
	TxTimestampF() PatternFlowCfmDmmTxTimestampF
	// SetTxTimestampF assigns PatternFlowCfmDmmTxTimestampF provided by user to FlowCfmDmm.
	// PatternFlowCfmDmmTxTimestampF is the timestamp at which the DMM frame is transmitted by the originating MEP (ITU-T Y.1731 Section 8.2).
	SetTxTimestampF(value PatternFlowCfmDmmTxTimestampF) FlowCfmDmm
	// HasTxTimestampF checks if TxTimestampF has been set in FlowCfmDmm
	HasTxTimestampF() bool
	// ReservedForDmmReceiver returns PatternFlowCfmDmmReservedForDmmReceiver, set in FlowCfmDmm.
	// PatternFlowCfmDmmReservedForDmmReceiver is reserved for the DMM receiving equipment to record its receive timestamp when forming the DMR. Set to zero by the DMM originator (ITU-T Y.1731 Section 8.2).
	ReservedForDmmReceiver() PatternFlowCfmDmmReservedForDmmReceiver
	// SetReservedForDmmReceiver assigns PatternFlowCfmDmmReservedForDmmReceiver provided by user to FlowCfmDmm.
	// PatternFlowCfmDmmReservedForDmmReceiver is reserved for the DMM receiving equipment to record its receive timestamp when forming the DMR. Set to zero by the DMM originator (ITU-T Y.1731 Section 8.2).
	SetReservedForDmmReceiver(value PatternFlowCfmDmmReservedForDmmReceiver) FlowCfmDmm
	// HasReservedForDmmReceiver checks if ReservedForDmmReceiver has been set in FlowCfmDmm
	HasReservedForDmmReceiver() bool
	// ReservedForDmr returns PatternFlowCfmDmmReservedForDmr, set in FlowCfmDmm.
	// PatternFlowCfmDmmReservedForDmr is reserved for the DMM receiving equipment to record its transmit timestamp when forming the DMR. Set to zero by the DMM originator (ITU-T Y.1731 Section 8.2).
	ReservedForDmr() PatternFlowCfmDmmReservedForDmr
	// SetReservedForDmr assigns PatternFlowCfmDmmReservedForDmr provided by user to FlowCfmDmm.
	// PatternFlowCfmDmmReservedForDmr is reserved for the DMM receiving equipment to record its transmit timestamp when forming the DMR. Set to zero by the DMM originator (ITU-T Y.1731 Section 8.2).
	SetReservedForDmr(value PatternFlowCfmDmmReservedForDmr) FlowCfmDmm
	// HasReservedForDmr checks if ReservedForDmr has been set in FlowCfmDmm
	HasReservedForDmr() bool
	// ReservedForDmrReceiver returns PatternFlowCfmDmmReservedForDmrReceiver, set in FlowCfmDmm.
	// PatternFlowCfmDmmReservedForDmrReceiver is reserved for the DMR receiving equipment to record its receive timestamp when the DMR arrives back at the originating MEP. Set to zero by the DMM originator (ITU-T Y.1731 Section 8.2).
	ReservedForDmrReceiver() PatternFlowCfmDmmReservedForDmrReceiver
	// SetReservedForDmrReceiver assigns PatternFlowCfmDmmReservedForDmrReceiver provided by user to FlowCfmDmm.
	// PatternFlowCfmDmmReservedForDmrReceiver is reserved for the DMR receiving equipment to record its receive timestamp when the DMR arrives back at the originating MEP. Set to zero by the DMM originator (ITU-T Y.1731 Section 8.2).
	SetReservedForDmrReceiver(value PatternFlowCfmDmmReservedForDmrReceiver) FlowCfmDmm
	// HasReservedForDmrReceiver checks if ReservedForDmrReceiver has been set in FlowCfmDmm
	HasReservedForDmrReceiver() bool
	// Tlvs returns FlowCfmDmmFlowCfmTlvsIterIter, set in FlowCfmDmm
	Tlvs() FlowCfmDmmFlowCfmTlvsIter
	// AppendEndTlv returns bool, set in FlowCfmDmm.
	AppendEndTlv() bool
	// SetAppendEndTlv assigns bool provided by user to FlowCfmDmm
	SetAppendEndTlv(value bool) FlowCfmDmm
	// HasAppendEndTlv checks if AppendEndTlv has been set in FlowCfmDmm
	HasAppendEndTlv() bool
	setNil()
}

// description is TBD
// TxTimestampF returns a PatternFlowCfmDmmTxTimestampF
func (obj *flowCfmDmm) TxTimestampF() PatternFlowCfmDmmTxTimestampF {
	if obj.obj.TxTimestampF == nil {
		obj.obj.TxTimestampF = NewPatternFlowCfmDmmTxTimestampF().msg()
	}
	if obj.txTimestampFHolder == nil {
		obj.txTimestampFHolder = &patternFlowCfmDmmTxTimestampF{obj: obj.obj.TxTimestampF}
	}
	return obj.txTimestampFHolder
}

// description is TBD
// TxTimestampF returns a PatternFlowCfmDmmTxTimestampF
func (obj *flowCfmDmm) HasTxTimestampF() bool {
	return obj.obj.TxTimestampF != nil
}

// description is TBD
// SetTxTimestampF sets the PatternFlowCfmDmmTxTimestampF value in the FlowCfmDmm object
func (obj *flowCfmDmm) SetTxTimestampF(value PatternFlowCfmDmmTxTimestampF) FlowCfmDmm {

	obj.txTimestampFHolder = nil
	obj.obj.TxTimestampF = value.msg()

	return obj
}

// description is TBD
// ReservedForDmmReceiver returns a PatternFlowCfmDmmReservedForDmmReceiver
func (obj *flowCfmDmm) ReservedForDmmReceiver() PatternFlowCfmDmmReservedForDmmReceiver {
	if obj.obj.ReservedForDmmReceiver == nil {
		obj.obj.ReservedForDmmReceiver = NewPatternFlowCfmDmmReservedForDmmReceiver().msg()
	}
	if obj.reservedForDmmReceiverHolder == nil {
		obj.reservedForDmmReceiverHolder = &patternFlowCfmDmmReservedForDmmReceiver{obj: obj.obj.ReservedForDmmReceiver}
	}
	return obj.reservedForDmmReceiverHolder
}

// description is TBD
// ReservedForDmmReceiver returns a PatternFlowCfmDmmReservedForDmmReceiver
func (obj *flowCfmDmm) HasReservedForDmmReceiver() bool {
	return obj.obj.ReservedForDmmReceiver != nil
}

// description is TBD
// SetReservedForDmmReceiver sets the PatternFlowCfmDmmReservedForDmmReceiver value in the FlowCfmDmm object
func (obj *flowCfmDmm) SetReservedForDmmReceiver(value PatternFlowCfmDmmReservedForDmmReceiver) FlowCfmDmm {

	obj.reservedForDmmReceiverHolder = nil
	obj.obj.ReservedForDmmReceiver = value.msg()

	return obj
}

// description is TBD
// ReservedForDmr returns a PatternFlowCfmDmmReservedForDmr
func (obj *flowCfmDmm) ReservedForDmr() PatternFlowCfmDmmReservedForDmr {
	if obj.obj.ReservedForDmr == nil {
		obj.obj.ReservedForDmr = NewPatternFlowCfmDmmReservedForDmr().msg()
	}
	if obj.reservedForDmrHolder == nil {
		obj.reservedForDmrHolder = &patternFlowCfmDmmReservedForDmr{obj: obj.obj.ReservedForDmr}
	}
	return obj.reservedForDmrHolder
}

// description is TBD
// ReservedForDmr returns a PatternFlowCfmDmmReservedForDmr
func (obj *flowCfmDmm) HasReservedForDmr() bool {
	return obj.obj.ReservedForDmr != nil
}

// description is TBD
// SetReservedForDmr sets the PatternFlowCfmDmmReservedForDmr value in the FlowCfmDmm object
func (obj *flowCfmDmm) SetReservedForDmr(value PatternFlowCfmDmmReservedForDmr) FlowCfmDmm {

	obj.reservedForDmrHolder = nil
	obj.obj.ReservedForDmr = value.msg()

	return obj
}

// description is TBD
// ReservedForDmrReceiver returns a PatternFlowCfmDmmReservedForDmrReceiver
func (obj *flowCfmDmm) ReservedForDmrReceiver() PatternFlowCfmDmmReservedForDmrReceiver {
	if obj.obj.ReservedForDmrReceiver == nil {
		obj.obj.ReservedForDmrReceiver = NewPatternFlowCfmDmmReservedForDmrReceiver().msg()
	}
	if obj.reservedForDmrReceiverHolder == nil {
		obj.reservedForDmrReceiverHolder = &patternFlowCfmDmmReservedForDmrReceiver{obj: obj.obj.ReservedForDmrReceiver}
	}
	return obj.reservedForDmrReceiverHolder
}

// description is TBD
// ReservedForDmrReceiver returns a PatternFlowCfmDmmReservedForDmrReceiver
func (obj *flowCfmDmm) HasReservedForDmrReceiver() bool {
	return obj.obj.ReservedForDmrReceiver != nil
}

// description is TBD
// SetReservedForDmrReceiver sets the PatternFlowCfmDmmReservedForDmrReceiver value in the FlowCfmDmm object
func (obj *flowCfmDmm) SetReservedForDmrReceiver(value PatternFlowCfmDmmReservedForDmrReceiver) FlowCfmDmm {

	obj.reservedForDmrReceiverHolder = nil
	obj.obj.ReservedForDmrReceiver = value.msg()

	return obj
}

// description is TBD
// Tlvs returns a []FlowCfmTlvs
func (obj *flowCfmDmm) Tlvs() FlowCfmDmmFlowCfmTlvsIter {
	if len(obj.obj.Tlvs) == 0 {
		obj.obj.Tlvs = []*otg.FlowCfmTlvs{}
	}
	if obj.tlvsHolder == nil {
		obj.tlvsHolder = newFlowCfmDmmFlowCfmTlvsIter(&obj.obj.Tlvs).setMsg(obj)
	}
	return obj.tlvsHolder
}

type flowCfmDmmFlowCfmTlvsIter struct {
	obj              *flowCfmDmm
	flowCfmTlvsSlice []FlowCfmTlvs
	fieldPtr         *[]*otg.FlowCfmTlvs
}

func newFlowCfmDmmFlowCfmTlvsIter(ptr *[]*otg.FlowCfmTlvs) FlowCfmDmmFlowCfmTlvsIter {
	return &flowCfmDmmFlowCfmTlvsIter{fieldPtr: ptr}
}

type FlowCfmDmmFlowCfmTlvsIter interface {
	setMsg(*flowCfmDmm) FlowCfmDmmFlowCfmTlvsIter
	Items() []FlowCfmTlvs
	Add() FlowCfmTlvs
	Append(items ...FlowCfmTlvs) FlowCfmDmmFlowCfmTlvsIter
	Set(index int, newObj FlowCfmTlvs) FlowCfmDmmFlowCfmTlvsIter
	Clear() FlowCfmDmmFlowCfmTlvsIter
	clearHolderSlice() FlowCfmDmmFlowCfmTlvsIter
	appendHolderSlice(item FlowCfmTlvs) FlowCfmDmmFlowCfmTlvsIter
}

func (obj *flowCfmDmmFlowCfmTlvsIter) setMsg(msg *flowCfmDmm) FlowCfmDmmFlowCfmTlvsIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowCfmTlvs{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowCfmDmmFlowCfmTlvsIter) Items() []FlowCfmTlvs {
	return obj.flowCfmTlvsSlice
}

func (obj *flowCfmDmmFlowCfmTlvsIter) Add() FlowCfmTlvs {
	newObj := &otg.FlowCfmTlvs{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowCfmTlvs{obj: newObj}
	newLibObj.setDefault()
	obj.flowCfmTlvsSlice = append(obj.flowCfmTlvsSlice, newLibObj)
	return newLibObj
}

func (obj *flowCfmDmmFlowCfmTlvsIter) Append(items ...FlowCfmTlvs) FlowCfmDmmFlowCfmTlvsIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowCfmTlvsSlice = append(obj.flowCfmTlvsSlice, item)
	}
	return obj
}

func (obj *flowCfmDmmFlowCfmTlvsIter) Set(index int, newObj FlowCfmTlvs) FlowCfmDmmFlowCfmTlvsIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowCfmTlvsSlice[index] = newObj
	return obj
}
func (obj *flowCfmDmmFlowCfmTlvsIter) Clear() FlowCfmDmmFlowCfmTlvsIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowCfmTlvs{}
		obj.flowCfmTlvsSlice = []FlowCfmTlvs{}
	}
	return obj
}
func (obj *flowCfmDmmFlowCfmTlvsIter) clearHolderSlice() FlowCfmDmmFlowCfmTlvsIter {
	if len(obj.flowCfmTlvsSlice) > 0 {
		obj.flowCfmTlvsSlice = []FlowCfmTlvs{}
	}
	return obj
}
func (obj *flowCfmDmmFlowCfmTlvsIter) appendHolderSlice(item FlowCfmTlvs) FlowCfmDmmFlowCfmTlvsIter {
	obj.flowCfmTlvsSlice = append(obj.flowCfmTlvsSlice, item)
	return obj
}

// When true, the implementation appends an End TLV (Type 0, IEEE 802.1Q-2018 Table 21-17) after all TLVs in the tlvs list. Set to false only when testing DUT behaviour with a deliberately missing End TLV.
// AppendEndTlv returns a bool
func (obj *flowCfmDmm) AppendEndTlv() bool {

	return *obj.obj.AppendEndTlv

}

// When true, the implementation appends an End TLV (Type 0, IEEE 802.1Q-2018 Table 21-17) after all TLVs in the tlvs list. Set to false only when testing DUT behaviour with a deliberately missing End TLV.
// AppendEndTlv returns a bool
func (obj *flowCfmDmm) HasAppendEndTlv() bool {
	return obj.obj.AppendEndTlv != nil
}

// When true, the implementation appends an End TLV (Type 0, IEEE 802.1Q-2018 Table 21-17) after all TLVs in the tlvs list. Set to false only when testing DUT behaviour with a deliberately missing End TLV.
// SetAppendEndTlv sets the bool value in the FlowCfmDmm object
func (obj *flowCfmDmm) SetAppendEndTlv(value bool) FlowCfmDmm {

	obj.obj.AppendEndTlv = &value
	return obj
}

func (obj *flowCfmDmm) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.TxTimestampF != nil {

		obj.TxTimestampF().validateObj(vObj, set_default)
	}

	if obj.obj.ReservedForDmmReceiver != nil {

		obj.ReservedForDmmReceiver().validateObj(vObj, set_default)
	}

	if obj.obj.ReservedForDmr != nil {

		obj.ReservedForDmr().validateObj(vObj, set_default)
	}

	if obj.obj.ReservedForDmrReceiver != nil {

		obj.ReservedForDmrReceiver().validateObj(vObj, set_default)
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

func (obj *flowCfmDmm) setDefault() {
	if obj.obj.AppendEndTlv == nil {
		obj.SetAppendEndTlv(true)
	}

}
