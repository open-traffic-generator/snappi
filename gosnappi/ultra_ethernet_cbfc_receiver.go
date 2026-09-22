package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UltraEthernetCbfcReceiver *****
type ultraEthernetCbfcReceiver struct {
	validation
	obj                   *otg.UltraEthernetCbfcReceiver
	marshaller            marshalUltraEthernetCbfcReceiver
	unMarshaller          unMarshalUltraEthernetCbfcReceiver
	portCreditHolder      UltraEthernetCbfcPortCredit
	ccMessageHolder       UltraEthernetCbfcReceiverCcMessage
	virtualChannelsHolder UltraEthernetCbfcReceiverUltraEthernetCbfcVcIter
}

func NewUltraEthernetCbfcReceiver() UltraEthernetCbfcReceiver {
	obj := ultraEthernetCbfcReceiver{obj: &otg.UltraEthernetCbfcReceiver{}}
	obj.setDefault()
	return &obj
}

func (obj *ultraEthernetCbfcReceiver) msg() *otg.UltraEthernetCbfcReceiver {
	return obj.obj
}

func (obj *ultraEthernetCbfcReceiver) setMsg(msg *otg.UltraEthernetCbfcReceiver) UltraEthernetCbfcReceiver {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalultraEthernetCbfcReceiver struct {
	obj *ultraEthernetCbfcReceiver
}

type marshalUltraEthernetCbfcReceiver interface {
	// ToProto marshals UltraEthernetCbfcReceiver to protobuf object *otg.UltraEthernetCbfcReceiver
	ToProto() (*otg.UltraEthernetCbfcReceiver, error)
	// ToPbText marshals UltraEthernetCbfcReceiver to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UltraEthernetCbfcReceiver to YAML text
	ToYaml() (string, error)
	// ToJson marshals UltraEthernetCbfcReceiver to JSON text
	ToJson() (string, error)
}

type unMarshalultraEthernetCbfcReceiver struct {
	obj *ultraEthernetCbfcReceiver
}

type unMarshalUltraEthernetCbfcReceiver interface {
	// FromProto unmarshals UltraEthernetCbfcReceiver from protobuf object *otg.UltraEthernetCbfcReceiver
	FromProto(msg *otg.UltraEthernetCbfcReceiver) (UltraEthernetCbfcReceiver, error)
	// FromPbText unmarshals UltraEthernetCbfcReceiver from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UltraEthernetCbfcReceiver from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UltraEthernetCbfcReceiver from JSON text
	FromJson(value string) error
}

func (obj *ultraEthernetCbfcReceiver) Marshal() marshalUltraEthernetCbfcReceiver {
	if obj.marshaller == nil {
		obj.marshaller = &marshalultraEthernetCbfcReceiver{obj: obj}
	}
	return obj.marshaller
}

func (obj *ultraEthernetCbfcReceiver) Unmarshal() unMarshalUltraEthernetCbfcReceiver {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalultraEthernetCbfcReceiver{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalultraEthernetCbfcReceiver) ToProto() (*otg.UltraEthernetCbfcReceiver, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalultraEthernetCbfcReceiver) FromProto(msg *otg.UltraEthernetCbfcReceiver) (UltraEthernetCbfcReceiver, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalultraEthernetCbfcReceiver) ToPbText() (string, error) {
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

func (m *unMarshalultraEthernetCbfcReceiver) FromPbText(value string) error {
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

func (m *marshalultraEthernetCbfcReceiver) ToYaml() (string, error) {
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

func (m *unMarshalultraEthernetCbfcReceiver) FromYaml(value string) error {
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

func (m *marshalultraEthernetCbfcReceiver) ToJson() (string, error) {
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

func (m *unMarshalultraEthernetCbfcReceiver) FromJson(value string) error {
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

func (obj *ultraEthernetCbfcReceiver) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfcReceiver) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfcReceiver) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ultraEthernetCbfcReceiver) Clone() (UltraEthernetCbfcReceiver, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUltraEthernetCbfcReceiver()
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

func (obj *ultraEthernetCbfcReceiver) setNil() {
	obj.portCreditHolder = nil
	obj.ccMessageHolder = nil
	obj.virtualChannelsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// UltraEthernetCbfcReceiver is cBFC receiver direction configuration.
type UltraEthernetCbfcReceiver interface {
	Validation
	// msg marshals UltraEthernetCbfcReceiver to protobuf object *otg.UltraEthernetCbfcReceiver
	// and doesn't set defaults
	msg() *otg.UltraEthernetCbfcReceiver
	// setMsg unmarshals UltraEthernetCbfcReceiver from protobuf object *otg.UltraEthernetCbfcReceiver
	// and doesn't set defaults
	setMsg(*otg.UltraEthernetCbfcReceiver) UltraEthernetCbfcReceiver
	// provides marshal interface
	Marshal() marshalUltraEthernetCbfcReceiver
	// provides unmarshal interface
	Unmarshal() unMarshalUltraEthernetCbfcReceiver
	// validate validates UltraEthernetCbfcReceiver
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UltraEthernetCbfcReceiver, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Enable returns bool, set in UltraEthernetCbfcReceiver.
	Enable() bool
	// SetEnable assigns bool provided by user to UltraEthernetCbfcReceiver
	SetEnable(value bool) UltraEthernetCbfcReceiver
	// HasEnable checks if Enable has been set in UltraEthernetCbfcReceiver
	HasEnable() bool
	// PortCredit returns UltraEthernetCbfcPortCredit, set in UltraEthernetCbfcReceiver.
	// UltraEthernetCbfcPortCredit is cBFC port level credit configuration. Credits are the unit used to track
	// available buffer space at the receiver.
	//
	// Reference: UE-Specification-1.0.3 Table 5-14.
	PortCredit() UltraEthernetCbfcPortCredit
	// SetPortCredit assigns UltraEthernetCbfcPortCredit provided by user to UltraEthernetCbfcReceiver.
	// UltraEthernetCbfcPortCredit is cBFC port level credit configuration. Credits are the unit used to track
	// available buffer space at the receiver.
	//
	// Reference: UE-Specification-1.0.3 Table 5-14.
	SetPortCredit(value UltraEthernetCbfcPortCredit) UltraEthernetCbfcReceiver
	// HasPortCredit checks if PortCredit has been set in UltraEthernetCbfcReceiver
	HasPortCredit() bool
	// CcMessage returns UltraEthernetCbfcReceiverCcMessage, set in UltraEthernetCbfcReceiver.
	// UltraEthernetCbfcReceiverCcMessage is cBFC CC_Update (credit consumed) message reception configuration.
	CcMessage() UltraEthernetCbfcReceiverCcMessage
	// SetCcMessage assigns UltraEthernetCbfcReceiverCcMessage provided by user to UltraEthernetCbfcReceiver.
	// UltraEthernetCbfcReceiverCcMessage is cBFC CC_Update (credit consumed) message reception configuration.
	SetCcMessage(value UltraEthernetCbfcReceiverCcMessage) UltraEthernetCbfcReceiver
	// HasCcMessage checks if CcMessage has been set in UltraEthernetCbfcReceiver
	HasCcMessage() bool
	// VirtualChannels returns UltraEthernetCbfcReceiverUltraEthernetCbfcVcIterIter, set in UltraEthernetCbfcReceiver
	VirtualChannels() UltraEthernetCbfcReceiverUltraEthernetCbfcVcIter
	setNil()
}

// Enable the CBFC receiver on the port.
// Enable returns a bool
func (obj *ultraEthernetCbfcReceiver) Enable() bool {

	return *obj.obj.Enable

}

// Enable the CBFC receiver on the port.
// Enable returns a bool
func (obj *ultraEthernetCbfcReceiver) HasEnable() bool {
	return obj.obj.Enable != nil
}

// Enable the CBFC receiver on the port.
// SetEnable sets the bool value in the UltraEthernetCbfcReceiver object
func (obj *ultraEthernetCbfcReceiver) SetEnable(value bool) UltraEthernetCbfcReceiver {

	obj.obj.Enable = &value
	return obj
}

// Port level credit configuration (credit mode, total credits, cell size and
// per-packet overhead).
// PortCredit returns a UltraEthernetCbfcPortCredit
func (obj *ultraEthernetCbfcReceiver) PortCredit() UltraEthernetCbfcPortCredit {
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
func (obj *ultraEthernetCbfcReceiver) HasPortCredit() bool {
	return obj.obj.PortCredit != nil
}

// Port level credit configuration (credit mode, total credits, cell size and
// per-packet overhead).
// SetPortCredit sets the UltraEthernetCbfcPortCredit value in the UltraEthernetCbfcReceiver object
func (obj *ultraEthernetCbfcReceiver) SetPortCredit(value UltraEthernetCbfcPortCredit) UltraEthernetCbfcReceiver {

	obj.portCreditHolder = nil
	obj.obj.PortCredit = value.msg()

	return obj
}

// CC_Update (credit consumed) message reception configuration.
// CcMessage returns a UltraEthernetCbfcReceiverCcMessage
func (obj *ultraEthernetCbfcReceiver) CcMessage() UltraEthernetCbfcReceiverCcMessage {
	if obj.obj.CcMessage == nil {
		obj.obj.CcMessage = NewUltraEthernetCbfcReceiverCcMessage().msg()
	}
	if obj.ccMessageHolder == nil {
		obj.ccMessageHolder = &ultraEthernetCbfcReceiverCcMessage{obj: obj.obj.CcMessage}
	}
	return obj.ccMessageHolder
}

// CC_Update (credit consumed) message reception configuration.
// CcMessage returns a UltraEthernetCbfcReceiverCcMessage
func (obj *ultraEthernetCbfcReceiver) HasCcMessage() bool {
	return obj.obj.CcMessage != nil
}

// CC_Update (credit consumed) message reception configuration.
// SetCcMessage sets the UltraEthernetCbfcReceiverCcMessage value in the UltraEthernetCbfcReceiver object
func (obj *ultraEthernetCbfcReceiver) SetCcMessage(value UltraEthernetCbfcReceiverCcMessage) UltraEthernetCbfcReceiver {

	obj.ccMessageHolder = nil
	obj.obj.CcMessage = value.msg()

	return obj
}

// The list of virtual channels (VCs). A maximum of 32 VCs are supported per
// port, of which a maximum of 4 may be configured as lossless.
// VirtualChannels returns a []UltraEthernetCbfcVc
func (obj *ultraEthernetCbfcReceiver) VirtualChannels() UltraEthernetCbfcReceiverUltraEthernetCbfcVcIter {
	if len(obj.obj.VirtualChannels) == 0 {
		obj.obj.VirtualChannels = []*otg.UltraEthernetCbfcVc{}
	}
	if obj.virtualChannelsHolder == nil {
		obj.virtualChannelsHolder = newUltraEthernetCbfcReceiverUltraEthernetCbfcVcIter(&obj.obj.VirtualChannels).setMsg(obj)
	}
	return obj.virtualChannelsHolder
}

type ultraEthernetCbfcReceiverUltraEthernetCbfcVcIter struct {
	obj                      *ultraEthernetCbfcReceiver
	ultraEthernetCbfcVcSlice []UltraEthernetCbfcVc
	fieldPtr                 *[]*otg.UltraEthernetCbfcVc
}

func newUltraEthernetCbfcReceiverUltraEthernetCbfcVcIter(ptr *[]*otg.UltraEthernetCbfcVc) UltraEthernetCbfcReceiverUltraEthernetCbfcVcIter {
	return &ultraEthernetCbfcReceiverUltraEthernetCbfcVcIter{fieldPtr: ptr}
}

type UltraEthernetCbfcReceiverUltraEthernetCbfcVcIter interface {
	setMsg(*ultraEthernetCbfcReceiver) UltraEthernetCbfcReceiverUltraEthernetCbfcVcIter
	Items() []UltraEthernetCbfcVc
	Add() UltraEthernetCbfcVc
	Append(items ...UltraEthernetCbfcVc) UltraEthernetCbfcReceiverUltraEthernetCbfcVcIter
	Set(index int, newObj UltraEthernetCbfcVc) UltraEthernetCbfcReceiverUltraEthernetCbfcVcIter
	Clear() UltraEthernetCbfcReceiverUltraEthernetCbfcVcIter
	clearHolderSlice() UltraEthernetCbfcReceiverUltraEthernetCbfcVcIter
	appendHolderSlice(item UltraEthernetCbfcVc) UltraEthernetCbfcReceiverUltraEthernetCbfcVcIter
}

func (obj *ultraEthernetCbfcReceiverUltraEthernetCbfcVcIter) setMsg(msg *ultraEthernetCbfcReceiver) UltraEthernetCbfcReceiverUltraEthernetCbfcVcIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ultraEthernetCbfcVc{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ultraEthernetCbfcReceiverUltraEthernetCbfcVcIter) Items() []UltraEthernetCbfcVc {
	return obj.ultraEthernetCbfcVcSlice
}

func (obj *ultraEthernetCbfcReceiverUltraEthernetCbfcVcIter) Add() UltraEthernetCbfcVc {
	newObj := &otg.UltraEthernetCbfcVc{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ultraEthernetCbfcVc{obj: newObj}
	newLibObj.setDefault()
	obj.ultraEthernetCbfcVcSlice = append(obj.ultraEthernetCbfcVcSlice, newLibObj)
	return newLibObj
}

func (obj *ultraEthernetCbfcReceiverUltraEthernetCbfcVcIter) Append(items ...UltraEthernetCbfcVc) UltraEthernetCbfcReceiverUltraEthernetCbfcVcIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ultraEthernetCbfcVcSlice = append(obj.ultraEthernetCbfcVcSlice, item)
	}
	return obj
}

func (obj *ultraEthernetCbfcReceiverUltraEthernetCbfcVcIter) Set(index int, newObj UltraEthernetCbfcVc) UltraEthernetCbfcReceiverUltraEthernetCbfcVcIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ultraEthernetCbfcVcSlice[index] = newObj
	return obj
}
func (obj *ultraEthernetCbfcReceiverUltraEthernetCbfcVcIter) Clear() UltraEthernetCbfcReceiverUltraEthernetCbfcVcIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.UltraEthernetCbfcVc{}
		obj.ultraEthernetCbfcVcSlice = []UltraEthernetCbfcVc{}
	}
	return obj
}
func (obj *ultraEthernetCbfcReceiverUltraEthernetCbfcVcIter) clearHolderSlice() UltraEthernetCbfcReceiverUltraEthernetCbfcVcIter {
	if len(obj.ultraEthernetCbfcVcSlice) > 0 {
		obj.ultraEthernetCbfcVcSlice = []UltraEthernetCbfcVc{}
	}
	return obj
}
func (obj *ultraEthernetCbfcReceiverUltraEthernetCbfcVcIter) appendHolderSlice(item UltraEthernetCbfcVc) UltraEthernetCbfcReceiverUltraEthernetCbfcVcIter {
	obj.ultraEthernetCbfcVcSlice = append(obj.ultraEthernetCbfcVcSlice, item)
	return obj
}

func (obj *ultraEthernetCbfcReceiver) validateObj(vObj *validation, set_default bool) {
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

func (obj *ultraEthernetCbfcReceiver) setDefault() {
	if obj.obj.Enable == nil {
		obj.SetEnable(false)
	}

}
