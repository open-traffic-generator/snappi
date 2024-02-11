package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpV6EvpnEvis *****
type bgpV6EvpnEvis struct {
	validation
	obj            *otg.BgpV6EvpnEvis
	marshaller     marshalBgpV6EvpnEvis
	unMarshaller   unMarshalBgpV6EvpnEvis
	eviVxlanHolder BgpV6EviVxlan
}

func NewBgpV6EvpnEvis() BgpV6EvpnEvis {
	obj := bgpV6EvpnEvis{obj: &otg.BgpV6EvpnEvis{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpV6EvpnEvis) msg() *otg.BgpV6EvpnEvis {
	return obj.obj
}

func (obj *bgpV6EvpnEvis) setMsg(msg *otg.BgpV6EvpnEvis) BgpV6EvpnEvis {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpV6EvpnEvis struct {
	obj *bgpV6EvpnEvis
}

type marshalBgpV6EvpnEvis interface {
	// ToProto marshals BgpV6EvpnEvis to protobuf object *otg.BgpV6EvpnEvis
	ToProto() (*otg.BgpV6EvpnEvis, error)
	// ToPbText marshals BgpV6EvpnEvis to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpV6EvpnEvis to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpV6EvpnEvis to JSON text
	ToJson() (string, error)
}

type unMarshalbgpV6EvpnEvis struct {
	obj *bgpV6EvpnEvis
}

type unMarshalBgpV6EvpnEvis interface {
	// FromProto unmarshals BgpV6EvpnEvis from protobuf object *otg.BgpV6EvpnEvis
	FromProto(msg *otg.BgpV6EvpnEvis) (BgpV6EvpnEvis, error)
	// FromPbText unmarshals BgpV6EvpnEvis from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpV6EvpnEvis from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpV6EvpnEvis from JSON text
	FromJson(value string) error
}

func (obj *bgpV6EvpnEvis) Marshal() marshalBgpV6EvpnEvis {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpV6EvpnEvis{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpV6EvpnEvis) Unmarshal() unMarshalBgpV6EvpnEvis {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpV6EvpnEvis{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpV6EvpnEvis) ToProto() (*otg.BgpV6EvpnEvis, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpV6EvpnEvis) FromProto(msg *otg.BgpV6EvpnEvis) (BgpV6EvpnEvis, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpV6EvpnEvis) ToPbText() (string, error) {
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

func (m *unMarshalbgpV6EvpnEvis) FromPbText(value string) error {
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

func (m *marshalbgpV6EvpnEvis) ToYaml() (string, error) {
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

func (m *unMarshalbgpV6EvpnEvis) FromYaml(value string) error {
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

func (m *marshalbgpV6EvpnEvis) ToJson() (string, error) {
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

func (m *unMarshalbgpV6EvpnEvis) FromJson(value string) error {
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

func (obj *bgpV6EvpnEvis) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpV6EvpnEvis) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpV6EvpnEvis) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpV6EvpnEvis) Clone() (BgpV6EvpnEvis, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpV6EvpnEvis()
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

func (obj *bgpV6EvpnEvis) setNil() {
	obj.eviVxlanHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpV6EvpnEvis is this contains a list of different flavors of EVPN.
// For example EVPN over VXLAN or EVPN over MPLS etc to be configured per Ethernet segment.
// Need to instantiate correct type of EVPN instance as per requirement.
type BgpV6EvpnEvis interface {
	Validation
	// msg marshals BgpV6EvpnEvis to protobuf object *otg.BgpV6EvpnEvis
	// and doesn't set defaults
	msg() *otg.BgpV6EvpnEvis
	// setMsg unmarshals BgpV6EvpnEvis from protobuf object *otg.BgpV6EvpnEvis
	// and doesn't set defaults
	setMsg(*otg.BgpV6EvpnEvis) BgpV6EvpnEvis
	// provides marshal interface
	Marshal() marshalBgpV6EvpnEvis
	// provides unmarshal interface
	Unmarshal() unMarshalBgpV6EvpnEvis
	// validate validates BgpV6EvpnEvis
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpV6EvpnEvis, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns BgpV6EvpnEvisChoiceEnum, set in BgpV6EvpnEvis
	Choice() BgpV6EvpnEvisChoiceEnum
	// setChoice assigns BgpV6EvpnEvisChoiceEnum provided by user to BgpV6EvpnEvis
	setChoice(value BgpV6EvpnEvisChoiceEnum) BgpV6EvpnEvis
	// HasChoice checks if Choice has been set in BgpV6EvpnEvis
	HasChoice() bool
	// EviVxlan returns BgpV6EviVxlan, set in BgpV6EvpnEvis.
	// BgpV6EviVxlan is configuration for BGP EVPN EVI. Advertises following routes -
	//
	// # Type 3 - Inclusive Multicast Ethernet Tag Route
	//
	// Type 1 -  Ethernet Auto-discovery Route (Per EVI)
	//
	// Type 1 -  Ethernet Auto-discovery Route (Per ES)
	EviVxlan() BgpV6EviVxlan
	// SetEviVxlan assigns BgpV6EviVxlan provided by user to BgpV6EvpnEvis.
	// BgpV6EviVxlan is configuration for BGP EVPN EVI. Advertises following routes -
	//
	// # Type 3 - Inclusive Multicast Ethernet Tag Route
	//
	// Type 1 -  Ethernet Auto-discovery Route (Per EVI)
	//
	// Type 1 -  Ethernet Auto-discovery Route (Per ES)
	SetEviVxlan(value BgpV6EviVxlan) BgpV6EvpnEvis
	// HasEviVxlan checks if EviVxlan has been set in BgpV6EvpnEvis
	HasEviVxlan() bool
	setNil()
}

type BgpV6EvpnEvisChoiceEnum string

// Enum of Choice on BgpV6EvpnEvis
var BgpV6EvpnEvisChoice = struct {
	EVI_VXLAN BgpV6EvpnEvisChoiceEnum
}{
	EVI_VXLAN: BgpV6EvpnEvisChoiceEnum("evi_vxlan"),
}

func (obj *bgpV6EvpnEvis) Choice() BgpV6EvpnEvisChoiceEnum {
	return BgpV6EvpnEvisChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *bgpV6EvpnEvis) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *bgpV6EvpnEvis) setChoice(value BgpV6EvpnEvisChoiceEnum) BgpV6EvpnEvis {
	intValue, ok := otg.BgpV6EvpnEvis_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpV6EvpnEvisChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpV6EvpnEvis_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.EviVxlan = nil
	obj.eviVxlanHolder = nil

	if value == BgpV6EvpnEvisChoice.EVI_VXLAN {
		obj.obj.EviVxlan = NewBgpV6EviVxlan().msg()
	}

	return obj
}

// EVPN VXLAN instance to be configured per Ethernet Segment.
// EviVxlan returns a BgpV6EviVxlan
func (obj *bgpV6EvpnEvis) EviVxlan() BgpV6EviVxlan {
	if obj.obj.EviVxlan == nil {
		obj.setChoice(BgpV6EvpnEvisChoice.EVI_VXLAN)
	}
	if obj.eviVxlanHolder == nil {
		obj.eviVxlanHolder = &bgpV6EviVxlan{obj: obj.obj.EviVxlan}
	}
	return obj.eviVxlanHolder
}

// EVPN VXLAN instance to be configured per Ethernet Segment.
// EviVxlan returns a BgpV6EviVxlan
func (obj *bgpV6EvpnEvis) HasEviVxlan() bool {
	return obj.obj.EviVxlan != nil
}

// EVPN VXLAN instance to be configured per Ethernet Segment.
// SetEviVxlan sets the BgpV6EviVxlan value in the BgpV6EvpnEvis object
func (obj *bgpV6EvpnEvis) SetEviVxlan(value BgpV6EviVxlan) BgpV6EvpnEvis {
	obj.setChoice(BgpV6EvpnEvisChoice.EVI_VXLAN)
	obj.eviVxlanHolder = nil
	obj.obj.EviVxlan = value.msg()

	return obj
}

func (obj *bgpV6EvpnEvis) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.EviVxlan != nil {

		obj.EviVxlan().validateObj(vObj, set_default)
	}

}

func (obj *bgpV6EvpnEvis) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(BgpV6EvpnEvisChoice.EVI_VXLAN)

	}

}
