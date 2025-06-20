package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MACRouteAddress *****
type mACRouteAddress struct {
	validation
	obj          *otg.MACRouteAddress
	marshaller   marshalMACRouteAddress
	unMarshaller unMarshalMACRouteAddress
}

func NewMACRouteAddress() MACRouteAddress {
	obj := mACRouteAddress{obj: &otg.MACRouteAddress{}}
	obj.setDefault()
	return &obj
}

func (obj *mACRouteAddress) msg() *otg.MACRouteAddress {
	return obj.obj
}

func (obj *mACRouteAddress) setMsg(msg *otg.MACRouteAddress) MACRouteAddress {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmACRouteAddress struct {
	obj *mACRouteAddress
}

type marshalMACRouteAddress interface {
	// ToProto marshals MACRouteAddress to protobuf object *otg.MACRouteAddress
	ToProto() (*otg.MACRouteAddress, error)
	// ToPbText marshals MACRouteAddress to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MACRouteAddress to YAML text
	ToYaml() (string, error)
	// ToJson marshals MACRouteAddress to JSON text
	ToJson() (string, error)
}

type unMarshalmACRouteAddress struct {
	obj *mACRouteAddress
}

type unMarshalMACRouteAddress interface {
	// FromProto unmarshals MACRouteAddress from protobuf object *otg.MACRouteAddress
	FromProto(msg *otg.MACRouteAddress) (MACRouteAddress, error)
	// FromPbText unmarshals MACRouteAddress from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MACRouteAddress from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MACRouteAddress from JSON text
	FromJson(value string) error
}

func (obj *mACRouteAddress) Marshal() marshalMACRouteAddress {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmACRouteAddress{obj: obj}
	}
	return obj.marshaller
}

func (obj *mACRouteAddress) Unmarshal() unMarshalMACRouteAddress {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmACRouteAddress{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmACRouteAddress) ToProto() (*otg.MACRouteAddress, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmACRouteAddress) FromProto(msg *otg.MACRouteAddress) (MACRouteAddress, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmACRouteAddress) ToPbText() (string, error) {
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

func (m *unMarshalmACRouteAddress) FromPbText(value string) error {
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

func (m *marshalmACRouteAddress) ToYaml() (string, error) {
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

func (m *unMarshalmACRouteAddress) FromYaml(value string) error {
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

func (m *marshalmACRouteAddress) ToJson() (string, error) {
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

func (m *unMarshalmACRouteAddress) FromJson(value string) error {
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

func (obj *mACRouteAddress) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *mACRouteAddress) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *mACRouteAddress) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *mACRouteAddress) Clone() (MACRouteAddress, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMACRouteAddress()
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

// MACRouteAddress is a container for MAC route addresses.
type MACRouteAddress interface {
	Validation
	// msg marshals MACRouteAddress to protobuf object *otg.MACRouteAddress
	// and doesn't set defaults
	msg() *otg.MACRouteAddress
	// setMsg unmarshals MACRouteAddress from protobuf object *otg.MACRouteAddress
	// and doesn't set defaults
	setMsg(*otg.MACRouteAddress) MACRouteAddress
	// provides marshal interface
	Marshal() marshalMACRouteAddress
	// provides unmarshal interface
	Unmarshal() unMarshalMACRouteAddress
	// validate validates MACRouteAddress
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MACRouteAddress, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Address returns string, set in MACRouteAddress.
	Address() string
	// SetAddress assigns string provided by user to MACRouteAddress
	SetAddress(value string) MACRouteAddress
	// Prefix returns uint32, set in MACRouteAddress.
	Prefix() uint32
	// SetPrefix assigns uint32 provided by user to MACRouteAddress
	SetPrefix(value uint32) MACRouteAddress
	// HasPrefix checks if Prefix has been set in MACRouteAddress
	HasPrefix() bool
	// Count returns uint32, set in MACRouteAddress.
	Count() uint32
	// SetCount assigns uint32 provided by user to MACRouteAddress
	SetCount(value uint32) MACRouteAddress
	// HasCount checks if Count has been set in MACRouteAddress
	HasCount() bool
	// Step returns uint32, set in MACRouteAddress.
	Step() uint32
	// SetStep assigns uint32 provided by user to MACRouteAddress
	SetStep(value uint32) MACRouteAddress
	// HasStep checks if Step has been set in MACRouteAddress
	HasStep() bool
}

// The starting address of the MAC Range.
// Address returns a string
func (obj *mACRouteAddress) Address() string {

	return *obj.obj.Address

}

// The starting address of the MAC Range.
// SetAddress sets the string value in the MACRouteAddress object
func (obj *mACRouteAddress) SetAddress(value string) MACRouteAddress {

	obj.obj.Address = &value
	return obj
}

// The MAC prefix length to be applied to the address.
// Prefix returns a uint32
func (obj *mACRouteAddress) Prefix() uint32 {

	return *obj.obj.Prefix

}

// The MAC prefix length to be applied to the address.
// Prefix returns a uint32
func (obj *mACRouteAddress) HasPrefix() bool {
	return obj.obj.Prefix != nil
}

// The MAC prefix length to be applied to the address.
// SetPrefix sets the uint32 value in the MACRouteAddress object
func (obj *mACRouteAddress) SetPrefix(value uint32) MACRouteAddress {

	obj.obj.Prefix = &value
	return obj
}

// The total number of mac addresses in the range.
// Count returns a uint32
func (obj *mACRouteAddress) Count() uint32 {

	return *obj.obj.Count

}

// The total number of mac addresses in the range.
// Count returns a uint32
func (obj *mACRouteAddress) HasCount() bool {
	return obj.obj.Count != nil
}

// The total number of mac addresses in the range.
// SetCount sets the uint32 value in the MACRouteAddress object
func (obj *mACRouteAddress) SetCount(value uint32) MACRouteAddress {

	obj.obj.Count = &value
	return obj
}

// Increments the mac address prefixes within a mac range  where multiple routes are present.  The value is incremented according to the mac prefix Length and Step.
// Step returns a uint32
func (obj *mACRouteAddress) Step() uint32 {

	return *obj.obj.Step

}

// Increments the mac address prefixes within a mac range  where multiple routes are present.  The value is incremented according to the mac prefix Length and Step.
// Step returns a uint32
func (obj *mACRouteAddress) HasStep() bool {
	return obj.obj.Step != nil
}

// Increments the mac address prefixes within a mac range  where multiple routes are present.  The value is incremented according to the mac prefix Length and Step.
// SetStep sets the uint32 value in the MACRouteAddress object
func (obj *mACRouteAddress) SetStep(value uint32) MACRouteAddress {

	obj.obj.Step = &value
	return obj
}

func (obj *mACRouteAddress) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Address is required
	if obj.obj.Address == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Address is required field on interface MACRouteAddress")
	}
	if obj.obj.Address != nil {

		err := obj.validateMac(obj.Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on MACRouteAddress.Address"))
		}

	}

	if obj.obj.Prefix != nil {

		if *obj.obj.Prefix > 48 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= MACRouteAddress.Prefix <= 48 but Got %d", *obj.obj.Prefix))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count < 1 || *obj.obj.Count > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= MACRouteAddress.Count <= 4294967295 but Got %d", *obj.obj.Count))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step < 1 || *obj.obj.Step > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= MACRouteAddress.Step <= 4294967295 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *mACRouteAddress) setDefault() {
	if obj.obj.Prefix == nil {
		obj.SetPrefix(48)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}

}
