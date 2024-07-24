package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesNextHop *****
type bgpAttributesNextHop struct {
	validation
	obj                    *otg.BgpAttributesNextHop
	marshaller             marshalBgpAttributesNextHop
	unMarshaller           unMarshalBgpAttributesNextHop
	ipv6TwoAddressesHolder BgpAttributesNextHopIpv6TwoAddresses
}

func NewBgpAttributesNextHop() BgpAttributesNextHop {
	obj := bgpAttributesNextHop{obj: &otg.BgpAttributesNextHop{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesNextHop) msg() *otg.BgpAttributesNextHop {
	return obj.obj
}

func (obj *bgpAttributesNextHop) setMsg(msg *otg.BgpAttributesNextHop) BgpAttributesNextHop {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesNextHop struct {
	obj *bgpAttributesNextHop
}

type marshalBgpAttributesNextHop interface {
	// ToProto marshals BgpAttributesNextHop to protobuf object *otg.BgpAttributesNextHop
	ToProto() (*otg.BgpAttributesNextHop, error)
	// ToPbText marshals BgpAttributesNextHop to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesNextHop to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesNextHop to JSON text
	ToJson() (string, error)
}

type unMarshalbgpAttributesNextHop struct {
	obj *bgpAttributesNextHop
}

type unMarshalBgpAttributesNextHop interface {
	// FromProto unmarshals BgpAttributesNextHop from protobuf object *otg.BgpAttributesNextHop
	FromProto(msg *otg.BgpAttributesNextHop) (BgpAttributesNextHop, error)
	// FromPbText unmarshals BgpAttributesNextHop from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesNextHop from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesNextHop from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesNextHop) Marshal() marshalBgpAttributesNextHop {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesNextHop{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesNextHop) Unmarshal() unMarshalBgpAttributesNextHop {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesNextHop{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesNextHop) ToProto() (*otg.BgpAttributesNextHop, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesNextHop) FromProto(msg *otg.BgpAttributesNextHop) (BgpAttributesNextHop, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesNextHop) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesNextHop) FromPbText(value string) error {
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

func (m *marshalbgpAttributesNextHop) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesNextHop) FromYaml(value string) error {
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

func (m *marshalbgpAttributesNextHop) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesNextHop) FromJson(value string) error {
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

func (obj *bgpAttributesNextHop) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesNextHop) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesNextHop) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesNextHop) Clone() (BgpAttributesNextHop, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesNextHop()
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

func (obj *bgpAttributesNextHop) setNil() {
	obj.ipv6TwoAddressesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpAttributesNextHop is next hop to be sent inside MP_REACH NLRI or as the NEXT_HOP attribute if advertised as traditional NLRI.
type BgpAttributesNextHop interface {
	Validation
	// msg marshals BgpAttributesNextHop to protobuf object *otg.BgpAttributesNextHop
	// and doesn't set defaults
	msg() *otg.BgpAttributesNextHop
	// setMsg unmarshals BgpAttributesNextHop from protobuf object *otg.BgpAttributesNextHop
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesNextHop) BgpAttributesNextHop
	// provides marshal interface
	Marshal() marshalBgpAttributesNextHop
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesNextHop
	// validate validates BgpAttributesNextHop
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesNextHop, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns BgpAttributesNextHopChoiceEnum, set in BgpAttributesNextHop
	Choice() BgpAttributesNextHopChoiceEnum
	// setChoice assigns BgpAttributesNextHopChoiceEnum provided by user to BgpAttributesNextHop
	setChoice(value BgpAttributesNextHopChoiceEnum) BgpAttributesNextHop
	// Ipv4 returns string, set in BgpAttributesNextHop.
	Ipv4() string
	// SetIpv4 assigns string provided by user to BgpAttributesNextHop
	SetIpv4(value string) BgpAttributesNextHop
	// HasIpv4 checks if Ipv4 has been set in BgpAttributesNextHop
	HasIpv4() bool
	// Ipv6 returns string, set in BgpAttributesNextHop.
	Ipv6() string
	// SetIpv6 assigns string provided by user to BgpAttributesNextHop
	SetIpv6(value string) BgpAttributesNextHop
	// HasIpv6 checks if Ipv6 has been set in BgpAttributesNextHop
	HasIpv6() bool
	// Ipv6TwoAddresses returns BgpAttributesNextHopIpv6TwoAddresses, set in BgpAttributesNextHop.
	// BgpAttributesNextHopIpv6TwoAddresses is there is a specific scenario in which it is possible to receive a Global and Link Local address in the Next Hop
	// field in a MP_REACH attribute or in the NEXT_HOP attribute(RFC2545: Section 3).
	Ipv6TwoAddresses() BgpAttributesNextHopIpv6TwoAddresses
	// SetIpv6TwoAddresses assigns BgpAttributesNextHopIpv6TwoAddresses provided by user to BgpAttributesNextHop.
	// BgpAttributesNextHopIpv6TwoAddresses is there is a specific scenario in which it is possible to receive a Global and Link Local address in the Next Hop
	// field in a MP_REACH attribute or in the NEXT_HOP attribute(RFC2545: Section 3).
	SetIpv6TwoAddresses(value BgpAttributesNextHopIpv6TwoAddresses) BgpAttributesNextHop
	// HasIpv6TwoAddresses checks if Ipv6TwoAddresses has been set in BgpAttributesNextHop
	HasIpv6TwoAddresses() bool
	setNil()
}

type BgpAttributesNextHopChoiceEnum string

// Enum of Choice on BgpAttributesNextHop
var BgpAttributesNextHopChoice = struct {
	IPV4               BgpAttributesNextHopChoiceEnum
	IPV6               BgpAttributesNextHopChoiceEnum
	IPV6_TWO_ADDRESSES BgpAttributesNextHopChoiceEnum
}{
	IPV4:               BgpAttributesNextHopChoiceEnum("ipv4"),
	IPV6:               BgpAttributesNextHopChoiceEnum("ipv6"),
	IPV6_TWO_ADDRESSES: BgpAttributesNextHopChoiceEnum("ipv6_two_addresses"),
}

func (obj *bgpAttributesNextHop) Choice() BgpAttributesNextHopChoiceEnum {
	return BgpAttributesNextHopChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *bgpAttributesNextHop) setChoice(value BgpAttributesNextHopChoiceEnum) BgpAttributesNextHop {
	intValue, ok := otg.BgpAttributesNextHop_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpAttributesNextHopChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpAttributesNextHop_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Ipv6TwoAddresses = nil
	obj.ipv6TwoAddressesHolder = nil
	obj.obj.Ipv6 = nil
	obj.obj.Ipv4 = nil

	if value == BgpAttributesNextHopChoice.IPV4 {
		defaultValue := "0.0.0.0"
		obj.obj.Ipv4 = &defaultValue
	}

	if value == BgpAttributesNextHopChoice.IPV6 {
		defaultValue := "0::0"
		obj.obj.Ipv6 = &defaultValue
	}

	if value == BgpAttributesNextHopChoice.IPV6_TWO_ADDRESSES {
		obj.obj.Ipv6TwoAddresses = NewBgpAttributesNextHopIpv6TwoAddresses().msg()
	}

	return obj
}

// The 4 byte IPv4 address of the next-hop from which the route was received.
// Ipv4 returns a string
func (obj *bgpAttributesNextHop) Ipv4() string {

	if obj.obj.Ipv4 == nil {
		obj.setChoice(BgpAttributesNextHopChoice.IPV4)
	}

	return *obj.obj.Ipv4

}

// The 4 byte IPv4 address of the next-hop from which the route was received.
// Ipv4 returns a string
func (obj *bgpAttributesNextHop) HasIpv4() bool {
	return obj.obj.Ipv4 != nil
}

// The 4 byte IPv4 address of the next-hop from which the route was received.
// SetIpv4 sets the string value in the BgpAttributesNextHop object
func (obj *bgpAttributesNextHop) SetIpv4(value string) BgpAttributesNextHop {
	obj.setChoice(BgpAttributesNextHopChoice.IPV4)
	obj.obj.Ipv4 = &value
	return obj
}

// The 16 byte IPv6 address of the next-hop from which the route was received.
// Ipv6 returns a string
func (obj *bgpAttributesNextHop) Ipv6() string {

	if obj.obj.Ipv6 == nil {
		obj.setChoice(BgpAttributesNextHopChoice.IPV6)
	}

	return *obj.obj.Ipv6

}

// The 16 byte IPv6 address of the next-hop from which the route was received.
// Ipv6 returns a string
func (obj *bgpAttributesNextHop) HasIpv6() bool {
	return obj.obj.Ipv6 != nil
}

// The 16 byte IPv6 address of the next-hop from which the route was received.
// SetIpv6 sets the string value in the BgpAttributesNextHop object
func (obj *bgpAttributesNextHop) SetIpv6(value string) BgpAttributesNextHop {
	obj.setChoice(BgpAttributesNextHopChoice.IPV6)
	obj.obj.Ipv6 = &value
	return obj
}

// description is TBD
// Ipv6TwoAddresses returns a BgpAttributesNextHopIpv6TwoAddresses
func (obj *bgpAttributesNextHop) Ipv6TwoAddresses() BgpAttributesNextHopIpv6TwoAddresses {
	if obj.obj.Ipv6TwoAddresses == nil {
		obj.setChoice(BgpAttributesNextHopChoice.IPV6_TWO_ADDRESSES)
	}
	if obj.ipv6TwoAddressesHolder == nil {
		obj.ipv6TwoAddressesHolder = &bgpAttributesNextHopIpv6TwoAddresses{obj: obj.obj.Ipv6TwoAddresses}
	}
	return obj.ipv6TwoAddressesHolder
}

// description is TBD
// Ipv6TwoAddresses returns a BgpAttributesNextHopIpv6TwoAddresses
func (obj *bgpAttributesNextHop) HasIpv6TwoAddresses() bool {
	return obj.obj.Ipv6TwoAddresses != nil
}

// description is TBD
// SetIpv6TwoAddresses sets the BgpAttributesNextHopIpv6TwoAddresses value in the BgpAttributesNextHop object
func (obj *bgpAttributesNextHop) SetIpv6TwoAddresses(value BgpAttributesNextHopIpv6TwoAddresses) BgpAttributesNextHop {
	obj.setChoice(BgpAttributesNextHopChoice.IPV6_TWO_ADDRESSES)
	obj.ipv6TwoAddressesHolder = nil
	obj.obj.Ipv6TwoAddresses = value.msg()

	return obj
}

func (obj *bgpAttributesNextHop) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface BgpAttributesNextHop")
	}

	if obj.obj.Ipv4 != nil {

		err := obj.validateIpv4(obj.Ipv4())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpAttributesNextHop.Ipv4"))
		}

	}

	if obj.obj.Ipv6 != nil {

		err := obj.validateIpv6(obj.Ipv6())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpAttributesNextHop.Ipv6"))
		}

	}

	if obj.obj.Ipv6TwoAddresses != nil {

		obj.Ipv6TwoAddresses().validateObj(vObj, set_default)
	}

}

func (obj *bgpAttributesNextHop) setDefault() {
	var choices_set int = 0
	var choice BgpAttributesNextHopChoiceEnum

	if obj.obj.Ipv4 != nil {
		choices_set += 1
		choice = BgpAttributesNextHopChoice.IPV4
	}

	if obj.obj.Ipv6 != nil {
		choices_set += 1
		choice = BgpAttributesNextHopChoice.IPV6
	}

	if obj.obj.Ipv6TwoAddresses != nil {
		choices_set += 1
		choice = BgpAttributesNextHopChoice.IPV6_TWO_ADDRESSES
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in BgpAttributesNextHop")
			}
		} else {
			intVal := otg.BgpAttributesNextHop_Choice_Enum_value[string(choice)]
			enumValue := otg.BgpAttributesNextHop_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.Ipv4 == nil && choice == BgpAttributesNextHopChoice.IPV4 {
		obj.SetIpv4("0.0.0.0")
	}
	if obj.obj.Ipv6 == nil && choice == BgpAttributesNextHopChoice.IPV6 {
		obj.SetIpv6("0::0")
	}

}
