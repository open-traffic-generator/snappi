package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2PerPortSettings *****
type rocev2PerPortSettings struct {
	validation
	obj                 *otg.Rocev2PerPortSettings
	marshaller          marshalRocev2PerPortSettings
	unMarshaller        unMarshalRocev2PerPortSettings
	cnpHolder           Rocev2CNP
	ackHolder           Rocev2ACK
	nackHolder          Rocev2NACK
	dcqcnSettingsHolder Rocev2DCQCN
}

func NewRocev2PerPortSettings() Rocev2PerPortSettings {
	obj := rocev2PerPortSettings{obj: &otg.Rocev2PerPortSettings{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2PerPortSettings) msg() *otg.Rocev2PerPortSettings {
	return obj.obj
}

func (obj *rocev2PerPortSettings) setMsg(msg *otg.Rocev2PerPortSettings) Rocev2PerPortSettings {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2PerPortSettings struct {
	obj *rocev2PerPortSettings
}

type marshalRocev2PerPortSettings interface {
	// ToProto marshals Rocev2PerPortSettings to protobuf object *otg.Rocev2PerPortSettings
	ToProto() (*otg.Rocev2PerPortSettings, error)
	// ToPbText marshals Rocev2PerPortSettings to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2PerPortSettings to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2PerPortSettings to JSON text
	ToJson() (string, error)
}

type unMarshalrocev2PerPortSettings struct {
	obj *rocev2PerPortSettings
}

type unMarshalRocev2PerPortSettings interface {
	// FromProto unmarshals Rocev2PerPortSettings from protobuf object *otg.Rocev2PerPortSettings
	FromProto(msg *otg.Rocev2PerPortSettings) (Rocev2PerPortSettings, error)
	// FromPbText unmarshals Rocev2PerPortSettings from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2PerPortSettings from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2PerPortSettings from JSON text
	FromJson(value string) error
}

func (obj *rocev2PerPortSettings) Marshal() marshalRocev2PerPortSettings {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2PerPortSettings{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2PerPortSettings) Unmarshal() unMarshalRocev2PerPortSettings {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2PerPortSettings{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2PerPortSettings) ToProto() (*otg.Rocev2PerPortSettings, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2PerPortSettings) FromProto(msg *otg.Rocev2PerPortSettings) (Rocev2PerPortSettings, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2PerPortSettings) ToPbText() (string, error) {
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

func (m *unMarshalrocev2PerPortSettings) FromPbText(value string) error {
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

func (m *marshalrocev2PerPortSettings) ToYaml() (string, error) {
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

func (m *unMarshalrocev2PerPortSettings) FromYaml(value string) error {
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

func (m *marshalrocev2PerPortSettings) ToJson() (string, error) {
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

func (m *unMarshalrocev2PerPortSettings) FromJson(value string) error {
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

func (obj *rocev2PerPortSettings) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2PerPortSettings) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2PerPortSettings) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2PerPortSettings) Clone() (Rocev2PerPortSettings, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2PerPortSettings()
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

func (obj *rocev2PerPortSettings) setNil() {
	obj.cnpHolder = nil
	obj.ackHolder = nil
	obj.nackHolder = nil
	obj.dcqcnSettingsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Rocev2PerPortSettings is a high level data plane traffic flow.
type Rocev2PerPortSettings interface {
	Validation
	// msg marshals Rocev2PerPortSettings to protobuf object *otg.Rocev2PerPortSettings
	// and doesn't set defaults
	msg() *otg.Rocev2PerPortSettings
	// setMsg unmarshals Rocev2PerPortSettings from protobuf object *otg.Rocev2PerPortSettings
	// and doesn't set defaults
	setMsg(*otg.Rocev2PerPortSettings) Rocev2PerPortSettings
	// provides marshal interface
	Marshal() marshalRocev2PerPortSettings
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2PerPortSettings
	// validate validates Rocev2PerPortSettings
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2PerPortSettings, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Cnp returns Rocev2CNP, set in Rocev2PerPortSettings.
	// Rocev2CNP is cNP parameters.
	Cnp() Rocev2CNP
	// SetCnp assigns Rocev2CNP provided by user to Rocev2PerPortSettings.
	// Rocev2CNP is cNP parameters.
	SetCnp(value Rocev2CNP) Rocev2PerPortSettings
	// HasCnp checks if Cnp has been set in Rocev2PerPortSettings
	HasCnp() bool
	// Ack returns Rocev2ACK, set in Rocev2PerPortSettings.
	// Rocev2ACK is aCK parameters.
	Ack() Rocev2ACK
	// SetAck assigns Rocev2ACK provided by user to Rocev2PerPortSettings.
	// Rocev2ACK is aCK parameters.
	SetAck(value Rocev2ACK) Rocev2PerPortSettings
	// HasAck checks if Ack has been set in Rocev2PerPortSettings
	HasAck() bool
	// Nack returns Rocev2NACK, set in Rocev2PerPortSettings.
	// Rocev2NACK is nACK parameters.
	Nack() Rocev2NACK
	// SetNack assigns Rocev2NACK provided by user to Rocev2PerPortSettings.
	// Rocev2NACK is nACK parameters.
	SetNack(value Rocev2NACK) Rocev2PerPortSettings
	// HasNack checks if Nack has been set in Rocev2PerPortSettings
	HasNack() bool
	// CnpDelayTimer returns uint32, set in Rocev2PerPortSettings.
	CnpDelayTimer() uint32
	// SetCnpDelayTimer assigns uint32 provided by user to Rocev2PerPortSettings
	SetCnpDelayTimer(value uint32) Rocev2PerPortSettings
	// HasCnpDelayTimer checks if CnpDelayTimer has been set in Rocev2PerPortSettings
	HasCnpDelayTimer() bool
	// EnableRetransmissionAckTimeout returns bool, set in Rocev2PerPortSettings.
	EnableRetransmissionAckTimeout() bool
	// SetEnableRetransmissionAckTimeout assigns bool provided by user to Rocev2PerPortSettings
	SetEnableRetransmissionAckTimeout(value bool) Rocev2PerPortSettings
	// HasEnableRetransmissionAckTimeout checks if EnableRetransmissionAckTimeout has been set in Rocev2PerPortSettings
	HasEnableRetransmissionAckTimeout() bool
	// AckTimeout returns uint32, set in Rocev2PerPortSettings.
	AckTimeout() uint32
	// SetAckTimeout assigns uint32 provided by user to Rocev2PerPortSettings
	SetAckTimeout(value uint32) Rocev2PerPortSettings
	// HasAckTimeout checks if AckTimeout has been set in Rocev2PerPortSettings
	HasAckTimeout() bool
	// RetransmissionRetryCount returns uint32, set in Rocev2PerPortSettings.
	RetransmissionRetryCount() uint32
	// SetRetransmissionRetryCount assigns uint32 provided by user to Rocev2PerPortSettings
	SetRetransmissionRetryCount(value uint32) Rocev2PerPortSettings
	// HasRetransmissionRetryCount checks if RetransmissionRetryCount has been set in Rocev2PerPortSettings
	HasRetransmissionRetryCount() bool
	// DcqcnSettings returns Rocev2DCQCN, set in Rocev2PerPortSettings.
	// Rocev2DCQCN is rocev2 DCQCN Settings.
	DcqcnSettings() Rocev2DCQCN
	// SetDcqcnSettings assigns Rocev2DCQCN provided by user to Rocev2PerPortSettings.
	// Rocev2DCQCN is rocev2 DCQCN Settings.
	SetDcqcnSettings(value Rocev2DCQCN) Rocev2PerPortSettings
	// HasDcqcnSettings checks if DcqcnSettings has been set in Rocev2PerPortSettings
	HasDcqcnSettings() bool
	setNil()
}

// description is TBD
// Cnp returns a Rocev2CNP
func (obj *rocev2PerPortSettings) Cnp() Rocev2CNP {
	if obj.obj.Cnp == nil {
		obj.obj.Cnp = NewRocev2CNP().msg()
	}
	if obj.cnpHolder == nil {
		obj.cnpHolder = &rocev2CNP{obj: obj.obj.Cnp}
	}
	return obj.cnpHolder
}

// description is TBD
// Cnp returns a Rocev2CNP
func (obj *rocev2PerPortSettings) HasCnp() bool {
	return obj.obj.Cnp != nil
}

// description is TBD
// SetCnp sets the Rocev2CNP value in the Rocev2PerPortSettings object
func (obj *rocev2PerPortSettings) SetCnp(value Rocev2CNP) Rocev2PerPortSettings {

	obj.cnpHolder = nil
	obj.obj.Cnp = value.msg()

	return obj
}

// description is TBD
// Ack returns a Rocev2ACK
func (obj *rocev2PerPortSettings) Ack() Rocev2ACK {
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
func (obj *rocev2PerPortSettings) HasAck() bool {
	return obj.obj.Ack != nil
}

// description is TBD
// SetAck sets the Rocev2ACK value in the Rocev2PerPortSettings object
func (obj *rocev2PerPortSettings) SetAck(value Rocev2ACK) Rocev2PerPortSettings {

	obj.ackHolder = nil
	obj.obj.Ack = value.msg()

	return obj
}

// description is TBD
// Nack returns a Rocev2NACK
func (obj *rocev2PerPortSettings) Nack() Rocev2NACK {
	if obj.obj.Nack == nil {
		obj.obj.Nack = NewRocev2NACK().msg()
	}
	if obj.nackHolder == nil {
		obj.nackHolder = &rocev2NACK{obj: obj.obj.Nack}
	}
	return obj.nackHolder
}

// description is TBD
// Nack returns a Rocev2NACK
func (obj *rocev2PerPortSettings) HasNack() bool {
	return obj.obj.Nack != nil
}

// description is TBD
// SetNack sets the Rocev2NACK value in the Rocev2PerPortSettings object
func (obj *rocev2PerPortSettings) SetNack(value Rocev2NACK) Rocev2PerPortSettings {

	obj.nackHolder = nil
	obj.obj.Nack = value.msg()

	return obj
}

// Amount of time to wait between the generation of successive CNP packets. Time in microseconds.
// CnpDelayTimer returns a uint32
func (obj *rocev2PerPortSettings) CnpDelayTimer() uint32 {

	return *obj.obj.CnpDelayTimer

}

// Amount of time to wait between the generation of successive CNP packets. Time in microseconds.
// CnpDelayTimer returns a uint32
func (obj *rocev2PerPortSettings) HasCnpDelayTimer() bool {
	return obj.obj.CnpDelayTimer != nil
}

// Amount of time to wait between the generation of successive CNP packets. Time in microseconds.
// SetCnpDelayTimer sets the uint32 value in the Rocev2PerPortSettings object
func (obj *rocev2PerPortSettings) SetCnpDelayTimer(value uint32) Rocev2PerPortSettings {

	obj.obj.CnpDelayTimer = &value
	return obj
}

// Enable Retransmission on ACK Timeout.
// EnableRetransmissionAckTimeout returns a bool
func (obj *rocev2PerPortSettings) EnableRetransmissionAckTimeout() bool {

	return *obj.obj.EnableRetransmissionAckTimeout

}

// Enable Retransmission on ACK Timeout.
// EnableRetransmissionAckTimeout returns a bool
func (obj *rocev2PerPortSettings) HasEnableRetransmissionAckTimeout() bool {
	return obj.obj.EnableRetransmissionAckTimeout != nil
}

// Enable Retransmission on ACK Timeout.
// SetEnableRetransmissionAckTimeout sets the bool value in the Rocev2PerPortSettings object
func (obj *rocev2PerPortSettings) SetEnableRetransmissionAckTimeout(value bool) Rocev2PerPortSettings {

	obj.obj.EnableRetransmissionAckTimeout = &value
	return obj
}

// Amount of time to wait before attempting to retransmit on non receipt of ACK. Time in milliseconds.
// AckTimeout returns a uint32
func (obj *rocev2PerPortSettings) AckTimeout() uint32 {

	return *obj.obj.AckTimeout

}

// Amount of time to wait before attempting to retransmit on non receipt of ACK. Time in milliseconds.
// AckTimeout returns a uint32
func (obj *rocev2PerPortSettings) HasAckTimeout() bool {
	return obj.obj.AckTimeout != nil
}

// Amount of time to wait before attempting to retransmit on non receipt of ACK. Time in milliseconds.
// SetAckTimeout sets the uint32 value in the Rocev2PerPortSettings object
func (obj *rocev2PerPortSettings) SetAckTimeout(value uint32) Rocev2PerPortSettings {

	obj.obj.AckTimeout = &value
	return obj
}

// Number of times to retry retransmission before giving up on non receipt of ACK.
// RetransmissionRetryCount returns a uint32
func (obj *rocev2PerPortSettings) RetransmissionRetryCount() uint32 {

	return *obj.obj.RetransmissionRetryCount

}

// Number of times to retry retransmission before giving up on non receipt of ACK.
// RetransmissionRetryCount returns a uint32
func (obj *rocev2PerPortSettings) HasRetransmissionRetryCount() bool {
	return obj.obj.RetransmissionRetryCount != nil
}

// Number of times to retry retransmission before giving up on non receipt of ACK.
// SetRetransmissionRetryCount sets the uint32 value in the Rocev2PerPortSettings object
func (obj *rocev2PerPortSettings) SetRetransmissionRetryCount(value uint32) Rocev2PerPortSettings {

	obj.obj.RetransmissionRetryCount = &value
	return obj
}

// description is TBD
// DcqcnSettings returns a Rocev2DCQCN
func (obj *rocev2PerPortSettings) DcqcnSettings() Rocev2DCQCN {
	if obj.obj.DcqcnSettings == nil {
		obj.obj.DcqcnSettings = NewRocev2DCQCN().msg()
	}
	if obj.dcqcnSettingsHolder == nil {
		obj.dcqcnSettingsHolder = &rocev2DCQCN{obj: obj.obj.DcqcnSettings}
	}
	return obj.dcqcnSettingsHolder
}

// description is TBD
// DcqcnSettings returns a Rocev2DCQCN
func (obj *rocev2PerPortSettings) HasDcqcnSettings() bool {
	return obj.obj.DcqcnSettings != nil
}

// description is TBD
// SetDcqcnSettings sets the Rocev2DCQCN value in the Rocev2PerPortSettings object
func (obj *rocev2PerPortSettings) SetDcqcnSettings(value Rocev2DCQCN) Rocev2PerPortSettings {

	obj.dcqcnSettingsHolder = nil
	obj.obj.DcqcnSettings = value.msg()

	return obj
}

func (obj *rocev2PerPortSettings) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Cnp != nil {

		obj.Cnp().validateObj(vObj, set_default)
	}

	if obj.obj.Ack != nil {

		obj.Ack().validateObj(vObj, set_default)
	}

	if obj.obj.Nack != nil {

		obj.Nack().validateObj(vObj, set_default)
	}

	if obj.obj.CnpDelayTimer != nil {

		if *obj.obj.CnpDelayTimer > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2PerPortSettings.CnpDelayTimer <= 255 but Got %d", *obj.obj.CnpDelayTimer))
		}

	}

	if obj.obj.AckTimeout != nil {

		if *obj.obj.AckTimeout > 5369 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2PerPortSettings.AckTimeout <= 5369 but Got %d", *obj.obj.AckTimeout))
		}

	}

	if obj.obj.RetransmissionRetryCount != nil {

		if *obj.obj.RetransmissionRetryCount > 254 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2PerPortSettings.RetransmissionRetryCount <= 254 but Got %d", *obj.obj.RetransmissionRetryCount))
		}

	}

	if obj.obj.DcqcnSettings != nil {

		obj.DcqcnSettings().validateObj(vObj, set_default)
	}

}

func (obj *rocev2PerPortSettings) setDefault() {
	if obj.obj.CnpDelayTimer == nil {
		obj.SetCnpDelayTimer(55)
	}
	if obj.obj.EnableRetransmissionAckTimeout == nil {
		obj.SetEnableRetransmissionAckTimeout(true)
	}
	if obj.obj.AckTimeout == nil {
		obj.SetAckTimeout(1)
	}
	if obj.obj.RetransmissionRetryCount == nil {
		obj.SetRetransmissionRetryCount(3)
	}

}
