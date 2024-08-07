package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ServerOptionsBootFileParams *****
type dhcpv6ServerOptionsBootFileParams struct {
	validation
	obj          *otg.Dhcpv6ServerOptionsBootFileParams
	marshaller   marshalDhcpv6ServerOptionsBootFileParams
	unMarshaller unMarshalDhcpv6ServerOptionsBootFileParams
}

func NewDhcpv6ServerOptionsBootFileParams() Dhcpv6ServerOptionsBootFileParams {
	obj := dhcpv6ServerOptionsBootFileParams{obj: &otg.Dhcpv6ServerOptionsBootFileParams{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ServerOptionsBootFileParams) msg() *otg.Dhcpv6ServerOptionsBootFileParams {
	return obj.obj
}

func (obj *dhcpv6ServerOptionsBootFileParams) setMsg(msg *otg.Dhcpv6ServerOptionsBootFileParams) Dhcpv6ServerOptionsBootFileParams {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ServerOptionsBootFileParams struct {
	obj *dhcpv6ServerOptionsBootFileParams
}

type marshalDhcpv6ServerOptionsBootFileParams interface {
	// ToProto marshals Dhcpv6ServerOptionsBootFileParams to protobuf object *otg.Dhcpv6ServerOptionsBootFileParams
	ToProto() (*otg.Dhcpv6ServerOptionsBootFileParams, error)
	// ToPbText marshals Dhcpv6ServerOptionsBootFileParams to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ServerOptionsBootFileParams to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ServerOptionsBootFileParams to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpv6ServerOptionsBootFileParams struct {
	obj *dhcpv6ServerOptionsBootFileParams
}

type unMarshalDhcpv6ServerOptionsBootFileParams interface {
	// FromProto unmarshals Dhcpv6ServerOptionsBootFileParams from protobuf object *otg.Dhcpv6ServerOptionsBootFileParams
	FromProto(msg *otg.Dhcpv6ServerOptionsBootFileParams) (Dhcpv6ServerOptionsBootFileParams, error)
	// FromPbText unmarshals Dhcpv6ServerOptionsBootFileParams from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ServerOptionsBootFileParams from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ServerOptionsBootFileParams from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ServerOptionsBootFileParams) Marshal() marshalDhcpv6ServerOptionsBootFileParams {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ServerOptionsBootFileParams{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ServerOptionsBootFileParams) Unmarshal() unMarshalDhcpv6ServerOptionsBootFileParams {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ServerOptionsBootFileParams{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ServerOptionsBootFileParams) ToProto() (*otg.Dhcpv6ServerOptionsBootFileParams, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ServerOptionsBootFileParams) FromProto(msg *otg.Dhcpv6ServerOptionsBootFileParams) (Dhcpv6ServerOptionsBootFileParams, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ServerOptionsBootFileParams) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ServerOptionsBootFileParams) FromPbText(value string) error {
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

func (m *marshaldhcpv6ServerOptionsBootFileParams) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ServerOptionsBootFileParams) FromYaml(value string) error {
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

func (m *marshaldhcpv6ServerOptionsBootFileParams) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ServerOptionsBootFileParams) FromJson(value string) error {
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

func (obj *dhcpv6ServerOptionsBootFileParams) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ServerOptionsBootFileParams) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ServerOptionsBootFileParams) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ServerOptionsBootFileParams) Clone() (Dhcpv6ServerOptionsBootFileParams, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ServerOptionsBootFileParams()
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

// Dhcpv6ServerOptionsBootFileParams is the option code is 60. They are used to specify parameters for the boot file (similar to the command line arguments in most  modern operating systems).  For example, these parameters could be used to specify the root file system of the OS kernel, or  the location from which a second-stage boot-loader program can download its configuration file.
type Dhcpv6ServerOptionsBootFileParams interface {
	Validation
	// msg marshals Dhcpv6ServerOptionsBootFileParams to protobuf object *otg.Dhcpv6ServerOptionsBootFileParams
	// and doesn't set defaults
	msg() *otg.Dhcpv6ServerOptionsBootFileParams
	// setMsg unmarshals Dhcpv6ServerOptionsBootFileParams from protobuf object *otg.Dhcpv6ServerOptionsBootFileParams
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ServerOptionsBootFileParams) Dhcpv6ServerOptionsBootFileParams
	// provides marshal interface
	Marshal() marshalDhcpv6ServerOptionsBootFileParams
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ServerOptionsBootFileParams
	// validate validates Dhcpv6ServerOptionsBootFileParams
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ServerOptionsBootFileParams, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Parameter returns string, set in Dhcpv6ServerOptionsBootFileParams.
	Parameter() string
	// SetParameter assigns string provided by user to Dhcpv6ServerOptionsBootFileParams
	SetParameter(value string) Dhcpv6ServerOptionsBootFileParams
}

// UTF-8 strings are parameters needed for booting, e.g., kernel parameters.
// Parameter returns a string
func (obj *dhcpv6ServerOptionsBootFileParams) Parameter() string {

	return *obj.obj.Parameter

}

// UTF-8 strings are parameters needed for booting, e.g., kernel parameters.
// SetParameter sets the string value in the Dhcpv6ServerOptionsBootFileParams object
func (obj *dhcpv6ServerOptionsBootFileParams) SetParameter(value string) Dhcpv6ServerOptionsBootFileParams {

	obj.obj.Parameter = &value
	return obj
}

func (obj *dhcpv6ServerOptionsBootFileParams) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Parameter is required
	if obj.obj.Parameter == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Parameter is required field on interface Dhcpv6ServerOptionsBootFileParams")
	}
}

func (obj *dhcpv6ServerOptionsBootFileParams) setDefault() {

}
