package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesCommunity *****
type bgpAttributesCommunity struct {
	validation
	obj                   *otg.BgpAttributesCommunity
	marshaller            marshalBgpAttributesCommunity
	unMarshaller          unMarshalBgpAttributesCommunity
	customCommunityHolder BgpAttributesCustomCommunity
}

func NewBgpAttributesCommunity() BgpAttributesCommunity {
	obj := bgpAttributesCommunity{obj: &otg.BgpAttributesCommunity{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesCommunity) msg() *otg.BgpAttributesCommunity {
	return obj.obj
}

func (obj *bgpAttributesCommunity) setMsg(msg *otg.BgpAttributesCommunity) BgpAttributesCommunity {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesCommunity struct {
	obj *bgpAttributesCommunity
}

type marshalBgpAttributesCommunity interface {
	// ToProto marshals BgpAttributesCommunity to protobuf object *otg.BgpAttributesCommunity
	ToProto() (*otg.BgpAttributesCommunity, error)
	// ToPbText marshals BgpAttributesCommunity to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesCommunity to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesCommunity to JSON text
	ToJson() (string, error)
}

type unMarshalbgpAttributesCommunity struct {
	obj *bgpAttributesCommunity
}

type unMarshalBgpAttributesCommunity interface {
	// FromProto unmarshals BgpAttributesCommunity from protobuf object *otg.BgpAttributesCommunity
	FromProto(msg *otg.BgpAttributesCommunity) (BgpAttributesCommunity, error)
	// FromPbText unmarshals BgpAttributesCommunity from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesCommunity from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesCommunity from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesCommunity) Marshal() marshalBgpAttributesCommunity {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesCommunity{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesCommunity) Unmarshal() unMarshalBgpAttributesCommunity {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesCommunity{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesCommunity) ToProto() (*otg.BgpAttributesCommunity, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesCommunity) FromProto(msg *otg.BgpAttributesCommunity) (BgpAttributesCommunity, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesCommunity) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesCommunity) FromPbText(value string) error {
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

func (m *marshalbgpAttributesCommunity) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesCommunity) FromYaml(value string) error {
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

func (m *marshalbgpAttributesCommunity) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesCommunity) FromJson(value string) error {
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

func (obj *bgpAttributesCommunity) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesCommunity) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesCommunity) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesCommunity) Clone() (BgpAttributesCommunity, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesCommunity()
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

func (obj *bgpAttributesCommunity) setNil() {
	obj.customCommunityHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpAttributesCommunity is the COMMUNITY attribute provide additional capability for tagging routes and for modifying BGP routing policy on
// upstream and downstream routers. BGP community is a 32-bit number which is broken into 16-bit AS number and a
// 16-bit custom value or it contains some pre-defined well known values.
type BgpAttributesCommunity interface {
	Validation
	// msg marshals BgpAttributesCommunity to protobuf object *otg.BgpAttributesCommunity
	// and doesn't set defaults
	msg() *otg.BgpAttributesCommunity
	// setMsg unmarshals BgpAttributesCommunity from protobuf object *otg.BgpAttributesCommunity
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesCommunity) BgpAttributesCommunity
	// provides marshal interface
	Marshal() marshalBgpAttributesCommunity
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesCommunity
	// validate validates BgpAttributesCommunity
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesCommunity, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns BgpAttributesCommunityChoiceEnum, set in BgpAttributesCommunity
	Choice() BgpAttributesCommunityChoiceEnum
	// setChoice assigns BgpAttributesCommunityChoiceEnum provided by user to BgpAttributesCommunity
	setChoice(value BgpAttributesCommunityChoiceEnum) BgpAttributesCommunity
	// getter for NoAdvertised to set choice.
	NoAdvertised()
	// getter for NoLlgr to set choice.
	NoLlgr()
	// getter for LlgrStale to set choice.
	LlgrStale()
	// getter for NoExportSubconfed to set choice.
	NoExportSubconfed()
	// getter for NoExport to set choice.
	NoExport()
	// CustomCommunity returns BgpAttributesCustomCommunity, set in BgpAttributesCommunity.
	// BgpAttributesCustomCommunity is user defined COMMUNITY attribute containing 2 byte AS and custom 2 byte value defined by the administrator of the domain.
	CustomCommunity() BgpAttributesCustomCommunity
	// SetCustomCommunity assigns BgpAttributesCustomCommunity provided by user to BgpAttributesCommunity.
	// BgpAttributesCustomCommunity is user defined COMMUNITY attribute containing 2 byte AS and custom 2 byte value defined by the administrator of the domain.
	SetCustomCommunity(value BgpAttributesCustomCommunity) BgpAttributesCommunity
	// HasCustomCommunity checks if CustomCommunity has been set in BgpAttributesCommunity
	HasCustomCommunity() bool
	setNil()
}

type BgpAttributesCommunityChoiceEnum string

// Enum of Choice on BgpAttributesCommunity
var BgpAttributesCommunityChoice = struct {
	CUSTOM_COMMUNITY    BgpAttributesCommunityChoiceEnum
	NO_EXPORT           BgpAttributesCommunityChoiceEnum
	NO_ADVERTISED       BgpAttributesCommunityChoiceEnum
	NO_EXPORT_SUBCONFED BgpAttributesCommunityChoiceEnum
	LLGR_STALE          BgpAttributesCommunityChoiceEnum
	NO_LLGR             BgpAttributesCommunityChoiceEnum
}{
	CUSTOM_COMMUNITY:    BgpAttributesCommunityChoiceEnum("custom_community"),
	NO_EXPORT:           BgpAttributesCommunityChoiceEnum("no_export"),
	NO_ADVERTISED:       BgpAttributesCommunityChoiceEnum("no_advertised"),
	NO_EXPORT_SUBCONFED: BgpAttributesCommunityChoiceEnum("no_export_subconfed"),
	LLGR_STALE:          BgpAttributesCommunityChoiceEnum("llgr_stale"),
	NO_LLGR:             BgpAttributesCommunityChoiceEnum("no_llgr"),
}

func (obj *bgpAttributesCommunity) Choice() BgpAttributesCommunityChoiceEnum {
	return BgpAttributesCommunityChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for NoAdvertised to set choice
func (obj *bgpAttributesCommunity) NoAdvertised() {
	obj.setChoice(BgpAttributesCommunityChoice.NO_ADVERTISED)
}

// getter for NoLlgr to set choice
func (obj *bgpAttributesCommunity) NoLlgr() {
	obj.setChoice(BgpAttributesCommunityChoice.NO_LLGR)
}

// getter for LlgrStale to set choice
func (obj *bgpAttributesCommunity) LlgrStale() {
	obj.setChoice(BgpAttributesCommunityChoice.LLGR_STALE)
}

// getter for NoExportSubconfed to set choice
func (obj *bgpAttributesCommunity) NoExportSubconfed() {
	obj.setChoice(BgpAttributesCommunityChoice.NO_EXPORT_SUBCONFED)
}

// getter for NoExport to set choice
func (obj *bgpAttributesCommunity) NoExport() {
	obj.setChoice(BgpAttributesCommunityChoice.NO_EXPORT)
}

func (obj *bgpAttributesCommunity) setChoice(value BgpAttributesCommunityChoiceEnum) BgpAttributesCommunity {
	intValue, ok := otg.BgpAttributesCommunity_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpAttributesCommunityChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpAttributesCommunity_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.CustomCommunity = nil
	obj.customCommunityHolder = nil

	if value == BgpAttributesCommunityChoice.CUSTOM_COMMUNITY {
		obj.obj.CustomCommunity = NewBgpAttributesCustomCommunity().msg()
	}

	return obj
}

// description is TBD
// CustomCommunity returns a BgpAttributesCustomCommunity
func (obj *bgpAttributesCommunity) CustomCommunity() BgpAttributesCustomCommunity {
	if obj.obj.CustomCommunity == nil {
		obj.setChoice(BgpAttributesCommunityChoice.CUSTOM_COMMUNITY)
	}
	if obj.customCommunityHolder == nil {
		obj.customCommunityHolder = &bgpAttributesCustomCommunity{obj: obj.obj.CustomCommunity}
	}
	return obj.customCommunityHolder
}

// description is TBD
// CustomCommunity returns a BgpAttributesCustomCommunity
func (obj *bgpAttributesCommunity) HasCustomCommunity() bool {
	return obj.obj.CustomCommunity != nil
}

// description is TBD
// SetCustomCommunity sets the BgpAttributesCustomCommunity value in the BgpAttributesCommunity object
func (obj *bgpAttributesCommunity) SetCustomCommunity(value BgpAttributesCustomCommunity) BgpAttributesCommunity {
	obj.setChoice(BgpAttributesCommunityChoice.CUSTOM_COMMUNITY)
	obj.customCommunityHolder = nil
	obj.obj.CustomCommunity = value.msg()

	return obj
}

func (obj *bgpAttributesCommunity) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface BgpAttributesCommunity")
	}

	if obj.obj.CustomCommunity != nil {

		obj.CustomCommunity().validateObj(vObj, set_default)
	}

}

func (obj *bgpAttributesCommunity) setDefault() {
	var choices_set int = 0
	var choice BgpAttributesCommunityChoiceEnum

	if obj.obj.CustomCommunity != nil {
		choices_set += 1
		choice = BgpAttributesCommunityChoice.CUSTOM_COMMUNITY
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in BgpAttributesCommunity")
			}
		} else {
			intVal := otg.BgpAttributesCommunity_Choice_Enum_value[string(choice)]
			enumValue := otg.BgpAttributesCommunity_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
