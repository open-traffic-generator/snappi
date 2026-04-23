package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfmCcm *****
type flowCfmCcm struct {
	validation
	obj                        *otg.FlowCfmCcm
	marshaller                 marshalFlowCfmCcm
	unMarshaller               unMarshalFlowCfmCcm
	flagsHolder                PatternFlowCfmCcmFlags
	sequenceNumberHolder       PatternFlowCfmCcmSequenceNumber
	maEndpointIdentifierHolder PatternFlowCfmCcmMaEndpointIdentifier
	ethernetOamProtocolHolder  FlowCfmOamProtocol
	tlvsHolder                 FlowCfmCcmFlowCfmTlvsIter
}

func NewFlowCfmCcm() FlowCfmCcm {
	obj := flowCfmCcm{obj: &otg.FlowCfmCcm{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfmCcm) msg() *otg.FlowCfmCcm {
	return obj.obj
}

func (obj *flowCfmCcm) setMsg(msg *otg.FlowCfmCcm) FlowCfmCcm {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfmCcm struct {
	obj *flowCfmCcm
}

type marshalFlowCfmCcm interface {
	// ToProto marshals FlowCfmCcm to protobuf object *otg.FlowCfmCcm
	ToProto() (*otg.FlowCfmCcm, error)
	// ToPbText marshals FlowCfmCcm to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfmCcm to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfmCcm to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfmCcm struct {
	obj *flowCfmCcm
}

type unMarshalFlowCfmCcm interface {
	// FromProto unmarshals FlowCfmCcm from protobuf object *otg.FlowCfmCcm
	FromProto(msg *otg.FlowCfmCcm) (FlowCfmCcm, error)
	// FromPbText unmarshals FlowCfmCcm from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfmCcm from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfmCcm from JSON text
	FromJson(value string) error
}

func (obj *flowCfmCcm) Marshal() marshalFlowCfmCcm {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfmCcm{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfmCcm) Unmarshal() unMarshalFlowCfmCcm {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfmCcm{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfmCcm) ToProto() (*otg.FlowCfmCcm, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfmCcm) FromProto(msg *otg.FlowCfmCcm) (FlowCfmCcm, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfmCcm) ToPbText() (string, error) {
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

func (m *unMarshalflowCfmCcm) FromPbText(value string) error {
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

func (m *marshalflowCfmCcm) ToYaml() (string, error) {
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

func (m *unMarshalflowCfmCcm) FromYaml(value string) error {
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

func (m *marshalflowCfmCcm) ToJson() (string, error) {
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

func (m *unMarshalflowCfmCcm) FromJson(value string) error {
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

func (obj *flowCfmCcm) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfmCcm) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfmCcm) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfmCcm) Clone() (FlowCfmCcm, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfmCcm()
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

func (obj *flowCfmCcm) setNil() {
	obj.flagsHolder = nil
	obj.sequenceNumberHolder = nil
	obj.maEndpointIdentifierHolder = nil
	obj.ethernetOamProtocolHolder = nil
	obj.tlvsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowCfmCcm is cCM is used to monitor connectivity and configuration errors. Each CCM message has a unique sequence number (Session ID) and unique flow-identifier.
type FlowCfmCcm interface {
	Validation
	// msg marshals FlowCfmCcm to protobuf object *otg.FlowCfmCcm
	// and doesn't set defaults
	msg() *otg.FlowCfmCcm
	// setMsg unmarshals FlowCfmCcm from protobuf object *otg.FlowCfmCcm
	// and doesn't set defaults
	setMsg(*otg.FlowCfmCcm) FlowCfmCcm
	// provides marshal interface
	Marshal() marshalFlowCfmCcm
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfmCcm
	// validate validates FlowCfmCcm
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfmCcm, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns PatternFlowCfmCcmFlags, set in FlowCfmCcm.
	// PatternFlowCfmCcmFlags is defines operational flags for CCM.
	Flags() PatternFlowCfmCcmFlags
	// SetFlags assigns PatternFlowCfmCcmFlags provided by user to FlowCfmCcm.
	// PatternFlowCfmCcmFlags is defines operational flags for CCM.
	SetFlags(value PatternFlowCfmCcmFlags) FlowCfmCcm
	// HasFlags checks if Flags has been set in FlowCfmCcm
	HasFlags() bool
	// SequenceNumber returns PatternFlowCfmCcmSequenceNumber, set in FlowCfmCcm.
	// PatternFlowCfmCcmSequenceNumber is cCM unique sequence number
	SequenceNumber() PatternFlowCfmCcmSequenceNumber
	// SetSequenceNumber assigns PatternFlowCfmCcmSequenceNumber provided by user to FlowCfmCcm.
	// PatternFlowCfmCcmSequenceNumber is cCM unique sequence number
	SetSequenceNumber(value PatternFlowCfmCcmSequenceNumber) FlowCfmCcm
	// HasSequenceNumber checks if SequenceNumber has been set in FlowCfmCcm
	HasSequenceNumber() bool
	// MaEndpointIdentifier returns PatternFlowCfmCcmMaEndpointIdentifier, set in FlowCfmCcm.
	// PatternFlowCfmCcmMaEndpointIdentifier is maintenance association endPoint identifier
	MaEndpointIdentifier() PatternFlowCfmCcmMaEndpointIdentifier
	// SetMaEndpointIdentifier assigns PatternFlowCfmCcmMaEndpointIdentifier provided by user to FlowCfmCcm.
	// PatternFlowCfmCcmMaEndpointIdentifier is maintenance association endPoint identifier
	SetMaEndpointIdentifier(value PatternFlowCfmCcmMaEndpointIdentifier) FlowCfmCcm
	// HasMaEndpointIdentifier checks if MaEndpointIdentifier has been set in FlowCfmCcm
	HasMaEndpointIdentifier() bool
	// EthernetOamProtocol returns FlowCfmOamProtocol, set in FlowCfmCcm.
	// FlowCfmOamProtocol is ethernet OAM standard. Currently only Connectivity Fault Management (CFM) is supported.
	EthernetOamProtocol() FlowCfmOamProtocol
	// SetEthernetOamProtocol assigns FlowCfmOamProtocol provided by user to FlowCfmCcm.
	// FlowCfmOamProtocol is ethernet OAM standard. Currently only Connectivity Fault Management (CFM) is supported.
	SetEthernetOamProtocol(value FlowCfmOamProtocol) FlowCfmCcm
	// HasEthernetOamProtocol checks if EthernetOamProtocol has been set in FlowCfmCcm
	HasEthernetOamProtocol() bool
	// Tlvs returns FlowCfmCcmFlowCfmTlvsIterIter, set in FlowCfmCcm
	Tlvs() FlowCfmCcmFlowCfmTlvsIter
	setNil()
}

// description is TBD
// Flags returns a PatternFlowCfmCcmFlags
func (obj *flowCfmCcm) Flags() PatternFlowCfmCcmFlags {
	if obj.obj.Flags == nil {
		obj.obj.Flags = NewPatternFlowCfmCcmFlags().msg()
	}
	if obj.flagsHolder == nil {
		obj.flagsHolder = &patternFlowCfmCcmFlags{obj: obj.obj.Flags}
	}
	return obj.flagsHolder
}

// description is TBD
// Flags returns a PatternFlowCfmCcmFlags
func (obj *flowCfmCcm) HasFlags() bool {
	return obj.obj.Flags != nil
}

// description is TBD
// SetFlags sets the PatternFlowCfmCcmFlags value in the FlowCfmCcm object
func (obj *flowCfmCcm) SetFlags(value PatternFlowCfmCcmFlags) FlowCfmCcm {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// description is TBD
// SequenceNumber returns a PatternFlowCfmCcmSequenceNumber
func (obj *flowCfmCcm) SequenceNumber() PatternFlowCfmCcmSequenceNumber {
	if obj.obj.SequenceNumber == nil {
		obj.obj.SequenceNumber = NewPatternFlowCfmCcmSequenceNumber().msg()
	}
	if obj.sequenceNumberHolder == nil {
		obj.sequenceNumberHolder = &patternFlowCfmCcmSequenceNumber{obj: obj.obj.SequenceNumber}
	}
	return obj.sequenceNumberHolder
}

// description is TBD
// SequenceNumber returns a PatternFlowCfmCcmSequenceNumber
func (obj *flowCfmCcm) HasSequenceNumber() bool {
	return obj.obj.SequenceNumber != nil
}

// description is TBD
// SetSequenceNumber sets the PatternFlowCfmCcmSequenceNumber value in the FlowCfmCcm object
func (obj *flowCfmCcm) SetSequenceNumber(value PatternFlowCfmCcmSequenceNumber) FlowCfmCcm {

	obj.sequenceNumberHolder = nil
	obj.obj.SequenceNumber = value.msg()

	return obj
}

// description is TBD
// MaEndpointIdentifier returns a PatternFlowCfmCcmMaEndpointIdentifier
func (obj *flowCfmCcm) MaEndpointIdentifier() PatternFlowCfmCcmMaEndpointIdentifier {
	if obj.obj.MaEndpointIdentifier == nil {
		obj.obj.MaEndpointIdentifier = NewPatternFlowCfmCcmMaEndpointIdentifier().msg()
	}
	if obj.maEndpointIdentifierHolder == nil {
		obj.maEndpointIdentifierHolder = &patternFlowCfmCcmMaEndpointIdentifier{obj: obj.obj.MaEndpointIdentifier}
	}
	return obj.maEndpointIdentifierHolder
}

// description is TBD
// MaEndpointIdentifier returns a PatternFlowCfmCcmMaEndpointIdentifier
func (obj *flowCfmCcm) HasMaEndpointIdentifier() bool {
	return obj.obj.MaEndpointIdentifier != nil
}

// description is TBD
// SetMaEndpointIdentifier sets the PatternFlowCfmCcmMaEndpointIdentifier value in the FlowCfmCcm object
func (obj *flowCfmCcm) SetMaEndpointIdentifier(value PatternFlowCfmCcmMaEndpointIdentifier) FlowCfmCcm {

	obj.maEndpointIdentifierHolder = nil
	obj.obj.MaEndpointIdentifier = value.msg()

	return obj
}

// description is TBD
// EthernetOamProtocol returns a FlowCfmOamProtocol
func (obj *flowCfmCcm) EthernetOamProtocol() FlowCfmOamProtocol {
	if obj.obj.EthernetOamProtocol == nil {
		obj.obj.EthernetOamProtocol = NewFlowCfmOamProtocol().msg()
	}
	if obj.ethernetOamProtocolHolder == nil {
		obj.ethernetOamProtocolHolder = &flowCfmOamProtocol{obj: obj.obj.EthernetOamProtocol}
	}
	return obj.ethernetOamProtocolHolder
}

// description is TBD
// EthernetOamProtocol returns a FlowCfmOamProtocol
func (obj *flowCfmCcm) HasEthernetOamProtocol() bool {
	return obj.obj.EthernetOamProtocol != nil
}

// description is TBD
// SetEthernetOamProtocol sets the FlowCfmOamProtocol value in the FlowCfmCcm object
func (obj *flowCfmCcm) SetEthernetOamProtocol(value FlowCfmOamProtocol) FlowCfmCcm {

	obj.ethernetOamProtocolHolder = nil
	obj.obj.EthernetOamProtocol = value.msg()

	return obj
}

// description is TBD
// Tlvs returns a []FlowCfmTlvs
func (obj *flowCfmCcm) Tlvs() FlowCfmCcmFlowCfmTlvsIter {
	if len(obj.obj.Tlvs) == 0 {
		obj.obj.Tlvs = []*otg.FlowCfmTlvs{}
	}
	if obj.tlvsHolder == nil {
		obj.tlvsHolder = newFlowCfmCcmFlowCfmTlvsIter(&obj.obj.Tlvs).setMsg(obj)
	}
	return obj.tlvsHolder
}

type flowCfmCcmFlowCfmTlvsIter struct {
	obj              *flowCfmCcm
	flowCfmTlvsSlice []FlowCfmTlvs
	fieldPtr         *[]*otg.FlowCfmTlvs
}

func newFlowCfmCcmFlowCfmTlvsIter(ptr *[]*otg.FlowCfmTlvs) FlowCfmCcmFlowCfmTlvsIter {
	return &flowCfmCcmFlowCfmTlvsIter{fieldPtr: ptr}
}

type FlowCfmCcmFlowCfmTlvsIter interface {
	setMsg(*flowCfmCcm) FlowCfmCcmFlowCfmTlvsIter
	Items() []FlowCfmTlvs
	Add() FlowCfmTlvs
	Append(items ...FlowCfmTlvs) FlowCfmCcmFlowCfmTlvsIter
	Set(index int, newObj FlowCfmTlvs) FlowCfmCcmFlowCfmTlvsIter
	Clear() FlowCfmCcmFlowCfmTlvsIter
	clearHolderSlice() FlowCfmCcmFlowCfmTlvsIter
	appendHolderSlice(item FlowCfmTlvs) FlowCfmCcmFlowCfmTlvsIter
}

func (obj *flowCfmCcmFlowCfmTlvsIter) setMsg(msg *flowCfmCcm) FlowCfmCcmFlowCfmTlvsIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowCfmTlvs{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowCfmCcmFlowCfmTlvsIter) Items() []FlowCfmTlvs {
	return obj.flowCfmTlvsSlice
}

func (obj *flowCfmCcmFlowCfmTlvsIter) Add() FlowCfmTlvs {
	newObj := &otg.FlowCfmTlvs{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowCfmTlvs{obj: newObj}
	newLibObj.setDefault()
	obj.flowCfmTlvsSlice = append(obj.flowCfmTlvsSlice, newLibObj)
	return newLibObj
}

func (obj *flowCfmCcmFlowCfmTlvsIter) Append(items ...FlowCfmTlvs) FlowCfmCcmFlowCfmTlvsIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowCfmTlvsSlice = append(obj.flowCfmTlvsSlice, item)
	}
	return obj
}

func (obj *flowCfmCcmFlowCfmTlvsIter) Set(index int, newObj FlowCfmTlvs) FlowCfmCcmFlowCfmTlvsIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowCfmTlvsSlice[index] = newObj
	return obj
}
func (obj *flowCfmCcmFlowCfmTlvsIter) Clear() FlowCfmCcmFlowCfmTlvsIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowCfmTlvs{}
		obj.flowCfmTlvsSlice = []FlowCfmTlvs{}
	}
	return obj
}
func (obj *flowCfmCcmFlowCfmTlvsIter) clearHolderSlice() FlowCfmCcmFlowCfmTlvsIter {
	if len(obj.flowCfmTlvsSlice) > 0 {
		obj.flowCfmTlvsSlice = []FlowCfmTlvs{}
	}
	return obj
}
func (obj *flowCfmCcmFlowCfmTlvsIter) appendHolderSlice(item FlowCfmTlvs) FlowCfmCcmFlowCfmTlvsIter {
	obj.flowCfmTlvsSlice = append(obj.flowCfmTlvsSlice, item)
	return obj
}

func (obj *flowCfmCcm) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(vObj, set_default)
	}

	if obj.obj.SequenceNumber != nil {

		obj.SequenceNumber().validateObj(vObj, set_default)
	}

	if obj.obj.MaEndpointIdentifier != nil {

		obj.MaEndpointIdentifier().validateObj(vObj, set_default)
	}

	if obj.obj.EthernetOamProtocol != nil {

		obj.EthernetOamProtocol().validateObj(vObj, set_default)
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

func (obj *flowCfmCcm) setDefault() {

}
