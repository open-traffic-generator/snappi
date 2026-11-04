package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfmLbm *****
type flowCfmLbm struct {
	validation
	obj                 *otg.FlowCfmLbm
	marshaller          marshalFlowCfmLbm
	unMarshaller        unMarshalFlowCfmLbm
	transactionIdHolder PatternFlowCfmLbmTransactionId
	tlvsHolder          FlowCfmLbmFlowCfmTlvsIter
}

func NewFlowCfmLbm() FlowCfmLbm {
	obj := flowCfmLbm{obj: &otg.FlowCfmLbm{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfmLbm) msg() *otg.FlowCfmLbm {
	return obj.obj
}

func (obj *flowCfmLbm) setMsg(msg *otg.FlowCfmLbm) FlowCfmLbm {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfmLbm struct {
	obj *flowCfmLbm
}

type marshalFlowCfmLbm interface {
	// ToProto marshals FlowCfmLbm to protobuf object *otg.FlowCfmLbm
	ToProto() (*otg.FlowCfmLbm, error)
	// ToPbText marshals FlowCfmLbm to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfmLbm to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfmLbm to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfmLbm struct {
	obj *flowCfmLbm
}

type unMarshalFlowCfmLbm interface {
	// FromProto unmarshals FlowCfmLbm from protobuf object *otg.FlowCfmLbm
	FromProto(msg *otg.FlowCfmLbm) (FlowCfmLbm, error)
	// FromPbText unmarshals FlowCfmLbm from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfmLbm from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfmLbm from JSON text
	FromJson(value string) error
}

func (obj *flowCfmLbm) Marshal() marshalFlowCfmLbm {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfmLbm{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfmLbm) Unmarshal() unMarshalFlowCfmLbm {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfmLbm{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfmLbm) ToProto() (*otg.FlowCfmLbm, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfmLbm) FromProto(msg *otg.FlowCfmLbm) (FlowCfmLbm, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfmLbm) ToPbText() (string, error) {
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

func (m *unMarshalflowCfmLbm) FromPbText(value string) error {
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

func (m *marshalflowCfmLbm) ToYaml() (string, error) {
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

func (m *unMarshalflowCfmLbm) FromYaml(value string) error {
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

func (m *marshalflowCfmLbm) ToJson() (string, error) {
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

func (m *unMarshalflowCfmLbm) FromJson(value string) error {
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

func (obj *flowCfmLbm) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfmLbm) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfmLbm) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfmLbm) Clone() (FlowCfmLbm, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfmLbm()
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

func (obj *flowCfmLbm) setNil() {
	obj.transactionIdHolder = nil
	obj.tlvsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowCfmLbm is loopback Message (LBM) is a diagnostic frame used to verify connectivity between Maintenance End Points (MEPs) and Maintenance Intermediate Points (MIPs). LBM flag bits are reserved and set to zero.
type FlowCfmLbm interface {
	Validation
	// msg marshals FlowCfmLbm to protobuf object *otg.FlowCfmLbm
	// and doesn't set defaults
	msg() *otg.FlowCfmLbm
	// setMsg unmarshals FlowCfmLbm from protobuf object *otg.FlowCfmLbm
	// and doesn't set defaults
	setMsg(*otg.FlowCfmLbm) FlowCfmLbm
	// provides marshal interface
	Marshal() marshalFlowCfmLbm
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfmLbm
	// validate validates FlowCfmLbm
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfmLbm, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// TransactionId returns PatternFlowCfmLbmTransactionId, set in FlowCfmLbm.
	// PatternFlowCfmLbmTransactionId is transaction identifier copied verbatim into the corresponding LBR, allowing the sender to match the reply to this request (IEEE 802.1Q-2018 Figure 21-5).
	TransactionId() PatternFlowCfmLbmTransactionId
	// SetTransactionId assigns PatternFlowCfmLbmTransactionId provided by user to FlowCfmLbm.
	// PatternFlowCfmLbmTransactionId is transaction identifier copied verbatim into the corresponding LBR, allowing the sender to match the reply to this request (IEEE 802.1Q-2018 Figure 21-5).
	SetTransactionId(value PatternFlowCfmLbmTransactionId) FlowCfmLbm
	// HasTransactionId checks if TransactionId has been set in FlowCfmLbm
	HasTransactionId() bool
	// Tlvs returns FlowCfmLbmFlowCfmTlvsIterIter, set in FlowCfmLbm
	Tlvs() FlowCfmLbmFlowCfmTlvsIter
	// AppendEndTlv returns bool, set in FlowCfmLbm.
	AppendEndTlv() bool
	// SetAppendEndTlv assigns bool provided by user to FlowCfmLbm
	SetAppendEndTlv(value bool) FlowCfmLbm
	// HasAppendEndTlv checks if AppendEndTlv has been set in FlowCfmLbm
	HasAppendEndTlv() bool
	setNil()
}

// description is TBD
// TransactionId returns a PatternFlowCfmLbmTransactionId
func (obj *flowCfmLbm) TransactionId() PatternFlowCfmLbmTransactionId {
	if obj.obj.TransactionId == nil {
		obj.obj.TransactionId = NewPatternFlowCfmLbmTransactionId().msg()
	}
	if obj.transactionIdHolder == nil {
		obj.transactionIdHolder = &patternFlowCfmLbmTransactionId{obj: obj.obj.TransactionId}
	}
	return obj.transactionIdHolder
}

// description is TBD
// TransactionId returns a PatternFlowCfmLbmTransactionId
func (obj *flowCfmLbm) HasTransactionId() bool {
	return obj.obj.TransactionId != nil
}

// description is TBD
// SetTransactionId sets the PatternFlowCfmLbmTransactionId value in the FlowCfmLbm object
func (obj *flowCfmLbm) SetTransactionId(value PatternFlowCfmLbmTransactionId) FlowCfmLbm {

	obj.transactionIdHolder = nil
	obj.obj.TransactionId = value.msg()

	return obj
}

// description is TBD
// Tlvs returns a []FlowCfmTlvs
func (obj *flowCfmLbm) Tlvs() FlowCfmLbmFlowCfmTlvsIter {
	if len(obj.obj.Tlvs) == 0 {
		obj.obj.Tlvs = []*otg.FlowCfmTlvs{}
	}
	if obj.tlvsHolder == nil {
		obj.tlvsHolder = newFlowCfmLbmFlowCfmTlvsIter(&obj.obj.Tlvs).setMsg(obj)
	}
	return obj.tlvsHolder
}

type flowCfmLbmFlowCfmTlvsIter struct {
	obj              *flowCfmLbm
	flowCfmTlvsSlice []FlowCfmTlvs
	fieldPtr         *[]*otg.FlowCfmTlvs
}

func newFlowCfmLbmFlowCfmTlvsIter(ptr *[]*otg.FlowCfmTlvs) FlowCfmLbmFlowCfmTlvsIter {
	return &flowCfmLbmFlowCfmTlvsIter{fieldPtr: ptr}
}

type FlowCfmLbmFlowCfmTlvsIter interface {
	setMsg(*flowCfmLbm) FlowCfmLbmFlowCfmTlvsIter
	Items() []FlowCfmTlvs
	Add() FlowCfmTlvs
	Append(items ...FlowCfmTlvs) FlowCfmLbmFlowCfmTlvsIter
	Set(index int, newObj FlowCfmTlvs) FlowCfmLbmFlowCfmTlvsIter
	Clear() FlowCfmLbmFlowCfmTlvsIter
	clearHolderSlice() FlowCfmLbmFlowCfmTlvsIter
	appendHolderSlice(item FlowCfmTlvs) FlowCfmLbmFlowCfmTlvsIter
}

func (obj *flowCfmLbmFlowCfmTlvsIter) setMsg(msg *flowCfmLbm) FlowCfmLbmFlowCfmTlvsIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowCfmTlvs{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowCfmLbmFlowCfmTlvsIter) Items() []FlowCfmTlvs {
	return obj.flowCfmTlvsSlice
}

func (obj *flowCfmLbmFlowCfmTlvsIter) Add() FlowCfmTlvs {
	newObj := &otg.FlowCfmTlvs{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowCfmTlvs{obj: newObj}
	newLibObj.setDefault()
	obj.flowCfmTlvsSlice = append(obj.flowCfmTlvsSlice, newLibObj)
	return newLibObj
}

func (obj *flowCfmLbmFlowCfmTlvsIter) Append(items ...FlowCfmTlvs) FlowCfmLbmFlowCfmTlvsIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowCfmTlvsSlice = append(obj.flowCfmTlvsSlice, item)
	}
	return obj
}

func (obj *flowCfmLbmFlowCfmTlvsIter) Set(index int, newObj FlowCfmTlvs) FlowCfmLbmFlowCfmTlvsIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowCfmTlvsSlice[index] = newObj
	return obj
}
func (obj *flowCfmLbmFlowCfmTlvsIter) Clear() FlowCfmLbmFlowCfmTlvsIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowCfmTlvs{}
		obj.flowCfmTlvsSlice = []FlowCfmTlvs{}
	}
	return obj
}
func (obj *flowCfmLbmFlowCfmTlvsIter) clearHolderSlice() FlowCfmLbmFlowCfmTlvsIter {
	if len(obj.flowCfmTlvsSlice) > 0 {
		obj.flowCfmTlvsSlice = []FlowCfmTlvs{}
	}
	return obj
}
func (obj *flowCfmLbmFlowCfmTlvsIter) appendHolderSlice(item FlowCfmTlvs) FlowCfmLbmFlowCfmTlvsIter {
	obj.flowCfmTlvsSlice = append(obj.flowCfmTlvsSlice, item)
	return obj
}

// When true, the implementation appends an End TLV (Type 0, IEEE 802.1Q-2018 Table 21-17) after all TLVs in the tlvs list. Set to false only when testing DUT behaviour with a deliberately missing End TLV.
// AppendEndTlv returns a bool
func (obj *flowCfmLbm) AppendEndTlv() bool {

	return *obj.obj.AppendEndTlv

}

// When true, the implementation appends an End TLV (Type 0, IEEE 802.1Q-2018 Table 21-17) after all TLVs in the tlvs list. Set to false only when testing DUT behaviour with a deliberately missing End TLV.
// AppendEndTlv returns a bool
func (obj *flowCfmLbm) HasAppendEndTlv() bool {
	return obj.obj.AppendEndTlv != nil
}

// When true, the implementation appends an End TLV (Type 0, IEEE 802.1Q-2018 Table 21-17) after all TLVs in the tlvs list. Set to false only when testing DUT behaviour with a deliberately missing End TLV.
// SetAppendEndTlv sets the bool value in the FlowCfmLbm object
func (obj *flowCfmLbm) SetAppendEndTlv(value bool) FlowCfmLbm {

	obj.obj.AppendEndTlv = &value
	return obj
}

func (obj *flowCfmLbm) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
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

func (obj *flowCfmLbm) setDefault() {
	if obj.obj.AppendEndTlv == nil {
		obj.SetAppendEndTlv(true)
	}

}
