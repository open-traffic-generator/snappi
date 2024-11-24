package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LldpCapabilityState *****
type lldpCapabilityState struct {
	validation
	obj          *otg.LldpCapabilityState
	marshaller   marshalLldpCapabilityState
	unMarshaller unMarshalLldpCapabilityState
}

func NewLldpCapabilityState() LldpCapabilityState {
	obj := lldpCapabilityState{obj: &otg.LldpCapabilityState{}}
	obj.setDefault()
	return &obj
}

func (obj *lldpCapabilityState) msg() *otg.LldpCapabilityState {
	return obj.obj
}

func (obj *lldpCapabilityState) setMsg(msg *otg.LldpCapabilityState) LldpCapabilityState {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallldpCapabilityState struct {
	obj *lldpCapabilityState
}

type marshalLldpCapabilityState interface {
	// ToProto marshals LldpCapabilityState to protobuf object *otg.LldpCapabilityState
	ToProto() (*otg.LldpCapabilityState, error)
	// ToPbText marshals LldpCapabilityState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LldpCapabilityState to YAML text
	ToYaml() (string, error)
	// ToJson marshals LldpCapabilityState to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals LldpCapabilityState to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshallldpCapabilityState struct {
	obj *lldpCapabilityState
}

type unMarshalLldpCapabilityState interface {
	// FromProto unmarshals LldpCapabilityState from protobuf object *otg.LldpCapabilityState
	FromProto(msg *otg.LldpCapabilityState) (LldpCapabilityState, error)
	// FromPbText unmarshals LldpCapabilityState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LldpCapabilityState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LldpCapabilityState from JSON text
	FromJson(value string) error
}

func (obj *lldpCapabilityState) Marshal() marshalLldpCapabilityState {
	if obj.marshaller == nil {
		obj.marshaller = &marshallldpCapabilityState{obj: obj}
	}
	return obj.marshaller
}

func (obj *lldpCapabilityState) Unmarshal() unMarshalLldpCapabilityState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallldpCapabilityState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallldpCapabilityState) ToProto() (*otg.LldpCapabilityState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallldpCapabilityState) FromProto(msg *otg.LldpCapabilityState) (LldpCapabilityState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallldpCapabilityState) ToPbText() (string, error) {
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

func (m *unMarshallldpCapabilityState) FromPbText(value string) error {
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

func (m *marshallldpCapabilityState) ToYaml() (string, error) {
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

func (m *unMarshallldpCapabilityState) FromYaml(value string) error {
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

func (m *marshallldpCapabilityState) ToJsonRaw() (string, error) {
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

func (m *marshallldpCapabilityState) ToJson() (string, error) {
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

func (m *unMarshallldpCapabilityState) FromJson(value string) error {
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

func (obj *lldpCapabilityState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lldpCapabilityState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lldpCapabilityState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lldpCapabilityState) Clone() (LldpCapabilityState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLldpCapabilityState()
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

// LldpCapabilityState is lLDP system capability advertised by the neighbor
type LldpCapabilityState interface {
	Validation
	// msg marshals LldpCapabilityState to protobuf object *otg.LldpCapabilityState
	// and doesn't set defaults
	msg() *otg.LldpCapabilityState
	// setMsg unmarshals LldpCapabilityState from protobuf object *otg.LldpCapabilityState
	// and doesn't set defaults
	setMsg(*otg.LldpCapabilityState) LldpCapabilityState
	// provides marshal interface
	Marshal() marshalLldpCapabilityState
	// provides unmarshal interface
	Unmarshal() unMarshalLldpCapabilityState
	// validate validates LldpCapabilityState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LldpCapabilityState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// CapabilityName returns LldpCapabilityStateCapabilityNameEnum, set in LldpCapabilityState
	CapabilityName() LldpCapabilityStateCapabilityNameEnum
	// SetCapabilityName assigns LldpCapabilityStateCapabilityNameEnum provided by user to LldpCapabilityState
	SetCapabilityName(value LldpCapabilityStateCapabilityNameEnum) LldpCapabilityState
	// HasCapabilityName checks if CapabilityName has been set in LldpCapabilityState
	HasCapabilityName() bool
	// CapabilityEnabled returns bool, set in LldpCapabilityState.
	CapabilityEnabled() bool
	// SetCapabilityEnabled assigns bool provided by user to LldpCapabilityState
	SetCapabilityEnabled(value bool) LldpCapabilityState
	// HasCapabilityEnabled checks if CapabilityEnabled has been set in LldpCapabilityState
	HasCapabilityEnabled() bool
}

type LldpCapabilityStateCapabilityNameEnum string

// Enum of CapabilityName on LldpCapabilityState
var LldpCapabilityStateCapabilityName = struct {
	MAC_BRIDGE          LldpCapabilityStateCapabilityNameEnum
	TWO_PORT_MAC_RELAY  LldpCapabilityStateCapabilityNameEnum
	REPEATER            LldpCapabilityStateCapabilityNameEnum
	DOCSIS_CABLE_DEVICE LldpCapabilityStateCapabilityNameEnum
	S_VLAN              LldpCapabilityStateCapabilityNameEnum
	TELEPHONE           LldpCapabilityStateCapabilityNameEnum
	OTHER               LldpCapabilityStateCapabilityNameEnum
	ROUTER              LldpCapabilityStateCapabilityNameEnum
	C_VLAN              LldpCapabilityStateCapabilityNameEnum
	STATION_ONLY        LldpCapabilityStateCapabilityNameEnum
	WLAN_ACCESS_POINT   LldpCapabilityStateCapabilityNameEnum
}{
	MAC_BRIDGE:          LldpCapabilityStateCapabilityNameEnum("mac_bridge"),
	TWO_PORT_MAC_RELAY:  LldpCapabilityStateCapabilityNameEnum("two_port_mac_relay"),
	REPEATER:            LldpCapabilityStateCapabilityNameEnum("repeater"),
	DOCSIS_CABLE_DEVICE: LldpCapabilityStateCapabilityNameEnum("docsis_cable_device"),
	S_VLAN:              LldpCapabilityStateCapabilityNameEnum("s_vlan"),
	TELEPHONE:           LldpCapabilityStateCapabilityNameEnum("telephone"),
	OTHER:               LldpCapabilityStateCapabilityNameEnum("other"),
	ROUTER:              LldpCapabilityStateCapabilityNameEnum("router"),
	C_VLAN:              LldpCapabilityStateCapabilityNameEnum("c_vlan"),
	STATION_ONLY:        LldpCapabilityStateCapabilityNameEnum("station_only"),
	WLAN_ACCESS_POINT:   LldpCapabilityStateCapabilityNameEnum("wlan_access_point"),
}

func (obj *lldpCapabilityState) CapabilityName() LldpCapabilityStateCapabilityNameEnum {
	return LldpCapabilityStateCapabilityNameEnum(obj.obj.CapabilityName.Enum().String())
}

// Name of the system capability advertised by the neighbor. Capabilities are represented in a bitmap that defines the primary functions of the system.  The capabilities are defined in IEEE 802.1AB.
// CapabilityName returns a string
func (obj *lldpCapabilityState) HasCapabilityName() bool {
	return obj.obj.CapabilityName != nil
}

func (obj *lldpCapabilityState) SetCapabilityName(value LldpCapabilityStateCapabilityNameEnum) LldpCapabilityState {
	intValue, ok := otg.LldpCapabilityState_CapabilityName_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on LldpCapabilityStateCapabilityNameEnum", string(value)))
		return obj
	}
	enumValue := otg.LldpCapabilityState_CapabilityName_Enum(intValue)
	obj.obj.CapabilityName = &enumValue

	return obj
}

// Indicates whether the corresponding system capability is enabled on the neighbor.
// CapabilityEnabled returns a bool
func (obj *lldpCapabilityState) CapabilityEnabled() bool {

	return *obj.obj.CapabilityEnabled

}

// Indicates whether the corresponding system capability is enabled on the neighbor.
// CapabilityEnabled returns a bool
func (obj *lldpCapabilityState) HasCapabilityEnabled() bool {
	return obj.obj.CapabilityEnabled != nil
}

// Indicates whether the corresponding system capability is enabled on the neighbor.
// SetCapabilityEnabled sets the bool value in the LldpCapabilityState object
func (obj *lldpCapabilityState) SetCapabilityEnabled(value bool) LldpCapabilityState {

	obj.obj.CapabilityEnabled = &value
	return obj
}

func (obj *lldpCapabilityState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *lldpCapabilityState) setDefault() {

}
