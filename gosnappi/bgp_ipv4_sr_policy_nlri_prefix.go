package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpIpv4SrPolicyNLRIPrefix *****
type bgpIpv4SrPolicyNLRIPrefix struct {
	validation
	obj          *otg.BgpIpv4SrPolicyNLRIPrefix
	marshaller   marshalBgpIpv4SrPolicyNLRIPrefix
	unMarshaller unMarshalBgpIpv4SrPolicyNLRIPrefix
}

func NewBgpIpv4SrPolicyNLRIPrefix() BgpIpv4SrPolicyNLRIPrefix {
	obj := bgpIpv4SrPolicyNLRIPrefix{obj: &otg.BgpIpv4SrPolicyNLRIPrefix{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpIpv4SrPolicyNLRIPrefix) msg() *otg.BgpIpv4SrPolicyNLRIPrefix {
	return obj.obj
}

func (obj *bgpIpv4SrPolicyNLRIPrefix) setMsg(msg *otg.BgpIpv4SrPolicyNLRIPrefix) BgpIpv4SrPolicyNLRIPrefix {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpIpv4SrPolicyNLRIPrefix struct {
	obj *bgpIpv4SrPolicyNLRIPrefix
}

type marshalBgpIpv4SrPolicyNLRIPrefix interface {
	// ToProto marshals BgpIpv4SrPolicyNLRIPrefix to protobuf object *otg.BgpIpv4SrPolicyNLRIPrefix
	ToProto() (*otg.BgpIpv4SrPolicyNLRIPrefix, error)
	// ToPbText marshals BgpIpv4SrPolicyNLRIPrefix to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpIpv4SrPolicyNLRIPrefix to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpIpv4SrPolicyNLRIPrefix to JSON text
	ToJson() (string, error)
}

type unMarshalbgpIpv4SrPolicyNLRIPrefix struct {
	obj *bgpIpv4SrPolicyNLRIPrefix
}

type unMarshalBgpIpv4SrPolicyNLRIPrefix interface {
	// FromProto unmarshals BgpIpv4SrPolicyNLRIPrefix from protobuf object *otg.BgpIpv4SrPolicyNLRIPrefix
	FromProto(msg *otg.BgpIpv4SrPolicyNLRIPrefix) (BgpIpv4SrPolicyNLRIPrefix, error)
	// FromPbText unmarshals BgpIpv4SrPolicyNLRIPrefix from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpIpv4SrPolicyNLRIPrefix from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpIpv4SrPolicyNLRIPrefix from JSON text
	FromJson(value string) error
}

func (obj *bgpIpv4SrPolicyNLRIPrefix) Marshal() marshalBgpIpv4SrPolicyNLRIPrefix {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpIpv4SrPolicyNLRIPrefix{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpIpv4SrPolicyNLRIPrefix) Unmarshal() unMarshalBgpIpv4SrPolicyNLRIPrefix {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpIpv4SrPolicyNLRIPrefix{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpIpv4SrPolicyNLRIPrefix) ToProto() (*otg.BgpIpv4SrPolicyNLRIPrefix, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpIpv4SrPolicyNLRIPrefix) FromProto(msg *otg.BgpIpv4SrPolicyNLRIPrefix) (BgpIpv4SrPolicyNLRIPrefix, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpIpv4SrPolicyNLRIPrefix) ToPbText() (string, error) {
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

func (m *unMarshalbgpIpv4SrPolicyNLRIPrefix) FromPbText(value string) error {
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

func (m *marshalbgpIpv4SrPolicyNLRIPrefix) ToYaml() (string, error) {
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

func (m *unMarshalbgpIpv4SrPolicyNLRIPrefix) FromYaml(value string) error {
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

func (m *marshalbgpIpv4SrPolicyNLRIPrefix) ToJson() (string, error) {
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

func (m *unMarshalbgpIpv4SrPolicyNLRIPrefix) FromJson(value string) error {
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

func (obj *bgpIpv4SrPolicyNLRIPrefix) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpIpv4SrPolicyNLRIPrefix) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpIpv4SrPolicyNLRIPrefix) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpIpv4SrPolicyNLRIPrefix) Clone() (BgpIpv4SrPolicyNLRIPrefix, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpIpv4SrPolicyNLRIPrefix()
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

// BgpIpv4SrPolicyNLRIPrefix is iPv4 Segment Routing Policy NLRI Prefix.
type BgpIpv4SrPolicyNLRIPrefix interface {
	Validation
	// msg marshals BgpIpv4SrPolicyNLRIPrefix to protobuf object *otg.BgpIpv4SrPolicyNLRIPrefix
	// and doesn't set defaults
	msg() *otg.BgpIpv4SrPolicyNLRIPrefix
	// setMsg unmarshals BgpIpv4SrPolicyNLRIPrefix from protobuf object *otg.BgpIpv4SrPolicyNLRIPrefix
	// and doesn't set defaults
	setMsg(*otg.BgpIpv4SrPolicyNLRIPrefix) BgpIpv4SrPolicyNLRIPrefix
	// provides marshal interface
	Marshal() marshalBgpIpv4SrPolicyNLRIPrefix
	// provides unmarshal interface
	Unmarshal() unMarshalBgpIpv4SrPolicyNLRIPrefix
	// validate validates BgpIpv4SrPolicyNLRIPrefix
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpIpv4SrPolicyNLRIPrefix, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Distinguisher returns uint32, set in BgpIpv4SrPolicyNLRIPrefix.
	Distinguisher() uint32
	// SetDistinguisher assigns uint32 provided by user to BgpIpv4SrPolicyNLRIPrefix
	SetDistinguisher(value uint32) BgpIpv4SrPolicyNLRIPrefix
	// HasDistinguisher checks if Distinguisher has been set in BgpIpv4SrPolicyNLRIPrefix
	HasDistinguisher() bool
	// Color returns uint32, set in BgpIpv4SrPolicyNLRIPrefix.
	Color() uint32
	// SetColor assigns uint32 provided by user to BgpIpv4SrPolicyNLRIPrefix
	SetColor(value uint32) BgpIpv4SrPolicyNLRIPrefix
	// HasColor checks if Color has been set in BgpIpv4SrPolicyNLRIPrefix
	HasColor() bool
	// Endpoint returns string, set in BgpIpv4SrPolicyNLRIPrefix.
	Endpoint() string
	// SetEndpoint assigns string provided by user to BgpIpv4SrPolicyNLRIPrefix
	SetEndpoint(value string) BgpIpv4SrPolicyNLRIPrefix
	// HasEndpoint checks if Endpoint has been set in BgpIpv4SrPolicyNLRIPrefix
	HasEndpoint() bool
}

// The 4-octet value uniquely identifying the policy in the context of <color, endpoint> tuple.  The distinguisher has no semantic value and is solely used by the SR Policy originator to make unique (from an NLRI perspective)  both for multiple candidate paths of the same SR Policy as well as candidate paths of different SR Policies  (i.e. with different segment lists) with the same Color and Endpoint but meant for different headends.
// Distinguisher returns a uint32
func (obj *bgpIpv4SrPolicyNLRIPrefix) Distinguisher() uint32 {

	return *obj.obj.Distinguisher

}

// The 4-octet value uniquely identifying the policy in the context of <color, endpoint> tuple.  The distinguisher has no semantic value and is solely used by the SR Policy originator to make unique (from an NLRI perspective)  both for multiple candidate paths of the same SR Policy as well as candidate paths of different SR Policies  (i.e. with different segment lists) with the same Color and Endpoint but meant for different headends.
// Distinguisher returns a uint32
func (obj *bgpIpv4SrPolicyNLRIPrefix) HasDistinguisher() bool {
	return obj.obj.Distinguisher != nil
}

// The 4-octet value uniquely identifying the policy in the context of <color, endpoint> tuple.  The distinguisher has no semantic value and is solely used by the SR Policy originator to make unique (from an NLRI perspective)  both for multiple candidate paths of the same SR Policy as well as candidate paths of different SR Policies  (i.e. with different segment lists) with the same Color and Endpoint but meant for different headends.
// SetDistinguisher sets the uint32 value in the BgpIpv4SrPolicyNLRIPrefix object
func (obj *bgpIpv4SrPolicyNLRIPrefix) SetDistinguisher(value uint32) BgpIpv4SrPolicyNLRIPrefix {

	obj.obj.Distinguisher = &value
	return obj
}

// 4-octet value identifying (with the endpoint) the policy. The color is used to match the color of the destination  prefixes to steer traffic into the SR Policy as specified in section 8 of RFC9256.
// Color returns a uint32
func (obj *bgpIpv4SrPolicyNLRIPrefix) Color() uint32 {

	return *obj.obj.Color

}

// 4-octet value identifying (with the endpoint) the policy. The color is used to match the color of the destination  prefixes to steer traffic into the SR Policy as specified in section 8 of RFC9256.
// Color returns a uint32
func (obj *bgpIpv4SrPolicyNLRIPrefix) HasColor() bool {
	return obj.obj.Color != nil
}

// 4-octet value identifying (with the endpoint) the policy. The color is used to match the color of the destination  prefixes to steer traffic into the SR Policy as specified in section 8 of RFC9256.
// SetColor sets the uint32 value in the BgpIpv4SrPolicyNLRIPrefix object
func (obj *bgpIpv4SrPolicyNLRIPrefix) SetColor(value uint32) BgpIpv4SrPolicyNLRIPrefix {

	obj.obj.Color = &value
	return obj
}

// Identifies the endpoint of a policy.  The Endpoint is an IPv4 address and can be either a unicast or an unspecified address (0.0.0.0) as specified in section 2.1 of RFC9256.
// Endpoint returns a string
func (obj *bgpIpv4SrPolicyNLRIPrefix) Endpoint() string {

	return *obj.obj.Endpoint

}

// Identifies the endpoint of a policy.  The Endpoint is an IPv4 address and can be either a unicast or an unspecified address (0.0.0.0) as specified in section 2.1 of RFC9256.
// Endpoint returns a string
func (obj *bgpIpv4SrPolicyNLRIPrefix) HasEndpoint() bool {
	return obj.obj.Endpoint != nil
}

// Identifies the endpoint of a policy.  The Endpoint is an IPv4 address and can be either a unicast or an unspecified address (0.0.0.0) as specified in section 2.1 of RFC9256.
// SetEndpoint sets the string value in the BgpIpv4SrPolicyNLRIPrefix object
func (obj *bgpIpv4SrPolicyNLRIPrefix) SetEndpoint(value string) BgpIpv4SrPolicyNLRIPrefix {

	obj.obj.Endpoint = &value
	return obj
}

func (obj *bgpIpv4SrPolicyNLRIPrefix) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Endpoint != nil {

		err := obj.validateIpv4(obj.Endpoint())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpIpv4SrPolicyNLRIPrefix.Endpoint"))
		}

	}

}

func (obj *bgpIpv4SrPolicyNLRIPrefix) setDefault() {
	if obj.obj.Distinguisher == nil {
		obj.SetDistinguisher(1)
	}
	if obj.obj.Color == nil {
		obj.SetColor(1)
	}
	if obj.obj.Endpoint == nil {
		obj.SetEndpoint("0.0.0.0")
	}

}
