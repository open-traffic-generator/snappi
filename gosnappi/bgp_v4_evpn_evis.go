package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpV4EvpnEvis *****
type bgpV4EvpnEvis struct {
	validation
	obj            *otg.BgpV4EvpnEvis
	marshaller     marshalBgpV4EvpnEvis
	unMarshaller   unMarshalBgpV4EvpnEvis
	eviVxlanHolder BgpV4EviVxlan
}

func NewBgpV4EvpnEvis() BgpV4EvpnEvis {
	obj := bgpV4EvpnEvis{obj: &otg.BgpV4EvpnEvis{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpV4EvpnEvis) msg() *otg.BgpV4EvpnEvis {
	return obj.obj
}

func (obj *bgpV4EvpnEvis) setMsg(msg *otg.BgpV4EvpnEvis) BgpV4EvpnEvis {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpV4EvpnEvis struct {
	obj *bgpV4EvpnEvis
}

type marshalBgpV4EvpnEvis interface {
	// ToProto marshals BgpV4EvpnEvis to protobuf object *otg.BgpV4EvpnEvis
	ToProto() (*otg.BgpV4EvpnEvis, error)
	// ToPbText marshals BgpV4EvpnEvis to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpV4EvpnEvis to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpV4EvpnEvis to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpV4EvpnEvis to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpV4EvpnEvis struct {
	obj *bgpV4EvpnEvis
}

type unMarshalBgpV4EvpnEvis interface {
	// FromProto unmarshals BgpV4EvpnEvis from protobuf object *otg.BgpV4EvpnEvis
	FromProto(msg *otg.BgpV4EvpnEvis) (BgpV4EvpnEvis, error)
	// FromPbText unmarshals BgpV4EvpnEvis from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpV4EvpnEvis from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpV4EvpnEvis from JSON text
	FromJson(value string) error
}

func (obj *bgpV4EvpnEvis) Marshal() marshalBgpV4EvpnEvis {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpV4EvpnEvis{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpV4EvpnEvis) Unmarshal() unMarshalBgpV4EvpnEvis {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpV4EvpnEvis{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpV4EvpnEvis) ToProto() (*otg.BgpV4EvpnEvis, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpV4EvpnEvis) FromProto(msg *otg.BgpV4EvpnEvis) (BgpV4EvpnEvis, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpV4EvpnEvis) ToPbText() (string, error) {
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

func (m *unMarshalbgpV4EvpnEvis) FromPbText(value string) error {
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

func (m *marshalbgpV4EvpnEvis) ToYaml() (string, error) {
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

func (m *unMarshalbgpV4EvpnEvis) FromYaml(value string) error {
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

func (m *marshalbgpV4EvpnEvis) ToJsonRaw() (string, error) {
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

func (m *marshalbgpV4EvpnEvis) ToJson() (string, error) {
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

func (m *unMarshalbgpV4EvpnEvis) FromJson(value string) error {
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

func (obj *bgpV4EvpnEvis) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpV4EvpnEvis) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpV4EvpnEvis) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpV4EvpnEvis) Clone() (BgpV4EvpnEvis, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpV4EvpnEvis()
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

func (obj *bgpV4EvpnEvis) setNil() {
	obj.eviVxlanHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpV4EvpnEvis is this contains a list of different flavors of EVPN.
// For example EVPN over VXLAN or EVPN over MPLS etc to be configured per Ethernet segment.
// Need to instantiate correct type of EVPN instance as per requirement.
type BgpV4EvpnEvis interface {
	Validation
	// msg marshals BgpV4EvpnEvis to protobuf object *otg.BgpV4EvpnEvis
	// and doesn't set defaults
	msg() *otg.BgpV4EvpnEvis
	// setMsg unmarshals BgpV4EvpnEvis from protobuf object *otg.BgpV4EvpnEvis
	// and doesn't set defaults
	setMsg(*otg.BgpV4EvpnEvis) BgpV4EvpnEvis
	// provides marshal interface
	Marshal() marshalBgpV4EvpnEvis
	// provides unmarshal interface
	Unmarshal() unMarshalBgpV4EvpnEvis
	// validate validates BgpV4EvpnEvis
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpV4EvpnEvis, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns BgpV4EvpnEvisChoiceEnum, set in BgpV4EvpnEvis
	Choice() BgpV4EvpnEvisChoiceEnum
	// setChoice assigns BgpV4EvpnEvisChoiceEnum provided by user to BgpV4EvpnEvis
	setChoice(value BgpV4EvpnEvisChoiceEnum) BgpV4EvpnEvis
	// HasChoice checks if Choice has been set in BgpV4EvpnEvis
	HasChoice() bool
	// EviVxlan returns BgpV4EviVxlan, set in BgpV4EvpnEvis.
	// BgpV4EviVxlan is configuration for BGP EVPN EVI. Advertises following routes -
	//
	// # Type 3 - Inclusive Multicast Ethernet Tag Route
	//
	// Type 1 -  Ethernet Auto-discovery Route (Per EVI)
	//
	// Type 1 -  Ethernet Auto-discovery Route (Per ES)
	EviVxlan() BgpV4EviVxlan
	// SetEviVxlan assigns BgpV4EviVxlan provided by user to BgpV4EvpnEvis.
	// BgpV4EviVxlan is configuration for BGP EVPN EVI. Advertises following routes -
	//
	// # Type 3 - Inclusive Multicast Ethernet Tag Route
	//
	// Type 1 -  Ethernet Auto-discovery Route (Per EVI)
	//
	// Type 1 -  Ethernet Auto-discovery Route (Per ES)
	SetEviVxlan(value BgpV4EviVxlan) BgpV4EvpnEvis
	// HasEviVxlan checks if EviVxlan has been set in BgpV4EvpnEvis
	HasEviVxlan() bool
	setNil()
}

type BgpV4EvpnEvisChoiceEnum string

// Enum of Choice on BgpV4EvpnEvis
var BgpV4EvpnEvisChoice = struct {
	EVI_VXLAN BgpV4EvpnEvisChoiceEnum
}{
	EVI_VXLAN: BgpV4EvpnEvisChoiceEnum("evi_vxlan"),
}

func (obj *bgpV4EvpnEvis) Choice() BgpV4EvpnEvisChoiceEnum {
	return BgpV4EvpnEvisChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *bgpV4EvpnEvis) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *bgpV4EvpnEvis) setChoice(value BgpV4EvpnEvisChoiceEnum) BgpV4EvpnEvis {
	intValue, ok := otg.BgpV4EvpnEvis_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpV4EvpnEvisChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpV4EvpnEvis_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.EviVxlan = nil
	obj.eviVxlanHolder = nil

	if value == BgpV4EvpnEvisChoice.EVI_VXLAN {
		obj.obj.EviVxlan = NewBgpV4EviVxlan().msg()
	}

	return obj
}

// EVPN VXLAN instance to be configured per Ethernet Segment.
// EviVxlan returns a BgpV4EviVxlan
func (obj *bgpV4EvpnEvis) EviVxlan() BgpV4EviVxlan {
	if obj.obj.EviVxlan == nil {
		obj.setChoice(BgpV4EvpnEvisChoice.EVI_VXLAN)
	}
	if obj.eviVxlanHolder == nil {
		obj.eviVxlanHolder = &bgpV4EviVxlan{obj: obj.obj.EviVxlan}
	}
	return obj.eviVxlanHolder
}

// EVPN VXLAN instance to be configured per Ethernet Segment.
// EviVxlan returns a BgpV4EviVxlan
func (obj *bgpV4EvpnEvis) HasEviVxlan() bool {
	return obj.obj.EviVxlan != nil
}

// EVPN VXLAN instance to be configured per Ethernet Segment.
// SetEviVxlan sets the BgpV4EviVxlan value in the BgpV4EvpnEvis object
func (obj *bgpV4EvpnEvis) SetEviVxlan(value BgpV4EviVxlan) BgpV4EvpnEvis {
	obj.setChoice(BgpV4EvpnEvisChoice.EVI_VXLAN)
	obj.eviVxlanHolder = nil
	obj.obj.EviVxlan = value.msg()

	return obj
}

func (obj *bgpV4EvpnEvis) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.EviVxlan != nil {

		obj.EviVxlan().validateObj(vObj, set_default)
	}

}

func (obj *bgpV4EvpnEvis) setDefault() {
	var choices_set int = 0
	var choice BgpV4EvpnEvisChoiceEnum

	if obj.obj.EviVxlan != nil {
		choices_set += 1
		choice = BgpV4EvpnEvisChoice.EVI_VXLAN
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(BgpV4EvpnEvisChoice.EVI_VXLAN)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in BgpV4EvpnEvis")
			}
		} else {
			intVal := otg.BgpV4EvpnEvis_Choice_Enum_value[string(choice)]
			enumValue := otg.BgpV4EvpnEvis_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
