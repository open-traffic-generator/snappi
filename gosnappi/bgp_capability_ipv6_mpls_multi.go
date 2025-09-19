package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpCapabilityIpv6MplsMulti *****
type bgpCapabilityIpv6MplsMulti struct {
	validation
	obj          *otg.BgpCapabilityIpv6MplsMulti
	marshaller   marshalBgpCapabilityIpv6MplsMulti
	unMarshaller unMarshalBgpCapabilityIpv6MplsMulti
}

func NewBgpCapabilityIpv6MplsMulti() BgpCapabilityIpv6MplsMulti {
	obj := bgpCapabilityIpv6MplsMulti{obj: &otg.BgpCapabilityIpv6MplsMulti{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpCapabilityIpv6MplsMulti) msg() *otg.BgpCapabilityIpv6MplsMulti {
	return obj.obj
}

func (obj *bgpCapabilityIpv6MplsMulti) setMsg(msg *otg.BgpCapabilityIpv6MplsMulti) BgpCapabilityIpv6MplsMulti {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpCapabilityIpv6MplsMulti struct {
	obj *bgpCapabilityIpv6MplsMulti
}

type marshalBgpCapabilityIpv6MplsMulti interface {
	// ToProto marshals BgpCapabilityIpv6MplsMulti to protobuf object *otg.BgpCapabilityIpv6MplsMulti
	ToProto() (*otg.BgpCapabilityIpv6MplsMulti, error)
	// ToPbText marshals BgpCapabilityIpv6MplsMulti to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpCapabilityIpv6MplsMulti to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpCapabilityIpv6MplsMulti to JSON text
	ToJson() (string, error)
}

type unMarshalbgpCapabilityIpv6MplsMulti struct {
	obj *bgpCapabilityIpv6MplsMulti
}

type unMarshalBgpCapabilityIpv6MplsMulti interface {
	// FromProto unmarshals BgpCapabilityIpv6MplsMulti from protobuf object *otg.BgpCapabilityIpv6MplsMulti
	FromProto(msg *otg.BgpCapabilityIpv6MplsMulti) (BgpCapabilityIpv6MplsMulti, error)
	// FromPbText unmarshals BgpCapabilityIpv6MplsMulti from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpCapabilityIpv6MplsMulti from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpCapabilityIpv6MplsMulti from JSON text
	FromJson(value string) error
}

func (obj *bgpCapabilityIpv6MplsMulti) Marshal() marshalBgpCapabilityIpv6MplsMulti {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpCapabilityIpv6MplsMulti{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpCapabilityIpv6MplsMulti) Unmarshal() unMarshalBgpCapabilityIpv6MplsMulti {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpCapabilityIpv6MplsMulti{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpCapabilityIpv6MplsMulti) ToProto() (*otg.BgpCapabilityIpv6MplsMulti, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpCapabilityIpv6MplsMulti) FromProto(msg *otg.BgpCapabilityIpv6MplsMulti) (BgpCapabilityIpv6MplsMulti, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpCapabilityIpv6MplsMulti) ToPbText() (string, error) {
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

func (m *unMarshalbgpCapabilityIpv6MplsMulti) FromPbText(value string) error {
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

func (m *marshalbgpCapabilityIpv6MplsMulti) ToYaml() (string, error) {
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

func (m *unMarshalbgpCapabilityIpv6MplsMulti) FromYaml(value string) error {
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

func (m *marshalbgpCapabilityIpv6MplsMulti) ToJson() (string, error) {
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

func (m *unMarshalbgpCapabilityIpv6MplsMulti) FromJson(value string) error {
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

func (obj *bgpCapabilityIpv6MplsMulti) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpCapabilityIpv6MplsMulti) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpCapabilityIpv6MplsMulti) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpCapabilityIpv6MplsMulti) Clone() (BgpCapabilityIpv6MplsMulti, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpCapabilityIpv6MplsMulti()
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

// BgpCapabilityIpv6MplsMulti is the MPLS multiple Lable capability is advertised in Optional Capability under Multiple Label Capability (code 8). Reference: https://datatracker.ietf.org/doc/html/rfc8277#section-2.1. For Multiple Labels binding to the IPv6 address prefix, BGP speaker will advertise triple of the form <AFI=2, SAFI=4, Count=n> in BGP Open Message under the Multiprotocol  extensions capability, where n = ipv6_mpls_multi, the numbers of labels.
type BgpCapabilityIpv6MplsMulti interface {
	Validation
	// msg marshals BgpCapabilityIpv6MplsMulti to protobuf object *otg.BgpCapabilityIpv6MplsMulti
	// and doesn't set defaults
	msg() *otg.BgpCapabilityIpv6MplsMulti
	// setMsg unmarshals BgpCapabilityIpv6MplsMulti from protobuf object *otg.BgpCapabilityIpv6MplsMulti
	// and doesn't set defaults
	setMsg(*otg.BgpCapabilityIpv6MplsMulti) BgpCapabilityIpv6MplsMulti
	// provides marshal interface
	Marshal() marshalBgpCapabilityIpv6MplsMulti
	// provides unmarshal interface
	Unmarshal() unMarshalBgpCapabilityIpv6MplsMulti
	// validate validates BgpCapabilityIpv6MplsMulti
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpCapabilityIpv6MplsMulti, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LabelCount returns uint32, set in BgpCapabilityIpv6MplsMulti.
	LabelCount() uint32
	// SetLabelCount assigns uint32 provided by user to BgpCapabilityIpv6MplsMulti
	SetLabelCount(value uint32) BgpCapabilityIpv6MplsMulti
	// HasLabelCount checks if LabelCount has been set in BgpCapabilityIpv6MplsMulti
	HasLabelCount() bool
}

// Number of Labels (n) in the triplet format <AFI=2, SAFI=4, Count=n>.
// LabelCount returns a uint32
func (obj *bgpCapabilityIpv6MplsMulti) LabelCount() uint32 {

	return *obj.obj.LabelCount

}

// Number of Labels (n) in the triplet format <AFI=2, SAFI=4, Count=n>.
// LabelCount returns a uint32
func (obj *bgpCapabilityIpv6MplsMulti) HasLabelCount() bool {
	return obj.obj.LabelCount != nil
}

// Number of Labels (n) in the triplet format <AFI=2, SAFI=4, Count=n>.
// SetLabelCount sets the uint32 value in the BgpCapabilityIpv6MplsMulti object
func (obj *bgpCapabilityIpv6MplsMulti) SetLabelCount(value uint32) BgpCapabilityIpv6MplsMulti {

	obj.obj.LabelCount = &value
	return obj
}

func (obj *bgpCapabilityIpv6MplsMulti) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.LabelCount != nil {

		if *obj.obj.LabelCount < 2 || *obj.obj.LabelCount > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("2 <= BgpCapabilityIpv6MplsMulti.LabelCount <= 255 but Got %d", *obj.obj.LabelCount))
		}

	}

}

func (obj *bgpCapabilityIpv6MplsMulti) setDefault() {
	if obj.obj.LabelCount == nil {
		obj.SetLabelCount(2)
	}

}
