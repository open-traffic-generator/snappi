package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** V6RouteAddress *****
type v6RouteAddress struct {
	validation
	obj          *otg.V6RouteAddress
	marshaller   marshalV6RouteAddress
	unMarshaller unMarshalV6RouteAddress
}

func NewV6RouteAddress() V6RouteAddress {
	obj := v6RouteAddress{obj: &otg.V6RouteAddress{}}
	obj.setDefault()
	return &obj
}

func (obj *v6RouteAddress) msg() *otg.V6RouteAddress {
	return obj.obj
}

func (obj *v6RouteAddress) setMsg(msg *otg.V6RouteAddress) V6RouteAddress {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalv6RouteAddress struct {
	obj *v6RouteAddress
}

type marshalV6RouteAddress interface {
	// ToProto marshals V6RouteAddress to protobuf object *otg.V6RouteAddress
	ToProto() (*otg.V6RouteAddress, error)
	// ToPbText marshals V6RouteAddress to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals V6RouteAddress to YAML text
	ToYaml() (string, error)
	// ToJson marshals V6RouteAddress to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals V6RouteAddress to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalv6RouteAddress struct {
	obj *v6RouteAddress
}

type unMarshalV6RouteAddress interface {
	// FromProto unmarshals V6RouteAddress from protobuf object *otg.V6RouteAddress
	FromProto(msg *otg.V6RouteAddress) (V6RouteAddress, error)
	// FromPbText unmarshals V6RouteAddress from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals V6RouteAddress from YAML text
	FromYaml(value string) error
	// FromJson unmarshals V6RouteAddress from JSON text
	FromJson(value string) error
}

func (obj *v6RouteAddress) Marshal() marshalV6RouteAddress {
	if obj.marshaller == nil {
		obj.marshaller = &marshalv6RouteAddress{obj: obj}
	}
	return obj.marshaller
}

func (obj *v6RouteAddress) Unmarshal() unMarshalV6RouteAddress {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalv6RouteAddress{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalv6RouteAddress) ToProto() (*otg.V6RouteAddress, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalv6RouteAddress) FromProto(msg *otg.V6RouteAddress) (V6RouteAddress, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalv6RouteAddress) ToPbText() (string, error) {
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

func (m *unMarshalv6RouteAddress) FromPbText(value string) error {
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

func (m *marshalv6RouteAddress) ToYaml() (string, error) {
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

func (m *unMarshalv6RouteAddress) FromYaml(value string) error {
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

func (m *marshalv6RouteAddress) ToJsonRaw() (string, error) {
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

func (m *marshalv6RouteAddress) ToJson() (string, error) {
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

func (m *unMarshalv6RouteAddress) FromJson(value string) error {
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

func (obj *v6RouteAddress) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *v6RouteAddress) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *v6RouteAddress) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *v6RouteAddress) Clone() (V6RouteAddress, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewV6RouteAddress()
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

// V6RouteAddress is a container for IPv6 route addresses.
type V6RouteAddress interface {
	Validation
	// msg marshals V6RouteAddress to protobuf object *otg.V6RouteAddress
	// and doesn't set defaults
	msg() *otg.V6RouteAddress
	// setMsg unmarshals V6RouteAddress from protobuf object *otg.V6RouteAddress
	// and doesn't set defaults
	setMsg(*otg.V6RouteAddress) V6RouteAddress
	// provides marshal interface
	Marshal() marshalV6RouteAddress
	// provides unmarshal interface
	Unmarshal() unMarshalV6RouteAddress
	// validate validates V6RouteAddress
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (V6RouteAddress, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Address returns string, set in V6RouteAddress.
	Address() string
	// SetAddress assigns string provided by user to V6RouteAddress
	SetAddress(value string) V6RouteAddress
	// Prefix returns uint32, set in V6RouteAddress.
	Prefix() uint32
	// SetPrefix assigns uint32 provided by user to V6RouteAddress
	SetPrefix(value uint32) V6RouteAddress
	// HasPrefix checks if Prefix has been set in V6RouteAddress
	HasPrefix() bool
	// Count returns uint32, set in V6RouteAddress.
	Count() uint32
	// SetCount assigns uint32 provided by user to V6RouteAddress
	SetCount(value uint32) V6RouteAddress
	// HasCount checks if Count has been set in V6RouteAddress
	HasCount() bool
	// Step returns uint32, set in V6RouteAddress.
	Step() uint32
	// SetStep assigns uint32 provided by user to V6RouteAddress
	SetStep(value uint32) V6RouteAddress
	// HasStep checks if Step has been set in V6RouteAddress
	HasStep() bool
}

// The starting address of the network.
// Address returns a string
func (obj *v6RouteAddress) Address() string {

	return *obj.obj.Address

}

// The starting address of the network.
// SetAddress sets the string value in the V6RouteAddress object
func (obj *v6RouteAddress) SetAddress(value string) V6RouteAddress {

	obj.obj.Address = &value
	return obj
}

// The IPv6 network prefix length to be applied to the address.
// Prefix returns a uint32
func (obj *v6RouteAddress) Prefix() uint32 {

	return *obj.obj.Prefix

}

// The IPv6 network prefix length to be applied to the address.
// Prefix returns a uint32
func (obj *v6RouteAddress) HasPrefix() bool {
	return obj.obj.Prefix != nil
}

// The IPv6 network prefix length to be applied to the address.
// SetPrefix sets the uint32 value in the V6RouteAddress object
func (obj *v6RouteAddress) SetPrefix(value uint32) V6RouteAddress {

	obj.obj.Prefix = &value
	return obj
}

// The total number of addresses in the range.
// Count returns a uint32
func (obj *v6RouteAddress) Count() uint32 {

	return *obj.obj.Count

}

// The total number of addresses in the range.
// Count returns a uint32
func (obj *v6RouteAddress) HasCount() bool {
	return obj.obj.Count != nil
}

// The total number of addresses in the range.
// SetCount sets the uint32 value in the V6RouteAddress object
func (obj *v6RouteAddress) SetCount(value uint32) V6RouteAddress {

	obj.obj.Count = &value
	return obj
}

// Increments the network address prefixes within a route range  where multiple routes are present.  The value is incremented according to the Prefix Length and Step.
// Step returns a uint32
func (obj *v6RouteAddress) Step() uint32 {

	return *obj.obj.Step

}

// Increments the network address prefixes within a route range  where multiple routes are present.  The value is incremented according to the Prefix Length and Step.
// Step returns a uint32
func (obj *v6RouteAddress) HasStep() bool {
	return obj.obj.Step != nil
}

// Increments the network address prefixes within a route range  where multiple routes are present.  The value is incremented according to the Prefix Length and Step.
// SetStep sets the uint32 value in the V6RouteAddress object
func (obj *v6RouteAddress) SetStep(value uint32) V6RouteAddress {

	obj.obj.Step = &value
	return obj
}

func (obj *v6RouteAddress) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Address is required
	if obj.obj.Address == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Address is required field on interface V6RouteAddress")
	}
	if obj.obj.Address != nil {

		err := obj.validateIpv6(obj.Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on V6RouteAddress.Address"))
		}

	}

	if obj.obj.Prefix != nil {

		if *obj.obj.Prefix > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= V6RouteAddress.Prefix <= 128 but Got %d", *obj.obj.Prefix))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count < 1 || *obj.obj.Count > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= V6RouteAddress.Count <= 4294967295 but Got %d", *obj.obj.Count))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step < 1 || *obj.obj.Step > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= V6RouteAddress.Step <= 4294967295 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *v6RouteAddress) setDefault() {
	if obj.obj.Prefix == nil {
		obj.SetPrefix(64)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}

}
