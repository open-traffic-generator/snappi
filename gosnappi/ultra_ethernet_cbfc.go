package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UltraEthernetCbfc *****
type ultraEthernetCbfc struct {
	validation
	obj            *otg.UltraEthernetCbfc
	marshaller     marshalUltraEthernetCbfc
	unMarshaller   unMarshalUltraEthernetCbfc
	senderHolder   UltraEthernetCbfcSender
	receiverHolder UltraEthernetCbfcReceiver
}

func NewUltraEthernetCbfc() UltraEthernetCbfc {
	obj := ultraEthernetCbfc{obj: &otg.UltraEthernetCbfc{}}
	obj.setDefault()
	return &obj
}

func (obj *ultraEthernetCbfc) msg() *otg.UltraEthernetCbfc {
	return obj.obj
}

func (obj *ultraEthernetCbfc) setMsg(msg *otg.UltraEthernetCbfc) UltraEthernetCbfc {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalultraEthernetCbfc struct {
	obj *ultraEthernetCbfc
}

type marshalUltraEthernetCbfc interface {
	// ToProto marshals UltraEthernetCbfc to protobuf object *otg.UltraEthernetCbfc
	ToProto() (*otg.UltraEthernetCbfc, error)
	// ToPbText marshals UltraEthernetCbfc to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UltraEthernetCbfc to YAML text
	ToYaml() (string, error)
	// ToJson marshals UltraEthernetCbfc to JSON text
	ToJson() (string, error)
}

type unMarshalultraEthernetCbfc struct {
	obj *ultraEthernetCbfc
}

type unMarshalUltraEthernetCbfc interface {
	// FromProto unmarshals UltraEthernetCbfc from protobuf object *otg.UltraEthernetCbfc
	FromProto(msg *otg.UltraEthernetCbfc) (UltraEthernetCbfc, error)
	// FromPbText unmarshals UltraEthernetCbfc from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UltraEthernetCbfc from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UltraEthernetCbfc from JSON text
	FromJson(value string) error
}

func (obj *ultraEthernetCbfc) Marshal() marshalUltraEthernetCbfc {
	if obj.marshaller == nil {
		obj.marshaller = &marshalultraEthernetCbfc{obj: obj}
	}
	return obj.marshaller
}

func (obj *ultraEthernetCbfc) Unmarshal() unMarshalUltraEthernetCbfc {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalultraEthernetCbfc{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalultraEthernetCbfc) ToProto() (*otg.UltraEthernetCbfc, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalultraEthernetCbfc) FromProto(msg *otg.UltraEthernetCbfc) (UltraEthernetCbfc, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalultraEthernetCbfc) ToPbText() (string, error) {
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

func (m *unMarshalultraEthernetCbfc) FromPbText(value string) error {
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

func (m *marshalultraEthernetCbfc) ToYaml() (string, error) {
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

func (m *unMarshalultraEthernetCbfc) FromYaml(value string) error {
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

func (m *marshalultraEthernetCbfc) ToJson() (string, error) {
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

func (m *unMarshalultraEthernetCbfc) FromJson(value string) error {
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

func (obj *ultraEthernetCbfc) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfc) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfc) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ultraEthernetCbfc) Clone() (UltraEthernetCbfc, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUltraEthernetCbfc()
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

func (obj *ultraEthernetCbfc) setNil() {
	obj.senderHolder = nil
	obj.receiverHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// UltraEthernetCbfc is ultra Ethernet Credit-based Flow Control (CBFC) settings. CBFC is a per
// virtual channel (VC) credit based flow control mechanism and is an alternative
// to priority based flow control (PFC). The sender and receiver directions of a
// link are configured independently.
//
// Reference: UE-Specification-1.0.3 Section 5.2.
type UltraEthernetCbfc interface {
	Validation
	// msg marshals UltraEthernetCbfc to protobuf object *otg.UltraEthernetCbfc
	// and doesn't set defaults
	msg() *otg.UltraEthernetCbfc
	// setMsg unmarshals UltraEthernetCbfc from protobuf object *otg.UltraEthernetCbfc
	// and doesn't set defaults
	setMsg(*otg.UltraEthernetCbfc) UltraEthernetCbfc
	// provides marshal interface
	Marshal() marshalUltraEthernetCbfc
	// provides unmarshal interface
	Unmarshal() unMarshalUltraEthernetCbfc
	// validate validates UltraEthernetCbfc
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UltraEthernetCbfc, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Sender returns UltraEthernetCbfcSender, set in UltraEthernetCbfc.
	// UltraEthernetCbfcSender is cBFC sender direction configuration.
	Sender() UltraEthernetCbfcSender
	// SetSender assigns UltraEthernetCbfcSender provided by user to UltraEthernetCbfc.
	// UltraEthernetCbfcSender is cBFC sender direction configuration.
	SetSender(value UltraEthernetCbfcSender) UltraEthernetCbfc
	// HasSender checks if Sender has been set in UltraEthernetCbfc
	HasSender() bool
	// Receiver returns UltraEthernetCbfcReceiver, set in UltraEthernetCbfc.
	// UltraEthernetCbfcReceiver is cBFC receiver direction configuration.
	Receiver() UltraEthernetCbfcReceiver
	// SetReceiver assigns UltraEthernetCbfcReceiver provided by user to UltraEthernetCbfc.
	// UltraEthernetCbfcReceiver is cBFC receiver direction configuration.
	SetReceiver(value UltraEthernetCbfcReceiver) UltraEthernetCbfc
	// HasReceiver checks if Receiver has been set in UltraEthernetCbfc
	HasReceiver() bool
	setNil()
}

// CBFC sender direction settings. The sender transmits packets on lossless
// VCs only when enough credits are available at the receiver, and periodically
// transmits CC_Update messages.
// Sender returns a UltraEthernetCbfcSender
func (obj *ultraEthernetCbfc) Sender() UltraEthernetCbfcSender {
	if obj.obj.Sender == nil {
		obj.obj.Sender = NewUltraEthernetCbfcSender().msg()
	}
	if obj.senderHolder == nil {
		obj.senderHolder = &ultraEthernetCbfcSender{obj: obj.obj.Sender}
	}
	return obj.senderHolder
}

// CBFC sender direction settings. The sender transmits packets on lossless
// VCs only when enough credits are available at the receiver, and periodically
// transmits CC_Update messages.
// Sender returns a UltraEthernetCbfcSender
func (obj *ultraEthernetCbfc) HasSender() bool {
	return obj.obj.Sender != nil
}

// CBFC sender direction settings. The sender transmits packets on lossless
// VCs only when enough credits are available at the receiver, and periodically
// transmits CC_Update messages.
// SetSender sets the UltraEthernetCbfcSender value in the UltraEthernetCbfc object
func (obj *ultraEthernetCbfc) SetSender(value UltraEthernetCbfcSender) UltraEthernetCbfc {

	obj.senderHolder = nil
	obj.obj.Sender = value.msg()

	return obj
}

// CBFC receiver direction settings. The receiver returns credits to the
// sender using CF_Update messages as buffer space is freed.
// Receiver returns a UltraEthernetCbfcReceiver
func (obj *ultraEthernetCbfc) Receiver() UltraEthernetCbfcReceiver {
	if obj.obj.Receiver == nil {
		obj.obj.Receiver = NewUltraEthernetCbfcReceiver().msg()
	}
	if obj.receiverHolder == nil {
		obj.receiverHolder = &ultraEthernetCbfcReceiver{obj: obj.obj.Receiver}
	}
	return obj.receiverHolder
}

// CBFC receiver direction settings. The receiver returns credits to the
// sender using CF_Update messages as buffer space is freed.
// Receiver returns a UltraEthernetCbfcReceiver
func (obj *ultraEthernetCbfc) HasReceiver() bool {
	return obj.obj.Receiver != nil
}

// CBFC receiver direction settings. The receiver returns credits to the
// sender using CF_Update messages as buffer space is freed.
// SetReceiver sets the UltraEthernetCbfcReceiver value in the UltraEthernetCbfc object
func (obj *ultraEthernetCbfc) SetReceiver(value UltraEthernetCbfcReceiver) UltraEthernetCbfc {

	obj.receiverHolder = nil
	obj.obj.Receiver = value.msg()

	return obj
}

func (obj *ultraEthernetCbfc) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Sender != nil {

		obj.Sender().validateObj(vObj, set_default)
	}

	if obj.obj.Receiver != nil {

		obj.Receiver().validateObj(vObj, set_default)
	}

}

func (obj *ultraEthernetCbfc) setDefault() {

}
