package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionResponseProtocol *****
type actionResponseProtocol struct {
	validation
	obj          *otg.ActionResponseProtocol
	marshaller   marshalActionResponseProtocol
	unMarshaller unMarshalActionResponseProtocol
	ipv4Holder   ActionResponseProtocolIpv4
	ipv6Holder   ActionResponseProtocolIpv6
}

func NewActionResponseProtocol() ActionResponseProtocol {
	obj := actionResponseProtocol{obj: &otg.ActionResponseProtocol{}}
	obj.setDefault()
	return &obj
}

func (obj *actionResponseProtocol) msg() *otg.ActionResponseProtocol {
	return obj.obj
}

func (obj *actionResponseProtocol) setMsg(msg *otg.ActionResponseProtocol) ActionResponseProtocol {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionResponseProtocol struct {
	obj *actionResponseProtocol
}

type marshalActionResponseProtocol interface {
	// ToProto marshals ActionResponseProtocol to protobuf object *otg.ActionResponseProtocol
	ToProto() (*otg.ActionResponseProtocol, error)
	// ToPbText marshals ActionResponseProtocol to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionResponseProtocol to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionResponseProtocol to JSON text
	ToJson() (string, error)
}

type unMarshalactionResponseProtocol struct {
	obj *actionResponseProtocol
}

type unMarshalActionResponseProtocol interface {
	// FromProto unmarshals ActionResponseProtocol from protobuf object *otg.ActionResponseProtocol
	FromProto(msg *otg.ActionResponseProtocol) (ActionResponseProtocol, error)
	// FromPbText unmarshals ActionResponseProtocol from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionResponseProtocol from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionResponseProtocol from JSON text
	FromJson(value string) error
}

func (obj *actionResponseProtocol) Marshal() marshalActionResponseProtocol {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionResponseProtocol{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionResponseProtocol) Unmarshal() unMarshalActionResponseProtocol {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionResponseProtocol{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionResponseProtocol) ToProto() (*otg.ActionResponseProtocol, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionResponseProtocol) FromProto(msg *otg.ActionResponseProtocol) (ActionResponseProtocol, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionResponseProtocol) ToPbText() (string, error) {
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

func (m *unMarshalactionResponseProtocol) FromPbText(value string) error {
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

func (m *marshalactionResponseProtocol) ToYaml() (string, error) {
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

func (m *unMarshalactionResponseProtocol) FromYaml(value string) error {
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

func (m *marshalactionResponseProtocol) ToJson() (string, error) {
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

func (m *unMarshalactionResponseProtocol) FromJson(value string) error {
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

func (obj *actionResponseProtocol) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionResponseProtocol) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionResponseProtocol) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionResponseProtocol) Clone() (ActionResponseProtocol, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionResponseProtocol()
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

func (obj *actionResponseProtocol) setNil() {
	obj.ipv4Holder = nil
	obj.ipv6Holder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ActionResponseProtocol is response for actions associated with protocols on configured resources.
type ActionResponseProtocol interface {
	Validation
	// msg marshals ActionResponseProtocol to protobuf object *otg.ActionResponseProtocol
	// and doesn't set defaults
	msg() *otg.ActionResponseProtocol
	// setMsg unmarshals ActionResponseProtocol from protobuf object *otg.ActionResponseProtocol
	// and doesn't set defaults
	setMsg(*otg.ActionResponseProtocol) ActionResponseProtocol
	// provides marshal interface
	Marshal() marshalActionResponseProtocol
	// provides unmarshal interface
	Unmarshal() unMarshalActionResponseProtocol
	// validate validates ActionResponseProtocol
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionResponseProtocol, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ActionResponseProtocolChoiceEnum, set in ActionResponseProtocol
	Choice() ActionResponseProtocolChoiceEnum
	// setChoice assigns ActionResponseProtocolChoiceEnum provided by user to ActionResponseProtocol
	setChoice(value ActionResponseProtocolChoiceEnum) ActionResponseProtocol
	// Ipv4 returns ActionResponseProtocolIpv4, set in ActionResponseProtocol.
	// ActionResponseProtocolIpv4 is response for actions associated with IPv4 on configured resources.
	Ipv4() ActionResponseProtocolIpv4
	// SetIpv4 assigns ActionResponseProtocolIpv4 provided by user to ActionResponseProtocol.
	// ActionResponseProtocolIpv4 is response for actions associated with IPv4 on configured resources.
	SetIpv4(value ActionResponseProtocolIpv4) ActionResponseProtocol
	// HasIpv4 checks if Ipv4 has been set in ActionResponseProtocol
	HasIpv4() bool
	// Ipv6 returns ActionResponseProtocolIpv6, set in ActionResponseProtocol.
	// ActionResponseProtocolIpv6 is response for actions associated with IPv6 on configured resources.
	Ipv6() ActionResponseProtocolIpv6
	// SetIpv6 assigns ActionResponseProtocolIpv6 provided by user to ActionResponseProtocol.
	// ActionResponseProtocolIpv6 is response for actions associated with IPv6 on configured resources.
	SetIpv6(value ActionResponseProtocolIpv6) ActionResponseProtocol
	// HasIpv6 checks if Ipv6 has been set in ActionResponseProtocol
	HasIpv6() bool
	setNil()
}

type ActionResponseProtocolChoiceEnum string

// Enum of Choice on ActionResponseProtocol
var ActionResponseProtocolChoice = struct {
	IPV4 ActionResponseProtocolChoiceEnum
	IPV6 ActionResponseProtocolChoiceEnum
}{
	IPV4: ActionResponseProtocolChoiceEnum("ipv4"),
	IPV6: ActionResponseProtocolChoiceEnum("ipv6"),
}

func (obj *actionResponseProtocol) Choice() ActionResponseProtocolChoiceEnum {
	return ActionResponseProtocolChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *actionResponseProtocol) setChoice(value ActionResponseProtocolChoiceEnum) ActionResponseProtocol {
	intValue, ok := otg.ActionResponseProtocol_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ActionResponseProtocolChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ActionResponseProtocol_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Ipv6 = nil
	obj.ipv6Holder = nil
	obj.obj.Ipv4 = nil
	obj.ipv4Holder = nil

	if value == ActionResponseProtocolChoice.IPV4 {
		obj.obj.Ipv4 = NewActionResponseProtocolIpv4().msg()
	}

	if value == ActionResponseProtocolChoice.IPV6 {
		obj.obj.Ipv6 = NewActionResponseProtocolIpv6().msg()
	}

	return obj
}

// description is TBD
// Ipv4 returns a ActionResponseProtocolIpv4
func (obj *actionResponseProtocol) Ipv4() ActionResponseProtocolIpv4 {
	if obj.obj.Ipv4 == nil {
		obj.setChoice(ActionResponseProtocolChoice.IPV4)
	}
	if obj.ipv4Holder == nil {
		obj.ipv4Holder = &actionResponseProtocolIpv4{obj: obj.obj.Ipv4}
	}
	return obj.ipv4Holder
}

// description is TBD
// Ipv4 returns a ActionResponseProtocolIpv4
func (obj *actionResponseProtocol) HasIpv4() bool {
	return obj.obj.Ipv4 != nil
}

// description is TBD
// SetIpv4 sets the ActionResponseProtocolIpv4 value in the ActionResponseProtocol object
func (obj *actionResponseProtocol) SetIpv4(value ActionResponseProtocolIpv4) ActionResponseProtocol {
	obj.setChoice(ActionResponseProtocolChoice.IPV4)
	obj.ipv4Holder = nil
	obj.obj.Ipv4 = value.msg()

	return obj
}

// description is TBD
// Ipv6 returns a ActionResponseProtocolIpv6
func (obj *actionResponseProtocol) Ipv6() ActionResponseProtocolIpv6 {
	if obj.obj.Ipv6 == nil {
		obj.setChoice(ActionResponseProtocolChoice.IPV6)
	}
	if obj.ipv6Holder == nil {
		obj.ipv6Holder = &actionResponseProtocolIpv6{obj: obj.obj.Ipv6}
	}
	return obj.ipv6Holder
}

// description is TBD
// Ipv6 returns a ActionResponseProtocolIpv6
func (obj *actionResponseProtocol) HasIpv6() bool {
	return obj.obj.Ipv6 != nil
}

// description is TBD
// SetIpv6 sets the ActionResponseProtocolIpv6 value in the ActionResponseProtocol object
func (obj *actionResponseProtocol) SetIpv6(value ActionResponseProtocolIpv6) ActionResponseProtocol {
	obj.setChoice(ActionResponseProtocolChoice.IPV6)
	obj.ipv6Holder = nil
	obj.obj.Ipv6 = value.msg()

	return obj
}

func (obj *actionResponseProtocol) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface ActionResponseProtocol")
	}

	if obj.obj.Ipv4 != nil {

		obj.Ipv4().validateObj(vObj, set_default)
	}

	if obj.obj.Ipv6 != nil {

		obj.Ipv6().validateObj(vObj, set_default)
	}

}

func (obj *actionResponseProtocol) setDefault() {

}
