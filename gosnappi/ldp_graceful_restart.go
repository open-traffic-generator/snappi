package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LdpGracefulRestart *****
type ldpGracefulRestart struct {
	validation
	obj          *otg.LdpGracefulRestart
	marshaller   marshalLdpGracefulRestart
	unMarshaller unMarshalLdpGracefulRestart
}

func NewLdpGracefulRestart() LdpGracefulRestart {
	obj := ldpGracefulRestart{obj: &otg.LdpGracefulRestart{}}
	obj.setDefault()
	return &obj
}

func (obj *ldpGracefulRestart) msg() *otg.LdpGracefulRestart {
	return obj.obj
}

func (obj *ldpGracefulRestart) setMsg(msg *otg.LdpGracefulRestart) LdpGracefulRestart {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalldpGracefulRestart struct {
	obj *ldpGracefulRestart
}

type marshalLdpGracefulRestart interface {
	// ToProto marshals LdpGracefulRestart to protobuf object *otg.LdpGracefulRestart
	ToProto() (*otg.LdpGracefulRestart, error)
	// ToPbText marshals LdpGracefulRestart to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LdpGracefulRestart to YAML text
	ToYaml() (string, error)
	// ToJson marshals LdpGracefulRestart to JSON text
	ToJson() (string, error)
}

type unMarshalldpGracefulRestart struct {
	obj *ldpGracefulRestart
}

type unMarshalLdpGracefulRestart interface {
	// FromProto unmarshals LdpGracefulRestart from protobuf object *otg.LdpGracefulRestart
	FromProto(msg *otg.LdpGracefulRestart) (LdpGracefulRestart, error)
	// FromPbText unmarshals LdpGracefulRestart from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LdpGracefulRestart from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LdpGracefulRestart from JSON text
	FromJson(value string) error
}

func (obj *ldpGracefulRestart) Marshal() marshalLdpGracefulRestart {
	if obj.marshaller == nil {
		obj.marshaller = &marshalldpGracefulRestart{obj: obj}
	}
	return obj.marshaller
}

func (obj *ldpGracefulRestart) Unmarshal() unMarshalLdpGracefulRestart {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalldpGracefulRestart{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalldpGracefulRestart) ToProto() (*otg.LdpGracefulRestart, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalldpGracefulRestart) FromProto(msg *otg.LdpGracefulRestart) (LdpGracefulRestart, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalldpGracefulRestart) ToPbText() (string, error) {
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

func (m *unMarshalldpGracefulRestart) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalldpGracefulRestart) ToYaml() (string, error) {
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

func (m *unMarshalldpGracefulRestart) FromYaml(value string) error {
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

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalldpGracefulRestart) ToJson() (string, error) {
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

func (m *unMarshalldpGracefulRestart) FromJson(value string) error {
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

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *ldpGracefulRestart) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ldpGracefulRestart) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ldpGracefulRestart) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ldpGracefulRestart) Clone() (LdpGracefulRestart, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLdpGracefulRestart()
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

// LdpGracefulRestart is lDP graceful restart parameters (RFC 3478). When this object is present, the emulated LSR includes the Fault Tolerant (FT) Session TLV (type 0x0503) with the L flag set in its Initialization message (RFC 3478 Section 2). Both timers are configured in seconds; they are carried in milliseconds on the wire, and the implementation converts them.
type LdpGracefulRestart interface {
	Validation
	// msg marshals LdpGracefulRestart to protobuf object *otg.LdpGracefulRestart
	// and doesn't set defaults
	msg() *otg.LdpGracefulRestart
	// setMsg unmarshals LdpGracefulRestart from protobuf object *otg.LdpGracefulRestart
	// and doesn't set defaults
	setMsg(*otg.LdpGracefulRestart) LdpGracefulRestart
	// provides marshal interface
	Marshal() marshalLdpGracefulRestart
	// provides unmarshal interface
	Unmarshal() unMarshalLdpGracefulRestart
	// validate validates LdpGracefulRestart
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LdpGracefulRestart, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ReconnectTime returns uint32, set in LdpGracefulRestart.
	ReconnectTime() uint32
	// SetReconnectTime assigns uint32 provided by user to LdpGracefulRestart
	SetReconnectTime(value uint32) LdpGracefulRestart
	// HasReconnectTime checks if ReconnectTime has been set in LdpGracefulRestart
	HasReconnectTime() bool
	// RecoveryTime returns uint32, set in LdpGracefulRestart.
	RecoveryTime() uint32
	// SetRecoveryTime assigns uint32 provided by user to LdpGracefulRestart
	SetRecoveryTime(value uint32) LdpGracefulRestart
	// HasRecoveryTime checks if RecoveryTime has been set in LdpGracefulRestart
	HasRecoveryTime() bool
}

// The FT Reconnect Timeout in seconds advertised in the FT Session TLV (RFC 3478 Section 2): the time the peer should wait for this LSR to re-establish the session after a restart. A value of 0 indicates that this LSR does not preserve its forwarding state across a restart.
// ReconnectTime returns a uint32
func (obj *ldpGracefulRestart) ReconnectTime() uint32 {

	return *obj.obj.ReconnectTime

}

// The FT Reconnect Timeout in seconds advertised in the FT Session TLV (RFC 3478 Section 2): the time the peer should wait for this LSR to re-establish the session after a restart. A value of 0 indicates that this LSR does not preserve its forwarding state across a restart.
// ReconnectTime returns a uint32
func (obj *ldpGracefulRestart) HasReconnectTime() bool {
	return obj.obj.ReconnectTime != nil
}

// The FT Reconnect Timeout in seconds advertised in the FT Session TLV (RFC 3478 Section 2): the time the peer should wait for this LSR to re-establish the session after a restart. A value of 0 indicates that this LSR does not preserve its forwarding state across a restart.
// SetReconnectTime sets the uint32 value in the LdpGracefulRestart object
func (obj *ldpGracefulRestart) SetReconnectTime(value uint32) LdpGracefulRestart {

	obj.obj.ReconnectTime = &value
	return obj
}

// The Recovery Time in seconds advertised in the FT Session TLV after a restart (RFC 3478 Section 2): the time for which this LSR retains the MPLS forwarding state it preserved across the restart. A value of 0 indicates that the forwarding state was not preserved.
// RecoveryTime returns a uint32
func (obj *ldpGracefulRestart) RecoveryTime() uint32 {

	return *obj.obj.RecoveryTime

}

// The Recovery Time in seconds advertised in the FT Session TLV after a restart (RFC 3478 Section 2): the time for which this LSR retains the MPLS forwarding state it preserved across the restart. A value of 0 indicates that the forwarding state was not preserved.
// RecoveryTime returns a uint32
func (obj *ldpGracefulRestart) HasRecoveryTime() bool {
	return obj.obj.RecoveryTime != nil
}

// The Recovery Time in seconds advertised in the FT Session TLV after a restart (RFC 3478 Section 2): the time for which this LSR retains the MPLS forwarding state it preserved across the restart. A value of 0 indicates that the forwarding state was not preserved.
// SetRecoveryTime sets the uint32 value in the LdpGracefulRestart object
func (obj *ldpGracefulRestart) SetRecoveryTime(value uint32) LdpGracefulRestart {

	obj.obj.RecoveryTime = &value
	return obj
}

func (obj *ldpGracefulRestart) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ReconnectTime != nil {

		if *obj.obj.ReconnectTime > 3600 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= LdpGracefulRestart.ReconnectTime <= 3600 but Got %d", *obj.obj.ReconnectTime))
		}

	}

	if obj.obj.RecoveryTime != nil {

		if *obj.obj.RecoveryTime > 3600 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= LdpGracefulRestart.RecoveryTime <= 3600 but Got %d", *obj.obj.RecoveryTime))
		}

	}

}

func (obj *ldpGracefulRestart) setDefault() {
	if obj.obj.ReconnectTime == nil {
		obj.SetReconnectTime(120)
	}
	if obj.obj.RecoveryTime == nil {
		obj.SetRecoveryTime(120)
	}

}
