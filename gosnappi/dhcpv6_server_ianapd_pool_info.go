package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ServerIanapdPoolInfo *****
type dhcpv6ServerIanapdPoolInfo struct {
	validation
	obj          *otg.Dhcpv6ServerIanapdPoolInfo
	marshaller   marshalDhcpv6ServerIanapdPoolInfo
	unMarshaller unMarshalDhcpv6ServerIanapdPoolInfo
	ianaHolder   Dhcpv6ServerPoolInfo
	iapdHolder   Dhcpv6ServerIapdPoolInfo
}

func NewDhcpv6ServerIanapdPoolInfo() Dhcpv6ServerIanapdPoolInfo {
	obj := dhcpv6ServerIanapdPoolInfo{obj: &otg.Dhcpv6ServerIanapdPoolInfo{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ServerIanapdPoolInfo) msg() *otg.Dhcpv6ServerIanapdPoolInfo {
	return obj.obj
}

func (obj *dhcpv6ServerIanapdPoolInfo) setMsg(msg *otg.Dhcpv6ServerIanapdPoolInfo) Dhcpv6ServerIanapdPoolInfo {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ServerIanapdPoolInfo struct {
	obj *dhcpv6ServerIanapdPoolInfo
}

type marshalDhcpv6ServerIanapdPoolInfo interface {
	// ToProto marshals Dhcpv6ServerIanapdPoolInfo to protobuf object *otg.Dhcpv6ServerIanapdPoolInfo
	ToProto() (*otg.Dhcpv6ServerIanapdPoolInfo, error)
	// ToPbText marshals Dhcpv6ServerIanapdPoolInfo to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ServerIanapdPoolInfo to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ServerIanapdPoolInfo to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Dhcpv6ServerIanapdPoolInfo to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldhcpv6ServerIanapdPoolInfo struct {
	obj *dhcpv6ServerIanapdPoolInfo
}

type unMarshalDhcpv6ServerIanapdPoolInfo interface {
	// FromProto unmarshals Dhcpv6ServerIanapdPoolInfo from protobuf object *otg.Dhcpv6ServerIanapdPoolInfo
	FromProto(msg *otg.Dhcpv6ServerIanapdPoolInfo) (Dhcpv6ServerIanapdPoolInfo, error)
	// FromPbText unmarshals Dhcpv6ServerIanapdPoolInfo from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ServerIanapdPoolInfo from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ServerIanapdPoolInfo from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ServerIanapdPoolInfo) Marshal() marshalDhcpv6ServerIanapdPoolInfo {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ServerIanapdPoolInfo{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ServerIanapdPoolInfo) Unmarshal() unMarshalDhcpv6ServerIanapdPoolInfo {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ServerIanapdPoolInfo{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ServerIanapdPoolInfo) ToProto() (*otg.Dhcpv6ServerIanapdPoolInfo, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ServerIanapdPoolInfo) FromProto(msg *otg.Dhcpv6ServerIanapdPoolInfo) (Dhcpv6ServerIanapdPoolInfo, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ServerIanapdPoolInfo) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ServerIanapdPoolInfo) FromPbText(value string) error {
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

func (m *marshaldhcpv6ServerIanapdPoolInfo) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ServerIanapdPoolInfo) FromYaml(value string) error {
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

func (m *marshaldhcpv6ServerIanapdPoolInfo) ToJsonRaw() (string, error) {
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

func (m *marshaldhcpv6ServerIanapdPoolInfo) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ServerIanapdPoolInfo) FromJson(value string) error {
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

func (obj *dhcpv6ServerIanapdPoolInfo) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ServerIanapdPoolInfo) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ServerIanapdPoolInfo) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ServerIanapdPoolInfo) Clone() (Dhcpv6ServerIanapdPoolInfo, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ServerIanapdPoolInfo()
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

func (obj *dhcpv6ServerIanapdPoolInfo) setNil() {
	obj.ianaHolder = nil
	obj.iapdHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Dhcpv6ServerIanapdPoolInfo is the container for pool configurations for IA type ianapd.
type Dhcpv6ServerIanapdPoolInfo interface {
	Validation
	// msg marshals Dhcpv6ServerIanapdPoolInfo to protobuf object *otg.Dhcpv6ServerIanapdPoolInfo
	// and doesn't set defaults
	msg() *otg.Dhcpv6ServerIanapdPoolInfo
	// setMsg unmarshals Dhcpv6ServerIanapdPoolInfo from protobuf object *otg.Dhcpv6ServerIanapdPoolInfo
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ServerIanapdPoolInfo) Dhcpv6ServerIanapdPoolInfo
	// provides marshal interface
	Marshal() marshalDhcpv6ServerIanapdPoolInfo
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ServerIanapdPoolInfo
	// validate validates Dhcpv6ServerIanapdPoolInfo
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ServerIanapdPoolInfo, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Iana returns Dhcpv6ServerPoolInfo, set in Dhcpv6ServerIanapdPoolInfo.
	// Dhcpv6ServerPoolInfo is the container for pool configurations for IA types iana and iata.
	Iana() Dhcpv6ServerPoolInfo
	// SetIana assigns Dhcpv6ServerPoolInfo provided by user to Dhcpv6ServerIanapdPoolInfo.
	// Dhcpv6ServerPoolInfo is the container for pool configurations for IA types iana and iata.
	SetIana(value Dhcpv6ServerPoolInfo) Dhcpv6ServerIanapdPoolInfo
	// HasIana checks if Iana has been set in Dhcpv6ServerIanapdPoolInfo
	HasIana() bool
	// Iapd returns Dhcpv6ServerIapdPoolInfo, set in Dhcpv6ServerIanapdPoolInfo.
	// Dhcpv6ServerIapdPoolInfo is the container for prefix pool configurations for IA type iapd.
	Iapd() Dhcpv6ServerIapdPoolInfo
	// SetIapd assigns Dhcpv6ServerIapdPoolInfo provided by user to Dhcpv6ServerIanapdPoolInfo.
	// Dhcpv6ServerIapdPoolInfo is the container for prefix pool configurations for IA type iapd.
	SetIapd(value Dhcpv6ServerIapdPoolInfo) Dhcpv6ServerIanapdPoolInfo
	// HasIapd checks if Iapd has been set in Dhcpv6ServerIanapdPoolInfo
	HasIapd() bool
	setNil()
}

// The pool configurations for IA types iana in ianapd.
// Iana returns a Dhcpv6ServerPoolInfo
func (obj *dhcpv6ServerIanapdPoolInfo) Iana() Dhcpv6ServerPoolInfo {
	if obj.obj.Iana == nil {
		obj.obj.Iana = NewDhcpv6ServerPoolInfo().msg()
	}
	if obj.ianaHolder == nil {
		obj.ianaHolder = &dhcpv6ServerPoolInfo{obj: obj.obj.Iana}
	}
	return obj.ianaHolder
}

// The pool configurations for IA types iana in ianapd.
// Iana returns a Dhcpv6ServerPoolInfo
func (obj *dhcpv6ServerIanapdPoolInfo) HasIana() bool {
	return obj.obj.Iana != nil
}

// The pool configurations for IA types iana in ianapd.
// SetIana sets the Dhcpv6ServerPoolInfo value in the Dhcpv6ServerIanapdPoolInfo object
func (obj *dhcpv6ServerIanapdPoolInfo) SetIana(value Dhcpv6ServerPoolInfo) Dhcpv6ServerIanapdPoolInfo {

	obj.ianaHolder = nil
	obj.obj.Iana = value.msg()

	return obj
}

// The pool configurations for IA types iapd in ianapd.
// Iapd returns a Dhcpv6ServerIapdPoolInfo
func (obj *dhcpv6ServerIanapdPoolInfo) Iapd() Dhcpv6ServerIapdPoolInfo {
	if obj.obj.Iapd == nil {
		obj.obj.Iapd = NewDhcpv6ServerIapdPoolInfo().msg()
	}
	if obj.iapdHolder == nil {
		obj.iapdHolder = &dhcpv6ServerIapdPoolInfo{obj: obj.obj.Iapd}
	}
	return obj.iapdHolder
}

// The pool configurations for IA types iapd in ianapd.
// Iapd returns a Dhcpv6ServerIapdPoolInfo
func (obj *dhcpv6ServerIanapdPoolInfo) HasIapd() bool {
	return obj.obj.Iapd != nil
}

// The pool configurations for IA types iapd in ianapd.
// SetIapd sets the Dhcpv6ServerIapdPoolInfo value in the Dhcpv6ServerIanapdPoolInfo object
func (obj *dhcpv6ServerIanapdPoolInfo) SetIapd(value Dhcpv6ServerIapdPoolInfo) Dhcpv6ServerIanapdPoolInfo {

	obj.iapdHolder = nil
	obj.obj.Iapd = value.msg()

	return obj
}

func (obj *dhcpv6ServerIanapdPoolInfo) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Iana != nil {

		obj.Iana().validateObj(vObj, set_default)
	}

	if obj.obj.Iapd != nil {

		obj.Iapd().validateObj(vObj, set_default)
	}

}

func (obj *dhcpv6ServerIanapdPoolInfo) setDefault() {

}
