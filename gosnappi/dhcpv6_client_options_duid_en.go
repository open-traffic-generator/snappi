package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ClientOptionsDuidEn *****
type dhcpv6ClientOptionsDuidEn struct {
	validation
	obj          *otg.Dhcpv6ClientOptionsDuidEn
	marshaller   marshalDhcpv6ClientOptionsDuidEn
	unMarshaller unMarshalDhcpv6ClientOptionsDuidEn
}

func NewDhcpv6ClientOptionsDuidEn() Dhcpv6ClientOptionsDuidEn {
	obj := dhcpv6ClientOptionsDuidEn{obj: &otg.Dhcpv6ClientOptionsDuidEn{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ClientOptionsDuidEn) msg() *otg.Dhcpv6ClientOptionsDuidEn {
	return obj.obj
}

func (obj *dhcpv6ClientOptionsDuidEn) setMsg(msg *otg.Dhcpv6ClientOptionsDuidEn) Dhcpv6ClientOptionsDuidEn {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ClientOptionsDuidEn struct {
	obj *dhcpv6ClientOptionsDuidEn
}

type marshalDhcpv6ClientOptionsDuidEn interface {
	// ToProto marshals Dhcpv6ClientOptionsDuidEn to protobuf object *otg.Dhcpv6ClientOptionsDuidEn
	ToProto() (*otg.Dhcpv6ClientOptionsDuidEn, error)
	// ToPbText marshals Dhcpv6ClientOptionsDuidEn to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ClientOptionsDuidEn to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ClientOptionsDuidEn to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Dhcpv6ClientOptionsDuidEn to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldhcpv6ClientOptionsDuidEn struct {
	obj *dhcpv6ClientOptionsDuidEn
}

type unMarshalDhcpv6ClientOptionsDuidEn interface {
	// FromProto unmarshals Dhcpv6ClientOptionsDuidEn from protobuf object *otg.Dhcpv6ClientOptionsDuidEn
	FromProto(msg *otg.Dhcpv6ClientOptionsDuidEn) (Dhcpv6ClientOptionsDuidEn, error)
	// FromPbText unmarshals Dhcpv6ClientOptionsDuidEn from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ClientOptionsDuidEn from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ClientOptionsDuidEn from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ClientOptionsDuidEn) Marshal() marshalDhcpv6ClientOptionsDuidEn {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ClientOptionsDuidEn{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ClientOptionsDuidEn) Unmarshal() unMarshalDhcpv6ClientOptionsDuidEn {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ClientOptionsDuidEn{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ClientOptionsDuidEn) ToProto() (*otg.Dhcpv6ClientOptionsDuidEn, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ClientOptionsDuidEn) FromProto(msg *otg.Dhcpv6ClientOptionsDuidEn) (Dhcpv6ClientOptionsDuidEn, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ClientOptionsDuidEn) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsDuidEn) FromPbText(value string) error {
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

func (m *marshaldhcpv6ClientOptionsDuidEn) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsDuidEn) FromYaml(value string) error {
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

func (m *marshaldhcpv6ClientOptionsDuidEn) ToJsonRaw() (string, error) {
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

func (m *marshaldhcpv6ClientOptionsDuidEn) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsDuidEn) FromJson(value string) error {
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

func (obj *dhcpv6ClientOptionsDuidEn) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsDuidEn) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsDuidEn) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ClientOptionsDuidEn) Clone() (Dhcpv6ClientOptionsDuidEn, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ClientOptionsDuidEn()
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

// Dhcpv6ClientOptionsDuidEn is dUID assigned by vendor based on enterprise number.
type Dhcpv6ClientOptionsDuidEn interface {
	Validation
	// msg marshals Dhcpv6ClientOptionsDuidEn to protobuf object *otg.Dhcpv6ClientOptionsDuidEn
	// and doesn't set defaults
	msg() *otg.Dhcpv6ClientOptionsDuidEn
	// setMsg unmarshals Dhcpv6ClientOptionsDuidEn from protobuf object *otg.Dhcpv6ClientOptionsDuidEn
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ClientOptionsDuidEn) Dhcpv6ClientOptionsDuidEn
	// provides marshal interface
	Marshal() marshalDhcpv6ClientOptionsDuidEn
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ClientOptionsDuidEn
	// validate validates Dhcpv6ClientOptionsDuidEn
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ClientOptionsDuidEn, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EnterpriseNumber returns uint32, set in Dhcpv6ClientOptionsDuidEn.
	EnterpriseNumber() uint32
	// SetEnterpriseNumber assigns uint32 provided by user to Dhcpv6ClientOptionsDuidEn
	SetEnterpriseNumber(value uint32) Dhcpv6ClientOptionsDuidEn
	// Identifier returns string, set in Dhcpv6ClientOptionsDuidEn.
	Identifier() string
	// SetIdentifier assigns string provided by user to Dhcpv6ClientOptionsDuidEn
	SetIdentifier(value string) Dhcpv6ClientOptionsDuidEn
}

// Vendor's registered private enterprise number as maintained by IANA.
// EnterpriseNumber returns a uint32
func (obj *dhcpv6ClientOptionsDuidEn) EnterpriseNumber() uint32 {

	return *obj.obj.EnterpriseNumber

}

// Vendor's registered private enterprise number as maintained by IANA.
// SetEnterpriseNumber sets the uint32 value in the Dhcpv6ClientOptionsDuidEn object
func (obj *dhcpv6ClientOptionsDuidEn) SetEnterpriseNumber(value uint32) Dhcpv6ClientOptionsDuidEn {

	obj.obj.EnterpriseNumber = &value
	return obj
}

// The unique identifier assigned by the vendor.
// Identifier returns a string
func (obj *dhcpv6ClientOptionsDuidEn) Identifier() string {

	return *obj.obj.Identifier

}

// The unique identifier assigned by the vendor.
// SetIdentifier sets the string value in the Dhcpv6ClientOptionsDuidEn object
func (obj *dhcpv6ClientOptionsDuidEn) SetIdentifier(value string) Dhcpv6ClientOptionsDuidEn {

	obj.obj.Identifier = &value
	return obj
}

func (obj *dhcpv6ClientOptionsDuidEn) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// EnterpriseNumber is required
	if obj.obj.EnterpriseNumber == nil {
		vObj.validationErrors = append(vObj.validationErrors, "EnterpriseNumber is required field on interface Dhcpv6ClientOptionsDuidEn")
	}
	if obj.obj.EnterpriseNumber != nil {

		if *obj.obj.EnterpriseNumber > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Dhcpv6ClientOptionsDuidEn.EnterpriseNumber <= 4294967295 but Got %d", *obj.obj.EnterpriseNumber))
		}

	}

	// Identifier is required
	if obj.obj.Identifier == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Identifier is required field on interface Dhcpv6ClientOptionsDuidEn")
	}
	if obj.obj.Identifier != nil {

		if len(*obj.obj.Identifier) < 1 || len(*obj.obj.Identifier) > 8 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"1 <= length of Dhcpv6ClientOptionsDuidEn.Identifier <= 8 but Got %d",
					len(*obj.obj.Identifier)))
		}

	}

}

func (obj *dhcpv6ClientOptionsDuidEn) setDefault() {

}
