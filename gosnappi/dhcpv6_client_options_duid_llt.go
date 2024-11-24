package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ClientOptionsDuidLlt *****
type dhcpv6ClientOptionsDuidLlt struct {
	validation
	obj                    *otg.Dhcpv6ClientOptionsDuidLlt
	marshaller             marshalDhcpv6ClientOptionsDuidLlt
	unMarshaller           unMarshalDhcpv6ClientOptionsDuidLlt
	linkLayerAddressHolder Dhcpv6ClientOptionsLinkLayerAddress
}

func NewDhcpv6ClientOptionsDuidLlt() Dhcpv6ClientOptionsDuidLlt {
	obj := dhcpv6ClientOptionsDuidLlt{obj: &otg.Dhcpv6ClientOptionsDuidLlt{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ClientOptionsDuidLlt) msg() *otg.Dhcpv6ClientOptionsDuidLlt {
	return obj.obj
}

func (obj *dhcpv6ClientOptionsDuidLlt) setMsg(msg *otg.Dhcpv6ClientOptionsDuidLlt) Dhcpv6ClientOptionsDuidLlt {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ClientOptionsDuidLlt struct {
	obj *dhcpv6ClientOptionsDuidLlt
}

type marshalDhcpv6ClientOptionsDuidLlt interface {
	// ToProto marshals Dhcpv6ClientOptionsDuidLlt to protobuf object *otg.Dhcpv6ClientOptionsDuidLlt
	ToProto() (*otg.Dhcpv6ClientOptionsDuidLlt, error)
	// ToPbText marshals Dhcpv6ClientOptionsDuidLlt to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ClientOptionsDuidLlt to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ClientOptionsDuidLlt to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Dhcpv6ClientOptionsDuidLlt to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldhcpv6ClientOptionsDuidLlt struct {
	obj *dhcpv6ClientOptionsDuidLlt
}

type unMarshalDhcpv6ClientOptionsDuidLlt interface {
	// FromProto unmarshals Dhcpv6ClientOptionsDuidLlt from protobuf object *otg.Dhcpv6ClientOptionsDuidLlt
	FromProto(msg *otg.Dhcpv6ClientOptionsDuidLlt) (Dhcpv6ClientOptionsDuidLlt, error)
	// FromPbText unmarshals Dhcpv6ClientOptionsDuidLlt from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ClientOptionsDuidLlt from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ClientOptionsDuidLlt from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ClientOptionsDuidLlt) Marshal() marshalDhcpv6ClientOptionsDuidLlt {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ClientOptionsDuidLlt{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ClientOptionsDuidLlt) Unmarshal() unMarshalDhcpv6ClientOptionsDuidLlt {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ClientOptionsDuidLlt{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ClientOptionsDuidLlt) ToProto() (*otg.Dhcpv6ClientOptionsDuidLlt, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ClientOptionsDuidLlt) FromProto(msg *otg.Dhcpv6ClientOptionsDuidLlt) (Dhcpv6ClientOptionsDuidLlt, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ClientOptionsDuidLlt) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsDuidLlt) FromPbText(value string) error {
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

func (m *marshaldhcpv6ClientOptionsDuidLlt) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsDuidLlt) FromYaml(value string) error {
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

func (m *marshaldhcpv6ClientOptionsDuidLlt) ToJsonRaw() (string, error) {
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

func (m *marshaldhcpv6ClientOptionsDuidLlt) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsDuidLlt) FromJson(value string) error {
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

func (obj *dhcpv6ClientOptionsDuidLlt) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsDuidLlt) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsDuidLlt) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ClientOptionsDuidLlt) Clone() (Dhcpv6ClientOptionsDuidLlt, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ClientOptionsDuidLlt()
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

func (obj *dhcpv6ClientOptionsDuidLlt) setNil() {
	obj.linkLayerAddressHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Dhcpv6ClientOptionsDuidLlt is dUID based on Link Layer address plus time. Hardware Type will be auto assigned to ethernet type.
type Dhcpv6ClientOptionsDuidLlt interface {
	Validation
	// msg marshals Dhcpv6ClientOptionsDuidLlt to protobuf object *otg.Dhcpv6ClientOptionsDuidLlt
	// and doesn't set defaults
	msg() *otg.Dhcpv6ClientOptionsDuidLlt
	// setMsg unmarshals Dhcpv6ClientOptionsDuidLlt from protobuf object *otg.Dhcpv6ClientOptionsDuidLlt
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ClientOptionsDuidLlt) Dhcpv6ClientOptionsDuidLlt
	// provides marshal interface
	Marshal() marshalDhcpv6ClientOptionsDuidLlt
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ClientOptionsDuidLlt
	// validate validates Dhcpv6ClientOptionsDuidLlt
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ClientOptionsDuidLlt, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Time returns uint32, set in Dhcpv6ClientOptionsDuidLlt.
	Time() uint32
	// SetTime assigns uint32 provided by user to Dhcpv6ClientOptionsDuidLlt
	SetTime(value uint32) Dhcpv6ClientOptionsDuidLlt
	// LinkLayerAddress returns Dhcpv6ClientOptionsLinkLayerAddress, set in Dhcpv6ClientOptionsDuidLlt.
	// Dhcpv6ClientOptionsLinkLayerAddress is the link-layer address configured in DUID llt or DUID ll.
	LinkLayerAddress() Dhcpv6ClientOptionsLinkLayerAddress
	// SetLinkLayerAddress assigns Dhcpv6ClientOptionsLinkLayerAddress provided by user to Dhcpv6ClientOptionsDuidLlt.
	// Dhcpv6ClientOptionsLinkLayerAddress is the link-layer address configured in DUID llt or DUID ll.
	SetLinkLayerAddress(value Dhcpv6ClientOptionsLinkLayerAddress) Dhcpv6ClientOptionsDuidLlt
	setNil()
}

// The time value is the time that the DUID is generated represented in seconds since midnight (UTC), January 1,
// 2000, modulo 2^32. The DUID generatation time will the current time when dhcpv6 client contacts the server.
// Time returns a uint32
func (obj *dhcpv6ClientOptionsDuidLlt) Time() uint32 {

	return *obj.obj.Time

}

// The time value is the time that the DUID is generated represented in seconds since midnight (UTC), January 1,
// 2000, modulo 2^32. The DUID generatation time will the current time when dhcpv6 client contacts the server.
// SetTime sets the uint32 value in the Dhcpv6ClientOptionsDuidLlt object
func (obj *dhcpv6ClientOptionsDuidLlt) SetTime(value uint32) Dhcpv6ClientOptionsDuidLlt {

	obj.obj.Time = &value
	return obj
}

// The link-layer address is stored in canonical form, as described in RFC 2464.
// LinkLayerAddress returns a Dhcpv6ClientOptionsLinkLayerAddress
func (obj *dhcpv6ClientOptionsDuidLlt) LinkLayerAddress() Dhcpv6ClientOptionsLinkLayerAddress {
	if obj.obj.LinkLayerAddress == nil {
		obj.obj.LinkLayerAddress = NewDhcpv6ClientOptionsLinkLayerAddress().msg()
	}
	if obj.linkLayerAddressHolder == nil {
		obj.linkLayerAddressHolder = &dhcpv6ClientOptionsLinkLayerAddress{obj: obj.obj.LinkLayerAddress}
	}
	return obj.linkLayerAddressHolder
}

// The link-layer address is stored in canonical form, as described in RFC 2464.
// SetLinkLayerAddress sets the Dhcpv6ClientOptionsLinkLayerAddress value in the Dhcpv6ClientOptionsDuidLlt object
func (obj *dhcpv6ClientOptionsDuidLlt) SetLinkLayerAddress(value Dhcpv6ClientOptionsLinkLayerAddress) Dhcpv6ClientOptionsDuidLlt {

	obj.linkLayerAddressHolder = nil
	obj.obj.LinkLayerAddress = value.msg()

	return obj
}

func (obj *dhcpv6ClientOptionsDuidLlt) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Time is required
	if obj.obj.Time == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Time is required field on interface Dhcpv6ClientOptionsDuidLlt")
	}

	// LinkLayerAddress is required
	if obj.obj.LinkLayerAddress == nil {
		vObj.validationErrors = append(vObj.validationErrors, "LinkLayerAddress is required field on interface Dhcpv6ClientOptionsDuidLlt")
	}

	if obj.obj.LinkLayerAddress != nil {

		obj.LinkLayerAddress().validateObj(vObj, set_default)
	}

}

func (obj *dhcpv6ClientOptionsDuidLlt) setDefault() {

}
