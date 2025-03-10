package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpOneTraditionalNlriPrefix *****
type bgpOneTraditionalNlriPrefix struct {
	validation
	obj          *otg.BgpOneTraditionalNlriPrefix
	marshaller   marshalBgpOneTraditionalNlriPrefix
	unMarshaller unMarshalBgpOneTraditionalNlriPrefix
	pathIdHolder BgpNLRIPrefixPathId
}

func NewBgpOneTraditionalNlriPrefix() BgpOneTraditionalNlriPrefix {
	obj := bgpOneTraditionalNlriPrefix{obj: &otg.BgpOneTraditionalNlriPrefix{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpOneTraditionalNlriPrefix) msg() *otg.BgpOneTraditionalNlriPrefix {
	return obj.obj
}

func (obj *bgpOneTraditionalNlriPrefix) setMsg(msg *otg.BgpOneTraditionalNlriPrefix) BgpOneTraditionalNlriPrefix {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpOneTraditionalNlriPrefix struct {
	obj *bgpOneTraditionalNlriPrefix
}

type marshalBgpOneTraditionalNlriPrefix interface {
	// ToProto marshals BgpOneTraditionalNlriPrefix to protobuf object *otg.BgpOneTraditionalNlriPrefix
	ToProto() (*otg.BgpOneTraditionalNlriPrefix, error)
	// ToPbText marshals BgpOneTraditionalNlriPrefix to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpOneTraditionalNlriPrefix to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpOneTraditionalNlriPrefix to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpOneTraditionalNlriPrefix to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpOneTraditionalNlriPrefix struct {
	obj *bgpOneTraditionalNlriPrefix
}

type unMarshalBgpOneTraditionalNlriPrefix interface {
	// FromProto unmarshals BgpOneTraditionalNlriPrefix from protobuf object *otg.BgpOneTraditionalNlriPrefix
	FromProto(msg *otg.BgpOneTraditionalNlriPrefix) (BgpOneTraditionalNlriPrefix, error)
	// FromPbText unmarshals BgpOneTraditionalNlriPrefix from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpOneTraditionalNlriPrefix from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpOneTraditionalNlriPrefix from JSON text
	FromJson(value string) error
}

func (obj *bgpOneTraditionalNlriPrefix) Marshal() marshalBgpOneTraditionalNlriPrefix {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpOneTraditionalNlriPrefix{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpOneTraditionalNlriPrefix) Unmarshal() unMarshalBgpOneTraditionalNlriPrefix {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpOneTraditionalNlriPrefix{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpOneTraditionalNlriPrefix) ToProto() (*otg.BgpOneTraditionalNlriPrefix, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpOneTraditionalNlriPrefix) FromProto(msg *otg.BgpOneTraditionalNlriPrefix) (BgpOneTraditionalNlriPrefix, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpOneTraditionalNlriPrefix) ToPbText() (string, error) {
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

func (m *unMarshalbgpOneTraditionalNlriPrefix) FromPbText(value string) error {
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

func (m *marshalbgpOneTraditionalNlriPrefix) ToYaml() (string, error) {
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

func (m *unMarshalbgpOneTraditionalNlriPrefix) FromYaml(value string) error {
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

func (m *marshalbgpOneTraditionalNlriPrefix) ToJsonRaw() (string, error) {
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

func (m *marshalbgpOneTraditionalNlriPrefix) ToJson() (string, error) {
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

func (m *unMarshalbgpOneTraditionalNlriPrefix) FromJson(value string) error {
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

func (obj *bgpOneTraditionalNlriPrefix) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpOneTraditionalNlriPrefix) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpOneTraditionalNlriPrefix) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpOneTraditionalNlriPrefix) Clone() (BgpOneTraditionalNlriPrefix, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpOneTraditionalNlriPrefix()
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

func (obj *bgpOneTraditionalNlriPrefix) setNil() {
	obj.pathIdHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpOneTraditionalNlriPrefix is tRADITIONAL_NLRI is an optional part of the the BGP Update which can carry only IPv4 prefix information as defined in https://www.rfc-editor.org/rfc/rfc4271.html#section-4.3
// and extended by https://datatracker.ietf.org/doc/html/rfc7911#section-3 to carry additional Path Id information per prefix.
type BgpOneTraditionalNlriPrefix interface {
	Validation
	// msg marshals BgpOneTraditionalNlriPrefix to protobuf object *otg.BgpOneTraditionalNlriPrefix
	// and doesn't set defaults
	msg() *otg.BgpOneTraditionalNlriPrefix
	// setMsg unmarshals BgpOneTraditionalNlriPrefix from protobuf object *otg.BgpOneTraditionalNlriPrefix
	// and doesn't set defaults
	setMsg(*otg.BgpOneTraditionalNlriPrefix) BgpOneTraditionalNlriPrefix
	// provides marshal interface
	Marshal() marshalBgpOneTraditionalNlriPrefix
	// provides unmarshal interface
	Unmarshal() unMarshalBgpOneTraditionalNlriPrefix
	// validate validates BgpOneTraditionalNlriPrefix
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpOneTraditionalNlriPrefix, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Address returns string, set in BgpOneTraditionalNlriPrefix.
	Address() string
	// SetAddress assigns string provided by user to BgpOneTraditionalNlriPrefix
	SetAddress(value string) BgpOneTraditionalNlriPrefix
	// HasAddress checks if Address has been set in BgpOneTraditionalNlriPrefix
	HasAddress() bool
	// Prefix returns uint32, set in BgpOneTraditionalNlriPrefix.
	Prefix() uint32
	// SetPrefix assigns uint32 provided by user to BgpOneTraditionalNlriPrefix
	SetPrefix(value uint32) BgpOneTraditionalNlriPrefix
	// HasPrefix checks if Prefix has been set in BgpOneTraditionalNlriPrefix
	HasPrefix() bool
	// PathId returns BgpNLRIPrefixPathId, set in BgpOneTraditionalNlriPrefix.
	// BgpNLRIPrefixPathId is optional field in the NLRI carrying Path Id of the prefix.
	PathId() BgpNLRIPrefixPathId
	// SetPathId assigns BgpNLRIPrefixPathId provided by user to BgpOneTraditionalNlriPrefix.
	// BgpNLRIPrefixPathId is optional field in the NLRI carrying Path Id of the prefix.
	SetPathId(value BgpNLRIPrefixPathId) BgpOneTraditionalNlriPrefix
	// HasPathId checks if PathId has been set in BgpOneTraditionalNlriPrefix
	HasPathId() bool
	setNil()
}

// The IPv4 address of the network.
// Address returns a string
func (obj *bgpOneTraditionalNlriPrefix) Address() string {

	return *obj.obj.Address

}

// The IPv4 address of the network.
// Address returns a string
func (obj *bgpOneTraditionalNlriPrefix) HasAddress() bool {
	return obj.obj.Address != nil
}

// The IPv4 address of the network.
// SetAddress sets the string value in the BgpOneTraditionalNlriPrefix object
func (obj *bgpOneTraditionalNlriPrefix) SetAddress(value string) BgpOneTraditionalNlriPrefix {

	obj.obj.Address = &value
	return obj
}

// The IPv4 network prefix length to be applied to the address.
// Prefix returns a uint32
func (obj *bgpOneTraditionalNlriPrefix) Prefix() uint32 {

	return *obj.obj.Prefix

}

// The IPv4 network prefix length to be applied to the address.
// Prefix returns a uint32
func (obj *bgpOneTraditionalNlriPrefix) HasPrefix() bool {
	return obj.obj.Prefix != nil
}

// The IPv4 network prefix length to be applied to the address.
// SetPrefix sets the uint32 value in the BgpOneTraditionalNlriPrefix object
func (obj *bgpOneTraditionalNlriPrefix) SetPrefix(value uint32) BgpOneTraditionalNlriPrefix {

	obj.obj.Prefix = &value
	return obj
}

// description is TBD
// PathId returns a BgpNLRIPrefixPathId
func (obj *bgpOneTraditionalNlriPrefix) PathId() BgpNLRIPrefixPathId {
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
func (obj *bgpOneTraditionalNlriPrefix) HasPathId() bool {
	return obj.obj.PathId != nil
}

// description is TBD
// SetPathId sets the BgpNLRIPrefixPathId value in the BgpOneTraditionalNlriPrefix object
func (obj *bgpOneTraditionalNlriPrefix) SetPathId(value BgpNLRIPrefixPathId) BgpOneTraditionalNlriPrefix {

	obj.pathIdHolder = nil
	obj.obj.PathId = value.msg()

	return obj
}

func (obj *bgpOneTraditionalNlriPrefix) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Address != nil {

		err := obj.validateIpv4(obj.Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpOneTraditionalNlriPrefix.Address"))
		}

	}

	if obj.obj.Prefix != nil {

		if *obj.obj.Prefix > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpOneTraditionalNlriPrefix.Prefix <= 32 but Got %d", *obj.obj.Prefix))
		}

	}

	if obj.obj.PathId != nil {

		obj.PathId().validateObj(vObj, set_default)
	}

}

func (obj *bgpOneTraditionalNlriPrefix) setDefault() {
	if obj.obj.Address == nil {
		obj.SetAddress("0.0.0.0")
	}
	if obj.obj.Prefix == nil {
		obj.SetPrefix(24)
	}

}
