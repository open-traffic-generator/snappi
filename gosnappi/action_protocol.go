package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionProtocol *****
type actionProtocol struct {
	validation
	obj          *otg.ActionProtocol
	marshaller   marshalActionProtocol
	unMarshaller unMarshalActionProtocol
	ipv4Holder   ActionProtocolIpv4
	ipv6Holder   ActionProtocolIpv6
	bgpHolder    ActionProtocolBgp
	isisHolder   ActionProtocolIsis
}

func NewActionProtocol() ActionProtocol {
	obj := actionProtocol{obj: &otg.ActionProtocol{}}
	obj.setDefault()
	return &obj
}

func (obj *actionProtocol) msg() *otg.ActionProtocol {
	return obj.obj
}

func (obj *actionProtocol) setMsg(msg *otg.ActionProtocol) ActionProtocol {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionProtocol struct {
	obj *actionProtocol
}

type marshalActionProtocol interface {
	// ToProto marshals ActionProtocol to protobuf object *otg.ActionProtocol
	ToProto() (*otg.ActionProtocol, error)
	// ToPbText marshals ActionProtocol to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionProtocol to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionProtocol to JSON text
	ToJson() (string, error)
}

type unMarshalactionProtocol struct {
	obj *actionProtocol
}

type unMarshalActionProtocol interface {
	// FromProto unmarshals ActionProtocol from protobuf object *otg.ActionProtocol
	FromProto(msg *otg.ActionProtocol) (ActionProtocol, error)
	// FromPbText unmarshals ActionProtocol from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionProtocol from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionProtocol from JSON text
	FromJson(value string) error
}

func (obj *actionProtocol) Marshal() marshalActionProtocol {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionProtocol{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionProtocol) Unmarshal() unMarshalActionProtocol {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionProtocol{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionProtocol) ToProto() (*otg.ActionProtocol, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionProtocol) FromProto(msg *otg.ActionProtocol) (ActionProtocol, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionProtocol) ToPbText() (string, error) {
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

func (m *unMarshalactionProtocol) FromPbText(value string) error {
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

func (m *marshalactionProtocol) ToYaml() (string, error) {
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

func (m *unMarshalactionProtocol) FromYaml(value string) error {
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

func (m *marshalactionProtocol) ToJson() (string, error) {
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

func (m *unMarshalactionProtocol) FromJson(value string) error {
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

func (obj *actionProtocol) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionProtocol) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionProtocol) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionProtocol) Clone() (ActionProtocol, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionProtocol()
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

func (obj *actionProtocol) setNil() {
	obj.ipv4Holder = nil
	obj.ipv6Holder = nil
	obj.bgpHolder = nil
	obj.isisHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ActionProtocol is actions associated with protocols on configured resources.
type ActionProtocol interface {
	Validation
	// msg marshals ActionProtocol to protobuf object *otg.ActionProtocol
	// and doesn't set defaults
	msg() *otg.ActionProtocol
	// setMsg unmarshals ActionProtocol from protobuf object *otg.ActionProtocol
	// and doesn't set defaults
	setMsg(*otg.ActionProtocol) ActionProtocol
	// provides marshal interface
	Marshal() marshalActionProtocol
	// provides unmarshal interface
	Unmarshal() unMarshalActionProtocol
	// validate validates ActionProtocol
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionProtocol, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ActionProtocolChoiceEnum, set in ActionProtocol
	Choice() ActionProtocolChoiceEnum
	// setChoice assigns ActionProtocolChoiceEnum provided by user to ActionProtocol
	setChoice(value ActionProtocolChoiceEnum) ActionProtocol
	// Ipv4 returns ActionProtocolIpv4, set in ActionProtocol.
	// ActionProtocolIpv4 is actions associated with IPv4 on configured resources.
	Ipv4() ActionProtocolIpv4
	// SetIpv4 assigns ActionProtocolIpv4 provided by user to ActionProtocol.
	// ActionProtocolIpv4 is actions associated with IPv4 on configured resources.
	SetIpv4(value ActionProtocolIpv4) ActionProtocol
	// HasIpv4 checks if Ipv4 has been set in ActionProtocol
	HasIpv4() bool
	// Ipv6 returns ActionProtocolIpv6, set in ActionProtocol.
	// ActionProtocolIpv6 is actions associated with IPv6 on configured resources.
	Ipv6() ActionProtocolIpv6
	// SetIpv6 assigns ActionProtocolIpv6 provided by user to ActionProtocol.
	// ActionProtocolIpv6 is actions associated with IPv6 on configured resources.
	SetIpv6(value ActionProtocolIpv6) ActionProtocol
	// HasIpv6 checks if Ipv6 has been set in ActionProtocol
	HasIpv6() bool
	// Bgp returns ActionProtocolBgp, set in ActionProtocol.
	// ActionProtocolBgp is actions associated with BGP on configured resources.
	Bgp() ActionProtocolBgp
	// SetBgp assigns ActionProtocolBgp provided by user to ActionProtocol.
	// ActionProtocolBgp is actions associated with BGP on configured resources.
	SetBgp(value ActionProtocolBgp) ActionProtocol
	// HasBgp checks if Bgp has been set in ActionProtocol
	HasBgp() bool
	// Isis returns ActionProtocolIsis, set in ActionProtocol.
	// ActionProtocolIsis is actions associated with IS-IS on configured resources.
	Isis() ActionProtocolIsis
	// SetIsis assigns ActionProtocolIsis provided by user to ActionProtocol.
	// ActionProtocolIsis is actions associated with IS-IS on configured resources.
	SetIsis(value ActionProtocolIsis) ActionProtocol
	// HasIsis checks if Isis has been set in ActionProtocol
	HasIsis() bool
	setNil()
}

type ActionProtocolChoiceEnum string

// Enum of Choice on ActionProtocol
var ActionProtocolChoice = struct {
	IPV4 ActionProtocolChoiceEnum
	IPV6 ActionProtocolChoiceEnum
	BGP  ActionProtocolChoiceEnum
	ISIS ActionProtocolChoiceEnum
}{
	IPV4: ActionProtocolChoiceEnum("ipv4"),
	IPV6: ActionProtocolChoiceEnum("ipv6"),
	BGP:  ActionProtocolChoiceEnum("bgp"),
	ISIS: ActionProtocolChoiceEnum("isis"),
}

func (obj *actionProtocol) Choice() ActionProtocolChoiceEnum {
	return ActionProtocolChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *actionProtocol) setChoice(value ActionProtocolChoiceEnum) ActionProtocol {
	intValue, ok := otg.ActionProtocol_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ActionProtocolChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ActionProtocol_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Isis = nil
	obj.isisHolder = nil
	obj.obj.Bgp = nil
	obj.bgpHolder = nil
	obj.obj.Ipv6 = nil
	obj.ipv6Holder = nil
	obj.obj.Ipv4 = nil
	obj.ipv4Holder = nil

	if value == ActionProtocolChoice.IPV4 {
		obj.obj.Ipv4 = NewActionProtocolIpv4().msg()
	}

	if value == ActionProtocolChoice.IPV6 {
		obj.obj.Ipv6 = NewActionProtocolIpv6().msg()
	}

	if value == ActionProtocolChoice.BGP {
		obj.obj.Bgp = NewActionProtocolBgp().msg()
	}

	if value == ActionProtocolChoice.ISIS {
		obj.obj.Isis = NewActionProtocolIsis().msg()
	}

	return obj
}

// description is TBD
// Ipv4 returns a ActionProtocolIpv4
func (obj *actionProtocol) Ipv4() ActionProtocolIpv4 {
	if obj.obj.Ipv4 == nil {
		obj.setChoice(ActionProtocolChoice.IPV4)
	}
	if obj.ipv4Holder == nil {
		obj.ipv4Holder = &actionProtocolIpv4{obj: obj.obj.Ipv4}
	}
	return obj.ipv4Holder
}

// description is TBD
// Ipv4 returns a ActionProtocolIpv4
func (obj *actionProtocol) HasIpv4() bool {
	return obj.obj.Ipv4 != nil
}

// description is TBD
// SetIpv4 sets the ActionProtocolIpv4 value in the ActionProtocol object
func (obj *actionProtocol) SetIpv4(value ActionProtocolIpv4) ActionProtocol {
	obj.setChoice(ActionProtocolChoice.IPV4)
	obj.ipv4Holder = nil
	obj.obj.Ipv4 = value.msg()

	return obj
}

// description is TBD
// Ipv6 returns a ActionProtocolIpv6
func (obj *actionProtocol) Ipv6() ActionProtocolIpv6 {
	if obj.obj.Ipv6 == nil {
		obj.setChoice(ActionProtocolChoice.IPV6)
	}
	if obj.ipv6Holder == nil {
		obj.ipv6Holder = &actionProtocolIpv6{obj: obj.obj.Ipv6}
	}
	return obj.ipv6Holder
}

// description is TBD
// Ipv6 returns a ActionProtocolIpv6
func (obj *actionProtocol) HasIpv6() bool {
	return obj.obj.Ipv6 != nil
}

// description is TBD
// SetIpv6 sets the ActionProtocolIpv6 value in the ActionProtocol object
func (obj *actionProtocol) SetIpv6(value ActionProtocolIpv6) ActionProtocol {
	obj.setChoice(ActionProtocolChoice.IPV6)
	obj.ipv6Holder = nil
	obj.obj.Ipv6 = value.msg()

	return obj
}

// description is TBD
// Bgp returns a ActionProtocolBgp
func (obj *actionProtocol) Bgp() ActionProtocolBgp {
	if obj.obj.Bgp == nil {
		obj.setChoice(ActionProtocolChoice.BGP)
	}
	if obj.bgpHolder == nil {
		obj.bgpHolder = &actionProtocolBgp{obj: obj.obj.Bgp}
	}
	return obj.bgpHolder
}

// description is TBD
// Bgp returns a ActionProtocolBgp
func (obj *actionProtocol) HasBgp() bool {
	return obj.obj.Bgp != nil
}

// description is TBD
// SetBgp sets the ActionProtocolBgp value in the ActionProtocol object
func (obj *actionProtocol) SetBgp(value ActionProtocolBgp) ActionProtocol {
	obj.setChoice(ActionProtocolChoice.BGP)
	obj.bgpHolder = nil
	obj.obj.Bgp = value.msg()

	return obj
}

// description is TBD
// Isis returns a ActionProtocolIsis
func (obj *actionProtocol) Isis() ActionProtocolIsis {
	if obj.obj.Isis == nil {
		obj.setChoice(ActionProtocolChoice.ISIS)
	}
	if obj.isisHolder == nil {
		obj.isisHolder = &actionProtocolIsis{obj: obj.obj.Isis}
	}
	return obj.isisHolder
}

// description is TBD
// Isis returns a ActionProtocolIsis
func (obj *actionProtocol) HasIsis() bool {
	return obj.obj.Isis != nil
}

// description is TBD
// SetIsis sets the ActionProtocolIsis value in the ActionProtocol object
func (obj *actionProtocol) SetIsis(value ActionProtocolIsis) ActionProtocol {
	obj.setChoice(ActionProtocolChoice.ISIS)
	obj.isisHolder = nil
	obj.obj.Isis = value.msg()

	return obj
}

func (obj *actionProtocol) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface ActionProtocol")
	}

	if obj.obj.Ipv4 != nil {

		obj.Ipv4().validateObj(vObj, set_default)
	}

	if obj.obj.Ipv6 != nil {

		obj.Ipv6().validateObj(vObj, set_default)
	}

	if obj.obj.Bgp != nil {

		obj.Bgp().validateObj(vObj, set_default)
	}

	if obj.obj.Isis != nil {

		obj.Isis().validateObj(vObj, set_default)
	}

}

func (obj *actionProtocol) setDefault() {
	var choices_set int = 0
	var choice ActionProtocolChoiceEnum

	if obj.obj.Ipv4 != nil {
		choices_set += 1
		choice = ActionProtocolChoice.IPV4
	}

	if obj.obj.Ipv6 != nil {
		choices_set += 1
		choice = ActionProtocolChoice.IPV6
	}

	if obj.obj.Bgp != nil {
		choices_set += 1
		choice = ActionProtocolChoice.BGP
	}

	if obj.obj.Isis != nil {
		choices_set += 1
		choice = ActionProtocolChoice.ISIS
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ActionProtocol")
			}
		} else {
			intVal := otg.ActionProtocol_Choice_Enum_value[string(choice)]
			enumValue := otg.ActionProtocol_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
