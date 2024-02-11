package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisInterfaceAdvanced *****
type isisInterfaceAdvanced struct {
	validation
	obj          *otg.IsisInterfaceAdvanced
	marshaller   marshalIsisInterfaceAdvanced
	unMarshaller unMarshalIsisInterfaceAdvanced
}

func NewIsisInterfaceAdvanced() IsisInterfaceAdvanced {
	obj := isisInterfaceAdvanced{obj: &otg.IsisInterfaceAdvanced{}}
	obj.setDefault()
	return &obj
}

func (obj *isisInterfaceAdvanced) msg() *otg.IsisInterfaceAdvanced {
	return obj.obj
}

func (obj *isisInterfaceAdvanced) setMsg(msg *otg.IsisInterfaceAdvanced) IsisInterfaceAdvanced {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisInterfaceAdvanced struct {
	obj *isisInterfaceAdvanced
}

type marshalIsisInterfaceAdvanced interface {
	// ToProto marshals IsisInterfaceAdvanced to protobuf object *otg.IsisInterfaceAdvanced
	ToProto() (*otg.IsisInterfaceAdvanced, error)
	// ToPbText marshals IsisInterfaceAdvanced to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisInterfaceAdvanced to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisInterfaceAdvanced to JSON text
	ToJson() (string, error)
}

type unMarshalisisInterfaceAdvanced struct {
	obj *isisInterfaceAdvanced
}

type unMarshalIsisInterfaceAdvanced interface {
	// FromProto unmarshals IsisInterfaceAdvanced from protobuf object *otg.IsisInterfaceAdvanced
	FromProto(msg *otg.IsisInterfaceAdvanced) (IsisInterfaceAdvanced, error)
	// FromPbText unmarshals IsisInterfaceAdvanced from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisInterfaceAdvanced from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisInterfaceAdvanced from JSON text
	FromJson(value string) error
}

func (obj *isisInterfaceAdvanced) Marshal() marshalIsisInterfaceAdvanced {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisInterfaceAdvanced{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisInterfaceAdvanced) Unmarshal() unMarshalIsisInterfaceAdvanced {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisInterfaceAdvanced{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisInterfaceAdvanced) ToProto() (*otg.IsisInterfaceAdvanced, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisInterfaceAdvanced) FromProto(msg *otg.IsisInterfaceAdvanced) (IsisInterfaceAdvanced, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisInterfaceAdvanced) ToPbText() (string, error) {
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

func (m *unMarshalisisInterfaceAdvanced) FromPbText(value string) error {
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

func (m *marshalisisInterfaceAdvanced) ToYaml() (string, error) {
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

func (m *unMarshalisisInterfaceAdvanced) FromYaml(value string) error {
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

func (m *marshalisisInterfaceAdvanced) ToJson() (string, error) {
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

func (m *unMarshalisisInterfaceAdvanced) FromJson(value string) error {
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

func (obj *isisInterfaceAdvanced) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisInterfaceAdvanced) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisInterfaceAdvanced) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisInterfaceAdvanced) Clone() (IsisInterfaceAdvanced, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisInterfaceAdvanced()
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

// IsisInterfaceAdvanced is optional container for advanced interface properties.
type IsisInterfaceAdvanced interface {
	Validation
	// msg marshals IsisInterfaceAdvanced to protobuf object *otg.IsisInterfaceAdvanced
	// and doesn't set defaults
	msg() *otg.IsisInterfaceAdvanced
	// setMsg unmarshals IsisInterfaceAdvanced from protobuf object *otg.IsisInterfaceAdvanced
	// and doesn't set defaults
	setMsg(*otg.IsisInterfaceAdvanced) IsisInterfaceAdvanced
	// provides marshal interface
	Marshal() marshalIsisInterfaceAdvanced
	// provides unmarshal interface
	Unmarshal() unMarshalIsisInterfaceAdvanced
	// validate validates IsisInterfaceAdvanced
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisInterfaceAdvanced, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// AutoAdjustMtu returns bool, set in IsisInterfaceAdvanced.
	AutoAdjustMtu() bool
	// SetAutoAdjustMtu assigns bool provided by user to IsisInterfaceAdvanced
	SetAutoAdjustMtu(value bool) IsisInterfaceAdvanced
	// HasAutoAdjustMtu checks if AutoAdjustMtu has been set in IsisInterfaceAdvanced
	HasAutoAdjustMtu() bool
	// AutoAdjustArea returns bool, set in IsisInterfaceAdvanced.
	AutoAdjustArea() bool
	// SetAutoAdjustArea assigns bool provided by user to IsisInterfaceAdvanced
	SetAutoAdjustArea(value bool) IsisInterfaceAdvanced
	// HasAutoAdjustArea checks if AutoAdjustArea has been set in IsisInterfaceAdvanced
	HasAutoAdjustArea() bool
	// AutoAdjustSupportedProtocols returns bool, set in IsisInterfaceAdvanced.
	AutoAdjustSupportedProtocols() bool
	// SetAutoAdjustSupportedProtocols assigns bool provided by user to IsisInterfaceAdvanced
	SetAutoAdjustSupportedProtocols(value bool) IsisInterfaceAdvanced
	// HasAutoAdjustSupportedProtocols checks if AutoAdjustSupportedProtocols has been set in IsisInterfaceAdvanced
	HasAutoAdjustSupportedProtocols() bool
	// Enable3WayHandshake returns bool, set in IsisInterfaceAdvanced.
	Enable3WayHandshake() bool
	// SetEnable3WayHandshake assigns bool provided by user to IsisInterfaceAdvanced
	SetEnable3WayHandshake(value bool) IsisInterfaceAdvanced
	// HasEnable3WayHandshake checks if Enable3WayHandshake has been set in IsisInterfaceAdvanced
	HasEnable3WayHandshake() bool
	// P2PHellosToUnicastMac returns bool, set in IsisInterfaceAdvanced.
	P2PHellosToUnicastMac() bool
	// SetP2PHellosToUnicastMac assigns bool provided by user to IsisInterfaceAdvanced
	SetP2PHellosToUnicastMac(value bool) IsisInterfaceAdvanced
	// HasP2PHellosToUnicastMac checks if P2PHellosToUnicastMac has been set in IsisInterfaceAdvanced
	HasP2PHellosToUnicastMac() bool
}

// If a padded Hello message is received on the interface, the length of
// the Hello packets sent out on that interface is adjusted to match.
// AutoAdjustMtu returns a bool
func (obj *isisInterfaceAdvanced) AutoAdjustMtu() bool {

	return *obj.obj.AutoAdjustMtu

}

// If a padded Hello message is received on the interface, the length of
// the Hello packets sent out on that interface is adjusted to match.
// AutoAdjustMtu returns a bool
func (obj *isisInterfaceAdvanced) HasAutoAdjustMtu() bool {
	return obj.obj.AutoAdjustMtu != nil
}

// If a padded Hello message is received on the interface, the length of
// the Hello packets sent out on that interface is adjusted to match.
// SetAutoAdjustMtu sets the bool value in the IsisInterfaceAdvanced object
func (obj *isisInterfaceAdvanced) SetAutoAdjustMtu(value bool) IsisInterfaceAdvanced {

	obj.obj.AutoAdjustMtu = &value
	return obj
}

// If a Level 1 Hello is received on this emulated router for an area
// not currently in its area list, an area from the received Hello is
// added to that list. This ensures an area match for all future
// Level 1 Hellos from the source L1 router.
// AutoAdjustArea returns a bool
func (obj *isisInterfaceAdvanced) AutoAdjustArea() bool {

	return *obj.obj.AutoAdjustArea

}

// If a Level 1 Hello is received on this emulated router for an area
// not currently in its area list, an area from the received Hello is
// added to that list. This ensures an area match for all future
// Level 1 Hellos from the source L1 router.
// AutoAdjustArea returns a bool
func (obj *isisInterfaceAdvanced) HasAutoAdjustArea() bool {
	return obj.obj.AutoAdjustArea != nil
}

// If a Level 1 Hello is received on this emulated router for an area
// not currently in its area list, an area from the received Hello is
// added to that list. This ensures an area match for all future
// Level 1 Hellos from the source L1 router.
// SetAutoAdjustArea sets the bool value in the IsisInterfaceAdvanced object
func (obj *isisInterfaceAdvanced) SetAutoAdjustArea(value bool) IsisInterfaceAdvanced {

	obj.obj.AutoAdjustArea = &value
	return obj
}

// If a Hello message listing supported protocols is received on this
// emulated router, the supported protocols advertised by this router
// are changed to match exactly.
// AutoAdjustSupportedProtocols returns a bool
func (obj *isisInterfaceAdvanced) AutoAdjustSupportedProtocols() bool {

	return *obj.obj.AutoAdjustSupportedProtocols

}

// If a Hello message listing supported protocols is received on this
// emulated router, the supported protocols advertised by this router
// are changed to match exactly.
// AutoAdjustSupportedProtocols returns a bool
func (obj *isisInterfaceAdvanced) HasAutoAdjustSupportedProtocols() bool {
	return obj.obj.AutoAdjustSupportedProtocols != nil
}

// If a Hello message listing supported protocols is received on this
// emulated router, the supported protocols advertised by this router
// are changed to match exactly.
// SetAutoAdjustSupportedProtocols sets the bool value in the IsisInterfaceAdvanced object
func (obj *isisInterfaceAdvanced) SetAutoAdjustSupportedProtocols(value bool) IsisInterfaceAdvanced {

	obj.obj.AutoAdjustSupportedProtocols = &value
	return obj
}

// If it is true, the Point-to-Point circuit will include 3-way TLV in its Point-to-Point IIH  and attempt to establish the adjacency as specified in RFC 5303. This field is not applicable if network_type is set to 'broadcast' type in ISIS interface.
// Enable3WayHandshake returns a bool
func (obj *isisInterfaceAdvanced) Enable3WayHandshake() bool {

	return *obj.obj.Enable_3WayHandshake

}

// If it is true, the Point-to-Point circuit will include 3-way TLV in its Point-to-Point IIH  and attempt to establish the adjacency as specified in RFC 5303. This field is not applicable if network_type is set to 'broadcast' type in ISIS interface.
// Enable3WayHandshake returns a bool
func (obj *isisInterfaceAdvanced) HasEnable3WayHandshake() bool {
	return obj.obj.Enable_3WayHandshake != nil
}

// If it is true, the Point-to-Point circuit will include 3-way TLV in its Point-to-Point IIH  and attempt to establish the adjacency as specified in RFC 5303. This field is not applicable if network_type is set to 'broadcast' type in ISIS interface.
// SetEnable3WayHandshake sets the bool value in the IsisInterfaceAdvanced object
func (obj *isisInterfaceAdvanced) SetEnable3WayHandshake(value bool) IsisInterfaceAdvanced {

	obj.obj.Enable_3WayHandshake = &value
	return obj
}

// If it is true, the Point-to-Point Hello messages will be sent to the unicast MAC address.
// P2PHellosToUnicastMac returns a bool
func (obj *isisInterfaceAdvanced) P2PHellosToUnicastMac() bool {

	return *obj.obj.P2PHellosToUnicastMac

}

// If it is true, the Point-to-Point Hello messages will be sent to the unicast MAC address.
// P2PHellosToUnicastMac returns a bool
func (obj *isisInterfaceAdvanced) HasP2PHellosToUnicastMac() bool {
	return obj.obj.P2PHellosToUnicastMac != nil
}

// If it is true, the Point-to-Point Hello messages will be sent to the unicast MAC address.
// SetP2PHellosToUnicastMac sets the bool value in the IsisInterfaceAdvanced object
func (obj *isisInterfaceAdvanced) SetP2PHellosToUnicastMac(value bool) IsisInterfaceAdvanced {

	obj.obj.P2PHellosToUnicastMac = &value
	return obj
}

func (obj *isisInterfaceAdvanced) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisInterfaceAdvanced) setDefault() {
	if obj.obj.AutoAdjustMtu == nil {
		obj.SetAutoAdjustMtu(true)
	}
	if obj.obj.AutoAdjustArea == nil {
		obj.SetAutoAdjustArea(true)
	}
	if obj.obj.AutoAdjustSupportedProtocols == nil {
		obj.SetAutoAdjustSupportedProtocols(false)
	}
	if obj.obj.Enable_3WayHandshake == nil {
		obj.SetEnable3WayHandshake(true)
	}
	if obj.obj.P2PHellosToUnicastMac == nil {
		obj.SetP2PHellosToUnicastMac(false)
	}

}
