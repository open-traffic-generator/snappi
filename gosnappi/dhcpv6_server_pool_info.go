package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ServerPoolInfo *****
type dhcpv6ServerPoolInfo struct {
	validation
	obj          *otg.Dhcpv6ServerPoolInfo
	marshaller   marshalDhcpv6ServerPoolInfo
	unMarshaller unMarshalDhcpv6ServerPoolInfo
}

func NewDhcpv6ServerPoolInfo() Dhcpv6ServerPoolInfo {
	obj := dhcpv6ServerPoolInfo{obj: &otg.Dhcpv6ServerPoolInfo{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ServerPoolInfo) msg() *otg.Dhcpv6ServerPoolInfo {
	return obj.obj
}

func (obj *dhcpv6ServerPoolInfo) setMsg(msg *otg.Dhcpv6ServerPoolInfo) Dhcpv6ServerPoolInfo {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ServerPoolInfo struct {
	obj *dhcpv6ServerPoolInfo
}

type marshalDhcpv6ServerPoolInfo interface {
	// ToProto marshals Dhcpv6ServerPoolInfo to protobuf object *otg.Dhcpv6ServerPoolInfo
	ToProto() (*otg.Dhcpv6ServerPoolInfo, error)
	// ToPbText marshals Dhcpv6ServerPoolInfo to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ServerPoolInfo to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ServerPoolInfo to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpv6ServerPoolInfo struct {
	obj *dhcpv6ServerPoolInfo
}

type unMarshalDhcpv6ServerPoolInfo interface {
	// FromProto unmarshals Dhcpv6ServerPoolInfo from protobuf object *otg.Dhcpv6ServerPoolInfo
	FromProto(msg *otg.Dhcpv6ServerPoolInfo) (Dhcpv6ServerPoolInfo, error)
	// FromPbText unmarshals Dhcpv6ServerPoolInfo from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ServerPoolInfo from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ServerPoolInfo from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ServerPoolInfo) Marshal() marshalDhcpv6ServerPoolInfo {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ServerPoolInfo{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ServerPoolInfo) Unmarshal() unMarshalDhcpv6ServerPoolInfo {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ServerPoolInfo{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ServerPoolInfo) ToProto() (*otg.Dhcpv6ServerPoolInfo, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ServerPoolInfo) FromProto(msg *otg.Dhcpv6ServerPoolInfo) (Dhcpv6ServerPoolInfo, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ServerPoolInfo) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ServerPoolInfo) FromPbText(value string) error {
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

func (m *marshaldhcpv6ServerPoolInfo) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ServerPoolInfo) FromYaml(value string) error {
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

func (m *marshaldhcpv6ServerPoolInfo) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ServerPoolInfo) FromJson(value string) error {
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

func (obj *dhcpv6ServerPoolInfo) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ServerPoolInfo) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ServerPoolInfo) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ServerPoolInfo) Clone() (Dhcpv6ServerPoolInfo, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ServerPoolInfo()
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

// Dhcpv6ServerPoolInfo is the container for pool configurations for IA types iana and iata.
type Dhcpv6ServerPoolInfo interface {
	Validation
	// msg marshals Dhcpv6ServerPoolInfo to protobuf object *otg.Dhcpv6ServerPoolInfo
	// and doesn't set defaults
	msg() *otg.Dhcpv6ServerPoolInfo
	// setMsg unmarshals Dhcpv6ServerPoolInfo from protobuf object *otg.Dhcpv6ServerPoolInfo
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ServerPoolInfo) Dhcpv6ServerPoolInfo
	// provides marshal interface
	Marshal() marshalDhcpv6ServerPoolInfo
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ServerPoolInfo
	// validate validates Dhcpv6ServerPoolInfo
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ServerPoolInfo, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// StartAddress returns string, set in Dhcpv6ServerPoolInfo.
	StartAddress() string
	// SetStartAddress assigns string provided by user to Dhcpv6ServerPoolInfo
	SetStartAddress(value string) Dhcpv6ServerPoolInfo
	// HasStartAddress checks if StartAddress has been set in Dhcpv6ServerPoolInfo
	HasStartAddress() bool
	// PrefixLen returns uint32, set in Dhcpv6ServerPoolInfo.
	PrefixLen() uint32
	// SetPrefixLen assigns uint32 provided by user to Dhcpv6ServerPoolInfo
	SetPrefixLen(value uint32) Dhcpv6ServerPoolInfo
	// HasPrefixLen checks if PrefixLen has been set in Dhcpv6ServerPoolInfo
	HasPrefixLen() bool
	// Size returns uint32, set in Dhcpv6ServerPoolInfo.
	Size() uint32
	// SetSize assigns uint32 provided by user to Dhcpv6ServerPoolInfo
	SetSize(value uint32) Dhcpv6ServerPoolInfo
	// HasSize checks if Size has been set in Dhcpv6ServerPoolInfo
	HasSize() bool
	// Step returns uint32, set in Dhcpv6ServerPoolInfo.
	Step() uint32
	// SetStep assigns uint32 provided by user to Dhcpv6ServerPoolInfo
	SetStep(value uint32) Dhcpv6ServerPoolInfo
	// HasStep checks if Step has been set in Dhcpv6ServerPoolInfo
	HasStep() bool
}

// The first IP address of the lease pool.
// StartAddress returns a string
func (obj *dhcpv6ServerPoolInfo) StartAddress() string {

	return *obj.obj.StartAddress

}

// The first IP address of the lease pool.
// StartAddress returns a string
func (obj *dhcpv6ServerPoolInfo) HasStartAddress() bool {
	return obj.obj.StartAddress != nil
}

// The first IP address of the lease pool.
// SetStartAddress sets the string value in the Dhcpv6ServerPoolInfo object
func (obj *dhcpv6ServerPoolInfo) SetStartAddress(value string) Dhcpv6ServerPoolInfo {

	obj.obj.StartAddress = &value
	return obj
}

// The IPv6 network prefix length is used for incrementing the lease address within the lease pool where multiple addresses are configured by using the size field. The address is incremented using the configured Prefix Length and Step.
// PrefixLen returns a uint32
func (obj *dhcpv6ServerPoolInfo) PrefixLen() uint32 {

	return *obj.obj.PrefixLen

}

// The IPv6 network prefix length is used for incrementing the lease address within the lease pool where multiple addresses are configured by using the size field. The address is incremented using the configured Prefix Length and Step.
// PrefixLen returns a uint32
func (obj *dhcpv6ServerPoolInfo) HasPrefixLen() bool {
	return obj.obj.PrefixLen != nil
}

// The IPv6 network prefix length is used for incrementing the lease address within the lease pool where multiple addresses are configured by using the size field. The address is incremented using the configured Prefix Length and Step.
// SetPrefixLen sets the uint32 value in the Dhcpv6ServerPoolInfo object
func (obj *dhcpv6ServerPoolInfo) SetPrefixLen(value uint32) Dhcpv6ServerPoolInfo {

	obj.obj.PrefixLen = &value
	return obj
}

// The total number of addresses in the pool.
// Size returns a uint32
func (obj *dhcpv6ServerPoolInfo) Size() uint32 {

	return *obj.obj.Size

}

// The total number of addresses in the pool.
// Size returns a uint32
func (obj *dhcpv6ServerPoolInfo) HasSize() bool {
	return obj.obj.Size != nil
}

// The total number of addresses in the pool.
// SetSize sets the uint32 value in the Dhcpv6ServerPoolInfo object
func (obj *dhcpv6ServerPoolInfo) SetSize(value uint32) Dhcpv6ServerPoolInfo {

	obj.obj.Size = &value
	return obj
}

// The increment value for the lease address within the lease pool where multiple addresses are present. The value is incremented according to the configured Prefix Length and Step.
// Step returns a uint32
func (obj *dhcpv6ServerPoolInfo) Step() uint32 {

	return *obj.obj.Step

}

// The increment value for the lease address within the lease pool where multiple addresses are present. The value is incremented according to the configured Prefix Length and Step.
// Step returns a uint32
func (obj *dhcpv6ServerPoolInfo) HasStep() bool {
	return obj.obj.Step != nil
}

// The increment value for the lease address within the lease pool where multiple addresses are present. The value is incremented according to the configured Prefix Length and Step.
// SetStep sets the uint32 value in the Dhcpv6ServerPoolInfo object
func (obj *dhcpv6ServerPoolInfo) SetStep(value uint32) Dhcpv6ServerPoolInfo {

	obj.obj.Step = &value
	return obj
}

func (obj *dhcpv6ServerPoolInfo) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.StartAddress != nil {

		err := obj.validateIpv6(obj.StartAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Dhcpv6ServerPoolInfo.StartAddress"))
		}

	}

	if obj.obj.PrefixLen != nil {

		if *obj.obj.PrefixLen > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Dhcpv6ServerPoolInfo.PrefixLen <= 128 but Got %d", *obj.obj.PrefixLen))
		}

	}

	if obj.obj.Size != nil {

		if *obj.obj.Size < 1 || *obj.obj.Size > 1000000 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= Dhcpv6ServerPoolInfo.Size <= 1000000 but Got %d", *obj.obj.Size))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step < 1 || *obj.obj.Step > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= Dhcpv6ServerPoolInfo.Step <= 4294967295 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *dhcpv6ServerPoolInfo) setDefault() {
	if obj.obj.PrefixLen == nil {
		obj.SetPrefixLen(64)
	}
	if obj.obj.Size == nil {
		obj.SetSize(1)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}

}
