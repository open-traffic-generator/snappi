package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6OptionsVendorSpecificOptions *****
type dhcpv6OptionsVendorSpecificOptions struct {
	validation
	obj          *otg.Dhcpv6OptionsVendorSpecificOptions
	marshaller   marshalDhcpv6OptionsVendorSpecificOptions
	unMarshaller unMarshalDhcpv6OptionsVendorSpecificOptions
}

func NewDhcpv6OptionsVendorSpecificOptions() Dhcpv6OptionsVendorSpecificOptions {
	obj := dhcpv6OptionsVendorSpecificOptions{obj: &otg.Dhcpv6OptionsVendorSpecificOptions{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6OptionsVendorSpecificOptions) msg() *otg.Dhcpv6OptionsVendorSpecificOptions {
	return obj.obj
}

func (obj *dhcpv6OptionsVendorSpecificOptions) setMsg(msg *otg.Dhcpv6OptionsVendorSpecificOptions) Dhcpv6OptionsVendorSpecificOptions {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6OptionsVendorSpecificOptions struct {
	obj *dhcpv6OptionsVendorSpecificOptions
}

type marshalDhcpv6OptionsVendorSpecificOptions interface {
	// ToProto marshals Dhcpv6OptionsVendorSpecificOptions to protobuf object *otg.Dhcpv6OptionsVendorSpecificOptions
	ToProto() (*otg.Dhcpv6OptionsVendorSpecificOptions, error)
	// ToPbText marshals Dhcpv6OptionsVendorSpecificOptions to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6OptionsVendorSpecificOptions to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6OptionsVendorSpecificOptions to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Dhcpv6OptionsVendorSpecificOptions to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldhcpv6OptionsVendorSpecificOptions struct {
	obj *dhcpv6OptionsVendorSpecificOptions
}

type unMarshalDhcpv6OptionsVendorSpecificOptions interface {
	// FromProto unmarshals Dhcpv6OptionsVendorSpecificOptions from protobuf object *otg.Dhcpv6OptionsVendorSpecificOptions
	FromProto(msg *otg.Dhcpv6OptionsVendorSpecificOptions) (Dhcpv6OptionsVendorSpecificOptions, error)
	// FromPbText unmarshals Dhcpv6OptionsVendorSpecificOptions from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6OptionsVendorSpecificOptions from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6OptionsVendorSpecificOptions from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6OptionsVendorSpecificOptions) Marshal() marshalDhcpv6OptionsVendorSpecificOptions {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6OptionsVendorSpecificOptions{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6OptionsVendorSpecificOptions) Unmarshal() unMarshalDhcpv6OptionsVendorSpecificOptions {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6OptionsVendorSpecificOptions{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6OptionsVendorSpecificOptions) ToProto() (*otg.Dhcpv6OptionsVendorSpecificOptions, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6OptionsVendorSpecificOptions) FromProto(msg *otg.Dhcpv6OptionsVendorSpecificOptions) (Dhcpv6OptionsVendorSpecificOptions, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6OptionsVendorSpecificOptions) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6OptionsVendorSpecificOptions) FromPbText(value string) error {
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

func (m *marshaldhcpv6OptionsVendorSpecificOptions) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6OptionsVendorSpecificOptions) FromYaml(value string) error {
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

func (m *marshaldhcpv6OptionsVendorSpecificOptions) ToJsonRaw() (string, error) {
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

func (m *marshaldhcpv6OptionsVendorSpecificOptions) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6OptionsVendorSpecificOptions) FromJson(value string) error {
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

func (obj *dhcpv6OptionsVendorSpecificOptions) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6OptionsVendorSpecificOptions) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6OptionsVendorSpecificOptions) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6OptionsVendorSpecificOptions) Clone() (Dhcpv6OptionsVendorSpecificOptions, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6OptionsVendorSpecificOptions()
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

// Dhcpv6OptionsVendorSpecificOptions is the encapsulated vendor-specific options field is encoded as a sequence of code/length/value fields of  identical format to the DHCP options field. The option codes are defined by the vendor identified in the  enterprise-number field and are not managed by IANA.
type Dhcpv6OptionsVendorSpecificOptions interface {
	Validation
	// msg marshals Dhcpv6OptionsVendorSpecificOptions to protobuf object *otg.Dhcpv6OptionsVendorSpecificOptions
	// and doesn't set defaults
	msg() *otg.Dhcpv6OptionsVendorSpecificOptions
	// setMsg unmarshals Dhcpv6OptionsVendorSpecificOptions from protobuf object *otg.Dhcpv6OptionsVendorSpecificOptions
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6OptionsVendorSpecificOptions) Dhcpv6OptionsVendorSpecificOptions
	// provides marshal interface
	Marshal() marshalDhcpv6OptionsVendorSpecificOptions
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6OptionsVendorSpecificOptions
	// validate validates Dhcpv6OptionsVendorSpecificOptions
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6OptionsVendorSpecificOptions, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Code returns uint32, set in Dhcpv6OptionsVendorSpecificOptions.
	Code() uint32
	// SetCode assigns uint32 provided by user to Dhcpv6OptionsVendorSpecificOptions
	SetCode(value uint32) Dhcpv6OptionsVendorSpecificOptions
	// Data returns string, set in Dhcpv6OptionsVendorSpecificOptions.
	Data() string
	// SetData assigns string provided by user to Dhcpv6OptionsVendorSpecificOptions
	SetData(value string) Dhcpv6OptionsVendorSpecificOptions
}

// The code for the encapsulated option.
// Code returns a uint32
func (obj *dhcpv6OptionsVendorSpecificOptions) Code() uint32 {

	return *obj.obj.Code

}

// The code for the encapsulated option.
// SetCode sets the uint32 value in the Dhcpv6OptionsVendorSpecificOptions object
func (obj *dhcpv6OptionsVendorSpecificOptions) SetCode(value uint32) Dhcpv6OptionsVendorSpecificOptions {

	obj.obj.Code = &value
	return obj
}

// The data for the encapsulated option.
// Data returns a string
func (obj *dhcpv6OptionsVendorSpecificOptions) Data() string {

	return *obj.obj.Data

}

// The data for the encapsulated option.
// SetData sets the string value in the Dhcpv6OptionsVendorSpecificOptions object
func (obj *dhcpv6OptionsVendorSpecificOptions) SetData(value string) Dhcpv6OptionsVendorSpecificOptions {

	obj.obj.Data = &value
	return obj
}

func (obj *dhcpv6OptionsVendorSpecificOptions) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Code is required
	if obj.obj.Code == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Code is required field on interface Dhcpv6OptionsVendorSpecificOptions")
	}
	if obj.obj.Code != nil {

		if *obj.obj.Code > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Dhcpv6OptionsVendorSpecificOptions.Code <= 65535 but Got %d", *obj.obj.Code))
		}

	}

	// Data is required
	if obj.obj.Data == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Data is required field on interface Dhcpv6OptionsVendorSpecificOptions")
	}
}

func (obj *dhcpv6OptionsVendorSpecificOptions) setDefault() {

}
