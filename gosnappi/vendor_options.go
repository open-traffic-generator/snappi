package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** VendorOptions *****
type vendorOptions struct {
	validation
	obj          *otg.VendorOptions
	marshaller   marshalVendorOptions
	unMarshaller unMarshalVendorOptions
}

func NewVendorOptions() VendorOptions {
	obj := vendorOptions{obj: &otg.VendorOptions{}}
	obj.setDefault()
	return &obj
}

func (obj *vendorOptions) msg() *otg.VendorOptions {
	return obj.obj
}

func (obj *vendorOptions) setMsg(msg *otg.VendorOptions) VendorOptions {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalvendorOptions struct {
	obj *vendorOptions
}

type marshalVendorOptions interface {
	// ToProto marshals VendorOptions to protobuf object *otg.VendorOptions
	ToProto() (*otg.VendorOptions, error)
	// ToPbText marshals VendorOptions to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals VendorOptions to YAML text
	ToYaml() (string, error)
	// ToJson marshals VendorOptions to JSON text
	ToJson() (string, error)
}

type unMarshalvendorOptions struct {
	obj *vendorOptions
}

type unMarshalVendorOptions interface {
	// FromProto unmarshals VendorOptions from protobuf object *otg.VendorOptions
	FromProto(msg *otg.VendorOptions) (VendorOptions, error)
	// FromPbText unmarshals VendorOptions from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals VendorOptions from YAML text
	FromYaml(value string) error
	// FromJson unmarshals VendorOptions from JSON text
	FromJson(value string) error
}

func (obj *vendorOptions) Marshal() marshalVendorOptions {
	if obj.marshaller == nil {
		obj.marshaller = &marshalvendorOptions{obj: obj}
	}
	return obj.marshaller
}

func (obj *vendorOptions) Unmarshal() unMarshalVendorOptions {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalvendorOptions{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalvendorOptions) ToProto() (*otg.VendorOptions, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalvendorOptions) FromProto(msg *otg.VendorOptions) (VendorOptions, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalvendorOptions) ToPbText() (string, error) {
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

func (m *unMarshalvendorOptions) FromPbText(value string) error {
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

func (m *marshalvendorOptions) ToYaml() (string, error) {
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

func (m *unMarshalvendorOptions) FromYaml(value string) error {
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

func (m *marshalvendorOptions) ToJson() (string, error) {
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

func (m *unMarshalvendorOptions) FromJson(value string) error {
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

func (obj *vendorOptions) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *vendorOptions) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *vendorOptions) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *vendorOptions) Clone() (VendorOptions, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewVendorOptions()
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

// VendorOptions is description is TBD
type VendorOptions interface {
	Validation
	// msg marshals VendorOptions to protobuf object *otg.VendorOptions
	// and doesn't set defaults
	msg() *otg.VendorOptions
	// setMsg unmarshals VendorOptions from protobuf object *otg.VendorOptions
	// and doesn't set defaults
	setMsg(*otg.VendorOptions) VendorOptions
	// provides marshal interface
	Marshal() marshalVendorOptions
	// provides unmarshal interface
	Unmarshal() unMarshalVendorOptions
	// validate validates VendorOptions
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (VendorOptions, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// AdditionalFeatures returns []string, set in VendorOptions.
	AdditionalFeatures() []string
	// SetAdditionalFeatures assigns []string provided by user to VendorOptions
	SetAdditionalFeatures(value []string) VendorOptions
	// Deviations returns []string, set in VendorOptions.
	Deviations() []string
	// SetDeviations assigns []string provided by user to VendorOptions
	SetDeviations(value []string) VendorOptions
	// Vendor returns string, set in VendorOptions.
	Vendor() string
	// SetVendor assigns string provided by user to VendorOptions
	SetVendor(value string) VendorOptions
	// HasVendor checks if Vendor has been set in VendorOptions
	HasVendor() bool
}

// OTG Application require some additional property to run the config. That additional property may not be generic enough to address in OTG model. So, OTG expose this backdoor access to define a feature name must be understandable by OTG application.
// AdditionalFeatures returns a []string
func (obj *vendorOptions) AdditionalFeatures() []string {
	if obj.obj.AdditionalFeatures == nil {
		obj.obj.AdditionalFeatures = make([]string, 0)
	}
	return obj.obj.AdditionalFeatures
}

// OTG Application require some additional property to run the config. That additional property may not be generic enough to address in OTG model. So, OTG expose this backdoor access to define a feature name must be understandable by OTG application.
// SetAdditionalFeatures sets the []string value in the VendorOptions object
func (obj *vendorOptions) SetAdditionalFeatures(value []string) VendorOptions {

	if obj.obj.AdditionalFeatures == nil {
		obj.obj.AdditionalFeatures = make([]string, 0)
	}
	obj.obj.AdditionalFeatures = value

	return obj
}

// Name of the deviating feature names when OTG model deviates from the application requirements.
// Deviations returns a []string
func (obj *vendorOptions) Deviations() []string {
	if obj.obj.Deviations == nil {
		obj.obj.Deviations = make([]string, 0)
	}
	return obj.obj.Deviations
}

// Name of the deviating feature names when OTG model deviates from the application requirements.
// SetDeviations sets the []string value in the VendorOptions object
func (obj *vendorOptions) SetDeviations(value []string) VendorOptions {

	if obj.obj.Deviations == nil {
		obj.obj.Deviations = make([]string, 0)
	}
	obj.obj.Deviations = value

	return obj
}

// Names of the vendor.  Vendor specific options will be applicable for all vendor if this field is not defined.
// Vendor returns a string
func (obj *vendorOptions) Vendor() string {

	return *obj.obj.Vendor

}

// Names of the vendor.  Vendor specific options will be applicable for all vendor if this field is not defined.
// Vendor returns a string
func (obj *vendorOptions) HasVendor() bool {
	return obj.obj.Vendor != nil
}

// Names of the vendor.  Vendor specific options will be applicable for all vendor if this field is not defined.
// SetVendor sets the string value in the VendorOptions object
func (obj *vendorOptions) SetVendor(value string) VendorOptions {

	obj.obj.Vendor = &value
	return obj
}

func (obj *vendorOptions) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *vendorOptions) setDefault() {

}
