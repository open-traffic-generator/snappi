package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpOneIpv4NLRIPrefix *****
type bgpOneIpv4NLRIPrefix struct {
	validation
	obj          *otg.BgpOneIpv4NLRIPrefix
	marshaller   marshalBgpOneIpv4NLRIPrefix
	unMarshaller unMarshalBgpOneIpv4NLRIPrefix
	pathIdHolder BgpNLRIPrefixPathId
}

func NewBgpOneIpv4NLRIPrefix() BgpOneIpv4NLRIPrefix {
	obj := bgpOneIpv4NLRIPrefix{obj: &otg.BgpOneIpv4NLRIPrefix{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpOneIpv4NLRIPrefix) msg() *otg.BgpOneIpv4NLRIPrefix {
	return obj.obj
}

func (obj *bgpOneIpv4NLRIPrefix) setMsg(msg *otg.BgpOneIpv4NLRIPrefix) BgpOneIpv4NLRIPrefix {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpOneIpv4NLRIPrefix struct {
	obj *bgpOneIpv4NLRIPrefix
}

type marshalBgpOneIpv4NLRIPrefix interface {
	// ToProto marshals BgpOneIpv4NLRIPrefix to protobuf object *otg.BgpOneIpv4NLRIPrefix
	ToProto() (*otg.BgpOneIpv4NLRIPrefix, error)
	// ToPbText marshals BgpOneIpv4NLRIPrefix to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpOneIpv4NLRIPrefix to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpOneIpv4NLRIPrefix to JSON text
	ToJson() (string, error)
}

type unMarshalbgpOneIpv4NLRIPrefix struct {
	obj *bgpOneIpv4NLRIPrefix
}

type unMarshalBgpOneIpv4NLRIPrefix interface {
	// FromProto unmarshals BgpOneIpv4NLRIPrefix from protobuf object *otg.BgpOneIpv4NLRIPrefix
	FromProto(msg *otg.BgpOneIpv4NLRIPrefix) (BgpOneIpv4NLRIPrefix, error)
	// FromPbText unmarshals BgpOneIpv4NLRIPrefix from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpOneIpv4NLRIPrefix from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpOneIpv4NLRIPrefix from JSON text
	FromJson(value string) error
}

func (obj *bgpOneIpv4NLRIPrefix) Marshal() marshalBgpOneIpv4NLRIPrefix {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpOneIpv4NLRIPrefix{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpOneIpv4NLRIPrefix) Unmarshal() unMarshalBgpOneIpv4NLRIPrefix {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpOneIpv4NLRIPrefix{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpOneIpv4NLRIPrefix) ToProto() (*otg.BgpOneIpv4NLRIPrefix, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpOneIpv4NLRIPrefix) FromProto(msg *otg.BgpOneIpv4NLRIPrefix) (BgpOneIpv4NLRIPrefix, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpOneIpv4NLRIPrefix) ToPbText() (string, error) {
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

func (m *unMarshalbgpOneIpv4NLRIPrefix) FromPbText(value string) error {
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

func (m *marshalbgpOneIpv4NLRIPrefix) ToYaml() (string, error) {
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

func (m *unMarshalbgpOneIpv4NLRIPrefix) FromYaml(value string) error {
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

func (m *marshalbgpOneIpv4NLRIPrefix) ToJson() (string, error) {
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

func (m *unMarshalbgpOneIpv4NLRIPrefix) FromJson(value string) error {
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

func (obj *bgpOneIpv4NLRIPrefix) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpOneIpv4NLRIPrefix) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpOneIpv4NLRIPrefix) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpOneIpv4NLRIPrefix) Clone() (BgpOneIpv4NLRIPrefix, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpOneIpv4NLRIPrefix()
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

func (obj *bgpOneIpv4NLRIPrefix) setNil() {
	obj.pathIdHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpOneIpv4NLRIPrefix is one IPv4 NLRI Prefix.
type BgpOneIpv4NLRIPrefix interface {
	Validation
	// msg marshals BgpOneIpv4NLRIPrefix to protobuf object *otg.BgpOneIpv4NLRIPrefix
	// and doesn't set defaults
	msg() *otg.BgpOneIpv4NLRIPrefix
	// setMsg unmarshals BgpOneIpv4NLRIPrefix from protobuf object *otg.BgpOneIpv4NLRIPrefix
	// and doesn't set defaults
	setMsg(*otg.BgpOneIpv4NLRIPrefix) BgpOneIpv4NLRIPrefix
	// provides marshal interface
	Marshal() marshalBgpOneIpv4NLRIPrefix
	// provides unmarshal interface
	Unmarshal() unMarshalBgpOneIpv4NLRIPrefix
	// validate validates BgpOneIpv4NLRIPrefix
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpOneIpv4NLRIPrefix, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Address returns string, set in BgpOneIpv4NLRIPrefix.
	Address() string
	// SetAddress assigns string provided by user to BgpOneIpv4NLRIPrefix
	SetAddress(value string) BgpOneIpv4NLRIPrefix
	// HasAddress checks if Address has been set in BgpOneIpv4NLRIPrefix
	HasAddress() bool
	// Prefix returns uint32, set in BgpOneIpv4NLRIPrefix.
	Prefix() uint32
	// SetPrefix assigns uint32 provided by user to BgpOneIpv4NLRIPrefix
	SetPrefix(value uint32) BgpOneIpv4NLRIPrefix
	// HasPrefix checks if Prefix has been set in BgpOneIpv4NLRIPrefix
	HasPrefix() bool
	// PathId returns BgpNLRIPrefixPathId, set in BgpOneIpv4NLRIPrefix.
	// BgpNLRIPrefixPathId is optional field in the NLRI carrying Path Id of the prefix.
	PathId() BgpNLRIPrefixPathId
	// SetPathId assigns BgpNLRIPrefixPathId provided by user to BgpOneIpv4NLRIPrefix.
	// BgpNLRIPrefixPathId is optional field in the NLRI carrying Path Id of the prefix.
	SetPathId(value BgpNLRIPrefixPathId) BgpOneIpv4NLRIPrefix
	// HasPathId checks if PathId has been set in BgpOneIpv4NLRIPrefix
	HasPathId() bool
	setNil()
}

// The IPv4 address of the network.
// Address returns a string
func (obj *bgpOneIpv4NLRIPrefix) Address() string {

	return *obj.obj.Address

}

// The IPv4 address of the network.
// Address returns a string
func (obj *bgpOneIpv4NLRIPrefix) HasAddress() bool {
	return obj.obj.Address != nil
}

// The IPv4 address of the network.
// SetAddress sets the string value in the BgpOneIpv4NLRIPrefix object
func (obj *bgpOneIpv4NLRIPrefix) SetAddress(value string) BgpOneIpv4NLRIPrefix {

	obj.obj.Address = &value
	return obj
}

// The IPv4 network prefix length to be applied to the address.
// Prefix returns a uint32
func (obj *bgpOneIpv4NLRIPrefix) Prefix() uint32 {

	return *obj.obj.Prefix

}

// The IPv4 network prefix length to be applied to the address.
// Prefix returns a uint32
func (obj *bgpOneIpv4NLRIPrefix) HasPrefix() bool {
	return obj.obj.Prefix != nil
}

// The IPv4 network prefix length to be applied to the address.
// SetPrefix sets the uint32 value in the BgpOneIpv4NLRIPrefix object
func (obj *bgpOneIpv4NLRIPrefix) SetPrefix(value uint32) BgpOneIpv4NLRIPrefix {

	obj.obj.Prefix = &value
	return obj
}

// description is TBD
// PathId returns a BgpNLRIPrefixPathId
func (obj *bgpOneIpv4NLRIPrefix) PathId() BgpNLRIPrefixPathId {
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
func (obj *bgpOneIpv4NLRIPrefix) HasPathId() bool {
	return obj.obj.PathId != nil
}

// description is TBD
// SetPathId sets the BgpNLRIPrefixPathId value in the BgpOneIpv4NLRIPrefix object
func (obj *bgpOneIpv4NLRIPrefix) SetPathId(value BgpNLRIPrefixPathId) BgpOneIpv4NLRIPrefix {

	obj.pathIdHolder = nil
	obj.obj.PathId = value.msg()

	return obj
}

func (obj *bgpOneIpv4NLRIPrefix) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Address != nil {

		err := obj.validateIpv4(obj.Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpOneIpv4NLRIPrefix.Address"))
		}

	}

	if obj.obj.Prefix != nil {

		if *obj.obj.Prefix > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpOneIpv4NLRIPrefix.Prefix <= 32 but Got %d", *obj.obj.Prefix))
		}

	}

	if obj.obj.PathId != nil {

		obj.PathId().validateObj(vObj, set_default)
	}

}

func (obj *bgpOneIpv4NLRIPrefix) setDefault() {
	if obj.obj.Address == nil {
		obj.SetAddress("0.0.0.0")
	}
	if obj.obj.Prefix == nil {
		obj.SetPrefix(24)
	}

}
