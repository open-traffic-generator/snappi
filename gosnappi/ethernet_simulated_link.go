package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** EthernetSimulatedLink *****
type ethernetSimulatedLink struct {
	validation
	obj          *otg.EthernetSimulatedLink
	marshaller   marshalEthernetSimulatedLink
	unMarshaller unMarshalEthernetSimulatedLink
}

func NewEthernetSimulatedLink() EthernetSimulatedLink {
	obj := ethernetSimulatedLink{obj: &otg.EthernetSimulatedLink{}}
	obj.setDefault()
	return &obj
}

func (obj *ethernetSimulatedLink) msg() *otg.EthernetSimulatedLink {
	return obj.obj
}

func (obj *ethernetSimulatedLink) setMsg(msg *otg.EthernetSimulatedLink) EthernetSimulatedLink {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalethernetSimulatedLink struct {
	obj *ethernetSimulatedLink
}

type marshalEthernetSimulatedLink interface {
	// ToProto marshals EthernetSimulatedLink to protobuf object *otg.EthernetSimulatedLink
	ToProto() (*otg.EthernetSimulatedLink, error)
	// ToPbText marshals EthernetSimulatedLink to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals EthernetSimulatedLink to YAML text
	ToYaml() (string, error)
	// ToJson marshals EthernetSimulatedLink to JSON text
	ToJson() (string, error)
}

type unMarshalethernetSimulatedLink struct {
	obj *ethernetSimulatedLink
}

type unMarshalEthernetSimulatedLink interface {
	// FromProto unmarshals EthernetSimulatedLink from protobuf object *otg.EthernetSimulatedLink
	FromProto(msg *otg.EthernetSimulatedLink) (EthernetSimulatedLink, error)
	// FromPbText unmarshals EthernetSimulatedLink from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals EthernetSimulatedLink from YAML text
	FromYaml(value string) error
	// FromJson unmarshals EthernetSimulatedLink from JSON text
	FromJson(value string) error
}

func (obj *ethernetSimulatedLink) Marshal() marshalEthernetSimulatedLink {
	if obj.marshaller == nil {
		obj.marshaller = &marshalethernetSimulatedLink{obj: obj}
	}
	return obj.marshaller
}

func (obj *ethernetSimulatedLink) Unmarshal() unMarshalEthernetSimulatedLink {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalethernetSimulatedLink{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalethernetSimulatedLink) ToProto() (*otg.EthernetSimulatedLink, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalethernetSimulatedLink) FromProto(msg *otg.EthernetSimulatedLink) (EthernetSimulatedLink, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalethernetSimulatedLink) ToPbText() (string, error) {
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

func (m *unMarshalethernetSimulatedLink) FromPbText(value string) error {
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

func (m *marshalethernetSimulatedLink) ToYaml() (string, error) {
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

func (m *unMarshalethernetSimulatedLink) FromYaml(value string) error {
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

func (m *marshalethernetSimulatedLink) ToJson() (string, error) {
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

func (m *unMarshalethernetSimulatedLink) FromJson(value string) error {
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

func (obj *ethernetSimulatedLink) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ethernetSimulatedLink) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ethernetSimulatedLink) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ethernetSimulatedLink) Clone() (EthernetSimulatedLink, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewEthernetSimulatedLink()
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

// EthernetSimulatedLink is details of the internal link which can be used to create simulated device topologies behind an emulated router. MAC, VLAN and MTU information for the internal links are not used for purposes of emulating Simulated Topologies ( e.g. by ISIS Emulated Router behind which this is configured )
type EthernetSimulatedLink interface {
	Validation
	// msg marshals EthernetSimulatedLink to protobuf object *otg.EthernetSimulatedLink
	// and doesn't set defaults
	msg() *otg.EthernetSimulatedLink
	// setMsg unmarshals EthernetSimulatedLink from protobuf object *otg.EthernetSimulatedLink
	// and doesn't set defaults
	setMsg(*otg.EthernetSimulatedLink) EthernetSimulatedLink
	// provides marshal interface
	Marshal() marshalEthernetSimulatedLink
	// provides unmarshal interface
	Unmarshal() unMarshalEthernetSimulatedLink
	// validate validates EthernetSimulatedLink
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (EthernetSimulatedLink, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RemoteSimulatedLink returns string, set in EthernetSimulatedLink.
	RemoteSimulatedLink() string
	// SetRemoteSimulatedLink assigns string provided by user to EthernetSimulatedLink
	SetRemoteSimulatedLink(value string) EthernetSimulatedLink
	// HasRemoteSimulatedLink checks if RemoteSimulatedLink has been set in EthernetSimulatedLink
	HasRemoteSimulatedLink() bool
	// LinkType returns EthernetSimulatedLinkLinkTypeEnum, set in EthernetSimulatedLink
	LinkType() EthernetSimulatedLinkLinkTypeEnum
	// SetLinkType assigns EthernetSimulatedLinkLinkTypeEnum provided by user to EthernetSimulatedLink
	SetLinkType(value EthernetSimulatedLinkLinkTypeEnum) EthernetSimulatedLink
	// HasLinkType checks if LinkType has been set in EthernetSimulatedLink
	HasLinkType() bool
}

// Name of the remote end of the simulated interface which also must be a simulated_link on a device which might be acting either as an unconnected device in a simulated topology
// ( all ethernet links of type simulated_link ) or an emulated device connected to the Device Under Test (has at atleast one ethernet interface with connection to the port or
// lag connected to the DUT)
//
// x-constraint:
// - #/components/schemas/Device.Ethernet/properties/name
//
// x-constraint:
// - #/components/schemas/Device.Ethernet/properties/name
//
// RemoteSimulatedLink returns a string
func (obj *ethernetSimulatedLink) RemoteSimulatedLink() string {

	return *obj.obj.RemoteSimulatedLink

}

// Name of the remote end of the simulated interface which also must be a simulated_link on a device which might be acting either as an unconnected device in a simulated topology
// ( all ethernet links of type simulated_link ) or an emulated device connected to the Device Under Test (has at atleast one ethernet interface with connection to the port or
// lag connected to the DUT)
//
// x-constraint:
// - #/components/schemas/Device.Ethernet/properties/name
//
// x-constraint:
// - #/components/schemas/Device.Ethernet/properties/name
//
// RemoteSimulatedLink returns a string
func (obj *ethernetSimulatedLink) HasRemoteSimulatedLink() bool {
	return obj.obj.RemoteSimulatedLink != nil
}

// Name of the remote end of the simulated interface which also must be a simulated_link on a device which might be acting either as an unconnected device in a simulated topology
// ( all ethernet links of type simulated_link ) or an emulated device connected to the Device Under Test (has at atleast one ethernet interface with connection to the port or
// lag connected to the DUT)
//
// x-constraint:
// - #/components/schemas/Device.Ethernet/properties/name
//
// x-constraint:
// - #/components/schemas/Device.Ethernet/properties/name
//
// SetRemoteSimulatedLink sets the string value in the EthernetSimulatedLink object
func (obj *ethernetSimulatedLink) SetRemoteSimulatedLink(value string) EthernetSimulatedLink {

	obj.obj.RemoteSimulatedLink = &value
	return obj
}

type EthernetSimulatedLinkLinkTypeEnum string

// Enum of LinkType on EthernetSimulatedLink
var EthernetSimulatedLinkLinkType = struct {
	PRIMARY   EthernetSimulatedLinkLinkTypeEnum
	SECONDARY EthernetSimulatedLinkLinkTypeEnum
}{
	PRIMARY:   EthernetSimulatedLinkLinkTypeEnum("primary"),
	SECONDARY: EthernetSimulatedLinkLinkTypeEnum("secondary"),
}

func (obj *ethernetSimulatedLink) LinkType() EthernetSimulatedLinkLinkTypeEnum {
	return EthernetSimulatedLinkLinkTypeEnum(obj.obj.LinkType.Enum().String())
}

// By default, simulated links are treated as Primary links , which means that the intention is for connected device to advertise this and full topology of devices connected to it.
// e.g. when advertised as ISIS Simulated Topology.
//
// All simulated links inside one topology subset would normally can point to only other unconnected devices in the same topology or to the 'root' emulated device.
// If a link is designated as secondary , only that link information will be advertised by the IGP e.g. ISIS , and not the entire topology behind it.
// The optional secondary option allows emulation of external link scenarios where a simulated device (e.g. part of a ISIS simulated topology ) is advertised as reachable part of the topology
// by the emulated router behind which this is configured , as well as the other end of the secondary link which could be
// - 1) either a simulated device behind a different emulated router.
// - 2) or an emulated router on same or different port.
// This allows emulation of scenarios where one device/router is emulated to be reachable from different Emulated Routers connected to the Device Under Test. (e.g. for FRR scenarios)
//
// If an implementation does not support multiple primary links from same simulated topology i.e. full topology advertised via multiple emulated routers, it should return an error
// during set_config operation with such a topology.
// LinkType returns a string
func (obj *ethernetSimulatedLink) HasLinkType() bool {
	return obj.obj.LinkType != nil
}

func (obj *ethernetSimulatedLink) SetLinkType(value EthernetSimulatedLinkLinkTypeEnum) EthernetSimulatedLink {
	intValue, ok := otg.EthernetSimulatedLink_LinkType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on EthernetSimulatedLinkLinkTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.EthernetSimulatedLink_LinkType_Enum(intValue)
	obj.obj.LinkType = &enumValue

	return obj
}

func (obj *ethernetSimulatedLink) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ethernetSimulatedLink) setDefault() {
	if obj.obj.LinkType == nil {
		obj.SetLinkType(EthernetSimulatedLinkLinkType.PRIMARY)

	}

}
