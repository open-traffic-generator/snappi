package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UltraEthernetCbfcSender *****
type ultraEthernetCbfcSender struct {
	validation
	obj                   *otg.UltraEthernetCbfcSender
	marshaller            marshalUltraEthernetCbfcSender
	unMarshaller          unMarshalUltraEthernetCbfcSender
	portCreditHolder      UltraEthernetCbfcPortCredit
	ccMessageHolder       UltraEthernetCbfcSenderCcMessage
	virtualChannelsHolder UltraEthernetCbfcSenderUltraEthernetCbfcVcIter
}

func NewUltraEthernetCbfcSender() UltraEthernetCbfcSender {
	obj := ultraEthernetCbfcSender{obj: &otg.UltraEthernetCbfcSender{}}
	obj.setDefault()
	return &obj
}

func (obj *ultraEthernetCbfcSender) msg() *otg.UltraEthernetCbfcSender {
	return obj.obj
}

func (obj *ultraEthernetCbfcSender) setMsg(msg *otg.UltraEthernetCbfcSender) UltraEthernetCbfcSender {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalultraEthernetCbfcSender struct {
	obj *ultraEthernetCbfcSender
}

type marshalUltraEthernetCbfcSender interface {
	// ToProto marshals UltraEthernetCbfcSender to protobuf object *otg.UltraEthernetCbfcSender
	ToProto() (*otg.UltraEthernetCbfcSender, error)
	// ToPbText marshals UltraEthernetCbfcSender to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UltraEthernetCbfcSender to YAML text
	ToYaml() (string, error)
	// ToJson marshals UltraEthernetCbfcSender to JSON text
	ToJson() (string, error)
}

type unMarshalultraEthernetCbfcSender struct {
	obj *ultraEthernetCbfcSender
}

type unMarshalUltraEthernetCbfcSender interface {
	// FromProto unmarshals UltraEthernetCbfcSender from protobuf object *otg.UltraEthernetCbfcSender
	FromProto(msg *otg.UltraEthernetCbfcSender) (UltraEthernetCbfcSender, error)
	// FromPbText unmarshals UltraEthernetCbfcSender from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UltraEthernetCbfcSender from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UltraEthernetCbfcSender from JSON text
	FromJson(value string) error
}

func (obj *ultraEthernetCbfcSender) Marshal() marshalUltraEthernetCbfcSender {
	if obj.marshaller == nil {
		obj.marshaller = &marshalultraEthernetCbfcSender{obj: obj}
	}
	return obj.marshaller
}

func (obj *ultraEthernetCbfcSender) Unmarshal() unMarshalUltraEthernetCbfcSender {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalultraEthernetCbfcSender{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalultraEthernetCbfcSender) ToProto() (*otg.UltraEthernetCbfcSender, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalultraEthernetCbfcSender) FromProto(msg *otg.UltraEthernetCbfcSender) (UltraEthernetCbfcSender, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalultraEthernetCbfcSender) ToPbText() (string, error) {
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

func (m *unMarshalultraEthernetCbfcSender) FromPbText(value string) error {
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

func (m *marshalultraEthernetCbfcSender) ToYaml() (string, error) {
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

func (m *unMarshalultraEthernetCbfcSender) FromYaml(value string) error {
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

func (m *marshalultraEthernetCbfcSender) ToJson() (string, error) {
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

func (m *unMarshalultraEthernetCbfcSender) FromJson(value string) error {
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

func (obj *ultraEthernetCbfcSender) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfcSender) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfcSender) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ultraEthernetCbfcSender) Clone() (UltraEthernetCbfcSender, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUltraEthernetCbfcSender()
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

func (obj *ultraEthernetCbfcSender) setNil() {
	obj.portCreditHolder = nil
	obj.ccMessageHolder = nil
	obj.virtualChannelsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// UltraEthernetCbfcSender is cBFC sender direction configuration.
type UltraEthernetCbfcSender interface {
	Validation
	// msg marshals UltraEthernetCbfcSender to protobuf object *otg.UltraEthernetCbfcSender
	// and doesn't set defaults
	msg() *otg.UltraEthernetCbfcSender
	// setMsg unmarshals UltraEthernetCbfcSender from protobuf object *otg.UltraEthernetCbfcSender
	// and doesn't set defaults
	setMsg(*otg.UltraEthernetCbfcSender) UltraEthernetCbfcSender
	// provides marshal interface
	Marshal() marshalUltraEthernetCbfcSender
	// provides unmarshal interface
	Unmarshal() unMarshalUltraEthernetCbfcSender
	// validate validates UltraEthernetCbfcSender
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UltraEthernetCbfcSender, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Enable returns bool, set in UltraEthernetCbfcSender.
	Enable() bool
	// SetEnable assigns bool provided by user to UltraEthernetCbfcSender
	SetEnable(value bool) UltraEthernetCbfcSender
	// HasEnable checks if Enable has been set in UltraEthernetCbfcSender
	HasEnable() bool
	// PortCredit returns UltraEthernetCbfcPortCredit, set in UltraEthernetCbfcSender.
	// UltraEthernetCbfcPortCredit is cBFC port level credit configuration. Credits are the unit used to track
	// available buffer space at the receiver.
	//
	// Reference: UE-Specification-1.0.3 Table 5-14.
	PortCredit() UltraEthernetCbfcPortCredit
	// SetPortCredit assigns UltraEthernetCbfcPortCredit provided by user to UltraEthernetCbfcSender.
	// UltraEthernetCbfcPortCredit is cBFC port level credit configuration. Credits are the unit used to track
	// available buffer space at the receiver.
	//
	// Reference: UE-Specification-1.0.3 Table 5-14.
	SetPortCredit(value UltraEthernetCbfcPortCredit) UltraEthernetCbfcSender
	// HasPortCredit checks if PortCredit has been set in UltraEthernetCbfcSender
	HasPortCredit() bool
	// CcMessage returns UltraEthernetCbfcSenderCcMessage, set in UltraEthernetCbfcSender.
	// UltraEthernetCbfcSenderCcMessage is cBFC CC_Update (credit consumed) message configuration. CC_Update messages are
	// periodically generated by the sender to recover credits that may have leaked
	// due to packet drops from link errors.
	CcMessage() UltraEthernetCbfcSenderCcMessage
	// SetCcMessage assigns UltraEthernetCbfcSenderCcMessage provided by user to UltraEthernetCbfcSender.
	// UltraEthernetCbfcSenderCcMessage is cBFC CC_Update (credit consumed) message configuration. CC_Update messages are
	// periodically generated by the sender to recover credits that may have leaked
	// due to packet drops from link errors.
	SetCcMessage(value UltraEthernetCbfcSenderCcMessage) UltraEthernetCbfcSender
	// HasCcMessage checks if CcMessage has been set in UltraEthernetCbfcSender
	HasCcMessage() bool
	// VirtualChannels returns UltraEthernetCbfcSenderUltraEthernetCbfcVcIterIter, set in UltraEthernetCbfcSender
	VirtualChannels() UltraEthernetCbfcSenderUltraEthernetCbfcVcIter
	setNil()
}

// Enable the CBFC sender on the port.
// Enable returns a bool
func (obj *ultraEthernetCbfcSender) Enable() bool {

	return *obj.obj.Enable

}

// Enable the CBFC sender on the port.
// Enable returns a bool
func (obj *ultraEthernetCbfcSender) HasEnable() bool {
	return obj.obj.Enable != nil
}

// Enable the CBFC sender on the port.
// SetEnable sets the bool value in the UltraEthernetCbfcSender object
func (obj *ultraEthernetCbfcSender) SetEnable(value bool) UltraEthernetCbfcSender {

	obj.obj.Enable = &value
	return obj
}

// Port level credit configuration (credit mode, total credits, cell size and
// per-packet overhead).
// PortCredit returns a UltraEthernetCbfcPortCredit
func (obj *ultraEthernetCbfcSender) PortCredit() UltraEthernetCbfcPortCredit {
	if obj.obj.PortCredit == nil {
		obj.obj.PortCredit = NewUltraEthernetCbfcPortCredit().msg()
	}
	if obj.portCreditHolder == nil {
		obj.portCreditHolder = &ultraEthernetCbfcPortCredit{obj: obj.obj.PortCredit}
	}
	return obj.portCreditHolder
}

// Port level credit configuration (credit mode, total credits, cell size and
// per-packet overhead).
// PortCredit returns a UltraEthernetCbfcPortCredit
func (obj *ultraEthernetCbfcSender) HasPortCredit() bool {
	return obj.obj.PortCredit != nil
}

// Port level credit configuration (credit mode, total credits, cell size and
// per-packet overhead).
// SetPortCredit sets the UltraEthernetCbfcPortCredit value in the UltraEthernetCbfcSender object
func (obj *ultraEthernetCbfcSender) SetPortCredit(value UltraEthernetCbfcPortCredit) UltraEthernetCbfcSender {

	obj.portCreditHolder = nil
	obj.obj.PortCredit = value.msg()

	return obj
}

// CC_Update (credit consumed) message configuration.
// CcMessage returns a UltraEthernetCbfcSenderCcMessage
func (obj *ultraEthernetCbfcSender) CcMessage() UltraEthernetCbfcSenderCcMessage {
	if obj.obj.CcMessage == nil {
		obj.obj.CcMessage = NewUltraEthernetCbfcSenderCcMessage().msg()
	}
	if obj.ccMessageHolder == nil {
		obj.ccMessageHolder = &ultraEthernetCbfcSenderCcMessage{obj: obj.obj.CcMessage}
	}
	return obj.ccMessageHolder
}

// CC_Update (credit consumed) message configuration.
// CcMessage returns a UltraEthernetCbfcSenderCcMessage
func (obj *ultraEthernetCbfcSender) HasCcMessage() bool {
	return obj.obj.CcMessage != nil
}

// CC_Update (credit consumed) message configuration.
// SetCcMessage sets the UltraEthernetCbfcSenderCcMessage value in the UltraEthernetCbfcSender object
func (obj *ultraEthernetCbfcSender) SetCcMessage(value UltraEthernetCbfcSenderCcMessage) UltraEthernetCbfcSender {

	obj.ccMessageHolder = nil
	obj.obj.CcMessage = value.msg()

	return obj
}

// The list of virtual channels (VCs). A maximum of 32 VCs are supported per
// port, of which a maximum of 4 may be configured as lossless.
// VirtualChannels returns a []UltraEthernetCbfcVc
func (obj *ultraEthernetCbfcSender) VirtualChannels() UltraEthernetCbfcSenderUltraEthernetCbfcVcIter {
	if len(obj.obj.VirtualChannels) == 0 {
		obj.obj.VirtualChannels = []*otg.UltraEthernetCbfcVc{}
	}
	if obj.virtualChannelsHolder == nil {
		obj.virtualChannelsHolder = newUltraEthernetCbfcSenderUltraEthernetCbfcVcIter(&obj.obj.VirtualChannels).setMsg(obj)
	}
	return obj.virtualChannelsHolder
}

type ultraEthernetCbfcSenderUltraEthernetCbfcVcIter struct {
	obj                      *ultraEthernetCbfcSender
	ultraEthernetCbfcVcSlice []UltraEthernetCbfcVc
	fieldPtr                 *[]*otg.UltraEthernetCbfcVc
}

func newUltraEthernetCbfcSenderUltraEthernetCbfcVcIter(ptr *[]*otg.UltraEthernetCbfcVc) UltraEthernetCbfcSenderUltraEthernetCbfcVcIter {
	return &ultraEthernetCbfcSenderUltraEthernetCbfcVcIter{fieldPtr: ptr}
}

type UltraEthernetCbfcSenderUltraEthernetCbfcVcIter interface {
	setMsg(*ultraEthernetCbfcSender) UltraEthernetCbfcSenderUltraEthernetCbfcVcIter
	Items() []UltraEthernetCbfcVc
	Add() UltraEthernetCbfcVc
	Append(items ...UltraEthernetCbfcVc) UltraEthernetCbfcSenderUltraEthernetCbfcVcIter
	Set(index int, newObj UltraEthernetCbfcVc) UltraEthernetCbfcSenderUltraEthernetCbfcVcIter
	Clear() UltraEthernetCbfcSenderUltraEthernetCbfcVcIter
	clearHolderSlice() UltraEthernetCbfcSenderUltraEthernetCbfcVcIter
	appendHolderSlice(item UltraEthernetCbfcVc) UltraEthernetCbfcSenderUltraEthernetCbfcVcIter
}

func (obj *ultraEthernetCbfcSenderUltraEthernetCbfcVcIter) setMsg(msg *ultraEthernetCbfcSender) UltraEthernetCbfcSenderUltraEthernetCbfcVcIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ultraEthernetCbfcVc{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ultraEthernetCbfcSenderUltraEthernetCbfcVcIter) Items() []UltraEthernetCbfcVc {
	return obj.ultraEthernetCbfcVcSlice
}

func (obj *ultraEthernetCbfcSenderUltraEthernetCbfcVcIter) Add() UltraEthernetCbfcVc {
	newObj := &otg.UltraEthernetCbfcVc{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ultraEthernetCbfcVc{obj: newObj}
	newLibObj.setDefault()
	obj.ultraEthernetCbfcVcSlice = append(obj.ultraEthernetCbfcVcSlice, newLibObj)
	return newLibObj
}

func (obj *ultraEthernetCbfcSenderUltraEthernetCbfcVcIter) Append(items ...UltraEthernetCbfcVc) UltraEthernetCbfcSenderUltraEthernetCbfcVcIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ultraEthernetCbfcVcSlice = append(obj.ultraEthernetCbfcVcSlice, item)
	}
	return obj
}

func (obj *ultraEthernetCbfcSenderUltraEthernetCbfcVcIter) Set(index int, newObj UltraEthernetCbfcVc) UltraEthernetCbfcSenderUltraEthernetCbfcVcIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ultraEthernetCbfcVcSlice[index] = newObj
	return obj
}
func (obj *ultraEthernetCbfcSenderUltraEthernetCbfcVcIter) Clear() UltraEthernetCbfcSenderUltraEthernetCbfcVcIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.UltraEthernetCbfcVc{}
		obj.ultraEthernetCbfcVcSlice = []UltraEthernetCbfcVc{}
	}
	return obj
}
func (obj *ultraEthernetCbfcSenderUltraEthernetCbfcVcIter) clearHolderSlice() UltraEthernetCbfcSenderUltraEthernetCbfcVcIter {
	if len(obj.ultraEthernetCbfcVcSlice) > 0 {
		obj.ultraEthernetCbfcVcSlice = []UltraEthernetCbfcVc{}
	}
	return obj
}
func (obj *ultraEthernetCbfcSenderUltraEthernetCbfcVcIter) appendHolderSlice(item UltraEthernetCbfcVc) UltraEthernetCbfcSenderUltraEthernetCbfcVcIter {
	obj.ultraEthernetCbfcVcSlice = append(obj.ultraEthernetCbfcVcSlice, item)
	return obj
}

func (obj *ultraEthernetCbfcSender) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PortCredit != nil {

		obj.PortCredit().validateObj(vObj, set_default)
	}

	if obj.obj.CcMessage != nil {

		obj.CcMessage().validateObj(vObj, set_default)
	}

	if len(obj.obj.VirtualChannels) != 0 {

		if set_default {
			obj.VirtualChannels().clearHolderSlice()
			for _, item := range obj.obj.VirtualChannels {
				obj.VirtualChannels().appendHolderSlice(&ultraEthernetCbfcVc{obj: item})
			}
		}
		for _, item := range obj.VirtualChannels().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *ultraEthernetCbfcSender) setDefault() {
	if obj.obj.Enable == nil {
		obj.SetEnable(false)
	}

}
