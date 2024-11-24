package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ClientOptionsDuidLl *****
type dhcpv6ClientOptionsDuidLl struct {
	validation
	obj                    *otg.Dhcpv6ClientOptionsDuidLl
	marshaller             marshalDhcpv6ClientOptionsDuidLl
	unMarshaller           unMarshalDhcpv6ClientOptionsDuidLl
	linkLayerAddressHolder Dhcpv6ClientOptionsLinkLayerAddress
}

func NewDhcpv6ClientOptionsDuidLl() Dhcpv6ClientOptionsDuidLl {
	obj := dhcpv6ClientOptionsDuidLl{obj: &otg.Dhcpv6ClientOptionsDuidLl{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ClientOptionsDuidLl) msg() *otg.Dhcpv6ClientOptionsDuidLl {
	return obj.obj
}

func (obj *dhcpv6ClientOptionsDuidLl) setMsg(msg *otg.Dhcpv6ClientOptionsDuidLl) Dhcpv6ClientOptionsDuidLl {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ClientOptionsDuidLl struct {
	obj *dhcpv6ClientOptionsDuidLl
}

type marshalDhcpv6ClientOptionsDuidLl interface {
	// ToProto marshals Dhcpv6ClientOptionsDuidLl to protobuf object *otg.Dhcpv6ClientOptionsDuidLl
	ToProto() (*otg.Dhcpv6ClientOptionsDuidLl, error)
	// ToPbText marshals Dhcpv6ClientOptionsDuidLl to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ClientOptionsDuidLl to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ClientOptionsDuidLl to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Dhcpv6ClientOptionsDuidLl to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldhcpv6ClientOptionsDuidLl struct {
	obj *dhcpv6ClientOptionsDuidLl
}

type unMarshalDhcpv6ClientOptionsDuidLl interface {
	// FromProto unmarshals Dhcpv6ClientOptionsDuidLl from protobuf object *otg.Dhcpv6ClientOptionsDuidLl
	FromProto(msg *otg.Dhcpv6ClientOptionsDuidLl) (Dhcpv6ClientOptionsDuidLl, error)
	// FromPbText unmarshals Dhcpv6ClientOptionsDuidLl from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ClientOptionsDuidLl from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ClientOptionsDuidLl from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ClientOptionsDuidLl) Marshal() marshalDhcpv6ClientOptionsDuidLl {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ClientOptionsDuidLl{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ClientOptionsDuidLl) Unmarshal() unMarshalDhcpv6ClientOptionsDuidLl {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ClientOptionsDuidLl{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ClientOptionsDuidLl) ToProto() (*otg.Dhcpv6ClientOptionsDuidLl, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ClientOptionsDuidLl) FromProto(msg *otg.Dhcpv6ClientOptionsDuidLl) (Dhcpv6ClientOptionsDuidLl, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ClientOptionsDuidLl) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsDuidLl) FromPbText(value string) error {
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

func (m *marshaldhcpv6ClientOptionsDuidLl) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsDuidLl) FromYaml(value string) error {
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

func (m *marshaldhcpv6ClientOptionsDuidLl) ToJsonRaw() (string, error) {
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

func (m *marshaldhcpv6ClientOptionsDuidLl) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsDuidLl) FromJson(value string) error {
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

func (obj *dhcpv6ClientOptionsDuidLl) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsDuidLl) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsDuidLl) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ClientOptionsDuidLl) Clone() (Dhcpv6ClientOptionsDuidLl, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ClientOptionsDuidLl()
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

func (obj *dhcpv6ClientOptionsDuidLl) setNil() {
	obj.linkLayerAddressHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Dhcpv6ClientOptionsDuidLl is dUID based on Link Layer address. Hardware Type will be auto assigned to ethernet type.
type Dhcpv6ClientOptionsDuidLl interface {
	Validation
	// msg marshals Dhcpv6ClientOptionsDuidLl to protobuf object *otg.Dhcpv6ClientOptionsDuidLl
	// and doesn't set defaults
	msg() *otg.Dhcpv6ClientOptionsDuidLl
	// setMsg unmarshals Dhcpv6ClientOptionsDuidLl from protobuf object *otg.Dhcpv6ClientOptionsDuidLl
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ClientOptionsDuidLl) Dhcpv6ClientOptionsDuidLl
	// provides marshal interface
	Marshal() marshalDhcpv6ClientOptionsDuidLl
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ClientOptionsDuidLl
	// validate validates Dhcpv6ClientOptionsDuidLl
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ClientOptionsDuidLl, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LinkLayerAddress returns Dhcpv6ClientOptionsLinkLayerAddress, set in Dhcpv6ClientOptionsDuidLl.
	// Dhcpv6ClientOptionsLinkLayerAddress is the link-layer address configured in DUID llt or DUID ll.
	LinkLayerAddress() Dhcpv6ClientOptionsLinkLayerAddress
	// SetLinkLayerAddress assigns Dhcpv6ClientOptionsLinkLayerAddress provided by user to Dhcpv6ClientOptionsDuidLl.
	// Dhcpv6ClientOptionsLinkLayerAddress is the link-layer address configured in DUID llt or DUID ll.
	SetLinkLayerAddress(value Dhcpv6ClientOptionsLinkLayerAddress) Dhcpv6ClientOptionsDuidLl
	setNil()
}

// The link-layer address is stored in canonical form, as described in RFC 2464.
// LinkLayerAddress returns a Dhcpv6ClientOptionsLinkLayerAddress
func (obj *dhcpv6ClientOptionsDuidLl) LinkLayerAddress() Dhcpv6ClientOptionsLinkLayerAddress {
	if obj.obj.LinkLayerAddress == nil {
		obj.obj.LinkLayerAddress = NewDhcpv6ClientOptionsLinkLayerAddress().msg()
	}
	if obj.linkLayerAddressHolder == nil {
		obj.linkLayerAddressHolder = &dhcpv6ClientOptionsLinkLayerAddress{obj: obj.obj.LinkLayerAddress}
	}
	return obj.linkLayerAddressHolder
}

// The link-layer address is stored in canonical form, as described in RFC 2464.
// SetLinkLayerAddress sets the Dhcpv6ClientOptionsLinkLayerAddress value in the Dhcpv6ClientOptionsDuidLl object
func (obj *dhcpv6ClientOptionsDuidLl) SetLinkLayerAddress(value Dhcpv6ClientOptionsLinkLayerAddress) Dhcpv6ClientOptionsDuidLl {

	obj.linkLayerAddressHolder = nil
	obj.obj.LinkLayerAddress = value.msg()

	return obj
}

func (obj *dhcpv6ClientOptionsDuidLl) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// LinkLayerAddress is required
	if obj.obj.LinkLayerAddress == nil {
		vObj.validationErrors = append(vObj.validationErrors, "LinkLayerAddress is required field on interface Dhcpv6ClientOptionsDuidLl")
	}

	if obj.obj.LinkLayerAddress != nil {

		obj.LinkLayerAddress().validateObj(vObj, set_default)
	}

}

func (obj *dhcpv6ClientOptionsDuidLl) setDefault() {

}
