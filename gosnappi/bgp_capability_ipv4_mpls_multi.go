package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpCapabilityIpv4MplsMulti *****
type bgpCapabilityIpv4MplsMulti struct {
	validation
	obj          *otg.BgpCapabilityIpv4MplsMulti
	marshaller   marshalBgpCapabilityIpv4MplsMulti
	unMarshaller unMarshalBgpCapabilityIpv4MplsMulti
}

func NewBgpCapabilityIpv4MplsMulti() BgpCapabilityIpv4MplsMulti {
	obj := bgpCapabilityIpv4MplsMulti{obj: &otg.BgpCapabilityIpv4MplsMulti{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpCapabilityIpv4MplsMulti) msg() *otg.BgpCapabilityIpv4MplsMulti {
	return obj.obj
}

func (obj *bgpCapabilityIpv4MplsMulti) setMsg(msg *otg.BgpCapabilityIpv4MplsMulti) BgpCapabilityIpv4MplsMulti {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpCapabilityIpv4MplsMulti struct {
	obj *bgpCapabilityIpv4MplsMulti
}

type marshalBgpCapabilityIpv4MplsMulti interface {
	// ToProto marshals BgpCapabilityIpv4MplsMulti to protobuf object *otg.BgpCapabilityIpv4MplsMulti
	ToProto() (*otg.BgpCapabilityIpv4MplsMulti, error)
	// ToPbText marshals BgpCapabilityIpv4MplsMulti to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpCapabilityIpv4MplsMulti to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpCapabilityIpv4MplsMulti to JSON text
	ToJson() (string, error)
}

type unMarshalbgpCapabilityIpv4MplsMulti struct {
	obj *bgpCapabilityIpv4MplsMulti
}

type unMarshalBgpCapabilityIpv4MplsMulti interface {
	// FromProto unmarshals BgpCapabilityIpv4MplsMulti from protobuf object *otg.BgpCapabilityIpv4MplsMulti
	FromProto(msg *otg.BgpCapabilityIpv4MplsMulti) (BgpCapabilityIpv4MplsMulti, error)
	// FromPbText unmarshals BgpCapabilityIpv4MplsMulti from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpCapabilityIpv4MplsMulti from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpCapabilityIpv4MplsMulti from JSON text
	FromJson(value string) error
}

func (obj *bgpCapabilityIpv4MplsMulti) Marshal() marshalBgpCapabilityIpv4MplsMulti {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpCapabilityIpv4MplsMulti{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpCapabilityIpv4MplsMulti) Unmarshal() unMarshalBgpCapabilityIpv4MplsMulti {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpCapabilityIpv4MplsMulti{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpCapabilityIpv4MplsMulti) ToProto() (*otg.BgpCapabilityIpv4MplsMulti, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpCapabilityIpv4MplsMulti) FromProto(msg *otg.BgpCapabilityIpv4MplsMulti) (BgpCapabilityIpv4MplsMulti, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpCapabilityIpv4MplsMulti) ToPbText() (string, error) {
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

func (m *unMarshalbgpCapabilityIpv4MplsMulti) FromPbText(value string) error {
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

func (m *marshalbgpCapabilityIpv4MplsMulti) ToYaml() (string, error) {
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

func (m *unMarshalbgpCapabilityIpv4MplsMulti) FromYaml(value string) error {
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

func (m *marshalbgpCapabilityIpv4MplsMulti) ToJson() (string, error) {
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

func (m *unMarshalbgpCapabilityIpv4MplsMulti) FromJson(value string) error {
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

func (obj *bgpCapabilityIpv4MplsMulti) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpCapabilityIpv4MplsMulti) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpCapabilityIpv4MplsMulti) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpCapabilityIpv4MplsMulti) Clone() (BgpCapabilityIpv4MplsMulti, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpCapabilityIpv4MplsMulti()
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

// BgpCapabilityIpv4MplsMulti is the MPLS multiple Lable capability is advertised in Optional Capability under Multiple Label Capability (code 8). Reference: https://datatracker.ietf.org/doc/html/rfc8277#section-2.1. For Multiple Labels binding to the IPv6 address prefix, BGP speaker will advertise triple of the form <AFI=2, SAFI=4, Count=n> in BGP Open Message under the Multiprotocol  extensions capability, where n = ipv6_mpls_multi, the numbers of labels.
type BgpCapabilityIpv4MplsMulti interface {
	Validation
	// msg marshals BgpCapabilityIpv4MplsMulti to protobuf object *otg.BgpCapabilityIpv4MplsMulti
	// and doesn't set defaults
	msg() *otg.BgpCapabilityIpv4MplsMulti
	// setMsg unmarshals BgpCapabilityIpv4MplsMulti from protobuf object *otg.BgpCapabilityIpv4MplsMulti
	// and doesn't set defaults
	setMsg(*otg.BgpCapabilityIpv4MplsMulti) BgpCapabilityIpv4MplsMulti
	// provides marshal interface
	Marshal() marshalBgpCapabilityIpv4MplsMulti
	// provides unmarshal interface
	Unmarshal() unMarshalBgpCapabilityIpv4MplsMulti
	// validate validates BgpCapabilityIpv4MplsMulti
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpCapabilityIpv4MplsMulti, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LabelCount returns uint32, set in BgpCapabilityIpv4MplsMulti.
	LabelCount() uint32
	// SetLabelCount assigns uint32 provided by user to BgpCapabilityIpv4MplsMulti
	SetLabelCount(value uint32) BgpCapabilityIpv4MplsMulti
	// HasLabelCount checks if LabelCount has been set in BgpCapabilityIpv4MplsMulti
	HasLabelCount() bool
}

// Number of Labels (n) in the triplet format <AFI=2, SAFI=4, Count=n>.
// LabelCount returns a uint32
func (obj *bgpCapabilityIpv4MplsMulti) LabelCount() uint32 {

	return *obj.obj.LabelCount

}

// Number of Labels (n) in the triplet format <AFI=2, SAFI=4, Count=n>.
// LabelCount returns a uint32
func (obj *bgpCapabilityIpv4MplsMulti) HasLabelCount() bool {
	return obj.obj.LabelCount != nil
}

// Number of Labels (n) in the triplet format <AFI=2, SAFI=4, Count=n>.
// SetLabelCount sets the uint32 value in the BgpCapabilityIpv4MplsMulti object
func (obj *bgpCapabilityIpv4MplsMulti) SetLabelCount(value uint32) BgpCapabilityIpv4MplsMulti {

	obj.obj.LabelCount = &value
	return obj
}

func (obj *bgpCapabilityIpv4MplsMulti) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.LabelCount != nil {

		if *obj.obj.LabelCount < 2 || *obj.obj.LabelCount > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("2 <= BgpCapabilityIpv4MplsMulti.LabelCount <= 255 but Got %d", *obj.obj.LabelCount))
		}

	}

}

func (obj *bgpCapabilityIpv4MplsMulti) setDefault() {
	if obj.obj.LabelCount == nil {
		obj.SetLabelCount(2)
	}

}
