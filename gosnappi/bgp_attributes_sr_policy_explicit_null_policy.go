package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesSrPolicyExplicitNullPolicy *****
type bgpAttributesSrPolicyExplicitNullPolicy struct {
	validation
	obj          *otg.BgpAttributesSrPolicyExplicitNullPolicy
	marshaller   marshalBgpAttributesSrPolicyExplicitNullPolicy
	unMarshaller unMarshalBgpAttributesSrPolicyExplicitNullPolicy
}

func NewBgpAttributesSrPolicyExplicitNullPolicy() BgpAttributesSrPolicyExplicitNullPolicy {
	obj := bgpAttributesSrPolicyExplicitNullPolicy{obj: &otg.BgpAttributesSrPolicyExplicitNullPolicy{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesSrPolicyExplicitNullPolicy) msg() *otg.BgpAttributesSrPolicyExplicitNullPolicy {
	return obj.obj
}

func (obj *bgpAttributesSrPolicyExplicitNullPolicy) setMsg(msg *otg.BgpAttributesSrPolicyExplicitNullPolicy) BgpAttributesSrPolicyExplicitNullPolicy {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesSrPolicyExplicitNullPolicy struct {
	obj *bgpAttributesSrPolicyExplicitNullPolicy
}

type marshalBgpAttributesSrPolicyExplicitNullPolicy interface {
	// ToProto marshals BgpAttributesSrPolicyExplicitNullPolicy to protobuf object *otg.BgpAttributesSrPolicyExplicitNullPolicy
	ToProto() (*otg.BgpAttributesSrPolicyExplicitNullPolicy, error)
	// ToPbText marshals BgpAttributesSrPolicyExplicitNullPolicy to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesSrPolicyExplicitNullPolicy to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesSrPolicyExplicitNullPolicy to JSON text
	ToJson() (string, error)
}

type unMarshalbgpAttributesSrPolicyExplicitNullPolicy struct {
	obj *bgpAttributesSrPolicyExplicitNullPolicy
}

type unMarshalBgpAttributesSrPolicyExplicitNullPolicy interface {
	// FromProto unmarshals BgpAttributesSrPolicyExplicitNullPolicy from protobuf object *otg.BgpAttributesSrPolicyExplicitNullPolicy
	FromProto(msg *otg.BgpAttributesSrPolicyExplicitNullPolicy) (BgpAttributesSrPolicyExplicitNullPolicy, error)
	// FromPbText unmarshals BgpAttributesSrPolicyExplicitNullPolicy from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesSrPolicyExplicitNullPolicy from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesSrPolicyExplicitNullPolicy from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesSrPolicyExplicitNullPolicy) Marshal() marshalBgpAttributesSrPolicyExplicitNullPolicy {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesSrPolicyExplicitNullPolicy{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesSrPolicyExplicitNullPolicy) Unmarshal() unMarshalBgpAttributesSrPolicyExplicitNullPolicy {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesSrPolicyExplicitNullPolicy{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesSrPolicyExplicitNullPolicy) ToProto() (*otg.BgpAttributesSrPolicyExplicitNullPolicy, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesSrPolicyExplicitNullPolicy) FromProto(msg *otg.BgpAttributesSrPolicyExplicitNullPolicy) (BgpAttributesSrPolicyExplicitNullPolicy, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesSrPolicyExplicitNullPolicy) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesSrPolicyExplicitNullPolicy) FromPbText(value string) error {
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

func (m *marshalbgpAttributesSrPolicyExplicitNullPolicy) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesSrPolicyExplicitNullPolicy) FromYaml(value string) error {
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

func (m *marshalbgpAttributesSrPolicyExplicitNullPolicy) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesSrPolicyExplicitNullPolicy) FromJson(value string) error {
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

func (obj *bgpAttributesSrPolicyExplicitNullPolicy) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesSrPolicyExplicitNullPolicy) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesSrPolicyExplicitNullPolicy) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesSrPolicyExplicitNullPolicy) Clone() (BgpAttributesSrPolicyExplicitNullPolicy, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesSrPolicyExplicitNullPolicy()
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

// BgpAttributesSrPolicyExplicitNullPolicy is this is an optional sub-tlv (Type 14) which indicates whether an Explicit NULL Label must be pushed on an unlabeled IP
// packet before other labels for IPv4 or IPv6 flows.
// - It is defined in Section 2.4.5 of draft-ietf-idr-sr-policy-safi-02.
type BgpAttributesSrPolicyExplicitNullPolicy interface {
	Validation
	// msg marshals BgpAttributesSrPolicyExplicitNullPolicy to protobuf object *otg.BgpAttributesSrPolicyExplicitNullPolicy
	// and doesn't set defaults
	msg() *otg.BgpAttributesSrPolicyExplicitNullPolicy
	// setMsg unmarshals BgpAttributesSrPolicyExplicitNullPolicy from protobuf object *otg.BgpAttributesSrPolicyExplicitNullPolicy
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesSrPolicyExplicitNullPolicy) BgpAttributesSrPolicyExplicitNullPolicy
	// provides marshal interface
	Marshal() marshalBgpAttributesSrPolicyExplicitNullPolicy
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesSrPolicyExplicitNullPolicy
	// validate validates BgpAttributesSrPolicyExplicitNullPolicy
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesSrPolicyExplicitNullPolicy, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns BgpAttributesSrPolicyExplicitNullPolicyChoiceEnum, set in BgpAttributesSrPolicyExplicitNullPolicy
	Choice() BgpAttributesSrPolicyExplicitNullPolicyChoiceEnum
	// setChoice assigns BgpAttributesSrPolicyExplicitNullPolicyChoiceEnum provided by user to BgpAttributesSrPolicyExplicitNullPolicy
	setChoice(value BgpAttributesSrPolicyExplicitNullPolicyChoiceEnum) BgpAttributesSrPolicyExplicitNullPolicy
	// HasChoice checks if Choice has been set in BgpAttributesSrPolicyExplicitNullPolicy
	HasChoice() bool
	// getter for PushIpv6 to set choice.
	PushIpv6()
	// getter for PushIpv4AndIpv6 to set choice.
	PushIpv4AndIpv6()
	// getter for PushIpv4 to set choice.
	PushIpv4()
	// getter for Unknown to set choice.
	Unknown()
	// getter for DonotPush to set choice.
	DonotPush()
}

type BgpAttributesSrPolicyExplicitNullPolicyChoiceEnum string

// Enum of Choice on BgpAttributesSrPolicyExplicitNullPolicy
var BgpAttributesSrPolicyExplicitNullPolicyChoice = struct {
	UNKNOWN            BgpAttributesSrPolicyExplicitNullPolicyChoiceEnum
	PUSH_IPV4          BgpAttributesSrPolicyExplicitNullPolicyChoiceEnum
	PUSH_IPV6          BgpAttributesSrPolicyExplicitNullPolicyChoiceEnum
	PUSH_IPV4_AND_IPV6 BgpAttributesSrPolicyExplicitNullPolicyChoiceEnum
	DONOT_PUSH         BgpAttributesSrPolicyExplicitNullPolicyChoiceEnum
}{
	UNKNOWN:            BgpAttributesSrPolicyExplicitNullPolicyChoiceEnum("unknown"),
	PUSH_IPV4:          BgpAttributesSrPolicyExplicitNullPolicyChoiceEnum("push_ipv4"),
	PUSH_IPV6:          BgpAttributesSrPolicyExplicitNullPolicyChoiceEnum("push_ipv6"),
	PUSH_IPV4_AND_IPV6: BgpAttributesSrPolicyExplicitNullPolicyChoiceEnum("push_ipv4_and_ipv6"),
	DONOT_PUSH:         BgpAttributesSrPolicyExplicitNullPolicyChoiceEnum("donot_push"),
}

func (obj *bgpAttributesSrPolicyExplicitNullPolicy) Choice() BgpAttributesSrPolicyExplicitNullPolicyChoiceEnum {
	return BgpAttributesSrPolicyExplicitNullPolicyChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for PushIpv6 to set choice
func (obj *bgpAttributesSrPolicyExplicitNullPolicy) PushIpv6() {
	obj.setChoice(BgpAttributesSrPolicyExplicitNullPolicyChoice.PUSH_IPV6)
}

// getter for PushIpv4AndIpv6 to set choice
func (obj *bgpAttributesSrPolicyExplicitNullPolicy) PushIpv4AndIpv6() {
	obj.setChoice(BgpAttributesSrPolicyExplicitNullPolicyChoice.PUSH_IPV4_AND_IPV6)
}

// getter for PushIpv4 to set choice
func (obj *bgpAttributesSrPolicyExplicitNullPolicy) PushIpv4() {
	obj.setChoice(BgpAttributesSrPolicyExplicitNullPolicyChoice.PUSH_IPV4)
}

// getter for Unknown to set choice
func (obj *bgpAttributesSrPolicyExplicitNullPolicy) Unknown() {
	obj.setChoice(BgpAttributesSrPolicyExplicitNullPolicyChoice.UNKNOWN)
}

// getter for DonotPush to set choice
func (obj *bgpAttributesSrPolicyExplicitNullPolicy) DonotPush() {
	obj.setChoice(BgpAttributesSrPolicyExplicitNullPolicyChoice.DONOT_PUSH)
}

// The Explicit NULL Label policy.
// Choice returns a string
func (obj *bgpAttributesSrPolicyExplicitNullPolicy) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *bgpAttributesSrPolicyExplicitNullPolicy) setChoice(value BgpAttributesSrPolicyExplicitNullPolicyChoiceEnum) BgpAttributesSrPolicyExplicitNullPolicy {
	intValue, ok := otg.BgpAttributesSrPolicyExplicitNullPolicy_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpAttributesSrPolicyExplicitNullPolicyChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpAttributesSrPolicyExplicitNullPolicy_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

func (obj *bgpAttributesSrPolicyExplicitNullPolicy) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *bgpAttributesSrPolicyExplicitNullPolicy) setDefault() {
	var choices_set int = 0
	var choice BgpAttributesSrPolicyExplicitNullPolicyChoiceEnum
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(BgpAttributesSrPolicyExplicitNullPolicyChoice.PUSH_IPV4_AND_IPV6)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in BgpAttributesSrPolicyExplicitNullPolicy")
			}
		} else {
			intVal := otg.BgpAttributesSrPolicyExplicitNullPolicy_Choice_Enum_value[string(choice)]
			enumValue := otg.BgpAttributesSrPolicyExplicitNullPolicy_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
