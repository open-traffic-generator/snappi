package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LdpIpv4BindingState *****
type ldpIpv4BindingState struct {
	validation
	obj          *otg.LdpIpv4BindingState
	marshaller   marshalLdpIpv4BindingState
	unMarshaller unMarshalLdpIpv4BindingState
}

func NewLdpIpv4BindingState() LdpIpv4BindingState {
	obj := ldpIpv4BindingState{obj: &otg.LdpIpv4BindingState{}}
	obj.setDefault()
	return &obj
}

func (obj *ldpIpv4BindingState) msg() *otg.LdpIpv4BindingState {
	return obj.obj
}

func (obj *ldpIpv4BindingState) setMsg(msg *otg.LdpIpv4BindingState) LdpIpv4BindingState {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalldpIpv4BindingState struct {
	obj *ldpIpv4BindingState
}

type marshalLdpIpv4BindingState interface {
	// ToProto marshals LdpIpv4BindingState to protobuf object *otg.LdpIpv4BindingState
	ToProto() (*otg.LdpIpv4BindingState, error)
	// ToPbText marshals LdpIpv4BindingState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LdpIpv4BindingState to YAML text
	ToYaml() (string, error)
	// ToJson marshals LdpIpv4BindingState to JSON text
	ToJson() (string, error)
}

type unMarshalldpIpv4BindingState struct {
	obj *ldpIpv4BindingState
}

type unMarshalLdpIpv4BindingState interface {
	// FromProto unmarshals LdpIpv4BindingState from protobuf object *otg.LdpIpv4BindingState
	FromProto(msg *otg.LdpIpv4BindingState) (LdpIpv4BindingState, error)
	// FromPbText unmarshals LdpIpv4BindingState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LdpIpv4BindingState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LdpIpv4BindingState from JSON text
	FromJson(value string) error
}

func (obj *ldpIpv4BindingState) Marshal() marshalLdpIpv4BindingState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalldpIpv4BindingState{obj: obj}
	}
	return obj.marshaller
}

func (obj *ldpIpv4BindingState) Unmarshal() unMarshalLdpIpv4BindingState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalldpIpv4BindingState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalldpIpv4BindingState) ToProto() (*otg.LdpIpv4BindingState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalldpIpv4BindingState) FromProto(msg *otg.LdpIpv4BindingState) (LdpIpv4BindingState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalldpIpv4BindingState) ToPbText() (string, error) {
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

func (m *unMarshalldpIpv4BindingState) FromPbText(value string) error {
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

func (m *marshalldpIpv4BindingState) ToYaml() (string, error) {
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

func (m *unMarshalldpIpv4BindingState) FromYaml(value string) error {
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

func (m *marshalldpIpv4BindingState) ToJson() (string, error) {
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

func (m *unMarshalldpIpv4BindingState) FromJson(value string) error {
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

func (obj *ldpIpv4BindingState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ldpIpv4BindingState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ldpIpv4BindingState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ldpIpv4BindingState) Clone() (LdpIpv4BindingState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLdpIpv4BindingState()
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

// LdpIpv4BindingState is one IPv4 Prefix FEC element (RFC 5036 Section 3.4.1) and the label bound to it by a peer (RFC 5036 Section 3.4.2.1).
type LdpIpv4BindingState interface {
	Validation
	// msg marshals LdpIpv4BindingState to protobuf object *otg.LdpIpv4BindingState
	// and doesn't set defaults
	msg() *otg.LdpIpv4BindingState
	// setMsg unmarshals LdpIpv4BindingState from protobuf object *otg.LdpIpv4BindingState
	// and doesn't set defaults
	setMsg(*otg.LdpIpv4BindingState) LdpIpv4BindingState
	// provides marshal interface
	Marshal() marshalLdpIpv4BindingState
	// provides unmarshal interface
	Unmarshal() unMarshalLdpIpv4BindingState
	// validate validates LdpIpv4BindingState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LdpIpv4BindingState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Prefix returns string, set in LdpIpv4BindingState.
	Prefix() string
	// SetPrefix assigns string provided by user to LdpIpv4BindingState
	SetPrefix(value string) LdpIpv4BindingState
	// HasPrefix checks if Prefix has been set in LdpIpv4BindingState
	HasPrefix() bool
	// PrefixLength returns uint32, set in LdpIpv4BindingState.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to LdpIpv4BindingState
	SetPrefixLength(value uint32) LdpIpv4BindingState
	// HasPrefixLength checks if PrefixLength has been set in LdpIpv4BindingState
	HasPrefixLength() bool
	// Label returns uint32, set in LdpIpv4BindingState.
	Label() uint32
	// SetLabel assigns uint32 provided by user to LdpIpv4BindingState
	SetLabel(value uint32) LdpIpv4BindingState
	// HasLabel checks if Label has been set in LdpIpv4BindingState
	HasLabel() bool
	// PeerLsrId returns string, set in LdpIpv4BindingState.
	PeerLsrId() string
	// SetPeerLsrId assigns string provided by user to LdpIpv4BindingState
	SetPeerLsrId(value string) LdpIpv4BindingState
	// HasPeerLsrId checks if PeerLsrId has been set in LdpIpv4BindingState
	HasPeerLsrId() bool
	// PeerLabelSpaceId returns uint32, set in LdpIpv4BindingState.
	PeerLabelSpaceId() uint32
	// SetPeerLabelSpaceId assigns uint32 provided by user to LdpIpv4BindingState
	SetPeerLabelSpaceId(value uint32) LdpIpv4BindingState
	// HasPeerLabelSpaceId checks if PeerLabelSpaceId has been set in LdpIpv4BindingState
	HasPeerLabelSpaceId() bool
}

// The IPv4 prefix of the FEC element.
// Prefix returns a string
func (obj *ldpIpv4BindingState) Prefix() string {

	return *obj.obj.Prefix

}

// The IPv4 prefix of the FEC element.
// Prefix returns a string
func (obj *ldpIpv4BindingState) HasPrefix() bool {
	return obj.obj.Prefix != nil
}

// The IPv4 prefix of the FEC element.
// SetPrefix sets the string value in the LdpIpv4BindingState object
func (obj *ldpIpv4BindingState) SetPrefix(value string) LdpIpv4BindingState {

	obj.obj.Prefix = &value
	return obj
}

// The prefix length of the FEC element.
// PrefixLength returns a uint32
func (obj *ldpIpv4BindingState) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// The prefix length of the FEC element.
// PrefixLength returns a uint32
func (obj *ldpIpv4BindingState) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// The prefix length of the FEC element.
// SetPrefixLength sets the uint32 value in the LdpIpv4BindingState object
func (obj *ldpIpv4BindingState) SetPrefixLength(value uint32) LdpIpv4BindingState {

	obj.obj.PrefixLength = &value
	return obj
}

// The label bound to the FEC by the peer. 3 is Implicit NULL and 0 is IPv4 Explicit NULL (RFC 3032 Section 2.1).
// Label returns a uint32
func (obj *ldpIpv4BindingState) Label() uint32 {

	return *obj.obj.Label

}

// The label bound to the FEC by the peer. 3 is Implicit NULL and 0 is IPv4 Explicit NULL (RFC 3032 Section 2.1).
// Label returns a uint32
func (obj *ldpIpv4BindingState) HasLabel() bool {
	return obj.obj.Label != nil
}

// The label bound to the FEC by the peer. 3 is Implicit NULL and 0 is IPv4 Explicit NULL (RFC 3032 Section 2.1).
// SetLabel sets the uint32 value in the LdpIpv4BindingState object
func (obj *ldpIpv4BindingState) SetLabel(value uint32) LdpIpv4BindingState {

	obj.obj.Label = &value
	return obj
}

// The LSR ID of the peer that advertised the binding (RFC 5036 Section 2.2.2).
// PeerLsrId returns a string
func (obj *ldpIpv4BindingState) PeerLsrId() string {

	return *obj.obj.PeerLsrId

}

// The LSR ID of the peer that advertised the binding (RFC 5036 Section 2.2.2).
// PeerLsrId returns a string
func (obj *ldpIpv4BindingState) HasPeerLsrId() bool {
	return obj.obj.PeerLsrId != nil
}

// The LSR ID of the peer that advertised the binding (RFC 5036 Section 2.2.2).
// SetPeerLsrId sets the string value in the LdpIpv4BindingState object
func (obj *ldpIpv4BindingState) SetPeerLsrId(value string) LdpIpv4BindingState {

	obj.obj.PeerLsrId = &value
	return obj
}

// The label space identifier of the peer that advertised the binding (RFC 5036 Section 2.2.2).
// PeerLabelSpaceId returns a uint32
func (obj *ldpIpv4BindingState) PeerLabelSpaceId() uint32 {

	return *obj.obj.PeerLabelSpaceId

}

// The label space identifier of the peer that advertised the binding (RFC 5036 Section 2.2.2).
// PeerLabelSpaceId returns a uint32
func (obj *ldpIpv4BindingState) HasPeerLabelSpaceId() bool {
	return obj.obj.PeerLabelSpaceId != nil
}

// The label space identifier of the peer that advertised the binding (RFC 5036 Section 2.2.2).
// SetPeerLabelSpaceId sets the uint32 value in the LdpIpv4BindingState object
func (obj *ldpIpv4BindingState) SetPeerLabelSpaceId(value uint32) LdpIpv4BindingState {

	obj.obj.PeerLabelSpaceId = &value
	return obj
}

func (obj *ldpIpv4BindingState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Prefix != nil {

		err := obj.validateIpv4(obj.Prefix())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on LdpIpv4BindingState.Prefix"))
		}

	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= LdpIpv4BindingState.PrefixLength <= 32 but Got %d", *obj.obj.PrefixLength))
		}

	}

	if obj.obj.Label != nil {

		if *obj.obj.Label > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= LdpIpv4BindingState.Label <= 1048575 but Got %d", *obj.obj.Label))
		}

	}

	if obj.obj.PeerLsrId != nil {

		err := obj.validateIpv4(obj.PeerLsrId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on LdpIpv4BindingState.PeerLsrId"))
		}

	}

	if obj.obj.PeerLabelSpaceId != nil {

		if *obj.obj.PeerLabelSpaceId > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= LdpIpv4BindingState.PeerLabelSpaceId <= 65535 but Got %d", *obj.obj.PeerLabelSpaceId))
		}

	}

}

func (obj *ldpIpv4BindingState) setDefault() {

}
