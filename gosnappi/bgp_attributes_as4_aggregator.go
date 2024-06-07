package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesAs4Aggregator *****
type bgpAttributesAs4Aggregator struct {
	validation
	obj          *otg.BgpAttributesAs4Aggregator
	marshaller   marshalBgpAttributesAs4Aggregator
	unMarshaller unMarshalBgpAttributesAs4Aggregator
}

func NewBgpAttributesAs4Aggregator() BgpAttributesAs4Aggregator {
	obj := bgpAttributesAs4Aggregator{obj: &otg.BgpAttributesAs4Aggregator{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesAs4Aggregator) msg() *otg.BgpAttributesAs4Aggregator {
	return obj.obj
}

func (obj *bgpAttributesAs4Aggregator) setMsg(msg *otg.BgpAttributesAs4Aggregator) BgpAttributesAs4Aggregator {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesAs4Aggregator struct {
	obj *bgpAttributesAs4Aggregator
}

type marshalBgpAttributesAs4Aggregator interface {
	// ToProto marshals BgpAttributesAs4Aggregator to protobuf object *otg.BgpAttributesAs4Aggregator
	ToProto() (*otg.BgpAttributesAs4Aggregator, error)
	// ToPbText marshals BgpAttributesAs4Aggregator to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesAs4Aggregator to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesAs4Aggregator to JSON text
	ToJson() (string, error)
}

type unMarshalbgpAttributesAs4Aggregator struct {
	obj *bgpAttributesAs4Aggregator
}

type unMarshalBgpAttributesAs4Aggregator interface {
	// FromProto unmarshals BgpAttributesAs4Aggregator from protobuf object *otg.BgpAttributesAs4Aggregator
	FromProto(msg *otg.BgpAttributesAs4Aggregator) (BgpAttributesAs4Aggregator, error)
	// FromPbText unmarshals BgpAttributesAs4Aggregator from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesAs4Aggregator from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesAs4Aggregator from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesAs4Aggregator) Marshal() marshalBgpAttributesAs4Aggregator {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesAs4Aggregator{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesAs4Aggregator) Unmarshal() unMarshalBgpAttributesAs4Aggregator {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesAs4Aggregator{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesAs4Aggregator) ToProto() (*otg.BgpAttributesAs4Aggregator, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesAs4Aggregator) FromProto(msg *otg.BgpAttributesAs4Aggregator) (BgpAttributesAs4Aggregator, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesAs4Aggregator) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesAs4Aggregator) FromPbText(value string) error {
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

func (m *marshalbgpAttributesAs4Aggregator) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesAs4Aggregator) FromYaml(value string) error {
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

func (m *marshalbgpAttributesAs4Aggregator) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesAs4Aggregator) FromJson(value string) error {
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

func (obj *bgpAttributesAs4Aggregator) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesAs4Aggregator) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesAs4Aggregator) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesAs4Aggregator) Clone() (BgpAttributesAs4Aggregator, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesAs4Aggregator()
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

// BgpAttributesAs4Aggregator is optional AS4_AGGREGATOR attribute which maybe be added by a BGP speaker in one of two cases:
// - If it is a new BGP speaker speaking to an old BGP speaker and needs to send a 4 byte value for the AS number of the BGP route aggregator.
// - If it is a old BGP speaker speaking to a new BGP speaker and has to transparently forward a received AS4_AGGREGATOR from some other peer.
// Its usage is described in RFC4893.
type BgpAttributesAs4Aggregator interface {
	Validation
	// msg marshals BgpAttributesAs4Aggregator to protobuf object *otg.BgpAttributesAs4Aggregator
	// and doesn't set defaults
	msg() *otg.BgpAttributesAs4Aggregator
	// setMsg unmarshals BgpAttributesAs4Aggregator from protobuf object *otg.BgpAttributesAs4Aggregator
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesAs4Aggregator) BgpAttributesAs4Aggregator
	// provides marshal interface
	Marshal() marshalBgpAttributesAs4Aggregator
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesAs4Aggregator
	// validate validates BgpAttributesAs4Aggregator
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesAs4Aggregator, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// AsNum returns uint32, set in BgpAttributesAs4Aggregator.
	AsNum() uint32
	// SetAsNum assigns uint32 provided by user to BgpAttributesAs4Aggregator
	SetAsNum(value uint32) BgpAttributesAs4Aggregator
	// HasAsNum checks if AsNum has been set in BgpAttributesAs4Aggregator
	HasAsNum() bool
	// Ipv4Address returns string, set in BgpAttributesAs4Aggregator.
	Ipv4Address() string
	// SetIpv4Address assigns string provided by user to BgpAttributesAs4Aggregator
	SetIpv4Address(value string) BgpAttributesAs4Aggregator
	// HasIpv4Address checks if Ipv4Address has been set in BgpAttributesAs4Aggregator
	HasIpv4Address() bool
}

// The value of the 4 byte AS number of the BGP speaker which aggregated the route.
// AsNum returns a uint32
func (obj *bgpAttributesAs4Aggregator) AsNum() uint32 {

	return *obj.obj.AsNum

}

// The value of the 4 byte AS number of the BGP speaker which aggregated the route.
// AsNum returns a uint32
func (obj *bgpAttributesAs4Aggregator) HasAsNum() bool {
	return obj.obj.AsNum != nil
}

// The value of the 4 byte AS number of the BGP speaker which aggregated the route.
// SetAsNum sets the uint32 value in the BgpAttributesAs4Aggregator object
func (obj *bgpAttributesAs4Aggregator) SetAsNum(value uint32) BgpAttributesAs4Aggregator {

	obj.obj.AsNum = &value
	return obj
}

// The IPv4 address of the BGP speaker which aggregated the route.
// Ipv4Address returns a string
func (obj *bgpAttributesAs4Aggregator) Ipv4Address() string {

	return *obj.obj.Ipv4Address

}

// The IPv4 address of the BGP speaker which aggregated the route.
// Ipv4Address returns a string
func (obj *bgpAttributesAs4Aggregator) HasIpv4Address() bool {
	return obj.obj.Ipv4Address != nil
}

// The IPv4 address of the BGP speaker which aggregated the route.
// SetIpv4Address sets the string value in the BgpAttributesAs4Aggregator object
func (obj *bgpAttributesAs4Aggregator) SetIpv4Address(value string) BgpAttributesAs4Aggregator {

	obj.obj.Ipv4Address = &value
	return obj
}

func (obj *bgpAttributesAs4Aggregator) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv4Address != nil {

		err := obj.validateIpv4(obj.Ipv4Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpAttributesAs4Aggregator.Ipv4Address"))
		}

	}

}

func (obj *bgpAttributesAs4Aggregator) setDefault() {
	if obj.obj.Ipv4Address == nil {
		obj.SetIpv4Address("0.0.0.0")
	}

}
