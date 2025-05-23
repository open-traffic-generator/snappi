package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2AckAndNak *****
type rocev2AckAndNak struct {
	validation
	obj          *otg.Rocev2AckAndNak
	marshaller   marshalRocev2AckAndNak
	unMarshaller unMarshalRocev2AckAndNak
	ackHolder    Rocev2ACK
	nakHolder    Rocev2NAK
}

func NewRocev2AckAndNak() Rocev2AckAndNak {
	obj := rocev2AckAndNak{obj: &otg.Rocev2AckAndNak{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2AckAndNak) msg() *otg.Rocev2AckAndNak {
	return obj.obj
}

func (obj *rocev2AckAndNak) setMsg(msg *otg.Rocev2AckAndNak) Rocev2AckAndNak {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2AckAndNak struct {
	obj *rocev2AckAndNak
}

type marshalRocev2AckAndNak interface {
	// ToProto marshals Rocev2AckAndNak to protobuf object *otg.Rocev2AckAndNak
	ToProto() (*otg.Rocev2AckAndNak, error)
	// ToPbText marshals Rocev2AckAndNak to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2AckAndNak to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2AckAndNak to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Rocev2AckAndNak to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalrocev2AckAndNak struct {
	obj *rocev2AckAndNak
}

type unMarshalRocev2AckAndNak interface {
	// FromProto unmarshals Rocev2AckAndNak from protobuf object *otg.Rocev2AckAndNak
	FromProto(msg *otg.Rocev2AckAndNak) (Rocev2AckAndNak, error)
	// FromPbText unmarshals Rocev2AckAndNak from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2AckAndNak from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2AckAndNak from JSON text
	FromJson(value string) error
}

func (obj *rocev2AckAndNak) Marshal() marshalRocev2AckAndNak {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2AckAndNak{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2AckAndNak) Unmarshal() unMarshalRocev2AckAndNak {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2AckAndNak{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2AckAndNak) ToProto() (*otg.Rocev2AckAndNak, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2AckAndNak) FromProto(msg *otg.Rocev2AckAndNak) (Rocev2AckAndNak, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2AckAndNak) ToPbText() (string, error) {
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

func (m *unMarshalrocev2AckAndNak) FromPbText(value string) error {
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

func (m *marshalrocev2AckAndNak) ToYaml() (string, error) {
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

func (m *unMarshalrocev2AckAndNak) FromYaml(value string) error {
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

func (m *marshalrocev2AckAndNak) ToJsonRaw() (string, error) {
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

func (m *marshalrocev2AckAndNak) ToJson() (string, error) {
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

func (m *unMarshalrocev2AckAndNak) FromJson(value string) error {
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

func (obj *rocev2AckAndNak) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2AckAndNak) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2AckAndNak) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2AckAndNak) Clone() (Rocev2AckAndNak, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2AckAndNak()
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

func (obj *rocev2AckAndNak) setNil() {
	obj.ackHolder = nil
	obj.nakHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Rocev2AckAndNak is defines the ACK and NAK settings for RoCEv2. This configuration ensures reliable data delivery by controlling how the system responds to successful and failed packet transmissions.
type Rocev2AckAndNak interface {
	Validation
	// msg marshals Rocev2AckAndNak to protobuf object *otg.Rocev2AckAndNak
	// and doesn't set defaults
	msg() *otg.Rocev2AckAndNak
	// setMsg unmarshals Rocev2AckAndNak from protobuf object *otg.Rocev2AckAndNak
	// and doesn't set defaults
	setMsg(*otg.Rocev2AckAndNak) Rocev2AckAndNak
	// provides marshal interface
	Marshal() marshalRocev2AckAndNak
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2AckAndNak
	// validate validates Rocev2AckAndNak
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2AckAndNak, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ack returns Rocev2ACK, set in Rocev2AckAndNak.
	// Rocev2ACK is aCK parameters.
	Ack() Rocev2ACK
	// SetAck assigns Rocev2ACK provided by user to Rocev2AckAndNak.
	// Rocev2ACK is aCK parameters.
	SetAck(value Rocev2ACK) Rocev2AckAndNak
	// HasAck checks if Ack has been set in Rocev2AckAndNak
	HasAck() bool
	// Nak returns Rocev2NAK, set in Rocev2AckAndNak.
	// Rocev2NAK is nAK parameters.
	Nak() Rocev2NAK
	// SetNak assigns Rocev2NAK provided by user to Rocev2AckAndNak.
	// Rocev2NAK is nAK parameters.
	SetNak(value Rocev2NAK) Rocev2AckAndNak
	// HasNak checks if Nak has been set in Rocev2AckAndNak
	HasNak() bool
	// EnableRetransmissionTimeout returns bool, set in Rocev2AckAndNak.
	EnableRetransmissionTimeout() bool
	// SetEnableRetransmissionTimeout assigns bool provided by user to Rocev2AckAndNak
	SetEnableRetransmissionTimeout(value bool) Rocev2AckAndNak
	// HasEnableRetransmissionTimeout checks if EnableRetransmissionTimeout has been set in Rocev2AckAndNak
	HasEnableRetransmissionTimeout() bool
	// RetransmissionTimeoutValue returns uint32, set in Rocev2AckAndNak.
	RetransmissionTimeoutValue() uint32
	// SetRetransmissionTimeoutValue assigns uint32 provided by user to Rocev2AckAndNak
	SetRetransmissionTimeoutValue(value uint32) Rocev2AckAndNak
	// HasRetransmissionTimeoutValue checks if RetransmissionTimeoutValue has been set in Rocev2AckAndNak
	HasRetransmissionTimeoutValue() bool
	// RetransmissionRetryCount returns uint32, set in Rocev2AckAndNak.
	RetransmissionRetryCount() uint32
	// SetRetransmissionRetryCount assigns uint32 provided by user to Rocev2AckAndNak
	SetRetransmissionRetryCount(value uint32) Rocev2AckAndNak
	// HasRetransmissionRetryCount checks if RetransmissionRetryCount has been set in Rocev2AckAndNak
	HasRetransmissionRetryCount() bool
	setNil()
}

// description is TBD
// Ack returns a Rocev2ACK
func (obj *rocev2AckAndNak) Ack() Rocev2ACK {
	if obj.obj.Ack == nil {
		obj.obj.Ack = NewRocev2ACK().msg()
	}
	if obj.ackHolder == nil {
		obj.ackHolder = &rocev2ACK{obj: obj.obj.Ack}
	}
	return obj.ackHolder
}

// description is TBD
// Ack returns a Rocev2ACK
func (obj *rocev2AckAndNak) HasAck() bool {
	return obj.obj.Ack != nil
}

// description is TBD
// SetAck sets the Rocev2ACK value in the Rocev2AckAndNak object
func (obj *rocev2AckAndNak) SetAck(value Rocev2ACK) Rocev2AckAndNak {

	obj.ackHolder = nil
	obj.obj.Ack = value.msg()

	return obj
}

// description is TBD
// Nak returns a Rocev2NAK
func (obj *rocev2AckAndNak) Nak() Rocev2NAK {
	if obj.obj.Nak == nil {
		obj.obj.Nak = NewRocev2NAK().msg()
	}
	if obj.nakHolder == nil {
		obj.nakHolder = &rocev2NAK{obj: obj.obj.Nak}
	}
	return obj.nakHolder
}

// description is TBD
// Nak returns a Rocev2NAK
func (obj *rocev2AckAndNak) HasNak() bool {
	return obj.obj.Nak != nil
}

// description is TBD
// SetNak sets the Rocev2NAK value in the Rocev2AckAndNak object
func (obj *rocev2AckAndNak) SetNak(value Rocev2NAK) Rocev2AckAndNak {

	obj.nakHolder = nil
	obj.obj.Nak = value.msg()

	return obj
}

// Enable Retransmission on ACK Timeout.
// EnableRetransmissionTimeout returns a bool
func (obj *rocev2AckAndNak) EnableRetransmissionTimeout() bool {

	return *obj.obj.EnableRetransmissionTimeout

}

// Enable Retransmission on ACK Timeout.
// EnableRetransmissionTimeout returns a bool
func (obj *rocev2AckAndNak) HasEnableRetransmissionTimeout() bool {
	return obj.obj.EnableRetransmissionTimeout != nil
}

// Enable Retransmission on ACK Timeout.
// SetEnableRetransmissionTimeout sets the bool value in the Rocev2AckAndNak object
func (obj *rocev2AckAndNak) SetEnableRetransmissionTimeout(value bool) Rocev2AckAndNak {

	obj.obj.EnableRetransmissionTimeout = &value
	return obj
}

// The duration to wait before retrying transmission upon not receiving an acknowledgment (ACK) or negative acknowledgment (NAK) is specified in milliseconds.
// RetransmissionTimeoutValue returns a uint32
func (obj *rocev2AckAndNak) RetransmissionTimeoutValue() uint32 {

	return *obj.obj.RetransmissionTimeoutValue

}

// The duration to wait before retrying transmission upon not receiving an acknowledgment (ACK) or negative acknowledgment (NAK) is specified in milliseconds.
// RetransmissionTimeoutValue returns a uint32
func (obj *rocev2AckAndNak) HasRetransmissionTimeoutValue() bool {
	return obj.obj.RetransmissionTimeoutValue != nil
}

// The duration to wait before retrying transmission upon not receiving an acknowledgment (ACK) or negative acknowledgment (NAK) is specified in milliseconds.
// SetRetransmissionTimeoutValue sets the uint32 value in the Rocev2AckAndNak object
func (obj *rocev2AckAndNak) SetRetransmissionTimeoutValue(value uint32) Rocev2AckAndNak {

	obj.obj.RetransmissionTimeoutValue = &value
	return obj
}

// Number of retransmission attempts before stopping due to missing ACK/NAK.
// RetransmissionRetryCount returns a uint32
func (obj *rocev2AckAndNak) RetransmissionRetryCount() uint32 {

	return *obj.obj.RetransmissionRetryCount

}

// Number of retransmission attempts before stopping due to missing ACK/NAK.
// RetransmissionRetryCount returns a uint32
func (obj *rocev2AckAndNak) HasRetransmissionRetryCount() bool {
	return obj.obj.RetransmissionRetryCount != nil
}

// Number of retransmission attempts before stopping due to missing ACK/NAK.
// SetRetransmissionRetryCount sets the uint32 value in the Rocev2AckAndNak object
func (obj *rocev2AckAndNak) SetRetransmissionRetryCount(value uint32) Rocev2AckAndNak {

	obj.obj.RetransmissionRetryCount = &value
	return obj
}

func (obj *rocev2AckAndNak) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ack != nil {

		obj.Ack().validateObj(vObj, set_default)
	}

	if obj.obj.Nak != nil {

		obj.Nak().validateObj(vObj, set_default)
	}

	if obj.obj.RetransmissionTimeoutValue != nil {

		if *obj.obj.RetransmissionTimeoutValue > 5369 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2AckAndNak.RetransmissionTimeoutValue <= 5369 but Got %d", *obj.obj.RetransmissionTimeoutValue))
		}

	}

	if obj.obj.RetransmissionRetryCount != nil {

		if *obj.obj.RetransmissionRetryCount > 254 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2AckAndNak.RetransmissionRetryCount <= 254 but Got %d", *obj.obj.RetransmissionRetryCount))
		}

	}

}

func (obj *rocev2AckAndNak) setDefault() {
	if obj.obj.EnableRetransmissionTimeout == nil {
		obj.SetEnableRetransmissionTimeout(true)
	}
	if obj.obj.RetransmissionTimeoutValue == nil {
		obj.SetRetransmissionTimeoutValue(1)
	}
	if obj.obj.RetransmissionRetryCount == nil {
		obj.SetRetransmissionRetryCount(3)
	}

}
