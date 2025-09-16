package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** CaptureFilter *****
type captureFilter struct {
	validation
	obj            *otg.CaptureFilter
	marshaller     marshalCaptureFilter
	unMarshaller   unMarshalCaptureFilter
	customHolder   CaptureCustom
	ethernetHolder CaptureEthernet
	vlanHolder     CaptureVlan
	ipv4Holder     CaptureIpv4
	ipv6Holder     CaptureIpv6
}

func NewCaptureFilter() CaptureFilter {
	obj := captureFilter{obj: &otg.CaptureFilter{}}
	obj.setDefault()
	return &obj
}

func (obj *captureFilter) msg() *otg.CaptureFilter {
	return obj.obj
}

func (obj *captureFilter) setMsg(msg *otg.CaptureFilter) CaptureFilter {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalcaptureFilter struct {
	obj *captureFilter
}

type marshalCaptureFilter interface {
	// ToProto marshals CaptureFilter to protobuf object *otg.CaptureFilter
	ToProto() (*otg.CaptureFilter, error)
	// ToPbText marshals CaptureFilter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals CaptureFilter to YAML text
	ToYaml() (string, error)
	// ToJson marshals CaptureFilter to JSON text
	ToJson() (string, error)
}

type unMarshalcaptureFilter struct {
	obj *captureFilter
}

type unMarshalCaptureFilter interface {
	// FromProto unmarshals CaptureFilter from protobuf object *otg.CaptureFilter
	FromProto(msg *otg.CaptureFilter) (CaptureFilter, error)
	// FromPbText unmarshals CaptureFilter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals CaptureFilter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals CaptureFilter from JSON text
	FromJson(value string) error
}

func (obj *captureFilter) Marshal() marshalCaptureFilter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalcaptureFilter{obj: obj}
	}
	return obj.marshaller
}

func (obj *captureFilter) Unmarshal() unMarshalCaptureFilter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalcaptureFilter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalcaptureFilter) ToProto() (*otg.CaptureFilter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalcaptureFilter) FromProto(msg *otg.CaptureFilter) (CaptureFilter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalcaptureFilter) ToPbText() (string, error) {
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

func (m *unMarshalcaptureFilter) FromPbText(value string) error {
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

func (m *marshalcaptureFilter) ToYaml() (string, error) {
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

func (m *unMarshalcaptureFilter) FromYaml(value string) error {
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

func (m *marshalcaptureFilter) ToJson() (string, error) {
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

func (m *unMarshalcaptureFilter) FromJson(value string) error {
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

func (obj *captureFilter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *captureFilter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *captureFilter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *captureFilter) Clone() (CaptureFilter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewCaptureFilter()
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

func (obj *captureFilter) setNil() {
	obj.customHolder = nil
	obj.ethernetHolder = nil
	obj.vlanHolder = nil
	obj.ipv4Holder = nil
	obj.ipv6Holder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// CaptureFilter is configuration for capture filters
type CaptureFilter interface {
	Validation
	// msg marshals CaptureFilter to protobuf object *otg.CaptureFilter
	// and doesn't set defaults
	msg() *otg.CaptureFilter
	// setMsg unmarshals CaptureFilter from protobuf object *otg.CaptureFilter
	// and doesn't set defaults
	setMsg(*otg.CaptureFilter) CaptureFilter
	// provides marshal interface
	Marshal() marshalCaptureFilter
	// provides unmarshal interface
	Unmarshal() unMarshalCaptureFilter
	// validate validates CaptureFilter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (CaptureFilter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns CaptureFilterChoiceEnum, set in CaptureFilter
	Choice() CaptureFilterChoiceEnum
	// setChoice assigns CaptureFilterChoiceEnum provided by user to CaptureFilter
	setChoice(value CaptureFilterChoiceEnum) CaptureFilter
	// HasChoice checks if Choice has been set in CaptureFilter
	HasChoice() bool
	// Custom returns CaptureCustom, set in CaptureFilter.
	// CaptureCustom is description is TBD
	Custom() CaptureCustom
	// SetCustom assigns CaptureCustom provided by user to CaptureFilter.
	// CaptureCustom is description is TBD
	SetCustom(value CaptureCustom) CaptureFilter
	// HasCustom checks if Custom has been set in CaptureFilter
	HasCustom() bool
	// Ethernet returns CaptureEthernet, set in CaptureFilter.
	// CaptureEthernet is description is TBD
	Ethernet() CaptureEthernet
	// SetEthernet assigns CaptureEthernet provided by user to CaptureFilter.
	// CaptureEthernet is description is TBD
	SetEthernet(value CaptureEthernet) CaptureFilter
	// HasEthernet checks if Ethernet has been set in CaptureFilter
	HasEthernet() bool
	// Vlan returns CaptureVlan, set in CaptureFilter.
	// CaptureVlan is description is TBD
	Vlan() CaptureVlan
	// SetVlan assigns CaptureVlan provided by user to CaptureFilter.
	// CaptureVlan is description is TBD
	SetVlan(value CaptureVlan) CaptureFilter
	// HasVlan checks if Vlan has been set in CaptureFilter
	HasVlan() bool
	// Ipv4 returns CaptureIpv4, set in CaptureFilter.
	// CaptureIpv4 is description is TBD
	Ipv4() CaptureIpv4
	// SetIpv4 assigns CaptureIpv4 provided by user to CaptureFilter.
	// CaptureIpv4 is description is TBD
	SetIpv4(value CaptureIpv4) CaptureFilter
	// HasIpv4 checks if Ipv4 has been set in CaptureFilter
	HasIpv4() bool
	// Ipv6 returns CaptureIpv6, set in CaptureFilter.
	// CaptureIpv6 is description is TBD
	Ipv6() CaptureIpv6
	// SetIpv6 assigns CaptureIpv6 provided by user to CaptureFilter.
	// CaptureIpv6 is description is TBD
	SetIpv6(value CaptureIpv6) CaptureFilter
	// HasIpv6 checks if Ipv6 has been set in CaptureFilter
	HasIpv6() bool
	setNil()
}

type CaptureFilterChoiceEnum string

// Enum of Choice on CaptureFilter
var CaptureFilterChoice = struct {
	CUSTOM   CaptureFilterChoiceEnum
	ETHERNET CaptureFilterChoiceEnum
	VLAN     CaptureFilterChoiceEnum
	IPV4     CaptureFilterChoiceEnum
	IPV6     CaptureFilterChoiceEnum
}{
	CUSTOM:   CaptureFilterChoiceEnum("custom"),
	ETHERNET: CaptureFilterChoiceEnum("ethernet"),
	VLAN:     CaptureFilterChoiceEnum("vlan"),
	IPV4:     CaptureFilterChoiceEnum("ipv4"),
	IPV6:     CaptureFilterChoiceEnum("ipv6"),
}

func (obj *captureFilter) Choice() CaptureFilterChoiceEnum {
	return CaptureFilterChoiceEnum(obj.obj.Choice.Enum().String())
}

// The type of capture filter.
// Choice returns a string
func (obj *captureFilter) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *captureFilter) setChoice(value CaptureFilterChoiceEnum) CaptureFilter {
	intValue, ok := otg.CaptureFilter_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on CaptureFilterChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.CaptureFilter_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Ipv6 = nil
	obj.ipv6Holder = nil
	obj.obj.Ipv4 = nil
	obj.ipv4Holder = nil
	obj.obj.Vlan = nil
	obj.vlanHolder = nil
	obj.obj.Ethernet = nil
	obj.ethernetHolder = nil
	obj.obj.Custom = nil
	obj.customHolder = nil

	if value == CaptureFilterChoice.CUSTOM {
		obj.obj.Custom = NewCaptureCustom().msg()
	}

	if value == CaptureFilterChoice.ETHERNET {
		obj.obj.Ethernet = NewCaptureEthernet().msg()
	}

	if value == CaptureFilterChoice.VLAN {
		obj.obj.Vlan = NewCaptureVlan().msg()
	}

	if value == CaptureFilterChoice.IPV4 {
		obj.obj.Ipv4 = NewCaptureIpv4().msg()
	}

	if value == CaptureFilterChoice.IPV6 {
		obj.obj.Ipv6 = NewCaptureIpv6().msg()
	}

	return obj
}

// Offset from last filter in the list. If no filters are present it is offset from position 0. Multiple custom filters can be present, the length of each custom filter is the length of the value being filtered.
// Custom returns a CaptureCustom
func (obj *captureFilter) Custom() CaptureCustom {
	if obj.obj.Custom == nil {
		obj.setChoice(CaptureFilterChoice.CUSTOM)
	}
	if obj.customHolder == nil {
		obj.customHolder = &captureCustom{obj: obj.obj.Custom}
	}
	return obj.customHolder
}

// Offset from last filter in the list. If no filters are present it is offset from position 0. Multiple custom filters can be present, the length of each custom filter is the length of the value being filtered.
// Custom returns a CaptureCustom
func (obj *captureFilter) HasCustom() bool {
	return obj.obj.Custom != nil
}

// Offset from last filter in the list. If no filters are present it is offset from position 0. Multiple custom filters can be present, the length of each custom filter is the length of the value being filtered.
// SetCustom sets the CaptureCustom value in the CaptureFilter object
func (obj *captureFilter) SetCustom(value CaptureCustom) CaptureFilter {
	obj.setChoice(CaptureFilterChoice.CUSTOM)
	obj.customHolder = nil
	obj.obj.Custom = value.msg()

	return obj
}

// description is TBD
// Ethernet returns a CaptureEthernet
func (obj *captureFilter) Ethernet() CaptureEthernet {
	if obj.obj.Ethernet == nil {
		obj.setChoice(CaptureFilterChoice.ETHERNET)
	}
	if obj.ethernetHolder == nil {
		obj.ethernetHolder = &captureEthernet{obj: obj.obj.Ethernet}
	}
	return obj.ethernetHolder
}

// description is TBD
// Ethernet returns a CaptureEthernet
func (obj *captureFilter) HasEthernet() bool {
	return obj.obj.Ethernet != nil
}

// description is TBD
// SetEthernet sets the CaptureEthernet value in the CaptureFilter object
func (obj *captureFilter) SetEthernet(value CaptureEthernet) CaptureFilter {
	obj.setChoice(CaptureFilterChoice.ETHERNET)
	obj.ethernetHolder = nil
	obj.obj.Ethernet = value.msg()

	return obj
}

// description is TBD
// Vlan returns a CaptureVlan
func (obj *captureFilter) Vlan() CaptureVlan {
	if obj.obj.Vlan == nil {
		obj.setChoice(CaptureFilterChoice.VLAN)
	}
	if obj.vlanHolder == nil {
		obj.vlanHolder = &captureVlan{obj: obj.obj.Vlan}
	}
	return obj.vlanHolder
}

// description is TBD
// Vlan returns a CaptureVlan
func (obj *captureFilter) HasVlan() bool {
	return obj.obj.Vlan != nil
}

// description is TBD
// SetVlan sets the CaptureVlan value in the CaptureFilter object
func (obj *captureFilter) SetVlan(value CaptureVlan) CaptureFilter {
	obj.setChoice(CaptureFilterChoice.VLAN)
	obj.vlanHolder = nil
	obj.obj.Vlan = value.msg()

	return obj
}

// description is TBD
// Ipv4 returns a CaptureIpv4
func (obj *captureFilter) Ipv4() CaptureIpv4 {
	if obj.obj.Ipv4 == nil {
		obj.setChoice(CaptureFilterChoice.IPV4)
	}
	if obj.ipv4Holder == nil {
		obj.ipv4Holder = &captureIpv4{obj: obj.obj.Ipv4}
	}
	return obj.ipv4Holder
}

// description is TBD
// Ipv4 returns a CaptureIpv4
func (obj *captureFilter) HasIpv4() bool {
	return obj.obj.Ipv4 != nil
}

// description is TBD
// SetIpv4 sets the CaptureIpv4 value in the CaptureFilter object
func (obj *captureFilter) SetIpv4(value CaptureIpv4) CaptureFilter {
	obj.setChoice(CaptureFilterChoice.IPV4)
	obj.ipv4Holder = nil
	obj.obj.Ipv4 = value.msg()

	return obj
}

// description is TBD
// Ipv6 returns a CaptureIpv6
func (obj *captureFilter) Ipv6() CaptureIpv6 {
	if obj.obj.Ipv6 == nil {
		obj.setChoice(CaptureFilterChoice.IPV6)
	}
	if obj.ipv6Holder == nil {
		obj.ipv6Holder = &captureIpv6{obj: obj.obj.Ipv6}
	}
	return obj.ipv6Holder
}

// description is TBD
// Ipv6 returns a CaptureIpv6
func (obj *captureFilter) HasIpv6() bool {
	return obj.obj.Ipv6 != nil
}

// description is TBD
// SetIpv6 sets the CaptureIpv6 value in the CaptureFilter object
func (obj *captureFilter) SetIpv6(value CaptureIpv6) CaptureFilter {
	obj.setChoice(CaptureFilterChoice.IPV6)
	obj.ipv6Holder = nil
	obj.obj.Ipv6 = value.msg()

	return obj
}

func (obj *captureFilter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {

		obj.Custom().validateObj(vObj, set_default)
	}

	if obj.obj.Ethernet != nil {

		obj.Ethernet().validateObj(vObj, set_default)
	}

	if obj.obj.Vlan != nil {

		obj.Vlan().validateObj(vObj, set_default)
	}

	if obj.obj.Ipv4 != nil {

		obj.Ipv4().validateObj(vObj, set_default)
	}

	if obj.obj.Ipv6 != nil {

		obj.Ipv6().validateObj(vObj, set_default)
	}

}

func (obj *captureFilter) setDefault() {
	var choices_set int = 0
	var choice CaptureFilterChoiceEnum

	if obj.obj.Custom != nil {
		choices_set += 1
		choice = CaptureFilterChoice.CUSTOM
	}

	if obj.obj.Ethernet != nil {
		choices_set += 1
		choice = CaptureFilterChoice.ETHERNET
	}

	if obj.obj.Vlan != nil {
		choices_set += 1
		choice = CaptureFilterChoice.VLAN
	}

	if obj.obj.Ipv4 != nil {
		choices_set += 1
		choice = CaptureFilterChoice.IPV4
	}

	if obj.obj.Ipv6 != nil {
		choices_set += 1
		choice = CaptureFilterChoice.IPV6
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(CaptureFilterChoice.CUSTOM)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in CaptureFilter")
			}
		} else {
			intVal := otg.CaptureFilter_Choice_Enum_value[string(choice)]
			enumValue := otg.CaptureFilter_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
