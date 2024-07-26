package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesNextHopIpv6TwoAddresses *****
type bgpAttributesNextHopIpv6TwoAddresses struct {
	validation
	obj          *otg.BgpAttributesNextHopIpv6TwoAddresses
	marshaller   marshalBgpAttributesNextHopIpv6TwoAddresses
	unMarshaller unMarshalBgpAttributesNextHopIpv6TwoAddresses
}

func NewBgpAttributesNextHopIpv6TwoAddresses() BgpAttributesNextHopIpv6TwoAddresses {
	obj := bgpAttributesNextHopIpv6TwoAddresses{obj: &otg.BgpAttributesNextHopIpv6TwoAddresses{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesNextHopIpv6TwoAddresses) msg() *otg.BgpAttributesNextHopIpv6TwoAddresses {
	return obj.obj
}

func (obj *bgpAttributesNextHopIpv6TwoAddresses) setMsg(msg *otg.BgpAttributesNextHopIpv6TwoAddresses) BgpAttributesNextHopIpv6TwoAddresses {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesNextHopIpv6TwoAddresses struct {
	obj *bgpAttributesNextHopIpv6TwoAddresses
}

type marshalBgpAttributesNextHopIpv6TwoAddresses interface {
	// ToProto marshals BgpAttributesNextHopIpv6TwoAddresses to protobuf object *otg.BgpAttributesNextHopIpv6TwoAddresses
	ToProto() (*otg.BgpAttributesNextHopIpv6TwoAddresses, error)
	// ToPbText marshals BgpAttributesNextHopIpv6TwoAddresses to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesNextHopIpv6TwoAddresses to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesNextHopIpv6TwoAddresses to JSON text
	ToJson() (string, error)
}

type unMarshalbgpAttributesNextHopIpv6TwoAddresses struct {
	obj *bgpAttributesNextHopIpv6TwoAddresses
}

type unMarshalBgpAttributesNextHopIpv6TwoAddresses interface {
	// FromProto unmarshals BgpAttributesNextHopIpv6TwoAddresses from protobuf object *otg.BgpAttributesNextHopIpv6TwoAddresses
	FromProto(msg *otg.BgpAttributesNextHopIpv6TwoAddresses) (BgpAttributesNextHopIpv6TwoAddresses, error)
	// FromPbText unmarshals BgpAttributesNextHopIpv6TwoAddresses from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesNextHopIpv6TwoAddresses from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesNextHopIpv6TwoAddresses from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesNextHopIpv6TwoAddresses) Marshal() marshalBgpAttributesNextHopIpv6TwoAddresses {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesNextHopIpv6TwoAddresses{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesNextHopIpv6TwoAddresses) Unmarshal() unMarshalBgpAttributesNextHopIpv6TwoAddresses {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesNextHopIpv6TwoAddresses{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesNextHopIpv6TwoAddresses) ToProto() (*otg.BgpAttributesNextHopIpv6TwoAddresses, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesNextHopIpv6TwoAddresses) FromProto(msg *otg.BgpAttributesNextHopIpv6TwoAddresses) (BgpAttributesNextHopIpv6TwoAddresses, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesNextHopIpv6TwoAddresses) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesNextHopIpv6TwoAddresses) FromPbText(value string) error {
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

func (m *marshalbgpAttributesNextHopIpv6TwoAddresses) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesNextHopIpv6TwoAddresses) FromYaml(value string) error {
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

func (m *marshalbgpAttributesNextHopIpv6TwoAddresses) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesNextHopIpv6TwoAddresses) FromJson(value string) error {
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

func (obj *bgpAttributesNextHopIpv6TwoAddresses) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesNextHopIpv6TwoAddresses) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesNextHopIpv6TwoAddresses) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesNextHopIpv6TwoAddresses) Clone() (BgpAttributesNextHopIpv6TwoAddresses, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesNextHopIpv6TwoAddresses()
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

// BgpAttributesNextHopIpv6TwoAddresses is there is a specific scenario in which it is possible to receive a Global and Link Local address in the Next Hop
// field in a MP_REACH attribute or in the NEXT_HOP attribute(RFC2545: Section 3).
type BgpAttributesNextHopIpv6TwoAddresses interface {
	Validation
	// msg marshals BgpAttributesNextHopIpv6TwoAddresses to protobuf object *otg.BgpAttributesNextHopIpv6TwoAddresses
	// and doesn't set defaults
	msg() *otg.BgpAttributesNextHopIpv6TwoAddresses
	// setMsg unmarshals BgpAttributesNextHopIpv6TwoAddresses from protobuf object *otg.BgpAttributesNextHopIpv6TwoAddresses
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesNextHopIpv6TwoAddresses) BgpAttributesNextHopIpv6TwoAddresses
	// provides marshal interface
	Marshal() marshalBgpAttributesNextHopIpv6TwoAddresses
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesNextHopIpv6TwoAddresses
	// validate validates BgpAttributesNextHopIpv6TwoAddresses
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesNextHopIpv6TwoAddresses, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// First returns string, set in BgpAttributesNextHopIpv6TwoAddresses.
	First() string
	// SetFirst assigns string provided by user to BgpAttributesNextHopIpv6TwoAddresses
	SetFirst(value string) BgpAttributesNextHopIpv6TwoAddresses
	// HasFirst checks if First has been set in BgpAttributesNextHopIpv6TwoAddresses
	HasFirst() bool
	// Second returns string, set in BgpAttributesNextHopIpv6TwoAddresses.
	Second() string
	// SetSecond assigns string provided by user to BgpAttributesNextHopIpv6TwoAddresses
	SetSecond(value string) BgpAttributesNextHopIpv6TwoAddresses
	// HasSecond checks if Second has been set in BgpAttributesNextHopIpv6TwoAddresses
	HasSecond() bool
}

// The first IPv6 next hop in the 32 byte IPv6 Next Hop.
// First returns a string
func (obj *bgpAttributesNextHopIpv6TwoAddresses) First() string {

	return *obj.obj.First

}

// The first IPv6 next hop in the 32 byte IPv6 Next Hop.
// First returns a string
func (obj *bgpAttributesNextHopIpv6TwoAddresses) HasFirst() bool {
	return obj.obj.First != nil
}

// The first IPv6 next hop in the 32 byte IPv6 Next Hop.
// SetFirst sets the string value in the BgpAttributesNextHopIpv6TwoAddresses object
func (obj *bgpAttributesNextHopIpv6TwoAddresses) SetFirst(value string) BgpAttributesNextHopIpv6TwoAddresses {

	obj.obj.First = &value
	return obj
}

// The second IPv6 next hop in the 32 byte IPv6 Next Hop.
// Second returns a string
func (obj *bgpAttributesNextHopIpv6TwoAddresses) Second() string {

	return *obj.obj.Second

}

// The second IPv6 next hop in the 32 byte IPv6 Next Hop.
// Second returns a string
func (obj *bgpAttributesNextHopIpv6TwoAddresses) HasSecond() bool {
	return obj.obj.Second != nil
}

// The second IPv6 next hop in the 32 byte IPv6 Next Hop.
// SetSecond sets the string value in the BgpAttributesNextHopIpv6TwoAddresses object
func (obj *bgpAttributesNextHopIpv6TwoAddresses) SetSecond(value string) BgpAttributesNextHopIpv6TwoAddresses {

	obj.obj.Second = &value
	return obj
}

func (obj *bgpAttributesNextHopIpv6TwoAddresses) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.First != nil {

		err := obj.validateIpv6(obj.First())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpAttributesNextHopIpv6TwoAddresses.First"))
		}

	}

	if obj.obj.Second != nil {

		err := obj.validateIpv6(obj.Second())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpAttributesNextHopIpv6TwoAddresses.Second"))
		}

	}

}

func (obj *bgpAttributesNextHopIpv6TwoAddresses) setDefault() {
	if obj.obj.First == nil {
		obj.SetFirst("0::0")
	}
	if obj.obj.Second == nil {
		obj.SetSecond("0::0")
	}

}
