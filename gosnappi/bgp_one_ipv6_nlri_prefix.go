package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpOneIpv6NLRIPrefix *****
type bgpOneIpv6NLRIPrefix struct {
	validation
	obj          *otg.BgpOneIpv6NLRIPrefix
	marshaller   marshalBgpOneIpv6NLRIPrefix
	unMarshaller unMarshalBgpOneIpv6NLRIPrefix
	pathIdHolder BgpNLRIPrefixPathId
}

func NewBgpOneIpv6NLRIPrefix() BgpOneIpv6NLRIPrefix {
	obj := bgpOneIpv6NLRIPrefix{obj: &otg.BgpOneIpv6NLRIPrefix{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpOneIpv6NLRIPrefix) msg() *otg.BgpOneIpv6NLRIPrefix {
	return obj.obj
}

func (obj *bgpOneIpv6NLRIPrefix) setMsg(msg *otg.BgpOneIpv6NLRIPrefix) BgpOneIpv6NLRIPrefix {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpOneIpv6NLRIPrefix struct {
	obj *bgpOneIpv6NLRIPrefix
}

type marshalBgpOneIpv6NLRIPrefix interface {
	// ToProto marshals BgpOneIpv6NLRIPrefix to protobuf object *otg.BgpOneIpv6NLRIPrefix
	ToProto() (*otg.BgpOneIpv6NLRIPrefix, error)
	// ToPbText marshals BgpOneIpv6NLRIPrefix to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpOneIpv6NLRIPrefix to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpOneIpv6NLRIPrefix to JSON text
	ToJson() (string, error)
}

type unMarshalbgpOneIpv6NLRIPrefix struct {
	obj *bgpOneIpv6NLRIPrefix
}

type unMarshalBgpOneIpv6NLRIPrefix interface {
	// FromProto unmarshals BgpOneIpv6NLRIPrefix from protobuf object *otg.BgpOneIpv6NLRIPrefix
	FromProto(msg *otg.BgpOneIpv6NLRIPrefix) (BgpOneIpv6NLRIPrefix, error)
	// FromPbText unmarshals BgpOneIpv6NLRIPrefix from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpOneIpv6NLRIPrefix from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpOneIpv6NLRIPrefix from JSON text
	FromJson(value string) error
}

func (obj *bgpOneIpv6NLRIPrefix) Marshal() marshalBgpOneIpv6NLRIPrefix {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpOneIpv6NLRIPrefix{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpOneIpv6NLRIPrefix) Unmarshal() unMarshalBgpOneIpv6NLRIPrefix {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpOneIpv6NLRIPrefix{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpOneIpv6NLRIPrefix) ToProto() (*otg.BgpOneIpv6NLRIPrefix, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpOneIpv6NLRIPrefix) FromProto(msg *otg.BgpOneIpv6NLRIPrefix) (BgpOneIpv6NLRIPrefix, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpOneIpv6NLRIPrefix) ToPbText() (string, error) {
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

func (m *unMarshalbgpOneIpv6NLRIPrefix) FromPbText(value string) error {
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

func (m *marshalbgpOneIpv6NLRIPrefix) ToYaml() (string, error) {
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

func (m *unMarshalbgpOneIpv6NLRIPrefix) FromYaml(value string) error {
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

func (m *marshalbgpOneIpv6NLRIPrefix) ToJson() (string, error) {
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

func (m *unMarshalbgpOneIpv6NLRIPrefix) FromJson(value string) error {
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

func (obj *bgpOneIpv6NLRIPrefix) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpOneIpv6NLRIPrefix) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpOneIpv6NLRIPrefix) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpOneIpv6NLRIPrefix) Clone() (BgpOneIpv6NLRIPrefix, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpOneIpv6NLRIPrefix()
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

func (obj *bgpOneIpv6NLRIPrefix) setNil() {
	obj.pathIdHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpOneIpv6NLRIPrefix is one IPv6 NLRI Prefix.
type BgpOneIpv6NLRIPrefix interface {
	Validation
	// msg marshals BgpOneIpv6NLRIPrefix to protobuf object *otg.BgpOneIpv6NLRIPrefix
	// and doesn't set defaults
	msg() *otg.BgpOneIpv6NLRIPrefix
	// setMsg unmarshals BgpOneIpv6NLRIPrefix from protobuf object *otg.BgpOneIpv6NLRIPrefix
	// and doesn't set defaults
	setMsg(*otg.BgpOneIpv6NLRIPrefix) BgpOneIpv6NLRIPrefix
	// provides marshal interface
	Marshal() marshalBgpOneIpv6NLRIPrefix
	// provides unmarshal interface
	Unmarshal() unMarshalBgpOneIpv6NLRIPrefix
	// validate validates BgpOneIpv6NLRIPrefix
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpOneIpv6NLRIPrefix, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Address returns string, set in BgpOneIpv6NLRIPrefix.
	Address() string
	// SetAddress assigns string provided by user to BgpOneIpv6NLRIPrefix
	SetAddress(value string) BgpOneIpv6NLRIPrefix
	// HasAddress checks if Address has been set in BgpOneIpv6NLRIPrefix
	HasAddress() bool
	// Prefix returns uint32, set in BgpOneIpv6NLRIPrefix.
	Prefix() uint32
	// SetPrefix assigns uint32 provided by user to BgpOneIpv6NLRIPrefix
	SetPrefix(value uint32) BgpOneIpv6NLRIPrefix
	// HasPrefix checks if Prefix has been set in BgpOneIpv6NLRIPrefix
	HasPrefix() bool
	// PathId returns BgpNLRIPrefixPathId, set in BgpOneIpv6NLRIPrefix.
	// BgpNLRIPrefixPathId is optional field in the NLRI carrying Path Id of the prefix.
	PathId() BgpNLRIPrefixPathId
	// SetPathId assigns BgpNLRIPrefixPathId provided by user to BgpOneIpv6NLRIPrefix.
	// BgpNLRIPrefixPathId is optional field in the NLRI carrying Path Id of the prefix.
	SetPathId(value BgpNLRIPrefixPathId) BgpOneIpv6NLRIPrefix
	// HasPathId checks if PathId has been set in BgpOneIpv6NLRIPrefix
	HasPathId() bool
	setNil()
}

// The IPv6 address of the network.
// Address returns a string
func (obj *bgpOneIpv6NLRIPrefix) Address() string {

	return *obj.obj.Address

}

// The IPv6 address of the network.
// Address returns a string
func (obj *bgpOneIpv6NLRIPrefix) HasAddress() bool {
	return obj.obj.Address != nil
}

// The IPv6 address of the network.
// SetAddress sets the string value in the BgpOneIpv6NLRIPrefix object
func (obj *bgpOneIpv6NLRIPrefix) SetAddress(value string) BgpOneIpv6NLRIPrefix {

	obj.obj.Address = &value
	return obj
}

// The IPv6 network prefix length to be applied to the address.
// Prefix returns a uint32
func (obj *bgpOneIpv6NLRIPrefix) Prefix() uint32 {

	return *obj.obj.Prefix

}

// The IPv6 network prefix length to be applied to the address.
// Prefix returns a uint32
func (obj *bgpOneIpv6NLRIPrefix) HasPrefix() bool {
	return obj.obj.Prefix != nil
}

// The IPv6 network prefix length to be applied to the address.
// SetPrefix sets the uint32 value in the BgpOneIpv6NLRIPrefix object
func (obj *bgpOneIpv6NLRIPrefix) SetPrefix(value uint32) BgpOneIpv6NLRIPrefix {

	obj.obj.Prefix = &value
	return obj
}

// description is TBD
// PathId returns a BgpNLRIPrefixPathId
func (obj *bgpOneIpv6NLRIPrefix) PathId() BgpNLRIPrefixPathId {
	if obj.obj.PathId == nil {
		obj.obj.PathId = NewBgpNLRIPrefixPathId().msg()
	}
	if obj.pathIdHolder == nil {
		obj.pathIdHolder = &bgpNLRIPrefixPathId{obj: obj.obj.PathId}
	}
	return obj.pathIdHolder
}

// description is TBD
// PathId returns a BgpNLRIPrefixPathId
func (obj *bgpOneIpv6NLRIPrefix) HasPathId() bool {
	return obj.obj.PathId != nil
}

// description is TBD
// SetPathId sets the BgpNLRIPrefixPathId value in the BgpOneIpv6NLRIPrefix object
func (obj *bgpOneIpv6NLRIPrefix) SetPathId(value BgpNLRIPrefixPathId) BgpOneIpv6NLRIPrefix {

	obj.pathIdHolder = nil
	obj.obj.PathId = value.msg()

	return obj
}

func (obj *bgpOneIpv6NLRIPrefix) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Address != nil {

		err := obj.validateIpv6(obj.Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpOneIpv6NLRIPrefix.Address"))
		}

	}

	if obj.obj.Prefix != nil {

		if *obj.obj.Prefix > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpOneIpv6NLRIPrefix.Prefix <= 128 but Got %d", *obj.obj.Prefix))
		}

	}

	if obj.obj.PathId != nil {

		obj.PathId().validateObj(vObj, set_default)
	}

}

func (obj *bgpOneIpv6NLRIPrefix) setDefault() {
	if obj.obj.Address == nil {
		obj.SetAddress("0::0")
	}
	if obj.obj.Prefix == nil {
		obj.SetPrefix(64)
	}

}
