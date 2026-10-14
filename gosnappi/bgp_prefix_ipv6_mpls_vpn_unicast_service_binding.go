package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpPrefixIpv6MplsVpnUnicastServiceBinding *****
type bgpPrefixIpv6MplsVpnUnicastServiceBinding struct {
	validation
	obj          *otg.BgpPrefixIpv6MplsVpnUnicastServiceBinding
	marshaller   marshalBgpPrefixIpv6MplsVpnUnicastServiceBinding
	unMarshaller unMarshalBgpPrefixIpv6MplsVpnUnicastServiceBinding
}

func NewBgpPrefixIpv6MplsVpnUnicastServiceBinding() BgpPrefixIpv6MplsVpnUnicastServiceBinding {
	obj := bgpPrefixIpv6MplsVpnUnicastServiceBinding{obj: &otg.BgpPrefixIpv6MplsVpnUnicastServiceBinding{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastServiceBinding) msg() *otg.BgpPrefixIpv6MplsVpnUnicastServiceBinding {
	return obj.obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastServiceBinding) setMsg(msg *otg.BgpPrefixIpv6MplsVpnUnicastServiceBinding) BgpPrefixIpv6MplsVpnUnicastServiceBinding {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpPrefixIpv6MplsVpnUnicastServiceBinding struct {
	obj *bgpPrefixIpv6MplsVpnUnicastServiceBinding
}

type marshalBgpPrefixIpv6MplsVpnUnicastServiceBinding interface {
	// ToProto marshals BgpPrefixIpv6MplsVpnUnicastServiceBinding to protobuf object *otg.BgpPrefixIpv6MplsVpnUnicastServiceBinding
	ToProto() (*otg.BgpPrefixIpv6MplsVpnUnicastServiceBinding, error)
	// ToPbText marshals BgpPrefixIpv6MplsVpnUnicastServiceBinding to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpPrefixIpv6MplsVpnUnicastServiceBinding to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpPrefixIpv6MplsVpnUnicastServiceBinding to JSON text
	ToJson() (string, error)
}

type unMarshalbgpPrefixIpv6MplsVpnUnicastServiceBinding struct {
	obj *bgpPrefixIpv6MplsVpnUnicastServiceBinding
}

type unMarshalBgpPrefixIpv6MplsVpnUnicastServiceBinding interface {
	// FromProto unmarshals BgpPrefixIpv6MplsVpnUnicastServiceBinding from protobuf object *otg.BgpPrefixIpv6MplsVpnUnicastServiceBinding
	FromProto(msg *otg.BgpPrefixIpv6MplsVpnUnicastServiceBinding) (BgpPrefixIpv6MplsVpnUnicastServiceBinding, error)
	// FromPbText unmarshals BgpPrefixIpv6MplsVpnUnicastServiceBinding from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpPrefixIpv6MplsVpnUnicastServiceBinding from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpPrefixIpv6MplsVpnUnicastServiceBinding from JSON text
	FromJson(value string) error
}

func (obj *bgpPrefixIpv6MplsVpnUnicastServiceBinding) Marshal() marshalBgpPrefixIpv6MplsVpnUnicastServiceBinding {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpPrefixIpv6MplsVpnUnicastServiceBinding{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpPrefixIpv6MplsVpnUnicastServiceBinding) Unmarshal() unMarshalBgpPrefixIpv6MplsVpnUnicastServiceBinding {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpPrefixIpv6MplsVpnUnicastServiceBinding{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpPrefixIpv6MplsVpnUnicastServiceBinding) ToProto() (*otg.BgpPrefixIpv6MplsVpnUnicastServiceBinding, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpPrefixIpv6MplsVpnUnicastServiceBinding) FromProto(msg *otg.BgpPrefixIpv6MplsVpnUnicastServiceBinding) (BgpPrefixIpv6MplsVpnUnicastServiceBinding, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpPrefixIpv6MplsVpnUnicastServiceBinding) ToPbText() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6MplsVpnUnicastServiceBinding) FromPbText(value string) error {
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

func (m *marshalbgpPrefixIpv6MplsVpnUnicastServiceBinding) ToYaml() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6MplsVpnUnicastServiceBinding) FromYaml(value string) error {
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

func (m *marshalbgpPrefixIpv6MplsVpnUnicastServiceBinding) ToJson() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6MplsVpnUnicastServiceBinding) FromJson(value string) error {
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

func (obj *bgpPrefixIpv6MplsVpnUnicastServiceBinding) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv6MplsVpnUnicastServiceBinding) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv6MplsVpnUnicastServiceBinding) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpPrefixIpv6MplsVpnUnicastServiceBinding) Clone() (BgpPrefixIpv6MplsVpnUnicastServiceBinding, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpPrefixIpv6MplsVpnUnicastServiceBinding()
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

// BgpPrefixIpv6MplsVpnUnicastServiceBinding is the VPN dataplane binding received for a VPN-IPv6 route, mirroring the Bgp.L3vpn.V6ServiceBinding choice on the configured route range. Currently the only defined choice is the VPN MPLS label (RFC 4364 Section 3, RFC 4659); kept as a choice, independent of the shared label field, so a future dataplane binding (for example an SRv6 Service SID, RFC 9252) can be reported as a new sibling choice value without restructuring this schema.
type BgpPrefixIpv6MplsVpnUnicastServiceBinding interface {
	Validation
	// msg marshals BgpPrefixIpv6MplsVpnUnicastServiceBinding to protobuf object *otg.BgpPrefixIpv6MplsVpnUnicastServiceBinding
	// and doesn't set defaults
	msg() *otg.BgpPrefixIpv6MplsVpnUnicastServiceBinding
	// setMsg unmarshals BgpPrefixIpv6MplsVpnUnicastServiceBinding from protobuf object *otg.BgpPrefixIpv6MplsVpnUnicastServiceBinding
	// and doesn't set defaults
	setMsg(*otg.BgpPrefixIpv6MplsVpnUnicastServiceBinding) BgpPrefixIpv6MplsVpnUnicastServiceBinding
	// provides marshal interface
	Marshal() marshalBgpPrefixIpv6MplsVpnUnicastServiceBinding
	// provides unmarshal interface
	Unmarshal() unMarshalBgpPrefixIpv6MplsVpnUnicastServiceBinding
	// validate validates BgpPrefixIpv6MplsVpnUnicastServiceBinding
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpPrefixIpv6MplsVpnUnicastServiceBinding, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns BgpPrefixIpv6MplsVpnUnicastServiceBindingChoiceEnum, set in BgpPrefixIpv6MplsVpnUnicastServiceBinding
	Choice() BgpPrefixIpv6MplsVpnUnicastServiceBindingChoiceEnum
	// setChoice assigns BgpPrefixIpv6MplsVpnUnicastServiceBindingChoiceEnum provided by user to BgpPrefixIpv6MplsVpnUnicastServiceBinding
	setChoice(value BgpPrefixIpv6MplsVpnUnicastServiceBindingChoiceEnum) BgpPrefixIpv6MplsVpnUnicastServiceBinding
	// HasChoice checks if Choice has been set in BgpPrefixIpv6MplsVpnUnicastServiceBinding
	HasChoice() bool
	// Label returns uint32, set in BgpPrefixIpv6MplsVpnUnicastServiceBinding.
	Label() uint32
	// SetLabel assigns uint32 provided by user to BgpPrefixIpv6MplsVpnUnicastServiceBinding
	SetLabel(value uint32) BgpPrefixIpv6MplsVpnUnicastServiceBinding
	// HasLabel checks if Label has been set in BgpPrefixIpv6MplsVpnUnicastServiceBinding
	HasLabel() bool
}

type BgpPrefixIpv6MplsVpnUnicastServiceBindingChoiceEnum string

// Enum of Choice on BgpPrefixIpv6MplsVpnUnicastServiceBinding
var BgpPrefixIpv6MplsVpnUnicastServiceBindingChoice = struct {
	LABEL BgpPrefixIpv6MplsVpnUnicastServiceBindingChoiceEnum
}{
	LABEL: BgpPrefixIpv6MplsVpnUnicastServiceBindingChoiceEnum("label"),
}

func (obj *bgpPrefixIpv6MplsVpnUnicastServiceBinding) Choice() BgpPrefixIpv6MplsVpnUnicastServiceBindingChoiceEnum {
	return BgpPrefixIpv6MplsVpnUnicastServiceBindingChoiceEnum(obj.obj.Choice.Enum().String())
}

// The VPN dataplane encoding received for this route.
// Choice returns a string
func (obj *bgpPrefixIpv6MplsVpnUnicastServiceBinding) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *bgpPrefixIpv6MplsVpnUnicastServiceBinding) setChoice(value BgpPrefixIpv6MplsVpnUnicastServiceBindingChoiceEnum) BgpPrefixIpv6MplsVpnUnicastServiceBinding {
	intValue, ok := otg.BgpPrefixIpv6MplsVpnUnicastServiceBinding_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpPrefixIpv6MplsVpnUnicastServiceBindingChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpPrefixIpv6MplsVpnUnicastServiceBinding_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Label = nil
	return obj
}

// The single VPN dataplane MPLS label (RFC 4364 Section 3, RFC 4659) bound to this VPN-IPv6 prefix. RFC 4364 assigns exactly one label per route, matching the configured mpls_labels range on Bgp.L3vpn.V6RouteRange.
// Label returns a uint32
func (obj *bgpPrefixIpv6MplsVpnUnicastServiceBinding) Label() uint32 {

	if obj.obj.Label == nil {
		obj.setChoice(BgpPrefixIpv6MplsVpnUnicastServiceBindingChoice.LABEL)
	}

	return *obj.obj.Label

}

// The single VPN dataplane MPLS label (RFC 4364 Section 3, RFC 4659) bound to this VPN-IPv6 prefix. RFC 4364 assigns exactly one label per route, matching the configured mpls_labels range on Bgp.L3vpn.V6RouteRange.
// Label returns a uint32
func (obj *bgpPrefixIpv6MplsVpnUnicastServiceBinding) HasLabel() bool {
	return obj.obj.Label != nil
}

// The single VPN dataplane MPLS label (RFC 4364 Section 3, RFC 4659) bound to this VPN-IPv6 prefix. RFC 4364 assigns exactly one label per route, matching the configured mpls_labels range on Bgp.L3vpn.V6RouteRange.
// SetLabel sets the uint32 value in the BgpPrefixIpv6MplsVpnUnicastServiceBinding object
func (obj *bgpPrefixIpv6MplsVpnUnicastServiceBinding) SetLabel(value uint32) BgpPrefixIpv6MplsVpnUnicastServiceBinding {
	obj.setChoice(BgpPrefixIpv6MplsVpnUnicastServiceBindingChoice.LABEL)
	obj.obj.Label = &value
	return obj
}

func (obj *bgpPrefixIpv6MplsVpnUnicastServiceBinding) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Label != nil {

		if *obj.obj.Label > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpPrefixIpv6MplsVpnUnicastServiceBinding.Label <= 1048575 but Got %d", *obj.obj.Label))
		}

	}

}

func (obj *bgpPrefixIpv6MplsVpnUnicastServiceBinding) setDefault() {
	var choices_set int = 0
	var choice BgpPrefixIpv6MplsVpnUnicastServiceBindingChoiceEnum

	if obj.obj.Label != nil {
		choices_set += 1
		choice = BgpPrefixIpv6MplsVpnUnicastServiceBindingChoice.LABEL
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(BgpPrefixIpv6MplsVpnUnicastServiceBindingChoice.LABEL)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in BgpPrefixIpv6MplsVpnUnicastServiceBinding")
			}
		} else {
			intVal := otg.BgpPrefixIpv6MplsVpnUnicastServiceBinding_Choice_Enum_value[string(choice)]
			enumValue := otg.BgpPrefixIpv6MplsVpnUnicastServiceBinding_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
