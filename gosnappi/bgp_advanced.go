package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAdvanced *****
type bgpAdvanced struct {
	validation
	obj          *otg.BgpAdvanced
	marshaller   marshalBgpAdvanced
	unMarshaller unMarshalBgpAdvanced
}

func NewBgpAdvanced() BgpAdvanced {
	obj := bgpAdvanced{obj: &otg.BgpAdvanced{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAdvanced) msg() *otg.BgpAdvanced {
	return obj.obj
}

func (obj *bgpAdvanced) setMsg(msg *otg.BgpAdvanced) BgpAdvanced {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAdvanced struct {
	obj *bgpAdvanced
}

type marshalBgpAdvanced interface {
	// ToProto marshals BgpAdvanced to protobuf object *otg.BgpAdvanced
	ToProto() (*otg.BgpAdvanced, error)
	// ToPbText marshals BgpAdvanced to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAdvanced to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAdvanced to JSON text
	ToJson() (string, error)
}

type unMarshalbgpAdvanced struct {
	obj *bgpAdvanced
}

type unMarshalBgpAdvanced interface {
	// FromProto unmarshals BgpAdvanced from protobuf object *otg.BgpAdvanced
	FromProto(msg *otg.BgpAdvanced) (BgpAdvanced, error)
	// FromPbText unmarshals BgpAdvanced from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAdvanced from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAdvanced from JSON text
	FromJson(value string) error
}

func (obj *bgpAdvanced) Marshal() marshalBgpAdvanced {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAdvanced{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAdvanced) Unmarshal() unMarshalBgpAdvanced {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAdvanced{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAdvanced) ToProto() (*otg.BgpAdvanced, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAdvanced) FromProto(msg *otg.BgpAdvanced) (BgpAdvanced, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAdvanced) ToPbText() (string, error) {
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

func (m *unMarshalbgpAdvanced) FromPbText(value string) error {
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

func (m *marshalbgpAdvanced) ToYaml() (string, error) {
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

func (m *unMarshalbgpAdvanced) FromYaml(value string) error {
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

func (m *marshalbgpAdvanced) ToJson() (string, error) {
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

func (m *unMarshalbgpAdvanced) FromJson(value string) error {
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

func (obj *bgpAdvanced) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAdvanced) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAdvanced) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAdvanced) Clone() (BgpAdvanced, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAdvanced()
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

// BgpAdvanced is configuration for BGP advanced settings.
type BgpAdvanced interface {
	Validation
	// msg marshals BgpAdvanced to protobuf object *otg.BgpAdvanced
	// and doesn't set defaults
	msg() *otg.BgpAdvanced
	// setMsg unmarshals BgpAdvanced from protobuf object *otg.BgpAdvanced
	// and doesn't set defaults
	setMsg(*otg.BgpAdvanced) BgpAdvanced
	// provides marshal interface
	Marshal() marshalBgpAdvanced
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAdvanced
	// validate validates BgpAdvanced
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAdvanced, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// HoldTimeInterval returns uint32, set in BgpAdvanced.
	HoldTimeInterval() uint32
	// SetHoldTimeInterval assigns uint32 provided by user to BgpAdvanced
	SetHoldTimeInterval(value uint32) BgpAdvanced
	// HasHoldTimeInterval checks if HoldTimeInterval has been set in BgpAdvanced
	HasHoldTimeInterval() bool
	// KeepAliveInterval returns uint32, set in BgpAdvanced.
	KeepAliveInterval() uint32
	// SetKeepAliveInterval assigns uint32 provided by user to BgpAdvanced
	SetKeepAliveInterval(value uint32) BgpAdvanced
	// HasKeepAliveInterval checks if KeepAliveInterval has been set in BgpAdvanced
	HasKeepAliveInterval() bool
	// UpdateInterval returns uint32, set in BgpAdvanced.
	UpdateInterval() uint32
	// SetUpdateInterval assigns uint32 provided by user to BgpAdvanced
	SetUpdateInterval(value uint32) BgpAdvanced
	// HasUpdateInterval checks if UpdateInterval has been set in BgpAdvanced
	HasUpdateInterval() bool
	// TimeToLive returns uint32, set in BgpAdvanced.
	TimeToLive() uint32
	// SetTimeToLive assigns uint32 provided by user to BgpAdvanced
	SetTimeToLive(value uint32) BgpAdvanced
	// HasTimeToLive checks if TimeToLive has been set in BgpAdvanced
	HasTimeToLive() bool
	// Md5Key returns string, set in BgpAdvanced.
	Md5Key() string
	// SetMd5Key assigns string provided by user to BgpAdvanced
	SetMd5Key(value string) BgpAdvanced
	// HasMd5Key checks if Md5Key has been set in BgpAdvanced
	HasMd5Key() bool
	// PassiveMode returns bool, set in BgpAdvanced.
	PassiveMode() bool
	// SetPassiveMode assigns bool provided by user to BgpAdvanced
	SetPassiveMode(value bool) BgpAdvanced
	// HasPassiveMode checks if PassiveMode has been set in BgpAdvanced
	HasPassiveMode() bool
	// ListenPort returns uint32, set in BgpAdvanced.
	ListenPort() uint32
	// SetListenPort assigns uint32 provided by user to BgpAdvanced
	SetListenPort(value uint32) BgpAdvanced
	// HasListenPort checks if ListenPort has been set in BgpAdvanced
	HasListenPort() bool
	// NeighborPort returns uint32, set in BgpAdvanced.
	NeighborPort() uint32
	// SetNeighborPort assigns uint32 provided by user to BgpAdvanced
	SetNeighborPort(value uint32) BgpAdvanced
	// HasNeighborPort checks if NeighborPort has been set in BgpAdvanced
	HasNeighborPort() bool
}

// Number of seconds the sender proposes for the value of the Hold Timer.
// HoldTimeInterval returns a uint32
func (obj *bgpAdvanced) HoldTimeInterval() uint32 {

	return *obj.obj.HoldTimeInterval

}

// Number of seconds the sender proposes for the value of the Hold Timer.
// HoldTimeInterval returns a uint32
func (obj *bgpAdvanced) HasHoldTimeInterval() bool {
	return obj.obj.HoldTimeInterval != nil
}

// Number of seconds the sender proposes for the value of the Hold Timer.
// SetHoldTimeInterval sets the uint32 value in the BgpAdvanced object
func (obj *bgpAdvanced) SetHoldTimeInterval(value uint32) BgpAdvanced {

	obj.obj.HoldTimeInterval = &value
	return obj
}

// Number of seconds between transmissions of Keepalive messages by this peer.
// KeepAliveInterval returns a uint32
func (obj *bgpAdvanced) KeepAliveInterval() uint32 {

	return *obj.obj.KeepAliveInterval

}

// Number of seconds between transmissions of Keepalive messages by this peer.
// KeepAliveInterval returns a uint32
func (obj *bgpAdvanced) HasKeepAliveInterval() bool {
	return obj.obj.KeepAliveInterval != nil
}

// Number of seconds between transmissions of Keepalive messages by this peer.
// SetKeepAliveInterval sets the uint32 value in the BgpAdvanced object
func (obj *bgpAdvanced) SetKeepAliveInterval(value uint32) BgpAdvanced {

	obj.obj.KeepAliveInterval = &value
	return obj
}

// The time interval at which Update messages are sent to the DUT, expressed as the number of milliseconds between Update messages. The update interval 0 implies to send all the updates as fast as possible.
// UpdateInterval returns a uint32
func (obj *bgpAdvanced) UpdateInterval() uint32 {

	return *obj.obj.UpdateInterval

}

// The time interval at which Update messages are sent to the DUT, expressed as the number of milliseconds between Update messages. The update interval 0 implies to send all the updates as fast as possible.
// UpdateInterval returns a uint32
func (obj *bgpAdvanced) HasUpdateInterval() bool {
	return obj.obj.UpdateInterval != nil
}

// The time interval at which Update messages are sent to the DUT, expressed as the number of milliseconds between Update messages. The update interval 0 implies to send all the updates as fast as possible.
// SetUpdateInterval sets the uint32 value in the BgpAdvanced object
func (obj *bgpAdvanced) SetUpdateInterval(value uint32) BgpAdvanced {

	obj.obj.UpdateInterval = &value
	return obj
}

// The limited number of iterations that a unit of data can experience before the data is discarded. This is placed in the TTL field in the IP header of the  transmitted packets.
// TimeToLive returns a uint32
func (obj *bgpAdvanced) TimeToLive() uint32 {

	return *obj.obj.TimeToLive

}

// The limited number of iterations that a unit of data can experience before the data is discarded. This is placed in the TTL field in the IP header of the  transmitted packets.
// TimeToLive returns a uint32
func (obj *bgpAdvanced) HasTimeToLive() bool {
	return obj.obj.TimeToLive != nil
}

// The limited number of iterations that a unit of data can experience before the data is discarded. This is placed in the TTL field in the IP header of the  transmitted packets.
// SetTimeToLive sets the uint32 value in the BgpAdvanced object
func (obj *bgpAdvanced) SetTimeToLive(value uint32) BgpAdvanced {

	obj.obj.TimeToLive = &value
	return obj
}

// The value to be used as a secret MD5 key for authentication. If not configured, MD5 authentication will not be enabled.
// Md5Key returns a string
func (obj *bgpAdvanced) Md5Key() string {

	return *obj.obj.Md5Key

}

// The value to be used as a secret MD5 key for authentication. If not configured, MD5 authentication will not be enabled.
// Md5Key returns a string
func (obj *bgpAdvanced) HasMd5Key() bool {
	return obj.obj.Md5Key != nil
}

// The value to be used as a secret MD5 key for authentication. If not configured, MD5 authentication will not be enabled.
// SetMd5Key sets the string value in the BgpAdvanced object
func (obj *bgpAdvanced) SetMd5Key(value string) BgpAdvanced {

	obj.obj.Md5Key = &value
	return obj
}

// If set to true, the local BGP peer will wait for the remote peer to initiate the BGP session
// by establishing the TCP connection, rather than initiating sessions from the local peer.
// PassiveMode returns a bool
func (obj *bgpAdvanced) PassiveMode() bool {

	return *obj.obj.PassiveMode

}

// If set to true, the local BGP peer will wait for the remote peer to initiate the BGP session
// by establishing the TCP connection, rather than initiating sessions from the local peer.
// PassiveMode returns a bool
func (obj *bgpAdvanced) HasPassiveMode() bool {
	return obj.obj.PassiveMode != nil
}

// If set to true, the local BGP peer will wait for the remote peer to initiate the BGP session
// by establishing the TCP connection, rather than initiating sessions from the local peer.
// SetPassiveMode sets the bool value in the BgpAdvanced object
func (obj *bgpAdvanced) SetPassiveMode(value bool) BgpAdvanced {

	obj.obj.PassiveMode = &value
	return obj
}

// The TCP port number on which to accept BGP connections from the remote peer.
// ListenPort returns a uint32
func (obj *bgpAdvanced) ListenPort() uint32 {

	return *obj.obj.ListenPort

}

// The TCP port number on which to accept BGP connections from the remote peer.
// ListenPort returns a uint32
func (obj *bgpAdvanced) HasListenPort() bool {
	return obj.obj.ListenPort != nil
}

// The TCP port number on which to accept BGP connections from the remote peer.
// SetListenPort sets the uint32 value in the BgpAdvanced object
func (obj *bgpAdvanced) SetListenPort(value uint32) BgpAdvanced {

	obj.obj.ListenPort = &value
	return obj
}

// Destination TCP port number of the BGP peer when initiating a
// session from the local BGP peer.
// NeighborPort returns a uint32
func (obj *bgpAdvanced) NeighborPort() uint32 {

	return *obj.obj.NeighborPort

}

// Destination TCP port number of the BGP peer when initiating a
// session from the local BGP peer.
// NeighborPort returns a uint32
func (obj *bgpAdvanced) HasNeighborPort() bool {
	return obj.obj.NeighborPort != nil
}

// Destination TCP port number of the BGP peer when initiating a
// session from the local BGP peer.
// SetNeighborPort sets the uint32 value in the BgpAdvanced object
func (obj *bgpAdvanced) SetNeighborPort(value uint32) BgpAdvanced {

	obj.obj.NeighborPort = &value
	return obj
}

func (obj *bgpAdvanced) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.TimeToLive != nil {

		if *obj.obj.TimeToLive > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpAdvanced.TimeToLive <= 255 but Got %d", *obj.obj.TimeToLive))
		}

	}

	if obj.obj.ListenPort != nil {

		if *obj.obj.ListenPort > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpAdvanced.ListenPort <= 65535 but Got %d", *obj.obj.ListenPort))
		}

	}

	if obj.obj.NeighborPort != nil {

		if *obj.obj.NeighborPort > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpAdvanced.NeighborPort <= 65535 but Got %d", *obj.obj.NeighborPort))
		}

	}

}

func (obj *bgpAdvanced) setDefault() {
	if obj.obj.HoldTimeInterval == nil {
		obj.SetHoldTimeInterval(90)
	}
	if obj.obj.KeepAliveInterval == nil {
		obj.SetKeepAliveInterval(30)
	}
	if obj.obj.UpdateInterval == nil {
		obj.SetUpdateInterval(0)
	}
	if obj.obj.TimeToLive == nil {
		obj.SetTimeToLive(64)
	}
	if obj.obj.PassiveMode == nil {
		obj.SetPassiveMode(false)
	}
	if obj.obj.ListenPort == nil {
		obj.SetListenPort(179)
	}
	if obj.obj.NeighborPort == nil {
		obj.SetNeighborPort(179)
	}

}
