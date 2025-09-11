package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpMPlsLabelBindingsCapabilities *****
type bgpMPlsLabelBindingsCapabilities struct {
	validation
	obj                  *otg.BgpMPlsLabelBindingsCapabilities
	marshaller           marshalBgpMPlsLabelBindingsCapabilities
	unMarshaller         unMarshalBgpMPlsLabelBindingsCapabilities
	singleLabelHolder    BgpCapabilitiesMplsSingleLabel
	multipleLabelsHolder BgpCapabilitiesMplsMultipleLabels
}

func NewBgpMPlsLabelBindingsCapabilities() BgpMPlsLabelBindingsCapabilities {
	obj := bgpMPlsLabelBindingsCapabilities{obj: &otg.BgpMPlsLabelBindingsCapabilities{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpMPlsLabelBindingsCapabilities) msg() *otg.BgpMPlsLabelBindingsCapabilities {
	return obj.obj
}

func (obj *bgpMPlsLabelBindingsCapabilities) setMsg(msg *otg.BgpMPlsLabelBindingsCapabilities) BgpMPlsLabelBindingsCapabilities {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpMPlsLabelBindingsCapabilities struct {
	obj *bgpMPlsLabelBindingsCapabilities
}

type marshalBgpMPlsLabelBindingsCapabilities interface {
	// ToProto marshals BgpMPlsLabelBindingsCapabilities to protobuf object *otg.BgpMPlsLabelBindingsCapabilities
	ToProto() (*otg.BgpMPlsLabelBindingsCapabilities, error)
	// ToPbText marshals BgpMPlsLabelBindingsCapabilities to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpMPlsLabelBindingsCapabilities to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpMPlsLabelBindingsCapabilities to JSON text
	ToJson() (string, error)
}

type unMarshalbgpMPlsLabelBindingsCapabilities struct {
	obj *bgpMPlsLabelBindingsCapabilities
}

type unMarshalBgpMPlsLabelBindingsCapabilities interface {
	// FromProto unmarshals BgpMPlsLabelBindingsCapabilities from protobuf object *otg.BgpMPlsLabelBindingsCapabilities
	FromProto(msg *otg.BgpMPlsLabelBindingsCapabilities) (BgpMPlsLabelBindingsCapabilities, error)
	// FromPbText unmarshals BgpMPlsLabelBindingsCapabilities from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpMPlsLabelBindingsCapabilities from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpMPlsLabelBindingsCapabilities from JSON text
	FromJson(value string) error
}

func (obj *bgpMPlsLabelBindingsCapabilities) Marshal() marshalBgpMPlsLabelBindingsCapabilities {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpMPlsLabelBindingsCapabilities{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpMPlsLabelBindingsCapabilities) Unmarshal() unMarshalBgpMPlsLabelBindingsCapabilities {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpMPlsLabelBindingsCapabilities{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpMPlsLabelBindingsCapabilities) ToProto() (*otg.BgpMPlsLabelBindingsCapabilities, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpMPlsLabelBindingsCapabilities) FromProto(msg *otg.BgpMPlsLabelBindingsCapabilities) (BgpMPlsLabelBindingsCapabilities, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpMPlsLabelBindingsCapabilities) ToPbText() (string, error) {
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

func (m *unMarshalbgpMPlsLabelBindingsCapabilities) FromPbText(value string) error {
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

func (m *marshalbgpMPlsLabelBindingsCapabilities) ToYaml() (string, error) {
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

func (m *unMarshalbgpMPlsLabelBindingsCapabilities) FromYaml(value string) error {
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

func (m *marshalbgpMPlsLabelBindingsCapabilities) ToJson() (string, error) {
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

func (m *unMarshalbgpMPlsLabelBindingsCapabilities) FromJson(value string) error {
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

func (obj *bgpMPlsLabelBindingsCapabilities) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpMPlsLabelBindingsCapabilities) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpMPlsLabelBindingsCapabilities) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpMPlsLabelBindingsCapabilities) Clone() (BgpMPlsLabelBindingsCapabilities, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpMPlsLabelBindingsCapabilities()
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

func (obj *bgpMPlsLabelBindingsCapabilities) setNil() {
	obj.singleLabelHolder = nil
	obj.multipleLabelsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpMPlsLabelBindingsCapabilities is container for configuring capabilities for carrying Label Information in BGP Open Message for RFC3107 and RFC8277 for MPLS label bindings.
type BgpMPlsLabelBindingsCapabilities interface {
	Validation
	// msg marshals BgpMPlsLabelBindingsCapabilities to protobuf object *otg.BgpMPlsLabelBindingsCapabilities
	// and doesn't set defaults
	msg() *otg.BgpMPlsLabelBindingsCapabilities
	// setMsg unmarshals BgpMPlsLabelBindingsCapabilities from protobuf object *otg.BgpMPlsLabelBindingsCapabilities
	// and doesn't set defaults
	setMsg(*otg.BgpMPlsLabelBindingsCapabilities) BgpMPlsLabelBindingsCapabilities
	// provides marshal interface
	Marshal() marshalBgpMPlsLabelBindingsCapabilities
	// provides unmarshal interface
	Unmarshal() unMarshalBgpMPlsLabelBindingsCapabilities
	// validate validates BgpMPlsLabelBindingsCapabilities
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpMPlsLabelBindingsCapabilities, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns BgpMPlsLabelBindingsCapabilitiesChoiceEnum, set in BgpMPlsLabelBindingsCapabilities
	Choice() BgpMPlsLabelBindingsCapabilitiesChoiceEnum
	// setChoice assigns BgpMPlsLabelBindingsCapabilitiesChoiceEnum provided by user to BgpMPlsLabelBindingsCapabilities
	setChoice(value BgpMPlsLabelBindingsCapabilitiesChoiceEnum) BgpMPlsLabelBindingsCapabilities
	// HasChoice checks if Choice has been set in BgpMPlsLabelBindingsCapabilities
	HasChoice() bool
	// SingleLabel returns BgpCapabilitiesMplsSingleLabel, set in BgpMPlsLabelBindingsCapabilities.
	// BgpCapabilitiesMplsSingleLabel is container for configuring capabilities for carrying single Label Information. The Capabilities Optional Parameter is a triple that includes a one-octet Capability Code,  a one-octet Capability length, and a variable-length Capability Value.
	SingleLabel() BgpCapabilitiesMplsSingleLabel
	// SetSingleLabel assigns BgpCapabilitiesMplsSingleLabel provided by user to BgpMPlsLabelBindingsCapabilities.
	// BgpCapabilitiesMplsSingleLabel is container for configuring capabilities for carrying single Label Information. The Capabilities Optional Parameter is a triple that includes a one-octet Capability Code,  a one-octet Capability length, and a variable-length Capability Value.
	SetSingleLabel(value BgpCapabilitiesMplsSingleLabel) BgpMPlsLabelBindingsCapabilities
	// HasSingleLabel checks if SingleLabel has been set in BgpMPlsLabelBindingsCapabilities
	HasSingleLabel() bool
	// MultipleLabels returns BgpCapabilitiesMplsMultipleLabels, set in BgpMPlsLabelBindingsCapabilities.
	// BgpCapabilitiesMplsMultipleLabels is container for configuring capabilities for multiple carrying Labels Information as per RFC8277. The MPLS multiple Lable capability is advertised in Optional Capability under Multiple Label Capability (code 8). Reference: https://datatracker.ietf.org/doc/html/rfc8277#section-2.1.
	MultipleLabels() BgpCapabilitiesMplsMultipleLabels
	// SetMultipleLabels assigns BgpCapabilitiesMplsMultipleLabels provided by user to BgpMPlsLabelBindingsCapabilities.
	// BgpCapabilitiesMplsMultipleLabels is container for configuring capabilities for multiple carrying Labels Information as per RFC8277. The MPLS multiple Lable capability is advertised in Optional Capability under Multiple Label Capability (code 8). Reference: https://datatracker.ietf.org/doc/html/rfc8277#section-2.1.
	SetMultipleLabels(value BgpCapabilitiesMplsMultipleLabels) BgpMPlsLabelBindingsCapabilities
	// HasMultipleLabels checks if MultipleLabels has been set in BgpMPlsLabelBindingsCapabilities
	HasMultipleLabels() bool
	setNil()
}

type BgpMPlsLabelBindingsCapabilitiesChoiceEnum string

// Enum of Choice on BgpMPlsLabelBindingsCapabilities
var BgpMPlsLabelBindingsCapabilitiesChoice = struct {
	SINGLE_LABEL    BgpMPlsLabelBindingsCapabilitiesChoiceEnum
	MULTIPLE_LABELS BgpMPlsLabelBindingsCapabilitiesChoiceEnum
}{
	SINGLE_LABEL:    BgpMPlsLabelBindingsCapabilitiesChoiceEnum("single_label"),
	MULTIPLE_LABELS: BgpMPlsLabelBindingsCapabilitiesChoiceEnum("multiple_labels"),
}

func (obj *bgpMPlsLabelBindingsCapabilities) Choice() BgpMPlsLabelBindingsCapabilitiesChoiceEnum {
	return BgpMPlsLabelBindingsCapabilitiesChoiceEnum(obj.obj.Choice.Enum().String())
}

// BGP to Bind MPLS Labels to Address Prefixes as per RFC3107 or RFC8277. - single_label: A BGP speaker that uses Multiprotocol Extensions to carry label
// mapping information should use the Capabilities Optional Parameter,
// as defined in Capabilities Advertisement in RFC2842, to inform its peers about this capability.
// The MP_EXT Capability Code, as defined in Multiprotocol Extensions RFC2858, is used to
// advertise the (AFI, SAFI) pairs available on a particular connection.
// Reference: https://datatracker.ietf.org/doc/html/rfc8277#section-2.1 &
// https://datatracker.ietf.org/doc/html/rfc3107#section-5.
// - multiple_labels: As RFC8277 is known as Multiple Labels Capability - A BGP
// speaker can include a Capabilities Optional Parameter in a BGP OPEN
// message. The Capabilities Optional Parameter is a triple that
// includes a one-octet Capability Code, a one-octet Capability length,
// and a variable-length Capability Value.
// Reference: https://datatracker.ietf.org/doc/html/rfc8277#section-2.1.
// Choice returns a string
func (obj *bgpMPlsLabelBindingsCapabilities) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *bgpMPlsLabelBindingsCapabilities) setChoice(value BgpMPlsLabelBindingsCapabilitiesChoiceEnum) BgpMPlsLabelBindingsCapabilities {
	intValue, ok := otg.BgpMPlsLabelBindingsCapabilities_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpMPlsLabelBindingsCapabilitiesChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpMPlsLabelBindingsCapabilities_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.MultipleLabels = nil
	obj.multipleLabelsHolder = nil
	obj.obj.SingleLabel = nil
	obj.singleLabelHolder = nil

	if value == BgpMPlsLabelBindingsCapabilitiesChoice.SINGLE_LABEL {
		obj.obj.SingleLabel = NewBgpCapabilitiesMplsSingleLabel().msg()
	}

	if value == BgpMPlsLabelBindingsCapabilitiesChoice.MULTIPLE_LABELS {
		obj.obj.MultipleLabels = NewBgpCapabilitiesMplsMultipleLabels().msg()
	}

	return obj
}

// description is TBD
// SingleLabel returns a BgpCapabilitiesMplsSingleLabel
func (obj *bgpMPlsLabelBindingsCapabilities) SingleLabel() BgpCapabilitiesMplsSingleLabel {
	if obj.obj.SingleLabel == nil {
		obj.setChoice(BgpMPlsLabelBindingsCapabilitiesChoice.SINGLE_LABEL)
	}
	if obj.singleLabelHolder == nil {
		obj.singleLabelHolder = &bgpCapabilitiesMplsSingleLabel{obj: obj.obj.SingleLabel}
	}
	return obj.singleLabelHolder
}

// description is TBD
// SingleLabel returns a BgpCapabilitiesMplsSingleLabel
func (obj *bgpMPlsLabelBindingsCapabilities) HasSingleLabel() bool {
	return obj.obj.SingleLabel != nil
}

// description is TBD
// SetSingleLabel sets the BgpCapabilitiesMplsSingleLabel value in the BgpMPlsLabelBindingsCapabilities object
func (obj *bgpMPlsLabelBindingsCapabilities) SetSingleLabel(value BgpCapabilitiesMplsSingleLabel) BgpMPlsLabelBindingsCapabilities {
	obj.setChoice(BgpMPlsLabelBindingsCapabilitiesChoice.SINGLE_LABEL)
	obj.singleLabelHolder = nil
	obj.obj.SingleLabel = value.msg()

	return obj
}

// description is TBD
// MultipleLabels returns a BgpCapabilitiesMplsMultipleLabels
func (obj *bgpMPlsLabelBindingsCapabilities) MultipleLabels() BgpCapabilitiesMplsMultipleLabels {
	if obj.obj.MultipleLabels == nil {
		obj.setChoice(BgpMPlsLabelBindingsCapabilitiesChoice.MULTIPLE_LABELS)
	}
	if obj.multipleLabelsHolder == nil {
		obj.multipleLabelsHolder = &bgpCapabilitiesMplsMultipleLabels{obj: obj.obj.MultipleLabels}
	}
	return obj.multipleLabelsHolder
}

// description is TBD
// MultipleLabels returns a BgpCapabilitiesMplsMultipleLabels
func (obj *bgpMPlsLabelBindingsCapabilities) HasMultipleLabels() bool {
	return obj.obj.MultipleLabels != nil
}

// description is TBD
// SetMultipleLabels sets the BgpCapabilitiesMplsMultipleLabels value in the BgpMPlsLabelBindingsCapabilities object
func (obj *bgpMPlsLabelBindingsCapabilities) SetMultipleLabels(value BgpCapabilitiesMplsMultipleLabels) BgpMPlsLabelBindingsCapabilities {
	obj.setChoice(BgpMPlsLabelBindingsCapabilitiesChoice.MULTIPLE_LABELS)
	obj.multipleLabelsHolder = nil
	obj.obj.MultipleLabels = value.msg()

	return obj
}

func (obj *bgpMPlsLabelBindingsCapabilities) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SingleLabel != nil {

		obj.SingleLabel().validateObj(vObj, set_default)
	}

	if obj.obj.MultipleLabels != nil {

		obj.MultipleLabels().validateObj(vObj, set_default)
	}

}

func (obj *bgpMPlsLabelBindingsCapabilities) setDefault() {
	var choices_set int = 0
	var choice BgpMPlsLabelBindingsCapabilitiesChoiceEnum

	if obj.obj.SingleLabel != nil {
		choices_set += 1
		choice = BgpMPlsLabelBindingsCapabilitiesChoice.SINGLE_LABEL
	}

	if obj.obj.MultipleLabels != nil {
		choices_set += 1
		choice = BgpMPlsLabelBindingsCapabilitiesChoice.MULTIPLE_LABELS
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(BgpMPlsLabelBindingsCapabilitiesChoice.SINGLE_LABEL)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in BgpMPlsLabelBindingsCapabilities")
			}
		} else {
			intVal := otg.BgpMPlsLabelBindingsCapabilities_Choice_Enum_value[string(choice)]
			enumValue := otg.BgpMPlsLabelBindingsCapabilities_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
