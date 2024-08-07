package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ServerIapdPoolInfo *****
type dhcpv6ServerIapdPoolInfo struct {
	validation
	obj          *otg.Dhcpv6ServerIapdPoolInfo
	marshaller   marshalDhcpv6ServerIapdPoolInfo
	unMarshaller unMarshalDhcpv6ServerIapdPoolInfo
}

func NewDhcpv6ServerIapdPoolInfo() Dhcpv6ServerIapdPoolInfo {
	obj := dhcpv6ServerIapdPoolInfo{obj: &otg.Dhcpv6ServerIapdPoolInfo{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ServerIapdPoolInfo) msg() *otg.Dhcpv6ServerIapdPoolInfo {
	return obj.obj
}

func (obj *dhcpv6ServerIapdPoolInfo) setMsg(msg *otg.Dhcpv6ServerIapdPoolInfo) Dhcpv6ServerIapdPoolInfo {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ServerIapdPoolInfo struct {
	obj *dhcpv6ServerIapdPoolInfo
}

type marshalDhcpv6ServerIapdPoolInfo interface {
	// ToProto marshals Dhcpv6ServerIapdPoolInfo to protobuf object *otg.Dhcpv6ServerIapdPoolInfo
	ToProto() (*otg.Dhcpv6ServerIapdPoolInfo, error)
	// ToPbText marshals Dhcpv6ServerIapdPoolInfo to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ServerIapdPoolInfo to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ServerIapdPoolInfo to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpv6ServerIapdPoolInfo struct {
	obj *dhcpv6ServerIapdPoolInfo
}

type unMarshalDhcpv6ServerIapdPoolInfo interface {
	// FromProto unmarshals Dhcpv6ServerIapdPoolInfo from protobuf object *otg.Dhcpv6ServerIapdPoolInfo
	FromProto(msg *otg.Dhcpv6ServerIapdPoolInfo) (Dhcpv6ServerIapdPoolInfo, error)
	// FromPbText unmarshals Dhcpv6ServerIapdPoolInfo from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ServerIapdPoolInfo from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ServerIapdPoolInfo from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ServerIapdPoolInfo) Marshal() marshalDhcpv6ServerIapdPoolInfo {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ServerIapdPoolInfo{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ServerIapdPoolInfo) Unmarshal() unMarshalDhcpv6ServerIapdPoolInfo {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ServerIapdPoolInfo{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ServerIapdPoolInfo) ToProto() (*otg.Dhcpv6ServerIapdPoolInfo, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ServerIapdPoolInfo) FromProto(msg *otg.Dhcpv6ServerIapdPoolInfo) (Dhcpv6ServerIapdPoolInfo, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ServerIapdPoolInfo) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ServerIapdPoolInfo) FromPbText(value string) error {
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

func (m *marshaldhcpv6ServerIapdPoolInfo) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ServerIapdPoolInfo) FromYaml(value string) error {
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

func (m *marshaldhcpv6ServerIapdPoolInfo) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ServerIapdPoolInfo) FromJson(value string) error {
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

func (obj *dhcpv6ServerIapdPoolInfo) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ServerIapdPoolInfo) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ServerIapdPoolInfo) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ServerIapdPoolInfo) Clone() (Dhcpv6ServerIapdPoolInfo, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ServerIapdPoolInfo()
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

// Dhcpv6ServerIapdPoolInfo is the container for prefix pool configurations for IA type iapd.
type Dhcpv6ServerIapdPoolInfo interface {
	Validation
	// msg marshals Dhcpv6ServerIapdPoolInfo to protobuf object *otg.Dhcpv6ServerIapdPoolInfo
	// and doesn't set defaults
	msg() *otg.Dhcpv6ServerIapdPoolInfo
	// setMsg unmarshals Dhcpv6ServerIapdPoolInfo from protobuf object *otg.Dhcpv6ServerIapdPoolInfo
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ServerIapdPoolInfo) Dhcpv6ServerIapdPoolInfo
	// provides marshal interface
	Marshal() marshalDhcpv6ServerIapdPoolInfo
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ServerIapdPoolInfo
	// validate validates Dhcpv6ServerIapdPoolInfo
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ServerIapdPoolInfo, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// StartPrefixAddress returns string, set in Dhcpv6ServerIapdPoolInfo.
	StartPrefixAddress() string
	// SetStartPrefixAddress assigns string provided by user to Dhcpv6ServerIapdPoolInfo
	SetStartPrefixAddress(value string) Dhcpv6ServerIapdPoolInfo
	// HasStartPrefixAddress checks if StartPrefixAddress has been set in Dhcpv6ServerIapdPoolInfo
	HasStartPrefixAddress() bool
	// Prefix returns uint32, set in Dhcpv6ServerIapdPoolInfo.
	Prefix() uint32
	// SetPrefix assigns uint32 provided by user to Dhcpv6ServerIapdPoolInfo
	SetPrefix(value uint32) Dhcpv6ServerIapdPoolInfo
	// HasPrefix checks if Prefix has been set in Dhcpv6ServerIapdPoolInfo
	HasPrefix() bool
	// Size returns uint32, set in Dhcpv6ServerIapdPoolInfo.
	Size() uint32
	// SetSize assigns uint32 provided by user to Dhcpv6ServerIapdPoolInfo
	SetSize(value uint32) Dhcpv6ServerIapdPoolInfo
	// HasSize checks if Size has been set in Dhcpv6ServerIapdPoolInfo
	HasSize() bool
	// Step returns uint32, set in Dhcpv6ServerIapdPoolInfo.
	Step() uint32
	// SetStep assigns uint32 provided by user to Dhcpv6ServerIapdPoolInfo
	SetStep(value uint32) Dhcpv6ServerIapdPoolInfo
	// HasStep checks if Step has been set in Dhcpv6ServerIapdPoolInfo
	HasStep() bool
}

// The IP address of the first prefix pool.
// StartPrefixAddress returns a string
func (obj *dhcpv6ServerIapdPoolInfo) StartPrefixAddress() string {

	return *obj.obj.StartPrefixAddress

}

// The IP address of the first prefix pool.
// StartPrefixAddress returns a string
func (obj *dhcpv6ServerIapdPoolInfo) HasStartPrefixAddress() bool {
	return obj.obj.StartPrefixAddress != nil
}

// The IP address of the first prefix pool.
// SetStartPrefixAddress sets the string value in the Dhcpv6ServerIapdPoolInfo object
func (obj *dhcpv6ServerIapdPoolInfo) SetStartPrefixAddress(value string) Dhcpv6ServerIapdPoolInfo {

	obj.obj.StartPrefixAddress = &value
	return obj
}

// The IPv6 network prefix length to be applied to the address.
// Prefix returns a uint32
func (obj *dhcpv6ServerIapdPoolInfo) Prefix() uint32 {

	return *obj.obj.Prefix

}

// The IPv6 network prefix length to be applied to the address.
// Prefix returns a uint32
func (obj *dhcpv6ServerIapdPoolInfo) HasPrefix() bool {
	return obj.obj.Prefix != nil
}

// The IPv6 network prefix length to be applied to the address.
// SetPrefix sets the uint32 value in the Dhcpv6ServerIapdPoolInfo object
func (obj *dhcpv6ServerIapdPoolInfo) SetPrefix(value uint32) Dhcpv6ServerIapdPoolInfo {

	obj.obj.Prefix = &value
	return obj
}

// The total number of addresses in the pool.
// Size returns a uint32
func (obj *dhcpv6ServerIapdPoolInfo) Size() uint32 {

	return *obj.obj.Size

}

// The total number of addresses in the pool.
// Size returns a uint32
func (obj *dhcpv6ServerIapdPoolInfo) HasSize() bool {
	return obj.obj.Size != nil
}

// The total number of addresses in the pool.
// SetSize sets the uint32 value in the Dhcpv6ServerIapdPoolInfo object
func (obj *dhcpv6ServerIapdPoolInfo) SetSize(value uint32) Dhcpv6ServerIapdPoolInfo {

	obj.obj.Size = &value
	return obj
}

// The increment value for the lease address within the lease pool where multiple addresses are present. The value is incremented according to the Prefix Length and Step.
// Step returns a uint32
func (obj *dhcpv6ServerIapdPoolInfo) Step() uint32 {

	return *obj.obj.Step

}

// The increment value for the lease address within the lease pool where multiple addresses are present. The value is incremented according to the Prefix Length and Step.
// Step returns a uint32
func (obj *dhcpv6ServerIapdPoolInfo) HasStep() bool {
	return obj.obj.Step != nil
}

// The increment value for the lease address within the lease pool where multiple addresses are present. The value is incremented according to the Prefix Length and Step.
// SetStep sets the uint32 value in the Dhcpv6ServerIapdPoolInfo object
func (obj *dhcpv6ServerIapdPoolInfo) SetStep(value uint32) Dhcpv6ServerIapdPoolInfo {

	obj.obj.Step = &value
	return obj
}

func (obj *dhcpv6ServerIapdPoolInfo) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.StartPrefixAddress != nil {

		err := obj.validateIpv6(obj.StartPrefixAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Dhcpv6ServerIapdPoolInfo.StartPrefixAddress"))
		}

	}

	if obj.obj.Prefix != nil {

		if *obj.obj.Prefix > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Dhcpv6ServerIapdPoolInfo.Prefix <= 128 but Got %d", *obj.obj.Prefix))
		}

	}

	if obj.obj.Size != nil {

		if *obj.obj.Size < 1 || *obj.obj.Size > 1000000 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= Dhcpv6ServerIapdPoolInfo.Size <= 1000000 but Got %d", *obj.obj.Size))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step < 1 || *obj.obj.Step > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= Dhcpv6ServerIapdPoolInfo.Step <= 4294967295 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *dhcpv6ServerIapdPoolInfo) setDefault() {
	if obj.obj.Prefix == nil {
		obj.SetPrefix(64)
	}
	if obj.obj.Size == nil {
		obj.SetSize(10)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}

}
