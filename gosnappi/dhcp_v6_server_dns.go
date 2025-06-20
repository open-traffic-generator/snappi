package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DhcpV6ServerDns *****
type dhcpV6ServerDns struct {
	validation
	obj                *otg.DhcpV6ServerDns
	marshaller         marshalDhcpV6ServerDns
	unMarshaller       unMarshalDhcpV6ServerDns
	secondaryDnsHolder DhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter
}

func NewDhcpV6ServerDns() DhcpV6ServerDns {
	obj := dhcpV6ServerDns{obj: &otg.DhcpV6ServerDns{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpV6ServerDns) msg() *otg.DhcpV6ServerDns {
	return obj.obj
}

func (obj *dhcpV6ServerDns) setMsg(msg *otg.DhcpV6ServerDns) DhcpV6ServerDns {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpV6ServerDns struct {
	obj *dhcpV6ServerDns
}

type marshalDhcpV6ServerDns interface {
	// ToProto marshals DhcpV6ServerDns to protobuf object *otg.DhcpV6ServerDns
	ToProto() (*otg.DhcpV6ServerDns, error)
	// ToPbText marshals DhcpV6ServerDns to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DhcpV6ServerDns to YAML text
	ToYaml() (string, error)
	// ToJson marshals DhcpV6ServerDns to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpV6ServerDns struct {
	obj *dhcpV6ServerDns
}

type unMarshalDhcpV6ServerDns interface {
	// FromProto unmarshals DhcpV6ServerDns from protobuf object *otg.DhcpV6ServerDns
	FromProto(msg *otg.DhcpV6ServerDns) (DhcpV6ServerDns, error)
	// FromPbText unmarshals DhcpV6ServerDns from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DhcpV6ServerDns from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DhcpV6ServerDns from JSON text
	FromJson(value string) error
}

func (obj *dhcpV6ServerDns) Marshal() marshalDhcpV6ServerDns {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpV6ServerDns{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpV6ServerDns) Unmarshal() unMarshalDhcpV6ServerDns {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpV6ServerDns{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpV6ServerDns) ToProto() (*otg.DhcpV6ServerDns, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpV6ServerDns) FromProto(msg *otg.DhcpV6ServerDns) (DhcpV6ServerDns, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpV6ServerDns) ToPbText() (string, error) {
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

func (m *unMarshaldhcpV6ServerDns) FromPbText(value string) error {
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

func (m *marshaldhcpV6ServerDns) ToYaml() (string, error) {
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

func (m *unMarshaldhcpV6ServerDns) FromYaml(value string) error {
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

func (m *marshaldhcpV6ServerDns) ToJson() (string, error) {
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

func (m *unMarshaldhcpV6ServerDns) FromJson(value string) error {
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

func (obj *dhcpV6ServerDns) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpV6ServerDns) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpV6ServerDns) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpV6ServerDns) Clone() (DhcpV6ServerDns, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpV6ServerDns()
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

func (obj *dhcpV6ServerDns) setNil() {
	obj.secondaryDnsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DhcpV6ServerDns is optional Dns configuration for DHCPv6 server.
type DhcpV6ServerDns interface {
	Validation
	// msg marshals DhcpV6ServerDns to protobuf object *otg.DhcpV6ServerDns
	// and doesn't set defaults
	msg() *otg.DhcpV6ServerDns
	// setMsg unmarshals DhcpV6ServerDns from protobuf object *otg.DhcpV6ServerDns
	// and doesn't set defaults
	setMsg(*otg.DhcpV6ServerDns) DhcpV6ServerDns
	// provides marshal interface
	Marshal() marshalDhcpV6ServerDns
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpV6ServerDns
	// validate validates DhcpV6ServerDns
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DhcpV6ServerDns, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Primary returns string, set in DhcpV6ServerDns.
	Primary() string
	// SetPrimary assigns string provided by user to DhcpV6ServerDns
	SetPrimary(value string) DhcpV6ServerDns
	// SecondaryDns returns DhcpV6ServerDnsDhcpV6ServerSecondaryDnsIterIter, set in DhcpV6ServerDns
	SecondaryDns() DhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter
	setNil()
}

// The primary DNS server address that is offered to DHCP clients that request this information through a TLV option.
// Primary returns a string
func (obj *dhcpV6ServerDns) Primary() string {

	return *obj.obj.Primary

}

// The primary DNS server address that is offered to DHCP clients that request this information through a TLV option.
// SetPrimary sets the string value in the DhcpV6ServerDns object
func (obj *dhcpV6ServerDns) SetPrimary(value string) DhcpV6ServerDns {

	obj.obj.Primary = &value
	return obj
}

// DHCP server secondary dns configuration options. If included secondary DNS server address will be offered to
// DHCP clients that request this information through a TLV option.
// SecondaryDns returns a []DhcpV6ServerSecondaryDns
func (obj *dhcpV6ServerDns) SecondaryDns() DhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter {
	if len(obj.obj.SecondaryDns) == 0 {
		obj.obj.SecondaryDns = []*otg.DhcpV6ServerSecondaryDns{}
	}
	if obj.secondaryDnsHolder == nil {
		obj.secondaryDnsHolder = newDhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter(&obj.obj.SecondaryDns).setMsg(obj)
	}
	return obj.secondaryDnsHolder
}

type dhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter struct {
	obj                           *dhcpV6ServerDns
	dhcpV6ServerSecondaryDnsSlice []DhcpV6ServerSecondaryDns
	fieldPtr                      *[]*otg.DhcpV6ServerSecondaryDns
}

func newDhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter(ptr *[]*otg.DhcpV6ServerSecondaryDns) DhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter {
	return &dhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter{fieldPtr: ptr}
}

type DhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter interface {
	setMsg(*dhcpV6ServerDns) DhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter
	Items() []DhcpV6ServerSecondaryDns
	Add() DhcpV6ServerSecondaryDns
	Append(items ...DhcpV6ServerSecondaryDns) DhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter
	Set(index int, newObj DhcpV6ServerSecondaryDns) DhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter
	Clear() DhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter
	clearHolderSlice() DhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter
	appendHolderSlice(item DhcpV6ServerSecondaryDns) DhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter
}

func (obj *dhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter) setMsg(msg *dhcpV6ServerDns) DhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&dhcpV6ServerSecondaryDns{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *dhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter) Items() []DhcpV6ServerSecondaryDns {
	return obj.dhcpV6ServerSecondaryDnsSlice
}

func (obj *dhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter) Add() DhcpV6ServerSecondaryDns {
	newObj := &otg.DhcpV6ServerSecondaryDns{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &dhcpV6ServerSecondaryDns{obj: newObj}
	newLibObj.setDefault()
	obj.dhcpV6ServerSecondaryDnsSlice = append(obj.dhcpV6ServerSecondaryDnsSlice, newLibObj)
	return newLibObj
}

func (obj *dhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter) Append(items ...DhcpV6ServerSecondaryDns) DhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.dhcpV6ServerSecondaryDnsSlice = append(obj.dhcpV6ServerSecondaryDnsSlice, item)
	}
	return obj
}

func (obj *dhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter) Set(index int, newObj DhcpV6ServerSecondaryDns) DhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.dhcpV6ServerSecondaryDnsSlice[index] = newObj
	return obj
}
func (obj *dhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter) Clear() DhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.DhcpV6ServerSecondaryDns{}
		obj.dhcpV6ServerSecondaryDnsSlice = []DhcpV6ServerSecondaryDns{}
	}
	return obj
}
func (obj *dhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter) clearHolderSlice() DhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter {
	if len(obj.dhcpV6ServerSecondaryDnsSlice) > 0 {
		obj.dhcpV6ServerSecondaryDnsSlice = []DhcpV6ServerSecondaryDns{}
	}
	return obj
}
func (obj *dhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter) appendHolderSlice(item DhcpV6ServerSecondaryDns) DhcpV6ServerDnsDhcpV6ServerSecondaryDnsIter {
	obj.dhcpV6ServerSecondaryDnsSlice = append(obj.dhcpV6ServerSecondaryDnsSlice, item)
	return obj
}

func (obj *dhcpV6ServerDns) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Primary is required
	if obj.obj.Primary == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Primary is required field on interface DhcpV6ServerDns")
	}
	if obj.obj.Primary != nil {

		err := obj.validateIpv6(obj.Primary())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on DhcpV6ServerDns.Primary"))
		}

	}

	if len(obj.obj.SecondaryDns) != 0 {

		if set_default {
			obj.SecondaryDns().clearHolderSlice()
			for _, item := range obj.obj.SecondaryDns {
				obj.SecondaryDns().appendHolderSlice(&dhcpV6ServerSecondaryDns{obj: item})
			}
		}
		for _, item := range obj.SecondaryDns().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *dhcpV6ServerDns) setDefault() {

}
