package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DhcpV6ServerLease *****
type dhcpV6ServerLease struct {
	validation
	obj          *otg.DhcpV6ServerLease
	marshaller   marshalDhcpV6ServerLease
	unMarshaller unMarshalDhcpV6ServerLease
	iaTypeHolder Dhcpv6ServerIaType
}

func NewDhcpV6ServerLease() DhcpV6ServerLease {
	obj := dhcpV6ServerLease{obj: &otg.DhcpV6ServerLease{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpV6ServerLease) msg() *otg.DhcpV6ServerLease {
	return obj.obj
}

func (obj *dhcpV6ServerLease) setMsg(msg *otg.DhcpV6ServerLease) DhcpV6ServerLease {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpV6ServerLease struct {
	obj *dhcpV6ServerLease
}

type marshalDhcpV6ServerLease interface {
	// ToProto marshals DhcpV6ServerLease to protobuf object *otg.DhcpV6ServerLease
	ToProto() (*otg.DhcpV6ServerLease, error)
	// ToPbText marshals DhcpV6ServerLease to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DhcpV6ServerLease to YAML text
	ToYaml() (string, error)
	// ToJson marshals DhcpV6ServerLease to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpV6ServerLease struct {
	obj *dhcpV6ServerLease
}

type unMarshalDhcpV6ServerLease interface {
	// FromProto unmarshals DhcpV6ServerLease from protobuf object *otg.DhcpV6ServerLease
	FromProto(msg *otg.DhcpV6ServerLease) (DhcpV6ServerLease, error)
	// FromPbText unmarshals DhcpV6ServerLease from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DhcpV6ServerLease from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DhcpV6ServerLease from JSON text
	FromJson(value string) error
}

func (obj *dhcpV6ServerLease) Marshal() marshalDhcpV6ServerLease {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpV6ServerLease{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpV6ServerLease) Unmarshal() unMarshalDhcpV6ServerLease {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpV6ServerLease{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpV6ServerLease) ToProto() (*otg.DhcpV6ServerLease, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpV6ServerLease) FromProto(msg *otg.DhcpV6ServerLease) (DhcpV6ServerLease, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpV6ServerLease) ToPbText() (string, error) {
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

func (m *unMarshaldhcpV6ServerLease) FromPbText(value string) error {
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

func (m *marshaldhcpV6ServerLease) ToYaml() (string, error) {
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

func (m *unMarshaldhcpV6ServerLease) FromYaml(value string) error {
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

func (m *marshaldhcpV6ServerLease) ToJson() (string, error) {
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

func (m *unMarshaldhcpV6ServerLease) FromJson(value string) error {
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

func (obj *dhcpV6ServerLease) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpV6ServerLease) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpV6ServerLease) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpV6ServerLease) Clone() (DhcpV6ServerLease, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpV6ServerLease()
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

func (obj *dhcpV6ServerLease) setNil() {
	obj.iaTypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DhcpV6ServerLease is one DHCP pool configuration on a server.
type DhcpV6ServerLease interface {
	Validation
	// msg marshals DhcpV6ServerLease to protobuf object *otg.DhcpV6ServerLease
	// and doesn't set defaults
	msg() *otg.DhcpV6ServerLease
	// setMsg unmarshals DhcpV6ServerLease from protobuf object *otg.DhcpV6ServerLease
	// and doesn't set defaults
	setMsg(*otg.DhcpV6ServerLease) DhcpV6ServerLease
	// provides marshal interface
	Marshal() marshalDhcpV6ServerLease
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpV6ServerLease
	// validate validates DhcpV6ServerLease
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DhcpV6ServerLease, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LeaseTime returns uint32, set in DhcpV6ServerLease.
	LeaseTime() uint32
	// SetLeaseTime assigns uint32 provided by user to DhcpV6ServerLease
	SetLeaseTime(value uint32) DhcpV6ServerLease
	// HasLeaseTime checks if LeaseTime has been set in DhcpV6ServerLease
	HasLeaseTime() bool
	// IaType returns Dhcpv6ServerIaType, set in DhcpV6ServerLease.
	// Dhcpv6ServerIaType is description is TBD
	IaType() Dhcpv6ServerIaType
	// SetIaType assigns Dhcpv6ServerIaType provided by user to DhcpV6ServerLease.
	// Dhcpv6ServerIaType is description is TBD
	SetIaType(value Dhcpv6ServerIaType) DhcpV6ServerLease
	setNil()
}

// The Life Time length in seconds that is assigned to a lease if the requesting DHCP client does not specify a specific expiration time.
// LeaseTime returns a uint32
func (obj *dhcpV6ServerLease) LeaseTime() uint32 {

	return *obj.obj.LeaseTime

}

// The Life Time length in seconds that is assigned to a lease if the requesting DHCP client does not specify a specific expiration time.
// LeaseTime returns a uint32
func (obj *dhcpV6ServerLease) HasLeaseTime() bool {
	return obj.obj.LeaseTime != nil
}

// The Life Time length in seconds that is assigned to a lease if the requesting DHCP client does not specify a specific expiration time.
// SetLeaseTime sets the uint32 value in the DhcpV6ServerLease object
func (obj *dhcpV6ServerLease) SetLeaseTime(value uint32) DhcpV6ServerLease {

	obj.obj.LeaseTime = &value
	return obj
}

// description is TBD
// IaType returns a Dhcpv6ServerIaType
func (obj *dhcpV6ServerLease) IaType() Dhcpv6ServerIaType {
	if obj.obj.IaType == nil {
		obj.obj.IaType = NewDhcpv6ServerIaType().msg()
	}
	if obj.iaTypeHolder == nil {
		obj.iaTypeHolder = &dhcpv6ServerIaType{obj: obj.obj.IaType}
	}
	return obj.iaTypeHolder
}

// description is TBD
// SetIaType sets the Dhcpv6ServerIaType value in the DhcpV6ServerLease object
func (obj *dhcpV6ServerLease) SetIaType(value Dhcpv6ServerIaType) DhcpV6ServerLease {

	obj.iaTypeHolder = nil
	obj.obj.IaType = value.msg()

	return obj
}

func (obj *dhcpV6ServerLease) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.LeaseTime != nil {

		if *obj.obj.LeaseTime < 300 || *obj.obj.LeaseTime > 30000000 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("300 <= DhcpV6ServerLease.LeaseTime <= 30000000 but Got %d", *obj.obj.LeaseTime))
		}

	}

	// IaType is required
	if obj.obj.IaType == nil {
		vObj.validationErrors = append(vObj.validationErrors, "IaType is required field on interface DhcpV6ServerLease")
	}

	if obj.obj.IaType != nil {

		obj.IaType().validateObj(vObj, set_default)
	}

}

func (obj *dhcpV6ServerLease) setDefault() {
	if obj.obj.LeaseTime == nil {
		obj.SetLeaseTime(86400)
	}

}
