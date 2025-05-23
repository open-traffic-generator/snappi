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
	// ToJsonRaw marshals Dhcpv6ServerIapdPoolInfo to raw JSON text
	ToJsonRaw() (string, error)
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

func (m *marshaldhcpv6ServerIapdPoolInfo) ToJsonRaw() (string, error) {
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
	// ConfiguredPrefixLen returns uint32, set in Dhcpv6ServerIapdPoolInfo.
	ConfiguredPrefixLen() uint32
	// SetConfiguredPrefixLen assigns uint32 provided by user to Dhcpv6ServerIapdPoolInfo
	SetConfiguredPrefixLen(value uint32) Dhcpv6ServerIapdPoolInfo
	// HasConfiguredPrefixLen checks if ConfiguredPrefixLen has been set in Dhcpv6ServerIapdPoolInfo
	HasConfiguredPrefixLen() bool
	// PrefixSize returns uint32, set in Dhcpv6ServerIapdPoolInfo.
	PrefixSize() uint32
	// SetPrefixSize assigns uint32 provided by user to Dhcpv6ServerIapdPoolInfo
	SetPrefixSize(value uint32) Dhcpv6ServerIapdPoolInfo
	// HasPrefixSize checks if PrefixSize has been set in Dhcpv6ServerIapdPoolInfo
	HasPrefixSize() bool
	// PrefixStep returns uint32, set in Dhcpv6ServerIapdPoolInfo.
	PrefixStep() uint32
	// SetPrefixStep assigns uint32 provided by user to Dhcpv6ServerIapdPoolInfo
	SetPrefixStep(value uint32) Dhcpv6ServerIapdPoolInfo
	// HasPrefixStep checks if PrefixStep has been set in Dhcpv6ServerIapdPoolInfo
	HasPrefixStep() bool
	// AdvertisedPrefixLen returns uint32, set in Dhcpv6ServerIapdPoolInfo.
	AdvertisedPrefixLen() uint32
	// SetAdvertisedPrefixLen assigns uint32 provided by user to Dhcpv6ServerIapdPoolInfo
	SetAdvertisedPrefixLen(value uint32) Dhcpv6ServerIapdPoolInfo
	// HasAdvertisedPrefixLen checks if AdvertisedPrefixLen has been set in Dhcpv6ServerIapdPoolInfo
	HasAdvertisedPrefixLen() bool
}

// The first IP address of the prefix pool.
// StartPrefixAddress returns a string
func (obj *dhcpv6ServerIapdPoolInfo) StartPrefixAddress() string {

	return *obj.obj.StartPrefixAddress

}

// The first IP address of the prefix pool.
// StartPrefixAddress returns a string
func (obj *dhcpv6ServerIapdPoolInfo) HasStartPrefixAddress() bool {
	return obj.obj.StartPrefixAddress != nil
}

// The first IP address of the prefix pool.
// SetStartPrefixAddress sets the string value in the Dhcpv6ServerIapdPoolInfo object
func (obj *dhcpv6ServerIapdPoolInfo) SetStartPrefixAddress(value string) Dhcpv6ServerIapdPoolInfo {

	obj.obj.StartPrefixAddress = &value
	return obj
}

// The configured_prefix_len ( in conjunction with the prefix_step) can be used to increment the IPv6 lease addresses  to be assigned to the requested clients when multiple addresses are configured by using the size field in the pool.  e.g. This can be used to assign multiple IPv6 host addresses within the same IPv6 subnet ( defined by advertised_prefix_len )  to multiple requesting clients.
// ConfiguredPrefixLen returns a uint32
func (obj *dhcpv6ServerIapdPoolInfo) ConfiguredPrefixLen() uint32 {

	return *obj.obj.ConfiguredPrefixLen

}

// The configured_prefix_len ( in conjunction with the prefix_step) can be used to increment the IPv6 lease addresses  to be assigned to the requested clients when multiple addresses are configured by using the size field in the pool.  e.g. This can be used to assign multiple IPv6 host addresses within the same IPv6 subnet ( defined by advertised_prefix_len )  to multiple requesting clients.
// ConfiguredPrefixLen returns a uint32
func (obj *dhcpv6ServerIapdPoolInfo) HasConfiguredPrefixLen() bool {
	return obj.obj.ConfiguredPrefixLen != nil
}

// The configured_prefix_len ( in conjunction with the prefix_step) can be used to increment the IPv6 lease addresses  to be assigned to the requested clients when multiple addresses are configured by using the size field in the pool.  e.g. This can be used to assign multiple IPv6 host addresses within the same IPv6 subnet ( defined by advertised_prefix_len )  to multiple requesting clients.
// SetConfiguredPrefixLen sets the uint32 value in the Dhcpv6ServerIapdPoolInfo object
func (obj *dhcpv6ServerIapdPoolInfo) SetConfiguredPrefixLen(value uint32) Dhcpv6ServerIapdPoolInfo {

	obj.obj.ConfiguredPrefixLen = &value
	return obj
}

// The total number of addresses in the pool.
// PrefixSize returns a uint32
func (obj *dhcpv6ServerIapdPoolInfo) PrefixSize() uint32 {

	return *obj.obj.PrefixSize

}

// The total number of addresses in the pool.
// PrefixSize returns a uint32
func (obj *dhcpv6ServerIapdPoolInfo) HasPrefixSize() bool {
	return obj.obj.PrefixSize != nil
}

// The total number of addresses in the pool.
// SetPrefixSize sets the uint32 value in the Dhcpv6ServerIapdPoolInfo object
func (obj *dhcpv6ServerIapdPoolInfo) SetPrefixSize(value uint32) Dhcpv6ServerIapdPoolInfo {

	obj.obj.PrefixSize = &value
	return obj
}

// The increment value for the lease address within the lease pool where multiple addresses are present.  The value of the advertised IPv6 prefixes are incremented according to the configured_prefix_len and prefix_step.
// PrefixStep returns a uint32
func (obj *dhcpv6ServerIapdPoolInfo) PrefixStep() uint32 {

	return *obj.obj.PrefixStep

}

// The increment value for the lease address within the lease pool where multiple addresses are present.  The value of the advertised IPv6 prefixes are incremented according to the configured_prefix_len and prefix_step.
// PrefixStep returns a uint32
func (obj *dhcpv6ServerIapdPoolInfo) HasPrefixStep() bool {
	return obj.obj.PrefixStep != nil
}

// The increment value for the lease address within the lease pool where multiple addresses are present.  The value of the advertised IPv6 prefixes are incremented according to the configured_prefix_len and prefix_step.
// SetPrefixStep sets the uint32 value in the Dhcpv6ServerIapdPoolInfo object
func (obj *dhcpv6ServerIapdPoolInfo) SetPrefixStep(value uint32) Dhcpv6ServerIapdPoolInfo {

	obj.obj.PrefixStep = &value
	return obj
}

// The prefix length of the IPv6 prefix that the Dhcpv6 server offers to the Dhcpv6 client.
// AdvertisedPrefixLen returns a uint32
func (obj *dhcpv6ServerIapdPoolInfo) AdvertisedPrefixLen() uint32 {

	return *obj.obj.AdvertisedPrefixLen

}

// The prefix length of the IPv6 prefix that the Dhcpv6 server offers to the Dhcpv6 client.
// AdvertisedPrefixLen returns a uint32
func (obj *dhcpv6ServerIapdPoolInfo) HasAdvertisedPrefixLen() bool {
	return obj.obj.AdvertisedPrefixLen != nil
}

// The prefix length of the IPv6 prefix that the Dhcpv6 server offers to the Dhcpv6 client.
// SetAdvertisedPrefixLen sets the uint32 value in the Dhcpv6ServerIapdPoolInfo object
func (obj *dhcpv6ServerIapdPoolInfo) SetAdvertisedPrefixLen(value uint32) Dhcpv6ServerIapdPoolInfo {

	obj.obj.AdvertisedPrefixLen = &value
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

	if obj.obj.ConfiguredPrefixLen != nil {

		if *obj.obj.ConfiguredPrefixLen > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Dhcpv6ServerIapdPoolInfo.ConfiguredPrefixLen <= 128 but Got %d", *obj.obj.ConfiguredPrefixLen))
		}

	}

	if obj.obj.PrefixSize != nil {

		if *obj.obj.PrefixSize < 1 || *obj.obj.PrefixSize > 1000000 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= Dhcpv6ServerIapdPoolInfo.PrefixSize <= 1000000 but Got %d", *obj.obj.PrefixSize))
		}

	}

	if obj.obj.PrefixStep != nil {

		if *obj.obj.PrefixStep < 1 || *obj.obj.PrefixStep > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= Dhcpv6ServerIapdPoolInfo.PrefixStep <= 4294967295 but Got %d", *obj.obj.PrefixStep))
		}

	}

	if obj.obj.AdvertisedPrefixLen != nil {

		if *obj.obj.AdvertisedPrefixLen > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Dhcpv6ServerIapdPoolInfo.AdvertisedPrefixLen <= 128 but Got %d", *obj.obj.AdvertisedPrefixLen))
		}

	}

}

func (obj *dhcpv6ServerIapdPoolInfo) setDefault() {
	if obj.obj.ConfiguredPrefixLen == nil {
		obj.SetConfiguredPrefixLen(128)
	}
	if obj.obj.PrefixSize == nil {
		obj.SetPrefixSize(10)
	}
	if obj.obj.PrefixStep == nil {
		obj.SetPrefixStep(1)
	}
	if obj.obj.AdvertisedPrefixLen == nil {
		obj.SetAdvertisedPrefixLen(64)
	}

}
