package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ClientOptionsLinkLayerAddress *****
type dhcpv6ClientOptionsLinkLayerAddress struct {
	validation
	obj          *otg.Dhcpv6ClientOptionsLinkLayerAddress
	marshaller   marshalDhcpv6ClientOptionsLinkLayerAddress
	unMarshaller unMarshalDhcpv6ClientOptionsLinkLayerAddress
}

func NewDhcpv6ClientOptionsLinkLayerAddress() Dhcpv6ClientOptionsLinkLayerAddress {
	obj := dhcpv6ClientOptionsLinkLayerAddress{obj: &otg.Dhcpv6ClientOptionsLinkLayerAddress{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ClientOptionsLinkLayerAddress) msg() *otg.Dhcpv6ClientOptionsLinkLayerAddress {
	return obj.obj
}

func (obj *dhcpv6ClientOptionsLinkLayerAddress) setMsg(msg *otg.Dhcpv6ClientOptionsLinkLayerAddress) Dhcpv6ClientOptionsLinkLayerAddress {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ClientOptionsLinkLayerAddress struct {
	obj *dhcpv6ClientOptionsLinkLayerAddress
}

type marshalDhcpv6ClientOptionsLinkLayerAddress interface {
	// ToProto marshals Dhcpv6ClientOptionsLinkLayerAddress to protobuf object *otg.Dhcpv6ClientOptionsLinkLayerAddress
	ToProto() (*otg.Dhcpv6ClientOptionsLinkLayerAddress, error)
	// ToPbText marshals Dhcpv6ClientOptionsLinkLayerAddress to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ClientOptionsLinkLayerAddress to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ClientOptionsLinkLayerAddress to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Dhcpv6ClientOptionsLinkLayerAddress to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldhcpv6ClientOptionsLinkLayerAddress struct {
	obj *dhcpv6ClientOptionsLinkLayerAddress
}

type unMarshalDhcpv6ClientOptionsLinkLayerAddress interface {
	// FromProto unmarshals Dhcpv6ClientOptionsLinkLayerAddress from protobuf object *otg.Dhcpv6ClientOptionsLinkLayerAddress
	FromProto(msg *otg.Dhcpv6ClientOptionsLinkLayerAddress) (Dhcpv6ClientOptionsLinkLayerAddress, error)
	// FromPbText unmarshals Dhcpv6ClientOptionsLinkLayerAddress from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ClientOptionsLinkLayerAddress from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ClientOptionsLinkLayerAddress from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ClientOptionsLinkLayerAddress) Marshal() marshalDhcpv6ClientOptionsLinkLayerAddress {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ClientOptionsLinkLayerAddress{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ClientOptionsLinkLayerAddress) Unmarshal() unMarshalDhcpv6ClientOptionsLinkLayerAddress {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ClientOptionsLinkLayerAddress{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ClientOptionsLinkLayerAddress) ToProto() (*otg.Dhcpv6ClientOptionsLinkLayerAddress, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ClientOptionsLinkLayerAddress) FromProto(msg *otg.Dhcpv6ClientOptionsLinkLayerAddress) (Dhcpv6ClientOptionsLinkLayerAddress, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ClientOptionsLinkLayerAddress) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsLinkLayerAddress) FromPbText(value string) error {
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

func (m *marshaldhcpv6ClientOptionsLinkLayerAddress) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsLinkLayerAddress) FromYaml(value string) error {
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

func (m *marshaldhcpv6ClientOptionsLinkLayerAddress) ToJsonRaw() (string, error) {
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

func (m *marshaldhcpv6ClientOptionsLinkLayerAddress) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsLinkLayerAddress) FromJson(value string) error {
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

func (obj *dhcpv6ClientOptionsLinkLayerAddress) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsLinkLayerAddress) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsLinkLayerAddress) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ClientOptionsLinkLayerAddress) Clone() (Dhcpv6ClientOptionsLinkLayerAddress, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ClientOptionsLinkLayerAddress()
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

// Dhcpv6ClientOptionsLinkLayerAddress is the link-layer address configured in DUID llt or DUID ll.
type Dhcpv6ClientOptionsLinkLayerAddress interface {
	Validation
	// msg marshals Dhcpv6ClientOptionsLinkLayerAddress to protobuf object *otg.Dhcpv6ClientOptionsLinkLayerAddress
	// and doesn't set defaults
	msg() *otg.Dhcpv6ClientOptionsLinkLayerAddress
	// setMsg unmarshals Dhcpv6ClientOptionsLinkLayerAddress from protobuf object *otg.Dhcpv6ClientOptionsLinkLayerAddress
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ClientOptionsLinkLayerAddress) Dhcpv6ClientOptionsLinkLayerAddress
	// provides marshal interface
	Marshal() marshalDhcpv6ClientOptionsLinkLayerAddress
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ClientOptionsLinkLayerAddress
	// validate validates Dhcpv6ClientOptionsLinkLayerAddress
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ClientOptionsLinkLayerAddress, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Value returns string, set in Dhcpv6ClientOptionsLinkLayerAddress.
	Value() string
	// SetValue assigns string provided by user to Dhcpv6ClientOptionsLinkLayerAddress
	SetValue(value string) Dhcpv6ClientOptionsLinkLayerAddress
}

// The MAC address that becomes part of DUID llt or DUID ll.
// Value returns a string
func (obj *dhcpv6ClientOptionsLinkLayerAddress) Value() string {

	return *obj.obj.Value

}

// The MAC address that becomes part of DUID llt or DUID ll.
// SetValue sets the string value in the Dhcpv6ClientOptionsLinkLayerAddress object
func (obj *dhcpv6ClientOptionsLinkLayerAddress) SetValue(value string) Dhcpv6ClientOptionsLinkLayerAddress {

	obj.obj.Value = &value
	return obj
}

func (obj *dhcpv6ClientOptionsLinkLayerAddress) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Value is required
	if obj.obj.Value == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Value is required field on interface Dhcpv6ClientOptionsLinkLayerAddress")
	}
	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Dhcpv6ClientOptionsLinkLayerAddress.Value"))
		}

	}

}

func (obj *dhcpv6ClientOptionsLinkLayerAddress) setDefault() {

}
