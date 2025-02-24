package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpLearnedInformationFilter *****
type bgpLearnedInformationFilter struct {
	validation
	obj          *otg.BgpLearnedInformationFilter
	marshaller   marshalBgpLearnedInformationFilter
	unMarshaller unMarshalBgpLearnedInformationFilter
}

func NewBgpLearnedInformationFilter() BgpLearnedInformationFilter {
	obj := bgpLearnedInformationFilter{obj: &otg.BgpLearnedInformationFilter{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpLearnedInformationFilter) msg() *otg.BgpLearnedInformationFilter {
	return obj.obj
}

func (obj *bgpLearnedInformationFilter) setMsg(msg *otg.BgpLearnedInformationFilter) BgpLearnedInformationFilter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpLearnedInformationFilter struct {
	obj *bgpLearnedInformationFilter
}

type marshalBgpLearnedInformationFilter interface {
	// ToProto marshals BgpLearnedInformationFilter to protobuf object *otg.BgpLearnedInformationFilter
	ToProto() (*otg.BgpLearnedInformationFilter, error)
	// ToPbText marshals BgpLearnedInformationFilter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpLearnedInformationFilter to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpLearnedInformationFilter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpLearnedInformationFilter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpLearnedInformationFilter struct {
	obj *bgpLearnedInformationFilter
}

type unMarshalBgpLearnedInformationFilter interface {
	// FromProto unmarshals BgpLearnedInformationFilter from protobuf object *otg.BgpLearnedInformationFilter
	FromProto(msg *otg.BgpLearnedInformationFilter) (BgpLearnedInformationFilter, error)
	// FromPbText unmarshals BgpLearnedInformationFilter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpLearnedInformationFilter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpLearnedInformationFilter from JSON text
	FromJson(value string) error
}

func (obj *bgpLearnedInformationFilter) Marshal() marshalBgpLearnedInformationFilter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpLearnedInformationFilter{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpLearnedInformationFilter) Unmarshal() unMarshalBgpLearnedInformationFilter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpLearnedInformationFilter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpLearnedInformationFilter) ToProto() (*otg.BgpLearnedInformationFilter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpLearnedInformationFilter) FromProto(msg *otg.BgpLearnedInformationFilter) (BgpLearnedInformationFilter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpLearnedInformationFilter) ToPbText() (string, error) {
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

func (m *unMarshalbgpLearnedInformationFilter) FromPbText(value string) error {
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

func (m *marshalbgpLearnedInformationFilter) ToYaml() (string, error) {
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

func (m *unMarshalbgpLearnedInformationFilter) FromYaml(value string) error {
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

func (m *marshalbgpLearnedInformationFilter) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalbgpLearnedInformationFilter) ToJson() (string, error) {
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

func (m *unMarshalbgpLearnedInformationFilter) FromJson(value string) error {
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

func (obj *bgpLearnedInformationFilter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpLearnedInformationFilter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpLearnedInformationFilter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpLearnedInformationFilter) Clone() (BgpLearnedInformationFilter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpLearnedInformationFilter()
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

// BgpLearnedInformationFilter is configuration for controlling storage of BGP learned information recieved from the peer.
type BgpLearnedInformationFilter interface {
	Validation
	// msg marshals BgpLearnedInformationFilter to protobuf object *otg.BgpLearnedInformationFilter
	// and doesn't set defaults
	msg() *otg.BgpLearnedInformationFilter
	// setMsg unmarshals BgpLearnedInformationFilter from protobuf object *otg.BgpLearnedInformationFilter
	// and doesn't set defaults
	setMsg(*otg.BgpLearnedInformationFilter) BgpLearnedInformationFilter
	// provides marshal interface
	Marshal() marshalBgpLearnedInformationFilter
	// provides unmarshal interface
	Unmarshal() unMarshalBgpLearnedInformationFilter
	// validate validates BgpLearnedInformationFilter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpLearnedInformationFilter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// UnicastIpv4Prefix returns bool, set in BgpLearnedInformationFilter.
	UnicastIpv4Prefix() bool
	// SetUnicastIpv4Prefix assigns bool provided by user to BgpLearnedInformationFilter
	SetUnicastIpv4Prefix(value bool) BgpLearnedInformationFilter
	// HasUnicastIpv4Prefix checks if UnicastIpv4Prefix has been set in BgpLearnedInformationFilter
	HasUnicastIpv4Prefix() bool
	// UnicastIpv6Prefix returns bool, set in BgpLearnedInformationFilter.
	UnicastIpv6Prefix() bool
	// SetUnicastIpv6Prefix assigns bool provided by user to BgpLearnedInformationFilter
	SetUnicastIpv6Prefix(value bool) BgpLearnedInformationFilter
	// HasUnicastIpv6Prefix checks if UnicastIpv6Prefix has been set in BgpLearnedInformationFilter
	HasUnicastIpv6Prefix() bool
}

// If enabled, will store the information related to Unicast IPv4 Prefixes recieved from the peer.
// UnicastIpv4Prefix returns a bool
func (obj *bgpLearnedInformationFilter) UnicastIpv4Prefix() bool {

	return *obj.obj.UnicastIpv4Prefix

}

// If enabled, will store the information related to Unicast IPv4 Prefixes recieved from the peer.
// UnicastIpv4Prefix returns a bool
func (obj *bgpLearnedInformationFilter) HasUnicastIpv4Prefix() bool {
	return obj.obj.UnicastIpv4Prefix != nil
}

// If enabled, will store the information related to Unicast IPv4 Prefixes recieved from the peer.
// SetUnicastIpv4Prefix sets the bool value in the BgpLearnedInformationFilter object
func (obj *bgpLearnedInformationFilter) SetUnicastIpv4Prefix(value bool) BgpLearnedInformationFilter {

	obj.obj.UnicastIpv4Prefix = &value
	return obj
}

// If enabled, will store the information related to Unicast IPv6 Prefixes recieved from the peer.
// UnicastIpv6Prefix returns a bool
func (obj *bgpLearnedInformationFilter) UnicastIpv6Prefix() bool {

	return *obj.obj.UnicastIpv6Prefix

}

// If enabled, will store the information related to Unicast IPv6 Prefixes recieved from the peer.
// UnicastIpv6Prefix returns a bool
func (obj *bgpLearnedInformationFilter) HasUnicastIpv6Prefix() bool {
	return obj.obj.UnicastIpv6Prefix != nil
}

// If enabled, will store the information related to Unicast IPv6 Prefixes recieved from the peer.
// SetUnicastIpv6Prefix sets the bool value in the BgpLearnedInformationFilter object
func (obj *bgpLearnedInformationFilter) SetUnicastIpv6Prefix(value bool) BgpLearnedInformationFilter {

	obj.obj.UnicastIpv6Prefix = &value
	return obj
}

func (obj *bgpLearnedInformationFilter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *bgpLearnedInformationFilter) setDefault() {
	if obj.obj.UnicastIpv4Prefix == nil {
		obj.SetUnicastIpv4Prefix(false)
	}
	if obj.obj.UnicastIpv6Prefix == nil {
		obj.SetUnicastIpv6Prefix(false)
	}

}
