package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesAggregator *****
type bgpAttributesAggregator struct {
	validation
	obj          *otg.BgpAttributesAggregator
	marshaller   marshalBgpAttributesAggregator
	unMarshaller unMarshalBgpAttributesAggregator
}

func NewBgpAttributesAggregator() BgpAttributesAggregator {
	obj := bgpAttributesAggregator{obj: &otg.BgpAttributesAggregator{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesAggregator) msg() *otg.BgpAttributesAggregator {
	return obj.obj
}

func (obj *bgpAttributesAggregator) setMsg(msg *otg.BgpAttributesAggregator) BgpAttributesAggregator {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesAggregator struct {
	obj *bgpAttributesAggregator
}

type marshalBgpAttributesAggregator interface {
	// ToProto marshals BgpAttributesAggregator to protobuf object *otg.BgpAttributesAggregator
	ToProto() (*otg.BgpAttributesAggregator, error)
	// ToPbText marshals BgpAttributesAggregator to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesAggregator to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesAggregator to JSON text
	ToJson() (string, error)
}

type unMarshalbgpAttributesAggregator struct {
	obj *bgpAttributesAggregator
}

type unMarshalBgpAttributesAggregator interface {
	// FromProto unmarshals BgpAttributesAggregator from protobuf object *otg.BgpAttributesAggregator
	FromProto(msg *otg.BgpAttributesAggregator) (BgpAttributesAggregator, error)
	// FromPbText unmarshals BgpAttributesAggregator from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesAggregator from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesAggregator from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesAggregator) Marshal() marshalBgpAttributesAggregator {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesAggregator{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesAggregator) Unmarshal() unMarshalBgpAttributesAggregator {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesAggregator{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesAggregator) ToProto() (*otg.BgpAttributesAggregator, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesAggregator) FromProto(msg *otg.BgpAttributesAggregator) (BgpAttributesAggregator, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesAggregator) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesAggregator) FromPbText(value string) error {
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

func (m *marshalbgpAttributesAggregator) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesAggregator) FromYaml(value string) error {
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

func (m *marshalbgpAttributesAggregator) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesAggregator) FromJson(value string) error {
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

func (obj *bgpAttributesAggregator) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesAggregator) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesAggregator) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesAggregator) Clone() (BgpAttributesAggregator, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesAggregator()
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

// BgpAttributesAggregator is optional AGGREGATOR attribute which maybe be added by a BGP speaker which performs route aggregation.
// When AGGREGATOR attribute is being sent to a new BGP speaker , the AS number is encoded as a 4 byte value.
// When AGGREGATOR attribute is being exchanged between a new and an old BGP speaker or between two old BGP speakers,
// the AS number is encoded as a 2 byte value.
// It contain the AS number and IP address of the speaker performing the aggregation.
type BgpAttributesAggregator interface {
	Validation
	// msg marshals BgpAttributesAggregator to protobuf object *otg.BgpAttributesAggregator
	// and doesn't set defaults
	msg() *otg.BgpAttributesAggregator
	// setMsg unmarshals BgpAttributesAggregator from protobuf object *otg.BgpAttributesAggregator
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesAggregator) BgpAttributesAggregator
	// provides marshal interface
	Marshal() marshalBgpAttributesAggregator
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesAggregator
	// validate validates BgpAttributesAggregator
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesAggregator, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns BgpAttributesAggregatorChoiceEnum, set in BgpAttributesAggregator
	Choice() BgpAttributesAggregatorChoiceEnum
	// setChoice assigns BgpAttributesAggregatorChoiceEnum provided by user to BgpAttributesAggregator
	setChoice(value BgpAttributesAggregatorChoiceEnum) BgpAttributesAggregator
	// HasChoice checks if Choice has been set in BgpAttributesAggregator
	HasChoice() bool
	// FourByteAs returns uint32, set in BgpAttributesAggregator.
	FourByteAs() uint32
	// SetFourByteAs assigns uint32 provided by user to BgpAttributesAggregator
	SetFourByteAs(value uint32) BgpAttributesAggregator
	// HasFourByteAs checks if FourByteAs has been set in BgpAttributesAggregator
	HasFourByteAs() bool
	// TwoByteAs returns uint32, set in BgpAttributesAggregator.
	TwoByteAs() uint32
	// SetTwoByteAs assigns uint32 provided by user to BgpAttributesAggregator
	SetTwoByteAs(value uint32) BgpAttributesAggregator
	// HasTwoByteAs checks if TwoByteAs has been set in BgpAttributesAggregator
	HasTwoByteAs() bool
	// Ipv4Address returns string, set in BgpAttributesAggregator.
	Ipv4Address() string
	// SetIpv4Address assigns string provided by user to BgpAttributesAggregator
	SetIpv4Address(value string) BgpAttributesAggregator
	// HasIpv4Address checks if Ipv4Address has been set in BgpAttributesAggregator
	HasIpv4Address() bool
}

type BgpAttributesAggregatorChoiceEnum string

// Enum of Choice on BgpAttributesAggregator
var BgpAttributesAggregatorChoice = struct {
	FOUR_BYTE_AS BgpAttributesAggregatorChoiceEnum
	TWO_BYTE_AS  BgpAttributesAggregatorChoiceEnum
}{
	FOUR_BYTE_AS: BgpAttributesAggregatorChoiceEnum("four_byte_as"),
	TWO_BYTE_AS:  BgpAttributesAggregatorChoiceEnum("two_byte_as"),
}

func (obj *bgpAttributesAggregator) Choice() BgpAttributesAggregatorChoiceEnum {
	return BgpAttributesAggregatorChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *bgpAttributesAggregator) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *bgpAttributesAggregator) setChoice(value BgpAttributesAggregatorChoiceEnum) BgpAttributesAggregator {
	intValue, ok := otg.BgpAttributesAggregator_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpAttributesAggregatorChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpAttributesAggregator_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.TwoByteAs = nil
	obj.obj.FourByteAs = nil

	if value == BgpAttributesAggregatorChoice.FOUR_BYTE_AS {
		defaultValue := uint32(65536)
		obj.obj.FourByteAs = &defaultValue
	}

	if value == BgpAttributesAggregatorChoice.TWO_BYTE_AS {
		defaultValue := uint32(1)
		obj.obj.TwoByteAs = &defaultValue
	}

	return obj
}

// The value of the 4 byte  AS number of the BGP speaker which aggregated the route. If the value of the AS number is less than 2 octets ( 65535 or less), the AS4_AGGREGATOR object should not be sent.
// FourByteAs returns a uint32
func (obj *bgpAttributesAggregator) FourByteAs() uint32 {

	if obj.obj.FourByteAs == nil {
		obj.setChoice(BgpAttributesAggregatorChoice.FOUR_BYTE_AS)
	}

	return *obj.obj.FourByteAs

}

// The value of the 4 byte  AS number of the BGP speaker which aggregated the route. If the value of the AS number is less than 2 octets ( 65535 or less), the AS4_AGGREGATOR object should not be sent.
// FourByteAs returns a uint32
func (obj *bgpAttributesAggregator) HasFourByteAs() bool {
	return obj.obj.FourByteAs != nil
}

// The value of the 4 byte  AS number of the BGP speaker which aggregated the route. If the value of the AS number is less than 2 octets ( 65535 or less), the AS4_AGGREGATOR object should not be sent.
// SetFourByteAs sets the uint32 value in the BgpAttributesAggregator object
func (obj *bgpAttributesAggregator) SetFourByteAs(value uint32) BgpAttributesAggregator {
	obj.setChoice(BgpAttributesAggregatorChoice.FOUR_BYTE_AS)
	obj.obj.FourByteAs = &value
	return obj
}

// The value of the 2 byte AS number of the BGP speaker which aggregated the route.
// TwoByteAs returns a uint32
func (obj *bgpAttributesAggregator) TwoByteAs() uint32 {

	if obj.obj.TwoByteAs == nil {
		obj.setChoice(BgpAttributesAggregatorChoice.TWO_BYTE_AS)
	}

	return *obj.obj.TwoByteAs

}

// The value of the 2 byte AS number of the BGP speaker which aggregated the route.
// TwoByteAs returns a uint32
func (obj *bgpAttributesAggregator) HasTwoByteAs() bool {
	return obj.obj.TwoByteAs != nil
}

// The value of the 2 byte AS number of the BGP speaker which aggregated the route.
// SetTwoByteAs sets the uint32 value in the BgpAttributesAggregator object
func (obj *bgpAttributesAggregator) SetTwoByteAs(value uint32) BgpAttributesAggregator {
	obj.setChoice(BgpAttributesAggregatorChoice.TWO_BYTE_AS)
	obj.obj.TwoByteAs = &value
	return obj
}

// The IPv4 address of the BGP speaker which aggregated the route.
// Ipv4Address returns a string
func (obj *bgpAttributesAggregator) Ipv4Address() string {

	return *obj.obj.Ipv4Address

}

// The IPv4 address of the BGP speaker which aggregated the route.
// Ipv4Address returns a string
func (obj *bgpAttributesAggregator) HasIpv4Address() bool {
	return obj.obj.Ipv4Address != nil
}

// The IPv4 address of the BGP speaker which aggregated the route.
// SetIpv4Address sets the string value in the BgpAttributesAggregator object
func (obj *bgpAttributesAggregator) SetIpv4Address(value string) BgpAttributesAggregator {

	obj.obj.Ipv4Address = &value
	return obj
}

func (obj *bgpAttributesAggregator) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.TwoByteAs != nil {

		if *obj.obj.TwoByteAs > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpAttributesAggregator.TwoByteAs <= 65535 but Got %d", *obj.obj.TwoByteAs))
		}

	}

	if obj.obj.Ipv4Address != nil {

		err := obj.validateIpv4(obj.Ipv4Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpAttributesAggregator.Ipv4Address"))
		}

	}

}

func (obj *bgpAttributesAggregator) setDefault() {
	var choices_set int = 0
	var choice BgpAttributesAggregatorChoiceEnum

	if obj.obj.FourByteAs != nil {
		choices_set += 1
		choice = BgpAttributesAggregatorChoice.FOUR_BYTE_AS
	}

	if obj.obj.TwoByteAs != nil {
		choices_set += 1
		choice = BgpAttributesAggregatorChoice.TWO_BYTE_AS
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(BgpAttributesAggregatorChoice.FOUR_BYTE_AS)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in BgpAttributesAggregator")
			}
		} else {
			intVal := otg.BgpAttributesAggregator_Choice_Enum_value[string(choice)]
			enumValue := otg.BgpAttributesAggregator_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.Ipv4Address == nil {
		obj.SetIpv4Address("0.0.0.0")
	}

}
