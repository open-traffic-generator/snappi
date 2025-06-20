package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** V4RouteAddress *****
type v4RouteAddress struct {
	validation
	obj          *otg.V4RouteAddress
	marshaller   marshalV4RouteAddress
	unMarshaller unMarshalV4RouteAddress
}

func NewV4RouteAddress() V4RouteAddress {
	obj := v4RouteAddress{obj: &otg.V4RouteAddress{}}
	obj.setDefault()
	return &obj
}

func (obj *v4RouteAddress) msg() *otg.V4RouteAddress {
	return obj.obj
}

func (obj *v4RouteAddress) setMsg(msg *otg.V4RouteAddress) V4RouteAddress {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalv4RouteAddress struct {
	obj *v4RouteAddress
}

type marshalV4RouteAddress interface {
	// ToProto marshals V4RouteAddress to protobuf object *otg.V4RouteAddress
	ToProto() (*otg.V4RouteAddress, error)
	// ToPbText marshals V4RouteAddress to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals V4RouteAddress to YAML text
	ToYaml() (string, error)
	// ToJson marshals V4RouteAddress to JSON text
	ToJson() (string, error)
}

type unMarshalv4RouteAddress struct {
	obj *v4RouteAddress
}

type unMarshalV4RouteAddress interface {
	// FromProto unmarshals V4RouteAddress from protobuf object *otg.V4RouteAddress
	FromProto(msg *otg.V4RouteAddress) (V4RouteAddress, error)
	// FromPbText unmarshals V4RouteAddress from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals V4RouteAddress from YAML text
	FromYaml(value string) error
	// FromJson unmarshals V4RouteAddress from JSON text
	FromJson(value string) error
}

func (obj *v4RouteAddress) Marshal() marshalV4RouteAddress {
	if obj.marshaller == nil {
		obj.marshaller = &marshalv4RouteAddress{obj: obj}
	}
	return obj.marshaller
}

func (obj *v4RouteAddress) Unmarshal() unMarshalV4RouteAddress {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalv4RouteAddress{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalv4RouteAddress) ToProto() (*otg.V4RouteAddress, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalv4RouteAddress) FromProto(msg *otg.V4RouteAddress) (V4RouteAddress, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalv4RouteAddress) ToPbText() (string, error) {
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

func (m *unMarshalv4RouteAddress) FromPbText(value string) error {
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

func (m *marshalv4RouteAddress) ToYaml() (string, error) {
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

func (m *unMarshalv4RouteAddress) FromYaml(value string) error {
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

func (m *marshalv4RouteAddress) ToJson() (string, error) {
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

func (m *unMarshalv4RouteAddress) FromJson(value string) error {
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

func (obj *v4RouteAddress) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *v4RouteAddress) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *v4RouteAddress) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *v4RouteAddress) Clone() (V4RouteAddress, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewV4RouteAddress()
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

// V4RouteAddress is a container for IPv4 route addresses.
type V4RouteAddress interface {
	Validation
	// msg marshals V4RouteAddress to protobuf object *otg.V4RouteAddress
	// and doesn't set defaults
	msg() *otg.V4RouteAddress
	// setMsg unmarshals V4RouteAddress from protobuf object *otg.V4RouteAddress
	// and doesn't set defaults
	setMsg(*otg.V4RouteAddress) V4RouteAddress
	// provides marshal interface
	Marshal() marshalV4RouteAddress
	// provides unmarshal interface
	Unmarshal() unMarshalV4RouteAddress
	// validate validates V4RouteAddress
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (V4RouteAddress, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Address returns string, set in V4RouteAddress.
	Address() string
	// SetAddress assigns string provided by user to V4RouteAddress
	SetAddress(value string) V4RouteAddress
	// Prefix returns uint32, set in V4RouteAddress.
	Prefix() uint32
	// SetPrefix assigns uint32 provided by user to V4RouteAddress
	SetPrefix(value uint32) V4RouteAddress
	// HasPrefix checks if Prefix has been set in V4RouteAddress
	HasPrefix() bool
	// Count returns uint32, set in V4RouteAddress.
	Count() uint32
	// SetCount assigns uint32 provided by user to V4RouteAddress
	SetCount(value uint32) V4RouteAddress
	// HasCount checks if Count has been set in V4RouteAddress
	HasCount() bool
	// Step returns uint32, set in V4RouteAddress.
	Step() uint32
	// SetStep assigns uint32 provided by user to V4RouteAddress
	SetStep(value uint32) V4RouteAddress
	// HasStep checks if Step has been set in V4RouteAddress
	HasStep() bool
}

// The starting address of the network.
// Address returns a string
func (obj *v4RouteAddress) Address() string {

	return *obj.obj.Address

}

// The starting address of the network.
// SetAddress sets the string value in the V4RouteAddress object
func (obj *v4RouteAddress) SetAddress(value string) V4RouteAddress {

	obj.obj.Address = &value
	return obj
}

// The IPv4 network prefix length to be applied to the address.
// Prefix returns a uint32
func (obj *v4RouteAddress) Prefix() uint32 {

	return *obj.obj.Prefix

}

// The IPv4 network prefix length to be applied to the address.
// Prefix returns a uint32
func (obj *v4RouteAddress) HasPrefix() bool {
	return obj.obj.Prefix != nil
}

// The IPv4 network prefix length to be applied to the address.
// SetPrefix sets the uint32 value in the V4RouteAddress object
func (obj *v4RouteAddress) SetPrefix(value uint32) V4RouteAddress {

	obj.obj.Prefix = &value
	return obj
}

// The total number of addresses in the range.
// Count returns a uint32
func (obj *v4RouteAddress) Count() uint32 {

	return *obj.obj.Count

}

// The total number of addresses in the range.
// Count returns a uint32
func (obj *v4RouteAddress) HasCount() bool {
	return obj.obj.Count != nil
}

// The total number of addresses in the range.
// SetCount sets the uint32 value in the V4RouteAddress object
func (obj *v4RouteAddress) SetCount(value uint32) V4RouteAddress {

	obj.obj.Count = &value
	return obj
}

// Increments the network address prefixes within a route range  where multiple routes are present.  The value is incremented according to the Prefix Length and Step.
// Step returns a uint32
func (obj *v4RouteAddress) Step() uint32 {

	return *obj.obj.Step

}

// Increments the network address prefixes within a route range  where multiple routes are present.  The value is incremented according to the Prefix Length and Step.
// Step returns a uint32
func (obj *v4RouteAddress) HasStep() bool {
	return obj.obj.Step != nil
}

// Increments the network address prefixes within a route range  where multiple routes are present.  The value is incremented according to the Prefix Length and Step.
// SetStep sets the uint32 value in the V4RouteAddress object
func (obj *v4RouteAddress) SetStep(value uint32) V4RouteAddress {

	obj.obj.Step = &value
	return obj
}

func (obj *v4RouteAddress) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Address is required
	if obj.obj.Address == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Address is required field on interface V4RouteAddress")
	}
	if obj.obj.Address != nil {

		err := obj.validateIpv4(obj.Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on V4RouteAddress.Address"))
		}

	}

	if obj.obj.Prefix != nil {

		if *obj.obj.Prefix > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= V4RouteAddress.Prefix <= 32 but Got %d", *obj.obj.Prefix))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count < 1 || *obj.obj.Count > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= V4RouteAddress.Count <= 4294967295 but Got %d", *obj.obj.Count))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step < 1 || *obj.obj.Step > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= V4RouteAddress.Step <= 4294967295 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *v4RouteAddress) setDefault() {
	if obj.obj.Prefix == nil {
		obj.SetPrefix(24)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}

}
