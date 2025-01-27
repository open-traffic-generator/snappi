package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesTunnelEncapsulation *****
type bgpAttributesTunnelEncapsulation struct {
	validation
	obj            *otg.BgpAttributesTunnelEncapsulation
	marshaller     marshalBgpAttributesTunnelEncapsulation
	unMarshaller   unMarshalBgpAttributesTunnelEncapsulation
	srPolicyHolder BgpAttributesSegmentRoutingPolicy
}

func NewBgpAttributesTunnelEncapsulation() BgpAttributesTunnelEncapsulation {
	obj := bgpAttributesTunnelEncapsulation{obj: &otg.BgpAttributesTunnelEncapsulation{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesTunnelEncapsulation) msg() *otg.BgpAttributesTunnelEncapsulation {
	return obj.obj
}

func (obj *bgpAttributesTunnelEncapsulation) setMsg(msg *otg.BgpAttributesTunnelEncapsulation) BgpAttributesTunnelEncapsulation {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesTunnelEncapsulation struct {
	obj *bgpAttributesTunnelEncapsulation
}

type marshalBgpAttributesTunnelEncapsulation interface {
	// ToProto marshals BgpAttributesTunnelEncapsulation to protobuf object *otg.BgpAttributesTunnelEncapsulation
	ToProto() (*otg.BgpAttributesTunnelEncapsulation, error)
	// ToPbText marshals BgpAttributesTunnelEncapsulation to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesTunnelEncapsulation to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesTunnelEncapsulation to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpAttributesTunnelEncapsulation to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpAttributesTunnelEncapsulation struct {
	obj *bgpAttributesTunnelEncapsulation
}

type unMarshalBgpAttributesTunnelEncapsulation interface {
	// FromProto unmarshals BgpAttributesTunnelEncapsulation from protobuf object *otg.BgpAttributesTunnelEncapsulation
	FromProto(msg *otg.BgpAttributesTunnelEncapsulation) (BgpAttributesTunnelEncapsulation, error)
	// FromPbText unmarshals BgpAttributesTunnelEncapsulation from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesTunnelEncapsulation from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesTunnelEncapsulation from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesTunnelEncapsulation) Marshal() marshalBgpAttributesTunnelEncapsulation {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesTunnelEncapsulation{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesTunnelEncapsulation) Unmarshal() unMarshalBgpAttributesTunnelEncapsulation {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesTunnelEncapsulation{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesTunnelEncapsulation) ToProto() (*otg.BgpAttributesTunnelEncapsulation, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesTunnelEncapsulation) FromProto(msg *otg.BgpAttributesTunnelEncapsulation) (BgpAttributesTunnelEncapsulation, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesTunnelEncapsulation) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesTunnelEncapsulation) FromPbText(value string) error {
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

func (m *marshalbgpAttributesTunnelEncapsulation) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesTunnelEncapsulation) FromYaml(value string) error {
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

func (m *marshalbgpAttributesTunnelEncapsulation) ToJsonRaw() (string, error) {
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

func (m *marshalbgpAttributesTunnelEncapsulation) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesTunnelEncapsulation) FromJson(value string) error {
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

func (obj *bgpAttributesTunnelEncapsulation) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesTunnelEncapsulation) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesTunnelEncapsulation) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesTunnelEncapsulation) Clone() (BgpAttributesTunnelEncapsulation, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesTunnelEncapsulation()
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

func (obj *bgpAttributesTunnelEncapsulation) setNil() {
	obj.srPolicyHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpAttributesTunnelEncapsulation is the TUNNEL_ENCAPSULATION  attribute is used by a BGP speaker to inform other BGP speakers how to encapsulate packets that need to be sent to it.
// It is defined in RFC9012 and is assigned a Type code of 23.
type BgpAttributesTunnelEncapsulation interface {
	Validation
	// msg marshals BgpAttributesTunnelEncapsulation to protobuf object *otg.BgpAttributesTunnelEncapsulation
	// and doesn't set defaults
	msg() *otg.BgpAttributesTunnelEncapsulation
	// setMsg unmarshals BgpAttributesTunnelEncapsulation from protobuf object *otg.BgpAttributesTunnelEncapsulation
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesTunnelEncapsulation) BgpAttributesTunnelEncapsulation
	// provides marshal interface
	Marshal() marshalBgpAttributesTunnelEncapsulation
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesTunnelEncapsulation
	// validate validates BgpAttributesTunnelEncapsulation
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesTunnelEncapsulation, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns BgpAttributesTunnelEncapsulationChoiceEnum, set in BgpAttributesTunnelEncapsulation
	Choice() BgpAttributesTunnelEncapsulationChoiceEnum
	// setChoice assigns BgpAttributesTunnelEncapsulationChoiceEnum provided by user to BgpAttributesTunnelEncapsulation
	setChoice(value BgpAttributesTunnelEncapsulationChoiceEnum) BgpAttributesTunnelEncapsulation
	// HasChoice checks if Choice has been set in BgpAttributesTunnelEncapsulation
	HasChoice() bool
	// SrPolicy returns BgpAttributesSegmentRoutingPolicy, set in BgpAttributesTunnelEncapsulation.
	// BgpAttributesSegmentRoutingPolicy is optional Segment Routing Policy information as defined in draft-ietf-idr-sr-policy-safi-02.
	// This information is carried in TUNNEL_ENCAPSULATION attribute with type set to  SR Policy (15).
	SrPolicy() BgpAttributesSegmentRoutingPolicy
	// SetSrPolicy assigns BgpAttributesSegmentRoutingPolicy provided by user to BgpAttributesTunnelEncapsulation.
	// BgpAttributesSegmentRoutingPolicy is optional Segment Routing Policy information as defined in draft-ietf-idr-sr-policy-safi-02.
	// This information is carried in TUNNEL_ENCAPSULATION attribute with type set to  SR Policy (15).
	SetSrPolicy(value BgpAttributesSegmentRoutingPolicy) BgpAttributesTunnelEncapsulation
	// HasSrPolicy checks if SrPolicy has been set in BgpAttributesTunnelEncapsulation
	HasSrPolicy() bool
	setNil()
}

type BgpAttributesTunnelEncapsulationChoiceEnum string

// Enum of Choice on BgpAttributesTunnelEncapsulation
var BgpAttributesTunnelEncapsulationChoice = struct {
	SR_POLICY BgpAttributesTunnelEncapsulationChoiceEnum
}{
	SR_POLICY: BgpAttributesTunnelEncapsulationChoiceEnum("sr_policy"),
}

func (obj *bgpAttributesTunnelEncapsulation) Choice() BgpAttributesTunnelEncapsulationChoiceEnum {
	return BgpAttributesTunnelEncapsulationChoiceEnum(obj.obj.Choice.Enum().String())
}

// Identifies a type of tunnel. The field contains values from the IANA registry "BGP Tunnel Encapsulation Attribute Tunnel Types".
// Choice returns a string
func (obj *bgpAttributesTunnelEncapsulation) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *bgpAttributesTunnelEncapsulation) setChoice(value BgpAttributesTunnelEncapsulationChoiceEnum) BgpAttributesTunnelEncapsulation {
	intValue, ok := otg.BgpAttributesTunnelEncapsulation_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpAttributesTunnelEncapsulationChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpAttributesTunnelEncapsulation_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.SrPolicy = nil
	obj.srPolicyHolder = nil

	if value == BgpAttributesTunnelEncapsulationChoice.SR_POLICY {
		obj.obj.SrPolicy = NewBgpAttributesSegmentRoutingPolicy().msg()
	}

	return obj
}

// description is TBD
// SrPolicy returns a BgpAttributesSegmentRoutingPolicy
func (obj *bgpAttributesTunnelEncapsulation) SrPolicy() BgpAttributesSegmentRoutingPolicy {
	if obj.obj.SrPolicy == nil {
		obj.setChoice(BgpAttributesTunnelEncapsulationChoice.SR_POLICY)
	}
	if obj.srPolicyHolder == nil {
		obj.srPolicyHolder = &bgpAttributesSegmentRoutingPolicy{obj: obj.obj.SrPolicy}
	}
	return obj.srPolicyHolder
}

// description is TBD
// SrPolicy returns a BgpAttributesSegmentRoutingPolicy
func (obj *bgpAttributesTunnelEncapsulation) HasSrPolicy() bool {
	return obj.obj.SrPolicy != nil
}

// description is TBD
// SetSrPolicy sets the BgpAttributesSegmentRoutingPolicy value in the BgpAttributesTunnelEncapsulation object
func (obj *bgpAttributesTunnelEncapsulation) SetSrPolicy(value BgpAttributesSegmentRoutingPolicy) BgpAttributesTunnelEncapsulation {
	obj.setChoice(BgpAttributesTunnelEncapsulationChoice.SR_POLICY)
	obj.srPolicyHolder = nil
	obj.obj.SrPolicy = value.msg()

	return obj
}

func (obj *bgpAttributesTunnelEncapsulation) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SrPolicy != nil {

		obj.SrPolicy().validateObj(vObj, set_default)
	}

}

func (obj *bgpAttributesTunnelEncapsulation) setDefault() {
	var choices_set int = 0
	var choice BgpAttributesTunnelEncapsulationChoiceEnum

	if obj.obj.SrPolicy != nil {
		choices_set += 1
		choice = BgpAttributesTunnelEncapsulationChoice.SR_POLICY
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(BgpAttributesTunnelEncapsulationChoice.SR_POLICY)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in BgpAttributesTunnelEncapsulation")
			}
		} else {
			intVal := otg.BgpAttributesTunnelEncapsulation_Choice_Enum_value[string(choice)]
			enumValue := otg.BgpAttributesTunnelEncapsulation_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
