package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfmDmr *****
type flowCfmDmr struct {
	validation
	obj                          *otg.FlowCfmDmr
	marshaller                   marshalFlowCfmDmr
	unMarshaller                 unMarshalFlowCfmDmr
	txTimestampFHolder           PatternFlowCfmDmrTxTimestampF
	rxTimestampFHolder           PatternFlowCfmDmrRxTimestampF
	txTimestampBHolder           PatternFlowCfmDmrTxTimestampB
	reservedForDmrReceiverHolder PatternFlowCfmDmrReservedForDmrReceiver
	tlvsHolder                   FlowCfmDmrFlowCfmTlvsIter
}

func NewFlowCfmDmr() FlowCfmDmr {
	obj := flowCfmDmr{obj: &otg.FlowCfmDmr{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfmDmr) msg() *otg.FlowCfmDmr {
	return obj.obj
}

func (obj *flowCfmDmr) setMsg(msg *otg.FlowCfmDmr) FlowCfmDmr {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfmDmr struct {
	obj *flowCfmDmr
}

type marshalFlowCfmDmr interface {
	// ToProto marshals FlowCfmDmr to protobuf object *otg.FlowCfmDmr
	ToProto() (*otg.FlowCfmDmr, error)
	// ToPbText marshals FlowCfmDmr to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfmDmr to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfmDmr to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfmDmr struct {
	obj *flowCfmDmr
}

type unMarshalFlowCfmDmr interface {
	// FromProto unmarshals FlowCfmDmr from protobuf object *otg.FlowCfmDmr
	FromProto(msg *otg.FlowCfmDmr) (FlowCfmDmr, error)
	// FromPbText unmarshals FlowCfmDmr from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfmDmr from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfmDmr from JSON text
	FromJson(value string) error
}

func (obj *flowCfmDmr) Marshal() marshalFlowCfmDmr {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfmDmr{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfmDmr) Unmarshal() unMarshalFlowCfmDmr {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfmDmr{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfmDmr) ToProto() (*otg.FlowCfmDmr, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfmDmr) FromProto(msg *otg.FlowCfmDmr) (FlowCfmDmr, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfmDmr) ToPbText() (string, error) {
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

func (m *unMarshalflowCfmDmr) FromPbText(value string) error {
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

func (m *marshalflowCfmDmr) ToYaml() (string, error) {
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

func (m *unMarshalflowCfmDmr) FromYaml(value string) error {
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

func (m *marshalflowCfmDmr) ToJson() (string, error) {
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

func (m *unMarshalflowCfmDmr) FromJson(value string) error {
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

func (obj *flowCfmDmr) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfmDmr) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfmDmr) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfmDmr) Clone() (FlowCfmDmr, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfmDmr()
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

func (obj *flowCfmDmr) setNil() {
	obj.txTimestampFHolder = nil
	obj.rxTimestampFHolder = nil
	obj.txTimestampBHolder = nil
	obj.reservedForDmrReceiverHolder = nil
	obj.tlvsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowCfmDmr is delay Measurement Reply (DMR) is the response frame sent by a peer Maintenance End Point (MEP) back to the originating MEP in reply to a Delay Measurement Message (DMM), as defined in ITU-T Y.1731 Section 8.2. DMR flag bits are reserved and set to zero.
type FlowCfmDmr interface {
	Validation
	// msg marshals FlowCfmDmr to protobuf object *otg.FlowCfmDmr
	// and doesn't set defaults
	msg() *otg.FlowCfmDmr
	// setMsg unmarshals FlowCfmDmr from protobuf object *otg.FlowCfmDmr
	// and doesn't set defaults
	setMsg(*otg.FlowCfmDmr) FlowCfmDmr
	// provides marshal interface
	Marshal() marshalFlowCfmDmr
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfmDmr
	// validate validates FlowCfmDmr
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfmDmr, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// TxTimestampF returns PatternFlowCfmDmrTxTimestampF, set in FlowCfmDmr.
	// PatternFlowCfmDmrTxTimestampF is transmit timestamp recorded by the DMM originator, copied verbatim by the responder from the received DMM into the DMR (ITU-T Y.1731 Section 8.2).
	TxTimestampF() PatternFlowCfmDmrTxTimestampF
	// SetTxTimestampF assigns PatternFlowCfmDmrTxTimestampF provided by user to FlowCfmDmr.
	// PatternFlowCfmDmrTxTimestampF is transmit timestamp recorded by the DMM originator, copied verbatim by the responder from the received DMM into the DMR (ITU-T Y.1731 Section 8.2).
	SetTxTimestampF(value PatternFlowCfmDmrTxTimestampF) FlowCfmDmr
	// HasTxTimestampF checks if TxTimestampF has been set in FlowCfmDmr
	HasTxTimestampF() bool
	// RxTimestampF returns PatternFlowCfmDmrRxTimestampF, set in FlowCfmDmr.
	// PatternFlowCfmDmrRxTimestampF is receive timestamp recorded by the responder when the DMM frame arrived (ITU-T Y.1731 Section 8.2).
	RxTimestampF() PatternFlowCfmDmrRxTimestampF
	// SetRxTimestampF assigns PatternFlowCfmDmrRxTimestampF provided by user to FlowCfmDmr.
	// PatternFlowCfmDmrRxTimestampF is receive timestamp recorded by the responder when the DMM frame arrived (ITU-T Y.1731 Section 8.2).
	SetRxTimestampF(value PatternFlowCfmDmrRxTimestampF) FlowCfmDmr
	// HasRxTimestampF checks if RxTimestampF has been set in FlowCfmDmr
	HasRxTimestampF() bool
	// TxTimestampB returns PatternFlowCfmDmrTxTimestampB, set in FlowCfmDmr.
	// PatternFlowCfmDmrTxTimestampB is transmit timestamp recorded by the responder when the DMR frame is transmitted back to the originating MEP (ITU-T Y.1731 Section 8.2).
	TxTimestampB() PatternFlowCfmDmrTxTimestampB
	// SetTxTimestampB assigns PatternFlowCfmDmrTxTimestampB provided by user to FlowCfmDmr.
	// PatternFlowCfmDmrTxTimestampB is transmit timestamp recorded by the responder when the DMR frame is transmitted back to the originating MEP (ITU-T Y.1731 Section 8.2).
	SetTxTimestampB(value PatternFlowCfmDmrTxTimestampB) FlowCfmDmr
	// HasTxTimestampB checks if TxTimestampB has been set in FlowCfmDmr
	HasTxTimestampB() bool
	// ReservedForDmrReceiver returns PatternFlowCfmDmrReservedForDmrReceiver, set in FlowCfmDmr.
	// PatternFlowCfmDmrReservedForDmrReceiver is reserved for the DMR receiving equipment (the originating MEP) to record its receive timestamp when the DMR arrives. Set to zero by the DMR originator (ITU-T Y.1731 Section 8.2).
	ReservedForDmrReceiver() PatternFlowCfmDmrReservedForDmrReceiver
	// SetReservedForDmrReceiver assigns PatternFlowCfmDmrReservedForDmrReceiver provided by user to FlowCfmDmr.
	// PatternFlowCfmDmrReservedForDmrReceiver is reserved for the DMR receiving equipment (the originating MEP) to record its receive timestamp when the DMR arrives. Set to zero by the DMR originator (ITU-T Y.1731 Section 8.2).
	SetReservedForDmrReceiver(value PatternFlowCfmDmrReservedForDmrReceiver) FlowCfmDmr
	// HasReservedForDmrReceiver checks if ReservedForDmrReceiver has been set in FlowCfmDmr
	HasReservedForDmrReceiver() bool
	// Tlvs returns FlowCfmDmrFlowCfmTlvsIterIter, set in FlowCfmDmr
	Tlvs() FlowCfmDmrFlowCfmTlvsIter
	// AppendEndTlv returns bool, set in FlowCfmDmr.
	AppendEndTlv() bool
	// SetAppendEndTlv assigns bool provided by user to FlowCfmDmr
	SetAppendEndTlv(value bool) FlowCfmDmr
	// HasAppendEndTlv checks if AppendEndTlv has been set in FlowCfmDmr
	HasAppendEndTlv() bool
	setNil()
}

// description is TBD
// TxTimestampF returns a PatternFlowCfmDmrTxTimestampF
func (obj *flowCfmDmr) TxTimestampF() PatternFlowCfmDmrTxTimestampF {
	if obj.obj.TxTimestampF == nil {
		obj.obj.TxTimestampF = NewPatternFlowCfmDmrTxTimestampF().msg()
	}
	if obj.txTimestampFHolder == nil {
		obj.txTimestampFHolder = &patternFlowCfmDmrTxTimestampF{obj: obj.obj.TxTimestampF}
	}
	return obj.txTimestampFHolder
}

// description is TBD
// TxTimestampF returns a PatternFlowCfmDmrTxTimestampF
func (obj *flowCfmDmr) HasTxTimestampF() bool {
	return obj.obj.TxTimestampF != nil
}

// description is TBD
// SetTxTimestampF sets the PatternFlowCfmDmrTxTimestampF value in the FlowCfmDmr object
func (obj *flowCfmDmr) SetTxTimestampF(value PatternFlowCfmDmrTxTimestampF) FlowCfmDmr {

	obj.txTimestampFHolder = nil
	obj.obj.TxTimestampF = value.msg()

	return obj
}

// description is TBD
// RxTimestampF returns a PatternFlowCfmDmrRxTimestampF
func (obj *flowCfmDmr) RxTimestampF() PatternFlowCfmDmrRxTimestampF {
	if obj.obj.RxTimestampF == nil {
		obj.obj.RxTimestampF = NewPatternFlowCfmDmrRxTimestampF().msg()
	}
	if obj.rxTimestampFHolder == nil {
		obj.rxTimestampFHolder = &patternFlowCfmDmrRxTimestampF{obj: obj.obj.RxTimestampF}
	}
	return obj.rxTimestampFHolder
}

// description is TBD
// RxTimestampF returns a PatternFlowCfmDmrRxTimestampF
func (obj *flowCfmDmr) HasRxTimestampF() bool {
	return obj.obj.RxTimestampF != nil
}

// description is TBD
// SetRxTimestampF sets the PatternFlowCfmDmrRxTimestampF value in the FlowCfmDmr object
func (obj *flowCfmDmr) SetRxTimestampF(value PatternFlowCfmDmrRxTimestampF) FlowCfmDmr {

	obj.rxTimestampFHolder = nil
	obj.obj.RxTimestampF = value.msg()

	return obj
}

// description is TBD
// TxTimestampB returns a PatternFlowCfmDmrTxTimestampB
func (obj *flowCfmDmr) TxTimestampB() PatternFlowCfmDmrTxTimestampB {
	if obj.obj.TxTimestampB == nil {
		obj.obj.TxTimestampB = NewPatternFlowCfmDmrTxTimestampB().msg()
	}
	if obj.txTimestampBHolder == nil {
		obj.txTimestampBHolder = &patternFlowCfmDmrTxTimestampB{obj: obj.obj.TxTimestampB}
	}
	return obj.txTimestampBHolder
}

// description is TBD
// TxTimestampB returns a PatternFlowCfmDmrTxTimestampB
func (obj *flowCfmDmr) HasTxTimestampB() bool {
	return obj.obj.TxTimestampB != nil
}

// description is TBD
// SetTxTimestampB sets the PatternFlowCfmDmrTxTimestampB value in the FlowCfmDmr object
func (obj *flowCfmDmr) SetTxTimestampB(value PatternFlowCfmDmrTxTimestampB) FlowCfmDmr {

	obj.txTimestampBHolder = nil
	obj.obj.TxTimestampB = value.msg()

	return obj
}

// description is TBD
// ReservedForDmrReceiver returns a PatternFlowCfmDmrReservedForDmrReceiver
func (obj *flowCfmDmr) ReservedForDmrReceiver() PatternFlowCfmDmrReservedForDmrReceiver {
	if obj.obj.ReservedForDmrReceiver == nil {
		obj.obj.ReservedForDmrReceiver = NewPatternFlowCfmDmrReservedForDmrReceiver().msg()
	}
	if obj.reservedForDmrReceiverHolder == nil {
		obj.reservedForDmrReceiverHolder = &patternFlowCfmDmrReservedForDmrReceiver{obj: obj.obj.ReservedForDmrReceiver}
	}
	return obj.reservedForDmrReceiverHolder
}

// description is TBD
// ReservedForDmrReceiver returns a PatternFlowCfmDmrReservedForDmrReceiver
func (obj *flowCfmDmr) HasReservedForDmrReceiver() bool {
	return obj.obj.ReservedForDmrReceiver != nil
}

// description is TBD
// SetReservedForDmrReceiver sets the PatternFlowCfmDmrReservedForDmrReceiver value in the FlowCfmDmr object
func (obj *flowCfmDmr) SetReservedForDmrReceiver(value PatternFlowCfmDmrReservedForDmrReceiver) FlowCfmDmr {

	obj.reservedForDmrReceiverHolder = nil
	obj.obj.ReservedForDmrReceiver = value.msg()

	return obj
}

// description is TBD
// Tlvs returns a []FlowCfmTlvs
func (obj *flowCfmDmr) Tlvs() FlowCfmDmrFlowCfmTlvsIter {
	if len(obj.obj.Tlvs) == 0 {
		obj.obj.Tlvs = []*otg.FlowCfmTlvs{}
	}
	if obj.tlvsHolder == nil {
		obj.tlvsHolder = newFlowCfmDmrFlowCfmTlvsIter(&obj.obj.Tlvs).setMsg(obj)
	}
	return obj.tlvsHolder
}

type flowCfmDmrFlowCfmTlvsIter struct {
	obj              *flowCfmDmr
	flowCfmTlvsSlice []FlowCfmTlvs
	fieldPtr         *[]*otg.FlowCfmTlvs
}

func newFlowCfmDmrFlowCfmTlvsIter(ptr *[]*otg.FlowCfmTlvs) FlowCfmDmrFlowCfmTlvsIter {
	return &flowCfmDmrFlowCfmTlvsIter{fieldPtr: ptr}
}

type FlowCfmDmrFlowCfmTlvsIter interface {
	setMsg(*flowCfmDmr) FlowCfmDmrFlowCfmTlvsIter
	Items() []FlowCfmTlvs
	Add() FlowCfmTlvs
	Append(items ...FlowCfmTlvs) FlowCfmDmrFlowCfmTlvsIter
	Set(index int, newObj FlowCfmTlvs) FlowCfmDmrFlowCfmTlvsIter
	Clear() FlowCfmDmrFlowCfmTlvsIter
	clearHolderSlice() FlowCfmDmrFlowCfmTlvsIter
	appendHolderSlice(item FlowCfmTlvs) FlowCfmDmrFlowCfmTlvsIter
}

func (obj *flowCfmDmrFlowCfmTlvsIter) setMsg(msg *flowCfmDmr) FlowCfmDmrFlowCfmTlvsIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowCfmTlvs{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowCfmDmrFlowCfmTlvsIter) Items() []FlowCfmTlvs {
	return obj.flowCfmTlvsSlice
}

func (obj *flowCfmDmrFlowCfmTlvsIter) Add() FlowCfmTlvs {
	newObj := &otg.FlowCfmTlvs{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowCfmTlvs{obj: newObj}
	newLibObj.setDefault()
	obj.flowCfmTlvsSlice = append(obj.flowCfmTlvsSlice, newLibObj)
	return newLibObj
}

func (obj *flowCfmDmrFlowCfmTlvsIter) Append(items ...FlowCfmTlvs) FlowCfmDmrFlowCfmTlvsIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowCfmTlvsSlice = append(obj.flowCfmTlvsSlice, item)
	}
	return obj
}

func (obj *flowCfmDmrFlowCfmTlvsIter) Set(index int, newObj FlowCfmTlvs) FlowCfmDmrFlowCfmTlvsIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowCfmTlvsSlice[index] = newObj
	return obj
}
func (obj *flowCfmDmrFlowCfmTlvsIter) Clear() FlowCfmDmrFlowCfmTlvsIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowCfmTlvs{}
		obj.flowCfmTlvsSlice = []FlowCfmTlvs{}
	}
	return obj
}
func (obj *flowCfmDmrFlowCfmTlvsIter) clearHolderSlice() FlowCfmDmrFlowCfmTlvsIter {
	if len(obj.flowCfmTlvsSlice) > 0 {
		obj.flowCfmTlvsSlice = []FlowCfmTlvs{}
	}
	return obj
}
func (obj *flowCfmDmrFlowCfmTlvsIter) appendHolderSlice(item FlowCfmTlvs) FlowCfmDmrFlowCfmTlvsIter {
	obj.flowCfmTlvsSlice = append(obj.flowCfmTlvsSlice, item)
	return obj
}

// When true, the implementation appends an End TLV (Type 0, IEEE 802.1Q-2018 Table 21-17) after all TLVs in the tlvs list. Set to false only when testing DUT behaviour with a deliberately missing End TLV.
// AppendEndTlv returns a bool
func (obj *flowCfmDmr) AppendEndTlv() bool {

	return *obj.obj.AppendEndTlv

}

// When true, the implementation appends an End TLV (Type 0, IEEE 802.1Q-2018 Table 21-17) after all TLVs in the tlvs list. Set to false only when testing DUT behaviour with a deliberately missing End TLV.
// AppendEndTlv returns a bool
func (obj *flowCfmDmr) HasAppendEndTlv() bool {
	return obj.obj.AppendEndTlv != nil
}

// When true, the implementation appends an End TLV (Type 0, IEEE 802.1Q-2018 Table 21-17) after all TLVs in the tlvs list. Set to false only when testing DUT behaviour with a deliberately missing End TLV.
// SetAppendEndTlv sets the bool value in the FlowCfmDmr object
func (obj *flowCfmDmr) SetAppendEndTlv(value bool) FlowCfmDmr {

	obj.obj.AppendEndTlv = &value
	return obj
}

func (obj *flowCfmDmr) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.TxTimestampF != nil {

		obj.TxTimestampF().validateObj(vObj, set_default)
	}

	if obj.obj.RxTimestampF != nil {

		obj.RxTimestampF().validateObj(vObj, set_default)
	}

	if obj.obj.TxTimestampB != nil {

		obj.TxTimestampB().validateObj(vObj, set_default)
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

func (obj *flowCfmDmr) setDefault() {
	if obj.obj.AppendEndTlv == nil {
		obj.SetAppendEndTlv(true)
	}

}
