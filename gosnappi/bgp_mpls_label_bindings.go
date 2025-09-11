package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpMplsLabelBindings *****
type bgpMplsLabelBindings struct {
	validation
	obj           *otg.BgpMplsLabelBindings
	marshaller    marshalBgpMplsLabelBindings
	unMarshaller  unMarshalBgpMplsLabelBindings
	labelHolder   BgpMplsLabelBindingsSingleLabelValue
	srLabelHolder BgpMplsLabelBindingsSingleLabelIndex
}

func NewBgpMplsLabelBindings() BgpMplsLabelBindings {
	obj := bgpMplsLabelBindings{obj: &otg.BgpMplsLabelBindings{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpMplsLabelBindings) msg() *otg.BgpMplsLabelBindings {
	return obj.obj
}

func (obj *bgpMplsLabelBindings) setMsg(msg *otg.BgpMplsLabelBindings) BgpMplsLabelBindings {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpMplsLabelBindings struct {
	obj *bgpMplsLabelBindings
}

type marshalBgpMplsLabelBindings interface {
	// ToProto marshals BgpMplsLabelBindings to protobuf object *otg.BgpMplsLabelBindings
	ToProto() (*otg.BgpMplsLabelBindings, error)
	// ToPbText marshals BgpMplsLabelBindings to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpMplsLabelBindings to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpMplsLabelBindings to JSON text
	ToJson() (string, error)
}

type unMarshalbgpMplsLabelBindings struct {
	obj *bgpMplsLabelBindings
}

type unMarshalBgpMplsLabelBindings interface {
	// FromProto unmarshals BgpMplsLabelBindings from protobuf object *otg.BgpMplsLabelBindings
	FromProto(msg *otg.BgpMplsLabelBindings) (BgpMplsLabelBindings, error)
	// FromPbText unmarshals BgpMplsLabelBindings from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpMplsLabelBindings from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpMplsLabelBindings from JSON text
	FromJson(value string) error
}

func (obj *bgpMplsLabelBindings) Marshal() marshalBgpMplsLabelBindings {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpMplsLabelBindings{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpMplsLabelBindings) Unmarshal() unMarshalBgpMplsLabelBindings {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpMplsLabelBindings{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpMplsLabelBindings) ToProto() (*otg.BgpMplsLabelBindings, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpMplsLabelBindings) FromProto(msg *otg.BgpMplsLabelBindings) (BgpMplsLabelBindings, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpMplsLabelBindings) ToPbText() (string, error) {
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

func (m *unMarshalbgpMplsLabelBindings) FromPbText(value string) error {
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

func (m *marshalbgpMplsLabelBindings) ToYaml() (string, error) {
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

func (m *unMarshalbgpMplsLabelBindings) FromYaml(value string) error {
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

func (m *marshalbgpMplsLabelBindings) ToJson() (string, error) {
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

func (m *unMarshalbgpMplsLabelBindings) FromJson(value string) error {
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

func (obj *bgpMplsLabelBindings) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpMplsLabelBindings) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpMplsLabelBindings) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpMplsLabelBindings) Clone() (BgpMplsLabelBindings, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpMplsLabelBindings()
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

func (obj *bgpMplsLabelBindings) setNil() {
	obj.labelHolder = nil
	obj.srLabelHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpMplsLabelBindings is bGP may be used to advertise that a particular node (N) has bound a particular MPLS label, or a particular sequence of MPLS labels,
// to a particular address prefix.
// This is done by sending a Multiprotocol BGP UPDATE message with with an MP_REACH_NLRI attribute.
// The Network Address of Next Hop field of that attribute contains an IP address of node N.
// References: https://datatracker.ietf.org/doc/html/rfc3107
// & https://datatracker.ietf.org/doc/html/rfc8277.
type BgpMplsLabelBindings interface {
	Validation
	// msg marshals BgpMplsLabelBindings to protobuf object *otg.BgpMplsLabelBindings
	// and doesn't set defaults
	msg() *otg.BgpMplsLabelBindings
	// setMsg unmarshals BgpMplsLabelBindings from protobuf object *otg.BgpMplsLabelBindings
	// and doesn't set defaults
	setMsg(*otg.BgpMplsLabelBindings) BgpMplsLabelBindings
	// provides marshal interface
	Marshal() marshalBgpMplsLabelBindings
	// provides unmarshal interface
	Unmarshal() unMarshalBgpMplsLabelBindings
	// validate validates BgpMplsLabelBindings
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpMplsLabelBindings, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns BgpMplsLabelBindingsChoiceEnum, set in BgpMplsLabelBindings
	Choice() BgpMplsLabelBindingsChoiceEnum
	// setChoice assigns BgpMplsLabelBindingsChoiceEnum provided by user to BgpMplsLabelBindings
	setChoice(value BgpMplsLabelBindingsChoiceEnum) BgpMplsLabelBindings
	// HasChoice checks if Choice has been set in BgpMplsLabelBindings
	HasChoice() bool
	// Label returns BgpMplsLabelBindingsSingleLabelValue, set in BgpMplsLabelBindings.
	// BgpMplsLabelBindingsSingleLabelValue is container for Single MPLS Value.
	Label() BgpMplsLabelBindingsSingleLabelValue
	// SetLabel assigns BgpMplsLabelBindingsSingleLabelValue provided by user to BgpMplsLabelBindings.
	// BgpMplsLabelBindingsSingleLabelValue is container for Single MPLS Value.
	SetLabel(value BgpMplsLabelBindingsSingleLabelValue) BgpMplsLabelBindings
	// HasLabel checks if Label has been set in BgpMplsLabelBindings
	HasLabel() bool
	// SrLabel returns BgpMplsLabelBindingsSingleLabelIndex, set in BgpMplsLabelBindings.
	// BgpMplsLabelBindingsSingleLabelIndex is container for Single MPLS SR Index.
	SrLabel() BgpMplsLabelBindingsSingleLabelIndex
	// SetSrLabel assigns BgpMplsLabelBindingsSingleLabelIndex provided by user to BgpMplsLabelBindings.
	// BgpMplsLabelBindingsSingleLabelIndex is container for Single MPLS SR Index.
	SetSrLabel(value BgpMplsLabelBindingsSingleLabelIndex) BgpMplsLabelBindings
	// HasSrLabel checks if SrLabel has been set in BgpMplsLabelBindings
	HasSrLabel() bool
	setNil()
}

type BgpMplsLabelBindingsChoiceEnum string

// Enum of Choice on BgpMplsLabelBindings
var BgpMplsLabelBindingsChoice = struct {
	LABEL    BgpMplsLabelBindingsChoiceEnum
	SR_LABEL BgpMplsLabelBindingsChoiceEnum
}{
	LABEL:    BgpMplsLabelBindingsChoiceEnum("label"),
	SR_LABEL: BgpMplsLabelBindingsChoiceEnum("sr_label"),
}

func (obj *bgpMplsLabelBindings) Choice() BgpMplsLabelBindingsChoiceEnum {
	return BgpMplsLabelBindingsChoiceEnum(obj.obj.Choice.Enum().String())
}

// Bind one or more MPLS Labels to Address Prefixes.
// - label: Label Value is to be configure here.
// - sr_label: The mpls_lebel_capability under BGP capability is chosen as single_label. At least one SRGB range has been configured in mpls_srgb_ranges under device/bgp.
// Choice returns a string
func (obj *bgpMplsLabelBindings) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *bgpMplsLabelBindings) setChoice(value BgpMplsLabelBindingsChoiceEnum) BgpMplsLabelBindings {
	intValue, ok := otg.BgpMplsLabelBindings_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpMplsLabelBindingsChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpMplsLabelBindings_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.SrLabel = nil
	obj.srLabelHolder = nil
	obj.obj.Label = nil
	obj.labelHolder = nil

	if value == BgpMplsLabelBindingsChoice.LABEL {
		obj.obj.Label = NewBgpMplsLabelBindingsSingleLabelValue().msg()
	}

	if value == BgpMplsLabelBindingsChoice.SR_LABEL {
		obj.obj.SrLabel = NewBgpMplsLabelBindingsSingleLabelIndex().msg()
	}

	return obj
}

// Group of MPLS Label value bind to a prefix
// Label returns a BgpMplsLabelBindingsSingleLabelValue
func (obj *bgpMplsLabelBindings) Label() BgpMplsLabelBindingsSingleLabelValue {
	if obj.obj.Label == nil {
		obj.setChoice(BgpMplsLabelBindingsChoice.LABEL)
	}
	if obj.labelHolder == nil {
		obj.labelHolder = &bgpMplsLabelBindingsSingleLabelValue{obj: obj.obj.Label}
	}
	return obj.labelHolder
}

// Group of MPLS Label value bind to a prefix
// Label returns a BgpMplsLabelBindingsSingleLabelValue
func (obj *bgpMplsLabelBindings) HasLabel() bool {
	return obj.obj.Label != nil
}

// Group of MPLS Label value bind to a prefix
// SetLabel sets the BgpMplsLabelBindingsSingleLabelValue value in the BgpMplsLabelBindings object
func (obj *bgpMplsLabelBindings) SetLabel(value BgpMplsLabelBindingsSingleLabelValue) BgpMplsLabelBindings {
	obj.setChoice(BgpMplsLabelBindingsChoice.LABEL)
	obj.labelHolder = nil
	obj.obj.Label = value.msg()

	return obj
}

// Group of MPLS SR Label Index bind to a prefix
// SrLabel returns a BgpMplsLabelBindingsSingleLabelIndex
func (obj *bgpMplsLabelBindings) SrLabel() BgpMplsLabelBindingsSingleLabelIndex {
	if obj.obj.SrLabel == nil {
		obj.setChoice(BgpMplsLabelBindingsChoice.SR_LABEL)
	}
	if obj.srLabelHolder == nil {
		obj.srLabelHolder = &bgpMplsLabelBindingsSingleLabelIndex{obj: obj.obj.SrLabel}
	}
	return obj.srLabelHolder
}

// Group of MPLS SR Label Index bind to a prefix
// SrLabel returns a BgpMplsLabelBindingsSingleLabelIndex
func (obj *bgpMplsLabelBindings) HasSrLabel() bool {
	return obj.obj.SrLabel != nil
}

// Group of MPLS SR Label Index bind to a prefix
// SetSrLabel sets the BgpMplsLabelBindingsSingleLabelIndex value in the BgpMplsLabelBindings object
func (obj *bgpMplsLabelBindings) SetSrLabel(value BgpMplsLabelBindingsSingleLabelIndex) BgpMplsLabelBindings {
	obj.setChoice(BgpMplsLabelBindingsChoice.SR_LABEL)
	obj.srLabelHolder = nil
	obj.obj.SrLabel = value.msg()

	return obj
}

func (obj *bgpMplsLabelBindings) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Label != nil {

		obj.Label().validateObj(vObj, set_default)
	}

	if obj.obj.SrLabel != nil {

		obj.SrLabel().validateObj(vObj, set_default)
	}

}

func (obj *bgpMplsLabelBindings) setDefault() {
	var choices_set int = 0
	var choice BgpMplsLabelBindingsChoiceEnum

	if obj.obj.Label != nil {
		choices_set += 1
		choice = BgpMplsLabelBindingsChoice.LABEL
	}

	if obj.obj.SrLabel != nil {
		choices_set += 1
		choice = BgpMplsLabelBindingsChoice.SR_LABEL
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(BgpMplsLabelBindingsChoice.LABEL)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in BgpMplsLabelBindings")
			}
		} else {
			intVal := otg.BgpMplsLabelBindings_Choice_Enum_value[string(choice)]
			enumValue := otg.BgpMplsLabelBindings_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
