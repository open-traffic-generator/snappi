package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpL3VpnV6ServiceBinding *****
type bgpL3VpnV6ServiceBinding struct {
	validation
	obj              *otg.BgpL3VpnV6ServiceBinding
	marshaller       marshalBgpL3VpnV6ServiceBinding
	unMarshaller     unMarshalBgpL3VpnV6ServiceBinding
	mplsLabelsHolder BgpMplsLabelBindings
}

func NewBgpL3VpnV6ServiceBinding() BgpL3VpnV6ServiceBinding {
	obj := bgpL3VpnV6ServiceBinding{obj: &otg.BgpL3VpnV6ServiceBinding{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpL3VpnV6ServiceBinding) msg() *otg.BgpL3VpnV6ServiceBinding {
	return obj.obj
}

func (obj *bgpL3VpnV6ServiceBinding) setMsg(msg *otg.BgpL3VpnV6ServiceBinding) BgpL3VpnV6ServiceBinding {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpL3VpnV6ServiceBinding struct {
	obj *bgpL3VpnV6ServiceBinding
}

type marshalBgpL3VpnV6ServiceBinding interface {
	// ToProto marshals BgpL3VpnV6ServiceBinding to protobuf object *otg.BgpL3VpnV6ServiceBinding
	ToProto() (*otg.BgpL3VpnV6ServiceBinding, error)
	// ToPbText marshals BgpL3VpnV6ServiceBinding to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpL3VpnV6ServiceBinding to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpL3VpnV6ServiceBinding to JSON text
	ToJson() (string, error)
}

type unMarshalbgpL3VpnV6ServiceBinding struct {
	obj *bgpL3VpnV6ServiceBinding
}

type unMarshalBgpL3VpnV6ServiceBinding interface {
	// FromProto unmarshals BgpL3VpnV6ServiceBinding from protobuf object *otg.BgpL3VpnV6ServiceBinding
	FromProto(msg *otg.BgpL3VpnV6ServiceBinding) (BgpL3VpnV6ServiceBinding, error)
	// FromPbText unmarshals BgpL3VpnV6ServiceBinding from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpL3VpnV6ServiceBinding from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpL3VpnV6ServiceBinding from JSON text
	FromJson(value string) error
}

func (obj *bgpL3VpnV6ServiceBinding) Marshal() marshalBgpL3VpnV6ServiceBinding {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpL3VpnV6ServiceBinding{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpL3VpnV6ServiceBinding) Unmarshal() unMarshalBgpL3VpnV6ServiceBinding {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpL3VpnV6ServiceBinding{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpL3VpnV6ServiceBinding) ToProto() (*otg.BgpL3VpnV6ServiceBinding, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpL3VpnV6ServiceBinding) FromProto(msg *otg.BgpL3VpnV6ServiceBinding) (BgpL3VpnV6ServiceBinding, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpL3VpnV6ServiceBinding) ToPbText() (string, error) {
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

func (m *unMarshalbgpL3VpnV6ServiceBinding) FromPbText(value string) error {
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

func (m *marshalbgpL3VpnV6ServiceBinding) ToYaml() (string, error) {
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

func (m *unMarshalbgpL3VpnV6ServiceBinding) FromYaml(value string) error {
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

func (m *marshalbgpL3VpnV6ServiceBinding) ToJson() (string, error) {
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

func (m *unMarshalbgpL3VpnV6ServiceBinding) FromJson(value string) error {
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

func (obj *bgpL3VpnV6ServiceBinding) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpL3VpnV6ServiceBinding) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpL3VpnV6ServiceBinding) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpL3VpnV6ServiceBinding) Clone() (BgpL3VpnV6ServiceBinding, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpL3VpnV6ServiceBinding()
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

func (obj *bgpL3VpnV6ServiceBinding) setNil() {
	obj.mplsLabelsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpL3VpnV6ServiceBinding is selects how a VPN-IPv6 route range's dataplane binding is advertised. Currently the only defined choice is the traditional VPN MPLS label (RFC 4364 Section 3); the choice discriminator is kept independent of the shared Bgp.MplsLabelBindings schema so a future dataplane binding (for example an SRv6 Service SID, RFC 9252) can be added as a new sibling choice value without restructuring this schema or affecting existing mpls_labels configs.
type BgpL3VpnV6ServiceBinding interface {
	Validation
	// msg marshals BgpL3VpnV6ServiceBinding to protobuf object *otg.BgpL3VpnV6ServiceBinding
	// and doesn't set defaults
	msg() *otg.BgpL3VpnV6ServiceBinding
	// setMsg unmarshals BgpL3VpnV6ServiceBinding from protobuf object *otg.BgpL3VpnV6ServiceBinding
	// and doesn't set defaults
	setMsg(*otg.BgpL3VpnV6ServiceBinding) BgpL3VpnV6ServiceBinding
	// provides marshal interface
	Marshal() marshalBgpL3VpnV6ServiceBinding
	// provides unmarshal interface
	Unmarshal() unMarshalBgpL3VpnV6ServiceBinding
	// validate validates BgpL3VpnV6ServiceBinding
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpL3VpnV6ServiceBinding, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns BgpL3VpnV6ServiceBindingChoiceEnum, set in BgpL3VpnV6ServiceBinding
	Choice() BgpL3VpnV6ServiceBindingChoiceEnum
	// setChoice assigns BgpL3VpnV6ServiceBindingChoiceEnum provided by user to BgpL3VpnV6ServiceBinding
	setChoice(value BgpL3VpnV6ServiceBindingChoiceEnum) BgpL3VpnV6ServiceBinding
	// HasChoice checks if Choice has been set in BgpL3VpnV6ServiceBinding
	HasChoice() bool
	// MplsLabels returns BgpMplsLabelBindings, set in BgpL3VpnV6ServiceBinding.
	// BgpMplsLabelBindings is bGP may be used to advertise that a particular node (N) has bound a particular MPLS label, or a particular sequence of MPLS labels,
	// to a particular address prefix.
	// This is done by sending a Multiprotocol BGP UPDATE message with with an MP_REACH_NLRI attribute.
	// The Network Address of Next Hop field of that attribute contains an IP address of node N.
	// References: https://datatracker.ietf.org/doc/html/rfc3107
	// & https://datatracker.ietf.org/doc/html/rfc8277.
	MplsLabels() BgpMplsLabelBindings
	// SetMplsLabels assigns BgpMplsLabelBindings provided by user to BgpL3VpnV6ServiceBinding.
	// BgpMplsLabelBindings is bGP may be used to advertise that a particular node (N) has bound a particular MPLS label, or a particular sequence of MPLS labels,
	// to a particular address prefix.
	// This is done by sending a Multiprotocol BGP UPDATE message with with an MP_REACH_NLRI attribute.
	// The Network Address of Next Hop field of that attribute contains an IP address of node N.
	// References: https://datatracker.ietf.org/doc/html/rfc3107
	// & https://datatracker.ietf.org/doc/html/rfc8277.
	SetMplsLabels(value BgpMplsLabelBindings) BgpL3VpnV6ServiceBinding
	// HasMplsLabels checks if MplsLabels has been set in BgpL3VpnV6ServiceBinding
	HasMplsLabels() bool
	setNil()
}

type BgpL3VpnV6ServiceBindingChoiceEnum string

// Enum of Choice on BgpL3VpnV6ServiceBinding
var BgpL3VpnV6ServiceBindingChoice = struct {
	MPLS_LABELS BgpL3VpnV6ServiceBindingChoiceEnum
}{
	MPLS_LABELS: BgpL3VpnV6ServiceBindingChoiceEnum("mpls_labels"),
}

func (obj *bgpL3VpnV6ServiceBinding) Choice() BgpL3VpnV6ServiceBindingChoiceEnum {
	return BgpL3VpnV6ServiceBindingChoiceEnum(obj.obj.Choice.Enum().String())
}

// The VPN dataplane encoding advertised for this route range's routes.
// Choice returns a string
func (obj *bgpL3VpnV6ServiceBinding) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *bgpL3VpnV6ServiceBinding) setChoice(value BgpL3VpnV6ServiceBindingChoiceEnum) BgpL3VpnV6ServiceBinding {
	intValue, ok := otg.BgpL3VpnV6ServiceBinding_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpL3VpnV6ServiceBindingChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpL3VpnV6ServiceBinding_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.MplsLabels = nil
	obj.mplsLabelsHolder = nil

	if value == BgpL3VpnV6ServiceBindingChoice.MPLS_LABELS {
		obj.obj.MplsLabels = NewBgpMplsLabelBindings().msg()
	}

	return obj
}

// Optional configuration for a BGP speaker to bind an address prefix to one or more MPLS labels (RFC 3107/8277, RFC 4364 Section 3).
// MplsLabels returns a BgpMplsLabelBindings
func (obj *bgpL3VpnV6ServiceBinding) MplsLabels() BgpMplsLabelBindings {
	if obj.obj.MplsLabels == nil {
		obj.setChoice(BgpL3VpnV6ServiceBindingChoice.MPLS_LABELS)
	}
	if obj.mplsLabelsHolder == nil {
		obj.mplsLabelsHolder = &bgpMplsLabelBindings{obj: obj.obj.MplsLabels}
	}
	return obj.mplsLabelsHolder
}

// Optional configuration for a BGP speaker to bind an address prefix to one or more MPLS labels (RFC 3107/8277, RFC 4364 Section 3).
// MplsLabels returns a BgpMplsLabelBindings
func (obj *bgpL3VpnV6ServiceBinding) HasMplsLabels() bool {
	return obj.obj.MplsLabels != nil
}

// Optional configuration for a BGP speaker to bind an address prefix to one or more MPLS labels (RFC 3107/8277, RFC 4364 Section 3).
// SetMplsLabels sets the BgpMplsLabelBindings value in the BgpL3VpnV6ServiceBinding object
func (obj *bgpL3VpnV6ServiceBinding) SetMplsLabels(value BgpMplsLabelBindings) BgpL3VpnV6ServiceBinding {
	obj.setChoice(BgpL3VpnV6ServiceBindingChoice.MPLS_LABELS)
	obj.mplsLabelsHolder = nil
	obj.obj.MplsLabels = value.msg()

	return obj
}

func (obj *bgpL3VpnV6ServiceBinding) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MplsLabels != nil {

		obj.MplsLabels().validateObj(vObj, set_default)
	}

}

func (obj *bgpL3VpnV6ServiceBinding) setDefault() {
	var choices_set int = 0
	var choice BgpL3VpnV6ServiceBindingChoiceEnum

	if obj.obj.MplsLabels != nil {
		choices_set += 1
		choice = BgpL3VpnV6ServiceBindingChoice.MPLS_LABELS
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(BgpL3VpnV6ServiceBindingChoice.MPLS_LABELS)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in BgpL3VpnV6ServiceBinding")
			}
		} else {
			intVal := otg.BgpL3VpnV6ServiceBinding_Choice_Enum_value[string(choice)]
			enumValue := otg.BgpL3VpnV6ServiceBinding_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
