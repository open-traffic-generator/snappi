package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UltraEthernetCbfcVcMapping *****
type ultraEthernetCbfcVcMapping struct {
	validation
	obj              *otg.UltraEthernetCbfcVcMapping
	marshaller       marshalUltraEthernetCbfcVcMapping
	unMarshaller     unMarshalUltraEthernetCbfcVcMapping
	vlanPcpDeiHolder UltraEthernetCbfcVcMappingVlanPcpDei
	dscpHolder       UltraEthernetCbfcVcMappingDscp
}

func NewUltraEthernetCbfcVcMapping() UltraEthernetCbfcVcMapping {
	obj := ultraEthernetCbfcVcMapping{obj: &otg.UltraEthernetCbfcVcMapping{}}
	obj.setDefault()
	return &obj
}

func (obj *ultraEthernetCbfcVcMapping) msg() *otg.UltraEthernetCbfcVcMapping {
	return obj.obj
}

func (obj *ultraEthernetCbfcVcMapping) setMsg(msg *otg.UltraEthernetCbfcVcMapping) UltraEthernetCbfcVcMapping {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalultraEthernetCbfcVcMapping struct {
	obj *ultraEthernetCbfcVcMapping
}

type marshalUltraEthernetCbfcVcMapping interface {
	// ToProto marshals UltraEthernetCbfcVcMapping to protobuf object *otg.UltraEthernetCbfcVcMapping
	ToProto() (*otg.UltraEthernetCbfcVcMapping, error)
	// ToPbText marshals UltraEthernetCbfcVcMapping to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UltraEthernetCbfcVcMapping to YAML text
	ToYaml() (string, error)
	// ToJson marshals UltraEthernetCbfcVcMapping to JSON text
	ToJson() (string, error)
}

type unMarshalultraEthernetCbfcVcMapping struct {
	obj *ultraEthernetCbfcVcMapping
}

type unMarshalUltraEthernetCbfcVcMapping interface {
	// FromProto unmarshals UltraEthernetCbfcVcMapping from protobuf object *otg.UltraEthernetCbfcVcMapping
	FromProto(msg *otg.UltraEthernetCbfcVcMapping) (UltraEthernetCbfcVcMapping, error)
	// FromPbText unmarshals UltraEthernetCbfcVcMapping from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UltraEthernetCbfcVcMapping from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UltraEthernetCbfcVcMapping from JSON text
	FromJson(value string) error
}

func (obj *ultraEthernetCbfcVcMapping) Marshal() marshalUltraEthernetCbfcVcMapping {
	if obj.marshaller == nil {
		obj.marshaller = &marshalultraEthernetCbfcVcMapping{obj: obj}
	}
	return obj.marshaller
}

func (obj *ultraEthernetCbfcVcMapping) Unmarshal() unMarshalUltraEthernetCbfcVcMapping {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalultraEthernetCbfcVcMapping{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalultraEthernetCbfcVcMapping) ToProto() (*otg.UltraEthernetCbfcVcMapping, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalultraEthernetCbfcVcMapping) FromProto(msg *otg.UltraEthernetCbfcVcMapping) (UltraEthernetCbfcVcMapping, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalultraEthernetCbfcVcMapping) ToPbText() (string, error) {
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

func (m *unMarshalultraEthernetCbfcVcMapping) FromPbText(value string) error {
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

func (m *marshalultraEthernetCbfcVcMapping) ToYaml() (string, error) {
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

func (m *unMarshalultraEthernetCbfcVcMapping) FromYaml(value string) error {
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

func (m *marshalultraEthernetCbfcVcMapping) ToJson() (string, error) {
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

func (m *unMarshalultraEthernetCbfcVcMapping) FromJson(value string) error {
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

func (obj *ultraEthernetCbfcVcMapping) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfcVcMapping) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfcVcMapping) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ultraEthernetCbfcVcMapping) Clone() (UltraEthernetCbfcVcMapping, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUltraEthernetCbfcVcMapping()
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

func (obj *ultraEthernetCbfcVcMapping) setNil() {
	obj.vlanPcpDeiHolder = nil
	obj.dscpHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// UltraEthernetCbfcVcMapping is the packet header field used to classify packets into a virtual channel, and
// the associated values. At a minimum, mapping from the mac.vlan.pcp_dei and
// ip.dscp header fields is supported.
type UltraEthernetCbfcVcMapping interface {
	Validation
	// msg marshals UltraEthernetCbfcVcMapping to protobuf object *otg.UltraEthernetCbfcVcMapping
	// and doesn't set defaults
	msg() *otg.UltraEthernetCbfcVcMapping
	// setMsg unmarshals UltraEthernetCbfcVcMapping from protobuf object *otg.UltraEthernetCbfcVcMapping
	// and doesn't set defaults
	setMsg(*otg.UltraEthernetCbfcVcMapping) UltraEthernetCbfcVcMapping
	// provides marshal interface
	Marshal() marshalUltraEthernetCbfcVcMapping
	// provides unmarshal interface
	Unmarshal() unMarshalUltraEthernetCbfcVcMapping
	// validate validates UltraEthernetCbfcVcMapping
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UltraEthernetCbfcVcMapping, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns UltraEthernetCbfcVcMappingChoiceEnum, set in UltraEthernetCbfcVcMapping
	Choice() UltraEthernetCbfcVcMappingChoiceEnum
	// setChoice assigns UltraEthernetCbfcVcMappingChoiceEnum provided by user to UltraEthernetCbfcVcMapping
	setChoice(value UltraEthernetCbfcVcMappingChoiceEnum) UltraEthernetCbfcVcMapping
	// HasChoice checks if Choice has been set in UltraEthernetCbfcVcMapping
	HasChoice() bool
	// VlanPcpDei returns UltraEthernetCbfcVcMappingVlanPcpDei, set in UltraEthernetCbfcVcMapping.
	// UltraEthernetCbfcVcMappingVlanPcpDei is vLAN PCP/DEI values that classify packets into this virtual channel.
	VlanPcpDei() UltraEthernetCbfcVcMappingVlanPcpDei
	// SetVlanPcpDei assigns UltraEthernetCbfcVcMappingVlanPcpDei provided by user to UltraEthernetCbfcVcMapping.
	// UltraEthernetCbfcVcMappingVlanPcpDei is vLAN PCP/DEI values that classify packets into this virtual channel.
	SetVlanPcpDei(value UltraEthernetCbfcVcMappingVlanPcpDei) UltraEthernetCbfcVcMapping
	// HasVlanPcpDei checks if VlanPcpDei has been set in UltraEthernetCbfcVcMapping
	HasVlanPcpDei() bool
	// Dscp returns UltraEthernetCbfcVcMappingDscp, set in UltraEthernetCbfcVcMapping.
	// UltraEthernetCbfcVcMappingDscp is iPv4/IPv6 DSCP values that classify packets into this virtual channel.
	Dscp() UltraEthernetCbfcVcMappingDscp
	// SetDscp assigns UltraEthernetCbfcVcMappingDscp provided by user to UltraEthernetCbfcVcMapping.
	// UltraEthernetCbfcVcMappingDscp is iPv4/IPv6 DSCP values that classify packets into this virtual channel.
	SetDscp(value UltraEthernetCbfcVcMappingDscp) UltraEthernetCbfcVcMapping
	// HasDscp checks if Dscp has been set in UltraEthernetCbfcVcMapping
	HasDscp() bool
	setNil()
}

type UltraEthernetCbfcVcMappingChoiceEnum string

// Enum of Choice on UltraEthernetCbfcVcMapping
var UltraEthernetCbfcVcMappingChoice = struct {
	VLAN_PCP_DEI UltraEthernetCbfcVcMappingChoiceEnum
	DSCP         UltraEthernetCbfcVcMappingChoiceEnum
}{
	VLAN_PCP_DEI: UltraEthernetCbfcVcMappingChoiceEnum("vlan_pcp_dei"),
	DSCP:         UltraEthernetCbfcVcMappingChoiceEnum("dscp"),
}

func (obj *ultraEthernetCbfcVcMapping) Choice() UltraEthernetCbfcVcMappingChoiceEnum {
	return UltraEthernetCbfcVcMappingChoiceEnum(obj.obj.Choice.Enum().String())
}

// The packet header field used to classify packets into this VC.
//
// - vlan_pcp_dei: classify by the mac.vlan.pcp_dei value(s).
// - dscp: classify by the ip.dscp value(s).
// Choice returns a string
func (obj *ultraEthernetCbfcVcMapping) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *ultraEthernetCbfcVcMapping) setChoice(value UltraEthernetCbfcVcMappingChoiceEnum) UltraEthernetCbfcVcMapping {
	intValue, ok := otg.UltraEthernetCbfcVcMapping_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on UltraEthernetCbfcVcMappingChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.UltraEthernetCbfcVcMapping_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Dscp = nil
	obj.dscpHolder = nil
	obj.obj.VlanPcpDei = nil
	obj.vlanPcpDeiHolder = nil

	if value == UltraEthernetCbfcVcMappingChoice.VLAN_PCP_DEI {
		obj.obj.VlanPcpDei = NewUltraEthernetCbfcVcMappingVlanPcpDei().msg()
	}

	if value == UltraEthernetCbfcVcMappingChoice.DSCP {
		obj.obj.Dscp = NewUltraEthernetCbfcVcMappingDscp().msg()
	}

	return obj
}

// description is TBD
// VlanPcpDei returns a UltraEthernetCbfcVcMappingVlanPcpDei
func (obj *ultraEthernetCbfcVcMapping) VlanPcpDei() UltraEthernetCbfcVcMappingVlanPcpDei {
	if obj.obj.VlanPcpDei == nil {
		obj.setChoice(UltraEthernetCbfcVcMappingChoice.VLAN_PCP_DEI)
	}
	if obj.vlanPcpDeiHolder == nil {
		obj.vlanPcpDeiHolder = &ultraEthernetCbfcVcMappingVlanPcpDei{obj: obj.obj.VlanPcpDei}
	}
	return obj.vlanPcpDeiHolder
}

// description is TBD
// VlanPcpDei returns a UltraEthernetCbfcVcMappingVlanPcpDei
func (obj *ultraEthernetCbfcVcMapping) HasVlanPcpDei() bool {
	return obj.obj.VlanPcpDei != nil
}

// description is TBD
// SetVlanPcpDei sets the UltraEthernetCbfcVcMappingVlanPcpDei value in the UltraEthernetCbfcVcMapping object
func (obj *ultraEthernetCbfcVcMapping) SetVlanPcpDei(value UltraEthernetCbfcVcMappingVlanPcpDei) UltraEthernetCbfcVcMapping {
	obj.setChoice(UltraEthernetCbfcVcMappingChoice.VLAN_PCP_DEI)
	obj.vlanPcpDeiHolder = nil
	obj.obj.VlanPcpDei = value.msg()

	return obj
}

// description is TBD
// Dscp returns a UltraEthernetCbfcVcMappingDscp
func (obj *ultraEthernetCbfcVcMapping) Dscp() UltraEthernetCbfcVcMappingDscp {
	if obj.obj.Dscp == nil {
		obj.setChoice(UltraEthernetCbfcVcMappingChoice.DSCP)
	}
	if obj.dscpHolder == nil {
		obj.dscpHolder = &ultraEthernetCbfcVcMappingDscp{obj: obj.obj.Dscp}
	}
	return obj.dscpHolder
}

// description is TBD
// Dscp returns a UltraEthernetCbfcVcMappingDscp
func (obj *ultraEthernetCbfcVcMapping) HasDscp() bool {
	return obj.obj.Dscp != nil
}

// description is TBD
// SetDscp sets the UltraEthernetCbfcVcMappingDscp value in the UltraEthernetCbfcVcMapping object
func (obj *ultraEthernetCbfcVcMapping) SetDscp(value UltraEthernetCbfcVcMappingDscp) UltraEthernetCbfcVcMapping {
	obj.setChoice(UltraEthernetCbfcVcMappingChoice.DSCP)
	obj.dscpHolder = nil
	obj.obj.Dscp = value.msg()

	return obj
}

func (obj *ultraEthernetCbfcVcMapping) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.VlanPcpDei != nil {

		obj.VlanPcpDei().validateObj(vObj, set_default)
	}

	if obj.obj.Dscp != nil {

		obj.Dscp().validateObj(vObj, set_default)
	}

}

func (obj *ultraEthernetCbfcVcMapping) setDefault() {
	var choices_set int = 0
	var choice UltraEthernetCbfcVcMappingChoiceEnum

	if obj.obj.VlanPcpDei != nil {
		choices_set += 1
		choice = UltraEthernetCbfcVcMappingChoice.VLAN_PCP_DEI
	}

	if obj.obj.Dscp != nil {
		choices_set += 1
		choice = UltraEthernetCbfcVcMappingChoice.DSCP
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(UltraEthernetCbfcVcMappingChoice.VLAN_PCP_DEI)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in UltraEthernetCbfcVcMapping")
			}
		} else {
			intVal := otg.UltraEthernetCbfcVcMapping_Choice_Enum_value[string(choice)]
			enumValue := otg.UltraEthernetCbfcVcMapping_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
